# CEM LSP Development Guidelines

## Instructions to Claude

### LSP Standards
Use LSP standards first; prefer `$/setTrace` for logging; use standard `window/logMessage`/`window/showMessage`

### Architecture
Each LSP method in its own package following `methods/textDocument/methodName/` pattern. Use adapter pattern with focused context interfaces.

### Documentation
- Document user-facing features in `../docs/content/docs/lsp.md` and `README.md`
- Focus on benefits and usage, not implementation details
- Update `lsp/CLAUDE.md` and `lsp/ARCHITECTURE.md` for major changes

### Testing Requirements
- New features need integration tests; bug fixes need regression tests
- Use `package lsp_test` to test only public APIs
- Use fixture files in organized directories, not inline content
- Use `testhelpers.NewMockDocument(content)` for mock documents

### Benchmark Validation
- Benchmarks in `lsp/benchmark/` validate LSP correctness and performance
- Run `make bench-lsp` before major releases
- Check `lsp/benchmark/PERFORMANCE_ANALYSIS.md` for baselines
- Each LSP operation has its own benchmark module with P50/P95/P99 analysis

## Common Pitfalls

### Nil Safety
Check for nil `DocumentManager` before use; test with nil parameters

### Protocol Compliance
Match LSP return types exactly; declare all capabilities in `initialize`

### Tree-Sitter Concurrency
**CRITICAL**: Never share `QueryMatcher` between goroutines. Cache queries, create fresh cursors per operation.

### Integration
Implement diagnostics + code actions together; update server capabilities when adding methods

## Repository Conventions

### File Organization
- LSP methods: `methods/textDocument/methodName/methodName.go`
- Tests: `methodName_test.go` in same package, using `package methodName_test`
- Fixtures: Organized in `test-fixtures/` or `methodName-test/` directories
- Types: Shared types in `lsp/types/` package
- Helpers: Utility functions in `lsp/helpers/` package

### Code Style
- Use `defer func() { _ = resource.Close() }()` for resource cleanup
- Error strings should be lowercase and not end with punctuation
- Function parameters: error should be last parameter
- Prefer built-in functions (e.g., `max()` in Go 1.25+) over custom implementations

### Tree-Sitter Patterns
- Cache queries at startup, create fresh cursors per operation
- Use `queries.GetGlobalQueryManager()` for shared query infrastructure
- Prefer tree-sitter parsing over regex for HTML/TypeScript content
- Use parser pools for performance

### Testing the Server
```bash
go build -o cem .        # Build
./cem lsp --help         # Verify command
go test ./lsp/...        # Run tests
./cem lsp                # Start server (stdio)
```

