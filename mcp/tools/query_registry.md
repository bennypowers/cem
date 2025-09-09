---
name: query_registry
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "Optional: specific element tag name to query"
    filter:
      type: string
      description: "Filter criteria (e.g., 'has-slots', 'has-events', 'interactive')"
    search:
      type: string
      description: "Search term for element names, descriptions, or functionality"
  required: []
---

Query and explore the custom elements registry to discover available components, understand their capabilities, and find elements suitable for specific use cases. Provides comprehensive element information and helps with component discovery.

Enables powerful component exploration through:
- Complete element listing with descriptions and capabilities summary
- Filtered searches by element features (slots, events, CSS properties, etc.)
- Text search across element names, descriptions, and documentation
- Element comparison and similarity discovery
- Package and workspace organization information
- Usage examples and integration patterns for discovered elements
- Accessibility features and compliance information for each element

Use this tool when exploring available custom elements, finding components for specific use cases, understanding component capabilities before implementation, or discovering elements with particular features like accessibility support or specific interaction patterns. Essential for component discovery and architecture planning.