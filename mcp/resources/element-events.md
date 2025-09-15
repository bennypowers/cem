---
uri: cem://element/{tagName}/events
name: element-events
mimeType: text/markdown
uriTemplate: true
dataFetchers:
  - name: element
    type: manifest_element
    path: ""
    required: true
  - name: schema
    type: schema_definitions
    path: "definitions"
    required: false
template: element_events
---

Detailed event documentation including triggers, data payloads, and JavaScript integration patterns for a specific custom element. This resource provides comprehensive event information optimized for AI-assisted event handling implementation.

Provides extensive event details including:
- Event triggers and when they fire
- Event data structures and payload information
- JavaScript event listener setup and best practices
- Custom event patterns and bubbling behavior
- Event timing and lifecycle considerations
- Cross-component event communication patterns
- Accessibility implications of event handling
- Design system integration for event-driven interactions

Use this resource when you need focused information about element events, their triggers, data payloads, and proper event handling patterns according to manifest definitions and JavaScript best practices.