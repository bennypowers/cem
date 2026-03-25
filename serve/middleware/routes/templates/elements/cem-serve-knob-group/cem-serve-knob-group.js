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

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-serve-knob-group/cem-serve-knob-group.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n"'));
var cem_serve_knob_group_default = s;

// elements/cem-serve-knob-group/cem-serve-knob-group.ts
import "../pf-v6-text-input-group/pf-v6-text-input-group.js";
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
      const formGroup = button.closest("pf-v6-form-group");
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
      const textInputGroup = button.closest("pf-v6-text-input-group");
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
  const formGroup = control.closest("pf-v6-form-group");
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
  if (control.tagName === "PF-V6-SWITCH") {
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
  KnobPropertyChangeEvent,
  KnobPropertyClearEvent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWtub2ItZ3JvdXAvY2VtLXNlcnZlLWtub2ItZ3JvdXAudHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9jZW0tc2VydmUta25vYi1ncm91cC9jZW0tc2VydmUta25vYi1ncm91cC5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWtub2ItZ3JvdXAuY3NzJztcblxuaW1wb3J0ICcuLi9wZi12Ni10ZXh0LWlucHV0LWdyb3VwL3BmLXY2LXRleHQtaW5wdXQtZ3JvdXAuanMnO1xuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIGEga25vYiBhdHRyaWJ1dGUgY2hhbmdlc1xuICovXG5leHBvcnQgY2xhc3MgS25vYkF0dHJpYnV0ZUNoYW5nZUV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBuYW1lOiBzdHJpbmc7XG4gIHZhbHVlOiB1bmtub3duO1xuICBjb25zdHJ1Y3RvcihuYW1lOiBzdHJpbmcsIHZhbHVlOiB1bmtub3duKSB7XG4gICAgc3VwZXIoJ2tub2I6YXR0cmlidXRlLWNoYW5nZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLm5hbWUgPSBuYW1lO1xuICAgIHRoaXMudmFsdWUgPSB2YWx1ZTtcbiAgfVxufVxuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIGEga25vYiBwcm9wZXJ0eSBjaGFuZ2VzXG4gKi9cbmV4cG9ydCBjbGFzcyBLbm9iUHJvcGVydHlDaGFuZ2VFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgbmFtZTogc3RyaW5nO1xuICB2YWx1ZTogdW5rbm93bjtcbiAgY29uc3RydWN0b3IobmFtZTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bikge1xuICAgIHN1cGVyKCdrbm9iOnByb3BlcnR5LWNoYW5nZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLm5hbWUgPSBuYW1lO1xuICAgIHRoaXMudmFsdWUgPSB2YWx1ZTtcbiAgfVxufVxuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIGEga25vYiBDU1MgcHJvcGVydHkgY2hhbmdlc1xuICovXG5leHBvcnQgY2xhc3MgS25vYkNzc1Byb3BlcnR5Q2hhbmdlRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIG5hbWU6IHN0cmluZztcbiAgdmFsdWU6IHN0cmluZztcbiAgY29uc3RydWN0b3IobmFtZTogc3RyaW5nLCB2YWx1ZTogc3RyaW5nKSB7XG4gICAgc3VwZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNoYW5nZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLm5hbWUgPSBuYW1lO1xuICAgIHRoaXMudmFsdWUgPSB2YWx1ZTtcbiAgfVxufVxuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIGEga25vYiBhdHRyaWJ1dGUgaXMgY2xlYXJlZFxuICovXG5leHBvcnQgY2xhc3MgS25vYkF0dHJpYnV0ZUNsZWFyRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIG5hbWU6IHN0cmluZztcbiAgY29uc3RydWN0b3IobmFtZTogc3RyaW5nKSB7XG4gICAgc3VwZXIoJ2tub2I6YXR0cmlidXRlLWNsZWFyJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMubmFtZSA9IG5hbWU7XG4gIH1cbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBhIGtub2IgcHJvcGVydHkgaXMgY2xlYXJlZFxuICovXG5leHBvcnQgY2xhc3MgS25vYlByb3BlcnR5Q2xlYXJFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgbmFtZTogc3RyaW5nO1xuICBjb25zdHJ1Y3RvcihuYW1lOiBzdHJpbmcpIHtcbiAgICBzdXBlcigna25vYjpwcm9wZXJ0eS1jbGVhcicsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLm5hbWUgPSBuYW1lO1xuICB9XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gYSBrbm9iIENTUyBwcm9wZXJ0eSBpcyBjbGVhcmVkXG4gKi9cbmV4cG9ydCBjbGFzcyBLbm9iQ3NzUHJvcGVydHlDbGVhckV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBuYW1lOiBzdHJpbmc7XG4gIGNvbnN0cnVjdG9yKG5hbWU6IHN0cmluZykge1xuICAgIHN1cGVyKCdrbm9iOmNzcy1wcm9wZXJ0eS1jbGVhcicsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLm5hbWUgPSBuYW1lO1xuICB9XG59XG5cbi8qKlxuICogQ0VNIFNlcnZlIEtub2IgR3JvdXAgQ29tcG9uZW50XG4gKlxuICogSGFuZGxlcyBldmVudCBkZWxlZ2F0aW9uIGFuZCBkZWJvdW5jaW5nIGZvciBmb3JtIGNvbnRyb2xzIHRoYXQgbW9kaWZ5IGRlbW8gZWxlbWVudHMuXG4gKiBDb250cm9scyBtdXN0IGhhdmUgZGF0YS1rbm9iLXR5cGUgYW5kIGRhdGEta25vYi1uYW1lIGF0dHJpYnV0ZXMuXG4gKlxuICogQGF0dHIge3N0cmluZ30gZm9yIC0gSUQgb2YgdGhlIHRhcmdldCBlbGVtZW50IHRvIGNvbnRyb2xcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBrbm9iIGNvbnRyb2xzXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tc2VydmUta25vYi1ncm91cCcpXG5leHBvcnQgY2xhc3MgQ2VtU2VydmVLbm9iR3JvdXAgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlLCBhdHRyaWJ1dGU6ICdmb3InIH0pXG4gIGFjY2Vzc29yIGh0bWxGb3I6IHN0cmluZyB8IHVuZGVmaW5lZDtcblxuICAjZGVib3VuY2VUaW1lcnMgPSBuZXcgTWFwPHN0cmluZywgUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4+KCk7XG4gICNkZWJvdW5jZURlbGF5ID0gMjUwO1xuICAjY29sb3JCdXR0b25MaXN0ZW5lcnMgPSBuZXcgV2Vha01hcDxFbGVtZW50LCBFdmVudExpc3RlbmVyPigpO1xuICAjY2xlYXJCdXR0b25MaXN0ZW5lcnMgPSBuZXcgV2Vha01hcDxFbGVtZW50LCBFdmVudExpc3RlbmVyPigpO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGA8c2xvdCBAc2xvdGNoYW5nZT0ke3RoaXMuI29uU2xvdENoYW5nZX0+PC9zbG90PmA7XG4gIH1cblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignaW5wdXQnLCB0aGlzLiNoYW5kbGVJbnB1dCk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdjaGFuZ2UnLCB0aGlzLiNoYW5kbGVDaGFuZ2UpO1xuICB9XG5cbiAgZmlyc3RVcGRhdGVkKCkge1xuICAgIHRoaXMuI2F0dGFjaENvbG9yQnV0dG9uTGlzdGVuZXJzKCk7XG4gICAgdGhpcy4jYXR0YWNoQ2xlYXJCdXR0b25MaXN0ZW5lcnMoKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgZm9yIChjb25zdCB0aW1lciBvZiB0aGlzLiNkZWJvdW5jZVRpbWVycy52YWx1ZXMoKSkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRpbWVyKTtcbiAgICB9XG4gICAgdGhpcy4jZGVib3VuY2VUaW1lcnMuY2xlYXIoKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2lucHV0JywgdGhpcy4jaGFuZGxlSW5wdXQpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcignY2hhbmdlJywgdGhpcy4jaGFuZGxlQ2hhbmdlKTtcbiAgICB0aGlzLiNyZW1vdmVDb2xvckJ1dHRvbkxpc3RlbmVycygpO1xuICAgIHRoaXMuI3JlbW92ZUNsZWFyQnV0dG9uTGlzdGVuZXJzKCk7XG4gIH1cblxuICAjb25TbG90Q2hhbmdlKCkge1xuICAgIHRoaXMuI2F0dGFjaENvbG9yQnV0dG9uTGlzdGVuZXJzKCk7XG4gICAgdGhpcy4jYXR0YWNoQ2xlYXJCdXR0b25MaXN0ZW5lcnMoKTtcbiAgfVxuXG4gICNhdHRhY2hDb2xvckJ1dHRvbkxpc3RlbmVycygpIHtcbiAgICBjb25zdCBidXR0b25zID0gdGhpcy5xdWVyeVNlbGVjdG9yQWxsKCcuY29sb3ItcGlja2VyLWJ1dHRvbicpO1xuICAgIGZvciAoY29uc3QgYnV0dG9uIG9mIGJ1dHRvbnMpIHtcbiAgICAgIGlmICh0aGlzLiNjb2xvckJ1dHRvbkxpc3RlbmVycy5oYXMoYnV0dG9uKSkgY29udGludWU7XG4gICAgICBjb25zdCBoYW5kbGVyID0gKGU6IEV2ZW50KSA9PiB0aGlzLiNoYW5kbGVDb2xvckJ1dHRvbkNsaWNrKGUsIGJ1dHRvbik7XG4gICAgICB0aGlzLiNjb2xvckJ1dHRvbkxpc3RlbmVycy5zZXQoYnV0dG9uLCBoYW5kbGVyKTtcbiAgICAgIGJ1dHRvbi5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsIGhhbmRsZXIpO1xuICAgIH1cbiAgfVxuXG4gICNyZW1vdmVDb2xvckJ1dHRvbkxpc3RlbmVycygpIHtcbiAgICBjb25zdCBidXR0b25zID0gdGhpcy5xdWVyeVNlbGVjdG9yQWxsKCcuY29sb3ItcGlja2VyLWJ1dHRvbicpO1xuICAgIGZvciAoY29uc3QgYnV0dG9uIG9mIGJ1dHRvbnMpIHtcbiAgICAgIGNvbnN0IGhhbmRsZXIgPSB0aGlzLiNjb2xvckJ1dHRvbkxpc3RlbmVycy5nZXQoYnV0dG9uKTtcbiAgICAgIGlmIChoYW5kbGVyKSB7XG4gICAgICAgIGJ1dHRvbi5yZW1vdmVFdmVudExpc3RlbmVyKCdjbGljaycsIGhhbmRsZXIpO1xuICAgICAgICB0aGlzLiNjb2xvckJ1dHRvbkxpc3RlbmVycy5kZWxldGUoYnV0dG9uKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjYXR0YWNoQ2xlYXJCdXR0b25MaXN0ZW5lcnMoKSB7XG4gICAgY29uc3QgYnV0dG9ucyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCgnLmtub2ItY2xlYXItYnV0dG9uJyk7XG4gICAgZm9yIChjb25zdCBidXR0b24gb2YgYnV0dG9ucykge1xuICAgICAgaWYgKHRoaXMuI2NsZWFyQnV0dG9uTGlzdGVuZXJzLmhhcyhidXR0b24pKSBjb250aW51ZTtcbiAgICAgIGNvbnN0IGhhbmRsZXIgPSAoZTogRXZlbnQpID0+IHRoaXMuI2hhbmRsZUNsZWFyQnV0dG9uQ2xpY2soZSwgYnV0dG9uIGFzIEhUTUxFbGVtZW50KTtcbiAgICAgIHRoaXMuI2NsZWFyQnV0dG9uTGlzdGVuZXJzLnNldChidXR0b24sIGhhbmRsZXIpO1xuICAgICAgYnV0dG9uLmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgaGFuZGxlcik7XG4gICAgfVxuICB9XG5cbiAgI3JlbW92ZUNsZWFyQnV0dG9uTGlzdGVuZXJzKCkge1xuICAgIGNvbnN0IGJ1dHRvbnMgPSB0aGlzLnF1ZXJ5U2VsZWN0b3JBbGwoJy5rbm9iLWNsZWFyLWJ1dHRvbicpO1xuICAgIGZvciAoY29uc3QgYnV0dG9uIG9mIGJ1dHRvbnMpIHtcbiAgICAgIGNvbnN0IGhhbmRsZXIgPSB0aGlzLiNjbGVhckJ1dHRvbkxpc3RlbmVycy5nZXQoYnV0dG9uKTtcbiAgICAgIGlmIChoYW5kbGVyKSB7XG4gICAgICAgIGJ1dHRvbi5yZW1vdmVFdmVudExpc3RlbmVyKCdjbGljaycsIGhhbmRsZXIpO1xuICAgICAgICB0aGlzLiNjbGVhckJ1dHRvbkxpc3RlbmVycy5kZWxldGUoYnV0dG9uKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjaGFuZGxlQ2xlYXJCdXR0b25DbGljayA9IChlOiBFdmVudCwgYnV0dG9uOiBIVE1MRWxlbWVudCkgPT4ge1xuICAgIGUucHJldmVudERlZmF1bHQoKTtcblxuICAgIGNvbnN0IGtub2JUeXBlID0gYnV0dG9uLmRhdGFzZXQua25vYlR5cGU7XG4gICAgY29uc3Qga25vYk5hbWUgPSBidXR0b24uZGF0YXNldC5rbm9iTmFtZTtcblxuICAgIGlmICgha25vYlR5cGUgfHwgIWtub2JOYW1lKSByZXR1cm47XG5cbiAgICBjb25zdCBmb3JtR3JvdXAgPSBidXR0b24uY2xvc2VzdCgncGYtdjYtZm9ybS1ncm91cCcpO1xuICAgIGlmICghZm9ybUdyb3VwKSByZXR1cm47XG5cbiAgICBjb25zdCBjb250cm9sID0gZm9ybUdyb3VwLnF1ZXJ5U2VsZWN0b3IoXG4gICAgICBgW2RhdGEta25vYi10eXBlPVwiJHtrbm9iVHlwZX1cIl1bZGF0YS1rbm9iLW5hbWU9XCIke2tub2JOYW1lfVwiXWAsXG4gICAgKSBhcyBIVE1MSW5wdXRFbGVtZW50IHwgbnVsbDtcbiAgICBpZiAoIWNvbnRyb2wpIHJldHVybjtcblxuICAgIGlmICh0aGlzLiNpc0Jvb2xlYW5Db250cm9sKGNvbnRyb2wpKSB7XG4gICAgICBjb250cm9sLmNoZWNrZWQgPSBmYWxzZTtcbiAgICB9IGVsc2Uge1xuICAgICAgY29udHJvbC52YWx1ZSA9ICcnO1xuICAgIH1cblxuICAgIGJ1dHRvbi5oaWRkZW4gPSB0cnVlO1xuXG4gICAgY29udHJvbC5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnaW5wdXQnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuXG4gICAgc3dpdGNoIChrbm9iVHlwZSkge1xuICAgICAgY2FzZSAnYXR0cmlidXRlJzpcbiAgICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBLbm9iQXR0cmlidXRlQ2xlYXJFdmVudChrbm9iTmFtZSkpO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ3Byb3BlcnR5JzpcbiAgICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBLbm9iUHJvcGVydHlDbGVhckV2ZW50KGtub2JOYW1lKSk7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnY3NzLXByb3BlcnR5JzpcbiAgICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBLbm9iQ3NzUHJvcGVydHlDbGVhckV2ZW50KGtub2JOYW1lKSk7XG4gICAgICAgIGJyZWFrO1xuICAgIH1cbiAgfTtcblxuICAjdXBkYXRlQ2xlYXJCdXR0b25WaXNpYmlsaXR5KGNvbnRyb2w6IEhUTUxFbGVtZW50KSB7XG4gICAgY29uc3Qga25vYlR5cGUgPSBjb250cm9sLmRhdGFzZXQua25vYlR5cGU7XG4gICAgY29uc3Qga25vYk5hbWUgPSBjb250cm9sLmRhdGFzZXQua25vYk5hbWU7XG5cbiAgICBpZiAoIWtub2JUeXBlIHx8ICFrbm9iTmFtZSkgcmV0dXJuO1xuXG4gICAgY29uc3QgZm9ybUdyb3VwID0gY29udHJvbC5jbG9zZXN0KCdwZi12Ni1mb3JtLWdyb3VwJyk7XG4gICAgaWYgKCFmb3JtR3JvdXApIHJldHVybjtcblxuICAgIGNvbnN0IGNsZWFyQnV0dG9uID0gZm9ybUdyb3VwLnF1ZXJ5U2VsZWN0b3IoXG4gICAgICBgLmtub2ItY2xlYXItYnV0dG9uW2RhdGEta25vYi10eXBlPVwiJHtrbm9iVHlwZX1cIl1bZGF0YS1rbm9iLW5hbWU9XCIke2tub2JOYW1lfVwiXWAsXG4gICAgKSBhcyBIVE1MRWxlbWVudCB8IG51bGw7XG4gICAgaWYgKCFjbGVhckJ1dHRvbikgcmV0dXJuO1xuXG4gICAgY29uc3QgaGFzVmFsdWUgPSB0aGlzLiNpc0Jvb2xlYW5Db250cm9sKGNvbnRyb2wpXG4gICAgICA/IChjb250cm9sIGFzIEhUTUxJbnB1dEVsZW1lbnQpLmNoZWNrZWRcbiAgICAgIDogKGNvbnRyb2wgYXMgSFRNTElucHV0RWxlbWVudCkudmFsdWUgIT09ICcnO1xuXG4gICAgY2xlYXJCdXR0b24uaGlkZGVuID0gIWhhc1ZhbHVlO1xuICB9XG5cbiAgI2hhbmRsZUlucHV0ID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgY29udHJvbCA9IGUudGFyZ2V0IGFzIEhUTUxJbnB1dEVsZW1lbnQ7XG4gICAgY29uc3Qga25vYlR5cGUgPSBjb250cm9sLmRhdGFzZXQ/Lmtub2JUeXBlO1xuICAgIGNvbnN0IGtub2JOYW1lID0gY29udHJvbC5kYXRhc2V0Py5rbm9iTmFtZTtcblxuICAgIGlmICgha25vYlR5cGUgfHwgIWtub2JOYW1lKSByZXR1cm47XG5cbiAgICB0aGlzLiN1cGRhdGVDbGVhckJ1dHRvblZpc2liaWxpdHkoY29udHJvbCk7XG5cbiAgICBpZiAodGhpcy4jaXNCb29sZWFuQ29udHJvbChjb250cm9sKSkge1xuICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5Q2hhbmdlKGtub2JUeXBlLCBrbm9iTmFtZSwgY29udHJvbC5jaGVja2VkKTtcbiAgICB9XG5cbiAgICBpZiAoY29udHJvbC50YWdOYW1lID09PSAnU0VMRUNUJykge1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IGtleSA9IGAke2tub2JUeXBlfS0ke2tub2JOYW1lfWA7XG4gICAgY2xlYXJUaW1lb3V0KHRoaXMuI2RlYm91bmNlVGltZXJzLmdldChrZXkpKTtcblxuICAgIHRoaXMuI2RlYm91bmNlVGltZXJzLnNldChrZXksIHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgdGhpcy4jYXBwbHlDaGFuZ2Uoa25vYlR5cGUsIGtub2JOYW1lLCBjb250cm9sLnZhbHVlKTtcbiAgICB9LCB0aGlzLiNkZWJvdW5jZURlbGF5KSk7XG4gIH07XG5cbiAgI2hhbmRsZUNoYW5nZSA9IChlOiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IGNvbnRyb2wgPSBlLnRhcmdldCBhcyBIVE1MSW5wdXRFbGVtZW50O1xuICAgIGNvbnN0IGtub2JUeXBlID0gY29udHJvbC5kYXRhc2V0Py5rbm9iVHlwZTtcbiAgICBjb25zdCBrbm9iTmFtZSA9IGNvbnRyb2wuZGF0YXNldD8ua25vYk5hbWU7XG4gICAgaWYgKCFrbm9iVHlwZSB8fCAha25vYk5hbWUpIHJldHVybjtcblxuICAgIHRoaXMuI3VwZGF0ZUNsZWFyQnV0dG9uVmlzaWJpbGl0eShjb250cm9sKTtcblxuICAgIGNvbnN0IGtleSA9IGAke2tub2JUeXBlfS0ke2tub2JOYW1lfWA7XG4gICAgY2xlYXJUaW1lb3V0KHRoaXMuI2RlYm91bmNlVGltZXJzLmdldChrZXkpKTtcbiAgICBjb25zdCB2YWx1ZSA9IHRoaXMuI2lzQm9vbGVhbkNvbnRyb2woY29udHJvbCkgPyBjb250cm9sLmNoZWNrZWQgOiBjb250cm9sLnZhbHVlO1xuICAgIHRoaXMuI2FwcGx5Q2hhbmdlKGtub2JUeXBlLCBrbm9iTmFtZSwgdmFsdWUpO1xuICB9O1xuXG4gICNpc0Jvb2xlYW5Db250cm9sKGNvbnRyb2w6IEhUTUxFbGVtZW50KTogYm9vbGVhbiB7XG4gICAgaWYgKGNvbnRyb2wudGFnTmFtZSA9PT0gJ1BGLVY2LVNXSVRDSCcpIHtcbiAgICAgIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICBpZiAoY29udHJvbC50YWdOYW1lID09PSAnSU5QVVQnICYmIChjb250cm9sIGFzIEhUTUxJbnB1dEVsZW1lbnQpLnR5cGUgPT09ICdjaGVja2JveCcpIHtcbiAgICAgIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICByZXR1cm4gZmFsc2U7XG4gIH1cblxuICAjaGFuZGxlQ29sb3JCdXR0b25DbGljayA9IGFzeW5jIChlOiBFdmVudCwgYnV0dG9uOiBFbGVtZW50KSA9PiB7XG4gICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuXG4gICAgY29uc3QgdGV4dElucHV0R3JvdXAgPSBidXR0b24uY2xvc2VzdCgncGYtdjYtdGV4dC1pbnB1dC1ncm91cCcpIGFzIEhUTUxFbGVtZW50ICYgeyB2YWx1ZT86IHN0cmluZyB9IHwgbnVsbDtcbiAgICBpZiAoIXRleHRJbnB1dEdyb3VwKSByZXR1cm47XG5cbiAgICBjb25zdCBjdXJyZW50VmFsdWUgPSB0ZXh0SW5wdXRHcm91cC52YWx1ZSB8fCAnIzAwMDAwMCc7XG5cbiAgICBpZiAoJ0V5ZURyb3BwZXInIGluIHdpbmRvdykge1xuICAgICAgdHJ5IHtcbiAgICAgICAgY29uc3QgZXllRHJvcHBlciA9IG5ldyAod2luZG93IGFzIGFueSkuRXllRHJvcHBlcigpO1xuICAgICAgICBjb25zdCByZXN1bHQgPSBhd2FpdCBleWVEcm9wcGVyLm9wZW4oKTtcblxuICAgICAgICB0ZXh0SW5wdXRHcm91cC52YWx1ZSA9IHJlc3VsdC5zUkdCSGV4O1xuICAgICAgICB0ZXh0SW5wdXRHcm91cC5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnaW5wdXQnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICAgICAgfSBjYXRjaCAoZXJyOiBhbnkpIHtcbiAgICAgICAgaWYgKGVyci5uYW1lICE9PSAnQWJvcnRFcnJvcicpIHtcbiAgICAgICAgICBjb25zb2xlLndhcm4oJ1tLbm9iR3JvdXBdIEV5ZURyb3BwZXIgZXJyb3I6JywgZXJyKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICBjb25zdCBjb2xvcklucHV0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnaW5wdXQnKTtcbiAgICAgIGNvbG9ySW5wdXQudHlwZSA9ICdjb2xvcic7XG4gICAgICBjb2xvcklucHV0LnZhbHVlID0gY3VycmVudFZhbHVlO1xuICAgICAgY29sb3JJbnB1dC5zdHlsZS5wb3NpdGlvbiA9ICdhYnNvbHV0ZSc7XG4gICAgICBjb2xvcklucHV0LnN0eWxlLm9wYWNpdHkgPSAnMCc7XG4gICAgICBjb2xvcklucHV0LnN0eWxlLnBvaW50ZXJFdmVudHMgPSAnbm9uZSc7XG5cbiAgICAgIGRvY3VtZW50LmJvZHkuYXBwZW5kQ2hpbGQoY29sb3JJbnB1dCk7XG5cbiAgICAgIGNvbG9ySW5wdXQuYWRkRXZlbnRMaXN0ZW5lcignaW5wdXQnLCAoKSA9PiB7XG4gICAgICAgIHRleHRJbnB1dEdyb3VwLnZhbHVlID0gY29sb3JJbnB1dC52YWx1ZTtcbiAgICAgICAgdGV4dElucHV0R3JvdXAuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQoJ2lucHV0JywgeyBidWJibGVzOiB0cnVlIH0pKTtcbiAgICAgIH0pO1xuXG4gICAgICBjb2xvcklucHV0LmFkZEV2ZW50TGlzdGVuZXIoJ2NoYW5nZScsICgpID0+IHtcbiAgICAgICAgdGV4dElucHV0R3JvdXAudmFsdWUgPSBjb2xvcklucHV0LnZhbHVlO1xuICAgICAgICB0ZXh0SW5wdXRHcm91cC5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnaW5wdXQnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICAgICAgICBpZiAoY29sb3JJbnB1dC5wYXJlbnROb2RlKSB7XG4gICAgICAgICAgZG9jdW1lbnQuYm9keS5yZW1vdmVDaGlsZChjb2xvcklucHV0KTtcbiAgICAgICAgfVxuICAgICAgfSk7XG5cbiAgICAgIGNvbG9ySW5wdXQuYWRkRXZlbnRMaXN0ZW5lcignYmx1cicsICgpID0+IHtcbiAgICAgICAgaWYgKGNvbG9ySW5wdXQucGFyZW50Tm9kZSkge1xuICAgICAgICAgIGRvY3VtZW50LmJvZHkucmVtb3ZlQ2hpbGQoY29sb3JJbnB1dCk7XG4gICAgICAgIH1cbiAgICAgIH0pO1xuXG4gICAgICBjb2xvcklucHV0LmNsaWNrKCk7XG4gICAgfVxuICB9O1xuXG4gICNhcHBseUNoYW5nZSh0eXBlOiBzdHJpbmcsIG5hbWU6IHN0cmluZywgdmFsdWU6IHVua25vd24pIHtcbiAgICBzd2l0Y2ggKHR5cGUpIHtcbiAgICAgIGNhc2UgJ2F0dHJpYnV0ZSc6XG4gICAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgS25vYkF0dHJpYnV0ZUNoYW5nZUV2ZW50KG5hbWUsIHZhbHVlKSk7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAncHJvcGVydHknOlxuICAgICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IEtub2JQcm9wZXJ0eUNoYW5nZUV2ZW50KG5hbWUsIHRoaXMuI3BhcnNlVmFsdWUodmFsdWUpKSk7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnY3NzLXByb3BlcnR5JzpcbiAgICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBLbm9iQ3NzUHJvcGVydHlDaGFuZ2VFdmVudChuYW1lLCB2YWx1ZSBhcyBzdHJpbmcpKTtcbiAgICAgICAgYnJlYWs7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICBjb25zb2xlLndhcm4oYFtLbm9iR3JvdXBdIFVua25vd24ga25vYiB0eXBlOiAke3R5cGV9YCk7XG4gICAgICAgIHJldHVybjtcbiAgICB9XG4gIH1cblxuICAjcGFyc2VWYWx1ZSh2YWx1ZTogdW5rbm93bik6IHVua25vd24ge1xuICAgIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgcmV0dXJuIHZhbHVlO1xuICAgIH1cblxuICAgIGlmICh2YWx1ZSA9PT0gJ3RydWUnKSByZXR1cm4gdHJ1ZTtcbiAgICBpZiAodmFsdWUgPT09ICdmYWxzZScpIHJldHVybiBmYWxzZTtcbiAgICBpZiAodmFsdWUgPT09ICdudWxsJykgcmV0dXJuIG51bGw7XG4gICAgaWYgKHZhbHVlID09PSAnJykgcmV0dXJuICcnO1xuXG4gICAgY29uc3QgbnVtID0gTnVtYmVyKHZhbHVlKTtcbiAgICBpZiAoIWlzTmFOKG51bSkgJiYgdmFsdWUgIT09ICcnKSB7XG4gICAgICByZXR1cm4gbnVtO1xuICAgIH1cblxuICAgIHJldHVybiB2YWx1ZTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tc2VydmUta25vYi1ncm91cCc6IENlbVNlcnZlS25vYkdyb3VwO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7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
