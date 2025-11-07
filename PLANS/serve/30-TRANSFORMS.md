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
