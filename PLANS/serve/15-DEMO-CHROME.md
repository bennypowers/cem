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
