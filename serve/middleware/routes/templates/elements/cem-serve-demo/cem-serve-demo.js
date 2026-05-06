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
var _rendering_dec, _a, _CemServeDemo_decorators, _init, _rendering, _iframeReady, _pendingMessages, _CemServeDemo_instances, iframe_get, iframeSrc_fn, _onChildReady, getElementInstance_fn, postKnobChange_fn, applyAttributeChange_fn, applyPropertyChange_fn, applyCSSPropertyChange_fn;
_CemServeDemo_decorators = [customElement("cem-serve-demo")];
var CemServeDemo = class extends (_a = LitElement, _rendering_dec = [property({ reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeDemo_instances);
    __privateAdd(this, _rendering, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
    __privateAdd(this, _iframeReady, false);
    __privateAdd(this, _pendingMessages, []);
    __privateAdd(this, _onChildReady, (event) => {
      if (event.origin !== window.location.origin) return;
      if (event.data?.type !== "cem-iframe-ready") return;
      __privateSet(this, _iframeReady, true);
      const cw = __privateGet(this, _CemServeDemo_instances, iframe_get)?.contentWindow;
      if (cw) {
        for (const msg of __privateGet(this, _pendingMessages))
          cw.postMessage(msg, window.location.origin);
      }
      __privateSet(this, _pendingMessages, []);
      this.dispatchEvent(new Event("iframe-ready"));
    });
  }
  connectedCallback() {
    super.connectedCallback();
    window.addEventListener("message", __privateGet(this, _onChildReady));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener("message", __privateGet(this, _onChildReady));
  }
  render() {
    if (this.rendering === "iframe") {
      __privateSet(this, _iframeReady, false);
      return html`<iframe part="iframe"
                          src="${__privateMethod(this, _CemServeDemo_instances, iframeSrc_fn).call(this)}"></iframe>`;
    }
    return html`<slot></slot>`;
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
_onChildReady = new WeakMap();
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8udHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tc2VydmUtZGVtby9jZW0tc2VydmUtZGVtby5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWRlbW8uY3NzJztcblxuLyoqXG4gKiBEZW1vIHdyYXBwZXIgY29tcG9uZW50IGZvciBrbm9icyBpbnRlZ3JhdGlvbi5cbiAqXG4gKiBJbiBsaWdodC9zaGFkb3cgbW9kZSwgcmVuZGVycyBkZW1vIGNvbnRlbnQgdmlhIGEgZGVmYXVsdCBzbG90LlxuICogSW4gaWZyYW1lIG1vZGUsIGxvYWRzIHRoZSBkZW1vIGluIGFuIGlzb2xhdGVkIGlmcmFtZSBhbmQgYnJpZGdlc1xuICoga25vYiBjaGFuZ2VzIHZpYSBwb3N0TWVzc2FnZS5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgZGVtbyBjb250ZW50IChsaWdodC9zaGFkb3cgbW9kZSBvbmx5KVxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXNlcnZlLWRlbW9cbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1zZXJ2ZS1kZW1vJylcbmV4cG9ydCBjbGFzcyBDZW1TZXJ2ZURlbW8gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pIGFjY2Vzc29yIHJlbmRlcmluZzogc3RyaW5nIHwgdW5kZWZpbmVkO1xuXG4gICNpZnJhbWVSZWFkeSA9IGZhbHNlO1xuICAjcGVuZGluZ01lc3NhZ2VzOiBBcnJheTxSZWNvcmQ8c3RyaW5nLCB1bmtub3duPj4gPSBbXTtcblxuICBnZXQgI2lmcmFtZSgpOiBIVE1MSUZyYW1lRWxlbWVudCB8IG51bGwge1xuICAgIHJldHVybiB0aGlzLnJlbmRlclJvb3QucXVlcnlTZWxlY3RvcignaWZyYW1lJyk7XG4gIH1cblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHdpbmRvdy5hZGRFdmVudExpc3RlbmVyKCdtZXNzYWdlJywgdGhpcy4jb25DaGlsZFJlYWR5KTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgd2luZG93LnJlbW92ZUV2ZW50TGlzdGVuZXIoJ21lc3NhZ2UnLCB0aGlzLiNvbkNoaWxkUmVhZHkpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIGlmICh0aGlzLnJlbmRlcmluZyA9PT0gJ2lmcmFtZScpIHtcbiAgICAgIHRoaXMuI2lmcmFtZVJlYWR5ID0gZmFsc2U7XG4gICAgICByZXR1cm4gaHRtbGA8aWZyYW1lIHBhcnQ9XCJpZnJhbWVcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICBzcmM9XCIke3RoaXMuI2lmcmFtZVNyYygpfVwiPjwvaWZyYW1lPmA7XG4gICAgfVxuICAgIHJldHVybiBodG1sYDxzbG90Pjwvc2xvdD5gO1xuICB9XG5cbiAgLyoqXG4gICAqIEFwcGx5IGEga25vYiBjaGFuZ2UgdG8gYW4gZWxlbWVudCBpbiB0aGUgZGVtby5cbiAgICogQ2FsbGVkIGJ5IHBhcmVudCBjaHJvbWUgZWxlbWVudCB3aGVuIGtub2IgZXZlbnRzIG9jY3VyLlxuICAgKiBJbiBpZnJhbWUgbW9kZSwgYnJpZGdlcyB2aWEgcG9zdE1lc3NhZ2UgaW5zdGVhZCBvZiBkaXJlY3QgRE9NIGFjY2Vzcy5cbiAgICovXG4gIGFwcGx5S25vYkNoYW5nZShcbiAgICB0eXBlOiBzdHJpbmcsXG4gICAgbmFtZTogc3RyaW5nLFxuICAgIHZhbHVlOiB1bmtub3duLFxuICAgIHRhZ05hbWU6IHN0cmluZyxcbiAgICBpbnN0YW5jZUluZGV4ID0gMCxcbiAgKTogYm9vbGVhbiB7XG4gICAgaWYgKHRoaXMucmVuZGVyaW5nID09PSAnaWZyYW1lJykge1xuICAgICAgcmV0dXJuIHRoaXMuI3Bvc3RLbm9iQ2hhbmdlKHR5cGUsIG5hbWUsIHZhbHVlLCB0YWdOYW1lLCBpbnN0YW5jZUluZGV4KTtcbiAgICB9XG5cbiAgICBjb25zdCBlbGVtZW50ID0gdGhpcy4jZ2V0RWxlbWVudEluc3RhbmNlKHRhZ05hbWUsIGluc3RhbmNlSW5kZXgpO1xuICAgIGlmICghZWxlbWVudCkge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWRlbW9dIEVsZW1lbnQgbm90IGZvdW5kOicsIHRhZ05hbWUsICdhdCBpbmRleCcsIGluc3RhbmNlSW5kZXgpO1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cblxuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnYXR0cmlidXRlJzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5QXR0cmlidXRlQ2hhbmdlKGVsZW1lbnQsIG5hbWUsIHZhbHVlKTtcbiAgICAgIGNhc2UgJ3Byb3BlcnR5JzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5UHJvcGVydHlDaGFuZ2UoXG4gICAgICAgICAgZWxlbWVudCxcbiAgICAgICAgICBuYW1lIGFzIGtleW9mIEVsZW1lbnQsXG4gICAgICAgICAgdmFsdWUgYXMgRWxlbWVudFtrZXlvZiBFbGVtZW50XSxcbiAgICAgICAgKTtcbiAgICAgIGNhc2UgJ2Nzcy1wcm9wZXJ0eSc6XG4gICAgICAgIHJldHVybiB0aGlzLiNhcHBseUNTU1Byb3BlcnR5Q2hhbmdlKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQsIG5hbWUsIHZhbHVlKTtcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1kZW1vXSBVbmtub3duIGtub2IgdHlwZTonLCB0eXBlKTtcbiAgICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBTZXQgYW4gYXR0cmlidXRlIG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICovXG4gIHNldERlbW9BdHRyaWJ1dGUoc2VsZWN0b3I6IHN0cmluZywgYXR0cmlidXRlOiBzdHJpbmcsIHZhbHVlOiBzdHJpbmcgfCBib29sZWFuIHwgbnVsbCk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMucXVlcnlTZWxlY3RvcihzZWxlY3Rvcik7XG4gICAgaWYgKCF0YXJnZXQpIHJldHVybiBmYWxzZTtcblxuICAgIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgdGFyZ2V0LnRvZ2dsZUF0dHJpYnV0ZShhdHRyaWJ1dGUsIHZhbHVlKTtcbiAgICB9IGVsc2UgaWYgKHZhbHVlID09PSAnJyB8fCB2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICB0YXJnZXQucmVtb3ZlQXR0cmlidXRlKGF0dHJpYnV0ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRhcmdldC5zZXRBdHRyaWJ1dGUoYXR0cmlidXRlLCB2YWx1ZSk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAvKipcbiAgICogU2V0IGEgcHJvcGVydHkgb24gYW4gZWxlbWVudCBpbiB0aGUgZGVtb1xuICAgKi9cbiAgc2V0RGVtb1Byb3BlcnR5PFQgZXh0ZW5kcyBFbGVtZW50PihcbiAgICBzZWxlY3Rvcjogc3RyaW5nLFxuICAgIHByb3BlcnR5OiBrZXlvZiBULFxuICAgIHZhbHVlOiBUW2tleW9mIFRdLFxuICApOiBib29sZWFuIHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLnF1ZXJ5U2VsZWN0b3I8VD4oc2VsZWN0b3IpO1xuICAgIGlmICh0YXJnZXQpIHtcbiAgICAgICh0YXJnZXQpW3Byb3BlcnR5XSA9IHZhbHVlO1xuICAgICAgcmV0dXJuIHRydWU7XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTZXQgYSBDU1MgY3VzdG9tIHByb3BlcnR5IG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICovXG4gIHNldERlbW9Dc3NDdXN0b21Qcm9wZXJ0eShzZWxlY3Rvcjogc3RyaW5nLCBjc3NQcm9wZXJ0eTogc3RyaW5nLCB2YWx1ZTogc3RyaW5nKTogYm9vbGVhbiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gdGhpcy5xdWVyeVNlbGVjdG9yKHNlbGVjdG9yKSBhcyBIVE1MRWxlbWVudCB8IG51bGw7XG4gICAgaWYgKHRhcmdldCkge1xuICAgICAgY29uc3QgcHJvcGVydHlOYW1lID0gY3NzUHJvcGVydHkuc3RhcnRzV2l0aCgnLS0nKSA/IGNzc1Byb3BlcnR5IDogYC0tJHtjc3NQcm9wZXJ0eX1gO1xuICAgICAgdGFyZ2V0LnN0eWxlLnNldFByb3BlcnR5KHByb3BlcnR5TmFtZSwgdmFsdWUpO1xuICAgICAgcmV0dXJuIHRydWU7XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbiAgfVxuXG4gICNpZnJhbWVTcmMoKTogc3RyaW5nIHtcbiAgICBjb25zdCB1cmwgPSBuZXcgVVJMKHdpbmRvdy5sb2NhdGlvbi5ocmVmKTtcbiAgICB1cmwuc2VhcmNoUGFyYW1zLnNldCgncmVuZGVyaW5nJywgJ2Nocm9tZWxlc3MnKTtcbiAgICByZXR1cm4gdXJsLnRvU3RyaW5nKCk7XG4gIH1cblxuICAjb25DaGlsZFJlYWR5ID0gKGV2ZW50OiBNZXNzYWdlRXZlbnQpID0+IHtcbiAgICBpZiAoZXZlbnQub3JpZ2luICE9PSB3aW5kb3cubG9jYXRpb24ub3JpZ2luKSByZXR1cm47XG4gICAgaWYgKGV2ZW50LmRhdGE/LnR5cGUgIT09ICdjZW0taWZyYW1lLXJlYWR5JykgcmV0dXJuO1xuICAgIHRoaXMuI2lmcmFtZVJlYWR5ID0gdHJ1ZTtcbiAgICBjb25zdCBjdyA9IHRoaXMuI2lmcmFtZT8uY29udGVudFdpbmRvdztcbiAgICBpZiAoY3cpIHtcbiAgICAgIGZvciAoY29uc3QgbXNnIG9mIHRoaXMuI3BlbmRpbmdNZXNzYWdlcylcbiAgICAgICAgY3cucG9zdE1lc3NhZ2UobXNnLCB3aW5kb3cubG9jYXRpb24ub3JpZ2luKTtcbiAgICB9XG4gICAgdGhpcy4jcGVuZGluZ01lc3NhZ2VzID0gW107XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnaWZyYW1lLXJlYWR5JykpO1xuICB9O1xuXG4gIC8qKlxuICAgKiBGaW5kIHRoZSBOdGggaW5zdGFuY2Ugb2YgYW4gZWxlbWVudCBieSB0YWcgbmFtZVxuICAgKi9cbiAgI2dldEVsZW1lbnRJbnN0YW5jZSh0YWdOYW1lOiBzdHJpbmcsIGluc3RhbmNlSW5kZXggPSAwKTogRWxlbWVudCB8IG51bGwge1xuICAgIGNvbnN0IGVsZW1lbnRzID0gdGhpcy5xdWVyeVNlbGVjdG9yQWxsKHRhZ05hbWUpO1xuICAgIHJldHVybiBlbGVtZW50c1tpbnN0YW5jZUluZGV4XSB8fCBudWxsO1xuICB9XG5cbiAgI3Bvc3RLbm9iQ2hhbmdlKFxuICAgIGtub2JUeXBlOiBzdHJpbmcsXG4gICAgbmFtZTogc3RyaW5nLFxuICAgIHZhbHVlOiB1bmtub3duLFxuICAgIHRhZ05hbWU6IHN0cmluZyxcbiAgICBpbnN0YW5jZUluZGV4OiBudW1iZXIsXG4gICk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IG1zZyA9IHsgdHlwZTogJ2NlbS1rbm9iLWNoYW5nZScsIGtub2JUeXBlLCBuYW1lLCB2YWx1ZSwgdGFnTmFtZSwgaW5zdGFuY2VJbmRleCB9O1xuICAgIGlmICghdGhpcy4jaWZyYW1lUmVhZHkpIHtcbiAgICAgIHRoaXMuI3BlbmRpbmdNZXNzYWdlcy5wdXNoKG1zZyk7XG4gICAgICByZXR1cm4gdHJ1ZTtcbiAgICB9XG4gICAgY29uc3QgaWZyYW1lID0gdGhpcy4jaWZyYW1lO1xuICAgIGlmICghaWZyYW1lPy5jb250ZW50V2luZG93KSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtZGVtb10gSWZyYW1lIG5vdCByZWFkeSBmb3IgcG9zdE1lc3NhZ2UnKTtcbiAgICAgIHJldHVybiBmYWxzZTtcbiAgICB9XG4gICAgaWZyYW1lLmNvbnRlbnRXaW5kb3cucG9zdE1lc3NhZ2UobXNnLCB3aW5kb3cubG9jYXRpb24ub3JpZ2luKTtcbiAgICByZXR1cm4gdHJ1ZTtcbiAgfVxuXG4gICNhcHBseUF0dHJpYnV0ZUNoYW5nZShlbGVtZW50OiBFbGVtZW50LCBuYW1lOiBzdHJpbmcsIHZhbHVlOiB1bmtub3duKTogYm9vbGVhbiB7XG4gICAgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ2Jvb2xlYW4nKSB7XG4gICAgICBlbGVtZW50LnRvZ2dsZUF0dHJpYnV0ZShuYW1lLCB2YWx1ZSk7XG4gICAgfSBlbHNlIGlmICh2YWx1ZSA9PT0gJycgfHwgdmFsdWUgPT09IG51bGwgfHwgdmFsdWUgPT09IHVuZGVmaW5lZCkge1xuICAgICAgZWxlbWVudC5yZW1vdmVBdHRyaWJ1dGUobmFtZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIGVsZW1lbnQuc2V0QXR0cmlidXRlKG5hbWUsIFN0cmluZyh2YWx1ZSkpO1xuICAgIH1cbiAgICByZXR1cm4gdHJ1ZTtcbiAgfVxuXG4gICNhcHBseVByb3BlcnR5Q2hhbmdlPFQgZXh0ZW5kcyBFbGVtZW50PihcbiAgICBlbGVtZW50OiBULFxuICAgIG5hbWU6IGtleW9mIFQsXG4gICAgdmFsdWU6IFRba2V5b2YgVF0sXG4gICk6IGJvb2xlYW4ge1xuICAgIGlmICh2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICB0cnkge1xuICAgICAgICBkZWxldGUgKGVsZW1lbnQpW25hbWVdO1xuICAgICAgfSBjYXRjaCB7XG4gICAgICAgIChlbGVtZW50KVtuYW1lXSA9IHZhbHVlO1xuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICAoZWxlbWVudClbbmFtZV0gPSB2YWx1ZTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAjYXBwbHlDU1NQcm9wZXJ0eUNoYW5nZShlbGVtZW50OiBIVE1MRWxlbWVudCwgbmFtZTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bik6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHByb3BlcnR5TmFtZSA9IG5hbWUuc3RhcnRzV2l0aCgnLS0nKSA/IG5hbWUgOiBgLS0ke25hbWV9YDtcbiAgICBpZiAodmFsdWUgPT09ICcnIHx8IHZhbHVlID09PSBudWxsIHx8IHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgIGVsZW1lbnQuc3R5bGUucmVtb3ZlUHJvcGVydHkocHJvcGVydHlOYW1lKTtcbiAgICB9IGVsc2Uge1xuICAgICAgZWxlbWVudC5zdHlsZS5zZXRQcm9wZXJ0eShwcm9wZXJ0eU5hbWUsIFN0cmluZyh2YWx1ZSkpO1xuICAgIH1cbiAgICByZXR1cm4gdHJ1ZTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tc2VydmUtZGVtbyc6IENlbVNlcnZlRGVtbztcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbn1cXFxcblxcXFxuaWZyYW1lIHtcXFxcbiAgYm9yZGVyOiBub25lO1xcXFxuICB3aWR0aDogMTAwJTtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7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
