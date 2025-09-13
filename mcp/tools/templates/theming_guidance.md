## 🌈 Theming Guidance

{{if .Theme}}### {{title .Theme}} Theme Considerations
{{else}}### Multi-Theme Support
{{end}}
When implementing themes for this element, consider:

{{if .CssProperties}}**Custom Properties:** Each property has specific constraints
{{range .CssProperties}}- `{{.Name}}`{{if .Syntax}} (syntax: {{.Syntax}}){{end}}{{if .Description}}: {{.Description}}{{end}}
{{end}}

{{end}}{{if .CssParts}}**Styleable Parts:** Consider visual hierarchy and relationships
{{range .CssParts}}- `{{.Name}}`{{if .Description}}: {{.Description}}{{end}}
{{end}}

{{end}}{{if .CssStates}}**Interactive States:** Consider user feedback and transitions
{{range .CssStates}}- `:{{.Name}}`{{if .Description}}: {{.Description}}{{end}}
{{end}}

{{end}}**Best Practices:**
- Use semantic color names and design tokens
- Ensure sufficient contrast ratios for accessibility
- Test themes with the element's intended content
- Consider animation and transition preferences