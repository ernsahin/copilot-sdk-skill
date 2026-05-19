# Copilot SDK Skill

[![skills.sh](https://skills.sh/b/ernsahin/skills)](https://skills.sh/ernsahin/skills/copilot-sdk)
[![CI](https://github.com/ernsahin/skills/actions/workflows/ci.yml/badge.svg)](https://github.com/ernsahin/skills/actions/workflows/ci.yml)

Build Copilot SDK agents safely with verified Go, TypeScript, and Python workflows.

Agent skill for designing, building, reviewing, and improving systems with [`github/copilot-sdk`](https://github.com/github/copilot-sdk).

The skill is intentionally opinionated. It is now a Builder Kit, not only a guardrail document. It combines a concise verified API ledger, workflow playbooks, minimal starters, and eval coverage so Copilot SDK work is treated as a real agent product: runtime host, workflow definitions, agent boundaries, tool contracts, permissions, state, observability, failure behavior, and risk-based validation.

## Install

```bash
npx skills add https://github.com/ernsahin/skills
```

Optional explicit install:

```bash
npx skills add https://github.com/ernsahin/skills --skill copilot-sdk
```

## Use When

Use this skill when working with:

- GitHub Copilot SDK systems.
- Go, TypeScript/Node.js, or Python Copilot SDK implementation.
- Copilot SDK sessions, tools, hooks, MCP servers, custom agents, skills, BYOK/auth, streaming, persistence, or telemetry.
- Agent products such as code reviewers, code patchers, PR workflows, documentation update agents, security review agents, or multi-agent workflow hosts.
- Professional directive rewrites for Copilot SDK projects.

Deep implementation support is scoped to Go, TypeScript/Node.js, and Python. .NET, Java, and Rust are official source-map only in this version unless current source is verified during the task.

## Design Standard

The core rule:

```text
Design the Copilot SDK runtime host as reusable infrastructure and the requested agent behavior as the first workflow, unless the user explicitly asks for a throwaway prototype.
```

Every non-trivial agent workflow must define:

1. Capability ownership.
2. Access policy.
3. Evidence requirements.
4. Action targeting.
5. Observability.
6. State continuity.
7. Failure semantics.

This prevents shallow outputs such as empty SDLC folders, giant prompts with no runtime policy, hardcoded demo answers, or one-off agents that cannot accept a second workflow.

## Package Layout

```text
SKILL.md
skills/copilot-sdk/
  SKILL.md
  references/
    verified-api-ledger.md
    workflows/
  examples/
    starters/
  evals/evals.json
  assets/
docs/
  benchmark-summary.md
  evaluation.md
  release-checklist.md
scripts/
  validate_skill.py
```

## Quality Controls

Run local validation:

```bash
python scripts/validate_skill.py skills/copilot-sdk
```

The validator checks:

- Required skill frontmatter.
- Skill name and directory consistency.
- Description length and trigger usefulness.
- Reference links from `SKILL.md`.
- Verified API ledger for Go, TypeScript, and Python.
- Workflow playbooks and minimal starters.
- Eval JSON schema shape.
- Example freshness labels.
- Prohibited placeholder language.

CI runs the same validation on every push and pull request.

CI also checks that official upstream source links in the API ledger still resolve:

```bash
python scripts/check_upstream_sources.py
```

## Evaluation Status

The skill includes eval prompts covering:

- Copilot SDK code review agents.
- Directive rewriting.
- Production backend architecture.
- Shallow prompt and approve-all anti-patterns.
- Generic skill quality.
- SDLC-as-workflow behavior.
- Hardcoded shortcut rejection.
- Runtime extensibility for future agents.
- Autonomous recovery instead of unnecessary user handoff.
- Source verification for exact SDK API guidance.
- Go, TypeScript, and Python starter-level implementation guidance.
- MCP-backed agent workflows.
- BYOK backend workflows.
- Negative tests for unsupported exact .NET/Rust code without source verification.

See [docs/evaluation.md](docs/evaluation.md) for the baseline vs with-skill runbook and [docs/benchmark-summary.md](docs/benchmark-summary.md) for the current manual benchmark summary.

Behavior eval tooling is included under [eval-harness](eval-harness/). It uses GitHub Copilot SDK itself:

```bash
cd eval-harness
go run . -limit 4
```

The harness runs paired sessions:

1. Baseline: same custom eval agent with no skill preloaded.
2. With skill: same custom eval agent with `Skills: ["copilot-sdk"]` and `SkillDirectories` pointing at this repository's `skills/` directory.

Raw run outputs are written to `eval-results/`, which is ignored by git. The repository does not claim benchmark completion until outputs are manually graded and summarized.

The harness also writes `aggregate-grading.json` so manual grades can be summarized without pretending that keyword checks prove quality.

## Builder Kit Policy

Files under `skills/copilot-sdk/examples/starters/` are minimal starter templates verified against `skills/copilot-sdk/references/verified-api-ledger.md`. They are intended as small implementation starting points, not production applications.

Files directly under `skills/copilot-sdk/examples/` remain shape references, not source verification. For exact SDK imports, types, fields, event names, permission result names, or setup code, inspect current upstream docs, installed SDK source, or the verified API ledger for starter-level Go, TypeScript, and Python guidance.

## Field Test

The skill was tested against a small local Go Copilot SDK CLI, documented in [docs/field-test-copilot-go-chat.md](docs/field-test-copilot-go-chat.md). The test verified that the skill pushed the implementation toward runtime prompts, SDK session lifecycle, permission policy, event observability, timeout behavior, and policy tests instead of a hardcoded demo.

## Source Policy

The Copilot SDK is public preview. The skill requires upstream verification before version-sensitive API guidance.

Primary sources:

- [github/copilot-sdk](https://github.com/github/copilot-sdk)
- [Copilot SDK docs index](https://raw.githubusercontent.com/github/copilot-sdk/main/docs/index.md)
- Target language source, especially `go/`, `nodejs/`, `python/`, `dotnet/`, `java/`, and `rust/`

The repository root `SKILL.md` is the install entrypoint for `npx skills add owner/repo`. The canonical skill body is `skills/copilot-sdk/SKILL.md`. Any sibling or local copied `copilot-sdk/` directory outside this repository should be treated as a stale workspace copy unless explicitly synchronized.

## License

MIT. See [LICENSE](LICENSE).
