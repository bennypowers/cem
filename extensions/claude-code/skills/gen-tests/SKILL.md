---
name: gen-tests
description: >
  Generate test scaffolding for custom elements using cem serve's chromeless
  demo UI with Playwright or Puppeteer. Use when the user asks to
  "generate tests", "scaffold tests", "create test boilerplate",
  "write tests for element", "playwright tests", "puppeteer tests",
  or mentions "test coverage" for custom elements.
tools: Read, Glob, Grep, Bash
---

# Test Scaffolding Generator

Generate browser-based test scaffolding for custom elements using the
Page Object Model pattern, testing against `cem serve`'s chromeless demo UI
with Playwright or Puppeteer.

## Workflow

### Phase 1: Gather Element Data

Read the target element's full manifest data:

```text
cem://element/{tagName}
cem://element/{tagName}/attributes
cem://element/{tagName}/slots
cem://element/{tagName}/events
cem://element/{tagName}/css/parts
cem://element/{tagName}/css/custom-properties
cem://element/{tagName}/css/states
```

Check which demos exist for the element — the manifest's `demos` array defines
the available demo pages and their URLs.

### Phase 2: Detect Test Environment

Search the project for existing test patterns:

1. Look for existing test files to match conventions:
   - `*.test.ts`, `*.spec.ts` patterns
   - Test framework: Playwright (`@playwright/test`) or Puppeteer (`puppeteer`)
   - Existing `playwright.config.ts` or `puppeteer` setup
   - Existing page objects in `tests/pages/`, `tests/models/`, etc.

2. Check `package.json` for test dependencies and scripts

3. If no test framework is present, recommend Playwright and generate a config

### Phase 3: Generate Playwright Config (if needed)

If the project doesn't have a Playwright config, generate one that starts
`cem serve` in chromeless mode:

```ts
// playwright.config.ts
import { defineConfig } from '@playwright/test';

export default defineConfig({
  webServer: {
    command: 'cem serve --rendering=chromeless',
    port: 8000,
    reuseExistingServer: !process.env.CI,
  },
  use: {
    baseURL: 'http://localhost:8000',
  },
});
```

For Puppeteer, generate equivalent setup/teardown that spawns `cem serve`.

### Phase 4: Generate Page Object

Generate a page object class for the element that encapsulates all locators
and interactions derived from the manifest. The page object is the primary
output — tests become thin and readable when the page object does the work.

#### Page Object Structure

```ts
// tests/pages/MyButtonPage.ts
import type { Locator, Page } from '@playwright/test';

export class MyButtonPage {
  readonly page: Page;

  /** The <my-button> element */
  readonly host: Locator;

  // --- Slots (from manifest slots) ---

  /** Default slot content */
  readonly slotDefault: Locator;
  /** The icon slot */
  readonly slotIcon: Locator;

  constructor(page: Page) {
    this.page = page;
    this.host = page.locator('my-button');
    this.slotDefault = this.host.locator(':scope > :not([slot])');
    this.slotIcon = this.host.locator('[slot="icon"]');
  }

  // --- Navigation ---

  async goto(demo = '') {
    const path = demo
      ? `/elements/my-button/demo/${demo}`
      : '/elements/my-button/demo/';
    await this.page.goto(path);
  }

  // --- Attribute helpers (from manifest attributes) ---

  async variant(): Promise<string | null> {
    return this.host.getAttribute('variant');
  }

  async setVariant(value: 'primary' | 'secondary' | 'danger') {
    await this.host.evaluate((el, v) => el.setAttribute('variant', v), value);
  }

  async disabled(): Promise<boolean> {
    return this.host.evaluate((el) => el.hasAttribute('disabled'));
  }

  async setDisabled(value: boolean) {
    await this.host.evaluate(
      (el, v) => el.toggleAttribute('disabled', v),
      value,
    );
  }

  // --- Property helpers (for non-attribute properties) ---

  async setComplexData(value: unknown) {
    await this.host.evaluate((el, v) => (el as any).complexData = v, value);
  }

  // --- Event helpers (from manifest events) ---

  /**
   * Returns a promise that resolves when the element fires the given event.
   * Call this BEFORE triggering the interaction.
   */
  async waitForEvent<T = unknown>(eventName: string): Promise<T> {
    return this.host.evaluate(
      (el, name) =>
        new Promise<T>((resolve) =>
          el.addEventListener(name, (ev) => resolve(ev as T), { once: true }),
        ),
      eventName,
    );
  }

  // --- CSS helpers (from manifest CSS custom properties) ---

  async setCssProperty(name: string, value: string) {
    await this.host.evaluate(
      (el, [n, v]) => el.style.setProperty(n, v),
      [name, value] as const,
    );
  }

  async computedStyle(property: string, internalSelector?: string): Promise<string> {
    return this.host.evaluate(
      (el, [prop, selector]) => {
        const target = selector
          ? el.shadowRoot?.querySelector(selector) ?? el
          : el;
        return getComputedStyle(target).getPropertyValue(prop);
      },
      [property, internalSelector ?? null] as const,
    );
  }

  // --- State helpers (from manifest CSS states) ---

  async matchesState(state: string): Promise<boolean> {
    return this.host.evaluate((el, s) => el.matches(`:state(${s})`), state);
  }

  // --- Interaction helpers ---

  async click() {
    await this.host.click();
  }

  async focus() {
    await this.host.focus();
  }
}
```

#### Page Object Generation Rules

For each manifest feature, generate the corresponding page object members:

| Manifest Feature | Page Object Member |
|------------------|--------------------|
| Each attribute | Getter method + setter method (typed for enums) |
| Each property (no attribute) | Setter via `evaluate` |
| Each named slot | `Locator` field: `this.host.locator('[slot="name"]')` |
| Default slot | `Locator` field: `this.host.locator(':scope > :not([slot])')` |
| Each CSS part | Helper method using `evaluate` to query shadow root |
| Each CSS custom property | Helper via `setCssProperty` / `computedStyle` |
| Each CSS state | `matchesState(name)` boolean helper |
| Each event | Typed `waitForEvent` usage or dedicated method |
| Each demo URL | Named navigation method or param to `goto()` |

### Phase 5: Generate Test Cases

Tests use the page object — they should read almost like plain English.

```ts
// tests/my-button.spec.ts
import { test, expect } from '@playwright/test';
import { MyButtonPage } from './pages/MyButtonPage';

test.describe('<my-button>', () => {
  let button: MyButtonPage;

  test.beforeEach(async ({ page }) => {
    button = new MyButtonPage(page);
    await button.goto();
  });

  test.describe('attributes', () => {
    test('has primary variant by default', async () => {
      await expect(button.host).toHaveAttribute('variant', 'primary');
    });

    test('reflects variant attribute', async () => {
      await button.setVariant('secondary');
      await expect(button.host).toHaveAttribute('variant', 'secondary');
    });

    test('supports all variant values', async () => {
      for (const variant of ['primary', 'secondary', 'danger'] as const) {
        await button.setVariant(variant);
        await expect(button.host).toHaveAttribute('variant', variant);
      }
    });

    test('toggles disabled', async () => {
      expect(await button.disabled()).toBe(false);
      await button.setDisabled(true);
      expect(await button.disabled()).toBe(true);
    });
  });

  test.describe('slots', () => {
    test('projects default slot content', async () => {
      await expect(button.slotDefault.first()).toBeVisible();
    });

    test('projects icon slot content', async () => {
      // Only test if the demo includes icon slot content
      await expect(button.slotIcon).toBeAttached();
    });
  });

  test.describe('events', () => {
    test('fires click event', async () => {
      const event = button.waitForEvent('click');
      await button.click();
      expect(await event).toBeTruthy();
    });
  });

  test.describe('css custom properties', () => {
    test('--my-button-color overrides text color', async () => {
      await button.setCssProperty('--my-button-color', 'red');
      const color = await button.computedStyle('color', '#button');
      expect(color).toBe('rgb(255, 0, 0)');
    });
  });

  test.describe('css states', () => {
    test('matches :state(loading) when loading', async () => {
      await button.host.evaluate((el: any) => (el.loading = true));
      expect(await button.matchesState('loading')).toBe(true);
    });
  });

  test.describe('accessibility', () => {
    test('has button role', async () => {
      await expect(button.host).toHaveRole('button');
    });

    test('is keyboard focusable', async ({ page }) => {
      await page.keyboard.press('Tab');
      await expect(button.host).toBeFocused();
    });
  });

  test.describe('visual regression', () => {
    test('default appearance', async () => {
      await expect(button.host).toHaveScreenshot('my-button-default.png');
    });
  });
});
```

### Phase 6: File Layout

Generate both files:

```text
tests/
├── pages/
│   └── MyButtonPage.ts       # Page object
└── my-button.spec.ts          # Test cases
```

Or match the project's existing test directory structure if one exists.

## Chromeless Demo Reference

`cem serve` provides a chromeless rendering mode designed for testing:

- **Start**: `cem serve --rendering=chromeless`
- **Demo URL**: `http://localhost:8000/elements/{tag-name}/demo/`
- **What it provides**: Bare demo HTML with live reload, import maps, and
  TypeScript/CSS transforms — no nav, knobs panel, or dev server UI
- **Query param override**: Append `?rendering=chromeless` to any demo URL
  to get chromeless mode even when the server runs in default mode

## Guidelines

- **Always generate a page object**: The page object is the main deliverable — tests should be thin wrappers
- **One page object per element**: Named `{PascalCaseTag}Page` (e.g., `MyButtonPage`)
- **Type attribute setters with manifest enums**: If `variant` accepts `'primary' | 'secondary'`, type the setter parameter
- **Default to Playwright**: Unless the project already uses Puppeteer
- **Navigate to real demos**: Use the element's actual demo pages, not blank pages with manually constructed HTML
- **Test behavior, not implementation**: Test the public API via the browser, not internal DOM structure
- **One assertion per test** (where practical): Makes failures easy to diagnose
- **Descriptive test names**: `'reflects the variant attribute'` not `'works'`
- **Mark TODOs**: Where the manifest doesn't provide enough info to write a
  complete assertion, add a `// TODO:` comment explaining what to fill in
- **Don't over-generate**: Skip categories that don't apply (e.g., no CSS
  state tests if the element has no custom states, no slot tests if no slots)
- **Use demo variants**: If the manifest has multiple demos, test the relevant
  variant for each feature (e.g., test loading state against a loading demo)
- **For Puppeteer**: Adapt the page object to use `page.$eval`, `page.waitForSelector`, etc. instead of Playwright locators, but keep the same class structure and method names
