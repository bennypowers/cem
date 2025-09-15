# HTML Generation for `{{.Element.TagName}}`

{{if .Element.Description}}{{.Element.Description}}{{end}}

## Generated HTML

```html
{{.GeneratedHTML}}
```

## Element Overview

**APIs Available:**
{{if gt (len .Element.Attributes) 0}}- **{{len .Element.Attributes}} Attributes** ({{range $i, $attr := .Element.Attributes}}{{if $i}}, {{end}}`{{$attr.Name}}`{{end}}){{end}}
{{if gt (len .Element.Slots) 0}}- **{{len .Element.Slots}} Slots** ({{range $i, $slot := .Element.Slots}}{{if $i}}, {{end}}{{if $slot.Name}}`{{$slot.Name}}`{{else}}default{{end}}{{end}}){{end}}
{{if gt (len .Element.Events) 0}}- **{{len .Element.Events}} Events** ({{range $i, $event := .Element.Events}}{{if $i}}, {{end}}`{{$event.Name}}`{{end}}){{end}}
{{if gt (len .Element.CssProperties) 0}}- **{{len .Element.CssProperties}} CSS Properties** for theming{{end}}
{{if gt (len .Element.CssParts) 0}}- **{{len .Element.CssParts}} CSS Parts** for styling{{end}}
{{if gt (len .Element.CssStates) 0}}- **{{len .Element.CssStates}} CSS States** for state-based styling{{end}}

## Detailed Information

For comprehensive information about this element's APIs:

{{if gt (len .Element.Attributes) 0}}- **`element_attributes`** - Detailed attribute documentation, types, and usage guidelines{{end}}
{{if gt (len .Element.Slots) 0}}- **`element_slots`** - Slot usage patterns and content guidelines{{end}}
{{if gt (len .Element.Events) 0}}- **`element_events`** - Event details and JavaScript integration{{end}}
{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}- **`element_styling`** - CSS customization with properties, parts, and states{{end}}
- **`element_details`** - Complete element reference with all APIs and schema information

Use **`suggest_css_integration`** for specific styling guidance and design system integration.