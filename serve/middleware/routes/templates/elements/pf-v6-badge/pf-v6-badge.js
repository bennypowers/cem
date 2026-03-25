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

// elements/pf-v6-badge/pf-v6-badge.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/pf-v6-badge/pf-v6-badge.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --pf-v6-c-badge--BorderColor: transparent;\\n  --pf-v6-c-badge--BorderWidth: var(--pf-t--global--border--width--regular);\\n  --pf-v6-c-badge--BorderRadius: var(--pf-t--global--border--radius--pill);\\n  --pf-v6-c-badge--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --pf-v6-c-badge--FontWeight: var(--pf-t--global--font--weight--body--bold);\\n  --pf-v6-c-badge--PaddingInlineEnd: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-badge--PaddingInlineStart: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-badge--Color: var(--pf-t--global--text--color--nonstatus--on-gray--default);\\n  --pf-v6-c-badge--MinWidth: var(--pf-t--global--spacer--xl);\\n  --pf-v6-c-badge--BackgroundColor: var(--pf-t--global--color--nonstatus--gray--default);\\n\\n  position: relative;\\n  display: inline-block;\\n  min-width: var(--pf-v6-c-badge--MinWidth);\\n  padding-inline-start: var(--pf-v6-c-badge--PaddingInlineStart);\\n  padding-inline-end: var(--pf-v6-c-badge--PaddingInlineEnd);\\n  font-size: var(--pf-v6-c-badge--FontSize);\\n  font-weight: var(--pf-v6-c-badge--FontWeight);\\n  color: var(--pf-v6-c-badge--Color);\\n  text-align: center;\\n  white-space: nowrap;\\n  background-color: var(--pf-v6-c-badge--BackgroundColor);\\n  border-radius: var(--pf-v6-c-badge--BorderRadius);\\n}\\n\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border: var(--pf-v6-c-badge--BorderWidth) solid var(--pf-v6-c-badge--BorderColor);\\n  border-radius: inherit;\\n}\\n\\n:host([read]) {\\n  --pf-v6-c-badge--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --pf-v6-c-badge--Color: var(--pf-t--global--text--color--nonstatus--on-gray--default);\\n  --pf-v6-c-badge--BackgroundColor: var(--pf-t--global--color--nonstatus--gray--default);\\n}\\n\\n:host([unread]) {\\n  --pf-v6-c-badge--Color: var(--pf-t--global--text--color--on-brand--default);\\n  --pf-v6-c-badge--BackgroundColor: var(--pf-t--global--color--brand--default);\\n}\\n\\n:host([disabled]) {\\n  --pf-v6-c-badge--Color: var(--pf-t--global--text--color--on-disabled);\\n  --pf-v6-c-badge--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n}\\n\\n:host([disabled])::after {\\n  border-color: var(--pf-t--global--border--color--disabled);\\n}\\n\\n#sr-text {\\n  position: absolute;\\n  width: 1px;\\n  height: 1px;\\n  padding: 0;\\n  margin: -1px;\\n  overflow: hidden;\\n  clip: rect(0, 0, 0, 0);\\n  white-space: nowrap;\\n  border: 0;\\n}\\n"'));
var pf_v6_badge_default = s;

// elements/pf-v6-badge/pf-v6-badge.ts
var _screenReaderText_dec, _disabled_dec, _unread_dec, _read_dec, _a, _PfV6Badge_decorators, _init, _read, _unread, _disabled, _screenReaderText;
_PfV6Badge_decorators = [customElement("pf-v6-badge")];
var PfV6Badge = class extends (_a = LitElement, _read_dec = [property({ type: Boolean, reflect: true })], _unread_dec = [property({ type: Boolean, reflect: true })], _disabled_dec = [property({ type: Boolean, reflect: true })], _screenReaderText_dec = [property({ attribute: "screen-reader-text" })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _read, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _unread, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateAdd(this, _disabled, __runInitializers(_init, 16, this, false)), __runInitializers(_init, 19, this);
    __privateAdd(this, _screenReaderText, __runInitializers(_init, 20, this)), __runInitializers(_init, 23, this);
  }
  render() {
    return html`
      <slot></slot>
      ${this.screenReaderText ? html`
        <span id="sr-text">${this.screenReaderText}</span>
      ` : ""}
    `;
  }
};
_init = __decoratorStart(_a);
_read = new WeakMap();
_unread = new WeakMap();
_disabled = new WeakMap();
_screenReaderText = new WeakMap();
__decorateElement(_init, 4, "read", _read_dec, PfV6Badge, _read);
__decorateElement(_init, 4, "unread", _unread_dec, PfV6Badge, _unread);
__decorateElement(_init, 4, "disabled", _disabled_dec, PfV6Badge, _disabled);
__decorateElement(_init, 4, "screenReaderText", _screenReaderText_dec, PfV6Badge, _screenReaderText);
PfV6Badge = __decorateElement(_init, 0, "PfV6Badge", _PfV6Badge_decorators, PfV6Badge);
__publicField(PfV6Badge, "styles", pf_v6_badge_default);
__runInitializers(_init, 1, PfV6Badge);
export {
  PfV6Badge
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtYmFkZ2UvcGYtdjYtYmFkZ2UudHMiLCAibGl0LWNzczplbGVtZW50cy9wZi12Ni1iYWRnZS9wZi12Ni1iYWRnZS5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtYmFkZ2UuY3NzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IEJhZGdlXG4gKlxuICogQHNsb3QgLSBCYWRnZSBjb250ZW50ICh0eXBpY2FsbHkgbnVtYmVycyBvciBzaG9ydCB0ZXh0KVxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtYmFkZ2UnKVxuZXhwb3J0IGNsYXNzIFBmVjZCYWRnZSBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgcmVhZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgdW5yZWFkID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBkaXNhYmxlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ3NjcmVlbi1yZWFkZXItdGV4dCcgfSlcbiAgYWNjZXNzb3Igc2NyZWVuUmVhZGVyVGV4dDogc3RyaW5nIHwgdW5kZWZpbmVkO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxzbG90Pjwvc2xvdD5cbiAgICAgICR7dGhpcy5zY3JlZW5SZWFkZXJUZXh0ID8gaHRtbGBcbiAgICAgICAgPHNwYW4gaWQ9XCJzci10ZXh0XCI+JHt0aGlzLnNjcmVlblJlYWRlclRleHR9PC9zcGFuPlxuICAgICAgYCA6ICcnfVxuICAgIGA7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAncGYtdjYtYmFkZ2UnOiBQZlY2QmFkZ2U7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcblxcXFxuICAtLXBmLXY2LWMtYmFkZ2UtLUJvcmRlckNvbG9yOiB0cmFuc3BhcmVudDtcXFxcbiAgLS1wZi12Ni1jLWJhZGdlLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1yZWd1bGFyKTtcXFxcbiAgLS1wZi12Ni1jLWJhZGdlLS1Cb3JkZXJSYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXBpbGwpO1xcXFxuICAtLXBmLXY2LWMtYmFkZ2UtLUZvbnRTaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtKTtcXFxcbiAgLS1wZi12Ni1jLWJhZGdlLS1Gb250V2VpZ2h0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXdlaWdodC0tYm9keS0tYm9sZCk7XFxcXG4gIC0tcGYtdjYtYy1iYWRnZS0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgLS1wZi12Ni1jLWJhZGdlLS1QYWRkaW5nSW5saW5lU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIC0tcGYtdjYtYy1iYWRnZS0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW5vbnN0YXR1cy0tb24tZ3JheS0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1iYWRnZS0tTWluV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS14bCk7XFxcXG4gIC0tcGYtdjYtYy1iYWRnZS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1ub25zdGF0dXMtLWdyYXktLWRlZmF1bHQpO1xcXFxuXFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICBtaW4td2lkdGg6IHZhcigtLXBmLXY2LWMtYmFkZ2UtLU1pbldpZHRoKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLXBmLXY2LWMtYmFkZ2UtLVBhZGRpbmdJbmxpbmVTdGFydCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLWVuZDogdmFyKC0tcGYtdjYtYy1iYWRnZS0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdjYtYy1iYWRnZS0tRm9udFNpemUpO1xcXFxuICBmb250LXdlaWdodDogdmFyKC0tcGYtdjYtYy1iYWRnZS0tRm9udFdlaWdodCk7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi12Ni1jLWJhZGdlLS1Db2xvcik7XFxcXG4gIHRleHQtYWxpZ246IGNlbnRlcjtcXFxcbiAgd2hpdGUtc3BhY2U6IG5vd3JhcDtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdjYtYy1iYWRnZS0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdjYtYy1iYWRnZS0tQm9yZGVyUmFkaXVzKTtcXFxcbn1cXFxcblxcXFxuOmhvc3Q6OmFmdGVyIHtcXFxcbiAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICBpbnNldDogMDtcXFxcbiAgcG9pbnRlci1ldmVudHM6IG5vbmU7XFxcXG4gIGNvbnRlbnQ6IFxcXFxcXFwiXFxcXFxcXCI7XFxcXG4gIGJvcmRlcjogdmFyKC0tcGYtdjYtYy1iYWRnZS0tQm9yZGVyV2lkdGgpIHNvbGlkIHZhcigtLXBmLXY2LWMtYmFkZ2UtLUJvcmRlckNvbG9yKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogaW5oZXJpdDtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3JlYWRdKSB7XFxcXG4gIC0tcGYtdjYtYy1iYWRnZS0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0taGlnaC1jb250cmFzdCk7XFxcXG4gIC0tcGYtdjYtYy1iYWRnZS0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW5vbnN0YXR1cy0tb24tZ3JheS0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1iYWRnZS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1ub25zdGF0dXMtLWdyYXktLWRlZmF1bHQpO1xcXFxufVxcXFxuXFxcXG46aG9zdChbdW5yZWFkXSkge1xcXFxuICAtLXBmLXY2LWMtYmFkZ2UtLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1icmFuZC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1iYWRnZS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtkaXNhYmxlZF0pIHtcXFxcbiAgLS1wZi12Ni1jLWJhZGdlLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tZGlzYWJsZWQpO1xcXFxuICAtLXBmLXY2LWMtYmFkZ2UtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tZGlzYWJsZWQtLWRlZmF1bHQpO1xcXFxufVxcXFxuXFxcXG46aG9zdChbZGlzYWJsZWRdKTo6YWZ0ZXIge1xcXFxuICBib3JkZXItY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGlzYWJsZWQpO1xcXFxufVxcXFxuXFxcXG4jc3ItdGV4dCB7XFxcXG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgd2lkdGg6IDFweDtcXFxcbiAgaGVpZ2h0OiAxcHg7XFxcXG4gIHBhZGRpbmc6IDA7XFxcXG4gIG1hcmdpbjogLTFweDtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgY2xpcDogcmVjdCgwLCAwLCAwLCAwKTtcXFxcbiAgd2hpdGUtc3BhY2U6IG5vd3JhcDtcXFxcbiAgYm9yZGVyOiAwO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLHE5RUFBeTlFLENBQUM7QUFBRSxJQUFPLHNCQUFROzs7QURBaGlGO0FBV0EseUJBQUMsY0FBYyxhQUFhO0FBQ3JCLElBQU0sWUFBTixlQUF3QixpQkFHN0IsYUFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBRzFDLGVBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxpQkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBRzFDLHlCQUFDLFNBQVMsRUFBRSxXQUFXLHFCQUFxQixDQUFDLElBWmhCLElBQVc7QUFBQSxFQUFuQztBQUFBO0FBSUwsdUJBQVMsT0FBTyxrQkFBaEIsZ0JBQWdCLFNBQWhCO0FBR0EsdUJBQVMsU0FBUyxrQkFBbEIsaUJBQWtCLFNBQWxCO0FBR0EsdUJBQVMsV0FBVyxrQkFBcEIsaUJBQW9CLFNBQXBCO0FBR0EsdUJBQVMsbUJBQVQ7QUFBQTtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQTtBQUFBLFFBRUgsS0FBSyxtQkFBbUI7QUFBQSw2QkFDSCxLQUFLLGdCQUFnQjtBQUFBLFVBQ3hDLEVBQUU7QUFBQTtBQUFBLEVBRVY7QUFDRjtBQXZCTztBQUlJO0FBR0E7QUFHQTtBQUdBO0FBVFQsNEJBQVMsUUFEVCxXQUhXLFdBSUY7QUFHVCw0QkFBUyxVQURULGFBTlcsV0FPRjtBQUdULDRCQUFTLFlBRFQsZUFUVyxXQVVGO0FBR1QsNEJBQVMsb0JBRFQsdUJBWlcsV0FhRjtBQWJFLFlBQU4seUNBRFAsdUJBQ2E7QUFDWCxjQURXLFdBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
