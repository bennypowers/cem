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
| `--design-tokens`               | string             | Path, npm, or jsr specifier for DTCG-format design tokens. Package specifiers fall back to esm.sh if not installed locally |
| `--design-tokens-prefix`        | string             | CSS custom property prefix for design tokens                                                      |
| `--demo-discovery-file-glob`    | string             | Glob pattern for discovering demo files                                                           |
| `--demo-discovery-url-pattern`  | string             | URLPattern with named parameters (`:param`) for matching demo file paths                          |
| `--demo-discovery-url-template` | string             | Go template with functions for generating canonical demo URLs                                     |
| `--source-control-root-url`     | string             | Canonical public source control URL for repository root                                           |
| `--project-dir`                 | string             | **Deprecated:** Use `--package` instead                                                           |

By default, `.d.ts` TypeScript declaration files are excluded. Use `--no-default-excludes` to include all matching files.

## JSDoc Tags

Document your custom elements using these JSDoc tags. See [Documenting Components][documenting] for examples.

### Class/Element Level Tags

These tags apply to the custom element class:

- `@alias` — Alternative name for the element
- `@attr` / `@attribute` — Custom element attributes
- `@csspart` — CSS shadow parts
- `@cssprop` / `@cssproperty` — Custom CSS properties
- `@cssstate` — Custom CSS states
- `@customElement` / `@element` / `@tagName` — Tag name (when `@customElement` decorator or `customElements.define` are not in use)
- `@demo` — Demo URL with optional description
- `@deprecated` — Marks element as deprecated
- `@event` / `@fires` — Custom events dispatched by the element
- `@example` — Code examples with optional captions
- `@slot` — Named or default slots
- `@summary` — Short summary for documentation

### Property Level Tags

These tags apply to class properties:

- `@default` / `@defaultvalue` — Default value for the field. Not needed when the field has an initializer. If a field has both, the jsdoc tag wins.
- `@deprecated` — Marks property as deprecated
- `@example` — Code examples for property usage
- `@ignore` / `@internal` — Exclude property from the manifest
- `@summary` — Short summary for the property
- `@type` — Property type annotation

### Method Level Tags

These tags apply to class methods:

- `@deprecated` — Marks method as deprecated
- `@example` — Code examples for method usage
- `@ignore` / `@internal` — Exclude method from the manifest
- `@param` / `@parameter` — Method parameter documentation
- `@return` / `@returns` — Return value documentation
- `@summary` — Short summary for the method

### CSS Property Level Tags

These tags apply to CSS custom properties in stylesheets:

- `@deprecated` — Marks CSS property as deprecated
- `@example` — Code examples for CSS property usage
- `@summary` — Short summary for the CSS property
- `@syntax` — CSS syntax/type definition

### Usage Examples

#### Element Level

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

#### Method Level

```typescript
/**
 * Calculates the sum of two numbers
 *
 * @summary Add two numbers
 * @param {number} a - First number
 * @param {number} b - Second number
 * @returns {number} The sum of a and b
 * @example
 * ```typescript
 * element.add(2, 3); // returns 5
 * ```
 */
add(a: number, b: number): number {
  return a + b;
}
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
  designTokensPrefix: "my-ds"
  demoDiscovery:
    fileGlob: "src/**/demo/*.html"
    urlPattern: "/src/:tag/demo/:demo.html"
    urlTemplate: "https://example.com/{{.tag}}/{{.demo}}/"
```

## Workspace Mode

In a monorepo with `workspaces` in `package.json`, running `cem generate` at the
workspace root generates a manifest for each package that has a `customElements`
field. Each manifest is written to the path specified in that package's
`customElements` field.

```bash
# Generate manifests for all workspace packages
cem generate

# Target a single package
cem generate -p packages/button
```

The `--output` flag is not available in workspace mode. See
[Configuration: Workspace Mode][workspace-mode] for details on config cascade
and package selection.

[workspace-mode]: {{< relref "/docs/reference/configuration#monorepo--workspace-mode" >}}

## Output Format

Generates JSON conforming to the [Custom Elements Manifest][cem] schema. HTML-sensitive characters are escaped using standard JSON unicode sequences (e.g., `<` becomes `\u003c`) for security.

### Function Overloads

When a module declares TypeScript function overloads, only the overload signatures are emitted, one declaration per signature, all sharing the function's name. The trailing implementation signature is not callable by TypeScript consumers, so it is omitted. When the implementation carries JSDoc that a kept overload signature lacks, the implementation's description, summary, and deprecation merge onto the first such signature; other tags (for example `@param` or `@returns`) are not merged. This diverges intentionally from the reference `@custom-elements-manifest/analyzer`, which emits the implementation signature as an additional declaration.

### Generic Type Parameters

The CEM schema cannot represent generic type parameter lists, so `typeParameters` fields are not emitted. Type text is published exactly as written in source, and consumers that need type semantics should run `tsc` against the linked source.

When a parameter, return, or field type references a generic type parameter of its declaration, cem attaches a `source` reference to that type pointing at the parameter's declaration, anchored by line and column (for example `#L1C26`), following GitHub's fragment convention. Given:

```typescript
export function identity<Value extends string = string>(value: Value): Value {
  return value;
}
```

the `value` parameter's type is emitted as:

```json
{
  "text": "Value",
  "source": {
    "href": "https://github.com/your/repo/tree/main/src/identity.ts#L1C26"
  }
}
```

The anchor is the line and column of `Value` in `<Value extends string = string>`. Declaration-level `source` references stay line-only (`#L2`); type-level references add a column (`#L1C26`). Line and column are 1-based, and the column counts Unicode characters. Custom element attributes share their field's type reference. Anchors address the identifier as written when the manifest was generated; if the source file changes, regenerate the manifest. Matching a type text to its declaring parameter is best-effort and lexical: the first parameter in scope whose name appears as a whole word in the type text wins, and a method's own type parameters are considered before its class's, so a method parameter shadows a class parameter even when the names differ. A method's types consider both the method's own and its class's type parameters; field types consider the class's. Source references require `sourceControlRootUrl` to be configured.

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
