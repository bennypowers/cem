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

// elements/cem-pf-v6-select/cem-pf-v6-select.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-select/cem-pf-v6-select.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse(`":host {\\n\\n  /* Form control custom properties */\\n  --cem-pf-v6-c-form-control--ColumnGap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-form-control--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-form-control--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-form-control--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-form-control--Resize: none;\\n  --cem-pf-v6-c-form-control--OutlineOffset: -6px;\\n  --cem-pf-v6-c-form-control--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-form-control--before--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-form-control--before--BorderStyle: solid;\\n  --cem-pf-v6-c-form-control--before--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-form-control--before--BorderRadius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  --cem-pf-v6-c-form-control--after--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-form-control--after--BorderStyle: solid;\\n  --cem-pf-v6-c-form-control--after--BorderColor: transparent;\\n  --cem-pf-v6-c-form-control--after--BorderRadius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  --cem-pf-v6-c-form-control--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-form-control--Width: 100%;\\n  --cem-pf-v6-c-form-control--inset--base: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-form-control--PaddingBlockStart--base: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-form-control--PaddingBlockEnd--base: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-form-control--PaddingInlineEnd--base: var(--cem-pf-v6-c-form-control--inset--base);\\n  --cem-pf-v6-c-form-control--PaddingInlineStart--base: var(--cem-pf-v6-c-form-control--inset--base);\\n  --cem-pf-v6-c-form-control--PaddingBlockStart: var(--cem-pf-v6-c-form-control__select--PaddingBlockStart);\\n  --cem-pf-v6-c-form-control--PaddingBlockEnd: var(--cem-pf-v6-c-form-control__select--PaddingBlockEnd);\\n  --cem-pf-v6-c-form-control--PaddingInlineStart: var(--cem-pf-v6-c-form-control__select--PaddingInlineStart);\\n  --cem-pf-v6-c-form-control--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__select--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__select--PaddingBlockStart: var(--cem-pf-v6-c-form-control--PaddingBlockStart--base);\\n  --cem-pf-v6-c-form-control__select--PaddingBlockEnd: var(--cem-pf-v6-c-form-control--PaddingBlockEnd--base);\\n  --cem-pf-v6-c-form-control__select--PaddingInlineEnd: var(--cem-pf-v6-c-form-control--PaddingInlineEnd--base);\\n  --cem-pf-v6-c-form-control__select--PaddingInlineStart: var(--cem-pf-v6-c-form-control--PaddingInlineStart--base);\\n  --cem-pf-v6-c-form-control__utilities--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__utilities--select--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__utilities--select--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__select--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control--hover--after--BorderWidth: var(--pf-t--global--border--width--control--hover);\\n  --cem-pf-v6-c-form-control--hover--after--BorderColor: var(--pf-t--global--border--color--hover);\\n\\n  position: relative;\\n  display: grid;\\n  grid-template-columns: 1fr auto;\\n  column-gap: var(--cem-pf-v6-c-form-control--ColumnGap);\\n  align-items: start;\\n  width: var(--cem-pf-v6-c-form-control--Width);\\n  font-size: var(--cem-pf-v6-c-form-control--FontSize);\\n  line-height: var(--cem-pf-v6-c-form-control--LineHeight);\\n  resize: var(--cem-pf-v6-c-form-control--Resize);\\n  background-color: var(--cem-pf-v6-c-form-control--BackgroundColor);\\n  border-radius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n}\\n\\n:host::before,\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n}\\n\\n:host::before {\\n  border-color: var(--cem-pf-v6-c-form-control--before--BorderColor);\\n  border-style: var(--cem-pf-v6-c-form-control--before--BorderStyle);\\n  border-width: var(--cem-pf-v6-c-form-control--before--BorderWidth);\\n  border-radius: var(--cem-pf-v6-c-form-control--before--BorderRadius);\\n}\\n\\n:host::after {\\n  border: var(--cem-pf-v6-c-form-control--after--BorderWidth) var(--cem-pf-v6-c-form-control--after--BorderStyle) var(--cem-pf-v6-c-form-control--after--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-form-control--before--BorderRadius);\\n}\\n\\n:host(:hover) {\\n  --cem-pf-v6-c-form-control--after--BorderColor: var(--cem-pf-v6-c-form-control--hover--after--BorderColor);\\n  --cem-pf-v6-c-form-control--after--BorderWidth: var(--cem-pf-v6-c-form-control--hover--after--BorderWidth);\\n}\\n\\n#select {\\n  grid-row: 1 / 2;\\n  grid-column: 1 / -1;\\n  padding-block-start: var(--cem-pf-v6-c-form-control--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-form-control--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-form-control--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-form-control--PaddingInlineEnd);\\n  margin: 0;\\n  font-family: inherit;\\n  font-size: 100%;\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--cem-pf-v6-c-form-control--Color);\\n  appearance: none;\\n  background-color: var(--cem-pf-v6-c-form-control--BackgroundColor);\\n  border: none;\\n  border-radius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  outline-offset: var(--cem-pf-v6-c-form-control--OutlineOffset);\\n  cursor: pointer;\\n  background-image: url(\\"data:image/svg+xml,%3Csvg width='10' height='6' viewBox='0 0 10 6' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M1 1L5 5L9 1' stroke='%236a6e73' stroke-width='1.5' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E\\");\\n  background-repeat: no-repeat;\\n  background-position: right var(--pf-t--global--spacer--md) center;\\n}\\n\\n#select * {\\n  color: var(--cem-pf-v6-c-form-control--Color);\\n}\\n"`));
var cem_pf_v6_select_default = s;

// elements/cem-pf-v6-select/cem-pf-v6-select.ts
var _invalid_dec, _disabled_dec, _value_dec, _a, _PfV6Select_decorators, _internals, _observer, _init, _value, _disabled, _invalid, _PfV6Select_instances, selectEl_get, cloneOptions_fn, _onChange, _onInput;
_PfV6Select_decorators = [customElement("cem-pf-v6-select")];
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
__publicField(PfV6Select, "styles", cem_pf_v6_select_default);
__runInitializers(_init, 1, PfV6Select);
export {
  PfV6Select
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXNlbGVjdC9jZW0tcGYtdjYtc2VsZWN0LnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LXNlbGVjdC9jZW0tcGYtdjYtc2VsZWN0LmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtc2VsZWN0LmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBTZWxlY3RcbiAqXG4gKiBBIGZvcm0tYXNzb2NpYXRlZCBzZWxlY3QgY29udHJvbCB0aGF0IGNsb25lcyBsaWdodCBET00gYDxvcHRpb24+YCBlbGVtZW50c1xuICogaW50byBhbiBpbnRlcm5hbCBzaGFkb3cgYDxzZWxlY3Q+YC5cbiAqXG4gKiBAZmlyZXMgaW5wdXQgLSBGaXJlZCB3aGVuIHRoZSBzZWxlY3QgdmFsdWUgY2hhbmdlc1xuICogQGZpcmVzIGNoYW5nZSAtIEZpcmVkIHdoZW4gdGhlIHNlbGVjdCB2YWx1ZSBpcyBjb21taXR0ZWRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1wZi12Ni1zZWxlY3QnKVxuZXhwb3J0IGNsYXNzIFBmVjZTZWxlY3QgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHJlYWRvbmx5IGZvcm1Bc3NvY2lhdGVkID0gdHJ1ZTtcblxuICBzdGF0aWMgc2hhZG93Um9vdE9wdGlvbnM6IFNoYWRvd1Jvb3RJbml0ID0ge1xuICAgIC4uLkxpdEVsZW1lbnQuc2hhZG93Um9vdE9wdGlvbnMsXG4gICAgZGVsZWdhdGVzRm9jdXM6IHRydWUsXG4gIH07XG5cbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcbiAgI29ic2VydmVyID0gbmV3IE11dGF0aW9uT2JzZXJ2ZXIoKCkgPT4gdGhpcy4jY2xvbmVPcHRpb25zKCkpO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgdmFsdWUgPSAnJztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGRpc2FibGVkID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBpbnZhbGlkID0gZmFsc2U7XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPHNlbGVjdCBpZD1cInNlbGVjdFwiXG4gICAgICAgICAgICAgID9kaXNhYmxlZD0ke3RoaXMuZGlzYWJsZWR9XG4gICAgICAgICAgICAgIGFyaWEtaW52YWxpZD0ke3RoaXMuaW52YWxpZCA/ICd0cnVlJyA6ICdmYWxzZSd9XG4gICAgICAgICAgICAgIEBjaGFuZ2U9JHt0aGlzLiNvbkNoYW5nZX1cbiAgICAgICAgICAgICAgQGlucHV0PSR7dGhpcy4jb25JbnB1dH0+PC9zZWxlY3Q+XG4gICAgYDtcbiAgfVxuXG4gIGZpcnN0VXBkYXRlZCgpIHtcbiAgICB0aGlzLiNjbG9uZU9wdGlvbnMoKTtcblxuICAgIGNvbnN0IHNlbGVjdCA9IHRoaXMuI3NlbGVjdEVsO1xuICAgIGlmIChzZWxlY3QgJiYgdGhpcy52YWx1ZSkge1xuICAgICAgc2VsZWN0LnZhbHVlID0gdGhpcy52YWx1ZTtcbiAgICB9XG5cbiAgICB0aGlzLiNpbnRlcm5hbHMuc2V0Rm9ybVZhbHVlKHNlbGVjdD8udmFsdWUgfHwgbnVsbCk7XG5cbiAgICB0aGlzLiNvYnNlcnZlci5vYnNlcnZlKHRoaXMsIHtcbiAgICAgIGNoaWxkTGlzdDogdHJ1ZSxcbiAgICAgIHN1YnRyZWU6IHRydWUsXG4gICAgICBhdHRyaWJ1dGVzOiB0cnVlLFxuICAgIH0pO1xuICB9XG5cbiAgZ2V0ICNzZWxlY3RFbCgpOiBIVE1MU2VsZWN0RWxlbWVudCB8IG51bGwge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdzZWxlY3QnKSBhcyBIVE1MU2VsZWN0RWxlbWVudCB8IG51bGw7XG4gIH1cblxuICAjY2xvbmVPcHRpb25zKCkge1xuICAgIGNvbnN0IHNlbGVjdCA9IHRoaXMuI3NlbGVjdEVsO1xuICAgIGlmICghc2VsZWN0KSByZXR1cm47XG5cbiAgICBjb25zdCBjdXJyZW50VmFsdWUgPSBzZWxlY3QudmFsdWU7XG4gICAgc2VsZWN0LmlubmVySFRNTCA9ICcnO1xuXG4gICAgY29uc3Qgb3B0aW9ucyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCgnb3B0aW9uJyk7XG4gICAgb3B0aW9ucy5mb3JFYWNoKG9wdGlvbiA9PiB7XG4gICAgICBzZWxlY3QuYXBwZW5kQ2hpbGQob3B0aW9uLmNsb25lTm9kZSh0cnVlKSk7XG4gICAgfSk7XG5cbiAgICBzZWxlY3QudmFsdWUgPSBjdXJyZW50VmFsdWU7XG4gIH1cblxuICAjb25DaGFuZ2UgPSAoKSA9PiB7XG4gICAgY29uc3Qgc2VsZWN0ID0gdGhpcy4jc2VsZWN0RWw7XG4gICAgaWYgKCFzZWxlY3QpIHJldHVybjtcbiAgICB0aGlzLnZhbHVlID0gc2VsZWN0LnZhbHVlO1xuICAgIHRoaXMuI2ludGVybmFscy5zZXRGb3JtVmFsdWUoc2VsZWN0LnZhbHVlKTtcbiAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IEV2ZW50KCdjaGFuZ2UnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICB9O1xuXG4gICNvbklucHV0ID0gKCkgPT4ge1xuICAgIGNvbnN0IHNlbGVjdCA9IHRoaXMuI3NlbGVjdEVsO1xuICAgIGlmICghc2VsZWN0KSByZXR1cm47XG4gICAgdGhpcy52YWx1ZSA9IHNlbGVjdC52YWx1ZTtcbiAgICB0aGlzLiNpbnRlcm5hbHMuc2V0Rm9ybVZhbHVlKHNlbGVjdC52YWx1ZSk7XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnaW5wdXQnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICB9O1xuXG4gIHVwZGF0ZWQoY2hhbmdlZDogTWFwPHN0cmluZywgdW5rbm93bj4pIHtcbiAgICBpZiAoY2hhbmdlZC5oYXMoJ3ZhbHVlJykpIHtcbiAgICAgIGNvbnN0IHNlbGVjdCA9IHRoaXMuI3NlbGVjdEVsO1xuICAgICAgaWYgKHNlbGVjdCAmJiBzZWxlY3QudmFsdWUgIT09IHRoaXMudmFsdWUpIHtcbiAgICAgICAgc2VsZWN0LnZhbHVlID0gdGhpcy52YWx1ZTtcbiAgICAgIH1cbiAgICAgIHRoaXMuI2ludGVybmFscy5zZXRGb3JtVmFsdWUodGhpcy52YWx1ZSB8fCBudWxsKTtcbiAgICB9XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuI29ic2VydmVyLmRpc2Nvbm5lY3QoKTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcGYtdjYtc2VsZWN0JzogUGZWNlNlbGVjdDtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuXFxcXG4gIC8qIEZvcm0gY29udHJvbCBjdXN0b20gcHJvcGVydGllcyAqL1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tQ29sdW1uR2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tZ2FwLS10ZXh0LXRvLWVsZW1lbnQtLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tRm9udFNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1MaW5lSGVpZ2h0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLWxpbmUtaGVpZ2h0LS1ib2R5KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVJlc2l6ZTogbm9uZTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLU91dGxpbmVPZmZzZXQ6IC02cHg7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1Cb3JkZXJSYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXNtYWxsKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWJlZm9yZS0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tY29udHJvbC0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1iZWZvcmUtLUJvcmRlclN0eWxlOiBzb2xpZDtcXFxcbiAgLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWJlZm9yZS0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1iZWZvcmUtLUJvcmRlclJhZGl1czogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1Cb3JkZXJSYWRpdXMpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYWZ0ZXItLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWNvbnRyb2wtLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYWZ0ZXItLUJvcmRlclN0eWxlOiBzb2xpZDtcXFxcbiAgLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWFmdGVyLS1Cb3JkZXJDb2xvcjogdHJhbnNwYXJlbnQ7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1hZnRlci0tQm9yZGVyUmFkaXVzOiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJvcmRlclJhZGl1cyk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWNvbnRyb2wtLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tV2lkdGg6IDEwMCU7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1pbnNldC0tYmFzZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLWhvcml6b250YWwtLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0Jsb2NrU3RhcnQtLWJhc2U6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nQmxvY2tFbmQtLWJhc2U6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nSW5saW5lRW5kLS1iYXNlOiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWluc2V0LS1iYXNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVBhZGRpbmdJbmxpbmVTdGFydC0tYmFzZTogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1pbnNldC0tYmFzZSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sX19zZWxlY3QtLVBhZGRpbmdCbG9ja1N0YXJ0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVBhZGRpbmdCbG9ja0VuZDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sX19zZWxlY3QtLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nSW5saW5lU3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nSW5saW5lU3RhcnQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sX19zZWxlY3QtLVBhZGRpbmdJbmxpbmVFbmQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nQmxvY2tTdGFydC0tYmFzZSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sX19zZWxlY3QtLVBhZGRpbmdCbG9ja0VuZDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nQmxvY2tFbmQtLWJhc2UpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nSW5saW5lRW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLVBhZGRpbmdJbmxpbmVFbmQtLWJhc2UpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbF9fc2VsZWN0LS1QYWRkaW5nSW5saW5lU3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0lubGluZVN0YXJ0LS1iYXNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2xfX3V0aWxpdGllcy0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sX191dGlsaXRpZXMtLXNlbGVjdC0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sX191dGlsaXRpZXMtLXNlbGVjdC0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sX19zZWxlY3QtLVBhZGRpbmdJbmxpbmVFbmQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0taG92ZXItLWFmdGVyLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1jb250cm9sLS1ob3Zlcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1ob3Zlci0tYWZ0ZXItLUJvcmRlckNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWhvdmVyKTtcXFxcblxcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIGdyaWQtdGVtcGxhdGUtY29sdW1uczogMWZyIGF1dG87XFxcXG4gIGNvbHVtbi1nYXA6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tQ29sdW1uR2FwKTtcXFxcbiAgYWxpZ24taXRlbXM6IHN0YXJ0O1xcXFxuICB3aWR0aDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1XaWR0aCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1Gb250U2l6ZSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUxpbmVIZWlnaHQpO1xcXFxuICByZXNpemU6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUmVzaXplKTtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJvcmRlclJhZGl1cyk7XFxcXG59XFxcXG5cXFxcbjpob3N0OjpiZWZvcmUsXFxcXG46aG9zdDo6YWZ0ZXIge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIGluc2V0OiAwO1xcXFxuICBwb2ludGVyLWV2ZW50czogbm9uZTtcXFxcbiAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcIjtcXFxcbn1cXFxcblxcXFxuOmhvc3Q6OmJlZm9yZSB7XFxcXG4gIGJvcmRlci1jb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1iZWZvcmUtLUJvcmRlckNvbG9yKTtcXFxcbiAgYm9yZGVyLXN0eWxlOiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWJlZm9yZS0tQm9yZGVyU3R5bGUpO1xcXFxuICBib3JkZXItd2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYmVmb3JlLS1Cb3JkZXJXaWR0aCk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYmVmb3JlLS1Cb3JkZXJSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG46aG9zdDo6YWZ0ZXIge1xcXFxuICBib3JkZXI6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tYWZ0ZXItLUJvcmRlcldpZHRoKSB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWFmdGVyLS1Cb3JkZXJTdHlsZSkgdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1hZnRlci0tQm9yZGVyQ29sb3IpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLWJlZm9yZS0tQm9yZGVyUmFkaXVzKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoOmhvdmVyKSB7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1hZnRlci0tQm9yZGVyQ29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0taG92ZXItLWFmdGVyLS1Cb3JkZXJDb2xvcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1hZnRlci0tQm9yZGVyV2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0taG92ZXItLWFmdGVyLS1Cb3JkZXJXaWR0aCk7XFxcXG59XFxcXG5cXFxcbiNzZWxlY3Qge1xcXFxuICBncmlkLXJvdzogMSAvIDI7XFxcXG4gIGdyaWQtY29sdW1uOiAxIC8gLTE7XFxcXG4gIHBhZGRpbmctYmxvY2stc3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0Jsb2NrU3RhcnQpO1xcXFxuICBwYWRkaW5nLWJsb2NrLWVuZDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1QYWRkaW5nSW5saW5lU3RhcnQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1lbmQ6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgZm9udC1mYW1pbHk6IGluaGVyaXQ7XFxcXG4gIGZvbnQtc2l6ZTogMTAwJTtcXFxcbiAgbGluZS1oZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tbGluZS1oZWlnaHQtLWJvZHkpO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtZm9ybS1jb250cm9sLS1Db2xvcik7XFxcXG4gIGFwcGVhcmFuY2U6IG5vbmU7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgYm9yZGVyOiBub25lO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLUJvcmRlclJhZGl1cyk7XFxcXG4gIG91dGxpbmUtb2Zmc2V0OiB2YXIoLS1jZW0tcGYtdjYtYy1mb3JtLWNvbnRyb2wtLU91dGxpbmVPZmZzZXQpO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gIGJhY2tncm91bmQtaW1hZ2U6IHVybChcXFxcXFxcImRhdGE6aW1hZ2Uvc3ZnK3htbCwlM0Nzdmcgd2lkdGg9JzEwJyBoZWlnaHQ9JzYnIHZpZXdCb3g9JzAgMCAxMCA2JyBmaWxsPSdub25lJyB4bWxucz0naHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmcnJTNFJTNDcGF0aCBkPSdNMSAxTDUgNUw5IDEnIHN0cm9rZT0nJTIzNmE2ZTczJyBzdHJva2Utd2lkdGg9JzEuNScgc3Ryb2tlLWxpbmVjYXA9J3JvdW5kJyBzdHJva2UtbGluZWpvaW49J3JvdW5kJy8lM0UlM0Mvc3ZnJTNFXFxcXFxcXCIpO1xcXFxuICBiYWNrZ3JvdW5kLXJlcGVhdDogbm8tcmVwZWF0O1xcXFxuICBiYWNrZ3JvdW5kLXBvc2l0aW9uOiByaWdodCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpIGNlbnRlcjtcXFxcbn1cXFxcblxcXFxuI3NlbGVjdCAqIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLWZvcm0tY29udHJvbC0tQ29sb3IpO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElB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
