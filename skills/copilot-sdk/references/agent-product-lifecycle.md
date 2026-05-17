---
title: Copilot SDK Agent Product Lifecycle
impact: HIGH
impactDescription: Prevents static scaffolds and turns vague agent requests into executable product workflows
tags: agents, lifecycle, product, sdlc, code-review, workflow
---

# Copilot SDK Agent Product Lifecycle

Use this reference when the user asks for an agent, code reviewer, code patcher, SDLC lifecycle, or agent-based project using GitHub Copilot SDK.

## Product Rule

An agent system is a product workflow, not a folder layout.

Folders, documents, and templates are supporting artifacts only. They do not count as implementation unless they are connected to executable behavior, state, review, or delivery flow.

## Runtime Extensibility

Design the Copilot SDK runtime host as reusable infrastructure. Treat the user's first requested agent as the first workflow running on that host, not as the entire architecture.

Keep these boundaries separate:

1. Runtime host: owns SDK client/session lifecycle, tool registry, permission policy, event stream, state store, configuration, and shutdown behavior.
2. Workflow definition: owns triggers, lifecycle states, routing, retries, continuation, and completion criteria.
3. Agent definition: owns role, instructions, allowed capabilities, evidence requirements, and output contract.
4. Tool contract: owns inputs, outputs, side effects, failure modes, and audit fields.
5. Integration boundary: owns external credentials, API limits, retries, deduplication, and write-back behavior.
6. Observability layer: owns progress, decisions, inspected evidence, skipped scope, errors, and confidence limits.

Adding a new agent, workflow, tool, or external integration should usually add configuration and contracts, not rewrite the runtime host.

Do not overbuild a platform before the first workflow works. The correct balance is:

1. Generic runtime boundaries.
2. Specific first workflow behavior.
3. Clear extension points.
4. No speculative abstractions without an expected second use.

## Required Shape

Define these before implementation:

1. User outcome: what the system helps the user accomplish.
2. Input boundary: what starts a run.
3. Agent responsibility: what the agent owns and what it must not own.
4. Tool boundary: what capabilities are exposed and how they are permissioned.
5. Lifecycle states: how work moves from intake to result.
6. State ownership: what survives between turns or sessions.
7. Human control points: what requires approval or review.
8. Failure behavior: what happens when the workflow cannot proceed.
9. Observability: how progress, errors, tool use, and confidence are visible.
10. Validation: what proves the workflow works at its current maturity level.

## General Agent Quality Gates

Apply these to any agent product, regardless of domain:

1. Capability ownership: assign each responsibility to the host app, session, agent, tool, policy layer, or persistence layer.
2. Access policy: define which capabilities each agent can use, when access is allowed, and what requires approval.
3. Evidence requirements: define what the agent must inspect before making decisions or producing outputs.
4. Action targeting: define how the system verifies that side effects target the correct resource.
5. Observability: expose progress, inspected evidence, skipped scope, tool use, errors, and confidence limits.
6. State continuity: persist only the context needed for retries, resume, deduplication, and follow-up work.
7. Failure semantics: define what happens when evidence is incomplete, stale, ambiguous, or conflicting.

Do not add domain-specific rules to the main skill when a general quality gate can cover the case.

## Enforcement Pattern

Use this pattern to prevent shallow completion:

1. Define the intended capability in one sentence.
2. Decide whether the capability needs a reusable mechanism or a one-off result.
3. If future inputs are possible, implement the mechanism.
4. If only the current answer is needed, state that the solution is intentionally narrow.
5. Before completion, report unresolved gates or prototype limits.

Shallow shortcuts often look correct for the first example but fail the capability. A computed, inspectable, reusable path is usually stronger than a fixed output when the user is asking for a system.

## Lifecycle Interpretation

When the user says SDLC, interpret it as behavior:

1. Discovery: collect requirements and repository context.
2. Analysis: inspect code and identify constraints.
3. Design: choose architecture and responsibilities.
4. Implementation: create or modify code through controlled tools.
5. Review: classify findings, risks, and approval needs.
6. Verification: run tests or manual checks justified by risk.
7. Release readiness: report state, unresolved risks, and next actions.
8. Maintenance: persist useful context and support continuation.

Do not implement this as empty directories unless the directories are part of a real workflow contract.

## Code Review Agent Standard

A Copilot SDK code review agent should define:

1. What code is reviewed: repo, diff, pull request, branch, file set, or user-selected scope.
2. How context is gathered: repository search, file reads, dependency/config inspection, and project instructions.
3. What findings contain: severity, file reference, root cause, impact, and suggested fix.
4. What the agent may do: read-only review, proposed patches, or approved edits.
5. What must be blocked: broad destructive actions, unapproved writes, secret exposure, and unsupported external calls.
6. How uncertainty is handled: missing context, oversized input, conflicting instructions, and low confidence.
7. How the run is observed: progress events, tool calls, skipped areas, errors, and final coverage summary.

## Completion Check

Before saying the project is done, confirm:

1. The workflow has an executable path.
2. Lifecycle states are tied to behavior.
3. Permissions are enforced outside prompt text.
4. Failure and continuation are defined.
5. Observability exists at the product boundary.
6. Runtime boundaries can accept a second workflow without a rewrite.
7. Validation matches the risk level.
