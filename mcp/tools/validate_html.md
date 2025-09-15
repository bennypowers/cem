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

Validate custom element usage based on manifest guidelines and best practices. Focuses exclusively on custom element-specific validation, not general HTML or accessibility auditing.

Performs custom element validation including:
- **Slot Content Guidelines**: Validates slotted content against manifest slot descriptions and content rules
- **Attribute Conflicts**: Detects contradictory attribute combinations (e.g., `loading="eager"` + `lazy="true"`)
- **Content/Attribute Redundancy**: Identifies when slot content overrides or conflicts with attribute values
- **Manifest Compliance**: Ensures custom elements are used according to their documented constraints and requirements
- **Custom Element Accessibility**: Validates accessibility patterns specific to custom element usage

Use this tool when reviewing custom element usage in HTML, ensuring proper slot content, validating attribute combinations, or checking compliance with manifest-defined guidelines. The tool provides specific feedback on custom element usage patterns and suggests corrections based on manifest documentation.

This tool does NOT validate general HTML structure, standard accessibility requirements, or semantic HTML patterns - it focuses solely on custom element usage validation.

## Reference Resources

For detailed information about element constraints and validation rules, consult these resources:

- **`cem://elements`** - Discover all available custom elements and their capabilities
- **`cem://element/{tagName}`** - Complete element reference including all validation constraints
- **`cem://element/{tagName}/attributes`** - Attribute types, valid values, and constraint validation
- **`cem://element/{tagName}/slots`** - Slot content guidelines and usage patterns for validation
- **`cem://guidelines`** - Design system guidelines and component usage best practices
- **`cem://accessibility`** - Accessibility patterns and validation rules for component compliance

These resources provide the foundational information needed to perform comprehensive custom element validation based on manifest definitions and design system requirements.