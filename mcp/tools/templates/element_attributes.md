# Attributes Reference: `{{.Element.TagName}}`

{{if .Element.Description}}{{.Element.Description}}{{end}}

{{if gt (len .Element.Attributes) 0}}
## Available Attributes

{{schemaDesc .SchemaDefinitions "Attribute"}}

{{range .Element.Attributes}}
### `{{.Name}}`{{if .Required}} *(required)*{{end}}

{{if .Description}}{{.Description}}{{end}}

**Type:** `{{.Type}}`{{if .Default}} | **Default:** `{{.Default}}`{{end}}{{if .Values}} | **Values:** {{range $i, $v := .Values}}{{if $i}}, {{end}}`{{$v}}`{{end}}{{end}}

{{if .Guidelines}}
**Guidelines:**
{{range .Guidelines}}- {{.}}
{{end}}{{end}}

{{if .Examples}}
**Examples:**
{{range .Examples}}- {{.}}
{{end}}{{end}}

{{end}}

## HTML Usage Examples

```html
<!-- Basic usage -->
<{{.Element.TagName}}></{{.Element.TagName}}>

<!-- With attributes -->
<{{.Element.TagName}}{{range .Element.Attributes}}{{if .Required}} {{.Name}}={{if .Default}}"{{.Default}}"{{else}}{{if .Values}}"{{index .Values 0}}"{{else}}"{{.Type}}"{{end}}{{end}}{{end}}{{end}}></{{.Element.TagName}}>

<!-- With optional attributes -->
<{{.Element.TagName}}{{range .Element.Attributes}}{{if not .Required}} {{.Name}}={{if .Default}}"{{.Default}}"{{else}}{{if .Values}}"{{index .Values 0}}"{{else}}"{{.Type}}"{{end}}{{end}}{{end}}{{end}}></{{.Element.TagName}}>
```

{{else}}
## No Attributes Available

This element does not define any custom attributes. It works with standard HTML attributes only.

## HTML Usage

```html
<{{.Element.TagName}}></{{.Element.TagName}}>
```
{{end}}

---

For related API information, use:
{{if gt (len .Element.Slots) 0}}- **`element_slots`** - Slot usage patterns and content guidelines{{end}}
{{if gt (len .Element.Events) 0}}- **`element_events`** - Event details and JavaScript integration{{end}}
{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}- **`element_styling`** - CSS customization with properties, parts, and states{{end}}
- **`element_details`** - Complete element reference with all APIs
- **`generate_html`** - HTML generation with proper attribute usage