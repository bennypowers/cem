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

// elements/cem-pf-v6-form-group/cem-pf-v6-form-group.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-pf-v6-form-group/cem-pf-v6-form-group.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: grid;\\n  grid-template-columns: subgrid;\\n  grid-column: 1 / -1;\\n  align-items: baseline;\\n}\\n\\n#label {\\n  grid-column: 1 / 2;\\n}\\n\\n#control {\\n  grid-column: 2 / 3;\\n  display: flex;\\n  gap: var(--pf-t--global--spacer--sm);\\n  align-items: start;\\n}\\n\\n#helper-text {\\n  grid-column: 1 / -1;\\n  font-size: var(--pf-t--global--font--size--sm);\\n  color: var(--pf-t--global--text--color--subtle);\\n  margin-top: var(--pf-t--global--spacer--xs);\\n\\n  \\u0026:empty {\\n    display: none;\\n  }\\n}\\n\\n::slotted([slot=control]:first-child) {\\n  flex: 1;\\n}\\n"'));
var cem_pf_v6_form_group_default = s;

// elements/cem-pf-v6-form-group/cem-pf-v6-form-group.ts
var _PfV6FormGroup_decorators, _init, _a;
_PfV6FormGroup_decorators = [customElement("cem-pf-v6-form-group")];
var PfV6FormGroup = class extends (_a = LitElement) {
  static styles = cem_pf_v6_form_group_default;
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LWZvcm0tZ3JvdXAvY2VtLXBmLXY2LWZvcm0tZ3JvdXAudHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tcGYtdjYtZm9ybS1ncm91cC9jZW0tcGYtdjYtZm9ybS1ncm91cC5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1wZi12Ni1mb3JtLWdyb3VwLmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBGb3JtIEdyb3VwIENvbXBvbmVudFxuICpcbiAqIEdyb3VwcyBhIGxhYmVsLCBmb3JtIGNvbnRyb2wsIGFuZCBvcHRpb25hbCBoZWxwZXIgdGV4dCB0b2dldGhlci5cbiAqXG4gKiBAc2xvdCBsYWJlbCAtIExhYmVsIGNvbnRlbnQgKHR5cGljYWxseSBjZW0tcGYtdjYtZm9ybS1sYWJlbClcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBmb3JtIGNvbnRyb2wgZWxlbWVudFxuICogQHNsb3QgaGVscGVyIC0gSGVscGVyIHRleHRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1wZi12Ni1mb3JtLWdyb3VwJylcbmV4cG9ydCBjbGFzcyBQZlY2Rm9ybUdyb3VwIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cImxhYmVsXCI+XG4gICAgICAgIDxzbG90IG5hbWU9XCJsYWJlbFwiPjwvc2xvdD5cbiAgICAgIDwvZGl2PlxuICAgICAgPGRpdiBpZD1cImNvbnRyb2xcIj5cbiAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgPC9kaXY+XG4gICAgICA8ZGl2IGlkPVwiaGVscGVyLXRleHRcIj5cbiAgICAgICAgPHNsb3QgbmFtZT1cImhlbHBlclwiPjwvc2xvdD5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXBmLXY2LWZvcm0tZ3JvdXAnOiBQZlY2Rm9ybUdyb3VwO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIGdyaWQtdGVtcGxhdGUtY29sdW1uczogc3ViZ3JpZDtcXFxcbiAgZ3JpZC1jb2x1bW46IDEgLyAtMTtcXFxcbiAgYWxpZ24taXRlbXM6IGJhc2VsaW5lO1xcXFxufVxcXFxuXFxcXG4jbGFiZWwge1xcXFxuICBncmlkLWNvbHVtbjogMSAvIDI7XFxcXG59XFxcXG5cXFxcbiNjb250cm9sIHtcXFxcbiAgZ3JpZC1jb2x1bW46IDIgLyAzO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBnYXA6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIGFsaWduLWl0ZW1zOiBzdGFydDtcXFxcbn1cXFxcblxcXFxuI2hlbHBlci10ZXh0IHtcXFxcbiAgZ3JpZC1jb2x1bW46IDEgLyAtMTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLXNtKTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIG1hcmdpbi10b3A6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS14cyk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmVtcHR5IHtcXFxcbiAgICBkaXNwbGF5OiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbjo6c2xvdHRlZChbc2xvdD1jb250cm9sXTpmaXJzdC1jaGlsZCkge1xcXFxuICBmbGV4OiAxO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7OztBQ0Q5QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxtbUJBQXFtQixDQUFDO0FBQUUsSUFBTywrQkFBUTs7O0FEQTVxQjtBQWNBLDZCQUFDLGNBQWMsc0JBQXNCO0FBQzlCLElBQU0sZ0JBQU4sZUFBNEIsaUJBQVc7QUFBQSxFQUM1QyxPQUFPLFNBQVM7QUFBQSxFQUVoQixTQUFTO0FBQ1AsV0FBTztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFXVDtBQUNGO0FBaEJPO0FBQU0sZ0JBQU4sNkNBRFAsMkJBQ2E7QUFBTiw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
