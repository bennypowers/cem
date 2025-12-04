---
title: Getting Started
layout: docs
weight: 10
---

This guide will walk you through setting up `cem serve` for your custom element project.

## Prerequisites

Before using `cem serve`, you need:

1. **Custom elements manifest** - Generate one with [`cem generate`](/docs/commands/generate/)
2. **Demo files** - HTML files in `demo/` directories or named `*.demo.html`

## The Workflow

`cem serve` follows a manifest-driven development workflow:

1. **Write your component** with JSDoc documentation
2. **Generate manifest** with `cem generate`
3. **Add demos** in `demo/` directories (automatically discovered)
4. **Run serve** with `cem serve`
5. **Iterate** - edit, test, and changes trigger live reload

## Quick Start

### 1. Install cem

```sh
npm install --save-dev @pwrs/cem
```

### 2. Generate a manifest

```sh
cem generate
```

This creates `custom-elements.json` with your component metadata.

### 3. Add a demo file

Create a demo HTML file that uses your component. Demos are HTML partials (not full pages) - the dev server wraps them in demo chrome:

**`demo/index.html`**:
```html
<my-element>
  Hello from my custom element!
</my-element>

<script type="module">
  import '@my-package/my-element.js';
</script>
```

### 4. Link demos to your component

The dev server will automatically discover your demo file since it's in the `demo/` directory - no additional configuration needed!

See [Demo File Discovery](#demo-file-discovery) for details on how demos are found.

{{<tip "info">}}
For explicit control over demo descriptions and ordering, use the `@demo` JSDoc tag in your component. See the [generate docs](/docs/commands/generate/) for details.
{{</tip>}}

### 5. Run the dev server

```sh
cem serve
```

The server starts on `http://localhost:8000` and opens your browser to the element listing page.

## Live Reload

Changes to your source files, demos, or manifest automatically refresh your browser. No manual reloads needed - just save and see your changes instantly.

**Disable reload** with `--no-reload` flag if needed.

## Error Overlay

When errors occur in your code, `cem serve` displays a developer-friendly overlay with:
- Full error message and stack trace
- Source-mapped locations (when using TypeScript)
- Syntax-highlighted code context

Click outside the overlay or press Escape to dismiss.

## Demo File Discovery

`cem serve` discovers demos in two ways:

1. **Manifest `demos` field** - Explicit demos listed in your manifest (recommended)
2. **File pattern matching** - Automatically finds files matching:
   - `demo/` or `demos/` directories
   - `*.demo.html` files

The manifest approach is more explicit and gives you control over demo descriptions and ordering.

{{/* TODO: Add link to example project */}}

## What's Next?

- **[Knobs](knobs/)** - Interactive controls for testing your components
- **[Buildless Development](buildless/)** - Write TypeScript and import CSS without build steps
- **[Import Maps](import-maps/)** - Use npm packages without bundling
- **[Configuration](/docs/configuration/)** - Configuration reference
