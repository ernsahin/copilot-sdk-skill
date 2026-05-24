# Language Strategy

The skill covers official Copilot SDK concepts, but deep Builder Kit support is scoped to Go, TypeScript/Node.js, and Python for this version.

## Selection Rule

1. Use the language already present in the repository.
2. If there is no repository signal, ask for the language only when implementation depends on it.
3. If the user wants a concrete recommendation and has no preference, use Go for backend/CLI/system tools, TypeScript for web app integration, and Python for async service prototypes or Python-heavy repositories.
4. Verify the target language docs/source before writing code.
5. For .NET, Java, and Rust, provide source-map guidance only unless current source was inspected during the task.

## Builder Kit Languages

Use the verified API ledger before exact starter-level code:

1. Go: `references/verified-api-ledger.md`, module `github.com/github/copilot-sdk/go`.
2. TypeScript/Node.js: `references/verified-api-ledger.md`, package `@github/copilot-sdk`.
3. Python: `references/verified-api-ledger.md`, import package `copilot`.

## Source-Map Only Languages

Use the upstream repository as the source and do not provide exact code unless verified during the task:

1. .NET: `dotnet/`, package `GitHub.Copilot.SDK`.
2. Java: official Java package linked from the main repository.
3. Rust: `rust/`, technical preview.
4. .NET: `dotnet/`, package `GitHub.Copilot.SDK`.
5. Java: official Java package linked from the main repository.
6. Rust: `rust/`, technical preview.

## Cross-Language Concepts

Most SDK designs should map these concepts across languages:

1. Client lifecycle.
2. Session configuration.
3. Permission handling.
4. Custom tools.
5. Hooks.
6. MCP servers.
7. Custom agents.
8. Skill directories.
9. Streaming and session events.
10. Auth and BYOK providers.
11. Session persistence and resume.

Avoid claiming exact field names across languages until verified.
