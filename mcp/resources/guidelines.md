---
uri: cem://guidelines
name: guidelines
mimeType: text/markdown
dataFetchers:
  - name: guidelines
    source: registry
    path: "elements"
    required: false
  - name: manifest_overview
    source: registry
    path: ""
    filter: packages_with_metadata
    required: false
template: guidelines
---

Design system guidelines and best practices extracted from component manifests and documentation. This resource provides comprehensive usage guidance for building consistent, accessible user interfaces.

Comprehensive guideline coverage including:
- Element-specific usage guidelines and best practices
- Attribute usage patterns and value recommendations
- Accessibility requirements and ARIA implementation
- CSS integration patterns and theming guidelines
- Component composition and layout recommendations
- Cross-component consistency patterns
- Common anti-patterns and their solutions
- Design token usage and system integration
- Framework-specific implementation guidance
- Performance optimization recommendations

Use this resource to ensure consistent component usage, maintain design system compliance, implement accessibility best practices, or understand the design philosophy behind component APIs.