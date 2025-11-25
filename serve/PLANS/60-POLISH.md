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
    - [x] `pf-v6-button` - Button component with variants
    - [x] `pf-v6-switch` - Toggle switch component
    - [x] `pf-v6-text-input` - Text input component
    - [x] `pf-v6-select` - Select/dropdown component
    - [x] `pf-v6-tabs` - Tabs container component (refactored with IDs and parts)
    - [x] `pf-v6-tab` - Individual tab component
    - [x] `pf-v6-card` - Card component for knobs panels
    - [x] `pf-v6-label` - Label/badge component with status variants (used in server logs)
    - [x] `pf-v6-navigation` - Horizontal navigation with scroll buttons
    - [x] `pf-v6-nav-list`, `pf-v6-nav-item`, `pf-v6-nav-link`, `pf-v6-nav-group` - Navigation structure
    - [x] `pf-v6-modal` - Modal dialog component
    - [x] `pf-v6-page`, `pf-v6-page-main`, `pf-v6-page-sidebar` - Page layout structure
    - [x] `pf-v6-masthead` - Application header
    - [x] `pf-v6-skip-to-content` - Accessibility component
    - [x] `pf-v6-toolbar`, `pf-v6-toolbar-group`, `pf-v6-toolbar-item` - Toolbar structure
    - [x] `pf-v6-toggle-group`, `pf-v6-toggle-group-item` - Toggle button group
    - [x] `pf-v6-form`, `pf-v6-form-group`, `pf-v6-form-label` - Form layout components
    - [x] `pf-v6-popover` - Popover component with CSS Anchor Positioning
  - [x] Custom CEM components:
    - [x] `cem-drawer` - Drawer component for footer panel
    - [x] `cem-serve-chrome` - Main dev server UI wrapper
    - [x] `cem-serve-knobs` - Multi-instance knobs panel container
    - [x] `cem-serve-knob-group` - Unified knob control component (consolidated from 3 separate components)
    - [x] `cem-serve-demo` - Demo container with knob event handling
    - [x] `cem-color-scheme-toggle` - Light/dark mode toggle
    - [x] `cem-connection-status` - WebSocket connection indicator
    - [x] `cem-reconnection-content` - Reconnection modal content
    - [x] `cem-transform-error-overlay` - Transform error display

  **Phase 2: Dev Server UI Integration** ✅ **COMPLETE**
  - [x] Integrated pfv6 components into dev server UI:
    - [x] **Page Layout** - `pf-v6-page`, `pf-v6-masthead`, `pf-v6-page-sidebar` for chrome structure
    - [x] **Tabs** - `pf-v6-tabs` and `pf-v6-tab` in footer drawer (Knobs/Server Logs)
    - [x] **Navigation** - `pf-v6-navigation` with horizontal scroll buttons in knobs panel
    - [x] **Cards** - `pf-v6-card` for knobs group containers
    - [x] **Labels** - `pf-v6-label` for server log badges (INFO/WARN/ERROR/DEBUG)
    - [x] **Form controls** - `pf-v6-switch`, `pf-v6-text-input`, `pf-v6-select` in knobs
    - [x] **Modal** - `pf-v6-modal` for reconnection dialog
    - [x] **Drawer** - Custom `cem-drawer` for footer panel
    - [x] **Toolbar** - `pf-v6-toolbar`, `pf-v6-toolbar-group`, `pf-v6-toolbar-item`

  **Phase 2.5: Recent Polish Work** ✅ **COMPLETE**
  - [x] Sticky navigation in knobs panel with proper positioning
  - [x] Fixed WebSocket hijack errors in middleware chain
  - [x] Removed debug logging from pf-v6-navigation
  - [x] Optimized ResizeObserver to only observe necessary elements
  - [x] Fixed knobs scrolling to only scroll within #knobs container
  - [x] Refactored pf-v6-tabs to use IDs instead of classes and expose parts
  - [x] Implemented pf-v6-label with all PatternFly v6 CSS custom properties
  - [x] Converted log timestamps to semantic `<time>` elements with datetime attributes
  - [x] Converted pf-tokens.css to use light-dark() CSS function

  **Phase 2.6: Form Components & Polish** ✅ **COMPLETE**
  - [x] Fixed `pf-v6-select` styles
    - [x] Cleaned up CSS to match PatternFly v6 design tokens
    - [x] Added proper form control wrapper structure
  - [x] Fixed `pf-v6-text-input` styles
    - [x] Created shared `pf-v6-form-control.css` for common form control styles
    - [x] Simplified component CSS to use shared form control base
    - [x] Added proper wrapper structure
  - [x] Fixed `pf-v6-page` sidebar on desktop
    - [x] Corrected sidebar visibility logic for desktop viewport
  - [x] Fixed knob templates
    - [x] Refactored `cem-serve-knob-base.js` to be cleaner
    - [x] Added proper HTML templates for knob components
  - [x] Fixed floating/popup `light-dark()` tokens
    - [x] Updated `pf-tokens.css` with proper light-dark tokens for floating elements
    - [x] Added 261 lines of color tokens for modals, popovers, tooltips, etc.
  - [x] Implemented form layout components
    - [x] Created `pf-v6-form` component with horizontal layout support
    - [x] Created `pf-v6-form-group` component for form field grouping
    - [x] Created `pf-v6-form-label` component with required indicator
    - [x] Added form control IDs to `pf-v6-select`, `pf-v6-switch`, `pf-v6-text-input`
  - [x] Refactored knobs system (decruft)
    - [x] Consolidated three separate knob components into unified `cem-serve-knob-group`
    - [x] Removed `cem-serve-knob-attribute`, `cem-serve-knob-property`, `cem-serve-knob-css-property`
    - [x] Removed unused `knobs-multi.html` template
    - [x] Simplified knobs.html template with unified approach
    - [x] Reduced codebase by 737 lines (deleted), added 564 lines (consolidated)

  **Phase 2.7: Popover Component & Knobs System** ✅ **COMPLETE**
  - [x] Implemented `pf-v6-popover` component with CSS Anchor Positioning
    - [x] Shadow DOM with IDs instead of classes
    - [x] CSS Anchor Positioning with polyfill support for Firefox
    - [x] All position variants (top, bottom, left, right, and -start/-end variants)
    - [x] Arrow positioning with PatternFly CSS variables
    - [x] Closeable button with proper ARIA attributes
    - [x] Empty slot handling (no margin when header/footer empty)
    - [x] Dark mode support with `light-dark()` box-shadow tokens
    - [x] Reset Firefox UA styles for `[popover]` attribute
    - [x] Imported CSS Anchor Positioning polyfill (6,411 lines for function variant, 6,415 for auto)
    - [x] Updated `CemElement` base class to support polyfill integration
  - [x] Fixed knobs dataset binding
    - [x] Changed from `getAttribute('data-*')` to `dataset` property access
    - [x] Fixed multi-instance knob targeting (button knobs were applying to alert)
    - [x] Updated `cem-serve-demo` to look for `data-is-element-knob="true"` marker
    - [x] Removed unused/incomplete `KnobsManager` class (371 lines deleted)
    - [x] Cleaned up knob system architecture

  **Phase 3: Remaining Components & Polish** ⬅️ **CURRENT**
  - [x] **pf-v6-text-input CSS polish** ✅
    - [x] Created shared `pf-v6-form-control.css` for common form control styles
    - [x] Matches PatternFly v6 design tokens
  - [ ] Additional PF v6 components needed:
    - [ ] **pf-v6-expandable-section** - ExpandableSection/disclosure component
      - **Use cases identified**:
        - Debug Information modal: Import map disclosure (`<details id="debug-importmap-details">`)
        - Listing page: Custom Elements list disclosure
        - Knobs panel: CSS custom properties section (should be collapsed by default)
      - Currently using native `<details>` elements, should be replaced with PF v6 component
    - [ ] **pf-v6-number-input** - NumberInput component for number knobs with increment/decrement (nice to have)
  - [ ] Template System Enhancement: Attribute-Dependent SSR
    - [x] **Phase 1: Core Infrastructure** ✅ **ALREADY COMPLETE**
      - [x] Shadowroot middleware already extracts attributes via `attributesToTemplateData()`
      - [x] Templates receive `.Attributes` map (raw attributes) and camelCased keys (`.Href`, `.Variant`, etc.)
      - [x] Supports nested component processing
      - [x] No infrastructure changes needed!
    - [ ] **Phase 2: Component Template Revisions**
      - **Strategy**: Review each component's client-side JS code to identify DOM manipulations that should be moved to SSR templates
      - [x] **pf-v6-button** (`pf-v6-button.js` lines 52-71)
        - Current: JS replaces `<button>` with `<a>` if `href` attribute exists
        - Template: Conditional `<a>` vs `<button>` based on `.Attributes.href`
        - Remove client-side DOM manipulation after SSR implemented
      - [x] **pf-v6-label** (check for removable/close button logic)
        - Review JS for dynamic close button insertion
        - Template: Conditional close button based on `.Attributes.removable`
      - [x] **pf-v6-switch** (check for attribute-dependent structure)
        - Review JS for any attribute-based DOM changes
        - Template: Add any conditional elements
      - [ ] **pf-v6-select** (check for attribute-dependent structure)
        - Review JS for any attribute-based DOM changes
        - Template: Add any conditional elements
      - [ ] **pf-v6-text-input** (check for attribute-dependent structure)
        - Review JS for any attribute-based DOM changes
        - Template: Add any conditional elements
      - [ ] **pf-v6-card** (check for expandable/collapsible logic)
        - Review JS for expanded state handling
        - Template: Conditional `open` attribute on details if using native disclosure
      - [ ] **All components**: Review for shadow DOM class management
        - Identify classes added/removed based on attributes
        - Move to template: `class="base-class {{if .Attributes.compact}}pf-m-compact{{end}}"`
    - [ ] **Phase 3: Testing & Documentation**
      - [ ] Create golden file tests for SSR output with various attribute combinations
      - [ ] Verify client-side hydration doesn't conflict with SSR structure
      - [ ] Update `16-PATTERNFLY-COMPONENTS.md` with SSR conditional rendering patterns
      - [ ] Document template data structure (`.Attributes` map, camelCase keys)
      - [ ] Performance comparison: SSR vs client-side rendering
  - [ ] UI/UX Polish:
    - [ ] Mobile/responsive testing and fixes
    - [ ] Accessibility audit (keyboard nav, screen readers, ARIA)
    - [ ] Visual consistency check across all components
  - [ ] Performance & Testing:
    - [ ] Test with large projects (100+ components)
    - [ ] Browser compatibility testing

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
