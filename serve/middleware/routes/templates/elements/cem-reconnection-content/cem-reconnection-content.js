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

// elements/cem-reconnection-content/cem-reconnection-content.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { state } from "/__cem/vendor/lit/decorators/state.js";

// lit-css:elements/cem-reconnection-content/cem-reconnection-content.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\np {\\n  margin: 0 0 var(--pf-t--global--spacer--md) 0;\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--regular);\\n}\\n\\n#retry-info {\\n  background: var(--pf-t--global--background--color--secondary--default);\\n  padding: var(--pf-t--global--spacer--md);\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  margin: 0;\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  color: var(--pf-t--global--text--color--subtle);\\n  border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n}\\n"'));
var cem_reconnection_content_default = s;

// elements/cem-reconnection-content/cem-reconnection-content.ts
var _retryText_dec, _a, _CemReconnectionContent_decorators, _init, _retryText, _b, retryText_get, retryText_set, _CemReconnectionContent_instances, _countdownInterval, _remainingMs;
_CemReconnectionContent_decorators = [customElement("cem-reconnection-content")];
var CemReconnectionContent = class extends (_a = LitElement, _retryText_dec = [state()], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemReconnectionContent_instances);
    __privateAdd(this, _retryText, __runInitializers(_init, 8, this, "")), __runInitializers(_init, 11, this);
    __privateAdd(this, _countdownInterval, null);
    __privateAdd(this, _remainingMs, 0);
  }
  render() {
    return html`
      <p>
        The connection to the development server was lost.
        Automatically retrying connection...
      </p>
      <div id="retry-info">${__privateGet(this, _CemReconnectionContent_instances, retryText_get)}</div>
    `;
  }
  /** Clear the countdown interval */
  clearCountdown() {
    if (__privateGet(this, _countdownInterval)) {
      clearInterval(__privateGet(this, _countdownInterval));
      __privateSet(this, _countdownInterval, null);
    }
  }
  /**
   * Update the retry information with countdown
   * @param attempt - Current retry attempt number
   * @param delay - Delay in milliseconds until next retry
   */
  updateRetryInfo(attempt, delay) {
    this.clearCountdown();
    __privateSet(this, _remainingMs, delay);
    if (__privateGet(this, _remainingMs) <= 0) {
      __privateSet(this, _CemReconnectionContent_instances, `Attempt #${attempt}. Connecting...`, retryText_set);
      return;
    }
    const seconds = Math.ceil(__privateGet(this, _remainingMs) / 1e3);
    __privateSet(this, _CemReconnectionContent_instances, `Attempt #${attempt}. Retrying in ${seconds}s...`, retryText_set);
    __privateSet(this, _countdownInterval, setInterval(() => {
      __privateSet(this, _remainingMs, __privateGet(this, _remainingMs) - 100);
      if (__privateGet(this, _remainingMs) < 100) {
        this.clearCountdown();
        __privateSet(this, _CemReconnectionContent_instances, `Attempt #${attempt}. Connecting...`, retryText_set);
        return;
      }
      const seconds2 = Math.ceil(__privateGet(this, _remainingMs) / 1e3);
      __privateSet(this, _CemReconnectionContent_instances, `Attempt #${attempt}. Retrying in ${seconds2}s...`, retryText_set);
    }, 100));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.clearCountdown();
  }
};
_init = __decoratorStart(_a);
_retryText = new WeakMap();
_CemReconnectionContent_instances = new WeakSet();
_countdownInterval = new WeakMap();
_remainingMs = new WeakMap();
_b = __decorateElement(_init, 20, "#retryText", _retryText_dec, _CemReconnectionContent_instances, _retryText), retryText_get = _b.get, retryText_set = _b.set;
CemReconnectionContent = __decorateElement(_init, 0, "CemReconnectionContent", _CemReconnectionContent_decorators, CemReconnectionContent);
__publicField(CemReconnectionContent, "styles", cem_reconnection_content_default);
__runInitializers(_init, 1, CemReconnectionContent);
export {
  CemReconnectionContent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXJlY29ubmVjdGlvbi1jb250ZW50L2NlbS1yZWNvbm5lY3Rpb24tY29udGVudC50cyIsICJsaXQtY3NzOmVsZW1lbnRzL2NlbS1yZWNvbm5lY3Rpb24tY29udGVudC9jZW0tcmVjb25uZWN0aW9uLWNvbnRlbnQuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1yZWNvbm5lY3Rpb24tY29udGVudC5jc3MnO1xuXG4vKipcbiAqIFJlY29ubmVjdGlvbiBkaWFsb2cgY29udGVudCBjb21wb25lbnQuXG4gKiBEaXNwbGF5cyByZXRyeSBpbmZvcm1hdGlvbiBhbmQgY291bnRkb3duIHdoZW4gc2VydmVyIGNvbm5lY3Rpb24gaXMgbG9zdC5cbiAqXG4gKiBAY3VzdG9tRWxlbWVudCBjZW0tcmVjb25uZWN0aW9uLWNvbnRlbnRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1yZWNvbm5lY3Rpb24tY29udGVudCcpXG5leHBvcnQgY2xhc3MgQ2VtUmVjb25uZWN0aW9uQ29udGVudCBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBzdGF0ZSgpXG4gIGFjY2Vzc29yICNyZXRyeVRleHQgPSAnJztcblxuICAjY291bnRkb3duSW50ZXJ2YWw6IFJldHVyblR5cGU8dHlwZW9mIHNldEludGVydmFsPiB8IG51bGwgPSBudWxsO1xuICAjcmVtYWluaW5nTXMgPSAwO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxwPlxuICAgICAgICBUaGUgY29ubmVjdGlvbiB0byB0aGUgZGV2ZWxvcG1lbnQgc2VydmVyIHdhcyBsb3N0LlxuICAgICAgICBBdXRvbWF0aWNhbGx5IHJldHJ5aW5nIGNvbm5lY3Rpb24uLi5cbiAgICAgIDwvcD5cbiAgICAgIDxkaXYgaWQ9XCJyZXRyeS1pbmZvXCI+JHt0aGlzLiNyZXRyeVRleHR9PC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gIC8qKiBDbGVhciB0aGUgY291bnRkb3duIGludGVydmFsICovXG4gIGNsZWFyQ291bnRkb3duKCk6IHZvaWQge1xuICAgIGlmICh0aGlzLiNjb3VudGRvd25JbnRlcnZhbCkge1xuICAgICAgY2xlYXJJbnRlcnZhbCh0aGlzLiNjb3VudGRvd25JbnRlcnZhbCk7XG4gICAgICB0aGlzLiNjb3VudGRvd25JbnRlcnZhbCA9IG51bGw7XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIFVwZGF0ZSB0aGUgcmV0cnkgaW5mb3JtYXRpb24gd2l0aCBjb3VudGRvd25cbiAgICogQHBhcmFtIGF0dGVtcHQgLSBDdXJyZW50IHJldHJ5IGF0dGVtcHQgbnVtYmVyXG4gICAqIEBwYXJhbSBkZWxheSAtIERlbGF5IGluIG1pbGxpc2Vjb25kcyB1bnRpbCBuZXh0IHJldHJ5XG4gICAqL1xuICB1cGRhdGVSZXRyeUluZm8oYXR0ZW1wdDogbnVtYmVyLCBkZWxheTogbnVtYmVyKTogdm9pZCB7XG4gICAgdGhpcy5jbGVhckNvdW50ZG93bigpO1xuXG4gICAgdGhpcy4jcmVtYWluaW5nTXMgPSBkZWxheTtcblxuICAgIC8vIElmIGRlbGF5IGlzIDAgb3IgbmVnYXRpdmUsIHNob3cgY29ubmVjdGluZyBpbW1lZGlhdGVseVxuICAgIGlmICh0aGlzLiNyZW1haW5pbmdNcyA8PSAwKSB7XG4gICAgICB0aGlzLiNyZXRyeVRleHQgPSBgQXR0ZW1wdCAjJHthdHRlbXB0fS4gQ29ubmVjdGluZy4uLmA7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgLy8gVXBkYXRlIGltbWVkaWF0ZWx5XG4gICAgY29uc3Qgc2Vjb25kcyA9IE1hdGguY2VpbCh0aGlzLiNyZW1haW5pbmdNcyAvIDEwMDApO1xuICAgIHRoaXMuI3JldHJ5VGV4dCA9IGBBdHRlbXB0ICMke2F0dGVtcHR9LiBSZXRyeWluZyBpbiAke3NlY29uZHN9cy4uLmA7XG5cbiAgICAvLyBVcGRhdGUgY291bnRkb3duIGV2ZXJ5IDEwMG1zIGZvciBzbW9vdGggZGlzcGxheVxuICAgIHRoaXMuI2NvdW50ZG93bkludGVydmFsID0gc2V0SW50ZXJ2YWwoKCkgPT4ge1xuICAgICAgdGhpcy4jcmVtYWluaW5nTXMgLT0gMTAwO1xuXG4gICAgICAvLyBVc2UgPCAxMDAgaW5zdGVhZCBvZiA8PSAwIHRvIGhhbmRsZSBjYXNlcyB3aGVyZSByZW1haW5pbmcgdGltZVxuICAgICAgLy8gaXMgbGVzcyB0aGFuIG9uZSBpbnRlcnZhbCBwZXJpb2QgKGF2b2lkcyBzaG93aW5nIFwiUmV0cnlpbmcgaW4gMXNcIiB3aGVuIDwgMTAwbXMgbGVmdClcbiAgICAgIGlmICh0aGlzLiNyZW1haW5pbmdNcyA8IDEwMCkge1xuICAgICAgICB0aGlzLmNsZWFyQ291bnRkb3duKCk7XG4gICAgICAgIHRoaXMuI3JldHJ5VGV4dCA9IGBBdHRlbXB0ICMke2F0dGVtcHR9LiBDb25uZWN0aW5nLi4uYDtcbiAgICAgICAgcmV0dXJuO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBzZWNvbmRzID0gTWF0aC5jZWlsKHRoaXMuI3JlbWFpbmluZ01zIC8gMTAwMCk7XG4gICAgICB0aGlzLiNyZXRyeVRleHQgPSBgQXR0ZW1wdCAjJHthdHRlbXB0fS4gUmV0cnlpbmcgaW4gJHtzZWNvbmRzfXMuLi5gO1xuICAgIH0sIDEwMCk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpOiB2b2lkIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuY2xlYXJDb3VudGRvd24oKTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcmVjb25uZWN0aW9uLWNvbnRlbnQnOiBDZW1SZWNvbm5lY3Rpb25Db250ZW50O1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxufVxcXFxuXFxcXG5wIHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKSAwO1xcXFxuICBsaW5lLWhlaWdodDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1saW5lLWhlaWdodC0tYm9keSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbn1cXFxcblxcXFxuI3JldHJ5LWluZm8ge1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWRlZmF1bHQpO1xcXFxuICBwYWRkaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCk7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtKTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGJvcmRlcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1yZWd1bGFyKSBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGFBQWE7OztBQ0Z0QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxvcEJBQXNwQixDQUFDO0FBQUUsSUFBTyxtQ0FBUTs7O0FEQTd0QjtBQVlBLHNDQUFDLGNBQWMsMEJBQTBCO0FBQ2xDLElBQU0seUJBQU4sZUFBcUMsaUJBRzFDLGtCQUFDLE1BQU0sSUFIbUMsSUFBVztBQUFBLEVBQWhEO0FBQUE7QUFBQTtBQUlMLHVCQUFTLFlBQWEsa0JBQXRCLGdCQUFzQixNQUF0QjtBQUVBLDJDQUE0RDtBQUM1RCxxQ0FBZTtBQUFBO0FBQUEsRUFFZixTQUFTO0FBQ1AsV0FBTztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsNkJBS2tCLG1CQUFLLGlEQUFVO0FBQUE7QUFBQSxFQUUxQztBQUFBO0FBQUEsRUFHQSxpQkFBdUI7QUFDckIsUUFBSSxtQkFBSyxxQkFBb0I7QUFDM0Isb0JBQWMsbUJBQUssbUJBQWtCO0FBQ3JDLHlCQUFLLG9CQUFxQjtBQUFBLElBQzVCO0FBQUEsRUFDRjtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQU9BLGdCQUFnQixTQUFpQixPQUFxQjtBQUNwRCxTQUFLLGVBQWU7QUFFcEIsdUJBQUssY0FBZTtBQUdwQixRQUFJLG1CQUFLLGlCQUFnQixHQUFHO0FBQzFCLHlCQUFLLG1DQUFhLFlBQVksT0FBTyxtQkFBaEM7QUFDTDtBQUFBLElBQ0Y7QUFHQSxVQUFNLFVBQVUsS0FBSyxLQUFLLG1CQUFLLGdCQUFlLEdBQUk7QUFDbEQsdUJBQUssbUNBQWEsWUFBWSxPQUFPLGlCQUFpQixPQUFPLFFBQXhEO0FBR0wsdUJBQUssb0JBQXFCLFlBQVksTUFBTTtBQUMxQyx5QkFBSyxjQUFMLG1CQUFLLGdCQUFnQjtBQUlyQixVQUFJLG1CQUFLLGdCQUFlLEtBQUs7QUFDM0IsYUFBSyxlQUFlO0FBQ3BCLDJCQUFLLG1DQUFhLFlBQVksT0FBTyxtQkFBaEM7QUFDTDtBQUFBLE1BQ0Y7QUFFQSxZQUFNQSxXQUFVLEtBQUssS0FBSyxtQkFBSyxnQkFBZSxHQUFJO0FBQ2xELHlCQUFLLG1DQUFhLFlBQVksT0FBTyxpQkFBaUJBLFFBQU8sUUFBeEQ7QUFBQSxJQUNQLEdBQUcsR0FBRztBQUFBLEVBQ1I7QUFBQSxFQUVBLHVCQUE2QjtBQUMzQixVQUFNLHFCQUFxQjtBQUMzQixTQUFLLGVBQWU7QUFBQSxFQUN0QjtBQUNGO0FBcEVPO0FBSUk7QUFKSjtBQU1MO0FBQ0E7QUFIQSxnREFEQSxnQkFDUyxnRUFBVCxRQUFTLGdCQUFUO0FBSlcseUJBQU4sc0RBRFAsb0NBQ2E7QUFDWCxjQURXLHdCQUNKLFVBQVM7QUFEWCw0QkFBTTsiLAogICJuYW1lcyI6IFsic2Vjb25kcyJdCn0K
