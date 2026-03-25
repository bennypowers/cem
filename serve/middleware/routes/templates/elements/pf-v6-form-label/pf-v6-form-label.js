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

// elements/pf-v6-form-label/pf-v6-form-label.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-form-label/pf-v6-form-label.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n}\\n\\nlabel {\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  font-size: var(--pf-t--global--font--size--body--default);\\n  color: var(--pf-t--global--text--color--regular);\\n  margin: 0;\\n  cursor: pointer;\\n}\\n\\nspan {\\n  display: inline;\\n}\\n"'));
var pf_v6_form_label_default = s;

// elements/pf-v6-form-label/pf-v6-form-label.ts
var _PfV6FormLabel_decorators, _init, _a;
_PfV6FormLabel_decorators = [customElement("pf-v6-form-label")];
var PfV6FormLabel = class extends (_a = LitElement) {
  static styles = pf_v6_form_label_default;
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
    const formGroup = this.closest("pf-v6-form-group");
    if (!formGroup) {
      console.warn("pf-v6-form-label must be inside pf-v6-form-group");
      return null;
    }
    const control = formGroup.querySelector(":not([slot])");
    if (!control) {
      console.warn("No control found in pf-v6-form-group");
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtZm9ybS1sYWJlbC9wZi12Ni1mb3JtLWxhYmVsLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvcGYtdjYtZm9ybS1sYWJlbC9wZi12Ni1mb3JtLWxhYmVsLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtZm9ybS1sYWJlbC5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgRm9ybSBMYWJlbCBDb21wb25lbnRcbiAqXG4gKiBBdXRvbWF0aWNhbGx5IGFzc29jaWF0ZXMgd2l0aCBjb250cm9sIGluIHBhcmVudCBwZi12Ni1mb3JtLWdyb3VwIGJ5OlxuICogLSBTZXR0aW5nIGFyaWEtbGFiZWwgb24gdGhlIGNvbnRyb2wgZnJvbSBsYWJlbCB0ZXh0IGNvbnRlbnRcbiAqIC0gRm9yd2FyZGluZyBjbGlja3MgdG8gZm9jdXMgdGhlIGNvbnRyb2xcbiAqXG4gKiBAc2xvdCAtIExhYmVsIHRleHQgY29udGVudFxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtZm9ybS1sYWJlbCcpXG5leHBvcnQgY2xhc3MgUGZWNkZvcm1MYWJlbCBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxsYWJlbCBAY2xpY2s9JHt0aGlzLiNoYW5kbGVDbGlja30+XG4gICAgICAgIDxzcGFuPjxzbG90IEBzbG90Y2hhbmdlPSR7dGhpcy4jdXBkYXRlQ29udHJvbExhYmVsfT48L3Nsb3Q+PC9zcGFuPlxuICAgICAgPC9sYWJlbD5cbiAgICBgO1xuICB9XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICAvLyBEZWZlciBpbml0aWFsIGxhYmVsIHNldHVwIHVudGlsIG5leHQgZnJhbWUgc28gc2xvdHRlZCBjb250ZW50IGlzIGF2YWlsYWJsZVxuICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB0aGlzLiN1cGRhdGVDb250cm9sTGFiZWwoKSk7XG4gIH1cblxuICAjaGFuZGxlQ2xpY2soZTogRXZlbnQpIHtcbiAgICBlLnByZXZlbnREZWZhdWx0KCk7XG4gICAgY29uc3QgY29udHJvbCA9IHRoaXMuI2dldENvbnRyb2woKTtcbiAgICBpZiAoY29udHJvbCAmJiB0eXBlb2YgY29udHJvbC5mb2N1cyA9PT0gJ2Z1bmN0aW9uJykge1xuICAgICAgY29udHJvbC5mb2N1cygpO1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBVcGRhdGVzIHRoZSBhcmlhLWxhYmVsIG9uIHRoZSBjb250cm9sJ3MgaW50ZXJuYWwgaW5wdXQgZWxlbWVudC5cbiAgICpcbiAgICogU3VwcG9ydHMgdHdvIGFwcHJvYWNoZXM6XG4gICAqIDEuIFByZWZlcnJlZDogQ29udHJvbCBpbXBsZW1lbnRzIHNldEFjY2Vzc2libGVMYWJlbCh0ZXh0KSBtZXRob2QgKENlbUZvcm1Db250cm9sIEFQSSlcbiAgICogMi4gRmFsbGJhY2s6IERpcmVjdCBzaGFkb3dSb290IGFjY2VzcyBmb3IgbGVnYWN5IGNvbnRyb2xzXG4gICAqXG4gICAqIENvbXBhdGlibGUgY29udHJvbHMgc2hvdWxkIGV4dGVuZCBDZW1Gb3JtQ29udHJvbCBvciBpbXBsZW1lbnQgc2V0QWNjZXNzaWJsZUxhYmVsKCkuXG4gICAqL1xuICAjdXBkYXRlQ29udHJvbExhYmVsID0gKCkgPT4ge1xuICAgIGNvbnN0IGNvbnRyb2wgPSB0aGlzLiNnZXRDb250cm9sKCk7XG4gICAgY29uc3QgbGFiZWxUZXh0ID0gdGhpcy50ZXh0Q29udGVudD8udHJpbSgpIHx8ICcnO1xuXG4gICAgaWYgKCFjb250cm9sIHx8ICFsYWJlbFRleHQpIHtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICAvLyBQcmVmZXJyZWQ6IFVzZSBzdGFuZGFyZCBBUEkgaWYgYXZhaWxhYmxlXG4gICAgaWYgKHR5cGVvZiAoY29udHJvbCBhcyBhbnkpLnNldEFjY2Vzc2libGVMYWJlbCA9PT0gJ2Z1bmN0aW9uJykge1xuICAgICAgKGNvbnRyb2wgYXMgYW55KS5zZXRBY2Nlc3NpYmxlTGFiZWwobGFiZWxUZXh0KTtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICAvLyBGYWxsYmFjazogRGlyZWN0IHNoYWRvd1Jvb3QgYWNjZXNzIGZvciBsZWdhY3kgY29udHJvbHNcbiAgICBjb25zdCBzaGFkb3dSb290ID0gY29udHJvbC5zaGFkb3dSb290O1xuICAgIGlmICghc2hhZG93Um9vdCkge1xuICAgICAgY29uc29sZS53YXJuKCdDb250cm9sIGhhcyBubyBzaGFkb3cgcm9vdCBhbmQgbm8gc2V0QWNjZXNzaWJsZUxhYmVsIG1ldGhvZCcpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IGludGVybmFsSW5wdXQgPSBzaGFkb3dSb290LnF1ZXJ5U2VsZWN0b3IoJ2lucHV0LCBzZWxlY3QsIHRleHRhcmVhJyk7XG4gICAgaWYgKGludGVybmFsSW5wdXQpIHtcbiAgICAgIGludGVybmFsSW5wdXQuc2V0QXR0cmlidXRlKCdhcmlhLWxhYmVsJywgbGFiZWxUZXh0KTtcbiAgICB9IGVsc2Uge1xuICAgICAgY29uc29sZS53YXJuKCdDb3VsZCBub3QgZmluZCBpbnRlcm5hbCBpbnB1dCBpbiBjb250cm9sIHNoYWRvdyBET00nKTtcbiAgICB9XG4gIH07XG5cbiAgI2dldENvbnRyb2woKTogRWxlbWVudCB8IG51bGwge1xuICAgIC8vIEZpbmQgdGhlIGNvbnRyb2wgZWxlbWVudCBpbiB0aGUgcGFyZW50IGZvcm0gZ3JvdXBcbiAgICBjb25zdCBmb3JtR3JvdXAgPSB0aGlzLmNsb3Nlc3QoJ3BmLXY2LWZvcm0tZ3JvdXAnKTtcbiAgICBpZiAoIWZvcm1Hcm91cCkge1xuICAgICAgY29uc29sZS53YXJuKCdwZi12Ni1mb3JtLWxhYmVsIG11c3QgYmUgaW5zaWRlIHBmLXY2LWZvcm0tZ3JvdXAnKTtcbiAgICAgIHJldHVybiBudWxsO1xuICAgIH1cblxuICAgIC8vIEdldCB0aGUgY29udHJvbCBzbG90IChkZWZhdWx0IHNsb3QgY29udGVudClcbiAgICBjb25zdCBjb250cm9sID0gZm9ybUdyb3VwLnF1ZXJ5U2VsZWN0b3IoJzpub3QoW3Nsb3RdKScpO1xuICAgIGlmICghY29udHJvbCkge1xuICAgICAgY29uc29sZS53YXJuKCdObyBjb250cm9sIGZvdW5kIGluIHBmLXY2LWZvcm0tZ3JvdXAnKTtcbiAgICAgIHJldHVybiBudWxsO1xuICAgIH1cblxuICAgIHJldHVybiBjb250cm9sO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LWZvcm0tbGFiZWwnOiBQZlY2Rm9ybUxhYmVsO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG59XFxcXG5cXFxcbmxhYmVsIHtcXFxcbiAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0td2VpZ2h0LS1ib2R5LS1ib2xkKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWRlZmF1bHQpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxufVxcXFxuXFxcXG5zcGFuIHtcXFxcbiAgZGlzcGxheTogaW5saW5lO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7OztBQ0Q5QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxpU0FBbVMsQ0FBQztBQUFFLElBQU8sMkJBQVE7OztBREExVztBQWNBLDZCQUFDLGNBQWMsa0JBQWtCO0FBQzFCLElBQU0sZ0JBQU4sZUFBNEIsaUJBQVc7QUFBQSxFQUM1QyxPQUFPLFNBQVM7QUFBQSxFQUVoQixTQUFTO0FBQ1AsV0FBTztBQUFBLHNCQUNXLEtBQUssWUFBWTtBQUFBLGtDQUNMLEtBQUssbUJBQW1CO0FBQUE7QUFBQTtBQUFBLEVBR3hEO0FBQUEsRUFFQSxvQkFBb0I7QUFDbEIsVUFBTSxrQkFBa0I7QUFFeEIsMEJBQXNCLE1BQU0sS0FBSyxvQkFBb0IsQ0FBQztBQUFBLEVBQ3hEO0FBQUEsRUFFQSxhQUFhLEdBQVU7QUFDckIsTUFBRSxlQUFlO0FBQ2pCLFVBQU0sVUFBVSxLQUFLLFlBQVk7QUFDakMsUUFBSSxXQUFXLE9BQU8sUUFBUSxVQUFVLFlBQVk7QUFDbEQsY0FBUSxNQUFNO0FBQUEsSUFDaEI7QUFBQSxFQUNGO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFXQSxzQkFBc0IsTUFBTTtBQUMxQixVQUFNLFVBQVUsS0FBSyxZQUFZO0FBQ2pDLFVBQU0sWUFBWSxLQUFLLGFBQWEsS0FBSyxLQUFLO0FBRTlDLFFBQUksQ0FBQyxXQUFXLENBQUMsV0FBVztBQUMxQjtBQUFBLElBQ0Y7QUFHQSxRQUFJLE9BQVEsUUFBZ0IsdUJBQXVCLFlBQVk7QUFDN0QsTUFBQyxRQUFnQixtQkFBbUIsU0FBUztBQUM3QztBQUFBLElBQ0Y7QUFHQSxVQUFNLGFBQWEsUUFBUTtBQUMzQixRQUFJLENBQUMsWUFBWTtBQUNmLGNBQVEsS0FBSyw2REFBNkQ7QUFDMUU7QUFBQSxJQUNGO0FBRUEsVUFBTSxnQkFBZ0IsV0FBVyxjQUFjLHlCQUF5QjtBQUN4RSxRQUFJLGVBQWU7QUFDakIsb0JBQWMsYUFBYSxjQUFjLFNBQVM7QUFBQSxJQUNwRCxPQUFPO0FBQ0wsY0FBUSxLQUFLLHFEQUFxRDtBQUFBLElBQ3BFO0FBQUEsRUFDRjtBQUFBLEVBRUEsY0FBOEI7QUFFNUIsVUFBTSxZQUFZLEtBQUssUUFBUSxrQkFBa0I7QUFDakQsUUFBSSxDQUFDLFdBQVc7QUFDZCxjQUFRLEtBQUssa0RBQWtEO0FBQy9ELGFBQU87QUFBQSxJQUNUO0FBR0EsVUFBTSxVQUFVLFVBQVUsY0FBYyxjQUFjO0FBQ3RELFFBQUksQ0FBQyxTQUFTO0FBQ1osY0FBUSxLQUFLLHNDQUFzQztBQUNuRCxhQUFPO0FBQUEsSUFDVDtBQUVBLFdBQU87QUFBQSxFQUNUO0FBQ0Y7QUFoRk87QUFBTSxnQkFBTiw2Q0FEUCwyQkFDYTtBQUFOLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
