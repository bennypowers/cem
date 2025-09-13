## 🔄 CSS Custom States

This element supports {{len .CssStates}} custom states:

{{range .CssStates}}
### `:{{.Name}}`
{{if .Description}}{{.Description}}

{{end}}**Usage:**
```css
{{$.TagName}}:{{.Name}} {
{{if .Description}}  /* {{.Description}} */
{{else}}  /* Styles when element is in {{.Name}} state */
{{end}}  /* Add your styling here */
}
```

{{end}}