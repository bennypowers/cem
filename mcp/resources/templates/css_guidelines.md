# CSS Integration Guidelines

## Design Tokens and Theming
{{if .ThemingGuidelines}}
### Tokens
{{range $type, $guidance := .ThemingGuidelines.Tokens}}
- **{{title $type}}**: {{$guidance}}
{{end}}

### Custom Properties
{{range $aspect, $guidance := .ThemingGuidelines.CustomProperties}}
- **{{title $aspect}}**: {{$guidance}}
{{end}}

{{if gt (len .ThemingGuidelines.CustomProperties.Available) 0}}
**Available Custom Properties:**
{{range .ThemingGuidelines.CustomProperties.Available}}
- `{{.}}`
{{end}}
{{end}}

### CSS Parts
{{range $aspect, $guidance := .ThemingGuidelines.Parts}}
- **{{title $aspect}}**: {{$guidance}}
{{end}}

### CSS States
{{range $aspect, $guidance := .ThemingGuidelines.States}}
- **{{title $aspect}}**: {{$guidance}}
{{end}}
{{end}}

## Performance Guidelines
{{if .PerformanceGuidelines}}
### Loading
{{range $strategy, $guidance := .PerformanceGuidelines.Loading}}
- **{{title $strategy}}**: {{$guidance}}
{{end}}

### Rendering
{{range $technique, $guidance := .PerformanceGuidelines.Rendering}}
- **{{title $technique}}**: {{$guidance}}
{{end}}

### Memory Management
{{range $aspect, $guidance := .PerformanceGuidelines.Memory}}
- **{{title $aspect}}**: {{$guidance}}
{{end}}
{{end}}

## Layout Patterns
{{if .LayoutPatterns}}
### Containers
{{range .LayoutPatterns.Containers}}
- {{.}}
{{end}}

### Grid
{{range .LayoutPatterns.Grid}}
- {{.}}
{{end}}

### Flexbox
{{range .LayoutPatterns.Flexbox}}
- {{.}}
{{end}}
{{end}}

## Responsive Design
{{if .ResponsiveGuidelines}}
{{range $guideline, $description := .ResponsiveGuidelines}}
- **{{title $guideline}}**: {{$description}}
{{end}}
{{end}}

## CSS Architecture Best Practices
{{if gt (len .CSSBestPractices) 0}}
{{range .CSSBestPractices}}
- {{.}}
{{end}}
{{end}}