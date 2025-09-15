---
name: element_attributes
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to get attribute information for"
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
  - name: attributes
    type: attribute_collection
    path: ""
    required: true
---

Provides detailed attribute documentation including types, defaults, valid values, and usage guidelines. Focuses specifically on element attributes with comprehensive information for implementation.

Returns detailed information about:
- Attribute types with union type definitions and value constraints
- Default values and behavioral implications
- Required vs optional attributes with validation guidance
- Usage guidelines extracted from attribute descriptions
- Schema documentation for attribute field types
- Examples of proper attribute usage in HTML
- Accessibility considerations for attribute values

Use this tool when implementing element attributes, need to understand attribute constraints, or want detailed documentation for a specific attribute API.