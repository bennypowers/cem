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
    urlPattern: "/components/:element/demo/:demo.html"
    # A template to construct the canonical URL for a demo.
    # Uses {{.param}} syntax to interpolate URLPattern parameters.
    urlTemplate: "https://example.com/components/{{.element}}/demo/{{.demo}}/"

# Configuration for validation warnings.
warnings:
  disable:
    # Disable entire categories
    - "lifecycle"
    - "private"
    # Or disable specific warning rules
    - "lifecycle-lit-render"
    - "implementation-static-styles"
```

## Demo Discovery Features

The demo discovery system supports multiple ways to control how demos are associated with elements and how their URLs are generated.

### HTML5 Microdata

Demos can use HTML5 microdata to explicitly declare their URLs and associations:

```html
<!-- Explicit URL declaration -->
<meta itemprop="demo-url" content="/elements/call-to-action/demo/">
<meta itemprop="description" content="Primary variant demonstration">

<!-- Explicit element association -->
<meta itemprop="demo-for" content="rh-button pf-button">
```

### Association Priority

The system uses the following priority order to associate demos with elements:

1. **Explicit microdata**: `<meta itemprop="demo-for" content="element-name">`
2. **Magic comments**: `<!-- @tag element-name -->`
3. **Path-based**: Elements whose aliases appear in demo file paths
4. **Content-based**: Custom elements found in the demo HTML

### URL Generation Priority

URLs are generated using the following priority:

1. **Explicit microdata**: `<meta itemprop="demo-url" content="/path/to/demo/">`
2. **URLPattern fallback**: Using `urlPattern` and `urlTemplate` configuration
3. **No URL**: Demo is skipped if no pattern matches

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
  urlTemplate: "https://site.com/components/{{.element}}/demo/{{.demo}}/"
```

**Complex multi-site example:**
```yaml
demoDiscovery:
  fileGlob: src/components/**/demos/*.html  
  urlPattern: "/src/components/:component/demos/:variant.html"
  urlTemplate: "https://{{.component}}.examples.com/{{.variant}}/"
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

