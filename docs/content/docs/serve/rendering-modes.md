---
title: Rendering Modes
weight: 15
---

The dev server supports three rendering modes for serving demos, allowing you to choose the right environment for your workflow.

## Light Mode (Default)

Full development UI with all features enabled. This is the standard mode for component development.

**Features:**
- PatternFly-based UI with sidebar navigation
- Header with branding, theme toggle, and debug info
- Interactive knobs for attributes, properties, and CSS custom properties
- Real-time event monitoring tab
- Server logs viewer with filtering
- Manifest browser
- Live reload with visual connection status
- Error overlays for transform failures

**Usage:**
```bash
cem serve
# or explicitly:
cem serve --rendering=light
```

**Config:**
```yaml
serve:
  demos:
    rendering: light
```

## Shadow Mode

Same UI as light mode, but renders demo content in a Shadow DOM. Use this when testing components that need to be encapsulated in a shadow root.

**Use cases:**
- Testing components inside Shadow DOM contexts
- Verifying CSS encapsulation
- Testing `:host` selectors and shadow piercing
- Debugging shadow DOM-specific behaviors

**Usage:**
```bash
cem serve --rendering=shadow
```

**Config:**
```yaml
serve:
  demos:
    rendering: shadow
```

## Chromeless Mode

Minimal rendering with no dev server UI. Demos are served as standalone pages with only live reload functionality.

**Features:**
- ✅ Live reload (silent, console-only)
- ✅ File watching with smart dependency tracking
- ✅ TypeScript/CSS transforms
- ✅ Import map injection
- ❌ No visual UI chrome (header, sidebar, drawer)
- ❌ No knobs panel
- ❌ No error overlays (errors logged to console)
- ❌ No connection status indicators

**Use cases:**
- **Automated testing**: Run Playwright/Puppeteer tests without UI interference
- **Isolated development**: Focus on component behavior without distractions
- **Embedding**: Embed demos in documentation sites or other applications
- **Simple demos**: Share clean demo URLs without dev tooling visible
- **Screenshots**: Capture component demos without chrome UI elements

**Usage:**
```bash
cem serve --rendering=chromeless
```

**Config:**
```yaml
serve:
  demos:
    rendering: chromeless
```

**Per-demo override:**
```
http://localhost:8000/elements/button/demo/?rendering=chromeless
```

## Playwright Integration Example

Chromeless mode is ideal for end-to-end testing with Playwright:

```javascript
// playwright.config.js
export default {
  webServer: {
    command: 'cem serve --rendering=chromeless',
    port: 8000,
    reuseExistingServer: !process.env.CI,
  },
  use: {
    baseURL: 'http://localhost:8000',
  },
};

// tests/button.spec.js
import { test, expect } from '@playwright/test';

test('button component renders correctly', async ({ page }) => {
  await page.goto('/elements/my-button/demo/');

  // Test the component directly without UI chrome interference
  const button = page.locator('my-button');
  await expect(button).toBeVisible();
  await button.click();
  await expect(button).toHaveAttribute('aria-pressed', 'true');
});
```

## Query Parameter Override

You can override the default rendering mode for any demo using the `?rendering=` query parameter:

```
# Use chromeless for a specific demo
http://localhost:8000/elements/button/demo/?rendering=chromeless

# Test in shadow mode temporarily
http://localhost:8000/elements/button/demo/?rendering=shadow
```

This is useful for:
- Testing different rendering contexts without restarting the server
- Sharing clean demo links
- Quick comparisons between modes

## Live Reload Behavior

All rendering modes support live reload with smart dependency tracking:

- **Light & Shadow**: Visual connection status with reconnection modals and toasts
- **Chromeless**: Silent reload with console logging only

The WebSocket client tracks which files the demo imports and only triggers reloads when relevant files change.

## Error Handling

Error handling varies by mode:

**Light & Shadow:**
- Full-screen error overlays for TypeScript/CSS transform errors
- Visual error indicators in the UI
- Error details in the logs tab

**Chromeless:**
- Errors logged to browser console with `[CEM]` prefix
- No visual overlays
- Clean demo presentation even during errors

{{<tip "info">}}
Check the browser console when using chromeless mode - all dev server messages and errors are logged there.
{{</tip>}}
