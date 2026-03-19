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

// elements/pf-v6-form-group/pf-v6-form-group.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-form-group/pf-v6-form-group.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: grid;\\n  grid-template-columns: subgrid;\\n  grid-column: 1 / -1;\\n  align-items: baseline;\\n}\\n\\n#label {\\n  grid-column: 1 / 2;\\n}\\n\\n#control {\\n  grid-column: 2 / 3;\\n  display: flex;\\n  gap: var(--pf-t--global--spacer--sm);\\n  align-items: start;\\n}\\n\\n#helper-text {\\n  grid-column: 1 / -1;\\n  font-size: var(--pf-t--global--font--size--sm);\\n  color: var(--pf-t--global--text--color--subtle);\\n  margin-top: var(--pf-t--global--spacer--xs);\\n\\n  \\u0026:empty {\\n    display: none;\\n  }\\n}\\n\\n::slotted([slot=control]:first-child) {\\n  flex: 1;\\n}\\n"'));
var pf_v6_form_group_default = s;

// elements/pf-v6-form-group/pf-v6-form-group.ts
var _PfV6FormGroup_decorators, _init, _a;
_PfV6FormGroup_decorators = [customElement("pf-v6-form-group")];
var PfV6FormGroup = class extends (_a = LitElement) {
  static styles = pf_v6_form_group_default;
  render() {
    return html`
      <div id="label">
        <slot name="label"></slot>
      </div>
      <div id="control">
        <slot></slot>
      </div>
      <div id="helper-text">
        <slot name="helper"></slot>
      </div>
    `;
  }
};
_init = __decoratorStart(_a);
PfV6FormGroup = __decorateElement(_init, 0, "PfV6FormGroup", _PfV6FormGroup_decorators, PfV6FormGroup);
__runInitializers(_init, 1, PfV6FormGroup);
export {
  PfV6FormGroup
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtZm9ybS1ncm91cC9wZi12Ni1mb3JtLWdyb3VwLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvcGYtdjYtZm9ybS1ncm91cC9wZi12Ni1mb3JtLWdyb3VwLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtZm9ybS1ncm91cC5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgRm9ybSBHcm91cCBDb21wb25lbnRcbiAqXG4gKiBHcm91cHMgYSBsYWJlbCwgZm9ybSBjb250cm9sLCBhbmQgb3B0aW9uYWwgaGVscGVyIHRleHQgdG9nZXRoZXIuXG4gKlxuICogQHNsb3QgbGFiZWwgLSBMYWJlbCBjb250ZW50ICh0eXBpY2FsbHkgcGYtdjYtZm9ybS1sYWJlbClcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBmb3JtIGNvbnRyb2wgZWxlbWVudFxuICogQHNsb3QgaGVscGVyIC0gSGVscGVyIHRleHRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ3BmLXY2LWZvcm0tZ3JvdXAnKVxuZXhwb3J0IGNsYXNzIFBmVjZGb3JtR3JvdXAgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8ZGl2IGlkPVwibGFiZWxcIj5cbiAgICAgICAgPHNsb3QgbmFtZT1cImxhYmVsXCI+PC9zbG90PlxuICAgICAgPC9kaXY+XG4gICAgICA8ZGl2IGlkPVwiY29udHJvbFwiPlxuICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICA8L2Rpdj5cbiAgICAgIDxkaXYgaWQ9XCJoZWxwZXItdGV4dFwiPlxuICAgICAgICA8c2xvdCBuYW1lPVwiaGVscGVyXCI+PC9zbG90PlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdwZi12Ni1mb3JtLWdyb3VwJzogUGZWNkZvcm1Hcm91cDtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBncmlkO1xcXFxuICBncmlkLXRlbXBsYXRlLWNvbHVtbnM6IHN1YmdyaWQ7XFxcXG4gIGdyaWQtY29sdW1uOiAxIC8gLTE7XFxcXG4gIGFsaWduLWl0ZW1zOiBiYXNlbGluZTtcXFxcbn1cXFxcblxcXFxuI2xhYmVsIHtcXFxcbiAgZ3JpZC1jb2x1bW46IDEgLyAyO1xcXFxufVxcXFxuXFxcXG4jY29udHJvbCB7XFxcXG4gIGdyaWQtY29sdW1uOiAyIC8gMztcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICBhbGlnbi1pdGVtczogc3RhcnQ7XFxcXG59XFxcXG5cXFxcbiNoZWxwZXItdGV4dCB7XFxcXG4gIGdyaWQtY29sdW1uOiAxIC8gLTE7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1zbSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBtYXJnaW4tdG9wOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teHMpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjplbXB0eSB7XFxcXG4gICAgZGlzcGxheTogbm9uZTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG46OnNsb3R0ZWQoW3Nsb3Q9Y29udHJvbF06Zmlyc3QtY2hpbGQpIHtcXFxcbiAgZmxleDogMTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCOzs7QUNEOUIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sbW1CQUFxbUIsQ0FBQztBQUFFLElBQU8sMkJBQVE7OztBREE1cUI7QUFjQSw2QkFBQyxjQUFjLGtCQUFrQjtBQUMxQixJQUFNLGdCQUFOLGVBQTRCLGlCQUFXO0FBQUEsRUFDNUMsT0FBTyxTQUFTO0FBQUEsRUFFaEIsU0FBUztBQUNQLFdBQU87QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBV1Q7QUFDRjtBQWhCTztBQUFNLGdCQUFOLDZDQURQLDJCQUNhO0FBQU4sNEJBQU07IiwKICAibmFtZXMiOiBbXQp9Cg==
