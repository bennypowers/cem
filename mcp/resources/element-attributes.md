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

Types that reference a generic type parameter carry a `source` reference pointing at the parameter's declaration. Its href is anchored `#L{line}C{column}` (1-based, following GitHub's fragment convention) so the identifier can be located in source. The `cem://element/{tagName}/attributes/{attributeName}` sub-resource returns raw JSON including `type.source`; the upstream CEM schema served at `cem://schema` does not define cem's anchor policy.

Use to implement element attributes correctly.
