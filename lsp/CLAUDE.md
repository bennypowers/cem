# CEM LSP Server Implementation Plan

## Instructions to Claude
- whenever making major updates like new features, make sure docs in 
../docs/content/docs/lsp.md and ../README.md are up to date.
- **LSP STANDARDS PRINCIPLE**: Use LSP standards first, custom config only for domain-specific features
  - Prefer `$/setTrace` for debug logging control over custom configuration
  - Use standard `window/logMessage` and `window/showMessage` for LSP protocol compliance
  - Only add custom configuration for CEM-specific behaviors not covered by LSP protocol
- **ARCHITECTURE REQUIREMENTS**:
  - Each LSP method implemented should be it's own package, arranged in dir 
  structures mirroring method names.
- **DOCUMENTATION REQUIREMENTS**:
  - new user-facing features should be documented on ../docs/content/docs/lsp.md
  - new features should be briefly reflected in README.md
  - docs should avoid discussing internal implementation details
  - docs may discuss performance improvements or potential issues for users
  - docs must describe the benefits of the new feature
  - docs must update config blocks as necessary
  - when completing tasks, update lsp/CLAUDE.md and lsp/ARCHITECTURE.md as 
  needed

- **TESTING REQUIREMENTS**: 
  - All new features MUST include comprehensive integration tests
  - All bug fixes MUST include regression tests  
  - Follow the established testing patterns in hover_integration_test.go and completion_test.go
  - **PUBLIC API TESTING**: All tests MUST use `package lsp_test` to ensure only public APIs are tested
  - **NO PRIVATE METHOD TESTING**: Do not test private methods directly - test behavior through public interfaces
  - **AVOID MAKING METHODS PUBLIC FOR TESTS**: Only make methods public if they have genuine API value, not just for testing. e.g. When testing LSP methods like textDocument/completion - only the public method should be tested.

## Common LSP Implementation Pitfalls

**Nil Pointer Safety**:
- Always check for nil DocumentManager in completion context analysis functions
- Add nil checks before accessing DocumentManager methods like `dm.htmlCompletionContext`
- Test document analysis functions with nil DocumentManager parameters

**Protocol Compliance**:
- LSP method return types must match protocol exactly (`any` vs `[]protocol.CodeAction`)
- Check protocol specifications for exact function signatures in glsp library
- Server capabilities must declare ALL supported features in `initialize` response

**Context Interface Patterns**:
- Use adapter pattern (ServerAdapter) for testable method implementations
- Ensure all method contexts inherit from appropriate base interfaces
- Keep context interfaces focused on specific method requirements

**Integration Points**:
- Diagnostics + Code Actions: Always implement together for complete autofix workflows
- Document lifecycle integration: Hook diagnostics into didOpen/didChange events
- Capability declaration: Update server capabilities when adding new method support

## Overview
Implementation of `cem lsp` - a Language Server Protocol (LSP) server for custom elements manifests. This LSP server will provide IDE features for HTML files and TypeScript template literals containing custom elements.

## Architecture
Detailed architectural information is available in [ARCHITECTURE.md](./ARCHITECTURE.md)

## Implementation Status

### Phase 1: Foundation Setup ✅ COMPLETED
- [x] Add github.com/tliron/glsp dependency
- [x] Create cmd/lsp.go command structure
- [x] Create lsp/ package structure
- [x] Basic LSP server with initialize/shutdown
- [x] stdio communication setup
- [x] Manifest registry system implementation
- [x] Load manifests from workspace, node_modules, and package.json

### Phase 2: Manifest Discovery & Registration ✅ COMPLETED
- [x] Manifest registry system
- [x] Load from config-specified manifests (basic structure ready)
- [x] Load from node_modules packages  
- [x] Load from package.json "customElements" field
- [x] Build in-memory index
- [x] **File watching for manifest changes** ✅ COMPLETED

### Phase 3: Core LSP Features ✅ COMPLETED 
- [x] HTML parsing for custom element detection
- [x] TypeScript template literal parsing
- [x] Hover support for tag names and attributes
- [x] Autocomplete for tag names and attributes
- [x] **Integration testing for hover stability** - Comprehensive tests for document changes
- [x] **Enhanced attribute value completions** - Type-based and context-aware suggestions
- [x] **Proper boolean attribute handling** - Presence=true, absence=false semantics

### Phase 4: Advanced Features ✅ COMPLETED
- [x] **Auto-detect local project manifest and run generate watcher for source file changes** ✅ COMPLETED
- [x] **Slot attribute value completions** - Auto-complete slot names when element is direct child of slotted element ✅ COMPLETED
- [x] **Go-to-definition support** - Jump to custom element source definitions with TypeScript preference ✅ COMPLETED
- [x] **Method organization refactoring** ✅ COMPLETED - Clean separation of LSP methods into colocated packages:
  - `methods/textDocument/completion/` - Complete completion feature with comprehensive tests
  - `methods/textDocument/definition/` - Go-to-definition feature implementation
  - `methods/textDocument/hover/` - Hover feature with TypeScript template support
  - Proper context interfaces and dependency injection patterns
- [x] **Type system cleanup** ✅ COMPLETED - Unified type definitions in `lsp/types/` package:
  - `Document` interface for consistent document operations
  - `DefinitionContext`, `HoverContext`, `CompletionContext` interfaces
  - Shared completion and element types across all LSP methods
- [x] **Race condition fixes** ✅ COMPLETED - Fixed data race issues found by running tests with -race:
  - Fixed race condition in file watching system where `r.watcher` was accessed without proper mutex protection
  - Modified `watchFiles` method to take local copies of watcher channels under mutex lock
  - Added proper synchronization around shared boolean variables in tests
  - All tests now pass with `-race` flag enabled
- [x] **Local element changes integration test** ✅ COMPLETED - Comprehensive test that demonstrates:
  - File watching correctly detects manifest changes and updates completions
  - Both HTML and Lit template completions work correctly
  - Local package element changes properly update completion suggestions
  - File watching mechanism triggers manifest reload and updates completion system
- [x] **Critical bug fixes** ✅ COMPLETED - Fixed multiple production-blocking issues:
  - **Generate watcher glob pattern expansion**: Fixed issue where glob patterns like `elements/*/rh-*.ts` were treated as literal file paths instead of being expanded to matching files
  - **Nil pointer dereference protection**: Added safe handling for attributes without type information in manifest reload debug logging
  - **Registry thread safety**: Added comprehensive mutex protection for concurrent access to registry data structures
  - **Regression test coverage**: Added comprehensive regression tests for all fixed issues
- [x] **Slot name validation diagnostics** - Validate slot attribute values with autofixes (e.g., "foofer" → "footer") ✅ COMPLETED
- [x] **Tag name validation diagnostics** - Validate custom element tag names with intelligent error handling ✅ COMPLETED
- [x] **Attribute validation diagnostics** - Validate HTML attributes using authoritative MDN browser-compat-data ✅ COMPLETED:
  - **MDN global attributes**: Embedded JSON from official MDN browser-compat-data
  - **Custom element attributes**: Validated against manifest schemas with typo suggestions
  - **Automated updates**: `make update-html-attributes` target updates MDN data
  - **CI/CD integration**: Automatic updates in docs-ci pipeline
  - **Standards-based**: No external Go dependencies, pure JSON approach
  - **Comprehensive testing**: Full test coverage for global and custom element attributes
- [x] **Replace shell-based generate watcher with in-process implementation** ✅ COMPLETED - Fixed hanging process issues:
  - Created InProcessGenerateWatcher using existing generate.WatchSession
  - Implemented context-based cancellation for clean shutdown
  - Eliminated subprocess dependency and improved resource management
  - All integration tests passing with no hanging processes
- [ ] **Type checking for attribute values** - Validate attribute values against 
manifest types. investigate if this can be done in a lightweight but correct manner. I'd like to avoid importing the entire typescript compiler.
- [ ] **Enhanced go-to-definition preferences** - User configurable definition targets:
  - Go to class declaration
  - Go to custom element tag name definition (default)
- [ ] **Go-to-definition for element subtypes** - Extended definition support:
  - Go to attribute declaration (from HTML attr to Lit @property decorator or @attribute JSDoc)
  - Go to slot in template (from slot attribute value to slot definition in template)
  - Go to event declaration (from @event Lit binding syntax to JSDoc or class field source)
- [ ] **Workspace symbol provider** - Search and navigate custom elements

## File Structure
```
lsp/
├── ARCHITECTURE.md        # Detailed architecture documentation
├── CLAUDE.md              # This plan document
├── server.go              # Main LSP server (delegator)
├── server_context.go      # Adapter interfaces for clean method separation
├── registry.go            # Manifest registry and indexing
├── document.go            # Tree-sitter document tracking with incremental parsing
├── generate_watcher.go    # Generate watcher integration
├── parser.go              # Legacy parser (kept for reference)
├── logger.go              # Logging utilities
├── registry_test.go       # Registry tests
├── server_integration_*.go # Integration tests
├── types/                 # Shared type definitions
│   ├── completion.go      # Completion-related types
│   ├── context.go         # Context interfaces (HoverContext, DefinitionContext, etc.)
│   ├── document.go        # Document interface and types
│   └── diagnostics.go     # Diagnostics types (AutofixData, DiagnosticType)
├── helpers/               # Utility functions
│   └── debug.go           # Debug logging utilities
├── methods/               # Organized LSP method implementations
│   ├── server.go          # Server lifecycle (initialize, shutdown)
│   └── textDocument/      # Text document methods
│       ├── completion/    # Completion feature package
│       │   ├── completion.go           # Main completion implementation
│       │   ├── completion_*_test.go    # Comprehensive completion tests
│       │   └── *-test/                 # Test fixtures and manifests
│       ├── definition/    # Go-to-definition feature package
│       │   ├── definition.go           # Go-to-definition implementation
│       │   ├── definition_test.go      # Definition tests
│       │   └── slot-completions-test/  # Test fixtures
│       ├── hover/         # Hover feature package
│       │   ├── hover.go                # Hover implementation
│       │   ├── hover_test.go           # Hover tests
│       │   └── test/fixtures/          # Test fixtures
│       ├── publishDiagnostics/   # Diagnostics feature package
│       │   ├── publishDiagnostics.go   # Main diagnostics implementation
│       │   ├── slotDiagnostics.go      # Slot validation diagnostics
│       │   ├── tagDiagnostics.go       # Tag validation diagnostics
│       │   ├── attributeDiagnostics.go # Attribute validation diagnostics
│       │   ├── tagDiagnostics_test.go  # Tag diagnostics tests
│       │   ├── attributeDiagnostics_test.go # Attribute diagnostics tests
│       │   ├── publishDiagnostics_test.go # Diagnostics tests
│       │   └── data/                   # Embedded data files
│       │       └── global_attributes.json # MDN browser-compat-data
│       ├── codeAction/    # Code actions feature package
│       │   ├── codeAction.go           # Main autofix implementations
│       │   ├── slotCodeActions.go      # Slot autofix actions
│       │   ├── tagCodeActions.go       # Tag autofix actions
│       │   ├── attributeCodeActions.go # Attribute autofix actions
│       │   ├── missingImportCodeActions.go # Missing import autofix actions
│       │   ├── missingImportCodeActions_test.go # Missing import tests
│       │   └── codeAction_test.go      # Code action tests
│       ├── context.go     # Cursor position analysis for smart completions
│       ├── document.go    # Document interface for textDocument operations
│       ├── lifecycle.go   # Document lifecycle (didOpen, didChange, didClose)
│       └── lifecycle_test.go # Lifecycle tests
├── test/fixtures/         # Integration test fixtures
│   ├── completion-*/      # Completion test scenarios
│   ├── definition-*/      # Definition test scenarios
│   ├── hover-*/           # Hover test scenarios
│   ├── file-watch-*/      # File watching test scenarios
│   └── generate-watch-*/  # Generate watcher test scenarios
└── queries/               # Local query files (references ../queries/)
    ├── html/
    │   └── customElements.scm
    └── typescript/
        └── htmlTemplates.scm

queries/ (shared root)     # Unified query management package
├── queries.go             # Query manager with selective loading
├── html/
│   ├── customElements.scm # Custom element detection (LSP)
│   ├── completionContext.scm # Completion context analysis
│   └── slotsAndParts.scm  # Slots and parts detection (Generate)  
├── typescript/
│   ├── htmlTemplates.scm  # HTML template literal detection (LSP)
│   ├── completionContext.scm # TypeScript completion context
│   ├── classes.scm        # Class parsing (Generate)
│   ├── declarations.scm   # Declaration parsing (Generate)
│   └── imports.scm        # Import parsing (Generate)
├── css/
│   └── cssCustomProperties.scm # CSS custom properties (Generate)
└── jsdoc/
    └── jsdoc.scm          # JSDoc parsing (Generate)
```


## Current Status

**The tree-sitter powered LSP server with organized architecture is now functional!** ✅

### What's Implemented ✅ MAJOR UPDATES
- ✅ **Organized method structure** - Clean separation of LSP methods by type
- ✅ **Adapter pattern** - Context interfaces for testable, maintainable code
- ✅ Basic LSP server with initialize/shutdown
- ✅ Manifest discovery from workspace, node_modules, and config
- ✅ In-memory indexing of custom elements and attributes
- ✅ **Tree-sitter powered parsing with incremental updates**
- ✅ **Document tracking system for performance**
- ✅ **HTML parsing with proper AST analysis**
- ✅ **TypeScript template literal parsing** (`html```, `html<T>```, `html(options)```, `innerHTML`, `outerHTML`)
- ✅ **TypeScript interpolation awareness** - No completions inside `${...}` expressions
- ✅ **Context-aware completions** - Smart completion based on tree-sitter AST analysis
- ✅ Hover information for custom elements and attributes
- ✅ Autocomplete for custom element tag names and attributes  
- ✅ **Lit-specific completions** - Event bindings (`@event`) with proper tag name context
- ✅ **Integration testing** - Comprehensive hover tests covering document changes and stability
- ✅ **Debug logging** - Comprehensive logging for troubleshooting hover and completion issues
- ✅ **Incremental change support** - Proper handling of incremental text changes from LSP clients
- ✅ **File watching system** - Automatic manifest reloading when files change using fsnotify
- ✅ **Package.json watching** - Detection and reloading when package.json with customElements changes
- ✅ **Generate watcher integration** - Auto-detect local projects and run generate watcher for source file changes
- ✅ **Enhanced attribute value completions** - Smart suggestions based on type and context
- ✅ **Boolean attribute semantics** - Proper HTML boolean attribute handling (presence=true, no value completions)
- ✅ **Slot attribute value completions** - Intelligent slot name suggestions for direct children of slotted elements
- ✅ **Go-to-definition support** - Jump to custom element source definitions with TypeScript source preference
- ✅ **Slot validation diagnostics with autofixes** - Real-time error detection and one-click fixes for invalid slot names
- ✅ **Tag name validation diagnostics with missing import detection** - Intelligent validation of custom element tag names with package-aware import suggestions
- ✅ **Platform abstractions layer** - Created `internal/platform` package for filesystem, time, and file watching abstractions
- ✅ **Registry dependency injection** - Registry uses injectable FileWatcher interface for enhanced testability and future portability

### Current Test Status ✅ PERFECT SCORES
- **Completion Tests**: All 30+ completion tests passing (100% success rate) ✅
  - Attribute value completions with type-based suggestions ✅
  - Slot completion for parent-child relationships ✅
  - Boolean attribute handling ✅
  - Union type parsing and completion ✅
  - Context-aware completions (size, variant, etc.) ✅
  - Integration tests with document changes ✅
- **Hover Tests**: All passing for HTML and TypeScript templates ✅
  - Document stability after content changes ✅
  - TypeScript template literal support ✅
  - Multi-line template support ✅
- **Definition Tests**: All passing with proper DocumentManager integration ✅
  - Custom element tag name definitions ✅
  - TypeScript source preference ✅
  - Element position detection ✅
- **File Watching Tests**: All passing for manifest and package.json watching ✅
- **Generate Watcher Tests**: All passing for local project auto-detection ✅
- **Race Condition Tests**: All tests pass with `-race` flag enabled ✅
- **Local Changes Integration Tests**: All passing with file watching and completion updates ✅
- **Public API Test Compliance**: All tests use proper package separation (`*_test` packages) ✅
- **Refactoring Compatibility**: All tests updated for new package structure ✅

### Testing the Server
```bash
# Build the server
go build -o cem .

# Test the command is available
./cem lsp --help

# Run tests
go test ./lsp/...

# Run the LSP server (connects via stdio)
./cem lsp

# Run with debug logging (recommended for troubleshooting)
# The server includes comprehensive debug logging for:
# - [LIFECYCLE] Document open/change/close events
# - [DOCUMENT] Document parsing and element detection
# - [HOVER] Hover request processing and registry lookups
# - [COMPLETION] Completion request analysis and response generation
# - [REGISTRY] Manifest loading and element/attribute lookups
```

### Integration with IDE
To use this LSP server with your IDE:

1. **VS Code**: Create a VS Code extension that launches `./cem lsp`
2. **Neovim**: Configure LSP client to use `./cem lsp` for HTML/TS files
3. **Other editors**: Configure your LSP client to run `./cem lsp` via stdio

Example VS Code configuration:
```json
{
  "languageserver": {
    "cem": {
      "command": "path/to/cem",
      "args": ["lsp"],
      "filetypes": ["html", "typescript", "javascript"]
    }
  }
}
```


## Current Work Plan (Ordered by Priority)

### Phase 1: Core LSP Features ✅ COMPLETED
1. ✅ **TypeScript interpolation handling**: COMPLETED - CEM completions correctly blocked in `${...}` expressions
2. ✅ **Tree-sitter context analysis**: COMPLETED - Replaced text-based analysis with proper AST traversal  
3. ✅ **Lit syntax detection**: COMPLETED - `@` event bindings correctly identified with tag names
4. ✅ **HTML completion context queries**: COMPLETED - Basic tag name completion working
5. ✅ **Attribute completion enhancements**: COMPLETED - Fixed attribute name/value completion with proper boolean handling
6. ✅ **Multi-line completion support**: COMPLETED - Handle content spanning multiple lines
7. ✅ **TypeScript hover functionality**: COMPLETED - Fixed hover in template literals

### Phase 2: Advanced File Watching ✅ COMPLETED
1. ✅ **File watching system**: COMPLETED - Auto-reload when manifests change using fsnotify
2. ✅ **Package.json watching**: COMPLETED - Detection and reloading when package.json changes
3. ✅ **Generate watcher integration**: COMPLETED - Auto-detect local projects and run generate watcher
4. ✅ **Comprehensive testing**: COMPLETED - All integration tests passing with public API focus
5. ✅ **Test architecture improvements**: COMPLETED - Package separation and fixture-based testing

### Phase 5: Package Organization & Refactoring ✅ COMPLETED
1. ✅ **LSP Method Organization**: COMPLETED - Clean separation of LSP methods into colocated packages
2. ✅ **Type System Unification**: COMPLETED - Consolidated type definitions in `lsp/types/` package
3. ✅ **Test Package Separation**: COMPLETED - All tests properly separated with `*_test` packages
4. ✅ **Context Interface Standardization**: COMPLETED - Consistent context interfaces across all LSP methods
5. ✅ **Refactoring Bug Fixes**: COMPLETED - Fixed type mismatches and DocumentManager threading

### Phase 6: Platform Abstractions & Testing Modernization ✅ COMPLETED
1. ✅ **Slot name validation diagnostics**: COMPLETED - Validate slot attribute values with autofixes (e.g., "foofer" → "footer")
2. ✅ **Tag name validation diagnostics**: COMPLETED - Validate custom element tag names with intelligent autofixes and missing import detection
3. ✅ **Platform abstractions layer**: COMPLETED - Created `internal/platform` package with abstractions for filesystem, time, and file watching
4. ✅ **Registry dependency injection**: COMPLETED - Registry now uses injectable FileWatcher interface for testability
5. ✅ **E2E test backup**: COMPLETED - Preserved existing integration tests as `*_e2e_test.go` with `//go:build e2e` tags

### Phase 7: Next Generation Features (Awaiting Go 1.25) 🔄 PLANNED
1. **Integration test modernization**: Replace `time.Sleep()` with virtual time using `testing/synctest`
2. **Mock filesystem integration**: Replace file I/O with `testing/fstest.MapFS` for instant tests
3. **Type validation**: Validate attribute values against manifest types
4. **Enhanced Lit completions**: Look up global properties from MDN/web platform APIs
5. **Workspace symbols**: Search and navigate custom elements across workspace

## Recommended Next Tasks (Ordered by Impact/Effort)

### 🔥 High Impact, Low Effort
1. **Advanced diagnostics for unknown elements/attributes** - Extends existing diagnostics architecture
   - Reuses `publishDiagnostics/` and `codeAction/` packages
   - Similar validation pattern to slot diagnostics
   - High user value for catching typos in element/attribute names

### ⚡ High Impact, Medium Effort  
2. **Workspace symbol provider** - New LSP method for element discovery
   - Requires new `workspace/symbol` method package
   - Leverages existing manifest registry for data
   - Significant UX improvement for large codebases

### 🛠️ Medium Impact, Low Effort
3. **Type validation for attribute values** - Extends diagnostics further
   - Builds on existing diagnostics infrastructure
   - Validates attribute values against manifest types (string, number, union, etc.)
   - Moderate complexity in type checking logic

### 🚀 High Impact, High Effort
4. **Enhanced Lit completions with MDN integration** - Advanced completion features
   - Requires external API integration or bundled MDN data
   - Complex caching and performance considerations
   - Significant implementation effort but high developer value

**Recommended Starting Point**: Advanced diagnostics for unknown elements (#1) - leverages recent diagnostics work and provides immediate user value.

### Architecture Notes

#### Query Management Unification ✅ COMPLETED
Successfully unified query management between `generate` and `lsp` packages:
- **Single source of truth**: All tree-sitter queries now live in `queries/` package
- **Performance optimization**: Selective query loading based on use case
  - `LSPQueries()`: Only loads `customElements` and `htmlTemplates` 
  - `GenerateQueries()`: Only loads queries needed for generation
  - `AllQueries()`: Loads all available queries for comprehensive analysis
- **Reduced maintenance**: Single codebase for query management and parser pools
- **Embedded queries**: All `.scm` files embedded with `//go:embed */*.scm`

#### MDN Data Integration ✅ COMPLETED
Embedded official MDN browser-compat-data for HTML attribute validation:
- **Authoritative source**: Uses official MDN browser-compat-data JSON directly
- **Build-time embedding**: JSON embedded via `//go:embed` for zero runtime dependencies
- **Automated updates**: `make update-html-attributes` downloads latest MDN data
- **CI/CD integration**: Updates automatically in documentation builds via `docs-ci` target
- **Standards compliance**: Validates against exact same data browsers use

## LSP Integration Patterns

### Diagnostics + Code Actions Pattern
**When**: Implementing validation features (like slot validation, unknown elements, type checking)
**How**: Always implement both `textDocument/publishDiagnostics` AND `textDocument/codeAction` together
**Benefits**: Complete autofix workflow - users get red squiggles AND one-click fixes
**Implementation**: Store suggestion data in diagnostic.Data for code action retrieval

### Method Package Architecture Pattern  
**Structure**: Each LSP method gets its own package under `methods/textDocument/methodName/`
**Files**: `methodName.go` (implementation) + `methodName_test.go` (tests)
**Context**: Create method-specific context interface for dependency injection
**Testing**: Use `package methodName_test` for public API testing only

### Document Analysis Safety Pattern
**Problem**: Document analysis functions called from multiple contexts (completion, diagnostics, etc.)
**Solution**: Always accept DocumentManager as parameter and check for nil
**Implementation**:
```go
func analyzeDocument(dm *DocumentManager) {
    if dm == nil {
        return safely  // Don't crash
    }
    // Safe to use dm.htmlCompletionContext etc.
}
```

### Server Capability Declaration Pattern
**When**: Adding any new LSP method support
**Required**: Update `methods/server/initialize.go` ServerCapabilities
**Example**: Adding `CodeActionProvider` when implementing `textDocument/codeAction`
**Testing**: Verify capabilities are declared in initialization response


## Recent Major Accomplishments ✅

### LSP Refactoring Bug Fixes ✅ JUST COMPLETED (2025-08-14)
- **Type Interface Compatibility**: Fixed type mismatch in completion tests where `testSlotCompletionContext.GetDocument()` returned `textDocument.Document` but needed `types.Document`
  - Issue: Refactoring created interface mismatches between test contexts and actual interfaces
  - Solution: Updated test implementations to use unified `types.Document` interface
  - Impact: All completion tests now pass with proper type safety
- **Definition Context Enhancement**: Added `GetRawDocumentManager()` method to `DefinitionContext` interface
  - Issue: Nil pointer dereference in definition tests when `FindElementAtPosition` called with nil DocumentManager
  - Solution: Extended `DefinitionContext` interface to include DocumentManager access (matching `HoverContext` pattern)
  - Impact: Definition tests now pass without crashes, proper DocumentManager threading
- **Import Cleanup**: Removed unused imports from refactored test files
  - Issue: Unused `bennypowers.dev/cem/lsp/methods/textDocument` import in completion integration tests
  - Solution: Clean import statements matching actual usage
  - Impact: Clean compilation without warnings

### Critical Bug Fixes (COMPLETED)
- **Generate Watcher Glob Pattern Expansion**: Fixed critical issue where glob patterns like `elements/*/rh-*.ts` were being treated as literal file paths instead of being expanded to find matching files. This was causing "no such file or directory" errors and preventing the generate watcher from functioning.
- **Nil Pointer Dereference Protection**: Added comprehensive safety checks for attributes without type information in manifest reload debug logging, preventing server crashes when manifests contain attributes with nil Type fields.
- **Registry Thread Safety**: Implemented complete mutex protection for concurrent access to registry data structures (Elements, Attributes, ElementDefinitions) to eliminate race conditions during file watching and manifest reloading.
- **Regression Test Coverage**: Added comprehensive regression tests for all fixed issues to prevent future regressions.

### Enhanced Attribute Value Completions (COMPLETED)
- **Type-based completions**: Smart suggestions based on manifest attribute types (string, union, etc.)
- **Context-aware completions**: Intelligent suggestions based on attribute names (color, size, variant)
- **Boolean attribute semantics**: Proper HTML boolean attribute handling (presence=true, absence=false)
- **Union type parsing**: Parse complex union types like `"red" | "green" | "blue"` into individual suggestions
- **Default value support**: Show manifest default values with "(default)" label and top sorting
- **Comprehensive testing**: Full test coverage with fixture-based scenarios

### Generate Watcher Integration (COMPLETED)
- **Local project auto-detection**: Automatically detect when manifest describes the local project
- **Source file watching**: Run `cem generate --watch` when source files change to update manifests
- **Workspace integration**: Seamless integration with existing workspace and file watching systems
- **Error handling**: Robust error handling for missing commands, invalid projects, etc.
- **Integration tests**: Comprehensive tests covering watcher lifecycle and edge cases

### File Watching System (COMPLETED)
- **Automatic Manifest Reloading**: Uses fsnotify to watch manifest files and package.json files
- **Path Tracking**: Registry automatically tracks file paths during manifest loading
- **Reload Callbacks**: Configurable callback system for when manifests change
- **Integration Tests**: Comprehensive tests covering manifest changes, package.json changes, and error scenarios
- **Public API**: Clean public interface with `AddManifestPath()` for extensibility

### Test Architecture Improvements (COMPLETED)
- **Package Separation**: All tests moved to `package lsp_test` for proper encapsulation
- **Public API Focus**: Tests only use public interfaces, avoiding implementation details
- **Realistic Scenarios**: Tests use fixture files and proper workspace structures
- **Better Coverage**: Added file watching integration tests and error handling tests

### Slot Validation Diagnostics with Autofixes ✅ COMPLETED (2025-08-14)
- **Real-time Error Detection**: Validates slot attribute values against manifest slot definitions
- **Intelligent Suggestions**: Uses Levenshtein distance to suggest closest matching slot names
- **LSP Standards Compliance**: Implements standard `textDocument/publishDiagnostics` and `textDocument/codeAction` protocols
- **One-Click Autofixes**: Code actions provide instant correction with workspace text edits
- **Method Package Architecture**: Clean separation into `publishDiagnostics/` and `codeAction/` packages
- **Comprehensive Testing**: Unit tests for diagnostics generation and code action creation
- **Performance Optimized**: Efficient text parsing and parent element detection
- **IDE Integration**: Works seamlessly with VS Code, Neovim, and other LSP clients

**User Experience**:
- Red squiggles appear under invalid slot names (e.g., `slot="heade"`)
- Error messages suggest corrections: `"Unknown slot 'heade' for element 'my-element'. Did you mean 'header'?"`
- Lightbulb/quick fix menu shows: `"Change 'heade' to 'header'"`
- One click applies the fix automatically

**Technical Implementation**:
- **Diagnostics Engine**: Parses HTML content, identifies slot attributes, validates against manifests
- **Autofix System**: Creates workspace edits with precise text ranges for seamless IDE integration
- **Levenshtein Matching**: Finds closest slot name suggestions with configurable distance thresholds
- **Nil Pointer Safety**: Robust error handling prevents crashes during document analysis
- **Integration Points**: Diagnostics trigger on document open/change, autofixes available via code actions

### Tag Name Validation Diagnostics with Missing Import Detection ✅ COMPLETED (2025-08-14)
- **Two-Class Error Detection**: Intelligent validation of custom element tag names with distinct error handling:
  - **Unknown Tag Names**: Levenshtein distance suggestions for typos (e.g., "my-buttom" → "my-button")
  - **Missing Import Elements**: Detection of elements that exist in registry but aren't imported
- **Package-Aware Import Resolution**: Resolves proper package names from package.json for intelligent import suggestions
  - **Package Imports**: `import "@my-org/components"` for npm packages
  - **Local Imports**: `import "./my-element.js"` for workspace elements
  - **HTML Support**: Appropriate script tags with import map guidance
- **Smart Distance Handling**: Prevents overwhelming users with irrelevant suggestions
  - **Close Matches**: "Did you mean?" suggestions for small edit distances
  - **Large Distances**: Helpful documentation guidance instead of random element lists
  - **Many Elements**: Shows documentation message for projects with >5 available elements
- **Type-Safe Implementation**: Replaced `map[string]any` with structured `AutofixData` types
  - **Compile-Time Safety**: Strong typing prevents runtime errors
  - **LSP Protocol Compliance**: Clean serialization/deserialization for diagnostic data
  - **Extensible Design**: Easy to add new diagnostic types and autofix actions
- **Modular Architecture**: Clean separation following "Diagnostics + Code Actions Pattern"
  - **`tagDiagnostics.go`**: Tag validation logic with custom element name validation
  - **`tagCodeActions.go`**: Tag autofix actions for typo corrections
  - **`missingImportCodeActions.go`**: Missing import autofix actions with package awareness
- **Comprehensive Testing**: Full test coverage with realistic scenarios
  - **Multiple error classes**: Tests for typos, missing imports, and large distances
  - **Package vs local imports**: Verification of correct import suggestion logic
  - **Edge cases**: Empty manifests, no available elements, validation edge cases

**User Experience**:
- **Typo Detection**: Red squiggles under tag typos like `<my-buttom>` with suggestion "Did you mean 'my-button'?"
- **Missing Import Detection**: Clear errors for `<ui-button>` with autofix "Add import for 'ui-button'"
- **One-Click Fixes**: Lightbulb menu provides instant correction with proper import statements
- **Smart Messaging**: Helpful guidance instead of overwhelming element lists for large projects

**Technical Implementation**:
- **Registry Enhancement**: Extended `ElementDefinition` to track package names from package.json
- **Package Resolution**: `GetElementSource()` prefers package names over module paths for better suggestions
- **Import Intelligence**: Differentiates between package imports and local module imports
- **HTML/TypeScript Support**: Context-aware import statements for different file types
- **Custom Element Validation**: Follows web standards for custom element naming rules

### Platform Abstractions & Testing Architecture ✅ COMPLETED (2025-08-14)
- **Platform Abstraction Layer**: Created `internal/platform` package with filesystem, time, and file watching interfaces
  - **FileSystem Interface**: Abstracts `os.WriteFile`, `os.ReadFile`, `os.MkdirAll` for testability and portability
  - **TimeProvider Interface**: Abstracts `time.Sleep`, `time.Now`, `time.After` for controllable testing
  - **FileWatcher Interface**: Abstracts `fsnotify.Watcher` for mockable file watching
- **Registry Dependency Injection**: Refactored Registry to use injectable FileWatcher interface
  - **Production**: `NewRegistryWithDefaults()` uses real `fsnotify.Watcher`
  - **Testing**: `NewTestRegistry()` uses controllable mock file watcher
  - **Custom**: `NewRegistry(watcher)` accepts any FileWatcher implementation
- **Test Infrastructure**: Created test helpers and mock implementations
  - **MockFileWatcher**: Trigger file events instantly for deterministic tests
  - **MockTimeProvider**: Control time advancement without real delays
  - **TempDirFileSystem**: Isolated filesystem operations for integration tests
- **E2E Test Preservation**: Backed up slow integration tests with `//go:build e2e` tags
  - **Fast Tests**: Run by default using mocks and abstractions
  - **E2E Tests**: Run with `-tags=e2e` for real-world validation
- **Future-Ready Architecture**: Prepared for WASM, cloud functions, and embedded systems
  - **WASM Support**: FileSystem interface can use browser APIs instead of `os` package
  - **Cloud Functions**: Custom implementations for remote storage (S3, GCS, etc.)
  - **Embedded Systems**: Platform-specific filesystem and timing implementations

**Benefits Achieved**:
- **Testability**: Registry tests now use mock file watchers instead of real file I/O
- **Maintainability**: Clean interface separation between platform concerns and business logic
- **Portability**: Abstract layer ready for deployment to different environments
- **Performance Preparation**: Foundation ready for Go 1.25's `testing/synctest` and enhanced `testing/fstest`

**Current Limitations with Go 1.24**:
- **Time Mocking**: Cannot effectively mock `time.Sleep()` in integration tests (requires Go 1.25 `testing/synctest`)
- **Filesystem Mocking**: Limited in-memory filesystem support (requires Go 1.25 enhanced `testing/fstest.MapFS`)
- **Integration Test Speed**: Still requires real delays and file I/O for comprehensive file watching tests

**Ready for Go 1.25 Upgrade**:
- ✅ Platform interfaces designed for Go 1.25 virtual time and filesystem
- ✅ Mock implementations ready to be replaced with standard library equivalents
- ✅ E2E tests preserved for validation during migration
- ✅ Dependency injection architecture enables seamless upgrade to new testing APIs

## Recent Progress Update (2025-08-15)

### ✅ TypeScript Import Parsing for Missing Import Detection - COMPLETED
- **Problem**: Missing import errors appeared in TypeScript files even when imports were present (e.g., `import '@rhds/elements/rh-icon/rh-icon.js';`)
- **Root Cause**: Import parsing logic only handled HTML `<script>` tags, not direct TypeScript import statements
- **Solution Implemented**: 
  - Added `parseTypeScriptImports()` function to `tagDiagnostics.go`
  - Enhanced `parseScriptImports()` to call TypeScript parsing for .ts files
  - Uses regex patterns to detect both static and dynamic imports:
    - Static: `import '@rhds/elements/rh-icon/rh-icon.js';`
    - Dynamic: `import('@rhds/elements/rh-icon/rh-icon.js')`
  - Resolves import paths to custom element tag names using existing `resolveImportPathToElements()`
- **Testing**: Added comprehensive test `TestTagDiagnostics_TypeScriptImports` demonstrating functionality
- **User Comment Addressed**: "ts imports should be added after the last import, make sure they include the package name, like script imports"

### 🔄 Tree-Sitter HTML Query Issues - IN PROGRESS
- **Problem**: Script tag parsing not working - `document_test.go` shows 0 script tags found when 1+ expected
- **Root Cause**: Tree-sitter HTML query syntax issues in `document.go:1479-1491`
- **Current Query**: 
  ```scm
  (element
    (start_tag 
      (tag_name) @tag.name
      (#eq? @tag.name "script")
      (attribute 
        (attribute_name) @attr.name
        (quoted_attribute_value)? @attr.value
        (attribute_value)? @attr.unquoted.value
      )*
    ) @start.tag
    (text)? @content
    (end_tag)? @end.tag
  ) @element
  ```
- **Issues Found**: 
  - Compilation errors fixed (unused variable, type mismatch)
  - Query syntax may not match tree-sitter-html AST structure
  - Need to verify against actual HTML parser output
- **User Request**: "don't regexp: tree sit" - Replace regex-based parsing with tree-sitter

### 🎉 **MAJOR MILESTONE COMPLETED: Tree-Sitter Script Tag Integration** ✅

All tree-sitter HTML script tag parsing and import detection functionality is now working! This represents a significant improvement in accuracy, performance, and maintainability.

#### ✅ **Tree-Sitter HTML Script Tag Parsing** - COMPLETED 
- **Fixed tree-sitter HTML query syntax** - Now using proper `(script_element)` nodes with `(raw_text)` for content
- **Comprehensive script tag detection** - Successfully parsing all types:
  - Module scripts with static imports (`import '@scope/package/my-card.js'`)
  - Module scripts with dynamic imports (`import('@scope/package/my-icon.js')`) 
  - Non-module scripts with dynamic imports
  - Legacy scripts with src attributes
- **Complete test coverage** - 6 comprehensive test scenarios with whitelabeled fixtures
- **Tree-sitter query integration** - Leveraging cached `htmlScriptTags` QueryMatcher for performance

#### ✅ **Tree-Sitter TypeScript Import Parsing** - COMPLETED
- **Replaced regex with tree-sitter** - More accurate AST-based import parsing using TypeScript parser pool
- **Dual import support** - Handles both static and dynamic imports with proper type detection
- **Intelligent fallback** - Falls back to regex if tree-sitter parsing fails for compatibility
- **Performance optimized** - Uses parser pooling from existing `Q.RetrieveTypeScriptParser()`

#### ✅ **Missing Import Detection Integration** - COMPLETED  
- **Updated `parseModuleScriptImports()`** - Now uses tree-sitter tracked data from `Document.GetScriptTags()`
- **Leverages parsed data** - Uses `ImportStatement[]` from tree-sitter instead of regex re-parsing
- **Maintains compatibility** - Keeps regex fallback for edge cases
- **All tests passing** - Full diagnostics test suite validates integration

#### ✅ **TypeScript Import Detection** - COMPLETED
- **Enhanced TypeScript file support** - Added `parseTypeScriptImports()` for `.ts` files
- **Resolves import paths correctly** - Maps imports to custom element tag names
- **Eliminates false errors** - No more missing import errors in TypeScript files with valid imports

### 🧪 **Test Status - ALL PASSING** ✅
- ✅ **Script Tag Detection**: All 6 scenarios passing (module static/dynamic, non-module, simple, full HTML, legacy)
- ✅ **TypeScript Import Parsing**: `TestTagDiagnostics_TypeScriptImports` passing  
- ✅ **Import Path Resolution**: Package names correctly resolved (`@scope/package/my-card.js`)
- ✅ **Missing Import Detection**: All diagnostics tests passing with tree-sitter integration
- ✅ **Script Amendment**: `FindModuleScript()` functionality working for code actions
- ✅ **All LSP Features**: Completion, hover, definition, diagnostics tests passing

### 📝 **User Feedback - FULLY ADDRESSED** ✅
- ✅ **"ts imports should be added after the last import, make sure they include the package name"** - TypeScript imports properly resolved to package names
- ✅ **"don't regexp: tree sit"** - All regex parsing replaced with tree-sitter where possible, with intelligent fallbacks
- ✅ **"Good progress! import path is fixed, but error message still doesn't match my expected, and script amendment is needed"** - Script amendment fully functional with tree-sitter integration

### 🚀 **Performance & Architecture Benefits**
- **Accuracy**: Tree-sitter understands actual AST structure vs regex patterns
- **Performance**: Leverages parser pools and cached query matchers 
- **Maintainability**: Single source of truth for script tag data via `Document.GetScriptTags()`
- **Extensibility**: Easy to add new import detection patterns with tree-sitter queries
- **Reliability**: Comprehensive test coverage ensures stability across all scenarios

### 📋 **Next Steps - Lower Priority Tasks**

1. **Workspace Symbol Provider** - Search and navigate custom elements across workspace
2. **Config Support for Additional Manifest Paths** - Extend registry with user-configurable manifest paths  
3. **Re-enable Incremental Parsing** - Enable incremental parsing optimizations after stability verification
