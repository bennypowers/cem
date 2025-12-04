---
title: Import Maps
layout: docs
weight: 40
---

`cem serve` automatically generates [import maps](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script/type/importmap) from your `package.json`, enabling bare module specifiers in your demos.

## What Are Import Maps?

Import maps let you use package names in imports instead of full paths:

```js
// With import maps:
import { LitElement } from 'lit';

// Without import maps:
import { LitElement } from '../node_modules/lit/index.js';
```

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

## How It Works

### Single Package Repo

For a single package, the dev server:
1. Reads `package.json` dependencies
2. Resolves entry points from `node_modules/*/package.json`
3. Generates import map
4. Injects it into demo HTML

### NPM Workspaces Monorepo

For npm workspaces, the dev server:
1. Discovers all workspace packages
2. Generates import map for each package's dependencies
3. Aggregates into a single import map
4. Makes workspace packages importable by name

**Example workspace import map:**
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

You can customize the import map behavior in your config file (see **[Configuration](/docs/configuration/)**).

## Benefits

- **No bundler needed** - Use npm packages directly in browser
- **Fast development** - No build step for dependencies
- **Standard syntax** - Same import statements as bundled apps
- **Workspace-aware** - Cross-package imports work automatically

## What's Next?

- **[Configuration](/docs/configuration/)** - Configuration reference
- **[Buildless Development](/docs/serve/buildless/)** - Write TypeScript and import CSS without build steps
- **[Getting Started](/docs/serve/getting-started/)** - Set up your first demo
