## New Features

### Generate

#### Flattened Union Types

When generating the manifest, `cem` will now try to compute the concrete types
of unions. For example, if a button element defines its `type` field like so:

```ts
type ButtonType = 'button' | 'submit' | 'reset';

@customElement('schwifty-button')
export class SchwiftyButton extends LitElement {
  type?: ButtonType;
}
```

The generated manifest will enumerate the concrete types, with a reference to 
their definition:

```json
{
  "name": "type",
  "type": {
    "text": "'button' | 'submit' | 'reset'",
    "references": [
      {
        "name": "ButtonType",
        "module": "src/schwifty-button.js"
      }
    ]
  }
}
```

Types will resolve across modules in your project, but will not yet resolve 
across packages (e.g. from node_modules).

This also means that these fields' dev server knobs will render as `<select>` 
widgets, instead of plain text inputs.

### Dev Server

#### Rendering Modes

Added configurable rendering modes: `light`, `shadow`, `iframe`, `chromeless`. 
Set via config or flags, or you can override with a query param. In **Light** 
mode (the default), your demos are rendered in the page's light DOM. **Shadow** 
mode renders your demo in a shadow root, **iFrame** mode does what you would 
expect, and **chromeless** mode renders your demo without any extraneous dev 
server UI - perfect for Puppeteer/Playwright tests.

#### URL Rewrites

Users can now set up URL rewrite mappings, for example to map `/dist/` to 
`/src/`, or to rewrite static asset URLs to their locations on disk.

#### Hot Config Reloads

Changing your `cem.yaml` config file will now reload the page with the new 
config.

### Bug Fixes

- `go install bennypowers.dev/cem@latest` now works as expected
- Dev server: Fixed the layout of knobs
- Dev server: Improved state persistence
- Dev server: buffer logs until the live log area initializes
- LSP: fixed duplicate results from the `textDocument/references` method

### Refactoring

- Tests: Consolidated fixture-driven tests for completions, hover, references, and diagnostics
- Hooks: moved the pre-commit hook to pre-push
- Dev server: split up server code into logical sections
- Dev server: resolve deadlocks, improve performance under load
