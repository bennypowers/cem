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

// elements/pf-v6-form/pf-v6-form.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/pf-v6-form/pf-v6-form.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\n#form {\\n  display: block;\\n}\\n\\n:host([horizontal]) #form {\\n  --_form-column-gap: var(--pf-t--global--spacer--md);\\n  display: grid;\\n  grid-template-columns: minmax(0, auto) 1fr;\\n  column-gap: var(--_form-column-gap);\\n  row-gap: var(--pf-t--global--spacer--sm);\\n\\n  \\u0026 ::slotted(pf-v6-form-field-group) {\\n    grid-column: 1 / -1;\\n  }\\n}\\n"'));
var pf_v6_form_default = s;

// elements/pf-v6-form/pf-v6-form.ts
var _horizontal_dec, _a, _PfV6Form_decorators, _init, _horizontal, _PfV6Form_instances, onSubmit_fn, formElement_get;
_PfV6Form_decorators = [customElement("pf-v6-form")];
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
__publicField(PfV6Form, "styles", pf_v6_form_default);
__runInitializers(_init, 1, PfV6Form);
export {
  PfV6Form
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtZm9ybS9wZi12Ni1mb3JtLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvcGYtdjYtZm9ybS9wZi12Ni1mb3JtLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9wZi12Ni1mb3JtLmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBGb3JtIENvbXBvbmVudFxuICpcbiAqIFdyYXBzIGZvcm0gY29udHJvbHMgaW4gYSBzZW1hbnRpYyBgPGZvcm0+YCBlbGVtZW50IHdpdGggUGF0dGVybkZseSBzdHlsaW5nLlxuICogU3VwcG9ydHMgc3RhbmRhcmQgZm9ybSBzdWJtaXNzaW9uIGFuZCB2YWxpZGF0aW9uIEFQSXMuXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGZvcm0gY29udHJvbHNcbiAqXG4gKiBAYXR0ciB7Ym9vbGVhbn0gaG9yaXpvbnRhbCAtIEVuYWJsZXMgaG9yaXpvbnRhbCBsYXlvdXQgd2l0aCBncmlkXG4gKiBAZmlyZXMgc3VibWl0IC0gRm9yd2FyZGVkIGZyb20gaW50ZXJuYWwgYDxmb3JtPmAgZWxlbWVudFxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtZm9ybScpXG5leHBvcnQgY2xhc3MgUGZWNkZvcm0gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGhvcml6b250YWwgPSBmYWxzZTtcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8Zm9ybSBpZD1cImZvcm1cIlxuICAgICAgICAgICAgQHN1Ym1pdD0ke3RoaXMuI29uU3VibWl0fT5cbiAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgPC9mb3JtPlxuICAgIGA7XG4gIH1cblxuICAjb25TdWJtaXQoZTogU3VibWl0RXZlbnQpIHtcbiAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IFN1Ym1pdEV2ZW50KCdzdWJtaXQnLCB7XG4gICAgICBidWJibGVzOiB0cnVlLFxuICAgICAgY2FuY2VsYWJsZTogdHJ1ZSxcbiAgICAgIHN1Ym1pdHRlcjogZS5zdWJtaXR0ZXIsXG4gICAgfSkpO1xuICB9XG5cbiAgLyoqIFN1Ym1pdHMgdGhlIGZvcm0gcHJvZ3JhbW1hdGljYWxseSAqL1xuICBzdWJtaXQoKSB7XG4gICAgdGhpcy4jZm9ybUVsZW1lbnQ/LnN1Ym1pdCgpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJlcXVlc3RzIGZvcm0gc3VibWlzc2lvbiB3aXRoIHZhbGlkYXRpb25cbiAgICogQHBhcmFtIHN1Ym1pdHRlciAtIE9wdGlvbmFsIHN1Ym1pdHRlciBlbGVtZW50XG4gICAqL1xuICByZXF1ZXN0U3VibWl0KHN1Ym1pdHRlcj86IEhUTUxFbGVtZW50KSB7XG4gICAgcmV0dXJuIHRoaXMuI2Zvcm1FbGVtZW50Py5yZXF1ZXN0U3VibWl0KHN1Ym1pdHRlcik7XG4gIH1cblxuICAvKiogUmVzZXRzIHRoZSBmb3JtIHRvIGRlZmF1bHQgdmFsdWVzICovXG4gIHJlc2V0KCkge1xuICAgIHRoaXMuI2Zvcm1FbGVtZW50Py5yZXNldCgpO1xuICB9XG5cbiAgLyoqIENoZWNrcyBmb3JtIHZhbGlkaXR5ICovXG4gIGNoZWNrVmFsaWRpdHkoKTogYm9vbGVhbiB7XG4gICAgcmV0dXJuIHRoaXMuI2Zvcm1FbGVtZW50Py5jaGVja1ZhbGlkaXR5KCkgPz8gdHJ1ZTtcbiAgfVxuXG4gIC8qKiBSZXBvcnRzIGZvcm0gdmFsaWRpdHkgKHNob3dzIHZhbGlkYXRpb24gbWVzc2FnZXMpICovXG4gIHJlcG9ydFZhbGlkaXR5KCk6IGJvb2xlYW4ge1xuICAgIHJldHVybiB0aGlzLiNmb3JtRWxlbWVudD8ucmVwb3J0VmFsaWRpdHkoKSA/PyB0cnVlO1xuICB9XG5cbiAgZ2V0ICNmb3JtRWxlbWVudCgpOiBIVE1MRm9ybUVsZW1lbnQgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnZm9ybScpIGFzIEhUTUxGb3JtRWxlbWVudCB8IG51bGw7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAncGYtdjYtZm9ybSc6IFBmVjZGb3JtO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxufVxcXFxuXFxcXG4jZm9ybSB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxufVxcXFxuXFxcXG46aG9zdChbaG9yaXpvbnRhbF0pICNmb3JtIHtcXFxcbiAgLS1fZm9ybS1jb2x1bW4tZ2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICBkaXNwbGF5OiBncmlkO1xcXFxuICBncmlkLXRlbXBsYXRlLWNvbHVtbnM6IG1pbm1heCgwLCBhdXRvKSAxZnI7XFxcXG4gIGNvbHVtbi1nYXA6IHZhcigtLV9mb3JtLWNvbHVtbi1nYXApO1xcXFxuICByb3ctZ2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuXFxcXG4gIFxcXFx1MDAyNiA6OnNsb3R0ZWQocGYtdjYtZm9ybS1maWVsZC1ncm91cCkge1xcXFxuICAgIGdyaWQtY29sdW1uOiAxIC8gLTE7XFxcXG4gIH1cXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSx1WkFBeVosQ0FBQztBQUFFLElBQU8scUJBQVE7OztBREFoZTtBQWlCQSx3QkFBQyxjQUFjLFlBQVk7QUFDcEIsSUFBTSxXQUFOLGVBQXVCLGlCQUc1QixtQkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBSGQsSUFBVztBQUFBLEVBQWxDO0FBQUE7QUFBQTtBQUlMLHVCQUFTLGFBQWEsa0JBQXRCLGdCQUFzQixTQUF0QjtBQUFBO0FBQUEsRUFFQSxTQUFTO0FBQ1AsV0FBTztBQUFBO0FBQUEsc0JBRVcsc0JBQUssaUNBQVM7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQUlsQztBQUFBO0FBQUEsRUFXQSxTQUFTO0FBQ1AsdUJBQUssdUNBQWMsT0FBTztBQUFBLEVBQzVCO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQU1BLGNBQWMsV0FBeUI7QUFDckMsV0FBTyxtQkFBSyx1Q0FBYyxjQUFjLFNBQVM7QUFBQSxFQUNuRDtBQUFBO0FBQUEsRUFHQSxRQUFRO0FBQ04sdUJBQUssdUNBQWMsTUFBTTtBQUFBLEVBQzNCO0FBQUE7QUFBQSxFQUdBLGdCQUF5QjtBQUN2QixXQUFPLG1CQUFLLHVDQUFjLGNBQWMsS0FBSztBQUFBLEVBQy9DO0FBQUE7QUFBQSxFQUdBLGlCQUEwQjtBQUN4QixXQUFPLG1CQUFLLHVDQUFjLGVBQWUsS0FBSztBQUFBLEVBQ2hEO0FBS0Y7QUF0RE87QUFJSTtBQUpKO0FBZUwsY0FBUyxTQUFDLEdBQWdCO0FBQ3hCLE9BQUssY0FBYyxJQUFJLFlBQVksVUFBVTtBQUFBLElBQzNDLFNBQVM7QUFBQSxJQUNULFlBQVk7QUFBQSxJQUNaLFdBQVcsRUFBRTtBQUFBLEVBQ2YsQ0FBQyxDQUFDO0FBQ0o7QUE4Qkksa0JBQVksV0FBMkI7QUFDekMsU0FBTyxLQUFLLFlBQVksZUFBZSxNQUFNO0FBQy9DO0FBakRBLDRCQUFTLGNBRFQsaUJBSFcsVUFJRjtBQUpFLFdBQU4sd0NBRFAsc0JBQ2E7QUFDWCxjQURXLFVBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
