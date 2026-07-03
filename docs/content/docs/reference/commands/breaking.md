---
title: Breaking
description: Detect breaking API changes between two custom-elements manifests
---

{{< tip >}}
**TL;DR**: Run `cem breaking` to detect breaking API changes between manifest versions. Use `--fail-on breaking` in CI to block breaking changes, `--format markdown` for PR comments.
{{< /tip >}}

The `cem breaking` command compares two `custom-elements.json` manifests and detects breaking changes in component APIs. Changes are classified as **breaking**, **dangerous**, or **safe**.

```bash
cem breaking [old-manifest.json new-manifest.json] [flags]
```

## Options

| Flag | Type | Description |
| ---- | ---- | ----------- |
| `--base` | string | Git ref for baseline (default: latest semver tag) |
| `--head` | string | Git ref for head (default: working tree manifest) |
| `--format` | string | Output format: `text` (default), `json`, or `markdown` |
| `--fail-on` | string | Exit 1 if changes at this severity or above: `breaking` or `dangerous` |
| `--disable` | string (repeatable) | Disable specific breaking change rules |
| `--package`, `-p` | string | Path to a package directory |

## Severity Classification

| Severity | Meaning | Examples |
|----------|---------|----------|
| **Breaking** | Downstream consumers will break | Removed element, removed attribute, changed type |
| **Dangerous** | May break depending on usage | Changed default value, narrowed type |
| **Safe** | Non-breaking additions | New element, new attribute, new slot |

## Detection Rules

| Rule ID | Severity | Detects |
|---------|----------|---------|
| `element-removed` | Breaking | Tag name removed from manifest |
| `element-added` | Safe | New tag name added |
| `attribute-removed` | Breaking | Attribute removed from element |
| `attribute-added` | Safe | New attribute added |
| `attribute-type-changed` | Breaking | Attribute type changed |
| `attribute-default-changed` | Dangerous | Attribute default value changed |
| `slot-removed` | Breaking | Named slot removed |
| `slot-added` | Safe | New slot added |
| `css-custom-property-removed` | Breaking | CSS custom property removed |
| `css-custom-property-added` | Safe | New CSS custom property |
| `css-custom-property-default-changed` | Dangerous | CSS property default changed |
| `css-part-removed` | Breaking | CSS part removed |
| `css-part-added` | Safe | New CSS part added |
| `css-state-removed` | Breaking | CSS state removed |
| `css-state-added` | Safe | New CSS state added |
| `event-removed` | Breaking | Event removed |
| `event-added` | Safe | New event added |
| `event-type-changed` | Breaking | Event type changed |
| `method-removed` | Breaking | Public method removed |
| `method-added` | Safe | New public method |
| `method-return-type-changed` | Breaking | Method return type changed |
| `method-parameter-changed` | Breaking | Method parameters changed |
| `field-removed` | Breaking | Public field removed |
| `field-type-changed` | Dangerous | Public field type changed |

## Examples

### Compare against latest git tag

```bash
cem breaking
```

### Compare against specific git ref

```bash
cem breaking --base v1.0.0
```

### Compare two manifest files directly

```bash
cem breaking old-manifest.json new-manifest.json
```

### CI mode: fail on breaking changes

```bash
cem breaking --base v1.0.0 --fail-on breaking
```

### Machine-readable output

```bash
cem breaking --base v1.0.0 --format json
```

### Markdown for PR comments

```bash
cem breaking --base v1.0.0 --format markdown
```

### Disable specific rules

```bash
cem breaking --disable css-custom-property-removed --disable slot-removed
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | No breaking changes (or `--fail-on` not triggered) |
| 1 | Breaking changes detected (when `--fail-on` is set) |
| 2 | Input error (missing files, invalid manifests, etc.) |

## Configuration

Rules can be disabled via `.cem.yaml`:

```yaml
breaking:
  disable:
    - css-custom-property-removed
    - slot-removed
```

CLI `--disable` flags are merged with config-file rules.
