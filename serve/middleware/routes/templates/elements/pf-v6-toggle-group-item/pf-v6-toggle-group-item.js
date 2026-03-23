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

// elements/pf-v6-toggle-group-item/pf-v6-toggle-group-item.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-toggle-group-item/pf-v6-toggle-group-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  position: relative;\\n  z-index: var(--pf-v6-c-toggle-group__button--ZIndex);\\n  display: inline-flex;\\n  align-items: center;\\n  padding-block-start: var(--pf-v6-c-toggle-group__button--PaddingBlockStart);\\n  padding-block-end: var(--pf-v6-c-toggle-group__button--PaddingBlockEnd);\\n  padding-inline-start: var(--pf-v6-c-toggle-group__button--PaddingInlineStart);\\n  padding-inline-end: var(--pf-v6-c-toggle-group__button--PaddingInlineEnd);\\n  font-size: var(--pf-v6-c-toggle-group__button--FontSize);\\n  line-height: var(--pf-v6-c-toggle-group__button--LineHeight);\\n  color: var(--pf-v6-c-toggle-group__button--Color);\\n  background-color: var(--pf-v6-c-toggle-group__button--BackgroundColor);\\n  border: 0;\\n  cursor: pointer;\\n  user-select: none;\\n}\\n\\n:host(:not(:first-child)) {\\n  margin-inline-start: var(--pf-v6-c-toggle-group__item--item--MarginInlineStart);\\n}\\n\\n:host::before,\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border-style: solid;\\n  border-radius: inherit;\\n}\\n\\n:host::before {\\n  border-width: var(--pf-v6-c-toggle-group__button--before--BorderWidth);\\n  border-block-start-color: var(--pf-v6-c-toggle-group__button--before--BorderBlockStartColor, var(--pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-block-end-color: var(--pf-v6-c-toggle-group__button--before--BorderBlockEndColor, var(--pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-inline-start-color: var(--pf-v6-c-toggle-group__button--before--BorderInlineStartColor, var(--pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-inline-end-color: var(--pf-v6-c-toggle-group__button--before--BorderInlineEndColor, var(--pf-v6-c-toggle-group__button--before--BorderColor));\\n}\\n\\n:host::after {\\n  inset: var(--pf-v6-c-toggle-group__button--before--BorderWidth);\\n  border-color: var(--pf-v6-c-toggle-group__button--after--BorderColor);\\n  border-width: var(--pf-v6-c-toggle-group__button--after--BorderWidth);\\n}\\n\\n:host(:hover:not([disabled])) {\\n  --pf-v6-c-toggle-group__button--BackgroundColor: var(--pf-v6-c-toggle-group__button--hover--BackgroundColor);\\n  --pf-v6-c-toggle-group__button--ZIndex: var(--pf-v6-c-toggle-group__button--hover--ZIndex);\\n  --pf-v6-c-toggle-group__button--before--BorderColor: var(--pf-v6-c-toggle-group__button--hover--before--BorderColor);\\n  --pf-v6-c-toggle-group__button--after--BorderWidth: var(--pf-v6-c-toggle-group__button--hover--after--BorderWidth);\\n}\\n\\n:host(:focus-visible) {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: 2px;\\n  --pf-v6-c-toggle-group__button--ZIndex: var(--pf-v6-c-toggle-group__button--hover--ZIndex);\\n}\\n\\n:host([selected]) {\\n  --pf-v6-c-toggle-group__button--BackgroundColor: var(--pf-v6-c-toggle-group__button--m-selected--BackgroundColor);\\n  --pf-v6-c-toggle-group__button--Color: var(--pf-v6-c-toggle-group__button--m-selected--Color);\\n  --pf-v6-c-toggle-group__button--before--BorderColor: var(--pf-v6-c-toggle-group__button--m-selected--before--BorderColor);\\n  --pf-v6-c-toggle-group__button--ZIndex: var(--pf-v6-c-toggle-group__button--m-selected--ZIndex);\\n  --pf-v6-c-toggle-group__button--after--BorderWidth: var(--pf-v6-c-toggle-group__button--m-selected--after--BorderWidth);\\n}\\n\\n:host([selected]) + :host([selected]) {\\n  --pf-v6-c-toggle-group__button--before--BorderInlineStartColor: var(--pf-v6-c-toggle-group__button--m-selected-selected--before--BorderInlineStartColor);\\n}\\n\\n:host([disabled]) {\\n  --pf-v6-c-toggle-group__button--BackgroundColor: var(--pf-v6-c-toggle-group__button--disabled--BackgroundColor);\\n  --pf-v6-c-toggle-group__button--Color: var(--pf-v6-c-toggle-group__button--disabled--Color);\\n  --pf-v6-c-toggle-group__button--before--BorderColor: var(--pf-v6-c-toggle-group__button--disabled--before--BorderColor);\\n  --pf-v6-c-toggle-group__button--ZIndex: var(--pf-v6-c-toggle-group__button--disabled--ZIndex);\\n  cursor: not-allowed;\\n  pointer-events: none;\\n}\\n\\n:host([disabled]) + :host([disabled]) {\\n  --pf-v6-c-toggle-group__button--before--BorderInlineStartColor: var(--pf-v6-c-toggle-group__button--disabled-disabled--before--BorderInlineStartColor);\\n}\\n\\n:host(:first-child) {\\n  border-start-start-radius: var(--pf-v6-c-toggle-group__item--first-child__button--BorderStartStartRadius);\\n  border-end-start-radius: var(--pf-v6-c-toggle-group__item--first-child__button--BorderEndStartRadius);\\n}\\n\\n:host(:last-child) {\\n  border-start-end-radius: var(--pf-v6-c-toggle-group__item--last-child__button--BorderStartEndRadius);\\n  border-end-end-radius: var(--pf-v6-c-toggle-group__item--last-child__button--BorderEndEndRadius);\\n}\\n\\n#wrapper {\\n  display: var(--_has-content, inline-flex);\\n  gap: var(--pf-v6-c-toggle-group__icon--text--MarginInlineStart);\\n  align-items: center;\\n}\\n\\n::slotted(svg) {\\n  width: 1em;\\n  height: 1em;\\n  vertical-align: -0.125em;\\n  fill: currentColor;\\n}\\n"'));
var pf_v6_toggle_group_item_default = s;

// elements/pf-v6-toggle-group-item/pf-v6-toggle-group-item.ts
var ToggleGroupItemSelectEvent = class extends Event {
  item;
  selected;
  value;
  constructor(item, selected, value) {
    super("pf-v6-toggle-group-item-select", { bubbles: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
};
var _value_dec, _disabled_dec, _selected_dec, _a, _PfV6ToggleGroupItem_decorators, _internals, _init, _selected, _disabled, _value, _PfV6ToggleGroupItem_instances, updateTabindex_fn, _handleClick, _handleKeydown, _handleFocus, selectItem_fn, focusAndSelect_fn, navigateItems_fn, navigateToEnd_fn, updateRovingTabindex_fn;
_PfV6ToggleGroupItem_decorators = [customElement("pf-v6-toggle-group-item")];
var PfV6ToggleGroupItem = class extends (_a = LitElement, _selected_dec = [property({ type: Boolean, reflect: true })], _disabled_dec = [property({ type: Boolean, reflect: true })], _value_dec = [property({ reflect: true })], _a) {
  constructor() {
    super();
    __privateAdd(this, _PfV6ToggleGroupItem_instances);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _selected, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _disabled, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateAdd(this, _value, __runInitializers(_init, 16, this)), __runInitializers(_init, 19, this);
    __privateAdd(this, _handleClick, (event) => {
      if (this.disabled) {
        event.preventDefault();
        event.stopImmediatePropagation();
        return;
      }
      this.focus();
      __privateMethod(this, _PfV6ToggleGroupItem_instances, selectItem_fn).call(this);
    });
    __privateAdd(this, _handleKeydown, (event) => {
      if (this.disabled) {
        event.preventDefault();
        return;
      }
      switch (event.key) {
        case " ":
        case "Enter":
          event.preventDefault();
          __privateMethod(this, _PfV6ToggleGroupItem_instances, selectItem_fn).call(this);
          break;
        case "ArrowLeft":
        case "ArrowRight":
          event.preventDefault();
          __privateMethod(this, _PfV6ToggleGroupItem_instances, navigateItems_fn).call(this, event.key === "ArrowLeft" ? -1 : 1);
          break;
        case "Home":
          event.preventDefault();
          __privateMethod(this, _PfV6ToggleGroupItem_instances, navigateToEnd_fn).call(this, true);
          break;
        case "End":
          event.preventDefault();
          __privateMethod(this, _PfV6ToggleGroupItem_instances, navigateToEnd_fn).call(this, false);
          break;
      }
    });
    __privateAdd(this, _handleFocus, () => {
      __privateMethod(this, _PfV6ToggleGroupItem_instances, updateRovingTabindex_fn).call(this);
    });
    __privateGet(this, _internals).role = "radio";
    if (!this.hasAttribute("tabindex")) {
      this.setAttribute("tabindex", "-1");
    }
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("click", __privateGet(this, _handleClick));
    this.addEventListener("keydown", __privateGet(this, _handleKeydown));
    this.addEventListener("focus", __privateGet(this, _handleFocus));
    queueMicrotask(() => {
      const parent = this.parentElement;
      if (parent) {
        const items = Array.from(parent.querySelectorAll("pf-v6-toggle-group-item"));
        const isFirstItem = items[0] === this;
        if (this.selected || isFirstItem) {
          __privateMethod(this, _PfV6ToggleGroupItem_instances, updateRovingTabindex_fn).call(this);
        }
      }
    });
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("click", __privateGet(this, _handleClick));
    this.removeEventListener("keydown", __privateGet(this, _handleKeydown));
    this.removeEventListener("focus", __privateGet(this, _handleFocus));
  }
  updated(changed) {
    if (changed.has("selected")) {
      __privateGet(this, _internals).ariaChecked = String(this.selected);
      __privateMethod(this, _PfV6ToggleGroupItem_instances, updateRovingTabindex_fn).call(this);
    }
    if (changed.has("disabled")) {
      __privateGet(this, _internals).ariaDisabled = String(this.disabled);
      __privateMethod(this, _PfV6ToggleGroupItem_instances, updateTabindex_fn).call(this);
    }
  }
  render() {
    return html`
      <span id="wrapper">
        <span id="icon-start"><slot name="icon"></slot></span>
        <span id="text"><slot></slot></span>
        <span id="icon-end"><slot name="icon-end"></slot></span>
      </span>
    `;
  }
};
_init = __decoratorStart(_a);
_internals = new WeakMap();
_selected = new WeakMap();
_disabled = new WeakMap();
_value = new WeakMap();
_PfV6ToggleGroupItem_instances = new WeakSet();
updateTabindex_fn = function() {
  if (this.disabled) {
    this.setAttribute("tabindex", "-1");
    return;
  }
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("pf-v6-toggle-group-item"));
  const isFirstItem = items[0] === this;
  const hasSelectedItem = items.some((item) => item.hasAttribute("selected"));
  const shouldBeTabbable = this.selected || isFirstItem && !hasSelectedItem;
  this.setAttribute("tabindex", shouldBeTabbable ? "0" : "-1");
};
_handleClick = new WeakMap();
_handleKeydown = new WeakMap();
_handleFocus = new WeakMap();
selectItem_fn = function() {
  if (!this.selected) {
    this.selected = true;
    this.dispatchEvent(new ToggleGroupItemSelectEvent(
      this,
      true,
      this.getAttribute("value")
    ));
  }
};
focusAndSelect_fn = function(item) {
  item.focus();
  item.selected = true;
  item.dispatchEvent(new ToggleGroupItemSelectEvent(
    item,
    true,
    item.getAttribute("value")
  ));
};
navigateItems_fn = function(direction) {
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("pf-v6-toggle-group-item")).filter((item) => !item.disabled);
  const currentIndex = items.indexOf(this);
  if (currentIndex === -1) return;
  let newIndex = currentIndex + direction;
  if (newIndex < 0) newIndex = items.length - 1;
  else if (newIndex >= items.length) newIndex = 0;
  const targetItem = items[newIndex];
  if (targetItem) __privateMethod(this, _PfV6ToggleGroupItem_instances, focusAndSelect_fn).call(this, targetItem);
};
navigateToEnd_fn = function(toStart) {
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("pf-v6-toggle-group-item")).filter((item) => !item.disabled);
  const targetItem = toStart ? items[0] : items[items.length - 1];
  if (targetItem) __privateMethod(this, _PfV6ToggleGroupItem_instances, focusAndSelect_fn).call(this, targetItem);
};
updateRovingTabindex_fn = function() {
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("pf-v6-toggle-group-item"));
  items.forEach((item) => {
    item.setAttribute("tabindex", item === this ? "0" : "-1");
  });
};
__decorateElement(_init, 4, "selected", _selected_dec, PfV6ToggleGroupItem, _selected);
__decorateElement(_init, 4, "disabled", _disabled_dec, PfV6ToggleGroupItem, _disabled);
__decorateElement(_init, 4, "value", _value_dec, PfV6ToggleGroupItem, _value);
PfV6ToggleGroupItem = __decorateElement(_init, 0, "PfV6ToggleGroupItem", _PfV6ToggleGroupItem_decorators, PfV6ToggleGroupItem);
__publicField(PfV6ToggleGroupItem, "styles", pf_v6_toggle_group_item_default);
__runInitializers(_init, 1, PfV6ToggleGroupItem);
export {
  PfV6ToggleGroupItem,
  ToggleGroupItemSelectEvent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0vcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0udHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0uY3NzJztcblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZm9yIHRvZ2dsZSBncm91cCBpdGVtIHNlbGVjdGlvblxuICovXG5leHBvcnQgY2xhc3MgVG9nZ2xlR3JvdXBJdGVtU2VsZWN0RXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIGl0ZW06IEVsZW1lbnQ7XG4gIHNlbGVjdGVkOiBib29sZWFuO1xuICB2YWx1ZTogc3RyaW5nIHwgbnVsbDtcbiAgY29uc3RydWN0b3IoaXRlbTogRWxlbWVudCwgc2VsZWN0ZWQ6IGJvb2xlYW4sIHZhbHVlOiBzdHJpbmcgfCBudWxsKSB7XG4gICAgc3VwZXIoJ3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtLXNlbGVjdCcsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLml0ZW0gPSBpdGVtO1xuICAgIHRoaXMuc2VsZWN0ZWQgPSBzZWxlY3RlZDtcbiAgICB0aGlzLnZhbHVlID0gdmFsdWU7XG4gIH1cbn1cblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IFRvZ2dsZSBHcm91cCBJdGVtXG4gKlxuICogQW4gaW5kaXZpZHVhbCB0b2dnbGUgYnV0dG9uIHdpdGhpbiBhIHRvZ2dsZSBncm91cC5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgYnV0dG9uIHRleHRcbiAqIEBzbG90IGljb24gLSBMZWFkaW5nIGljb25cbiAqIEBzbG90IGljb24tZW5kIC0gVHJhaWxpbmcgaWNvblxuICpcbiAqIEBmaXJlcyBwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS1zZWxlY3QgLSBGaXJlcyB3aGVuIGl0ZW0gaXMgc2VsZWN0ZWRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJylcbmV4cG9ydCBjbGFzcyBQZlY2VG9nZ2xlR3JvdXBJdGVtIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgI2ludGVybmFscyA9IHRoaXMuYXR0YWNoSW50ZXJuYWxzKCk7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBzZWxlY3RlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZGlzYWJsZWQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIHZhbHVlPzogc3RyaW5nO1xuXG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCk7XG4gICAgdGhpcy4jaW50ZXJuYWxzLnJvbGUgPSAncmFkaW8nO1xuICAgIGlmICghdGhpcy5oYXNBdHRyaWJ1dGUoJ3RhYmluZGV4JykpIHtcbiAgICAgIHRoaXMuc2V0QXR0cmlidXRlKCd0YWJpbmRleCcsICctMScpO1xuICAgIH1cbiAgfVxuXG4gIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmNvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsIHRoaXMuI2hhbmRsZUNsaWNrKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNoYW5kbGVLZXlkb3duKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2ZvY3VzJywgdGhpcy4jaGFuZGxlRm9jdXMpO1xuXG4gICAgcXVldWVNaWNyb3Rhc2soKCkgPT4ge1xuICAgICAgY29uc3QgcGFyZW50ID0gdGhpcy5wYXJlbnRFbGVtZW50O1xuICAgICAgaWYgKHBhcmVudCkge1xuICAgICAgICBjb25zdCBpdGVtcyA9IEFycmF5LmZyb20ocGFyZW50LnF1ZXJ5U2VsZWN0b3JBbGwoJ3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJykpO1xuICAgICAgICBjb25zdCBpc0ZpcnN0SXRlbSA9IGl0ZW1zWzBdID09PSB0aGlzO1xuICAgICAgICBpZiAodGhpcy5zZWxlY3RlZCB8fCBpc0ZpcnN0SXRlbSkge1xuICAgICAgICAgIHRoaXMuI3VwZGF0ZVJvdmluZ1RhYmluZGV4KCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdjbGljaycsIHRoaXMuI2hhbmRsZUNsaWNrKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNoYW5kbGVLZXlkb3duKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2ZvY3VzJywgdGhpcy4jaGFuZGxlRm9jdXMpO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygnc2VsZWN0ZWQnKSkge1xuICAgICAgdGhpcy4jaW50ZXJuYWxzLmFyaWFDaGVja2VkID0gU3RyaW5nKHRoaXMuc2VsZWN0ZWQpO1xuICAgICAgdGhpcy4jdXBkYXRlUm92aW5nVGFiaW5kZXgoKTtcbiAgICB9XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdkaXNhYmxlZCcpKSB7XG4gICAgICB0aGlzLiNpbnRlcm5hbHMuYXJpYURpc2FibGVkID0gU3RyaW5nKHRoaXMuZGlzYWJsZWQpO1xuICAgICAgdGhpcy4jdXBkYXRlVGFiaW5kZXgoKTtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8c3BhbiBpZD1cIndyYXBwZXJcIj5cbiAgICAgICAgPHNwYW4gaWQ9XCJpY29uLXN0YXJ0XCI+PHNsb3QgbmFtZT1cImljb25cIj48L3Nsb3Q+PC9zcGFuPlxuICAgICAgICA8c3BhbiBpZD1cInRleHRcIj48c2xvdD48L3Nsb3Q+PC9zcGFuPlxuICAgICAgICA8c3BhbiBpZD1cImljb24tZW5kXCI+PHNsb3QgbmFtZT1cImljb24tZW5kXCI+PC9zbG90Pjwvc3Bhbj5cbiAgICAgIDwvc3Bhbj5cbiAgICBgO1xuICB9XG5cbiAgI3VwZGF0ZVRhYmluZGV4KCkge1xuICAgIGlmICh0aGlzLmRpc2FibGVkKSB7XG4gICAgICB0aGlzLnNldEF0dHJpYnV0ZSgndGFiaW5kZXgnLCAnLTEnKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG4gICAgY29uc3QgcGFyZW50ID0gdGhpcy5wYXJlbnRFbGVtZW50O1xuICAgIGlmICghcGFyZW50KSByZXR1cm47XG4gICAgY29uc3QgaXRlbXMgPSBBcnJheS5mcm9tKHBhcmVudC5xdWVyeVNlbGVjdG9yQWxsKCdwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbScpKTtcbiAgICBjb25zdCBpc0ZpcnN0SXRlbSA9IGl0ZW1zWzBdID09PSB0aGlzO1xuICAgIGNvbnN0IGhhc1NlbGVjdGVkSXRlbSA9IGl0ZW1zLnNvbWUoaXRlbSA9PiBpdGVtLmhhc0F0dHJpYnV0ZSgnc2VsZWN0ZWQnKSk7XG4gICAgY29uc3Qgc2hvdWxkQmVUYWJiYWJsZSA9IHRoaXMuc2VsZWN0ZWQgfHwgKGlzRmlyc3RJdGVtICYmICFoYXNTZWxlY3RlZEl0ZW0pO1xuICAgIHRoaXMuc2V0QXR0cmlidXRlKCd0YWJpbmRleCcsIHNob3VsZEJlVGFiYmFibGUgPyAnMCcgOiAnLTEnKTtcbiAgfVxuXG4gICNoYW5kbGVDbGljayA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBpZiAodGhpcy5kaXNhYmxlZCkge1xuICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgIGV2ZW50LnN0b3BJbW1lZGlhdGVQcm9wYWdhdGlvbigpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICB0aGlzLmZvY3VzKCk7XG4gICAgdGhpcy4jc2VsZWN0SXRlbSgpO1xuICB9O1xuXG4gICNoYW5kbGVLZXlkb3duID0gKGV2ZW50OiBLZXlib2FyZEV2ZW50KSA9PiB7XG4gICAgaWYgKHRoaXMuZGlzYWJsZWQpIHtcbiAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICByZXR1cm47XG4gICAgfVxuICAgIHN3aXRjaCAoZXZlbnQua2V5KSB7XG4gICAgICBjYXNlICcgJzpcbiAgICAgIGNhc2UgJ0VudGVyJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgdGhpcy4jc2VsZWN0SXRlbSgpO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ0Fycm93TGVmdCc6XG4gICAgICBjYXNlICdBcnJvd1JpZ2h0JzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgdGhpcy4jbmF2aWdhdGVJdGVtcyhldmVudC5rZXkgPT09ICdBcnJvd0xlZnQnID8gLTEgOiAxKTtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdIb21lJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgdGhpcy4jbmF2aWdhdGVUb0VuZCh0cnVlKTtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdFbmQnOlxuICAgICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICB0aGlzLiNuYXZpZ2F0ZVRvRW5kKGZhbHNlKTtcbiAgICAgICAgYnJlYWs7XG4gICAgfVxuICB9O1xuXG4gICNoYW5kbGVGb2N1cyA9ICgpID0+IHtcbiAgICB0aGlzLiN1cGRhdGVSb3ZpbmdUYWJpbmRleCgpO1xuICB9O1xuXG4gICNzZWxlY3RJdGVtKCkge1xuICAgIGlmICghdGhpcy5zZWxlY3RlZCkge1xuICAgICAgdGhpcy5zZWxlY3RlZCA9IHRydWU7XG4gICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IFRvZ2dsZUdyb3VwSXRlbVNlbGVjdEV2ZW50KFxuICAgICAgICB0aGlzLFxuICAgICAgICB0cnVlLFxuICAgICAgICB0aGlzLmdldEF0dHJpYnV0ZSgndmFsdWUnKVxuICAgICAgKSk7XG4gICAgfVxuICB9XG5cbiAgI2ZvY3VzQW5kU2VsZWN0KGl0ZW06IFBmVjZUb2dnbGVHcm91cEl0ZW0pIHtcbiAgICBpdGVtLmZvY3VzKCk7XG4gICAgaXRlbS5zZWxlY3RlZCA9IHRydWU7XG4gICAgaXRlbS5kaXNwYXRjaEV2ZW50KG5ldyBUb2dnbGVHcm91cEl0ZW1TZWxlY3RFdmVudChcbiAgICAgIGl0ZW0sXG4gICAgICB0cnVlLFxuICAgICAgaXRlbS5nZXRBdHRyaWJ1dGUoJ3ZhbHVlJylcbiAgICApKTtcbiAgfVxuXG4gICNuYXZpZ2F0ZUl0ZW1zKGRpcmVjdGlvbjogbnVtYmVyKSB7XG4gICAgY29uc3QgcGFyZW50ID0gdGhpcy5wYXJlbnRFbGVtZW50O1xuICAgIGlmICghcGFyZW50KSByZXR1cm47XG4gICAgY29uc3QgaXRlbXMgPSBBcnJheS5mcm9tKHBhcmVudC5xdWVyeVNlbGVjdG9yQWxsPFBmVjZUb2dnbGVHcm91cEl0ZW0+KCdwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbScpKVxuICAgICAgLmZpbHRlcihpdGVtID0+ICFpdGVtLmRpc2FibGVkKTtcbiAgICBjb25zdCBjdXJyZW50SW5kZXggPSBpdGVtcy5pbmRleE9mKHRoaXMpO1xuICAgIGlmIChjdXJyZW50SW5kZXggPT09IC0xKSByZXR1cm47XG4gICAgbGV0IG5ld0luZGV4ID0gY3VycmVudEluZGV4ICsgZGlyZWN0aW9uO1xuICAgIGlmIChuZXdJbmRleCA8IDApIG5ld0luZGV4ID0gaXRlbXMubGVuZ3RoIC0gMTtcbiAgICBlbHNlIGlmIChuZXdJbmRleCA+PSBpdGVtcy5sZW5ndGgpIG5ld0luZGV4ID0gMDtcbiAgICBjb25zdCB0YXJnZXRJdGVtID0gaXRlbXNbbmV3SW5kZXhdO1xuICAgIGlmICh0YXJnZXRJdGVtKSB0aGlzLiNmb2N1c0FuZFNlbGVjdCh0YXJnZXRJdGVtKTtcbiAgfVxuXG4gICNuYXZpZ2F0ZVRvRW5kKHRvU3RhcnQ6IGJvb2xlYW4pIHtcbiAgICBjb25zdCBwYXJlbnQgPSB0aGlzLnBhcmVudEVsZW1lbnQ7XG4gICAgaWYgKCFwYXJlbnQpIHJldHVybjtcbiAgICBjb25zdCBpdGVtcyA9IEFycmF5LmZyb20ocGFyZW50LnF1ZXJ5U2VsZWN0b3JBbGw8UGZWNlRvZ2dsZUdyb3VwSXRlbT4oJ3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJykpXG4gICAgICAuZmlsdGVyKGl0ZW0gPT4gIWl0ZW0uZGlzYWJsZWQpO1xuICAgIGNvbnN0IHRhcmdldEl0ZW0gPSB0b1N0YXJ0ID8gaXRlbXNbMF0gOiBpdGVtc1tpdGVtcy5sZW5ndGggLSAxXTtcbiAgICBpZiAodGFyZ2V0SXRlbSkgdGhpcy4jZm9jdXNBbmRTZWxlY3QodGFyZ2V0SXRlbSk7XG4gIH1cblxuICAjdXBkYXRlUm92aW5nVGFiaW5kZXgoKSB7XG4gICAgY29uc3QgcGFyZW50ID0gdGhpcy5wYXJlbnRFbGVtZW50O1xuICAgIGlmICghcGFyZW50KSByZXR1cm47XG4gICAgY29uc3QgaXRlbXMgPSBBcnJheS5mcm9tKHBhcmVudC5xdWVyeVNlbGVjdG9yQWxsKCdwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbScpKTtcbiAgICBpdGVtcy5mb3JFYWNoKGl0ZW0gPT4ge1xuICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3RhYmluZGV4JywgaXRlbSA9PT0gdGhpcyA/ICcwJyA6ICctMScpO1xuICAgIH0pO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJzogUGZWNlRvZ2dsZUdyb3VwSXRlbTtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIHotaW5kZXg6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLVpJbmRleCk7XFxcXG4gIGRpc3BsYXk6IGlubGluZS1mbGV4O1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1Gb250U2l6ZSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1MaW5lSGVpZ2h0KTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLUNvbG9yKTtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgYm9yZGVyOiAwO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gIHVzZXItc2VsZWN0OiBub25lO1xcXFxufVxcXFxuXFxcXG46aG9zdCg6bm90KDpmaXJzdC1jaGlsZCkpIHtcXFxcbiAgbWFyZ2luLWlubGluZS1zdGFydDogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2l0ZW0tLWl0ZW0tLU1hcmdpbklubGluZVN0YXJ0KTtcXFxcbn1cXFxcblxcXFxuOmhvc3Q6OmJlZm9yZSxcXFxcbjpob3N0OjphZnRlciB7XFxcXG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgaW5zZXQ6IDA7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxuICBjb250ZW50OiBcXFxcXFxcIlxcXFxcXFwiO1xcXFxuICBib3JkZXItc3R5bGU6IHNvbGlkO1xcXFxuICBib3JkZXItcmFkaXVzOiBpbmhlcml0O1xcXFxufVxcXFxuXFxcXG46aG9zdDo6YmVmb3JlIHtcXFxcbiAgYm9yZGVyLXdpZHRoOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlcldpZHRoKTtcXFxcbiAgYm9yZGVyLWJsb2NrLXN0YXJ0LWNvbG9yOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlckJsb2NrU3RhcnRDb2xvciwgdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJDb2xvcikpO1xcXFxuICBib3JkZXItYmxvY2stZW5kLWNvbG9yOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlckJsb2NrRW5kQ29sb3IsIHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWJlZm9yZS0tQm9yZGVyQ29sb3IpKTtcXFxcbiAgYm9yZGVyLWlubGluZS1zdGFydC1jb2xvcjogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJJbmxpbmVTdGFydENvbG9yLCB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlckNvbG9yKSk7XFxcXG4gIGJvcmRlci1pbmxpbmUtZW5kLWNvbG9yOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlcklubGluZUVuZENvbG9yLCB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlckNvbG9yKSk7XFxcXG59XFxcXG5cXFxcbjpob3N0OjphZnRlciB7XFxcXG4gIGluc2V0OiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlcldpZHRoKTtcXFxcbiAgYm9yZGVyLWNvbG9yOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1hZnRlci0tQm9yZGVyQ29sb3IpO1xcXFxuICBib3JkZXItd2lkdGg6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWFmdGVyLS1Cb3JkZXJXaWR0aCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KDpob3Zlcjpub3QoW2Rpc2FibGVkXSkpIHtcXFxcbiAgLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWhvdmVyLS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICAtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLVpJbmRleDogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0taG92ZXItLVpJbmRleCk7XFxcXG4gIC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0taG92ZXItLWJlZm9yZS0tQm9yZGVyQ29sb3IpO1xcXFxuICAtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWFmdGVyLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0taG92ZXItLWFmdGVyLS1Cb3JkZXJXaWR0aCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KDpmb2N1cy12aXNpYmxlKSB7XFxcXG4gIG91dGxpbmU6IDJweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG4gIG91dGxpbmUtb2Zmc2V0OiAycHg7XFxcXG4gIC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tWkluZGV4OiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1ob3Zlci0tWkluZGV4KTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3NlbGVjdGVkXSkge1xcXFxuICAtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tbS1zZWxlY3RlZC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1Db2xvcjogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tbS1zZWxlY3RlZC0tQ29sb3IpO1xcXFxuICAtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWJlZm9yZS0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLW0tc2VsZWN0ZWQtLWJlZm9yZS0tQm9yZGVyQ29sb3IpO1xcXFxuICAtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLVpJbmRleDogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tbS1zZWxlY3RlZC0tWkluZGV4KTtcXFxcbiAgLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1hZnRlci0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLW0tc2VsZWN0ZWQtLWFmdGVyLS1Cb3JkZXJXaWR0aCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtzZWxlY3RlZF0pICsgOmhvc3QoW3NlbGVjdGVkXSkge1xcXFxuICAtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWJlZm9yZS0tQm9yZGVySW5saW5lU3RhcnRDb2xvcjogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tbS1zZWxlY3RlZC1zZWxlY3RlZC0tYmVmb3JlLS1Cb3JkZXJJbmxpbmVTdGFydENvbG9yKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2Rpc2FibGVkXSkge1xcXFxuICAtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tZGlzYWJsZWQtLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tQ29sb3I6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWRpc2FibGVkLS1Db2xvcik7XFxcXG4gIC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tZGlzYWJsZWQtLWJlZm9yZS0tQm9yZGVyQ29sb3IpO1xcXFxuICAtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLVpJbmRleDogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tZGlzYWJsZWQtLVpJbmRleCk7XFxcXG4gIGN1cnNvcjogbm90LWFsbG93ZWQ7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxufVxcXFxuXFxcXG46aG9zdChbZGlzYWJsZWRdKSArIDpob3N0KFtkaXNhYmxlZF0pIHtcXFxcbiAgLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlcklubGluZVN0YXJ0Q29sb3I6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWRpc2FibGVkLWRpc2FibGVkLS1iZWZvcmUtLUJvcmRlcklubGluZVN0YXJ0Q29sb3IpO1xcXFxufVxcXFxuXFxcXG46aG9zdCg6Zmlyc3QtY2hpbGQpIHtcXFxcbiAgYm9yZGVyLXN0YXJ0LXN0YXJ0LXJhZGl1czogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2l0ZW0tLWZpcnN0LWNoaWxkX19idXR0b24tLUJvcmRlclN0YXJ0U3RhcnRSYWRpdXMpO1xcXFxuICBib3JkZXItZW5kLXN0YXJ0LXJhZGl1czogdmFyKC0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2l0ZW0tLWZpcnN0LWNoaWxkX19idXR0b24tLUJvcmRlckVuZFN0YXJ0UmFkaXVzKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoOmxhc3QtY2hpbGQpIHtcXFxcbiAgYm9yZGVyLXN0YXJ0LWVuZC1yYWRpdXM6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19pdGVtLS1sYXN0LWNoaWxkX19idXR0b24tLUJvcmRlclN0YXJ0RW5kUmFkaXVzKTtcXFxcbiAgYm9yZGVyLWVuZC1lbmQtcmFkaXVzOiB2YXIoLS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9faXRlbS0tbGFzdC1jaGlsZF9fYnV0dG9uLS1Cb3JkZXJFbmRFbmRSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG4jd3JhcHBlciB7XFxcXG4gIGRpc3BsYXk6IHZhcigtLV9oYXMtY29udGVudCwgaW5saW5lLWZsZXgpO1xcXFxuICBnYXA6IHZhcigtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19pY29uLS10ZXh0LS1NYXJnaW5JbmxpbmVTdGFydCk7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG59XFxcXG5cXFxcbjo6c2xvdHRlZChzdmcpIHtcXFxcbiAgd2lkdGg6IDFlbTtcXFxcbiAgaGVpZ2h0OiAxZW07XFxcXG4gIHZlcnRpY2FsLWFsaWduOiAtMC4xMjVlbTtcXFxcbiAgZmlsbDogY3VycmVudENvbG9yO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLG04SkFBdThKLENBQUM7QUFBRSxJQUFPLGtDQUFROzs7QURTdmdLLElBQU0sNkJBQU4sY0FBeUMsTUFBTTtBQUFBLEVBQ3BEO0FBQUEsRUFDQTtBQUFBLEVBQ0E7QUFBQSxFQUNBLFlBQVksTUFBZSxVQUFtQixPQUFzQjtBQUNsRSxVQUFNLGtDQUFrQyxFQUFFLFNBQVMsS0FBSyxDQUFDO0FBQ3pELFNBQUssT0FBTztBQUNaLFNBQUssV0FBVztBQUNoQixTQUFLLFFBQVE7QUFBQSxFQUNmO0FBQ0Y7QUFuQkE7QUFnQ0EsbUNBQUMsY0FBYyx5QkFBeUI7QUFDakMsSUFBTSxzQkFBTixlQUFrQyxpQkFLdkMsaUJBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxpQkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBRzFDLGNBQUMsU0FBUyxFQUFFLFNBQVMsS0FBSyxDQUFDLElBWFksSUFBVztBQUFBLEVBY2xELGNBQWM7QUFDWixVQUFNO0FBZkg7QUFHTCxtQ0FBYSxLQUFLLGdCQUFnQjtBQUdsQyx1QkFBUyxXQUFXLGtCQUFwQixnQkFBb0IsU0FBcEI7QUFHQSx1QkFBUyxXQUFXLGtCQUFwQixpQkFBb0IsU0FBcEI7QUFHQSx1QkFBUyxRQUFUO0FBc0VBLHFDQUFlLENBQUMsVUFBaUI7QUFDL0IsVUFBSSxL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
