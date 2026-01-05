---
title: Rendering Modes
weight: 100
---

The [dev server][workflow] supports three rendering modes that control how your [demos][demos] are presented: light mode (full development UI), shadow mode (demos in [shadow DOM][shadowdom]), and chromeless mode (minimal standalone pages). Choose the mode that fits your workflow—light for interactive development with [knobs][knobs] and live monitoring, shadow for testing CSS encapsulation, or chromeless for [automated testing](#chromeless-mode) and clean demo sharing.

All modes include [live reload][websocket] with smart dependency tracking and support [query parameter overrides](#query-parameter-override) for quick testing. The rendering mode affects UI presentation and error handling but preserves core functionality like TypeScript transforms, [import map][importmaps] injection, and file watching across all modes.

## Light Mode (Default)

Light mode provides the full development experience with a PatternFly-based UI that includes sidebar navigation, header with theme toggle and debug info, interactive knobs for attributes and properties, real-time event monitoring, server logs with filtering, manifest browser, and visual live reload status. Error overlays appear for transform failures, making debugging straightforward. This is the standard mode for component development.

Use `cem serve` (light mode is the default) or configure it explicitly with `serve.demos.rendering: light` in `.config/cem.yaml`.

## Shadow Mode

Shadow mode provides the same full UI as light mode but renders demo content inside a shadow root, letting you test components in encapsulated contexts. Use this when verifying CSS encapsulation, testing `:host` selectors and shadow piercing, or debugging shadow DOM-specific behaviors.

Start the server with `cem serve --rendering=shadow` or configure it with `serve.demos.rendering: shadow`.

## Chromeless Mode

Chromeless mode strips away all dev server UI, serving demos as clean standalone pages ideal for automated testing with [Playwright][playwright] or Puppeteer, embedding in documentation sites, sharing clean demo URLs, or capturing screenshots. Core functionality remains—live reload, file watching, TypeScript/CSS transforms, and import map injection—but errors log to the console instead of showing overlays, and there's no visual chrome, knobs panel, or connection status indicators.

Start with `cem serve --rendering=chromeless` or configure `serve.demos.rendering: chromeless`. Override per-demo with the `?rendering=chromeless` query parameter.

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

```
http://localhost:8000/elements/button/demo/?rendering=chromeless
http://localhost:8000/elements/button/demo/?rendering=shadow
```

## Live Reload and Error Handling

All modes use [WebSocket][websocket]-based live reload with smart dependency tracking that only triggers reloads when relevant imported files change. Light and shadow modes show visual connection status with reconnection modals and toasts, while chromeless mode reloads silently with console-only logging.

Error handling adapts to each mode: light and shadow display full-screen overlays for TypeScript/CSS transform errors with visual indicators and detailed logs in the UI, while chromeless logs errors to the browser console with a `[CEM]` prefix to maintain clean demo presentation even during errors.

{{<tip "info">}}
Check the browser console when using chromeless mode—all dev server messages and errors are logged there.
{{</tip>}}

[workflow]: ../workflow/
[demos]: ../demos/
[shadowdom]: https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_shadow_DOM
[knobs]: ../knobs/
[websocket]: https://developer.mozilla.org/en-US/docs/Web/API/WebSocket
[importmaps]: ../import-maps/
[playwright]: https://playwright.dev/
