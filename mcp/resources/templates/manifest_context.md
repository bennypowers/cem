{{.Overview}}

{{if .SchemaVersion}}
*Using {{.SchemaDefinitions.title}}{{if .SchemaVersion}} ({{.SchemaVersion}}){{end}}*
{{if .SchemaDefinitions.description}}

{{.SchemaDefinitions.description}}
{{end}}
{{else}}
*Note: No schema version detected in your manifests. Consider adding schema version for better AI understanding.*
{{end}}

## Custom Elements in Your Project

{{- $customElement := schemaDesc .SchemaDefinitions "CustomElementDeclaration"}}
{{- if $customElement}}
Custom elements are JavaScript classes that extend HTML with new functionality. {{$customElement}}
{{- end}}

{{- if gt (len .Elements) 0}}
{{- range .Elements}}

### `{{.TagName}}`
{{- if .Description}}

{{.Description}}
{{- end}}

{{- if gt (len .Attributes) 0}}
**Attributes:**
{{- range .Attributes}}
- `{{.Name}}`{{if .Type}} ({{.Type}}){{end}}{{if .Description}} - {{.Description}}{{end}}
{{- end}}
{{- end}}

{{- if gt (len .Slots) 0}}
**Slots:**
{{- range .Slots}}
- `{{if .Name}}{{.Name}}{{else}}default{{end}}`{{if .Description}} - {{.Description}}{{end}}
{{- end}}
{{- end}}

{{- if gt (len .CssProperties) 0}}
**CSS Custom Properties:**
{{- range .CssProperties}}
- `{{.Name}}`{{if .Syntax}} ({{.Syntax}}){{end}}{{if .Description}} - {{.Description}}{{end}}
{{- end}}
{{- end}}

{{- if gt (len .CssParts) 0}}
**CSS Parts:**
{{- range .CssParts}}
- `{{.Name}}`{{if .Description}} - {{.Description}}{{end}}
{{- end}}
{{- end}}

{{- if gt (len .CssStates) 0}}
**CSS States:**
{{- range .CssStates}}
- `{{.Name}}`{{if .Description}} - {{.Description}}{{end}}
{{- end}}
{{- end}}

{{- if gt (len .Events) 0}}
**Events:**
{{- range .Events}}
- `{{.Name}}`{{if .Description}} - {{.Description}}{{end}}
{{- end}}
{{- end}}
{{- end}}
{{- else}}
*No custom elements found in your manifests.*
{{- end}}

{{- if gt (len .CommonPrefixes) 0}}

## Naming Conventions

Your elements follow these naming patterns:
{{- range .CommonPrefixes}}
- `{{.}}-*` prefix
{{- end}}
{{- end}}

{{- if gt (len .ExtractedGuidelines) 0}}

## Guidelines from Your Manifest Descriptions

These guidelines were extracted from your component and attribute descriptions using RFC 2119 keywords:
{{- range .ExtractedGuidelines}}

**{{.Source}}** ({{.Type}}): {{.Guideline}}
{{- end}}
{{- end}}

## Schema Reference

{{- $attributeSchema := schemaDesc .SchemaDefinitions "Attribute"}}
{{- if $attributeSchema}}
**Attribute Properties**: {{$attributeSchema}}
{{- $type := schemaFieldDesc .SchemaDefinitions "Attribute" "type"}}
{{- if $type}}
- **type**: {{$type}}
{{- end}}
{{- $default := schemaFieldDesc .SchemaDefinitions "Attribute" "default"}}
{{- if $default}}
- **default**: {{$default}}
{{- end}}
{{- end}}

{{- $slotSchema := schemaDesc .SchemaDefinitions "Slot"}}
{{- if $slotSchema}}
**Slot Properties**: {{$slotSchema}}
{{- $name := schemaFieldDesc .SchemaDefinitions "Slot" "name"}}
{{- if $name}}
- **name**: {{$name}}
{{- end}}
{{- end}}

{{- $cssPropertySchema := schemaDesc .SchemaDefinitions "CssCustomProperty"}}
{{- if $cssPropertySchema}}
**CSS Custom Property Properties**: {{$cssPropertySchema}}
{{- $name := schemaFieldDesc .SchemaDefinitions "CssCustomProperty" "name"}}
{{- if $name}}
- **name**: {{$name}}
{{- end}}
{{- $syntax := schemaFieldDesc .SchemaDefinitions "CssCustomProperty" "syntax"}}
{{- if $syntax}}
- **syntax**: {{$syntax}}
{{- end}}
{{- end}}

## How to Use This Context

Each element above shows its complete API surface. When working with custom elements:
- Use the exact tag names and attribute names shown
- Respect the slot structure for content placement
- Use CSS custom properties for theming and styling
- Follow the documented guidelines and constraints
- Reference the schema properties for understanding data types and requirements