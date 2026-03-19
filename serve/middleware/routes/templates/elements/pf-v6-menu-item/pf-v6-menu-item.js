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

// elements/pf-v6-menu-item/pf-v6-menu-item.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-menu-item/pf-v6-menu-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --pf-v6-c-menu__list-item--Color: var(--pf-t--global--text--color--regular);\\n  --pf-v6-c-menu__list-item--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --pf-v6-c-menu__list-item--BorderWidth: var(--pf-t--global--border--width--action--plain--default);\\n  --pf-v6-c-menu__list-item--hover--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --pf-v6-c-menu__list-item--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --pf-v6-c-menu__list-item--TransitionDuration: var(--pf-t--global--motion--duration--fade--short);\\n  --pf-v6-c-menu__list-item--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --pf-v6-c-menu__list-item--TransitionProperty: background-color;\\n  --pf-v6-c-menu__list-item--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --pf-v6-c-menu__item--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-menu__item--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-menu__item--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-menu__item--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-menu__item--FontSize: var(--pf-t--global--font--size--body--default);\\n  --pf-v6-c-menu__item--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --pf-v6-c-menu__item--FontWeight: var(--pf-t--global--font--weight--body--default);\\n  --pf-v6-c-menu__item--Color: var(--pf-t--global--text--color--regular);\\n  --pf-v6-c-menu__item--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --pf-v6-c-menu__item--m-disabled--Color: var(--pf-t--global--text--color--disabled);\\n  --pf-v6-c-menu__item-main--ColumnGap: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-menu__item-description--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --pf-v6-c-menu__item-description--Color: var(--pf-t--global--text--color--subtle);\\n  --pf-v6-c-menu--OutlineOffset: calc(var(--pf-t--global--border--width--control--default) * -3);\\n\\n  position: relative;\\n  display: flex;\\n  align-items: baseline;\\n  min-width: 0;\\n  padding-block-start: var(--pf-v6-c-menu__item--PaddingBlockStart);\\n  padding-block-end: var(--pf-v6-c-menu__item--PaddingBlockEnd);\\n  padding-inline-start: var(--pf-v6-c-menu__item--PaddingInlineStart);\\n  padding-inline-end: var(--pf-v6-c-menu__item--PaddingInlineEnd);\\n  font-size: var(--pf-v6-c-menu__item--FontSize);\\n  font-weight: var(--pf-v6-c-menu__item--FontWeight);\\n  line-height: var(--pf-v6-c-menu__item--LineHeight);\\n  color: var(--pf-v6-c-menu__item--Color);\\n  text-align: start;\\n  text-decoration-line: none;\\n  background-color: transparent;\\n  border: 0;\\n  outline-offset: var(--pf-v6-c-menu--OutlineOffset);\\n  cursor: pointer;\\n  user-select: none;\\n  transition-timing-function: var(--pf-v6-c-menu__list-item--TransitionTimingFunction);\\n  transition-duration: var(--pf-v6-c-menu__list-item--TransitionDuration);\\n  transition-property: var(--pf-v6-c-menu__list-item--TransitionProperty);\\n  gap: var(--pf-v6-c-menu__item-main--ColumnGap);\\n}\\n\\n:host::before {\\n  position: absolute;\\n  inset: 0;\\n  content: \\"\\";\\n  background-color: var(--pf-v6-c-menu__list-item--BackgroundColor);\\n  border-block-start: var(--pf-v6-c-menu__list-item--BorderWidth) solid var(--pf-v6-c-menu__list-item--BorderColor);\\n  border-block-end: var(--pf-v6-c-menu__list-item--BorderWidth) solid var(--pf-v6-c-menu__list-item--BorderColor);\\n  transition: inherit;\\n}\\n\\n:host(:hover:not([disabled])),\\n:host(:focus-visible) {\\n  --pf-v6-c-menu__list-item--BackgroundColor: var(--pf-v6-c-menu__list-item--hover--BackgroundColor);\\n  --pf-v6-c-menu__list-item--BorderWidth: var(--pf-v6-c-menu__list-item--hover--BorderWidth);\\n}\\n\\n:host(:focus-visible) {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: -2px;\\n}\\n\\n:host([disabled]) {\\n  --pf-v6-c-menu__item--Color: var(--pf-v6-c-menu__item--m-disabled--Color);\\n  --pf-v6-c-menu__list-item--BackgroundColor: transparent;\\n  --pf-v6-c-menu__list-item--hover--BackgroundColor: transparent;\\n  cursor: not-allowed;\\n  pointer-events: none;\\n}\\n\\n#check,\\n#text,\\n#description {\\n  position: relative;\\n  z-index: 1;\\n}\\n\\n#check {\\n  display: flex;\\n  align-items: center;\\n  flex-shrink: 0;\\n\\n  \\u0026 .pf-v6-c-check {\\n    --pf-v6-c-check__input--TranslateY: none;\\n  }\\n}\\n\\n#text {\\n  overflow: hidden;\\n  text-overflow: ellipsis;\\n  white-space: nowrap;\\n  flex-grow: 1;\\n}\\n\\n#description {\\n  display: block;\\n  font-size: var(--pf-v6-c-menu__item-description--FontSize);\\n  color: var(--pf-v6-c-menu__item-description--Color);\\n  word-break: break-word;\\n  margin-block-start: 0.25rem;\\n}\\n\\nslot {\\n  display: inline;\\n}\\n"'));
var pf_v6_menu_item_default = s;

// elements/pf-v6-menu-item/pf-v6-menu-item.ts
var PfMenuItemSelectEvent = class extends Event {
  value;
  checked;
  constructor(value, checked) {
    super("select", { bubbles: true });
    this.value = value;
    this.checked = checked;
  }
};
var _description_dec, _value_dec, _variant_dec, _checked_dec, _disabled_dec, _a, _PfV6MenuItem_decorators, _internals, _init, _disabled, _checked, _variant, _value, _description, _handleClick, _handleKeydown;
_PfV6MenuItem_decorators = [customElement("pf-v6-menu-item")];
var PfV6MenuItem = class extends (_a = LitElement, _disabled_dec = [property({ type: Boolean, reflect: true })], _checked_dec = [property({ type: Boolean, reflect: true })], _variant_dec = [property({ reflect: true })], _value_dec = [property({ reflect: true })], _description_dec = [property({ reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _disabled, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _checked, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateAdd(this, _variant, __runInitializers(_init, 16, this, "default")), __runInitializers(_init, 19, this);
    __privateAdd(this, _value, __runInitializers(_init, 20, this, "")), __runInitializers(_init, 23, this);
    __privateAdd(this, _description, __runInitializers(_init, 24, this)), __runInitializers(_init, 27, this);
    __privateAdd(this, _handleClick, (event) => {
      if (this.disabled) {
        event.preventDefault();
        event.stopImmediatePropagation();
        return;
      }
      if (this.variant === "checkbox") {
        this.checked = !this.checked;
      }
      this.dispatchEvent(new PfMenuItemSelectEvent(this.value, this.checked));
    });
    __privateAdd(this, _handleKeydown, (event) => {
      if (this.disabled) return;
      if (event.key === " " || event.key === "Enter") {
        event.preventDefault();
        __privateGet(this, _handleClick).call(this, event);
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("click", __privateGet(this, _handleClick));
    this.addEventListener("keydown", __privateGet(this, _handleKeydown));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("click", __privateGet(this, _handleClick));
    this.removeEventListener("keydown", __privateGet(this, _handleKeydown));
  }
  updated(changed) {
    if (changed.has("variant")) {
      __privateGet(this, _internals).role = this.variant === "checkbox" ? "menuitemcheckbox" : "menuitem";
    }
    if (changed.has("checked") || changed.has("variant")) {
      __privateGet(this, _internals).ariaChecked = this.variant === "checkbox" ? String(this.checked) : null;
    }
    if (changed.has("disabled")) {
      __privateGet(this, _internals).ariaDisabled = this.disabled ? "true" : null;
      this.setAttribute("tabindex", "-1");
    }
  }
  render() {
    return html`
      ${this.variant === "checkbox" ? html`
        <span id="check"
              class="pf-v6-c-check pf-m-standalone">
          <input id="input"
                 class="pf-v6-c-check__input"
                 role="presentation"
                 tabindex="-1"
                 type="checkbox"
                 .checked=${this.checked}
                 ?disabled=${this.disabled}>
        </span>
      ` : nothing}
      <span id="text"><slot></slot></span>
      ${this.description ? html`
        <span id="description">${this.description}</span>
      ` : nothing}
    `;
  }
};
_init = __decoratorStart(_a);
_internals = new WeakMap();
_disabled = new WeakMap();
_checked = new WeakMap();
_variant = new WeakMap();
_value = new WeakMap();
_description = new WeakMap();
_handleClick = new WeakMap();
_handleKeydown = new WeakMap();
__decorateElement(_init, 4, "disabled", _disabled_dec, PfV6MenuItem, _disabled);
__decorateElement(_init, 4, "checked", _checked_dec, PfV6MenuItem, _checked);
__decorateElement(_init, 4, "variant", _variant_dec, PfV6MenuItem, _variant);
__decorateElement(_init, 4, "value", _value_dec, PfV6MenuItem, _value);
__decorateElement(_init, 4, "description", _description_dec, PfV6MenuItem, _description);
PfV6MenuItem = __decorateElement(_init, 0, "PfV6MenuItem", _PfV6MenuItem_decorators, PfV6MenuItem);
__publicField(PfV6MenuItem, "styles", pf_v6_menu_item_default);
__runInitializers(_init, 1, PfV6MenuItem);
export {
  PfMenuItemSelectEvent,
  PfV6MenuItem
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtbWVudS1pdGVtL3BmLXY2LW1lbnUtaXRlbS50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL3BmLXY2LW1lbnUtaXRlbS9wZi12Ni1tZW51LWl0ZW0uY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sLCBub3RoaW5nIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL3BmLXY2LW1lbnUtaXRlbS5jc3MnO1xuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIG1lbnUgaXRlbSBpcyBzZWxlY3RlZC9jaGVja2VkXG4gKi9cbmV4cG9ydCBjbGFzcyBQZk1lbnVJdGVtU2VsZWN0RXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIHZhbHVlOiBzdHJpbmc7XG4gIGNoZWNrZWQ6IGJvb2xlYW47XG4gIGNvbnN0cnVjdG9yKHZhbHVlOiBzdHJpbmcsIGNoZWNrZWQ6IGJvb2xlYW4pIHtcbiAgICBzdXBlcignc2VsZWN0JywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMudmFsdWUgPSB2YWx1ZTtcbiAgICB0aGlzLmNoZWNrZWQgPSBjaGVja2VkO1xuICB9XG59XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBNZW51IEl0ZW1cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgaXRlbSB0ZXh0XG4gKlxuICogQGZpcmVzIHNlbGVjdCAtIEZpcmVzIHdoZW4gaXRlbSBpcyBzZWxlY3RlZC9jaGVja2VkXG4gKi9cbkBjdXN0b21FbGVtZW50KCdwZi12Ni1tZW51LWl0ZW0nKVxuZXhwb3J0IGNsYXNzIFBmVjZNZW51SXRlbSBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gICNpbnRlcm5hbHMgPSB0aGlzLmF0dGFjaEludGVybmFscygpO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZGlzYWJsZWQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGNoZWNrZWQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIHZhcmlhbnQ6ICdkZWZhdWx0JyB8ICdjaGVja2JveCcgPSAnZGVmYXVsdCc7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciB2YWx1ZSA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZGVzY3JpcHRpb24/OiBzdHJpbmc7XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgdGhpcy4jaGFuZGxlQ2xpY2spO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI2hhbmRsZUtleWRvd24pO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgdGhpcy4jaGFuZGxlQ2xpY2spO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI2hhbmRsZUtleWRvd24pO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygndmFyaWFudCcpKSB7XG4gICAgICB0aGlzLiNpbnRlcm5hbHMucm9sZSA9IHRoaXMudmFyaWFudCA9PT0gJ2NoZWNrYm94JyA/ICdtZW51aXRlbWNoZWNrYm94JyA6ICdtZW51aXRlbSc7XG4gICAgfVxuICAgIGlmIChjaGFuZ2VkLmhhcygnY2hlY2tlZCcpIHx8IGNoYW5nZWQuaGFzKCd2YXJpYW50JykpIHtcbiAgICAgIHRoaXMuI2ludGVybmFscy5hcmlhQ2hlY2tlZCA9IHRoaXMudmFyaWFudCA9PT0gJ2NoZWNrYm94JyA/IFN0cmluZyh0aGlzLmNoZWNrZWQpIDogbnVsbDtcbiAgICB9XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdkaXNhYmxlZCcpKSB7XG4gICAgICB0aGlzLiNpbnRlcm5hbHMuYXJpYURpc2FibGVkID0gdGhpcy5kaXNhYmxlZCA/ICd0cnVlJyA6IG51bGw7XG4gICAgICB0aGlzLnNldEF0dHJpYnV0ZSgndGFiaW5kZXgnLCAnLTEnKTtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICAke3RoaXMudmFyaWFudCA9PT0gJ2NoZWNrYm94JyA/IGh0bWxgXG4gICAgICAgIDxzcGFuIGlkPVwiY2hlY2tcIlxuICAgICAgICAgICAgICBjbGFzcz1cInBmLXY2LWMtY2hlY2sgcGYtbS1zdGFuZGFsb25lXCI+XG4gICAgICAgICAgPGlucHV0IGlkPVwiaW5wdXRcIlxuICAgICAgICAgICAgICAgICBjbGFzcz1cInBmLXY2LWMtY2hlY2tfX2lucHV0XCJcbiAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiXG4gICAgICAgICAgICAgICAgIHRhYmluZGV4PVwiLTFcIlxuICAgICAgICAgICAgICAgICB0eXBlPVwiY2hlY2tib3hcIlxuICAgICAgICAgICAgICAgICAuY2hlY2tlZD0ke3RoaXMuY2hlY2tlZH1cbiAgICAgICAgICAgICAgICAgP2Rpc2FibGVkPSR7dGhpcy5kaXNhYmxlZH0+XG4gICAgICAgIDwvc3Bhbj5cbiAgICAgIGAgOiBub3RoaW5nfVxuICAgICAgPHNwYW4gaWQ9XCJ0ZXh0XCI+PHNsb3Q+PC9zbG90Pjwvc3Bhbj5cbiAgICAgICR7dGhpcy5kZXNjcmlwdGlvbiA/IGh0bWxgXG4gICAgICAgIDxzcGFuIGlkPVwiZGVzY3JpcHRpb25cIj4ke3RoaXMuZGVzY3JpcHRpb259PC9zcGFuPlxuICAgICAgYCA6IG5vdGhpbmd9XG4gICAgYDtcbiAgfVxuXG4gICNoYW5kbGVDbGljayA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBpZiAodGhpcy5kaXNhYmxlZCkge1xuICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgIGV2ZW50LnN0b3BJbW1lZGlhdGVQcm9wYWdhdGlvbigpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICBpZiAodGhpcy52YXJpYW50ID09PSAnY2hlY2tib3gnKSB7XG4gICAgICB0aGlzLmNoZWNrZWQgPSAhdGhpcy5jaGVja2VkO1xuICAgIH1cbiAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IFBmTWVudUl0ZW1TZWxlY3RFdmVudCh0aGlzLnZhbHVlLCB0aGlzLmNoZWNrZWQpKTtcbiAgfTtcblxuICAjaGFuZGxlS2V5ZG93biA9IChldmVudDogS2V5Ym9hcmRFdmVudCkgPT4ge1xuICAgIGlmICh0aGlzLmRpc2FibGVkKSByZXR1cm47XG4gICAgaWYgKGV2ZW50LmtleSA9PT0gJyAnIHx8IGV2ZW50LmtleSA9PT0gJ0VudGVyJykge1xuICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgIHRoaXMuI2hhbmRsZUNsaWNrKGV2ZW50KTtcbiAgICB9XG4gIH07XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LW1lbnUtaXRlbSc6IFBmVjZNZW51SXRlbTtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuXFxcXG4gIC0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tYWN0aW9uLS1wbGFpbi0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLWhvdmVyLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1hY3Rpb24tLXBsYWluLS1ob3Zlcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLUJvcmRlckNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWhpZ2gtY29udHJhc3QpO1xcXFxuICAtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1UcmFuc2l0aW9uRHVyYXRpb246IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpO1xcXFxuICAtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1UcmFuc2l0aW9uVGltaW5nRnVuY3Rpb246IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1UcmFuc2l0aW9uUHJvcGVydHk6IGJhY2tncm91bmQtY29sb3I7XFxcXG4gIC0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLWhvdmVyLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWhvdmVyKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2l0ZW0tLVBhZGRpbmdCbG9ja1N0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICAtLXBmLXY2LWMtbWVudV9faXRlbS0tUGFkZGluZ0Jsb2NrRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICAtLXBmLXY2LWMtbWVudV9faXRlbS0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLXBmLXY2LWMtbWVudV9faXRlbS0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2l0ZW0tLUZvbnRTaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudV9faXRlbS0tTGluZUhlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSk7XFxcXG4gIC0tcGYtdjYtYy1tZW51X19pdGVtLS1Gb250V2VpZ2h0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXdlaWdodC0tYm9keS0tZGVmYXVsdCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51X19pdGVtLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51X19pdGVtLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbWVudV9faXRlbS0tbS1kaXNhYmxlZC0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLWRpc2FibGVkKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2l0ZW0tbWFpbi0tQ29sdW1uR2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICAtLXBmLXY2LWMtbWVudV9faXRlbS1kZXNjcmlwdGlvbi0tRm9udFNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxuICAtLXBmLXY2LWMtbWVudV9faXRlbS1kZXNjcmlwdGlvbi0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1PdXRsaW5lT2Zmc2V0OiBjYWxjKHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tY29udHJvbC0tZGVmYXVsdCkgKiAtMyk7XFxcXG5cXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBhbGlnbi1pdGVtczogYmFzZWxpbmU7XFxcXG4gIG1pbi13aWR0aDogMDtcXFxcbiAgcGFkZGluZy1ibG9jay1zdGFydDogdmFyKC0tcGYtdjYtYy1tZW51X19pdGVtLS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1wZi12Ni1jLW1lbnVfX2l0ZW0tLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLW1lbnVfX2l0ZW0tLVBhZGRpbmdJbmxpbmVTdGFydCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLWVuZDogdmFyKC0tcGYtdjYtYy1tZW51X19pdGVtLS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi12Ni1jLW1lbnVfX2l0ZW0tLUZvbnRTaXplKTtcXFxcbiAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXY2LWMtbWVudV9faXRlbS0tRm9udFdlaWdodCk7XFxcXG4gIGxpbmUtaGVpZ2h0OiB2YXIoLS1wZi12Ni1jLW1lbnVfX2l0ZW0tLUxpbmVIZWlnaHQpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdjYtYy1tZW51X19pdGVtLS1Db2xvcik7XFxcXG4gIHRleHQtYWxpZ246IHN0YXJ0O1xcXFxuICB0ZXh0LWRlY29yYXRpb24tbGluZTogbm9uZTtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdHJhbnNwYXJlbnQ7XFxcXG4gIGJvcmRlcjogMDtcXFxcbiAgb3V0bGluZS1vZmZzZXQ6IHZhcigtLXBmLXY2LWMtbWVudS0tT3V0bGluZU9mZnNldCk7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG4gIHRyYW5zaXRpb24tdGltaW5nLWZ1bmN0aW9uOiB2YXIoLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tVHJhbnNpdGlvblRpbWluZ0Z1bmN0aW9uKTtcXFxcbiAgdHJhbnNpdGlvbi1kdXJhdGlvbjogdmFyKC0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLVRyYW5zaXRpb25EdXJhdGlvbik7XFxcXG4gIHRyYW5zaXRpb24tcHJvcGVydHk6IHZhcigtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1UcmFuc2l0aW9uUHJvcGVydHkpO1xcXFxuICBnYXA6IHZhcigtLXBmLXY2LWMtbWVudV9faXRlbS1tYWluLS1Db2x1bW5HYXApO1xcXFxufVxcXFxuXFxcXG46aG9zdDo6YmVmb3JlIHtcXFxcbiAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICBpbnNldDogMDtcXFxcbiAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcIjtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIGJvcmRlci1ibG9jay1zdGFydDogdmFyKC0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQm9yZGVyQ29sb3IpO1xcXFxuICBib3JkZXItYmxvY2stZW5kOiB2YXIoLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQm9yZGVyV2lkdGgpIHNvbGlkIHZhcigtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1Cb3JkZXJDb2xvcik7XFxcXG4gIHRyYW5zaXRpb246IGluaGVyaXQ7XFxcXG59XFxcXG5cXFxcbjpob3N0KDpob3Zlcjpub3QoW2Rpc2FibGVkXSkpLFxcXFxuOmhvc3QoOmZvY3VzLXZpc2libGUpIHtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0taG92ZXItLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLUJvcmRlcldpZHRoOiB2YXIoLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0taG92ZXItLUJvcmRlcldpZHRoKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoOmZvY3VzLXZpc2libGUpIHtcXFxcbiAgb3V0bGluZTogMnB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgb3V0bGluZS1vZmZzZXQ6IC0ycHg7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtkaXNhYmxlZF0pIHtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2l0ZW0tLUNvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnVfX2l0ZW0tLW0tZGlzYWJsZWQtLUNvbG9yKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQmFja2dyb3VuZENvbG9yOiB0cmFuc3BhcmVudDtcXFxcbiAgLS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0taG92ZXItLUJhY2tncm91bmRDb2xvcjogdHJhbnNwYXJlbnQ7XFxcXG4gIGN1cnNvcjogbm90LWFsbG93ZWQ7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxufVxcXFxuXFxcXG4jY2hlY2ssXFxcXG4jdGV4dCxcXFxcbiNkZXNjcmlwdGlvbiB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgei1pbmRleDogMTtcXFxcbn1cXFxcblxcXFxuI2NoZWNrIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgZmxleC1zaHJpbms6IDA7XFxcXG5cXFxcbiAgXFxcXHUwMDI2IC5wZi12Ni1jLWNoZWNrIHtcXFxcbiAgICAtLXBmLXY2LWMtY2hlY2tfX2lucHV0LS1UcmFuc2xhdGVZOiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiN0ZXh0IHtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgdGV4dC1vdmVyZmxvdzogZWxsaXBzaXM7XFxcXG4gIHdoaXRlLXNwYWNlOiBub3dyYXA7XFxcXG4gIGZsZXgtZ3JvdzogMTtcXFxcbn1cXFxcblxcXFxuI2Rlc2NyaXB0aW9uIHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdjYtYy1tZW51X19pdGVtLWRlc2NyaXB0aW9uLS1Gb250U2l6ZSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnVfX2l0ZW0tZGVzY3JpcHRpb24tLUNvbG9yKTtcXFxcbiAgd29yZC1icmVhazogYnJlYWstd29yZDtcXFxcbiAgbWFyZ2luLWJsb2NrLXN0YXJ0OiAwLjI1cmVtO1xcXFxufVxcXFxuXFxcXG5zbG90IHtcXFxcbiAgZGlzcGxheTogaW5saW5lO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksTUFBTSxlQUFlO0FBQzFDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUs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
