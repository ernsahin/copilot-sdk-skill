# ernsahin Skills

[![skills.sh](https://skills.sh/b/ernsahin/skills)](https://skills.sh/ernsahin/skills)
[![CI](https://github.com/ernsahin/skills/actions/workflows/ci.yml/badge.svg)](https://github.com/ernsahin/skills/actions/workflows/ci.yml)

Agent skills maintained by Ernsahin.

This repository is a skill collection. Each installable skill lives under `skills/<skill-name>/` with its own `SKILL.md`, references, examples, evals, and assets as needed. The repository intentionally does not keep a root `SKILL.md`, so `skills.sh` and skill installers should treat the packages under `skills/` as the source of truth.

## Skills

| Skill | Purpose | Install |
| --- | --- | --- |
| `copilot-sdk-kit` | Build, plan, review, and improve production-shaped systems using `github/copilot-sdk`, with verified Go, TypeScript/Node.js, and Python workflows. | `npx skills add https://github.com/ernsahin/skills --skill copilot-sdk-kit` |

## Layout

```text
skills/
  copilot-sdk-kit/
    SKILL.md
    references/
    examples/
    evals/
    assets/
docs/
  benchmark-summary.md
  evaluation.md
  release-checklist.md
scripts/
  validate_skill.py
  check_upstream_sources.py
eval-harness/
```

## Quality Controls

Validate the current skill package:

```bash
python3 scripts/validate_skill.py skills/copilot-sdk-kit
python3 scripts/check_upstream_sources.py
cd eval-harness && go test ./...
```

CI runs the same validation on push and pull request.

## Current Skill

`copilot-sdk-kit` is the first skill in this collection. It is a Builder Kit for Copilot SDK systems, not only a guardrail document. It combines a verified API ledger, workflow playbooks, minimal starters, and eval coverage so Copilot SDK work is treated as a durable agent product: runtime host, workflow definitions, agent boundaries, tool contracts, permissions, state, observability, failure behavior, and risk-based validation.

Deep implementation support is scoped to Go, TypeScript/Node.js, and Python. .NET, Java, and Rust are official source-map only unless current source is verified during the task.

## Evaluation

Behavior eval tooling is included under `eval-harness/`. It runs paired sessions:

1. Baseline: same custom eval agent with no skill preloaded.
2. With skill: same custom eval agent with `Skills: ["copilot-sdk-kit"]` and `SkillDirectories` pointing at this repository's `skills/` directory.

Raw run outputs are written to `eval-results/`, which is ignored by git. See [docs/evaluation.md](docs/evaluation.md) and [docs/benchmark-summary.md](docs/benchmark-summary.md).

## License

MIT. See [LICENSE](LICENSE).
