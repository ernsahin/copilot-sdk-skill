# Language Strategy

The skill covers all official Copilot SDK languages. Use Go as the deepest default path, but do not force Go when the user's project clearly uses another SDK.

## Selection Rule

1. Use the language already present in the repository.
2. If there is no repository signal, ask for the language only when implementation depends on it.
3. If the user wants a concrete recommendation and has no preference, use Go for backend/CLI/system tools and TypeScript for web app integration.
4. Verify the target language docs/source before writing code.

## Official SDKs

Use the upstream repository as the source:

1. TypeScript/Node.js: `nodejs/`, package `@github/copilot-sdk`.
2. Python: `python/`, package `github-copilot-sdk`.
3. Go: `go/`, module `github.com/github/copilot-sdk/go`.
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

