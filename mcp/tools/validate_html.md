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
      description: "Validation context for custom elements"
    tagName:
      type: string
      description: "Optional: specific custom element tag name to focus validation on"
  required: ["html"]
---

Validate custom element usage against manifest guidelines. Focuses on custom element validation, not general HTML.

Validates:
- Slot content against guidelines
- Attribute conflicts and requirements
- Manifest compliance
- Custom element accessibility patterns

Use to validate custom element usage in HTML. Does NOT validate general HTML structure.

## Reference Resources

For detailed information about element constraints and validation rules, consult these resources:

- **`cem://elements`** - Discover all available custom elements and their capabilities
- **`cem://element/{tagName}`** - Complete element reference including all validation constraints
- **`cem://element/{tagName}/attributes`** - Attribute types, valid values, and constraint validation
- **`cem://element/{tagName}/slots`** - Slot content guidelines and usage patterns for validation
- **`cem://guidelines`** - Design system guidelines and component usage best practices
- **`cem://accessibility`** - Accessibility patterns and validation rules for component compliance

These resources provide the foundational information needed to perform comprehensive custom element validation based on manifest definitions and design system requirements.