# Serve Architecture

The CEM dev server renders demo pages with Declarative Shadow DOM (DSD)
for instant first paint, then hydrates with Lit on the client.

## SSR Pipeline

1. **Build time** (`go generate`): esbuild bundles all LitElement chrome
   components + `@lit-labs/ssr` shims into `ssr-bundle.js`, then compiles
   to QuickJS bytecode (`ssr-bundle.qbc`) for faster init.

2. **Server startup**: `litssr.go` loads the bytecode into a wazero-backed
   QuickJS worker pool (via `bennypowers.dev/lit-ssr-wasm/go`).

3. **Per request**: The shadowroot middleware captures HTML responses,
   passes them through the lit-ssr renderer, and returns HTML with
   `<template shadowrootmode="open">` blocks injected.

4. **Client hydration**: The browser parses DSD immediately (styled first
   paint). `@lit-labs/ssr-client/lit-element-hydrate-support.js` patches
   LitElement to reuse existing shadow roots instead of re-rendering.

## Component Architecture

Chrome components are LitElement classes in `serve/elements/`:

```
serve/elements/
  pf-v6-badge/
    pf-v6-badge.ts     # LitElement source
    pf-v6-badge.css    # Styles (imported with { type: 'css' })
  cem-serve-chrome/
    cem-serve-chrome.ts # Top-level chrome wrapper
    cem-serve-chrome.css
  ...
```

At `go generate` time, `generate_elements.go`:
- Vendors lit into `templates/vendor/` for browser serving
- Transpiles each `.ts` to `.js` in `templates/elements/` (browser modules)
- Bundles all components into `templates/ssr-bundle.js` (SSR, with lit-ssr shims)
- Compiles to `templates/ssr-bundle.qbc` (pre-compiled bytecode)
- Generates custom-elements.json manifests via cem generate API

## Middleware Chain

Applied in reverse order (requests flow top-to-bottom):

1. **Routes** - `/__cem/*` endpoints (demos, manifest, WebSocket, vendor assets)
2. **TypeScript Transform** - Transpiles `.ts` on the fly
3. **CSS Transform** - Resolves CSS `@import`, serves CSS modules
4. **Import Map Injection** - Generates and injects import maps
5. **WebSocket Injection** - Adds live reload client script
6. **Shadow Root Injection** - Passes HTML through lit-ssr-wasm renderer
7. **Static Files** - Serves files from the watch directory

## Import Attributes

The dev server supports `import styles from './foo.css' with { type: 'css' }`:

1. Tree-sitter rewrites import attributes to query params before esbuild
2. CSS middleware detects `__cem-import-attrs` and serves CSS as JS module
3. Returns a `CSSStyleSheet` object for `adoptedStyleSheets`

## Key Files

| File | Role |
|------|------|
| `serve/litssr.go` | SSR init (bytecode or source fallback) |
| `serve/generate_elements.go` | Build pipeline (vendor, transpile, bundle, bytecode, manifests) |
| `serve/middleware/shadowroot/shadowroot.go` | HTML interceptor, passes through lit-ssr |
| `serve/middleware/routes/routes.go` | HTTP handler, `/__cem/*` endpoints |
| `serve/middleware/routes/templates.go` | Embed directives, template registry |
| `serve/elements/cem-serve-chrome/` | Top-level chrome component |
