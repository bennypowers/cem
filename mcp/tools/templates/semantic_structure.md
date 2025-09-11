## Semantic Structure Validation

{{if gt (len .SemanticIssues) 0}}### ❌ Semantic Issues

{{range .SemanticIssues}}{{if eq .Priority "error"}}- **{{.Message}}**
{{else}}- {{.Message}}
{{end}}{{end}}
{{end}}{{if gt (len .SemanticSuggestions) 0}}### 💡 Semantic Suggestions

{{range .SemanticSuggestions}}{{if eq .Type "semantic-elements"}}- **Semantic Elements**: {{.Message}}
{{else if eq .Type "list-structure"}}- **List Structure**: {{.Message}}
{{else if eq .Type "document-structure"}}- **Document Structure**: {{.Message}}
{{else}}- **{{.Type}}**: {{.Message}}
{{end}}{{end}}
{{end}}{{if eq (len .SemanticIssues) 0}}{{if eq (len .SemanticSuggestions) 0}}### ✅ Good Semantic Structure

{{end}}{{end}}