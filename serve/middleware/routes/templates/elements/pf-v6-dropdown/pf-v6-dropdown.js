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
document.addEventListener("click", (event) => {
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtZHJvcGRvd24vcGYtdjYtZHJvcGRvd24udHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9wZi12Ni1kcm9wZG93bi9wZi12Ni1kcm9wZG93bi5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtZHJvcGRvd24uY3NzJztcblxuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbWVudS9wZi12Ni1tZW51LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbWVudS1pdGVtL3BmLXY2LW1lbnUtaXRlbS5qcyc7XG5cbmltcG9ydCB0eXBlIHsgUGZWNk1lbnUgfSBmcm9tICcuLi9wZi12Ni1tZW51L3BmLXY2LW1lbnUuanMnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgRHJvcGRvd24gY29tcG9uZW50XG4gKlxuICogQGZpcmVzIHtFdmVudH0gZXhwYW5kIC0gRmlyZXMgd2hlbiBkcm9wZG93biBvcGVuc1xuICogQGZpcmVzIHtFdmVudH0gY29sbGFwc2UgLSBGaXJlcyB3aGVuIGRyb3Bkb3duIGNsb3Nlc1xuICpcbiAqIEBzbG90IHRvZ2dsZS10ZXh0IC0gVGV4dCBjb250ZW50IGZvciB0b2dnbGUgYnV0dG9uXG4gKiBAc2xvdCAtIE1lbnUgaXRlbXMgKHBmLXY2LW1lbnUtaXRlbSBlbGVtZW50cylcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ3BmLXY2LWRyb3Bkb3duJylcbmV4cG9ydCBjbGFzcyBQZlY2RHJvcGRvd24gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHJlYWRvbmx5ICNpbnN0YW5jZXMgPSBuZXcgU2V0PFBmVjZEcm9wZG93bj4oKTtcblxuICBzdGF0aWMge1xuICAgIGRvY3VtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKGV2ZW50KSA9PiB7XG4gICAgICBmb3IgKGNvbnN0IGluc3RhbmNlIG9mIFBmVjZEcm9wZG93bi4jaW5zdGFuY2VzKSB7XG4gICAgICAgIGlmIChpbnN0YW5jZS5leHBhbmRlZCAmJiAhZXZlbnQuY29tcG9zZWRQYXRoKCkuaW5jbHVkZXMoaW5zdGFuY2UpKSB7XG4gICAgICAgICAgaW5zdGFuY2UuY29sbGFwc2UoKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0pO1xuICB9XG5cbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGV4cGFuZGVkID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBkaXNhYmxlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSgpXG4gIGFjY2Vzc29yIGxhYmVsID0gJyc7XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICBQZlY2RHJvcGRvd24uI2luc3RhbmNlcy5hZGQodGhpcyk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrZXlkb3duJywgdGhpcy4jb25LZXlkb3duKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgUGZWNkRyb3Bkb3duLiNpbnN0YW5jZXMuZGVsZXRlKHRoaXMpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI29uS2V5ZG93bik7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8cGYtdjYtYnV0dG9uIGlkPVwidG9nZ2xlXCJcbiAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJ0ZXJ0aWFyeVwiXG4gICAgICAgICAgICAgICAgICAgICBhcmlhLWhhc3BvcHVwPVwidHJ1ZVwiXG4gICAgICAgICAgICAgICAgICAgICBhcmlhLWV4cGFuZGVkPVwiJHt0aGlzLmV4cGFuZGVkfVwiXG4gICAgICAgICAgICAgICAgICAgICA/ZGlzYWJsZWQ9JHt0aGlzLmRpc2FibGVkfVxuICAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25Ub2dnbGVDbGlja30+XG4gICAgICAgIDxzbG90IG5hbWU9XCJ0b2dnbGUtdGV4dFwiPlRvZ2dsZTwvc2xvdD5cbiAgICAgICAgPHN2ZyBzbG90PVwiaWNvbi1lbmRcIlxuICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMzIwIDUxMlwiXG4gICAgICAgICAgICAgYXJpYS1oaWRkZW49XCJ0cnVlXCI+XG4gICAgICAgICAgPHBhdGggZD1cIk0zMS4zIDE5MmgyNTcuM2MxNy44IDAgMjYuNyAyMS41IDE0LjEgMzQuMUwxNzQuMSAzNTQuOGMtNy44IDcuOC0yMC41IDcuOC0yOC4zIDBMMTcuMiAyMjYuMUM0LjYgMjEzLjUgMTMuNSAxOTIgMzEuMyAxOTJ6XCIvPlxuICAgICAgICA8L3N2Zz5cbiAgICAgIDwvcGYtdjYtYnV0dG9uPlxuXG4gICAgICA8ZGl2IGlkPVwibWVudS1jb250YWluZXJcIlxuICAgICAgICAgICA/aGlkZGVuPSR7IXRoaXMuZXhwYW5kZWR9PlxuICAgICAgICA8cGYtdjYtbWVudSBpZD1cIm1lbnVcIlxuICAgICAgICAgICAgICAgICAgICAgbGFiZWw9JHt0aGlzLmxhYmVsfT5cbiAgICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICAgIDwvcGYtdjYtbWVudT5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cblxuICB1cGRhdGVkKGNoYW5nZWQ6IE1hcDxzdHJpbmcsIHVua25vd24+KSB7XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdleHBhbmRlZCcpKSB7XG4gICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IEV2ZW50KHRoaXMuZXhwYW5kZWQgPyAnZXhwYW5kJyA6ICdjb2xsYXBzZScsIHsgYnViYmxlczogdHJ1ZSB9KSk7XG4gICAgICBpZiAodGhpcy5leHBhbmRlZCkge1xuICAgICAgICByZXF1ZXN0QW5pbWF0aW9uRnJhbWUoKCkgPT4ge1xuICAgICAgICAgIGNvbnN0IG1lbnUgPSB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdtZW51JykgYXMgUGZWNk1lbnUgfCBudWxsO1xuICAgICAgICAgIG1lbnU/LmZvY3VzRmlyc3RJdGVtKCk7XG4gICAgICAgIH0pO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNvbktleWRvd24gPSAoZXZlbnQ6IEtleWJvYXJkRXZlbnQpID0+IHtcbiAgICBpZiAoZXZlbnQua2V5ID09PSAnRXNjYXBlJyAmJiB0aGlzLmV4cGFuZGVkKSB7XG4gICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgdGhpcy5jb2xsYXBzZSgpO1xuICAgICAgKHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ3RvZ2dsZScpIGFzIEhUTUxFbGVtZW50KT8uZm9jdXMoKTtcbiAgICB9XG4gIH07XG5cbiAgI29uVG9nZ2xlQ2xpY2soKSB7XG4gICAgaWYgKHRoaXMuZGlzYWJsZWQpIHJldHVybjtcbiAgICB0aGlzLnRvZ2dsZSgpO1xuICB9XG5cbiAgLyoqIFRvZ2dsZSBleHBhbmRlZCBzdGF0ZSAqL1xuICB0b2dnbGUoKSB7XG4gICAgdGhpcy5leHBhbmRlZCA9ICF0aGlzLmV4cGFuZGVkO1xuICB9XG5cbiAgLyoqIEV4cGFuZCB0aGUgZHJvcGRvd24gKi9cbiAgZXhwYW5kKCkge1xuICAgIHRoaXMuZXhwYW5kZWQgPSB0cnVlO1xuICB9XG5cbiAgLyoqIENvbGxhcHNlIHRoZSBkcm9wZG93biAqL1xuICBjb2xsYXBzZSgpIHtcbiAgICB0aGlzLmV4cGFuZGVkID0gZmFsc2U7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAncGYtdjYtZHJvcGRvd24nOiBQZlY2RHJvcGRvd247XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG59XFxcXG5cXFxcbltoaWRkZW5dIHtcXFxcbiAgZGlzcGxheTogbm9uZSAhaW1wb3J0YW50O1xcXFxufVxcXFxuXFxcXG4jdG9nZ2xlIHtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1HYXA6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1nYXAtLXRleHQtdG8tZWxlbWVudC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tUGFkZGluZ0Jsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLWhvcml6b250YWwtLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdCbG9ja0VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLXZlcnRpY2FsLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nSW5saW5lU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS1ob3Jpem9udGFsLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1NaW5XaWR0aDogY2FsYyh2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Gb250U2l6ZSkgKiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1MaW5lSGVpZ2h0KSArIHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdCbG9ja1N0YXJ0KSArIHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdCbG9ja0VuZCkpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUZvbnRTaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1MaW5lSGVpZ2h0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLWxpbmUtaGVpZ2h0LS1ib2R5KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWNvbnRyb2wtLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlclJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlckNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWNvbnRyb2wtLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLWJvcmRlci0tWkluZGV4OiB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLXhzKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1UcmFuc2l0aW9uVGltaW5nRnVuY3Rpb246IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLVRyYW5zaXRpb25EdXJhdGlvbjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1zaG9ydCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tVHJhbnNpdGlvblByb3BlcnR5OiBjb2xvciwgYmFja2dyb3VuZC1jb2xvciwgYm9yZGVyLXdpZHRoLCBib3JkZXItY29sb3I7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0taG92ZXItLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1jb250cm9sLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tY29udHJvbC0taG92ZXIpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLWhvdmVyLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1ob3Zlcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tZXhwYW5kZWQtLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1jb250cm9sLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tY29udHJvbC0tY2xpY2tlZCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tZXhwYW5kZWQtLUJvcmRlckNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWNsaWNrZWQpO1xcXFxuICAtLXBmLXY2LWMtbWVudS10b2dnbGVfX3RvZ2dsZS1pY29uLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1pY29uLS1jb2xvci0tcmVndWxhcik7XFxcXG5cXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tQ29sb3I6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLUNvbG9yKTtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1Cb3JkZXJSYWRpdXM6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlclJhZGl1cyk7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLVBhZGRpbmdCbG9ja1N0YXJ0OiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLVBhZGRpbmdCbG9ja0VuZDogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tUGFkZGluZ0Jsb2NrRW5kKTtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nSW5saW5lU3RhcnQpO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1QYWRkaW5nSW5saW5lRW5kOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tRm9udFNpemU6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLUZvbnRTaXplKTtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tTGluZUhlaWdodDogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tTGluZUhlaWdodCk7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLU1pbldpZHRoOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1NaW5XaWR0aCk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OjphZnRlciB7XFxcXG4gICAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICAgIGluc2V0OiAwO1xcXFxuICAgIHotaW5kZXg6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLWJvcmRlci0tWkluZGV4KTtcXFxcbiAgICBwb2ludGVyLWV2ZW50czogbm9uZTtcXFxcbiAgICBjb250ZW50OiBcXFxcXFxcIlxcXFxcXFwiO1xcXFxuICAgIGJvcmRlcjogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyV2lkdGgpIHNvbGlkIHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlckNvbG9yKTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJSYWRpdXMpO1xcXFxuICAgIHRyYW5zaXRpb246IGluaGVyaXQ7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXI6bm90KFtkaXNhYmxlZF0pIHtcXFxcbiAgICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUNvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQ29sb3IpO1xcXFxuICAgIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgICAtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlcldpZHRoOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQm9yZGVyV2lkdGgpO1xcXFxuICAgIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLWhvdmVyLS1Cb3JkZXJDb2xvcik7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuOmhvc3QoW2V4cGFuZGVkXSkgI3RvZ2dsZSB7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQ29sb3I6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGUtLWV4cGFuZGVkLS1Db2xvcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tZXhwYW5kZWQtLUJvcmRlcldpZHRoKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tZXhwYW5kZWQtLUJvcmRlckNvbG9yKTtcXFxcbn1cXFxcblxcXFxuI3RvZ2dsZSBzdmdbc2xvdD1cXFxcXFxcImljb24tZW5kXFxcXFxcXCJdIHtcXFxcbiAgd2lkdGg6IDAuNzVlbTtcXFxcbiAgaGVpZ2h0OiAwLjc1ZW07XFxcXG4gIGZpbGw6IGN1cnJlbnRDb2xvcjtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXY2LWMtbWVudS10b2dnbGVfX3RvZ2dsZS1pY29uLS1Db2xvcik7XFxcXG4gIHRyYW5zaXRpb246IHRyYW5zZm9ybSAxNTBtcyBlYXNlLWluLW91dDtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2V4cGFuZGVkXSkgI3RvZ2dsZSBzdmdbc2xvdD1cXFxcXFxcImljb24tZW5kXFxcXFxcXCJdIHtcXFxcbiAgdHJhbnNmb3JtOiByb3RhdGUoMTgwZGVnKTtcXFxcbn1cXFxcblxcXFxuI21lbnUtY29udGFpbmVyIHtcXFxcbiAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICB0b3A6IDEwMCU7XFxcXG4gIGxlZnQ6IDA7XFxcXG4gIG1hcmdpbi10b3A6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS14cyk7XFxcXG4gIHotaW5kZXg6IHZhcigtLXBmLXQtLWdsb2JhbC0tei1pbmRleC0tc20sIDEwMDApO1xcXFxufVxcXFxuXFxcXG46aG9zdChbZGlzYWJsZWRdKSB7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxuICBvcGFjaXR5OiAwLjY7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sK25LQUF1b0ssQ0FBQztBQUFFLElBQU8seUJBQVE7OztBRE05c0ssT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBUlA7QUFxQkEsNEJBQUMsY0FBYyxnQkFBZ0I7QUFDeEIsSUFBTSxnQkFBTixNQUFNLHVCQUFxQixpQkFlaEMsaUJBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxpQkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBRzFDLGNBQUMsU0FBUyxJQXJCc0IsSUFBVztBQUFBLEVBQXRDO0FBQUE7QUFBQTtBQWdCTCx1QkFBUyxXQUFXLGtCQUFwQixnQkFBb0IsU0FBcEI7QUFHQSx1QkFBUyxXQUFXLGtCQUFwQixpQkFBb0IsU0FBcEI7QUFHQSx1QkFBUyxR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
