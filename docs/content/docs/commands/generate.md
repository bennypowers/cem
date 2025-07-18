---
title: Generate
layout: page
---

- **Generates CEM files** from source code using syntax analysis powered by
[go][go] and [tree-sitter][treesitter].
- Identifies custom elements, classes, variables, functions, and exports.
- Supports elements written in idiomatic style using Lit and TypeScript, with a
`@customElement` decorator, and `@property` decorators on class fields.

{{<tip "warning">}}

`cem generate` best supports LitElements written in idiomatic style with
TypeScript decorators. There is rudimentary support for `extends HTMLElement`,
but it is not a high priority for development. If you need something more
specific [open an issue](https://github.com/bennypowers/cem/issues/new).

{{</tip>}}

## JSDoc
Use JSDoc comments to add metadata to your element classes, similar to other
tools. Add a description by separating the name of the item with ` - `

- `@attr` / `@attribute` — Custom element attributes
- `@csspart` — CSS shadow parts. See [Slots and Parts](#slots-and-parts)
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@demo` — Demo URL
- `@deprecated` — Marks a feature or member as deprecated
- `@event` — Custom events dispatched by the element
- `@slot` — Named or default slots. See [Slots and Parts](#slots-and-parts)
- `@summary` — Short summary for documentation

See the [generate test fixtures](https://github.com/bennypowers/cem/tree/main/generate/test/fixtures/) directory for examples

<a id="html-template-analysis-for-slots-and-parts"></a>
## Slots and Parts

- **Automatically detects `<slot>` elements and `part` attributes in your element’s
`render()` template.**
- Merges slot and part information found in templates with any provided
via JSDoc, ensuring comprehensive documentation in the generated manifest.
- **Supports documenting slots and parts inline in your template HTML** using
HTML comments.
- If the comment is a plain string, it will be used as the `description` for the element. In cases where an element is both a slot and a part, the description will only be applied to the slot.
- For more detailed documentation, you can use YAML inside the comment.
- Comments are not necessary to detect slots and parts, but help in
documenting them for your users.

### Examples

#### Plain String Comment
A simple comment will be treated as the description.
```html
<!-- This is the default slot. -->
<slot></slot>
```

#### YAML Comment
For more complex metadata, use YAML syntax.
```html
<!--
  summary: The main slot for content
  description: |
    This slot displays user-provided content.
    Supports multiline **markdown**.
  deprecated: true
-->
<slot></slot>
```

#### Documenting Slots and Parts together
When an element is both a slot and a part, you can document both in a single comment.
```html
<!-- slot:
       summary: Named slot summary
     part:
       summary: Part summary
-->
<slot name="info" part="info-part"></slot>
```

## CSS Custom Properties
Supports CSS Custom Properties by scanning css files and css tagged-template-literals

- Custom properties beginning with `_` will be ignored (treated as "private")
e.g. `var(--_private)`
- If you provide a [Design Tokens Community Group][dtcg] format module (JSON) to
`cem` via the `--design-tokens` flag,
`cem` will add metadata from your design system to any matching css variables it
finds in your elements
- You can use jsdoc-like comment syntax before each var call to document your
variables

### Example

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

<a id="element-demos"></a>
## Demos

`cem generate` supports documenting your elements' demos by linking directly
from JSDoc, or by configurable file-system based discovery.

### JSDoc `@demo` Tag

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

When using `@alias` tags, the alias will be [slugified](https://www.npmjs.com/package/slug)
for use in the URL.

### Automatic Demo Discovery

`cem` can automatically discover demos from your codebase based on your
repository structure and configuration. Demos that are co-located with their
component's source module will be prioritized in the generated manifest.

<a id="demo-discovery-options"></a>
## Demo Discovery

The `urlPattern` is a flexible Go regular expression with named capture groups.
You can use it to match complex file paths and extract the parts you need to
build your demo URLs.

For example, if your demos are in subdirectories like `src/my-element/demos/foo.html`,
you could use a pattern like this:

```yaml
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"
generate:
  demoDiscovery:
    fileGlob: "src/**/demos/*.html"
    urlPattern: "src/(?P<tag>[\w-]+)/demos/(?P<demo>[\w-]+).html"
    urlTemplate: "https://example.com/elements/{tag}/{demo}/"
```

**Demo discovery options:**

| Option                 | Type   | Description                                                                                  |
| ---------------------- | ------ | -------------------------------------------------------------------------------------------- |
| `fileGlob`             | string | Glob pattern for discovering demo files.                                                     |
| `sourceControlRootUrl` | string | Canonical public source control URL for your repository root (on the main branch).           |
| `urlPattern`           | string | Go Regexp pattern with named capture groups for generating canonical demo urls.              |
| `urlTemplate`          | string | (optional) Alternative URL template for demo links.                                          |

## Monorepos

If you are planning to use `cem` in an npm or yarn monorepo, the best way for now
is to create a new `.config/cem.yaml` file for each package you want to generate 
for, instead of using a top-level config file.

### Example

Root package.json:
```json
{
  "scripts": {
    "generate": "npm run generate --workspaces"
  },
  "workspaces": [
    "./core",
    "./elements"
  ]
}
```

core/.config/cem.yaml:
```yaml
generate:
  files:
    - './**/*.ts'
```

core/package.json
```json
{
  "scripts": {
    "generate": "cem generate"
  }
}
```

elements/.config/cem.yaml:
```yaml
generate:
  files:
    - './**/*.ts'
```

elements/package.json
```json
{
  "scripts": {
    "generate": "cem generate"
  }
}
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

{{<tip "note">}}
The `generate` command does not support remote packages. To inspect a remote package's manifest, use the `cem list` command.
{{</tip>}}

<a id="command-line-arguments"></a>
### Arguments

| Argument                        | Type               | Description                                                                                       |
| ------------------------------- | ------------------ | ------------------------------------------------------------------------------------------------- |
| `<files or globs>`              | positional (array) | Files or glob patterns to include                                                                 |
| `--package, -p`                 | string             | Path to a package directory.                                                                      |
| `--output, -o`                  | string             | Write the manifest to this file instead of stdout                                                 |
| `--exclude, -e`                 | array              | Files or glob patterns to exclude                                                                 |
| `--no-default-excludes`         | bool               | Do not exclude files by default (e.g., `.d.ts` files will be included unless excluded explicitly) |
| `--design-tokens`               | string             | Path or npm specifier for DTCG-format design tokens                                               |
| `--design-tokens-prefix`        | string             | CSS custom property prefix for design tokens                                                      |
| `--demo-discovery-file-glob`    | string             | Glob pattern for discovering demo files                                                           |
| `--demo-discovery-url-pattern`  | string             | Go Regexp pattern with named capture groups for generating canonical demo urls                    |
| `--demo-discovery-url-template` | string             | URL pattern string using {groupName} syntax to interpolate named captures from the URL pattern    |
| `--source-control-root-url`     | string             | Glob pattern for discovering demo files                                                           |
| `--project-dir`                 | string             | **Deprecated:** Use `--package` instead.                                                          |

By default, some files (like `.d.ts` TypeScript declaration files) are excluded from the manifest.
Use `--no-default-excludes` if you want to include all matching files and manage excludes yourself.

[cem]: https://github.com/webcomponents/custom-elements-manifest
[dtcg]: https://tr.designtokens.org/format/
[go]: https://go.dev
[treesitter]: https://tree-sitter.github.io/tree-sitter/
[gpl3]: https://www.gnu.org/licenses/gpl-3.0.html
[contributingmd]: ./CONTRIBUTING.md
[issuenew]: https://github.com/bennypowers/cem/issues/new
[installationdocs]: https://bennypowers.github.io/cem/installation/
[generatedocs]: https://bennypowers.github.io/cem/commands/generate/
