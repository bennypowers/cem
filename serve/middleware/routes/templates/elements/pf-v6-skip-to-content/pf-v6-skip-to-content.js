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

// elements/pf-v6-skip-to-content/pf-v6-skip-to-content.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-skip-to-content/pf-v6-skip-to-content.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-inline-start: 50%;\\n  z-index: var(--pf-t--global--z-index--xl);\\n  transform: translateX(-50%);\\n}\\n\\n#wrapper {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-inline-start: 50%;\\n  z-index: var(--pf-t--global--z-index--xl);\\n  transform: translate(-50%, -100%);\\n  transition: transform var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n\\n  \\u0026:focus-within {\\n    transform: translate(-50%, 0);\\n  }\\n}\\n\\n#link {\\n  display: inline-flex;\\n  align-items: center;\\n  justify-content: center;\\n  padding-block-start: var(--pf-t--global--spacer--control--vertical--default);\\n  padding-block-end: var(--pf-t--global--spacer--control--vertical--default);\\n  padding-inline-start: var(--pf-t--global--spacer--control--horizontal--default);\\n  padding-inline-end: var(--pf-t--global--spacer--control--horizontal--default);\\n  font-family: var(--pf-t--global--font--family--body);\\n  font-size: var(--pf-t--global--font--size--body--default);\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--inverse);\\n  text-align: center;\\n  text-decoration: none;\\n  white-space: nowrap;\\n  background-color: var(--pf-t--global--color--brand--default);\\n  border: 0;\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  cursor: pointer;\\n\\n  \\u0026:hover {\\n    background-color: var(--pf-t--global--color--brand--hover);\\n  }\\n\\n  \\u0026:active {\\n    background-color: var(--pf-t--global--color--brand--clicked);\\n  }\\n\\n  \\u0026:focus {\\n    outline: var(--pf-t--global--border--width--action--default) solid var(--pf-t--global--focus-ring--color--default);\\n    outline-offset: var(--pf-t--global--focus-ring--position--offset);\\n  }\\n}\\n"'));
var pf_v6_skip_to_content_default = s;

// elements/pf-v6-skip-to-content/pf-v6-skip-to-content.ts
var _href_dec, _a, _PfV6SkipToContent_decorators, _init, _href;
_PfV6SkipToContent_decorators = [customElement("pf-v6-skip-to-content")];
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
__publicField(PfV6SkipToContent, "styles", pf_v6_skip_to_content_default);
__runInitializers(_init, 1, PfV6SkipToContent);
export {
  PfV6SkipToContent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtc2tpcC10by1jb250ZW50L3BmLXY2LXNraXAtdG8tY29udGVudC50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL3BmLXY2LXNraXAtdG8tY29udGVudC9wZi12Ni1za2lwLXRvLWNvbnRlbnQuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL3BmLXY2LXNraXAtdG8tY29udGVudC5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgU2tpcCB0byBDb250ZW50XG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIHNraXAgbGluayB0ZXh0IChkZWZhdWx0cyB0byBcIlNraXAgdG8gY29udGVudFwiKVxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtc2tpcC10by1jb250ZW50JylcbmV4cG9ydCBjbGFzcyBQZlY2U2tpcFRvQ29udGVudCBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIC8qKiBUYXJnZXQgYW5jaG9yIElEIHRvIHNraXAgdG8gKGUuZy4sIFwiI21haW4tY29udGVudFwiKSAqL1xuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGhyZWYgPSAnI21haW4tY29udGVudCc7XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cIndyYXBwZXJcIj5cbiAgICAgICAgPGEgaWQ9XCJsaW5rXCJcbiAgICAgICAgICAgcGFydD1cImxpbmtcIlxuICAgICAgICAgICBocmVmPVwiJHt0aGlzLmhyZWZ9XCI+XG4gICAgICAgICAgPHNwYW4gaWQ9XCJ0ZXh0XCI+XG4gICAgICAgICAgICA8c2xvdD5Ta2lwIHRvIGNvbnRlbnQ8L3Nsb3Q+XG4gICAgICAgICAgPC9zcGFuPlxuICAgICAgICA8L2E+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LXNraXAtdG8tY29udGVudCc6IFBmVjZTa2lwVG9Db250ZW50O1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgaW5zZXQtYmxvY2stc3RhcnQ6IDA7XFxcXG4gIGluc2V0LWlubGluZS1zdGFydDogNTAlO1xcXFxuICB6LWluZGV4OiB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLXhsKTtcXFxcbiAgdHJhbnNmb3JtOiB0cmFuc2xhdGVYKC01MCUpO1xcXFxufVxcXFxuXFxcXG4jd3JhcHBlciB7XFxcXG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgaW5zZXQtYmxvY2stc3RhcnQ6IDA7XFxcXG4gIGluc2V0LWlubGluZS1zdGFydDogNTAlO1xcXFxuICB6LWluZGV4OiB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLXhsKTtcXFxcbiAgdHJhbnNmb3JtOiB0cmFuc2xhdGUoLTUwJSwgLTEwMCUpO1xcXFxuICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1zaG9ydCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmZvY3VzLXdpdGhpbiB7XFxcXG4gICAgdHJhbnNmb3JtOiB0cmFuc2xhdGUoLTUwJSwgMCk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI2xpbmsge1xcXFxuICBkaXNwbGF5OiBpbmxpbmUtZmxleDtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAganVzdGlmeS1jb250ZW50OiBjZW50ZXI7XFxcXG4gIHBhZGRpbmctYmxvY2stc3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tY29udHJvbC0tdmVydGljYWwtLWRlZmF1bHQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLWhvcml6b250YWwtLWRlZmF1bHQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1lbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS1ob3Jpem9udGFsLS1kZWZhdWx0KTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tZmFtaWx5LS1ib2R5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWRlZmF1bHQpO1xcXFxuICBmb250LXdlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS13ZWlnaHQtLWJvZHktLWJvbGQpO1xcXFxuICBsaW5lLWhlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1pbnZlcnNlKTtcXFxcbiAgdGV4dC1hbGlnbjogY2VudGVyO1xcXFxuICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG4gIHdoaXRlLXNwYWNlOiBub3dyYXA7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyOiAwO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCk7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1ob3Zlcik7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6YWN0aXZlIHtcXFxcbiAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0tY2xpY2tlZCk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6Zm9jdXMge1xcXFxuICAgIG91dGxpbmU6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tYWN0aW9uLS1kZWZhdWx0KSBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWZvY3VzLXJpbmctLWNvbG9yLS1kZWZhdWx0KTtcXFxcbiAgICBvdXRsaW5lLW9mZnNldDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb2N1cy1yaW5nLS1wb3NpdGlvbi0tb2Zmc2V0KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLGs2REFBbzZELENBQUM7QUFBRSxJQUFPLGdDQUFROzs7QURBMytEO0FBV0EsaUNBQUMsY0FBYyx1QkFBdUI7QUFDL0IsSUFBTSxvQkFBTixlQUFnQyxpQkFJckMsYUFBQyxTQUFTLEVBQUUsU0FBUyxLQUFLLENBQUMsSUFKVSxJQUFXO0FBQUEsRUFBM0M7QUFBQTtBQUtMLHVCQUFTLE9BQU8sa0JBQWhCLGdCQUFnQixtQkFBaEI7QUFBQTtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQTtBQUFBO0FBQUE7QUFBQSxtQkFJUSxLQUFLLElBQUk7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQU8xQjtBQUNGO0FBcEJPO0FBS0k7QUFBVCw0QkFBUyxRQURULFdBSlcsbUJBS0Y7QUFMRSxvQkFBTixpREFEUCwrQkFDYTtBQUNYLGNBRFcsbUJBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
