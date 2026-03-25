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

// elements/pf-v6-masthead/pf-v6-masthead.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";

// lit-css:elements/pf-v6-masthead/pf-v6-masthead.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  --pf-v6-c-masthead--ColumnGap: var(--pf-v6-c-masthead--m-display-stack--ColumnGap);\\n  --pf-v6-c-masthead--GridTemplateColumns: var(--pf-v6-c-masthead--m-display-stack--GridTemplateColumns);\\n  --pf-v6-c-masthead__toggle--GridColumn: var(--pf-v6-c-masthead--m-display-stack__toggle--GridColumn);\\n  --pf-v6-c-masthead__brand--GridColumn: var(--pf-v6-c-masthead--m-display-stack__brand--GridColumn);\\n  --pf-v6-c-masthead__brand--Order: var(--pf-v6-c-masthead--m-display-stack__brand--Order);\\n  --pf-v6-c-masthead__brand--PaddingBlockEnd: var(--pf-v6-c-masthead--m-display-stack__brand--PaddingBlockEnd);\\n  --pf-v6-c-masthead__brand--BorderBlockEnd: var(--pf-v6-c-masthead--m-display-stack__brand--BorderBlockEnd);\\n  --pf-v6-c-masthead__main--GridColumn: unset;\\n  --pf-v6-c-masthead__content--GridColumn: var(--pf-v6-c-masthead--m-display-stack__content--GridColumn);\\n  --pf-v6-c-masthead__content--Order: var(--pf-v6-c-masthead--m-display-stack__content--Order);\\n  --pf-v6-c-masthead__main--Display: var(--pf-v6-c-masthead--m-display-stack__main--Display);\\n  --pf-v6-c-masthead__main--ColumnGap: var(--pf-v6-c-masthead--m-display-stack__main--ColumnGap);\\n\\n  position: relative;\\n  z-index: var(--pf-v6-c-page--c-masthead--ZIndex);\\n  grid-area: header;\\n  display: grid;\\n  grid-template-columns: var(--pf-v6-c-masthead--GridTemplateColumns);\\n  row-gap: var(--pf-v6-c-masthead--RowGap);\\n  column-gap: var(--pf-v6-c-masthead--ColumnGap);\\n  align-items: start;\\n  min-width: 0;\\n  padding-block-start: var(--pf-v6-c-masthead--PaddingBlockStart, var(--pf-v6-c-masthead--PaddingBlock));\\n  padding-block-end: var(--pf-v6-c-masthead--PaddingBlockEnd, var(--pf-v6-c-masthead--PaddingBlock));\\n  padding-inline-start: var(--pf-v6-c-masthead--PaddingInlineStart, var(--pf-v6-c-masthead--PaddingInline));\\n  padding-inline-end: var(--pf-v6-c-masthead--PaddingInlineEnd, var(--pf-v6-c-masthead--PaddingInline));\\n  background-color: var(--pf-v6-c-masthead--BackgroundColor);\\n}\\n\\n#main {\\n  display: var(--pf-v6-c-masthead__main--Display);\\n  grid-column: var(--pf-v6-c-masthead__main--GridColumn);\\n  column-gap: var(--pf-v6-c-masthead__main--ColumnGap);\\n  align-items: center;\\n}\\n\\n#toggle {\\n  grid-column: var(--pf-v6-c-masthead__toggle--GridColumn);\\n  display: flex;\\n  align-items: center;\\n}\\n\\n#brand {\\n  position: relative;\\n  grid-column: var(--pf-v6-c-masthead__brand--GridColumn);\\n  order: var(--pf-v6-c-masthead__brand--Order);\\n  padding-block-end: var(--pf-v6-c-masthead__brand--PaddingBlockEnd);\\n  border-block-end: var(--pf-v6-c-masthead__brand--BorderBlockEnd);\\n}\\n\\n#content {\\n  grid-column: var(--pf-v6-c-masthead__content--GridColumn);\\n  order: var(--pf-v6-c-masthead__content--Order);\\n  column-gap: var(--pf-v6-c-masthead__content--ColumnGap);\\n}\\n\\n::slotted(a[slot=\\"logo\\"]),\\n::slotted(img[slot=\\"logo\\"]),\\n::slotted(svg[slot=\\"logo\\"]) {\\n  position: relative;\\n  width: var(--pf-v6-c-masthead__logo--Width);\\n  max-height: var(--pf-v6-c-masthead__logo--MaxHeight);\\n}\\n\\n#toggle pf-v6-button {\\n  font-size: var(--pf-v6-c-masthead__toggle--Size);\\n}\\n\\n::slotted(pf-v6-toolbar) {\\n  --pf-v6-c-toolbar--Width: var(--pf-v6-c-masthead--c-toolbar--Width);\\n  --pf-v6-c-toolbar--PaddingBlockEnd: var(--pf-v6-c-masthead--c-toolbar--PaddingBlockEnd);\\n  --pf-v6-c-toolbar__content--MinWidth: var(--pf-v6-c-masthead--c-toolbar__content--MinWidth);\\n  --pf-v6-c-toolbar__content--PaddingInlineStart: var(--pf-v6-c-masthead--c-toolbar__content--PaddingInlineStart);\\n  --pf-v6-c-toolbar__content--PaddingInlineEnd: var(--pf-v6-c-masthead--c-toolbar__content--PaddingInlineEnd);\\n  --pf-v6-c-toolbar__content-section--FlexWrap: var(--pf-v6-c-masthead--c-toolbar__content-section--FlexWrap);\\n}\\n"'));
var pf_v6_masthead_default = s;

// elements/pf-v6-masthead/pf-v6-masthead.ts
import "../pf-v6-button/pf-v6-button.js";
var SidebarToggleEvent = class extends Event {
  expanded;
  constructor(expanded) {
    super("sidebar-toggle", { bubbles: true });
    this.expanded = expanded;
  }
};
var _sidebarExpanded_dec, _a, _PfV6Masthead_decorators, _internals, _init, _sidebarExpanded, _PfV6Masthead_instances, onToggle_fn;
_PfV6Masthead_decorators = [customElement("pf-v6-masthead")];
var PfV6Masthead = class extends (_a = LitElement, _sidebarExpanded_dec = [property({ type: Boolean, reflect: true, attribute: "sidebar-expanded" })], _a) {
  constructor() {
    super();
    __privateAdd(this, _PfV6Masthead_instances);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _sidebarExpanded, __runInitializers(_init, 8, this, false)), __runInitializers(_init, 11, this);
    __privateGet(this, _internals).role = "banner";
  }
  render() {
    return html`
      <div id="main">
        <div id="toggle">
          <pf-v6-button id="toggle-button"
                        variant="plain"
                        aria-label="Toggle global navigation"
                        aria-expanded=${String(this.sidebarExpanded)}
                        @click=${__privateMethod(this, _PfV6Masthead_instances, onToggle_fn)}>
            <svg id="hamburger"
                 role="presentation"
                 viewBox="0 0 10 10"
                 width="1em"
                 height="1em"
                 fill="none"
                 stroke="currentColor"
                 stroke-width="1.5"
                 stroke-linecap="round">
              <path d="M1,1 L9,1" />
              <path d="M1,5 L9,5" />
              <path d="M9,9 L1,9" />
            </svg>
          </pf-v6-button>
        </div>
        <div id="brand">
          <slot name="logo"></slot>
        </div>
      </div>
      <div id="content">
        <slot name="toolbar"></slot>
      </div>
    `;
  }
};
_init = __decoratorStart(_a);
_internals = new WeakMap();
_sidebarExpanded = new WeakMap();
_PfV6Masthead_instances = new WeakSet();
onToggle_fn = function() {
  this.dispatchEvent(new SidebarToggleEvent(!this.sidebarExpanded));
};
__decorateElement(_init, 4, "sidebarExpanded", _sidebarExpanded_dec, PfV6Masthead, _sidebarExpanded);
PfV6Masthead = __decorateElement(_init, 0, "PfV6Masthead", _PfV6Masthead_decorators, PfV6Masthead);
__publicField(PfV6Masthead, "styles", pf_v6_masthead_default);
__runInitializers(_init, 1, PfV6Masthead);
export {
  PfV6Masthead,
  SidebarToggleEvent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtbWFzdGhlYWQvcGYtdjYtbWFzdGhlYWQudHMiLCAibGl0LWNzczplbGVtZW50cy9wZi12Ni1tYXN0aGVhZC9wZi12Ni1tYXN0aGVhZC5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcbmltcG9ydCB7IHByb3BlcnR5IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvcHJvcGVydHkuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vcGYtdjYtbWFzdGhlYWQuY3NzJztcblxuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZm9yIHNpZGViYXIgdG9nZ2xlXG4gKi9cbmV4cG9ydCBjbGFzcyBTaWRlYmFyVG9nZ2xlRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIGV4cGFuZGVkOiBib29sZWFuO1xuICBjb25zdHJ1Y3RvcihleHBhbmRlZDogYm9vbGVhbikge1xuICAgIHN1cGVyKCdzaWRlYmFyLXRvZ2dsZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLmV4cGFuZGVkID0gZXhwYW5kZWQ7XG4gIH1cbn1cblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IE1hc3RoZWFkXG4gKlxuICogQHNsb3QgbG9nbyAtIExvZ28gb3IgYnJhbmQgY29udGVudFxuICogQHNsb3QgdG9vbGJhciAtIFRvb2xiYXIgY29udGVudCAoYWN0aW9ucywgbWVudXMsIGV0Yy4pXG4gKlxuICogQGZpcmVzIHtTaWRlYmFyVG9nZ2xlRXZlbnR9IHNpZGViYXItdG9nZ2xlIC0gV2hlbiB0aGUgaGFtYnVyZ2VyIHRvZ2dsZSBpcyBjbGlja2VkXG4gKi9cbkBjdXN0b21FbGVtZW50KCdwZi12Ni1tYXN0aGVhZCcpXG5leHBvcnQgY2xhc3MgUGZWNk1hc3RoZWFkIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgI2ludGVybmFscyA9IHRoaXMuYXR0YWNoSW50ZXJuYWxzKCk7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSwgYXR0cmlidXRlOiAnc2lkZWJhci1leHBhbmRlZCcgfSlcbiAgYWNjZXNzb3Igc2lkZWJhckV4cGFuZGVkID0gZmFsc2U7XG5cbiAgY29uc3RydWN0b3IoKSB7XG4gICAgc3VwZXIoKTtcbiAgICB0aGlzLiNpbnRlcm5hbHMucm9sZSA9ICdiYW5uZXInO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cIm1haW5cIj5cbiAgICAgICAgPGRpdiBpZD1cInRvZ2dsZVwiPlxuICAgICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJ0b2dnbGUtYnV0dG9uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiVG9nZ2xlIGdsb2JhbCBuYXZpZ2F0aW9uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtZXhwYW5kZWQ9JHtTdHJpbmcodGhpcy5zaWRlYmFyRXhwYW5kZWQpfVxuICAgICAgICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25Ub2dnbGV9PlxuICAgICAgICAgICAgPHN2ZyBpZD1cImhhbWJ1cmdlclwiXG4gICAgICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIlxuICAgICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDEwIDEwXCJcbiAgICAgICAgICAgICAgICAgd2lkdGg9XCIxZW1cIlxuICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxZW1cIlxuICAgICAgICAgICAgICAgICBmaWxsPVwibm9uZVwiXG4gICAgICAgICAgICAgICAgIHN0cm9rZT1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICAgIHN0cm9rZS13aWR0aD1cIjEuNVwiXG4gICAgICAgICAgICAgICAgIHN0cm9rZS1saW5lY2FwPVwicm91bmRcIj5cbiAgICAgICAgICAgICAgPHBhdGggZD1cIk0xLDEgTDksMVwiIC8+XG4gICAgICAgICAgICAgIDxwYXRoIGQ9XCJNMSw1IEw5LDVcIiAvPlxuICAgICAgICAgICAgICA8cGF0aCBkPVwiTTksOSBMMSw5XCIgLz5cbiAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgIDwvcGYtdjYtYnV0dG9uPlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBpZD1cImJyYW5kXCI+XG4gICAgICAgICAgPHNsb3QgbmFtZT1cImxvZ29cIj48L3Nsb3Q+XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9kaXY+XG4gICAgICA8ZGl2IGlkPVwiY29udGVudFwiPlxuICAgICAgICA8c2xvdCBuYW1lPVwidG9vbGJhclwiPjwvc2xvdD5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cblxuICAjb25Ub2dnbGUoKSB7XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBTaWRlYmFyVG9nZ2xlRXZlbnQoIXRoaXMuc2lkZWJhckV4cGFuZGVkKSk7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAncGYtdjYtbWFzdGhlYWQnOiBQZlY2TWFzdGhlYWQ7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgLS1wZi12Ni1jLW1hc3RoZWFkLS1Db2x1bW5HYXA6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLW0tZGlzcGxheS1zdGFjay0tQ29sdW1uR2FwKTtcXFxcbiAgLS1wZi12Ni1jLW1hc3RoZWFkLS1HcmlkVGVtcGxhdGVDb2x1bW5zOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1tLWRpc3BsYXktc3RhY2stLUdyaWRUZW1wbGF0ZUNvbHVtbnMpO1xcXFxuICAtLXBmLXY2LWMtbWFzdGhlYWRfX3RvZ2dsZS0tR3JpZENvbHVtbjogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZC0tbS1kaXNwbGF5LXN0YWNrX190b2dnbGUtLUdyaWRDb2x1bW4pO1xcXFxuICAtLXBmLXY2LWMtbWFzdGhlYWRfX2JyYW5kLS1HcmlkQ29sdW1uOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1tLWRpc3BsYXktc3RhY2tfX2JyYW5kLS1HcmlkQ29sdW1uKTtcXFxcbiAgLS1wZi12Ni1jLW1hc3RoZWFkX19icmFuZC0tT3JkZXI6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLW0tZGlzcGxheS1zdGFja19fYnJhbmQtLU9yZGVyKTtcXFxcbiAgLS1wZi12Ni1jLW1hc3RoZWFkX19icmFuZC0tUGFkZGluZ0Jsb2NrRW5kOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1tLWRpc3BsYXktc3RhY2tfX2JyYW5kLS1QYWRkaW5nQmxvY2tFbmQpO1xcXFxuICAtLXBmLXY2LWMtbWFzdGhlYWRfX2JyYW5kLS1Cb3JkZXJCbG9ja0VuZDogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZC0tbS1kaXNwbGF5LXN0YWNrX19icmFuZC0tQm9yZGVyQmxvY2tFbmQpO1xcXFxuICAtLXBmLXY2LWMtbWFzdGhlYWRfX21haW4tLUdyaWRDb2x1bW46IHVuc2V0O1xcXFxuICAtLXBmLXY2LWMtbWFzdGhlYWRfX2NvbnRlbnQtLUdyaWRDb2x1bW46IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLW0tZGlzcGxheS1zdGFja19fY29udGVudC0tR3JpZENvbHVtbik7XFxcXG4gIC0tcGYtdjYtYy1tYXN0aGVhZF9fY29udGVudC0tT3JkZXI6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLW0tZGlzcGxheS1zdGFja19fY29udGVudC0tT3JkZXIpO1xcXFxuICAtLXBmLXY2LWMtbWFzdGhlYWRfX21haW4tLURpc3BsYXk6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLW0tZGlzcGxheS1zdGFja19fbWFpbi0tRGlzcGxheSk7XFxcXG4gIC0tcGYtdjYtYy1tYXN0aGVhZF9fbWFpbi0tQ29sdW1uR2FwOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1tLWRpc3BsYXktc3RhY2tfX21haW4tLUNvbHVtbkdhcCk7XFxcXG5cXFxcbiAgcG9zaXRpb246IHJlbGF0aXZlO1xcXFxuICB6LWluZGV4OiB2YXIoLS1wZi12Ni1jLXBhZ2UtLWMtbWFzdGhlYWQtLVpJbmRleCk7XFxcXG4gIGdyaWQtYXJlYTogaGVhZGVyO1xcXFxuICBkaXNwbGF5OiBncmlkO1xcXFxuICBncmlkLXRlbXBsYXRlLWNvbHVtbnM6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLUdyaWRUZW1wbGF0ZUNvbHVtbnMpO1xcXFxuICByb3ctZ2FwOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1Sb3dHYXApO1xcXFxuICBjb2x1bW4tZ2FwOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1Db2x1bW5HYXApO1xcXFxuICBhbGlnbi1pdGVtczogc3RhcnQ7XFxcXG4gIG1pbi13aWR0aDogMDtcXFxcbiAgcGFkZGluZy1ibG9jay1zdGFydDogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZC0tUGFkZGluZ0Jsb2NrU3RhcnQsIHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLVBhZGRpbmdCbG9jaykpO1xcXFxuICBwYWRkaW5nLWJsb2NrLWVuZDogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZC0tUGFkZGluZ0Jsb2NrRW5kLCB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1QYWRkaW5nQmxvY2spKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLVBhZGRpbmdJbmxpbmVTdGFydCwgdmFyKC0tcGYtdjYtYy1tYXN0aGVhZC0tUGFkZGluZ0lubGluZSkpO1xcXFxuICBwYWRkaW5nLWlubGluZS1lbmQ6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLVBhZGRpbmdJbmxpbmVFbmQsIHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLVBhZGRpbmdJbmxpbmUpKTtcXFxcbiAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZC0tQmFja2dyb3VuZENvbG9yKTtcXFxcbn1cXFxcblxcXFxuI21haW4ge1xcXFxuICBkaXNwbGF5OiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkX19tYWluLS1EaXNwbGF5KTtcXFxcbiAgZ3JpZC1jb2x1bW46IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX21haW4tLUdyaWRDb2x1bW4pO1xcXFxuICBjb2x1bW4tZ2FwOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkX19tYWluLS1Db2x1bW5HYXApO1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxufVxcXFxuXFxcXG4jdG9nZ2xlIHtcXFxcbiAgZ3JpZC1jb2x1bW46IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX3RvZ2dsZS0tR3JpZENvbHVtbik7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG59XFxcXG5cXFxcbiNicmFuZCB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgZ3JpZC1jb2x1bW46IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX2JyYW5kLS1HcmlkQ29sdW1uKTtcXFxcbiAgb3JkZXI6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX2JyYW5kLS1PcmRlcik7XFxcXG4gIHBhZGRpbmctYmxvY2stZW5kOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkX19icmFuZC0tUGFkZGluZ0Jsb2NrRW5kKTtcXFxcbiAgYm9yZGVyLWJsb2NrLWVuZDogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZF9fYnJhbmQtLUJvcmRlckJsb2NrRW5kKTtcXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQge1xcXFxuICBncmlkLWNvbHVtbjogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZF9fY29udGVudC0tR3JpZENvbHVtbik7XFxcXG4gIG9yZGVyOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkX19jb250ZW50LS1PcmRlcik7XFxcXG4gIGNvbHVtbi1nYXA6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX2NvbnRlbnQtLUNvbHVtbkdhcCk7XFxcXG59XFxcXG5cXFxcbjo6c2xvdHRlZChhW3Nsb3Q9XFxcXFxcXCJsb2dvXFxcXFxcXCJdKSxcXFxcbjo6c2xvdHRlZChpbWdbc2xvdD1cXFxcXFxcImxvZ29cXFxcXFxcIl0pLFxcXFxuOjpzbG90dGVkKHN2Z1tzbG90PVxcXFxcXFwibG9nb1xcXFxcXFwiXSkge1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIHdpZHRoOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkX19sb2dvLS1XaWR0aCk7XFxcXG4gIG1heC1oZWlnaHQ6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX2xvZ28tLU1heEhlaWdodCk7XFxcXG59XFxcXG5cXFxcbiN0b2dnbGUgcGYtdjYtYnV0dG9uIHtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkX190b2dnbGUtLVNpemUpO1xcXFxufVxcXFxuXFxcXG46OnNsb3R0ZWQocGYtdjYtdG9vbGJhcikge1xcXFxuICAtLXBmLXY2LWMtdG9vbGJhci0tV2lkdGg6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLWMtdG9vbGJhci0tV2lkdGgpO1xcXFxuICAtLXBmLXY2LWMtdG9vbGJhci0tUGFkZGluZ0Jsb2NrRW5kOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1jLXRvb2xiYXItLVBhZGRpbmdCbG9ja0VuZCk7XFxcXG4gIC0tcGYtdjYtYy10b29sYmFyX19jb250ZW50LS1NaW5XaWR0aDogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZC0tYy10b29sYmFyX19jb250ZW50LS1NaW5XaWR0aCk7XFxcXG4gIC0tcGYtdjYtYy10b29sYmFyX19jb250ZW50LS1QYWRkaW5nSW5saW5lU3RhcnQ6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLWMtdG9vbGJhcl9fY29udGVudC0tUGFkZGluZ0lubGluZVN0YXJ0KTtcXFxcbiAgLS1wZi12Ni1jLXRvb2xiYXJfX2NvbnRlbnQtLVBhZGRpbmdJbmxpbmVFbmQ6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWQtLWMtdG9vbGJhcl9fY29udGVudC0tUGFkZGluZ0lubGluZUVuZCk7XFxcXG4gIC0tcGYtdjYtYy10b29sYmFyX19jb250ZW50LXNlY3Rpb24tLUZsZXhXcmFwOiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkLS1jLXRvb2xiYXJfX2NvbnRlbnQtc2VjdGlvbi0tRmxleFdyYXApO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLGd2SEFBd3ZILENBQUM7QUFBRSxJQUFPLHlCQUFROzs7QURNL3pILE9BQU87QUFLQSxJQUFNLHFCQUFOLGNBQWlDLE1BQU07QUFBQSxFQUM1QztBQUFBLEVBQ0EsWUFBWSxVQUFtQjtBQUM3QixVQUFNLGtCQUFrQixFQUFFLFNBQVMsS0FBSyxDQUFDO0FBQ3pDLFNBQUssV0FBVztBQUFBLEVBQ2xCO0FBQ0Y7QUFqQkE7QUEyQkEsNEJBQUMsY0FBYyxnQkFBZ0I7QUFDeEIsSUFBTSxlQUFOLGVBQTJCLGlCQUtoQyx3QkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFNBQVMsTUFBTSxXQUFXLG1CQUFtQixDQUFDLElBTHpDLElBQVc7QUFBQSxFQVEzQyxjQUFjO0FBQ1osVUFBTTtBQVRIO0FBR0wsbUNBQWEsS0FBSyxnQkFBZ0I7QUFHbEMsdUJBQVMsa0JBQWtCLGtCQUEzQixnQkFBMkIsU0FBM0I7QUFJRSx1QkFBSyxZQUFXLE9BQU87QUFBQSxFQUN6QjtBQUFBLEVBRUEsU0FBUztBQUNQLFdBQU87QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsd0NBTTZCLE9BQU8sS0FBSyxlQUFlLENBQUM7QUFBQSxpQ0FDbkMsc0JBQUsscUNBQVM7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsRUF3QjdDO0FBS0Y7QUFsRE87QUFHTDtBQUdTO0FBTko7QUErQ0wsY0FBUyxXQUFHO0FBQ1YsT0FBSyxjQUFjLElBQUksbUJBQW1CLENBQUMsS0FBSyxlQUFlLENBQUM7QUFDbEU7QUEzQ0EsNEJBQVMsbUJBRFQsc0JBTFcsY0FNRjtBQU5FLGVBQU4sNENBRFAsMEJBQ2E7QUFDWCxjQURXLGNBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
