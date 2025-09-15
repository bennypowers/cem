---
uri: cem://element/{tagName}/attributes
name: element-attributes
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
  - name: attributes
    type: attribute_collection
    path: ""
    required: true
template: element_attributes
---

Detailed attribute documentation for a specific custom element, including types, constraints, default values, and usage patterns. This resource provides comprehensive attribute information optimized for AI-assisted development.

Provides extensive attribute details including:
- Complete type definitions with union types and constraints
- Default values and behavioral implications
- Required vs optional attributes with validation guidance
- Usage guidelines extracted from attribute descriptions
- Schema documentation for attribute field types
- Examples of proper attribute usage in HTML
- Accessibility considerations for attribute values
- Design system integration and consistency patterns

Use this resource when you need focused information about element attributes, their valid values, constraints, and proper usage patterns according to manifest definitions and design system guidelines.