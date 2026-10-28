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

Names and values follow HTML parsing semantics: tag and attribute names are matched case-insensitively (ASCII lowercase), character references are decoded, and when an attribute is repeated the first value is the one validated.

Deprecated elements, deprecated attributes, deprecated slots, and unknown slots are reported by the exported `validate/markup` API but are not rendered by this tool. Attribute conflicts honor manifest types: boolean and untyped attributes are presence-based, while an enum containing both `true` and `false` treats a literal `false` as disabling. Content-attribute redundancy is reported only when a child element is actually assigned to the matching slot.

## Reference Resources

For detailed information about element constraints and validation rules, consult these resources:

- **`cem://elements`** - Discover all available custom elements and their capabilities
- **`cem://element/{tagName}`** - Complete element reference including all validation constraints
- **`cem://element/{tagName}/attributes`** - Attribute types, valid values, and constraint validation
- **`cem://element/{tagName}/slots`** - Slot content guidelines and usage patterns for validation
- **`cem://element/{tagName}/methods`** - Method signatures for post-render scripting
- **`cem://guidelines`** - Design system guidelines and component usage best practices
- **`cem://accessibility`** - Accessibility patterns and validation rules for component compliance

These resources provide the foundational information needed to perform comprehensive custom element validation based on manifest definitions and design system requirements.