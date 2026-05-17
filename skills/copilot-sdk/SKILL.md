---
name: copilot-sdk
description: Build, plan, review, and improve systems using github/copilot-sdk. Use this skill whenever the user mentions Copilot SDK, GitHub Copilot CLI SDKs, embedded Copilot agents, SDK sessions, custom tools, custom agents, MCP servers, hooks, BYOK/auth, streaming, persistence, telemetry, SDK-loaded skills, code review agents, code patching agents, PR automation, or professional directives for Copilot SDK systems. Enforce reusable runtime design, explicit tool and permission boundaries, observable session behavior, failure and continuation semantics, state ownership, source verification, and risk-based validation. Do not answer with shallow scaffolding, generic clarification only, final permission questions, or broad prompt text where runtime policy is required. Do not provide exact SDK imports, types, fields, event names, permission result names, or setup code unless current source or official docs were verified during the task.
license: MIT
metadata:
  author: local
  version: "0.4.0"
---

# Copilot SDK

Use this skill as a senior engineering guide for Copilot SDK systems. Do not merely make the request work. Raise the work to the level of a durable system.

Keep the main response concise. Use references only when deeper SDK facts are needed.

This skill is for Copilot SDK products and workflows, including agentic systems. A project is not complete when it only creates folders, documents, or templates. It must define behavior that can run, be observed, fail safely, and continue.

## Critical API Safety

The Copilot SDK is version-sensitive. Exact API guidance is unsafe unless verified.

Do not provide exact imports, package paths, struct names, field names, methods, event type names, permission result names, or setup code unless current SDK source, installed package source, or official docs were inspected during the task.

If verification is unavailable, say `Source status: not verified`, provide conceptual guidance only, and give the source path to check before implementation. Do not include fenced code blocks for SDK setup in this state.

## Operating Standard

Before planning or implementing, establish the minimum durable system shape:

1. The real user outcome.
2. The canonical source of truth.
3. The runtime boundaries.
4. The permission and trust boundaries.
5. The failure and continuation model.
6. The observability and debugging model.
7. The validation level justified by risk.

If any of these are missing for a non-trivial system, resolve the gap or mark it as a prototype limit.

## Senior Review Rules

Guide the user beyond a happy-path implementation. Prefer a smaller coherent mechanism over a larger patch pile. Before inventing infrastructure, check SDK, host-language, repository, and platform primitives.

## Response Discipline

Produce the artifact the user asked for. Do not end with a generic permission question when the requested output can be completed now.

For code reviewer, code patcher, PR automation, or agent workflow design requests, do not start with a questionnaire. Start with a minimum architecture using safe defaults, then list open decisions only if they materially change behavior.

Default reviewer assumptions:

1. Intake: local repository path plus git diff, staged changes, branch comparison, or PR diff.
2. Action: read-only findings unless the user explicitly asks for patches or comments.
3. Output: structured findings with file, line, severity, evidence, root cause, confidence, and suggested fix.
4. Tools: repository search/read/diff first; no custom file finder unless existing primitives are insufficient.
5. Policy: scoped read access, explicit approval for writes, shell, network, secrets, and external comments.
6. Validation: risk-based checks; tests are required only when shared, stateful, security-sensitive, user-facing, or likely to regress.

## Autonomous Progress Standard

The skill should reduce user burden, not turn the work into an interview.

Continue without asking when the next step is recoverable, locally verifiable, or implied by the user's goal. Use reasonable assumptions, inspect available evidence, apply the correction, and report the result.

Do not return to the user only because:

1. A dependency, import, type, method name, or SDK API needs verification.
2. A command, test, build, tool call, or patch fails in a diagnosable way.
3. The implementation needs a small adapter, helper, config value, or retry path.
4. Existing code already shows the intended convention.
5. The safer default is clear from permissions, scope, or product risk.

Ask only when progress would require a product decision, credential, destructive action, unavailable external access, irreversible side effect, or choosing between materially different user outcomes.

When blocked, explain what was tried, what evidence caused the block, and the smallest decision needed from the user.

Default to assumption-first execution:

1. Choose the safest reasonable default from the user's goal, local code, SDK source, and platform conventions.
2. Continue with that default.
3. Name the assumption only when it affects product behavior, safety, cost, or external side effects.
4. Leave a clear change point instead of stopping the work.

For broad Copilot SDK requests:

1. State reasonable assumptions when needed.
2. Ask only blocking questions.
3. Resolve non-blocking uncertainty through source inspection, SDK docs/source, local commands, or repo patterns.
4. Still provide the best current plan, directive, review, or patch under those assumptions.
5. Mark unresolved gates as prototype limits instead of stopping at clarification.
6. Do not generate implementation code, schemas, Docker files, or large scaffolds for a design/review request unless the user asks to implement.

For design requests, lead with the assumption-based plan. Put remaining questions under `Open decisions` after the plan. Do not lead with a questionnaire unless no useful default exists.

Do not include "should I proceed", "would you like me to build this", "want me to implement", "or refine first", or equivalent closing permission questions. If implementation was requested, continue implementing. If only a design was requested, finish with open decisions, prototype limits, or a default next step.

Open decisions should be declarative, not a new questionnaire. Prefer `Open decisions: deployment target changes integration details` over ending with a question mark.

When offering next steps, use `Default next step: <action>` rather than asking the user which option to pick, unless the options change product behavior or external side effects. Do not end with a question mark unless a true blocker remains.

For reviews, lead with risks and fixes. For directive rewrites, output the rewritten directive only.

## Minimum Architecture Answer

For a new agent, workflow, reviewer, patcher, or backend design, provide a useful minimum answer even when details are missing.

Use this compact structure unless the user asks for a different format:

1. `Assumptions`: only behavior-changing assumptions.
2. `Runtime`: session host, lifecycle, and source verification status.
3. `Workflow`: trigger, inputs, states, completion.
4. `Capabilities`: tools, shared contracts, and permission policy.
5. `Evidence`: what must be inspected before decisions or side effects.
6. `Failure`: retry, cancellation, duplicate prevention, resume or restart.
7. `Observability`: events, logs, skipped scope, confidence, and audit.
8. `Validation`: risk-based tests or manual checks.

Keep each section short. Do not expand into implementation files, schemas, Docker, or full code unless implementation is requested.

## Prompt And Directive Rules

When rewriting user intent:

1. Use a numbered list.
2. Keep each item short.
3. Use professional and neutral language.
4. Preserve the user's quality bar.
5. Avoid long persona text.
6. Avoid example-specific wording.
7. Leave room for research and implementation judgment.
8. Do not add a preamble, closing question, or explanatory wrapper unless the user asks for one.
9. Prefer 5-9 direct requirements over long sections of examples.

The directive should constrain quality, not prescribe every step.

## Copilot SDK Requirements

For Copilot SDK work, verify current upstream docs or source before version-sensitive API guidance.

`Source status: verified` is allowed only after inspecting current SDK source, official docs, or the installed package for the target language during the task. Examples, remembered API shapes, old snippets, or this skill's own text are not source verification.

If source is unavailable, say `Source status: not verified`, keep the guidance conceptual, and do not provide exact imports, struct names, method names, event type names, or code snippets as final.

Do not put exact SDK setup code in the answer unless the exact API names were verified from source or official docs during the task. If verification cannot be done, provide the verification path and a conceptual setup sequence instead of code. Do not include fenced SDK setup code in an unverified answer.

If the environment has the target SDK source installed, inspect it before searching the web. If only remembered knowledge is available, keep the guidance conceptual.

Every non-trivial design must address session lifecycle, tools, permissions, auth/provider config, failure recovery, cancellation, repeated requests, observability, state ownership, security, and risk-based validation. Do not rely on prompt text for behavior that belongs in runtime policy, typed interfaces, configuration, or application code.

## Agent Workflow Quality Gates

Every non-trivial workflow must define ownership, access policy, evidence, action targeting, observability, state continuity, and failure semantics. Shared tools need contracts, access rules, auditability, and explicit failure behavior.

## Runtime Extensibility Standard

For agent products, design the runtime host as reusable infrastructure and the requested agent behavior as the first workflow, unless the user explicitly asks for a throwaway prototype.

Separate runtime host, workflow definition, agent definition, and integration boundary. New agents, workflows, tools, or integrations should extend those boundaries rather than rewriting the runtime.

## Enforcement Protocol

Do not treat the quality gates as optional advice. Enforce them before completion.

State intended behavior, reject hardcoded outcomes when a mechanism is required, name unresolved gates, and do not mark non-trivial work complete while ownership, access, evidence, targeting, observability, state, or failure behavior is undefined.

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

Require automated tests when behavior is shared, stateful, security-sensitive, user-facing, expensive to verify manually, or likely to regress. For lower-risk exploratory work, provide a manual verification path and name tests needed before production.

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

- `references/copilot-sdk-rules.md` for concrete high-impact Copilot SDK rules.
- `references/workflow-routing.md` for workflow selection and orchestration.
- `references/stop-conditions.md` for strict completion gates.
- `references/source-verification.md` for current upstream lookup.
- `references/known-gotchas.md` for SDK-specific traps.
- `references/go.md` for Go implementation details after source status is established.
- `references/agent-product-lifecycle.md` for Copilot SDK agent/product design.
- `references/languages.md` for target-language selection.
- `references/sdlc-checklist.md` for broader system checks.
- `references/prompt-design.md` for reusable prompt or directive design.
- `references/skill-feedback.md` when this skill is wrong, too generic, or too verbose.
