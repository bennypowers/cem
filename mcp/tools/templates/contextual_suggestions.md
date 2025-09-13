{{if eq .Context "accessibility"}}## ♿ Accessibility Recommendations

Consider these accessibility attributes for `{{.TagName}}`:
- `role`: Define the element's purpose for screen readers
- `aria-label`: Provide accessible name if not obvious from content
- `aria-describedby`: Reference detailed description element
- `aria-expanded`: For collapsible elements (true/false)
- `aria-hidden`: Hide decorative elements from screen readers (true/false)
- `tabindex`: Control keyboard navigation (0 for focusable, -1 for programmatic focus)

{{end}}{{if eq .Context "form"}}## 📝 Form Context Recommendations

For form usage with `{{.TagName}}`, consider:
- `name`: For form submission and data collection
- `required`: Mark required fields for validation
- `disabled`: Disable when not available or applicable
- `aria-invalid`: Indicate validation state (true/false)
- `aria-describedby`: Link to error messages or help text

{{end}}{{if eq .Context "interactive"}}## 🖱️ Interactive Element Recommendations

For interactive `{{.TagName}}` elements, consider:
- `disabled`: Control interaction state availability
- `aria-pressed`: For toggle buttons (true/false)
- `aria-expanded`: For expandable controls (true/false)
- `aria-controls`: Reference controlled element by ID

{{end}}{{if eq .Context "styling"}}{{if gt (len .CssProperties) 0}}## 🎨 CSS Customization

This element supports {{len .CssProperties}} CSS custom properties:
{{range .CssProperties}}- `{{.Name}}`: {{.Description}}
{{end}}
{{end}}{{end}}