---
title: Generate
layout: page
---

Generates custom elements manifest from source code using [tree-sitter][treesitter] syntax analysis. Identifies custom elements, properties, attributes, slots, events, CSS parts, and CSS custom properties. See [Documenting Components][documenting] for usage guide.

## Syntax

```bash
cem generate [files...] [flags]
```

Generate from specific files or globs:

```bash
cem generate "src/**/*.ts" --exclude "src/**/*.test.ts"
```

Generate using configuration file:

```bash
cem generate
```

{{<tip "warning">}}
Best supports LitElements written in idiomatic style with TypeScript decorators. Rudimentary support for `extends HTMLElement`. See [issue tracker][issues] for feature requests.
{{</tip>}}

## Arguments

| Argument                        | Type               | Description                                                                                       |
| ------------------------------- | ------------------ | ------------------------------------------------------------------------------------------------- |
| `<files or globs>`              | positional (array) | Files or glob patterns to include                                                                 |
| `--package, -p`                 | string             | Path to a package directory                                                                       |
| `--output, -o`                  | string             | Write the manifest to this file instead of stdout                                                 |
| `--watch, -w`                   | bool               | Watch files for changes and regenerate automatically                                              |
| `--exclude, -e`                 | array              | Files or glob patterns to exclude                                                                 |
| `--no-default-excludes`         | bool               | Do not exclude `.d.ts` files by default                                                           |
| `--design-tokens`               | string             | Path or npm specifier for DTCG-format design tokens                                               |
| `--design-tokens-prefix`        | string             | CSS custom property prefix for design tokens                                                      |
| `--demo-discovery-file-glob`    | string             | Glob pattern for discovering demo files                                                           |
| `--demo-discovery-url-pattern`  | string             | URLPattern with named parameters (`:param`) for matching demo file paths                          |
| `--demo-discovery-url-template` | string             | Go template with functions for generating canonical demo URLs                                     |
| `--source-control-root-url`     | string             | Canonical public source control URL for repository root                                           |
| `--project-dir`                 | string             | **Deprecated:** Use `--package` instead                                                           |

By default, `.d.ts` TypeScript declaration files are excluded. Use `--no-default-excludes` to include all matching files.

## JSDoc Tags

Document your custom elements using these JSDoc tags. See [Documenting Components][documenting] for examples.

### Element and Member Tags

- `@attr` / `@attribute` — Custom element attributes
- `@deprecated` — Marks a feature or member as deprecated
- `@event` — Custom events dispatched by the element
- `@example` — Code examples with optional captions
- `@summary` — Short summary for documentation

### CSS and Shadow DOM Tags

- `@csspart` — CSS shadow parts
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@slot` — Named or default slots

### Demo and Documentation Tags

- `@demo` — Demo URL with optional description

### Usage Examples

```typescript
/**
 * A button component with variants
 *
 * @summary Interactive button element
 * @slot - Button content and icon
 * @csspart button - The button element
 * @cssprop --button-bg - Background color
 * @fires click - Dispatched when button is clicked
 * @example
 * ```html
 * <my-button variant="primary">Click me</my-button>
 * ```
 */
@customElement('my-button')
class MyButton extends LitElement { }
```

See the [generate test fixtures][fixtures] for comprehensive examples.

## Configuration

Configure generation in `.config/cem.yaml`:

```yaml
sourceControlRootUrl: "https://github.com/your/repo/tree/main/"
generate:
  files:
    - "src/**/*.ts"
  exclude:
    - "src/**/*.test.ts"
  output: custom-elements.json
  designTokens: "npm:@my-ds/tokens/tokens.json"
  designTokensPrefix: "--my-ds"
  demoDiscovery:
    fileGlob: "src/**/demo/*.html"
    urlPattern: "/src/:tag/demo/:demo.html"
    urlTemplate: "https://example.com/{{.tag}}/{{.demo}}/"
```

## Output Format

Generates JSON conforming to the [Custom Elements Manifest][cem] schema. HTML-sensitive characters are escaped using standard JSON unicode sequences (e.g., `<` becomes `\u003c`) for security.

## See Also

- **[Documenting Components][documenting]** - JSDoc usage guide and examples
- **[Development Workflow][workflow]** - When to regenerate the manifest
- **[Configuration Reference][config]** - All config options

[treesitter]: https://tree-sitter.github.io/tree-sitter/
[issues]: https://github.com/bennypowers/cem/issues/new
[documenting]: /docs/usage/documenting-components/
[fixtures]: https://github.com/bennypowers/cem/tree/main/generate/testdata/fixtures/
[cem]: https://github.com/webcomponents/custom-elements-manifest
[workflow]: /docs/usage/workflow/
[config]: /docs/reference/configuration/
