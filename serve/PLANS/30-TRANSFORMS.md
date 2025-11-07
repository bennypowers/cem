## Build Transformations

### TypeScript (esbuild Go API)
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
- Inline source maps for debugging
- Responds to `.js` URLs but reads `.ts` files (like pfe-tools router)
- Cache transformed output (invalidate on file change)
- Target is configurable via switch/config combo

### CSS Transformation (Optional)
```go
// Transform .css files to CSSStyleSheet
const sheet = new CSSStyleSheet()
sheet.replaceSync(originalCss);
export default sheet;
```

**Implementation:**
- Simple template wrapper (no complex parsing)
- Content-type: `application/javascript; charset=utf-8`
- Optional flag: `--transform-css` takes a list of glob patterns to include (doublestar, negations)
  - e.g. user wants to transform element css (added to lit static styles list)  
    but does not want to transform page css (linked in document)

**No plugin system** - these are the only transforms, hard-coded and fast.

### Config
under `.serve.transforms`
```yaml
typescript:
  enabled: true
  target: es2022
css:
  enabled: true
  include:
    # - elements/**/*.css
  exclude:
    # - docs/styles/*.css
```

---

## Open Questions & Concerns

### TypeScript Transform Without Type Checking
**Issue:** esbuild transforms TypeScript to JavaScript **without type checking** (lines 4-11).
- Fast but silently hides type errors
- Developers may not realize code has type errors until separate type check

**Example:**
```typescript
// This transforms successfully but has type error
const x: string = 123; // esbuild doesn't complain
```

**Recommendation:**
- Add warning in logs when TypeScript files are served
- Suggest running `tsc --noEmit` separately
- Consider optional type checking mode (slower but safer)

### Transform Cache Dependencies
**Issue:** "Cache transformed output (invalidate on file change)" (line 19) doesn't handle dependencies.

**Problem:** If `a.ts` imports `b.ts`:
```typescript
// a.ts
import { foo } from './b.js';
```

When `b.ts` changes:
- Cache for `b.ts` is invalidated ✓
- Cache for `a.ts` is NOT invalidated ✗ (stale!)

**Recommendation:**
- Build dependency graph from imports
- Invalidate transitive dependents when file changes
- Or: simpler approach - invalidate entire cache on any change (trade memory for correctness)

### Cache Key Strategy
**Missing:** How cache keys are generated:
- File path only? (misses content changes)
- File path + mtime? (can miss rapid changes within same second)
- File path + content hash? (expensive)

**Recommendation:**
```go
type CacheKey struct {
    Path     string
    ModTime  time.Time
    Size     int64
}
```
This is fast and reliable enough for dev server.

### Cache Memory Limits
**Issue:** No bounds on cache size (lines 19, 00-OVERVIEW.md:68).

**Footgun:** Large codebase + many transforms = OOM.

**Example scenario:**
- 1000 TypeScript files
- Average 50KB transformed
- 1000 × 50KB = 50MB (manageable)

But:
- 10,000 files × 100KB = 1GB (not manageable)

**Recommendation:**
- Add max cache size (e.g., 500MB)
- Implement LRU eviction
- Log cache stats (hits, misses, size) periodically

### CSS Transform Browser Compatibility
**Issue:** CSS transform uses Constructable Stylesheets (lines 23-28):
```javascript
const sheet = new CSSStyleSheet();
sheet.replaceSync(originalCss);
```

**Browser support:**
- ✓ Chrome 73+
- ✓ Safari 16.4+
- ✓ Firefox 101+
- ✗ IE 11
- ✗ Older mobile browsers

**Footgun:** Demos using transformed CSS won't work in older browsers.

**Recommendation:**
- Document browser requirements
- Provide fallback for older browsers (style tag injection)
- Or add flag: `--css-transform-style=constructable|link`

### CSS URL Rewriting
**Issue:** CSS transform wraps raw CSS (line 26) but doesn't handle relative URLs.

**Problem:**
```css
/* elements/button/button.css */
.button {
  background: url('../shared/icon.svg');
}
```

When transformed to JS module, relative URL breaks (resolved from wrong base).

**Recommendation:**
- Parse CSS for `url()` references
- Rewrite to absolute paths based on CSS file location
- Or document that CSS transform doesn't support relative URLs

### CSS Transform Error Handling
**Issue:** `sheet.replaceSync(originalCss)` (line 26) throws on invalid CSS.

**Example:**
```css
/* Malformed CSS */
.foo { color: ; }
```

**Result:** JavaScript error in browser, no graceful degradation.

**Recommendation:**
- Validate CSS before transformation
- Return error to client (don't crash browser)
- Show error overlay with CSS syntax issue

### Transform Target Configuration
**Issue:** Config shows `target: es2022` (line 44) but:
- No validation of target values
- What if user provides invalid target? (`target: es2099`)
- What's the default if not specified?

**Recommendation:**
- Validate target against esbuild's allowed values
- Document supported targets (ES2015-ES2023, ESNext)
- Set sensible default (ES2020 for broad compatibility)

### Source Map Handling
**Issue:** Plan specifies "Inline source maps" (line 10) but:
- Inline source maps increase file size significantly
- No option for external source maps
- What about original source serving? (debugger needs it)

**Recommendation:**
- Default to inline for simplicity
- Add `--source-maps=inline|external|none` flag
- Serve original `.ts` files at `/__source/` for debugger

### TypeScript Configuration Ignored
**Issue:** esbuild transform doesn't read `tsconfig.json`.
- User's `compilerOptions` are ignored
- `paths` mappings don't work
- `jsx` settings might not match

**Example:**
```json
// tsconfig.json
{
  "compilerOptions": {
    "jsx": "react-jsx",
    "paths": {
      "@/*": ["src/*"]
    }
  }
}
```

**Recommendation:**
- Document that tsconfig.json is not used
- Consider parsing tsconfig.json for `target` and `jsx` settings
- Or tell users to use esbuild config format

### Transform Concurrency
**Missing:** Limits on concurrent transform operations.
- What if 100 requests arrive for different TypeScript files?
- esbuild is CPU-bound, parallel transforms compete

**Recommendation:**
- Add worker pool for transforms (e.g., max 4 concurrent)
- Queue excess requests
- Return 503 if queue is full

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