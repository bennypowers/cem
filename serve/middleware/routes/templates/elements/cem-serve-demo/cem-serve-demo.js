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

// elements/cem-serve-demo/cem-serve-demo.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-serve-demo/cem-serve-demo.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\niframe {\\n  border: none;\\n  width: 100%;\\n  height: 100%;\\n}\\n"'));
var cem_serve_demo_default = s;

// elements/cem-serve-demo/cem-serve-demo.ts
var _rendering_dec, _a, _CemServeDemo_decorators, _init, _rendering, _CemServeDemo_instances, iframe_get, iframeSrc_fn, onIframeLoad_fn, getElementInstance_fn, postKnobChange_fn, applyAttributeChange_fn, applyPropertyChange_fn, applyCSSPropertyChange_fn;
_CemServeDemo_decorators = [customElement("cem-serve-demo")];
var CemServeDemo = class extends (_a = LitElement, _rendering_dec = [property({ reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeDemo_instances);
    __privateAdd(this, _rendering, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
  }
  render() {
    return this.rendering === "iframe" ? html`
      <iframe part="iframe"
              src="${__privateMethod(this, _CemServeDemo_instances, iframeSrc_fn).call(this)}"
              @load="${__privateMethod(this, _CemServeDemo_instances, onIframeLoad_fn)}"></iframe>` : html`
      <slot></slot>`;
  }
  /**
   * Apply a knob change to an element in the demo.
   * Called by parent chrome element when knob events occur.
   * In iframe mode, bridges via postMessage instead of direct DOM access.
   */
  applyKnobChange(type, name, value, tagName, instanceIndex = 0) {
    if (this.rendering === "iframe") {
      return __privateMethod(this, _CemServeDemo_instances, postKnobChange_fn).call(this, type, name, value, tagName, instanceIndex);
    }
    const element = __privateMethod(this, _CemServeDemo_instances, getElementInstance_fn).call(this, tagName, instanceIndex);
    if (!element) {
      console.warn("[cem-serve-demo] Element not found:", tagName, "at index", instanceIndex);
      return false;
    }
    switch (type) {
      case "attribute":
        return __privateMethod(this, _CemServeDemo_instances, applyAttributeChange_fn).call(this, element, name, value);
      case "property":
        return __privateMethod(this, _CemServeDemo_instances, applyPropertyChange_fn).call(this, element, name, value);
      case "css-property":
        return __privateMethod(this, _CemServeDemo_instances, applyCSSPropertyChange_fn).call(this, element, name, value);
      default:
        console.warn("[cem-serve-demo] Unknown knob type:", type);
        return false;
    }
  }
  /**
   * Set an attribute on an element in the demo
   */
  setDemoAttribute(selector, attribute, value) {
    const target = this.querySelector(selector);
    if (!target) return false;
    if (typeof value === "boolean") {
      target.toggleAttribute(attribute, value);
    } else if (value === "" || value === null || value === void 0) {
      target.removeAttribute(attribute);
    } else {
      target.setAttribute(attribute, value);
    }
    return true;
  }
  /**
   * Set a property on an element in the demo
   */
  setDemoProperty(selector, property2, value) {
    const target = this.querySelector(selector);
    if (target) {
      target[property2] = value;
      return true;
    }
    return false;
  }
  /**
   * Set a CSS custom property on an element in the demo
   */
  setDemoCssCustomProperty(selector, cssProperty, value) {
    const target = this.querySelector(selector);
    if (target) {
      const propertyName = cssProperty.startsWith("--") ? cssProperty : `--${cssProperty}`;
      target.style.setProperty(propertyName, value);
      return true;
    }
    return false;
  }
};
_init = __decoratorStart(_a);
_rendering = new WeakMap();
_CemServeDemo_instances = new WeakSet();
iframe_get = function() {
  return this.renderRoot.querySelector("iframe");
};
iframeSrc_fn = function() {
  const url = new URL(window.location.href);
  url.searchParams.set("rendering", "chromeless");
  return url.toString();
};
onIframeLoad_fn = function() {
  this.dispatchEvent(new Event("iframe-ready"));
};
/**
 * Find the Nth instance of an element by tag name
 */
getElementInstance_fn = function(tagName, instanceIndex = 0) {
  const elements = this.querySelectorAll(tagName);
  return elements[instanceIndex] || null;
};
postKnobChange_fn = function(knobType, name, value, tagName, instanceIndex) {
  const iframe = __privateGet(this, _CemServeDemo_instances, iframe_get);
  if (!iframe?.contentWindow) {
    console.warn("[cem-serve-demo] Iframe not ready for postMessage");
    return false;
  }
  iframe.contentWindow.postMessage(
    { type: "cem-knob-change", knobType, name, value, tagName, instanceIndex },
    window.location.origin
  );
  return true;
};
applyAttributeChange_fn = function(element, name, value) {
  if (typeof value === "boolean") {
    element.toggleAttribute(name, value);
  } else if (value === "" || value === null || value === void 0) {
    element.removeAttribute(name);
  } else {
    element.setAttribute(name, String(value));
  }
  return true;
};
applyPropertyChange_fn = function(element, name, value) {
  if (value === void 0) {
    try {
      delete element[name];
    } catch {
      element[name] = value;
    }
  } else {
    element[name] = value;
  }
  return true;
};
applyCSSPropertyChange_fn = function(element, name, value) {
  const propertyName = name.startsWith("--") ? name : `--${name}`;
  if (value === "" || value === null || value === void 0) {
    element.style.removeProperty(propertyName);
  } else {
    element.style.setProperty(propertyName, String(value));
  }
  return true;
};
__decorateElement(_init, 4, "rendering", _rendering_dec, CemServeDemo, _rendering);
CemServeDemo = __decorateElement(_init, 0, "CemServeDemo", _CemServeDemo_decorators, CemServeDemo);
__publicField(CemServeDemo, "styles", cem_serve_demo_default);
__runInitializers(_init, 1, CemServeDemo);
export {
  CemServeDemo
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8udHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tc2VydmUtZGVtby9jZW0tc2VydmUtZGVtby5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWRlbW8uY3NzJztcblxuLyoqXG4gKiBEZW1vIHdyYXBwZXIgY29tcG9uZW50IGZvciBrbm9icyBpbnRlZ3JhdGlvbi5cbiAqXG4gKiBJbiBsaWdodC9zaGFkb3cgbW9kZSwgcmVuZGVycyBkZW1vIGNvbnRlbnQgdmlhIGEgZGVmYXVsdCBzbG90LlxuICogSW4gaWZyYW1lIG1vZGUsIGxvYWRzIHRoZSBkZW1vIGluIGFuIGlzb2xhdGVkIGlmcmFtZSBhbmQgYnJpZGdlc1xuICoga25vYiBjaGFuZ2VzIHZpYSBwb3N0TWVzc2FnZS5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgZGVtbyBjb250ZW50IChsaWdodC9zaGFkb3cgbW9kZSBvbmx5KVxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXNlcnZlLWRlbW9cbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1zZXJ2ZS1kZW1vJylcbmV4cG9ydCBjbGFzcyBDZW1TZXJ2ZURlbW8gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pIGFjY2Vzc29yIHJlbmRlcmluZzogc3RyaW5nIHwgdW5kZWZpbmVkO1xuXG4gIGdldCAjaWZyYW1lKCk6IEhUTUxJRnJhbWVFbGVtZW50IHwgbnVsbCB7IHJldHVybiB0aGlzLnJlbmRlclJvb3QucXVlcnlTZWxlY3RvcignaWZyYW1lJyk7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuICh0aGlzLnJlbmRlcmluZyA9PT0gJ2lmcmFtZScpID8gIGh0bWxgXG4gICAgICA8aWZyYW1lIHBhcnQ9XCJpZnJhbWVcIlxuICAgICAgICAgICAgICBzcmM9XCIke3RoaXMuI2lmcmFtZVNyYygpfVwiXG4gICAgICAgICAgICAgIEBsb2FkPVwiJHt0aGlzLiNvbklmcmFtZUxvYWR9XCI+PC9pZnJhbWU+YDogaHRtbGBcbiAgICAgIDxzbG90Pjwvc2xvdD5gO1xuICB9XG5cbiAgLyoqXG4gICAqIEFwcGx5IGEga25vYiBjaGFuZ2UgdG8gYW4gZWxlbWVudCBpbiB0aGUgZGVtby5cbiAgICogQ2FsbGVkIGJ5IHBhcmVudCBjaHJvbWUgZWxlbWVudCB3aGVuIGtub2IgZXZlbnRzIG9jY3VyLlxuICAgKiBJbiBpZnJhbWUgbW9kZSwgYnJpZGdlcyB2aWEgcG9zdE1lc3NhZ2UgaW5zdGVhZCBvZiBkaXJlY3QgRE9NIGFjY2Vzcy5cbiAgICovXG4gIGFwcGx5S25vYkNoYW5nZShcbiAgICB0eXBlOiBzdHJpbmcsXG4gICAgbmFtZTogc3RyaW5nLFxuICAgIHZhbHVlOiB1bmtub3duLFxuICAgIHRhZ05hbWU6IHN0cmluZyxcbiAgICBpbnN0YW5jZUluZGV4ID0gMCxcbiAgKTogYm9vbGVhbiB7XG4gICAgaWYgKHRoaXMucmVuZGVyaW5nID09PSAnaWZyYW1lJykge1xuICAgICAgcmV0dXJuIHRoaXMuI3Bvc3RLbm9iQ2hhbmdlKHR5cGUsIG5hbWUsIHZhbHVlLCB0YWdOYW1lLCBpbnN0YW5jZUluZGV4KTtcbiAgICB9XG5cbiAgICBjb25zdCBlbGVtZW50ID0gdGhpcy4jZ2V0RWxlbWVudEluc3RhbmNlKHRhZ05hbWUsIGluc3RhbmNlSW5kZXgpO1xuICAgIGlmICghZWxlbWVudCkge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWRlbW9dIEVsZW1lbnQgbm90IGZvdW5kOicsIHRhZ05hbWUsICdhdCBpbmRleCcsIGluc3RhbmNlSW5kZXgpO1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cblxuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnYXR0cmlidXRlJzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5QXR0cmlidXRlQ2hhbmdlKGVsZW1lbnQsIG5hbWUsIHZhbHVlKTtcbiAgICAgIGNhc2UgJ3Byb3BlcnR5JzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5UHJvcGVydHlDaGFuZ2UoXG4gICAgICAgICAgZWxlbWVudCxcbiAgICAgICAgICBuYW1lIGFzIGtleW9mIEVsZW1lbnQsXG4gICAgICAgICAgdmFsdWUgYXMgRWxlbWVudFtrZXlvZiBFbGVtZW50XSxcbiAgICAgICAgKTtcbiAgICAgIGNhc2UgJ2Nzcy1wcm9wZXJ0eSc6XG4gICAgICAgIHJldHVybiB0aGlzLiNhcHBseUNTU1Byb3BlcnR5Q2hhbmdlKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQsIG5hbWUsIHZhbHVlKTtcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1kZW1vXSBVbmtub3duIGtub2IgdHlwZTonLCB0eXBlKTtcbiAgICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBTZXQgYW4gYXR0cmlidXRlIG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICovXG4gIHNldERlbW9BdHRyaWJ1dGUoc2VsZWN0b3I6IHN0cmluZywgYXR0cmlidXRlOiBzdHJpbmcsIHZhbHVlOiBzdHJpbmcgfCBib29sZWFuIHwgbnVsbCk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMucXVlcnlTZWxlY3RvcihzZWxlY3Rvcik7XG4gICAgaWYgKCF0YXJnZXQpIHJldHVybiBmYWxzZTtcblxuICAgIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgdGFyZ2V0LnRvZ2dsZUF0dHJpYnV0ZShhdHRyaWJ1dGUsIHZhbHVlKTtcbiAgICB9IGVsc2UgaWYgKHZhbHVlID09PSAnJyB8fCB2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICB0YXJnZXQucmVtb3ZlQXR0cmlidXRlKGF0dHJpYnV0ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRhcmdldC5zZXRBdHRyaWJ1dGUoYXR0cmlidXRlLCB2YWx1ZSk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAvKipcbiAgICogU2V0IGEgcHJvcGVydHkgb24gYW4gZWxlbWVudCBpbiB0aGUgZGVtb1xuICAgKi9cbiAgc2V0RGVtb1Byb3BlcnR5PFQgZXh0ZW5kcyBFbGVtZW50PihcbiAgICBzZWxlY3Rvcjogc3RyaW5nLFxuICAgIHByb3BlcnR5OiBrZXlvZiBULFxuICAgIHZhbHVlOiBUW2tleW9mIFRdLFxuICApOiBib29sZWFuIHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLnF1ZXJ5U2VsZWN0b3I8VD4oc2VsZWN0b3IpO1xuICAgIGlmICh0YXJnZXQpIHtcbiAgICAgICh0YXJnZXQpW3Byb3BlcnR5XSA9IHZhbHVlO1xuICAgICAgcmV0dXJuIHRydWU7XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTZXQgYSBDU1MgY3VzdG9tIHByb3BlcnR5IG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICovXG4gIHNldERlbW9Dc3NDdXN0b21Qcm9wZXJ0eShzZWxlY3Rvcjogc3RyaW5nLCBjc3NQcm9wZXJ0eTogc3RyaW5nLCB2YWx1ZTogc3RyaW5nKTogYm9vbGVhbiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gdGhpcy5xdWVyeVNlbGVjdG9yKHNlbGVjdG9yKSBhcyBIVE1MRWxlbWVudCB8IG51bGw7XG4gICAgaWYgKHRhcmdldCkge1xuICAgICAgY29uc3QgcHJvcGVydHlOYW1lID0gY3NzUHJvcGVydHkuc3RhcnRzV2l0aCgnLS0nKSA/IGNzc1Byb3BlcnR5IDogYC0tJHtjc3NQcm9wZXJ0eX1gO1xuICAgICAgdGFyZ2V0LnN0eWxlLnNldFByb3BlcnR5KHByb3BlcnR5TmFtZSwgdmFsdWUpO1xuICAgICAgcmV0dXJuIHRydWU7XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbiAgfVxuXG4gICNpZnJhbWVTcmMoKTogc3RyaW5nIHtcbiAgICBjb25zdCB1cmwgPSBuZXcgVVJMKHdpbmRvdy5sb2NhdGlvbi5ocmVmKTtcbiAgICB1cmwuc2VhcmNoUGFyYW1zLnNldCgncmVuZGVyaW5nJywgJ2Nocm9tZWxlc3MnKTtcbiAgICByZXR1cm4gdXJsLnRvU3RyaW5nKCk7XG4gIH1cblxuICAjb25JZnJhbWVMb2FkKCkge1xuICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQoJ2lmcmFtZS1yZWFkeScpKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBGaW5kIHRoZSBOdGggaW5zdGFuY2Ugb2YgYW4gZWxlbWVudCBieSB0YWcgbmFtZVxuICAgKi9cbiAgI2dldEVsZW1lbnRJbnN0YW5jZSh0YWdOYW1lOiBzdHJpbmcsIGluc3RhbmNlSW5kZXggPSAwKTogRWxlbWVudCB8IG51bGwge1xuICAgIGNvbnN0IGVsZW1lbnRzID0gdGhpcy5xdWVyeVNlbGVjdG9yQWxsKHRhZ05hbWUpO1xuICAgIHJldHVybiBlbGVtZW50c1tpbnN0YW5jZUluZGV4XSB8fCBudWxsO1xuICB9XG5cbiAgI3Bvc3RLbm9iQ2hhbmdlKFxuICAgIGtub2JUeXBlOiBzdHJpbmcsXG4gICAgbmFtZTogc3RyaW5nLFxuICAgIHZhbHVlOiB1bmtub3duLFxuICAgIHRhZ05hbWU6IHN0cmluZyxcbiAgICBpbnN0YW5jZUluZGV4OiBudW1iZXIsXG4gICk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IGlmcmFtZSA9IHRoaXMuI2lmcmFtZTtcbiAgICBpZiAoIWlmcmFtZT8uY29udGVudFdpbmRvdykge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWRlbW9dIElmcmFtZSBub3QgcmVhZHkgZm9yIHBvc3RNZXNzYWdlJyk7XG4gICAgICByZXR1cm4gZmFsc2U7XG4gICAgfVxuICAgIGlmcmFtZS5jb250ZW50V2luZG93LnBvc3RNZXNzYWdlKFxuICAgICAgeyB0eXBlOiAnY2VtLWtub2ItY2hhbmdlJywga25vYlR5cGUsIG5hbWUsIHZhbHVlLCB0YWdOYW1lLCBpbnN0YW5jZUluZGV4IH0sXG4gICAgICB3aW5kb3cubG9jYXRpb24ub3JpZ2luLFxuICAgICk7XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAjYXBwbHlBdHRyaWJ1dGVDaGFuZ2UoZWxlbWVudDogRWxlbWVudCwgbmFtZTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bik6IGJvb2xlYW4ge1xuICAgIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgZWxlbWVudC50b2dnbGVBdHRyaWJ1dGUobmFtZSwgdmFsdWUpO1xuICAgIH0gZWxzZSBpZiAodmFsdWUgPT09ICcnIHx8IHZhbHVlID09PSBudWxsIHx8IHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgIGVsZW1lbnQucmVtb3ZlQXR0cmlidXRlKG5hbWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICBlbGVtZW50LnNldEF0dHJpYnV0ZShuYW1lLCBTdHJpbmcodmFsdWUpKTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAjYXBwbHlQcm9wZXJ0eUNoYW5nZTxUIGV4dGVuZHMgRWxlbWVudD4oXG4gICAgZWxlbWVudDogVCxcbiAgICBuYW1lOiBrZXlvZiBULFxuICAgIHZhbHVlOiBUW2tleW9mIFRdLFxuICApOiBib29sZWFuIHtcbiAgICBpZiAodmFsdWUgPT09IHVuZGVmaW5lZCkge1xuICAgICAgdHJ5IHtcbiAgICAgICAgZGVsZXRlIChlbGVtZW50KVtuYW1lXTtcbiAgICAgIH0gY2F0Y2gge1xuICAgICAgICAoZWxlbWVudClbbmFtZV0gPSB2YWx1ZTtcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgKGVsZW1lbnQpW25hbWVdID0gdmFsdWU7XG4gICAgfVxuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgI2FwcGx5Q1NTUHJvcGVydHlDaGFuZ2UoZWxlbWVudDogSFRNTEVsZW1lbnQsIG5hbWU6IHN0cmluZywgdmFsdWU6IHVua25vd24pOiBib29sZWFuIHtcbiAgICBjb25zdCBwcm9wZXJ0eU5hbWUgPSBuYW1lLnN0YXJ0c1dpdGgoJy0tJykgPyBuYW1lIDogYC0tJHtuYW1lfWA7XG4gICAgaWYgKHZhbHVlID09PSAnJyB8fCB2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICBlbGVtZW50LnN0eWxlLnJlbW92ZVByb3BlcnR5KHByb3BlcnR5TmFtZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIGVsZW1lbnQuc3R5bGUuc2V0UHJvcGVydHkocHJvcGVydHlOYW1lLCBTdHJpbmcodmFsdWUpKTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXNlcnZlLWRlbW8nOiBDZW1TZXJ2ZURlbW87XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG59XFxcXG5cXFxcbmlmcmFtZSB7XFxcXG4gIGJvcmRlcjogbm9uZTtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlB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T0FBSyxjQUFjLElBQUksTUFBTSxjQUFjLENBQUM7QUFDOUM7QUFBQTtBQUFBO0FBQUE7QUFLQSx3QkFBbUIsU0FBQyxTQUFpQixnQkFBZ0IsR0FBbUI7QUFDdEUsUUFBTSxXQUFXLEtBQUssaUJBQWlCLE9BQU87QUFDOUMsU0FBTyxTQUFTLGFBQWEsS0FBSztBQUNwQztBQUVBLG9CQUFlLFNBQ2IsVUFDQSxNQUNBLE9BQ0EsU0FDQSxlQUNTO0FBQ1QsUUFBTSxTQUFTLG1CQUFLO0FBQ3BCLE1BQUksQ0FBQyxRQUFRLGVBQWU7QUFDMUIsWUFBUSxLQUFLLG1EQUFtRDtBQUNoRSxXQUFPO0FBQUEsRUFDVDtBQUNBLFNBQU8sY0FBYztBQUFBLElBQ25CLEVBQUUsTUFBTSxtQkFBbUIsVUFBVSxNQUFNLE9BQU8sU0FBUyxjQUFjO0FBQUEsSUFDekUsT0FBTyxTQUFTO0FBQUEsRUFDbEI7QUFDQSxTQUFPO0FBQ1Q7QUFFQSwwQkFBcUIsU0FBQyxTQUFrQixNQUFjLE9BQXlCO0FBQzdFLE1BQUksT0FBTyxVQUFVLFdBQVc7QUFDOUIsWUFBUSxnQkFBZ0IsTUFBTSxLQUFLO0FBQUEsRUFDckMsV0FBVyxVQUFVLE1BQU0sVUFBVSxRQUFRLFVBQVUsUUFBVztBQUNoRSxZQUFRLGdCQUFnQixJQUFJO0FBQUEsRUFDOUIsT0FBTztBQUNMLFlBQVEsYUFBYSxNQUFNLE9BQU8sS0FBSyxDQUFDO0FBQUEsRUFDMUM7QUFDQSxTQUFPO0FBQ1Q7QUFFQSx5QkFBdUMsU0FDckMsU0FDQSxNQUNBLE9BQ1M7QUFDVCxNQUFJLFVBQVUsUUFBVztBQUN2QixRQUFJO0FBQ0YsYUFBUSxRQUFTLElBQUk7QUFBQSxJQUN2QixRQUFRO0FBQ04sTUFBQyxRQUFTLElBQUksSUFBSTtBQUFBLElBQ3BCO0FBQUEsRUFDRixPQUFPO0FBQ0wsSUFBQyxRQUFTLElBQUksSUFBSTtBQUFBLEVBQ3BCO0FBQ0EsU0FBTztBQUNUO0FBRUEsNEJBQXVCLFNBQUMsU0FBc0IsTUFBYyxPQUF5QjtBQUNuRixRQUFNLGVBQWUsS0FBSyxXQUFXLElBQUksSUFBSSxPQUFPLEtBQUssSUFBSTtBQUM3RCxNQUFJLFVBQVUsTUFBTSxVQUFVLFFBQVEsVUFBVSxRQUFXO0FBQ3pELFlBQVEsTUFBTSxlQUFlLFlBQVk7QUFBQSxFQUMzQyxPQUFPO0FBQ0wsWUFBUSxNQUFNLFlBQVksY0FBYyxPQUFPLEtBQUssQ0FBQztBQUFBLEVBQ3ZEO0FBQ0EsU0FBTztBQUNUO0FBNUs2Qiw0QkFBUyxhQUF0QyxnQkFIVyxjQUcyQjtBQUgzQixlQUFOLDRDQURQLDBCQUNhO0FBQ1gsY0FEVyxjQUNKLFVBQVM7QUFEWCw0QkFBTTsiLAogICJuYW1lcyI6IFsicHJvcGVydHkiXQp9Cg==
