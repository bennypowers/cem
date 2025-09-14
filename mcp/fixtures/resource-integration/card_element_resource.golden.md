# card-element Usage Examples

## Basic Usage
```html
<card-element></card-element>
```




## With Attributes

### elevation Attribute
Card elevation level
- **Type**: number
- **Default**: 1

```html
<card-element elevation=1></card-element>
```




## With Slots

### Default Slot
Card content

```html
<card-element>
  <p>Default slot content</p>
</card-element>
```

### header Slot
Card header

```html
<card-element>
  <span slot="header">header content</span>
</card-element>
```

### footer Slot
Card footer

```html
<card-element>
  <span slot="footer">footer content</span>
</card-element>
```






## CSS Custom Properties

### --card-background
Card background color
- **Syntax**: `<color>`
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



