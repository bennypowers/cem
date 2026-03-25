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

// elements/pf-v6-nav-item/pf-v6-nav-item.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-nav-item/pf-v6-nav-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  display: grid;\\n}\\n\\n:host([expanded]:last-child) ::slotted(pf-v6-nav-group) {\\n  margin-block-end: calc(var(--pf-v6-c-nav__subnav--PaddingBlockEnd, var(--pf-t--global--spacer--sm)) * -1);\\n}\\n"'));
var pf_v6_nav_item_default = s;

// elements/pf-v6-nav-item/pf-v6-nav-item.ts
var _expanded_dec, _a, _PfV6NavItem_decorators, _init, _expanded, _internals, _handleToggle, _PfV6NavItem_instances, updateExpandedState_fn;
_PfV6NavItem_decorators = [customElement("pf-v6-nav-item")];
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
    const navLink = this.querySelector(":scope > pf-v6-nav-link[expandable]");
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
  this.querySelector(":scope > pf-v6-nav-group")?.toggleAttribute("hidden", !isExpanded);
  this.querySelector(":scope > pf-v6-nav-link[expandable]")?.toggleAttribute("expanded", isExpanded);
};
__decorateElement(_init, 4, "expanded", _expanded_dec, PfV6NavItem, _expanded);
PfV6NavItem = __decorateElement(_init, 0, "PfV6NavItem", _PfV6NavItem_decorators, PfV6NavItem);
__publicField(PfV6NavItem, "styles", pf_v6_nav_item_default);
__runInitializers(_init, 1, PfV6NavItem);
export {
  PfV6NavItem
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtbmF2LWl0ZW0vcGYtdjYtbmF2LWl0ZW0udHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9wZi12Ni1uYXYtaXRlbS9wZi12Ni1uYXYtaXRlbS5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtbmF2LWl0ZW0uY3NzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IE5hdmlnYXRpb24gSXRlbVxuICpcbiAqIENvbnRhaW5lciBmb3IgbmF2aWdhdGlvbiBsaW5rcyB3aXRoIG9wdGlvbmFsIGV4cGFuZGFibGUgZ3JvdXBzLlxuICogVGhlIGNoaWxkIHBmLXY2LW5hdi1saW5rIGhhbmRsZXMgdGhlIGN1cnJlbnQgc3RhdGUgaW5kZXBlbmRlbnRseS5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgbmF2LWxpbmsgYW5kIG9wdGlvbmFsIG5hdi1ncm91cFxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtbmF2LWl0ZW0nKVxuZXhwb3J0IGNsYXNzIFBmVjZOYXZJdGVtIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBleHBhbmRlZCA9IGZhbHNlO1xuXG4gICNpbnRlcm5hbHMgPSB0aGlzLmF0dGFjaEludGVybmFscygpO1xuXG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCk7XG4gICAgdGhpcy4jaW50ZXJuYWxzLnJvbGUgPSAnbGlzdGl0ZW0nO1xuICB9XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ3BmLW5hdi10b2dnbGUnLCB0aGlzLiNoYW5kbGVUb2dnbGUgYXMgRXZlbnRMaXN0ZW5lcik7XG4gICAgLy8gSHlkcmF0ZTogc3luYyBuYXYtaXRlbSBleHBhbmRlZCBzdGF0ZSBmcm9tIGRpcmVjdCBjaGlsZCBuYXYtbGluayAoU1NSIHNldHMgaXQgdGhlcmUpXG4gICAgY29uc3QgbmF2TGluayA9IHRoaXMucXVlcnlTZWxlY3RvcignOnNjb3BlID4gcGYtdjYtbmF2LWxpbmtbZXhwYW5kYWJsZV0nKTtcbiAgICBpZiAobmF2TGluaz8uaGFzQXR0cmlidXRlKCdleHBhbmRlZCcpICYmICF0aGlzLmhhc0F0dHJpYnV0ZSgnZXhwYW5kZWQnKSkge1xuICAgICAgdGhpcy5leHBhbmRlZCA9IHRydWU7XG4gICAgfVxuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ3BmLW5hdi10b2dnbGUnLCB0aGlzLiNoYW5kbGVUb2dnbGUgYXMgRXZlbnRMaXN0ZW5lcik7XG4gIH1cblxuICB1cGRhdGVkKGNoYW5nZWQ6IE1hcDxzdHJpbmcsIHVua25vd24+KSB7XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdleHBhbmRlZCcpKSB7XG4gICAgICB0aGlzLiN1cGRhdGVFeHBhbmRlZFN0YXRlKCk7XG4gICAgfVxuICB9XG5cbiAgI2hhbmRsZVRvZ2dsZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBldmVudC5zdG9wUHJvcGFnYXRpb24oKTtcbiAgICB0aGlzLmV4cGFuZGVkID0gIXRoaXMuZXhwYW5kZWQ7XG4gIH07XG5cbiAgI3VwZGF0ZUV4cGFuZGVkU3RhdGUoKSB7XG4gICAgY29uc3QgaXNFeHBhbmRlZCA9IHRoaXMuZXhwYW5kZWQ7XG5cbiAgICAvLyBVcGRhdGUgZGlyZWN0IGNoaWxkIG5hdi1ncm91cCBvbmx5IChub3QgbmVzdGVkIG9uZXMpXG4gICAgdGhpcy5xdWVyeVNlbGVjdG9yKCc6c2NvcGUgPiBwZi12Ni1uYXYtZ3JvdXAnKVxuICAgICAgPy50b2dnbGVBdHRyaWJ1dGUoJ2hpZGRlbicsICFpc0V4cGFuZGVkKTtcblxuICAgIC8vIFVwZGF0ZSBkaXJlY3QgY2hpbGQgbmF2LWxpbmsgZXhwYW5kZWQgc3RhdGUgb25seVxuICAgIC8vIFRoZSBuYXYtbGluayB3aWxsIGZvcndhcmQgdGhpcyB0byBhcmlhLWV4cGFuZGVkIG9uIGl0cyBpbnRlcm5hbCBlbGVtZW50XG4gICAgdGhpcy5xdWVyeVNlbGVjdG9yKCc6c2NvcGUgPiBwZi12Ni1uYXYtbGlua1tleHBhbmRhYmxlXScpXG4gICAgICA/LnRvZ2dsZUF0dHJpYnV0ZSgnZXhwYW5kZWQnLCBpc0V4cGFuZGVkKTtcbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGA8c2xvdD48L3Nsb3Q+YDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdwZi12Ni1uYXYtaXRlbSc6IFBmVjZOYXZJdGVtO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG5cXFxcbiAgZGlzcGxheTogZ3JpZDtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2V4cGFuZGVkXTpsYXN0LWNoaWxkKSA6OnNsb3R0ZWQocGYtdjYtbmF2LWdyb3VwKSB7XFxcXG4gIG1hcmdpbi1ibG9jay1lbmQ6IGNhbGModmFyKC0tcGYtdjYtYy1uYXZfX3N1Ym5hdi0tUGFkZGluZ0Jsb2NrRW5kLCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pKSAqIC0xKTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSwwTkFBNE4sQ0FBQztBQUFFLElBQU8seUJBQVE7OztBREFuUztBQWNBLDJCQUFDLGNBQWMsZ0JBQWdCO0FBQ3hCLElBQU0sY0FBTixlQUEwQixpQkFHL0IsaUJBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUhYLElBQVc7QUFBQSxFQVExQyxjQUFjO0FBQ1osVUFBTTtBQVRIO0FBSUwsdUJBQVMsV0FBVyxrQkFBcEIsZ0JBQW9CLFNBQXBCO0FBRUEsbUNBQWEsS0FBSyxnQkFBZ0I7QUE0QmxDLHNDQUFnQixDQUFDLFVBQWlCO0FBQ2hDLFlBQU0sZ0JBQWdCO0FBQ3RCLFdBQUssV0FBVyxDQUFDLEtBQUs7QUFBQSxJQUN4QjtBQTNCRSx1QkFBSyxZQUFXLE9BQU87QUFBQSxFQUN6QjtBQUFBLEVBRUEsb0JBQW9CO0FBQ2xCLFVBQU0sa0JBQWtCO0FBQ3hCLFNBQUssaUJBQWlCLGlCQUFpQixtQkFBSyxjQUE4QjtBQUUxRSxVQUFNLFVBQVUsS0FBSyxjQUFjLHFDQUFxQztBQUN4RSxRQUFJLFNBQVMsYUFBYSxVQUFVLEtBQUssQ0FBQyxLQUFLLGFBQWEsVUFBVSxHQUFHO0FBQ3ZFLFdBQUssV0FBVztBQUFBLElBQ2xCO0FBQUEsRUFDRjtBQUFBLEVBRUEsdUJBQXVCO0FBQ3JCLFVBQU0scUJBQXFCO0FBQzNCLFNBQUssb0JBQW9CLGlCQUFpQixtQkFBSyxjQUE4QjtBQUFBLEVBQy9FO0FBQUEsRUFFQSxRQUFRLFNBQStCO0FBQ3JDLFFBQUksUUFBUSxJQUFJLFVBQVUsR0FBRztBQUMzQiw0QkFBSyxnREFBTDtBQUFBLElBQ0Y7QUFBQSxFQUNGO0FBQUEsRUFvQkEsU0FBUztBQUNQLFdBQU87QUFBQSxFQUNUO0FBQ0Y7QUF2RE87QUFJSTtBQUVUO0FBNEJBO0FBbENLO0FBdUNMLHlCQUFvQixXQUFHO0FBQ3JCLFFBQU0sYUFBYSxLQUFLO0FBR3hCLE9BQUssY0FBYywwQkFBMEIsR0FDekMsZ0JBQWdCLFVBQVUsQ0FBQyxVQUFVO0FBSXpDLE9BQUssY0FBYyxxQ0FBcUMsR0FDcEQsZ0JBQWdCLFlBQVksVUFBVTtBQUM1QztBQTlDQSw0QkFBUyxZQURULGVBSFcsYUFJRjtBQUpFLGNBQU4sMkNBRFAseUJBQ2E7QUFDWCxjQURXLGFBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
