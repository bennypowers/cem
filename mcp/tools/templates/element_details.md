# Element Details: `{{.TagName}}`

{{if .Description}}**Description:** {{.Description}}

{{end}}{{if .Package}}**Package:** {{.Package}}

{{end}}{{if gt (len .Attributes) 0}}## Attributes ({{len .Attributes}})

{{range .Attributes}}- **{{.Name}}**{{if .Type}} (_{{.Type}}_){{end}}{{if .Required}} ⚠️ Required{{end}}{{if .Description}}: {{.Description}}{{end}}
{{end}}
{{end}}{{if gt (len .Slots) 0}}## Slots ({{len .Slots}})

{{range .Slots}}- **{{if .Name}}{{.Name}}{{else}}default{{end}}**{{if .Description}}: {{.Description}}{{end}}
{{end}}
{{end}}{{if gt (len .Events) 0}}## Events ({{len .Events}})

{{range .Events}}- **{{.Name}}**{{if .Description}}: {{.Description}}{{end}}
{{end}}
{{end}}{{if gt (len .CssProperties) 0}}## CSS Custom Properties ({{len .CssProperties}})

{{range .CssProperties}}- **{{.Name}}**{{if .Description}}: {{.Description}}{{end}}{{if .Syntax}} ({{.Syntax}}){{end}}
{{end}}
{{end}}{{if gt (len .CssParts) 0}}## CSS Parts ({{len .CssParts}})

{{range .CssParts}}- **{{.Name}}**{{if .Description}}: {{.Description}}{{end}}
{{end}}
{{end}}{{if gt (len .CssStates) 0}}## CSS Custom States ({{len .CssStates}})

{{range .CssStates}}- **{{.Name}}**{{if .Description}}: {{.Description}}{{end}}
{{end}}
{{end}}## Usage Example

```html
<{{.TagName}}{{range .Attributes}}{{if .Required}}{{if .Default}} {{.Name}}="{{.Default}}"{{else if gt (len .Values) 0}} {{.Name}}="{{index .Values 0}}"{{else}} {{.Name}}="value"{{end}}{{end}}{{end}}>
{{if gt (len .Slots) 0}}  <!-- Slot content here -->
{{range .Slots}}  {{if .Name}}<div slot="{{.Name}}">{{.Name}} content</div>{{else}}<div>Default slot content</div>{{end}}
{{end}}{{end}}</{{.TagName}}>
```