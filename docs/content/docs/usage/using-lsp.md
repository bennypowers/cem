---
title: Using LSP Features
weight: 90
---

Once you've set up the [LSP integration](/docs/installation/lsp/), you can leverage powerful editor features for working with custom elements.

## Prerequisites

- CEM LSP configured in your editor (see [LSP Integration](/docs/installation/lsp/))
- A `custom-elements.json` manifest in your project (`cem generate`)

## Autocomplete

The LSP provides intelligent autocomplete for custom elements in HTML and TypeScript files.

### Tag Name Completion

Start typing a tag name and see suggestions for your custom elements:

```html
<my-bu|
```

Press Ctrl+Space (or your editor's autocomplete trigger) to see:
- `my-button` - A clickable button component
- `my-button-group` - Container for related buttons

### Attribute Completion

Type a space after a tag name to see available attributes:

```html
<my-button |
```

Suggestions include:
- `variant` - Button style variant
- `size` - Button size
- `disabled` - Disable the button
- `loading` - Show loading state

### Attribute Value Completion

For attributes with known values (like enums), get value suggestions:

```html
<my-button variant="|"
```

Suggestions show:
- `primary` - Primary action button
- `secondary` - Secondary action button
- `danger` - Destructive action button

### Slot Attribute Completion

When elements have named slots, autocomplete suggests valid slot names:

```html
<my-card>
  <div slot="|">
```

Suggestions:
- `header` - Card header content
- `footer` - Card footer content
- `media` - Card media content

### In Lit Templates

The LSP works in Lit `html` template literals:

```typescript
render() {
  return html`
    <my-button variant="primary">
      <span slot="start">📝</span>
      Click Me
    </my-button>
  `;
}
```

Get autocomplete for:
- Tag names
- Attributes
- Slot names
- Event names (with `@` prefix)
- Properties (with `.` prefix)
- Boolean attributes (with `?` prefix)

## Hover Documentation

Hover over custom elements and attributes to see inline documentation.

### Element Hover

Hover over a tag name to see:

```html
<my-button>
     ↑ hover here
```

**Documentation shows**:
- Element summary and description
- Complete API (properties, attributes, slots, events)
- CSS custom properties and parts
- Links to source code

### Attribute Hover

Hover over an attribute to see:

```html
<my-button variant="primary">
           ↑ hover here
```

**Documentation shows**:
- Attribute description
- Type information
- Default value
- Valid values (for enums)
- Deprecation warnings

### CSS Part Hover

Hover over `::part()` selectors to see styling guidance:

```css
my-button::part(label) {
                ↑ hover here
  font-weight: bold;
}
```

## Go-to-Definition

Jump to the source code of custom elements.

### From Tag Names

1. **Position cursor** on a tag name: `<my-button>`
2. **Trigger go-to-definition** (F12 in VS Code, `gd` in Neovim)
3. **Jump to source** - Opens the component source file

### From Attributes

1. **Position cursor** on an attribute: `variant="primary"`
2. **Trigger go-to-definition**
3. **Jump to property definition** in the component class

Works for:
- Properties
- Attributes
- Reflected properties
- Custom events (future)
- Slots (future)

## Find References

Find all usages of a custom element across your workspace.

### How to Use

1. **Position cursor** anywhere on a custom element tag
2. **Trigger find-references** (Shift+F12 in VS Code, `gr` in Neovim)
3. **See all usages** across HTML, TypeScript, and JavaScript files

### What You Get

- Workspace-wide search
- Filtered by `.gitignore` (excludes `node_modules/`, etc.)
- Shows only start tags (avoids duplicates)
- Works in template literals

**Example results**:
```
elements/my-app/demo/index.html (3 matches)
elements/my-dashboard/my-dashboard.ts (1 match)
docs/examples/buttons.html (2 matches)
```

## Workspace Symbols

Search for custom elements across your entire workspace.

### How to Use

1. **Open symbol search** (Ctrl+T in VS Code, `:Telescope lsp_workspace_symbols` in Neovim)
2. **Type element name** (fuzzy matching)
3. **Jump to definition**

**Example**:
- Search: `btn`
- Finds: `my-button`, `icon-button`, `button-group`

## Error Detection & Quick Fixes

The LSP validates your HTML and provides one-click fixes.

### Invalid Slot Names

```html
<my-card>
  <div slot="heade">Title</div>
  <!-- ❌ Red squiggles appear -->
</my-card>
```

**Quick fix available**:
1. Position cursor on the error
2. Trigger quick fix (Ctrl+. in VS Code, `<leader>ca` in Neovim)
3. Select "Change 'heade' to 'header'"
4. Fixed automatically!

### Typos in Tag Names

```html
<my-buttom>Click</my-buttom>
<!-- ❌ Typo detected -->
```

**Quick fixes**:
- "Change 'my-buttom' to 'my-button'"
- "Add import for my-button"

### Invalid Attribute Names

```html
<my-button varient="primary">
<!-- ❌ Attribute typo -->
```

**Quick fix**:
- "Change 'varient' to 'variant'"

### Invalid Attribute Values

```html
<my-button variant="primar">
<!-- ❌ Invalid value (should be "primary") -->
```

**Quick fix**:
- "Change 'primar' to 'primary'"

### Missing Imports

```html
<my-button>Click</my-button>
<!-- ❌ Element exists but not imported -->
```

**Quick fix**:
- "Add import for my-button from '@my-lib/button'"

## Tips & Tricks

### Enable Verbose Logging

See what the LSP is doing:

**VS Code**: Set `"cem.lsp.trace.server": "verbose"`
**Neovim**: Set `trace = 'verbose'` in LSP config

Logs show:
- Manifest loading
- File parsing
- Completion requests
- Validation results

### Regenerate Manifest for Fresh Data

The LSP watches your manifest file, but you may need to regenerate:

```sh
cem generate
```

The LSP picks up changes automatically.

### Use in Monorepos

The LSP discovers manifests from:
- Local `custom-elements.json`
- Dependencies in `package.json`
- Workspace packages

No special configuration needed.

### Combine with Dev Server

Run both together for the best experience:

**Terminal 1**:
```sh
cem serve
```

**Terminal 2** (optional, for manifest regeneration):
```sh
cem generate --watch
```

Now you have:
- Live preview in browser
- Autocomplete and validation in editor
- Auto-regenerating manifest

## Troubleshooting

### No Autocomplete

1. **Check manifest exists**: `ls custom-elements.json`
2. **Verify LSP is running**: Check editor's LSP status
3. **Regenerate manifest**: `cem generate`
4. **Restart editor**: Reload window

### Outdated Suggestions

1. **Regenerate manifest**: `cem generate`
2. **Check file watcher**: Ensure manifest changes are detected
3. **Restart LSP**: Reload editor or restart LSP client

### Validation Errors Don't Appear

1. **Check diagnostics are enabled** in your editor
2. **Verify manifest has element schemas**
3. **Enable verbose logging** to see validation results

### Performance Issues

For large projects:
- Limit workspace scope
- Exclude build directories in `.gitignore`
- Disable features you don't use

## See Also

- **[LSP Integration](/docs/installation/lsp/)** - Setup instructions
- **[LSP Protocol Reference](/docs/reference/lsp/)** - Technical details
- **[Development Workflow](/docs/usage/workflow/)** - How LSP fits into the dev cycle
- **[Editor Configuration](/docs/installation/editors/)** - Editor-specific tips
