---
name: suggest_css_integration
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to get CSS integration suggestions for"
    styleTarget:
      type: string
      description: "What to style (e.g., 'element', 'parts', 'states', 'properties')"
    context:
      type: string
      description: "Styling context (e.g., 'theme', 'responsive', 'dark-mode')"
  required: ["tagName"]
---

Recommend CSS integration patterns for custom elements including custom properties, CSS parts, custom states, and proper styling approaches. Provides guidance on theming, responsive design, and accessibility-compliant styling.

Offers comprehensive styling guidance for:
- CSS custom properties (variables) with syntax definitions and usage patterns
- CSS parts for styling internal element structure with ::part() selectors
- CSS custom states for conditional styling with :--state selectors
- Proper cascade and inheritance patterns for custom elements
- Theming strategies using CSS custom properties and design tokens
- Responsive design approaches that work with component architecture
- Accessibility considerations for color contrast, focus styles, and motion

Use this tool when integrating custom elements into design systems, creating component themes, implementing responsive designs with custom elements, or ensuring CSS follows accessibility guidelines. The tool helps maintain consistent styling patterns and proper CSS architecture.