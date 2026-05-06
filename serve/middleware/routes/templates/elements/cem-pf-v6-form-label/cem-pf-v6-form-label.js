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

// elements/cem-pf-v6-form-label/cem-pf-v6-form-label.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-pf-v6-form-label/cem-pf-v6-form-label.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n}\\n\\nlabel {\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  font-size: var(--pf-t--global--font--size--body--default);\\n  color: var(--pf-t--global--text--color--regular);\\n  margin: 0;\\n  cursor: pointer;\\n}\\n\\nspan {\\n  display: inline;\\n}\\n"'));
var cem_pf_v6_form_label_default = s;

// elements/cem-pf-v6-form-label/cem-pf-v6-form-label.ts
var _PfV6FormLabel_decorators, _init, _a;
_PfV6FormLabel_decorators = [customElement("cem-pf-v6-form-label")];
var PfV6FormLabel = class extends (_a = LitElement) {
  static styles = cem_pf_v6_form_label_default;
  render() {
    return html`
      <label @click=${this.#handleClick}>
        <span><slot @slotchange=${this.#updateControlLabel}></slot></span>
      </label>
    `;
  }
  connectedCallback() {
    super.connectedCallback();
    requestAnimationFrame(() => this.#updateControlLabel());
  }
  #handleClick(e) {
    e.preventDefault();
    const control = this.#getControl();
    if (control && typeof control.focus === "function") {
      control.focus();
    }
  }
  /**
   * Updates the aria-label on the control's internal input element.
   *
   * Supports two approaches:
   * 1. Preferred: Control implements setAccessibleLabel(text) method (CemFormControl API)
   * 2. Fallback: Direct shadowRoot access for legacy controls
   *
   * Compatible controls should extend CemFormControl or implement setAccessibleLabel().
   */
  #updateControlLabel = () => {
    const control = this.#getControl();
    const labelText = this.textContent?.trim() || "";
    if (!control || !labelText) {
      return;
    }
    if (typeof control.setAccessibleLabel === "function") {
      control.setAccessibleLabel(labelText);
      return;
    }
    const shadowRoot = control.shadowRoot;
    if (!shadowRoot) {
      console.warn("Control has no shadow root and no setAccessibleLabel method");
      return;
    }
    const internalInput = shadowRoot.querySelector("input, select, textarea");
    if (internalInput) {
      internalInput.setAttribute("aria-label", labelText);
    } else {
      console.warn("Could not find internal input in control shadow DOM");
    }
  };
  #getControl() {
    const formGroup = this.closest("cem-pf-v6-form-group");
    if (!formGroup) {
      console.warn("cem-pf-v6-form-label must be inside cem-pf-v6-form-group");
      return null;
    }
    const control = formGroup.querySelector(":not([slot])");
    if (!control) {
      console.warn("No control found in cem-pf-v6-form-group");
      return null;
    }
    return control;
  }
};
_init = __decoratorStart(_a);
PfV6FormLabel = __decorateElement(_init, 0, "PfV6FormLabel", _PfV6FormLabel_decorators, PfV6FormLabel);
__runInitializers(_init, 1, PfV6FormLabel);
export {
  PfV6FormLabel
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LWZvcm0tbGFiZWwvY2VtLXBmLXY2LWZvcm0tbGFiZWwudHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tcGYtdjYtZm9ybS1sYWJlbC9jZW0tcGYtdjYtZm9ybS1sYWJlbC5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1wZi12Ni1mb3JtLWxhYmVsLmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBGb3JtIExhYmVsIENvbXBvbmVudFxuICpcbiAqIEF1dG9tYXRpY2FsbHkgYXNzb2NpYXRlcyB3aXRoIGNvbnRyb2wgaW4gcGFyZW50IGNlbS1wZi12Ni1mb3JtLWdyb3VwIGJ5OlxuICogLSBTZXR0aW5nIGFyaWEtbGFiZWwgb24gdGhlIGNvbnRyb2wgZnJvbSBsYWJlbCB0ZXh0IGNvbnRlbnRcbiAqIC0gRm9yd2FyZGluZyBjbGlja3MgdG8gZm9jdXMgdGhlIGNvbnRyb2xcbiAqXG4gKiBAc2xvdCAtIExhYmVsIHRleHQgY29udGVudFxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXBmLXY2LWZvcm0tbGFiZWwnKVxuZXhwb3J0IGNsYXNzIFBmVjZGb3JtTGFiZWwgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8bGFiZWwgQGNsaWNrPSR7dGhpcy4jaGFuZGxlQ2xpY2t9PlxuICAgICAgICA8c3Bhbj48c2xvdCBAc2xvdGNoYW5nZT0ke3RoaXMuI3VwZGF0ZUNvbnRyb2xMYWJlbH0+PC9zbG90Pjwvc3Bhbj5cbiAgICAgIDwvbGFiZWw+XG4gICAgYDtcbiAgfVxuXG4gIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmNvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgLy8gRGVmZXIgaW5pdGlhbCBsYWJlbCBzZXR1cCB1bnRpbCBuZXh0IGZyYW1lIHNvIHNsb3R0ZWQgY29udGVudCBpcyBhdmFpbGFibGVcbiAgICByZXF1ZXN0QW5pbWF0aW9uRnJhbWUoKCkgPT4gdGhpcy4jdXBkYXRlQ29udHJvbExhYmVsKCkpO1xuICB9XG5cbiAgI2hhbmRsZUNsaWNrKGU6IEV2ZW50KSB7XG4gICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuICAgIGNvbnN0IGNvbnRyb2wgPSB0aGlzLiNnZXRDb250cm9sKCk7XG4gICAgaWYgKGNvbnRyb2wgJiYgdHlwZW9mIGNvbnRyb2wuZm9jdXMgPT09ICdmdW5jdGlvbicpIHtcbiAgICAgIGNvbnRyb2wuZm9jdXMoKTtcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogVXBkYXRlcyB0aGUgYXJpYS1sYWJlbCBvbiB0aGUgY29udHJvbCdzIGludGVybmFsIGlucHV0IGVsZW1lbnQuXG4gICAqXG4gICAqIFN1cHBvcnRzIHR3byBhcHByb2FjaGVzOlxuICAgKiAxLiBQcmVmZXJyZWQ6IENvbnRyb2wgaW1wbGVtZW50cyBzZXRBY2Nlc3NpYmxlTGFiZWwodGV4dCkgbWV0aG9kIChDZW1Gb3JtQ29udHJvbCBBUEkpXG4gICAqIDIuIEZhbGxiYWNrOiBEaXJlY3Qgc2hhZG93Um9vdCBhY2Nlc3MgZm9yIGxlZ2FjeSBjb250cm9sc1xuICAgKlxuICAgKiBDb21wYXRpYmxlIGNvbnRyb2xzIHNob3VsZCBleHRlbmQgQ2VtRm9ybUNvbnRyb2wgb3IgaW1wbGVtZW50IHNldEFjY2Vzc2libGVMYWJlbCgpLlxuICAgKi9cbiAgI3VwZGF0ZUNvbnRyb2xMYWJlbCA9ICgpID0+IHtcbiAgICBjb25zdCBjb250cm9sID0gdGhpcy4jZ2V0Q29udHJvbCgpO1xuICAgIGNvbnN0IGxhYmVsVGV4dCA9IHRoaXMudGV4dENvbnRlbnQ/LnRyaW0oKSB8fCAnJztcblxuICAgIGlmICghY29udHJvbCB8fCAhbGFiZWxUZXh0KSB7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgLy8gUHJlZmVycmVkOiBVc2Ugc3RhbmRhcmQgQVBJIGlmIGF2YWlsYWJsZVxuICAgIGlmICh0eXBlb2YgKGNvbnRyb2wgYXMgYW55KS5zZXRBY2Nlc3NpYmxlTGFiZWwgPT09ICdmdW5jdGlvbicpIHtcbiAgICAgIChjb250cm9sIGFzIGFueSkuc2V0QWNjZXNzaWJsZUxhYmVsKGxhYmVsVGV4dCk7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgLy8gRmFsbGJhY2s6IERpcmVjdCBzaGFkb3dSb290IGFjY2VzcyBmb3IgbGVnYWN5IGNvbnRyb2xzXG4gICAgY29uc3Qgc2hhZG93Um9vdCA9IGNvbnRyb2wuc2hhZG93Um9vdDtcbiAgICBpZiAoIXNoYWRvd1Jvb3QpIHtcbiAgICAgIGNvbnNvbGUud2FybignQ29udHJvbCBoYXMgbm8gc2hhZG93IHJvb3QgYW5kIG5vIHNldEFjY2Vzc2libGVMYWJlbCBtZXRob2QnKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICBjb25zdCBpbnRlcm5hbElucHV0ID0gc2hhZG93Um9vdC5xdWVyeVNlbGVjdG9yKCdpbnB1dCwgc2VsZWN0LCB0ZXh0YXJlYScpO1xuICAgIGlmIChpbnRlcm5hbElucHV0KSB7XG4gICAgICBpbnRlcm5hbElucHV0LnNldEF0dHJpYnV0ZSgnYXJpYS1sYWJlbCcsIGxhYmVsVGV4dCk7XG4gICAgfSBlbHNlIHtcbiAgICAgIGNvbnNvbGUud2FybignQ291bGQgbm90IGZpbmQgaW50ZXJuYWwgaW5wdXQgaW4gY29udHJvbCBzaGFkb3cgRE9NJyk7XG4gICAgfVxuICB9O1xuXG4gICNnZXRDb250cm9sKCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICAvLyBGaW5kIHRoZSBjb250cm9sIGVsZW1lbnQgaW4gdGhlIHBhcmVudCBmb3JtIGdyb3VwXG4gICAgY29uc3QgZm9ybUdyb3VwID0gdGhpcy5jbG9zZXN0KCdjZW0tcGYtdjYtZm9ybS1ncm91cCcpO1xuICAgIGlmICghZm9ybUdyb3VwKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ2NlbS1wZi12Ni1mb3JtLWxhYmVsIG11c3QgYmUgaW5zaWRlIGNlbS1wZi12Ni1mb3JtLWdyb3VwJyk7XG4gICAgICByZXR1cm4gbnVsbDtcbiAgICB9XG5cbiAgICAvLyBHZXQgdGhlIGNvbnRyb2wgc2xvdCAoZGVmYXVsdCBzbG90IGNvbnRlbnQpXG4gICAgY29uc3QgY29udHJvbCA9IGZvcm1Hcm91cC5xdWVyeVNlbGVjdG9yKCc6bm90KFtzbG90XSknKTtcbiAgICBpZiAoIWNvbnRyb2wpIHtcbiAgICAgIGNvbnNvbGUud2FybignTm8gY29udHJvbCBmb3VuZCBpbiBjZW0tcGYtdjYtZm9ybS1ncm91cCcpO1xuICAgICAgcmV0dXJuIG51bGw7XG4gICAgfVxuXG4gICAgcmV0dXJuIGNvbnRyb2w7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXBmLXY2LWZvcm0tbGFiZWwnOiBQZlY2Rm9ybUxhYmVsO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG59XFxcXG5cXFxcbmxhYmVsIHtcXFxcbiAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0td2VpZ2h0LS1ib2R5LS1ib2xkKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWRlZmF1bHQpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxufVxcXFxuXFxcXG5zcGFuIHtcXFxcbiAgZGlzcGxheTogaW5saW5lO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7OztBQ0Q5QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxpU0FBbVMsQ0FBQztBQUFFLElBQU8sK0JBQVE7OztBREExVztBQWNBLDZCQUFDLGNBQWMsc0JBQXNCO0FBQzlCLElBQU0sZ0JBQU4sZUFBNEIsaUJBQVc7QUFBQSxFQUM1QyxPQUFPLFNBQVM7QUFBQSxFQUVoQixTQUFTO0FBQ1AsV0FBTztBQUFBLHNCQUNXLEtBQUssWUFBWTtBQUFBLGtDQUNMLEtBQUssbUJBQW1CO0FBQUE7QUFBQTtBQUFBLEVBR3hEO0FBQUEsRUFFQSxvQkFBb0I7QUFDbEIsVUFBTSxrQkFBa0I7QUFFeEIsMEJBQXNCLE1BQU0sS0FBSyxvQkFBb0IsQ0FBQztBQUFBLEVBQ3hEO0FBQUEsRUFFQSxhQUFhLEdBQVU7QUFDckIsTUFBRSxlQUFlO0FBQ2pCLFVBQU0sVUFBVSxLQUFLLFlBQVk7QUFDakMsUUFBSSxXQUFXLE9BQU8sUUFBUSxVQUFVLFlBQVk7QUFDbEQsY0FBUSxNQUFNO0FBQUEsSUFDaEI7QUFBQSxFQUNGO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFXQSxzQkFBc0IsTUFBTTtBQUMxQixVQUFNLFVBQVUsS0FBSyxZQUFZO0FBQ2pDLFVBQU0sWUFBWSxLQUFLLGFBQWEsS0FBSyxLQUFLO0FBRTlDLFFBQUksQ0FBQyxXQUFXLENBQUMsV0FBVztBQUMxQjtBQUFBLElBQ0Y7QUFHQSxRQUFJLE9BQVEsUUFBZ0IsdUJBQXVCLFlBQVk7QUFDN0QsTUFBQyxRQUFnQixtQkFBbUIsU0FBUztBQUM3QztBQUFBLElBQ0Y7QUFHQSxVQUFNLGFBQWEsUUFBUTtBQUMzQixRQUFJLENBQUMsWUFBWTtBQUNmLGNBQVEsS0FBSyw2REFBNkQ7QUFDMUU7QUFBQSxJQUNGO0FBRUEsVUFBTSxnQkFBZ0IsV0FBVyxjQUFjLHlCQUF5QjtBQUN4RSxRQUFJLGVBQWU7QUFDakIsb0JBQWMsYUFBYSxjQUFjLFNBQVM7QUFBQSxJQUNwRCxPQUFPO0FBQ0wsY0FBUSxLQUFLLHFEQUFxRDtBQUFBLElBQ3BFO0FBQUEsRUFDRjtBQUFBLEVBRUEsY0FBOEI7QUFFNUIsVUFBTSxZQUFZLEtBQUssUUFBUSxzQkFBc0I7QUFDckQsUUFBSSxDQUFDLFdBQVc7QUFDZCxjQUFRLEtBQUssMERBQTBEO0FBQ3ZFLGFBQU87QUFBQSxJQUNUO0FBR0EsVUFBTSxVQUFVLFVBQVUsY0FBYyxjQUFjO0FBQ3RELFFBQUksQ0FBQyxTQUFTO0FBQ1osY0FBUSxLQUFLLDBDQUEwQztBQUN2RCxhQUFPO0FBQUEsSUFDVDtBQUVBLFdBQU87QUFBQSxFQUNUO0FBQ0Y7QUFoRk87QUFBTSxnQkFBTiw2Q0FEUCwyQkFDYTtBQUFOLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
