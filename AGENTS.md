## Go

Getter methods should be named `Foo()`, not `GetFoo()`.

When adding debug logs in go code, always use `bennypowers.dev/cem/internal/logging`, to avoid polluting stdio.

Run `go vet` to surface gopls suggestions. Common examples:
- replace `interface{}` with `any`
- replace `if/else` with `min`
- replace `m[k]=v` loop with `maps.Copy` [mapsloop]
- Loop can be simplified using slices.Contains [slicescontains]

## Debugging

When debugging Go code, always use the logger `logger.Debug`, etc. Don't use `fmt.Printf`, which pollutes stdio, breaking the LSP and MCP commands.
To print debug logs to the console, pass `-vvv`

## Testing

Practice TDD. When writing tests, always use the fixture/golden patterns we've established:

- **Fixtures**: Input test data - the content directories/files your code operates on (e.g., `serve/testdata/transforms/http-typescript/simple-greeting.ts`). Use `NewFixtureFS()` helper from `internal/platform/testutil` instead of `os.ReadFile()` to load fixtures into in-memory file systems.
- **Goldens**: Expected output files to compare against (e.g., `serve/middleware/routes/testdata/chrome-rendering/expected-basic.html`). Tests should support `--update` flag to regenerate golden files when intentional changes occur.
- Always use Makefile targets for running tests or builds, since they export the necessary env vars.
- For LSP tests, ALWAYS use `testutil.RunLSPFixtures()` - NEVER use direct `os.ReadFile()` calls.

### LSP Testing Patterns

- **Fixture Structure**: Each test scenario is a subdirectory containing:
  - `input.html` or `input.ts` (required)
  - `manifest.json` (optional)
  - `expected.json` or `expected-*.json` (optional, for assertions)
  - `package.json` pointing `customElements` block to `manifest.json`

- **RunLSPFixtures Pattern**:
  ```go
  testutil.RunLSPFixtures(t, "testdata/my-test-suite", func(t *testing.T, fixture *testutil.LSPFixture) {
      // Setup context
      ctx := testhelpers.NewMockServerContext()

      // Use fixture.InputContent, fixture.Manifest, etc.
      // Load expected data with fixture.GetExpected("key", &expected)
  })
  ```

- **Multiple Expected Files**: Use `expected-variant.json`, `expected-size.json` pattern for testing multiple positions/cases in one fixture

- **Regression Test Isolation**: Keep regression test fixtures in separate directories (e.g., `testdata-regression/`) to avoid interference with standard test discovery

## Per-package guidelines

- LSP: When working on LSP features, refer to `lsp/CLAUDE.md` for LSP-specific guidelines.
- MCP: When working on MCP features, refer to `mcp/CLAUDE.md` for MCP-specific guidelines.
- Docs: When working on the docs site, refer to `docs/CLAUDE.md` for docs-specific guidelines.

## Git

When commit messages mention AI agents, always use `Assisted-By`, never use `Co-Authored-By`.

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
