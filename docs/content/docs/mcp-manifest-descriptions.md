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

## Complete Example: Complex Component Documentation

Here's a comprehensive example showing how to properly document a complex component with all its APIs:

### Multi-Select Component: Full Documentation

This example demonstrates complete documentation for a sophisticated multi-select component with rich interactions, accessibility features, and extensive customization options.

```json
{
  "name": "multi-select",
  "tagName": "multi-select",
  "description": "An accessible multi-select dropdown with search, keyboard navigation, and bulk operations. Implements WCAG 2.1 AA standards with comprehensive screen reader support and keyboard interaction patterns. Use for selection from large datasets where users need to choose multiple related items, such as tag selection, user assignment, or category filtering.",
  
  "attributes": [
    {
      "name": "placeholder",
      "type": "string",
      "description": "Placeholder text shown when no selections are made. Should be descriptive and action-oriented (e.g., 'Select team members' rather than 'Choose options'). Announced by screen readers when the component receives focus.",
      "default": "Select options..."
    },
    {
      "name": "searchable",
      "type": "boolean", 
      "description": "Enables search/filter functionality within the dropdown. When true, users can type to filter available options, improving usability for large option sets. Search is case-insensitive and matches partial strings. Essential for lists with more than 10 options.",
      "default": "false"
    },
    {
      "name": "max-selections",
      "type": "number",
      "description": "Maximum number of items that can be selected. When limit is reached, remaining options become disabled with appropriate ARIA states. Use to prevent overwhelming interfaces or enforce business rules. Value of 0 means unlimited selections.",
      "default": "0"
    },
    {
      "name": "required",
      "type": "boolean",
      "description": "Marks the component as required for form validation. Displays visual indicator and updates ARIA attributes for screen readers. Prevents form submission when no selections are made. Should be paired with clear error messaging.",
      "default": "false"
    },
    {
      "name": "disabled",
      "type": "boolean",
      "description": "Disables all interaction with the component. Sets appropriate ARIA states and visual styling. Use for read-only states or when dependent form fields haven't been completed. Maintains selection visibility for context.",
      "default": "false"
    },
    {
      "name": "variant",
      "type": "\"default\" | \"compact\" | \"detailed\"",
      "description": "Controls layout density and information display. 'default' shows standard spacing and basic option text. 'compact' reduces spacing for dense interfaces. 'detailed' shows additional option metadata and descriptions. Choose based on interface density and content complexity.",
      "default": "default"
    },
    {
      "name": "selection-display",
      "type": "\"tags\" | \"count\" | \"list\"",
      "description": "How selected items are displayed in the closed state. 'tags' shows individual removable chips (best for <= 5 selections). 'count' shows summary like '3 selected' (best for large selections). 'list' shows comma-separated names (best for 2-4 selections). Automatically adjusts for accessibility and screen space.",
      "default": "tags"
    }
  ],

  "slots": [
    {
      "name": "",
      "description": "Default slot for option elements. Should contain option elements with value attributes and descriptive text content. Options support disabled state and optional description text. Maintain semantic markup for screen readers - each option should be self-contained and descriptive."
    },
    {
      "name": "empty-state", 
      "description": "Content shown when no options match current search or when options list is empty. Should provide helpful guidance or alternative actions. Include clear messaging about search refinement or data loading states. Will be announced by screen readers when displayed."
    },
    {
      "name": "header",
      "description": "Optional header content above the options list. Use for category labels, action buttons (like 'Select All'), or contextual help. Should not contain form controls that conflict with option selection. Announced as part of dropdown context by screen readers."
    },
    {
      "name": "footer",
      "description": "Footer content below options list. Ideal for summary information ('X of Y selected'), action buttons, or links to manage the option set externally. Maintains focus context within the dropdown for keyboard navigation."
    }
  ],

  "events": [
    {
      "name": "selection-change",
      "type": "CustomEvent<{selectedValues: string[], selectedOptions: Element[], action: 'add'|'remove'|'clear', changedValue?: string}>",
      "description": "Fired when selection changes through any interaction method (click, keyboard, programmatic). Event detail includes complete selection state and the specific change that occurred. Use for form validation, dependent field updates, or analytics tracking. Debounced during rapid keyboard navigation to prevent excessive firing."
    },
    {
      "name": "search-input",
      "type": "CustomEvent<{query: string, filteredCount: number}>",
      "description": "Fired when user types in search field (when searchable is true). Event detail contains current search query and number of matching options. Use for external filtering, analytics, or implementing custom search behavior. Fires after built-in filtering is applied."
    },
    {
      "name": "dropdown-toggle",
      "type": "CustomEvent<{isOpen: boolean, trigger: 'click'|'keyboard'|'programmatic'}>",
      "description": "Fired when dropdown opens or closes. Event detail indicates new state and what triggered the change. Use for coordinating with external UI elements, implementing custom positioning, or managing page-level focus. Essential for modal-style behaviors."
    },
    {
      "name": "max-reached",
      "type": "CustomEvent<{maxSelections: number, attemptedValue: string}>",
      "description": "Fired when user attempts to select beyond max-selections limit. Use to show helpful error messages or suggest alternative actions. Event is fired before selection is rejected, allowing for custom handling or user education."
    }
  ],

  "cssProperties": [
    {
      "name": "--multi-select-width",
      "syntax": "<length> | <percentage>",
      "description": "Controls component width. Use design system width tokens or specific measurements. Affects dropdown positioning and text wrapping behavior. Consider content length and container constraints when setting.",
      "inherits": false,
      "initial": "auto"
    },
    {
      "name": "--multi-select-max-height",
      "syntax": "<length>",
      "description": "Maximum height for the dropdown options list. When exceeded, the list becomes scrollable with keyboard navigation support. Use viewport-relative units for responsive behavior. Recommended range: 200px to 60vh for optimal usability.",
      "inherits": false,
      "initial": "300px"
    },
    {
      "name": "--selection-tag-color",
      "syntax": "<color>",
      "description": "Background color for selected item tags (when selection-display='tags'). Should maintain sufficient contrast with text color for accessibility. Use semantic colors from design system that indicate selection state clearly.",
      "inherits": true,
      "initial": "var(--color-primary-100)"
    },
    {
      "name": "--option-hover-bg",
      "syntax": "<color>",
      "description": "Background color for option hover state. Must meet WCAG contrast requirements against option text. Provides visual feedback for keyboard and mouse navigation. Should be distinct from selection and focus states.",
      "inherits": true,
      "initial": "var(--color-gray-100)"
    },
    {
      "name": "--dropdown-border-radius",
      "syntax": "<length>",
      "description": "Border radius for dropdown container. Use design system radius tokens for consistency. Affects visual integration with page design and brand consistency. Consider relationship to trigger element styling.",
      "inherits": true,
      "initial": "var(--radius-md)"
    }
  ],

  "cssParts": [
    {
      "name": "trigger",
      "description": "Main clickable area that opens the dropdown. Style for brand consistency and clear affordance. Includes focus indicators and disabled states. Should maintain minimum 44px touch target for mobile accessibility."
    },
    {
      "name": "selections",
      "description": "Container for displaying selected items. Layout varies based on selection-display attribute. Style for clear visual hierarchy and easy scanning. Handles overflow with appropriate text truncation."
    },
    {
      "name": "dropdown",
      "description": "Dropdown container including search, options, and footer areas. Handles positioning, shadows, and border styling. Consider z-index stacking context and responsive positioning behavior."
    },
    {
      "name": "search-input",
      "description": "Search input field within dropdown (when searchable=true). Style for clear focus states and placeholder visibility. Should integrate seamlessly with dropdown design while remaining clearly functional."
    },
    {
      "name": "option",
      "description": "Individual option elements within the list. Style for clear selection states, hover feedback, and disabled appearance. Must maintain accessibility contrast requirements across all states."
    },
    {
      "name": "selected-tag",
      "description": "Individual selection chips (when selection-display='tags'). Include close button styling and hover states. Should be easily scannable and clearly removable for keyboard users."
    }
  ],

  "cssStates": [
    {
      "name": "open",
      "description": "Applied when dropdown is expanded. Use for styling state transitions, positioning adjustments, or coordinating with external elements. Helps manage z-index and overlay behaviors."
    },
    {
      "name": "searching",
      "description": "Applied when user is actively typing in search field. Use for loading indicators, highlighting matched text, or adjusting dropdown layout. Provides visual feedback during search operations."
    },
    {
      "name": "max-reached",
      "description": "Applied when maximum selection limit is reached. Use for visual feedback, disabling remaining options, or showing helpful messaging. Should clearly communicate the constraint to users."
    },
    {
      "name": "required-empty",
      "description": "Applied when component is required but has no selections. Use for error styling, validation messaging, or form submission prevention. Should clearly indicate the validation state."
    }
  ]
}
```

### Key Documentation Patterns Demonstrated

This example showcases several critical documentation patterns:

1. **Comprehensive Context**: Each description explains not just what something does, but when and why to use it
2. **Accessibility Integration**: Every interactive element includes accessibility considerations and screen reader behavior
3. **Design System Alignment**: CSS properties reference design tokens and explain visual hierarchy
4. **Error Prevention**: Descriptions anticipate common mistakes and provide guidance
5. **Usage Scaling**: Recommendations change based on data size and interface density
6. **Progressive Enhancement**: Features work together but can be used independently

## Examples by Use Case

### Form Components

Focus on validation, labeling, and form submission integration:

```json
{
  "name": "email-input",
  "description": "An email input field with built-in validation and accessibility features. Automatically provides email format validation, proper input type for mobile keyboards, and integrates with form submission. Include proper labeling and error message association for screen readers.",
  
  "attributes": [
    {
      "name": "validate-on",
      "type": "\"blur\" | \"input\" | \"submit\"",
      "description": "When to trigger email format validation. 'blur' validates when field loses focus (recommended for better UX), 'input' validates on every keystroke (use sparingly), 'submit' validates only on form submission. Choose based on user experience needs and form complexity.",
      "default": "blur"
    },
    {
      "name": "suggest-domains",
      "type": "boolean",
      "description": "Enables common email domain suggestions (gmail.com, outlook.com, etc.) when user types '@'. Improves accuracy and reduces typos. Suggestions are keyboard navigable and announced to screen readers. Particularly useful for public-facing forms.",
      "default": "false"
    }
  ],

  "events": [
    {
      "name": "email-validated",
      "type": "CustomEvent<{isValid: boolean, email: string, errors: string[]}>",
      "description": "Fired after email validation completes. Event detail includes validation state, cleaned email value, and any error messages for display. Use for coordinating form validation state or implementing custom error display patterns."
    }
  ]
}
```

### Data Display Components  

Emphasize content structure and relationships:

```json
{
  "name": "data-table", 
  "description": "An accessible data table with sorting, filtering, and keyboard navigation. Maintains proper table semantics for screen readers and provides column header associations. Use for structured data display where relationships between data points are important.",
  
  "attributes": [
    {
      "name": "selectable",
      "type": "\"none\" | \"single\" | \"multiple\"",
      "description": "Row selection behavior. 'none' for read-only tables, 'single' for choosing one item, 'multiple' for bulk operations. When enabled, adds proper ARIA attributes and keyboard selection patterns. Include clear visual indicators for selected state.",
      "default": "none"
    },
    {
      "name": "sticky-header",
      "type": "boolean",
      "description": "Makes table header sticky during vertical scrolling. Improves usability for long tables by maintaining column context. Automatically adjusts for keyboard navigation and screen reader column announcements. Consider header height in page layout.",
      "default": "false"
    }
  ],

  "slots": [
    {
      "name": "toolbar",
      "description": "Actions and controls above the table (search, filters, bulk actions). Content should relate to table operations and maintain logical tab order. Use for functionality that affects the entire table view rather than individual rows."
    },
    {
      "name": "empty-state",
      "description": "Content shown when table has no data or filtered results. Should provide clear next steps or alternative actions. Include helpful messaging about data loading, permission requirements, or how to add content."
    }
  ]
}
```

### Modal and Overlay Components

Focus on focus management and escape patterns:

```json
{
  "name": "modal-dialog",
  "description": "A modal dialog that traps focus and provides accessible close mechanisms. Automatically manages focus restoration, escape key handling, and backdrop click behavior. Include proper heading structure and action buttons for clarity.",
  
  "attributes": [
    {
      "name": "close-on-backdrop",
      "type": "boolean",
      "description": "Allows closing dialog by clicking outside content area. Enable for non-critical dialogs where accidental closure isn't problematic. Disable for forms or confirmation dialogs where data loss is a concern. Always provide explicit close button regardless.",
      "default": "true"
    },
    {
      "name": "size",
      "type": "\"small\" | \"medium\" | \"large\" | \"fullscreen\"",
      "description": "Modal size affecting width and responsive behavior. 'small' for confirmations and simple forms, 'medium' for detailed forms, 'large' for complex workflows, 'fullscreen' for immersive experiences. Automatically adjusts on mobile devices.",
      "default": "medium"
    }
  ],

  "events": [
    {
      "name": "modal-closed",
      "type": "CustomEvent<{reason: 'backdrop' | 'escape' | 'button' | 'programmatic'}>",
      "description": "Fired when modal closes through any method. Event detail indicates how the modal was closed for analytics or conditional behavior. Use to save draft data, clear form state, or coordinate with external UI elements."
    }
  ],

  "slots": [
    {
      "name": "header",
      "description": "Modal header typically containing title and close button. Should include a heading element (h1-h6) for proper document structure and screen reader navigation. Close button should be clearly labeled and keyboard accessible."
    },
    {
      "name": "footer",
      "description": "Modal footer for action buttons (Save, Cancel, etc.). Buttons should have clear labels and logical tab order. Include primary action styling to guide user attention. Consider button placement for cultural reading patterns."
    }
  ]
}
```

### Real-World Complex Example: Product Card

Here's how a complex e-commerce product card might be fully documented:

```json
{
  "name": "product-card",
  "description": "A comprehensive product display card for e-commerce interfaces with image gallery, pricing, availability, and quick actions. Implements structured data markup for SEO and maintains full accessibility for screen readers and keyboard users. Use in product grids, search results, and recommendation sections.",
  
  "attributes": [
    {
      "name": "product-id",
      "type": "string",
      "description": "Unique identifier for the product. Used for tracking, analytics, and cart operations. Should correspond to your product database ID for consistency across systems.",
      "required": true
    },
    {
      "name": "layout",
      "type": "\"standard\" | \"compact\" | \"featured\"",
      "description": "Card layout variant. 'standard' for regular grid displays, 'compact' for dense listing views, 'featured' for hero sections with enhanced imagery and details. Affects image aspect ratio and information hierarchy.",
      "default": "standard"
    },
    {
      "name": "show-quick-add",
      "type": "boolean",
      "description": "Enables quick add-to-cart functionality without navigation. Shows quantity selector and add button on hover/focus. Disable for products requiring configuration or when cart integration isn't available. Includes proper loading states and error handling.",
      "default": "false"
    }
  ],

  "events": [
    {
      "name": "add-to-cart",
      "type": "CustomEvent<{productId: string, quantity: number, variant?: string}>",
      "description": "Fired when user adds product to cart via quick-add or action button. Event detail includes product information for cart integration. Use to update cart state, show confirmation, or track conversion events."
    },
    {
      "name": "product-view",
      "type": "CustomEvent<{productId: string, source: 'card-click' | 'image-click' | 'title-click'}>",
      "description": "Fired when user clicks to view product details. Event detail indicates which element was clicked for analytics. Use to track engagement patterns and optimize card design for better conversion."
    }
  ],

  "slots": [
    {
      "name": "badges",
      "description": "Product status badges (Sale, New, Limited, etc.). Content should be concise and meaningful. Use semantic colors and ensure badges don't overwhelm the product information. Will be announced by screen readers as part of product description."
    },
    {
      "name": "actions",
      "description": "Additional action buttons beyond standard add-to-cart (wishlist, compare, share). Actions should be clearly labeled and follow consistent interaction patterns. Consider mobile touch targets and keyboard navigation order."
    }
  ]
}
```

Following these comprehensive documentation patterns ensures your custom elements manifest provides rich, actionable information that enables both human developers and AI systems to use your components correctly, accessibly, and effectively. The key is balancing thoroughness with clarity, always prioritizing the information needed for proper implementation.