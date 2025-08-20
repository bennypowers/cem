# CEM Language Server Demo Project

This is a demonstration project showcasing the capabilities of the CEM Language Server with realistic custom elements built using Lit.

## 🎯 What This Demonstrates

When you open this project in an editor with CEM Language Server support, you'll experience:

### Autocomplete & IntelliSense
- Element name completion (try typing `<button-` or `<card-`)
- Attribute suggestions with documentation
- Slot name completion in HTML
- CSS parts and custom properties

### Hover Documentation
- Rich documentation for elements, attributes, and properties
- Type information and default values
- Slot descriptions and usage examples

### Go-to-Definition
- Jump from HTML usage to TypeScript component definitions
- Navigate between related files and imports

### Error Detection
- Invalid element names
- Unknown attributes
- Incorrect attribute types
- Missing required properties

## 🏗️ Project Structure

```
demo-project/
├── src/
│   ├── components/
│   │   ├── button-element.ts    # Versatile button component
│   │   ├── card-element.ts      # Layout card with slots
│   │   ├── dialog-element.ts    # Modal dialog component
│   │   └── index.ts             # Component exports
│   └── main.ts                  # Application entry point
├── index.html                   # Demo page with examples
├── package.json
├── tsconfig.json
└── custom-elements.json         # Generated CEM file
```

## 🚀 Getting Started

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Generate the Custom Elements Manifest:**
   ```bash
   npm run build
   ```

3. **Start the LSP server:**
   ```bash
   npm run lsp
   ```

4. **Open in your editor:**
   - Open `index.html` to see LSP features in action
   - Try editing component usage to see autocomplete
   - Hover over elements and attributes for documentation
   - Use go-to-definition on element names

## 🎨 Components Overview

### `<button-element>`
A versatile button with multiple variants, sizes, and states.

**Key Features:**
- Variants: primary, secondary, success, danger, warning
- Sizes: small, medium, large
- States: disabled, loading, full-width
- Icon slot support
- CSS parts for styling

### `<card-element>`
A layout component for structured content display.

**Key Features:**
- Header, content, footer, and actions slots
- Theme variants with visual styling
- Elevated appearance option
- Centered content layout
- CSS parts for custom styling

### `<dialog-element>`
A modal dialog with backdrop and focus management.

**Key Features:**
- Backdrop click to close (configurable)
- ESC key support for cancellation
- Keyboard navigation and focus trapping
- Customizable header, content, and footer
- CSS parts for complete styling control

## 🔍 LSP Demo Scenarios

### Scenario 1: Element Discovery
1. Open `index.html`
2. Go to an empty line in the body
3. Type `<` and see custom elements appear
4. Select an element and see attribute suggestions

### Scenario 2: Documentation Access
1. Hover over any `button-element` tag
2. See comprehensive documentation
3. Hover over attributes like `variant` or `size`
4. Notice type information and valid values

### Scenario 3: Error Detection
1. Try typing `<button-element invalid-attr="test">`
2. See the error highlighting for unknown attribute
3. Try `<unknown-element>` and see element error

### Scenario 4: Go-to-Definition
1. Ctrl/Cmd+click on any element name
2. Navigate directly to the TypeScript source
3. Explore the component implementation

### Scenario 5: CSS Parts & Properties
1. Open the `<style>` section in `index.html`
2. Type `.my-element::part(` and see part suggestions
3. Try CSS custom properties with `--card-` prefix

## 🛠️ Development Commands

- `npm run build` - Generate custom-elements.json
- `npm run dev` - Watch for changes and regenerate manifest
- `npm run lsp` - Start the CEM Language Server

## 💡 Tips for Recording

- Use the larger font sizes configured in the recording script
- Focus on the autocomplete popups and hover documentation
- Show both successful completion and error detection
- Demonstrate go-to-definition with obvious visual navigation
- Highlight the seamless workflow between HTML and TypeScript