# Go Minimal Starter

Verified against API ledger: 2026-05-18.

This is a minimal Copilot SDK starter, not a full application. Recheck the installed SDK before production use.

Run shape:

```bash
go mod init example.com/copilot-starter
go get github.com/github/copilot-sdk/go
go run .
```

It demonstrates client/session lifecycle, a conservative permission handler, streaming event capture, prompt sending, and cleanup.
