# Package: `{{.Args.package}}`

{{range .FetchedData.package.packages}}
{{if eq .name $.Args.package}}
## Package Information

**Elements:** {{.elementCount}} | **Modules:** {{.moduleCount}}

## Elements in Package

| Element | Module | Summary |
| ------- | ------ | ------- |
{{range .elements}}{{$tagName := .}}{{range $.FetchedData.elements.elements}}{{if eq .tagName $tagName}}| `{{.tagName}}` | `{{.module}}` | {{if .summary}}{{.summary}}{{else if .description}}{{.description}}{{else}}-{{end}} |
{{end}}{{end}}{{end}}

## Module Structure

{{if gt .moduleCount 0}}
{{range .modules}}
- `{{.}}`
{{end}}
{{else}}
No modules defined for this package.
{{end}}

## Package Usage

### Import Patterns
```javascript
{{range .modules}}
// Import from {{.}}
import './{{.}}';
{{end}}
```

### Element Usage
```html
{{range .elements}}
<!-- {{.}} -->
<{{.}}></{{.}}>
{{end}}
```

---

For detailed element information within this package, use:
{{range .elements}}- **`cem://element/{{.}}`** - Complete reference for {{.}}
{{end}}

{{else}}
## Package Not Found

Package `{{$.Args.package}}` was not found in the workspace.

Available packages:
{{range $.FetchedData.package.packages}}- `{{.name}}`
{{end}}
{{end}}
{{end}}