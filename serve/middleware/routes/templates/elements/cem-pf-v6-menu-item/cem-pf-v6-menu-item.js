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

// elements/cem-pf-v6-menu-item/cem-pf-v6-menu-item.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-menu-item/cem-pf-v6-menu-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-menu__list-item--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu__list-item--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-menu__list-item--BorderWidth: var(--pf-t--global--border--width--action--plain--default);\\n  --cem-pf-v6-c-menu__list-item--hover--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --cem-pf-v6-c-menu__list-item--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-menu__list-item--TransitionDuration: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-menu__list-item--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-menu__list-item--TransitionProperty: background-color;\\n  --cem-pf-v6-c-menu__list-item--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --cem-pf-v6-c-menu__item--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu__item--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu__item--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-menu__item--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-menu__item--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-menu__item--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-menu__item--FontWeight: var(--pf-t--global--font--weight--body--default);\\n  --cem-pf-v6-c-menu__item--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu__item--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-menu__item--m-disabled--Color: var(--pf-t--global--text--color--disabled);\\n  --cem-pf-v6-c-menu__item-main--ColumnGap: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu__item-description--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --cem-pf-v6-c-menu__item-description--Color: var(--pf-t--global--text--color--subtle);\\n  --cem-pf-v6-c-menu--OutlineOffset: calc(var(--pf-t--global--border--width--control--default) * -3);\\n\\n  position: relative;\\n  display: flex;\\n  align-items: baseline;\\n  min-width: 0;\\n  padding-block-start: var(--cem-pf-v6-c-menu__item--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-menu__item--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-menu__item--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-menu__item--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-menu__item--FontSize);\\n  font-weight: var(--cem-pf-v6-c-menu__item--FontWeight);\\n  line-height: var(--cem-pf-v6-c-menu__item--LineHeight);\\n  color: var(--cem-pf-v6-c-menu__item--Color);\\n  text-align: start;\\n  text-decoration-line: none;\\n  background-color: transparent;\\n  border: 0;\\n  outline-offset: var(--cem-pf-v6-c-menu--OutlineOffset);\\n  cursor: pointer;\\n  user-select: none;\\n  transition-timing-function: var(--cem-pf-v6-c-menu__list-item--TransitionTimingFunction);\\n  transition-duration: var(--cem-pf-v6-c-menu__list-item--TransitionDuration);\\n  transition-property: var(--cem-pf-v6-c-menu__list-item--TransitionProperty);\\n  gap: var(--cem-pf-v6-c-menu__item-main--ColumnGap);\\n}\\n\\n:host::before {\\n  position: absolute;\\n  inset: 0;\\n  content: \\"\\";\\n  background-color: var(--cem-pf-v6-c-menu__list-item--BackgroundColor);\\n  border-block-start: var(--cem-pf-v6-c-menu__list-item--BorderWidth) solid var(--cem-pf-v6-c-menu__list-item--BorderColor);\\n  border-block-end: var(--cem-pf-v6-c-menu__list-item--BorderWidth) solid var(--cem-pf-v6-c-menu__list-item--BorderColor);\\n  transition: inherit;\\n}\\n\\n:host(:hover:not([disabled])),\\n:host(:focus-visible) {\\n  --cem-pf-v6-c-menu__list-item--BackgroundColor: var(--cem-pf-v6-c-menu__list-item--hover--BackgroundColor);\\n  --cem-pf-v6-c-menu__list-item--BorderWidth: var(--cem-pf-v6-c-menu__list-item--hover--BorderWidth);\\n}\\n\\n:host(:focus-visible) {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: -2px;\\n}\\n\\n:host([disabled]) {\\n  --cem-pf-v6-c-menu__item--Color: var(--cem-pf-v6-c-menu__item--m-disabled--Color);\\n  --cem-pf-v6-c-menu__list-item--BackgroundColor: transparent;\\n  --cem-pf-v6-c-menu__list-item--hover--BackgroundColor: transparent;\\n  cursor: not-allowed;\\n  pointer-events: none;\\n}\\n\\n#check,\\n#text,\\n#description {\\n  position: relative;\\n  z-index: 1;\\n}\\n\\n#check {\\n  display: flex;\\n  align-items: center;\\n  flex-shrink: 0;\\n\\n  \\u0026 .cem-pf-v6-c-check {\\n    --cem-pf-v6-c-check__input--TranslateY: none;\\n  }\\n}\\n\\n#text {\\n  overflow: hidden;\\n  text-overflow: ellipsis;\\n  white-space: nowrap;\\n  flex-grow: 1;\\n}\\n\\n#description {\\n  display: block;\\n  font-size: var(--cem-pf-v6-c-menu__item-description--FontSize);\\n  color: var(--cem-pf-v6-c-menu__item-description--Color);\\n  word-break: break-word;\\n  margin-block-start: 0.25rem;\\n}\\n\\nslot {\\n  display: inline;\\n}\\n"'));
var cem_pf_v6_menu_item_default = s;

// elements/cem-pf-v6-menu-item/cem-pf-v6-menu-item.ts
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
_PfV6MenuItem_decorators = [customElement("cem-pf-v6-menu-item")];
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
              class="cem-pf-v6-c-check pf-m-standalone">
          <input id="input"
                 class="cem-pf-v6-c-check__input"
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
__publicField(PfV6MenuItem, "styles", cem_pf_v6_menu_item_default);
__runInitializers(_init, 1, PfV6MenuItem);
export {
  PfMenuItemSelectEvent,
  PfV6MenuItem
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LW1lbnUtaXRlbS9jZW0tcGYtdjYtbWVudS1pdGVtLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LW1lbnUtaXRlbS9jZW0tcGYtdjYtbWVudS1pdGVtLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tcGYtdjYtbWVudS1pdGVtLmNzcyc7XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gbWVudSBpdGVtIGlzIHNlbGVjdGVkL2NoZWNrZWRcbiAqL1xuZXhwb3J0IGNsYXNzIFBmTWVudUl0ZW1TZWxlY3RFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgdmFsdWU6IHN0cmluZztcbiAgY2hlY2tlZDogYm9vbGVhbjtcbiAgY29uc3RydWN0b3IodmFsdWU6IHN0cmluZywgY2hlY2tlZDogYm9vbGVhbikge1xuICAgIHN1cGVyKCdzZWxlY3QnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gICAgdGhpcy52YWx1ZSA9IHZhbHVlO1xuICAgIHRoaXMuY2hlY2tlZCA9IGNoZWNrZWQ7XG4gIH1cbn1cblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IE1lbnUgSXRlbVxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBpdGVtIHRleHRcbiAqXG4gKiBAZmlyZXMgc2VsZWN0IC0gRmlyZXMgd2hlbiBpdGVtIGlzIHNlbGVjdGVkL2NoZWNrZWRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1wZi12Ni1tZW51LWl0ZW0nKVxuZXhwb3J0IGNsYXNzIFBmVjZNZW51SXRlbSBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gICNpbnRlcm5hbHMgPSB0aGlzLmF0dGFjaEludGVybmFscygpO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZGlzYWJsZWQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGNoZWNrZWQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIHZhcmlhbnQ6ICdkZWZhdWx0JyB8ICdjaGVja2JveCcgPSAnZGVmYXVsdCc7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciB2YWx1ZSA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgZGVzY3JpcHRpb24/OiBzdHJpbmc7XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgdGhpcy4jaGFuZGxlQ2xpY2spO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI2hhbmRsZUtleWRvd24pO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgdGhpcy4jaGFuZGxlQ2xpY2spO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI2hhbmRsZUtleWRvd24pO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygndmFyaWFudCcpKSB7XG4gICAgICB0aGlzLiNpbnRlcm5hbHMucm9sZSA9IHRoaXMudmFyaWFudCA9PT0gJ2NoZWNrYm94JyA/ICdtZW51aXRlbWNoZWNrYm94JyA6ICdtZW51aXRlbSc7XG4gICAgfVxuICAgIGlmIChjaGFuZ2VkLmhhcygnY2hlY2tlZCcpIHx8IGNoYW5nZWQuaGFzKCd2YXJpYW50JykpIHtcbiAgICAgIHRoaXMuI2ludGVybmFscy5hcmlhQ2hlY2tlZCA9IHRoaXMudmFyaWFudCA9PT0gJ2NoZWNrYm94JyA/IFN0cmluZyh0aGlzLmNoZWNrZWQpIDogbnVsbDtcbiAgICB9XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdkaXNhYmxlZCcpKSB7XG4gICAgICB0aGlzLiNpbnRlcm5hbHMuYXJpYURpc2FibGVkID0gdGhpcy5kaXNhYmxlZCA/ICd0cnVlJyA6IG51bGw7XG4gICAgICB0aGlzLnNldEF0dHJpYnV0ZSgndGFiaW5kZXgnLCAnLTEnKTtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICAke3RoaXMudmFyaWFudCA9PT0gJ2NoZWNrYm94JyA/IGh0bWxgXG4gICAgICAgIDxzcGFuIGlkPVwiY2hlY2tcIlxuICAgICAgICAgICAgICBjbGFzcz1cImNlbS1wZi12Ni1jLWNoZWNrIHBmLW0tc3RhbmRhbG9uZVwiPlxuICAgICAgICAgIDxpbnB1dCBpZD1cImlucHV0XCJcbiAgICAgICAgICAgICAgICAgY2xhc3M9XCJjZW0tcGYtdjYtYy1jaGVja19faW5wdXRcIlxuICAgICAgICAgICAgICAgICByb2xlPVwicHJlc2VudGF0aW9uXCJcbiAgICAgICAgICAgICAgICAgdGFiaW5kZXg9XCItMVwiXG4gICAgICAgICAgICAgICAgIHR5cGU9XCJjaGVja2JveFwiXG4gICAgICAgICAgICAgICAgIC5jaGVja2VkPSR7dGhpcy5jaGVja2VkfVxuICAgICAgICAgICAgICAgICA/ZGlzYWJsZWQ9JHt0aGlzLmRpc2FibGVkfT5cbiAgICAgICAgPC9zcGFuPlxuICAgICAgYCA6IG5vdGhpbmd9XG4gICAgICA8c3BhbiBpZD1cInRleHRcIj48c2xvdD48L3Nsb3Q+PC9zcGFuPlxuICAgICAgJHt0aGlzLmRlc2NyaXB0aW9uID8gaHRtbGBcbiAgICAgICAgPHNwYW4gaWQ9XCJkZXNjcmlwdGlvblwiPiR7dGhpcy5kZXNjcmlwdGlvbn08L3NwYW4+XG4gICAgICBgIDogbm90aGluZ31cbiAgICBgO1xuICB9XG5cbiAgI2hhbmRsZUNsaWNrID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGlmICh0aGlzLmRpc2FibGVkKSB7XG4gICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgZXZlbnQuc3RvcEltbWVkaWF0ZVByb3BhZ2F0aW9uKCk7XG4gICAgICByZXR1cm47XG4gICAgfVxuICAgIGlmICh0aGlzLnZhcmlhbnQgPT09ICdjaGVja2JveCcpIHtcbiAgICAgIHRoaXMuY2hlY2tlZCA9ICF0aGlzLmNoZWNrZWQ7XG4gICAgfVxuICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgUGZNZW51SXRlbVNlbGVjdEV2ZW50KHRoaXMudmFsdWUsIHRoaXMuY2hlY2tlZCkpO1xuICB9O1xuXG4gICNoYW5kbGVLZXlkb3duID0gKGV2ZW50OiBLZXlib2FyZEV2ZW50KSA9PiB7XG4gICAgaWYgKHRoaXMuZGlzYWJsZWQpIHJldHVybjtcbiAgICBpZiAoZXZlbnQua2V5ID09PSAnICcgfHwgZXZlbnQua2V5ID09PSAnRW50ZXInKSB7XG4gICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgdGhpcy4jaGFuZGxlQ2xpY2soZXZlbnQpO1xuICAgIH1cbiAgfTtcbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXBmLXY2LW1lbnUtaXRlbSc6IFBmVjZNZW51SXRlbTtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuXFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tYWN0aW9uLS1wbGFpbi0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1ob3Zlci0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tYWN0aW9uLS1wbGFpbi0taG92ZXIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0taGlnaC1jb250cmFzdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1UcmFuc2l0aW9uRHVyYXRpb246IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tVHJhbnNpdGlvblRpbWluZ0Z1bmN0aW9uOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLVRyYW5zaXRpb25Qcm9wZXJ0eTogYmFja2dyb3VuZC1jb2xvcjtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLWhvdmVyLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWhvdmVyKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1QYWRkaW5nQmxvY2tFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2l0ZW0tLVBhZGRpbmdJbmxpbmVFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS0tRm9udFNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS0tTGluZUhlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS0tRm9udFdlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS13ZWlnaHQtLWJvZHktLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2l0ZW0tLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2l0ZW0tLW0tZGlzYWJsZWQtLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1kaXNhYmxlZCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS1tYWluLS1Db2x1bW5HYXA6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS1kZXNjcmlwdGlvbi0tRm9udFNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2l0ZW0tZGVzY3JpcHRpb24tLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtLU91dGxpbmVPZmZzZXQ6IGNhbGModmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1jb250cm9sLS1kZWZhdWx0KSAqIC0zKTtcXFxcblxcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBiYXNlbGluZTtcXFxcbiAgbWluLXdpZHRoOiAwO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1Gb250U2l6ZSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1Gb250V2VpZ2h0KTtcXFxcbiAgbGluZS1oZWlnaHQ6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnVfX2l0ZW0tLUxpbmVIZWlnaHQpO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS0tQ29sb3IpO1xcXFxuICB0ZXh0LWFsaWduOiBzdGFydDtcXFxcbiAgdGV4dC1kZWNvcmF0aW9uLWxpbmU6IG5vbmU7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHRyYW5zcGFyZW50O1xcXFxuICBib3JkZXI6IDA7XFxcXG4gIG91dGxpbmUtb2Zmc2V0OiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LS1PdXRsaW5lT2Zmc2V0KTtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcbiAgdHJhbnNpdGlvbi10aW1pbmctZnVuY3Rpb246IHZhcigtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tVHJhbnNpdGlvblRpbWluZ0Z1bmN0aW9uKTtcXFxcbiAgdHJhbnNpdGlvbi1kdXJhdGlvbjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1UcmFuc2l0aW9uRHVyYXRpb24pO1xcXFxuICB0cmFuc2l0aW9uLXByb3BlcnR5OiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLVRyYW5zaXRpb25Qcm9wZXJ0eSk7XFxcXG4gIGdhcDogdmFyKC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS1tYWluLS1Db2x1bW5HYXApO1xcXFxufVxcXFxuXFxcXG46aG9zdDo6YmVmb3JlIHtcXFxcbiAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICBpbnNldDogMDtcXFxcbiAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcIjtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudV9fbGlzdC1pdGVtLS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICBib3JkZXItYmxvY2stc3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQm9yZGVyV2lkdGgpIHNvbGlkIHZhcigtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQm9yZGVyQ29sb3IpO1xcXFxuICBib3JkZXItYmxvY2stZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLUJvcmRlckNvbG9yKTtcXFxcbiAgdHJhbnNpdGlvbjogaW5oZXJpdDtcXFxcbn1cXFxcblxcXFxuOmhvc3QoOmhvdmVyOm5vdChbZGlzYWJsZWRdKSksXFxcXG46aG9zdCg6Zm9jdXMtdmlzaWJsZSkge1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLWhvdmVyLS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQm9yZGVyV2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0taG92ZXItLUJvcmRlcldpZHRoKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoOmZvY3VzLXZpc2libGUpIHtcXFxcbiAgb3V0bGluZTogMnB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgb3V0bGluZS1vZmZzZXQ6IC0ycHg7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtkaXNhYmxlZF0pIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51X19pdGVtLS1Db2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS0tbS1kaXNhYmxlZC0tQ29sb3IpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnVfX2xpc3QtaXRlbS0tQmFja2dyb3VuZENvbG9yOiB0cmFuc3BhcmVudDtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51X19saXN0LWl0ZW0tLWhvdmVyLS1CYWNrZ3JvdW5kQ29sb3I6IHRyYW5zcGFyZW50O1xcXFxuICBjdXJzb3I6IG5vdC1hbGxvd2VkO1xcXFxuICBwb2ludGVyLWV2ZW50czogbm9uZTtcXFxcbn1cXFxcblxcXFxuI2NoZWNrLFxcXFxuI3RleHQsXFxcXG4jZGVzY3JpcHRpb24ge1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIHotaW5kZXg6IDE7XFxcXG59XFxcXG5cXFxcbiNjaGVjayB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGZsZXgtc2hyaW5rOiAwO1xcXFxuXFxcXG4gIFxcXFx1MDAyNiAuY2VtLXBmLXY2LWMtY2hlY2sge1xcXFxuICAgIC0tY2VtLXBmLXY2LWMtY2hlY2tfX2lucHV0LS1UcmFuc2xhdGVZOiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiN0ZXh0IHtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgdGV4dC1vdmVyZmxvdzogZWxsaXBzaXM7XFxcXG4gIHdoaXRlLXNwYWNlOiBub3dyYXA7XFxcXG4gIGZsZXgtZ3JvdzogMTtcXFxcbn1cXFxcblxcXFxuI2Rlc2NyaXB0aW9uIHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS1kZXNjcmlwdGlvbi0tRm9udFNpemUpO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudV9faXRlbS1kZXNjcmlwdGlvbi0tQ29sb3IpO1xcXFxuICB3b3JkLWJyZWFrOiBicmVhay13b3JkO1xcXFxuICBtYXJnaW4tYmxvY2stc3RhcnQ6IDAuMjVyZW07XFxcXG59XFxcXG5cXFxcbnNsb3Qge1xcXFxuICBkaXNwbGF5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
