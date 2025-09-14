# HTML Generation for `{{.Element.TagName}}`

{{if .Element.Description}}**Element Description:** {{.Element.Description}}
{{end}}
## Generated HTML

```html
{{.GeneratedHTML}}
```

## Element Structure and Guidelines

### Tag Information
- **Element:** `{{.Element.TagName}}`{{if .Element.Package}}
- **Package:** {{.Element.Package}}{{end}}
{{if .Element.Description}}
- **Purpose:** {{.Element.Description}}{{end}}

{{if gt (len .Element.Attributes) 0}}### Attributes Available

{{- $attributeSchema := schemaDesc $.SchemaDefinitions "Attribute"}}
{{- $typeSchema := schemaFieldDesc $.SchemaDefinitions "Attribute" "type"}}
{{- $defaultSchema := schemaFieldDesc $.SchemaDefinitions "Attribute" "default"}}
{{- $nameSchema := schemaFieldDesc $.SchemaDefinitions "Attribute" "name"}}
{{- $descriptionSchema := schemaFieldDesc $.SchemaDefinitions "Attribute" "description"}}

**What are Attributes?** {{if $attributeSchema}}{{$attributeSchema}}{{else}}Attributes are properties that can be set on the custom element via HTML attributes or JavaScript properties.{{end}}

**Understanding Attribute Fields:**
{{if $nameSchema}}- **Name** ({{$nameSchema}}){{end}}
{{if $descriptionSchema}}- **Description** ({{$descriptionSchema}}){{end}}
{{if $typeSchema}}- **Type** ({{$typeSchema}}){{end}}
{{if $defaultSchema}}- **Default** ({{$defaultSchema}}){{end}}

{{end}}
{{range .Element.Attributes}}#### `{{.Name}}`{{if .Required}} ⚠️ **Required**{{end}}
{{if .Description}}{{.Description}}
{{end}}{{if .Type}}**Type:** `{{.Type}}`
{{end}}{{if .Default}}**Default:** `{{.Default}}`
{{end}}{{if gt (len .Values) 0}}**Valid Values:** {{range $i, $v := .Values}}{{if $i}}, {{end}}`{{$v}}`{{end}}
{{end}}

{{end}}{{end}}

{{if gt (len .Element.Slots) 0}}### Slots Available

{{- $slotSchema := schemaDesc $.SchemaDefinitions "Slot"}}
{{- $slotNameSchema := schemaFieldDesc $.SchemaDefinitions "Slot" "name"}}
{{- $slotDescriptionSchema := schemaFieldDesc $.SchemaDefinitions "Slot" "description"}}

**What are Slots?** {{if $slotSchema}}{{$slotSchema}}{{else}}Slots are placeholders where users can insert their own content into the component. They allow custom elements to accept and position user-provided content within their shadow DOM structure.{{end}}

**Understanding Slot Fields:**
{{if $slotNameSchema}}- **Name** ({{$slotNameSchema}}){{end}}
{{if $slotDescriptionSchema}}- **Description** ({{$slotDescriptionSchema}}){{end}}
- **Usage** (How to target this slot with content)

{{end}}
{{range .Element.Slots}}#### {{if .Name}}`{{.Name}}` slot{{else}}Default slot{{end}}
{{if .Description}}{{.Description}}
{{end}}{{if .Summary}}**Summary:** {{.Summary}}
{{end}}{{if .Deprecated}}⚠️ **Deprecated**{{if ne .Deprecated true}}: {{.Deprecated}}{{end}}
{{end}}{{if .InheritedFrom}}**Inherited From:** {{.InheritedFrom.name}}
{{end}}{{if .Name}}**Usage:** `<element slot="{{.Name}}">content</element>`
{{else}}**Usage:** Content placed directly inside the element
{{end}}**Guidelines:** Follow semantic HTML practices and maintain accessibility within this slot

{{end}}{{end}}

{{if gt (len .Element.Events) 0}}### Events Available

{{- $eventSchema := schemaDesc $.SchemaDefinitions "Event"}}
{{- $eventNameSchema := schemaFieldDesc $.SchemaDefinitions "Event" "name"}}
{{- $eventDescriptionSchema := schemaFieldDesc $.SchemaDefinitions "Event" "description"}}
{{- $eventTypeSchema := schemaFieldDesc $.SchemaDefinitions "Event" "type"}}

**What are Events?** {{if $eventSchema}}{{$eventSchema}}{{else}}Events are dispatched by the custom element to notify external code of state changes, user interactions, or important lifecycle moments. They enable communication between the component and the surrounding application.{{end}}

**Understanding Event Fields:**
{{if $eventNameSchema}}- **Name** ({{$eventNameSchema}}){{end}}
{{if $eventDescriptionSchema}}- **Description** ({{$eventDescriptionSchema}}){{end}}
{{if $eventTypeSchema}}- **Type** ({{$eventTypeSchema}}){{end}}
- **JavaScript** (How to listen for this event)

{{end}}
{{range .Element.Events}}#### `{{.Name}}`
{{if .Description}}{{.Description}}
{{end}}{{if .Summary}}**Summary:** {{.Summary}}
{{end}}{{if .Type}}**Type:** `{{.Type}}`
{{end}}{{if .Deprecated}}⚠️ **Deprecated**{{if ne .Deprecated true}}: {{.Deprecated}}{{end}}
{{end}}{{if .InheritedFrom}}**Inherited From:** {{.InheritedFrom.name}}
{{end}}**JavaScript:** `element.addEventListener('{{.Name}}', (event) => { /* handle event */ });`

{{end}}{{end}}

## Implementation Notes

### Best Practices
- Use semantic HTML within slots for better accessibility
- Respect the element's intended purpose and constraints
- Check element documentation for specific usage guidelines
{{if gt (len .Element.CssProperties) 0}}
### Styling
This element supports {{len .Element.CssProperties}} CSS custom properties for theming.

{{- $cssPropertySchema := schemaDesc $.SchemaDefinitions "CssCustomProperty"}}
{{- $cssPropertyNameSchema := schemaFieldDesc $.SchemaDefinitions "CssCustomProperty" "name"}}
{{- $cssPropertyDescriptionSchema := schemaFieldDesc $.SchemaDefinitions "CssCustomProperty" "description"}}
{{- $cssPropertySyntaxSchema := schemaFieldDesc $.SchemaDefinitions "CssCustomProperty" "syntax"}}
{{- $cssPropertyInitialSchema := schemaFieldDesc $.SchemaDefinitions "CssCustomProperty" "initial"}}

**What are CSS Custom Properties?** {{if $cssPropertySchema}}{{$cssPropertySchema}}{{else}}CSS custom properties (also known as CSS variables) are properties defined by the component that allow external CSS to customize the component's appearance and behavior. They provide a controlled interface for theming and styling customization.{{end}}

**Understanding CSS Property Fields:**
{{if $cssPropertyNameSchema}}- **Name** ({{$cssPropertyNameSchema}}){{end}}
{{if $cssPropertyDescriptionSchema}}- **Description** ({{$cssPropertyDescriptionSchema}}){{end}}
{{if $cssPropertySyntaxSchema}}- **Syntax** ({{$cssPropertySyntaxSchema}}){{end}}
{{if $cssPropertyInitialSchema}}- **Default Value** ({{$cssPropertyInitialSchema}}){{end}}
- **CSS Usage** (How to use this property in your stylesheets)

**Available CSS Custom Properties:**
{{range .Element.CssProperties}}#### `{{.Name}}`
{{if .Description}}{{.Description}}
{{end}}{{if .Summary}}**Summary:** {{.Summary}}
{{end}}{{if .Syntax}}**Syntax:** `{{.Syntax}}`
{{end}}{{if .Initial}}**Default Value:** `{{.Initial}}`
{{end}}{{if .Inherits}}**Inherits:** Yes
{{end}}{{if .Deprecated}}⚠️ **Deprecated**{{if ne .Deprecated true}}: {{.Deprecated}}{{end}}
{{end}}{{if .InheritedFrom}}**Inherited From:** {{.InheritedFrom.name}}
{{end}}**CSS Usage:** `selector { {{.Name}}: /* your value */; }`

{{end}}

Use the `suggest_css_integration` tool for detailed styling guidance.
{{end}}

{{if gt (len .Element.CssParts) 0}}### CSS Parts Available

{{- $cssPartSchema := schemaDesc $.SchemaDefinitions "CssPart"}}
{{- $cssPartNameSchema := schemaFieldDesc $.SchemaDefinitions "CssPart" "name"}}
{{- $cssPartDescriptionSchema := schemaFieldDesc $.SchemaDefinitions "CssPart" "description"}}

**What are CSS Parts?** {{if $cssPartSchema}}{{$cssPartSchema}}{{else}}CSS parts expose internal elements of the component's shadow DOM for external styling. They provide controlled access points where external CSS can apply styles to specific parts of the component's internal structure using the ::part() pseudo-element.{{end}}

**Understanding CSS Part Fields:**
{{if $cssPartNameSchema}}- **Name** ({{$cssPartNameSchema}}){{end}}
{{if $cssPartDescriptionSchema}}- **Description** ({{$cssPartDescriptionSchema}}){{end}}
- **CSS Usage** (How to style this part with ::part() pseudo-element)

**Available CSS Parts:**
{{range .Element.CssParts}}#### `{{.Name}}`
{{if .Description}}{{.Description}}
{{end}}{{if .Summary}}**Summary:** {{.Summary}}
{{end}}{{if .Deprecated}}⚠️ **Deprecated**{{if ne .Deprecated true}}: {{.Deprecated}}{{end}}
{{end}}{{if .InheritedFrom}}**Inherited From:** {{.InheritedFrom.name}}
{{end}}**CSS Usage:** `{{$.Element.TagName}}::part({{.Name}}) { /* your styles */ }`

{{end}}
{{end}}

{{if gt (len .Element.CssStates) 0}}### CSS States Available

{{- $cssStateSchema := schemaDesc $.SchemaDefinitions "CssCustomState"}}
{{- $cssStateNameSchema := schemaFieldDesc $.SchemaDefinitions "CssCustomState" "name"}}
{{- $cssStateDescriptionSchema := schemaFieldDesc $.SchemaDefinitions "CssCustomState" "description"}}

**What are CSS Custom States?** {{if $cssStateSchema}}{{$cssStateSchema}}{{else}}CSS custom states are user-defined states that represent the component's current condition or mode. They allow external CSS to style the component differently based on its internal state (like loading, expanded, invalid, etc.) using custom state pseudo-classes.{{end}}

**Understanding CSS State Fields:**
{{if $cssStateNameSchema}}- **Name** ({{$cssStateNameSchema}}){{end}}
{{if $cssStateDescriptionSchema}}- **Description** ({{$cssStateDescriptionSchema}}){{end}}
- **CSS Usage** (How to style based on this state using :--state pseudo-class)

**Available CSS States:**
{{range .Element.CssStates}}#### `:--{{.Name}}`
{{if .Description}}{{.Description}}
{{end}}{{if .Summary}}**Summary:** {{.Summary}}
{{end}}{{if .Deprecated}}⚠️ **Deprecated**{{if ne .Deprecated true}}: {{.Deprecated}}{{end}}
{{end}}{{if .InheritedFrom}}**Inherited From:** {{.InheritedFrom.name}}
{{end}}**CSS Usage:** `{{$.Element.TagName}}:--{{.Name}} { /* styles when in this state */ }`

{{end}}
{{end}}

### Accessibility
- This element may implement accessibility features via ElementInternals
- Test with assistive technologies to verify accessibility implementation  
- Avoid adding conflicting ARIA attributes without understanding the element's built-in semantics
- Check shadow DOM implementation for accessibility patterns
