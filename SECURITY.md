# Security Policy

This repository contains agent skill instructions, examples, and validation scripts.

## Reporting

Report security issues privately to the repository owner. Do not open public issues for suspected malicious content, credential leaks, or unsafe automation behavior.

## Security Expectations

Skill content must not:

- Exfiltrate credentials or private data.
- Encourage bypassing permission systems.
- Hide side effects from the user.
- Recommend approve-all permissions for production workflows.
- Treat prompt text as a substitute for runtime access control.

## Review Standard

Changes to `skills/copilot-sdk/` should be reviewed for:

- Prompt injection risk.
- Unsafe tool or shell guidance.
- Incorrect permission boundary guidance.
- Stale Copilot SDK API claims.
- Instructions that could cause destructive external side effects without explicit approval.

