---
title: Import Maps
layout: docs
weight: 60
---

The dev server automatically generates [import maps][importmaps] from your package.json, letting you use npm packages in demos without bundling. Import maps enable standard bare specifiers like `import { LitElement } from 'lit'` instead of relative paths to node_modules, supporting both single packages and [npm workspaces][npmworkspaces] monorepos. This complements [buildless TypeScript development][buildlessdevelopment] by eliminating the need for bundlers while preserving the same import syntax you'd use in production code.

Import maps work by mapping package names to file paths, so `import { LitElement } from 'lit'` resolves to `/node_modules/lit/index.js`. The dev server reads your package.json dependencies, resolves entry points from each package's exports or main field, and injects the generated import map into demo HTML automatically.

## Automatic Generation

The dev server reads your `package.json` and generates an import map for all dependencies:

**package.json:**
```json
{
  "dependencies": {
    "lit": "^3.0.0",
    "@patternfly/elements": "^3.0.0"
  }
}
```

**Generated import map:**
```json
{
  "imports": {
    "lit": "/node_modules/lit/index.js",
    "lit/": "/node_modules/lit/",
    "@patternfly/elements": "/node_modules/@patternfly/elements/pfe.min.js",
    "@patternfly/elements/": "/node_modules/@patternfly/elements/"
  }
}
```

## NPM Workspaces Support

For npm workspaces monorepos, the dev server discovers all workspace packages, generates import maps for each package's dependencies, then aggregates them into a single import map. Workspace packages become importable by name, so `import { MyButton } from '@my-org/button'` works across packages in your monorepo just like external dependencies:

```json
{
  "imports": {
    "lit": "/node_modules/lit/index.js",
    "@my-org/button": "/pkg/@my-org/button/src/index.js",
    "@my-org/card": "/pkg/@my-org/card/src/index.js"
  }
}
```

## Package Exports Support

The dev server respects `package.json` `exports` field:

```json
{
  "name": "my-package",
  "exports": {
    ".": "./dist/index.js",
    "./button": "./dist/button.js",
    "./card": "./dist/card.js"
  }
}
```

**Result:**
```json
{
  "imports": {
    "my-package": "/node_modules/my-package/dist/index.js",
    "my-package/button": "/node_modules/my-package/dist/button.js",
    "my-package/card": "/node_modules/my-package/dist/card.js"
  }
}
```

## Prefix Matching

The dev server includes prefix mappings (trailing `/`) for subpath imports:

```js
// Both work:
import { MyElement } from '@my-org/elements';
import { utils } from '@my-org/elements/utils.js';
```

## Debugging

### View the import map

Look in the HTML source of any demo page:

```html
<script type="importmap">
{
  "imports": {
    "lit": "/node_modules/lit/index.js",
    ...
  }
}
</script>
```

### Check server logs

With `--verbose` flag, see import map generation:

```sh
cem serve --verbose
```

Look for log entries like:
```
[INFO] Generated import map with 23 entries
[DEBUG] Import map entry: lit -> /node_modules/lit/index.js
```

## Custom Overrides

You can customize import map entries to point packages to CDNs, fix broken package exports, or use local development builds. The dev server resolves entries with this priority: auto-generated from package.json (lowest), override file entries (medium), and config overrides (highest priority wins).

Specify overrides directly in `.config/cem.yaml` using the full [import map format][importmaps]:

```yaml
serve:
  importMap:
    override:
      imports:
        'lit': 'https://cdn.jsdelivr.net/npm/lit@3/+esm'
        'lit/': 'https://cdn.jsdelivr.net/npm/lit@3/'
        '@my-org/components': '/dist/components.js'
      scopes:
        '/demos/legacy/':
          'lit': 'https://cdn.jsdelivr.net/npm/lit@2/+esm'
```

Alternatively, use an override file for complex mappings. Create `.config/importmap.json` with standard import map format, then reference it in your config:

```yaml
serve:
  importMap:
    overrideFile: '.config/importmap.json'
```

To disable automatic generation and use only your override file, set `generate: false` or use the `--no-import-map-generate` flag.

## What's Next?

- **[Buildless Development][buildlessdevelopment]** - TypeScript and CSS without build steps
- **[Configuration][configuration]** - Complete configuration reference
- **[Getting Started][gettingstarted]** - Set up your first demo

[importmaps]: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script/type/importmap
[npmworkspaces]: https://docs.npmjs.com/cli/using-npm/workspaces
[buildlessdevelopment]: ../buildless-development/
[configuration]: /docs/reference/configuration/
[gettingstarted]: ../getting-started/
