# Attributes Reference: `{{.Element.TagName}}`

{{if .Element.Description}}{{.Element.Description}}{{end}}

{{if gt (len .Element.Attributes) 0}}
## Available Attributes

{{schemaDesc .SchemaDefinitions "Attribute"}}

| Attribute | Type | Default | Summary |
| --------- | ---- | ------- | ------- |
{{range .Element.Attributes}}| `{{.Name}}` | {{if .Type}}`{{.Type.Text}}`{{else}}-{{end}} | {{if .Default}}`{{.Default}}`{{else}}-{{end}} | {{if .Summary}}{{.Summary}}{{else if .Description}}{{.Description}}{{else}}-{{end}} |
{{end}}

{{if gt (len .Element.Attributes) 0}}
## Detailed Attribute Information

{{range .Element.Attributes}}
### `{{.Name}}`
{{if .Summary}}{{.Summary}}{{end}}
{{if .Description}}

{{.Description}}{{end}}

**Type:** `{{.Type}}`{{if .Default}} | **Default:** `{{.Default}}`{{end}}{{if .IsEnum}} | **Values:** {{range $i, $v := .EnumValues}}{{if $i}}, {{end}}`{{$v}}`{{end}}{{end}}

{{end}}
{{end}}

## HTML Usage Examples

```html
<!-- Basic usage -->
<{{.Element.TagName}}></{{.Element.TagName}}>

<!-- With attributes -->
<{{.Element.TagName}}{{range .Element.Attributes}}{{if false}} {{.Name}}={{if .Default}}"{{.Default}}"{{else}}{{if .IsEnum}}"{{index .EnumValues 0}}"{{else}}"{{.Type}}"{{end}}{{end}}{{end}}{{end}}></{{.Element.TagName}}>

<!-- With optional attributes -->
<{{.Element.TagName}}{{range .Element.Attributes}}{{if true}} {{.Name}}={{if .Default}}"{{.Default}}"{{else}}{{if .IsEnum}}"{{index .EnumValues 0}}"{{else}}"{{.Type}}"{{end}}{{end}}{{end}}{{end}}></{{.Element.TagName}}>
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
