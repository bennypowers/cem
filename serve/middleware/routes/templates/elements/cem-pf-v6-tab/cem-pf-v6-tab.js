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

// elements/cem-pf-v6-tab/cem-pf-v6-tab.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-tab/cem-pf-v6-tab.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n  padding: var(--pf-t--global--spacer--md);\\n  height: 100%;\\n  box-sizing: border-box;\\n}\\n"'));
var cem_pf_v6_tab_default = s;

// elements/cem-pf-v6-tab/cem-pf-v6-tab.ts
var _title_dec, _a, _PfV6Tab_decorators, _init, _title;
_PfV6Tab_decorators = [customElement("cem-pf-v6-tab")];
var PfV6Tab = class extends (_a = LitElement, _title_dec = [property({ reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _title, __runInitializers(_init, 8, this, "")), __runInitializers(_init, 11, this);
  }
  connectedCallback() {
    super.connectedCallback();
    this.dispatchEvent(new Event("cem-pf-v6-tab-connected", { bubbles: true }));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.dispatchEvent(new Event("cem-pf-v6-tab-disconnected", { bubbles: true }));
  }
  updated(changed) {
    if (changed.has("title")) {
      this.dispatchEvent(new Event("cem-pf-v6-tab-title-changed", { bubbles: true }));
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
__publicField(PfV6Tab, "styles", cem_pf_v6_tab_default);
__runInitializers(_init, 1, PfV6Tab);
export {
  PfV6Tab
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXRhYi9jZW0tcGYtdjYtdGFiLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LXRhYi9jZW0tcGYtdjYtdGFiLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtdGFiLmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBUYWIgcGFuZWwgY29udGVudFxuICpcbiAqIFVzZWQgYXMgYSBjaGlsZCBvZiBgY2VtLXBmLXY2LXRhYnNgIHRvIGRlZmluZSB0YWIgcGFuZWxzLlxuICogVGhlIGB0aXRsZWAgYXR0cmlidXRlIHByb3ZpZGVzIHRoZSB0YWIgYnV0dG9uIGxhYmVsLlxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciB0YWIgcGFuZWwgY29udGVudFxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXBmLXY2LXRhYicpXG5leHBvcnQgY2xhc3MgUGZWNlRhYiBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgb3ZlcnJpZGUgYWNjZXNzb3IgdGl0bGUgPSAnJztcblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQoJ2NlbS1wZi12Ni10YWItY29ubmVjdGVkJywgeyBidWJibGVzOiB0cnVlIH0pKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnY2VtLXBmLXY2LXRhYi1kaXNjb25uZWN0ZWQnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygndGl0bGUnKSkge1xuICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCgnY2VtLXBmLXY2LXRhYi10aXRsZS1jaGFuZ2VkJywgeyBidWJibGVzOiB0cnVlIH0pKTtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgPHNsb3Q+PC9zbG90PmA7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXBmLXY2LXRhYic6IFBmVjZUYWI7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCk7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbiAgYm94LXNpemluZzogYm9yZGVyLWJveDtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxrSUFBb0ksQ0FBQztBQUFFLElBQU8sd0JBQVE7OztBREEzTTtBQWNBLHVCQUFDLGNBQWMsZUFBZTtBQUN2QixJQUFNLFVBQU4sZUFBc0IsaUJBRzNCLGNBQUMsU0FBUyxFQUFFLFNBQVMsS0FBSyxDQUFDLElBSEEsSUFBVztBQUFBLEVBQWpDO0FBQUE7QUFJTCx1QkFBa0IsUUFBUSxrQkFBMUIsZ0JBQTBCLE1BQTFCO0FBQUE7QUFBQSxFQUVBLG9CQUFvQjtBQUNsQixVQUFNLGtCQUFrQjtBQUN4QixTQUFLLGNBQWMsSUFBSSxNQUFNLDJCQUEyQixFQUFFLFNBQVMsS0FBSyxDQUFDLENBQUM7QUFBQSxFQUM1RTtBQUFBLEVBRUEsdUJBQXVCO0FBQ3JCLFVBQU0scUJBQXFCO0FBQzNCLFNBQUssY0FBYyxJQUFJLE1BQU0sOEJBQThCLEVBQUUsU0FBUyxLQUFLLENBQUMsQ0FBQztBQUFBLEVBQy9FO0FBQUEsRUFFQSxRQUFRLFNBQStCO0FBQ3JDLFFBQUksUUFBUSxJQUFJLE9BQU8sR0FBRztBQUN4QixXQUFLLGNBQWMsSUFBSSxNQUFNLCtCQUErQixFQUFFLFNBQVMsS0FBSyxDQUFDLENBQUM7QUFBQSxJQUNoRjtBQUFBLEVBQ0Y7QUFBQSxFQUVBLFNBQVM7QUFDUCxXQUFPO0FBQUEsRUFDVDtBQUNGO0FBekJPO0FBSWE7QUFBbEIsNEJBQWtCLFNBRGxCLFlBSFcsU0FJTztBQUpQLFVBQU4sdUNBRFAscUJBQ2E7QUFDWCxjQURXLFNBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
