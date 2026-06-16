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

// elements/cem-serve-chrome/cem-serve-chrome.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";
import { ifDefined } from "/__cem/vendor/lit/directives/if-defined.js";

// lit-css:elements/cem-serve-chrome/cem-serve-chrome.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse(`":host {\\n  display: block;\\n  height: 100vh;\\n  overflow: hidden;\\n  --cem-pf-v6-c-masthead__logo--Width: max-content;\\n  --cem-pf-v6-c-masthead__toggle--Size: 1rem;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n/* Masthead logo styles */\\n.masthead-logo {\\n  display: flex;\\n  align-items: center;\\n  text-decoration: none;\\n  color: inherit;\\n  max-height: var(--cem-pf-v6-c-masthead__logo--MaxHeight);\\n  gap: 4px;\\n  \\u0026 img {\\n    display: block;\\n    max-height: var(--cem-pf-v6-c-masthead__logo--MaxHeight);\\n    width: auto;\\n  }\\n  \\u0026 ::slotted([slot=\\"title\\"]) {\\n    margin: 0;\\n    font-size: 1.125rem;\\n    font-weight: 600;\\n    color: var(--pf-t--global--text--color--regular);\\n  }\\n  \\u0026 h1 {\\n    margin: 0;\\n    font-size: 18px;\\n  }\\n}\\n\\n/* Toolbar group alignment */\\ncem-pf-v6-toolbar-group[variant=\\"action-group\\"] {\\n  margin-inline-start: auto;\\n  align-items: center;\\n}\\n\\n.debug-panel {\\n  background: var(--pf-t--global--background--color--primary--default);\\n  border: 1px solid var(--pf-t--global--border--color--default);\\n  border-radius: 6px;\\n  padding: 1.5rem;\\n  max-width: 600px;\\n  width: 90%;\\n  max-height: 80vh;\\n  overflow-y: auto;\\n\\n  h2 {\\n    margin: 0 0 1rem 0;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-size: 1.125rem;\\n    font-weight: 600;\\n  }\\n\\n  dl {\\n    margin: 0;\\n  }\\n\\n  dt {\\n    color: var(--pf-t--global--text--color--subtle);\\n    font-size: 0.875rem;\\n    margin-top: 0.5rem;\\n    font-weight: 500;\\n  }\\n\\n  dd {\\n    margin: 0 0 0.5rem 0;\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n    color: var(--pf-t--global--text--color--regular);\\n  }\\n\\n  details {\\n    margin-top: 1rem;\\n\\n    summary {\\n      cursor: pointer;\\n      color: var(--pf-t--global--text--color--regular);\\n      font-size: 0.875rem;\\n      font-weight: 500;\\n      list-style: none;\\n      display: flex;\\n      align-items: center;\\n      gap: 0.5rem;\\n      user-select: none;\\n\\n      \\u0026::-webkit-details-marker {\\n        display: none;\\n      }\\n\\n      \\u0026::before {\\n        content: '\\\\25B8';\\n        display: inline-block;\\n        transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n        color: var(--pf-t--global--text--color--subtle);\\n      }\\n    }\\n\\n    \\u0026[open] summary::before {\\n      transform: rotate(90deg);\\n    }\\n\\n    pre {\\n      margin-top: 0.5rem;\\n      margin-left: 1.5rem;\\n      padding: 0.5rem;\\n      background: var(--pf-t--global--background--color--secondary--default);\\n      border-radius: 6px;\\n      font-size: 0.875rem;\\n      overflow-x: auto;\\n      color: var(--pf-t--global--text--color--regular);\\n    }\\n  }\\n\\n  .button-container {\\n    display: flex;\\n    gap: 0.5rem;\\n    margin-top: 1rem;\\n  }\\n\\n  p {\\n    color: var(--pf-t--global--text--color--subtle);\\n    font-size: 0.875rem;\\n  }\\n\\n  button {\\n    margin-top: 1rem;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--color--brand--default);\\n    color: var(--pf-t--global--text--color--on-brand);\\n    border: none;\\n    border-radius: 6px;\\n    font-size: 0.875rem;\\n    font-weight: 400;\\n    cursor: pointer;\\n    transition: all 200ms cubic-bezier(0.645, 0.045, 0.355, 1);\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--color--brand--hover);\\n    }\\n  }\\n}\\n\\n/* Content area padding for demo */\\ncem-pf-v6-page-main {\\n  min-height: calc(100dvh - 72px - var(--pf-t--global--spacer--inset--page-chrome));\\n  display: flex;\\n  flex-direction: column;\\n  \\u0026 \\u003e ::slotted(:not([slot=knobs])) {\\n    padding: var(--pf-t--global--spacer--lg);\\n    flex: 1;\\n  }\\n}\\n\\ncem-drawer {\\n  cem-pf-v6-tabs {\\n    cem-pf-v6-tab {\\n      padding-block-end: 0;\\n    }\\n  }\\n}\\n\\n/* Element descriptions in listing */\\n.element-summary {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n.element-description {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n/* Card footer demo navigation */\\n.card-demos {\\n  display: flex;\\n  flex-wrap: wrap;\\n  gap: var(--pf-t--global--spacer--gap--action-to-action--default);\\n  padding: 0;\\n  margin: 0;\\n}\\n\\n.package-name {\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n/* Knobs container - fills tab panel height */\\n#knobs-container {\\n  height: 100%;\\n  display: flex;\\n  flex-direction: column;\\n  \\u0026 ::slotted([slot=\\"knobs\\"]) {\\n    flex: 1;\\n    min-height: 0;\\n    overflow: hidden;\\n  }\\n}\\n\\n.knobs-empty {\\n  color: var(--cem-dev-server-text-muted);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  text-align: center;\\n  padding: var(--cem-dev-server-spacing-lg);\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    padding: 2px 6px;\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n}\\n\\n.instance-tag {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  color: var(--cem-dev-server-accent-color);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.instance-label {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.knob-group {\\n  margin-bottom: var(--cem-dev-server-spacing-lg);\\n\\n  \\u0026:last-child {\\n    margin-bottom: 0;\\n  }\\n}\\n\\n/* PatternFly v6 form - horizontal layout */\\ncem-pf-v6-form[horizontal] cem-pf-v6-form-field-group {\\n  grid-column: span 2\\n}\\n\\n.knob-group-title {\\n  grid-column: 1 / -1;\\n  margin: 0 0 var(--cem-dev-server-spacing-md) 0;\\n  color: var(--cem-dev-server-text-primary);\\n  font-size: var(--cem-dev-server-font-size-base);\\n  font-weight: 600;\\n  border-bottom: 1px solid var(--cem-dev-server-border-color);\\n  padding-bottom: var(--cem-dev-server-spacing-sm);\\n}\\n\\n.knob-control {\\n  margin-bottom: var(--cem-dev-server-spacing-md);\\n}\\n\\n.knob-label {\\n  display: flex;\\n  flex-direction: column;\\n  gap: var(--cem-dev-server-spacing-xs);\\n  cursor: pointer;\\n}\\n\\n.knob-name {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-primary);\\n  font-weight: 500;\\n}\\n\\n.knob-description {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  line-height: 1.5;\\n\\n  p {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n  }\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n\\n  a {\\n    color: var(--cem-dev-server-accent-color);\\n    text-decoration: none;\\n\\n    \\u0026:hover {\\n      text-decoration: underline;\\n    }\\n  }\\n\\n  strong {\\n    font-weight: 600;\\n    color: var(--cem-dev-server-text-primary);\\n  }\\n\\n  ul, ol {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n    padding-left: var(--cem-dev-server-spacing-lg);\\n  }\\n}\\n\\nfooter.pf-m-sticky-bottom {\\n  view-transition-name: dev-server-footer;\\n  position: sticky;\\n  bottom: 0;\\n  background: var(--pf-t--global--background--color--primary--default);\\n  border-top: 1px solid var(--pf-t--global--border--color--default);\\n  z-index: var(--cem-pf-v6-c-page--section--m-sticky-bottom--ZIndex, var(--pf-t--global--z-index--md));\\n  box-shadow: var(--cem-pf-v6-c-page--section--m-sticky-bottom--BoxShadow, var(--pf-t--global--box-shadow--sm--top));\\n}\\n\\n.footer-description {\\n  padding: 1.5rem;\\n\\n  \\u0026.empty {\\n    display: none;\\n  }\\n}\\n\\nfooter ::slotted([slot=\\"description\\"]) {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  line-height: 1.6;\\n  font-size: 0.875rem;\\n\\n  code {\\n    background: var(--pf-t--global--background--color--primary--hover);\\n    padding: 2px 6px;\\n    border-radius: 3px;\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n  }\\n}\\n\\n.logs-wrapper {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n#log-container {\\n  flex-grow: 1;\\n  overflow-y: auto;\\n}\\n\\n.log-entry {\\n  padding: var(--cem-dev-server-spacing-xs) var(--cem-dev-server-spacing-md);\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-sm);\\n  align-items: baseline;\\n  cem-pf-v6-label {\\n    flex-shrink: 0;\\n  }\\n}\\n\\n.log-time,\\n.log-message {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.log-time {\\n  color: var(--cem-dev-server-text-muted);\\n  flex-shrink: 0;\\n  font-size: 11px;\\n}\\n\\n.log-message {\\n  color: var(--cem-dev-server-text-primary);\\n  word-break: break-word;\\n}\\n\\n/* Navigation content (light DOM slotted content for cem-pf-v6-page-sidebar) */\\n.nav-package {\\n  margin-bottom: var(--cem-dev-server-spacing-md);\\n\\n  \\u0026 \\u003e summary {\\n    cursor: pointer;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--background--color--secondary--default);\\n    border-radius: 6px;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-weight: 600;\\n    font-size: 0.875rem;\\n    list-style: none;\\n    transition: background 200ms cubic-bezier(0.4, 0, 0.2, 1);\\n    margin-bottom: 0.5rem;\\n    display: flex;\\n    align-items: center;\\n    gap: 0.5rem;\\n    user-select: none;\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--background--color--secondary--hover);\\n    }\\n\\n    \\u0026::-webkit-details-marker {\\n      display: none;\\n    }\\n\\n    \\u0026::before {\\n      content: '\\\\25B8';\\n      display: inline-block;\\n      transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n      color: var(--pf-t--global--text--color--subtle);\\n    }\\n  }\\n\\n  \\u0026[open] \\u003e summary::before {\\n    transform: rotate(90deg);\\n  }\\n}\\n\\n.nav-element {\\n  margin-bottom: var(--cem-dev-server-spacing-sm);\\n  margin-inline-start: var(--cem-dev-server-spacing-md);\\n\\n  summary {\\n    cursor: pointer;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--background--color--secondary--default);\\n    border-radius: 6px;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n    list-style: none;\\n    transition: background 200ms cubic-bezier(0.4, 0, 0.2, 1);\\n    display: flex;\\n    align-items: center;\\n    gap: 0.5rem;\\n    user-select: none;\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--background--color--secondary--hover);\\n    }\\n\\n    \\u0026::-webkit-details-marker {\\n      display: none;\\n    }\\n\\n    \\u0026::before {\\n      content: '\\\\25B8';\\n      display: inline-block;\\n      transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n      color: var(--pf-t--global--text--color--subtle);\\n    }\\n  }\\n\\n  \\u0026[open] summary::before {\\n    transform: rotate(90deg);\\n  }\\n}\\n\\n.nav-element-title {\\n  user-select: none;\\n}\\n\\n.nav-demo-list {\\n  list-style: none;\\n  padding: 0;\\n  margin: var(--cem-dev-server-spacing-sm) 0 0 0;\\n  display: grid;\\n  gap: var(--cem-dev-server-spacing-xs);\\n}\\n\\n.nav-demo-link {\\n  color: var(--cem-dev-server-text-primary);\\n  text-decoration: none;\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  padding-inline-start: calc(var(--cem-dev-server-spacing-md) * 2);\\n  background: var(--cem-dev-server-bg-tertiary);\\n  border-radius: var(--cem-dev-server-border-radius);\\n  display: block;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  transition: background 0.2s, color 0.2s;\\n\\n  \\u0026:hover {\\n    background: var(--cem-dev-server-accent-color);\\n    color: var(--pf-t--global--text--color--on-brand);\\n\\n    .nav-package-name {\\n      color: rgba(255, 255, 255, 0.8);\\n    }\\n  }\\n\\n  \\u0026[aria-current=\\"page\\"] {\\n    background: var(--cem-dev-server-accent-color);\\n    color: var(--pf-t--global--text--color--on-brand);\\n\\n    .nav-package-name {\\n      color: rgba(255, 255, 255, 0.8);\\n    }\\n  }\\n}\\n\\n.nav-package-name {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n/* Info button popover triggers in knobs - override plain button padding */\\ncem-pf-v6-popover cem-pf-v6-button[variant=\\"plain\\"] {\\n  --cem-pf-v6-c-button--m-plain--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-button--m-plain--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-button--MinWidth: auto;\\n}\\n\\n/* Knob description content (slotted in form group helper) */\\ncem-pf-v6-form-group [slot=\\"helper\\"] {\\n  p {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n  }\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n\\n  a {\\n    color: var(--cem-dev-server-accent-color);\\n    text-decoration: none;\\n\\n    \\u0026:hover {\\n      text-decoration: underline;\\n    }\\n  }\\n\\n  strong {\\n    font-weight: 600;\\n  }\\n\\n  ul, ol {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n    padding-left: var(--cem-dev-server-spacing-lg);\\n  }\\n}\\n\\n/* Syntax highlighting (chroma - themable via CSS custom properties) */\\ncem-pf-v6-form-group [slot=\\"helper\\"] {\\n  .chroma {\\n    background-color: var(--cem-dev-server-bg-tertiary);\\n    padding: var(--cem-dev-server-spacing-sm);\\n    border-radius: var(--cem-dev-server-border-radius);\\n    overflow-x: auto;\\n\\n    \\u0026 .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }\\n    \\u0026 .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; }\\n    \\u0026 .hl { background-color: var(--cem-dev-server-syntax-highlight) }\\n    \\u0026 .lnt,\\n    \\u0026 .ln {\\n      white-space: pre;\\n      user-select: none;\\n      margin-right: 0.4em;\\n      padding: 0 0.4em 0 0.4em;\\n      color: var(--cem-dev-server-text-muted);\\n    }\\n    \\u0026 .line { display: flex; }\\n\\n    /* Keywords */\\n    \\u0026 .k,\\n    \\u0026 .kc,\\n    \\u0026 .kd,\\n    \\u0026 .kn,\\n    \\u0026 .kp,\\n    \\u0026 .kr {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n    \\u0026 .kt { color: var(--cem-dev-server-syntax-type); font-weight: bold; }\\n\\n    /* Names */\\n    \\u0026 .na,\\n    \\u0026 .nb,\\n    \\u0026 .no,\\n    \\u0026 .nv,\\n    \\u0026 .vc,\\n    \\u0026 .vg,\\n    \\u0026 .vi {\\n      color: var(--cem-dev-server-syntax-name);\\n    }\\n    \\u0026 .bp { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .nc { color: var(--cem-dev-server-syntax-class); font-weight: bold; }\\n    \\u0026 .nd { color: var(--cem-dev-server-syntax-decorator); font-weight: bold; }\\n    \\u0026 .ni,\\n    \\u0026 .ss {\\n      color: var(--cem-dev-server-syntax-special);\\n    }\\n    \\u0026 .ne,\\n    \\u0026 .nl {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n    \\u0026 .nf { color: var(--cem-dev-server-syntax-function); font-weight: bold; }\\n    \\u0026 .nn { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .nt { color: var(--cem-dev-server-syntax-tag) }\\n\\n    /* Strings */\\n    \\u0026 .s,\\n    \\u0026 .sa,\\n    \\u0026 .sb,\\n    \\u0026 .sc,\\n    \\u0026 .dl,\\n    \\u0026 .sd,\\n    \\u0026 .s2,\\n    \\u0026 .se,\\n    \\u0026 .sh,\\n    \\u0026 .si,\\n    \\u0026 .sx,\\n    \\u0026 .s1 {\\n      color: var(--cem-dev-server-syntax-string);\\n    }\\n    \\u0026 .sr { color: var(--cem-dev-server-syntax-tag) }\\n\\n    /* Numbers */\\n    \\u0026 .m,\\n    \\u0026 .mb,\\n    \\u0026 .mf,\\n    \\u0026 .mh,\\n    \\u0026 .mi,\\n    \\u0026 .il,\\n    \\u0026 .mo {\\n      color: var(--cem-dev-server-syntax-number);\\n    }\\n\\n    /* Operators */\\n    \\u0026 .o,\\n    \\u0026 .ow {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n\\n    /* Comments */\\n    \\u0026 .c,\\n    \\u0026 .ch,\\n    \\u0026 .cm,\\n    \\u0026 .c1 {\\n      color: var(--cem-dev-server-text-muted);\\n      font-style: italic;\\n    }\\n    \\u0026 .cs,\\n    \\u0026 .cp,\\n    \\u0026 .cpf {\\n      color: var(--cem-dev-server-text-muted);\\n      font-weight: bold;\\n      font-style: italic;\\n    }\\n\\n    /* Errors */\\n    \\u0026 .err {\\n      color: var(--cem-dev-server-syntax-error);\\n      background-color: var(--cem-dev-server-syntax-error-bg);\\n    }\\n\\n    /* Generics */\\n    \\u0026 .gd {\\n      color: var(--cem-dev-server-syntax-deleted);\\n      background-color: var(--cem-dev-server-syntax-deleted-bg);\\n    }\\n    \\u0026 .ge { font-style: italic; }\\n    \\u0026 .gr { color: var(--cem-dev-server-syntax-error) }\\n    \\u0026 .gh { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gi {\\n      color: var(--cem-dev-server-syntax-inserted);\\n      background-color: var(--cem-dev-server-syntax-inserted-bg);\\n    }\\n    \\u0026 .go { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gp { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gs { font-weight: bold; }\\n    \\u0026 .gu { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gt { color: var(--cem-dev-server-syntax-error) }\\n    \\u0026 .gl { text-decoration: underline; }\\n    \\u0026 .w { color: var(--cem-dev-server-text-muted) }\\n  }\\n}\\n\\n/* Events tab styling - Primary-detail layout */\\n.events-wrapper {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n#event-drawer {\\n  flex: 1;\\n  min-height: 0;\\n}\\n\\n/* Event list (primary panel) */\\n#event-list {\\n  overflow-y: auto;\\n  height: 100%;\\n}\\n\\n.event-list-item {\\n  /* Reset button styles */\\n  appearance: none;\\n  background: none;\\n  border: none;\\n  border-left: 3px solid transparent;\\n  margin: 0;\\n  font: inherit;\\n  color: inherit;\\n  text-align: inherit;\\n  width: 100%;\\n\\n  /* Component styles */\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-sm);\\n  align-items: center;\\n  cursor: pointer;\\n  transition: background 100ms ease-in-out, border-color 100ms ease-in-out;\\n\\n  cem-pf-v6-label {\\n    flex-shrink: 0;\\n  }\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--background--color--primary--hover);\\n  }\\n\\n  \\u0026:focus {\\n    outline: 2px solid var(--pf-t--global--border--color--clicked);\\n    outline-offset: -2px;\\n  }\\n\\n  \\u0026.selected {\\n    background: var(--pf-t--global--background--color--action--plain--selected);\\n    border-left-color: var(--pf-t--global--border--color--brand--default);\\n  }\\n}\\n\\n.event-time,\\n.event-element {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.event-time {\\n  color: var(--cem-dev-server-text-muted);\\n  flex-shrink: 0;\\n  font-size: 11px;\\n}\\n\\n.event-element {\\n  color: var(--cem-dev-server-text-muted);\\n  font-weight: 400;\\n}\\n\\n/* Event detail panel */\\n.event-detail-header-content {\\n  padding: var(--cem-dev-server-spacing-md);\\n  border-bottom: var(--cem-dev-server-border-width) solid var(--cem-dev-server-border-color);\\n}\\n\\n.event-detail-name {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-lg);\\n  font-weight: 600;\\n  color: var(--cem-dev-server-text-primary);\\n}\\n\\n.event-detail-summary {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-secondary);\\n  line-height: 1.5;\\n  white-space: pre-wrap;\\n}\\n\\n.event-detail-description {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-secondary);\\n  line-height: 1.5;\\n  white-space: pre-wrap;\\n}\\n\\n.event-detail-meta {\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-md);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.event-detail-time {\\n  color: var(--cem-dev-server-text-muted);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n}\\n\\n.event-detail-element {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n}\\n\\n.event-detail-properties-heading {\\n  margin: var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  font-size: var(--cem-dev-server-font-size-base);\\n  font-weight: 600;\\n  color: var(--cem-dev-server-text-primary);\\n}\\n\\n.event-detail-properties {\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  background: var(--cem-dev-server-bg-secondary);\\n  border: var(--cem-dev-server-border-width) solid var(--cem-dev-server-border-color);\\n  border-radius: var(--cem-dev-server-border-radius);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: 12px;\\n  line-height: 1.6;\\n  margin: 0 var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md);\\n}\\n\\n.event-property-tree {\\n  list-style: none;\\n  padding: 0;\\n  margin: 0;\\n\\n  \\u0026.nested {\\n    padding-left: 1.5em;\\n    margin-top: 0.25em;\\n  }\\n}\\n\\n.property-item {\\n  padding: 0.125em 0;\\n}\\n\\n.property-key {\\n  color: var(--cem-dev-server-accent-color);\\n  font-weight: 500;\\n}\\n\\n.property-colon {\\n  color: var(--cem-dev-server-text-muted);\\n}\\n\\n.property-value {\\n  \\u0026.null,\\n  \\u0026.undefined {\\n    color: var(--cem-dev-server-text-muted);\\n    font-style: italic;\\n  }\\n\\n  \\u0026.boolean {\\n    color: var(--cem-dev-server-color-boolean);\\n  }\\n\\n  \\u0026.number {\\n    color: var(--cem-dev-server-color-number);\\n  }\\n\\n  \\u0026.string {\\n    color: var(--cem-dev-server-color-string);\\n  }\\n\\n  \\u0026.array,\\n  \\u0026.object {\\n    color: var(--cem-dev-server-text-secondary);\\n  }\\n}\\n\\n#debug-modal {\\n  container-type: inline-size;\\n}\\n"`));
var cem_serve_chrome_default = s;

// elements/cem-serve-chrome/cem-serve-chrome.ts
import "../cem-color-scheme-toggle/cem-color-scheme-toggle.js";
import "../cem-drawer/cem-drawer.js";
import "../cem-health-panel/cem-health-panel.js";
import "../cem-manifest-browser/cem-manifest-browser.js";
import "../cem-reconnection-content/cem-reconnection-content.js";
import "../cem-serve-demo/cem-serve-demo.js";
import "../cem-serve-knob-group/cem-serve-knob-group.js";
import "../cem-serve-knobs/cem-serve-knobs.js";
import "../cem-transform-error-overlay/cem-transform-error-overlay.js";
import "../cem-pf-v6-alert/cem-pf-v6-alert.js";
import "../cem-pf-v6-alert-group/cem-pf-v6-alert-group.js";
import "../cem-pf-v6-button/cem-pf-v6-button.js";
import "../cem-pf-v6-card/cem-pf-v6-card.js";
import "../cem-pf-v6-badge/cem-pf-v6-badge.js";
import "../cem-pf-v6-dropdown/cem-pf-v6-dropdown.js";
import "../cem-pf-v6-expandable-section/cem-pf-v6-expandable-section.js";
import "../cem-pf-v6-label/cem-pf-v6-label.js";
import "../cem-pf-v6-masthead/cem-pf-v6-masthead.js";
import "../cem-pf-v6-modal/cem-pf-v6-modal.js";
import "../cem-pf-v6-nav-group/cem-pf-v6-nav-group.js";
import "../cem-pf-v6-nav-item/cem-pf-v6-nav-item.js";
import "../cem-pf-v6-nav-link/cem-pf-v6-nav-link.js";
import "../cem-pf-v6-nav-list/cem-pf-v6-nav-list.js";
import "../cem-pf-v6-navigation/cem-pf-v6-navigation.js";
import "../cem-pf-v6-page-main/cem-pf-v6-page-main.js";
import "../cem-pf-v6-page-sidebar/cem-pf-v6-page-sidebar.js";
import "../cem-pf-v6-page/cem-pf-v6-page.js";
import "../cem-pf-v6-popover/cem-pf-v6-popover.js";
import "../cem-pf-v6-select/cem-pf-v6-select.js";
import "../cem-pf-v6-skip-to-content/cem-pf-v6-skip-to-content.js";
import "../cem-pf-v6-switch/cem-pf-v6-switch.js";
import "../cem-pf-v6-tab/cem-pf-v6-tab.js";
import "../cem-pf-v6-tabs/cem-pf-v6-tabs.js";
import "../cem-pf-v6-text-input-group/cem-pf-v6-text-input-group.js";
import "../cem-pf-v6-text-input/cem-pf-v6-text-input.js";
import "../cem-pf-v6-toggle-group-item/cem-pf-v6-toggle-group-item.js";
import "../cem-pf-v6-toggle-group/cem-pf-v6-toggle-group.js";
import "../cem-pf-v6-toolbar-group/cem-pf-v6-toolbar-group.js";
import "../cem-pf-v6-toolbar-item/cem-pf-v6-toolbar-item.js";
import "../cem-pf-v6-toolbar/cem-pf-v6-toolbar.js";
import "../cem-pf-v6-tree-item/cem-pf-v6-tree-item.js";
import "../cem-pf-v6-tree-view/cem-pf-v6-tree-view.js";
var CEMReloadClient;
var StatePersistence;
var CemLogsEvent = class extends Event {
  logs;
  constructor(logs) {
    super("cem:logs", { bubbles: true });
    this.logs = logs;
  }
};
var _hasDescription_dec, _sidebar_dec, _tabsSelected_dec, _drawerHeight_dec, _drawer_dec, _knobs_dec, _sourceURL_dec, _canonicalURL_dec, _packageName_dec, _demoTitle_dec, _primaryTagName_dec, _a, _CemServeChrome_decorators, _demoInfoTemplate, _demoGroupTemplate, _demoListTemplate, _logEntryTemplate, _eventEntryTemplate, _init, _primaryTagName, _demoTitle, _packageName, _canonicalURL, _sourceURL, _knobs, _drawer, _drawerHeight, _tabsSelected, _sidebar, _hasDescription, _CemServeChrome_instances, $_fn, $$_fn, _logContainer, _drawerOpen, _initialLogsFetched, _isInitialLoad, _debugData, _elementEventMap, _manifest, _capturedEvents, _maxCapturedEvents, _eventList, _eventDetailHeader, _eventDetailBody, _selectedEventId, _eventsFilterValue, _eventsFilterDebounceTimer, _eventTypeFilters, _elementFilters, _discoveredElements, _handleLogsEvent, _handleTreeExpand, _handleTreeCollapse, _handleTreeSelect, _copyLogsFeedbackTimeout, _copyDebugFeedbackTimeout, _copyEventsFeedbackTimeout, _logsFilterValue, _logsFilterDebounceTimer, _logLevelFilters, _logLevelDropdown, _observer, _wsClient, initWsClient_fn, renderSourceButton_fn, _modulesReady, loadClientModules_fn, fetchDebugInfo_fn, populateDemoUrls_fn, setupLogListener_fn, filterLogs_fn, getLogTypeFromEntry_fn, loadLogFilterState_fn, syncCheckboxStates_fn, saveLogFilterState_fn, _handleLogFilterChange, copyLogs_fn, setupDebugOverlay_fn, setupFooterDrawer_fn, detectBrowser_fn, copyDebugInfo_fn, renderLogs_fn, getLogBadge_fn, applyLogLabelAttrs_fn, scrollLatestIntoView_fn, scrollLogsToBottom_fn, migrateFromLocalStorageIfNeeded_fn, setupColorSchemeToggle_fn, applyColorScheme_fn, setupKnobCoordination_fn, _onKnobChange, _onKnobClear, getKnobTarget_fn, getKnobTypeFromEvent_fn, getKnobTypeFromClearEvent_fn, setupTreeStatePersistence_fn, applyTreeState_fn, setupSidebarStatePersistence_fn, findTreeItemById_fn, getTreeNodeId_fn, discoverElementEvents_fn, setupEventCapture_fn, attachEventListeners_fn, observeDemoMutations_fn, _handleElementEvent, getEventDocumentation_fn, findTypeDeclaration_fn, captureEvent_fn, extractEventProperties_fn, renderEvent_fn, selectEvent_fn, buildPropertiesForDisplay_fn, buildPropertyTree_fn, scrollEventsToBottom_fn, isEventsTabActive_fn, filterEvents_fn, eventMatchesTextFilter_fn, getEventRecordById_fn, updateEventFilters_fn, _handleEventTypeFilterChange, _handleElementFilterChange, loadEventFiltersFromStorage_fn, saveEventFilters_fn, clearEvents_fn, copyEvents_fn, setupEventListeners_fn;
_CemServeChrome_decorators = [customElement("cem-serve-chrome")];
var _CemServeChrome = class _CemServeChrome extends (_a = LitElement, _primaryTagName_dec = [property({ attribute: "primary-tag-name" })], _demoTitle_dec = [property({ attribute: "demo-title" })], _packageName_dec = [property({ attribute: "package-name" })], _canonicalURL_dec = [property({ attribute: "canonical-url" })], _sourceURL_dec = [property({ attribute: "source-url" })], _knobs_dec = [property()], _drawer_dec = [property()], _drawerHeight_dec = [property({ attribute: "drawer-height" })], _tabsSelected_dec = [property({ attribute: "tabs-selected" })], _sidebar_dec = [property()], _hasDescription_dec = [property({ type: Boolean, attribute: "has-description" })], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeChrome_instances);
    __privateAdd(this, _primaryTagName, __runInitializers(_init, 8, this, "")), __runInitializers(_init, 11, this);
    __privateAdd(this, _demoTitle, __runInitializers(_init, 12, this, "")), __runInitializers(_init, 15, this);
    __privateAdd(this, _packageName, __runInitializers(_init, 16, this, "")), __runInitializers(_init, 19, this);
    __privateAdd(this, _canonicalURL, __runInitializers(_init, 20, this, "")), __runInitializers(_init, 23, this);
    __privateAdd(this, _sourceURL, __runInitializers(_init, 24, this, "")), __runInitializers(_init, 27, this);
    __privateAdd(this, _knobs, __runInitializers(_init, 28, this, "")), __runInitializers(_init, 31, this);
    __privateAdd(this, _drawer, __runInitializers(_init, 32, this, "")), __runInitializers(_init, 35, this);
    __privateAdd(this, _drawerHeight, __runInitializers(_init, 36, this, "")), __runInitializers(_init, 39, this);
    __privateAdd(this, _tabsSelected, __runInitializers(_init, 40, this, "")), __runInitializers(_init, 43, this);
    __privateAdd(this, _sidebar, __runInitializers(_init, 44, this, "")), __runInitializers(_init, 47, this);
    __privateAdd(this, _hasDescription, __runInitializers(_init, 48, this, false)), __runInitializers(_init, 51, this);
    __privateAdd(this, _logContainer, null);
    __privateAdd(this, _drawerOpen, false);
    __privateAdd(this, _initialLogsFetched, false);
    __privateAdd(this, _isInitialLoad, true);
    __privateAdd(this, _debugData, null);
    // Element event tracking
    __privateAdd(this, _elementEventMap, null);
    __privateAdd(this, _manifest, null);
    __privateAdd(this, _capturedEvents, []);
    __privateAdd(this, _maxCapturedEvents, 1e3);
    __privateAdd(this, _eventList, null);
    __privateAdd(this, _eventDetailHeader, null);
    __privateAdd(this, _eventDetailBody, null);
    __privateAdd(this, _selectedEventId, null);
    __privateAdd(this, _eventsFilterValue, "");
    __privateAdd(this, _eventsFilterDebounceTimer, null);
    __privateAdd(this, _eventTypeFilters, /* @__PURE__ */ new Set());
    __privateAdd(this, _elementFilters, /* @__PURE__ */ new Set());
    __privateAdd(this, _discoveredElements, /* @__PURE__ */ new Set());
    // Event listener references for cleanup
    __privateAdd(this, _handleLogsEvent, null);
    __privateAdd(this, _handleTreeExpand, null);
    __privateAdd(this, _handleTreeCollapse, null);
    __privateAdd(this, _handleTreeSelect, null);
    // Timeout IDs for cleanup
    __privateAdd(this, _copyLogsFeedbackTimeout, null);
    __privateAdd(this, _copyDebugFeedbackTimeout, null);
    __privateAdd(this, _copyEventsFeedbackTimeout, null);
    // Log filter state
    __privateAdd(this, _logsFilterValue, "");
    __privateAdd(this, _logsFilterDebounceTimer, null);
    __privateAdd(this, _logLevelFilters, /* @__PURE__ */ new Set(["info", "warn", "error", "success", "debug", "trace"]));
    __privateAdd(this, _logLevelDropdown, null);
    // Watch for dynamically added elements
    /* c8 ignore start - MutationObserver callback tested via integration */
    __privateAdd(this, _observer, new MutationObserver((mutations) => {
      let needsUpdate = false;
      for (const mutation of mutations) {
        for (const node of mutation.addedNodes) {
          if (node instanceof HTMLElement) {
            const tagName = node.tagName.toLowerCase();
            if (__privateGet(this, _elementEventMap)?.has(tagName) && !node.dataset.cemEventsAttached) {
              const eventInfo = __privateGet(this, _elementEventMap).get(tagName);
              for (const eventName of eventInfo.eventNames) {
                node.addEventListener(eventName, __privateGet(this, _handleElementEvent), { capture: true });
              }
              node.dataset.cemEventsAttached = "true";
              __privateGet(this, _discoveredElements).add(tagName);
              needsUpdate = true;
            }
          }
        }
      }
      if (needsUpdate) {
        __privateMethod(this, _CemServeChrome_instances, updateEventFilters_fn).call(this);
      }
    }));
    /* c8 ignore stop */
    __privateAdd(this, _wsClient);
    /** Resolves when client-only modules are loaded. */
    __privateAdd(this, _modulesReady, __privateMethod(this, _CemServeChrome_instances, loadClientModules_fn).call(this));
    __privateAdd(this, _handleLogFilterChange, (event) => {
      const { value, checked } = event;
      if (checked) {
        __privateGet(this, _logLevelFilters).add(value);
      } else {
        __privateGet(this, _logLevelFilters).delete(value);
      }
      __privateMethod(this, _CemServeChrome_instances, saveLogFilterState_fn).call(this);
      __privateMethod(this, _CemServeChrome_instances, filterLogs_fn).call(this, __privateGet(this, _logsFilterValue));
    });
    __privateAdd(this, _onKnobChange, (event) => {
      const target = __privateMethod(this, _CemServeChrome_instances, getKnobTarget_fn).call(this, event);
      if (!target) {
        console.warn("[cem-serve-chrome] Could not find knob target info in event path");
        return;
      }
      const { tagName, instanceIndex } = target;
      const demo = this.demo;
      if (!demo) return;
      const knobType = __privateMethod(this, _CemServeChrome_instances, getKnobTypeFromEvent_fn).call(this, event);
      const success = demo.applyKnobChange(
        knobType,
        event.name,
        event.value,
        tagName,
        instanceIndex
      );
      if (!success) {
        console.warn("[cem-serve-chrome] Failed to apply knob change:", {
          type: knobType,
          name: event.name,
          tagName,
          instanceIndex
        });
      }
    });
    __privateAdd(this, _onKnobClear, (event) => {
      const target = __privateMethod(this, _CemServeChrome_instances, getKnobTarget_fn).call(this, event);
      if (!target) {
        console.warn("[cem-serve-chrome] Could not find knob target info in event path");
        return;
      }
      const { tagName, instanceIndex } = target;
      const demo = this.demo;
      if (!demo) return;
      const knobType = __privateMethod(this, _CemServeChrome_instances, getKnobTypeFromClearEvent_fn).call(this, event);
      const clearValue = knobType === "property" ? void 0 : "";
      const success = demo.applyKnobChange(
        knobType,
        event.name,
        clearValue,
        tagName,
        instanceIndex
      );
      if (!success) {
        console.warn("[cem-serve-chrome] Failed to clear knob:", {
          type: knobType,
          name: event.name,
          tagName,
          instanceIndex
        });
      }
    });
    __privateAdd(this, _handleElementEvent, (event) => {
      const element = event.currentTarget;
      if (!(element instanceof HTMLElement)) return;
      const tagName = element.tagName.toLowerCase();
      const eventInfo = __privateGet(this, _elementEventMap)?.get(tagName);
      if (!eventInfo || !eventInfo.eventNames.has(event.type)) return;
      __privateGet(this, _discoveredElements).add(tagName);
      __privateMethod(this, _CemServeChrome_instances, captureEvent_fn).call(this, event, element, tagName, eventInfo);
    });
    __privateAdd(this, _handleEventTypeFilterChange, (event) => {
      const { value, checked } = event;
      if (!value) return;
      if (checked) {
        __privateGet(this, _eventTypeFilters).add(value);
      } else {
        __privateGet(this, _eventTypeFilters).delete(value);
      }
      __privateMethod(this, _CemServeChrome_instances, saveEventFilters_fn).call(this);
      __privateMethod(this, _CemServeChrome_instances, filterEvents_fn).call(this, __privateGet(this, _eventsFilterValue));
    });
    __privateAdd(this, _handleElementFilterChange, (event) => {
      const { value, checked } = event;
      if (!value) return;
      if (checked) {
        __privateGet(this, _elementFilters).add(value);
      } else {
        __privateGet(this, _elementFilters).delete(value);
      }
      __privateMethod(this, _CemServeChrome_instances, saveEventFilters_fn).call(this);
      __privateMethod(this, _CemServeChrome_instances, filterEvents_fn).call(this, __privateGet(this, _eventsFilterValue));
    });
  }
  get demo() {
    return this.querySelector("cem-serve-demo");
  }
  render() {
    return html`
      <link rel="stylesheet" href="/__cem/cem-pf-v6-c-description-list.css">
      <link rel="stylesheet" href="/__cem/pf-lightdom.css">

      <cem-pf-v6-page ?sidebar-collapsed=${this.sidebar === "collapsed"}>
        <cem-pf-v6-skip-to-content slot="skip-to-content"
                               href="#main-content">
          Skip to content
        </cem-pf-v6-skip-to-content>

        <cem-pf-v6-masthead slot="masthead">
          <a class="masthead-logo"
             href="/"
             slot="logo">
            <img alt="CEM Dev Server"
                 src="/__cem/logo.svg">
            ${this.packageName ? html`<h1>${this.packageName}</h1>` : nothing}
          </a>
          <cem-pf-v6-toolbar slot="toolbar">
            <cem-pf-v6-toolbar-group variant="action-group">
              ${__privateMethod(this, _CemServeChrome_instances, renderSourceButton_fn).call(this)}
              <cem-pf-v6-toolbar-item>
                <cem-pf-v6-button id="debug-info"
                              variant="plain"
                              aria-label="Debug info">
                  <svg width="16"
                       height="16"
                       viewBox="0 0 16 16"
                       fill="currentColor"
                       role="presentation">
                    <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
                    <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
                  </svg>
                </cem-pf-v6-button>
              </cem-pf-v6-toolbar-item>
              <cem-pf-v6-toolbar-item>
                <cem-pf-v6-toggle-group class="color-scheme-toggle"
                                    aria-label="Color scheme">
                  <cem-pf-v6-toggle-group-item value="light">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="Light mode">
                      <path d="M8 11a3 3 0 1 1 0-6 3 3 0 0 1 0 6zm0 1a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z"/>
                    </svg>
                  </cem-pf-v6-toggle-group-item>
                  <cem-pf-v6-toggle-group-item value="dark">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="Dark mode">
                      <path d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278zM4.858 1.311A7.269 7.269 0 0 0 1.025 7.71c0 4.02 3.279 7.276 7.319 7.276a7.316 7.316 0 0 0 5.205-2.162c-.337.042-.68.063-1.029.063-4.61 0-8.343-3.714-8.343-8.29 0-1.167.242-2.278.681-3.286z"/>
                    </svg>
                  </cem-pf-v6-toggle-group-item>
                  <cem-pf-v6-toggle-group-item value="system">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="System preference">
                      <path d="M0 1.5A1.5 1.5 0 0 1 1.5 0h13A1.5 1.5 0 0 1 16 1.5v8a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 9.5v-8zM1.5 1a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-8a.5.5 0 0 0-.5-.5h-13z"/>
                      <path d="M2.5 12h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1 0-1zm0 2h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1 0-1z"/>
                    </svg>
                  </cem-pf-v6-toggle-group-item>
                </cem-pf-v6-toggle-group>
              </cem-pf-v6-toolbar-item>
            </cem-pf-v6-toolbar-group>
          </cem-pf-v6-toolbar>
        </cem-pf-v6-masthead>

        <cem-pf-v6-page-sidebar slot="sidebar"
                            ?expanded=${this.sidebar === "expanded"}
                            ?collapsed=${this.sidebar !== "expanded"}>
          <slot name="navigation"></slot>
        </cem-pf-v6-page-sidebar>

        <cem-pf-v6-page-main slot="main" id="main-content">
          <slot></slot>
          <footer class="pf-m-sticky-bottom">
            <div class="footer-description${this.hasDescription ? "" : " empty"}">
              <slot name="description"></slot>
            </div>
            <cem-drawer ?open=${this.drawer === "expanded"}
                        drawer-height="${this.drawerHeight || "400"}">
              <cem-pf-v6-tabs selected="${this.tabsSelected || "0"}">
                <cem-pf-v6-tab title="Knobs">
                  <div id="knobs-container">
                    <slot name="knobs">
                      <p class="knobs-empty">No knobs available for this element.</p>
                    </slot>
                  </div>
                </cem-pf-v6-tab>
                <cem-pf-v6-tab title="Manifest Browser">
                  <cem-manifest-browser>
                    <slot name="manifest-tree" slot="manifest-tree"></slot>
                    <slot name="manifest-name" slot="manifest-name"></slot>
                    <slot name="manifest-details" slot="manifest-details"></slot>
                  </cem-manifest-browser>
                </cem-pf-v6-tab>
                <cem-pf-v6-tab title="Server Logs">
                  <div class="logs-wrapper">
                    <cem-pf-v6-toolbar sticky>
                      <cem-pf-v6-toolbar-group>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-text-input-group id="logs-filter"
                                                  placeholder="Filter logs..."
                                                  icon>
                            <svg slot="icon"
                                 role="presentation"
                                 fill="currentColor"
                                 height="1em"
                                 width="1em"
                                 viewBox="0 0 512 512">
                              <path d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                            </svg>
                          </cem-pf-v6-text-input-group>
                        </cem-pf-v6-toolbar-item>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-dropdown id="log-level-filter"
                                          label="Filter log levels">
                            <span slot="toggle-text">Log Levels</span>
                            <cem-pf-v6-menu-item variant="checkbox" value="error" checked>Errors</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="warn" checked>Warnings</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="success" checked>Success</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="info" checked>Info</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="debug" checked>Debug</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="trace" checked>Trace</cem-pf-v6-menu-item>
                          </cem-pf-v6-dropdown>
                        </cem-pf-v6-toolbar-item>
                      </cem-pf-v6-toolbar-group>
                      <cem-pf-v6-toolbar-group variant="action-group">
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-button id="copy-logs"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M13 0H6a2 2 0 0 0-2 2 2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2 2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm0 13V4a2 2 0 0 0-2-2H5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1zM3 13V4a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>
                            </svg>
                            Copy Logs
                          </cem-pf-v6-button>
                        </cem-pf-v6-toolbar-item>
                      </cem-pf-v6-toolbar-group>
                    </cem-pf-v6-toolbar>
                    <div id="log-container"></div>
                  </div>
                </cem-pf-v6-tab>
                <cem-pf-v6-tab title="Events">
                  <div class="events-wrapper">
                    <cem-pf-v6-toolbar sticky>
                      <cem-pf-v6-toolbar-group>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-text-input-group id="events-filter"
                                                  placeholder="Filter events..."
                                                  icon>
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z"/>
                            </svg>
                          </cem-pf-v6-text-input-group>
                        </cem-pf-v6-toolbar-item>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-dropdown id="event-type-filter"
                                          label="Filter event types">
                            <span slot="toggle-text">Event Types</span>
                          </cem-pf-v6-dropdown>
                        </cem-pf-v6-toolbar-item>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-dropdown id="element-filter"
                                          label="Filter elements">
                            <span slot="toggle-text">Elements</span>
                          </cem-pf-v6-dropdown>
                        </cem-pf-v6-toolbar-item>
                      </cem-pf-v6-toolbar-group>
                      <cem-pf-v6-toolbar-group variant="action-group">
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-button id="clear-events"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                              <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                            </svg>
                            Clear Events
                          </cem-pf-v6-button>
                        </cem-pf-v6-toolbar-item>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-button id="copy-events"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M13 0H6a2 2 0 0 0-2 2 2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2 2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm0 13V4a2 2 0 0 0-2-2H5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1zM3 13V4a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>
                            </svg>
                            Copy Events
                          </cem-pf-v6-button>
                        </cem-pf-v6-toolbar-item>
                      </cem-pf-v6-toolbar-group>
                    </cem-pf-v6-toolbar>
                    <cem-pf-v6-drawer id="event-drawer" expanded>
                      <div id="event-list"></div>
                      <div id="event-detail-header" slot="panel-header"></div>
                      <div id="event-detail-body" slot="panel-body"></div>
                    </cem-pf-v6-drawer>
                  </div>
                </cem-pf-v6-tab>
                <cem-pf-v6-tab title="Health">
                  <cem-health-panel component=${ifDefined(this.primaryTagName)}>
                  </cem-health-panel>
                </cem-pf-v6-tab>
              </cem-pf-v6-tabs>
            </cem-drawer>
          </footer>
        </cem-pf-v6-page-main>
      </cem-pf-v6-page>

      <cem-pf-v6-modal id="debug-modal" variant="large">
        <h2 slot="header">Debug Information</h2>
        <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Server Version</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-version">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Server OS</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-os">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Watch Directory</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-watch-dir">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Manifest Size</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-manifest-size">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Demos Found</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-demo-count">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Tag Name</dt>
            <dd class="cem-pf-v6-c-description-list__description">${this.primaryTagName || "-"}</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Demo Title</dt>
            <dd class="cem-pf-v6-c-description-list__description">${this.demoTitle || "-"}</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Browser</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-browser">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">User Agent</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-ua">-</dd>
          </div>
        </dl>
        <div id="demo-urls-container"></div>
        <cem-pf-v6-expandable-section id="debug-importmap-details"
                                  toggle-text="Show Import Map">
          <pre id="debug-importmap">-</pre>
        </cem-pf-v6-expandable-section>
        <div slot="footer" class="button-container">
          <cem-pf-v6-button class="debug-copy" variant="primary">
            Copy Debug Info
          </cem-pf-v6-button>
          <cem-pf-v6-button class="debug-close" variant="secondary">
            Close
          </cem-pf-v6-button>
        </div>
      </cem-pf-v6-modal>

      <!-- Reconnection modal -->
      <cem-pf-v6-modal id="reconnection-modal" variant="large">
        <h2 slot="header">Development Server Disconnected</h2>
        <cem-reconnection-content id="reconnection-content"></cem-reconnection-content>
        <cem-pf-v6-button id="reload-button"
                      slot="footer"
                      variant="primary">Reload Page</cem-pf-v6-button>
        <cem-pf-v6-button id="retry-button"
                      slot="footer"
                      variant="secondary">Retry Now</cem-pf-v6-button>
      </cem-pf-v6-modal>

      <!-- Transform error overlay -->
      <cem-transform-error-overlay id="error-overlay">
      </cem-transform-error-overlay>
    `;
  }
  async connectedCallback() {
    await __privateGet(this, _modulesReady);
    super.connectedCallback();
    if (__privateGet(this, _wsClient) == null) {
      __privateMethod(this, _CemServeChrome_instances, initWsClient_fn).call(this);
    }
    __privateMethod(this, _CemServeChrome_instances, migrateFromLocalStorageIfNeeded_fn).call(this);
  }
  firstUpdated() {
    __privateMethod(this, _CemServeChrome_instances, setupDebugOverlay_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupColorSchemeToggle_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupFooterDrawer_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupLogListener_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupKnobCoordination_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupTreeStatePersistence_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupSidebarStatePersistence_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupEventCapture_fn).call(this).then(() => {
      __privateMethod(this, _CemServeChrome_instances, setupEventListeners_fn).call(this);
    });
    __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "reload-button")?.addEventListener("click", () => {
      window.location.reload();
    });
    __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "retry-button")?.addEventListener("click", () => {
      __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "reconnection-modal")?.close();
      __privateGet(this, _wsClient).retry();
    });
    __privateGet(this, _wsClient).init();
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("knob:attribute-change", __privateGet(this, _onKnobChange));
    this.removeEventListener("knob:property-change", __privateGet(this, _onKnobChange));
    this.removeEventListener("knob:css-property-change", __privateGet(this, _onKnobChange));
    this.removeEventListener("knob:attribute-clear", __privateGet(this, _onKnobClear));
    this.removeEventListener("knob:property-clear", __privateGet(this, _onKnobClear));
    this.removeEventListener("knob:css-property-clear", __privateGet(this, _onKnobClear));
    if (__privateGet(this, _handleTreeExpand)) {
      this.removeEventListener("expand", __privateGet(this, _handleTreeExpand));
    }
    if (__privateGet(this, _handleTreeCollapse)) {
      this.removeEventListener("collapse", __privateGet(this, _handleTreeCollapse));
    }
    if (__privateGet(this, _handleTreeSelect)) {
      this.removeEventListener("select", __privateGet(this, _handleTreeSelect));
    }
    if (__privateGet(this, _handleLogsEvent)) {
      window.removeEventListener("cem:logs", __privateGet(this, _handleLogsEvent));
    }
    if (__privateGet(this, _copyLogsFeedbackTimeout)) {
      clearTimeout(__privateGet(this, _copyLogsFeedbackTimeout));
      __privateSet(this, _copyLogsFeedbackTimeout, null);
    }
    if (__privateGet(this, _copyDebugFeedbackTimeout)) {
      clearTimeout(__privateGet(this, _copyDebugFeedbackTimeout));
      __privateSet(this, _copyDebugFeedbackTimeout, null);
    }
    if (__privateGet(this, _copyEventsFeedbackTimeout)) {
      clearTimeout(__privateGet(this, _copyEventsFeedbackTimeout));
      __privateSet(this, _copyEventsFeedbackTimeout, null);
    }
    __privateGet(this, _observer).disconnect();
    if (__privateGet(this, _wsClient)) {
      __privateGet(this, _wsClient).destroy();
    }
  }
};
_init = __decoratorStart(_a);
_demoInfoTemplate = new WeakMap();
_demoGroupTemplate = new WeakMap();
_demoListTemplate = new WeakMap();
_logEntryTemplate = new WeakMap();
_eventEntryTemplate = new WeakMap();
_primaryTagName = new WeakMap();
_demoTitle = new WeakMap();
_packageName = new WeakMap();
_canonicalURL = new WeakMap();
_sourceURL = new WeakMap();
_knobs = new WeakMap();
_drawer = new WeakMap();
_drawerHeight = new WeakMap();
_tabsSelected = new WeakMap();
_sidebar = new WeakMap();
_hasDescription = new WeakMap();
_CemServeChrome_instances = new WeakSet();
$_fn = function(id) {
  return this.shadowRoot?.getElementById(id);
};
$$_fn = function(selector) {
  return this.shadowRoot?.querySelectorAll(selector) ?? [];
};
_logContainer = new WeakMap();
_drawerOpen = new WeakMap();
_initialLogsFetched = new WeakMap();
_isInitialLoad = new WeakMap();
_debugData = new WeakMap();
_elementEventMap = new WeakMap();
_manifest = new WeakMap();
_capturedEvents = new WeakMap();
_maxCapturedEvents = new WeakMap();
_eventList = new WeakMap();
_eventDetailHeader = new WeakMap();
_eventDetailBody = new WeakMap();
_selectedEventId = new WeakMap();
_eventsFilterValue = new WeakMap();
_eventsFilterDebounceTimer = new WeakMap();
_eventTypeFilters = new WeakMap();
_elementFilters = new WeakMap();
_discoveredElements = new WeakMap();
_handleLogsEvent = new WeakMap();
_handleTreeExpand = new WeakMap();
_handleTreeCollapse = new WeakMap();
_handleTreeSelect = new WeakMap();
_copyLogsFeedbackTimeout = new WeakMap();
_copyDebugFeedbackTimeout = new WeakMap();
_copyEventsFeedbackTimeout = new WeakMap();
_logsFilterValue = new WeakMap();
_logsFilterDebounceTimer = new WeakMap();
_logLevelFilters = new WeakMap();
_logLevelDropdown = new WeakMap();
_observer = new WeakMap();
_wsClient = new WeakMap();
initWsClient_fn = function() {
  __privateSet(this, _wsClient, new CEMReloadClient({
    jitterMax: 1e3,
    overlayThreshold: 15,
    badgeFadeDelay: 2e3,
    /* c8 ignore start - WebSocket callbacks tested via integration */
    callbacks: {
      onOpen: () => {
        __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "reconnection-modal")?.close();
      },
      onError: (errorData) => {
        if (errorData?.title && errorData?.message) {
          console.error("[cem-serve] Server error:", errorData);
          __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "error-overlay")?.show(errorData.title, errorData.message, errorData.file);
        } else {
          console.error("[cem-serve] WebSocket error:", errorData);
        }
      },
      onReconnecting: ({ attempt, delay }) => {
        if (attempt >= 15) {
          __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "reconnection-modal")?.showModal();
          __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "reconnection-content")?.updateRetryInfo(attempt, delay);
        }
      },
      onReload: () => {
        const errorOverlay = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "error-overlay");
        if (errorOverlay?.hasAttribute("open")) {
          errorOverlay.hide();
        }
        window.location.reload();
      },
      onShutdown: () => {
        __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "reconnection-modal")?.showModal();
        __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "reconnection-content")?.updateRetryInfo(30, 3e4);
      },
      onLogs: (logs) => {
        window.dispatchEvent(new CemLogsEvent(logs));
      }
    }
    /* c8 ignore stop */
  }));
};
renderSourceButton_fn = function() {
  if (!this.sourceURL) return nothing;
  let label = "Version Control";
  let path = "M5.854 4.854a.5.5 0 1 0-.708-.708l-3.5 3.5a.5.5 0 0 0 0 .708l3.5 3.5a.5.5 0 0 0 .708-.708L2.707 8l3.147-3.146zm4.292 0a.5.5 0 0 1 .708-.708l3.5 3.5a.5.5 0 0 1 0 .708l-3.5 3.5a.5.5 0 0 1-.708-.708L13.293 8l-3.147-3.146z";
  if (this.sourceURL.includes("github.com")) {
    label = "GitHub.com";
    path = "M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z";
  } else if (this.sourceURL.includes("gitlab.com")) {
    label = "GitLab";
    path = "m15.734 6.1-.022-.058L13.534.358a.568.568 0 0 0-.563-.356.583.583 0 0 0-.328.122.582.582 0 0 0-.193.294l-1.47 4.499H5.025l-1.47-4.5A.572.572 0 0 0 3.360.174a.572.572 0 0 0-.328-.172.582.582 0 0 0-.563.357L.29 6.04l-.022.057A4.044 4.044 0 0 0 1.61 10.77l.007.006.02.014 3.318 2.485 1.64 1.242 1 .755a.673.673 0 0 0 .814 0l1-.755 1.64-1.242 3.338-2.5.009-.007a4.05 4.05 0 0 0 1.34-4.668Z";
  } else if (this.sourceURL.includes("bitbucket.org")) {
    label = "Bitbucket";
    path = "M0 1.5A1.5 1.5 0 0 1 1.5 0h13A1.5 1.5 0 0 1 16 1.5v13a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 14.5v-13zM2.5 3a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-11zm5.038 1.435a.5.5 0 0 1 .924 0l1.42 3.37H8.78l-.243.608-.243-.608H5.082l1.42-3.37zM8 9.143l-.743 1.857H4.743L6.076 7.608 8 9.143z";
  }
  return html`
      <cem-pf-v6-toolbar-item>
        <cem-pf-v6-button href="${this.sourceURL}"
                      target="_blank"
                      rel="noopener noreferrer"
                      variant="plain"
                      aria-label="View source file">
          <svg aria-label="${label}"
               width="1rem"
               height="1rem"
               fill="currentColor"
               viewBox="0 0 16 16">
            <path d="${path}"/>
          </svg>
        </cem-pf-v6-button>
      </cem-pf-v6-toolbar-item>
    `;
};
_modulesReady = new WeakMap();
loadClientModules_fn = function() {
  return Promise.all([
    // @ts-ignore -- plain JS modules served at runtime by Go server
    import("/__cem/websocket-client.js"),
    // @ts-ignore
    import("/__cem/state-persistence.js")
  ]).then(([ws, sp]) => {
    CEMReloadClient = ws.CEMReloadClient;
    StatePersistence = sp.StatePersistence;
    import("/__cem/health-badges.js").catch((e) => console.error("[cem-serve] Failed to load health-badges:", e));
  }).catch((e) => {
    console.error("[cem-serve] Failed to load client modules:", e);
    CEMReloadClient ??= class {
      init() {
      }
      retry() {
      }
      destroy() {
      }
    };
    StatePersistence ??= {
      getState: () => ({ colorScheme: "system" }),
      updateState() {
      },
      getTreeState: () => ({ expanded: [], selected: null }),
      setTreeState() {
      },
      updateTreeState() {
      },
      migrateFromLocalStorage() {
      }
    };
  });
};
fetchDebugInfo_fn = async function() {
  const browserEl = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-browser");
  const uaEl = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-ua");
  if (browserEl) {
    const browser = __privateMethod(this, _CemServeChrome_instances, detectBrowser_fn).call(this);
    browserEl.textContent = browser;
  }
  if (uaEl) {
    uaEl.textContent = navigator.userAgent;
  }
  const data = await fetch("/__cem/debug").then((res) => res.json()).catch((err) => {
    console.error("[cem-serve-chrome] Failed to fetch debug info:", err);
  });
  if (!data) return;
  const versionEl = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-version");
  const osEl = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-os");
  const watchDirEl = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-watch-dir");
  const manifestSizeEl = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-manifest-size");
  const demoCountEl = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-demo-count");
  const demoUrlsContainer = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "demo-urls-container");
  const importMapEl = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-importmap");
  if (versionEl) versionEl.textContent = data.version || "-";
  if (osEl) osEl.textContent = data.os || "-";
  if (watchDirEl) watchDirEl.textContent = data.watchDir || "-";
  if (manifestSizeEl) manifestSizeEl.textContent = data.manifestSize || "-";
  if (demoCountEl) demoCountEl.textContent = data.demoCount || "0";
  if (demoUrlsContainer) {
    __privateMethod(this, _CemServeChrome_instances, populateDemoUrls_fn).call(this, demoUrlsContainer, data.demos);
  }
  if (importMapEl && data.importMap) {
    const importMapJSON = JSON.stringify(data.importMap, null, 2);
    importMapEl.textContent = importMapJSON;
    data.importMapJSON = importMapJSON;
  } else if (importMapEl) {
    importMapEl.textContent = "No import map generated";
  }
  __privateSet(this, _debugData, data);
};
populateDemoUrls_fn = function(container, demos) {
  if (!demos?.length) {
    container.textContent = "No demos found in manifest";
    return;
  }
  const currentTagName = this.primaryTagName || "";
  const isOnDemoPage = !!currentTagName;
  if (isOnDemoPage) {
    const currentDemo = demos.find((demo) => demo.tagName === currentTagName);
    if (!currentDemo) {
      container.textContent = "Current demo not found in manifest";
      return;
    }
    const fragment = __privateGet(_CemServeChrome, _demoInfoTemplate).content.cloneNode(true);
    fragment.querySelector('[data-field="tagName"]').textContent = currentDemo.tagName;
    fragment.querySelector('[data-field="canonicalURL"]').textContent = currentDemo.canonicalURL;
    fragment.querySelector('[data-field="localRoute"]').textContent = currentDemo.localRoute;
    const descriptionGroup = fragment.querySelector('[data-field-group="description"]');
    if (currentDemo.description) {
      fragment.querySelector('[data-field="description"]').textContent = currentDemo.description;
    } else {
      descriptionGroup?.remove();
    }
    const filepathGroup = fragment.querySelector('[data-field-group="filepath"]');
    if (currentDemo.filepath) {
      fragment.querySelector('[data-field="filepath"]').textContent = currentDemo.filepath;
    } else {
      filepathGroup?.remove();
    }
    container.replaceChildren(fragment);
  } else {
    const listFragment = __privateGet(_CemServeChrome, _demoListTemplate).content.cloneNode(true);
    const groupsContainer = listFragment.querySelector('[data-container="groups"]');
    for (const demo of demos) {
      const groupFragment = __privateGet(_CemServeChrome, _demoGroupTemplate).content.cloneNode(true);
      groupFragment.querySelector('[data-field="tagName"]').textContent = demo.tagName;
      groupFragment.querySelector('[data-field="description"]').textContent = demo.description || "(no description)";
      groupFragment.querySelector('[data-field="canonicalURL"]').textContent = demo.canonicalURL;
      groupFragment.querySelector('[data-field="localRoute"]').textContent = demo.localRoute;
      const filepathGroup = groupFragment.querySelector('[data-field-group="filepath"]');
      if (demo.filepath) {
        groupFragment.querySelector('[data-field="filepath"]').textContent = demo.filepath;
      } else {
        filepathGroup?.remove();
      }
      groupsContainer.appendChild(groupFragment);
    }
    container.replaceChildren(listFragment);
  }
};
setupLogListener_fn = function() {
  __privateSet(this, _logContainer, __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "log-container"));
  const logsFilter = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "logs-filter");
  if (logsFilter) {
    logsFilter.addEventListener("input", () => {
      const { value = "" } = logsFilter;
      clearTimeout(__privateGet(this, _logsFilterDebounceTimer));
      __privateSet(this, _logsFilterDebounceTimer, setTimeout(() => {
        __privateMethod(this, _CemServeChrome_instances, filterLogs_fn).call(this, value);
      }, 300));
    });
  }
  __privateSet(this, _logLevelDropdown, this.shadowRoot?.querySelector("#log-level-filter") ?? null);
  if (__privateGet(this, _logLevelDropdown)) {
    requestAnimationFrame(() => {
      __privateMethod(this, _CemServeChrome_instances, loadLogFilterState_fn).call(this);
    });
    __privateGet(this, _logLevelDropdown).addEventListener("select", __privateGet(this, _handleLogFilterChange));
  }
  __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "copy-logs")?.addEventListener("click", () => {
    __privateMethod(this, _CemServeChrome_instances, copyLogs_fn).call(this);
  });
  __privateSet(this, _handleLogsEvent, ((event) => {
    const logs = event.logs;
    if (logs) {
      __privateMethod(this, _CemServeChrome_instances, renderLogs_fn).call(this, logs);
    }
  }));
  window.addEventListener("cem:logs", __privateGet(this, _handleLogsEvent));
};
filterLogs_fn = function(query) {
  __privateSet(this, _logsFilterValue, query.toLowerCase());
  if (!__privateGet(this, _logContainer)) return;
  for (const entry of __privateGet(this, _logContainer).children) {
    const text = entry.textContent?.toLowerCase() ?? "";
    const textMatch = !__privateGet(this, _logsFilterValue) || text.includes(__privateGet(this, _logsFilterValue));
    const logType = __privateMethod(this, _CemServeChrome_instances, getLogTypeFromEntry_fn).call(this, entry);
    const levelMatch = __privateGet(this, _logLevelFilters).has(logType);
    entry.hidden = !(textMatch && levelMatch);
  }
};
getLogTypeFromEntry_fn = function(entry) {
  for (const cls of entry.classList) {
    if (["info", "warning", "error", "success", "debug", "trace"].includes(cls)) {
      return cls === "warning" ? "warn" : cls;
    }
  }
  return "info";
};
loadLogFilterState_fn = function() {
  try {
    const saved = localStorage.getItem("cem-serve-log-filters");
    if (saved) {
      __privateSet(this, _logLevelFilters, new Set(JSON.parse(saved)));
    }
  } catch (e) {
    console.debug("[cem-serve-chrome] localStorage unavailable, using default log filters");
  }
  __privateMethod(this, _CemServeChrome_instances, syncCheckboxStates_fn).call(this);
};
syncCheckboxStates_fn = function() {
  if (!__privateGet(this, _logLevelDropdown)) return;
  const menuItems = __privateGet(this, _logLevelDropdown).querySelectorAll("cem-pf-v6-menu-item");
  menuItems.forEach((item) => {
    const value = item.value;
    item.checked = __privateGet(this, _logLevelFilters).has(value);
  });
};
saveLogFilterState_fn = function() {
  try {
    localStorage.setItem(
      "cem-serve-log-filters",
      JSON.stringify([...__privateGet(this, _logLevelFilters)])
    );
  } catch (e) {
  }
};
_handleLogFilterChange = new WeakMap();
copyLogs_fn = async function() {
  if (!__privateGet(this, _logContainer)) return;
  const logs = Array.from(__privateGet(this, _logContainer).children).filter((entry) => !entry.hidden).map((entry) => {
    const type = entry.querySelector('[data-field="label"]')?.textContent?.trim() || "INFO";
    const time = entry.querySelector('[data-field="time"]')?.textContent?.trim() || "";
    const message = entry.querySelector('[data-field="message"]')?.textContent?.trim() || "";
    return `[${type}] ${time} ${message}`;
  }).join("\n");
  if (!logs) return;
  try {
    await navigator.clipboard.writeText(logs);
    const btn = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "copy-logs");
    if (btn) {
      const textNode = Array.from(btn.childNodes).find(
        (n) => n.nodeType === Node.TEXT_NODE && (n.textContent?.trim().length ?? 0) > 0
      );
      if (textNode) {
        const original = textNode.textContent;
        textNode.textContent = "Copied!";
        if (__privateGet(this, _copyLogsFeedbackTimeout)) {
          clearTimeout(__privateGet(this, _copyLogsFeedbackTimeout));
        }
        __privateSet(this, _copyLogsFeedbackTimeout, setTimeout(() => {
          if (this.isConnected && textNode.parentNode) {
            textNode.textContent = original;
          }
          __privateSet(this, _copyLogsFeedbackTimeout, null);
        }, 2e3));
      }
    }
  } catch (err) {
    console.error("[cem-serve-chrome] Failed to copy logs:", err);
  }
};
setupDebugOverlay_fn = function() {
  const debugButton = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-info");
  const debugModal = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "debug-modal");
  const debugClose = this.shadowRoot?.querySelector(".debug-close");
  const debugCopy = this.shadowRoot?.querySelector(".debug-copy");
  if (debugButton && debugModal) {
    debugButton.addEventListener("click", () => {
      __privateMethod(this, _CemServeChrome_instances, fetchDebugInfo_fn).call(this);
      debugModal.showModal();
    });
    debugClose?.addEventListener("click", () => debugModal.close());
    debugCopy?.addEventListener("click", () => {
      __privateMethod(this, _CemServeChrome_instances, copyDebugInfo_fn).call(this);
    });
  }
};
setupFooterDrawer_fn = function() {
  const drawer = this.shadowRoot?.querySelector("cem-drawer");
  const tabs = this.shadowRoot?.querySelector("cem-pf-v6-tabs");
  if (!drawer || !tabs) return;
  __privateSet(this, _drawerOpen, drawer.open);
  drawer.addEventListener("change", (e) => {
    __privateSet(this, _drawerOpen, e.open);
    StatePersistence.updateState({
      drawer: { open: e.open }
    });
    if (e.open) {
      __privateMethod(this, _CemServeChrome_instances, scrollLogsToBottom_fn).call(this);
    }
  });
  drawer.addEventListener("resize", (e) => {
    drawer.setAttribute("drawer-height", e.height);
    StatePersistence.updateState({
      drawer: { height: e.height }
    });
  });
  tabs.addEventListener("change", (e) => {
    StatePersistence.updateState({
      tabs: { selectedIndex: e.selectedIndex }
    });
    if (e.selectedIndex === 2 && drawer.open) {
      __privateMethod(this, _CemServeChrome_instances, scrollLogsToBottom_fn).call(this);
    }
    if (e.selectedIndex === 3 && drawer.open) {
      __privateMethod(this, _CemServeChrome_instances, scrollEventsToBottom_fn).call(this);
    }
  });
};
detectBrowser_fn = function() {
  const ua = navigator.userAgent;
  if (ua.includes("Firefox/")) {
    const match = ua.match(/Firefox\/(\d+)/);
    return match ? `Firefox ${match[1]}` : "Firefox";
  } else if (ua.includes("Edg/")) {
    const match = ua.match(/Edg\/(\d+)/);
    return match ? `Edge ${match[1]}` : "Edge";
  } else if (ua.includes("Chrome/")) {
    const match = ua.match(/Chrome\/(\d+)/);
    return match ? `Chrome ${match[1]}` : "Chrome";
  } else if (ua.includes("Safari/") && !ua.includes("Chrome")) {
    const match = ua.match(/Version\/(\d+)/);
    return match ? `Safari ${match[1]}` : "Safari";
  }
  return "Unknown";
};
copyDebugInfo_fn = async function() {
  const info = Array.from(__privateMethod(this, _CemServeChrome_instances, $$_fn).call(this, "#debug-modal dl dt")).map((dt) => {
    const dd = dt.nextElementSibling;
    if (dd && dd.tagName === "DD") {
      return `${dt.textContent}: ${dd.textContent}`;
    }
    return "";
  }).join("\n");
  let importMapSection = "";
  if (__privateGet(this, _debugData)?.importMapJSON) {
    importMapSection = `
${"=".repeat(40)}
Import Map:
${"=".repeat(40)}
${__privateGet(this, _debugData).importMapJSON}
`;
  }
  const debugText = `CEM Serve Debug Information
${"=".repeat(40)}
${info}${importMapSection}
${"=".repeat(40)}
Generated: ${(/* @__PURE__ */ new Date()).toISOString()}`;
  try {
    await navigator.clipboard.writeText(debugText);
    const copyButton = this.shadowRoot?.querySelector(".debug-copy");
    if (copyButton) {
      const originalText = copyButton.textContent;
      copyButton.textContent = "Copied!";
      if (__privateGet(this, _copyDebugFeedbackTimeout)) {
        clearTimeout(__privateGet(this, _copyDebugFeedbackTimeout));
      }
      __privateSet(this, _copyDebugFeedbackTimeout, setTimeout(() => {
        if (this.isConnected && copyButton.parentNode) {
          copyButton.textContent = originalText;
        }
        __privateSet(this, _copyDebugFeedbackTimeout, null);
      }, 2e3));
    }
  } catch (err) {
    console.error("[cem-serve-chrome] Failed to copy debug info:", err);
  }
};
renderLogs_fn = function(logs) {
  if (!__privateGet(this, _logContainer)) return;
  const logElements = logs.map((log) => {
    const fragment = __privateGet(_CemServeChrome, _logEntryTemplate).content.cloneNode(true);
    const date = new Date(log.date);
    const time = date.toLocaleTimeString();
    const container = fragment.querySelector('[data-field="container"]');
    container.classList.add(log.type);
    container.setAttribute("data-log-id", log.date);
    const typeLabel = __privateMethod(this, _CemServeChrome_instances, getLogBadge_fn).call(this, log.type);
    const searchText = `${typeLabel} ${time} ${log.message}`.toLowerCase();
    const textMatch = !__privateGet(this, _logsFilterValue) || searchText.includes(__privateGet(this, _logsFilterValue));
    const logTypeForFilter = log.type === "warning" ? "warn" : log.type;
    const levelMatch = __privateGet(this, _logLevelFilters).has(logTypeForFilter);
    if (!(textMatch && levelMatch)) {
      container.setAttribute("hidden", "");
    }
    const label = fragment.querySelector('[data-field="label"]');
    label.textContent = __privateMethod(this, _CemServeChrome_instances, getLogBadge_fn).call(this, log.type);
    __privateMethod(this, _CemServeChrome_instances, applyLogLabelAttrs_fn).call(this, label, log.type);
    const timeEl = fragment.querySelector('[data-field="time"]');
    timeEl.setAttribute("datetime", log.date);
    timeEl.textContent = time;
    fragment.querySelector('[data-field="message"]').textContent = log.message;
    return fragment;
  });
  if (!__privateGet(this, _initialLogsFetched)) {
    __privateGet(this, _logContainer).replaceChildren(...logElements);
    __privateSet(this, _initialLogsFetched, true);
    if (__privateGet(this, _drawerOpen)) {
      __privateMethod(this, _CemServeChrome_instances, scrollLatestIntoView_fn).call(this);
    }
  } else {
    __privateGet(this, _logContainer).append(...logElements);
    if (__privateGet(this, _drawerOpen)) {
      __privateMethod(this, _CemServeChrome_instances, scrollLatestIntoView_fn).call(this);
    }
  }
};
getLogBadge_fn = function(type) {
  switch (type) {
    case "info":
      return "Info";
    case "warning":
      return "Warn";
    case "error":
      return "Error";
    case "success":
      return "OK";
    case "debug":
      return "Debug";
    case "trace":
      return "Trace";
    default:
      return type.toUpperCase();
  }
};
applyLogLabelAttrs_fn = function(label, type) {
  switch (type) {
    case "info":
      return label.setAttribute("status", "info");
    case "warning":
      return label.setAttribute("status", "warning");
    case "error":
      return label.setAttribute("status", "danger");
    case "success":
      return label.setAttribute("status", "success");
    case "debug":
      return label.setAttribute("color", "purple");
    case "trace":
      return label.setAttribute("color", "grey");
    default:
      label.setAttribute("color", "grey");
  }
};
scrollLatestIntoView_fn = function() {
  if (!__privateGet(this, _logContainer)) return;
  requestAnimationFrame(() => {
    const lastLog = __privateGet(this, _logContainer).lastElementChild;
    if (lastLog) {
      lastLog.scrollIntoView({ behavior: "auto", block: "end" });
    }
  });
};
scrollLogsToBottom_fn = function() {
  if (!__privateGet(this, _logContainer)) return;
  if (__privateGet(this, _isInitialLoad)) {
    __privateMethod(this, _CemServeChrome_instances, scrollLatestIntoView_fn).call(this);
  } else {
    setTimeout(() => {
      __privateMethod(this, _CemServeChrome_instances, scrollLatestIntoView_fn).call(this);
    }, 350);
  }
};
migrateFromLocalStorageIfNeeded_fn = function() {
  try {
    const hasLocalStorage = localStorage.getItem("cem-serve-color-scheme") !== null || localStorage.getItem("cem-serve-drawer-open") !== null || localStorage.getItem("cem-serve-drawer-height") !== null || localStorage.getItem("cem-serve-active-tab") !== null;
    if (hasLocalStorage) {
      const migrated = localStorage.getItem("cem-serve-migrated-to-cookies");
      if (!migrated) {
        StatePersistence.migrateFromLocalStorage();
        localStorage.setItem("cem-serve-migrated-to-cookies", "true");
        setTimeout(() => window.location.reload(), 100);
      }
    }
  } catch (e) {
  }
};
setupColorSchemeToggle_fn = function() {
  const toggleGroup = this.shadowRoot?.querySelector(".color-scheme-toggle");
  if (!toggleGroup) return;
  const state = StatePersistence.getState();
  __privateMethod(this, _CemServeChrome_instances, applyColorScheme_fn).call(this, state.colorScheme);
  const items = toggleGroup.querySelectorAll("cem-pf-v6-toggle-group-item");
  items.forEach((item) => {
    if (item.value === state.colorScheme) {
      item.setAttribute("selected", "");
    }
  });
  toggleGroup.addEventListener("cem-pf-v6-toggle-group-change", (e) => {
    const scheme = e.value;
    __privateMethod(this, _CemServeChrome_instances, applyColorScheme_fn).call(this, scheme);
    StatePersistence.updateState({ colorScheme: scheme });
  });
};
applyColorScheme_fn = function(scheme) {
  const body = document.body;
  switch (scheme) {
    case "light":
      body.style.colorScheme = "light";
      break;
    case "dark":
      body.style.colorScheme = "dark";
      break;
    case "system":
    default:
      body.style.colorScheme = "light dark";
      break;
  }
};
setupKnobCoordination_fn = function() {
  this.addEventListener("knob:attribute-change", __privateGet(this, _onKnobChange));
  this.addEventListener("knob:property-change", __privateGet(this, _onKnobChange));
  this.addEventListener("knob:css-property-change", __privateGet(this, _onKnobChange));
  this.addEventListener("knob:attribute-clear", __privateGet(this, _onKnobClear));
  this.addEventListener("knob:property-clear", __privateGet(this, _onKnobClear));
  this.addEventListener("knob:css-property-clear", __privateGet(this, _onKnobClear));
};
_onKnobChange = new WeakMap();
_onKnobClear = new WeakMap();
getKnobTarget_fn = function(event) {
  const defaultTagName = this.primaryTagName || "";
  if (event.composedPath) {
    for (const element of event.composedPath()) {
      if (!(element instanceof Element)) continue;
      if (element.dataset?.isElementKnob === "true") {
        const tagName = element.dataset.tagName || defaultTagName;
        let instanceIndex = Number.parseInt(element.dataset.instanceIndex ?? "", 10);
        if (Number.isNaN(instanceIndex)) instanceIndex = 0;
        return { tagName, instanceIndex };
      }
    }
  }
  return { tagName: defaultTagName, instanceIndex: 0 };
};
getKnobTypeFromEvent_fn = function(event) {
  switch (event.type) {
    case "knob:attribute-change":
      return "attribute";
    case "knob:property-change":
      return "property";
    case "knob:css-property-change":
      return "css-property";
    default:
      return "unknown";
  }
};
getKnobTypeFromClearEvent_fn = function(event) {
  switch (event.type) {
    case "knob:attribute-clear":
      return "attribute";
    case "knob:property-clear":
      return "property";
    case "knob:css-property-clear":
      return "css-property";
    default:
      return "unknown";
  }
};
setupTreeStatePersistence_fn = function() {
  __privateSet(this, _handleTreeExpand, (e) => {
    if (e.target?.tagName !== "CEM-PF-V6-TREE-ITEM") return;
    const nodeId = __privateMethod(this, _CemServeChrome_instances, getTreeNodeId_fn).call(this, e.target);
    const treeState = StatePersistence.getTreeState();
    if (!treeState.expanded.includes(nodeId)) {
      treeState.expanded.push(nodeId);
      StatePersistence.setTreeState(treeState);
    }
  });
  this.addEventListener("expand", __privateGet(this, _handleTreeExpand));
  __privateSet(this, _handleTreeCollapse, (e) => {
    if (e.target?.tagName !== "CEM-PF-V6-TREE-ITEM") return;
    const nodeId = __privateMethod(this, _CemServeChrome_instances, getTreeNodeId_fn).call(this, e.target);
    const treeState = StatePersistence.getTreeState();
    const index = treeState.expanded.indexOf(nodeId);
    if (index > -1) {
      treeState.expanded.splice(index, 1);
      StatePersistence.setTreeState(treeState);
    }
  });
  this.addEventListener("collapse", __privateGet(this, _handleTreeCollapse));
  __privateSet(this, _handleTreeSelect, (e) => {
    if (e.target?.tagName !== "CEM-PF-V6-TREE-ITEM") return;
    const nodeId = __privateMethod(this, _CemServeChrome_instances, getTreeNodeId_fn).call(this, e.target);
    StatePersistence.updateTreeState({ selected: nodeId });
  });
  this.addEventListener("select", __privateGet(this, _handleTreeSelect));
  __privateMethod(this, _CemServeChrome_instances, applyTreeState_fn).call(this);
};
applyTreeState_fn = function() {
  const treeState = StatePersistence.getTreeState();
  for (const nodeId of treeState.expanded) {
    const treeItem = __privateMethod(this, _CemServeChrome_instances, findTreeItemById_fn).call(this, nodeId);
    if (treeItem && !treeItem.hasAttribute("expanded")) {
      treeItem.setAttribute("expanded", "");
    }
  }
  if (treeState.selected) {
    const treeItem = __privateMethod(this, _CemServeChrome_instances, findTreeItemById_fn).call(this, treeState.selected);
    if (treeItem && !treeItem.hasAttribute("current")) {
      treeItem.setAttribute("current", "");
    }
  }
};
setupSidebarStatePersistence_fn = function() {
  const page = this.shadowRoot?.querySelector("cem-pf-v6-page");
  if (!page) return;
  page.addEventListener("sidebar-toggle", (event) => {
    const collapsed = !event.expanded;
    StatePersistence.updateState({
      sidebar: { collapsed }
    });
  });
};
findTreeItemById_fn = function(nodeId) {
  const parts = nodeId.split(":");
  const [type, modulePath, tagName, name] = parts;
  let attrSuffix = "";
  if (tagName) {
    attrSuffix += `[data-tag-name="${CSS.escape(tagName)}"]`;
  }
  if (name) {
    attrSuffix += `[data-name="${CSS.escape(name)}"]`;
  }
  let selector = `cem-pf-v6-tree-item[data-type="${CSS.escape(type)}"]`;
  if (modulePath) {
    const escapedModulePath = CSS.escape(modulePath);
    const escapedType = CSS.escape(type);
    const selector1 = `cem-pf-v6-tree-item[data-type="${escapedType}"][data-module-path="${escapedModulePath}"]${attrSuffix}`;
    const selector2 = `cem-pf-v6-tree-item[data-type="${escapedType}"][data-path="${escapedModulePath}"]${attrSuffix}`;
    selector = `${selector1}, ${selector2}`;
  } else {
    selector += attrSuffix;
  }
  return this.querySelector(selector);
};
getTreeNodeId_fn = function(treeItem) {
  const type = treeItem.getAttribute("data-type");
  const modulePath = treeItem.getAttribute("data-module-path") || treeItem.getAttribute("data-path");
  const tagName = treeItem.getAttribute("data-tag-name");
  const name = treeItem.getAttribute("data-name");
  const category = treeItem.getAttribute("data-category");
  const parts = [type];
  if (modulePath) parts.push(modulePath);
  if (tagName) parts.push(tagName);
  if (category) {
    parts.push(category);
  } else if (name) {
    parts.push(name);
  }
  return parts.join(":");
};
discoverElementEvents_fn = async function() {
  try {
    const response = await fetch("/custom-elements.json");
    if (!response.ok) {
      console.warn("[cem-serve-chrome] No manifest available for event discovery");
      return /* @__PURE__ */ new Map();
    }
    const manifest = await response.json();
    __privateSet(this, _manifest, manifest);
    const eventMap = /* @__PURE__ */ new Map();
    for (const module of manifest.modules || []) {
      for (const declaration of module.declarations || []) {
        if (declaration.customElement && declaration.tagName) {
          const tagName = declaration.tagName;
          const events = declaration.events || [];
          if (events.length > 0) {
            const eventNames = new Set(events.map((e) => e.name));
            eventMap.set(tagName, {
              eventNames,
              events
            });
          }
        }
      }
    }
    return eventMap;
  } catch (error) {
    console.warn("[cem-serve-chrome] Error loading manifest for event discovery:", error);
    return /* @__PURE__ */ new Map();
  }
};
setupEventCapture_fn = async function() {
  __privateSet(this, _elementEventMap, await __privateMethod(this, _CemServeChrome_instances, discoverElementEvents_fn).call(this));
  if (__privateGet(this, _elementEventMap).size === 0) return;
  __privateMethod(this, _CemServeChrome_instances, attachEventListeners_fn).call(this);
  __privateMethod(this, _CemServeChrome_instances, updateEventFilters_fn).call(this);
  __privateMethod(this, _CemServeChrome_instances, observeDemoMutations_fn).call(this);
};
attachEventListeners_fn = function() {
  const demo = this.demo;
  if (!demo) return;
  const root = demo.shadowRoot ?? demo;
  for (const [tagName, eventInfo] of __privateGet(this, _elementEventMap)) {
    const elements = root.querySelectorAll(tagName);
    for (const element of elements) {
      for (const eventName of eventInfo.eventNames) {
        element.addEventListener(eventName, __privateGet(this, _handleElementEvent), { capture: true });
      }
      element.dataset.cemEventsAttached = "true";
      __privateGet(this, _discoveredElements).add(tagName);
    }
  }
};
observeDemoMutations_fn = function() {
  const demo = this.demo;
  if (!demo) return;
  const root = demo.shadowRoot ?? demo;
  __privateGet(this, _observer).observe(root, {
    childList: true,
    subtree: true
  });
};
_handleElementEvent = new WeakMap();
getEventDocumentation_fn = function(manifestEvent) {
  if (!manifestEvent) {
    return { summary: null, description: null };
  }
  let summary = manifestEvent.summary || null;
  let description = manifestEvent.description || null;
  if (manifestEvent.type?.text && __privateGet(this, _manifest)) {
    const typeName = manifestEvent.type.text;
    const typeDeclaration = __privateMethod(this, _CemServeChrome_instances, findTypeDeclaration_fn).call(this, typeName);
    if (typeDeclaration) {
      if (!summary && typeDeclaration.summary) {
        summary = typeDeclaration.summary;
      } else if (typeDeclaration.summary && typeDeclaration.summary !== summary) {
        summary = summary ? `${summary}

From ${typeName}: ${typeDeclaration.summary}` : typeDeclaration.summary;
      }
      if (!description && typeDeclaration.description) {
        description = typeDeclaration.description;
      } else if (typeDeclaration.description && typeDeclaration.description !== description) {
        description = description ? `${description}

${typeDeclaration.description}` : typeDeclaration.description;
      }
    }
  }
  return { summary, description };
};
findTypeDeclaration_fn = function(typeName) {
  if (!__privateGet(this, _manifest)) return null;
  for (const module of __privateGet(this, _manifest).modules || []) {
    for (const declaration of module.declarations || []) {
      if (declaration.name === typeName && (declaration.kind === "class" || declaration.kind === "interface")) {
        return declaration;
      }
    }
  }
  return null;
};
captureEvent_fn = function(event, target, tagName, eventInfo) {
  const manifestEvent = eventInfo.events.find((e) => e.name === event.type);
  const eventDocs = __privateMethod(this, _CemServeChrome_instances, getEventDocumentation_fn).call(this, manifestEvent);
  const customProperties = __privateMethod(this, _CemServeChrome_instances, extractEventProperties_fn).call(this, event);
  const eventRecord = {
    id: `${Date.now()}-${Math.random()}`,
    timestamp: /* @__PURE__ */ new Date(),
    eventName: event.type,
    tagName,
    elementId: target.id || null,
    elementClass: target.className || null,
    customProperties,
    manifestType: manifestEvent?.type?.text || null,
    summary: eventDocs.summary,
    description: eventDocs.description,
    bubbles: event.bubbles,
    composed: event.composed,
    cancelable: event.cancelable,
    defaultPrevented: event.defaultPrevented
  };
  __privateGet(this, _capturedEvents).push(eventRecord);
  if (__privateGet(this, _capturedEvents).length > __privateGet(this, _maxCapturedEvents)) {
    __privateGet(this, _capturedEvents).shift();
  }
  __privateMethod(this, _CemServeChrome_instances, renderEvent_fn).call(this, eventRecord);
};
extractEventProperties_fn = function(event) {
  const properties = {};
  const eventPrototypeKeys = new Set(Object.getOwnPropertyNames(Event.prototype));
  const serializeValue = (value) => {
    try {
      return JSON.parse(JSON.stringify(value));
    } catch (e) {
      try {
        return String(value);
      } catch (stringErr) {
        return "[Not serializable]";
      }
    }
  };
  if (event instanceof CustomEvent && event.detail !== void 0) {
    properties.detail = serializeValue(event.detail);
  }
  for (const key of Object.getOwnPropertyNames(event)) {
    if (!eventPrototypeKeys.has(key) && !key.startsWith("_") && !properties.hasOwnProperty(key)) {
      properties[key] = serializeValue(event[key]);
    }
  }
  return properties;
};
renderEvent_fn = function(eventRecord) {
  if (!__privateGet(this, _eventList)) return;
  const fragment = __privateGet(_CemServeChrome, _eventEntryTemplate).content.cloneNode(true);
  const time = eventRecord.timestamp.toLocaleTimeString();
  const container = fragment.querySelector('[data-field="container"]');
  container.dataset.eventId = eventRecord.id;
  container.dataset.eventType = eventRecord.eventName;
  container.dataset.elementType = eventRecord.tagName;
  const textMatch = __privateMethod(this, _CemServeChrome_instances, eventMatchesTextFilter_fn).call(this, eventRecord);
  const typeMatch = __privateGet(this, _eventTypeFilters).size === 0 || __privateGet(this, _eventTypeFilters).has(eventRecord.eventName);
  const elementMatch = __privateGet(this, _elementFilters).size === 0 || __privateGet(this, _elementFilters).has(eventRecord.tagName);
  if (!(textMatch && typeMatch && elementMatch)) {
    container.setAttribute("hidden", "");
  }
  const label = fragment.querySelector('[data-field="label"]');
  label.textContent = eventRecord.eventName;
  label.setAttribute("status", "info");
  const timeEl = fragment.querySelector('[data-field="time"]');
  timeEl.setAttribute("datetime", eventRecord.timestamp.toISOString());
  timeEl.textContent = time;
  const elementEl = fragment.querySelector('[data-field="element"]');
  let elementText = `<${eventRecord.tagName}>`;
  if (eventRecord.elementId) {
    elementText += `#${eventRecord.elementId}`;
  }
  elementEl.textContent = elementText;
  __privateGet(this, _eventList).append(fragment);
  if (!__privateGet(this, _selectedEventId)) {
    __privateMethod(this, _CemServeChrome_instances, selectEvent_fn).call(this, eventRecord.id);
  }
  if (__privateGet(this, _drawerOpen) && __privateMethod(this, _CemServeChrome_instances, isEventsTabActive_fn).call(this)) {
    __privateMethod(this, _CemServeChrome_instances, scrollEventsToBottom_fn).call(this);
  }
};
selectEvent_fn = function(eventId) {
  const eventRecord = __privateMethod(this, _CemServeChrome_instances, getEventRecordById_fn).call(this, eventId);
  if (!eventRecord) return;
  __privateSet(this, _selectedEventId, eventId);
  const allItems = __privateGet(this, _eventList)?.querySelectorAll(".event-list-item");
  allItems?.forEach((item) => {
    if (item.dataset.eventId === eventId) {
      item.classList.add("selected");
      item.setAttribute("aria-selected", "true");
    } else {
      item.classList.remove("selected");
      item.setAttribute("aria-selected", "false");
    }
  });
  if (__privateGet(this, _eventDetailHeader)) {
    __privateGet(this, _eventDetailHeader).innerHTML = "";
    const headerContent = document.createElement("div");
    headerContent.className = "event-detail-header-content";
    const eventName = document.createElement("h3");
    eventName.textContent = eventRecord.eventName;
    eventName.className = "event-detail-name";
    headerContent.appendChild(eventName);
    if (eventRecord.summary) {
      const summary = document.createElement("p");
      summary.textContent = eventRecord.summary;
      summary.className = "event-detail-summary";
      headerContent.appendChild(summary);
    }
    if (eventRecord.description) {
      const description = document.createElement("p");
      description.textContent = eventRecord.description;
      description.className = "event-detail-description";
      headerContent.appendChild(description);
    }
    const meta = document.createElement("div");
    meta.className = "event-detail-meta";
    const timeEl = document.createElement("time");
    timeEl.setAttribute("datetime", eventRecord.timestamp.toISOString());
    timeEl.textContent = eventRecord.timestamp.toLocaleTimeString();
    timeEl.className = "event-detail-time";
    const element = document.createElement("span");
    let elementText = `<${eventRecord.tagName}>`;
    if (eventRecord.elementId) {
      elementText += `#${eventRecord.elementId}`;
    }
    element.textContent = elementText;
    element.className = "event-detail-element";
    meta.appendChild(timeEl);
    meta.appendChild(element);
    headerContent.appendChild(meta);
    __privateGet(this, _eventDetailHeader).appendChild(headerContent);
  }
  if (__privateGet(this, _eventDetailBody)) {
    __privateGet(this, _eventDetailBody).innerHTML = "";
    const propertiesHeading = document.createElement("h4");
    propertiesHeading.textContent = "Properties";
    propertiesHeading.className = "event-detail-properties-heading";
    const propertiesContainer = document.createElement("div");
    propertiesContainer.className = "event-detail-properties";
    const eventProperties = __privateMethod(this, _CemServeChrome_instances, buildPropertiesForDisplay_fn).call(this, eventRecord);
    if (Object.keys(eventProperties).length > 0) {
      propertiesContainer.appendChild(__privateMethod(this, _CemServeChrome_instances, buildPropertyTree_fn).call(this, eventProperties));
    } else {
      propertiesContainer.textContent = "No properties to display";
    }
    __privateGet(this, _eventDetailBody).appendChild(propertiesHeading);
    __privateGet(this, _eventDetailBody).appendChild(propertiesContainer);
  }
};
buildPropertiesForDisplay_fn = function(eventRecord) {
  const properties = {};
  if (eventRecord.customProperties) {
    Object.assign(properties, eventRecord.customProperties);
  }
  properties.bubbles = eventRecord.bubbles;
  properties.cancelable = eventRecord.cancelable;
  properties.defaultPrevented = eventRecord.defaultPrevented;
  properties.composed = eventRecord.composed;
  if (eventRecord.manifestType) {
    properties.type = eventRecord.manifestType;
  }
  return properties;
};
buildPropertyTree_fn = function(obj, depth = 0) {
  const ul = document.createElement("ul");
  ul.className = "event-property-tree";
  if (depth > 0) {
    ul.classList.add("nested");
  }
  for (const [key, value] of Object.entries(obj)) {
    const li = document.createElement("li");
    li.className = "property-item";
    const keySpan = document.createElement("span");
    keySpan.className = "property-key";
    keySpan.textContent = key;
    const colonSpan = document.createElement("span");
    colonSpan.className = "property-colon";
    colonSpan.textContent = ": ";
    li.appendChild(keySpan);
    li.appendChild(colonSpan);
    if (value === null || value === void 0) {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value null";
      valueSpan.textContent = String(value);
      li.appendChild(valueSpan);
    } else if (typeof value === "boolean") {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value boolean";
      valueSpan.textContent = String(value);
      li.appendChild(valueSpan);
    } else if (typeof value === "number") {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value number";
      valueSpan.textContent = String(value);
      li.appendChild(valueSpan);
    } else if (typeof value === "string") {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value string";
      valueSpan.textContent = `"${value}"`;
      li.appendChild(valueSpan);
    } else if (Array.isArray(value)) {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value array";
      valueSpan.textContent = `Array(${value.length})`;
      li.appendChild(valueSpan);
      if (value.length > 0 && depth < 3) {
        const nestedObj = {};
        value.forEach((item, index) => {
          nestedObj[index] = item;
        });
        li.appendChild(__privateMethod(this, _CemServeChrome_instances, buildPropertyTree_fn).call(this, nestedObj, depth + 1));
      }
    } else if (typeof value === "object") {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value object";
      const keys = Object.keys(value);
      valueSpan.textContent = keys.length > 0 ? `Object` : `{}`;
      li.appendChild(valueSpan);
      if (keys.length > 0 && depth < 3) {
        li.appendChild(__privateMethod(this, _CemServeChrome_instances, buildPropertyTree_fn).call(this, value, depth + 1));
      }
    } else {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value";
      valueSpan.textContent = String(value);
      li.appendChild(valueSpan);
    }
    ul.appendChild(li);
  }
  return ul;
};
scrollEventsToBottom_fn = function() {
  if (!__privateGet(this, _eventList)) return;
  requestAnimationFrame(() => {
    const lastEvent = __privateGet(this, _eventList).lastElementChild;
    if (lastEvent) {
      lastEvent.scrollIntoView({ behavior: "auto", block: "end" });
    }
  });
};
isEventsTabActive_fn = function() {
  const tabs = this.shadowRoot?.querySelector("cem-pf-v6-tabs");
  if (!tabs) return false;
  const selectedIndex = parseInt(tabs.getAttribute("selected") || "0", 10);
  return selectedIndex === 3;
};
filterEvents_fn = function(query) {
  __privateSet(this, _eventsFilterValue, query.toLowerCase());
  if (!__privateGet(this, _eventList)) return;
  for (const entry of __privateGet(this, _eventList).children) {
    const eventRecord = __privateMethod(this, _CemServeChrome_instances, getEventRecordById_fn).call(this, entry.dataset.eventId);
    if (!eventRecord) continue;
    const textMatch = __privateMethod(this, _CemServeChrome_instances, eventMatchesTextFilter_fn).call(this, eventRecord);
    const typeMatch = __privateGet(this, _eventTypeFilters).size === 0 || __privateGet(this, _eventTypeFilters).has(eventRecord.eventName);
    const elementMatch = __privateGet(this, _elementFilters).size === 0 || __privateGet(this, _elementFilters).has(eventRecord.tagName);
    entry.hidden = !(textMatch && typeMatch && elementMatch);
  }
};
eventMatchesTextFilter_fn = function(eventRecord) {
  if (!__privateGet(this, _eventsFilterValue)) return true;
  const searchText = [
    eventRecord.tagName,
    eventRecord.eventName,
    eventRecord.elementId || "",
    JSON.stringify(eventRecord.customProperties || {})
  ].join(" ").toLowerCase();
  return searchText.includes(__privateGet(this, _eventsFilterValue));
};
getEventRecordById_fn = function(id) {
  return __privateGet(this, _capturedEvents).find((e) => e.id === id);
};
updateEventFilters_fn = function() {
  const savedPreferences = __privateMethod(this, _CemServeChrome_instances, loadEventFiltersFromStorage_fn).call(this);
  const eventTypeFilter = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "event-type-filter");
  if (eventTypeFilter && __privateGet(this, _elementEventMap)) {
    let menu = eventTypeFilter.querySelector("cem-pf-v6-menu");
    if (!menu) {
      menu = document.createElement("cem-pf-v6-menu");
      eventTypeFilter.appendChild(menu);
    }
    const existingItems = menu.querySelectorAll("cem-pf-v6-menu-item");
    existingItems.forEach((item) => item.remove());
    const allEventTypes = /* @__PURE__ */ new Set();
    for (const [tagName, eventInfo] of __privateGet(this, _elementEventMap)) {
      if (__privateGet(this, _discoveredElements).has(tagName)) {
        for (const eventName of eventInfo.eventNames) {
          allEventTypes.add(eventName);
        }
      }
    }
    if (savedPreferences.eventTypes) {
      __privateSet(this, _eventTypeFilters, savedPreferences.eventTypes.intersection(allEventTypes));
    } else {
      __privateSet(this, _eventTypeFilters, new Set(allEventTypes));
    }
    for (const eventName of allEventTypes) {
      const item = document.createElement("cem-pf-v6-menu-item");
      item.setAttribute("variant", "checkbox");
      item.setAttribute("value", eventName);
      if (__privateGet(this, _eventTypeFilters).has(eventName)) {
        item.setAttribute("checked", "");
      }
      item.textContent = eventName;
      menu.appendChild(item);
    }
  }
  const elementFilter = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "element-filter");
  if (elementFilter && __privateGet(this, _elementEventMap)) {
    let menu = elementFilter.querySelector("cem-pf-v6-menu");
    if (!menu) {
      menu = document.createElement("cem-pf-v6-menu");
      elementFilter.appendChild(menu);
    }
    const existingItems = menu.querySelectorAll("cem-pf-v6-menu-item");
    existingItems.forEach((item) => item.remove());
    const allElements = /* @__PURE__ */ new Set();
    for (const tagName of __privateGet(this, _elementEventMap).keys()) {
      if (__privateGet(this, _discoveredElements).has(tagName)) {
        allElements.add(tagName);
      }
    }
    if (savedPreferences.elements) {
      __privateSet(this, _elementFilters, savedPreferences.elements.intersection(allElements));
    } else {
      __privateSet(this, _elementFilters, new Set(allElements));
    }
    for (const tagName of allElements) {
      const item = document.createElement("cem-pf-v6-menu-item");
      item.setAttribute("variant", "checkbox");
      item.setAttribute("value", tagName);
      if (__privateGet(this, _elementFilters).has(tagName)) {
        item.setAttribute("checked", "");
      }
      item.textContent = `<${tagName}>`;
      menu.appendChild(item);
    }
  }
};
_handleEventTypeFilterChange = new WeakMap();
_handleElementFilterChange = new WeakMap();
loadEventFiltersFromStorage_fn = function() {
  const preferences = {
    eventTypes: null,
    elements: null
  };
  try {
    const savedEventTypes = localStorage.getItem("cem-serve-event-type-filters");
    if (savedEventTypes) {
      preferences.eventTypes = new Set(JSON.parse(savedEventTypes));
    }
    const savedElements = localStorage.getItem("cem-serve-element-filters");
    if (savedElements) {
      preferences.elements = new Set(JSON.parse(savedElements));
    }
  } catch (e) {
    console.debug("[cem-serve-chrome] localStorage unavailable for event filters");
  }
  return preferences;
};
saveEventFilters_fn = function() {
  try {
    localStorage.setItem(
      "cem-serve-event-type-filters",
      JSON.stringify([...__privateGet(this, _eventTypeFilters)])
    );
    localStorage.setItem(
      "cem-serve-element-filters",
      JSON.stringify([...__privateGet(this, _elementFilters)])
    );
  } catch (e) {
  }
};
clearEvents_fn = function() {
  __privateSet(this, _capturedEvents, []);
  __privateSet(this, _selectedEventId, null);
  if (__privateGet(this, _eventList)) {
    __privateGet(this, _eventList).replaceChildren();
  }
  if (__privateGet(this, _eventDetailHeader)) {
    __privateGet(this, _eventDetailHeader).innerHTML = "";
  }
  if (__privateGet(this, _eventDetailBody)) {
    __privateGet(this, _eventDetailBody).innerHTML = "";
  }
};
copyEvents_fn = async function() {
  if (!__privateGet(this, _eventList)) return;
  const visibleEvents = Array.from(__privateGet(this, _eventList).children).filter((entry) => !entry.hidden).map((entry) => {
    const id = entry.dataset.eventId;
    return __privateMethod(this, _CemServeChrome_instances, getEventRecordById_fn).call(this, id);
  }).filter((event) => !!event).map((event) => {
    const time = event.timestamp.toLocaleTimeString();
    const element = event.elementId ? `#${event.elementId}` : event.tagName;
    const props = event.customProperties && Object.keys(event.customProperties).length > 0 ? ` - ${JSON.stringify(event.customProperties)}` : "";
    return `[${time}] <${event.tagName}> ${element} \u2192 ${event.eventName}${props}`;
  }).join("\n");
  if (!visibleEvents) return;
  try {
    await navigator.clipboard.writeText(visibleEvents);
    const btn = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "copy-events");
    if (btn) {
      const textNode = Array.from(btn.childNodes).find(
        (n) => n.nodeType === Node.TEXT_NODE && (n.textContent?.trim().length ?? 0) > 0
      );
      if (textNode) {
        const original = textNode.textContent;
        textNode.textContent = "Copied!";
        if (__privateGet(this, _copyEventsFeedbackTimeout)) {
          clearTimeout(__privateGet(this, _copyEventsFeedbackTimeout));
        }
        __privateSet(this, _copyEventsFeedbackTimeout, setTimeout(() => {
          if (this.isConnected && textNode.parentNode) {
            textNode.textContent = original;
          }
          __privateSet(this, _copyEventsFeedbackTimeout, null);
        }, 2e3));
      }
    }
  } catch (err) {
    console.error("[cem-serve-chrome] Failed to copy events:", err);
  }
};
setupEventListeners_fn = function() {
  __privateSet(this, _eventList, __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "event-list"));
  __privateSet(this, _eventDetailHeader, __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "event-detail-header"));
  __privateSet(this, _eventDetailBody, __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "event-detail-body"));
  if (__privateGet(this, _eventList)) {
    __privateGet(this, _eventList).addEventListener("click", (e) => {
      const listItem = e.target.closest(".event-list-item");
      if (listItem) {
        const eventId = listItem.dataset.eventId;
        if (eventId) {
          __privateMethod(this, _CemServeChrome_instances, selectEvent_fn).call(this, eventId);
        }
      }
    });
  }
  const eventsFilter = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "events-filter");
  if (eventsFilter) {
    eventsFilter.addEventListener("input", (e) => {
      const { value = "" } = e.target;
      clearTimeout(__privateGet(this, _eventsFilterDebounceTimer));
      __privateSet(this, _eventsFilterDebounceTimer, setTimeout(() => {
        __privateMethod(this, _CemServeChrome_instances, filterEvents_fn).call(this, value);
      }, 300));
    });
  }
  const eventTypeFilter = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "event-type-filter");
  if (eventTypeFilter) {
    eventTypeFilter.addEventListener("select", __privateGet(this, _handleEventTypeFilterChange));
  }
  const elementFilter = __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "element-filter");
  if (elementFilter) {
    elementFilter.addEventListener("select", __privateGet(this, _handleElementFilterChange));
  }
  __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "clear-events")?.addEventListener("click", () => {
    __privateMethod(this, _CemServeChrome_instances, clearEvents_fn).call(this);
  });
  __privateMethod(this, _CemServeChrome_instances, $_fn).call(this, "copy-events")?.addEventListener("click", () => {
    __privateMethod(this, _CemServeChrome_instances, copyEvents_fn).call(this);
  });
};
__decorateElement(_init, 4, "primaryTagName", _primaryTagName_dec, _CemServeChrome, _primaryTagName);
__decorateElement(_init, 4, "demoTitle", _demoTitle_dec, _CemServeChrome, _demoTitle);
__decorateElement(_init, 4, "packageName", _packageName_dec, _CemServeChrome, _packageName);
__decorateElement(_init, 4, "canonicalURL", _canonicalURL_dec, _CemServeChrome, _canonicalURL);
__decorateElement(_init, 4, "sourceURL", _sourceURL_dec, _CemServeChrome, _sourceURL);
__decorateElement(_init, 4, "knobs", _knobs_dec, _CemServeChrome, _knobs);
__decorateElement(_init, 4, "drawer", _drawer_dec, _CemServeChrome, _drawer);
__decorateElement(_init, 4, "drawerHeight", _drawerHeight_dec, _CemServeChrome, _drawerHeight);
__decorateElement(_init, 4, "tabsSelected", _tabsSelected_dec, _CemServeChrome, _tabsSelected);
__decorateElement(_init, 4, "sidebar", _sidebar_dec, _CemServeChrome, _sidebar);
__decorateElement(_init, 4, "hasDescription", _hasDescription_dec, _CemServeChrome, _hasDescription);
_CemServeChrome = __decorateElement(_init, 0, "CemServeChrome", _CemServeChrome_decorators, _CemServeChrome);
__publicField(_CemServeChrome, "styles", cem_serve_chrome_default);
// Static templates for DOM cloning (logs, events, debug info)
__privateAdd(_CemServeChrome, _demoInfoTemplate, (() => {
  const t = document.createElement("template");
  t.innerHTML = `
      <h3>Demo Information</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group">
          <dt class="cem-pf-v6-c-description-list__term">Tag Name</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="tagName"></dd>
        </div>
        <div class="cem-pf-v6-c-description-list__group" data-field-group="description">
          <dt class="cem-pf-v6-c-description-list__term">Description</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="description"></dd>
        </div>
        <div class="cem-pf-v6-c-description-list__group" data-field-group="filepath">
          <dt class="cem-pf-v6-c-description-list__term">File Path</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="filepath"></dd>
        </div>
        <div class="cem-pf-v6-c-description-list__group">
          <dt class="cem-pf-v6-c-description-list__term">Canonical URL</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="canonicalURL"></dd>
        </div>
        <div class="cem-pf-v6-c-description-list__group">
          <dt class="cem-pf-v6-c-description-list__term">Local Route</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="localRoute"></dd>
        </div>
      </dl>`;
  return t;
})());
__privateAdd(_CemServeChrome, _demoGroupTemplate, (() => {
  const t = document.createElement("template");
  t.innerHTML = `
      <div class="cem-pf-v6-c-description-list__group">
        <dt class="cem-pf-v6-c-description-list__term" data-field="tagName"></dt>
        <dd class="cem-pf-v6-c-description-list__description">
          <span data-field="description"></span><br>
          <small data-field-group="filepath">File: <span data-field="filepath"></span></small>
          <small>Canonical: <span data-field="canonicalURL"></span></small><br>
          <small>Local: <span data-field="localRoute"></span></small>
        </dd>
      </div>`;
  return t;
})());
__privateAdd(_CemServeChrome, _demoListTemplate, (() => {
  const t = document.createElement("template");
  t.innerHTML = `
      <cem-pf-v6-expandable-section id="debug-demos-section"
                                toggle-text="Show Demos Info">
        <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact" data-container="groups"></dl>
      </cem-pf-v6-expandable-section>`;
  return t;
})());
__privateAdd(_CemServeChrome, _logEntryTemplate, (() => {
  const t = document.createElement("template");
  t.innerHTML = `
      <div class="log-entry" data-field="container">
        <cem-pf-v6-label compact data-field="label"></cem-pf-v6-label>
        <time class="log-time" data-field="time"></time>
        <span class="log-message" data-field="message"></span>
      </div>`;
  return t;
})());
__privateAdd(_CemServeChrome, _eventEntryTemplate, (() => {
  const t = document.createElement("template");
  t.innerHTML = `
      <button class="event-list-item" data-field="container">
        <cem-pf-v6-label compact data-field="label"></cem-pf-v6-label>
        <time class="event-time" data-field="time"></time>
        <span class="event-element" data-field="element"></span>
      </button>`;
  return t;
})());
__runInitializers(_init, 1, _CemServeChrome);
var CemServeChrome = _CemServeChrome;
export {
  CemLogsEvent,
  CemServeChrome
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcbmltcG9ydCB7IGlmRGVmaW5lZCB9IGZyb20gJ2xpdC9kaXJlY3RpdmVzL2lmLWRlZmluZWQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWNocm9tZS5jc3MnO1xuXG5pbXBvcnQgJy4uL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlLmpzJztcbmltcG9ydCAnLi4vY2VtLWRyYXdlci9jZW0tZHJhd2VyLmpzJztcbmltcG9ydCAnLi4vY2VtLWhlYWx0aC1wYW5lbC9jZW0taGVhbHRoLXBhbmVsLmpzJztcbmltcG9ydCAnLi4vY2VtLW1hbmlmZXN0LWJyb3dzZXIvY2VtLW1hbmlmZXN0LWJyb3dzZXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcmVjb25uZWN0aW9uLWNvbnRlbnQvY2VtLXJlY29ubmVjdGlvbi1jb250ZW50LmpzJztcbmltcG9ydCAnLi4vY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8uanMnO1xuaW1wb3J0ICcuLi9jZW0tc2VydmUta25vYi1ncm91cC9jZW0tc2VydmUta25vYi1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1zZXJ2ZS1rbm9icy9jZW0tc2VydmUta25vYnMuanMnO1xuaW1wb3J0ICcuLi9jZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXkvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LWFsZXJ0L2NlbS1wZi12Ni1hbGVydC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1hbGVydC1ncm91cC9jZW0tcGYtdjYtYWxlcnQtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtYnV0dG9uL2NlbS1wZi12Ni1idXR0b24uanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtY2FyZC9jZW0tcGYtdjYtY2FyZC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1iYWRnZS9jZW0tcGYtdjYtYmFkZ2UuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtZHJvcGRvd24vY2VtLXBmLXY2LWRyb3Bkb3duLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbi9jZW0tcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LWxhYmVsL2NlbS1wZi12Ni1sYWJlbC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1tYXN0aGVhZC9jZW0tcGYtdjYtbWFzdGhlYWQuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtbW9kYWwvY2VtLXBmLXY2LW1vZGFsLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LW5hdi1ncm91cC9jZW0tcGYtdjYtbmF2LWdyb3VwLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LW5hdi1pdGVtL2NlbS1wZi12Ni1uYXYtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1uYXYtbGluay9jZW0tcGYtdjYtbmF2LWxpbmsuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtbmF2LWxpc3QvY2VtLXBmLXY2LW5hdi1saXN0LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LW5hdmlnYXRpb24vY2VtLXBmLXY2LW5hdmlnYXRpb24uanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtcGFnZS1tYWluL2NlbS1wZi12Ni1wYWdlLW1haW4uanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtcGFnZS1zaWRlYmFyL2NlbS1wZi12Ni1wYWdlLXNpZGViYXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtcGFnZS9jZW0tcGYtdjYtcGFnZS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1wb3BvdmVyL2NlbS1wZi12Ni1wb3BvdmVyLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXNlbGVjdC9jZW0tcGYtdjYtc2VsZWN0LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXNraXAtdG8tY29udGVudC9jZW0tcGYtdjYtc2tpcC10by1jb250ZW50LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXN3aXRjaC9jZW0tcGYtdjYtc3dpdGNoLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRhYi9jZW0tcGYtdjYtdGFiLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRhYnMvY2VtLXBmLXY2LXRhYnMuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cC9jZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10ZXh0LWlucHV0L2NlbS1wZi12Ni10ZXh0LWlucHV0LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtL2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAvY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10b29sYmFyLWdyb3VwL2NlbS1wZi12Ni10b29sYmFyLWdyb3VwLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRvb2xiYXItaXRlbS9jZW0tcGYtdjYtdG9vbGJhci1pdGVtLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRvb2xiYXIvY2VtLXBmLXY2LXRvb2xiYXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtdHJlZS1pdGVtL2NlbS1wZi12Ni10cmVlLWl0ZW0uanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtdHJlZS12aWV3L2NlbS1wZi12Ni10cmVlLXZpZXcuanMnO1xuXG4vLyBDbGllbnQtb25seSBtb2R1bGVzIGxvYWRlZCBkeW5hbWljYWxseSB0byBhdm9pZCBicmVha2luZyBTU1IuXG4vLyBUaGVzZSBhcmUgcGxhaW4gSlMgbW9kdWxlcyBzZXJ2ZWQgYXQgcnVudGltZSBieSB0aGUgR28gc2VydmVyLlxudHlwZSBDRU1SZWxvYWRDbGllbnRUeXBlID0geyBuZXcob3B0czogYW55KTogYW55IH07XG50eXBlIFN0YXRlUGVyc2lzdGVuY2VUeXBlID0ge1xuICBnZXRTdGF0ZSgpOiBhbnk7XG4gIHVwZGF0ZVN0YXRlKHM6IGFueSk6IHZvaWQ7XG4gIGdldFRyZWVTdGF0ZSgpOiBhbnk7XG4gIHNldFRyZWVTdGF0ZShzOiBhbnkpOiB2b2lkO1xuICB1cGRhdGVUcmVlU3RhdGUoczogYW55KTogdm9pZDtcbiAgbWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2UoKTogdm9pZDtcbn07XG5sZXQgQ0VNUmVsb2FkQ2xpZW50OiBDRU1SZWxvYWRDbGllbnRUeXBlO1xubGV0IFN0YXRlUGVyc2lzdGVuY2U6IFN0YXRlUGVyc2lzdGVuY2VUeXBlO1xuXG5pbnRlcmZhY2UgRXZlbnRJbmZvIHtcbiAgZXZlbnROYW1lczogU2V0PHN0cmluZz47XG4gIGV2ZW50czogQXJyYXk8eyBuYW1lOiBzdHJpbmc7IHR5cGU/OiB7IHRleHQ6IHN0cmluZyB9OyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9Pjtcbn1cblxuaW50ZXJmYWNlIEV2ZW50UmVjb3JkIHtcbiAgaWQ6IHN0cmluZztcbiAgdGltZXN0YW1wOiBEYXRlO1xuICBldmVudE5hbWU6IHN0cmluZztcbiAgdGFnTmFtZTogc3RyaW5nO1xuICBlbGVtZW50SWQ6IHN0cmluZyB8IG51bGw7XG4gIGVsZW1lbnRDbGFzczogc3RyaW5nIHwgbnVsbDtcbiAgY3VzdG9tUHJvcGVydGllczogUmVjb3JkPHN0cmluZywgdW5rbm93bj47XG4gIG1hbmlmZXN0VHlwZTogc3RyaW5nIHwgbnVsbDtcbiAgc3VtbWFyeTogc3RyaW5nIHwgbnVsbDtcbiAgZGVzY3JpcHRpb246IHN0cmluZyB8IG51bGw7XG4gIGJ1YmJsZXM6IGJvb2xlYW47XG4gIGNvbXBvc2VkOiBib29sZWFuO1xuICBjYW5jZWxhYmxlOiBib29sZWFuO1xuICBkZWZhdWx0UHJldmVudGVkOiBib29sZWFuO1xufVxuXG5pbnRlcmZhY2UgRGVidWdEYXRhIHtcbiAgdmVyc2lvbj86IHN0cmluZztcbiAgb3M/OiBzdHJpbmc7XG4gIHdhdGNoRGlyPzogc3RyaW5nO1xuICBtYW5pZmVzdFNpemU/OiBzdHJpbmc7XG4gIGRlbW9Db3VudD86IG51bWJlcjtcbiAgZGVtb3M/OiBBcnJheTx7XG4gICAgdGFnTmFtZTogc3RyaW5nO1xuICAgIGRlc2NyaXB0aW9uPzogc3RyaW5nO1xuICAgIGZpbGVwYXRoPzogc3RyaW5nO1xuICAgIGNhbm9uaWNhbFVSTDogc3RyaW5nO1xuICAgIGxvY2FsUm91dGU6IHN0cmluZztcbiAgfT47XG4gIGltcG9ydE1hcD86IFJlY29yZDxzdHJpbmcsIHVua25vd24+O1xuICBpbXBvcnRNYXBKU09OPzogc3RyaW5nO1xufVxuXG5pbnRlcmZhY2UgTWFuaWZlc3Qge1xuICBtb2R1bGVzPzogQXJyYXk8e1xuICAgIGRlY2xhcmF0aW9ucz86IEFycmF5PHtcbiAgICAgIGN1c3RvbUVsZW1lbnQ/OiBib29sZWFuO1xuICAgICAgdGFnTmFtZT86IHN0cmluZztcbiAgICAgIG5hbWU/OiBzdHJpbmc7XG4gICAgICBraW5kPzogc3RyaW5nO1xuICAgICAgZXZlbnRzPzogQXJyYXk8eyBuYW1lOiBzdHJpbmc7IHR5cGU/OiB7IHRleHQ6IHN0cmluZyB9OyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9PjtcbiAgICB9PjtcbiAgfT47XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gbG9ncyBhcmUgcmVjZWl2ZWRcbiAqL1xuZXhwb3J0IGNsYXNzIENlbUxvZ3NFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgbG9nczogQXJyYXk8eyB0eXBlOiBzdHJpbmc7IGRhdGU6IHN0cmluZzsgbWVzc2FnZTogc3RyaW5nIH0+O1xuICBjb25zdHJ1Y3Rvcihsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT4pIHtcbiAgICBzdXBlcignY2VtOmxvZ3MnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gICAgdGhpcy5sb2dzID0gbG9ncztcbiAgfVxufVxuXG4vKipcbiAqIENFTSBTZXJ2ZSBDaHJvbWUgLSBNYWluIGRlbW8gd3JhcHBlciBjb21wb25lbnRcbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgZGVtbyBjb250ZW50XG4gKiBAc2xvdCBuYXZpZ2F0aW9uIC0gTmF2aWdhdGlvbiBzaWRlYmFyIGNvbnRlbnRcbiAqIEBzbG90IGtub2JzIC0gS25vYiBjb250cm9sc1xuICogQHNsb3QgZGVzY3JpcHRpb24gLSBEZW1vIGRlc2NyaXB0aW9uXG4gKiBAc2xvdCBtYW5pZmVzdC10cmVlIC0gTWFuaWZlc3QgdHJlZSB2aWV3XG4gKiBAc2xvdCBtYW5pZmVzdC1uYW1lIC0gTWFuaWZlc3QgbmFtZSBkaXNwbGF5XG4gKiBAc2xvdCBtYW5pZmVzdC1kZXRhaWxzIC0gTWFuaWZlc3QgZGV0YWlscyBkaXNwbGF5XG4gKlxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXNlcnZlLWNocm9tZVxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXNlcnZlLWNocm9tZScpXG5leHBvcnQgY2xhc3MgQ2VtU2VydmVDaHJvbWUgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAvLyBTdGF0aWMgdGVtcGxhdGVzIGZvciBET00gY2xvbmluZyAobG9ncywgZXZlbnRzLCBkZWJ1ZyBpbmZvKVxuICBzdGF0aWMgI2RlbW9JbmZvVGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPGgzPkRlbW8gSW5mb3JtYXRpb248L2gzPlxuICAgICAgPGRsIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdCBwZi1tLWhvcml6b250YWwgcGYtbS1jb21wYWN0XCI+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5UYWcgTmFtZTwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwidGFnTmFtZVwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIiBkYXRhLWZpZWxkLWdyb3VwPVwiZGVzY3JpcHRpb25cIj5cbiAgICAgICAgICA8ZHQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+RGVzY3JpcHRpb248L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cImRlc2NyaXB0aW9uXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiIGRhdGEtZmllbGQtZ3JvdXA9XCJmaWxlcGF0aFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5GaWxlIFBhdGg8L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cImZpbGVwYXRoXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5DYW5vbmljYWwgVVJMPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJjYW5vbmljYWxVUkxcIj48L2RkPlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkxvY2FsIFJvdXRlPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICA8L2RsPmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgc3RhdGljICNkZW1vR3JvdXBUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiIGRhdGEtZmllbGQ9XCJ0YWdOYW1lXCI+PC9kdD5cbiAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIj5cbiAgICAgICAgICA8c3BhbiBkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIj48L3NwYW4+PGJyPlxuICAgICAgICAgIDxzbWFsbCBkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIj5GaWxlOiA8c3BhbiBkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIj48L3NwYW4+PC9zbWFsbD5cbiAgICAgICAgICA8c21hbGw+Q2Fub25pY2FsOiA8c3BhbiBkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCI+PC9zcGFuPjwvc21hbGw+PGJyPlxuICAgICAgICAgIDxzbWFsbD5Mb2NhbDogPHNwYW4gZGF0YS1maWVsZD1cImxvY2FsUm91dGVcIj48L3NwYW4+PC9zbWFsbD5cbiAgICAgICAgPC9kZD5cbiAgICAgIDwvZGl2PmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgc3RhdGljICNkZW1vTGlzdFRlbXBsYXRlID0gKCgpID0+IHtcbiAgICBjb25zdCB0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGVtcGxhdGUnKTtcbiAgICB0LmlubmVySFRNTCA9IGBcbiAgICAgIDxjZW0tcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uIGlkPVwiZGVidWctZGVtb3Mtc2VjdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHRvZ2dsZS10ZXh0PVwiU2hvdyBEZW1vcyBJbmZvXCI+XG4gICAgICAgIDxkbCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3QgcGYtbS1ob3Jpem9udGFsIHBmLW0tY29tcGFjdFwiIGRhdGEtY29udGFpbmVyPVwiZ3JvdXBzXCI+PC9kbD5cbiAgICAgIDwvY2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbj5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjbG9nRW50cnlUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8ZGl2IGNsYXNzPVwibG9nLWVudHJ5XCIgZGF0YS1maWVsZD1cImNvbnRhaW5lclwiPlxuICAgICAgICA8Y2VtLXBmLXY2LWxhYmVsIGNvbXBhY3QgZGF0YS1maWVsZD1cImxhYmVsXCI+PC9jZW0tcGYtdjYtbGFiZWw+XG4gICAgICAgIDx0aW1lIGNsYXNzPVwibG9nLXRpbWVcIiBkYXRhLWZpZWxkPVwidGltZVwiPjwvdGltZT5cbiAgICAgICAgPHNwYW4gY2xhc3M9XCJsb2ctbWVzc2FnZVwiIGRhdGEtZmllbGQ9XCJtZXNzYWdlXCI+PC9zcGFuPlxuICAgICAgPC9kaXY+YDtcbiAgICByZXR1cm4gdDtcbiAgfSkoKTtcblxuICBzdGF0aWMgI2V2ZW50RW50cnlUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8YnV0dG9uIGNsYXNzPVwiZXZlbnQtbGlzdC1pdGVtXCIgZGF0YS1maWVsZD1cImNvbnRhaW5lclwiPlxuICAgICAgICA8Y2VtLXBmLXY2LWxhYmVsIGNvbXBhY3QgZGF0YS1maWVsZD1cImxhYmVsXCI+PC9jZW0tcGYtdjYtbGFiZWw+XG4gICAgICAgIDx0aW1lIGNsYXNzPVwiZXZlbnQtdGltZVwiIGRhdGEtZmllbGQ9XCJ0aW1lXCI+PC90aW1lPlxuICAgICAgICA8c3BhbiBjbGFzcz1cImV2ZW50LWVsZW1lbnRcIiBkYXRhLWZpZWxkPVwiZWxlbWVudFwiPjwvc3Bhbj5cbiAgICAgIDwvYnV0dG9uPmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAncHJpbWFyeS10YWctbmFtZScgfSlcbiAgYWNjZXNzb3IgcHJpbWFyeVRhZ05hbWUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdkZW1vLXRpdGxlJyB9KVxuICBhY2Nlc3NvciBkZW1vVGl0bGUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdwYWNrYWdlLW5hbWUnIH0pXG4gIGFjY2Vzc29yIHBhY2thZ2VOYW1lID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnY2Fub25pY2FsLXVybCcgfSlcbiAgYWNjZXNzb3IgY2Fub25pY2FsVVJMID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnc291cmNlLXVybCcgfSlcbiAgYWNjZXNzb3Igc291cmNlVVJMID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3Iga25vYnMgPSAnJztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBkcmF3ZXI6ICdleHBhbmRlZCcgfCAnY29sbGFwc2VkJyB8ICcnID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnZHJhd2VyLWhlaWdodCcgfSlcbiAgYWNjZXNzb3IgZHJhd2VySGVpZ2h0ID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAndGFicy1zZWxlY3RlZCcgfSlcbiAgYWNjZXNzb3IgdGFic1NlbGVjdGVkID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3Igc2lkZWJhcjogJ2V4cGFuZGVkJyB8ICdjb2xsYXBzZWQnIHwgJycgPSAnJztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCBhdHRyaWJ1dGU6ICdoYXMtZGVzY3JpcHRpb24nIH0pXG4gIGFjY2Vzc29yIGhhc0Rlc2NyaXB0aW9uID0gZmFsc2U7XG5cbiAgIyQoaWQ6IHN0cmluZykge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKGlkKTtcbiAgfVxuXG4gICMkJChzZWxlY3Rvcjogc3RyaW5nKSB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvckFsbChzZWxlY3RvcikgPz8gW107XG4gIH1cblxuICAjbG9nQ29udGFpbmVyOiBIVE1MRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjZHJhd2VyT3BlbiA9IGZhbHNlO1xuICAjaW5pdGlhbExvZ3NGZXRjaGVkID0gZmFsc2U7XG4gICNpc0luaXRpYWxMb2FkID0gdHJ1ZTtcbiAgI2RlYnVnRGF0YTogRGVidWdEYXRhIHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gRWxlbWVudCBldmVudCB0cmFja2luZ1xuICAjZWxlbWVudEV2ZW50TWFwOiBNYXA8c3RyaW5nLCBFdmVudEluZm8+IHwgbnVsbCA9IG51bGw7XG4gICNtYW5pZmVzdDogTWFuaWZlc3QgfCBudWxsID0gbnVsbDtcbiAgI2NhcHR1cmVkRXZlbnRzOiBFdmVudFJlY29yZFtdID0gW107XG4gICNtYXhDYXB0dXJlZEV2ZW50cyA9IDEwMDA7XG4gICNldmVudExpc3Q6IEhUTUxFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNldmVudERldGFpbEhlYWRlcjogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50RGV0YWlsQm9keTogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI3NlbGVjdGVkRXZlbnRJZDogc3RyaW5nIHwgbnVsbCA9IG51bGw7XG4gICNldmVudHNGaWx0ZXJWYWx1ZSA9ICcnO1xuICAjZXZlbnRzRmlsdGVyRGVib3VuY2VUaW1lcjogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50VHlwZUZpbHRlcnMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgI2VsZW1lbnRGaWx0ZXJzID0gbmV3IFNldDxzdHJpbmc+KCk7XG4gICNkaXNjb3ZlcmVkRWxlbWVudHMgPSBuZXcgU2V0PHN0cmluZz4oKTtcblxuICAvLyBFdmVudCBsaXN0ZW5lciByZWZlcmVuY2VzIGZvciBjbGVhbnVwXG4gICNoYW5kbGVMb2dzRXZlbnQ6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZUV4cGFuZDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG4gICNoYW5kbGVUcmVlQ29sbGFwc2U6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZVNlbGVjdDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gVGltZW91dCBJRHMgZm9yIGNsZWFudXBcbiAgI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjY29weURlYnVnRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dDogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcblxuICAvLyBMb2cgZmlsdGVyIHN0YXRlXG4gICNsb2dzRmlsdGVyVmFsdWUgPSAnJztcbiAgI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyOiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjbG9nTGV2ZWxGaWx0ZXJzID0gbmV3IFNldChbJ2luZm8nLCAnd2FybicsICdlcnJvcicsICdzdWNjZXNzJywgJ2RlYnVnJywgJ3RyYWNlJ10pO1xuICAjbG9nTGV2ZWxEcm9wZG93bjogRWxlbWVudCB8IG51bGwgPSBudWxsO1xuXG4gIC8vIFdhdGNoIGZvciBkeW5hbWljYWxseSBhZGRlZCBlbGVtZW50c1xuICAvKiBjOCBpZ25vcmUgc3RhcnQgLSBNdXRhdGlvbk9ic2VydmVyIGNhbGxiYWNrIHRlc3RlZCB2aWEgaW50ZWdyYXRpb24gKi9cbiAgI29ic2VydmVyID0gbmV3IE11dGF0aW9uT2JzZXJ2ZXIoKG11dGF0aW9ucykgPT4ge1xuICAgIGxldCBuZWVkc1VwZGF0ZSA9IGZhbHNlO1xuXG4gICAgZm9yIChjb25zdCBtdXRhdGlvbiBvZiBtdXRhdGlvbnMpIHtcbiAgICAgIGZvciAoY29uc3Qgbm9kZSBvZiBtdXRhdGlvbi5hZGRlZE5vZGVzKSB7XG4gICAgICAgIGlmIChub2RlIGluc3RhbmNlb2YgSFRNTEVsZW1lbnQpIHtcbiAgICAgICAgICBjb25zdCB0YWdOYW1lID0gbm9kZS50YWdOYW1lLnRvTG93ZXJDYXNlKCk7XG4gICAgICAgICAgaWYgKHRoaXMuI2VsZW1lbnRFdmVudE1hcD8uaGFzKHRhZ05hbWUpICYmICFub2RlLmRhdGFzZXQuY2VtRXZlbnRzQXR0YWNoZWQpIHtcbiAgICAgICAgICAgIGNvbnN0IGV2ZW50SW5mbyA9IHRoaXMuI2VsZW1lbnRFdmVudE1hcC5nZXQodGFnTmFtZSkhO1xuICAgICAgICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgZXZlbnRJbmZvLmV2ZW50TmFtZXMpIHtcbiAgICAgICAgICAgICAgbm9kZS5hZGRFdmVudExpc3RlbmVyKGV2ZW50TmFtZSwgdGhpcy4jaGFuZGxlRWxlbWVudEV2ZW50LCB7IGNhcHR1cmU6IHRydWUgfSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgICBub2RlLmRhdGFzZXQuY2VtRXZlbnRzQXR0YWNoZWQgPSAndHJ1ZSc7XG4gICAgICAgICAgICB0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuYWRkKHRhZ05hbWUpO1xuICAgICAgICAgICAgbmVlZHNVcGRhdGUgPSB0cnVlO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIGlmIChuZWVkc1VwZGF0ZSkge1xuICAgICAgdGhpcy4jdXBkYXRlRXZlbnRGaWx0ZXJzKCk7XG4gICAgfVxuICB9KTtcbiAgLyogYzggaWdub3JlIHN0b3AgKi9cblxuICAjd3NDbGllbnQ6IGFueTtcblxuICAjaW5pdFdzQ2xpZW50KCkge1xuICAgIHRoaXMuI3dzQ2xpZW50ID0gbmV3IENFTVJlbG9hZENsaWVudCh7XG4gICAgaml0dGVyTWF4OiAxMDAwLFxuICAgIG92ZXJsYXlUaHJlc2hvbGQ6IDE1LFxuICAgIGJhZGdlRmFkZURlbGF5OiAyMDAwLFxuICAgIC8qIGM4IGlnbm9yZSBzdGFydCAtIFdlYlNvY2tldCBjYWxsYmFja3MgdGVzdGVkIHZpYSBpbnRlZ3JhdGlvbiAqL1xuICAgIGNhbGxiYWNrczoge1xuICAgICAgb25PcGVuOiAoKSA9PiB7XG4gICAgICAgIHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpPy5jbG9zZSgpO1xuICAgICAgfSxcbiAgICAgIG9uRXJyb3I6IChlcnJvckRhdGE6IHsgdGl0bGU/OiBzdHJpbmc7IG1lc3NhZ2U/OiBzdHJpbmc7IGZpbGU/OiBzdHJpbmcgfSkgPT4ge1xuICAgICAgICBpZiAoZXJyb3JEYXRhPy50aXRsZSAmJiBlcnJvckRhdGE/Lm1lc3NhZ2UpIHtcbiAgICAgICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlXSBTZXJ2ZXIgZXJyb3I6JywgZXJyb3JEYXRhKTtcbiAgICAgICAgICAodGhpcy4jJCgnZXJyb3Itb3ZlcmxheScpIGFzIGFueSk/LnNob3coZXJyb3JEYXRhLnRpdGxlLCBlcnJvckRhdGEubWVzc2FnZSwgZXJyb3JEYXRhLmZpbGUpO1xuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmVdIFdlYlNvY2tldCBlcnJvcjonLCBlcnJvckRhdGEpO1xuICAgICAgICB9XG4gICAgICB9LFxuICAgICAgb25SZWNvbm5lY3Rpbmc6ICh7IGF0dGVtcHQsIGRlbGF5IH06IHsgYXR0ZW1wdDogbnVtYmVyOyBkZWxheTogbnVtYmVyIH0pID0+IHtcbiAgICAgICAgaWYgKGF0dGVtcHQgPj0gMTUpIHtcbiAgICAgICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLW1vZGFsJykgYXMgYW55KT8uc2hvd01vZGFsKCk7XG4gICAgICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1jb250ZW50JykgYXMgYW55KT8udXBkYXRlUmV0cnlJbmZvKGF0dGVtcHQsIGRlbGF5KTtcbiAgICAgICAgfVxuICAgICAgfSxcbiAgICAgIG9uUmVsb2FkOiAoKSA9PiB7XG4gICAgICAgIGNvbnN0IGVycm9yT3ZlcmxheSA9IHRoaXMuIyQoJ2Vycm9yLW92ZXJsYXknKTtcbiAgICAgICAgaWYgKGVycm9yT3ZlcmxheT8uaGFzQXR0cmlidXRlKCdvcGVuJykpIHtcbiAgICAgICAgICAoZXJyb3JPdmVybGF5IGFzIGFueSkuaGlkZSgpO1xuICAgICAgICB9XG4gICAgICAgIHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKTtcbiAgICAgIH0sXG4gICAgICBvblNodXRkb3duOiAoKSA9PiB7XG4gICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tbW9kYWwnKSBhcyBhbnkpPy5zaG93TW9kYWwoKTtcbiAgICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1jb250ZW50JykgYXMgYW55KT8udXBkYXRlUmV0cnlJbmZvKDMwLCAzMDAwMCk7XG4gICAgICB9LFxuICAgICAgb25Mb2dzOiAobG9nczogQXJyYXk8eyB0eXBlOiBzdHJpbmc7IGRhdGU6IHN0cmluZzsgbWVzc2FnZTogc3RyaW5nIH0+KSA9PiB7XG4gICAgICAgIHdpbmRvdy5kaXNwYXRjaEV2ZW50KG5ldyBDZW1Mb2dzRXZlbnQobG9ncykpO1xuICAgICAgfVxuICAgIH1cbiAgICAvKiBjOCBpZ25vcmUgc3RvcCAqL1xuICAgIH0pO1xuICB9XG5cbiAgZ2V0IGRlbW8oKTogRWxlbWVudCB8IG51bGwge1xuICAgIHJldHVybiB0aGlzLnF1ZXJ5U2VsZWN0b3IoJ2NlbS1zZXJ2ZS1kZW1vJyk7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8bGluayByZWw9XCJzdHlsZXNoZWV0XCIgaHJlZj1cIi9fX2NlbS9jZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0LmNzc1wiPlxuICAgICAgPGxpbmsgcmVsPVwic3R5bGVzaGVldFwiIGhyZWY9XCIvX19jZW0vcGYtbGlnaHRkb20uY3NzXCI+XG5cbiAgICAgIDxjZW0tcGYtdjYtcGFnZSA/c2lkZWJhci1jb2xsYXBzZWQ9JHt0aGlzLnNpZGViYXIgPT09ICdjb2xsYXBzZWQnfT5cbiAgICAgICAgPGNlbS1wZi12Ni1za2lwLXRvLWNvbnRlbnQgc2xvdD1cInNraXAtdG8tY29udGVudFwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaHJlZj1cIiNtYWluLWNvbnRlbnRcIj5cbiAgICAgICAgICBTa2lwIHRvIGNvbnRlbnRcbiAgICAgICAgPC9jZW0tcGYtdjYtc2tpcC10by1jb250ZW50PlxuXG4gICAgICAgIDxjZW0tcGYtdjYtbWFzdGhlYWQgc2xvdD1cIm1hc3RoZWFkXCI+XG4gICAgICAgICAgPGEgY2xhc3M9XCJtYXN0aGVhZC1sb2dvXCJcbiAgICAgICAgICAgICBocmVmPVwiL1wiXG4gICAgICAgICAgICAgc2xvdD1cImxvZ29cIj5cbiAgICAgICAgICAgIDxpbWcgYWx0PVwiQ0VNIERldiBTZXJ2ZXJcIlxuICAgICAgICAgICAgICAgICBzcmM9XCIvX19jZW0vbG9nby5zdmdcIj5cbiAgICAgICAgICAgICR7dGhpcy5wYWNrYWdlTmFtZSA/IGh0bWxgPGgxPiR7dGhpcy5wYWNrYWdlTmFtZX08L2gxPmAgOiBub3RoaW5nfVxuICAgICAgICAgIDwvYT5cbiAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXIgc2xvdD1cInRvb2xiYXJcIj5cbiAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1ncm91cCB2YXJpYW50PVwiYWN0aW9uLWdyb3VwXCI+XG4gICAgICAgICAgICAgICR7dGhpcy4jcmVuZGVyU291cmNlQnV0dG9uKCl9XG4gICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtYnV0dG9uIGlkPVwiZGVidWctaW5mb1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwicGxhaW5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkRlYnVnIGluZm9cIj5cbiAgICAgICAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiXG4gICAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjE2XCJcbiAgICAgICAgICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCAxNiAxNlwiXG4gICAgICAgICAgICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgICAgICAgICAgICByb2xlPVwicHJlc2VudGF0aW9uXCI+XG4gICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNOCAxNUE3IDcgMCAxIDEgOCAxYTcgNyAwIDAgMSAwIDE0em0wIDFBOCA4IDAgMSAwIDggMGE4IDggMCAwIDAgMCAxNnpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJtOC45MyA2LjU4OC0yLjI5LjI4Ny0uMDgyLjM4LjQ1LjA4M2MuMjk0LjA3LjM1Mi4xNzYuMjg4LjQ2OWwtLjczOCAzLjQ2OGMtLjE5NC44OTcuMTA1IDEuMzE5LjgwOCAxLjMxOS41NDUgMCAxLjE3OC0uMjUyIDEuNDY1LS41OThsLjA4OC0uNDE2Yy0uMi4xNzYtLjQ5Mi4yNDYtLjY4Ni4yNDYtLjI3NSAwLS4zNzUtLjE5My0uMzA0LS41MzNMOC45MyA2LjU4OHpNOSA0LjVhMSAxIDAgMSAxLTIgMCAxIDEgMCAwIDEgMiAwelwiLz5cbiAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvZ2dsZS1ncm91cCBjbGFzcz1cImNvbG9yLXNjaGVtZS10b2dnbGVcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkNvbG9yIHNjaGVtZVwiPlxuICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cImxpZ2h0XCI+XG4gICAgICAgICAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIiBhcmlhLWxhYmVsPVwiTGlnaHQgbW9kZVwiPlxuICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNOCAxMWEzIDMgMCAxIDEgMC02IDMgMyAwIDAgMSAwIDZ6bTAgMWE0IDQgMCAxIDAgMC04IDQgNCAwIDAgMCAwIDh6TTggMGEuNS41IDAgMCAxIC41LjV2MmEuNS41IDAgMCAxLTEgMHYtMkEuNS41IDAgMCAxIDggMHptMCAxM2EuNS41IDAgMCAxIC41LjV2MmEuNS41IDAgMCAxLTEgMHYtMkEuNS41IDAgMCAxIDggMTN6bTgtNWEuNS41IDAgMCAxLS41LjVoLTJhLjUuNSAwIDAgMSAwLTFoMmEuNS41IDAgMCAxIC41LjV6TTMgOGEuNS41IDAgMCAxLS41LjVoLTJhLjUuNSAwIDAgMSAwLTFoMkEuNS41IDAgMCAxIDMgOHptMTAuNjU3LTUuNjU3YS41LjUgMCAwIDEgMCAuNzA3bC0xLjQxNCAxLjQxNWEuNS41IDAgMSAxLS43MDctLjcwOGwxLjQxNC0xLjQxNGEuNS41IDAgMCAxIC43MDcgMHptLTkuMTkzIDkuMTkzYS41LjUgMCAwIDEgMCAuNzA3TDMuMDUgMTMuNjU3YS41LjUgMCAwIDEtLjcwNy0uNzA3bDEuNDE0LTEuNDE0YS41LjUgMCAwIDEgLjcwNyAwem05LjE5MyAyLjEyMWEuNS41IDAgMCAxLS43MDcgMGwtMS40MTQtMS40MTRhLjUuNSAwIDAgMSAuNzA3LS43MDdsMS40MTQgMS40MTRhLjUuNSAwIDAgMSAwIC43MDd6TTQuNDY0IDQuNDY1YS41LjUgMCAwIDEtLjcwNyAwTDIuMzQzIDMuMDVhLjUuNSAwIDEgMSAuNzA3LS43MDdsMS40MTQgMS40MTRhLjUuNSAwIDAgMSAwIC43MDh6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtPlxuICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cImRhcmtcIj5cbiAgICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiIGFyaWEtbGFiZWw9XCJEYXJrIG1vZGVcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTYgLjI3OGEuNzY4Ljc2OCAwIDAgMSAuMDguODU4IDcuMjA4IDcuMjA4IDAgMCAwLS44NzggMy40NmMwIDQuMDIxIDMuMjc4IDcuMjc3IDcuMzE4IDcuMjc3LjUyNyAwIDEuMDQtLjA1NSAxLjUzMy0uMTZhLjc4Ny43ODcgMCAwIDEgLjgxLjMxNi43MzMuNzMzIDAgMCAxLS4wMzEuODkzQTguMzQ5IDguMzQ5IDAgMCAxIDguMzQ0IDE2QzMuNzM0IDE2IDAgMTIuMjg2IDAgNy43MSAwIDQuMjY2IDIuMTE0IDEuMzEyIDUuMTI0LjA2QS43NTIuNzUyIDAgMCAxIDYgLjI3OHpNNC44NTggMS4zMTFBNy4yNjkgNy4yNjkgMCAwIDAgMS4wMjUgNy43MWMwIDQuMDIgMy4yNzkgNy4yNzYgNy4zMTkgNy4yNzZhNy4zMTYgNy4zMTYgMCAwIDAgNS4yMDUtMi4xNjJjLS4zMzcuMDQyLS42OC4wNjMtMS4wMjkuMDYzLTQuNjEgMC04LjM0My0zLjcxNC04LjM0My04LjI5IDAtMS4xNjcuMjQyLTIuMjc4LjY4MS0zLjI4NnpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0+XG4gICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtIHZhbHVlPVwic3lzdGVtXCI+XG4gICAgICAgICAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIiBhcmlhLWxhYmVsPVwiU3lzdGVtIHByZWZlcmVuY2VcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTAgMS41QTEuNSAxLjUgMCAwIDEgMS41IDBoMTNBMS41IDEuNSAwIDAgMSAxNiAxLjV2OGExLjUgMS41IDAgMCAxLTEuNSAxLjVoLTEzQTEuNSAxLjUgMCAwIDEgMCA5LjV2LTh6TTEuNSAxYS41LjUgMCAwIDAtLjUuNXY4YS41LjUgMCAwIDAgLjUuNWgxM2EuNS41IDAgMCAwIC41LS41di04YS41LjUgMCAwIDAtLjUtLjVoLTEzelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTIuNSAxMmgxMWEuNS41IDAgMCAxIDAgMWgtMTFhLjUuNSAwIDAgMSAwLTF6bTAgMmgxMWEuNS41IDAgMCAxIDAgMWgtMTFhLjUuNSAwIDAgMSAwLTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtPlxuICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvZ2dsZS1ncm91cD5cbiAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyPlxuICAgICAgICA8L2NlbS1wZi12Ni1tYXN0aGVhZD5cblxuICAgICAgICA8Y2VtLXBmLXY2LXBhZ2Utc2lkZWJhciBzbG90PVwic2lkZWJhclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgP2V4cGFuZGVkPSR7dGhpcy5zaWRlYmFyID09PSAnZXhwYW5kZWQnfVxuICAgICAgICAgICAgICAgICAgICAgICAgICAgID9jb2xsYXBzZWQ9JHt0aGlzLnNpZGViYXIgIT09ICdleHBhbmRlZCd9PlxuICAgICAgICAgIDxzbG90IG5hbWU9XCJuYXZpZ2F0aW9uXCI+PC9zbG90PlxuICAgICAgICA8L2NlbS1wZi12Ni1wYWdlLXNpZGViYXI+XG5cbiAgICAgICAgPGNlbS1wZi12Ni1wYWdlLW1haW4gc2xvdD1cIm1haW5cIiBpZD1cIm1haW4tY29udGVudFwiPlxuICAgICAgICAgIDxzbG90Pjwvc2xvdD5cbiAgICAgICAgICA8Zm9vdGVyIGNsYXNzPVwicGYtbS1zdGlja3ktYm90dG9tXCI+XG4gICAgICAgICAgICA8ZGl2IGNsYXNzPVwiZm9vdGVyLWRlc2NyaXB0aW9uJHt0aGlzLmhhc0Rlc2NyaXB0aW9uID8gJycgOiAnIGVtcHR5J31cIj5cbiAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cImRlc2NyaXB0aW9uXCI+PC9zbG90PlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICA8Y2VtLWRyYXdlciA/b3Blbj0ke3RoaXMuZHJhd2VyID09PSAnZXhwYW5kZWQnfVxuICAgICAgICAgICAgICAgICAgICAgICAgZHJhd2VyLWhlaWdodD1cIiR7dGhpcy5kcmF3ZXJIZWlnaHQgfHwgJzQwMCd9XCI+XG4gICAgICAgICAgICAgIDxjZW0tcGYtdjYtdGFicyBzZWxlY3RlZD1cIiR7dGhpcy50YWJzU2VsZWN0ZWQgfHwgJzAnfVwiPlxuICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdGFiIHRpdGxlPVwiS25vYnNcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgaWQ9XCJrbm9icy1jb250YWluZXJcIj5cbiAgICAgICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cImtub2JzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHAgY2xhc3M9XCJrbm9icy1lbXB0eVwiPk5vIGtub2JzIGF2YWlsYWJsZSBmb3IgdGhpcyBlbGVtZW50LjwvcD5cbiAgICAgICAgICAgICAgICAgICAgPC9zbG90PlxuICAgICAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdGFiPlxuICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdGFiIHRpdGxlPVwiTWFuaWZlc3QgQnJvd3NlclwiPlxuICAgICAgICAgICAgICAgICAgPGNlbS1tYW5pZmVzdC1icm93c2VyPlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtdHJlZVwiIHNsb3Q9XCJtYW5pZmVzdC10cmVlXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtbmFtZVwiIHNsb3Q9XCJtYW5pZmVzdC1uYW1lXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtZGV0YWlsc1wiIHNsb3Q9XCJtYW5pZmVzdC1kZXRhaWxzXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgPC9jZW0tbWFuaWZlc3QtYnJvd3Nlcj5cbiAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10YWI+XG4gICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10YWIgdGl0bGU9XCJTZXJ2ZXIgTG9nc1wiPlxuICAgICAgICAgICAgICAgICAgPGRpdiBjbGFzcz1cImxvZ3Mtd3JhcHBlclwiPlxuICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXIgc3RpY2t5PlxuICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRleHQtaW5wdXQtZ3JvdXAgaWQ9XCJsb2dzLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHBsYWNlaG9sZGVyPVwiRmlsdGVyIGxvZ3MuLi5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpY29uPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDUxMiA1MTJcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNTA1IDQ0Mi43TDQwNS4zIDM0M2MtNC41LTQuNS0xMC42LTctMTctN0gzNzJjMjcuNi0zNS4zIDQ0LTc5LjcgNDQtMTI4QzQxNiA5My4xIDMyMi45IDAgMjA4IDBTMCA5My4xIDAgMjA4czkzLjEgMjA4IDIwOCAyMDhjNDguMyAwIDkyLjctMTYuNCAxMjgtNDR2MTYuM2MwIDYuNCAyLjUgMTIuNSA3IDE3bDk5LjcgOTkuN2M5LjQgOS40IDI0LjYgOS40IDMzLjkgMGwyOC4zLTI4LjNjOS40LTkuNCA5LjQtMjQuNi4xLTM0ek0yMDggMzM2Yy03MC43IDAtMTI4LTU3LjItMTI4LTEyOCAwLTcwLjcgNTcuMi0xMjggMTI4LTEyOCA3MC43IDAgMTI4IDU3LjIgMTI4IDEyOCAwIDcwLjctNTcuMiAxMjgtMTI4IDEyOHpcIj48L3BhdGg+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRleHQtaW5wdXQtZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni1kcm9wZG93biBpZD1cImxvZy1sZXZlbC1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgbG9nIGxldmVsc1wiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzcGFuIHNsb3Q9XCJ0b2dnbGUtdGV4dFwiPkxvZyBMZXZlbHM8L3NwYW4+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJlcnJvclwiIGNoZWNrZWQ+RXJyb3JzPC9jZW0tcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwid2FyblwiIGNoZWNrZWQ+V2FybmluZ3M8L2NlbS1wZi12Ni1tZW51LWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJzdWNjZXNzXCIgY2hlY2tlZD5TdWNjZXNzPC9jZW0tcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwiaW5mb1wiIGNoZWNrZWQ+SW5mbzwvY2VtLXBmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LW1lbnUtaXRlbSB2YXJpYW50PVwiY2hlY2tib3hcIiB2YWx1ZT1cImRlYnVnXCIgY2hlY2tlZD5EZWJ1ZzwvY2VtLXBmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LW1lbnUtaXRlbSB2YXJpYW50PVwiY2hlY2tib3hcIiB2YWx1ZT1cInRyYWNlXCIgY2hlY2tlZD5UcmFjZTwvY2VtLXBmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtZHJvcGRvd24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItZ3JvdXAgdmFyaWFudD1cImFjdGlvbi1ncm91cFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtYnV0dG9uIGlkPVwiY29weS1sb2dzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0xMyAwSDZhMiAyIDAgMCAwLTIgMiAyIDIgMCAwIDAtMiAydjEwYTIgMiAwIDAgMCAyIDJoN2EyIDIgMCAwIDAgMi0yIDIgMiAwIDAgMCAyLTJWMmEyIDIgMCAwIDAtMi0yem0wIDEzVjRhMiAyIDAgMCAwLTItMkg1YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjEwYTEgMSAwIDAgMS0xIDF6TTMgMTNWNGExIDEgMCAwIDEgMS0xaDdhMSAxIDAgMCAxIDEgMXY5YTEgMSAwIDAgMS0xIDFINGExIDEgMCAwIDEtMS0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDb3B5IExvZ3NcbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtYnV0dG9uPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXI+XG4gICAgICAgICAgICAgICAgICAgIDxkaXYgaWQ9XCJsb2ctY29udGFpbmVyXCI+PC9kaXY+XG4gICAgICAgICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10YWI+XG4gICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10YWIgdGl0bGU9XCJFdmVudHNcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJldmVudHMtd3JhcHBlclwiPlxuICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXIgc3RpY2t5PlxuICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRleHQtaW5wdXQtZ3JvdXAgaWQ9XCJldmVudHMtZmlsdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgcGxhY2Vob2xkZXI9XCJGaWx0ZXIgZXZlbnRzLi4uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaWNvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCIgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNMTEuNzQyIDEwLjM0NGE2LjUgNi41IDAgMSAwLTEuMzk3IDEuMzk4aC0uMDAxYy4wMy4wNC4wNjIuMDc4LjA5OC4xMTVsMy44NSAzLjg1YTEgMSAwIDAgMCAxLjQxNS0xLjQxNGwtMy44NS0zLjg1YTEuMDA3IDEuMDA3IDAgMCAwLS4xMTUtLjF6TTEyIDYuNWE1LjUgNS41IDAgMSAxLTExIDAgNS41IDUuNSAwIDAgMSAxMSAwelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LWRyb3Bkb3duIGlkPVwiZXZlbnQtdHlwZS1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgZXZlbnQgdHlwZXNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5FdmVudCBUeXBlczwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtZHJvcGRvd24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni1kcm9wZG93biBpZD1cImVsZW1lbnQtZmlsdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGxhYmVsPVwiRmlsdGVyIGVsZW1lbnRzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHNwYW4gc2xvdD1cInRvZ2dsZS10ZXh0XCI+RWxlbWVudHM8L3NwYW4+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LWRyb3Bkb3duPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWdyb3VwIHZhcmlhbnQ9XCJhY3Rpb24tZ3JvdXBcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBpZD1cImNsZWFyLWV2ZW50c1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInRlcnRpYXJ5XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBzaXplPVwic21hbGxcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCIgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNS41IDUuNUEuNS41IDAgMCAxIDYgNnY2YS41LjUgMCAwIDEtMSAwVjZhLjUuNSAwIDAgMSAuNS0uNXptMi41IDBhLjUuNSAwIDAgMSAuNS41djZhLjUuNSAwIDAgMS0xIDBWNmEuNS41IDAgMCAxIC41LS41em0zIC41YS41LjUgMCAwIDAtMSAwdjZhLjUuNSAwIDAgMCAxIDBWNnpcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBmaWxsLXJ1bGU9XCJldmVub2RkXCIgZD1cIk0xNC41IDNhMSAxIDAgMCAxLTEgMUgxM3Y5YTIgMiAwIDAgMS0yIDJINWEyIDIgMCAwIDEtMi0yVjRoLS41YTEgMSAwIDAgMS0xLTFWMmExIDEgMCAwIDEgMS0xSDZhMSAxIDAgMCAxIDEtMWgyYTEgMSAwIDAgMSAxIDFoMy41YTEgMSAwIDAgMSAxIDF2MXpNNC4xMTggNCA0IDQuMDU5VjEzYTEgMSAwIDAgMCAxIDFoNmExIDEgMCAwIDAgMS0xVjQuMDU5TDExLjg4MiA0SDQuMTE4ek0yLjUgM1YyaDExdjFoLTExelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDbGVhciBFdmVudHNcbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtYnV0dG9uPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtYnV0dG9uIGlkPVwiY29weS1ldmVudHNcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJ0ZXJ0aWFyeVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgc2l6ZT1cInNtYWxsXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHN2ZyBzbG90PVwiaWNvblwiIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTEzIDBINmEyIDIgMCAwIDAtMiAyIDIgMiAwIDAgMC0yIDJ2MTBhMiAyIDAgMCAwIDIgMmg3YTIgMiAwIDAgMCAyLTIgMiAyIDAgMCAwIDItMlYyYTIgMiAwIDAgMC0yLTJ6bTAgMTNWNGEyIDIgMCAwIDAtMi0ySDVhMSAxIDAgMCAxIDEtMWg3YTEgMSAwIDAgMSAxIDF2MTBhMSAxIDAgMCAxLTEgMXpNMyAxM1Y0YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjlhMSAxIDAgMCAxLTEgMUg0YTEgMSAwIDAgMS0xLTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIENvcHkgRXZlbnRzXG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyPlxuICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LWRyYXdlciBpZD1cImV2ZW50LWRyYXdlclwiIGV4cGFuZGVkPlxuICAgICAgICAgICAgICAgICAgICAgIDxkaXYgaWQ9XCJldmVudC1saXN0XCI+PC9kaXY+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWRldGFpbC1oZWFkZXJcIiBzbG90PVwicGFuZWwtaGVhZGVyXCI+PC9kaXY+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWRldGFpbC1ib2R5XCIgc2xvdD1cInBhbmVsLWJvZHlcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtZHJhd2VyPlxuICAgICAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdGFiPlxuICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdGFiIHRpdGxlPVwiSGVhbHRoXCI+XG4gICAgICAgICAgICAgICAgICA8Y2VtLWhlYWx0aC1wYW5lbCBjb21wb25lbnQ9JHtpZkRlZmluZWQodGhpcy5wcmltYXJ5VGFnTmFtZSl9PlxuICAgICAgICAgICAgICAgICAgPC9jZW0taGVhbHRoLXBhbmVsPlxuICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRhYj5cbiAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdGFicz5cbiAgICAgICAgICAgIDwvY2VtLWRyYXdlcj5cbiAgICAgICAgICA8L2Zvb3Rlcj5cbiAgICAgICAgPC9jZW0tcGYtdjYtcGFnZS1tYWluPlxuICAgICAgPC9jZW0tcGYtdjYtcGFnZT5cblxuICAgICAgPGNlbS1wZi12Ni1tb2RhbCBpZD1cImRlYnVnLW1vZGFsXCIgdmFyaWFudD1cImxhcmdlXCI+XG4gICAgICAgIDxoMiBzbG90PVwiaGVhZGVyXCI+RGVidWcgSW5mb3JtYXRpb248L2gyPlxuICAgICAgICA8ZGwgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0IHBmLW0taG9yaXpvbnRhbCBwZi1tLWNvbXBhY3RcIj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5TZXJ2ZXIgVmVyc2lvbjwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctdmVyc2lvblwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlNlcnZlciBPUzwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctb3NcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5XYXRjaCBEaXJlY3Rvcnk8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLXdhdGNoLWRpclwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPk1hbmlmZXN0IFNpemU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLW1hbmlmZXN0LXNpemVcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5EZW1vcyBGb3VuZDwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctZGVtby1jb3VudFwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlRhZyBOYW1lPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCI+JHt0aGlzLnByaW1hcnlUYWdOYW1lIHx8ICctJ308L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkRlbW8gVGl0bGU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIj4ke3RoaXMuZGVtb1RpdGxlIHx8ICctJ308L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkJyb3dzZXI8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLWJyb3dzZXJcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5Vc2VyIEFnZW50PC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy11YVwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICA8L2RsPlxuICAgICAgICA8ZGl2IGlkPVwiZGVtby11cmxzLWNvbnRhaW5lclwiPjwvZGl2PlxuICAgICAgICA8Y2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbiBpZD1cImRlYnVnLWltcG9ydG1hcC1kZXRhaWxzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB0b2dnbGUtdGV4dD1cIlNob3cgSW1wb3J0IE1hcFwiPlxuICAgICAgICAgIDxwcmUgaWQ9XCJkZWJ1Zy1pbXBvcnRtYXBcIj4tPC9wcmU+XG4gICAgICAgIDwvY2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbj5cbiAgICAgICAgPGRpdiBzbG90PVwiZm9vdGVyXCIgY2xhc3M9XCJidXR0b24tY29udGFpbmVyXCI+XG4gICAgICAgICAgPGNlbS1wZi12Ni1idXR0b24gY2xhc3M9XCJkZWJ1Zy1jb3B5XCIgdmFyaWFudD1cInByaW1hcnlcIj5cbiAgICAgICAgICAgIENvcHkgRGVidWcgSW5mb1xuICAgICAgICAgIDwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBjbGFzcz1cImRlYnVnLWNsb3NlXCIgdmFyaWFudD1cInNlY29uZGFyeVwiPlxuICAgICAgICAgICAgQ2xvc2VcbiAgICAgICAgICA8L2NlbS1wZi12Ni1idXR0b24+XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9jZW0tcGYtdjYtbW9kYWw+XG5cbiAgICAgIDwhLS0gUmVjb25uZWN0aW9uIG1vZGFsIC0tPlxuICAgICAgPGNlbS1wZi12Ni1tb2RhbCBpZD1cInJlY29ubmVjdGlvbi1tb2RhbFwiIHZhcmlhbnQ9XCJsYXJnZVwiPlxuICAgICAgICA8aDIgc2xvdD1cImhlYWRlclwiPkRldmVsb3BtZW50IFNlcnZlciBEaXNjb25uZWN0ZWQ8L2gyPlxuICAgICAgICA8Y2VtLXJlY29ubmVjdGlvbi1jb250ZW50IGlkPVwicmVjb25uZWN0aW9uLWNvbnRlbnRcIj48L2NlbS1yZWNvbm5lY3Rpb24tY29udGVudD5cbiAgICAgICAgPGNlbS1wZi12Ni1idXR0b24gaWQ9XCJyZWxvYWQtYnV0dG9uXCJcbiAgICAgICAgICAgICAgICAgICAgICBzbG90PVwiZm9vdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwicHJpbWFyeVwiPlJlbG9hZCBQYWdlPC9jZW0tcGYtdjYtYnV0dG9uPlxuICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBpZD1cInJldHJ5LWJ1dHRvblwiXG4gICAgICAgICAgICAgICAgICAgICAgc2xvdD1cImZvb3RlclwiXG4gICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInNlY29uZGFyeVwiPlJldHJ5IE5vdzwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgIDwvY2VtLXBmLXY2LW1vZGFsPlxuXG4gICAgICA8IS0tIFRyYW5zZm9ybSBlcnJvciBvdmVybGF5IC0tPlxuICAgICAgPGNlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheSBpZD1cImVycm9yLW92ZXJsYXlcIj5cbiAgICAgIDwvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5PlxuICAgIGA7XG4gIH1cblxuICAjcmVuZGVyU291cmNlQnV0dG9uKCkge1xuICAgIGlmICghdGhpcy5zb3VyY2VVUkwpIHJldHVybiBub3RoaW5nO1xuXG4gICAgbGV0IGxhYmVsID0gJ1ZlcnNpb24gQ29udHJvbCc7XG4gICAgbGV0IHBhdGggPSAnTTUuODU0IDQuODU0YS41LjUgMCAxIDAtLjcwOC0uNzA4bC0zLjUgMy41YS41LjUgMCAwIDAgMCAuNzA4bDMuNSAzLjVhLjUuNSAwIDAgMCAuNzA4LS43MDhMMi43MDcgOGwzLjE0Ny0zLjE0NnptNC4yOTIgMGEuNS41IDAgMCAxIC43MDgtLjcwOGwzLjUgMy41YS41LjUgMCAwIDEgMCAuNzA4bC0zLjUgMy41YS41LjUgMCAwIDEtLjcwOC0uNzA4TDEzLjI5MyA4bC0zLjE0Ny0zLjE0NnonO1xuXG4gICAgaWYgKHRoaXMuc291cmNlVVJMLmluY2x1ZGVzKCdnaXRodWIuY29tJykpIHtcbiAgICAgIGxhYmVsID0gJ0dpdEh1Yi5jb20nO1xuICAgICAgcGF0aCA9ICdNOCAwQzMuNTggMCAwIDMuNTggMCA4YzAgMy41NCAyLjI5IDYuNTMgNS40NyA3LjU5LjQuMDcuNTUtLjE3LjU1LS4zOCAwLS4xOS0uMDEtLjgyLS4wMS0xLjQ5LTIuMDEuMzctMi41My0uNDktMi42OS0uOTQtLjA5LS4yMy0uNDgtLjk0LS44Mi0xLjEzLS4yOC0uMTUtLjY4LS41Mi0uMDEtLjUzLjYzLS4wMSAxLjA4LjU4IDEuMjMuODIuNzIgMS4yMSAxLjg3Ljg3IDIuMzMuNjYuMDctLjUyLjI4LS44Ny41MS0xLjA3LTEuNzgtLjItMy42NC0uODktMy42NC0zLjk1IDAtLjg3LjMxLTEuNTkuODItMi4xNS0uMDgtLjItLjM2LTEuMDIuMDgtMi4xMiAwIDAgLjY3LS4yMSAyLjIuODIuNjQtLjE4IDEuMzItLjI3IDItLjI3LjY4IDAgMS4zNi4wOSAyIC4yNyAxLjUzLTEuMDQgMi4yLS44MiAyLjItLjgyLjQ0IDEuMS4xNiAxLjkyLjA4IDIuMTIuNTEuNTYuODIgMS4yNy44MiAyLjE1IDAgMy4wNy0xLjg3IDMuNzUtMy42NSAzLjk1LjI5LjI1LjU0LjczLjU0IDEuNDggMCAxLjA3LS4wMSAxLjkzLS4wMSAyLjIgMCAuMjEuMTUuNDYuNTUuMzhBOC4wMTMgOC4wMTMgMCAwMDE2IDhjMC00LjQyLTMuNTgtOC04LTh6JztcbiAgICB9IGVsc2UgaWYgKHRoaXMuc291cmNlVVJMLmluY2x1ZGVzKCdnaXRsYWIuY29tJykpIHtcbiAgICAgIGxhYmVsID0gJ0dpdExhYic7XG4gICAgICBwYXRoID0gJ20xNS43MzQgNi4xLS4wMjItLjA1OEwxMy41MzQuMzU4YS41NjguNTY4IDAgMCAwLS41NjMtLjM1Ni41ODMuNTgzIDAgMCAwLS4zMjguMTIyLjU4Mi41ODIgMCAwIDAtLjE5My4yOTRsLTEuNDcgNC40OTlINS4wMjVsLTEuNDctNC41QS41NzIuNTcyIDAgMCAwIDMuMzYwLjE3NGEuNTcyLjU3MiAwIDAgMC0uMzI4LS4xNzIuNTgyLjU4MiAwIDAgMC0uNTYzLjM1N0wuMjkgNi4wNGwtLjAyMi4wNTdBNC4wNDQgNC4wNDQgMCAwIDAgMS42MSAxMC43N2wuMDA3LjAwNi4wMi4wMTQgMy4zMTggMi40ODUgMS42NCAxLjI0MiAxIC43NTVhLjY3My42NzMgMCAwIDAgLjgxNCAwbDEtLjc1NSAxLjY0LTEuMjQyIDMuMzM4LTIuNS4wMDktLjAwN2E0LjA1IDQuMDUgMCAwIDAgMS4zNC00LjY2OFonO1xuICAgIH0gZWxzZSBpZiAodGhpcy5zb3VyY2VVUkwuaW5jbHVkZXMoJ2JpdGJ1Y2tldC5vcmcnKSkge1xuICAgICAgbGFiZWwgPSAnQml0YnVja2V0JztcbiAgICAgIHBhdGggPSAnTTAgMS41QTEuNSAxLjUgMCAwIDEgMS41IDBoMTNBMS41IDEuNSAwIDAgMSAxNiAxLjV2MTNhMS41IDEuNSAwIDAgMS0xLjUgMS41aC0xM0ExLjUgMS41IDAgMCAxIDAgMTQuNXYtMTN6TTIuNSAzYS41LjUgMCAwIDAtLjUuNXY5YS41LjUgMCAwIDAgLjUuNWgxMWEuNS41IDAgMCAwIC41LS41di05YS41LjUgMCAwIDAtLjUtLjVoLTExem01LjAzOCAxLjQzNWEuNS41IDAgMCAxIC45MjQgMGwxLjQyIDMuMzdIOC43OGwtLjI0My42MDgtLjI0My0uNjA4SDUuMDgybDEuNDItMy4zN3pNOCA5LjE0M2wtLjc0MyAxLjg1N0g0Ljc0M0w2LjA3NiA3LjYwOCA4IDkuMTQzeic7XG4gICAgfVxuXG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgPGNlbS1wZi12Ni1idXR0b24gaHJlZj1cIiR7dGhpcy5zb3VyY2VVUkx9XCJcbiAgICAgICAgICAgICAgICAgICAgICB0YXJnZXQ9XCJfYmxhbmtcIlxuICAgICAgICAgICAgICAgICAgICAgIHJlbD1cIm5vb3BlbmVyIG5vcmVmZXJyZXJcIlxuICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIlZpZXcgc291cmNlIGZpbGVcIj5cbiAgICAgICAgICA8c3ZnIGFyaWEtbGFiZWw9XCIke2xhYmVsfVwiXG4gICAgICAgICAgICAgICB3aWR0aD1cIjFyZW1cIlxuICAgICAgICAgICAgICAgaGVpZ2h0PVwiMXJlbVwiXG4gICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIj5cbiAgICAgICAgICAgIDxwYXRoIGQ9XCIke3BhdGh9XCIvPlxuICAgICAgICAgIDwvc3ZnPlxuICAgICAgICA8L2NlbS1wZi12Ni1idXR0b24+XG4gICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgYDtcbiAgfVxuXG4gIC8qKiBSZXNvbHZlcyB3aGVuIGNsaWVudC1vbmx5IG1vZHVsZXMgYXJlIGxvYWRlZC4gKi9cbiAgI21vZHVsZXNSZWFkeTogUHJvbWlzZTx2b2lkPiA9IHRoaXMuI2xvYWRDbGllbnRNb2R1bGVzKCk7XG5cbiAgI2xvYWRDbGllbnRNb2R1bGVzKCk6IFByb21pc2U8dm9pZD4ge1xuICAgIHJldHVybiBQcm9taXNlLmFsbChbXG4gICAgICAvLyBAdHMtaWdub3JlIC0tIHBsYWluIEpTIG1vZHVsZXMgc2VydmVkIGF0IHJ1bnRpbWUgYnkgR28gc2VydmVyXG4gICAgICBpbXBvcnQoJy9fX2NlbS93ZWJzb2NrZXQtY2xpZW50LmpzJyksXG4gICAgICAvLyBAdHMtaWdub3JlXG4gICAgICBpbXBvcnQoJy9fX2NlbS9zdGF0ZS1wZXJzaXN0ZW5jZS5qcycpLFxuICAgIF0pLnRoZW4oKFt3cywgc3BdKSA9PiB7XG4gICAgICBDRU1SZWxvYWRDbGllbnQgPSB3cy5DRU1SZWxvYWRDbGllbnQ7XG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlID0gc3AuU3RhdGVQZXJzaXN0ZW5jZTtcbiAgICAgIC8vIEB0cy1pZ25vcmVcbiAgICAgIGltcG9ydCgnL19fY2VtL2hlYWx0aC1iYWRnZXMuanMnKS5jYXRjaCgoZTogdW5rbm93bikgPT5cbiAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZV0gRmFpbGVkIHRvIGxvYWQgaGVhbHRoLWJhZGdlczonLCBlKSk7XG4gICAgfSkuY2F0Y2goKGU6IHVua25vd24pID0+IHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmVdIEZhaWxlZCB0byBsb2FkIGNsaWVudCBtb2R1bGVzOicsIGUpO1xuICAgICAgLy8gRGVncmFkZWQgZmFsbGJhY2tzIHNvIHRoZSBjb21wb25lbnQgc3RpbGwgcmVuZGVyc1xuICAgICAgQ0VNUmVsb2FkQ2xpZW50ID8/PSBjbGFzcyB7IGluaXQoKSB7fSByZXRyeSgpIHt9IGRlc3Ryb3koKSB7fSB9IGFzIGFueTtcbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UgPz89IHtcbiAgICAgICAgZ2V0U3RhdGU6ICgpID0+ICh7IGNvbG9yU2NoZW1lOiAnc3lzdGVtJyB9KSxcbiAgICAgICAgdXBkYXRlU3RhdGUoKSB7fSxcbiAgICAgICAgZ2V0VHJlZVN0YXRlOiAoKSA9PiAoeyBleHBhbmRlZDogW10sIHNlbGVjdGVkOiBudWxsIH0pLFxuICAgICAgICBzZXRUcmVlU3RhdGUoKSB7fSxcbiAgICAgICAgdXBkYXRlVHJlZVN0YXRlKCkge30sXG4gICAgICAgIG1pZ3JhdGVGcm9tTG9jYWxTdG9yYWdlKCkge30sXG4gICAgICB9O1xuICAgIH0pO1xuICB9XG5cbiAgYXN5bmMgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgLy8gTW9kdWxlcyBtdXN0IGxvYWQgYmVmb3JlIHN1cGVyIHNvIGZpcnN0VXBkYXRlZCgpIGNhbiB1c2UgU3RhdGVQZXJzaXN0ZW5jZVxuICAgIGF3YWl0IHRoaXMuI21vZHVsZXNSZWFkeTtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuXG4gICAgaWYgKHRoaXMuI3dzQ2xpZW50ID09IG51bGwpIHtcbiAgICAgIHRoaXMuI2luaXRXc0NsaWVudCgpO1xuICAgIH1cbiAgICB0aGlzLiNtaWdyYXRlRnJvbUxvY2FsU3RvcmFnZUlmTmVlZGVkKCk7XG4gIH1cblxuICBmaXJzdFVwZGF0ZWQoKSB7XG4gICAgLy8gU2V0IHVwIGRlYnVnIG92ZXJsYXlcbiAgICB0aGlzLiNzZXR1cERlYnVnT3ZlcmxheSgpO1xuXG4gICAgLy8gU2V0IHVwIGNvbG9yIHNjaGVtZSB0b2dnbGVcbiAgICB0aGlzLiNzZXR1cENvbG9yU2NoZW1lVG9nZ2xlKCk7XG5cbiAgICAvLyBTZXQgdXAgZm9vdGVyIGRyYXdlciBhbmQgdGFic1xuICAgIHRoaXMuI3NldHVwRm9vdGVyRHJhd2VyKCk7XG5cbiAgICAvLyBMaXN0ZW4gZm9yIHNlcnZlciBsb2cgbWVzc2FnZXMgZnJvbSBXZWJTb2NrZXRcbiAgICB0aGlzLiNzZXR1cExvZ0xpc3RlbmVyKCk7XG5cbiAgICAvLyBTZXQgdXAga25vYiBldmVudCBjb29yZGluYXRpb25cbiAgICB0aGlzLiNzZXR1cEtub2JDb29yZGluYXRpb24oKTtcblxuICAgIC8vIFNldCB1cCB0cmVlIHN0YXRlIHBlcnNpc3RlbmNlXG4gICAgdGhpcy4jc2V0dXBUcmVlU3RhdGVQZXJzaXN0ZW5jZSgpO1xuXG4gICAgLy8gU2V0IHVwIHNpZGViYXIgc3RhdGUgcGVyc2lzdGVuY2VcbiAgICB0aGlzLiNzZXR1cFNpZGViYXJTdGF0ZVBlcnNpc3RlbmNlKCk7XG5cbiAgICAvLyBTZXQgdXAgZWxlbWVudCBldmVudCBjYXB0dXJlXG4gICAgdGhpcy4jc2V0dXBFdmVudENhcHR1cmUoKS50aGVuKCgpID0+IHtcbiAgICAgIHRoaXMuI3NldHVwRXZlbnRMaXN0ZW5lcnMoKTtcbiAgICB9KTtcblxuICAgIC8vIFNldCB1cCByZWNvbm5lY3Rpb24gbW9kYWwgYnV0dG9uIGhhbmRsZXJzXG4gICAgLyogYzggaWdub3JlIHN0YXJ0IC0gd2luZG93LmxvY2F0aW9uLnJlbG9hZCB3b3VsZCByZWxvYWQgdGVzdCBwYWdlICovXG4gICAgdGhpcy4jJCgncmVsb2FkLWJ1dHRvbicpPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgIHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKTtcbiAgICB9KTtcbiAgICAvKiBjOCBpZ25vcmUgc3RvcCAqL1xuXG4gICAgdGhpcy4jJCgncmV0cnktYnV0dG9uJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LmNsb3NlKCk7XG4gICAgICB0aGlzLiN3c0NsaWVudC5yZXRyeSgpO1xuICAgIH0pO1xuXG4gICAgLy8gSW5pdGlhbGl6ZSBXZWJTb2NrZXQgY29ubmVjdGlvblxuICAgIHRoaXMuI3dzQ2xpZW50LmluaXQoKTtcbiAgfVxuXG4gIGFzeW5jICNmZXRjaERlYnVnSW5mbygpIHtcbiAgICAvLyBQb3B1bGF0ZSBicm93c2VyIGluZm8gaW1tZWRpYXRlbHlcbiAgICBjb25zdCBicm93c2VyRWwgPSB0aGlzLiMkKCdkZWJ1Zy1icm93c2VyJyk7XG4gICAgY29uc3QgdWFFbCA9IHRoaXMuIyQoJ2RlYnVnLXVhJyk7XG4gICAgaWYgKGJyb3dzZXJFbCkge1xuICAgICAgY29uc3QgYnJvd3NlciA9IHRoaXMuI2RldGVjdEJyb3dzZXIoKTtcbiAgICAgIGJyb3dzZXJFbC50ZXh0Q29udGVudCA9IGJyb3dzZXI7XG4gICAgfVxuICAgIGlmICh1YUVsKSB7XG4gICAgICB1YUVsLnRleHRDb250ZW50ID0gbmF2aWdhdG9yLnVzZXJBZ2VudDtcbiAgICB9XG5cbiAgICAvLyBGZXRjaCBzZXJ2ZXIgZGVidWcgaW5mb1xuICAgIGNvbnN0IGRhdGEgPSBhd2FpdCBmZXRjaCgnL19fY2VtL2RlYnVnJylcbiAgICAgIC50aGVuKHJlcyA9PiByZXMuanNvbigpKVxuICAgICAgLmNhdGNoKChlcnI6IEVycm9yKSA9PiB7XG4gICAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gZmV0Y2ggZGVidWcgaW5mbzonLCBlcnIpO1xuICAgICAgfSk7XG5cbiAgICBpZiAoIWRhdGEpIHJldHVybjtcbiAgICBjb25zdCB2ZXJzaW9uRWwgPSB0aGlzLiMkKCdkZWJ1Zy12ZXJzaW9uJyk7XG4gICAgY29uc3Qgb3NFbCA9IHRoaXMuIyQoJ2RlYnVnLW9zJyk7XG4gICAgY29uc3Qgd2F0Y2hEaXJFbCA9IHRoaXMuIyQoJ2RlYnVnLXdhdGNoLWRpcicpO1xuICAgIGNvbnN0IG1hbmlmZXN0U2l6ZUVsID0gdGhpcy4jJCgnZGVidWctbWFuaWZlc3Qtc2l6ZScpO1xuICAgIGNvbnN0IGRlbW9Db3VudEVsID0gdGhpcy4jJCgnZGVidWctZGVtby1jb3VudCcpO1xuICAgIGNvbnN0IGRlbW9VcmxzQ29udGFpbmVyID0gdGhpcy4jJCgnZGVtby11cmxzLWNvbnRhaW5lcicpO1xuICAgIGNvbnN0IGltcG9ydE1hcEVsID0gdGhpcy4jJCgnZGVidWctaW1wb3J0bWFwJyk7XG5cbiAgICBpZiAodmVyc2lvbkVsKSB2ZXJzaW9uRWwudGV4dENvbnRlbnQgPSBkYXRhLnZlcnNpb24gfHwgJy0nO1xuICAgIGlmIChvc0VsKSBvc0VsLnRleHRDb250ZW50ID0gZGF0YS5vcyB8fCAnLSc7XG4gICAgaWYgKHdhdGNoRGlyRWwpIHdhdGNoRGlyRWwudGV4dENvbnRlbnQgPSBkYXRhLndhdGNoRGlyIHx8ICctJztcbiAgICBpZiAobWFuaWZlc3RTaXplRWwpIG1hbmlmZXN0U2l6ZUVsLnRleHRDb250ZW50ID0gZGF0YS5tYW5pZmVzdFNpemUgfHwgJy0nO1xuICAgIGlmIChkZW1vQ291bnRFbCkgZGVtb0NvdW50RWwudGV4dENvbnRlbnQgPSBkYXRhLmRlbW9Db3VudCB8fCAnMCc7XG5cbiAgICBpZiAoZGVtb1VybHNDb250YWluZXIpIHtcbiAgICAgIHRoaXMuI3BvcHVsYXRlRGVtb1VybHMoZGVtb1VybHNDb250YWluZXIsIGRhdGEuZGVtb3MpO1xuICAgIH1cblxuICAgIGlmIChpbXBvcnRNYXBFbCAmJiBkYXRhLmltcG9ydE1hcCkge1xuICAgICAgY29uc3QgaW1wb3J0TWFwSlNPTiA9IEpTT04uc3RyaW5naWZ5KGRhdGEuaW1wb3J0TWFwLCBudWxsLCAyKTtcbiAgICAgIGltcG9ydE1hcEVsLnRleHRDb250ZW50ID0gaW1wb3J0TWFwSlNPTjtcbiAgICAgIGRhdGEuaW1wb3J0TWFwSlNPTiA9IGltcG9ydE1hcEpTT047XG4gICAgfSBlbHNlIGlmIChpbXBvcnRNYXBFbCkge1xuICAgICAgaW1wb3J0TWFwRWwudGV4dENvbnRlbnQgPSAnTm8gaW1wb3J0IG1hcCBnZW5lcmF0ZWQnO1xuICAgIH1cblxuICAgIHRoaXMuI2RlYnVnRGF0YSA9IGRhdGE7XG4gIH1cblxuICAjcG9wdWxhdGVEZW1vVXJscyhjb250YWluZXI6IEhUTUxFbGVtZW50LCBkZW1vczogRGVidWdEYXRhWydkZW1vcyddKSB7XG4gICAgaWYgKCFkZW1vcz8ubGVuZ3RoKSB7XG4gICAgICBjb250YWluZXIudGV4dENvbnRlbnQgPSAnTm8gZGVtb3MgZm91bmQgaW4gbWFuaWZlc3QnO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IGN1cnJlbnRUYWdOYW1lID0gdGhpcy5wcmltYXJ5VGFnTmFtZSB8fCAnJztcbiAgICBjb25zdCBpc09uRGVtb1BhZ2UgPSAhIWN1cnJlbnRUYWdOYW1lO1xuXG4gICAgaWYgKGlzT25EZW1vUGFnZSkge1xuICAgICAgY29uc3QgY3VycmVudERlbW8gPSBkZW1vcy5maW5kKGRlbW8gPT4gZGVtby50YWdOYW1lID09PSBjdXJyZW50VGFnTmFtZSk7XG4gICAgICBpZiAoIWN1cnJlbnREZW1vKSB7XG4gICAgICAgIGNvbnRhaW5lci50ZXh0Q29udGVudCA9ICdDdXJyZW50IGRlbW8gbm90IGZvdW5kIGluIG1hbmlmZXN0JztcbiAgICAgICAgcmV0dXJuO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBmcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vSW5mb1RlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGFnTmFtZVwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLnRhZ05hbWU7XG4gICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNhbm9uaWNhbFVSTFwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmNhbm9uaWNhbFVSTDtcbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibG9jYWxSb3V0ZVwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmxvY2FsUm91dGU7XG5cbiAgICAgIGNvbnN0IGRlc2NyaXB0aW9uR3JvdXAgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImRlc2NyaXB0aW9uXCJdJyk7XG4gICAgICBpZiAoY3VycmVudERlbW8uZGVzY3JpcHRpb24pIHtcbiAgICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJkZXNjcmlwdGlvblwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmRlc2NyaXB0aW9uO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgZGVzY3JpcHRpb25Hcm91cD8ucmVtb3ZlKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGZpbGVwYXRoR3JvdXAgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImZpbGVwYXRoXCJdJyk7XG4gICAgICBpZiAoY3VycmVudERlbW8uZmlsZXBhdGgpIHtcbiAgICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJmaWxlcGF0aFwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmZpbGVwYXRoO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgZmlsZXBhdGhHcm91cD8ucmVtb3ZlKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnRhaW5lci5yZXBsYWNlQ2hpbGRyZW4oZnJhZ21lbnQpO1xuICAgIH0gZWxzZSB7XG4gICAgICBjb25zdCBsaXN0RnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZGVtb0xpc3RUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICBjb25zdCBncm91cHNDb250YWluZXIgPSBsaXN0RnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtY29udGFpbmVyPVwiZ3JvdXBzXCJdJykhO1xuXG4gICAgICBmb3IgKGNvbnN0IGRlbW8gb2YgZGVtb3MpIHtcbiAgICAgICAgY29uc3QgZ3JvdXBGcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vR3JvdXBUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0YWdOYW1lXCJdJykhLnRleHRDb250ZW50ID0gZGVtby50YWdOYW1lO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIl0nKSEudGV4dENvbnRlbnQgPVxuICAgICAgICAgIGRlbW8uZGVzY3JpcHRpb24gfHwgJyhubyBkZXNjcmlwdGlvbiknO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5jYW5vbmljYWxVUkw7XG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5sb2NhbFJvdXRlO1xuXG4gICAgICAgIGNvbnN0IGZpbGVwYXRoR3JvdXAgPSBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIl0nKTtcbiAgICAgICAgaWYgKGRlbW8uZmlsZXBhdGgpIHtcbiAgICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIl0nKSEudGV4dENvbnRlbnQgPSBkZW1vLmZpbGVwYXRoO1xuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIGZpbGVwYXRoR3JvdXA/LnJlbW92ZSgpO1xuICAgICAgICB9XG5cbiAgICAgICAgZ3JvdXBzQ29udGFpbmVyLmFwcGVuZENoaWxkKGdyb3VwRnJhZ21lbnQpO1xuICAgICAgfVxuXG4gICAgICBjb250YWluZXIucmVwbGFjZUNoaWxkcmVuKGxpc3RGcmFnbWVudCk7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwTG9nTGlzdGVuZXIoKSB7XG4gICAgdGhpcy4jbG9nQ29udGFpbmVyID0gdGhpcy4jJCgnbG9nLWNvbnRhaW5lcicpO1xuXG4gICAgY29uc3QgbG9nc0ZpbHRlciA9IHRoaXMuIyQoJ2xvZ3MtZmlsdGVyJykgYXMgSFRNTElucHV0RWxlbWVudCB8IG51bGw7XG4gICAgaWYgKGxvZ3NGaWx0ZXIpIHtcbiAgICAgIGxvZ3NGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignaW5wdXQnLCAoKSA9PiB7XG4gICAgICAgIGNvbnN0IHsgdmFsdWUgPSAnJyB9ID0gbG9nc0ZpbHRlcjtcbiAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyISk7XG4gICAgICAgIHRoaXMuI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgdGhpcy4jZmlsdGVyTG9ncyh2YWx1ZSk7XG4gICAgICAgIH0sIDMwMCk7XG4gICAgICB9KTtcbiAgICB9XG5cbiAgICB0aGlzLiNsb2dMZXZlbERyb3Bkb3duID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcjbG9nLWxldmVsLWZpbHRlcicpID8/IG51bGw7XG4gICAgaWYgKHRoaXMuI2xvZ0xldmVsRHJvcGRvd24pIHtcbiAgICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICAgIHRoaXMuI2xvYWRMb2dGaWx0ZXJTdGF0ZSgpO1xuICAgICAgfSk7XG4gICAgICB0aGlzLiNsb2dMZXZlbERyb3Bkb3duLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZUxvZ0ZpbHRlckNoYW5nZSBhcyBFdmVudExpc3RlbmVyKTtcbiAgICB9XG5cbiAgICB0aGlzLiMkKCdjb3B5LWxvZ3MnKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB0aGlzLiNjb3B5TG9ncygpO1xuICAgIH0pO1xuXG4gICAgdGhpcy4jaGFuZGxlTG9nc0V2ZW50ID0gKChldmVudDogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IGxvZ3MgPSAoZXZlbnQgYXMgQ2VtTG9nc0V2ZW50KS5sb2dzO1xuICAgICAgaWYgKGxvZ3MpIHtcbiAgICAgICAgdGhpcy4jcmVuZGVyTG9ncyhsb2dzKTtcbiAgICAgIH1cbiAgICB9KTtcbiAgICB3aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcignY2VtOmxvZ3MnLCB0aGlzLiNoYW5kbGVMb2dzRXZlbnQpO1xuICB9XG5cbiAgI2ZpbHRlckxvZ3MocXVlcnk6IHN0cmluZykge1xuICAgIHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgZm9yIChjb25zdCBlbnRyeSBvZiB0aGlzLiNsb2dDb250YWluZXIuY2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IHRleHQgPSBlbnRyeS50ZXh0Q29udGVudD8udG9Mb3dlckNhc2UoKSA/PyAnJztcbiAgICAgIGNvbnN0IHRleHRNYXRjaCA9ICF0aGlzLiNsb2dzRmlsdGVyVmFsdWUgfHwgdGV4dC5pbmNsdWRlcyh0aGlzLiNsb2dzRmlsdGVyVmFsdWUpO1xuXG4gICAgICBjb25zdCBsb2dUeXBlID0gdGhpcy4jZ2V0TG9nVHlwZUZyb21FbnRyeShlbnRyeSk7XG4gICAgICBjb25zdCBsZXZlbE1hdGNoID0gdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmhhcyhsb2dUeXBlKTtcblxuICAgICAgKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5oaWRkZW4gPSAhKHRleHRNYXRjaCAmJiBsZXZlbE1hdGNoKTtcbiAgICB9XG4gIH1cblxuICAjZ2V0TG9nVHlwZUZyb21FbnRyeShlbnRyeTogRWxlbWVudCk6IHN0cmluZyB7XG4gICAgZm9yIChjb25zdCBjbHMgb2YgZW50cnkuY2xhc3NMaXN0KSB7XG4gICAgICBpZiAoWydpbmZvJywgJ3dhcm5pbmcnLCAnZXJyb3InLCAnc3VjY2VzcycsICdkZWJ1ZycsICd0cmFjZSddLmluY2x1ZGVzKGNscykpIHtcbiAgICAgICAgcmV0dXJuIGNscyA9PT0gJ3dhcm5pbmcnID8gJ3dhcm4nIDogY2xzO1xuICAgICAgfVxuICAgIH1cbiAgICByZXR1cm4gJ2luZm8nO1xuICB9XG5cbiAgI2xvYWRMb2dGaWx0ZXJTdGF0ZSgpIHtcbiAgICB0cnkge1xuICAgICAgY29uc3Qgc2F2ZWQgPSBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWxvZy1maWx0ZXJzJyk7XG4gICAgICBpZiAoc2F2ZWQpIHtcbiAgICAgICAgdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzID0gbmV3IFNldChKU09OLnBhcnNlKHNhdmVkKSk7XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgY29uc29sZS5kZWJ1ZygnW2NlbS1zZXJ2ZS1jaHJvbWVdIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSwgdXNpbmcgZGVmYXVsdCBsb2cgZmlsdGVycycpO1xuICAgIH1cbiAgICB0aGlzLiNzeW5jQ2hlY2tib3hTdGF0ZXMoKTtcbiAgfVxuXG4gICNzeW5jQ2hlY2tib3hTdGF0ZXMoKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dMZXZlbERyb3Bkb3duKSByZXR1cm47XG4gICAgY29uc3QgbWVudUl0ZW1zID0gdGhpcy4jbG9nTGV2ZWxEcm9wZG93bi5xdWVyeVNlbGVjdG9yQWxsKCdjZW0tcGYtdjYtbWVudS1pdGVtJyk7XG4gICAgbWVudUl0ZW1zLmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBjb25zdCB2YWx1ZSA9IChpdGVtIGFzIGFueSkudmFsdWU7XG4gICAgICAoaXRlbSBhcyBhbnkpLmNoZWNrZWQgPSB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuaGFzKHZhbHVlKTtcbiAgICB9KTtcbiAgfVxuXG4gICNzYXZlTG9nRmlsdGVyU3RhdGUoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdjZW0tc2VydmUtbG9nLWZpbHRlcnMnLFxuICAgICAgICBKU09OLnN0cmluZ2lmeShbLi4udGhpcy4jbG9nTGV2ZWxGaWx0ZXJzXSkpO1xuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIC8vIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSAocHJpdmF0ZSBtb2RlKSwgc2lsZW50bHkgY29udGludWVcbiAgICB9XG4gIH1cblxuICAjaGFuZGxlTG9nRmlsdGVyQ2hhbmdlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHsgdmFsdWUsIGNoZWNrZWQgfSA9IGV2ZW50IGFzIEV2ZW50ICYgeyB2YWx1ZTogc3RyaW5nOyBjaGVja2VkOiBib29sZWFuIH07XG5cbiAgICBpZiAoY2hlY2tlZCkge1xuICAgICAgdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2xvZ0xldmVsRmlsdGVycy5kZWxldGUodmFsdWUpO1xuICAgIH1cblxuICAgIHRoaXMuI3NhdmVMb2dGaWx0ZXJTdGF0ZSgpO1xuICAgIHRoaXMuI2ZpbHRlckxvZ3ModGhpcy4jbG9nc0ZpbHRlclZhbHVlKTtcbiAgfTtcblxuICBhc3luYyAjY29weUxvZ3MoKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGNvbnN0IGxvZ3MgPSBBcnJheS5mcm9tKHRoaXMuI2xvZ0NvbnRhaW5lci5jaGlsZHJlbilcbiAgICAgIC5maWx0ZXIoZW50cnkgPT4gIShlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuKVxuICAgICAgLm1hcChlbnRyeSA9PiB7XG4gICAgICAgIGNvbnN0IHR5cGUgPSBlbnRyeS5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxhYmVsXCJdJyk/LnRleHRDb250ZW50Py50cmltKCkgfHwgJ0lORk8nO1xuICAgICAgICBjb25zdCB0aW1lID0gZW50cnkucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0aW1lXCJdJyk/LnRleHRDb250ZW50Py50cmltKCkgfHwgJyc7XG4gICAgICAgIGNvbnN0IG1lc3NhZ2UgPSBlbnRyeS5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cIm1lc3NhZ2VcIl0nKT8udGV4dENvbnRlbnQ/LnRyaW0oKSB8fCAnJztcbiAgICAgICAgcmV0dXJuIGBbJHt0eXBlfV0gJHt0aW1lfSAke21lc3NhZ2V9YDtcbiAgICAgIH0pLmpvaW4oJ1xcbicpO1xuXG4gICAgaWYgKCFsb2dzKSByZXR1cm47XG5cbiAgICB0cnkge1xuICAgICAgYXdhaXQgbmF2aWdhdG9yLmNsaXBib2FyZC53cml0ZVRleHQobG9ncyk7XG4gICAgICBjb25zdCBidG4gPSB0aGlzLiMkKCdjb3B5LWxvZ3MnKTtcbiAgICAgIGlmIChidG4pIHtcbiAgICAgICAgY29uc3QgdGV4dE5vZGUgPSBBcnJheS5mcm9tKGJ0bi5jaGlsZE5vZGVzKS5maW5kKFxuICAgICAgICAgIG4gPT4gbi5ub2RlVHlwZSA9PT0gTm9kZS5URVhUX05PREUgJiYgKG4udGV4dENvbnRlbnQ/LnRyaW0oKS5sZW5ndGggPz8gMCkgPiAwXG4gICAgICAgICk7XG4gICAgICAgIGlmICh0ZXh0Tm9kZSkge1xuICAgICAgICAgIGNvbnN0IG9yaWdpbmFsID0gdGV4dE5vZGUudGV4dENvbnRlbnQ7XG4gICAgICAgICAgdGV4dE5vZGUudGV4dENvbnRlbnQgPSAnQ29waWVkISc7XG5cbiAgICAgICAgICBpZiAodGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgdGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICAgIGlmICh0aGlzLmlzQ29ubmVjdGVkICYmIHRleHROb2RlLnBhcmVudE5vZGUpIHtcbiAgICAgICAgICAgICAgdGV4dE5vZGUudGV4dENvbnRlbnQgPSBvcmlnaW5hbDtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICAgICAgICB9LCAyMDAwKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gY2F0Y2ggKGVycikge1xuICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjb3B5IGxvZ3M6JywgZXJyKTtcbiAgICB9XG4gIH1cblxuICAjc2V0dXBEZWJ1Z092ZXJsYXkoKSB7XG4gICAgY29uc3QgZGVidWdCdXR0b24gPSB0aGlzLiMkKCdkZWJ1Zy1pbmZvJyk7XG4gICAgY29uc3QgZGVidWdNb2RhbCA9IHRoaXMuIyQoJ2RlYnVnLW1vZGFsJyk7XG4gICAgY29uc3QgZGVidWdDbG9zZSA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignLmRlYnVnLWNsb3NlJyk7XG4gICAgY29uc3QgZGVidWdDb3B5ID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuZGVidWctY29weScpO1xuXG4gICAgaWYgKGRlYnVnQnV0dG9uICYmIGRlYnVnTW9kYWwpIHtcbiAgICAgIGRlYnVnQnV0dG9uLmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgICB0aGlzLiNmZXRjaERlYnVnSW5mbygpO1xuICAgICAgICAoZGVidWdNb2RhbCBhcyBhbnkpLnNob3dNb2RhbCgpO1xuICAgICAgfSk7XG5cbiAgICAgIGRlYnVnQ2xvc2U/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4gKGRlYnVnTW9kYWwgYXMgYW55KS5jbG9zZSgpKTtcblxuICAgICAgZGVidWdDb3B5Py5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgICAgdGhpcy4jY29weURlYnVnSW5mbygpO1xuICAgICAgfSk7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwRm9vdGVyRHJhd2VyKCkge1xuICAgIGNvbnN0IGRyYXdlciA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignY2VtLWRyYXdlcicpO1xuICAgIGNvbnN0IHRhYnMgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ2NlbS1wZi12Ni10YWJzJyk7XG5cbiAgICBpZiAoIWRyYXdlciB8fCAhdGFicykgcmV0dXJuO1xuXG4gICAgdGhpcy4jZHJhd2VyT3BlbiA9IChkcmF3ZXIgYXMgYW55KS5vcGVuO1xuXG4gICAgZHJhd2VyLmFkZEV2ZW50TGlzdGVuZXIoJ2NoYW5nZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgdGhpcy4jZHJhd2VyT3BlbiA9IChlIGFzIGFueSkub3BlbjtcblxuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIGRyYXdlcjogeyBvcGVuOiAoZSBhcyBhbnkpLm9wZW4gfVxuICAgICAgfSk7XG5cbiAgICAgIGlmICgoZSBhcyBhbnkpLm9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTG9nc1RvQm90dG9tKCk7XG4gICAgICB9XG4gICAgfSk7XG5cbiAgICBkcmF3ZXIuYWRkRXZlbnRMaXN0ZW5lcigncmVzaXplJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICAoZHJhd2VyIGFzIGFueSkuc2V0QXR0cmlidXRlKCdkcmF3ZXItaGVpZ2h0JywgKGUgYXMgYW55KS5oZWlnaHQpO1xuXG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHtcbiAgICAgICAgZHJhd2VyOiB7IGhlaWdodDogKGUgYXMgYW55KS5oZWlnaHQgfVxuICAgICAgfSk7XG4gICAgfSk7XG5cbiAgICB0YWJzLmFkZEV2ZW50TGlzdGVuZXIoJ2NoYW5nZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIHRhYnM6IHsgc2VsZWN0ZWRJbmRleDogKGUgYXMgYW55KS5zZWxlY3RlZEluZGV4IH1cbiAgICAgIH0pO1xuXG4gICAgICBpZiAoKGUgYXMgYW55KS5zZWxlY3RlZEluZGV4ID09PSAyICYmIChkcmF3ZXIgYXMgYW55KS5vcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExvZ3NUb0JvdHRvbSgpO1xuICAgICAgfVxuXG4gICAgICBpZiAoKGUgYXMgYW55KS5zZWxlY3RlZEluZGV4ID09PSAzICYmIChkcmF3ZXIgYXMgYW55KS5vcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbEV2ZW50c1RvQm90dG9tKCk7XG4gICAgICB9XG4gICAgfSk7XG4gIH1cblxuICAjZGV0ZWN0QnJvd3NlcigpOiBzdHJpbmcge1xuICAgIGNvbnN0IHVhID0gbmF2aWdhdG9yLnVzZXJBZ2VudDtcbiAgICBpZiAodWEuaW5jbHVkZXMoJ0ZpcmVmb3gvJykpIHtcbiAgICAgIGNvbnN0IG1hdGNoID0gdWEubWF0Y2goL0ZpcmVmb3hcXC8oXFxkKykvKTtcbiAgICAgIHJldHVybiBtYXRjaCA/IGBGaXJlZm94ICR7bWF0Y2hbMV19YCA6ICdGaXJlZm94JztcbiAgICB9IGVsc2UgaWYgKHVhLmluY2x1ZGVzKCdFZGcvJykpIHtcbiAgICAgIGNvbnN0IG1hdGNoID0gdWEubWF0Y2goL0VkZ1xcLyhcXGQrKS8pO1xuICAgICAgcmV0dXJuIG1hdGNoID8gYEVkZ2UgJHttYXRjaFsxXX1gIDogJ0VkZ2UnO1xuICAgIH0gZWxzZSBpZiAodWEuaW5jbHVkZXMoJ0Nocm9tZS8nKSkge1xuICAgICAgY29uc3QgbWF0Y2ggPSB1YS5tYXRjaCgvQ2hyb21lXFwvKFxcZCspLyk7XG4gICAgICByZXR1cm4gbWF0Y2ggPyBgQ2hyb21lICR7bWF0Y2hbMV19YCA6ICdDaHJvbWUnO1xuICAgIH0gZWxzZSBpZiAodWEuaW5jbHVkZXMoJ1NhZmFyaS8nKSAmJiAhdWEuaW5jbHVkZXMoJ0Nocm9tZScpKSB7XG4gICAgICBjb25zdCBtYXRjaCA9IHVhLm1hdGNoKC9WZXJzaW9uXFwvKFxcZCspLyk7XG4gICAgICByZXR1cm4gbWF0Y2ggPyBgU2FmYXJpICR7bWF0Y2hbMV19YCA6ICdTYWZhcmknO1xuICAgIH1cbiAgICByZXR1cm4gJ1Vua25vd24nO1xuICB9XG5cbiAgYXN5bmMgI2NvcHlEZWJ1Z0luZm8oKSB7XG4gICAgY29uc3QgaW5mbyA9IEFycmF5LmZyb20odGhpcy4jJCQoJyNkZWJ1Zy1tb2RhbCBkbCBkdCcpKS5tYXAoZHQgPT4ge1xuICAgICAgY29uc3QgZGQgPSBkdC5uZXh0RWxlbWVudFNpYmxpbmc7XG4gICAgICBpZiAoZGQgJiYgZGQudGFnTmFtZSA9PT0gJ0REJykge1xuICAgICAgICByZXR1cm4gYCR7ZHQudGV4dENvbnRlbnR9OiAke2RkLnRleHRDb250ZW50fWA7XG4gICAgICB9XG4gICAgICByZXR1cm4gJyc7XG4gICAgfSkuam9pbignXFxuJyk7XG5cbiAgICBsZXQgaW1wb3J0TWFwU2VjdGlvbiA9ICcnO1xuICAgIGlmICh0aGlzLiNkZWJ1Z0RhdGE/LmltcG9ydE1hcEpTT04pIHtcbiAgICAgIGltcG9ydE1hcFNlY3Rpb24gPSBgXFxuJHsnPScucmVwZWF0KDQwKX1cXG5JbXBvcnQgTWFwOlxcbiR7Jz0nLnJlcGVhdCg0MCl9XFxuJHt0aGlzLiNkZWJ1Z0RhdGEuaW1wb3J0TWFwSlNPTn1cXG5gO1xuICAgIH1cblxuICAgIGNvbnN0IGRlYnVnVGV4dCA9IGBDRU0gU2VydmUgRGVidWcgSW5mb3JtYXRpb25cbiR7Jz0nLnJlcGVhdCg0MCl9XG4ke2luZm99JHtpbXBvcnRNYXBTZWN0aW9ufVxuJHsnPScucmVwZWF0KDQwKX1cbkdlbmVyYXRlZDogJHtuZXcgRGF0ZSgpLnRvSVNPU3RyaW5nKCl9YDtcblxuICAgIHRyeSB7XG4gICAgICBhd2FpdCBuYXZpZ2F0b3IuY2xpcGJvYXJkLndyaXRlVGV4dChkZWJ1Z1RleHQpO1xuICAgICAgY29uc3QgY29weUJ1dHRvbiA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignLmRlYnVnLWNvcHknKTtcbiAgICAgIGlmIChjb3B5QnV0dG9uKSB7XG4gICAgICAgIGNvbnN0IG9yaWdpbmFsVGV4dCA9IGNvcHlCdXR0b24udGV4dENvbnRlbnQ7XG4gICAgICAgIGNvcHlCdXR0b24udGV4dENvbnRlbnQgPSAnQ29waWVkISc7XG5cbiAgICAgICAgaWYgKHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgICB9XG5cbiAgICAgICAgdGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgaWYgKHRoaXMuaXNDb25uZWN0ZWQgJiYgY29weUJ1dHRvbi5wYXJlbnROb2RlKSB7XG4gICAgICAgICAgICBjb3B5QnV0dG9uLnRleHRDb250ZW50ID0gb3JpZ2luYWxUZXh0O1xuICAgICAgICAgIH1cbiAgICAgICAgICB0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgICAgICB9LCAyMDAwKTtcbiAgICAgIH1cbiAgICB9IGNhdGNoIChlcnIpIHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY29weSBkZWJ1ZyBpbmZvOicsIGVycik7XG4gICAgfVxuICB9XG5cbiAgI3JlbmRlckxvZ3MobG9nczogQXJyYXk8eyB0eXBlOiBzdHJpbmc7IGRhdGU6IHN0cmluZzsgbWVzc2FnZTogc3RyaW5nIH0+KSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGNvbnN0IGxvZ0VsZW1lbnRzID0gbG9ncy5tYXAobG9nID0+IHtcbiAgICAgIGNvbnN0IGZyYWdtZW50ID0gQ2VtU2VydmVDaHJvbWUuI2xvZ0VudHJ5VGVtcGxhdGUuY29udGVudC5jbG9uZU5vZGUodHJ1ZSkgYXMgRG9jdW1lbnRGcmFnbWVudDtcblxuICAgICAgY29uc3QgZGF0ZSA9IG5ldyBEYXRlKGxvZy5kYXRlKTtcbiAgICAgIGNvbnN0IHRpbWUgPSBkYXRlLnRvTG9jYWxlVGltZVN0cmluZygpO1xuXG4gICAgICBjb25zdCBjb250YWluZXIgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNvbnRhaW5lclwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgICAgY29udGFpbmVyLmNsYXNzTGlzdC5hZGQobG9nLnR5cGUpO1xuICAgICAgY29udGFpbmVyLnNldEF0dHJpYnV0ZSgnZGF0YS1sb2ctaWQnLCBsb2cuZGF0ZSk7XG5cbiAgICAgIGNvbnN0IHR5cGVMYWJlbCA9IHRoaXMuI2dldExvZ0JhZGdlKGxvZy50eXBlKTtcbiAgICAgIGNvbnN0IHNlYXJjaFRleHQgPSBgJHt0eXBlTGFiZWx9ICR7dGltZX0gJHtsb2cubWVzc2FnZX1gLnRvTG93ZXJDYXNlKCk7XG4gICAgICBjb25zdCB0ZXh0TWF0Y2ggPSAhdGhpcy4jbG9nc0ZpbHRlclZhbHVlIHx8IHNlYXJjaFRleHQuaW5jbHVkZXModGhpcy4jbG9nc0ZpbHRlclZhbHVlKTtcblxuICAgICAgY29uc3QgbG9nVHlwZUZvckZpbHRlciA9IGxvZy50eXBlID09PSAnd2FybmluZycgPyAnd2FybicgOiBsb2cudHlwZTtcbiAgICAgIGNvbnN0IGxldmVsTWF0Y2ggPSB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuaGFzKGxvZ1R5cGVGb3JGaWx0ZXIpO1xuXG4gICAgICBpZiAoISh0ZXh0TWF0Y2ggJiYgbGV2ZWxNYXRjaCkpIHtcbiAgICAgICAgY29udGFpbmVyLnNldEF0dHJpYnV0ZSgnaGlkZGVuJywgJycpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBsYWJlbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibGFiZWxcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICAgIGxhYmVsLnRleHRDb250ZW50ID0gdGhpcy4jZ2V0TG9nQmFkZ2UobG9nLnR5cGUpO1xuICAgICAgdGhpcy4jYXBwbHlMb2dMYWJlbEF0dHJzKGxhYmVsLCBsb2cudHlwZSk7XG5cbiAgICAgIGNvbnN0IHRpbWVFbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGltZVwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgICAgdGltZUVsLnNldEF0dHJpYnV0ZSgnZGF0ZXRpbWUnLCBsb2cuZGF0ZSk7XG4gICAgICB0aW1lRWwudGV4dENvbnRlbnQgPSB0aW1lO1xuXG4gICAgICAoZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJtZXNzYWdlXCJdJykgYXMgSFRNTEVsZW1lbnQpLnRleHRDb250ZW50ID0gbG9nLm1lc3NhZ2U7XG5cbiAgICAgIHJldHVybiBmcmFnbWVudDtcbiAgICB9KTtcblxuICAgIGlmICghdGhpcy4jaW5pdGlhbExvZ3NGZXRjaGVkKSB7XG4gICAgICB0aGlzLiNsb2dDb250YWluZXIucmVwbGFjZUNoaWxkcmVuKC4uLmxvZ0VsZW1lbnRzKTtcbiAgICAgIHRoaXMuI2luaXRpYWxMb2dzRmV0Y2hlZCA9IHRydWU7XG5cbiAgICAgIGlmICh0aGlzLiNkcmF3ZXJPcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExhdGVzdEludG9WaWV3KCk7XG4gICAgICB9XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2xvZ0NvbnRhaW5lci5hcHBlbmQoLi4ubG9nRWxlbWVudHMpO1xuXG4gICAgICBpZiAodGhpcy4jZHJhd2VyT3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNnZXRMb2dCYWRnZSh0eXBlOiBzdHJpbmcpOiBzdHJpbmcge1xuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnaW5mbyc6IHJldHVybiAnSW5mbyc7XG4gICAgICBjYXNlICd3YXJuaW5nJzogcmV0dXJuICdXYXJuJztcbiAgICAgIGNhc2UgJ2Vycm9yJzogcmV0dXJuICdFcnJvcic7XG4gICAgICBjYXNlICdzdWNjZXNzJzogcmV0dXJuICdPSyc7XG4gICAgICBjYXNlICdkZWJ1Zyc6IHJldHVybiAnRGVidWcnO1xuICAgICAgY2FzZSAndHJhY2UnOiByZXR1cm4gJ1RyYWNlJztcbiAgICAgIGRlZmF1bHQ6IHJldHVybiB0eXBlLnRvVXBwZXJDYXNlKCk7XG4gICAgfVxuICB9XG5cbiAgI2FwcGx5TG9nTGFiZWxBdHRycyhsYWJlbDogSFRNTEVsZW1lbnQsIHR5cGU6IHN0cmluZykge1xuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnaW5mbyc6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICdpbmZvJyk7XG4gICAgICBjYXNlICd3YXJuaW5nJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnc3RhdHVzJywgJ3dhcm5pbmcnKTtcbiAgICAgIGNhc2UgJ2Vycm9yJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnc3RhdHVzJywgJ2RhbmdlcicpO1xuICAgICAgY2FzZSAnc3VjY2Vzcyc6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICdzdWNjZXNzJyk7XG4gICAgICBjYXNlICdkZWJ1Zyc6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ2NvbG9yJywgJ3B1cnBsZScpO1xuICAgICAgY2FzZSAndHJhY2UnOlxuICAgICAgICByZXR1cm4gbGFiZWwuc2V0QXR0cmlidXRlKCdjb2xvcicsICdncmV5Jyk7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICBsYWJlbC5zZXRBdHRyaWJ1dGUoJ2NvbG9yJywgJ2dyZXknKTtcbiAgICB9XG4gIH1cblxuICAjc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICBjb25zdCBsYXN0TG9nID0gdGhpcy4jbG9nQ29udGFpbmVyIS5sYXN0RWxlbWVudENoaWxkO1xuICAgICAgaWYgKGxhc3RMb2cpIHtcbiAgICAgICAgbGFzdExvZy5zY3JvbGxJbnRvVmlldyh7IGJlaGF2aW9yOiAnYXV0bycsIGJsb2NrOiAnZW5kJyB9KTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNzY3JvbGxMb2dzVG9Cb3R0b20oKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGlmICh0aGlzLiNpc0luaXRpYWxMb2FkKSB7XG4gICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgIH0gZWxzZSB7XG4gICAgICBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKTtcbiAgICAgIH0sIDM1MCk7XG4gICAgfVxuICB9XG5cbiAgI21pZ3JhdGVGcm9tTG9jYWxTdG9yYWdlSWZOZWVkZWQoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IGhhc0xvY2FsU3RvcmFnZSA9XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtY29sb3Itc2NoZW1lJykgIT09IG51bGwgfHxcbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1kcmF3ZXItb3BlbicpICE9PSBudWxsIHx8XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtZHJhd2VyLWhlaWdodCcpICE9PSBudWxsIHx8XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtYWN0aXZlLXRhYicpICE9PSBudWxsO1xuXG4gICAgICBpZiAoaGFzTG9jYWxTdG9yYWdlKSB7XG4gICAgICAgIGNvbnN0IG1pZ3JhdGVkID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1taWdyYXRlZC10by1jb29raWVzJyk7XG4gICAgICAgIGlmICghbWlncmF0ZWQpIHtcbiAgICAgICAgICBTdGF0ZVBlcnNpc3RlbmNlLm1pZ3JhdGVGcm9tTG9jYWxTdG9yYWdlKCk7XG4gICAgICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oJ2NlbS1zZXJ2ZS1taWdyYXRlZC10by1jb29raWVzJywgJ3RydWUnKTtcbiAgICAgICAgICBzZXRUaW1lb3V0KCgpID0+IHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKSwgMTAwKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIC8vIGxvY2FsU3RvcmFnZSBub3QgYXZhaWxhYmxlLCBza2lwIG1pZ3JhdGlvblxuICAgIH1cbiAgfVxuXG4gICNzZXR1cENvbG9yU2NoZW1lVG9nZ2xlKCkge1xuICAgIGNvbnN0IHRvZ2dsZUdyb3VwID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuY29sb3Itc2NoZW1lLXRvZ2dsZScpO1xuICAgIGlmICghdG9nZ2xlR3JvdXApIHJldHVybjtcblxuICAgIGNvbnN0IHN0YXRlID0gU3RhdGVQZXJzaXN0ZW5jZS5nZXRTdGF0ZSgpO1xuXG4gICAgdGhpcy4jYXBwbHlDb2xvclNjaGVtZShzdGF0ZS5jb2xvclNjaGVtZSk7XG5cbiAgICBjb25zdCBpdGVtcyA9IHRvZ2dsZUdyb3VwLnF1ZXJ5U2VsZWN0b3JBbGwoJ2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbScpO1xuICAgIGl0ZW1zLmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBpZiAoKGl0ZW0gYXMgYW55KS52YWx1ZSA9PT0gc3RhdGUuY29sb3JTY2hlbWUpIHtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3NlbGVjdGVkJywgJycpO1xuICAgICAgfVxuICAgIH0pO1xuXG4gICAgdG9nZ2xlR3JvdXAuYWRkRXZlbnRMaXN0ZW5lcignY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1jaGFuZ2UnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IHNjaGVtZSA9IChlIGFzIGFueSkudmFsdWU7XG4gICAgICB0aGlzLiNhcHBseUNvbG9yU2NoZW1lKHNjaGVtZSk7XG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHsgY29sb3JTY2hlbWU6IHNjaGVtZSB9KTtcbiAgICB9KTtcbiAgfVxuXG4gICNhcHBseUNvbG9yU2NoZW1lKHNjaGVtZTogc3RyaW5nKSB7XG4gICAgY29uc3QgYm9keSA9IGRvY3VtZW50LmJvZHk7XG5cbiAgICBzd2l0Y2ggKHNjaGVtZSkge1xuICAgICAgY2FzZSAnbGlnaHQnOlxuICAgICAgICBib2R5LnN0eWxlLmNvbG9yU2NoZW1lID0gJ2xpZ2h0JztcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdkYXJrJzpcbiAgICAgICAgYm9keS5zdHlsZS5jb2xvclNjaGVtZSA9ICdkYXJrJztcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdzeXN0ZW0nOlxuICAgICAgZGVmYXVsdDpcbiAgICAgICAgYm9keS5zdHlsZS5jb2xvclNjaGVtZSA9ICdsaWdodCBkYXJrJztcbiAgICAgICAgYnJlYWs7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwS25vYkNvb3JkaW5hdGlvbigpIHtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOnByb3BlcnR5LWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOmNzcy1wcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjphdHRyaWJ1dGUtY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOnByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gIH1cblxuICAjb25Lbm9iQ2hhbmdlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMuI2dldEtub2JUYXJnZXQoZXZlbnQpO1xuICAgIGlmICghdGFyZ2V0KSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBDb3VsZCBub3QgZmluZCBrbm9iIHRhcmdldCBpbmZvIGluIGV2ZW50IHBhdGgnKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICBjb25zdCB7IHRhZ05hbWUsIGluc3RhbmNlSW5kZXggfSA9IHRhcmdldDtcblxuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCBrbm9iVHlwZSA9IHRoaXMuI2dldEtub2JUeXBlRnJvbUV2ZW50KGV2ZW50KTtcblxuICAgIGNvbnN0IHN1Y2Nlc3MgPSAoZGVtbyBhcyBhbnkpLmFwcGx5S25vYkNoYW5nZShcbiAgICAgIGtub2JUeXBlLFxuICAgICAgKGV2ZW50IGFzIGFueSkubmFtZSxcbiAgICAgIChldmVudCBhcyBhbnkpLnZhbHVlLFxuICAgICAgdGFnTmFtZSxcbiAgICAgIGluc3RhbmNlSW5kZXhcbiAgICApO1xuXG4gICAgaWYgKCFzdWNjZXNzKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gYXBwbHkga25vYiBjaGFuZ2U6Jywge1xuICAgICAgICB0eXBlOiBrbm9iVHlwZSxcbiAgICAgICAgbmFtZTogKGV2ZW50IGFzIGFueSkubmFtZSxcbiAgICAgICAgdGFnTmFtZSxcbiAgICAgICAgaW5zdGFuY2VJbmRleFxuICAgICAgfSk7XG4gICAgfVxuICB9O1xuXG4gICNvbktub2JDbGVhciA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLiNnZXRLbm9iVGFyZ2V0KGV2ZW50KTtcbiAgICBpZiAoIXRhcmdldCkge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gQ291bGQgbm90IGZpbmQga25vYiB0YXJnZXQgaW5mbyBpbiBldmVudCBwYXRoJyk7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgY29uc3QgeyB0YWdOYW1lLCBpbnN0YW5jZUluZGV4IH0gPSB0YXJnZXQ7XG5cbiAgICBjb25zdCBkZW1vID0gdGhpcy5kZW1vO1xuICAgIGlmICghZGVtbykgcmV0dXJuO1xuXG4gICAgY29uc3Qga25vYlR5cGUgPSB0aGlzLiNnZXRLbm9iVHlwZUZyb21DbGVhckV2ZW50KGV2ZW50KTtcbiAgICBjb25zdCBjbGVhclZhbHVlID0ga25vYlR5cGUgPT09ICdwcm9wZXJ0eScgPyB1bmRlZmluZWQgOiAnJztcblxuICAgIGNvbnN0IHN1Y2Nlc3MgPSAoZGVtbyBhcyBhbnkpLmFwcGx5S25vYkNoYW5nZShcbiAgICAgIGtub2JUeXBlLFxuICAgICAgKGV2ZW50IGFzIGFueSkubmFtZSxcbiAgICAgIGNsZWFyVmFsdWUsXG4gICAgICB0YWdOYW1lLFxuICAgICAgaW5zdGFuY2VJbmRleFxuICAgICk7XG5cbiAgICBpZiAoIXN1Y2Nlc3MpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjbGVhciBrbm9iOicsIHtcbiAgICAgICAgdHlwZToga25vYlR5cGUsXG4gICAgICAgIG5hbWU6IChldmVudCBhcyBhbnkpLm5hbWUsXG4gICAgICAgIHRhZ05hbWUsXG4gICAgICAgIGluc3RhbmNlSW5kZXhcbiAgICAgIH0pO1xuICAgIH1cbiAgfTtcblxuICAjZ2V0S25vYlRhcmdldChldmVudDogRXZlbnQpOiB7IHRhZ05hbWU6IHN0cmluZzsgaW5zdGFuY2VJbmRleDogbnVtYmVyIH0gfCBudWxsIHtcbiAgICBjb25zdCBkZWZhdWx0VGFnTmFtZSA9IHRoaXMucHJpbWFyeVRhZ05hbWUgfHwgJyc7XG5cbiAgICBpZiAoZXZlbnQuY29tcG9zZWRQYXRoKSB7XG4gICAgICBmb3IgKGNvbnN0IGVsZW1lbnQgb2YgZXZlbnQuY29tcG9zZWRQYXRoKCkpIHtcbiAgICAgICAgaWYgKCEoZWxlbWVudCBpbnN0YW5jZW9mIEVsZW1lbnQpKSBjb250aW51ZTtcblxuICAgICAgICBpZiAoKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQ/LmlzRWxlbWVudEtub2IgPT09ICd0cnVlJykge1xuICAgICAgICAgIGNvbnN0IHRhZ05hbWUgPSAoZWxlbWVudCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC50YWdOYW1lIHx8IGRlZmF1bHRUYWdOYW1lO1xuICAgICAgICAgIGxldCBpbnN0YW5jZUluZGV4ID0gTnVtYmVyLnBhcnNlSW50KChlbGVtZW50IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0Lmluc3RhbmNlSW5kZXggPz8gJycsIDEwKTtcbiAgICAgICAgICBpZiAoTnVtYmVyLmlzTmFOKGluc3RhbmNlSW5kZXgpKSBpbnN0YW5jZUluZGV4ID0gMDtcbiAgICAgICAgICByZXR1cm4geyB0YWdOYW1lLCBpbnN0YW5jZUluZGV4IH07XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICByZXR1cm4geyB0YWdOYW1lOiBkZWZhdWx0VGFnTmFtZSwgaW5zdGFuY2VJbmRleDogMCB9O1xuICB9XG5cbiAgI2dldEtub2JUeXBlRnJvbUV2ZW50KGV2ZW50OiBFdmVudCk6IHN0cmluZyB7XG4gICAgc3dpdGNoIChldmVudC50eXBlKSB7XG4gICAgICBjYXNlICdrbm9iOmF0dHJpYnV0ZS1jaGFuZ2UnOlxuICAgICAgICByZXR1cm4gJ2F0dHJpYnV0ZSc7XG4gICAgICBjYXNlICdrbm9iOnByb3BlcnR5LWNoYW5nZSc6XG4gICAgICAgIHJldHVybiAncHJvcGVydHknO1xuICAgICAgY2FzZSAna25vYjpjc3MtcHJvcGVydHktY2hhbmdlJzpcbiAgICAgICAgcmV0dXJuICdjc3MtcHJvcGVydHknO1xuICAgICAgZGVmYXVsdDpcbiAgICAgICAgcmV0dXJuICd1bmtub3duJztcbiAgICB9XG4gIH1cblxuICAjZ2V0S25vYlR5cGVGcm9tQ2xlYXJFdmVudChldmVudDogRXZlbnQpOiBzdHJpbmcge1xuICAgIHN3aXRjaCAoZXZlbnQudHlwZSkge1xuICAgICAgY2FzZSAna25vYjphdHRyaWJ1dGUtY2xlYXInOlxuICAgICAgICByZXR1cm4gJ2F0dHJpYnV0ZSc7XG4gICAgICBjYXNlICdrbm9iOnByb3BlcnR5LWNsZWFyJzpcbiAgICAgICAgcmV0dXJuICdwcm9wZXJ0eSc7XG4gICAgICBjYXNlICdrbm9iOmNzcy1wcm9wZXJ0eS1jbGVhcic6XG4gICAgICAgIHJldHVybiAnY3NzLXByb3BlcnR5JztcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIHJldHVybiAndW5rbm93bic7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwVHJlZVN0YXRlUGVyc2lzdGVuY2UoKSB7XG4gICAgdGhpcy4jaGFuZGxlVHJlZUV4cGFuZCA9IChlOiBFdmVudCkgPT4ge1xuICAgICAgaWYgKChlLnRhcmdldCBhcyBFbGVtZW50KT8udGFnTmFtZSAhPT0gJ0NFTS1QRi1WNi1UUkVFLUlURU0nKSByZXR1cm47XG5cbiAgICAgIGNvbnN0IG5vZGVJZCA9IHRoaXMuI2dldFRyZWVOb2RlSWQoZS50YXJnZXQgYXMgRWxlbWVudCk7XG4gICAgICBjb25zdCB0cmVlU3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFRyZWVTdGF0ZSgpO1xuICAgICAgaWYgKCF0cmVlU3RhdGUuZXhwYW5kZWQuaW5jbHVkZXMobm9kZUlkKSkge1xuICAgICAgICB0cmVlU3RhdGUuZXhwYW5kZWQucHVzaChub2RlSWQpO1xuICAgICAgICBTdGF0ZVBlcnNpc3RlbmNlLnNldFRyZWVTdGF0ZSh0cmVlU3RhdGUpO1xuICAgICAgfVxuICAgIH07XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdleHBhbmQnLCB0aGlzLiNoYW5kbGVUcmVlRXhwYW5kKTtcblxuICAgIHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSA9IChlOiBFdmVudCkgPT4ge1xuICAgICAgaWYgKChlLnRhcmdldCBhcyBFbGVtZW50KT8udGFnTmFtZSAhPT0gJ0NFTS1QRi1WNi1UUkVFLUlURU0nKSByZXR1cm47XG5cbiAgICAgIGNvbnN0IG5vZGVJZCA9IHRoaXMuI2dldFRyZWVOb2RlSWQoZS50YXJnZXQgYXMgRWxlbWVudCk7XG4gICAgICBjb25zdCB0cmVlU3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFRyZWVTdGF0ZSgpO1xuICAgICAgY29uc3QgaW5kZXggPSB0cmVlU3RhdGUuZXhwYW5kZWQuaW5kZXhPZihub2RlSWQpO1xuICAgICAgaWYgKGluZGV4ID4gLTEpIHtcbiAgICAgICAgdHJlZVN0YXRlLmV4cGFuZGVkLnNwbGljZShpbmRleCwgMSk7XG4gICAgICAgIFN0YXRlUGVyc2lzdGVuY2Uuc2V0VHJlZVN0YXRlKHRyZWVTdGF0ZSk7XG4gICAgICB9XG4gICAgfTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2NvbGxhcHNlJywgdGhpcy4jaGFuZGxlVHJlZUNvbGxhcHNlKTtcblxuICAgIHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QgPSAoZTogRXZlbnQpID0+IHtcbiAgICAgIGlmICgoZS50YXJnZXQgYXMgRWxlbWVudCk/LnRhZ05hbWUgIT09ICdDRU0tUEYtVjYtVFJFRS1JVEVNJykgcmV0dXJuO1xuXG4gICAgICBjb25zdCBub2RlSWQgPSB0aGlzLiNnZXRUcmVlTm9kZUlkKGUudGFyZ2V0IGFzIEVsZW1lbnQpO1xuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVUcmVlU3RhdGUoeyBzZWxlY3RlZDogbm9kZUlkIH0pO1xuICAgIH07XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdzZWxlY3QnLCB0aGlzLiNoYW5kbGVUcmVlU2VsZWN0KTtcblxuICAgIHRoaXMuI2FwcGx5VHJlZVN0YXRlKCk7XG4gIH1cblxuICAjYXBwbHlUcmVlU3RhdGUoKSB7XG4gICAgY29uc3QgdHJlZVN0YXRlID0gU3RhdGVQZXJzaXN0ZW5jZS5nZXRUcmVlU3RhdGUoKTtcblxuICAgIGZvciAoY29uc3Qgbm9kZUlkIG9mIHRyZWVTdGF0ZS5leHBhbmRlZCkge1xuICAgICAgY29uc3QgdHJlZUl0ZW0gPSB0aGlzLiNmaW5kVHJlZUl0ZW1CeUlkKG5vZGVJZCk7XG4gICAgICBpZiAodHJlZUl0ZW0gJiYgIXRyZWVJdGVtLmhhc0F0dHJpYnV0ZSgnZXhwYW5kZWQnKSkge1xuICAgICAgICB0cmVlSXRlbS5zZXRBdHRyaWJ1dGUoJ2V4cGFuZGVkJywgJycpO1xuICAgICAgfVxuICAgIH1cblxuICAgIGlmICh0cmVlU3RhdGUuc2VsZWN0ZWQpIHtcbiAgICAgIGNvbnN0IHRyZWVJdGVtID0gdGhpcy4jZmluZFRyZWVJdGVtQnlJZCh0cmVlU3RhdGUuc2VsZWN0ZWQpO1xuICAgICAgaWYgKHRyZWVJdGVtICYmICF0cmVlSXRlbS5oYXNBdHRyaWJ1dGUoJ2N1cnJlbnQnKSkge1xuICAgICAgICB0cmVlSXRlbS5zZXRBdHRyaWJ1dGUoJ2N1cnJlbnQnLCAnJyk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI3NldHVwU2lkZWJhclN0YXRlUGVyc2lzdGVuY2UoKSB7XG4gICAgY29uc3QgcGFnZSA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignY2VtLXBmLXY2LXBhZ2UnKTtcblxuICAgIGlmICghcGFnZSkgcmV0dXJuO1xuXG4gICAgcGFnZS5hZGRFdmVudExpc3RlbmVyKCdzaWRlYmFyLXRvZ2dsZScsIChldmVudDogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IGNvbGxhcHNlZCA9ICEoZXZlbnQgYXMgYW55KS5leHBhbmRlZDtcblxuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIHNpZGViYXI6IHsgY29sbGFwc2VkIH1cbiAgICAgIH0pO1xuICAgIH0pO1xuICB9XG5cbiAgI2ZpbmRUcmVlSXRlbUJ5SWQobm9kZUlkOiBzdHJpbmcpOiBFbGVtZW50IHwgbnVsbCB7XG4gICAgY29uc3QgcGFydHMgPSBub2RlSWQuc3BsaXQoJzonKTtcbiAgICBjb25zdCBbdHlwZSwgbW9kdWxlUGF0aCwgdGFnTmFtZSwgbmFtZV0gPSBwYXJ0cztcblxuICAgIGxldCBhdHRyU3VmZml4ID0gJyc7XG4gICAgaWYgKHRhZ05hbWUpIHtcbiAgICAgIGF0dHJTdWZmaXggKz0gYFtkYXRhLXRhZy1uYW1lPVwiJHtDU1MuZXNjYXBlKHRhZ05hbWUpfVwiXWA7XG4gICAgfVxuICAgIGlmIChuYW1lKSB7XG4gICAgICBhdHRyU3VmZml4ICs9IGBbZGF0YS1uYW1lPVwiJHtDU1MuZXNjYXBlKG5hbWUpfVwiXWA7XG4gICAgfVxuXG4gICAgbGV0IHNlbGVjdG9yID0gYGNlbS1wZi12Ni10cmVlLWl0ZW1bZGF0YS10eXBlPVwiJHtDU1MuZXNjYXBlKHR5cGUpfVwiXWA7XG4gICAgaWYgKG1vZHVsZVBhdGgpIHtcbiAgICAgIGNvbnN0IGVzY2FwZWRNb2R1bGVQYXRoID0gQ1NTLmVzY2FwZShtb2R1bGVQYXRoKTtcbiAgICAgIGNvbnN0IGVzY2FwZWRUeXBlID0gQ1NTLmVzY2FwZSh0eXBlKTtcbiAgICAgIGNvbnN0IHNlbGVjdG9yMSA9IGBjZW0tcGYtdjYtdHJlZS1pdGVtW2RhdGEtdHlwZT1cIiR7ZXNjYXBlZFR5cGV9XCJdW2RhdGEtbW9kdWxlLXBhdGg9XCIke2VzY2FwZWRNb2R1bGVQYXRofVwiXSR7YXR0clN1ZmZpeH1gO1xuICAgICAgY29uc3Qgc2VsZWN0b3IyID0gYGNlbS1wZi12Ni10cmVlLWl0ZW1bZGF0YS10eXBlPVwiJHtlc2NhcGVkVHlwZX1cIl1bZGF0YS1wYXRoPVwiJHtlc2NhcGVkTW9kdWxlUGF0aH1cIl0ke2F0dHJTdWZmaXh9YDtcbiAgICAgIHNlbGVjdG9yID0gYCR7c2VsZWN0b3IxfSwgJHtzZWxlY3RvcjJ9YDtcbiAgICB9IGVsc2Uge1xuICAgICAgc2VsZWN0b3IgKz0gYXR0clN1ZmZpeDtcbiAgICB9XG5cbiAgICByZXR1cm4gdGhpcy5xdWVyeVNlbGVjdG9yKHNlbGVjdG9yKTtcbiAgfVxuXG4gICNnZXRUcmVlTm9kZUlkKHRyZWVJdGVtOiBFbGVtZW50KTogc3RyaW5nIHtcbiAgICBjb25zdCB0eXBlID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXR5cGUnKTtcbiAgICBjb25zdCBtb2R1bGVQYXRoID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLW1vZHVsZS1wYXRoJykgfHwgdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXBhdGgnKTtcbiAgICBjb25zdCB0YWdOYW1lID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXRhZy1uYW1lJyk7XG4gICAgY29uc3QgbmFtZSA9IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS1uYW1lJyk7XG4gICAgY29uc3QgY2F0ZWdvcnkgPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtY2F0ZWdvcnknKTtcblxuICAgIGNvbnN0IHBhcnRzID0gW3R5cGVdO1xuICAgIGlmIChtb2R1bGVQYXRoKSBwYXJ0cy5wdXNoKG1vZHVsZVBhdGgpO1xuICAgIGlmICh0YWdOYW1lKSBwYXJ0cy5wdXNoKHRhZ05hbWUpO1xuICAgIGlmIChjYXRlZ29yeSkge1xuICAgICAgcGFydHMucHVzaChjYXRlZ29yeSk7XG4gICAgfSBlbHNlIGlmIChuYW1lKSB7XG4gICAgICBwYXJ0cy5wdXNoKG5hbWUpO1xuICAgIH1cblxuICAgIHJldHVybiBwYXJ0cy5qb2luKCc6Jyk7XG4gIH1cblxuICAvLyBFdmVudCBEaXNjb3ZlcnkgJiBDYXB0dXJlIE1ldGhvZHNcblxuICBhc3luYyAjZGlzY292ZXJFbGVtZW50RXZlbnRzKCk6IFByb21pc2U8TWFwPHN0cmluZywgRXZlbnRJbmZvPj4ge1xuICAgIHRyeSB7XG4gICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IGZldGNoKCcvY3VzdG9tLWVsZW1lbnRzLmpzb24nKTtcbiAgICAgIGlmICghcmVzcG9uc2Uub2spIHtcbiAgICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gTm8gbWFuaWZlc3QgYXZhaWxhYmxlIGZvciBldmVudCBkaXNjb3ZlcnknKTtcbiAgICAgICAgcmV0dXJuIG5ldyBNYXAoKTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgbWFuaWZlc3QgPSBhd2FpdCByZXNwb25zZS5qc29uKCkgYXMgTWFuaWZlc3Q7XG4gICAgICB0aGlzLiNtYW5pZmVzdCA9IG1hbmlmZXN0O1xuXG4gICAgICBjb25zdCBldmVudE1hcCA9IG5ldyBNYXA8c3RyaW5nLCBFdmVudEluZm8+KCk7XG5cbiAgICAgIGZvciAoY29uc3QgbW9kdWxlIG9mIG1hbmlmZXN0Lm1vZHVsZXMgfHwgW10pIHtcbiAgICAgICAgZm9yIChjb25zdCBkZWNsYXJhdGlvbiBvZiBtb2R1bGUuZGVjbGFyYXRpb25zIHx8IFtdKSB7XG4gICAgICAgICAgaWYgKGRlY2xhcmF0aW9uLmN1c3RvbUVsZW1lbnQgJiYgZGVjbGFyYXRpb24udGFnTmFtZSkge1xuICAgICAgICAgICAgY29uc3QgdGFnTmFtZSA9IGRlY2xhcmF0aW9uLnRhZ05hbWU7XG4gICAgICAgICAgICBjb25zdCBldmVudHMgPSBkZWNsYXJhdGlvbi5ldmVudHMgfHwgW107XG5cbiAgICAgICAgICAgIGlmIChldmVudHMubGVuZ3RoID4gMCkge1xuICAgICAgICAgICAgICBjb25zdCBldmVudE5hbWVzID0gbmV3IFNldChldmVudHMubWFwKGUgPT4gZS5uYW1lKSk7XG4gICAgICAgICAgICAgIGV2ZW50TWFwLnNldCh0YWdOYW1lLCB7XG4gICAgICAgICAgICAgICAgZXZlbnROYW1lcyxcbiAgICAgICAgICAgICAgICBldmVudHM6IGV2ZW50c1xuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgcmV0dXJuIGV2ZW50TWFwO1xuICAgIH0gY2F0Y2ggKGVycm9yKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBFcnJvciBsb2FkaW5nIG1hbmlmZXN0IGZvciBldmVudCBkaXNjb3Zlcnk6JywgZXJyb3IpO1xuICAgICAgcmV0dXJuIG5ldyBNYXAoKTtcbiAgICB9XG4gIH1cblxuICBhc3luYyAjc2V0dXBFdmVudENhcHR1cmUoKSB7XG4gICAgdGhpcy4jZWxlbWVudEV2ZW50TWFwID0gYXdhaXQgdGhpcy4jZGlzY292ZXJFbGVtZW50RXZlbnRzKCk7XG5cbiAgICBpZiAodGhpcy4jZWxlbWVudEV2ZW50TWFwLnNpemUgPT09IDApIHJldHVybjtcblxuICAgIHRoaXMuI2F0dGFjaEV2ZW50TGlzdGVuZXJzKCk7XG4gICAgdGhpcy4jdXBkYXRlRXZlbnRGaWx0ZXJzKCk7XG4gICAgdGhpcy4jb2JzZXJ2ZURlbW9NdXRhdGlvbnMoKTtcbiAgfVxuXG4gICNhdHRhY2hFdmVudExpc3RlbmVycygpIHtcbiAgICBjb25zdCBkZW1vID0gdGhpcy5kZW1vO1xuICAgIGlmICghZGVtbykgcmV0dXJuO1xuXG4gICAgY29uc3Qgcm9vdCA9IGRlbW8uc2hhZG93Um9vdCA/PyBkZW1vO1xuXG4gICAgZm9yIChjb25zdCBbdGFnTmFtZSwgZXZlbnRJbmZvXSBvZiB0aGlzLiNlbGVtZW50RXZlbnRNYXAhKSB7XG4gICAgICBjb25zdCBlbGVtZW50cyA9IHJvb3QucXVlcnlTZWxlY3RvckFsbCh0YWdOYW1lKTtcblxuICAgICAgZm9yIChjb25zdCBlbGVtZW50IG9mIGVsZW1lbnRzKSB7XG4gICAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGV2ZW50SW5mby5ldmVudE5hbWVzKSB7XG4gICAgICAgICAgZWxlbWVudC5hZGRFdmVudExpc3RlbmVyKGV2ZW50TmFtZSwgdGhpcy4jaGFuZGxlRWxlbWVudEV2ZW50LCB7IGNhcHR1cmU6IHRydWUgfSk7XG4gICAgICAgIH1cbiAgICAgICAgKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuY2VtRXZlbnRzQXR0YWNoZWQgPSAndHJ1ZSc7XG4gICAgICAgIHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5hZGQodGFnTmFtZSk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI29ic2VydmVEZW1vTXV0YXRpb25zKCkge1xuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCByb290ID0gZGVtby5zaGFkb3dSb290ID8/IGRlbW87XG5cbiAgICB0aGlzLiNvYnNlcnZlci5vYnNlcnZlKHJvb3QsIHtcbiAgICAgIGNoaWxkTGlzdDogdHJ1ZSxcbiAgICAgIHN1YnRyZWU6IHRydWVcbiAgICB9KTtcbiAgfVxuXG4gICNoYW5kbGVFbGVtZW50RXZlbnQgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgZWxlbWVudCA9IGV2ZW50LmN1cnJlbnRUYXJnZXQ7XG4gICAgaWYgKCEoZWxlbWVudCBpbnN0YW5jZW9mIEhUTUxFbGVtZW50KSkgcmV0dXJuO1xuXG4gICAgY29uc3QgdGFnTmFtZSA9IGVsZW1lbnQudGFnTmFtZS50b0xvd2VyQ2FzZSgpO1xuICAgIGNvbnN0IGV2ZW50SW5mbyA9IHRoaXMuI2VsZW1lbnRFdmVudE1hcD8uZ2V0KHRhZ05hbWUpO1xuXG4gICAgaWYgKCFldmVudEluZm8gfHwgIWV2ZW50SW5mby5ldmVudE5hbWVzLmhhcyhldmVudC50eXBlKSkgcmV0dXJuO1xuXG4gICAgdGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICB0aGlzLiNjYXB0dXJlRXZlbnQoZXZlbnQsIGVsZW1lbnQsIHRhZ05hbWUsIGV2ZW50SW5mbyk7XG4gIH07XG5cbiAgI2dldEV2ZW50RG9jdW1lbnRhdGlvbihtYW5pZmVzdEV2ZW50OiBFdmVudEluZm9bJ2V2ZW50cyddWzBdIHwgdW5kZWZpbmVkKSB7XG4gICAgaWYgKCFtYW5pZmVzdEV2ZW50KSB7XG4gICAgICByZXR1cm4geyBzdW1tYXJ5OiBudWxsLCBkZXNjcmlwdGlvbjogbnVsbCB9O1xuICAgIH1cblxuICAgIGxldCBzdW1tYXJ5ID0gbWFuaWZlc3RFdmVudC5zdW1tYXJ5IHx8IG51bGw7XG4gICAgbGV0IGRlc2NyaXB0aW9uID0gbWFuaWZlc3RFdmVudC5kZXNjcmlwdGlvbiB8fCBudWxsO1xuXG4gICAgaWYgKG1hbmlmZXN0RXZlbnQudHlwZT8udGV4dCAmJiB0aGlzLiNtYW5pZmVzdCkge1xuICAgICAgY29uc3QgdHlwZU5hbWUgPSBtYW5pZmVzdEV2ZW50LnR5cGUudGV4dDtcbiAgICAgIGNvbnN0IHR5cGVEZWNsYXJhdGlvbiA9IHRoaXMuI2ZpbmRUeXBlRGVjbGFyYXRpb24odHlwZU5hbWUpO1xuXG4gICAgICBpZiAodHlwZURlY2xhcmF0aW9uKSB7XG4gICAgICAgIGlmICghc3VtbWFyeSAmJiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSkge1xuICAgICAgICAgIHN1bW1hcnkgPSB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeTtcbiAgICAgICAgfSBlbHNlIGlmICh0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSAmJiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSAhPT0gc3VtbWFyeSkge1xuICAgICAgICAgIHN1bW1hcnkgPSBzdW1tYXJ5ID8gYCR7c3VtbWFyeX1cXG5cXG5Gcm9tICR7dHlwZU5hbWV9OiAke3R5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5fWAgOiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeTtcbiAgICAgICAgfVxuXG4gICAgICAgIGlmICghZGVzY3JpcHRpb24gJiYgdHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uKSB7XG4gICAgICAgICAgZGVzY3JpcHRpb24gPSB0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb247XG4gICAgICAgIH0gZWxzZSBpZiAodHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uICYmIHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbiAhPT0gZGVzY3JpcHRpb24pIHtcbiAgICAgICAgICBkZXNjcmlwdGlvbiA9IGRlc2NyaXB0aW9uID8gYCR7ZGVzY3JpcHRpb259XFxuXFxuJHt0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb259YCA6IHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbjtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiB7IHN1bW1hcnksIGRlc2NyaXB0aW9uIH07XG4gIH1cblxuICAjZmluZFR5cGVEZWNsYXJhdGlvbih0eXBlTmFtZTogc3RyaW5nKSB7XG4gICAgaWYgKCF0aGlzLiNtYW5pZmVzdCkgcmV0dXJuIG51bGw7XG5cbiAgICBmb3IgKGNvbnN0IG1vZHVsZSBvZiB0aGlzLiNtYW5pZmVzdC5tb2R1bGVzIHx8IFtdKSB7XG4gICAgICBmb3IgKGNvbnN0IGRlY2xhcmF0aW9uIG9mIG1vZHVsZS5kZWNsYXJhdGlvbnMgfHwgW10pIHtcbiAgICAgICAgaWYgKGRlY2xhcmF0aW9uLm5hbWUgPT09IHR5cGVOYW1lICYmXG4gICAgICAgICAgICAoZGVjbGFyYXRpb24ua2luZCA9PT0gJ2NsYXNzJyB8fCBkZWNsYXJhdGlvbi5raW5kID09PSAnaW50ZXJmYWNlJykpIHtcbiAgICAgICAgICByZXR1cm4gZGVjbGFyYXRpb24gYXMgeyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9O1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIG51bGw7XG4gIH1cblxuICAjY2FwdHVyZUV2ZW50KGV2ZW50OiBFdmVudCwgdGFyZ2V0OiBIVE1MRWxlbWVudCwgdGFnTmFtZTogc3RyaW5nLCBldmVudEluZm86IEV2ZW50SW5mbykge1xuICAgIGNvbnN0IG1hbmlmZXN0RXZlbnQgPSBldmVudEluZm8uZXZlbnRzLmZpbmQoZSA9PiBlLm5hbWUgPT09IGV2ZW50LnR5cGUpO1xuXG4gICAgY29uc3QgZXZlbnREb2NzID0gdGhpcy4jZ2V0RXZlbnREb2N1bWVudGF0aW9uKG1hbmlmZXN0RXZlbnQpO1xuXG4gICAgY29uc3QgY3VzdG9tUHJvcGVydGllcyA9IHRoaXMuI2V4dHJhY3RFdmVudFByb3BlcnRpZXMoZXZlbnQpO1xuXG4gICAgY29uc3QgZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkID0ge1xuICAgICAgaWQ6IGAke0RhdGUubm93KCl9LSR7TWF0aC5yYW5kb20oKX1gLFxuICAgICAgdGltZXN0YW1wOiBuZXcgRGF0ZSgpLFxuICAgICAgZXZlbnROYW1lOiBldmVudC50eXBlLFxuICAgICAgdGFnTmFtZTogdGFnTmFtZSxcbiAgICAgIGVsZW1lbnRJZDogdGFyZ2V0LmlkIHx8IG51bGwsXG4gICAgICBlbGVtZW50Q2xhc3M6IHRhcmdldC5jbGFzc05hbWUgfHwgbnVsbCxcbiAgICAgIGN1c3RvbVByb3BlcnRpZXM6IGN1c3RvbVByb3BlcnRpZXMsXG4gICAgICBtYW5pZmVzdFR5cGU6IG1hbmlmZXN0RXZlbnQ/LnR5cGU/LnRleHQgfHwgbnVsbCxcbiAgICAgIHN1bW1hcnk6IGV2ZW50RG9jcy5zdW1tYXJ5LFxuICAgICAgZGVzY3JpcHRpb246IGV2ZW50RG9jcy5kZXNjcmlwdGlvbixcbiAgICAgIGJ1YmJsZXM6IGV2ZW50LmJ1YmJsZXMsXG4gICAgICBjb21wb3NlZDogZXZlbnQuY29tcG9zZWQsXG4gICAgICBjYW5jZWxhYmxlOiBldmVudC5jYW5jZWxhYmxlLFxuICAgICAgZGVmYXVsdFByZXZlbnRlZDogZXZlbnQuZGVmYXVsdFByZXZlbnRlZFxuICAgIH07XG5cbiAgICB0aGlzLiNjYXB0dXJlZEV2ZW50cy5wdXNoKGV2ZW50UmVjb3JkKTtcblxuICAgIGlmICh0aGlzLiNjYXB0dXJlZEV2ZW50cy5sZW5ndGggPiB0aGlzLiNtYXhDYXB0dXJlZEV2ZW50cykge1xuICAgICAgdGhpcy4jY2FwdHVyZWRFdmVudHMuc2hpZnQoKTtcbiAgICB9XG5cbiAgICB0aGlzLiNyZW5kZXJFdmVudChldmVudFJlY29yZCk7XG4gIH1cblxuICAjZXh0cmFjdEV2ZW50UHJvcGVydGllcyhldmVudDogRXZlbnQpOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiB7XG4gICAgY29uc3QgcHJvcGVydGllczogUmVjb3JkPHN0cmluZywgdW5rbm93bj4gPSB7fTtcbiAgICBjb25zdCBldmVudFByb3RvdHlwZUtleXMgPSBuZXcgU2V0KE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKEV2ZW50LnByb3RvdHlwZSkpO1xuXG4gICAgY29uc3Qgc2VyaWFsaXplVmFsdWUgPSAodmFsdWU6IHVua25vd24pOiB1bmtub3duID0+IHtcbiAgICAgIHRyeSB7XG4gICAgICAgIHJldHVybiBKU09OLnBhcnNlKEpTT04uc3RyaW5naWZ5KHZhbHVlKSk7XG4gICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgIHRyeSB7XG4gICAgICAgICAgcmV0dXJuIFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIH0gY2F0Y2ggKHN0cmluZ0Vycikge1xuICAgICAgICAgIHJldHVybiAnW05vdCBzZXJpYWxpemFibGVdJztcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH07XG5cbiAgICBpZiAoZXZlbnQgaW5zdGFuY2VvZiBDdXN0b21FdmVudCAmJiBldmVudC5kZXRhaWwgIT09IHVuZGVmaW5lZCkge1xuICAgICAgcHJvcGVydGllcy5kZXRhaWwgPSBzZXJpYWxpemVWYWx1ZShldmVudC5kZXRhaWwpO1xuICAgIH1cblxuICAgIGZvciAoY29uc3Qga2V5IG9mIE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKGV2ZW50KSkge1xuICAgICAgaWYgKCFldmVudFByb3RvdHlwZUtleXMuaGFzKGtleSkgJiYgIWtleS5zdGFydHNXaXRoKCdfJykgJiYgIXByb3BlcnRpZXMuaGFzT3duUHJvcGVydHkoa2V5KSkge1xuICAgICAgICBwcm9wZXJ0aWVzW2tleV0gPSBzZXJpYWxpemVWYWx1ZSgoZXZlbnQgYXMgYW55KVtrZXldKTtcbiAgICAgIH1cbiAgICB9XG5cbiAgICByZXR1cm4gcHJvcGVydGllcztcbiAgfVxuXG4gICNyZW5kZXJFdmVudChldmVudFJlY29yZDogRXZlbnRSZWNvcmQpIHtcbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgY29uc3QgZnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZXZlbnRFbnRyeVRlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICBjb25zdCB0aW1lID0gZXZlbnRSZWNvcmQudGltZXN0YW1wLnRvTG9jYWxlVGltZVN0cmluZygpO1xuXG4gICAgY29uc3QgY29udGFpbmVyID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJjb250YWluZXJcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICBjb250YWluZXIuZGF0YXNldC5ldmVudElkID0gZXZlbnRSZWNvcmQuaWQ7XG4gICAgY29udGFpbmVyLmRhdGFzZXQuZXZlbnRUeXBlID0gZXZlbnRSZWNvcmQuZXZlbnROYW1lO1xuICAgIGNvbnRhaW5lci5kYXRhc2V0LmVsZW1lbnRUeXBlID0gZXZlbnRSZWNvcmQudGFnTmFtZTtcblxuICAgIGNvbnN0IHRleHRNYXRjaCA9IHRoaXMuI2V2ZW50TWF0Y2hlc1RleHRGaWx0ZXIoZXZlbnRSZWNvcmQpO1xuICAgIGNvbnN0IHR5cGVNYXRjaCA9IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmhhcyhldmVudFJlY29yZC5ldmVudE5hbWUpO1xuICAgIGNvbnN0IGVsZW1lbnRNYXRjaCA9IHRoaXMuI2VsZW1lbnRGaWx0ZXJzLnNpemUgPT09IDAgfHwgdGhpcy4jZWxlbWVudEZpbHRlcnMuaGFzKGV2ZW50UmVjb3JkLnRhZ05hbWUpO1xuXG4gICAgaWYgKCEodGV4dE1hdGNoICYmIHR5cGVNYXRjaCAmJiBlbGVtZW50TWF0Y2gpKSB7XG4gICAgICBjb250YWluZXIuc2V0QXR0cmlidXRlKCdoaWRkZW4nLCAnJyk7XG4gICAgfVxuXG4gICAgY29uc3QgbGFiZWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxhYmVsXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgbGFiZWwudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC5ldmVudE5hbWU7XG4gICAgbGFiZWwuc2V0QXR0cmlidXRlKCdzdGF0dXMnLCAnaW5mbycpO1xuXG4gICAgY29uc3QgdGltZUVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0aW1lXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgdGltZUVsLnNldEF0dHJpYnV0ZSgnZGF0ZXRpbWUnLCBldmVudFJlY29yZC50aW1lc3RhbXAudG9JU09TdHJpbmcoKSk7XG4gICAgdGltZUVsLnRleHRDb250ZW50ID0gdGltZTtcblxuICAgIGNvbnN0IGVsZW1lbnRFbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZWxlbWVudFwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgIGxldCBlbGVtZW50VGV4dCA9IGA8JHtldmVudFJlY29yZC50YWdOYW1lfT5gO1xuICAgIGlmIChldmVudFJlY29yZC5lbGVtZW50SWQpIHtcbiAgICAgIGVsZW1lbnRUZXh0ICs9IGAjJHtldmVudFJlY29yZC5lbGVtZW50SWR9YDtcbiAgICB9XG4gICAgZWxlbWVudEVsLnRleHRDb250ZW50ID0gZWxlbWVudFRleHQ7XG5cbiAgICB0aGlzLiNldmVudExpc3QuYXBwZW5kKGZyYWdtZW50KTtcblxuICAgIGlmICghdGhpcy4jc2VsZWN0ZWRFdmVudElkKSB7XG4gICAgICB0aGlzLiNzZWxlY3RFdmVudChldmVudFJlY29yZC5pZCk7XG4gICAgfVxuXG4gICAgaWYgKHRoaXMuI2RyYXdlck9wZW4gJiYgdGhpcy4jaXNFdmVudHNUYWJBY3RpdmUoKSkge1xuICAgICAgdGhpcy4jc2Nyb2xsRXZlbnRzVG9Cb3R0b20oKTtcbiAgICB9XG4gIH1cblxuICAjc2VsZWN0RXZlbnQoZXZlbnRJZDogc3RyaW5nKSB7XG4gICAgY29uc3QgZXZlbnRSZWNvcmQgPSB0aGlzLiNnZXRFdmVudFJlY29yZEJ5SWQoZXZlbnRJZCk7XG4gICAgaWYgKCFldmVudFJlY29yZCkgcmV0dXJuO1xuXG4gICAgdGhpcy4jc2VsZWN0ZWRFdmVudElkID0gZXZlbnRJZDtcblxuICAgIGNvbnN0IGFsbEl0ZW1zID0gdGhpcy4jZXZlbnRMaXN0Py5xdWVyeVNlbGVjdG9yQWxsKCcuZXZlbnQtbGlzdC1pdGVtJyk7XG4gICAgYWxsSXRlbXM/LmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBpZiAoKGl0ZW0gYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuZXZlbnRJZCA9PT0gZXZlbnRJZCkge1xuICAgICAgICBpdGVtLmNsYXNzTGlzdC5hZGQoJ3NlbGVjdGVkJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCdhcmlhLXNlbGVjdGVkJywgJ3RydWUnKTtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGl0ZW0uY2xhc3NMaXN0LnJlbW92ZSgnc2VsZWN0ZWQnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2FyaWEtc2VsZWN0ZWQnLCAnZmFsc2UnKTtcbiAgICAgIH1cbiAgICB9KTtcblxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEhlYWRlcikge1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIuaW5uZXJIVE1MID0gJyc7XG5cbiAgICAgIGNvbnN0IGhlYWRlckNvbnRlbnQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIGhlYWRlckNvbnRlbnQuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1oZWFkZXItY29udGVudCc7XG5cbiAgICAgIGNvbnN0IGV2ZW50TmFtZSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2gzJyk7XG4gICAgICBldmVudE5hbWUudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC5ldmVudE5hbWU7XG4gICAgICBldmVudE5hbWUuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1uYW1lJztcbiAgICAgIGhlYWRlckNvbnRlbnQuYXBwZW5kQ2hpbGQoZXZlbnROYW1lKTtcblxuICAgICAgaWYgKGV2ZW50UmVjb3JkLnN1bW1hcnkpIHtcbiAgICAgICAgY29uc3Qgc3VtbWFyeSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3AnKTtcbiAgICAgICAgc3VtbWFyeS50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLnN1bW1hcnk7XG4gICAgICAgIHN1bW1hcnkuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1zdW1tYXJ5JztcbiAgICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChzdW1tYXJ5KTtcbiAgICAgIH1cblxuICAgICAgaWYgKGV2ZW50UmVjb3JkLmRlc2NyaXB0aW9uKSB7XG4gICAgICAgIGNvbnN0IGRlc2NyaXB0aW9uID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncCcpO1xuICAgICAgICBkZXNjcmlwdGlvbi50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLmRlc2NyaXB0aW9uO1xuICAgICAgICBkZXNjcmlwdGlvbi5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLWRlc2NyaXB0aW9uJztcbiAgICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChkZXNjcmlwdGlvbik7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IG1ldGEgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIG1ldGEuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1tZXRhJztcblxuICAgICAgY29uc3QgdGltZUVsID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGltZScpO1xuICAgICAgdGltZUVsLnNldEF0dHJpYnV0ZSgnZGF0ZXRpbWUnLCBldmVudFJlY29yZC50aW1lc3RhbXAudG9JU09TdHJpbmcoKSk7XG4gICAgICB0aW1lRWwudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC50aW1lc3RhbXAudG9Mb2NhbGVUaW1lU3RyaW5nKCk7XG4gICAgICB0aW1lRWwuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC10aW1lJztcblxuICAgICAgY29uc3QgZWxlbWVudCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgIGxldCBlbGVtZW50VGV4dCA9IGA8JHtldmVudFJlY29yZC50YWdOYW1lfT5gO1xuICAgICAgaWYgKGV2ZW50UmVjb3JkLmVsZW1lbnRJZCkge1xuICAgICAgICBlbGVtZW50VGV4dCArPSBgIyR7ZXZlbnRSZWNvcmQuZWxlbWVudElkfWA7XG4gICAgICB9XG4gICAgICBlbGVtZW50LnRleHRDb250ZW50ID0gZWxlbWVudFRleHQ7XG4gICAgICBlbGVtZW50LmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtZWxlbWVudCc7XG5cbiAgICAgIG1ldGEuYXBwZW5kQ2hpbGQodGltZUVsKTtcbiAgICAgIG1ldGEuYXBwZW5kQ2hpbGQoZWxlbWVudCk7XG5cbiAgICAgIGhlYWRlckNvbnRlbnQuYXBwZW5kQ2hpbGQobWV0YSk7XG5cbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyLmFwcGVuZENoaWxkKGhlYWRlckNvbnRlbnQpO1xuICAgIH1cblxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEJvZHkpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5pbm5lckhUTUwgPSAnJztcblxuICAgICAgY29uc3QgcHJvcGVydGllc0hlYWRpbmcgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdoNCcpO1xuICAgICAgcHJvcGVydGllc0hlYWRpbmcudGV4dENvbnRlbnQgPSAnUHJvcGVydGllcyc7XG4gICAgICBwcm9wZXJ0aWVzSGVhZGluZy5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMtaGVhZGluZyc7XG5cbiAgICAgIGNvbnN0IHByb3BlcnRpZXNDb250YWluZXIgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIHByb3BlcnRpZXNDb250YWluZXIuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1wcm9wZXJ0aWVzJztcblxuICAgICAgY29uc3QgZXZlbnRQcm9wZXJ0aWVzID0gdGhpcy4jYnVpbGRQcm9wZXJ0aWVzRm9yRGlzcGxheShldmVudFJlY29yZCk7XG4gICAgICBpZiAoT2JqZWN0LmtleXMoZXZlbnRQcm9wZXJ0aWVzKS5sZW5ndGggPiAwKSB7XG4gICAgICAgIHByb3BlcnRpZXNDb250YWluZXIuYXBwZW5kQ2hpbGQodGhpcy4jYnVpbGRQcm9wZXJ0eVRyZWUoZXZlbnRQcm9wZXJ0aWVzKSk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBwcm9wZXJ0aWVzQ29udGFpbmVyLnRleHRDb250ZW50ID0gJ05vIHByb3BlcnRpZXMgdG8gZGlzcGxheSc7XG4gICAgICB9XG5cbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5hcHBlbmRDaGlsZChwcm9wZXJ0aWVzSGVhZGluZyk7XG4gICAgICB0aGlzLiNldmVudERldGFpbEJvZHkuYXBwZW5kQ2hpbGQocHJvcGVydGllc0NvbnRhaW5lcik7XG4gICAgfVxuICB9XG5cbiAgI2J1aWxkUHJvcGVydGllc0ZvckRpc3BsYXkoZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkKTogUmVjb3JkPHN0cmluZywgdW5rbm93bj4ge1xuICAgIGNvbnN0IHByb3BlcnRpZXM6IFJlY29yZDxzdHJpbmcsIHVua25vd24+ID0ge307XG5cbiAgICBpZiAoZXZlbnRSZWNvcmQuY3VzdG9tUHJvcGVydGllcykge1xuICAgICAgT2JqZWN0LmFzc2lnbihwcm9wZXJ0aWVzLCBldmVudFJlY29yZC5jdXN0b21Qcm9wZXJ0aWVzKTtcbiAgICB9XG5cbiAgICBwcm9wZXJ0aWVzLmJ1YmJsZXMgPSBldmVudFJlY29yZC5idWJibGVzO1xuICAgIHByb3BlcnRpZXMuY2FuY2VsYWJsZSA9IGV2ZW50UmVjb3JkLmNhbmNlbGFibGU7XG4gICAgcHJvcGVydGllcy5kZWZhdWx0UHJldmVudGVkID0gZXZlbnRSZWNvcmQuZGVmYXVsdFByZXZlbnRlZDtcbiAgICBwcm9wZXJ0aWVzLmNvbXBvc2VkID0gZXZlbnRSZWNvcmQuY29tcG9zZWQ7XG5cbiAgICBpZiAoZXZlbnRSZWNvcmQubWFuaWZlc3RUeXBlKSB7XG4gICAgICBwcm9wZXJ0aWVzLnR5cGUgPSBldmVudFJlY29yZC5tYW5pZmVzdFR5cGU7XG4gICAgfVxuXG4gICAgcmV0dXJuIHByb3BlcnRpZXM7XG4gIH1cblxuICAjYnVpbGRQcm9wZXJ0eVRyZWUob2JqOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiwgZGVwdGggPSAwKTogSFRNTFVMaXN0RWxlbWVudCB7XG4gICAgY29uc3QgdWwgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd1bCcpO1xuICAgIHVsLmNsYXNzTmFtZSA9ICdldmVudC1wcm9wZXJ0eS10cmVlJztcbiAgICBpZiAoZGVwdGggPiAwKSB7XG4gICAgICB1bC5jbGFzc0xpc3QuYWRkKCduZXN0ZWQnKTtcbiAgICB9XG5cbiAgICBmb3IgKGNvbnN0IFtrZXksIHZhbHVlXSBvZiBPYmplY3QuZW50cmllcyhvYmopKSB7XG4gICAgICBjb25zdCBsaSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2xpJyk7XG4gICAgICBsaS5jbGFzc05hbWUgPSAncHJvcGVydHktaXRlbSc7XG5cbiAgICAgIGNvbnN0IGtleVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICBrZXlTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS1rZXknO1xuICAgICAga2V5U3Bhbi50ZXh0Q29udGVudCA9IGtleTtcblxuICAgICAgY29uc3QgY29sb25TcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgY29sb25TcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS1jb2xvbic7XG4gICAgICBjb2xvblNwYW4udGV4dENvbnRlbnQgPSAnOiAnO1xuXG4gICAgICBsaS5hcHBlbmRDaGlsZChrZXlTcGFuKTtcbiAgICAgIGxpLmFwcGVuZENoaWxkKGNvbG9uU3Bhbik7XG5cbiAgICAgIGlmICh2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBudWxsJztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH0gZWxzZSBpZiAodHlwZW9mIHZhbHVlID09PSAnYm9vbGVhbicpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIGJvb2xlYW4nO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdudW1iZXInKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBudW1iZXInO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdzdHJpbmcnKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBzdHJpbmcnO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBgXCIke3ZhbHVlfVwiYDtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH0gZWxzZSBpZiAoQXJyYXkuaXNBcnJheSh2YWx1ZSkpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIGFycmF5JztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gYEFycmF5KCR7dmFsdWUubGVuZ3RofSlgO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuXG4gICAgICAgIGlmICh2YWx1ZS5sZW5ndGggPiAwICYmIGRlcHRoIDwgMykge1xuICAgICAgICAgIGNvbnN0IG5lc3RlZE9iajogUmVjb3JkPHN0cmluZywgdW5rbm93bj4gPSB7fTtcbiAgICAgICAgICB2YWx1ZS5mb3JFYWNoKChpdGVtLCBpbmRleCkgPT4ge1xuICAgICAgICAgICAgbmVzdGVkT2JqW2luZGV4XSA9IGl0ZW07XG4gICAgICAgICAgfSk7XG4gICAgICAgICAgbGkuYXBwZW5kQ2hpbGQodGhpcy4jYnVpbGRQcm9wZXJ0eVRyZWUobmVzdGVkT2JqLCBkZXB0aCArIDEpKTtcbiAgICAgICAgfVxuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdvYmplY3QnKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBvYmplY3QnO1xuICAgICAgICBjb25zdCBrZXlzID0gT2JqZWN0LmtleXModmFsdWUgYXMgUmVjb3JkPHN0cmluZywgdW5rbm93bj4pO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBrZXlzLmxlbmd0aCA+IDAgPyBgT2JqZWN0YCA6IGB7fWA7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG5cbiAgICAgICAgaWYgKGtleXMubGVuZ3RoID4gMCAmJiBkZXB0aCA8IDMpIHtcbiAgICAgICAgICBsaS5hcHBlbmRDaGlsZCh0aGlzLiNidWlsZFByb3BlcnR5VHJlZSh2YWx1ZSBhcyBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiwgZGVwdGggKyAxKSk7XG4gICAgICAgIH1cbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSc7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9XG5cbiAgICAgIHVsLmFwcGVuZENoaWxkKGxpKTtcbiAgICB9XG5cbiAgICByZXR1cm4gdWw7XG4gIH1cblxuICAjc2Nyb2xsRXZlbnRzVG9Cb3R0b20oKSB7XG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICBjb25zdCBsYXN0RXZlbnQgPSB0aGlzLiNldmVudExpc3QhLmxhc3RFbGVtZW50Q2hpbGQ7XG4gICAgICBpZiAobGFzdEV2ZW50KSB7XG4gICAgICAgIGxhc3RFdmVudC5zY3JvbGxJbnRvVmlldyh7IGJlaGF2aW9yOiAnYXV0bycsIGJsb2NrOiAnZW5kJyB9KTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNpc0V2ZW50c1RhYkFjdGl2ZSgpOiBib29sZWFuIHtcbiAgICBjb25zdCB0YWJzID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdjZW0tcGYtdjYtdGFicycpO1xuICAgIGlmICghdGFicykgcmV0dXJuIGZhbHNlO1xuXG4gICAgY29uc3Qgc2VsZWN0ZWRJbmRleCA9IHBhcnNlSW50KHRhYnMuZ2V0QXR0cmlidXRlKCdzZWxlY3RlZCcpIHx8ICcwJywgMTApO1xuICAgIHJldHVybiBzZWxlY3RlZEluZGV4ID09PSAzO1xuICB9XG5cbiAgI2ZpbHRlckV2ZW50cyhxdWVyeTogc3RyaW5nKSB7XG4gICAgdGhpcy4jZXZlbnRzRmlsdGVyVmFsdWUgPSBxdWVyeS50b0xvd2VyQ2FzZSgpO1xuXG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIGZvciAoY29uc3QgZW50cnkgb2YgdGhpcy4jZXZlbnRMaXN0LmNoaWxkcmVuKSB7XG4gICAgICBjb25zdCBldmVudFJlY29yZCA9IHRoaXMuI2dldEV2ZW50UmVjb3JkQnlJZCgoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuZXZlbnRJZCEpO1xuXG4gICAgICBpZiAoIWV2ZW50UmVjb3JkKSBjb250aW51ZTtcblxuICAgICAgY29uc3QgdGV4dE1hdGNoID0gdGhpcy4jZXZlbnRNYXRjaGVzVGV4dEZpbHRlcihldmVudFJlY29yZCk7XG4gICAgICBjb25zdCB0eXBlTWF0Y2ggPSB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLnNpemUgPT09IDAgfHwgdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5oYXMoZXZlbnRSZWNvcmQuZXZlbnROYW1lKTtcbiAgICAgIGNvbnN0IGVsZW1lbnRNYXRjaCA9IHRoaXMuI2VsZW1lbnRGaWx0ZXJzLnNpemUgPT09IDAgfHwgdGhpcy4jZWxlbWVudEZpbHRlcnMuaGFzKGV2ZW50UmVjb3JkLnRhZ05hbWUpO1xuXG4gICAgICAoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmhpZGRlbiA9ICEodGV4dE1hdGNoICYmIHR5cGVNYXRjaCAmJiBlbGVtZW50TWF0Y2gpO1xuICAgIH1cbiAgfVxuXG4gICNldmVudE1hdGNoZXNUZXh0RmlsdGVyKGV2ZW50UmVjb3JkOiBFdmVudFJlY29yZCk6IGJvb2xlYW4ge1xuICAgIGlmICghdGhpcy4jZXZlbnRzRmlsdGVyVmFsdWUpIHJldHVybiB0cnVlO1xuXG4gICAgY29uc3Qgc2VhcmNoVGV4dCA9IFtcbiAgICAgIGV2ZW50UmVjb3JkLnRhZ05hbWUsXG4gICAgICBldmVudFJlY29yZC5ldmVudE5hbWUsXG4gICAgICBldmVudFJlY29yZC5lbGVtZW50SWQgfHwgJycsXG4gICAgICBKU09OLnN0cmluZ2lmeShldmVudFJlY29yZC5jdXN0b21Qcm9wZXJ0aWVzIHx8IHt9KVxuICAgIF0uam9pbignICcpLnRvTG93ZXJDYXNlKCk7XG5cbiAgICByZXR1cm4gc2VhcmNoVGV4dC5pbmNsdWRlcyh0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSk7XG4gIH1cblxuICAjZ2V0RXZlbnRSZWNvcmRCeUlkKGlkOiBzdHJpbmcpOiBFdmVudFJlY29yZCB8IHVuZGVmaW5lZCB7XG4gICAgcmV0dXJuIHRoaXMuI2NhcHR1cmVkRXZlbnRzLmZpbmQoZSA9PiBlLmlkID09PSBpZCk7XG4gIH1cblxuICAjdXBkYXRlRXZlbnRGaWx0ZXJzKCkge1xuICAgIGNvbnN0IHNhdmVkUHJlZmVyZW5jZXMgPSB0aGlzLiNsb2FkRXZlbnRGaWx0ZXJzRnJvbVN0b3JhZ2UoKTtcblxuICAgIGNvbnN0IGV2ZW50VHlwZUZpbHRlciA9IHRoaXMuIyQoJ2V2ZW50LXR5cGUtZmlsdGVyJyk7XG4gICAgaWYgKGV2ZW50VHlwZUZpbHRlciAmJiB0aGlzLiNlbGVtZW50RXZlbnRNYXApIHtcbiAgICAgIGxldCBtZW51ID0gZXZlbnRUeXBlRmlsdGVyLnF1ZXJ5U2VsZWN0b3IoJ2NlbS1wZi12Ni1tZW51Jyk7XG4gICAgICBpZiAoIW1lbnUpIHtcbiAgICAgICAgbWVudSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2NlbS1wZi12Ni1tZW51Jyk7XG4gICAgICAgIGV2ZW50VHlwZUZpbHRlci5hcHBlbmRDaGlsZChtZW51KTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgZXhpc3RpbmdJdGVtcyA9IG1lbnUucXVlcnlTZWxlY3RvckFsbCgnY2VtLXBmLXY2LW1lbnUtaXRlbScpO1xuICAgICAgZXhpc3RpbmdJdGVtcy5mb3JFYWNoKGl0ZW0gPT4gaXRlbS5yZW1vdmUoKSk7XG5cbiAgICAgIGNvbnN0IGFsbEV2ZW50VHlwZXMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgICAgIGZvciAoY29uc3QgW3RhZ05hbWUsIGV2ZW50SW5mb10gb2YgdGhpcy4jZWxlbWVudEV2ZW50TWFwKSB7XG4gICAgICAgIGlmICh0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuaGFzKHRhZ05hbWUpKSB7XG4gICAgICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgZXZlbnRJbmZvLmV2ZW50TmFtZXMpIHtcbiAgICAgICAgICAgIGFsbEV2ZW50VHlwZXMuYWRkKGV2ZW50TmFtZSk7XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIGlmIChzYXZlZFByZWZlcmVuY2VzLmV2ZW50VHlwZXMpIHtcbiAgICAgICAgdGhpcy4jZXZlbnRUeXBlRmlsdGVycyA9IChzYXZlZFByZWZlcmVuY2VzLmV2ZW50VHlwZXMgYXMgU2V0PHN0cmluZz4gJiB7IGludGVyc2VjdGlvbjogKG90aGVyOiBTZXQ8c3RyaW5nPikgPT4gU2V0PHN0cmluZz4gfSkuaW50ZXJzZWN0aW9uKGFsbEV2ZW50VHlwZXMpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgdGhpcy4jZXZlbnRUeXBlRmlsdGVycyA9IG5ldyBTZXQoYWxsRXZlbnRUeXBlcyk7XG4gICAgICB9XG5cbiAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGFsbEV2ZW50VHlwZXMpIHtcbiAgICAgICAgY29uc3QgaXRlbSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2NlbS1wZi12Ni1tZW51LWl0ZW0nKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3ZhcmlhbnQnLCAnY2hlY2tib3gnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3ZhbHVlJywgZXZlbnROYW1lKTtcbiAgICAgICAgaWYgKHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuaGFzKGV2ZW50TmFtZSkpIHtcbiAgICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnY2hlY2tlZCcsICcnKTtcbiAgICAgICAgfVxuICAgICAgICBpdGVtLnRleHRDb250ZW50ID0gZXZlbnROYW1lO1xuICAgICAgICBtZW51LmFwcGVuZENoaWxkKGl0ZW0pO1xuICAgICAgfVxuICAgIH1cblxuICAgIGNvbnN0IGVsZW1lbnRGaWx0ZXIgPSB0aGlzLiMkKCdlbGVtZW50LWZpbHRlcicpO1xuICAgIGlmIChlbGVtZW50RmlsdGVyICYmIHRoaXMuI2VsZW1lbnRFdmVudE1hcCkge1xuICAgICAgbGV0IG1lbnUgPSBlbGVtZW50RmlsdGVyLnF1ZXJ5U2VsZWN0b3IoJ2NlbS1wZi12Ni1tZW51Jyk7XG4gICAgICBpZiAoIW1lbnUpIHtcbiAgICAgICAgbWVudSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2NlbS1wZi12Ni1tZW51Jyk7XG4gICAgICAgIGVsZW1lbnRGaWx0ZXIuYXBwZW5kQ2hpbGQobWVudSk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGV4aXN0aW5nSXRlbXMgPSBtZW51LnF1ZXJ5U2VsZWN0b3JBbGwoJ2NlbS1wZi12Ni1tZW51LWl0ZW0nKTtcbiAgICAgIGV4aXN0aW5nSXRlbXMuZm9yRWFjaChpdGVtID0+IGl0ZW0ucmVtb3ZlKCkpO1xuXG4gICAgICBjb25zdCBhbGxFbGVtZW50cyA9IG5ldyBTZXQ8c3RyaW5nPigpO1xuICAgICAgZm9yIChjb25zdCB0YWdOYW1lIG9mIHRoaXMuI2VsZW1lbnRFdmVudE1hcC5rZXlzKCkpIHtcbiAgICAgICAgaWYgKHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5oYXModGFnTmFtZSkpIHtcbiAgICAgICAgICBhbGxFbGVtZW50cy5hZGQodGFnTmFtZSk7XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgaWYgKHNhdmVkUHJlZmVyZW5jZXMuZWxlbWVudHMpIHtcbiAgICAgICAgdGhpcy4jZWxlbWVudEZpbHRlcnMgPSAoc2F2ZWRQcmVmZXJlbmNlcy5lbGVtZW50cyBhcyBTZXQ8c3RyaW5nPiAmIHsgaW50ZXJzZWN0aW9uOiAob3RoZXI6IFNldDxzdHJpbmc+KSA9PiBTZXQ8c3RyaW5nPiB9KS5pbnRlcnNlY3Rpb24oYWxsRWxlbWVudHMpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgdGhpcy4jZWxlbWVudEZpbHRlcnMgPSBuZXcgU2V0KGFsbEVsZW1lbnRzKTtcbiAgICAgIH1cblxuICAgICAgZm9yIChjb25zdCB0YWdOYW1lIG9mIGFsbEVsZW1lbnRzKSB7XG4gICAgICAgIGNvbnN0IGl0ZW0gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdjZW0tcGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YXJpYW50JywgJ2NoZWNrYm94Jyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YWx1ZScsIHRhZ05hbWUpO1xuICAgICAgICBpZiAodGhpcy4jZWxlbWVudEZpbHRlcnMuaGFzKHRhZ05hbWUpKSB7XG4gICAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2NoZWNrZWQnLCAnJyk7XG4gICAgICAgIH1cbiAgICAgICAgaXRlbS50ZXh0Q29udGVudCA9IGA8JHt0YWdOYW1lfT5gO1xuICAgICAgICBtZW51LmFwcGVuZENoaWxkKGl0ZW0pO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNoYW5kbGVFdmVudFR5cGVGaWx0ZXJDaGFuZ2UgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgeyB2YWx1ZSwgY2hlY2tlZCB9ID0gZXZlbnQgYXMgRXZlbnQgJiB7IHZhbHVlOiBzdHJpbmc7IGNoZWNrZWQ6IGJvb2xlYW4gfTtcblxuICAgIGlmICghdmFsdWUpIHJldHVybjtcblxuICAgIGlmIChjaGVja2VkKSB7XG4gICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuZGVsZXRlKHZhbHVlKTtcbiAgICB9XG5cbiAgICB0aGlzLiNzYXZlRXZlbnRGaWx0ZXJzKCk7XG4gICAgdGhpcy4jZmlsdGVyRXZlbnRzKHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKTtcbiAgfTtcblxuICAjaGFuZGxlRWxlbWVudEZpbHRlckNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB7IHZhbHVlLCBjaGVja2VkIH0gPSBldmVudCBhcyBFdmVudCAmIHsgdmFsdWU6IHN0cmluZzsgY2hlY2tlZDogYm9vbGVhbiB9O1xuXG4gICAgaWYgKCF2YWx1ZSkgcmV0dXJuO1xuXG4gICAgaWYgKGNoZWNrZWQpIHtcbiAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmRlbGV0ZSh2YWx1ZSk7XG4gICAgfVxuXG4gICAgdGhpcy4jc2F2ZUV2ZW50RmlsdGVycygpO1xuICAgIHRoaXMuI2ZpbHRlckV2ZW50cyh0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSk7XG4gIH07XG5cbiAgI2xvYWRFdmVudEZpbHRlcnNGcm9tU3RvcmFnZSgpOiB7IGV2ZW50VHlwZXM6IFNldDxzdHJpbmc+IHwgbnVsbDsgZWxlbWVudHM6IFNldDxzdHJpbmc+IHwgbnVsbCB9IHtcbiAgICBjb25zdCBwcmVmZXJlbmNlczogeyBldmVudFR5cGVzOiBTZXQ8c3RyaW5nPiB8IG51bGw7IGVsZW1lbnRzOiBTZXQ8c3RyaW5nPiB8IG51bGwgfSA9IHtcbiAgICAgIGV2ZW50VHlwZXM6IG51bGwsXG4gICAgICBlbGVtZW50czogbnVsbFxuICAgIH07XG5cbiAgICB0cnkge1xuICAgICAgY29uc3Qgc2F2ZWRFdmVudFR5cGVzID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1ldmVudC10eXBlLWZpbHRlcnMnKTtcbiAgICAgIGlmIChzYXZlZEV2ZW50VHlwZXMpIHtcbiAgICAgICAgcHJlZmVyZW5jZXMuZXZlbnRUeXBlcyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZEV2ZW50VHlwZXMpKTtcbiAgICAgIH1cblxuICAgICAgY29uc3Qgc2F2ZWRFbGVtZW50cyA9IGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtZWxlbWVudC1maWx0ZXJzJyk7XG4gICAgICBpZiAoc2F2ZWRFbGVtZW50cykge1xuICAgICAgICBwcmVmZXJlbmNlcy5lbGVtZW50cyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZEVsZW1lbnRzKSk7XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgY29uc29sZS5kZWJ1ZygnW2NlbS1zZXJ2ZS1jaHJvbWVdIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSBmb3IgZXZlbnQgZmlsdGVycycpO1xuICAgIH1cblxuICAgIHJldHVybiBwcmVmZXJlbmNlcztcbiAgfVxuXG4gICNzYXZlRXZlbnRGaWx0ZXJzKCkge1xuICAgIHRyeSB7XG4gICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLWV2ZW50LXR5cGUtZmlsdGVycycsXG4gICAgICAgIEpTT04uc3RyaW5naWZ5KFsuLi50aGlzLiNldmVudFR5cGVGaWx0ZXJzXSkpO1xuICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oJ2NlbS1zZXJ2ZS1lbGVtZW50LWZpbHRlcnMnLFxuICAgICAgICBKU09OLnN0cmluZ2lmeShbLi4udGhpcy4jZWxlbWVudEZpbHRlcnNdKSk7XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgLy8gbG9jYWxTdG9yYWdlIHVuYXZhaWxhYmxlIChwcml2YXRlIG1vZGUpLCBzaWxlbnRseSBjb250aW51ZVxuICAgIH1cbiAgfVxuXG4gICNjbGVhckV2ZW50cygpIHtcbiAgICB0aGlzLiNjYXB0dXJlZEV2ZW50cyA9IFtdO1xuICAgIHRoaXMuI3NlbGVjdGVkRXZlbnRJZCA9IG51bGw7XG4gICAgaWYgKHRoaXMuI2V2ZW50TGlzdCkge1xuICAgICAgdGhpcy4jZXZlbnRMaXN0LnJlcGxhY2VDaGlsZHJlbigpO1xuICAgIH1cbiAgICBpZiAodGhpcy4jZXZlbnREZXRhaWxIZWFkZXIpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyLmlubmVySFRNTCA9ICcnO1xuICAgIH1cbiAgICBpZiAodGhpcy4jZXZlbnREZXRhaWxCb2R5KSB7XG4gICAgICB0aGlzLiNldmVudERldGFpbEJvZHkuaW5uZXJIVE1MID0gJyc7XG4gICAgfVxuICB9XG5cbiAgYXN5bmMgI2NvcHlFdmVudHMoKSB7XG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIGNvbnN0IHZpc2libGVFdmVudHMgPSBBcnJheS5mcm9tKHRoaXMuI2V2ZW50TGlzdC5jaGlsZHJlbilcbiAgICAgIC5maWx0ZXIoZW50cnkgPT4gIShlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuKVxuICAgICAgLm1hcChlbnRyeSA9PiB7XG4gICAgICAgIGNvbnN0IGlkID0gKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LmV2ZW50SWQhO1xuICAgICAgICByZXR1cm4gdGhpcy4jZ2V0RXZlbnRSZWNvcmRCeUlkKGlkKTtcbiAgICAgIH0pXG4gICAgICAuZmlsdGVyKChldmVudCk6IGV2ZW50IGlzIEV2ZW50UmVjb3JkID0+ICEhZXZlbnQpXG4gICAgICAubWFwKGV2ZW50ID0+IHtcbiAgICAgICAgY29uc3QgdGltZSA9IGV2ZW50LnRpbWVzdGFtcC50b0xvY2FsZVRpbWVTdHJpbmcoKTtcbiAgICAgICAgY29uc3QgZWxlbWVudCA9IGV2ZW50LmVsZW1lbnRJZCA/IGAjJHtldmVudC5lbGVtZW50SWR9YCA6IGV2ZW50LnRhZ05hbWU7XG4gICAgICAgIGNvbnN0IHByb3BzID0gZXZlbnQuY3VzdG9tUHJvcGVydGllcyAmJiBPYmplY3Qua2V5cyhldmVudC5jdXN0b21Qcm9wZXJ0aWVzKS5sZW5ndGggPiAwXG4gICAgICAgICAgPyBgIC0gJHtKU09OLnN0cmluZ2lmeShldmVudC5jdXN0b21Qcm9wZXJ0aWVzKX1gXG4gICAgICAgICAgOiAnJztcbiAgICAgICAgcmV0dXJuIGBbJHt0aW1lfV0gPCR7ZXZlbnQudGFnTmFtZX0+ICR7ZWxlbWVudH0gXFx1MjE5MiAke2V2ZW50LmV2ZW50TmFtZX0ke3Byb3BzfWA7XG4gICAgICB9KVxuICAgICAgLmpvaW4oJ1xcbicpO1xuXG4gICAgaWYgKCF2aXNpYmxlRXZlbnRzKSByZXR1cm47XG5cbiAgICB0cnkge1xuICAgICAgYXdhaXQgbmF2aWdhdG9yLmNsaXBib2FyZC53cml0ZVRleHQodmlzaWJsZUV2ZW50cyk7XG4gICAgICBjb25zdCBidG4gPSB0aGlzLiMkKCdjb3B5LWV2ZW50cycpO1xuICAgICAgaWYgKGJ0bikge1xuICAgICAgICBjb25zdCB0ZXh0Tm9kZSA9IEFycmF5LmZyb20oYnRuLmNoaWxkTm9kZXMpLmZpbmQoXG4gICAgICAgICAgbiA9PiBuLm5vZGVUeXBlID09PSBOb2RlLlRFWFRfTk9ERSAmJiAobi50ZXh0Q29udGVudD8udHJpbSgpLmxlbmd0aCA/PyAwKSA+IDBcbiAgICAgICAgKTtcbiAgICAgICAgaWYgKHRleHROb2RlKSB7XG4gICAgICAgICAgY29uc3Qgb3JpZ2luYWwgPSB0ZXh0Tm9kZS50ZXh0Q29udGVudDtcbiAgICAgICAgICB0ZXh0Tm9kZS50ZXh0Q29udGVudCA9ICdDb3BpZWQhJztcblxuICAgICAgICAgIGlmICh0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgICAgaWYgKHRoaXMuaXNDb25uZWN0ZWQgJiYgdGV4dE5vZGUucGFyZW50Tm9kZSkge1xuICAgICAgICAgICAgICB0ZXh0Tm9kZS50ZXh0Q29udGVudCA9IG9yaWdpbmFsO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgICAgICAgfSwgMjAwMCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9IGNhdGNoIChlcnIpIHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY29weSBldmVudHM6JywgZXJyKTtcbiAgICB9XG4gIH1cblxuICAjc2V0dXBFdmVudExpc3RlbmVycygpIHtcbiAgICB0aGlzLiNldmVudExpc3QgPSB0aGlzLiMkKCdldmVudC1saXN0Jyk7XG4gICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIgPSB0aGlzLiMkKCdldmVudC1kZXRhaWwtaGVhZGVyJyk7XG4gICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5ID0gdGhpcy4jJCgnZXZlbnQtZGV0YWlsLWJvZHknKTtcblxuICAgIGlmICh0aGlzLiNldmVudExpc3QpIHtcbiAgICAgIHRoaXMuI2V2ZW50TGlzdC5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsIChlOiBFdmVudCkgPT4ge1xuICAgICAgICBjb25zdCBsaXN0SXRlbSA9IChlLnRhcmdldCBhcyBFbGVtZW50KS5jbG9zZXN0KCcuZXZlbnQtbGlzdC1pdGVtJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICAgIGlmIChsaXN0SXRlbSkge1xuICAgICAgICAgIGNvbnN0IGV2ZW50SWQgPSBsaXN0SXRlbS5kYXRhc2V0LmV2ZW50SWQ7XG4gICAgICAgICAgaWYgKGV2ZW50SWQpIHtcbiAgICAgICAgICAgIHRoaXMuI3NlbGVjdEV2ZW50KGV2ZW50SWQpO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfSk7XG4gICAgfVxuXG4gICAgY29uc3QgZXZlbnRzRmlsdGVyID0gdGhpcy4jJCgnZXZlbnRzLWZpbHRlcicpO1xuICAgIGlmIChldmVudHNGaWx0ZXIpIHtcbiAgICAgIGV2ZW50c0ZpbHRlci5hZGRFdmVudExpc3RlbmVyKCdpbnB1dCcsIChlOiBFdmVudCkgPT4ge1xuICAgICAgICBjb25zdCB7IHZhbHVlID0gJycgfSA9IGUudGFyZ2V0IGFzIEhUTUxJbnB1dEVsZW1lbnQ7XG4gICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNldmVudHNGaWx0ZXJEZWJvdW5jZVRpbWVyISk7XG4gICAgICAgIHRoaXMuI2V2ZW50c0ZpbHRlckRlYm91bmNlVGltZXIgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICB0aGlzLiNmaWx0ZXJFdmVudHModmFsdWUpO1xuICAgICAgICB9LCAzMDApO1xuICAgICAgfSk7XG4gICAgfVxuXG4gICAgY29uc3QgZXZlbnRUeXBlRmlsdGVyID0gdGhpcy4jJCgnZXZlbnQtdHlwZS1maWx0ZXInKTtcbiAgICBpZiAoZXZlbnRUeXBlRmlsdGVyKSB7XG4gICAgICBldmVudFR5cGVGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlRXZlbnRUeXBlRmlsdGVyQ2hhbmdlIGFzIEV2ZW50TGlzdGVuZXIpO1xuICAgIH1cblxuICAgIGNvbnN0IGVsZW1lbnRGaWx0ZXIgPSB0aGlzLiMkKCdlbGVtZW50LWZpbHRlcicpO1xuICAgIGlmIChlbGVtZW50RmlsdGVyKSB7XG4gICAgICBlbGVtZW50RmlsdGVyLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZUVsZW1lbnRGaWx0ZXJDaGFuZ2UgYXMgRXZlbnRMaXN0ZW5lcik7XG4gICAgfVxuXG4gICAgdGhpcy4jJCgnY2xlYXItZXZlbnRzJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgdGhpcy4jY2xlYXJFdmVudHMoKTtcbiAgICB9KTtcblxuICAgIHRoaXMuIyQoJ2NvcHktZXZlbnRzJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgdGhpcy4jY29weUV2ZW50cygpO1xuICAgIH0pO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcblxuICAgIC8vIENsZWFuIHVwIGtub2IgbGlzdGVuZXJzXG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOmF0dHJpYnV0ZS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuXG4gICAgLy8gQ2xlYW4gdXAgdHJlZSBzdGF0ZSBsaXN0ZW5lcnNcbiAgICBpZiAodGhpcy4jaGFuZGxlVHJlZUV4cGFuZCkge1xuICAgICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdleHBhbmQnLCB0aGlzLiNoYW5kbGVUcmVlRXhwYW5kKTtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSkge1xuICAgICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdjb2xsYXBzZScsIHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSk7XG4gICAgfVxuICAgIGlmICh0aGlzLiNoYW5kbGVUcmVlU2VsZWN0KSB7XG4gICAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QpO1xuICAgIH1cblxuICAgIC8vIENsZWFuIHVwIHdpbmRvdyBsaXN0ZW5lclxuICAgIGlmICh0aGlzLiNoYW5kbGVMb2dzRXZlbnQpIHtcbiAgICAgIHdpbmRvdy5yZW1vdmVFdmVudExpc3RlbmVyKCdjZW06bG9ncycsIHRoaXMuI2hhbmRsZUxvZ3NFdmVudCk7XG4gICAgfVxuXG4gICAgLy8gQ2xlYXIgcGVuZGluZyBmZWVkYmFjayB0aW1lb3V0c1xuICAgIGlmICh0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgIHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICB0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgIH1cbiAgICBpZiAodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgfVxuXG4gICAgLy8gRGlzY29ubmVjdCBtdXRhdGlvbiBvYnNlcnZlclxuICAgIHRoaXMuI29ic2VydmVyLmRpc2Nvbm5lY3QoKTtcblxuICAgIC8vIENsb3NlIFdlYlNvY2tldCBjb25uZWN0aW9uXG4gICAgaWYgKHRoaXMuI3dzQ2xpZW50KSB7XG4gICAgICB0aGlzLiN3c0NsaWVudC5kZXN0cm95KCk7XG4gICAgfVxuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1zZXJ2ZS1jaHJvbWUnOiBDZW1TZXJ2ZUNocm9tZTtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbiAgaGVpZ2h0OiAxMDB2aDtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgLS1jZW0tcGYtdjYtYy1tYXN0aGVhZF9fbG9nby0tV2lkdGg6IG1heC1jb250ZW50O1xcXFxuICAtLWNlbS1wZi12Ni1jLW1hc3RoZWFkX190b2dnbGUtLVNpemU6IDFyZW07XFxcXG59XFxcXG5cXFxcbltoaWRkZW5dIHtcXFxcbiAgZGlzcGxheTogbm9uZSAhaW1wb3J0YW50O1xcXFxufVxcXFxuXFxcXG4vKiBNYXN0aGVhZCBsb2dvIHN0eWxlcyAqL1xcXFxuLm1hc3RoZWFkLWxvZ28ge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG4gIGNvbG9yOiBpbmhlcml0O1xcXFxuICBtYXgtaGVpZ2h0OiB2YXIoLS1jZW0tcGYtdjYtYy1tYXN0aGVhZF9fbG9nby0tTWF4SGVpZ2h0KTtcXFxcbiAgZ2FwOiA0cHg7XFxcXG4gIFxcXFx1MDAyNiBpbWcge1xcXFxuICAgIGRpc3BsYXk6IGJsb2NrO1xcXFxuICAgIG1heC1oZWlnaHQ6IHZhcigtLWNlbS1wZi12Ni1jLW1hc3RoZWFkX19sb2dvLS1NYXhIZWlnaHQpO1xcXFxuICAgIHdpZHRoOiBhdXRvO1xcXFxuICB9XFxcXG4gIFxcXFx1MDAyNiA6OnNsb3R0ZWQoW3Nsb3Q9XFxcXFxcXCJ0aXRsZVxcXFxcXFwiXSkge1xcXFxuICAgIG1hcmdpbjogMDtcXFxcbiAgICBmb250LXNpemU6IDEuMTI1cmVtO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICB9XFxcXG4gIFxcXFx1MDAyNiBoMSB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICAgIGZvbnQtc2l6ZTogMThweDtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBUb29sYmFyIGdyb3VwIGFsaWdubWVudCAqL1xcXFxuY2VtLXBmLXY2LXRvb2xiYXItZ3JvdXBbdmFyaWFudD1cXFxcXFxcImFjdGlvbi1ncm91cFxcXFxcXFwiXSB7XFxcXG4gIG1hcmdpbi1pbmxpbmUtc3RhcnQ6IGF1dG87XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG59XFxcXG5cXFxcbi5kZWJ1Zy1wYW5lbCB7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWRlZmF1bHQpO1xcXFxuICBib3JkZXI6IDFweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxuICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gIHBhZGRpbmc6IDEuNXJlbTtcXFxcbiAgbWF4LXdpZHRoOiA2MDBweDtcXFxcbiAgd2lkdGg6IDkwJTtcXFxcbiAgbWF4LWhlaWdodDogODB2aDtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcblxcXFxuICBoMiB7XFxcXG4gICAgbWFyZ2luOiAwIDAgMXJlbSAwO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICBmb250LXNpemU6IDEuMTI1cmVtO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIH1cXFxcblxcXFxuICBkbCB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICB9XFxcXG5cXFxcbiAgZHQge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgbWFyZ2luLXRvcDogMC41cmVtO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG4gIH1cXFxcblxcXFxuICBkZCB7XFxcXG4gICAgbWFyZ2luOiAwIDAgMC41cmVtIDA7XFxcXG4gICAgZm9udC1mYW1pbHk6IHVpLW1vbm9zcGFjZSwgJ0Nhc2NhZGlhIENvZGUnLCAnU291cmNlIENvZGUgUHJvJywgTWVubG8sIENvbnNvbGFzLCAnRGVqYVZ1IFNhbnMgTW9ubycsIG1vbm9zcGFjZTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgfVxcXFxuXFxcXG4gIGRldGFpbHMge1xcXFxuICAgIG1hcmdpbi10b3A6IDFyZW07XFxcXG5cXFxcbiAgICBzdW1tYXJ5IHtcXFxcbiAgICAgIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgICBmb250LXdlaWdodDogNTAwO1xcXFxuICAgICAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICAgICAgZ2FwOiAwLjVyZW07XFxcXG4gICAgICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcblxcXFxuICAgICAgXFxcXHUwMDI2Ojotd2Via2l0LWRldGFpbHMtbWFya2VyIHtcXFxcbiAgICAgICAgZGlzcGxheTogbm9uZTtcXFxcbiAgICAgIH1cXFxcblxcXFxuICAgICAgXFxcXHUwMDI2OjpiZWZvcmUge1xcXFxuICAgICAgICBjb250ZW50OiAnXFxcXFxcXFwyNUI4JztcXFxcbiAgICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICAgICAgICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gMTAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgICB9XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2W29wZW5dIHN1bW1hcnk6OmJlZm9yZSB7XFxcXG4gICAgICB0cmFuc2Zvcm06IHJvdGF0ZSg5MGRlZyk7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgcHJlIHtcXFxcbiAgICAgIG1hcmdpbi10b3A6IDAuNXJlbTtcXFxcbiAgICAgIG1hcmdpbi1sZWZ0OiAxLjVyZW07XFxcXG4gICAgICBwYWRkaW5nOiAwLjVyZW07XFxcXG4gICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWRlZmF1bHQpO1xcXFxuICAgICAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICAgICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICAgIG92ZXJmbG93LXg6IGF1dG87XFxcXG4gICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgLmJ1dHRvbi1jb250YWluZXIge1xcXFxuICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgZ2FwOiAwLjVyZW07XFxcXG4gICAgbWFyZ2luLXRvcDogMXJlbTtcXFxcbiAgfVxcXFxuXFxcXG4gIHAge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gIH1cXFxcblxcXFxuICBidXR0b24ge1xcXFxuICAgIG1hcmdpbi10b3A6IDFyZW07XFxcXG4gICAgcGFkZGluZzogMC41cmVtIDFyZW07XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQpO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1icmFuZCk7XFxcXG4gICAgYm9yZGVyOiBub25lO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA0MDA7XFxcXG4gICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICAgIHRyYW5zaXRpb246IGFsbCAyMDBtcyBjdWJpYy1iZXppZXIoMC42NDUsIDAuMDQ1LCAwLjM1NSwgMSk7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWhvdmVyKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogQ29udGVudCBhcmVhIHBhZGRpbmcgZm9yIGRlbW8gKi9cXFxcbmNlbS1wZi12Ni1wYWdlLW1haW4ge1xcXFxuICBtaW4taGVpZ2h0OiBjYWxjKDEwMGR2aCAtIDcycHggLSB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0taW5zZXQtLXBhZ2UtY2hyb21lKSk7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIFxcXFx1MDAyNiBcXFxcdTAwM2UgOjpzbG90dGVkKDpub3QoW3Nsb3Q9a25vYnNdKSkge1xcXFxuICAgIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gICAgZmxleDogMTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG5jZW0tZHJhd2VyIHtcXFxcbiAgY2VtLXBmLXY2LXRhYnMge1xcXFxuICAgIGNlbS1wZi12Ni10YWIge1xcXFxuICAgICAgcGFkZGluZy1ibG9jay1lbmQ6IDA7XFxcXG4gICAgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIEVsZW1lbnQgZGVzY3JpcHRpb25zIGluIGxpc3RpbmcgKi9cXFxcbi5lbGVtZW50LXN1bW1hcnkge1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxufVxcXFxuXFxcXG4uZWxlbWVudC1kZXNjcmlwdGlvbiB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG59XFxcXG5cXFxcbi8qIENhcmQgZm9vdGVyIGRlbW8gbmF2aWdhdGlvbiAqL1xcXFxuLmNhcmQtZGVtb3Mge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LXdyYXA6IHdyYXA7XFxcXG4gIGdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWdhcC0tYWN0aW9uLXRvLWFjdGlvbi0tZGVmYXVsdCk7XFxcXG4gIHBhZGRpbmc6IDA7XFxcXG4gIG1hcmdpbjogMDtcXFxcbn1cXFxcblxcXFxuLnBhY2thZ2UtbmFtZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxufVxcXFxuXFxcXG4vKiBLbm9icyBjb250YWluZXIgLSBmaWxscyB0YWIgcGFuZWwgaGVpZ2h0ICovXFxcXG4ja25vYnMtY29udGFpbmVyIHtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBcXFxcdTAwMjYgOjpzbG90dGVkKFtzbG90PVxcXFxcXFwia25vYnNcXFxcXFxcIl0pIHtcXFxcbiAgICBmbGV4OiAxO1xcXFxuICAgIG1pbi1oZWlnaHQ6IDA7XFxcXG4gICAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ua25vYnMtZW1wdHkge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgdGV4dC1hbGlnbjogY2VudGVyO1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLWxnKTtcXFxcblxcXFxuICBjb2RlIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gICAgcGFkZGluZzogMnB4IDZweDtcXFxcbiAgICBib3JkZXItcmFkaXVzOiAzcHg7XFxcXG4gICAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5pbnN0YW5jZS10YWcge1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5pbnN0YW5jZS1sYWJlbCB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLmtub2ItZ3JvdXAge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLWxnKTtcXFxcblxcXFxuICBcXFxcdTAwMjY6bGFzdC1jaGlsZCB7XFxcXG4gICAgbWFyZ2luLWJvdHRvbTogMDtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBQYXR0ZXJuRmx5IHY2IGZvcm0gLSBob3Jpem9udGFsIGxheW91dCAqL1xcXFxuY2VtLXBmLXY2LWZvcm1baG9yaXpvbnRhbF0gY2VtLXBmLXY2LWZvcm0tZmllbGQtZ3JvdXAge1xcXFxuICBncmlkLWNvbHVtbjogc3BhbiAyXFxcXG59XFxcXG5cXFxcbi5rbm9iLWdyb3VwLXRpdGxlIHtcXFxcbiAgZ3JpZC1jb2x1bW46IDEgLyAtMTtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgMDtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLWJhc2UpO1xcXFxuICBmb250LXdlaWdodDogNjAwO1xcXFxuICBib3JkZXItYm90dG9tOiAxcHggc29saWQgdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLWNvbG9yKTtcXFxcbiAgcGFkZGluZy1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxufVxcXFxuXFxcXG4ua25vYi1jb250cm9sIHtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWxhYmVsIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKTtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxufVxcXFxuXFxcXG4ua25vYi1uYW1lIHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxuICBmb250LXdlaWdodDogNTAwO1xcXFxufVxcXFxuXFxcXG4ua25vYi1kZXNjcmlwdGlvbiB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgbGluZS1oZWlnaHQ6IDEuNTtcXFxcblxcXFxuICBwIHtcXFxcbiAgICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIDA7XFxcXG4gIH1cXFxcblxcXFxuICBjb2RlIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogM3B4O1xcXFxuICAgIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgfVxcXFxuXFxcXG4gIGEge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICAgIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICB0ZXh0LWRlY29yYXRpb246IHVuZGVybGluZTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBzdHJvbmcge1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIH1cXFxcblxcXFxuICB1bCwgb2wge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgICBwYWRkaW5nLWxlZnQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbGcpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbmZvb3Rlci5wZi1tLXN0aWNreS1ib3R0b20ge1xcXFxuICB2aWV3LXRyYW5zaXRpb24tbmFtZTogZGV2LXNlcnZlci1mb290ZXI7XFxcXG4gIHBvc2l0aW9uOiBzdGlja3k7XFxcXG4gIGJvdHRvbTogMDtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlci10b3A6IDFweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxuICB6LWluZGV4OiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlLS1zZWN0aW9uLS1tLXN0aWNreS1ib3R0b20tLVpJbmRleCwgdmFyKC0tcGYtdC0tZ2xvYmFsLS16LWluZGV4LS1tZCkpO1xcXFxuICBib3gtc2hhZG93OiB2YXIoLS1jZW0tcGYtdjYtYy1wYWdlLS1zZWN0aW9uLS1tLXN0aWNreS1ib3R0b20tLUJveFNoYWRvdywgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3gtc2hhZG93LS1zbS0tdG9wKSk7XFxcXG59XFxcXG5cXFxcbi5mb290ZXItZGVzY3JpcHRpb24ge1xcXFxuICBwYWRkaW5nOiAxLjVyZW07XFxcXG5cXFxcbiAgXFxcXHUwMDI2LmVtcHR5IHtcXFxcbiAgICBkaXNwbGF5OiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbmZvb3RlciA6OnNsb3R0ZWQoW3Nsb3Q9XFxcXFxcXCJkZXNjcmlwdGlvblxcXFxcXFwiXSkge1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBsaW5lLWhlaWdodDogMS42O1xcXFxuICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWhvdmVyKTtcXFxcbiAgICBwYWRkaW5nOiAycHggNnB4O1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDNweDtcXFxcbiAgICBmb250LWZhbWlseTogdWktbW9ub3NwYWNlLCAnQ2FzY2FkaWEgQ29kZScsICdTb3VyY2UgQ29kZSBQcm8nLCBNZW5sbywgQ29uc29sYXMsICdEZWphVnUgU2FucyBNb25vJywgbW9ub3NwYWNlO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmxvZ3Mtd3JhcHBlciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxuI2xvZy1jb250YWluZXIge1xcXFxuICBmbGV4LWdyb3c6IDE7XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG59XFxcXG5cXFxcbi5sb2ctZW50cnkge1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgYWxpZ24taXRlbXM6IGJhc2VsaW5lO1xcXFxuICBjZW0tcGYtdjYtbGFiZWwge1xcXFxuICAgIGZsZXgtc2hyaW5rOiAwO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5sb2ctdGltZSxcXFxcbi5sb2ctbWVzc2FnZSB7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4ubG9nLXRpbWUge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZsZXgtc2hyaW5rOiAwO1xcXFxuICBmb250LXNpemU6IDExcHg7XFxcXG59XFxcXG5cXFxcbi5sb2ctbWVzc2FnZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxuICB3b3JkLWJyZWFrOiBicmVhay13b3JkO1xcXFxufVxcXFxuXFxcXG4vKiBOYXZpZ2F0aW9uIGNvbnRlbnQgKGxpZ2h0IERPTSBzbG90dGVkIGNvbnRlbnQgZm9yIGNlbS1wZi12Ni1wYWdlLXNpZGViYXIpICovXFxcXG4ubmF2LXBhY2thZ2Uge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcblxcXFxuICBcXFxcdTAwMjYgXFxcXHUwMDNlIHN1bW1hcnkge1xcXFxuICAgIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgICBwYWRkaW5nOiAwLjVyZW0gMXJlbTtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWRlZmF1bHQpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gICAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAyMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgIG1hcmdpbi1ib3R0b206IDAuNXJlbTtcXFxcbiAgICBkaXNwbGF5OiBmbGV4O1xcXFxuICAgIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gICAgZ2FwOiAwLjVyZW07XFxcXG4gICAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1ob3Zlcik7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2Ojotd2Via2l0LWRldGFpbHMtbWFya2VyIHtcXFxcbiAgICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2OjpiZWZvcmUge1xcXFxuICAgICAgY29udGVudDogJ1xcXFxcXFxcMjVCOCc7XFxcXG4gICAgICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XFxcXG4gICAgICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gMTAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNltvcGVuXSBcXFxcdTAwM2Ugc3VtbWFyeTo6YmVmb3JlIHtcXFxcbiAgICB0cmFuc2Zvcm06IHJvdGF0ZSg5MGRlZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLm5hdi1lbGVtZW50IHtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG4gIG1hcmdpbi1pbmxpbmUtc3RhcnQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuXFxcXG4gIHN1bW1hcnkge1xcXFxuICAgIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgICBwYWRkaW5nOiAwLjVyZW0gMXJlbTtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWRlZmF1bHQpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgZm9udC1mYW1pbHk6IHVpLW1vbm9zcGFjZSwgJ0Nhc2NhZGlhIENvZGUnLCAnU291cmNlIENvZGUgUHJvJywgTWVubG8sIENvbnNvbGFzLCAnRGVqYVZ1IFNhbnMgTW9ubycsIG1vbm9zcGFjZTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gICAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAyMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgICBnYXA6IDAuNXJlbTtcXFxcbiAgICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWhvdmVyKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6Oi13ZWJraXQtZGV0YWlscy1tYXJrZXIge1xcXFxuICAgICAgZGlzcGxheTogbm9uZTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgICBjb250ZW50OiAnXFxcXFxcXFwyNUI4JztcXFxcbiAgICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgICAgIHRyYW5zaXRpb246IHRyYW5zZm9ybSAxMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2W29wZW5dIHN1bW1hcnk6OmJlZm9yZSB7XFxcXG4gICAgdHJhbnNmb3JtOiByb3RhdGUoOTBkZWcpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5uYXYtZWxlbWVudC10aXRsZSB7XFxcXG4gIHVzZXItc2VsZWN0OiBub25lO1xcXFxufVxcXFxuXFxcXG4ubmF2LWRlbW8tbGlzdCB7XFxcXG4gIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gIHBhZGRpbmc6IDA7XFxcXG4gIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMCAwIDA7XFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cyk7XFxcXG59XFxcXG5cXFxcbi5uYXYtZGVtby1saW5rIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLXN0YXJ0OiBjYWxjKHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpICogMik7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXJhZGl1cyk7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIHRyYW5zaXRpb246IGJhY2tncm91bmQgMC4ycywgY29sb3IgMC4ycztcXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLWJyYW5kKTtcXFxcblxcXFxuICAgIC5uYXYtcGFja2FnZS1uYW1lIHtcXFxcbiAgICAgIGNvbG9yOiByZ2JhKDI1NSwgMjU1LCAyNTUsIDAuOCk7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2W2FyaWEtY3VycmVudD1cXFxcXFxcInBhZ2VcXFxcXFxcIl0ge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLWJyYW5kKTtcXFxcblxcXFxuICAgIC5uYXYtcGFja2FnZS1uYW1lIHtcXFxcbiAgICAgIGNvbG9yOiByZ2JhKDI1NSwgMjU1LCAyNTUsIDAuOCk7XFxcXG4gICAgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5uYXYtcGFja2FnZS1uYW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4vKiBJbmZvIGJ1dHRvbiBwb3BvdmVyIHRyaWdnZXJzIGluIGtub2JzIC0gb3ZlcnJpZGUgcGxhaW4gYnV0dG9uIHBhZGRpbmcgKi9cXFxcbmNlbS1wZi12Ni1wb3BvdmVyIGNlbS1wZi12Ni1idXR0b25bdmFyaWFudD1cXFxcXFxcInBsYWluXFxcXFxcXCJdIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy1idXR0b24tLW0tcGxhaW4tLVBhZGRpbmdJbmxpbmVFbmQ6IDA7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1tLXBsYWluLS1QYWRkaW5nSW5saW5lU3RhcnQ6IDA7XFxcXG4gIC0tY2VtLXBmLXY2LWMtYnV0dG9uLS1NaW5XaWR0aDogYXV0bztcXFxcbn1cXFxcblxcXFxuLyogS25vYiBkZXNjcmlwdGlvbiBjb250ZW50IChzbG90dGVkIGluIGZvcm0gZ3JvdXAgaGVscGVyKSAqL1xcXFxuY2VtLXBmLXY2LWZvcm0tZ3JvdXAgW3Nsb3Q9XFxcXFxcXCJoZWxwZXJcXFxcXFxcIl0ge1xcXFxuICBwIHtcXFxcbiAgICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIDA7XFxcXG4gIH1cXFxcblxcXFxuICBjb2RlIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogM3B4O1xcXFxuICAgIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgfVxcXFxuXFxcXG4gIGEge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICAgIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICB0ZXh0LWRlY29yYXRpb246IHVuZGVybGluZTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBzdHJvbmcge1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIH1cXFxcblxcXFxuICB1bCwgb2wge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgICBwYWRkaW5nLWxlZnQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbGcpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIFN5bnRheCBoaWdobGlnaHRpbmcgKGNocm9tYSAtIHRoZW1hYmxlIHZpYSBDU1MgY3VzdG9tIHByb3BlcnRpZXMpICovXFxcXG5jZW0tcGYtdjYtZm9ybS1ncm91cCBbc2xvdD1cXFxcXFxcImhlbHBlclxcXFxcXFwiXSB7XFxcXG4gIC5jaHJvbWEge1xcXFxuICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItcmFkaXVzKTtcXFxcbiAgICBvdmVyZmxvdy14OiBhdXRvO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2IC5sbnRkIHsgdmVydGljYWwtYWxpZ246IHRvcDsgcGFkZGluZzogMDsgbWFyZ2luOiAwOyBib3JkZXI6IDA7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmxudGFibGUgeyBib3JkZXItc3BhY2luZzogMDsgcGFkZGluZzogMDsgbWFyZ2luOiAwOyBib3JkZXI6IDA7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmhsIHsgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWhpZ2hsaWdodCkgfVxcXFxuICAgIFxcXFx1MDAyNiAubG50LFxcXFxuICAgIFxcXFx1MDAyNiAubG4ge1xcXFxuICAgICAgd2hpdGUtc3BhY2U6IHByZTtcXFxcbiAgICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuICAgICAgbWFyZ2luLXJpZ2h0OiAwLjRlbTtcXFxcbiAgICAgIHBhZGRpbmc6IDAgMC40ZW0gMCAwLjRlbTtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5saW5lIHsgZGlzcGxheTogZmxleDsgfVxcXFxuXFxcXG4gICAgLyogS2V5d29yZHMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmssXFxcXG4gICAgXFxcXHUwMDI2IC5rYyxcXFxcbiAgICBcXFxcdTAwMjYgLmtkLFxcXFxuICAgIFxcXFx1MDAyNiAua24sXFxcXG4gICAgXFxcXHUwMDI2IC5rcCxcXFxcbiAgICBcXFxcdTAwMjYgLmtyIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgta2V5d29yZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5rdCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtdHlwZSk7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG5cXFxcbiAgICAvKiBOYW1lcyAqL1xcXFxuICAgIFxcXFx1MDAyNiAubmEsXFxcXG4gICAgXFxcXHUwMDI2IC5uYixcXFxcbiAgICBcXFxcdTAwMjYgLm5vLFxcXFxuICAgIFxcXFx1MDAyNiAubnYsXFxcXG4gICAgXFxcXHUwMDI2IC52YyxcXFxcbiAgICBcXFxcdTAwMjYgLnZnLFxcXFxuICAgIFxcXFx1MDAyNiAudmkge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1uYW1lKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5icCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAubmMgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWNsYXNzKTsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLm5kIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1kZWNvcmF0b3IpOyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuICAgIFxcXFx1MDAyNiAubmksXFxcXG4gICAgXFxcXHUwMDI2IC5zcyB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXNwZWNpYWwpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLm5lLFxcXFxuICAgIFxcXFx1MDAyNiAubmwge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1rZXl3b3JkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLm5mIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1mdW5jdGlvbik7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5ubiB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAubnQgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXRhZykgfVxcXFxuXFxcXG4gICAgLyogU3RyaW5ncyAqL1xcXFxuICAgIFxcXFx1MDAyNiAucyxcXFxcbiAgICBcXFxcdTAwMjYgLnNhLFxcXFxuICAgIFxcXFx1MDAyNiAuc2IsXFxcXG4gICAgXFxcXHUwMDI2IC5zYyxcXFxcbiAgICBcXFxcdTAwMjYgLmRsLFxcXFxuICAgIFxcXFx1MDAyNiAuc2QsXFxcXG4gICAgXFxcXHUwMDI2IC5zMixcXFxcbiAgICBcXFxcdTAwMjYgLnNlLFxcXFxuICAgIFxcXFx1MDAyNiAuc2gsXFxcXG4gICAgXFxcXHUwMDI2IC5zaSxcXFxcbiAgICBcXFxcdTAwMjYgLnN4LFxcXFxuICAgIFxcXFx1MDAyNiAuczEge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1zdHJpbmcpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLnNyIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC10YWcpIH1cXFxcblxcXFxuICAgIC8qIE51bWJlcnMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLm0sXFxcXG4gICAgXFxcXHUwMDI2IC5tYixcXFxcbiAgICBcXFxcdTAwMjYgLm1mLFxcXFxuICAgIFxcXFx1MDAyNiAubWgsXFxcXG4gICAgXFxcXHUwMDI2IC5taSxcXFxcbiAgICBcXFxcdTAwMjYgLmlsLFxcXFxuICAgIFxcXFx1MDAyNiAubW8ge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1udW1iZXIpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIE9wZXJhdG9ycyAqL1xcXFxuICAgIFxcXFx1MDAyNiAubyxcXFxcbiAgICBcXFxcdTAwMjYgLm93IHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgta2V5d29yZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBDb21tZW50cyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuYyxcXFxcbiAgICBcXFxcdTAwMjYgLmNoLFxcXFxuICAgIFxcXFx1MDAyNiAuY20sXFxcXG4gICAgXFxcXHUwMDI2IC5jMSB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gICAgICBmb250LXN0eWxlOiBpdGFsaWM7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuY3MsXFxcXG4gICAgXFxcXHUwMDI2IC5jcCxcXFxcbiAgICBcXFxcdTAwMjYgLmNwZiB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBFcnJvcnMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmVyciB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWVycm9yKTtcXFxcbiAgICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvci1iZyk7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgLyogR2VuZXJpY3MgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmdkIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZGVsZXRlZCk7XFxcXG4gICAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZGVsZXRlZC1iZyk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2UgeyBmb250LXN0eWxlOiBpdGFsaWM7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmdyIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvcikgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2ggeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdpIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtaW5zZXJ0ZWQpO1xcXFxuICAgICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWluc2VydGVkLWJnKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5nbyB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3AgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdzIHsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmd1IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5ndCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZXJyb3IpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdsIHsgdGV4dC1kZWNvcmF0aW9uOiB1bmRlcmxpbmU7IH1cXFxcbiAgICBcXFxcdTAwMjYgLncgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCkgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIEV2ZW50cyB0YWIgc3R5bGluZyAtIFByaW1hcnktZGV0YWlsIGxheW91dCAqL1xcXFxuLmV2ZW50cy13cmFwcGVyIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4jZXZlbnQtZHJhd2VyIHtcXFxcbiAgZmxleDogMTtcXFxcbiAgbWluLWhlaWdodDogMDtcXFxcbn1cXFxcblxcXFxuLyogRXZlbnQgbGlzdCAocHJpbWFyeSBwYW5lbCkgKi9cXFxcbiNldmVudC1saXN0IHtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtbGlzdC1pdGVtIHtcXFxcbiAgLyogUmVzZXQgYnV0dG9uIHN0eWxlcyAqL1xcXFxuICBhcHBlYXJhbmNlOiBub25lO1xcXFxuICBiYWNrZ3JvdW5kOiBub25lO1xcXFxuICBib3JkZXI6IG5vbmU7XFxcXG4gIGJvcmRlci1sZWZ0OiAzcHggc29saWQgdHJhbnNwYXJlbnQ7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgZm9udDogaW5oZXJpdDtcXFxcbiAgY29sb3I6IGluaGVyaXQ7XFxcXG4gIHRleHQtYWxpZ246IGluaGVyaXQ7XFxcXG4gIHdpZHRoOiAxMDAlO1xcXFxuXFxcXG4gIC8qIENvbXBvbmVudCBzdHlsZXMgKi9cXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAxMDBtcyBlYXNlLWluLW91dCwgYm9yZGVyLWNvbG9yIDEwMG1zIGVhc2UtaW4tb3V0O1xcXFxuXFxcXG4gIGNlbS1wZi12Ni1sYWJlbCB7XFxcXG4gICAgZmxleC1zaHJpbms6IDA7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWhvdmVyKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjpmb2N1cyB7XFxcXG4gICAgb3V0bGluZTogMnB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tY2xpY2tlZCk7XFxcXG4gICAgb3V0bGluZS1vZmZzZXQ6IC0ycHg7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYuc2VsZWN0ZWQge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLXNlbGVjdGVkKTtcXFxcbiAgICBib3JkZXItbGVmdC1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmV2ZW50LXRpbWUsXFxcXG4uZXZlbnQtZWxlbWVudCB7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtdGltZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZmxleC1zaHJpbms6IDA7XFxcXG4gIGZvbnQtc2l6ZTogMTFweDtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWVsZW1lbnQge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA0MDA7XFxcXG59XFxcXG5cXFxcbi8qIEV2ZW50IGRldGFpbCBwYW5lbCAqL1xcXFxuLmV2ZW50LWRldGFpbC1oZWFkZXItY29udGVudCB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBib3JkZXItYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItd2lkdGgpIHNvbGlkIHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1jb2xvcik7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtbmFtZSB7XFxcXG4gIG1hcmdpbjogMCAwIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIDA7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLWxnKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtc3VtbWFyeSB7XFxcXG4gIG1hcmdpbjogMCAwIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIDA7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgbGluZS1oZWlnaHQ6IDEuNTtcXFxcbiAgd2hpdGUtc3BhY2U6IHByZS13cmFwO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLWRlc2NyaXB0aW9uIHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxuICB3aGl0ZS1zcGFjZTogcHJlLXdyYXA7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtbWV0YSB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC10aW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtZWxlbWVudCB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1wcm9wZXJ0aWVzLWhlYWRpbmcge1xcXFxuICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1iYXNlKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtcHJvcGVydGllcyB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy1zZWNvbmRhcnkpO1xcXFxuICBib3JkZXI6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci13aWR0aCkgc29saWQgdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLWNvbG9yKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXJhZGl1cyk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgZm9udC1zaXplOiAxMnB4O1xcXFxuICBsaW5lLWhlaWdodDogMS42O1xcXFxuICBtYXJnaW46IDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1wcm9wZXJ0eS10cmVlIHtcXFxcbiAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuXFxcXG4gIFxcXFx1MDAyNi5uZXN0ZWQge1xcXFxuICAgIHBhZGRpbmctbGVmdDogMS41ZW07XFxcXG4gICAgbWFyZ2luLXRvcDogMC4yNWVtO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5wcm9wZXJ0eS1pdGVtIHtcXFxcbiAgcGFkZGluZzogMC4xMjVlbSAwO1xcXFxufVxcXFxuXFxcXG4ucHJvcGVydHkta2V5IHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG59XFxcXG5cXFxcbi5wcm9wZXJ0eS1jb2xvbiB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LXZhbHVlIHtcXFxcbiAgXFxcXHUwMDI2Lm51bGwsXFxcXG4gIFxcXFx1MDAyNi51bmRlZmluZWQge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICBmb250LXN0eWxlOiBpdGFsaWM7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYuYm9vbGVhbiB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWNvbG9yLWJvb2xlYW4pO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2Lm51bWJlciB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWNvbG9yLW51bWJlcik7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYuc3RyaW5nIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItY29sb3Itc3RyaW5nKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5hcnJheSxcXFxcbiAgXFxcXHUwMDI2Lm9iamVjdCB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jZGVidWctbW9kYWwge1xcXFxuICBjb250YWluZXItdHlwZTogaW5saW5lLXNpemU7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxNQUFNLGVBQWU7QUFDMUMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7QUFFekIsU0FBUyxpQkFBaUI7OztBQ0oxQixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSw0dHNCQUE4dXNCLENBQUM7QUFBRSxJQUFPLDJCQUFROzs7QURRcnpzQixPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87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
