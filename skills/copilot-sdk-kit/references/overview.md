# Copilot SDK Overview

Use this reference for architecture-level decisions.

## Current Upstream Baseline

The GitHub Copilot SDK is a public-preview SDK for embedding Copilot agent workflows in applications. It supports TypeScript/Node.js, Python, Go, .NET, Java, and Rust technical preview. All SDKs communicate with a Copilot CLI server through JSON-RPC.

Because the SDK is public preview, verify upstream before final code or production guidance:

- Repository: https://github.com/github/copilot-sdk
- Docs index: https://raw.githubusercontent.com/github/copilot-sdk/main/docs/index.md
- Features docs: https://github.com/github/copilot-sdk/tree/main/docs/features
- Setup docs: https://github.com/github/copilot-sdk/tree/main/docs/setup
- Target language source: `nodejs/`, `python/`, `go/`, `dotnet/`, `java/`, or `rust/`

## Architecture Model

A Copilot SDK app usually has these layers:

1. Host application: owns product flow, auth, persistence, tenancy, and UI.
2. SDK client: starts or connects to the Copilot CLI server.
3. Session: owns model, tools, permissions, hooks, custom agents, MCP servers, skills, streaming, and persistence.
4. Tool boundary: exposes host capabilities to the agent through typed contracts.
5. Permission boundary: decides whether tool execution is allowed.
6. Event boundary: streams agent messages, reasoning, tool activity, sub-agent lifecycle, and idle/completion states.

Keep these layers distinct. Do not solve permissions in the prompt, persistence in tool handlers, or UI policy in the SDK session unless that is the canonical owner.

## Planning Checks

Before designing:

1. Identify the canonical source of truth.
2. Identify the minimum SDK primitives needed.
3. Define explicit failure behavior.
4. Decide which behavior is user-facing and which remains internal.
5. Use existing search, tool, hook, event, and permission primitives before creating custom infrastructure.

## Shallow-Path Warnings

Stop and reconsider if the solution:

1. Creates a custom file crawler when code search is available.
2. Stores session state in multiple incompatible places.
3. Uses prompt text to enforce security instead of permission handlers.
4. Adds a custom agent router before evaluating SDK custom agents.
5. Builds a subprocess bridge where MCP already fits.
6. Claims success without handling denial, timeout, auth failure, cancellation, or repeated requests.
7. Adds tests only for one observed example instead of a recurring invariant.

