# Claude Development Notes

## Docs Layout Refactoring

### Theme System Migration
- **Goal**: Replace SCSS/Sass with vanilla CSS using CSS custom properties
- **Status**: Planning phase
- **Next Steps**:
  1. Define CSS custom properties for color scheme (light/dark mode)
  2. Create centralized theme variables for:
     - Colors (background, text, borders, accent)
     - Typography (font sizes, weights, line heights)
     - Spacing (margins, padding, gaps)
     - Layout (container widths, breakpoints)
  3. Migrate existing SCSS files to vanilla CSS with nesting
  4. Update web components to reference theme variables instead of hardcoded colors

### Web Components Strategy
- **Current**: Starting with `<fuse-search>` component using declarative shadow DOM
- **Approach**: 
  - Use simple CSS selectors within shadow DOM (no need for complex class names)
  - Leverage CSS nesting for cleaner stylesheets
  - Reference CSS custom properties from host document for theming
  - Progressive enhancement pattern: fallback to external search, upgrade to local search

### Components to Refactor
- **liteyoutube**: Convert the current Hugo shortcode + external scripts to a modern web component
  - Replace `lite-yt-embed.js` with a custom element
  - Use declarative shadow DOM for better performance
  - Eliminate need for runtime shortcode detection in head

### Legacy Theme Cleanup
- **hooks/head**: Currently restored for compatibility, plan to replace with modern component system
  - Will replace with explicit import map management and component registration
  - Move from Hugo partials to declarative JavaScript modules

### Color Scheme Variables (Future)
```css
:root {
  --color-bg-primary: light-dark(#ffffff, #1a1a1a);
  --color-bg-secondary: light-dark(#f5f5f5, #2a2a2a);
  --color-text-primary: light-dark(#333333, #ffffff);
  --color-text-secondary: light-dark(#666666, #cccccc);
  --color-border: light-dark(#cccccc, #555555);
  --color-accent: light-dark(#0066cc, #4d9fff);
}
```