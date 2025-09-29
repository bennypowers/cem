# CEM LSP Server Architecture

## Overview

The CEM LSP server provides Language Server Protocol support for custom elements in HTML files and TypeScript template literals. It uses tree-sitter for parsing and leverages custom-elements manifest data for intelligent IDE features.

## Core Components

### Server (`server.go`)
LSP protocol handler that delegates requests to organized method handlers using an adapter pattern for clean interfaces.

### Registry (`registry.go`)
Indexes custom elements manifests from workspace, node_modules, and config files. Provides O(1) lookups with automatic file watching for manifest changes. Thread-safe with comprehensive mutex protection.

### Document Manager (`document.go`)
Manages document lifecycle with tree-sitter parsing, incremental updates, parser pooling, and query caching. Handles HTML files directly and extracts TypeScript template literals.

### Method Organization (`methods/`)
Organized by LSP method type with colocated packages:
- `textDocument/completion/` - Autocompletion with type-based suggestions
- `textDocument/hover/` - Element and attribute information
- `textDocument/definition/` - Go-to-definition for custom elements
- `textDocument/publishDiagnostics/` - Real-time validation (slots, tags, attributes, values)
- `textDocument/codeAction/` - One-click autofixes for diagnostics
- `textDocument/references/` - Find all element usages
- `workspace/symbol/` - Search custom elements across workspace

## Data Flow

1. **Initialization**: Load manifests → build indexes → cache queries → start file watching
2. **Document Events**: Parse documents → extract custom elements → cache ASTs
3. **LSP Requests**: Analyze cursor position → query indexes → return results
4. **File Changes**: Incremental reparse → update cached data → reload manifests

## Key Features

### Parsing
- **Tree-sitter Integration**: Shared query system with selective loading for performance
- **Template Literal Support**: Detects `html`` templates, `innerHTML` assignments, etc.
- **Incremental Updates**: Efficient document change handling

### Validation & Autofixes
- **Slot Validation**: Validates slot names with typo suggestions
- **Tag Validation**: Detects typos and missing imports
- **Attribute Validation**: Uses embedded MDN browser-compat-data
- **Value Validation**: Type-based validation (union, literal, number, boolean)

### Navigation
- **Go-to-Definition**: Jump to custom element source definitions
- **Find References**: Workspace-wide element usage search with gitignore filtering
- **Workspace Symbols**: Fuzzy search across all custom elements

## Interface Design

Uses adapter pattern with unified type system:
- **Context Interfaces**: Method-specific contexts for dependency injection
- **Shared Types**: Unified document, element, and attribute interfaces in `lsp/types/`
- **Public API Testing**: All tests use public interfaces only

## Performance

- **Query Caching**: Pre-compiled tree-sitter queries cached at startup
- **O(1) Lookups**: Hash table indexes for manifest data
- **Parser Pooling**: Shared parser instances across operations
- **Incremental Parsing**: Updates only changed document sections