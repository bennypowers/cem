---
uri: cem://element/{tagName}
name: element
mimeType: text/markdown
uriTemplate: true
dataFetchers:
  - name: element
    type: manifest_element
    path: ""
    required: true
template: element
---

Detailed information about a specific custom element including usage patterns, examples, and integration guidance. This resource provides comprehensive element documentation optimized for AI-assisted development.

Provides extensive element details including:
- Complete attribute definitions with types, defaults, and validation
- Slot structure and content recommendations
- Event definitions and usage patterns
- CSS custom properties, parts, and states with examples
- Accessibility requirements and ARIA patterns
- Integration guidelines and best practices
- Usage examples in different contexts (forms, layouts, themes)
- Common attribute combinations and patterns
- Framework-specific integration notes
- Performance considerations and optimization tips

Use this resource when working with a specific element to understand its API, get usage examples, or ensure proper implementation according to design system guidelines.

## Available Element Tools

For focused information about specific aspects of an element, use these specialized tools:

- **`element_details`** - Complete reference with all APIs and comprehensive usage guidance
- **`element_attributes`** - Detailed attribute documentation with types, constraints, and usage patterns
- **`element_slots`** - Content guidelines and accessibility considerations for proper slot usage
- **`element_events`** - Event triggers, data payloads, and JavaScript integration patterns
- **`element_styling`** - CSS customization guidance including custom properties, parts, and states

Each tool combines element-specific information with schema definitions and design system context to provide rich, actionable guidance for AI-assisted development.