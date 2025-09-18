# Events from `{{.Element.TagName}}` 📡

{{if .Element.Description}}{{.Element.Description}}

{{end}}{{if gt (len .Element.Events) 0}}This element fires **{{len .Element.Events}} different events** to keep you informed about what's happening. Listen in to react to user interactions and state changes!

{{schemaDesc .SchemaDefinitions "Event"}}

## What's Broadcasting 📻

| Event | When it fires | Details |
| ----- | ------------- | ------- |
{{range .Element.Events}}| `{{.Name}}` | {{if .Summary}}{{.Summary}}{{else if .Description}}{{.Description}}{{else}}When something happens{{end}} | {{if .Type}}{{.Type.Text}}{{else}}Event data{{end}} |
{{end}}

## The Full Story 📖

{{range .Element.Events}}
### `{{.Name}}` event

{{if .Summary}}{{.Summary}}{{end}}
{{if .Description}}

{{.Description}}{{end}}

**Triggers:** User interactions and state changes
**Event data:** {{if .Type}}`{{.Type.Text}}`{{else}}Standard event object{{end}}

{{end}}

## Hook Into the Action 🎣

```javascript
// Listen for events
const element = document.querySelector('{{.Element.TagName}}');

{{range .Element.Events}}
// React to {{.Name}}
element.addEventListener('{{.Name}}', (event) => {
  console.log('{{.Name}} fired!', event.detail);
  // Your logic here
});

{{end}}
```

### 🚫 Event Listening Pitfalls
- **Don't** forget to remove event listeners when elements are destroyed
- **Don't** assume event data structure without checking the manifest
- **Don't** listen for events that don't exist (check spelling!)
- **Don't** ignore browser support for custom events in older browsers

{{else}}
## Quiet Element 🤫

This element doesn't fire any custom events — it keeps things simple and lets the standard HTML events do the talking.

```javascript
// You can still listen for standard events
const element = document.querySelector('{{.Element.TagName}}');
element.addEventListener('click', () => {
  console.log('Clicked!');
});
```
{{end}}

## Keep Exploring 🌟

{{if gt (len .Element.Attributes) 0}}• **[Attributes](cem://element/{{.Element.TagName}}/attributes)** — Control the behavior{{end}}
{{if gt (len .Element.Slots) 0}}• **[Content Slots](cem://element/{{.Element.TagName}}/slots)** — Fill with your HTML{{end}}
{{if or (gt (len .Element.CssProperties) 0) (gt (len .Element.CssParts) 0) (gt (len .Element.CssStates) 0)}}• **[Styling](cem://element/{{.Element.TagName}}/css)** — Customize the appearance{{end}}
• **[Full Reference](cem://element/{{.Element.TagName}})** — Complete element guide
• **[Generate HTML](cem://tools/generate_html)** — Get markup with proper event handling
