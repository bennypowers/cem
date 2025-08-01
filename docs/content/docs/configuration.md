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
    # A Go regexp with named capture groups to extract information from demo file paths.
    urlPattern: "src/(?P<tag>[\w-]+)/demos/(?P<demo>[\w-]+).html"
    # A template to construct the canonical URL for a demo.
    # Uses `{groupName}` syntax to interpolate captures from `urlPattern`.
    urlTemplate: "https://example.com/elements/{tag}/{demo}/"

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

