## 🔧 CSS Parts

This element exposes {{len .CssParts}} CSS parts for granular styling:

{{range .CssParts}}
### `{{.Name}}`
{{if .Description}}{{.Description}}

{{end}}**Usage:**
```css
{{$.TagName}}::part({{.Name}}) {
{{if .Description}}  /* {{.Description}} */
{{end}}  /* Add your styling here */
}
```

{{end}}
{{if gt (len .CssParts) 1}}### Styling Multiple Parts
Consider how these parts work together in the element's design according to their descriptions.
Each part represents a part of the element's shadow DOM which the element author desired to be made available to be styled independently.

{{end}}
