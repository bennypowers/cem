# CEM MCP Server

An MCP (Model Context Protocol) server that provides AI assistants with intelligent access to Custom Elements Manifest data for generating accurate HTML with web components.

## What It Does

The CEM MCP Server helps AI assistants understand your custom elements and generate correct HTML by providing:

- **Component discovery**: Find available custom elements in your project
- **Attribute validation**: Get accurate attribute names, types, and values
- **HTML generation**: Generate valid HTML with proper component usage
- **Documentation access**: Access component descriptions and usage guidelines

## Setup

1. **Install CEM**: First make sure you have the CEM CLI installed
   ```bash
   npm install -g @pwrs/cem
   ```

2. **Generate manifests**: Run CEM in your project to generate Custom Elements Manifests
   ```bash
   cem analyze
   ```

3. **Start the MCP server**: Run the CEM MCP server in your project directory
   ```bash
   cem mcp
   ```

4. **Connect to AI assistant**: Configure your AI assistant (Claude Desktop, etc.) to use the MCP server

## Usage Examples

### Discovering Components

Ask your AI assistant about available components:

> "What custom elements are available in this project?"

The MCP server will provide a list of all custom elements found in your manifests.

### Getting Component Information

Ask for details about specific components:

> "Tell me about the `my-button` component"

You'll get information about:
- Available attributes and their types
- Slots for content projection
- CSS custom properties for styling
- Events the component fires
- CSS parts for external styling

### Generating HTML

Ask the AI to generate HTML using your components:

> "Create a form with my-input and my-button components"

The AI will generate valid HTML with correct attribute names and values based on your component definitions.

## Configuration

### Claude Desktop

Add to your Claude Desktop MCP configuration:

```json
{
  "mcpServers": {
    "cem": {
      "command": "cem",
      "args": ["mcp"]
    }
  }
}
```

### Other AI Assistants

The CEM MCP server follows the standard MCP protocol and should work with any MCP-compatible AI assistant.

## What You Get

### Accurate Component Usage

The AI assistant will have access to:
- **Exact attribute names and types** from your component definitions
- **Slot information** for proper content placement
- **CSS custom properties** for theme-aware styling
- **Event details** for interactive components

### Smart HTML Generation

Instead of guessing, the AI can generate HTML like:

```html
<!-- AI knows the correct attributes and their types -->
<my-button variant="primary" size="large">
  Click me
</my-button>

<!-- AI understands slot usage -->
<my-card>
  <h2 slot="title">Card Title</h2>
  <p>Card content goes in the default slot</p>
  <button slot="action">Learn More</button>
</my-card>
```

### Design System Integration

The AI can help with styling using your component's CSS API:

```css
my-button {
  --button-background: var(--primary-color);
  --button-radius: var(--border-radius-md);
}
```

## Benefits

- **No more guessing** - AI uses your actual component definitions
- **Consistent usage** - Components are used according to their documented API
- **Better DX** - Less time looking up component documentation
- **Design system compliance** - Automatic adherence to your design tokens
