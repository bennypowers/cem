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

// elements/pf-v6-page-sidebar/pf-v6-page-sidebar.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-page-sidebar/pf-v6-page-sidebar.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  z-index: var(--pf-v6-c-page__sidebar--ZIndex);\\n  display: flex;\\n  flex-direction: column;\\n  grid-area: sidebar;\\n  grid-row-start: 2;\\n  grid-column-start: 1;\\n  width: var(--pf-v6-c-page__sidebar--Width);\\n  padding-block-start: 0;\\n  padding-block-end: var(--pf-v6-c-page__sidebar--PaddingBlockEnd);\\n  padding-inline-start: var(--pf-v6-c-page__sidebar--PaddingInlineStart);\\n  padding-inline-end: var(--pf-v6-c-page__sidebar--PaddingInlineEnd);\\n  margin-inline-end: var(--pf-v6-c-page__sidebar--MarginInlineEnd);\\n  overflow-x: hidden;\\n  overflow-y: auto;\\n  -webkit-overflow-scrolling: touch;\\n  background-color: var(--pf-v6-c-page__sidebar--BackgroundColor);\\n  border-inline-end: var(--pf-v6-c-page__sidebar--BorderInlineEndWidth) solid var(--pf-v6-c-page__sidebar--BorderInlineEndColor);\\n  opacity: var(--pf-v6-c-page__sidebar--Opacity);\\n  transition: var(--pf-v6-c-page__sidebar--TransitionProperty) var(--pf-v6-c-page__sidebar--TransitionDuration) var(--pf-v6-c-page__sidebar--TransitionTimingFunction);\\n  transform: translateX(var(--pf-v6-c-page__sidebar--TranslateX)) translateZ(var(--pf-v6-c-page__sidebar--TranslateZ));\\n}\\n\\n#body {\\n  padding-inline-start: var(--pf-v6-c-page__sidebar-body--PaddingInlineStart);\\n  padding-inline-end: var(--pf-v6-c-page__sidebar-body--PaddingInlineEnd);\\n  flex-grow: 1;\\n}\\n\\n:host([expanded]) {\\n  --pf-v6-c-page__sidebar--Opacity: var(--pf-v6-c-page__sidebar--m-expanded--Opacity);\\n  --pf-v6-c-page__sidebar--TranslateX: var(--pf-v6-c-page__sidebar--m-expanded--TranslateX);\\n  box-shadow: var(--pf-v6-c-page__sidebar--BoxShadow);\\n\\n  @media screen and (min-width: 75rem) {\\n    --pf-v6-c-page__sidebar--BoxShadow: 0;\\n  }\\n}\\n\\n:host([collapsed]) {\\n  --pf-v6-c-page__sidebar--Opacity: 0;\\n  --pf-v6-c-page__sidebar--TranslateX: -100%;\\n  pointer-events: none;\\n}\\n"'));
var pf_v6_page_sidebar_default = s;

// elements/pf-v6-page-sidebar/pf-v6-page-sidebar.ts
var _expanded_dec, _collapsed_dec, _a, _PfV6PageSidebar_decorators, _internals, _init, _collapsed, _expanded;
_PfV6PageSidebar_decorators = [customElement("pf-v6-page-sidebar")];
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
__publicField(PfV6PageSidebar, "styles", pf_v6_page_sidebar_default);
__runInitializers(_init, 1, PfV6PageSidebar);
export {
  PfV6PageSidebar
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtcGFnZS1zaWRlYmFyL3BmLXY2LXBhZ2Utc2lkZWJhci50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL3BmLXY2LXBhZ2Utc2lkZWJhci9wZi12Ni1wYWdlLXNpZGViYXIuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL3BmLXY2LXBhZ2Utc2lkZWJhci5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgUGFnZSBTaWRlYmFyXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIHNpZGViYXIgY29udGVudCAodHlwaWNhbGx5IHBmLXY2LW5hdmlnYXRpb24pXG4gKi9cbkBjdXN0b21FbGVtZW50KCdwZi12Ni1wYWdlLXNpZGViYXInKVxuZXhwb3J0IGNsYXNzIFBmVjZQYWdlU2lkZWJhciBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gICNpbnRlcm5hbHMgPSB0aGlzLmF0dGFjaEludGVybmFscygpO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgY29sbGFwc2VkID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBleHBhbmRlZCA9IGZhbHNlO1xuXG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCk7XG4gICAgdGhpcy4jaW50ZXJuYWxzLnJvbGUgPSAnbmF2aWdhdGlvbic7XG4gIH1cblxuICB1cGRhdGVkKGNoYW5nZWQ6IE1hcDxzdHJpbmcsIHVua25vd24+KSB7XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdjb2xsYXBzZWQnKSkge1xuICAgICAgdGhpcy5leHBhbmRlZCA9ICF0aGlzLmNvbGxhcHNlZDtcbiAgICB9IGVsc2UgaWYgKGNoYW5nZWQuaGFzKCdleHBhbmRlZCcpKSB7XG4gICAgICB0aGlzLmNvbGxhcHNlZCA9ICF0aGlzLmV4cGFuZGVkO1xuICAgIH1cbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJib2R5XCI+XG4gICAgICAgIDxzbG90Pjwvc2xvdD5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAncGYtdjYtcGFnZS1zaWRlYmFyJzogUGZWNlBhZ2VTaWRlYmFyO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIHotaW5kZXg6IHZhcigtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tWkluZGV4KTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgZ3JpZC1hcmVhOiBzaWRlYmFyO1xcXFxuICBncmlkLXJvdy1zdGFydDogMjtcXFxcbiAgZ3JpZC1jb2x1bW4tc3RhcnQ6IDE7XFxcXG4gIHdpZHRoOiB2YXIoLS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLVdpZHRoKTtcXFxcbiAgcGFkZGluZy1ibG9jay1zdGFydDogMDtcXFxcbiAgcGFkZGluZy1ibG9jay1lbmQ6IHZhcigtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tUGFkZGluZ0Jsb2NrRW5kKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLVBhZGRpbmdJbmxpbmVFbmQpO1xcXFxuICBtYXJnaW4taW5saW5lLWVuZDogdmFyKC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1NYXJnaW5JbmxpbmVFbmQpO1xcXFxuICBvdmVyZmxvdy14OiBoaWRkZW47XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG4gIC13ZWJraXQtb3ZlcmZsb3ctc2Nyb2xsaW5nOiB0b3VjaDtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICBib3JkZXItaW5saW5lLWVuZDogdmFyKC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1Cb3JkZXJJbmxpbmVFbmRXaWR0aCkgc29saWQgdmFyKC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1Cb3JkZXJJbmxpbmVFbmRDb2xvcik7XFxcXG4gIG9wYWNpdHk6IHZhcigtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tT3BhY2l0eSk7XFxcXG4gIHRyYW5zaXRpb246IHZhcigtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tVHJhbnNpdGlvblByb3BlcnR5KSB2YXIoLS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLVRyYW5zaXRpb25EdXJhdGlvbikgdmFyKC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1UcmFuc2l0aW9uVGltaW5nRnVuY3Rpb24pO1xcXFxuICB0cmFuc2Zvcm06IHRyYW5zbGF0ZVgodmFyKC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1UcmFuc2xhdGVYKSkgdHJhbnNsYXRlWih2YXIoLS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLVRyYW5zbGF0ZVopKTtcXFxcbn1cXFxcblxcXFxuI2JvZHkge1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLWJvZHktLVBhZGRpbmdJbmxpbmVTdGFydCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLWVuZDogdmFyKC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLWJvZHktLVBhZGRpbmdJbmxpbmVFbmQpO1xcXFxuICBmbGV4LWdyb3c6IDE7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtleHBhbmRlZF0pIHtcXFxcbiAgLS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLU9wYWNpdHk6IHZhcigtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tbS1leHBhbmRlZC0tT3BhY2l0eSk7XFxcXG4gIC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1UcmFuc2xhdGVYOiB2YXIoLS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLW0tZXhwYW5kZWQtLVRyYW5zbGF0ZVgpO1xcXFxuICBib3gtc2hhZG93OiB2YXIoLS1wZi12Ni1jLXBhZ2VfX3NpZGViYXItLUJveFNoYWRvdyk7XFxcXG5cXFxcbiAgQG1lZGlhIHNjcmVlbiBhbmQgKG1pbi13aWR0aDogNzVyZW0pIHtcXFxcbiAgICAtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tQm94U2hhZG93OiAwO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbjpob3N0KFtjb2xsYXBzZWRdKSB7XFxcXG4gIC0tcGYtdjYtYy1wYWdlX19zaWRlYmFyLS1PcGFjaXR5OiAwO1xcXFxuICAtLXBmLXY2LWMtcGFnZV9fc2lkZWJhci0tVHJhbnNsYXRlWDogLTEwMCU7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLGcyREFBazJELENBQUM7QUFBRSxJQUFPLDZCQUFROzs7QURBejZEO0FBV0EsK0JBQUMsY0FBYyxvQkFBb0I7QUFDNUIsSUFBTSxrQkFBTixlQUE4QixpQkFLbkMsa0JBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxpQkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBUlAsSUFBVztBQUFBLEVBVzlDLGNBQWM7QUFDWixVQUFNO0FBVFIsbUNBQWEsS0FBSyxnQkFBZ0I7QUFHbEMsdUJBQVMsWUFBWSxrQkFBckIsZ0JBQXFCLFNBQXJCO0FBR0EsdUJBQVMsV0FBVyxrQkFBcEIsaUJBQW9CLFNBQXBCO0FBSUUsdUJBQUssWUFBVyxPQUFPO0FBQUEsRUFDekI7QUFBQSxFQUVBLFFBQVEsU0FBK0I7QUFDckMsUUFBSSxRQUFRLElBQUksV0FBVyxHQUFHO0FBQzVCLFdBQUssV0FBVyxDQUFDLEtBQUs7QUFBQSxJQUN4QixXQUFXLFFBQVEsSUFBSSxVQUFVLEdBQUc7QUFDbEMsV0FBSyxZQUFZLENBQUMsS0FBSztBQUFBLElBQ3pCO0FBQUEsRUFDRjtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBS1Q7QUFDRjtBQS9CTztBQUdMO0FBR1M7QUFHQTtBQUhULDRCQUFTLGFBRFQsZ0JBTFcsaUJBTUY7QUFHVCw0QkFBUyxZQURULGVBUlcsaUJBU0Y7QUFURSxrQkFBTiwrQ0FEUCw2QkFDYTtBQUNYLGNBRFcsaUJBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
