# AI-Friendly Documentation Example

This example demonstrates comprehensive, AI-friendly JSDoc documentation for custom elements. It shows how to write descriptions that help both human developers and AI assistants understand and use your components correctly.

## Overview

The `<multi-select>` component in this example showcases best practices for documenting complex components with:

- **Rich element descriptions** with purpose, accessibility features, and usage context
- **Detailed event documentation** including trigger conditions and data structures
- **CSS parts descriptions** with styling guidance and accessibility requirements
- **Slot descriptions** explaining content types and relationships
- **CSS custom properties** linking to design system concepts
- **CSS states** with behavioral context
- **RFC 2119 keywords** (MUST, SHOULD, AVOID) for actionable guidelines

## Key Documentation Patterns

### 1. Element Description Structure

```typescript
/**
 * [Component purpose and overview]
 * [Accessibility features and WCAG compliance]
 * [Primary use cases with specific examples]
 *
 * @summary [Brief one-liner]
 * @event ... - [Detailed event documentation]
 * @csspart ... - [Styling guidance for each part]
 * @slot ... - [Content expectations and accessibility]
 * @cssprop ... - [Design system integration]
 * @cssstate ... - [Behavioral context]
 */
```

### 2. RFC 2119 Keywords

The documentation uses RFC 2119 keywords to create structured guidelines:

- **MUST** / **must** - Required behavior
- **SHOULD** / **should** - Recommended practices
- **AVOID** / **avoid** - Discouraged patterns

Example:
> "Should maintain minimum 44px touch target for mobile accessibility"

AI systems extract these as actionable guidelines.

### 3. Event Documentation

Events include:
- **Trigger conditions** - What causes the event to fire
- **Event detail contents** - Data structure and fields
- **Use cases** - When to listen for this event
- **Performance notes** - Debouncing, throttling behavior

### 4. Slot Documentation

Slots describe:
- **Expected content types** - What elements should go in the slot
- **Accessibility considerations** - Screen reader behavior
- **Content relationships** - How slots work together
- **Anti-patterns** - What to avoid

### 5. CSS Property Documentation

CSS properties explain:
- **Visual aspect controlled** - What the property changes
- **Design system integration** - Links to tokens and scales
- **Recommended values** - Practical ranges and examples
- **Functional impact** - Effects beyond visual appearance

## Running the Example

### Install Dependencies

From the repository root:

```bash
npm install  # or pnpm install
```

### Generate the Manifest

```bash
cd examples/ai-friendly-docs
npm run analyze
```

This creates `custom-elements.json` with comprehensive metadata.

### Start the Dev Server

```bash
npm run serve
```

### Test with AI

Use the MCP server to test AI understanding:

```bash
npm run mcp
```

Then ask an AI assistant to:
- Generate HTML using the multi-select component
- Suggest appropriate attribute values for different contexts
- Explain accessibility considerations
- Recommend usage patterns

## Documentation Quality Checklist

Use this checklist when writing component documentation:

- [ ] **Purpose is clear** - New developers understand the component's role
- [ ] **Usage context provided** - When and where to use it
- [ ] **Accessibility included** - Screen reader and keyboard behavior
- [ ] **Examples or guidelines** - Specific guidance with RFC 2119 keywords
- [ ] **Design system integration** - How it fits with other components
- [ ] **Error prevention** - Common mistakes identified
- [ ] **Event data structures** - What information events provide
- [ ] **Slot relationships** - How slot content works together
- [ ] **CSS property ranges** - Recommended values and constraints

## Key Concepts Demonstrated

### Comprehensive Element Documentation

The multi-select component shows how to document a complex component with multiple interaction patterns, state management, accessibility features, and customization options.

### Design System Integration

Documentation references design tokens, elevation scales, and consistency patterns that help AI understand how components fit into larger design systems.

### Accessibility-First Approach

Every aspect of the documentation includes accessibility considerations:
- Screen reader announcements
- Keyboard navigation patterns
- WCAG compliance levels
- Focus management
- ARIA attribute usage

### Progressive Disclosure

Documentation is structured for different audiences:
- **Summary** - Quick overview
- **Description** - Detailed explanation
- **Examples** - Specific use cases
- **Guidelines** - Best practices with RFC 2119 keywords

## Learn More

- **[Effective Writing for AI](/docs/usage/effective-mcp-descriptions/)** - Complete guide to AI-friendly documentation
- **[MCP Integration](/docs/installation/mcp/)** - Setup AI assistants
- **[RFC 2119](https://www.rfc-editor.org/rfc/rfc2119)** - Keyword specifications
- **[Custom Elements Manifest](https://github.com/webcomponents/custom-elements-json)** - Schema specification

## Related Examples

- **[minimal](../minimal/)** - Simplest possible component
- **[intermediate](../intermediate/)** - Practical multi-component project
- **[kitchen-sink](../kitchen-sink/)** - All CEM features
