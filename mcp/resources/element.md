---
uri: cem://element/{tagName}
name: element
mimeType: text/markdown
uriTemplate: true
dataFetchers:
  - name: element
    source: elementInfo
    path: ""
    required: true
template: element
---

Overview of a specific custom element including API summary and usage guidance.

Provides:
- Attribute, slot, event, and CSS API counts
- Integration guidance
- Related resource links

Use for element API overview before implementing.

## Available Element Tools

For focused information about specific aspects of an element, use these specialized tools:

- **`element_details`** - Complete reference with all APIs and comprehensive usage guidance
- **`element_attributes`** - Detailed attribute documentation with types, constraints, and usage patterns
- **`element_slots`** - Content guidelines and accessibility considerations for proper slot usage
- **`element_events`** - Event triggers, data payloads, and JavaScript integration patterns
- **`element_styling`** - CSS customization guidance including custom properties, parts, and states

Each tool combines element-specific information with schema definitions and design system context to provide rich, actionable guidance for AI-assisted development.