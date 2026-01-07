---
title: Serve
weight: 20
---

The `cem serve` command starts a development server for custom element development with live reload, TypeScript transformation, and automatic import map generation.

```sh
cem serve [flags]
```

## Command Flags

| Flag | Description |
| ---- | ----------- |
| `--port` | Port to listen on (default: `8000`) |
| `--rendering` | Demo rendering mode: `light` (full UI), `shadow` (Shadow DOM), or `chromeless` (minimal, no UI) (default: `light`) |
| `--no-reload` | Disable live reload |
| `--target` | TypeScript/JavaScript transform target: `es2015`, `es2016`, `es2017`, `es2018`, `es2019`, `es2020`, `es2021`, `es2022`, `es2023`, `esnext` (default: `es2022`) |
| `--no-import-map-generate` | Disable automatic import map generation from package.json |
| `--import-map-override-file` | Path to JSON file with custom import map entries (merged with auto-generated map) |
| `--css-transform` | Glob patterns for CSS files to transform to JavaScript modules (opt-in, e.g., `src/**/*.css,elements/**/*.css`) |
| `--css-transform-exclude` | Glob patterns for CSS files to exclude from transformation (e.g., `demo/**/*.css`) |
| `--watch-ignore` | Glob patterns to ignore in file watcher (comma-separated, e.g., `_site/**,dist/**`) |

### Global Flags

| Flag | Description |
| ---- | ----------- |
| `--config` | Path to config file (default: `.config/cem.yaml`) |
| `--package`, `-p` | Deno-style package specifier (e.g., `npm:@scope/package`) or path to package directory |
| `--source-control-root-url` | Canonical public source control URL for primary branch (e.g., `https://github.com/user/repo/tree/main/`) |
| `--quiet`, `-q` | Quiet output (only warnings and errors) |
| `--verbose`, `-v` | Verbose logging output |

## Examples

### Basic usage

```sh
cem serve
```

Starts the server on `http://localhost:8000`.

### Common options

```sh
# Use a different port
cem serve --port 3000

# Disable live reload
cem serve --no-reload

# Use chromeless mode for testing
cem serve --rendering=chromeless

# Configure TypeScript target
cem serve --target es2020

# Ignore build directories in watcher
cem serve --watch-ignore 'dist/**,_site/**'
```

## Configuration

All command-line flags have corresponding configuration file options. See **[Configuration](/docs/reference/configuration/)** for the complete reference.

Example `.config/cem.yaml`:

```yaml
serve:
  port: 3000
  demos:
    rendering: shadow
  target: es2020
  transforms:
    css:
      include:
        - 'src/**/*.css'
      exclude:
        - 'demo/**/*.css'
  importMap:
    generate: true
    overrideFile: '.config/importmap.json'
  urlRewrites:
    - urlPattern: "/dist/:path*"
      urlTemplate: "/src/{{.path}}"
  watch:
    ignore:
      - 'dist/**'
      - '_site/**'
```

## See Also

- **[Development Workflow](/docs/usage/workflow/)** - Using the dev server in your workflow
- **[Buildless Development](/docs/usage/buildless-development/)** - TypeScript and CSS transformation
- **[Rendering Modes](/docs/usage/rendering-modes/)** - Light, shadow, and chromeless modes
- **[Import Maps](/docs/usage/import-maps/)** - Using npm packages without bundling
- **[Configuration](/docs/reference/configuration/)** - Complete configuration reference
