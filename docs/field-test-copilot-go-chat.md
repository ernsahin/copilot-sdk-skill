# Field Test: Copilot Go Chat

This field test was created to verify that the skill pushes toward an executable Copilot SDK mechanism rather than a hardcoded demo.

## Test Artifact

Local path:

```text
C:\Users\Eren\Videos\skills\copilot-go-chat
```

The artifact is intentionally not part of this skill package. It is a separate test project used to validate the skill's design pressure.

## What Was Built

A minimal Go CLI that:

- Accepts runtime prompts.
- Starts a Copilot SDK client and session.
- Uses streaming session events.
- Applies a conservative permission policy.
- Logs tool events, permission requests, and session errors.
- Uses timeout/cancellation.
- Includes tests for permission policy behavior.

## Verification Run

Commands run:

```powershell
go test ./...
go build ./cmd/copilot-chat
.\copilot-chat.exe -prompt "Say hello in one short sentence." -timeout 45s
.\copilot-chat.exe -prompt "Inspect the working directory at a high level and summarize what this Go project does. Do not modify files." -timeout 90s
```

Observed behavior:

- The CLI connected to GitHub Copilot SDK successfully.
- Copilot responded through the session.
- Built-in tools such as `view` and `glob` were used by the runtime.
- Read permission requests were approved by policy.
- Tool start/complete and permission decisions were visible in logs.

## What This Proved

The skill's standard pushed the implementation toward:

- A reusable mechanism instead of a hardcoded response.
- Runtime permission enforcement instead of prompt-only safety.
- Observable tool and permission behavior.
- Minimal but real verification.

## What This Did Not Prove

This was not a full benchmark of the skill. It tested one implementation path and did not compare baseline vs with-skill outputs.
