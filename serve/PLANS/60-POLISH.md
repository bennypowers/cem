# Phase 6: Polish

**Status**: 🚧 **IN PROGRESS**

## Overview

Phase 6 focuses on final polish, bug fixes, documentation, and ensuring excellent developer experience across different project types.

**Prerequisites**: All previous phases complete (0-5b) ✅

## Already Implemented (During Phases 1-5)

The following polish features were implemented proactively during earlier phases:

- ✅ Error overlay with source-mapped stack traces
- ✅ Colored logging with structured badges (INFO, WARN, ERROR, DEBUG)
- ✅ Navigation drawer for demo browsing
- ✅ 404 error page with helpful navigation
- ✅ View Transitions API support
- ✅ Favicon support
- ✅ Color scheme toggle (light/system/dark with localStorage persistence)

## Known Issues to Address

### File Watcher: Ignore Build Output Directories

**Issue**: When running in RHDS (Red Hat Design System), file changes under the `_site` directory cause filewatch notifications, even though they are not source files or dependencies of the demos.

**Impact**:
- Unnecessary manifest regeneration
- Performance degradation
- Log noise from non-source file changes
- Potential infinite loops if dev server writes to watched directories

**Root Cause**: File watcher currently watches all files in the working directory tree, including build output directories.

**Solution**: Add configurable ignore patterns for file watching

**Implementation Details**:
- Add `--watch-ignore` flag (comma-separated glob patterns)
- Default ignores: `node_modules/`, `dist/`, `build/`, `_site/`, `.git/`, common build output patterns
- Allow users to customize via flag or config file
- Document common patterns for different project types

**Acceptance Criteria**:
- [x] File watcher respects ignore patterns ✅
- [x] Default patterns ignore common build directories (`.git`, `node_modules`, `dist`, `build`, `_site`, `.cache`) ✅
- [x] `--watch-ignore` flag allows custom patterns (comma-separated) ✅
- [x] Config file supports `watchIgnore` array ✅
- [x] Custom patterns completely override defaults (not merged) ✅
- [x] Changes to ignored files don't trigger manifest regeneration ✅
- [x] Log message indicates when files are ignored (DEBUG level) ✅
- [x] Test coverage for ignore functionality ✅
- [x] Documentation includes examples for monorepos, Eleventy projects, etc. ✅ (see Usage Examples below)

## Usage Examples

### File Watcher Ignore Patterns

**Via Command Line Flag:**
```bash
# Single pattern
cem serve --watch-ignore "_site/**"

# Multiple patterns (comma-separated)
cem serve --watch-ignore "_site/**,dist/**,coverage/**"
```

**Via Config File (`cem.config.yaml` or `cem.config.json`):**
```yaml
serve:
  watchIgnore:
    - "_site/**"
    - "dist/**"
    - ".eleventy/**"
    - "coverage/**"
```

**Common Use Cases:**

1. **Eleventy/11ty Projects:**
   ```bash
   cem serve --watch-ignore "_site/**"
   ```

2. **Monorepos with Build Outputs:**
   ```yaml
   serve:
     watchIgnore:
       - "packages/*/dist/**"
       - "packages/*/build/**"
       - ".nx/**"
       - ".turbo/**"
   ```

3. **Coverage and Test Artifacts:**
   ```yaml
   serve:
     watchIgnore:
       - "coverage/**"
       - ".nyc_output/**"
       - "test-results/**"
   ```

**Default Patterns (used when no custom patterns specified):**
- `.git` - Git repository metadata
- `node_modules` - Node.js dependencies
- `dist` - Common build output directory
- `build` - Common build output directory
- `_site` - Eleventy/Jekyll output directory
- `.cache` - Cache directories

**Important:** Custom patterns completely override defaults. If you specify custom patterns, the defaults are not applied. To keep defaults and add more patterns, explicitly include them:

```yaml
serve:
  watchIgnore:
    - ".git"
    - "node_modules"
    - "dist"
    - "build"
    - "_site"
    - ".cache"
    - "my-custom-dir/**"  # Your additional pattern
```

## Remaining Work

### Documentation

- [ ] Comprehensive README for `cem serve`
- [ ] Flag reference documentation
- [ ] Configuration file schema and examples
- [ ] Troubleshooting guide
- [ ] Architecture documentation (linking to PLANS/)
- [ ] Migration guide from other dev servers

### Example Projects

- [ ] Basic component example
- [ ] Complex composition example (tabs, accordion, etc.)
- [ ] Monorepo/workspace example
- [ ] TypeScript + CSS transforms example
- [ ] Custom knobs demonstration

### Performance Tuning

- [ ] Benchmark manifest regeneration with large projects
- [ ] Optimize transform cache hit rates
- [ ] Profile WebSocket broadcast performance
- [ ] Memory usage analysis with long-running sessions
- [ ] File watcher efficiency with many files

### Final Polish

- [ ] Consistent error messages across all failure modes
- [ ] Graceful degradation when manifest incomplete
- [ ] Keyboard shortcuts documentation
- [ ] Browser compatibility testing
- [ ] Mobile/responsive demo chrome
- [ ] Accessibility audit of demo chrome UI
- [ ] **PatternFly v6 UI Components Implementation** 🚧 **IN PROGRESS**

  **Phase 1: Infrastructure & Core Components** ✅ **COMPLETE**
  - [x] Hybrid SSR/client-side rendering infrastructure
    - [x] Template transformation middleware (Go templates → HTML comments)
    - [x] Constructable Stylesheets utility for CSS efficiency
    - [x] Client-side fallback pattern with error handling
    - [x] Transform tests (6 passing tests)
  - [x] Core pfv6 components implemented with hybrid rendering:
    - [x] `pfv6-button` - Button component with variants
    - [x] `pfv6-switch` - Toggle switch component
    - [x] `pfv6-text-input` - Text input component
    - [x] `pfv6-select` - Select/dropdown component
    - [x] `pfv6-tabs` - Tabs container component
    - [x] `pfv6-tab` - Individual tab component

  **Phase 2: Dev Server UI Integration** ⬅️ **NEXT STEP**
  - [ ] Integrate existing pfv6 components into dev server UI:
    - [ ] **Buttons** - Replace button styles with `<pfv6-button>` in:
      - [ ] Nav drawer (`cem-nav-drawer.html`)
      - [ ] Debug panel (`demo-chrome.html`)
      - [ ] Chrome controls (`demo-chrome.html`)
    - [ ] **Switch** - Replace checkboxes with `<pfv6-switch>` in:
      - [ ] Knobs panel (`templates/knobs.html`)
      - [ ] Color scheme toggle (if applicable)
    - [ ] **Form controls** - Use `<pfv6-text-input>` and `<pfv6-select>` in:
      - [ ] Knobs panel for text/number inputs
      - [ ] Knobs panel for select/enum inputs
    - [ ] **Tabs** - Use `<pfv6-tabs>` and `<pfv6-tab>` where applicable

  **Phase 3: Additional Components** (after Phase 2)
  - [ ] Develop missing pfv6 components:
    - [ ] **pfv6-drawer** - For navigation drawer pattern (`<cem-nav-drawer>`)
    - [ ] **pfv6-header** - Masthead/chrome pattern (`<cem-serve-chrome>`)
    - [ ] **pfv6-disclosure** - Expandable sections for debug panel
    - [ ] Other form controls as needed (checkbox, radio, etc.)
  - [ ] Integrate new components into dev server UI

  **Implementation approach:**
    - ✅ Create local web components in `serve/middleware/routes/templates/elements/`
    - ✅ Use existing template/SSR setup (Declarative Shadow DOM)
    - Follow PF v6 design tokens and patterns
    - No external dependencies - self-contained implementations
    - Maintain current accessibility features (aria-labels, keyboard nav, etc.)
    - Preserve light/dark mode support
    - When implementing web components for patternfly, they must follow the
    patternfly designs, but the APIs do not have to be 1:1 with the react components, for example, we should prefer slots or events to render props.

## Testing Strategy

### Integration Tests

Test the complete developer workflow:
- Start server in various project types
- Verify all features work together
- Test common user scenarios

### Performance Tests

- Large projects (100+ components)
- Rapid file changes (save storms)
- Memory leak detection (24hr runs)
- Transform cache effectiveness

### Cross-Platform Tests

- Linux ✅ (development platform)
- macOS (common developer platform)
- Windows (compatibility check)

## Acceptance Criteria

### Critical (Must Complete)

- [x] File watcher ignores build output directories ⭐ **DONE**
- [ ] PatternFly v6 UI components implemented for all dev server elements
- [ ] Documentation complete and reviewed
- [ ] All known bugs fixed
- [ ] Performance acceptable on large projects (100+ components)
- [ ] No memory leaks in long-running sessions

### Nice to Have

- [ ] Example projects published
- [ ] Video/GIF demonstrations
- [ ] Blog post or announcement
- [ ] Migration scripts from other dev servers

## Out of Scope (Stretch Goals)

See [99-STRETCH-GOALS.md](./99-STRETCH-GOALS.md) for features explicitly deferred:
- HMR (hot module replacement)
- HTTPS support
- Proxy configuration
- Custom HTTP headers
- JSON structured logging
- Plugin API
