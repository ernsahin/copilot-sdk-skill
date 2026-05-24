---
title: Skill Loaded Custom Agent Workflow
impact: HIGH
impactDescription: Defines how Copilot SDK custom agents should load skills without context leaks or inheritance mistakes
tags: workflow, custom-agents, skills, context
---

# Skill Loaded Custom Agent Workflow

Use this when the user asks for custom agents, sub-agents, skill directories, agent-specific expertise, or multi-agent workflows.

## Verified Behavior

Official docs state that:

1. `skillDirectories` points to the parent directory containing named skill folders.
2. Custom agents can declare a `skills` list.
3. Listed skills are eagerly injected into that custom agent's context at startup.
4. Agents receive no skills by default.
5. Sub-agents do not inherit skills from the parent.

Recheck current upstream docs or installed source before exact code.

## Runtime Shape

1. Host owns the skill directory path and version of local skill content.
2. Session owns `skillDirectories`, `customAgents`, selected `agent`, tool filters, and event stream.
3. Each custom agent owns its prompt, tool scope, optional MCP servers, and optional skills.
4. Parent session owns orchestration and final answer synthesis.

## Workflow

1. Define agent responsibilities before prompts.
2. Give each agent a specific `description` for runtime selection.
3. Preload only the skills that agent needs.
4. Keep large durable instructions in skills, not in one giant session prompt.
5. Use explicit tool lists for least privilege.
6. Log sub-agent start, completion, failure, skipped scope, and final contribution.

## Failure

Handle missing skill directories, misspelled skill names, excessive context from broad skills, poor agent selection from vague descriptions, sub-agent failure, and parent/child disagreement.
