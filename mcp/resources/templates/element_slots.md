# Slots Reference: `{{.Element.TagName}}`

{{if .Element.Description}}{{.Element.Description}}{{end}}

{{if gt (len .Element.Slots) 0}}
## Available Slots

{{schemaDesc .SchemaDefinitions "Slot"}}

{{range .Element.Slots}}
### {{if .Name}}`{{.Name}}` slot{{else}}Default slot{{end}}

{{if .Summary}}{{.Summary}}{{end}}
{{if .Description}}

{{.Description}}{{end}}

{{end}}

## HTML Usage Examples

{{range .Element.Slots}}
{{if .Name}}
### Using `{{.Name}}` slot

```html
<{{$.Element.TagName}}>
  <div slot="{{.Name}}">{{if .Description}}{{.Description}}{{else}}Content for {{.Name}} slot{{end}}</div>
</{{$.Element.TagName}}>
```
{{else}}
### Using default slot

```html
<{{$.Element.TagName}}>
  {{if .Description}}{{.Description}}{{else}}Default content{{end}}
</{{$.Element.TagName}}>
```
{{end}}
{{end}}

### Complete example with all slots

```html
<{{.Element.TagName}}>
{{range .Element.Slots}}{{if .Name}}  <div slot="{{.Name}}">{{if .Description}}{{.Description}}{{else}}{{.Name}} content{{end}}</div>
{{else}}  {{if .Description}}{{.Description}}{{else}}Default content{{end}}
{{end}}{{end}}</{{.Element.TagName}}>
```

{{else}}
## No Named Slots Available

This element does not define any named slots. It may accept content through its default slot or work as a self-contained element.

## HTML Usage

```html
<!-- If the element accepts content -->
<{{.Element.TagName}}>Content here</{{.Element.TagName}}>

<!-- If the element is self-contained -->
<{{.Element.TagName}}></{{.Element.TagName}}>
```
{{end}}

---

For related API information, use:
{{if gt (len .Element.Attributes) 0}}- **`element_attributes`** - Detailed attribute documentation and usage{{end}}
{{if gt (len .Element.Events) 0}}- **`element_events`** - Event details and JavaScript integration{{end}}
{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}- **`element_styling`** - CSS customization with properties, parts, and states{{end}}
- **`element_details`** - Complete element reference with all APIs
- **`generate_html`** - HTML generation with proper slot usage