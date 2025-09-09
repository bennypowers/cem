---
name: validate_html
inputSchema:
  type: object
  properties:
    html:
      type: string
      description: "The HTML content to validate"
    context:
      type: string
      description: "Validation context (e.g., 'accessibility', 'semantic', 'manifest-compliance')"
    tagName:
      type: string
      description: "Optional: specific custom element tag name to focus validation on"
  required: ["html"]
---

Validate HTML structure for accessibility compliance, semantic correctness, and proper custom element usage patterns. Analyzes HTML against custom elements manifest definitions and web accessibility guidelines.

Performs comprehensive validation including:
- Custom element attribute validation against manifest definitions
- Slot content structure and semantic appropriateness
- ARIA attributes and accessibility compliance checking
- Semantic HTML structure and landmark usage
- CSS custom properties and parts usage validation
- Focus management and keyboard navigation patterns
- Color contrast and visual accessibility requirements

Use this tool when reviewing HTML for accessibility issues, validating custom element usage against manifest specifications, checking semantic HTML structure, or ensuring compliance with web accessibility guidelines. The tool provides detailed feedback on accessibility violations and suggests corrections.