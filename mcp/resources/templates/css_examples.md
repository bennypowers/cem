# CSS Integration Context

## CSS Custom Properties
{{range .CssProperties}}
### {{.Name}}
{{if .Description}}{{.Description}}{{end}}

**Technical Details:**
- **Syntax**: {{.Syntax}}{{if .Initial}}
- **Initial Value**: {{.Initial}}{{end}}{{if .Inherits}}
- **Inherits**: Yes{{else}}
- **Inherits**: No{{end}}

**Usage Context:**
```css
element {
  {{.Name}}: {{if .Initial}}{{.Initial}}{{else}}/* {{.Syntax}} value */{{end}};
}
```

{{if .Guidelines}}
**Guidelines:**
{{range .Guidelines}}
- {{.}}{{end}}
{{end}}
{{end}}

{{if gt (len .CssParts) 0}}
## CSS Parts
{{range .CssParts}}
### {{.Name}} Part
{{if .Description}}{{.Description}}{{end}}

**Styling Context:**
```css
element::part({{.Name}}) {
  /* Style the {{.Name}} part */
  /* Common properties to consider: */
  /* color, background, border, padding, margin */
}
```

{{if .Guidelines}}
**Styling Guidelines:**
{{range .Guidelines}}
- {{.}}{{end}}
{{end}}
{{end}}
{{end}}

{{if gt (len .CssStates) 0}}
## CSS Custom States
{{range .CssStates}}
### {{.Name}} State
{{if .Description}}{{.Description}}{{end}}

**State Context:**
```css
element:--{{.Name}} {
  /* Styles applied when element is in {{.Name}} state */
  /* Consider: visual feedback, accessibility, transitions */
}
```

{{if .Guidelines}}
**State Guidelines:**
{{range .Guidelines}}
- {{.}}{{end}}
{{end}}
{{end}}
{{end}}