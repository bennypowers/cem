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
- `compilerOptions.rootDir` and `compilerOptions.outDir` (for src/dist separation)
- `include` and `exclude` patterns

{{<tip "info">}}
Command-line `--target` flag overrides `tsconfig.json` settings.
{{</tip>}}

### Separate `dist` Directory

If your project uses separate source and output directories (e.g., `src/` and `dist/`), the dev server automatically handles path resolution so you can reference compiled output paths while serving source files.

**How it works:**

1. The dev server reads your `tsconfig.json` to detect `rootDir` and `outDir`
2. It creates path mappings automatically (e.g., `"/dist/" → "/src/"`)
3. Requests to `/dist/components/button.js` resolve to `/src/components/button.ts`
4. Your demos can reference the output path, but the server serves the source

**Example tsconfig.json:**
```json
{
  "compilerOptions": {
    "rootDir": "./src",
    "outDir": "./dist",
    "target": "ES2022"
  }
}
```

With this configuration:
- Request: `/dist/components/button.js`
- Resolves to: `/src/components/button.ts`
- Server transforms and serves the TypeScript source

**Benefits:**
- Write demos using the same paths your production build uses
- No need to change import paths between development and production
- Works seamlessly with TypeScript's project references
- Supports monorepo/workspace setups with per-package `tsconfig.json` files

**Fallback behavior:**

If the dev server can't find a source file via path mappings, it tries co-located files (in-place compilation). This ensures backward compatibility with projects that compile TypeScript in the same directory as source files.

**Manual configuration:**

If you need custom URL rewrites beyond what `tsconfig.json` provides, configure them using URLPattern syntax and Go templates:

```yaml
serve:
  urlRewrites:
    # Simple prefix mapping
    - urlPattern: "/dist/:path*"
      urlTemplate: "/src/{{.path}}"

    # Custom library paths
    - urlPattern: "/lib/:path*"
      urlTemplate: "/sources/{{.path}}"

    # Pattern with template function
    - urlPattern: "/api/:version/:endpoint*"
      urlTemplate: "/{{.version | lower}}/api/{{.endpoint}}"
```

**Pattern syntax:**

URL patterns use [URLPattern](https://developer.mozilla.org/en-US/docs/Web/API/URL_Pattern_API) syntax:
- `:param` - Matches a single path segment and captures it as a named parameter
- `**` - Matches any number of path segments (wildcard)
- `*` - Matches characters within a single segment

URL templates use Go template syntax:
- `{{.param}}` - Inserts a captured parameter
- `{{.param | lower}}` - Applies the `lower` function (converts to lowercase)
- `{{.param | upper}}` - Applies the `upper` function (converts to uppercase)
- `{{.param | slug}}` - Applies the `slug` function (URL-safe slug)

**Example:** Element name transformation for demos

If your project serves demos at URLs that don't match the on-disk structure:
- **On disk**: `elements/my-card/demo/card.html`
- **Served URL**: `/elements/card/demo/card/`
- **CSS reference in demo**: `<link rel="stylesheet" href="../my-card-lightdom.css">`
- **Browser resolves to**: `/elements/card/demo/my-card-lightdom.css` ❌

This can happen when your demo files load files by relative path from, since the
dev server prettifies their URLs (`card.html` => `card/index.html`).

Configure URL rewrites to fix the mismatch:

```yaml
serve:
  urlRewrites:
    - urlPattern: "/elements/:slug/demo/:rest*.css"
      urlTemplate: "/elements/my-{{.slug}}/{{.rest}}.css"
```

This resolves `/elements/card/demo/my-card-lightdom.css` → `elements/my-card/my-card-lightdom.css` ✓

See **[Configuration](/docs/reference/configuration/)** for details.

## Import CSS as Modules

Import CSS files directly in your components using modern CSS module syntax:

```js
import styles from './my-element.css' with { type: 'css' };

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

### Import Attributes (Modern Syntax)

Use the standard [import attributes](https://github.com/tc39/proposal-import-attributes) syntax with the `with` keyword:

```js
import styles from './my-element.css' with { type: 'css' };

class MyElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.adoptedStyleSheets = [styles];
  }
}
```

This is the modern, standards-based syntax for importing CSS as modules. The dev server automatically handles this syntax during TypeScript transformation - import attributes are preserved and CSS files are transformed to JavaScript modules that export `CSSStyleSheet` objects.

{{<tip "info">}}
Import attributes with `with { type: 'css' }` **bypass** include/exclude patterns. When you use this syntax, the dev server knows you want that CSS file transformed to a module, regardless of glob patterns.
{{</tip>}}

**How it works:**

1. You write: `import styles from './foo.css' with { type: 'css' }`
2. The dev server rewrites it to: `import styles from './foo.css?__cem-import-attrs[type]=css'`
3. When the browser requests the CSS file with those query parameters, the server transforms it to a JavaScript module
4. Your code receives a ready-to-use `CSSStyleSheet` object

This approach works around esbuild's current limitation with CSS import attributes while preserving the standard syntax in your source code.

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

**Via import attributes:**
```js
// This CSS file will be transformed regardless of include/exclude patterns
import styles from './my-styles.css' with { type: 'css' };
```

{{<tip "warning">}}
Without include patterns or import attributes, **no CSS files are transformed**. This prevents `<link rel="stylesheet">` tags from breaking unexpectedly.
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

See **[Configuration](/docs/reference/configuration/)** for all options.

## What's Next?

- **[Import Maps](../import-maps/)** - Use npm packages without bundling
- **[Configuration](/docs/reference/configuration/)** - Configuration reference
- **[Getting Started](../getting-started/)** - Set up your first demo
