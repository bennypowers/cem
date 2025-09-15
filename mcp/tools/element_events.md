---
name: element_events
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to get event information for"
    context:
      type: string
      description: "Usage context for additional guidance"
  required: ["tagName"]
---

Provides detailed event documentation including triggers, data payloads, and JavaScript integration patterns. Focuses specifically on element events with comprehensive information for event handling implementation.

Returns detailed information about:
- Event triggers and when they fire
- Event data structures and payload information
- JavaScript event listener setup and best practices
- Event bubbling and propagation behavior
- Schema documentation for event field types
- Code examples for event handling
- Common event handling patterns and use cases

Use this tool when implementing event handlers, need to understand event data structures, or want detailed documentation for JavaScript integration with element events.