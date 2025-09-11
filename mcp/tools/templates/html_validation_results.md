# HTML Validation Results

{{if gt (len .FoundElements) 0}}## Custom Elements Found

{{range .FoundElements}}- **`{{.TagName}}`**: {{if .Description}}{{.Description}}{{else}}Custom element{{end}}
{{end}}
{{end}}{{if eq .Context "semantic"}}## Semantic Structure Validation

{{if gt (len .SemanticIssues) 0}}### ❌ Semantic Issues

{{range .SemanticIssues}}{{if eq .Priority "error"}}- **{{.Message}}**
{{else}}- {{.Message}}
{{end}}{{end}}
{{end}}{{if gt (len .SemanticSuggestions) 0}}### 💡 Semantic Suggestions

{{range .SemanticSuggestions}}{{if eq .Type "semantic-elements"}}- **Semantic Elements**: {{.Message}}
{{else if eq .Type "list-structure"}}- **List Structure**: {{.Message}}
{{else if eq .Type "document-structure"}}- **Document Structure**: {{.Message}}
{{else}}- **{{.Type}}**: {{.Message}}
{{end}}{{end}}
{{end}}{{if eq (len .SemanticIssues) 0}}{{if eq (len .SemanticSuggestions) 0}}### ✅ Good Semantic Structure

{{end}}{{end}}{{else if eq .Context "manifest-compliance"}}## Manifest Compliance Validation

{{if gt (len .ManifestIssues) 0}}### ❌ Manifest Compliance Issues

{{range .ManifestIssues}}{{if eq .Type "missing-attribute"}}- **Missing Required Attribute**: `{{.Attribute}}` is required for `<{{.Element}}>`
{{else if eq .Type "invalid-value"}}- **Invalid Attribute Value**: `{{.Attribute}}="{{.Actual}}"` in `<{{.Element}}>`. {{if .Expected}}Valid values: {{.Expected}}{{end}}
{{else}}- **{{.Type}}**: {{.Message}}
{{end}}{{end}}
{{else}}### ✅ No Manifest Compliance Issues Found

{{end}}{{if gt (len .ManifestFeatures) 0}}### 💡 Manifest-Defined Features

{{range .ManifestFeatures}}{{if eq .Type "slots"}}- **`<{{.Element}}>`** supports slots: {{.Details}}
{{else if eq .Type "guidelines"}}- **`<{{.Element}}>`** guidelines: {{.Details}}
{{else if eq .Type "css-apis"}}- **`<{{.Element}}>`** supports {{.Details}} CSS customization options
{{else}}- **{{.Type}}**: {{.Details}}
{{end}}{{end}}
{{end}}{{else}}## Manifest Compliance Validation

{{if gt (len .ManifestIssues) 0}}### ❌ Manifest Compliance Issues

{{range .ManifestIssues}}{{if eq .Type "missing-attribute"}}- **Missing Required Attribute**: `{{.Attribute}}` is required for `<{{.Element}}>`
{{else if eq .Type "invalid-value"}}- **Invalid Attribute Value**: `{{.Attribute}}="{{.Actual}}"` in `<{{.Element}}>`. {{if .Expected}}Valid values: {{.Expected}}{{end}}
{{else}}- **{{.Type}}**: {{.Message}}
{{end}}{{end}}
{{else}}### ✅ No Manifest Compliance Issues Found

{{end}}{{if gt (len .ManifestFeatures) 0}}### 💡 Manifest-Defined Features

{{range .ManifestFeatures}}{{if eq .Type "slots"}}- **`<{{.Element}}>`** supports slots: {{.Details}}
{{else if eq .Type "guidelines"}}- **`<{{.Element}}>`** guidelines: {{.Details}}
{{else if eq .Type "css-apis"}}- **`<{{.Element}}>`** supports {{.Details}} CSS customization options
{{else}}- **{{.Type}}**: {{.Details}}
{{end}}{{end}}
{{end}}## Custom Element Content Validation

{{if gt (len .SlotContentIssues) 0}}### ❌ Slot Content Issues

{{range .SlotContentIssues}}- **`<{{.ElementTagName}}>`** slot `{{.SlotName}}`: {{.ViolationMessage}}
  - **Guideline**: {{.Guideline}}
{{end}}
{{end}}{{if gt (len .AttributeConflicts) 0}}### ⚠️ Attribute Conflicts

{{range .AttributeConflicts}}- **`<{{.ElementTagName}}>`**: `{{.Attribute1}}="{{.Value1}}"` conflicts with `{{.Attribute2}}="{{.Value2}}"`
  - **Reason**: {{.ConflictReason}}
{{end}}
{{end}}{{if gt (len .ContentAttributeRedundancies) 0}}### 🔄 Content/Attribute Redundancy

{{range .ContentAttributeRedundancies}}- **`<{{.ElementTagName}}>`**: Attribute `{{.AttributeName}}="{{.AttributeValue}}"` is overridden by slot content
  - **Slot**: `{{.SlotName}}` contains: `{{.SlotContent}}`
{{end}}
{{end}}{{if eq (len .SlotContentIssues) 0}}{{if eq (len .AttributeConflicts) 0}}{{if eq (len .ContentAttributeRedundancies) 0}}### ✅ Custom Element Usage Looks Good

No custom element content or attribute issues found.
{{end}}{{end}}{{end}}{{end}}{{if .SpecificElement}}## Element-Specific Validation: `{{.SpecificElement.TagName}}`

{{if .SpecificElement.ElementFound}}**Description:** {{.SpecificElement.Description}}

**Found {{len .SpecificElement.Usages}} usage(s):**

{{range $i, $usage := .SpecificElement.Usages}}{{$index := add $i 1}}{{$index}}. `{{$usage.Html}}`
{{if gt (len $usage.Issues) 0}}   - **Issues:** {{join $usage.Issues "; "}}
{{else}}   - ✅ No issues found
{{end}}{{end}}

{{if gt (len .SpecificElement.Attributes) 0}}**Available Attributes:**
{{range .SpecificElement.Attributes}}- **`{{.Name}}`**{{if .Required}} ⚠️ (required){{end}}{{if .Description}}: {{.Description}}{{end}}
{{end}}{{end}}

{{if gt (len .SpecificElement.Slots) 0}}**Available Slots:**
{{range .SpecificElement.Slots}}- {{if .Name}}**`{{.Name}}`**{{else}}**Default slot**{{end}}{{if .Description}}: {{.Description}}{{end}}
{{end}}{{end}}

{{if gt (len .SpecificElement.Events) 0}}**Available Events:**
{{range .SpecificElement.Events}}- **`{{.Name}}`**{{if .Description}}: {{.Description}}{{end}}
{{end}}{{end}}
{{else}}❌ Element `{{.SpecificElement.TagName}}` not found in HTML

{{end}}{{end}}