---
title: Verified API Ledger
impact: CRITICAL
impactDescription: Concise verified Copilot SDK API names for starter-level Go, TypeScript, and Python guidance
tags: source-verification, api-ledger, go, typescript, python
checked: 2026-05-18
---

# Verified API Ledger

Use this file only as a concise verification ledger. Do not treat it as a tutorial.

## Verification Rule

1. Prefer installed SDK source or lockfiles over this ledger.
2. Use this ledger for starter-level Go, TypeScript, and Python guidance.
3. Recheck upstream or installed source before production implementation.
4. Do not provide exact .NET, Java, or Rust code unless current source was inspected during the task.

## Sources Checked

1. Go: https://raw.githubusercontent.com/github/copilot-sdk/main/go/README.md
2. TypeScript: https://raw.githubusercontent.com/github/copilot-sdk/main/nodejs/README.md
3. Python: https://raw.githubusercontent.com/github/copilot-sdk/main/python/README.md
4. Skills: https://raw.githubusercontent.com/github/copilot-sdk/main/docs/features/skills.md
5. Custom agents: https://raw.githubusercontent.com/github/copilot-sdk/main/docs/features/custom-agents.md
6. MCP: https://raw.githubusercontent.com/github/copilot-sdk/main/docs/features/mcp.md
7. BYOK: https://raw.githubusercontent.com/github/copilot-sdk/main/docs/auth/byok.md
8. Local Go module: `github.com/github/copilot-sdk/go v0.3.0`

## Go

1. Install: `go get github.com/github/copilot-sdk/go`.
2. Import: `github.com/github/copilot-sdk/go`.
3. Client: `copilot.NewClient`, `ClientOptions`, `Start`, `Stop`.
4. Session: `CreateSession`, `SessionConfig`, `Disconnect`.
5. Send: `Send`, `SendAndWait`, `MessageOptions`.
6. Events: `Streaming`, `SessionEvent`, `AssistantMessageDeltaData`, `AssistantMessageData`, `SessionIdleData`.
7. Permissions: `OnPermissionRequest`, `PermissionRequestKindRead`, `PermissionRequestKindWrite`, `PermissionRequestKindShell`, `PermissionRequestResultKindApproved`, `PermissionRequestResultKindRejected`, `PermissionRequestResultKindUserNotAvailable`.
8. Skills: `SkillDirectories`.
9. Custom agents: `CustomAgents`, `CustomAgentConfig`, `Skills`.
10. MCP/BYOK: `MCPServers`, `MCPStdioServerConfig`, `Provider`, `ProviderConfig`.
11. Gotcha: Older docs may show aliases such as `KindShell`; installed `v0.3.0` exposes explicit `PermissionRequestKind*` constants.

## TypeScript

1. Install: `npm install @github/copilot-sdk`.
2. Import: `@github/copilot-sdk`.
3. Client: `CopilotClient`, `start`, `stop`.
4. Session: `createSession`, `disconnect`.
5. Send: `send`, `sendAndWait`.
6. Events: `session.on("assistant.message", ...)`, `session.on("session.idle", ...)`.
7. Permissions: `onPermissionRequest`, request kinds `"read"`, `"write"`, `"shell"`, `"mcp"`, `"custom-tool"`, `"url"`, `"memory"`, `"hook"`.
8. Permission results: `"approved"`, `"denied-interactively-by-user"`, `"denied-no-approval-rule-and-could-not-request-from-user"`, `"denied-by-rules"`.
9. Skills: `skillDirectories`.
10. Custom agents: `customAgents`, `skills`.
11. MCP/BYOK: `mcpServers`, `provider`, `baseUrl`, `wireApi`, `apiKey`.
12. Gotcha: Custom-agent skills are explicit; sub-agents do not inherit parent skills by default.

## Python

1. Install: `pip install copilot-sdk[telemetry]`.
2. Import: `from copilot import CopilotClient`.
3. Client: `CopilotClient`, `start`, `stop`, async context manager.
4. Session: `create_session`, async context manager.
5. Send: `send`, `send_and_wait`.
6. Events: `session.on(...)`, `AssistantMessageData`, `SessionIdleData`.
7. Permissions: `on_permission_request`, `PermissionRequestResult`.
8. Permission values: `"read"`, `"write"`, `"shell"`, `"mcp"`, `"custom-tool"`, `"url"`, `"memory"`, `"hook"`, `"approved"`, `"denied-no-approval-rule-and-could-not-request-from-user"`.
9. Skills: `skill_directories`.
10. Custom agents: `custom_agents`, `skills`.
11. MCP/BYOK: `mcp_servers`, `provider`, `base_url`, `wire_api`, `api_key`.
12. Gotcha: Use project lockfiles or installed package metadata when available; Python packaging details may change during preview.
