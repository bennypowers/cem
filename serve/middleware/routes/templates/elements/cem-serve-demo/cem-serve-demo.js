var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __knownSymbol = (name, symbol) => (symbol = Symbol[name]) ? symbol : Symbol.for("Symbol." + name);
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
    const root = this.shadowRoot ?? this;
    const elements = root.querySelectorAll(tagName);
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
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
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
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
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
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8udHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9jZW0tc2VydmUtZGVtby9jZW0tc2VydmUtZGVtby5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1zZXJ2ZS1kZW1vLmNzcyc7XG5cbi8qKlxuICogRGVtbyB3cmFwcGVyIGNvbXBvbmVudCBmb3Iga25vYnMgaW50ZWdyYXRpb24uXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGRlbW8gY29udGVudFxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXNlcnZlLWRlbW9cbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1zZXJ2ZS1kZW1vJylcbmV4cG9ydCBjbGFzcyBDZW1TZXJ2ZURlbW8gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgPHNsb3Q+PC9zbG90PmA7XG4gIH1cblxuICAvKipcbiAgICogRmluZCB0aGUgTnRoIGluc3RhbmNlIG9mIGFuIGVsZW1lbnQgYnkgdGFnIG5hbWVcbiAgICovXG4gICNnZXRFbGVtZW50SW5zdGFuY2UodGFnTmFtZTogc3RyaW5nLCBpbnN0YW5jZUluZGV4ID0gMCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICBjb25zdCByb290ID0gdGhpcy5zaGFkb3dSb290ID8/IHRoaXM7XG4gICAgY29uc3QgZWxlbWVudHMgPSByb290LnF1ZXJ5U2VsZWN0b3JBbGwodGFnTmFtZSk7XG4gICAgcmV0dXJuIGVsZW1lbnRzW2luc3RhbmNlSW5kZXhdIHx8IG51bGw7XG4gIH1cblxuICAvKipcbiAgICogQXBwbHkgYSBrbm9iIGNoYW5nZSB0byBhbiBlbGVtZW50IGluIHRoZSBkZW1vLlxuICAgKiBDYWxsZWQgYnkgcGFyZW50IGNocm9tZSBlbGVtZW50IHdoZW4ga25vYiBldmVudHMgb2NjdXIuXG4gICAqIEBwYXJhbSB0eXBlIC0gJ2F0dHJpYnV0ZScsICdwcm9wZXJ0eScsIG9yICdjc3MtcHJvcGVydHknXG4gICAqIEBwYXJhbSBuYW1lIC0gQXR0cmlidXRlL3Byb3BlcnR5L0NTUyBuYW1lXG4gICAqIEBwYXJhbSB2YWx1ZSAtIFZhbHVlIHRvIGFwcGx5XG4gICAqIEBwYXJhbSB0YWdOYW1lIC0gVGFyZ2V0IGVsZW1lbnQgdGFnIG5hbWVcbiAgICogQHBhcmFtIGluc3RhbmNlSW5kZXggLSBXaGljaCBpbnN0YW5jZSBvZiB0aGUgZWxlbWVudCAoMC1iYXNlZClcbiAgICogQHJldHVybnMgV2hldGhlciB0aGUgb3BlcmF0aW9uIHN1Y2NlZWRlZFxuICAgKi9cbiAgYXBwbHlLbm9iQ2hhbmdlKFxuICAgIHR5cGU6IHN0cmluZyxcbiAgICBuYW1lOiBzdHJpbmcsXG4gICAgdmFsdWU6IHVua25vd24sXG4gICAgdGFnTmFtZTogc3RyaW5nLFxuICAgIGluc3RhbmNlSW5kZXggPSAwLFxuICApOiBib29sZWFuIHtcbiAgICBjb25zdCBlbGVtZW50ID0gdGhpcy4jZ2V0RWxlbWVudEluc3RhbmNlKHRhZ05hbWUsIGluc3RhbmNlSW5kZXgpO1xuICAgIGlmICghZWxlbWVudCkge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWRlbW9dIEVsZW1lbnQgbm90IGZvdW5kOicsIHRhZ05hbWUsICdhdCBpbmRleCcsIGluc3RhbmNlSW5kZXgpO1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cblxuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnYXR0cmlidXRlJzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5QXR0cmlidXRlQ2hhbmdlKGVsZW1lbnQsIG5hbWUsIHZhbHVlKTtcbiAgICAgIGNhc2UgJ3Byb3BlcnR5JzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5UHJvcGVydHlDaGFuZ2UoZWxlbWVudCwgbmFtZSwgdmFsdWUpO1xuICAgICAgY2FzZSAnY3NzLXByb3BlcnR5JzpcbiAgICAgICAgcmV0dXJuIHRoaXMuI2FwcGx5Q1NTUHJvcGVydHlDaGFuZ2UoZWxlbWVudCBhcyBIVE1MRWxlbWVudCwgbmFtZSwgdmFsdWUpO1xuICAgICAgZGVmYXVsdDpcbiAgICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWRlbW9dIFVua25vd24ga25vYiB0eXBlOicsIHR5cGUpO1xuICAgICAgICByZXR1cm4gZmFsc2U7XG4gICAgfVxuICB9XG5cbiAgI2FwcGx5QXR0cmlidXRlQ2hhbmdlKGVsZW1lbnQ6IEVsZW1lbnQsIG5hbWU6IHN0cmluZywgdmFsdWU6IHVua25vd24pOiBib29sZWFuIHtcbiAgICBpZiAodHlwZW9mIHZhbHVlID09PSAnYm9vbGVhbicpIHtcbiAgICAgIGVsZW1lbnQudG9nZ2xlQXR0cmlidXRlKG5hbWUsIHZhbHVlKTtcbiAgICB9IGVsc2UgaWYgKHZhbHVlID09PSAnJyB8fCB2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICBlbGVtZW50LnJlbW92ZUF0dHJpYnV0ZShuYW1lKTtcbiAgICB9IGVsc2Uge1xuICAgICAgZWxlbWVudC5zZXRBdHRyaWJ1dGUobmFtZSwgU3RyaW5nKHZhbHVlKSk7XG4gICAgfVxuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgI2FwcGx5UHJvcGVydHlDaGFuZ2UoZWxlbWVudDogRWxlbWVudCwgbmFtZTogc3RyaW5nLCB2YWx1ZTogdW5rbm93bik6IGJvb2xlYW4ge1xuICAgIGlmICh2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICB0cnkge1xuICAgICAgICBkZWxldGUgKGVsZW1lbnQgYXMgUmVjb3JkPHN0cmluZywgdW5rbm93bj4pW25hbWVdO1xuICAgICAgfSBjYXRjaCB7XG4gICAgICAgIChlbGVtZW50IGFzIFJlY29yZDxzdHJpbmcsIHVua25vd24+KVtuYW1lXSA9IHVuZGVmaW5lZDtcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgKGVsZW1lbnQgYXMgUmVjb3JkPHN0cmluZywgdW5rbm93bj4pW25hbWVdID0gdmFsdWU7XG4gICAgfVxuICAgIHJldHVybiB0cnVlO1xuICB9XG5cbiAgI2FwcGx5Q1NTUHJvcGVydHlDaGFuZ2UoZWxlbWVudDogSFRNTEVsZW1lbnQsIG5hbWU6IHN0cmluZywgdmFsdWU6IHVua25vd24pOiBib29sZWFuIHtcbiAgICBjb25zdCBwcm9wZXJ0eU5hbWUgPSBuYW1lLnN0YXJ0c1dpdGgoJy0tJykgPyBuYW1lIDogYC0tJHtuYW1lfWA7XG4gICAgaWYgKHZhbHVlID09PSAnJyB8fCB2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICBlbGVtZW50LnN0eWxlLnJlbW92ZVByb3BlcnR5KHByb3BlcnR5TmFtZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIGVsZW1lbnQuc3R5bGUuc2V0UHJvcGVydHkocHJvcGVydHlOYW1lLCBTdHJpbmcodmFsdWUpKTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH1cblxuICAvKipcbiAgICogU2V0IGFuIGF0dHJpYnV0ZSBvbiBhbiBlbGVtZW50IGluIHRoZSBkZW1vXG4gICAqIEBwYXJhbSBzZWxlY3RvciAtIENTUyBzZWxlY3RvciBmb3IgdGFyZ2V0IGVsZW1lbnRcbiAgICogQHBhcmFtIGF0dHJpYnV0ZSAtIEF0dHJpYnV0ZSBuYW1lXG4gICAqIEBwYXJhbSB2YWx1ZSAtIEF0dHJpYnV0ZSB2YWx1ZSAoYm9vbGVhbiBmb3IgcHJlc2VuY2UvYWJzZW5jZSlcbiAgICogQHJldHVybnMgV2hldGhlciB0aGUgb3BlcmF0aW9uIHN1Y2NlZWRlZFxuICAgKi9cbiAgc2V0RGVtb0F0dHJpYnV0ZShzZWxlY3Rvcjogc3RyaW5nLCBhdHRyaWJ1dGU6IHN0cmluZywgdmFsdWU6IHN0cmluZyB8IGJvb2xlYW4pOiBib29sZWFuIHtcbiAgICBjb25zdCByb290ID0gdGhpcy5zaGFkb3dSb290ID8/IHRoaXM7XG4gICAgY29uc3QgdGFyZ2V0ID0gcm9vdC5xdWVyeVNlbGVjdG9yKHNlbGVjdG9yKTtcbiAgICBpZiAoIXRhcmdldCkgcmV0dXJuIGZhbHNlO1xuXG4gICAgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ2Jvb2xlYW4nKSB7XG4gICAgICB0YXJnZXQudG9nZ2xlQXR0cmlidXRlKGF0dHJpYnV0ZSwgdmFsdWUpO1xuICAgIH0gZWxzZSBpZiAodmFsdWUgPT09ICcnIHx8IHZhbHVlID09PSBudWxsIHx8IHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgIHRhcmdldC5yZW1vdmVBdHRyaWJ1dGUoYXR0cmlidXRlKTtcbiAgICB9IGVsc2Uge1xuICAgICAgdGFyZ2V0LnNldEF0dHJpYnV0ZShhdHRyaWJ1dGUsIHZhbHVlKTtcbiAgICB9XG5cbiAgICByZXR1cm4gdHJ1ZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBTZXQgYSBwcm9wZXJ0eSBvbiBhbiBlbGVtZW50IGluIHRoZSBkZW1vXG4gICAqIEBwYXJhbSBzZWxlY3RvciAtIENTUyBzZWxlY3RvciBmb3IgdGFyZ2V0IGVsZW1lbnRcbiAgICogQHBhcmFtIHByb3BlcnR5IC0gUHJvcGVydHkgbmFtZVxuICAgKiBAcGFyYW0gdmFsdWUgLSBQcm9wZXJ0eSB2YWx1ZVxuICAgKiBAcmV0dXJucyBXaGV0aGVyIHRoZSBvcGVyYXRpb24gc3VjY2VlZGVkXG4gICAqL1xuICBzZXREZW1vUHJvcGVydHkoc2VsZWN0b3I6IHN0cmluZywgcHJvcGVydHk6IHN0cmluZywgdmFsdWU6IHVua25vd24pOiBib29sZWFuIHtcbiAgICBjb25zdCByb290ID0gdGhpcy5zaGFkb3dSb290ID8/IHRoaXM7XG4gICAgY29uc3QgdGFyZ2V0ID0gcm9vdC5xdWVyeVNlbGVjdG9yKHNlbGVjdG9yKTtcbiAgICBpZiAodGFyZ2V0KSB7XG4gICAgICAodGFyZ2V0IGFzIFJlY29yZDxzdHJpbmcsIHVua25vd24+KVtwcm9wZXJ0eV0gPSB2YWx1ZTtcbiAgICAgIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICByZXR1cm4gZmFsc2U7XG4gIH1cblxuICAvKipcbiAgICogU2V0IGEgQ1NTIGN1c3RvbSBwcm9wZXJ0eSBvbiBhbiBlbGVtZW50IGluIHRoZSBkZW1vXG4gICAqIEBwYXJhbSBzZWxlY3RvciAtIENTUyBzZWxlY3RvciBmb3IgdGFyZ2V0IGVsZW1lbnRcbiAgICogQHBhcmFtIGNzc1Byb3BlcnR5IC0gQ1NTIGN1c3RvbSBwcm9wZXJ0eSBuYW1lICh3aXRoIG9yIHdpdGhvdXQgLS0pXG4gICAqIEBwYXJhbSB2YWx1ZSAtIENTUyBwcm9wZXJ0eSB2YWx1ZVxuICAgKiBAcmV0dXJucyBXaGV0aGVyIHRoZSBvcGVyYXRpb24gc3VjY2VlZGVkXG4gICAqL1xuICBzZXREZW1vQ3NzQ3VzdG9tUHJvcGVydHkoc2VsZWN0b3I6IHN0cmluZywgY3NzUHJvcGVydHk6IHN0cmluZywgdmFsdWU6IHN0cmluZyk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHJvb3QgPSB0aGlzLnNoYWRvd1Jvb3QgPz8gdGhpcztcbiAgICBjb25zdCB0YXJnZXQgPSByb290LnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpIGFzIEhUTUxFbGVtZW50IHwgbnVsbDtcbiAgICBpZiAodGFyZ2V0KSB7XG4gICAgICBjb25zdCBwcm9wZXJ0eU5hbWUgPSBjc3NQcm9wZXJ0eS5zdGFydHNXaXRoKCctLScpID8gY3NzUHJvcGVydHkgOiBgLS0ke2Nzc1Byb3BlcnR5fWA7XG4gICAgICB0YXJnZXQuc3R5bGUuc2V0UHJvcGVydHkocHJvcGVydHlOYW1lLCB2YWx1ZSk7XG4gICAgICByZXR1cm4gdHJ1ZTtcbiAgICB9XG4gICAgcmV0dXJuIGZhbHNlO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1zZXJ2ZS1kZW1vJzogQ2VtU2VydmVEZW1vO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7OztBQ0Q5QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxzQ0FBd0MsQ0FBQztBQUFFLElBQU8seUJBQVE7OztBREEvRztBQVdBLDRCQUFDLGNBQWMsZ0JBQWdCO0FBQ3hCLElBQU0sZUFBTixlQUEyQixpQkFBVztBQUFBLEVBQzNDLE9BQU8sU0FBUztBQUFBLEVBRWhCLFNBQVM7QUFDUCxXQUFPO0FBQUEsRUFDVDtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBS0Esb0JBQW9CLFNBQWlCLGdCQUFnQixHQUFtQjtBQUN0RSxVQUFNLE9BQU8sS0FBSyxjQUFjO0FBQ2hDLFVBQU0sV0FBVyxLQUFLLGlCQUFpQixPQUFPO0FBQzlDLFdBQU8sU0FBUyxhQUFhLEtBQUs7QUFBQSxFQUNwQztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFZQSxnQkFDRSxNQUNBLE1BQ0EsT0FDQSxTQUNBLGdCQUFnQixHQUNQO0FBQ1QsVUFBTSxVQUFVLEtBQUssb0JBQW9CLFNBQVMsYUFBYTtBQUMvRCxRQUFJLENBQUMsU0FBUztBQUNaLGNBQVEsS0FBSyx1Q0FBdUMsU0FBUyxZQUFZLGFBQWE7QUFDdEYsYUFBTztBQUFBLElBQ1Q7QUFFQSxZQUFRLE1BQU07QUFB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
