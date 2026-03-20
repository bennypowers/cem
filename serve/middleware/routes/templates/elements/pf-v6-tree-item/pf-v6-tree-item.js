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

// elements/pf-v6-tree-item/pf-v6-tree-item.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";
import { classMap } from "/__cem/vendor/lit/directives/class-map.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-tree-item/pf-v6-tree-item.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n\\n  --_node-padding-block: var(--pf-t--global--spacer--sm, 0.5rem);\\n  --_node-padding-inline: var(--pf-t--global--spacer--sm, 0.5rem);\\n  --_indent: var(--pf-t--global--spacer--lg, 1rem);\\n\\n  --_node-color: var(--pf-t--global--text--color--subtle, #6a6e73);\\n  --_node-bg: transparent;\\n  --_node-hover-bg: var(--pf-t--global--background--color--primary--hover, #f0f0f0);\\n  --_node-current-color: var(--pf-t--global--text--color--regular, #151515);\\n  --_node-current-bg: var(--pf-t--global--background--color--primary--clicked, #e0e0e0);\\n\\n  --_node-border-radius: var(--pf-t--global--border--radius--small, 0.1875rem);\\n  --_node-border-color: var(--pf-t--global--border--color--high-contrast, transparent);\\n  --_node-border-width: var(--pf-t--global--border--width--action--plain--default, 0);\\n  --_node-hover-border-width: var(--pf-t--global--border--width--action--plain--hover, 0.0625rem);\\n  --_node-current-border-width: var(--pf-t--global--border--width--action--plain--clicked, 0.125rem);\\n\\n  --_toggle-color: var(--pf-t--global--icon--color--subtle, #6a6e73);\\n  --_toggle-expanded-color: var(--pf-t--global--icon--color--regular, #151515);\\n  --_toggle-padding: var(--pf-t--global--spacer--sm, 0.5rem) var(--pf-t--global--spacer--md, 1rem);\\n\\n  --_icon-color: var(--pf-t--global--icon--color--subtle, #6a6e73);\\n  --_icon-spacing: var(--pf-t--global--spacer--sm, 0.5rem);\\n  --_badge-spacing: var(--pf-t--global--spacer--sm, 0.5rem);\\n\\n  --_transition-duration: var(--pf-t--global--motion--duration--fade--short, 150ms);\\n  --_transition-timing: var(--pf-t--global--motion--timing-function--default, ease);\\n\\n  display: block;\\n}\\n\\n:host([hidden]) {\\n  display: none !important;\\n}\\n\\n#item {\\n  position: relative;\\n  margin: 0;\\n  padding: 0;\\n  list-style: none;\\n}\\n\\n#content {\\n  display: flex;\\n  align-items: center;\\n  background-color: var(--_node-bg);\\n  border-radius: var(--_node-border-radius);\\n\\n  \\u0026:hover {\\n    background-color: var(--_node-hover-bg);\\n  }\\n}\\n\\n#content:has(\\u003e #node.current) {\\n  background-color: var(--_node-current-bg);\\n}\\n\\n#node {\\n  position: relative;\\n  display: flex;\\n  flex: 1;\\n  align-items: flex-start;\\n  min-width: 0;\\n  padding-block: var(--_node-padding-block);\\n  padding-inline: var(--_node-padding-inline);\\n  color: var(--_node-color);\\n  background: transparent;\\n  border: 0;\\n  text-align: start;\\n  cursor: pointer;\\n\\n  \\u0026::after {\\n    position: absolute;\\n    inset: 0;\\n    pointer-events: none;\\n    content: \\"\\";\\n    border: var(--_node-border-width) solid var(--_node-border-color);\\n    border-radius: var(--_node-border-radius);\\n  }\\n\\n  \\u0026.current {\\n    color: var(--_node-current-color);\\n\\n    \\u0026::after {\\n      border-width: var(--_node-current-border-width);\\n    }\\n  }\\n}\\n\\n#content:hover #node::after {\\n  border-width: var(--_node-hover-border-width);\\n}\\n\\n#node-container {\\n  display: contents;\\n  flex-grow: 1;\\n}\\n\\n#toggle {\\n  display: none;\\n  align-items: center;\\n  justify-content: center;\\n  padding: var(--_toggle-padding);\\n  margin-block: calc(var(--_node-padding-block) * -1);\\n  color: var(--_toggle-color);\\n  background: transparent;\\n  border: 0;\\n  cursor: pointer;\\n}\\n\\n:host([has-children]) #toggle {\\n  display: inline-flex;\\n}\\n\\n:host([expanded]) #toggle {\\n  color: var(--_toggle-expanded-color);\\n}\\n\\n#toggle-icon {\\n  display: inline-block;\\n  min-width: 1em;\\n  text-align: center;\\n  transition: transform var(--_transition-duration) var(--_transition-timing);\\n  transform: rotate(0deg);\\n}\\n\\n:host([expanded]) #toggle-icon {\\n  transform: rotate(90deg);\\n}\\n\\n:host(:dir(rtl)) #toggle-icon {\\n  scale: -1 1;\\n}\\n\\n#icon {\\n  padding-inline-end: var(--_icon-spacing);\\n  color: var(--_icon-color);\\n\\n  \\u0026:empty {\\n    display: none;\\n  }\\n}\\n\\n#node-text {\\n  font-weight: inherit;\\n  color: inherit;\\n}\\n\\n#badge-container {\\n  margin-inline-start: var(--_badge-spacing);\\n\\n  \\u0026:empty {\\n    display: none;\\n  }\\n}\\n\\n#children {\\n  margin: 0;\\n  padding: 0;\\n  padding-inline-start: var(--_indent);\\n  list-style: none;\\n  max-height: 0;\\n  overflow: hidden;\\n  opacity: 0;\\n  transition:\\n    opacity var(--_transition-duration) var(--_transition-timing),\\n    max-height var(--_transition-duration) var(--_transition-timing);\\n}\\n\\n:host([expanded]) #children {\\n  max-height: 99999px;\\n  opacity: 1;\\n}\\n\\n@media (prefers-reduced-motion: reduce) {\\n  * {\\n    transition-duration: 0.01ms !important;\\n  }\\n}\\n"'));
var pf_v6_tree_item_default = s;

// elements/pf-v6-tree-item/pf-v6-tree-item.ts
import "../pf-v6-badge/pf-v6-badge.js";
var PfTreeItemSelectEvent = class extends Event {
  constructor() {
    super("select", { bubbles: true });
  }
};
var PfTreeItemExpandEvent = class extends Event {
  constructor() {
    super("expand", { bubbles: true });
  }
};
var PfTreeItemCollapseEvent = class extends Event {
  constructor() {
    super("collapse", { bubbles: true });
  }
};
var _hasChildren_dec, _current_dec, _expanded_dec, _badge_dec, _icon_dec, _label_dec, _a, _PfV6TreeItem_decorators, _init, _label, _icon, _badge, _expanded, _current, _hasChildren, _explicitHasChildren, _PfV6TreeItem_instances, onToggleClick_fn, onNodeClick_fn, onSlotChange_fn, updateIcon_fn;
_PfV6TreeItem_decorators = [customElement("pf-v6-tree-item")];
var PfV6TreeItem = class extends (_a = LitElement, _label_dec = [property({ reflect: true })], _icon_dec = [property()], _badge_dec = [property()], _expanded_dec = [property({ type: Boolean, reflect: true })], _current_dec = [property({ type: Boolean, reflect: true })], _hasChildren_dec = [property({ type: Boolean, reflect: true, attribute: "has-children" })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6TreeItem_instances);
    __privateAdd(this, _label, __runInitializers(_init, 8, this)), __runInitializers(_init, 11, this);
    __privateAdd(this, _icon, __runInitializers(_init, 12, this)), __runInitializers(_init, 15, this);
    __privateAdd(this, _badge, __runInitializers(_init, 16, this)), __runInitializers(_init, 19, this);
    __privateAdd(this, _expanded, __runInitializers(_init, 20, this, false)), __runInitializers(_init, 23, this);
    __privateAdd(this, _current, __runInitializers(_init, 24, this, false)), __runInitializers(_init, 27, this);
    __privateAdd(this, _hasChildren, __runInitializers(_init, 28, this, false)), __runInitializers(_init, 31, this);
    /** Track if has-children was explicitly set by the user */
    __privateAdd(this, _explicitHasChildren, false);
  }
  connectedCallback() {
    super.connectedCallback();
    if (this.hasAttribute("has-children")) {
      __privateSet(this, _explicitHasChildren, true);
    }
  }
  render() {
    return html`
      <li id="item"
          role="treeitem"
          tabindex="-1"
          aria-expanded=${this.hasChildren ? String(this.expanded) : nothing}
          aria-selected=${String(this.current)}
          aria-label=${this.label ?? nothing}>
        <div id="content">
          <button id="toggle"
                  type="button"
                  aria-label="Toggle"
                  @click=${__privateMethod(this, _PfV6TreeItem_instances, onToggleClick_fn)}>
            <span id="toggle-icon">
              <svg class="pf-v6-svg"
                   viewBox="0 0 256 512"
                   fill="currentColor"
                   role="presentation"
                   width="1em"
                   height="1em">
                <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
              </svg>
            </span>
          </button>
          <button id="node"
                  class=${classMap({ current: this.current })}
                  @click=${__privateMethod(this, _PfV6TreeItem_instances, onNodeClick_fn)}>
            <span id="node-container">
              <span id="icon"></span>
              <span id="node-text">
                <slot name="label">${this.label}</slot>
              </span>
              <span id="badge-container">
                <pf-v6-badge id="badge">
                  <slot name="badge">${this.badge ?? nothing}</slot>
                </pf-v6-badge>
              </span>
            </span>
          </button>
        </div>
        <ul id="children" role="group">
          <slot @slotchange=${__privateMethod(this, _PfV6TreeItem_instances, onSlotChange_fn)}></slot>
        </ul>
      </li>
    `;
  }
  updated(changed) {
    if (changed.has("icon")) {
      __privateMethod(this, _PfV6TreeItem_instances, updateIcon_fn).call(this);
    }
    if (changed.has("expanded") && changed.get("expanded") !== void 0) {
      this.dispatchEvent(
        this.expanded ? new PfTreeItemExpandEvent() : new PfTreeItemCollapseEvent()
      );
    }
  }
  /** Toggle expanded state */
  toggle() {
    if (!this.hasChildren) return;
    this.expanded = !this.expanded;
  }
  /** Expand the item */
  expand() {
    if (!this.hasChildren) return;
    this.expanded = true;
  }
  /** Collapse the item */
  collapse() {
    if (!this.hasChildren) return;
    this.expanded = false;
  }
  /** Select this item */
  select() {
    this.current = true;
    this.dispatchEvent(new PfTreeItemSelectEvent());
  }
  /** Deselect this item */
  deselect() {
    this.current = false;
  }
  /** Set tabindex for roving tabindex pattern */
  setTabindex(value) {
    const item = this.shadowRoot?.getElementById("item");
    item?.setAttribute("tabindex", String(value));
  }
  /** Focus this item */
  focusItem() {
    const item = this.shadowRoot?.getElementById("item");
    item?.focus();
  }
};
_init = __decoratorStart(_a);
_label = new WeakMap();
_icon = new WeakMap();
_badge = new WeakMap();
_expanded = new WeakMap();
_current = new WeakMap();
_hasChildren = new WeakMap();
_explicitHasChildren = new WeakMap();
_PfV6TreeItem_instances = new WeakSet();
onToggleClick_fn = function(e) {
  e.stopPropagation();
  this.toggle();
};
onNodeClick_fn = function() {
  if (this.hasChildren) {
    this.toggle();
  }
  this.select();
};
onSlotChange_fn = function() {
  if (__privateGet(this, _explicitHasChildren)) return;
  const slot = this.shadowRoot?.querySelector("#children slot");
  if (!slot) return;
  const children = slot.assignedElements();
  this.hasChildren = children.length > 0;
};
updateIcon_fn = function() {
  const iconEl = this.shadowRoot?.getElementById("icon");
  if (!iconEl) return;
  iconEl.innerHTML = this.icon ?? "";
};
__decorateElement(_init, 4, "label", _label_dec, PfV6TreeItem, _label);
__decorateElement(_init, 4, "icon", _icon_dec, PfV6TreeItem, _icon);
__decorateElement(_init, 4, "badge", _badge_dec, PfV6TreeItem, _badge);
__decorateElement(_init, 4, "expanded", _expanded_dec, PfV6TreeItem, _expanded);
__decorateElement(_init, 4, "current", _current_dec, PfV6TreeItem, _current);
__decorateElement(_init, 4, "hasChildren", _hasChildren_dec, PfV6TreeItem, _hasChildren);
PfV6TreeItem = __decorateElement(_init, 0, "PfV6TreeItem", _PfV6TreeItem_decorators, PfV6TreeItem);
__publicField(PfV6TreeItem, "styles", pf_v6_tree_item_default);
__runInitializers(_init, 1, PfV6TreeItem);
export {
  PfTreeItemCollapseEvent,
  PfTreeItemExpandEvent,
  PfTreeItemSelectEvent,
  PfV6TreeItem
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtdHJlZS1pdGVtL3BmLXY2LXRyZWUtaXRlbS50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL3BmLXY2LXRyZWUtaXRlbS9wZi12Ni10cmVlLWl0ZW0uY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sLCBub3RoaW5nIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5pbXBvcnQgeyBwcm9wZXJ0eSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3Byb3BlcnR5LmpzJztcbmltcG9ydCB7IGNsYXNzTWFwIH0gZnJvbSAnbGl0L2RpcmVjdGl2ZXMvY2xhc3MtbWFwLmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL3BmLXY2LXRyZWUtaXRlbS5jc3MnO1xuXG5pbXBvcnQgJy4uL3BmLXY2LWJhZGdlL3BmLXY2LWJhZGdlLmpzJztcblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiB0cmVlIGl0ZW0gaXMgc2VsZWN0ZWRcbiAqL1xuZXhwb3J0IGNsYXNzIFBmVHJlZUl0ZW1TZWxlY3RFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgY29uc3RydWN0b3IoKSB7XG4gICAgc3VwZXIoJ3NlbGVjdCcsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgfVxufVxuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIHRyZWUgaXRlbSBpcyBleHBhbmRlZFxuICovXG5leHBvcnQgY2xhc3MgUGZUcmVlSXRlbUV4cGFuZEV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBjb25zdHJ1Y3RvcigpIHtcbiAgICBzdXBlcignZXhwYW5kJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICB9XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gdHJlZSBpdGVtIGlzIGNvbGxhcHNlZFxuICovXG5leHBvcnQgY2xhc3MgUGZUcmVlSXRlbUNvbGxhcHNlRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIGNvbnN0cnVjdG9yKCkge1xuICAgIHN1cGVyKCdjb2xsYXBzZScsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgfVxufVxuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgVHJlZSBJdGVtXG4gKlxuICogSW5kaXZpZHVhbCB0cmVlIG5vZGUgdGhhdCBjYW4gYmUgbmVzdGVkLlxuICpcbiAqIEBzbG90IC0gQ2hpbGQgdHJlZSBpdGVtc1xuICogQHNsb3QgbGFiZWwgLSBJdGVtIGxhYmVsIHRleHRcbiAqXG4gKiBAZmlyZXMgc2VsZWN0IC0gRmlyZXMgd2hlbiBpdGVtIGlzIHNlbGVjdGVkXG4gKiBAZmlyZXMgZXhwYW5kIC0gRmlyZXMgd2hlbiBpdGVtIGlzIGV4cGFuZGVkXG4gKiBAZmlyZXMgY29sbGFwc2UgLSBGaXJlcyB3aGVuIGl0ZW0gaXMgY29sbGFwc2VkXG4gKi9cbkBjdXN0b21FbGVtZW50KCdwZi12Ni10cmVlLWl0ZW0nKVxuZXhwb3J0IGNsYXNzIFBmVjZUcmVlSXRlbSBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgbGFiZWw/OiBzdHJpbmc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3IgaWNvbj86IHN0cmluZztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBiYWRnZT86IHN0cmluZztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCByZWZsZWN0OiB0cnVlIH0pXG4gIGFjY2Vzc29yIGV4cGFuZGVkID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBjdXJyZW50ID0gZmFsc2U7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgcmVmbGVjdDogdHJ1ZSwgYXR0cmlidXRlOiAnaGFzLWNoaWxkcmVuJyB9KVxuICBhY2Nlc3NvciBoYXNDaGlsZHJlbiA9IGZhbHNlO1xuXG4gIC8qKiBUcmFjayBpZiBoYXMtY2hpbGRyZW4gd2FzIGV4cGxpY2l0bHkgc2V0IGJ5IHRoZSB1c2VyICovXG4gICNleHBsaWNpdEhhc0NoaWxkcmVuID0gZmFsc2U7XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICBpZiAodGhpcy5oYXNBdHRyaWJ1dGUoJ2hhcy1jaGlsZHJlbicpKSB7XG4gICAgICB0aGlzLiNleHBsaWNpdEhhc0NoaWxkcmVuID0gdHJ1ZTtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8bGkgaWQ9XCJpdGVtXCJcbiAgICAgICAgICByb2xlPVwidHJlZWl0ZW1cIlxuICAgICAgICAgIHRhYmluZGV4PVwiLTFcIlxuICAgICAgICAgIGFyaWEtZXhwYW5kZWQ9JHt0aGlzLmhhc0NoaWxkcmVuID8gU3RyaW5nKHRoaXMuZXhwYW5kZWQpIDogbm90aGluZ31cbiAgICAgICAgICBhcmlhLXNlbGVjdGVkPSR7U3RyaW5nKHRoaXMuY3VycmVudCl9XG4gICAgICAgICAgYXJpYS1sYWJlbD0ke3RoaXMubGFiZWwgPz8gbm90aGluZ30+XG4gICAgICAgIDxkaXYgaWQ9XCJjb250ZW50XCI+XG4gICAgICAgICAgPGJ1dHRvbiBpZD1cInRvZ2dsZVwiXG4gICAgICAgICAgICAgICAgICB0eXBlPVwiYnV0dG9uXCJcbiAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJUb2dnbGVcIlxuICAgICAgICAgICAgICAgICAgQGNsaWNrPSR7dGhpcy4jb25Ub2dnbGVDbGlja30+XG4gICAgICAgICAgICA8c3BhbiBpZD1cInRvZ2dsZS1pY29uXCI+XG4gICAgICAgICAgICAgIDxzdmcgY2xhc3M9XCJwZi12Ni1zdmdcIlxuICAgICAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMjU2IDUxMlwiXG4gICAgICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiXG4gICAgICAgICAgICAgICAgICAgd2lkdGg9XCIxZW1cIlxuICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiPlxuICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNMjI0LjMgMjczbC0xMzYgMTM2Yy05LjQgOS40LTI0LjYgOS40LTMzLjkgMGwtMjIuNi0yMi42Yy05LjQtOS40LTkuNC0yNC42IDAtMzMuOWw5Ni40LTk2LjQtOTYuNC05Ni40Yy05LjQtOS40LTkuNC0yNC42IDAtMzMuOUw1NC4zIDEwM2M5LjQtOS40IDI0LjYtOS40IDMzLjkgMGwxMzYgMTM2YzkuNSA5LjQgOS41IDI0LjYuMSAzNHpcIj48L3BhdGg+XG4gICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgPC9zcGFuPlxuICAgICAgICAgIDwvYnV0dG9uPlxuICAgICAgICAgIDxidXR0b24gaWQ9XCJub2RlXCJcbiAgICAgICAgICAgICAgICAgIGNsYXNzPSR7Y2xhc3NNYXAoeyBjdXJyZW50OiB0aGlzLmN1cnJlbnQgfSl9XG4gICAgICAgICAgICAgICAgICBAY2xpY2s9JHt0aGlzLiNvbk5vZGVDbGlja30+XG4gICAgICAgICAgICA8c3BhbiBpZD1cIm5vZGUtY29udGFpbmVyXCI+XG4gICAgICAgICAgICAgIDxzcGFuIGlkPVwiaWNvblwiPjwvc3Bhbj5cbiAgICAgICAgICAgICAgPHNwYW4gaWQ9XCJub2RlLXRleHRcIj5cbiAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibGFiZWxcIj4ke3RoaXMubGFiZWx9PC9zbG90PlxuICAgICAgICAgICAgICA8L3NwYW4+XG4gICAgICAgICAgICAgIDxzcGFuIGlkPVwiYmFkZ2UtY29udGFpbmVyXCI+XG4gICAgICAgICAgICAgICAgPHBmLXY2LWJhZGdlIGlkPVwiYmFkZ2VcIj5cbiAgICAgICAgICAgICAgICAgIDxzbG90IG5hbWU9XCJiYWRnZVwiPiR7dGhpcy5iYWRnZSA/PyBub3RoaW5nfTwvc2xvdD5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LWJhZGdlPlxuICAgICAgICAgICAgICA8L3NwYW4+XG4gICAgICAgICAgICA8L3NwYW4+XG4gICAgICAgICAgPC9idXR0b24+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8dWwgaWQ9XCJjaGlsZHJlblwiIHJvbGU9XCJncm91cFwiPlxuICAgICAgICAgIDxzbG90IEBzbG90Y2hhbmdlPSR7dGhpcy4jb25TbG90Q2hhbmdlfT48L3Nsb3Q+XG4gICAgICAgIDwvdWw+XG4gICAgICA8L2xpPlxuICAgIGA7XG4gIH1cblxuICB1cGRhdGVkKGNoYW5nZWQ6IE1hcDxzdHJpbmcsIHVua25vd24+KSB7XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdpY29uJykpIHtcbiAgICAgIHRoaXMuI3VwZGF0ZUljb24oKTtcbiAgICB9XG4gICAgaWYgKGNoYW5nZWQuaGFzKCdleHBhbmRlZCcpICYmIGNoYW5nZWQuZ2V0KCdleHBhbmRlZCcpICE9PSB1bmRlZmluZWQpIHtcbiAgICAgIHRoaXMuZGlzcGF0Y2hFdmVudChcbiAgICAgICAgdGhpcy5leHBhbmRlZCA/IG5ldyBQZlRyZWVJdGVtRXhwYW5kRXZlbnQoKSA6IG5ldyBQZlRyZWVJdGVtQ29sbGFwc2VFdmVudCgpXG4gICAgICApO1xuICAgIH1cbiAgfVxuXG4gICNvblRvZ2dsZUNsaWNrKGU6IEV2ZW50KSB7XG4gICAgZS5zdG9wUHJvcGFnYXRpb24oKTtcbiAgICB0aGlzLnRvZ2dsZSgpO1xuICB9XG5cbiAgI29uTm9kZUNsaWNrKCkge1xuICAgIGlmICh0aGlzLmhhc0NoaWxkcmVuKSB7XG4gICAgICB0aGlzLnRvZ2dsZSgpO1xuICAgIH1cbiAgICB0aGlzLnNlbGVjdCgpO1xuICB9XG5cbiAgI29uU2xvdENoYW5nZSgpIHtcbiAgICBpZiAodGhpcy4jZXhwbGljaXRIYXNDaGlsZHJlbikgcmV0dXJuO1xuICAgIGNvbnN0IHNsb3QgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJyNjaGlsZHJlbiBzbG90JykgYXMgSFRNTFNsb3RFbGVtZW50IHwgbnVsbDtcbiAgICBpZiAoIXNsb3QpIHJldHVybjtcbiAgICBjb25zdCBjaGlsZHJlbiA9IHNsb3QuYXNzaWduZWRFbGVtZW50cygpO1xuICAgIHRoaXMuaGFzQ2hpbGRyZW4gPSBjaGlsZHJlbi5sZW5ndGggPiAwO1xuICB9XG5cbiAgI3VwZGF0ZUljb24oKSB7XG4gICAgY29uc3QgaWNvbkVsID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnaWNvbicpO1xuICAgIGlmICghaWNvbkVsKSByZXR1cm47XG4gICAgaWNvbkVsLmlubmVySFRNTCA9IHRoaXMuaWNvbiA/PyAnJztcbiAgfVxuXG4gIC8qKiBUb2dnbGUgZXhwYW5kZWQgc3RhdGUgKi9cbiAgdG9nZ2xlKCkge1xuICAgIGlmICghdGhpcy5oYXNDaGlsZHJlbikgcmV0dXJuO1xuICAgIHRoaXMuZXhwYW5kZWQgPSAhdGhpcy5leHBhbmRlZDtcbiAgfVxuXG4gIC8qKiBFeHBhbmQgdGhlIGl0ZW0gKi9cbiAgZXhwYW5kKCkge1xuICAgIGlmICghdGhpcy5oYXNDaGlsZHJlbikgcmV0dXJuO1xuICAgIHRoaXMuZXhwYW5kZWQgPSB0cnVlO1xuICB9XG5cbiAgLyoqIENvbGxhcHNlIHRoZSBpdGVtICovXG4gIGNvbGxhcHNlKCkge1xuICAgIGlmICghdGhpcy5oYXNDaGlsZHJlbikgcmV0dXJuO1xuICAgIHRoaXMuZXhwYW5kZWQgPSBmYWxzZTtcbiAgfVxuXG4gIC8qKiBTZWxlY3QgdGhpcyBpdGVtICovXG4gIHNlbGVjdCgpIHtcbiAgICB0aGlzLmN1cnJlbnQgPSB0cnVlO1xuICAgIHRoaXMuZGlzcGF0Y2hFdmVudChuZXcgUGZUcmVlSXRlbVNlbGVjdEV2ZW50KCkpO1xuICB9XG5cbiAgLyoqIERlc2VsZWN0IHRoaXMgaXRlbSAqL1xuICBkZXNlbGVjdCgpIHtcbiAgICB0aGlzLmN1cnJlbnQgPSBmYWxzZTtcbiAgfVxuXG4gIC8qKiBTZXQgdGFiaW5kZXggZm9yIHJvdmluZyB0YWJpbmRleCBwYXR0ZXJuICovXG4gIHNldFRhYmluZGV4KHZhbHVlOiBudW1iZXIpIHtcbiAgICBjb25zdCBpdGVtID0gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZCgnaXRlbScpO1xuICAgIGl0ZW0/LnNldEF0dHJpYnV0ZSgndGFiaW5kZXgnLCBTdHJpbmcodmFsdWUpKTtcbiAgfVxuXG4gIC8qKiBGb2N1cyB0aGlzIGl0ZW0gKi9cbiAgZm9jdXNJdGVtKCkge1xuICAgIGNvbnN0IGl0ZW0gPSB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCdpdGVtJyk7XG4gICAgaXRlbT8uZm9jdXMoKTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdwZi12Ni10cmVlLWl0ZW0nOiBQZlY2VHJlZUl0ZW07XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcblxcXFxuICAtLV9ub2RlLXBhZGRpbmctYmxvY2s6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSwgMC41cmVtKTtcXFxcbiAgLS1fbm9kZS1wYWRkaW5nLWlubGluZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtLCAwLjVyZW0pO1xcXFxuICAtLV9pbmRlbnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZywgMXJlbSk7XFxcXG5cXFxcbiAgLS1fbm9kZS1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlLCAjNmE2ZTczKTtcXFxcbiAgLS1fbm9kZS1iZzogdHJhbnNwYXJlbnQ7XFxcXG4gIC0tX25vZGUtaG92ZXItYmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWhvdmVyLCAjZjBmMGYwKTtcXFxcbiAgLS1fbm9kZS1jdXJyZW50LWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyLCAjMTUxNTE1KTtcXFxcbiAgLS1fbm9kZS1jdXJyZW50LWJnOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1jbGlja2VkLCAjZTBlMGUwKTtcXFxcblxcXFxuICAtLV9ub2RlLWJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXNtYWxsLCAwLjE4NzVyZW0pO1xcXFxuICAtLV9ub2RlLWJvcmRlci1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1oaWdoLWNvbnRyYXN0LCB0cmFuc3BhcmVudCk7XFxcXG4gIC0tX25vZGUtYm9yZGVyLXdpZHRoOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0td2lkdGgtLWFjdGlvbi0tcGxhaW4tLWRlZmF1bHQsIDApO1xcXFxuICAtLV9ub2RlLWhvdmVyLWJvcmRlci13aWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1hY3Rpb24tLXBsYWluLS1ob3ZlciwgMC4wNjI1cmVtKTtcXFxcbiAgLS1fbm9kZS1jdXJyZW50LWJvcmRlci13aWR0aDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLXdpZHRoLS1hY3Rpb24tLXBsYWluLS1jbGlja2VkLCAwLjEyNXJlbSk7XFxcXG5cXFxcbiAgLS1fdG9nZ2xlLWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWljb24tLWNvbG9yLS1zdWJ0bGUsICM2YTZlNzMpO1xcXFxuICAtLV90b2dnbGUtZXhwYW5kZWQtY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLXJlZ3VsYXIsICMxNTE1MTUpO1xcXFxuICAtLV90b2dnbGUtcGFkZGluZzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtLCAwLjVyZW0pIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCwgMXJlbSk7XFxcXG5cXFxcbiAgLS1faWNvbi1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1pY29uLS1jb2xvci0tc3VidGxlLCAjNmE2ZTczKTtcXFxcbiAgLS1faWNvbi1zcGFjaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tc20sIDAuNXJlbSk7XFxcXG4gIC0tX2JhZGdlLXNwYWNpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1zbSwgMC41cmVtKTtcXFxcblxcXFxuICAtLV90cmFuc2l0aW9uLWR1cmF0aW9uOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tZHVyYXRpb24tLWZhZGUtLXNob3J0LCAxNTBtcyk7XFxcXG4gIC0tX3RyYW5zaXRpb24tdGltaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLW1vdGlvbi0tdGltaW5nLWZ1bmN0aW9uLS1kZWZhdWx0LCBlYXNlKTtcXFxcblxcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2hpZGRlbl0pIHtcXFxcbiAgZGlzcGxheTogbm9uZSAhaW1wb3J0YW50O1xcXFxufVxcXFxuXFxcXG4jaXRlbSB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBsaXN0LXN0eWxlOiBub25lO1xcXFxufVxcXFxuXFxcXG4jY29udGVudCB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLV9ub2RlLWJnKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tX25vZGUtYm9yZGVyLXJhZGl1cyk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1fbm9kZS1ob3Zlci1iZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQ6aGFzKFxcXFx1MDAzZSAjbm9kZS5jdXJyZW50KSB7XFxcXG4gIGJhY2tncm91bmQtY29sb3I6IHZhcigtLV9ub2RlLWN1cnJlbnQtYmcpO1xcXFxufVxcXFxuXFxcXG4jbm9kZSB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleDogMTtcXFxcbiAgYWxpZ24taXRlbXM6IGZsZXgtc3RhcnQ7XFxcXG4gIG1pbi13aWR0aDogMDtcXFxcbiAgcGFkZGluZy1ibG9jazogdmFyKC0tX25vZGUtcGFkZGluZy1ibG9jayk7XFxcXG4gIHBhZGRpbmctaW5saW5lOiB2YXIoLS1fbm9kZS1wYWRkaW5nLWlubGluZSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1fbm9kZS1jb2xvcik7XFxcXG4gIGJhY2tncm91bmQ6IHRyYW5zcGFyZW50O1xcXFxuICBib3JkZXI6IDA7XFxcXG4gIHRleHQtYWxpZ246IHN0YXJ0O1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OjphZnRlciB7XFxcXG4gICAgcG9zaXRpb246IGFic29sdXRlO1xcXFxuICAgIGluc2V0OiAwO1xcXFxuICAgIHBvaW50ZXItZXZlbnRzOiBub25lO1xcXFxuICAgIGNvbnRlbnQ6IFxcXFxcXFwiXFxcXFxcXCI7XFxcXG4gICAgYm9yZGVyOiB2YXIoLS1fbm9kZS1ib3JkZXItd2lkdGgpIHNvbGlkIHZhcigtLV9ub2RlLWJvcmRlci1jb2xvcik7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogdmFyKC0tX25vZGUtYm9yZGVyLXJhZGl1cyk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYuY3VycmVudCB7XFxcXG4gICAgY29sb3I6IHZhcigtLV9ub2RlLWN1cnJlbnQtY29sb3IpO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OjphZnRlciB7XFxcXG4gICAgICBib3JkZXItd2lkdGg6IHZhcigtLV9ub2RlLWN1cnJlbnQtYm9yZGVyLXdpZHRoKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI2NvbnRlbnQ6aG92ZXIgI25vZGU6OmFmdGVyIHtcXFxcbiAgYm9yZGVyLXdpZHRoOiB2YXIoLS1fbm9kZS1ob3Zlci1ib3JkZXItd2lkdGgpO1xcXFxufVxcXFxuXFxcXG4jbm9kZS1jb250YWluZXIge1xcXFxuICBkaXNwbGF5OiBjb250ZW50cztcXFxcbiAgZmxleC1ncm93OiAxO1xcXFxufVxcXFxuXFxcXG4jdG9nZ2xlIHtcXFxcbiAgZGlzcGxheTogbm9uZTtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAganVzdGlmeS1jb250ZW50OiBjZW50ZXI7XFxcXG4gIHBhZGRpbmc6IHZhcigtLV90b2dnbGUtcGFkZGluZyk7XFxcXG4gIG1hcmdpbi1ibG9jazogY2FsYyh2YXIoLS1fbm9kZS1wYWRkaW5nLWJsb2NrKSAqIC0xKTtcXFxcbiAgY29sb3I6IHZhcigtLV90b2dnbGUtY29sb3IpO1xcXFxuICBiYWNrZ3JvdW5kOiB0cmFuc3BhcmVudDtcXFxcbiAgYm9yZGVyOiAwO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtoYXMtY2hpbGRyZW5dKSAjdG9nZ2xlIHtcXFxcbiAgZGlzcGxheTogaW5saW5lLWZsZXg7XFxcXG59XFxcXG5cXFxcbjpob3N0KFtleHBhbmRlZF0pICN0b2dnbGUge1xcXFxuICBjb2xvcjogdmFyKC0tX3RvZ2dsZS1leHBhbmRlZC1jb2xvcik7XFxcXG59XFxcXG5cXFxcbiN0b2dnbGUtaWNvbiB7XFxcXG4gIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgbWluLXdpZHRoOiAxZW07XFxcXG4gIHRleHQtYWxpZ246IGNlbnRlcjtcXFxcbiAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIHZhcigtLV90cmFuc2l0aW9uLWR1cmF0aW9uKSB2YXIoLS1fdHJhbnNpdGlvbi10aW1pbmcpO1xcXFxuICB0cmFuc2Zvcm06IHJvdGF0ZSgwZGVnKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2V4cGFuZGVkXSkgI3RvZ2dsZS1pY29uIHtcXFxcbiAgdHJhbnNmb3JtOiByb3RhdGUoOTBkZWcpO1xcXFxufVxcXFxuXFxcXG46aG9zdCg6ZGlyKHJ0bCkpICN0b2dnbGUtaWNvbiB7XFxcXG4gIHNjYWxlOiAtMSAxO1xcXFxufVxcXFxuXFxcXG4jaWNvbiB7XFxcXG4gIHBhZGRpbmctaW5saW5lLWVuZDogdmFyKC0tX2ljb24tc3BhY2luZyk7XFxcXG4gIGNvbG9yOiB2YXIoLS1faWNvbi1jb2xvcik7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmVtcHR5IHtcXFxcbiAgICBkaXNwbGF5OiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNub2RlLXRleHQge1xcXFxuICBmb250LXdlaWdodDogaW5oZXJpdDtcXFxcbiAgY29sb3I6IGluaGVyaXQ7XFxcXG59XFxcXG5cXFxcbiNiYWRnZS1jb250YWluZXIge1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiB2YXIoLS1fYmFkZ2Utc3BhY2luZyk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmVtcHR5IHtcXFxcbiAgICBkaXNwbGF5OiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNjaGlsZHJlbiB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLV9pbmRlbnQpO1xcXFxuICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICBtYXgtaGVpZ2h0OiAwO1xcXFxuICBvdmVyZmxvdzogaGlkZGVuO1xcXFxuICBvcGFjaXR5OiAwO1xcXFxuICB0cmFuc2l0aW9uOlxcXFxuICAgIG9wYWNpdHkgdmFyKC0tX3RyYW5zaXRpb24tZHVyYXRpb24pIHZhcigtLV90cmFuc2l0aW9uLXRpbWluZyksXFxcXG4gICAgbWF4LWhlaWdodCB2YXIoLS1fdHJhbnNpdGlvbi1kdXJhdGlvbikgdmFyKC0tX3RyYW5zaXRpb24tdGltaW5nKTtcXFxcbn1cXFxcblxcXFxuOmhvc3QoW2V4cGFuZGVkXSkgI2NoaWxkcmVuIHtcXFxcbiAgbWF4LWhlaWdodDogOTk5OTlweDtcXFxcbiAgb3BhY2l0eTogMTtcXFxcbn1cXFxcblxcXFxuQG1lZGlhIChwcmVmZXJzLXJlZHVjZWQtbW90aW9uOiByZWR1Y2UpIHtcXFxcbiAgKiB7XFxcXG4gICAgdHJhbnNpdGlvbi1kdXJhdGlvbjogMC4wMW1zICFpbXBvcnRhbnQ7XFxcXG4gIH1cXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLE1BQU0sZUFBZTtBQUMxQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjtBQUN6QixTQUFTLGdCQUFnQjs7O0FDSHpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLHFsSkFBeWxKLENBQUM7QUFBRSxJQUFPLDBCQUFROzs7QURPaHFKLE9BQU87QUFLQSxJQUFNLHdCQUFOLGNBQW9DLE1BQU07QUFBQSxFQUMvQyxjQUFjO0FBQ1osVUFBTSxVQUFVLEVBQUUsU0FBUyxLQUFLLENBQUM7QUFBQSxFQUNuQztBQUNGO0FBS08sSUFBTSx3QkFBTixjQUFvQyxNQUFNO0FBQUEsRUFDL0MsY0FBYztBQUNaLFVBQU0sVUFBVSxFQUFFLFNBQVMsS0FBSyxDQUFDO0FBQUEsRUFDbkM7QUFDRjtBQUtPLElBQU0sMEJBQU4sY0FBc0MsTUFBTTtBQUFBLEVBQ2pELGNBQWM7QUFDWixVQUFNLFlBQVksRUFBRSxTQUFTLEtBQUssQ0FBQztBQUFBLEVBQ3JDO0FBQ0Y7QUFsQ0E7QUFnREEsNEJBQUMsY0FBYyxpQkFBaUI7QUFDekIsSUFBTSxlQUFOLGVBQTJCLGlCQUdoQyxjQUFDLFNBQVMsRUFBRSxTQUFTLEtBQUssQ0FBQyxJ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
