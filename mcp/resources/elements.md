---
uri: cem://elements
name: elements
mimeType: application/json
dataFetchers:
  - name: elements
    type: all_elements_with_capabilities
    path: ""
    required: true
responseType: json
---

Listing of all custom elements across packages for discovery and planning.

Provides:
- Element tag names and capabilities
- Package organization
- Element APIs (attributes, slots, events, CSS)
- Accessibility features

Use to discover elements and plan component selection.

When choosing elements, use `cem://element/{tagName}` for more information about available APIs.
