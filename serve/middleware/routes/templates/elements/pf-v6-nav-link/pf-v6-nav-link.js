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

// elements/pf-v6-nav-link/pf-v6-nav-link.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-nav-link/pf-v6-nav-link.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --pf-v6-c-nav__toggle--TranslateY: 0;\\n\\n  position: relative;\\n  display: flex;\\n  column-gap: var(--pf-v6-c-nav__link--ColumnGap, var(--pf-t--global--spacer--sm));\\n  align-items: baseline;\\n}\\n\\n#link {\\n  position: relative;\\n  display: flex;\\n  column-gap: inherit;\\n  align-items: inherit;\\n  padding-block-start: var(--pf-v6-c-nav__link--PaddingBlockStart, var(--pf-t--global--spacer--sm));\\n  padding-block-end: var(--pf-v6-c-nav__link--PaddingBlockEnd, var(--pf-t--global--spacer--sm));\\n  padding-inline-start: var(--pf-v6-c-nav__link--PaddingInlineStart, var(--pf-t--global--spacer--md));\\n  padding-inline-end: var(--pf-v6-c-nav__link--PaddingInlineEnd, var(--pf-t--global--spacer--md));\\n  color: var(--pf-v6-c-nav__link--Color, var(--pf-t--global--text--color--subtle));\\n  text-align: start;\\n  text-decoration: none;\\n  background-color: var(--pf-v6-c-nav__link--BackgroundColor, var(--pf-t--global--background--color--action--plain--default));\\n  border: none;\\n  border-radius: var(--pf-v6-c-nav__link--BorderRadius, var(--pf-t--global--border--radius--small));\\n  transition: background-color var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default),\\n              color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n  cursor: pointer;\\n  width: 100%;\\n  font: inherit;\\n}\\n\\n#link::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border: var(--pf-v6-c-nav__link--BorderWidth, var(--pf-t--global--border--width--action--plain--default)) solid var(--pf-v6-c-nav__link--BorderColor, var(--pf-t--global--border--color--high-contrast));\\n  border-radius: inherit;\\n}\\n\\n#link:hover,\\n#link:focus {\\n  --pf-v6-c-nav__link--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  color: var(--pf-v6-c-nav__link--hover--Color, var(--pf-t--global--text--color--regular));\\n  background-color: var(--pf-v6-c-nav__link--hover--BackgroundColor, var(--pf-t--global--background--color--action--plain--alt--hover));\\n}\\n\\n:host([current]) #link,\\n:host([current]) #link:hover {\\n  --pf-v6-c-nav__link--BorderWidth: var(--pf-t--global--border--width--action--plain--clicked);\\n  --pf-v6-c-nav__link-icon--Color: var(--pf-t--global--icon--color--regular);\\n  color: var(--pf-v6-c-nav__link--m-current--Color, var(--pf-t--global--text--color--regular));\\n  background-color: var(--pf-v6-c-nav__link--m-current--BackgroundColor, var(--pf-t--global--background--color--action--plain--alt--clicked));\\n}\\n\\n::slotted(svg) {\\n  color: var(--pf-v6-c-nav__link-icon--Color, var(--pf-t--global--icon--color--subtle));\\n}\\n\\n/* Toggle wrapper for expandable items */\\n#toggle {\\n  flex: none;\\n  align-self: start;\\n  margin-inline-start: auto;\\n  transform: translateY(var(--pf-v6-c-nav__toggle--TranslateY));\\n}\\n\\n/* Hide toggle wrapper when not expandable */\\n:host(:not([expandable])) #toggle {\\n  display: none;\\n}\\n\\n#toggle-icon {\\n  display: inline-flex;\\n  transition: transform var(--pf-t--global--motion--duration--icon--default, 0.2s) var(--pf-t--global--motion--timing-function--default, ease);\\n  transform: rotate(var(--pf-v6-c-nav__toggle-icon--Rotate, 0));\\n}\\n\\n#link:where([aria-expanded=\\"true\\"]) #toggle-icon {\\n  transform: rotate(var(--pf-v6-c-nav__item--m-expanded__toggle-icon--Rotate, 90deg));\\n}\\n\\n.pf-v6-svg {\\n  width: 1em;\\n  height: 1em;\\n  vertical-align: -0.125em;\\n}\\n"'));
var pf_v6_nav_link_default = s;

// elements/pf-v6-nav-link/pf-v6-nav-link.ts
var PfNavToggleEvent = class extends Event {
  expanded;
  constructor(expanded) {
    super("pf-nav-toggle", { bubbles: true });
    this.expanded = expanded;
  }
};
var _expanded_dec, _expandable_dec, _current_dec, _label_dec, _href_dec, _a, _PfV6NavLink_decorators, _init, _href, _label, _current, _expandable, _expanded, _PfV6NavLink_instances, onClick_fn, markCurrentIfMatches_fn, renderLinkContent_fn;
_PfV6NavLink_decorators = [customElement("pf-v6-nav-link")];
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
          <svg class="pf-v6-svg"
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
__publicField(PfV6NavLink, "styles", pf_v6_nav_link_default);
__runInitializers(_init, 1, PfV6NavLink);
export {
  PfNavToggleEvent,
  PfV6NavLink
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtbmF2LWxpbmsvcGYtdjYtbmF2LWxpbmsudHMiLCAibGl0LWNzczovaG9tZS9iZW5ueXAvRGV2ZWxvcGVyL2NlbS9zZXJ2ZS9lbGVtZW50cy9wZi12Ni1uYXYtbGluay9wZi12Ni1uYXYtbGluay5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwsIG5vdGhpbmcgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtbmF2LWxpbmsuY3NzJztcblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZm9yIG5hdmlnYXRpb24gdG9nZ2xlXG4gKi9cbmV4cG9ydCBjbGFzcyBQZk5hdlRvZ2dsZUV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBleHBhbmRlZDogYm9vbGVhbjtcbiAgY29uc3RydWN0b3IoZXhwYW5kZWQ6IGJvb2xlYW4pIHtcbiAgICBzdXBlcigncGYtbmF2LXRvZ2dsZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLmV4cGFuZGVkID0gZXhwYW5kZWQ7XG4gIH1cbn1cblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IE5hdmlnYXRpb24gTGlua1xuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBsaW5rIHRleHRcbiAqIEBzbG90IGljb24tc3RhcnQgLSBJY29uIGJlZm9yZSB0ZXh0XG4gKiBAc2xvdCBpY29uLWVuZCAtIEljb24gYWZ0ZXIgdGV4dFxuICpcbiAqIEBmaXJlcyB7UGZOYXZUb2dnbGVFdmVudH0gcGYtbmF2LXRvZ2dsZSAtIFdoZW4gZXhwYW5kYWJsZSBsaW5rIGlzIGNsaWNrZWRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ3BmLXY2LW5hdi1saW5rJylcbmV4cG9ydCBjbGFzcyBQZlY2TmF2TGluayBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgaHJlZj86IHN0cmluZztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBsYWJlbD86IHN0cmluZztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGN1cnJlbnQgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGV4cGFuZGFibGUgPSBmYWxzZTtcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGV4cGFuZGVkID0gZmFsc2U7XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLiNtYXJrQ3VycmVudElmTWF0Y2hlcygpO1xuICB9XG5cbiAgI29uQ2xpY2soZTogRXZlbnQpIHtcbiAgICBpZiAodGhpcy5leHBhbmRhYmxlKSB7XG4gICAgICBlLnByZXZlbnREZWZhdWx0KCk7XG4gICAgICB0aGlzLmRpc3BhdGNoRXZlbnQobmV3IFBmTmF2VG9nZ2xlRXZlbnQoIXRoaXMuZXhwYW5kZWQpKTtcbiAgICB9XG4gIH1cblxuICAjbWFya0N1cnJlbnRJZk1hdGNoZXMoKSB7XG4gICAgaWYgKHRoaXMuaHJlZiAmJiB0aGlzLmhyZWYgPT09IHdpbmRvdy5sb2NhdGlvbi5wYXRobmFtZSkge1xuICAgICAgdGhpcy5jdXJyZW50ID0gdHJ1ZTtcbiAgICB9XG4gIH1cblxuICAjcmVuZGVyTGlua0NvbnRlbnQoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8c2xvdCBuYW1lPVwiaWNvbi1zdGFydFwiPjwvc2xvdD5cbiAgICAgIDxzbG90Pjwvc2xvdD5cbiAgICAgIDxzbG90IG5hbWU9XCJpY29uLWVuZFwiPjwvc2xvdD5cbiAgICAgIDxzcGFuIGlkPVwidG9nZ2xlXCJcbiAgICAgICAgICAgIHBhcnQ9XCJ0b2dnbGVcIj5cbiAgICAgICAgPHNwYW4gaWQ9XCJ0b2dnbGUtaWNvblwiXG4gICAgICAgICAgICAgIHBhcnQ9XCJ0b2dnbGUtaWNvblwiPlxuICAgICAgICAgIDxzdmcgY2xhc3M9XCJwZi12Ni1zdmdcIlxuICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCAyNTYgNTEyXCJcbiAgICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiPlxuICAgICAgICAgICAgPHBhdGggZD1cIk0yMjQuMyAyNzNsLTEzNiAxMzZjLTkuNCA5LjQtMjQuNiA5LjQtMzMuOSAwbC0yMi42LTIyLjZjLTkuNC05LjQtOS40LTI0LjYgMC0zMy45bDk2LjQtOTYuNC05Ni40LTk2LjRjLTkuNC05LjQtOS40LTI0LjYgMC0zMy45TDU0LjMgMTAzYzkuNC05LjQgMjQuNi05LjQgMzMuOSAwbDEzNiAxMzZjOS41IDkuNCA5LjUgMjQuNi4xIDM0elwiPjwvcGF0aD5cbiAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgPC9zcGFuPlxuICAgICAgPC9zcGFuPlxuICAgIGA7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgaWYgKHRoaXMuaHJlZikge1xuICAgICAgcmV0dXJuIGh0bWxgXG4gICAgICAgIDxhIGlkPVwibGlua1wiXG4gICAgICAgICAgIHBhcnQ9XCJsaW5rXCJcbiAgICAgICAgICAgaHJlZj0ke3RoaXMuaHJlZn1cbiAgICAgICAgICAgYXJpYS1sYWJlbD0ke3RoaXMubGFiZWwgPz8gbm90aGluZ31cbiAgICAgICAgICAgYXJpYS1jdXJyZW50PSR7dGhpcy5jdXJyZW50ID8gJ3BhZ2UnIDogbm90aGluZ31cbiAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25DbGlja30+XG4gICAgICAgICAgJHt0aGlzLiNyZW5kZXJMaW5rQ29udGVudCgpfVxuICAgICAgICA8L2E+XG4gICAgICBgO1xuICAgIH1cbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxidXR0b24gaWQ9XCJsaW5rXCJcbiAgICAgICAgICAgICAgcGFydD1cImxpbmtcIlxuICAgICAgICAgICAgICB0eXBlPVwiYnV0dG9uXCJcbiAgICAgICAgICAgICAgYXJpYS1sYWJlbD0ke3RoaXMubGFiZWwgPz8gbm90aGluZ31cbiAgICAgICAgICAgICAgYXJpYS1leHBhbmRlZD0ke3RoaXMuZXhwYW5kYWJsZSA/IFN0cmluZyh0aGlzLmV4cGFuZGVkKSA6IG5vdGhpbmd9XG4gICAgICAgICAgICAgIEBjbGljaz0ke3RoaXMuI29uQ2xpY2t9PlxuICAgICAgICAke3RoaXMuI3JlbmRlckxpbmtDb250ZW50KCl9XG4gICAgICA8L2J1dHRvbj5cbiAgICBgO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ3BmLXY2LW5hdi1saW5rJzogUGZWNk5hdkxpbms7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcblxcXFxuICAtLXBmLXY2LWMtbmF2X190b2dnbGUtLVRyYW5zbGF0ZVk6IDA7XFxcXG5cXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBjb2x1bW4tZ2FwOiB2YXIoLS1wZi12Ni1jLW5hdl9fbGluay0tQ29sdW1uR2FwLCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pKTtcXFxcbiAgYWxpZ24taXRlbXM6IGJhc2VsaW5lO1xcXFxufVxcXFxuXFxcXG4jbGluayB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgY29sdW1uLWdhcDogaW5oZXJpdDtcXFxcbiAgYWxpZ24taXRlbXM6IGluaGVyaXQ7XFxcXG4gIHBhZGRpbmctYmxvY2stc3RhcnQ6IHZhcigtLXBmLXY2LWMtbmF2X19saW5rLS1QYWRkaW5nQmxvY2tTdGFydCwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKSk7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1wZi12Ni1jLW5hdl9fbGluay0tUGFkZGluZ0Jsb2NrRW5kLCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20pKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLXBmLXY2LWMtbmF2X19saW5rLS1QYWRkaW5nSW5saW5lU3RhcnQsIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCkpO1xcXFxuICBwYWRkaW5nLWlubGluZS1lbmQ6IHZhcigtLXBmLXY2LWMtbmF2X19saW5rLS1QYWRkaW5nSW5saW5lRW5kLCB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpKTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXY2LWMtbmF2X19saW5rLS1Db2xvciwgdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKSk7XFxcXG4gIHRleHQtYWxpZ246IHN0YXJ0O1xcXFxuICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXY2LWMtbmF2X19saW5rLS1CYWNrZ3JvdW5kQ29sb3IsIHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLWRlZmF1bHQpKTtcXFxcbiAgYm9yZGVyOiBub25lO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1wZi12Ni1jLW5hdl9fbGluay0tQm9yZGVyUmFkaXVzLCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tcmFkaXVzLS1zbWFsbCkpO1xcXFxuICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kLWNvbG9yIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tZGVmYXVsdCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLXRpbWluZy1mdW5jdGlvbi0tZGVmYXVsdCksXFxcXG4gICAgICAgICAgICAgIGNvbG9yIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS1kdXJhdGlvbi0tZmFkZS0tc2hvcnQpIHZhcigtLXBmLXQtLWdsb2JhbC0tbW90aW9uLS10aW1pbmctZnVuY3Rpb24tLWRlZmF1bHQpO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gIHdpZHRoOiAxMDAlO1xcXFxuICBmb250OiBpbmhlcml0O1xcXFxufVxcXFxuXFxcXG4jbGluazo6YWZ0ZXIge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIGluc2V0OiAwO1xcXFxuICBwb2ludGVyLWV2ZW50czogbm9uZTtcXFxcbiAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcIjtcXFxcbiAgYm9yZGVyOiB2YXIoLS1wZi12Ni1jLW5hdl9fbGluay0tQm9yZGVyV2lkdGgsIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tYWN0aW9uLS1wbGFpbi0tZGVmYXVsdCkpIHNvbGlkIHZhcigtLXBmLXY2LWMtbmF2X19saW5rLS1Cb3JkZXJDb2xvciwgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1oaWdoLWNvbnRyYXN0KSk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IGluaGVyaXQ7XFxcXG59XFxcXG5cXFxcbiNsaW5rOmhvdmVyLFxcXFxuI2xpbms6Zm9jdXMge1xcXFxuICAtLXBmLXY2LWMtbmF2X19saW5rLS1Cb3JkZXJXaWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1hY3Rpb24tLXBsYWluLS1ob3Zlcik7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi12Ni1jLW5hdl9fbGluay0taG92ZXItLUNvbG9yLCB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKSk7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXY2LWMtbmF2X19saW5rLS1ob3Zlci0tQmFja2dyb3VuZENvbG9yLCB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1hbHQtLWhvdmVyKSk7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtjdXJyZW50XSkgI2xpbmssXFxcXG46aG9zdChbY3VycmVudF0pICNsaW5rOmhvdmVyIHtcXFxcbiAgLS1wZi12Ni1jLW5hdl9fbGluay0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS13aWR0aC0tYWN0aW9uLS1wbGFpbi0tY2xpY2tlZCk7XFxcXG4gIC0tcGYtdjYtYy1uYXZfX2xpbmstaWNvbi0tQ29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdjYtYy1uYXZfX2xpbmstLW0tY3VycmVudC0tQ29sb3IsIHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpKTtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdjYtYy1uYXZfX2xpbmstLW0tY3VycmVudC0tQmFja2dyb3VuZENvbG9yLCB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1hbHQtLWNsaWNrZWQpKTtcXFxcbn1cXFxcblxcXFxuOjpzbG90dGVkKHN2Zykge1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdjYtYy1uYXZfX2xpbmstaWNvbi0tQ29sb3IsIHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLXN1YnRsZSkpO1xcXFxufVxcXFxuXFxcXG4vKiBUb2dnbGUgd3JhcHBlciBmb3IgZXhwYW5kYWJsZSBpdGVtcyAqL1xcXFxuI3RvZ2dsZSB7XFxcXG4gIGZsZXg6IG5vbmU7XFxcXG4gIGFsaWduLXNlbGY6IHN0YXJ0O1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiBhdXRvO1xcXFxuICB0cmFuc2Zvcm06IHRyYW5zbGF0ZVkodmFyKC0tcGYtdjYtYy1uYXZfX3RvZ2dsZS0tVHJhbnNsYXRlWSkpO1xcXFxufVxcXFxuXFxcXG4vKiBIaWRlIHRvZ2dsZSB3cmFwcGVyIHdoZW4gbm90IGV4cGFuZGFibGUgKi9cXFxcbjpob3N0KDpub3QoW2V4cGFuZGFibGVdKSkgI3RvZ2dsZSB7XFxcXG4gIGRpc3BsYXk6IG5vbmU7XFxcXG59XFxcXG5cXFxcbiN0b2dnbGUtaWNvbiB7XFxcXG4gIGRpc3BsYXk6IGlubGluZS1mbGV4O1xcXFxuICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gdmFyKC0tcGYtdC0tZ2xvYmFsLS1tb3Rpb24tLWR1cmF0aW9uLS1pY29uLS1kZWZhdWx0LCAwLjJzKSB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0LCBlYXNlKTtcXFxcbiAgdHJhbnNmb3JtOiByb3RhdGUodmFyKC0tcGYtdjYtYy1uYXZfX3RvZ2dsZS1pY29uLS1Sb3RhdGUsIDApKTtcXFxcbn1cXFxcblxcXFxuI2xpbms6d2hlcmUoW2FyaWEtZXhwYW5kZWQ9XFxcXFxcXCJ0cnVlXFxcXFxcXCJdKSAjdG9nZ2xlLWljb24ge1xcXFxuICB0cmFuc2Zvcm06IHJvdGF0ZSh2YXIoLS1wZi12Ni1jLW5hdl9faXRlbS0tbS1leHBhbmRlZF9fdG9nZ2xlLWljb24tLVJvdGF0ZSwgOTBkZWcpKTtcXFxcbn1cXFxcblxcXFxuLnBmLXY2LXN2ZyB7XFxcXG4gIHdpZHRoOiAxZW07XFxcXG4gIGhlaWdodDogMWVtO1xcXFxuICB2ZXJ0aWNhbC1hbGlnbjogLTAuMTI1ZW07XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7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
