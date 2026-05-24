# Benchmark Summary

This repository uses GitHub Copilot SDK itself to run paired behavior evals.

The eval harness compares:

1. Baseline: same custom eval agent with no skill preloaded.
2. With skill: same custom eval agent with `Skills: ["copilot-sdk-kit"]` and this repository's `skills/` directory.

Raw eval outputs are intentionally ignored by git under `eval-results/`. They contain model transcripts, event logs, permission logs, aggregate grading summaries, and run metadata that are useful locally but too noisy for the published skill package.

## Latest Manual Read

Recorded during local iteration on 2026-05-18.

Run shape:

1. Runner: `eval-harness/`, using GitHub Copilot SDK Go package.
2. Conditions: paired baseline and with-skill sessions using the same custom eval agent.
3. Baseline: no skill preloaded.
4. With skill: `Skills: ["copilot-sdk-kit"]` and this repository's `skills/` directory.
5. Grading: manual expectation review against saved `response.md`, `events.jsonl`, `permissions.jsonl`, and `grading.json`.
6. Raw outputs: intentionally not committed.

The strongest full manual read showed:

| Condition | Score |
| --- | ---: |
| Baseline | 8 / 16 |
| With skill | 13 / 16 |

This proves useful behavior lift, not final excellence.

The score should be read as directional evidence. It does not replace repeated runs, blind review, or automated aggregation.

## Targeted Fixes After That Run

Two weak areas were isolated and patched:

1. Code-review agent prompts were too likely to start with clarification questions.
2. Source-sensitive Go SDK guidance could still present exact APIs without true source verification.

After targeted iteration:

1. Reviewer design prompts now start with assumptions and a minimum architecture instead of a questionnaire.
2. Unverified SDK guidance now stays conceptual and avoids exact setup code.
3. Example files were downgraded to shape references so they cannot be treated as source verification.

## Top-Tier Claim Gate

Do not claim top-tier benchmark quality until all of these are true:

1. At least three runs are completed.
2. At least eight evals are graded per run.
3. Go, TypeScript, and Python implementation evals are included.
4. MCP and BYOK workflow evals are included.
5. The negative unsupported-language eval passes.
6. With-skill output materially beats baseline across architecture, implementation, source verification, and workflow categories.
7. Failed expectations are classified and either fixed or documented as out of scope.

## Remaining Honest Limits

Before claiming top-tier benchmark quality:

1. Run the expanded eval set across multiple attempts to account for model variance.
2. Use `aggregate-grading.json` for manual grade summaries.
3. Add trigger-selection evals separate from behavior-quality evals.
4. Keep release notes explicit about what was verified and what remains qualitative.
