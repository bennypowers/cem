Practice TDD. When writing tests, always use the fixture/golden patterns we've established:

- **Fixtures**: Input test data - the content directories/files your code operates on (e.g., `serve/testdata/transforms/http-typescript/simple-greeting.ts`). Use `NewFixtureFS()` helper from `internal/platform/testutil` instead of `os.ReadFile()` to load fixtures into in-memory file systems.
- **Goldens**: Expected output files to compare against (e.g., `serve/middleware/routes/testdata/chrome-rendering/expected-basic.html`). Tests should support `--update` flag to regenerate golden files when intentional changes occur.

Getter methods should be named Foo(), not GetFoo().

Always use Makefile targets for running tests or builds, since they export the necessary env vars.

When commit messages mention AI agents, always use `Assisted-By`, never use `Co-Authored-By`.

When writing CSS, use modern native syntax including nesting, :has, layers, and light-dark()
