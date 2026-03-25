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

// elements/pf-v6-select/pf-v6-select.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/pf-v6-select/pf-v6-select.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse(`":host {\\n\\n  /* Form control custom properties */\\n  --pf-v6-c-form-control--ColumnGap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --pf-v6-c-form-control--Color: var(--pf-t--global--text--color--regular);\\n  --pf-v6-c-form-control--FontSize: var(--pf-t--global--font--size--body--default);\\n  --pf-v6-c-form-control--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --pf-v6-c-form-control--Resize: none;\\n  --pf-v6-c-form-control--OutlineOffset: -6px;\\n  --pf-v6-c-form-control--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --pf-v6-c-form-control--before--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --pf-v6-c-form-control--before--BorderStyle: solid;\\n  --pf-v6-c-form-control--before--BorderColor: var(--pf-t--global--border--color--default);\\n  --pf-v6-c-form-control--before--BorderRadius: var(--pf-v6-c-form-control--BorderRadius);\\n  --pf-v6-c-form-control--after--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --pf-v6-c-form-control--after--BorderStyle: solid;\\n  --pf-v6-c-form-control--after--BorderColor: transparent;\\n  --pf-v6-c-form-control--after--BorderRadius: var(--pf-v6-c-form-control--BorderRadius);\\n  --pf-v6-c-form-control--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --pf-v6-c-form-control--Width: 100%;\\n  --pf-v6-c-form-control--inset--base: var(--pf-t--global--spacer--control--horizontal--default);\\n  --pf-v6-c-form-control--PaddingBlockStart--base: var(--pf-t--global--spacer--control--vertical--default);\\n  --pf-v6-c-form-control--PaddingBlockEnd--base: var(--pf-t--global--spacer--control--vertical--default);\\n  --pf-v6-c-form-control--PaddingInlineEnd--base: var(--pf-v6-c-form-control--inset--base);\\n  --pf-v6-c-form-control--PaddingInlineStart--base: var(--pf-v6-c-form-control--inset--base);\\n  --pf-v6-c-form-control--PaddingBlockStart: var(--pf-v6-c-form-control__select--PaddingBlockStart);\\n  --pf-v6-c-form-control--PaddingBlockEnd: var(--pf-v6-c-form-control__select--PaddingBlockEnd);\\n  --pf-v6-c-form-control--PaddingInlineStart: var(--pf-v6-c-form-control__select--PaddingInlineStart);\\n  --pf-v6-c-form-control--PaddingInlineEnd: var(--pf-v6-c-form-control__select--PaddingInlineEnd);\\n  --pf-v6-c-form-control__select--PaddingBlockStart: var(--pf-v6-c-form-control--PaddingBlockStart--base);\\n  --pf-v6-c-form-control__select--PaddingBlockEnd: var(--pf-v6-c-form-control--PaddingBlockEnd--base);\\n  --pf-v6-c-form-control__select--PaddingInlineEnd: var(--pf-v6-c-form-control--PaddingInlineEnd--base);\\n  --pf-v6-c-form-control__select--PaddingInlineStart: var(--pf-v6-c-form-control--PaddingInlineStart--base);\\n  --pf-v6-c-form-control__utilities--PaddingInlineEnd: var(--pf-v6-c-form-control__utilities--select--PaddingInlineEnd);\\n  --pf-v6-c-form-control__utilities--select--PaddingInlineEnd: var(--pf-v6-c-form-control__select--PaddingInlineEnd);\\n  --pf-v6-c-form-control--hover--after--BorderWidth: var(--pf-t--global--border--width--control--hover);\\n  --pf-v6-c-form-control--hover--after--BorderColor: var(--pf-t--global--border--color--hover);\\n\\n  position: relative;\\n  display: grid;\\n  grid-template-columns: 1fr auto;\\n  column-gap: var(--pf-v6-c-form-control--ColumnGap);\\n  align-items: start;\\n  width: var(--pf-v6-c-form-control--Width);\\n  font-size: var(--pf-v6-c-form-control--FontSize);\\n  line-height: var(--pf-v6-c-form-control--LineHeight);\\n  resize: var(--pf-v6-c-form-control--Resize);\\n  background-color: var(--pf-v6-c-form-control--BackgroundColor);\\n  border-radius: var(--pf-v6-c-form-control--BorderRadius);\\n}\\n\\n:host::before,\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n}\\n\\n:host::before {\\n  border-color: var(--pf-v6-c-form-control--before--BorderColor);\\n  border-style: var(--pf-v6-c-form-control--before--BorderStyle);\\n  border-width: var(--pf-v6-c-form-control--before--BorderWidth);\\n  border-radius: var(--pf-v6-c-form-control--before--BorderRadius);\\n}\\n\\n:host::after {\\n  border: var(--pf-v6-c-form-control--after--BorderWidth) var(--pf-v6-c-form-control--after--BorderStyle) var(--pf-v6-c-form-control--after--BorderColor);\\n  border-radius: var(--pf-v6-c-form-control--before--BorderRadius);\\n}\\n\\n:host(:hover) {\\n  --pf-v6-c-form-control--after--BorderColor: var(--pf-v6-c-form-control--hover--after--BorderColor);\\n  --pf-v6-c-form-control--after--BorderWidth: var(--pf-v6-c-form-control--hover--after--BorderWidth);\\n}\\n\\n#select {\\n  grid-row: 1 / 2;\\n  grid-column: 1 / -1;\\n  padding-block-start: var(--pf-v6-c-form-control--PaddingBlockStart);\\n  padding-block-end: var(--pf-v6-c-form-control--PaddingBlockEnd);\\n  padding-inline-start: var(--pf-v6-c-form-control--PaddingInlineStart);\\n  padding-inline-end: var(--pf-v6-c-form-control--PaddingInlineEnd);\\n  margin: 0;\\n  font-family: inherit;\\n  font-size: 100%;\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-v6-c-form-control--Color);\\n  appearance: none;\\n  background-color: var(--pf-v6-c-form-control--BackgroundColor);\\n  border: none;\\n  border-radius: var(--pf-v6-c-form-control--BorderRadius);\\n  outline-offset: var(--pf-v6-c-form-control--OutlineOffset);\\n  cursor: pointer;\\n  background-image: url(\\"data:image/svg+xml,%3Csvg width='10' height='6' viewBox='0 0 10 6' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M1 1L5 5L9 1' stroke='%236a6e73' stroke-width='1.5' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E\\");\\n  background-repeat: no-repeat;\\n  background-position: right var(--pf-t--global--spacer--md) center;\\n}\\n\\n#select * {\\n  color: var(--pf-v6-c-form-control--Color);\\n}\\n"`));
var pf_v6_select_default = s;

// elements/pf-v6-select/pf-v6-select.ts
var _invalid_dec, _disabled_dec, _value_dec, _a, _PfV6Select_decorators, _internals, _observer, _init, _value, _disabled, _invalid, _PfV6Select_instances, selectEl_get, cloneOptions_fn, _onChange, _onInput;
_PfV6Select_decorators = [customElement("pf-v6-select")];
var PfV6Select = class extends (_a = LitElement, _value_dec = [property({ reflect: true })], _disabled_dec = [property({ type: Boolean, reflect: true })], _invalid_dec = [property({ type: Boolean, reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Select_instances);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _observer, new MutationObserver(() => __privateMethod(this, _PfV6Select_instances, cloneOptions_fn).call(this)));
    __privateAdd(this, _value, __runInitializers(_init, 8, this, "")), __runInitializers(_init, 11, this);
    __privateAdd(this, _disabled, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateAdd(this, _invalid, __runInitializers(_init, 16, this, false)), __runInitializers(_init, 19, this);
    __privateAdd(this, _onChange, () => {
      const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
      if (!select) return;
      this.value = select.value;
      __privateGet(this, _internals).setFormValue(select.value);
      this.dispatchEvent(new Event("change", { bubbles: true }));
    });
    __privateAdd(this, _onInput, () => {
      const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
      if (!select) return;
      this.value = select.value;
      __privateGet(this, _internals).setFormValue(select.value);
      this.dispatchEvent(new Event("input", { bubbles: true }));
    });
  }
  render() {
    return html`
      <select id="select"
              ?disabled=${this.disabled}
              aria-invalid=${this.invalid ? "true" : "false"}
              @change=${__privateGet(this, _onChange)}
              @input=${__privateGet(this, _onInput)}></select>
    `;
  }
  firstUpdated() {
    __privateMethod(this, _PfV6Select_instances, cloneOptions_fn).call(this);
    const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
    if (select && this.value) {
      select.value = this.value;
    }
    __privateGet(this, _internals).setFormValue(select?.value || null);
    __privateGet(this, _observer).observe(this, {
      childList: true,
      subtree: true,
      attributes: true
    });
  }
  updated(changed) {
    if (changed.has("value")) {
      const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
      if (select && select.value !== this.value) {
        select.value = this.value;
      }
      __privateGet(this, _internals).setFormValue(this.value || null);
    }
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(this, _observer).disconnect();
  }
};
_init = __decoratorStart(_a);
_internals = new WeakMap();
_observer = new WeakMap();
_value = new WeakMap();
_disabled = new WeakMap();
_invalid = new WeakMap();
_PfV6Select_instances = new WeakSet();
selectEl_get = function() {
  return this.shadowRoot?.getElementById("select");
};
cloneOptions_fn = function() {
  const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
  if (!select) return;
  const currentValue = select.value;
  select.innerHTML = "";
  const options = this.querySelectorAll("option");
  options.forEach((option) => {
    select.appendChild(option.cloneNode(true));
  });
  select.value = currentValue;
};
_onChange = new WeakMap();
_onInput = new WeakMap();
__decorateElement(_init, 4, "value", _value_dec, PfV6Select, _value);
__decorateElement(_init, 4, "disabled", _disabled_dec, PfV6Select, _disabled);
__decorateElement(_init, 4, "invalid", _invalid_dec, PfV6Select, _invalid);
PfV6Select = __decorateElement(_init, 0, "PfV6Select", _PfV6Select_decorators, PfV6Select);
__publicField(PfV6Select, "formAssociated", true);
__publicField(PfV6Select, "shadowRootOptions", {
  ...LitElement.shadowRootOptions,
  delegatesFocus: true
});
__publicField(PfV6Select, "styles", pf_v6_select_default);
__runInitializers(_init, 1, PfV6Select);
export {
  PfV6Select
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtc2VsZWN0L3BmLXY2LXNlbGVjdC50cyIsICJsaXQtY3NzOmVsZW1lbnRzL3BmLXY2LXNlbGVjdC9wZi12Ni1zZWxlY3QuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL3BmLXY2LXNlbGVjdC5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgU2VsZWN0XG4gKlxuICogQSBmb3JtLWFzc29jaWF0ZWQgc2VsZWN0IGNvbnRyb2wgdGhhdCBjbG9uZXMgbGlnaHQgRE9NIGA8b3B0aW9uPmAgZWxlbWVudHNcbiAqIGludG8gYW4gaW50ZXJuYWwgc2hhZG93IGA8c2VsZWN0PmAuXG4gKlxuICogQGZpcmVzIGlucHV0IC0gRmlyZWQgd2hlbiB0aGUgc2VsZWN0IHZhbHVlIGNoYW5nZXNcbiAqIEBmaXJlcyBjaGFuZ2UgLSBGaXJlZCB3aGVuIHRoZSBzZWxlY3QgdmFsdWUgaXMgY29tbWl0dGVkXG4gKi9cbkBjdXN0b21FbGVtZW50KCdwZi12Ni1zZWxlY3QnKVxuZXhwb3J0IGNsYXNzIFBmVjZTZWxlY3QgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHJlYWRvbmx5IGZvcm1Bc3NvY2lhdGVkID0gdHJ1ZTtcblxuICBzdGF0aWMgc2hhZG93Um9vdE9wdGlvbnM6IFNoYWRvd1Jvb3RJbml0ID0ge1xuICAgIC4uLkxpdEVsZW1lbnQuc2hhZG93Um9vdE9wdGlvbnMsXG4gICAgZGVsZWdhdGVzRm9jdXM6IHRydWUsXG4gIH07XG5cbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcbiAgI29ic2VydmVyID0gbmV3IE11dGF0aW9uT2JzZXJ2ZXIoKCkgPT4gdGhpcy4jY2xvbmVPcHRpb25zKCkpO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgdmFsdWUgPSAnJztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGRpc2FibGVkID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBpbnZhbGlkID0gZmFsc2U7XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPHNlbGVjdCBpZD1cInNlbGVjdFwiXG4gICAgICAgICAgICAgID9kaXNhYmxlZD0ke3RoaXMuZGlzYWJsZWR9XG4gICAgICAgICAgICAgIGFyaWEtaW52YWxpZD0ke3RoaXMuaW52YWxpZCA/ICd0cnVlJyA6ICdmYWxzZSd9XG4gICAgICAgICAgICAgIEBjaGFuZ2U9JHt0aGlzLiNvbkNoYW5nZX1cbiAgICAgICAgICAgICAgQGlucHV0PSR7dGhpcy4jb25JbnB1dH0+PC9zZWxlY3Q+XG4gICAgYDtcbiAgfVxuXG4gIGZpcnN0VXBkYXRlZCgpIHtcbiAgICB0aGlzLiNjbG9uZU9wdGlvbnMoKTtcblxuICAgIGNvbnN0IHNlbGVjdCA9IHRoaXMuI3NlbGVjdEVsO1xuICAgIGlmIChzZWxlY3QgJiYgdGhpcy52YWx1ZSkge1xuICAgICAgc2VsZWN0LnZhbHVlID0gdGhpcy52YWx1ZTtcbiAgICB9XG5cbiAgICB0aGlzLiNpbnRlcm5hbHMuc2V0Rm9ybVZhbHVlKHNlbGVjdD8udmFsdWUgfHwgbnVsbCk7XG5cbiAgICB0aGlzLiNvYnNlcnZlci5vYnNlcnZlKHRoaXMsIHtcbiAgICAgIGNoaWxkTGlzdDogdHJ1ZSxcbiAgICAgIHN1YnRyZWU6IHRydWUsXG4gICAgICBhdHRyaWJ1dGVzOiB0cnVlLFxuICAgIH0pO1xuICB9XG5cbiAgZ2V0ICNzZWxlY3RFbCgpOiBIVE1MU2VsZWN0RWxlbWVudCB8IG51bGwge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdzZWxlY3QnKSBhcyBIVE1MU2VsZWN0RWxlbWVudCB8IG51bGw7XG4gIH1cblxuICAjY2xvbmVPcHRpb25zKCkge1xuICAgIGNvbnN0IHNlbGVjdCA9IHRoaXMuI3NlbGVjdEVsO1xuICAgIGlmICghc2VsZWN0KSByZXR1cm47XG5cbiAgICBjb25zdCBjdXJyZW50VmFsdWUgPSBzZWxlY3QudmFsdWU7XG4gICAgc2VsZWN0LmlubmVySFRNTCA9ICcnO1xuXG4gICAgY29uc3Qgb3B0aW9ucyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCgnb3B0aW9uJyk7XG4gICAgb3B0aW9ucy5mb3JFYWNoKG9wdGlvbiA9PiB7XG4gICAgICBzZWxlY3QuYXBwZW5kQ2hpbGQob3B0aW9uLmNsb25lTm9kZSh0cnVlKSk7XG4gICAgfSk7XG5cbiAgICBzZWxlY3QudmFsdWUgPSBjdXJyZW50VmFsdWU7XG4gIH1cblxuICAjb25DaGFuZ2UgPSAoKSA9PiB7XG4gICAgY29uc3Qgc2VsZWN0ID0gdGhpcy4jc2VsZWN0RWw7XG4gICAgaWYgKCFzZWxlY3QpIHJldHVybjtcbiAgICB0aGlzLnZhbHVlID0gc2VsZWN0LnZhbHVlO1xuICAgIHRoaXMuI2ludGVybmFscy5zZXRGb3JtVmFsdWUoc2VsZWN0LnZhbHVlKTtcbiAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IEV2ZW50KCdjaGFuZ2UnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICB9O1xuXG4gICNvbklucHV0ID0gKCkgPT4ge1xuICAgIGNvbnN0IHNlbGVjdCA9IHRoaXMuI3NlbGVjdEVsO1xuICAgIGlmICghc2VsZWN0KSByZXR1cm47XG4gICAgdGhpcy52YWx1ZSA9IHNlbGVjdC52YWx1ZTtcbiAgICB0aGlzLiNpbnRlcm5hbHMuc2V0Rm9ybVZhbHVlKHNlbGVjdC52YWx1ZSk7XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnaW5wdXQnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICB9O1xuXG4gIHVwZGF0ZWQoY2hhbmdlZDogTWFwPHN0cmluZywgdW5rbm93bj4pIHtcbiAgICBpZiAoY2hhbmdlZC5oYXMoJ3ZhbHVlJykpIHtcbiAgICAgIGNvbnN0IHNlbGVjdCA9IHRoaXMuI3NlbGVjdEVsO1xuICAgICAgaWYgKHNlbGVjdCAmJiBzZWxlY3QudmFsdWUgIT09IHRoaXMudmFsdWUpIHtcbiAgICAgICAgc2VsZWN0LnZhbHVlID0gdGhpcy52YWx1ZTtcbiAgICAgIH1cbiAgICAgIHRoaXMuI2ludGVybmFscy5zZXRGb3JtVmFsdWUodGhpcy52YWx1ZSB8fCBudWxsKTtcbiAgICB9XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuI29ic2VydmVyLmRpc2Nvbm5lY3QoKTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdwZi12Ni1zZWxlY3QnOiBQZlY2U2VsZWN0O1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG5cXFxcbiAgLyogRm9ybSBjb250cm9sIGN1c3RvbSBwcm9wZXJ0aWVzICovXFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUNvbHVtbkdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWdhcC0tdGV4dC10by1lbGVtZW50LS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1Gb250U2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tTGluZUhlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVJlc2l6ZTogbm9uZTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tT3V0bGluZU9mZnNldDogLTZweDtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tQm9yZGVyUmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWJlZm9yZS0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tY29udHJvbC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWJlZm9yZS0tQm9yZGVyU3R5bGU6IHNvbGlkO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1iZWZvcmUtLUJvcmRlckNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1iZWZvcmUtLUJvcmRlclJhZGl1czogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJvcmRlclJhZGl1cyk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWFmdGVyLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1jb250cm9sLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYWZ0ZXItLUJvcmRlclN0eWxlOiBzb2xpZDtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYWZ0ZXItLUJvcmRlckNvbG9yOiB0cmFuc3BhcmVudDtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYWZ0ZXItLUJvcmRlclJhZGl1czogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJvcmRlclJhZGl1cyk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tY29udHJvbC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVdpZHRoOiAxMDAlO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1pbnNldC0tYmFzZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLWhvcml6b250YWwtLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nQmxvY2tTdGFydC0tYmFzZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLXZlcnRpY2FsLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0Jsb2NrRW5kLS1iYXNlOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tY29udHJvbC0tdmVydGljYWwtLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nSW5saW5lRW5kLS1iYXNlOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0taW5zZXQtLWJhc2UpO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nSW5saW5lU3RhcnQtLWJhc2U6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1pbnNldC0tYmFzZSk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVBhZGRpbmdCbG9ja1N0YXJ0OiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVBhZGRpbmdCbG9ja0VuZDogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2xfX3NlbGVjdC0tUGFkZGluZ0Jsb2NrRW5kKTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nSW5saW5lU3RhcnQpO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nSW5saW5lRW5kOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVBhZGRpbmdCbG9ja1N0YXJ0LS1iYXNlKTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nQmxvY2tFbmQ6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nQmxvY2tFbmQtLWJhc2UpO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sX19zZWxlY3QtLVBhZGRpbmdJbmxpbmVFbmQ6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nSW5saW5lRW5kLS1iYXNlKTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nSW5saW5lU3RhcnQ6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nSW5saW5lU3RhcnQtLWJhc2UpO1xcXFxuICAtLXBmLXY2LWMtZm9ybS1jb250cm9sX191dGlsaXRpZXMtLVBhZGRpbmdJbmxpbmVFbmQ6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sX191dGlsaXRpZXMtLXNlbGVjdC0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2xfX3V0aWxpdGllcy0tc2VsZWN0LS1QYWRkaW5nSW5saW5lRW5kOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0taG92ZXItLWFmdGVyLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1jb250cm9sLS1ob3Zlcik7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWhvdmVyLS1hZnRlci0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0taG92ZXIpO1xcXFxuXFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgZGlzcGxheTogZ3JpZDtcXFxcbiAgZ3JpZC10ZW1wbGF0ZS1jb2x1bW5zOiAxZnIgYXV0bztcXFxcbiAgY29sdW1uLWdhcDogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUNvbHVtbkdhcCk7XFxcXG4gIGFsaWduLWl0ZW1zOiBzdGFydDtcXFxcbiAgd2lkdGg6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1XaWR0aCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUZvbnRTaXplKTtcXFxcbiAgbGluZS1oZWlnaHQ6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1MaW5lSGVpZ2h0KTtcXFxcbiAgcmVzaXplOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUmVzaXplKTtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1Cb3JkZXJSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG46aG9zdDo6YmVmb3JlLFxcXFxuOmhvc3Q6OmFmdGVyIHtcXFxcbiAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICBpbnNldDogMDtcXFxcbiAgcG9pbnRlci1ldmVudHM6IG5vbmU7XFxcXG4gIGNvbnRlbnQ6IFxcXFxcXFwiXFxcXFxcXCI7XFxcXG59XFxcXG5cXFxcbjpob3N0OjpiZWZvcmUge1xcXFxuICBib3JkZXItY29sb3I6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1iZWZvcmUtLUJvcmRlckNvbG9yKTtcXFxcbiAgYm9yZGVyLXN0eWxlOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYmVmb3JlLS1Cb3JkZXJTdHlsZSk7XFxcXG4gIGJvcmRlci13aWR0aDogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWJlZm9yZS0tQm9yZGVyV2lkdGgpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYmVmb3JlLS1Cb3JkZXJSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG46aG9zdDo6YWZ0ZXIge1xcXFxuICBib3JkZXI6IHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1hZnRlci0tQm9yZGVyV2lkdGgpIHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1hZnRlci0tQm9yZGVyU3R5bGUpIHZhcigtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1hZnRlci0tQm9yZGVyQ29sb3IpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYmVmb3JlLS1Cb3JkZXJSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG46aG9zdCg6aG92ZXIpIHtcXFxcbiAgLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYWZ0ZXItLUJvcmRlckNvbG9yOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0taG92ZXItLWFmdGVyLS1Cb3JkZXJDb2xvcik7XFxcXG4gIC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWFmdGVyLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWhvdmVyLS1hZnRlci0tQm9yZGVyV2lkdGgpO1xcXFxufVxcXFxuXFxcXG4jc2VsZWN0IHtcXFxcbiAgZ3JpZC1yb3c6IDEgLyAyO1xcXFxuICBncmlkLWNvbHVtbjogMSAvIC0xO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0Jsb2NrU3RhcnQpO1xcXFxuICBwYWRkaW5nLWJsb2NrLWVuZDogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgZm9udC1mYW1pbHk6IGluaGVyaXQ7XFxcXG4gIGZvbnQtc2l6ZTogMTAwJTtcXFxcbiAgbGluZS1oZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tbGluZS1oZWlnaHQtLWJvZHkpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUNvbG9yKTtcXFxcbiAgYXBwZWFyYW5jZTogbm9uZTtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIGJvcmRlcjogbm9uZTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJvcmRlclJhZGl1cyk7XFxcXG4gIG91dGxpbmUtb2Zmc2V0OiB2YXIoLS1wZi12Ni1jLWZvcm0tY29udHJvbC0tT3V0bGluZU9mZnNldCk7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgYmFja2dyb3VuZC1pbWFnZTogdXJsKFxcXFxcXFwiZGF0YTppbWFnZS9zdmcreG1sLCUzQ3N2ZyB3aWR0aD0nMTAnIGhlaWdodD0nNicgdmlld0JveD0nMCAwIDEwIDYnIGZpbGw9J25vbmUnIHhtbG5zPSdodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyclM0UlM0NwYXRoIGQ9J00xIDFMNSA1TDkgMScgc3Ryb2tlPSclMjM2YTZlNzMnIHN0cm9rZS13aWR0aD0nMS41JyBzdHJva2UtbGluZWNhcD0ncm91bmQnIHN0cm9rZS1saW5lam9pbj0ncm91bmQnLyUzRSUzQy9zdmclM0VcXFxcXFxcIik7XFxcXG4gIGJhY2tncm91bmQtcmVwZWF0OiBuby1yZXBlYXQ7XFxcXG4gIGJhY2tncm91bmQtcG9zaXRpb246IHJpZ2h0IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCkgY2VudGVyO1xcXFxufVxcXFxuXFxcXG4jc2VsZWN0ICoge1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUNvbG9yKTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSw2ckxBQW1zTCxDQUFDO0FBQUUsSUFBTyx1QkFBUTs7O0FEQTF3TDtBQWVBLDBCQUFDLGNBQWMsY0FBYztBQUN0QixJQUFNLGFBQU4sZUFBeUIsaUJBYTlCLGNBQUMsU0FBUyxFQUFFLFNBQVMsS0FBSyxDQUFDLElBRzNCLGlCQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFHMUMsZ0JBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQW5CWixJQUFXO0FBQUEsRUFBcEM7QUFBQTtBQUFBO0FBVUwsbUNBQWEsS0FBSyxnQkFBZ0I7QUFDbEMsa0NBQVksSUFBSSxpQkFB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
