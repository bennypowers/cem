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

// elements/cem-pf-v6-nav-item/cem-pf-v6-nav-item.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-nav-item/cem-pf-v6-nav-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  display: grid;\\n}\\n\\n:host([expanded]:last-child) ::slotted(cem-pf-v6-nav-group) {\\n  margin-block-end: calc(var(--cem-pf-v6-c-nav__subnav--PaddingBlockEnd, var(--pf-t--global--spacer--sm)) * -1);\\n}\\n"'));
var cem_pf_v6_nav_item_default = s;

// elements/cem-pf-v6-nav-item/cem-pf-v6-nav-item.ts
var _expanded_dec, _a, _PfV6NavItem_decorators, _init, _expanded, _internals, _handleToggle, _PfV6NavItem_instances, updateExpandedState_fn;
_PfV6NavItem_decorators = [customElement("cem-pf-v6-nav-item")];
var PfV6NavItem = class extends (_a = LitElement, _expanded_dec = [property({ type: Boolean, reflect: true })], _a) {
  constructor() {
    super();
    __privateAdd(this, _PfV6NavItem_instances);
    __privateAdd(this, _expanded, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _handleToggle, (event) => {
      event.stopPropagation();
      this.expanded = !this.expanded;
    });
    __privateGet(this, _internals).role = "listitem";
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("pf-nav-toggle", __privateGet(this, _handleToggle));
    const navLink = this.querySelector(":scope > cem-pf-v6-nav-link[expandable]");
    if (navLink?.hasAttribute("expanded") && !this.hasAttribute("expanded")) {
      this.expanded = true;
    }
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("pf-nav-toggle", __privateGet(this, _handleToggle));
  }
  updated(changed) {
    if (changed.has("expanded")) {
      __privateMethod(this, _PfV6NavItem_instances, updateExpandedState_fn).call(this);
    }
  }
  render() {
    return html`<slot></slot>`;
  }
};
_init = __decoratorStart(_a);
_expanded = new WeakMap();
_internals = new WeakMap();
_handleToggle = new WeakMap();
_PfV6NavItem_instances = new WeakSet();
updateExpandedState_fn = function() {
  const isExpanded = this.expanded;
  this.querySelector(":scope > cem-pf-v6-nav-group")?.toggleAttribute("hidden", !isExpanded);
  this.querySelector(":scope > cem-pf-v6-nav-link[expandable]")?.toggleAttribute("expanded", isExpanded);
};
__decorateElement(_init, 4, "expanded", _expanded_dec, PfV6NavItem, _expanded);
PfV6NavItem = __decorateElement(_init, 0, "PfV6NavItem", _PfV6NavItem_decorators, PfV6NavItem);
__publicField(PfV6NavItem, "styles", cem_pf_v6_nav_item_default);
__runInitializers(_init, 1, PfV6NavItem);
export {
  PfV6NavItem
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LW5hdi1pdGVtL2NlbS1wZi12Ni1uYXYtaXRlbS50cyIsICJsaXQtY3NzOmVsZW1lbnRzL2NlbS1wZi12Ni1uYXYtaXRlbS9jZW0tcGYtdjYtbmF2LWl0ZW0uY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1wZi12Ni1uYXYtaXRlbS5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgTmF2aWdhdGlvbiBJdGVtXG4gKlxuICogQ29udGFpbmVyIGZvciBuYXZpZ2F0aW9uIGxpbmtzIHdpdGggb3B0aW9uYWwgZXhwYW5kYWJsZSBncm91cHMuXG4gKiBUaGUgY2hpbGQgY2VtLXBmLXY2LW5hdi1saW5rIGhhbmRsZXMgdGhlIGN1cnJlbnQgc3RhdGUgaW5kZXBlbmRlbnRseS5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgbmF2LWxpbmsgYW5kIG9wdGlvbmFsIG5hdi1ncm91cFxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXBmLXY2LW5hdi1pdGVtJylcbmV4cG9ydCBjbGFzcyBQZlY2TmF2SXRlbSBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZXhwYW5kZWQgPSBmYWxzZTtcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcblxuICBjb25zdHJ1Y3RvcigpIHtcbiAgICBzdXBlcigpO1xuICAgIHRoaXMuI2ludGVybmFscy5yb2xlID0gJ2xpc3RpdGVtJztcbiAgfVxuXG4gIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmNvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdwZi1uYXYtdG9nZ2xlJywgdGhpcy4jaGFuZGxlVG9nZ2xlIGFzIEV2ZW50TGlzdGVuZXIpO1xuICAgIC8vIEh5ZHJhdGU6IHN5bmMgbmF2LWl0ZW0gZXhwYW5kZWQgc3RhdGUgZnJvbSBkaXJlY3QgY2hpbGQgbmF2LWxpbmsgKFNTUiBzZXRzIGl0IHRoZXJlKVxuICAgIGNvbnN0IG5hdkxpbmsgPSB0aGlzLnF1ZXJ5U2VsZWN0b3IoJzpzY29wZSA+IGNlbS1wZi12Ni1uYXYtbGlua1tleHBhbmRhYmxlXScpO1xuICAgIGlmIChuYXZMaW5rPy5oYXNBdHRyaWJ1dGUoJ2V4cGFuZGVkJykgJiYgIXRoaXMuaGFzQXR0cmlidXRlKCdleHBhbmRlZCcpKSB7XG4gICAgICB0aGlzLmV4cGFuZGVkID0gdHJ1ZTtcbiAgICB9XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigncGYtbmF2LXRvZ2dsZScsIHRoaXMuI2hhbmRsZVRvZ2dsZSBhcyBFdmVudExpc3RlbmVyKTtcbiAgfVxuXG4gIHVwZGF0ZWQoY2hhbmdlZDogTWFwPHN0cmluZywgdW5rbm93bj4pIHtcbiAgICBpZiAoY2hhbmdlZC5oYXMoJ2V4cGFuZGVkJykpIHtcbiAgICAgIHRoaXMuI3VwZGF0ZUV4cGFuZGVkU3RhdGUoKTtcbiAgICB9XG4gIH1cblxuICAjaGFuZGxlVG9nZ2xlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGV2ZW50LnN0b3BQcm9wYWdhdGlvbigpO1xuICAgIHRoaXMuZXhwYW5kZWQgPSAhdGhpcy5leHBhbmRlZDtcbiAgfTtcblxuICAjdXBkYXRlRXhwYW5kZWRTdGF0ZSgpIHtcbiAgICBjb25zdCBpc0V4cGFuZGVkID0gdGhpcy5leHBhbmRlZDtcblxuICAgIC8vIFVwZGF0ZSBkaXJlY3QgY2hpbGQgbmF2LWdyb3VwIG9ubHkgKG5vdCBuZXN0ZWQgb25lcylcbiAgICB0aGlzLnF1ZXJ5U2VsZWN0b3IoJzpzY29wZSA+IGNlbS1wZi12Ni1uYXYtZ3JvdXAnKVxuICAgICAgPy50b2dnbGVBdHRyaWJ1dGUoJ2hpZGRlbicsICFpc0V4cGFuZGVkKTtcblxuICAgIC8vIFVwZGF0ZSBkaXJlY3QgY2hpbGQgbmF2LWxpbmsgZXhwYW5kZWQgc3RhdGUgb25seVxuICAgIC8vIFRoZSBuYXYtbGluayB3aWxsIGZvcndhcmQgdGhpcyB0byBhcmlhLWV4cGFuZGVkIG9uIGl0cyBpbnRlcm5hbCBlbGVtZW50XG4gICAgdGhpcy5xdWVyeVNlbGVjdG9yKCc6c2NvcGUgPiBjZW0tcGYtdjYtbmF2LWxpbmtbZXhwYW5kYWJsZV0nKVxuICAgICAgPy50b2dnbGVBdHRyaWJ1dGUoJ2V4cGFuZGVkJywgaXNFeHBhbmRlZCk7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgPHNsb3Q+PC9zbG90PmA7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXBmLXY2LW5hdi1pdGVtJzogUGZWNk5hdkl0ZW07XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcblxcXFxuICBkaXNwbGF5OiBncmlkO1xcXFxufVxcXFxuXFxcXG46aG9zdChbZXhwYW5kZWRdOmxhc3QtY2hpbGQpIDo6c2xvdHRlZChjZW0tcGYtdjYtbmF2LWdyb3VwKSB7XFxcXG4gIG1hcmdpbi1ibG9jay1lbmQ6IGNhbGModmFyKC0tY2VtLXBmLXY2LWMtbmF2X19zdWJuYXYtLVBhZGRpbmdCbG9ja0VuZCwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKSkgKiAtMSk7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sa09BQW9PLENBQUM7QUFBRSxJQUFPLDZCQUFROzs7QURBM1M7QUFjQSwyQkFBQyxjQUFjLG9CQUFvQjtBQUM1QixJQUFNLGNBQU4sZUFBMEIsaUJBRy9CLGlCQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFIWCxJQUFXO0FBQUEsRUFRMUMsY0FBYztBQUNaLFVBQU07QUFUSDtBQUlMLHVCQUFTLFdBQVcsa0JBQXBCLGdCQUFvQixTQUFwQjtBQUVBLG1DQUFhLEtBQUssZ0JBQWdCO0FBNEJsQyxzQ0FBZ0IsQ0FBQyxVQUFpQjtBQUNoQyxZQUFNLGdCQUFnQjtBQUN0QixXQUFLLFdBQVcsQ0FBQyxLQUFLO0FBQUEsSUFDeEI7QUEzQkUsdUJBQUssWUFBVyxPQUFPO0FBQUEsRUFDekI7QUFBQSxFQUVBLG9CQUFvQjtBQUNsQixVQUFNLGtCQUFrQjtBQUN4QixTQUFLLGlCQUFpQixpQkFBaUIsbUJBQUssY0FBOEI7QUFFMUUsVUFBTSxVQUFVLEtBQUssY0FBYyx5Q0FBeUM7QUFDNUUsUUFBSSxTQUFTLGFBQWEsVUFBVSxLQUFLLENBQUMsS0FBSyxhQUFhLFVBQVUsR0FBRztBQUN2RSxXQUFLLFdBQVc7QUFBQSxJQUNsQjtBQUFBLEVBQ0Y7QUFBQSxFQUVBLHVCQUF1QjtBQUNyQixVQUFNLHFCQUFxQjtBQUMzQixTQUFLLG9CQUFvQixpQkFBaUIsbUJBQUssY0FBOEI7QUFBQSxFQUMvRTtBQUFBLEVBRUEsUUFBUSxTQUErQjtBQUNyQyxRQUFJLFFBQVEsSUFBSSxVQUFVLEdBQUc7QUFDM0IsNEJBQUssZ0RBQUw7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUFBLEVBb0JBLFNBQVM7QUFDUCxXQUFPO0FBQUEsRUFDVDtBQUNGO0FBdkRPO0FBSUk7QUFFVDtBQTRCQTtBQWxDSztBQXVDTCx5QkFBb0IsV0FBRztBQUNyQixRQUFNLGFBQWEsS0FBSztBQUd4QixPQUFLLGNBQWMsOEJBQThCLEdBQzdDLGdCQUFnQixVQUFVLENBQUMsVUFBVTtBQUl6QyxPQUFLLGNBQWMseUNBQXlDLEdBQ3hELGdCQUFnQixZQUFZLFVBQVU7QUFDNUM7QUE5Q0EsNEJBQVMsWUFEVCxlQUhXLGFBSUY7QUFKRSxjQUFOLDJDQURQLHlCQUNhO0FBQ1gsY0FEVyxhQUNKLFVBQVM7QUFEWCw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
