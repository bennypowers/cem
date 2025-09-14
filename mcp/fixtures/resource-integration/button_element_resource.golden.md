# button-element Usage Examples

## Basic Usage
```html
&lt;button-element>&lt;/button-element>
```




## With Attributes

### variant Attribute
Button variant
- **Type**: &#34;primary&#34; | &#34;secondary&#34; | &#34;ghost&#34;
- **Default**: &#34;primary&#34;
- **Options**: &#34;primary&#34;, &#34;secondary&#34;, &#34;ghost&#34;

```html
&lt;button-element variant="&#34;primary&#34;">&lt;/button-element>
```

### size Attribute
Button size
- **Type**: &#34;small&#34; | &#34;medium&#34; | &#34;large&#34;
- **Default**: &#34;medium&#34;
- **Options**: &#34;small&#34;, &#34;medium&#34;, &#34;large&#34;

```html
&lt;button-element size="&#34;medium&#34;">&lt;/button-element>
```

### disabled Attribute
Whether button is disabled
- **Type**: boolean

```html
&lt;button-element disabled="/* your boolean value */">&lt;/button-element>
```




## With Slots

### Default Slot
Button content

```html
&lt;button-element>
  <p>Default slot content</p>
&lt;/button-element>
```

### icon Slot
Button icon

```html
&lt;button-element>
  <span slot="icon">icon content</span>
&lt;/button-element>
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
- **Syntax**: &lt;color&gt;
- **Initial**: blue

```css
button-element {
  --button-color: blue;
}
```

### --button-padding
Button padding
- **Syntax**: &lt;length&gt;
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



