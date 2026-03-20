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

// elements/pf-v6-menu/pf-v6-menu.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-menu/pf-v6-menu.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  --pf-v6-c-menu--RowGap: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-menu--Width: auto;\\n  --pf-v6-c-menu--MinWidth: auto;\\n  --pf-v6-c-menu--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-menu--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --pf-v6-c-menu--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --pf-v6-c-menu--BoxShadow: var(--pf-t--global--box-shadow--md);\\n  --pf-v6-c-menu--Color: var(--pf-t--global--text--color--regular);\\n  --pf-v6-c-menu--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --pf-v6-c-menu--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --pf-v6-c-menu--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --pf-v6-c-menu--OutlineOffset: calc(var(--pf-t--global--border--width--control--default) * -3);\\n  --pf-v6-c-menu--ZIndex: var(--pf-t--global--z-index--sm);\\n  --pf-v6-c-menu--TransitionDuration: 0s;\\n  --pf-v6-c-menu--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n\\n  @media (prefers-reduced-motion: no-preference) {\\n    --pf-v6-c-menu--TransitionDuration: var(--pf-t--global--motion--duration--fade--default);\\n  }\\n\\n  display: grid;\\n  row-gap: var(--pf-v6-c-menu--RowGap);\\n  width: var(--pf-v6-c-menu--Width);\\n  min-width: var(--pf-v6-c-menu--MinWidth);\\n  padding-block-start: var(--pf-v6-c-menu--PaddingBlockStart);\\n  padding-block-end: var(--pf-v6-c-menu--PaddingBlockEnd);\\n  overflow: hidden;\\n  color: var(--pf-v6-c-menu--Color);\\n  background-color: var(--pf-v6-c-menu--BackgroundColor);\\n  border: var(--pf-v6-c-menu--BorderWidth) solid var(--pf-v6-c-menu--BorderColor);\\n  border-radius: var(--pf-v6-c-menu--BorderRadius);\\n  box-shadow: var(--pf-v6-c-menu--BoxShadow);\\n  transition-timing-function: var(--pf-v6-c-menu--TransitionTimingFunction) !important;\\n  transition-duration: var(--pf-v6-c-menu--TransitionDuration) !important;\\n}\\n\\nslot {\\n  display: contents;\\n}\\n"'));
var pf_v6_menu_default = s;

// elements/pf-v6-menu/pf-v6-menu.ts
var _label_dec, _a, _PfV6Menu_decorators, _internals, _init, _label, _PfV6Menu_instances, onSlotChange_fn, _onKeydown, getMenuItems_fn, initializeTabindex_fn, focusItem_fn;
_PfV6Menu_decorators = [customElement("pf-v6-menu")];
var PfV6Menu = class extends (_a = LitElement, _label_dec = [property()], _a) {
  constructor() {
    super();
    __privateAdd(this, _PfV6Menu_instances);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _label, __runInitializers(_init, 8, this, "")), __runInitializers(_init, 11, this);
    __privateAdd(this, _onKeydown, (event) => {
      const { key, target } = event;
      if (target.tagName !== "PF-V6-MENU-ITEM") return;
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
    (el) => el.tagName === "PF-V6-MENU-ITEM" && !el.disabled
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
__publicField(PfV6Menu, "styles", pf_v6_menu_default);
__runInitializers(_init, 1, PfV6Menu);
export {
  PfV6Menu
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtbWVudS9wZi12Ni1tZW51LnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvcGYtdjYtbWVudS9wZi12Ni1tZW51LmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9wZi12Ni1tZW51LmNzcyc7XG5cbi8qKlxuICogUGF0dGVybkZseSB2NiBNZW51IGNvbXBvbmVudFxuICpcbiAqIENvbnRhaW5lciBmb3IgcGYtdjYtbWVudS1pdGVtIGVsZW1lbnRzLiBJbXBsZW1lbnRzIGtleWJvYXJkIG5hdmlnYXRpb25cbiAqIHVzaW5nIHJvdmluZyB0YWJpbmRleCBwYXR0ZXJuLlxuICpcbiAqIEBzbG90IC0gTWVudSBpdGVtcyAocGYtdjYtbWVudS1pdGVtIGVsZW1lbnRzKVxuICpcbiAqIEBhdHRyIHtzdHJpbmd9IGxhYmVsIC0gQWNjZXNzaWJsZSBsYWJlbCBmb3IgdGhlIG1lbnVcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ3BmLXY2LW1lbnUnKVxuZXhwb3J0IGNsYXNzIFBmVjZNZW51IGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgI2ludGVybmFscyA9IHRoaXMuYXR0YWNoSW50ZXJuYWxzKCk7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3IgbGFiZWwgPSAnJztcblxuICBjb25zdHJ1Y3RvcigpIHtcbiAgICBzdXBlcigpO1xuICAgIHRoaXMuI2ludGVybmFscy5yb2xlID0gJ21lbnUnO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYDxzbG90IEBzbG90Y2hhbmdlPSR7dGhpcy4jb25TbG90Q2hhbmdlfT48L3Nsb3Q+YDtcbiAgfVxuXG4gIHVwZGF0ZWQoY2hhbmdlZDogTWFwPHN0cmluZywgdW5rbm93bj4pIHtcbiAgICBpZiAoY2hhbmdlZC5oYXMoJ2xhYmVsJykpIHtcbiAgICAgIHRoaXMuI2ludGVybmFscy5hcmlhTGFiZWwgPSB0aGlzLmxhYmVsIHx8IG51bGw7XG4gICAgfVxuICB9XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNvbktleWRvd24pO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNvbktleWRvd24pO1xuICB9XG5cbiAgI29uU2xvdENoYW5nZSgpIHtcbiAgICByZXF1ZXN0QW5pbWF0aW9uRnJhbWUoKCkgPT4ge1xuICAgICAgdGhpcy4jaW5pdGlhbGl6ZVRhYmluZGV4KCk7XG4gICAgfSk7XG4gIH1cblxuICAjb25LZXlkb3duID0gKGV2ZW50OiBLZXlib2FyZEV2ZW50KSA9PiB7XG4gICAgY29uc3QgeyBrZXksIHRhcmdldCB9ID0gZXZlbnQ7XG5cbiAgICBpZiAoKHRhcmdldCBhcyBFbGVtZW50KS50YWdOYW1lICE9PSAnUEYtVjYtTUVOVS1JVEVNJykgcmV0dXJuO1xuXG4gICAgY29uc3QgaXRlbXMgPSB0aGlzLiNnZXRNZW51SXRlbXMoKTtcbiAgICBjb25zdCBjdXJyZW50SW5kZXggPSBpdGVtcy5pbmRleE9mKHRhcmdldCBhcyBFbGVtZW50KTtcblxuICAgIHN3aXRjaCAoa2V5KSB7XG4gICAgICBjYXNlICdBcnJvd0Rvd24nOlxuICAgICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICB0aGlzLiNmb2N1c0l0ZW0oaXRlbXNbY3VycmVudEluZGV4IDwgaXRlbXMubGVuZ3RoIC0gMSA/IGN1cnJlbnRJbmRleCArIDEgOiAwXSk7XG4gICAgICAgIGJyZWFrO1xuXG4gICAgICBjYXNlICdBcnJvd1VwJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zW2N1cnJlbnRJbmRleCA+IDAgPyBjdXJyZW50SW5kZXggLSAxIDogaXRlbXMubGVuZ3RoIC0gMV0pO1xuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnSG9tZSc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmIChpdGVtcy5sZW5ndGggPiAwKSB0aGlzLiNmb2N1c0l0ZW0oaXRlbXNbMF0pO1xuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnRW5kJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgaWYgKGl0ZW1zLmxlbmd0aCA+IDApIHRoaXMuI2ZvY3VzSXRlbShpdGVtc1tpdGVtcy5sZW5ndGggLSAxXSk7XG4gICAgICAgIGJyZWFrO1xuICAgIH1cbiAgfTtcblxuICAjZ2V0TWVudUl0ZW1zKCk6IEVsZW1lbnRbXSB7XG4gICAgY29uc3Qgc2xvdCA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3Rvcignc2xvdCcpIGFzIEhUTUxTbG90RWxlbWVudCB8IG51bGw7XG4gICAgaWYgKCFzbG90KSByZXR1cm4gW107XG4gICAgcmV0dXJuIHNsb3QuYXNzaWduZWRFbGVtZW50cygpLmZpbHRlcihcbiAgICAgIGVsID0+IGVsLnRhZ05hbWUgPT09ICdQRi1WNi1NRU5VLUlURU0nICYmICEoZWwgYXMgSFRNTEVsZW1lbnQgJiB7IGRpc2FibGVkPzogYm9vbGVhbiB9KS5kaXNhYmxlZFxuICAgICk7XG4gIH1cblxuICAjaW5pdGlhbGl6ZVRhYmluZGV4KCkge1xuICAgIGNvbnN0IGl0ZW1zID0gdGhpcy4jZ2V0TWVudUl0ZW1zKCk7XG4gICAgaWYgKGl0ZW1zLmxlbmd0aCA9PT0gMCkgcmV0dXJuO1xuICAgIGl0ZW1zLmZvckVhY2goKGl0ZW0sIGluZGV4KSA9PiB7XG4gICAgICBpdGVtLnNldEF0dHJpYnV0ZSgndGFiaW5kZXgnLCBpbmRleCA9PT0gMCA/ICcwJyA6ICctMScpO1xuICAgIH0pO1xuICB9XG5cbiAgI2ZvY3VzSXRlbShpdGVtOiBFbGVtZW50KSB7XG4gICAgaWYgKCFpdGVtKSByZXR1cm47XG4gICAgY29uc3QgaXRlbXMgPSB0aGlzLiNnZXRNZW51SXRlbXMoKTtcbiAgICBpdGVtcy5mb3JFYWNoKGkgPT4ge1xuICAgICAgaS5zZXRBdHRyaWJ1dGUoJ3RhYmluZGV4JywgaSA9PT0gaXRlbSA/ICcwJyA6ICctMScpO1xuICAgIH0pO1xuICAgIChpdGVtIGFzIEhUTUxFbGVtZW50KS5mb2N1cygpO1xuICB9XG5cbiAgLyoqXG4gICAqIEZvY3VzIHRoZSBmaXJzdCBtZW51IGl0ZW0uXG4gICAqIENhbGxlZCB3aGVuIGRyb3Bkb3duIG9wZW5zLlxuICAgKi9cbiAgZm9jdXNGaXJzdEl0ZW0oKSB7XG4gICAgY29uc3QgaXRlbXMgPSB0aGlzLiNnZXRNZW51SXRlbXMoKTtcbiAgICBpZiAoaXRlbXMubGVuZ3RoID4gMCkge1xuICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zWzBdKTtcbiAgICB9XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAncGYtdjYtbWVudSc6IFBmVjZNZW51O1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1Sb3dHYXA6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1XaWR0aDogYXV0bztcXFxcbiAgLS1wZi12Ni1jLW1lbnUtLU1pbldpZHRoOiBhdXRvO1xcXFxuICAtLXBmLXY2LWMtbWVudS0tUGFkZGluZ0Jsb2NrU3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1QYWRkaW5nQmxvY2tFbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1CYWNrZ3JvdW5kQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWZsb2F0aW5nLS1kZWZhdWx0KTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtLUJveFNoYWRvdzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3gtc2hhZG93LS1tZCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1Db2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1oaWdoLWNvbnRyYXN0LS1yZWd1bGFyKTtcXFxcbiAgLS1wZi12Ni1jLW1lbnUtLUJvcmRlckNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWhpZ2gtY29udHJhc3QpO1xcXFxuICAtLXBmLXY2LWMtbWVudS0tQm9yZGVyUmFkaXVzOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1PdXRsaW5lT2Zmc2V0OiBjYWxjKHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tY29udHJvbC0tZGVmYXVsdCkgKiAtMyk7XFxcXG4gIC0tcGYtdjYtYy1tZW51LS1aSW5kZXg6IHZhcigtLXBmLXQtLWdsb2JhbC0tei1pbmRleC0tc20pO1xcXFxuICAtLXBmLXY2LWMtbWVudS0tVHJhbnNpdGlvbkR1cmF0aW9uOiAwcztcXFxcbiAgLS1wZi12Ni1jLW1lbnUtLVRyYW5zaXRpb25UaW1pbmdGdW5jdGlvbjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCk7XFxcXG5cXFxcbiAgQG1lZGlhIChwcmVmZXJzLXJlZHVjZWQtbW90aW9uOiBuby1wcmVmZXJlbmNlKSB7XFxcXG4gICAgLS1wZi12Ni1jLW1lbnUtLVRyYW5zaXRpb25EdXJhdGlvbjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1kZWZhdWx0KTtcXFxcbiAgfVxcXFxuXFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIHJvdy1nYXA6IHZhcigtLXBmLXY2LWMtbWVudS0tUm93R2FwKTtcXFxcbiAgd2lkdGg6IHZhcigtLXBmLXY2LWMtbWVudS0tV2lkdGgpO1xcXFxuICBtaW4td2lkdGg6IHZhcigtLXBmLXY2LWMtbWVudS0tTWluV2lkdGgpO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi12Ni1jLW1lbnUtLVBhZGRpbmdCbG9ja1N0YXJ0KTtcXFxcbiAgcGFkZGluZy1ibG9jay1lbmQ6IHZhcigtLXBmLXY2LWMtbWVudS0tUGFkZGluZ0Jsb2NrRW5kKTtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXY2LWMtbWVudS0tQ29sb3IpO1xcXFxuICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1wZi12Ni1jLW1lbnUtLUJhY2tncm91bmRDb2xvcik7XFxcXG4gIGJvcmRlcjogdmFyKC0tcGYtdjYtYy1tZW51LS1Cb3JkZXJXaWR0aCkgc29saWQgdmFyKC0tcGYtdjYtYy1tZW51LS1Cb3JkZXJDb2xvcik7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXY2LWMtbWVudS0tQm9yZGVyUmFkaXVzKTtcXFxcbiAgYm94LXNoYWRvdzogdmFyKC0tcGYtdjYtYy1tZW51LS1Cb3hTaGFkb3cpO1xcXFxuICB0cmFuc2l0aW9uLXRpbWluZy1mdW5jdGlvbjogdmFyKC0tcGYtdjYtYy1tZW51LS1UcmFuc2l0aW9uVGltaW5nRnVuY3Rpb24pICFpbXBvcnRhbnQ7XFxcXG4gIHRyYW5zaXRpb24tZHVyYXRpb246IHZhcigtLXBmLXY2LWMtbWVudS0tVHJhbnNpdGlvbkR1cmF0aW9uKSAhaW1wb3J0YW50O1xcXFxufVxcXFxuXFxcXG5zbG90IHtcXFxcbiAgZGlzcGxheTogY29udGVudHM7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCOzs7QUNGekIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0saS9EQUFtL0QsQ0FBQztBQUFFLElBQU8scUJBQVE7OztBREExakU7QUFnQkEsd0JBQUMsY0FBYyxZQUFZO0FBQ3BCLElBQU0sV0FBTixlQUF1QixpQkFLNUIsY0FBQyxTQUFTLElBTGtCLElBQVc7QUFBQSxFQVF2QyxjQUFjO0FBQ1osVUFBTTtBQVRIO0FBR0wsbUNBQWEsS0FBSyxnQkFBZ0I7QUFHbEMsdUJBQVMsUUFBUSxrQkFBakIsZ0JBQWlCLE1BQWpCO0FBaUNBLG1DQUFhLENBQUMsVUFBeUI7QUFDckMsWUFBTSxFQUFFLEtBQUssT0FBTyxJQUFJO0FBRXhCLFVBQUssT0FBbUIsWUFBWSxrQkFBbUI7QUFFdkQsWUFBTSxRQUFRLHNCQUFLLHNDQUFMO0FBQ2QsWUFBTSxlQUFlLE1BQU0sUUFBUSxNQUFpQjtBQUVwRCxjQUFRLEtBQUs7QUFBQSxRQUNYLEtBQUs7QUFDSCxnQkFBTSxlQUFlO0FBQ3JCLGdDQUFLLG1DQUFMLFdBQWdCLE1BQU0sZUFBZSxNQUFNLFNBQVMsSUFB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
