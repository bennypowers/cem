# Vanilla Example

Web Components built with vanilla JavaScript/TypeScript, no framework required.

## Overview

This example demonstrates how to create custom elements using only native Web APIs (HTMLElement), without any framework like Lit. It shows that CEM works with any approach to building Web Components.

Perfect for:

- Learning Web Components fundamentals
- Projects that want zero dependencies
- Understanding what frameworks do under the hood
- Maximum performance and minimal bundle size

## Features Demonstrated

- **Native `HTMLElement`**: Extending the base class directly
- **Observed Attributes**: Using `static observedAttributes` and `attributeChangedCallback`
- **Properties**: Getters/setters that sync with attributes
- **Custom Events**: Using `CustomEvent` with typed details
- **Shadow DOM**: Manual shadow root creation and rendering
- **Slots**: Content projection with `<slot>`
- **CSS Custom Properties**: Themeable styling
- **CSS Parts**: Shadow DOM styling with `::part()`
- **JSDoc**: Complete documentation using `@element`, `@attr`, `@fires`, etc.

## Project Structure

```
vanilla/
├── .config/
│   └── cem.yaml                     # CEM configuration
├── elements/
│   └── vanilla-element/
│       └── vanilla-element.ts       # Vanilla custom element
├── package.json
├── tsconfig.json
└── README.md
```

## Running the Example

### Install Dependencies

From the repository root:

```bash
npm install  # or pnpm install
```

### Generate the Manifest

```bash
cd examples/vanilla
npm run analyze
```

### Start the Dev Server

```bash
npm run serve
```

## Using the Component

### In HTML

```html
<!DOCTYPE html>
<html>
<head>
  <script type="module" src="/elements/vanilla-element/vanilla-element.js"></script>
  <style>
    vanilla-element {
      --vanilla-bg: #e3f2fd;
      --vanilla-color: #1976d2;
      --vanilla-padding: 1.5rem;
    }

    vanilla-element::part(message) {
      font-weight: bold;
    }
  </style>
</head>
<body>
  <vanilla-element message="Hello Vanilla!"></vanilla-element>

  <vanilla-element message="!sdrawkcab ma I" reversed>
    <p>This is slotted content!</p>
  </vanilla-element>

  <script>
    const el = document.querySelector('vanilla-element');
    el.addEventListener('message-changed', (e) => {
      console.log('Message changed:', e.detail.message);
    });
  </script>
</body>
</html>
```

### In JavaScript

```javascript
const element = document.createElement('vanilla-element');
element.message = 'Dynamically created!';
element.reversed = true;

element.addEventListener('message-changed', (e) => {
  console.log('New message:', e.detail.message);
});

document.body.appendChild(element);
```

## Key Differences from Lit

This vanilla implementation requires you to manually:

- Call `attachShadow()` in the constructor
- Implement `connectedCallback()` for initialization
- Define `observedAttributes` and `attributeChangedCallback()` for reactive attributes
- Create getters/setters for properties
- Manually update the DOM in a `render()` method
- Create and dispatch custom events

Lit automates all of this with decorators like `@customElement`, `@property`, and reactive rendering.

## What's in the Manifest?

The generated manifest includes:

- Element name (`vanilla-element`)
- Attributes (`message`, `reversed`) with types and descriptions
- Properties (getters/setters) with types
- Events (`message-changed`) with detail type
- Slots (default slot)
- CSS custom properties with descriptions
- CSS parts for shadow DOM styling

All extracted from JSDoc comments!

## When to Use Vanilla Web Components

**Pros:**
- Zero runtime dependencies
- Maximum performance
- Full control over everything
- Smallest possible bundle size
- Works in any environment

**Cons:**
- More boilerplate code
- Manual reactivity management
- No template syntax helpers
- Have to handle rendering yourself

## Next Steps

Compare this with the framework-based examples:

- **[minimal](../minimal/)** - Simplest Lit element
- **[intermediate](../intermediate/)** - Multiple Lit components
- **[kitchen-sink](../kitchen-sink/)** - Full Lit features
- **[typescript-paths](../typescript-paths/)** - Advanced build config

## Learn More

- [Web Components MDN Guide](https://developer.mozilla.org/en-US/docs/Web/API/Web_components)
- [HTMLElement API](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement)
- [Custom Elements Spec](https://html.spec.whatwg.org/multipage/custom-elements.html)
- [Shadow DOM](https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_shadow_DOM)
