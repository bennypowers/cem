# HTML Generation for `button-element`

**What is a Custom Element:** A description of a custom element class.

Custom elements are JavaScript classes, so this extends from `ClassDeclaration` and adds custom-element-specific features like attributes, events, and slots.

Note that `tagName` in this interface is optional. Tag names are not neccessarily part of a custom element class, but belong to the definition (often called the "registration") or the `customElements.define()` call.

Because classes and tag names can only be registered once, there's a one-to-one relationship between classes and tag names. For ease of use, we allow the tag name here.

Some packages define and register custom elements in separate modules. In these cases one `Module` should contain the `CustomElement` without a tagName, and another `Module` should contain the `CustomElementExport`.

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
Attributes can have the following fields: **Type** (The type that the attribute will be serialized/deserialized as.), **Default** (The default value of the attribute, if any.

As attributes are always strings, this is the actual value, not a human readable description.).


#### `variant`
**Description:** Button variant
**Type:** `"primary" | "secondary" | "ghost"`
**Default:** `"primary"`
**Valid Values:** `"primary"`, `"secondary"`, `"ghost"`


#### `size`
**Description:** Button size
**Type:** `"small" | "medium" | "large"`
**Default:** `"medium"`
**Valid Values:** `"small"`, `"medium"`, `"large"`


#### `disabled`
**Description:** Whether button is disabled
**Type:** `boolean`




### Slots Available
#### Default slot
**Description:** Button content
**Usage:** Content placed directly inside the element


#### `icon` slot
**Description:** Button icon
**Usage:** `<element slot="icon">content</element>`




### Events Available
#### `button-click`
**Description:** Button click event
**Type:** `CustomEvent`



## Implementation Notes

### Best Practices
- Use semantic HTML within slots for better accessibility
- Respect the element's intended purpose and constraints
- Check element documentation for specific usage guidelines

### Styling
This element supports 2 CSS custom properties for theming.

**Available CSS Custom Properties:**
- **`--button-color`**: Button color (Syntax: `<color>`) — Default: `blue`
- **`--button-padding`**: Button padding (Syntax: `<length>`) — Default: `8px`


Use the `suggest_css_integration` tool for detailed styling guidance.

### Accessibility
- This element may implement accessibility features via ElementInternals
- Test with assistive technologies to verify accessibility implementation  
- Avoid adding conflicting ARIA attributes without understanding the element's built-in semantics
- Check shadow DOM implementation for accessibility patterns