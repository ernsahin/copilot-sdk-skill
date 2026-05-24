---
title: BYOK Backend Workflow
impact: HIGH
impactDescription: Defines hosted or enterprise Copilot SDK backends with provider, auth, persistence, and telemetry boundaries
tags: workflow, byok, backend, auth, telemetry
---

# BYOK Backend Workflow

Use this when the user asks for a hosted Copilot SDK backend, enterprise deployment, custom provider, Azure/OpenAI-compatible endpoint, Anthropic, local provider, billing isolation, streaming API, or resumable sessions.

## Runtime Shape

1. Host application owns users, tenants, auth, provider secrets, rate limits, persistence, and HTTP/WebSocket/SSE API.
2. SDK client owns Copilot CLI connection lifecycle and provider configuration.
3. Session owns model, tools, permissions, hooks, custom agents, MCP servers, streaming, and resume behavior.
4. Storage owns product state. Do not let tools independently create a second source of truth.

## Workflow

1. Resolve provider type, model/deployment name, base URL, wire API, and secret source.
2. Require model when custom provider is used.
3. Define session ID strategy for resume, cancellation, repeated requests, and tenant isolation.
4. Stream assistant messages, tool activity, permission decisions, and errors to the client.
5. Capture telemetry without sensitive content unless explicitly approved.

## Failure

Handle missing API key, invalid provider config, unsupported model, rate limits, auth expiration, MCP startup failure, denied permissions, session resume misses, duplicate sends, cancellation, and partial results.

## Validation

Require automated tests for provider config parsing, permission policy, tenant/session isolation, duplicate request handling, and event stream completion. Manual checks are acceptable only for local prototypes.
