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

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-drawer/cem-drawer.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse(`":host {\\n  display: flex;\\n  flex-direction: column;\\n}\\n\\n#resize-handle {\\n  height: 4px;\\n  width: 100%;\\n  cursor: ns-resize;\\n  background: transparent;\\n  position: relative;\\n  display: none;\\n\\n  :host([open]) \\u0026 {\\n    display: block;\\n  }\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    top: 50%;\\n    left: 50%;\\n    transform: translate(-50%, -50%);\\n    width: 40px;\\n    height: 4px;\\n    background: var(--pf-t--global--border--color--default);\\n    border-radius: var(--pf-t--global--border--radius--small);\\n    transition: background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n  }\\n\\n  \\u0026:hover::before,\\n  \\u0026:active::before {\\n    background: var(--pf-t--global--color--brand--default);\\n  }\\n\\n  \\u0026:focus-visible {\\n    outline: 2px solid var(--pf-t--global--color--brand--default);\\n    outline-offset: 2px;\\n  }\\n\\n  \\u0026:focus-visible::before {\\n    background: var(--pf-t--global--color--brand--default);\\n  }\\n}\\n\\n#toggle {\\n  background: var(--pf-t--global--background--color--action--plain--default);\\n  border: var(--pf-t--global--border--width--action--plain--hover) solid transparent;\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  color: var(--pf-t--global--text--color--subtle);\\n  padding: var(--pf-t--global--spacer--xs) var(--pf-t--global--spacer--sm);\\n  cursor: pointer;\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n  width: 100%;\\n  transition:\\n    background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default),\\n    color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default),\\n    transform var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--background--color--action--plain--hover);\\n    color: var(--pf-t--global--text--color--regular);\\n    border-color: var(--pf-t--global--border--color--high-contrast);\\n  }\\n\\n  \\u0026 svg {\\n    rotate: 180deg;\\n    transition: rotate var(--pf-t--global--motion--duration--md) var(--pf-t--global--motion--timing-function--decelerate);\\n  }\\n\\n  :host([open]) \\u0026 svg {\\n    rotate: 0deg;\\n  }\\n}\\n\\n#content {\\n  view-transition-name: dev-server-drawer;\\n  height: 0;\\n  overflow: hidden;\\n  display: flex;\\n  flex-direction: column;\\n  opacity: 0;\\n  contain: layout style paint;\\n\\n  \\u0026.transitions-enabled {\\n    transition:\\n      height var(--pf-t--global--motion--duration--slide-in--short) var(--pf-t--global--motion--timing-function--decelerate),\\n      opacity var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default);\\n  }\\n\\n  \\u0026.resizing {\\n    will-change: height;\\n    contain: layout style paint size;\\n  }\\n\\n  :host([open]) \\u0026 {\\n    opacity: 1;\\n  }\\n}\\n"`));
var cem_drawer_default = s;

// elements/cem-drawer/cem-drawer.ts
import "../pf-v6-button/pf-v6-button.js";
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
      <pf-v6-button id="toggle"
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
      </pf-v6-button>
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLWRyYXdlci9jZW0tZHJhd2VyLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvY2VtLWRyYXdlci9jZW0tZHJhd2VyLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tZHJhd2VyLmNzcyc7XG5cbmltcG9ydCAnLi4vcGYtdjYtYnV0dG9uL3BmLXY2LWJ1dHRvbi5qcyc7XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gdGhlIGRyYXdlciBvcGVucyBvciBjbG9zZXNcbiAqL1xuZXhwb3J0IGNsYXNzIENlbURyYXdlckNoYW5nZUV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBvcGVuOiBib29sZWFuO1xuICBjb25zdHJ1Y3RvcihvcGVuOiBib29sZWFuKSB7XG4gICAgc3VwZXIoJ2NoYW5nZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLm9wZW4gPSBvcGVuO1xuICB9XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gdGhlIGRyYXdlciBpcyByZXNpemVkXG4gKi9cbmV4cG9ydCBjbGFzcyBDZW1EcmF3ZXJSZXNpemVFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgaGVpZ2h0OiBudW1iZXI7XG4gIGNvbnN0cnVjdG9yKGhlaWdodDogbnVtYmVyKSB7XG4gICAgc3VwZXIoJ3Jlc2l6ZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLmhlaWdodCA9IGhlaWdodDtcbiAgfVxufVxuXG4vKipcbiAqIENFTSBTZXJ2ZSBEcmF3ZXIgLSBDb2xsYXBzaWJsZSwgcmVzaXphYmxlIGRyYXdlciBjb21wb25lbnRcbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgZHJhd2VyIGNvbnRlbnRcbiAqXG4gKiBAZmlyZXMgY2hhbmdlIC0gRmlyZXMgd2hlbiB0aGUgZHJhd2VyIG9wZW5zIG9yIGNsb3Nlc1xuICogQGZpcmVzIHJlc2l6ZSAtIEZpcmVzIHdoZW4gdGhlIGRyYXdlciBpcyByZXNpemVkXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tZHJhd2VyJylcbmV4cG9ydCBjbGFzcyBDZW1TZXJ2ZURyYXdlciBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3Igb3BlbiA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IE51bWJlciwgcmVmbGVjdDogdHJ1ZSwgYXR0cmlidXRlOiAnZHJhd2VyLWhlaWdodCcgfSlcbiAgYWNjZXNzb3IgZHJhd2VySGVpZ2h0ID0gNDAwO1xuXG4gICNpc0RyYWdnaW5nID0gZmFsc2U7XG4gICNzdGFydFkgPSAwO1xuICAjc3RhcnRIZWlnaHQgPSAwO1xuICAjbWF4SGVpZ2h0OiBudW1iZXIgfCBudWxsID0gbnVsbDtcbiAgI3JhZklkOiBudW1iZXIgfCBudWxsID0gbnVsbDtcbiAgI3BlbmRpbmdIZWlnaHQ6IG51bWJlciB8IG51bGwgPSBudWxsO1xuXG4gICMkKGlkOiBzdHJpbmcpIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZChpZCk7XG4gIH1cblxuICAvKipcbiAgICogUmV0dXJucyB0aGUgbWF4aW11bSBzYWZlIGhlaWdodCBmb3IgdGhlIGRyYXdlciBjb250ZW50IGluIHBpeGVscy5cbiAgICpcbiAgICogVGhlIHRvZ2dsZSBidXR0b24gbXVzdCBhbHdheXMgcmVtYWluIHZpc2libGUgYmVsb3cgdGhlIG1hc3RoZWFkIHNvIHRoZVxuICAgKiB1c2VyIGNhbiBjbG9zZSBvciByZXNpemUgdGhlIGRyYXdlciBhdCBhbnkgdGltZS5cbiAgICovXG4gICNnZXRNYXhIZWlnaHQoKTogbnVtYmVyIHtcbiAgICBjb25zdCB0b2dnbGUgPSB0aGlzLiMkKCd0b2dnbGUnKTtcbiAgICBjb25zdCBoYW5kbGUgPSB0aGlzLiMkKCdyZXNpemUtaGFuZGxlJyk7XG4gICAgY29uc3QgbWFzdGhlYWRIZWlnaHQgPSA1NjtcbiAgICBjb25zdCB0b2dnbGVIZWlnaHQgPSB0b2dnbGU/Lm9mZnNldEhlaWdodCA/PyAzMjtcbiAgICBjb25zdCBoYW5kbGVIZWlnaHQgPSBoYW5kbGU/Lm9mZnNldEhlaWdodCA/PyA0O1xuICAgIHJldHVybiBNYXRoLm1heCgxMDAsIHdpbmRvdy5pbm5lckhlaWdodCAtIG1hc3RoZWFkSGVpZ2h0IC0gdG9nZ2xlSGVpZ2h0IC0gaGFuZGxlSGVpZ2h0KTtcbiAgfVxuXG4gIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmNvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgd2luZG93LmFkZEV2ZW50TGlzdGVuZXIoJ3Jlc2l6ZScsIHRoaXMuI2hhbmRsZVdpbmRvd1Jlc2l6ZSk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHdpbmRvdy5yZW1vdmVFdmVudExpc3RlbmVyKCdyZXNpemUnLCB0aGlzLiNoYW5kbGVXaW5kb3dSZXNpemUpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cInJlc2l6ZS1oYW5kbGVcIlxuICAgICAgICAgICByb2xlPVwic2VwYXJhdG9yXCJcbiAgICAgICAgICAgYXJpYS1vcmllbnRhdGlvbj1cImhvcml6b250YWxcIlxuICAgICAgICAgICBhcmlhLWxhYmVsPVwiUmVzaXplIGRyYXdlclwiXG4gICAgICAgICAgIHRhYmluZGV4PVwiMFwiXG4gICAgICAgICAgIGFyaWEtY29udHJvbHM9XCJjb250ZW50XCJcbiAgICAgICAgICAgYXJpYS12YWx1ZW1pbj1cIjEwMFwiXG4gICAgICAgICAgIGFyaWEtdmFsdWVtYXg9XCIxMDAwXCJcbiAgICAgICAgICAgYXJpYS12YWx1ZW5vdz1cIjQwMFwiXG4gICAgICAgICAgIEBtb3VzZWRvd249JHt0aGlzLiNzdGFydFJlc2l6ZX1cbiAgICAgICAgICAgQGtleWRvd249JHt0aGlzLiNoYW5kbGVLZXlkb3dufT48L2Rpdj5cbiAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJ0b2dnbGVcIlxuICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInBsYWluXCJcbiAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJUb2dnbGUgZHJhd2VyXCJcbiAgICAgICAgICAgICAgICAgICAgIGFyaWEtZXhwYW5kZWQ9XCIke3RoaXMub3Blbn1cIlxuICAgICAgICAgICAgICAgICAgICAgYXJpYS1jb250cm9scz1cImNvbnRlbnRcIlxuICAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25Ub2dnbGVDbGlja30+XG4gICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiXG4gICAgICAgICAgICAgaGVpZ2h0PVwiMTZcIlxuICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIlxuICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIj5cbiAgICAgICAgICA8cGF0aCBmaWxsLXJ1bGU9XCJldmVub2RkXCJcbiAgICAgICAgICAgICAgICBkPVwiTTEuNjQ2IDQuNjQ2YS41LjUgMCAwIDEgLjcwOCAwTDggMTAuMjkzbDUuNjQ2LTUuNjQ3YS41LjUgMCAwIDEgLjcwOC43MDhsLTYgNmEuNS41IDAgMCAxLS43MDggMGwtNi02YS41LjUgMCAwIDEgMC0uNzA4elwiLz5cbiAgICAgICAgPC9zdmc+XG4gICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgIDxkaXYgaWQ9XCJjb250ZW50XCJcbiAgICAgICAgICAgc3R5bGU9XCJoZWlnaHQ6ICR7dGhpcy5vcGVuID8gdGhpcy5kcmF3ZXJIZWlnaHQgOiAwfXB4O1wiPlxuICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygnb3BlbicpKSB7XG4gICAgICBjb25zdCBjb250ZW50ID0gdGhpcy4jJCgnY29udGVudCcpO1xuICAgICAgaWYgKGNvbnRlbnQgJiYgdGhpcy5vcGVuKSB7XG4gICAgICAgIGxldCBoZWlnaHQgPSB0aGlzLmRyYXdlckhlaWdodDtcbiAgICAgICAgbGV0IG5lZWRzUGVyc2lzdCA9IGZhbHNlO1xuICAgICAgICBpZiAoaGVpZ2h0IDw9IDAgfHwgaXNOYU4oaGVpZ2h0KSkge1xuICAgICAgICAgIGhlaWdodCA9IDQwMDtcbiAgICAgICAgICBuZWVkc1BlcnNpc3QgPSB0cnVlO1xuICAgICAgICB9XG4gICAgICAgIGNvbnN0IG1heEhlaWdodCA9IHRoaXMuI2dldE1heEhlaWdodCgpO1xuICAgICAgICBpZiAoaGVpZ2h0ID4gbWF4SGVpZ2h0KSB7XG4gICAgICAgICAgaGVpZ2h0ID0gbWF4SGVpZ2h0O1xuICAgICAgICAgIG5lZWRzUGVyc2lzdCA9IHRydWU7XG4gICAgICAgIH1cbiAgICAgICAgaWYgKG5lZWRzUGVyc2lzdCkge1xuICAgICAgICAgIHRoaXMuZHJhd2VySGVpZ2h0ID0gaGVpZ2h0O1xuICAgICAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgQ2VtRHJhd2VyUmVzaXplRXZlbnQoaGVpZ2h0KSk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgQ2VtRHJhd2VyQ2hhbmdlRXZlbnQodGhpcy5vcGVuKSk7XG4gICAgfVxuICB9XG5cbiAgI29uVG9nZ2xlQ2xpY2sgPSAoKSA9PiB7XG4gICAgY29uc3QgY29udGVudCA9IHRoaXMuIyQoJ2NvbnRlbnQnKTtcbiAgICBpZiAoY29udGVudCkge1xuICAgICAgY29udGVudC5jbGFzc0xpc3QuYWRkKCd0cmFuc2l0aW9ucy1lbmFibGVkJyk7XG4gICAgfVxuICAgIHRoaXMudG9nZ2xlKCk7XG4gIH07XG5cbiAgdG9nZ2xlKCkge1xuICAgIHRoaXMub3BlbiA9ICF0aGlzLm9wZW47XG4gIH1cblxuICBvcGVuRHJhd2VyKCkge1xuICAgIHRoaXMub3BlbiA9IHRydWU7XG4gIH1cblxuICBjbG9zZSgpIHtcbiAgICB0aGlzLm9wZW4gPSBmYWxzZTtcbiAgfVxuXG4gICNzdGFydFJlc2l6ZSA9IChlOiBNb3VzZUV2ZW50KSA9PiB7XG4gICAgdGhpcy4jaXNEcmFnZ2luZyA9IHRydWU7XG4gICAgdGhpcy4jc3RhcnRZID0gZS5jbGllbnRZO1xuICAgIGNvbnN0IGNvbnRlbnQgPSB0aGlzLiMkKCdjb250ZW50Jyk7XG4gICAgaWYgKCFjb250ZW50KSByZXR1cm47XG4gICAgdGhpcy4jc3RhcnRIZWlnaHQgPSBjb250ZW50Lm9mZnNldEhlaWdodDtcbiAgICB0aGlzLiNtYXhIZWlnaHQgPSB0aGlzLiNnZXRNYXhIZWlnaHQoKTtcblxuICAgIGNvbnRlbnQuY2xhc3NMaXN0LnJlbW92ZSgndHJhbnNpdGlvbnMtZW5hYmxlZCcpO1xuICAgIGNvbnRlbnQuY2xhc3NMaXN0LmFkZCgncmVzaXppbmcnKTtcblxuICAgIGRvY3VtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ21vdXNlbW92ZScsIHRoaXMuI2hhbmRsZVJlc2l6ZSwgeyBwYXNzaXZlOiB0cnVlIH0pO1xuICAgIGRvY3VtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ21vdXNldXAnLCB0aGlzLiNzdG9wUmVzaXplKTtcblxuICAgIGUucHJldmVudERlZmF1bHQoKTtcbiAgfTtcblxuICAjaGFuZGxlUmVzaXplID0gKGU6IE1vdXNlRXZlbnQpID0+IHtcbiAgICBpZiAoIXRoaXMuI2lzRHJhZ2dpbmcpIHJldHVybjtcblxuICAgIGNvbnN0IGRlbHRhWSA9IHRoaXMuI3N0YXJ0WSAtIGUuY2xpZW50WTtcbiAgICBjb25zdCBuZXdIZWlnaHQgPSBNYXRoLm1heCgxMDAsIE1hdGgubWluKHRoaXMuI21heEhlaWdodCEsIHRoaXMuI3N0YXJ0SGVpZ2h0ICsgZGVsdGFZKSk7XG5cbiAgICB0aGlzLiNwZW5kaW5nSGVpZ2h0ID0gbmV3SGVpZ2h0O1xuXG4gICAgaWYgKCF0aGlzLiNyYWZJZCkge1xuICAgICAgdGhpcy4jcmFmSWQgPSByZXF1ZXN0QW5pbWF0aW9uRnJhbWUodGhpcy4jYXBwbHlSZXNpemUpO1xuICAgIH1cbiAgfTtcblxuICAjYXBwbHlSZXNpemUgPSAoKSA9PiB7XG4gICAgaWYgKHRoaXMuI3BlbmRpbmdIZWlnaHQgPT09IG51bGwpIHJldHVybjtcblxuICAgIGNvbnN0IGNvbnRlbnQgPSB0aGlzLiMkKCdjb250ZW50Jyk7XG4gICAgaWYgKGNvbnRlbnQpIHtcbiAgICAgIGNvbnRlbnQuc3R5bGUuaGVpZ2h0ID0gYCR7dGhpcy4jcGVuZGluZ0hlaWdodH1weGA7XG4gICAgfVxuXG4gICAgdGhpcy4jcGVuZGluZ0hlaWdodCA9IG51bGw7XG4gICAgdGhpcy4jcmFmSWQgPSBudWxsO1xuICB9O1xuXG4gICNoYW5kbGVLZXlkb3duID0gKGU6IEtleWJvYXJkRXZlbnQpID0+IHtcbiAgICBjb25zdCBjb250ZW50ID0gdGhpcy4jJCgnY29udGVudCcpO1xuICAgIGlmICghY29udGVudCkgcmV0dXJuO1xuXG4gICAgY29uc3Qgc3RlcCA9IGUuc2hpZnRLZXkgPyA1MCA6IDEwO1xuICAgIGxldCBjdXJyZW50SGVpZ2h0ID0gcGFyc2VJbnQoY29udGVudC5zdHlsZS5oZWlnaHQsIDEwKSB8fCA0MDA7XG4gICAgbGV0IG5ld0hlaWdodCA9IGN1cnJlbnRIZWlnaHQ7XG5cbiAgICBzd2l0Y2ggKGUua2V5KSB7XG4gICAgICBjYXNlICdBcnJvd1VwJzpcbiAgICAgICAgZS5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBuZXdIZWlnaHQgPSBjdXJyZW50SGVpZ2h0ICsgc3RlcDtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdBcnJvd0Rvd24nOlxuICAgICAgICBlLnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIG5ld0hlaWdodCA9IGN1cnJlbnRIZWlnaHQgLSBzdGVwO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ0hvbWUnOlxuICAgICAgICBlLnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIG5ld0hlaWdodCA9IDEwMDtcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdFbmQnOlxuICAgICAgICBlLnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIG5ld0hlaWdodCA9IHRoaXMuI2dldE1heEhlaWdodCgpO1xuICAgICAgICBicmVhaztcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICBuZXdIZWlnaHQgPSBNYXRoLm1heCgxMDAsIE1hdGgubWluKHRoaXMuI2dldE1heEhlaWdodCgpLCBuZXdIZWlnaHQpKTtcblxuICAgIGNvbnRlbnQuc3R5bGUuaGVpZ2h0ID0gYCR7bmV3SGVpZ2h0fXB4YDtcbiAgICB0aGlzLiN1cGRhdGVBcmlhVmFsdWVOb3cobmV3SGVpZ2h0KTtcbiAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IENlbURyYXdlclJlc2l6ZUV2ZW50KG5ld0hlaWdodCkpO1xuICB9O1xuXG4gICN1cGRhdGVBcmlhVmFsdWVOb3coaGVpZ2h0OiBudW1iZXIpIHtcbiAgICBjb25zdCByZXNpemVIYW5kbGUgPSB0aGlzLiMkKCdyZXNpemUtaGFuZGxlJyk7XG4gICAgaWYgKHJlc2l6ZUhhbmRsZSkge1xuICAgICAgcmVzaXplSGFuZGxlLnNldEF0dHJpYnV0ZSgnYXJpYS12YWx1ZW5vdycsIFN0cmluZyhNYXRoLnJvdW5kKGhlaWdodCkpKTtcbiAgICAgIHJlc2l6ZUhhbmRsZS5zZXRBdHRyaWJ1dGUoJ2FyaWEtdmFsdWVtYXgnLCBTdHJpbmcoTWF0aC5yb3VuZCh0aGlzLiNnZXRNYXhIZWlnaHQoKSkpKTtcbiAgICB9XG4gIH1cblxuICAjc3RvcFJlc2l6ZSA9ICgpID0+IHtcbiAgICB0aGlzLiNpc0RyYWdnaW5nID0gZmFsc2U7XG5cbiAgICBpZiAodGhpcy4jcmFmSWQpIHtcbiAgICAgIGNhbmNlbEFuaW1hdGlvbkZyYW1lKHRoaXMuI3JhZklkKTtcbiAgICAgIHRoaXMuI3JhZklkID0gbnVsbDtcbiAgICB9XG5cbiAgICBjb25zdCBjb250ZW50ID0gdGhpcy4jJCgnY29udGVudCcpO1xuICAgIGlmICghY29udGVudCkgcmV0dXJuO1xuICAgIGNvbnRlbnQuY2xhc3NMaXN0LmFkZCgndHJhbnNpdGlvbnMtZW5hYmxlZCcpO1xuICAgIGNvbnRlbnQuY2xhc3NMaXN0LnJlbW92ZSgncmVzaXppbmcnKTtcblxuICAgIGNvbnN0IGhlaWdodCA9IHBhcnNlSW50KGNvbnRlbnQuc3R5bGUuaGVpZ2h0LCAxMCk7XG4gICAgdGhpcy4jdXBkYXRlQXJpYVZhbHVlTm93KGhlaWdodCk7XG5cbiAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IENlbURyYXdlclJlc2l6ZUV2ZW50KGhlaWdodCkpO1xuXG4gICAgZG9jdW1lbnQucmVtb3ZlRXZlbnRMaXN0ZW5lcignbW91c2Vtb3ZlJywgdGhpcy4jaGFuZGxlUmVzaXplKTtcbiAgICBkb2N1bWVudC5yZW1vdmVFdmVudExpc3RlbmVyKCdtb3VzZXVwJywgdGhpcy4jc3RvcFJlc2l6ZSk7XG4gIH07XG5cbiAgI2hhbmRsZVdpbmRvd1Jlc2l6ZSA9ICgpID0+IHtcbiAgICBpZiAoIXRoaXMub3BlbikgcmV0dXJuO1xuICAgIGNvbnN0IGNvbnRlbnQgPSB0aGlzLiMkKCdjb250ZW50Jyk7XG4gICAgaWYgKCFjb250ZW50KSByZXR1cm47XG4gICAgY29uc3QgY3VycmVudEhlaWdodCA9IHBhcnNlSW50KGNvbnRlbnQuc3R5bGUuaGVpZ2h0KSB8fCAwO1xuICAgIGNvbnN0IG1heEhlaWdodCA9IHRoaXMuI2dldE1heEhlaWdodCgpO1xuICAgIGlmIChjdXJyZW50SGVpZ2h0ID4gbWF4SGVpZ2h0KSB7XG4gICAgICBjb250ZW50LnN0eWxlLmhlaWdodCA9IGAke21heEhlaWdodH1weGA7XG4gICAgICB0aGlzLmRyYXdlckhlaWdodCA9IG1heEhlaWdodDtcbiAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgQ2VtRHJhd2VyUmVzaXplRXZlbnQobWF4SGVpZ2h0KSk7XG4gICAgfVxuICB9O1xufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tZHJhd2VyJzogQ2VtU2VydmVEcmF3ZXI7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbn1cXFxcblxcXFxuI3Jlc2l6ZS1oYW5kbGUge1xcXFxuICBoZWlnaHQ6IDRweDtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG4gIGN1cnNvcjogbnMtcmVzaXplO1xcXFxuICBiYWNrZ3JvdW5kOiB0cmFuc3BhcmVudDtcXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxuICBkaXNwbGF5OiBub25lO1xcXFxuXFxcXG4gIDpob3N0KFtvcGVuXSkgXFxcXHUwMDI2IHtcXFxcbiAgICBkaXNwbGF5OiBibG9jaztcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjo6YmVmb3JlIHtcXFxcbiAgICBjb250ZW50OiAnJztcXFxcbiAgICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gICAgdG9wOiA1MCU7XFxcXG4gICAgbGVmdDogNTAlO1xcXFxuICAgIHRyYW5zZm9ybTogdHJhbnNsYXRlKC01MCUsIC01MCUpO1xcXFxuICAgIHdpZHRoOiA0MHB4O1xcXFxuICAgIGhlaWdodDogNHB4O1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICAgIHRyYW5zaXRpb246IGJhY2tncm91bmQtY29sb3IgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1zaG9ydCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXI6OmJlZm9yZSxcXFxcbiAgXFxcXHUwMDI2OmFjdGl2ZTo6YmVmb3JlIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6Zm9jdXMtdmlzaWJsZSB7XFxcXG4gICAgb3V0bGluZTogMnB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgICBvdXRsaW5lLW9mZnNldDogMnB4O1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmZvY3VzLXZpc2libGU6OmJlZm9yZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiN0b2dnbGUge1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWFjdGlvbi0tcGxhaW4tLWhvdmVyKSBzb2xpZCB0cmFuc3BhcmVudDtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgcGFkZGluZzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXhzKSB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGp1c3RpZnktY29udGVudDogY2VudGVyO1xcXFxuICB3aWR0aDogMTAwJTtcXFxcbiAgdHJhbnNpdGlvbjpcXFxcbiAgICBiYWNrZ3JvdW5kLWNvbG9yIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpLFxcXFxuICAgIGNvbG9yIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpLFxcXFxuICAgIHRyYW5zZm9ybSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLXNob3J0KSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KTtcXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWhvdmVyKTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgYm9yZGVyLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWhpZ2gtY29udHJhc3QpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2IHN2ZyB7XFxcXG4gICAgcm90YXRlOiAxODBkZWc7XFxcXG4gICAgdHJhbnNpdGlvbjogcm90YXRlIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tbWQpIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlY2VsZXJhdGUpO1xcXFxuICB9XFxcXG5cXFxcbiAgOmhvc3QoW29wZW5dKSBcXFxcdTAwMjYgc3ZnIHtcXFxcbiAgICByb3RhdGU6IDBkZWc7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQge1xcXFxuICB2aWV3LXRyYW5zaXRpb24tbmFtZTogZGV2LXNlcnZlci1kcmF3ZXI7XFxcXG4gIGhlaWdodDogMDtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgb3BhY2l0eTogMDtcXFxcbiAgY29udGFpbjogbGF5b3V0IHN0eWxlIHBhaW50O1xcXFxuXFxcXG4gIFxcXFx1MDAyNi50cmFuc2l0aW9ucy1lbmFibGVkIHtcXFxcbiAgICB0cmFuc2l0aW9uOlxcXFxuICAgICAgaGVpZ2h0IHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tc2xpZGUtaW4tLXNob3J0KSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWNlbGVyYXRlKSxcXFxcbiAgICAgIG9wYWNpdHkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1kZWZhdWx0KSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5yZXNpemluZyB7XFxcXG4gICAgd2lsbC1jaGFuZ2U6IGhlaWdodDtcXFxcbiAgICBjb250YWluOiBsYXlvdXQgc3R5bGUgcGFpbnQgc2l6ZTtcXFxcbiAgfVxcXFxuXFxcXG4gIDpob3N0KFtvcGVuXSkgXFxcXHUwMDI2IHtcXFxcbiAgICBvcGFjaXR5OiAxO1xcXFxuICB9XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sMmlHQUE2aUcsQ0FBQztBQUFFLElBQU8scUJBQVE7OztBRE1wbkcsT0FBTztBQUtBLElBQU0sdUJBQU4sY0FBbUMsTUFBTTtBQUFBLEVBQzlDO0FBQUEsRUFDQSxZQUFZLE1BQWU7QUFDekIsVUFBTSxVQUFVLEVBQUUsU0FBUyxLQUFLLENBQUM7QUFDakMsU0FBSyxPQUFPO0FBQUEsRUFDZDtBQUNGO0FBS08sSUFBTSx1QkFBTixjQUFtQyxNQUFNO0FBQUEsRUFDOUM7QUFBQSxFQUNBLFlBQVksUUFBZ0I7QUFDMUIsVUFBTSxVQUFVLEVBQUUsU0FBUyxLQUFLLENBQUM7QUFDakMsU0FBSyxTQUFTO0FBQUEsRUFDaEI7QUFDRjtBQTVCQTtBQXNDQSw4QkFBQyxjQUFjLFlBQVk7QUFDcEIsSUFBTSxpQkFBTixlQUE2QixpQkFHbEMsYUFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsS0FBSyxDQUFDLElBRzFDLHFCQUFDLFNBQVMsRUFBRSxNQUFNLFFBQVEsU0FBUyxNQUFNLFdBQVcsZ0JBQWdCLENBQUMsSUFObkMsSUFBVztBQUFBLEVBQXhDO0FBQUE7QUFBQTtBQUlMLHVCQUFTLE9BQU8sa0JBQWhCLGdCQUFnQixTQUFoQjtBQUdBLHVCQUFTLGVBQWUsa0JBQXhCLGlCQUF3QixPQUF4QjtBQUVBLG9DQUFjO0FBQ2QsZ0NBQVU7QUFDVixxQ0FBZTtBQUNmLG1DQUE0QjtBQUM1QiwrQkFBd0I7QUFDeEIsdUNBQWdDO0FBMEZoQyx1Q0FBaUIsTUFBTTtBQUNyQixZQUFNLFVBQVUsc0JBQUssaUNBQUwsV0FBUTtBQUN4QixVQUFJLFNBQVM7QUFDWCxnQkFBUSxVQUFVLElBQUkscUJBQXFCO0FBQUEsTUFDN0M7QUFDQSxXQUFLLE9BQU87QUFBQSxJQUNkO0FBY0EscUNBQWUsQ0FBQyxNQUFrQjtBQUNoQyx5QkFBSyxhQUFjO0FBQ25CLHlCQUFLLFNBQVUsRUFBRTtBQUNqQixZQUFNLFVBQVUsc0JBQUssaUNBQUwsV0FBUTtBQUN4QixVQUFJLENBQUMsUUFBUztBQUNkLHlCQUFLLGNBQWUsUUFBUTtBQUM1Qix5QkFBSyxZQUFhLHNCQUFLLDRDQUFMO0FBRWxCLGNBQVEsVUFBVSxPQUFPLHFCQUFxQjtBQUM5QyxjQUFRLFVBQVUsSUFBSSxVQUFVO0FBRWhDLGVBQVMsaUJBQWlCLGFBQWEsbUJBQUssZ0JBQWUsRUFBRSxTQUFTLEtBQUssQ0FBQztBQUM1RSxlQUFTLGlCQUFpQixXQUFXLG1CQUFLLFlBQVc7QUFFckQsUUFBRSxlQUFlO0FBQUEsSUFDbkI7QUFFQSxzQ0FBZ0IsQ0FBQyxNQUFrQjtBQUNqQyxVQUFJLENBQUMsbUJBQUssYUFBYTtBQUV2QixZQUFNLFNBQVMsbUJBQUssV0FBVSxFQUFFO0FBQ2hDLFlBQU0sWUFBWSxLQUFLLElBQUksS0FBSyxLQUFLLElBQUksbUJBQUssYUFBYSxtQkFBSyxnQkFBZSxNQUFNLENBQUM7QUFFdEYseUJBQUssZ0JBQWlCO0FBRXRCLFVBQUksQ0FBQyxtQkFBSyxTQUFRO0FBQ2hCLDJCQUFLLFFBQVMsc0JBQXNCLG1CQUFLLGFBQVk7QUFBQSxNQUN2RDtBQUFBLElBQ0Y7QUFFQSxxQ0FBZSxNQUFNO0FBQ25CLFVBQUksbUJBQUssb0JBQW1CLEtBQU07QUFFbEMsWUFBTSxVQUFVLHNCQUFLLGlDQUFMLFdBQVE7QUFDeEIsVUFBSSxTQUFTO0FBQ1gsZ0JBQVEsTUFBTSxTQUFTLEdBQUcsbUJBQUssZUFBYztBQUFBLE1BQy9DO0FBRUEseUJBQUssZ0JBQWlCO0FBQ3RCLHlCQUFLLFFBQVM7QUFBQSxJQUNoQjtBQUVBLHVDQUFpQixDQUFDLE1BQXFCO0FBQ3JDLFlBQU0sVUFBVSxzQkFBSyxpQ0FBTCxXQUFRO0FBQ3hCLFVBQUksQ0FBQyxRQUFTO0FBRWQsWUFBTSxPQUFPLEVBQUUsV0FBVyxLQUFLO0FBQy9CLFVBQUksZ0JBQWdCLFNBQVMsUUFBUSxNQUFNLFFBQVEsRUFBRSxLQUFLO0FBQzFELFVBQUksWUFBWTtBQUVoQixjQUFRLEVBQUUsS0FBSztBQUFBLFFBQ2IsS0FBSztBQUNILFlBQUUsZUFBZTtBQUNqQixzQkFBWSxnQkFBZ0I7QUFDNUI7QUFBQSxRQUNGLEtBQUs7QUFDSCxZQUFFLGVBQWU7QUFDakIsc0JBQVksZ0JBQWdCO0FBQzVCO0FBQUEsUUFDRixLQUFLO0FBQ0gsWUFBRSxlQUFlO0FBQ2pCLHNCQUFZO0FBQ1o7QUFBQSxRQUNGLEtBQUs7QUFDSCxZQUFFLGVBQWU7QUFDakIsc0JBQVksc0JBQUssNENBQUw7QUFDWjtBQUFBLFFBQ0Y7QUFDRTtBQUFBLE1BQ0o7QUFFQSxrQkFBWSxLQUFLLElBQUksS0FBSyxLQUFLLElBQUksc0JBQUssNENBQUwsWUFBc0IsU0FBUyxDQUFDO0FBRW5FLGNBQVEsTUFBTSxTQUFTLEdBQUcsU0FBUztBQUNuQyw0QkFBSyxrREFBTCxXQUF5QjtBQUN6QixXQUFLLGNBQWMsSUFBSSxxQkFBcUIsU0FBUyxDQUFDO0FBQUEsSUFDeEQ7QUFVQSxvQ0FBYyxNQUFNO0FBQ2xCLHlCQUFLLGFBQWM7QUFFbkIsVUFBSSxtQkFBSyxTQUFRO0FBQ2YsNkJBQXFCLG1CQUFLLE9BQU07QUFDaEMsMkJBQUssUUFBUztBQUFBLE1BQ2hCO0FBRUEsWUFBTSxVQUFVLHNCQUFLLGlDQUFMLFdBQVE7QUFDeEIsVUFBSSxDQUFDLFFBQVM7QUFDZCxjQUFRLFVBQVUsSUFBSSxxQkFBcUI7QUFDM0MsY0FBUSxVQUFVLE9BQU8sVUFBVTtBQUVuQyxZQUFNLFNBQVMsU0FBUyxRQUFRLE1BQU0sUUFBUSxFQUFFO0FBQ2hELDRCQUFLLGtEQUFMLFdBQXlCO0FBRXpCLFdBQUssY0FBYyxJQUFJLHFCQUFxQixNQUFNLENBQUM7QUFFbkQsZUFBUyxvQkFBb0IsYUFBYSxtQkFBSyxjQUFhO0FBQzVELGVBQVMsb0JBQW9CLFdBQVcsbUJBQUssWUFBVztBQUFBLElBQzFEO0FBRUEsNENBQXNCLE1BQU07QUFDMUIsVUFBSSxDQUFDLEtBQUssS0FBTTtBQUNoQixZQUFNLFVBQVUsc0JBQUssaUNBQUwsV0FBUTtBQUN4QixVQUFJLENBQUMsUUFBUztBQUNkLFlBQU0sZ0JBQWdCLFNBQVMsUUFBUSxNQUFNLE1BQU0sS0FBSztBQUN4RCxZQUFNLFlBQVksc0JBQUssNENBQUw7QUFDbEIsVUFBSSxnQkFBZ0IsV0FBVztBQUM3QixnQkFBUSxNQUFNLFNBQVMsR0FBRyxTQUFTO0FBQ25DLGFBQUssZUFBZTtBQUNwQixhQUFLLGNBQWMsSUFBSSxxQkFBcUIsU0FBUyxDQUFDO0FBQUEsTUFDeEQ7QUFBQSxJQUNGO0FBQUE7QUFBQSxFQWhOQSxvQkFBb0I7QUFDbEIsVUFBTSxrQkFBa0I7QUFDeEIsV0FBTyxpQkFBaUIsVUFBVSxtQkFBSyxvQkFBbUI7QUFBQSxFQUM1RDtBQUFBLEVBRUEsdUJBQXVCO0FBQ3JCLFVBQU0scUJBQXFCO0FBQzNCLFdBQU8sb0JBQW9CLFVBQVUsbUJBQUssb0JBQW1CO0FBQUEsRUFDL0Q7QUFBQSxFQUVBLFNBQVM7QUFDUCxXQUFPO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsd0JBVWEsbUJBQUssYUFBWTtBQUFBLHNCQUNuQixtQkFBSyxlQUFjO0FBQUE7QUFBQTtBQUFBO0FBQUEsc0NBSUgsS0FBSyxJQUFJO0FBQUE7QUFBQSw4QkFFakIsbUJBQUssZUFBYztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsNEJBV3JCLEtBQUssT0FBTyxLQUFLLGVBQWUsQ0FBQztBQUFBO0FBQUE7QUFBQTtBQUFBLEVBSTNEO0FBQUEsRUFFQSxRQUFRLFNBQStCO0FBQ3JDLFFBQUksUUFBUSxJQUFJLE1BQU0sR0FBRztBQUN2QixZQUFNLFVBQVUsc0JBQUssaUNBQUwsV0FBUTtBQUN4QixVQUFJLFdBQVcsS0FBSyxNQUFNO0FBQ3hCLFlBQUksU0FBUyxLQUFLO0FBQ2xCLFlBQUksZUFBZTtBQUNuQixZQUFJLFVBQVUsS0FBSyxNQUFNLE1BQU0sR0FBRztBQUNoQyxtQkFBUztBQUNULHlCQUFlO0FBQUEsUUFDakI7QUFDQSxjQUFNLFlBQVksc0JBQUssNENBQUw7QUFDbEIsWUFBSSxTQUFTLFdBQVc7QUFDdEIsbUJBQVM7QUFDVCx5QkFBZTtBQUFBLFFBQ2pCO0FBQ0EsWUFBSSxjQUFjO0FBQ2hCLGVBQUssZUFBZTtBQUNwQixlQUFLLGNBQWMsSUFBSSxxQkFBcUIsTUFBTSxDQUFDO0FBQUEsUUFDckQ7QUFBQSxNQUNGO0FBQ0EsV0FBSyxjQUFjLElBQUkscUJBQXFCLEtBQUssSUFBSSxDQUFDO0FBQUEsSUFDeEQ7QUFBQSxFQUNGO0FBQUEsRUFVQSxTQUFTO0FBQ1AsU0FBSyxPQUFPLENBQUMsS0FBSztBQUFBLEVBQ3BCO0FBQUEsRUFFQSxhQUFhO0FBQ1gsU0FBSyxPQUFPO0FBQUEsRUFDZDtBQUFBLEVBRUEsUUFBUTtBQUNOLFNBQUssT0FBTztBQUFBLEVBQ2Q7QUEwSEY7QUFwUE87QUFJSTtBQUdBO0FBRVQ7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBZEs7QUFnQkwsT0FBRSxTQUFDLElBQVk7QUFDYixTQUFPLEtBQUssWUFBWSxlQUFlLEVBQUU7QUFDM0M7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFRQSxrQkFBYSxXQUFXO0FBQ3RCLFFBQU0sU0FBUyxzQkFBSyxpQ0FBTCxXQUFRO0FBQ3ZCLFFBQU0sU0FBUyxzQkFBSyxpQ0FBTCxXQUFRO0FBQ3ZCLFFBQU0saUJBQWlCO0FBQ3ZCLFFBQU0sZUFBZSxRQUFRLGdCQUFnQjtBQUM3QyxRQUFNLGVBQWUsUUFBUSxnQkFBZ0I7QUFDN0MsU0FBTyxLQUFLLElBQUksS0FBSyxPQUFPLGNBQWMsaUJBQWlCLGVBQWUsWUFBWTtBQUN4RjtBQXVFQTtBQW9CQTtBQWlCQTtBQWFBO0FBWUE7QUFvQ0Esd0JBQW1CLFNBQUMsUUFBZ0I7QUFDbEMsUUFBTSxlQUFlLHNCQUFLLGlDQUFMLFdBQVE7QUFDN0IsTUFBSSxjQUFjO0FBQ2hCLGlCQUFhLGFBQWEsaUJBQWlCLE9BQU8sS0FBSyxNQUFNLE1BQU0sQ0FBQyxDQUFDO0FBQ3JFLGlCQUFhLGFBQWEsaUJBQWlCLE9BQU8sS0FBSyxNQUFNLHNCQUFLLDRDQUFMLFVBQW9CLENBQUMsQ0FBQztBQUFBLEVBQ3JGO0FBQ0Y7QUFFQTtBQXNCQTtBQXBPQSw0QkFBUyxRQURULFdBSFcsZ0JBSUY7QUFHVCw0QkFBUyxnQkFEVCxtQkFOVyxnQkFPRjtBQVBFLGlCQUFOLDhDQURQLDRCQUNhO0FBQ1gsY0FEVyxnQkFDSixVQUFTO0FBRFgsNEJBQU07IiwKICAibmFtZXMiOiBbXQp9Cg==
