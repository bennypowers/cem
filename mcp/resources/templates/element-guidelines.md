# Element Usage Guidelines: `{{.Element.TagName}}`

{{if .Element}}
## Element Guidelines

### Element Description
{{if .Element.Description}}{{.Element.Description | sanitize}}{{else}}No description available.{{end}}

{{if gt (len .Element.Attributes) 0}}
## Attribute Guidelines

{{range .Element.Attributes}}
### `{{.Name}}` Attribute
{{if .Summary}}{{.Summary | sanitize}}{{end}}
{{if .Description}}

{{.Description | sanitize}}{{end}}
{{if .Type}}{{if .Type.Text}}
**Type:** `{{.Type.Text | sanitize}}`{{end}}{{end}}
{{if .Default}}
**Default:** `{{.Default | sanitize}}`{{end}}
{{end}}
{{end}}

{{if gt (len .Element.Slots) 0}}
## Slot Content Guidelines

{{range .Element.Slots}}
### {{if .Name}}`{{.Name}}` Slot{{else}}Default Slot{{end}}
{{if .Summary}}{{.Summary | sanitize}}{{end}}
{{if .Description}}

{{.Description | sanitize}}{{end}}
{{end}}
{{end}}

{{if gt (len .Element.Events) 0}}
## Event Usage Guidelines

{{range .Element.Events}}
### `{{.Name}}` Event
{{if .Summary}}{{.Summary | sanitize}}{{end}}
{{if .Description}}

{{.Description | sanitize}}{{end}}
{{if .Type}}{{if .Type.Text}}
**Type:** `{{.Type.Text | sanitize}}`{{end}}{{end}}
{{end}}
{{end}}

{{else}}
No element data available for guidelines.
{{end}}