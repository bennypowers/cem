# cem

**cem** is a command-line tool for generating [Custom Elements Manifest][cem]
(CEM) files from TypeScript sources. It analyzes your codebase and generates
rich metadata for your custom elements, facilitating documentation, tooling, and
integration.

> [!NOTE]
> `cem` best supports LitElements written in idiomatic style with
> TypeScript decorators. There is rudimentary support for `extends HTMLElement`,
> but it is not a high priority for development. If you need something more
> specific [open an issue][issuenew].

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

## Shell Completion

`cem` supports shell completion for Bash, Zsh, Fish, and PowerShell. The completion scripts provide tab completion for commands, flags, and file paths.

### Setup Instructions

#### Bash

To load completions in your current shell session:

```bash
source <(cem completion bash)
```

To load completions for every new session, execute once:

**Linux:**
```bash
cem completion bash > /etc/bash_completion.d/cem
```

**macOS:**
```bash
cem completion bash > $(brew --prefix)/etc/bash_completion.d/cem
```

#### Zsh

To load completions in your current shell session:

```zsh
source <(cem completion zsh)
```

To load completions for every new session, execute once:

```zsh
cem completion zsh > "${fpath[1]}/_cem"
```

You will need to start a new shell for this setup to take effect.

#### Fish

To load completions in your current shell session:

```fish
cem completion fish | source
```

To load completions for every new session, execute once:

```fish
cem completion fish > ~/.config/fish/completions/cem.fish
```

#### PowerShell

To load completions in your current shell session:

```powershell
cem completion powershell | Out-String | Invoke-Expression
```

To load completions for every new session, add the output of the above command to your PowerShell profile.

## Features

### `cem generate`

- **Generates CEM files** from source code using syntax analysis powered by
[go][go] and [tree-sitter][treesitter].
- Identifies custom elements, classes, variables, functions, and exports.
- Supports elements written in idiomatic Lit typescript style, with a
`@customElement` decorator, and `@property` decorators on class fields.

#### JSDoc
Use JSDoc comments to add metadata to your element classes, similar to other
tools. Add a description by separating the name of the item with ` - `

- `@attr` / `@attribute` — Custom element attributes
- `@csspart` — CSS shadow parts
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@demo` — Demo URL
- `@deprecated` — Marks a feature or member as deprecated
- `@event` — Custom events dispatched by the element
- `@slot` — Named or default slots
- `@summary` — Short summary for documentation

See the [test-fixtures](tree/main/generate/test-fixtures/) directory for examples

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

---

#### Element Demos

`cem generate` supports documenting your elements' demos by linking directly
from JSDoc, or by configurable file-system based discovery.

##### 1. JSDoc `@demo` Tag

Add demos directly to your element class or members with the `@demo` tag:

```ts
/**
 * @demo https://example.com/my-element-plain/
 * @demo https://example.com/my-element-fancy/ - A fancier demo with description
 */
@customElement('my-element')
class MyElement extends LitElement {
  // ...
}
```

Demos defined this way will always appear in your manifest for the element.

##### 2. Demo Discovery

`cem` can automatically discover demos from your codebase based on your
repository structure and configuration.

#### Demo Discovery Options

Configure demo discovery with the `demoDiscovery` key in your `.config/cem.yaml` file

```yaml
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"
generate:
  demoDiscovery:
    fileGlob: "demos/**/*.html"
    urlPattern: "demos/(?P<tag>[\w-]+)/(?P<demo>[\w-]+).html"
    urlTemplate: "https://example.com/elements/{tag}/{demo}/"
```

**Demo discovery options:**

| Option                 | Type   | Description                                                                                  |
| ---------------------- | ------ | -------------------------------------------------------------------------------------------- |
| `fileGlob`             | string | Glob pattern for discovering demo files.                                                     |
| `sourceControlRootUrl` | string | Canonical public source control URL for your repository root (on the main branch).           |
| `urlPattern`           | string | Pattern for generating demo URLs, e.g. `"demos/{tag}.html"`. `{tag}` is replaced by tag name.|
| `urlTemplate`          | string | (optional) Alternative URL template for demo links.                                          |

---

## Configuration Reference

You can configure CEM via `.config/cem.yaml`, relative to your project root,
or via CLI flags.

### Example Configuration

```yaml
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"
generate:
  files:
    - "src/**/*.ts"
  exclude:
    - "src/**/*.test.ts"
  output: "custom-elements.json"
  noDefaultExcludes: false
  designTokens:
    spec: "npm:@my-ds/tokens/tokens.json"
    prefix: "--my-ds"
  demoDiscovery:
    fileGlob: "demos/**/*.html"
    urlPattern: "demos/(?P<tag>[\w-]+)/(?P<demo>[\w-]+).html"
    urlTemplate: "https://example.com/elements/{tag}/{demo}/"
```

---

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

### Command Line Arguments

| Argument                        | Type               | Description                                                                                       |
| ------------------------------- | ------------------ | ------------------------------------------------------------------------------------------------- |
| `<files or globs>`              | positional (array) | Files or glob patterns to include                                                                 |
| `--output, -o`                  | string             | Write the manifest to this file instead of stdout                                                 |
| `--exclude, -e`                 | array              | Files or glob patterns to exclude                                                                 |
| `--no-default-excludes`         | bool               | Do not exclude files by default (e.g., `.d.ts` files will be included unless excluded explicitly) |
| `--design-tokens, -t`           | string             | Path or npm specifier for DTCG-format design tokens                                               |
| `--design-tokens-prefix, -p`    | string             | CSS custom property prefix for design tokens                                                      |
| `--demo-discovery-file-glob`    | string             | Glob pattern for discovering demo files                                                           |
| `--demo-discovery-url-pattern`  | string             | Go Regexp pattern with named capture groups for generating canonical demo urls                    |
| `--demo-discovery-url-template` | string             | URL pattern string using {groupName} syntax to interpolate named captures from the URL pattern    |
| `--source-control-root-url`     | string             | Glob pattern for discovering demo files                                                           |
| `--project-dir`                 | string             | Specify the project root directory to use for resolving relative paths and configuration.         |


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
