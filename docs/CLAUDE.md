# Claude Development Notes

## Docs Layout Refactoring

### Completed Work ✅
- **Search Component**: Refactored as modern `<fuse-search>` web component with declarative shadow DOM
  - Removed Algolia integration, kept Fuse.js only
  - Implemented proper ARIA accessibility patterns
  - Added DuckDuckGo fallback form
  - Uses import maps and ESM for dependencies

- **Color Mode Toggle**: Refactored as modern `<color-mode-toggle>` web component
  - Three-way toggle (light/auto/dark) using actual radio inputs
  - Proper accessibility with keyboard navigation
  - localStorage persistence and system preference detection
  - Icon-only design with custom SVG icons
  - Integrates with Shoelace theme switching

### Theme System Migration ✅
- **Goal**: Replace SCSS/Sass with vanilla CSS using CSS custom properties
- **Status**: Completed - fully modernized CSS architecture
- **Completed**:
  - ✅ Implemented `color-scheme: light dark` CSS property
  - ✅ Added `light-dark()` CSS function usage throughout
  - ✅ Created centralized CSS custom properties system
  - ✅ Converted all SASS files to modern CSS with nesting
  - ✅ Updated web components to use semantic color variables
  - ✅ Integrated theme switching for external libraries (Shoelace, highlight.js)
  - ✅ Fixed Hugo syntax highlighting to use CSS classes instead of inline styles
  - ✅ Consolidated all theme switching logic in color-mode-toggle component

### Components Removed ✅
- **liteyoutube**: Removed completely - shortcode, scripts, and CSS files deleted

### Legacy Theme Cleanup ✅
- **hooks/head**: Ejected - removed hooks system completely
  - Replaced with direct component registration via import maps
  - Modern declarative JavaScript modules approach

### DSD Component Migration Plan

#### Components for DSD Conversion 🚧
**High Priority - Enhanced functionality:**
- **`<tip>`** → `<info-callout>` web component
  - Expandable/collapsible behavior
  - Icon variants (info, warning, success, error)
  - Animation states
  - Encapsulated styling prevents conflicts

**Keep as Hugo shortcodes (SSR preferred):**
- **`<cta>`** (renamed from `<button>`) - Simple styled links
- **`<icon>`** - Static SVG inclusion via Hugo asset pipeline
- **`<image>`/`<picture>`** - Static media with Hugo's image processing
- **`<mermaid>`** - SSR'd SVG diagrams, no JS interaction needed
- **`<youtube>`** - Simple iframe embed

**Remove completely (not needed):**
- **`<tabs>/<tab>`** - Complex interactions not required
- **`<gallery>`** - Lightbox functionality not needed
- **`<chart>`** - Prefer SSR'd SVG charts over Chart.js
- **`<grid>/<column>`** - CSS Grid/Flexbox sufficient
- **Utility classes** - Prefer semantic CSS over utility classes (future cleanup)

#### Light DOM Components (No DSD needed)
- **Navigation links** - Complex semantic structure better in light DOM
- **Breadcrumbs** - SEO and accessibility benefit from light DOM
- **Basic text formatting** (`<block>`, `<partial>`)

### Modern Web Components Architecture
- **Approach**: 
  - Use declarative shadow DOM with `<template shadowrootmode="open">`
  - Simple CSS selectors within shadow DOM (no complex class names needed)
  - Leverage CSS nesting for cleaner stylesheets
  - Reference CSS custom properties from host document for theming
  - Progressive enhancement pattern: fallback functionality, upgrade with JavaScript
  - ESM imports with import maps for external dependencies

#### Shadow-Class Pattern ✨
**Pattern**: Use component attributes to set classes on shadow DOM containers for styling.

**Implementation**: 
```html
<info-callout type="warning">
  <template shadowrootmode="open">
    <style>
      #content {
        &.info { --callout-color: var(--color-theme); }
        &.warning { --callout-color: #fbbf24; }
        &.error { --callout-color: #ef4444; }
      }
    </style>
    <div id="content" class="{{ $type }}">
```

**Benefits**:
- Clean CSS nesting instead of complex `:host([attr])` selectors
- Easier to read and maintain than attribute-based styling
- Leverages familiar class-based CSS patterns
- Works well with CSS preprocessing and nesting

**Tradeoffs**:
- Requires template logic to map attributes to classes
- Less semantic than pure attribute-based styling
- Can blur the line between component API and implementation

**Best for**: Complex styling variations, theme systems, and cases where CSS nesting significantly improves readability.