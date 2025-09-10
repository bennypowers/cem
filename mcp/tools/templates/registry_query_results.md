# {{if .Filter}}{{if .Search}}Filtered {{end}}{{end}}Custom Elements Registry ({{len .Elements}} elements)

{{if .Filter}}**Filter:** {{.Filter}}
{{end}}{{if .Search}}**Search:** {{.Search}}
{{end}}
{{if eq (len .Elements) 0}}No elements found matching the specified criteria.
{{else}}{{range $category, $elementList := .Categories}}{{if gt (len $elementList) 0}}## {{$category}} ({{len $elementList}} elements)

{{range $elementList}}### `{{.TagName}}`

{{if .Description}}{{.Description}}
{{end}}
{{if gt (len .Capabilities) 0}}**Capabilities:** {{join .Capabilities " • "}}
{{end}}{{if .Package}}**Package:** {{.Package}}
{{end}}
{{end}}{{end}}{{end}}
## Usage Suggestions

Use `query_registry` with specific parameters to explore elements:
- `tagName`: Get detailed info about a specific element
- `filter`: Filter by capabilities (has-slots, has-events, interactive, form)
- `search`: Search element names and descriptions
{{end}}