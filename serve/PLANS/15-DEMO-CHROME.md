## Demo Chrome

## TDD First Steps

Phase 3 has no stubs in existing tests. Follow standard TDD: write tests first, implement to make them pass.

### Server Side

```html
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>{{.DemoTitle}} - {{.TagName}}</title>
  <script type="importmap">{{.ImportMap}}</script>
  <script src="/__cem-serve-chrome.js"></script>
</head>
<body>
  <cem-serve-chrome role="none"
                    knobs="{{.EnabledKnobs}}"
                    tag-name="{{.TagName}}">
    <template shadowrootmode="open">
      <style>/*...*/</style>
      <header>
        <slot name="title"></slot>
        <nav><!-- demo switcher --></nav>
        {{if .Demo.Source}}
        <a href="{{.Demo.Source}}"
           target="_blank"
           title="View Source on {{.Demo.Source | deriveVCSHostNameFromURL}}">{{.Demo.Source | deriveVCSIconFromURL}}</a>
        {{end}}
      </header>
      {{if ne .EnabledKnobs ""}}
      <aside id="knobs">
        {{range .ElementsInDemo}}
          {{ template "knob-group" . }}
        {{end}}
      </aside>
      {{end}}
      <main>
        <slot></slot>
      </main>
      <footer>
        <slot name="description"></slot>
      </footer>
    </template>
    <h1 slot="title">{{.TagName}}</h1>
    {{if .Demo.Description}}
    <section slot="description">{{.Demo.Description | markdown}}</section>
    {{end}}

    <!-- Demo in light DOM by default -->
    <cem-serve-demo id="demo">
      {{.DemoHTML}}
    </cem-serve-demo>
  </cem-serve-chrome>
</body>
</html>
```

- `EnabledKnobs` is a space-separated list of well known strings, computed from the config and url params, `attributes slots properties css-properties`.
- `deriveVSC*` functions are pure, and support GitHub, GitLab, others down the 
line.
- `markdown` does GFM, escaping HTML. Relative links are resolved against file 
path.

### Client side

embedded `__cem-serve-chrome.js`:

```typescript
class CemServeDemo extends HTMLElement {
  // ...
  setAttribute(attribute: string, value: string): boolean
  setProperty(property: string, value: unknown): boolean
  setCssCustomProperty(cssCustomProperty: string, value: string): boolean
  // ...
}

class CemServeChrome extends HTMLElement {
  static is = 'cem-serve-chrome';

  static { customElements.define(this.is, this); }

  // Chrome UI in shadow root
  #knobs = this.shadowRoot.getElementById('knobs');

  // Demo in light DOM (direct child)
  #demo = this.querySelector('#demo');

  // ...
}
```

## Demo Isolation Strategy

### Light DOM by Default

Demos render in **light DOM** for ergonomic script access:

**Rationale:**

1. **Script ergonomics** - Demo scripts use `document.querySelector()` naturally:
   ```javascript
   <script type="module">
     import '@rhds/elements/rh-tabs/rh-tabs.js';

     const form = document.querySelector('form');  // ✅ Works
     const tabs = document.querySelector('rh-tabs'); // ✅ Works

     form.addEventListener('input', () => {
       tabs.setAttribute('variant', ...);
     });
   </script>
   ```

2. **Module scripts work** - No `document.currentScript` issues (it's always null in modules)

3. **Chrome still isolated** - Chrome UI in shadow root prevents:
   - Chrome styles leaking to demo
   - Demo styles affecting chrome
   - Demo scripts accidentally selecting chrome elements

4. **No migration burden** - Existing demos (like RHDS `rh-tabs/demo/color-context.html`) work as-is

5. **Familiar patterns** - Matches how developers write HTML/JS naturally

### Shadow Mode for Isolation Testing

Optional `?shadow=true` query parameter wraps demo in shadow root:

```html
<!-- When ?shadow=true -->
<cem-serve-demo id="demo">
  <template shadowrootmode="open">
    {{.DemoHTML}} <!-- Now isolated in shadow -->
  </template>
</cem-serve-demo>
```

**Use cases:**
- Testing component works in nested shadow roots
- Verifying component doesn't leak styles
- Debugging shadow DOM-specific issues
- Ensuring component is truly encapsulated

**Trade-offs:**
- ❌ Demo scripts must use `getRootNode()` for queries
- ❌ Can't use `document.currentScript` in module scripts
- ✅ Complete isolation of demo from document
- ✅ Tests real shadow DOM encapsulation

### Knobs in Shadow Mode

When `?shadow=true` is enabled, knobs use `getRootNode()` to access demo elements inside the shadow root.

**Implementation in CemServeDemo:**

```typescript
class CemServeDemo extends HTMLElement {
  #getShadowRoot(): ShadowRoot | null {
    // Check if demo content is in shadow root
    const template = this.querySelector('template[shadowrootmode]');
    return template ? this.shadowRoot : null;
  }

  setAttribute(attribute: string, value: string): boolean {
    const root = this.#getShadowRoot() || this;
    const target = root.querySelector(`[${attribute}]`);
    if (target) {
      target.setAttribute(attribute, value);
      return true;
    }
    return false;
  }

  setProperty(property: string, value: unknown): boolean {
    const root = this.#getShadowRoot() || this;
    const target = root.querySelector('[data-knob-target]');
    if (target) {
      target[property] = value;
      return true;
    }
    return false;
  }

  setCssCustomProperty(cssProperty: string, value: string): boolean {
    const root = this.#getShadowRoot() || this;
    const target = root.querySelector('[data-knob-target]');
    if (target) {
      target.style.setProperty(cssProperty, value);
      return true;
    }
    return false;
  }
}
```

**Behavior:**
- **Light DOM mode** (default): `root = this`, queries work normally
- **Shadow mode** (`?shadow=true`): `root = this.shadowRoot`, queries work inside shadow root
- Knobs continue to function in both modes without user intervention

See [50-KNOBS-CORE.md](./50-KNOBS-CORE.md) for complete knobs implementation details.

### Custom Element Name Collisions
Not an issue. cem-serve- is a specific-enough prefix.

### Template Injection Security
**Issue:** Server renders user-provided demo HTML (line 46) in templates.
- Are demos sanitized against XSS? It's dev server, XSS isn't a huge concern
- Can demos inject script into chrome template? Yes, in fact they need to to 
load elements or use their JS APIs.
- Template functions like `deriveVCSHostNameFromURL` (line 26) - are they safe? 
  takes string and returns string.

**Recommendation:**
- Document template escaping strategy
- Use Go's `html/template` auto-escaping
- Validate demo metadata before rendering

### Browser support
We target baseline support, so no need to polyfill DSD, ecma private fields, etc.
We can issue a browser warning based on user agent if someone tries to load a dev server page with ie6 or whatever.

### Demo Switcher Navigation
- Nav is SSRd by package and module (grouped by normalized primary tag name)
- Current demo is highlighted `aria-active`

---

## Acceptance Criteria

- [ ] Demo chrome template renders with SSR (Declarative Shadow DOM)
- [ ] Chrome UI in shadow root, demo in light DOM
- [ ] `<cem-serve-chrome>` custom element defined and registered
- [ ] `<cem-serve-demo>` custom element defined with helper methods
- [ ] Import map injected into `<head>` before demo scripts
- [ ] `__cem-serve-chrome.js` script loaded and executed
- [ ] Demo switcher nav shows all demos for current element
- [ ] Current demo highlighted in demo switcher
- [ ] VCS source link renders with correct icon (GitHub, GitLab)
- [ ] Demo description renders as GFM markdown with HTML escaping
- [ ] Relative links in markdown resolve against demo file path
- [ ] Query parameter `?shadow=true` wraps demo in shadow root
- [ ] Light DOM mode (default) allows `document.querySelector()` in demo scripts
- [ ] Module scripts in demos work without `currentScript` issues
- [ ] Knobs panel renders when knobs enabled
- [ ] Server errors logged to browser console
- [ ] Browser compatibility check warns on unsupported browsers (DSD requirement)
- [ ] Tests cover template rendering, shadow/light DOM modes
