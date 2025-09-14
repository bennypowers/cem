# HTML Generation for `button-element`

## Generated HTML

```html
&amp;lt;button-element&gt;
  Test Button
  &lt;span slot=&#34;icon&#34;&gt;Icon content&lt;/span&gt;
&amp;lt;/button-element&gt;
```

## Element Structure and Guidelines

### Tag Information
- **Element:** `button-element`


### Attributes Available

#### `variant`
Button variant
**Type:** `&#34;primary&#34; | &#34;secondary&#34; | &#34;ghost&#34;`
**Default:** `&#34;primary&#34;`
**Valid Values:** `&#34;primary&#34;`, `&#34;secondary&#34;`, `&#34;ghost&#34;`

#### `size`
Button size
**Type:** `&#34;small&#34; | &#34;medium&#34; | &#34;large&#34;`
**Default:** `&#34;medium&#34;`
**Valid Values:** `&#34;small&#34;`, `&#34;medium&#34;`, `&#34;large&#34;`

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