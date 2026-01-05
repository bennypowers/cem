---
title: Serve
weight: 20
---

{{< tip >}}
**TL;DR**: Run `cem serve` to start the dev server on port 8000 with live reload and TypeScript transformation. Use `--port 3000` to change port, `--rendering=chromeless` for testing, and `--css-transform 'src/**/*.css'` to import CSS as modules.
{{< /tip >}}

The `cem serve` command starts a development server specifically designed for custom element development.

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

See **[Buildless Development](/docs/usage/buildless-development/)** for details on CSS module imports.

### Disable automatic import map generation

```sh
# Skip import map generation from package.json
cem serve --no-import-map-generate
```

### Override import map entries

```sh
# Merge custom import map with auto-generated entries
cem serve --import-map-override-file custom-imports.json
```

Example `custom-imports.json`:
```json
{
  "imports": {
    "lit": "https://cdn.skypack.dev/lit@3.0.0",
    "@my-scope/": "/node_modules/@my-scope/"
  }
}
```

See **[Import Maps](/docs/usage/import-maps/)** for detailed documentation on configuring import maps.

### Ignore build directories in watcher

```sh
cem serve --watch-ignore 'dist/**,_site/**'
```

### Use chromeless mode for testing

```sh
# Perfect for Playwright/Puppeteer tests
cem serve --rendering=chromeless
```

### Test with Shadow DOM

```sh
# Verify Shadow DOM encapsulation
cem serve --rendering=shadow
```

## Demo Rendering Modes

The dev server supports multiple rendering modes for demo pages, affecting both the UI chrome and how demo content is rendered.

### Available Modes

| Mode | Description | When to Use |
| ---- | ----------- | ----------- |
| `light` | Full dev UI with demo in light DOM (default) | Standard development with all features enabled |
| `shadow` | Full dev UI with demo in shadow DOM | Testing encapsulation, isolated styling, shadow DOM behavior |
| `chromeless` | Minimal HTML with live reload, no UI chrome | Automated testing, embedding, clean screenshots, isolated development |
| `iframe` | *(Not yet implemented)* Full UI with demo in iframe | Future support for complete document isolation |

See **[Rendering Modes](/docs/usage/rendering-modes/)** for detailed documentation.

### Configuring Default Mode

Set the default rendering mode in your `.config/cem.yaml`:

```yaml
serve:
  demos:
    rendering: chromeless  # or "light" (default), "shadow"
```

If the `rendering` option is omitted or empty, demos default to `light` mode.

**Note:** Specifying `iframe` mode in the config will cause the server to fail at startup since it's not yet implemented.

### Per-Demo Override

Override the rendering mode for a specific demo using the `?rendering` query parameter:

```
http://localhost:8000/demos/my-element/basic.html?rendering=chromeless
```

Valid values: `light`, `shadow`, `chromeless`

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

## URL Rewrites

The dev server automatically resolves TypeScript source files from compiled output paths when using src/dist separation.

### Automatic Detection

If your project has a `tsconfig.json` with `rootDir` and `outDir`:

```json
{
  "compilerOptions": {
    "rootDir": "./src",
    "outDir": "./dist"
  }
}
```

The server automatically rewrites `/dist/` requests to `/src/` source files:
- Request: `/dist/components/button.js`
- URL pattern matches: `/dist/:path*` with `path=components/button.js`
- Rewrites to: `/src/components/button.ts`
- Transforms and serves on-demand

### Manual Configuration

Override or extend automatic rewrites in `.config/cem.yaml`:

```yaml
serve:
  urlRewrites:
    - urlPattern: "/dist/:path*"
      urlTemplate: "/src/{{.path}}"
    - urlPattern: "/lib/:path*"
      urlTemplate: "/sources/{{.path}}"
```

See **[Configuration > URL Rewrites](/docs/reference/configuration/#url-rewrites)** for detailed documentation including:
- tsconfig.json inheritance
- Workspace/monorepo support
- Edge cases and debugging

## See Also

- **[Getting Started](/docs/usage/getting-started/)** - Set up your first demo
- **[Buildless Development](/docs/usage/buildless-development/)** - Write TypeScript and import CSS without build steps
- **[Knobs](/docs/usage/knobs/)** - Interactive testing controls
- **[Import Maps](/docs/usage/import-maps/)** - Use npm packages without bundling
- **[Configuration](/docs/reference/configuration/)** - Configuration reference
