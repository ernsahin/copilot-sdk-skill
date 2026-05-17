---
title: Stop Conditions
impact: CRITICAL
impactDescription: Prevents premature completion when core agent-system quality gates are unresolved
tags: enforcement, quality, completion, stop
---

# Stop Conditions

Use this file when the agent is about to finalize a plan, review, or implementation.

## Hard Stops

Do not mark non-trivial Copilot SDK work complete if any of these are unresolved:

1. The intended behavior is not stated.
2. The solution is a hardcoded outcome where a reusable mechanism is required.
3. Runtime host responsibilities are mixed into a single agent prompt.
4. Tool access is broad, implicit, or unaudited.
5. Permission policy is replaced by prompt text.
6. Evidence requirements are undefined.
7. External side effects cannot be target-verified.
8. Failure behavior is missing for ambiguous, stale, incomplete, or contradictory evidence.
9. Observability is limited to vague logging.
10. State ownership is duplicated or unclear.
11. Validation is either absent or excessive for the risk level.
12. SDK API names or setup claims were not verified against current upstream sources.
13. The response is only a clarification prompt when a useful assumption-based plan, review, or directive could be provided now.
14. The response ends with a generic "should I proceed" question after producing the requested artifact.
15. The agent stops at a diagnosable error instead of inspecting the cause, applying a reasonable fix, and continuing.
16. The agent asks the user to choose an implementation detail that is already implied by local code, SDK source, platform conventions, or the user's stated outcome.
17. A broad design answer expands into large code, schemas, deployment files, or exhaustive configuration before the core decisions are settled.
18. Version-sensitive SDK guidance names exact APIs without a source status.
19. The response claims source verification from memory, generated examples, previous answers, or skill text.
20. A design response starts with a questionnaire when a reasonable default architecture can be provided first.
21. The response includes exact SDK setup code while source status is not verified.
22. The response asks for permission to proceed after the user already asked to build, fix, create, or continue.

## Recovery Bias

Prefer recovery over handoff when the failure is reversible and locally testable.

Use this order:

1. Inspect the error and surrounding context.
2. Check local code, SDK source, docs, or existing conventions.
3. Apply the smallest coherent correction.
4. Re-run the relevant verification.
5. Report only the final state and any remaining real blocker.

Do not hide repeated failure. If the same class of fix fails twice, stop and summarize the evidence.

## Clarification Boundary

Clarifying questions are allowed only when the answer changes the user-facing product behavior, external side effects, credential/security model, cost profile, or irreversible action.

Otherwise, proceed with a stated assumption and leave a change point.

For design work, place non-blocking questions after the useful plan as `Open decisions`.

## Acceptable Prototype Limit

If the user explicitly wants a prototype, unresolved gates may remain only when they are named clearly.

Use this format:

```text
Prototype limit:
- This version demonstrates <capability>.
- It does not yet handle <unresolved gate>.
- Before production, implement <required next control>.
```

## Completion Statement

Before finalizing, state:

1. What is implemented or planned.
2. Which quality gates are satisfied.
3. Which gates remain prototype limits, if any.
4. What verification was actually run.
