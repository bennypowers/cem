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

// elements/cem-pf-v6-page-main/cem-pf-v6-page-main.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-pf-v6-page-main/cem-pf-v6-page-main.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  z-index: var(--cem-pf-v6-c-page__main-container--ZIndex);\\n  display: flex;\\n  flex-direction: column;\\n  grid-area: var(--cem-pf-v6-c-page__main-container--GridArea);\\n  align-self: var(--cem-pf-v6-c-page__main-container--AlignSelf);\\n  max-height: var(--cem-pf-v6-c-page__main-container--MaxHeight);\\n  margin-inline-start: var(--cem-pf-v6-c-page__main-container--MarginInlineStart);\\n  margin-inline-end: var(--cem-pf-v6-c-page__main-container--MarginInlineEnd);\\n  overflow-x: hidden;\\n  overflow-y: auto;\\n  -webkit-overflow-scrolling: touch;\\n  background: var(--cem-pf-v6-c-page__main-container--BackgroundColor);\\n  border: solid var(--cem-pf-v6-c-page__main-container--BorderColor);\\n  border-block-start-width: var(--cem-pf-v6-c-page__main-container--BorderBlockStartWidth);\\n  border-block-end-width: var(--cem-pf-v6-c-page__main-container--BorderBlockEndWidth);\\n  border-inline-start-width: var(--cem-pf-v6-c-page__main-container--BorderInlineStartWidth);\\n  border-inline-end-width: var(--cem-pf-v6-c-page__main-container--BorderInlineEndWidth);\\n  border-radius: var(--cem-pf-v6-c-page__main-container--BorderRadius);\\n}\\n\\nslot {\\n  display: flex;\\n  flex: 1;\\n  flex-direction: column;\\n  outline: 0;\\n}\\n"'));
var cem_pf_v6_page_main_default = s;

// elements/cem-pf-v6-page-main/cem-pf-v6-page-main.ts
var _PfV6PageMain_decorators, _init, _a;
_PfV6PageMain_decorators = [customElement("cem-pf-v6-page-main")];
var PfV6PageMain = class extends (_a = LitElement) {
  static styles = cem_pf_v6_page_main_default;
  #internals = this.attachInternals();
  constructor() {
    super();
    this.#internals.role = "main";
  }
  render() {
    return html`<slot></slot>`;
  }
};
_init = __decoratorStart(_a);
PfV6PageMain = __decorateElement(_init, 0, "PfV6PageMain", _PfV6PageMain_decorators, PfV6PageMain);
__runInitializers(_init, 1, PfV6PageMain);
export {
  PfV6PageMain
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXBhZ2UtbWFpbi9jZW0tcGYtdjYtcGFnZS1tYWluLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LXBhZ2UtbWFpbi9jZW0tcGYtdjYtcGFnZS1tYWluLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXBmLXY2LXBhZ2UtbWFpbi5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgUGFnZSBNYWluIENvbnRlbnRcbiAqXG4gKiBXcmFwcyBtYWluIGNvbnRlbnQgYXJlYSB3aXRoIHByb3BlciBzZW1hbnRpY3MgYW5kIHNraXAtdG8tY29udGVudCBzdXBwb3J0LlxuICogVGhlIGlkIGF0dHJpYnV0ZSBzaG91bGQgYmUgc2V0IG9uIHRoZSBob3N0IGVsZW1lbnQgZm9yIHNraXAtdG8tY29udGVudCBsaW5rcy5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgbWFpbiBjb250ZW50XG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tcGYtdjYtcGFnZS1tYWluJylcbmV4cG9ydCBjbGFzcyBQZlY2UGFnZU1haW4gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcblxuICBjb25zdHJ1Y3RvcigpIHtcbiAgICBzdXBlcigpO1xuICAgIHRoaXMuI2ludGVybmFscy5yb2xlID0gJ21haW4nO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYDxzbG90Pjwvc2xvdD5gO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1wZi12Ni1wYWdlLW1haW4nOiBQZlY2UGFnZU1haW47XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgei1pbmRleDogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLVpJbmRleCk7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGdyaWQtYXJlYTogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLUdyaWRBcmVhKTtcXFxcbiAgYWxpZ24tc2VsZjogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLUFsaWduU2VsZik7XFxcXG4gIG1heC1oZWlnaHQ6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX21haW4tY29udGFpbmVyLS1NYXhIZWlnaHQpO1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tTWFyZ2luSW5saW5lU3RhcnQpO1xcXFxuICBtYXJnaW4taW5saW5lLWVuZDogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLU1hcmdpbklubGluZUVuZCk7XFxcXG4gIG92ZXJmbG93LXg6IGhpZGRlbjtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbiAgLXdlYmtpdC1vdmVyZmxvdy1zY3JvbGxpbmc6IHRvdWNoO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgYm9yZGVyOiBzb2xpZCB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tQm9yZGVyQ29sb3IpO1xcXFxuICBib3JkZXItYmxvY2stc3RhcnQtd2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX21haW4tY29udGFpbmVyLS1Cb3JkZXJCbG9ja1N0YXJ0V2lkdGgpO1xcXFxuICBib3JkZXItYmxvY2stZW5kLXdpZHRoOiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tQm9yZGVyQmxvY2tFbmRXaWR0aCk7XFxcXG4gIGJvcmRlci1pbmxpbmUtc3RhcnQtd2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX21haW4tY29udGFpbmVyLS1Cb3JkZXJJbmxpbmVTdGFydFdpZHRoKTtcXFxcbiAgYm9yZGVyLWlubGluZS1lbmQtd2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX21haW4tY29udGFpbmVyLS1Cb3JkZXJJbmxpbmVFbmRXaWR0aCk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX21haW4tY29udGFpbmVyLS1Cb3JkZXJSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG5zbG90IHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleDogMTtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgb3V0bGluZTogMDtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCOzs7QUNEOUIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sZ3ZDQUFrdkMsQ0FBQztBQUFFLElBQU8sOEJBQVE7OztBREF6ekM7QUFhQSw0QkFBQyxjQUFjLHFCQUFxQjtBQUM3QixJQUFNLGVBQU4sZUFBMkIsaUJBQVc7QUFBQSxFQUMzQyxPQUFPLFNBQVM7QUFBQSxFQUVoQixhQUFhLEtBQUssZ0JBQWdCO0FBQUEsRUFFbEMsY0FBYztBQUNaLFVBQU07QUFDTixTQUFLLFdBQVcsT0FBTztBQUFBLEVBQ3pCO0FBQUEsRUFFQSxTQUFTO0FBQ1AsV0FBTztBQUFBLEVBQ1Q7QUFDRjtBQWJPO0FBQU0sZUFBTiw0Q0FEUCwwQkFDYTtBQUFOLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
