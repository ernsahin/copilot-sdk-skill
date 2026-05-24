---
title: Copilot SDK Workflow Routing
impact: HIGH
impactDescription: Routes user intent to concrete Copilot SDK workflows before architecture or code generation
tags: routing, workflow, orchestration, agents
---

# Copilot SDK Workflow Routing

Use this before designing or implementing. Pick one primary workflow and only add secondary workflows when the user explicitly asks or the system cannot work without them.

## Route Table

| User intent | Primary workflow | Required outcome |
| --- | --- | --- |
| Create an agent product | Agent product lifecycle | Runtime host plus first executable workflow |
| Review code or PRs | Review workflow | Intake, evidence, findings, permissions, state, observability |
| Patch code | Patch workflow | Proposed changes, approval points, rollback/failure behavior |
| Add docs updater | Content workflow | Source evidence, target files, write policy, duplication control |
| Add security agent | Security workflow | Scope, evidence, severity model, tool restrictions, escalation |
| Integrate with existing app | Integration workflow | Minimal SDK boundary aligned with repo patterns |
| Deploy or host | Operations workflow | Auth, secrets, model/provider, telemetry, rollback/failure path |
| Improve prompt/agent | Evaluation workflow | Concrete failures, versioned instruction changes, benchmark result |

## Workflow Contract

Every workflow must define:

1. Trigger.
2. Inputs.
3. States.
4. Agent/tool ownership.
5. Permission boundaries.
6. Observable events.
7. Persistent state.
8. Failure behavior.
9. Completion criteria.
10. Validation level.

## Routing Rules

1. Do not implement static folders when the user asked for lifecycle behavior.
2. Do not implement a single overloaded agent when the user expects future workflows.
3. Do not create a generic platform before the first workflow is executable.
4. Do not write to external systems until action targeting and permission policy are defined.
5. Do not skip source verification for SDK APIs, auth, provider, tools, hooks, MCP, or skills.
