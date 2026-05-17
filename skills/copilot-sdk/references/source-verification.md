---
title: Verify Current Copilot SDK Sources
impact: CRITICAL
impactDescription: Prevents stale API names, outdated setup advice, and unsafe production claims for a public-preview SDK
tags: verification, docs, public-preview, source
---

## Verify Current Copilot SDK Sources

The Copilot SDK is public preview. Treat version-sensitive facts as untrusted until checked against upstream.

**Incorrect:**

```text
Use a remembered API shape from an example and write final code without checking the current target language source.
```

**Correct:**

```text
Check the current repository docs and target language source, then write only the API shape that matches the verified version.
```

## Lookup Order

1. Main repository: `https://github.com/github/copilot-sdk`
2. Docs index: `https://raw.githubusercontent.com/github/copilot-sdk/main/docs/index.md`
3. Target language directory:
   - TypeScript/Node.js: `nodejs/`
   - Python: `python/`
   - Go: `go/`
   - .NET: `dotnet/`
   - Java: linked Java package/repository from the main README
   - Rust: `rust/`
4. Feature docs:
   - `docs/features/skills.md`
   - `docs/features/mcp.md`
   - `docs/features/custom-agents.md`
   - `docs/features/hooks.md`
   - `docs/features/streaming-events.md`
   - `docs/features/session-persistence.md`
5. Auth/setup docs:
   - `docs/auth/index.md`
   - `docs/auth/byok.md`
   - `docs/setup/index.md`
6. Troubleshooting and compatibility docs.
7. Releases and issues when docs conflict with observed behavior.

## Verification Output

State only the facts needed for the task:

1. Which source was checked.
2. Which API names or setup rules were confirmed.
3. Which assumptions remain uncertain.
4. Which parts should be rechecked before production.

