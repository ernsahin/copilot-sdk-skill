# Copilot SDK Skill

[![skills.sh](https://skills.sh/b/ernsahin/copilot-sdk-skill)](https://skills.sh/ernsahin/copilot-sdk-skill)
[![CI](https://github.com/ernsahin/copilot-sdk-skill/actions/workflows/ci.yml/badge.svg)](https://github.com/ernsahin/copilot-sdk-skill/actions/workflows/ci.yml)

Agent skill for designing, building, reviewing, and improving systems with [`github/copilot-sdk`](https://github.com/github/copilot-sdk).

The skill is intentionally opinionated. It is not a code-snippet collection. It forces Copilot SDK work to be treated as a real agent product: runtime host, workflow definitions, agent boundaries, tool contracts, permissions, state, observability, failure behavior, and risk-based validation.

## Install

```bash
npx skills add https://github.com/ernsahin/copilot-sdk-skill --skill copilot-sdk
```

Or, if the repository is discoverable by the Skills CLI:

```bash
npx skills add ernsahin/copilot-sdk-skill --skill copilot-sdk
```

## Use When

Use this skill when working with:

- GitHub Copilot SDK systems.
- Copilot SDK sessions, tools, hooks, MCP servers, custom agents, skills, BYOK/auth, streaming, persistence, or telemetry.
- Agent products such as code reviewers, code patchers, PR workflows, documentation update agents, security review agents, or multi-agent workflow hosts.
- Professional directive rewrites for Copilot SDK projects.

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
skills/copilot-sdk/
  SKILL.md
  references/
  examples/
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
- Eval JSON schema shape.
- Example freshness labels.
- Prohibited placeholder language.

CI runs the same validation on every push and pull request.

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

## Examples Policy

Files under `skills/copilot-sdk/examples/` are shape references, not source verification. They show the kind of runtime boundary to design, but they must not be copied as final SDK API code.

For exact SDK imports, types, fields, event names, permission result names, or setup code, inspect current upstream docs or the installed SDK source for the target version.

## Field Test

The skill was tested against a small local Go Copilot SDK CLI, documented in [docs/field-test-copilot-go-chat.md](docs/field-test-copilot-go-chat.md). The test verified that the skill pushed the implementation toward runtime prompts, SDK session lifecycle, permission policy, event observability, timeout behavior, and policy tests instead of a hardcoded demo.

## Source Policy

The Copilot SDK is public preview. The skill requires upstream verification before version-sensitive API guidance.

Primary sources:

- [github/copilot-sdk](https://github.com/github/copilot-sdk)
- [Copilot SDK docs index](https://raw.githubusercontent.com/github/copilot-sdk/main/docs/index.md)
- Target language source, especially `go/`, `nodejs/`, `python/`, `dotnet/`, `java/`, and `rust/`

## License

MIT. See [LICENSE](LICENSE).
