---
title: Copilot SDK Known Gotchas
impact: CRITICAL
impactDescription: Prevents shallow implementations that appear to work but fail under real SDK usage
tags: gotchas, permissions, mcp, skills, custom-agents, cli
---

## Copilot SDK Known Gotchas

Use this file before architecture, implementation, or review work.

## Public Preview

The SDK is useful but not a stable long-term contract. Do not promise production safety or API stability without checking current docs.

## CLI Lifecycle

All SDKs communicate with the Copilot CLI server. Lifecycle ownership matters:

1. Bundled CLI, local CLI, and external server modes have different deployment and cleanup behavior.
2. Go and Rust may require manual CLI installation or explicit application-level bundling.
3. External server mode changes process ownership; do not design spawn/stop behavior blindly.

## Permissions

Permission handlers are a runtime safety boundary.

**Incorrect:**

```text
Tell the agent in the prompt not to run dangerous tools, then approve every tool call.
```

**Correct:**

```text
Use a scoped permission policy that allows safe reads, asks for risky operations, and denies unsupported actions.
```

## Skills And Custom Agents

Do not assume skill content automatically applies everywhere.

1. Verify how `SkillDirectories` are resolved in the target language.
2. Verify whether custom agents need explicit skill configuration.
3. Keep durable instructions in skills, not giant session prompts.
4. Keep scoped expertise and tool restrictions in custom agents.

## MCP

MCP is the right boundary for external tool ecosystems.

1. Confirm local vs remote server type.
2. Confirm command, args, env, cwd, URL, headers, timeout, and enabled tools.
3. Handle server startup failure and remote connection failure.
4. Do not build a bespoke subprocess bridge unless MCP is a poor fit.

## Events And Repeated Requests

Prefer streaming and lifecycle events over polling. Define behavior for:

1. Duplicate sends.
2. User spam.
3. Cancellation.
4. Session resume.
5. Partial result display.
6. Idle/completion detection.

