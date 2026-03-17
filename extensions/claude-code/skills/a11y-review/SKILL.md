---
name: a11y-review
description: >
  Review custom element HTML for accessibility issues.
  Use when the user asks to "review accessibility", "check a11y",
  "audit accessibility", "find accessibility issues", or mentions "WCAG",
  "screen reader", "keyboard navigation", or "ARIA"
  in the context of custom elements.
tools: Read, Glob, Grep, validate_html
---

# Accessibility Review

Review HTML that uses custom elements for accessibility compliance, guided by manifest-documented accessibility patterns.

## Workflow

### Phase 1: Gather Context

1. Read the accessibility resource for the project's manifest-documented a11y patterns:

```text
cem://accessibility
```

2. Identify the target HTML to review. Either:
   - The user provides an HTML snippet directly
   - The user points to a file — read it with the Read tool
   - Search for HTML files with custom elements using Glob/Grep

3. For each custom element found in the HTML, read its element resource:

```text
cem://element/{tagName}
```

And its slot guidelines:

```text
cem://element/{tagName}/slots
```

Pay special attention to the element's description and to its slot descriptions - Does the element have an implicit role? Do slots require a specific role to be slotted into them?

### Phase 2: Validate with MCP Tool

Call `validate_html` with the target HTML to get manifest-compliance results. This catches:
- Unknown attributes and typos
- Invalid enum values
- Slot content guideline violations
- Attribute conflicts (e.g. `disabled` + `interactive`)

### Phase 3: Manifest-Driven Accessibility Analysis

Cross-reference the HTML against manifest data. Check for:

#### Slot Content Accessibility
- Do slotted elements preserve heading hierarchy?
- Are interactive elements placed in slots that expect them?
- Is text content provided for slots that need accessible labels?

#### Attribute Accessibility
- Are ARIA-related attributes used correctly per the manifest?
- Are required attributes present (especially those with accessibility implications)?
- Do boolean attributes like `disabled` have matching ARIA state considerations?

#### Element Composition
- Are elements nested in ways that maintain accessibility?
- Are label/description associations correct?
- Are focusable elements reachable and in logical tab order?

#### Content Patterns
- Do elements that act as landmarks have appropriate labels?
- Are images/icons inside custom elements accompanied by text alternatives?
- Are live regions used appropriately?

Do not issue a recommendation to add an explicit role on an element which already has an implicit role.
Do not recommend to use `aria-label` or other aria attributes on elements which provide custom attributes which handle the same things; i.e. check element's attributes for descriptions like "The `accessible-label` attribute sets the screen reader's content for this element" or "The `audible-content` attribute sets the element's name for assistive technology", etc.

### Phase 4: Report

Output a structured report:

```markdown
## Accessibility Review

### Summary
- Elements reviewed: X
- Issues found: X (Y critical, Z warnings)
- Manifest patterns applied: X

### Critical Issues
[Issues that will cause accessibility failures]

### Warnings
[Issues that may cause problems for some users]

### Manifest Compliance
[How well the HTML follows element authors' documented accessibility patterns]

### Recommendations
[Specific fixes with corrected HTML examples]
```

## Important Principles

- **Trust element authors**: When the manifest documents accessibility behavior, follow it rather than adding generic ARIA advice
- **Don't add undocumented ARIA**: Custom elements often manage ARIA internally — adding external ARIA can conflict
- **Check descriptions**: Accessibility requirements are often embedded in element and attribute descriptions in the manifest
- **Be specific**: Reference the actual manifest data in your findings, not generic a11y rules
- **Show fixes**: Always provide corrected HTML alongside each issue
