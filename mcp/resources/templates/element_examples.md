# {{.TagName}} Usage Examples

## Basic Usage
```html
<{{.TagName}}></{{.TagName}}>
```

{{if .Description}}
## Description
{{.Description}}
{{end}}

{{if gt (len .Attributes) 0}}
## With Attributes
{{range .Attributes}}
### {{.Name}} Attribute
{{if .Description}}{{.Description}}{{end}}
- **Type**: {{.Type}}{{if .Default}}
- **Default**: {{.Default}}{{end}}{{if .Required}}
- **Required**: Yes{{end}}{{if .Values}}
- **Options**: {{join .Values ", "}}{{end}}

```html
<{{$.TagName}} {{.Name}}="{{if .Default}}{{.Default}}{{else}}{{if .Values}}{{index .Values 0}}{{else}}/* your {{.Type}} value */{{end}}{{end}}"></{{$.TagName}}>
```
{{end}}
{{end}}

{{if gt (len .Slots) 0}}
## With Slots
{{range .Slots}}
### {{if .Name}}{{.Name}} Slot{{else}}Default Slot{{end}}
{{if .Description}}{{.Description}}{{end}}

```html
<{{$.TagName}}>
  {{if .Name}}<span slot="{{.Name}}">{{.Name}} content</span>{{else}}<p>Default slot content</p>{{end}}
</{{$.TagName}}>
```
{{end}}
{{end}}

{{if gt (len .Events) 0}}
## Event Handling
{{range .Events}}
### {{.Name}} Event
{{if .Description}}{{.Description}}{{end}}
{{if .Type}}- **Type**: {{.Type}}{{end}}

```javascript
element.addEventListener('{{.Name}}', (event) => {
  // Handle {{.Name}} event
  console.log('{{.Name}} fired:', event.detail);
});
```
{{end}}
{{end}}

{{if gt (len .CssProperties) 0}}
## CSS Custom Properties
{{range .CssProperties}}
### {{.Name}}
{{if .Description}}{{.Description}}{{end}}
{{if .Syntax}}- **Syntax**: {{.Syntax}}{{end}}
{{if .Initial}}- **Initial**: {{.Initial}}{{end}}

```css
{{$.TagName}} {
  {{.Name}}: {{if .Initial}}{{.Initial}}{{else}}/* your {{.Syntax}} value */{{end}};
}
```
{{end}}
{{end}}

{{if gt (len .CssParts) 0}}
## CSS Parts
{{range .CssParts}}
### {{.Name}} Part
{{if .Description}}{{.Description}}{{end}}

```css
{{$.TagName}}::part({{.Name}}) {
  /* Style the {{.Name}} part */
  color: /* your color value */;
}
```
{{end}}
{{end}}

{{if gt (len .CssStates) 0}}
## CSS Custom States
{{range .CssStates}}
### {{.Name}} State
{{if .Description}}{{.Description}}{{end}}

```css
{{$.TagName}}:--{{.Name}} {
  /* Styles when {{$.TagName}} is in {{.Name}} state */
  opacity: /* your opacity value */;
}
```
{{end}}
{{end}}