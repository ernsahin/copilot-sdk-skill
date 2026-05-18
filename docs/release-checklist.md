# Release Checklist

Use this checklist before publishing a new version.

## Skill Content

- [ ] `skills/copilot-sdk/SKILL.md` is concise and trigger-rich.
- [ ] Root `SKILL.md` installs with `npx skills add owner/repo` and points to the canonical skill body.
- [ ] Domain details live in `references/`, not the main skill body.
- [ ] The verified API ledger has a checked date and current official URLs for Go, TypeScript, and Python.
- [ ] Workflow playbooks exist for code review, code patching, MCP-backed agents, BYOK backends, and skill-loaded custom agents.
- [ ] Minimal starters exist for Go, TypeScript, and Python and match the API ledger.
- [ ] New recurring behavior standards have eval coverage.
- [ ] Examples are either verified against the current SDK or clearly marked as shape references.
- [ ] Version in `SKILL.md` matches `CHANGELOG.md`.

## Validation

- [ ] `python scripts/validate_skill.py skills/copilot-sdk` passes.
- [ ] `python scripts/check_upstream_sources.py` passes.
- [ ] `cd eval-harness && go test ./...` passes.
- [ ] CI passes on the release branch.
- [ ] Links to official Copilot SDK sources still resolve.
- [ ] Eval prompts are valid JSON and have expectations.

## Evaluation

- [ ] Copilot SDK eval harness outputs exist for at least eight evals.
- [ ] At least three runs are compared before claiming top-tier benchmark quality.
- [ ] With-skill outputs materially improve over baseline.
- [ ] Any failed expectations are explained.
- [ ] Raw `eval-results/` directories are not committed.
- [ ] `aggregate-grading.json` exists for the latest meaningful run.
- [ ] `docs/benchmark-summary.md` reflects the latest meaningful manual read.
- [ ] Release notes honestly state benchmark coverage.

## Publishing

- [ ] Repository is public.
- [ ] Release gate: install command works on a clean machine or clean temp workspace:

```bash
npx skills add https://github.com/ernsahin/copilot-sdk-skill
npx skills add https://github.com/ernsahin/copilot-sdk-skill --skill copilot-sdk
```

- [ ] README includes purpose, install command, quality standard, and source policy.
- [ ] README states that Go, TypeScript, and Python are deep Builder Kit languages and .NET/Java/Rust are source-map only unless verified.
- [ ] Git tag created for the release.
