# Copilot SDK Eval Harness

This harness evaluates the `copilot-sdk` skill using GitHub Copilot SDK itself.

For each eval prompt, it runs two isolated Copilot SDK sessions:

1. **Baseline**: the same custom eval agent with no skills preloaded.
2. **With skill**: the same custom eval agent with `Skills: ["copilot-sdk"]` and `SkillDirectories` pointing at the repository `skills/` directory.

By default, both sessions use a clean workspace created outside the skill repository. This keeps the baseline from reading the skill repository by accident. Use `-workdir` only when an eval intentionally needs a real repository.

The custom agent is intentional. In the Copilot SDK, `SkillDirectories` discovers skills, while a custom agent's `Skills` field preloads named skill content into that agent's context. This harness keeps the agent prompt identical between baseline and with-skill so the measured difference is the skill.

It saves:

- `baseline/response.md`
- `baseline/events.jsonl`
- `baseline/permissions.jsonl`
- `with-skill/response.md`
- `with-skill/events.jsonl`
- `with-skill/permissions.jsonl`
- `grading.json`

The harness does not auto-grade. Manual grading is intentional because these evals check architectural quality, not simple keyword presence.

## Run

From this directory:

```bash
go run . -limit 4
```

Run specific evals:

```bash
go run . -ids 1,10 -limit 0
```

To test the skill content separately from automatic skill-selection behavior, run:

```bash
go run . -limit 4 -force-skill
```

This keeps the baseline unchanged and adds a short "use the copilot-sdk skill" instruction only to the with-skill prompts. It is usually unnecessary because the with-skill eval agent already preloads the skill.

Windows PowerShell:

```powershell
go run . -limit 4
```

Optional:

```bash
go run . \
  -evals ../skills/copilot-sdk/evals/evals.json \
  -skills-dir ../skills \
  -out ../eval-results/manual-run \
  -ids 1,10 \
  -limit 4 \
  -force-skill \
  -timeout 3m
```

## Requirements

- Go.
- GitHub Copilot SDK runtime access.
- GitHub Copilot CLI authentication or another authentication method supported by the SDK runtime.

## Policy

The harness uses a conservative permission policy:

- Read requests are approved.
- Shell, write, URL, MCP, memory, custom-tool, and hook permissions are denied.

This keeps evals focused on planning and architecture quality rather than allowing side effects.
