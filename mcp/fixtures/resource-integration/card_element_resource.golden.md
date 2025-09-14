# card-element Usage Examples

## Basic Usage
```html
&lt;card-element>&lt;/card-element>
```




## With Attributes

### elevation Attribute
Card elevation level
- **Type**: number
- **Default**: 1

```html
&lt;card-element elevation="1">&lt;/card-element>
```




## With Slots

### Default Slot
Card content

```html
&lt;card-element>
  <p>Default slot content</p>
&lt;/card-element>
```

### header Slot
Card header

```html
&lt;card-element>
  <span slot="header">header content</span>
&lt;/card-element>
```

### footer Slot
Card footer

```html
&lt;card-element>
  <span slot="footer">footer content</span>
&lt;/card-element>
```






## CSS Custom Properties

### --card-background
Card background color
- **Syntax**: &lt;color&gt;
- **Initial**: white

```css
card-element {
  --card-background: white;
}
```




## CSS Parts

### container Part
Card container

```css
card-element::part(container) {
  /* Style the container part */
  color: /* your color value */;
}
```

### header Part
Card header part

```css
card-element::part(header) {
  /* Style the header part */
  color: /* your color value */;
}
```

### content Part
Card content part

```css
card-element::part(content) {
  /* Style the content part */
  color: /* your color value */;
}
```

### footer Part
Card footer part

```css
card-element::part(footer) {
  /* Style the footer part */
  color: /* your color value */;
}
```



