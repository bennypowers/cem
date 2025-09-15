---
name: generate_html
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to generate HTML for"
    content:
      type: string
      description: "Content to place inside the element"
    attributes:
      type: object
      description: "Attributes to include (key: attribute name, value: attribute value)"
    context:
      type: string
      description: "Usage context for additional guidance"
  required: ["tagName"]
---

Generate HTML for custom elements with proper slot usage and required attributes. Inspects the element's APIs as listed in the manifest, and leverages any built-in accessibility features the element has.

Creates HTML structures that:
- Place content in correct slots based on manifest definitions
- Include required attributes with appropriate default values
- Respect any of the element's built-in accessibility features (e.g. implicit role)
- Provide usage notes about available attributes and slots

Use this tool when you need to generate HTML for custom elements, create component examples, or build markup templates. The tool focuses on correct element structure while respecting modern custom element accessibility patterns.

## Resource Discovery

When choosing elements and understanding their capabilities, use these resources:

- **`cem://elements`** - Complete listing of all available custom elements for discovery and planning
- **`cem://element/{tagName}`** - Detailed element information including all APIs and usage guidance
- **`cem://element/{tagName}/attributes`** - Focused attribute documentation with types and constraints
- **`cem://element/{tagName}/slots`** - Content guidelines and accessibility considerations for slot usage
- **`cem://element/{tagName}/events`** - Event patterns and JavaScript integration guidance
- **`cem://element/{tagName}/css/parts`** - CSS parts styling guidance for element customization
- **`cem://element/{tagName}/css/custom-properties`** - CSS custom properties for theming and styling
- **`cem://element/{tagName}/css/states`** - CSS custom states for interactive styling patterns

These resources provide the context needed to generate proper HTML with correct attributes, slots, and styling integration.
