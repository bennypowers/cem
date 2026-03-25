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

// elements/pf-v6-page-main/pf-v6-page-main.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-page-main/pf-v6-page-main.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  z-index: var(--pf-v6-c-page__main-container--ZIndex);\\n  display: flex;\\n  flex-direction: column;\\n  grid-area: var(--pf-v6-c-page__main-container--GridArea);\\n  align-self: var(--pf-v6-c-page__main-container--AlignSelf);\\n  max-height: var(--pf-v6-c-page__main-container--MaxHeight);\\n  margin-inline-start: var(--pf-v6-c-page__main-container--MarginInlineStart);\\n  margin-inline-end: var(--pf-v6-c-page__main-container--MarginInlineEnd);\\n  overflow-x: hidden;\\n  overflow-y: auto;\\n  -webkit-overflow-scrolling: touch;\\n  background: var(--pf-v6-c-page__main-container--BackgroundColor);\\n  border: solid var(--pf-v6-c-page__main-container--BorderColor);\\n  border-block-start-width: var(--pf-v6-c-page__main-container--BorderBlockStartWidth);\\n  border-block-end-width: var(--pf-v6-c-page__main-container--BorderBlockEndWidth);\\n  border-inline-start-width: var(--pf-v6-c-page__main-container--BorderInlineStartWidth);\\n  border-inline-end-width: var(--pf-v6-c-page__main-container--BorderInlineEndWidth);\\n  border-radius: var(--pf-v6-c-page__main-container--BorderRadius);\\n}\\n\\nslot {\\n  display: flex;\\n  flex: 1;\\n  flex-direction: column;\\n  outline: 0;\\n}\\n"'));
var pf_v6_page_main_default = s;

// elements/pf-v6-page-main/pf-v6-page-main.ts
var _PfV6PageMain_decorators, _init, _a;
_PfV6PageMain_decorators = [customElement("pf-v6-page-main")];
var PfV6PageMain = class extends (_a = LitElement) {
  static styles = pf_v6_page_main_default;
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtcGFnZS1tYWluL3BmLXY2LXBhZ2UtbWFpbi50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL3BmLXY2LXBhZ2UtbWFpbi9wZi12Ni1wYWdlLW1haW4uY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9wZi12Ni1wYWdlLW1haW4uY3NzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IFBhZ2UgTWFpbiBDb250ZW50XG4gKlxuICogV3JhcHMgbWFpbiBjb250ZW50IGFyZWEgd2l0aCBwcm9wZXIgc2VtYW50aWNzIGFuZCBza2lwLXRvLWNvbnRlbnQgc3VwcG9ydC5cbiAqIFRoZSBpZCBhdHRyaWJ1dGUgc2hvdWxkIGJlIHNldCBvbiB0aGUgaG9zdCBlbGVtZW50IGZvciBza2lwLXRvLWNvbnRlbnQgbGlua3MuXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIG1haW4gY29udGVudFxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtcGFnZS1tYWluJylcbmV4cG9ydCBjbGFzcyBQZlY2UGFnZU1haW4gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcblxuICBjb25zdHJ1Y3RvcigpIHtcbiAgICBzdXBlcigpO1xuICAgIHRoaXMuI2ludGVybmFscy5yb2xlID0gJ21haW4nO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYDxzbG90Pjwvc2xvdD5gO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LXBhZ2UtbWFpbic6IFBmVjZQYWdlTWFpbjtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICB6LWluZGV4OiB2YXIoLS1wZi12Ni1jLXBhZ2VfX21haW4tY29udGFpbmVyLS1aSW5kZXgpO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBncmlkLWFyZWE6IHZhcigtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLUdyaWRBcmVhKTtcXFxcbiAgYWxpZ24tc2VsZjogdmFyKC0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tQWxpZ25TZWxmKTtcXFxcbiAgbWF4LWhlaWdodDogdmFyKC0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tTWF4SGVpZ2h0KTtcXFxcbiAgbWFyZ2luLWlubGluZS1zdGFydDogdmFyKC0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tTWFyZ2luSW5saW5lU3RhcnQpO1xcXFxuICBtYXJnaW4taW5saW5lLWVuZDogdmFyKC0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tTWFyZ2luSW5saW5lRW5kKTtcXFxcbiAgb3ZlcmZsb3cteDogaGlkZGVuO1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuICAtd2Via2l0LW92ZXJmbG93LXNjcm9sbGluZzogdG91Y2g7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIGJvcmRlcjogc29saWQgdmFyKC0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tQm9yZGVyQ29sb3IpO1xcXFxuICBib3JkZXItYmxvY2stc3RhcnQtd2lkdGg6IHZhcigtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLUJvcmRlckJsb2NrU3RhcnRXaWR0aCk7XFxcXG4gIGJvcmRlci1ibG9jay1lbmQtd2lkdGg6IHZhcigtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLUJvcmRlckJsb2NrRW5kV2lkdGgpO1xcXFxuICBib3JkZXItaW5saW5lLXN0YXJ0LXdpZHRoOiB2YXIoLS1wZi12Ni1jLXBhZ2VfX21haW4tY29udGFpbmVyLS1Cb3JkZXJJbmxpbmVTdGFydFdpZHRoKTtcXFxcbiAgYm9yZGVyLWlubGluZS1lbmQtd2lkdGg6IHZhcigtLXBmLXY2LWMtcGFnZV9fbWFpbi1jb250YWluZXItLUJvcmRlcklubGluZUVuZFdpZHRoKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdjYtYy1wYWdlX19tYWluLWNvbnRhaW5lci0tQm9yZGVyUmFkaXVzKTtcXFxcbn1cXFxcblxcXFxuc2xvdCB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXg6IDE7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIG91dGxpbmU6IDA7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjs7O0FDRDlCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLDRyQ0FBOHJDLENBQUM7QUFBRSxJQUFPLDBCQUFROzs7QURBcndDO0FBYUEsNEJBQUMsY0FBYyxpQkFBaUI7QUFDekIsSUFBTSxlQUFOLGVBQTJCLGlCQUFXO0FBQUEsRUFDM0MsT0FBTyxTQUFTO0FBQUEsRUFFaEIsYUFBYSxLQUFLLGdCQUFnQjtBQUFBLEVBRWxDLGNBQWM7QUFDWixVQUFNO0FBQ04sU0FBSyxXQUFXLE9BQU87QUFBQSxFQUN6QjtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQSxFQUNUO0FBQ0Y7QUFiTztBQUFNLGVBQU4sNENBRFAsMEJBQ2E7QUFBTiw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
