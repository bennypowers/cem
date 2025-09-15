# Element Reference: `{{.Element.TagName | sanitize}}`

{{if .Element.Name}}**Name:** {{.Element.Name | sanitize}}{{end}}
{{if .Element.Summary}}{{.Element.Summary | sanitize}}{{end}}
{{if .Element.Description}}

{{.Element.Description | sanitize}}{{end}}

{{if .Element.Package}}**Package:** `{{.Element.Package | sanitize}}`{{end}}{{if .Element.Module}} | **Module:** `{{.Element.Module | sanitize}}`{{end}}

## API Overview

{{if gt (len .Element.Attributes) 0}}**Attributes:** {{len .Element.Attributes}} configurable properties{{end}}
{{if gt (len .Element.Slots) 0}}**Slots:** {{len .Element.Slots}} content areas{{end}}
{{if gt (len .Element.Events) 0}}**Events:** {{len .Element.Events}} custom events{{end}}
{{if gt (len .Element.CssProperties) 0}}**CSS Properties:** {{len .Element.CssProperties}} custom properties{{end}}
{{if gt (len .Element.CssParts) 0}}**CSS Parts:** {{len .Element.CssParts}} styleable parts{{end}}
{{if gt (len .Element.CssStates) 0}}**CSS States:** {{len .Element.CssStates}} custom states{{end}}

## Integration

{{if .Element.Module}}**Module:** `{{.Element.Module | sanitize}}`

### Import
```javascript
import '{{.Element.Module | sanitize}}';
```
{{end}}

## Related Resources

- **Attributes:** [`cem://element/{{.Element.TagName}}/attributes`](cem://element/{{.Element.TagName}}/attributes)
- **Slots:** [`cem://element/{{.Element.TagName}}/slots`](cem://element/{{.Element.TagName}}/slots)
- **Events:** [`cem://element/{{.Element.TagName}}/events`](cem://element/{{.Element.TagName}}/events)
- **Styling:** [`cem://element/{{.Element.TagName}}/css`](cem://element/{{.Element.TagName}}/css)

---

## Detailed API Documentation

For comprehensive API details, use these focused resources:

{{if gt (len .Element.Attributes) 0}}- **`cem://element/{{.Element.TagName}}/attributes`** - Complete attribute reference with types and constraints{{end}}
{{if gt (len .Element.Slots) 0}}- **`cem://element/{{.Element.TagName}}/slots`** - Slot documentation and content guidelines{{end}}
{{if gt (len .Element.Events) 0}}- **`cem://element/{{.Element.TagName}}/events`** - Event details and JavaScript integration{{end}}
{{if gt (len .Element.CssProperties) 0}}- **`cem://element/{{.Element.TagName}}/css/custom-properties`** - CSS custom properties for theming{{end}}
{{if gt (len .Element.CssParts) 0}}- **`cem://element/{{.Element.TagName}}/css/parts`** - CSS parts for targeted styling{{end}}
{{if gt (len .Element.CssStates) 0}}- **`cem://element/{{.Element.CssStates}}/css/states`** - CSS custom states for conditional styling{{end}}

## Related Resources

- **`cem://elements`** - Browse all available elements
- **`cem://packages`** - Package organization and structure
- **`cem://guidelines`** - Usage guidelines and best practices
- **`cem://accessibility`** - Accessibility patterns and requirements