---
uri: cem://element/{tagName}/methods
name: element-methods
mimeType: text/markdown
uriTemplate: true
dataFetchers:
  - name: element
    source: elementInfo
    path: ""
    required: true
  - name: methods
    source: registry
    path: "elements.$.tagName.methods"
    required: false
  - name: fields
    source: registry
    path: "elements.$.tagName.fields"
    required: false
  - name: schema
    source: schema
    path: "definitions"
    required: false
template: element-methods
---

Class method and non-attribute field documentation for a specific custom element. This resource provides the element's JavaScript API surface, which is not covered by the attribute, slot, event, or CSS resources.

Provides extensive member details including:
- Method signatures with parameters, types, and return types
- Parameter defaults, optional and rest markers
- Static, privacy, readonly, and deprecated markers
- Non-attribute class fields with types and default values
- Inheritance provenance for mixin and superclass members
- Source links to the declaring code

Use this resource when you need to call element methods from JavaScript, extend the element, or inspect its class API beyond attributes. Deep links like `cem://element/{tagName}/methods/{methodName}` return a single member as JSON.
