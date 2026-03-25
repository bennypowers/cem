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

// elements/cem-connection-status/cem-connection-status.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-connection-status/cem-connection-status.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  position: fixed;\\n  bottom: 20px;\\n  right: 20px;\\n  padding: 8px 12px;\\n  border-radius: 6px;\\n  font-family: system-ui, -apple-system, sans-serif;\\n  font-size: 12px;\\n  font-weight: 500;\\n  z-index: 999999;\\n  display: flex;\\n  align-items: center;\\n  gap: 8px;\\n  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);\\n  transition: opacity 0.3s ease;\\n}\\n\\n:host([state=\\"connected\\"]) {\\n  color: #10b981;\\n  background-color: #064e3b;\\n  border: 1px solid #10b981;\\n}\\n\\n:host([state=\\"reconnecting\\"]) {\\n  color: #fbbf24;\\n  background-color: #78350f;\\n  border: 1px solid #fbbf24;\\n}\\n\\n:host([state=\\"disconnected\\"]) {\\n  color: #ef4444;\\n  background-color: #7f1d1d;\\n  border: 1px solid #ef4444;\\n}\\n\\n:host([state=\\"connected\\"][faded]) {\\n  opacity: 0.3;\\n}\\n"'));
var cem_connection_status_default = s;

// elements/cem-connection-status/cem-connection-status.ts
var ICONS = {
  connected: "\u{1F7E2}",
  reconnecting: "\u{1F7E1}",
  disconnected: "\u{1F534}"
};
var _message_dec, _faded_dec, _state_dec, _a, _CemConnectionStatus_decorators, _init, _state, _faded, _message, _fadeTimeout, _hideTimeout;
_CemConnectionStatus_decorators = [customElement("cem-connection-status")];
var CemConnectionStatus = class extends (_a = LitElement, _state_dec = [property({ reflect: true })], _faded_dec = [property({ type: Boolean, reflect: true })], _message_dec = [property()], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _state, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
    __privateAdd(this, _faded, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateAdd(this, _message, __runInitializers(_init, 16, this, "")), __runInitializers(_init, 19, this);
    __privateAdd(this, _fadeTimeout, null);
    __privateAdd(this, _hideTimeout, null);
  }
  render() {
    return html`
      <span id="icon">${this.state ? ICONS[this.state] : ""}</span>
      <span id="message">${this.message}</span>
    `;
  }
  show(state, message, { fadeDelay = 0 } = {}) {
    if (__privateGet(this, _fadeTimeout) != null) {
      clearTimeout(__privateGet(this, _fadeTimeout));
      __privateSet(this, _fadeTimeout, null);
    }
    if (__privateGet(this, _hideTimeout) != null) {
      clearTimeout(__privateGet(this, _hideTimeout));
      __privateSet(this, _hideTimeout, null);
    }
    this.state = state;
    this.faded = false;
    this.message = message;
    if (fadeDelay > 0 && state === "connected") {
      __privateSet(this, _fadeTimeout, setTimeout(() => {
        this.faded = true;
      }, fadeDelay));
    }
  }
  hide() {
    this.style.opacity = "0";
    __privateSet(this, _hideTimeout, setTimeout(() => this.remove(), 300));
  }
};
_init = __decoratorStart(_a);
_state = new WeakMap();
_faded = new WeakMap();
_message = new WeakMap();
_fadeTimeout = new WeakMap();
_hideTimeout = new WeakMap();
__decorateElement(_init, 4, "state", _state_dec, CemConnectionStatus, _state);
__decorateElement(_init, 4, "faded", _faded_dec, CemConnectionStatus, _faded);
__decorateElement(_init, 4, "message", _message_dec, CemConnectionStatus, _message);
CemConnectionStatus = __decorateElement(_init, 0, "CemConnectionStatus", _CemConnectionStatus_decorators, CemConnectionStatus);
__publicField(CemConnectionStatus, "styles", cem_connection_status_default);
__runInitializers(_init, 1, CemConnectionStatus);
export {
  CemConnectionStatus
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLWNvbm5lY3Rpb24tc3RhdHVzL2NlbS1jb25uZWN0aW9uLXN0YXR1cy50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL2NlbS1jb25uZWN0aW9uLXN0YXR1cy9jZW0tY29ubmVjdGlvbi1zdGF0dXMuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1jb25uZWN0aW9uLXN0YXR1cy5jc3MnO1xuXG50eXBlIENvbm5lY3Rpb25TdGF0ZSA9ICdjb25uZWN0ZWQnIHwgJ3JlY29ubmVjdGluZycgfCAnZGlzY29ubmVjdGVkJztcblxuY29uc3QgSUNPTlM6IFJlY29yZDxDb25uZWN0aW9uU3RhdGUsIHN0cmluZz4gPSB7XG4gIGNvbm5lY3RlZDogJ1xcdXsxRjdFMn0nLFxuICByZWNvbm5lY3Rpbmc6ICdcXHV7MUY3RTF9JyxcbiAgZGlzY29ubmVjdGVkOiAnXFx1ezFGNTM0fScsXG59O1xuXG4vKipcbiAqIENvbm5lY3Rpb24gc3RhdHVzIHRvYXN0IGZvciBXZWJTb2NrZXQgY29ubmVjdGlvbiBzdGF0ZXMuXG4gKlxuICogQGN1c3RvbUVsZW1lbnQgY2VtLWNvbm5lY3Rpb24tc3RhdHVzXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tY29ubmVjdGlvbi1zdGF0dXMnKVxuZXhwb3J0IGNsYXNzIENlbUNvbm5lY3Rpb25TdGF0dXMgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIHN0YXRlOiBDb25uZWN0aW9uU3RhdGUgfCB1bmRlZmluZWQ7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBmYWRlZCA9IGZhbHNlO1xuXG4gIEBwcm9wZXJ0eSgpXG4gIGFjY2Vzc29yIG1lc3NhZ2UgPSAnJztcblxuICAjZmFkZVRpbWVvdXQ6IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+IHwgbnVsbCA9IG51bGw7XG4gICNoaWRlVGltZW91dDogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8c3BhbiBpZD1cImljb25cIj4ke3RoaXMuc3RhdGUgPyBJQ09OU1t0aGlzLnN0YXRlXSA6ICcnfTwvc3Bhbj5cbiAgICAgIDxzcGFuIGlkPVwibWVzc2FnZVwiPiR7dGhpcy5tZXNzYWdlfTwvc3Bhbj5cbiAgICBgO1xuICB9XG5cbiAgc2hvdyhzdGF0ZTogQ29ubmVjdGlvblN0YXRlLCBtZXNzYWdlOiBzdHJpbmcsIHsgZmFkZURlbGF5ID0gMCB9ID0ge30pOiB2b2lkIHtcbiAgICBpZiAodGhpcy4jZmFkZVRpbWVvdXQgIT0gbnVsbCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2ZhZGVUaW1lb3V0KTtcbiAgICAgIHRoaXMuI2ZhZGVUaW1lb3V0ID0gbnVsbDtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2hpZGVUaW1lb3V0ICE9IG51bGwpIHtcbiAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNoaWRlVGltZW91dCk7XG4gICAgICB0aGlzLiNoaWRlVGltZW91dCA9IG51bGw7XG4gICAgfVxuICAgIHRoaXMuc3RhdGUgPSBzdGF0ZTtcbiAgICB0aGlzLmZhZGVkID0gZmFsc2U7XG4gICAgdGhpcy5tZXNzYWdlID0gbWVzc2FnZTtcblxuICAgIGlmIChmYWRlRGVsYXkgPiAwICYmIHN0YXRlID09PSAnY29ubmVjdGVkJykge1xuICAgICAgdGhpcy4jZmFkZVRpbWVvdXQgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgdGhpcy5mYWRlZCA9IHRydWU7XG4gICAgICB9LCBmYWRlRGVsYXkpO1xuICAgIH1cbiAgfVxuXG4gIGhpZGUoKTogdm9pZCB7XG4gICAgdGhpcy5zdHlsZS5vcGFjaXR5ID0gJzAnO1xuICAgIHRoaXMuI2hpZGVUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB0aGlzLnJlbW92ZSgpLCAzMDApO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1jb25uZWN0aW9uLXN0YXR1cyc6IENlbUNvbm5lY3Rpb25TdGF0dXM7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgcG9zaXRpb246IGZpeGVkO1xcXFxuICBib3R0b206IDIwcHg7XFxcXG4gIHJpZ2h0OiAyMHB4O1xcXFxuICBwYWRkaW5nOiA4cHggMTJweDtcXFxcbiAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICBmb250LWZhbWlseTogc3lzdGVtLXVpLCAtYXBwbGUtc3lzdGVtLCBzYW5zLXNlcmlmO1xcXFxuICBmb250LXNpemU6IDEycHg7XFxcXG4gIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG4gIHotaW5kZXg6IDk5OTk5OTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgZ2FwOiA4cHg7XFxcXG4gIGJveC1zaGFkb3c6IDAgNHB4IDZweCByZ2JhKDAsIDAsIDAsIDAuMyk7XFxcXG4gIHRyYW5zaXRpb246IG9wYWNpdHkgMC4zcyBlYXNlO1xcXFxufVxcXFxuXFxcXG46aG9zdChbc3RhdGU9XFxcXFxcXCJjb25uZWN0ZWRcXFxcXFxcIl0pIHtcXFxcbiAgY29sb3I6ICMxMGI5ODE7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6ICMwNjRlM2I7XFxcXG4gIGJvcmRlcjogMXB4IHNvbGlkICMxMGI5ODE7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtzdGF0ZT1cXFxcXFxcInJlY29ubmVjdGluZ1xcXFxcXFwiXSkge1xcXFxuICBjb2xvcjogI2ZiYmYyNDtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogIzc4MzUwZjtcXFxcbiAgYm9yZGVyOiAxcHggc29saWQgI2ZiYmYyNDtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW3N0YXRlPVxcXFxcXFwiZGlzY29ubmVjdGVkXFxcXFxcXCJdKSB7XFxcXG4gIGNvbG9yOiAjZWY0NDQ0O1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiAjN2YxZDFkO1xcXFxuICBib3JkZXI6IDFweCBzb2xpZCAjZWY0NDQ0O1xcXFxufVxcXFxuXFxcXG46aG9zdChbc3RhdGU9XFxcXFxcXCJjb25uZWN0ZWRcXFxcXFxcIl1bZmFkZWRdKSB7XFxcXG4gIG9wYWNpdHk6IDAuMztcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSx3ekJBQWswQixDQUFDO0FBQUUsSUFBTyxnQ0FBUTs7O0FEUXo0QixJQUFNLFFBQXlDO0FBQUEsRUFDN0MsV0FBVztBQUFBLEVBQ1gsY0FBYztBQUFBLEVBQ2QsY0FBYztBQUNoQjtBQVpBO0FBbUJBLG1DQUFDLGNBQWMsdUJBQXVCO0FBQy9CLElBQU0sc0JBQU4sZUFBa0MsaUJBR3ZDLGNBQUMsU0FBUyxFQUFFLFNBQVMsS0FBSyxDQUFDLElBRzNCLGNBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxnQkFBQyxTQUFTLElBVDZCLElBQVc7QUFBQSxFQUE3QztBQUFBO0FBSUwsdUJBQVMsUUFBVDtBQUdBLHVCQUFTLFFBQVEsa0JBQWpCLGlCQUFpQixTQUFqQjtBQUdBLHVCQUFTLFVBQVUsa0JBQW5CLGlCQUFtQixNQUFuQjtBQUVBLHFDQUFxRDtBQUNyRCxxQ0FBcUQ7QUFBQTtBQUFBLEVBRXJELFNBQVM7QUFDUCxXQUFPO0FBQUEsd0JBQ2EsS0FBSyxRQUFRLE1BQU0sS0FBSyxLQUFLLElBQUksRUFBRTtBQUFBLDJCQUNoQyxLQUFLLE9BQU87QUFBQTtBQUFBLEVBRXJDO0FBQUEsRUFFQSxLQUFLLE9BQXdCLFNBQWlCLEVBQUUsWUFBWSxFQUFFLElBQUksQ0FBQyxHQUFTO0FBQzFFLFFBQUksbUJBQUssaUJBQWdCLE1BQU07QUFDN0IsbUJBQWEsbUJBQUssYUFBWTtBQUM5Qix5QkFBSyxjQUFlO0FBQUEsSUFDdEI7QUFDQSxRQUFJLG1CQUFLLGlCQUFnQixNQUFNO0FBQzdCLG1CQUFhLG1CQUFLLGFBQVk7QUFDOUIseUJBQUssY0FBZTtBQUFBLElBQ3RCO0FBQ0EsU0FBSyxRQUFRO0FBQ2IsU0FBSyxRQUFRO0FBQ2IsU0FBSyxVQUFVO0FBRWYsUUFBSSxZQUFZLEtBQUssVUFBVSxhQUFhO0FBQzFDLHlCQUFLLGNBQWUsV0FBVyxNQUFNO0FBQ25DLGFBQUssUUFBUTtBQUFBLE1BQ2YsR0FBRyxTQUFTO0FBQUEsSUFDZDtBQUFBLEVBQ0Y7QUFBQSxFQUVBLE9BQWE7QUFDWCxTQUFLLE1BQU0sVUFBVTtBQUNyQix1QkFBSyxjQUFlLFdBQVcsTUFBTSxLQUFLLE9BQU8sR0FBRyxHQUFHO0FBQUEsRUFDekQ7QUFDRjtBQTlDTztBQUlJO0FBR0E7QUFHQTtBQUVUO0FBQ0E7QUFUQSw0QkFBUyxTQURULFlBSFcscUJBSUY7QUFHVCw0QkFBUyxTQURULFlBTlcscUJBT0Y7QUFHVCw0QkFBUyxXQURULGNBVFcscUJBVUY7QUFWRSxzQkFBTixtREFEUCxpQ0FDYTtBQUNYLGNBRFcscUJBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
