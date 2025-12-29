# Meet `{{.Element.TagName | sanitize}}`

{{if .Element.Name}}**{{.Element.Name | sanitize}}** — {{end}}{{if .Element.Summary}}{{.Element.Summary | sanitize}}{{else}}Your next favorite custom element{{end}}
{{if .Element.Description}}

{{.Element.Description | sanitize}}{{end}}

{{if .Element.Package}}**From the `{{.Element.Package | sanitize}}` package**{{end}}{{if .Element.Module}} • **Import from** `{{.Element.Module | sanitize}}`{{end}}

## API Overview

This element comes loaded with:
{{if gt (len .Element.Attributes) 0}}• **{{len .Element.Attributes}} attributes** to configure its behavior{{end}}
{{if gt (len .Element.Slots) 0}}• **{{len .Element.Slots}} content slots** for your HTML{{end}}
{{if gt (len .Element.Events) 0}}• **{{len .Element.Events}} events** to listen for{{end}}
{{if gt (len .Element.CssProperties) 0}}• **{{len .Element.CssProperties}} CSS properties** for theming{{end}}
{{if gt (len .Element.CssParts) 0}}• **{{len .Element.CssParts}} CSS parts** for precise styling{{end}}
{{if gt (len .Element.CssStates) 0}}• **{{len .Element.CssStates}} CSS states** for conditional styling{{end}}

## Ready to Use It? 🚀

{{if .Element.Module}}**First, bring it into your project:**
```javascript
import '{{.Element.Module | sanitize}}';
```

Now you're ready to drop `<{{.Element.TagName}}>` tags into your HTML!
{{else}}This element is ready to use right out of the box — no imports needed!
{{end}}

## Dive Deeper 🏊‍♀️

Want to master this element? These resources have everything you need:

{{if gt (len .Element.Attributes) 0}}• **[Attributes](cem://element/{{.Element.TagName}}/attributes)** — Learn every setting and option{{end}}
{{if gt (len .Element.Slots) 0}}• **[Slots](cem://element/{{.Element.TagName}}/slots)** — Discover where your content goes{{end}}
{{if gt (len .Element.Events) 0}}• **[Events](cem://element/{{.Element.TagName}}/events)** — Hook into what's happening{{end}}
{{if gt (len .Element.CssProperties) 0}}• **[CSS Properties](cem://element/{{.Element.TagName}}/css/custom-properties)** — Make it look exactly right{{end}}
{{if gt (len .Element.CssParts) 0}}• **[CSS Parts](cem://element/{{.Element.TagName}}/css/parts)** — Style individual pieces{{end}}
{{if gt (len .Element.CssStates) 0}}• **[CSS States](cem://element/{{.Element.TagName}}/css/states)** — Style different modes and states{{end}}

## Explore the Ecosystem 🌐

• **[All Elements](cem://elements)** — See what else is available
• **[Package Guide](cem://packages)** — Understand the component library structure
• **[Best Practices](cem://guidelines)** — Learn the recommended patterns
• **[Accessibility Tips](cem://accessibility)** — Build inclusive experiences