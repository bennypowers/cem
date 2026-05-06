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

// elements/cem-drawer/cem-drawer.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-drawer/cem-drawer.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse(`":host {\\n  display: flex;\\n  flex-direction: column;\\n}\\n\\n#resize-handle {\\n  height: 4px;\\n  width: 100%;\\n  cursor: ns-resize;\\n  background: transparent;\\n  position: relative;\\n  display: none;\\n\\n  :host([open]) \\u0026 {\\n    display: block;\\n  }\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    top: 50%;\\n    left: 50%;\\n    transform: translate(-50%, -50%);\\n    width: 40px;\\n    height: 4px;\\n    background: var(--pf-t--global--border--color--default);\\n    border-radius: var(--pf-t--global--border--radius--small);\\n    transition: background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n  }\\n\\n  \\u0026:hover::before,\\n  \\u0026:active::before {\\n    background: var(--pf-t--global--color--brand--default);\\n  }\\n\\n  \\u0026:focus-visible {\\n    outline: 2px solid var(--pf-t--global--color--brand--default);\\n    outline-offset: 2px;\\n  }\\n\\n  \\u0026:focus-visible::before {\\n    background: var(--pf-t--global--color--brand--default);\\n  }\\n}\\n\\n#toggle {\\n  background: var(--pf-t--global--background--color--action--plain--default);\\n  border: var(--pf-t--global--border--width--action--plain--hover) solid transparent;\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  color: var(--pf-t--global--text--color--subtle);\\n  padding: var(--pf-t--global--spacer--xs) var(--pf-t--global--spacer--sm);\\n  cursor: pointer;\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n  width: 100%;\\n  transition:\\n    background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default),\\n    color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default),\\n    transform var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--background--color--action--plain--hover);\\n    color: var(--pf-t--global--text--color--regular);\\n    border-color: var(--pf-t--global--border--color--high-contrast);\\n  }\\n\\n  \\u0026 svg {\\n    rotate: 180deg;\\n    transition: rotate var(--pf-t--global--motion--duration--md) var(--pf-t--global--motion--timing-function--decelerate);\\n  }\\n\\n  :host([open]) \\u0026 svg {\\n    rotate: 0deg;\\n  }\\n}\\n\\n#content {\\n  view-transition-name: dev-server-drawer;\\n  height: 0;\\n  overflow: hidden;\\n  display: flex;\\n  flex-direction: column;\\n  opacity: 0;\\n  contain: layout style paint;\\n\\n  \\u0026.transitions-enabled {\\n    transition:\\n      height var(--pf-t--global--motion--duration--slide-in--short) var(--pf-t--global--motion--timing-function--decelerate),\\n      opacity var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default);\\n  }\\n\\n  \\u0026.resizing {\\n    will-change: height;\\n    contain: layout style paint size;\\n  }\\n\\n  :host([open]) \\u0026 {\\n    opacity: 1;\\n  }\\n}\\n"`));
var cem_drawer_default = s;

// elements/cem-drawer/cem-drawer.ts
import "../cem-pf-v6-button/cem-pf-v6-button.js";
var CemDrawerChangeEvent = class extends Event {
  open;
  constructor(open) {
    super("change", { bubbles: true });
    this.open = open;
  }
};
var CemDrawerResizeEvent = class extends Event {
  height;
  constructor(height) {
    super("resize", { bubbles: true });
    this.height = height;
  }
};
var _drawerHeight_dec, _open_dec, _a, _CemServeDrawer_decorators, _init, _open, _drawerHeight, _isDragging, _startY, _startHeight, _maxHeight, _rafId, _pendingHeight, _CemServeDrawer_instances, $_fn, getMaxHeight_fn, _onToggleClick, _startResize, _handleResize, _applyResize, _handleKeydown, updateAriaValueNow_fn, _stopResize, _handleWindowResize;
_CemServeDrawer_decorators = [customElement("cem-drawer")];
var CemServeDrawer = class extends (_a = LitElement, _open_dec = [property({ type: Boolean, reflect: true })], _drawerHeight_dec = [property({ type: Number, reflect: true, attribute: "drawer-height" })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeDrawer_instances);
    __privateAdd(this, _open, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _drawerHeight, __runInitializers(_init, 12, this, 400)), __runInitializers(_init, 15, this);
    __privateAdd(this, _isDragging, false);
    __privateAdd(this, _startY, 0);
    __privateAdd(this, _startHeight, 0);
    __privateAdd(this, _maxHeight, null);
    __privateAdd(this, _rafId, null);
    __privateAdd(this, _pendingHeight, null);
    __privateAdd(this, _onToggleClick, () => {
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (content) {
        content.classList.add("transitions-enabled");
      }
      this.toggle();
    });
    __privateAdd(this, _startResize, (e) => {
      __privateSet(this, _isDragging, true);
      __privateSet(this, _startY, e.clientY);
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (!content) return;
      __privateSet(this, _startHeight, content.offsetHeight);
      __privateSet(this, _maxHeight, __privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this));
      content.classList.remove("transitions-enabled");
      content.classList.add("resizing");
      document.addEventListener("mousemove", __privateGet(this, _handleResize), { passive: true });
      document.addEventListener("mouseup", __privateGet(this, _stopResize));
      e.preventDefault();
    });
    __privateAdd(this, _handleResize, (e) => {
      if (!__privateGet(this, _isDragging)) return;
      const deltaY = __privateGet(this, _startY) - e.clientY;
      const newHeight = Math.max(100, Math.min(__privateGet(this, _maxHeight), __privateGet(this, _startHeight) + deltaY));
      __privateSet(this, _pendingHeight, newHeight);
      if (!__privateGet(this, _rafId)) {
        __privateSet(this, _rafId, requestAnimationFrame(__privateGet(this, _applyResize)));
      }
    });
    __privateAdd(this, _applyResize, () => {
      if (__privateGet(this, _pendingHeight) === null) return;
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (content) {
        content.style.height = `${__privateGet(this, _pendingHeight)}px`;
      }
      __privateSet(this, _pendingHeight, null);
      __privateSet(this, _rafId, null);
    });
    __privateAdd(this, _handleKeydown, (e) => {
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (!content) return;
      const step = e.shiftKey ? 50 : 10;
      let currentHeight = parseInt(content.style.height, 10) || 400;
      let newHeight = currentHeight;
      switch (e.key) {
        case "ArrowUp":
          e.preventDefault();
          newHeight = currentHeight + step;
          break;
        case "ArrowDown":
          e.preventDefault();
          newHeight = currentHeight - step;
          break;
        case "Home":
          e.preventDefault();
          newHeight = 100;
          break;
        case "End":
          e.preventDefault();
          newHeight = __privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this);
          break;
        default:
          return;
      }
      newHeight = Math.max(100, Math.min(__privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this), newHeight));
      content.style.height = `${newHeight}px`;
      __privateMethod(this, _CemServeDrawer_instances, updateAriaValueNow_fn).call(this, newHeight);
      this.dispatchEvent(new CemDrawerResizeEvent(newHeight));
    });
    __privateAdd(this, _stopResize, () => {
      __privateSet(this, _isDragging, false);
      if (__privateGet(this, _rafId)) {
        cancelAnimationFrame(__privateGet(this, _rafId));
        __privateSet(this, _rafId, null);
      }
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (!content) return;
      content.classList.add("transitions-enabled");
      content.classList.remove("resizing");
      const height = parseInt(content.style.height, 10);
      __privateMethod(this, _CemServeDrawer_instances, updateAriaValueNow_fn).call(this, height);
      this.dispatchEvent(new CemDrawerResizeEvent(height));
      document.removeEventListener("mousemove", __privateGet(this, _handleResize));
      document.removeEventListener("mouseup", __privateGet(this, _stopResize));
    });
    __privateAdd(this, _handleWindowResize, () => {
      if (!this.open) return;
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (!content) return;
      const currentHeight = parseInt(content.style.height) || 0;
      const maxHeight = __privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this);
      if (currentHeight > maxHeight) {
        content.style.height = `${maxHeight}px`;
        this.drawerHeight = maxHeight;
        this.dispatchEvent(new CemDrawerResizeEvent(maxHeight));
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    window.addEventListener("resize", __privateGet(this, _handleWindowResize));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener("resize", __privateGet(this, _handleWindowResize));
  }
  render() {
    return html`
      <div id="resize-handle"
           role="separator"
           aria-orientation="horizontal"
           aria-label="Resize drawer"
           tabindex="0"
           aria-controls="content"
           aria-valuemin="100"
           aria-valuemax="1000"
           aria-valuenow="400"
           @mousedown=${__privateGet(this, _startResize)}
           @keydown=${__privateGet(this, _handleKeydown)}></div>
      <cem-pf-v6-button id="toggle"
                     variant="plain"
                     aria-label="Toggle drawer"
                     aria-expanded="${this.open}"
                     aria-controls="content"
                     @click=${__privateGet(this, _onToggleClick)}>
        <svg width="16"
             height="16"
             viewBox="0 0 16 16"
             fill="currentColor"
             role="presentation">
          <path fill-rule="evenodd"
                d="M1.646 4.646a.5.5 0 0 1 .708 0L8 10.293l5.646-5.647a.5.5 0 0 1 .708.708l-6 6a.5.5 0 0 1-.708 0l-6-6a.5.5 0 0 1 0-.708z"/>
        </svg>
      </cem-pf-v6-button>
      <div id="content"
           style="height: ${this.open ? this.drawerHeight : 0}px;">
        <slot></slot>
      </div>
    `;
  }
  updated(changed) {
    if (changed.has("open")) {
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (content && this.open) {
        let height = this.drawerHeight;
        let needsPersist = false;
        if (height <= 0 || isNaN(height)) {
          height = 400;
          needsPersist = true;
        }
        const maxHeight = __privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this);
        if (height > maxHeight) {
          height = maxHeight;
          needsPersist = true;
        }
        if (needsPersist) {
          this.drawerHeight = height;
          this.dispatchEvent(new CemDrawerResizeEvent(height));
        }
      }
      this.dispatchEvent(new CemDrawerChangeEvent(this.open));
    }
  }
  toggle() {
    this.open = !this.open;
  }
  openDrawer() {
    this.open = true;
  }
  close() {
    this.open = false;
  }
};
_init = __decoratorStart(_a);
_open = new WeakMap();
_drawerHeight = new WeakMap();
_isDragging = new WeakMap();
_startY = new WeakMap();
_startHeight = new WeakMap();
_maxHeight = new WeakMap();
_rafId = new WeakMap();
_pendingHeight = new WeakMap();
_CemServeDrawer_instances = new WeakSet();
$_fn = function(id) {
  return this.shadowRoot?.getElementById(id);
};
/**
 * Returns the maximum safe height for the drawer content in pixels.
 *
 * The toggle button must always remain visible below the masthead so the
 * user can close or resize the drawer at any time.
 */
getMaxHeight_fn = function() {
  const toggle = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "toggle");
  const handle = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "resize-handle");
  const mastheadHeight = 56;
  const toggleHeight = toggle?.offsetHeight ?? 32;
  const handleHeight = handle?.offsetHeight ?? 4;
  return Math.max(100, window.innerHeight - mastheadHeight - toggleHeight - handleHeight);
};
_onToggleClick = new WeakMap();
_startResize = new WeakMap();
_handleResize = new WeakMap();
_applyResize = new WeakMap();
_handleKeydown = new WeakMap();
updateAriaValueNow_fn = function(height) {
  const resizeHandle = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "resize-handle");
  if (resizeHandle) {
    resizeHandle.setAttribute("aria-valuenow", String(Math.round(height)));
    resizeHandle.setAttribute("aria-valuemax", String(Math.round(__privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this))));
  }
};
_stopResize = new WeakMap();
_handleWindowResize = new WeakMap();
__decorateElement(_init, 4, "open", _open_dec, CemServeDrawer, _open);
__decorateElement(_init, 4, "drawerHeight", _drawerHeight_dec, CemServeDrawer, _drawerHeight);
CemServeDrawer = __decorateElement(_init, 0, "CemServeDrawer", _CemServeDrawer_decorators, CemServeDrawer);
__publicField(CemServeDrawer, "styles", cem_drawer_default);
__runInitializers(_init, 1, CemServeDrawer);
export {
  CemDrawerChangeEvent,
  CemDrawerResizeEvent,
  CemServeDrawer
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLWRyYXdlci9jZW0tZHJhd2VyLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLWRyYXdlci9jZW0tZHJhd2VyLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tZHJhd2VyLmNzcyc7XG5cbmltcG9ydCAnLi4vY2VtLXBmLXY2LWJ1dHRvbi9jZW0tcGYtdjYtYnV0dG9uLmpzJztcblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiB0aGUgZHJhd2VyIG9wZW5zIG9yIGNsb3Nlc1xuICovXG5leHBvcnQgY2xhc3MgQ2VtRHJhd2VyQ2hhbmdlRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIG9wZW46IGJvb2xlYW47XG4gIGNvbnN0cnVjdG9yKG9wZW46IGJvb2xlYW4pIHtcbiAgICBzdXBlcignY2hhbmdlJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMub3BlbiA9IG9wZW47XG4gIH1cbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiB0aGUgZHJhd2VyIGlzIHJlc2l6ZWRcbiAqL1xuZXhwb3J0IGNsYXNzIENlbURyYXdlclJlc2l6ZUV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBoZWlnaHQ6IG51bWJlcjtcbiAgY29uc3RydWN0b3IoaGVpZ2h0OiBudW1iZXIpIHtcbiAgICBzdXBlcigncmVzaXplJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMuaGVpZ2h0ID0gaGVpZ2h0O1xuICB9XG59XG5cbi8qKlxuICogQ0VNIFNlcnZlIERyYXdlciAtIENvbGxhcHNpYmxlLCByZXNpemFibGUgZHJhd2VyIGNvbXBvbmVudFxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBkcmF3ZXIgY29udGVudFxuICpcbiAqIEBmaXJlcyBjaGFuZ2UgLSBGaXJlcyB3aGVuIHRoZSBkcmF3ZXIgb3BlbnMgb3IgY2xvc2VzXG4gKiBAZmlyZXMgcmVzaXplIC0gRmlyZXMgd2hlbiB0aGUgZHJhd2VyIGlzIHJlc2l6ZWRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1kcmF3ZXInKVxuZXhwb3J0IGNsYXNzIENlbVNlcnZlRHJhd2VyIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBvcGVuID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogTnVtYmVyLCByZWZsZWN0OiB0cnVlLCBhdHRyaWJ1dGU6ICdkcmF3ZXItaGVpZ2h0JyB9KVxuICBhY2Nlc3NvciBkcmF3ZXJIZWlnaHQgPSA0MDA7XG5cbiAgI2lzRHJhZ2dpbmcgPSBmYWxzZTtcbiAgI3N0YXJ0WSA9IDA7XG4gICNzdGFydEhlaWdodCA9IDA7XG4gICNtYXhIZWlnaHQ6IG51bWJlciB8IG51bGwgPSBudWxsO1xuICAjcmFmSWQ6IG51bWJlciB8IG51bGwgPSBudWxsO1xuICAjcGVuZGluZ0hlaWdodDogbnVtYmVyIHwgbnVsbCA9IG51bGw7XG5cbiAgIyQoaWQ6IHN0cmluZykge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKGlkKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBSZXR1cm5zIHRoZSBtYXhpbXVtIHNhZmUgaGVpZ2h0IGZvciB0aGUgZHJhd2VyIGNvbnRlbnQgaW4gcGl4ZWxzLlxuICAgKlxuICAgKiBUaGUgdG9nZ2xlIGJ1dHRvbiBtdXN0IGFsd2F5cyByZW1haW4gdmlzaWJsZSBiZWxvdyB0aGUgbWFzdGhlYWQgc28gdGhlXG4gICAqIHVzZXIgY2FuIGNsb3NlIG9yIHJlc2l6ZSB0aGUgZHJhd2VyIGF0IGFueSB0aW1lLlxuICAgKi9cbiAgI2dldE1heEhlaWdodCgpOiBudW1iZXIge1xuICAgIGNvbnN0IHRvZ2dsZSA9IHRoaXMuIyQoJ3RvZ2dsZScpO1xuICAgIGNvbnN0IGhhbmRsZSA9IHRoaXMuIyQoJ3Jlc2l6ZS1oYW5kbGUnKTtcbiAgICBjb25zdCBtYXN0aGVhZEhlaWdodCA9IDU2O1xuICAgIGNvbnN0IHRvZ2dsZUhlaWdodCA9IHRvZ2dsZT8ub2Zmc2V0SGVpZ2h0ID8/IDMyO1xuICAgIGNvbnN0IGhhbmRsZUhlaWdodCA9IGhhbmRsZT8ub2Zmc2V0SGVpZ2h0ID8/IDQ7XG4gICAgcmV0dXJuIE1hdGgubWF4KDEwMCwgd2luZG93LmlubmVySGVpZ2h0IC0gbWFzdGhlYWRIZWlnaHQgLSB0b2dnbGVIZWlnaHQgLSBoYW5kbGVIZWlnaHQpO1xuICB9XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB3aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcigncmVzaXplJywgdGhpcy4jaGFuZGxlV2luZG93UmVzaXplKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgd2luZG93LnJlbW92ZUV2ZW50TGlzdGVuZXIoJ3Jlc2l6ZScsIHRoaXMuI2hhbmRsZVdpbmRvd1Jlc2l6ZSk7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8ZGl2IGlkPVwicmVzaXplLWhhbmRsZVwiXG4gICAgICAgICAgIHJvbGU9XCJzZXBhcmF0b3JcIlxuICAgICAgICAgICBhcmlhLW9yaWVudGF0aW9uPVwiaG9yaXpvbnRhbFwiXG4gICAgICAgICAgIGFyaWEtbGFiZWw9XCJSZXNpemUgZHJhd2VyXCJcbiAgICAgICAgICAgdGFiaW5kZXg9XCIwXCJcbiAgICAgICAgICAgYXJpYS1jb250cm9scz1cImNvbnRlbnRcIlxuICAgICAgICAgICBhcmlhLXZhbHVlbWluPVwiMTAwXCJcbiAgICAgICAgICAgYXJpYS12YWx1ZW1heD1cIjEwMDBcIlxuICAgICAgICAgICBhcmlhLXZhbHVlbm93PVwiNDAwXCJcbiAgICAgICAgICAgQG1vdXNlZG93bj0ke3RoaXMuI3N0YXJ0UmVzaXplfVxuICAgICAgICAgICBAa2V5ZG93bj0ke3RoaXMuI2hhbmRsZUtleWRvd259PjwvZGl2PlxuICAgICAgPGNlbS1wZi12Ni1idXR0b24gaWQ9XCJ0b2dnbGVcIlxuICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInBsYWluXCJcbiAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJUb2dnbGUgZHJhd2VyXCJcbiAgICAgICAgICAgICAgICAgICAgIGFyaWEtZXhwYW5kZWQ9XCIke3RoaXMub3Blbn1cIlxuICAgICAgICAgICAgICAgICAgICAgYXJpYS1jb250cm9scz1cImNvbnRlbnRcIlxuICAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25Ub2dnbGVDbGlja30+XG4gICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiXG4gICAgICAgICAgICAgaGVpZ2h0PVwiMTZcIlxuICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIlxuICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIj5cbiAgICAgICAgICA8cGF0aCBmaWxsLXJ1bGU9XCJldmVub2RkXCJcbiAgICAgICAgICAgICAgICBkPVwiTTEuNjQ2IDQuNjQ2YS41LjUgMCAwIDEgLjcwOCAwTDggMTAuMjkzbDUuNjQ2LTUuNjQ3YS41LjUgMCAwIDEgLjcwOC43MDhsLTYgNmEuNS41IDAgMCAxLS43MDggMGwtNi02YS41LjUgMCAwIDEgMC0uNzA4elwiLz5cbiAgICAgICAgPC9zdmc+XG4gICAgICA8L2NlbS1wZi12Ni1idXR0b24+XG4gICAgICA8ZGl2IGlkPVwiY29udGVudFwiXG4gICAgICAgICAgIHN0eWxlPVwiaGVpZ2h0OiAke3RoaXMub3BlbiA/IHRoaXMuZHJhd2VySGVpZ2h0IDogMH1weDtcIj5cbiAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gIHVwZGF0ZWQoY2hhbmdlZDogTWFwPHN0cmluZywgdW5rbm93bj4pIHtcbiAgICBpZiAoY2hhbmdlZC5oYXMoJ29wZW4nKSkge1xuICAgICAgY29uc3QgY29udGVudCA9IHRoaXMuIyQoJ2NvbnRlbnQnKTtcbiAgICAgIGlmIChjb250ZW50ICYmIHRoaXMub3Blbikge1xuICAgICAgICBsZXQgaGVpZ2h0ID0gdGhpcy5kcmF3ZXJIZWlnaHQ7XG4gICAgICAgIGxldCBuZWVkc1BlcnNpc3QgPSBmYWxzZTtcbiAgICAgICAgaWYgKGhlaWdodCA8PSAwIHx8IGlzTmFOKGhlaWdodCkpIHtcbiAgICAgICAgICBoZWlnaHQgPSA0MDA7XG4gICAgICAgICAgbmVlZHNQZXJzaXN0ID0gdHJ1ZTtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCBtYXhIZWlnaHQgPSB0aGlzLiNnZXRNYXhIZWlnaHQoKTtcbiAgICAgICAgaWYgKGhlaWdodCA+IG1heEhlaWdodCkge1xuICAgICAgICAgIGhlaWdodCA9IG1heEhlaWdodDtcbiAgICAgICAgICBuZWVkc1BlcnNpc3QgPSB0cnVlO1xuICAgICAgICB9XG4gICAgICAgIGlmIChuZWVkc1BlcnNpc3QpIHtcbiAgICAgICAgICB0aGlzLmRyYXdlckhlaWdodCA9IGhlaWdodDtcbiAgICAgICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IENlbURyYXdlclJlc2l6ZUV2ZW50KGhlaWdodCkpO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IENlbURyYXdlckNoYW5nZUV2ZW50KHRoaXMub3BlbikpO1xuICAgIH1cbiAgfVxuXG4gICNvblRvZ2dsZUNsaWNrID0gKCkgPT4ge1xuICAgIGNvbnN0IGNvbnRlbnQgPSB0aGlzLiMkKCdjb250ZW50Jyk7XG4gICAgaWYgKGNvbnRlbnQpIHtcbiAgICAgIGNvbnRlbnQuY2xhc3NMaXN0LmFkZCgndHJhbnNpdGlvbnMtZW5hYmxlZCcpO1xuICAgIH1cbiAgICB0aGlzLnRvZ2dsZSgpO1xuICB9O1xuXG4gIHRvZ2dsZSgpIHtcbiAgICB0aGlzLm9wZW4gPSAhdGhpcy5vcGVuO1xuICB9XG5cbiAgb3BlbkRyYXdlcigpIHtcbiAgICB0aGlzLm9wZW4gPSB0cnVlO1xuICB9XG5cbiAgY2xvc2UoKSB7XG4gICAgdGhpcy5vcGVuID0gZmFsc2U7XG4gIH1cblxuICAjc3RhcnRSZXNpemUgPSAoZTogTW91c2VFdmVudCkgPT4ge1xuICAgIHRoaXMuI2lzRHJhZ2dpbmcgPSB0cnVlO1xuICAgIHRoaXMuI3N0YXJ0WSA9IGUuY2xpZW50WTtcbiAgICBjb25zdCBjb250ZW50ID0gdGhpcy4jJCgnY29udGVudCcpO1xuICAgIGlmICghY29udGVudCkgcmV0dXJuO1xuICAgIHRoaXMuI3N0YXJ0SGVpZ2h0ID0gY29udGVudC5vZmZzZXRIZWlnaHQ7XG4gICAgdGhpcy4jbWF4SGVpZ2h0ID0gdGhpcy4jZ2V0TWF4SGVpZ2h0KCk7XG5cbiAgICBjb250ZW50LmNsYXNzTGlzdC5yZW1vdmUoJ3RyYW5zaXRpb25zLWVuYWJsZWQnKTtcbiAgICBjb250ZW50LmNsYXNzTGlzdC5hZGQoJ3Jlc2l6aW5nJyk7XG5cbiAgICBkb2N1bWVudC5hZGRFdmVudExpc3RlbmVyKCdtb3VzZW1vdmUnLCB0aGlzLiNoYW5kbGVSZXNpemUsIHsgcGFzc2l2ZTogdHJ1ZSB9KTtcbiAgICBkb2N1bWVudC5hZGRFdmVudExpc3RlbmVyKCdtb3VzZXVwJywgdGhpcy4jc3RvcFJlc2l6ZSk7XG5cbiAgICBlLnByZXZlbnREZWZhdWx0KCk7XG4gIH07XG5cbiAgI2hhbmRsZVJlc2l6ZSA9IChlOiBNb3VzZUV2ZW50KSA9PiB7XG4gICAgaWYgKCF0aGlzLiNpc0RyYWdnaW5nKSByZXR1cm47XG5cbiAgICBjb25zdCBkZWx0YVkgPSB0aGlzLiNzdGFydFkgLSBlLmNsaWVudFk7XG4gICAgY29uc3QgbmV3SGVpZ2h0ID0gTWF0aC5tYXgoMTAwLCBNYXRoLm1pbih0aGlzLiNtYXhIZWlnaHQhLCB0aGlzLiNzdGFydEhlaWdodCArIGRlbHRhWSkpO1xuXG4gICAgdGhpcy4jcGVuZGluZ0hlaWdodCA9IG5ld0hlaWdodDtcblxuICAgIGlmICghdGhpcy4jcmFmSWQpIHtcbiAgICAgIHRoaXMuI3JhZklkID0gcmVxdWVzdEFuaW1hdGlvbkZyYW1lKHRoaXMuI2FwcGx5UmVzaXplKTtcbiAgICB9XG4gIH07XG5cbiAgI2FwcGx5UmVzaXplID0gKCkgPT4ge1xuICAgIGlmICh0aGlzLiNwZW5kaW5nSGVpZ2h0ID09PSBudWxsKSByZXR1cm47XG5cbiAgICBjb25zdCBjb250ZW50ID0gdGhpcy4jJCgnY29udGVudCcpO1xuICAgIGlmIChjb250ZW50KSB7XG4gICAgICBjb250ZW50LnN0eWxlLmhlaWdodCA9IGAke3RoaXMuI3BlbmRpbmdIZWlnaHR9cHhgO1xuICAgIH1cblxuICAgIHRoaXMuI3BlbmRpbmdIZWlnaHQgPSBudWxsO1xuICAgIHRoaXMuI3JhZklkID0gbnVsbDtcbiAgfTtcblxuICAjaGFuZGxlS2V5ZG93biA9IChlOiBLZXlib2FyZEV2ZW50KSA9PiB7XG4gICAgY29uc3QgY29udGVudCA9IHRoaXMuIyQoJ2NvbnRlbnQnKTtcbiAgICBpZiAoIWNvbnRlbnQpIHJldHVybjtcblxuICAgIGNvbnN0IHN0ZXAgPSBlLnNoaWZ0S2V5ID8gNTAgOiAxMDtcbiAgICBsZXQgY3VycmVudEhlaWdodCA9IHBhcnNlSW50KGNvbnRlbnQuc3R5bGUuaGVpZ2h0LCAxMCkgfHwgNDAwO1xuICAgIGxldCBuZXdIZWlnaHQgPSBjdXJyZW50SGVpZ2h0O1xuXG4gICAgc3dpdGNoIChlLmtleSkge1xuICAgICAgY2FzZSAnQXJyb3dVcCc6XG4gICAgICAgIGUucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgbmV3SGVpZ2h0ID0gY3VycmVudEhlaWdodCArIHN0ZXA7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnQXJyb3dEb3duJzpcbiAgICAgICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBuZXdIZWlnaHQgPSBjdXJyZW50SGVpZ2h0IC0gc3RlcDtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdIb21lJzpcbiAgICAgICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBuZXdIZWlnaHQgPSAxMDA7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnRW5kJzpcbiAgICAgICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBuZXdIZWlnaHQgPSB0aGlzLiNnZXRNYXhIZWlnaHQoKTtcbiAgICAgICAgYnJlYWs7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgbmV3SGVpZ2h0ID0gTWF0aC5tYXgoMTAwLCBNYXRoLm1pbih0aGlzLiNnZXRNYXhIZWlnaHQoKSwgbmV3SGVpZ2h0KSk7XG5cbiAgICBjb250ZW50LnN0eWxlLmhlaWdodCA9IGAke25ld0hlaWdodH1weGA7XG4gICAgdGhpcy4jdXBkYXRlQXJpYVZhbHVlTm93KG5ld0hlaWdodCk7XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBDZW1EcmF3ZXJSZXNpemVFdmVudChuZXdIZWlnaHQpKTtcbiAgfTtcblxuICAjdXBkYXRlQXJpYVZhbHVlTm93KGhlaWdodDogbnVtYmVyKSB7XG4gICAgY29uc3QgcmVzaXplSGFuZGxlID0gdGhpcy4jJCgncmVzaXplLWhhbmRsZScpO1xuICAgIGlmIChyZXNpemVIYW5kbGUpIHtcbiAgICAgIHJlc2l6ZUhhbmRsZS5zZXRBdHRyaWJ1dGUoJ2FyaWEtdmFsdWVub3cnLCBTdHJpbmcoTWF0aC5yb3VuZChoZWlnaHQpKSk7XG4gICAgICByZXNpemVIYW5kbGUuc2V0QXR0cmlidXRlKCdhcmlhLXZhbHVlbWF4JywgU3RyaW5nKE1hdGgucm91bmQodGhpcy4jZ2V0TWF4SGVpZ2h0KCkpKSk7XG4gICAgfVxuICB9XG5cbiAgI3N0b3BSZXNpemUgPSAoKSA9PiB7XG4gICAgdGhpcy4jaXNEcmFnZ2luZyA9IGZhbHNlO1xuXG4gICAgaWYgKHRoaXMuI3JhZklkKSB7XG4gICAgICBjYW5jZWxBbmltYXRpb25GcmFtZSh0aGlzLiNyYWZJZCk7XG4gICAgICB0aGlzLiNyYWZJZCA9IG51bGw7XG4gICAgfVxuXG4gICAgY29uc3QgY29udGVudCA9IHRoaXMuIyQoJ2NvbnRlbnQnKTtcbiAgICBpZiAoIWNvbnRlbnQpIHJldHVybjtcbiAgICBjb250ZW50LmNsYXNzTGlzdC5hZGQoJ3RyYW5zaXRpb25zLWVuYWJsZWQnKTtcbiAgICBjb250ZW50LmNsYXNzTGlzdC5yZW1vdmUoJ3Jlc2l6aW5nJyk7XG5cbiAgICBjb25zdCBoZWlnaHQgPSBwYXJzZUludChjb250ZW50LnN0eWxlLmhlaWdodCwgMTApO1xuICAgIHRoaXMuI3VwZGF0ZUFyaWFWYWx1ZU5vdyhoZWlnaHQpO1xuXG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBDZW1EcmF3ZXJSZXNpemVFdmVudChoZWlnaHQpKTtcblxuICAgIGRvY3VtZW50LnJlbW92ZUV2ZW50TGlzdGVuZXIoJ21vdXNlbW92ZScsIHRoaXMuI2hhbmRsZVJlc2l6ZSk7XG4gICAgZG9jdW1lbnQucmVtb3ZlRXZlbnRMaXN0ZW5lcignbW91c2V1cCcsIHRoaXMuI3N0b3BSZXNpemUpO1xuICB9O1xuXG4gICNoYW5kbGVXaW5kb3dSZXNpemUgPSAoKSA9PiB7XG4gICAgaWYgKCF0aGlzLm9wZW4pIHJldHVybjtcbiAgICBjb25zdCBjb250ZW50ID0gdGhpcy4jJCgnY29udGVudCcpO1xuICAgIGlmICghY29udGVudCkgcmV0dXJuO1xuICAgIGNvbnN0IGN1cnJlbnRIZWlnaHQgPSBwYXJzZUludChjb250ZW50LnN0eWxlLmhlaWdodCkgfHwgMDtcbiAgICBjb25zdCBtYXhIZWlnaHQgPSB0aGlzLiNnZXRNYXhIZWlnaHQoKTtcbiAgICBpZiAoY3VycmVudEhlaWdodCA+IG1heEhlaWdodCkge1xuICAgICAgY29udGVudC5zdHlsZS5oZWlnaHQgPSBgJHttYXhIZWlnaHR9cHhgO1xuICAgICAgdGhpcy5kcmF3ZXJIZWlnaHQgPSBtYXhIZWlnaHQ7XG4gICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IENlbURyYXdlclJlc2l6ZUV2ZW50KG1heEhlaWdodCkpO1xuICAgIH1cbiAgfTtcbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLWRyYXdlcic6IENlbVNlcnZlRHJhd2VyO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG59XFxcXG5cXFxcbiNyZXNpemUtaGFuZGxlIHtcXFxcbiAgaGVpZ2h0OiA0cHg7XFxcXG4gIHdpZHRoOiAxMDAlO1xcXFxuICBjdXJzb3I6IG5zLXJlc2l6ZTtcXFxcbiAgYmFja2dyb3VuZDogdHJhbnNwYXJlbnQ7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgZGlzcGxheTogbm9uZTtcXFxcblxcXFxuICA6aG9zdChbb3Blbl0pIFxcXFx1MDAyNiB7XFxcXG4gICAgZGlzcGxheTogYmxvY2s7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgY29udGVudDogJyc7XFxcXG4gICAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICAgIHRvcDogNTAlO1xcXFxuICAgIGxlZnQ6IDUwJTtcXFxcbiAgICB0cmFuc2Zvcm06IHRyYW5zbGF0ZSgtNTAlLCAtNTAlKTtcXFxcbiAgICB3aWR0aDogNDBweDtcXFxcbiAgICBoZWlnaHQ6IDRweDtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXNtYWxsKTtcXFxcbiAgICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kLWNvbG9yIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyOjpiZWZvcmUsXFxcXG4gIFxcXFx1MDAyNjphY3RpdmU6OmJlZm9yZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmZvY3VzLXZpc2libGUge1xcXFxuICAgIG91dGxpbmU6IDJweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG4gICAgb3V0bGluZS1vZmZzZXQ6IDJweDtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjpmb2N1cy12aXNpYmxlOjpiZWZvcmUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jdG9nZ2xlIHtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1hY3Rpb24tLXBsYWluLS1ob3Zlcikgc29saWQgdHJhbnNwYXJlbnQ7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXNtYWxsKTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS14cykgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBqdXN0aWZ5LWNvbnRlbnQ6IGNlbnRlcjtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG4gIHRyYW5zaXRpb246XFxcXG4gICAgYmFja2dyb3VuZC1jb2xvciB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLXNob3J0KSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KSxcXFxcbiAgICBjb2xvciB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLXNob3J0KSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KSxcXFxcbiAgICB0cmFuc2Zvcm0gdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1zaG9ydCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1ob3Zlcik7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIGJvcmRlci1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1oaWdoLWNvbnRyYXN0KTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNiBzdmcge1xcXFxuICAgIHJvdGF0ZTogMTgwZGVnO1xcXFxuICAgIHRyYW5zaXRpb246IHJvdGF0ZSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLW1kKSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWNlbGVyYXRlKTtcXFxcbiAgfVxcXFxuXFxcXG4gIDpob3N0KFtvcGVuXSkgXFxcXHUwMDI2IHN2ZyB7XFxcXG4gICAgcm90YXRlOiAwZGVnO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNjb250ZW50IHtcXFxcbiAgdmlldy10cmFuc2l0aW9uLW5hbWU6IGRldi1zZXJ2ZXItZHJhd2VyO1xcXFxuICBoZWlnaHQ6IDA7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIG9wYWNpdHk6IDA7XFxcXG4gIGNvbnRhaW46IGxheW91dCBzdHlsZSBwYWludDtcXFxcblxcXFxuICBcXFxcdTAwMjYudHJhbnNpdGlvbnMtZW5hYmxlZCB7XFxcXG4gICAgdHJhbnNpdGlvbjpcXFxcbiAgICAgIGhlaWdodCB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLXNsaWRlLWluLS1zaG9ydCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVjZWxlcmF0ZSksXFxcXG4gICAgICBvcGFjaXR5IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tZGVmYXVsdCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYucmVzaXppbmcge1xcXFxuICAgIHdpbGwtY2hhbmdlOiBoZWlnaHQ7XFxcXG4gICAgY29udGFpbjogbGF5b3V0IHN0eWxlIHBhaW50IHNpemU7XFxcXG4gIH1cXFxcblxcXFxuICA6aG9zdChbb3Blbl0pIFxcXFx1MDAyNiB7XFxcXG4gICAgb3BhY2l0eTogMTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLDJpR0FBNmlHLENBQUM7QUFBRSxJQUFPLHFCQUFROzs7QURNcG5HLE9BQU87QUFLQSxJQUFNLHVCQUFOLGNBQW1DLE1BQU07QUFBQSxFQUM5QztBQUFBLEVBQ0EsWUFBWSxNQUFlO0FBQ3pCLFVBQU0sVUFBVSxFQUFFLFNBQVMsS0FBSyxDQUFDO0FBQ2pDLFNBQUssT0FBTztBQUFBLEVBQ2Q7QUFDRjtBQUtPLElBQU0sdUJBQU4sY0FBbUMsTUFBTTtBQUFBLEVBQzlDO0FBQUEsRUFDQSxZQUFZLFFBQWdCO0FBQzFCLFVBQU0sVUFBVSxFQUFFLFNBQVMsS0FBSyxDQUFDO0FBQ2pDLFNBQUssU0FBUztBQUFBLEVBQ2hCO0FBQ0Y7QUE1QkE7QUFzQ0EsOEJBQUMsY0FBYyxZQUFZO0FBQ3BCLElBQU0saUJBQU4sZUFBNkIsaUJBR2xDLGFBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxxQkFBQyxTQUFTLEVBQUUsTUFBTSxRQUFRLFNBQVMsTUFBTSxXQUFXLGdCQUFnQixDQUFDLElBTm5DLElBQVc7QUFBQSxFQUF4QztBQUFBO0FBQUE7QUFJTCx1QkFBUyxPQUFPLGtCQUFoQixnQkFBZ0IsU0FBaEI7QUFHQSx1QkFBUyxlQUFl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
