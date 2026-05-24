---
title: Code Patcher Workflow
impact: HIGH
impactDescription: Defines safe Copilot SDK patching behavior with scoped writes and verification
tags: workflow, patching, permissions, validation
---

# Code Patcher Workflow

Use this when the user asks for a Copilot SDK agent that edits code, fixes tests, migrates code, or opens changes.

## Runtime Shape

1. Host owns repository path, write scope, branch/worktree policy, test commands, and rollback strategy.
2. Session owns plan, edits, verification loop, and final summary.
3. Tools expose read/search before edit. Write tools must be scoped to the selected repository and intended file set.
4. Permission policy denies broad shell/write by default and allows only targeted, auditable operations.

## Workflow

1. Diagnose first: inspect failing output, relevant source, tests, and existing patterns.
2. Patch narrowly: change the smallest coherent set of files that addresses the root cause.
3. Verify: run the targeted test/build command first, then broader checks when risk justifies it.
4. Recover: if import/API/build errors are diagnosable, inspect source and fix without asking the user.
5. Stop only for product choices, credentials, destructive actions, unavailable external services, or conflicting user edits.

## Output Contract

Return:

1. What changed.
2. Why the change fixes the root cause.
3. Verification commands and results.
4. Remaining risk or checks that could not run.

Do not end with a non-blocking permission question after the user already asked for implementation.

## Failure

Handle denied writes, test command failure, flaky tests, merge conflicts, generated files, formatters that rewrite unrelated files, and duplicate external side effects. Never revert user changes unless explicitly instructed.
