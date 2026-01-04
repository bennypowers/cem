# Minimal Example

The simplest possible custom element using CEM (Custom Elements Manifest).

## Overview

This example demonstrates the absolute minimum needed to create a custom element and generate a manifest with CEM. It's perfect for:

- First-time users learning custom elements
- Understanding the basics of CEM manifest generation
- Quick prototyping of new components

## Features Demonstrated

- **Custom Element**: Basic `@customElement` decorator
- **Reactive Property**: Single `@property` with default value
- **Slot**: Default slot for content projection
- **CSS Custom Property**: One CSS variable for styling
- **CSS Part**: One part for shadow DOM styling
- **JSDoc**: Basic documentation with `@summary` and `@slot` tags

## Project Structure

```text
minimal/
├── .config/
│   └── cem.yaml              # CEM configuration
├── elements/
│   └── hello-world/
│       └── hello-world.ts    # The component
├── package.json              # Package configuration
├── tsconfig.json             # TypeScript settings
└── README.md                 # This file
```

## Running the Example

### Install Dependencies

From the repository root:

```bash
npm install  # or pnpm install
```

### Generate the Manifest

```bash
cd examples/minimal
npm run analyze
```

This will create `custom-elements.json` in this directory.

### Start the Dev Server

```bash
npm run serve
```

Then open your browser and navigate to the component file to see it in action.

## Using the Component

### In HTML

```html
<!DOCTYPE html>
<html>
<head>
  <script type="module" src="/elements/hello-world/hello-world.js"></script>
</head>
<body>
  <hello-world name="CEM"></hello-world>
  <hello-world>Have a great day!</hello-world>
</body>
</html>
```

### Styling with CSS

```css
hello-world {
  --hello-world-color: blue;
}

hello-world::part(greeting) {
  font-weight: bold;
}
```

## What's in the Manifest?

After running `npm run analyze`, check `custom-elements.json` to see:

- Component name and tag name (`hello-world`)
- Properties (`name` with type and default value)
- Slots (default slot)
- CSS custom properties (`--hello-world-color`)
- CSS parts (`greeting`)
- JSDoc descriptions

## Next Steps

Once you're comfortable with this minimal example, check out:

- **[intermediate](../intermediate/)** - Multiple components with more features
- **[vanilla](../vanilla/)** - Web components without Lit
- **[kitchen-sink](../kitchen-sink/)** - All CEM features demonstrated
- **[typescript-paths](../typescript-paths/)** - Advanced build setup

## Learn More

- [Custom Elements Manifest Documentation](https://custom-elements-manifest.open-wc.org/)
- [Lit Documentation](https://lit.dev/)
- [Web Components Basics](https://developer.mozilla.org/en-US/docs/Web/API/Web_components)
