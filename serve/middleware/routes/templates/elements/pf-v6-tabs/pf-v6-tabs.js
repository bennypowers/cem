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
var __publicField = (obj, key, value) => __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
var __accessCheck = (obj, member, msg) => member.has(obj) || __typeError("Cannot " + msg);
var __privateIn = (member, obj) => Object(obj) !== obj ? __typeError('Cannot use the "in" operator on this value') : member.has(obj);
var __privateGet = (obj, member, getter) => (__accessCheck(obj, member, "read from private field"), getter ? getter.call(obj) : member.get(obj));
var __privateAdd = (obj, member, value) => member.has(obj) ? __typeError("Cannot add the same private member more than once") : member instanceof WeakSet ? member.add(obj) : member.set(obj, value);
var __privateSet = (obj, member, value, setter) => (__accessCheck(obj, member, "write to private field"), setter ? setter.call(obj, value) : member.set(obj, value), value);
var __privateMethod = (obj, member, method) => (__accessCheck(obj, member, "access private method"), method);

// elements/pf-v6-tabs/pf-v6-tabs.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";
import { state } from "/__cem/vendor/lit/decorators/state.js";

// lit-css:elements/pf-v6-tabs/pf-v6-tabs.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse(`":host {\\n  display: block;\\n  height: 100%;\\n  position: relative;\\n  overflow: hidden;\\n}\\n\\n#tabs-container {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n@property --_pf-v6-tabs--link-accent--length {\\n  syntax: \\"\\u003clength\\u003e\\";\\n  inherits: true;\\n  initial-value: 0px;\\n}\\n\\n#tabs {\\n  --_tab--Color: var(--pf-t--global--text--color--subtle);\\n  --_tab--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --_tab--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --_tab--BorderColor: transparent;\\n  --_tab--hover--Color: var(--pf-t--global--text--color--regular);\\n  --_tab--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --_tab--current--Color: var(--pf-t--global--text--color--regular);\\n  --_tab--current--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --_tab--disabled--Color: var(--pf-t--global--text--color--on-disabled);\\n  --_tab--disabled--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n\\n  --_pf-v6-tabs--link-accent--start: 0px;\\n  --_pf-v6-tabs--link-accent--length: 0px;\\n  --_pf-v6-tabs--link-accent--color: var(--pf-t--global--border--color--clicked);\\n  --_pf-v6-tabs--link-accent--border-size: var(--pf-t--global--border--width--extra-strong);\\n\\n  position: sticky;\\n  inset-block-start: 0;\\n  z-index: var(--pf-t--global--z-index--md);\\n  display: flex;\\n  overflow-x: auto;\\n  overflow-y: hidden;\\n  -webkit-overflow-scrolling: touch;\\n  scroll-behavior: smooth;\\n  flex-shrink: 0;\\n  background-color: var(--pf-t--global--background--color--primary--default);\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    inset-block-end: 0;\\n    inset-inline-start: 0;\\n    inset-inline-end: 0;\\n    border-block-end: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n  }\\n\\n  \\u0026::after {\\n    content: '';\\n    position: absolute;\\n    inset-block-start: auto;\\n    inset-block-end: 0;\\n    inset-inline-start: 0;\\n    width: var(--_pf-v6-tabs--link-accent--length);\\n    height: 0;\\n    border: 0 solid var(--_pf-v6-tabs--link-accent--color);\\n    border-block-end-width: var(--_pf-v6-tabs--link-accent--border-size);\\n    transform-origin: 0 center;\\n    translate: var(--_pf-v6-tabs--link-accent--start) 0;\\n    transition-property: width, translate;\\n    transition-duration: var(--pf-t--global--motion--duration--md);\\n    transition-timing-function: var(--pf-t--global--motion--timing-function--decelerate);\\n  }\\n}\\n\\n#panels {\\n  flex: 1;\\n  min-height: 0;\\n  overflow: hidden;\\n}\\n\\n.panel {\\n  display: block;\\n  height: 100%;\\n  overflow-y: auto;\\n  box-sizing: border-box;\\n  contain: layout style paint;\\n\\n  \\u0026[hidden] {\\n    display: none;\\n    content-visibility: hidden;\\n  }\\n}\\n\\n.tab {\\n  position: relative;\\n  display: flex;\\n  flex: none;\\n  align-items: center;\\n  padding: calc(var(--pf-t--global--spacer--sm) + var(--pf-t--global--spacer--xs))\\n           calc(var(--pf-t--global--spacer--sm) + var(--pf-t--global--spacer--sm));\\n  font-family: inherit;\\n  font-size: var(--pf-t--global--font--size--sm);\\n  font-weight: 400;\\n  color: var(--_tab--Color);\\n  text-decoration: none;\\n  background-color: transparent;\\n  border: 0;\\n  cursor: pointer;\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    inset-block-start: 0;\\n    inset-inline-start: 0;\\n    width: calc(100% - calc(var(--pf-t--global--spacer--sm) * 2));\\n    height: calc(100% - calc(var(--pf-t--global--spacer--sm) * 2));\\n    translate: var(--pf-t--global--spacer--sm) var(--pf-t--global--spacer--sm);\\n    background-color: var(--_tab--BackgroundColor);\\n    border: var(--_tab--BorderWidth) solid var(--_tab--BorderColor);\\n    border-radius: var(--pf-t--global--border--radius--small);\\n    transition: background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n    z-index: -1;\\n  }\\n\\n  \\u0026:where(:hover, :focus) {\\n    --_tab--Color: var(--_tab--hover--Color);\\n    --_tab--BackgroundColor: var(--_tab--hover--BackgroundColor);\\n  }\\n\\n  \\u0026[aria-selected=\\"true\\"] {\\n    --_tab--Color: var(--_tab--current--Color);\\n    --_tab--BackgroundColor: var(--_tab--current--BackgroundColor);\\n  }\\n\\n  \\u0026:disabled {\\n    --_tab--Color: var(--_tab--disabled--Color);\\n    --_tab--BackgroundColor: var(--_tab--disabled--BackgroundColor);\\n    cursor: not-allowed;\\n    pointer-events: none;\\n  }\\n}\\n"`));
var pf_v6_tabs_default = s;

// elements/pf-v6-tabs/pf-v6-tabs.ts
import "../pf-v6-tab/pf-v6-tab.js";
var PfTabsChangeEvent = class extends Event {
  selectedIndex;
  constructor(selectedIndex) {
    super("change", { bubbles: true });
    this.selectedIndex = selectedIndex;
  }
};
var __tabs_dec, _selected_dec, _a, _PfV6Tabs_decorators, _init, _selected, __tabs, _mutationObserver, _onTabsChanged, _PfV6Tabs_instances, syncTabs_fn, handleTabClick_fn, handleKeyDown_fn, updateAccentLine_fn;
_PfV6Tabs_decorators = [customElement("pf-v6-tabs")];
var PfV6Tabs = class extends (_a = LitElement, _selected_dec = [property({ type: Number, reflect: true })], __tabs_dec = [state()], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Tabs_instances);
    __privateAdd(this, _selected, __runInitializers(_init, 8, this, 0)), __runInitializers(_init, 11, this);
    __privateAdd(this, __tabs, __runInitializers(_init, 12, this, [])), __runInitializers(_init, 15, this);
    __privateAdd(this, _mutationObserver);
    __privateAdd(this, _onTabsChanged, () => {
      __privateMethod(this, _PfV6Tabs_instances, syncTabs_fn).call(this);
    });
  }
  get selectedIndex() {
    return this.selected;
  }
  set selectedIndex(index) {
    this.selected = Math.max(0, Math.min(index, this._tabs.length - 1));
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("pf-v6-tab-title-changed", __privateGet(this, _onTabsChanged));
    this.addEventListener("pf-v6-tab-connected", __privateGet(this, _onTabsChanged));
    this.addEventListener("pf-v6-tab-disconnected", __privateGet(this, _onTabsChanged));
    __privateSet(this, _mutationObserver, new MutationObserver(__privateGet(this, _onTabsChanged)));
    __privateGet(this, _mutationObserver).observe(this, { childList: true });
  }
  firstUpdated() {
    __privateMethod(this, _PfV6Tabs_instances, syncTabs_fn).call(this);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(this, _mutationObserver)?.disconnect();
    this.removeEventListener("pf-v6-tab-title-changed", __privateGet(this, _onTabsChanged));
    this.removeEventListener("pf-v6-tab-connected", __privateGet(this, _onTabsChanged));
    this.removeEventListener("pf-v6-tab-disconnected", __privateGet(this, _onTabsChanged));
  }
  updated(changed) {
    if (changed.has("selected") || changed.has("_tabs")) {
      this._tabs.forEach((tab, i) => {
        tab.setAttribute("slot", `panel-${i}`);
      });
      requestAnimationFrame(() => __privateMethod(this, _PfV6Tabs_instances, updateAccentLine_fn).call(this));
    }
  }
  render() {
    return html`
      <div id="tabs-container">
        <div id="tabs"
             role="tablist"
             aria-label="Tabs"
             part="tabs"
             @keydown=${__privateMethod(this, _PfV6Tabs_instances, handleKeyDown_fn)}>
          ${this._tabs.map((tab, index) => html`
            <button class="tab"
                    type="button"
                    role="tab"
                    id="tab-${index}"
                    aria-controls="panel-${index}"
                    aria-selected="${index === this.selected ? "true" : "false"}"
                    tabindex="${index === this.selected ? 0 : -1}"
                    @click=${() => __privateMethod(this, _PfV6Tabs_instances, handleTabClick_fn).call(this, index)}>
              ${tab.title || `Tab ${index + 1}`}
            </button>
          `)}
        </div>
        <div id="panels"
             part="panels">
          ${this._tabs.map((_, index) => html`
            <div class="panel"
                 id="panel-${index}"
                 role="tabpanel"
                 aria-labelledby="tab-${index}"
                 ?hidden=${index !== this.selected}>
              <slot name="panel-${index}"></slot>
            </div>
          `)}
        </div>
      </div>
    `;
  }
};
_init = __decoratorStart(_a);
_selected = new WeakMap();
__tabs = new WeakMap();
_mutationObserver = new WeakMap();
_onTabsChanged = new WeakMap();
_PfV6Tabs_instances = new WeakSet();
syncTabs_fn = function() {
  this._tabs = Array.from(this.querySelectorAll("pf-v6-tab"));
  this.requestUpdate();
};
handleTabClick_fn = function(index) {
  const oldIndex = this.selected;
  this.selected = index;
  if (oldIndex !== index) {
    this.dispatchEvent(new PfTabsChangeEvent(index));
  }
};
handleKeyDown_fn = function(e) {
  const tabsEl = this.shadowRoot?.getElementById("tabs");
  if (!tabsEl) return;
  const buttons = Array.from(tabsEl.querySelectorAll(".tab"));
  const currentIndex = buttons.findIndex((btn) => btn === this.shadowRoot?.activeElement);
  if (currentIndex === -1) return;
  let nextIndex = currentIndex;
  switch (e.key) {
    case "ArrowLeft":
      e.preventDefault();
      nextIndex = currentIndex > 0 ? currentIndex - 1 : buttons.length - 1;
      break;
    case "ArrowRight":
      e.preventDefault();
      nextIndex = currentIndex < buttons.length - 1 ? currentIndex + 1 : 0;
      break;
    case "Home":
      e.preventDefault();
      nextIndex = 0;
      break;
    case "End":
      e.preventDefault();
      nextIndex = buttons.length - 1;
      break;
    default:
      return;
  }
  __privateMethod(this, _PfV6Tabs_instances, handleTabClick_fn).call(this, nextIndex);
  buttons[nextIndex].focus();
};
updateAccentLine_fn = function() {
  const tabsEl = this.shadowRoot?.getElementById("tabs");
  if (!tabsEl) return;
  const buttons = Array.from(tabsEl.querySelectorAll(".tab"));
  const activeButton = buttons[this.selected];
  if (!activeButton) return;
  const containerRect = tabsEl.getBoundingClientRect();
  const buttonRect = activeButton.getBoundingClientRect();
  const inset = 8;
  const start = buttonRect.left - containerRect.left + inset;
  const length = buttonRect.width - inset * 2;
  tabsEl.style.setProperty("--_pf-v6-tabs--link-accent--start", `${start}px`);
  tabsEl.style.setProperty("--_pf-v6-tabs--link-accent--length", `${length}px`);
};
__decorateElement(_init, 4, "selected", _selected_dec, PfV6Tabs, _selected);
__decorateElement(_init, 4, "_tabs", __tabs_dec, PfV6Tabs, __tabs);
PfV6Tabs = __decorateElement(_init, 0, "PfV6Tabs", _PfV6Tabs_decorators, PfV6Tabs);
__publicField(PfV6Tabs, "styles", pf_v6_tabs_default);
__runInitializers(_init, 1, PfV6Tabs);
export {
  PfTabsChangeEvent,
  PfV6Tabs
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtdGFicy9wZi12Ni10YWJzLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvcGYtdjYtdGFicy9wZi12Ni10YWJzLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL3BmLXY2LXRhYnMuY3NzJztcblxuaW1wb3J0ICcuLi9wZi12Ni10YWIvcGYtdjYtdGFiLmpzJztcblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiB0YWIgc2VsZWN0aW9uIGNoYW5nZXNcbiAqL1xuZXhwb3J0IGNsYXNzIFBmVGFic0NoYW5nZUV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBzZWxlY3RlZEluZGV4OiBudW1iZXI7XG4gIGNvbnN0cnVjdG9yKHNlbGVjdGVkSW5kZXg6IG51bWJlcikge1xuICAgIHN1cGVyKCdjaGFuZ2UnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gICAgdGhpcy5zZWxlY3RlZEluZGV4ID0gc2VsZWN0ZWRJbmRleDtcbiAgfVxufVxuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgVGFic1xuICpcbiAqIEEgdGFiYmVkIGNvbnRlbnQgY29udGFpbmVyLiBBZGQgYHBmLXY2LXRhYmAgY2hpbGRyZW4gd2l0aCBgdGl0bGVgIGF0dHJpYnV0ZXNcbiAqIHRvIGRlZmluZSB0YWJzIGFuZCB0aGVpciBwYW5lbCBjb250ZW50LlxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBgcGYtdjYtdGFiYCBjaGlsZHJlblxuICpcbiAqIEBmaXJlcyBjaGFuZ2UgLSBGaXJlZCB3aGVuIHRoZSBzZWxlY3RlZCB0YWIgY2hhbmdlc1xuICpcbiAqIEBjc3NwYXJ0IHRhYnMgLSBUaGUgdGFiIGJ1dHRvbiBjb250YWluZXJcbiAqIEBjc3NwYXJ0IHBhbmVscyAtIFRoZSB0YWIgcGFuZWwgY29udGFpbmVyXG4gKi9cbkBjdXN0b21FbGVtZW50KCdwZi12Ni10YWJzJylcbmV4cG9ydCBjbGFzcyBQZlY2VGFicyBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IE51bWJlciwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBzZWxlY3RlZCA9IDA7XG5cbiAgQHN0YXRlKClcbiAgYWNjZXNzb3IgX3RhYnM6IEVsZW1lbnRbXSA9IFtdO1xuXG4gICNtdXRhdGlvbk9ic2VydmVyPzogTXV0YXRpb25PYnNlcnZlcjtcblxuICBnZXQgc2VsZWN0ZWRJbmRleCgpOiBudW1iZXIge1xuICAgIHJldHVybiB0aGlzLnNlbGVjdGVkO1xuICB9XG5cbiAgc2V0IHNlbGVjdGVkSW5kZXgoaW5kZXg6IG51bWJlcikge1xuICAgIHRoaXMuc2VsZWN0ZWQgPSBNYXRoLm1heCgwLCBNYXRoLm1pbihpbmRleCwgdGhpcy5fdGFicy5sZW5ndGggLSAxKSk7XG4gIH1cblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigncGYtdjYtdGFiLXRpdGxlLWNoYW5nZWQnLCB0aGlzLiNvblRhYnNDaGFuZ2VkKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ3BmLXY2LXRhYi1jb25uZWN0ZWQnLCB0aGlzLiNvblRhYnNDaGFuZ2VkKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ3BmLXY2LXRhYi1kaXNjb25uZWN0ZWQnLCB0aGlzLiNvblRhYnNDaGFuZ2VkKTtcbiAgICB0aGlzLiNtdXRhdGlvbk9ic2VydmVyID0gbmV3IE11dGF0aW9uT2JzZXJ2ZXIodGhpcy4jb25UYWJzQ2hhbmdlZCk7XG4gICAgdGhpcy4jbXV0YXRpb25PYnNlcnZlci5vYnNlcnZlKHRoaXMsIHsgY2hpbGRMaXN0OiB0cnVlIH0pO1xuICB9XG5cbiAgZmlyc3RVcGRhdGVkKCkge1xuICAgIHRoaXMuI3N5bmNUYWJzKCk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuI211dGF0aW9uT2JzZXJ2ZXI/LmRpc2Nvbm5lY3QoKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ3BmLXY2LXRhYi10aXRsZS1jaGFuZ2VkJywgdGhpcy4jb25UYWJzQ2hhbmdlZCk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdwZi12Ni10YWItY29ubmVjdGVkJywgdGhpcy4jb25UYWJzQ2hhbmdlZCk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdwZi12Ni10YWItZGlzY29ubmVjdGVkJywgdGhpcy4jb25UYWJzQ2hhbmdlZCk7XG4gIH1cblxuICAjb25UYWJzQ2hhbmdlZCA9ICgpID0+IHtcbiAgICB0aGlzLiNzeW5jVGFicygpO1xuICB9O1xuXG4gICNzeW5jVGFicygpIHtcbiAgICB0aGlzLl90YWJzID0gQXJyYXkuZnJvbSh0aGlzLnF1ZXJ5U2VsZWN0b3JBbGwoJ3BmLXY2LXRhYicpKTtcbiAgICB0aGlzLnJlcXVlc3RVcGRhdGUoKTtcbiAgfVxuXG4gIHVwZGF0ZWQoY2hhbmdlZDogTWFwPHN0cmluZywgdW5rbm93bj4pIHtcbiAgICBpZiAoY2hhbmdlZC5oYXMoJ3NlbGVjdGVkJykgfHwgY2hhbmdlZC5oYXMoJ190YWJzJykpIHtcbiAgICAgIC8vIEFzc2lnbiBzbG90IG5hbWVzIHRvIHRhYiBjaGlsZHJlblxuICAgICAgdGhpcy5fdGFicy5mb3JFYWNoKCh0YWIsIGkpID0+IHtcbiAgICAgICAgdGFiLnNldEF0dHJpYnV0ZSgnc2xvdCcsIGBwYW5lbC0ke2l9YCk7XG4gICAgICB9KTtcbiAgICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB0aGlzLiN1cGRhdGVBY2NlbnRMaW5lKCkpO1xuICAgIH1cbiAgfVxuXG4gICNoYW5kbGVUYWJDbGljayhpbmRleDogbnVtYmVyKSB7XG4gICAgY29uc3Qgb2xkSW5kZXggPSB0aGlzLnNlbGVjdGVkO1xuICAgIHRoaXMuc2VsZWN0ZWQgPSBpbmRleDtcbiAgICBpZiAob2xkSW5kZXggIT09IGluZGV4KSB7XG4gICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IFBmVGFic0NoYW5nZUV2ZW50KGluZGV4KSk7XG4gICAgfVxuICB9XG5cbiAgI2hhbmRsZUtleURvd24oZTogS2V5Ym9hcmRFdmVudCkge1xuICAgIGNvbnN0IHRhYnNFbCA9IHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ3RhYnMnKTtcbiAgICBpZiAoIXRhYnNFbCkgcmV0dXJuO1xuICAgIGNvbnN0IGJ1dHRvbnMgPSBBcnJheS5mcm9tKHRhYnNFbC5xdWVyeVNlbGVjdG9yQWxsKCcudGFiJykpIGFzIEhUTUxCdXR0b25FbGVtZW50W107XG4gICAgY29uc3QgY3VycmVudEluZGV4ID0gYnV0dG9ucy5maW5kSW5kZXgoYnRuID0+IGJ0biA9PT0gdGhpcy5zaGFkb3dSb290Py5hY3RpdmVFbGVtZW50KTtcbiAgICBpZiAoY3VycmVudEluZGV4ID09PSAtMSkgcmV0dXJuO1xuXG4gICAgbGV0IG5leHRJbmRleCA9IGN1cnJlbnRJbmRleDtcblxuICAgIHN3aXRjaCAoZS5rZXkpIHtcbiAgICAgIGNhc2UgJ0Fycm93TGVmdCc6XG4gICAgICAgIGUucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgbmV4dEluZGV4ID0gY3VycmVudEluZGV4ID4gMCA/IGN1cnJlbnRJbmRleCAtIDEgOiBidXR0b25zLmxlbmd0aCAtIDE7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnQXJyb3dSaWdodCc6XG4gICAgICAgIGUucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgbmV4dEluZGV4ID0gY3VycmVudEluZGV4IDwgYnV0dG9ucy5sZW5ndGggLSAxID8gY3VycmVudEluZGV4ICsgMSA6IDA7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnSG9tZSc6XG4gICAgICAgIGUucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgbmV4dEluZGV4ID0gMDtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdFbmQnOlxuICAgICAgICBlLnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIG5leHRJbmRleCA9IGJ1dHRvbnMubGVuZ3RoIC0gMTtcbiAgICAgICAgYnJlYWs7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgdGhpcy4jaGFuZGxlVGFiQ2xpY2sobmV4dEluZGV4KTtcbiAgICBidXR0b25zW25leHRJbmRleF0uZm9jdXMoKTtcbiAgfVxuXG4gICN1cGRhdGVBY2NlbnRMaW5lKCkge1xuICAgIGNvbnN0IHRhYnNFbCA9IHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ3RhYnMnKTtcbiAgICBpZiAoIXRhYnNFbCkgcmV0dXJuO1xuICAgIGNvbnN0IGJ1dHRvbnMgPSBBcnJheS5mcm9tKHRhYnNFbC5xdWVyeVNlbGVjdG9yQWxsKCcudGFiJykpO1xuICAgIGNvbnN0IGFjdGl2ZUJ1dHRvbiA9IGJ1dHRvbnNbdGhpcy5zZWxlY3RlZF0gYXMgSFRNTEVsZW1lbnQgfCB1bmRlZmluZWQ7XG4gICAgaWYgKCFhY3RpdmVCdXR0b24pIHJldHVybjtcblxuICAgIGNvbnN0IGNvbnRhaW5lclJlY3QgPSB0YWJzRWwuZ2V0Qm91bmRpbmdDbGllbnRSZWN0KCk7XG4gICAgY29uc3QgYnV0dG9uUmVjdCA9IGFjdGl2ZUJ1dHRvbi5nZXRCb3VuZGluZ0NsaWVudFJlY3QoKTtcbiAgICBjb25zdCBpbnNldCA9IDg7IC8vIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSlcbiAgICBjb25zdCBzdGFydCA9IChidXR0b25SZWN0LmxlZnQgLSBjb250YWluZXJSZWN0LmxlZnQpICsgaW5zZXQ7XG4gICAgY29uc3QgbGVuZ3RoID0gYnV0dG9uUmVjdC53aWR0aCAtIChpbnNldCAqIDIpO1xuXG4gICAgdGFic0VsLnN0eWxlLnNldFByb3BlcnR5KCctLV9wZi12Ni10YWJzLS1saW5rLWFjY2VudC0tc3RhcnQnLCBgJHtzdGFydH1weGApO1xuICAgIHRhYnNFbC5zdHlsZS5zZXRQcm9wZXJ0eSgnLS1fcGYtdjYtdGFicy0tbGluay1hY2NlbnQtLWxlbmd0aCcsIGAke2xlbmd0aH1weGApO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cInRhYnMtY29udGFpbmVyXCI+XG4gICAgICAgIDxkaXYgaWQ9XCJ0YWJzXCJcbiAgICAgICAgICAgICByb2xlPVwidGFibGlzdFwiXG4gICAgICAgICAgICAgYXJpYS1sYWJlbD1cIlRhYnNcIlxuICAgICAgICAgICAgIHBhcnQ9XCJ0YWJzXCJcbiAgICAgICAgICAgICBAa2V5ZG93bj0ke3RoaXMuI2hhbmRsZUtleURvd259PlxuICAgICAgICAgICR7dGhpcy5fdGFicy5tYXAoKHRhYiwgaW5kZXgpID0+IGh0bWxgXG4gICAgICAgICAgICA8YnV0dG9uIGNsYXNzPVwidGFiXCJcbiAgICAgICAgICAgICAgICAgICAgdHlwZT1cImJ1dHRvblwiXG4gICAgICAgICAgICAgICAgICAgIHJvbGU9XCJ0YWJcIlxuICAgICAgICAgICAgICAgICAgICBpZD1cInRhYi0ke2luZGV4fVwiXG4gICAgICAgICAgICAgICAgICAgIGFyaWEtY29udHJvbHM9XCJwYW5lbC0ke2luZGV4fVwiXG4gICAgICAgICAgICAgICAgICAgIGFyaWEtc2VsZWN0ZWQ9XCIke2luZGV4ID09PSB0aGlzLnNlbGVjdGVkID8gJ3RydWUnIDogJ2ZhbHNlJ31cIlxuICAgICAgICAgICAgICAgICAgICB0YWJpbmRleD1cIiR7aW5kZXggPT09IHRoaXMuc2VsZWN0ZWQgPyAwIDogLTF9XCJcbiAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7KCkgPT4gdGhpcy4jaGFuZGxlVGFiQ2xpY2soaW5kZXgpfT5cbiAgICAgICAgICAgICAgJHsodGFiIGFzIEhUTUxFbGVtZW50KS50aXRsZSB8fCBgVGFiICR7aW5kZXggKyAxfWB9XG4gICAgICAgICAgICA8L2J1dHRvbj5cbiAgICAgICAgICBgKX1cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgaWQ9XCJwYW5lbHNcIlxuICAgICAgICAgICAgIHBhcnQ9XCJwYW5lbHNcIj5cbiAgICAgICAgICAke3RoaXMuX3RhYnMubWFwKChfLCBpbmRleCkgPT4gaHRtbGBcbiAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJwYW5lbFwiXG4gICAgICAgICAgICAgICAgIGlkPVwicGFuZWwtJHtpbmRleH1cIlxuICAgICAgICAgICAgICAgICByb2xlPVwidGFicGFuZWxcIlxuICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsbGVkYnk9XCJ0YWItJHtpbmRleH1cIlxuICAgICAgICAgICAgICAgICA/aGlkZGVuPSR7aW5kZXggIT09IHRoaXMuc2VsZWN0ZWR9PlxuICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwicGFuZWwtJHtpbmRleH1cIj48L3Nsb3Q+XG4gICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICBgKX1cbiAgICAgICAgPC9kaXY+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LXRhYnMnOiBQZlY2VGFicztcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG59XFxcXG5cXFxcbiN0YWJzLWNvbnRhaW5lciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxuQHByb3BlcnR5IC0tX3BmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1sZW5ndGgge1xcXFxuICBzeW50YXg6IFxcXFxcXFwiXFxcXHUwMDNjbGVuZ3RoXFxcXHUwMDNlXFxcXFxcXCI7XFxcXG4gIGluaGVyaXRzOiB0cnVlO1xcXFxuICBpbml0aWFsLXZhbHVlOiAwcHg7XFxcXG59XFxcXG5cXFxcbiN0YWJzIHtcXFxcbiAgLS1fdGFiLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgLS1fdGFiLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWRlZmF1bHQpO1xcXFxuICAtLV90YWItLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWFjdGlvbi0tcGxhaW4tLWhvdmVyKTtcXFxcbiAgLS1fdGFiLS1Cb3JkZXJDb2xvcjogdHJhbnNwYXJlbnQ7XFxcXG4gIC0tX3RhYi0taG92ZXItLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1fdGFiLS1ob3Zlci0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1ob3Zlcik7XFxcXG4gIC0tX3RhYi0tY3VycmVudC0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAtLV90YWItLWN1cnJlbnQtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0tZGVmYXVsdCk7XFxcXG4gIC0tX3RhYi0tZGlzYWJsZWQtLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1kaXNhYmxlZCk7XFxcXG4gIC0tX3RhYi0tZGlzYWJsZWQtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tZGlzYWJsZWQtLWRlZmF1bHQpO1xcXFxuXFxcXG4gIC0tX3BmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1zdGFydDogMHB4O1xcXFxuICAtLV9wZi12Ni10YWJzLS1saW5rLWFjY2VudC0tbGVuZ3RoOiAwcHg7XFxcXG4gIC0tX3BmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1jbGlja2VkKTtcXFxcbiAgLS1fcGYtdjYtdGFicy0tbGluay1hY2NlbnQtLWJvcmRlci1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWV4dHJhLXN0cm9uZyk7XFxcXG5cXFxcbiAgcG9zaXRpb246IHN0aWNreTtcXFxcbiAgaW5zZXQtYmxvY2stc3RhcnQ6IDA7XFxcXG4gIHotaW5kZXg6IHZhcigtLXBmLXQtLWdsb2JhbC0tei1pbmRleC0tbWQpO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBvdmVyZmxvdy14OiBhdXRvO1xcXFxuICBvdmVyZmxvdy15OiBoaWRkZW47XFxcXG4gIC13ZWJraXQtb3ZlcmZsb3ctc2Nyb2xsaW5nOiB0b3VjaDtcXFxcbiAgc2Nyb2xsLWJlaGF2aW9yOiBzbW9vdGg7XFxcXG4gIGZsZXgtc2hyaW5rOiAwO1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1kZWZhdWx0KTtcXFxcblxcXFxuICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgY29udGVudDogJyc7XFxcXG4gICAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICAgIGluc2V0LWJsb2NrLWVuZDogMDtcXFxcbiAgICBpbnNldC1pbmxpbmUtc3RhcnQ6IDA7XFxcXG4gICAgaW5zZXQtaW5saW5lLWVuZDogMDtcXFxcbiAgICBib3JkZXItYmxvY2stZW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLXJlZ3VsYXIpIHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6OmFmdGVyIHtcXFxcbiAgICBjb250ZW50OiAnJztcXFxcbiAgICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gICAgaW5zZXQtYmxvY2stc3RhcnQ6IGF1dG87XFxcXG4gICAgaW5zZXQtYmxvY2stZW5kOiAwO1xcXFxuICAgIGluc2V0LWlubGluZS1zdGFydDogMDtcXFxcbiAgICB3aWR0aDogdmFyKC0tX3BmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1sZW5ndGgpO1xcXFxuICAgIGhlaWdodDogMDtcXFxcbiAgICBib3JkZXI6IDAgc29saWQgdmFyKC0tX3BmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1jb2xvcik7XFxcXG4gICAgYm9yZGVyLWJsb2NrLWVuZC13aWR0aDogdmFyKC0tX3BmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1ib3JkZXItc2l6ZSk7XFxcXG4gICAgdHJhbnNmb3JtLW9yaWdpbjogMCBjZW50ZXI7XFxcXG4gICAgdHJhbnNsYXRlOiB2YXIoLS1fcGYtdjYtdGFicy0tbGluay1hY2NlbnQtLXN0YXJ0KSAwO1xcXFxuICAgIHRyYW5zaXRpb24tcHJvcGVydHk6IHdpZHRoLCB0cmFuc2xhdGU7XFxcXG4gICAgdHJhbnNpdGlvbi1kdXJhdGlvbjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1tZCk7XFxcXG4gICAgdHJhbnNpdGlvbi10aW1pbmctZnVuY3Rpb246IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlY2VsZXJhdGUpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNwYW5lbHMge1xcXFxuICBmbGV4OiAxO1xcXFxuICBtaW4taGVpZ2h0OiAwO1xcXFxuICBvdmVyZmxvdzogaGlkZGVuO1xcXFxufVxcXFxuXFxcXG4ucGFuZWwge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuICBib3gtc2l6aW5nOiBib3JkZXItYm94O1xcXFxuICBjb250YWluOiBsYXlvdXQgc3R5bGUgcGFpbnQ7XFxcXG5cXFxcbiAgXFxcXHUwMDI2W2hpZGRlbl0ge1xcXFxuICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gICAgY29udGVudC12aXNpYmlsaXR5OiBoaWRkZW47XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLnRhYiB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleDogbm9uZTtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgcGFkZGluZzogY2FsYyh2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pICsgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXhzKSlcXFxcbiAgICAgICAgICAgY2FsYyh2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pICsgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKSk7XFxcXG4gIGZvbnQtZmFtaWx5OiBpbmhlcml0O1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tc20pO1xcXFxuICBmb250LXdlaWdodDogNDAwO1xcXFxuICBjb2xvcjogdmFyKC0tX3RhYi0tQ29sb3IpO1xcXFxuICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHRyYW5zcGFyZW50O1xcXFxuICBib3JkZXI6IDA7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcblxcXFxuICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgY29udGVudDogJyc7XFxcXG4gICAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICAgIGluc2V0LWJsb2NrLXN0YXJ0OiAwO1xcXFxuICAgIGluc2V0LWlubGluZS1zdGFydDogMDtcXFxcbiAgICB3aWR0aDogY2FsYygxMDAlIC0gY2FsYyh2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pICogMikpO1xcXFxuICAgIGhlaWdodDogY2FsYygxMDAlIC0gY2FsYyh2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pICogMikpO1xcXFxuICAgIHRyYW5zbGF0ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKSB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLV90YWItLUJhY2tncm91bmRDb2xvcik7XFxcXG4gICAgYm9yZGVyOiB2YXIoLS1fdGFiLS1Cb3JkZXJXaWR0aCkgc29saWQgdmFyKC0tX3RhYi0tQm9yZGVyQ29sb3IpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXNtYWxsKTtcXFxcbiAgICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kLWNvbG9yIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpO1xcXFxuICAgIHotaW5kZXg6IC0xO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OndoZXJlKDpob3ZlciwgOmZvY3VzKSB7XFxcXG4gICAgLS1fdGFiLS1Db2xvcjogdmFyKC0tX3RhYi0taG92ZXItLUNvbG9yKTtcXFxcbiAgICAtLV90YWItLUJhY2tncm91bmRDb2xvcjogdmFyKC0tX3RhYi0taG92ZXItLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjZbYXJpYS1zZWxlY3RlZD1cXFxcXFxcInRydWVcXFxcXFxcIl0ge1xcXFxuICAgIC0tX3RhYi0tQ29sb3I6IHZhcigtLV90YWItLWN1cnJlbnQtLUNvbG9yKTtcXFxcbiAgICAtLV90YWItLUJhY2tncm91bmRDb2xvcjogdmFyKC0tX3RhYi0tY3VycmVudC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjpkaXNhYmxlZCB7XFxcXG4gICAgLS1fdGFiLS1Db2xvcjogdmFyKC0tX3RhYi0tZGlzYWJsZWQtLUNvbG9yKTtcXFxcbiAgICAtLV90YWItLUJhY2tncm91bmRDb2xvcjogdmFyKC0tX3RhYi0tZGlzYWJsZWQtLUJhY2tncm91bmRDb2xvcik7XFxcXG4gICAgY3Vyc29yOiBub3QtYWxsb3dlZDtcXFxcbiAgICBwb2ludGVyLWV2ZW50czogbm9uZTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjtBQUN6QixTQUFTLGFBQWE7OztBQ0h0QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSwwbkpBQWdvSixDQUFDO0FBQUUsSUFBTyxxQkFBUTs7O0FET3ZzSixPQUFPO0FBS0EsSUFBTSxvQkFBTixjQUFnQyxNQUFNO0FBQUEsRUFDM0M7QUFBQSxFQUNBLFlBQVksZUFBdUI7QUFDakMsVUFBTSxVQUFVLEVBQUUsU0FBUyxLQUFLLENBQUM7QUFDakMsU0FBSyxn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
