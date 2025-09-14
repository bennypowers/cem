# Element Usage Guidelines

{{if gt (len .ElementGuidelines) 0}}
## Element-Specific Guidelines

{{range $tagName, $guidelines := .ElementGuidelines}}
### `{{$tagName}}`
{{if .Description}}{{.Description}}{{end}}

{{if gt (len .Attributes) 0}}**Attributes:**
{{range .Attributes}}
- **`{{.Name}}`**{{if .Required}} ⚠️ (required){{end}}{{if .Type}} - {{.Type}}{{end}}
  {{if .Description}}{{.Description}}{{end}}
  {{if .Default}}Default: `{{.Default}}`{{end}}
  {{if gt (len .Values) 0}}Valid values: {{join .Values ", "}}{{end}}
  {{if gt (len .Guidance) 0}}{{range .Guidance}}
  - {{.}}{{end}}{{end}}
{{end}}
{{end}}

{{if gt (len .Slots) 0}}**Slots:**
{{range .Slots}}
- {{if .Name}}**`{{.Name}}`**{{else}}**Default slot**{{end}}{{if .Description}}: {{.Description}}{{end}}
  {{if gt (len .Guidance) 0}}{{range .Guidance}}
  - {{.}}{{end}}{{end}}
{{end}}
{{end}}

{{if gt (len .Examples) 0}}**Usage Examples:**
{{range .Examples}}
```html
{{.}}
```
{{end}}
{{end}}

{{if gt (len .Guidelines) 0}}**Guidelines:**
{{range .Guidelines}}
- {{.}}
{{end}}
{{end}}

---
{{end}}
{{end}}

## Naming Conventions
{{if .NamingGuidelines}}
### Elements
- **Format**: {{.NamingGuidelines.Elements.Format}}
- **Pattern**: {{.NamingGuidelines.Elements.Pattern}}
{{if gt (len .NamingGuidelines.Elements.Guidelines) 0}}
**Guidelines:**
{{range .NamingGuidelines.Elements.Guidelines}}
- {{.}}
{{end}}
{{end}}

{{if gt (len .NamingGuidelines.Elements.Examples) 0}}
**Examples:**
{{range .NamingGuidelines.Elements.Examples}}
- `{{.}}`
{{end}}
{{end}}

### Attributes
- **Format**: {{.NamingGuidelines.Attributes.Format}}
{{if gt (len .NamingGuidelines.Attributes.Guidelines) 0}}
**Guidelines:**
{{range .NamingGuidelines.Attributes.Guidelines}}
- {{.}}
{{end}}
{{end}}
{{end}}

## Composition Guidelines
{{if .CompositionGuidelines}}
### Slots
{{range $key, $value := .CompositionGuidelines.Slots}}
- **{{title $key}}**: {{$value}}
{{end}}

### Component Nesting
{{range $key, $value := .CompositionGuidelines.Nesting}}
- **{{title $key}}**: {{$value}}
{{end}}

### Component Communication
{{range $key, $value := .CompositionGuidelines.Communication}}
- **{{title $key}}**: {{$value}}
{{end}}
{{end}}
