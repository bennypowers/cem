---
title: Writing Effective Descriptions for AI
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

## Description Length Limits

The CEM MCP server enforces a **2000 character limit** on description fields to ensure optimal AI performance and prevent abuse. Descriptions exceeding this limit are automatically truncated with "..." appended.

This limit applies to:
- Element descriptions
- Attribute descriptions
- Slot descriptions
- CSS property descriptions
- Event descriptions
- CSS part descriptions

**Best Practice**: Most effective descriptions are much shorter than 2000 characters. Aim for 200-400 characters for core information, using the full limit only for complex components requiring extensive guidance.

### Customizing Length Limits

For projects requiring different description length limits, you can configure this in your CEM configuration file:

```yaml
mcp:
  maxDescriptionLength: 5000
```

Or use the command line flag:

```bash
cem mcp --max-description-length 5000
```

## Declarative Tool Architecture

The CEM MCP server uses a **declarative tool framework** that transforms your manifest data into intelligent, context-aware responses for AI systems. Understanding this architecture helps you write more effective descriptions.

### How the Declarative Framework Works

The MCP server uses a data-driven approach where tools are defined through YAML configuration rather than hardcoded logic:

1. **Data Fetchers**: Extract specific information from your manifests using JSON path queries
2. **Template Rendering**: Combine manifest data with Go templates to create rich, contextual responses
3. **Schema Integration**: Provide JSON schema definitions that help AI understand component constraints
4. **Intelligent Context**: Present comprehensive information while letting AI make smart decisions

### What This Means for Your Descriptions

The declarative framework **amplifies** the impact of your manifest descriptions by:

- **Contextualizing** your descriptions with related schema information
- **Combining** element descriptions with attribute, slot, and CSS property details
- **Presenting** comprehensive usage patterns derived from your component API definitions
- **Enabling** AI to understand relationships between different parts of your component

**Key insight**: Your descriptions don't work in isolation. They're combined with type information, schema definitions, and usage patterns to create comprehensive guidance that helps AI make intelligent decisions about component usage.

### Writing for the Declarative Framework

Since the framework combines multiple data sources, focus your descriptions on:

1. **Purpose and Context**: What the component/attribute/slot is **for**, not just what it **is**
2. **Usage Guidelines**: When and how to use it effectively
3. **Relationships**: How it works with other component features
4. **Constraints**: What to avoid or be careful about

The framework will automatically combine this with:
- Type definitions and valid values
- Schema descriptions and relationships
- Usage patterns from your component API
- Cross-component consistency information

## General Principles

### 1. Use RFC 2119 Keywords for Actionable Guidelines

The CEM MCP server extracts actionable guidelines from your descriptions by looking for [RFC 2119](https://tools.ietf.org/html/rfc2119) keywords. This creates structured guidance that AI systems can use for intelligent code generation and validation.

The guideline extraction system looks for these keywords in your descriptions:

- **MUST** / **must** - Required behavior or values
- **SHOULD** / **should** - Recommended practices
- **AVOID** / **avoid** - Discouraged patterns
- **USE** / **use** - Preferred approaches

#### Effective RFC 2119 Usage

You can use either uppercase or lowercase - both work equally well:

```typescript
/**
 * Button size affecting touch targets and visual hierarchy.
 * Must be 'large' for primary mobile actions to meet accessibility requirements.
 * Should use 'medium' for standard desktop interfaces.
 * Avoid 'small' for important actions as it reduces accessibility.
 */
@property()
size: 'small' | 'medium' | 'large' = 'medium'
```

This generates both type suggestions ('small', 'medium', 'large') and guideline recommendations from the must/should/avoid statements.

#### Less Effective Patterns

```typescript
/**
 * This property controls the button size. Different sizes are available
 * for different use cases. Pick the one that works best.
 */
@property()
size: 'small' | 'medium' | 'large' = 'medium'
```

While the type information is extracted, no guidelines are generated because there are no RFC 2119 keywords.

### 2. Write for Both Humans and AI

Good descriptions serve dual purposes:
- **Human developers** need clear, concise information for manual coding
- **AI systems** need structured, actionable guidance for automatic generation

<figure class="do">
  <blockquote>A navigation button that supports keyboard interaction and screen readers. Use for primary actions in forms and dialogs.</blockquote>
  <figcaption>✅ Good: Clear and actionable</figcaption>
</figure>

<figure class="dont">
  <blockquote>A custom button element with click handling</blockquote>
  <figcaption>❌ Poor: Too brief and technical</figcaption>
</figure>

### 3. Include Usage Context and Guidelines

AI systems benefit from understanding **when** and **how** to use components:

<figure class="do">
  <blockquote>A card container for grouping related content. Should include a header slot for titles and supports elevation levels 0-5 for visual hierarchy. Use sparingly to avoid visual clutter.</blockquote>
  <figcaption>✅ Good: Includes context and guidelines</figcaption>
</figure>

<figure class="dont">
  <blockquote>A card component with slots</blockquote>
  <figcaption>❌ Poor: Just describes what it is</figcaption>
</figure>

### 4. Emphasize Accessibility Requirements

Since accessibility is a core priority, include accessibility guidance in descriptions:

<figure class="do">
  <blockquote>An interactive toggle button that announces state changes to screen readers. Must include aria-pressed attribute and supports keyboard activation with Space or Enter keys.</blockquote>
  <figcaption>✅ Good: Accessibility-focused</figcaption>
</figure>

<figure class="dont">
  <blockquote>A toggle button component</blockquote>
  <figcaption>❌ Poor: No accessibility information</figcaption>
</figure>

### 5. Combine Types and Guidelines for Maximum AI Effectiveness

The most AI-friendly approach combines clear TypeScript union types with guideline-rich descriptions:

```typescript
/**
 * Visual style variant affecting semantic meaning and accessibility.
 * Use 'primary' for main call-to-action buttons in forms and dialogs.
 * Use 'secondary' for supporting actions that complement primary actions.
 * Use 'danger' for destructive actions that require user confirmation.
 * Avoid using 'danger' for non-destructive actions to prevent confusion.
 * Each variant provides appropriate color contrast for accessibility.
 */
@property()
variant: 'primary' | 'secondary' | 'danger' | 'ghost' = 'primary'
```

This provides:
- **Type suggestions**: AI understands all valid values from the union type
- **Usage guidelines**: When to use each variant (extracted from RFC 2119 keywords)
- **Accessibility context**: Color contrast considerations
- **Anti-patterns**: What to avoid and why

## Declarative Tool Integration

The declarative framework provides several specialized tools that present your descriptions in focused, context-aware formats. Understanding how these tools work helps you write descriptions that shine in each context.

### Element-Focused Tools

The framework includes specialized tools for different aspects of your components:

- **`element_details`**: Comprehensive overview with all APIs and usage guidance
- **`element_attributes`**: Focused attribute documentation with type constraints
- **`element_slots`**: Content guidelines and accessibility considerations for slots
- **`element_events`**: Event triggers, data payloads, and integration patterns
- **`element_styling`**: CSS customization with design system integration

### How Tools Present Your Content

Each tool combines your descriptions with relevant schema information:

```
Your Description + Type Information + Schema Context = Rich AI Guidance
```

**Example**: For an attribute description like "Button size affecting accessibility", the `element_attributes` tool combines this with:
- Union type values (`'small' | 'medium' | 'large'`)
- Default value information (`'medium'`)
- Schema descriptions for size concepts
- Cross-component consistency patterns

**Result**: AI receives comprehensive guidance like:
> "Button size affects touch targets and accessibility. Use 'small' for compact layouts, 'medium' for standard interfaces, 'large' for mobile-first actions. Each size ensures minimum touch target requirements for accessibility compliance."

### Writing for Tool Specialization

Since tools focus on specific aspects, tailor your descriptions accordingly:

#### For Element Descriptions (used by `element_details`)
Focus on overall purpose, main use cases, and key relationships:
```typescript
/**
 * An accessible card container for grouping related content with semantic markup.
 * Use for dashboard widgets, content previews, and information grouping.
 * Supports responsive design and maintains proper heading hierarchy.
 */
```

#### For Attribute Descriptions (enhanced by `element_attributes`)
Focus on purpose, constraints, and usage context:
```typescript
/**
 * Controls card prominence in visual hierarchy. Use 'flat' for embedded content,
 * 'raised' for interactive cards, 'floating' for modal-style overlays.
 */
elevation: 'flat' | 'raised' | 'floating'
```

#### For Slot Descriptions (used by `element_slots`)
Focus on content types, accessibility, and relationships:
```typescript
/**
 * Card header area for titles and actions. Should contain heading elements (h2-h4)
 * to maintain document structure. Avoid interactive elements that conflict with
 * card-level actions. Announced first by screen readers.
 */
```

The declarative framework ensures that your focused descriptions are automatically enhanced with relevant schema information, type constraints, and cross-component patterns to provide comprehensive AI guidance.

## Element Descriptions

Element descriptions should answer: **What is this component and when should I use it?**

### Structure Your Element Descriptions

Use this pattern for comprehensive element descriptions:

```
[Component Purpose] [Accessibility Features] [Usage Guidelines] [Context Information]
```

### Examples by Component Type

#### Interactive Components (Buttons, Inputs, Controls)

> A primary action button with built-in loading states and accessibility
> features. Automatically includes proper ARIA attributes and keyboard support.
> Use for important actions like form submission, confirmation dialogs, and
> primary navigation. Supports variant styling for different contexts.

**Key elements:**
- Purpose: "primary action button"
- Accessibility: "ARIA attributes and keyboard support"  
- Usage context: "form submission, confirmation dialogs"
- Features: "loading states", "variant styling"

#### Layout Components (Cards, Containers, Grids)

> A flexible content container with semantic markup for grouping related
> information. Includes slots for header, body, and footer content with proper
> heading hierarchy. Use for dashboard widgets, product listings, and content
> previews. Supports responsive design and maintains accessibility across
> screen sizes.

**Key elements:**
- Purpose: "content container for grouping"
- Structure: "slots for header, body, footer"
- Semantic info: "proper heading hierarchy"
- Usage examples: "dashboard widgets, product listings"
- Responsive: "maintains accessibility across screen sizes"

#### Navigation Components (Menus, Tabs, Breadcrumbs)

> An accessible tab navigation component that manages focus and announces
> selections to screen readers. Implements WCAG navigation patterns with arrow
> key support and automatic ARIA relationships. Use for organizing content into
> logical sections within a single view.

**Key elements:**
- Accessibility: "manages focus", "announces selections", "WCAG patterns"
- Interaction: "arrow key support", "ARIA relationships"
- Usage context: "organizing content into logical sections"

## Attribute Descriptions

Attribute descriptions should answer: **What does this control and how should I use it?**

### Include Type Information and Valid Values

```ts
@property() variant: 'primary' | 'secondary' | 'danger' | 'ghost'
```

<figure class="do">
  <blockquote>
    Button visual style and semantic meaning. Use 'primary' for main actions,
    'secondary' for supporting actions, 'danger' for destructive actions, and
    'ghost' for subtle actions. Affects color contrast and accessibility
    announcements.
  </blockquote>
  <figcaption>✅ Good: Complete attribute information</figcaption>
</figure>

<figure class="dont">
  <blockquote>Button style</blockquote>
  <figcaption>❌ Poor: Minimal information</figcaption>
</figure>

### Explain Impact and Usage Context

```ts
@property() size: 'small' | 'medium' | 'large'
```

<figure class="do">
  <blockquote>
    Controls button dimensions and touch target size for accessibility. Use
    'small' for compact layouts, 'medium' for standard interfaces, 'large'
    for mobile-first or prominent actions. Ensures minimum 44px touch target on
    mobile devices.
  </blockquote>
  <figcaption>✅ Good: Explains impact and context</figcaption>
</figure>

### Include Default Behavior

```ts
@property({ type: Boolean }) loading = false
```

<figure class="do">
  <blockquote>
    Shows loading spinner and disables interaction during async operations.
    Automatically announces loading state to screen readers. Set to true during
    form submission or data fetching to prevent duplicate actions.
  </blockquote>
  <figcaption>✅ Good: Explains default and when to change</figcaption>
</figure>

## Slot Descriptions

Slot descriptions should answer: **What content belongs here and how should it be structured?**

### Describe Expected Content Types

For a `header` slot:
<figure class="do">
  <blockquote>
    Card header content, typically a heading element (h2-h4) or title text.
    Should establish the card's purpose and maintain proper heading hierarchy.
    Avoid interactive elements that might conflict with card-level actions.
  </blockquote>
  <figcaption>✅ Good: Specific content guidance</figcaption>
</figure>

<figure class="dont">
  <blockquote>Header content</blockquote>
  <figcaption>❌ Poor: Generic description</figcaption>
</figure>

### Include Accessibility Considerations

For a card element's `action` slot:

<figure class="do">
  <blockquote>
    Primary action area for buttons or links. Content should be descriptive for
    screen readers (avoid 'click here' or 'read more'). Interactive elements
    will receive proper focus management within the card context.
  </blockquote>
  <figcaption>✅ Good: Accessibility-aware slot description</figcaption>
</figure>

### Explain Content Relationships

For a card element's default slot:
<figure class="do">
  <blockquote>
    Main card content area for body text, lists, or media. Content should
    complement the header and any action elements. Will be announced as the
    primary card content by screen readers after the header.
  </blockquote>
  <figcaption>✅ Good: Explains how content works together</figcaption>
</figure>

## CSS Property Descriptions

CSS property descriptions should answer: **What visual aspect does this control and how does it fit the design system?**

### Link to Design System Concepts

```css
--card-elevation: 2px;
```

<figure class="do">
  <blockquote>
    Controls card shadow depth using design system elevation scale. Maps to
    elevation tokens 0-5 where 0 is flat and 5 is highest prominence. Affects
    both visual hierarchy and accessibility by indicating interactive priority.
  </blockquote>
  <figcaption>✅ Good: Design system integration</figcaption>
</figure>

<figure class="dont">
  <blockquote>Card shadow depth</blockquote>
  <figcaption>❌ Poor: No context or guidance</figcaption>
</figure>

### Explain Visual and Functional Impact

```css
--button-radius: 4px;
```

<figure class="do">
  <blockquote>
    Border radius for button corners affecting visual style and brand
    consistency. Use design system radius tokens (<code>--radius-sm</code>,
    <code>--radius-md</code>, <code>--radius-lg</code>) for consistency.
    Impacts touch target perception on mobile devices.
  </blockquote>
  <figcaption>✅ Good: Visual and functional explanation</figcaption>
</figure>

<figure class="dont">
  <blockquote>Button border radius</blockquote>
  <figcaption>❌ Poor: No usage guidance</figcaption>
</figure>

## Event Descriptions

Event descriptions should answer: **When does this fire and what data does it provide?**

### Describe Trigger Conditions and Data

For a `selection-change` event:

<figure class="do">
  <blockquote>
    Fired when user selects or deselects items, including keyboard and mouse
    interactions. Event detail contains <code>selectedItems</code> array and
    <code>previousSelection</code> for accessibility announcements. Debounced
    to prevent excessive firing during rapid selection changes.
  </blockquote>
  <figcaption>✅ Good: Complete event information</figcaption>
</figure>

<figure class="dont">
  <blockquote>Fired when selection changes</blockquote>
  <figcaption>❌ Poor: No detail about triggers or data</figcaption>
</figure>

### Include Accessibility Information

For an `expand-toggle` event:

<figure class="do">
  <blockquote>
    Fired when collapsible content expands or collapses. Event detail includes
    expanded state and target element for screen reader announcements. Use to
    coordinate ARIA attributes and manage focus when content visibility changes.
  </blockquote>
  <figcaption>✅ Good: Accessibility integration</figcaption>
</figure>

<figure class="dont">
  <blockquote>Fired when expanding or collapsing</blockquote>
  <figcaption>❌ Poor: No accessibility context</figcaption>
</figure>

## CSS Parts and States

### CSS Parts: Describe Styling Purpose

```html
<input part="input" type="text">
```

<figure class="do">
  <blockquote>
    The input field element for styling text appearance, borders, and focus
    states. Inherits form-level styling but can be customized for specific use
    cases. Maintain sufficient color contrast for accessibility.
  </blockquote>
  <figcaption>✅ Good: Clear styling guidance</figcaption>
</figure>

<figure class="dont">
  <blockquote>Input field part</blockquote>
  <figcaption>❌ Poor: No styling guidance or context</figcaption>
</figure>

### CSS States: Explain Behavioral Context

```typescript
@property({ type: Boolean, reflect: true }) loading = false;
```

<figure class="do">
  <blockquote>
    Applied during async operations when the component is waiting for data or
    processing requests. Use to style loading indicators and adjust interaction
    states. Should provide visual feedback while maintaining accessibility.
  </blockquote>
  <figcaption>✅ Good: Behavioral context</figcaption>
</figure>

<figure class="dont">
  <blockquote>Loading state</blockquote>
  <figcaption>❌ Poor: No behavioral guidance or usage context</figcaption>
</figure>

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

```typescript
/**
 * An accessible multi-select dropdown with search, keyboard navigation, and bulk operations.
 * Implements WCAG 2.1 AA standards with comprehensive screen reader support and keyboard
 * interaction patterns. Use for selection from large datasets where users need to choose
 * multiple related items, such as tag selection, user assignment, or category filtering.
 *
 * @event selection-change - Fired when selection changes through any interaction method (click, keyboard, programmatic). Event detail includes complete selection state and the specific change that occurred. Use for form validation, dependent field updates, or analytics tracking. Debounced during rapid keyboard navigation to prevent excessive firing.
 * @event search-input - Fired when user types in search field (when searchable is true). Event detail contains current search query and number of matching options. Use for external filtering, analytics, or implementing custom search behavior. Fires after built-in filtering is applied.
 * @event dropdown-toggle - Fired when dropdown opens or closes. Event detail indicates new state and what triggered the change. Use for coordinating with external UI elements, implementing custom positioning, or managing page-level focus. Essential for modal-style behaviors.
 * @event max-reached - Fired when user attempts to select beyond max-selections limit. Use to show helpful error messages or suggest alternative actions. Event is fired before selection is rejected, allowing for custom handling or user education.
 *
 * @cssstate open - Applied when dropdown is expanded. Use for styling state transitions, positioning adjustments, or coordinating with external elements. Helps manage z-index and overlay behaviors.
 * @cssstate searching - Applied when user is actively typing in search field. Use for loading indicators, highlighting matched text, or adjusting dropdown layout. Provides visual feedback during search operations.
 * @cssstate max-reached - Applied when maximum selection limit is reached. Use for visual feedback, disabling remaining options, or showing helpful messaging. Should clearly communicate the constraint to users.
 * @cssstate required-empty - Applied when component is required but has no selections. Use for error styling, validation messaging, or form submission prevention. Should clearly indicate the validation state.
 */
@customElement('multi-select')
export class MultiSelectElement extends LitElement {
  static styles = css`
    :host {
      /**
       * Controls component width. Use design system width tokens or specific
       * measurements. Affects dropdown positioning and text wrapping behavior.
       * Consider content length and container constraints when setting.
       */
      --multi-select-width: auto;

      /**
       * Maximum height for the dropdown options list. When exceeded, the list
       * becomes scrollable with keyboard navigation support. Use viewport-relative
       * units for responsive behavior. Recommended range: 200px to 60vh for optimal usability.
       */
      --multi-select-max-height: 300px;

      /**
       * Background color for selected item tags (when selection-display='tags').
       * Should maintain sufficient contrast with text color for accessibility.
       * Use semantic colors from design system that indicate selection state clearly.
       */
      --selection-tag-color: var(--color-primary-100);

      /**
       * Background color for option hover state. Must meet WCAG contrast requirements
       * against option text. Provides visual feedback for keyboard and mouse navigation.
       * Should be distinct from selection and focus states.
       */
      --option-hover-bg: var(--color-gray-100);

      /**
       * Border radius for dropdown container. Use design system radius tokens for
       * consistency. Affects visual integration with page design and brand consistency.
       * Consider relationship to trigger element styling.
       */
      --dropdown-border-radius: var(--radius-md);
    }
  `;

  /**
   * Placeholder text shown when no selections are made. Should be descriptive
   * and action-oriented (e.g., 'Select team members' rather than 'Choose options').
   * Announced by screen readers when the component receives focus.
   */
  @property({ type: String })
  placeholder = 'Select options...';

  /**
   * Enables search/filter functionality within the dropdown. When true, users can
   * type to filter available options, improving usability for large option sets.
   * Search is case-insensitive and matches partial strings. Essential for lists
   * with more than 10 options.
   */
  @property({ type: Boolean })
  searchable = false;

  /**
   * Maximum number of items that can be selected. When limit is reached, remaining
   * options become disabled with appropriate ARIA states. Use to prevent overwhelming
   * interfaces or enforce business rules. Value of 0 means unlimited selections.
   */
  @property({ type: Number, attribute: 'max-selections' })
  maxSelections = 0;

  /**
   * Marks the component as required for form validation. Displays visual indicator
   * and updates ARIA attributes for screen readers. Prevents form submission when
   * no selections are made. Should be paired with clear error messaging.
   */
  @property({ type: Boolean })
  required = false;

  /**
   * Disables all interaction with the component. Sets appropriate ARIA states and
   * visual styling. Use for read-only states or when dependent form fields haven't
   * been completed. Maintains selection visibility for context.
   */
  @property({ type: Boolean })
  disabled = false;

  /**
   * Controls layout density and information display. 'default' shows standard spacing
   * and basic option text. 'compact' reduces spacing for dense interfaces. 'detailed'
   * shows additional option metadata and descriptions. Choose based on interface
   * density and content complexity.
   */
  @property({ type: String })
  variant: 'default' | 'compact' | 'detailed' = 'default';

  /**
   * How selected items are displayed in the closed state. 'tags' shows individual
   * removable chips (best for <= 5 selections). 'count' shows summary like '3 selected'
   * (best for large selections). 'list' shows comma-separated names (best for 2-4
   * selections). Automatically adjusts for accessibility and screen space.
   */
  @property({ type: String, attribute: 'selection-display' })
  selectionDisplay: 'tags' | 'count' | 'list' = 'tags';

  render() {
    return html`
      <!-- description: >
             Main clickable area that opens the dropdown. Style for brand
             consistency and clear affordance. Includes focus indicators
             and disabled states. Should maintain minimum 44px touch
             target for mobile accessibility. -->
      <div part="trigger">
        <!-- description: >
               Container for displaying selected items. Layout varies based on
               selection-display attribute. Style for clear visual hierarchy
               and easy scanning. Handles overflow with appropriate text
               truncation. -->
        <div part="selections">
          <!-- Selected items display -->
        </div>
      </div>

      <!-- summary: Dropdown container including search, options, and footer areas.
           description: >
             Handles positioning, shadows, and border styling. Consider z-index
             stacking context and responsive positioning behavior. -->
      <div part="dropdown" ?hidden="${!this.open}">
        ${!this.searchable ? '' : html`
          <!-- summary: Search input field within dropdown (when searchable).
               description: >
                 Style for clear focus states and placeholder visibility. Should
                 integrate seamlessly with dropdown design while remaining
                 clearly functional. -->
          <input part="search-input" type="search" placeholder="Search options...">`}

        <!-- summary: Optional header content above the options list
             description: >
               Use for category labels, action buttons (like 'Select All'), or
               contextual help. Should not contain form controls that conflict
               with option selection. Announced as part of dropdown context by
               screen readers. -->
        <slot name="header"></slot>

        <div class="options-container">
          <!-- summary: Default slot for option elements.
               description: >
                 Should contain option elements with value attributes and
                 descriptive text content. Options support disabled state and
                 optional description text. Maintain semantic markup for screen
                 readers - each option should be self-contained and descriptive. -->
          <!-- description: >
                 Individual option elements within the list. Style for clear
                 selection states, hover feedback, and disabled appearance.
                 Must maintain accessibility contrast requirements across all
                 states. -->
          <slot part="option"></slot>

          <!-- summary: Content shown when no options match
               description: >
                 Content shown when no options match current search or when options
                 list is empty. Should provide helpful guidance or alternative
                 actions. Include clear messaging about search refinement or
                 data loading states. Will be announced by screen readers when
                 displayed. -->
          <slot name="empty-state"></slot>
        </div>

        <!-- summary: Footer content below options list
             description: >
               Ideal for summary information ('X of Y selected'), action
               buttons, or links to manage the option set externally. Maintains
               focus context within the dropdown for keyboard navigation. -->
        <slot name="footer"></slot>
      </div>

      <!-- summary: Individual selection chips when selection-display='tags'
           description: >
             Individual selection chips (when selection-display='tags').
             Include close button styling and hover states. Should be easily
             scannable and clearly removable for keyboard users. -->
      <div part="selected-tag" ?hidden="${!this.selected}"></div>
    `;
  }
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

Following these comprehensive documentation patterns ensures your custom elements manifest provides rich, actionable information that enables both human developers and AI systems to use your components correctly, accessibly, and effectively. The key is balancing thoroughness with clarity, always prioritizing the information needed for proper implementation.
