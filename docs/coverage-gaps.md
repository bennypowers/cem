# Coverage Gaps Tracker
Describes work to do, not work done. Edit this file as you go, updating scores and _removing_ covered sections.

Target: 80% overall. Current: 61.4% unit / 68.2% combined.

## In-package tests

Unexported functions requiring in-package tests:

| Package    | File                            | Reason                                                                                                                         |
| ---------- | ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| generate/  | html_helpers_test.go            | unexported: dedentYaml, getInnerComment, etc.                                                                                  |
| generate/  | classMemberDeclarations_test.go | unexported: isIgnoredMember, isPropertyField, etc.                                                                             |
| generate/  | external_types_helpers_test.go  | unexported: isQuoted, isNumericLiteral, etc.                                                                                   |
| generate/  | reexports_helpers_test.go       | unexported: exportKeySet, statementTextBefore                                                                                  |
| generate/  | typeResolver_test.go            | unexported: findModuleBySpec                                                                                                   |
| manifest/  | clone_test.go                   | Clone on value types (in-package for field access)                                                                             |
| manifest/  | deprecatable_test.go            | IsDeprecated nil safety                                                                                                        |
| validate/  | validate_helpers_test.go        | unexported: isSemverLessThan, filterWarningsByConfig                                                                           |
| workspace/ | helpers_test.go                 | unexported: parseNpmSpecifier, isGlobPattern                                                                                   |
| mcp/       | context_helpers_test.go         | unexported methods: computeCommonPrefixes, selectBestSchemaVersion; standalone: extractGuidelinesFromElement, generateCacheKey |

## Remaining work by package

### High priority (biggest impact on overall %)

| Package                   | Current % | Nature                                                                                                                                                          |
| ------------------------- | --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| internal/workspace/       | 50.1%     | context.go (GetWorkspaceContext needs cobra.Command), findProjectRoot (needs os.Getwd), OutputWriter (writes to disk)                                           |
| serve/                    | 44.3%     | Build() orchestrator (needs full server), filewatcher.go (shouldIgnore, WatchPaths, processEvents). build.go refactored to use s.fs; helpers tested with MapFS. |
| internal/modulegraph/     | 64.8%     | BuildFromWorkspace (needs real FS), resolveImportPathToFilesRelativeTo, processFileWithDependencies                                                             |
| internal/platform/        | 56.4%     | OSFileSystem methods (needs real FS), FSNotifyFileWatcher (needs real watchers)                                                                                 |
| lsp/                      | 58.1%     | registry pathsMatch/normalizePath (unexported on receiver), server_context.go URI helpers                                                                       |
| lsp/methods/textDocument/ | 82.6%     | Remaining: completion/definition/diagnostics subpackage handlers                                                                                               |

### Medium priority

| Package                  | Current % | Nature                                                                                                      |
| ------------------------ | --------- | ----------------------------------------------------------------------------------------------------------- |
| generate/                | 68.9%     | Tier 2 functions remaining                                                                                  |
| mcp/                     | 74.1%     | Remaining: LoadManifests, convertElement, LSP registry integration                                          |
| export/                  | 84.5%     | Export() refactored to accept FS interface; tested with MapFS. Remaining: template_helpers funcMap coverage |
| validate/                | 73.1%     | display.go (Tier 3 terminal output), remaining rule branches                                                |
| manifest/                | 74.8%     | inheritance merge description dedup branches, Renderable Children() population                              |
| health/                  | 88.5%     | display.go PrintHealthResult (Tier 3, terminal output)                                                      |
| serve/middleware/routes/ | 46.0%     | Route serving handlers (need full middleware chain), listing renderers (need TemplateRegistry)               |
| lsp/document/            | 57.2%     | incremental.go wrappers (0%, thin delegation to lsp/ level), remaining UpdateDocumentWithChanges branches   |
| lsp/document/base/       | 75.9%     | Remaining: handler delegation methods (FindElementAtPosition, FindCustomElements with real handlers)        |

### Lower priority (small packages or Tier 2/3)

| Package                     | Nature                              |
| --------------------------- | ----------------------------------- |
| serve/middleware/transform/ | Cache, engine helpers               |
| serve/middleware/importmap/ | Import map generation (Tier 2)      |
| mcp/tools/                  | Tool handlers (Tier 2)              |
| internal/treesitter/        | QueryManager (cross-package tested) |
| mcp/resources/              | Data source converters              |
| lsp/document/html/          | Document.go internals (Tier 2)      |
| lsp/document/typescript/    | Template extraction (Tier 2)        |
| lsp/ephemeral/              | Ephemeral registry                  |
| lsp/document/tsx/           | Completion analysis (Tier 2)        |
| generate/demodiscovery/     | Discovery helpers                   |

## Skipped packages (not counting toward target)

- cmd/ -- covered by e2e tests
- internal/languages/* -- Tier 3 grammar registration
- internal/logging, internal/version -- infrastructure
- internal/platform/testutil, lsp/testhelpers -- test helpers
- mcp/types, lsp/types -- type definitions
- mcp/tools/gen-adapters -- code generator
- serve/logger -- terminal output wrapper
- main.go -- entry point
