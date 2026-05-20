---
name: copilot-sdk
description: "Build, plan, review, and improve production-shaped systems using github/copilot-sdk. Use this skill whenever the user mentions Copilot SDK, GitHub Copilot CLI SDKs, embedded Copilot agents, SDK sessions, Go, TypeScript/Node.js, Python, custom tools, custom agents, MCP servers, hooks, BYOK/auth, streaming, persistence, telemetry, SDK-loaded skills, code review agents, code patching agents, PR automation, or professional directives for Copilot SDK systems. Treat this as a Builder Kit: route to the verified API ledger, workflow playbooks, and minimal starters before inventing code. Enforce reusable runtime design, explicit tool and permission boundaries, observable session behavior, failure and continuation semantics, state ownership, source verification, and risk-based validation. Do not provide exact SDK imports, types, fields, event names, permission result names, or setup code unless current source, installed package source, or the verified API ledger was inspected during the task."
license: MIT
metadata:
  author: local
  version: "0.5.0"
---

# Copilot SDK Builder Kit

Use this skill to design, build, review, and improve production-shaped systems that use `github/copilot-sdk`.

Treat Copilot SDK work as agent product engineering, not prompt-only scaffolding. Every non-trivial answer should address runtime ownership, tools, permissions, evidence, failure behavior, observability, state, and validation.

This root `SKILL.md` is the audited install entrypoint for `npx skills add ernsahin/skills`. The named skill package also exists at `skills/copilot-sdk/SKILL.md` for agents that load skill directories.

## Core Rules

The Copilot SDK is version-sensitive. Exact API guidance is unsafe unless verified.

1. Do not provide exact imports, package paths, structs, fields, methods, event names, permission result names, or setup code unless current SDK source, installed package source, official docs, or `skills/copilot-sdk/references/verified-api-ledger.md` were inspected during the task.
2. If source verification is unavailable, state `Source status: not verified`, provide conceptual guidance only, and name the source path to check.
3. For Go, TypeScript/Node.js, or Python implementation guidance, inspect the verified API ledger first. Installed project source or lockfiles outrank the ledger.
4. For .NET, Java, or Rust, give conceptual guidance and official lookup paths unless current source was inspected during the task.
5. Do not use broad prompts as substitutes for runtime controls, typed tool contracts, permission policy, persistence, or audit logs.

## Workflow Standard

Before planning or implementing, establish:

1. User outcome and source of truth.
2. Runtime host, session lifecycle, and integration boundary.
3. Tool ownership, permission policy, and trust boundary.
4. Evidence required before reads, writes, or external side effects.
5. Failure, retry, cancellation, duplicate prevention, and resume behavior.
6. Observability through events, logs, skipped scope, confidence, and audit records.
7. Risk-based validation.

If a non-trivial design leaves these unresolved, mark the gap as a prototype limit instead of claiming production readiness.

## Routing

Read only the reference needed for the task:

1. Exact Go, TypeScript, or Python APIs: `skills/copilot-sdk/references/verified-api-ledger.md`.
2. Go implementation details: `skills/copilot-sdk/references/go.md`.
3. Code reviewer agents: `skills/copilot-sdk/references/workflows/code-reviewer-workflow.md`.
4. Code patcher agents: `skills/copilot-sdk/references/workflows/code-patcher-workflow.md`.
5. MCP-backed agents: `skills/copilot-sdk/references/workflows/mcp-backed-agent-workflow.md`.
6. BYOK or hosted backends: `skills/copilot-sdk/references/workflows/byok-backend-workflow.md`.
7. Skill-loaded custom agents or sub-agents: `skills/copilot-sdk/references/workflows/skill-loaded-custom-agent-workflow.md`.
8. Broader product design: `skills/copilot-sdk/references/agent-product-lifecycle.md`.
9. Strict completion checks: `skills/copilot-sdk/references/stop-conditions.md`.

## Response Discipline

Produce the artifact the user asked for. Do not end with a generic permission question when the requested output can be completed now.

For design requests, lead with a concise minimum architecture. Use these headings only when helpful: `Assumptions`, `Runtime`, `Workflow`, `Capabilities`, `Evidence`, `Failure`, `Observability`, and `Validation`.

Continue without asking when the next step is recoverable, locally verifiable, or implied by the user's goal. Ask only when progress requires a product decision, credentials, destructive action, unavailable external access, irreversible side effect, or a materially different user outcome.

For reviews, lead with risks and fixes. For directive rewrites, output the rewritten directive only.

## Completion Standard

Before finalizing, check that the response does not:

1. Solve only the happy path.
2. Move complexity without naming ownership.
3. Create duplicate sources of truth.
4. Use prompt text instead of runtime controls.
5. Ignore failure, recovery, observability, or validation.
6. Claim production readiness without source verification and risk-based checks.

When the request is vague, turn it into clear system requirements before building. Before inventing custom infrastructure, inspect whether the SDK, host language, repository, or execution environment already provides the needed primitive.
