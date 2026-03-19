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
s.replaceSync(JSON.parse('"/* Manifest Browser - Component API navigation tree */\\n\\n:host {\\n  display: block;\\n  height: 100%;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n#drawer {\\n  height: 100%;\\n}\\n\\n#drawer-content {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n  overflow: hidden;\\n}\\n\\n#tree-wrapper {\\n  flex: 1;\\n  overflow: auto;\\n  min-height: 0;\\n\\n  \\u0026::slotted(pf-v6-tree-view) {\\n    margin: var(--pf-t--global--spacer--md, 1rem);\\n  }\\n}\\n\\n#panel-title {\\n  margin: 0;\\n}\\n\\n#details-content {\\n  padding: var(--pf-t--global--spacer--md, 1rem);\\n\\n  h3 {\\n    margin-top: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n    font-size: var(--pf-t--global--font--size--heading--lg, 1.25rem);\\n  }\\n\\n  dl {\\n    margin: 0;\\n  }\\n\\n  dt {\\n    font-weight: var(--pf-t--global--font--weight--body--bold, 600);\\n    margin-top: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  dd {\\n    margin-left: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  .empty-state {\\n    color: var(--pf-t--global--text--color--subtle, #6c757d);\\n    font-style: italic;\\n  }\\n}\\n"'));
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLW1hbmlmZXN0LWJyb3dzZXIvY2VtLW1hbmlmZXN0LWJyb3dzZXIudHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9jZW0tbWFuaWZlc3QtYnJvd3Nlci9jZW0tbWFuaWZlc3QtYnJvd3Nlci5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1tYW5pZmVzdC1icm93c2VyLmNzcyc7XG5cbmltcG9ydCAnLi4vcGYtdjYtYmFkZ2UvcGYtdjYtYmFkZ2UuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtZHJhd2VyL3BmLXY2LWRyYXdlci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRleHQtaW5wdXQtZ3JvdXAvcGYtdjYtdGV4dC1pbnB1dC1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvb2xiYXIvcGYtdjYtdG9vbGJhci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvb2xiYXItZ3JvdXAvcGYtdjYtdG9vbGJhci1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvb2xiYXItaXRlbS9wZi12Ni10b29sYmFyLWl0ZW0uanMnO1xuXG4vLyBjZW0tZGV0YWlsLXBhbmVsIGFuZCBjZW0tdmlydHVhbC10cmVlIGFyZSBzdGlsbCBDZW1FbGVtZW50LWJhc2VkLFxuLy8gaW1wb3J0ZWQgYXMgc2lkZS1lZmZlY3QgSlMgZnJvbSB0aGVpciB0ZW1wbGF0ZSBsb2NhdGlvbnMuXG5pbXBvcnQgJy4uL2NlbS1kZXRhaWwtcGFuZWwvY2VtLWRldGFpbC1wYW5lbC5qcyc7XG5pbXBvcnQgJy4uL2NlbS12aXJ0dWFsLXRyZWUvY2VtLXZpcnR1YWwtdHJlZS5qcyc7XG5cbi8vIFJlLWV4cG9ydCBmb3IgdHlwZSB1c2FnZTsgQ2VtVmlydHVhbFRyZWUgcHJvdmlkZXMgbG9hZE1hbmlmZXN0KClcbi8vIEB0cy1pZ25vcmUgLS0gQ2VtRWxlbWVudC1iYXNlZCBKUywgbm8gdHlwZXMgYXZhaWxhYmxlXG5pbXBvcnQgeyBDZW1WaXJ0dWFsVHJlZSB9IGZyb20gJy4uL2NlbS12aXJ0dWFsLXRyZWUvY2VtLXZpcnR1YWwtdHJlZS5qcyc7XG5cbmludGVyZmFjZSBUcmVlSXRlbSB7XG4gIGlkOiBudW1iZXI7XG4gIHR5cGU6IHN0cmluZztcbiAgbGFiZWw6IHN0cmluZztcbiAgW2tleTogc3RyaW5nXTogdW5rbm93bjtcbn1cblxuaW50ZXJmYWNlIERldGFpbFBhbmVsIGV4dGVuZHMgSFRNTEVsZW1lbnQge1xuICByZW5kZXJJdGVtKGl0ZW06IFRyZWVJdGVtLCBtYW5pZmVzdDogdW5rbm93bik6IFByb21pc2U8dm9pZD47XG59XG5cbmludGVyZmFjZSBWaXJ0dWFsVHJlZSBleHRlbmRzIEhUTUxFbGVtZW50IHtcbiAgc2VhcmNoKHF1ZXJ5OiBzdHJpbmcpOiB2b2lkO1xuICBleHBhbmRBbGwoKTogdm9pZDtcbiAgY29sbGFwc2VBbGwoKTogdm9pZDtcbn1cblxuaW50ZXJmYWNlIEl0ZW1TZWxlY3RFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgaXRlbTogVHJlZUl0ZW07XG59XG5cbi8qKlxuICogTWFuaWZlc3QgQnJvd3NlciB3aXRoIHRyZWUgbmF2aWdhdGlvbiBhbmQgZGV0YWlsIGRyYXdlclxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90ICh1bnVzZWQpXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tbWFuaWZlc3QtYnJvd3NlcicpXG5leHBvcnQgY2xhc3MgQ2VtTWFuaWZlc3RCcm93c2VyIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgI3NlYXJjaERlYm91bmNlVGltZXI6IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+IHwgbnVsbCA9IG51bGw7XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cImRyYXdlci1jb250ZW50XCIgc2xvdD1cImNvbnRlbnRcIj5cbiAgICAgICAgPHBmLXY2LXRvb2xiYXIgc3RpY2t5PlxuICAgICAgICAgIDxwZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgPHBmLXY2LXRleHQtaW5wdXQtZ3JvdXAgaWQ9XCJzZWFyY2hcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB0eXBlPVwic2VhcmNoXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgcGxhY2Vob2xkZXI9XCJTZWFyY2ggbWFuaWZlc3QuLi5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiU2VhcmNoIG1hbmlmZXN0XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaWNvblxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBAaW5wdXQ9JHt0aGlzLiNvblNlYXJjaElucHV0fT5cbiAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCJcbiAgICAgICAgICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIlxuICAgICAgICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxZW1cIlxuICAgICAgICAgICAgICAgICAgICAgd2lkdGg9XCIxZW1cIlxuICAgICAgICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCA1MTIgNTEyXCI+XG4gICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTUwNSA0NDIuN0w0MDUuMyAzNDNjLTQuNS00LjUtMTAuNi03LTE3LTdIMzcyYzI3LjYtMzUuMyA0NC03OS43IDQ0LTEyOEM0MTYgOTMuMSAzMjIuOSAwIDIwOCAwUzAgOTMuMSAwIDIwOHM5My4xIDIwOCAyMDggMjA4YzQ4LjMgMCA5Mi43LTE2LjQgMTI4LTQ0djE2LjNjMCA2LjQgMi41IDEyLjUgNyAxN2w5OS43IDk5LjdjOS40IDkuNCAyNC42IDkuNCAzMy45IDBsMjguMy0yOC4zYzkuNC05LjQgOS40LTI0LjYuMS0zNHpNMjA4IDMzNmMtNzAuNyAwLTEyOC01Ny4yLTEyOC0xMjggMC03MC43IDU3LjItMTI4IDEyOC0xMjggNzAuNyAwIDEyOCA1Ny4yIDEyOCAxMjggMCA3MC43LTU3LjIgMTI4LTEyOCAxMjh6XCI+PC9wYXRoPlxuICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgIDxwZi12Ni1iYWRnZSBzbG90PVwidXRpbGl0aWVzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaWQ9XCJzZWFyY2gtY291bnRcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICBjb21wYWN0XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhpZGRlbj5cbiAgICAgICAgICAgICAgICAgIDA8c3BhbiBjbGFzcz1cInBmLXY2LXNjcmVlbi1yZWFkZXJcIj4gcmVzdWx0czwvc3Bhbj5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LWJhZGdlPlxuICAgICAgICAgICAgICAgIDxwZi12Ni1idXR0b24gc2xvdD1cInV0aWxpdGllc1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpZD1cInNlYXJjaC1jbGVhclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwicGxhaW5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkNsZWFyIHNlYXJjaFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICBoaWRkZW5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIEBjbGljaz0ke3RoaXMuI29uU2VhcmNoQ2xlYXJ9PlxuICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMzUyIDUxMlwiXG4gICAgICAgICAgICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWhpZGRlbj1cInRydWVcIj5cbiAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0yNDIuNzIgMjU2bDEwMC4wNy0xMDAuMDdjMTIuMjgtMTIuMjggMTIuMjgtMzIuMTkgMC00NC40OGwtMjIuMjQtMjIuMjRjLTEyLjI4LTEyLjI4LTMyLjE5LTEyLjI4LTQ0LjQ4IDBMMTc2IDE4OS4yOCA3NS45MyA4OS4yMWMtMTIuMjgtMTIuMjgtMzIuMTktMTIuMjgtNDQuNDggMEw5LjIxIDExMS40NWMtMTIuMjggMTIuMjgtMTIuMjggMzIuMTkgMCA0NC40OEwxMDkuMjggMjU2IDkuMjEgMzU2LjA3Yy0xMi4yOCAxMi4yOC0xMi4yOCAzMi4xOSAwIDQ0LjQ4bDIyLjI0IDIyLjI0YzEyLjI4IDEyLjI4IDMyLjIgMTIuMjggNDQuNDggMEwxNzYgMzIyLjcybDEwMC4wNyAxMDAuMDdjMTIuMjggMTIuMjggMzIuMiAxMi4yOCA0NC40OCAwbDIyLjI0LTIyLjI0YzEyLjI4LTEyLjI4IDEyLjI4LTMyLjE5IDAtNDQuNDhMMjQyLjcyIDI1NnpcIj48L3BhdGg+XG4gICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgPC9wZi12Ni10ZXh0LWlucHV0LWdyb3VwPlxuICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgPC9wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgIDxwZi12Ni10b29sYmFyLWdyb3VwIHZhcmlhbnQ9XCJhY3Rpb24tZ3JvdXBcIj5cbiAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJleHBhbmQtYWxsXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkV4cGFuZCBhbGwgdHJlZSBpdGVtc1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25FeHBhbmRBbGx9PlxuICAgICAgICAgICAgICAgIEV4cGFuZCBhbGxcbiAgICAgICAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJjb2xsYXBzZS1hbGxcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJ0ZXJ0aWFyeVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgc2l6ZT1cInNtYWxsXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiQ29sbGFwc2UgYWxsIHRyZWUgaXRlbXNcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIEBjbGljaz0ke3RoaXMuI29uQ29sbGFwc2VBbGx9PlxuICAgICAgICAgICAgICAgIENvbGxhcHNlIGFsbFxuICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgPC9wZi12Ni10b29sYmFyPlxuICAgICAgICA8cGYtdjYtZHJhd2VyIGlkPVwiZHJhd2VyXCI+XG4gICAgICAgICAgPGRpdiBpZD1cInRyZWUtd3JhcHBlclwiPlxuICAgICAgICAgICAgPGNlbS12aXJ0dWFsLXRyZWUgaWQ9XCJ2aXJ0dWFsLXRyZWVcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgQGl0ZW0tc2VsZWN0PSR7dGhpcy4jb25JdGVtU2VsZWN0fT48L2NlbS12aXJ0dWFsLXRyZWU+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGNlbS1kZXRhaWwtcGFuZWwgaWQ9XCJkZXRhaWwtcGFuZWxcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNsb3Q9XCJwYW5lbC1ib2R5XCI+PC9jZW0tZGV0YWlsLXBhbmVsPlxuICAgICAgICA8L3BmLXY2LWRyYXdlcj5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cblxuICBnZXQgI2RyYXdlcigpIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnZHJhd2VyJykgYXMgSFRNTEVsZW1lbnQgJiB7IGV4cGFuZGVkOiBib29sZWFuIH0gfCBudWxsO1xuICB9XG5cbiAgZ2V0ICN2aXJ0dWFsVHJlZSgpOiBWaXJ0dWFsVHJlZSB8IG51bGwge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCd2aXJ0dWFsLXRyZWUnKSBhcyBWaXJ0dWFsVHJlZSB8IG51bGw7XG4gIH1cblxuICBnZXQgI2RldGFpbFBhbmVsKCk6IERldGFpbFBhbmVsIHwgbnVsbCB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ2RldGFpbC1wYW5lbCcpIGFzIERldGFpbFBhbmVsIHwgbnVsbDtcbiAgfVxuXG4gIGdldCAjc2VhcmNoSW5wdXQoKSB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ3NlYXJjaCcpIGFzIEhUTUxFbGVtZW50ICYgeyB2YWx1ZTogc3RyaW5nIH0gfCBudWxsO1xuICB9XG5cbiAgZ2V0ICNzZWFyY2hDb3VudCgpIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnc2VhcmNoLWNvdW50JykgYXMgSFRNTEVsZW1lbnQgfCBudWxsO1xuICB9XG5cbiAgZ2V0ICNzZWFyY2hDbGVhcigpIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnc2VhcmNoLWNsZWFyJykgYXMgSFRNTEVsZW1lbnQgfCBudWxsO1xuICB9XG5cbiAgYXN5bmMgI29uSXRlbVNlbGVjdChlOiBJdGVtU2VsZWN0RXZlbnQpIHtcbiAgICBjb25zdCBpdGVtID0gZS5pdGVtO1xuICAgIGlmICghaXRlbSkgcmV0dXJuO1xuXG4gICAgY29uc3QgbWFuaWZlc3QgPSBhd2FpdCBDZW1WaXJ0dWFsVHJlZS5sb2FkTWFuaWZlc3QoKTtcbiAgICBpZiAodGhpcy4jZGV0YWlsUGFuZWwgJiYgbWFuaWZlc3QpIHtcbiAgICAgIGF3YWl0IHRoaXMuI2RldGFpbFBhbmVsLnJlbmRlckl0ZW0oaXRlbSwgbWFuaWZlc3QpO1xuICAgICAgaWYgKHRoaXMuI2RyYXdlcikge1xuICAgICAgICB0aGlzLiNkcmF3ZXIuZXhwYW5kZWQgPSB0cnVlO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNvblNlYXJjaElucHV0KCkge1xuICAgIGNvbnN0IHZhbHVlID0gdGhpcy4jc2VhcmNoSW5wdXQ/LnZhbHVlIHx8ICcnO1xuXG4gICAgLy8gU2hvdy9oaWRlIGNsZWFyIGJ1dHRvbiBiYXNlZCBvbiB3aGV0aGVyIHRoZXJlJ3MgYSB2YWx1ZVxuICAgIGlmICh0aGlzLiNzZWFyY2hDbGVhcikge1xuICAgICAgdGhpcy4jc2VhcmNoQ2xlYXIuaGlkZGVuID0gIXZhbHVlO1xuICAgIH1cblxuICAgIC8vIERlYm91bmNlIHNlYXJjaCAtIHdhaXQgMzAwbXMgYWZ0ZXIgdXNlciBzdG9wcyB0eXBpbmdcbiAgICBpZiAodGhpcy4jc2VhcmNoRGVib3VuY2VUaW1lciAhPSBudWxsKSB7XG4gICAgICBjbGVhclRpbWVvdXQodGhpcy4jc2VhcmNoRGVib3VuY2VUaW1lcik7XG4gICAgfVxuICAgIHRoaXMuI3NlYXJjaERlYm91bmNlVGltZXIgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgIHRoaXMuI2hhbmRsZVNlYXJjaCh2YWx1ZSk7XG4gICAgfSwgMzAwKTtcbiAgfVxuXG4gICNvblNlYXJjaENsZWFyKCkge1xuICAgIGlmICh0aGlzLiNzZWFyY2hJbnB1dCkge1xuICAgICAgdGhpcy4jc2VhcmNoSW5wdXQudmFsdWUgPSAnJztcbiAgICB9XG4gICAgaWYgKHRoaXMuI3NlYXJjaENsZWFyKSB7XG4gICAgICB0aGlzLiNzZWFyY2hDbGVhci5oaWRkZW4gPSB0cnVlO1xuICAgIH1cbiAgICBpZiAodGhpcy4jc2VhcmNoQ291bnQpIHtcbiAgICAgIHRoaXMuI3NlYXJjaENvdW50LmhpZGRlbiA9IHRydWU7XG4gICAgfVxuICAgIHRoaXMuI2hhbmRsZVNlYXJjaCgnJyk7XG4gIH1cblxuICAjaGFuZGxlU2VhcmNoKHF1ZXJ5OiBzdHJpbmcpIHtcbiAgICB0aGlzLiN2aXJ0dWFsVHJlZT8uc2VhcmNoKHF1ZXJ5KTtcbiAgICBpZiAodGhpcy4jc2VhcmNoQ291bnQpIHtcbiAgICAgIHRoaXMuI3NlYXJjaENvdW50LmhpZGRlbiA9ICFxdWVyeTtcbiAgICB9XG4gIH1cblxuICAjb25FeHBhbmRBbGwoKSB7XG4gICAgdGhpcy4jdmlydHVhbFRyZWU/LmV4cGFuZEFsbCgpO1xuICB9XG5cbiAgI29uQ29sbGFwc2VBbGwoKSB7XG4gICAgdGhpcy4jdmlydHVhbFRyZWU/LmNvbGxhcHNlQWxsKCk7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLW1hbmlmZXN0LWJyb3dzZXInOiBDZW1NYW5pZmVzdEJyb3dzZXI7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIi8qIE1hbmlmZXN0IEJyb3dzZXIgLSBDb21wb25lbnQgQVBJIG5hdmlnYXRpb24gdHJlZSAqL1xcXFxuXFxcXG46aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbltoaWRkZW5dIHtcXFxcbiAgZGlzcGxheTogbm9uZSAhaW1wb3J0YW50O1xcXFxufVxcXFxuXFxcXG4jZHJhd2VyIHtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4jZHJhd2VyLWNvbnRlbnQge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG59XFxcXG5cXFxcbiN0cmVlLXdyYXBwZXIge1xcXFxuICBmbGV4OiAxO1xcXFxuICBvdmVyZmxvdzogYXV0bztcXFxcbiAgbWluLWhlaWdodDogMDtcXFxcblxcXFxuICBcXFxcdTAwMjY6OnNsb3R0ZWQocGYtdjYtdHJlZS12aWV3KSB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQsIDFyZW0pO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNwYW5lbC10aXRsZSB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbn1cXFxcblxcXFxuI2RldGFpbHMtY29udGVudCB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCwgMXJlbSk7XFxcXG5cXFxcbiAgaDMge1xcXFxuICAgIG1hcmdpbi10b3A6IDA7XFxcXG4gICAgbWFyZ2luLWJvdHRvbTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtLCAwLjVyZW0pO1xcXFxuICAgIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1oZWFkaW5nLS1sZywgMS4yNXJlbSk7XFxcXG4gIH1cXFxcblxcXFxuICBkbCB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICB9XFxcXG5cXFxcbiAgZHQge1xcXFxuICAgIGZvbnQtd2VpZ2h0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXdlaWdodC0tYm9keS0tYm9sZCwgNjAwKTtcXFxcbiAgICBtYXJnaW4tdG9wOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20sIDAuNXJlbSk7XFxcXG4gIH1cXFxcblxcXFxuICBkZCB7XFxcXG4gICAgbWFyZ2luLWxlZnQ6IDA7XFxcXG4gICAgbWFyZ2luLWJvdHRvbTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtLCAwLjVyZW0pO1xcXFxuICB9XFxcXG5cXFxcbiAgLmVtcHR5LXN0YXRlIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlLCAjNmM3NTdkKTtcXFxcbiAgICBmb250LXN0eWxlOiBpdGFsaWM7XFxcXG4gIH1cXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCOzs7QUNEOUIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sdXFDQUF5cUMsQ0FBQztBQUFFLElBQU8sK0JBQVE7OztBREtodkMsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUlQLE9BQU87QUFDUCxPQUFPO0FBSVAsU0FBUyxzQkFBc0I7QUFwQi9CO0FBZ0RBLGtDQUFDLGNBQWMsc0JBQXNCO0FBQzlCLElBQU0scUJBQU4sZUFBaUMsaUJBQVc7QUFBQSxFQUNqRCxPQUFP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
