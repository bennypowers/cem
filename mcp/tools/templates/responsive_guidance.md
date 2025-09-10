## 📱 Responsive Design

### Container Queries (Recommended)
Use container queries to make the element responsive to its container size:

```css
{{.TagName}} {
  container-type: inline-size;
}

/* Small container */
@container (max-width: 400px) {
  {{.TagName}} {
    /* Compact styles */
  }
}

/* Medium container */
@container (min-width: 401px) and (max-width: 800px) {
  {{.TagName}} {
    /* Standard styles */
  }
}

/* Large container */
@container (min-width: 801px) {
  {{.TagName}} {
    /* Expanded styles */
  }
}
```

### Media Queries (Fallback)
For broader browser support, use traditional media queries:

```css
/* Mobile first approach */
{{.TagName}} {
  /* Mobile styles */
}

@media (min-width: 768px) {
  {{.TagName}} {
    /* Tablet styles */
  }
}

@media (min-width: 1024px) {
  {{.TagName}} {
    /* Desktop styles */
  }
}
```

### Responsive Considerations
{{if .CssProperties}}- Adjust CSS custom properties for different screen sizes
{{range .CssProperties}}- `{{.Name}}`: Consider responsive values{{if .Description}} ({{.Description}}){{end}}
{{end}}

{{end}}{{if .CssParts}}- Style parts differently at various breakpoints:
{{range .CssParts}}- `{{.Name}}` part: {{if .Description}}{{.Description}}{{else}}Consider responsive layout changes{{end}}
{{end}}

{{end}}- Test across different device sizes and orientations
- Consider touch targets and interaction methods
- Ensure content remains accessible at all sizes