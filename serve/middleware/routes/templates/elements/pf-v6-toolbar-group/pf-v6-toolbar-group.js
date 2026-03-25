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

// elements/pf-v6-toolbar-group/pf-v6-toolbar-group.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-toolbar-group/pf-v6-toolbar-group.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --pf-v6-hidden-visible--visible--Display: flex;\\n  --pf-v6-hidden-visible--hidden--Display: none;\\n  --pf-v6-hidden-visible--Display: var(--pf-v6-hidden-visible--visible--Display);\\n\\n  display: var(--pf-v6-hidden-visible--Display);\\n  row-gap: var(--pf-v6-c-toolbar__group--RowGap);\\n  column-gap: var(--pf-v6-c-toolbar__group--ColumnGap);\\n  align-items: baseline;\\n}\\n\\n:host([variant=\\"filter-group\\"]) {\\n  column-gap: var(--pf-v6-c-toolbar__group--m-filter-group--ColumnGap);\\n}\\n\\n:host([variant=\\"label-group\\"]) {\\n  flex: 1;\\n  flex-wrap: wrap;\\n  column-gap: var(--pf-v6-c-toolbar__group--m-label-group--ColumnGap);\\n}\\n\\n:host([variant=\\"action-group\\"]) {\\n  column-gap: var(--pf-v6-c-toolbar__group--m-action-group--ColumnGap);\\n}\\n\\n:host([variant=\\"action-group-plain\\"]) {\\n  column-gap: var(--pf-v6-c-toolbar__group--m-action-group-plain--ColumnGap);\\n}\\n\\n:host([variant=\\"action-group-inline\\"]) {\\n  flex-wrap: wrap;\\n  column-gap: var(--pf-v6-c-toolbar__group--m-action-group-inline--ColumnGap);\\n}\\n\\n:host([variant=\\"overflow-container\\"]) {\\n  flex: 1;\\n  min-width: var(--pf-v6-c-toolbar__group--m-overflow-container--MinWidth);\\n  overflow: auto;\\n}\\n"'));
var pf_v6_toolbar_group_default = s;

// elements/pf-v6-toolbar-group/pf-v6-toolbar-group.ts
var _variant_dec, _a, _PfV6ToolbarGroup_decorators, _init, _variant;
_PfV6ToolbarGroup_decorators = [customElement("pf-v6-toolbar-group")];
var PfV6ToolbarGroup = class extends (_a = LitElement, _variant_dec = [property({ reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _variant, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
  }
  render() {
    return html`<slot></slot>`;
  }
};
_init = __decoratorStart(_a);
_variant = new WeakMap();
__decorateElement(_init, 4, "variant", _variant_dec, PfV6ToolbarGroup, _variant);
PfV6ToolbarGroup = __decorateElement(_init, 0, "PfV6ToolbarGroup", _PfV6ToolbarGroup_decorators, PfV6ToolbarGroup);
__publicField(PfV6ToolbarGroup, "styles", pf_v6_toolbar_group_default);
__runInitializers(_init, 1, PfV6ToolbarGroup);
export {
  PfV6ToolbarGroup
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtdG9vbGJhci1ncm91cC9wZi12Ni10b29sYmFyLWdyb3VwLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvcGYtdjYtdG9vbGJhci1ncm91cC9wZi12Ni10b29sYmFyLWdyb3VwLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9wZi12Ni10b29sYmFyLWdyb3VwLmNzcyc7XG5cbnR5cGUgVG9vbGJhckdyb3VwVmFyaWFudCA9XG4gIHwgJ2ZpbHRlci1ncm91cCdcbiAgfCAnbGFiZWwtZ3JvdXAnXG4gIHwgJ2FjdGlvbi1ncm91cCdcbiAgfCAnYWN0aW9uLWdyb3VwLXBsYWluJ1xuICB8ICdhY3Rpb24tZ3JvdXAtaW5saW5lJ1xuICB8ICdvdmVyZmxvdy1jb250YWluZXInO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgVG9vbGJhciBHcm91cFxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciB0b29sYmFyIGl0ZW1zXG4gKi9cbkBjdXN0b21FbGVtZW50KCdwZi12Ni10b29sYmFyLWdyb3VwJylcbmV4cG9ydCBjbGFzcyBQZlY2VG9vbGJhckdyb3VwIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciB2YXJpYW50PzogVG9vbGJhckdyb3VwVmFyaWFudDtcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgPHNsb3Q+PC9zbG90PmA7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAncGYtdjYtdG9vbGJhci1ncm91cCc6IFBmVjZUb29sYmFyR3JvdXA7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcblxcXFxuICAtLXBmLXY2LWhpZGRlbi12aXNpYmxlLS12aXNpYmxlLS1EaXNwbGF5OiBmbGV4O1xcXFxuICAtLXBmLXY2LWhpZGRlbi12aXNpYmxlLS1oaWRkZW4tLURpc3BsYXk6IG5vbmU7XFxcXG4gIC0tcGYtdjYtaGlkZGVuLXZpc2libGUtLURpc3BsYXk6IHZhcigtLXBmLXY2LWhpZGRlbi12aXNpYmxlLS12aXNpYmxlLS1EaXNwbGF5KTtcXFxcblxcXFxuICBkaXNwbGF5OiB2YXIoLS1wZi12Ni1oaWRkZW4tdmlzaWJsZS0tRGlzcGxheSk7XFxcXG4gIHJvdy1nYXA6IHZhcigtLXBmLXY2LWMtdG9vbGJhcl9fZ3JvdXAtLVJvd0dhcCk7XFxcXG4gIGNvbHVtbi1nYXA6IHZhcigtLXBmLXY2LWMtdG9vbGJhcl9fZ3JvdXAtLUNvbHVtbkdhcCk7XFxcXG4gIGFsaWduLWl0ZW1zOiBiYXNlbGluZTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3ZhcmlhbnQ9XFxcXFxcXCJmaWx0ZXItZ3JvdXBcXFxcXFxcIl0pIHtcXFxcbiAgY29sdW1uLWdhcDogdmFyKC0tcGYtdjYtYy10b29sYmFyX19ncm91cC0tbS1maWx0ZXItZ3JvdXAtLUNvbHVtbkdhcCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KFt2YXJpYW50PVxcXFxcXFwibGFiZWwtZ3JvdXBcXFxcXFxcIl0pIHtcXFxcbiAgZmxleDogMTtcXFxcbiAgZmxleC13cmFwOiB3cmFwO1xcXFxuICBjb2x1bW4tZ2FwOiB2YXIoLS1wZi12Ni1jLXRvb2xiYXJfX2dyb3VwLS1tLWxhYmVsLWdyb3VwLS1Db2x1bW5HYXApO1xcXFxufVxcXFxuXFxcXG46aG9zdChbdmFyaWFudD1cXFxcXFxcImFjdGlvbi1ncm91cFxcXFxcXFwiXSkge1xcXFxuICBjb2x1bW4tZ2FwOiB2YXIoLS1wZi12Ni1jLXRvb2xiYXJfX2dyb3VwLS1tLWFjdGlvbi1ncm91cC0tQ29sdW1uR2FwKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3ZhcmlhbnQ9XFxcXFxcXCJhY3Rpb24tZ3JvdXAtcGxhaW5cXFxcXFxcIl0pIHtcXFxcbiAgY29sdW1uLWdhcDogdmFyKC0tcGYtdjYtYy10b29sYmFyX19ncm91cC0tbS1hY3Rpb24tZ3JvdXAtcGxhaW4tLUNvbHVtbkdhcCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KFt2YXJpYW50PVxcXFxcXFwiYWN0aW9uLWdyb3VwLWlubGluZVxcXFxcXFwiXSkge1xcXFxuICBmbGV4LXdyYXA6IHdyYXA7XFxcXG4gIGNvbHVtbi1nYXA6IHZhcigtLXBmLXY2LWMtdG9vbGJhcl9fZ3JvdXAtLW0tYWN0aW9uLWdyb3VwLWlubGluZS0tQ29sdW1uR2FwKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3ZhcmlhbnQ9XFxcXFxcXCJvdmVyZmxvdy1jb250YWluZXJcXFxcXFxcIl0pIHtcXFxcbiAgZmxleDogMTtcXFxcbiAgbWluLXdpZHRoOiB2YXIoLS1wZi12Ni1jLXRvb2xiYXJfX2dyb3VwLS1tLW92ZXJmbG93LWNvbnRhaW5lci0tTWluV2lkdGgpO1xcXFxuICBvdmVyZmxvdzogYXV0bztcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSwrdENBQTZ1QyxDQUFDO0FBQUUsSUFBTyw4QkFBUTs7O0FEQXB6QztBQW1CQSxnQ0FBQyxjQUFjLHFCQUFxQjtBQUM3QixJQUFNLG1CQUFOLGVBQStCLGlCQUdwQyxnQkFBQyxTQUFTLEVBQUUsU0FBUyxLQUFLLENBQUMsSUFIUyxJQUFXO0FBQUEsRUFBMUM7QUFBQTtBQUlMLHVCQUFTLFVBQVQ7QUFBQTtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQSxFQUNUO0FBQ0Y7QUFUTztBQUlJO0FBQVQsNEJBQVMsV0FEVCxjQUhXLGtCQUlGO0FBSkUsbUJBQU4sZ0RBRFAsOEJBQ2E7QUFDWCxjQURXLGtCQUNKLFVBQVM7QUFEWCw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
