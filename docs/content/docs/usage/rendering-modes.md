---
title: Rendering Modes
weight: 100
---

The [dev server][workflow] supports three rendering modes that control how your [demos][demos] are presented: light mode (full development UI), shadow mode (demos in [shadow DOM][shadowdom]), and chromeless mode (minimal standalone pages). Choose the mode that fits your workflow—light for interactive development with [knobs][knobs] and live monitoring, shadow for testing CSS encapsulation, or chromeless for [automated testing](#chromeless-mode) and clean demo sharing.

All modes include [live reload][websocket] with smart dependency tracking and support [query parameter overrides](#query-parameter-override) for quick testing. The rendering mode affects UI presentation and error handling but preserves core functionality like TypeScript transforms, [import map][importmaps] injection, and file watching across all modes.

## Light Mode (Default)

Light mode provides the full development experience with a PatternFly-based UI that includes sidebar navigation, header with theme toggle and debug info, interactive knobs for attributes and properties, real-time event monitoring, server logs with filtering, manifest browser, and visual live reload status. Error overlays appear for transform failures, making debugging straightforward. This is the standard mode for component development.

Use `cem serve` (light mode is the default) or configure it explicitly in `.config/cem.yaml`:

```yaml
serve:
  demos:
    rendering: light
```

If the `rendering` option is omitted or empty, demos default to `light` mode.

## Shadow Mode

Shadow mode provides the same full UI as light mode but renders demo content inside a shadow root, letting you test components in encapsulated contexts. Use this when verifying CSS encapsulation, testing `:host` selectors and shadow piercing, or debugging shadow DOM-specific behaviors.

Start the server with `cem serve --rendering=shadow` or configure it in `.config/cem.yaml`:

```yaml
serve:
  demos:
    rendering: shadow
```

## Chromeless Mode

Chromeless mode strips away all dev server UI, serving demos as clean standalone pages ideal for automated testing with [Playwright][playwright] or Puppeteer, embedding in documentation sites, sharing clean demo URLs, or capturing screenshots. Core functionality remains—live reload, file watching, TypeScript/CSS transforms, and import map injection—but errors log to the console instead of showing overlays, and there's no visual chrome, knobs panel, or connection status indicators.

Start with `cem serve --rendering=chromeless` or configure it in `.config/cem.yaml`:

```yaml
serve:
  demos:
    rendering: chromeless
```

Override per-demo with the `?rendering=chromeless` query parameter.

### Playwright Integration

Configure Playwright to use chromeless mode for clean component testing without UI interference:

```javascript
// playwright.config.js
export default {
  webServer: {
    command: 'cem serve --rendering=chromeless',
    port: 8000,
    reuseExistingServer: !process.env.CI,
  },
  use: { baseURL: 'http://localhost:8000' },
};

// tests/button.spec.js
import { test, expect } from '@playwright/test';

test('button component', async ({ page }) => {
  await page.goto('/elements/my-button/demo/');
  const button = page.locator('my-button');
  await expect(button).toBeVisible();
  await button.click();
  await expect(button).toHaveAttribute('aria-pressed', 'true');
});
```

## Query Parameter Override

Override the default rendering mode for any demo with the `?rendering=` query parameter—useful for testing different contexts without restarting the server, sharing clean demo links, or quick mode comparisons:

```text
http://localhost:8000/elements/button/demo/?rendering=chromeless
http://localhost:8000/elements/button/demo/?rendering=shadow
http://localhost:8000/elements/button/demo/?rendering=light
```

Valid values are `light`, `shadow`, and `chromeless`. Invalid values are ignored and the configured default is used.

{{<tip "warning">}}
**Iframe mode not yet implemented:** Specifying `iframe` mode in the config will cause the server to fail at startup. If `iframe` is requested via query parameter, the server logs a warning, broadcasts an error overlay, and falls back to `shadow` mode.
{{</tip>}}

### Backward Compatibility

The legacy `?shadow=true` query parameter is still supported and will override to shadow mode:

```text
http://localhost:8000/elements/button/demo/?shadow=true
```

## Live Reload and Error Handling

All modes use [WebSocket][websocket]-based live reload with smart dependency tracking that only triggers reloads when relevant imported files change. Light and shadow modes show visual connection status with reconnection modals and toasts, while chromeless mode reloads silently with console-only logging.

Error handling adapts to each mode: light and shadow display full-screen overlays for TypeScript/CSS transform errors with visual indicators and detailed logs in the UI, while chromeless logs errors to the browser console with a `[CEM]` prefix to maintain clean demo presentation even during errors.

{{<tip "info">}}
Check the browser console when using chromeless mode—all dev server messages and errors are logged there.
{{</tip>}}

## Use Cases by Mode

**Light DOM (default):**
- General component testing with full dev UI
- Demos that rely on global styles
- Testing how components integrate with the parent page
- Interactive development with knobs and event monitoring

**Shadow DOM:**
- Testing encapsulation behavior
- Verifying CSS custom properties penetrate shadow boundaries
- Testing `::part()` and `::slotted()` selectors
- Ensuring styles don't leak in or out

**Chromeless:**
- Automated testing with Playwright/Puppeteer
- Embedding demos in documentation sites
- Sharing clean demo URLs without dev UI
- Capturing screenshots for documentation

## Configuration Examples

**Default all demos to shadow mode:**

```yaml
serve:
  demos:
    rendering: shadow
```

Then override specific demos back to light mode when needed:

```text
http://localhost:8000/elements/integration-test/demo/?rendering=light
```

**Use chromeless for CI testing:**

```yaml
# .config/cem.ci.yaml
serve:
  demos:
    rendering: chromeless
```

```sh
# In CI pipeline
cem serve --config .config/cem.ci.yaml
```

[workflow]: ../workflow/
[demos]: ../demos/
[shadowdom]: https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_shadow_DOM
[knobs]: ../knobs/
[websocket]: https://developer.mozilla.org/en-US/docs/Web/API/WebSocket
[importmaps]: ../import-maps/
[playwright]: https://playwright.dev/
