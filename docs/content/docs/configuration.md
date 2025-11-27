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

