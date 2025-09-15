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

Generate HTML for custom elements with proper slot usage and attributes based on manifest definitions.

Creates HTML that:
- Places content in correct slots
- Includes required attributes with defaults
- Respects accessibility features
- Provides usage notes

Use to generate correct HTML structure for custom elements.

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
