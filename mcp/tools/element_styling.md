---
name: element_styling
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to get styling information for"
    context:
      type: string
      description: "Usage context for additional guidance"
  required: ["tagName"]
dataFetchers:
  - name: element
    type: manifest_element
    path: ""
    required: true
  - name: schema
    type: schema_definitions
    path: "definitions"
    required: false
---

Provides comprehensive CSS customization guidance including custom properties, parts, and states. Focuses specifically on element styling with detailed information for theme integration and design system compliance.

Returns detailed information about:
- CSS custom properties with syntax definitions and usage patterns
- CSS parts for targeted styling and design system integration
- CSS custom states for interactive and conditional styling
- Schema documentation for CSS field types and constraints
- CSS examples with proper syntax and best practices
- Design token integration and theming guidance
- Shadow DOM styling considerations and accessibility

Use this tool when implementing element styling, integrating with design systems, or need detailed guidance for CSS customization and theming.