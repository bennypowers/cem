---
name: migrate
description: >
  Identify breaking changes between custom element manifest versions and suggest
  migration steps. Use when the user asks to "migrate", "upgrade components",
  "find breaking changes", "compare manifests", "what changed between versions",
  or mentions "deprecation" or "migration guide".
tools: Read, Glob, Grep, Bash
---

# Migration Assistant

Compare custom element manifests across versions to identify breaking changes and generate migration guidance.

## Workflow

### Phase 1: Identify Versions to Compare

Determine the two versions to compare. Sources:

1. **Git-based**: Compare the current manifest against a previous git ref
   ```bash
   git show {ref}:path/to/custom-elements.json
   ```

2. **File-based**: User provides two manifest files or paths

3. **Package-based**: Compare installed version against a target version
   - Read current manifest from `node_modules/{package}/custom-elements.json`
   - Check the target version's manifest

Read the current project's elements for context:

```
cem://elements
```

### Phase 2: Diff Analysis

Compare the two manifests and categorize changes:

#### Removed (Breaking)
- **Removed elements**: Tag names that no longer exist
- **Removed attributes**: Attributes dropped from an element
- **Removed slots**: Named slots that no longer exist
- **Removed events**: Events that are no longer fired
- **Removed CSS parts**: Parts no longer exposed
- **Removed CSS custom properties**: Theming properties removed

#### Changed (Potentially Breaking)
- **Renamed attributes**: Attribute name changed (detect via similar names + same type)
- **Type changes**: Attribute type changed (e.g., `string` to `enum`, `boolean` to `string`)
- **Enum value changes**: Valid values added or removed
- **Default value changes**: Different default behavior
- **Slot renamed**: Named slot changed (detect via similar names)
- **Event detail changes**: Event detail type modified
- **CSS property renamed**: Custom property name changed

#### Added (Non-Breaking)
- **New elements**: Newly available tag names
- **New attributes**: Additional configuration options
- **New slots**: Additional content areas
- **New events**: Additional observable state changes
- **New CSS parts/properties/states**: Additional styling hooks

#### Deprecated
- **Deprecated attributes**: Marked for future removal
- **Deprecated elements**: Elements scheduled for removal
- Look for `deprecated` field or deprecation notes in descriptions

### Phase 3: Impact Assessment

For each breaking change, search the project for affected usage:

- Search HTML files for removed/changed element tag names
- Search for removed attribute names on affected elements
- Search for removed event listeners
- Search CSS for removed parts and custom properties

Report the number of affected files and locations.

### Phase 4: Migration Report

````markdown
## Migration Report: {package} {old-version} -> {new-version}

### Summary
- Breaking changes: X
- Potentially breaking: X
- New features: X
- Deprecations: X
- Files affected: X

### Breaking Changes

#### 1. `<old-element>` removed
**Impact**: Used in X files
**Migration**: Replace with `<new-element>`
```diff
- <old-element attr="value">content</old-element>
+ <new-element attr="value">content</new-element>
```
**Files affected**:
- `src/pages/login.html:15`
- `src/components/form.html:42`

#### 2. `variant` attribute on `<prefix-button>` renamed to `appearance`
**Impact**: Used in X files
**Migration**: Rename the attribute
```diff
- <prefix-button variant="primary">
+ <prefix-button appearance="primary">
```

### Deprecation Warnings

[Items that still work but will break in a future version]

### New Features

[Brief summary of what's new — the user might want to adopt these]

### Recommended Migration Order

1. [Fix breaking changes that cause build/runtime errors first]
2. [Address deprecation warnings]
3. [Optionally adopt new features]
````

## Guidelines

- **Be conservative with "renamed" detection**: Only suggest a rename if the types match and names are similar — don't guess
- **Show concrete diffs**: Every breaking change should include a before/after code example
- **Find all affected files**: Don't just describe the change abstractly — show where it matters in the project
- **Distinguish runtime breaks from cosmetic changes**: A removed attribute is worse than a changed default
- **Suggest incremental migration**: If there are many changes, recommend an order that keeps the project functional at each step
