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

Attribute documentation for custom elements including types, defaults, and constraints.

Provides:
- Type definitions and validation rules
- Default values and requirements
- Usage examples and guidelines
- Accessibility considerations

Use to implement element attributes correctly.