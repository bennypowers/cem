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
import "../pf-v6-button/pf-v6-button.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-transform-error-overlay/cem-transform-error-overlay.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('"/* Transform Error Overlay - displays server-side compilation errors */\\n\\n:host {\\n  display: none;\\n  position: fixed;\\n  inset: 0;\\n  z-index: var(--pf-t--global--z-index--2xl, 1000);\\n  background: rgba(0, 0, 0, 0.9);\\n  backdrop-filter: blur(4px);\\n  animation: fadeIn 0.2s ease-out;\\n}\\n\\n:host([open]) {\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n  padding: var(--pf-t--global--spacer--lg);\\n}\\n\\n@keyframes fadeIn {\\n  from {\\n    opacity: 0;\\n  }\\n  to {\\n    opacity: 1;\\n  }\\n}\\n\\n#overlay-content {\\n  background: var(--pf-t--global--background--color--floating--default);\\n  color: var(--pf-t--global--text--color--regular);\\n  border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--color--status--danger--default);\\n  border-radius: var(--pf-t--global--border--radius--medium);\\n  max-width: 800px;\\n  width: 100%;\\n  max-height: 90vh;\\n  display: flex;\\n  flex-direction: column;\\n  box-shadow: var(--pf-t--global--box-shadow--xl);\\n  font-family: var(--pf-t--global--font--family--mono);\\n}\\n\\n#header {\\n  display: flex;\\n  align-items: center;\\n  justify-content: space-between;\\n  padding: var(--pf-t--global--spacer--md) var(--pf-t--global--spacer--lg);\\n  background: var(--pf-t--global--color--status--danger--default);\\n  color: var(--pf-t--global--text--color--on-status--danger);\\n  border-radius: var(--pf-t--global--border--radius--small) var(--pf-t--global--border--radius--small) 0 0;\\n}\\n\\n#title-container {\\n  display: flex;\\n  align-items: center;\\n  gap: var(--pf-t--global--spacer--sm);\\n  font-size: var(--pf-t--global--font--size--body--lg);\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  margin: 0;\\n}\\n\\n#error-icon {\\n  font-size: var(--pf-t--global--font--size--heading--sm);\\n}\\n\\n#close {\\n  --pf-v6-c-button--Color: var(--pf-t--global--text--color--on-status--danger);\\n  --pf-v6-c-button--BackgroundColor: rgba(255, 255, 255, 0.2);\\n  --pf-v6-c-button--BorderColor: transparent;\\n  --pf-v6-c-button--hover--BackgroundColor: rgba(255, 255, 255, 0.3);\\n  --pf-v6-c-button--hover--Color: var(--pf-t--global--text--color--on-status--danger);\\n}\\n\\n#body {\\n  padding: var(--pf-t--global--spacer--lg);\\n  overflow-y: auto;\\n  flex: 1;\\n}\\n\\n#file {\\n  background: var(--pf-t--global--background--color--secondary--default);\\n  padding: var(--pf-t--global--spacer--sm) var(--pf-t--global--spacer--md);\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  margin-bottom: var(--pf-t--global--spacer--md);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  color: var(--pf-t--global--color--brand--default);\\n  border-inline-start: var(--pf-t--global--border--width--strong, 3px) solid var(--pf-t--global--color--status--danger--default);\\n}\\n\\n#file:empty {\\n  display: none;\\n}\\n\\n#message {\\n  background: var(--pf-t--global--background--color--primary--default);\\n  padding: var(--pf-t--global--spacer--md);\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  white-space: pre-wrap;\\n  word-break: break-word;\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--regular);\\n  border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n}\\n\\n#footer {\\n  padding: var(--pf-t--global--spacer--md) var(--pf-t--global--spacer--lg);\\n  background: var(--pf-t--global--background--color--secondary--default);\\n  border-block-start: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n  border-radius: 0 0 var(--pf-t--global--border--radius--small) var(--pf-t--global--border--radius--small);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  color: var(--pf-t--global--text--color--subtle);\\n}\\n"'));
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
          <pf-v6-button id="close"
                        variant="plain"
                        @click=${this.hide}>Dismiss</pf-v6-button>
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5L2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheS50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheS9jZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXkuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheS5jc3MnO1xuXG4vKipcbiAqIFRyYW5zZm9ybSBFcnJvciBPdmVybGF5IENvbXBvbmVudC5cbiAqIERpc3BsYXlzIFR5cGVTY3JpcHQgdHJhbnNmb3JtIGVycm9ycyBhbmQgb3RoZXIgc2VydmVyLXNpZGUgY29tcGlsYXRpb24gZXJyb3JzLlxuICpcbiAqIEBjdXN0b21FbGVtZW50IGNlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheVxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5JylcbmV4cG9ydCBjbGFzcyBDZW1UcmFuc2Zvcm1FcnJvck92ZXJsYXkgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIG9wZW4gPSBmYWxzZTtcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciB0aXRsZSA9ICcnO1xuXG4gIEBwcm9wZXJ0eSgpXG4gIGFjY2Vzc29yIGZpbGUgPSAnJztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBtZXNzYWdlID0gJyc7XG5cbiAgI2hhbmRsZUtleWRvd24gPSAoZTogS2V5Ym9hcmRFdmVudCkgPT4ge1xuICAgIGlmIChlLmtleSA9PT0gJ0VzY2FwZScgJiYgdGhpcy5vcGVuKSB7XG4gICAgICB0aGlzLmhpZGUoKTtcbiAgICB9XG4gIH07XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKTogdm9pZCB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICBkb2N1bWVudC5hZGRFdmVudExpc3RlbmVyKCdrZXlkb3duJywgdGhpcy4jaGFuZGxlS2V5ZG93bik7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpOiB2b2lkIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIGRvY3VtZW50LnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNoYW5kbGVLZXlkb3duKTtcbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJvdmVybGF5LWNvbnRlbnRcIj5cbiAgICAgICAgPGRpdiBpZD1cImhlYWRlclwiPlxuICAgICAgICAgIDxoMiBpZD1cInRpdGxlLWNvbnRhaW5lclwiPlxuICAgICAgICAgICAgPHNwYW4gaWQ9XCJlcnJvci1pY29uXCI+XFx1MjZBMFxcdUZFMEY8L3NwYW4+XG4gICAgICAgICAgICA8c3Bhbj4ke3RoaXMudGl0bGV9PC9zcGFuPlxuICAgICAgICAgIDwvaDI+XG4gICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImNsb3NlXCJcbiAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLmhpZGV9PkRpc21pc3M8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgaWQ9XCJib2R5XCI+XG4gICAgICAgICAgPGRpdiBpZD1cImZpbGVcIj4ke3RoaXMuZmlsZSA/IGBGaWxlOiAke3RoaXMuZmlsZX1gIDogJyd9PC9kaXY+XG4gICAgICAgICAgPGRpdiBpZD1cIm1lc3NhZ2VcIj4ke3RoaXMubWVzc2FnZX08L2Rpdj5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgaWQ9XCJmb290ZXJcIj5cbiAgICAgICAgICBUaGlzIGVycm9yIHdpbGwgYXV0b21hdGljYWxseSBkaXNtaXNzIHdoZW4gdGhlIGlzc3VlIGlzIGZpeGVkLlxuICAgICAgICA8L2Rpdj5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cblxuICAvKipcbiAgICogU2hvdyB0aGUgZXJyb3Igb3ZlcmxheS5cbiAgICogQHBhcmFtIHRpdGxlIC0gRXJyb3IgdGl0bGVcbiAgICogQHBhcmFtIG1lc3NhZ2UgLSBFcnJvciBtZXNzYWdlXG4gICAqIEBwYXJhbSBmaWxlIC0gT3B0aW9uYWwgZmlsZSBwYXRoIHdoZXJlIGVycm9yIG9jY3VycmVkXG4gICAqL1xuICBzaG93KHRpdGxlOiBzdHJpbmcsIG1lc3NhZ2U6IHN0cmluZywgZmlsZSA9ICcnKTogdm9pZCB7XG4gICAgdGhpcy50aXRsZSA9IHRpdGxlO1xuICAgIHRoaXMubWVzc2FnZSA9IG1lc3NhZ2U7XG4gICAgdGhpcy5maWxlID0gZmlsZTtcbiAgICB0aGlzLm9wZW4gPSB0cnVlO1xuICB9XG5cbiAgLyoqIEhpZGUgdGhlIGVycm9yIG92ZXJsYXkgKi9cbiAgaGlkZSgpOiB2b2lkIHtcbiAgICB0aGlzLm9wZW4gPSBmYWxzZTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXknOiBDZW1UcmFuc2Zvcm1FcnJvck92ZXJsYXk7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIi8qIFRyYW5zZm9ybSBFcnJvciBPdmVybGF5IC0gZGlzcGxheXMgc2VydmVyLXNpZGUgY29tcGlsYXRpb24gZXJyb3JzICovXFxcXG5cXFxcbjpob3N0IHtcXFxcbiAgZGlzcGxheTogbm9uZTtcXFxcbiAgcG9zaXRpb246IGZpeGVkO1xcXFxuICBpbnNldDogMDtcXFxcbiAgei1pbmRleDogdmFyKC0tcGYtdC0tZ2xvYmFsLS16LWluZGV4LS0yeGwsIDEwMDApO1xcXFxuICBiYWNrZ3JvdW5kOiByZ2JhKDAsIDAsIDAsIDAuOSk7XFxcXG4gIGJhY2tkcm9wLWZpbHRlcjogYmx1cig0cHgpO1xcXFxuICBhbmltYXRpb246IGZhZGVJbiAwLjJzIGVhc2Utb3V0O1xcXFxufVxcXFxuXFxcXG46aG9zdChbb3Blbl0pIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAganVzdGlmeS1jb250ZW50OiBjZW50ZXI7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG59XFxcXG5cXFxcbkBrZXlmcmFtZXMgZmFkZUluIHtcXFxcbiAgZnJvbSB7XFxcXG4gICAgb3BhY2l0eTogMDtcXFxcbiAgfVxcXFxuICB0byB7XFxcXG4gICAgb3BhY2l0eTogMTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jb3ZlcmxheS1jb250ZW50IHtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tZmxvYXRpbmctLWRlZmF1bHQpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIGJvcmRlcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1yZWd1bGFyKSBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1zdGF0dXMtLWRhbmdlci0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLW1lZGl1bSk7XFxcXG4gIG1heC13aWR0aDogODAwcHg7XFxcXG4gIHdpZHRoOiAxMDAlO1xcXFxuICBtYXgtaGVpZ2h0OiA5MHZoO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBib3gtc2hhZG93OiB2YXIoLS1wZi10LS1nbG9iYWwtLWJveC1zaGFkb3ctLXhsKTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tZmFtaWx5LS1tb25vKTtcXFxcbn1cXFxcblxcXFxuI2hlYWRlciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGp1c3RpZnktY29udGVudDogc3BhY2UtYmV0d2VlbjtcXFxcbiAgcGFkZGluZzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKSB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1zdGF0dXMtLWRhbmdlci0tZGVmYXVsdCk7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1zdGF0dXMtLWRhbmdlcik7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXNtYWxsKSB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCkgMCAwO1xcXFxufVxcXFxuXFxcXG4jdGl0bGUtY29udGFpbmVyIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgZ2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tbGcpO1xcXFxuICBmb250LXdlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS13ZWlnaHQtLWJvZHktLWJvbGQpO1xcXFxuICBtYXJnaW46IDA7XFxcXG59XFxcXG5cXFxcbiNlcnJvci1pY29uIHtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWhlYWRpbmctLXNtKTtcXFxcbn1cXFxcblxcXFxuI2Nsb3NlIHtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLXN0YXR1cy0tZGFuZ2VyKTtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tQmFja2dyb3VuZENvbG9yOiByZ2JhKDI1NSwgMjU1LCAyNTUsIDAuMik7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLUJvcmRlckNvbG9yOiB0cmFuc3BhcmVudDtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0taG92ZXItLUJhY2tncm91bmRDb2xvcjogcmdiYSgyNTUsIDI1NSwgMjU1LCAwLjMpO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1ob3Zlci0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLXN0YXR1cy0tZGFuZ2VyKTtcXFxcbn1cXFxcblxcXFxuI2JvZHkge1xcXFxuICBwYWRkaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuICBmbGV4OiAxO1xcXFxufVxcXFxuXFxcXG4jZmlsZSB7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0tZGVmYXVsdCk7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQpO1xcXFxuICBib3JkZXItaW5saW5lLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLXN0cm9uZywgM3B4KSBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1zdGF0dXMtLWRhbmdlci0tZGVmYXVsdCk7XFxcXG59XFxcXG5cXFxcbiNmaWxlOmVtcHR5IHtcXFxcbiAgZGlzcGxheTogbm9uZTtcXFxcbn1cXFxcblxcXFxuI21lc3NhZ2Uge1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1kZWZhdWx0KTtcXFxcbiAgcGFkZGluZzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICB3aGl0ZS1zcGFjZTogcHJlLXdyYXA7XFxcXG4gIHdvcmQtYnJlYWs6IGJyZWFrLXdvcmQ7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLWxpbmUtaGVpZ2h0LS1ib2R5KTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICBib3JkZXI6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tcmVndWxhcikgc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbn1cXFxcblxcXFxuI2Zvb3RlciB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLXJlZ3VsYXIpIHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IDAgMCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7QUFFekIsT0FBTzs7O0FDSlAsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0scTNIQUF1M0gsQ0FBQztBQUFFLElBQU8sc0NBQVE7OztBREE5N0g7QUFjQSx3Q0FBQyxjQUFjLDZCQUE2QjtBQUNyQyxJQUFNLDJCQUFOLGVBQXVDLGlCQUc1QyxhQUFDLFNBQVMsRUFBRSxNQUFNLFNBQVMsU0FBUyxLQUFLLENBQUMsSUFHMUMsY0FBQyxTQUFTLElBR1YsYUFBQyxTQUFTLElBR1YsZ0JBQUMsU0FBUyxJQVprQyxJQUFXO0FBQUEsRUFBbEQ7QUFBQTtBQUlMLHVCQUFTLE9BQU8sa0JBQWhCLGdCQUFnQixTQUFoQjtBQUdBLHVCQUFTLFFBQVEsa0JBQWpCLGlCQUFpQixNQUFqQjtBQUdBLHVCQUFTLE9BQU8sa0JBQWhCLGlCQUFnQixNQUFoQjtBQUdBLHVCQUFTLFVBQVUsa0JBQW5CLGlCQUFtQixNQUFuQjtBQUVBLHVDQUFpQixDQUFDLE1BQXFCO0FBQ3JDLFVBQUksRUFBRSxRQUFRLFlBQVksS0FBSyxNQUFNO0FBQ25DLGFBQUssS0FBSztBQUFBLE1BQ1o7QUFBQSxJQUNGO0FBQUE7QUFBQSxFQUVBLG9CQUEwQjtBQUN4QixVQUFNLGtCQUFrQjtBQUN4QixhQUFTLGlCQUFpQixXQUFXLG1CQUFLLGVBQWM7QUFBQSxFQUMxRDtBQUFBLEVBRUEsdUJBQTZCO0FBQzNCLFVBQU0scUJBQXFCO0FBQzNCLGFBQVMsb0JBQW9CLFdBQVcsbUJBQUssZUFBYztBQUFBLEVBQzdEO0FBQUEsRUFFQSxTQUFTO0FBQ1AsV0FBTztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsb0JBS1MsS0FBSyxLQUFLO0FBQUE7QUFBQTtBQUFBO0FBQUEsaUNBSUcsS0FBSyxJQUFJO0FBQUE7QUFBQTtBQUFBLDJCQUdmLEtBQUssT0FBTyxTQUFTLEtBQUssSUFBSSxLQUFLLEVBQUU7QUFBQSw4QkFDbEMsS0FBSyxPQUFPO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFPeEM7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQVFBLEtBQUssT0FBZSxTQUFpQixPQUFPLElBQVU7QUFDcEQsU0FBSyxRQUFRO0FBQ2IsU0FBSyxVQUFVO0FBQ2YsU0FBSyxPQUFPO0FBQ1osU0FBSyxPQUFPO0FBQUEsRUFDZDtBQUFBO0FBQUEsRUFHQSxPQUFhO0FBQ1gsU0FBSyxPQUFPO0FBQUEsRUFDZDtBQUNGO0FBdkVPO0FBSUk7QUFHQTtBQUdBO0FBR0E7QUFFVDtBQVhBLDRCQUFTLFFBRFQsV0FIVywwQkFJRjtBQUdULDRCQUFTLFNBRFQsWUFOVywwQkFPRjtBQUdULDRCQUFTLFFBRFQsV0FUVywwQkFVRjtBQUdULDRCQUFTLFdBRFQsY0FaVywwQkFhRjtBQWJFLDJCQUFOLHdEQURQLHNDQUNhO0FBQ1gsY0FEVywwQkFDSixVQUFTO0FBRFgsNEJBQU07IiwKICAibmFtZXMiOiBbXQp9Cg==
