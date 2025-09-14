{{.Overview}}

## Schema Context: {{.SchemaDefinitions.title}}{{if .SchemaVersion}} ({{.SchemaVersion}}){{end}}

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

{{$customElement := schemaDesc .SchemaDefinitions "CustomElementDeclaration"}}
{{if $customElement}}
#### CustomElementDeclaration
{{$customElement}}

Core properties:
{{$tagName := schemaFieldDesc .SchemaDefinitions "CustomElementDeclaration" "tagName"}}
{{if $tagName}}- **tagName**: {{$tagName}}{{end}}
{{$attributes := schemaFieldDesc .SchemaDefinitions "CustomElementDeclaration" "attributes"}}
{{if $attributes}}- **attributes**: {{$attributes}}{{end}}
{{$slots := schemaFieldDesc .SchemaDefinitions "CustomElementDeclaration" "slots"}}
{{if $slots}}- **slots**: {{$slots}}{{end}}
{{$events := schemaFieldDesc .SchemaDefinitions "CustomElementDeclaration" "events"}}
{{if $events}}- **events**: {{$events}}{{end}}
{{$cssProperties := schemaFieldDesc .SchemaDefinitions "CustomElementDeclaration" "cssProperties"}}
{{if $cssProperties}}- **cssProperties**: {{$cssProperties}}{{end}}
{{$cssParts := schemaFieldDesc .SchemaDefinitions "CustomElementDeclaration" "cssParts"}}
{{if $cssParts}}- **cssParts**: {{$cssParts}}{{end}}
{{$cssStates := schemaFieldDesc .SchemaDefinitions "CustomElementDeclaration" "cssStates"}}
{{if $cssStates}}- **cssStates**: {{$cssStates}}{{end}}
{{end}}

{{if gt (len .AttributePatterns) 0}}
#### Attribute
Element attributes define the interface between HTML and your custom element. Each attribute has a name, type constraints, and optional default values.

Key properties:
{{$name := schemaFieldDesc .SchemaDefinitions "Attribute" "name"}}
{{if $name}}- **name**: {{$name}}{{end}}
{{$type := schemaFieldDesc .SchemaDefinitions "Attribute" "type"}}
{{if $type}}- **type**: {{$type}}{{end}}
{{$default := schemaFieldDesc .SchemaDefinitions "Attribute" "default"}}
{{if $default}}- **default**: {{$default}}{{end}}
{{$required := schemaFieldDesc .SchemaDefinitions "Attribute" "required"}}
{{if $required}}- **required**: {{$required}}{{end}}
{{$fieldName := schemaFieldDesc .SchemaDefinitions "Attribute" "fieldName"}}
{{if $fieldName}}- **fieldName**: {{$fieldName}}{{end}}
{{end}}

{{if gt (len .SlotPatterns) 0}}
#### Slot
Content slots define where child content can be placed within your custom element's shadow DOM. Slots enable flexible, composable component designs.

Key properties:
{{$name := schemaFieldDesc .SchemaDefinitions "Slot" "name"}}
{{if $name}}- **name**: {{$name}}{{end}}
{{end}}

{{if gt (len .CSSProperties) 0}}
#### CssCustomProperty
CSS custom properties (CSS variables) provide a styling API for your custom elements. They allow theme customization and design system integration.

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

{{if gt (len .CssPartPatterns) 0}}
#### CssPart
CSS parts provide styling hooks for internal components of your custom elements. They enable precise styling without breaking encapsulation.

Key properties:
{{$name := schemaFieldDesc .SchemaDefinitions "CssPart" "name"}}
{{if $name}}- **name**: {{$name}}{{end}}
{{end}}

{{if gt (len .CssStatePatterns) 0}}
#### CssCustomState
CSS custom states represent dynamic conditions of your custom elements. They provide semantic styling hooks for interactive states.

Key properties:
{{$name := schemaFieldDesc .SchemaDefinitions "CssCustomState" "name"}}
{{if $name}}- **name**: {{$name}}{{end}}
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

{{if gt (len .CssPartPatterns) 0}}### Common CSS Parts

{{range .CssPartPatterns}}
- `{{.Name}}` - {{.Description}}
{{end}}
{{end}}

{{if gt (len .CssStatePatterns) 0}}### Common CSS States

{{range .CssStatePatterns}}
- `{{.Name}}` - {{.Description}}
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