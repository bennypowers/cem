---
name: cem-serve
description: >
  Set up cem serve for custom element development. Use when the user asks to
  "start dev server", "set up cem serve", "preview elements", "run demos",
  or needs help configuring the dev server, import maps, or CSS transforms.
tools: Read, Bash, Glob, Grep
---

# Set Up CEM Serve for Development

Configure and run `cem serve` to preview custom elements with live reload,
auto-generated import maps, and interactive knobs.

## Best Practices

### Package Dependencies

Runtime imports must be in `dependencies`, not `devDependencies`. The import
map generator only resolves packages from `dependencies`.

```json
{
  "dependencies": {
    "lit": "^3.0.0"
  },
  "devDependencies": {
    "@pwrs/cem": "^0.9.0"
  }
}
```

### Project Config

Create `.config/cem.yaml` with generate and serve sections:

```yaml
generate:
  files:
    - elements/*/my-*.ts
  output: custom-elements.json
  designTokens:
    spec: tokens/my-tokens.json
    prefix: my-prefix
  demoDiscovery:
    fileGlob: elements/*/demo/*.html
    urlPattern: /elements/:tag/demo/:demo.html
    urlTemplate: /elements/{{.tag}}/demo/{{.demo}}/
serve:
  transforms:
    css:
      include:
        - "elements/**/*.css"
```

- `generate.files`: globs for element source files
- `generate.output`: manifest output path
- `generate.designTokens`: DTCG token file and CSS custom property prefix
- `generate.demoDiscovery`: how to find and route demo files
- `serve.transforms.css.include`: globs for CSS files to serve as JS modules (enables `import styles from './el.css' with { type: 'css' }`)

### Demo Files

Demos are **HTML partials**, not full documents. No `<!DOCTYPE>`, `<html>`, or
`<head>`. Include inline `<script type="module">` and `<style>` as needed.

The `index.html` demo is the default — keep it minimal (simplest possible usage).
Additional demos show specific features or variants.

Use `<meta itemprop="name">` and `<meta itemprop="description">` on non-index
demos for display in the dev server UI.

For demo authoring conventions, see the `write-docs` and `design-api` skills.

### Testing with Chromeless Mode

`cem serve --rendering=chromeless` strips the dev server UI, serving bare demo
HTML with import maps and transforms — ideal for browser-based testing with
Playwright or Puppeteer. Append `?rendering=chromeless` to any demo URL for
per-request chromeless mode without restarting.

See the `gen-tests` skill for full test scaffolding with page objects against
chromeless demos.

## Workflow

### 1. Generate the Manifest

```bash
cem generate
```

Verify demos appear without warnings. If you see "No URL configured for demo",
check that `demoDiscovery.urlPattern` uses URLPattern syntax (`:param` captures,
not regex).

### 2. Start the Server

```bash
cem serve
```

Default port is 8000. Use `--port <n>` to change.

### 3. Verify

Open `http://localhost:8000`. The UI should show element demos in the sidebar.
Check browser console for import errors — if a bare specifier fails, the
package needs to be in `dependencies`.

## Troubleshooting

### Bare specifier errors
Move the package from `devDependencies` to `dependencies` and reinstall.

### Duplicate demo route warnings
Delete the cached manifest and regenerate: `rm custom-elements.json && cem generate`

### Port in use
`lsof -ti:8000 | xargs kill` or use `--port`.

### Stale manifest on background regen
`cem serve` regenerates the manifest in the background on file changes. If the
background regen produces a smaller/empty manifest, ensure `generate.files` in
the config matches your source globs — the background regen uses the config,
not CLI args.
