---
title: search
summary: Search through custom elements manifests for any element by keyword or regex pattern
---

# `cem search`

Search through the custom elements manifest for any element matching the given pattern.

## Usage

```bash
cem search [pattern] [flags]
```

## Description

The `search` command allows you to find any element in your custom elements manifest by searching through names, descriptions, summaries, and labels. This includes:

- Custom element tags
- Modules and files  
- Attributes
- Slots
- CSS properties, parts, and states
- Events
- Methods and fields
- Functions and variables
- Demos

The search pattern is treated as a **regular expression by default**, allowing for powerful and flexible searches. If the regex is invalid, it automatically falls back to literal string matching.

All searches are **case-insensitive**.

## Options

- `-f, --format string` - Output format: `table` (default) or `tree`

## Examples

### Basic Search

Search for anything containing "button":
```bash
cem search button
```

### Regular Expression Search

Find elements starting with "my-" and ending with "button":
```bash
cem search "^my-.*button$"
```

Find elements containing "click" OR "hover":
```bash
cem search "click|hover"
```

### Format Options

Display results as a tree:
```bash
cem search button --format tree
```

Display results as a table (default):
```bash
cem search button --format table
```

### Specific Content Types

Find CSS-related properties:
```bash
cem search "css.*property"
```

Find header-related slots:
```bash
cem search "slot.*header"
```

Find deprecated items:
```bash
cem search deprecated --format tree
```

## Output Formats

### Table Format (Default)

The table format shows detailed information organized by sections:

```
# <root>

# module rh-button/rh-button.js

# `<rh-button>`

Triggers actions on the page or in the background

## Demos

URL                                            | Source
-----------------------------------------------|--------
`https://ux.redhat.com/elements/button/demo/` | https://...

## CSS Parts

Name
--------
`button`
```

### Tree Format

The tree format shows a hierarchical view of matching elements:

```
Search Results for: button
<root>
├─┬ module rh-button/rh-button.js
│ └─┬ <rh-button>
│   └─┬ Parts
│     └── button
```

## Regular Expression Features

Since patterns are treated as regular expressions, you can use:

- `^` - Start of string anchor
- `$` - End of string anchor  
- `.*` - Match any characters
- `|` - OR operator
- `[abc]` - Character class
- `\d`, `\w`, `\s` - Character shortcuts

If your regex is invalid, the search automatically falls back to literal string matching.

## Global Options

The search command also supports all global `cem` options:

- `-p, --package string` - Package specifier or path
- `--config string` - Config file path
- `-v, --verbose` - Verbose output

## See Also

- [`cem list`](./list.md) - List specific types of elements from the manifest
- [`cem validate`](./validate.md) - Validate manifest files
- [`cem generate`](./generate.md) - Generate manifest files