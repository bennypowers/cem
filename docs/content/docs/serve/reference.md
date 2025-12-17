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

## Demo Rendering Modes

The dev server supports three rendering modes for demo pages, which affect how demo content is rendered in the browser.

### Available Modes

| Mode | Description | When to Use |
| ---- | ----------- | ----------- |
| `light` | Renders demo content in the light DOM (default) | Most demos, allows CSS from parent document to affect demo |
| `shadow` | Wraps demo in declarative shadow DOM (`<template shadowrootmode="open">`) | Testing encapsulation, isolated styling, shadow DOM behavior |
| `iframe` | *(Not yet implemented)* Renders demo in an isolated iframe | Future support for complete document isolation |

### Configuring Default Mode

Set the default rendering mode in your `.config/cem.yaml`:

```yaml
serve:
  demos:
    rendering: shadow  # or "light" (default)
```

If the `rendering` option is omitted or empty, demos default to `light` mode.

**Note:** Specifying `iframe` mode in the config will cause the server to fail at startup since it's not yet implemented.

### Per-Demo Override

Override the rendering mode for a specific demo using the `?rendering` query parameter:

```
http://localhost:8000/demos/my-element/basic.html?rendering=shadow
```

Valid values: `shadow`, `light`

- Invalid values are ignored and the config default is used
- If `iframe` is requested via query parameter, the server logs a warning, broadcasts an error overlay, and falls back to `shadow` mode

### Backward Compatibility

The legacy `?shadow=true` query parameter is still supported and will override to shadow mode:

```
http://localhost:8000/demos/my-element/basic.html?shadow=true
```

### Use Cases

**Light DOM (default):**
- General component testing
- Demos that rely on global styles
- Testing how components integrate with the parent page

**Shadow DOM:**
- Testing encapsulation behavior
- Verifying CSS custom properties penetrate shadow boundaries
- Testing `::part()` and `::slotted()` selectors
- Ensuring styles don't leak in or out

**Example Configuration:**

```yaml
# Default all demos to shadow mode
serve:
  demos:
    rendering: shadow
```

Then override specific demos back to light mode when needed:
```
/demos/integration-test.html?rendering=light
```

## See Also

- **[Getting Started](getting-started/)** - Set up your first demo
- **[Buildless Development](buildless/)** - Write TypeScript and import CSS without build steps
- **[Knobs](knobs/)** - Interactive testing controls
- **[Import Maps](import-maps/)** - Use npm packages without bundling
- **[Configuration](/docs/configuration/)** - Configuration reference
