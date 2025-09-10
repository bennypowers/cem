## 📋 Usage Examples

### Basic Usage
```html
<{{.TagName}}{{range .Attributes}}{{if .Required}}{{if .Default}} {{.Name}}="{{.Default}}"{{else if gt (len .Values) 0}} {{.Name}}="{{index .Values 0}}"{{else}} {{.Name}}="value"{{end}}{{end}}{{end}}></{{.TagName}}>
```

{{if eq .Context "accessibility"}}### Accessible Implementation
```html
<{{.TagName}} role="button" aria-label="Accessible {{.TagName}}" tabindex="0">
  Content
</{{.TagName}}>
```

{{end}}{{if eq .Context "form"}}### Form Integration
```html
<form>
  <label for="my-{{.TagName}}">Label:</label>
  <{{.TagName}} id="my-{{.TagName}}" name="field-name" required>
    Content
  </{{.TagName}}>
</form>
```

{{end}}{{if gt (len .Attributes) 1}}### Full Configuration
```html
<{{.TagName}}{{range .Attributes}}{{if .Default}} {{.Name}}="{{.Default}}"{{else if gt (len .Values) 0}} {{.Name}}="{{index .Values 0}}"{{end}}{{end}}></{{.TagName}}>
```
{{end}}