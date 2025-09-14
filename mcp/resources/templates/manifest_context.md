{{.Overview}}

## Schema Context: {{.SchemaDefinitions.title}}

{{if .SchemaVersion}}
{{if .SchemaDefinitions.description}}{{.SchemaDefinitions.description}}{{end}}

### Schema Field Definitions

{{if .SchemaDefinitions.properties}}
{{$modules := schemaDesc .SchemaDefinitions "modules"}}
{{if $modules}}
#### modules
{{$modules}}
{{end}}
{{end}}

{{$customElement := schemaDesc .SchemaDefinitions "CustomElement"}}
{{if $customElement}}
#### CustomElement
{{$customElement}}

Core properties:
{{$tagName := schemaFieldDesc .SchemaDefinitions "CustomElement" "tagName"}}
{{if $tagName}}- **tagName**: {{$tagName}}{{end}}
{{$attributes := schemaFieldDesc .SchemaDefinitions "CustomElement" "attributes"}}
{{if $attributes}}- **attributes**: {{$attributes}}{{end}}
{{$slots := schemaFieldDesc .SchemaDefinitions "CustomElement" "slots"}}
{{if $slots}}- **slots**: {{$slots}}{{end}}
{{$events := schemaFieldDesc .SchemaDefinitions "CustomElement" "events"}}
{{if $events}}- **events**: {{$events}}{{end}}
{{$cssProperties := schemaFieldDesc .SchemaDefinitions "CustomElement" "cssProperties"}}
{{if $cssProperties}}- **cssProperties**: {{$cssProperties}}{{end}}
{{$cssParts := schemaFieldDesc .SchemaDefinitions "CustomElement" "cssParts"}}
{{if $cssParts}}- **cssParts**: {{$cssParts}}{{end}}
{{$cssStates := schemaFieldDesc .SchemaDefinitions "CustomElement" "cssStates"}}
{{if $cssStates}}- **cssStates**: {{$cssStates}}{{end}}
{{end}}

{{$attribute := schemaDesc .SchemaDefinitions "Attribute"}}
{{if $attribute}}
#### Attribute
{{$attribute}}

Key constraints:
{{$type := schemaFieldDesc .SchemaDefinitions "Attribute" "type"}}
{{if $type}}- **type**: {{$type}}{{end}}
{{$default := schemaFieldDesc .SchemaDefinitions "Attribute" "default"}}
{{if $default}}- **default**: {{$default}}{{end}}
{{$required := schemaFieldDesc .SchemaDefinitions "Attribute" "required"}}
{{if $required}}- **required**: {{$required}}{{end}}
{{end}}

{{$slot := schemaDesc .SchemaDefinitions "Slot"}}
{{if $slot}}
#### Slot
{{$slot}}
{{end}}

{{$cssCustomProperty := schemaDesc .SchemaDefinitions "CssCustomProperty"}}
{{if $cssCustomProperty}}
#### CssCustomProperty
{{$cssCustomProperty}}

Properties:
{{$name := schemaFieldDesc .SchemaDefinitions "CssCustomProperty" "name"}}
{{if $name}}- **name**: {{$name}}{{end}}
{{$syntax := schemaFieldDesc .SchemaDefinitions "CssCustomProperty" "syntax"}}
{{if $syntax}}- **syntax**: {{$syntax}}{{end}}
{{$inherits := schemaFieldDesc .SchemaDefinitions "CssCustomProperty" "inherits"}}
{{if $inherits}}- **inherits**: {{$inherits}}{{end}}
{{$initialValue := schemaFieldDesc .SchemaDefinitions "CssCustomProperty" "initialValue"}}
{{if $initialValue}}- **initialValue**: {{$initialValue}}{{end}}
{{end}}

This schema provides the semantic framework for understanding your specific component data below.
{{else}}
Note: No schema version detected in your manifests. Consider adding schema version for better AI understanding.
{{end}}

## Your Component Data

{{if gt (len .ElementPatterns) 0}}### Element Naming Patterns
{{range .ElementPatterns}}
- **{{.Description}}**: {{join .Examples ", "}}
{{end}}
{{end}}

{{if gt (len .AttributePatterns) 0}}### Common Attributes
{{range .AttributePatterns}}
- `{{.Name}}` - {{.Description}}
{{end}}
{{end}}

{{if gt (len .SlotPatterns) 0}}### Common Slots
{{range .SlotPatterns}}
- `{{.Name}}` - {{.Description}}
{{end}}
{{end}}

{{if gt (len .CSSProperties) 0}}### Your CSS Custom Properties

Your components define **{{len .CSSProperties}} CSS custom properties**:

{{range .CSSProperties}}
- `{{.}}`
{{end}}
{{end}}

{{if gt (len .ExtractedGuidelines) 0}}### Guidelines from Your Manifest Descriptions

These guidelines were extracted from your component and attribute descriptions using RFC 2119 keywords:

{{range .ExtractedGuidelines}}
**{{.Source}}** ({{.Type}}): {{.Guideline}}

{{end}}
{{end}}

## How to Use This Context

This information helps AI understand:
- **Your naming conventions** ({{join .CommonPrefixes ", "}} prefixes)
- **Your component patterns** (common attributes and slots)
- **Your CSS architecture** (custom properties and design tokens)
- **Your documented constraints** (guidelines from descriptions)

When asking for component help, the AI can now reference your actual manifest data and the schema definitions that explain what each field means.