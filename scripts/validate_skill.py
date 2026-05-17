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


def main() -> None:
    if len(sys.argv) != 2:
        fail("Usage: validate_skill.py <skill-dir>")
    validate_skill(Path(sys.argv[1]).resolve())
    print("Skill package is valid.")


if __name__ == "__main__":
    main()

