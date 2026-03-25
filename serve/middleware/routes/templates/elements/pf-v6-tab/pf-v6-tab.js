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

// elements/pf-v6-tab/pf-v6-tab.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/pf-v6-tab/pf-v6-tab.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n  padding: var(--pf-t--global--spacer--md);\\n  height: 100%;\\n  box-sizing: border-box;\\n}\\n"'));
var pf_v6_tab_default = s;

// elements/pf-v6-tab/pf-v6-tab.ts
var _title_dec, _a, _PfV6Tab_decorators, _init, _title;
_PfV6Tab_decorators = [customElement("pf-v6-tab")];
var PfV6Tab = class extends (_a = LitElement, _title_dec = [property({ reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _title, __runInitializers(_init, 8, this, "")), __runInitializers(_init, 11, this);
  }
  connectedCallback() {
    super.connectedCallback();
    this.dispatchEvent(new Event("pf-v6-tab-connected", { bubbles: true }));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.dispatchEvent(new Event("pf-v6-tab-disconnected", { bubbles: true }));
  }
  updated(changed) {
    if (changed.has("title")) {
      this.dispatchEvent(new Event("pf-v6-tab-title-changed", { bubbles: true }));
    }
  }
  render() {
    return html`<slot></slot>`;
  }
};
_init = __decoratorStart(_a);
_title = new WeakMap();
__decorateElement(_init, 4, "title", _title_dec, PfV6Tab, _title);
PfV6Tab = __decorateElement(_init, 0, "PfV6Tab", _PfV6Tab_decorators, PfV6Tab);
__publicField(PfV6Tab, "styles", pf_v6_tab_default);
__runInitializers(_init, 1, PfV6Tab);
export {
  PfV6Tab
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtdGFiL3BmLXY2LXRhYi50cyIsICJsaXQtY3NzOmVsZW1lbnRzL3BmLXY2LXRhYi9wZi12Ni10YWIuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL3BmLXY2LXRhYi5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgVGFiIHBhbmVsIGNvbnRlbnRcbiAqXG4gKiBVc2VkIGFzIGEgY2hpbGQgb2YgYHBmLXY2LXRhYnNgIHRvIGRlZmluZSB0YWIgcGFuZWxzLlxuICogVGhlIGB0aXRsZWAgYXR0cmlidXRlIHByb3ZpZGVzIHRoZSB0YWIgYnV0dG9uIGxhYmVsLlxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciB0YWIgcGFuZWwgY29udGVudFxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtdGFiJylcbmV4cG9ydCBjbGFzcyBQZlY2VGFiIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KVxuICBvdmVycmlkZSBhY2Nlc3NvciB0aXRsZSA9ICcnO1xuXG4gIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmNvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgncGYtdjYtdGFiLWNvbm5lY3RlZCcsIHsgYnViYmxlczogdHJ1ZSB9KSk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQoJ3BmLXY2LXRhYi1kaXNjb25uZWN0ZWQnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygndGl0bGUnKSkge1xuICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgncGYtdjYtdGFiLXRpdGxlLWNoYW5nZWQnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICAgIH1cbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGA8c2xvdD48L3Nsb3Q+YDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdwZi12Ni10YWInOiBQZlY2VGFiO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBwYWRkaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG4gIGJveC1zaXppbmc6IGJvcmRlci1ib3g7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sa0lBQW9JLENBQUM7QUFBRSxJQUFPLG9CQUFROzs7QURBM007QUFjQSx1QkFBQyxjQUFjLFdBQVc7QUFDbkIsSUFBTSxVQUFOLGVBQXNCLGlCQUczQixjQUFDLFNBQVMsRUFBRSxTQUFTLEtBQUssQ0FBQyxJQUhBLElBQVc7QUFBQSxFQUFqQztBQUFBO0FBSUwsdUJBQWtCLFFBQVEsa0JBQTFCLGdCQUEwQixNQUExQjtBQUFBO0FBQUEsRUFFQSxvQkFBb0I7QUFDbEIsVUFBTSxrQkFBa0I7QUFDeEIsU0FBSyxjQUFjLElBQUksTUFBTSx1QkFBdUIsRUFBRSxTQUFTLEtBQUssQ0FBQyxDQUFDO0FBQUEsRUFDeEU7QUFBQSxFQUVBLHVCQUF1QjtBQUNyQixVQUFNLHFCQUFxQjtBQUMzQixTQUFLLGNBQWMsSUFBSSxNQUFNLDBCQUEwQixFQUFFLFNBQVMsS0FBSyxDQUFDLENBQUM7QUFBQSxFQUMzRTtBQUFBLEVBRUEsUUFBUSxTQUErQjtBQUNyQyxRQUFJLFFBQVEsSUFBSSxPQUFPLEdBQUc7QUFDeEIsV0FBSyxjQUFjLElBQUksTUFBTSwyQkFBMkIsRUFBRSxTQUFTLEtBQUssQ0FBQyxDQUFDO0FBQUEsSUFDNUU7QUFBQSxFQUNGO0FBQUEsRUFFQSxTQUFTO0FBQ1AsV0FBTztBQUFBLEVBQ1Q7QUFDRjtBQXpCTztBQUlhO0FBQWxCLDRCQUFrQixTQURsQixZQUhXLFNBSU87QUFKUCxVQUFOLHVDQURQLHFCQUNhO0FBQ1gsY0FEVyxTQUNKLFVBQVM7QUFEWCw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
