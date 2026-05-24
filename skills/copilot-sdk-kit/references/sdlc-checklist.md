# SDLC Checklist

Use this checklist to keep plans concise and rigorous.

## Discovery

1. Define the user outcome.
2. Identify the host application boundary.
3. Identify the canonical source of truth.
4. Identify current repo patterns and available tools.
5. Verify upstream SDK APIs before version-sensitive guidance.

## Architecture

1. Use SDK primitives before custom orchestration.
2. Keep session, tool, permission, event, and persistence ownership separate.
3. Prefer one coherent path over parallel temporary paths.
4. Avoid overfitting to one prompt, file, tool call, or demo scenario.
5. Design explicit failure behavior before implementation.

## Implementation

1. Keep prompts concise and general.
2. Use typed tool contracts where possible.
3. Scope built-in tools and MCP tools deliberately.
4. Handle cancellation, timeout, repeated requests, and partial results.
5. Clean up sessions and client processes.

## Validation

Tests are risk-based.

Require automated tests for:

1. Shared tool contracts.
2. Permission policies.
3. State transitions.
4. Hooks that modify behavior.
5. Auth/provider selection.
6. Security-sensitive paths.

Use manual verification for low-risk prototypes, and state what must be tested before production.

## Review

Before declaring work complete:

1. Check whether the root problem is solved.
2. Check whether complexity was reduced or merely moved.
3. Check whether the design generalizes.
4. Check whether failure modes are explicit.
5. Check whether the implementation used existing SDK and host primitives.

