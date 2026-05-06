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

// elements/cem-pf-v6-nav-list/cem-pf-v6-nav-list.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-pf-v6-nav-list/cem-pf-v6-nav-list.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  display: grid;\\n  row-gap: var(--cem-pf-v6-c-nav__list--RowGap, var(--pf-t--global--spacer--sm));\\n  column-gap: var(--cem-pf-v6-c-nav__list--ColumnGap, var(--pf-t--global--spacer--xs));\\n  min-height: 0;\\n}\\n"'));
var cem_pf_v6_nav_list_default = s;

// elements/cem-pf-v6-nav-list/cem-pf-v6-nav-list.ts
var _PfV6NavList_decorators, _init, _a;
_PfV6NavList_decorators = [customElement("cem-pf-v6-nav-list")];
var PfV6NavList = class extends (_a = LitElement) {
  static styles = cem_pf_v6_nav_list_default;
  #internals = this.attachInternals();
  constructor() {
    super();
    this.#internals.role = "list";
  }
  render() {
    return html`<slot></slot>`;
  }
};
_init = __decoratorStart(_a);
PfV6NavList = __decorateElement(_init, 0, "PfV6NavList", _PfV6NavList_decorators, PfV6NavList);
__runInitializers(_init, 1, PfV6NavList);
export {
  PfV6NavList
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LW5hdi1saXN0L2NlbS1wZi12Ni1uYXYtbGlzdC50cyIsICJsaXQtY3NzOmVsZW1lbnRzL2NlbS1wZi12Ni1uYXYtbGlzdC9jZW0tcGYtdjYtbmF2LWxpc3QuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtbmF2LWxpc3QuY3NzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IE5hdmlnYXRpb24gTGlzdFxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBuYXYtaXRlbSBvciBuYXYtZ3JvdXAgZWxlbWVudHNcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1wZi12Ni1uYXYtbGlzdCcpXG5leHBvcnQgY2xhc3MgUGZWNk5hdkxpc3QgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcblxuICBjb25zdHJ1Y3RvcigpIHtcbiAgICBzdXBlcigpO1xuICAgIHRoaXMuI2ludGVybmFscy5yb2xlID0gJ2xpc3QnO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYDxzbG90Pjwvc2xvdD5gO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1wZi12Ni1uYXYtbGlzdCc6IFBmVjZOYXZMaXN0O1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG5cXFxcbiAgZGlzcGxheTogZ3JpZDtcXFxcbiAgcm93LWdhcDogdmFyKC0tY2VtLXBmLXY2LWMtbmF2X19saXN0LS1Sb3dHYXAsIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSkpO1xcXFxuICBjb2x1bW4tZ2FwOiB2YXIoLS1jZW0tcGYtdjYtYy1uYXZfX2xpc3QtLUNvbHVtbkdhcCwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXhzKSk7XFxcXG4gIG1pbi1oZWlnaHQ6IDA7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjs7O0FDRDlCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLHlPQUEyTyxDQUFDO0FBQUUsSUFBTyw2QkFBUTs7O0FEQWxUO0FBVUEsMkJBQUMsY0FBYyxvQkFBb0I7QUFDNUIsSUFBTSxjQUFOLGVBQTBCLGlCQUFXO0FBQUEsRUFDMUMsT0FBTyxTQUFTO0FBQUEsRUFFaEIsYUFBYSxLQUFLLGdCQUFnQjtBQUFBLEVBRWxDLGNBQWM7QUFDWixVQUFNO0FBQ04sU0FBSyxXQUFXLE9BQU87QUFBQSxFQUN6QjtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQSxFQUNUO0FBQ0Y7QUFiTztBQUFNLGNBQU4sMkNBRFAseUJBQ2E7QUFBTiw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
