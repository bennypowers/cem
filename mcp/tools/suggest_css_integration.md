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

Provide CSS integration guidance for custom elements by analyzing their manifest-defined CSS APIs. Uses actual property descriptions, syntax constraints, and initial values from the manifest to guide intelligent styling decisions.

Extracts and presents:
- CSS custom properties with their syntax definitions, descriptions, and initial values from the manifest
- CSS parts with their descriptions and intended styling purposes
- CSS custom states with their descriptions and behavioral context
- Theming considerations based on property constraints and semantics
- Styling guidance that respects the element's documented design intent

The tool avoids hardcoded styling suggestions and instead presents manifest data to help make informed styling decisions. Use this when you need to understand how to properly style a custom element, implement themes that respect element constraints, or integrate components into design systems while following their documented CSS APIs.