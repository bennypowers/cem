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

// elements/cem-serve-knob-group/cem-serve-knob-group.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-serve-knob-group/cem-serve-knob-group.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n"'));
var cem_serve_knob_group_default = s;

// elements/cem-serve-knob-group/cem-serve-knob-group.ts
import "../cem-pf-v6-text-input-group/cem-pf-v6-text-input-group.js";
var KnobAttributeChangeEvent = class extends Event {
  name;
  value;
  constructor(name, value) {
    super("knob:attribute-change", { bubbles: true });
    this.name = name;
    this.value = value;
  }
};
var KnobPropertyChangeEvent = class extends Event {
  name;
  value;
  constructor(name, value) {
    super("knob:property-change", { bubbles: true });
    this.name = name;
    this.value = value;
  }
};
var KnobCssPropertyChangeEvent = class extends Event {
  name;
  value;
  constructor(name, value) {
    super("knob:css-property-change", { bubbles: true });
    this.name = name;
    this.value = value;
  }
};
var KnobAttributeClearEvent = class extends Event {
  name;
  constructor(name) {
    super("knob:attribute-clear", { bubbles: true });
    this.name = name;
  }
};
var KnobPropertyClearEvent = class extends Event {
  name;
  constructor(name) {
    super("knob:property-clear", { bubbles: true });
    this.name = name;
  }
};
var KnobCssPropertyClearEvent = class extends Event {
  name;
  constructor(name) {
    super("knob:css-property-clear", { bubbles: true });
    this.name = name;
  }
};
var KnobCssStateChangeEvent = class extends Event {
  name;
  value;
  constructor(name, value) {
    super("knob:css-state-change", { bubbles: true });
    this.name = name;
    this.value = value;
  }
};
var KnobCssStateClearEvent = class extends Event {
  name;
  constructor(name) {
    super("knob:css-state-clear", { bubbles: true });
    this.name = name;
  }
};
var _htmlFor_dec, _a, _CemServeKnobGroup_decorators, _init, _htmlFor, _debounceTimers, _debounceDelay, _colorButtonListeners, _clearButtonListeners, _CemServeKnobGroup_instances, onSlotChange_fn, attachColorButtonListeners_fn, removeColorButtonListeners_fn, attachClearButtonListeners_fn, removeClearButtonListeners_fn, _handleClearButtonClick, updateClearButtonVisibility_fn, _handleInput, _handleChange, isBooleanControl_fn, _handleColorButtonClick, applyChange_fn, parseValue_fn;
_CemServeKnobGroup_decorators = [customElement("cem-serve-knob-group")];
var CemServeKnobGroup = class extends (_a = LitElement, _htmlFor_dec = [property({ reflect: true, attribute: "for" })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeKnobGroup_instances);
    __privateAdd(this, _htmlFor, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
    __privateAdd(this, _debounceTimers, /* @__PURE__ */ new Map());
    __privateAdd(this, _debounceDelay, 250);
    __privateAdd(this, _colorButtonListeners, /* @__PURE__ */ new WeakMap());
    __privateAdd(this, _clearButtonListeners, /* @__PURE__ */ new WeakMap());
    __privateAdd(this, _handleClearButtonClick, (e, button) => {
      e.preventDefault();
      const knobType = button.dataset.knobType;
      const knobName = button.dataset.knobName;
      if (!knobType || !knobName) return;
      const formGroup = button.closest("cem-pf-v6-form-group");
      if (!formGroup) return;
      const control = formGroup.querySelector(
        `[data-knob-type="${knobType}"][data-knob-name="${knobName}"]`
      );
      if (!control) return;
      if (__privateMethod(this, _CemServeKnobGroup_instances, isBooleanControl_fn).call(this, control)) {
        control.checked = false;
      } else {
        control.value = "";
      }
      button.hidden = true;
      control.dispatchEvent(new Event("input", { bubbles: true }));
      switch (knobType) {
        case "attribute":
          this.dispatchEvent(new KnobAttributeClearEvent(knobName));
          break;
        case "property":
          this.dispatchEvent(new KnobPropertyClearEvent(knobName));
          break;
        case "css-property":
          this.dispatchEvent(new KnobCssPropertyClearEvent(knobName));
          break;
        case "css-state":
          this.dispatchEvent(new KnobCssStateClearEvent(knobName));
          break;
      }
    });
    __privateAdd(this, _handleInput, (e) => {
      const control = e.target;
      const knobType = control.dataset?.knobType;
      const knobName = control.dataset?.knobName;
      if (!knobType || !knobName) return;
      __privateMethod(this, _CemServeKnobGroup_instances, updateClearButtonVisibility_fn).call(this, control);
      if (__privateMethod(this, _CemServeKnobGroup_instances, isBooleanControl_fn).call(this, control)) {
        return __privateMethod(this, _CemServeKnobGroup_instances, applyChange_fn).call(this, knobType, knobName, control.checked);
      }
      if (control.tagName === "SELECT") {
        return;
      }
      const key = `${knobType}-${knobName}`;
      clearTimeout(__privateGet(this, _debounceTimers).get(key));
      __privateGet(this, _debounceTimers).set(key, setTimeout(() => {
        __privateMethod(this, _CemServeKnobGroup_instances, applyChange_fn).call(this, knobType, knobName, control.value);
      }, __privateGet(this, _debounceDelay)));
    });
    __privateAdd(this, _handleChange, (e) => {
      const control = e.target;
      const knobType = control.dataset?.knobType;
      const knobName = control.dataset?.knobName;
      if (!knobType || !knobName) return;
      __privateMethod(this, _CemServeKnobGroup_instances, updateClearButtonVisibility_fn).call(this, control);
      const key = `${knobType}-${knobName}`;
      clearTimeout(__privateGet(this, _debounceTimers).get(key));
      const value = __privateMethod(this, _CemServeKnobGroup_instances, isBooleanControl_fn).call(this, control) ? control.checked : control.value;
      __privateMethod(this, _CemServeKnobGroup_instances, applyChange_fn).call(this, knobType, knobName, value);
    });
    __privateAdd(this, _handleColorButtonClick, async (e, button) => {
      e.preventDefault();
      const textInputGroup = button.closest("cem-pf-v6-text-input-group");
      if (!textInputGroup) return;
      const currentValue = textInputGroup.value || "#000000";
      if ("EyeDropper" in window) {
        try {
          const eyeDropper = new window.EyeDropper();
          const result = await eyeDropper.open();
          textInputGroup.value = result.sRGBHex;
          textInputGroup.dispatchEvent(new Event("input", { bubbles: true }));
        } catch (err) {
          if (err.name !== "AbortError") {
            console.warn("[KnobGroup] EyeDropper error:", err);
          }
        }
      } else {
        const colorInput = document.createElement("input");
        colorInput.type = "color";
        colorInput.value = currentValue;
        colorInput.style.position = "absolute";
        colorInput.style.opacity = "0";
        colorInput.style.pointerEvents = "none";
        document.body.appendChild(colorInput);
        colorInput.addEventListener("input", () => {
          textInputGroup.value = colorInput.value;
          textInputGroup.dispatchEvent(new Event("input", { bubbles: true }));
        });
        colorInput.addEventListener("change", () => {
          textInputGroup.value = colorInput.value;
          textInputGroup.dispatchEvent(new Event("input", { bubbles: true }));
          if (colorInput.parentNode) {
            document.body.removeChild(colorInput);
          }
        });
        colorInput.addEventListener("blur", () => {
          if (colorInput.parentNode) {
            document.body.removeChild(colorInput);
          }
        });
        colorInput.click();
      }
    });
  }
  render() {
    return html`<slot @slotchange=${__privateMethod(this, _CemServeKnobGroup_instances, onSlotChange_fn)}></slot>`;
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("input", __privateGet(this, _handleInput));
    this.addEventListener("change", __privateGet(this, _handleChange));
  }
  firstUpdated() {
    __privateMethod(this, _CemServeKnobGroup_instances, attachColorButtonListeners_fn).call(this);
    __privateMethod(this, _CemServeKnobGroup_instances, attachClearButtonListeners_fn).call(this);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    for (const timer of __privateGet(this, _debounceTimers).values()) {
      clearTimeout(timer);
    }
    __privateGet(this, _debounceTimers).clear();
    this.removeEventListener("input", __privateGet(this, _handleInput));
    this.removeEventListener("change", __privateGet(this, _handleChange));
    __privateMethod(this, _CemServeKnobGroup_instances, removeColorButtonListeners_fn).call(this);
    __privateMethod(this, _CemServeKnobGroup_instances, removeClearButtonListeners_fn).call(this);
  }
};
_init = __decoratorStart(_a);
_htmlFor = new WeakMap();
_debounceTimers = new WeakMap();
_debounceDelay = new WeakMap();
_colorButtonListeners = new WeakMap();
_clearButtonListeners = new WeakMap();
_CemServeKnobGroup_instances = new WeakSet();
onSlotChange_fn = function() {
  __privateMethod(this, _CemServeKnobGroup_instances, attachColorButtonListeners_fn).call(this);
  __privateMethod(this, _CemServeKnobGroup_instances, attachClearButtonListeners_fn).call(this);
};
attachColorButtonListeners_fn = function() {
  const buttons = this.querySelectorAll(".color-picker-button");
  for (const button of buttons) {
    if (__privateGet(this, _colorButtonListeners).has(button)) continue;
    const handler = (e) => __privateGet(this, _handleColorButtonClick).call(this, e, button);
    __privateGet(this, _colorButtonListeners).set(button, handler);
    button.addEventListener("click", handler);
  }
};
removeColorButtonListeners_fn = function() {
  const buttons = this.querySelectorAll(".color-picker-button");
  for (const button of buttons) {
    const handler = __privateGet(this, _colorButtonListeners).get(button);
    if (handler) {
      button.removeEventListener("click", handler);
      __privateGet(this, _colorButtonListeners).delete(button);
    }
  }
};
attachClearButtonListeners_fn = function() {
  const buttons = this.querySelectorAll(".knob-clear-button");
  for (const button of buttons) {
    if (__privateGet(this, _clearButtonListeners).has(button)) continue;
    const handler = (e) => __privateGet(this, _handleClearButtonClick).call(this, e, button);
    __privateGet(this, _clearButtonListeners).set(button, handler);
    button.addEventListener("click", handler);
  }
};
removeClearButtonListeners_fn = function() {
  const buttons = this.querySelectorAll(".knob-clear-button");
  for (const button of buttons) {
    const handler = __privateGet(this, _clearButtonListeners).get(button);
    if (handler) {
      button.removeEventListener("click", handler);
      __privateGet(this, _clearButtonListeners).delete(button);
    }
  }
};
_handleClearButtonClick = new WeakMap();
updateClearButtonVisibility_fn = function(control) {
  const knobType = control.dataset.knobType;
  const knobName = control.dataset.knobName;
  if (!knobType || !knobName) return;
  const formGroup = control.closest("cem-pf-v6-form-group");
  if (!formGroup) return;
  const clearButton = formGroup.querySelector(
    `.knob-clear-button[data-knob-type="${knobType}"][data-knob-name="${knobName}"]`
  );
  if (!clearButton) return;
  const hasValue = __privateMethod(this, _CemServeKnobGroup_instances, isBooleanControl_fn).call(this, control) ? control.checked : control.value !== "";
  clearButton.hidden = !hasValue;
};
_handleInput = new WeakMap();
_handleChange = new WeakMap();
isBooleanControl_fn = function(control) {
  if (control.tagName === "CEM-PF-V6-SWITCH") {
    return true;
  }
  if (control.tagName === "INPUT" && control.type === "checkbox") {
    return true;
  }
  return false;
};
_handleColorButtonClick = new WeakMap();
applyChange_fn = function(type, name, value) {
  switch (type) {
    case "attribute":
      this.dispatchEvent(new KnobAttributeChangeEvent(name, value));
      break;
    case "property":
      this.dispatchEvent(new KnobPropertyChangeEvent(name, __privateMethod(this, _CemServeKnobGroup_instances, parseValue_fn).call(this, value)));
      break;
    case "css-property":
      this.dispatchEvent(new KnobCssPropertyChangeEvent(name, value));
      break;
    case "css-state":
      this.dispatchEvent(new KnobCssStateChangeEvent(name, value));
      break;
    default:
      console.warn(`[KnobGroup] Unknown knob type: ${type}`);
      return;
  }
};
parseValue_fn = function(value) {
  if (typeof value === "boolean") {
    return value;
  }
  if (value === "true") return true;
  if (value === "false") return false;
  if (value === "null") return null;
  if (value === "") return "";
  const num = Number(value);
  if (!isNaN(num) && value !== "") {
    return num;
  }
  return value;
};
__decorateElement(_init, 4, "htmlFor", _htmlFor_dec, CemServeKnobGroup, _htmlFor);
CemServeKnobGroup = __decorateElement(_init, 0, "CemServeKnobGroup", _CemServeKnobGroup_decorators, CemServeKnobGroup);
__publicField(CemServeKnobGroup, "styles", cem_serve_knob_group_default);
__runInitializers(_init, 1, CemServeKnobGroup);
export {
  CemServeKnobGroup,
  KnobAttributeChangeEvent,
  KnobAttributeClearEvent,
  KnobCssPropertyChangeEvent,
  KnobCssPropertyClearEvent,
  KnobCssStateChangeEvent,
  KnobCssStateClearEvent,
  KnobPropertyChangeEvent,
  KnobPropertyClearEvent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWtub2ItZ3JvdXAvY2VtLXNlcnZlLWtub2ItZ3JvdXAudHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tc2VydmUta25vYi1ncm91cC9jZW0tc2VydmUta25vYi1ncm91cC5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWtub2ItZ3JvdXAuY3NzJztcblxuaW1wb3J0ICcuLi9jZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cC9jZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cC5qcyc7XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gYSBrbm9iIGF0dHJpYnV0ZSBjaGFuZ2VzXG4gKi9cbmV4cG9ydCBjbGFzcyBLbm9iQXR0cmlidXRlQ2hhbmdlRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIG5hbWU6IHN0cmluZztcbiAgdmFsdWU6IHVua25vd247XG4gIGNvbnN0cnVjdG9yKG5hbWU6IHN0cmluZywgdmFsdWU6IHVua25vd24pIHtcbiAgICBzdXBlcigna25vYjphdHRyaWJ1dGUtY2hhbmdlJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMubmFtZSA9IG5hbWU7XG4gICAgdGhpcy52YWx1ZSA9IHZhbHVlO1xuICB9XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gYSBrbm9iIHByb3BlcnR5IGNoYW5nZXNcbiAqL1xuZXhwb3J0IGNsYXNzIEtub2JQcm9wZXJ0eUNoYW5nZUV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBuYW1lOiBzdHJpbmc7XG4gIHZhbHVlOiB1bmtub3duO1xuICBjb25zdHJ1Y3RvcihuYW1lOiBzdHJpbmcsIHZhbHVlOiB1bmtub3duKSB7XG4gICAgc3VwZXIoJ2tub2I6cHJvcGVydHktY2hhbmdlJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMubmFtZSA9IG5hbWU7XG4gICAgdGhpcy52YWx1ZSA9IHZhbHVlO1xuICB9XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gYSBrbm9iIENTUyBwcm9wZXJ0eSBjaGFuZ2VzXG4gKi9cbmV4cG9ydCBjbGFzcyBLbm9iQ3NzUHJvcGVydHlDaGFuZ2VFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgbmFtZTogc3RyaW5nO1xuICB2YWx1ZTogc3RyaW5nO1xuICBjb25zdHJ1Y3RvcihuYW1lOiBzdHJpbmcsIHZhbHVlOiBzdHJpbmcpIHtcbiAgICBzdXBlcigna25vYjpjc3MtcHJvcGVydHktY2hhbmdlJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMubmFtZSA9IG5hbWU7XG4gICAgdGhpcy52YWx1ZSA9IHZhbHVlO1xuICB9XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gYSBrbm9iIGF0dHJpYnV0ZSBpcyBjbGVhcmVkXG4gKi9cbmV4cG9ydCBjbGFzcyBLbm9iQXR0cmlidXRlQ2xlYXJFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgbmFtZTogc3RyaW5nO1xuICBjb25zdHJ1Y3RvcihuYW1lOiBzdHJpbmcpIHtcbiAgICBzdXBlcigna25vYjphdHRyaWJ1dGUtY2xlYXInLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gICAgdGhpcy5uYW1lID0gbmFtZTtcbiAgfVxufVxuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIGEga25vYiBwcm9wZXJ0eSBpcyBjbGVhcmVkXG4gKi9cbmV4cG9ydCBjbGFzcyBLbm9iUHJvcGVydHlDbGVhckV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBuYW1lOiBzdHJpbmc7XG4gIGNvbnN0cnVjdG9yKG5hbWU6IHN0cmluZykge1xuICAgIHN1cGVyKCdrbm9iOnByb3BlcnR5LWNsZWFyJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMubmFtZSA9IG5hbWU7XG4gIH1cbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBhIGtub2IgQ1NTIHByb3BlcnR5IGlzIGNsZWFyZWRcbiAqL1xuZXhwb3J0IGNsYXNzIEtub2JDc3NQcm9wZXJ0eUNsZWFyRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIG5hbWU6IHN0cmluZztcbiAgY29uc3RydWN0b3IobmFtZTogc3RyaW5nKSB7XG4gICAgc3VwZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNsZWFyJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMubmFtZSA9IG5hbWU7XG4gIH1cbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBhIGtub2IgQ1NTIGN1c3RvbSBzdGF0ZSBjaGFuZ2VzXG4gKi9cbmV4cG9ydCBjbGFzcyBLbm9iQ3NzU3RhdGVDaGFuZ2VFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgbmFtZTogc3RyaW5nO1xuICB2YWx1ZTogYm9vbGVhbjtcbiAgY29uc3RydWN0b3IobmFtZTogc3RyaW5nLCB2YWx1ZTogYm9vbGVhbikge1xuICAgIHN1cGVyKCdrbm9iOmNzcy1zdGF0ZS1jaGFuZ2UnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gICAgdGhpcy5uYW1lID0gbmFtZTtcbiAgICB0aGlzLnZhbHVlID0gdmFsdWU7XG4gIH1cbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBhIGtub2IgQ1NTIGN1c3RvbSBzdGF0ZSBpcyBjbGVhcmVkXG4gKi9cbmV4cG9ydCBjbGFzcyBLbm9iQ3NzU3RhdGVDbGVhckV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBuYW1lOiBzdHJpbmc7XG4gIGNvbnN0cnVjdG9yKG5hbWU6IHN0cmluZykge1xuICAgIHN1cGVyKCdrbm9iOmNzcy1zdGF0ZS1jbGVhcicsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLm5hbWUgPSBuYW1lO1xuICB9XG59XG5cbi8qKlxuICogQ0VNIFNlcnZlIEtub2IgR3JvdXAgQ29tcG9uZW50XG4gKlxuICogSGFuZGxlcyBldmVudCBkZWxlZ2F0aW9uIGFuZCBkZWJvdW5jaW5nIGZvciBmb3JtIGNvbnRyb2xzIHRoYXQgbW9kaWZ5IGRlbW8gZWxlbWVudHMuXG4gKiBDb250cm9scyBtdXN0IGhhdmUgZGF0YS1rbm9iLXR5cGUgYW5kIGRhdGEta25vYi1uYW1lIGF0dHJpYnV0ZXMuXG4gKlxuICogQGF0dHIge3N0cmluZ30gZm9yIC0gSUQgb2YgdGhlIHRhcmdldCBlbGVtZW50IHRvIGNvbnRyb2xcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBrbm9iIGNvbnRyb2xzXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tc2VydmUta25vYi1ncm91cCcpXG5leHBvcnQgY2xhc3MgQ2VtU2VydmVLbm9iR3JvdXAgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlLCBhdHRyaWJ1dGU6ICdmb3InIH0pXG4gIGFjY2Vzc29yIGh0bWxGb3I6IHN0cmluZyB8IHVuZGVmaW5lZDtcblxuICAjZGVib3VuY2VUaW1lcnMgPSBuZXcgTWFwPHN0cmluZywgUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4+KCk7XG4gICNkZWJvdW5jZURlbGF5ID0gMjUwO1xuICAjY29sb3JCdXR0b25MaXN0ZW5lcnMgPSBuZXcgV2Vha01hcDxFbGVtZW50LCBFdmVudExpc3RlbmVyPigpO1xuICAjY2xlYXJCdXR0b25MaXN0ZW5lcnMgPSBuZXcgV2Vha01hcDxFbGVtZW50LCBFdmVudExpc3RlbmVyPigpO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGA8c2xvdCBAc2xvdGNoYW5nZT0ke3RoaXMuI29uU2xvdENoYW5nZX0+PC9zbG90PmA7XG4gIH1cblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignaW5wdXQnLCB0aGlzLiNoYW5kbGVJbnB1dCk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdjaGFuZ2UnLCB0aGlzLiNoYW5kbGVDaGFuZ2UpO1xuICB9XG5cbiAgZmlyc3RVcGRhdGVkKCkge1xuICAgIHRoaXMuI2F0dGFjaENvbG9yQnV0dG9uTGlzdGVuZXJzKCk7XG4gICAgdGhpcy4jYXR0YWNoQ2xlYXJCdXR0b25MaXN0ZW5lcnMoKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgZm9yIChjb25zdCB0aW1lciBvZiB0aGlzLiNkZWJvdW5jZVRpbWVycy52YWx1ZXMoKSkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRpbWVyKTtcbiAgICB9XG4gICAgdGhpcy4jZGVib3VuY2VUaW1lcnMuY2xlYXIoKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2lucHV0JywgdGhpcy4jaGFuZGxlSW5wdXQpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcignY2hhbmdlJywgdGhpcy4jaGFuZGxlQ2hhbmdlKTtcbiAgICB0aGlzLiNyZW1vdmVDb2xvckJ1dHRvbkxpc3RlbmVycygpO1xuICAgIHRoaXMuI3JlbW92ZUNsZWFyQnV0dG9uTGlzdGVuZXJzKCk7XG4gIH1cblxuICAjb25TbG90Q2hhbmdlKCkge1xuICAgIHRoaXMuI2F0dGFjaENvbG9yQnV0dG9uTGlzdGVuZXJzKCk7XG4gICAgdGhpcy4jYXR0YWNoQ2xlYXJCdXR0b25MaXN0ZW5lcnMoKTtcbiAgfVxuXG4gICNhdHRhY2hDb2xvckJ1dHRvbkxpc3RlbmVycygpIHtcbiAgICBjb25zdCBidXR0b25zID0gdGhpcy5xdWVyeVNlbGVjdG9yQWxsKCcuY29sb3ItcGlja2VyLWJ1dHRvbicpO1xuICAgIGZvciAoY29uc3QgYnV0dG9uIG9mIGJ1dHRvbnMpIHtcbiAgICAgIGlmICh0aGlzLiNjb2xvckJ1dHRvbkxpc3RlbmVycy5oYXMoYnV0dG9uKSkgY29udGludWU7XG4gICAgICBjb25zdCBoYW5kbGVyID0gKGU6IEV2ZW50KSA9PiB0aGlzLiNoYW5kbGVDb2xvckJ1dHRvbkNsaWNrKGUsIGJ1dHRvbik7XG4gICAgICB0aGlzLiNjb2xvckJ1dHRvbkxpc3RlbmVycy5zZXQoYnV0dG9uLCBoYW5kbGVyKTtcbiAgICAgIGJ1dHRvbi5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsIGhhbmRsZXIpO1xuICAgIH1cbiAgfVxuXG4gICNyZW1vdmVDb2xvckJ1dHRvbkxpc3RlbmVycygpIHtcbiAgICBjb25zdCBidXR0b25zID0gdGhpcy5xdWVyeVNlbGVjdG9yQWxsKCcuY29sb3ItcGlja2VyLWJ1dHRvbicpO1xuICAgIGZvciAoY29uc3QgYnV0dG9uIG9mIGJ1dHRvbnMpIHtcbiAgICAgIGNvbnN0IGhhbmRsZXIgPSB0aGlzLiNjb2xvckJ1dHRvbkxpc3RlbmVycy5nZXQoYnV0dG9uKTtcbiAgICAgIGlmIChoYW5kbGVyKSB7XG4gICAgICAgIGJ1dHRvbi5yZW1vdmVFdmVudExpc3RlbmVyKCdjbGljaycsIGhhbmRsZXIpO1xuICAgICAgICB0aGlzLiNjb2xvckJ1dHRvbkxpc3RlbmVycy5kZWxldGUoYnV0dG9uKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjYXR0YWNoQ2xlYXJCdXR0b25MaXN0ZW5lcnMoKSB7XG4gICAgY29uc3QgYnV0dG9ucyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCgnLmtub2ItY2xlYXItYnV0dG9uJyk7XG4gICAgZm9yIChjb25zdCBidXR0b24gb2YgYnV0dG9ucykge1xuICAgICAgaWYgKHRoaXMuI2NsZWFyQnV0dG9uTGlzdGVuZXJzLmhhcyhidXR0b24pKSBjb250aW51ZTtcbiAgICAgIGNvbnN0IGhhbmRsZXIgPSAoZTogRXZlbnQpID0+IHRoaXMuI2hhbmRsZUNsZWFyQnV0dG9uQ2xpY2soZSwgYnV0dG9uIGFzIEhUTUxFbGVtZW50KTtcbiAgICAgIHRoaXMuI2NsZWFyQnV0dG9uTGlzdGVuZXJzLnNldChidXR0b24sIGhhbmRsZXIpO1xuICAgICAgYnV0dG9uLmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgaGFuZGxlcik7XG4gICAgfVxuICB9XG5cbiAgI3JlbW92ZUNsZWFyQnV0dG9uTGlzdGVuZXJzKCkge1xuICAgIGNvbnN0IGJ1dHRvbnMgPSB0aGlzLnF1ZXJ5U2VsZWN0b3JBbGwoJy5rbm9iLWNsZWFyLWJ1dHRvbicpO1xuICAgIGZvciAoY29uc3QgYnV0dG9uIG9mIGJ1dHRvbnMpIHtcbiAgICAgIGNvbnN0IGhhbmRsZXIgPSB0aGlzLiNjbGVhckJ1dHRvbkxpc3RlbmVycy5nZXQoYnV0dG9uKTtcbiAgICAgIGlmIChoYW5kbGVyKSB7XG4gICAgICAgIGJ1dHRvbi5yZW1vdmVFdmVudExpc3RlbmVyKCdjbGljaycsIGhhbmRsZXIpO1xuICAgICAgICB0aGlzLiNjbGVhckJ1dHRvbkxpc3RlbmVycy5kZWxldGUoYnV0dG9uKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjaGFuZGxlQ2xlYXJCdXR0b25DbGljayA9IChlOiBFdmVudCwgYnV0dG9uOiBIVE1MRWxlbWVudCkgPT4ge1xuICAgIGUucHJldmVudERlZmF1bHQoKTtcblxuICAgIGNvbnN0IGtub2JUeXBlID0gYnV0dG9uLmRhdGFzZXQua25vYlR5cGU7XG4gICAgY29uc3Qga25vYk5hbWUgPSBidXR0b24uZGF0YXNldC5rbm9iTmFtZTtcblxuICAgIGlmICgha25vYlR5cGUgfHwgIWtub2JOYW1lKSByZXR1cm47XG5cbiAgICBjb25zdCBmb3JtR3JvdXAgPSBidXR0b24uY2xvc2VzdCgnY2VtLXBmLXY2LWZvcm0tZ3JvdXAnKTtcbiAgICBpZiAoIWZvcm1Hcm91cCkgcmV0dXJuO1xuXG4gICAgY29uc3QgY29udHJvbCA9IGZvcm1Hcm91cC5xdWVyeVNlbGVjdG9yKFxuICAgICAgYFtkYXRhLWtub2ItdHlwZT1cIiR7a25vYlR5cGV9XCJdW2RhdGEta25vYi1uYW1lPVwiJHtrbm9iTmFtZX1cIl1gLFxuICAgICkgYXMgSFRNTElucHV0RWxlbWVudCB8IG51bGw7XG4gICAgaWYgKCFjb250cm9sKSByZXR1cm47XG5cbiAgICBpZiAodGhpcy4jaXNCb29sZWFuQ29udHJvbChjb250cm9sKSkge1xuICAgICAgY29udHJvbC5jaGVja2VkID0gZmFsc2U7XG4gICAgfSBlbHNlIHtcbiAgICAgIGNvbnRyb2wudmFsdWUgPSAnJztcbiAgICB9XG5cbiAgICBidXR0b24uaGlkZGVuID0gdHJ1ZTtcblxuICAgIGNvbnRyb2wuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQoJ2lucHV0JywgeyBidWJibGVzOiB0cnVlIH0pKTtcblxuICAgIHN3aXRjaCAoa25vYlR5cGUpIHtcbiAgICAgIGNhc2UgJ2F0dHJpYnV0ZSc6XG4gICAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgS25vYkF0dHJpYnV0ZUNsZWFyRXZlbnQoa25vYk5hbWUpKTtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdwcm9wZXJ0eSc6XG4gICAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgS25vYlByb3BlcnR5Q2xlYXJFdmVudChrbm9iTmFtZSkpO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ2Nzcy1wcm9wZXJ0eSc6XG4gICAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgS25vYkNzc1Byb3BlcnR5Q2xlYXJFdmVudChrbm9iTmFtZSkpO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ2Nzcy1zdGF0ZSc6XG4gICAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgS25vYkNzc1N0YXRlQ2xlYXJFdmVudChrbm9iTmFtZSkpO1xuICAgICAgICBicmVhaztcbiAgICB9XG4gIH07XG5cbiAgI3VwZGF0ZUNsZWFyQnV0dG9uVmlzaWJpbGl0eShjb250cm9sOiBIVE1MRWxlbWVudCkge1xuICAgIGNvbnN0IGtub2JUeXBlID0gY29udHJvbC5kYXRhc2V0Lmtub2JUeXBlO1xuICAgIGNvbnN0IGtub2JOYW1lID0gY29udHJvbC5kYXRhc2V0Lmtub2JOYW1lO1xuXG4gICAgaWYgKCFrbm9iVHlwZSB8fCAha25vYk5hbWUpIHJldHVybjtcblxuICAgIGNvbnN0IGZvcm1Hcm91cCA9IGNvbnRyb2wuY2xvc2VzdCgnY2VtLXBmLXY2LWZvcm0tZ3JvdXAnKTtcbiAgICBpZiAoIWZvcm1Hcm91cCkgcmV0dXJuO1xuXG4gICAgY29uc3QgY2xlYXJCdXR0b24gPSBmb3JtR3JvdXAucXVlcnlTZWxlY3RvcihcbiAgICAgIGAua25vYi1jbGVhci1idXR0b25bZGF0YS1rbm9iLXR5cGU9XCIke2tub2JUeXBlfVwiXVtkYXRhLWtub2ItbmFtZT1cIiR7a25vYk5hbWV9XCJdYCxcbiAgICApIGFzIEhUTUxFbGVtZW50IHwgbnVsbDtcbiAgICBpZiAoIWNsZWFyQnV0dG9uKSByZXR1cm47XG5cbiAgICBjb25zdCBoYXNWYWx1ZSA9IHRoaXMuI2lzQm9vbGVhbkNvbnRyb2woY29udHJvbClcbiAgICAgID8gKGNvbnRyb2wgYXMgSFRNTElucHV0RWxlbWVudCkuY2hlY2tlZFxuICAgICAgOiAoY29udHJvbCBhcyBIVE1MSW5wdXRFbGVtZW50KS52YWx1ZSAhPT0gJyc7XG5cbiAgICBjbGVhckJ1dHRvbi5oaWRkZW4gPSAhaGFzVmFsdWU7XG4gIH1cblxuICAjaGFuZGxlSW5wdXQgPSAoZTogRXZlbnQpID0+IHtcbiAgICBjb25zdCBjb250cm9sID0gZS50YXJnZXQgYXMgSFRNTElucHV0RWxlbWVudDtcbiAgICBjb25zdCBrbm9iVHlwZSA9IGNvbnRyb2wuZGF0YXNldD8ua25vYlR5cGU7XG4gICAgY29uc3Qga25vYk5hbWUgPSBjb250cm9sLmRhdGFzZXQ/Lmtub2JOYW1lO1xuXG4gICAgaWYgKCFrbm9iVHlwZSB8fCAha25vYk5hbWUpIHJldHVybjtcblxuICAgIHRoaXMuI3VwZGF0ZUNsZWFyQnV0dG9uVmlzaWJpbGl0eShjb250cm9sKTtcblxuICAgIGlmICh0aGlzLiNpc0Jvb2xlYW5Db250cm9sKGNvbnRyb2wpKSB7XG4gICAgICByZXR1cm4gdGhpcy4jYXBwbHlDaGFuZ2Uoa25vYlR5cGUsIGtub2JOYW1lLCBjb250cm9sLmNoZWNrZWQpO1xuICAgIH1cblxuICAgIGlmIChjb250cm9sLnRhZ05hbWUgPT09ICdTRUxFQ1QnKSB7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgY29uc3Qga2V5ID0gYCR7a25vYlR5cGV9LSR7a25vYk5hbWV9YDtcbiAgICBjbGVhclRpbWVvdXQodGhpcy4jZGVib3VuY2VUaW1lcnMuZ2V0KGtleSkpO1xuXG4gICAgdGhpcy4jZGVib3VuY2VUaW1lcnMuc2V0KGtleSwgc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICB0aGlzLiNhcHBseUNoYW5nZShrbm9iVHlwZSwga25vYk5hbWUsIGNvbnRyb2wudmFsdWUpO1xuICAgIH0sIHRoaXMuI2RlYm91bmNlRGVsYXkpKTtcbiAgfTtcblxuICAjaGFuZGxlQ2hhbmdlID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgY29udHJvbCA9IGUudGFyZ2V0IGFzIEhUTUxJbnB1dEVsZW1lbnQ7XG4gICAgY29uc3Qga25vYlR5cGUgPSBjb250cm9sLmRhdGFzZXQ/Lmtub2JUeXBlO1xuICAgIGNvbnN0IGtub2JOYW1lID0gY29udHJvbC5kYXRhc2V0Py5rbm9iTmFtZTtcbiAgICBpZiAoIWtub2JUeXBlIHx8ICFrbm9iTmFtZSkgcmV0dXJuO1xuXG4gICAgdGhpcy4jdXBkYXRlQ2xlYXJCdXR0b25WaXNpYmlsaXR5KGNvbnRyb2wpO1xuXG4gICAgY29uc3Qga2V5ID0gYCR7a25vYlR5cGV9LSR7a25vYk5hbWV9YDtcbiAgICBjbGVhclRpbWVvdXQodGhpcy4jZGVib3VuY2VUaW1lcnMuZ2V0KGtleSkpO1xuICAgIGNvbnN0IHZhbHVlID0gdGhpcy4jaXNCb29sZWFuQ29udHJvbChjb250cm9sKSA/IGNvbnRyb2wuY2hlY2tlZCA6IGNvbnRyb2wudmFsdWU7XG4gICAgdGhpcy4jYXBwbHlDaGFuZ2Uoa25vYlR5cGUsIGtub2JOYW1lLCB2YWx1ZSk7XG4gIH07XG5cbiAgI2lzQm9vbGVhbkNvbnRyb2woY29udHJvbDogSFRNTEVsZW1lbnQpOiBib29sZWFuIHtcbiAgICBpZiAoY29udHJvbC50YWdOYW1lID09PSAnQ0VNLVBGLVY2LVNXSVRDSCcpIHtcbiAgICAgIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICBpZiAoY29udHJvbC50YWdOYW1lID09PSAnSU5QVVQnICYmIChjb250cm9sIGFzIEhUTUxJbnB1dEVsZW1lbnQpLnR5cGUgPT09ICdjaGVja2JveCcpIHtcbiAgICAgIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICByZXR1cm4gZmFsc2U7XG4gIH1cblxuICAjaGFuZGxlQ29sb3JCdXR0b25DbGljayA9IGFzeW5jIChlOiBFdmVudCwgYnV0dG9uOiBFbGVtZW50KSA9PiB7XG4gICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuXG4gICAgY29uc3QgdGV4dElucHV0R3JvdXAgPSBidXR0b24uY2xvc2VzdCgnY2VtLXBmLXY2LXRleHQtaW5wdXQtZ3JvdXAnKSBhcyBIVE1MRWxlbWVudCAmIHsgdmFsdWU/OiBzdHJpbmcgfSB8IG51bGw7XG4gICAgaWYgKCF0ZXh0SW5wdXRHcm91cCkgcmV0dXJuO1xuXG4gICAgY29uc3QgY3VycmVudFZhbHVlID0gdGV4dElucHV0R3JvdXAudmFsdWUgfHwgJyMwMDAwMDAnO1xuXG4gICAgaWYgKCdFeWVEcm9wcGVyJyBpbiB3aW5kb3cpIHtcbiAgICAgIHRyeSB7XG4gICAgICAgIGNvbnN0IGV5ZURyb3BwZXIgPSBuZXcgKHdpbmRvdyBhcyBhbnkpLkV5ZURyb3BwZXIoKTtcbiAgICAgICAgY29uc3QgcmVzdWx0ID0gYXdhaXQgZXllRHJvcHBlci5vcGVuKCk7XG5cbiAgICAgICAgdGV4dElucHV0R3JvdXAudmFsdWUgPSByZXN1bHQuc1JHQkhleDtcbiAgICAgICAgdGV4dElucHV0R3JvdXAuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQoJ2lucHV0JywgeyBidWJibGVzOiB0cnVlIH0pKTtcbiAgICAgIH0gY2F0Y2ggKGVycjogYW55KSB7XG4gICAgICAgIGlmIChlcnIubmFtZSAhPT0gJ0Fib3J0RXJyb3InKSB7XG4gICAgICAgICAgY29uc29sZS53YXJuKCdbS25vYkdyb3VwXSBFeWVEcm9wcGVyIGVycm9yOicsIGVycik7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgY29uc3QgY29sb3JJbnB1dCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2lucHV0Jyk7XG4gICAgICBjb2xvcklucHV0LnR5cGUgPSAnY29sb3InO1xuICAgICAgY29sb3JJbnB1dC52YWx1ZSA9IGN1cnJlbnRWYWx1ZTtcbiAgICAgIGNvbG9ySW5wdXQuc3R5bGUucG9zaXRpb24gPSAnYWJzb2x1dGUnO1xuICAgICAgY29sb3JJbnB1dC5zdHlsZS5vcGFjaXR5ID0gJzAnO1xuICAgICAgY29sb3JJbnB1dC5zdHlsZS5wb2ludGVyRXZlbnRzID0gJ25vbmUnO1xuXG4gICAgICBkb2N1bWVudC5ib2R5LmFwcGVuZENoaWxkKGNvbG9ySW5wdXQpO1xuXG4gICAgICBjb2xvcklucHV0LmFkZEV2ZW50TGlzdGVuZXIoJ2lucHV0JywgKCkgPT4ge1xuICAgICAgICB0ZXh0SW5wdXRHcm91cC52YWx1ZSA9IGNvbG9ySW5wdXQudmFsdWU7XG4gICAgICAgIHRleHRJbnB1dEdyb3VwLmRpc3BhdGNoRXZlbnQobmV3IEV2ZW50KCdpbnB1dCcsIHsgYnViYmxlczogdHJ1ZSB9KSk7XG4gICAgICB9KTtcblxuICAgICAgY29sb3JJbnB1dC5hZGRFdmVudExpc3RlbmVyKCdjaGFuZ2UnLCAoKSA9PiB7XG4gICAgICAgIHRleHRJbnB1dEdyb3VwLnZhbHVlID0gY29sb3JJbnB1dC52YWx1ZTtcbiAgICAgICAgdGV4dElucHV0R3JvdXAuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQoJ2lucHV0JywgeyBidWJibGVzOiB0cnVlIH0pKTtcbiAgICAgICAgaWYgKGNvbG9ySW5wdXQucGFyZW50Tm9kZSkge1xuICAgICAgICAgIGRvY3VtZW50LmJvZHkucmVtb3ZlQ2hpbGQoY29sb3JJbnB1dCk7XG4gICAgICAgIH1cbiAgICAgIH0pO1xuXG4gICAgICBjb2xvcklucHV0LmFkZEV2ZW50TGlzdGVuZXIoJ2JsdXInLCAoKSA9PiB7XG4gICAgICAgIGlmIChjb2xvcklucHV0LnBhcmVudE5vZGUpIHtcbiAgICAgICAgICBkb2N1bWVudC5ib2R5LnJlbW92ZUNoaWxkKGNvbG9ySW5wdXQpO1xuICAgICAgICB9XG4gICAgICB9KTtcblxuICAgICAgY29sb3JJbnB1dC5jbGljaygpO1xuICAgIH1cbiAgfTtcblxuICAjYXBwbHlDaGFuZ2UodHlwZTogc3RyaW5nLCBuYW1lOiBzdHJpbmcsIHZhbHVlOiB1bmtub3duKSB7XG4gICAgc3dpdGNoICh0eXBlKSB7XG4gICAgICBjYXNlICdhdHRyaWJ1dGUnOlxuICAgICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IEtub2JBdHRyaWJ1dGVDaGFuZ2VFdmVudChuYW1lLCB2YWx1ZSkpO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ3Byb3BlcnR5JzpcbiAgICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBLbm9iUHJvcGVydHlDaGFuZ2VFdmVudChuYW1lLCB0aGlzLiNwYXJzZVZhbHVlKHZhbHVlKSkpO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ2Nzcy1wcm9wZXJ0eSc6XG4gICAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgS25vYkNzc1Byb3BlcnR5Q2hhbmdlRXZlbnQobmFtZSwgdmFsdWUgYXMgc3RyaW5nKSk7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnY3NzLXN0YXRlJzpcbiAgICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBLbm9iQ3NzU3RhdGVDaGFuZ2VFdmVudChuYW1lLCB2YWx1ZSBhcyBib29sZWFuKSk7XG4gICAgICAgIGJyZWFrO1xuICAgICAgZGVmYXVsdDpcbiAgICAgICAgY29uc29sZS53YXJuKGBbS25vYkdyb3VwXSBVbmtub3duIGtub2IgdHlwZTogJHt0eXBlfWApO1xuICAgICAgICByZXR1cm47XG4gICAgfVxuICB9XG5cbiAgI3BhcnNlVmFsdWUodmFsdWU6IHVua25vd24pOiB1bmtub3duIHtcbiAgICBpZiAodHlwZW9mIHZhbHVlID09PSAnYm9vbGVhbicpIHtcbiAgICAgIHJldHVybiB2YWx1ZTtcbiAgICB9XG5cbiAgICBpZiAodmFsdWUgPT09ICd0cnVlJykgcmV0dXJuIHRydWU7XG4gICAgaWYgKHZhbHVlID09PSAnZmFsc2UnKSByZXR1cm4gZmFsc2U7XG4gICAgaWYgKHZhbHVlID09PSAnbnVsbCcpIHJldHVybiBudWxsO1xuICAgIGlmICh2YWx1ZSA9PT0gJycpIHJldHVybiAnJztcblxuICAgIGNvbnN0IG51bSA9IE51bWJlcih2YWx1ZSk7XG4gICAgaWYgKCFpc05hTihudW0pICYmIHZhbHVlICE9PSAnJykge1xuICAgICAgcmV0dXJuIG51bTtcbiAgICB9XG5cbiAgICByZXR1cm4gdmFsdWU7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXNlcnZlLWtub2ItZ3JvdXAnOiBDZW1TZXJ2ZUtub2JHcm91cDtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7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
