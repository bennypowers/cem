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
var __publicField = (obj, key, value) => __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
var __accessCheck = (obj, member, msg) => member.has(obj) || __typeError("Cannot " + msg);
var __privateIn = (member, obj) => Object(obj) !== obj ? __typeError('Cannot use the "in" operator on this value') : member.has(obj);
var __privateGet = (obj, member, getter) => (__accessCheck(obj, member, "read from private field"), getter ? getter.call(obj) : member.get(obj));
var __privateAdd = (obj, member, value) => member.has(obj) ? __typeError("Cannot add the same private member more than once") : member instanceof WeakSet ? member.add(obj) : member.set(obj, value);
var __privateSet = (obj, member, value, setter) => (__accessCheck(obj, member, "write to private field"), setter ? setter.call(obj, value) : member.set(obj, value), value);
var __privateMethod = (obj, member, method) => (__accessCheck(obj, member, "access private method"), method);

// elements/pf-v6-toolbar-item/pf-v6-toolbar-item.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-toolbar-item/pf-v6-toolbar-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --pf-v6-hidden-visible--visible--Display: flex;\\n  --pf-v6-hidden-visible--hidden--Display: none;\\n  --pf-v6-hidden-visible--Display: var(--pf-v6-hidden-visible--visible--Display);\\n\\n  display: var(--pf-v6-hidden-visible--Display);\\n  row-gap: var(--pf-v6-c-toolbar__item--RowGap);\\n  column-gap: var(--pf-v6-c-toolbar__item--ColumnGap);\\n  width: var(--pf-v6-c-toolbar__item--Width);\\n  min-width: var(--pf-v6-c-toolbar__item--MinWidth);\\n  align-items: baseline;\\n}\\n\\n:host([variant=\\"label\\"]) {\\n  font-weight: var(--pf-v6-c-toolbar__item--m-label--FontWeight);\\n}\\n\\n:host([variant=\\"pagination\\"]) {\\n  margin-inline-start: auto;\\n}\\n\\n:host([variant=\\"overflow-container\\"]) {\\n  flex: 1;\\n  min-width: var(--pf-v6-c-toolbar__item--m-overflow-container--MinWidth);\\n  overflow: auto;\\n}\\n\\n:host([variant=\\"expand-all\\"][expanded]) ::slotted(.pf-v6-c-toolbar__expand-all-icon) {\\n  transform: rotate(var(--pf-v6-c-toolbar__item--m-expand-all--m-expanded__expand-all-icon--Rotate));\\n}\\n"'));
var pf_v6_toolbar_item_default = s;

// elements/pf-v6-toolbar-item/pf-v6-toolbar-item.ts
var _expanded_dec, _variant_dec, _a, _PfV6ToolbarItem_decorators, _init, _variant, _expanded;
_PfV6ToolbarItem_decorators = [customElement("pf-v6-toolbar-item")];
var PfV6ToolbarItem = class extends (_a = LitElement, _variant_dec = [property({ reflect: true })], _expanded_dec = [property({ type: Boolean, reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _variant, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
    __privateAdd(this, _expanded, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
  }
  render() {
    return html`<slot></slot>`;
  }
};
_init = __decoratorStart(_a);
_variant = new WeakMap();
_expanded = new WeakMap();
__decorateElement(_init, 4, "variant", _variant_dec, PfV6ToolbarItem, _variant);
__decorateElement(_init, 4, "expanded", _expanded_dec, PfV6ToolbarItem, _expanded);
PfV6ToolbarItem = __decorateElement(_init, 0, "PfV6ToolbarItem", _PfV6ToolbarItem_decorators, PfV6ToolbarItem);
__publicField(PfV6ToolbarItem, "styles", pf_v6_toolbar_item_default);
__runInitializers(_init, 1, PfV6ToolbarItem);
export {
  PfV6ToolbarItem
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtdG9vbGJhci1pdGVtL3BmLXY2LXRvb2xiYXItaXRlbS50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL3BmLXY2LXRvb2xiYXItaXRlbS9wZi12Ni10b29sYmFyLWl0ZW0uY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL3BmLXY2LXRvb2xiYXItaXRlbS5jc3MnO1xuXG50eXBlIFRvb2xiYXJJdGVtVmFyaWFudCA9XG4gIHwgJ2xhYmVsJ1xuICB8ICdwYWdpbmF0aW9uJ1xuICB8ICdvdmVyZmxvdy1jb250YWluZXInXG4gIHwgJ2V4cGFuZC1hbGwnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgVG9vbGJhciBJdGVtXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGl0ZW0gY29udGVudFxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtdG9vbGJhci1pdGVtJylcbmV4cG9ydCBjbGFzcyBQZlY2VG9vbGJhckl0ZW0gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIHZhcmlhbnQ/OiBUb29sYmFySXRlbVZhcmlhbnQ7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBleHBhbmRlZCA9IGZhbHNlO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGA8c2xvdD48L3Nsb3Q+YDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdwZi12Ni10b29sYmFyLWl0ZW0nOiBQZlY2VG9vbGJhckl0ZW07XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcblxcXFxuICAtLXBmLXY2LWhpZGRlbi12aXNpYmxlLS12aXNpYmxlLS1EaXNwbGF5OiBmbGV4O1xcXFxuICAtLXBmLXY2LWhpZGRlbi12aXNpYmxlLS1oaWRkZW4tLURpc3BsYXk6IG5vbmU7XFxcXG4gIC0tcGYtdjYtaGlkZGVuLXZpc2libGUtLURpc3BsYXk6IHZhcigtLXBmLXY2LWhpZGRlbi12aXNpYmxlLS12aXNpYmxlLS1EaXNwbGF5KTtcXFxcblxcXFxuICBkaXNwbGF5OiB2YXIoLS1wZi12Ni1oaWRkZW4tdmlzaWJsZS0tRGlzcGxheSk7XFxcXG4gIHJvdy1nYXA6IHZhcigtLXBmLXY2LWMtdG9vbGJhcl9faXRlbS0tUm93R2FwKTtcXFxcbiAgY29sdW1uLWdhcDogdmFyKC0tcGYtdjYtYy10b29sYmFyX19pdGVtLS1Db2x1bW5HYXApO1xcXFxuICB3aWR0aDogdmFyKC0tcGYtdjYtYy10b29sYmFyX19pdGVtLS1XaWR0aCk7XFxcXG4gIG1pbi13aWR0aDogdmFyKC0tcGYtdjYtYy10b29sYmFyX19pdGVtLS1NaW5XaWR0aCk7XFxcXG4gIGFsaWduLWl0ZW1zOiBiYXNlbGluZTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3ZhcmlhbnQ9XFxcXFxcXCJsYWJlbFxcXFxcXFwiXSkge1xcXFxuICBmb250LXdlaWdodDogdmFyKC0tcGYtdjYtYy10b29sYmFyX19pdGVtLS1tLWxhYmVsLS1Gb250V2VpZ2h0KTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3ZhcmlhbnQ9XFxcXFxcXCJwYWdpbmF0aW9uXFxcXFxcXCJdKSB7XFxcXG4gIG1hcmdpbi1pbmxpbmUtc3RhcnQ6IGF1dG87XFxcXG59XFxcXG5cXFxcbjpob3N0KFt2YXJpYW50PVxcXFxcXFwib3ZlcmZsb3ctY29udGFpbmVyXFxcXFxcXCJdKSB7XFxcXG4gIGZsZXg6IDE7XFxcXG4gIG1pbi13aWR0aDogdmFyKC0tcGYtdjYtYy10b29sYmFyX19pdGVtLS1tLW92ZXJmbG93LWNvbnRhaW5lci0tTWluV2lkdGgpO1xcXFxuICBvdmVyZmxvdzogYXV0bztcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3ZhcmlhbnQ9XFxcXFxcXCJleHBhbmQtYWxsXFxcXFxcXCJdW2V4cGFuZGVkXSkgOjpzbG90dGVkKC5wZi12Ni1jLXRvb2xiYXJfX2V4cGFuZC1hbGwtaWNvbikge1xcXFxuICB0cmFuc2Zvcm06IHJvdGF0ZSh2YXIoLS1wZi12Ni1jLXRvb2xiYXJfX2l0ZW0tLW0tZXhwYW5kLWFsbC0tbS1leHBhbmRlZF9fZXhwYW5kLWFsbC1pY29uLS1Sb3RhdGUpKTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSwyaENBQXFpQyxDQUFDO0FBQUUsSUFBTyw2QkFBUTs7O0FEQTVtQztBQWlCQSwrQkFBQyxjQUFjLG9CQUFvQjtBQUM1QixJQUFNLGtCQUFOLGVBQThCLGlCQUduQyxnQkFBQyxTQUFTLEVBQUUsU0FBUyxLQUFLLENBQUMsSUFHM0IsaUJBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQU5QLElBQVc7QUFBQSxFQUF6QztBQUFBO0FBSUwsdUJBQVMsVUFBVDtBQUdBLHVCQUFTLFdBQVcsa0JBQXBCLGlCQUFvQixTQUFwQjtBQUFBO0FBQUEsRUFFQSxTQUFTO0FBQ1AsV0FBTztBQUFBLEVBQ1Q7QUFDRjtBQVpPO0FBSUk7QUFHQTtBQUhULDRCQUFTLFdBRFQsY0FIVyxpQkFJRjtBQUdULDRCQUFTLFlBRFQsZUFOVyxpQkFPRjtBQVBFLGtCQUFOLCtDQURQLDZCQUNhO0FBQ1gsY0FEVyxpQkFDSixVQUFTO0FBRFgsNEJBQU07IiwKICAibmFtZXMiOiBbXQp9Cg==
