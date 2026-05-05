# Coverage Gaps Tracker

Target: 80% overall. Current: 57.1% unit / 65.6% combined.

## Strategy

Most uncovered functions fall into categories:
- **Interface implementations** (Name, Label, IsDeprecated, Children, ToTreeNode, etc.) -- repeated across every manifest type. Compact tests that construct a type and call all interface methods cover these in bulk.
- **Clone methods** -- test by constructing, cloning, and asserting independence.
- **Pure helpers** -- table-driven tests.
- **Tier 2 logic** -- needs fixtures/context (skip for 80% target unless high-value).

## Remaining work by package

### High priority (biggest impact on overall %)

| Package | Uncovered | Current % | Nature |
|---------|-----------|-----------|--------|
| manifest/ | 280 | 67.5% | Mostly Renderable/Deprecatable interface impls + Clone. Add tests to existing test files. |
| generate/ | 107 | 68.1% | Pure helpers: html.go (getInnerComment, parseYamlComment), typeResolver.go (findModuleBySpec), classMemberDeclarations.go (isIgnoredMember, isPropertyField), external_types.go (isQuoted, reconstructTemplate), css.go (sortCustomProperty), reexports.go (exportKeySet) |
| validate/ | 76 | 63.2% | types.go getter methods (30), warnings.go rule check functions (18), navigator.go path traversal |
| lsp/ | 76 | 56.8% | registry.go helpers (already exercised cross-package), incremental.go text manipulation, server_context.go URI helpers |
| lsp/methods/textDocument/ | 77 | 31.1% | lifecycle.go dispatch (Tier 2), completion/definition/diagnostics helpers (some already tested) |
| internal/workspace/ | 79 | 42.5% | discovery helpers, config merging, URL manipulation |
| internal/modulegraph/ | 69 | 44.1% | tracking operations, path normalization, metrics |
| serve/ | 78 | 43.0% | build.go helpers (partially tested), filewatcher.go pattern matching, server lifecycle (Tier 2/3) |

### Medium priority

| Package | Uncovered | Current % | Nature |
|---------|-----------|-----------|--------|
| health/ | 35 | 85.9% | Warning rules, display helpers |
| serve/middleware/routes/ | 45 | 38.8% | Route building (Tier 2), remaining pure helpers |
| lsp/document/ | 36 | 3.4% | manager.go lifecycle (Tier 2), position helpers (moved to helpers/) |
| lsp/document/base/ | 29 | 42.6% | BaseDocument methods (partially cross-package tested) |
| export/ | 20 | 69.6% | Template helpers, Export orchestrator |
| mcp/ | 26 | 68.5% | context.go element helpers, cache key generation |
| internal/platform/ | 31 | 51.4% | Filesystem helpers, MapFS operations |

### Lower priority (small packages or Tier 2/3)

| Package | Uncovered | Nature |
|---------|-----------|--------|
| serve/middleware/transform/ | 24 | Cache, engine helpers |
| serve/middleware/importmap/ | 23 | Import map generation (Tier 2) |
| mcp/tools/ | 22 | Tool handlers (Tier 2) |
| internal/treesitter/ | 22 | QueryManager (cross-package tested) |
| mcp/resources/ | 21 | Data source converters |
| lsp/document/html/ | 21 | Document.go internals (Tier 2) |
| lsp/document/typescript/ | 20 | Template extraction (Tier 2) |
| lsp/ephemeral/ | 16 | Ephemeral registry |
| lsp/document/tsx/ | 16 | Completion analysis (Tier 2) |
| generate/demodiscovery/ | 11 | Discovery helpers |
| generate/jsdoc/ | 35 | JSDoc parsing (already 92.4% after recent work) |

## Agents identified these specific functions (from failed worktree runs)

### generate/ (from agent research)
- `html.go`: getInnerComment, isYamlComment, unescapeBackticks, parseYamlComment, dedentYaml
- `typeResolver.go`: findModuleBySpec
- `classMemberDeclarations.go`: isIgnoredMember, isPropertyField, getMemberKindFromCaptures, isStaticToTypeFlag
- `external_types.go`: isQuoted, isNumericLiteral, reconstructTemplate, validateExpressionValues, resolveTextParts
- `generate_helpers.go`: matchesAnyPattern, ColorizeDuration
- `css.go`: sortCustomProperty
- `reexports.go`: exportKeySet, statementTextBefore

### serve/ + mcp/ + export/ + health/ (from agent research)
- `serve/build.go`: rewriteJSONPaths, rewriteBasePath
- `serve/filewatcher.go`: getDefaultIgnorePatterns, shouldIgnore
- `mcp/context.go`: extractGuidelinesFromElement, generateCacheKey, selectBestSchemaVersion, computeCommonPrefixes
- `export/template_helpers.go`: jsKey
- `health/rules.go`: status, proportionalScore, percentage, buildBar, containsWord, generateRecommendations

### internal/workspace/ + internal/modulegraph/ (from agent research)
- Pure functions in workspace: URL manipulation, config merging helpers
- Pure functions in modulegraph: tracking operations, path normalization

### lsp/ (from agent research)
- `incremental.go`: calculateOldLength, lspPositionToTreeSitterPoint, calculateNewEndPoint
- `registry.go`: pathsMatch, normalizePath (already have some tests)
- `document/html/document.go`: CompletionPrefix
- `document/tsx/document.go`: CompletionPrefix
- `document/typescript/document.go`: CompletionPrefix

## Skipped packages (not counting toward target)

- cmd/ -- covered by e2e tests
- internal/languages/* -- Tier 3 grammar registration
- internal/logging, internal/version -- infrastructure
- internal/platform/testutil, lsp/testhelpers -- test helpers
- mcp/types, lsp/types -- type definitions
- mcp/tools/gen-adapters -- code generator
- serve/logger -- terminal output wrapper
- main.go -- entry point
