---
title: Code Reviewer Workflow
impact: HIGH
impactDescription: Turns Copilot SDK review requests into executable repository review workflows
tags: workflow, code-review, repository, permissions
---

# Code Reviewer Workflow

Use this when the user asks for a Copilot SDK code reviewer, PR reviewer, audit agent, or repository review workflow.

## Runtime Shape

1. Host owns repository intake, branch/diff selection, credentials, result storage, and final publishing.
2. Session owns review reasoning, tool calls, custom agents, streaming events, and completion.
3. Tools expose repository search/read/diff operations first. Avoid custom crawlers unless native search is unavailable.
4. Permission policy defaults to read-only. Writing comments, patches, labels, or external statuses requires explicit target verification.

## Workflow

1. Intake: repository path plus one of working tree diff, staged diff, branch comparison, or PR diff.
2. Evidence: inspect changed files, related tests, call sites, configs, and ownership boundaries.
3. Analysis: prioritize bugs, regressions, security issues, data loss, user-facing breakage, and missing tests.
4. Output: findings first, ordered by severity, with file/line, evidence, root cause, confidence, and suggested fix.
5. Completion: include skipped scope and residual risk. Do not claim full review when evidence was unavailable.

## Custom Agents

Use custom agents only when they reduce context or risk:

1. `repository-researcher`: read-only tools, broad search, no writes.
2. `finding-checker`: validates suspected issues against code and tests.
3. `publisher`: optional external comment/status writer with narrow target permissions.

Keep the parent session responsible for final severity ordering and user-facing output.

## Failure

Handle missing git metadata, denied read permission, large diffs, tool errors, test unavailability, duplicate publish attempts, and cancellation. If publish fails, preserve the review result locally and report the exact failed target.
