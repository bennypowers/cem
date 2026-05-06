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

// elements/cem-pf-v6-nav-link/cem-pf-v6-nav-link.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/cem-pf-v6-nav-link/cem-pf-v6-nav-link.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-nav__toggle--TranslateY: 0;\\n\\n  position: relative;\\n  display: flex;\\n  column-gap: var(--cem-pf-v6-c-nav__link--ColumnGap, var(--pf-t--global--spacer--sm));\\n  align-items: baseline;\\n}\\n\\n#link {\\n  position: relative;\\n  display: flex;\\n  column-gap: inherit;\\n  align-items: inherit;\\n  padding-block-start: var(--cem-pf-v6-c-nav__link--PaddingBlockStart, var(--pf-t--global--spacer--sm));\\n  padding-block-end: var(--cem-pf-v6-c-nav__link--PaddingBlockEnd, var(--pf-t--global--spacer--sm));\\n  padding-inline-start: var(--cem-pf-v6-c-nav__link--PaddingInlineStart, var(--pf-t--global--spacer--md));\\n  padding-inline-end: var(--cem-pf-v6-c-nav__link--PaddingInlineEnd, var(--pf-t--global--spacer--md));\\n  color: var(--cem-pf-v6-c-nav__link--Color, var(--pf-t--global--text--color--subtle));\\n  text-align: start;\\n  text-decoration: none;\\n  background-color: var(--cem-pf-v6-c-nav__link--BackgroundColor, var(--pf-t--global--background--color--action--plain--default));\\n  border: none;\\n  border-radius: var(--cem-pf-v6-c-nav__link--BorderRadius, var(--pf-t--global--border--radius--small));\\n  transition: background-color var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default),\\n              color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n  cursor: pointer;\\n  width: 100%;\\n  font: inherit;\\n}\\n\\n#link::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border: var(--cem-pf-v6-c-nav__link--BorderWidth, var(--pf-t--global--border--width--action--plain--default)) solid var(--cem-pf-v6-c-nav__link--BorderColor, var(--pf-t--global--border--color--high-contrast));\\n  border-radius: inherit;\\n}\\n\\n#link:hover,\\n#link:focus {\\n  --cem-pf-v6-c-nav__link--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  color: var(--cem-pf-v6-c-nav__link--hover--Color, var(--pf-t--global--text--color--regular));\\n  background-color: var(--cem-pf-v6-c-nav__link--hover--BackgroundColor, var(--pf-t--global--background--color--action--plain--alt--hover));\\n}\\n\\n:host([current]) #link,\\n:host([current]) #link:hover {\\n  --cem-pf-v6-c-nav__link--BorderWidth: var(--pf-t--global--border--width--action--plain--clicked);\\n  --cem-pf-v6-c-nav__link-icon--Color: var(--pf-t--global--icon--color--regular);\\n  color: var(--cem-pf-v6-c-nav__link--m-current--Color, var(--pf-t--global--text--color--regular));\\n  background-color: var(--cem-pf-v6-c-nav__link--m-current--BackgroundColor, var(--pf-t--global--background--color--action--plain--alt--clicked));\\n}\\n\\n::slotted(svg) {\\n  color: var(--cem-pf-v6-c-nav__link-icon--Color, var(--pf-t--global--icon--color--subtle));\\n}\\n\\n/* Toggle wrapper for expandable items */\\n#toggle {\\n  flex: none;\\n  align-self: start;\\n  margin-inline-start: auto;\\n  transform: translateY(var(--cem-pf-v6-c-nav__toggle--TranslateY));\\n}\\n\\n/* Hide toggle wrapper when not expandable */\\n:host(:not([expandable])) #toggle {\\n  display: none;\\n}\\n\\n#toggle-icon {\\n  display: inline-flex;\\n  transition: transform var(--pf-t--global--motion--duration--icon--default, 0.2s) var(--pf-t--global--motion--timing-function--default, ease);\\n  transform: rotate(var(--cem-pf-v6-c-nav__toggle-icon--Rotate, 0));\\n}\\n\\n#link:where([aria-expanded=\\"true\\"]) #toggle-icon {\\n  transform: rotate(var(--cem-pf-v6-c-nav__item--m-expanded__toggle-icon--Rotate, 90deg));\\n}\\n\\n.cem-pf-v6-svg {\\n  width: 1em;\\n  height: 1em;\\n  vertical-align: -0.125em;\\n}\\n"'));
var cem_pf_v6_nav_link_default = s;

// elements/cem-pf-v6-nav-link/cem-pf-v6-nav-link.ts
var PfNavToggleEvent = class extends Event {
  expanded;
  constructor(expanded) {
    super("pf-nav-toggle", { bubbles: true });
    this.expanded = expanded;
  }
};
var _expanded_dec, _expandable_dec, _current_dec, _label_dec, _href_dec, _a, _PfV6NavLink_decorators, _init, _href, _label, _current, _expandable, _expanded, _PfV6NavLink_instances, onClick_fn, markCurrentIfMatches_fn, renderLinkContent_fn;
_PfV6NavLink_decorators = [customElement("cem-pf-v6-nav-link")];
var PfV6NavLink = class extends (_a = LitElement, _href_dec = [property({ reflect: true })], _label_dec = [property()], _current_dec = [property({ type: Boolean, reflect: true })], _expandable_dec = [property({ type: Boolean, reflect: true })], _expanded_dec = [property({ type: Boolean, reflect: true })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6NavLink_instances);
    __privateAdd(this, _href, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
    __privateAdd(this, _label, __runInitializers(_init, 12, this)), __runInitializers(_init, 15, this);
    __privateAdd(this, _current, __runInitializers(_init, 16, this, false)), __runInitializers(_init, 19, this);
    __privateAdd(this, _expandable, __runInitializers(_init, 20, this, false)), __runInitializers(_init, 23, this);
    __privateAdd(this, _expanded, __runInitializers(_init, 24, this, false)), __runInitializers(_init, 27, this);
  }
  connectedCallback() {
    super.connectedCallback();
    __privateMethod(this, _PfV6NavLink_instances, markCurrentIfMatches_fn).call(this);
  }
  render() {
    if (this.href) {
      return html`
        <a id="link"
           part="link"
           href=${this.href}
           aria-label=${this.label ?? nothing}
           aria-current=${this.current ? "page" : nothing}
           @click=${__privateMethod(this, _PfV6NavLink_instances, onClick_fn)}>
          ${__privateMethod(this, _PfV6NavLink_instances, renderLinkContent_fn).call(this)}
        </a>
      `;
    }
    return html`
      <button id="link"
              part="link"
              type="button"
              aria-label=${this.label ?? nothing}
              aria-expanded=${this.expandable ? String(this.expanded) : nothing}
              @click=${__privateMethod(this, _PfV6NavLink_instances, onClick_fn)}>
        ${__privateMethod(this, _PfV6NavLink_instances, renderLinkContent_fn).call(this)}
      </button>
    `;
  }
};
_init = __decoratorStart(_a);
_href = new WeakMap();
_label = new WeakMap();
_current = new WeakMap();
_expandable = new WeakMap();
_expanded = new WeakMap();
_PfV6NavLink_instances = new WeakSet();
onClick_fn = function(e) {
  if (this.expandable) {
    e.preventDefault();
    this.dispatchEvent(new PfNavToggleEvent(!this.expanded));
  }
};
markCurrentIfMatches_fn = function() {
  if (this.href && this.href === window.location.pathname) {
    this.current = true;
  }
};
renderLinkContent_fn = function() {
  return html`
      <slot name="icon-start"></slot>
      <slot></slot>
      <slot name="icon-end"></slot>
      <span id="toggle"
            part="toggle">
        <span id="toggle-icon"
              part="toggle-icon">
          <svg class="cem-pf-v6-svg"
               viewBox="0 0 256 512"
               fill="currentColor"
               role="presentation">
            <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
          </svg>
        </span>
      </span>
    `;
};
__decorateElement(_init, 4, "href", _href_dec, PfV6NavLink, _href);
__decorateElement(_init, 4, "label", _label_dec, PfV6NavLink, _label);
__decorateElement(_init, 4, "current", _current_dec, PfV6NavLink, _current);
__decorateElement(_init, 4, "expandable", _expandable_dec, PfV6NavLink, _expandable);
__decorateElement(_init, 4, "expanded", _expanded_dec, PfV6NavLink, _expanded);
PfV6NavLink = __decorateElement(_init, 0, "PfV6NavLink", _PfV6NavLink_decorators, PfV6NavLink);
__publicField(PfV6NavLink, "styles", cem_pf_v6_nav_link_default);
__runInitializers(_init, 1, PfV6NavLink);
export {
  PfNavToggleEvent,
  PfV6NavLink
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LW5hdi1saW5rL2NlbS1wZi12Ni1uYXYtbGluay50cyIsICJsaXQtY3NzOmVsZW1lbnRzL2NlbS1wZi12Ni1uYXYtbGluay9jZW0tcGYtdjYtbmF2LWxpbmsuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sLCBub3RoaW5nIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1wZi12Ni1uYXYtbGluay5jc3MnO1xuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmb3IgbmF2aWdhdGlvbiB0b2dnbGVcbiAqL1xuZXhwb3J0IGNsYXNzIFBmTmF2VG9nZ2xlRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIGV4cGFuZGVkOiBib29sZWFuO1xuICBjb25zdHJ1Y3RvcihleHBhbmRlZDogYm9vbGVhbikge1xuICAgIHN1cGVyKCdwZi1uYXYtdG9nZ2xlJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMuZXhwYW5kZWQgPSBleHBhbmRlZDtcbiAgfVxufVxuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgTmF2aWdhdGlvbiBMaW5rXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGxpbmsgdGV4dFxuICogQHNsb3QgaWNvbi1zdGFydCAtIEljb24gYmVmb3JlIHRleHRcbiAqIEBzbG90IGljb24tZW5kIC0gSWNvbiBhZnRlciB0ZXh0XG4gKlxuICogQGZpcmVzIHtQZk5hdlRvZ2dsZUV2ZW50fSBwZi1uYXYtdG9nZ2xlIC0gV2hlbiBleHBhbmRhYmxlIGxpbmsgaXMgY2xpY2tlZFxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXBmLXY2LW5hdi1saW5rJylcbmV4cG9ydCBjbGFzcyBQZlY2TmF2TGluayBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgaHJlZj86IHN0cmluZztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBsYWJlbD86IHN0cmluZztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGN1cnJlbnQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGV4cGFuZGFibGUgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGV4cGFuZGVkID0gZmFsc2U7XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLiNtYXJrQ3VycmVudElmTWF0Y2hlcygpO1xuICB9XG5cbiAgI29uQ2xpY2soZTogRXZlbnQpIHtcbiAgICBpZiAodGhpcy5leHBhbmRhYmxlKSB7XG4gICAgICBlLnByZXZlbnREZWZhdWx0KCk7XG4gICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IFBmTmF2VG9nZ2xlRXZlbnQoIXRoaXMuZXhwYW5kZWQpKTtcbiAgICB9XG4gIH1cblxuICAjbWFya0N1cnJlbnRJZk1hdGNoZXMoKSB7XG4gICAgaWYgKHRoaXMuaHJlZiAmJiB0aGlzLmhyZWYgPT09IHdpbmRvdy5sb2NhdGlvbi5wYXRobmFtZSkge1xuICAgICAgdGhpcy5jdXJyZW50ID0gdHJ1ZTtcbiAgICB9XG4gIH1cblxuICAjcmVuZGVyTGlua0NvbnRlbnQoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8c2xvdCBuYW1lPVwiaWNvbi1zdGFydFwiPjwvc2xvdD5cbiAgICAgIDxzbG90Pjwvc2xvdD5cbiAgICAgIDxzbG90IG5hbWU9XCJpY29uLWVuZFwiPjwvc2xvdD5cbiAgICAgIDxzcGFuIGlkPVwidG9nZ2xlXCJcbiAgICAgICAgICAgIHBhcnQ9XCJ0b2dnbGVcIj5cbiAgICAgICAgPHNwYW4gaWQ9XCJ0b2dnbGUtaWNvblwiXG4gICAgICAgICAgICAgIHBhcnQ9XCJ0b2dnbGUtaWNvblwiPlxuICAgICAgICAgIDxzdmcgY2xhc3M9XCJjZW0tcGYtdjYtc3ZnXCJcbiAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMjU2IDUxMlwiXG4gICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIj5cbiAgICAgICAgICAgIDxwYXRoIGQ9XCJNMjI0LjMgMjczbC0xMzYgMTM2Yy05LjQgOS40LTI0LjYgOS40LTMzLjkgMGwtMjIuNi0yMi42Yy05LjQtOS40LTkuNC0yNC42IDAtMzMuOWw5Ni40LTk2LjQtOTYuNC05Ni40Yy05LjQtOS40LTkuNC0yNC42IDAtMzMuOUw1NC4zIDEwM2M5LjQtOS40IDI0LjYtOS40IDMzLjkgMGwxMzYgMTM2YzkuNSA5LjQgOS41IDI0LjYuMSAzNHpcIj48L3BhdGg+XG4gICAgICAgICAgPC9zdmc+XG4gICAgICAgIDwvc3Bhbj5cbiAgICAgIDwvc3Bhbj5cbiAgICBgO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIGlmICh0aGlzLmhyZWYpIHtcbiAgICAgIHJldHVybiBodG1sYFxuICAgICAgICA8YSBpZD1cImxpbmtcIlxuICAgICAgICAgICBwYXJ0PVwibGlua1wiXG4gICAgICAgICAgIGhyZWY9JHt0aGlzLmhyZWZ9XG4gICAgICAgICAgIGFyaWEtbGFiZWw9JHt0aGlzLmxhYmVsID8/IG5vdGhpbmd9XG4gICAgICAgICAgIGFyaWEtY3VycmVudD0ke3RoaXMuY3VycmVudCA/ICdwYWdlJyA6IG5vdGhpbmd9XG4gICAgICAgICAgIEBjbGljaz0ke3RoaXMuI29uQ2xpY2t9PlxuICAgICAgICAgICR7dGhpcy4jcmVuZGVyTGlua0NvbnRlbnQoKX1cbiAgICAgICAgPC9hPlxuICAgICAgYDtcbiAgICB9XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8YnV0dG9uIGlkPVwibGlua1wiXG4gICAgICAgICAgICAgIHBhcnQ9XCJsaW5rXCJcbiAgICAgICAgICAgICAgdHlwZT1cImJ1dHRvblwiXG4gICAgICAgICAgICAgIGFyaWEtbGFiZWw9JHt0aGlzLmxhYmVsID8/IG5vdGhpbmd9XG4gICAgICAgICAgICAgIGFyaWEtZXhwYW5kZWQ9JHt0aGlzLmV4cGFuZGFibGUgPyBTdHJpbmcodGhpcy5leHBhbmRlZCkgOiBub3RoaW5nfVxuICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLiNvbkNsaWNrfT5cbiAgICAgICAgJHt0aGlzLiNyZW5kZXJMaW5rQ29udGVudCgpfVxuICAgICAgPC9idXR0b24+XG4gICAgYDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcGYtdjYtbmF2LWxpbmsnOiBQZlY2TmF2TGluaztcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuXFxcXG4gIC0tY2VtLXBmLXY2LWMtbmF2X190b2dnbGUtLVRyYW5zbGF0ZVk6IDA7XFxcXG5cXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBjb2x1bW4tZ2FwOiB2YXIoLS1jZW0tcGYtdjYtYy1uYXZfX2xpbmstLUNvbHVtbkdhcCwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKSk7XFxcXG4gIGFsaWduLWl0ZW1zOiBiYXNlbGluZTtcXFxcbn1cXFxcblxcXFxuI2xpbmsge1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGNvbHVtbi1nYXA6IGluaGVyaXQ7XFxcXG4gIGFsaWduLWl0ZW1zOiBpbmhlcml0O1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1jZW0tcGYtdjYtYy1uYXZfX2xpbmstLVBhZGRpbmdCbG9ja1N0YXJ0LCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pKTtcXFxcbiAgcGFkZGluZy1ibG9jay1lbmQ6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fbGluay0tUGFkZGluZ0Jsb2NrRW5kLCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fbGluay0tUGFkZGluZ0lubGluZVN0YXJ0LCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtZW5kOiB2YXIoLS1jZW0tcGYtdjYtYy1uYXZfX2xpbmstLVBhZGRpbmdJbmxpbmVFbmQsIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCkpO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbmF2X19saW5rLS1Db2xvciwgdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKSk7XFxcXG4gIHRleHQtYWxpZ246IHN0YXJ0O1xcXFxuICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fbGluay0tQmFja2dyb3VuZENvbG9yLCB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1kZWZhdWx0KSk7XFxcXG4gIGJvcmRlcjogbm9uZTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tY2VtLXBmLXY2LWMtbmF2X19saW5rLS1Cb3JkZXJSYWRpdXMsIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXNtYWxsKSk7XFxcXG4gIHRyYW5zaXRpb246IGJhY2tncm91bmQtY29sb3IgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1kZWZhdWx0KSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0KSxcXFxcbiAgICAgICAgICAgICAgY29sb3IgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1mYWRlLS1zaG9ydCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCk7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG4gIGZvbnQ6IGluaGVyaXQ7XFxcXG59XFxcXG5cXFxcbiNsaW5rOjphZnRlciB7XFxcXG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXFxcbiAgaW5zZXQ6IDA7XFxcXG4gIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxuICBjb250ZW50OiBcXFxcXFxcIlxcXFxcXFwiO1xcXFxuICBib3JkZXI6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fbGluay0tQm9yZGVyV2lkdGgsIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tYWN0aW9uLS1wbGFpbi0tZGVmYXVsdCkpIHNvbGlkIHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fbGluay0tQm9yZGVyQ29sb3IsIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0taGlnaC1jb250cmFzdCkpO1xcXFxuICBib3JkZXItcmFkaXVzOiBpbmhlcml0O1xcXFxufVxcXFxuXFxcXG4jbGluazpob3ZlcixcXFxcbiNsaW5rOmZvY3VzIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy1uYXZfX2xpbmstLUJvcmRlcldpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWFjdGlvbi0tcGxhaW4tLWhvdmVyKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fbGluay0taG92ZXItLUNvbG9yLCB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKSk7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fbGluay0taG92ZXItLUJhY2tncm91bmRDb2xvciwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0tYWx0LS1ob3ZlcikpO1xcXFxufVxcXFxuXFxcXG46aG9zdChbY3VycmVudF0pICNsaW5rLFxcXFxuOmhvc3QoW2N1cnJlbnRdKSAjbGluazpob3ZlciB7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbmF2X19saW5rLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1hY3Rpb24tLXBsYWluLS1jbGlja2VkKTtcXFxcbiAgLS1jZW0tcGYtdjYtYy1uYXZfX2xpbmstaWNvbi0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtbmF2X19saW5rLS1tLWN1cnJlbnQtLUNvbG9yLCB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKSk7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fbGluay0tbS1jdXJyZW50LS1CYWNrZ3JvdW5kQ29sb3IsIHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWFsdC0tY2xpY2tlZCkpO1xcXFxufVxcXFxuXFxcXG46OnNsb3R0ZWQoc3ZnKSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tcGYtdjYtYy1uYXZfX2xpbmstaWNvbi0tQ29sb3IsIHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLXN1YnRsZSkpO1xcXFxufVxcXFxuXFxcXG4vKiBUb2dnbGUgd3JhcHBlciBmb3IgZXhwYW5kYWJsZSBpdGVtcyAqL1xcXFxuI3RvZ2dsZSB7XFxcXG4gIGZsZXg6IG5vbmU7XFxcXG4gIGFsaWduLXNlbGY6IHN0YXJ0O1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiBhdXRvO1xcXFxuICB0cmFuc2Zvcm06IHRyYW5zbGF0ZVkodmFyKC0tY2VtLXBmLXY2LWMtbmF2X190b2dnbGUtLVRyYW5zbGF0ZVkpKTtcXFxcbn1cXFxcblxcXFxuLyogSGlkZSB0b2dnbGUgd3JhcHBlciB3aGVuIG5vdCBleHBhbmRhYmxlICovXFxcXG46aG9zdCg6bm90KFtleHBhbmRhYmxlXSkpICN0b2dnbGUge1xcXFxuICBkaXNwbGF5OiBub25lO1xcXFxufVxcXFxuXFxcXG4jdG9nZ2xlLWljb24ge1xcXFxuICBkaXNwbGF5OiBpbmxpbmUtZmxleDtcXFxcbiAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0taWNvbi0tZGVmYXVsdCwgMC4ycykgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCwgZWFzZSk7XFxcXG4gIHRyYW5zZm9ybTogcm90YXRlKHZhcigtLWNlbS1wZi12Ni1jLW5hdl9fdG9nZ2xlLWljb24tLVJvdGF0ZSwgMCkpO1xcXFxufVxcXFxuXFxcXG4jbGluazp3aGVyZShbYXJpYS1leHBhbmRlZD1cXFxcXFxcInRydWVcXFxcXFxcIl0pICN0b2dnbGUtaWNvbiB7XFxcXG4gIHRyYW5zZm9ybTogcm90YXRlKHZhcigtLWNlbS1wZi12Ni1jLW5hdl9faXRlbS0tbS1leHBhbmRlZF9fdG9nZ2xlLWljb24tLVJvdGF0ZSwgOTBkZWcpKTtcXFxcbn1cXFxcblxcXFxuLmNlbS1wZi12Ni1zdmcge1xcXFxuICB3aWR0aDogMWVtO1xcXFxuICBoZWlnaHQ6IDFlbTtcXFxcbiAgdmVydGljYWwtYWxpZ246IC0wLjEyNWVtO1xcXFxufVxcXFxu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
