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

// elements/cem-pf-v6-popover/cem-pf-v6-popover.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-popover/cem-pf-v6-popover.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: inline-block;\\n  position: relative;\\n  align-self: start;\\n  justify-self: start;\\n\\n  --distance: 8px;\\n  --min-width: 100px;\\n  --max-width: 300px;\\n\\n  --cem-pf-v6-c-popover--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --cem-pf-v6-c-popover--BoxShadow: var(--pf-t--global--box-shadow--md);\\n  --cem-pf-v6-c-popover--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-popover--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-popover--BorderRadius: var(--pf-t--global--border--radius--medium);\\n  --cem-pf-v6-c-popover__content--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --cem-pf-v6-c-popover__content--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__content--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__content--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__content--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__close--InsetBlockStart: calc(var(--cem-pf-v6-c-popover__content--PaddingBlockStart) - var(--pf-t--global--spacer--control--vertical--default));\\n  --cem-pf-v6-c-popover__close--InsetInlineEnd: var(--cem-pf-v6-c-popover__content--PaddingInlineEnd);\\n  --cem-pf-v6-c-popover__close--sibling--PaddingInlineEnd: var(--pf-t--global--spacer--2xl);\\n  --cem-pf-v6-c-popover__header--MarginBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-popover__title-text--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-popover__title-text--FontWeight: var(--pf-t--global--font--weight--body--bold);\\n  --cem-pf-v6-c-popover__title-text--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-popover__footer--MarginBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__arrow--Width: 1.5rem;\\n  --cem-pf-v6-c-popover__arrow--Height: 1.5rem;\\n  --cem-pf-v6-c-popover__arrow--BackgroundColor: var(--cem-pf-v6-c-popover__content--BackgroundColor);\\n  --cem-pf-v6-c-popover__arrow--BoxShadow: var(--cem-pf-v6-c-popover--BoxShadow);\\n}\\n\\n#trigger {\\n  anchor-name: --trigger;\\n  display: inline-block;\\n}\\n\\n#content {\\n  background-color: transparent;\\n  padding: 0;\\n  overflow: initial;\\n  position: fixed;\\n  position-anchor: --trigger;\\n  min-width: var(--min-width);\\n  max-width: var(--max-width);\\n  margin: 0;\\n  font-size: var(--cem-pf-v6-c-popover--FontSize);\\n  border: var(--cem-pf-v6-c-popover--BorderWidth) solid var(--cem-pf-v6-c-popover--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-popover--BorderRadius);\\n  box-shadow: var(--cem-pf-v6-c-popover--BoxShadow);\\n  top: auto;\\n  bottom: auto;\\n  left: auto;\\n  right: auto;\\n}\\n\\n#popover {\\n  position: relative;\\n  padding-block-start: var(--cem-pf-v6-c-popover__content--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-popover__content--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-popover__content--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-popover__content--PaddingInlineEnd);\\n  background-color: var(--cem-pf-v6-c-popover__content--BackgroundColor);\\n  border-radius: var(--cem-pf-v6-c-popover--BorderRadius);\\n}\\n\\n#arrow {\\n  position: absolute;\\n  top: var(--cem-pf-v6-c-popover__arrow--InsetBlockStart, auto);\\n  right: var(--cem-pf-v6-c-popover__arrow--InsetInlineEnd, auto);\\n  bottom: var(--cem-pf-v6-c-popover__arrow--InsetBlockEnd, auto);\\n  left: var(--cem-pf-v6-c-popover__arrow--InsetInlineStart, auto);\\n  width: var(--cem-pf-v6-c-popover__arrow--Width);\\n  height: var(--cem-pf-v6-c-popover__arrow--Height);\\n  pointer-events: none;\\n  background-color: var(--cem-pf-v6-c-popover__arrow--BackgroundColor);\\n  border: var(--cem-pf-v6-c-popover--BorderWidth) solid var(--cem-pf-v6-c-popover--BorderColor);\\n  box-shadow: var(--cem-pf-v6-c-popover__arrow--BoxShadow);\\n  transform: translateX(var(--cem-pf-v6-c-popover__arrow--TranslateX, 0)) translateY(var(--cem-pf-v6-c-popover__arrow--TranslateY, 0)) rotate(var(--cem-pf-v6-c-popover__arrow--Rotate, 0));\\n}\\n\\n#close {\\n  position: absolute;\\n  inset-block-start: var(--cem-pf-v6-c-popover__close--InsetBlockStart);\\n  inset-inline-end: var(--cem-pf-v6-c-popover__close--InsetInlineEnd);\\n  padding: var(--pf-t--global--spacer--control--vertical--default) var(--pf-t--global--spacer--control--horizontal--plain--default);\\n  background: transparent;\\n  border: none;\\n  cursor: pointer;\\n  color: var(--pf-t--global--icon--color--regular);\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n\\n  \\u0026:hover {\\n    background-color: var(--pf-t--global--background--color--action--plain--hover);\\n  }\\n}\\n\\n:host(:not([closeable])) #close,\\n:host([closeable=\\"false\\"]) #close {\\n  display: none;\\n}\\n\\n:host([closeable]) #header,\\n:host([closeable=\\"\\"]) #header {\\n  padding-inline-end: var(--cem-pf-v6-c-popover__close--sibling--PaddingInlineEnd);\\n}\\n\\n#header {\\n  margin-block-end: 0;\\n\\n  \\u0026.has-content {\\n    margin-block-end: var(--cem-pf-v6-c-popover__header--MarginBlockEnd);\\n  }\\n}\\n\\n#header-content {\\n  display: flex;\\n  flex: 0 0 auto;\\n\\n  ::slotted(h1),\\n  ::slotted(h2),\\n  ::slotted(h3),\\n  ::slotted(h4),\\n  ::slotted(h5),\\n  ::slotted(h6) {\\n    min-width: 0 !important;\\n    font-size: var(--cem-pf-v6-c-popover__title-text--FontSize) !important;\\n    font-weight: var(--cem-pf-v6-c-popover__title-text--FontWeight) !important;\\n    color: var(--cem-pf-v6-c-popover__title-text--Color) !important;\\n    overflow-wrap: break-word !important;\\n    margin: 0 !important;\\n  }\\n}\\n\\n#body {\\n  word-wrap: break-word;\\n}\\n\\n#footer {\\n  margin-block-start: 0;\\n\\n  \\u0026.has-content {\\n    margin-block-start: var(--cem-pf-v6-c-popover__footer--MarginBlockStart);\\n  }\\n}\\n\\n:host([has-auto-width]) #content {\\n  min-width: unset;\\n  max-width: unset;\\n  width: max-content;\\n}\\n\\n/* Position: top (default) */\\n#content.position-top {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-bottom: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: 50%;\\n  --cem-pf-v6-c-popover__arrow--TranslateX: -50%;\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom {\\n  top: anchor(--trigger bottom);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-top: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: 50%;\\n  --cem-pf-v6-c-popover__arrow--TranslateX: -50%;\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left {\\n  top: anchor(--trigger center);\\n  right: anchor(--trigger left);\\n  translate: 0 -50%;\\n  margin-right: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: 50%;\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--TranslateY: -50%;\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right {\\n  top: anchor(--trigger center);\\n  left: anchor(--trigger right);\\n  translate: 0 -50%;\\n  margin-left: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: 50%;\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--TranslateY: -50%;\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-top-start {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger left);\\n  margin-bottom: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-top-end {\\n  bottom: anchor(--trigger top);\\n  right: anchor(--trigger right);\\n  margin-bottom: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom-start {\\n  top: anchor(--trigger bottom);\\n  left: anchor(--trigger left);\\n  margin-top: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom-end {\\n  top: anchor(--trigger bottom);\\n  right: anchor(--trigger right);\\n  margin-top: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left-start {\\n  top: anchor(--trigger top);\\n  right: anchor(--trigger left);\\n  margin-right: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left-end {\\n  bottom: anchor(--trigger bottom);\\n  right: anchor(--trigger left);\\n  margin-right: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right-start {\\n  top: anchor(--trigger top);\\n  left: anchor(--trigger right);\\n  margin-left: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right-end {\\n  bottom: anchor(--trigger bottom);\\n  left: anchor(--trigger right);\\n  margin-left: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-auto {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-bottom: var(--distance);\\n}\\n"'));
var cem_pf_v6_popover_default = s;

// elements/cem-pf-v6-popover/cem-pf-v6-popover.ts
var PfPopoverShowEvent = class extends Event {
  open = true;
  constructor() {
    super("pf-popover-show", { bubbles: true });
  }
};
var PfPopoverHideEvent = class extends Event {
  open = false;
  constructor() {
    super("pf-popover-hide", { bubbles: true });
  }
};
var supportsAnchorPositioning = globalThis.CSS?.supports?.("anchor-name: --test") ?? false;
var _hasAutoWidth_dec, _maxWidth_dec, _minWidth_dec, _closeButtonLabel_dec, _closeable_dec, _triggerAction_dec, _distance_dec, _position_dec, _open_dec, _a, _PfV6Popover_decorators, _init, _open, _position, _distance, _triggerAction, _closeable, _closeButtonLabel, _minWidth, _maxWidth, _hasAutoWidth, _hoverShowTimeout, _hoverHideTimeout, _hasHeaderContent, _hasFooterContent, _PfV6Popover_instances, contentStyle_get, showPopover_fn, hidePopover_fn, _handleTriggerClick, _handlePointerEnter, _handlePointerLeave, _handlePopoverPointerEnter, _handlePopoverPointerLeave, _handlePopoverToggle, _handleClose, onHeaderSlotChange_fn, onFooterSlotChange_fn;
_PfV6Popover_decorators = [customElement("cem-pf-v6-popover")];
var PfV6Popover = class extends (_a = LitElement, _open_dec = [property({ type: Boolean, reflect: true })], _position_dec = [property({ reflect: true })], _distance_dec = [property({ type: Number, reflect: true })], _triggerAction_dec = [property({ reflect: true, attribute: "trigger-action" })], _closeable_dec = [property({ type: Boolean, reflect: true })], _closeButtonLabel_dec = [property({ attribute: "close-button-label" })], _minWidth_dec = [property({ attribute: "min-width" })], _maxWidth_dec = [property({ attribute: "max-width" })], _hasAutoWidth_dec = [property({ type: Boolean, reflect: true, attribute: "has-auto-width" })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Popover_instances);
    __privateAdd(this, _open, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _position, __runInitializers(_init, 12, this, "top")), __runInitializers(_init, 15, this);
    __privateAdd(this, _distance, __runInitializers(_init, 16, this, 8)), __runInitializers(_init, 19, this);
    __privateAdd(this, _triggerAction, __runInitializers(_init, 20, this, "click")), __runInitializers(_init, 23, this);
    __privateAdd(this, _closeable, __runInitializers(_init, 24, this, true)), __runInitializers(_init, 27, this);
    __privateAdd(this, _closeButtonLabel, __runInitializers(_init, 28, this, "Close")), __runInitializers(_init, 31, this);
    __privateAdd(this, _minWidth, __runInitializers(_init, 32, this)), __runInitializers(_init, 35, this);
    __privateAdd(this, _maxWidth, __runInitializers(_init, 36, this)), __runInitializers(_init, 39, this);
    __privateAdd(this, _hasAutoWidth, __runInitializers(_init, 40, this, false)), __runInitializers(_init, 43, this);
    __privateAdd(this, _hoverShowTimeout);
    __privateAdd(this, _hoverHideTimeout);
    __privateAdd(this, _hasHeaderContent, false);
    __privateAdd(this, _hasFooterContent, false);
    __privateAdd(this, _handleTriggerClick, (e) => {
      const triggerSlot = this.shadowRoot?.querySelector('slot[name="trigger"]');
      const assigned = triggerSlot?.assignedElements();
      if (assigned && assigned.length > 0) {
        e.stopPropagation();
        this.toggle();
      }
    });
    __privateAdd(this, _handlePointerEnter, () => {
      clearTimeout(__privateGet(this, _hoverHideTimeout));
      __privateSet(this, _hoverShowTimeout, setTimeout(() => {
        __privateMethod(this, _PfV6Popover_instances, showPopover_fn).call(this);
      }, 150));
    });
    __privateAdd(this, _handlePointerLeave, () => {
      clearTimeout(__privateGet(this, _hoverShowTimeout));
      __privateSet(this, _hoverHideTimeout, setTimeout(() => {
        __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
      }, 300));
    });
    __privateAdd(this, _handlePopoverPointerEnter, () => {
      clearTimeout(__privateGet(this, _hoverHideTimeout));
    });
    __privateAdd(this, _handlePopoverPointerLeave, () => {
      __privateSet(this, _hoverHideTimeout, setTimeout(() => {
        __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
      }, 300));
    });
    __privateAdd(this, _handlePopoverToggle, (e) => {
      const isOpen = e.newState === "open";
      if (isOpen !== this.open) {
        this.open = isOpen;
      }
      const triggerSlot = this.shadowRoot?.querySelector('slot[name="trigger"]');
      const triggerButton = triggerSlot?.assignedElements()?.[0];
      if (triggerButton) {
        triggerButton.setAttribute("aria-expanded", String(isOpen));
      }
      this.dispatchEvent(isOpen ? new PfPopoverShowEvent() : new PfPopoverHideEvent());
    });
    __privateAdd(this, _handleClose, (e) => {
      e.stopPropagation();
      __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
    });
  }
  render() {
    return html`
      <div id="trigger"
           @click=${this.triggerAction === "click" ? __privateGet(this, _handleTriggerClick) : nothing}
           @pointerenter=${this.triggerAction === "hover" ? __privateGet(this, _handlePointerEnter) : nothing}
           @pointerleave=${this.triggerAction === "hover" ? __privateGet(this, _handlePointerLeave) : nothing}>
        <slot name="trigger"></slot>
      </div>

      <div id="content"
           class="position-${this.position}"
           popover="manual"
           part="popover"
           style=${__privateGet(this, _PfV6Popover_instances, contentStyle_get)}
           @pointerenter=${this.triggerAction === "hover" ? __privateGet(this, _handlePopoverPointerEnter) : nothing}
           @pointerleave=${this.triggerAction === "hover" ? __privateGet(this, _handlePopoverPointerLeave) : nothing}
           @toggle=${__privateGet(this, _handlePopoverToggle)}>
        <div id="arrow"></div>
        <div id="popover"
             part="content">
          <button id="close"
                  part="close-button"
                  aria-label=${this.closeButtonLabel}
                  @click=${__privateGet(this, _handleClose)}>
            <svg class="cem-pf-v6-svg"
                 viewBox="0 0 352 512"
                 fill="currentColor"
                 role="presentation"
                 width="1em"
                 height="1em">
              <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
            </svg>
          </button>

          <header id="header"
                  class=${__privateGet(this, _hasHeaderContent) ? "has-content" : ""}
                  part="header">
            <div id="title">
              <div id="header-content">
                <slot name="header"
                      @slotchange=${__privateMethod(this, _PfV6Popover_instances, onHeaderSlotChange_fn)}></slot>
              </div>
            </div>
          </header>

          <div id="body"
               part="body">
            <slot></slot>
          </div>

          <footer id="footer"
                  class=${__privateGet(this, _hasFooterContent) ? "has-content" : ""}
                  part="footer">
            <slot name="footer"
                  @slotchange=${__privateMethod(this, _PfV6Popover_instances, onFooterSlotChange_fn)}></slot>
          </footer>
        </div>
      </div>
    `;
  }
  updated(changed) {
    if (changed.has("open")) {
      const contentEl = this.shadowRoot?.getElementById("content");
      if (!contentEl) return;
      if (this.open && !contentEl.matches(":popover-open")) {
        __privateMethod(this, _PfV6Popover_instances, showPopover_fn).call(this);
      } else if (!this.open && contentEl.matches(":popover-open")) {
        __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
      }
    }
  }
  toggle() {
    if (this.open) {
      __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
    } else {
      __privateMethod(this, _PfV6Popover_instances, showPopover_fn).call(this);
    }
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    clearTimeout(__privateGet(this, _hoverShowTimeout));
    clearTimeout(__privateGet(this, _hoverHideTimeout));
  }
};
_init = __decoratorStart(_a);
_open = new WeakMap();
_position = new WeakMap();
_distance = new WeakMap();
_triggerAction = new WeakMap();
_closeable = new WeakMap();
_closeButtonLabel = new WeakMap();
_minWidth = new WeakMap();
_maxWidth = new WeakMap();
_hasAutoWidth = new WeakMap();
_hoverShowTimeout = new WeakMap();
_hoverHideTimeout = new WeakMap();
_hasHeaderContent = new WeakMap();
_hasFooterContent = new WeakMap();
_PfV6Popover_instances = new WeakSet();
contentStyle_get = function() {
  const parts = [];
  if (this.minWidth) parts.push(`--min-width: ${this.minWidth}`);
  if (this.maxWidth) parts.push(`--max-width: ${this.maxWidth}`);
  parts.push(`--distance: ${this.distance}px`);
  return parts.join("; ");
};
showPopover_fn = async function() {
  const contentEl = this.shadowRoot?.getElementById("content");
  if (!contentEl) return;
  try {
    contentEl.showPopover();
    if (!supportsAnchorPositioning) {
      await new Promise((resolve) => requestAnimationFrame(resolve));
      const trigger = this.shadowRoot?.getElementById("trigger");
      if (!trigger) return;
      const triggerRect = trigger.getBoundingClientRect();
      const popoverRect = contentEl.getBoundingClientRect();
      const dist = this.distance;
      let top, left;
      switch (this.position) {
        case "top":
          top = triggerRect.top - popoverRect.height - dist;
          left = triggerRect.left + triggerRect.width / 2 - popoverRect.width / 2;
          break;
        case "top-start":
          top = triggerRect.top - popoverRect.height - dist;
          left = triggerRect.left;
          break;
        case "top-end":
          top = triggerRect.top - popoverRect.height - dist;
          left = triggerRect.right - popoverRect.width;
          break;
        case "bottom":
          top = triggerRect.bottom + dist;
          left = triggerRect.left + triggerRect.width / 2 - popoverRect.width / 2;
          break;
        case "bottom-start":
          top = triggerRect.bottom + dist;
          left = triggerRect.left;
          break;
        case "bottom-end":
          top = triggerRect.bottom + dist;
          left = triggerRect.right - popoverRect.width;
          break;
        case "left":
          top = triggerRect.top + triggerRect.height / 2 - popoverRect.height / 2;
          left = triggerRect.left - popoverRect.width - dist;
          break;
        case "left-start":
          top = triggerRect.top;
          left = triggerRect.left - popoverRect.width - dist;
          break;
        case "left-end":
          top = triggerRect.bottom - popoverRect.height;
          left = triggerRect.left - popoverRect.width - dist;
          break;
        case "right":
          top = triggerRect.top + triggerRect.height / 2 - popoverRect.height / 2;
          left = triggerRect.right + dist;
          break;
        case "right-start":
          top = triggerRect.top;
          left = triggerRect.right + dist;
          break;
        case "right-end":
          top = triggerRect.bottom - popoverRect.height;
          left = triggerRect.right + dist;
          break;
        default:
          top = triggerRect.top - popoverRect.height - dist;
          left = triggerRect.left + triggerRect.width / 2 - popoverRect.width / 2;
      }
      contentEl.style.top = `${top}px`;
      contentEl.style.left = `${left}px`;
      contentEl.style.bottom = "auto";
      contentEl.style.right = "auto";
      contentEl.style.translate = "none";
    }
  } catch (e) {
    console.warn("Failed to show popover:", e);
  }
};
hidePopover_fn = function() {
  const contentEl = this.shadowRoot?.getElementById("content");
  try {
    contentEl?.hidePopover();
  } catch (e) {
    console.warn("Failed to hide popover:", e);
  }
};
_handleTriggerClick = new WeakMap();
_handlePointerEnter = new WeakMap();
_handlePointerLeave = new WeakMap();
_handlePopoverPointerEnter = new WeakMap();
_handlePopoverPointerLeave = new WeakMap();
_handlePopoverToggle = new WeakMap();
_handleClose = new WeakMap();
onHeaderSlotChange_fn = function(e) {
  const slot = e.target;
  __privateSet(this, _hasHeaderContent, slot.assignedNodes().length > 0);
  this.requestUpdate();
};
onFooterSlotChange_fn = function(e) {
  const slot = e.target;
  __privateSet(this, _hasFooterContent, slot.assignedNodes().length > 0);
  this.requestUpdate();
};
__decorateElement(_init, 4, "open", _open_dec, PfV6Popover, _open);
__decorateElement(_init, 4, "position", _position_dec, PfV6Popover, _position);
__decorateElement(_init, 4, "distance", _distance_dec, PfV6Popover, _distance);
__decorateElement(_init, 4, "triggerAction", _triggerAction_dec, PfV6Popover, _triggerAction);
__decorateElement(_init, 4, "closeable", _closeable_dec, PfV6Popover, _closeable);
__decorateElement(_init, 4, "closeButtonLabel", _closeButtonLabel_dec, PfV6Popover, _closeButtonLabel);
__decorateElement(_init, 4, "minWidth", _minWidth_dec, PfV6Popover, _minWidth);
__decorateElement(_init, 4, "maxWidth", _maxWidth_dec, PfV6Popover, _maxWidth);
__decorateElement(_init, 4, "hasAutoWidth", _hasAutoWidth_dec, PfV6Popover, _hasAutoWidth);
PfV6Popover = __decorateElement(_init, 0, "PfV6Popover", _PfV6Popover_decorators, PfV6Popover);
__publicField(PfV6Popover, "styles", cem_pf_v6_popover_default);
__runInitializers(_init, 1, PfV6Popover);
export {
  PfPopoverHideEvent,
  PfPopoverShowEvent,
  PfV6Popover
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXBvcG92ZXIvY2VtLXBmLXY2LXBvcG92ZXIudHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tcGYtdjYtcG9wb3Zlci9jZW0tcGYtdjYtcG9wb3Zlci5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwsIG5vdGhpbmcgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXBmLXY2LXBvcG92ZXIuY3NzJztcblxudHlwZSBQb3BvdmVyUG9zaXRpb24gPVxuICB8ICdhdXRvJyB8ICd0b3AnIHwgJ2JvdHRvbScgfCAnbGVmdCcgfCAncmlnaHQnXG4gIHwgJ3RvcC1zdGFydCcgfCAndG9wLWVuZCcgfCAnYm90dG9tLXN0YXJ0JyB8ICdib3R0b20tZW5kJ1xuICB8ICdsZWZ0LXN0YXJ0JyB8ICdsZWZ0LWVuZCcgfCAncmlnaHQtc3RhcnQnIHwgJ3JpZ2h0LWVuZCc7XG5cbnR5cGUgVHJpZ2dlckFjdGlvbiA9ICdjbGljaycgfCAnaG92ZXInO1xuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIHBvcG92ZXIgc2hvd3NcbiAqL1xuZXhwb3J0IGNsYXNzIFBmUG9wb3ZlclNob3dFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgb3BlbiA9IHRydWU7XG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCdwZi1wb3BvdmVyLXNob3cnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gIH1cbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBwb3BvdmVyIGhpZGVzXG4gKi9cbmV4cG9ydCBjbGFzcyBQZlBvcG92ZXJIaWRlRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIG9wZW4gPSBmYWxzZTtcbiAgY29uc3RydWN0b3IoKSB7XG4gICAgc3VwZXIoJ3BmLXBvcG92ZXItaGlkZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgfVxufVxuXG4vLyBDYWNoZSBmZWF0dXJlIGRldGVjdGlvbiBmb3IgQ1NTIEFuY2hvciBQb3NpdGlvbmluZ1xuY29uc3Qgc3VwcG9ydHNBbmNob3JQb3NpdGlvbmluZyA9IGdsb2JhbFRoaXMuQ1NTPy5zdXBwb3J0cz8uKCdhbmNob3ItbmFtZTogLS10ZXN0JykgPz8gZmFsc2U7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBQb3BvdmVyXG4gKlxuICogQSBwb3BvdmVyIGNvbXBvbmVudCB1c2luZyB0aGUgbmF0aXZlIFBvcG92ZXIgQVBJIGFuZCBDU1MgQW5jaG9yIFBvc2l0aW9uaW5nLlxuICpcbiAqIEBzbG90IHRyaWdnZXIgLSBUaGUgZWxlbWVudCB0aGF0IHRyaWdnZXJzIHRoZSBwb3BvdmVyICh1c3VhbGx5IGEgYnV0dG9uKVxuICogQHNsb3QgaGVhZGVyIC0gSGVhZGVyIGNvbnRlbnRcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBwb3BvdmVyIGJvZHkgY29udGVudFxuICogQHNsb3QgZm9vdGVyIC0gRm9vdGVyIGNvbnRlbnRcbiAqXG4gKiBAZmlyZXMgcGYtcG9wb3Zlci1zaG93IC0gRmlyZWQgd2hlbiBwb3BvdmVyIG9wZW5zXG4gKiBAZmlyZXMgcGYtcG9wb3Zlci1oaWRlIC0gRmlyZWQgd2hlbiBwb3BvdmVyIGNsb3Nlc1xuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXBmLXY2LXBvcG92ZXInKVxuZXhwb3J0IGNsYXNzIFBmVjZQb3BvdmVyIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBvcGVuID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBwb3NpdGlvbjogUG9wb3ZlclBvc2l0aW9uID0gJ3RvcCc7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogTnVtYmVyLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGRpc3RhbmNlID0gODtcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlLCBhdHRyaWJ1dGU6ICd0cmlnZ2VyLWFjdGlvbicgfSlcbiAgYWNjZXNzb3IgdHJpZ2dlckFjdGlvbjogVHJpZ2dlckFjdGlvbiA9ICdjbGljayc7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBjbG9zZWFibGUgPSB0cnVlO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ2Nsb3NlLWJ1dHRvbi1sYWJlbCcgfSlcbiAgYWNjZXNzb3IgY2xvc2VCdXR0b25MYWJlbCA9ICdDbG9zZSc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnbWluLXdpZHRoJyB9KVxuICBhY2Nlc3NvciBtaW5XaWR0aD86IHN0cmluZztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdtYXgtd2lkdGgnIH0pXG4gIGFjY2Vzc29yIG1heFdpZHRoPzogc3RyaW5nO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUsIGF0dHJpYnV0ZTogJ2hhcy1hdXRvLXdpZHRoJyB9KVxuICBhY2Nlc3NvciBoYXNBdXRvV2lkdGggPSBmYWxzZTtcblxuICAjaG92ZXJTaG93VGltZW91dD86IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+O1xuICAjaG92ZXJIaWRlVGltZW91dD86IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+O1xuICAjaGFzSGVhZGVyQ29udGVudCA9IGZhbHNlO1xuICAjaGFzRm9vdGVyQ29udGVudCA9IGZhbHNlO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJ0cmlnZ2VyXCJcbiAgICAgICAgICAgQGNsaWNrPSR7dGhpcy50cmlnZ2VyQWN0aW9uID09PSAnY2xpY2snID8gdGhpcy4jaGFuZGxlVHJpZ2dlckNsaWNrIDogbm90aGluZ31cbiAgICAgICAgICAgQHBvaW50ZXJlbnRlcj0ke3RoaXMudHJpZ2dlckFjdGlvbiA9PT0gJ2hvdmVyJyA/IHRoaXMuI2hhbmRsZVBvaW50ZXJFbnRlciA6IG5vdGhpbmd9XG4gICAgICAgICAgIEBwb2ludGVybGVhdmU9JHt0aGlzLnRyaWdnZXJBY3Rpb24gPT09ICdob3ZlcicgPyB0aGlzLiNoYW5kbGVQb2ludGVyTGVhdmUgOiBub3RoaW5nfT5cbiAgICAgICAgPHNsb3QgbmFtZT1cInRyaWdnZXJcIj48L3Nsb3Q+XG4gICAgICA8L2Rpdj5cblxuICAgICAgPGRpdiBpZD1cImNvbnRlbnRcIlxuICAgICAgICAgICBjbGFzcz1cInBvc2l0aW9uLSR7dGhpcy5wb3NpdGlvbn1cIlxuICAgICAgICAgICBwb3BvdmVyPVwibWFudWFsXCJcbiAgICAgICAgICAgcGFydD1cInBvcG92ZXJcIlxuICAgICAgICAgICBzdHlsZT0ke3RoaXMuI2NvbnRlbnRTdHlsZX1cbiAgICAgICAgICAgQHBvaW50ZXJlbnRlcj0ke3RoaXMudHJpZ2dlckFjdGlvbiA9PT0gJ2hvdmVyJyA/IHRoaXMuI2hhbmRsZVBvcG92ZXJQb2ludGVyRW50ZXIgOiBub3RoaW5nfVxuICAgICAgICAgICBAcG9pbnRlcmxlYXZlPSR7dGhpcy50cmlnZ2VyQWN0aW9uID09PSAnaG92ZXInID8gdGhpcy4jaGFuZGxlUG9wb3ZlclBvaW50ZXJMZWF2ZSA6IG5vdGhpbmd9XG4gICAgICAgICAgIEB0b2dnbGU9JHt0aGlzLiNoYW5kbGVQb3BvdmVyVG9nZ2xlfT5cbiAgICAgICAgPGRpdiBpZD1cImFycm93XCI+PC9kaXY+XG4gICAgICAgIDxkaXYgaWQ9XCJwb3BvdmVyXCJcbiAgICAgICAgICAgICBwYXJ0PVwiY29udGVudFwiPlxuICAgICAgICAgIDxidXR0b24gaWQ9XCJjbG9zZVwiXG4gICAgICAgICAgICAgICAgICBwYXJ0PVwiY2xvc2UtYnV0dG9uXCJcbiAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9JHt0aGlzLmNsb3NlQnV0dG9uTGFiZWx9XG4gICAgICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLiNoYW5kbGVDbG9zZX0+XG4gICAgICAgICAgICA8c3ZnIGNsYXNzPVwiY2VtLXBmLXY2LXN2Z1wiXG4gICAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMzUyIDUxMlwiXG4gICAgICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgICAgICByb2xlPVwicHJlc2VudGF0aW9uXCJcbiAgICAgICAgICAgICAgICAgd2lkdGg9XCIxZW1cIlxuICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxZW1cIj5cbiAgICAgICAgICAgICAgPHBhdGggZD1cIk0yNDIuNzIgMjU2bDEwMC4wNy0xMDAuMDdjMTIuMjgtMTIuMjggMTIuMjgtMzIuMTkgMC00NC40OGwtMjIuMjQtMjIuMjRjLTEyLjI4LTEyLjI4LTMyLjE5LTEyLjI4LTQ0LjQ4IDBMMTc2IDE4OS4yOCA3NS45MyA4OS4yMWMtMTIuMjgtMTIuMjgtMzIuMTktMTIuMjgtNDQuNDggMEw5LjIxIDExMS40NWMtMTIuMjggMTIuMjgtMTIuMjggMzIuMTkgMCA0NC40OEwxMDkuMjggMjU2IDkuMjEgMzU2LjA3Yy0xMi4yOCAxMi4yOC0xMi4yOCAzMi4xOSAwIDQ0LjQ4bDIyLjI0IDIyLjI0YzEyLjI4IDEyLjI4IDMyLjIgMTIuMjggNDQuNDggMEwxNzYgMzIyLjcybDEwMC4wNyAxMDAuMDdjMTIuMjggMTIuMjggMzIuMiAxMi4yOCA0NC40OCAwbDIyLjI0LTIyLjI0YzEyLjI4LTEyLjI4IDEyLjI4LTMyLjE5IDAtNDQuNDhMMjQyLjcyIDI1NnpcIj48L3BhdGg+XG4gICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICA8L2J1dHRvbj5cblxuICAgICAgICAgIDxoZWFkZXIgaWQ9XCJoZWFkZXJcIlxuICAgICAgICAgICAgICAgICAgY2xhc3M9JHt0aGlzLiNoYXNIZWFkZXJDb250ZW50ID8gJ2hhcy1jb250ZW50JyA6ICcnfVxuICAgICAgICAgICAgICAgICAgcGFydD1cImhlYWRlclwiPlxuICAgICAgICAgICAgPGRpdiBpZD1cInRpdGxlXCI+XG4gICAgICAgICAgICAgIDxkaXYgaWQ9XCJoZWFkZXItY29udGVudFwiPlxuICAgICAgICAgICAgICAgIDxzbG90IG5hbWU9XCJoZWFkZXJcIlxuICAgICAgICAgICAgICAgICAgICAgIEBzbG90Y2hhbmdlPSR7dGhpcy4jb25IZWFkZXJTbG90Q2hhbmdlfT48L3Nsb3Q+XG4gICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPC9oZWFkZXI+XG5cbiAgICAgICAgICA8ZGl2IGlkPVwiYm9keVwiXG4gICAgICAgICAgICAgICBwYXJ0PVwiYm9keVwiPlxuICAgICAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgICAgIDwvZGl2PlxuXG4gICAgICAgICAgPGZvb3RlciBpZD1cImZvb3RlclwiXG4gICAgICAgICAgICAgICAgICBjbGFzcz0ke3RoaXMuI2hhc0Zvb3RlckNvbnRlbnQgPyAnaGFzLWNvbnRlbnQnIDogJyd9XG4gICAgICAgICAgICAgICAgICBwYXJ0PVwiZm9vdGVyXCI+XG4gICAgICAgICAgICA8c2xvdCBuYW1lPVwiZm9vdGVyXCJcbiAgICAgICAgICAgICAgICAgIEBzbG90Y2hhbmdlPSR7dGhpcy4jb25Gb290ZXJTbG90Q2hhbmdlfT48L3Nsb3Q+XG4gICAgICAgICAgPC9mb290ZXI+XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gIGdldCAjY29udGVudFN0eWxlKCk6IHN0cmluZyB7XG4gICAgY29uc3QgcGFydHM6IHN0cmluZ1tdID0gW107XG4gICAgaWYgKHRoaXMubWluV2lkdGgpIHBhcnRzLnB1c2goYC0tbWluLXdpZHRoOiAke3RoaXMubWluV2lkdGh9YCk7XG4gICAgaWYgKHRoaXMubWF4V2lkdGgpIHBhcnRzLnB1c2goYC0tbWF4LXdpZHRoOiAke3RoaXMubWF4V2lkdGh9YCk7XG4gICAgcGFydHMucHVzaChgLS1kaXN0YW5jZTogJHt0aGlzLmRpc3RhbmNlfXB4YCk7XG4gICAgcmV0dXJuIHBhcnRzLmpvaW4oJzsgJyk7XG4gIH1cblxuICB1cGRhdGVkKGNoYW5nZWQ6IE1hcDxzdHJpbmcsIHVua25vd24+KSB7XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdvcGVuJykpIHtcbiAgICAgIGNvbnN0IGNvbnRlbnRFbCA9IHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ2NvbnRlbnQnKSBhcyBIVE1MRWxlbWVudCAmIHsgc2hvd1BvcG92ZXIoKTogdm9pZDsgaGlkZVBvcG92ZXIoKTogdm9pZDsgbWF0Y2hlcyhzOiBzdHJpbmcpOiBib29sZWFuIH0gfCBudWxsO1xuICAgICAgaWYgKCFjb250ZW50RWwpIHJldHVybjtcbiAgICAgIGlmICh0aGlzLm9wZW4gJiYgIWNvbnRlbnRFbC5tYXRjaGVzKCc6cG9wb3Zlci1vcGVuJykpIHtcbiAgICAgICAgdGhpcy4jc2hvd1BvcG92ZXIoKTtcbiAgICAgIH0gZWxzZSBpZiAoIXRoaXMub3BlbiAmJiBjb250ZW50RWwubWF0Y2hlcygnOnBvcG92ZXItb3BlbicpKSB7XG4gICAgICAgIHRoaXMuI2hpZGVQb3BvdmVyKCk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgdG9nZ2xlKCkge1xuICAgIGlmICh0aGlzLm9wZW4pIHtcbiAgICAgIHRoaXMuI2hpZGVQb3BvdmVyKCk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI3Nob3dQb3BvdmVyKCk7XG4gICAgfVxuICB9XG5cbiAgYXN5bmMgI3Nob3dQb3BvdmVyKCkge1xuICAgIGNvbnN0IGNvbnRlbnRFbCA9IHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ2NvbnRlbnQnKSBhcyBIVE1MRWxlbWVudCAmIHsgc2hvd1BvcG92ZXIoKTogdm9pZCB9IHwgbnVsbDtcbiAgICBpZiAoIWNvbnRlbnRFbCkgcmV0dXJuO1xuICAgIHRyeSB7XG4gICAgICBjb250ZW50RWwuc2hvd1BvcG92ZXIoKTtcblxuICAgICAgaWYgKCFzdXBwb3J0c0FuY2hvclBvc2l0aW9uaW5nKSB7XG4gICAgICAgIGF3YWl0IG5ldyBQcm9taXNlKHJlc29sdmUgPT4gcmVxdWVzdEFuaW1hdGlvbkZyYW1lKHJlc29sdmUpKTtcblxuICAgICAgICBjb25zdCB0cmlnZ2VyID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgndHJpZ2dlcicpO1xuICAgICAgICBpZiAoIXRyaWdnZXIpIHJldHVybjtcbiAgICAgICAgY29uc3QgdHJpZ2dlclJlY3QgPSB0cmlnZ2VyLmdldEJvdW5kaW5nQ2xpZW50UmVjdCgpO1xuICAgICAgICBjb25zdCBwb3BvdmVyUmVjdCA9IGNvbnRlbnRFbC5nZXRCb3VuZGluZ0NsaWVudFJlY3QoKTtcbiAgICAgICAgY29uc3QgZGlzdCA9IHRoaXMuZGlzdGFuY2U7XG5cbiAgICAgICAgbGV0IHRvcDogbnVtYmVyLCBsZWZ0OiBudW1iZXI7XG5cbiAgICAgICAgc3dpdGNoICh0aGlzLnBvc2l0aW9uKSB7XG4gICAgICAgICAgY2FzZSAndG9wJzpcbiAgICAgICAgICAgIHRvcCA9IHRyaWdnZXJSZWN0LnRvcCAtIHBvcG92ZXJSZWN0LmhlaWdodCAtIGRpc3Q7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QubGVmdCArICh0cmlnZ2VyUmVjdC53aWR0aCAvIDIpIC0gKHBvcG92ZXJSZWN0LndpZHRoIC8gMik7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBjYXNlICd0b3Atc3RhcnQnOlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QudG9wIC0gcG9wb3ZlclJlY3QuaGVpZ2h0IC0gZGlzdDtcbiAgICAgICAgICAgIGxlZnQgPSB0cmlnZ2VyUmVjdC5sZWZ0O1xuICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgY2FzZSAndG9wLWVuZCc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC50b3AgLSBwb3BvdmVyUmVjdC5oZWlnaHQgLSBkaXN0O1xuICAgICAgICAgICAgbGVmdCA9IHRyaWdnZXJSZWN0LnJpZ2h0IC0gcG9wb3ZlclJlY3Qud2lkdGg7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBjYXNlICdib3R0b20nOlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QuYm90dG9tICsgZGlzdDtcbiAgICAgICAgICAgIGxlZnQgPSB0cmlnZ2VyUmVjdC5sZWZ0ICsgKHRyaWdnZXJSZWN0LndpZHRoIC8gMikgLSAocG9wb3ZlclJlY3Qud2lkdGggLyAyKTtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgIGNhc2UgJ2JvdHRvbS1zdGFydCc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC5ib3R0b20gKyBkaXN0O1xuICAgICAgICAgICAgbGVmdCA9IHRyaWdnZXJSZWN0LmxlZnQ7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBjYXNlICdib3R0b20tZW5kJzpcbiAgICAgICAgICAgIHRvcCA9IHRyaWdnZXJSZWN0LmJvdHRvbSArIGRpc3Q7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QucmlnaHQgLSBwb3BvdmVyUmVjdC53aWR0aDtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgIGNhc2UgJ2xlZnQnOlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QudG9wICsgKHRyaWdnZXJSZWN0LmhlaWdodCAvIDIpIC0gKHBvcG92ZXJSZWN0LmhlaWdodCAvIDIpO1xuICAgICAgICAgICAgbGVmdCA9IHRyaWdnZXJSZWN0LmxlZnQgLSBwb3BvdmVyUmVjdC53aWR0aCAtIGRpc3Q7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBjYXNlICdsZWZ0LXN0YXJ0JzpcbiAgICAgICAgICAgIHRvcCA9IHRyaWdnZXJSZWN0LnRvcDtcbiAgICAgICAgICAgIGxlZnQgPSB0cmlnZ2VyUmVjdC5sZWZ0IC0gcG9wb3ZlclJlY3Qud2lkdGggLSBkaXN0O1xuICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgY2FzZSAnbGVmdC1lbmQnOlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QuYm90dG9tIC0gcG9wb3ZlclJlY3QuaGVpZ2h0O1xuICAgICAgICAgICAgbGVmdCA9IHRyaWdnZXJSZWN0LmxlZnQgLSBwb3BvdmVyUmVjdC53aWR0aCAtIGRpc3Q7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBjYXNlICdyaWdodCc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC50b3AgKyAodHJpZ2dlclJlY3QuaGVpZ2h0IC8gMikgLSAocG9wb3ZlclJlY3QuaGVpZ2h0IC8gMik7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QucmlnaHQgKyBkaXN0O1xuICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgY2FzZSAncmlnaHQtc3RhcnQnOlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QudG9wO1xuICAgICAgICAgICAgbGVmdCA9IHRyaWdnZXJSZWN0LnJpZ2h0ICsgZGlzdDtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgIGNhc2UgJ3JpZ2h0LWVuZCc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC5ib3R0b20gLSBwb3BvdmVyUmVjdC5oZWlnaHQ7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QucmlnaHQgKyBkaXN0O1xuICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgZGVmYXVsdDpcbiAgICAgICAgICAgIHRvcCA9IHRyaWdnZXJSZWN0LnRvcCAtIHBvcG92ZXJSZWN0LmhlaWdodCAtIGRpc3Q7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QubGVmdCArICh0cmlnZ2VyUmVjdC53aWR0aCAvIDIpIC0gKHBvcG92ZXJSZWN0LndpZHRoIC8gMik7XG4gICAgICAgIH1cblxuICAgICAgICBjb250ZW50RWwuc3R5bGUudG9wID0gYCR7dG9wfXB4YDtcbiAgICAgICAgY29udGVudEVsLnN0eWxlLmxlZnQgPSBgJHtsZWZ0fXB4YDtcbiAgICAgICAgY29udGVudEVsLnN0eWxlLmJvdHRvbSA9ICdhdXRvJztcbiAgICAgICAgY29udGVudEVsLnN0eWxlLnJpZ2h0ID0gJ2F1dG8nO1xuICAgICAgICBjb250ZW50RWwuc3R5bGUudHJhbnNsYXRlID0gJ25vbmUnO1xuICAgICAgfVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIGNvbnNvbGUud2FybignRmFpbGVkIHRvIHNob3cgcG9wb3ZlcjonLCBlKTtcbiAgICB9XG4gIH1cblxuICAjaGlkZVBvcG92ZXIoKSB7XG4gICAgY29uc3QgY29udGVudEVsID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnY29udGVudCcpIGFzIEhUTUxFbGVtZW50ICYgeyBoaWRlUG9wb3ZlcigpOiB2b2lkIH0gfCBudWxsO1xuICAgIHRyeSB7XG4gICAgICBjb250ZW50RWw/LmhpZGVQb3BvdmVyKCk7XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgY29uc29sZS53YXJuKCdGYWlsZWQgdG8gaGlkZSBwb3BvdmVyOicsIGUpO1xuICAgIH1cbiAgfVxuXG4gICNoYW5kbGVUcmlnZ2VyQ2xpY2sgPSAoZTogRXZlbnQpID0+IHtcbiAgICBjb25zdCB0cmlnZ2VyU2xvdCA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3Rvcignc2xvdFtuYW1lPVwidHJpZ2dlclwiXScpIGFzIEhUTUxTbG90RWxlbWVudCB8IG51bGw7XG4gICAgY29uc3QgYXNzaWduZWQgPSB0cmlnZ2VyU2xvdD8uYXNzaWduZWRFbGVtZW50cygpO1xuICAgIGlmIChhc3NpZ25lZCAmJiBhc3NpZ25lZC5sZW5ndGggPiAwKSB7XG4gICAgICBlLnN0b3BQcm9wYWdhdGlvbigpO1xuICAgICAgdGhpcy50b2dnbGUoKTtcbiAgICB9XG4gIH07XG5cbiAgI2hhbmRsZVBvaW50ZXJFbnRlciA9ICgpID0+IHtcbiAgICBjbGVhclRpbWVvdXQodGhpcy4jaG92ZXJIaWRlVGltZW91dCk7XG4gICAgdGhpcy4jaG92ZXJTaG93VGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgdGhpcy4jc2hvd1BvcG92ZXIoKTtcbiAgICB9LCAxNTApO1xuICB9O1xuXG4gICNoYW5kbGVQb2ludGVyTGVhdmUgPSAoKSA9PiB7XG4gICAgY2xlYXJUaW1lb3V0KHRoaXMuI2hvdmVyU2hvd1RpbWVvdXQpO1xuICAgIHRoaXMuI2hvdmVySGlkZVRpbWVvdXQgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgIHRoaXMuI2hpZGVQb3BvdmVyKCk7XG4gICAgfSwgMzAwKTtcbiAgfTtcblxuICAjaGFuZGxlUG9wb3ZlclBvaW50ZXJFbnRlciA9ICgpID0+IHtcbiAgICBjbGVhclRpbWVvdXQodGhpcy4jaG92ZXJIaWRlVGltZW91dCk7XG4gIH07XG5cbiAgI2hhbmRsZVBvcG92ZXJQb2ludGVyTGVhdmUgPSAoKSA9PiB7XG4gICAgdGhpcy4jaG92ZXJIaWRlVGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgdGhpcy4jaGlkZVBvcG92ZXIoKTtcbiAgICB9LCAzMDApO1xuICB9O1xuXG4gICNoYW5kbGVQb3BvdmVyVG9nZ2xlID0gKGU6IFRvZ2dsZUV2ZW50KSA9PiB7XG4gICAgY29uc3QgaXNPcGVuID0gZS5uZXdTdGF0ZSA9PT0gJ29wZW4nO1xuXG4gICAgaWYgKGlzT3BlbiAhPT0gdGhpcy5vcGVuKSB7XG4gICAgICB0aGlzLm9wZW4gPSBpc09wZW47XG4gICAgfVxuXG4gICAgY29uc3QgdHJpZ2dlclNsb3QgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ3Nsb3RbbmFtZT1cInRyaWdnZXJcIl0nKSBhcyBIVE1MU2xvdEVsZW1lbnQgfCBudWxsO1xuICAgIGNvbnN0IHRyaWdnZXJCdXR0b24gPSB0cmlnZ2VyU2xvdD8uYXNzaWduZWRFbGVtZW50cygpPy5bMF07XG4gICAgaWYgKHRyaWdnZXJCdXR0b24pIHtcbiAgICAgIHRyaWdnZXJCdXR0b24uc2V0QXR0cmlidXRlKCdhcmlhLWV4cGFuZGVkJywgU3RyaW5nKGlzT3BlbikpO1xuICAgIH1cblxuICAgIHRoaXMuZGlzcGF0Y2hFdmVudChpc09wZW4gPyBuZXcgUGZQb3BvdmVyU2hvd0V2ZW50KCkgOiBuZXcgUGZQb3BvdmVySGlkZUV2ZW50KCkpO1xuICB9O1xuXG4gICNoYW5kbGVDbG9zZSA9IChlOiBFdmVudCkgPT4ge1xuICAgIGUuc3RvcFByb3BhZ2F0aW9uKCk7XG4gICAgdGhpcy4jaGlkZVBvcG92ZXIoKTtcbiAgfTtcblxuICAjb25IZWFkZXJTbG90Q2hhbmdlKGU6IEV2ZW50KSB7XG4gICAgY29uc3Qgc2xvdCA9IGUudGFyZ2V0IGFzIEhUTUxTbG90RWxlbWVudDtcbiAgICB0aGlzLiNoYXNIZWFkZXJDb250ZW50ID0gc2xvdC5hc3NpZ25lZE5vZGVzKCkubGVuZ3RoID4gMDtcbiAgICB0aGlzLnJlcXVlc3RVcGRhdGUoKTtcbiAgfVxuXG4gICNvbkZvb3RlclNsb3RDaGFuZ2UoZTogRXZlbnQpIHtcbiAgICBjb25zdCBzbG90ID0gZS50YXJnZXQgYXMgSFRNTFNsb3RFbGVtZW50O1xuICAgIHRoaXMuI2hhc0Zvb3RlckNvbnRlbnQgPSBzbG90LmFzc2lnbmVkTm9kZXMoKS5sZW5ndGggPiAwO1xuICAgIHRoaXMucmVxdWVzdFVwZGF0ZSgpO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICBjbGVhclRpbWVvdXQodGhpcy4jaG92ZXJTaG93VGltZW91dCk7XG4gICAgY2xlYXJUaW1lb3V0KHRoaXMuI2hvdmVySGlkZVRpbWVvdXQpO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1wZi12Ni1wb3BvdmVyJzogUGZWNlBvcG92ZXI7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIGFsaWduLXNlbGY6IHN0YXJ0O1xcXFxuICBqdXN0aWZ5LXNlbGY6IHN0YXJ0O1xcXFxuXFxcXG4gIC0tZGlzdGFuY2U6IDhweDtcXFxcbiAgLS1taW4td2lkdGg6IDEwMHB4O1xcXFxuICAtLW1heC13aWR0aDogMzAwcHg7XFxcXG5cXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyLS1Gb250U2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlci0tQm94U2hhZG93OiB2YXIoLS1wZi10LS1nbG9iYWwtLWJveC1zaGFkb3ctLW1kKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1oaWdoLWNvbnRyYXN0LS1yZWd1bGFyKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1oaWdoLWNvbnRyYXN0KTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJSYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLW1lZGl1bSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1mbG9hdGluZy0tZGVmYXVsdCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0Jsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19jb250ZW50LS1QYWRkaW5nQmxvY2tFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Nsb3NlLS1JbnNldEJsb2NrU3RhcnQ6IGNhbGModmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0Jsb2NrU3RhcnQpIC0gdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLXZlcnRpY2FsLS1kZWZhdWx0KSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY2xvc2UtLUluc2V0SW5saW5lRW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19jb250ZW50LS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19jbG9zZS0tc2libGluZy0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLTJ4bCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9faGVhZGVyLS1NYXJnaW5CbG9ja0VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX190aXRsZS10ZXh0LS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fdGl0bGUtdGV4dC0tRm9udFdlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS13ZWlnaHQtLWJvZHktLWJvbGQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX3RpdGxlLXRleHQtLUZvbnRTaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Zvb3Rlci0tTWFyZ2luQmxvY2tTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tV2lkdGg6IDEuNXJlbTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSGVpZ2h0OiAxLjVyZW07XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUJhY2tncm91bmRDb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tQm94U2hhZG93OiB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyLS1Cb3hTaGFkb3cpO1xcXFxufVxcXFxuXFxcXG4jdHJpZ2dlciB7XFxcXG4gIGFuY2hvci1uYW1lOiAtLXRyaWdnZXI7XFxcXG4gIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQge1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB0cmFuc3BhcmVudDtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgb3ZlcmZsb3c6IGluaXRpYWw7XFxcXG4gIHBvc2l0aW9uOiBmaXhlZDtcXFxcbiAgcG9zaXRpb24tYW5jaG9yOiAtLXRyaWdnZXI7XFxcXG4gIG1pbi13aWR0aDogdmFyKC0tbWluLXdpZHRoKTtcXFxcbiAgbWF4LXdpZHRoOiB2YXIoLS1tYXgtd2lkdGgpO1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlci0tRm9udFNpemUpO1xcXFxuICBib3JkZXI6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXItLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJDb2xvcik7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXItLUJvcmRlclJhZGl1cyk7XFxcXG4gIGJveC1zaGFkb3c6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXItLUJveFNoYWRvdyk7XFxcXG4gIHRvcDogYXV0bztcXFxcbiAgYm90dG9tOiBhdXRvO1xcXFxuICBsZWZ0OiBhdXRvO1xcXFxuICByaWdodDogYXV0bztcXFxcbn1cXFxcblxcXFxuI3BvcG92ZXIge1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIHBhZGRpbmctYmxvY2stc3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2NvbnRlbnQtLVBhZGRpbmdCbG9ja1N0YXJ0KTtcXFxcbiAgcGFkZGluZy1ibG9jay1lbmQ6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2NvbnRlbnQtLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19jb250ZW50LS1QYWRkaW5nSW5saW5lU3RhcnQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1lbmQ6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2NvbnRlbnQtLVBhZGRpbmdJbmxpbmVFbmQpO1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19jb250ZW50LS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG4jYXJyb3cge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIHRvcDogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0QmxvY2tTdGFydCwgYXV0byk7XFxcXG4gIHJpZ2h0OiB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVFbmQsIGF1dG8pO1xcXFxuICBib3R0b206IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldEJsb2NrRW5kLCBhdXRvKTtcXFxcbiAgbGVmdDogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lU3RhcnQsIGF1dG8pO1xcXFxuICB3aWR0aDogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVdpZHRoKTtcXFxcbiAgaGVpZ2h0OiB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSGVpZ2h0KTtcXFxcbiAgcG9pbnRlci1ldmVudHM6IG5vbmU7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICBib3JkZXI6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXItLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJDb2xvcik7XFxcXG4gIGJveC1zaGFkb3c6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1Cb3hTaGFkb3cpO1xcXFxuICB0cmFuc2Zvcm06IHRyYW5zbGF0ZVgodmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVRyYW5zbGF0ZVgsIDApKSB0cmFuc2xhdGVZKHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1UcmFuc2xhdGVZLCAwKSkgcm90YXRlKHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1Sb3RhdGUsIDApKTtcXFxcbn1cXFxcblxcXFxuI2Nsb3NlIHtcXFxcbiAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICBpbnNldC1ibG9jay1zdGFydDogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY2xvc2UtLUluc2V0QmxvY2tTdGFydCk7XFxcXG4gIGluc2V0LWlubGluZS1lbmQ6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Nsb3NlLS1JbnNldElubGluZUVuZCk7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLWhvcml6b250YWwtLXBsYWluLS1kZWZhdWx0KTtcXFxcbiAgYmFja2dyb3VuZDogdHJhbnNwYXJlbnQ7XFxcXG4gIGJvcmRlcjogbm9uZTtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1pY29uLS1jb2xvci0tcmVndWxhcik7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGp1c3RpZnktY29udGVudDogY2VudGVyO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0taG92ZXIpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbjpob3N0KDpub3QoW2Nsb3NlYWJsZV0pKSAjY2xvc2UsXFxcXG46aG9zdChbY2xvc2VhYmxlPVxcXFxcXFwiZmFsc2VcXFxcXFxcIl0pICNjbG9zZSB7XFxcXG4gIGRpc3BsYXk6IG5vbmU7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtjbG9zZWFibGVdKSAjaGVhZGVyLFxcXFxuOmhvc3QoW2Nsb3NlYWJsZT1cXFxcXFxcIlxcXFxcXFwiXSkgI2hlYWRlciB7XFxcXG4gIHBhZGRpbmctaW5saW5lLWVuZDogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fY2xvc2UtLXNpYmxpbmctLVBhZGRpbmdJbmxpbmVFbmQpO1xcXFxufVxcXFxuXFxcXG4jaGVhZGVyIHtcXFxcbiAgbWFyZ2luLWJsb2NrLWVuZDogMDtcXFxcblxcXFxuICBcXFxcdTAwMjYuaGFzLWNvbnRlbnQge1xcXFxuICAgIG1hcmdpbi1ibG9jay1lbmQ6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2hlYWRlci0tTWFyZ2luQmxvY2tFbmQpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNoZWFkZXItY29udGVudCB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXg6IDAgMCBhdXRvO1xcXFxuXFxcXG4gIDo6c2xvdHRlZChoMSksXFxcXG4gIDo6c2xvdHRlZChoMiksXFxcXG4gIDo6c2xvdHRlZChoMyksXFxcXG4gIDo6c2xvdHRlZChoNCksXFxcXG4gIDo6c2xvdHRlZChoNSksXFxcXG4gIDo6c2xvdHRlZChoNikge1xcXFxuICAgIG1pbi13aWR0aDogMCAhaW1wb3J0YW50O1xcXFxuICAgIGZvbnQtc2l6ZTogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fdGl0bGUtdGV4dC0tRm9udFNpemUpICFpbXBvcnRhbnQ7XFxcXG4gICAgZm9udC13ZWlnaHQ6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX3RpdGxlLXRleHQtLUZvbnRXZWlnaHQpICFpbXBvcnRhbnQ7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX3RpdGxlLXRleHQtLUNvbG9yKSAhaW1wb3J0YW50O1xcXFxuICAgIG92ZXJmbG93LXdyYXA6IGJyZWFrLXdvcmQgIWltcG9ydGFudDtcXFxcbiAgICBtYXJnaW46IDAgIWltcG9ydGFudDtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jYm9keSB7XFxcXG4gIHdvcmQtd3JhcDogYnJlYWstd29yZDtcXFxcbn1cXFxcblxcXFxuI2Zvb3RlciB7XFxcXG4gIG1hcmdpbi1ibG9jay1zdGFydDogMDtcXFxcblxcXFxuICBcXFxcdTAwMjYuaGFzLWNvbnRlbnQge1xcXFxuICAgIG1hcmdpbi1ibG9jay1zdGFydDogdmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fZm9vdGVyLS1NYXJnaW5CbG9ja1N0YXJ0KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG46aG9zdChbaGFzLWF1dG8td2lkdGhdKSAjY29udGVudCB7XFxcXG4gIG1pbi13aWR0aDogdW5zZXQ7XFxcXG4gIG1heC13aWR0aDogdW5zZXQ7XFxcXG4gIHdpZHRoOiBtYXgtY29udGVudDtcXFxcbn1cXFxcblxcXFxuLyogUG9zaXRpb246IHRvcCAoZGVmYXVsdCkgKi9cXFxcbiNjb250ZW50LnBvc2l0aW9uLXRvcCB7XFxcXG4gIGJvdHRvbTogYW5jaG9yKC0tdHJpZ2dlciB0b3ApO1xcXFxuICBsZWZ0OiBhbmNob3IoLS10cmlnZ2VyIGNlbnRlcik7XFxcXG4gIHRyYW5zbGF0ZTogLTUwJSAwO1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0QmxvY2tFbmQ6IGNhbGModmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUhlaWdodCkgLyAtMik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lU3RhcnQ6IDUwJTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tVHJhbnNsYXRlWDogLTUwJTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tYm90dG9tIHtcXFxcbiAgdG9wOiBhbmNob3IoLS10cmlnZ2VyIGJvdHRvbSk7XFxcXG4gIGxlZnQ6IGFuY2hvcigtLXRyaWdnZXIgY2VudGVyKTtcXFxcbiAgdHJhbnNsYXRlOiAtNTAlIDA7XFxcXG4gIG1hcmdpbi10b3A6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiBjYWxjKHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1IZWlnaHQpIC8gLTIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldElubGluZVN0YXJ0OiA1MCU7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVRyYW5zbGF0ZVg6IC01MCU7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLWxlZnQge1xcXFxuICB0b3A6IGFuY2hvcigtLXRyaWdnZXIgY2VudGVyKTtcXFxcbiAgcmlnaHQ6IGFuY2hvcigtLXRyaWdnZXIgbGVmdCk7XFxcXG4gIHRyYW5zbGF0ZTogMCAtNTAlO1xcXFxuICBtYXJnaW4tcmlnaHQ6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiA1MCU7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lRW5kOiBjYWxjKHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1XaWR0aCkgLyAtMik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVRyYW5zbGF0ZVk6IC01MCU7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLXJpZ2h0IHtcXFxcbiAgdG9wOiBhbmNob3IoLS10cmlnZ2VyIGNlbnRlcik7XFxcXG4gIGxlZnQ6IGFuY2hvcigtLXRyaWdnZXIgcmlnaHQpO1xcXFxuICB0cmFuc2xhdGU6IDAgLTUwJTtcXFxcbiAgbWFyZ2luLWxlZnQ6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiA1MCU7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lU3RhcnQ6IGNhbGModmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVdpZHRoKSAvIC0yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tVHJhbnNsYXRlWTogLTUwJTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tdG9wLXN0YXJ0IHtcXFxcbiAgYm90dG9tOiBhbmNob3IoLS10cmlnZ2VyIHRvcCk7XFxcXG4gIGxlZnQ6IGFuY2hvcigtLXRyaWdnZXIgbGVmdCk7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja0VuZDogY2FsYyh2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSGVpZ2h0KSAvIC0yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tdG9wLWVuZCB7XFxcXG4gIGJvdHRvbTogYW5jaG9yKC0tdHJpZ2dlciB0b3ApO1xcXFxuICByaWdodDogYW5jaG9yKC0tdHJpZ2dlciByaWdodCk7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja0VuZDogY2FsYyh2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSGVpZ2h0KSAvIC0yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLWJvdHRvbS1zdGFydCB7XFxcXG4gIHRvcDogYW5jaG9yKC0tdHJpZ2dlciBib3R0b20pO1xcXFxuICBsZWZ0OiBhbmNob3IoLS10cmlnZ2VyIGxlZnQpO1xcXFxuICBtYXJnaW4tdG9wOiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0QmxvY2tTdGFydDogY2FsYyh2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSGVpZ2h0KSAvIC0yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tYm90dG9tLWVuZCB7XFxcXG4gIHRvcDogYW5jaG9yKC0tdHJpZ2dlciBib3R0b20pO1xcXFxuICByaWdodDogYW5jaG9yKC0tdHJpZ2dlciByaWdodCk7XFxcXG4gIG1hcmdpbi10b3A6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiBjYWxjKHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1IZWlnaHQpIC8gLTIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldElubGluZUVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tbGVmdC1zdGFydCB7XFxcXG4gIHRvcDogYW5jaG9yKC0tdHJpZ2dlciB0b3ApO1xcXFxuICByaWdodDogYW5jaG9yKC0tdHJpZ2dlciBsZWZ0KTtcXFxcbiAgbWFyZ2luLXJpZ2h0OiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0QmxvY2tTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVFbmQ6IGNhbGModmFyKC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVdpZHRoKSAvIC0yKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tbGVmdC1lbmQge1xcXFxuICBib3R0b206IGFuY2hvcigtLXRyaWdnZXIgYm90dG9tKTtcXFxcbiAgcmlnaHQ6IGFuY2hvcigtLXRyaWdnZXIgbGVmdCk7XFxcXG4gIG1hcmdpbi1yaWdodDogdmFyKC0tZGlzdGFuY2UpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldEJsb2NrRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldElubGluZUVuZDogY2FsYyh2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tV2lkdGgpIC8gLTIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1Sb3RhdGU6IDQ1ZGVnO1xcXFxufVxcXFxuXFxcXG4jY29udGVudC5wb3NpdGlvbi1yaWdodC1zdGFydCB7XFxcXG4gIHRvcDogYW5jaG9yKC0tdHJpZ2dlciB0b3ApO1xcXFxuICBsZWZ0OiBhbmNob3IoLS10cmlnZ2VyIHJpZ2h0KTtcXFxcbiAgbWFyZ2luLWxlZnQ6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldElubGluZVN0YXJ0OiBjYWxjKHZhcigtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1XaWR0aCkgLyAtMik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLXJpZ2h0LWVuZCB7XFxcXG4gIGJvdHRvbTogYW5jaG9yKC0tdHJpZ2dlciBib3R0b20pO1xcXFxuICBsZWZ0OiBhbmNob3IoLS10cmlnZ2VyIHJpZ2h0KTtcXFxcbiAgbWFyZ2luLWxlZnQ6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja0VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVTdGFydDogY2FsYyh2YXIoLS1jZW0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tV2lkdGgpIC8gLTIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1Sb3RhdGU6IDQ1ZGVnO1xcXFxufVxcXFxuXFxcXG4jY29udGVudC5wb3NpdGlvbi1hdXRvIHtcXFxcbiAgYm90dG9tOiBhbmNob3IoLS10cmlnZ2VyIHRvcCk7XFxcXG4gIGxlZnQ6IGFuY2hvcigtLXRyaWdnZXIgY2VudGVyKTtcXFxcbiAgdHJhbnNsYXRlOiAtNTAlIDA7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWRpc3RhbmNlKTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLE1BQU0sZUFBZTtBQUMxQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLGt5VkFBd3lWLENBQUM7QUFBRSxJQUFPLDRCQUFROzs7QURnQngyVixJQUFNLHFCQUFOLGNBQWlDLE1BQU07QUFBQSxFQUM1QyxPQUFPO0FBQUEsRUFDUCxjQUFjO0FBQ1osVUFBTSxtQkFBbUIsRUFBRSxTQUFTLEtBQUssQ0FBQztBQUFBLEVBQzVDO0FBQ0Y7QUFLTyxJQUFNLHFCQUFOLGNBQWlDLE1BQU07QUFBQSxFQUM1QyxPQUFPO0FBQUEsRUFDUCxjQUFjO0FBQ1osVUFBTSxtQkFBbUIsRUFBRSxTQUFTLEtBQUssQ0FBQztBQUFBLEVBQzVDO0FBQ0Y7QUFHQSxJQUFNLDRCQUE0QixXQUFXLEtBQUssV0FBVyxxQkFBcUIsS0FBSztBQWxDdkY7QUFpREEsMkJBQUMsY0FBYyxtQkFBbUI7QUFDM0IsSUFBTSxjQUFOLGVBQTBCLGlCQUcvQixhQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFHMUMsaUJBQUMsU0FBUyxFQUFFLFNBQVMsS0FBSyxDQUFDLElBRzNCLGlCQUFDLFNBQVMsRUFBRSxNQUFNLFFBQVEsU0FBUyxLQUFLLENBQUMsSUFHekMsc0JBQUMsU0FBUyxFQUFFLFNBQVMsTUFBTSxXQUFXLGlCQUFpQixDQUFDLElBR3hELGtCQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFHMUMseUJBQUMsU0FBUyxFQUFFLFdBQVcscUJBQXFCLENBQUMsSUFHN0MsaUJBQUMs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
