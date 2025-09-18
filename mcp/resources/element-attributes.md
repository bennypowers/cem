---
uri: cem://element/{tagName}/attributes
name: element-attributes
mimeType: text/markdown
uriTemplate: true
dataFetchers:
  - name: element
    source: elementInfo
    path: ""
    required: true
  - name: schema
    source: schema
    path: definitions
    required: false
  - name: attributes
    source: elementInfo
    path: ""
    filter: attributes
    required: true
template: element-attributes
---

Attribute documentation for custom elements including types, defaults, and constraints.

Provides:
- Type definitions and validation rules
- Default values and requirements
- Usage examples and guidelines
- Accessibility considerations

Use to implement element attributes correctly.
