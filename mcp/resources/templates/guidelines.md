# Design System Guidelines

## Summary
{{.FetchedData.guidelines.metadata.totalGuidelines}} guidelines extracted from manifest documentation

## Element Guidelines

{{range .FetchedData.guidelines.guidelines}}
{{if eq .type "element"}}
### `{{.source}}` Element
{{.guideline}}

---
{{end}}
{{end}}

## Attribute Guidelines

{{range .FetchedData.guidelines.guidelines}}
{{if eq .type "attribute"}}
### `{{.source}}` Attribute
{{.guideline}}

---
{{end}}
{{end}}

## Guidelines Application

These guidelines are extracted from element and attribute descriptions in your manifests. Use them to:

### Development Standards
- **Consistent usage** - Follow documented patterns for element implementation
- **Attribute validation** - Ensure attributes are used according to their intended purpose
- **Content guidelines** - Understand proper element content and structure
- **Integration patterns** - Learn how elements should work together

### Design System Compliance
- **Component behavior** - Understand expected element behavior and interactions
- **Content strategy** - Follow documented content guidelines and constraints
- **Accessibility requirements** - Implement proper accessibility based on element guidance
- **Performance considerations** - Follow optimization guidelines from element documentation

### Code Review Standards
- **Manifest compliance** - Verify implementations follow documented guidelines
- **Consistency checking** - Ensure similar elements are used similarly across projects
- **Best practice validation** - Apply documented best practices consistently
- **Documentation completeness** - Ensure new elements include proper guidelines

## Guideline Categories

Guidelines are categorized by their source and scope:

- **Element guidelines** - Overall element usage, behavior, and integration patterns
- **Attribute guidelines** - Specific attribute usage, validation, and value constraints
- **Content guidelines** - Slot content requirements and content structure patterns
- **Styling guidelines** - CSS integration patterns and theming requirements

---

For implementation guidance, use:
- **`cem://elements`** - Browse elements with guideline context
- **`cem://element/{tagName}`** - Detailed element information including guidelines
- **`cem://accessibility`** - Accessibility-specific patterns and requirements
- **`cem://packages`** - Package-level organization and consistency patterns