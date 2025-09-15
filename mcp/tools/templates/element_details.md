# Complete Reference: `{{.Element.TagName}}`

{{if .Element.Description}}{{.Element.Description}}{{end}}

## Element APIs Overview

This element provides:
{{if gt (len .Element.Attributes) 0}}- **{{len .Element.Attributes}} Attributes** for configuration and behavior{{end}}
{{if gt (len .Element.Slots) 0}}- **{{len .Element.Slots}} Slots** for content placement{{end}}
{{if gt (len .Element.Events) 0}}- **{{len .Element.Events}} Events** for interaction handling{{end}}
{{if gt (len .Element.CssProperties) 0}}- **{{len .Element.CssProperties}} CSS Properties** for theming{{end}}
{{if gt (len .Element.CssParts) 0}}- **{{len .Element.CssParts}} CSS Parts** for styling{{end}}
{{if gt (len .Element.CssStates) 0}}- **{{len .Element.CssStates}} CSS States** for state-based styling{{end}}

{{if gt (len .Element.Attributes) 0}}
## Attributes

{{schemaDesc .SchemaDefinitions "Attribute"}}

{{range .Element.Attributes}}
### `{{.Name}}`{{if .Required}} *(required)*{{end}}

{{if .Description}}{{.Description}}{{end}}

**Type:** `{{.Type}}`{{if .Default}} | **Default:** `{{.Default}}`{{end}}{{if .Values}} | **Values:** {{range $i, $v := .Values}}{{if $i}}, {{end}}`{{$v}}`{{end}}{{end}}
{{end}}
{{end}}

{{if gt (len .Element.Slots) 0}}
## Slots

{{schemaDesc .SchemaDefinitions "Slot"}}

{{range .Element.Slots}}
### {{if .Name}}`{{.Name}}`{{else}}Default slot{{end}}

{{if .Description}}{{.Description}}{{end}}
{{end}}
{{end}}

{{if gt (len .Element.Events) 0}}
## Events

{{schemaDesc .SchemaDefinitions "Event"}}

{{range .Element.Events}}
### `{{.Name}}`

{{if .Description}}{{.Description}}{{end}}

**Type:** `{{.Type}}`
{{end}}
{{end}}

{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}
## CSS Customization

{{if gt (len .Element.CssProperties) 0}}
### CSS Properties

{{schemaDesc .SchemaDefinitions "CssCustomProperty"}}

{{range .Element.CssProperties}}
#### `{{.Name}}`

{{if .Description}}{{.Description}}{{end}}

**Syntax:** `{{.Syntax}}`{{if .Initial}} | **Initial:** `{{.Initial}}`{{end}}
{{end}}
{{end}}

{{if gt (len .Element.CssParts) 0}}
### CSS Parts

{{schemaDesc .SchemaDefinitions "CssPart"}}

{{range .Element.CssParts}}
#### `{{.Name}}`

{{if .Description}}{{.Description}}{{end}}
{{end}}
{{end}}

{{if gt (len .Element.CssStates) 0}}
### CSS States

{{schemaDesc .SchemaDefinitions "CssCustomState"}}

{{range .Element.CssStates}}
#### `{{.Name}}`

{{if .Description}}{{.Description}}{{end}}
{{end}}
{{end}}
{{end}}

## Additional Information

**Module:** {{.Element.Module}}

**Package:** {{.Element.Package}}


---

For focused API information, use:
{{if gt (len .Element.Attributes) 0}}- **`element_attributes`** - Detailed attribute documentation{{end}}
{{if gt (len .Element.Slots) 0}}- **`element_slots`** - Slot usage patterns and guidelines{{end}}
{{if gt (len .Element.Events) 0}}- **`element_events`** - Event details and JavaScript integration{{end}}
{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}- **`element_styling`** - CSS customization guidance{{end}}
- **`generate_html`** - HTML generation with proper structure
- **`suggest_css_integration`** - Design system integration guidance