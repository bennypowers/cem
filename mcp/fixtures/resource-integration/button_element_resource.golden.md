# button-element Usage Examples

## Basic Usage
```html
<button-element></button-element>
```




## With Attributes

### variant Attribute
Button variant
- **Type**: "primary" | "secondary" | "ghost"
- **Default**: "primary"
- **Options**: "primary", "secondary", "ghost"

```html
<button-element variant="primary"></button-element>
```

### size Attribute
Button size
- **Type**: "small" | "medium" | "large"
- **Default**: "medium"
- **Options**: "small", "medium", "large"

```html
<button-element size="medium"></button-element>
```

### disabled Attribute
Whether button is disabled
- **Type**: boolean

```html
<button-element disabled></button-element>
```




## With Slots

### Default Slot
Button content

```html
<button-element>
  <p>Default slot content</p>
</button-element>
```

### icon Slot
Button icon

```html
<button-element>
  <span slot="icon">icon content</span>
</button-element>
```




## Event Handling

### button-click Event
Button click event
- **Type**: CustomEvent

```javascript
element.addEventListener('button-click', (event) => {
  // Handle button-click event
  console.log('button-click fired:', event.detail);
});
```




## CSS Custom Properties

### --button-color
Button color
- **Syntax**: `<color>`
- **Initial**: blue

```css
button-element {
  --button-color: blue;
}
```

### --button-padding
Button padding
- **Syntax**: `<length>`
- **Initial**: 8px

```css
button-element {
  --button-padding: 8px;
}
```




## CSS Parts

### button Part
The button element

```css
button-element::part(button) {
  /* Style the button part */
  color: /* your color value */;
}
```



