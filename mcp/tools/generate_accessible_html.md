---
name: generate_accessible_html
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to generate HTML for"
    content:
      type: object
      description: "Content for named slots (key: slot name, value: content)"
    attributes:
      type: object
      description: "Attributes to include (key: attribute name, value: attribute value)"
    context:
      type: string
      description: "Usage context (e.g., 'form', 'navigation', 'content') for accessibility optimization"
  required: ["tagName"]
---

Generate semantically correct HTML for custom elements with proper ARIA attributes, slot usage, and accessibility compliance. Analyzes element descriptions for usage guidelines and applies accessibility best practices automatically.

Creates complete HTML structures that:
- Place content in correct slots based on manifest definitions
- Include proper attributes with type validation and enum value checking
- Add appropriate ARIA attributes for accessibility compliance
- Follow guidelines extracted from element descriptions and documentation
- Integrate CSS custom properties and parts correctly for styling
- Ensure proper semantic HTML structure within slot content

Use this tool when users ask to generate HTML for custom elements, create component examples, build accessible markup for design system components, or need properly structured HTML templates. The tool automatically applies accessibility best practices and follows element-specific usage guidelines.