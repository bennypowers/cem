---
name: write-demos
description: >
  Write demo HTML files for custom elements that work with cem serve.
  Use when the user asks to "write a demo", "create a demo", "add a demo",
  "demo file", "demo page", or mentions "cem serve demos".
tools: Read, Glob, Grep, generate_html
---

# Write Element Demos

## Critical: Demos Are Partials

Never include `<!DOCTYPE>`, `<html>`, `<head>`, or `<body>`.
The dev server wraps partials with its own document shell, import maps,
and live reload. Adding document structure breaks rendering.

## Workflow

### 1. Gather Context

Read the element's manifest data (`cem://element/{tagName}` and sub-resources)
and check `.config/cem.yaml` for:

- **`generate.demoDiscovery`** — `fileGlob`, `urlPattern`, `urlTemplate`
  determine where demos live and how URLs are generated
- **`serve.urlRewrites`** (if present) — maps request URLs to filesystem paths;
  when configured, imports must use rewritten paths, not raw filesystem paths
- **`serve.importMap`** — bare specifiers need to be in `dependencies` or the
  import map config

### 2. Write Demos

Place files where `fileGlob` expects them (typically `elements/my-el/demo/`).

**Index demo** — simplest possible usage, no metadata needed:

```html
<script type="module">
  import '../my-button.js';
</script>

<my-button>Click Me</my-button>
```

**Feature demos** — one feature per file, with metadata:

```html
<meta itemprop="name"
      content="With Icon">
<meta itemprop="description"
      content="Button with a leading icon in the icon slot">

<script type="module">
  import '../my-button.js';
  import '../../my-icon/my-icon.js';
</script>

<my-button variant="primary">
  <my-icon slot="icon"
           name="check"></my-icon>
  Save Changes
</my-button>
```

**Multi-element demos** — use `demo-for` to associate with specific elements:

```html
<meta itemprop="demo-for"
      content="my-button my-icon">
```

### 3. Validate

Call `validate_html` on the demo content, then verify:
- No document shell tags
- Script imports resolve (relative paths from demo file to source)
- `<meta>` tags use `itemprop`, not `name` or `property`

## Common Mistakes

| Mistake | Fix |
|---------|-----|
| Adding `<!DOCTYPE>` / `<html>` | Demos are partials |
| `<meta name="description">` | `<meta itemprop="description">` |
| Absolute import paths | Relative from demo file to source |
| Importing from `devDependencies` | Move to `dependencies` |
| External CSS/JS links | Inline styles and scripts |
| Everything in one file | One feature per demo |
| Missing `type="module"` | Always `<script type="module">` |

## Guidelines

- **Index is minimal**: default demo should be copy-paste simple
- **Self-contained**: inline all CSS and JS
- **One feature per file**: variants, states, features get separate demos
- **Use `generate_html`**: produce valid element markup via the MCP tool
