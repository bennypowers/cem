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

// elements/cem-pf-v6-switch/cem-pf-v6-switch.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";
import { live } from "/__cem/vendor/lit/directives/live.js";

// lit-css:elements/cem-pf-v6-switch/cem-pf-v6-switch.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: inline-flex;\\n}\\n\\n#switch-label {\\n  display: inline-flex;\\n  align-items: center;\\n  gap: 0.5rem;\\n  cursor: pointer;\\n  user-select: none;\\n}\\n\\n#switch-input {\\n  position: absolute;\\n  opacity: 0;\\n  pointer-events: none;\\n}\\n\\n#switch-toggle {\\n  position: relative;\\n  display: inline-block;\\n  width: calc(calc(var(--pf-t--global--font--size--sm) * var(--pf-t--global--font--line-height--body)) + 0.125rem + calc(var(--pf-t--global--font--size--sm) - 0.125rem));\\n  height: calc(var(--pf-t--global--font--size--sm) * var(--pf-t--global--font--line-height--body));\\n  background-color: var(--pf-t--global--background--color--control--default);\\n  border-radius: var(--pf-t--global--border--radius--pill);\\n\\n  \\u0026::before {\\n    position: absolute;\\n    inset-block-start: 50%;\\n    inset-inline-start: calc((calc(var(--pf-t--global--font--size--sm) * var(--pf-t--global--font--line-height--body)) - calc(var(--pf-t--global--font--size--sm) - 0.125rem)) / 2);\\n    display: block;\\n    width: calc(var(--pf-t--global--font--size--sm) - 0.125rem);\\n    height: calc(var(--pf-t--global--font--size--sm) - 0.125rem);\\n    content: \\"\\";\\n    background-color: var(--pf-t--global--icon--color--subtle);\\n    border: var(--pf-t--global--border--width--regular) solid transparent;\\n    border-radius: var(--pf-t--global--border--radius--pill);\\n    transition: transform var(--pf-t--global--motion--duration--md) var(--pf-t--global--motion--timing-function--default),\\n                background-color var(--pf-t--global--motion--duration--md) var(--pf-t--global--motion--timing-function--default);\\n    transform: translateY(-50%);\\n  }\\n\\n  \\u0026::after {\\n    position: absolute;\\n    inset: 0;\\n    content: \\"\\";\\n    border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n    border-radius: var(--pf-t--global--border--radius--pill);\\n  }\\n}\\n\\n.switch-toggle-icon {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-block-end: 0;\\n  inset-inline-start: calc(var(--pf-t--global--font--size--sm) * 0.4);\\n  display: none;\\n  align-items: center;\\n  font-size: calc(var(--pf-t--global--font--size--sm) * 0.625);\\n  color: var(--pf-t--global--icon--color--on-brand--default);\\n\\n  \\u0026 svg {\\n    width: 1em;\\n    height: 1em;\\n  }\\n}\\n\\n#switch-input:checked + #switch-toggle {\\n  background-color: var(--pf-t--global--color--brand--default);\\n\\n  \\u0026::before {\\n    transform: translate(calc(100% + 0.125rem), -50%);\\n    background-color: var(--pf-t--global--icon--color--inverse);\\n  }\\n\\n  \\u0026::after {\\n    border-color: transparent;\\n  }\\n\\n  \\u0026 .switch-toggle-icon {\\n    display: flex;\\n  }\\n}\\n\\n#switch-input:not(:checked) + #switch-toggle .switch-toggle-icon {\\n  display: none;\\n}\\n\\n#switch-input:focus-visible + #switch-toggle {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: 2px;\\n}\\n\\n#switch-input:disabled + #switch-toggle {\\n  background-color: var(--pf-t--global--background--color--disabled--default);\\n  cursor: not-allowed;\\n\\n  \\u0026::before {\\n    background-color: var(--pf-t--global--text--color--disabled);\\n  }\\n\\n  \\u0026::after {\\n    border-color: var(--pf-t--global--border--color--disabled);\\n  }\\n\\n  \\u0026 .switch-toggle-icon {\\n    color: var(--pf-t--global--icon--color--disabled);\\n  }\\n}\\n\\n#switch-input:disabled ~ #switch-text {\\n  color: var(--pf-t--global--text--color--disabled);\\n  cursor: not-allowed;\\n}\\n\\n:host([disabled]) #switch-label {\\n  cursor: not-allowed;\\n}\\n\\n#switch-text {\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: 0.875rem;\\n  line-height: 1.5;\\n}\\n\\n#switch-input:checked ~ #switch-text {\\n  color: var(--pf-t--global--text--color--regular);\\n}\\n\\n:host([label-position=\\"start\\"]) #switch-label {\\n  flex-direction: row-reverse;\\n  justify-content: flex-end;\\n}\\n"'));
var cem_pf_v6_switch_default = s;

// elements/cem-pf-v6-switch/cem-pf-v6-switch.ts
var _labelPosition_dec, _disabled_dec, _checked_dec, _a, _PfV6Switch_decorators, _internals, _init, _checked, _disabled, _labelPosition, _PfV6Switch_instances, onInput_fn;
_PfV6Switch_decorators = [customElement("cem-pf-v6-switch")];
var PfV6Switch = class extends (_a = LitElement, _checked_dec = [property({ type: Boolean, reflect: true })], _disabled_dec = [property({ type: Boolean, reflect: true })], _labelPosition_dec = [property({ reflect: true, attribute: "label-position" })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Switch_instances);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _checked, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _disabled, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateAdd(this, _labelPosition, __runInitializers(_init, 16, this)), __runInitializers(_init, 19, this);
  }
  render() {
    return html`
      <label id="switch-label">
        <input type="checkbox"
               id="switch-input"
               role="switch"
               .checked=${live(this.checked)}
               ?disabled=${this.disabled}
               @input=${__privateMethod(this, _PfV6Switch_instances, onInput_fn)}>
        <span id="switch-toggle"
              aria-hidden="true">
          <span class="switch-toggle-icon">
            <svg width="10"
                 height="10"
                 viewBox="0 0 512 512"
                 fill="currentColor">
              <path d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"/>
            </svg>
          </span>
        </span>
        <span id="switch-text"><slot></slot></span>
      </label>
    `;
  }
};
_init = __decoratorStart(_a);
_internals = new WeakMap();
_checked = new WeakMap();
_disabled = new WeakMap();
_labelPosition = new WeakMap();
_PfV6Switch_instances = new WeakSet();
onInput_fn = function() {
  const input = this.shadowRoot?.getElementById("switch-input");
  if (input) {
    this.checked = input.checked;
    __privateGet(this, _internals).setFormValue(input.checked ? "on" : null);
  }
};
__decorateElement(_init, 4, "checked", _checked_dec, PfV6Switch, _checked);
__decorateElement(_init, 4, "disabled", _disabled_dec, PfV6Switch, _disabled);
__decorateElement(_init, 4, "labelPosition", _labelPosition_dec, PfV6Switch, _labelPosition);
PfV6Switch = __decorateElement(_init, 0, "PfV6Switch", _PfV6Switch_decorators, PfV6Switch);
__publicField(PfV6Switch, "formAssociated", true);
__publicField(PfV6Switch, "styles", cem_pf_v6_switch_default);
__runInitializers(_init, 1, PfV6Switch);
export {
  PfV6Switch
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXN3aXRjaC9jZW0tcGYtdjYtc3dpdGNoLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LXN3aXRjaC9jZW0tcGYtdjYtc3dpdGNoLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBsaXZlIH0gZnJvbSAnbGl0L2RpcmVjdGl2ZXMvbGl2ZS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtc3dpdGNoLmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBTd2l0Y2hcbiAqXG4gKiBBIHRvZ2dsZSBzd2l0Y2ggZm9ybSBjb250cm9sLlxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBzd2l0Y2ggbGFiZWwgdGV4dFxuICpcbiAqIEBmaXJlcyBpbnB1dCAtIEZpcmVkIHdoZW4gdGhlIHN3aXRjaCBpcyB0b2dnbGVkXG4gKiBAZmlyZXMgY2hhbmdlIC0gRmlyZWQgd2hlbiB0aGUgc3dpdGNoIGlzIHRvZ2dsZWRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1wZi12Ni1zd2l0Y2gnKVxuZXhwb3J0IGNsYXNzIFBmVjZTd2l0Y2ggZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHJlYWRvbmx5IGZvcm1Bc3NvY2lhdGVkID0gdHJ1ZTtcblxuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gICNpbnRlcm5hbHMgPSB0aGlzLmF0dGFjaEludGVybmFscygpO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgY2hlY2tlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZGlzYWJsZWQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlLCBhdHRyaWJ1dGU6ICdsYWJlbC1wb3NpdGlvbicgfSlcbiAgYWNjZXNzb3IgbGFiZWxQb3NpdGlvbj86ICdzdGFydCc7XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGxhYmVsIGlkPVwic3dpdGNoLWxhYmVsXCI+XG4gICAgICAgIDxpbnB1dCB0eXBlPVwiY2hlY2tib3hcIlxuICAgICAgICAgICAgICAgaWQ9XCJzd2l0Y2gtaW5wdXRcIlxuICAgICAgICAgICAgICAgcm9sZT1cInN3aXRjaFwiXG4gICAgICAgICAgICAgICAuY2hlY2tlZD0ke2xpdmUodGhpcy5jaGVja2VkKX1cbiAgICAgICAgICAgICAgID9kaXNhYmxlZD0ke3RoaXMuZGlzYWJsZWR9XG4gICAgICAgICAgICAgICBAaW5wdXQ9JHt0aGlzLiNvbklucHV0fT5cbiAgICAgICAgPHNwYW4gaWQ9XCJzd2l0Y2gtdG9nZ2xlXCJcbiAgICAgICAgICAgICAgYXJpYS1oaWRkZW49XCJ0cnVlXCI+XG4gICAgICAgICAgPHNwYW4gY2xhc3M9XCJzd2l0Y2gtdG9nZ2xlLWljb25cIj5cbiAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxMFwiXG4gICAgICAgICAgICAgICAgIGhlaWdodD1cIjEwXCJcbiAgICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCA1MTIgNTEyXCJcbiAgICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICA8cGF0aCBkPVwiTTE3My44OTggNDM5LjQwNGwtMTY2LjQtMTY2LjRjLTkuOTk3LTkuOTk3LTkuOTk3LTI2LjIwNiAwLTM2LjIwNGwzNi4yMDMtMzYuMjA0YzkuOTk3LTkuOTk4IDI2LjIwNy05Ljk5OCAzNi4yMDQgMEwxOTIgMzEyLjY5IDQzMi4wOTUgNzIuNTk2YzkuOTk3LTkuOTk3IDI2LjIwNy05Ljk5NyAzNi4yMDQgMGwzNi4yMDMgMzYuMjA0YzkuOTk3IDkuOTk3IDkuOTk3IDI2LjIwNiAwIDM2LjIwNGwtMjk0LjQgMjk0LjQwMWMtOS45OTggOS45OTctMjYuMjA3IDkuOTk3LTM2LjIwNC0uMDAxelwiLz5cbiAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgIDwvc3Bhbj5cbiAgICAgICAgPC9zcGFuPlxuICAgICAgICA8c3BhbiBpZD1cInN3aXRjaC10ZXh0XCI+PHNsb3Q+PC9zbG90Pjwvc3Bhbj5cbiAgICAgIDwvbGFiZWw+XG4gICAgYDtcbiAgfVxuXG4gICNvbklucHV0KCkge1xuICAgIGNvbnN0IGlucHV0ID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnc3dpdGNoLWlucHV0JykgYXMgSFRNTElucHV0RWxlbWVudCB8IG51bGw7XG4gICAgaWYgKGlucHV0KSB7XG4gICAgICB0aGlzLmNoZWNrZWQgPSBpbnB1dC5jaGVja2VkO1xuICAgICAgdGhpcy4jaW50ZXJuYWxzLnNldEZvcm1WYWx1ZShpbnB1dC5jaGVja2VkID8gJ29uJyA6IG51bGwpO1xuICAgIH1cbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcGYtdjYtc3dpdGNoJzogUGZWNlN3aXRjaDtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBpbmxpbmUtZmxleDtcXFxcbn1cXFxcblxcXFxuI3N3aXRjaC1sYWJlbCB7XFxcXG4gIGRpc3BsYXk6IGlubGluZS1mbGV4O1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBnYXA6IDAuNXJlbTtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcbn1cXFxcblxcXFxuI3N3aXRjaC1pbnB1dCB7XFxcXG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgb3BhY2l0eTogMDtcXFxcbiAgcG9pbnRlci1ldmVudHM6IG5vbmU7XFxcXG59XFxcXG5cXFxcbiNzd2l0Y2gtdG9nZ2xlIHtcXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxuICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XFxcXG4gIHdpZHRoOiBjYWxjKGNhbGModmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1zbSkgKiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLWxpbmUtaGVpZ2h0LS1ib2R5KSkgKyAwLjEyNXJlbSArIGNhbGModmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1zbSkgLSAwLjEyNXJlbSkpO1xcXFxuICBoZWlnaHQ6IGNhbGModmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1zbSkgKiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLWxpbmUtaGVpZ2h0LS1ib2R5KSk7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWNvbnRyb2wtLWRlZmF1bHQpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1waWxsKTtcXFxcblxcXFxuICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICAgIGluc2V0LWJsb2NrLXN0YXJ0OiA1MCU7XFxcXG4gICAgaW5zZXQtaW5saW5lLXN0YXJ0OiBjYWxjKChjYWxjKHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tc20pICogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSkpIC0gY2FsYyh2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLXNtKSAtIDAuMTI1cmVtKSkgLyAyKTtcXFxcbiAgICBkaXNwbGF5OiBibG9jaztcXFxcbiAgICB3aWR0aDogY2FsYyh2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLXNtKSAtIDAuMTI1cmVtKTtcXFxcbiAgICBoZWlnaHQ6IGNhbGModmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1zbSkgLSAwLjEyNXJlbSk7XFxcXG4gICAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcIjtcXFxcbiAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWljb24tLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIGJvcmRlcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1yZWd1bGFyKSBzb2xpZCB0cmFuc3BhcmVudDtcXFxcbiAgICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1waWxsKTtcXFxcbiAgICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1tZCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCksXFxcXG4gICAgICAgICAgICAgICAgYmFja2dyb3VuZC1jb2xvciB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLW1kKSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KTtcXFxcbiAgICB0cmFuc2Zvcm06IHRyYW5zbGF0ZVkoLTUwJSk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6OmFmdGVyIHtcXFxcbiAgICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gICAgaW5zZXQ6IDA7XFxcXG4gICAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcIjtcXFxcbiAgICBib3JkZXI6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tcmVndWxhcikgc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1waWxsKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4uc3dpdGNoLXRvZ2dsZS1pY29uIHtcXFxcbiAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICBpbnNldC1ibG9jay1zdGFydDogMDtcXFxcbiAgaW5zZXQtYmxvY2stZW5kOiAwO1xcXFxuICBpbnNldC1pbmxpbmUtc3RhcnQ6IGNhbGModmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1zbSkgKiAwLjQpO1xcXFxuICBkaXNwbGF5OiBub25lO1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBmb250LXNpemU6IGNhbGModmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1zbSkgKiAwLjYyNSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWljb24tLWNvbG9yLS1vbi1icmFuZC0tZGVmYXVsdCk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2IHN2ZyB7XFxcXG4gICAgd2lkdGg6IDFlbTtcXFxcbiAgICBoZWlnaHQ6IDFlbTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jc3dpdGNoLWlucHV0OmNoZWNrZWQgKyAjc3dpdGNoLXRvZ2dsZSB7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcblxcXFxuICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgdHJhbnNmb3JtOiB0cmFuc2xhdGUoY2FsYygxMDAlICsgMC4xMjVyZW0pLCAtNTAlKTtcXFxcbiAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWljb24tLWNvbG9yLS1pbnZlcnNlKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjo6YWZ0ZXIge1xcXFxuICAgIGJvcmRlci1jb2xvcjogdHJhbnNwYXJlbnQ7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYgLnN3aXRjaC10b2dnbGUtaWNvbiB7XFxcXG4gICAgZGlzcGxheTogZmxleDtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jc3dpdGNoLWlucHV0Om5vdCg6Y2hlY2tlZCkgKyAjc3dpdGNoLXRvZ2dsZSAuc3dpdGNoLXRvZ2dsZS1pY29uIHtcXFxcbiAgZGlzcGxheTogbm9uZTtcXFxcbn1cXFxcblxcXFxuI3N3aXRjaC1pbnB1dDpmb2N1cy12aXNpYmxlICsgI3N3aXRjaC10b2dnbGUge1xcXFxuICBvdXRsaW5lOiAycHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQpO1xcXFxuICBvdXRsaW5lLW9mZnNldDogMnB4O1xcXFxufVxcXFxuXFxcXG4jc3dpdGNoLWlucHV0OmRpc2FibGVkICsgI3N3aXRjaC10b2dnbGUge1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1kaXNhYmxlZC0tZGVmYXVsdCk7XFxcXG4gIGN1cnNvcjogbm90LWFsbG93ZWQ7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OjpiZWZvcmUge1xcXFxuICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLWRpc2FibGVkKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjo6YWZ0ZXIge1xcXFxuICAgIGJvcmRlci1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kaXNhYmxlZCk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYgLnN3aXRjaC10b2dnbGUtaWNvbiB7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLWRpc2FibGVkKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jc3dpdGNoLWlucHV0OmRpc2FibGVkIH4gI3N3aXRjaC10ZXh0IHtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLWRpc2FibGVkKTtcXFxcbiAgY3Vyc29yOiBub3QtYWxsb3dlZDtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2Rpc2FibGVkXSkgI3N3aXRjaC1sYWJlbCB7XFxcXG4gIGN1cnNvcjogbm90LWFsbG93ZWQ7XFxcXG59XFxcXG5cXFxcbiNzd2l0Y2gtdGV4dCB7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxufVxcXFxuXFxcXG4jc3dpdGNoLWlucHV0OmNoZWNrZWQgfiAjc3dpdGNoLXRleHQge1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtsYWJlbC1wb3NpdGlvbj1cXFxcXFxcInN0YXJ0XFxcXFxcXCJdKSAjc3dpdGNoLWxhYmVsIHtcXFxcbiAgZmxleC1kaXJlY3Rpb246IHJvdy1yZXZlcnNlO1xcXFxuICBqdXN0aWZ5LWNvbnRlbnQ6IGZsZXgtZW5kO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjtBQUN6QixTQUFTLFlBQVk7OztBQ0hyQixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSx3OEhBQWc5SCxDQUFDO0FBQUUsSUFBTywyQkFBUTs7O0FEQXZoSTtBQWlCQSwwQkFBQyxjQUFjLGtCQUFrQjtBQUMxQixJQUFNLGFBQU4sZUFBeUIsaUJBTzlCLGdCQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFHMUMsaUJBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxzQkFBQyxTQUFTLEVBQUUsU0FBUyxNQUFNLFdBQVcsaUJBQWlCLENBQUMsSUFiMUIsSUFBVztBQUFBLEVBQXBDO0FBQUE7QUFBQTtBQUtMLG1DQUFhLEtBQUssZ0JBQWdCO0FBR2xDLHVCQUFTLFVBQVUsa0JBQW5CLGdCQUFtQixTQUFuQjtBQUdBLHVCQUFTLFdBQVcsa0JBQXBCLGlCQUFvQixTQUFwQjtBQUdBLHVCQUFTLGdCQUFUO0FBQUE7QUFBQSxFQUVBLFNBQVM7QUFDUCxXQUFPO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSwwQkFLZSxLQUFLLEtBQUssT0FBTyxDQUFDO0FBQUEsMkJBQ2pCLEtBQUssUUFBUTtBQUFBLHdCQUNoQixzQkFBSyxrQ0FBUTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQWVuQztBQVNGO0FBaERPO0FBS0w7QUFHUztBQUdBO0FBR0E7QUFkSjtBQXlDTCxhQUFRLFdBQUc7QUFDVCxRQUFNLFFBQVEsS0FBSyxZQUFZLGVBQWUsY0FBYztBQUM1RCxNQUFJLE9BQU87QUFDVCxTQUFLLFVBQVUsTUFBTTtBQUNyQix1QkFBSyxZQUFXLGFBQWEsTUFBTSxVQUFVLE9BQU8sSUFBSTtBQUFBLEVBQzFEO0FBQ0Y7QUF2Q0EsNEJBQVMsV0FEVCxjQVBXLFlBUUY7QUFHVCw0QkFBUyxZQURULGVBVlcsWUFXRjtBQUdULDRCQUFTLGlCQURULG9CQWJXLFlBY0Y7QUFkRSxhQUFOLDBDQURQLHdCQUNhO0FBQ1gsY0FEVyxZQUNLLGtCQUFpQjtBQUVqQyxjQUhXLFlBR0osVUFBUztBQUhYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
