## Demo Chrome

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
           title="View Source on {{.Demo.Source | deriveVCSHostNameFromURL}}">{{.Demo.Source | deriveIconFromURL}}</a>
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
    <cem-serve-demo id="demo">
      <template shadowrootmode="open">
        {{.DemoHTML}}
      </template>
    </cem-serve-demo>
    {{if .Demo.Description}}
    <section slot="description">{{.Demo.Description | markdown}}</section>
    {{end}}
  </cem-serve-chrome>
</body>
</html>
```

`EnabledKnobs` is a space-separated list of well known strings, computed from the config and url params, `attributes slots properties css-properties`.

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

  // already present due to SSR
  #demo = this.shadowRoot.getElementById('demo');
  #knobs = this.shadowRoot.getElementById('knobs');

  // ...
}
```

---

## Open Questions & Concerns

### Shadow DOM Boundary Issues
**Issue:** Demo HTML is wrapped in shadow DOM (lines 44-48):
```html
<cem-serve-demo id="demo">
  <template shadowrootmode="open">
    {{.DemoHTML}}
  </template>
</cem-serve-demo>
```

**Problems this creates:**
1. **Style inheritance broken** - Demos expecting document-level styles won't work
2. **Slots don't cross boundaries** - If demo uses slotting, won't work across shadow root
3. **Global CSS selectors fail** - Demos can't target elements outside shadow root
4. **Focus management** - `document.activeElement` won't reflect focused element in shadow

**Examples that break:**
```css
/* Won't work - outside shadow root */
body { font-family: sans-serif; }
```

```html
<!-- Won't work - slot can't cross shadow boundary -->
<template id="user-template">
  <slot name="avatar"></slot>
</template>
```

**Recommendation:** Provide both shadow and light DOM modes:
- Default: light DOM for maximum compatibility
- `?shadow-demo=true` for isolation testing

### Custom Element Name Collisions
**Footgun:** Chrome uses `<cem-serve-chrome>` and `<cem-serve-demo>` (lines 15, 44).
- What if user's demos define elements with same names?
- Custom elements registry is global - redefinition throws error

**Example collision:**
```javascript
// User demo file
customElements.define('cem-serve-chrome', class extends HTMLElement {});
// Error: 'cem-serve-chrome' has already been defined
```

**Recommendation:** Use unlikely-to-collide prefix:
- `<__cem-internal-chrome>`
- `<__cem-internal-demo>`

Or use UUIDs in name: `<cem-serve-chrome-a3f2b1>`

### Template Injection Security
**Issue:** Server renders user-provided demo HTML (line 46) in templates.
- Are demos sanitized against XSS?
- Can demos inject script into chrome template?
- Template functions like `deriveVCSHostNameFromURL` (line 26) - are they safe?

**Attack vector:**
```html
<!-- Malicious demo -->
<meta itemprop="name" content="</template><script>steal()</script>">
```

**Recommendation:**
- Document template escaping strategy
- Use Go's `html/template` auto-escaping
- Validate demo metadata before rendering

### Source Link VCS Icon Derivation
**Issue:** Line 26 shows `{{.Demo.Source | deriveIconFromURL}}` but:
- No specification for which VCS hosts are supported
- What if URL is invalid?
- What if icon resource fails to load?

**Recommendation:** Define supported VCS list, fallback icon strategy.

### Demo Description Markdown Rendering
**Issue:** Line 50 shows `{{.Demo.Description | markdown}}` but:
- Which markdown flavor? (CommonMark? GFM?)
- Is HTML in markdown allowed? (XSS risk)
- What about relative links in markdown?

**Recommendation:**
- Use CommonMark with strict HTML escaping
- Document markdown subset supported
- Resolve relative links against demo file path

### Client-Side Hydration
**Issue:** Templates use Declarative Shadow DOM (DSD) which requires:
- Browser support (Chrome 90+, Safari 16.4+, Firefox 123+)
- No polyfill available for older browsers

**Footgun:** Demos won't render in older browsers without warning.

**Recommendation:**
- Add browser compatibility check on page load
- Show error message in unsupported browsers
- Document minimum browser versions

### Demo Switcher Navigation
**Issue:** Line 22 shows `<nav><!-- demo switcher -->` but no details:
- How are demos discovered for switcher?
- Client-side or server-side rendered?
- How to handle many demos (>50)?

**Recommendation:** Specify demo switcher implementation:
- Server renders list from manifest
- Highlight current demo
- Group by category if metadata available