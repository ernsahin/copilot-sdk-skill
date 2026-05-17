---
name: copilot-sdk
description: Build, plan, review, and improve systems using github/copilot-sdk. Use this skill whenever the user mentions Copilot SDK, GitHub Copilot CLI SDKs, embedded Copilot agents, SDK sessions, custom tools, custom agents, MCP servers, hooks, BYOK/auth, streaming, persistence, telemetry, SDK-loaded skills, or wants concise professional technical directives for a Copilot SDK system.
license: MIT
metadata:
  author: local
  version: "0.4.0"
---

# Copilot SDK

Use this skill as a senior engineering guide for Copilot SDK systems. Do not merely make the request work. Raise the work to the level of a durable system.

Keep the main response concise. Use references only when deeper SDK facts are needed.

This skill is for Copilot SDK products and workflows, including agentic systems. A project is not complete when it only creates folders, documents, or templates. It must define behavior that can run, be observed, fail safely, and continue.

## Operating Standard

Before planning or implementing, establish:

1. The real user outcome.
2. The canonical source of truth.
3. The runtime boundaries.
4. The permission and trust boundaries.
5. The failure and continuation model.
6. The observability and debugging model.
7. The validation level justified by risk.

If any of these are missing for a non-trivial system, do not present the work as complete.

## Senior Review Rules

Guide the user beyond a junior happy-path implementation.

Require the solution to answer:

1. What should own each responsibility?
2. What can fail, and what should happen next?
3. What must be visible when the system misbehaves?
4. What must be configurable, persistent, or versioned?
5. What should be verified now, and what can wait?
6. What existing SDK or platform primitive should be used first?
7. What complexity is accidental and should not become architecture?

Prefer a smaller, coherent design over a larger design made from patches.

## Prompt And Directive Rules

When rewriting user intent:

1. Use a numbered list.
2. Keep each item short.
3. Use professional and neutral language.
4. Preserve the user's quality bar.
5. Avoid long persona text.
6. Avoid example-specific wording.
7. Leave room for research and implementation judgment.

The directive should constrain quality, not prescribe every step.

## Copilot SDK Requirements

For Copilot SDK work, verify current upstream docs or source before version-sensitive API guidance.

Every non-trivial design must explicitly address:

1. Session lifecycle.
2. Tool and capability boundaries.
3. Permission policy.
4. Authentication or provider configuration.
5. Failure handling and recovery.
6. Cancellation, repeated requests, and continuation.
7. Operational visibility.
8. State ownership and persistence.
9. Security and data exposure.
10. Risk-based validation.

Do not rely on prompt text for behavior that belongs in runtime policy, typed interfaces, configuration, or application code.

## Agent Workflow Quality Gates

Every non-trivial agent workflow must define:

1. Capability ownership: which component or agent owns each responsibility.
2. Access policy: which capabilities are available, restricted, shared, or denied.
3. Evidence requirements: what inspected inputs justify each decision.
4. Action targeting: how external side effects are aimed at the correct target before execution.
5. Observability: what progress, decisions, skipped work, failures, and confidence limits are visible.
6. State continuity: what context persists across turns, sessions, retries, and repeated requests.
7. Failure semantics: how the system behaves when evidence is missing, ambiguous, stale, or contradictory.

Shared tools must have explicit contracts, access rules, and auditability. Agent decisions must be traceable to inspected evidence. External actions must be verifiable before they are executed.

## Runtime Extensibility Standard

For agent products, design the runtime host as reusable infrastructure and the requested agent behavior as the first workflow, unless the user explicitly asks for a throwaway prototype.

Separate:

1. Runtime host: session lifecycle, tool registry, permissions, events, state, and configuration.
2. Workflow definition: trigger, states, routing, continuation, and completion criteria.
3. Agent definition: role, instructions, allowed capabilities, and evidence requirements.
4. Integration boundary: external systems, credentials, side effects, and retry behavior.

New agents, workflows, tools, or integrations should extend the system through these boundaries instead of requiring a runtime rewrite.

## Enforcement Protocol

Do not treat the quality gates as optional advice. Enforce them before completion.

For non-trivial work:

1. State the intended behavior before choosing implementation details.
2. Identify whether the solution is a reusable mechanism or a hardcoded outcome.
3. Reject hardcoded outcomes when a general mechanism is required.
4. Name any quality gate that is unresolved.
5. Either resolve the gap or explicitly limit the work to a prototype.
6. Do not mark the task complete when ownership, access, evidence, targeting, observability, state, or failure behavior is undefined.

Prefer simple general mechanisms over brittle shortcuts. A small reusable function, policy, schema, or workflow is better than a one-off value that only satisfies the current example.

## Intent Routing

Match the user's request to the correct workflow before designing:

1. New Copilot SDK agent product: define the user-facing workflow, lifecycle states, tools, permissions, state, observability, and validation.
2. Code reviewer or code patcher: define repository intake, analysis scope, review output, change policy, approval points, and continuation behavior.
3. Existing project integration: inspect the codebase first, preserve established patterns, and add only the SDK boundary needed.
4. Deployment or hosting: identify environment, auth, model/provider, secrets, telemetry, and operational failure paths.
5. Prompt or agent improvement: use evaluation, traces, concrete failures, and versioned instruction changes rather than rewriting prompts blindly.

If the user asks for an SDLC lifecycle, treat it as executable workflow governance, not static directory scaffolding.

## Native Primitive Rule

Before inventing custom infrastructure, inspect whether the SDK, host language, repository, or execution environment already provides the needed primitive.

Use custom code only when it has a clear owner, contract, failure behavior, and validation path.

## Test Policy

Tests are not automatically required.

Require automated tests when the behavior is:

1. Shared.
2. Stateful.
3. Security-sensitive.
4. User-facing.
5. Expensive to verify manually.
6. Likely to regress.

For lower-risk exploratory work, provide a manual verification path and name the tests needed before production use.

## Completion Standard

A response is not ready if it:

1. Solves only the happy path.
2. Moves complexity without naming it.
3. Creates duplicate sources of truth.
4. Uses a broad prompt instead of runtime controls.
5. Ignores failure, recovery, or observability.
6. Overfits to one example.
7. Claims production readiness without verification.

When the request is vague, elevate it into clear system requirements before building.

## References

Read only what the task needs:

- `references/source-verification.md` for current upstream lookup.
- `references/known-gotchas.md` for SDK-specific traps.
- `references/go.md` for Go implementation details.
- `references/agent-product-lifecycle.md` for Copilot SDK agent/product design.
- `references/languages.md` for target-language selection.
- `references/sdlc-checklist.md` for broader system checks.
- `references/prompt-design.md` for reusable prompt or directive design.
- `references/skill-feedback.md` when this skill is wrong, too generic, or too verbose.
