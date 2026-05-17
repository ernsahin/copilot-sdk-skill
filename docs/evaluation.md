# Evaluation Runbook

The skill includes eval prompts in `skills/copilot-sdk/evals/evals.json`.

The goal is to verify that the skill changes agent behavior, not just that the files are valid.

## Current Status

- Eval prompts: defined.
- Objective expectations: defined.
- Baseline vs with-skill runner: available at `eval-harness/`.
- Baseline vs with-skill manual summary: available in `docs/benchmark-summary.md`.
- CI smoke validation: enabled through `scripts/validate_skill.py`.
- Field test: documented in `docs/field-test-copilot-go-chat.md`.

The current benchmark status is intentionally conservative. Local paired runs show behavior lift, but the repository should not claim top-tier benchmark completion until the expanded eval set is run repeatedly and graded consistently.

## Run Paired Evals With Copilot SDK

The behavior harness uses GitHub Copilot SDK itself. This matches the domain of the skill and avoids testing through unrelated Claude/Codex CLI behavior.

```bash
cd eval-harness
go run . -limit 4
```

Run targeted evals while iterating on a specific behavior:

```bash
cd eval-harness
go run . -ids 1,10 -limit 0
```

Use `-force-skill` when the goal is to evaluate the skill instructions themselves rather than automatic skill selection:

```bash
cd eval-harness
go run . -limit 4 -force-skill
```

This leaves the baseline prompt unchanged and only prefixes the with-skill prompt with a short instruction to use the `copilot-sdk` skill. It is usually unnecessary for content-quality evals because the with-skill custom agent already preloads the skill.

For each eval, the harness creates:

1. A baseline Copilot SDK session using a custom eval agent with no skills preloaded.
2. A with-skill Copilot SDK session using the same custom eval agent with `Skills: ["copilot-sdk"]` and `SkillDirectories` pointing to this repository's `skills/` directory.

By default both sessions run in an isolated workspace outside the skill repository. This prevents the baseline session from reading the skill files from the repository and contaminating the comparison. Pass `-workdir` only when an eval intentionally needs a real repository as input.

This follows the SDK behavior verified from the Go SDK source and e2e tests: `SkillDirectories` discovers skills, and the custom agent `Skills` field preloads named skill content into that agent's context.

```text
eval-results/<timestamp>/
  eval-01/
    baseline/
      response.md
      events.jsonl
      permissions.jsonl
      run.json
    with-skill/
      response.md
      events.jsonl
      permissions.jsonl
      run.json
    grading.json
```

`eval-results/` is ignored by git. Keep raw outputs local and summarize meaningful results in `docs/benchmark-summary.md`.

The harness uses a conservative permission policy:

- Read requests are approved.
- Shell, write, URL, MCP, memory, custom-tool, and hook requests are denied.

The harness does not auto-grade the outputs. Grading is intentionally explicit so the repository does not pretend quality was proven by keyword matching.

## Evaluation Standard

For each eval:

1. Run the prompt without this skill.
2. Run the same prompt with `skills/copilot-sdk`.
3. Save both outputs.
4. Grade each expectation as pass/fail with evidence.
5. Compare whether the skill materially improves the output.
6. Patch the skill only when the failure reveals a reusable rule or reference gap.

## Expected Improvements

The with-skill output should:

- Reject hardcoded demo behavior when a reusable workflow is needed.
- Treat SDLC as executable workflow governance, not folder scaffolding.
- Separate runtime host, workflows, agents, tools, permissions, state, integrations, and observability.
- Require upstream verification before version-sensitive Copilot SDK guidance.
- Use risk-based validation instead of requiring tests everywhere.
- Prefer SDK and platform primitives before custom infrastructure.

## Result Format

Local raw outputs are stored under:

```text
eval-results/
  <timestamp>/
    eval-01/
      prompt.txt
      baseline/
        response.md
        events.jsonl
        permissions.jsonl
        run.json
      with-skill/
        response.md
        events.jsonl
        permissions.jsonl
        run.json
      grading.json
    summary.json
```

Do not commit raw result directories. Commit only a concise summary when the result changes release confidence.

Use this grading shape:

```json
{
  "eval_id": 1,
  "baseline_passed": 1,
  "with_skill_passed": 4,
  "expectations": [
    {
      "text": "The output rejects hardcoded outcomes when future inputs are implied.",
      "baseline": false,
      "with_skill": true,
      "evidence": "The with-skill output explicitly chose a reusable workflow mechanism."
    }
  ],
  "notes": "Patch the skill only if with-skill misses a recurring standard."
}
```

## Publishing Gate

Before a public release:

1. At least four evals must have baseline and with-skill outputs.
2. With-skill output should pass more expectations than baseline on each run.
3. Any failed expectation must be classified as one of:
   - Skill wording gap.
   - Eval expectation too broad.
   - User prompt intentionally out of scope.
   - Upstream SDK ambiguity.
4. The release notes must mention evaluation status honestly.
