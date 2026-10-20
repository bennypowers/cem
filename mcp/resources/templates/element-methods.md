# JavaScript API for `{{.Element.TagName}}` 🔧
{{if .Element.Description}}
{{.Element.Description}}
{{end}}
{{- if gt (len .Element.Methods) 0}}
This element declares **{{len .Element.Methods}} class methods**. They are the verbs of the element: everything it does beyond reacting to attributes and slots. Only public, non-deprecated instance methods are callable on the element; static, protected, private, and deprecated members are listed for completeness.

## Quick Reference 📋

| Method | Returns | Modifiers | What it does |
| ------ | ------- | --------- | ------------ |
{{range .Element.Methods}}| `{{.Name}}` | {{if and .Return .Return.Type}}`{{.Return.Type.Text}}`{{else}}-{{end}} | {{if .Static}}static {{end}}{{if and .Privacy (ne .Privacy "public")}}{{.Privacy}} {{end}}{{if .IsDeprecated}}deprecated{{end}} | {{if .Summary}}{{.Summary | sanitize}}{{else if .Description}}{{.Description | sanitize}}{{else}}-{{end}} |
{{end}}
## The Full Story 📖
{{range .Element.Methods}}
### `{{.Name}}({{range $i, $p := .Parameters}}{{if $i}}, {{end}}{{if $p.Rest}}...{{end}}{{$p.Name}}{{if $p.Type}}: {{$p.Type.Text}}{{end}}{{if $p.Default}} = {{$p.Default}}{{end}}{{end}}){{if and .Return .Return.Type}}: {{.Return.Type.Text}}{{end}}`{{if .IsDeprecated}} ⚠️ *deprecated*{{end}}
{{if .Summary}}
{{.Summary}}
{{end}}{{if .Description}}
{{.Description}}
{{end}}{{if .IsDeprecated}}
**Deprecated**{{if and .Deprecated (ne (printf "%v" .Deprecated.Value) "true")}}: {{.Deprecated.Value}}{{end}}
{{end}}{{if .Parameters}}
**Parameters:**
{{range .Parameters}}- `{{if .Rest}}...{{end}}{{.Name}}`{{if .Optional}} (optional){{end}}{{if .Type}}: `{{.Type.Text}}`{{end}}{{if .Default}} — default `{{.Default}}`{{end}}{{if .Description}} — {{.Description}}{{end}}
{{end}}{{end}}{{if and .Return .Return.Type}}
**Returns:** `{{.Return.Type.Text}}`{{if .Return.Description}} — {{.Return.Description}}{{end}}
{{end}}{{if .InheritedFrom}}
Inherited from `{{.InheritedFrom.Name}}`{{if .InheritedFrom.Module}} ({{.InheritedFrom.Module}}){{end}}.
{{end}}{{if .Source}}
[See the declaration]({{.Source.Href}})
{{end}}{{end}}
## Call It From JavaScript 💻

```javascript
const el = document.querySelector('{{.Element.TagName}}');
{{range .Element.Methods}}{{if and (not .Static) (not .IsDeprecated) (or (not .Privacy) (eq .Privacy "public"))}}el.{{.Name}}({{range $i, $p := .Parameters}}{{if $i}}, {{end}}{{if $p.Default}}{{$p.Default}}{{else}}/* {{if $p.Type}}{{$p.Type.Text}}{{else}}value{{end}} */{{end}}{{end}});
{{end}}{{end}}```
{{end}}
{{- if gt (len .Element.Fields) 0}}{{if gt (len .Element.Methods) 0}}
---
{{end}}
This element also has **{{len .Element.Fields}} class fields** that are not backed by attributes. Unlike attributes, they are only reachable from JavaScript.

| Field | Type | Default | Modifiers | What it holds |
| ----- | ---- | ------- | --------- | ------------- |
{{range .Element.Fields}}| `{{.Name}}` | {{if .Type}}`{{.Type.Text}}`{{else}}-{{end}} | {{if .Default}}`{{.Default}}`{{else}}-{{end}} | {{if .Static}}static {{end}}{{if .Readonly}}readonly {{end}}{{if and .Privacy (ne .Privacy "public")}}{{.Privacy}} {{end}}{{if .IsDeprecated}}deprecated{{end}} | {{if .Summary}}{{.Summary | sanitize}}{{else if .Description}}{{.Description | sanitize}}{{else}}-{{end}} |
{{end}}
{{range .Element.Fields}}{{if or .Description .Summary}}- **`{{.Name}}`**{{if .Type}}: `{{.Type.Text}}`{{end}}{{if .Readonly}} *(readonly)*{{end}}{{if .Static}} *(static)*{{end}}{{if and .Privacy (ne .Privacy "public")}} *({{.Privacy}})*{{end}}{{if .InheritedFrom}} *(inherited from `{{.InheritedFrom.Name}}`)*{{end}} — {{if .Description}}{{.Description | sanitize}}{{else}}{{.Summary | sanitize}}{{end}}
{{end}}{{end}}{{end}}
{{- if and (eq (len .Element.Methods) 0) (eq (len .Element.Fields) 0)}}
## Nothing to Call 🦥

This element keeps its class API private: no public methods or non-attribute fields are declared in the manifest. Configure it with attributes and interact through events instead.
{{end}}
## Keep Exploring 🧭

{{if gt (len .Element.Attributes) 0}}• **[Attributes](cem://element/{{.Element.TagName}}/attributes)** — Configure the behavior
{{end}}{{if gt (len .Element.Events) 0}}• **[Events](cem://element/{{.Element.TagName}}/events)** — React to interactions
{{end}}{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}• **[Styling](cem://element/{{.Element.TagName}}/css)** — Make it look right
{{end}}• **[Full Reference](cem://element/{{.Element.TagName}})** — Everything about this element
• **[Generate HTML](cem://tools/generate_html)** — Let AI create the markup for you
