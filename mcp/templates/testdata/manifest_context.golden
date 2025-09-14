Test component overview with 5 elements

## Schema Context: Understanding Custom Elements Structure

This information follows the **Custom Elements Manifest v2.1.0** specification, which defines the semantic meaning of each field:

### Core Concepts from the Schema

**Custom Elements** are reusable web components with:
- **tagName**: The HTML tag used to invoke the component (kebab-case)
- **attributes**: Properties that control component behavior, with types and constraints
- **slots**: Named content insertion points for composition
- **events**: Custom events the component fires for communication
- **cssProperties**: CSS custom properties for styling customization
- **cssParts**: CSS parts for styling internal component structure
- **cssStates**: CSS custom states for styling component states

**Attributes** have semantic constraints:
- **type**: TypeScript-style type definitions (union types, primitives, references)
- **default**: Default value when not specified
- **required**: Whether the attribute must be provided
- **description**: Human-readable explanation with usage guidelines

**Slots** enable composition:
- **name**: Slot identifier for targeted content insertion
- **description**: Expected content type and usage guidance

This schema provides the semantic framework for understanding your specific component data below.

## Common Patterns in Your Components

### Element Naming Patterns

- **Elements with 'my-' prefix**: my-button, my-card

### Common Attributes

- `variant` - Used in 3 elements

## CSS Custom Properties

**Schema Context**: CSS custom properties are component-specific CSS variables that enable theming and customization. Each property has:
- **name**: The CSS variable name (typically prefixed with component name)
- **syntax**: CSS syntax definition (e.g., `<color>`, `<length>`, `<number>`)
- **inherits**: Whether the property inherits from parent elements
- **initialValue**: Default value when not specified

**Your 2 CSS Custom Properties**:

- `--my-color`
- `--my-size`

These properties allow styling customization while respecting component design constraints.

## Guidelines from Your Manifest Descriptions

These guidelines were extracted from your component and attribute descriptions using RFC 2119 keywords:

### my-button (element)
Use primary variant for main actions.

## How to Use This Context

This information helps AI understand:
- **Your naming conventions** (my, app prefixes)
- **Your component patterns** (common attributes and slots)
- **Your CSS architecture** (custom properties and design tokens)
- **Your documented constraints** (guidelines from descriptions)

When asking for component help, the AI can now reference your actual manifest data rather than generic best practices.