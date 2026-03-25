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

// elements/pf-v6-dropdown/pf-v6-dropdown.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-dropdown/pf-v6-dropdown.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: inline-block;\\n  position: relative;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n#toggle {\\n  --pf-v6-c-menu-toggle--Gap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --pf-v6-c-menu-toggle--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--default);\\n  --pf-v6-c-menu-toggle--PaddingInlineEnd: var(--pf-t--global--spacer--control--horizontal--default);\\n  --pf-v6-c-menu-toggle--PaddingBlockEnd: var(--pf-t--global--spacer--control--vertical--default);\\n  --pf-v6-c-menu-toggle--PaddingInlineStart: var(--pf-t--global--spacer--control--horizontal--default);\\n  --pf-v6-c-menu-toggle--MinWidth: calc(var(--pf-v6-c-menu-toggle--FontSize) * var(--pf-v6-c-menu-toggle--LineHeight) + var(--pf-v6-c-menu-toggle--PaddingBlockStart) + var(--pf-v6-c-menu-toggle--PaddingBlockEnd));\\n  --pf-v6-c-menu-toggle--FontSize: var(--pf-t--global--font--size--body--default);\\n  --pf-v6-c-menu-toggle--Color: var(--pf-t--global--text--color--regular);\\n  --pf-v6-c-menu-toggle--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --pf-v6-c-menu-toggle--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --pf-v6-c-menu-toggle--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --pf-v6-c-menu-toggle--BorderColor: var(--pf-t--global--border--color--default);\\n  --pf-v6-c-menu-toggle--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --pf-v6-c-menu-toggle--border--ZIndex: var(--pf-t--global--z-index--xs);\\n  --pf-v6-c-menu-toggle--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --pf-v6-c-menu-toggle--TransitionDuration: var(--pf-t--global--motion--duration--fade--short);\\n  --pf-v6-c-menu-toggle--TransitionProperty: color, background-color, border-width, border-color;\\n  --pf-v6-c-menu-toggle--hover--Color: var(--pf-t--global--text--color--regular);\\n  --pf-v6-c-menu-toggle--hover--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --pf-v6-c-menu-toggle--hover--BorderWidth: var(--pf-t--global--border--width--control--hover);\\n  --pf-v6-c-menu-toggle--hover--BorderColor: var(--pf-t--global--border--color--hover);\\n  --pf-v6-c-menu-toggle--expanded--Color: var(--pf-t--global--text--color--regular);\\n  --pf-v6-c-menu-toggle--expanded--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --pf-v6-c-menu-toggle--expanded--BorderWidth: var(--pf-t--global--border--width--control--clicked);\\n  --pf-v6-c-menu-toggle--expanded--BorderColor: var(--pf-t--global--border--color--clicked);\\n  --pf-v6-c-menu-toggle__toggle-icon--Color: var(--pf-t--global--icon--color--regular);\\n\\n  --pf-v6-c-button--Color: var(--pf-v6-c-menu-toggle--Color);\\n  --pf-v6-c-button--BackgroundColor: var(--pf-v6-c-menu-toggle--BackgroundColor);\\n  --pf-v6-c-button--BorderRadius: var(--pf-v6-c-menu-toggle--BorderRadius);\\n  --pf-v6-c-button--PaddingBlockStart: var(--pf-v6-c-menu-toggle--PaddingBlockStart);\\n  --pf-v6-c-button--PaddingBlockEnd: var(--pf-v6-c-menu-toggle--PaddingBlockEnd);\\n  --pf-v6-c-button--PaddingInlineStart: var(--pf-v6-c-menu-toggle--PaddingInlineStart);\\n  --pf-v6-c-button--PaddingInlineEnd: var(--pf-v6-c-menu-toggle--PaddingInlineEnd);\\n  --pf-v6-c-button--FontSize: var(--pf-v6-c-menu-toggle--FontSize);\\n  --pf-v6-c-button--LineHeight: var(--pf-v6-c-menu-toggle--LineHeight);\\n  --pf-v6-c-button--MinWidth: var(--pf-v6-c-menu-toggle--MinWidth);\\n\\n  \\u0026::after {\\n    position: absolute;\\n    inset: 0;\\n    z-index: var(--pf-v6-c-menu-toggle--border--ZIndex);\\n    pointer-events: none;\\n    content: \\"\\";\\n    border: var(--pf-v6-c-menu-toggle--BorderWidth) solid var(--pf-v6-c-menu-toggle--BorderColor);\\n    border-radius: var(--pf-v6-c-menu-toggle--BorderRadius);\\n    transition: inherit;\\n  }\\n\\n  \\u0026:hover:not([disabled]) {\\n    --pf-v6-c-menu-toggle--Color: var(--pf-v6-c-menu-toggle--hover--Color);\\n    --pf-v6-c-menu-toggle--BackgroundColor: var(--pf-v6-c-menu-toggle--hover--BackgroundColor);\\n    --pf-v6-c-menu-toggle--BorderWidth: var(--pf-v6-c-menu-toggle--hover--BorderWidth);\\n    --pf-v6-c-menu-toggle--BorderColor: var(--pf-v6-c-menu-toggle--hover--BorderColor);\\n  }\\n}\\n\\n:host([expanded]) #toggle {\\n  --pf-v6-c-menu-toggle--Color: var(--pf-v6-c-menu-toggle--expanded--Color);\\n  --pf-v6-c-menu-toggle--BackgroundColor: var(--pf-v6-c-menu-toggle--expanded--BackgroundColor);\\n  --pf-v6-c-menu-toggle--BorderWidth: var(--pf-v6-c-menu-toggle--expanded--BorderWidth);\\n  --pf-v6-c-menu-toggle--BorderColor: var(--pf-v6-c-menu-toggle--expanded--BorderColor);\\n}\\n\\n#toggle svg[slot=\\"icon-end\\"] {\\n  width: 0.75em;\\n  height: 0.75em;\\n  fill: currentColor;\\n  color: var(--pf-v6-c-menu-toggle__toggle-icon--Color);\\n  transition: transform 150ms ease-in-out;\\n}\\n\\n:host([expanded]) #toggle svg[slot=\\"icon-end\\"] {\\n  transform: rotate(180deg);\\n}\\n\\n#menu-container {\\n  position: absolute;\\n  top: 100%;\\n  left: 0;\\n  margin-top: var(--pf-t--global--spacer--xs);\\n  z-index: var(--pf-t--global--z-index--sm, 1000);\\n}\\n\\n:host([disabled]) {\\n  pointer-events: none;\\n  opacity: 0.6;\\n}\\n"'));
var pf_v6_dropdown_default = s;

// elements/pf-v6-dropdown/pf-v6-dropdown.ts
import "../pf-v6-button/pf-v6-button.js";
import "../pf-v6-menu/pf-v6-menu.js";
import "../pf-v6-menu-item/pf-v6-menu-item.js";
var _label_dec, _disabled_dec, _expanded_dec, _a, _PfV6Dropdown_decorators, _instances, _init, _expanded, _disabled, _label, _onKeydown, _PfV6Dropdown_instances, onToggleClick_fn;
_PfV6Dropdown_decorators = [customElement("pf-v6-dropdown")];
var _PfV6Dropdown = class _PfV6Dropdown extends (_a = LitElement, _expanded_dec = [property({ type: Boolean, reflect: true })], _disabled_dec = [property({ type: Boolean, reflect: true })], _label_dec = [property()], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Dropdown_instances);
    __privateAdd(this, _expanded, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _disabled, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateAdd(this, _label, __runInitializers(_init, 16, this, "")), __runInitializers(_init, 19, this);
    __privateAdd(this, _onKeydown, (event) => {
      if (event.key === "Escape" && this.expanded) {
        event.preventDefault();
        this.collapse();
        this.shadowRoot?.getElementById("toggle")?.focus();
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    __privateGet(_PfV6Dropdown, _instances).add(this);
    this.addEventListener("keydown", __privateGet(this, _onKeydown));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(_PfV6Dropdown, _instances).delete(this);
    this.removeEventListener("keydown", __privateGet(this, _onKeydown));
  }
  render() {
    return html`
      <pf-v6-button id="toggle"
                     variant="tertiary"
                     aria-haspopup="true"
                     aria-expanded="${this.expanded}"
                     ?disabled=${this.disabled}
                     @click=${__privateMethod(this, _PfV6Dropdown_instances, onToggleClick_fn)}>
        <slot name="toggle-text">Toggle</slot>
        <svg slot="icon-end"
             viewBox="0 0 320 512"
             aria-hidden="true">
          <path d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z"/>
        </svg>
      </pf-v6-button>

      <div id="menu-container"
           ?hidden=${!this.expanded}>
        <pf-v6-menu id="menu"
                     label=${this.label}>
          <slot></slot>
        </pf-v6-menu>
      </div>
    `;
  }
  updated(changed) {
    if (changed.has("expanded")) {
      this.dispatchEvent(new Event(this.expanded ? "expand" : "collapse", { bubbles: true }));
      if (this.expanded) {
        requestAnimationFrame(() => {
          const menu = this.shadowRoot?.getElementById("menu");
          menu?.focusFirstItem();
        });
      }
    }
  }
  /** Toggle expanded state */
  toggle() {
    this.expanded = !this.expanded;
  }
  /** Expand the dropdown */
  expand() {
    this.expanded = true;
  }
  /** Collapse the dropdown */
  collapse() {
    this.expanded = false;
  }
};
_init = __decoratorStart(_a);
_instances = new WeakMap();
_expanded = new WeakMap();
_disabled = new WeakMap();
_label = new WeakMap();
_onKeydown = new WeakMap();
_PfV6Dropdown_instances = new WeakSet();
onToggleClick_fn = function() {
  if (this.disabled) return;
  this.toggle();
};
__decorateElement(_init, 4, "expanded", _expanded_dec, _PfV6Dropdown, _expanded);
__decorateElement(_init, 4, "disabled", _disabled_dec, _PfV6Dropdown, _disabled);
__decorateElement(_init, 4, "label", _label_dec, _PfV6Dropdown, _label);
_PfV6Dropdown = __decorateElement(_init, 0, "PfV6Dropdown", _PfV6Dropdown_decorators, _PfV6Dropdown);
__privateAdd(_PfV6Dropdown, _instances, /* @__PURE__ */ new Set());
document?.addEventListener?.("click", (event) => {
  for (const instance of __privateGet(_PfV6Dropdown, _instances)) {
    if (instance.expanded && !event.composedPath().includes(instance)) {
      instance.collapse();
    }
  }
});
__publicField(_PfV6Dropdown, "styles", pf_v6_dropdown_default);
__runInitializers(_init, 1, _PfV6Dropdown);
var PfV6Dropdown = _PfV6Dropdown;
export {
  PfV6Dropdown
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtZHJvcGRvd24vcGYtdjYtZHJvcGRvd24udHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9wZi12Ni1kcm9wZG93bi9wZi12Ni1kcm9wZG93bi5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtZHJvcGRvd24uY3NzJztcblxuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbWVudS9wZi12Ni1tZW51LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbWVudS1pdGVtL3BmLXY2LW1lbnUtaXRlbS5qcyc7XG5cbmltcG9ydCB0eXBlIHsgUGZWNk1lbnUgfSBmcm9tICcuLi9wZi12Ni1tZW51L3BmLXY2LW1lbnUuanMnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgRHJvcGRvd24gY29tcG9uZW50XG4gKlxuICogQGZpcmVzIHtFdmVudH0gZXhwYW5kIC0gRmlyZXMgd2hlbiBkcm9wZG93biBvcGVuc1xuICogQGZpcmVzIHtFdmVudH0gY29sbGFwc2UgLSBGaXJlcyB3aGVuIGRyb3Bkb3duIGNsb3Nlc1xuICpcbiAqIEBzbG90IHRvZ2dsZS10ZXh0IC0gVGV4dCBjb250ZW50IGZvciB0b2dnbGUgYnV0dG9uXG4gKiBAc2xvdCAtIE1lbnUgaXRlbXMgKHBmLXY2LW1lbnUtaXRlbSBlbGVtZW50cylcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ3BmLXY2LWRyb3Bkb3duJylcbmV4cG9ydCBjbGFzcyBQZlY2RHJvcGRvd24gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHJlYWRvbmx5ICNpbnN0YW5jZXMgPSBuZXcgU2V0PFBmVjZEcm9wZG93bj4oKTtcblxuICBzdGF0aWMge1xuICAgIGRvY3VtZW50Py5hZGRFdmVudExpc3RlbmVyPy4oJ2NsaWNrJywgKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgICAgZm9yIChjb25zdCBpbnN0YW5jZSBvZiBQZlY2RHJvcGRvd24uI2luc3RhbmNlcykge1xuICAgICAgICBpZiAoaW5zdGFuY2UuZXhwYW5kZWQgJiYgIWV2ZW50LmNvbXBvc2VkUGF0aCgpLmluY2x1ZGVzKGluc3RhbmNlKSkge1xuICAgICAgICAgIGluc3RhbmNlLmNvbGxhcHNlKCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBleHBhbmRlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZGlzYWJsZWQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBsYWJlbCA9ICcnO1xuXG4gIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmNvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgUGZWNkRyb3Bkb3duLiNpbnN0YW5jZXMuYWRkKHRoaXMpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI29uS2V5ZG93bik7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIFBmVjZEcm9wZG93bi4jaW5zdGFuY2VzLmRlbGV0ZSh0aGlzKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNvbktleWRvd24pO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cInRvZ2dsZVwiXG4gICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgYXJpYS1oYXNwb3B1cD1cInRydWVcIlxuICAgICAgICAgICAgICAgICAgICAgYXJpYS1leHBhbmRlZD1cIiR7dGhpcy5leHBhbmRlZH1cIlxuICAgICAgICAgICAgICAgICAgICAgP2Rpc2FibGVkPSR7dGhpcy5kaXNhYmxlZH1cbiAgICAgICAgICAgICAgICAgICAgIEBjbGljaz0ke3RoaXMuI29uVG9nZ2xlQ2xpY2t9PlxuICAgICAgICA8c2xvdCBuYW1lPVwidG9nZ2xlLXRleHRcIj5Ub2dnbGU8L3Nsb3Q+XG4gICAgICAgIDxzdmcgc2xvdD1cImljb24tZW5kXCJcbiAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDMyMCA1MTJcIlxuICAgICAgICAgICAgIGFyaWEtaGlkZGVuPVwidHJ1ZVwiPlxuICAgICAgICAgIDxwYXRoIGQ9XCJNMzEuMyAxOTJoMjU3LjNjMTcuOCAwIDI2LjcgMjEuNSAxNC4xIDM0LjFMMTc0LjEgMzU0LjhjLTcuOCA3LjgtMjAuNSA3LjgtMjguMyAwTDE3LjIgMjI2LjFDNC42IDIxMy41IDEzLjUgMTkyIDMxLjMgMTkyelwiLz5cbiAgICAgICAgPC9zdmc+XG4gICAgICA8L3BmLXY2LWJ1dHRvbj5cblxuICAgICAgPGRpdiBpZD1cIm1lbnUtY29udGFpbmVyXCJcbiAgICAgICAgICAgP2hpZGRlbj0keyF0aGlzLmV4cGFuZGVkfT5cbiAgICAgICAgPHBmLXY2LW1lbnUgaWQ9XCJtZW51XCJcbiAgICAgICAgICAgICAgICAgICAgIGxhYmVsPSR7dGhpcy5sYWJlbH0+XG4gICAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgICA8L3BmLXY2LW1lbnU+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygnZXhwYW5kZWQnKSkge1xuICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBFdmVudCh0aGlzLmV4cGFuZGVkID8gJ2V4cGFuZCcgOiAnY29sbGFwc2UnLCB7IGJ1YmJsZXM6IHRydWUgfSkpO1xuICAgICAgaWYgKHRoaXMuZXhwYW5kZWQpIHtcbiAgICAgICAgcmVxdWVzdEFuaW1hdGlvbkZyYW1lKCgpID0+IHtcbiAgICAgICAgICBjb25zdCBtZW51ID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnbWVudScpIGFzIFBmVjZNZW51IHwgbnVsbDtcbiAgICAgICAgICBtZW51Py5mb2N1c0ZpcnN0SXRlbSgpO1xuICAgICAgICB9KTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjb25LZXlkb3duID0gKGV2ZW50OiBLZXlib2FyZEV2ZW50KSA9PiB7XG4gICAgaWYgKGV2ZW50LmtleSA9PT0gJ0VzY2FwZScgJiYgdGhpcy5leHBhbmRlZCkge1xuICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgIHRoaXMuY29sbGFwc2UoKTtcbiAgICAgICh0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCd0b2dnbGUnKSBhcyBIVE1MRWxlbWVudCk/LmZvY3VzKCk7XG4gICAgfVxuICB9O1xuXG4gICNvblRvZ2dsZUNsaWNrKCkge1xuICAgIGlmICh0aGlzLmRpc2FibGVkKSByZXR1cm47XG4gICAgdGhpcy50b2dnbGUoKTtcbiAgfVxuXG4gIC8qKiBUb2dnbGUgZXhwYW5kZWQgc3RhdGUgKi9cbiAgdG9nZ2xlKCkge1xuICAgIHRoaXMuZXhwYW5kZWQgPSAhdGhpcy5leHBhbmRlZDtcbiAgfVxuXG4gIC8qKiBFeHBhbmQgdGhlIGRyb3Bkb3duICovXG4gIGV4cGFuZCgpIHtcbiAgICB0aGlzLmV4cGFuZGVkID0gdHJ1ZTtcbiAgfVxuXG4gIC8qKiBDb2xsYXBzZSB0aGUgZHJvcGRvd24gKi9cbiAgY29sbGFwc2UoKSB7XG4gICAgdGhpcy5leHBhbmRlZCA9IGZhbHNlO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LWRyb3Bkb3duJzogUGZWNkRyb3Bkb3duO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxufVxcXFxuXFxcXG5baGlkZGVuXSB7XFxcXG4gIGRpc3BsYXk6IG5vbmUgIWltcG9ydGFudDtcXFxcbn1cXFxcblxcXFxuI3RvZ2dsZSB7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tR2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tZ2FwLS10ZXh0LXRvLWVsZW1lbnQtLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdCbG9ja1N0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tY29udHJvbC0tdmVydGljYWwtLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdJbmxpbmVFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS1ob3Jpem9udGFsLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nQmxvY2tFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tY29udHJvbC0taG9yaXpvbnRhbC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tTWluV2lkdGg6IGNhbGModmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tRm9udFNpemUpICogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tTGluZUhlaWdodCkgKyB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nQmxvY2tTdGFydCkgKyB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nQmxvY2tFbmQpKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Gb250U2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tTGluZUhlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1jb250cm9sLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJSYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXNtYWxsKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1jb250cm9sLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ib3JkZXItLVpJbmRleDogdmFyKC0tcGYtdC0tZ2xvYmFsLS16LWluZGV4LS14cyk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tVHJhbnNpdGlvblRpbWluZ0Z1bmN0aW9uOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1UcmFuc2l0aW9uRHVyYXRpb246IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLVRyYW5zaXRpb25Qcm9wZXJ0eTogY29sb3IsIGJhY2tncm91bmQtY29sb3IsIGJvcmRlci13aWR0aCwgYm9yZGVyLWNvbG9yO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLWhvdmVyLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0taG92ZXItLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tY29udHJvbC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0taG92ZXItLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWNvbnRyb2wtLWhvdmVyKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0taG92ZXIpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLWV4cGFuZGVkLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tZXhwYW5kZWQtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tY29udHJvbC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tZXhwYW5kZWQtLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWNvbnRyb2wtLWNsaWNrZWQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLWV4cGFuZGVkLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1jbGlja2VkKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlX190b2dnbGUtaWNvbi0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuXFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLUNvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Db2xvcik7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tQm9yZGVyUmFkaXVzOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJSYWRpdXMpO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tUGFkZGluZ0Jsb2NrU3RhcnQpO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1QYWRkaW5nQmxvY2tFbmQ6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLVBhZGRpbmdJbmxpbmVTdGFydDogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLUZvbnRTaXplOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Gb250U2l6ZSk7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLUxpbmVIZWlnaHQ6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLUxpbmVIZWlnaHQpO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1NaW5XaWR0aDogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tTWluV2lkdGgpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjo6YWZ0ZXIge1xcXFxuICAgIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgICBpbnNldDogMDtcXFxcbiAgICB6LWluZGV4OiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ib3JkZXItLVpJbmRleCk7XFxcXG4gICAgcG9pbnRlci1ldmVudHM6IG5vbmU7XFxcXG4gICAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcIjtcXFxcbiAgICBib3JkZXI6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJDb2xvcik7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyUmFkaXVzKTtcXFxcbiAgICB0cmFuc2l0aW9uOiBpbmhlcml0O1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyOm5vdChbZGlzYWJsZWRdKSB7XFxcXG4gICAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Db2xvcjogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0taG92ZXItLUNvbG9yKTtcXFxcbiAgICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0taG92ZXItLUJhY2tncm91bmRDb2xvcik7XFxcXG4gICAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0taG92ZXItLUJvcmRlcldpZHRoKTtcXFxcbiAgICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlckNvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQm9yZGVyQ29sb3IpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbjpob3N0KFtleHBhbmRlZF0pICN0b2dnbGUge1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUNvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQ29sb3IpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tZXhwYW5kZWQtLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLWV4cGFuZGVkLS1Cb3JkZXJXaWR0aCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLWV4cGFuZGVkLS1Cb3JkZXJDb2xvcik7XFxcXG59XFxcXG5cXFxcbiN0b2dnbGUgc3ZnW3Nsb3Q9XFxcXFxcXCJpY29uLWVuZFxcXFxcXFwiXSB7XFxcXG4gIHdpZHRoOiAwLjc1ZW07XFxcXG4gIGhlaWdodDogMC43NWVtO1xcXFxuICBmaWxsOiBjdXJyZW50Q29sb3I7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlX190b2dnbGUtaWNvbi0tQ29sb3IpO1xcXFxuICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gMTUwbXMgZWFzZS1pbi1vdXQ7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtleHBhbmRlZF0pICN0b2dnbGUgc3ZnW3Nsb3Q9XFxcXFxcXCJpY29uLWVuZFxcXFxcXFwiXSB7XFxcXG4gIHRyYW5zZm9ybTogcm90YXRlKDE4MGRlZyk7XFxcXG59XFxcXG5cXFxcbiNtZW51LWNvbnRhaW5lciB7XFxcXG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgdG9wOiAxMDAlO1xcXFxuICBsZWZ0OiAwO1xcXFxuICBtYXJnaW4tdG9wOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teHMpO1xcXFxuICB6LWluZGV4OiB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLXNtLCAxMDAwKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2Rpc2FibGVkXSkge1xcXFxuICBwb2ludGVyLWV2ZW50czogbm9uZTtcXFxcbiAgb3BhY2l0eTogMC42O1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLCtuS0FBdW9LLENBQUM7QUFBRSxJQUFPLHlCQUFROzs7QURNOXNLLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQVJQO0FBcUJBLDRCQUFDLGNBQWMsZ0JBQWdCO0FBQ3hCLElBQU0sZ0JBQU4sTUFBTSx1QkFBcUIsaUJBZWhDLGlCQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFHMUMsaUJBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxjQUFDLFNBQVMsSUFyQnNCLElBQVc7QUFBQSxFQUF0QztBQUFBO0FBQUE7QUFnQkwsdUJBQVMsV0FBVyxrQkFBcEIsZ0JBQW9CLFNBQXBCO0FBR0EsdUJBQVMsV0FBVyxrQkFBcEIsaUJBQW9CLFNBQXBCO0FB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
