# Changelog

All notable changes to this skill are documented here.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project uses semantic versioning for skill content.

## Unreleased

### Changed

- Renamed the installable skill package to `copilot-sdk-kit`.
- Removed the repository-root `SKILL.md` entrypoint so the repository exposes a single skill package under `skills/copilot-sdk-kit/`.

## [0.5.4] - 2026-05-20

### Changed

- Made the repository-root `SKILL.md` a concise, self-contained audited entrypoint instead of a delegation-only bootstrap.
- Kept the named skill package for agents that load skill directories.
- Clarified source verification, workflow quality gates, routing, and completion standards in the root entrypoint to reduce audit ambiguity.

## [0.5.0] - 2026-05-18

### Added

- Builder Kit positioning for Go, TypeScript/Node.js, and Python Copilot SDK work.
- Concise verified API ledger for Go, TypeScript, and Python with checked upstream URLs.
- Repository-root skill entrypoint for `npx skills add owner/repo` compatibility.
- Workflow playbooks for code review agents, code patchers, MCP-backed agents, BYOK backends, and skill-loaded custom agents.
- Minimal starter templates for Go, TypeScript, and Python.
- Upstream source link checker for CI.
- Expanded eval coverage for starter implementation, MCP, BYOK, custom-agent skills, and unsupported exact .NET/Rust code.
- Eval harness aggregate grading output.

### Changed

- Clarified that .NET, Java, and Rust are source-map only unless current source is verified during the task.
- Updated release gates to require source link checks, starters, workflow playbooks, and stronger benchmark evidence before top-tier claims.

## [0.4.0] - 2026-05-18

### Added

- Repo-ready package structure for `skills.sh` installation.
- CI validation workflow.
- Local validator for frontmatter, references, evals, examples, and placeholders.
- Evaluation runbook for baseline vs with-skill comparisons.
- Release checklist for publishing and update quality.
- Copilot SDK eval harness that runs paired baseline and with-skill sessions through the Go SDK.
- Benchmark summary documenting the current manual read and remaining evaluation limits.
- Field-test notes for the local Go Copilot SDK CLI prototype.
- High-impact Copilot SDK rule references for permissions, state ownership, side-effect targeting, observability, native primitives, and extensibility.

### Changed

- Strengthened runtime extensibility standard for multi-agent Copilot SDK products.
- Clarified that the first requested agent is the first workflow on a reusable runtime host, not the whole architecture.
- Tightened source-verification rules so exact SDK APIs are not presented unless current source or official docs were inspected during the task.
- Reclassified example files as shape references rather than verified SDK setup code.
- Expanded eval coverage for autonomous recovery and source-sensitive SDK guidance.

## [0.3.0] - 2026-05-17

### Added

- Agent workflow quality gates: ownership, access, evidence, targeting, observability, state continuity, and failure semantics.
- Enforcement protocol to reject hardcoded outcomes when a reusable mechanism is required.
- Eval coverage for SDLC workflow behavior, hardcoded shortcut rejection, and runtime extensibility.

## [0.2.0] - 2026-05-17

### Added

- Source verification reference.
- Known Copilot SDK gotchas reference.
- Prompt design reference.
- Skill feedback workflow.

## [0.1.0] - 2026-05-17

### Added

- Initial Copilot SDK skill draft.
- Go-first references, examples, and initial eval prompts.
