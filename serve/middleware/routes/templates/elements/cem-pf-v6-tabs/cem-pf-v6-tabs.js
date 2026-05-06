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

// elements/cem-pf-v6-tabs/cem-pf-v6-tabs.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";
import { state } from "/__cem/vendor/lit/decorators/state.js";

// lit-css:elements/cem-pf-v6-tabs/cem-pf-v6-tabs.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse(`":host {\\n  display: block;\\n  height: 100%;\\n  position: relative;\\n  overflow: hidden;\\n}\\n\\n#tabs-container {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n@property --_cem-pf-v6-tabs--link-accent--length {\\n  syntax: \\"\\u003clength\\u003e\\";\\n  inherits: true;\\n  initial-value: 0px;\\n}\\n\\n#tabs {\\n  --_tab--Color: var(--pf-t--global--text--color--subtle);\\n  --_tab--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --_tab--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --_tab--BorderColor: transparent;\\n  --_tab--hover--Color: var(--pf-t--global--text--color--regular);\\n  --_tab--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --_tab--current--Color: var(--pf-t--global--text--color--regular);\\n  --_tab--current--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --_tab--disabled--Color: var(--pf-t--global--text--color--on-disabled);\\n  --_tab--disabled--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n\\n  --_cem-pf-v6-tabs--link-accent--start: 0px;\\n  --_cem-pf-v6-tabs--link-accent--length: 0px;\\n  --_cem-pf-v6-tabs--link-accent--color: var(--pf-t--global--border--color--clicked);\\n  --_cem-pf-v6-tabs--link-accent--border-size: var(--pf-t--global--border--width--extra-strong);\\n\\n  position: sticky;\\n  inset-block-start: 0;\\n  z-index: var(--pf-t--global--z-index--md);\\n  display: flex;\\n  overflow-x: auto;\\n  overflow-y: hidden;\\n  -webkit-overflow-scrolling: touch;\\n  scroll-behavior: smooth;\\n  flex-shrink: 0;\\n  background-color: var(--pf-t--global--background--color--primary--default);\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    inset-block-end: 0;\\n    inset-inline-start: 0;\\n    inset-inline-end: 0;\\n    border-block-end: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n  }\\n\\n  \\u0026::after {\\n    content: '';\\n    position: absolute;\\n    inset-block-start: auto;\\n    inset-block-end: 0;\\n    inset-inline-start: 0;\\n    width: var(--_cem-pf-v6-tabs--link-accent--length);\\n    height: 0;\\n    border: 0 solid var(--_cem-pf-v6-tabs--link-accent--color);\\n    border-block-end-width: var(--_cem-pf-v6-tabs--link-accent--border-size);\\n    transform-origin: 0 center;\\n    translate: var(--_cem-pf-v6-tabs--link-accent--start) 0;\\n    transition-property: width, translate;\\n    transition-duration: var(--pf-t--global--motion--duration--md);\\n    transition-timing-function: var(--pf-t--global--motion--timing-function--decelerate);\\n  }\\n}\\n\\n#panels {\\n  flex: 1;\\n  min-height: 0;\\n  overflow: hidden;\\n}\\n\\n.panel {\\n  display: block;\\n  height: 100%;\\n  overflow-y: auto;\\n  box-sizing: border-box;\\n  contain: layout style paint;\\n\\n  \\u0026[hidden] {\\n    display: none;\\n    content-visibility: hidden;\\n  }\\n}\\n\\n.tab {\\n  position: relative;\\n  display: flex;\\n  flex: none;\\n  align-items: center;\\n  padding: calc(var(--pf-t--global--spacer--sm) + var(--pf-t--global--spacer--xs))\\n           calc(var(--pf-t--global--spacer--sm) + var(--pf-t--global--spacer--sm));\\n  font-family: inherit;\\n  font-size: var(--pf-t--global--font--size--sm);\\n  font-weight: 400;\\n  color: var(--_tab--Color);\\n  text-decoration: none;\\n  background-color: transparent;\\n  border: 0;\\n  cursor: pointer;\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    inset-block-start: 0;\\n    inset-inline-start: 0;\\n    width: calc(100% - calc(var(--pf-t--global--spacer--sm) * 2));\\n    height: calc(100% - calc(var(--pf-t--global--spacer--sm) * 2));\\n    translate: var(--pf-t--global--spacer--sm) var(--pf-t--global--spacer--sm);\\n    background-color: var(--_tab--BackgroundColor);\\n    border: var(--_tab--BorderWidth) solid var(--_tab--BorderColor);\\n    border-radius: var(--pf-t--global--border--radius--small);\\n    transition: background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n    z-index: -1;\\n  }\\n\\n  \\u0026:where(:hover, :focus) {\\n    --_tab--Color: var(--_tab--hover--Color);\\n    --_tab--BackgroundColor: var(--_tab--hover--BackgroundColor);\\n  }\\n\\n  \\u0026[aria-selected=\\"true\\"] {\\n    --_tab--Color: var(--_tab--current--Color);\\n    --_tab--BackgroundColor: var(--_tab--current--BackgroundColor);\\n  }\\n\\n  \\u0026:disabled {\\n    --_tab--Color: var(--_tab--disabled--Color);\\n    --_tab--BackgroundColor: var(--_tab--disabled--BackgroundColor);\\n    cursor: not-allowed;\\n    pointer-events: none;\\n  }\\n}\\n"`));
var cem_pf_v6_tabs_default = s;

// elements/cem-pf-v6-tabs/cem-pf-v6-tabs.ts
import "../cem-pf-v6-tab/cem-pf-v6-tab.js";
var PfTabsChangeEvent = class extends Event {
  selectedIndex;
  constructor(selectedIndex) {
    super("change", { bubbles: true });
    this.selectedIndex = selectedIndex;
  }
};
var __tabs_dec, _selected_dec, _a, _PfV6Tabs_decorators, _init, _selected, __tabs, _mutationObserver, _onTabsChanged, _PfV6Tabs_instances, syncTabs_fn, handleTabClick_fn, handleKeyDown_fn, updateAccentLine_fn;
_PfV6Tabs_decorators = [customElement("cem-pf-v6-tabs")];
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
    this.addEventListener("cem-pf-v6-tab-title-changed", __privateGet(this, _onTabsChanged));
    this.addEventListener("cem-pf-v6-tab-connected", __privateGet(this, _onTabsChanged));
    this.addEventListener("cem-pf-v6-tab-disconnected", __privateGet(this, _onTabsChanged));
    __privateSet(this, _mutationObserver, new MutationObserver(__privateGet(this, _onTabsChanged)));
    __privateGet(this, _mutationObserver).observe(this, { childList: true });
  }
  firstUpdated() {
    __privateMethod(this, _PfV6Tabs_instances, syncTabs_fn).call(this);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(this, _mutationObserver)?.disconnect();
    this.removeEventListener("cem-pf-v6-tab-title-changed", __privateGet(this, _onTabsChanged));
    this.removeEventListener("cem-pf-v6-tab-connected", __privateGet(this, _onTabsChanged));
    this.removeEventListener("cem-pf-v6-tab-disconnected", __privateGet(this, _onTabsChanged));
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
  this._tabs = Array.from(this.querySelectorAll("cem-pf-v6-tab"));
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
  tabsEl.style.setProperty("--_cem-pf-v6-tabs--link-accent--start", `${start}px`);
  tabsEl.style.setProperty("--_cem-pf-v6-tabs--link-accent--length", `${length}px`);
};
__decorateElement(_init, 4, "selected", _selected_dec, PfV6Tabs, _selected);
__decorateElement(_init, 4, "_tabs", __tabs_dec, PfV6Tabs, __tabs);
PfV6Tabs = __decorateElement(_init, 0, "PfV6Tabs", _PfV6Tabs_decorators, PfV6Tabs);
__publicField(PfV6Tabs, "styles", cem_pf_v6_tabs_default);
__runInitializers(_init, 1, PfV6Tabs);
export {
  PfTabsChangeEvent,
  PfV6Tabs
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXRhYnMvY2VtLXBmLXY2LXRhYnMudHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tcGYtdjYtdGFicy9jZW0tcGYtdjYtdGFicy5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuaW1wb3J0IHsgc3RhdGUgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9zdGF0ZS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtdGFicy5jc3MnO1xuXG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10YWIvY2VtLXBmLXY2LXRhYi5qcyc7XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gdGFiIHNlbGVjdGlvbiBjaGFuZ2VzXG4gKi9cbmV4cG9ydCBjbGFzcyBQZlRhYnNDaGFuZ2VFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgc2VsZWN0ZWRJbmRleDogbnVtYmVyO1xuICBjb25zdHJ1Y3RvcihzZWxlY3RlZEluZGV4OiBudW1iZXIpIHtcbiAgICBzdXBlcignY2hhbmdlJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMuc2VsZWN0ZWRJbmRleCA9IHNlbGVjdGVkSW5kZXg7XG4gIH1cbn1cblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IFRhYnNcbiAqXG4gKiBBIHRhYmJlZCBjb250ZW50IGNvbnRhaW5lci4gQWRkIGBjZW0tcGYtdjYtdGFiYCBjaGlsZHJlbiB3aXRoIGB0aXRsZWAgYXR0cmlidXRlc1xuICogdG8gZGVmaW5lIHRhYnMgYW5kIHRoZWlyIHBhbmVsIGNvbnRlbnQuXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGBjZW0tcGYtdjYtdGFiYCBjaGlsZHJlblxuICpcbiAqIEBmaXJlcyBjaGFuZ2UgLSBGaXJlZCB3aGVuIHRoZSBzZWxlY3RlZCB0YWIgY2hhbmdlc1xuICpcbiAqIEBjc3NwYXJ0IHRhYnMgLSBUaGUgdGFiIGJ1dHRvbiBjb250YWluZXJcbiAqIEBjc3NwYXJ0IHBhbmVscyAtIFRoZSB0YWIgcGFuZWwgY29udGFpbmVyXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tcGYtdjYtdGFicycpXG5leHBvcnQgY2xhc3MgUGZWNlRhYnMgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBOdW1iZXIsIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3Igc2VsZWN0ZWQgPSAwO1xuXG4gIEBzdGF0ZSgpXG4gIGFjY2Vzc29yIF90YWJzOiBFbGVtZW50W10gPSBbXTtcblxuICAjbXV0YXRpb25PYnNlcnZlcj86IE11dGF0aW9uT2JzZXJ2ZXI7XG5cbiAgZ2V0IHNlbGVjdGVkSW5kZXgoKTogbnVtYmVyIHtcbiAgICByZXR1cm4gdGhpcy5zZWxlY3RlZDtcbiAgfVxuXG4gIHNldCBzZWxlY3RlZEluZGV4KGluZGV4OiBudW1iZXIpIHtcbiAgICB0aGlzLnNlbGVjdGVkID0gTWF0aC5tYXgoMCwgTWF0aC5taW4oaW5kZXgsIHRoaXMuX3RhYnMubGVuZ3RoIC0gMSkpO1xuICB9XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2NlbS1wZi12Ni10YWItdGl0bGUtY2hhbmdlZCcsIHRoaXMuI29uVGFic0NoYW5nZWQpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignY2VtLXBmLXY2LXRhYi1jb25uZWN0ZWQnLCB0aGlzLiNvblRhYnNDaGFuZ2VkKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2NlbS1wZi12Ni10YWItZGlzY29ubmVjdGVkJywgdGhpcy4jb25UYWJzQ2hhbmdlZCk7XG4gICAgdGhpcy4jbXV0YXRpb25PYnNlcnZlciA9IG5ldyBNdXRhdGlvbk9ic2VydmVyKHRoaXMuI29uVGFic0NoYW5nZWQpO1xuICAgIHRoaXMuI211dGF0aW9uT2JzZXJ2ZXIub2JzZXJ2ZSh0aGlzLCB7IGNoaWxkTGlzdDogdHJ1ZSB9KTtcbiAgfVxuXG4gIGZpcnN0VXBkYXRlZCgpIHtcbiAgICB0aGlzLiNzeW5jVGFicygpO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLiNtdXRhdGlvbk9ic2VydmVyPy5kaXNjb25uZWN0KCk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdjZW0tcGYtdjYtdGFiLXRpdGxlLWNoYW5nZWQnLCB0aGlzLiNvblRhYnNDaGFuZ2VkKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2NlbS1wZi12Ni10YWItY29ubmVjdGVkJywgdGhpcy4jb25UYWJzQ2hhbmdlZCk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdjZW0tcGYtdjYtdGFiLWRpc2Nvbm5lY3RlZCcsIHRoaXMuI29uVGFic0NoYW5nZWQpO1xuICB9XG5cbiAgI29uVGFic0NoYW5nZWQgPSAoKSA9PiB7XG4gICAgdGhpcy4jc3luY1RhYnMoKTtcbiAgfTtcblxuICAjc3luY1RhYnMoKSB7XG4gICAgdGhpcy5fdGFicyA9IEFycmF5LmZyb20odGhpcy5xdWVyeVNlbGVjdG9yQWxsKCdjZW0tcGYtdjYtdGFiJykpO1xuICAgIHRoaXMucmVxdWVzdFVwZGF0ZSgpO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygnc2VsZWN0ZWQnKSB8fCBjaGFuZ2VkLmhhcygnX3RhYnMnKSkge1xuICAgICAgLy8gQXNzaWduIHNsb3QgbmFtZXMgdG8gdGFiIGNoaWxkcmVuXG4gICAgICB0aGlzLl90YWJzLmZvckVhY2goKHRhYiwgaSkgPT4ge1xuICAgICAgICB0YWIuc2V0QXR0cmlidXRlKCdzbG90JywgYHBhbmVsLSR7aX1gKTtcbiAgICAgIH0pO1xuICAgICAgcmVxdWVzdEFuaW1hdGlvbkZyYW1lKCgpID0+IHRoaXMuI3VwZGF0ZUFjY2VudExpbmUoKSk7XG4gICAgfVxuICB9XG5cbiAgI2hhbmRsZVRhYkNsaWNrKGluZGV4OiBudW1iZXIpIHtcbiAgICBjb25zdCBvbGRJbmRleCA9IHRoaXMuc2VsZWN0ZWQ7XG4gICAgdGhpcy5zZWxlY3RlZCA9IGluZGV4O1xuICAgIGlmIChvbGRJbmRleCAhPT0gaW5kZXgpIHtcbiAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgUGZUYWJzQ2hhbmdlRXZlbnQoaW5kZXgpKTtcbiAgICB9XG4gIH1cblxuICAjaGFuZGxlS2V5RG93bihlOiBLZXlib2FyZEV2ZW50KSB7XG4gICAgY29uc3QgdGFic0VsID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgndGFicycpO1xuICAgIGlmICghdGFic0VsKSByZXR1cm47XG4gICAgY29uc3QgYnV0dG9ucyA9IEFycmF5LmZyb20odGFic0VsLnF1ZXJ5U2VsZWN0b3JBbGwoJy50YWInKSkgYXMgSFRNTEJ1dHRvbkVsZW1lbnRbXTtcbiAgICBjb25zdCBjdXJyZW50SW5kZXggPSBidXR0b25zLmZpbmRJbmRleChidG4gPT4gYnRuID09PSB0aGlzLnNoYWRvd1Jvb3Q/LmFjdGl2ZUVsZW1lbnQpO1xuICAgIGlmIChjdXJyZW50SW5kZXggPT09IC0xKSByZXR1cm47XG5cbiAgICBsZXQgbmV4dEluZGV4ID0gY3VycmVudEluZGV4O1xuXG4gICAgc3dpdGNoIChlLmtleSkge1xuICAgICAgY2FzZSAnQXJyb3dMZWZ0JzpcbiAgICAgICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBuZXh0SW5kZXggPSBjdXJyZW50SW5kZXggPiAwID8gY3VycmVudEluZGV4IC0gMSA6IGJ1dHRvbnMubGVuZ3RoIC0gMTtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdBcnJvd1JpZ2h0JzpcbiAgICAgICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBuZXh0SW5kZXggPSBjdXJyZW50SW5kZXggPCBidXR0b25zLmxlbmd0aCAtIDEgPyBjdXJyZW50SW5kZXggKyAxIDogMDtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdIb21lJzpcbiAgICAgICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBuZXh0SW5kZXggPSAwO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ0VuZCc6XG4gICAgICAgIGUucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgbmV4dEluZGV4ID0gYnV0dG9ucy5sZW5ndGggLSAxO1xuICAgICAgICBicmVhaztcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICB0aGlzLiNoYW5kbGVUYWJDbGljayhuZXh0SW5kZXgpO1xuICAgIGJ1dHRvbnNbbmV4dEluZGV4XS5mb2N1cygpO1xuICB9XG5cbiAgI3VwZGF0ZUFjY2VudExpbmUoKSB7XG4gICAgY29uc3QgdGFic0VsID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgndGFicycpO1xuICAgIGlmICghdGFic0VsKSByZXR1cm47XG4gICAgY29uc3QgYnV0dG9ucyA9IEFycmF5LmZyb20odGFic0VsLnF1ZXJ5U2VsZWN0b3JBbGwoJy50YWInKSk7XG4gICAgY29uc3QgYWN0aXZlQnV0dG9uID0gYnV0dG9uc1t0aGlzLnNlbGVjdGVkXSBhcyBIVE1MRWxlbWVudCB8IHVuZGVmaW5lZDtcbiAgICBpZiAoIWFjdGl2ZUJ1dHRvbikgcmV0dXJuO1xuXG4gICAgY29uc3QgY29udGFpbmVyUmVjdCA9IHRhYnNFbC5nZXRCb3VuZGluZ0NsaWVudFJlY3QoKTtcbiAgICBjb25zdCBidXR0b25SZWN0ID0gYWN0aXZlQnV0dG9uLmdldEJvdW5kaW5nQ2xpZW50UmVjdCgpO1xuICAgIGNvbnN0IGluc2V0ID0gODsgLy8gdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKVxuICAgIGNvbnN0IHN0YXJ0ID0gKGJ1dHRvblJlY3QubGVmdCAtIGNvbnRhaW5lclJlY3QubGVmdCkgKyBpbnNldDtcbiAgICBjb25zdCBsZW5ndGggPSBidXR0b25SZWN0LndpZHRoIC0gKGluc2V0ICogMik7XG5cbiAgICB0YWJzRWwuc3R5bGUuc2V0UHJvcGVydHkoJy0tX2NlbS1wZi12Ni10YWJzLS1saW5rLWFjY2VudC0tc3RhcnQnLCBgJHtzdGFydH1weGApO1xuICAgIHRhYnNFbC5zdHlsZS5zZXRQcm9wZXJ0eSgnLS1fY2VtLXBmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1sZW5ndGgnLCBgJHtsZW5ndGh9cHhgKTtcbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJ0YWJzLWNvbnRhaW5lclwiPlxuICAgICAgICA8ZGl2IGlkPVwidGFic1wiXG4gICAgICAgICAgICAgcm9sZT1cInRhYmxpc3RcIlxuICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJUYWJzXCJcbiAgICAgICAgICAgICBwYXJ0PVwidGFic1wiXG4gICAgICAgICAgICAgQGtleWRvd249JHt0aGlzLiNoYW5kbGVLZXlEb3dufT5cbiAgICAgICAgICAke3RoaXMuX3RhYnMubWFwKCh0YWIsIGluZGV4KSA9PiBodG1sYFxuICAgICAgICAgICAgPGJ1dHRvbiBjbGFzcz1cInRhYlwiXG4gICAgICAgICAgICAgICAgICAgIHR5cGU9XCJidXR0b25cIlxuICAgICAgICAgICAgICAgICAgICByb2xlPVwidGFiXCJcbiAgICAgICAgICAgICAgICAgICAgaWQ9XCJ0YWItJHtpbmRleH1cIlxuICAgICAgICAgICAgICAgICAgICBhcmlhLWNvbnRyb2xzPVwicGFuZWwtJHtpbmRleH1cIlxuICAgICAgICAgICAgICAgICAgICBhcmlhLXNlbGVjdGVkPVwiJHtpbmRleCA9PT0gdGhpcy5zZWxlY3RlZCA/ICd0cnVlJyA6ICdmYWxzZSd9XCJcbiAgICAgICAgICAgICAgICAgICAgdGFiaW5kZXg9XCIke2luZGV4ID09PSB0aGlzLnNlbGVjdGVkID8gMCA6IC0xfVwiXG4gICAgICAgICAgICAgICAgICAgIEBjbGljaz0keygpID0+IHRoaXMuI2hhbmRsZVRhYkNsaWNrKGluZGV4KX0+XG4gICAgICAgICAgICAgICR7KHRhYiBhcyBIVE1MRWxlbWVudCkudGl0bGUgfHwgYFRhYiAke2luZGV4ICsgMX1gfVxuICAgICAgICAgICAgPC9idXR0b24+XG4gICAgICAgICAgYCl9XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGlkPVwicGFuZWxzXCJcbiAgICAgICAgICAgICBwYXJ0PVwicGFuZWxzXCI+XG4gICAgICAgICAgJHt0aGlzLl90YWJzLm1hcCgoXywgaW5kZXgpID0+IGh0bWxgXG4gICAgICAgICAgICA8ZGl2IGNsYXNzPVwicGFuZWxcIlxuICAgICAgICAgICAgICAgICBpZD1cInBhbmVsLSR7aW5kZXh9XCJcbiAgICAgICAgICAgICAgICAgcm9sZT1cInRhYnBhbmVsXCJcbiAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbGxlZGJ5PVwidGFiLSR7aW5kZXh9XCJcbiAgICAgICAgICAgICAgICAgP2hpZGRlbj0ke2luZGV4ICE9PSB0aGlzLnNlbGVjdGVkfT5cbiAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cInBhbmVsLSR7aW5kZXh9XCI+PC9zbG90PlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgYCl9XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcGYtdjYtdGFicyc6IFBmVjZUYWJzO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbn1cXFxcblxcXFxuI3RhYnMtY29udGFpbmVyIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG5AcHJvcGVydHkgLS1fY2VtLXBmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1sZW5ndGgge1xcXFxuICBzeW50YXg6IFxcXFxcXFwiXFxcXHUwMDNjbGVuZ3RoXFxcXHUwMDNlXFxcXFxcXCI7XFxcXG4gIGluaGVyaXRzOiB0cnVlO1xcXFxuICBpbml0aWFsLXZhbHVlOiAwcHg7XFxcXG59XFxcXG5cXFxcbiN0YWJzIHtcXFxcbiAgLS1fdGFiLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgLS1fdGFiLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWRlZmF1bHQpO1xcXFxuICAtLV90YWItLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWFjdGlvbi0tcGxhaW4tLWhvdmVyKTtcXFxcbiAgLS1fdGFiLS1Cb3JkZXJDb2xvcjogdHJhbnNwYXJlbnQ7XFxcXG4gIC0tX3RhYi0taG92ZXItLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1fdGFiLS1ob3Zlci0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1ob3Zlcik7XFxcXG4gIC0tX3RhYi0tY3VycmVudC0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAtLV90YWItLWN1cnJlbnQtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0tZGVmYXVsdCk7XFxcXG4gIC0tX3RhYi0tZGlzYWJsZWQtLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1kaXNhYmxlZCk7XFxcXG4gIC0tX3RhYi0tZGlzYWJsZWQtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tZGlzYWJsZWQtLWRlZmF1bHQpO1xcXFxuXFxcXG4gIC0tX2NlbS1wZi12Ni10YWJzLS1saW5rLWFjY2VudC0tc3RhcnQ6IDBweDtcXFxcbiAgLS1fY2VtLXBmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1sZW5ndGg6IDBweDtcXFxcbiAgLS1fY2VtLXBmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1jbGlja2VkKTtcXFxcbiAgLS1fY2VtLXBmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1ib3JkZXItc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1leHRyYS1zdHJvbmcpO1xcXFxuXFxcXG4gIHBvc2l0aW9uOiBzdGlja3k7XFxcXG4gIGluc2V0LWJsb2NrLXN0YXJ0OiAwO1xcXFxuICB6LWluZGV4OiB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLW1kKTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgb3ZlcmZsb3cteDogYXV0bztcXFxcbiAgb3ZlcmZsb3cteTogaGlkZGVuO1xcXFxuICAtd2Via2l0LW92ZXJmbG93LXNjcm9sbGluZzogdG91Y2g7XFxcXG4gIHNjcm9sbC1iZWhhdmlvcjogc21vb3RoO1xcXFxuICBmbGV4LXNocmluazogMDtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0tZGVmYXVsdCk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OjpiZWZvcmUge1xcXFxuICAgIGNvbnRlbnQ6ICcnO1xcXFxuICAgIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgICBpbnNldC1ibG9jay1lbmQ6IDA7XFxcXG4gICAgaW5zZXQtaW5saW5lLXN0YXJ0OiAwO1xcXFxuICAgIGluc2V0LWlubGluZS1lbmQ6IDA7XFxcXG4gICAgYm9yZGVyLWJsb2NrLWVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1yZWd1bGFyKSBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OjphZnRlciB7XFxcXG4gICAgY29udGVudDogJyc7XFxcXG4gICAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICAgIGluc2V0LWJsb2NrLXN0YXJ0OiBhdXRvO1xcXFxuICAgIGluc2V0LWJsb2NrLWVuZDogMDtcXFxcbiAgICBpbnNldC1pbmxpbmUtc3RhcnQ6IDA7XFxcXG4gICAgd2lkdGg6IHZhcigtLV9jZW0tcGYtdjYtdGFicy0tbGluay1hY2NlbnQtLWxlbmd0aCk7XFxcXG4gICAgaGVpZ2h0OiAwO1xcXFxuICAgIGJvcmRlcjogMCBzb2xpZCB2YXIoLS1fY2VtLXBmLXY2LXRhYnMtLWxpbmstYWNjZW50LS1jb2xvcik7XFxcXG4gICAgYm9yZGVyLWJsb2NrLWVuZC13aWR0aDogdmFyKC0tX2NlbS1wZi12Ni10YWJzLS1saW5rLWFjY2VudC0tYm9yZGVyLXNpemUpO1xcXFxuICAgIHRyYW5zZm9ybS1vcmlnaW46IDAgY2VudGVyO1xcXFxuICAgIHRyYW5zbGF0ZTogdmFyKC0tX2NlbS1wZi12Ni10YWJzLS1saW5rLWFjY2VudC0tc3RhcnQpIDA7XFxcXG4gICAgdHJhbnNpdGlvbi1wcm9wZXJ0eTogd2lkdGgsIHRyYW5zbGF0ZTtcXFxcbiAgICB0cmFuc2l0aW9uLWR1cmF0aW9uOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLW1kKTtcXFxcbiAgICB0cmFuc2l0aW9uLXRpbWluZy1mdW5jdGlvbjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVjZWxlcmF0ZSk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI3BhbmVscyB7XFxcXG4gIGZsZXg6IDE7XFxcXG4gIG1pbi1oZWlnaHQ6IDA7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG59XFxcXG5cXFxcbi5wYW5lbCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG4gIGJveC1zaXppbmc6IGJvcmRlci1ib3g7XFxcXG4gIGNvbnRhaW46IGxheW91dCBzdHlsZSBwYWludDtcXFxcblxcXFxuICBcXFxcdTAwMjZbaGlkZGVuXSB7XFxcXG4gICAgZGlzcGxheTogbm9uZTtcXFxcbiAgICBjb250ZW50LXZpc2liaWxpdHk6IGhpZGRlbjtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4udGFiIHtcXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4OiBub25lO1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBwYWRkaW5nOiBjYWxjKHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSkgKyB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teHMpKVxcXFxuICAgICAgICAgICBjYWxjKHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSkgKyB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pKTtcXFxcbiAgZm9udC1mYW1pbHk6IGluaGVyaXQ7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1zbSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA0MDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1fdGFiLS1Db2xvcik7XFxcXG4gIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdHJhbnNwYXJlbnQ7XFxcXG4gIGJvcmRlcjogMDtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjo6YmVmb3JlIHtcXFxcbiAgICBjb250ZW50OiAnJztcXFxcbiAgICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gICAgaW5zZXQtYmxvY2stc3RhcnQ6IDA7XFxcXG4gICAgaW5zZXQtaW5saW5lLXN0YXJ0OiAwO1xcXFxuICAgIHdpZHRoOiBjYWxjKDEwMCUgLSBjYWxjKHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSkgKiAyKSk7XFxcXG4gICAgaGVpZ2h0OiBjYWxjKDEwMCUgLSBjYWxjKHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSkgKiAyKSk7XFxcXG4gICAgdHJhbnNsYXRlOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tX3RhYi0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgICBib3JkZXI6IHZhcigtLV90YWItLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1fdGFiLS1Cb3JkZXJDb2xvcik7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICAgIHRyYW5zaXRpb246IGJhY2tncm91bmQtY29sb3IgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1zaG9ydCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCk7XFxcXG4gICAgei1pbmRleDogLTE7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6d2hlcmUoOmhvdmVyLCA6Zm9jdXMpIHtcXFxcbiAgICAtLV90YWItLUNvbG9yOiB2YXIoLS1fdGFiLS1ob3Zlci0tQ29sb3IpO1xcXFxuICAgIC0tX3RhYi0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1fdGFiLS1ob3Zlci0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNlthcmlhLXNlbGVjdGVkPVxcXFxcXFwidHJ1ZVxcXFxcXFwiXSB7XFxcXG4gICAgLS1fdGFiLS1Db2xvcjogdmFyKC0tX3RhYi0tY3VycmVudC0tQ29sb3IpO1xcXFxuICAgIC0tX3RhYi0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1fdGFiLS1jdXJyZW50LS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmRpc2FibGVkIHtcXFxcbiAgICAtLV90YWItLUNvbG9yOiB2YXIoLS1fdGFiLS1kaXNhYmxlZC0tQ29sb3IpO1xcXFxuICAgIC0tX3RhYi0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1fdGFiLS1kaXNhYmxlZC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgICBjdXJzb3I6IG5vdC1hbGxvd2VkO1xcXFxuICAgIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCO0FBQ3pCLFNBQVMsYUFBYTs7O0FDSHRCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLDhwSkFBb3FKLENBQUM7QUFBRSxJQUFPLHlCQUFROzs7QURP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
