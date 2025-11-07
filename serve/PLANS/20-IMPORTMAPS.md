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

---

## Open Questions & Concerns

### Unresolved Open Questions
**Critical:** Lines 3-8 list open questions but provide no answers:

1. **"Do we trace through sources like `@jspm/generator` does?"**
   - Tracing = analyze import statements in source files to discover transitive dependencies
   - Pro: Complete dependency graph, finds all needed packages
   - Con: Complex, slow, requires parsing all JavaScript
   - **Recommendation:** Start simple (package.json only), add tracing later if needed

2. **"Or Do we only generate based on dependencies in package.json, from node_modules?"**
   - Pro: Fast, simple, deterministic
   - Con: Misses dynamically imported modules not in package.json
   - **Recommendation:** This is sufficient for 90% of cases

3. **"How do we handle trailing slash?"**
   - Import `lit/` vs `lit` are different in import maps
   - Node resolution allows both, browsers require exact match
   - **Options:**
     - Generate both with and without trailing slash
     - Normalize all to no trailing slash
     - Match package.json `exports` field exactly
   - **Recommendation:** Follow package.json `exports`, fallback to no slash

4. **"Should we make import map generation out of band on the first pass?"**
   - Out of band = generate async, serve placeholder until ready
   - Pro: Faster server startup
   - Con: Demos may fail on first load
   - **Recommendation:** Generate synchronously on startup, cache aggressively

### package.json `exports` Field Resolution
**Missing:** The plan mentions "Resolve `exports` field or `main` field" (line 24) but:
- No handling for conditional exports (`node`, `browser`, `import`, `require`)
- No handling for subpath patterns (`./utils/*`)
- No handling for `exports: null` (blocks package access)

**Complex example:**
```json
{
  "exports": {
    ".": {
      "import": "./esm/index.js",
      "require": "./cjs/index.js"
    },
    "./utils/*": "./esm/utils/*.js"
  }
}
```

**Recommendation:**
- Use `import` condition for ESM (ignore `require`)
- Expand subpath patterns to specific mappings if possible
- Warn and skip packages with `exports: null`

### Missing or Malformed node_modules
**Issue:** Auto-generation assumes node_modules exists and is valid (line 24-25).
- What if `npm install` hasn't been run?
- What if package.json dependency is missing from node_modules?
- What if package has no `exports` or `main` field?

**Recommendation:**
- Validate node_modules exists before generation
- Warn about missing dependencies
- Fallback to user-provided import map for missing packages

### Scoped Package Path Structures
**Missing:** Scoped packages (@org/pkg) have different paths than unscoped.
```
/node_modules/@org/pkg/index.js  vs  /node_modules/pkg/index.js
```
- Need special handling for `@` in package names
- Path construction differs

**Recommendation:** Test with scoped packages, document path generation rules.

### Merge Strategy Details
**Issue:** "User overrides win, auto-generated fills gaps" (line 47) is ambiguous.

**Example conflict:**
```yaml
# Auto-generated
imports:
  lit: /node_modules/lit/index.js

# User override
imports:
  lit: https://esm.sh/lit
```

**Questions:**
- Is this a top-level merge or per-import key merge?
- What about scopes?
- What if user provides partial path that conflicts?

**Recommendation:** Specify merge algorithm:
```
1. Start with empty map
2. Add all auto-generated imports
3. For each user import:
   - If key exists, replace (user wins)
   - Else add new key
```

### Import Map Injection Timing
**Issue:** Plan says "Middleware injects `<script type="importmap">`" but:
- Import maps must appear before any module scripts
- What if HTML already has an import map?
- What if HTML is malformed (no `<head>`)?

**Recommendation:**
- Parse HTML to find `<head>`
- Insert import map as first child of `<head>`
- Merge with existing import map if present
- Error if no `<head>` found

### Workspace/Monorepo Support
**Issue:** Line 22 mentions "Scan `package.json` workspaces" but:
- No detail on how to resolve workspace packages
- What if workspace package has same name as npm package?
- How to handle cross-workspace imports?

**Example:**
```json
// Root package.json
{
  "workspaces": ["packages/*"]
}

// packages/ui/package.json
{
  "name": "@myapp/ui",
  "main": "dist/index.js"
}
```

**Recommendation:**
- Resolve workspace packages to local paths (not /node_modules)
- Prioritize workspace packages over node_modules
- Document workspace resolution order

### Bare Module Specifier Edge Cases
**Missing:** Handling for non-standard module specifiers:
- Relative imports already in source (`./foo.js`)
- Absolute URLs (`https://cdn.example.com/lib.js`)
- Data URLs (`data:text/javascript,...`)

**Recommendation:** Only process bare specifiers, ignore others.

### Config Merging: CLI vs File vs Default
**Missing:** When multiple config sources provide import maps:
- CLI flag: `--import-map ./dev.json`
- Config file: `serve.importMap.inputMap: ./prod.json`
- Default: auto-generated

**Which wins?** Need precedence order.

**Recommendation:**
```
Priority (highest to lowest):
1. CLI flags
2. Config file
3. Auto-generated
```
