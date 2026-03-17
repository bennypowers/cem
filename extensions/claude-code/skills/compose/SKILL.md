---
name: compose
description: >
  Suggest and generate custom element compositions for a UI goal.
  Use when the user asks to "compose elements", "build a layout",
  "combine components", "create a page with", "what elements should I use for",
  or describes a UI they want to build using custom elements.
tools: Read, generate_html, validate_html
---

# Component Composition Advisor

Given a UI goal, recommend which custom elements to use together and generate composed HTML.

## Workflow

### Phase 1: Understand the Goal

Clarify (if not already provided):
1. What UI is the user trying to build? (e.g., "login form", "data table", "card grid")
2. Any constraints? (e.g., specific elements they must use, layout requirements)
3. What level of detail? (quick wireframe-level vs production-ready)

### Phase 2: Survey Available Elements

Read the complete element library:

```text
cem://elements
```

Read project-level guidelines for composition patterns:

```text
cem://guidelines
```

For each candidate element, read its full API to understand capabilities and slot expectations:

```text
cem://element/{tagName}
cem://element/{tagName}/slots
cem://element/{tagName}/attributes
```

### Phase 3: Select Elements

Choose elements based on:

#### Semantic Fit
- Does the element's purpose match the role it would play in the composition?
- Is the element designed for this kind of content?

#### Slot Compatibility
- Can the elements nest correctly via slots?
- Do parent element slots accept the kind of content the child elements produce?
- Are there slot content guidelines that constrain composition?

#### Attribute Harmony
- Do the elements share compatible attribute patterns?
- Can they be configured to work together without conflicts?
- Do variant/size attributes compose sensibly?

#### Event Flow
- Can parent elements respond to child element events?
- Is the event propagation path logical?

### Phase 4: Generate Composed HTML

For each element in the composition, call `generate_html` to produce valid markup, then combine them.

Present the composition as:

```markdown
## Composition: [UI Goal]

### Elements Used

| Element          | Role       | Why                                    |
|------------------|------------|----------------------------------------|
| `<prefix-card>`  | Container  | Provides visual grouping and elevation |
| `<prefix-input>` | Form field | Text input with built-in validation    |


### HTML

[Complete composed HTML with comments explaining structure]

### How It Works
- [Explain the composition: which elements are slotted where]
- [Note any attribute interactions between elements]
- [Describe the event flow for interactive compositions]

### Styling

[CSS showing how to style the composition, using CSS custom properties and parts from the composed elements]

### Alternatives Considered
- [Other element combinations that could work, and why this one is better]
```

### Phase 5: Validate

Call `validate_html` on the final composed HTML to verify:
- All elements are used correctly
- Slot content follows guidelines
- No attribute conflicts
- Accessibility patterns are maintained

Report any issues and fix them before presenting the final result.

## Guidelines

- **Prefer composition over customization**: Use multiple elements together rather than forcing one element to do everything
- **Respect slot guidelines**: If a slot says "text content only", don't put a complex element in it
- **Keep nesting shallow**: Deeply nested compositions are fragile and hard to style
- **Show the minimal composition first**: Start simple, then offer enhanced versions with more elements
- **Consider responsive behavior**: Note if the composition needs different structure at different viewport sizes
- **Document the "why"**: Explain why each element was chosen so the user can make substitutions
