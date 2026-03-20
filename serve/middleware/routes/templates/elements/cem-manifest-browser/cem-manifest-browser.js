var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __knownSymbol = (name, symbol) => (symbol = Symbol[name]) ? symbol : Symbol.for("Symbol." + name);
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

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-manifest-browser/cem-manifest-browser.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('"/* Manifest Browser - Component API navigation tree */\\n\\n:host {\\n  display: block;\\n  height: 100%;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n#drawer {\\n  height: 100%;\\n}\\n\\n#drawer-content {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n  overflow: hidden;\\n}\\n\\n#tree-wrapper {\\n  flex: 1;\\n  overflow: auto;\\n  min-height: 0;\\n}\\n\\n::slotted(pf-v6-tree-view) {\\n  margin: var(--pf-t--global--spacer--md, 1rem);\\n}\\n\\n#panel-title {\\n  margin: 0;\\n}\\n\\n#details-content {\\n  padding: var(--pf-t--global--spacer--md, 1rem);\\n\\n  h3 {\\n    margin-top: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n    font-size: var(--pf-t--global--font--size--heading--lg, 1.25rem);\\n  }\\n\\n  dl {\\n    margin: 0;\\n  }\\n\\n  dt {\\n    font-weight: var(--pf-t--global--font--weight--body--bold, 600);\\n    margin-top: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  dd {\\n    margin-left: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  .empty-state {\\n    color: var(--pf-t--global--text--color--subtle, #6c757d);\\n    font-style: italic;\\n  }\\n}\\n"'));
var cem_manifest_browser_default = s;

// elements/cem-manifest-browser/cem-manifest-browser.ts
import "../pf-v6-badge/pf-v6-badge.js";
import "../pf-v6-button/pf-v6-button.js";
import "../pf-v6-drawer/pf-v6-drawer.js";
import "../pf-v6-text-input-group/pf-v6-text-input-group.js";
import "../pf-v6-toolbar/pf-v6-toolbar.js";
import "../pf-v6-toolbar-group/pf-v6-toolbar-group.js";
import "../pf-v6-toolbar-item/pf-v6-toolbar-item.js";
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
        <pf-v6-toolbar sticky>
          <pf-v6-toolbar-group>
            <pf-v6-toolbar-item>
              <pf-v6-text-input-group id="search"
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
                <pf-v6-badge slot="utilities"
                             id="search-count"
                             compact
                             hidden>
                  0<span class="pf-v6-screen-reader"> results</span>
                </pf-v6-badge>
                <pf-v6-button slot="utilities"
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
                </pf-v6-button>
              </pf-v6-text-input-group>
            </pf-v6-toolbar-item>
          </pf-v6-toolbar-group>
          <pf-v6-toolbar-group variant="action-group">
            <pf-v6-toolbar-item>
              <pf-v6-button id="expand-all"
                            variant="tertiary"
                            size="small"
                            aria-label="Expand all tree items"
                            @click=${this.#onExpandAll}>
                Expand all
              </pf-v6-button>
            </pf-v6-toolbar-item>
            <pf-v6-toolbar-item>
              <pf-v6-button id="collapse-all"
                            variant="tertiary"
                            size="small"
                            aria-label="Collapse all tree items"
                            @click=${this.#onCollapseAll}>
                Collapse all
              </pf-v6-button>
            </pf-v6-toolbar-item>
          </pf-v6-toolbar-group>
        </pf-v6-toolbar>
        <pf-v6-drawer id="drawer">
          <div id="tree-wrapper">
            <cem-virtual-tree id="virtual-tree"
                              @item-select=${this.#onItemSelect}></cem-virtual-tree>
          </div>
          <cem-detail-panel id="detail-panel"
                            slot="panel-body"></cem-detail-panel>
        </pf-v6-drawer>
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLW1hbmlmZXN0LWJyb3dzZXIvY2VtLW1hbmlmZXN0LWJyb3dzZXIudHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9jZW0tbWFuaWZlc3QtYnJvd3Nlci9jZW0tbWFuaWZlc3QtYnJvd3Nlci5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1tYW5pZmVzdC1icm93c2VyLmNzcyc7XG5cbmltcG9ydCAnLi4vcGYtdjYtYmFkZ2UvcGYtdjYtYmFkZ2UuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtZHJhd2VyL3BmLXY2LWRyYXdlci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRleHQtaW5wdXQtZ3JvdXAvcGYtdjYtdGV4dC1pbnB1dC1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvb2xiYXIvcGYtdjYtdG9vbGJhci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvb2xiYXItZ3JvdXAvcGYtdjYtdG9vbGJhci1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvb2xiYXItaXRlbS9wZi12Ni10b29sYmFyLWl0ZW0uanMnO1xuXG4vLyBjZW0tZGV0YWlsLXBhbmVsIGFuZCBjZW0tdmlydHVhbC10cmVlIGFyZSBzdGlsbCBDZW1FbGVtZW50LWJhc2VkLFxuLy8gaW1wb3J0ZWQgYXMgc2lkZS1lZmZlY3QgSlMgZnJvbSB0aGVpciB0ZW1wbGF0ZSBsb2NhdGlvbnMuXG5pbXBvcnQgJy4uL2NlbS1kZXRhaWwtcGFuZWwvY2VtLWRldGFpbC1wYW5lbC5qcyc7XG5pbXBvcnQgJy4uL2NlbS12aXJ0dWFsLXRyZWUvY2VtLXZpcnR1YWwtdHJlZS5qcyc7XG5cbi8vIFJlLWV4cG9ydCBmb3IgdHlwZSB1c2FnZTsgQ2VtVmlydHVhbFRyZWUgcHJvdmlkZXMgbG9hZE1hbmlmZXN0KClcbi8vIEB0cy1pZ25vcmUgLS0gQ2VtRWxlbWVudC1iYXNlZCBKUywgbm8gdHlwZXMgYXZhaWxhYmxlXG5pbXBvcnQgeyBDZW1WaXJ0dWFsVHJlZSB9IGZyb20gJy4uL2NlbS12aXJ0dWFsLXRyZWUvY2VtLXZpcnR1YWwtdHJlZS5qcyc7XG5cbmludGVyZmFjZSBUcmVlSXRlbSB7XG4gIGlkOiBudW1iZXI7XG4gIHR5cGU6IHN0cmluZztcbiAgbGFiZWw6IHN0cmluZztcbiAgW2tleTogc3RyaW5nXTogdW5rbm93bjtcbn1cblxuaW50ZXJmYWNlIERldGFpbFBhbmVsIGV4dGVuZHMgSFRNTEVsZW1lbnQge1xuICByZW5kZXJJdGVtKGl0ZW06IFRyZWVJdGVtLCBtYW5pZmVzdDogdW5rbm93bik6IFByb21pc2U8dm9pZD47XG59XG5cbmludGVyZmFjZSBWaXJ0dWFsVHJlZSBleHRlbmRzIEhUTUxFbGVtZW50IHtcbiAgc2VhcmNoKHF1ZXJ5OiBzdHJpbmcpOiB2b2lkO1xuICBleHBhbmRBbGwoKTogdm9pZDtcbiAgY29sbGFwc2VBbGwoKTogdm9pZDtcbn1cblxuaW50ZXJmYWNlIEl0ZW1TZWxlY3RFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgaXRlbTogVHJlZUl0ZW07XG59XG5cbi8qKlxuICogTWFuaWZlc3QgQnJvd3NlciB3aXRoIHRyZWUgbmF2aWdhdGlvbiBhbmQgZGV0YWlsIGRyYXdlclxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90ICh1bnVzZWQpXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tbWFuaWZlc3QtYnJvd3NlcicpXG5leHBvcnQgY2xhc3MgQ2VtTWFuaWZlc3RCcm93c2VyIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgI3NlYXJjaERlYm91bmNlVGltZXI6IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+IHwgbnVsbCA9IG51bGw7XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICBpZiAodGhpcy4jc2VhcmNoRGVib3VuY2VUaW1lciAhPSBudWxsKSB7XG4gICAgICBjbGVhclRpbWVvdXQodGhpcy4jc2VhcmNoRGVib3VuY2VUaW1lcik7XG4gICAgICB0aGlzLiNzZWFyY2hEZWJvdW5jZVRpbWVyID0gbnVsbDtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8ZGl2IGlkPVwiZHJhd2VyLWNvbnRlbnRcIiBzbG90PVwiY29udGVudFwiPlxuICAgICAgICA8cGYtdjYtdG9vbGJhciBzdGlja3k+XG4gICAgICAgICAgPHBmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICA8cGYtdjYtdGV4dC1pbnB1dC1ncm91cCBpZD1cInNlYXJjaFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHR5cGU9XCJzZWFyY2hcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBwbGFjZWhvbGRlcj1cIlNlYXJjaCBtYW5pZmVzdC4uLlwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJTZWFyY2ggbWFuaWZlc3RcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpY29uXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIEBpbnB1dD0ke3RoaXMuI29uU2VhcmNoSW5wdXR9PlxuICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIlxuICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDUxMiA1MTJcIj5cbiAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNTA1IDQ0Mi43TDQwNS4zIDM0M2MtNC41LTQuNS0xMC42LTctMTctN0gzNzJjMjcuNi0zNS4zIDQ0LTc5LjcgNDQtMTI4QzQxNiA5My4xIDMyMi45IDAgMjA4IDBTMCA5My4xIDAgMjA4czkzLjEgMjA4IDIwOCAyMDhjNDguMyAwIDkyLjctMTYuNCAxMjgtNDR2MTYuM2MwIDYuNCAyLjUgMTIuNSA3IDE3bDk5LjcgOTkuN2M5LjQgOS40IDI0LjYgOS40IDMzLjkgMGwyOC4zLTI4LjNjOS40LTkuNCA5LjQtMjQuNi4xLTM0ek0yMDggMzM2Yy03MC43IDAtMTI4LTU3LjItMTI4LTEyOCAwLTcwLjcgNTcuMi0xMjggMTI4LTEyOCA3MC43IDAgMTI4IDU3LjIgMTI4IDEyOCAwIDcwLjctNTcuMiAxMjgtMTI4IDEyOHpcIj48L3BhdGg+XG4gICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgPHBmLXY2LWJhZGdlIHNsb3Q9XCJ1dGlsaXRpZXNcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpZD1cInNlYXJjaC1jb3VudFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgIGNvbXBhY3RcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaGlkZGVuPlxuICAgICAgICAgICAgICAgICAgMDxzcGFuIGNsYXNzPVwicGYtdjYtc2NyZWVuLXJlYWRlclwiPiByZXN1bHRzPC9zcGFuPlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtYmFkZ2U+XG4gICAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBzbG90PVwidXRpbGl0aWVzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGlkPVwic2VhcmNoLWNsZWFyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiQ2xlYXIgc2VhcmNoXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhpZGRlblxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25TZWFyY2hDbGVhcn0+XG4gICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMWVtXCJcbiAgICAgICAgICAgICAgICAgICAgICAgaGVpZ2h0PVwiMWVtXCJcbiAgICAgICAgICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCAzNTIgNTEyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICAgICAgICAgIGFyaWEtaGlkZGVuPVwidHJ1ZVwiPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTI0Mi43MiAyNTZsMTAwLjA3LTEwMC4wN2MxMi4yOC0xMi4yOCAxMi4yOC0zMi4xOSAwLTQ0LjQ4bC0yMi4yNC0yMi4yNGMtMTIuMjgtMTIuMjgtMzIuMTktMTIuMjgtNDQuNDggMEwxNzYgMTg5LjI4IDc1LjkzIDg5LjIxYy0xMi4yOC0xMi4yOC0zMi4xOS0xMi4yOC00NC40OCAwTDkuMjEgMTExLjQ1Yy0xMi4yOCAxMi4yOC0xMi4yOCAzMi4xOSAwIDQ0LjQ4TDEwOS4yOCAyNTYgOS4yMSAzNTYuMDdjLTEyLjI4IDEyLjI4LTEyLjI4IDMyLjE5IDAgNDQuNDhsMjIuMjQgMjIuMjRjMTIuMjggMTIuMjggMzIuMiAxMi4yOCA0NC40OCAwTDE3NiAzMjIuNzJsMTAwLjA3IDEwMC4wN2MxMi4yOCAxMi4yOCAzMi4yIDEyLjI4IDQ0LjQ4IDBsMjIuMjQtMjIuMjRjMTIuMjgtMTIuMjggMTIuMjgtMzIuMTkgMC00NC40OEwyNDIuNzIgMjU2elwiPjwvcGF0aD5cbiAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtYnV0dG9uPlxuICAgICAgICAgICAgICA8L3BmLXY2LXRleHQtaW5wdXQtZ3JvdXA+XG4gICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgPHBmLXY2LXRvb2xiYXItZ3JvdXAgdmFyaWFudD1cImFjdGlvbi1ncm91cFwiPlxuICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImV4cGFuZC1hbGxcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJ0ZXJ0aWFyeVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgc2l6ZT1cInNtYWxsXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiRXhwYW5kIGFsbCB0cmVlIGl0ZW1zXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLiNvbkV4cGFuZEFsbH0+XG4gICAgICAgICAgICAgICAgRXhwYW5kIGFsbFxuICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImNvbGxhcHNlLWFsbFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInRlcnRpYXJ5XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBzaXplPVwic21hbGxcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJDb2xsYXBzZSBhbGwgdHJlZSBpdGVtc1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25Db2xsYXBzZUFsbH0+XG4gICAgICAgICAgICAgICAgQ29sbGFwc2UgYWxsXG4gICAgICAgICAgICAgIDwvcGYtdjYtYnV0dG9uPlxuICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgPC9wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICA8L3BmLXY2LXRvb2xiYXI+XG4gICAgICAgIDxwZi12Ni1kcmF3ZXIgaWQ9XCJkcmF3ZXJcIj5cbiAgICAgICAgICA8ZGl2IGlkPVwidHJlZS13cmFwcGVyXCI+XG4gICAgICAgICAgICA8Y2VtLXZpcnR1YWwtdHJlZSBpZD1cInZpcnR1YWwtdHJlZVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICBAaXRlbS1zZWxlY3Q9JHt0aGlzLiNvbkl0ZW1TZWxlY3R9PjwvY2VtLXZpcnR1YWwtdHJlZT5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8Y2VtLWRldGFpbC1wYW5lbCBpZD1cImRldGFpbC1wYW5lbFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgc2xvdD1cInBhbmVsLWJvZHlcIj48L2NlbS1kZXRhaWwtcGFuZWw+XG4gICAgICAgIDwvcGYtdjYtZHJhd2VyPlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gIGdldCAjZHJhd2VyKCkge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdkcmF3ZXInKSBhcyBIVE1MRWxlbWVudCAmIHsgZXhwYW5kZWQ6IGJvb2xlYW4gfSB8IG51bGw7XG4gIH1cblxuICBnZXQgI3ZpcnR1YWxUcmVlKCk6IFZpcnR1YWxUcmVlIHwgbnVsbCB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ3ZpcnR1YWwtdHJlZScpIGFzIFZpcnR1YWxUcmVlIHwgbnVsbDtcbiAgfVxuXG4gIGdldCAjZGV0YWlsUGFuZWwoKTogRGV0YWlsUGFuZWwgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnZGV0YWlsLXBhbmVsJykgYXMgRGV0YWlsUGFuZWwgfCBudWxsO1xuICB9XG5cbiAgZ2V0ICNzZWFyY2hJbnB1dCgpIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnc2VhcmNoJykgYXMgSFRNTEVsZW1lbnQgJiB7IHZhbHVlOiBzdHJpbmcgfSB8IG51bGw7XG4gIH1cblxuICBnZXQgI3NlYXJjaENvdW50KCkge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdzZWFyY2gtY291bnQnKSBhcyBIVE1MRWxlbWVudCB8IG51bGw7XG4gIH1cblxuICBnZXQgI3NlYXJjaENsZWFyKCkge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdzZWFyY2gtY2xlYXInKSBhcyBIVE1MRWxlbWVudCB8IG51bGw7XG4gIH1cblxuICBhc3luYyAjb25JdGVtU2VsZWN0KGU6IEl0ZW1TZWxlY3RFdmVudCkge1xuICAgIGNvbnN0IGl0ZW0gPSBlLml0ZW07XG4gICAgaWYgKCFpdGVtKSByZXR1cm47XG5cbiAgICBjb25zdCBtYW5pZmVzdCA9IGF3YWl0IENlbVZpcnR1YWxUcmVlLmxvYWRNYW5pZmVzdCgpO1xuICAgIGlmICh0aGlzLiNkZXRhaWxQYW5lbCAmJiBtYW5pZmVzdCkge1xuICAgICAgYXdhaXQgdGhpcy4jZGV0YWlsUGFuZWwucmVuZGVySXRlbShpdGVtLCBtYW5pZmVzdCk7XG4gICAgICBpZiAodGhpcy4jZHJhd2VyKSB7XG4gICAgICAgIHRoaXMuI2RyYXdlci5leHBhbmRlZCA9IHRydWU7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI29uU2VhcmNoSW5wdXQoKSB7XG4gICAgY29uc3QgdmFsdWUgPSB0aGlzLiNzZWFyY2hJbnB1dD8udmFsdWUgfHwgJyc7XG5cbiAgICAvLyBTaG93L2hpZGUgY2xlYXIgYnV0dG9uIGJhc2VkIG9uIHdoZXRoZXIgdGhlcmUncyBhIHZhbHVlXG4gICAgaWYgKHRoaXMuI3NlYXJjaENsZWFyKSB7XG4gICAgICB0aGlzLiNzZWFyY2hDbGVhci5oaWRkZW4gPSAhdmFsdWU7XG4gICAgfVxuXG4gICAgLy8gRGVib3VuY2Ugc2VhcmNoIC0gd2FpdCAzMDBtcyBhZnRlciB1c2VyIHN0b3BzIHR5cGluZ1xuICAgIGlmICh0aGlzLiNzZWFyY2hEZWJvdW5jZVRpbWVyICE9IG51bGwpIHtcbiAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNzZWFyY2hEZWJvdW5jZVRpbWVyKTtcbiAgICB9XG4gICAgdGhpcy4jc2VhcmNoRGVib3VuY2VUaW1lciA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgdGhpcy4jaGFuZGxlU2VhcmNoKHZhbHVlKTtcbiAgICB9LCAzMDApO1xuICB9XG5cbiAgI29uU2VhcmNoQ2xlYXIoKSB7XG4gICAgaWYgKHRoaXMuI3NlYXJjaElucHV0KSB7XG4gICAgICB0aGlzLiNzZWFyY2hJbnB1dC52YWx1ZSA9ICcnO1xuICAgIH1cbiAgICBpZiAodGhpcy4jc2VhcmNoQ2xlYXIpIHtcbiAgICAgIHRoaXMuI3NlYXJjaENsZWFyLmhpZGRlbiA9IHRydWU7XG4gICAgfVxuICAgIGlmICh0aGlzLiNzZWFyY2hDb3VudCkge1xuICAgICAgdGhpcy4jc2VhcmNoQ291bnQuaGlkZGVuID0gdHJ1ZTtcbiAgICB9XG4gICAgdGhpcy4jaGFuZGxlU2VhcmNoKCcnKTtcbiAgfVxuXG4gICNoYW5kbGVTZWFyY2gocXVlcnk6IHN0cmluZykge1xuICAgIHRoaXMuI3ZpcnR1YWxUcmVlPy5zZWFyY2gocXVlcnkpO1xuICAgIGlmICh0aGlzLiNzZWFyY2hDb3VudCkge1xuICAgICAgdGhpcy4jc2VhcmNoQ291bnQuaGlkZGVuID0gIXF1ZXJ5O1xuICAgIH1cbiAgfVxuXG4gICNvbkV4cGFuZEFsbCgpIHtcbiAgICB0aGlzLiN2aXJ0dWFsVHJlZT8uZXhwYW5kQWxsKCk7XG4gIH1cblxuICAjb25Db2xsYXBzZUFsbCgpIHtcbiAgICB0aGlzLiN2aXJ0dWFsVHJlZT8uY29sbGFwc2VBbGwoKTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tbWFuaWZlc3QtYnJvd3Nlcic6IENlbU1hbmlmZXN0QnJvd3NlcjtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiLyogTWFuaWZlc3QgQnJvd3NlciAtIENvbXBvbmVudCBBUEkgbmF2aWdhdGlvbiB0cmVlICovXFxcXG5cXFxcbjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxuW2hpZGRlbl0ge1xcXFxuICBkaXNwbGF5OiBub25lICFpbXBvcnRhbnQ7XFxcXG59XFxcXG5cXFxcbiNkcmF3ZXIge1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbiNkcmF3ZXItY29udGVudCB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbn1cXFxcblxcXFxuI3RyZWUtd3JhcHBlciB7XFxcXG4gIGZsZXg6IDE7XFxcXG4gIG92ZXJmbG93OiBhdXRvO1xcXFxuICBtaW4taGVpZ2h0OiAwO1xcXFxufVxcXFxuXFxcXG46OnNsb3R0ZWQocGYtdjYtdHJlZS12aWV3KSB7XFxcXG4gIG1hcmdpbjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kLCAxcmVtKTtcXFxcbn1cXFxcblxcXFxuI3BhbmVsLXRpdGxlIHtcXFxcbiAgbWFyZ2luOiAwO1xcXFxufVxcXFxuXFxcXG4jZGV0YWlscy1jb250ZW50IHtcXFxcbiAgcGFkZGluZzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kLCAxcmVtKTtcXFxcblxcXFxuICBoMyB7XFxcXG4gICAgbWFyZ2luLXRvcDogMDtcXFxcbiAgICBtYXJnaW4tYm90dG9tOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20sIDAuNXJlbSk7XFxcXG4gICAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWhlYWRpbmctLWxnLCAxLjI1cmVtKTtcXFxcbiAgfVxcXFxuXFxcXG4gIGRsIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gIH1cXFxcblxcXFxuICBkdCB7XFxcXG4gICAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0td2VpZ2h0LS1ib2R5LS1ib2xkLCA2MDApO1xcXFxuICAgIG1hcmdpbi10b3A6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSwgMC41cmVtKTtcXFxcbiAgfVxcXFxuXFxcXG4gIGRkIHtcXFxcbiAgICBtYXJnaW4tbGVmdDogMDtcXFxcbiAgICBtYXJnaW4tYm90dG9tOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20sIDAuNXJlbSk7XFxcXG4gIH1cXFxcblxcXFxuICAuZW1wdHktc3RhdGUge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUsICM2Yzc1N2QpO1xcXFxuICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgfVxcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7OztBQ0Q5QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSwwcENBQTRwQyxDQUFDO0FBQUUs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
