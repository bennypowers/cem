---
title: Health
description: Score documentation quality in a custom-elements manifest
---

{{< tip >}}
**TL;DR**: Run `cem health` to score how well your custom elements are documented. Use `--format markdown` for CI reports, `--fail-below 60` to enforce a minimum score, and `--disable demos` to skip categories you don't care about.
{{< /tip >}}

The `cem health` command analyzes a `custom-elements.json` manifest and scores the quality of its documentation across several categories. It produces actionable recommendations sorted by potential point gain.

```bash
cem health [path/to/custom-elements.json] [flags]
```

## Options

| Flag | Type | Description |
| ---- | ---- | ----------- |
| `--component` | string | Filter to a specific component by tag name or class name |
| `--module` | string (repeatable) | Filter to specific modules by path (can be used multiple times) |
| `--format` | string | Output format: `text` (default), `json`, or `markdown` |
| `--fail-below` | int | Exit 1 if overall percentage is below this threshold (0-100) |
| `--disable` | string (repeatable) | Disable specific health categories by ID |
| `--package`, `-p` | string | Path to a package directory |

## Scoring Categories

Each declaration is scored out of 100 points across six categories:

| ID | Category | Max Points | What it checks |
| -- | -------- | ---------: | -------------- |
| `description` | Element description | 25 | Has a description, optimal length, explains purpose, RFC 2119 keywords, accessibility notes, keyboard interaction |
| `attributes` | Attribute documentation | 20 | All attributes described, descriptions explain purpose, constraints documented |
| `slots` | Slot documentation | 15 | All slots described, content type guidelines, accessibility considerations |
| `css` | CSS documentation | 15 | CSS custom properties documented, CSS parts documented, design system notes |
| `events` | Event documentation | 15 | All events described, trigger conditions, event detail shape |
| `demos` | Demos | 10 | Has at least one demo, demo has a description |

Status thresholds: **pass** (>=80%), **warn** (>=40%), **fail** (<40%).

## Examples

### Basic usage

```bash
cem health
```

### Score a specific package

```bash
cem health -p packages/my-components
```

### Filter to specific modules

```bash
cem health --module elements/my-button/my-button.js \
           --module elements/my-card/my-card.js
```

### Markdown output for CI

```bash
cem health --format markdown > health_report.md
```

### Fail CI if documentation score is below 60%

```bash
cem health --fail-below 60
```

### Skip the demos category

```bash
cem health --disable demos
```

### JSON output for tooling

```bash
cem health --format json
```

```json
{
  "modules": [
    {
      "path": "elements/data-table/data-table.js",
      "score": 56,
      "maxScore": 100,
      "declarations": [
        {
          "tagName": "data-table",
          "name": "DataTable",
          "score": 56,
          "maxScore": 100,
          "categories": [
            {
              "id": "description",
              "category": "Element description",
              "points": 14,
              "maxPoints": 25,
              "status": "warn",
              "findings": [
                { "check": "Has description", "points": 5, "max": 5 },
                { "check": "Optimal length", "points": 3, "max": 3 },
                { "check": "Explains purpose/context", "points": 3, "max": 5,
                  "message": "describe purpose and context using words like 'for', 'when', 'provides', 'allows'" }
              ]
            }
          ]
        }
      ]
    }
  ],
  "overallScore": 56,
  "overallMax": 100,
  "recommendations": [
    "data-table: use RFC 2119 keywords (MUST, SHOULD, AVOID) to clarify requirements (Element description, +5 pts)"
  ]
}
```

## Configuration

### Configuration File

```yaml
# .cem.yaml
health:
  failBelow: 60
  disable:
    - demos
```

### Command Line

```bash
# Disable categories via flags (merged with config)
cem health --disable demos --disable css

# Override fail-below threshold
cem health --fail-below 80
```

## GitHub Actions

Two reusable workflows are available:

- **`cem-health-report.yml`** — for pull requests. Generates a manifest, runs `cem health` scoped to changed files, and posts a sticky PR comment.
- **`cem-health-dispatch.yml`** — for manual runs. Generates a manifest, runs `cem health` across the whole repo, and posts a job summary.

### PR workflow

```yaml
name: Health Report

on:
  pull_request:
    types: [opened, synchronize, reopened]

jobs:
  health:
    uses: bennypowers/cem/.github/workflows/cem-health-report.yml@main
    with:
      fail-below: 60
```

### Dispatch workflow

```yaml
name: Health Report (Dispatch)

on:
  workflow_dispatch:

jobs:
  health:
    uses: bennypowers/cem/.github/workflows/cem-health-dispatch.yml@main
    with:
      fail-below: 60
```

### Workflow inputs

Both workflows accept the same inputs:

| Input | Type | Default | Description |
| ----- | ---- | ------- | ----------- |
| `package-path` | string | `.` | Path to the package |
| `cem-version` | string | (auto) | cem version (semver, no `v` prefix). If empty, reads from `package.json` or uses latest. Minimum 0.9.5. |
| `fail-below` | number | `0` | Exit 1 if overall percentage is below this threshold |
| `generate-args` | string | | Additional arguments for `cem generate` |
| `health-args` | string | | Additional arguments for `cem health` |

Both workflows download the `cem` binary directly from GitHub releases — no Node.js or `node_modules` required.
