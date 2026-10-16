# Configuring `{{.Element.TagName}}` ⚙️

{{if .Element.Description}}{{.Element.Description}}

{{end}}{{if gt (len .Element.Attributes) 0}}This element comes with **{{len .Element.Attributes}} attributes** to customize its behavior. Think of them as the knobs and dials that control how your element works.

{{schemaDesc .SchemaDefinitions "Attribute"}}

## Quick Reference 📋

| Attribute | Type | Default | What it does |
| --------- | ---- | ------- | ------------ |
{{range .Element.Attributes}}| `{{.Name}}` | {{if .Type}}`{{.Type.Text}}`{{else}}-{{end}} | {{if .Default}}`{{.Default}}`{{else}}-{{end}} | {{if .Summary}}{{.Summary}}{{else if .Description}}{{.Description}}{{else}}-{{end}} |
{{end}}

## The Full Story 📖

{{range .Element.Attributes}}
### `{{.Name}}`
{{if .Summary}}{{.Summary}}{{end}}
{{if .Description}}

{{.Description}}{{end}}

**Type:** {{if .Type}}`{{.Type.Text}}`{{if .Type.Source}} ([declaration]({{.Type.Source.Href}})){{end}}{{else}}-{{end}}{{if .Default}} • **Default:** `{{.Default}}`{{end}}{{if .IsEnum}} • **Options:** {{range $i, $v := .EnumValues}}{{if $i}}, {{end}}`{{$v}}`{{end}}{{end}}

{{end}}

## See It in Action 💻

```html
<!-- Just the basics -->
<{{.Element.TagName}}></{{.Element.TagName}}>

<!-- With some personality -->
<{{.Element.TagName}}{{range .Element.Attributes}}{{if false}} {{.Name}}={{if .Default}}"{{.Default}}"{{else}}{{if .IsEnum}}"{{index .EnumValues 0}}"{{else}}"{{if .Type}}{{.Type.Text}}{{end}}"{{end}}{{end}}{{end}}{{end}}></{{.Element.TagName}}>

<!-- All dressed up -->
<{{.Element.TagName}}{{range .Element.Attributes}}{{if true}} {{.Name}}={{if .Default}}"{{.Default}}"{{else}}{{if .IsEnum}}"{{index .EnumValues 0}}"{{else}}"{{if .Type}}{{.Type.Text}}{{end}}"{{end}}{{end}}{{end}}{{end}}></{{.Element.TagName}}>
```

{{else}}
## Simple and Sweet 🍯

This element keeps things simple — no custom attributes needed! It works great with just standard HTML attributes.

```html
<{{.Element.TagName}}></{{.Element.TagName}}>
```
{{end}}

## Keep Exploring 🧭

{{if gt (len .Element.Slots) 0}}• **[Content Slots](cem://element/{{.Element.TagName}}/slots)** — Learn where your content goes{{end}}
{{if gt (len .Element.Events) 0}}• **[Events](cem://element/{{.Element.TagName}}/events)** — Hook into interactions{{end}}
{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}• **[Styling](cem://element/{{.Element.TagName}}/css)** — Make it look perfect{{end}}
• **[Full Reference](cem://element/{{.Element.TagName}})** — Everything about this element
• **[Generate HTML](cem://tools/generate_html)** — Let AI create the markup for you
