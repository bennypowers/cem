## Goals

`cem serve` will provide a local development server, specifically geared for web component development. The workflow and structure are based on the custom-elements manifest, particularly around the `"demos"` field.
Users will write demos for their elements, and the dev server will present those demos along with helpful utilities like error overlays, event logging, and element knobs (for attributes, css properties, etc).

### Non-Goals

`cem serve` is not meant as a general purpose development server, it's for developing web components in isolation. While it may be possible to develop quite complex `<my-app>`-style components with cem serve, it's primary purpose is *component development*
It's also not meant to run as a service, it's for local development
`cem serve`, like the rest of the `cem` project, does not expose a plugin API, preferring sensible defaults with limited, sensible configuration. We're not aiming to enable complex toolchains (non-standard css syntax transforms, etc). Rather, we encourage standards-based development. (TypeScript transforms are a major exception).

## Command Structure

```bash
cem serve [directory] [flags]
```

**Behavior:**
- Top-level command (like `cem lsp`)
- Reuses generate infrastructure for in-memory manifest (no temp files)
- Watch mode triggers incremental manifest regeneration
- Server stays running, manifest updates propagate via WebSocket

**Flags:**
```
--port, -p         Port to listen on (default: 8000)
--open, -o         Open browser on startup
--config           Path to cem.config.json
--target           Esbuild transform target for TypeScript files (default 2020? 2024?)
--import-map       Path to custom import map overrides
--transform-css    List of glob patterns of CSS files to transform to JavaScript modules
--no-reload        Disable live reload
--template-dir     Custom template directory for demo chrome
```

## Prior Art & References

### @web/dev-server - Buildless Workflow
- **URL**: https://modern-web.dev/docs/dev-server/overview/
- **GitHub**: https://github.com/modernweb-dev/web
- **Key feature**: Buildless ES module development with optional transformations

### Vite - Error Overlay UX
- **URL**: https://vite.dev/
- **GitHub**: https://github.com/vitejs/vite
- **Key feature**: Developer-friendly error overlay with source-mapped stack traces
- Reference for Phase 6 polish work

### PatternFly Elements Dev Server
- **Location**: `patternfly-elements/tools/pfe-tools/dev-server/`
- **Key features**:
  - **Import maps**: Auto-generation from package.json with workspace support
  - **Demo discovery**: Pattern matching and routing for demo files
  - **CSS transforms**: Native constructable stylesheets syntax

### open-wc HMR - Hot Module Replacement
- **URL**: https://open-wc.org/docs/development/hot-module-replacement/
- **GitHub**: https://github.com/open-wc/open-wc/tree/HEAD/packages/dev-server-hmr
- **npm**: `@open-wc/dev-server-hmr`
- **Key feature**: Web component-specific HMR with state preservation
- Reference for stretch goal (99-STRETCH-GOALS.md)

### CEM Project Internal
- **PR #98**: Dependency graph for incremental manifest updates → Phase 4 (transform cache)
- **PR #127**: Workspace/monorepo support → Phase 2 (import maps)
- **Existing packages:**
  - `generate/session.go`: Manifest generation → Phase 1
  - `generate/demodiscovery`: Demo URL discovery → Phase 3

## Architecture Overview

See [01-ARCHITECTURE.md](./01-ARCHITECTURE.md) for detailed cross-cutting architectural concerns:
- HTTP server design and middleware pipeline (7-stage request flow)
- File watching strategy (source files, demos, config)
- Manifest lifecycle management (in-memory generation, regeneration triggers)
- WebSocket live reload implementation (`/__cem-reload` endpoint)
- Error handling taxonomy (user code vs server errors, overlay rendering)
- Cache management (transform cache, LRU eviction, ETags)
- Graceful shutdown sequence (WebSocket cleanup, request draining)
- Port auto-increment behavior
- Logging format and examples

See [99-STRETCH-GOALS.md](./99-STRETCH-GOALS.md) for out-of-scope features: HTTPS, proxy config, custom headers, JSON logs.

---

## Development Workflow

To manage LLM review tool (CodeRabbit) context and enable incremental review:

### Branch Structure
- **Staging branch**: `feat/dev-server` (existing, already created)
- **Phase branches**: `feat/dev-server-phase-N-description`
  - Example: `feat/dev-server-phase-0-testing-infra`
  - Example: `feat/dev-server-phase-1-core-server`
  - Example: `feat/dev-server-phase-2-import-maps`

### PR Workflow
1. Each phase implemented on its own branch
2. Branches manage checklists in overview and plans.
3. Phase PR merges to `feat/dev-server` (staging)
4. After all phases complete, final PR: `feat/dev-server` → `main`

### Review Process
- **Phase PRs** (to staging): Full CodeRabbit review, focused and manageable
- **Final PR** (staging to main): Summary review only
  - Instruct CodeRabbit to skip line-by-line (already reviewed in phase PRs)
  - Focus on integration concerns and breaking changes

### Benefits
- ✅ Each phase PR is small and reviewable
- ✅ CodeRabbit reviews incrementally (avoids context overload)
- ✅ Main branch stays clean (no incomplete features)
- ✅ Releases on main not blocked during development
- ✅ Final PR manageable (just integration review)

### Example Flow
```
main (continues releases)

feat/dev-server (staging)
  ← PR #1: feat/dev-server-phase-0-testing-infra
  ← PR #2: feat/dev-server-phase-1-core-server
  ← PR #3: feat/dev-server-phase-2-import-maps
  ← PR #4: feat/dev-server-phase-3-demo-rendering
  ← PR #5: feat/dev-server-phase-4-transforms
  ← PR #6: feat/dev-server-phase-5a-knobs-core
  ← PR #7: feat/dev-server-phase-5b-knobs-advanced
  ← PR #8: feat/dev-server-phase-5c-knobs-custom (optional)
  ← PR #9: feat/dev-server-phase-6-polish

→ Final PR: feat/dev-server → main (after all phases)
   - CodeRabbit: summary only, skip line-by-line
   - Review: integration concerns, breaking changes
```

## Implementation Plan

- [x] **Phase 0: Test Infrastructure** ✅
  - Goal: Establish fixture-based testing patterns and benchmarks for all subsequent phases
  - Prerequisites: None
  - Provides: Test helpers, mock manifest, WebSocket test client, transform cache test doubles, benchmarks
  - Details: See acceptance criteria at end of this document

- [x] **Phase 1: Core Server** ✅
  - Goal: Basic HTTP server with live reload and middleware pipeline foundation
  - Prerequisites: Phase 0
  - Provides: HTTP server, middleware pipeline, WebSocket endpoint, manifest generation, file watching
  - Consumes: `generate/session.go` (manifest generation)
  - Affects: All subsequent phases
  - Details: [05-AUTORELOAD-HMR.md](./05-AUTORELOAD-HMR.md)
  - **Additions**: Interactive pterm UI with colored logs, persistent status line, WebSocket log broadcasting, default index.html with debug console

- [x] **Phase 2: Import Maps** ✅
  - Goal: Auto-generate import maps from package.json with workspace/monorepo support
  - Prerequisites: Phase 1 (HTTP server)
  - Provides: Import map JSON for HTML injection
  - Consumes: package.json, node_modules (leverages PR #127 for workspace support)
  - Affects: Phase 3 (demo rendering needs import maps)
  - Details: [20-IMPORTMAPS.md](./20-IMPORTMAPS.md)

- [x] **Phase 3: Demo Rendering**
  - Goal: URL routing and demo chrome templates with light DOM by default
  - Prerequisites: Phase 1 (HTTP server), Phase 2 (import maps)
  - Provides: Rendered demo HTML with chrome UI
  - Consumes: Import maps, manifest, demo HTML files
  - Affects: Phase 5 (knobs need chrome infrastructure)
  - Details: [10-URL-REWRITING.md](./10-URL-REWRITING.md), [15-DEMO-CHROME.md](./15-DEMO-CHROME.md)

- [x] **Phase 4: Transforms** ✅ (100% complete)
  - Goal: On-the-fly TypeScript and CSS transformation with dependency-aware caching
  - Prerequisites: Phase 1 (HTTP server, middleware pipeline)
  - Provides: Transformed JavaScript/CSS responses
  - Consumes: Source files, tsconfig.json (leverages PR #98 for dependency graph)
  - Affects: Demo rendering (transformed files loaded via import maps)
  - Note: Can be implemented in parallel with Phases 2 and 3
  - Details: [30-TRANSFORMS.md](./30-TRANSFORMS.md)
  - **Status**: All features implemented, including CSS glob filtering and YAML config

- [ ] **Phase 5: Knobs** (Sequential implementation: 5a → 5b → 5c)
  - Goal: Interactive controls for element attributes, properties, and CSS custom properties
  - Prerequisites: Phase 3 (chrome infrastructure)
  - **Implementation approach**: Complete each sub-phase fully before moving to the next
  - Sub-phases:
    - [x] **Phase 5a: Basic Knobs (REQUIRED)** ✅ - Implement first
      - Goal: Single element demos with basic control types
      - Provides: Knob UI in chrome sidebar
      - Consumes: Manifest data, demo chrome
      - Details: [50-KNOBS-CORE.md](./50-KNOBS-CORE.md)
      - **Status**: Complete with enhancements (type badges, dual color input, deduplication, state sync, markdown descriptions)
    - [ ] **Phase 5b: Advanced Knobs (REQUIRED)** - Implement after 5a is complete
      - Goal: Multiple elements and complex compositions with mutation observers
      - Extends: Phase 5a
      - Details: [51-KNOBS-ADVANCED.md](./51-KNOBS-ADVANCED.md)
    - [ ] **Phase 5c: Custom Templates (OPTIONAL)** - Implement after 5b is complete
      - Goal: User-provided knob templates for specialized controls
      - Extends: Phase 5a/5b
      - Details: [52-KNOBS-CUSTOM.md](./52-KNOBS-CUSTOM.md)

- [ ] **Phase 6: Polish** (Partially implemented early)
  - Goal: Documentation and examples
  - Prerequisites: All previous phases
  - Provides: Enhanced developer experience
  - **Already implemented** (during Phases 1-5):
    - ✅ Error overlay with source-mapped stack traces
    - ✅ Colored logging with structured badges (INFO, WARN, ERROR, DEBUG)
    - ✅ Navigation drawer for demo browsing
    - ✅ 404 error page with helpful navigation
    - ✅ View Transitions API support
    - ✅ Favicon support
    - ✅ Color scheme toggle (light/system/dark with localStorage persistence)
  - **Remaining work**:
    - [ ] Comprehensive documentation
    - [ ] Example projects showcasing features
    - [ ] Performance tuning and final polish

---

## Phase 0 Acceptance Criteria

Since Phase 0 is test infrastructure, its acceptance criteria are unique:

- [ ] Test helpers exist and can be imported by phase tests
- [ ] HTTP server test helper can start/stop test server
- [ ] Mock manifest generation creates valid CEM JSON
- [ ] Fixture pattern established (demo HTML + expected output)
- [ ] WebSocket test client can connect and receive events
- [ ] Transform cache test double implements cache interface
- [ ] Benchmarks defined (even if implementation is stub):
  - Manifest regeneration benchmark exists
  - Transform cache benchmark exists
  - WebSocket broadcast benchmark exists
- [ ] All benchmarks return output (timing may be unrealistic with stubs)
- [ ] Tests fail in "red" phase (prove they detect unimplemented features)
