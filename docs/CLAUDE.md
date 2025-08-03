# Claude Development Notes

## General rules
- Never leave meaningless whitespace at the end of a line
- HTML start tags that have one attribute should occupy one line
- if they have two attributes, one line or two lines, depending on the content 
length
- if there are more than two attributes, then the first attribute (id, class, 
name, part) on the start tag line, all subsequent attributes on a new line, with
the start of the attribute aligned with the first line's attribute:

```html
<some-tag id="whatever"
          aligned-attrs="true"
          max-single-line-attrs="2">
```

## Docs / HTML

- Use server-rendered web components (DSD) where appropriate

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
