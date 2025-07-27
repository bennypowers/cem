---
title: "cem validate"
description: "Validate your custom-elements.json against the official schema"
---

The `cem validate` command validates your `custom-elements.json` file against its corresponding JSON schema and provides intelligent warnings for potentially inappropriate manifest content.

```bash
cem validate [path/to/custom-elements.json]
```

By default, `cem validate` will look for a `custom-elements.json` file in the current directory. You can also provide a path to a different file.

## Options

- `--verbose`, `-v`: Show detailed information including schema version
- `--no-warnings`: Suppress validation warnings

## How it Works

The `validate` command reads the `schemaVersion` field from your manifest and fetches the corresponding schema from `https://unpkg.com/custom-elements-manifest@<version>/schema.json`. Schemas are cached locally for performance.

### Schema Validation
If the manifest is valid against the JSON schema, the command will exit with a `0` status code and print a success message. If the manifest is invalid, it will print detailed validation errors with contextual information and exit with a non-zero status code.

### Intelligent Warnings
Beyond basic schema validation, `cem validate` analyzes your manifest for patterns that are technically valid but may indicate issues with your API documentation:

#### Lifecycle Methods
- **Web Components lifecycle**: `connectedCallback`, `disconnectedCallback`, `attributeChangedCallback`, `adoptedCallback`
- **Lit Element lifecycle**: `firstUpdated`, `updated`, `willUpdate`, `getUpdateComplete`, `performUpdate`, `scheduleUpdate`, `requestUpdate`, `createRenderRoot`
- **Lit Element render method**: `render` (only in Lit elements)
- **Form-associated callbacks**: `formAssociatedCallback`, `formDisabledCallback`, `formResetCallback`, `formStateRestoreCallback`

#### Private Methods and Implementation Details
- **Private methods**: Methods starting with `_` or `#`
- **Static implementation fields**: `styles`, `shadowRootOptions`, `formAssociated`, `observedAttributes`
- **Internal utility methods**: `init`, `destroy`, `dispose`, `cleanup`, `debug`, `log`

#### Superclass Attribution
- **Built-in types**: Warns when built-in types like `HTMLElement` don't have `"module": "global:"`

#### Verbose Content
- **Large CSS defaults**: CSS properties with very long default values

## Configuration

You can disable specific warning rules using the `warnings.disable` configuration option:

```yaml
# .cem.yaml
warnings:
  disable:
    # Disable entire categories
    - lifecycle
    - private
    - implementation
    # Or disable specific rules
    - lifecycle-lit-render
    - implementation-static-styles
    - private-underscore-methods
```

### Available Warning Categories

- `lifecycle` - All lifecycle method warnings
- `private` - Private method warnings (underscore and hash prefixed)
- `implementation` - Implementation detail warnings (static fields)
- `superclass` - Superclass attribution warnings
- `verbose` - Verbose content warnings
- `internal` - Internal utility method warnings

### Specific Warning Rule IDs

#### Lifecycle Rules
- `lifecycle-web-components` - Web Components lifecycle methods
- `lifecycle-lit-methods` - Lit Element lifecycle methods
- `lifecycle-lit-render` - Lit Element render method
- `lifecycle-constructor` - Constructor method
- `lifecycle-form-callbacks` - Form-associated element callbacks

#### Private Method Rules
- `private-underscore-methods` - Methods starting with `_`
- `private-hash-methods` - Methods starting with `#`

#### Implementation Detail Rules
- `implementation-static-styles` - Static styles field
- `implementation-shadow-root-options` - shadowRootOptions field
- `implementation-form-associated` - formAssociated field
- `implementation-observed-attributes` - observedAttributes field

#### Other Rules
- `superclass-builtin-modules` - Built-in superclass module attribution
- `verbose-css-defaults` - Large CSS property defaults
- `internal-utility-methods` - Internal utility methods
