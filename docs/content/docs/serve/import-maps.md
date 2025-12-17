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

The dev server provides multiple ways to customize import map entries, allowing you to override auto-generated mappings for specific packages.

### Overview

Import map entries are resolved with the following priority (highest priority wins):

1. **Auto-generated** - Automatically created from `package.json` dependencies
2. **Override file** - Custom JSON file with import map entries
3. **Config overrides** - Individual overrides specified in `.config/cem.yaml`

### Using an Override File

Create a JSON file containing custom import map entries. This file follows the standard [import map format](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script/type/importmap).

**.config/importmap.json:**
```json
{
  "imports": {
    "lit": "https://cdn.jsdelivr.net/npm/lit@3/+esm",
    "lit/": "https://cdn.jsdelivr.net/npm/lit@3/",
    "@patternfly/elements/": "/node_modules/@patternfly/elements/"
  },
  "scopes": {
    "/demos/": {
      "lit": "/node_modules/lit/index.js"
    }
  }
}
```

Then reference it in your config:

**.config/cem.yaml:**
```yaml
serve:
  importMap:
    overrideFile: '.config/importmap.json'
```

The entries in this file will be deep-merged with the auto-generated import map, with your entries taking precedence.

### Using Config Overrides

You can specify overrides directly in your config using the full import map structure:

**.config/cem.yaml:**
```yaml
serve:
  importMap:
    override:
      imports:
        'lit': 'https://cdn.jsdelivr.net/npm/lit@3/+esm'
        'lit/': 'https://cdn.jsdelivr.net/npm/lit@3/'
        '@custom/library': '/vendor/custom-library.js'
      scopes:
        '/demos/legacy/':
          'lit': 'https://cdn.jsdelivr.net/npm/lit@2/+esm'
```

Config overrides have the highest priority and will override both auto-generated and file-based entries. They support both `imports` and `scopes`.

### Common Use Cases

#### Use CDN for Specific Packages

Useful for testing against different versions or reducing local bundle size:

```yaml
serve:
  importMap:
    override:
      imports:
        'lit': 'https://cdn.jsdelivr.net/npm/lit@3/+esm'
        'lit/': 'https://cdn.jsdelivr.net/npm/lit@3/'
```

#### Point to Local Development Build

Override a workspace package to use a built version:

```yaml
serve:
  importMap:
    override:
      imports:
        '@my-org/components': '/dist/components.js'
```

#### Fix Broken Package Exports

Some packages have incorrect or missing `package.json` exports. Override them:

```yaml
serve:
  importMap:
    override:
      imports:
        'some-lib/utils': '/node_modules/some-lib/src/utils/index.js'
```

#### Use Scoped Overrides

For scenarios where different paths need different resolutions, you can use scopes in either the override file or config:

**.config/cem.yaml:**
```yaml
serve:
  importMap:
    override:
      imports:
        'react': 'https://esm.sh/react@18'
      scopes:
        '/demos/legacy/':
          'react': 'https://esm.sh/react@17'
```

Or via override file:

**.config/importmap.json:**
```json
{
  "imports": {
    "react": "https://esm.sh/react@18"
  },
  "scopes": {
    "/demos/legacy/": {
      "react": "https://esm.sh/react@17"
    }
  }
}
```

#### Disable Automatic Generation

If you want complete control via an override file:

```yaml
serve:
  importMap:
    generate: false
    overrideFile: '.config/importmap.json'
```

Use the `--no-import-map-generate` flag to disable temporarily:
```sh
cem serve --no-import-map-generate --import-map-override-file .config/importmap.json
```

### Debugging Overrides

To verify your overrides are applied correctly:

1. **View the generated import map** in your browser's DevTools:
   - Open any demo page
   - View source (Ctrl+U or Cmd+U)
   - Find the `<script type="importmap">` tag
   - Verify your overrides appear in the `imports` or `scopes` sections

2. **Check server logs** with the `--verbose` flag:
   ```sh
   cem serve --verbose
   ```

   Look for entries like:
   ```
   [DEBUG] Applied 2 import map overrides
   [DEBUG] Import map entry: lit -> https://cdn.jsdelivr.net/npm/lit@3/+esm
   ```

For complete configuration reference, see **[Configuration](/docs/configuration/)**.

## Benefits

- **No bundler needed** - Use npm packages directly in browser
- **Fast development** - No build step for dependencies
- **Standard syntax** - Same import statements as bundled apps
- **Workspace-aware** - Cross-package imports work automatically

## What's Next?

- **[Configuration](/docs/configuration/)** - Configuration reference
- **[Buildless Development](/docs/serve/buildless/)** - Write TypeScript and import CSS without build steps
- **[Getting Started](/docs/serve/getting-started/)** - Set up your first demo
