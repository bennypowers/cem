---
title: Configuration
layout: docs
weight: 20
---

You can configure `cem` via a `cem.yaml` file in your project's `.config` directory, or by using command-line flags.

## Configuration File

Here is a complete example of a `.config/cem.yaml` file with all available options explained.

```yaml
# The canonical public source control URL for your repository root.
# Used for generating source links in the manifest.
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"

# Configuration for the `generate` command.
generate:
  # A list of glob patterns for files to include in the analysis.
  files:
    - "src/**/*.ts"

  # A list of glob patterns for files to exclude from the analysis.
  exclude:
    - "src/**/*.test.ts"

  # The path to write the final custom-elements.json manifest.
  # If omitted, the manifest is written to standard output.
  output: "custom-elements.json"

  # By default, certain files like TypeScript declaration files (`.d.ts`) are excluded.
  # Set to `true` to include all files matched by the `files` glob.
  noDefaultExcludes: false

  # Configuration for integrating Design Tokens.
  designTokens:
    # An npm specifier or local path to a DTCG-formatted JSON module.
    spec: "npm:@my-ds/tokens/tokens.json"
    # A CSS custom property prefix to apply to the design tokens.
    prefix: "--my-ds"

  # Configuration for discovering element demos.
  demoDiscovery:
    # A glob pattern to find demo files.
    fileGlob: "src/**/demos/*.html"
    # URLPattern for extracting parameters from demo file paths.
    # Uses standard URLPattern syntax with named parameters.
    urlPattern: "/src/:component/demos/:demo.html"
    # Go template to construct the canonical URL for a demo.
    # Uses {{.param}} syntax with optional template functions.
    # Available functions: alias, slug, lower, upper
    urlTemplate: "https://example.com/components/{{.component | alias}}/demo/{{.demo | slug}}/"

# Configuration for validation warnings.
warnings:
  disable:
    # Disable entire categories
    - "lifecycle"
    - "private"
    # Or disable specific warning rules
    - "lifecycle-lit-render"
    - "implementation-static-styles"

# Configuration for the `serve` command.
serve:
  # Port to listen on
  port: 8000

  # Disable live reload
  no-reload: false

  # Glob patterns to ignore in file watcher
  watchIgnore:
    - 'dist/**'
    - '_site/**'
    - 'node_modules/**'

  # Demo rendering configuration
  demos:
    # Default rendering mode for demos: "light" (default), "shadow", or "iframe"
    # Can be overridden per-demo with ?rendering=shadow|light query parameter
    # Note: "iframe" mode is not yet implemented
    rendering: light

  # Path mappings for src/dist separation
  # Maps URL prefixes to source directories for TypeScript resolution
  # Automatically detected from tsconfig.json (rootDir/outDir)
  # Manual mappings override automatic detection
  pathMappings:
    "/dist/": "/src/"
    "/lib/": "/sources/"

  # Import map configuration
  importMap:
    # Enable automatic import map generation (default: true)
    generate: true

    # Path to a JSON file containing custom import map entries
    # These entries are merged with auto-generated imports
    overrideFile: '.config/importmap.json'

    # Import map override from config (highest priority)
    # Full import map structure with imports and scopes
    override:
      imports:
        'lit': 'https://cdn.jsdelivr.net/npm/lit@3/+esm'
        'lit/': 'https://cdn.jsdelivr.net/npm/lit@3/'
      scopes:
        '/demos/':
          'lit': '/node_modules/lit/index.js'

  # Transform configuration
  transforms:
    # TypeScript transformation
    typescript:
      enabled: true
      target: es2022  # es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, esnext

    # CSS transformation (opt-in)
    css:
      enabled: true
      # Glob patterns for CSS files to transform to JavaScript modules
      include:
        - 'src/**/*.css'
        - 'elements/**/*.css'
      # Glob patterns to exclude from transformation
      exclude:
        - 'demo/**/*.css'
        - '**/*.min.css'
```

## Import Map Overrides

The dev server automatically generates import maps from `package.json`, but you can customize or override these mappings.

### Override Priority

Import map entries are merged with the following priority (highest wins):

1. **Auto-generated** - From `package.json` dependencies
2. **Override file** - Custom import map JSON file (via `serve.importMap.overrideFile`)
3. **Config overrides** - Individual overrides in config (via `serve.importMap.override`)

### Override File

Create a JSON file with custom import map entries:

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

```yaml
serve:
  importMap:
    overrideFile: '.config/importmap.json'
```

### Config Overrides

Specify overrides directly in your config using the full import map structure:

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

Config overrides support both `imports` and `scopes`, giving you full control over the import map structure.

### Use Cases

**Use CDN instead of local modules:**
```yaml
serve:
  importMap:
    override:
      imports:
        'lit': 'https://cdn.jsdelivr.net/npm/lit@3/+esm'
        'lit/': 'https://cdn.jsdelivr.net/npm/lit@3/'
```

**Point to local development version:**
```yaml
serve:
  importMap:
    override:
      imports:
        '@my-org/components': '/packages/components/src/index.js'
```

**Override specific subpaths:**
```yaml
serve:
  importMap:
    override:
      imports:
        'some-lib/broken-export': '/local-fixes/fixed-export.js'
```

**Use scoped overrides for different contexts:**
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

**Disable automatic generation (use only override file):**
```yaml
serve:
  importMap:
    generate: false
    overrideFile: '.config/importmap.json'
```

## Path Mappings

Path mappings enable buildless TypeScript development with src/dist separation. The dev server automatically resolves requests to compiled output paths (e.g., `/dist/foo.js`) to their TypeScript source files (e.g., `/src/foo.ts`).

### Automatic Detection from tsconfig.json

The dev server automatically reads your `tsconfig.json` and creates path mappings based on `rootDir` and `outDir`:

**tsconfig.json:**
```json
{
  "compilerOptions": {
    "rootDir": "./src",
    "outDir": "./dist"
  }
}
```

This automatically creates the mapping: `"/dist/" → "/src/"`

**How it works:**

1. Request arrives: `GET /dist/components/button.js`
2. Dev server checks path mappings: `/dist/` maps to `/src/`
3. Resolves to source file: `/src/components/button.ts`
4. Transforms TypeScript to JavaScript on-demand
5. Serves transformed code with source maps

### Manual Configuration

Override or extend automatic mappings in your config:

```yaml
serve:
  pathMappings:
    "/dist/": "/src/"
    "/lib/": "/sources/"
    "/build/": "/typescript/"
```

Manual mappings take precedence over automatic detection from `tsconfig.json`.

### Resolution Strategy

The dev server uses a multi-strategy approach to find source files:

1. **Co-located files** (backward compatibility)
   - Tries `/dist/button.ts` alongside `/dist/button.js`
   - Supports in-place compilation workflows

2. **Path mappings** (src/dist separation)
   - Applies configured mappings: `/dist/` → `/src/`
   - Resolves `/dist/button.js` to `/src/button.ts`

If neither strategy finds a source file, the request falls through to static file serving.

### tsconfig.json Inheritance

Path mappings work with TypeScript's `extends` feature:

**tsconfig.base.json:**
```json
{
  "compilerOptions": {
    "rootDir": "./src"
  }
}
```

**tsconfig.json:**
```json
{
  "extends": "./tsconfig.base.json",
  "compilerOptions": {
    "outDir": "./dist"
  }
}
```

The dev server correctly merges inherited values and creates the mapping `"/dist/" → "/src/"`.

### Workspace Mode

In workspace/monorepo setups, each package can have its own `tsconfig.json` with different path mappings:

```
packages/
  components/
    tsconfig.json  # rootDir: "src", outDir: "dist"
    src/
    dist/
  elements/
    tsconfig.json  # rootDir: "lib", outDir: "build"
    lib/
    build/
```

Each package's mappings are discovered independently and applied to requests within that package's path.

### Edge Cases

**Default rootDir with custom outDir:**

```json
{
  "compilerOptions": {
    "rootDir": ".",
    "outDir": "dist"
  }
}
```

Creates mapping: `"/dist/" → "/./"`

The dev server normalizes this correctly via `filepath.Join()`, resolving `/dist/foo.js` to `/foo.ts` at the project root.

**Same rootDir and outDir:**

```json
{
  "compilerOptions": {
    "rootDir": ".",
    "outDir": "."
  }
}
```

No path mapping is created (in-place compilation). The dev server falls back to co-located file resolution.

### Use Cases

**Standard src/dist separation:**
```yaml
# Automatic from tsconfig.json
# No config needed - just works!
```

**Custom build output directories:**
```yaml
serve:
  pathMappings:
    "/lib/": "/src/"
    "/esm/": "/src/"
    "/cjs/": "/src/"
```

**Legacy projects with non-standard structure:**
```yaml
serve:
  pathMappings:
    "/compiled/": "/source/typescript/"
    "/js/": "/ts/"
```

### Debugging Path Mappings

Use verbose logging to see path resolution in action:

```sh
cem serve --verbose
```

Look for log messages like:
```
PathResolver: mapped source found: /dist/button.js -> /src/button.ts (mapping: /dist/ -> /src/)
```

## Demo Discovery Features

The demo discovery system supports multiple ways to control how demos are associated with elements and how their URLs are generated.

### HTML5 Microdata

Demos can use HTML5 microdata to explicitly declare their URLs and associations:

```html
<!-- Explicit URL declaration -->
<meta itemprop="demo-url" content="/elements/call-to-action/demo/">
<meta itemprop="description" content="Primary variant demonstration">

<!-- Rich markdown description -->
<script type="text/markdown" itemprop="description">
# Call to Action Demo
Showcases primary variant with styling, accessibility, and interaction states.
</script>

<!-- Explicit element association -->
<meta itemprop="demo-for" content="rh-button pf-button">
```

### Association Priority

The system uses the following priority order to associate demos with elements:

1. **Explicit microdata**: `<meta itemprop="demo-for" content="element-name">`
2. **Path-based**: Elements whose aliases appear in demo file paths
3. **Content-based**: Custom elements found in the demo HTML

#### Path-Based Association with URLPattern

When a `urlPattern` is configured, path-based association becomes **parameter-position aware**. Element aliases only match when they appear in path positions that correspond to URLPattern parameters.

**Example Configuration:**
```yaml
demoDiscovery:
  fileGlob: "shop/**/demos/*.html"
  urlPattern: "/shop/:element/:demo.html"
  urlTemplate: "https://mysite.com/shop/{{.element | alias}}/{{.demo}}/"
```

**Element Aliases:**
```yaml
aliases:
  my-shop-button: "shop-button"
  my-accordion: "accordion"
  my-accordion-header: "accordion-header"
```

**Path Matching Behavior:**

- ✅ `/shop/shop-button/basic.html` → matches `my-shop-button` (alias "shop-button" in `:element` position)
- ❌ `/shop/my-shop-button/basic.html` → no match (alias "shop-button" ≠ "my-shop-button")
- ✅ `/shop/accordion-header/demo.html` → matches `my-accordion-header` only
- ❌ `/shop/accordion-header/demo.html` → does NOT match `my-accordion` (prevents false positives)

This precision matching prevents issues where shorter aliases (like "accordion") incorrectly match longer path segments (like "my-accordion-header").

**Legacy Behavior (no URLPattern):**
When `urlPattern` is not configured, the system falls back to compatibility mode with improved substring matching that prioritizes longer, more specific aliases over shorter ones.

### Description Sources

Demo descriptions are extracted exclusively from microdata:

- **Meta tags**: `<meta itemprop="description" content="Simple description">`
- **Script tags**: `<script type="text/markdown" itemprop="description">Rich **markdown** content</script>`

### URL Generation Priority

URLs are generated using the following priority:

1. **Explicit microdata**: `<meta itemprop="demo-url" content="/path/to/demo/">`
2. **URLPattern fallback**: Using `urlPattern` and `urlTemplate` configuration
3. **No URL**: Demo is skipped if no pattern matches

### URL Template Functions

The `urlTemplate` uses Go template syntax with a set of built-in functions for transforming URLPattern parameters:

| Function | Description | Example |
|----------|-------------|---------|
| `alias` | Apply element alias mapping from configuration | `{{.tag \| alias}}` |
| `slug` | Convert to URL-friendly slug format | `{{.demo \| slug}}` |
| `lower` | Convert to lowercase | `{{.component \| lower}}` |
| `upper` | Convert to uppercase | `{{.section \| upper}}` |

**Template Examples:**

```yaml
# Basic parameter interpolation
urlTemplate: "https://example.com/{{.component}}/{{.demo}}/"

# Apply alias transformation
urlTemplate: "https://example.com/{{.component | alias}}/{{.demo}}/"

# Chain multiple functions
urlTemplate: "https://example.com/{{.component | alias | slug}}/{{.demo | lower}}/"

# Function call syntax (alternative)
urlTemplate: "https://example.com/{{alias .component}}/{{slug .demo}}/"
```

**Important:** All transformations must be explicitly specified. Unlike previous versions, no automatic aliasing or slugification is applied unless explicitly requested in the template.

### Configuration Examples

**Minimal (microdata-driven):**
```yaml
demoDiscovery:
  fileGlob: elements/**/demo/*.html
```

**URLPattern with explicit configuration:**
```yaml
demoDiscovery:
  fileGlob: elements/**/demo/*.html
  urlPattern: "/elements/:element/demo/:demo.html"
  urlTemplate: "https://site.com/components/{{.element | alias}}/demo/{{.demo | slug}}/"
```

**Complex multi-site example:**
```yaml
demoDiscovery:
  fileGlob: src/components/**/demos/*.html  
  urlPattern: "/src/components/:component/demos/:variant.html"
  urlTemplate: "https://{{.component | alias | lower}}.examples.com/{{.variant | slug}}/"
```

## Global Flags

These flags can be used with any `cem` command.

| Flag            | Description                                                           |
| --------------- | --------------------------------------------------------------------- |
| `--config`      | Path to a custom config file.                                         |
| `--package`     | deno-style package specifier, or path to the local package directory. |
| `--verbose`, -v | Enable verbose logging output.                                        |
| `--help`, -h    | Show help for a command.                                              |

## Command-Line Flags

All configuration options can also be set via command-line flags. Flags will always override any values set in the configuration file.

For example, to override the `output` and `exclude` options for the `generate` command:

```sh
cem generate --output my-manifest.json --exclude "src/legacy/**"
```

### Validate Command Flags

The `validate` command supports additional flags:

| Flag                      | Description                                                    |
| ------------------------- | -------------------------------------------------------------- |
| `--disable`               | Disable specific warning rules or categories (repeatable).     |
| `--format`                | Output format: `text` (default) or `json`.                    |

For example, to validate with disabled warnings and JSON output:

```sh
cem validate --disable lifecycle --disable private --format json
```

