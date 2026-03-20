var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __knownSymbol = (name, symbol) => (symbol = Symbol[name]) ? symbol : Symbol.for("Symbol." + name);
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

// elements/pf-v6-alert-group/pf-v6-alert-group.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-alert-group/pf-v6-alert-group.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n\\n  --pf-v6-c-alert-group__item--MarginBlockStart: var(--pf-t--global--spacer--gap--group--vertical);\\n  --pf-v6-c-alert-group--m-toast--InsetBlockStart: var(--pf-t--global--spacer--2xl);\\n  --pf-v6-c-alert-group--m-toast--InsetInlineEnd: var(--pf-t--global--spacer--xl);\\n  --pf-v6-c-alert-group--m-toast--MaxWidth: 37.5rem;\\n  --pf-v6-c-alert-group--m-toast--ZIndex: var(--pf-t--global--z-index--2xl);\\n}\\n\\n:host([toast]) {\\n  position: fixed;\\n  inset-block-start: var(--pf-v6-c-alert-group--m-toast--InsetBlockStart);\\n  inset-inline-end: var(--pf-v6-c-alert-group--m-toast--InsetInlineEnd);\\n  z-index: var(--pf-v6-c-alert-group--m-toast--ZIndex);\\n  width: calc(100% - var(--pf-v6-c-alert-group--m-toast--InsetInlineEnd) * 2);\\n  max-width: var(--pf-v6-c-alert-group--m-toast--MaxWidth);\\n}\\n\\n:host ::slotted(*:not(:first-child)),\\n:host ::slotted(.pf-v6-c-alert-group__item:not(:first-child)) {\\n  margin-block-start: var(--pf-v6-c-alert-group__item--MarginBlockStart);\\n}\\n"'));
var pf_v6_alert_group_default = s;

// elements/pf-v6-alert-group/pf-v6-alert-group.ts
var _liveRegion_dec, _toast_dec, _a, _PfAlertGroup_decorators, _init, _toast, _liveRegion, _onClose, _PfAlertGroup_instances, removeAlert_fn;
_PfAlertGroup_decorators = [customElement("pf-v6-alert-group")];
var PfAlertGroup = class extends (_a = LitElement, _toast_dec = [property({ type: Boolean, reflect: true })], _liveRegion_dec = [property({ type: Boolean, reflect: true, attribute: "live-region" })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfAlertGroup_instances);
    __privateAdd(this, _toast, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateAdd(this, _liveRegion, __runInitializers(_init, 12, this, false)), __runInitializers(_init, 15, this);
    __privateAdd(this, _onClose, (e) => {
      const alert = e.target;
      if (alert?.tagName === "PF-V6-ALERT") {
        __privateMethod(this, _PfAlertGroup_instances, removeAlert_fn).call(this, alert);
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("close", __privateGet(this, _onClose));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("close", __privateGet(this, _onClose));
  }
  firstUpdated() {
    if (this.liveRegion) {
      this.setAttribute("aria-live", "polite");
      this.setAttribute("aria-atomic", "false");
    }
  }
  render() {
    return html`<slot></slot>`;
  }
  /**
   * Add an alert to the group
   */
  addAlert(alert) {
    if (this.toast) {
      alert.classList.add("pf-m-incoming");
      this.appendChild(alert);
      requestAnimationFrame(() => {
        alert.classList.remove("pf-m-incoming");
      });
    } else {
      this.appendChild(alert);
    }
  }
};
_init = __decoratorStart(_a);
_toast = new WeakMap();
_liveRegion = new WeakMap();
_onClose = new WeakMap();
_PfAlertGroup_instances = new WeakSet();
removeAlert_fn = function(alert) {
  if (!alert) return;
  if (this.toast) {
    alert.classList.add("pf-m-outgoing");
    alert.addEventListener("transitionend", () => {
      alert.remove();
    }, { once: true });
  } else {
    alert.remove();
  }
};
__decorateElement(_init, 4, "toast", _toast_dec, PfAlertGroup, _toast);
__decorateElement(_init, 4, "liveRegion", _liveRegion_dec, PfAlertGroup, _liveRegion);
PfAlertGroup = __decorateElement(_init, 0, "PfAlertGroup", _PfAlertGroup_decorators, PfAlertGroup);
__publicField(PfAlertGroup, "styles", pf_v6_alert_group_default);
__runInitializers(_init, 1, PfAlertGroup);
export {
  PfAlertGroup
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtYWxlcnQtZ3JvdXAvcGYtdjYtYWxlcnQtZ3JvdXAudHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9wZi12Ni1hbGVydC1ncm91cC9wZi12Ni1hbGVydC1ncm91cC5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtYWxlcnQtZ3JvdXAuY3NzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IEFsZXJ0IEdyb3VwIGNvbXBvbmVudFxuICpcbiAqIENvbnRhaW5lciBmb3IgbWFuYWdpbmcgbXVsdGlwbGUgYWxlcnRzIHdpdGggb3B0aW9uYWwgdG9hc3QgcG9zaXRpb25pbmdcbiAqIGFuZCBvdmVyZmxvdyBoYW5kbGluZy5cbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgcGYtdjYtYWxlcnQgZWxlbWVudHNcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ3BmLXY2LWFsZXJ0LWdyb3VwJylcbmV4cG9ydCBjbGFzcyBQZkFsZXJ0R3JvdXAgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIHRvYXN0ID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSwgYXR0cmlidXRlOiAnbGl2ZS1yZWdpb24nIH0pXG4gIGFjY2Vzc29yIGxpdmVSZWdpb24gPSBmYWxzZTtcblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignY2xvc2UnLCB0aGlzLiNvbkNsb3NlKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdjbG9zZScsIHRoaXMuI29uQ2xvc2UpO1xuICB9XG5cbiAgZmlyc3RVcGRhdGVkKCkge1xuICAgIGlmICh0aGlzLmxpdmVSZWdpb24pIHtcbiAgICAgIHRoaXMuc2V0QXR0cmlidXRlKCdhcmlhLWxpdmUnLCAncG9saXRlJyk7XG4gICAgICB0aGlzLnNldEF0dHJpYnV0ZSgnYXJpYS1hdG9taWMnLCAnZmFsc2UnKTtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgPHNsb3Q+PC9zbG90PmA7XG4gIH1cblxuICAvKipcbiAgICogQWRkIGFuIGFsZXJ0IHRvIHRoZSBncm91cFxuICAgKi9cbiAgYWRkQWxlcnQoYWxlcnQ6IEhUTUxFbGVtZW50KSB7XG4gICAgaWYgKHRoaXMudG9hc3QpIHtcbiAgICAgIGFsZXJ0LmNsYXNzTGlzdC5hZGQoJ3BmLW0taW5jb21pbmcnKTtcbiAgICAgIHRoaXMuYXBwZW5kQ2hpbGQoYWxlcnQpO1xuICAgICAgcmVxdWVzdEFuaW1hdGlvbkZyYW1lKCgpID0+IHtcbiAgICAgICAgYWxlcnQuY2xhc3NMaXN0LnJlbW92ZSgncGYtbS1pbmNvbWluZycpO1xuICAgICAgfSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuYXBwZW5kQ2hpbGQoYWxlcnQpO1xuICAgIH1cbiAgfVxuXG4gICNvbkNsb3NlID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgYWxlcnQgPSBlLnRhcmdldCBhcyBIVE1MRWxlbWVudDtcbiAgICBpZiAoYWxlcnQ/LnRhZ05hbWUgPT09ICdQRi1WNi1BTEVSVCcpIHtcbiAgICAgIHRoaXMuI3JlbW92ZUFsZXJ0KGFsZXJ0KTtcbiAgICB9XG4gIH07XG5cbiAgI3JlbW92ZUFsZXJ0KGFsZXJ0OiBIVE1MRWxlbWVudCkge1xuICAgIGlmICghYWxlcnQpIHJldHVybjtcblxuICAgIGlmICh0aGlzLnRvYXN0KSB7XG4gICAgICBhbGVydC5jbGFzc0xpc3QuYWRkKCdwZi1tLW91dGdvaW5nJyk7XG4gICAgICBhbGVydC5hZGRFdmVudExpc3RlbmVyKCd0cmFuc2l0aW9uZW5kJywgKCkgPT4ge1xuICAgICAgICBhbGVydC5yZW1vdmUoKTtcbiAgICAgIH0sIHsgb25jZTogdHJ1ZSB9KTtcbiAgICB9IGVsc2Uge1xuICAgICAgYWxlcnQucmVtb3ZlKCk7XG4gICAgfVxuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LWFsZXJ0LWdyb3VwJzogUGZBbGVydEdyb3VwO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuXFxcXG4gIC0tcGYtdjYtYy1hbGVydC1ncm91cF9faXRlbS0tTWFyZ2luQmxvY2tTdGFydDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWdhcC0tZ3JvdXAtLXZlcnRpY2FsKTtcXFxcbiAgLS1wZi12Ni1jLWFsZXJ0LWdyb3VwLS1tLXRvYXN0LS1JbnNldEJsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS0yeGwpO1xcXFxuICAtLXBmLXY2LWMtYWxlcnQtZ3JvdXAtLW0tdG9hc3QtLUluc2V0SW5saW5lRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teGwpO1xcXFxuICAtLXBmLXY2LWMtYWxlcnQtZ3JvdXAtLW0tdG9hc3QtLU1heFdpZHRoOiAzNy41cmVtO1xcXFxuICAtLXBmLXY2LWMtYWxlcnQtZ3JvdXAtLW0tdG9hc3QtLVpJbmRleDogdmFyKC0tcGYtdC0tZ2xvYmFsLS16LWluZGV4LS0yeGwpO1xcXFxufVxcXFxuXFxcXG46aG9zdChbdG9hc3RdKSB7XFxcXG4gIHBvc2l0aW9uOiBmaXhlZDtcXFxcbiAgaW5zZXQtYmxvY2stc3RhcnQ6IHZhcigtLXBmLXY2LWMtYWxlcnQtZ3JvdXAtLW0tdG9hc3QtLUluc2V0QmxvY2tTdGFydCk7XFxcXG4gIGluc2V0LWlubGluZS1lbmQ6IHZhcigtLXBmLXY2LWMtYWxlcnQtZ3JvdXAtLW0tdG9hc3QtLUluc2V0SW5saW5lRW5kKTtcXFxcbiAgei1pbmRleDogdmFyKC0tcGYtdjYtYy1hbGVydC1ncm91cC0tbS10b2FzdC0tWkluZGV4KTtcXFxcbiAgd2lkdGg6IGNhbGMoMTAwJSAtIHZhcigtLXBmLXY2LWMtYWxlcnQtZ3JvdXAtLW0tdG9hc3QtLUluc2V0SW5saW5lRW5kKSAqIDIpO1xcXFxuICBtYXgtd2lkdGg6IHZhcigtLXBmLXY2LWMtYWxlcnQtZ3JvdXAtLW0tdG9hc3QtLU1heFdpZHRoKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QgOjpzbG90dGVkKCo6bm90KDpmaXJzdC1jaGlsZCkpLFxcXFxuOmhvc3QgOjpzbG90dGVkKC5wZi12Ni1jLWFsZXJ0LWdyb3VwX19pdGVtOm5vdCg6Zmlyc3QtY2hpbGQpKSB7XFxcXG4gIG1hcmdpbi1ibG9jay1zdGFydDogdmFyKC0tcGYtdjYtYy1hbGVydC1ncm91cF9faXRlbS0tTWFyZ2luQmxvY2tTdGFydCk7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sOGdDQUFnaEMsQ0FBQztBQUFFLElBQU8sNEJBQVE7OztBREF2bEM7QUFjQSw0QkFBQyxjQUFjLG1CQUFtQjtBQUMzQixJQUFNLGVBQU4sZUFBMkIsaUJBR2hDLGNBQUMsU0FBUyxFQUFFLE1BQU0sU0FBUyxTQUFTLEtBQUssQ0FBQyxJQUcxQyxtQkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsTUFBTSxXQUFXLGNBQWMsQ0FBQyxJQU5wQyxJQUFXO0FBQUEsRUFBdEM7QUFBQTtBQUFBO0FBSUwsdUJBQVMsUUFBUSxrQkFBakIsZ0JBQWlCLFNBQWpCO0FBR0EsdUJBQVMsYUFBYSxrQkFBdEIsaUJBQXNCLFNBQXRCO0FBc0NBLGlDQUFXLENBQUMsTUFBYTtBQUN2QixZQUFNLFFBQVEsRUFBRTtBQUNoQixVQUFJLE9BQU8sWUFBWSxlQUFlO0FBQ3BDLDhCQUFLLHlDQUFMLFdBQWtCO0FBQUEsTUFDcEI7QUFBQSxJQUNGO0FBQUE7QUFBQSxFQXpDQSxvQkFBb0I7QUFDbEIsVUFBTSxrQkFBa0I7QUFDeEIsU0FBSyxpQkFBaUIsU0FBUyxtQkFBSyxTQUFRO0FBQUEsRUFDOUM7QUFBQSxFQUVBLHVCQUF1QjtBQUNyQixVQUFNLHFCQUFxQjtBQUMzQixTQUFLLG9CQUFvQixTQUFTLG1CQUFLLFNBQVE7QUFBQSxFQUNqRDtBQUFBLEVBRUEsZUFBZTtBQUNiLFFBQUksS0FBSyxZQUFZO0FBQ25CLFdBQUssYUFBYSxhQUFhLFFBQVE7QUFDdkMsV0FBSyxhQUFhLGVBQWUsT0FBTztBQUFBLElBQzFDO0FBQUEsRUFDRjtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQSxFQUNUO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUFLQSxTQUFTLE9BQW9CO0FBQzNCLFFBQUksS0FBSyxPQUFPO0FBQ2QsWUFBTSxVQUFVLElBQUksZUFBZTtBQUNuQyxXQUFLLFlBQVksS0FBSztBQUN0Qiw0QkFBc0IsTUFBTTtBQUMxQixjQUFNLFVBQVUsT0FBTyxlQUFlO0FBQUEsTUFDeEMsQ0FBQztBQUFBLElBQ0gsT0FBTztBQUNMLFdBQUssWUFBWSxLQUFLO0FBQUEsSUFDeEI7QUFBQSxFQUNGO0FBcUJGO0FBaEVPO0FBSUk7QUFHQTtBQXNDVDtBQTdDSztBQW9ETCxpQkFBWSxTQUFDLE9BQW9CO0FBQy9CLE1BQUksQ0FBQyxNQUFPO0FBRVosTUFBSSxLQUFLLE9BQU87QUFDZCxVQUFNLFVBQVUsSUFBSSxlQUFlO0FBQ25DLFVBQU0saUJBQWlCLGlCQUFpQixNQUFNO0FBQzVDLFlBQU0sT0FBTztBQUFBLElBQ2YsR0FBRyxFQUFFLE1BQU0sS0FBSyxDQUFDO0FBQUEsRUFDbkIsT0FBTztBQUNMLFVBQU0sT0FBTztBQUFBLEVBQ2Y7QUFDRjtBQTNEQSw0QkFBUyxTQURULFlBSFcsY0FJRjtBQUdULDRCQUFTLGNBRFQsaUJBTlcsY0FPRjtBQVBFLGVBQU4sNENBRFAsMEJBQ2E7QUFDWCxjQURXLGNBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
