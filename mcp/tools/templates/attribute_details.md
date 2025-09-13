### `{{.Name}}`{{if .Type}} _{{.Type}}_{{end}}

{{if .Description}}{{.Description}}

{{end}}{{if .Type}}**Type:** `{{.Type}}`
{{end}}{{if .Default}}**Default:** `{{.Default}}`
{{end}}**Required:** {{if .Required}}Yes ⚠️{{else}}No{{end}}

{{if gt (len .Values) 0}}**Valid Values:**
{{range .Values}}- `{{.}}`
{{end}}
{{end}}{{if gt (len .Examples) 0}}**Examples:**
{{range .Examples}}- `{{.}}`
{{end}}
{{end}}{{if gt (len .Guidelines) 0}}**Guidelines:**
{{range .Guidelines}}- {{.}}
{{end}}
{{end}}