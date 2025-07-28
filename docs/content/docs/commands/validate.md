---
title: Validate
description: Validate your custom-elements.json against the official schema
---

The `cem validate` command validates your `custom-elements.json` file against its corresponding JSON schema and provides intelligent warnings for potentially inappropriate manifest content.

```bash
cem validate [path/to/custom-elements.json]
```

By default, `cem validate` will look for a `custom-elements.json` file in the current directory. You can also provide a path to a different file.

## Options

- `--verbose`, `-v`: Show detailed information including schema version
- `--disable`: Disable specific warning rules or categories (can be used multiple times)
- `--format`: Output format, either `text` (default) or `json`

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

## Output Formats

### Text Format (Default)
The default text format provides human-readable output with colors and formatting.

### JSON Format
Use `--format json` to get machine-readable output suitable for CI/CD pipelines and tooling:

```bash
cem validate --format json custom-elements.json
```

JSON output structure:
```json
{
  "valid": true,
  "path": "custom-elements.json",
  "schemaVersion": "2.1.1",
  "errors": [
    {
      "id": "schema-required-property",
      "module": "my-element.js",
      "declaration": "class MyElement",
      "message": "required property 'name' is missing",
      "location": "/modules/0/declarations/0"
    }
  ],
  "warnings": [
    {
      "id": "lifecycle-lit-render",
      "module": "my-element.js", 
      "declaration": "class MyElement",
      "member": "method render",
      "message": "render method in Lit element should not be documented in public API",
      "category": "lifecycle"
    }
  ]
}
```

## Configuration

You can disable specific warning rules using configuration or command-line flags:

### Configuration File
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

### Command Line
```bash
# Disable entire categories
cem validate --disable lifecycle --disable private

# Disable specific rules
cem validate --disable lifecycle-lit-render --disable implementation-static-styles

# Combine with JSON output
cem validate --format json --disable lifecycle
```

The `--disable` flag can be used multiple times and will be merged with any disabled rules from your configuration file.

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

### Schema Validation Error IDs

Schema validation errors also include unique IDs for programmatic handling:

- `schema-required-property` - Missing required property
- `schema-additional-properties` - Unexpected additional property
- `schema-invalid-enum` - Invalid enum value
- `schema-invalid-kind` - Invalid declaration kind
- `schema-invalid-type` - Wrong data type
- `schema-invalid-format` - Invalid format (e.g., URI, email)
- `schema-invalid-pattern` - String doesn't match required pattern
- `schema-value-too-small` - Number below minimum
- `schema-value-too-large` - Number above maximum
- `schema-string-too-short` - String shorter than minLength
- `schema-string-too-long` - String longer than maxLength
- `schema-array-too-short` - Array shorter than minItems
- `schema-array-too-long` - Array longer than maxItems
- `schema-duplicate-items` - Array contains duplicate items
- `schema-validation-error` - Generic validation error
