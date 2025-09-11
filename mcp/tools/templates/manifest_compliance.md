## Manifest Compliance Validation

{{if gt (len .ManifestIssues) 0}}### ❌ Manifest Compliance Issues

{{range .ManifestIssues}}{{if eq .Type "missing-attribute"}}- **Missing Required Attribute**: `{{.Attribute}}` is required for `<{{.Element}}>`
{{else if eq .Type "invalid-value"}}- **Invalid Attribute Value**: `{{.Attribute}}="{{.Actual}}"` in `<{{.Element}}>`. {{if .Expected}}Valid values: {{.Expected}}{{end}}
{{else}}- **{{.Type}}**: {{.Message}}
{{end}}{{end}}
{{else}}### ✅ No Manifest Compliance Issues Found

{{end}}{{if gt (len .ManifestFeatures) 0}}### 💡 Manifest-Defined Features

{{range .ManifestFeatures}}{{if eq .Type "slots"}}- **`<{{.Element}}>`** supports slots: {{.Details}}
{{else if eq .Type "guidelines"}}- **`<{{.Element}}>`** guidelines: {{.Details}}
{{else if eq .Type "css-apis"}}- **`<{{.Element}}>`** supports {{.Details}} CSS customization options
{{else}}- **{{.Type}}**: {{.Details}}
{{end}}{{end}}
{{end}}