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

// elements/pf-v6-popover/pf-v6-popover.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-popover/pf-v6-popover.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: inline-block;\\n  position: relative;\\n  align-self: start;\\n  justify-self: start;\\n\\n  --distance: 8px;\\n  --min-width: 100px;\\n  --max-width: 300px;\\n\\n  --pf-v6-c-popover--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --pf-v6-c-popover--BoxShadow: var(--pf-t--global--box-shadow--md);\\n  --pf-v6-c-popover--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --pf-v6-c-popover--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --pf-v6-c-popover--BorderRadius: var(--pf-t--global--border--radius--medium);\\n  --pf-v6-c-popover__content--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --pf-v6-c-popover__content--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --pf-v6-c-popover__content--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --pf-v6-c-popover__content--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --pf-v6-c-popover__content--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --pf-v6-c-popover__close--InsetBlockStart: calc(var(--pf-v6-c-popover__content--PaddingBlockStart) - var(--pf-t--global--spacer--control--vertical--default));\\n  --pf-v6-c-popover__close--InsetInlineEnd: var(--pf-v6-c-popover__content--PaddingInlineEnd);\\n  --pf-v6-c-popover__close--sibling--PaddingInlineEnd: var(--pf-t--global--spacer--2xl);\\n  --pf-v6-c-popover__header--MarginBlockEnd: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-popover__title-text--Color: var(--pf-t--global--text--color--regular);\\n  --pf-v6-c-popover__title-text--FontWeight: var(--pf-t--global--font--weight--body--bold);\\n  --pf-v6-c-popover__title-text--FontSize: var(--pf-t--global--font--size--body--default);\\n  --pf-v6-c-popover__footer--MarginBlockStart: var(--pf-t--global--spacer--md);\\n  --pf-v6-c-popover__arrow--Width: 1.5rem;\\n  --pf-v6-c-popover__arrow--Height: 1.5rem;\\n  --pf-v6-c-popover__arrow--BackgroundColor: var(--pf-v6-c-popover__content--BackgroundColor);\\n  --pf-v6-c-popover__arrow--BoxShadow: var(--pf-v6-c-popover--BoxShadow);\\n}\\n\\n#trigger {\\n  anchor-name: --trigger;\\n  display: inline-block;\\n}\\n\\n#content {\\n  background-color: transparent;\\n  padding: 0;\\n  overflow: initial;\\n  position: fixed;\\n  position-anchor: --trigger;\\n  min-width: var(--min-width);\\n  max-width: var(--max-width);\\n  margin: 0;\\n  font-size: var(--pf-v6-c-popover--FontSize);\\n  border: var(--pf-v6-c-popover--BorderWidth) solid var(--pf-v6-c-popover--BorderColor);\\n  border-radius: var(--pf-v6-c-popover--BorderRadius);\\n  box-shadow: var(--pf-v6-c-popover--BoxShadow);\\n  top: auto;\\n  bottom: auto;\\n  left: auto;\\n  right: auto;\\n}\\n\\n#popover {\\n  position: relative;\\n  padding-block-start: var(--pf-v6-c-popover__content--PaddingBlockStart);\\n  padding-block-end: var(--pf-v6-c-popover__content--PaddingBlockEnd);\\n  padding-inline-start: var(--pf-v6-c-popover__content--PaddingInlineStart);\\n  padding-inline-end: var(--pf-v6-c-popover__content--PaddingInlineEnd);\\n  background-color: var(--pf-v6-c-popover__content--BackgroundColor);\\n  border-radius: var(--pf-v6-c-popover--BorderRadius);\\n}\\n\\n#arrow {\\n  position: absolute;\\n  top: var(--pf-v6-c-popover__arrow--InsetBlockStart, auto);\\n  right: var(--pf-v6-c-popover__arrow--InsetInlineEnd, auto);\\n  bottom: var(--pf-v6-c-popover__arrow--InsetBlockEnd, auto);\\n  left: var(--pf-v6-c-popover__arrow--InsetInlineStart, auto);\\n  width: var(--pf-v6-c-popover__arrow--Width);\\n  height: var(--pf-v6-c-popover__arrow--Height);\\n  pointer-events: none;\\n  background-color: var(--pf-v6-c-popover__arrow--BackgroundColor);\\n  border: var(--pf-v6-c-popover--BorderWidth) solid var(--pf-v6-c-popover--BorderColor);\\n  box-shadow: var(--pf-v6-c-popover__arrow--BoxShadow);\\n  transform: translateX(var(--pf-v6-c-popover__arrow--TranslateX, 0)) translateY(var(--pf-v6-c-popover__arrow--TranslateY, 0)) rotate(var(--pf-v6-c-popover__arrow--Rotate, 0));\\n}\\n\\n#close {\\n  position: absolute;\\n  inset-block-start: var(--pf-v6-c-popover__close--InsetBlockStart);\\n  inset-inline-end: var(--pf-v6-c-popover__close--InsetInlineEnd);\\n  padding: var(--pf-t--global--spacer--control--vertical--default) var(--pf-t--global--spacer--control--horizontal--plain--default);\\n  background: transparent;\\n  border: none;\\n  cursor: pointer;\\n  color: var(--pf-t--global--icon--color--regular);\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n\\n  \\u0026:hover {\\n    background-color: var(--pf-t--global--background--color--action--plain--hover);\\n  }\\n}\\n\\n:host(:not([closeable])) #close,\\n:host([closeable=\\"false\\"]) #close {\\n  display: none;\\n}\\n\\n:host([closeable]) #header,\\n:host([closeable=\\"\\"]) #header {\\n  padding-inline-end: var(--pf-v6-c-popover__close--sibling--PaddingInlineEnd);\\n}\\n\\n#header {\\n  margin-block-end: 0;\\n\\n  \\u0026.has-content {\\n    margin-block-end: var(--pf-v6-c-popover__header--MarginBlockEnd);\\n  }\\n}\\n\\n#header-content {\\n  display: flex;\\n  flex: 0 0 auto;\\n\\n  ::slotted(h1),\\n  ::slotted(h2),\\n  ::slotted(h3),\\n  ::slotted(h4),\\n  ::slotted(h5),\\n  ::slotted(h6) {\\n    min-width: 0 !important;\\n    font-size: var(--pf-v6-c-popover__title-text--FontSize) !important;\\n    font-weight: var(--pf-v6-c-popover__title-text--FontWeight) !important;\\n    color: var(--pf-v6-c-popover__title-text--Color) !important;\\n    overflow-wrap: break-word !important;\\n    margin: 0 !important;\\n  }\\n}\\n\\n#body {\\n  word-wrap: break-word;\\n}\\n\\n#footer {\\n  margin-block-start: 0;\\n\\n  \\u0026.has-content {\\n    margin-block-start: var(--pf-v6-c-popover__footer--MarginBlockStart);\\n  }\\n}\\n\\n:host([has-auto-width]) #content {\\n  min-width: unset;\\n  max-width: unset;\\n  width: max-content;\\n}\\n\\n/* Position: top (default) */\\n#content.position-top {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-bottom: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--pf-v6-c-popover__arrow--Height) / -2);\\n  --pf-v6-c-popover__arrow--InsetInlineStart: 50%;\\n  --pf-v6-c-popover__arrow--TranslateX: -50%;\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom {\\n  top: anchor(--trigger bottom);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-top: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--pf-v6-c-popover__arrow--Height) / -2);\\n  --pf-v6-c-popover__arrow--InsetInlineStart: 50%;\\n  --pf-v6-c-popover__arrow--TranslateX: -50%;\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left {\\n  top: anchor(--trigger center);\\n  right: anchor(--trigger left);\\n  translate: 0 -50%;\\n  margin-right: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockStart: 50%;\\n  --pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--pf-v6-c-popover__arrow--Width) / -2);\\n  --pf-v6-c-popover__arrow--TranslateY: -50%;\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right {\\n  top: anchor(--trigger center);\\n  left: anchor(--trigger right);\\n  translate: 0 -50%;\\n  margin-left: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockStart: 50%;\\n  --pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--pf-v6-c-popover__arrow--Width) / -2);\\n  --pf-v6-c-popover__arrow--TranslateY: -50%;\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-top-start {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger left);\\n  margin-bottom: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--pf-v6-c-popover__arrow--Height) / -2);\\n  --pf-v6-c-popover__arrow--InsetInlineStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-top-end {\\n  bottom: anchor(--trigger top);\\n  right: anchor(--trigger right);\\n  margin-bottom: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--pf-v6-c-popover__arrow--Height) / -2);\\n  --pf-v6-c-popover__arrow--InsetInlineEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom-start {\\n  top: anchor(--trigger bottom);\\n  left: anchor(--trigger left);\\n  margin-top: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--pf-v6-c-popover__arrow--Height) / -2);\\n  --pf-v6-c-popover__arrow--InsetInlineStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom-end {\\n  top: anchor(--trigger bottom);\\n  right: anchor(--trigger right);\\n  margin-top: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--pf-v6-c-popover__arrow--Height) / -2);\\n  --pf-v6-c-popover__arrow--InsetInlineEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left-start {\\n  top: anchor(--trigger top);\\n  right: anchor(--trigger left);\\n  margin-right: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--pf-v6-c-popover__arrow--Width) / -2);\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left-end {\\n  bottom: anchor(--trigger bottom);\\n  right: anchor(--trigger left);\\n  margin-right: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--pf-v6-c-popover__arrow--Width) / -2);\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right-start {\\n  top: anchor(--trigger top);\\n  left: anchor(--trigger right);\\n  margin-left: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--pf-v6-c-popover__arrow--Width) / -2);\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right-end {\\n  bottom: anchor(--trigger bottom);\\n  left: anchor(--trigger right);\\n  margin-left: var(--distance);\\n  --pf-v6-c-popover__arrow--InsetBlockEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--pf-v6-c-popover__arrow--Width) / -2);\\n  --pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-auto {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-bottom: var(--distance);\\n}\\n"'));
var pf_v6_popover_default = s;

// elements/pf-v6-popover/pf-v6-popover.ts
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
_PfV6Popover_decorators = [customElement("pf-v6-popover")];
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
            <svg class="pf-v6-svg"
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
__publicField(PfV6Popover, "styles", pf_v6_popover_default);
__runInitializers(_init, 1, PfV6Popover);
export {
  PfPopoverHideEvent,
  PfPopoverShowEvent,
  PfV6Popover
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtcG9wb3Zlci9wZi12Ni1wb3BvdmVyLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvcGYtdjYtcG9wb3Zlci9wZi12Ni1wb3BvdmVyLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9wZi12Ni1wb3BvdmVyLmNzcyc7XG5cbnR5cGUgUG9wb3ZlclBvc2l0aW9uID1cbiAgfCAnYXV0bycgfCAndG9wJyB8ICdib3R0b20nIHwgJ2xlZnQnIHwgJ3JpZ2h0J1xuICB8ICd0b3Atc3RhcnQnIHwgJ3RvcC1lbmQnIHwgJ2JvdHRvbS1zdGFydCcgfCAnYm90dG9tLWVuZCdcbiAgfCAnbGVmdC1zdGFydCcgfCAnbGVmdC1lbmQnIHwgJ3JpZ2h0LXN0YXJ0JyB8ICdyaWdodC1lbmQnO1xuXG50eXBlIFRyaWdnZXJBY3Rpb24gPSAnY2xpY2snIHwgJ2hvdmVyJztcblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBwb3BvdmVyIHNob3dzXG4gKi9cbmV4cG9ydCBjbGFzcyBQZlBvcG92ZXJTaG93RXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIG9wZW4gPSB0cnVlO1xuICBjb25zdHJ1Y3RvcigpIHtcbiAgICBzdXBlcigncGYtcG9wb3Zlci1zaG93JywgeyBidWJibGVzOiB0cnVlIH0pO1xuICB9XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gcG9wb3ZlciBoaWRlc1xuICovXG5leHBvcnQgY2xhc3MgUGZQb3BvdmVySGlkZUV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBvcGVuID0gZmFsc2U7XG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCdwZi1wb3BvdmVyLWhpZGUnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gIH1cbn1cblxuLy8gQ2FjaGUgZmVhdHVyZSBkZXRlY3Rpb24gZm9yIENTUyBBbmNob3IgUG9zaXRpb25pbmdcbmNvbnN0IHN1cHBvcnRzQW5jaG9yUG9zaXRpb25pbmcgPSBnbG9iYWxUaGlzLkNTUz8uc3VwcG9ydHM/LignYW5jaG9yLW5hbWU6IC0tdGVzdCcpID8/IGZhbHNlO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgUG9wb3ZlclxuICpcbiAqIEEgcG9wb3ZlciBjb21wb25lbnQgdXNpbmcgdGhlIG5hdGl2ZSBQb3BvdmVyIEFQSSBhbmQgQ1NTIEFuY2hvciBQb3NpdGlvbmluZy5cbiAqXG4gKiBAc2xvdCB0cmlnZ2VyIC0gVGhlIGVsZW1lbnQgdGhhdCB0cmlnZ2VycyB0aGUgcG9wb3ZlciAodXN1YWxseSBhIGJ1dHRvbilcbiAqIEBzbG90IGhlYWRlciAtIEhlYWRlciBjb250ZW50XG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgcG9wb3ZlciBib2R5IGNvbnRlbnRcbiAqIEBzbG90IGZvb3RlciAtIEZvb3RlciBjb250ZW50XG4gKlxuICogQGZpcmVzIHBmLXBvcG92ZXItc2hvdyAtIEZpcmVkIHdoZW4gcG9wb3ZlciBvcGVuc1xuICogQGZpcmVzIHBmLXBvcG92ZXItaGlkZSAtIEZpcmVkIHdoZW4gcG9wb3ZlciBjbG9zZXNcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ3BmLXY2LXBvcG92ZXInKVxuZXhwb3J0IGNsYXNzIFBmVjZQb3BvdmVyIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBvcGVuID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBwb3NpdGlvbjogUG9wb3ZlclBvc2l0aW9uID0gJ3RvcCc7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogTnVtYmVyLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGRpc3RhbmNlID0gODtcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlLCBhdHRyaWJ1dGU6ICd0cmlnZ2VyLWFjdGlvbicgfSlcbiAgYWNjZXNzb3IgdHJpZ2dlckFjdGlvbjogVHJpZ2dlckFjdGlvbiA9ICdjbGljayc7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBjbG9zZWFibGUgPSB0cnVlO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ2Nsb3NlLWJ1dHRvbi1sYWJlbCcgfSlcbiAgYWNjZXNzb3IgY2xvc2VCdXR0b25MYWJlbCA9ICdDbG9zZSc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnbWluLXdpZHRoJyB9KVxuICBhY2Nlc3NvciBtaW5XaWR0aD86IHN0cmluZztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdtYXgtd2lkdGgnIH0pXG4gIGFjY2Vzc29yIG1heFdpZHRoPzogc3RyaW5nO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUsIGF0dHJpYnV0ZTogJ2hhcy1hdXRvLXdpZHRoJyB9KVxuICBhY2Nlc3NvciBoYXNBdXRvV2lkdGggPSBmYWxzZTtcblxuICAjaG92ZXJTaG93VGltZW91dD86IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+O1xuICAjaG92ZXJIaWRlVGltZW91dD86IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+O1xuICAjaGFzSGVhZGVyQ29udGVudCA9IGZhbHNlO1xuICAjaGFzRm9vdGVyQ29udGVudCA9IGZhbHNlO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJ0cmlnZ2VyXCJcbiAgICAgICAgICAgQGNsaWNrPSR7dGhpcy50cmlnZ2VyQWN0aW9uID09PSAnY2xpY2snID8gdGhpcy4jaGFuZGxlVHJpZ2dlckNsaWNrIDogbm90aGluZ31cbiAgICAgICAgICAgQHBvaW50ZXJlbnRlcj0ke3RoaXMudHJpZ2dlckFjdGlvbiA9PT0gJ2hvdmVyJyA/IHRoaXMuI2hhbmRsZVBvaW50ZXJFbnRlciA6IG5vdGhpbmd9XG4gICAgICAgICAgIEBwb2ludGVybGVhdmU9JHt0aGlzLnRyaWdnZXJBY3Rpb24gPT09ICdob3ZlcicgPyB0aGlzLiNoYW5kbGVQb2ludGVyTGVhdmUgOiBub3RoaW5nfT5cbiAgICAgICAgPHNsb3QgbmFtZT1cInRyaWdnZXJcIj48L3Nsb3Q+XG4gICAgICA8L2Rpdj5cblxuICAgICAgPGRpdiBpZD1cImNvbnRlbnRcIlxuICAgICAgICAgICBjbGFzcz1cInBvc2l0aW9uLSR7dGhpcy5wb3NpdGlvbn1cIlxuICAgICAgICAgICBwb3BvdmVyPVwibWFudWFsXCJcbiAgICAgICAgICAgcGFydD1cInBvcG92ZXJcIlxuICAgICAgICAgICBzdHlsZT0ke3RoaXMuI2NvbnRlbnRTdHlsZX1cbiAgICAgICAgICAgQHBvaW50ZXJlbnRlcj0ke3RoaXMudHJpZ2dlckFjdGlvbiA9PT0gJ2hvdmVyJyA/IHRoaXMuI2hhbmRsZVBvcG92ZXJQb2ludGVyRW50ZXIgOiBub3RoaW5nfVxuICAgICAgICAgICBAcG9pbnRlcmxlYXZlPSR7dGhpcy50cmlnZ2VyQWN0aW9uID09PSAnaG92ZXInID8gdGhpcy4jaGFuZGxlUG9wb3ZlclBvaW50ZXJMZWF2ZSA6IG5vdGhpbmd9XG4gICAgICAgICAgIEB0b2dnbGU9JHt0aGlzLiNoYW5kbGVQb3BvdmVyVG9nZ2xlfT5cbiAgICAgICAgPGRpdiBpZD1cImFycm93XCI+PC9kaXY+XG4gICAgICAgIDxkaXYgaWQ9XCJwb3BvdmVyXCJcbiAgICAgICAgICAgICBwYXJ0PVwiY29udGVudFwiPlxuICAgICAgICAgIDxidXR0b24gaWQ9XCJjbG9zZVwiXG4gICAgICAgICAgICAgICAgICBwYXJ0PVwiY2xvc2UtYnV0dG9uXCJcbiAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9JHt0aGlzLmNsb3NlQnV0dG9uTGFiZWx9XG4gICAgICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLiNoYW5kbGVDbG9zZX0+XG4gICAgICAgICAgICA8c3ZnIGNsYXNzPVwicGYtdjYtc3ZnXCJcbiAgICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCAzNTIgNTEyXCJcbiAgICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIlxuICAgICAgICAgICAgICAgICB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiPlxuICAgICAgICAgICAgICA8cGF0aCBkPVwiTTI0Mi43MiAyNTZsMTAwLjA3LTEwMC4wN2MxMi4yOC0xMi4yOCAxMi4yOC0zMi4xOSAwLTQ0LjQ4bC0yMi4yNC0yMi4yNGMtMTIuMjgtMTIuMjgtMzIuMTktMTIuMjgtNDQuNDggMEwxNzYgMTg5LjI4IDc1LjkzIDg5LjIxYy0xMi4yOC0xMi4yOC0zMi4xOS0xMi4yOC00NC40OCAwTDkuMjEgMTExLjQ1Yy0xMi4yOCAxMi4yOC0xMi4yOCAzMi4xOSAwIDQ0LjQ4TDEwOS4yOCAyNTYgOS4yMSAzNTYuMDdjLTEyLjI4IDEyLjI4LTEyLjI4IDMyLjE5IDAgNDQuNDhsMjIuMjQgMjIuMjRjMTIuMjggMTIuMjggMzIuMiAxMi4yOCA0NC40OCAwTDE3NiAzMjIuNzJsMTAwLjA3IDEwMC4wN2MxMi4yOCAxMi4yOCAzMi4yIDEyLjI4IDQ0LjQ4IDBsMjIuMjQtMjIuMjRjMTIuMjgtMTIuMjggMTIuMjgtMzIuMTkgMC00NC40OEwyNDIuNzIgMjU2elwiPjwvcGF0aD5cbiAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgIDwvYnV0dG9uPlxuXG4gICAgICAgICAgPGhlYWRlciBpZD1cImhlYWRlclwiXG4gICAgICAgICAgICAgICAgICBjbGFzcz0ke3RoaXMuI2hhc0hlYWRlckNvbnRlbnQgPyAnaGFzLWNvbnRlbnQnIDogJyd9XG4gICAgICAgICAgICAgICAgICBwYXJ0PVwiaGVhZGVyXCI+XG4gICAgICAgICAgICA8ZGl2IGlkPVwidGl0bGVcIj5cbiAgICAgICAgICAgICAgPGRpdiBpZD1cImhlYWRlci1jb250ZW50XCI+XG4gICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cImhlYWRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgQHNsb3RjaGFuZ2U9JHt0aGlzLiNvbkhlYWRlclNsb3RDaGFuZ2V9Pjwvc2xvdD5cbiAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8L2hlYWRlcj5cblxuICAgICAgICAgIDxkaXYgaWQ9XCJib2R5XCJcbiAgICAgICAgICAgICAgIHBhcnQ9XCJib2R5XCI+XG4gICAgICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICAgICAgPC9kaXY+XG5cbiAgICAgICAgICA8Zm9vdGVyIGlkPVwiZm9vdGVyXCJcbiAgICAgICAgICAgICAgICAgIGNsYXNzPSR7dGhpcy4jaGFzRm9vdGVyQ29udGVudCA/ICdoYXMtY29udGVudCcgOiAnJ31cbiAgICAgICAgICAgICAgICAgIHBhcnQ9XCJmb290ZXJcIj5cbiAgICAgICAgICAgIDxzbG90IG5hbWU9XCJmb290ZXJcIlxuICAgICAgICAgICAgICAgICAgQHNsb3RjaGFuZ2U9JHt0aGlzLiNvbkZvb3RlclNsb3RDaGFuZ2V9Pjwvc2xvdD5cbiAgICAgICAgICA8L2Zvb3Rlcj5cbiAgICAgICAgPC9kaXY+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG5cbiAgZ2V0ICNjb250ZW50U3R5bGUoKTogc3RyaW5nIHtcbiAgICBjb25zdCBwYXJ0czogc3RyaW5nW10gPSBbXTtcbiAgICBpZiAodGhpcy5taW5XaWR0aCkgcGFydHMucHVzaChgLS1taW4td2lkdGg6ICR7dGhpcy5taW5XaWR0aH1gKTtcbiAgICBpZiAodGhpcy5tYXhXaWR0aCkgcGFydHMucHVzaChgLS1tYXgtd2lkdGg6ICR7dGhpcy5tYXhXaWR0aH1gKTtcbiAgICBwYXJ0cy5wdXNoKGAtLWRpc3RhbmNlOiAke3RoaXMuZGlzdGFuY2V9cHhgKTtcbiAgICByZXR1cm4gcGFydHMuam9pbignOyAnKTtcbiAgfVxuXG4gIHVwZGF0ZWQoY2hhbmdlZDogTWFwPHN0cmluZywgdW5rbm93bj4pIHtcbiAgICBpZiAoY2hhbmdlZC5oYXMoJ29wZW4nKSkge1xuICAgICAgY29uc3QgY29udGVudEVsID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnY29udGVudCcpIGFzIEhUTUxFbGVtZW50ICYgeyBzaG93UG9wb3ZlcigpOiB2b2lkOyBoaWRlUG9wb3ZlcigpOiB2b2lkOyBtYXRjaGVzKHM6IHN0cmluZyk6IGJvb2xlYW4gfSB8IG51bGw7XG4gICAgICBpZiAoIWNvbnRlbnRFbCkgcmV0dXJuO1xuICAgICAgaWYgKHRoaXMub3BlbiAmJiAhY29udGVudEVsLm1hdGNoZXMoJzpwb3BvdmVyLW9wZW4nKSkge1xuICAgICAgICB0aGlzLiNzaG93UG9wb3ZlcigpO1xuICAgICAgfSBlbHNlIGlmICghdGhpcy5vcGVuICYmIGNvbnRlbnRFbC5tYXRjaGVzKCc6cG9wb3Zlci1vcGVuJykpIHtcbiAgICAgICAgdGhpcy4jaGlkZVBvcG92ZXIoKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICB0b2dnbGUoKSB7XG4gICAgaWYgKHRoaXMub3Blbikge1xuICAgICAgdGhpcy4jaGlkZVBvcG92ZXIoKTtcbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy4jc2hvd1BvcG92ZXIoKTtcbiAgICB9XG4gIH1cblxuICBhc3luYyAjc2hvd1BvcG92ZXIoKSB7XG4gICAgY29uc3QgY29udGVudEVsID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnY29udGVudCcpIGFzIEhUTUxFbGVtZW50ICYgeyBzaG93UG9wb3ZlcigpOiB2b2lkIH0gfCBudWxsO1xuICAgIGlmICghY29udGVudEVsKSByZXR1cm47XG4gICAgdHJ5IHtcbiAgICAgIGNvbnRlbnRFbC5zaG93UG9wb3ZlcigpO1xuXG4gICAgICBpZiAoIXN1cHBvcnRzQW5jaG9yUG9zaXRpb25pbmcpIHtcbiAgICAgICAgYXdhaXQgbmV3IFByb21pc2UocmVzb2x2ZSA9PiByZXF1ZXN0QW5pbWF0aW9uRnJhbWUocmVzb2x2ZSkpO1xuXG4gICAgICAgIGNvbnN0IHRyaWdnZXIgPSB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCd0cmlnZ2VyJyk7XG4gICAgICAgIGlmICghdHJpZ2dlcikgcmV0dXJuO1xuICAgICAgICBjb25zdCB0cmlnZ2VyUmVjdCA9IHRyaWdnZXIuZ2V0Qm91bmRpbmdDbGllbnRSZWN0KCk7XG4gICAgICAgIGNvbnN0IHBvcG92ZXJSZWN0ID0gY29udGVudEVsLmdldEJvdW5kaW5nQ2xpZW50UmVjdCgpO1xuICAgICAgICBjb25zdCBkaXN0ID0gdGhpcy5kaXN0YW5jZTtcblxuICAgICAgICBsZXQgdG9wOiBudW1iZXIsIGxlZnQ6IG51bWJlcjtcblxuICAgICAgICBzd2l0Y2ggKHRoaXMucG9zaXRpb24pIHtcbiAgICAgICAgICBjYXNlICd0b3AnOlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QudG9wIC0gcG9wb3ZlclJlY3QuaGVpZ2h0IC0gZGlzdDtcbiAgICAgICAgICAgIGxlZnQgPSB0cmlnZ2VyUmVjdC5sZWZ0ICsgKHRyaWdnZXJSZWN0LndpZHRoIC8gMikgLSAocG9wb3ZlclJlY3Qud2lkdGggLyAyKTtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgIGNhc2UgJ3RvcC1zdGFydCc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC50b3AgLSBwb3BvdmVyUmVjdC5oZWlnaHQgLSBkaXN0O1xuICAgICAgICAgICAgbGVmdCA9IHRyaWdnZXJSZWN0LmxlZnQ7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBjYXNlICd0b3AtZW5kJzpcbiAgICAgICAgICAgIHRvcCA9IHRyaWdnZXJSZWN0LnRvcCAtIHBvcG92ZXJSZWN0LmhlaWdodCAtIGRpc3Q7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QucmlnaHQgLSBwb3BvdmVyUmVjdC53aWR0aDtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgIGNhc2UgJ2JvdHRvbSc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC5ib3R0b20gKyBkaXN0O1xuICAgICAgICAgICAgbGVmdCA9IHRyaWdnZXJSZWN0LmxlZnQgKyAodHJpZ2dlclJlY3Qud2lkdGggLyAyKSAtIChwb3BvdmVyUmVjdC53aWR0aCAvIDIpO1xuICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgY2FzZSAnYm90dG9tLXN0YXJ0JzpcbiAgICAgICAgICAgIHRvcCA9IHRyaWdnZXJSZWN0LmJvdHRvbSArIGRpc3Q7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QubGVmdDtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgIGNhc2UgJ2JvdHRvbS1lbmQnOlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QuYm90dG9tICsgZGlzdDtcbiAgICAgICAgICAgIGxlZnQgPSB0cmlnZ2VyUmVjdC5yaWdodCAtIHBvcG92ZXJSZWN0LndpZHRoO1xuICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgY2FzZSAnbGVmdCc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC50b3AgKyAodHJpZ2dlclJlY3QuaGVpZ2h0IC8gMikgLSAocG9wb3ZlclJlY3QuaGVpZ2h0IC8gMik7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QubGVmdCAtIHBvcG92ZXJSZWN0LndpZHRoIC0gZGlzdDtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgIGNhc2UgJ2xlZnQtc3RhcnQnOlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QudG9wO1xuICAgICAgICAgICAgbGVmdCA9IHRyaWdnZXJSZWN0LmxlZnQgLSBwb3BvdmVyUmVjdC53aWR0aCAtIGRpc3Q7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBjYXNlICdsZWZ0LWVuZCc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC5ib3R0b20gLSBwb3BvdmVyUmVjdC5oZWlnaHQ7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QubGVmdCAtIHBvcG92ZXJSZWN0LndpZHRoIC0gZGlzdDtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgIGNhc2UgJ3JpZ2h0JzpcbiAgICAgICAgICAgIHRvcCA9IHRyaWdnZXJSZWN0LnRvcCArICh0cmlnZ2VyUmVjdC5oZWlnaHQgLyAyKSAtIChwb3BvdmVyUmVjdC5oZWlnaHQgLyAyKTtcbiAgICAgICAgICAgIGxlZnQgPSB0cmlnZ2VyUmVjdC5yaWdodCArIGRpc3Q7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBjYXNlICdyaWdodC1zdGFydCc6XG4gICAgICAgICAgICB0b3AgPSB0cmlnZ2VyUmVjdC50b3A7XG4gICAgICAgICAgICBsZWZ0ID0gdHJpZ2dlclJlY3QucmlnaHQgKyBkaXN0O1xuICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgY2FzZSAncmlnaHQtZW5kJzpcbiAgICAgICAgICAgIHRvcCA9IHRyaWdnZXJSZWN0LmJvdHRvbSAtIHBvcG92ZXJSZWN0LmhlaWdodDtcbiAgICAgICAgICAgIGxlZnQgPSB0cmlnZ2VyUmVjdC5yaWdodCArIGRpc3Q7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgICBkZWZhdWx0OlxuICAgICAgICAgICAgdG9wID0gdHJpZ2dlclJlY3QudG9wIC0gcG9wb3ZlclJlY3QuaGVpZ2h0IC0gZGlzdDtcbiAgICAgICAgICAgIGxlZnQgPSB0cmlnZ2VyUmVjdC5sZWZ0ICsgKHRyaWdnZXJSZWN0LndpZHRoIC8gMikgLSAocG9wb3ZlclJlY3Qud2lkdGggLyAyKTtcbiAgICAgICAgfVxuXG4gICAgICAgIGNvbnRlbnRFbC5zdHlsZS50b3AgPSBgJHt0b3B9cHhgO1xuICAgICAgICBjb250ZW50RWwuc3R5bGUubGVmdCA9IGAke2xlZnR9cHhgO1xuICAgICAgICBjb250ZW50RWwuc3R5bGUuYm90dG9tID0gJ2F1dG8nO1xuICAgICAgICBjb250ZW50RWwuc3R5bGUucmlnaHQgPSAnYXV0byc7XG4gICAgICAgIGNvbnRlbnRFbC5zdHlsZS50cmFuc2xhdGUgPSAnbm9uZSc7XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgY29uc29sZS53YXJuKCdGYWlsZWQgdG8gc2hvdyBwb3BvdmVyOicsIGUpO1xuICAgIH1cbiAgfVxuXG4gICNoaWRlUG9wb3ZlcigpIHtcbiAgICBjb25zdCBjb250ZW50RWwgPSB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdjb250ZW50JykgYXMgSFRNTEVsZW1lbnQgJiB7IGhpZGVQb3BvdmVyKCk6IHZvaWQgfSB8IG51bGw7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnRlbnRFbD8uaGlkZVBvcG92ZXIoKTtcbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ0ZhaWxlZCB0byBoaWRlIHBvcG92ZXI6JywgZSk7XG4gICAgfVxuICB9XG5cbiAgI2hhbmRsZVRyaWdnZXJDbGljayA9IChlOiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHRyaWdnZXJTbG90ID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdzbG90W25hbWU9XCJ0cmlnZ2VyXCJdJykgYXMgSFRNTFNsb3RFbGVtZW50IHwgbnVsbDtcbiAgICBjb25zdCBhc3NpZ25lZCA9IHRyaWdnZXJTbG90Py5hc3NpZ25lZEVsZW1lbnRzKCk7XG4gICAgaWYgKGFzc2lnbmVkICYmIGFzc2lnbmVkLmxlbmd0aCA+IDApIHtcbiAgICAgIGUuc3RvcFByb3BhZ2F0aW9uKCk7XG4gICAgICB0aGlzLnRvZ2dsZSgpO1xuICAgIH1cbiAgfTtcblxuICAjaGFuZGxlUG9pbnRlckVudGVyID0gKCkgPT4ge1xuICAgIGNsZWFyVGltZW91dCh0aGlzLiNob3ZlckhpZGVUaW1lb3V0KTtcbiAgICB0aGlzLiNob3ZlclNob3dUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICB0aGlzLiNzaG93UG9wb3ZlcigpO1xuICAgIH0sIDE1MCk7XG4gIH07XG5cbiAgI2hhbmRsZVBvaW50ZXJMZWF2ZSA9ICgpID0+IHtcbiAgICBjbGVhclRpbWVvdXQodGhpcy4jaG92ZXJTaG93VGltZW91dCk7XG4gICAgdGhpcy4jaG92ZXJIaWRlVGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgdGhpcy4jaGlkZVBvcG92ZXIoKTtcbiAgICB9LCAzMDApO1xuICB9O1xuXG4gICNoYW5kbGVQb3BvdmVyUG9pbnRlckVudGVyID0gKCkgPT4ge1xuICAgIGNsZWFyVGltZW91dCh0aGlzLiNob3ZlckhpZGVUaW1lb3V0KTtcbiAgfTtcblxuICAjaGFuZGxlUG9wb3ZlclBvaW50ZXJMZWF2ZSA9ICgpID0+IHtcbiAgICB0aGlzLiNob3ZlckhpZGVUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICB0aGlzLiNoaWRlUG9wb3ZlcigpO1xuICAgIH0sIDMwMCk7XG4gIH07XG5cbiAgI2hhbmRsZVBvcG92ZXJUb2dnbGUgPSAoZTogVG9nZ2xlRXZlbnQpID0+IHtcbiAgICBjb25zdCBpc09wZW4gPSBlLm5ld1N0YXRlID09PSAnb3Blbic7XG5cbiAgICBpZiAoaXNPcGVuICE9PSB0aGlzLm9wZW4pIHtcbiAgICAgIHRoaXMub3BlbiA9IGlzT3BlbjtcbiAgICB9XG5cbiAgICBjb25zdCB0cmlnZ2VyU2xvdCA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3Rvcignc2xvdFtuYW1lPVwidHJpZ2dlclwiXScpIGFzIEhUTUxTbG90RWxlbWVudCB8IG51bGw7XG4gICAgY29uc3QgdHJpZ2dlckJ1dHRvbiA9IHRyaWdnZXJTbG90Py5hc3NpZ25lZEVsZW1lbnRzKCk/LlswXTtcbiAgICBpZiAodHJpZ2dlckJ1dHRvbikge1xuICAgICAgdHJpZ2dlckJ1dHRvbi5zZXRBdHRyaWJ1dGUoJ2FyaWEtZXhwYW5kZWQnLCBTdHJpbmcoaXNPcGVuKSk7XG4gICAgfVxuXG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KGlzT3BlbiA/IG5ldyBQZlBvcG92ZXJTaG93RXZlbnQoKSA6IG5ldyBQZlBvcG92ZXJIaWRlRXZlbnQoKSk7XG4gIH07XG5cbiAgI2hhbmRsZUNsb3NlID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgZS5zdG9wUHJvcGFnYXRpb24oKTtcbiAgICB0aGlzLiNoaWRlUG9wb3ZlcigpO1xuICB9O1xuXG4gICNvbkhlYWRlclNsb3RDaGFuZ2UoZTogRXZlbnQpIHtcbiAgICBjb25zdCBzbG90ID0gZS50YXJnZXQgYXMgSFRNTFNsb3RFbGVtZW50O1xuICAgIHRoaXMuI2hhc0hlYWRlckNvbnRlbnQgPSBzbG90LmFzc2lnbmVkTm9kZXMoKS5sZW5ndGggPiAwO1xuICAgIHRoaXMucmVxdWVzdFVwZGF0ZSgpO1xuICB9XG5cbiAgI29uRm9vdGVyU2xvdENoYW5nZShlOiBFdmVudCkge1xuICAgIGNvbnN0IHNsb3QgPSBlLnRhcmdldCBhcyBIVE1MU2xvdEVsZW1lbnQ7XG4gICAgdGhpcy4jaGFzRm9vdGVyQ29udGVudCA9IHNsb3QuYXNzaWduZWROb2RlcygpLmxlbmd0aCA+IDA7XG4gICAgdGhpcy5yZXF1ZXN0VXBkYXRlKCk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIGNsZWFyVGltZW91dCh0aGlzLiNob3ZlclNob3dUaW1lb3V0KTtcbiAgICBjbGVhclRpbWVvdXQodGhpcy4jaG92ZXJIaWRlVGltZW91dCk7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAncGYtdjYtcG9wb3Zlcic6IFBmVjZQb3BvdmVyO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxuICBhbGlnbi1zZWxmOiBzdGFydDtcXFxcbiAganVzdGlmeS1zZWxmOiBzdGFydDtcXFxcblxcXFxuICAtLWRpc3RhbmNlOiA4cHg7XFxcXG4gIC0tbWluLXdpZHRoOiAxMDBweDtcXFxcbiAgLS1tYXgtd2lkdGg6IDMwMHB4O1xcXFxuXFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyLS1Gb250U2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyLS1Cb3hTaGFkb3c6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm94LXNoYWRvdy0tbWQpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlci0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0taGlnaC1jb250cmFzdC0tcmVndWxhcik7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1oaWdoLWNvbnRyYXN0KTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXItLUJvcmRlclJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tbWVkaXVtKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2NvbnRlbnQtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tZmxvYXRpbmctLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0Jsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19jb250ZW50LS1QYWRkaW5nSW5saW5lRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0Jsb2NrRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fY2xvc2UtLUluc2V0QmxvY2tTdGFydDogY2FsYyh2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2NvbnRlbnQtLVBhZGRpbmdCbG9ja1N0YXJ0KSAtIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCkpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fY2xvc2UtLUluc2V0SW5saW5lRW5kOiB2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2NvbnRlbnQtLVBhZGRpbmdJbmxpbmVFbmQpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fY2xvc2UtLXNpYmxpbmctLVBhZGRpbmdJbmxpbmVFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS0yeGwpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9faGVhZGVyLS1NYXJnaW5CbG9ja0VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX3RpdGxlLXRleHQtLUNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX3RpdGxlLXRleHQtLUZvbnRXZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0td2VpZ2h0LS1ib2R5LS1ib2xkKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX3RpdGxlLXRleHQtLUZvbnRTaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fZm9vdGVyLS1NYXJnaW5CbG9ja1N0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVdpZHRoOiAxLjVyZW07XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSGVpZ2h0OiAxLjVyZW07XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tQmFja2dyb3VuZENvbG9yOiB2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2NvbnRlbnQtLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tQm94U2hhZG93OiB2YXIoLS1wZi12Ni1jLXBvcG92ZXItLUJveFNoYWRvdyk7XFxcXG59XFxcXG5cXFxcbiN0cmlnZ2VyIHtcXFxcbiAgYW5jaG9yLW5hbWU6IC0tdHJpZ2dlcjtcXFxcbiAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxufVxcXFxuXFxcXG4jY29udGVudCB7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHRyYW5zcGFyZW50O1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBvdmVyZmxvdzogaW5pdGlhbDtcXFxcbiAgcG9zaXRpb246IGZpeGVkO1xcXFxuICBwb3NpdGlvbi1hbmNob3I6IC0tdHJpZ2dlcjtcXFxcbiAgbWluLXdpZHRoOiB2YXIoLS1taW4td2lkdGgpO1xcXFxuICBtYXgtd2lkdGg6IHZhcigtLW1heC13aWR0aCk7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi12Ni1jLXBvcG92ZXItLUZvbnRTaXplKTtcXFxcbiAgYm9yZGVyOiB2YXIoLS1wZi12Ni1jLXBvcG92ZXItLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1wZi12Ni1jLXBvcG92ZXItLUJvcmRlckNvbG9yKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJSYWRpdXMpO1xcXFxuICBib3gtc2hhZG93OiB2YXIoLS1wZi12Ni1jLXBvcG92ZXItLUJveFNoYWRvdyk7XFxcXG4gIHRvcDogYXV0bztcXFxcbiAgYm90dG9tOiBhdXRvO1xcXFxuICBsZWZ0OiBhdXRvO1xcXFxuICByaWdodDogYXV0bztcXFxcbn1cXFxcblxcXFxuI3BvcG92ZXIge1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIHBhZGRpbmctYmxvY2stc3RhcnQ6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0Jsb2NrU3RhcnQpO1xcXFxuICBwYWRkaW5nLWJsb2NrLWVuZDogdmFyKC0tcGYtdjYtYy1wb3BvdmVyX19jb250ZW50LS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tcGYtdjYtYy1wb3BvdmVyX19jb250ZW50LS1QYWRkaW5nSW5saW5lU3RhcnQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1lbmQ6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fY29udGVudC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJSYWRpdXMpO1xcXFxufVxcXFxuXFxcXG4jYXJyb3cge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIHRvcDogdmFyKC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0LCBhdXRvKTtcXFxcbiAgcmlnaHQ6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lRW5kLCBhdXRvKTtcXFxcbiAgYm90dG9tOiB2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldEJsb2NrRW5kLCBhdXRvKTtcXFxcbiAgbGVmdDogdmFyKC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVTdGFydCwgYXV0byk7XFxcXG4gIHdpZHRoOiB2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1XaWR0aCk7XFxcXG4gIGhlaWdodDogdmFyKC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSGVpZ2h0KTtcXFxcbiAgcG9pbnRlci1ldmVudHM6IG5vbmU7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIGJvcmRlcjogdmFyKC0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJXaWR0aCkgc29saWQgdmFyKC0tcGYtdjYtYy1wb3BvdmVyLS1Cb3JkZXJDb2xvcik7XFxcXG4gIGJveC1zaGFkb3c6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUJveFNoYWRvdyk7XFxcXG4gIHRyYW5zZm9ybTogdHJhbnNsYXRlWCh2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1UcmFuc2xhdGVYLCAwKSkgdHJhbnNsYXRlWSh2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1UcmFuc2xhdGVZLCAwKSkgcm90YXRlKHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZSwgMCkpO1xcXFxufVxcXFxuXFxcXG4jY2xvc2Uge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIGluc2V0LWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Nsb3NlLS1JbnNldEJsb2NrU3RhcnQpO1xcXFxuICBpbnNldC1pbmxpbmUtZW5kOiB2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Nsb3NlLS1JbnNldElubGluZUVuZCk7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWNvbnRyb2wtLWhvcml6b250YWwtLXBsYWluLS1kZWZhdWx0KTtcXFxcbiAgYmFja2dyb3VuZDogdHJhbnNwYXJlbnQ7XFxcXG4gIGJvcmRlcjogbm9uZTtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1pY29uLS1jb2xvci0tcmVndWxhcik7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGp1c3RpZnktY29udGVudDogY2VudGVyO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0taG92ZXIpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbjpob3N0KDpub3QoW2Nsb3NlYWJsZV0pKSAjY2xvc2UsXFxcXG46aG9zdChbY2xvc2VhYmxlPVxcXFxcXFwiZmFsc2VcXFxcXFxcIl0pICNjbG9zZSB7XFxcXG4gIGRpc3BsYXk6IG5vbmU7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtjbG9zZWFibGVdKSAjaGVhZGVyLFxcXFxuOmhvc3QoW2Nsb3NlYWJsZT1cXFxcXFxcIlxcXFxcXFwiXSkgI2hlYWRlciB7XFxcXG4gIHBhZGRpbmctaW5saW5lLWVuZDogdmFyKC0tcGYtdjYtYy1wb3BvdmVyX19jbG9zZS0tc2libGluZy0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG59XFxcXG5cXFxcbiNoZWFkZXIge1xcXFxuICBtYXJnaW4tYmxvY2stZW5kOiAwO1xcXFxuXFxcXG4gIFxcXFx1MDAyNi5oYXMtY29udGVudCB7XFxcXG4gICAgbWFyZ2luLWJsb2NrLWVuZDogdmFyKC0tcGYtdjYtYy1wb3BvdmVyX19oZWFkZXItLU1hcmdpbkJsb2NrRW5kKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jaGVhZGVyLWNvbnRlbnQge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4OiAwIDAgYXV0bztcXFxcblxcXFxuICA6OnNsb3R0ZWQoaDEpLFxcXFxuICA6OnNsb3R0ZWQoaDIpLFxcXFxuICA6OnNsb3R0ZWQoaDMpLFxcXFxuICA6OnNsb3R0ZWQoaDQpLFxcXFxuICA6OnNsb3R0ZWQoaDUpLFxcXFxuICA6OnNsb3R0ZWQoaDYpIHtcXFxcbiAgICBtaW4td2lkdGg6IDAgIWltcG9ydGFudDtcXFxcbiAgICBmb250LXNpemU6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fdGl0bGUtdGV4dC0tRm9udFNpemUpICFpbXBvcnRhbnQ7XFxcXG4gICAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fdGl0bGUtdGV4dC0tRm9udFdlaWdodCkgIWltcG9ydGFudDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdjYtYy1wb3BvdmVyX190aXRsZS10ZXh0LS1Db2xvcikgIWltcG9ydGFudDtcXFxcbiAgICBvdmVyZmxvdy13cmFwOiBicmVhay13b3JkICFpbXBvcnRhbnQ7XFxcXG4gICAgbWFyZ2luOiAwICFpbXBvcnRhbnQ7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI2JvZHkge1xcXFxuICB3b3JkLXdyYXA6IGJyZWFrLXdvcmQ7XFxcXG59XFxcXG5cXFxcbiNmb290ZXIge1xcXFxuICBtYXJnaW4tYmxvY2stc3RhcnQ6IDA7XFxcXG5cXFxcbiAgXFxcXHUwMDI2Lmhhcy1jb250ZW50IHtcXFxcbiAgICBtYXJnaW4tYmxvY2stc3RhcnQ6IHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fZm9vdGVyLS1NYXJnaW5CbG9ja1N0YXJ0KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG46aG9zdChbaGFzLWF1dG8td2lkdGhdKSAjY29udGVudCB7XFxcXG4gIG1pbi13aWR0aDogdW5zZXQ7XFxcXG4gIG1heC13aWR0aDogdW5zZXQ7XFxcXG4gIHdpZHRoOiBtYXgtY29udGVudDtcXFxcbn1cXFxcblxcXFxuLyogUG9zaXRpb246IHRvcCAoZGVmYXVsdCkgKi9cXFxcbiNjb250ZW50LnBvc2l0aW9uLXRvcCB7XFxcXG4gIGJvdHRvbTogYW5jaG9yKC0tdHJpZ2dlciB0b3ApO1xcXFxuICBsZWZ0OiBhbmNob3IoLS10cmlnZ2VyIGNlbnRlcik7XFxcXG4gIHRyYW5zbGF0ZTogLTUwJSAwO1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja0VuZDogY2FsYyh2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1IZWlnaHQpIC8gLTIpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lU3RhcnQ6IDUwJTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1UcmFuc2xhdGVYOiAtNTAlO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLWJvdHRvbSB7XFxcXG4gIHRvcDogYW5jaG9yKC0tdHJpZ2dlciBib3R0b20pO1xcXFxuICBsZWZ0OiBhbmNob3IoLS10cmlnZ2VyIGNlbnRlcik7XFxcXG4gIHRyYW5zbGF0ZTogLTUwJSAwO1xcXFxuICBtYXJnaW4tdG9wOiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiBjYWxjKHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUhlaWdodCkgLyAtMik7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVTdGFydDogNTAlO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVRyYW5zbGF0ZVg6IC01MCU7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tbGVmdCB7XFxcXG4gIHRvcDogYW5jaG9yKC0tdHJpZ2dlciBjZW50ZXIpO1xcXFxuICByaWdodDogYW5jaG9yKC0tdHJpZ2dlciBsZWZ0KTtcXFxcbiAgdHJhbnNsYXRlOiAwIC01MCU7XFxcXG4gIG1hcmdpbi1yaWdodDogdmFyKC0tZGlzdGFuY2UpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0QmxvY2tTdGFydDogNTAlO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lRW5kOiBjYWxjKHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVdpZHRoKSAvIC0yKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1UcmFuc2xhdGVZOiAtNTAlO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLXJpZ2h0IHtcXFxcbiAgdG9wOiBhbmNob3IoLS10cmlnZ2VyIGNlbnRlcik7XFxcXG4gIGxlZnQ6IGFuY2hvcigtLXRyaWdnZXIgcmlnaHQpO1xcXFxuICB0cmFuc2xhdGU6IDAgLTUwJTtcXFxcbiAgbWFyZ2luLWxlZnQ6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldEJsb2NrU3RhcnQ6IDUwJTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldElubGluZVN0YXJ0OiBjYWxjKHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVdpZHRoKSAvIC0yKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1UcmFuc2xhdGVZOiAtNTAlO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLXRvcC1zdGFydCB7XFxcXG4gIGJvdHRvbTogYW5jaG9yKC0tdHJpZ2dlciB0b3ApO1xcXFxuICBsZWZ0OiBhbmNob3IoLS10cmlnZ2VyIGxlZnQpO1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja0VuZDogY2FsYyh2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1IZWlnaHQpIC8gLTIpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tdG9wLWVuZCB7XFxcXG4gIGJvdHRvbTogYW5jaG9yKC0tdHJpZ2dlciB0b3ApO1xcXFxuICByaWdodDogYW5jaG9yKC0tdHJpZ2dlciByaWdodCk7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldEJsb2NrRW5kOiBjYWxjKHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUhlaWdodCkgLyAtMik7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tYm90dG9tLXN0YXJ0IHtcXFxcbiAgdG9wOiBhbmNob3IoLS10cmlnZ2VyIGJvdHRvbSk7XFxcXG4gIGxlZnQ6IGFuY2hvcigtLXRyaWdnZXIgbGVmdCk7XFxcXG4gIG1hcmdpbi10b3A6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldEJsb2NrU3RhcnQ6IGNhbGModmFyKC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSGVpZ2h0KSAvIC0yKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldElubGluZVN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLWJvdHRvbS1lbmQge1xcXFxuICB0b3A6IGFuY2hvcigtLXRyaWdnZXIgYm90dG9tKTtcXFxcbiAgcmlnaHQ6IGFuY2hvcigtLXRyaWdnZXIgcmlnaHQpO1xcXFxuICBtYXJnaW4tdG9wOiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiBjYWxjKHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUhlaWdodCkgLyAtMik7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRJbmxpbmVFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tbGVmdC1zdGFydCB7XFxcXG4gIHRvcDogYW5jaG9yKC0tdHJpZ2dlciB0b3ApO1xcXFxuICByaWdodDogYW5jaG9yKC0tdHJpZ2dlciBsZWZ0KTtcXFxcbiAgbWFyZ2luLXJpZ2h0OiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lRW5kOiBjYWxjKHZhcigtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVdpZHRoKSAvIC0yKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1Sb3RhdGU6IDQ1ZGVnO1xcXFxufVxcXFxuXFxcXG4jY29udGVudC5wb3NpdGlvbi1sZWZ0LWVuZCB7XFxcXG4gIGJvdHRvbTogYW5jaG9yKC0tdHJpZ2dlciBib3R0b20pO1xcXFxuICByaWdodDogYW5jaG9yKC0tdHJpZ2dlciBsZWZ0KTtcXFxcbiAgbWFyZ2luLXJpZ2h0OiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja0VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldElubGluZUVuZDogY2FsYyh2YXIoLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1XaWR0aCkgLyAtMik7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tUm90YXRlOiA0NWRlZztcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQucG9zaXRpb24tcmlnaHQtc3RhcnQge1xcXFxuICB0b3A6IGFuY2hvcigtLXRyaWdnZXIgdG9wKTtcXFxcbiAgbGVmdDogYW5jaG9yKC0tdHJpZ2dlciByaWdodCk7XFxcXG4gIG1hcmdpbi1sZWZ0OiB2YXIoLS1kaXN0YW5jZSk7XFxcXG4gIC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tSW5zZXRCbG9ja1N0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lU3RhcnQ6IGNhbGModmFyKC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tV2lkdGgpIC8gLTIpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLXJpZ2h0LWVuZCB7XFxcXG4gIGJvdHRvbTogYW5jaG9yKC0tdHJpZ2dlciBib3R0b20pO1xcXFxuICBsZWZ0OiBhbmNob3IoLS10cmlnZ2VyIHJpZ2h0KTtcXFxcbiAgbWFyZ2luLWxlZnQ6IHZhcigtLWRpc3RhbmNlKTtcXFxcbiAgLS1wZi12Ni1jLXBvcG92ZXJfX2Fycm93LS1JbnNldEJsb2NrRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLUluc2V0SW5saW5lU3RhcnQ6IGNhbGModmFyKC0tcGYtdjYtYy1wb3BvdmVyX19hcnJvdy0tV2lkdGgpIC8gLTIpO1xcXFxuICAtLXBmLXY2LWMtcG9wb3Zlcl9fYXJyb3ctLVJvdGF0ZTogNDVkZWc7XFxcXG59XFxcXG5cXFxcbiNjb250ZW50LnBvc2l0aW9uLWF1dG8ge1xcXFxuICBib3R0b206IGFuY2hvcigtLXRyaWdnZXIgdG9wKTtcXFxcbiAgbGVmdDogYW5jaG9yKC0tdHJpZ2dlciBjZW50ZXIpO1xcXFxuICB0cmFuc2xhdGU6IC01MCUgMDtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tZGlzdGFuY2UpO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksTUFBTSxlQUFlO0FBQzFDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sMDJVQUFnM1UsQ0FBQztBQUFFLElBQU8sd0JBQVE7OztBRGdCaDdVLElBQU0scUJBQU4sY0FBaUMsTUFBTTtBQUFBLEVBQzVDLE9BQU87QUFBQSxFQUNQLGNBQWM7QUFDWixVQUFNLG1CQUFtQixFQUFFLFNBQVMsS0FBSyxDQUFDO0FBQUEsRUFDNUM7QUFDRjtBQUtPLElBQU0scUJBQU4sY0FBaUMsTUFB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
