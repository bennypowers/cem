---
title: Buildless Development
layout: docs
weight: 20
---

Write TypeScript and import CSS directly in your browser - no build step required. `cem serve` lets you develop without running `tsc`, CSS watchers, or bundlers.

## Write TypeScript Directly

Import TypeScript files directly in your demos without compilation:

```html
<!-- In your demo -->
<script type="module" src="../src/my-element.ts"></script>
```

No need to:
- Run `tsc --watch`
- Set up a build pipeline
- Wait for compilation before testing
- Manage separate source and output directories

The dev server handles TypeScript on-demand with full source map support for debugging.

### Browser Compatibility

Control which browsers your code supports using the `--target` flag:

```sh
cem serve --target es2020
```

The default target is `es2022`. See [esbuild's target documentation](https://esbuild.github.io/api/#target) for all available targets.

### Your tsconfig.json Works

The dev server respects your existing TypeScript configuration:
- `compilerOptions.target`
- `compilerOptions.module`
- `include` and `exclude` patterns
- Path mappings (via import maps)

{{<tip "info">}}
Command-line `--target` flag overrides `tsconfig.json` settings.
{{</tip>}}

## Import CSS as Modules

Import CSS files directly in your components using modern CSS module syntax:

```js
import styles from './my-element.css';

class MyElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.adoptedStyleSheets = [styles];
  }
}
```

No need to:
- Run CSS preprocessors or watchers
- Bundle stylesheets separately
- Use inline `<style>` tags
- Write CSS-in-JS template literals

Your imported CSS becomes a [Constructable Stylesheet](https://web.dev/constructable-stylesheets/) automatically, giving you better performance than inline styles.

### Opt-In for Safety

CSS module transformation is opt-in to protect traditional `<link>` stylesheets. Specify which CSS files to transform:

**Via command line:**
```sh
# Transform component CSS only
cem serve --css-transform 'src/**/*.css' --css-transform 'elements/**/*.css'
```

**Via config file:**
```yaml
serve:
  transforms:
    css:
      include:
        - 'src/**/*.css'
        - 'elements/**/*.css'
      exclude:
        - 'demo/**/*.css'
```

{{<tip "warning">}}
Without include patterns, **no CSS files are transformed**. This prevents `<link rel="stylesheet">` tags from breaking unexpectedly.
{{</tip>}}

### Example: CSS Module Import

**Your CSS file:**
```css
/* my-element.css */
:host {
  display: block;
  color: var(--text-color, #333);
}
```

**Import it:**
```js
import styles from './my-element.css';

class MyElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.adoptedStyleSheets = [styles];
  }
}
```

**What you get:**
A `CSSStyleSheet` object ready to use with `adoptedStyleSheets` - no build step, no runtime CSS parsing.

## What Works Without Building

You can use directly:
- TypeScript files (`.ts`, `.tsx`)
- CSS files as modules (with opt-in patterns)
- ES modules from npm via import maps
- Modern JavaScript syntax

You still write as-is:
- JavaScript files (`.js`, `.mjs`)
- HTML files
- JSON files

## Debugging

Source maps work automatically:
- **Stack traces** point to your original TypeScript
- **Browser DevTools** show your source files
- **Breakpoints** work in TypeScript, not generated JavaScript

## Configuration

### Via Command Line

```sh
cem serve --target es2020 --css-transform 'src/**/*.css'
```

### Via Config File

```yaml
serve:
  target: es2020
  transforms:
    css:
      include:
        - 'src/**/*.css'
```

See **[Configuration](/docs/configuration/)** for all options.

## What's Next?

- **[Import Maps](import-maps/)** - Use npm packages without bundling
- **[Configuration](/docs/configuration/)** - Configuration reference
- **[Getting Started](getting-started/)** - Set up your first demo
