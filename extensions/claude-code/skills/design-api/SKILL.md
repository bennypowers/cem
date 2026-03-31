---
name: design-api
description: >
  Design a custom element's public API e.g. attributes, properties, events,
  slots, CSS parts and custom properties. Use when the user asks to
  "design an element", "plan a component API", "create element API",
  "design component interface", or discusses naming conventions, API surface,
  or element architecture.
tools: Read
---

# Element API Design

Guide the design of a custom element's public API, ensuring consistency with existing elements and CEM conventions.

## Workflow

### Phase 1: Understand the Element

Ask the user (if not already provided):
1. What is the element's purpose?
2. What content does it render or contain?
3. What interactions does it support?
4. Are there existing elements it should be consistent with?

### Phase 2: Survey Existing Conventions

Read the existing elements to understand naming patterns:

```text
cem://elements
```

For elements the user mentions as related or similar, read their full APIs:

```text
cem://element/{tagName}
```

Also read any project-level guidelines:

```text
cem://guidelines
```

Note existing patterns for:
- **Tag naming**: prefix conventions, naming style
- **Attribute naming**: kebab-case patterns, boolean naming (e.g. `disabled` vs `is-disabled`)
- **Event naming**: prefix patterns, past tense vs present
- **Slot naming**: `header`, `footer`, `actions` patterns
- **CSS part naming**: conventions for internal structure
- **CSS custom property naming**: `--element-property` patterns

### Phase 3: Design the API

Work through each API surface systematically:

#### Tag Name
- Follow the project's tag prefix convention
- Name should describe what the element *is*, not what it *does*
- Avoid overly generic names

#### Attributes & Properties
- Use kebab-case for attribute names
- Use camelCase for corresponding property names
- Boolean attributes should be positive (`disabled`, not `enabled`)
- Enum attributes should list all valid values
- Consider which attributes reflect to properties and vice versa
- Provide sensible defaults
- Document type constraints
- Prefer CSS states to reflected state attributes

#### Slots
- Use the default slot for primary content
- Name slots by their semantic role (`header`, `actions`), not position (`top`, `bottom`)
- Write content guidelines for each slot
- Consider what content types are appropriate

#### Events
- Name events in past tense for state changes (`opened`, `closed`)
- Name events in present tense for ongoing actions (`input`, `scroll`)
- Use the element prefix if events could conflict with native events
- Design event detail types with the minimum necessary data
- Document what triggers each event
- Prefer to subclass `Event` for type safety, rather than to ship `CustomEvent`

#### CSS Custom Properties
- Follow `--element-name-property` naming convention
- Provide defaults that work in both light and dark themes
- Group related properties (color, spacing, typography)
- Consider which properties consumers will most likely want to override
- Prefer to use existing design tokens rather than to add element-specific props

#### CSS Nesting
- Use native CSS nesting for descendant selectors within shadow DOM
- Don't nest `:host` attribute selectors with `&` -- `&([attr])` inside `:host {}` is invalid because `&` substitutes `:host` literally, making `&([attr])` parse as a function call rather than an attribute selector
- Keep `:host([variant="..."])` blocks flat at the top level; nesting is valid for descendants within those blocks

```css
/* WRONG: &([attr]) is invalid */
:host {
  &([variant="primary"]) { ... }
}

/* RIGHT: flat :host selectors, nested descendants */
:host([variant="primary"]) {
  --_color: blue;
}

:host([variant="primary"]) button {
  &:hover { --_color: darkblue; }
  &:active { --_color: navy; }
}
```

#### CSS Parts
- Name parts after the internal element they expose (`button`, `label`, `icon`)
- Only expose parts that consumers need to style
- Don't expose implementation details that may change

#### CSS States
- Use custom states for element-specific states (`:state(loading)`, `:state(empty)`)
- Map to existing CSS pseudo-classes where appropriate
- Document state transitions

### Phase 4: Source Documentation

When presenting the designed API, show how each part should be documented
in source code.

!`cat ${CLAUDE_SKILL_DIR}/html-comment-syntax.md`

### Phase 5: Consistency Check

Compare the proposed API against existing elements:
- Are similar concepts named the same way?
- Do shared attribute names have compatible types?
- Are event patterns consistent?
- Do CSS property names follow the same scheme?

### Phase 6: Output

Present the designed API as a structured summary:

```markdown
## `<prefix-element-name>` API Design

### Description
[One-paragraph purpose statement]

### Attributes
| Name | Type | Default | Description |
|------|------|---------|-------------|
| `variant` | `'primary' \| 'secondary'` | `'primary'` | Visual style variant |

### Slots
| Name | Description | Content Guidelines |
|------|-------------|--------------------|
| (default) | Primary content | Text or inline elements |

### Events

| Name | Type | Description |
|------|-------------|-------------|
| `change` | `MyChangeEvent` | Fired when value changes |

### CSS Custom Properties
| Name | Default | Description |
|------|---------|-------------|
| `--prefix-element-color` | `light-dark(#000, #fff)` | Primary text color |

### CSS Parts
| Name | Description |
|------|-------------|
| `button` | The internal button element |

### CSS States
| Name | Description |
|------|-------------|
| `loading` | Applied while async content loads |
```

Optionally, generate a skeleton class definition or manifest snippet if the user wants to start implementing.

## Design Principles

- **Minimal API surface**: Don't add attributes "just in case" — each should serve a clear purpose
- **Consistency over novelty**: Match existing patterns in the project, even if you'd name things differently in isolation
- **Attributes for HTML, properties for JS**: Attributes should be serializable; complex data goes through properties only
- **Slots over attributes for rich content**: If content could be more than a plain string, use a slot
- **CSS custom properties for theming, parts for structure**: Properties for colors/spacing, parts for layout overrides
- **Clarify**: If an API can't be clearly and succinctly described in plain English, that likely indicates a design flaw.
