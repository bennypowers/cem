## 🎨 CSS Custom Properties

This element supports {{len .CssProperties}} CSS custom properties for theming:

{{range .CssProperties}}
### `{{.Name}}`
{{if .Description}}{{.Description}}

{{end}}{{if .Syntax}}**Syntax:** `{{.Syntax}}`
{{end}}{{if .Initial}}**Default Value:** `{{.Initial}}`
{{end}}**Inherits:** {{.Inherits}}

**Usage:**
```css
{{$.TagName}} {
{{if .Initial}}  {{.Name}}: {{.Initial}}; /* default value */
{{else}}  {{.Name}}: /* set appropriate value based on {{.Syntax}} */;
{{end}}}
```

{{end}}
### Theming Guidance
When customizing these properties, consider:
- The element's intended purpose and context
- Accessibility requirements (contrast, readability)
- Consistency with your design system
- The property's syntax requirements and constraints

**Example Template:**
```css
{{.TagName}} {
{{range .CssProperties}}{{if .Initial}}  {{.Name}}: {{.Initial}}; /* modify as needed */
{{else}}  {{.Name}}: /* provide value matching syntax: {{.Syntax}} */;
{{end}}{{end}}}
```