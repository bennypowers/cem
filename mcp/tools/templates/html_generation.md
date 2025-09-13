# HTML Generation for `{{.Element.TagName}}`

{{if .Element.Description}}**Element Description:** {{.Element.Description}}

{{end}}## Generated HTML

```html
{{.GeneratedHTML}}
```

## Element Structure and Guidelines

### Tag Information
- **Element:** `{{.Element.TagName}}`{{if .Element.Package}}
- **Package:** {{.Element.Package}}{{end}}
{{if .Element.Description}}
- **Purpose:** {{.Element.Description}}{{end}}

{{if gt (len .Element.Attributes) 0}}### Attributes Available

{{range .Element.Attributes}}#### `{{.Name}}`{{if .Required}} ⚠️ **Required**{{end}}
{{if .Description}}{{.Description}}
{{end}}{{if .Type}}**Type:** `{{.Type}}`
{{end}}{{if .Default}}**Default:** `{{.Default}}`
{{end}}{{if gt (len .Values) 0}}**Valid Values:** {{range $i, $v := .Values}}{{if $i}}, {{end}}`{{$v}}`{{end}}
{{end}}
{{end}}{{end}}

{{if gt (len .Element.Slots) 0}}### Slots Available

{{range .Element.Slots}}#### {{if .Name}}`{{.Name}}` slot{{else}}Default slot{{end}}
{{if .Description}}{{.Description}}
{{end}}{{if .Name}}**Usage:** `<element slot="{{.Name}}">content</element>`
{{else}}**Usage:** Content placed directly inside the element
{{end}}
{{end}}{{end}}

{{if gt (len .Element.Events) 0}}### Events Available

{{range .Element.Events}}#### `{{.Name}}`
{{if .Description}}{{.Description}}
{{end}}
{{end}}{{end}}

## Implementation Notes

### Best Practices
- Use semantic HTML within slots for better accessibility
- Respect the element's intended purpose and constraints
- Check element documentation for specific usage guidelines
{{if gt (len .Element.CssProperties) 0}}
### Styling
This element supports {{len .Element.CssProperties}} CSS custom properties for theming. Use the `suggest_css_integration` tool for styling guidance.
{{end}}
### Accessibility
- This element may implement accessibility features via ElementInternals
- Test with assistive technologies to verify accessibility implementation  
- Avoid adding conflicting ARIA attributes without understanding the element's built-in semantics
- Check shadow DOM implementation for accessibility patterns