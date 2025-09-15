# Events Reference: `{{.Element.TagName}}`

{{if .Element.Description}}{{.Element.Description}}{{end}}

{{if gt (len .Element.Events) 0}}
## Available Events

{{schemaDesc .SchemaDefinitions "Event"}}

{{range .Element.Events}}
### `{{.Name}}`

{{if .Description}}{{.Description}}{{end}}

**Type:** `{{.Type}}`

{{if .Guidelines}}
**Guidelines:**
{{range .Guidelines}}- {{.}}
{{end}}{{end}}

{{if .Examples}}
**Examples:**
{{range .Examples}}- {{.}}
{{end}}{{end}}

{{end}}

## JavaScript Integration

### Event Listeners

```javascript
// Get reference to the element
const element = document.querySelector('{{.Element.TagName}}');

{{range .Element.Events}}
// Listen for {{.Name}} event
element.addEventListener('{{.Name}}', (event) => {
  console.log('{{.Name}} fired:', event.detail);
  {{if .Description}}// {{.Description}}{{end}}
});

{{end}}
```

### Event Handler Functions

```javascript
{{range .Element.Events}}
function handle{{.Name}}Event(event) {
  {{if .Description}}// {{.Description}}{{end}}
  const data = event.detail;
  // Handle the {{.Name}} event
  console.log('Event data:', data);
}

{{end}}
// Attach event handlers
const element = document.querySelector('{{.Element.TagName}}');
{{range .Element.Events}}element.addEventListener('{{.Name}}', handle{{.Name}}Event);
{{end}}
```

### React Integration

```jsx
import React from 'react';

function MyComponent() {
  {{range .Element.Events}}
  const handle{{.Name}}Event = (event) => {
    {{if .Description}}// {{.Description}}{{end}}
    console.log('{{.Name}}:', event.detail);
  };

  {{end}}
  return (
    <{{.Element.TagName}}
      {{range .Element.Events}}on{{.Name}}={handle{{.Name}}Event}
      {{end}}
    >
      Content
    </{{.Element.TagName}}>
  );
}
```

{{else}}
## No Custom Events Available

This element does not emit any custom events. It may use standard DOM events only.

## Standard DOM Events

```javascript
// Standard DOM events still work
const element = document.querySelector('{{.Element.TagName}}');

element.addEventListener('click', (event) => {
  console.log('Element clicked');
});

element.addEventListener('focus', (event) => {
  console.log('Element focused');
});
```
{{end}}

---

For related API information, use:
{{if gt (len .Element.Attributes) 0}}- **`element_attributes`** - Detailed attribute documentation and usage{{end}}
{{if gt (len .Element.Slots) 0}}- **`element_slots`** - Slot usage patterns and content guidelines{{end}}
{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}- **`element_styling`** - CSS customization with properties, parts, and states{{end}}
- **`element_details`** - Complete element reference with all APIs
- **`generate_html`** - HTML generation with proper structure