---
title: Prompt And Directive Design
impact: MEDIUM
impactDescription: Keeps prompts useful without forcing brittle or overfitted agent behavior
tags: prompts, directives, concise, anti-shallow
---

## Prompt And Directive Design

Prompts should guide reasoning without replacing research. Good prompts are short enough for the agent to use and general enough to allow the best implementation path.

## Directive Pattern

Use this structure:

1. Goal.
2. Runtime boundary.
3. Source verification requirement.
4. SDK primitives to prefer.
5. Failure modes to handle.
6. Validation standard.

## Example

```text
1. Build a Copilot SDK code patching workflow for the existing repository.
2. Inspect the repository before choosing the implementation path.
3. Verify the current SDK API for the target language before writing code.
4. Use SDK-native sessions, tools, permissions, hooks, and events before custom orchestration.
5. Handle denied permissions, tool failures, cancellation, repeated requests, and partial results.
6. Add automated tests only for shared, risky, or regression-prone behavior; otherwise provide a manual verification path.
```

## Avoid

1. Long persona blocks.
2. Emotional quality language without operational checks.
3. One-case rules that only solve the example in the conversation.
4. Commands that force implementation before research.
5. Vague requests such as "make it advanced" without defining failure modes or source of truth.
6. Static folder structures that pretend to implement a workflow.

## Product Test

When the user asks for an agent system, the output should describe or build executable behavior, not only a file taxonomy.

Ask:

1. What action does the system perform?
2. What input causes that action?
3. What output proves it happened?
4. What state must survive between steps?
5. What happens when a step fails?

Folders are acceptable only when they support real workflow behavior.

## Review Questions

Before using a prompt:

1. Does it leave room to inspect the codebase?
2. Does it tell the agent which primitives to prefer?
3. Does it define failure behavior?
4. Does it avoid overfitting to one example?
5. Is every sentence worth its token cost?
