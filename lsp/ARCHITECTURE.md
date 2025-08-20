# CEM LSP Server Architecture

## Overview

The CEM LSP server provides Language Server Protocol support for custom elements in HTML files and TypeScript template literals. It uses tree-sitter for parsing and leverages existing custom-elements manifest data for intelligent features.

## Core Components

### Server (`server.go`)
- **Role**: LSP protocol handler and coordinator
- **Responsibilities**: Delegates LSP requests to organized method handlers
- **Dependencies**: Registry, DocumentManager, WorkspaceContext
- **Pattern**: Thin delegator layer with adapter pattern for clean interfaces

### Registry (`registry.go`) 
- **Role**: Custom elements manifest indexing and file watching
- **Data Sources**: Workspace manifests, node_modules packages, config-specified files
- **Indexing**: Tag name → element definition, element → attributes mapping
- **Performance**: In-memory indexes for O(1) lookups
- **File Watching**: Automatic manifest reload using fsnotify when files change
- **Path Tracking**: Monitors manifest files and package.json files for changes
- **Thread Safety**: Comprehensive mutex protection for all registry data structures
  - Main data mutex (`mu`) protects Elements, Attributes, ElementDefinitions maps
  - File watcher mutex (`watcherMu`) protects file watching operations
  - Generate watcher mutex (`generateMu`) protects generate watcher lifecycle
  - All read/write operations properly synchronized to prevent race conditions

### Document Manager (`document.go`)
- **Role**: Document lifecycle and tree-sitter parsing
- **Features**: Incremental parsing, parser pooling, query caching
- **Parsing**: HTML files directly, TypeScript template literals extracted and parsed as HTML
- **Cache**: Pre-parsed tree-sitter queries for performance

### Method Organization (`methods/`)
- **Structure**: Organized by LSP method type with colocated packages
- **Current Organization**:
  - `methods/server.go` - Server lifecycle methods (initialize, shutdown)
  - `methods/textDocument/` - All text document methods in separate packages:
    - `completion/` - Complete completion feature with tests and fixtures
    - `definition/` - Go-to-definition implementation with tests
    - `hover/` - Hover functionality with TypeScript template support
    - `lifecycle.go` - Document lifecycle (didOpen, didChange, didClose)
    - `context.go` - Cursor position analysis utilities
- **Pattern**: Context interfaces for dependency injection and clean testing
- **Benefits**: Clean separation, testable units, maintainable code, colocated tests

## Parsing Strategy

### Tree-sitter Integration
- **Query System**: Pre-compiled tree-sitter queries cached at startup via shared `queries/` package
- **Selective Loading**: Only loads queries needed for LSP operations (`customElements`, `htmlTemplates`)
- **HTML Parsing**: Direct parsing of HTML files for custom elements
- **TypeScript Parsing**: Extract HTML template literals then parse as HTML
- **Performance**: Shared parser pools between `generate` and `lsp` packages

### Template Literal Support
Detects and parses:
- `html`backtick`` tagged templates
- `html<Generic>`backtick`` tagged templates with generics
- `element.innerHTML = `backtick`` assignments
- `element.outerHTML = `backtick`` assignments

## Data Flow

1. **Initialization**: Load manifests → build indexes → cache queries → start file watching
2. **Document Open**: Parse document → extract custom elements → cache AST
3. **LSP Request**: Analyze cursor position → query indexes → return results
4. **Document Change**: Incremental reparse → update cached data
5. **Manifest Change**: File watcher detects change → reload manifest → rebuild indexes → notify clients

## Performance Characteristics

- **Query Parsing**: Once at startup (expensive operation avoided during requests)
- **Document Parsing**: Incremental updates using tree-sitter
- **Manifest Lookup**: O(1) hash table lookups
- **Memory Usage**: Cached ASTs and indexes, managed cleanup

## Interface Design

Uses adapter pattern to separate concerns with unified type system:
- **HoverContext**: Element and attribute lookup capabilities + raw DocumentManager access
- **CompletionContext**: Tag name enumeration, element details, and slot information
- **DefinitionContext**: Element definition lookup + raw DocumentManager access for position analysis
- **LifecycleContext**: Document management operations
- **ServerContext**: Resource cleanup and workspace access

**Type Unification in `lsp/types/`**:
- **Document**: Unified document interface for all LSP methods
- **CustomElementMatch**: Shared element match type
- **AttributeMatch**: Shared attribute match type
- **ElementDefinition**: Common element definition interface

This allows method implementations to be pure functions with clear dependencies and consistent types.

## Limitations

- **Template Positioning**: Range calculation in template literals is simplified
- ✅ **Incremental Parsing**: Currently does full reparse on document changes (FIXED - now handles incremental text changes properly)
- **Context Analysis**: Tree-sitter based template detection could be more sophisticated

## Testing Architecture

### Test Package Separation
- **Public API Testing**: All tests use `package lsp_test` to ensure only public APIs are tested
- **Implementation Detail Avoidance**: Private methods are not tested directly
- **Realistic Test Scenarios**: Tests focus on behavior through public interfaces

### Integration Test Strategy
- **Test Fixtures**: Realistic custom elements manifests with complete element definitions
- **Document Lifecycle Testing**: Comprehensive testing of document changes and hover stability
- **Position-based Testing**: Verification of cursor position analysis for elements and attributes
- **Adapter Pattern Testing**: Custom test contexts implementing LSP interfaces for isolated testing
- **File Watching Testing**: End-to-end tests for manifest file change detection and reloading

### Test Files Structure
```
cmd/fixture/lsp-hover-test/           # Integration test fixtures
├── package.json                     # Package with customElements reference
├── custom-elements.json             # Complete manifest with test elements
└── index.html                       # HTML file with custom elements

cmd/fixture/file-watch-test/          # File watching test fixtures
├── package.json                     # Package.json with customElements field
├── initial-manifest.json            # Initial manifest state for testing
└── updated-manifest.json            # Updated manifest for change detection

lsp/*_test.go                         # All tests use package lsp_test
├── file_watch_integration_test.go    # Comprehensive file watching tests
├── hover_integration_test.go         # Document change hover stability tests
├── local_changes_integration_test.go # Local element changes and completion updates
├── completion_test.go                # Completion context and integration tests
├── registry_unit_test.go             # Registry behavior tests
└── context_analysis_test.go          # Context analysis behavior tests
```

### Test Coverage
- ✅ **Hover Stability**: Tests that hover continues working after document edits
- ✅ **Document Management**: Tests document opening, updating, and content tracking
- ✅ **Registry Integration**: Tests manifest loading and element registration
- ✅ **Position Analysis**: Tests cursor position to element/attribute mapping
- ✅ **Interface Adapters**: Tests that different document interfaces work correctly
- ✅ **File Watching**: Tests manifest file change detection, reloading, and package.json watching
- ✅ **Race Condition Safety**: All tests pass with `-race` flag, ensuring thread safety
- ✅ **Local Element Changes**: Integration test demonstrating completion updates when local manifests change
- ✅ **Public API Compliance**: All tests use public APIs only, ensuring proper encapsulation
- ✅ **Refactoring Compatibility**: All tests updated for new package structure and type system
- ✅ **Method Package Testing**: Each feature package has comprehensive tests (completion/, definition/, hover/)
- ✅ **Context Interface Testing**: Thorough testing of all context interfaces and their implementations

### Debug Logging
- ✅ **Comprehensive Logging**: Added detailed debug logging throughout the system
  - **[LIFECYCLE]**: Document lifecycle events (open, change, close)
  - **[DOCUMENT]**: Document parsing, element detection, and position analysis
  - **[HOVER]**: Hover request processing and registry lookups
  - **[COMPLETION]**: Completion context analysis and item generation
  - **[REGISTRY]**: Manifest loading and element/attribute lookups
- ✅ **Troubleshooting Support**: Logs help identify issues with:
  - Document parsing after changes
  - Element position detection
  - Registry element/attribute lookups
  - Completion context detection
- ✅ **Incremental Change Handling**: Proper support for incremental text changes
  - **[LIFECYCLE]**: Detailed logging of change events and content synchronization
  - **Range-based changes**: Handles both full document and incremental range-based updates
  - **Multi-line changes**: Supports complex text edits across multiple lines
  - **Bounds checking**: Safe handling of out-of-bounds edit ranges

## Recent Critical Improvements

### Production Stability Fixes ✅ COMPLETED
- **Glob Pattern Expansion**: Fixed generate watcher glob pattern handling in `generate/generate.go`
  - Issue: Patterns like `elements/*/rh-*.ts` were treated as literal file paths
  - Solution: Added proper glob expansion using `ctx.Glob()` before processing
  - Impact: Generate watcher now correctly processes source file changes
- **Nil Pointer Safety**: Added comprehensive null checks in `lsp/server.go`
  - Issue: Debug logging crashed when attributes had nil Type fields
  - Solution: Safe type access with fallback empty strings
  - Impact: Server no longer crashes on manifests with incomplete type information
- **Thread Safety**: Complete registry synchronization in `lsp/registry.go`
  - Issue: Race conditions during concurrent manifest loading and file watching
  - Solution: Added main data mutex (`mu`) protecting all registry maps
  - Impact: All tests pass with `-race` flag, eliminates concurrency bugs

### Regression Test Coverage ✅ COMPLETED
- **Generate Glob Test**: `generate/glob_expansion_regression_test.go`
- **Nil Type Test**: `lsp/nil_type_regression_test.go`
- **Race Condition Prevention**: All fixes include comprehensive test coverage

## Recent Architectural Achievements ✅

- ✅ **Query Unification**: Share tree-sitter infrastructure between `generate` and `lsp` (COMPLETED)
- ✅ **Advanced Incremental Parsing**: Proper text synchronization and edit handling (COMPLETED)  
- ✅ **File Watching System**: Automatic manifest reloading on file changes (COMPLETED)
- ✅ **Race Condition Safety**: Thread-safe file watching and registry operations (COMPLETED)
- ✅ **Local Element Integration**: Comprehensive testing of local package changes updating completions (COMPLETED)
- ✅ **Source File Watching**: Auto-regenerate manifests when source files change (COMPLETED)
- ✅ **Go-to-Definition**: Jump to custom element source definitions (COMPLETED)
- ✅ **Production Stability**: Critical bug fixes for glob patterns, nil safety, and thread safety (COMPLETED)
- ✅ **Method Package Organization**: Clean separation of LSP methods into colocated packages (COMPLETED)
- ✅ **Type System Unification**: Consolidated type definitions in `lsp/types/` package (COMPLETED)

## Future Architecture Goals

### Immediate Improvements (High Priority)
- **Method Package Consistency**: Move server lifecycle methods (`initialize`, `shutdown`) to `methods/server/` package for consistency
- **Helper Reorganization**: Move logging utilities from `logger.go` to `helpers/` package
- **LSP-Compliant Logging**: Align logging with LSP standards using `logMessage`/`showMessage` for proper IDE integration
- **Workspace Method Organization**: Create `methods/workspace/` package for workspace-related LSP methods

### Medium-Term Goals
- **Context-Aware Analysis**: Smarter completion based on cursor position
- **Diagnostic System**: Real-time validation with proper error ranges
- **Enhanced Template Support**: Better TypeScript template literal hover support
- **Configuration Management**: Centralized configuration handling with LSP `workspace/configuration` support