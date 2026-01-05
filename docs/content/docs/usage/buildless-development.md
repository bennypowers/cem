---
title: Buildless Development
layout: docs
weight: 50
---

The `cem` dev server transforms TypeScript and CSS on-demand, letting you write modern code without running build tools. [Write TypeScript directly][writetypescriptdirectly] in your demos with full source map support, [import CSS as modules][importcssasmodules] using constructable stylesheets, and configure [automatic path rewrites][automaticpathrewrites] for projects with separate source and output directories. The server respects your `tsconfig.json` settings and provides [browser compatibility control][browsercompatibilitycontrol] through esbuild targets.

This buildless approach eliminates `tsc --watch`, CSS preprocessors, and bundlers from your development workflow. You write TypeScript and CSS, import them directly in demos, and the dev server handles transformation transparently with full debugging support.

The dev server handles TypeScript files (`.ts`), CSS files as modules (with opt-in patterns), ES modules from npm via import maps, and modern JavaScript syntax. JavaScript files (`.js`, `.mjs`), HTML files, and JSON files work as-is without building them ahead-of-time.

## Write TypeScript Directly

Import TypeScript files directly in your demos without compilation:

```html
<!-- In your demo -->
<script type="module" src="../src/my-element.ts"></script>
```

The dev server handles TypeScript transformation on-demand with full source map support for debugging. You don't need to run `tsc --watch`, set up build pipelines, wait for compilation before testing, or manage separate source and output directories.

### Browser Compatibility

Control which browsers your code supports using the `--target` flag:

```sh
cem serve --target es2020
```

The default target is `es2022`. See [esbuild's target documentation][esbuildstargetdocumentation] for all available targets.

### Your tsconfig.json Works

The dev server respects your existing TypeScript configuration including `compilerOptions.target`, `compilerOptions.module`, `compilerOptions.rootDir` and `compilerOptions.outDir` for src/dist separation, and `include` and `exclude` patterns. The command-line `--target` flag overrides `tsconfig.json` settings if provided.

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

With this configuration, a request to `/dist/components/button.js` resolves to `/src/components/button.ts`, which the server transforms and serves as TypeScript source.

This lets you write demos using the same paths your production build uses without changing import paths between development and production. It works seamlessly with TypeScript's project references and supports monorepo/workspace setups with per-package `tsconfig.json` files.

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

URL patterns use [URLPattern][urlpattern] syntax:
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

See **[Configuration][configuration]** for details.

## Import CSS as Modules

The dev server supports two ways to import CSS files as constructable stylesheets. Import attributes with `with { type: 'css' }` work automatically without configuration, while plain imports require glob pattern configuration.

### Import Attributes

Use the standard [import attributes][importattributes] syntax for automatic CSS transformation:

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

This modern, standards-based syntax always transforms CSS to modules—no configuration needed. The dev server rewrites `with { type: 'css' }` to query parameters during TypeScript transformation, then transforms the CSS file to a JavaScript module that exports a `CSSStyleSheet` object.

### Plain Imports (Requires Configuration)

Alternatively, use plain imports without the `with` keyword:

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

Plain imports only transform CSS files matching glob patterns that you configure via command line or config file:

**Via command line:**
```sh
cem serve --css-transform 'src/**/*.css' --css-transform 'elements/**/*.css'
```

**Via config file:**
```yaml
serve:
  transforms:
    css:
      include:
        - src/**/*.css
        - elements/**/*.css
      exclude:
        - demo/**/*.css
```

{{<tip "warning">}}
Without glob patterns, plain CSS imports won't transform—only imports with `with { type: 'css' }` work. This prevents `<link rel="stylesheet">` tags from breaking unexpectedly.
{{</tip>}}

## Debugging

Source maps work automatically, so stack traces point to your original TypeScript, browser DevTools show your source files, and breakpoints work in TypeScript rather than generated JavaScript.

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

See **[Configuration][configuration]** for all options.

## What's Next?

- **[Import Maps][importmaps]** - Use npm packages without bundling
- **[Configuration][configuration]** - Configuration reference
- **[Getting Started][gettingstarted]** - Set up your first demo

[importattributes]: https://github.com/tc39/proposal-import-attributes
[configuration]: /docs/reference/configuration/
[urlpattern]: https://developer.mozilla.org/en-US/docs/Web/API/URL_Pattern_API
[esbuildstargetdocumentation]: https://esbuild.github.io/api/#target
[writetypescriptdirectly]: #write-typescript-directly
[importcssasmodules]: #import-css-as-modules
[automaticpathrewrites]: #separate-dist-directory
[browsercompatibilitycontrol]: #browser-compatibility
[importmaps]: ../import-maps/
[gettingstarted]: ../getting-started/
