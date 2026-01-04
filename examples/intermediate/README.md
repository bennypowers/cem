# Intermediate Example

Practical, real-world custom elements demonstrating common UI patterns.

## Overview

This example contains multiple related components that you might actually use in a production application. It demonstrates:

- Multiple components working together
- Common UI patterns (buttons, cards, icons)
- Realistic features and styling
- Demo discovery with HTML metadata
- Progressive enhancement of features from the minimal example

Perfect for:

- Building production component libraries
- Understanding multi-component projects
- Learning practical Web Component patterns
- Seeing how CEM handles multiple elements in one manifest

## Components

### `<ui-button>`

A flexible button component with variants, sizes, and states.

**Features**:
- 4 color variants (primary, secondary, success, danger)
- 3 sizes (small, medium, large)
- Loading and disabled states
- Start/end slots for icons
- CSS parts for styling
- CSS custom properties for theming

### `<ui-card>`

A card component for grouping related content.

**Features**:
- Header, body, and footer slots
- Optional heading property
- Elevation via shadow
- CSS parts for each section
- Themeable via CSS custom properties

### `<ui-icon>`

A simple SVG icon component.

**Features**:
- Built-in icon set (check, x, chevron-right, heart, star)
- Size control
- Color customization
- Accessibility label support
- Warning for missing labels

## Project Structure

```text
intermediate/
├── .config/
│   └── cem.yaml                      # CEM config with demo discovery
├── elements/
│   ├── ui-button/
│   │   ├── ui-button.ts              # Button component
│   │   └── demo/
│   │       ├── variants.html         # Color variants demo
│   │       └── sizes.html            # Size variants demo
│   ├── ui-card/
│   │   └── ui-card.ts                # Card component
│   └── ui-icon/
│       └── ui-icon.ts                # Icon component
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
cd examples/intermediate
npm run analyze
```

This creates `custom-elements.json` with metadata for all three components.

### Start the Dev Server

```bash
npm run serve
```

Navigate to the component files and demo pages in your browser.

## Using the Components

### Basic Usage

```html
<script type="module">
  import "@cem-elements/intermediate/elements/ui-button/ui-button.js";
  import "@cem-elements/intermediate/elements/ui-card/ui-card.js";
  import "@cem-elements/intermediate/elements/ui-icon/ui-icon.js";
</script>

<ui-card heading="Welcome">
  <p>This is a card with a button</p>
  <div slot="footer">
    <ui-button variant="primary">
      <ui-icon name="check" label="Confirm" slot="start"></ui-icon>
      Confirm
    </ui-button>
    <ui-button variant="secondary">Cancel</ui-button>
  </div>
</ui-card>
```

### Styling with CSS Custom Properties

```css
ui-button {
  --button-radius: 999px;  /* Pill-shaped */
  --button-padding: 0.75rem 2rem;
}

ui-card {
  --card-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  --card-radius: 16px;
}

ui-icon {
  --icon-size: 24px;
  --icon-color: #0070f3;
}
```

### Using CSS Parts

```css
ui-button::part(base) {
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

ui-card::part(header) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}
```

## Demo Discovery

This example includes demo HTML files with microdata for automatic demo discovery:

```html
<meta itemprop="description"
      content="Description of this demo...">
```

When you run `npm run analyze`, CEM will automatically discover these demos and include them in the manifest.

## What's in the Manifest?

The generated manifest includes metadata for all three components:

- **Elements**: `ui-button`, `ui-card`, `ui-icon`
- **Properties**: All @property decorators with types and defaults
- **Attributes**: Reflected properties as attributes
- **Slots**: Named and default slots with descriptions
- **Events**: All dispatched events
- **CSS Parts**: All part attributes for styling
- **CSS Custom Properties**: All documented CSS variables
- **Demos**: Auto-discovered HTML demo files
- **Dependencies**: Component dependencies (ui-button depends on ui-spinner concept)

## Key Concepts Demonstrated

### 1. Multiple Components

Unlike the minimal example, this shows how to organize multiple related components in a single project. The manifest will contain all three elements.

### 2. Realistic Styling

Components include real CSS with transitions, hover states, and responsive design.

### 3. Component Composition

Components are designed to work together (e.g., ui-icon inside ui-button slots).

### 4. Accessibility Considerations

- Icon labels for screen readers
- Proper button semantics
- ARIA attributes where needed

### 5. Demo Discovery

HTML files with microdata enable automatic demo discovery, making it easy to showcase component variations.

## Next Steps

Ready for more advanced features?

- **[minimal](../minimal/)** - Start simpler
- **[vanilla](../vanilla/)** - Framework-free approach
- **[kitchen-sink](../kitchen-sink/)** - All CEM features
- **[typescript-paths](../typescript-paths/)** - Advanced build config

## Learn More

- [Lit Documentation](https://lit.dev/)
- [CSS Parts](https://developer.mozilla.org/en-US/docs/Web/CSS/::part)
- [CSS Custom Properties](https://developer.mozilla.org/en-US/docs/Web/CSS/--*)
- [Microdata](https://developer.mozilla.org/en-US/docs/Web/HTML/Microdata)
