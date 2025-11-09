## Import Maps

## TDD First Steps

1. Remove "not implemented" guard tests from existing tests for this phase
2. Write comprehensive tests for real functionality (will fail - true red phase)
3. Implement to make tests pass

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

**Merging:** User overrides win, auto-generated fills gaps. Algorithm is deep merge, keeping user overrides at each layer.

**Injection:** Middleware injects `<script type="importmap">` into HTML responses before any demo scripts are parsed.

### Injection Timing

Import map is injected during demo HTML template rendering (Phase 3):

1. Server renders demo chrome template (see [15-DEMO-CHROME.md](./15-DEMO-CHROME.md))
2. Import map JSON is generated (auto + overrides merged)
3. Import map is injected as **first child of `<head>`** via template variable `{{.ImportMap}}`
4. Demo HTML content is inserted into light DOM
5. Demo scripts execute with import map already available

This ensures import map is established before any module scripts in the demo execute.

**Important constraint**: User demos should not already contain an import map. If they do, the page will fail since browsers only honor the first import map. For this reason, import map generation is configurable (can be disabled via config).


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

- Use `import` condition for ESM (ignore `require`)
- Expand subpath patterns to specific mappings if possible, preferring to 
normalize to trailing-slash
- Skip packages with `exports: null`. Warn if requested


### Dependency tracing like `@jspm/generator`
Out of scope.
 - Tracing = analyze import statements in source files to discover transitive dependencies
 - Pro: Complete dependency graph, finds all needed packages
 - Con: Complex, slow, requires parsing all JavaScript
 - **Conclusion:** Start simple (package.json only), add tracing later if needed

### Missing or Malformed node_modules
**Issue:** Auto-generation assumes node_modules exists and is valid (line 24-25).
- What if `npm install` hasn't been run?
- What if package.json dependency is missing from node_modules?
- What if package has no `exports` or `main` field?

In all cases:
  - Warn the user
  - Fallback to user-provided import map for missing packages

### Workspace/Monorepo Support

PR #127 added workspace/monorepo support. We will rely on code there for resolving workspaces, possibly abstracting shared code to a common package.

**Recommendation:**
- Resolve workspace packages to local paths (not /node_modules)
- Prioritize workspace packages over node_modules
- Document workspace resolution order

### Module Specifier Edge Cases
Handling for non-standard module specifiers:
- Relative imports already in source (`./foo.js`)
- Absolute URLs (`https://cdn.example.com/lib.js`)
- Data URLs (`data:text/javascript,...`)

Only process bare specifiers, ignore others.

### Config Merging: CLI vs File vs Default

Priority (highest to lowest):
1. CLI flags
2. Config file
3. Auto-generated

---

## Acceptance Criteria

- [x] Import map auto-generated from package.json dependencies
- [x] Import map includes workspace packages (monorepo support)
- [x] package.json `exports` field resolved (prioritize `import` condition)
- [x] package.json `main` field used as fallback when `exports` missing
- [x] Subpath patterns in `exports` handled correctly
- [x] User override file (import-map.json) merges with auto-generated map
- [x] User overrides win in merge conflicts (deep merge algorithm)
- [x] Import map injected into HTML `<head>` as `<script type="importmap">` (default index)
- [x] Workspace packages resolve to local paths (not /node_modules)
- [x] Workspace packages prioritized over node_modules
- [x] Missing node_modules dependencies logged as warnings
- [x] Packages with no `exports` or `main` field logged as warnings
- [x] Only bare specifiers processed (relative, absolute URLs, data URLs ignored)
- [x] Config priority enforced: CLI > file > auto-generated
- [ ] Demos containing import maps fail with helpful error message (Phase 3 - demo rendering)
- [x] Tests cover exports resolution, workspace support, merge algorithm
