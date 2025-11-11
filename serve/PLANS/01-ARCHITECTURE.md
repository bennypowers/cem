# Development Server Architecture

This document describes cross-cutting architectural concerns that span multiple implementation phases.

## Middleware Pipeline

The HTTP server uses a middleware pipeline to process requests. Order is critical.

### Pipeline Stages

1. **Logging**
   - Type: Pass-through
   - Function: Log all requests/responses
   - Dependencies: None

2. **CORS Headers**
   - Type: Response modifier
   - Function: Add CORS headers for cross-origin development
   - Headers:
     ```
     Access-Control-Allow-Origin: *
     X-Content-Type-Options: nosniff
     ```

3. **Import Map Injection** (Phase 2)
   - Type: HTML response modifier
   - Function: Inject `<script type="importmap">` into HTML `<head>`
   - Dependencies: package.json → manifest
   - Short-circuits: Non-HTML responses

4. **esbuild Transform** (Phase 4)
   - Type: Subresource modifier (TypeScript → JavaScript)
   - Function: Transform `.ts` files on-the-fly
   - Dependencies: Minimal (single-file transforms)
   - Short-circuits: node_modules files (except monorepo siblings)

5. **CSS Transform** (Phase 4)
   - Type: Subresource modifier (CSS → JS module)
   - Function: Wrap CSS in CSSStyleSheet export
   - Dependencies: Config globs only
   - Short-circuits: Files not matching config patterns

6. **Demo Rendering** (Phase 3)
   - Type: HTML response generator
   - Function: Render demo HTML with chrome template
   - Dependencies: Import map (from stage 3)
   - Short-circuits: File not found (ENOENT)

7. **Static Fallback**
   - Type: File server
   - Function: Serve static files directly
   - Dependencies: None

### Dependency Chain

The critical dependency path is:

```
package.json changes
  → regenerate import map
    → affects rendered HTML
```

**First draft simplification**: Ignore changes to `package.json` files of node_modules dependencies (excluding monorepo siblings linked there).

### Extension Points

Each phase can register middleware at appropriate pipeline stages:
- **Phase 1**: Logging, CORS
- **Phase 2**: Import map injection
- **Phase 3**: Demo rendering
- **Phase 4**: Transform middlewares

### Implementation Status (as of Phase 4 completion)

**Fully Implemented**:
- ✅ Logging middleware with colored badges (INFO, WARN, ERROR, DEBUG)
- ✅ CORS headers with security headers
- ✅ Import map injection with workspace/monorepo support
- ✅ TypeScript transformation with tree-sitter dependency tracking
- ✅ CSS transformation to constructable stylesheets
- ✅ Demo rendering with chrome templates and navigation drawer
- ✅ Static file serving with proper MIME types (.mjs, .cjs support)

**Additional Features Implemented**:
- ✅ Error overlay with source-mapped stack traces
- ✅ 404 error page with helpful navigation
- ✅ WebSocket live reload with smart invalidation
- ✅ Transform cache with LRU eviction and worker pool
- ✅ Path traversal protection
- ✅ View Transitions API support

**Monorepo/Workspace Support** (implemented beyond plan):
- Multi-package discovery and routing
- Workspace-scoped import maps
- Cross-package dependency tracking
- Package conflict detection

## Manifest Integration

The manifest is central to multiple features. It must be managed carefully to avoid race conditions.

### Manifest Lifecycle

1. **Initial Generation**
   - On server startup, generate manifest in-memory
   - No file writes during serve mode
   - Share code with `generate/session.go`

2. **Watch Mode Triggers**
   - File changes trigger manifest regeneration
   - Demo discovery runs automatically on manifest update
   - Broadcast reload event after regeneration

3. **Concurrency Strategy**

Three options for handling manifest access during regeneration:

1. **Read-lock for handlers, write-lock for regeneration** (blocks requests)
   - Simple to implement
   - May cause request delays during regeneration

2. **Serve stale, swap atomically** (no blocking)
   - Better UX, requests never block
   - Requires atomic pointer swap

3. **Version-based with graceful degradation**
   - Most complex
   - Allows advanced features

**Recommended approach**: Option 2 (serve stale) with pre-regeneration relevance check.

**Optimization**: Before regenerating, determine if file changes are relevant to current pages. Only block regeneration for relevant manifest paths.

### Manifest Consumers

Phases that depend on manifest data:

- **Phase 2 (Import Maps)**: Reads dependencies, workspace structure
- **Phase 3 (Demo Rendering)**: Reads demo metadata, element info
- **Phase 5 (Knobs)**: Reads element attributes, properties, CSS custom properties

## File Watching

Watch source files and trigger appropriate updates.

### Categories of Watched Files

1. **Source files for elements** (from manifest)
   - Triggers manifest regeneration

2. **Source files for demos** (`generate/demodiscovery`)
   - Triggers manifest regeneration

3. **Config files** (`cem.config.yaml`, `package.json`)
   - Triggers full server reconfiguration

### Watch Behavior

- **Renames/deletes**: Regenerate manifest, reload page
- **Symlinks**: Follow symlinks to actual files
- **Debouncing**: Batch rapid changes (e.g., `git checkout`)
  - Recommended: 150ms debounce window
- **Limits**: Warn if exceeding reasonable max watched files

### Implementation Notes

- Use existing file watching infrastructure from `generate/session_watch.go`
- Coordinate with manifest regeneration strategy
- Clear transform cache for changed files (and their dependency tree)

## Error Handling

Errors can originate from multiple phases. The overlay (Phase 6) must handle all types.

### Error Types Taxonomy

1. **Compilation Errors** (Phase 4)
   - TypeScript syntax errors
   - CSS parse errors
   - Source: esbuild transform failures

2. **Template Runtime Errors** (Phase 3)
   - Go template execution failures
   - Demo HTML parse errors
   - Source: demo rendering failures

3. **Server Errors** (Phases 1-5)
   - File not found (404)
   - Internal server errors (500)
   - Source: HTTP handler failures

4. **Manifest Generation Errors** (Phase 0+)
   - Parse errors in source files
   - Validation failures
   - Source: generate package

### Error Display Strategy

**In browser overlay** (✅ Implemented in Phase 4):
- **User code errors** (compilation, templates): Primary block with source-mapped stack traces
- **Server errors**: Secondary block, less prominent
- **ENOENT**: Special 404 screen (not an overlay)

**Error message requirements** (✅ All implemented):
- ✅ Pretty formatting with syntax highlighting
- ✅ Stack traces point to original files (not transformed)
- ✅ Source maps for TypeScript errors
- ✅ Clear distinction between error types
- ✅ File path links in error overlay
- ✅ Error details preserved during navigation

### Error Collection Mechanism

Each phase should:
1. Catch errors at appropriate boundaries
2. Format errors with context (file, line, column)
3. Send to error aggregator
4. Error aggregator broadcasts via WebSocket

## WebSocket Events

Live reload uses WebSocket for server-to-client communication.

### Event Types

#### `reload`
Full page reload.

**Payload**:
```json
{
  "type": "reload",
  "reason": "file-change" | "manifest-update" | "config-change",
  "files": ["path/to/changed/file.ts"]
}
```

**Triggers**:
- Source file changes (most common)
- Manifest regeneration completes
- Config file changes

#### `hmr` (STRETCH GOAL)
Hot module replacement without full reload.

**Payload**:
```json
{
  "type": "hmr",
  "modules": ["@rhds/elements/rh-button.js"]
}
```

**Triggers**:
- Module-level changes when HMR is enabled
- See `99-STRETCH-GOALS.md` for implementation details

#### `manifest-updated`
Manifest data changed, UI may need updates.

**Payload**:
```json
{
  "type": "manifest-updated",
  "elementCount": 23,
  "demoCount": 47
}
```

**Triggers**:
- Manifest regeneration completes
- Used by manifest viewer UI (Phase 6, stretch)

#### `error`
Server error occurred, show in overlay.

**Payload**:
```json
{
  "type": "error",
  "category": "compilation" | "template" | "server" | "manifest",
  "message": "Error message",
  "file": "path/to/file.ts",
  "line": 42,
  "column": 10,
  "stack": "..."
}
```

**Triggers**:
- Any error type from taxonomy above
- Consumed by error overlay (Phase 6)

### WebSocket Lifecycle

**Server-side**:
- Endpoint: `/__cem-reload`
- Manage client connections (add/remove on connect/disconnect)
- Broadcast events to all connected clients
- Debounce broadcasts (150ms) to prevent spam

**Client-side**:
- Auto-reconnect on disconnect (exponential backoff)
- Handle each event type appropriately
- Injected via `<script>` tag in all HTML responses

**Connection limits**: Reasonable max (e.g., 100 concurrent clients for dev server).

See `05-AUTORELOAD-HMR.md` for detailed WebSocket implementation.

## Cache Management

Transform middleware (Phase 4) uses in-memory caching.

### Cache Strategy

**Cache key**:
```go
type CacheKey struct {
    Path     string
    ModTime  time.Time
    Size     int64
}
```

**Cache invalidation**:
- File changes invalidate cache for that file
- Dependency changes invalidate dependent files
- Use dependency graph from PR #98 for transitive invalidation

**Cache limits**:
- Max size: 500MB (configurable)
- Eviction: LRU (Least Recently Used)
- Metrics: Log cache hits/misses periodically

### Cache Interactions

- **Phase 1**: File watcher triggers cache invalidation
- **Phase 4**: Transform middleware reads/writes cache
- **Logging**: Cache hit/miss rates in metrics

## Graceful Shutdown

Handle SIGINT/SIGTERM cleanly.

### Shutdown Sequence

1. Stop accepting new connections
2. Close WebSocket connections gracefully
3. Wait for in-flight HTTP requests (timeout: 5s)
4. Clean up file watchers
5. Flush logs
6. Exit

**Timeout**: If requests don't complete within 5s, force shutdown.

## Port Auto-Increment

**Behavior**:
- Default port: 8000
- If taken, try 8001, 8002, ..., up to 8010
- Log actual port prominently on startup
- Exit with error if entire range exhausted

**User consideration**: Demos should not hardcode `localhost:8000` URLs.

## Logging Format

```
[cem serve] Starting dev server on http://localhost:8000
[cem serve] Manifest generated (23 elements, 47 demos)
[cem serve] Watching 128 files for changes
[cem serve] GET /components/pf-button/demo/variants/ 200 12ms
[cem serve] File changed: elements/pf-button/pf-button.ts
[cem serve] Manifest regenerated (1 file changed)
[cem serve] Broadcast reload event to 3 clients
[cem serve] Cache stats: 89% hit rate (234 hits, 28 misses)
```

**Machine-readable format**: Out of scope for initial release (see stretch goals).
