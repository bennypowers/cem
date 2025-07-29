# CEM Project Architecture Overview

This document provides a comprehensive architectural review of the **CEM (Custom Elements Manifest)** project - a command-line tool written in Go that generates, validates, and queries Custom Elements Manifest files for web components.

## Project Goals & Design Principles

### Core Design Goals

1. **Performance First**
   - Prioritize fast build times and responsive watch mode
   - 68% performance improvement with incremental rebuilds (1.1s → 350ms)
   - Efficient resource pooling and caching strategies
   - O(1) lookups and optimized data structures

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

These principles guide architectural decisions throughout the codebase, favoring practical utility over theoretical completeness.

## Project Overview

### Purpose & Mission
CEM is a high-performance Go-based tool that analyzes TypeScript/JavaScript codebases (particularly LitElement components) to generate standardized Custom Elements Manifest files. These manifests provide rich metadata for web components, enabling better documentation, tooling integration, and API discovery.

### Core Commands
1. **`cem generate`** - Analyzes source code and generates manifest files
2. **`cem list`** - Queries manifest files for component information  
3. **`cem validate`** - Validates manifest files against schemas and best practices

### Key Technologies
- **Language**: Go 1.24.3
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

#### `/workspace` - Abstraction Layer
- **Purpose**: Unified interface for local and remote package access
- **Implementations**: FileSystem and Remote workspace contexts
- **Key Features**: Path resolution, file operations, configuration management

#### `/validate` - Schema Validation & Linting
- **Purpose**: Manifest validation against schemas and best practices
- **Features**: Schema validation, intelligent warnings, deprecation detection
- **Schemas**: Support for CEM versions 0.1.0 through 2.1.1-speculative

#### `/list` - Query & Display Engine
- **Purpose**: Fast querying and flexible display of manifest data
- **Features**: Column filtering, table rendering, tree views
- **Output**: Markdown tables, structured data, custom formatting

#### `/queries` - Tree-sitter Query System
- **Purpose**: Declarative parsing of source code using tree-sitter
- **Languages**: TypeScript, CSS, HTML, JSDoc
- **Queries**: Structured query files (`.scm`) for extracting component metadata

### Supporting Packages

#### `/designtokens` - Design Tokens Integration
- **Purpose**: DTCG (Design Tokens Community Group) format support
- **Integration**: Merges design tokens into component manifests

#### `/demodiscovery` - Demo Discovery System
- **Purpose**: Automatically discovers and links component demos
- **Pattern**: Convention-based discovery with manifest integration

#### `/set` - Set Operations Utility
- **Purpose**: Generic set operations for data processing

## Technical Architecture Patterns

### Dependency Injection & Interface Design
The project extensively uses Go interfaces for:
- **WorkspaceContext**: Abstracts file system vs remote package access
- **CssCache**: Enables different caching strategies and testing
- **QueryManager**: Manages expensive tree-sitter query compilation

### Concurrent Processing
- **Worker Pools**: ModuleBatchProcessor abstracts parallel processing
- **Thread Safety**: sync.RWMutex protection for shared state
- **Resource Pooling**: Parser pooling to avoid expensive allocations

### Configuration Management
- **Viper Integration**: Flexible configuration with multiple sources
- **Workspace-Scoped**: Configuration tied to package context
- **CLI Flag Support**: Override configuration via command line

### Testing Strategy
- **Golden Tests**: Extensive use of golden file testing for manifest generation
- **E2E Tests**: End-to-end command testing with real projects
- **Benchmark Tests**: Performance regression detection
- **Fixture-Based**: Comprehensive test fixtures for different component patterns

## Generate System Deep Dive

### Core Performance Achievements

The CEM generate system has been optimized for development workflows with significant performance improvements:

- **Full rebuild**: ~1.1-1.5 seconds
- **Incremental rebuild**: ~350ms (68% improvement, 3x faster)
- **Watch mode**: Real-time updates with intelligent fallback strategies

### Key Architectural Components

#### 1. GenerateSession - Stateful Processing Core
The `GenerateSession` serves as the central orchestrator for generation cycles, maintaining expensive resources like tree-sitter parsers and query managers across multiple runs. This design enables efficient watch mode by avoiding repeated expensive initializations.

**Key Features:**
- Reusable QueryManager (expensive tree-sitter query loading)
- In-memory manifest with dual access patterns (shallow vs deep copy)
- Integrated dependency tracking and CSS caching
- Thread-safe concurrent access via sync.RWMutex

#### 2. FileDependencyTracker - Incremental Intelligence
Tracks file content hashes, module dependencies, and CSS reverse dependencies to enable intelligent incremental rebuilds. Uses SHA256 hashing for reliable change detection and maintains bidirectional dependency maps.

**Smart Fallback Strategy:**
- ≤3 affected modules: Incremental rebuild
- >3 affected modules: Full rebuild (prevents inefficiency)
- No base manifest: Full rebuild (first run)

#### 3. ModuleBatchProcessor - Parallel Processing Abstraction
Abstracts the worker pool pattern used throughout the codebase, eliminating ~200 lines of code duplication. Supports both dependency-tracking and simple processing modes with optimized worker allocation for small incremental builds.

#### 4. Workspace Path Abstraction
Provides clean separation between module paths (used in manifests) and file system paths (used for watching), enabling consistent file handling across different contexts.

**Key Methods:**
- `ModulePathToFS()` - Convert module paths to filesystem paths
- `FSPathToModule()` - Convert filesystem paths to module paths  
- `ResolveModuleDependency()` - Resolve relative dependencies

### File Organization

The generate system was refactored from a single 1079-line file into focused modules:

- **`session_core.go`**: Core GenerateSession functionality
- **`session_deps.go`**: Dependency tracking functionality  
- **`session_incremental.go`**: Incremental processing methods
- **`session_watch.go`**: Watch session functionality
- **`parallel.go`**: Parallel processing abstraction
- **`css.go`**: CSS cache interface and implementation

### Design Patterns

#### Dependency Injection
CSS cache and other setup objects use dependency injection for:
- Testing isolation
- Future cache strategy flexibility
- Clean separation of concerns

#### Worker Pool Pattern
Abstracted parallel processing eliminates code duplication while maintaining:
- Optimized worker allocation
- Thread-safe result aggregation
- Cancellation support

#### Interface Abstraction
Clean interfaces enable:
- Multiple cache implementations
- Testing with mock objects
- Future extensibility

### Thread Safety

All components are designed for concurrent access:
- `sync.RWMutex` protection for shared state
- Proper parser pooling and resource management
- Thread-safe cache operations
- Debounced file change detection (100ms)

## Performance Optimizations

### Tree-sitter Integration
- **Query Compilation**: Expensive queries compiled once per session
- **Parser Pooling**: Reuse parsers across worker threads
- **Memory Management**: Proper tree closure to prevent leaks

### Incremental Processing
- **SHA256 Hashing**: Reliable file change detection
- **Dependency Graphs**: Track module and CSS dependencies
- **Smart Fallbacks**: Automatic full rebuild when incremental is inefficient

### Caching Strategies
- **CSS Parsing Cache**: Avoid re-parsing unchanged CSS files
- **Session-Scoped**: Cache lifetime tied to generation sessions
- **Dependency Injection**: Testable and configurable cache implementations

### Manifest Access Patterns
The system implements dual manifest access patterns to balance performance and thread safety:

- **`GetInMemoryManifest()`**: Shallow copy (~microseconds) for watch mode and hot paths
- **`GetInMemoryManifestDeep()`**: Deep copy (~1-5ms) for future features requiring full isolation

### Module Indexing
O(1) module lookup via persistent module index, replacing O(n) linear searches for better performance in incremental scenarios.

## Quality Assurance

### Code Quality
- **Go Standards**: Idiomatic Go code with clear package boundaries
- **Error Handling**: Comprehensive error handling with context
- **Logging**: Structured logging with configurable verbosity

### Test Coverage
- **Unit Tests**: Package-level testing with mocks and fixtures
- **Integration Tests**: Cross-package functionality testing
- **E2E Tests**: Full command execution testing
- **Golden File Testing**: Regression testing for manifest output

### Error Handling & Resilience

The system implements graceful degradation:
- Fallback to full rebuilds on errors
- Post-processing failures don't break builds
- Detailed error reporting with context
- Resource cleanup on failures

## Documentation & Tooling

### Comprehensive Documentation
- **Hugo-based Website**: bennypowers.github.io/cem/
- **Command Documentation**: Embedded help and examples
- **Benchmark Tracking**: Performance regression monitoring

### Build & Distribution
- **Multi-platform**: Native binaries for major platforms
- **NPM Distribution**: Easy installation via npm/yarn
- **Container Support**: Containerfile for containerized builds

### Development Tools
- **Makefile**: Standardized build and test commands
- **Coverage Tracking**: Comprehensive test coverage reports
- **Benchmark Suite**: Performance testing with visual reports

## Extensibility & Future Development

### Setup Context Wrapper (Recommended)
For managing setup objects like CSS cache, a wrapper pattern is recommended:

```go
type GenerateSetupContext struct {
    W.WorkspaceContext
    cssCache     CssCache
    queryManager *Q.QueryManager
    depTracker   *FileDependencyTracker
}
```

This provides:
- Type safety
- Clean separation of concerns
- Composability for testing and future proofing
- Backward compatibility

### Plugin Architecture Potential
- **Interface-based Design**: Easy to extend with new implementations
- **Query System**: Extensible tree-sitter query definitions
- **Workspace Abstraction**: Support for different source types (Git, npm, etc.)

### Community Contributions
- **Clear Package Boundaries**: Easy to understand and modify
- **Comprehensive Testing**: Changes can be validated quickly
- **Documentation**: Well-documented APIs and architecture decisions

This architecture balances performance, maintainability, and extensibility while providing a solid foundation for web component tooling and future language server features.
