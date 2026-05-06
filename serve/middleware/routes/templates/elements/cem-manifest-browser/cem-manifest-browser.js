var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __knownSymbol = (name, symbol) => (symbol = Symbol[name]) ? symbol : /* @__PURE__ */ Symbol.for("Symbol." + name);
var __typeError = (msg) => {
  throw TypeError(msg);
};
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __name = (target, value) => __defProp(target, "name", { value, configurable: true });
var __decoratorStart = (base) => [, , , __create(base?.[__knownSymbol("metadata")] ?? null)];
var __decoratorStrings = ["class", "method", "getter", "setter", "accessor", "field", "value", "get", "set"];
var __expectFn = (fn) => fn !== void 0 && typeof fn !== "function" ? __typeError("Function expected") : fn;
var __decoratorContext = (kind, name, done, metadata, fns) => ({ kind: __decoratorStrings[kind], name, metadata, addInitializer: (fn) => done._ ? __typeError("Already initialized") : fns.push(__expectFn(fn || null)) });
var __decoratorMetadata = (array, target) => __defNormalProp(target, __knownSymbol("metadata"), array[3]);
var __runInitializers = (array, flags, self, value) => {
  for (var i = 0, fns = array[flags >> 1], n = fns && fns.length; i < n; i++) flags & 1 ? fns[i].call(self) : value = fns[i].call(self, value);
  return value;
};
var __decorateElement = (array, flags, name, decorators, target, extra) => {
  var fn, it, done, ctx, access, k = flags & 7, s2 = !!(flags & 8), p = !!(flags & 16);
  var j = k > 3 ? array.length + 1 : k ? s2 ? 1 : 2 : 0, key = __decoratorStrings[k + 5];
  var initializers = k > 3 && (array[j - 1] = []), extraInitializers = array[j] || (array[j] = []);
  var desc = k && (!p && !s2 && (target = target.prototype), k < 5 && (k > 3 || !p) && __getOwnPropDesc(k < 4 ? target : { get [name]() {
    return __privateGet(this, extra);
  }, set [name](x) {
    return __privateSet(this, extra, x);
  } }, name));
  k ? p && k < 4 && __name(extra, (k > 2 ? "set " : k > 1 ? "get " : "") + name) : __name(target, name);
  for (var i = decorators.length - 1; i >= 0; i--) {
    ctx = __decoratorContext(k, name, done = {}, array[3], extraInitializers);
    if (k) {
      ctx.static = s2, ctx.private = p, access = ctx.access = { has: p ? (x) => __privateIn(target, x) : (x) => name in x };
      if (k ^ 3) access.get = p ? (x) => (k ^ 1 ? __privateGet : __privateMethod)(x, target, k ^ 4 ? extra : desc.get) : (x) => x[name];
      if (k > 2) access.set = p ? (x, y) => __privateSet(x, target, y, k ^ 4 ? extra : desc.set) : (x, y) => x[name] = y;
    }
    it = (0, decorators[i])(k ? k < 4 ? p ? extra : desc[key] : k > 4 ? void 0 : { get: desc.get, set: desc.set } : target, ctx), done._ = 1;
    if (k ^ 4 || it === void 0) __expectFn(it) && (k > 4 ? initializers.unshift(it) : k ? p ? extra = it : desc[key] = it : target = it);
    else if (typeof it !== "object" || it === null) __typeError("Object expected");
    else __expectFn(fn = it.get) && (desc.get = fn), __expectFn(fn = it.set) && (desc.set = fn), __expectFn(fn = it.init) && initializers.unshift(fn);
  }
  return k || __decoratorMetadata(array, target), desc && __defProp(target, name, desc), p ? k ^ 4 ? extra : desc : target;
};
var __accessCheck = (obj, member, msg) => member.has(obj) || __typeError("Cannot " + msg);
var __privateIn = (member, obj) => Object(obj) !== obj ? __typeError('Cannot use the "in" operator on this value') : member.has(obj);
var __privateGet = (obj, member, getter) => (__accessCheck(obj, member, "read from private field"), getter ? getter.call(obj) : member.get(obj));
var __privateSet = (obj, member, value, setter) => (__accessCheck(obj, member, "write to private field"), setter ? setter.call(obj, value) : member.set(obj, value), value);
var __privateMethod = (obj, member, method) => (__accessCheck(obj, member, "access private method"), method);

// elements/cem-manifest-browser/cem-manifest-browser.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-manifest-browser/cem-manifest-browser.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('"/* Manifest Browser - Component API navigation tree */\\n\\n:host {\\n  display: block;\\n  height: 100%;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n#drawer {\\n  height: 100%;\\n}\\n\\n#drawer-content {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n  overflow: hidden;\\n}\\n\\n#tree-wrapper {\\n  flex: 1;\\n  overflow: auto;\\n  min-height: 0;\\n}\\n\\n::slotted(cem-pf-v6-tree-view) {\\n  margin: var(--pf-t--global--spacer--md, 1rem);\\n}\\n\\n#panel-title {\\n  margin: 0;\\n}\\n\\n#details-content {\\n  padding: var(--pf-t--global--spacer--md, 1rem);\\n\\n  h3 {\\n    margin-top: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n    font-size: var(--pf-t--global--font--size--heading--lg, 1.25rem);\\n  }\\n\\n  dl {\\n    margin: 0;\\n  }\\n\\n  dt {\\n    font-weight: var(--pf-t--global--font--weight--body--bold, 600);\\n    margin-top: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  dd {\\n    margin-left: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  .empty-state {\\n    color: var(--pf-t--global--text--color--subtle, #6c757d);\\n    font-style: italic;\\n  }\\n}\\n"'));
var cem_manifest_browser_default = s;

// elements/cem-manifest-browser/cem-manifest-browser.ts
import "../cem-pf-v6-badge/cem-pf-v6-badge.js";
import "../cem-pf-v6-button/cem-pf-v6-button.js";
import "../cem-pf-v6-drawer/cem-pf-v6-drawer.js";
import "../cem-pf-v6-text-input-group/cem-pf-v6-text-input-group.js";
import "../cem-pf-v6-toolbar/cem-pf-v6-toolbar.js";
import "../cem-pf-v6-toolbar-group/cem-pf-v6-toolbar-group.js";
import "../cem-pf-v6-toolbar-item/cem-pf-v6-toolbar-item.js";
import "../cem-detail-panel/cem-detail-panel.js";
import "../cem-virtual-tree/cem-virtual-tree.js";
import { CemVirtualTree } from "../cem-virtual-tree/cem-virtual-tree.js";
var _CemManifestBrowser_decorators, _init, _a;
_CemManifestBrowser_decorators = [customElement("cem-manifest-browser")];
var CemManifestBrowser = class extends (_a = LitElement) {
  static styles = cem_manifest_browser_default;
  #searchDebounceTimer = null;
  disconnectedCallback() {
    super.disconnectedCallback();
    if (this.#searchDebounceTimer != null) {
      clearTimeout(this.#searchDebounceTimer);
      this.#searchDebounceTimer = null;
    }
  }
  render() {
    return html`
      <div id="drawer-content" slot="content">
        <cem-pf-v6-toolbar sticky>
          <cem-pf-v6-toolbar-group>
            <cem-pf-v6-toolbar-item>
              <cem-pf-v6-text-input-group id="search"
                                      type="search"
                                      placeholder="Search manifest..."
                                      aria-label="Search manifest"
                                      icon
                                      @input=${this.#onSearchInput}>
                <svg slot="icon"
                     role="presentation"
                     fill="currentColor"
                     height="1em"
                     width="1em"
                     viewBox="0 0 512 512">
                  <path d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                </svg>
                <cem-pf-v6-badge slot="utilities"
                             id="search-count"
                             compact
                             hidden>
                  0<span class="cem-pf-v6-screen-reader"> results</span>
                </cem-pf-v6-badge>
                <cem-pf-v6-button slot="utilities"
                              id="search-clear"
                              variant="plain"
                              aria-label="Clear search"
                              hidden
                              @click=${this.#onSearchClear}>
                  <svg width="1em"
                       height="1em"
                       viewBox="0 0 352 512"
                       fill="currentColor"
                       aria-hidden="true">
                    <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
                  </svg>
                </cem-pf-v6-button>
              </cem-pf-v6-text-input-group>
            </cem-pf-v6-toolbar-item>
          </cem-pf-v6-toolbar-group>
          <cem-pf-v6-toolbar-group variant="action-group">
            <cem-pf-v6-toolbar-item>
              <cem-pf-v6-button id="expand-all"
                            variant="tertiary"
                            size="small"
                            aria-label="Expand all tree items"
                            @click=${this.#onExpandAll}>
                Expand all
              </cem-pf-v6-button>
            </cem-pf-v6-toolbar-item>
            <cem-pf-v6-toolbar-item>
              <cem-pf-v6-button id="collapse-all"
                            variant="tertiary"
                            size="small"
                            aria-label="Collapse all tree items"
                            @click=${this.#onCollapseAll}>
                Collapse all
              </cem-pf-v6-button>
            </cem-pf-v6-toolbar-item>
          </cem-pf-v6-toolbar-group>
        </cem-pf-v6-toolbar>
        <cem-pf-v6-drawer id="drawer">
          <div id="tree-wrapper">
            <cem-virtual-tree id="virtual-tree"
                              @item-select=${this.#onItemSelect}></cem-virtual-tree>
          </div>
          <cem-detail-panel id="detail-panel"
                            slot="panel-body"></cem-detail-panel>
        </cem-pf-v6-drawer>
      </div>
    `;
  }
  get #drawer() {
    return this.shadowRoot?.getElementById("drawer");
  }
  get #virtualTree() {
    return this.shadowRoot?.getElementById("virtual-tree");
  }
  get #detailPanel() {
    return this.shadowRoot?.getElementById("detail-panel");
  }
  get #searchInput() {
    return this.shadowRoot?.getElementById("search");
  }
  get #searchCount() {
    return this.shadowRoot?.getElementById("search-count");
  }
  get #searchClear() {
    return this.shadowRoot?.getElementById("search-clear");
  }
  async #onItemSelect(e) {
    const item = e.item;
    if (!item) return;
    const manifest = await CemVirtualTree.loadManifest();
    if (this.#detailPanel && manifest) {
      await this.#detailPanel.renderItem(item, manifest);
      if (this.#drawer) {
        this.#drawer.expanded = true;
      }
    }
  }
  #onSearchInput() {
    const value = this.#searchInput?.value || "";
    if (this.#searchClear) {
      this.#searchClear.hidden = !value;
    }
    if (this.#searchDebounceTimer != null) {
      clearTimeout(this.#searchDebounceTimer);
    }
    this.#searchDebounceTimer = setTimeout(() => {
      this.#handleSearch(value);
    }, 300);
  }
  #onSearchClear() {
    if (this.#searchInput) {
      this.#searchInput.value = "";
    }
    if (this.#searchClear) {
      this.#searchClear.hidden = true;
    }
    if (this.#searchCount) {
      this.#searchCount.hidden = true;
    }
    this.#handleSearch("");
  }
  #handleSearch(query) {
    this.#virtualTree?.search(query);
    if (this.#searchCount) {
      this.#searchCount.hidden = !query;
    }
  }
  #onExpandAll() {
    this.#virtualTree?.expandAll();
  }
  #onCollapseAll() {
    this.#virtualTree?.collapseAll();
  }
};
_init = __decoratorStart(_a);
CemManifestBrowser = __decorateElement(_init, 0, "CemManifestBrowser", _CemManifestBrowser_decorators, CemManifestBrowser);
__runInitializers(_init, 1, CemManifestBrowser);
export {
  CemManifestBrowser
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLW1hbmlmZXN0LWJyb3dzZXIvY2VtLW1hbmlmZXN0LWJyb3dzZXIudHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tbWFuaWZlc3QtYnJvd3Nlci9jZW0tbWFuaWZlc3QtYnJvd3Nlci5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1tYW5pZmVzdC1icm93c2VyLmNzcyc7XG5cbmltcG9ydCAnLi4vY2VtLXBmLXY2LWJhZGdlL2NlbS1wZi12Ni1iYWRnZS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1idXR0b24vY2VtLXBmLXY2LWJ1dHRvbi5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1kcmF3ZXIvY2VtLXBmLXY2LWRyYXdlci5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10ZXh0LWlucHV0LWdyb3VwL2NlbS1wZi12Ni10ZXh0LWlucHV0LWdyb3VwLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRvb2xiYXIvY2VtLXBmLXY2LXRvb2xiYXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtdG9vbGJhci1ncm91cC9jZW0tcGYtdjYtdG9vbGJhci1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10b29sYmFyLWl0ZW0vY2VtLXBmLXY2LXRvb2xiYXItaXRlbS5qcyc7XG5cbi8vIGNlbS1kZXRhaWwtcGFuZWwgYW5kIGNlbS12aXJ0dWFsLXRyZWUgYXJlIHN0aWxsIENlbUVsZW1lbnQtYmFzZWQsXG4vLyBpbXBvcnRlZCBhcyBzaWRlLWVmZmVjdCBKUyBmcm9tIHRoZWlyIHRlbXBsYXRlIGxvY2F0aW9ucy5cbmltcG9ydCAnLi4vY2VtLWRldGFpbC1wYW5lbC9jZW0tZGV0YWlsLXBhbmVsLmpzJztcbmltcG9ydCAnLi4vY2VtLXZpcnR1YWwtdHJlZS9jZW0tdmlydHVhbC10cmVlLmpzJztcblxuLy8gUmUtZXhwb3J0IGZvciB0eXBlIHVzYWdlOyBDZW1WaXJ0dWFsVHJlZSBwcm92aWRlcyBsb2FkTWFuaWZlc3QoKVxuLy8gQHRzLWlnbm9yZSAtLSBDZW1FbGVtZW50LWJhc2VkIEpTLCBubyB0eXBlcyBhdmFpbGFibGVcbmltcG9ydCB7IENlbVZpcnR1YWxUcmVlIH0gZnJvbSAnLi4vY2VtLXZpcnR1YWwtdHJlZS9jZW0tdmlydHVhbC10cmVlLmpzJztcblxuaW50ZXJmYWNlIFRyZWVJdGVtIHtcbiAgaWQ6IG51bWJlcjtcbiAgdHlwZTogc3RyaW5nO1xuICBsYWJlbDogc3RyaW5nO1xuICBba2V5OiBzdHJpbmddOiB1bmtub3duO1xufVxuXG5pbnRlcmZhY2UgRGV0YWlsUGFuZWwgZXh0ZW5kcyBIVE1MRWxlbWVudCB7XG4gIHJlbmRlckl0ZW0oaXRlbTogVHJlZUl0ZW0sIG1hbmlmZXN0OiB1bmtub3duKTogUHJvbWlzZTx2b2lkPjtcbn1cblxuaW50ZXJmYWNlIFZpcnR1YWxUcmVlIGV4dGVuZHMgSFRNTEVsZW1lbnQge1xuICBzZWFyY2gocXVlcnk6IHN0cmluZyk6IHZvaWQ7XG4gIGV4cGFuZEFsbCgpOiB2b2lkO1xuICBjb2xsYXBzZUFsbCgpOiB2b2lkO1xufVxuXG5pbnRlcmZhY2UgSXRlbVNlbGVjdEV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBpdGVtOiBUcmVlSXRlbTtcbn1cblxuLyoqXG4gKiBNYW5pZmVzdCBCcm93c2VyIHdpdGggdHJlZSBuYXZpZ2F0aW9uIGFuZCBkZXRhaWwgZHJhd2VyXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgKHVudXNlZClcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1tYW5pZmVzdC1icm93c2VyJylcbmV4cG9ydCBjbGFzcyBDZW1NYW5pZmVzdEJyb3dzZXIgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjc2VhcmNoRGVib3VuY2VUaW1lcjogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIGlmICh0aGlzLiNzZWFyY2hEZWJvdW5jZVRpbWVyICE9IG51bGwpIHtcbiAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNzZWFyY2hEZWJvdW5jZVRpbWVyKTtcbiAgICAgIHRoaXMuI3NlYXJjaERlYm91bmNlVGltZXIgPSBudWxsO1xuICAgIH1cbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJkcmF3ZXItY29udGVudFwiIHNsb3Q9XCJjb250ZW50XCI+XG4gICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhciBzdGlja3k+XG4gICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgIDxjZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cCBpZD1cInNlYXJjaFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHR5cGU9XCJzZWFyY2hcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBwbGFjZWhvbGRlcj1cIlNlYXJjaCBtYW5pZmVzdC4uLlwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJTZWFyY2ggbWFuaWZlc3RcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpY29uXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIEBpbnB1dD0ke3RoaXMuI29uU2VhcmNoSW5wdXR9PlxuICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIlxuICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDUxMiA1MTJcIj5cbiAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNTA1IDQ0Mi43TDQwNS4zIDM0M2MtNC41LTQuNS0xMC42LTctMTctN0gzNzJjMjcuNi0zNS4zIDQ0LTc5LjcgNDQtMTI4QzQxNiA5My4xIDMyMi45IDAgMjA4IDBTMCA5My4xIDAgMjA4czkzLjEgMjA4IDIwOCAyMDhjNDguMyAwIDkyLjctMTYuNCAxMjgtNDR2MTYuM2MwIDYuNCAyLjUgMTIuNSA3IDE3bDk5LjcgOTkuN2M5LjQgOS40IDI0LjYgOS40IDMzLjkgMGwyOC4zLTI4LjNjOS40LTkuNCA5LjQtMjQuNi4xLTM0ek0yMDggMzM2Yy03MC43IDAtMTI4LTU3LjItMTI4LTEyOCAwLTcwLjcgNTcuMi0xMjggMTI4LTEyOCA3MC43IDAgMTI4IDU3LjIgMTI4IDEyOCAwIDcwLjctNTcuMiAxMjgtMTI4IDEyOHpcIj48L3BhdGg+XG4gICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgPGNlbS1wZi12Ni1iYWRnZSBzbG90PVwidXRpbGl0aWVzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaWQ9XCJzZWFyY2gtY291bnRcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICBjb21wYWN0XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhpZGRlbj5cbiAgICAgICAgICAgICAgICAgIDA8c3BhbiBjbGFzcz1cImNlbS1wZi12Ni1zY3JlZW4tcmVhZGVyXCI+IHJlc3VsdHM8L3NwYW4+XG4gICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtYmFkZ2U+XG4gICAgICAgICAgICAgICAgPGNlbS1wZi12Ni1idXR0b24gc2xvdD1cInV0aWxpdGllc1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpZD1cInNlYXJjaC1jbGVhclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwicGxhaW5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkNsZWFyIHNlYXJjaFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICBoaWRkZW5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIEBjbGljaz0ke3RoaXMuI29uU2VhcmNoQ2xlYXJ9PlxuICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMzUyIDUxMlwiXG4gICAgICAgICAgICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWhpZGRlbj1cInRydWVcIj5cbiAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0yNDIuNzIgMjU2bDEwMC4wNy0xMDAuMDdjMTIuMjgtMTIuMjggMTIuMjgtMzIuMTkgMC00NC40OGwtMjIuMjQtMjIuMjRjLTEyLjI4LTEyLjI4LTMyLjE5LTEyLjI4LTQ0LjQ4IDBMMTc2IDE4OS4yOCA3NS45MyA4OS4yMWMtMTIuMjgtMTIuMjgtMzIuMTktMTIuMjgtNDQuNDggMEw5LjIxIDExMS40NWMtMTIuMjggMTIuMjgtMTIuMjggMzIuMTkgMCA0NC40OEwxMDkuMjggMjU2IDkuMjEgMzU2LjA3Yy0xMi4yOCAxMi4yOC0xMi4yOCAzMi4xOSAwIDQ0LjQ4bDIyLjI0IDIyLjI0YzEyLjI4IDEyLjI4IDMyLjIgMTIuMjggNDQuNDggMEwxNzYgMzIyLjcybDEwMC4wNyAxMDAuMDdjMTIuMjggMTIuMjggMzIuMiAxMi4yOCA0NC40OCAwbDIyLjI0LTIyLjI0YzEyLjI4LTEyLjI4IDEyLjI4LTMyLjE5IDAtNDQuNDhMMjQyLjcyIDI1NnpcIj48L3BhdGg+XG4gICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni1idXR0b24+XG4gICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRleHQtaW5wdXQtZ3JvdXA+XG4gICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItZ3JvdXAgdmFyaWFudD1cImFjdGlvbi1ncm91cFwiPlxuICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgIDxjZW0tcGYtdjYtYnV0dG9uIGlkPVwiZXhwYW5kLWFsbFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInRlcnRpYXJ5XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBzaXplPVwic21hbGxcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJFeHBhbmQgYWxsIHRyZWUgaXRlbXNcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIEBjbGljaz0ke3RoaXMuI29uRXhwYW5kQWxsfT5cbiAgICAgICAgICAgICAgICBFeHBhbmQgYWxsXG4gICAgICAgICAgICAgIDwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBpZD1cImNvbGxhcHNlLWFsbFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInRlcnRpYXJ5XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBzaXplPVwic21hbGxcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJDb2xsYXBzZSBhbGwgdHJlZSBpdGVtc1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25Db2xsYXBzZUFsbH0+XG4gICAgICAgICAgICAgICAgQ29sbGFwc2UgYWxsXG4gICAgICAgICAgICAgIDwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyPlxuICAgICAgICA8Y2VtLXBmLXY2LWRyYXdlciBpZD1cImRyYXdlclwiPlxuICAgICAgICAgIDxkaXYgaWQ9XCJ0cmVlLXdyYXBwZXJcIj5cbiAgICAgICAgICAgIDxjZW0tdmlydHVhbC10cmVlIGlkPVwidmlydHVhbC10cmVlXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIEBpdGVtLXNlbGVjdD0ke3RoaXMuI29uSXRlbVNlbGVjdH0+PC9jZW0tdmlydHVhbC10cmVlPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxjZW0tZGV0YWlsLXBhbmVsIGlkPVwiZGV0YWlsLXBhbmVsXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBzbG90PVwicGFuZWwtYm9keVwiPjwvY2VtLWRldGFpbC1wYW5lbD5cbiAgICAgICAgPC9jZW0tcGYtdjYtZHJhd2VyPlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gIGdldCAjZHJhd2VyKCkge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdkcmF3ZXInKSBhcyBIVE1MRWxlbWVudCAmIHsgZXhwYW5kZWQ6IGJvb2xlYW4gfSB8IG51bGw7XG4gIH1cblxuICBnZXQgI3ZpcnR1YWxUcmVlKCk6IFZpcnR1YWxUcmVlIHwgbnVsbCB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ3ZpcnR1YWwtdHJlZScpIGFzIFZpcnR1YWxUcmVlIHwgbnVsbDtcbiAgfVxuXG4gIGdldCAjZGV0YWlsUGFuZWwoKTogRGV0YWlsUGFuZWwgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnZGV0YWlsLXBhbmVsJykgYXMgRGV0YWlsUGFuZWwgfCBudWxsO1xuICB9XG5cbiAgZ2V0ICNzZWFyY2hJbnB1dCgpIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnc2VhcmNoJykgYXMgSFRNTEVsZW1lbnQgJiB7IHZhbHVlOiBzdHJpbmcgfSB8IG51bGw7XG4gIH1cblxuICBnZXQgI3NlYXJjaENvdW50KCkge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdzZWFyY2gtY291bnQnKSBhcyBIVE1MRWxlbWVudCB8IG51bGw7XG4gIH1cblxuICBnZXQgI3NlYXJjaENsZWFyKCkge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdzZWFyY2gtY2xlYXInKSBhcyBIVE1MRWxlbWVudCB8IG51bGw7XG4gIH1cblxuICBhc3luYyAjb25JdGVtU2VsZWN0KGU6IEl0ZW1TZWxlY3RFdmVudCkge1xuICAgIGNvbnN0IGl0ZW0gPSBlLml0ZW07XG4gICAgaWYgKCFpdGVtKSByZXR1cm47XG5cbiAgICBjb25zdCBtYW5pZmVzdCA9IGF3YWl0IENlbVZpcnR1YWxUcmVlLmxvYWRNYW5pZmVzdCgpO1xuICAgIGlmICh0aGlzLiNkZXRhaWxQYW5lbCAmJiBtYW5pZmVzdCkge1xuICAgICAgYXdhaXQgdGhpcy4jZGV0YWlsUGFuZWwucmVuZGVySXRlbShpdGVtLCBtYW5pZmVzdCk7XG4gICAgICBpZiAodGhpcy4jZHJhd2VyKSB7XG4gICAgICAgIHRoaXMuI2RyYXdlci5leHBhbmRlZCA9IHRydWU7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI29uU2VhcmNoSW5wdXQoKSB7XG4gICAgY29uc3QgdmFsdWUgPSB0aGlzLiNzZWFyY2hJbnB1dD8udmFsdWUgfHwgJyc7XG5cbiAgICAvLyBTaG93L2hpZGUgY2xlYXIgYnV0dG9uIGJhc2VkIG9uIHdoZXRoZXIgdGhlcmUncyBhIHZhbHVlXG4gICAgaWYgKHRoaXMuI3NlYXJjaENsZWFyKSB7XG4gICAgICB0aGlzLiNzZWFyY2hDbGVhci5oaWRkZW4gPSAhdmFsdWU7XG4gICAgfVxuXG4gICAgLy8gRGVib3VuY2Ugc2VhcmNoIC0gd2FpdCAzMDBtcyBhZnRlciB1c2VyIHN0b3BzIHR5cGluZ1xuICAgIGlmICh0aGlzLiNzZWFyY2hEZWJvdW5jZVRpbWVyICE9IG51bGwpIHtcbiAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNzZWFyY2hEZWJvdW5jZVRpbWVyKTtcbiAgICB9XG4gICAgdGhpcy4jc2VhcmNoRGVib3VuY2VUaW1lciA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgdGhpcy4jaGFuZGxlU2VhcmNoKHZhbHVlKTtcbiAgICB9LCAzMDApO1xuICB9XG5cbiAgI29uU2VhcmNoQ2xlYXIoKSB7XG4gICAgaWYgKHRoaXMuI3NlYXJjaElucHV0KSB7XG4gICAgICB0aGlzLiNzZWFyY2hJbnB1dC52YWx1ZSA9ICcnO1xuICAgIH1cbiAgICBpZiAodGhpcy4jc2VhcmNoQ2xlYXIpIHtcbiAgICAgIHRoaXMuI3NlYXJjaENsZWFyLmhpZGRlbiA9IHRydWU7XG4gICAgfVxuICAgIGlmICh0aGlzLiNzZWFyY2hDb3VudCkge1xuICAgICAgdGhpcy4jc2VhcmNoQ291bnQuaGlkZGVuID0gdHJ1ZTtcbiAgICB9XG4gICAgdGhpcy4jaGFuZGxlU2VhcmNoKCcnKTtcbiAgfVxuXG4gICNoYW5kbGVTZWFyY2gocXVlcnk6IHN0cmluZykge1xuICAgIHRoaXMuI3ZpcnR1YWxUcmVlPy5zZWFyY2gocXVlcnkpO1xuICAgIGlmICh0aGlzLiNzZWFyY2hDb3VudCkge1xuICAgICAgdGhpcy4jc2VhcmNoQ291bnQuaGlkZGVuID0gIXF1ZXJ5O1xuICAgIH1cbiAgfVxuXG4gICNvbkV4cGFuZEFsbCgpIHtcbiAgICB0aGlzLiN2aXJ0dWFsVHJlZT8uZXhwYW5kQWxsKCk7XG4gIH1cblxuICAjb25Db2xsYXBzZUFsbCgpIHtcbiAgICB0aGlzLiN2aXJ0dWFsVHJlZT8uY29sbGFwc2VBbGwoKTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tbWFuaWZlc3QtYnJvd3Nlcic6IENlbU1hbmlmZXN0QnJvd3NlcjtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiLyogTWFuaWZlc3QgQnJvd3NlciAtIENvbXBvbmVudCBBUEkgbmF2aWdhdGlvbiB0cmVlICovXFxcXG5cXFxcbjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxuW2hpZGRlbl0ge1xcXFxuICBkaXNwbGF5OiBub25lICFpbXBvcnRhbnQ7XFxcXG59XFxcXG5cXFxcbiNkcmF3ZXIge1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbiNkcmF3ZXItY29udGVudCB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbn1cXFxcblxcXFxuI3RyZWUtd3JhcHBlciB7XFxcXG4gIGZsZXg6IDE7XFxcXG4gIG92ZXJmbG93OiBhdXRvO1xcXFxuICBtaW4taGVpZ2h0OiAwO1xcXFxufVxcXFxuXFxcXG46OnNsb3R0ZWQoY2VtLXBmLXY2LXRyZWUtdmlldykge1xcXFxuICBtYXJnaW46IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCwgMXJlbSk7XFxcXG59XFxcXG5cXFxcbiNwYW5lbC10aXRsZSB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbn1cXFxcblxcXFxuI2RldGFpbHMtY29udGVudCB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCwgMXJlbSk7XFxcXG5cXFxcbiAgaDMge1xcXFxuICAgIG1hcmdpbi10b3A6IDA7XFxcXG4gICAgbWFyZ2luLWJvdHRvbTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtLCAwLjVyZW0pO1xcXFxuICAgIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1oZWFkaW5nLS1sZywgMS4yNXJlbSk7XFxcXG4gIH1cXFxcblxcXFxuICBkbCB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICB9XFxcXG5cXFxcbiAgZHQge1xcXFxuICAgIGZvbnQtd2VpZ2h0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXdlaWdodC0tYm9keS0tYm9sZCwgNjAwKTtcXFxcbiAgICBtYXJnaW4tdG9wOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20sIDAuNXJlbSk7XFxcXG4gIH1cXFxcblxcXFxuICBkZCB7XFxcXG4gICAgbWFyZ2luLWxlZnQ6IDA7XFxcXG4gICAgbWFyZ2luLWJvdHRvbTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtLCAwLjVyZW0pO1xcXFxuICB9XFxcXG5cXFxcbiAgLmVtcHR5LXN0YXRlIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlLCAjNmM3NTdkKTtcXFxcbiAgICBmb250LXN0eWxlOiBpdGFsaWM7XFxcXG4gIH1cXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7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
