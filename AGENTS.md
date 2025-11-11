Practice TDD. When writing tests, always write use the fixture/golden patterns we've established. Use the NewFixtureFS() helper instead of os.ReadFile() to create in-memory testing file systems from fixtures.

Getter methods should be named Foo(), not GetFoo().

Always use Makefile targets for running tests or builds, since they export the necessary env vars.

When commit messages mention AI agents, always use `Assisted-By`, never use `Co-Authored-By`.

When writing CSS, use modern native syntax including nesting, :has, layers, and light-dark()
