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

// elements/cem-pf-v6-menu/cem-pf-v6-menu.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-menu/cem-pf-v6-menu.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  --cem-pf-v6-c-menu--RowGap: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu--Width: auto;\\n  --cem-pf-v6-c-menu--MinWidth: auto;\\n  --cem-pf-v6-c-menu--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --cem-pf-v6-c-menu--BoxShadow: var(--pf-t--global--box-shadow--md);\\n  --cem-pf-v6-c-menu--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-menu--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-menu--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-menu--OutlineOffset: calc(var(--pf-t--global--border--width--control--default) * -3);\\n  --cem-pf-v6-c-menu--ZIndex: var(--pf-t--global--z-index--sm);\\n  --cem-pf-v6-c-menu--TransitionDuration: 0s;\\n  --cem-pf-v6-c-menu--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n\\n  @media (prefers-reduced-motion: no-preference) {\\n    --cem-pf-v6-c-menu--TransitionDuration: var(--pf-t--global--motion--duration--fade--default);\\n  }\\n\\n  display: grid;\\n  row-gap: var(--cem-pf-v6-c-menu--RowGap);\\n  width: var(--cem-pf-v6-c-menu--Width);\\n  min-width: var(--cem-pf-v6-c-menu--MinWidth);\\n  padding-block-start: var(--cem-pf-v6-c-menu--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-menu--PaddingBlockEnd);\\n  overflow: hidden;\\n  color: var(--cem-pf-v6-c-menu--Color);\\n  background-color: var(--cem-pf-v6-c-menu--BackgroundColor);\\n  border: var(--cem-pf-v6-c-menu--BorderWidth) solid var(--cem-pf-v6-c-menu--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-menu--BorderRadius);\\n  box-shadow: var(--cem-pf-v6-c-menu--BoxShadow);\\n  transition-timing-function: var(--cem-pf-v6-c-menu--TransitionTimingFunction) !important;\\n  transition-duration: var(--cem-pf-v6-c-menu--TransitionDuration) !important;\\n}\\n\\nslot {\\n  display: contents;\\n}\\n"'));
var cem_pf_v6_menu_default = s;

// elements/cem-pf-v6-menu/cem-pf-v6-menu.ts
var _label_dec, _a, _PfV6Menu_decorators, _internals, _init, _label, _PfV6Menu_instances, onSlotChange_fn, _onKeydown, getMenuItems_fn, initializeTabindex_fn, focusItem_fn;
_PfV6Menu_decorators = [customElement("cem-pf-v6-menu")];
var PfV6Menu = class extends (_a = LitElement, _label_dec = [property()], _a) {
  constructor() {
    super();
    __privateAdd(this, _PfV6Menu_instances);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _label, __runInitializers(_init, 8, this, "")), __runInitializers(_init, 11, this);
    __privateAdd(this, _onKeydown, (event) => {
      const { key, target } = event;
      if (target.tagName !== "CEM-PF-V6-MENU-ITEM") return;
      const items = __privateMethod(this, _PfV6Menu_instances, getMenuItems_fn).call(this);
      const currentIndex = items.indexOf(target);
      switch (key) {
        case "ArrowDown":
          event.preventDefault();
          __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[currentIndex < items.length - 1 ? currentIndex + 1 : 0]);
          break;
        case "ArrowUp":
          event.preventDefault();
          __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[currentIndex > 0 ? currentIndex - 1 : items.length - 1]);
          break;
        case "Home":
          event.preventDefault();
          if (items.length > 0) __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[0]);
          break;
        case "End":
          event.preventDefault();
          if (items.length > 0) __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[items.length - 1]);
          break;
      }
    });
    __privateGet(this, _internals).role = "menu";
  }
  render() {
    return html`<slot @slotchange=${__privateMethod(this, _PfV6Menu_instances, onSlotChange_fn)}></slot>`;
  }
  updated(changed) {
    if (changed.has("label")) {
      __privateGet(this, _internals).ariaLabel = this.label || null;
    }
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("keydown", __privateGet(this, _onKeydown));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("keydown", __privateGet(this, _onKeydown));
  }
  /**
   * Focus the first menu item.
   * Called when dropdown opens.
   */
  focusFirstItem() {
    const items = __privateMethod(this, _PfV6Menu_instances, getMenuItems_fn).call(this);
    if (items.length > 0) {
      __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[0]);
    }
  }
};
_init = __decoratorStart(_a);
_internals = new WeakMap();
_label = new WeakMap();
_PfV6Menu_instances = new WeakSet();
onSlotChange_fn = function() {
  requestAnimationFrame(() => {
    __privateMethod(this, _PfV6Menu_instances, initializeTabindex_fn).call(this);
  });
};
_onKeydown = new WeakMap();
getMenuItems_fn = function() {
  const slot = this.shadowRoot?.querySelector("slot");
  if (!slot) return [];
  return slot.assignedElements().filter(
    (el) => el.tagName === "CEM-PF-V6-MENU-ITEM" && !el.disabled
  );
};
initializeTabindex_fn = function() {
  const items = __privateMethod(this, _PfV6Menu_instances, getMenuItems_fn).call(this);
  if (items.length === 0) return;
  items.forEach((item, index) => {
    item.setAttribute("tabindex", index === 0 ? "0" : "-1");
  });
};
focusItem_fn = function(item) {
  if (!item) return;
  const items = __privateMethod(this, _PfV6Menu_instances, getMenuItems_fn).call(this);
  items.forEach((i) => {
    i.setAttribute("tabindex", i === item ? "0" : "-1");
  });
  item.focus();
};
__decorateElement(_init, 4, "label", _label_dec, PfV6Menu, _label);
PfV6Menu = __decorateElement(_init, 0, "PfV6Menu", _PfV6Menu_decorators, PfV6Menu);
__publicField(PfV6Menu, "styles", cem_pf_v6_menu_default);
__runInitializers(_init, 1, PfV6Menu);
export {
  PfV6Menu
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LW1lbnUvY2VtLXBmLXY2LW1lbnUudHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tcGYtdjYtbWVudS9jZW0tcGYtdjYtbWVudS5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXBmLXY2LW1lbnUuY3NzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IE1lbnUgY29tcG9uZW50XG4gKlxuICogQ29udGFpbmVyIGZvciBjZW0tcGYtdjYtbWVudS1pdGVtIGVsZW1lbnRzLiBJbXBsZW1lbnRzIGtleWJvYXJkIG5hdmlnYXRpb25cbiAqIHVzaW5nIHJvdmluZyB0YWJpbmRleCBwYXR0ZXJuLlxuICpcbiAqIEBzbG90IC0gTWVudSBpdGVtcyAoY2VtLXBmLXY2LW1lbnUtaXRlbSBlbGVtZW50cylcbiAqXG4gKiBAYXR0ciB7c3RyaW5nfSBsYWJlbCAtIEFjY2Vzc2libGUgbGFiZWwgZm9yIHRoZSBtZW51XG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tcGYtdjYtbWVudScpXG5leHBvcnQgY2xhc3MgUGZWNk1lbnUgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjaW50ZXJuYWxzID0gdGhpcy5hdHRhY2hJbnRlcm5hbHMoKTtcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBsYWJlbCA9ICcnO1xuXG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCk7XG4gICAgdGhpcy4jaW50ZXJuYWxzLnJvbGUgPSAnbWVudSc7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgPHNsb3QgQHNsb3RjaGFuZ2U9JHt0aGlzLiNvblNsb3RDaGFuZ2V9Pjwvc2xvdD5gO1xuICB9XG5cbiAgdXBkYXRlZChjaGFuZ2VkOiBNYXA8c3RyaW5nLCB1bmtub3duPikge1xuICAgIGlmIChjaGFuZ2VkLmhhcygnbGFiZWwnKSkge1xuICAgICAgdGhpcy4jaW50ZXJuYWxzLmFyaWFMYWJlbCA9IHRoaXMubGFiZWwgfHwgbnVsbDtcbiAgICB9XG4gIH1cblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI29uS2V5ZG93bik7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI29uS2V5ZG93bik7XG4gIH1cblxuICAjb25TbG90Q2hhbmdlKCkge1xuICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICB0aGlzLiNpbml0aWFsaXplVGFiaW5kZXgoKTtcbiAgICB9KTtcbiAgfVxuXG4gICNvbktleWRvd24gPSAoZXZlbnQ6IEtleWJvYXJkRXZlbnQpID0+IHtcbiAgICBjb25zdCB7IGtleSwgdGFyZ2V0IH0gPSBldmVudDtcblxuICAgIGlmICgodGFyZ2V0IGFzIEVsZW1lbnQpLnRhZ05hbWUgIT09ICdDRU0tUEYtVjYtTUVOVS1JVEVNJykgcmV0dXJuO1xuXG4gICAgY29uc3QgaXRlbXMgPSB0aGlzLiNnZXRNZW51SXRlbXMoKTtcbiAgICBjb25zdCBjdXJyZW50SW5kZXggPSBpdGVtcy5pbmRleE9mKHRhcmdldCBhcyBFbGVtZW50KTtcblxuICAgIHN3aXRjaCAoa2V5KSB7XG4gICAgICBjYXNlICdBcnJvd0Rvd24nOlxuICAgICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICB0aGlzLiNmb2N1c0l0ZW0oaXRlbXNbY3VycmVudEluZGV4IDwgaXRlbXMubGVuZ3RoIC0gMSA/IGN1cnJlbnRJbmRleCArIDEgOiAwXSk7XG4gICAgICAgIGJyZWFrO1xuXG4gICAgICBjYXNlICdBcnJvd1VwJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zW2N1cnJlbnRJbmRleCA+IDAgPyBjdXJyZW50SW5kZXggLSAxIDogaXRlbXMubGVuZ3RoIC0gMV0pO1xuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnSG9tZSc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmIChpdGVtcy5sZW5ndGggPiAwKSB0aGlzLiNmb2N1c0l0ZW0oaXRlbXNbMF0pO1xuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnRW5kJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgaWYgKGl0ZW1zLmxlbmd0aCA+IDApIHRoaXMuI2ZvY3VzSXRlbShpdGVtc1tpdGVtcy5sZW5ndGggLSAxXSk7XG4gICAgICAgIGJyZWFrO1xuICAgIH1cbiAgfTtcblxuICAjZ2V0TWVudUl0ZW1zKCk6IEVsZW1lbnRbXSB7XG4gICAgY29uc3Qgc2xvdCA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3Rvcignc2xvdCcpIGFzIEhUTUxTbG90RWxlbWVudCB8IG51bGw7XG4gICAgaWYgKCFzbG90KSByZXR1cm4gW107XG4gICAgcmV0dXJuIHNsb3QuYXNzaWduZWRFbGVtZW50cygpLmZpbHRlcihcbiAgICAgIGVsID0+IGVsLnRhZ05hbWUgPT09ICdDRU0tUEYtVjYtTUVOVS1JVEVNJyAmJiAhKGVsIGFzIEhUTUxFbGVtZW50ICYgeyBkaXNhYmxlZD86IGJvb2xlYW4gfSkuZGlzYWJsZWRcbiAgICApO1xuICB9XG5cbiAgI2luaXRpYWxpemVUYWJpbmRleCgpIHtcbiAgICBjb25zdCBpdGVtcyA9IHRoaXMuI2dldE1lbnVJdGVtcygpO1xuICAgIGlmIChpdGVtcy5sZW5ndGggPT09IDApIHJldHVybjtcbiAgICBpdGVtcy5mb3JFYWNoKChpdGVtLCBpbmRleCkgPT4ge1xuICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3RhYmluZGV4JywgaW5kZXggPT09IDAgPyAnMCcgOiAnLTEnKTtcbiAgICB9KTtcbiAgfVxuXG4gICNmb2N1c0l0ZW0oaXRlbTogRWxlbWVudCkge1xuICAgIGlmICghaXRlbSkgcmV0dXJuO1xuICAgIGNvbnN0IGl0ZW1zID0gdGhpcy4jZ2V0TWVudUl0ZW1zKCk7XG4gICAgaXRlbXMuZm9yRWFjaChpID0+IHtcbiAgICAgIGkuc2V0QXR0cmlidXRlKCd0YWJpbmRleCcsIGkgPT09IGl0ZW0gPyAnMCcgOiAnLTEnKTtcbiAgICB9KTtcbiAgICAoaXRlbSBhcyBIVE1MRWxlbWVudCkuZm9jdXMoKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBGb2N1cyB0aGUgZmlyc3QgbWVudSBpdGVtLlxuICAgKiBDYWxsZWQgd2hlbiBkcm9wZG93biBvcGVucy5cbiAgICovXG4gIGZvY3VzRmlyc3RJdGVtKCkge1xuICAgIGNvbnN0IGl0ZW1zID0gdGhpcy4jZ2V0TWVudUl0ZW1zKCk7XG4gICAgaWYgKGl0ZW1zLmxlbmd0aCA+IDApIHtcbiAgICAgIHRoaXMuI2ZvY3VzSXRlbShpdGVtc1swXSk7XG4gICAgfVxuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1wZi12Ni1tZW51JzogUGZWNk1lbnU7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LS1Sb3dHYXA6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS0tV2lkdGg6IGF1dG87XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS0tTWluV2lkdGg6IGF1dG87XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS0tUGFkZGluZ0Jsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS0tUGFkZGluZ0Jsb2NrRW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtLUJhY2tncm91bmRDb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tZmxvYXRpbmctLWRlZmF1bHQpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtLUJveFNoYWRvdzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3gtc2hhZG93LS1tZCk7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWhpZ2gtY29udHJhc3QtLXJlZ3VsYXIpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtLUJvcmRlckNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWhpZ2gtY29udHJhc3QpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtLUJvcmRlclJhZGl1czogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXJhZGl1cy0tc21hbGwpO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtLU91dGxpbmVPZmZzZXQ6IGNhbGModmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1jb250cm9sLS1kZWZhdWx0KSAqIC0zKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tZW51LS1aSW5kZXg6IHZhcigtLXBmLXQtLWdsb2JhbC0tei1pbmRleC0tc20pO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1lbnUtLVRyYW5zaXRpb25EdXJhdGlvbjogMHM7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWVudS0tVHJhbnNpdGlvblRpbWluZ0Z1bmN0aW9uOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KTtcXFxcblxcXFxuICBAbWVkaWEgKHByZWZlcnMtcmVkdWNlZC1tb3Rpb246IG5vLXByZWZlcmVuY2UpIHtcXFxcbiAgICAtLWNlbS1wZi12Ni1jLW1lbnUtLVRyYW5zaXRpb25EdXJhdGlvbjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1kZWZhdWx0KTtcXFxcbiAgfVxcXFxuXFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIHJvdy1nYXA6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtLVJvd0dhcCk7XFxcXG4gIHdpZHRoOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LS1XaWR0aCk7XFxcXG4gIG1pbi13aWR0aDogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS0tTWluV2lkdGgpO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LS1QYWRkaW5nQmxvY2tTdGFydCk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICBvdmVyZmxvdzogaGlkZGVuO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbWVudS0tQ29sb3IpO1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LS1CYWNrZ3JvdW5kQ29sb3IpO1xcXFxuICBib3JkZXI6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtLUJvcmRlcldpZHRoKSBzb2xpZCB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LS1Cb3JkZXJDb2xvcik7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtLUJvcmRlclJhZGl1cyk7XFxcXG4gIGJveC1zaGFkb3c6IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtLUJveFNoYWRvdyk7XFxcXG4gIHRyYW5zaXRpb24tdGltaW5nLWZ1bmN0aW9uOiB2YXIoLS1jZW0tcGYtdjYtYy1tZW51LS1UcmFuc2l0aW9uVGltaW5nRnVuY3Rpb24pICFpbXBvcnRhbnQ7XFxcXG4gIHRyYW5zaXRpb24tZHVyYXRpb246IHZhcigtLWNlbS1wZi12Ni1jLW1lbnUtLVRyYW5zaXRpb25EdXJhdGlvbikgIWltcG9ydGFudDtcXFxcbn1cXFxcblxcXFxuc2xvdCB7XFxcXG4gIGRpc3BsYXk6IGNvbnRlbnRzO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLHFtRUFBdW1FLENBQUM7QUFBRSxJQUFPLHlCQUFROzs7QURBOXFFO0FBZ0JBLHdCQUFDLGNBQWMsZ0JBQWdCO0FBQ3hCLElBQU0sV0FBTixlQUF1QixpQkFLNUIsY0FBQyxTQUFTLElBTGtCLElBQVc7QUFBQSxFQVF2QyxjQUFjO0FBQ1osVUFBTTtBQVRIO0FBR0wsbUNBQWEsS0FBSyxnQkFBZ0I7QUFHbEMsdUJBQVMsUUFBUSxrQkFBakIsZ0JBQWlCLE1BQWpCO0FBaUNBLG1DQUFhLENBQUMsVUFBeUI7QUFDckMsWUFBTSxFQUFFLEtBQUssT0FBTyxJQUFJO0FBRXhCLFVBQUssT0FBbUIsWUFBWSxzQkFBdUI7QUFFM0QsWUFBTSxRQUFRLHNCQUFLLHND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
