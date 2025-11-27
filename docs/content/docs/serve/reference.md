---
title: Command Reference
layout: docs
weight: 100
---

The `cem serve` command starts a development server specifically designed for custom element development.

```sh
cem serve [flags]
```

## Command Flags

| Flag | Description |
| ---- | ----------- |
| `--port` | Port to listen on (default: `8000`) |
| `--no-reload` | Disable live reload |
| `--target` | TypeScript/JavaScript transform target: `es2015`, `es2016`, `es2017`, `es2018`, `es2019`, `es2020`, `es2021`, `es2022`, `es2023`, `esnext` (default: `es2022`) |
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

## Usage Examples

### Start the dev server

```sh
cem serve
```

Starts the server on `http://localhost:8000` and opens your default browser to the element listing page.

### Use a different port

```sh
cem serve --port 3000
```

### Disable live reload

```sh
cem serve --no-reload
```

### Configure TypeScript target

```sh
cem serve --target es2020
```

### Enable CSS module transformation

```sh
# Transform component CSS to JavaScript modules
cem serve --css-transform 'src/**/*.css' --css-transform 'elements/**/*.css'
```

See **[Buildless Development](buildless/)** for details on CSS module imports.

### Ignore build directories in watcher

```sh
cem serve --watch-ignore 'dist/**,_site/**'
```

## See Also

- **[Getting Started](getting-started/)** - Set up your first demo
- **[Buildless Development](buildless/)** - Write TypeScript and import CSS without build steps
- **[Knobs](knobs/)** - Interactive testing controls
- **[Import Maps](import-maps/)** - Use npm packages without bundling
- **[Configuration](/docs/configuration/)** - Configuration reference
