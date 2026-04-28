# CEM Architecture

## Design Principles

1. **Performance First**
   - Fast builds, responsive watch mode, efficient resource pooling
   - Incremental rebuilds, O(1) lookups, parser and query caching

2. **"Good Enough" Idiomatic Support**
   - Focus on common, idiomatic TypeScript/LitElement patterns
   - Deep support for decorator-based components and standard practices
   - Pragmatic approach: cover 95% of real-world usage efficiently
   - Accept limitations rather than slow, comprehensive edge-case handling

3. **Serializable Configuration**
   - Pure JSON/YAML configuration without JavaScript plugins
   - Predictable, reproducible builds across environments
   - Easy integration with CI/CD pipelines and toolchains
   - Configuration as data, not code

4. **Real-World Applicability**
   - All-in-one solution for typical web component development workflows
   - Generate, validate, and query manifests in a single tool
   - Built for practical project needs, not academic completeness
   - Acknowledge that CEM may not be the right tool for every project

## Overview

### Commands
1. **`cem generate`** - Analyzes source code and generates manifest files
2. **`cem list`** - Queries manifest files for component information  
3. **`cem validate`** - Validates manifest files against schemas and best practices

### Key Technologies
- **Language**: Go 1.25
- **Parsing**: Tree-sitter for TypeScript, JavaScript, CSS, HTML, and JSDoc
- **CLI Framework**: Cobra with Viper for configuration
- **Architecture**: Modular package structure with clear separation of concerns

## Project Structure & Package Organization

### Core Packages

#### `/cmd` - Command Line Interface
- **Purpose**: CLI command definitions and entry points
- **Key Files**: `generate.go`, `list.go`, `validate.go`, `root.go`
- **Pattern**: Cobra command structure with configuration management

#### `/generate` - Manifest Generation Engine
- **Purpose**: Core manifest generation logic with incremental rebuild capabilities
- **Architecture**: Session-based processing with dependency tracking
- **Key Components**: GenerateSession, FileDependencyTracker, ModuleBatchProcessor
- **Performance**: 68% improvement with incremental rebuilds (1.1s → 350ms)

#### `/manifest` - Data Model & Serialization
- **Purpose**: Go structs representing Custom Elements Manifest schema
- **Features**: JSON marshaling/unmarshaling, table rendering, validation helpers
- **Schema Support**: CEM 2.1.0 with backward compatibility

#### `/validate` - Schema Validation & Linting
- **Purpose**: Manifest validation against schemas and best practices
- **Features**: Schema validation, intelligent warnings, deprecation detection
- **Schemas**: Support for CEM versions 0.1.0 through 2.1.1-speculative

#### `/list` - Query & Display Engine
- **Purpose**: Fast querying and flexible display of manifest data
- **Features**: Column filtering, table rendering, tree views
- **Output**: Markdown tables, structured data, custom formatting

#### `/lsp` - Language Server Protocol
- **Purpose**: LSP server for editor integration (completions, diagnostics, definitions, references)
- **Architecture**: Method-per-package pattern under `methods/textDocument/`
- **Document System**: BaseDocument embedding with language-specific document types (HTML, TSX, TypeScript)
- **Key Patterns**: Tree-sitter tree ref-counting, UTF-16 offset conversion, parser pools

#### `/mcp` - Model Context Protocol Server
- **Purpose**: MCP server exposing manifest data to AI coding agents
- **Features**: Resources (element listings, details) and tools (generate HTML, validate HTML)

#### `/serve` - Development Server
- **Purpose**: Local dev server with live reload for web component development
- **Features**: Import rewriting, TypeScript transform, frontend test runner integration

### Internal Packages (`/internal`)

#### `internal/workspace` - Workspace Abstraction
- **Purpose**: Unified interface for local and remote package access
- **Implementations**: FileSystem and Remote workspace contexts
- **Key Features**: Path resolution, file operations, configuration management

#### `internal/treesitter` - Tree-sitter Query System
- **Purpose**: Declarative parsing of source code using tree-sitter
- **Languages**: TypeScript, CSS, HTML, JSDoc
- **Components**: QueryManager (compiled query cache), QueryMatcher (per-operation cursors), embedded `.scm` query files

#### `internal/languages` - Per-Language Support
- **Purpose**: Language-specific tree-sitter query definitions and parsing logic
- **Structure**: `typescript/ecmascript.go` (shared JS/TS queries), `registry/` (language registry)

#### `internal/modulegraph` - Module Dependency Graph
- **Purpose**: Tracks import/export relationships between modules
- **Features**: Lazy loading, transitive dependency resolution, metrics

#### `internal/textutil` - Text Processing Utilities
- **Purpose**: UTF-16 code unit offset conversion for LSP protocol compliance
- **Key Function**: Converts between Go's UTF-8 byte offsets and LSP's UTF-16 character offsets

#### `internal/designtokens` - Design Tokens Integration
- **Purpose**: DTCG (Design Tokens Community Group) format support
- **Integration**: Merges design tokens into component manifests

#### `internal/validations` - HTML Validation Data
- **Purpose**: Global HTML attribute data for validation and diagnostics

#### `internal/set` - Set Operations Utility
- **Purpose**: Generic set operations for data processing

### Other Packages

#### `/generate/demodiscovery` - Demo Discovery System
- **Purpose**: Automatically discovers and links component demos
- **Pattern**: Convention-based discovery with manifest integration

## Generate System

`GenerateSession` orchestrates manifest generation cycles, maintaining tree-sitter parsers and query managers across runs for efficient watch mode. `FileDependencyTracker` enables incremental rebuilds via SHA256 content hashing and bidirectional dependency maps (incremental for ≤3 affected modules, full rebuild otherwise). `ModuleBatchProcessor` abstracts the worker pool pattern with dependency-tracking and simple processing modes.

## LSP Document System

All language-specific document types (HTML, TSX, TypeScript) embed `base.BaseDocument`, which provides tree-sitter tree lifecycle management, content storage, UTF-16 offset conversion, and parser pool integration. Language-specific documents implement only their unique parsing and query logic. Trees are ref-counted to prevent use-after-free when concurrent operations hold references during document updates.

## Key Performance Patterns

- Tree-sitter queries compiled once per session, parsers pooled across workers, trees ref-counted
- Incremental rebuilds via SHA256 content hashing and dependency graphs with smart fallback to full rebuild
- O(1) module lookup via persistent index
- Dual manifest access: shallow copy for hot paths, deep copy for isolation

## Development Conventions

### Commit Message Format

This project follows [Conventional Commits](https://www.conventionalcommits.org/) with scopes based on top-level commands:

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or modifying tests
- `chore`: Maintenance tasks

**Scopes:**
- `generate`: Manifest generation functionality
- `list`: Querying and display functionality  
- `validate`: Validation functionality
- `manifest`: Core manifest types and serialization
- `lsp`: Language server protocol
- `mcp`: Model context protocol server
- `serve`: Dev server
- `cmd`: CLI command definitions
- `deps`: Dependencies and build system

**Examples:**
```
feat(generate): add --watch flag for incremental rebuilds
fix(manifest): handle deprecated field unmarshaling
docs(validate): update schema validation examples
test(generate): add comprehensive round-trip tests
```

### Branch Naming Convention

**Format:** `<type>/<description-with-hyphens>`

**Examples:**
- `feat/watch-mode-implementation`
- `fix/manifest-serialization-bug`
- `test/marshal-unmarshal-coverage`
- `docs/architecture-documentation`
- `refactor/generate-session-cleanup`

### Pull Request Guidelines

1. **Create PRs for all non-trivial changes**
2. **Use descriptive titles** matching commit message format
3. **Include comprehensive description** with:
   - Summary of changes
   - Test plan or coverage
   - Breaking changes (if any)
4. **Ensure all tests pass** before requesting review
5. **Add `🤖 Generated with [Claude Code](https://claude.ai/code)` footer when AI-assisted**

### Code Quality Standards

- **Run tests**: `make test` before committing
- **Format code**: Pre-commit hooks handle `go fmt`
- **Race detection**: Tests run with `-race` flag in CI
- **Linting**: `golangci-lint` provides PR feedback
- **Documentation**: Update relevant docs with functional changes
