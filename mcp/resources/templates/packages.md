# Packages Overview

## Summary
{{.FetchedData.packages.metadata.totalPackages}} packages containing {{.FetchedData.packages.metadata.totalElements}} total elements

## Available Packages

| Package | Elements | Modules | Element List |
| ------- | -------- | ------- | ------------ |
{{range .FetchedData.packages.packages}}| `{{.name}}` | {{.elementCount}} | {{.moduleCount}} | {{range $i, $el := .elements}}{{if $i}}, {{end}}`{{$el}}`{{end}} |
{{end}}

## Package Discovery

This overview helps you understand the component ecosystem in your workspace:

- **Package organization** - See how elements are grouped and distributed
- **Element discovery** - Find components by package or module
- **Dependency analysis** - Understand package relationships and imports
- **Development planning** - Identify areas for component development

## Usage Patterns

### Import Resolution
Elements are typically imported from their module paths. Use the module information above to construct proper import statements.

### Package Dependencies
Some packages may depend on others. Consider package relationships when planning component usage and development.

### Component Selection
Choose elements from packages that align with your design system and architectural patterns.

---

For detailed package exploration, use:
- **`cem://elements`** - Browse all elements with capabilities
- **`cem://element/{tagName}`** - Detailed information for specific elements
- **`cem://schema`** - Understand manifest structure and validation