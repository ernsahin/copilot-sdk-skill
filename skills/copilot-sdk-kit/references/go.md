# Go Guidance

Use this reference when the target language is Go or when the user has not chosen a language and a concrete implementation path is needed.

This file is not a substitute for source verification. Treat field names, type names, event names, permission result names, and code snippets as conceptual until checked against the installed package or upstream Go source for the target version.

Verify current source before final code:

- Go README: https://raw.githubusercontent.com/github/copilot-sdk/main/go/README.md
- Go types: https://raw.githubusercontent.com/github/copilot-sdk/main/go/types.go
- Go package path: `github.com/github/copilot-sdk/go`

## Core Flow

A Go Copilot SDK app normally:

1. Creates or connects a Copilot SDK client.
2. Starts the client or relies on autostart.
3. Creates a session with the SDK's session configuration type.
4. Provides a runtime permission handler.
5. Registers event handlers before or immediately after session creation.
6. Sends messages with `Send`.
7. Disconnects the session and stops the client during cleanup.

Use `context.Context` consistently for cancellation and deadlines.

## Session Configuration

Review these fields before designing new abstractions:

1. `Model` and `ReasoningEffort` for model selection.
2. `Tools` for host-provided capabilities.
3. `AvailableTools` and `ExcludedTools` for built-in tool scope.
4. The permission handler field for tool approval policy.
5. `OnUserInputRequest` for user input requests.
6. `Hooks` for lifecycle, prompt, tool, and error interception.
7. `WorkingDirectory` and `ConfigDir` for execution and state boundaries.
8. `Streaming` and `OnEvent` for responsive UIs.
9. `Provider` for BYOK configuration.
10. `MCPServers` for external tool providers.
11. Custom-agent configuration fields for scoped agent behavior.
12. Skill-directory and disabled-skill fields for reusable instructions.
13. `InfiniteSessions` for persistent long-running workflows.

## Tool Design

Prefer typed tools.

1. Use the SDK's typed-tool helper when available for typed parameters and generated schemas.
2. Keep tools narrow, deterministic, and auditable.
3. Return concise data for the model and richer logs for the session/user when supported.
4. Do not expose broad filesystem, shell, network, or database access without a permission policy.
5. Treat tool arguments as untrusted input.
6. Add tests when a tool changes data, reads sensitive data, or becomes shared infrastructure.

## Permission Policy

Do not rely on instructions alone for safety.

Design permission behavior for:

1. Read-only operations.
2. File edits.
3. Shell commands.
4. Network access.
5. Secret access.
6. External system writes.
7. User-confirmed destructive operations.

Use approve-all only for local prototypes or controlled demos. Production paths need scoped rules, logging, and explicit denial behavior.

## Hooks And Events

Use hooks for boundaries:

1. `OnUserPromptSubmitted` for prompt normalization or extra context.
2. `OnPreToolUse` for policy, argument checks, and approval routing.
3. `OnPostToolUse` for redaction, audit, and result shaping.
4. Session-start hooks for context setup.
5. `OnSessionEnd` for summaries and cleanup.
6. `OnErrorOccurred` for retry, abort, or user notification.

Use events for UI updates and monitoring. Avoid polling when streaming or lifecycle events are available.

## Production Concerns

For production candidates, define:

1. Auth source and token lifetime.
2. BYOK provider config and model list behavior.
3. Tenant isolation and session storage boundaries.
4. Rate limits and repeated prompt behavior.
5. Timeout and cancellation policy.
6. Telemetry fields and content-capture policy.
7. Upgrade plan for public-preview API changes.
