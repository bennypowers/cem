## Command Structure

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

## Core Architecture

### HTTP Server (Go net/http)
Static file serving with smart MIME types and a middleware pipeline to handle requests

Out of scope for initial implementation:
- HTTPS support
- Proxy configuration
- Custom headers - CSP, Feature-Policy, etc.
- Request logging format - Machine-readable option (JSON logs)

See [stretch goals](./99-STRETCH-GOALS.md)

### Port Auto-Increment Behavior

- If port 8000 is taken, starts on 8001
- demos should not use localhost:8000 urls, so we can accept breakage as user error
- Log actual port prominently on startup
- Fail fast (exit 1) if entire range is exhausted

### Request Concurrency & Rate Limiting
It's a dev server, so we don't need yet to implement rate limiting

### Graceful Shutdown
SIGINT/SIGTERM handling strategy:
- Close WebSocket connections gracefully
- Wait for in-flight requests (with 5s timeout)
- Clean up file watchers
- Flush logs

### Middleware pipeline:
  1. Logging
     - passes through
  2. CORS headers
     - modifies response headers
  3. Import map injection
     - depends on package.json, which depends on manifest
     - modifies HTML response
  4. esbuild transform (TypeScript → JavaScript)
     - modifies responses for subresources
     - minimal dependencies since these are single-file transforms.
     - skip for node_modules files (excepting monorepo siblings linked there)
  5. CSS transform (optional: litcss-style)
     - modifies responses for subresources
     - minimal dependencies, just the config globs
     - short circuits when file doesn't match config globs
  6. Demo rendering and url transformation (with subresource loading)
     - depends on import map
     - short circuits when file ENOENT
  7. Static fallback

Middleware order is important, but there are middlewares which have minimal dependencies. The most important dependencies is package.json > import map > render HTML. When "dependencies" blocks in package.json files change, that affects the import map, which affects final rendered HTML. For the first draft, we can ignore changes to package.json files of node_modules dependencies (excluding monorepo siblings linked there)

### File Watching

#### Categories of files to watch

- Source files for elements (from manifest)
- Source files for demos (`generate/demodiscovery`)
- Config files

#### Considerations

- When source files are renamed or deleted, or entire watched directories, just regen manifest, then reload page
- Follow symlinks
- Debounce rapid changes (case: `git checkout`)
- keep a reasonable max. watched files limit, and warn in console

### Manifest Generation

- Share code with `generate/session.go`
- In-memory manifest (no file writes during serve)
- Watch mode triggers regeneration → broadcast reload event
- Demo discovery runs automatically on manifest update

#### Options

1. Read-lock for all HTTP handlers, write-lock during regeneration (blocks requests)
2. Serve stale manifest during regeneration, atomic swap when done (no blocking)
3. Version-based manifest with graceful degradation

It makes sense to block relevant changes on manifest generation. Perhaps a pre-regen step which determines if the file changes are relevant to the current page, and only blocks on regen of relevant manifest paths.

### Live Reload

- WebSocket endpoint at `/__cem-reload`
- Injected via `<script>` tag in HTML responses
- Debounced broadcasts (prevent spam)
- Events:
  - `reload`
  - `hmr`
  - `manifest-updated`
  - `error`

## Additional Features

### Error Handling
- Show compilation or template runtime errors in browser overlay
- Pretty error messages with stack traces relevant to user files
- Source-mapped errors point to original TypeScript
- ENOENT errors (missing demo) should show a 404 screen
- Distinguish use code errors (esbuidl compilation, template compilation) from runtime / server errors
  Place them in different blocks in the UI

### Cache Management
- ETag support for static assets
- `Cache-Control: no-cache` for demos (always fresh)
- In-memory cache for transformed files (cleared on change)
- Add max cache size (MB or entry count)
- Implement LRU eviction
- Add cache hit/miss metrics to logging

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

## Configuration File

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

##  Implementation Plan

### Phase 0: Test Infrastructure
- HTTP server test helpers (similar to `cmd/e2e_helpers_test.go`)
- Mock manifest generation
- Fixture-based demo HTML (per CLAUDE.md directive)
- WebSocket test client
- Transform cache test doubles
- set up benchmark / performance targets:
    - Manifest regeneration time limits (e.g., <500ms for incremental)
    - Max concurrent clients (e.g., 100)
    - Transform cache hit rate targets (e.g., >90%)
    - Memory usage limits (e.g., <500MB for typical project)

#### Acceptance Criteria
- Tests fail in red phase (command not implemented/stub implementation)
- Benchmarks return output (despite implementation being a stub)

### Phase 1: Core Server (cmd/serve.go, serve package)
- HTTP server with static file serving
- Watch mode integration (reuse generate/)
- [WebSocket live reload](./05-AUTORELOAD-HMR.md)
- Basic logging

#### Acceptance Criteria
- Server starts
- Core config tested
- Autoreload works
- Core Errors logged to console

### Phase 2: Import Maps (serve/importmap package)
- Auto-discovery from package.json
- Override file loading
- Merge strategy
- [HTML injection middleware](./20-IMPORTMAPS.md)

#### Acceptance Criteria
- Import map generated and appended to head

### Phase 3: Routing and Demo Rendering (serve/demo package)
- [URL routing](./10-URL-REWRITING.md)
- Template system (Go html/template)
- [Chrome templates](./15-DEMO-CHROME.md)

#### Acceptance Criteria
- Chrome implemented
- Server Errors logged to browser console

### Phase 4: Transforms ([serve/transform](./30-TRANSFORMS.md) package)
- esbuild integration (TypeScript)
- CSS transform
- Transform cache

#### Acceptance Criteria
- Tests pass
- "buildless workflow" enabled 
  1. User starts server with element.ts, element.css, and demo/index.html files
  2. Demo transforms sources to js
  3. Writing changes to any of the three files reloads the page

### Phase 5: Knobs ([serve/knobs](./40-KNOBS.md) package, embedded JS)
- Server-side control generation
- Client-side event handlers
- Template customization
- Type inference from CEM

#### Acceptance Criteria
- knobs display on demos
- knobs control multiple elements
- knobs have multiple kinds (enum, boolean)

### Phase 6: Polish
- Error overlay (serve/errors package)
- Manifest viewer UI (serve/manifest package) (stretch goal)
- Event logger (serve/events package)
- Documentation

#### Acceptance Criteria
- Server errors logged to overlay
- Custom element events logged
- Manifest browser implemented
