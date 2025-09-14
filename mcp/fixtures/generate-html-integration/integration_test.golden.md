# HTML Generation for `button-element`

## Generated HTML

```html
<button-element>
  Test Button
  <span slot="icon">Icon content</span>
</button-element>
```

## Element Structure and Guidelines

### Tag Information
- **Element:** `button-element`


### Attributes Available

#### `variant`
Button variant
**Type:** `"primary" | "secondary" | "ghost"`
**Default:** `"primary"`
**Valid Values:** `"primary"`, `"secondary"`, `"ghost"`

#### `size`
Button size
**Type:** `"small" | "medium" | "large"`
**Default:** `"medium"`
**Valid Values:** `"small"`, `"medium"`, `"large"`

#### `disabled`
Whether button is disabled
**Type:** `boolean`



### Slots Available

#### Default slot
Button content
**Usage:** Content placed directly inside the element

#### `icon` slot
Button icon
**Usage:** `<element slot="icon">content</element>`



### Events Available

#### `button-click`
Button click event



## Implementation Notes

### Best Practices
- Use semantic HTML within slots for better accessibility
- Respect the element's intended purpose and constraints
- Check element documentation for specific usage guidelines

### Styling
This element supports 2 CSS custom properties for theming. Use the `suggest_css_integration` tool for styling guidance.

### Accessibility
- This element may implement accessibility features via ElementInternals
- Test with assistive technologies to verify accessibility implementation  
- Avoid adding conflicting ARIA attributes without understanding the element's built-in semantics
- Check shadow DOM implementation for accessibility patterns