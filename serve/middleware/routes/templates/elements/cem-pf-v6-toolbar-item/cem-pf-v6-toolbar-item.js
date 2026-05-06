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
var __publicField = (obj, key, value) => __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
var __accessCheck = (obj, member, msg) => member.has(obj) || __typeError("Cannot " + msg);
var __privateIn = (member, obj) => Object(obj) !== obj ? __typeError('Cannot use the "in" operator on this value') : member.has(obj);
var __privateGet = (obj, member, getter) => (__accessCheck(obj, member, "read from private field"), getter ? getter.call(obj) : member.get(obj));
var __privateAdd = (obj, member, value) => member.has(obj) ? __typeError("Cannot add the same private member more than once") : member instanceof WeakSet ? member.add(obj) : member.set(obj, value);
var __privateSet = (obj, member, value, setter) => (__accessCheck(obj, member, "write to private field"), setter ? setter.call(obj, value) : member.set(obj, value), value);
var __privateMethod = (obj, member, method) => (__accessCheck(obj, member, "access private method"), method);

// elements/cem-pf-v6-toolbar-item/cem-pf-v6-toolbar-item.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-toolbar-item/cem-pf-v6-toolbar-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-hidden-visible--visible--Display: flex;\\n  --cem-pf-v6-hidden-visible--hidden--Display: none;\\n  --cem-pf-v6-hidden-visible--Display: var(--cem-pf-v6-hidden-visible--visible--Display);\\n\\n  display: var(--cem-pf-v6-hidden-visible--Display);\\n  row-gap: var(--cem-pf-v6-c-toolbar__item--RowGap);\\n  column-gap: var(--cem-pf-v6-c-toolbar__item--ColumnGap);\\n  width: var(--cem-pf-v6-c-toolbar__item--Width);\\n  min-width: var(--cem-pf-v6-c-toolbar__item--MinWidth);\\n  align-items: baseline;\\n}\\n\\n:host([variant=\\"label\\"]) {\\n  font-weight: var(--cem-pf-v6-c-toolbar__item--m-label--FontWeight);\\n}\\n\\n:host([variant=\\"pagination\\"]) {\\n  margin-inline-start: auto;\\n}\\n\\n:host([variant=\\"overflow-container\\"]) {\\n  flex: 1;\\n  min-width: var(--cem-pf-v6-c-toolbar__item--m-overflow-container--MinWidth);\\n  overflow: auto;\\n}\\n\\n:host([variant=\\"expand-all\\"][expanded]) ::slotted(.cem-pf-v6-c-toolbar__expand-all-icon) {\\n  transform: rotate(var(--cem-pf-v6-c-toolbar__item--m-expand-all--m-expanded__expand-all-icon--Rotate));\\n}\\n"'));
var cem_pf_v6_toolbar_item_default = s;

// elements/cem-pf-v6-toolbar-item/cem-pf-v6-toolbar-item.ts
var _expanded_dec, _variant_dec, _a, _PfV6ToolbarItem_decorators, _init, _variant, _expanded;
_PfV6ToolbarItem_decorators = [customElement("cem-pf-v6-toolbar-item")];
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
__publicField(PfV6ToolbarItem, "styles", cem_pf_v6_toolbar_item_default);
__runInitializers(_init, 1, PfV6ToolbarItem);
export {
  PfV6ToolbarItem
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXRvb2xiYXItaXRlbS9jZW0tcGYtdjYtdG9vbGJhci1pdGVtLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LXRvb2xiYXItaXRlbS9jZW0tcGYtdjYtdG9vbGJhci1pdGVtLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtdG9vbGJhci1pdGVtLmNzcyc7XG5cbnR5cGUgVG9vbGJhckl0ZW1WYXJpYW50ID1cbiAgfCAnbGFiZWwnXG4gIHwgJ3BhZ2luYXRpb24nXG4gIHwgJ292ZXJmbG93LWNvbnRhaW5lcidcbiAgfCAnZXhwYW5kLWFsbCc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBUb29sYmFyIEl0ZW1cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgaXRlbSBjb250ZW50XG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tcGYtdjYtdG9vbGJhci1pdGVtJylcbmV4cG9ydCBjbGFzcyBQZlY2VG9vbGJhckl0ZW0gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIHZhcmlhbnQ/OiBUb29sYmFySXRlbVZhcmlhbnQ7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBleHBhbmRlZCA9IGZhbHNlO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGA8c2xvdD48L3Nsb3Q+YDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcGYtdjYtdG9vbGJhci1pdGVtJzogUGZWNlRvb2xiYXJJdGVtO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG5cXFxcbiAgLS1jZW0tcGYtdjYtaGlkZGVuLXZpc2libGUtLXZpc2libGUtLURpc3BsYXk6IGZsZXg7XFxcXG4gIC0tY2VtLXBmLXY2LWhpZGRlbi12aXNpYmxlLS1oaWRkZW4tLURpc3BsYXk6IG5vbmU7XFxcXG4gIC0tY2VtLXBmLXY2LWhpZGRlbi12aXNpYmxlLS1EaXNwbGF5OiB2YXIoLS1jZW0tcGYtdjYtaGlkZGVuLXZpc2libGUtLXZpc2libGUtLURpc3BsYXkpO1xcXFxuXFxcXG4gIGRpc3BsYXk6IHZhcigtLWNlbS1wZi12Ni1oaWRkZW4tdmlzaWJsZS0tRGlzcGxheSk7XFxcXG4gIHJvdy1nYXA6IHZhcigtLWNlbS1wZi12Ni1jLXRvb2xiYXJfX2l0ZW0tLVJvd0dhcCk7XFxcXG4gIGNvbHVtbi1nYXA6IHZhcigtLWNlbS1wZi12Ni1jLXRvb2xiYXJfX2l0ZW0tLUNvbHVtbkdhcCk7XFxcXG4gIHdpZHRoOiB2YXIoLS1jZW0tcGYtdjYtYy10b29sYmFyX19pdGVtLS1XaWR0aCk7XFxcXG4gIG1pbi13aWR0aDogdmFyKC0tY2VtLXBmLXY2LWMtdG9vbGJhcl9faXRlbS0tTWluV2lkdGgpO1xcXFxuICBhbGlnbi1pdGVtczogYmFzZWxpbmU7XFxcXG59XFxcXG5cXFxcbjpob3N0KFt2YXJpYW50PVxcXFxcXFwibGFiZWxcXFxcXFxcIl0pIHtcXFxcbiAgZm9udC13ZWlnaHQ6IHZhcigtLWNlbS1wZi12Ni1jLXRvb2xiYXJfX2l0ZW0tLW0tbGFiZWwtLUZvbnRXZWlnaHQpO1xcXFxufVxcXFxuXFxcXG46aG9zdChbdmFyaWFudD1cXFxcXFxcInBhZ2luYXRpb25cXFxcXFxcIl0pIHtcXFxcbiAgbWFyZ2luLWlubGluZS1zdGFydDogYXV0bztcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3ZhcmlhbnQ9XFxcXFxcXCJvdmVyZmxvdy1jb250YWluZXJcXFxcXFxcIl0pIHtcXFxcbiAgZmxleDogMTtcXFxcbiAgbWluLXdpZHRoOiB2YXIoLS1jZW0tcGYtdjYtYy10b29sYmFyX19pdGVtLS1tLW92ZXJmbG93LWNvbnRhaW5lci0tTWluV2lkdGgpO1xcXFxuICBvdmVyZmxvdzogYXV0bztcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3ZhcmlhbnQ9XFxcXFxcXCJleHBhbmQtYWxsXFxcXFxcXCJdW2V4cGFuZGVkXSkgOjpzbG90dGVkKC5jZW0tcGYtdjYtYy10b29sYmFyX19leHBhbmQtYWxsLWljb24pIHtcXFxcbiAgdHJhbnNmb3JtOiByb3RhdGUodmFyKC0tY2VtLXBmLXY2LWMtdG9vbGJhcl9faXRlbS0tbS1leHBhbmQtYWxsLS1tLWV4cGFuZGVkX19leHBhbmQtYWxsLWljb24tLVJvdGF0ZSkpO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLCtrQ0FBeWxDLENBQUM7QUFBRSxJQUFPLGlDQUFROzs7QURBaHFDO0FBaUJBLCtCQUFDLGNBQWMsd0JBQXdCO0FBQ2hDLElBQU0sa0JBQU4sZUFBOEIsaUJBR25DLGdCQUFDLFNBQVMsRUFBRSxTQUFTLEtBQUssQ0FBQyxJQUczQixpQkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBTlAsSUFBVztBQUFBLEVBQXpDO0FBQUE7QUFJTCx1QkFBUyxVQUFUO0FBR0EsdUJBQVMsV0FBVyxrQkFBcEIsaUJBQW9CLFNBQXBCO0FBQUE7QUFBQSxFQUVBLFNBQVM7QUFDUCxXQUFPO0FBQUEsRUFDVDtBQUNGO0FBWk87QUFJSTtBQUdBO0FBSFQsNEJBQVMsV0FEVCxjQUhXLGlCQUlGO0FBR1QsNEJBQVMsWUFEVCxlQU5XLGlCQU9GO0FBUEUsa0JBQU4sK0NBRFAsNkJBQ2E7QUFDWCxjQURXLGlCQUNKLFVBQVM7QUFEWCw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
