---
name: element_slots
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to get slot information for"
    context:
      type: string
      description: "Usage context for additional guidance"
  required: ["tagName"]
---

Provides comprehensive slot usage patterns and content guidelines. Focuses specifically on element slots with detailed information for proper content placement and structure.

Returns detailed information about:
- Named slots with purpose and content guidelines
- Default slot behavior and expected content types
- Content placement patterns and semantic structure
- Accessibility considerations for slotted content
- Schema documentation for slot field types
- HTML examples showing proper slot usage
- Content hierarchy and nesting recommendations

Use this tool when implementing element content structure, need to understand slot patterns, or want detailed guidance for content placement within elements.