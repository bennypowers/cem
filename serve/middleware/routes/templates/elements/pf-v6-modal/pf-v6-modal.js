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

// elements/pf-v6-modal/pf-v6-modal.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/pf-v6-modal/pf-v6-modal.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n\\n  /* Modal box */\\n  --pf-v6-c-modal-box--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --pf-v6-c-modal-box--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --pf-v6-c-modal-box--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --pf-v6-c-modal-box--BorderRadius: var(--pf-t--global--border--radius--large);\\n  --pf-v6-c-modal-box--BoxShadow: var(--pf-t--global--box-shadow--lg);\\n  --pf-v6-c-modal-box--ZIndex: var(--pf-t--global--z-index--xl);\\n  --pf-v6-c-modal-box--Width: 100%;\\n  --pf-v6-c-modal-box--MaxWidth: calc(100% - var(--pf-t--global--spacer--xl));\\n  --pf-v6-c-modal-box--m-sm--MaxWidth: 35rem;\\n  --pf-v6-c-modal-box--m-md--Width: 52.5rem;\\n  --pf-v6-c-modal-box--m-lg--MaxWidth: 70rem;\\n  --pf-v6-c-modal-box--MaxHeight: calc(100% - var(--pf-t--global--spacer--2xl));\\n\\n  /* Align top position */\\n  --pf-v6-c-modal-box--m-align-top--spacer: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-modal-box--m-align-top--InsetBlockStart: var(--pf-v6-c-modal-box--m-align-top--spacer);\\n  --pf-v6-c-modal-box--m-align-top--MaxHeight: calc(100% - min(var(--pf-v6-c-modal-box--m-align-top--spacer), var(--pf-t--global--spacer--2xl)) - var(--pf-v6-c-modal-box--m-align-top--spacer));\\n  --pf-v6-c-modal-box--m-align-top--MaxWidth: calc(100% - min(var(--pf-v6-c-modal-box--m-align-top--spacer) * 2, var(--pf-t--global--spacer--xl)));\\n\\n  /* Header */\\n  --pf-v6-c-modal-box__header--PaddingBlockStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__header--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-modal-box__header--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__header--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__header--Gap: var(--pf-t--global--spacer--md);\\n\\n  /* Header main */\\n  --pf-v6-c-modal-box__header-main--Gap: var(--pf-t--global--spacer--md);\\n  --pf-v6-c-modal-box__header-main--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--default);\\n\\n  /* Title */\\n  --pf-v6-c-modal-box__title--LineHeight: var(--pf-t--global--font--line-height--heading);\\n  --pf-v6-c-modal-box__title--FontFamily: var(--pf-t--global--font--family--heading);\\n  --pf-v6-c-modal-box__title--FontWeight: var(--pf-t--global--font--weight--heading--default);\\n  --pf-v6-c-modal-box__title--FontSize: var(--pf-t--global--font--size--heading--md);\\n\\n  /* Description */\\n  --pf-v6-c-modal-box__description--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --pf-v6-c-modal-box__description--Color: var(--pf-t--global--text--color--subtle);\\n\\n  /* Body */\\n  --pf-v6-c-modal-box__body--MinHeight: calc(var(--pf-t--global--font--size--body--default) * var(--pf-t--global--font--line-height--body));\\n  --pf-v6-c-modal-box__body--PaddingBlockStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__body--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__body--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__body--PaddingBlockEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__header--body--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n\\n  /* Close button */\\n  --pf-v6-c-modal-box__close--InsetBlockStart: var(--pf-v6-c-modal-box__header--PaddingBlockStart);\\n  --pf-v6-c-modal-box__close--InsetInlineEnd: var(--pf-v6-c-modal-box__header--PaddingInlineEnd);\\n\\n  /* Footer */\\n  --pf-v6-c-modal-box__footer--PaddingBlockStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__footer--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__footer--PaddingBlockEnd: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__footer--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --pf-v6-c-modal-box__footer--Gap: var(--pf-t--global--spacer--md);\\n}\\n\\n@media (min-width: 1200px) {\\n  :host {\\n    --pf-v6-c-modal-box--m-align-top--spacer: var(--pf-t--global--spacer--xl);\\n  }\\n}\\n\\ndialog {\\n  position: fixed;\\n  inset: 0;\\n  z-index: var(--pf-v6-c-modal-box--ZIndex);\\n  width: var(--pf-v6-c-modal-box--Width);\\n  max-width: min(var(--pf-v6-c-modal-box--MaxWidth), 100vw);\\n  max-height: min(var(--pf-v6-c-modal-box--MaxHeight), 100vh);\\n  margin: auto;\\n  padding: 0;\\n  overflow: hidden;\\n  background-color: var(--pf-v6-c-modal-box--BackgroundColor);\\n  border: var(--pf-v6-c-modal-box--BorderWidth) solid var(--pf-v6-c-modal-box--BorderColor);\\n  border-radius: var(--pf-v6-c-modal-box--BorderRadius);\\n  box-shadow: var(--pf-v6-c-modal-box--BoxShadow);\\n\\n  \\u0026[open] {\\n    display: flex;\\n    flex-direction: column;\\n  }\\n\\n  \\u0026::backdrop {\\n    background-color: var(--pf-v6-c-backdrop--BackgroundColor, var(--pf-t--global--background--color--backdrop--default));\\n  }\\n\\n  :host([variant=\\"small\\"]) \\u0026 {\\n    --pf-v6-c-modal-box--Width: var(--pf-v6-c-modal-box--m-sm--MaxWidth);\\n  }\\n\\n  :host([variant=\\"medium\\"]) \\u0026 {\\n    --pf-v6-c-modal-box--Width: var(--pf-v6-c-modal-box--m-md--Width);\\n  }\\n\\n  :host([variant=\\"large\\"]) \\u0026 {\\n    --pf-v6-c-modal-box--Width: var(--pf-v6-c-modal-box--m-lg--MaxWidth);\\n  }\\n\\n  :host([position=\\"top\\"]) \\u0026 {\\n    margin-block-start: var(--pf-v6-c-modal-box--m-align-top--InsetBlockStart);\\n    margin-block-end: auto;\\n    --pf-v6-c-modal-box--MaxWidth: var(--pf-v6-c-modal-box--m-align-top--MaxWidth);\\n    --pf-v6-c-modal-box--MaxHeight: var(--pf-v6-c-modal-box--m-align-top--MaxHeight);\\n  }\\n}\\n\\n#header {\\n  display: flex;\\n  flex-direction: column;\\n  flex-shrink: 0;\\n  padding-block-start: var(--pf-v6-c-modal-box__header--PaddingBlockStart);\\n  padding-block-end: var(--pf-v6-c-modal-box__header--PaddingBlockEnd);\\n  padding-inline-start: var(--pf-v6-c-modal-box__header--PaddingInlineStart);\\n  padding-inline-end: var(--pf-v6-c-modal-box__header--PaddingInlineEnd);\\n  gap: var(--pf-v6-c-modal-box__header--Gap);\\n}\\n\\n#header-main {\\n  display: flex;\\n  flex-direction: column;\\n  gap: var(--pf-v6-c-modal-box__header-main--Gap);\\n  padding-block-start: var(--pf-v6-c-modal-box__header-main--PaddingBlockStart);\\n  flex-grow: 1;\\n  min-width: 0;\\n}\\n\\n::slotted([slot=\\"header\\"]) {\\n  margin: 0 !important;\\n  overflow: hidden;\\n  text-overflow: ellipsis;\\n  white-space: nowrap;\\n  font-family: var(--pf-v6-c-modal-box__title--FontFamily);\\n  font-size: var(--pf-v6-c-modal-box__title--FontSize);\\n  font-weight: var(--pf-v6-c-modal-box__title--FontWeight);\\n  line-height: var(--pf-v6-c-modal-box__title--LineHeight);\\n}\\n\\n#description {\\n  font-size: var(--pf-v6-c-modal-box__description--FontSize);\\n  color: var(--pf-v6-c-modal-box__description--Color);\\n}\\n\\n#body {\\n  flex: 1 1 auto;\\n  min-height: var(--pf-v6-c-modal-box__body--MinHeight);\\n  padding-block-start: var(--pf-v6-c-modal-box__body--PaddingBlockStart);\\n  padding-inline-end: var(--pf-v6-c-modal-box__body--PaddingInlineEnd);\\n  padding-inline-start: var(--pf-v6-c-modal-box__body--PaddingInlineStart);\\n  padding-block-end: var(--pf-v6-c-modal-box__body--PaddingBlockEnd);\\n  overflow-x: hidden;\\n  overflow-y: auto;\\n  overscroll-behavior: contain;\\n  word-break: break-word;\\n  -webkit-overflow-scrolling: touch;\\n}\\n\\n#header:not([hidden]) ~ #body {\\n  --pf-v6-c-modal-box__body--PaddingBlockStart: var(--pf-v6-c-modal-box__header--body--PaddingBlockStart);\\n}\\n\\n#footer {\\n  display: flex;\\n  flex: 0 0 auto;\\n  align-items: center;\\n  gap: var(--pf-v6-c-modal-box__footer--Gap);\\n  padding-block-start: var(--pf-v6-c-modal-box__footer--PaddingBlockStart);\\n  padding-inline-end: var(--pf-v6-c-modal-box__footer--PaddingInlineEnd);\\n  padding-block-end: var(--pf-v6-c-modal-box__footer--PaddingBlockEnd);\\n  padding-inline-start: var(--pf-v6-c-modal-box__footer--PaddingInlineStart);\\n}\\n\\n#close {\\n  position: absolute;\\n  inset-block-start: var(--pf-v6-c-modal-box__close--InsetBlockStart);\\n  inset-inline-end: var(--pf-v6-c-modal-box__close--InsetInlineEnd);\\n\\n  \\u0026 svg {\\n    width: 1rem;\\n    height: 1rem;\\n    fill: currentColor;\\n  }\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n"'));
var pf_v6_modal_default = s;

// elements/pf-v6-modal/pf-v6-modal.ts
import "../pf-v6-button/pf-v6-button.js";
var PfModalOpenEvent = class extends Event {
  constructor() {
    super("open", { bubbles: true });
  }
};
var PfModalCloseEvent = class extends Event {
  returnValue;
  constructor(returnValue) {
    super("close", { bubbles: true });
    this.returnValue = returnValue;
  }
};
var PfModalCancelEvent = class extends Event {
  constructor() {
    super("cancel", { bubbles: true });
  }
};
var _open_dec, _position_dec, _variant_dec, _a, _PfV6Modal_decorators, _dialog, _cancelling, _hasHeader, _hasDescription, _hasFooter, _init, _variant, _position, _open, _PfV6Modal_instances, syncDialogState_fn, _onDialogClose, _onDialogCancel, _onCloseClick, _onHeaderSlotChange, _onDescriptionSlotChange, _onFooterSlotChange;
_PfV6Modal_decorators = [customElement("pf-v6-modal")];
var PfV6Modal = class extends (_a = LitElement, _variant_dec = [property({ reflect: true })], _position_dec = [property({ reflect: true })], _open_dec = [property({ type: Boolean, reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Modal_instances);
    __privateAdd(this, _dialog, null);
    __privateAdd(this, _cancelling, false);
    __privateAdd(this, _hasHeader, false);
    __privateAdd(this, _hasDescription, false);
    __privateAdd(this, _hasFooter, false);
    /** @see https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/returnValue */
    __publicField(this, "returnValue", "");
    __privateAdd(this, _variant, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
    __privateAdd(this, _position, __runInitializers(_init, 12, this)), __runInitializers(_init, 15, this);
    __privateAdd(this, _open, __runInitializers(_init, 16, this, false)), __runInitializers(_init, 19, this);
    __privateAdd(this, _onDialogClose, () => {
      if (this.open) {
        this.open = false;
      }
      if (__privateGet(this, _cancelling)) {
        this.dispatchEvent(new PfModalCancelEvent());
        __privateSet(this, _cancelling, false);
      } else {
        this.dispatchEvent(new PfModalCloseEvent(this.returnValue));
      }
    });
    __privateAdd(this, _onDialogCancel, (e) => {
      e.preventDefault();
      __privateSet(this, _cancelling, true);
      this.close();
    });
    __privateAdd(this, _onCloseClick, () => {
      __privateSet(this, _cancelling, true);
      this.close("cancel");
    });
    __privateAdd(this, _onHeaderSlotChange, (e) => {
      const slot = e.target;
      const had = __privateGet(this, _hasHeader);
      __privateSet(this, _hasHeader, slot.assignedNodes().length > 0);
      if (had !== __privateGet(this, _hasHeader)) this.requestUpdate();
    });
    __privateAdd(this, _onDescriptionSlotChange, (e) => {
      const slot = e.target;
      const had = __privateGet(this, _hasDescription);
      __privateSet(this, _hasDescription, slot.assignedNodes().length > 0);
      if (had !== __privateGet(this, _hasDescription)) this.requestUpdate();
    });
    __privateAdd(this, _onFooterSlotChange, (e) => {
      const slot = e.target;
      const had = __privateGet(this, _hasFooter);
      __privateSet(this, _hasFooter, slot.assignedNodes().length > 0);
      if (had !== __privateGet(this, _hasFooter)) this.requestUpdate();
    });
  }
  render() {
    return html`
      <dialog id="dialog"
              part="dialog"
              @close=${__privateGet(this, _onDialogClose)}
              @cancel=${__privateGet(this, _onDialogCancel)}>
        <pf-v6-button id="close"
                      part="close"
                      variant="plain"
                      aria-label="Close"
                      @click=${__privateGet(this, _onCloseClick)}>
          <svg viewBox="0 0 352 512"
               role="presentation">
            <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
          </svg>
        </pf-v6-button>
        <header id="header"
                part="header"
                ?hidden=${!__privateGet(this, _hasHeader)}>
          <div id="header-main">
            <slot name="header"
                  @slotchange=${__privateGet(this, _onHeaderSlotChange)}></slot>
            <div id="description"
                 part="description"
                 ?hidden=${!__privateGet(this, _hasDescription)}>
              <slot name="description"
                    @slotchange=${__privateGet(this, _onDescriptionSlotChange)}></slot>
            </div>
          </div>
        </header>
        <div id="body"
             part="body">
          <slot></slot>
        </div>
        <footer id="footer"
                part="footer"
                ?hidden=${!__privateGet(this, _hasFooter)}>
          <slot name="footer"
                @slotchange=${__privateGet(this, _onFooterSlotChange)}></slot>
        </footer>
      </dialog>
    `;
  }
  updated(changed) {
    if (changed.has("open")) {
      __privateMethod(this, _PfV6Modal_instances, syncDialogState_fn).call(this);
    }
  }
  firstUpdated() {
    __privateSet(this, _dialog, this.shadowRoot.getElementById("dialog"));
    __privateMethod(this, _PfV6Modal_instances, syncDialogState_fn).call(this);
  }
  /**
   * Opens the modal
   */
  show() {
    this.showModal();
  }
  /**
   * Opens the modal as a modal dialog (with backdrop)
   */
  showModal() {
    this.open = true;
  }
  /**
   * Closes the modal
   * @param returnValue - Optional return value
   */
  close(returnValue) {
    if (typeof returnValue === "string") {
      this.returnValue = returnValue;
      if (__privateGet(this, _dialog)) {
        __privateGet(this, _dialog).returnValue = returnValue;
      }
    }
    this.open = false;
  }
  /**
   * Toggles the modal open/closed
   */
  toggle() {
    if (this.open) {
      this.close();
    } else {
      this.showModal();
    }
  }
};
_init = __decoratorStart(_a);
_dialog = new WeakMap();
_cancelling = new WeakMap();
_hasHeader = new WeakMap();
_hasDescription = new WeakMap();
_hasFooter = new WeakMap();
_variant = new WeakMap();
_position = new WeakMap();
_open = new WeakMap();
_PfV6Modal_instances = new WeakSet();
syncDialogState_fn = function() {
  if (!__privateGet(this, _dialog)) return;
  if (this.open) {
    if (!__privateGet(this, _dialog).open) {
      __privateGet(this, _dialog).showModal();
      this.dispatchEvent(new PfModalOpenEvent());
    }
  } else {
    if (__privateGet(this, _dialog).open) {
      __privateGet(this, _dialog).close();
    }
  }
};
_onDialogClose = new WeakMap();
_onDialogCancel = new WeakMap();
_onCloseClick = new WeakMap();
_onHeaderSlotChange = new WeakMap();
_onDescriptionSlotChange = new WeakMap();
_onFooterSlotChange = new WeakMap();
__decorateElement(_init, 4, "variant", _variant_dec, PfV6Modal, _variant);
__decorateElement(_init, 4, "position", _position_dec, PfV6Modal, _position);
__decorateElement(_init, 4, "open", _open_dec, PfV6Modal, _open);
PfV6Modal = __decorateElement(_init, 0, "PfV6Modal", _PfV6Modal_decorators, PfV6Modal);
__publicField(PfV6Modal, "shadowRootOptions", {
  ...LitElement.shadowRootOptions,
  delegatesFocus: true
});
__publicField(PfV6Modal, "styles", pf_v6_modal_default);
__runInitializers(_init, 1, PfV6Modal);
export {
  PfModalCancelEvent,
  PfModalCloseEvent,
  PfModalOpenEvent,
  PfV6Modal
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtbW9kYWwvcGYtdjYtbW9kYWwudHMiLCAibGl0LWNzczplbGVtZW50cy9wZi12Ni1tb2RhbC9wZi12Ni1tb2RhbC5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtbW9kYWwuY3NzJztcblxuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcblxudHlwZSBNb2RhbFZhcmlhbnQgPSAnc21hbGwnIHwgJ21lZGl1bScgfCAnbGFyZ2UnO1xudHlwZSBNb2RhbFBvc2l0aW9uID0gJ3RvcCc7XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gbW9kYWwgb3BlbnNcbiAqL1xuZXhwb3J0IGNsYXNzIFBmTW9kYWxPcGVuRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCdvcGVuJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICB9XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gbW9kYWwgY2xvc2VzIG5vcm1hbGx5XG4gKi9cbmV4cG9ydCBjbGFzcyBQZk1vZGFsQ2xvc2VFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgcmV0dXJuVmFsdWU6IHN0cmluZztcbiAgY29uc3RydWN0b3IocmV0dXJuVmFsdWU6IHN0cmluZykge1xuICAgIHN1cGVyKCdjbG9zZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLnJldHVyblZhbHVlID0gcmV0dXJuVmFsdWU7XG4gIH1cbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBtb2RhbCBpcyBjYW5jZWxsZWQgKEVTQyBrZXkpXG4gKi9cbmV4cG9ydCBjbGFzcyBQZk1vZGFsQ2FuY2VsRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCdjYW5jZWwnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gIH1cbn1cblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IE1vZGFsIENvbXBvbmVudFxuICpcbiAqIEEgbW9kYWwgZGlzcGxheXMgaW1wb3J0YW50IGluZm9ybWF0aW9uIHRvIGEgdXNlciB3aXRob3V0IHJlcXVpcmluZyB0aGVtIHRvXG4gKiBuYXZpZ2F0ZSB0byBhIG5ldyBwYWdlLiBVc2VzIG5hdGl2ZSBgPGRpYWxvZz5gIGVsZW1lbnQgZm9yIHByb3BlciBtb2RhbFxuICogYmVoYXZpb3IuXG4gKlxuICogQGF0dHIge3N0cmluZ30gdmFyaWFudCAtIE1vZGFsIHdpZHRoIHZhcmlhbnQ6IHNtYWxsLCBtZWRpdW0sIGxhcmdlXG4gKiBAYXR0ciB7c3RyaW5nfSBwb3NpdGlvbiAtIE1vZGFsIHBvc2l0aW9uOiB0b3AgKGRlZmF1bHQgaXMgY2VudGVyZWQpXG4gKiBAYXR0ciB7Ym9vbGVhbn0gb3BlbiAtIFdoZXRoZXIgdGhlIG1vZGFsIGlzIGN1cnJlbnRseSBvcGVuXG4gKlxuICogQHNsb3QgaGVhZGVyIC0gTW9kYWwgdGl0bGUgKHR5cGljYWxseSBoMi1oNilcbiAqIEBzbG90IGRlc2NyaXB0aW9uIC0gT3B0aW9uYWwgZGVzY3JpcHRpb24gYmVsb3cgdGl0bGVcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBib2R5IGNvbnRlbnRcbiAqIEBzbG90IGZvb3RlciAtIEFjdGlvbiBidXR0b25zXG4gKlxuICogQGZpcmVzIHtQZk1vZGFsT3BlbkV2ZW50fSBvcGVuIC0gV2hlbiBtb2RhbCBvcGVuc1xuICogQGZpcmVzIHtQZk1vZGFsQ2xvc2VFdmVudH0gY2xvc2UgLSBXaGVuIG1vZGFsIGNsb3NlcyBub3JtYWxseVxuICogQGZpcmVzIHtQZk1vZGFsQ2FuY2VsRXZlbnR9IGNhbmNlbCAtIFdoZW4gbW9kYWwgaXMgY2FuY2VsbGVkIChFU0Mga2V5KVxuICovXG5AY3VzdG9tRWxlbWVudCgncGYtdjYtbW9kYWwnKVxuZXhwb3J0IGNsYXNzIFBmVjZNb2RhbCBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc2hhZG93Um9vdE9wdGlvbnM6IFNoYWRvd1Jvb3RJbml0ID0ge1xuICAgIC4uLkxpdEVsZW1lbnQuc2hhZG93Um9vdE9wdGlvbnMsXG4gICAgZGVsZWdhdGVzRm9jdXM6IHRydWUsXG4gIH07XG5cbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjZGlhbG9nOiBIVE1MRGlhbG9nRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjY2FuY2VsbGluZyA9IGZhbHNlO1xuICAjaGFzSGVhZGVyID0gZmFsc2U7XG4gICNoYXNEZXNjcmlwdGlvbiA9IGZhbHNlO1xuICAjaGFzRm9vdGVyID0gZmFsc2U7XG5cbiAgLyoqIEBzZWUgaHR0cHM6Ly9kZXZlbG9wZXIubW96aWxsYS5vcmcvZW4tVVMvZG9jcy9XZWIvQVBJL0hUTUxEaWFsb2dFbGVtZW50L3JldHVyblZhbHVlICovXG4gIHJldHVyblZhbHVlID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciB2YXJpYW50PzogTW9kYWxWYXJpYW50O1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgcG9zaXRpb24/OiBNb2RhbFBvc2l0aW9uO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3Igb3BlbiA9IGZhbHNlO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaWFsb2cgaWQ9XCJkaWFsb2dcIlxuICAgICAgICAgICAgICBwYXJ0PVwiZGlhbG9nXCJcbiAgICAgICAgICAgICAgQGNsb3NlPSR7dGhpcy4jb25EaWFsb2dDbG9zZX1cbiAgICAgICAgICAgICAgQGNhbmNlbD0ke3RoaXMuI29uRGlhbG9nQ2FuY2VsfT5cbiAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImNsb3NlXCJcbiAgICAgICAgICAgICAgICAgICAgICBwYXJ0PVwiY2xvc2VcIlxuICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkNsb3NlXCJcbiAgICAgICAgICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLiNvbkNsb3NlQ2xpY2t9PlxuICAgICAgICAgIDxzdmcgdmlld0JveD1cIjAgMCAzNTIgNTEyXCJcbiAgICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIj5cbiAgICAgICAgICAgIDxwYXRoIGQ9XCJNMjQyLjcyIDI1NmwxMDAuMDctMTAwLjA3YzEyLjI4LTEyLjI4IDEyLjI4LTMyLjE5IDAtNDQuNDhsLTIyLjI0LTIyLjI0Yy0xMi4yOC0xMi4yOC0zMi4xOS0xMi4yOC00NC40OCAwTDE3NiAxODkuMjggNzUuOTMgODkuMjFjLTEyLjI4LTEyLjI4LTMyLjE5LTEyLjI4LTQ0LjQ4IDBMOS4yMSAxMTEuNDVjLTEyLjI4IDEyLjI4LTEyLjI4IDMyLjE5IDAgNDQuNDhMMTA5LjI4IDI1NiA5LjIxIDM1Ni4wN2MtMTIuMjggMTIuMjgtMTIuMjggMzIuMTkgMCA0NC40OGwyMi4yNCAyMi4yNGMxMi4yOCAxMi4yOCAzMi4yIDEyLjI4IDQ0LjQ4IDBMMTc2IDMyMi43MmwxMDAuMDcgMTAwLjA3YzEyLjI4IDEyLjI4IDMyLjIgMTIuMjggNDQuNDggMGwyMi4yNC0yMi4yNGMxMi4yOC0xMi4yOCAxMi4yOC0zMi4xOSAwLTQ0LjQ4TDI0Mi43MiAyNTZ6XCI+PC9wYXRoPlxuICAgICAgICAgIDwvc3ZnPlxuICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgPGhlYWRlciBpZD1cImhlYWRlclwiXG4gICAgICAgICAgICAgICAgcGFydD1cImhlYWRlclwiXG4gICAgICAgICAgICAgICAgP2hpZGRlbj0keyF0aGlzLiNoYXNIZWFkZXJ9PlxuICAgICAgICAgIDxkaXYgaWQ9XCJoZWFkZXItbWFpblwiPlxuICAgICAgICAgICAgPHNsb3QgbmFtZT1cImhlYWRlclwiXG4gICAgICAgICAgICAgICAgICBAc2xvdGNoYW5nZT0ke3RoaXMuI29uSGVhZGVyU2xvdENoYW5nZX0+PC9zbG90PlxuICAgICAgICAgICAgPGRpdiBpZD1cImRlc2NyaXB0aW9uXCJcbiAgICAgICAgICAgICAgICAgcGFydD1cImRlc2NyaXB0aW9uXCJcbiAgICAgICAgICAgICAgICAgP2hpZGRlbj0keyF0aGlzLiNoYXNEZXNjcmlwdGlvbn0+XG4gICAgICAgICAgICAgIDxzbG90IG5hbWU9XCJkZXNjcmlwdGlvblwiXG4gICAgICAgICAgICAgICAgICAgIEBzbG90Y2hhbmdlPSR7dGhpcy4jb25EZXNjcmlwdGlvblNsb3RDaGFuZ2V9Pjwvc2xvdD5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICA8L2hlYWRlcj5cbiAgICAgICAgPGRpdiBpZD1cImJvZHlcIlxuICAgICAgICAgICAgIHBhcnQ9XCJib2R5XCI+XG4gICAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGZvb3RlciBpZD1cImZvb3RlclwiXG4gICAgICAgICAgICAgICAgcGFydD1cImZvb3RlclwiXG4gICAgICAgICAgICAgICAgP2hpZGRlbj0keyF0aGlzLiNoYXNGb290ZXJ9PlxuICAgICAgICAgIDxzbG90IG5hbWU9XCJmb290ZXJcIlxuICAgICAgICAgICAgICAgIEBzbG90Y2hhbmdlPSR7dGhpcy4jb25Gb290ZXJTbG90Q2hhbmdlfT48L3Nsb3Q+XG4gICAgICAgIDwvZm9vdGVyPlxuICAgICAgPC9kaWFsb2c+XG4gICAgYDtcbiAgfVxuXG4gIHByb3RlY3RlZCB1cGRhdGVkKGNoYW5nZWQ6IE1hcDxzdHJpbmcsIHVua25vd24+KSB7XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdvcGVuJykpIHtcbiAgICAgIHRoaXMuI3N5bmNEaWFsb2dTdGF0ZSgpO1xuICAgIH1cbiAgfVxuXG4gIHByb3RlY3RlZCBmaXJzdFVwZGF0ZWQoKSB7XG4gICAgdGhpcy4jZGlhbG9nID0gdGhpcy5zaGFkb3dSb290IS5nZXRFbGVtZW50QnlJZCgnZGlhbG9nJykgYXMgSFRNTERpYWxvZ0VsZW1lbnQ7XG4gICAgdGhpcy4jc3luY0RpYWxvZ1N0YXRlKCk7XG4gIH1cblxuICAjc3luY0RpYWxvZ1N0YXRlKCkge1xuICAgIGlmICghdGhpcy4jZGlhbG9nKSByZXR1cm47XG4gICAgaWYgKHRoaXMub3Blbikge1xuICAgICAgaWYgKCF0aGlzLiNkaWFsb2cub3Blbikge1xuICAgICAgICB0aGlzLiNkaWFsb2cuc2hvd01vZGFsKCk7XG4gICAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgUGZNb2RhbE9wZW5FdmVudCgpKTtcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgaWYgKHRoaXMuI2RpYWxvZy5vcGVuKSB7XG4gICAgICAgIHRoaXMuI2RpYWxvZy5jbG9zZSgpO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNvbkRpYWxvZ0Nsb3NlID0gKCkgPT4ge1xuICAgIC8vIFN5bmMgb3BlbiBwcm9wZXJ0eSAod2hpY2ggcmVmbGVjdHMgdG8gYXR0cmlidXRlKVxuICAgIGlmICh0aGlzLm9wZW4pIHtcbiAgICAgIHRoaXMub3BlbiA9IGZhbHNlO1xuICAgIH1cblxuICAgIGlmICh0aGlzLiNjYW5jZWxsaW5nKSB7XG4gICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IFBmTW9kYWxDYW5jZWxFdmVudCgpKTtcbiAgICAgIHRoaXMuI2NhbmNlbGxpbmcgPSBmYWxzZTtcbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBQZk1vZGFsQ2xvc2VFdmVudCh0aGlzLnJldHVyblZhbHVlKSk7XG4gICAgfVxuICB9O1xuXG4gICNvbkRpYWxvZ0NhbmNlbCA9IChlOiBFdmVudCkgPT4ge1xuICAgIGUucHJldmVudERlZmF1bHQoKTtcbiAgICB0aGlzLiNjYW5jZWxsaW5nID0gdHJ1ZTtcbiAgICB0aGlzLmNsb3NlKCk7XG4gIH07XG5cbiAgI29uQ2xvc2VDbGljayA9ICgpID0+IHtcbiAgICB0aGlzLiNjYW5jZWxsaW5nID0gdHJ1ZTtcbiAgICB0aGlzLmNsb3NlKCdjYW5jZWwnKTtcbiAgfTtcblxuICAjb25IZWFkZXJTbG90Q2hhbmdlID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgY29uc3Qgc2xvdCA9IGUudGFyZ2V0IGFzIEhUTUxTbG90RWxlbWVudDtcbiAgICBjb25zdCBoYWQgPSB0aGlzLiNoYXNIZWFkZXI7XG4gICAgdGhpcy4jaGFzSGVhZGVyID0gc2xvdC5hc3NpZ25lZE5vZGVzKCkubGVuZ3RoID4gMDtcbiAgICBpZiAoaGFkICE9PSB0aGlzLiNoYXNIZWFkZXIpIHRoaXMucmVxdWVzdFVwZGF0ZSgpO1xuICB9O1xuXG4gICNvbkRlc2NyaXB0aW9uU2xvdENoYW5nZSA9IChlOiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHNsb3QgPSBlLnRhcmdldCBhcyBIVE1MU2xvdEVsZW1lbnQ7XG4gICAgY29uc3QgaGFkID0gdGhpcy4jaGFzRGVzY3JpcHRpb247XG4gICAgdGhpcy4jaGFzRGVzY3JpcHRpb24gPSBzbG90LmFzc2lnbmVkTm9kZXMoKS5sZW5ndGggPiAwO1xuICAgIGlmIChoYWQgIT09IHRoaXMuI2hhc0Rlc2NyaXB0aW9uKSB0aGlzLnJlcXVlc3RVcGRhdGUoKTtcbiAgfTtcblxuICAjb25Gb290ZXJTbG90Q2hhbmdlID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgY29uc3Qgc2xvdCA9IGUudGFyZ2V0IGFzIEhUTUxTbG90RWxlbWVudDtcbiAgICBjb25zdCBoYWQgPSB0aGlzLiNoYXNGb290ZXI7XG4gICAgdGhpcy4jaGFzRm9vdGVyID0gc2xvdC5hc3NpZ25lZE5vZGVzKCkubGVuZ3RoID4gMDtcbiAgICBpZiAoaGFkICE9PSB0aGlzLiNoYXNGb290ZXIpIHRoaXMucmVxdWVzdFVwZGF0ZSgpO1xuICB9O1xuXG4gIC8qKlxuICAgKiBPcGVucyB0aGUgbW9kYWxcbiAgICovXG4gIHNob3coKSB7XG4gICAgdGhpcy5zaG93TW9kYWwoKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBPcGVucyB0aGUgbW9kYWwgYXMgYSBtb2RhbCBkaWFsb2cgKHdpdGggYmFja2Ryb3ApXG4gICAqL1xuICBzaG93TW9kYWwoKSB7XG4gICAgdGhpcy5vcGVuID0gdHJ1ZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBDbG9zZXMgdGhlIG1vZGFsXG4gICAqIEBwYXJhbSByZXR1cm5WYWx1ZSAtIE9wdGlvbmFsIHJldHVybiB2YWx1ZVxuICAgKi9cbiAgY2xvc2UocmV0dXJuVmFsdWU/OiBzdHJpbmcpIHtcbiAgICBpZiAodHlwZW9mIHJldHVyblZhbHVlID09PSAnc3RyaW5nJykge1xuICAgICAgdGhpcy5yZXR1cm5WYWx1ZSA9IHJldHVyblZhbHVlO1xuICAgICAgaWYgKHRoaXMuI2RpYWxvZykge1xuICAgICAgICB0aGlzLiNkaWFsb2cucmV0dXJuVmFsdWUgPSByZXR1cm5WYWx1ZTtcbiAgICAgIH1cbiAgICB9XG4gICAgdGhpcy5vcGVuID0gZmFsc2U7XG4gIH1cblxuICAvKipcbiAgICogVG9nZ2xlcyB0aGUgbW9kYWwgb3Blbi9jbG9zZWRcbiAgICovXG4gIHRvZ2dsZSgpIHtcbiAgICBpZiAodGhpcy5vcGVuKSB7XG4gICAgICB0aGlzLmNsb3NlKCk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuc2hvd01vZGFsKCk7XG4gICAgfVxuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LW1vZGFsJzogUGZWNk1vZGFsO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuXFxcXG4gIC8qIE1vZGFsIGJveCAqL1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94LS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWZsb2F0aW5nLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveC0tQm9yZGVyQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0taGlnaC1jb250cmFzdCk7XFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3gtLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWhpZ2gtY29udHJhc3QtLXJlZ3VsYXIpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94LS1Cb3JkZXJSYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLWxhcmdlKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveC0tQm94U2hhZG93OiB2YXIoLS1wZi10LS1nbG9iYWwtLWJveC1zaGFkb3ctLWxnKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveC0tWkluZGV4OiB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLXhsKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveC0tV2lkdGg6IDEwMCU7XFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3gtLU1heFdpZHRoOiBjYWxjKDEwMCUgLSB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teGwpKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveC0tbS1zbS0tTWF4V2lkdGg6IDM1cmVtO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94LS1tLW1kLS1XaWR0aDogNTIuNXJlbTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveC0tbS1sZy0tTWF4V2lkdGg6IDcwcmVtO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94LS1NYXhIZWlnaHQ6IGNhbGMoMTAwJSAtIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS0yeGwpKTtcXFxcblxcXFxuICAvKiBBbGlnbiB0b3AgcG9zaXRpb24gKi9cXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveC0tbS1hbGlnbi10b3AtLXNwYWNlcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveC0tbS1hbGlnbi10b3AtLUluc2V0QmxvY2tTdGFydDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3gtLW0tYWxpZ24tdG9wLS1zcGFjZXIpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94LS1tLWFsaWduLXRvcC0tTWF4SGVpZ2h0OiBjYWxjKDEwMCUgLSBtaW4odmFyKC0tcGYtdjYtYy1tb2RhbC1ib3gtLW0tYWxpZ24tdG9wLS1zcGFjZXIpLCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tMnhsKSkgLSB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveC0tbS1hbGlnbi10b3AtLXNwYWNlcikpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94LS1tLWFsaWduLXRvcC0tTWF4V2lkdGg6IGNhbGMoMTAwJSAtIG1pbih2YXIoLS1wZi12Ni1jLW1vZGFsLWJveC0tbS1hbGlnbi10b3AtLXNwYWNlcikgKiAyLCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teGwpKSk7XFxcXG5cXFxcbiAgLyogSGVhZGVyICovXFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2hlYWRlci0tUGFkZGluZ0Jsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2hlYWRlci0tUGFkZGluZ0Jsb2NrRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X19oZWFkZXItLVBhZGRpbmdJbmxpbmVFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2hlYWRlci0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X19oZWFkZXItLUdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcblxcXFxuICAvKiBIZWFkZXIgbWFpbiAqL1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X19oZWFkZXItbWFpbi0tR2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X19oZWFkZXItbWFpbi0tUGFkZGluZ0Jsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1jb250cm9sLS12ZXJ0aWNhbC0tZGVmYXVsdCk7XFxcXG5cXFxcbiAgLyogVGl0bGUgKi9cXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveF9fdGl0bGUtLUxpbmVIZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tbGluZS1oZWlnaHQtLWhlYWRpbmcpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X190aXRsZS0tRm9udEZhbWlseTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1mYW1pbHktLWhlYWRpbmcpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X190aXRsZS0tRm9udFdlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS13ZWlnaHQtLWhlYWRpbmctLWRlZmF1bHQpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X190aXRsZS0tRm9udFNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0taGVhZGluZy0tbWQpO1xcXFxuXFxcXG4gIC8qIERlc2NyaXB0aW9uICovXFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2Rlc2NyaXB0aW9uLS1Gb250U2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2Rlc2NyaXB0aW9uLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcblxcXFxuICAvKiBCb2R5ICovXFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2JvZHktLU1pbkhlaWdodDogY2FsYyh2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWRlZmF1bHQpICogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSkpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X19ib2R5LS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveF9fYm9keS0tUGFkZGluZ0lubGluZUVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveF9fYm9keS0tUGFkZGluZ0lubGluZVN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X19ib2R5LS1QYWRkaW5nQmxvY2tFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2hlYWRlci0tYm9keS0tUGFkZGluZ0Jsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG5cXFxcbiAgLyogQ2xvc2UgYnV0dG9uICovXFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2Nsb3NlLS1JbnNldEJsb2NrU3RhcnQ6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94X19oZWFkZXItLVBhZGRpbmdCbG9ja1N0YXJ0KTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveF9fY2xvc2UtLUluc2V0SW5saW5lRW5kOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9faGVhZGVyLS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcblxcXFxuICAvKiBGb290ZXIgKi9cXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveF9fZm9vdGVyLS1QYWRkaW5nQmxvY2tTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveF9fZm9vdGVyLS1QYWRkaW5nSW5saW5lRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAtLXBmLXY2LWMtbW9kYWwtYm94X19mb290ZXItLVBhZGRpbmdCbG9ja0VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgLS1wZi12Ni1jLW1vZGFsLWJveF9fZm9vdGVyLS1QYWRkaW5nSW5saW5lU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2Zvb3Rlci0tR2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxufVxcXFxuXFxcXG5AbWVkaWEgKG1pbi13aWR0aDogMTIwMHB4KSB7XFxcXG4gIDpob3N0IHtcXFxcbiAgICAtLXBmLXY2LWMtbW9kYWwtYm94LS1tLWFsaWduLXRvcC0tc3BhY2VyOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teGwpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbmRpYWxvZyB7XFxcXG4gIHBvc2l0aW9uOiBmaXhlZDtcXFxcbiAgaW5zZXQ6IDA7XFxcXG4gIHotaW5kZXg6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94LS1aSW5kZXgpO1xcXFxuICB3aWR0aDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3gtLVdpZHRoKTtcXFxcbiAgbWF4LXdpZHRoOiBtaW4odmFyKC0tcGYtdjYtYy1tb2RhbC1ib3gtLU1heFdpZHRoKSwgMTAwdncpO1xcXFxuICBtYXgtaGVpZ2h0OiBtaW4odmFyKC0tcGYtdjYtYy1tb2RhbC1ib3gtLU1heEhlaWdodCksIDEwMHZoKTtcXFxcbiAgbWFyZ2luOiBhdXRvO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBvdmVyZmxvdzogaGlkZGVuO1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbiAgYm9yZGVyOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveC0tQm9yZGVyV2lkdGgpIHNvbGlkIHZhcigtLXBmLXY2LWMtbW9kYWwtYm94LS1Cb3JkZXJDb2xvcik7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94LS1Cb3JkZXJSYWRpdXMpO1xcXFxuICBib3gtc2hhZG93OiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveC0tQm94U2hhZG93KTtcXFxcblxcXFxuICBcXFxcdTAwMjZbb3Blbl0ge1xcXFxuICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjo6YmFja2Ryb3Age1xcXFxuICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXY2LWMtYmFja2Ryb3AtLUJhY2tncm91bmRDb2xvciwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYmFja2Ryb3AtLWRlZmF1bHQpKTtcXFxcbiAgfVxcXFxuXFxcXG4gIDpob3N0KFt2YXJpYW50PVxcXFxcXFwic21hbGxcXFxcXFxcIl0pIFxcXFx1MDAyNiB7XFxcXG4gICAgLS1wZi12Ni1jLW1vZGFsLWJveC0tV2lkdGg6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94LS1tLXNtLS1NYXhXaWR0aCk7XFxcXG4gIH1cXFxcblxcXFxuICA6aG9zdChbdmFyaWFudD1cXFxcXFxcIm1lZGl1bVxcXFxcXFwiXSkgXFxcXHUwMDI2IHtcXFxcbiAgICAtLXBmLXY2LWMtbW9kYWwtYm94LS1XaWR0aDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3gtLW0tbWQtLVdpZHRoKTtcXFxcbiAgfVxcXFxuXFxcXG4gIDpob3N0KFt2YXJpYW50PVxcXFxcXFwibGFyZ2VcXFxcXFxcIl0pIFxcXFx1MDAyNiB7XFxcXG4gICAgLS1wZi12Ni1jLW1vZGFsLWJveC0tV2lkdGg6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94LS1tLWxnLS1NYXhXaWR0aCk7XFxcXG4gIH1cXFxcblxcXFxuICA6aG9zdChbcG9zaXRpb249XFxcXFxcXCJ0b3BcXFxcXFxcIl0pIFxcXFx1MDAyNiB7XFxcXG4gICAgbWFyZ2luLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveC0tbS1hbGlnbi10b3AtLUluc2V0QmxvY2tTdGFydCk7XFxcXG4gICAgbWFyZ2luLWJsb2NrLWVuZDogYXV0bztcXFxcbiAgICAtLXBmLXY2LWMtbW9kYWwtYm94LS1NYXhXaWR0aDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3gtLW0tYWxpZ24tdG9wLS1NYXhXaWR0aCk7XFxcXG4gICAgLS1wZi12Ni1jLW1vZGFsLWJveC0tTWF4SGVpZ2h0OiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveC0tbS1hbGlnbi10b3AtLU1heEhlaWdodCk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI2hlYWRlciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGZsZXgtc2hyaW5rOiAwO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9faGVhZGVyLS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9faGVhZGVyLS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3hfX2hlYWRlci0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9faGVhZGVyLS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgZ2FwOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9faGVhZGVyLS1HYXApO1xcXFxufVxcXFxuXFxcXG4jaGVhZGVyLW1haW4ge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBnYXA6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94X19oZWFkZXItbWFpbi0tR2FwKTtcXFxcbiAgcGFkZGluZy1ibG9jay1zdGFydDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3hfX2hlYWRlci1tYWluLS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIGZsZXgtZ3JvdzogMTtcXFxcbiAgbWluLXdpZHRoOiAwO1xcXFxufVxcXFxuXFxcXG46OnNsb3R0ZWQoW3Nsb3Q9XFxcXFxcXCJoZWFkZXJcXFxcXFxcIl0pIHtcXFxcbiAgbWFyZ2luOiAwICFpbXBvcnRhbnQ7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG4gIHRleHQtb3ZlcmZsb3c6IGVsbGlwc2lzO1xcXFxuICB3aGl0ZS1zcGFjZTogbm93cmFwO1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3hfX3RpdGxlLS1Gb250RmFtaWx5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9fdGl0bGUtLUZvbnRTaXplKTtcXFxcbiAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94X190aXRsZS0tRm9udFdlaWdodCk7XFxcXG4gIGxpbmUtaGVpZ2h0OiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9fdGl0bGUtLUxpbmVIZWlnaHQpO1xcXFxufVxcXFxuXFxcXG4jZGVzY3JpcHRpb24ge1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94X19kZXNjcmlwdGlvbi0tRm9udFNpemUpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3hfX2Rlc2NyaXB0aW9uLS1Db2xvcik7XFxcXG59XFxcXG5cXFxcbiNib2R5IHtcXFxcbiAgZmxleDogMSAxIGF1dG87XFxcXG4gIG1pbi1oZWlnaHQ6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94X19ib2R5LS1NaW5IZWlnaHQpO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9fYm9keS0tUGFkZGluZ0Jsb2NrU3RhcnQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1lbmQ6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94X19ib2R5LS1QYWRkaW5nSW5saW5lRW5kKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLXBmLXY2LWMtbW9kYWwtYm94X19ib2R5LS1QYWRkaW5nSW5saW5lU3RhcnQpO1xcXFxuICBwYWRkaW5nLWJsb2NrLWVuZDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3hfX2JvZHktLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIG92ZXJmbG93LXg6IGhpZGRlbjtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbiAgb3ZlcnNjcm9sbC1iZWhhdmlvcjogY29udGFpbjtcXFxcbiAgd29yZC1icmVhazogYnJlYWstd29yZDtcXFxcbiAgLXdlYmtpdC1vdmVyZmxvdy1zY3JvbGxpbmc6IHRvdWNoO1xcXFxufVxcXFxuXFxcXG4jaGVhZGVyOm5vdChbaGlkZGVuXSkgfiAjYm9keSB7XFxcXG4gIC0tcGYtdjYtYy1tb2RhbC1ib3hfX2JvZHktLVBhZGRpbmdCbG9ja1N0YXJ0OiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9faGVhZGVyLS1ib2R5LS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG59XFxcXG5cXFxcbiNmb290ZXIge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4OiAwIDAgYXV0bztcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgZ2FwOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9fZm9vdGVyLS1HYXApO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9fZm9vdGVyLS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLWVuZDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3hfX2Zvb3Rlci0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9fZm9vdGVyLS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3hfX2Zvb3Rlci0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbn1cXFxcblxcXFxuI2Nsb3NlIHtcXFxcbiAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICBpbnNldC1ibG9jay1zdGFydDogdmFyKC0tcGYtdjYtYy1tb2RhbC1ib3hfX2Nsb3NlLS1JbnNldEJsb2NrU3RhcnQpO1xcXFxuICBpbnNldC1pbmxpbmUtZW5kOiB2YXIoLS1wZi12Ni1jLW1vZGFsLWJveF9fY2xvc2UtLUluc2V0SW5saW5lRW5kKTtcXFxcblxcXFxuICBcXFxcdTAwMjYgc3ZnIHtcXFxcbiAgICB3aWR0aDogMXJlbTtcXFxcbiAgICBoZWlnaHQ6IDFyZW07XFxcXG4gICAgZmlsbDogY3VycmVudENvbG9yO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbltoaWRkZW5dIHtcXFxcbiAgZGlzcGxheTogbm9uZSAhaW1wb3J0YW50O1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLCtuUUFBMm9RLENBQUM7QUFBRSxJQUFPLHNCQUFROzs7QURNbHRRLE9BQU87QUFRQSxJQUFNLG1CQUFOLGNBQStCLE1BQU07QUFBQSxFQUMxQyxjQUFjO0FBQ1osVUFBTSxRQUFRLEVBQUUsU0FBUyxLQUFLLENBQUM7QUFBQSxFQUNqQztBQUNGO0FBS08sSUFBTSxvQkFBTixjQUFnQyxNQUFNO0FBQUEsRUFDM0M7QUFBQSxFQUNBLFlBQVksYUFBcUI7QUFDL0IsVUFBTSxTQUFTLEVBQUUsU0FBUyxLQUFLLENBQUM7QUFDaEMsU0FBSyxjQUFjO0FBQUEsRUFDckI7QUFDRjtBQUtPLElBQU0scUJBQU4sY0FBaUMsTUFBTTtBQUFBLEVBQzVDLGNBQWM7QUFDWixVQUFNLFVBQVUsRUFBRSxTQUFTLEtBQUssQ0FBQztBQUFBLEVBQ25DO0FBQ0Y7QUF0Q0E7QUE0REEseUJBQUMsY0FBYyxhQUFhO0FBQ3JCLElBQU0sWUFBTixlQUF3QixpQkFpQjdCLGdCQUFDLFNBQVMsRUFBRSxTQUFTLEtBQUssQ0FBQyxJQUczQixpQkFBQyxTQUFTLEVBQUUsU0FBUyxLQUFLLENBQUMsSUFHM0IsYUFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBdkJiLElBQVc7QUFBQSxFQUFuQztBQUFBO0FBQUE7QUFRTCxnQ0FBb0M7QUFDcEMsb0NBQWM7QUFDZCxtQ0FBYTtBQUNiLHdDQUFrQjtBQUNsQixtQ0FBYTtBQUdiO0FBQUEsdUNBQWM7QUFHZCx1QkFBUyxVQUFUO0FBR0EsdUJBQVMsV0FBVDtBQUdBLHVCQUFTLE9BQU8sa0JBQWhCLGlCQUFnQixTQUFoQjtBQXVFQSx1Q0FBaUIsTUFBTTtBQUVyQixVQUFJLEtBQUssTUFBTTtBQUNiLGFBQUssT0FBTztBQUFBLE1BQ2Q7QUFFQSxVQUFJLG1CQUFLLGNBQWE7QUFDcEIsYUFBSyxjQUFjLElBQUksbUJBQW1CLENBQUM7QUFDM0MsMkJBQUssYUFBYztBQUFBLE1BQ3JCLE9BQU87QUFDTCxhQUFLLGNBQWMsSUFBSSxrQkFBa0IsS0FBSyxXQUFXLENBQUM7QUFBQSxNQUM1RDtBQUFBLElBQ0Y7QUFFQSx3Q0FBa0IsQ0FBQyxNQUFhO0FBQzlCLFFBQUUsZUFBZTtBQUNqQix5QkFBSyxhQUFjO0FBQ25CLFdBQUssTUFBTTtBQUFBLElBQ2I7QUFFQSxzQ0FBZ0IsTUFBTTtBQUNwQix5QkFBSyxhQUFjO0FBQ25CLFdBQUssTUFBTSxRQUFRO0FBQUEsSUFDckI7QUFFQSw0Q0FBc0IsQ0FBQyxNQUFhO0FBQ2xDLFlBQU0sT0FBTyxFQUFFO0FBQ2YsWUFBTSxNQUFNLG1CQUFLO0FBQ2pCLHlCQUFLLFlBQWEsS0FBSyxjQUFjLEVBQUUsU0FBUztBQUNoRCxVQUFJLFFBQVEsbUJBQUssWUFBWSxNQUFLLGNBQWM7QUFBQSxJQUNsRDtBQUVBLGlEQUEyQixDQUFDLE1BQWE7QUFDdkMsWUFBTSxPQUFPLEVBQUU7QUFDZixZQUFNLE1BQU0sbUJBQUs7QUFDakIseUJBQUssaUJBQWtCLEtBQUssY0FBYyxFQUFFLFNBQVM7QUFDckQsVUFBSSxRQUFRLG1CQUFLLGlCQUFpQixNQUFLLGNBQWM7QUFBQSxJQUN2RDtBQUVBLDRDQUFzQixDQUFDLE1BQWE7QUFDbEMsWUFBTSxPQUFPLEVBQUU7QUFDZixZQUFNLE1BQU0sbUJBQUs7QUFDakIseUJBQUssWUFBYSxLQUFLLGNBQWMsRUFBRSxTQUFTO0FBQ2hELFVBQUksUUFBUSxtQkFBSyxZQUFZLE1BQUssY0FBYztBQUFBLElBQ2xEO0FBQUE7QUFBQSxFQWpIQSxTQUFTO0FBQ1AsV0FBTztBQUFBO0FBQUE7QUFBQSx1QkFHWSxtQkFBSyxlQUFjO0FBQUEsd0JBQ2xCLG1CQUFLLGdCQUFlO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSwrQkFLYixtQkFBSyxjQUFhO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSwwQkFRdkIsQ0FBQyxtQkFBSyxXQUFVO0FBQUE7QUFBQTtBQUFBLGdDQUdWLG1CQUFLLG9CQUFtQjtBQUFBO0FBQUE7QUFBQSwyQkFHN0IsQ0FBQyxtQkFBSyxnQkFBZTtBQUFBO0FBQUEsa0NBRWQsbUJBQUsseUJBQXdCO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsMEJBVXJDLENBQUMsbUJBQUssV0FBVTtBQUFBO0FBQUEsOEJBRVosbUJBQUssb0JBQW1CO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFJcEQ7QUFBQSxFQUVVLFFBQVEsU0FBK0I7QUFDL0MsUUFBSSxRQUFRLElBQUksTUFBTSxHQUFHO0FBQ3ZCLDRCQUFLLDBDQUFMO0FBQUEsSUFDRjtBQUFBLEVBQ0Y7QUFBQSxFQUVVLGVBQWU7QUFDdkIsdUJBQUssU0FBVSxLQUFLLFdBQVksZUFBZSxRQUFRO0FBQ3ZELDBCQUFLLDBDQUFMO0FBQUEsRUFDRjtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBaUVBLE9BQU87QUFDTCxTQUFLLFVBQVU7QUFBQSxFQUNqQjtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBS0EsWUFBWTtBQUNWLFNBQUssT0FBTztBQUFBLEVBQ2Q7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBTUEsTUFBTSxhQUFzQjtBQUMxQixRQUFJLE9BQU8sZ0JBQWdCLFVBQVU7QUFDbkMsV0FBSyxjQUFjO0FBQ25CLFVBQUksbUJBQUssVUFBUztBQUNoQiwyQkFBSyxTQUFRLGNBQWM7QUFBQSxNQUM3QjtBQUFBLElBQ0Y7QUFDQSxTQUFLLE9BQU87QUFBQSxFQUNkO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFLQSxTQUFTO0FBQ1AsUUFBSSxLQUFLLE1BQU07QUFDYixXQUFLLE1BQU07QUFBQSxJQUNiLE9BQU87QUFDTCxXQUFLLFVBQVU7QUFBQSxJQUNqQjtBQUFBLEVBQ0Y7QUFDRjtBQW5MTztBQVFMO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFNUztBQUdBO0FBR0E7QUF4Qko7QUFpRkwscUJBQWdCLFdBQUc7QUFDakIsTUFBSSxDQUFDLG1CQUFLLFNBQVM7QUFDbkIsTUFBSSxLQUFLLE1BQU07QUFDYixRQUFJLENBQUMsbUJBQUssU0FBUSxNQUFNO0FBQ3RCLHlCQUFLLFNBQVEsVUFBVTtBQUN2QixXQUFLLGNBQWMsSUFBSSxpQkFBaUIsQ0FBQztBQUFBLElBQzNDO0FBQUEsRUFDRixPQUFPO0FBQ0wsUUFBSSxtQkFBSyxTQUFRLE1BQU07QUFDckIseUJBQUssU0FBUSxNQUFNO0FBQUEsSUFDckI7QUFBQSxFQUNGO0FBQ0Y7QUFFQTtBQWNBO0FBTUE7QUFLQTtBQU9BO0FBT0E7QUFwSEEsNEJBQVMsV0FEVCxjQWpCVyxXQWtCRjtBQUdULDRCQUFTLFlBRFQsZUFwQlcsV0FxQkY7QUFHVCw0QkFBUyxRQURULFdBdkJXLFdBd0JGO0FBeEJFLFlBQU4seUNBRFAsdUJBQ2E7QUFDWCxjQURXLFdBQ0oscUJBQW9DO0FBQUEsRUFDekMsR0FBRyxXQUFXO0FBQUEsRUFDZCxnQkFBZ0I7QUFDbEI7QUFFQSxjQU5XLFdBTUosVUFBUztBQU5YLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
