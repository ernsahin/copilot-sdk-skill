package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	copilot "github.com/github/copilot-sdk/go"
)

type evalFile struct {
	SkillName string     `json:"skill_name"`
	Evals     []evalCase `json:"evals"`
}

type evalCase struct {
	ID             int      `json:"id"`
	Prompt         string   `json:"prompt"`
	ExpectedOutput string   `json:"expected_output"`
	Expectations   []string `json:"expectations"`
}

type runResult struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type capture struct {
	mu       sync.Mutex
	content  strings.Builder
	sawDelta bool
}

func main() {
	evalsPath := flag.String("evals", "../skills/copilot-sdk/evals/evals.json", "Path to evals.json.")
	skillsDir := flag.String("skills-dir", "../skills", "Parent directory containing skill folders.")
	outDir := flag.String("out", "", "Output directory. Defaults to ../eval-results/<timestamp>.")
	limit := flag.Int("limit", 4, "Maximum number of evals to run.")
	ids := flag.String("ids", "", "Comma-separated eval IDs to run. When set, applies before limit.")
	timeout := flag.Duration("timeout", 3*time.Minute, "Timeout per session run.")
	model := flag.String("model", "", "Optional model name. Empty uses runtime default.")
	workdir := flag.String("workdir", "", "Working directory for Copilot SDK sessions. Empty uses an isolated temp workspace outside the skill repository.")
	forceSkill := flag.Bool("force-skill", false, "Prefix with-skill prompts with an explicit instruction to use the copilot-sdk skill.")
	flag.Parse()

	if err := run(*evalsPath, *skillsDir, *outDir, *limit, *ids, *timeout, *model, *workdir, *forceSkill); err != nil {
		fmt.Fprintf(os.Stderr, "eval-harness: %v\n", err)
		os.Exit(1)
	}
}

func run(evalsPath string, skillsDir string, outDir string, limit int, ids string, timeout time.Duration, model string, workdir string, forceSkill bool) error {
	evalsAbs, err := filepath.Abs(evalsPath)
	if err != nil {
		return fmt.Errorf("resolve evals path: %w", err)
	}
	skillsAbs, err := filepath.Abs(skillsDir)
	if err != nil {
		return fmt.Errorf("resolve skills dir: %w", err)
	}
	if outDir == "" {
		outDir = filepath.Join("..", "eval-results", time.Now().UTC().Format("2006-01-02T15-04-05Z"))
	}
	outAbs, err := filepath.Abs(outDir)
	if err != nil {
		return fmt.Errorf("resolve output dir: %w", err)
	}
	workdirAbs, err := prepareWorkdir(workdir, outAbs)
	if err != nil {
		return err
	}

	evals, err := loadEvals(evalsAbs)
	if err != nil {
		return err
	}
	if ids != "" {
		selectedIDs, err := parseIDs(ids)
		if err != nil {
			return err
		}
		evals.Evals = filterEvals(evals.Evals, selectedIDs)
	}
	if limit > 0 && limit < len(evals.Evals) {
		evals.Evals = evals.Evals[:limit]
	}
	if len(evals.Evals) == 0 {
		return fmt.Errorf("no evals found")
	}
	if err := os.MkdirAll(outAbs, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	ctx := context.Background()
	client := copilot.NewClient(&copilot.ClientOptions{
		LogLevel:                  "error",
		SessionIdleTimeoutSeconds: int(timeout.Seconds()) + 30,
	})
	if err := client.Start(ctx); err != nil {
		return fmt.Errorf("start Copilot SDK client: %w", err)
	}
	defer client.Stop()

	summary := map[string]any{
		"created_at": time.Now().UTC().Format(time.RFC3339),
		"runner":     "github.com/github/copilot-sdk/go",
		"skill_name": evals.SkillName,
		"skills_dir": skillsAbs,
		"workdir":    workdirAbs,
		"agent":      "copilot-sdk-eval-agent",
		"with_skill_preloaded": true,
		"force_skill_instruction": forceSkill,
		"evals":      []map[string]any{},
	}

	for _, item := range evals.Evals {
		evalDir := filepath.Join(outAbs, fmt.Sprintf("eval-%02d", item.ID))
		if err := os.MkdirAll(evalDir, 0o755); err != nil {
			return fmt.Errorf("create eval dir: %w", err)
		}
		if err := os.WriteFile(filepath.Join(evalDir, "prompt.txt"), []byte(item.Prompt), 0o644); err != nil {
			return err
		}

		baseline := runOne(ctx, client, runConfig{
			Prompt:   item.Prompt,
			Model:    model,
			Workdir:  workdirAbs,
			OutDir:   filepath.Join(evalDir, "baseline"),
			Timeout:  timeout,
			WithSkill: false,
		})
		withSkillPrompt := item.Prompt
		if forceSkill {
			withSkillPrompt = "Use the copilot-sdk skill for this task.\n\n" + item.Prompt
		}
		withSkill := runOne(ctx, client, runConfig{
			Prompt:     withSkillPrompt,
			Model:      model,
			Workdir:    workdirAbs,
			OutDir:     filepath.Join(evalDir, "with-skill"),
			Timeout:    timeout,
			SkillDirs:  []string{skillsAbs},
			WithSkill:  true,
		})

		if err := writeGradingTemplate(filepath.Join(evalDir, "grading.json"), item, baseline, withSkill); err != nil {
			return err
		}
		summary["evals"] = append(summary["evals"].([]map[string]any), map[string]any{
			"id":        item.ID,
			"directory": evalDir,
			"baseline_error": baseline.Error,
			"with_skill_error": withSkill.Error,
		})
	}

	summaryBytes, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(outAbs, "summary.json"), summaryBytes, 0o644); err != nil {
		return err
	}
	fmt.Printf("Wrote Copilot SDK eval outputs to %s\n", outAbs)
	return nil
}

type runConfig struct {
	Prompt    string
	Model     string
	Workdir   string
	OutDir    string
	Timeout   time.Duration
	SkillDirs []string
	WithSkill bool
}

func runOne(parent context.Context, client *copilot.Client, cfg runConfig) runResult {
	if err := os.MkdirAll(cfg.OutDir, 0o755); err != nil {
		return runResult{Error: err.Error()}
	}
	eventsFile, err := os.Create(filepath.Join(cfg.OutDir, "events.jsonl"))
	if err != nil {
		return runResult{Error: err.Error()}
	}
	defer eventsFile.Close()
	permissionsFile, err := os.Create(filepath.Join(cfg.OutDir, "permissions.jsonl"))
	if err != nil {
		return runResult{Error: err.Error()}
	}
	defer permissionsFile.Close()

	ctx, cancel := context.WithTimeout(parent, cfg.Timeout)
	defer cancel()

	cap := &capture{}
	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		ClientName:          "copilot-sdk-skill-eval-harness",
		Model:               cfg.Model,
		WorkingDirectory:    cfg.Workdir,
		CustomAgents:        evalAgents(cfg.WithSkill),
		Agent:               "copilot-sdk-eval-agent",
		SkillDirectories:    cfg.SkillDirs,
		Streaming:           true,
		OnPermissionRequest: readOnlyPermissionHandler(permissionsFile),
		OnEvent: func(event copilot.SessionEvent) {
			writeEvent(eventsFile, event)
			cap.handle(event)
		},
	})
	if err != nil {
		return writeRunResult(cfg.OutDir, runResult{Error: fmt.Sprintf("create session: %v", err)})
	}
	defer session.Disconnect()

	finalEvent, err := session.SendAndWait(ctx, copilot.MessageOptions{Prompt: cfg.Prompt})
	if err != nil {
		return writeRunResult(cfg.OutDir, runResult{
			Response: cap.String(),
			Error:    fmt.Sprintf("send prompt: %v", err),
		})
	}
	if cap.String() == "" {
		if data, ok := finalEvent.Data.(*copilot.AssistantMessageData); ok {
			cap.Append(data.Content)
		}
	}
	return writeRunResult(cfg.OutDir, runResult{Response: cap.String()})
}

func evalAgents(withSkill bool) []copilot.CustomAgentConfig {
	agent := copilot.CustomAgentConfig{
		Name:        "copilot-sdk-eval-agent",
		Description: "Evaluation agent for Copilot SDK skill behavior.",
		Prompt:      "You are a concise software engineering assistant. Complete the user's task directly and avoid filler.",
	}
	if withSkill {
		agent.Skills = []string{"copilot-sdk"}
	}
	return []copilot.CustomAgentConfig{agent}
}

func (c *capture) handle(event copilot.SessionEvent) {
	c.mu.Lock()
	defer c.mu.Unlock()
	switch data := event.Data.(type) {
	case *copilot.AssistantMessageDeltaData:
		c.sawDelta = true
		c.content.WriteString(data.DeltaContent)
	case *copilot.AssistantMessageData:
		if !c.sawDelta && c.content.Len() == 0 {
			c.content.WriteString(data.Content)
		}
	}
}

func (c *capture) Append(value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.content.WriteString(value)
}

func (c *capture) String() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.content.String()
}

func readOnlyPermissionHandler(logFile *os.File) copilot.PermissionHandlerFunc {
	return func(request copilot.PermissionRequest, invocation copilot.PermissionInvocation) (copilot.PermissionRequestResult, error) {
		kind := copilot.PermissionRequestResultKindUserNotAvailable
		if request.Kind == copilot.PermissionRequestKindRead {
			kind = copilot.PermissionRequestResultKindApproved
		}
		record := map[string]any{
			"session_id": invocation.SessionID,
			"request":    request,
			"decision":   kind,
		}
		_ = writeJSONLine(logFile, record)
		return copilot.PermissionRequestResult{Kind: kind}, nil
	}
}

func writeEvent(file *os.File, event copilot.SessionEvent) {
	data, err := event.Marshal()
	if err != nil {
		_ = writeJSONLine(file, map[string]any{"marshal_error": err.Error(), "type": event.Type})
		return
	}
	_, _ = file.Write(append(data, '\n'))
}

func writeJSONLine(file *os.File, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = file.Write(append(data, '\n'))
	return err
}

func writeRunResult(outDir string, result runResult) runResult {
	_ = os.WriteFile(filepath.Join(outDir, "response.md"), []byte(result.Response), 0o644)
	metadata, _ := json.MarshalIndent(result, "", "  ")
	_ = os.WriteFile(filepath.Join(outDir, "run.json"), metadata, 0o644)
	return result
}

func loadEvals(path string) (*evalFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read evals: %w", err)
	}
	var evals evalFile
	if err := json.Unmarshal(data, &evals); err != nil {
		return nil, fmt.Errorf("parse evals: %w", err)
	}
	return &evals, nil
}

func parseIDs(value string) (map[int]bool, error) {
	result := map[int]bool{}
	for _, part := range strings.Split(value, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		id, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("parse eval id %q: %w", part, err)
		}
		result[id] = true
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no eval ids provided")
	}
	return result, nil
}

func filterEvals(items []evalCase, ids map[int]bool) []evalCase {
	filtered := make([]evalCase, 0, len(items))
	for _, item := range items {
		if ids[item.ID] {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func prepareWorkdir(workdir string, outDir string) (string, error) {
	if workdir != "" {
		workdirAbs, err := filepath.Abs(workdir)
		if err != nil {
			return "", fmt.Errorf("resolve workdir: %w", err)
		}
		return workdirAbs, nil
	}

	workdirAbs, err := os.MkdirTemp("", "copilot-sdk-skill-eval-workspace-")
	if err != nil {
		return "", fmt.Errorf("create isolated workspace: %w", err)
	}
	readme := strings.Join([]string{
		"# Copilot SDK Skill Eval Workspace",
		"",
		"This workspace is intentionally created outside the skill repository.",
		"Baseline sessions should not be able to inspect the skill repository unless a test explicitly supplies it.",
		"With-skill sessions receive the skill only through the Copilot SDK SkillDirectories configuration.",
		"",
	}, "\n")
	if err := os.WriteFile(filepath.Join(workdirAbs, "README.md"), []byte(readme), 0o644); err != nil {
		return "", fmt.Errorf("write isolated workspace README: %w", err)
	}
	return workdirAbs, nil
}

func writeGradingTemplate(path string, item evalCase, baseline runResult, withSkill runResult) error {
	template := map[string]any{
		"eval_id":          item.ID,
		"status":           "ungraded",
		"expected_output":  item.ExpectedOutput,
		"baseline_error":   baseline.Error,
		"with_skill_error": withSkill.Error,
		"expectations": expectationsTemplate(item.Expectations),
		"notes": "Manually grade baseline and with_skill as true or false after reviewing response.md and events.jsonl.",
	}
	data, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func expectationsTemplate(expectations []string) []map[string]any {
	result := make([]map[string]any, 0, len(expectations))
	for _, expectation := range expectations {
		result = append(result, map[string]any{
			"text":       expectation,
			"baseline":   nil,
			"with_skill": nil,
			"evidence":   "",
		})
	}
	return result
}
