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

// elements/cem-pf-v6-dropdown/cem-pf-v6-dropdown.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-dropdown/cem-pf-v6-dropdown.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: inline-block;\\n  position: relative;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n#toggle {\\n  --cem-pf-v6-c-menu-toggle--Gap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-menu-toggle--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-menu-toggle--PaddingInlineEnd: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-menu-toggle--PaddingBlockEnd: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-menu-toggle--PaddingInlineStart: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-menu-toggle--MinWidth: calc(var(--cem-pf-v6-c-menu-toggle--FontSize) * var(--cem-pf-v6-c-menu-toggle--LineHeight) + var(--cem-pf-v6-c-menu-toggle--PaddingBlockStart) + var(--cem-pf-v6-c-menu-toggle--PaddingBlockEnd));\\n  --cem-pf-v6-c-menu-toggle--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-menu-toggle--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu-toggle--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-menu-toggle--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-menu-toggle--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-menu-toggle--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-menu-toggle--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-menu-toggle--border--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-menu-toggle--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-menu-toggle--TransitionDuration: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-menu-toggle--TransitionProperty: color, background-color, border-width, border-color;\\n  --cem-pf-v6-c-menu-toggle--hover--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu-toggle--hover--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-menu-toggle--hover--BorderWidth: var(--pf-t--global--border--width--control--hover);\\n  --cem-pf-v6-c-menu-toggle--hover--BorderColor: var(--pf-t--global--border--color--hover);\\n  --cem-pf-v6-c-menu-toggle--expanded--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu-toggle--expanded--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-menu-toggle--expanded--BorderWidth: var(--pf-t--global--border--width--control--clicked);\\n  --cem-pf-v6-c-menu-toggle--expanded--BorderColor: var(--pf-t--global--border--color--clicked);\\n  --cem-pf-v6-c-menu-toggle__toggle-icon--Color: var(--pf-t--global--icon--color--regular);\\n\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-menu-toggle--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-menu-toggle--BackgroundColor);\\n  --cem-pf-v6-c-button--BorderRadius: var(--cem-pf-v6-c-menu-toggle--BorderRadius);\\n  --cem-pf-v6-c-button--PaddingBlockStart: var(--cem-pf-v6-c-menu-toggle--PaddingBlockStart);\\n  --cem-pf-v6-c-button--PaddingBlockEnd: var(--cem-pf-v6-c-menu-toggle--PaddingBlockEnd);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--cem-pf-v6-c-menu-toggle--PaddingInlineStart);\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--cem-pf-v6-c-menu-toggle--PaddingInlineEnd);\\n  --cem-pf-v6-c-button--FontSize: var(--cem-pf-v6-c-menu-toggle--FontSize);\\n  --cem-pf-v6-c-button--LineHeight: var(--cem-pf-v6-c-menu-toggle--LineHeight);\\n  --cem-pf-v6-c-button--MinWidth: var(--cem-pf-v6-c-menu-toggle--MinWidth);\\n\\n  \\u0026::after {\\n    position: absolute;\\n    inset: 0;\\n    z-index: var(--cem-pf-v6-c-menu-toggle--border--ZIndex);\\n    pointer-events: none;\\n    content: \\"\\";\\n    border: var(--cem-pf-v6-c-menu-toggle--BorderWidth) solid var(--cem-pf-v6-c-menu-toggle--BorderColor);\\n    border-radius: var(--cem-pf-v6-c-menu-toggle--BorderRadius);\\n    transition: inherit;\\n  }\\n\\n  \\u0026:hover:not([disabled]) {\\n    --cem-pf-v6-c-menu-toggle--Color: var(--cem-pf-v6-c-menu-toggle--hover--Color);\\n    --cem-pf-v6-c-menu-toggle--BackgroundColor: var(--cem-pf-v6-c-menu-toggle--hover--BackgroundColor);\\n    --cem-pf-v6-c-menu-toggle--BorderWidth: var(--cem-pf-v6-c-menu-toggle--hover--BorderWidth);\\n    --cem-pf-v6-c-menu-toggle--BorderColor: var(--cem-pf-v6-c-menu-toggle--hover--BorderColor);\\n  }\\n}\\n\\n:host([expanded]) #toggle {\\n  --cem-pf-v6-c-menu-toggle--Color: var(--cem-pf-v6-c-menu-toggle--expanded--Color);\\n  --cem-pf-v6-c-menu-toggle--BackgroundColor: var(--cem-pf-v6-c-menu-toggle--expanded--BackgroundColor);\\n  --cem-pf-v6-c-menu-toggle--BorderWidth: var(--cem-pf-v6-c-menu-toggle--expanded--BorderWidth);\\n  --cem-pf-v6-c-menu-toggle--BorderColor: var(--cem-pf-v6-c-menu-toggle--expanded--BorderColor);\\n}\\n\\n#toggle svg[slot=\\"icon-end\\"] {\\n  width: 0.75em;\\n  height: 0.75em;\\n  fill: currentColor;\\n  color: var(--cem-pf-v6-c-menu-toggle__toggle-icon--Color);\\n  transition: transform 150ms ease-in-out;\\n}\\n\\n:host([expanded]) #toggle svg[slot=\\"icon-end\\"] {\\n  transform: rotate(180deg);\\n}\\n\\n#menu-container {\\n  position: absolute;\\n  top: 100%;\\n  left: 0;\\n  margin-top: var(--pf-t--global--spacer--xs);\\n  z-index: var(--pf-t--global--z-index--sm, 1000);\\n}\\n\\n:host([disabled]) {\\n  pointer-events: none;\\n  opacity: 0.6;\\n}\\n"'));
var cem_pf_v6_dropdown_default = s;

// elements/cem-pf-v6-dropdown/cem-pf-v6-dropdown.ts
import "../cem-pf-v6-button/cem-pf-v6-button.js";
import "../cem-pf-v6-menu/cem-pf-v6-menu.js";
import "../cem-pf-v6-menu-item/cem-pf-v6-menu-item.js";
var _label_dec, _disabled_dec, _expanded_dec, _a, _PfV6Dropdown_decorators, _instances, _init, _expanded, _disabled, _label, _onKeydown, _PfV6Dropdown_instances, onToggleClick_fn;
_PfV6Dropdown_decorators = [customElement("cem-pf-v6-dropdown")];
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
      <cem-pf-v6-button id="toggle"
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
      </cem-pf-v6-button>

      <div id="menu-container"
           ?hidden=${!this.expanded}>
        <cem-pf-v6-menu id="menu"
                     label=${this.label}>
          <slot></slot>
        </cem-pf-v6-menu>
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
__publicField(_PfV6Dropdown, "styles", cem_pf_v6_dropdown_default);
__runInitializers(_init, 1, _PfV6Dropdown);
var PfV6Dropdown = _PfV6Dropdown;
export {
  PfV6Dropdown
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LWRyb3Bkb3duL2NlbS1wZi12Ni1kcm9wZG93bi50cyIsICJsaXQtY3NzOmVsZW1lbnRzL2NlbS1wZi12Ni1kcm9wZG93bi9jZW0tcGYtdjYtZHJvcGRvd24uY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1wZi12Ni1kcm9wZG93bi5jc3MnO1xuXG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1idXR0b24vY2VtLXBmLXY2LWJ1dHRvbi5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1tZW51L2NlbS1wZi12Ni1tZW51LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LW1lbnUtaXRlbS9jZW0tcGYtdjYtbWVudS1pdGVtLmpzJztcblxuaW1wb3J0IHR5cGUgeyBQZlY2TWVudSB9IGZyb20gJy4uL2NlbS1wZi12Ni1tZW51L2NlbS1wZi12Ni1tZW51LmpzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IERyb3Bkb3duIGNvbXBvbmVudFxuICpcbiAqIEBmaXJlcyB7RXZlbnR9IGV4cGFuZCAtIEZpcmVzIHdoZW4gZHJvcGRvd24gb3BlbnNcbiAqIEBmaXJlcyB7RXZlbnR9IGNvbGxhcHNlIC0gRmlyZXMgd2hlbiBkcm9wZG93biBjbG9zZXNcbiAqXG4gKiBAc2xvdCB0b2dnbGUtdGV4dCAtIFRleHQgY29udGVudCBmb3IgdG9nZ2xlIGJ1dHRvblxuICogQHNsb3QgLSBNZW51IGl0ZW1zIChjZW0tcGYtdjYtbWVudS1pdGVtIGVsZW1lbnRzKVxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXBmLXY2LWRyb3Bkb3duJylcbmV4cG9ydCBjbGFzcyBQZlY2RHJvcGRvd24gZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHJlYWRvbmx5ICNpbnN0YW5jZXMgPSBuZXcgU2V0PFBmVjZEcm9wZG93bj4oKTtcblxuICBzdGF0aWMge1xuICAgIGRvY3VtZW50Py5hZGRFdmVudExpc3RlbmVyPy4oJ2NsaWNrJywgKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgICAgZm9yIChjb25zdCBpbnN0YW5jZSBvZiBQZlY2RHJvcGRvd24uI2luc3RhbmNlcykge1xuICAgICAgICBpZiAoaW5zdGFuY2UuZXhwYW5kZWQgJiYgIWV2ZW50LmNvbXBvc2VkUGF0aCgpLmluY2x1ZGVzKGluc3RhbmNlKSkge1xuICAgICAgICAgIGluc3RhbmNlLmNvbGxhcHNlKCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBleHBhbmRlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZGlzYWJsZWQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBsYWJlbCA9ICcnO1xuXG4gIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmNvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgUGZWNkRyb3Bkb3duLiNpbnN0YW5jZXMuYWRkKHRoaXMpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI29uS2V5ZG93bik7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIFBmVjZEcm9wZG93bi4jaW5zdGFuY2VzLmRlbGV0ZSh0aGlzKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNvbktleWRvd24pO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGNlbS1wZi12Ni1idXR0b24gaWQ9XCJ0b2dnbGVcIlxuICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInRlcnRpYXJ5XCJcbiAgICAgICAgICAgICAgICAgICAgIGFyaWEtaGFzcG9wdXA9XCJ0cnVlXCJcbiAgICAgICAgICAgICAgICAgICAgIGFyaWEtZXhwYW5kZWQ9XCIke3RoaXMuZXhwYW5kZWR9XCJcbiAgICAgICAgICAgICAgICAgICAgID9kaXNhYmxlZD0ke3RoaXMuZGlzYWJsZWR9XG4gICAgICAgICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLiNvblRvZ2dsZUNsaWNrfT5cbiAgICAgICAgPHNsb3QgbmFtZT1cInRvZ2dsZS10ZXh0XCI+VG9nZ2xlPC9zbG90PlxuICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uLWVuZFwiXG4gICAgICAgICAgICAgdmlld0JveD1cIjAgMCAzMjAgNTEyXCJcbiAgICAgICAgICAgICBhcmlhLWhpZGRlbj1cInRydWVcIj5cbiAgICAgICAgICA8cGF0aCBkPVwiTTMxLjMgMTkyaDI1Ny4zYzE3LjggMCAyNi43IDIxLjUgMTQuMSAzNC4xTDE3NC4xIDM1NC44Yy03LjggNy44LTIwLjUgNy44LTI4LjMgMEwxNy4yIDIyNi4xQzQuNiAyMTMuNSAxMy41IDE5MiAzMS4zIDE5MnpcIi8+XG4gICAgICAgIDwvc3ZnPlxuICAgICAgPC9jZW0tcGYtdjYtYnV0dG9uPlxuXG4gICAgICA8ZGl2IGlkPVwibWVudS1jb250YWluZXJcIlxuICAgICAgICAgICA/aGlkZGVuPSR7IXRoaXMuZXhwYW5kZWR9PlxuICAgICAgICA8Y2VtLXBmLXY2LW1lbnUgaWQ9XCJtZW51XCJcbiAgICAgICAgICAgICAgICAgICAgIGxhYmVsPSR7dGhpcy5sYWJlbH0+XG4gICAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgICA8L2NlbS1wZi12Ni1tZW51PlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gIHVwZGF0ZWQoY2hhbmdlZDogTWFwPHN0cmluZywgdW5rbm93bj4pIHtcbiAgICBpZiAoY2hhbmdlZC5oYXMoJ2V4cGFuZGVkJykpIHtcbiAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgRXZlbnQodGhpcy5leHBhbmRlZCA/ICdleHBhbmQnIDogJ2NvbGxhcHNlJywgeyBidWJibGVzOiB0cnVlIH0pKTtcbiAgICAgIGlmICh0aGlzLmV4cGFuZGVkKSB7XG4gICAgICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICAgICAgY29uc3QgbWVudSA9IHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ21lbnUnKSBhcyBQZlY2TWVudSB8IG51bGw7XG4gICAgICAgICAgbWVudT8uZm9jdXNGaXJzdEl0ZW0oKTtcbiAgICAgICAgfSk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI29uS2V5ZG93biA9IChldmVudDogS2V5Ym9hcmRFdmVudCkgPT4ge1xuICAgIGlmIChldmVudC5rZXkgPT09ICdFc2NhcGUnICYmIHRoaXMuZXhwYW5kZWQpIHtcbiAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICB0aGlzLmNvbGxhcHNlKCk7XG4gICAgICAodGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgndG9nZ2xlJykgYXMgSFRNTEVsZW1lbnQpPy5mb2N1cygpO1xuICAgIH1cbiAgfTtcblxuICAjb25Ub2dnbGVDbGljaygpIHtcbiAgICBpZiAodGhpcy5kaXNhYmxlZCkgcmV0dXJuO1xuICAgIHRoaXMudG9nZ2xlKCk7XG4gIH1cblxuICAvKiogVG9nZ2xlIGV4cGFuZGVkIHN0YXRlICovXG4gIHRvZ2dsZSgpIHtcbiAgICB0aGlzLmV4cGFuZGVkID0gIXRoaXMuZXhwYW5kZWQ7XG4gIH1cblxuICAvKiogRXhwYW5kIHRoZSBkcm9wZG93biAqL1xuICBleHBhbmQoKSB7XG4gICAgdGhpcy5leHBhbmRlZCA9IHRydWU7XG4gIH1cblxuICAvKiogQ29sbGFwc2UgdGhlIGRyb3Bkb3duICovXG4gIGNvbGxhcHNlKCkge1xuICAgIHRoaXMuZXhwYW5kZWQgPSBmYWxzZTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcGYtdjYtZHJvcGRvd24nOiBQZlY2RHJvcGRvd247XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG59XFxcXG5cXFxcbltoaWRkZW5dIHtcXFxcbiAgZGlzcGxheTogbm9uZSAhaW1wb3J0YW50O1xcXFxufVxcXFxuXFxcXG4jdG9nZ2xlIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tR2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tZ2FwLS10ZXh0LXRvLWVsZW1lbnQtLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLXZlcnRpY2FsLS1kZWZhdWx0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLWhvcml6b250YWwtLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nQmxvY2tFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdJbmxpbmVTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLWhvcml6b250YWwtLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1NaW5XaWR0aDogY2FsYyh2YXIoLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tRm9udFNpemUpICogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLUxpbmVIZWlnaHQpICsgdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdCbG9ja1N0YXJ0KSArIHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nQmxvY2tFbmQpKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tRm9udFNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tTGluZUhlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tY29udHJvbC0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlclJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tY29udHJvbC0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLWJvcmRlci0tWkluZGV4OiB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLXhzKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tVHJhbnNpdGlvblRpbWluZ0Z1bmN0aW9uOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tVHJhbnNpdGlvbkR1cmF0aW9uOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLXNob3J0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tVHJhbnNpdGlvblByb3BlcnR5OiBjb2xvciwgYmFja2dyb3VuZC1jb2xvciwgYm9yZGVyLXdpZHRoLCBib3JkZXItY29sb3I7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLWhvdmVyLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLWhvdmVyLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWNvbnRyb2wtLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tY29udHJvbC0taG92ZXIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0taG92ZXIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1jb250cm9sLS1kZWZhdWx0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tZXhwYW5kZWQtLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWNvbnRyb2wtLWNsaWNrZWQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tY2xpY2tlZCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGVfX3RvZ2dsZS1pY29uLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1pY29uLS1jb2xvci0tcmVndWxhcik7XFxcXG5cXFxcbiAgLS1jZW0tcGYtdjYtYy1idXR0b24tLUNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQ29sb3IpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWJ1dHRvbi0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1idXR0b24tLUJvcmRlclJhZGl1czogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlclJhZGl1cyk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdCbG9ja1N0YXJ0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1idXR0b24tLVBhZGRpbmdCbG9ja0VuZDogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1QYWRkaW5nSW5saW5lU3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1QYWRkaW5nSW5saW5lU3RhcnQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWJ1dHRvbi0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLVBhZGRpbmdJbmxpbmVFbmQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWJ1dHRvbi0tRm9udFNpemU6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Gb250U2l6ZSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1MaW5lSGVpZ2h0OiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tTGluZUhlaWdodCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1NaW5XaWR0aDogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLU1pbldpZHRoKTtcXFxcblxcXFxuICBcXFxcdTAwMjY6OmFmdGVyIHtcXFxcbiAgICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gICAgaW5zZXQ6IDA7XFxcXG4gICAgei1pbmRleDogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLWJvcmRlci0tWkluZGV4KTtcXFxcbiAgICBwb2ludGVyLWV2ZW50czogbm9uZTtcXFxcbiAgICBjb250ZW50OiBcXFxcXFxcIlxcXFxcXFwiO1xcXFxuICAgIGJvcmRlcjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyQ29sb3IpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJSYWRpdXMpO1xcXFxuICAgIHRyYW5zaXRpb246IGluaGVyaXQ7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXI6bm90KFtkaXNhYmxlZF0pIHtcXFxcbiAgICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Db2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLWhvdmVyLS1Db2xvcik7XFxcXG4gICAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0taG92ZXItLUJhY2tncm91bmRDb2xvcik7XFxcXG4gICAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyV2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1ob3Zlci0tQm9yZGVyV2lkdGgpO1xcXFxuICAgIC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLUJvcmRlckNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0taG92ZXItLUJvcmRlckNvbG9yKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG46aG9zdChbZXhwYW5kZWRdKSAjdG9nZ2xlIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQ29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQ29sb3IpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZS0tQm9yZGVyV2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1leHBhbmRlZC0tQm9yZGVyV2lkdGgpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtdG9nZ2xlLS1Cb3JkZXJDb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS10b2dnbGUtLWV4cGFuZGVkLS1Cb3JkZXJDb2xvcik7XFxcXG59XFxcXG5cXFxcbiN0b2dnbGUgc3ZnW3Nsb3Q9XFxcXFxcXCJpY29uLWVuZFxcXFxcXFwiXSB7XFxcXG4gIHdpZHRoOiAwLjc1ZW07XFxcXG4gIGhlaWdodDogMC43NWVtO1xcXFxuICBmaWxsOiBjdXJyZW50Q29sb3I7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LXRvZ2dsZV9fdG9nZ2xlLWljb24tLUNvbG9yKTtcXFxcbiAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIDE1MG1zIGVhc2UtaW4tb3V0O1xcXFxufVxcXFxuXFxcXG46aG9zdChbZXhwYW5kZWRdKSAjdG9nZ2xlIHN2Z1tzbG90PVxcXFxcXFwiaWNvbi1lbmRcXFxcXFxcIl0ge1xcXFxuICB0cmFuc2Zvcm06IHJvdGF0ZSgxODBkZWcpO1xcXFxufVxcXFxuXFxcXG4jbWVudS1jb250YWluZXIge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIHRvcDogMTAwJTtcXFxcbiAgbGVmdDogMDtcXFxcbiAgbWFyZ2luLXRvcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXhzKTtcXFxcbiAgei1pbmRleDogdmFyKC0tcGYtdC0tZ2xvYmFsLS16LWluZGV4LS1zbSwgMTAwMCk7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtkaXNhYmxlZF0pIHtcXFxcbiAgcG9pbnRlci1ldmVudHM6IG5vbmU7XFxcXG4gIG9wYWNpdHk6IDAuNjtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSwyNUtBQW02SyxD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
