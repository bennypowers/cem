---
title: Writing Effective Manifest Descriptions for AI
weight: 45
---

When using the `cem mcp` server, the quality of your custom elements manifest descriptions directly impacts how well AI systems can understand and use your components. This guide shows you how to write descriptions that work excellently with both human developers and AI assistants.

## Why Description Quality Matters

AI language models use your manifest descriptions to:
- Generate proper HTML with correct attributes and slots
- Suggest appropriate component usage patterns  
- Validate accessibility compliance
- Provide contextual recommendations
- Understand component relationships and hierarchies

Well-written descriptions enable AI to make intelligent decisions about component usage, while poor descriptions lead to incorrect or suboptimal code generation.

## General Principles

### 1. Write for Both Humans and AI

Good descriptions serve dual purposes:
- **Human developers** need clear, concise information for manual coding
- **AI systems** need structured, actionable guidance for automatic generation

```json
// Good: Clear and actionable
{
  "description": "A navigation button that supports keyboard interaction and screen readers. Use for primary actions in forms and dialogs."
}

// Poor: Too brief and technical
{
  "description": "A custom button element with click handling"
}
```

### 2. Include Usage Context and Guidelines

AI systems benefit from understanding **when** and **how** to use components:

```json
// Good: Includes context and guidelines
{
  "description": "A card container for grouping related content. Should include a header slot for titles and supports elevation levels 0-5 for visual hierarchy. Use sparingly to avoid visual clutter."
}

// Poor: Just describes what it is
{
  "description": "A card component with slots"
}
```

### 3. Emphasize Accessibility Requirements

Since accessibility is a core priority, include accessibility guidance in descriptions:

```json
// Good: Accessibility-focused
{
  "description": "An interactive toggle button that announces state changes to screen readers. Must include aria-pressed attribute and supports keyboard activation with Space or Enter keys."
}

// Poor: No accessibility information
{
  "description": "A toggle button component"
}
```

## Element Descriptions

Element descriptions should answer: **What is this component and when should I use it?**

### Structure Your Element Descriptions

Use this pattern for comprehensive element descriptions:

```
[Component Purpose] [Accessibility Features] [Usage Guidelines] [Context Information]
```

### Examples by Component Type

#### Interactive Components (Buttons, Inputs, Controls)

```json
{
  "name": "ActionButton",
  "description": "A primary action button with built-in loading states and accessibility features. Automatically includes proper ARIA attributes and keyboard support. Use for important actions like form submission, confirmation dialogs, and primary navigation. Supports variant styling for different contexts."
}
```

**Key elements:**
- Purpose: "primary action button"
- Accessibility: "ARIA attributes and keyboard support"  
- Usage context: "form submission, confirmation dialogs"
- Features: "loading states", "variant styling"

#### Layout Components (Cards, Containers, Grids)

```json
{
  "name": "ContentCard", 
  "description": "A flexible content container with semantic markup for grouping related information. Includes slots for header, body, and footer content with proper heading hierarchy. Use for dashboard widgets, product listings, and content previews. Supports responsive design and maintains accessibility across screen sizes."
}
```

**Key elements:**
- Purpose: "content container for grouping"
- Structure: "slots for header, body, footer"
- Semantic info: "proper heading hierarchy"
- Usage examples: "dashboard widgets, product listings"
- Responsive: "maintains accessibility across screen sizes"

#### Navigation Components (Menus, Tabs, Breadcrumbs)

```json
{
  "name": "TabNavigation",
  "description": "An accessible tab navigation component that manages focus and announces selections to screen readers. Implements WCAG navigation patterns with arrow key support and automatic ARIA relationships. Use for organizing content into logical sections within a single view."
}
```

**Key elements:**
- Accessibility: "manages focus", "announces selections", "WCAG patterns"
- Interaction: "arrow key support", "ARIA relationships"
- Usage context: "organizing content into logical sections"

## Attribute Descriptions

Attribute descriptions should answer: **What does this control and how should I use it?**

### Include Type Information and Valid Values

```json
// Good: Complete attribute information
{
  "name": "variant",
  "type": "\"primary\" | \"secondary\" | \"danger\" | \"ghost\"", 
  "description": "Button visual style and semantic meaning. Use 'primary' for main actions, 'secondary' for supporting actions, 'danger' for destructive actions, and 'ghost' for subtle actions. Affects color contrast and accessibility announcements."
}

// Poor: Minimal information
{
  "name": "variant",
  "type": "string",
  "description": "Button style"
}
```

### Explain Impact and Usage Context

```json
// Good: Explains impact and context
{
  "name": "size",
  "type": "\"small\" | \"medium\" | \"large\"",
  "description": "Controls button dimensions and touch target size for accessibility. Use 'small' for compact layouts, 'medium' for standard interfaces, 'large' for mobile-first or prominent actions. Ensures minimum 44px touch target on mobile devices."
}
```

### Include Default Behavior

```json
// Good: Explains default and when to change
{
  "name": "loading", 
  "type": "boolean",
  "description": "Shows loading spinner and disables interaction during async operations. Automatically announces loading state to screen readers. Set to true during form submission or data fetching to prevent duplicate actions.",
  "default": "false"
}
```

## Slot Descriptions

Slot descriptions should answer: **What content belongs here and how should it be structured?**

### Describe Expected Content Types

```json
// Good: Specific content guidance
{
  "name": "header",
  "description": "Card header content, typically a heading element (h2-h4) or title text. Should establish the card's purpose and maintain proper heading hierarchy. Avoid interactive elements that might conflict with card-level actions."
}

// Poor: Generic description
{
  "name": "header", 
  "description": "Header content"
}
```

### Include Accessibility Considerations

```json
// Good: Accessibility-aware slot description
{
  "name": "action",
  "description": "Primary action area for buttons or links. Content should be descriptive for screen readers (avoid 'click here' or 'read more'). Interactive elements will receive proper focus management within the card context."
}
```

### Explain Content Relationships

```json
// Good: Explains how content works together  
{
  "name": "",
  "description": "Main card content area for body text, lists, or media. Content should complement the header and any action elements. Will be announced as the primary card content by screen readers after the header."
}
```

## CSS Property Descriptions

CSS property descriptions should answer: **What visual aspect does this control and how does it fit the design system?**

### Link to Design System Concepts

```json
// Good: Design system integration
{
  "name": "--card-elevation",
  "syntax": "<length>",
  "description": "Controls card shadow depth using design system elevation scale. Maps to elevation tokens 0-5 where 0 is flat and 5 is highest prominence. Affects both visual hierarchy and accessibility by indicating interactive priority.",
  "inherits": false,
  "initial": "2px"
}
```

### Explain Visual and Functional Impact

```json
// Good: Visual and functional explanation
{
  "name": "--button-radius",
  "syntax": "<length> | <percentage>", 
  "description": "Border radius for button corners affecting visual style and brand consistency. Use design system radius tokens (--radius-sm, --radius-md, --radius-lg) for consistency. Impacts touch target perception on mobile devices.",
  "inherits": true,
  "initial": "4px"
}
```

## Event Descriptions

Event descriptions should answer: **When does this fire and what data does it provide?**

### Describe Trigger Conditions and Data

```json
// Good: Complete event information
{
  "name": "selection-change",
  "type": "CustomEvent",
  "description": "Fired when user selects or deselects items, including keyboard and mouse interactions. Event detail contains selectedItems array and previousSelection for accessibility announcements. Debounced to prevent excessive firing during rapid selection changes."
}
```

### Include Accessibility Information

```json
// Good: Accessibility integration
{
  "name": "expand-toggle", 
  "type": "CustomEvent",
  "description": "Fired when collapsible content expands or collapses. Event detail includes expanded state and target element for screen reader announcements. Use to coordinate ARIA attributes and manage focus when content visibility changes."
}
```

## CSS Parts and States

### CSS Parts: Describe Styling Purpose

```json
// Good: Clear styling guidance
{
  "name": "input",
  "description": "The input field element for styling text appearance, borders, and focus states. Inherits form-level styling but can be customized for specific use cases. Maintain sufficient color contrast for accessibility."
}
```

### CSS States: Explain Behavioral Context

```json
// Good: Behavioral context
{
  "name": "loading", 
  "description": "Applied during async operations when the component is waiting for data or processing requests. Use to style loading indicators and adjust interaction states. Should provide visual feedback while maintaining accessibility."
}
```

## Common Patterns and Anti-Patterns

### ✅ Good Patterns

1. **Action-oriented language**: "Use for...", "Include...", "Ensure..."
2. **Context information**: When, where, and why to use
3. **Accessibility integration**: Screen reader behavior, keyboard support
4. **Design system alignment**: Token usage, consistency guidelines
5. **Error prevention**: Common mistakes and how to avoid them

### ❌ Anti-Patterns

1. **Implementation details**: Avoid technical internals that don't help usage
2. **Vague language**: "Customizes appearance" instead of specific guidance
3. **Missing accessibility**: No mention of screen reader or keyboard behavior  
4. **No usage context**: Describing what something is without when to use it
5. **Inconsistent terminology**: Different names for the same concepts

## Testing Your Descriptions

### Manual Review Checklist

- [ ] **Purpose is clear**: A developer new to the component understands its role
- [ ] **Usage context provided**: When and where to use the component
- [ ] **Accessibility included**: Screen reader and keyboard behavior mentioned
- [ ] **Examples or guidelines**: Specific guidance for proper usage
- [ ] **Design system integration**: How it fits with other components
- [ ] **Error prevention**: Common mistakes identified and avoided

### AI Testing

Use the `cem mcp` server to test your descriptions:

1. Ask AI to generate HTML using your component
2. Request attribute suggestions for different contexts
3. Check if accessibility recommendations are appropriate
4. Verify that generated examples follow your intended usage patterns

The AI should be able to make intelligent decisions based solely on your manifest descriptions. If it generates inappropriate usage, consider refining your descriptions for clarity and completeness.

## Examples by Use Case

### Form Components

Focus on validation, labeling, and form submission integration:

```json
{
  "name": "email-input",
  "description": "An email input field with built-in validation and accessibility features. Automatically provides email format validation, proper input type for mobile keyboards, and integrates with form submission. Include proper labeling and error message association for screen readers."
}
```

### Data Display Components  

Emphasize content structure and relationships:

```json
{
  "name": "data-table", 
  "description": "An accessible data table with sorting, filtering, and keyboard navigation. Maintains proper table semantics for screen readers and provides column header associations. Use for structured data display where relationships between data points are important."
}
```

### Modal and Overlay Components

Focus on focus management and escape patterns:

```json
{
  "name": "modal-dialog",
  "description": "A modal dialog that traps focus and provides accessible close mechanisms. Automatically manages focus restoration, escape key handling, and backdrop click behavior. Include proper heading structure and action buttons for clarity."
}
```

Following these guidelines ensures your custom elements manifest provides rich, actionable information that enables both human developers and AI systems to use your components correctly, accessibly, and effectively.