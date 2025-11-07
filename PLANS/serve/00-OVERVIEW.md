## 1. Command Structure

```bash
cem serve [directory] [flags]
```

**Behavior:**
- Top-level command (like `cem lsp`)
- Reuses generate infrastructure for in-memory manifest (no temp files)
- Watch mode triggers incremental manifest regeneration
- Server stays running, manifest updates propagate via WebSocket

**Flags:**
```
--port, -p         Port to listen on (default: 8000)
--open, -o         Open browser on startup
--config           Path to cem.config.json
--target           Esbuild transform target for TypeScript files (default 2020? 2024?)
--import-map       Path to custom import map overrides
--transform-css    List of glob patterns of CSS files to transform to JavaScript modules
--no-reload        Disable live reload
--template-dir     Custom template directory for demo chrome
```

---

## 2. Core Architecture

### HTTP Server (Go net/http)
- Static file serving with smart MIME types
- Middleware pipeline:
  1. Logging
  2. CORS headers
  3. Import map injection
  4. esbuild transform (TypeScript → JavaScript)
  5. CSS transform (optional: litcss-style)
  6. Demo rendering and url transformation (with subresource loading)
  7. Static fallback

### Manifest Integration
- Share code with `generate/session.go`
- In-memory manifest (no file writes during serve)
- Watch mode triggers regeneration → broadcast reload event
- Demo discovery runs automatically on manifest update

### Live Reload
- WebSocket endpoint at `/__cem-reload`
- Injected via `<script>` tag in HTML responses
- Events:
  - `reload`
  - `hmr`
  - `manifest-updated`
  - `error`
- Debounced broadcasts (prevent spam)

---

## 7. Additional Features

### Error Handling
- Show compilation errors in browser overlay
- Pretty error messages with stack traces relevant to user files
- Source-mapped errors point to original TypeScript

### Cache Management
- ETag support for static assets
- `Cache-Control: no-cache` for demos (always fresh)
- In-memory cache for transformed files (cleared on change)

### CORS & Headers
```
Access-Control-Allow-Origin: *
X-Content-Type-Options: nosniff
```

### Logging
```
[cem serve] Starting dev server on http://localhost:8000
[cem serve] Manifest generated (23 elements, 47 demos)
[cem serve] Watching 128 files for changes
[cem serve] GET /components/pf-button/demo/variants/ 200 12ms
[cem serve] File changed: elements/pf-button/pf-button.ts
[cem serve] Manifest regenerated (1 file changed)
[cem serve] Broadcast reload event to 3 clients
```

---

## 8. Configuration File

**`.config/cem.yaml` additions:** (defaults)
```yaml
serve:
  port: 8000 # auto increments when port in use
  open: true # launches browser
  importMap: # see importMap section
  knobs: # see knobs section
  transforms: # see transform section
  reload: # see reload section
```

---

## 9. Implementation Plan

1. **Phase 1: Core Server** (cmd/serve.go, server package)
   - HTTP server with static file serving
   - Watch mode integration (reuse generate/)
   - WebSocket live reload
   - Basic logging

2. **Phase 2: Routing and Demo Rendering** (server/demo package)
   - URL routing
   - Template system (Go html/template)
   - Default chrome templates
   - Raw demo mode

3. **Phase 3: Import Maps** (server/importmap package)
   - Auto-discovery from package.json
   - Override file loading
   - Merge strategy
   - HTML injection middleware

4. **Phase 4: Transforms** (server/transform package)
   - esbuild integration (TypeScript)
   - CSS transform
   - Transform cache

5. **Phase 5: Knobs** (server/knobs package, embedded JS)
   - Server-side control generation
   - Client-side event handlers
   - Template customization
   - Type inference from CEM

6. **Phase 6: Polish**
   - Error overlay
   - Manifest viewer UI
   - Event logger
   - Documentation

