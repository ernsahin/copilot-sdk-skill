# Contributing

Contributions should improve agent behavior, not only add content.

## Before Changing The Skill

1. Identify the failing behavior or missing decision standard.
2. Decide whether the fix belongs in `SKILL.md` or a reference file.
3. Prefer a general quality gate over a domain-specific rule when it covers the same issue.
4. Add or update an eval when the change prevents a recurring failure.

## Local Checks

```bash
python scripts/validate_skill.py skills/copilot-sdk-kit
```

## Writing Rules

- Keep `SKILL.md` concise.
- Move SDK details, examples, and gotchas to `references/`.
- Do not add long persona prompts.
- Do not add rules that only solve one example when a general principle covers the case.
- Do not claim stable API behavior without upstream verification.

