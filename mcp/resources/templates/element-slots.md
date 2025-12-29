# Content Areas for `{{.Element.TagName}}` 📍

{{if .Element.Description}}{{.Element.Description}}

{{end}}{{if gt (len .Element.Slots) 0}}This element has **{{len .Element.Slots}} content areas** where you can place your HTML. Think of slots as designated parking spots for your content!

{{schemaDesc .SchemaDefinitions "Slot"}}

## Content Destinations 🎯

{{range .Element.Slots}}
### {{if .Name}}The `{{.Name}}` spot{{else}}Main content area{{end}}

{{if .Summary}}{{.Summary}}{{end}}
{{if .Description}}

{{.Description}}{{end}}

{{end}}

## Put It to Work 💼

{{range .Element.Slots}}
{{if .Name}}
### Drop content in the `{{.Name}}` spot:

```html
<{{$.Element.TagName}}>
  <div slot="{{.Name}}">Your {{.Name}} content goes here!</div>
</{{$.Element.TagName}}>
```
{{else}}
### Fill the main content area:

```html
<{{$.Element.TagName}}>
  Your main content goes here!
</{{$.Element.TagName}}>
```
{{end}}
{{end}}

### The full experience — all slots filled:

```html
<{{.Element.TagName}}>
{{range .Element.Slots}}{{if .Name}}  <div slot="{{.Name}}">Perfect {{.Name}} content</div>
{{else}}  Main content that makes sense
{{end}}{{end}}</{{.Element.TagName}}>
```

### 🚫 Common Slot Mistakes
- **Don't** forget the `slot` attribute — your content might end up in the wrong place
- **Don't** use slot names that aren't defined (they'll be ignored)
- **Don't** mix slotted and non-slotted content without understanding the fallback behavior

{{else}}
## Keep It Simple 🎯

This element doesn't use named slots — it either accepts content directly or works on its own.

```html
<!-- Try with content -->
<{{.Element.TagName}}>Your content here</{{.Element.TagName}}>

<!-- Or standalone -->
<{{.Element.TagName}}></{{.Element.TagName}}>
```
{{end}}

## Keep Learning 📚

{{if gt (len .Element.Attributes) 0}}• **[Attributes](cem://element/{{.Element.TagName}}/attributes)** — Configure the behavior{{end}}
{{if gt (len .Element.Events) 0}}• **[Events](cem://element/{{.Element.TagName}}/events)** — React to interactions{{end}}
{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}• **[Styling](cem://element/{{.Element.TagName}}/css)** — Make it beautiful{{end}}
• **[Full Reference](cem://element/{{.Element.TagName}})** — Everything in one place
• **[Generate HTML](cem://tools/generate_html)** — Get AI help with the markup
