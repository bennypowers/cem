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
var _rendering_dec, _a, _CemServeDemo_decorators, _init, _rendering, _iframeReady, _pendingMessages, _CemServeDemo_instances, iframe_get, iframeSrc_fn, onIframeLoad_fn, getElementInstance_fn, postKnobChange_fn, applyAttributeChange_fn, applyPropertyChange_fn, applyCSSPropertyChange_fn;
_CemServeDemo_decorators = [customElement("cem-serve-demo")];
var CemServeDemo = class extends (_a = LitElement, _rendering_dec = [property({ reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeDemo_instances);
    __privateAdd(this, _rendering, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
    __privateAdd(this, _iframeReady, false);
    __privateAdd(this, _pendingMessages, []);
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
_iframeReady = new WeakMap();
_pendingMessages = new WeakMap();
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
  __privateSet(this, _iframeReady, true);
  for (const msg of __privateGet(this, _pendingMessages))
    __privateGet(this, _CemServeDemo_instances, iframe_get)?.contentWindow?.postMessage(msg, window.location.origin);
  __privateSet(this, _pendingMessages, []);
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
  const msg = { type: "cem-knob-change", knobType, name, value, tagName, instanceIndex };
  if (!__privateGet(this, _iframeReady)) {
    __privateGet(this, _pendingMessages).push(msg);
    return true;
  }
  const iframe = __privateGet(this, _CemServeDemo_instances, iframe_get);
  if (!iframe?.contentWindow) {
    console.warn("[cem-serve-demo] Iframe not ready for postMessage");
    return false;
  }
  iframe.contentWindow.postMessage(msg, window.location.origin);
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8udHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tc2VydmUtZGVtby9jZW0tc2VydmUtZGVtby5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWRlbW8uY3NzJztcblxuLyoqXG4gKiBEZW1vIHdyYXBwZXIgY29tcG9uZW50IGZvciBrbm9icyBpbnRlZ3JhdGlvbi5cbiAqXG4gKiBJbiBsaWdodC9zaGFkb3cgbW9kZSwgcmVuZGVycyBkZW1vIGNvbnRlbnQgdmlhIGEgZGVmYXVsdCBzbG90LlxuICogSW4gaWZyYW1lIG1vZGUsIGxvYWRzIHRoZSBkZW1vIGluIGFuIGlzb2xhdGVkIGlmcmFtZSBhbmQgYnJpZGdlc1xuICoga25vYiBjaGFuZ2VzIHZpYSBwb3N0TWVzc2FnZS5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgZGVtbyBjb250ZW50IChsaWdodC9zaGFkb3cgbW9kZSBvbmx5KVxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXNlcnZlLWRlbW9cbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1zZXJ2ZS1kZW1vJylcbmV4cG9ydCBjbGFzcyBDZW1TZXJ2ZURlbW8gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pIGFjY2Vzc29yIHJlbmRlcmluZzogc3RyaW5nIHwgdW5kZWZpbmVkO1xuXG4gICNpZnJhbWVSZWFkeSA9IGZhbHNlO1xuICAjcGVuZGluZ01lc3NhZ2VzOiBBcnJheTxSZWNvcmQ8c3RyaW5nLCB1bmtub3duPj4gPSBbXTtcblxuICBnZXQgI2lmcmFtZSgpOiBIVE1MSUZyYW1lRWxlbWVudCB8IG51bGwge1xuICAgIHJldHVybiB0aGlzLnJlbmRlclJvb3QucXVlcnlTZWxlY3RvcignaWZyYW1lJyk7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuICh0aGlzLnJlbmRlcmluZyA9PT0gJ2lmcmFtZScpID8gIGh0bWxgXG4gICAgICA8aWZyYW1lIHBhcnQ9XCJpZnJhbWVcIlxuICAgICAgICAgICAgICBzcmM9XCIke3RoaXMuI2lmcmFtZVNyYygpfVwiXG4gICAgICAgICAgICAgIEBsb2FkPVwiJHt0aGlzLiNvbklmcmFtZUxvYWR9XCI+PC9pZnJhbWU+YDogaHRtbGBcbiAgICAgIDxzbG90Pjwvc2xvdD5gO1xuICB9XG5cbiAgLyoqXG4gICAqIEFwcGx5IGEga25vYiBjaGFuZ2UgdG8gYW4gZWxlbWVudCBpbiB0aGUgZGVtby5cbiAgICogQ2FsbGVkIGJ5IHBhcmVudCBjaHJvbWUgZWxlbWVudCB3aGVuIGtub2IgZXZlbnRzIG9jY3VyLlxuICAgKiBJbiBpZnJhbWUgbW9kZSwgYnJpZGdlcyB2aWEgcG9zdE1lc3NhZ2UgaW5zdGVhZCBvZiBkaXJlY3QgRE9NIGFjY2Vzcy5cbiAgICovXG4gIGFwcGx5S25vYkNoYW5nZShcbiAgICB0eXBlOiBzdHJpbmcsXG4gICAgbmFtZTogc3RyaW5nLFxuICAgIHZhbHVlOiB1bmtub3duLFxuICAgIHRhZ05hbWU6IHN0cmluZyxcbiAgICBpbnN0YW5jZUluZGV4ID0gMCxcbiAgKTogYm9vbGVhbiB7XG4gICAgaWYgKHRoaXMucmVuZGVyaW5nID09PSAnaWZyYW1lJykge1xuICAgICAgcmV0dXJuIHRoaXMuI3Bvc3RLbm9iQ2hhbmdlKHR5cGUsIG5hbWUsIHZhbHVlLCB0YWdOYW1lLCBpbnN0YW5jZUluZGV4KTtcbiAgICB9XG5cbiAgICBjb25zdCBlbGVtZW50ID0gdGhpcy4jZ2V0RWxlbWVudEluc3RhbmNlKHRhZ05hbWUsIGluc3RhbmNlSW5kZXgpO1xuICAgIGlmICghZWxlbWVudCkge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWRlbW9dIEVsZW1lbnQgbm90IGZvdW5kOicsIHRhZ05hbWUsICdhdCBpbmRleCcsIGluc3RhbmNlSW5kZXgpO1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cblxuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnYXR0cmlidXRlJzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5QXR0cmlidXRlQ2hhbmdlKGVsZW1lbnQsIG5hbWUsIHZhbHVlKTtcbiAgICAgIGNhc2UgJ3Byb3BlcnR5JzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5UHJvcGVydHlDaGFuZ2UoXG4gICAgICAgICAgZWxlbWVudCxcbiAgICAgICAgICBuYW1lIGFzIGtleW9mIEVsZW1lbnQsXG4gICAgICAgICAgdmFsdWUgYXMgRWxlbWVudFtrZXlvZiBFbGVtZW50XSxcbiAgICAgICAgKTtcbiAgICAgIGNhc2UgJ2Nzcy1wcm9wZXJ0eSc6XG4gICAgICAgIHJldHVybiB0aGlzLiNhcHBseUNTU1Byb3BlcnR5Q2hhbmdlKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQsIG5hbWUsIHZhbHVlKTtcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1kZW1vXSBVbmtub3duIGtub2IgdHlwZTonLCB0eXBlKTtcbiAgICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBTZXQgYW4gYXR0cmlidXRlIG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICovXG4gIHNldERlbW9BdHRyaWJ1dGUoc2VsZWN0b3I6IHN0cmluZywgYXR0cmlidXRlOiBzdHJpbmcsIHZhbHVlOiBzdHJpbmcgfCBib29sZWFuIHwgbnVsbCk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMucXVlcnlTZWxlY3RvcihzZWxlY3Rvcik7XG4gICAgaWYgKCF0YXJnZXQpIHJldHVybiBmYWxzZTtcblxuICAgIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgdGFyZ2V0LnRvZ2dsZUF0dHJpYnV0ZShhdHRyaWJ1dGUsIHZhbHVlKTtcbiAgICB9IGVsc2UgaWYgKHZhbHVlID09PSAnJyB8fCB2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICB0YXJnZXQucmVtb3ZlQXR0cmlidXRlKGF0dHJpYnV0ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRhcmdldC5zZXRBdHRyaWJ1dGUoYXR0cmlidXRlLCB2YWx1ZSk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAvKipcbiAgICogU2V0IGEgcHJvcGVydHkgb24gYW4gZWxlbWVudCBpbiB0aGUgZGVtb1xuICAgKi9cbiAgc2V0RGVtb1Byb3BlcnR5PFQgZXh0ZW5kcyBFbGVtZW50PihcbiAgICBzZWxlY3Rvcjogc3RyaW5nLFxuICAgIHByb3BlcnR5OiBrZXlvZiBULFxuICAgIHZhbHVlOiBUW2tleW9mIFRdLFxuICApOiBib29sZWFuIHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLnF1ZXJ5U2VsZWN0b3I8VD4oc2VsZWN0b3IpO1xuICAgIGlmICh0YXJnZXQpIHtcbiAgICAgICh0YXJnZXQpW3Byb3BlcnR5XSA9IHZhbHVlO1xuICAgICAgcmV0dXJuIHRydWU7XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTZXQgYSBDU1MgY3VzdG9tIHByb3BlcnR5IG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICovXG4gIHNldERlbW9Dc3NDdXN0b21Qcm9wZXJ0eShzZWxlY3Rvcjogc3RyaW5nLCBjc3NQcm9wZXJ0eTogc3RyaW5nLCB2YWx1ZTogc3RyaW5nKTogYm9vbGVhbiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gdGhpcy5xdWVyeVNlbGVjdG9yKHNlbGVjdG9yKSBhcyBIVE1MRWxlbWVudCB8IG51bGw7XG4gICAgaWYgKHRhcmdldCkge1xuICAgICAgY29uc3QgcHJvcGVydHlOYW1lID0gY3NzUHJvcGVydHkuc3RhcnRzV2l0aCgnLS0nKSA/IGNzc1Byb3BlcnR5IDogYC0tJHtjc3NQcm9wZXJ0eX1gO1xuICAgICAgdGFyZ2V0LnN0eWxlLnNldFByb3BlcnR5KHByb3BlcnR5TmFtZSwgdmFsdWUpO1xuICAgICAgcmV0dXJuIHRydWU7XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbiAgfVxuXG4gICNpZnJhbWVTcmMoKTogc3RyaW5nIHtcbiAgICBjb25zdCB1cmwgPSBuZXcgVVJMKHdpbmRvdy5sb2NhdGlvbi5ocmVmKTtcbiAgICB1cmwuc2VhcmNoUGFyYW1zLnNldCgncmVuZGVyaW5nJywgJ2Nocm9tZWxlc3MnKTtcbiAgICByZXR1cm4gdXJsLnRvU3RyaW5nKCk7XG4gIH1cblxuICAjb25JZnJhbWVMb2FkKCkge1xuICAgIHRoaXMuI2lmcmFtZVJlYWR5ID0gdHJ1ZTtcbiAgICBmb3IgKGNvbnN0IG1zZyBvZiB0aGlzLiNwZW5kaW5nTWVzc2FnZXMpXG4gICAgICB0aGlzLiNpZnJhbWU/LmNvbnRlbnRXaW5kb3c/LnBvc3RNZXNzYWdlKG1zZywgd2luZG93LmxvY2F0aW9uLm9yaWdpbik7XG4gICAgdGhpcy4jcGVuZGluZ01lc3NhZ2VzID0gW107XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnaWZyYW1lLXJlYWR5JykpO1xuICB9XG5cbiAgLyoqXG4gICAqIEZpbmQgdGhlIE50aCBpbnN0YW5jZSBvZiBhbiBlbGVtZW50IGJ5IHRhZyBuYW1lXG4gICAqL1xuICAjZ2V0RWxlbWVudEluc3RhbmNlKHRhZ05hbWU6IHN0cmluZywgaW5zdGFuY2VJbmRleCA9IDApOiBFbGVtZW50IHwgbnVsbCB7XG4gICAgY29uc3QgZWxlbWVudHMgPSB0aGlzLnF1ZXJ5U2VsZWN0b3JBbGwodGFnTmFtZSk7XG4gICAgcmV0dXJuIGVsZW1lbnRzW2luc3RhbmNlSW5kZXhdIHx8IG51bGw7XG4gIH1cblxuICAjcG9zdEtub2JDaGFuZ2UoXG4gICAga25vYlR5cGU6IHN0cmluZyxcbiAgICBuYW1lOiBzdHJpbmcsXG4gICAgdmFsdWU6IHVua25vd24sXG4gICAgdGFnTmFtZTogc3RyaW5nLFxuICAgIGluc3RhbmNlSW5kZXg6IG51bWJlcixcbiAgKTogYm9vbGVhbiB7XG4gICAgY29uc3QgbXNnID0geyB0eXBlOiAnY2VtLWtub2ItY2hhbmdlJywga25vYlR5cGUsIG5hbWUsIHZhbHVlLCB0YWdOYW1lLCBpbnN0YW5jZUluZGV4IH07XG4gICAgaWYgKCF0aGlzLiNpZnJhbWVSZWFkeSkge1xuICAgICAgdGhpcy4jcGVuZGluZ01lc3NhZ2VzLnB1c2gobXNnKTtcbiAgICAgIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICBjb25zdCBpZnJhbWUgPSB0aGlzLiNpZnJhbWU7XG4gICAgaWYgKCFpZnJhbWU/LmNvbnRlbnRXaW5kb3cpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1kZW1vXSBJZnJhbWUgbm90IHJlYWR5IGZvciBwb3N0TWVzc2FnZScpO1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgICBpZnJhbWUuY29udGVudFdpbmRvdy5wb3N0TWVzc2FnZShtc2csIHdpbmRvdy5sb2NhdGlvbi5vcmlnaW4pO1xuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgI2FwcGx5QXR0cmlidXRlQ2hhbmdlKGVsZW1lbnQ6IEVsZW1lbnQsIG5hbWU6IHN0cmluZywgdmFsdWU6IHVua25vd24pOiBib29sZWFuIHtcbiAgICBpZiAodHlwZW9mIHZhbHVlID09PSAnYm9vbGVhbicpIHtcbiAgICAgIGVsZW1lbnQudG9nZ2xlQXR0cmlidXRlKG5hbWUsIHZhbHVlKTtcbiAgICB9IGVsc2UgaWYgKHZhbHVlID09PSAnJyB8fCB2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICBlbGVtZW50LnJlbW92ZUF0dHJpYnV0ZShuYW1lKTtcbiAgICB9IGVsc2Uge1xuICAgICAgZWxlbWVudC5zZXRBdHRyaWJ1dGUobmFtZSwgU3RyaW5nKHZhbHVlKSk7XG4gICAgfVxuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgI2FwcGx5UHJvcGVydHlDaGFuZ2U8VCBleHRlbmRzIEVsZW1lbnQ+KFxuICAgIGVsZW1lbnQ6IFQsXG4gICAgbmFtZToga2V5b2YgVCxcbiAgICB2YWx1ZTogVFtrZXlvZiBUXSxcbiAgKTogYm9vbGVhbiB7XG4gICAgaWYgKHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgIHRyeSB7XG4gICAgICAgIGRlbGV0ZSAoZWxlbWVudClbbmFtZV07XG4gICAgICB9IGNhdGNoIHtcbiAgICAgICAgKGVsZW1lbnQpW25hbWVdID0gdmFsdWU7XG4gICAgICB9XG4gICAgfSBlbHNlIHtcbiAgICAgIChlbGVtZW50KVtuYW1lXSA9IHZhbHVlO1xuICAgIH1cbiAgICByZXR1cm4gdHJ1ZTtcbiAgfVxuXG4gICNhcHBseUNTU1Byb3BlcnR5Q2hhbmdlKGVsZW1lbnQ6IEhUTUxFbGVtZW50LCBuYW1lOiBzdHJpbmcsIHZhbHVlOiB1bmtub3duKTogYm9vbGVhbiB7XG4gICAgY29uc3QgcHJvcGVydHlOYW1lID0gbmFtZS5zdGFydHNXaXRoKCctLScpID8gbmFtZSA6IGAtLSR7bmFtZX1gO1xuICAgIGlmICh2YWx1ZSA9PT0gJycgfHwgdmFsdWUgPT09IG51bGwgfHwgdmFsdWUgPT09IHVuZGVmaW5lZCkge1xuICAgICAgZWxlbWVudC5zdHlsZS5yZW1vdmVQcm9wZXJ0eShwcm9wZXJ0eU5hbWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICBlbGVtZW50LnN0eWxlLnNldFByb3BlcnR5KHByb3BlcnR5TmFtZSwgU3RyaW5nKHZhbHVlKSk7XG4gICAgfVxuICAgIHJldHVybiB0cnVlO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1zZXJ2ZS1kZW1vJzogQ2VtU2VydmVEZW1vO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxufVxcXFxuXFxcXG5pZnJhbWUge1xcXFxuICBib3JkZXI6IG5vbmU7XFxcXG4gIHdpZHRoOiAxMDAlO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sNkdBQStHLENBQUM7QUFBRSxJQUFPLHlCQUFROzs7QURBdEw7QUFnQkEsNEJBQUMsY0FBYyxnQkFBZ0I7QUFDeEIsSUFBTSxlQUFOLGVBQTJCLGlCQUdoQyxrQkFBQyxTQUFTLEVBQUUsU0FBUyxLQUFLLENBQUMsSUFISyxJQUFXO0FBQUEsRUFBdEM7QUFBQTtBQUFBO0FBR3dCLHVCQUFTLFlBQVQ7QUFFN0IscUNBQWU7QUFDZix5Q0FBbUQsQ0FBQztBQUFBO0FBQUEsRUFNcEQsU0FBUztBQUNQLFdBQVEsS0FBSyxjQUFjLFdBQWE7QUFBQTtBQUFBLHFCQUV2QixzQkFBSyx1Q0FBTCxVQUFpQjtBQUFBLHVCQUNmLHNCQUFLLHlDQUFhLGdCQUFlO0FBQUE7QUFBQSxFQUV0RDtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQU9BLGdCQUNFLE1BQ0EsTUFDQSxPQUNBLFNBQ0EsZ0JBQWdCLEdBQ1A7QUFDVCxRQUFJLEtBQUssY0FBYyxVQUFVO0FBQy9CLGFBQU8s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
