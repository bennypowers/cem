# Testing Convention

## Pure Function Definition

A **pure function** is one where all observable effects depend on its formal parameters: given identical parameters, it always returns the same result and has no effect other than returning its result.

Pure functions are the easiest to test: given params, expect return.

## Testing Tiers

### Tier 1: Pure Functions

**Coverage:** exhaustive (all code paths exercised).

**Test style:** table-driven tests in the same package, using `testify/assert`.

**Example:** `lsp/helpers/position_test.go` tests `IsPositionInRange` with boundary positions, single/multi-line ranges, and edge cases.

**When to use:** any function that takes values, performs no I/O, and returns deterministic results. Functions taking protocol types (e.g., `protocol.Position`) still count as pure if they satisfy the definition.

**Deduplication:** when identical pure functions exist across packages, consolidate to a shared location (e.g., `lsp/helpers/`) and test once. Methods with unused receivers should be refactored to package-level functions.

### Tier 2: Logic with External State/Effects

**Coverage:** each distinct behavior exercised via fixture-based tests. No percentage target; coverage is a side effect of thorough fixtures, not a goal.

**Test style:** fixture HTML/TS files in `testdata/` directories, real `QueryManager` via `treesitter.NewQueryManager(treesitter.LSPQueries())`, validation functions or golden JSON for assertions.

**Example:** `lsp/document/html/handler_test.go` tests `ParseScriptTags` by loading fixture HTML, creating a real handler with a QueryManager, and validating the parsed output.

**When to use:** handler methods that use tree-sitter queries, functions that read state from document trees, methods with meaningful side effects.

**QueryManager setup pattern:**
```go
qm, err := treesitter.NewQueryManager(treesitter.LSPQueries())
require.NoError(t, err)
defer qm.Close()

handler, err := html.NewHandler(qm)
require.NoError(t, err)
defer handler.Close()
```

### Tier 3: Thin Wiring/Adapters

**Coverage:** exercised by parent-package integration tests or e2e tests. No separate unit tests required.

**Examples:** `CreateDocument` methods that wire handler creation to the document manager, protocol adapter methods that translate between types.

**When to use:** code that delegates to Tier 1/2 functions without adding logic. If the function's body is primarily calling other functions and assembling results, it's wiring.

**Cross-package coverage:** some packages (e.g., `lsp/document/blade/`) may show low per-package coverage because their logic is exercised by integration tests in `lsp_test`. This is expected for Tier 3 code and does not indicate missing tests.

## Decision Tree

```text
Is the function pure?
  (all effects depend on params, deterministic, no I/O)
    YES -> Tier 1: table-driven tests, in-package, exhaustive
    NO  -> Does it have testable behavior with fixture setup?
             YES -> Tier 2: fixture-based tests, real dependencies
             NO  -> Tier 3: covered by integration/e2e tests
```

## Merge Gate

New LSP methods or MCP tools must include at least Tier 2 fixture-based tests before merging. This applies after the testing convention is established.

## Existing Patterns

| Pattern | Location | Use for |
|---------|----------|---------|
| Table-driven pure function tests | `lsp/document/base/base_document_test.go` | Tier 1 |
| Fixture-driven LSP tests | `internal/platform/testutil/lsp.go` | Tier 2 (LSP features) |
| Mock server context | `lsp/testhelpers/mock_server_context.go` | Integration tests |
| E2e binary tests | `cmd/lsp_test.go` | Tier 3 / smoke tests |

## When to Write a New Test vs Rely on Integration

Write a new in-package test when:
- the function is pure (always)
- the function has distinct behaviors that integration tests don't specifically exercise
- a bug was found that integration tests didn't catch
- the function is complex enough that reading the test is faster than reading the implementation

Rely on integration coverage when:
- the code is thin wiring with no branching logic
- the function is already well-exercised by `lsp_test` integration tests
- adding an in-package test would duplicate what integration tests verify
