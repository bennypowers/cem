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

// elements/cem-pf-v6-page-sidebar/cem-pf-v6-page-sidebar.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-page-sidebar/cem-pf-v6-page-sidebar.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  z-index: var(--cem-pf-v6-c-page__sidebar--ZIndex);\\n  display: flex;\\n  flex-direction: column;\\n  grid-area: sidebar;\\n  grid-row-start: 2;\\n  grid-column-start: 1;\\n  width: var(--cem-pf-v6-c-page__sidebar--Width);\\n  padding-block-start: 0;\\n  padding-block-end: var(--cem-pf-v6-c-page__sidebar--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-page__sidebar--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-page__sidebar--PaddingInlineEnd);\\n  margin-inline-end: var(--cem-pf-v6-c-page__sidebar--MarginInlineEnd);\\n  overflow-x: hidden;\\n  overflow-y: auto;\\n  -webkit-overflow-scrolling: touch;\\n  background-color: var(--cem-pf-v6-c-page__sidebar--BackgroundColor);\\n  border-inline-end: var(--cem-pf-v6-c-page__sidebar--BorderInlineEndWidth) solid var(--cem-pf-v6-c-page__sidebar--BorderInlineEndColor);\\n  opacity: var(--cem-pf-v6-c-page__sidebar--Opacity);\\n  transition: var(--cem-pf-v6-c-page__sidebar--TransitionProperty) var(--cem-pf-v6-c-page__sidebar--TransitionDuration) var(--cem-pf-v6-c-page__sidebar--TransitionTimingFunction);\\n  transform: translateX(var(--cem-pf-v6-c-page__sidebar--TranslateX)) translateZ(var(--cem-pf-v6-c-page__sidebar--TranslateZ));\\n}\\n\\n#body {\\n  padding-inline-start: var(--cem-pf-v6-c-page__sidebar-body--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-page__sidebar-body--PaddingInlineEnd);\\n  flex-grow: 1;\\n}\\n\\n:host([expanded]) {\\n  --cem-pf-v6-c-page__sidebar--Opacity: var(--cem-pf-v6-c-page__sidebar--m-expanded--Opacity);\\n  --cem-pf-v6-c-page__sidebar--TranslateX: var(--cem-pf-v6-c-page__sidebar--m-expanded--TranslateX);\\n  box-shadow: var(--cem-pf-v6-c-page__sidebar--BoxShadow);\\n\\n  @media screen and (min-width: 75rem) {\\n    --cem-pf-v6-c-page__sidebar--BoxShadow: 0;\\n  }\\n}\\n\\n:host([collapsed]) {\\n  --cem-pf-v6-c-page__sidebar--Opacity: 0;\\n  --cem-pf-v6-c-page__sidebar--TranslateX: -100%;\\n  pointer-events: none;\\n}\\n"'));
var cem_pf_v6_page_sidebar_default = s;

// elements/cem-pf-v6-page-sidebar/cem-pf-v6-page-sidebar.ts
var _expanded_dec, _collapsed_dec, _a, _PfV6PageSidebar_decorators, _internals, _init, _collapsed, _expanded;
_PfV6PageSidebar_decorators = [customElement("cem-pf-v6-page-sidebar")];
var PfV6PageSidebar = class extends (_a = LitElement, _collapsed_dec = [property({ type: Boolean, reflect: true })], _expanded_dec = [property({ type: Boolean, reflect: true })], _a) {
  constructor() {
    super();
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _collapsed, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _expanded, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateGet(this, _internals).role = "navigation";
  }
  updated(changed) {
    if (changed.has("collapsed")) {
      this.expanded = !this.collapsed;
    } else if (changed.has("expanded")) {
      this.collapsed = !this.expanded;
    }
  }
  render() {
    return html`
      <div id="body">
        <slot></slot>
      </div>
    `;
  }
};
_init = __decoratorStart(_a);
_internals = new WeakMap();
_collapsed = new WeakMap();
_expanded = new WeakMap();
__decorateElement(_init, 4, "collapsed", _collapsed_dec, PfV6PageSidebar, _collapsed);
__decorateElement(_init, 4, "expanded", _expanded_dec, PfV6PageSidebar, _expanded);
PfV6PageSidebar = __decorateElement(_init, 0, "PfV6PageSidebar", _PfV6PageSidebar_decorators, PfV6PageSidebar);
__publicField(PfV6PageSidebar, "styles", cem_pf_v6_page_sidebar_default);
__runInitializers(_init, 1, PfV6PageSidebar);
export {
  PfV6PageSidebar
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXBhZ2Utc2lkZWJhci9jZW0tcGYtdjYtcGFnZS1zaWRlYmFyLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LXBhZ2Utc2lkZWJhci9jZW0tcGYtdjYtcGFnZS1zaWRlYmFyLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtcGFnZS1zaWRlYmFyLmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBQYWdlIFNpZGViYXJcbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3Igc2lkZWJhciBjb250ZW50ICh0eXBpY2FsbHkgY2VtLXBmLXY2LW5hdmlnYXRpb24pXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tcGYtdjYtcGFnZS1zaWRlYmFyJylcbmV4cG9ydCBjbGFzcyBQZlY2UGFnZVNpZGViYXIgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGNvbGxhcHNlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZXhwYW5kZWQgPSBmYWxzZTtcblxuICBjb25zdHJ1Y3RvcigpIHtcbiAgICBzdXBlcigpO1xuICAgIHRoaXMuI2ludGVybmFscy5yb2xlID0gJ25hdmlnYXRpb24nO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygnY29sbGFwc2VkJykpIHtcbiAgICAgIHRoaXMuZXhwYW5kZWQgPSAhdGhpcy5jb2xsYXBzZWQ7XG4gICAgfSBlbHNlIGlmIChjaGFuZ2VkLmhhcygnZXhwYW5kZWQnKSkge1xuICAgICAgdGhpcy5jb2xsYXBzZWQgPSAhdGhpcy5leHBhbmRlZDtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8ZGl2IGlkPVwiYm9keVwiPlxuICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1wZi12Ni1wYWdlLXNpZGViYXInOiBQZlY2UGFnZVNpZGViYXI7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgei1pbmRleDogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tWkluZGV4KTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgZ3JpZC1hcmVhOiBzaWRlYmFyO1xcXFxuICBncmlkLXJvdy1zdGFydDogMjtcXFxcbiAgZ3JpZC1jb2x1bW4tc3RhcnQ6IDE7XFxcXG4gIHdpZHRoOiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1XaWR0aCk7XFxcXG4gIHBhZGRpbmctYmxvY2stc3RhcnQ6IDA7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgbWFyZ2luLWlubGluZS1lbmQ6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLU1hcmdpbklubGluZUVuZCk7XFxcXG4gIG92ZXJmbG93LXg6IGhpZGRlbjtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbiAgLXdlYmtpdC1vdmVyZmxvdy1zY3JvbGxpbmc6IHRvdWNoO1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICBib3JkZXItaW5saW5lLWVuZDogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tQm9yZGVySW5saW5lRW5kV2lkdGgpIHNvbGlkIHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLUJvcmRlcklubGluZUVuZENvbG9yKTtcXFxcbiAgb3BhY2l0eTogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tT3BhY2l0eSk7XFxcXG4gIHRyYW5zaXRpb246IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLVRyYW5zaXRpb25Qcm9wZXJ0eSkgdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tVHJhbnNpdGlvbkR1cmF0aW9uKSB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1UcmFuc2l0aW9uVGltaW5nRnVuY3Rpb24pO1xcXFxuICB0cmFuc2Zvcm06IHRyYW5zbGF0ZVgodmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tVHJhbnNsYXRlWCkpIHRyYW5zbGF0ZVoodmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tVHJhbnNsYXRlWikpO1xcXFxufVxcXFxuXFxcXG4jYm9keSB7XFxcXG4gIHBhZGRpbmctaW5saW5lLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLWJvZHktLVBhZGRpbmdJbmxpbmVTdGFydCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLWVuZDogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci1ib2R5LS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgZmxleC1ncm93OiAxO1xcXFxufVxcXFxuXFxcXG46aG9zdChbZXhwYW5kZWRdKSB7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tT3BhY2l0eTogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tbS1leHBhbmRlZC0tT3BhY2l0eSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tVHJhbnNsYXRlWDogdmFyKC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tbS1leHBhbmRlZC0tVHJhbnNsYXRlWCk7XFxcXG4gIGJveC1zaGFkb3c6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLUJveFNoYWRvdyk7XFxcXG5cXFxcbiAgQG1lZGlhIHNjcmVlbiBhbmQgKG1pbi13aWR0aDogNzVyZW0pIHtcXFxcbiAgICAtLWNlbS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLUJveFNoYWRvdzogMDtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG46aG9zdChbY29sbGFwc2VkXSkge1xcXFxuICAtLWNlbS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLU9wYWNpdHk6IDA7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tVHJhbnNsYXRlWDogLTEwMCU7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLG84REFBczhELENBQUM7QUFBRSxJQUFPLGlDQUFROzs7QURBN2dFO0FBV0EsK0JBQUMsY0FBYyx3QkFBd0I7QUFDaEMsSUFBTSxrQkFBTixlQUE4QixpQkFLbkMsa0JBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxpQkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBUlAsSUFBVztBQUFBLEVBVzlDLGNBQWM7QUFDWixVQUFNO0FBVFIsbUNBQWEsS0FBSyxnQkFBZ0I7QUFHbEMsdUJBQVMsWUFBWSxrQkFBckIsZ0JBQXFCLFNBQXJCO0FBR0EsdUJBQVMsV0FBVyxrQkFBcEIsaUJBQW9CLFNBQXBCO0FBSUUsdUJBQUssWUFBVyxPQUFPO0FBQUEsRUFDekI7QUFBQSxFQUVBLFFBQVEsU0FBK0I7QUFDckMsUUFBSSxRQUFRLElBQUksV0FBVyxHQUFHO0FBQzVCLFdBQUssV0FBVyxDQUFDLEtBQUs7QUFBQSxJQUN4QixXQUFXLFFBQVEsSUFBSSxVQUFVLEdBQUc7QUFDbEMsV0FBSyxZQUFZLENBQUMsS0FBSztBQUFBLElBQ3pCO0FBQUEsRUFDRjtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBS1Q7QUFDRjtBQS9CTztBQUdMO0FBR1M7QUFHQTtBQUhULDRCQUFTLGFBRFQsZ0JBTFcsaUJBTUY7QUFHVCw0QkFBUyxZQURULGVBUlcsaUJBU0Y7QUFURSxrQkFBTiwrQ0FEUCw2QkFDYTtBQUNYLGNBRFcsaUJBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
