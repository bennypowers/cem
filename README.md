# cem

**cem** is a command-line tool for generating [Custom Elements Manifest][cem]
(CEM) files from TypeScript sources. It analyzes your codebase and generates
rich metadata for your custom elements, facilitating documentation, tooling, and
integration.

> [!NOTE]
> `cem` currently supports LitElements written in idiomatic style with
> TypeScript decorators. If you need something more specific,
> [open an issue][issuenew].

## Features

### `cem generate`

- **Generates CEM files** from source code using syntax analysis powered by
[go][go] and [tree-sitter][treesitter].
- Identifies custom elements, classes, variables, functions, and exports.
- Supports elements written in idiomatic Lit typescript style, with a
`@customElement` decorator, and `@property` decorators on class fields.

#### JSDoc
Use JSDoc comments to add metadata to your element classes, similar to other
tools.

- `@attr` / `@attribute` — Custom element attributes
- `@csspart` — CSS shadow parts
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@deprecated` — Marks a feature or member as deprecated
- `@event` — Custom events dispatched by the element
- `@slot` — Named or default slots
- `@summary` — Short summary for documentation

#### HTML Template Analysis for Slots and Parts

- **Automatically detects `<slot>` elements and `part` attributes in your element’s
`render()` template.**
- Merges slot and part information found in templates with any provided
via JSDoc, ensuring comprehensive documentation in the generated manifest.
- **Deprecation and other metadata** for slots and parts can be specified via
YAML in HTML comments.
- **Supports documenting slots and parts inline in your template HTML** using
HTML comments with YAML blocks.
- YAML comments are not necessary to detect slots and parts, but help in
documenting them for your users.

##### Examples
```html
<!--
  summary: The main slot for content
  description: |
    This slot displays user-provided content.
    Supports multiline **markdown**.
  deprecated: true
-->
<slot></slot>
<!-- slot:
       summary: Named slot summary
     part:
       summary: Part summary
-->
<slot name="info" part="info-part"></slot>
```

#### CSS Custom Properties
Supports CSS Custom Properties by scanning css files and css tagged-template-literals

- Custom properties beginning with `_` will be ignored (treated as "private")
e.g. `var(--_private)`
- If you provide a [Design Tokens Community Group][dtcg] format module (JSON) to
`cem` via the `--design-tokens` flag,
`cem` will add metadata from your design system to any matching css variables it
finds in your elements
- You can use jsdoc-like comment syntax before each var call to document your
variables

##### Example

```css
:host {
  color:
    /**
     * custom color for use in this element
     * @summary color
     * @deprecated just use the `color` property
     */
    var(--custom-color);
  border:
    1px
    solid
    /** Border color of the element */
    var(--border-color);
}
```

## Installation

For go binaries:
```sh
go install bennypowers.dev/cem@latest
```

For NPM projects:

```sh
npm install --save-dev @pwrs/cem
```

Or clone this repository and build from source:

```sh
git clone https://github.com/bennypowers/cem.git
cd cem
make
```

## Usage

Generate a custom elements manifest from your files:

```sh
cem generate \
  "src/**/*.ts" \
  --design-tokens npm:@my-ds/tokens/tokens.json \
  --exclude "src/**/*.test.ts" \
  --output custom-elements.json
```

For npm projects you can use `npx @pwrs/cem generate ...`.
### Arguments
| Argument                     | Type               | Description                                                                                       |
| ---------------------------- | ------------------ | ------------------------------------------------------------------------------------------------- |
| `<files or globs>`           | positional (array) | Files or glob patterns to include                                                                 |
| `--exclude, -e`              | array              | Files or glob patterns to exclude                                                                 |
| `--design-tokens, -t`        | string             | Path or npm specifier for DTCG-format design tokens                                               |
| `--design-tokens-prefix, -p` | string             | CSS custom property prefix for design tokens                                                      |
| `--output, -o`               | string             | Write the manifest to this file instead of stdout                                                 |
| `--no-default-excludes`      | bool               | Do not exclude files by default (e.g., `.d.ts` files will be included unless excluded explicitly) |

By default, some files (like `.d.ts` TypeScript declaration files) are excluded from the manifest.
Use `--no-default-excludes` if you want to include all matching files and manage excludes yourself.

## Contributing

For information on building and testing, please see
[CONTRIBUTING.md][contributingmd].

## License

This program is free software: you can redistribute it and/or modify it under
the terms of the [GNU General Public License v3.0][gpl3].

&copy; 2025 Benny Powers <web@bennypowers.com>

[cem]: https://github.com/webcomponents/custom-elements-manifest
[dtcg]: https://tr.designtokens.org/format/
[go]: https://go.dev
[treesitter]: https://tree-sitter.github.io/tree-sitter/
[gpl3]: https://www.gnu.org/licenses/gpl-3.0.html
[contributingmd]: ./CONTRIBUTING.md
[issuenew]: https://github.com/bennypowers/cem/issues/new
