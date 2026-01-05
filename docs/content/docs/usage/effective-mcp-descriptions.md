---
title: Effective Writing for AI
weight: 90
---

The [Model Context Protocol][mcp] integration uses your [Custom Elements 
Manifest][customelementsjson] descriptions to help AI systems generate correct 
HTML, suggest appropriate component usage, validate accessibility compliance, 
and provide contextual recommendations. After [setting up MCP][mcpsetup], AI 
assistants read your manifest to understand component relationships and 
hierarchies, so well-written descriptions enable intelligent code generation 
while poor descriptions lead to incorrect or suboptimal output. This complements 
the [development workflow][workflow] by providing AI assistants with the same 
component knowledge you document for humans.

The CEM MCP server uses a declarative tool framework that transforms your 
manifest data into context-aware responses for AI systems. Your JSDoc 
descriptions don't work in isolation—they're combined with type information, 
schema definitions, and usage patterns to create comprehensive guidance. Use 
[RFC 2119 keywords][rfc2119] (MUST, SHOULD, AVOID) in descriptions to create 
structured guidelines that AI systems extract for validation and code 
generation.

## Description Length Limits

The CEM MCP server enforces a 2000 character limit on all description fields 
(elements, attributes, slots, CSS properties, events, and parts) to ensure 
optimal AI performance. Descriptions exceeding this limit are automatically 
truncated with "..." appended. Most effective descriptions are much shorter—aim 
for 200-400 characters for core information, using the full limit only for 
complex components requiring extensive guidance.

Configure custom limits in `.config/cem.yaml` with `mcp.maxDescriptionLength: 
5000` or use the `--max-description-length` flag.

## Writing Principles

Focus descriptions on purpose and context (what the component/attribute/slot is 
**for**, not just what it **is**), usage guidelines (when and how to use it 
effectively), relationships (how it works with other features), and constraints 
(what to avoid). The MCP server automatically combines your descriptions with 
type definitions, valid values, schema descriptions, and cross-component 
consistency information to create comprehensive AI guidance.

## Use RFC 2119 Keywords

The MCP server extracts actionable guidelines by looking for RFC 2119 keywords 
(MUST/must, SHOULD/should, AVOID/avoid, USE/use) in your descriptions. For 
example, "Must be 'large' for primary mobile actions to meet accessibility 
requirements" creates structured guidance that AI uses for validation and code 
generation. Descriptions without these keywords provide only type information 
without behavioral guidance.

## Writing Effective Descriptions

Write descriptions that serve both human developers and AI systems by being 
clear and actionable. Include usage context explaining **when** and **how** to 
use components—for example, "A card container for grouping related content. 
Should include a header slot for titles and supports elevation levels 0-5 for 
visual hierarchy" is better than "A card component with slots."

Emphasize accessibility requirements since they're a core priority. Descriptions 
like "An interactive toggle button that announces state changes to screen 
readers. Must include aria-pressed attribute and supports keyboard activation 
with Space or Enter keys" help AI generate accessible code.

Combine TypeScript union types with guideline-rich descriptions for maximum 
effectiveness—AI extracts type suggestions from the union, usage guidelines from 
RFC 2119 keywords, accessibility context from your descriptions, and 
anti-patterns from AVOID statements.

### How `cem mcp` Presents Your Content

`cem mcp` provides AI agents with your content in the form of resources.
Each resource combines your descriptions with relevant schema information:

```
Your Description + Type Information + Schema Context = Rich AI Guidance
```

**Example**: For an attribute description like "Button size affecting 
accessibility", the `cem://element/{tagName}/attributes` resource combines this 
with:
- Union type values (`'small' | 'medium' | 'large'`)
- Default value information (`'medium'`)
- Schema descriptions for size concepts
- Cross-component consistency patterns

**Result**: AI receives comprehensive guidance like:

> Button size affects touch targets and accessibility. Use 'small' for compact 
> layouts, 'medium' for standard interfaces, 'large' for mobile-first actions. 
> Each size ensures minimum touch target requirements for accessibility 
> compliance.

### Writing for Resource Specialization

Since resources focus on specific aspects, tailor your descriptions accordingly:

#### For Element Descriptions (used by `element_details`)
Focus on overall purpose, main use cases, and key relationships:

> An accessible card container for grouping related content with semantic 
> markup. Use for dashboard widgets, content previews, and information 
> grouping. Supports responsive design and maintains proper heading hierarchy.

#### For Attribute Descriptions (enhanced by `cem://element/{tagName}/attributes`)
Focus on purpose, constraints, and usage context:

```typescript
/**
 * Controls card prominence in visual hierarchy. Use 'flat' for embedded content,
 * 'raised' for interactive cards, 'floating' for modal-style overlays.
 */
elevation: 'flat' | 'raised' | 'floating'
```

#### For Slot Descriptions (used by `cem://element/{tagName}/slots`)
Focus on content types, accessibility, and relationships:


> Card header area for titles and actions. Should contain heading elements 
> (h2-h4) to maintain document structure. Avoid interactive elements that 
> conflict with card-level actions. Announced first by screen readers.

The declarative framework ensures that your focused descriptions are 
automatically enhanced with relevant schema information, type constraints, and 
cross-component patterns to provide comprehensive AI guidance.

## Element Descriptions

Element descriptions should answer what the component is and when to use it. 
Structure descriptions as: component purpose, accessibility features, usage 
guidelines, and context information. For example, describe interactive 
components by their role and keyboard support, layout components by their 
semantic structure and responsive behavior, and navigation components by their 
WCAG patterns and focus management.

## Attribute Descriptions

Attribute descriptions should explain what the attribute controls and how to use 
it. Include type information, valid values, impact on behavior, and usage 
context. Explain default behavior and when to override defaults.

<figure class="do">
  <blockquote>
    Button visual style and semantic meaning. Use 'primary' for main actions,
    'secondary' for supporting actions, 'danger' for destructive actions, and
    'ghost' for subtle actions. Affects color contrast and accessibility
    announcements.
  </blockquote>
  <figcaption><strong>Good</strong>: Complete attribute information</figcaption>
</figure>

<figure class="dont">
  <blockquote>Button style</blockquote>
  <figcaption><strong>Poor</strong>: Minimal information</figcaption>
</figure>

## Slot Descriptions

Slot descriptions should explain what content belongs in the slot and how it should be structured. Describe expected content types (heading elements, interactive elements, text), accessibility considerations (screen reader announcements, focus management), and content relationships (how slot content works with other slots).

<figure class="do">
  <blockquote>
    Card header content, typically a heading element (h2-h4) or title text.
    Should establish the card's purpose and maintain proper heading hierarchy.
    Avoid interactive elements that might conflict with card-level actions.
  </blockquote>
  <figcaption><strong>Good</strong>: Specific content guidance</figcaption>
</figure>

<figure class="dont">
  <blockquote>Header content</blockquote>
  <figcaption><strong>Poor</strong>: Generic description</figcaption>
</figure>

## CSS Property Descriptions

CSS property descriptions should explain what visual aspect the property controls and how it fits the design system. Link to design system concepts (elevation scales, spacing tokens, color schemes), explain visual and functional impact, and reference design tokens for consistency.

<figure class="do">
  <blockquote>
    Controls card shadow depth using design system elevation scale. Maps to
    elevation tokens 0-5 where 0 is flat and 5 is highest prominence. Affects
    both visual hierarchy and accessibility by indicating interactive priority.
  </blockquote>
  <figcaption><strong>Good</strong>: Design system integration</figcaption>
</figure>

<figure class="dont">
  <blockquote>Card shadow depth</blockquote>
  <figcaption><strong>Poor</strong>: No context or guidance</figcaption>
</figure>

## Event Descriptions

Event descriptions should explain when the event fires and what data it provides. Describe trigger conditions (user interactions, state changes), event detail contents (data structure, accessibility information), and debouncing or throttling behavior.

<figure class="do">
  <blockquote>
    Fired when user selects or deselects items, including keyboard and mouse
    interactions. Event detail contains <code>selectedItems</code> array and
    <code>previousSelection</code> for accessibility announcements. Debounced
    to prevent excessive firing during rapid selection changes.
  </blockquote>
  <figcaption><strong>Good</strong>: Complete event information</figcaption>
</figure>

<figure class="dont">
  <blockquote>Fired when selection changes</blockquote>
  <figcaption><strong>Poor</strong>: No detail about triggers or data</figcaption>
</figure>

## CSS Parts and States

CSS parts descriptions should explain the styling purpose, inheritance behavior, and accessibility requirements for color contrast. CSS states descriptions should explain behavioral context, when the state is applied, and how to provide visual feedback while maintaining accessibility.

<figure class="do">
  <blockquote>
    The input field element for styling text appearance, borders, and focus
    states. Inherits form-level styling but can be customized for specific use
    cases. Maintain sufficient color contrast for accessibility.
  </blockquote>
  <figcaption><strong>Good</strong>: Clear styling guidance</figcaption>
</figure>

<figure class="dont">
  <blockquote>Input field part</blockquote>
  <figcaption><strong>Poor</strong>: No styling guidance or context</figcaption>
</figure>

## Testing Your Descriptions

Test your descriptions using the `cem mcp` server by asking AI to generate HTML using your components, requesting attribute suggestions for different contexts, and verifying that generated examples follow your intended usage patterns. If AI generates inappropriate usage, refine your descriptions for clarity and completeness. The key is balancing thoroughness with clarity, always prioritizing the information needed for proper implementation.

## See Also

- **[AI-Friendly Documentation Example][ai-example]** - Comprehensive multi-select component demonstrating best practices
- **[MCP Integration][mcpsetup]** - Setup instructions for AI assistants
- **[Development Workflow][workflow]** - How MCP fits into the dev cycle
- **[Generate Command][generate]** - JSDoc syntax for documenting components

[mcp]: https://modelcontextprotocol.io/
[customelementsjson]: https://github.com/webcomponents/custom-elements-json
[mcpsetup]: /docs/installation/mcp/
[workflow]: ../workflow/
[rfc2119]: https://www.rfc-editor.org/rfc/rfc2119
[generate]: /docs/reference/commands/generate/
[ai-example]: /examples/ai-friendly-docs/
