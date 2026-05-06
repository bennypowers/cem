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

// elements/cem-pf-v6-skip-to-content/cem-pf-v6-skip-to-content.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-skip-to-content/cem-pf-v6-skip-to-content.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-inline-start: 50%;\\n  z-index: var(--pf-t--global--z-index--xl);\\n  transform: translateX(-50%);\\n}\\n\\n#wrapper {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-inline-start: 50%;\\n  z-index: var(--pf-t--global--z-index--xl);\\n  transform: translate(-50%, -100%);\\n  transition: transform var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n\\n  \\u0026:focus-within {\\n    transform: translate(-50%, 0);\\n  }\\n}\\n\\n#link {\\n  display: inline-flex;\\n  align-items: center;\\n  justify-content: center;\\n  padding-block-start: var(--pf-t--global--spacer--control--vertical--default);\\n  padding-block-end: var(--pf-t--global--spacer--control--vertical--default);\\n  padding-inline-start: var(--pf-t--global--spacer--control--horizontal--default);\\n  padding-inline-end: var(--pf-t--global--spacer--control--horizontal--default);\\n  font-family: var(--pf-t--global--font--family--body);\\n  font-size: var(--pf-t--global--font--size--body--default);\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--inverse);\\n  text-align: center;\\n  text-decoration: none;\\n  white-space: nowrap;\\n  background-color: var(--pf-t--global--color--brand--default);\\n  border: 0;\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  cursor: pointer;\\n\\n  \\u0026:hover {\\n    background-color: var(--pf-t--global--color--brand--hover);\\n  }\\n\\n  \\u0026:active {\\n    background-color: var(--pf-t--global--color--brand--clicked);\\n  }\\n\\n  \\u0026:focus {\\n    outline: var(--pf-t--global--border--width--action--default) solid var(--pf-t--global--focus-ring--color--default);\\n    outline-offset: var(--pf-t--global--focus-ring--position--offset);\\n  }\\n}\\n"'));
var cem_pf_v6_skip_to_content_default = s;

// elements/cem-pf-v6-skip-to-content/cem-pf-v6-skip-to-content.ts
var _href_dec, _a, _PfV6SkipToContent_decorators, _init, _href;
_PfV6SkipToContent_decorators = [customElement("cem-pf-v6-skip-to-content")];
var PfV6SkipToContent = class extends (_a = LitElement, _href_dec = [property({ reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _href, __runInitializers(_init, 8, this, "#main-content")), __runInitializers(_init, 11, this);
  }
  render() {
    return html`
      <div id="wrapper">
        <a id="link"
           part="link"
           href="${this.href}">
          <span id="text">
            <slot>Skip to content</slot>
          </span>
        </a>
      </div>
    `;
  }
};
_init = __decoratorStart(_a);
_href = new WeakMap();
__decorateElement(_init, 4, "href", _href_dec, PfV6SkipToContent, _href);
PfV6SkipToContent = __decorateElement(_init, 0, "PfV6SkipToContent", _PfV6SkipToContent_decorators, PfV6SkipToContent);
__publicField(PfV6SkipToContent, "styles", cem_pf_v6_skip_to_content_default);
__runInitializers(_init, 1, PfV6SkipToContent);
export {
  PfV6SkipToContent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXNraXAtdG8tY29udGVudC9jZW0tcGYtdjYtc2tpcC10by1jb250ZW50LnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LXNraXAtdG8tY29udGVudC9jZW0tcGYtdjYtc2tpcC10by1jb250ZW50LmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtc2tpcC10by1jb250ZW50LmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBTa2lwIHRvIENvbnRlbnRcbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3Igc2tpcCBsaW5rIHRleHQgKGRlZmF1bHRzIHRvIFwiU2tpcCB0byBjb250ZW50XCIpXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tcGYtdjYtc2tpcC10by1jb250ZW50JylcbmV4cG9ydCBjbGFzcyBQZlY2U2tpcFRvQ29udGVudCBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIC8qKiBUYXJnZXQgYW5jaG9yIElEIHRvIHNraXAgdG8gKGUuZy4sIFwiI21haW4tY29udGVudFwiKSAqL1xuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGhyZWYgPSAnI21haW4tY29udGVudCc7XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cIndyYXBwZXJcIj5cbiAgICAgICAgPGEgaWQ9XCJsaW5rXCJcbiAgICAgICAgICAgcGFydD1cImxpbmtcIlxuICAgICAgICAgICBocmVmPVwiJHt0aGlzLmhyZWZ9XCI+XG4gICAgICAgICAgPHNwYW4gaWQ9XCJ0ZXh0XCI+XG4gICAgICAgICAgICA8c2xvdD5Ta2lwIHRvIGNvbnRlbnQ8L3Nsb3Q+XG4gICAgICAgICAgPC9zcGFuPlxuICAgICAgICA8L2E+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1wZi12Ni1za2lwLXRvLWNvbnRlbnQnOiBQZlY2U2tpcFRvQ29udGVudDtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIGluc2V0LWJsb2NrLXN0YXJ0OiAwO1xcXFxuICBpbnNldC1pbmxpbmUtc3RhcnQ6IDUwJTtcXFxcbiAgei1pbmRleDogdmFyKC0tcGYtdC0tZ2xvYmFsLS16LWluZGV4LS14bCk7XFxcXG4gIHRyYW5zZm9ybTogdHJhbnNsYXRlWCgtNTAlKTtcXFxcbn1cXFxcblxcXFxuI3dyYXBwZXIge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIGluc2V0LWJsb2NrLXN0YXJ0OiAwO1xcXFxuICBpbnNldC1pbmxpbmUtc3RhcnQ6IDUwJTtcXFxcbiAgei1pbmRleDogdmFyKC0tcGYtdC0tZ2xvYmFsLS16LWluZGV4LS14bCk7XFxcXG4gIHRyYW5zZm9ybTogdHJhbnNsYXRlKC01MCUsIC0xMDAlKTtcXFxcbiAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjpmb2N1cy13aXRoaW4ge1xcXFxuICAgIHRyYW5zZm9ybTogdHJhbnNsYXRlKC01MCUsIDApO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNsaW5rIHtcXFxcbiAgZGlzcGxheTogaW5saW5lLWZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGp1c3RpZnktY29udGVudDogY2VudGVyO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tY29udHJvbC0tdmVydGljYWwtLWRlZmF1bHQpO1xcXFxuICBwYWRkaW5nLWJsb2NrLWVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLXZlcnRpY2FsLS1kZWZhdWx0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS1ob3Jpem9udGFsLS1kZWZhdWx0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tY29udHJvbC0taG9yaXpvbnRhbC0tZGVmYXVsdCk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLWZhbWlseS0tYm9keSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1kZWZhdWx0KTtcXFxcbiAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0td2VpZ2h0LS1ib2R5LS1ib2xkKTtcXFxcbiAgbGluZS1oZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tbGluZS1oZWlnaHQtLWJvZHkpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0taW52ZXJzZSk7XFxcXG4gIHRleHQtYWxpZ246IGNlbnRlcjtcXFxcbiAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuICB3aGl0ZS1zcGFjZTogbm93cmFwO1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlcjogMDtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0taG92ZXIpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmFjdGl2ZSB7XFxcXG4gICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWNsaWNrZWQpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmZvY3VzIHtcXFxcbiAgICBvdXRsaW5lOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWFjdGlvbi0tZGVmYXVsdCkgc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb2N1cy1yaW5nLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gICAgb3V0bGluZS1vZmZzZXQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9jdXMtcmluZy0tcG9zaXRpb24tLW9mZnNldCk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxrNkRBQW82RCxDQUFDO0FBQUUsSUFBTyxvQ0FBUTs7O0FEQTMrRDtBQVdBLGlDQUFDLGNBQWMsMkJBQTJCO0FBQ25DLElBQU0sb0JBQU4sZUFBZ0MsaUJBSXJDLGFBQUMsU0FBUyxFQUFFLFNBQVMsS0FBSyxDQUFDLElBSlUsSUFBVztBQUFBLEVBQTNDO0FBQUE7QUFLTCx1QkFBUyxPQUFPLGtCQUFoQixnQkFBZ0IsbUJBQWhCO0FBQUE7QUFBQSxFQUVBLFNBQVM7QUFDUCxXQUFPO0FBQUE7QUFBQTtBQUFBO0FBQUEsbUJBSVEsS0FBSyxJQUFJO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFPMUI7QUFDRjtBQXBCTztBQUtJO0FBQVQsNEJBQVMsUUFEVCxXQUpXLG1CQUtGO0FBTEUsb0JBQU4saURBRFAsK0JBQ2E7QUFDWCxjQURXLG1CQUNKLFVBQVM7QUFEWCw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
