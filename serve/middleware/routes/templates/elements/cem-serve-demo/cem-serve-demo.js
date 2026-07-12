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
var _rendering_dec, _a, _CemServeDemo_decorators, _init, _rendering, _iframeReady, _pendingMessages, _CemServeDemo_instances, iframe_get, iframeSrc_fn, _onChildReady, getElementInstance_fn, postKnobChange_fn, applyAttributeChange_fn, applyPropertyChange_fn, applyCSSStateChange_fn, applyCSSPropertyChange_fn;
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
      case "css-state":
        return __privateMethod(this, _CemServeDemo_instances, applyCSSStateChange_fn).call(this, element, name, value);
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
applyCSSStateChange_fn = function(element, name, value) {
  const stateName = name.startsWith("--") ? name.slice(2) : name;
  const internals = globalThis._elementInternals?.get(element);
  const states = internals?.states;
  if (!states) return false;
  if (value) {
    states.add(stateName);
  } else {
    states.delete(stateName);
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8udHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tc2VydmUtZGVtby9jZW0tc2VydmUtZGVtby5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImRlY2xhcmUgZ2xvYmFsIHtcbiAgdmFyIF9lbGVtZW50SW50ZXJuYWxzOiBXZWFrTWFwPEVsZW1lbnQsIEVsZW1lbnRJbnRlcm5hbHM+IHwgdW5kZWZpbmVkO1xufVxuXG5pbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1zZXJ2ZS1kZW1vLmNzcyc7XG5cbi8qKlxuICogRGVtbyB3cmFwcGVyIGNvbXBvbmVudCBmb3Iga25vYnMgaW50ZWdyYXRpb24uXG4gKlxuICogSW4gbGlnaHQvc2hhZG93IG1vZGUsIHJlbmRlcnMgZGVtbyBjb250ZW50IHZpYSBhIGRlZmF1bHQgc2xvdC5cbiAqIEluIGlmcmFtZSBtb2RlLCBsb2FkcyB0aGUgZGVtbyBpbiBhbiBpc29sYXRlZCBpZnJhbWUgYW5kIGJyaWRnZXNcbiAqIGtub2IgY2hhbmdlcyB2aWEgcG9zdE1lc3NhZ2UuXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGRlbW8gY29udGVudCAobGlnaHQvc2hhZG93IG1vZGUgb25seSlcbiAqIEBjdXN0b21FbGVtZW50IGNlbS1zZXJ2ZS1kZW1vXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tc2VydmUtZGVtbycpXG5leHBvcnQgY2xhc3MgQ2VtU2VydmVEZW1vIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KSBhY2Nlc3NvciByZW5kZXJpbmc6IHN0cmluZyB8IHVuZGVmaW5lZDtcblxuICAjaWZyYW1lUmVhZHkgPSBmYWxzZTtcbiAgI3BlbmRpbmdNZXNzYWdlczogQXJyYXk8UmVjb3JkPHN0cmluZywgdW5rbm93bj4+ID0gW107XG5cbiAgZ2V0ICNpZnJhbWUoKTogSFRNTElGcmFtZUVsZW1lbnQgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5yZW5kZXJSb290LnF1ZXJ5U2VsZWN0b3IoJ2lmcmFtZScpO1xuICB9XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB3aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcignbWVzc2FnZScsIHRoaXMuI29uQ2hpbGRSZWFkeSk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHdpbmRvdy5yZW1vdmVFdmVudExpc3RlbmVyKCdtZXNzYWdlJywgdGhpcy4jb25DaGlsZFJlYWR5KTtcbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICBpZiAodGhpcy5yZW5kZXJpbmcgPT09ICdpZnJhbWUnKSB7XG4gICAgICB0aGlzLiNpZnJhbWVSZWFkeSA9IGZhbHNlO1xuICAgICAgcmV0dXJuIGh0bWxgPGlmcmFtZSBwYXJ0PVwiaWZyYW1lXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgc3JjPVwiJHt0aGlzLiNpZnJhbWVTcmMoKX1cIj48L2lmcmFtZT5gO1xuICAgIH1cbiAgICByZXR1cm4gaHRtbGA8c2xvdD48L3Nsb3Q+YDtcbiAgfVxuXG4gIC8qKlxuICAgKiBBcHBseSBhIGtub2IgY2hhbmdlIHRvIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW8uXG4gICAqIENhbGxlZCBieSBwYXJlbnQgY2hyb21lIGVsZW1lbnQgd2hlbiBrbm9iIGV2ZW50cyBvY2N1ci5cbiAgICogSW4gaWZyYW1lIG1vZGUsIGJyaWRnZXMgdmlhIHBvc3RNZXNzYWdlIGluc3RlYWQgb2YgZGlyZWN0IERPTSBhY2Nlc3MuXG4gICAqL1xuICBhcHBseUtub2JDaGFuZ2UoXG4gICAgdHlwZTogc3RyaW5nLFxuICAgIG5hbWU6IHN0cmluZyxcbiAgICB2YWx1ZTogdW5rbm93bixcbiAgICB0YWdOYW1lOiBzdHJpbmcsXG4gICAgaW5zdGFuY2VJbmRleCA9IDAsXG4gICk6IGJvb2xlYW4ge1xuICAgIGlmICh0aGlzLnJlbmRlcmluZyA9PT0gJ2lmcmFtZScpIHtcbiAgICAgIHJldHVybiB0aGlzLiNwb3N0S25vYkNoYW5nZSh0eXBlLCBuYW1lLCB2YWx1ZSwgdGFnTmFtZSwgaW5zdGFuY2VJbmRleCk7XG4gICAgfVxuXG4gICAgY29uc3QgZWxlbWVudCA9IHRoaXMuI2dldEVsZW1lbnRJbnN0YW5jZSh0YWdOYW1lLCBpbnN0YW5jZUluZGV4KTtcbiAgICBpZiAoIWVsZW1lbnQpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1kZW1vXSBFbGVtZW50IG5vdCBmb3VuZDonLCB0YWdOYW1lLCAnYXQgaW5kZXgnLCBpbnN0YW5jZUluZGV4KTtcbiAgICAgIHJldHVybiBmYWxzZTtcbiAgICB9XG5cbiAgICBzd2l0Y2ggKHR5cGUpIHtcbiAgICAgIGNhc2UgJ2F0dHJpYnV0ZSc6XG4gICAgICAgIHJldHVybiB0aGlzLiNhcHBseUF0dHJpYnV0ZUNoYW5nZShlbGVtZW50LCBuYW1lLCB2YWx1ZSk7XG4gICAgICBjYXNlICdwcm9wZXJ0eSc6XG4gICAgICAgIHJldHVybiB0aGlzLiNhcHBseVByb3BlcnR5Q2hhbmdlKFxuICAgICAgICAgIGVsZW1lbnQsXG4gICAgICAgICAgbmFtZSBhcyBrZXlvZiBFbGVtZW50LFxuICAgICAgICAgIHZhbHVlIGFzIEVsZW1lbnRba2V5b2YgRWxlbWVudF0sXG4gICAgICAgICk7XG4gICAgICBjYXNlICdjc3MtcHJvcGVydHknOlxuICAgICAgICByZXR1cm4gdGhpcy4jYXBwbHlDU1NQcm9wZXJ0eUNoYW5nZShlbGVtZW50IGFzIEhUTUxFbGVtZW50LCBuYW1lLCB2YWx1ZSk7XG4gICAgICBjYXNlICdjc3Mtc3RhdGUnOlxuICAgICAgICByZXR1cm4gdGhpcy4jYXBwbHlDU1NTdGF0ZUNoYW5nZShlbGVtZW50LCBuYW1lLCB2YWx1ZSk7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtZGVtb10gVW5rbm93biBrbm9iIHR5cGU6JywgdHlwZSk7XG4gICAgICAgIHJldHVybiBmYWxzZTtcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogU2V0IGFuIGF0dHJpYnV0ZSBvbiBhbiBlbGVtZW50IGluIHRoZSBkZW1vXG4gICAqL1xuICBzZXREZW1vQXR0cmlidXRlKHNlbGVjdG9yOiBzdHJpbmcsIGF0dHJpYnV0ZTogc3RyaW5nLCB2YWx1ZTogc3RyaW5nIHwgYm9vbGVhbiB8IG51bGwpOiBib29sZWFuIHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpO1xuICAgIGlmICghdGFyZ2V0KSByZXR1cm4gZmFsc2U7XG5cbiAgICBpZiAodHlwZW9mIHZhbHVlID09PSAnYm9vbGVhbicpIHtcbiAgICAgIHRhcmdldC50b2dnbGVBdHRyaWJ1dGUoYXR0cmlidXRlLCB2YWx1ZSk7XG4gICAgfSBlbHNlIGlmICh2YWx1ZSA9PT0gJycgfHwgdmFsdWUgPT09IG51bGwgfHwgdmFsdWUgPT09IHVuZGVmaW5lZCkge1xuICAgICAgdGFyZ2V0LnJlbW92ZUF0dHJpYnV0ZShhdHRyaWJ1dGUpO1xuICAgIH0gZWxzZSB7XG4gICAgICB0YXJnZXQuc2V0QXR0cmlidXRlKGF0dHJpYnV0ZSwgdmFsdWUpO1xuICAgIH1cblxuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgLyoqXG4gICAqIFNldCBhIHByb3BlcnR5IG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICovXG4gIHNldERlbW9Qcm9wZXJ0eTxUIGV4dGVuZHMgRWxlbWVudD4oXG4gICAgc2VsZWN0b3I6IHN0cmluZyxcbiAgICBwcm9wZXJ0eToga2V5b2YgVCxcbiAgICB2YWx1ZTogVFtrZXlvZiBUXSxcbiAgKTogYm9vbGVhbiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gdGhpcy5xdWVyeVNlbGVjdG9yPFQ+KHNlbGVjdG9yKTtcbiAgICBpZiAodGFyZ2V0KSB7XG4gICAgICAodGFyZ2V0KVtwcm9wZXJ0eV0gPSB2YWx1ZTtcbiAgICAgIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICByZXR1cm4gZmFsc2U7XG4gIH1cblxuICAvKipcbiAgICogU2V0IGEgQ1NTIGN1c3RvbSBwcm9wZXJ0eSBvbiBhbiBlbGVtZW50IGluIHRoZSBkZW1vXG4gICAqL1xuICBzZXREZW1vQ3NzQ3VzdG9tUHJvcGVydHkoc2VsZWN0b3I6IHN0cmluZywgY3NzUHJvcGVydHk6IHN0cmluZywgdmFsdWU6IHN0cmluZyk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMucXVlcnlTZWxlY3RvcihzZWxlY3RvcikgYXMgSFRNTEVsZW1lbnQgfCBudWxsO1xuICAgIGlmICh0YXJnZXQpIHtcbiAgICAgIGNvbnN0IHByb3BlcnR5TmFtZSA9IGNzc1Byb3BlcnR5LnN0YXJ0c1dpdGgoJy0tJykgPyBjc3NQcm9wZXJ0eSA6IGAtLSR7Y3NzUHJvcGVydHl9YDtcbiAgICAgIHRhcmdldC5zdHlsZS5zZXRQcm9wZXJ0eShwcm9wZXJ0eU5hbWUsIHZhbHVlKTtcbiAgICAgIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICByZXR1cm4gZmFsc2U7XG4gIH1cblxuICAjaWZyYW1lU3JjKCk6IHN0cmluZyB7XG4gICAgY29uc3QgdXJsID0gbmV3IFVSTCh3aW5kb3cubG9jYXRpb24uaHJlZik7XG4gICAgdXJsLnNlYXJjaFBhcmFtcy5zZXQoJ3JlbmRlcmluZycsICdjaHJvbWVsZXNzJyk7XG4gICAgcmV0dXJuIHVybC50b1N0cmluZygpO1xuICB9XG5cbiAgI29uQ2hpbGRSZWFkeSA9IChldmVudDogTWVzc2FnZUV2ZW50KSA9PiB7XG4gICAgaWYgKGV2ZW50Lm9yaWdpbiAhPT0gd2luZG93LmxvY2F0aW9uLm9yaWdpbikgcmV0dXJuO1xuICAgIGlmIChldmVudC5kYXRhPy50eXBlICE9PSAnY2VtLWlmcmFtZS1yZWFkeScpIHJldHVybjtcbiAgICB0aGlzLiNpZnJhbWVSZWFkeSA9IHRydWU7XG4gICAgY29uc3QgY3cgPSB0aGlzLiNpZnJhbWU/LmNvbnRlbnRXaW5kb3c7XG4gICAgaWYgKGN3KSB7XG4gICAgICBmb3IgKGNvbnN0IG1zZyBvZiB0aGlzLiNwZW5kaW5nTWVzc2FnZXMpXG4gICAgICAgIGN3LnBvc3RNZXNzYWdlKG1zZywgd2luZG93LmxvY2F0aW9uLm9yaWdpbik7XG4gICAgfVxuICAgIHRoaXMuI3BlbmRpbmdNZXNzYWdlcyA9IFtdO1xuICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQoJ2lmcmFtZS1yZWFkeScpKTtcbiAgfTtcblxuICAvKipcbiAgICogRmluZCB0aGUgTnRoIGluc3RhbmNlIG9mIGFuIGVsZW1lbnQgYnkgdGFnIG5hbWVcbiAgICovXG4gICNnZXRFbGVtZW50SW5zdGFuY2UodGFnTmFtZTogc3RyaW5nLCBpbnN0YW5jZUluZGV4ID0gMCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICBjb25zdCBlbGVtZW50cyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCh0YWdOYW1lKTtcbiAgICByZXR1cm4gZWxlbWVudHNbaW5zdGFuY2VJbmRleF0gfHwgbnVsbDtcbiAgfVxuXG4gICNwb3N0S25vYkNoYW5nZShcbiAgICBrbm9iVHlwZTogc3RyaW5nLFxuICAgIG5hbWU6IHN0cmluZyxcbiAgICB2YWx1ZTogdW5rbm93bixcbiAgICB0YWdOYW1lOiBzdHJpbmcsXG4gICAgaW5zdGFuY2VJbmRleDogbnVtYmVyLFxuICApOiBib29sZWFuIHtcbiAgICBjb25zdCBtc2cgPSB7IHR5cGU6ICdjZW0ta25vYi1jaGFuZ2UnLCBrbm9iVHlwZSwgbmFtZSwgdmFsdWUsIHRhZ05hbWUsIGluc3RhbmNlSW5kZXggfTtcbiAgICBpZiAoIXRoaXMuI2lmcmFtZVJlYWR5KSB7XG4gICAgICB0aGlzLiNwZW5kaW5nTWVzc2FnZXMucHVzaChtc2cpO1xuICAgICAgcmV0dXJuIHRydWU7XG4gICAgfVxuICAgIGNvbnN0IGlmcmFtZSA9IHRoaXMuI2lmcmFtZTtcbiAgICBpZiAoIWlmcmFtZT8uY29udGVudFdpbmRvdykge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWRlbW9dIElmcmFtZSBub3QgcmVhZHkgZm9yIHBvc3RNZXNzYWdlJyk7XG4gICAgICByZXR1cm4gZmFsc2U7XG4gICAgfVxuICAgIGlmcmFtZS5jb250ZW50V2luZG93LnBvc3RNZXNzYWdlKG1zZywgd2luZG93LmxvY2F0aW9uLm9yaWdpbik7XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAjYXBwbHlBdHRyaWJ1dGVDaGFuZ2UoZWxlbWVudDogRWxlbWVudCwgbmFtZTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bik6IGJvb2xlYW4ge1xuICAgIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgZWxlbWVudC50b2dnbGVBdHRyaWJ1dGUobmFtZSwgdmFsdWUpO1xuICAgIH0gZWxzZSBpZiAodmFsdWUgPT09ICcnIHx8IHZhbHVlID09PSBudWxsIHx8IHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgIGVsZW1lbnQucmVtb3ZlQXR0cmlidXRlKG5hbWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICBlbGVtZW50LnNldEF0dHJpYnV0ZShuYW1lLCBTdHJpbmcodmFsdWUpKTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAjYXBwbHlQcm9wZXJ0eUNoYW5nZTxUIGV4dGVuZHMgRWxlbWVudD4oXG4gICAgZWxlbWVudDogVCxcbiAgICBuYW1lOiBrZXlvZiBULFxuICAgIHZhbHVlOiBUW2tleW9mIFRdLFxuICApOiBib29sZWFuIHtcbiAgICBpZiAodmFsdWUgPT09IHVuZGVmaW5lZCkge1xuICAgICAgdHJ5IHtcbiAgICAgICAgZGVsZXRlIChlbGVtZW50KVtuYW1lXTtcbiAgICAgIH0gY2F0Y2gge1xuICAgICAgICAoZWxlbWVudClbbmFtZV0gPSB2YWx1ZTtcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgKGVsZW1lbnQpW25hbWVdID0gdmFsdWU7XG4gICAgfVxuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgI2FwcGx5Q1NTU3RhdGVDaGFuZ2UoZWxlbWVudDogRWxlbWVudCwgbmFtZTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bik6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHN0YXRlTmFtZSA9IG5hbWUuc3RhcnRzV2l0aCgnLS0nKSA/IG5hbWUuc2xpY2UoMikgOiBuYW1lO1xuICAgIGNvbnN0IGludGVybmFscyA9IGdsb2JhbFRoaXMuX2VsZW1lbnRJbnRlcm5hbHM/LmdldChlbGVtZW50KTtcbiAgICBjb25zdCBzdGF0ZXMgPSBpbnRlcm5hbHM/LnN0YXRlcztcbiAgICBpZiAoIXN0YXRlcykgcmV0dXJuIGZhbHNlO1xuICAgIGlmICh2YWx1ZSkge1xuICAgICAgc3RhdGVzLmFkZChzdGF0ZU5hbWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICBzdGF0ZXMuZGVsZXRlKHN0YXRlTmFtZSk7XG4gICAgfVxuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgI2FwcGx5Q1NTUHJvcGVydHlDaGFuZ2UoZWxlbWVudDogSFRNTEVsZW1lbnQsIG5hbWU6IHN0cmluZywgdmFsdWU6IHVua25vd24pOiBib29sZWFuIHtcbiAgICBjb25zdCBwcm9wZXJ0eU5hbWUgPSBuYW1lLnN0YXJ0c1dpdGgoJy0tJykgPyBuYW1lIDogYC0tJHtuYW1lfWA7XG4gICAgaWYgKHZhbHVlID09PSAnJyB8fCB2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICBlbGVtZW50LnN0eWxlLnJlbW92ZVByb3BlcnR5KHByb3BlcnR5TmFtZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIGVsZW1lbnQuc3R5bGUuc2V0UHJvcGVydHkocHJvcGVydHlOYW1lLCBTdHJpbmcodmFsdWUpKTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXNlcnZlLWRlbW8nOiBDZW1TZXJ2ZURlbW87XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG59XFxcXG5cXFxcbmlmcmFtZSB7XFxcXG4gIGJvcmRlcjogbm9uZTtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBSUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ056QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSw2R0FBK0csQ0FBQztBQUFFLElBQU8seUJBQVE7OztBREF0TDtBQW9CQSw0QkFBQyxjQUFjLGdCQUFnQjtBQUN4QixJQUFNLGVBQU4sZUFBMkIsaUJBR2hDLGtCQUFDLFNBQVMsRUFBRSxTQUFTLEtBQUssQ0FBQyxJQUhLLElBQVc7QUFBQSxFQUF0QztBQUFBO0FBQUE7QUFHd0IsdUJBQVMsWUFBVDtBQUU3QixxQ0FBZTtBQUNmLHlDQUFtRCxDQUFDO0FBdUhwRCxzQ0FBZ0IsQ0FBQyxVQUF3QjtBQUN2QyxVQUFJLE1BQU0sV0FBVyxPQUFPLFNBQVMsT0FBUTtBQUM3QyxVQUFJLE1B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
