---
name: docs-health
description: >
  Audit custom element manifest documentation for completeness and quality.
  Use when the user asks to "check docs", "audit documentation",
  "review manifest quality", "find missing descriptions", "docs health",
  or mentions "undocumented" attributes, slots, events, or CSS properties.
tools: Read, Bash
---

# Documentation Health Check

Audit custom element manifests for documentation completeness and quality using `cem health`.

## Workflow

### Phase 1: Run Health Analysis

Run `cem health` with JSON output to get structured scoring data:

```bash
cem health --format json
```

To filter to specific components or modules:

```bash
# Single component
cem health --format json --component my-button

# Specific modules
cem health --format json --module src/components/button.js
```

The JSON output has this structure:

```json
{
  "modules": [
    {
      "path": "src/components/button.js",
      "score": 18,
      "maxScore": 20,
      "declarations": [
        {
          "tagName": "my-button",
          "name": "MyButton",
          "score": 18,
          "maxScore": 20,
          "categories": [
            {
              "id": "description",
              "category": "Description",
              "points": 5,
              "maxPoints": 5,
              "status": "pass",
              "findings": [...]
            },
            {
              "id": "attributes",
              "category": "Attribute Documentation",
              "points": 3,
              "maxPoints": 5,
              "status": "warn",
              "findings": [
                {
                  "check": "attribute 'variant' has description",
                  "points": 1,
                  "max": 1
                },
                {
                  "check": "attribute 'size' has description",
                  "points": 0,
                  "max": 1,
                  "message": "Add a description for attribute 'size'"
                }
              ]
            }
          ]
        }
      ]
    }
  ],
  "overallScore": 18,
  "overallMax": 20,
  "recommendations": [
    "my-button: Add a description for attribute 'size' (Attribute Documentation, +1 pts)"
  ]
}
```

### Phase 2: Interpret Results

The health categories scored by `cem health` are:

| Category | What it checks |
|----------|---------------|
| `description` | Element has a description and summary |
| `attributes` | Every attribute has a description, type, and default |
| `slots` | Every slot has a description |
| `css` | CSS custom properties, parts, and states have descriptions |
| `events` | Every event has a description and type |
| `demos` | Element has demo links |

Status values:
- **pass** (>=80%): Well documented
- **warn** (40-79%): Has gaps
- **fail** (<40%): Major gaps

### Phase 3: Report

Present the results in a readable format. Use the JSON data to build:

```markdown
## Documentation Health Report

### Summary
- Components audited: X
- Overall score: X/Y (Z%)
- Components passing: X
- Components needing attention: X

### Per-Component Scores

| Component | Score | Description | Attrs | Slots | Events | CSS | Demos |
|-----------|-------|-------------|-------|-------|--------|-----|-------|
| `<my-button>` | 90% | pass | pass | pass | warn | pass | fail |

### Findings

#### `<my-button>` (18/20, 90%)
- **Attribute Documentation** (warn): Add a description for attribute 'size'
- **Demos** (fail): No demo links found

### Top Recommendations
[List the recommendations from the JSON output — these are already
sorted by point impact]
```

### Phase 4: Deeper Investigation (Optional)

If the user wants more detail on specific gaps, read the element's MCP resources:

```
cem://element/{tagName}
cem://element/{tagName}/attributes
cem://element/{tagName}/slots
```

This provides the full manifest data so you can show exactly what's missing and suggest what to document.

## Guidelines

- **Use `cem health` as the source of truth**: Don't reinvent the scoring — it already checks descriptions, types, defaults, etc.
- **Surface recommendations first**: The `recommendations` array is sorted by impact — present those prominently
- **Suggest, don't fabricate**: Flag gaps for the element author to fill, don't invent descriptions
- **Flag inconsistencies**: If some elements score 100% and others score 30%, highlight the gap
- **Use `--component` for focused audits**: When the user asks about a specific element, filter to avoid noise
