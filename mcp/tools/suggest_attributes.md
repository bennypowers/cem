---
name: suggest_attributes
inputSchema:
  type: object
  properties:
    tagName:
      type: string
      description: "The custom element tag name to get attribute suggestions for"
    context:
      type: string
      description: "Usage context for contextual attribute recommendations"
    partial:
      type: string
      description: "Partial attribute name for completion suggestions"
  required: ["tagName"]
---

Get intelligent attribute suggestions for custom elements based on manifest definitions, usage context, and accessibility requirements. Provides type-aware recommendations with valid values, defaults, and usage guidance.

Provides detailed attribute information including:
- Available attributes with their TypeScript types and descriptions
- Enum/union type values with explanations of when to use each
- Required vs optional attributes with usage recommendations
- Default values and their semantic meaning
- Accessibility-related attributes and ARIA requirements
- Context-appropriate suggestions based on element usage
- Common attribute combinations and patterns

Use this tool when writing HTML with custom elements, exploring available component APIs, understanding attribute options, or getting contextual recommendations for proper element configuration. The tool helps ensure correct attribute usage and accessibility compliance.