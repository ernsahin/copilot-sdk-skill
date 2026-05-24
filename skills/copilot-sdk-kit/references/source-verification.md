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

1. Installed dependency source or lockfile for the project, if present.
2. Concise verified API ledger for Go, TypeScript/Node.js, or Python:
   - `references/verified-api-ledger.md`
3. Main repository: `https://github.com/github/copilot-sdk`
4. Docs index: `https://raw.githubusercontent.com/github/copilot-sdk/main/docs/index.md`
5. Target language directory:
   - TypeScript/Node.js: `nodejs/`
   - Python: `python/`
   - Go: `go/`
   - .NET: `dotnet/`
   - Java: linked Java package/repository from the main README
   - Rust: `rust/`
6. Feature docs:
   - `docs/features/skills.md`
   - `docs/features/mcp.md`
   - `docs/features/custom-agents.md`
   - `docs/features/hooks.md`
   - `docs/features/streaming-events.md`
   - `docs/features/session-persistence.md`
7. Auth/setup docs:
   - `docs/auth/index.md`
   - `docs/auth/byok.md`
   - `docs/setup/index.md`
8. Troubleshooting and compatibility docs.
9. Releases and issues when docs conflict with observed behavior.

The ledger is a context reducer. It is acceptable source verification for starter-level guidance only when the task does not require a newer installed package or production guarantee.

## Version Discipline

Do not mix API shapes across versions or languages.

When implementation guidance depends on exact SDK APIs:

1. Identify the target language and installed or intended version.
2. Confirm field names, method names, config names, and event names from that version.
3. Prefer installed source over the ledger, and the ledger over examples that may lag the package.
4. If the version is unknown, write conceptual guidance and mark exact API names as unverified.

## Verification Claims

Do not claim verification from:

1. Memory.
2. A generated example.
3. A previous answer.
4. A skill instruction.
5. A package name guessed from convention.

Claim verification only from source or docs inspected during the task.

For the verified API ledger, state:

```text
Source status: verified from the API ledger checked on <date>; recheck installed SDK before production.
```

If verification is impossible, use this form:

```text
Source status: not verified.
Use this as conceptual guidance only. Before implementation, inspect the installed SDK or upstream docs for exact imports, types, fields, and event names.
```

## Verification Output

State only the facts needed for the task:

1. Which source was checked.
2. Which API names or setup rules were confirmed.
3. Which assumptions remain uncertain.
4. Which parts should be rechecked before production.
