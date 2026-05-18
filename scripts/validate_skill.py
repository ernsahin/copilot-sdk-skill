#!/usr/bin/env python3
from __future__ import annotations

import json
import re
import sys
from pathlib import Path


PLACEHOLDER_PATTERNS = [
    r"\bTBD\b",
    r"\bTODO\b",
    r"implement later",
    r"fill in details",
]


def fail(message: str) -> None:
    print(f"ERROR: {message}", file=sys.stderr)
    raise SystemExit(1)


def parse_frontmatter(text: str) -> tuple[dict[str, str], str]:
    if not text.startswith("---\n"):
        fail("SKILL.md must start with YAML frontmatter")
    end = text.find("\n---\n", 4)
    if end == -1:
        fail("SKILL.md frontmatter is not closed")
    raw = text[4:end]
    body = text[end + 5 :]
    data: dict[str, str] = {}
    current_key: str | None = None
    for line in raw.splitlines():
        if not line.strip():
            continue
        if line.startswith("  ") and current_key:
            continue
        match = re.match(r"^([A-Za-z0-9_-]+):\s*(.*)$", line)
        if match:
            key, value = match.groups()
            current_key = key
            data[key] = value.strip().strip('"')
    return data, body


def validate_skill(skill_dir: Path) -> None:
    skill_md = skill_dir / "SKILL.md"
    if not skill_md.exists():
        fail(f"Missing {skill_md}")

    text = skill_md.read_text(encoding="utf-8")
    frontmatter, body = parse_frontmatter(text)

    name = frontmatter.get("name")
    description = frontmatter.get("description")
    if not name:
        fail("Missing frontmatter name")
    if not re.match(r"^[a-z0-9](?:[a-z0-9-]{0,62}[a-z0-9])?$", name):
        fail(f"Invalid skill name: {name}")
    if name != skill_dir.name:
        fail(f"Skill name {name!r} must match directory {skill_dir.name!r}")
    if not description:
        fail("Missing frontmatter description")
    if len(description) > 1024:
        fail("Description must be 1024 characters or fewer")
    if "Use this skill" not in description and "Use when" not in description:
        fail("Description must include trigger guidance")

    if len(body.splitlines()) > 500:
        fail("SKILL.md body should stay under 500 lines")

    for pattern in PLACEHOLDER_PATTERNS:
        if re.search(pattern, text, flags=re.IGNORECASE):
            fail(f"Placeholder language found in SKILL.md: {pattern}")

    validate_references(skill_dir, text)
    validate_examples(skill_dir)
    validate_builder_kit(skill_dir)
    validate_evals(skill_dir, name)


def validate_references(skill_dir: Path, skill_text: str) -> None:
    references_dir = skill_dir / "references"
    if not references_dir.exists():
        fail("Missing references directory")

    referenced = re.findall(r"`(references/[^`]+\.md)`", skill_text)
    if not referenced:
        fail("SKILL.md should link to reference files")
    for rel in referenced:
        if not (skill_dir / rel).exists():
            fail(f"Referenced file does not exist: {rel}")

    sections = references_dir / "_sections.md"
    if not sections.exists():
        fail("Missing references/_sections.md")

    required_references = [
        "copilot-sdk-rules.md",
        "workflow-routing.md",
        "stop-conditions.md",
        "source-verification.md",
        "known-gotchas.md",
        "agent-product-lifecycle.md",
    ]
    for name in required_references:
        path = references_dir / name
        if not path.exists():
            fail(f"Missing required reference: references/{name}")
        text = path.read_text(encoding="utf-8")
        if "impact:" not in text:
            fail(f"Reference must declare impact metadata: references/{name}")


def validate_builder_kit(skill_dir: Path) -> None:
    workflow_dir = skill_dir / "references" / "workflows"
    starters_dir = skill_dir / "examples" / "starters"

    ledger = skill_dir / "references" / "verified-api-ledger.md"
    if not ledger.exists():
        fail("Missing verified API ledger: references/verified-api-ledger.md")
    ledger_text = ledger.read_text(encoding="utf-8")
    for required in ["checked:", "Sources Checked", "## Go", "## TypeScript", "## Python"]:
        if required not in ledger_text:
            fail(f"Verified API ledger missing {required!r}")
    if "raw.githubusercontent.com/github/copilot-sdk" not in ledger_text:
        fail("Verified API ledger must cite official upstream raw URLs")

    required_workflows = [
        "code-reviewer-workflow.md",
        "code-patcher-workflow.md",
        "mcp-backed-agent-workflow.md",
        "byok-backend-workflow.md",
        "skill-loaded-custom-agent-workflow.md",
    ]
    for name in required_workflows:
        path = workflow_dir / name
        if not path.exists():
            fail(f"Missing workflow playbook: references/workflows/{name}")
        text = path.read_text(encoding="utf-8")
        for required in ["impact:", "Runtime Shape", "Workflow", "Failure"]:
            if required not in text:
                fail(f"Workflow missing {required!r}: {path.relative_to(skill_dir)}")

    required_starters = {
        "go": ["README.md", "main.go"],
        "typescript": ["README.md", "index.ts"],
        "python": ["README.md", "main.py"],
    }
    for starter, files in required_starters.items():
        starter_dir = starters_dir / starter
        if not starter_dir.exists():
            fail(f"Missing starter directory: examples/starters/{starter}")
        readme = (starter_dir / "README.md").read_text(encoding="utf-8")
        if "Verified against API ledger:" not in readme:
            fail(f"Starter README must declare API ledger freshness: examples/starters/{starter}/README.md")
        for file_name in files:
            if not (starter_dir / file_name).exists():
                fail(f"Missing starter file: examples/starters/{starter}/{file_name}")


def validate_examples(skill_dir: Path) -> None:
    examples_dir = skill_dir / "examples"
    if not examples_dir.exists():
        fail("Missing examples directory")
    examples = list(examples_dir.glob("*.md"))
    if not examples:
        fail("At least one example markdown file is required")
    for example in examples:
        text = example.read_text(encoding="utf-8")
        if "Verify upstream" not in text and "shape reference" not in text:
            fail(f"Example must declare freshness policy: {example.name}")
        if "not source verification" not in text.lower():
            fail(f"Example must state it is not source verification: {example.name}")
        if "```go" in text:
            fail(f"Example must not include exact Go setup code: {example.name}")


def validate_evals(skill_dir: Path, skill_name: str) -> None:
    evals_path = skill_dir / "evals" / "evals.json"
    if not evals_path.exists():
        fail("Missing evals/evals.json")
    try:
        data = json.loads(evals_path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as exc:
        fail(f"Invalid evals JSON: {exc}")

    if data.get("skill_name") != skill_name:
        fail("evals.skill_name must match frontmatter name")
    evals = data.get("evals")
    if not isinstance(evals, list) or len(evals) < 4:
        fail("At least four evals are required")

    seen_ids: set[int] = set()
    for item in evals:
        eval_id = item.get("id")
        if not isinstance(eval_id, int):
            fail("Each eval must have an integer id")
        if eval_id in seen_ids:
            fail(f"Duplicate eval id: {eval_id}")
        seen_ids.add(eval_id)

        prompt = item.get("prompt")
        expected = item.get("expected_output")
        expectations = item.get("expectations")
        if not isinstance(prompt, str) or len(prompt.strip()) < 20:
            fail(f"Eval {eval_id} prompt is missing or too short")
        if not isinstance(expected, str) or len(expected.strip()) < 20:
            fail(f"Eval {eval_id} expected_output is missing or too short")
        if not isinstance(expectations, list) or len(expectations) < 3:
            fail(f"Eval {eval_id} must have at least three expectations")

    required_eval_phrases = [
        "hardcoded",
        "runtime host",
        "SDLC lifecycle",
        "approve all",
        "TypeScript",
        "Python",
        "API ledger",
        "MCP",
        "BYOK",
    ]
    eval_text = "\n".join(
        "\n".join(
            [
                item.get("prompt", ""),
                item.get("expected_output", ""),
                "\n".join(item.get("expectations", [])),
            ]
        )
        for item in evals
    )
    for phrase in required_eval_phrases:
        if phrase.lower() not in eval_text.lower():
            fail(f"Evals must cover phrase/theme: {phrase}")


def main() -> None:
    if len(sys.argv) != 2:
        fail("Usage: validate_skill.py <skill-dir>")
    skill_dir = Path(sys.argv[1]).resolve()
    validate_skill(skill_dir)
    repo_root = skill_dir.parent.parent
    validate_root_entrypoint(repo_root)
    validate_eval_harness(repo_root)
    print("Skill package is valid.")


def validate_root_entrypoint(repo_root: Path) -> None:
    root_skill = repo_root / "SKILL.md"
    if not root_skill.exists():
        fail("Missing repository-root SKILL.md entrypoint")
    text = root_skill.read_text(encoding="utf-8")
    frontmatter, body = parse_frontmatter(text)
    if frontmatter.get("name") != "copilot-sdk":
        fail("Root SKILL.md name must be copilot-sdk")
    description = frontmatter.get("description", "")
    if "Use this skill" not in description:
        fail("Root SKILL.md description must include trigger guidance")
    if "skills/copilot-sdk/SKILL.md" not in body:
        fail("Root SKILL.md must point to the canonical skill body")


def validate_eval_harness(repo_root: Path) -> None:
    harness_dir = repo_root / "eval-harness"
    if not harness_dir.exists():
        fail("Missing eval-harness directory")
    for rel in ["go.mod", "main.go", "README.md"]:
        if not (harness_dir / rel).exists():
            fail(f"Missing eval harness file: eval-harness/{rel}")
    gitignore = repo_root / ".gitignore"
    if not gitignore.exists() or "eval-results/" not in gitignore.read_text(encoding="utf-8"):
        fail(".gitignore must exclude eval-results/")


if __name__ == "__main__":
    main()
