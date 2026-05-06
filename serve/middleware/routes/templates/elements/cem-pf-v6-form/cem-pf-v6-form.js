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

// elements/cem-pf-v6-form/cem-pf-v6-form.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-form/cem-pf-v6-form.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\n#form {\\n  display: block;\\n}\\n\\n:host([horizontal]) #form {\\n  --_form-column-gap: var(--pf-t--global--spacer--md);\\n  display: grid;\\n  grid-template-columns: minmax(0, auto) 1fr;\\n  column-gap: var(--_form-column-gap);\\n  row-gap: var(--pf-t--global--spacer--sm);\\n\\n  \\u0026 ::slotted(cem-pf-v6-form-field-group) {\\n    grid-column: 1 / -1;\\n  }\\n}\\n"'));
var cem_pf_v6_form_default = s;

// elements/cem-pf-v6-form/cem-pf-v6-form.ts
var _horizontal_dec, _a, _PfV6Form_decorators, _init, _horizontal, _PfV6Form_instances, onSubmit_fn, formElement_get;
_PfV6Form_decorators = [customElement("cem-pf-v6-form")];
var PfV6Form = class extends (_a = LitElement, _horizontal_dec = [property({ type: Boolean, reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Form_instances);
    __privateAdd(this, _horizontal, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
  }
  render() {
    return html`
      <form id="form"
            @submit=${__privateMethod(this, _PfV6Form_instances, onSubmit_fn)}>
        <slot></slot>
      </form>
    `;
  }
  /** Submits the form programmatically */
  submit() {
    __privateGet(this, _PfV6Form_instances, formElement_get)?.submit();
  }
  /**
   * Requests form submission with validation
   * @param submitter - Optional submitter element
   */
  requestSubmit(submitter) {
    return __privateGet(this, _PfV6Form_instances, formElement_get)?.requestSubmit(submitter);
  }
  /** Resets the form to default values */
  reset() {
    __privateGet(this, _PfV6Form_instances, formElement_get)?.reset();
  }
  /** Checks form validity */
  checkValidity() {
    return __privateGet(this, _PfV6Form_instances, formElement_get)?.checkValidity() ?? true;
  }
  /** Reports form validity (shows validation messages) */
  reportValidity() {
    return __privateGet(this, _PfV6Form_instances, formElement_get)?.reportValidity() ?? true;
  }
};
_init = __decoratorStart(_a);
_horizontal = new WeakMap();
_PfV6Form_instances = new WeakSet();
onSubmit_fn = function(e) {
  this.dispatchEvent(new SubmitEvent("submit", {
    bubbles: true,
    cancelable: true,
    submitter: e.submitter
  }));
};
formElement_get = function() {
  return this.shadowRoot?.getElementById("form");
};
__decorateElement(_init, 4, "horizontal", _horizontal_dec, PfV6Form, _horizontal);
PfV6Form = __decorateElement(_init, 0, "PfV6Form", _PfV6Form_decorators, PfV6Form);
__publicField(PfV6Form, "styles", cem_pf_v6_form_default);
__runInitializers(_init, 1, PfV6Form);
export {
  PfV6Form
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LWZvcm0vY2VtLXBmLXY2LWZvcm0udHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tcGYtdjYtZm9ybS9jZW0tcGYtdjYtZm9ybS5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXBmLXY2LWZvcm0uY3NzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IEZvcm0gQ29tcG9uZW50XG4gKlxuICogV3JhcHMgZm9ybSBjb250cm9scyBpbiBhIHNlbWFudGljIGA8Zm9ybT5gIGVsZW1lbnQgd2l0aCBQYXR0ZXJuRmx5IHN0eWxpbmcuXG4gKiBTdXBwb3J0cyBzdGFuZGFyZCBmb3JtIHN1Ym1pc3Npb24gYW5kIHZhbGlkYXRpb24gQVBJcy5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgZm9ybSBjb250cm9sc1xuICpcbiAqIEBhdHRyIHtib29sZWFufSBob3Jpem9udGFsIC0gRW5hYmxlcyBob3Jpem9udGFsIGxheW91dCB3aXRoIGdyaWRcbiAqIEBmaXJlcyBzdWJtaXQgLSBGb3J3YXJkZWQgZnJvbSBpbnRlcm5hbCBgPGZvcm0+YCBlbGVtZW50XG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tcGYtdjYtZm9ybScpXG5leHBvcnQgY2xhc3MgUGZWNkZvcm0gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGhvcml6b250YWwgPSBmYWxzZTtcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8Zm9ybSBpZD1cImZvcm1cIlxuICAgICAgICAgICAgQHN1Ym1pdD0ke3RoaXMuI29uU3VibWl0fT5cbiAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgPC9mb3JtPlxuICAgIGA7XG4gIH1cblxuICAjb25TdWJtaXQoZTogU3VibWl0RXZlbnQpIHtcbiAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IFN1Ym1pdEV2ZW50KCdzdWJtaXQnLCB7XG4gICAgICBidWJibGVzOiB0cnVlLFxuICAgICAgY2FuY2VsYWJsZTogdHJ1ZSxcbiAgICAgIHN1Ym1pdHRlcjogZS5zdWJtaXR0ZXIsXG4gICAgfSkpO1xuICB9XG5cbiAgLyoqIFN1Ym1pdHMgdGhlIGZvcm0gcHJvZ3JhbW1hdGljYWxseSAqL1xuICBzdWJtaXQoKSB7XG4gICAgdGhpcy4jZm9ybUVsZW1lbnQ/LnN1Ym1pdCgpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJlcXVlc3RzIGZvcm0gc3VibWlzc2lvbiB3aXRoIHZhbGlkYXRpb25cbiAgICogQHBhcmFtIHN1Ym1pdHRlciAtIE9wdGlvbmFsIHN1Ym1pdHRlciBlbGVtZW50XG4gICAqL1xuICByZXF1ZXN0U3VibWl0KHN1Ym1pdHRlcj86IEhUTUxFbGVtZW50KSB7XG4gICAgcmV0dXJuIHRoaXMuI2Zvcm1FbGVtZW50Py5yZXF1ZXN0U3VibWl0KHN1Ym1pdHRlcik7XG4gIH1cblxuICAvKiogUmVzZXRzIHRoZSBmb3JtIHRvIGRlZmF1bHQgdmFsdWVzICovXG4gIHJlc2V0KCkge1xuICAgIHRoaXMuI2Zvcm1FbGVtZW50Py5yZXNldCgpO1xuICB9XG5cbiAgLyoqIENoZWNrcyBmb3JtIHZhbGlkaXR5ICovXG4gIGNoZWNrVmFsaWRpdHkoKTogYm9vbGVhbiB7XG4gICAgcmV0dXJuIHRoaXMuI2Zvcm1FbGVtZW50Py5jaGVja1ZhbGlkaXR5KCkgPz8gdHJ1ZTtcbiAgfVxuXG4gIC8qKiBSZXBvcnRzIGZvcm0gdmFsaWRpdHkgKHNob3dzIHZhbGlkYXRpb24gbWVzc2FnZXMpICovXG4gIHJlcG9ydFZhbGlkaXR5KCk6IGJvb2xlYW4ge1xuICAgIHJldHVybiB0aGlzLiNmb3JtRWxlbWVudD8ucmVwb3J0VmFsaWRpdHkoKSA/PyB0cnVlO1xuICB9XG5cbiAgZ2V0ICNmb3JtRWxlbWVudCgpOiBIVE1MRm9ybUVsZW1lbnQgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnZm9ybScpIGFzIEhUTUxGb3JtRWxlbWVudCB8IG51bGw7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXBmLXY2LWZvcm0nOiBQZlY2Rm9ybTtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbn1cXFxcblxcXFxuI2Zvcm0ge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2hvcml6b250YWxdKSAjZm9ybSB7XFxcXG4gIC0tX2Zvcm0tY29sdW1uLWdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbiAgZGlzcGxheTogZ3JpZDtcXFxcbiAgZ3JpZC10ZW1wbGF0ZS1jb2x1bW5zOiBtaW5tYXgoMCwgYXV0bykgMWZyO1xcXFxuICBjb2x1bW4tZ2FwOiB2YXIoLS1fZm9ybS1jb2x1bW4tZ2FwKTtcXFxcbiAgcm93LWdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcblxcXFxuICBcXFxcdTAwMjYgOjpzbG90dGVkKGNlbS1wZi12Ni1mb3JtLWZpZWxkLWdyb3VwKSB7XFxcXG4gICAgZ3JpZC1jb2x1bW46IDEgLyAtMTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLDJaQUE2WixDQUFDO0FBQUUsSUFBTyx5QkFBUTs7O0FEQXBlO0FBaUJBLHdCQUFDLGNBQWMsZ0JBQWdCO0FBQ3hCLElBQU0sV0FBTixlQUF1QixpQkFHNUIsbUJBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUhkLElBQVc7QUFBQSxFQUFsQztBQUFBO0FBQUE7QUFJTCx1QkFBUyxhQUFhLGtCQUF0QixnQkFBc0IsU0FBdEI7QUFBQTtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQTtBQUFBLHNCQUVXLHNCQUFLLGlDQUFTO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFJbEM7QUFBQTtBQUFBLEVBV0EsU0FBUztBQUNQLHVCQUFLLHVDQUFjLE9BQU87QUFBQSxFQUM1QjtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFNQSxjQUFjLFdBQXlCO0FBQ3JDLFdBQU8sbUJBQUssdUNBQWMsY0FBYyxTQUFTO0FBQUEsRUFDbkQ7QUFBQTtBQUFBLEVBR0EsUUFBUTtBQUNOLHVCQUFLLHVDQUFjLE1BQU07QUFBQSxFQUMzQjtBQUFBO0FBQUEsRUFHQSxnQkFBeUI7QUFDdkIsV0FBTyxtQkFBSyx1Q0FBYyxjQUFjLEtBQUs7QUFBQSxFQUMvQztBQUFBO0FBQUEsRUFHQSxpQkFBMEI7QUFDeEIsV0FBTyxtQkFBSyx1Q0FBYyxlQUFlLEtBQUs7QUFBQSxFQUNoRDtBQUtGO0FBdERPO0FBSUk7QUFKSjtBQWVMLGNBQVMsU0FBQyxHQUFnQjtBQUN4QixPQUFLLGNBQWMsSUFBSSxZQUFZLFVBQVU7QUFBQSxJQUMzQyxTQUFTO0FBQUEsSUFDVCxZQUFZO0FBQUEsSUFDWixXQUFXLEVBQUU7QUFBQSxFQUNmLENBQUMsQ0FBQztBQUNKO0FBOEJJLGtCQUFZLFdBQTJCO0FBQ3pDLFNBQU8sS0FBSyxZQUFZLGVBQWUsTUFBTTtBQUMvQztBQWpEQSw0QkFBUyxjQURULGlCQUhXLFVBSUY7QUFKRSxXQUFOLHdDQURQLHNCQUNhO0FBQ1gsY0FEVyxVQUNKLFVBQVM7QUFEWCw0QkFBTTsiLAogICJuYW1lcyI6IFtdCn0K
