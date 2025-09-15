---
name: element_details
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to get detailed information for"
    context:
      type: string
      description: "Usage context for additional guidance"
  required: ["tagName"]
---

Provides comprehensive detailed information about a custom element including all APIs, schema definitions, and usage guidance. This is the complete reference for understanding everything about an element.

Returns detailed information about:
- Element description and purpose
- Complete attribute documentation with types, defaults, and usage guidelines
- Slot definitions with content guidelines and accessibility considerations
- Event details with triggers, data, and JavaScript integration patterns
- CSS customization options (custom properties, parts, states)
- Schema documentation for all API types
- Accessibility features and compliance information
- Design system integration guidance

Use this tool when you need complete element documentation, are implementing complex element usage, or need to understand all available APIs for development planning.