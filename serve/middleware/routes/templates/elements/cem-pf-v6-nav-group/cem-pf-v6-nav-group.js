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
var __accessCheck = (obj, member, msg) => member.has(obj) || __typeError("Cannot " + msg);
var __privateIn = (member, obj) => Object(obj) !== obj ? __typeError('Cannot use the "in" operator on this value') : member.has(obj);
var __privateGet = (obj, member, getter) => (__accessCheck(obj, member, "read from private field"), getter ? getter.call(obj) : member.get(obj));
var __privateSet = (obj, member, value, setter) => (__accessCheck(obj, member, "write to private field"), setter ? setter.call(obj, value) : member.set(obj, value), value);
var __privateMethod = (obj, member, method) => (__accessCheck(obj, member, "access private method"), method);

// elements/cem-pf-v6-nav-group/cem-pf-v6-nav-group.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-pf-v6-nav-group/cem-pf-v6-nav-group.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  display: grid;\\n  max-height: 99999px;\\n  padding-block-start: var(--cem-pf-v6-c-nav__subnav--PaddingBlockStart, var(--pf-t--global--spacer--sm));\\n  padding-block-end: var(--cem-pf-v6-c-nav__subnav--PaddingBlockEnd, var(--pf-t--global--spacer--sm));\\n  padding-inline-start: var(--cem-pf-v6-c-nav__subnav--PaddingInlineStart, var(--pf-t--global--spacer--md));\\n  overflow-y: clip;\\n  visibility: visible;\\n  opacity: 1;\\n  translate: 0 var(--cem-pf-v6-c-nav__subnav--TranslateY, 0);\\n  transition: opacity var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default),\\n              translate var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default),\\n              visibility 0s 0s,\\n              max-height 0s 0s,\\n              padding-block-start 0s 0s,\\n              padding-block-end 0s 0s;\\n}\\n\\n@media (prefers-reduced-motion: no-preference) {\\n  :host {\\n    --cem-pf-v6-c-nav__subnav--TranslateY: 0;\\n  }\\n\\n  :host([hidden]) {\\n    --cem-pf-v6-c-nav__subnav--TranslateY: var(--cem-pf-v6-c-nav__subnav--hidden--TranslateY, -.5rem);\\n  }\\n}\\n\\n:host([hidden]) {\\n  --cem-pf-v6-c-nav__subnav--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-nav__subnav--PaddingBlockEnd: 0;\\n\\n  display: grid;\\n  max-height: 0;\\n  visibility: hidden;\\n  opacity: 0;\\n  transition-delay: var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default);\\n  transition-duration: var(--pf-t--global--motion--duration--fade--short), var(--pf-t--global--motion--duration--fade--short), 0s, 0s, 0s, 0s;\\n}\\n\\n#list {\\n  display: grid;\\n  row-gap: var(--cem-pf-v6-c-nav__subnav--RowGap, var(--pf-t--global--border--width--extra-strong));\\n}\\n"'));
var cem_pf_v6_nav_group_default = s;

// elements/cem-pf-v6-nav-group/cem-pf-v6-nav-group.ts
var _PfV6NavGroup_decorators, _init, _a;
_PfV6NavGroup_decorators = [customElement("cem-pf-v6-nav-group")];
var PfV6NavGroup = class extends (_a = LitElement) {
  static styles = cem_pf_v6_nav_group_default;
  #internals = this.attachInternals();
  constructor() {
    super();
    this.#internals.role = "region";
  }
  render() {
    return html`
      <section id="subnav"
               part="subnav">
        <div id="list"
             role="list">
          <slot></slot>
        </div>
      </section>
    `;
  }
};
_init = __decoratorStart(_a);
PfV6NavGroup = __decorateElement(_init, 0, "PfV6NavGroup", _PfV6NavGroup_decorators, PfV6NavGroup);
__runInitializers(_init, 1, PfV6NavGroup);
export {
  PfV6NavGroup
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LW5hdi1ncm91cC9jZW0tcGYtdjYtbmF2LWdyb3VwLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LW5hdi1ncm91cC9jZW0tcGYtdjYtbmF2LWdyb3VwLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtbmF2LWdyb3VwLmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBOYXZpZ2F0aW9uIEdyb3VwIChOZXN0ZWQvU3VibmF2KVxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBuYXYtbGlzdCBjb250YWluaW5nIG5hdi1pdGVtc1xuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXBmLXY2LW5hdi1ncm91cCcpXG5leHBvcnQgY2xhc3MgUGZWNk5hdkdyb3VwIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgI2ludGVybmFscyA9IHRoaXMuYXR0YWNoSW50ZXJuYWxzKCk7XG5cbiAgY29uc3RydWN0b3IoKSB7XG4gICAgc3VwZXIoKTtcbiAgICB0aGlzLiNpbnRlcm5hbHMucm9sZSA9ICdyZWdpb24nO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPHNlY3Rpb24gaWQ9XCJzdWJuYXZcIlxuICAgICAgICAgICAgICAgcGFydD1cInN1Ym5hdlwiPlxuICAgICAgICA8ZGl2IGlkPVwibGlzdFwiXG4gICAgICAgICAgICAgcm9sZT1cImxpc3RcIj5cbiAgICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9zZWN0aW9uPlxuICAgIGA7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXBmLXY2LW5hdi1ncm91cCc6IFBmVjZOYXZHcm91cDtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuXFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIG1heC1oZWlnaHQ6IDk5OTk5cHg7XFxcXG4gIHBhZGRpbmctYmxvY2stc3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fc3VibmF2LS1QYWRkaW5nQmxvY2tTdGFydCwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKSk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1uYXZfX3N1Ym5hdi0tUGFkZGluZ0Jsb2NrRW5kLCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fc3VibmF2LS1QYWRkaW5nSW5saW5lU3RhcnQsIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCkpO1xcXFxuICBvdmVyZmxvdy15OiBjbGlwO1xcXFxuICB2aXNpYmlsaXR5OiB2aXNpYmxlO1xcXFxuICBvcGFjaXR5OiAxO1xcXFxuICB0cmFuc2xhdGU6IDAgdmFyKC0tY2VtLXBmLXY2LWMtbmF2X19zdWJuYXYtLVRyYW5zbGF0ZVksIDApO1xcXFxuICB0cmFuc2l0aW9uOiBvcGFjaXR5IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tZGVmYXVsdCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCksXFxcXG4gICAgICAgICAgICAgIHRyYW5zbGF0ZSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLWRlZmF1bHQpIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpLFxcXFxuICAgICAgICAgICAgICB2aXNpYmlsaXR5IDBzIDBzLFxcXFxuICAgICAgICAgICAgICBtYXgtaGVpZ2h0IDBzIDBzLFxcXFxuICAgICAgICAgICAgICBwYWRkaW5nLWJsb2NrLXN0YXJ0IDBzIDBzLFxcXFxuICAgICAgICAgICAgICBwYWRkaW5nLWJsb2NrLWVuZCAwcyAwcztcXFxcbn1cXFxcblxcXFxuQG1lZGlhIChwcmVmZXJzLXJlZHVjZWQtbW90aW9uOiBuby1wcmVmZXJlbmNlKSB7XFxcXG4gIDpob3N0IHtcXFxcbiAgICAtLWNlbS1wZi12Ni1jLW5hdl9fc3VibmF2LS1UcmFuc2xhdGVZOiAwO1xcXFxuICB9XFxcXG5cXFxcbiAgOmhvc3QoW2hpZGRlbl0pIHtcXFxcbiAgICAtLWNlbS1wZi12Ni1jLW5hdl9fc3VibmF2LS1UcmFuc2xhdGVZOiB2YXIoLS1jZW0tcGYtdjYtYy1uYXZfX3N1Ym5hdi0taGlkZGVuLS1UcmFuc2xhdGVZLCAtLjVyZW0pO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbjpob3N0KFtoaWRkZW5dKSB7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbmF2X19zdWJuYXYtLVBhZGRpbmdCbG9ja1N0YXJ0OiAwO1xcXFxuICAtLWNlbS1wZi12Ni1jLW5hdl9fc3VibmF2LS1QYWRkaW5nQmxvY2tFbmQ6IDA7XFxcXG5cXFxcbiAgZGlzcGxheTogZ3JpZDtcXFxcbiAgbWF4LWhlaWdodDogMDtcXFxcbiAgdmlzaWJpbGl0eTogaGlkZGVuO1xcXFxuICBvcGFjaXR5OiAwO1xcXFxuICB0cmFuc2l0aW9uLWRlbGF5OiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLWRlZmF1bHQpLCB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLWRlZmF1bHQpLCB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLWRlZmF1bHQpLCB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLWRlZmF1bHQpLCB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLWRlZmF1bHQpLCB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLWRlZmF1bHQpO1xcXFxuICB0cmFuc2l0aW9uLWR1cmF0aW9uOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLXNob3J0KSwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1zaG9ydCksIDBzLCAwcywgMHMsIDBzO1xcXFxufVxcXFxuXFxcXG4jbGlzdCB7XFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIHJvdy1nYXA6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fc3VibmF2LS1Sb3dHYXAsIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tZXh0cmEtc3Ryb25nKSk7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjs7O0FDRDlCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLHEvREFBdS9ELENBQUM7QUFBRSxJQUFPLDhCQUFROzs7QURBOWpFO0FBV0EsNEJBQUMsY0FBYyxxQkFBcUI7QUFDN0IsSUFBTSxlQUFOLGVBQTJCLGlCQUFXO0FBQUEsRUFDM0MsT0FBTyxTQUFTO0FBQUEsRUFFaEIsYUFBYSxLQUFLLGdCQUFnQjtBQUFBLEVBRWxDLGNBQWM7QUFDWixVQUFNO0FBQ04sU0FBSyxXQUFXLE9BQU87QUFBQSxFQUN6QjtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFTVDtBQUNGO0FBckJPO0FBQU0sZUFBTiw0Q0FEUCwwQkFDYTtBQUFOLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
