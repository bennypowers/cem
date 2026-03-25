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
var __accessCheck = (obj, member, msg) => member.has(obj) || __typeError("Cannot " + msg);
var __privateIn = (member, obj) => Object(obj) !== obj ? __typeError('Cannot use the "in" operator on this value') : member.has(obj);
var __privateGet = (obj, member, getter) => (__accessCheck(obj, member, "read from private field"), getter ? getter.call(obj) : member.get(obj));
var __privateSet = (obj, member, value, setter) => (__accessCheck(obj, member, "write to private field"), setter ? setter.call(obj, value) : member.set(obj, value), value);
var __privateMethod = (obj, member, method) => (__accessCheck(obj, member, "access private method"), method);

// elements/cem-serve-demo/cem-serve-demo.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-serve-demo/cem-serve-demo.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n"'));
var cem_serve_demo_default = s;

// elements/cem-serve-demo/cem-serve-demo.ts
var _CemServeDemo_decorators, _init, _a;
_CemServeDemo_decorators = [customElement("cem-serve-demo")];
var CemServeDemo = class extends (_a = LitElement) {
  static styles = cem_serve_demo_default;
  render() {
    return html`<slot></slot>`;
  }
  /**
   * Find the Nth instance of an element by tag name
   */
  #getElementInstance(tagName, instanceIndex = 0) {
    const elements = this.querySelectorAll(tagName);
    return elements[instanceIndex] || null;
  }
  /**
   * Apply a knob change to an element in the demo.
   * Called by parent chrome element when knob events occur.
   * @param type - 'attribute', 'property', or 'css-property'
   * @param name - Attribute/property/CSS name
   * @param value - Value to apply
   * @param tagName - Target element tag name
   * @param instanceIndex - Which instance of the element (0-based)
   * @returns Whether the operation succeeded
   */
  applyKnobChange(type, name, value, tagName, instanceIndex = 0) {
    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn("[cem-serve-demo] Element not found:", tagName, "at index", instanceIndex);
      return false;
    }
    switch (type) {
      case "attribute":
        return this.#applyAttributeChange(element, name, value);
      case "property":
        return this.#applyPropertyChange(element, name, value);
      case "css-property":
        return this.#applyCSSPropertyChange(element, name, value);
      default:
        console.warn("[cem-serve-demo] Unknown knob type:", type);
        return false;
    }
  }
  #applyAttributeChange(element, name, value) {
    if (typeof value === "boolean") {
      element.toggleAttribute(name, value);
    } else if (value === "" || value === null || value === void 0) {
      element.removeAttribute(name);
    } else {
      element.setAttribute(name, String(value));
    }
    return true;
  }
  #applyPropertyChange(element, name, value) {
    if (value === void 0) {
      try {
        delete element[name];
      } catch {
        element[name] = void 0;
      }
    } else {
      element[name] = value;
    }
    return true;
  }
  #applyCSSPropertyChange(element, name, value) {
    const propertyName = name.startsWith("--") ? name : `--${name}`;
    if (value === "" || value === null || value === void 0) {
      element.style.removeProperty(propertyName);
    } else {
      element.style.setProperty(propertyName, String(value));
    }
    return true;
  }
  /**
   * Set an attribute on an element in the demo
   * @param selector - CSS selector for target element
   * @param attribute - Attribute name
   * @param value - Attribute value (boolean for presence/absence)
   * @returns Whether the operation succeeded
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
   * @param selector - CSS selector for target element
   * @param property - Property name
   * @param value - Property value
   * @returns Whether the operation succeeded
   */
  setDemoProperty(selector, property, value) {
    const target = this.querySelector(selector);
    if (target) {
      target[property] = value;
      return true;
    }
    return false;
  }
  /**
   * Set a CSS custom property on an element in the demo
   * @param selector - CSS selector for target element
   * @param cssProperty - CSS custom property name (with or without --)
   * @param value - CSS property value
   * @returns Whether the operation succeeded
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
CemServeDemo = __decorateElement(_init, 0, "CemServeDemo", _CemServeDemo_decorators, CemServeDemo);
__runInitializers(_init, 1, CemServeDemo);
export {
  CemServeDemo
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8udHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9jZW0tc2VydmUtZGVtby9jZW0tc2VydmUtZGVtby5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1zZXJ2ZS1kZW1vLmNzcyc7XG5cbi8qKlxuICogRGVtbyB3cmFwcGVyIGNvbXBvbmVudCBmb3Iga25vYnMgaW50ZWdyYXRpb24uXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGRlbW8gY29udGVudFxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXNlcnZlLWRlbW9cbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1zZXJ2ZS1kZW1vJylcbmV4cG9ydCBjbGFzcyBDZW1TZXJ2ZURlbW8gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgPHNsb3Q+PC9zbG90PmA7XG4gIH1cblxuICAvKipcbiAgICogRmluZCB0aGUgTnRoIGluc3RhbmNlIG9mIGFuIGVsZW1lbnQgYnkgdGFnIG5hbWVcbiAgICovXG4gICNnZXRFbGVtZW50SW5zdGFuY2UodGFnTmFtZTogc3RyaW5nLCBpbnN0YW5jZUluZGV4ID0gMCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICBjb25zdCBlbGVtZW50cyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCh0YWdOYW1lKTtcbiAgICByZXR1cm4gZWxlbWVudHNbaW5zdGFuY2VJbmRleF0gfHwgbnVsbDtcbiAgfVxuXG4gIC8qKlxuICAgKiBBcHBseSBhIGtub2IgY2hhbmdlIHRvIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW8uXG4gICAqIENhbGxlZCBieSBwYXJlbnQgY2hyb21lIGVsZW1lbnQgd2hlbiBrbm9iIGV2ZW50cyBvY2N1ci5cbiAgICogQHBhcmFtIHR5cGUgLSAnYXR0cmlidXRlJywgJ3Byb3BlcnR5Jywgb3IgJ2Nzcy1wcm9wZXJ0eSdcbiAgICogQHBhcmFtIG5hbWUgLSBBdHRyaWJ1dGUvcHJvcGVydHkvQ1NTIG5hbWVcbiAgICogQHBhcmFtIHZhbHVlIC0gVmFsdWUgdG8gYXBwbHlcbiAgICogQHBhcmFtIHRhZ05hbWUgLSBUYXJnZXQgZWxlbWVudCB0YWcgbmFtZVxuICAgKiBAcGFyYW0gaW5zdGFuY2VJbmRleCAtIFdoaWNoIGluc3RhbmNlIG9mIHRoZSBlbGVtZW50ICgwLWJhc2VkKVxuICAgKiBAcmV0dXJucyBXaGV0aGVyIHRoZSBvcGVyYXRpb24gc3VjY2VlZGVkXG4gICAqL1xuICBhcHBseUtub2JDaGFuZ2UoXG4gICAgdHlwZTogc3RyaW5nLFxuICAgIG5hbWU6IHN0cmluZyxcbiAgICB2YWx1ZTogdW5rbm93bixcbiAgICB0YWdOYW1lOiBzdHJpbmcsXG4gICAgaW5zdGFuY2VJbmRleCA9IDAsXG4gICk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IGVsZW1lbnQgPSB0aGlzLiNnZXRFbGVtZW50SW5zdGFuY2UodGFnTmFtZSwgaW5zdGFuY2VJbmRleCk7XG4gICAgaWYgKCFlbGVtZW50KSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtZGVtb10gRWxlbWVudCBub3QgZm91bmQ6JywgdGFnTmFtZSwgJ2F0IGluZGV4JywgaW5zdGFuY2VJbmRleCk7XG4gICAgICByZXR1cm4gZmFsc2U7XG4gICAgfVxuXG4gICAgc3dpdGNoICh0eXBlKSB7XG4gICAgICBjYXNlICdhdHRyaWJ1dGUnOlxuICAgICAgICByZXR1cm4gdGhpcy4jYXBwbHlBdHRyaWJ1dGVDaGFuZ2UoZWxlbWVudCwgbmFtZSwgdmFsdWUpO1xuICAgICAgY2FzZSAncHJvcGVydHknOlxuICAgICAgICByZXR1cm4gdGhpcy4jYXBwbHlQcm9wZXJ0eUNoYW5nZShlbGVtZW50LCBuYW1lLCB2YWx1ZSk7XG4gICAgICBjYXNlICdjc3MtcHJvcGVydHknOlxuICAgICAgICByZXR1cm4gdGhpcy4jYXBwbHlDU1NQcm9wZXJ0eUNoYW5nZShlbGVtZW50IGFzIEhUTUxFbGVtZW50LCBuYW1lLCB2YWx1ZSk7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtZGVtb10gVW5rbm93biBrbm9iIHR5cGU6JywgdHlwZSk7XG4gICAgICAgIHJldHVybiBmYWxzZTtcbiAgICB9XG4gIH1cblxuICAjYXBwbHlBdHRyaWJ1dGVDaGFuZ2UoZWxlbWVudDogRWxlbWVudCwgbmFtZTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bik6IGJvb2xlYW4ge1xuICAgIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgZWxlbWVudC50b2dnbGVBdHRyaWJ1dGUobmFtZSwgdmFsdWUpO1xuICAgIH0gZWxzZSBpZiAodmFsdWUgPT09ICcnIHx8IHZhbHVlID09PSBudWxsIHx8IHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgIGVsZW1lbnQucmVtb3ZlQXR0cmlidXRlKG5hbWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICBlbGVtZW50LnNldEF0dHJpYnV0ZShuYW1lLCBTdHJpbmcodmFsdWUpKTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAjYXBwbHlQcm9wZXJ0eUNoYW5nZShlbGVtZW50OiBFbGVtZW50LCBuYW1lOiBzdHJpbmcsIHZhbHVlOiB1bmtub3duKTogYm9vbGVhbiB7XG4gICAgaWYgKHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgIHRyeSB7XG4gICAgICAgIGRlbGV0ZSAoZWxlbWVudCBhcyBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPilbbmFtZV07XG4gICAgICB9IGNhdGNoIHtcbiAgICAgICAgKGVsZW1lbnQgYXMgUmVjb3JkPHN0cmluZywgdW5rbm93bj4pW25hbWVdID0gdW5kZWZpbmVkO1xuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICAoZWxlbWVudCBhcyBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPilbbmFtZV0gPSB2YWx1ZTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAjYXBwbHlDU1NQcm9wZXJ0eUNoYW5nZShlbGVtZW50OiBIVE1MRWxlbWVudCwgbmFtZTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bik6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHByb3BlcnR5TmFtZSA9IG5hbWUuc3RhcnRzV2l0aCgnLS0nKSA/IG5hbWUgOiBgLS0ke25hbWV9YDtcbiAgICBpZiAodmFsdWUgPT09ICcnIHx8IHZhbHVlID09PSBudWxsIHx8IHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgIGVsZW1lbnQuc3R5bGUucmVtb3ZlUHJvcGVydHkocHJvcGVydHlOYW1lKTtcbiAgICB9IGVsc2Uge1xuICAgICAgZWxlbWVudC5zdHlsZS5zZXRQcm9wZXJ0eShwcm9wZXJ0eU5hbWUsIFN0cmluZyh2YWx1ZSkpO1xuICAgIH1cbiAgICByZXR1cm4gdHJ1ZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTZXQgYW4gYXR0cmlidXRlIG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICogQHBhcmFtIHNlbGVjdG9yIC0gQ1NTIHNlbGVjdG9yIGZvciB0YXJnZXQgZWxlbWVudFxuICAgKiBAcGFyYW0gYXR0cmlidXRlIC0gQXR0cmlidXRlIG5hbWVcbiAgICogQHBhcmFtIHZhbHVlIC0gQXR0cmlidXRlIHZhbHVlIChib29sZWFuIGZvciBwcmVzZW5jZS9hYnNlbmNlKVxuICAgKiBAcmV0dXJucyBXaGV0aGVyIHRoZSBvcGVyYXRpb24gc3VjY2VlZGVkXG4gICAqL1xuICBzZXREZW1vQXR0cmlidXRlKHNlbGVjdG9yOiBzdHJpbmcsIGF0dHJpYnV0ZTogc3RyaW5nLCB2YWx1ZTogc3RyaW5nIHwgYm9vbGVhbiB8IG51bGwpOiBib29sZWFuIHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpO1xuICAgIGlmICghdGFyZ2V0KSByZXR1cm4gZmFsc2U7XG5cbiAgICBpZiAodHlwZW9mIHZhbHVlID09PSAnYm9vbGVhbicpIHtcbiAgICAgIHRhcmdldC50b2dnbGVBdHRyaWJ1dGUoYXR0cmlidXRlLCB2YWx1ZSk7XG4gICAgfSBlbHNlIGlmICh2YWx1ZSA9PT0gJycgfHwgdmFsdWUgPT09IG51bGwgfHwgdmFsdWUgPT09IHVuZGVmaW5lZCkge1xuICAgICAgdGFyZ2V0LnJlbW92ZUF0dHJpYnV0ZShhdHRyaWJ1dGUpO1xuICAgIH0gZWxzZSB7XG4gICAgICB0YXJnZXQuc2V0QXR0cmlidXRlKGF0dHJpYnV0ZSwgdmFsdWUpO1xuICAgIH1cblxuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgLyoqXG4gICAqIFNldCBhIHByb3BlcnR5IG9uIGFuIGVsZW1lbnQgaW4gdGhlIGRlbW9cbiAgICogQHBhcmFtIHNlbGVjdG9yIC0gQ1NTIHNlbGVjdG9yIGZvciB0YXJnZXQgZWxlbWVudFxuICAgKiBAcGFyYW0gcHJvcGVydHkgLSBQcm9wZXJ0eSBuYW1lXG4gICAqIEBwYXJhbSB2YWx1ZSAtIFByb3BlcnR5IHZhbHVlXG4gICAqIEByZXR1cm5zIFdoZXRoZXIgdGhlIG9wZXJhdGlvbiBzdWNjZWVkZWRcbiAgICovXG4gIHNldERlbW9Qcm9wZXJ0eShzZWxlY3Rvcjogc3RyaW5nLCBwcm9wZXJ0eTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bik6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMucXVlcnlTZWxlY3RvcihzZWxlY3Rvcik7XG4gICAgaWYgKHRhcmdldCkge1xuICAgICAgKHRhcmdldCBhcyBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPilbcHJvcGVydHldID0gdmFsdWU7XG4gICAgICByZXR1cm4gdHJ1ZTtcbiAgICB9XG4gICAgcmV0dXJuIGZhbHNlO1xuICB9XG5cbiAgLyoqXG4gICAqIFNldCBhIENTUyBjdXN0b20gcHJvcGVydHkgb24gYW4gZWxlbWVudCBpbiB0aGUgZGVtb1xuICAgKiBAcGFyYW0gc2VsZWN0b3IgLSBDU1Mgc2VsZWN0b3IgZm9yIHRhcmdldCBlbGVtZW50XG4gICAqIEBwYXJhbSBjc3NQcm9wZXJ0eSAtIENTUyBjdXN0b20gcHJvcGVydHkgbmFtZSAod2l0aCBvciB3aXRob3V0IC0tKVxuICAgKiBAcGFyYW0gdmFsdWUgLSBDU1MgcHJvcGVydHkgdmFsdWVcbiAgICogQHJldHVybnMgV2hldGhlciB0aGUgb3BlcmF0aW9uIHN1Y2NlZWRlZFxuICAgKi9cbiAgc2V0RGVtb0Nzc0N1c3RvbVByb3BlcnR5KHNlbGVjdG9yOiBzdHJpbmcsIGNzc1Byb3BlcnR5OiBzdHJpbmcsIHZhbHVlOiBzdHJpbmcpOiBib29sZWFuIHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpIGFzIEhUTUxFbGVtZW50IHwgbnVsbDtcbiAgICBpZiAodGFyZ2V0KSB7XG4gICAgICBjb25zdCBwcm9wZXJ0eU5hbWUgPSBjc3NQcm9wZXJ0eS5zdGFydHNXaXRoKCctLScpID8gY3NzUHJvcGVydHkgOiBgLS0ke2Nzc1Byb3BlcnR5fWA7XG4gICAgICB0YXJnZXQuc3R5bGUuc2V0UHJvcGVydHkocHJvcGVydHlOYW1lLCB2YWx1ZSk7XG4gICAgICByZXR1cm4gdHJ1ZTtcbiAgICB9XG4gICAgcmV0dXJuIGZhbHNlO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1zZXJ2ZS1kZW1vJzogQ2VtU2VydmVEZW1vO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4g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
