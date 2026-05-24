---
title: MCP Backed Agent Workflow
impact: HIGH
impactDescription: Defines safe MCP integration boundaries for Copilot SDK agents
tags: workflow, mcp, tools, external-systems
---

# MCP Backed Agent Workflow

Use this when the user asks for MCP tools, external tool ecosystems, database access, GitHub integrations, browser automation, or remote/local tool servers.

## Runtime Shape

1. Host owns MCP server configuration, credentials, network policy, and process lifecycle.
2. Session owns MCP server allowlist, tool routing, permission decisions, and event logging.
3. MCP servers should be configured explicitly: type, command or URL, args, env, cwd, headers, tools, and timeout.
4. Tool exposure should be least-privilege. Avoid `["*"]` outside controlled prototypes.

## Workflow

1. Identify whether each server is local stdio or remote HTTP/SSE.
2. Define startup behavior, timeout, retry, and failure message.
3. Restrict tools to the workflow's needs.
4. Route every write or external mutation through permission policy and target verification.
5. Capture server startup, tool invocation, denial, timeout, and result summary in logs/events.

## Failure

Handle server startup failure, missing command, invalid environment, auth failure, remote timeout, schema mismatch, unavailable tools, duplicate writes, and cancellation. Do not build a bespoke subprocess bridge while MCP fits the integration.
