---
title: Copilot SDK Rules By Impact
impact: CRITICAL
impactDescription: Converts broad quality standards into concrete Copilot SDK implementation rules
tags: rules, copilot-sdk, permissions, sessions, tools, observability
---

# Copilot SDK Rules By Impact

Use these rules when planning, implementing, or reviewing Copilot SDK systems.

## Critical Rules

### runtime-permission-boundary

Permissions must be enforced through runtime policy, not only prompt text.

Reject designs where prompts say "do not do dangerous actions" but the permission handler approves everything. For production-like work, define which operations are allowed, denied, interactive, or environment-specific.

### runtime-source-verification

Version-sensitive API guidance must be checked against current upstream docs or source before implementation.

The SDK is public preview. Do not rely on remembered field names, setup steps, or production caveats.

### runtime-state-owner

Session state, workflow state, external integration state, and persisted product state need explicit ownership.

Reject designs where multiple agents or tools independently write overlapping state without a canonical owner.

### runtime-side-effect-targeting

External side effects must be target-verifiable before execution.

Any write, comment, patch, deployment, shell command, network call, or external API mutation must know exactly what resource it targets and how the system prevents duplicate or wrong-target execution.

## High-Impact Rules

### workflow-first-agent-second

For agent products, design the workflow before agent prompts.

Define trigger, input boundary, states, routing, continuation, output contract, and failure behavior. Then define which agent or tool owns each step.

### shared-tool-contracts

Shared tools must have stable contracts, access policies, audit fields, and failure semantics.

Do not hide product behavior inside one agent prompt when future agents are expected to reuse the same capability.

### observable-agent-work

Agent work must be inspectable at the product boundary.

Expose progress, inspected evidence, skipped scope, tool calls, permission decisions, errors, confidence limits, and final coverage summary when the workflow is non-trivial.

### native-primitives-first

Use Copilot SDK and host-environment primitives before custom infrastructure.

Prefer built-in sessions, events, permission handlers, tools, skills, custom agents, MCP, and repository search/read tools before inventing replacements.

## Medium-Impact Rules

### risk-based-tests

Tests are required when behavior is shared, stateful, security-sensitive, user-facing, expensive to verify manually, or likely to regress.

For prototypes, provide a manual verification path and name what must be tested before production.

### compact-prompts

Prompts should define goal, constraints, boundaries, evidence, and failure expectations. They should not become giant operating manuals.

Durable instructions belong in skills or versioned agent definitions. Runtime controls belong in code and configuration.

### extension-without-rewrite

Adding a second workflow, agent, tool, or integration should usually extend contracts and configuration, not rewrite the runtime host.

Avoid both extremes: one-off hardcoded agents and speculative generic platforms with no working first workflow.
