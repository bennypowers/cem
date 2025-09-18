{{if eq .Context "accessibility"}}## ♿ Accessibility Guidelines for `{{.TagName}}`

{{if .ManifestGuidelines}}### 📋 Manifest Requirements
`{{.TagName}}` has specific accessibility guidelines documented in its manifest:

{{range .ManifestGuidelines}}
- **{{.Type}}**: {{.Guideline}}
{{end}}

### ✅ Follow the Manifest
- Implement exactly what the element author documented
- Don't add extra ARIA attributes unless specified
- Trust the element's internal accessibility implementation
- Focus on proper usage patterns as defined

### 🚫 Common Mistakes
- **Don't** override the element's built-in accessibility features
- **Don't** add ARIA attributes whose functionality is covered by other features
mentioned in the manifest
- **Don't** assume standard HTML accessibility patterns apply, consult the
manifest guidelines first
- **Don't** ignore manifest-specified accessibility requirements

{{else}}### 🤝 Trust the Element
`{{.TagName}}` handles its own accessibility internally. Focus on:
- Using the element as documented in its manifest
- Providing meaningful content in slots
- Following any usage guidelines in the element description
- Ensuring proper semantic context around the element

### 🚫 What to Avoid
- **Don't** add ARIA attributes unless the manifest specifies them
- **Don't** assume you need to manage accessibility manually
- **Don't** override the element's semantic behavior
- **Don't** ignore element-specific accessibility guidance in descriptions
{{end}}

{{end}}{{if eq .Context "form"}}## 📝 Form Context Recommendations

### ✅ Best Practices
- `name`: For form submission and data collection
- `required`: Mark required fields for validation
- `disabled`: Disable when not available or applicable
- `aria-invalid`: Indicate validation state (true/false)
- `aria-describedby`: Link to error messages or help text

### ❌ Common Pitfalls
- **Don't** forget the `name` attribute — your form data won't submit
- **Don't** use `placeholder` as the only label (screen readers may miss it)
- **Don't** set `aria-invalid="true"` on valid fields
- **Don't** disable form controls without providing alternative access

{{end}}{{if eq .Context "interactive"}}## 🖱️ Interactive Element Recommendations

### 💡 Smart Choices
- `disabled`: Control interaction state availability
- `aria-pressed`: For toggle buttons (true/false)
- `aria-expanded`: For expandable controls (true/false)
- `aria-controls`: Reference controlled element by ID

### 🚨 Watch Out For
- **Don't** make elements interactive without keyboard support
- **Don't** use `aria-pressed` on non-toggle elements
- **Don't** forget to update `aria-expanded` when content changes
- **Don't** reference non-existent IDs in `aria-controls`

{{end}}{{if eq .Context "styling"}}{{if gt (len .CssProperties) 0}}## 🎨 CSS Customization

This element supports {{len .CssProperties}} CSS custom properties:
{{range .CssProperties}}- `{{.Name}}`: {{.Description}}
{{end}}
{{end}}{{end}}