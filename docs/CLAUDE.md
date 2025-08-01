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

### Components to Refactor
- **liteyoutube**: Convert the current Hugo shortcode + external scripts to a modern web component
  - Replace `lite-yt-embed.js` with a custom element
  - Use declarative shadow DOM for better performance
  - Eliminate need for runtime shortcode detection in head

### Legacy Theme Cleanup ✅
- **hooks/head**: Ejected - removed hooks system completely
  - Replaced with direct component registration via import maps
  - Modern declarative JavaScript modules approach

### Modern Web Components Architecture
- **Approach**: 
  - Use declarative shadow DOM with `<template shadowrootmode="open">`
  - Simple CSS selectors within shadow DOM (no complex class names needed)
  - Leverage CSS nesting for cleaner stylesheets
  - Reference CSS custom properties from host document for theming
  - Progressive enhancement pattern: fallback functionality, upgrade with JavaScript
  - ESM imports with import maps for external dependencies