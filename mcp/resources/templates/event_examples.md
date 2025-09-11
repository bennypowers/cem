# Event Handling Context

{{range .Events}}
## {{.Name}} Event
{{if .Description}}{{.Description}}{{end}}

**Event Details:**
{{if .Type}}- **Type**: {{.Type}}{{end}}
- **Bubbles**: Check element documentation
- **Cancelable**: Check element documentation

**Event Handling Context:**
```javascript
// Event listener approach
element.addEventListener('{{.Name}}', (event) => {
  // Access event details through event.detail
  console.log('{{.Name}} event:', event.detail);
  
  // Your event handling logic here
});

// One-time listener
element.addEventListener('{{.Name}}', (event) => {
  // Handle once
}, { once: true });

// Passive listener for performance
element.addEventListener('{{.Name}}', (event) => {
  // Non-blocking handler
}, { passive: true });
```

**Framework Integration:**
```html
<!-- Lit -->
<element @{{.Name}}="${this.handle{{title .Name}}}"></element>

<!-- React -->
<Element on{{title .Name}}={this.handle{{title .Name}}} />

<!-- Vue -->
<element @{{.Name}}="handle{{title .Name}}"></element>
```

{{if .Guidelines}}
**Event Guidelines:**
{{range .Guidelines}}
- {{.}}{{end}}
{{end}}
{{end}}