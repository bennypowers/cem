# {{.Name}} Attribute Context

{{if .Description}}
## Description
{{.Description}}
{{end}}

## Type Information
- **Type**: {{.Type}}{{if .Default}}
- **Default Value**: {{.Default}}{{end}}{{if .Required}}
- **Required**: Yes{{end}}

{{if .Values}}
## Available Options
{{range .Values}}
- `{{.}}`{{end}}

### Usage Example
```html
<element {{$.Name}}="{{index .Values 0}}">
```
{{else}}
## Usage Context
The `{{.Name}}` attribute accepts {{.Type}} values.
{{if .Default}}When not specified, defaults to `{{.Default}}`.{{end}}

### Usage Example  
```html
<element {{.Name}}="{{if .Default}}{{.Default}}{{else}}/* your {{.Type}} value */{{end}}">
```
{{end}}

{{if .Guidelines}}
## Guidelines
{{range .Guidelines}}
- {{.}}{{end}}
{{end}}