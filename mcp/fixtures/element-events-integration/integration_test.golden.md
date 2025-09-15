# Events Reference: `button-element`




## Available Events




### `button-click`

Button click event

**Type:** `CustomEvent`







## JavaScript Integration

### Event Listeners

```javascript
// Get reference to the element
const element = document.querySelector('button-element');


// Listen for button-click event
element.addEventListener('button-click', (event) => {
  console.log('button-click fired:', event.detail);
  // Button click event
});


```

### Event Handler Functions

```javascript

function handlebutton-clickEvent(event) {
  // Button click event
  const data = event.detail;
  // Handle the button-click event
  console.log('Event data:', data);
}


// Attach event handlers
const element = document.querySelector('button-element');
element.addEventListener('button-click', handlebutton-clickEvent);

```

### React Integration

```jsx
import React from 'react';

function MyComponent() {
  
  const handlebutton-clickEvent = (event) => {
    // Button click event
    console.log('button-click:', event.detail);
  };

  
  return (
    <button-element
      onbutton-click={handlebutton-clickEvent}
      
    >
      Content
    </button-element>
  );
}
```



---

For related API information, use:
- **`element_attributes`** - Detailed attribute documentation and usage
- **`element_slots`** - Slot usage patterns and content guidelines
- **`element_styling`** - CSS customization with properties, parts, and states
- **`element_details`** - Complete element reference with all APIs
- **`generate_html`** - HTML generation with proper structure