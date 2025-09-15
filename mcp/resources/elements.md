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

Complete listing of all custom element declarations across all packages in the workspace. This resource provides a comprehensive view of available elements for discovery and planning.

Provides detailed elements information including:
- Element tag names and identifiers  
- Element descriptions and purpose
- Package and module organization
- Element capabilities summary (attributes, slots, events, CSS)
- Usage patterns and common combinations
- Framework compatibility information
- Accessibility features and requirements
- Performance characteristics and bundle impact

Use this resource to discover all available custom elements, understand the complete component vocabulary, plan element selection for projects, or analyze capabilities across the entire ecosystem.

When choosing elements, use `cem://element/{tagName}` for more information about available APIs.
