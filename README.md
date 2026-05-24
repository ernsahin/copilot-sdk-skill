# ernsahin Skills

[![skills.sh](https://skills.sh/b/ernsahin/skills)](https://skills.sh/ernsahin/skills)
[![CI](https://github.com/ernsahin/skills/actions/workflows/ci.yml/badge.svg)](https://github.com/ernsahin/skills/actions/workflows/ci.yml)

Skills live under `skills/`. There is no root `SKILL.md`.

## Skills

| Skill | Install |
| --- | --- |
| `copilot-sdk-kit` | `npx skills add https://github.com/ernsahin/skills --skill copilot-sdk-kit` |

## Checks

```bash
python3 scripts/validate_skill.py skills/copilot-sdk-kit
python3 scripts/check_upstream_sources.py
cd eval-harness && go test ./...
```

## License

MIT. See [LICENSE](LICENSE).
