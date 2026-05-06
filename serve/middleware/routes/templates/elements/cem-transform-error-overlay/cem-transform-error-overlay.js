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

// elements/cem-transform-error-overlay/cem-transform-error-overlay.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";
import "../cem-pf-v6-button/cem-pf-v6-button.js";

// lit-css:elements/cem-transform-error-overlay/cem-transform-error-overlay.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('"/* Transform Error Overlay - displays server-side compilation errors */\\n\\n:host {\\n  display: none;\\n  position: fixed;\\n  inset: 0;\\n  z-index: var(--pf-t--global--z-index--2xl, 1000);\\n  background: rgba(0, 0, 0, 0.9);\\n  backdrop-filter: blur(4px);\\n  animation: fadeIn 0.2s ease-out;\\n}\\n\\n:host([open]) {\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n  padding: var(--pf-t--global--spacer--lg);\\n}\\n\\n@keyframes fadeIn {\\n  from {\\n    opacity: 0;\\n  }\\n  to {\\n    opacity: 1;\\n  }\\n}\\n\\n#overlay-content {\\n  background: var(--pf-t--global--background--color--floating--default);\\n  color: var(--pf-t--global--text--color--regular);\\n  border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--color--status--danger--default);\\n  border-radius: var(--pf-t--global--border--radius--medium);\\n  max-width: 800px;\\n  width: 100%;\\n  max-height: 90vh;\\n  display: flex;\\n  flex-direction: column;\\n  box-shadow: var(--pf-t--global--box-shadow--xl);\\n  font-family: var(--pf-t--global--font--family--mono);\\n}\\n\\n#header {\\n  display: flex;\\n  align-items: center;\\n  justify-content: space-between;\\n  padding: var(--pf-t--global--spacer--md) var(--pf-t--global--spacer--lg);\\n  background: var(--pf-t--global--color--status--danger--default);\\n  color: var(--pf-t--global--text--color--on-status--danger);\\n  border-radius: var(--pf-t--global--border--radius--small) var(--pf-t--global--border--radius--small) 0 0;\\n}\\n\\n#title-container {\\n  display: flex;\\n  align-items: center;\\n  gap: var(--pf-t--global--spacer--sm);\\n  font-size: var(--pf-t--global--font--size--body--lg);\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  margin: 0;\\n}\\n\\n#error-icon {\\n  font-size: var(--pf-t--global--font--size--heading--sm);\\n}\\n\\n#close {\\n  --cem-pf-v6-c-button--Color: var(--pf-t--global--text--color--on-status--danger);\\n  --cem-pf-v6-c-button--BackgroundColor: rgba(255, 255, 255, 0.2);\\n  --cem-pf-v6-c-button--BorderColor: transparent;\\n  --cem-pf-v6-c-button--hover--BackgroundColor: rgba(255, 255, 255, 0.3);\\n  --cem-pf-v6-c-button--hover--Color: var(--pf-t--global--text--color--on-status--danger);\\n}\\n\\n#body {\\n  padding: var(--pf-t--global--spacer--lg);\\n  overflow-y: auto;\\n  flex: 1;\\n}\\n\\n#file {\\n  background: var(--pf-t--global--background--color--secondary--default);\\n  padding: var(--pf-t--global--spacer--sm) var(--pf-t--global--spacer--md);\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  margin-bottom: var(--pf-t--global--spacer--md);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  color: var(--pf-t--global--color--brand--default);\\n  border-inline-start: var(--pf-t--global--border--width--strong, 3px) solid var(--pf-t--global--color--status--danger--default);\\n}\\n\\n#file:empty {\\n  display: none;\\n}\\n\\n#message {\\n  background: var(--pf-t--global--background--color--primary--default);\\n  padding: var(--pf-t--global--spacer--md);\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  white-space: pre-wrap;\\n  word-break: break-word;\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--regular);\\n  border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n}\\n\\n#footer {\\n  padding: var(--pf-t--global--spacer--md) var(--pf-t--global--spacer--lg);\\n  background: var(--pf-t--global--background--color--secondary--default);\\n  border-block-start: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n  border-radius: 0 0 var(--pf-t--global--border--radius--small) var(--pf-t--global--border--radius--small);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  color: var(--pf-t--global--text--color--subtle);\\n}\\n"'));
var cem_transform_error_overlay_default = s;

// elements/cem-transform-error-overlay/cem-transform-error-overlay.ts
var _message_dec, _file_dec, _title_dec, _open_dec, _a, _CemTransformErrorOverlay_decorators, _init, _open, _title, _file, _message, _handleKeydown;
_CemTransformErrorOverlay_decorators = [customElement("cem-transform-error-overlay")];
var CemTransformErrorOverlay = class extends (_a = LitElement, _open_dec = [property({ type: Boolean, reflect: true })], _title_dec = [property()], _file_dec = [property()], _message_dec = [property()], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _open, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _title, __runInitializers(_init, 12, this, "")), __runInitializers(_init, 15, this);
    __privateAdd(this, _file, __runInitializers(_init, 16, this, "")), __runInitializers(_init, 19, this);
    __privateAdd(this, _message, __runInitializers(_init, 20, this, "")), __runInitializers(_init, 23, this);
    __privateAdd(this, _handleKeydown, (e) => {
      if (e.key === "Escape" && this.open) {
        this.hide();
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    document.addEventListener("keydown", __privateGet(this, _handleKeydown));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    document.removeEventListener("keydown", __privateGet(this, _handleKeydown));
  }
  render() {
    return html`
      <div id="overlay-content">
        <div id="header">
          <h2 id="title-container">
            <span id="error-icon">\u26A0\uFE0F</span>
            <span>${this.title}</span>
          </h2>
          <cem-pf-v6-button id="close"
                        variant="plain"
                        @click=${this.hide}>Dismiss</cem-pf-v6-button>
        </div>
        <div id="body">
          <div id="file">${this.file ? `File: ${this.file}` : ""}</div>
          <div id="message">${this.message}</div>
        </div>
        <div id="footer">
          This error will automatically dismiss when the issue is fixed.
        </div>
      </div>
    `;
  }
  /**
   * Show the error overlay.
   * @param title - Error title
   * @param message - Error message
   * @param file - Optional file path where error occurred
   */
  show(title, message, file = "") {
    this.title = title;
    this.message = message;
    this.file = file;
    this.open = true;
  }
  /** Hide the error overlay */
  hide() {
    this.open = false;
  }
};
_init = __decoratorStart(_a);
_open = new WeakMap();
_title = new WeakMap();
_file = new WeakMap();
_message = new WeakMap();
_handleKeydown = new WeakMap();
__decorateElement(_init, 4, "open", _open_dec, CemTransformErrorOverlay, _open);
__decorateElement(_init, 4, "title", _title_dec, CemTransformErrorOverlay, _title);
__decorateElement(_init, 4, "file", _file_dec, CemTransformErrorOverlay, _file);
__decorateElement(_init, 4, "message", _message_dec, CemTransformErrorOverlay, _message);
CemTransformErrorOverlay = __decorateElement(_init, 0, "CemTransformErrorOverlay", _CemTransformErrorOverlay_decorators, CemTransformErrorOverlay);
__publicField(CemTransformErrorOverlay, "styles", cem_transform_error_overlay_default);
__runInitializers(_init, 1, CemTransformErrorOverlay);
export {
  CemTransformErrorOverlay
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5L2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheS50cyIsICJsaXQtY3NzOmVsZW1lbnRzL2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheS9jZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXkuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0ICcuLi9jZW0tcGYtdjYtYnV0dG9uL2NlbS1wZi12Ni1idXR0b24uanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5LmNzcyc7XG5cbi8qKlxuICogVHJhbnNmb3JtIEVycm9yIE92ZXJsYXkgQ29tcG9uZW50LlxuICogRGlzcGxheXMgVHlwZVNjcmlwdCB0cmFuc2Zvcm0gZXJyb3JzIGFuZCBvdGhlciBzZXJ2ZXItc2lkZSBjb21waWxhdGlvbiBlcnJvcnMuXG4gKlxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5XG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXknKVxuZXhwb3J0IGNsYXNzIENlbVRyYW5zZm9ybUVycm9yT3ZlcmxheSBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3Igb3BlbiA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSgpXG4gIGFjY2Vzc29yIHRpdGxlID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3IgZmlsZSA9ICcnO1xuXG4gIEBwcm9wZXJ0eSgpXG4gIGFjY2Vzc29yIG1lc3NhZ2UgPSAnJztcblxuICAjaGFuZGxlS2V5ZG93biA9IChlOiBLZXlib2FyZEV2ZW50KSA9PiB7XG4gICAgaWYgKGUua2V5ID09PSAnRXNjYXBlJyAmJiB0aGlzLm9wZW4pIHtcbiAgICAgIHRoaXMuaGlkZSgpO1xuICAgIH1cbiAgfTtcblxuICBjb25uZWN0ZWRDYWxsYmFjaygpOiB2b2lkIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIGRvY3VtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNoYW5kbGVLZXlkb3duKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCk6IHZvaWQge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgZG9jdW1lbnQucmVtb3ZlRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI2hhbmRsZUtleWRvd24pO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cIm92ZXJsYXktY29udGVudFwiPlxuICAgICAgICA8ZGl2IGlkPVwiaGVhZGVyXCI+XG4gICAgICAgICAgPGgyIGlkPVwidGl0bGUtY29udGFpbmVyXCI+XG4gICAgICAgICAgICA8c3BhbiBpZD1cImVycm9yLWljb25cIj5cXHUyNkEwXFx1RkUwRjwvc3Bhbj5cbiAgICAgICAgICAgIDxzcGFuPiR7dGhpcy50aXRsZX08L3NwYW4+XG4gICAgICAgICAgPC9oMj5cbiAgICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBpZD1cImNsb3NlXCJcbiAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLmhpZGV9PkRpc21pc3M8L2NlbS1wZi12Ni1idXR0b24+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGlkPVwiYm9keVwiPlxuICAgICAgICAgIDxkaXYgaWQ9XCJmaWxlXCI+JHt0aGlzLmZpbGUgPyBgRmlsZTogJHt0aGlzLmZpbGV9YCA6ICcnfTwvZGl2PlxuICAgICAgICAgIDxkaXYgaWQ9XCJtZXNzYWdlXCI+JHt0aGlzLm1lc3NhZ2V9PC9kaXY+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGlkPVwiZm9vdGVyXCI+XG4gICAgICAgICAgVGhpcyBlcnJvciB3aWxsIGF1dG9tYXRpY2FsbHkgZGlzbWlzcyB3aGVuIHRoZSBpc3N1ZSBpcyBmaXhlZC5cbiAgICAgICAgPC9kaXY+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG5cbiAgLyoqXG4gICAqIFNob3cgdGhlIGVycm9yIG92ZXJsYXkuXG4gICAqIEBwYXJhbSB0aXRsZSAtIEVycm9yIHRpdGxlXG4gICAqIEBwYXJhbSBtZXNzYWdlIC0gRXJyb3IgbWVzc2FnZVxuICAgKiBAcGFyYW0gZmlsZSAtIE9wdGlvbmFsIGZpbGUgcGF0aCB3aGVyZSBlcnJvciBvY2N1cnJlZFxuICAgKi9cbiAgc2hvdyh0aXRsZTogc3RyaW5nLCBtZXNzYWdlOiBzdHJpbmcsIGZpbGUgPSAnJyk6IHZvaWQge1xuICAgIHRoaXMudGl0bGUgPSB0aXRsZTtcbiAgICB0aGlzLm1lc3NhZ2UgPSBtZXNzYWdlO1xuICAgIHRoaXMuZmlsZSA9IGZpbGU7XG4gICAgdGhpcy5vcGVuID0gdHJ1ZTtcbiAgfVxuXG4gIC8qKiBIaWRlIHRoZSBlcnJvciBvdmVybGF5ICovXG4gIGhpZGUoKTogdm9pZCB7XG4gICAgdGhpcy5vcGVuID0gZmFsc2U7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5JzogQ2VtVHJhbnNmb3JtRXJyb3JPdmVybGF5O1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCIvKiBUcmFuc2Zvcm0gRXJyb3IgT3ZlcmxheSAtIGRpc3BsYXlzIHNlcnZlci1zaWRlIGNvbXBpbGF0aW9uIGVycm9ycyAqL1xcXFxuXFxcXG46aG9zdCB7XFxcXG4gIGRpc3BsYXk6IG5vbmU7XFxcXG4gIHBvc2l0aW9uOiBmaXhlZDtcXFxcbiAgaW5zZXQ6IDA7XFxcXG4gIHotaW5kZXg6IHZhcigtLXBmLXQtLWdsb2JhbC0tei1pbmRleC0tMnhsLCAxMDAwKTtcXFxcbiAgYmFja2dyb3VuZDogcmdiYSgwLCAwLCAwLCAwLjkpO1xcXFxuICBiYWNrZHJvcC1maWx0ZXI6IGJsdXIoNHB4KTtcXFxcbiAgYW5pbWF0aW9uOiBmYWRlSW4gMC4ycyBlYXNlLW91dDtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW29wZW5dKSB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGp1c3RpZnktY29udGVudDogY2VudGVyO1xcXFxuICBwYWRkaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxufVxcXFxuXFxcXG5Aa2V5ZnJhbWVzIGZhZGVJbiB7XFxcXG4gIGZyb20ge1xcXFxuICAgIG9wYWNpdHk6IDA7XFxcXG4gIH1cXFxcbiAgdG8ge1xcXFxuICAgIG9wYWNpdHk6IDE7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI292ZXJsYXktY29udGVudCB7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWZsb2F0aW5nLS1kZWZhdWx0KTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICBib3JkZXI6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tcmVndWxhcikgc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tc3RhdHVzLS1kYW5nZXItLWRlZmF1bHQpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1tZWRpdW0pO1xcXFxuICBtYXgtd2lkdGg6IDgwMHB4O1xcXFxuICB3aWR0aDogMTAwJTtcXFxcbiAgbWF4LWhlaWdodDogOTB2aDtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgYm94LXNoYWRvdzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3gtc2hhZG93LS14bCk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLWZhbWlseS0tbW9ubyk7XFxcXG59XFxcXG5cXFxcbiNoZWFkZXIge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBqdXN0aWZ5LWNvbnRlbnQ6IHNwYWNlLWJldHdlZW47XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tc3RhdHVzLS1kYW5nZXItLWRlZmF1bHQpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tc3RhdHVzLS1kYW5nZXIpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpIDAgMDtcXFxcbn1cXFxcblxcXFxuI3RpdGxlLWNvbnRhaW5lciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLWxnKTtcXFxcbiAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0td2VpZ2h0LS1ib2R5LS1ib2xkKTtcXFxcbiAgbWFyZ2luOiAwO1xcXFxufVxcXFxuXFxcXG4jZXJyb3ItaWNvbiB7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1oZWFkaW5nLS1zbSk7XFxcXG59XFxcXG5cXFxcbiNjbG9zZSB7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tc3RhdHVzLS1kYW5nZXIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLWJ1dHRvbi0tQmFja2dyb3VuZENvbG9yOiByZ2JhKDI1NSwgMjU1LCAyNTUsIDAuMik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1Cb3JkZXJDb2xvcjogdHJhbnNwYXJlbnQ7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1ob3Zlci0tQmFja2dyb3VuZENvbG9yOiByZ2JhKDI1NSwgMjU1LCAyNTUsIDAuMyk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1ob3Zlci0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLXN0YXR1cy0tZGFuZ2VyKTtcXFxcbn1cXFxcblxcXFxuI2JvZHkge1xcXFxuICBwYWRkaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuICBmbGV4OiAxO1xcXFxufVxcXFxuXFxcXG4jZmlsZSB7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0tZGVmYXVsdCk7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQpO1xcXFxuICBib3JkZXItaW5saW5lLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLXN0cm9uZywgM3B4KSBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1zdGF0dXMtLWRhbmdlci0tZGVmYXVsdCk7XFxcXG59XFxcXG5cXFxcbiNmaWxlOmVtcHR5IHtcXFxcbiAgZGlzcGxheTogbm9uZTtcXFxcbn1cXFxcblxcXFxuI21lc3NhZ2Uge1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1kZWZhdWx0KTtcXFxcbiAgcGFkZGluZzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICB3aGl0ZS1zcGFjZTogcHJlLXdyYXA7XFxcXG4gIHdvcmQtYnJlYWs6IGJyZWFrLXdvcmQ7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLWxpbmUtaGVpZ2h0LS1ib2R5KTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICBib3JkZXI6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tcmVndWxhcikgc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbn1cXFxcblxcXFxuI2Zvb3RlciB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLXJlZ3VsYXIpIHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IDAgMCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7QUFFekIsT0FBTzs7O0FDSlAsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0seTRIQUEyNEgsQ0FBQztBQUFFLElBQU8sc0NBQVE7OztBREFsOUg7QUFjQSx3Q0FBQyxjQUFjLDZCQUE2QjtBQUNyQyxJQUFNLDJCQUFOLGVBQXVDLGlCQUc1QyxhQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFHMUMsY0FBQyxTQUFTLElBR1YsYUFBQyxTQUFTLElBR1YsZ0JBQUMsU0FBUyxJQVprQyxJQUFXO0FBQUEsRUFBbEQ7QUFBQTtBQUlMLHVCQUFTLE9BQU8sa0JBQWhCLGdCQUFnQixTQUFoQjtBQUdBLHVCQUFTLFFBQVEsa0JBQWpCLGlCQUFpQixNQUFqQjtBQUdBLHVCQUFTLE9BQU8sa0JBQWhCLGlCQUFnQixNQUFoQjtBQUdBLHVCQUFTLFVBQVUsa0JBQW5CLGlCQUFtQixNQUFuQjtBQUVBLHVDQUFpQixDQUFDLE1BQXFCO0FBQ3JDLFVBQUksRUFBRSxRQUFRLFlBQVksS0FBSyxNQUFNO0FBQ25DLGFBQUssS0FBSztBQUFBLE1BQ1o7QUFBQSxJQUNGO0FBQUE7QUFBQSxFQUVBLG9CQUEwQjtBQUN4QixVQUFNLGtCQUFrQjtBQUN4QixhQUFTLGlCQUFpQixXQUFXLG1CQUFLLGVBQWM7QUFBQSxFQUMxRDtBQUFBLEVBRUEsdUJBQTZCO0FBQzNCLFVBQU0scUJBQXFCO0FBQzNCLGFBQVMsb0JBQW9CLFdBQVcsbUJBQUssZUFBYztBQUFBLEVBQzdEO0FBQUEsRUFFQSxTQUFTO0FBQ1AsV0FBTztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsb0JBS1MsS0FBSyxLQUFLO0FBQUE7QUFBQTtBQUFBO0FBQUEsaUNBSUcsS0FBSyxJQUFJO0FBQUE7QUFBQTtBQUFBLDJCQUdmLEtBQUssT0FBTyxTQUFTLEtBQUssSUFBSSxLQUFLLEVBQUU7QUFBQSw4QkFDbEMsS0FBSyxPQUFPO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFPeEM7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQVFBLEtBQUssT0FBZSxTQUFpQixPQUFPLElBQVU7QUFDcEQsU0FBSyxRQUFRO0FBQ2IsU0FBSyxVQUFVO0FBQ2YsU0FBSyxPQUFPO0FBQ1osU0FBSyxPQUFPO0FBQUEsRUFDZDtBQUFBO0FBQUEsRUFHQSxPQUFhO0FBQ1gsU0FBSyxPQUFPO0FBQUEsRUFDZDtBQUNGO0FBdkVPO0FBSUk7QUFHQTtBQUdBO0FBR0E7QUFFVDtBQVhBLDRCQUFTLFFBRFQsV0FIVywwQkFJRjtBQUdULDRCQUFTLFNBRFQsWUFOVywwQkFPRjtBQUdULDRCQUFTLFFBRFQsV0FUVywwQkFVRjtBQUdULDRCQUFTLFdBRFQsY0FaVywwQkFhRjtBQWJFLDJCQUFOLHdEQURQLHNDQUNhO0FBQ1gsY0FEVywwQkFDSixVQUFTO0FBRFgsNEJBQU07IiwKICAibmFtZXMiOiBbXQp9Cg==
