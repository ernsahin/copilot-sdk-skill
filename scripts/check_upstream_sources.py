#!/usr/bin/env python3
from __future__ import annotations

import re
import sys
import urllib.request
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
LEDGER = ROOT / "skills" / "copilot-sdk" / "references" / "verified-api-ledger.md"
URL_RE = re.compile(r"https://raw\.githubusercontent\.com/github/copilot-sdk/[^\s)]+")


def fail(message: str) -> None:
    print(f"ERROR: {message}", file=sys.stderr)
    raise SystemExit(1)


def fetch_head(url: str) -> str:
    req = urllib.request.Request(url, headers={"User-Agent": "copilot-sdk-source-check"})
    with urllib.request.urlopen(req, timeout=20) as response:
        data = response.read(2048)
    return data.decode("utf-8", errors="replace")


def main() -> None:
    if not LEDGER.exists():
        fail(f"Missing API ledger: {LEDGER}")

    text = LEDGER.read_text(encoding="utf-8")
    if "checked:" not in text:
        fail("API ledger missing checked metadata")
    urls = sorted(set(URL_RE.findall(text)))
    if not urls:
        fail("API ledger has no official raw GitHub source URL")
    for url in urls:
        head = fetch_head(url)
        if not head.strip():
            fail(f"Source returned empty content: {url}")

    print(f"Checked {len(urls)} upstream Copilot SDK source URLs.")


if __name__ == "__main__":
    main()
