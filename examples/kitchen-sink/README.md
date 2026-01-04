# Kitchen Sink Example

Comprehensive demonstration of ALL Custom Elements Manifest features in one powerful example.

## Overview

This is the most feature-complete example, showcasing every capability of CEM manifest generation. The `demo-button` component is a production-ready, fully-featured button that demonstrates:

- All attribute types (string, boolean, enums, unions)
- Reflected and non-reflected properties
- Multiple slots with conditional rendering
- Custom events
- CSS parts for shadow DOM styling
- CSS custom properties for theming
- Design token integration
- Demo discovery via HTML microdata
- Public methods with JSDoc
- Comprehensive documentation
- Link and button modes
- Form integration capabilities

Perfect for:

- Understanding the full power of CEM
- Reference implementation for complex components
- Learning advanced Web Component patterns
- Seeing how all features work together

## The Demo Button Component

A comprehensive button component adapted from real-world design systems, whitelabeled for educational purposes.

### All Features Demonstrated

#### 1. Properties & Attributes

**Variant** (enum):
- `neutral` - Default gray
- `brand` - Primary brand color
- `success` - Positive actions
- `warning` - Cautionary actions
- `danger` - Destructive actions

**Appearance** (enum):
- `filled` - Solid background
- `outlined` - Border only
- `plain` - No background or border

**Size** (enum):
- `small` - Compact size
- `medium` - Standard size
- `large` - Prominent size

**Boolean Attributes**:
- `pill` - Fully rounded corners
- `disabled` - Prevents interaction
- `loading` - Shows spinner, prevents clicks

**Form Attributes**:
- `type` - button, submit, or reset
- `name` - Form field name
- `value` - Form field value

**Link Attributes**:
- `href` - Render as link
- `target` - Link target
- `rel` - Link relationship
- `download` - Download attribute

#### 2. Slots

- **Default**: Button label/content
- **start**: Leading content (icons, badges)
- **end**: Trailing content (icons, indicators)

#### 3. Events

- `click` - Button clicked
- `focus` - Button gained focus
- `blur` - Button lost focus
- `invalid` - Form validation failed

#### 4. CSS Parts

- `base` - The button element itself
- `start` - Start slot container
- `label` - Label/content container
- `end` - End slot container
- `spinner` - Loading spinner

#### 5. CSS Custom Properties

- `--button-bg` - Background color
- `--button-color` - Text color
- `--button-border` - Border style
- `--button-radius` - Border radius
- `--button-padding` - Internal padding
- `--button-font-size` - Font size
- `--button-font-weight` - Font weight
- `--button-transition` - Transition duration

#### 6. Public Methods

- `click()` - Programmatically click
- `focus(options)` - Set focus
- `blur()` - Remove focus

#### 7. Design Tokens

Integration with design token specification:
- Color tokens (brand, success, warning, danger, neutral)
- Spacing tokens (small, medium, large)
- Radius tokens (small, medium, pill)

Tokens are prefixed with `--demo` and can be used throughout the component.

## Project Structure

```text
kitchen-sink/
├── .config/
│   └── cem.yaml                      # Config with tokens & demos
├── elements/
│   └── demo-button/
│       ├── demo-button.ts            # Comprehensive button component
│       └── demo/
│           ├── variants.html         # All color variants
│           └── states.html           # Interactive states
├── design-tokens.json                # Design token specification
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
cd examples/kitchen-sink
npm run analyze
```

This will:
1. Parse the TypeScript component
2. Extract JSDoc metadata
3. Discover demo HTML files
4. Integrate design tokens
5. Generate comprehensive `custom-elements.json`

### Start the Dev Server

```bash
npm run serve
```

Browse to the component and demo files.

## Using the Component

### Basic Button

```html
<demo-button>Click Me</demo-button>
```

### With Variant and Size

```html
<demo-button variant="success" size="large">
  Save Changes
</demo-button>
```

### With Icons in Slots

```html
<demo-button variant="danger" appearance="outlined">
  <span slot="start">🗑️</span>
  Delete Item
</demo-button>
```

### Loading State

```html
<demo-button loading variant="brand">
  Saving...
</demo-button>
```

### As a Link

```html
<demo-button href="/dashboard" variant="brand">
  Go to Dashboard
</demo-button>
```

### In Forms

```html
<form>
  <demo-button type="submit" variant="success">
    Submit Form
  </demo-button>
  <demo-button type="reset" variant="neutral">
    Reset
  </demo-button>
</form>
```

## Styling Examples

### Using CSS Custom Properties

```css
demo-button {
  --button-bg: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  --button-padding: 1rem 2rem;
  --button-radius: 12px;
  --button-font-weight: 700;
}
```

### Using CSS Parts

```css
demo-button::part(base) {
  text-transform: uppercase;
  letter-spacing: 0.1em;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

demo-button::part(base):hover {
  box-shadow: 0 10px 15px rgba(0, 0, 0, 0.2);
}
```

### Using Design Tokens

The component integrates with design tokens defined in `design-tokens.json`:

```css
demo-button {
  --button-bg: var(--demo-color-brand);
  --button-radius: var(--demo-radius-medium);
  --button-padding: var(--demo-spacing-medium) var(--demo-spacing-large);
}
```

## Demo Discovery

The demo HTML files include microdata for automatic discovery:

```html
<meta itemprop="description"
      content="Comprehensive demonstration of all color variants...">
```

CEM automatically discovers these and adds them to the manifest under the `demos` field.

## What's in the Manifest?

The generated `custom-elements.json` includes:

**Element Metadata**:
- Name, tag, description
- Source code location
- Status and version

**Members**:
- 15+ properties with types and defaults
- Reflected attributes
- Private vs public members

**Slots**:
- Default, start, and end slots
- Descriptions and usage

**Events**:
- All dispatched events
- Event detail types

**Styling**:
- 8 CSS custom properties
- 5 CSS parts
- Design token integration

**Demos**:
- Discovered demo pages
- Descriptions and URLs

**Methods**:
- Public API methods
- Parameters and return types

## Design Token Integration

This example demonstrates design token integration per the [Design Tokens Format Module](https://tr.designtokens.org/format/):

1. Tokens defined in `design-tokens.json`
2. Referenced in CEM config with `--demo` prefix
3. Used in component CSS via custom properties
4. Included in generated manifest

Benefits:
- Consistent design language
- Easy theming
- Design-developer handoff
- Token documentation

## Advanced Patterns

### Dynamic Properties

Properties update DOM attributes for reflection:

```javascript
button.variant = 'success';  // Updates [variant="success"]
button.disabled = true;       // Adds [disabled] attribute
```

### Slot Detection

Component detects slotted content and applies appropriate styling:

```javascript
@state() private hasStartSlot = false;

private handleSlotChange() {
  this.hasStartSlot = this.querySelector('[slot="start"]') !== null;
}
```

### Button vs Link Rendering

Dynamically renders `<button>` or `<a>` based on `href`:

```typescript
const isLink = this.href !== '';
const tag = isLink ? literal`a` : literal`button`;
```

## Comparison with Other Examples

| Feature | Minimal | Intermediate | **Kitchen Sink** | Vanilla |
|---------|---------|--------------|------------------|---------|
| Components | 1 | 3 | 1 (comprehensive) | 1 |
| Properties | 1 | 5-7 | **15+** | 2 |
| Slots | 1 | 2-3 | **3** | 1 |
| CSS Parts | 1 | 2-4 | **5** | 2 |
| Events | 0 | 1-2 | **4** | 1 |
| Design Tokens | ❌ | ❌ | **✅** | ❌ |
| Demo Discovery | ❌ | ✅ | **✅** | ❌ |
| Forms | ❌ | ❌ | **✅** | ❌ |
| Methods | 0 | 0 | **3** | 1 |

## Next Steps

- **[minimal](../minimal/)** - Start with basics
- **[intermediate](../intermediate/)** - Multiple components
- **[vanilla](../vanilla/)** - No framework
- **[typescript-paths](../typescript-paths/)** - Build config

## Learn More

- [Custom Elements Manifest](https://custom-elements-manifest.open-wc.org/)
- [Design Tokens Format](https://tr.designtokens.org/format/)
- [Lit Documentation](https://lit.dev/)
- [CSS Shadow Parts](https://developer.mozilla.org/en-US/docs/Web/CSS/::part)
