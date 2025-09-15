# Styling Reference: `button-element`





## CSS Custom Properties




### `--button-color`

Button color

**Syntax:** `<color>` | **Initial:** `blue`






### `--button-padding`

Button padding

**Syntax:** `<length>` | **Initial:** `8px`







### Usage Examples

```css
/* Basic theming */
button-element {
  --button-color: /* your value (<color>) */;
  --button-padding: /* your value (<length>) */;
}

/* With custom values */
button-element {
  --button-color: blue;
  --button-padding: 8px;
}
```




## CSS Parts

The description of a CSS Part


### `button` part

The button element







### Usage Examples

```css
/* Style individual parts */
button-element::part(button) {
  /* The button element */
}


/* Multiple parts styling */
button-element::part(button) {
  font-family: var(--font-family);
}
```





## Complete Styling Example

```css
/* Complete styling with all available APIs */
button-element {
  --button-color: blue;
  --button-padding: 8px;
}


/* Style button part */
button-element::part(button) {
  /* The button element */
}



```



---

For related API information, use:
- **`element_attributes`** - Detailed attribute documentation and usage
- **`element_slots`** - Slot usage patterns and content guidelines
- **`element_events`** - Event details and JavaScript integration
- **`element_details`** - Complete element reference with all APIs
- **`suggest_css_integration`** - Design system integration guidance
- **`generate_html`** - HTML generation with proper structure