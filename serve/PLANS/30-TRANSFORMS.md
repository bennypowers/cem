## Build Transformations

### TypeScript (esbuild Go API)

Single file transforms without type checking (rely on tsserver)

```go
// Embedded esbuild transform
transform := api.Transform(source, api.TransformOptions{
    Loader: api.LoaderTS,
    Target: api.ES2020,
    Format: api.FormatESM,
    Sourcemap: api.SourceMapInline,
})
```

**Features:**
- Zero external dependencies
- On-the-fly `.ts` → `.js` transformation
- Inline source maps for debugging (size isn't an issue for a dev server)
- Responds to `.js` URLs but reads `.ts` files (like pfe-tools router)
- Target is configurable via switch/config combo
- tsconfig is read and passed to esbuild's tsconfigRaw arg.
- Parallellizing transforms:
  - Add worker pool for transforms (e.g., max 4 concurrent)
  - Queue excess requests
  - Return 503 if queue is full

#### Caching and invalidation

- Build dependency graph for modules (see prior art in PR #98, which creates a 
dependency graph for incremental manifest updates, this should be sufficient for us, perhaps with minor modifications)
- Cache transformed output (invalidate file and it's dependent tree on file change)

```go
type CacheKey struct {
    Path     string
    ModTime  time.Time
    Size     int64
}
```
This is fast and reliable enough for dev server.

- Add max cache size (e.g., 500MB)
- Implement LRU eviction
- Log cache stats (hits, misses, size) periodically

### Config
under `.serve.transforms`

```yaml
typescript:
  enabled: true
  target: es2022
  tsconfig: ./path/to/tsconfig.json
```

#### Transform Target Configuration
- Validate target against esbuild's allowed values
- Document supported targets (ES2015-ES2023, ESNext)
- Set sensible default (ES2020 or ES2022 for broad compatibility)

### CSS Transformation (Optional)
```go
// Transform .css files to CSSStyleSheet
const sheet = new CSSStyleSheet()
sheet.replaceSync(originalCss);
export default sheet;
```

- Browser support for older versions is less important - we target browser baseline 2024 and up.
- Errors in browser are fine for now.

**Implementation:**
- Simple template wrapper (no complex parsing)
- Content-type: `application/javascript; charset=utf-8`
- Optional flag: `--transform-css` takes a list of glob patterns to include (doublestar, negations)
  - e.g. user wants to transform element css (added to lit static styles list)  
    but does not want to transform page css (linked in document)
- document that CSS transform doesn't support relative URLs
- no postcss or sass pipeline, at least not for initial release

**No plugin system** - these are the only transforms, hard-coded and fast.

### Config
under `.serve.transforms`

```yaml
css:
  enabled: true
  include:
    # - elements/**/*.css
  exclude:
    # - docs/styles/*.css
```

### Testing Strategy for Transforms
**Issue:** Complex transform logic needs comprehensive tests:
- Cache hit/miss scenarios
- Dependency invalidation
- Error handling (malformed code)
- Source map correctness

**Recommendation:** Create test fixtures (per CLAUDE.local.md):
```
server/transform/test-fixtures/
  ├── basic/
  │   ├── input.ts
  │   ├── expected.js
  │   └── expected.js.map
  ├── imports/
  │   ├── a.ts (imports b.ts)
  │   ├── b.ts
  │   └── cache-invalidation-test.go
  └── errors/
      ├── syntax-error.ts
      ├── type-error.ts
      └── error-test.go
```
