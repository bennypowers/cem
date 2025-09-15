---
uri: cem://accessibility
name: accessibility
mimeType: text/markdown
dataFetchers:
  - name: accessibility_patterns
    type: accessibility_extraction
    path: ""
    required: false
  - name: manifest_context
    type: workspace_manifest_overview
    path: ""
    required: false
template: accessibility
---

Comprehensive accessibility patterns, guidelines, and validation rules for custom elements. This resource provides WCAG 2.1 AA compliance guidance and implementation patterns for building accessible web components.

Extensive accessibility coverage including:
- WCAG 2.1 AA compliance requirements and testing procedures
- ARIA implementation patterns for different component types
- Keyboard navigation requirements and focus management
- Screen reader compatibility and announcement patterns
- Color contrast requirements and visual accessibility
- Form accessibility and labeling best practices
- Dynamic content and live region implementation
- Custom element accessibility architecture patterns
- Cross-component accessibility consistency
- Automated testing integration and validation rules

Use this resource to implement accessible components, validate accessibility compliance, understand ARIA patterns, or ensure inclusive design practices throughout your component library.