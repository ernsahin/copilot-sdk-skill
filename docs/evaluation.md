# Evaluation Runbook

The skill includes eval prompts in `skills/copilot-sdk/evals/evals.json`.

The goal is to verify that the skill changes agent behavior, not just that the files are valid.

## Current Status

- Eval prompts: defined.
- Objective expectations: defined.
- Baseline vs with-skill runs: pending.
- CI smoke validation: enabled through `scripts/validate_skill.py`.

The pending status is intentional. A real benchmark requires paired agent runs and human review of outputs. Do not mark this skill as fully benchmarked until those runs exist.

## Evaluation Standard

For each eval:

1. Run the prompt without this skill.
2. Run the same prompt with `skills/copilot-sdk`.
3. Save both outputs.
4. Grade each expectation as pass/fail with evidence.
5. Compare whether the skill materially improves the output.
6. Patch the skill only when the failure reveals a reusable rule or reference gap.

## Expected Improvements

The with-skill output should:

- Reject hardcoded demo behavior when a reusable workflow is needed.
- Treat SDLC as executable workflow governance, not folder scaffolding.
- Separate runtime host, workflows, agents, tools, permissions, state, integrations, and observability.
- Require upstream verification before version-sensitive Copilot SDK guidance.
- Use risk-based validation instead of requiring tests everywhere.
- Prefer SDK and platform primitives before custom infrastructure.

## Result Format

Store benchmark results under:

```text
eval-results/
  YYYY-MM-DD-iteration-1/
    eval-01/
      baseline.md
      with-skill.md
      grading.json
```

Use this grading shape:

```json
{
  "eval_id": 1,
  "baseline_passed": 1,
  "with_skill_passed": 4,
  "expectations": [
    {
      "text": "The output rejects hardcoded outcomes when future inputs are implied.",
      "baseline": false,
      "with_skill": true,
      "evidence": "The with-skill output explicitly chose a reusable workflow mechanism."
    }
  ],
  "notes": "Patch the skill only if with-skill misses a recurring standard."
}
```

## Publishing Gate

Before a public release:

1. At least four evals must have baseline and with-skill outputs.
2. With-skill output should pass more expectations than baseline on each run.
3. Any failed expectation must be classified as one of:
   - Skill wording gap.
   - Eval expectation too broad.
   - User prompt intentionally out of scope.
   - Upstream SDK ambiguity.
4. The release notes must mention evaluation status honestly.

