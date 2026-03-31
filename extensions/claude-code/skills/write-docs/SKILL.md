---
name: write-docs
description: >
  Generate documentation for custom elements based on their manifest data.
  Use when the user asks to "write docs", "generate documentation",
  "document this element", "create API docs", or mentions "usage examples" or
  "component documentation".
tools: Read, Glob, Grep, generate_html
---

# Write Element Documentation

Generate documentation for custom elements based on their manifest API surface,
including usage examples and integration guidance.

## Workflow

### Phase 1: Gather Element Data

Read the target element's full manifest data:

```text
cem://element/{tagName}
```

And each sub-resource for complete coverage:

```text
cem://element/{tagName}/attributes
cem://element/{tagName}/slots
cem://element/{tagName}/events
cem://element/{tagName}/css/parts
cem://element/{tagName}/css/custom-properties
cem://element/{tagName}/css/states
```

Also read project-level guidelines for tone and conventions:

```text
cem://guidelines
```

### Phase 1b: Source Comment Syntax

!`cat ${CLAUDE_SKILL_DIR}/html-comment-syntax.md`

### Phase 2: Detect Documentation Conventions

Search the project for existing documentation patterns:

1. Look for existing docs files to match format and structure:
   - Markdown docs in `docs/`, `documentation/`, or alongside source
   - Storybook stories (`*.stories.ts`, `*.stories.js`)
   - README files in component directories
   - JSDoc/TSDoc in source files, including CSS files

2. Match the project's existing doc style:
   - Heading structure and depth
   - Code example format (HTML-only, with JS, framework-specific)
   - Whether docs include live demos or static examples
   - Tone (formal API reference vs conversational guide)

### Phase 3: Generate Documentation

#### Element Overview
- Purpose statement from the manifest description
- When to use this element (and when not to)
- Quick start example showing the simplest useful usage

#### Usage Examples

Generate HTML examples using `generate_html` for:
- **Basic usage**: Minimal attributes, default slot content
- **Common configurations**: Each major attribute combination
- **Slot variations**: Each named slot with appropriate content
- **Composition**: Element used with related elements from the same library

#### Attributes Reference

| Name | Type | Default | Description |
|------|------|---------|-------------|

For enum attributes, include a sub-table or list of valid values with descriptions.

#### Slots Reference

| Name | Description | Content Guidelines |
|------|-------------|--------------------|

Include example HTML showing content in each slot.

#### Events Reference

| Name | Detail | Description |
|------|--------|-------------|

Include JavaScript examples showing how to listen for and handle each event.

#### Styling

**CSS Custom Properties**:

| Property | Default | Description |
|----------|---------|-------------|

**CSS Parts**:

| Part | Description |
|------|-------------|

**CSS States**:

| State | Description |
|-------|-------------|

Include a CSS example showing common customizations.

#### Accessibility

Document any manifest-sourced accessibility patterns:
- ARIA role and keyboard interaction
- Required content for screen readers
- Focus management behavior

### Phase 4: Output Format

Adapt the output format to match what the user needs:

- **Markdown**: Default, suitable for docs sites or READMEs
- **HTML**: If the project uses HTML documentation
- **Storybook**: If stories exist, generate a `*.stories.ts` with knobs for each attribute

Ask the user which format they prefer if it's not clear from context.

## Guidelines

- **Use manifest data as the source of truth**: Don't invent behavior not documented in the manifest
- **Generate working examples**: Call `generate_html` to produce valid HTML, don't hand-write examples that might have errors
- **Match existing tone**: If the project's docs are terse, be terse. If they're tutorial-style, match that
- **Link to related elements**: When composition makes sense, reference other elements from `cem://elements`
- **Mark uncertainty**: If the manifest lacks enough detail for a section, note it rather than guessing
- **Docs feed the manifest**: When writing JSDoc, remember it generates the manifest. Use `@summary` for the short description, `@cssprop` / `@csspart` / `@slot` / `@fires` for API docs. Well-written JSDoc produces a better manifest, which produces better MCP and LSP support.
- **Write MCP-effective descriptions**: Descriptions should be concise, specific, and machine-readable. Avoid vague language. State what the attribute *does*, not just what it *is*. See the CEM guide on effective MCP descriptions.
- **Demo files are docs too**: Each demo should show one variant or feature. Use `<meta itemprop="name">` and `<meta itemprop="description">` for demo metadata. Keep demos self-contained with inline imports.
- **Escape backticks in template HTML comments**: HTML comments inside Lit `html` tagged template literals are used for slot/part docs. Backticks in these comments (e.g. `` `currentColor` ``) break cem's TS transform because they're parsed as template literal boundaries. Escape them with `\`` or avoid them.
