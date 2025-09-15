---
uri: cem://element/{tagName}/guidelines
name: element-guidelines
mimeType: text/markdown
uriTemplate: true
dataFetchers:
  - name: element
    type: manifest_element
    path: ""
    required: true
  - name: guidelines
    type: guidelines_extraction
    path: ""
    required: false
template: element_guidelines
---

Comprehensive guidelines and best practices for a specific custom element, including usage patterns, accessibility requirements, and design system compliance. This resource provides focused guidance for proper element implementation and integration.

Provides detailed guidelines including:
- Element-specific usage guidelines and best practices
- Attribute usage patterns and validation requirements
- Content guidelines for slots and nested elements
- Accessibility requirements and ARIA patterns
- CSS integration guidelines and theming best practices
- Framework-specific integration notes and considerations
- Common patterns and anti-patterns for this element
- Design system compliance and consistency requirements

Use this resource when implementing a specific element to ensure proper usage according to manifest definitions, design system guidelines, and accessibility requirements.

For comprehensive element information beyond guidelines, explore these related resources:

- **`cem://element/{tagName}`** - Complete element reference including all APIs
- **`cem://element/{tagName}/attributes`** - Detailed attribute documentation with constraints
- **`cem://element/{tagName}/slots`** - Content guidelines and slot usage patterns
- **`cem://guidelines`** - Global design system guidelines that apply across elements
- **`cem://accessibility`** - Accessibility patterns and validation rules