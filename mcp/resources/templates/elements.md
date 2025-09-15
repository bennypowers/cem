# Elements Discovery

## Summary
{{.FetchedData.elements.metadata.totalElements}} custom elements available

{{if .FetchedData.elements.metadata.categories}}
## Element Categories

{{range $category, $count := .FetchedData.elements.metadata.categories}}{{if gt $count 0}}**{{$category | title}}:** {{$count}}{{if eq $category "form-elements"}} elements{{else if eq $category "layout-elements"}} elements{{else if eq $category "navigation-elements"}} elements{{else}} elements{{end}}
{{end}}{{end}}{{end}}

## Available Elements

| Tag Name       | Package        | Module       | Summary                            |
| -------------- | -------------- | ------------ | ---------------------------------- |
{{range .FetchedData.elements.elements}}
| `{{.tagName | sanitize}}` | `{{.package | sanitize}}` | `{{.module | sanitize}}`| {{if .summary}}{{.summary | sanitize}}{{end}} |
{{end}}

## Element Capabilities

Use this overview to find elements by their functional capabilities:

- **Configurable** - Elements with custom attributes for behavior control
- **Content-slots** - Elements that accept slotted content
- **Interactive** - Elements that emit custom events
- **Themeable** - Elements with CSS custom properties for styling
- **Styleable** - Elements with CSS parts for targeted styling
- **Stateful** - Elements with CSS custom states for conditional styling
- **Form-elements** - Components for form construction and input handling
- **Layout-elements** - Components for page structure and content organization
- **Navigation-elements** - Components for site navigation and user guidance

## Development Guidance

### Element Selection
Choose elements based on:
1. **Functional requirements** - Match capabilities to your needs
2. **Package consistency** - Prefer elements from the same design system

### Integration Planning
- Elements with more attributes require more configuration
- Interactive elements need event handling setup
- Themeable elements benefit from design token integration
- Layout elements should align with your CSS architecture

---

For detailed element information, use:
- **`cem://element/{tagName}`** - Complete element reference
- **`cem://element/{tagName}/attributes`** - Focused attribute documentation
- **`cem://packages`** - Package-level organization and structure
- **`cem://guidelines`** - Usage guidelines and best practices
