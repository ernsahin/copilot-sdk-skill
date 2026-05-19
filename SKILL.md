---
name: copilot-sdk
description: "Build, plan, review, and improve production-shaped systems using github/copilot-sdk. Use this skill whenever the user mentions Copilot SDK, GitHub Copilot CLI SDKs, embedded Copilot agents, SDK sessions, Go, TypeScript/Node.js, Python, custom tools, custom agents, MCP servers, hooks, BYOK/auth, streaming, persistence, telemetry, SDK-loaded skills, code review agents, code patching agents, PR automation, or professional directives for Copilot SDK systems. Treat this as a Builder Kit: route to the verified API ledger, workflow playbooks, and minimal starters before inventing code. Enforce reusable runtime design, explicit tool and permission boundaries, observable session behavior, failure and continuation semantics, state ownership, source verification, and risk-based validation. Do not provide exact SDK imports, types, fields, event names, permission result names, or setup code unless current source, installed package source, or the verified API ledger was inspected during the task."
license: MIT
metadata:
  author: local
  version: "0.5.0"
---

# Copilot SDK Builder Kit

This repository-root skill exists so `npx skills add ernsahin/copilot-sdk` installs without extra flags.

Use `skills/copilot-sdk/SKILL.md` as the canonical skill body. Read that file immediately when this root entrypoint triggers, then follow its instructions.

Supporting files live under:

1. `skills/copilot-sdk/references/verified-api-ledger.md`
2. `skills/copilot-sdk/references/workflows/`
3. `skills/copilot-sdk/examples/starters/`
4. `skills/copilot-sdk/evals/evals.json`

Do not duplicate or reinterpret the rules here. The canonical body owns source verification, workflow routing, context discipline, exact API policy, starter usage, and completion standards.
