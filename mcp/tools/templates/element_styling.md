# Styling Reference: `{{.Element.TagName}}`

{{if .Element.Description}}{{.Element.Description}}{{end}}

{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}
{{if gt (len .Element.CssProperties) 0}}
## CSS Custom Properties

{{schemaDesc .SchemaDefinitions "CssCustomProperty"}}

{{range .Element.CssProperties}}
### `{{.Name}}`

{{if .Description}}{{.Description}}{{end}}

**Syntax:** `{{.Syntax}}`{{if .Initial}} | **Initial:** `{{.Initial}}`{{end}}

{{if .Guidelines}}
**Guidelines:**
{{range .Guidelines}}- {{.}}
{{end}}{{end}}

{{if .Examples}}
**Examples:**
{{range .Examples}}- {{.}}
{{end}}{{end}}

{{end}}

### Usage Examples

```css
/* Basic theming */
{{.Element.TagName}} {
{{range .Element.CssProperties}}  {{.Name}}: /* your value ({{.Syntax}}) */;
{{end}}}

/* With custom values */
{{.Element.TagName}} {
{{range .Element.CssProperties}}  {{.Name}}: {{if .Initial}}{{.Initial}}{{else}}/* your value */{{end}};
{{end}}}
```

{{end}}

{{if gt (len .Element.CssParts) 0}}
## CSS Parts

{{schemaDesc .SchemaDefinitions "CssPart"}}

{{range .Element.CssParts}}
### `{{.Name}}` part

{{if .Description}}{{.Description}}{{end}}

{{if .Guidelines}}
**Guidelines:**
{{range .Guidelines}}- {{.}}
{{end}}{{end}}

{{if .Examples}}
**Examples:**
{{range .Examples}}- {{.}}
{{end}}{{end}}

{{end}}

### Usage Examples

```css
/* Style individual parts */
{{range .Element.CssParts}}{{$.Element.TagName}}::part({{.Name}}) {
  /* {{if .Description}}{{.Description}}{{else}}Style the {{.Name}} part{{end}} */
}

{{end}}
/* Multiple parts styling */
{{.Element.TagName}}::part({{range $i, $part := .Element.CssParts}}{{if $i}}, {{end}}{{$part.Name}}{{end}}) {
  font-family: var(--font-family);
}
```

{{end}}

{{if gt (len .Element.CssStates) 0}}
## CSS Custom States

{{schemaDesc .SchemaDefinitions "CssCustomState"}}

{{range .Element.CssStates}}
### `{{.Name}}` state

{{if .Description}}{{.Description}}{{end}}

{{if .Guidelines}}
**Guidelines:**
{{range .Guidelines}}- {{.}}
{{end}}{{end}}

{{if .Examples}}
**Examples:**
{{range .Examples}}- {{.}}
{{end}}{{end}}

{{end}}

### Usage Examples

```css
/* Style based on custom states */
{{range .Element.CssStates}}{{$.Element.TagName}}:state({{.Name}}) {
  /* {{if .Description}}{{.Description}}{{else}}Style when {{.Name}} state is active{{end}} */
}

{{end}}
/* Combined state styling */
{{.Element.TagName}}:state({{range $i, $state := .Element.CssStates}}{{if $i}}, {{end}}{{$state.Name}}{{end}}) {
  transition: all 0.2s ease;
}
```

{{end}}

## Complete Styling Example

```css
/* Complete styling with all available APIs */
{{.Element.TagName}} {
{{range .Element.CssProperties}}  {{.Name}}: {{if .Initial}}{{.Initial}}{{else}}/* your value */{{end}};
{{end}}}

{{if gt (len .Element.CssParts) 0}}
{{range .Element.CssParts}}/* Style {{.Name}} part */
{{$.Element.TagName}}::part({{.Name}}) {
  /* {{if .Description}}{{.Description}}{{else}}Custom styling for {{.Name}}{{end}} */
}

{{end}}{{end}}
{{if gt (len .Element.CssStates) 0}}
{{range .Element.CssStates}}/* Style {{.Name}} state */
{{$.Element.TagName}}:state({{.Name}}) {
  /* {{if .Description}}{{.Description}}{{else}}Style when {{.Name}} state is active{{end}} */
}

{{end}}{{end}}
```

{{else}}
## No CSS Customization APIs Available

This element does not expose CSS custom properties, parts, or states for styling. It may rely on standard CSS selectors and properties.

## Standard CSS Styling

```css
/* Standard CSS styling */
{{.Element.TagName}} {
  display: block;
  /* Apply standard CSS properties */
}

/* Pseudo-class styling */
{{.Element.TagName}}:hover {
  /* Hover styles */
}

{{.Element.TagName}}:focus {
  /* Focus styles */
}
```
{{end}}

---

For related API information, use:
{{if gt (len .Element.Attributes) 0}}- **`element_attributes`** - Detailed attribute documentation and usage{{end}}
{{if gt (len .Element.Slots) 0}}- **`element_slots`** - Slot usage patterns and content guidelines{{end}}
{{if gt (len .Element.Events) 0}}- **`element_events`** - Event details and JavaScript integration{{end}}
- **`element_details`** - Complete element reference with all APIs
- **`suggest_css_integration`** - Design system integration guidance
- **`generate_html`** - HTML generation with proper structure