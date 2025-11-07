## Import Maps

### Open Questions

- Do we trace through sources like `@jspm/generator` does?
- Or Do we only generate based on dependencies in package.json, from node_modules, with input map as a fallback?
- How do we handle trailing slash?
- Should we make import map generation out of band on the first pass?

### Auto-Generation Strategy
```json
{
  "imports": {
    "@workspace/pkg-name": "/node_modules/@workspace/pkg-name/index.js",
    "lit": "/node_modules/lit/index.js",
    "./element.js": "/elements/my-element/element.js"
  }
}
```

**Auto-discovery:**
1. Scan `package.json` workspaces (if monorepo)
2. Parse `package.json` dependencies
3. Resolve `exports` field or `main` field
4. Generate `/node_modules/...` mappings

**Override File** (e.g., `import-map.json`):
```json
{
  "imports": {
    "react": "https://esm.sh/react@18"
  }
}
```

### Config keys
under `.serve.importMap`
```yaml
inputMap: ./import-map.json
# or
inputMap:
  imports:
    lit: https://esm.sh/lit
generate: true # based on package.json
```

**Merging:** User overrides win, auto-generated fills gaps.

**Injection:** Middleware injects `<script type="importmap">` into HTML responses.

