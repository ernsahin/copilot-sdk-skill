# Release Checklist

Use this checklist before publishing a new version.

## Skill Content

- [ ] `skills/copilot-sdk/SKILL.md` is concise and trigger-rich.
- [ ] Domain details live in `references/`, not the main skill body.
- [ ] New recurring behavior standards have eval coverage.
- [ ] Examples are either verified against the current SDK or clearly marked as shape references.
- [ ] Version in `SKILL.md` matches `CHANGELOG.md`.

## Validation

- [ ] `python scripts/validate_skill.py skills/copilot-sdk` passes.
- [ ] CI passes on the release branch.
- [ ] Links to official Copilot SDK sources still resolve.
- [ ] Eval prompts are valid JSON and have expectations.

## Evaluation

- [ ] Baseline vs with-skill outputs exist for at least four evals.
- [ ] With-skill outputs materially improve over baseline.
- [ ] Any failed expectations are explained.
- [ ] Release notes honestly state benchmark coverage.

## Publishing

- [ ] Repository is public.
- [ ] Install command works:

```bash
npx skills add https://github.com/ernsahin/copilot-sdk-skill --skill copilot-sdk
```

- [ ] README includes purpose, install command, quality standard, and source policy.
- [ ] Git tag created for the release.

