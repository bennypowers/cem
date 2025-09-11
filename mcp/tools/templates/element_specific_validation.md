## Element-Specific Validation: `{{.SpecificElement.TagName}}`

{{if .SpecificElement.ElementFound}}**Description:** {{.SpecificElement.Description}}

**Found {{len .SpecificElement.Usages}} usage(s):**

{{range $i, $usage := .SpecificElement.Usages}}{{$index := add $i 1}}{{$index}}. `{{$usage.Html}}`
{{if gt (len $usage.Issues) 0}}   - **Issues:** {{join $usage.Issues "; "}}
{{else}}   - ✅ No issues found
{{end}}{{end}}

{{if gt (len .SpecificElement.Attributes) 0}}**Available Attributes:**
{{range .SpecificElement.Attributes}}- **`{{.Name}}`**{{if .Required}} ⚠️ (required){{end}}{{if .Description}}: {{.Description}}{{end}}
{{end}}{{end}}

{{if gt (len .SpecificElement.Slots) 0}}**Available Slots:**
{{range .SpecificElement.Slots}}- {{if .Name}}**`{{.Name}}`**{{else}}**Default slot**{{end}}{{if .Description}}: {{.Description}}{{end}}
{{end}}{{end}}

{{if gt (len .SpecificElement.Events) 0}}**Available Events:**
{{range .SpecificElement.Events}}- **`{{.Name}}`**{{if .Description}}: {{.Description}}{{end}}
{{end}}{{end}}
{{else}}❌ Element `{{.SpecificElement.TagName}}` not found in HTML

{{end}}