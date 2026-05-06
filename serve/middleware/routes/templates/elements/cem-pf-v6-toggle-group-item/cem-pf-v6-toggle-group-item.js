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

// elements/cem-pf-v6-toggle-group-item/cem-pf-v6-toggle-group-item.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-toggle-group-item/cem-pf-v6-toggle-group-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  position: relative;\\n  z-index: var(--cem-pf-v6-c-toggle-group__button--ZIndex);\\n  display: inline-flex;\\n  align-items: center;\\n  padding-block-start: var(--cem-pf-v6-c-toggle-group__button--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-toggle-group__button--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-toggle-group__button--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-toggle-group__button--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-toggle-group__button--FontSize);\\n  line-height: var(--cem-pf-v6-c-toggle-group__button--LineHeight);\\n  color: var(--cem-pf-v6-c-toggle-group__button--Color);\\n  background-color: var(--cem-pf-v6-c-toggle-group__button--BackgroundColor);\\n  border: 0;\\n  cursor: pointer;\\n  user-select: none;\\n}\\n\\n:host(:not(:first-child)) {\\n  margin-inline-start: var(--cem-pf-v6-c-toggle-group__item--item--MarginInlineStart);\\n}\\n\\n:host::before,\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border-style: solid;\\n  border-radius: inherit;\\n}\\n\\n:host::before {\\n  border-width: var(--cem-pf-v6-c-toggle-group__button--before--BorderWidth);\\n  border-block-start-color: var(--cem-pf-v6-c-toggle-group__button--before--BorderBlockStartColor, var(--cem-pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-block-end-color: var(--cem-pf-v6-c-toggle-group__button--before--BorderBlockEndColor, var(--cem-pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-inline-start-color: var(--cem-pf-v6-c-toggle-group__button--before--BorderInlineStartColor, var(--cem-pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-inline-end-color: var(--cem-pf-v6-c-toggle-group__button--before--BorderInlineEndColor, var(--cem-pf-v6-c-toggle-group__button--before--BorderColor));\\n}\\n\\n:host::after {\\n  inset: var(--cem-pf-v6-c-toggle-group__button--before--BorderWidth);\\n  border-color: var(--cem-pf-v6-c-toggle-group__button--after--BorderColor);\\n  border-width: var(--cem-pf-v6-c-toggle-group__button--after--BorderWidth);\\n}\\n\\n:host(:hover:not([disabled])) {\\n  --cem-pf-v6-c-toggle-group__button--BackgroundColor: var(--cem-pf-v6-c-toggle-group__button--hover--BackgroundColor);\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: var(--cem-pf-v6-c-toggle-group__button--hover--ZIndex);\\n  --cem-pf-v6-c-toggle-group__button--before--BorderColor: var(--cem-pf-v6-c-toggle-group__button--hover--before--BorderColor);\\n  --cem-pf-v6-c-toggle-group__button--after--BorderWidth: var(--cem-pf-v6-c-toggle-group__button--hover--after--BorderWidth);\\n}\\n\\n:host(:focus-visible) {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: 2px;\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: var(--cem-pf-v6-c-toggle-group__button--hover--ZIndex);\\n}\\n\\n:host([selected]) {\\n  --cem-pf-v6-c-toggle-group__button--BackgroundColor: var(--cem-pf-v6-c-toggle-group__button--m-selected--BackgroundColor);\\n  --cem-pf-v6-c-toggle-group__button--Color: var(--cem-pf-v6-c-toggle-group__button--m-selected--Color);\\n  --cem-pf-v6-c-toggle-group__button--before--BorderColor: var(--cem-pf-v6-c-toggle-group__button--m-selected--before--BorderColor);\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: var(--cem-pf-v6-c-toggle-group__button--m-selected--ZIndex);\\n  --cem-pf-v6-c-toggle-group__button--after--BorderWidth: var(--cem-pf-v6-c-toggle-group__button--m-selected--after--BorderWidth);\\n}\\n\\n:host([selected]) + :host([selected]) {\\n  --cem-pf-v6-c-toggle-group__button--before--BorderInlineStartColor: var(--cem-pf-v6-c-toggle-group__button--m-selected-selected--before--BorderInlineStartColor);\\n}\\n\\n:host([disabled]) {\\n  --cem-pf-v6-c-toggle-group__button--BackgroundColor: var(--cem-pf-v6-c-toggle-group__button--disabled--BackgroundColor);\\n  --cem-pf-v6-c-toggle-group__button--Color: var(--cem-pf-v6-c-toggle-group__button--disabled--Color);\\n  --cem-pf-v6-c-toggle-group__button--before--BorderColor: var(--cem-pf-v6-c-toggle-group__button--disabled--before--BorderColor);\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: var(--cem-pf-v6-c-toggle-group__button--disabled--ZIndex);\\n  cursor: not-allowed;\\n  pointer-events: none;\\n}\\n\\n:host([disabled]) + :host([disabled]) {\\n  --cem-pf-v6-c-toggle-group__button--before--BorderInlineStartColor: var(--cem-pf-v6-c-toggle-group__button--disabled-disabled--before--BorderInlineStartColor);\\n}\\n\\n:host(:first-child) {\\n  border-start-start-radius: var(--cem-pf-v6-c-toggle-group__item--first-child__button--BorderStartStartRadius);\\n  border-end-start-radius: var(--cem-pf-v6-c-toggle-group__item--first-child__button--BorderEndStartRadius);\\n}\\n\\n:host(:last-child) {\\n  border-start-end-radius: var(--cem-pf-v6-c-toggle-group__item--last-child__button--BorderStartEndRadius);\\n  border-end-end-radius: var(--cem-pf-v6-c-toggle-group__item--last-child__button--BorderEndEndRadius);\\n}\\n\\n#wrapper {\\n  display: var(--_has-content, inline-flex);\\n  gap: var(--cem-pf-v6-c-toggle-group__icon--text--MarginInlineStart);\\n  align-items: center;\\n}\\n\\n::slotted(svg) {\\n  width: 1em;\\n  height: 1em;\\n  vertical-align: -0.125em;\\n  fill: currentColor;\\n}\\n"'));
var cem_pf_v6_toggle_group_item_default = s;

// elements/cem-pf-v6-toggle-group-item/cem-pf-v6-toggle-group-item.ts
var ToggleGroupItemSelectEvent = class extends Event {
  item;
  selected;
  value;
  constructor(item, selected, value) {
    super("cem-pf-v6-toggle-group-item-select", { bubbles: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
};
var _value_dec, _disabled_dec, _selected_dec, _a, _PfV6ToggleGroupItem_decorators, _internals, _init, _selected, _disabled, _value, _PfV6ToggleGroupItem_instances, updateTabindex_fn, _handleClick, _handleKeydown, _handleFocus, selectItem_fn, focusAndSelect_fn, navigateItems_fn, navigateToEnd_fn, updateRovingTabindex_fn;
_PfV6ToggleGroupItem_decorators = [customElement("cem-pf-v6-toggle-group-item")];
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
        const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item"));
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
  const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item"));
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
  const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item")).filter((item) => !item.disabled);
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
  const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item")).filter((item) => !item.disabled);
  const targetItem = toStart ? items[0] : items[items.length - 1];
  if (targetItem) __privateMethod(this, _PfV6ToggleGroupItem_instances, focusAndSelect_fn).call(this, targetItem);
};
updateRovingTabindex_fn = function() {
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item"));
  items.forEach((item) => {
    item.setAttribute("tabindex", item === this ? "0" : "-1");
  });
};
__decorateElement(_init, 4, "selected", _selected_dec, PfV6ToggleGroupItem, _selected);
__decorateElement(_init, 4, "disabled", _disabled_dec, PfV6ToggleGroupItem, _disabled);
__decorateElement(_init, 4, "value", _value_dec, PfV6ToggleGroupItem, _value);
PfV6ToggleGroupItem = __decorateElement(_init, 0, "PfV6ToggleGroupItem", _PfV6ToggleGroupItem_decorators, PfV6ToggleGroupItem);
__publicField(PfV6ToggleGroupItem, "styles", cem_pf_v6_toggle_group_item_default);
__runInitializers(_init, 1, PfV6ToggleGroupItem);
export {
  PfV6ToggleGroupItem,
  ToggleGroupItemSelectEvent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtL2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS50cyIsICJsaXQtY3NzOmVsZW1lbnRzL2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS9jZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0uY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS5jc3MnO1xuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmb3IgdG9nZ2xlIGdyb3VwIGl0ZW0gc2VsZWN0aW9uXG4gKi9cbmV4cG9ydCBjbGFzcyBUb2dnbGVHcm91cEl0ZW1TZWxlY3RFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgaXRlbTogRWxlbWVudDtcbiAgc2VsZWN0ZWQ6IGJvb2xlYW47XG4gIHZhbHVlOiBzdHJpbmcgfCBudWxsO1xuICBjb25zdHJ1Y3RvcihpdGVtOiBFbGVtZW50LCBzZWxlY3RlZDogYm9vbGVhbiwgdmFsdWU6IHN0cmluZyB8IG51bGwpIHtcbiAgICBzdXBlcignY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtLXNlbGVjdCcsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLml0ZW0gPSBpdGVtO1xuICAgIHRoaXMuc2VsZWN0ZWQgPSBzZWxlY3RlZDtcbiAgICB0aGlzLnZhbHVlID0gdmFsdWU7XG4gIH1cbn1cblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IFRvZ2dsZSBHcm91cCBJdGVtXG4gKlxuICogQW4gaW5kaXZpZHVhbCB0b2dnbGUgYnV0dG9uIHdpdGhpbiBhIHRvZ2dsZSBncm91cC5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgYnV0dG9uIHRleHRcbiAqIEBzbG90IGljb24gLSBMZWFkaW5nIGljb25cbiAqIEBzbG90IGljb24tZW5kIC0gVHJhaWxpbmcgaWNvblxuICpcbiAqIEBmaXJlcyBjZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0tc2VsZWN0IC0gRmlyZXMgd2hlbiBpdGVtIGlzIHNlbGVjdGVkXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0nKVxuZXhwb3J0IGNsYXNzIFBmVjZUb2dnbGVHcm91cEl0ZW0gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIHNlbGVjdGVkID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBkaXNhYmxlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgdmFsdWU/OiBzdHJpbmc7XG5cbiAgY29uc3RydWN0b3IoKSB7XG4gICAgc3VwZXIoKTtcbiAgICB0aGlzLiNpbnRlcm5hbHMucm9sZSA9ICdyYWRpbyc7XG4gICAgaWYgKCF0aGlzLmhhc0F0dHJpYnV0ZSgndGFiaW5kZXgnKSkge1xuICAgICAgdGhpcy5zZXRBdHRyaWJ1dGUoJ3RhYmluZGV4JywgJy0xJyk7XG4gICAgfVxuICB9XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgdGhpcy4jaGFuZGxlQ2xpY2spO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI2hhbmRsZUtleWRvd24pO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignZm9jdXMnLCB0aGlzLiNoYW5kbGVGb2N1cyk7XG5cbiAgICBxdWV1ZU1pY3JvdGFzaygoKSA9PiB7XG4gICAgICBjb25zdCBwYXJlbnQgPSB0aGlzLnBhcmVudEVsZW1lbnQ7XG4gICAgICBpZiAocGFyZW50KSB7XG4gICAgICAgIGNvbnN0IGl0ZW1zID0gQXJyYXkuZnJvbShwYXJlbnQucXVlcnlTZWxlY3RvckFsbCgnY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJykpO1xuICAgICAgICBjb25zdCBpc0ZpcnN0SXRlbSA9IGl0ZW1zWzBdID09PSB0aGlzO1xuICAgICAgICBpZiAodGhpcy5zZWxlY3RlZCB8fCBpc0ZpcnN0SXRlbSkge1xuICAgICAgICAgIHRoaXMuI3VwZGF0ZVJvdmluZ1RhYmluZGV4KCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdjbGljaycsIHRoaXMuI2hhbmRsZUNsaWNrKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNoYW5kbGVLZXlkb3duKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2ZvY3VzJywgdGhpcy4jaGFuZGxlRm9jdXMpO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygnc2VsZWN0ZWQnKSkge1xuICAgICAgdGhpcy4jaW50ZXJuYWxzLmFyaWFDaGVja2VkID0gU3RyaW5nKHRoaXMuc2VsZWN0ZWQpO1xuICAgICAgdGhpcy4jdXBkYXRlUm92aW5nVGFiaW5kZXgoKTtcbiAgICB9XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdkaXNhYmxlZCcpKSB7XG4gICAgICB0aGlzLiNpbnRlcm5hbHMuYXJpYURpc2FibGVkID0gU3RyaW5nKHRoaXMuZGlzYWJsZWQpO1xuICAgICAgdGhpcy4jdXBkYXRlVGFiaW5kZXgoKTtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8c3BhbiBpZD1cIndyYXBwZXJcIj5cbiAgICAgICAgPHNwYW4gaWQ9XCJpY29uLXN0YXJ0XCI+PHNsb3QgbmFtZT1cImljb25cIj48L3Nsb3Q+PC9zcGFuPlxuICAgICAgICA8c3BhbiBpZD1cInRleHRcIj48c2xvdD48L3Nsb3Q+PC9zcGFuPlxuICAgICAgICA8c3BhbiBpZD1cImljb24tZW5kXCI+PHNsb3QgbmFtZT1cImljb24tZW5kXCI+PC9zbG90Pjwvc3Bhbj5cbiAgICAgIDwvc3Bhbj5cbiAgICBgO1xuICB9XG5cbiAgI3VwZGF0ZVRhYmluZGV4KCkge1xuICAgIGlmICh0aGlzLmRpc2FibGVkKSB7XG4gICAgICB0aGlzLnNldEF0dHJpYnV0ZSgndGFiaW5kZXgnLCAnLTEnKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG4gICAgY29uc3QgcGFyZW50ID0gdGhpcy5wYXJlbnRFbGVtZW50O1xuICAgIGlmICghcGFyZW50KSByZXR1cm47XG4gICAgY29uc3QgaXRlbXMgPSBBcnJheS5mcm9tKHBhcmVudC5xdWVyeVNlbGVjdG9yQWxsKCdjZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0nKSk7XG4gICAgY29uc3QgaXNGaXJzdEl0ZW0gPSBpdGVtc1swXSA9PT0gdGhpcztcbiAgICBjb25zdCBoYXNTZWxlY3RlZEl0ZW0gPSBpdGVtcy5zb21lKGl0ZW0gPT4gaXRlbS5oYXNBdHRyaWJ1dGUoJ3NlbGVjdGVkJykpO1xuICAgIGNvbnN0IHNob3VsZEJlVGFiYmFibGUgPSB0aGlzLnNlbGVjdGVkIHx8IChpc0ZpcnN0SXRlbSAmJiAhaGFzU2VsZWN0ZWRJdGVtKTtcbiAgICB0aGlzLnNldEF0dHJpYnV0ZSgndGFiaW5kZXgnLCBzaG91bGRCZVRhYmJhYmxlID8gJzAnIDogJy0xJyk7XG4gIH1cblxuICAjaGFuZGxlQ2xpY2sgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgaWYgKHRoaXMuZGlzYWJsZWQpIHtcbiAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICBldmVudC5zdG9wSW1tZWRpYXRlUHJvcGFnYXRpb24oKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG4gICAgdGhpcy5mb2N1cygpO1xuICAgIHRoaXMuI3NlbGVjdEl0ZW0oKTtcbiAgfTtcblxuICAjaGFuZGxlS2V5ZG93biA9IChldmVudDogS2V5Ym9hcmRFdmVudCkgPT4ge1xuICAgIGlmICh0aGlzLmRpc2FibGVkKSB7XG4gICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICBzd2l0Y2ggKGV2ZW50LmtleSkge1xuICAgICAgY2FzZSAnICc6XG4gICAgICBjYXNlICdFbnRlcic6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIHRoaXMuI3NlbGVjdEl0ZW0oKTtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdBcnJvd0xlZnQnOlxuICAgICAgY2FzZSAnQXJyb3dSaWdodCc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIHRoaXMuI25hdmlnYXRlSXRlbXMoZXZlbnQua2V5ID09PSAnQXJyb3dMZWZ0JyA/IC0xIDogMSk7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnSG9tZSc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIHRoaXMuI25hdmlnYXRlVG9FbmQodHJ1ZSk7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnRW5kJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgdGhpcy4jbmF2aWdhdGVUb0VuZChmYWxzZSk7XG4gICAgICAgIGJyZWFrO1xuICAgIH1cbiAgfTtcblxuICAjaGFuZGxlRm9jdXMgPSAoKSA9PiB7XG4gICAgdGhpcy4jdXBkYXRlUm92aW5nVGFiaW5kZXgoKTtcbiAgfTtcblxuICAjc2VsZWN0SXRlbSgpIHtcbiAgICBpZiAoIXRoaXMuc2VsZWN0ZWQpIHtcbiAgICAgIHRoaXMuc2VsZWN0ZWQgPSB0cnVlO1xuICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBUb2dnbGVHcm91cEl0ZW1TZWxlY3RFdmVudChcbiAgICAgICAgdGhpcyxcbiAgICAgICAgdHJ1ZSxcbiAgICAgICAgdGhpcy5nZXRBdHRyaWJ1dGUoJ3ZhbHVlJylcbiAgICAgICkpO1xuICAgIH1cbiAgfVxuXG4gICNmb2N1c0FuZFNlbGVjdChpdGVtOiBQZlY2VG9nZ2xlR3JvdXBJdGVtKSB7XG4gICAgaXRlbS5mb2N1cygpO1xuICAgIGl0ZW0uc2VsZWN0ZWQgPSB0cnVlO1xuICAgIGl0ZW0uZGlzcGF0Y2hFdmVudChuZXcgVG9nZ2xlR3JvdXBJdGVtU2VsZWN0RXZlbnQoXG4gICAgICBpdGVtLFxuICAgICAgdHJ1ZSxcbiAgICAgIGl0ZW0uZ2V0QXR0cmlidXRlKCd2YWx1ZScpXG4gICAgKSk7XG4gIH1cblxuICAjbmF2aWdhdGVJdGVtcyhkaXJlY3Rpb246IG51bWJlcikge1xuICAgIGNvbnN0IHBhcmVudCA9IHRoaXMucGFyZW50RWxlbWVudDtcbiAgICBpZiAoIXBhcmVudCkgcmV0dXJuO1xuICAgIGNvbnN0IGl0ZW1zID0gQXJyYXkuZnJvbShwYXJlbnQucXVlcnlTZWxlY3RvckFsbDxQZlY2VG9nZ2xlR3JvdXBJdGVtPignY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJykpXG4gICAgICAuZmlsdGVyKGl0ZW0gPT4gIWl0ZW0uZGlzYWJsZWQpO1xuICAgIGNvbnN0IGN1cnJlbnRJbmRleCA9IGl0ZW1zLmluZGV4T2YodGhpcyk7XG4gICAgaWYgKGN1cnJlbnRJbmRleCA9PT0gLTEpIHJldHVybjtcbiAgICBsZXQgbmV3SW5kZXggPSBjdXJyZW50SW5kZXggKyBkaXJlY3Rpb247XG4gICAgaWYgKG5ld0luZGV4IDwgMCkgbmV3SW5kZXggPSBpdGVtcy5sZW5ndGggLSAxO1xuICAgIGVsc2UgaWYgKG5ld0luZGV4ID49IGl0ZW1zLmxlbmd0aCkgbmV3SW5kZXggPSAwO1xuICAgIGNvbnN0IHRhcmdldEl0ZW0gPSBpdGVtc1tuZXdJbmRleF07XG4gICAgaWYgKHRhcmdldEl0ZW0pIHRoaXMuI2ZvY3VzQW5kU2VsZWN0KHRhcmdldEl0ZW0pO1xuICB9XG5cbiAgI25hdmlnYXRlVG9FbmQodG9TdGFydDogYm9vbGVhbikge1xuICAgIGNvbnN0IHBhcmVudCA9IHRoaXMucGFyZW50RWxlbWVudDtcbiAgICBpZiAoIXBhcmVudCkgcmV0dXJuO1xuICAgIGNvbnN0IGl0ZW1zID0gQXJyYXkuZnJvbShwYXJlbnQucXVlcnlTZWxlY3RvckFsbDxQZlY2VG9nZ2xlR3JvdXBJdGVtPignY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJykpXG4gICAgICAuZmlsdGVyKGl0ZW0gPT4gIWl0ZW0uZGlzYWJsZWQpO1xuICAgIGNvbnN0IHRhcmdldEl0ZW0gPSB0b1N0YXJ0ID8gaXRlbXNbMF0gOiBpdGVtc1tpdGVtcy5sZW5ndGggLSAxXTtcbiAgICBpZiAodGFyZ2V0SXRlbSkgdGhpcy4jZm9jdXNBbmRTZWxlY3QodGFyZ2V0SXRlbSk7XG4gIH1cblxuICAjdXBkYXRlUm92aW5nVGFiaW5kZXgoKSB7XG4gICAgY29uc3QgcGFyZW50ID0gdGhpcy5wYXJlbnRFbGVtZW50O1xuICAgIGlmICghcGFyZW50KSByZXR1cm47XG4gICAgY29uc3QgaXRlbXMgPSBBcnJheS5mcm9tKHBhcmVudC5xdWVyeVNlbGVjdG9yQWxsKCdjZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0nKSk7XG4gICAgaXRlbXMuZm9yRWFjaChpdGVtID0+IHtcbiAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd0YWJpbmRleCcsIGl0ZW0gPT09IHRoaXMgPyAnMCcgOiAnLTEnKTtcbiAgICB9KTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0nOiBQZlY2VG9nZ2xlR3JvdXBJdGVtO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgei1pbmRleDogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLVpJbmRleCk7XFxcXG4gIGRpc3BsYXk6IGlubGluZS1mbGV4O1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tUGFkZGluZ0Jsb2NrU3RhcnQpO1xcXFxuICBwYWRkaW5nLWJsb2NrLWVuZDogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLUZvbnRTaXplKTtcXFxcbiAgbGluZS1oZWlnaHQ6IHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1MaW5lSGVpZ2h0KTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1Db2xvcik7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICBib3JkZXI6IDA7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG59XFxcXG5cXFxcbjpob3N0KDpub3QoOmZpcnN0LWNoaWxkKSkge1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2l0ZW0tLWl0ZW0tLU1hcmdpbklubGluZVN0YXJ0KTtcXFxcbn1cXFxcblxcXFxuOmhvc3Q6OmJlZm9yZSxcXFxcbjpob3N0OjphZnRlciB7XFxcXG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgaW5zZXQ6IDA7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxuICBjb250ZW50OiBcXFxcXFxcIlxcXFxcXFwiO1xcXFxuICBib3JkZXItc3R5bGU6IHNvbGlkO1xcXFxuICBib3JkZXItcmFkaXVzOiBpbmhlcml0O1xcXFxufVxcXFxuXFxcXG46aG9zdDo6YmVmb3JlIHtcXFxcbiAgYm9yZGVyLXdpZHRoOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJXaWR0aCk7XFxcXG4gIGJvcmRlci1ibG9jay1zdGFydC1jb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWJlZm9yZS0tQm9yZGVyQmxvY2tTdGFydENvbG9yLCB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJDb2xvcikpO1xcXFxuICBib3JkZXItYmxvY2stZW5kLWNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJCbG9ja0VuZENvbG9yLCB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJDb2xvcikpO1xcXFxuICBib3JkZXItaW5saW5lLXN0YXJ0LWNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJJbmxpbmVTdGFydENvbG9yLCB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJDb2xvcikpO1xcXFxuICBib3JkZXItaW5saW5lLWVuZC1jb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWJlZm9yZS0tQm9yZGVySW5saW5lRW5kQ29sb3IsIHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlckNvbG9yKSk7XFxcXG59XFxcXG5cXFxcbjpob3N0OjphZnRlciB7XFxcXG4gIGluc2V0OiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJXaWR0aCk7XFxcXG4gIGJvcmRlci1jb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWFmdGVyLS1Cb3JkZXJDb2xvcik7XFxcXG4gIGJvcmRlci13aWR0aDogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWFmdGVyLS1Cb3JkZXJXaWR0aCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KDpob3Zlcjpub3QoW2Rpc2FibGVkXSkpIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0taG92ZXItLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLVpJbmRleDogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWhvdmVyLS1aSW5kZXgpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1iZWZvcmUtLUJvcmRlckNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0taG92ZXItLWJlZm9yZS0tQm9yZGVyQ29sb3IpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1hZnRlci0tQm9yZGVyV2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1ob3Zlci0tYWZ0ZXItLUJvcmRlcldpZHRoKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoOmZvY3VzLXZpc2libGUpIHtcXFxcbiAgb3V0bGluZTogMnB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgb3V0bGluZS1vZmZzZXQ6IDJweDtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tWkluZGV4OiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0taG92ZXItLVpJbmRleCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtzZWxlY3RlZF0pIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tbS1zZWxlY3RlZC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tQ29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1tLXNlbGVjdGVkLS1Db2xvcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWJlZm9yZS0tQm9yZGVyQ29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9fYnV0dG9uLS1tLXNlbGVjdGVkLS1iZWZvcmUtLUJvcmRlckNvbG9yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tWkluZGV4OiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tbS1zZWxlY3RlZC0tWkluZGV4KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYWZ0ZXItLUJvcmRlcldpZHRoOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tbS1zZWxlY3RlZC0tYWZ0ZXItLUJvcmRlcldpZHRoKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3NlbGVjdGVkXSkgKyA6aG9zdChbc2VsZWN0ZWRdKSB7XFxcXG4gIC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWJlZm9yZS0tQm9yZGVySW5saW5lU3RhcnRDb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLW0tc2VsZWN0ZWQtc2VsZWN0ZWQtLWJlZm9yZS0tQm9yZGVySW5saW5lU3RhcnRDb2xvcik7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtkaXNhYmxlZF0pIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tZGlzYWJsZWQtLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLUNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tZGlzYWJsZWQtLUNvbG9yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJDb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19idXR0b24tLWRpc2FibGVkLS1iZWZvcmUtLUJvcmRlckNvbG9yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tWkluZGV4OiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tZGlzYWJsZWQtLVpJbmRleCk7XFxcXG4gIGN1cnNvcjogbm90LWFsbG93ZWQ7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxufVxcXFxuXFxcXG46aG9zdChbZGlzYWJsZWRdKSArIDpob3N0KFtkaXNhYmxlZF0pIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tYmVmb3JlLS1Cb3JkZXJJbmxpbmVTdGFydENvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2J1dHRvbi0tZGlzYWJsZWQtZGlzYWJsZWQtLWJlZm9yZS0tQm9yZGVySW5saW5lU3RhcnRDb2xvcik7XFxcXG59XFxcXG5cXFxcbjpob3N0KDpmaXJzdC1jaGlsZCkge1xcXFxuICBib3JkZXItc3RhcnQtc3RhcnQtcmFkaXVzOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2l0ZW0tLWZpcnN0LWNoaWxkX19idXR0b24tLUJvcmRlclN0YXJ0U3RhcnRSYWRpdXMpO1xcXFxuICBib3JkZXItZW5kLXN0YXJ0LXJhZGl1czogdmFyKC0tY2VtLXBmLXY2LWMtdG9nZ2xlLWdyb3VwX19pdGVtLS1maXJzdC1jaGlsZF9fYnV0dG9uLS1Cb3JkZXJFbmRTdGFydFJhZGl1cyk7XFxcXG59XFxcXG5cXFxcbjpob3N0KDpsYXN0LWNoaWxkKSB7XFxcXG4gIGJvcmRlci1zdGFydC1lbmQtcmFkaXVzOiB2YXIoLS1jZW0tcGYtdjYtYy10b2dnbGUtZ3JvdXBfX2l0ZW0tLWxhc3QtY2hpbGRfX2J1dHRvbi0tQm9yZGVyU3RhcnRFbmRSYWRpdXMpO1xcXFxuICBib3JkZXItZW5kLWVuZC1yYWRpdXM6IHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9faXRlbS0tbGFzdC1jaGlsZF9fYnV0dG9uLS1Cb3JkZXJFbmRFbmRSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG4jd3JhcHBlciB7XFxcXG4gIGRpc3BsYXk6IHZhcigtLV9oYXMtY29udGVudCwgaW5saW5lLWZsZXgpO1xcXFxuICBnYXA6IHZhcigtLWNlbS1wZi12Ni1jLXRvZ2dsZS1ncm91cF9faWNvbi0tdGV4dC0tTWFyZ2luSW5saW5lU3RhcnQpO1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxufVxcXFxuXFxcXG46OnNsb3R0ZWQoc3ZnKSB7XFxcXG4gIHdpZHRoOiAxZW07XFxcXG4gIGhlaWdodDogMWVtO1xcXFxuICB2ZXJ0aWNhbC1hbGlnbjogLTAuMTI1ZW07XFxcXG4gIGZpbGw6IGN1cnJlbnRDb2xvcjtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSwrcUtBQW1ySyxDQUFDO0FBQUUsSUFBTyxzQ0FBUTs7O0FEU252SyxJQUFNLDZCQUFOLGNBQXlDLE1BQU07QUFBQSxFQUNwRDtBQUFBLEVBQ0E7QUFBQSxFQUNBO0FBQUEsRUFDQSxZQUFZLE1BQWUsVUFBbUIsT0FBc0I7QUFDbEUsVUFBTSxzQ0FBc0MsRUFBRSxTQUFTLEtBQUssQ0FBQztBQUM3RCxTQUFLLE9BQU87QUFDWixTQUFLLFdBQVc7QUFDaEIsU0FBSyxRQUFRO0FBQUEsRUFDZjtBQUNGO0FBbkJBO0FBZ0NBLG1DQUFDLGNBQWMsNkJBQTZCO0FBQ3JDLElBQU0sc0JBQU4sZUFBa0MsaUJBS3ZDLGlCQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFHMUMsaUJBQUMs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
