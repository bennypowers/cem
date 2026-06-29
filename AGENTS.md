## cem binary usage

When running cem commands against the example projects, you don't need to `cd`, you can `cem ${cmd} -p examples/${example}`.

example:

```shell
$ make
$ dist/cem -p examples/kitchen-sink validate
```

## Go

Getter methods should be named `Foo()`, not `GetFoo()`.

All diagnostic logging (debug, info, warning, error, success, trace) must go through `bennypowers.dev/cem/internal/logging`. This ensures correct routing in CLI vs LSP modes and respects `-v`/`-vv`/`-q` flags. Never call `pterm.Warning`, `pterm.Info`, etc. directly for log messages. Direct pterm usage is reserved for terminal UI primitives only: spinners, live areas, colored/styled display formatting.

Verbosity levels: `-q` (quiet: warnings/errors only), default (+ success), `-v` (+ info), `-vv` (+ debug), `-vvv` (+ trace per-file detail). Use `logging.AtLevel(level)` as a guard before expensive debug computations.

Use `platform.FileSystem` for all filesystem I/O in library code. Never call `os.ReadFile`, `os.WriteFile`, `os.Stat`, `os.MkdirAll`, `os.Create`, `os.Open`, `os.Remove`, `os.Rename`, `os.ReadDir`, `filepath.Glob`, or `os.DirFS` directly -- thread `platform.FileSystem` through function parameters instead. Use `platform.DirFS(fsys, dir)` instead of `os.DirFS(dir)`. For `os.Getwd`/`os.UserHomeDir`, pass the result as a parameter from CLI entry points rather than calling from internal packages. Exception: `cmd/` package entry points may call `os.Getwd()`/`os.UserHomeDir()` since they are CLI-specific and pass results downward. Build tools with `//go:build ignore` (e.g., `serve/generate_elements.go`, `mcp/tools/gen-adapters/`) are also exempt.

Use `platform.WalkDir` for all directory traversals. It silently prunes `.git` directories and accepts additional dirs to skip via `set.Set[string]`. Never use `filepath.WalkDir` or `fs.WalkDir` directly -- those bypass `.git` pruning.

Run `go vet` to surface gopls suggestions. Common examples:
- replace `interface{}` with `any`
- replace `if/else` with `min`
- replace `m[k]=v` loop with `maps.Copy` [mapsloop]
- Loop can be simplified using slices.Contains [slicescontains]

### CLI commands

In cobra command `RunE` handlers, use `cmd.Println` / `cmd.PrintErrln` instead of `fmt.Println` / `fmt.Fprintln(os.Stderr, ...)`. This respects cobra's configured output streams and makes commands testable.

## Debugging

When debugging Go code, always use the logger `logger.Debug`, etc. Don't use `fmt.Printf`, which pollutes stdio, breaking the LSP and MCP commands.

To print debug logs to the console, pass `-vv` (debug) or `-vvv` (trace)

## Testing

Practice TDD. See new tests fail first, before implementing the fix/feature.

IMPORTANT: When writing tests, **always** use fixture/golden pattern. We **always** prefer goldens, STRONG justification needed to permit inline assertions. If using inline assertions, comment why in test func.

- **Fixtures**: Input test data - the content directories/files your code operates on (e.g., `serve/testdata/transforms/http-typescript/simple-greeting.ts`).
- **Goldens**: Expected output files to compare against (e.g., `serve/middleware/routes/testdata/chrome-rendering/expected-basic.html`). Tests should support `--update` flag to regenerate golden files when intentional changes occur.
- **Always use MapFS for test data** (unit/integration tests): Load all fixture and golden files into `platform.MapFileSystem` via `testutil.NewFixtureFS()` or `testutil.LoadTestdataFS()`. Never call `os.ReadFile()` or `os.Open()` on test data in test code. Frontloading disk I/O into MapFS at setup time ensures consistency across packages and improves CI performance.
- **Exception: E2E tests** (`//go:build e2e` in `cmd/`): These run the built binary as a subprocess and compare stdout against golden files on disk. Direct `os.ReadFile`/`os.WriteFile` is expected for golden comparison and `--update` regeneration since MapFS cannot cross process boundaries.
- Always use Makefile targets for running tests or builds, since they export the necessary env vars.
- For LSP tests, ALWAYS use `testutil.RunLSPFixtures()` - NEVER use direct `os.ReadFile()` calls.
- If frontend visual regression tests fail consistently, check if it's merely 
  an issue of font rendering.

### LSP Testing Patterns

- **Fixture Structure**: Each test scenario is a subdirectory containing:
  - `input.html` or `input.ts` (required)
  - `manifest.json` (optional)
  - `expected.json` or `expected-*.json` (optional, for assertions)
  - `package.json` pointing `customElements` block to `manifest.json`

- **Cursor Positions via `^cursor` Marker**: Position-dependent LSP tests (hover, completion, definition, references) embed cursor position directly in the fixture input file using a comment marker. The `^` column is the cursor character; the line above is the cursor line. Use language-appropriate comment syntax:

  HTML:
  ```html
  <test-element .testAttr="hello"></test-element>
                <!-- ^cursor -->
  ```

  TypeScript (inside Lit template literals, use HTML comments):
  ```typescript
  const tpl = html`
    <test-element .testAttr="hello"></test-element>
    <!--          ^cursor -->
  `;
  ```

  TypeScript (outside template literals, use JS comments):
  ```typescript
  import { html } from 'lit';
  /*        ^cursor */
  ```

  CSS:
  ```css
  my-element::part(button) {
                /* ^cursor */
  ```

  For HTML fixtures, the loader extracts cursors automatically. For TS fixtures, tests call `fixture.ParseCursor(testhelpers.TSCursorParser)` since tree-sitter is needed to find template literal boundaries. For CSS, use `fixture.ParseCursor(testhelpers.CSSCursorParser)`.

  The fixture loader strips the marker line from `InputContent` and populates `fixture.Cursor *protocol.Position`. Tests use `fixture.Cursor` directly instead of maintaining a separate Go map of cursor positions. When no marker is present, `Cursor` is nil.

  **Fallback**: When a comment marker isn't practical (cursor inside a value, ambiguous column), use YAML frontmatter. Avoid frontmatter in `.ts` files as it causes biome/eslint issues:
  ```html
  ---
  cursor:
    line: 0
    character: 14
  ---
  <test-element .testAttr="hello"></test-element>
  ```

- **RunLSPFixtures Pattern**:
  ```go
  testutil.RunLSPFixtures(t, "testdata/my-test-suite", func(t *testing.T, fixture *testutil.LSPFixture) {
      // Setup context
      ctx := testhelpers.NewMockServerContext()

      // Use fixture.InputContent, fixture.Manifest, fixture.Cursor, etc.
      // Load expected data with fixture.GetExpected("key", &expected)
  })
  ```

- **Multiple Expected Files**: Use `expected-variant.json`, `expected-size.json` pattern for testing multiple positions/cases in one fixture

- **Regression Test Isolation**: Keep regression test fixtures in separate directories (e.g., `testdata-regression/`) to avoid interference with standard test discovery

- **Filesystem in Tests**: Test code follows the same `platform.FileSystem` rules as production code. Load fixture data via `testutil.NewFixtureFS()`, `testutil.LoadTestdataFS()`, or `testutil.LoadFixtureFile()` -- never via direct `os.ReadFile`/`os.Open` calls. When constructing `MCPContext` or similar objects that accept a `platform.FileSystem`, pass the workspace's `MapFileSystem` instead of `nil` to exercise the injected FS path. Two exceptions: (1) OS integration subtests (e.g., filewatcher tests) that must write to real temp dirs via `t.TempDir()`, and (2) E2E tests (`//go:build e2e` in `cmd/`) that run the built binary as a subprocess and compare stdout against golden files on disk.

- **Error Path Testing**: When testing filesystem error paths that rely on OS permissions (e.g., unreadable files), use an error-injecting `FileSystem` wrapper instead of `os.Chmod`. MapFS does not enforce permissions, so permission-based error tests need a wrapper that returns errors for specific paths.

### Testing Tiers

See `docs/TESTING.md` for the full testing convention. In brief:
- **Tier 1 (pure functions):** exhaustive table-driven tests, in-package. A pure function is one where all effects depend on its formal parameters.
- **Tier 2 (logic with effects):** fixture-based tests with real dependencies (e.g., QueryManager). No percentage target.
- **Tier 3 (thin wiring):** covered by integration/e2e tests. No separate unit tests required.

New LSP methods or MCP tools must include at least Tier 2 tests before merging.

## Per-package guidelines

- Docs: When working on the docs site, refer to `docs/CLAUDE.md` for docs-specific guidelines.
- Generate and Validate: When touching code related to the custom-elements manifest schema, double-check with the canonical TypeScript schema.d.ts at https://github.com/webcomponents/custom-elements-manifest/. This is to avoid hallucinations. Note that different schema versions may have different types, and we'll have to deal with that if they do.

### LSP

- **Tree-sitter concurrency**: Never share `QueryMatcher` between goroutines. Cache queries, create fresh cursors per operation.
- **Nil safety**: Check for nil `DocumentManager` before use; test with nil parameters.
- **Protocol compliance**: Match LSP return types exactly; declare all capabilities in `initialize`.
- **Testing**: Use `package lsp_test` for public API tests. Use `testhelpers.NewMockServerContext()` for test context. Fixtures in `testdata/` dirs, regression fixtures in `testdata-regression/`.
- When adding LSP methods, implement diagnostics + code actions together and update server capabilities.

### MCP

- **Design principle**: Data + context + LLM decision-making. Provide rich manifest data and contextual templates; let the LLM decide. No hardcoded suggestions.
- **Respect element authors**: Use their documented constraints and intent. Follow documented accessibility patterns, don't add generic ARIA advice.
- **Template naming**: Use dashes (e.g., `element-attributes.md`), not underscores. Reference subresources (`cem://element/{tagName}/attributes`), not tool names.

### Config

When adding or modifying configuration fields, update all three locations:
1. **Config struct** in `internal/config/load.go` -- add the field with proper struct tags
2. **Init wizard** in `cmd/config_init.go` -- add form fields, detection, and apply logic
3. **Docs** in `docs/content/docs/reference/configuration.md` -- document the new option

If the field is detectable from the project (e.g. file patterns, git remote), add detection in `cmd/config_init_detect.go`. If it needs validation, add it to `internal/config/validate.go` and wire it into the wizard's `ValidateFn`.

## Git

- Always run `make lint` and address issues before committing
- Always check go LSP diagnostics and address them before committing
- When commit messages mention AI agents, always use `Assisted-By`, never use `Co-Authored-By`.

## Frontend

When writing CSS, use modern native syntax including nesting, :has, layers, and light-dark()

When writing web components (e.g. for the dev server)
- prefer ids to classes in shadow roots for unique elements
- don't use `class` on the host as a public api, use attributes instead
- don't dispatch new CustomEvent with details, instead create custom classes which extend Event, and have class fields for attached state.

## HTML

Format attributes vertically aligned and sort by importance, e.g.:

```html
<button id="button"
        part="button"
        class="${classMap({ loading })}"
        aria-disabled="${String(disabled)}"
        @click="${this.#onClick}">
```

When manipulating HTML via JavaScript, prefer template cloning and the DOM API
to the innerHTML setter.

## CSS
- Use native CSS nesting syntax.
- but don't nest under :host except when using host attrs like `:host([disabled]) { button { ... } }`
- Prefer `light-dark()` to `prefers-color-scheme`
- When writing stylesheets for shadow roots, prefer ID selectors to classes
