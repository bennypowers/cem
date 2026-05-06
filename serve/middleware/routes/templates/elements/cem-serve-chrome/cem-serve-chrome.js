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
    __privateAdd(this, _logLevelFilters, /* @__PURE__ */ new Set(["info", "warn", "error", "debug"]));
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
                            <cem-pf-v6-menu-item variant="checkbox" value="info" checked>Info</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="warn" checked>Warnings</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="error" checked>Errors</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="debug" checked>Debug</cem-pf-v6-menu-item>
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
    if (["info", "warning", "error", "debug"].includes(cls)) {
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
    case "debug":
      return "Debug";
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
    case "debug":
      return label.setAttribute("color", "purple");
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcbmltcG9ydCB7IGlmRGVmaW5lZCB9IGZyb20gJ2xpdC9kaXJlY3RpdmVzL2lmLWRlZmluZWQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWNocm9tZS5jc3MnO1xuXG5pbXBvcnQgJy4uL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlLmpzJztcbmltcG9ydCAnLi4vY2VtLWRyYXdlci9jZW0tZHJhd2VyLmpzJztcbmltcG9ydCAnLi4vY2VtLWhlYWx0aC1wYW5lbC9jZW0taGVhbHRoLXBhbmVsLmpzJztcbmltcG9ydCAnLi4vY2VtLW1hbmlmZXN0LWJyb3dzZXIvY2VtLW1hbmlmZXN0LWJyb3dzZXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcmVjb25uZWN0aW9uLWNvbnRlbnQvY2VtLXJlY29ubmVjdGlvbi1jb250ZW50LmpzJztcbmltcG9ydCAnLi4vY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8uanMnO1xuaW1wb3J0ICcuLi9jZW0tc2VydmUta25vYi1ncm91cC9jZW0tc2VydmUta25vYi1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1zZXJ2ZS1rbm9icy9jZW0tc2VydmUta25vYnMuanMnO1xuaW1wb3J0ICcuLi9jZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXkvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LWFsZXJ0L2NlbS1wZi12Ni1hbGVydC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1hbGVydC1ncm91cC9jZW0tcGYtdjYtYWxlcnQtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtYnV0dG9uL2NlbS1wZi12Ni1idXR0b24uanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtY2FyZC9jZW0tcGYtdjYtY2FyZC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1iYWRnZS9jZW0tcGYtdjYtYmFkZ2UuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtZHJvcGRvd24vY2VtLXBmLXY2LWRyb3Bkb3duLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbi9jZW0tcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LWxhYmVsL2NlbS1wZi12Ni1sYWJlbC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1tYXN0aGVhZC9jZW0tcGYtdjYtbWFzdGhlYWQuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtbW9kYWwvY2VtLXBmLXY2LW1vZGFsLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LW5hdi1ncm91cC9jZW0tcGYtdjYtbmF2LWdyb3VwLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LW5hdi1pdGVtL2NlbS1wZi12Ni1uYXYtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1uYXYtbGluay9jZW0tcGYtdjYtbmF2LWxpbmsuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtbmF2LWxpc3QvY2VtLXBmLXY2LW5hdi1saXN0LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LW5hdmlnYXRpb24vY2VtLXBmLXY2LW5hdmlnYXRpb24uanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtcGFnZS1tYWluL2NlbS1wZi12Ni1wYWdlLW1haW4uanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtcGFnZS1zaWRlYmFyL2NlbS1wZi12Ni1wYWdlLXNpZGViYXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtcGFnZS9jZW0tcGYtdjYtcGFnZS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1wb3BvdmVyL2NlbS1wZi12Ni1wb3BvdmVyLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXNlbGVjdC9jZW0tcGYtdjYtc2VsZWN0LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXNraXAtdG8tY29udGVudC9jZW0tcGYtdjYtc2tpcC10by1jb250ZW50LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXN3aXRjaC9jZW0tcGYtdjYtc3dpdGNoLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRhYi9jZW0tcGYtdjYtdGFiLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRhYnMvY2VtLXBmLXY2LXRhYnMuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cC9jZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10ZXh0LWlucHV0L2NlbS1wZi12Ni10ZXh0LWlucHV0LmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtL2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAvY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10b29sYmFyLWdyb3VwL2NlbS1wZi12Ni10b29sYmFyLWdyb3VwLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRvb2xiYXItaXRlbS9jZW0tcGYtdjYtdG9vbGJhci1pdGVtLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LXRvb2xiYXIvY2VtLXBmLXY2LXRvb2xiYXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtdHJlZS1pdGVtL2NlbS1wZi12Ni10cmVlLWl0ZW0uanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtdHJlZS12aWV3L2NlbS1wZi12Ni10cmVlLXZpZXcuanMnO1xuXG4vLyBDbGllbnQtb25seSBtb2R1bGVzIGxvYWRlZCBkeW5hbWljYWxseSB0byBhdm9pZCBicmVha2luZyBTU1IuXG4vLyBUaGVzZSBhcmUgcGxhaW4gSlMgbW9kdWxlcyBzZXJ2ZWQgYXQgcnVudGltZSBieSB0aGUgR28gc2VydmVyLlxudHlwZSBDRU1SZWxvYWRDbGllbnRUeXBlID0geyBuZXcob3B0czogYW55KTogYW55IH07XG50eXBlIFN0YXRlUGVyc2lzdGVuY2VUeXBlID0ge1xuICBnZXRTdGF0ZSgpOiBhbnk7XG4gIHVwZGF0ZVN0YXRlKHM6IGFueSk6IHZvaWQ7XG4gIGdldFRyZWVTdGF0ZSgpOiBhbnk7XG4gIHNldFRyZWVTdGF0ZShzOiBhbnkpOiB2b2lkO1xuICB1cGRhdGVUcmVlU3RhdGUoczogYW55KTogdm9pZDtcbiAgbWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2UoKTogdm9pZDtcbn07XG5sZXQgQ0VNUmVsb2FkQ2xpZW50OiBDRU1SZWxvYWRDbGllbnRUeXBlO1xubGV0IFN0YXRlUGVyc2lzdGVuY2U6IFN0YXRlUGVyc2lzdGVuY2VUeXBlO1xuXG5pbnRlcmZhY2UgRXZlbnRJbmZvIHtcbiAgZXZlbnROYW1lczogU2V0PHN0cmluZz47XG4gIGV2ZW50czogQXJyYXk8eyBuYW1lOiBzdHJpbmc7IHR5cGU/OiB7IHRleHQ6IHN0cmluZyB9OyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9Pjtcbn1cblxuaW50ZXJmYWNlIEV2ZW50UmVjb3JkIHtcbiAgaWQ6IHN0cmluZztcbiAgdGltZXN0YW1wOiBEYXRlO1xuICBldmVudE5hbWU6IHN0cmluZztcbiAgdGFnTmFtZTogc3RyaW5nO1xuICBlbGVtZW50SWQ6IHN0cmluZyB8IG51bGw7XG4gIGVsZW1lbnRDbGFzczogc3RyaW5nIHwgbnVsbDtcbiAgY3VzdG9tUHJvcGVydGllczogUmVjb3JkPHN0cmluZywgdW5rbm93bj47XG4gIG1hbmlmZXN0VHlwZTogc3RyaW5nIHwgbnVsbDtcbiAgc3VtbWFyeTogc3RyaW5nIHwgbnVsbDtcbiAgZGVzY3JpcHRpb246IHN0cmluZyB8IG51bGw7XG4gIGJ1YmJsZXM6IGJvb2xlYW47XG4gIGNvbXBvc2VkOiBib29sZWFuO1xuICBjYW5jZWxhYmxlOiBib29sZWFuO1xuICBkZWZhdWx0UHJldmVudGVkOiBib29sZWFuO1xufVxuXG5pbnRlcmZhY2UgRGVidWdEYXRhIHtcbiAgdmVyc2lvbj86IHN0cmluZztcbiAgb3M/OiBzdHJpbmc7XG4gIHdhdGNoRGlyPzogc3RyaW5nO1xuICBtYW5pZmVzdFNpemU/OiBzdHJpbmc7XG4gIGRlbW9Db3VudD86IG51bWJlcjtcbiAgZGVtb3M/OiBBcnJheTx7XG4gICAgdGFnTmFtZTogc3RyaW5nO1xuICAgIGRlc2NyaXB0aW9uPzogc3RyaW5nO1xuICAgIGZpbGVwYXRoPzogc3RyaW5nO1xuICAgIGNhbm9uaWNhbFVSTDogc3RyaW5nO1xuICAgIGxvY2FsUm91dGU6IHN0cmluZztcbiAgfT47XG4gIGltcG9ydE1hcD86IFJlY29yZDxzdHJpbmcsIHVua25vd24+O1xuICBpbXBvcnRNYXBKU09OPzogc3RyaW5nO1xufVxuXG5pbnRlcmZhY2UgTWFuaWZlc3Qge1xuICBtb2R1bGVzPzogQXJyYXk8e1xuICAgIGRlY2xhcmF0aW9ucz86IEFycmF5PHtcbiAgICAgIGN1c3RvbUVsZW1lbnQ/OiBib29sZWFuO1xuICAgICAgdGFnTmFtZT86IHN0cmluZztcbiAgICAgIG5hbWU/OiBzdHJpbmc7XG4gICAgICBraW5kPzogc3RyaW5nO1xuICAgICAgZXZlbnRzPzogQXJyYXk8eyBuYW1lOiBzdHJpbmc7IHR5cGU/OiB7IHRleHQ6IHN0cmluZyB9OyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9PjtcbiAgICB9PjtcbiAgfT47XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gbG9ncyBhcmUgcmVjZWl2ZWRcbiAqL1xuZXhwb3J0IGNsYXNzIENlbUxvZ3NFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgbG9nczogQXJyYXk8eyB0eXBlOiBzdHJpbmc7IGRhdGU6IHN0cmluZzsgbWVzc2FnZTogc3RyaW5nIH0+O1xuICBjb25zdHJ1Y3Rvcihsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT4pIHtcbiAgICBzdXBlcignY2VtOmxvZ3MnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gICAgdGhpcy5sb2dzID0gbG9ncztcbiAgfVxufVxuXG4vKipcbiAqIENFTSBTZXJ2ZSBDaHJvbWUgLSBNYWluIGRlbW8gd3JhcHBlciBjb21wb25lbnRcbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgZGVtbyBjb250ZW50XG4gKiBAc2xvdCBuYXZpZ2F0aW9uIC0gTmF2aWdhdGlvbiBzaWRlYmFyIGNvbnRlbnRcbiAqIEBzbG90IGtub2JzIC0gS25vYiBjb250cm9sc1xuICogQHNsb3QgZGVzY3JpcHRpb24gLSBEZW1vIGRlc2NyaXB0aW9uXG4gKiBAc2xvdCBtYW5pZmVzdC10cmVlIC0gTWFuaWZlc3QgdHJlZSB2aWV3XG4gKiBAc2xvdCBtYW5pZmVzdC1uYW1lIC0gTWFuaWZlc3QgbmFtZSBkaXNwbGF5XG4gKiBAc2xvdCBtYW5pZmVzdC1kZXRhaWxzIC0gTWFuaWZlc3QgZGV0YWlscyBkaXNwbGF5XG4gKlxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXNlcnZlLWNocm9tZVxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXNlcnZlLWNocm9tZScpXG5leHBvcnQgY2xhc3MgQ2VtU2VydmVDaHJvbWUgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAvLyBTdGF0aWMgdGVtcGxhdGVzIGZvciBET00gY2xvbmluZyAobG9ncywgZXZlbnRzLCBkZWJ1ZyBpbmZvKVxuICBzdGF0aWMgI2RlbW9JbmZvVGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPGgzPkRlbW8gSW5mb3JtYXRpb248L2gzPlxuICAgICAgPGRsIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdCBwZi1tLWhvcml6b250YWwgcGYtbS1jb21wYWN0XCI+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5UYWcgTmFtZTwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwidGFnTmFtZVwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIiBkYXRhLWZpZWxkLWdyb3VwPVwiZGVzY3JpcHRpb25cIj5cbiAgICAgICAgICA8ZHQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+RGVzY3JpcHRpb248L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cImRlc2NyaXB0aW9uXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiIGRhdGEtZmllbGQtZ3JvdXA9XCJmaWxlcGF0aFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5GaWxlIFBhdGg8L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cImZpbGVwYXRoXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5DYW5vbmljYWwgVVJMPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJjYW5vbmljYWxVUkxcIj48L2RkPlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkxvY2FsIFJvdXRlPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICA8L2RsPmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgc3RhdGljICNkZW1vR3JvdXBUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiIGRhdGEtZmllbGQ9XCJ0YWdOYW1lXCI+PC9kdD5cbiAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIj5cbiAgICAgICAgICA8c3BhbiBkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIj48L3NwYW4+PGJyPlxuICAgICAgICAgIDxzbWFsbCBkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIj5GaWxlOiA8c3BhbiBkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIj48L3NwYW4+PC9zbWFsbD5cbiAgICAgICAgICA8c21hbGw+Q2Fub25pY2FsOiA8c3BhbiBkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCI+PC9zcGFuPjwvc21hbGw+PGJyPlxuICAgICAgICAgIDxzbWFsbD5Mb2NhbDogPHNwYW4gZGF0YS1maWVsZD1cImxvY2FsUm91dGVcIj48L3NwYW4+PC9zbWFsbD5cbiAgICAgICAgPC9kZD5cbiAgICAgIDwvZGl2PmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgc3RhdGljICNkZW1vTGlzdFRlbXBsYXRlID0gKCgpID0+IHtcbiAgICBjb25zdCB0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGVtcGxhdGUnKTtcbiAgICB0LmlubmVySFRNTCA9IGBcbiAgICAgIDxjZW0tcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uIGlkPVwiZGVidWctZGVtb3Mtc2VjdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHRvZ2dsZS10ZXh0PVwiU2hvdyBEZW1vcyBJbmZvXCI+XG4gICAgICAgIDxkbCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3QgcGYtbS1ob3Jpem9udGFsIHBmLW0tY29tcGFjdFwiIGRhdGEtY29udGFpbmVyPVwiZ3JvdXBzXCI+PC9kbD5cbiAgICAgIDwvY2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbj5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjbG9nRW50cnlUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8ZGl2IGNsYXNzPVwibG9nLWVudHJ5XCIgZGF0YS1maWVsZD1cImNvbnRhaW5lclwiPlxuICAgICAgICA8Y2VtLXBmLXY2LWxhYmVsIGNvbXBhY3QgZGF0YS1maWVsZD1cImxhYmVsXCI+PC9jZW0tcGYtdjYtbGFiZWw+XG4gICAgICAgIDx0aW1lIGNsYXNzPVwibG9nLXRpbWVcIiBkYXRhLWZpZWxkPVwidGltZVwiPjwvdGltZT5cbiAgICAgICAgPHNwYW4gY2xhc3M9XCJsb2ctbWVzc2FnZVwiIGRhdGEtZmllbGQ9XCJtZXNzYWdlXCI+PC9zcGFuPlxuICAgICAgPC9kaXY+YDtcbiAgICByZXR1cm4gdDtcbiAgfSkoKTtcblxuICBzdGF0aWMgI2V2ZW50RW50cnlUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8YnV0dG9uIGNsYXNzPVwiZXZlbnQtbGlzdC1pdGVtXCIgZGF0YS1maWVsZD1cImNvbnRhaW5lclwiPlxuICAgICAgICA8Y2VtLXBmLXY2LWxhYmVsIGNvbXBhY3QgZGF0YS1maWVsZD1cImxhYmVsXCI+PC9jZW0tcGYtdjYtbGFiZWw+XG4gICAgICAgIDx0aW1lIGNsYXNzPVwiZXZlbnQtdGltZVwiIGRhdGEtZmllbGQ9XCJ0aW1lXCI+PC90aW1lPlxuICAgICAgICA8c3BhbiBjbGFzcz1cImV2ZW50LWVsZW1lbnRcIiBkYXRhLWZpZWxkPVwiZWxlbWVudFwiPjwvc3Bhbj5cbiAgICAgIDwvYnV0dG9uPmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAncHJpbWFyeS10YWctbmFtZScgfSlcbiAgYWNjZXNzb3IgcHJpbWFyeVRhZ05hbWUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdkZW1vLXRpdGxlJyB9KVxuICBhY2Nlc3NvciBkZW1vVGl0bGUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdwYWNrYWdlLW5hbWUnIH0pXG4gIGFjY2Vzc29yIHBhY2thZ2VOYW1lID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnY2Fub25pY2FsLXVybCcgfSlcbiAgYWNjZXNzb3IgY2Fub25pY2FsVVJMID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnc291cmNlLXVybCcgfSlcbiAgYWNjZXNzb3Igc291cmNlVVJMID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3Iga25vYnMgPSAnJztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBkcmF3ZXI6ICdleHBhbmRlZCcgfCAnY29sbGFwc2VkJyB8ICcnID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnZHJhd2VyLWhlaWdodCcgfSlcbiAgYWNjZXNzb3IgZHJhd2VySGVpZ2h0ID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAndGFicy1zZWxlY3RlZCcgfSlcbiAgYWNjZXNzb3IgdGFic1NlbGVjdGVkID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3Igc2lkZWJhcjogJ2V4cGFuZGVkJyB8ICdjb2xsYXBzZWQnIHwgJycgPSAnJztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCBhdHRyaWJ1dGU6ICdoYXMtZGVzY3JpcHRpb24nIH0pXG4gIGFjY2Vzc29yIGhhc0Rlc2NyaXB0aW9uID0gZmFsc2U7XG5cbiAgIyQoaWQ6IHN0cmluZykge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKGlkKTtcbiAgfVxuXG4gICMkJChzZWxlY3Rvcjogc3RyaW5nKSB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvckFsbChzZWxlY3RvcikgPz8gW107XG4gIH1cblxuICAjbG9nQ29udGFpbmVyOiBIVE1MRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjZHJhd2VyT3BlbiA9IGZhbHNlO1xuICAjaW5pdGlhbExvZ3NGZXRjaGVkID0gZmFsc2U7XG4gICNpc0luaXRpYWxMb2FkID0gdHJ1ZTtcbiAgI2RlYnVnRGF0YTogRGVidWdEYXRhIHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gRWxlbWVudCBldmVudCB0cmFja2luZ1xuICAjZWxlbWVudEV2ZW50TWFwOiBNYXA8c3RyaW5nLCBFdmVudEluZm8+IHwgbnVsbCA9IG51bGw7XG4gICNtYW5pZmVzdDogTWFuaWZlc3QgfCBudWxsID0gbnVsbDtcbiAgI2NhcHR1cmVkRXZlbnRzOiBFdmVudFJlY29yZFtdID0gW107XG4gICNtYXhDYXB0dXJlZEV2ZW50cyA9IDEwMDA7XG4gICNldmVudExpc3Q6IEhUTUxFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNldmVudERldGFpbEhlYWRlcjogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50RGV0YWlsQm9keTogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI3NlbGVjdGVkRXZlbnRJZDogc3RyaW5nIHwgbnVsbCA9IG51bGw7XG4gICNldmVudHNGaWx0ZXJWYWx1ZSA9ICcnO1xuICAjZXZlbnRzRmlsdGVyRGVib3VuY2VUaW1lcjogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50VHlwZUZpbHRlcnMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgI2VsZW1lbnRGaWx0ZXJzID0gbmV3IFNldDxzdHJpbmc+KCk7XG4gICNkaXNjb3ZlcmVkRWxlbWVudHMgPSBuZXcgU2V0PHN0cmluZz4oKTtcblxuICAvLyBFdmVudCBsaXN0ZW5lciByZWZlcmVuY2VzIGZvciBjbGVhbnVwXG4gICNoYW5kbGVMb2dzRXZlbnQ6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZUV4cGFuZDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG4gICNoYW5kbGVUcmVlQ29sbGFwc2U6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZVNlbGVjdDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gVGltZW91dCBJRHMgZm9yIGNsZWFudXBcbiAgI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjY29weURlYnVnRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dDogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcblxuICAvLyBMb2cgZmlsdGVyIHN0YXRlXG4gICNsb2dzRmlsdGVyVmFsdWUgPSAnJztcbiAgI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyOiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjbG9nTGV2ZWxGaWx0ZXJzID0gbmV3IFNldChbJ2luZm8nLCAnd2FybicsICdlcnJvcicsICdkZWJ1ZyddKTtcbiAgI2xvZ0xldmVsRHJvcGRvd246IEVsZW1lbnQgfCBudWxsID0gbnVsbDtcblxuICAvLyBXYXRjaCBmb3IgZHluYW1pY2FsbHkgYWRkZWQgZWxlbWVudHNcbiAgLyogYzggaWdub3JlIHN0YXJ0IC0gTXV0YXRpb25PYnNlcnZlciBjYWxsYmFjayB0ZXN0ZWQgdmlhIGludGVncmF0aW9uICovXG4gICNvYnNlcnZlciA9IG5ldyBNdXRhdGlvbk9ic2VydmVyKChtdXRhdGlvbnMpID0+IHtcbiAgICBsZXQgbmVlZHNVcGRhdGUgPSBmYWxzZTtcblxuICAgIGZvciAoY29uc3QgbXV0YXRpb24gb2YgbXV0YXRpb25zKSB7XG4gICAgICBmb3IgKGNvbnN0IG5vZGUgb2YgbXV0YXRpb24uYWRkZWROb2Rlcykge1xuICAgICAgICBpZiAobm9kZSBpbnN0YW5jZW9mIEhUTUxFbGVtZW50KSB7XG4gICAgICAgICAgY29uc3QgdGFnTmFtZSA9IG5vZGUudGFnTmFtZS50b0xvd2VyQ2FzZSgpO1xuICAgICAgICAgIGlmICh0aGlzLiNlbGVtZW50RXZlbnRNYXA/Lmhhcyh0YWdOYW1lKSAmJiAhbm9kZS5kYXRhc2V0LmNlbUV2ZW50c0F0dGFjaGVkKSB7XG4gICAgICAgICAgICBjb25zdCBldmVudEluZm8gPSB0aGlzLiNlbGVtZW50RXZlbnRNYXAuZ2V0KHRhZ05hbWUpITtcbiAgICAgICAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGV2ZW50SW5mby5ldmVudE5hbWVzKSB7XG4gICAgICAgICAgICAgIG5vZGUuYWRkRXZlbnRMaXN0ZW5lcihldmVudE5hbWUsIHRoaXMuI2hhbmRsZUVsZW1lbnRFdmVudCwgeyBjYXB0dXJlOiB0cnVlIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgbm9kZS5kYXRhc2V0LmNlbUV2ZW50c0F0dGFjaGVkID0gJ3RydWUnO1xuICAgICAgICAgICAgdGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICAgICAgICAgIG5lZWRzVXBkYXRlID0gdHJ1ZTtcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICBpZiAobmVlZHNVcGRhdGUpIHtcbiAgICAgIHRoaXMuI3VwZGF0ZUV2ZW50RmlsdGVycygpO1xuICAgIH1cbiAgfSk7XG4gIC8qIGM4IGlnbm9yZSBzdG9wICovXG5cbiAgI3dzQ2xpZW50OiBhbnk7XG5cbiAgI2luaXRXc0NsaWVudCgpIHtcbiAgICB0aGlzLiN3c0NsaWVudCA9IG5ldyBDRU1SZWxvYWRDbGllbnQoe1xuICAgIGppdHRlck1heDogMTAwMCxcbiAgICBvdmVybGF5VGhyZXNob2xkOiAxNSxcbiAgICBiYWRnZUZhZGVEZWxheTogMjAwMCxcbiAgICAvKiBjOCBpZ25vcmUgc3RhcnQgLSBXZWJTb2NrZXQgY2FsbGJhY2tzIHRlc3RlZCB2aWEgaW50ZWdyYXRpb24gKi9cbiAgICBjYWxsYmFja3M6IHtcbiAgICAgIG9uT3BlbjogKCkgPT4ge1xuICAgICAgICB0aGlzLiMkKCdyZWNvbm5lY3Rpb24tbW9kYWwnKT8uY2xvc2UoKTtcbiAgICAgIH0sXG4gICAgICBvbkVycm9yOiAoZXJyb3JEYXRhOiB7IHRpdGxlPzogc3RyaW5nOyBtZXNzYWdlPzogc3RyaW5nOyBmaWxlPzogc3RyaW5nIH0pID0+IHtcbiAgICAgICAgaWYgKGVycm9yRGF0YT8udGl0bGUgJiYgZXJyb3JEYXRhPy5tZXNzYWdlKSB7XG4gICAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZV0gU2VydmVyIGVycm9yOicsIGVycm9yRGF0YSk7XG4gICAgICAgICAgKHRoaXMuIyQoJ2Vycm9yLW92ZXJsYXknKSBhcyBhbnkpPy5zaG93KGVycm9yRGF0YS50aXRsZSwgZXJyb3JEYXRhLm1lc3NhZ2UsIGVycm9yRGF0YS5maWxlKTtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlXSBXZWJTb2NrZXQgZXJyb3I6JywgZXJyb3JEYXRhKTtcbiAgICAgICAgfVxuICAgICAgfSxcbiAgICAgIG9uUmVjb25uZWN0aW5nOiAoeyBhdHRlbXB0LCBkZWxheSB9OiB7IGF0dGVtcHQ6IG51bWJlcjsgZGVsYXk6IG51bWJlciB9KSA9PiB7XG4gICAgICAgIGlmIChhdHRlbXB0ID49IDE1KSB7XG4gICAgICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LnNob3dNb2RhbCgpO1xuICAgICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tY29udGVudCcpIGFzIGFueSk/LnVwZGF0ZVJldHJ5SW5mbyhhdHRlbXB0LCBkZWxheSk7XG4gICAgICAgIH1cbiAgICAgIH0sXG4gICAgICBvblJlbG9hZDogKCkgPT4ge1xuICAgICAgICBjb25zdCBlcnJvck92ZXJsYXkgPSB0aGlzLiMkKCdlcnJvci1vdmVybGF5Jyk7XG4gICAgICAgIGlmIChlcnJvck92ZXJsYXk/Lmhhc0F0dHJpYnV0ZSgnb3BlbicpKSB7XG4gICAgICAgICAgKGVycm9yT3ZlcmxheSBhcyBhbnkpLmhpZGUoKTtcbiAgICAgICAgfVxuICAgICAgICB3aW5kb3cubG9jYXRpb24ucmVsb2FkKCk7XG4gICAgICB9LFxuICAgICAgb25TaHV0ZG93bjogKCkgPT4ge1xuICAgICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLW1vZGFsJykgYXMgYW55KT8uc2hvd01vZGFsKCk7XG4gICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tY29udGVudCcpIGFzIGFueSk/LnVwZGF0ZVJldHJ5SW5mbygzMCwgMzAwMDApO1xuICAgICAgfSxcbiAgICAgIG9uTG9nczogKGxvZ3M6IEFycmF5PHsgdHlwZTogc3RyaW5nOyBkYXRlOiBzdHJpbmc7IG1lc3NhZ2U6IHN0cmluZyB9PikgPT4ge1xuICAgICAgICB3aW5kb3cuZGlzcGF0Y2hFdmVudChuZXcgQ2VtTG9nc0V2ZW50KGxvZ3MpKTtcbiAgICAgIH1cbiAgICB9XG4gICAgLyogYzggaWdub3JlIHN0b3AgKi9cbiAgICB9KTtcbiAgfVxuXG4gIGdldCBkZW1vKCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5xdWVyeVNlbGVjdG9yKCdjZW0tc2VydmUtZGVtbycpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGxpbmsgcmVsPVwic3R5bGVzaGVldFwiIGhyZWY9XCIvX19jZW0vY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdC5jc3NcIj5cbiAgICAgIDxsaW5rIHJlbD1cInN0eWxlc2hlZXRcIiBocmVmPVwiL19fY2VtL3BmLWxpZ2h0ZG9tLmNzc1wiPlxuXG4gICAgICA8Y2VtLXBmLXY2LXBhZ2UgP3NpZGViYXItY29sbGFwc2VkPSR7dGhpcy5zaWRlYmFyID09PSAnY29sbGFwc2VkJ30+XG4gICAgICAgIDxjZW0tcGYtdjYtc2tpcC10by1jb250ZW50IHNsb3Q9XCJza2lwLXRvLWNvbnRlbnRcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhyZWY9XCIjbWFpbi1jb250ZW50XCI+XG4gICAgICAgICAgU2tpcCB0byBjb250ZW50XG4gICAgICAgIDwvY2VtLXBmLXY2LXNraXAtdG8tY29udGVudD5cblxuICAgICAgICA8Y2VtLXBmLXY2LW1hc3RoZWFkIHNsb3Q9XCJtYXN0aGVhZFwiPlxuICAgICAgICAgIDxhIGNsYXNzPVwibWFzdGhlYWQtbG9nb1wiXG4gICAgICAgICAgICAgaHJlZj1cIi9cIlxuICAgICAgICAgICAgIHNsb3Q9XCJsb2dvXCI+XG4gICAgICAgICAgICA8aW1nIGFsdD1cIkNFTSBEZXYgU2VydmVyXCJcbiAgICAgICAgICAgICAgICAgc3JjPVwiL19fY2VtL2xvZ28uc3ZnXCI+XG4gICAgICAgICAgICAke3RoaXMucGFja2FnZU5hbWUgPyBodG1sYDxoMT4ke3RoaXMucGFja2FnZU5hbWV9PC9oMT5gIDogbm90aGluZ31cbiAgICAgICAgICA8L2E+XG4gICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyIHNsb3Q9XCJ0b29sYmFyXCI+XG4gICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItZ3JvdXAgdmFyaWFudD1cImFjdGlvbi1ncm91cFwiPlxuICAgICAgICAgICAgICAke3RoaXMuI3JlbmRlclNvdXJjZUJ1dHRvbigpfVxuICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBpZD1cImRlYnVnLWluZm9cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInBsYWluXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJEZWJ1ZyBpbmZvXCI+XG4gICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxNlwiXG4gICAgICAgICAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTggMTVBNyA3IDAgMSAxIDggMWE3IDcgMCAwIDEgMCAxNHptMCAxQTggOCAwIDEgMCA4IDBhOCA4IDAgMCAwIDAgMTZ6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwibTguOTMgNi41ODgtMi4yOS4yODctLjA4Mi4zOC40NS4wODNjLjI5NC4wNy4zNTIuMTc2LjI4OC40NjlsLS43MzggMy40NjhjLS4xOTQuODk3LjEwNSAxLjMxOS44MDggMS4zMTkuNTQ1IDAgMS4xNzgtLjI1MiAxLjQ2NS0uNTk4bC4wODgtLjQxNmMtLjIuMTc2LS40OTIuMjQ2LS42ODYuMjQ2LS4yNzUgMC0uMzc1LS4xOTMtLjMwNC0uNTMzTDguOTMgNi41ODh6TTkgNC41YTEgMSAwIDEgMS0yIDAgMSAxIDAgMCAxIDIgMHpcIi8+XG4gICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni1idXR0b24+XG4gICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b2dnbGUtZ3JvdXAgY2xhc3M9XCJjb2xvci1zY2hlbWUtdG9nZ2xlXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJDb2xvciBzY2hlbWVcIj5cbiAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0gdmFsdWU9XCJsaWdodFwiPlxuICAgICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCIgYXJpYS1sYWJlbD1cIkxpZ2h0IG1vZGVcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTggMTFhMyAzIDAgMSAxIDAtNiAzIDMgMCAwIDEgMCA2em0wIDFhNCA0IDAgMSAwIDAtOCA0IDQgMCAwIDAgMCA4ek04IDBhLjUuNSAwIDAgMSAuNS41djJhLjUuNSAwIDAgMS0xIDB2LTJBLjUuNSAwIDAgMSA4IDB6bTAgMTNhLjUuNSAwIDAgMSAuNS41djJhLjUuNSAwIDAgMS0xIDB2LTJBLjUuNSAwIDAgMSA4IDEzem04LTVhLjUuNSAwIDAgMS0uNS41aC0yYS41LjUgMCAwIDEgMC0xaDJhLjUuNSAwIDAgMSAuNS41ek0zIDhhLjUuNSAwIDAgMS0uNS41aC0yYS41LjUgMCAwIDEgMC0xaDJBLjUuNSAwIDAgMSAzIDh6bTEwLjY1Ny01LjY1N2EuNS41IDAgMCAxIDAgLjcwN2wtMS40MTQgMS40MTVhLjUuNSAwIDEgMS0uNzA3LS43MDhsMS40MTQtMS40MTRhLjUuNSAwIDAgMSAuNzA3IDB6bS05LjE5MyA5LjE5M2EuNS41IDAgMCAxIDAgLjcwN0wzLjA1IDEzLjY1N2EuNS41IDAgMCAxLS43MDctLjcwN2wxLjQxNC0xLjQxNGEuNS41IDAgMCAxIC43MDcgMHptOS4xOTMgMi4xMjFhLjUuNSAwIDAgMS0uNzA3IDBsLTEuNDE0LTEuNDE0YS41LjUgMCAwIDEgLjcwNy0uNzA3bDEuNDE0IDEuNDE0YS41LjUgMCAwIDEgMCAuNzA3ek00LjQ2NCA0LjQ2NWEuNS41IDAgMCAxLS43MDcgMEwyLjM0MyAzLjA1YS41LjUgMCAxIDEgLjcwNy0uNzA3bDEuNDE0IDEuNDE0YS41LjUgMCAwIDEgMCAuNzA4elwiLz5cbiAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0gdmFsdWU9XCJkYXJrXCI+XG4gICAgICAgICAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIiBhcmlhLWxhYmVsPVwiRGFyayBtb2RlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk02IC4yNzhhLjc2OC43NjggMCAwIDEgLjA4Ljg1OCA3LjIwOCA3LjIwOCAwIDAgMC0uODc4IDMuNDZjMCA0LjAyMSAzLjI3OCA3LjI3NyA3LjMxOCA3LjI3Ny41MjcgMCAxLjA0LS4wNTUgMS41MzMtLjE2YS43ODcuNzg3IDAgMCAxIC44MS4zMTYuNzMzLjczMyAwIDAgMS0uMDMxLjg5M0E4LjM0OSA4LjM0OSAwIDAgMSA4LjM0NCAxNkMzLjczNCAxNiAwIDEyLjI4NiAwIDcuNzEgMCA0LjI2NiAyLjExNCAxLjMxMiA1LjEyNC4wNkEuNzUyLjc1MiAwIDAgMSA2IC4yNzh6TTQuODU4IDEuMzExQTcuMjY5IDcuMjY5IDAgMCAwIDEuMDI1IDcuNzFjMCA0LjAyIDMuMjc5IDcuMjc2IDcuMzE5IDcuMjc2YTcuMzE2IDcuMzE2IDAgMCAwIDUuMjA1LTIuMTYyYy0uMzM3LjA0Mi0uNjguMDYzLTEuMDI5LjA2My00LjYxIDAtOC4zNDMtMy43MTQtOC4zNDMtOC4yOSAwLTEuMTY3LjI0Mi0yLjI3OC42ODEtMy4yODZ6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtPlxuICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cInN5c3RlbVwiPlxuICAgICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCIgYXJpYS1sYWJlbD1cIlN5c3RlbSBwcmVmZXJlbmNlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0wIDEuNUExLjUgMS41IDAgMCAxIDEuNSAwaDEzQTEuNSAxLjUgMCAwIDEgMTYgMS41djhhMS41IDEuNSAwIDAgMS0xLjUgMS41aC0xM0ExLjUgMS41IDAgMCAxIDAgOS41di04ek0xLjUgMWEuNS41IDAgMCAwLS41LjV2OGEuNS41IDAgMCAwIC41LjVoMTNhLjUuNSAwIDAgMCAuNS0uNXYtOGEuNS41IDAgMCAwLS41LS41aC0xM3pcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0yLjUgMTJoMTFhLjUuNSAwIDAgMSAwIDFoLTExYS41LjUgMCAwIDEgMC0xem0wIDJoMTFhLjUuNSAwIDAgMSAwIDFoLTExYS41LjUgMCAwIDEgMC0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b2dnbGUtZ3JvdXA+XG4gICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhcj5cbiAgICAgICAgPC9jZW0tcGYtdjYtbWFzdGhlYWQ+XG5cbiAgICAgICAgPGNlbS1wZi12Ni1wYWdlLXNpZGViYXIgc2xvdD1cInNpZGViYXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgID9leHBhbmRlZD0ke3RoaXMuc2lkZWJhciA9PT0gJ2V4cGFuZGVkJ31cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA/Y29sbGFwc2VkPSR7dGhpcy5zaWRlYmFyICE9PSAnZXhwYW5kZWQnfT5cbiAgICAgICAgICA8c2xvdCBuYW1lPVwibmF2aWdhdGlvblwiPjwvc2xvdD5cbiAgICAgICAgPC9jZW0tcGYtdjYtcGFnZS1zaWRlYmFyPlxuXG4gICAgICAgIDxjZW0tcGYtdjYtcGFnZS1tYWluIHNsb3Q9XCJtYWluXCIgaWQ9XCJtYWluLWNvbnRlbnRcIj5cbiAgICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICAgICAgPGZvb3RlciBjbGFzcz1cInBmLW0tc3RpY2t5LWJvdHRvbVwiPlxuICAgICAgICAgICAgPGRpdiBjbGFzcz1cImZvb3Rlci1kZXNjcmlwdGlvbiR7dGhpcy5oYXNEZXNjcmlwdGlvbiA/ICcnIDogJyBlbXB0eSd9XCI+XG4gICAgICAgICAgICAgIDxzbG90IG5hbWU9XCJkZXNjcmlwdGlvblwiPjwvc2xvdD5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgPGNlbS1kcmF3ZXIgP29wZW49JHt0aGlzLmRyYXdlciA9PT0gJ2V4cGFuZGVkJ31cbiAgICAgICAgICAgICAgICAgICAgICAgIGRyYXdlci1oZWlnaHQ9XCIke3RoaXMuZHJhd2VySGVpZ2h0IHx8ICc0MDAnfVwiPlxuICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRhYnMgc2VsZWN0ZWQ9XCIke3RoaXMudGFic1NlbGVjdGVkIHx8ICcwJ31cIj5cbiAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRhYiB0aXRsZT1cIktub2JzXCI+XG4gICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwia25vYnMtY29udGFpbmVyXCI+XG4gICAgICAgICAgICAgICAgICAgIDxzbG90IG5hbWU9XCJrbm9ic1wiPlxuICAgICAgICAgICAgICAgICAgICAgIDxwIGNsYXNzPVwia25vYnMtZW1wdHlcIj5ObyBrbm9icyBhdmFpbGFibGUgZm9yIHRoaXMgZWxlbWVudC48L3A+XG4gICAgICAgICAgICAgICAgICAgIDwvc2xvdD5cbiAgICAgICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRhYj5cbiAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRhYiB0aXRsZT1cIk1hbmlmZXN0IEJyb3dzZXJcIj5cbiAgICAgICAgICAgICAgICAgIDxjZW0tbWFuaWZlc3QtYnJvd3Nlcj5cbiAgICAgICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cIm1hbmlmZXN0LXRyZWVcIiBzbG90PVwibWFuaWZlc3QtdHJlZVwiPjwvc2xvdD5cbiAgICAgICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cIm1hbmlmZXN0LW5hbWVcIiBzbG90PVwibWFuaWZlc3QtbmFtZVwiPjwvc2xvdD5cbiAgICAgICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cIm1hbmlmZXN0LWRldGFpbHNcIiBzbG90PVwibWFuaWZlc3QtZGV0YWlsc1wiPjwvc2xvdD5cbiAgICAgICAgICAgICAgICAgIDwvY2VtLW1hbmlmZXN0LWJyb3dzZXI+XG4gICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdGFiPlxuICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdGFiIHRpdGxlPVwiU2VydmVyIExvZ3NcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJsb2dzLXdyYXBwZXJcIj5cbiAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyIHN0aWNreT5cbiAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10ZXh0LWlucHV0LWdyb3VwIGlkPVwibG9ncy1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBwbGFjZWhvbGRlcj1cIkZpbHRlciBsb2dzLi4uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaWNvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxZW1cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgd2lkdGg9XCIxZW1cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCA1MTIgNTEyXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTUwNSA0NDIuN0w0MDUuMyAzNDNjLTQuNS00LjUtMTAuNi03LTE3LTdIMzcyYzI3LjYtMzUuMyA0NC03OS43IDQ0LTEyOEM0MTYgOTMuMSAzMjIuOSAwIDIwOCAwUzAgOTMuMSAwIDIwOHM5My4xIDIwOCAyMDggMjA4YzQ4LjMgMCA5Mi43LTE2LjQgMTI4LTQ0djE2LjNjMCA2LjQgMi41IDEyLjUgNyAxN2w5OS43IDk5LjdjOS40IDkuNCAyNC42IDkuNCAzMy45IDBsMjguMy0yOC4zYzkuNC05LjQgOS40LTI0LjYuMS0zNHpNMjA4IDMzNmMtNzAuNyAwLTEyOC01Ny4yLTEyOC0xMjggMC03MC43IDU3LjItMTI4IDEyOC0xMjggNzAuNyAwIDEyOCA1Ny4yIDEyOCAxMjggMCA3MC43LTU3LjIgMTI4LTEyOCAxMjh6XCI+PC9wYXRoPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10ZXh0LWlucHV0LWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtZHJvcGRvd24gaWQ9XCJsb2ctbGV2ZWwtZmlsdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGxhYmVsPVwiRmlsdGVyIGxvZyBsZXZlbHNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5Mb2cgTGV2ZWxzPC9zcGFuPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwiaW5mb1wiIGNoZWNrZWQ+SW5mbzwvY2VtLXBmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LW1lbnUtaXRlbSB2YXJpYW50PVwiY2hlY2tib3hcIiB2YWx1ZT1cIndhcm5cIiBjaGVja2VkPldhcm5pbmdzPC9jZW0tcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwiZXJyb3JcIiBjaGVja2VkPkVycm9yczwvY2VtLXBmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LW1lbnUtaXRlbSB2YXJpYW50PVwiY2hlY2tib3hcIiB2YWx1ZT1cImRlYnVnXCIgY2hlY2tlZD5EZWJ1ZzwvY2VtLXBmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtZHJvcGRvd24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItZ3JvdXAgdmFyaWFudD1cImFjdGlvbi1ncm91cFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtYnV0dG9uIGlkPVwiY29weS1sb2dzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0xMyAwSDZhMiAyIDAgMCAwLTIgMiAyIDIgMCAwIDAtMiAydjEwYTIgMiAwIDAgMCAyIDJoN2EyIDIgMCAwIDAgMi0yIDIgMiAwIDAgMCAyLTJWMmEyIDIgMCAwIDAtMi0yem0wIDEzVjRhMiAyIDAgMCAwLTItMkg1YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjEwYTEgMSAwIDAgMS0xIDF6TTMgMTNWNGExIDEgMCAwIDEgMS0xaDdhMSAxIDAgMCAxIDEgMXY5YTEgMSAwIDAgMS0xIDFINGExIDEgMCAwIDEtMS0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDb3B5IExvZ3NcbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtYnV0dG9uPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXI+XG4gICAgICAgICAgICAgICAgICAgIDxkaXYgaWQ9XCJsb2ctY29udGFpbmVyXCI+PC9kaXY+XG4gICAgICAgICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10YWI+XG4gICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10YWIgdGl0bGU9XCJFdmVudHNcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJldmVudHMtd3JhcHBlclwiPlxuICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXIgc3RpY2t5PlxuICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRleHQtaW5wdXQtZ3JvdXAgaWQ9XCJldmVudHMtZmlsdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgcGxhY2Vob2xkZXI9XCJGaWx0ZXIgZXZlbnRzLi4uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaWNvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCIgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNMTEuNzQyIDEwLjM0NGE2LjUgNi41IDAgMSAwLTEuMzk3IDEuMzk4aC0uMDAxYy4wMy4wNC4wNjIuMDc4LjA5OC4xMTVsMy44NSAzLjg1YTEgMSAwIDAgMCAxLjQxNS0xLjQxNGwtMy44NS0zLjg1YTEuMDA3IDEuMDA3IDAgMCAwLS4xMTUtLjF6TTEyIDYuNWE1LjUgNS41IDAgMSAxLTExIDAgNS41IDUuNSAwIDAgMSAxMSAwelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdGV4dC1pbnB1dC1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LWRyb3Bkb3duIGlkPVwiZXZlbnQtdHlwZS1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgZXZlbnQgdHlwZXNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5FdmVudCBUeXBlczwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtZHJvcGRvd24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni1kcm9wZG93biBpZD1cImVsZW1lbnQtZmlsdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGxhYmVsPVwiRmlsdGVyIGVsZW1lbnRzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHNwYW4gc2xvdD1cInRvZ2dsZS10ZXh0XCI+RWxlbWVudHM8L3NwYW4+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LWRyb3Bkb3duPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWdyb3VwIHZhcmlhbnQ9XCJhY3Rpb24tZ3JvdXBcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBpZD1cImNsZWFyLWV2ZW50c1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInRlcnRpYXJ5XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBzaXplPVwic21hbGxcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCIgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNS41IDUuNUEuNS41IDAgMCAxIDYgNnY2YS41LjUgMCAwIDEtMSAwVjZhLjUuNSAwIDAgMSAuNS0uNXptMi41IDBhLjUuNSAwIDAgMSAuNS41djZhLjUuNSAwIDAgMS0xIDBWNmEuNS41IDAgMCAxIC41LS41em0zIC41YS41LjUgMCAwIDAtMSAwdjZhLjUuNSAwIDAgMCAxIDBWNnpcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBmaWxsLXJ1bGU9XCJldmVub2RkXCIgZD1cIk0xNC41IDNhMSAxIDAgMCAxLTEgMUgxM3Y5YTIgMiAwIDAgMS0yIDJINWEyIDIgMCAwIDEtMi0yVjRoLS41YTEgMSAwIDAgMS0xLTFWMmExIDEgMCAwIDEgMS0xSDZhMSAxIDAgMCAxIDEtMWgyYTEgMSAwIDAgMSAxIDFoMy41YTEgMSAwIDAgMSAxIDF2MXpNNC4xMTggNCA0IDQuMDU5VjEzYTEgMSAwIDAgMCAxIDFoNmExIDEgMCAwIDAgMS0xVjQuMDU5TDExLjg4MiA0SDQuMTE4ek0yLjUgM1YyaDExdjFoLTExelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDbGVhciBFdmVudHNcbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtYnV0dG9uPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPGNlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtYnV0dG9uIGlkPVwiY29weS1ldmVudHNcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJ0ZXJ0aWFyeVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgc2l6ZT1cInNtYWxsXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHN2ZyBzbG90PVwiaWNvblwiIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTEzIDBINmEyIDIgMCAwIDAtMiAyIDIgMiAwIDAgMC0yIDJ2MTBhMiAyIDAgMCAwIDIgMmg3YTIgMiAwIDAgMCAyLTIgMiAyIDAgMCAwIDItMlYyYTIgMiAwIDAgMC0yLTJ6bTAgMTNWNGEyIDIgMCAwIDAtMi0ySDVhMSAxIDAgMCAxIDEtMWg3YTEgMSAwIDAgMSAxIDF2MTBhMSAxIDAgMCAxLTEgMXpNMyAxM1Y0YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjlhMSAxIDAgMCAxLTEgMUg0YTEgMSAwIDAgMS0xLTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIENvcHkgRXZlbnRzXG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICA8L2NlbS1wZi12Ni10b29sYmFyPlxuICAgICAgICAgICAgICAgICAgICA8Y2VtLXBmLXY2LWRyYXdlciBpZD1cImV2ZW50LWRyYXdlclwiIGV4cGFuZGVkPlxuICAgICAgICAgICAgICAgICAgICAgIDxkaXYgaWQ9XCJldmVudC1saXN0XCI+PC9kaXY+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWRldGFpbC1oZWFkZXJcIiBzbG90PVwicGFuZWwtaGVhZGVyXCI+PC9kaXY+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWRldGFpbC1ib2R5XCIgc2xvdD1cInBhbmVsLWJvZHlcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtZHJhd2VyPlxuICAgICAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdGFiPlxuICAgICAgICAgICAgICAgIDxjZW0tcGYtdjYtdGFiIHRpdGxlPVwiSGVhbHRoXCI+XG4gICAgICAgICAgICAgICAgICA8Y2VtLWhlYWx0aC1wYW5lbCBjb21wb25lbnQ9JHtpZkRlZmluZWQodGhpcy5wcmltYXJ5VGFnTmFtZSl9PlxuICAgICAgICAgICAgICAgICAgPC9jZW0taGVhbHRoLXBhbmVsPlxuICAgICAgICAgICAgICAgIDwvY2VtLXBmLXY2LXRhYj5cbiAgICAgICAgICAgICAgPC9jZW0tcGYtdjYtdGFicz5cbiAgICAgICAgICAgIDwvY2VtLWRyYXdlcj5cbiAgICAgICAgICA8L2Zvb3Rlcj5cbiAgICAgICAgPC9jZW0tcGYtdjYtcGFnZS1tYWluPlxuICAgICAgPC9jZW0tcGYtdjYtcGFnZT5cblxuICAgICAgPGNlbS1wZi12Ni1tb2RhbCBpZD1cImRlYnVnLW1vZGFsXCIgdmFyaWFudD1cImxhcmdlXCI+XG4gICAgICAgIDxoMiBzbG90PVwiaGVhZGVyXCI+RGVidWcgSW5mb3JtYXRpb248L2gyPlxuICAgICAgICA8ZGwgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0IHBmLW0taG9yaXpvbnRhbCBwZi1tLWNvbXBhY3RcIj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5TZXJ2ZXIgVmVyc2lvbjwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctdmVyc2lvblwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlNlcnZlciBPUzwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctb3NcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5XYXRjaCBEaXJlY3Rvcnk8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLXdhdGNoLWRpclwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPk1hbmlmZXN0IFNpemU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLW1hbmlmZXN0LXNpemVcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5EZW1vcyBGb3VuZDwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctZGVtby1jb3VudFwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlRhZyBOYW1lPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCI+JHt0aGlzLnByaW1hcnlUYWdOYW1lIHx8ICctJ308L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkRlbW8gVGl0bGU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIj4ke3RoaXMuZGVtb1RpdGxlIHx8ICctJ308L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkJyb3dzZXI8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLWJyb3dzZXJcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2VtLXBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5Vc2VyIEFnZW50PC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy11YVwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICA8L2RsPlxuICAgICAgICA8ZGl2IGlkPVwiZGVtby11cmxzLWNvbnRhaW5lclwiPjwvZGl2PlxuICAgICAgICA8Y2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbiBpZD1cImRlYnVnLWltcG9ydG1hcC1kZXRhaWxzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB0b2dnbGUtdGV4dD1cIlNob3cgSW1wb3J0IE1hcFwiPlxuICAgICAgICAgIDxwcmUgaWQ9XCJkZWJ1Zy1pbXBvcnRtYXBcIj4tPC9wcmU+XG4gICAgICAgIDwvY2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbj5cbiAgICAgICAgPGRpdiBzbG90PVwiZm9vdGVyXCIgY2xhc3M9XCJidXR0b24tY29udGFpbmVyXCI+XG4gICAgICAgICAgPGNlbS1wZi12Ni1idXR0b24gY2xhc3M9XCJkZWJ1Zy1jb3B5XCIgdmFyaWFudD1cInByaW1hcnlcIj5cbiAgICAgICAgICAgIENvcHkgRGVidWcgSW5mb1xuICAgICAgICAgIDwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBjbGFzcz1cImRlYnVnLWNsb3NlXCIgdmFyaWFudD1cInNlY29uZGFyeVwiPlxuICAgICAgICAgICAgQ2xvc2VcbiAgICAgICAgICA8L2NlbS1wZi12Ni1idXR0b24+XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9jZW0tcGYtdjYtbW9kYWw+XG5cbiAgICAgIDwhLS0gUmVjb25uZWN0aW9uIG1vZGFsIC0tPlxuICAgICAgPGNlbS1wZi12Ni1tb2RhbCBpZD1cInJlY29ubmVjdGlvbi1tb2RhbFwiIHZhcmlhbnQ9XCJsYXJnZVwiPlxuICAgICAgICA8aDIgc2xvdD1cImhlYWRlclwiPkRldmVsb3BtZW50IFNlcnZlciBEaXNjb25uZWN0ZWQ8L2gyPlxuICAgICAgICA8Y2VtLXJlY29ubmVjdGlvbi1jb250ZW50IGlkPVwicmVjb25uZWN0aW9uLWNvbnRlbnRcIj48L2NlbS1yZWNvbm5lY3Rpb24tY29udGVudD5cbiAgICAgICAgPGNlbS1wZi12Ni1idXR0b24gaWQ9XCJyZWxvYWQtYnV0dG9uXCJcbiAgICAgICAgICAgICAgICAgICAgICBzbG90PVwiZm9vdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwicHJpbWFyeVwiPlJlbG9hZCBQYWdlPC9jZW0tcGYtdjYtYnV0dG9uPlxuICAgICAgICA8Y2VtLXBmLXY2LWJ1dHRvbiBpZD1cInJldHJ5LWJ1dHRvblwiXG4gICAgICAgICAgICAgICAgICAgICAgc2xvdD1cImZvb3RlclwiXG4gICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInNlY29uZGFyeVwiPlJldHJ5IE5vdzwvY2VtLXBmLXY2LWJ1dHRvbj5cbiAgICAgIDwvY2VtLXBmLXY2LW1vZGFsPlxuXG4gICAgICA8IS0tIFRyYW5zZm9ybSBlcnJvciBvdmVybGF5IC0tPlxuICAgICAgPGNlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheSBpZD1cImVycm9yLW92ZXJsYXlcIj5cbiAgICAgIDwvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5PlxuICAgIGA7XG4gIH1cblxuICAjcmVuZGVyU291cmNlQnV0dG9uKCkge1xuICAgIGlmICghdGhpcy5zb3VyY2VVUkwpIHJldHVybiBub3RoaW5nO1xuXG4gICAgbGV0IGxhYmVsID0gJ1ZlcnNpb24gQ29udHJvbCc7XG4gICAgbGV0IHBhdGggPSAnTTUuODU0IDQuODU0YS41LjUgMCAxIDAtLjcwOC0uNzA4bC0zLjUgMy41YS41LjUgMCAwIDAgMCAuNzA4bDMuNSAzLjVhLjUuNSAwIDAgMCAuNzA4LS43MDhMMi43MDcgOGwzLjE0Ny0zLjE0NnptNC4yOTIgMGEuNS41IDAgMCAxIC43MDgtLjcwOGwzLjUgMy41YS41LjUgMCAwIDEgMCAuNzA4bC0zLjUgMy41YS41LjUgMCAwIDEtLjcwOC0uNzA4TDEzLjI5MyA4bC0zLjE0Ny0zLjE0NnonO1xuXG4gICAgaWYgKHRoaXMuc291cmNlVVJMLmluY2x1ZGVzKCdnaXRodWIuY29tJykpIHtcbiAgICAgIGxhYmVsID0gJ0dpdEh1Yi5jb20nO1xuICAgICAgcGF0aCA9ICdNOCAwQzMuNTggMCAwIDMuNTggMCA4YzAgMy41NCAyLjI5IDYuNTMgNS40NyA3LjU5LjQuMDcuNTUtLjE3LjU1LS4zOCAwLS4xOS0uMDEtLjgyLS4wMS0xLjQ5LTIuMDEuMzctMi41My0uNDktMi42OS0uOTQtLjA5LS4yMy0uNDgtLjk0LS44Mi0xLjEzLS4yOC0uMTUtLjY4LS41Mi0uMDEtLjUzLjYzLS4wMSAxLjA4LjU4IDEuMjMuODIuNzIgMS4yMSAxLjg3Ljg3IDIuMzMuNjYuMDctLjUyLjI4LS44Ny41MS0xLjA3LTEuNzgtLjItMy42NC0uODktMy42NC0zLjk1IDAtLjg3LjMxLTEuNTkuODItMi4xNS0uMDgtLjItLjM2LTEuMDIuMDgtMi4xMiAwIDAgLjY3LS4yMSAyLjIuODIuNjQtLjE4IDEuMzItLjI3IDItLjI3LjY4IDAgMS4zNi4wOSAyIC4yNyAxLjUzLTEuMDQgMi4yLS44MiAyLjItLjgyLjQ0IDEuMS4xNiAxLjkyLjA4IDIuMTIuNTEuNTYuODIgMS4yNy44MiAyLjE1IDAgMy4wNy0xLjg3IDMuNzUtMy42NSAzLjk1LjI5LjI1LjU0LjczLjU0IDEuNDggMCAxLjA3LS4wMSAxLjkzLS4wMSAyLjIgMCAuMjEuMTUuNDYuNTUuMzhBOC4wMTMgOC4wMTMgMCAwMDE2IDhjMC00LjQyLTMuNTgtOC04LTh6JztcbiAgICB9IGVsc2UgaWYgKHRoaXMuc291cmNlVVJMLmluY2x1ZGVzKCdnaXRsYWIuY29tJykpIHtcbiAgICAgIGxhYmVsID0gJ0dpdExhYic7XG4gICAgICBwYXRoID0gJ20xNS43MzQgNi4xLS4wMjItLjA1OEwxMy41MzQuMzU4YS41NjguNTY4IDAgMCAwLS41NjMtLjM1Ni41ODMuNTgzIDAgMCAwLS4zMjguMTIyLjU4Mi41ODIgMCAwIDAtLjE5My4yOTRsLTEuNDcgNC40OTlINS4wMjVsLTEuNDctNC41QS41NzIuNTcyIDAgMCAwIDMuMzYwLjE3NGEuNTcyLjU3MiAwIDAgMC0uMzI4LS4xNzIuNTgyLjU4MiAwIDAgMC0uNTYzLjM1N0wuMjkgNi4wNGwtLjAyMi4wNTdBNC4wNDQgNC4wNDQgMCAwIDAgMS42MSAxMC43N2wuMDA3LjAwNi4wMi4wMTQgMy4zMTggMi40ODUgMS42NCAxLjI0MiAxIC43NTVhLjY3My42NzMgMCAwIDAgLjgxNCAwbDEtLjc1NSAxLjY0LTEuMjQyIDMuMzM4LTIuNS4wMDktLjAwN2E0LjA1IDQuMDUgMCAwIDAgMS4zNC00LjY2OFonO1xuICAgIH0gZWxzZSBpZiAodGhpcy5zb3VyY2VVUkwuaW5jbHVkZXMoJ2JpdGJ1Y2tldC5vcmcnKSkge1xuICAgICAgbGFiZWwgPSAnQml0YnVja2V0JztcbiAgICAgIHBhdGggPSAnTTAgMS41QTEuNSAxLjUgMCAwIDEgMS41IDBoMTNBMS41IDEuNSAwIDAgMSAxNiAxLjV2MTNhMS41IDEuNSAwIDAgMS0xLjUgMS41aC0xM0ExLjUgMS41IDAgMCAxIDAgMTQuNXYtMTN6TTIuNSAzYS41LjUgMCAwIDAtLjUuNXY5YS41LjUgMCAwIDAgLjUuNWgxMWEuNS41IDAgMCAwIC41LS41di05YS41LjUgMCAwIDAtLjUtLjVoLTExem01LjAzOCAxLjQzNWEuNS41IDAgMCAxIC45MjQgMGwxLjQyIDMuMzdIOC43OGwtLjI0My42MDgtLjI0My0uNjA4SDUuMDgybDEuNDItMy4zN3pNOCA5LjE0M2wtLjc0MyAxLjg1N0g0Ljc0M0w2LjA3NiA3LjYwOCA4IDkuMTQzeic7XG4gICAgfVxuXG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8Y2VtLXBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgPGNlbS1wZi12Ni1idXR0b24gaHJlZj1cIiR7dGhpcy5zb3VyY2VVUkx9XCJcbiAgICAgICAgICAgICAgICAgICAgICB0YXJnZXQ9XCJfYmxhbmtcIlxuICAgICAgICAgICAgICAgICAgICAgIHJlbD1cIm5vb3BlbmVyIG5vcmVmZXJyZXJcIlxuICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIlZpZXcgc291cmNlIGZpbGVcIj5cbiAgICAgICAgICA8c3ZnIGFyaWEtbGFiZWw9XCIke2xhYmVsfVwiXG4gICAgICAgICAgICAgICB3aWR0aD1cIjFyZW1cIlxuICAgICAgICAgICAgICAgaGVpZ2h0PVwiMXJlbVwiXG4gICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIj5cbiAgICAgICAgICAgIDxwYXRoIGQ9XCIke3BhdGh9XCIvPlxuICAgICAgICAgIDwvc3ZnPlxuICAgICAgICA8L2NlbS1wZi12Ni1idXR0b24+XG4gICAgICA8L2NlbS1wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgYDtcbiAgfVxuXG4gIC8qKiBSZXNvbHZlcyB3aGVuIGNsaWVudC1vbmx5IG1vZHVsZXMgYXJlIGxvYWRlZC4gKi9cbiAgI21vZHVsZXNSZWFkeTogUHJvbWlzZTx2b2lkPiA9IHRoaXMuI2xvYWRDbGllbnRNb2R1bGVzKCk7XG5cbiAgI2xvYWRDbGllbnRNb2R1bGVzKCk6IFByb21pc2U8dm9pZD4ge1xuICAgIHJldHVybiBQcm9taXNlLmFsbChbXG4gICAgICAvLyBAdHMtaWdub3JlIC0tIHBsYWluIEpTIG1vZHVsZXMgc2VydmVkIGF0IHJ1bnRpbWUgYnkgR28gc2VydmVyXG4gICAgICBpbXBvcnQoJy9fX2NlbS93ZWJzb2NrZXQtY2xpZW50LmpzJyksXG4gICAgICAvLyBAdHMtaWdub3JlXG4gICAgICBpbXBvcnQoJy9fX2NlbS9zdGF0ZS1wZXJzaXN0ZW5jZS5qcycpLFxuICAgIF0pLnRoZW4oKFt3cywgc3BdKSA9PiB7XG4gICAgICBDRU1SZWxvYWRDbGllbnQgPSB3cy5DRU1SZWxvYWRDbGllbnQ7XG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlID0gc3AuU3RhdGVQZXJzaXN0ZW5jZTtcbiAgICAgIC8vIEB0cy1pZ25vcmVcbiAgICAgIGltcG9ydCgnL19fY2VtL2hlYWx0aC1iYWRnZXMuanMnKS5jYXRjaCgoZTogdW5rbm93bikgPT5cbiAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZV0gRmFpbGVkIHRvIGxvYWQgaGVhbHRoLWJhZGdlczonLCBlKSk7XG4gICAgfSkuY2F0Y2goKGU6IHVua25vd24pID0+IHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmVdIEZhaWxlZCB0byBsb2FkIGNsaWVudCBtb2R1bGVzOicsIGUpO1xuICAgICAgLy8gRGVncmFkZWQgZmFsbGJhY2tzIHNvIHRoZSBjb21wb25lbnQgc3RpbGwgcmVuZGVyc1xuICAgICAgQ0VNUmVsb2FkQ2xpZW50ID8/PSBjbGFzcyB7IGluaXQoKSB7fSByZXRyeSgpIHt9IGRlc3Ryb3koKSB7fSB9IGFzIGFueTtcbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UgPz89IHtcbiAgICAgICAgZ2V0U3RhdGU6ICgpID0+ICh7IGNvbG9yU2NoZW1lOiAnc3lzdGVtJyB9KSxcbiAgICAgICAgdXBkYXRlU3RhdGUoKSB7fSxcbiAgICAgICAgZ2V0VHJlZVN0YXRlOiAoKSA9PiAoeyBleHBhbmRlZDogW10sIHNlbGVjdGVkOiBudWxsIH0pLFxuICAgICAgICBzZXRUcmVlU3RhdGUoKSB7fSxcbiAgICAgICAgdXBkYXRlVHJlZVN0YXRlKCkge30sXG4gICAgICAgIG1pZ3JhdGVGcm9tTG9jYWxTdG9yYWdlKCkge30sXG4gICAgICB9O1xuICAgIH0pO1xuICB9XG5cbiAgYXN5bmMgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgLy8gTW9kdWxlcyBtdXN0IGxvYWQgYmVmb3JlIHN1cGVyIHNvIGZpcnN0VXBkYXRlZCgpIGNhbiB1c2UgU3RhdGVQZXJzaXN0ZW5jZVxuICAgIGF3YWl0IHRoaXMuI21vZHVsZXNSZWFkeTtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuXG4gICAgaWYgKHRoaXMuI3dzQ2xpZW50ID09IG51bGwpIHtcbiAgICAgIHRoaXMuI2luaXRXc0NsaWVudCgpO1xuICAgIH1cbiAgICB0aGlzLiNtaWdyYXRlRnJvbUxvY2FsU3RvcmFnZUlmTmVlZGVkKCk7XG4gIH1cblxuICBmaXJzdFVwZGF0ZWQoKSB7XG4gICAgLy8gU2V0IHVwIGRlYnVnIG92ZXJsYXlcbiAgICB0aGlzLiNzZXR1cERlYnVnT3ZlcmxheSgpO1xuXG4gICAgLy8gU2V0IHVwIGNvbG9yIHNjaGVtZSB0b2dnbGVcbiAgICB0aGlzLiNzZXR1cENvbG9yU2NoZW1lVG9nZ2xlKCk7XG5cbiAgICAvLyBTZXQgdXAgZm9vdGVyIGRyYXdlciBhbmQgdGFic1xuICAgIHRoaXMuI3NldHVwRm9vdGVyRHJhd2VyKCk7XG5cbiAgICAvLyBMaXN0ZW4gZm9yIHNlcnZlciBsb2cgbWVzc2FnZXMgZnJvbSBXZWJTb2NrZXRcbiAgICB0aGlzLiNzZXR1cExvZ0xpc3RlbmVyKCk7XG5cbiAgICAvLyBTZXQgdXAga25vYiBldmVudCBjb29yZGluYXRpb25cbiAgICB0aGlzLiNzZXR1cEtub2JDb29yZGluYXRpb24oKTtcblxuICAgIC8vIFNldCB1cCB0cmVlIHN0YXRlIHBlcnNpc3RlbmNlXG4gICAgdGhpcy4jc2V0dXBUcmVlU3RhdGVQZXJzaXN0ZW5jZSgpO1xuXG4gICAgLy8gU2V0IHVwIHNpZGViYXIgc3RhdGUgcGVyc2lzdGVuY2VcbiAgICB0aGlzLiNzZXR1cFNpZGViYXJTdGF0ZVBlcnNpc3RlbmNlKCk7XG5cbiAgICAvLyBTZXQgdXAgZWxlbWVudCBldmVudCBjYXB0dXJlXG4gICAgdGhpcy4jc2V0dXBFdmVudENhcHR1cmUoKS50aGVuKCgpID0+IHtcbiAgICAgIHRoaXMuI3NldHVwRXZlbnRMaXN0ZW5lcnMoKTtcbiAgICB9KTtcblxuICAgIC8vIFNldCB1cCByZWNvbm5lY3Rpb24gbW9kYWwgYnV0dG9uIGhhbmRsZXJzXG4gICAgLyogYzggaWdub3JlIHN0YXJ0IC0gd2luZG93LmxvY2F0aW9uLnJlbG9hZCB3b3VsZCByZWxvYWQgdGVzdCBwYWdlICovXG4gICAgdGhpcy4jJCgncmVsb2FkLWJ1dHRvbicpPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgIHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKTtcbiAgICB9KTtcbiAgICAvKiBjOCBpZ25vcmUgc3RvcCAqL1xuXG4gICAgdGhpcy4jJCgncmV0cnktYnV0dG9uJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LmNsb3NlKCk7XG4gICAgICB0aGlzLiN3c0NsaWVudC5yZXRyeSgpO1xuICAgIH0pO1xuXG4gICAgLy8gSW5pdGlhbGl6ZSBXZWJTb2NrZXQgY29ubmVjdGlvblxuICAgIHRoaXMuI3dzQ2xpZW50LmluaXQoKTtcbiAgfVxuXG4gIGFzeW5jICNmZXRjaERlYnVnSW5mbygpIHtcbiAgICAvLyBQb3B1bGF0ZSBicm93c2VyIGluZm8gaW1tZWRpYXRlbHlcbiAgICBjb25zdCBicm93c2VyRWwgPSB0aGlzLiMkKCdkZWJ1Zy1icm93c2VyJyk7XG4gICAgY29uc3QgdWFFbCA9IHRoaXMuIyQoJ2RlYnVnLXVhJyk7XG4gICAgaWYgKGJyb3dzZXJFbCkge1xuICAgICAgY29uc3QgYnJvd3NlciA9IHRoaXMuI2RldGVjdEJyb3dzZXIoKTtcbiAgICAgIGJyb3dzZXJFbC50ZXh0Q29udGVudCA9IGJyb3dzZXI7XG4gICAgfVxuICAgIGlmICh1YUVsKSB7XG4gICAgICB1YUVsLnRleHRDb250ZW50ID0gbmF2aWdhdG9yLnVzZXJBZ2VudDtcbiAgICB9XG5cbiAgICAvLyBGZXRjaCBzZXJ2ZXIgZGVidWcgaW5mb1xuICAgIGNvbnN0IGRhdGEgPSBhd2FpdCBmZXRjaCgnL19fY2VtL2RlYnVnJylcbiAgICAgIC50aGVuKHJlcyA9PiByZXMuanNvbigpKVxuICAgICAgLmNhdGNoKChlcnI6IEVycm9yKSA9PiB7XG4gICAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gZmV0Y2ggZGVidWcgaW5mbzonLCBlcnIpO1xuICAgICAgfSk7XG5cbiAgICBpZiAoIWRhdGEpIHJldHVybjtcbiAgICBjb25zdCB2ZXJzaW9uRWwgPSB0aGlzLiMkKCdkZWJ1Zy12ZXJzaW9uJyk7XG4gICAgY29uc3Qgb3NFbCA9IHRoaXMuIyQoJ2RlYnVnLW9zJyk7XG4gICAgY29uc3Qgd2F0Y2hEaXJFbCA9IHRoaXMuIyQoJ2RlYnVnLXdhdGNoLWRpcicpO1xuICAgIGNvbnN0IG1hbmlmZXN0U2l6ZUVsID0gdGhpcy4jJCgnZGVidWctbWFuaWZlc3Qtc2l6ZScpO1xuICAgIGNvbnN0IGRlbW9Db3VudEVsID0gdGhpcy4jJCgnZGVidWctZGVtby1jb3VudCcpO1xuICAgIGNvbnN0IGRlbW9VcmxzQ29udGFpbmVyID0gdGhpcy4jJCgnZGVtby11cmxzLWNvbnRhaW5lcicpO1xuICAgIGNvbnN0IGltcG9ydE1hcEVsID0gdGhpcy4jJCgnZGVidWctaW1wb3J0bWFwJyk7XG5cbiAgICBpZiAodmVyc2lvbkVsKSB2ZXJzaW9uRWwudGV4dENvbnRlbnQgPSBkYXRhLnZlcnNpb24gfHwgJy0nO1xuICAgIGlmIChvc0VsKSBvc0VsLnRleHRDb250ZW50ID0gZGF0YS5vcyB8fCAnLSc7XG4gICAgaWYgKHdhdGNoRGlyRWwpIHdhdGNoRGlyRWwudGV4dENvbnRlbnQgPSBkYXRhLndhdGNoRGlyIHx8ICctJztcbiAgICBpZiAobWFuaWZlc3RTaXplRWwpIG1hbmlmZXN0U2l6ZUVsLnRleHRDb250ZW50ID0gZGF0YS5tYW5pZmVzdFNpemUgfHwgJy0nO1xuICAgIGlmIChkZW1vQ291bnRFbCkgZGVtb0NvdW50RWwudGV4dENvbnRlbnQgPSBkYXRhLmRlbW9Db3VudCB8fCAnMCc7XG5cbiAgICBpZiAoZGVtb1VybHNDb250YWluZXIpIHtcbiAgICAgIHRoaXMuI3BvcHVsYXRlRGVtb1VybHMoZGVtb1VybHNDb250YWluZXIsIGRhdGEuZGVtb3MpO1xuICAgIH1cblxuICAgIGlmIChpbXBvcnRNYXBFbCAmJiBkYXRhLmltcG9ydE1hcCkge1xuICAgICAgY29uc3QgaW1wb3J0TWFwSlNPTiA9IEpTT04uc3RyaW5naWZ5KGRhdGEuaW1wb3J0TWFwLCBudWxsLCAyKTtcbiAgICAgIGltcG9ydE1hcEVsLnRleHRDb250ZW50ID0gaW1wb3J0TWFwSlNPTjtcbiAgICAgIGRhdGEuaW1wb3J0TWFwSlNPTiA9IGltcG9ydE1hcEpTT047XG4gICAgfSBlbHNlIGlmIChpbXBvcnRNYXBFbCkge1xuICAgICAgaW1wb3J0TWFwRWwudGV4dENvbnRlbnQgPSAnTm8gaW1wb3J0IG1hcCBnZW5lcmF0ZWQnO1xuICAgIH1cblxuICAgIHRoaXMuI2RlYnVnRGF0YSA9IGRhdGE7XG4gIH1cblxuICAjcG9wdWxhdGVEZW1vVXJscyhjb250YWluZXI6IEhUTUxFbGVtZW50LCBkZW1vczogRGVidWdEYXRhWydkZW1vcyddKSB7XG4gICAgaWYgKCFkZW1vcz8ubGVuZ3RoKSB7XG4gICAgICBjb250YWluZXIudGV4dENvbnRlbnQgPSAnTm8gZGVtb3MgZm91bmQgaW4gbWFuaWZlc3QnO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IGN1cnJlbnRUYWdOYW1lID0gdGhpcy5wcmltYXJ5VGFnTmFtZSB8fCAnJztcbiAgICBjb25zdCBpc09uRGVtb1BhZ2UgPSAhIWN1cnJlbnRUYWdOYW1lO1xuXG4gICAgaWYgKGlzT25EZW1vUGFnZSkge1xuICAgICAgY29uc3QgY3VycmVudERlbW8gPSBkZW1vcy5maW5kKGRlbW8gPT4gZGVtby50YWdOYW1lID09PSBjdXJyZW50VGFnTmFtZSk7XG4gICAgICBpZiAoIWN1cnJlbnREZW1vKSB7XG4gICAgICAgIGNvbnRhaW5lci50ZXh0Q29udGVudCA9ICdDdXJyZW50IGRlbW8gbm90IGZvdW5kIGluIG1hbmlmZXN0JztcbiAgICAgICAgcmV0dXJuO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBmcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vSW5mb1RlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGFnTmFtZVwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLnRhZ05hbWU7XG4gICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNhbm9uaWNhbFVSTFwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmNhbm9uaWNhbFVSTDtcbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibG9jYWxSb3V0ZVwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmxvY2FsUm91dGU7XG5cbiAgICAgIGNvbnN0IGRlc2NyaXB0aW9uR3JvdXAgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImRlc2NyaXB0aW9uXCJdJyk7XG4gICAgICBpZiAoY3VycmVudERlbW8uZGVzY3JpcHRpb24pIHtcbiAgICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJkZXNjcmlwdGlvblwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmRlc2NyaXB0aW9uO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgZGVzY3JpcHRpb25Hcm91cD8ucmVtb3ZlKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGZpbGVwYXRoR3JvdXAgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImZpbGVwYXRoXCJdJyk7XG4gICAgICBpZiAoY3VycmVudERlbW8uZmlsZXBhdGgpIHtcbiAgICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJmaWxlcGF0aFwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmZpbGVwYXRoO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgZmlsZXBhdGhHcm91cD8ucmVtb3ZlKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnRhaW5lci5yZXBsYWNlQ2hpbGRyZW4oZnJhZ21lbnQpO1xuICAgIH0gZWxzZSB7XG4gICAgICBjb25zdCBsaXN0RnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZGVtb0xpc3RUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICBjb25zdCBncm91cHNDb250YWluZXIgPSBsaXN0RnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtY29udGFpbmVyPVwiZ3JvdXBzXCJdJykhO1xuXG4gICAgICBmb3IgKGNvbnN0IGRlbW8gb2YgZGVtb3MpIHtcbiAgICAgICAgY29uc3QgZ3JvdXBGcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vR3JvdXBUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0YWdOYW1lXCJdJykhLnRleHRDb250ZW50ID0gZGVtby50YWdOYW1lO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIl0nKSEudGV4dENvbnRlbnQgPVxuICAgICAgICAgIGRlbW8uZGVzY3JpcHRpb24gfHwgJyhubyBkZXNjcmlwdGlvbiknO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5jYW5vbmljYWxVUkw7XG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5sb2NhbFJvdXRlO1xuXG4gICAgICAgIGNvbnN0IGZpbGVwYXRoR3JvdXAgPSBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIl0nKTtcbiAgICAgICAgaWYgKGRlbW8uZmlsZXBhdGgpIHtcbiAgICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIl0nKSEudGV4dENvbnRlbnQgPSBkZW1vLmZpbGVwYXRoO1xuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIGZpbGVwYXRoR3JvdXA/LnJlbW92ZSgpO1xuICAgICAgICB9XG5cbiAgICAgICAgZ3JvdXBzQ29udGFpbmVyLmFwcGVuZENoaWxkKGdyb3VwRnJhZ21lbnQpO1xuICAgICAgfVxuXG4gICAgICBjb250YWluZXIucmVwbGFjZUNoaWxkcmVuKGxpc3RGcmFnbWVudCk7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwTG9nTGlzdGVuZXIoKSB7XG4gICAgdGhpcy4jbG9nQ29udGFpbmVyID0gdGhpcy4jJCgnbG9nLWNvbnRhaW5lcicpO1xuXG4gICAgY29uc3QgbG9nc0ZpbHRlciA9IHRoaXMuIyQoJ2xvZ3MtZmlsdGVyJykgYXMgSFRNTElucHV0RWxlbWVudCB8IG51bGw7XG4gICAgaWYgKGxvZ3NGaWx0ZXIpIHtcbiAgICAgIGxvZ3NGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignaW5wdXQnLCAoKSA9PiB7XG4gICAgICAgIGNvbnN0IHsgdmFsdWUgPSAnJyB9ID0gbG9nc0ZpbHRlcjtcbiAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyISk7XG4gICAgICAgIHRoaXMuI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgdGhpcy4jZmlsdGVyTG9ncyh2YWx1ZSk7XG4gICAgICAgIH0sIDMwMCk7XG4gICAgICB9KTtcbiAgICB9XG5cbiAgICB0aGlzLiNsb2dMZXZlbERyb3Bkb3duID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcjbG9nLWxldmVsLWZpbHRlcicpID8/IG51bGw7XG4gICAgaWYgKHRoaXMuI2xvZ0xldmVsRHJvcGRvd24pIHtcbiAgICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICAgIHRoaXMuI2xvYWRMb2dGaWx0ZXJTdGF0ZSgpO1xuICAgICAgfSk7XG4gICAgICB0aGlzLiNsb2dMZXZlbERyb3Bkb3duLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZUxvZ0ZpbHRlckNoYW5nZSBhcyBFdmVudExpc3RlbmVyKTtcbiAgICB9XG5cbiAgICB0aGlzLiMkKCdjb3B5LWxvZ3MnKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB0aGlzLiNjb3B5TG9ncygpO1xuICAgIH0pO1xuXG4gICAgdGhpcy4jaGFuZGxlTG9nc0V2ZW50ID0gKChldmVudDogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IGxvZ3MgPSAoZXZlbnQgYXMgQ2VtTG9nc0V2ZW50KS5sb2dzO1xuICAgICAgaWYgKGxvZ3MpIHtcbiAgICAgICAgdGhpcy4jcmVuZGVyTG9ncyhsb2dzKTtcbiAgICAgIH1cbiAgICB9KTtcbiAgICB3aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcignY2VtOmxvZ3MnLCB0aGlzLiNoYW5kbGVMb2dzRXZlbnQpO1xuICB9XG5cbiAgI2ZpbHRlckxvZ3MocXVlcnk6IHN0cmluZykge1xuICAgIHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgZm9yIChjb25zdCBlbnRyeSBvZiB0aGlzLiNsb2dDb250YWluZXIuY2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IHRleHQgPSBlbnRyeS50ZXh0Q29udGVudD8udG9Mb3dlckNhc2UoKSA/PyAnJztcbiAgICAgIGNvbnN0IHRleHRNYXRjaCA9ICF0aGlzLiNsb2dzRmlsdGVyVmFsdWUgfHwgdGV4dC5pbmNsdWRlcyh0aGlzLiNsb2dzRmlsdGVyVmFsdWUpO1xuXG4gICAgICBjb25zdCBsb2dUeXBlID0gdGhpcy4jZ2V0TG9nVHlwZUZyb21FbnRyeShlbnRyeSk7XG4gICAgICBjb25zdCBsZXZlbE1hdGNoID0gdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmhhcyhsb2dUeXBlKTtcblxuICAgICAgKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5oaWRkZW4gPSAhKHRleHRNYXRjaCAmJiBsZXZlbE1hdGNoKTtcbiAgICB9XG4gIH1cblxuICAjZ2V0TG9nVHlwZUZyb21FbnRyeShlbnRyeTogRWxlbWVudCk6IHN0cmluZyB7XG4gICAgZm9yIChjb25zdCBjbHMgb2YgZW50cnkuY2xhc3NMaXN0KSB7XG4gICAgICBpZiAoWydpbmZvJywgJ3dhcm5pbmcnLCAnZXJyb3InLCAnZGVidWcnXS5pbmNsdWRlcyhjbHMpKSB7XG4gICAgICAgIHJldHVybiBjbHMgPT09ICd3YXJuaW5nJyA/ICd3YXJuJyA6IGNscztcbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuICdpbmZvJztcbiAgfVxuXG4gICNsb2FkTG9nRmlsdGVyU3RhdGUoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHNhdmVkID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1sb2ctZmlsdGVycycpO1xuICAgICAgaWYgKHNhdmVkKSB7XG4gICAgICAgIHRoaXMuI2xvZ0xldmVsRmlsdGVycyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZCkpO1xuICAgICAgfVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIGNvbnNvbGUuZGVidWcoJ1tjZW0tc2VydmUtY2hyb21lXSBsb2NhbFN0b3JhZ2UgdW5hdmFpbGFibGUsIHVzaW5nIGRlZmF1bHQgbG9nIGZpbHRlcnMnKTtcbiAgICB9XG4gICAgdGhpcy4jc3luY0NoZWNrYm94U3RhdGVzKCk7XG4gIH1cblxuICAjc3luY0NoZWNrYm94U3RhdGVzKCkge1xuICAgIGlmICghdGhpcy4jbG9nTGV2ZWxEcm9wZG93bikgcmV0dXJuO1xuICAgIGNvbnN0IG1lbnVJdGVtcyA9IHRoaXMuI2xvZ0xldmVsRHJvcGRvd24ucXVlcnlTZWxlY3RvckFsbCgnY2VtLXBmLXY2LW1lbnUtaXRlbScpO1xuICAgIG1lbnVJdGVtcy5mb3JFYWNoKGl0ZW0gPT4ge1xuICAgICAgY29uc3QgdmFsdWUgPSAoaXRlbSBhcyBhbnkpLnZhbHVlO1xuICAgICAgKGl0ZW0gYXMgYW55KS5jaGVja2VkID0gdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmhhcyh2YWx1ZSk7XG4gICAgfSk7XG4gIH1cblxuICAjc2F2ZUxvZ0ZpbHRlclN0YXRlKCkge1xuICAgIHRyeSB7XG4gICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLWxvZy1maWx0ZXJzJyxcbiAgICAgICAgSlNPTi5zdHJpbmdpZnkoWy4uLnRoaXMuI2xvZ0xldmVsRmlsdGVyc10pKTtcbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICAvLyBsb2NhbFN0b3JhZ2UgdW5hdmFpbGFibGUgKHByaXZhdGUgbW9kZSksIHNpbGVudGx5IGNvbnRpbnVlXG4gICAgfVxuICB9XG5cbiAgI2hhbmRsZUxvZ0ZpbHRlckNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB7IHZhbHVlLCBjaGVja2VkIH0gPSBldmVudCBhcyBFdmVudCAmIHsgdmFsdWU6IHN0cmluZzsgY2hlY2tlZDogYm9vbGVhbiB9O1xuXG4gICAgaWYgKGNoZWNrZWQpIHtcbiAgICAgIHRoaXMuI2xvZ0xldmVsRmlsdGVycy5hZGQodmFsdWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuZGVsZXRlKHZhbHVlKTtcbiAgICB9XG5cbiAgICB0aGlzLiNzYXZlTG9nRmlsdGVyU3RhdGUoKTtcbiAgICB0aGlzLiNmaWx0ZXJMb2dzKHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSk7XG4gIH07XG5cbiAgYXN5bmMgI2NvcHlMb2dzKCkge1xuICAgIGlmICghdGhpcy4jbG9nQ29udGFpbmVyKSByZXR1cm47XG5cbiAgICBjb25zdCBsb2dzID0gQXJyYXkuZnJvbSh0aGlzLiNsb2dDb250YWluZXIuY2hpbGRyZW4pXG4gICAgICAuZmlsdGVyKGVudHJ5ID0+ICEoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmhpZGRlbilcbiAgICAgIC5tYXAoZW50cnkgPT4ge1xuICAgICAgICBjb25zdCB0eXBlID0gZW50cnkucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsYWJlbFwiXScpPy50ZXh0Q29udGVudD8udHJpbSgpIHx8ICdJTkZPJztcbiAgICAgICAgY29uc3QgdGltZSA9IGVudHJ5LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGltZVwiXScpPy50ZXh0Q29udGVudD8udHJpbSgpIHx8ICcnO1xuICAgICAgICBjb25zdCBtZXNzYWdlID0gZW50cnkucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJtZXNzYWdlXCJdJyk/LnRleHRDb250ZW50Py50cmltKCkgfHwgJyc7XG4gICAgICAgIHJldHVybiBgWyR7dHlwZX1dICR7dGltZX0gJHttZXNzYWdlfWA7XG4gICAgICB9KS5qb2luKCdcXG4nKTtcblxuICAgIGlmICghbG9ncykgcmV0dXJuO1xuXG4gICAgdHJ5IHtcbiAgICAgIGF3YWl0IG5hdmlnYXRvci5jbGlwYm9hcmQud3JpdGVUZXh0KGxvZ3MpO1xuICAgICAgY29uc3QgYnRuID0gdGhpcy4jJCgnY29weS1sb2dzJyk7XG4gICAgICBpZiAoYnRuKSB7XG4gICAgICAgIGNvbnN0IHRleHROb2RlID0gQXJyYXkuZnJvbShidG4uY2hpbGROb2RlcykuZmluZChcbiAgICAgICAgICBuID0+IG4ubm9kZVR5cGUgPT09IE5vZGUuVEVYVF9OT0RFICYmIChuLnRleHRDb250ZW50Py50cmltKCkubGVuZ3RoID8/IDApID4gMFxuICAgICAgICApO1xuICAgICAgICBpZiAodGV4dE5vZGUpIHtcbiAgICAgICAgICBjb25zdCBvcmlnaW5hbCA9IHRleHROb2RlLnRleHRDb250ZW50O1xuICAgICAgICAgIHRleHROb2RlLnRleHRDb250ZW50ID0gJ0NvcGllZCEnO1xuXG4gICAgICAgICAgaWYgKHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgICBpZiAodGhpcy5pc0Nvbm5lY3RlZCAmJiB0ZXh0Tm9kZS5wYXJlbnROb2RlKSB7XG4gICAgICAgICAgICAgIHRleHROb2RlLnRleHRDb250ZW50ID0gb3JpZ2luYWw7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgICB0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgICAgICAgfSwgMjAwMCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9IGNhdGNoIChlcnIpIHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY29weSBsb2dzOicsIGVycik7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwRGVidWdPdmVybGF5KCkge1xuICAgIGNvbnN0IGRlYnVnQnV0dG9uID0gdGhpcy4jJCgnZGVidWctaW5mbycpO1xuICAgIGNvbnN0IGRlYnVnTW9kYWwgPSB0aGlzLiMkKCdkZWJ1Zy1tb2RhbCcpO1xuICAgIGNvbnN0IGRlYnVnQ2xvc2UgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJy5kZWJ1Zy1jbG9zZScpO1xuICAgIGNvbnN0IGRlYnVnQ29weSA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignLmRlYnVnLWNvcHknKTtcblxuICAgIGlmIChkZWJ1Z0J1dHRvbiAmJiBkZWJ1Z01vZGFsKSB7XG4gICAgICBkZWJ1Z0J1dHRvbi5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgICAgdGhpcy4jZmV0Y2hEZWJ1Z0luZm8oKTtcbiAgICAgICAgKGRlYnVnTW9kYWwgYXMgYW55KS5zaG93TW9kYWwoKTtcbiAgICAgIH0pO1xuXG4gICAgICBkZWJ1Z0Nsb3NlPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IChkZWJ1Z01vZGFsIGFzIGFueSkuY2xvc2UoKSk7XG5cbiAgICAgIGRlYnVnQ29weT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICAgIHRoaXMuI2NvcHlEZWJ1Z0luZm8oKTtcbiAgICAgIH0pO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cEZvb3RlckRyYXdlcigpIHtcbiAgICBjb25zdCBkcmF3ZXIgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ2NlbS1kcmF3ZXInKTtcbiAgICBjb25zdCB0YWJzID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdjZW0tcGYtdjYtdGFicycpO1xuXG4gICAgaWYgKCFkcmF3ZXIgfHwgIXRhYnMpIHJldHVybjtcblxuICAgIHRoaXMuI2RyYXdlck9wZW4gPSAoZHJhd2VyIGFzIGFueSkub3BlbjtcblxuICAgIGRyYXdlci5hZGRFdmVudExpc3RlbmVyKCdjaGFuZ2UnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgIHRoaXMuI2RyYXdlck9wZW4gPSAoZSBhcyBhbnkpLm9wZW47XG5cbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoe1xuICAgICAgICBkcmF3ZXI6IHsgb3BlbjogKGUgYXMgYW55KS5vcGVuIH1cbiAgICAgIH0pO1xuXG4gICAgICBpZiAoKGUgYXMgYW55KS5vcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExvZ3NUb0JvdHRvbSgpO1xuICAgICAgfVxuICAgIH0pO1xuXG4gICAgZHJhd2VyLmFkZEV2ZW50TGlzdGVuZXIoJ3Jlc2l6ZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgKGRyYXdlciBhcyBhbnkpLnNldEF0dHJpYnV0ZSgnZHJhd2VyLWhlaWdodCcsIChlIGFzIGFueSkuaGVpZ2h0KTtcblxuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIGRyYXdlcjogeyBoZWlnaHQ6IChlIGFzIGFueSkuaGVpZ2h0IH1cbiAgICAgIH0pO1xuICAgIH0pO1xuXG4gICAgdGFicy5hZGRFdmVudExpc3RlbmVyKCdjaGFuZ2UnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoe1xuICAgICAgICB0YWJzOiB7IHNlbGVjdGVkSW5kZXg6IChlIGFzIGFueSkuc2VsZWN0ZWRJbmRleCB9XG4gICAgICB9KTtcblxuICAgICAgaWYgKChlIGFzIGFueSkuc2VsZWN0ZWRJbmRleCA9PT0gMiAmJiAoZHJhd2VyIGFzIGFueSkub3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxMb2dzVG9Cb3R0b20oKTtcbiAgICAgIH1cblxuICAgICAgaWYgKChlIGFzIGFueSkuc2VsZWN0ZWRJbmRleCA9PT0gMyAmJiAoZHJhd2VyIGFzIGFueSkub3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxFdmVudHNUb0JvdHRvbSgpO1xuICAgICAgfVxuICAgIH0pO1xuICB9XG5cbiAgI2RldGVjdEJyb3dzZXIoKTogc3RyaW5nIHtcbiAgICBjb25zdCB1YSA9IG5hdmlnYXRvci51c2VyQWdlbnQ7XG4gICAgaWYgKHVhLmluY2x1ZGVzKCdGaXJlZm94LycpKSB7XG4gICAgICBjb25zdCBtYXRjaCA9IHVhLm1hdGNoKC9GaXJlZm94XFwvKFxcZCspLyk7XG4gICAgICByZXR1cm4gbWF0Y2ggPyBgRmlyZWZveCAke21hdGNoWzFdfWAgOiAnRmlyZWZveCc7XG4gICAgfSBlbHNlIGlmICh1YS5pbmNsdWRlcygnRWRnLycpKSB7XG4gICAgICBjb25zdCBtYXRjaCA9IHVhLm1hdGNoKC9FZGdcXC8oXFxkKykvKTtcbiAgICAgIHJldHVybiBtYXRjaCA/IGBFZGdlICR7bWF0Y2hbMV19YCA6ICdFZGdlJztcbiAgICB9IGVsc2UgaWYgKHVhLmluY2x1ZGVzKCdDaHJvbWUvJykpIHtcbiAgICAgIGNvbnN0IG1hdGNoID0gdWEubWF0Y2goL0Nocm9tZVxcLyhcXGQrKS8pO1xuICAgICAgcmV0dXJuIG1hdGNoID8gYENocm9tZSAke21hdGNoWzFdfWAgOiAnQ2hyb21lJztcbiAgICB9IGVsc2UgaWYgKHVhLmluY2x1ZGVzKCdTYWZhcmkvJykgJiYgIXVhLmluY2x1ZGVzKCdDaHJvbWUnKSkge1xuICAgICAgY29uc3QgbWF0Y2ggPSB1YS5tYXRjaCgvVmVyc2lvblxcLyhcXGQrKS8pO1xuICAgICAgcmV0dXJuIG1hdGNoID8gYFNhZmFyaSAke21hdGNoWzFdfWAgOiAnU2FmYXJpJztcbiAgICB9XG4gICAgcmV0dXJuICdVbmtub3duJztcbiAgfVxuXG4gIGFzeW5jICNjb3B5RGVidWdJbmZvKCkge1xuICAgIGNvbnN0IGluZm8gPSBBcnJheS5mcm9tKHRoaXMuIyQkKCcjZGVidWctbW9kYWwgZGwgZHQnKSkubWFwKGR0ID0+IHtcbiAgICAgIGNvbnN0IGRkID0gZHQubmV4dEVsZW1lbnRTaWJsaW5nO1xuICAgICAgaWYgKGRkICYmIGRkLnRhZ05hbWUgPT09ICdERCcpIHtcbiAgICAgICAgcmV0dXJuIGAke2R0LnRleHRDb250ZW50fTogJHtkZC50ZXh0Q29udGVudH1gO1xuICAgICAgfVxuICAgICAgcmV0dXJuICcnO1xuICAgIH0pLmpvaW4oJ1xcbicpO1xuXG4gICAgbGV0IGltcG9ydE1hcFNlY3Rpb24gPSAnJztcbiAgICBpZiAodGhpcy4jZGVidWdEYXRhPy5pbXBvcnRNYXBKU09OKSB7XG4gICAgICBpbXBvcnRNYXBTZWN0aW9uID0gYFxcbiR7Jz0nLnJlcGVhdCg0MCl9XFxuSW1wb3J0IE1hcDpcXG4keyc9Jy5yZXBlYXQoNDApfVxcbiR7dGhpcy4jZGVidWdEYXRhLmltcG9ydE1hcEpTT059XFxuYDtcbiAgICB9XG5cbiAgICBjb25zdCBkZWJ1Z1RleHQgPSBgQ0VNIFNlcnZlIERlYnVnIEluZm9ybWF0aW9uXG4keyc9Jy5yZXBlYXQoNDApfVxuJHtpbmZvfSR7aW1wb3J0TWFwU2VjdGlvbn1cbiR7Jz0nLnJlcGVhdCg0MCl9XG5HZW5lcmF0ZWQ6ICR7bmV3IERhdGUoKS50b0lTT1N0cmluZygpfWA7XG5cbiAgICB0cnkge1xuICAgICAgYXdhaXQgbmF2aWdhdG9yLmNsaXBib2FyZC53cml0ZVRleHQoZGVidWdUZXh0KTtcbiAgICAgIGNvbnN0IGNvcHlCdXR0b24gPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJy5kZWJ1Zy1jb3B5Jyk7XG4gICAgICBpZiAoY29weUJ1dHRvbikge1xuICAgICAgICBjb25zdCBvcmlnaW5hbFRleHQgPSBjb3B5QnV0dG9uLnRleHRDb250ZW50O1xuICAgICAgICBjb3B5QnV0dG9uLnRleHRDb250ZW50ID0gJ0NvcGllZCEnO1xuXG4gICAgICAgIGlmICh0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgICAgfVxuXG4gICAgICAgIHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgIGlmICh0aGlzLmlzQ29ubmVjdGVkICYmIGNvcHlCdXR0b24ucGFyZW50Tm9kZSkge1xuICAgICAgICAgICAgY29weUJ1dHRvbi50ZXh0Q29udGVudCA9IG9yaWdpbmFsVGV4dDtcbiAgICAgICAgICB9XG4gICAgICAgICAgdGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICAgICAgfSwgMjAwMCk7XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZXJyKSB7XG4gICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGNvcHkgZGVidWcgaW5mbzonLCBlcnIpO1xuICAgIH1cbiAgfVxuXG4gICNyZW5kZXJMb2dzKGxvZ3M6IEFycmF5PHsgdHlwZTogc3RyaW5nOyBkYXRlOiBzdHJpbmc7IG1lc3NhZ2U6IHN0cmluZyB9Pikge1xuICAgIGlmICghdGhpcy4jbG9nQ29udGFpbmVyKSByZXR1cm47XG5cbiAgICBjb25zdCBsb2dFbGVtZW50cyA9IGxvZ3MubWFwKGxvZyA9PiB7XG4gICAgICBjb25zdCBmcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNsb2dFbnRyeVRlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICAgIGNvbnN0IGRhdGUgPSBuZXcgRGF0ZShsb2cuZGF0ZSk7XG4gICAgICBjb25zdCB0aW1lID0gZGF0ZS50b0xvY2FsZVRpbWVTdHJpbmcoKTtcblxuICAgICAgY29uc3QgY29udGFpbmVyID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJjb250YWluZXJcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICAgIGNvbnRhaW5lci5jbGFzc0xpc3QuYWRkKGxvZy50eXBlKTtcbiAgICAgIGNvbnRhaW5lci5zZXRBdHRyaWJ1dGUoJ2RhdGEtbG9nLWlkJywgbG9nLmRhdGUpO1xuXG4gICAgICBjb25zdCB0eXBlTGFiZWwgPSB0aGlzLiNnZXRMb2dCYWRnZShsb2cudHlwZSk7XG4gICAgICBjb25zdCBzZWFyY2hUZXh0ID0gYCR7dHlwZUxhYmVsfSAke3RpbWV9ICR7bG9nLm1lc3NhZ2V9YC50b0xvd2VyQ2FzZSgpO1xuICAgICAgY29uc3QgdGV4dE1hdGNoID0gIXRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSB8fCBzZWFyY2hUZXh0LmluY2x1ZGVzKHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSk7XG5cbiAgICAgIGNvbnN0IGxvZ1R5cGVGb3JGaWx0ZXIgPSBsb2cudHlwZSA9PT0gJ3dhcm5pbmcnID8gJ3dhcm4nIDogbG9nLnR5cGU7XG4gICAgICBjb25zdCBsZXZlbE1hdGNoID0gdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmhhcyhsb2dUeXBlRm9yRmlsdGVyKTtcblxuICAgICAgaWYgKCEodGV4dE1hdGNoICYmIGxldmVsTWF0Y2gpKSB7XG4gICAgICAgIGNvbnRhaW5lci5zZXRBdHRyaWJ1dGUoJ2hpZGRlbicsICcnKTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgbGFiZWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxhYmVsXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICBsYWJlbC50ZXh0Q29udGVudCA9IHRoaXMuI2dldExvZ0JhZGdlKGxvZy50eXBlKTtcbiAgICAgIHRoaXMuI2FwcGx5TG9nTGFiZWxBdHRycyhsYWJlbCwgbG9nLnR5cGUpO1xuXG4gICAgICBjb25zdCB0aW1lRWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cInRpbWVcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICAgIHRpbWVFbC5zZXRBdHRyaWJ1dGUoJ2RhdGV0aW1lJywgbG9nLmRhdGUpO1xuICAgICAgdGltZUVsLnRleHRDb250ZW50ID0gdGltZTtcblxuICAgICAgKGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibWVzc2FnZVwiXScpIGFzIEhUTUxFbGVtZW50KS50ZXh0Q29udGVudCA9IGxvZy5tZXNzYWdlO1xuXG4gICAgICByZXR1cm4gZnJhZ21lbnQ7XG4gICAgfSk7XG5cbiAgICBpZiAoIXRoaXMuI2luaXRpYWxMb2dzRmV0Y2hlZCkge1xuICAgICAgdGhpcy4jbG9nQ29udGFpbmVyLnJlcGxhY2VDaGlsZHJlbiguLi5sb2dFbGVtZW50cyk7XG4gICAgICB0aGlzLiNpbml0aWFsTG9nc0ZldGNoZWQgPSB0cnVlO1xuXG4gICAgICBpZiAodGhpcy4jZHJhd2VyT3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICB0aGlzLiNsb2dDb250YWluZXIuYXBwZW5kKC4uLmxvZ0VsZW1lbnRzKTtcblxuICAgICAgaWYgKHRoaXMuI2RyYXdlck9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjZ2V0TG9nQmFkZ2UodHlwZTogc3RyaW5nKTogc3RyaW5nIHtcbiAgICBzd2l0Y2ggKHR5cGUpIHtcbiAgICAgIGNhc2UgJ2luZm8nOiByZXR1cm4gJ0luZm8nO1xuICAgICAgY2FzZSAnd2FybmluZyc6IHJldHVybiAnV2Fybic7XG4gICAgICBjYXNlICdlcnJvcic6IHJldHVybiAnRXJyb3InO1xuICAgICAgY2FzZSAnZGVidWcnOiByZXR1cm4gJ0RlYnVnJztcbiAgICAgIGRlZmF1bHQ6IHJldHVybiB0eXBlLnRvVXBwZXJDYXNlKCk7XG4gICAgfVxuICB9XG5cbiAgI2FwcGx5TG9nTGFiZWxBdHRycyhsYWJlbDogSFRNTEVsZW1lbnQsIHR5cGU6IHN0cmluZykge1xuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnaW5mbyc6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICdpbmZvJyk7XG4gICAgICBjYXNlICd3YXJuaW5nJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnc3RhdHVzJywgJ3dhcm5pbmcnKTtcbiAgICAgIGNhc2UgJ2Vycm9yJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnc3RhdHVzJywgJ2RhbmdlcicpO1xuICAgICAgY2FzZSAnZGVidWcnOlxuICAgICAgICByZXR1cm4gbGFiZWwuc2V0QXR0cmlidXRlKCdjb2xvcicsICdwdXJwbGUnKTtcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIGxhYmVsLnNldEF0dHJpYnV0ZSgnY29sb3InLCAnZ3JleScpO1xuICAgIH1cbiAgfVxuXG4gICNzY3JvbGxMYXRlc3RJbnRvVmlldygpIHtcbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgcmVxdWVzdEFuaW1hdGlvbkZyYW1lKCgpID0+IHtcbiAgICAgIGNvbnN0IGxhc3RMb2cgPSB0aGlzLiNsb2dDb250YWluZXIhLmxhc3RFbGVtZW50Q2hpbGQ7XG4gICAgICBpZiAobGFzdExvZykge1xuICAgICAgICBsYXN0TG9nLnNjcm9sbEludG9WaWV3KHsgYmVoYXZpb3I6ICdhdXRvJywgYmxvY2s6ICdlbmQnIH0pO1xuICAgICAgfVxuICAgIH0pO1xuICB9XG5cbiAgI3Njcm9sbExvZ3NUb0JvdHRvbSgpIHtcbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgaWYgKHRoaXMuI2lzSW5pdGlhbExvYWQpIHtcbiAgICAgIHRoaXMuI3Njcm9sbExhdGVzdEludG9WaWV3KCk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgICAgfSwgMzUwKTtcbiAgICB9XG4gIH1cblxuICAjbWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2VJZk5lZWRlZCgpIHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgaGFzTG9jYWxTdG9yYWdlID1cbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1jb2xvci1zY2hlbWUnKSAhPT0gbnVsbCB8fFxuICAgICAgICBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWRyYXdlci1vcGVuJykgIT09IG51bGwgfHxcbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1kcmF3ZXItaGVpZ2h0JykgIT09IG51bGwgfHxcbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1hY3RpdmUtdGFiJykgIT09IG51bGw7XG5cbiAgICAgIGlmIChoYXNMb2NhbFN0b3JhZ2UpIHtcbiAgICAgICAgY29uc3QgbWlncmF0ZWQgPSBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLW1pZ3JhdGVkLXRvLWNvb2tpZXMnKTtcbiAgICAgICAgaWYgKCFtaWdyYXRlZCkge1xuICAgICAgICAgIFN0YXRlUGVyc2lzdGVuY2UubWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2UoKTtcbiAgICAgICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLW1pZ3JhdGVkLXRvLWNvb2tpZXMnLCAndHJ1ZScpO1xuICAgICAgICAgIHNldFRpbWVvdXQoKCkgPT4gd2luZG93LmxvY2F0aW9uLnJlbG9hZCgpLCAxMDApO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgLy8gbG9jYWxTdG9yYWdlIG5vdCBhdmFpbGFibGUsIHNraXAgbWlncmF0aW9uXG4gICAgfVxuICB9XG5cbiAgI3NldHVwQ29sb3JTY2hlbWVUb2dnbGUoKSB7XG4gICAgY29uc3QgdG9nZ2xlR3JvdXAgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJy5jb2xvci1zY2hlbWUtdG9nZ2xlJyk7XG4gICAgaWYgKCF0b2dnbGVHcm91cCkgcmV0dXJuO1xuXG4gICAgY29uc3Qgc3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFN0YXRlKCk7XG5cbiAgICB0aGlzLiNhcHBseUNvbG9yU2NoZW1lKHN0YXRlLmNvbG9yU2NoZW1lKTtcblxuICAgIGNvbnN0IGl0ZW1zID0gdG9nZ2xlR3JvdXAucXVlcnlTZWxlY3RvckFsbCgnY2VtLXBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJyk7XG4gICAgaXRlbXMuZm9yRWFjaChpdGVtID0+IHtcbiAgICAgIGlmICgoaXRlbSBhcyBhbnkpLnZhbHVlID09PSBzdGF0ZS5jb2xvclNjaGVtZSkge1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnc2VsZWN0ZWQnLCAnJyk7XG4gICAgICB9XG4gICAgfSk7XG5cbiAgICB0b2dnbGVHcm91cC5hZGRFdmVudExpc3RlbmVyKCdjZW0tcGYtdjYtdG9nZ2xlLWdyb3VwLWNoYW5nZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgY29uc3Qgc2NoZW1lID0gKGUgYXMgYW55KS52YWx1ZTtcbiAgICAgIHRoaXMuI2FwcGx5Q29sb3JTY2hlbWUoc2NoZW1lKTtcbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoeyBjb2xvclNjaGVtZTogc2NoZW1lIH0pO1xuICAgIH0pO1xuICB9XG5cbiAgI2FwcGx5Q29sb3JTY2hlbWUoc2NoZW1lOiBzdHJpbmcpIHtcbiAgICBjb25zdCBib2R5ID0gZG9jdW1lbnQuYm9keTtcblxuICAgIHN3aXRjaCAoc2NoZW1lKSB7XG4gICAgICBjYXNlICdsaWdodCc6XG4gICAgICAgIGJvZHkuc3R5bGUuY29sb3JTY2hlbWUgPSAnbGlnaHQnO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ2RhcmsnOlxuICAgICAgICBib2R5LnN0eWxlLmNvbG9yU2NoZW1lID0gJ2RhcmsnO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ3N5c3RlbSc6XG4gICAgICBkZWZhdWx0OlxuICAgICAgICBib2R5LnN0eWxlLmNvbG9yU2NoZW1lID0gJ2xpZ2h0IGRhcmsnO1xuICAgICAgICBicmVhaztcbiAgICB9XG4gIH1cblxuICAjc2V0dXBLbm9iQ29vcmRpbmF0aW9uKCkge1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjphdHRyaWJ1dGUtY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6cHJvcGVydHktY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOmF0dHJpYnV0ZS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6cHJvcGVydHktY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOmNzcy1wcm9wZXJ0eS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgfVxuXG4gICNvbktub2JDaGFuZ2UgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gdGhpcy4jZ2V0S25vYlRhcmdldChldmVudCk7XG4gICAgaWYgKCF0YXJnZXQpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIENvdWxkIG5vdCBmaW5kIGtub2IgdGFyZ2V0IGluZm8gaW4gZXZlbnQgcGF0aCcpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IHsgdGFnTmFtZSwgaW5zdGFuY2VJbmRleCB9ID0gdGFyZ2V0O1xuXG4gICAgY29uc3QgZGVtbyA9IHRoaXMuZGVtbztcbiAgICBpZiAoIWRlbW8pIHJldHVybjtcblxuICAgIGNvbnN0IGtub2JUeXBlID0gdGhpcy4jZ2V0S25vYlR5cGVGcm9tRXZlbnQoZXZlbnQpO1xuXG4gICAgY29uc3Qgc3VjY2VzcyA9IChkZW1vIGFzIGFueSkuYXBwbHlLbm9iQ2hhbmdlKFxuICAgICAga25vYlR5cGUsXG4gICAgICAoZXZlbnQgYXMgYW55KS5uYW1lLFxuICAgICAgKGV2ZW50IGFzIGFueSkudmFsdWUsXG4gICAgICB0YWdOYW1lLFxuICAgICAgaW5zdGFuY2VJbmRleFxuICAgICk7XG5cbiAgICBpZiAoIXN1Y2Nlc3MpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBhcHBseSBrbm9iIGNoYW5nZTonLCB7XG4gICAgICAgIHR5cGU6IGtub2JUeXBlLFxuICAgICAgICBuYW1lOiAoZXZlbnQgYXMgYW55KS5uYW1lLFxuICAgICAgICB0YWdOYW1lLFxuICAgICAgICBpbnN0YW5jZUluZGV4XG4gICAgICB9KTtcbiAgICB9XG4gIH07XG5cbiAgI29uS25vYkNsZWFyID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMuI2dldEtub2JUYXJnZXQoZXZlbnQpO1xuICAgIGlmICghdGFyZ2V0KSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBDb3VsZCBub3QgZmluZCBrbm9iIHRhcmdldCBpbmZvIGluIGV2ZW50IHBhdGgnKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICBjb25zdCB7IHRhZ05hbWUsIGluc3RhbmNlSW5kZXggfSA9IHRhcmdldDtcblxuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCBrbm9iVHlwZSA9IHRoaXMuI2dldEtub2JUeXBlRnJvbUNsZWFyRXZlbnQoZXZlbnQpO1xuICAgIGNvbnN0IGNsZWFyVmFsdWUgPSBrbm9iVHlwZSA9PT0gJ3Byb3BlcnR5JyA/IHVuZGVmaW5lZCA6ICcnO1xuXG4gICAgY29uc3Qgc3VjY2VzcyA9IChkZW1vIGFzIGFueSkuYXBwbHlLbm9iQ2hhbmdlKFxuICAgICAga25vYlR5cGUsXG4gICAgICAoZXZlbnQgYXMgYW55KS5uYW1lLFxuICAgICAgY2xlYXJWYWx1ZSxcbiAgICAgIHRhZ05hbWUsXG4gICAgICBpbnN0YW5jZUluZGV4XG4gICAgKTtcblxuICAgIGlmICghc3VjY2Vzcykge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGNsZWFyIGtub2I6Jywge1xuICAgICAgICB0eXBlOiBrbm9iVHlwZSxcbiAgICAgICAgbmFtZTogKGV2ZW50IGFzIGFueSkubmFtZSxcbiAgICAgICAgdGFnTmFtZSxcbiAgICAgICAgaW5zdGFuY2VJbmRleFxuICAgICAgfSk7XG4gICAgfVxuICB9O1xuXG4gICNnZXRLbm9iVGFyZ2V0KGV2ZW50OiBFdmVudCk6IHsgdGFnTmFtZTogc3RyaW5nOyBpbnN0YW5jZUluZGV4OiBudW1iZXIgfSB8IG51bGwge1xuICAgIGNvbnN0IGRlZmF1bHRUYWdOYW1lID0gdGhpcy5wcmltYXJ5VGFnTmFtZSB8fCAnJztcblxuICAgIGlmIChldmVudC5jb21wb3NlZFBhdGgpIHtcbiAgICAgIGZvciAoY29uc3QgZWxlbWVudCBvZiBldmVudC5jb21wb3NlZFBhdGgoKSkge1xuICAgICAgICBpZiAoIShlbGVtZW50IGluc3RhbmNlb2YgRWxlbWVudCkpIGNvbnRpbnVlO1xuXG4gICAgICAgIGlmICgoZWxlbWVudCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldD8uaXNFbGVtZW50S25vYiA9PT0gJ3RydWUnKSB7XG4gICAgICAgICAgY29uc3QgdGFnTmFtZSA9IChlbGVtZW50IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LnRhZ05hbWUgfHwgZGVmYXVsdFRhZ05hbWU7XG4gICAgICAgICAgbGV0IGluc3RhbmNlSW5kZXggPSBOdW1iZXIucGFyc2VJbnQoKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuaW5zdGFuY2VJbmRleCA/PyAnJywgMTApO1xuICAgICAgICAgIGlmIChOdW1iZXIuaXNOYU4oaW5zdGFuY2VJbmRleCkpIGluc3RhbmNlSW5kZXggPSAwO1xuICAgICAgICAgIHJldHVybiB7IHRhZ05hbWUsIGluc3RhbmNlSW5kZXggfTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiB7IHRhZ05hbWU6IGRlZmF1bHRUYWdOYW1lLCBpbnN0YW5jZUluZGV4OiAwIH07XG4gIH1cblxuICAjZ2V0S25vYlR5cGVGcm9tRXZlbnQoZXZlbnQ6IEV2ZW50KTogc3RyaW5nIHtcbiAgICBzd2l0Y2ggKGV2ZW50LnR5cGUpIHtcbiAgICAgIGNhc2UgJ2tub2I6YXR0cmlidXRlLWNoYW5nZSc6XG4gICAgICAgIHJldHVybiAnYXR0cmlidXRlJztcbiAgICAgIGNhc2UgJ2tub2I6cHJvcGVydHktY2hhbmdlJzpcbiAgICAgICAgcmV0dXJuICdwcm9wZXJ0eSc7XG4gICAgICBjYXNlICdrbm9iOmNzcy1wcm9wZXJ0eS1jaGFuZ2UnOlxuICAgICAgICByZXR1cm4gJ2Nzcy1wcm9wZXJ0eSc7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICByZXR1cm4gJ3Vua25vd24nO1xuICAgIH1cbiAgfVxuXG4gICNnZXRLbm9iVHlwZUZyb21DbGVhckV2ZW50KGV2ZW50OiBFdmVudCk6IHN0cmluZyB7XG4gICAgc3dpdGNoIChldmVudC50eXBlKSB7XG4gICAgICBjYXNlICdrbm9iOmF0dHJpYnV0ZS1jbGVhcic6XG4gICAgICAgIHJldHVybiAnYXR0cmlidXRlJztcbiAgICAgIGNhc2UgJ2tub2I6cHJvcGVydHktY2xlYXInOlxuICAgICAgICByZXR1cm4gJ3Byb3BlcnR5JztcbiAgICAgIGNhc2UgJ2tub2I6Y3NzLXByb3BlcnR5LWNsZWFyJzpcbiAgICAgICAgcmV0dXJuICdjc3MtcHJvcGVydHknO1xuICAgICAgZGVmYXVsdDpcbiAgICAgICAgcmV0dXJuICd1bmtub3duJztcbiAgICB9XG4gIH1cblxuICAjc2V0dXBUcmVlU3RhdGVQZXJzaXN0ZW5jZSgpIHtcbiAgICB0aGlzLiNoYW5kbGVUcmVlRXhwYW5kID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgICBpZiAoKGUudGFyZ2V0IGFzIEVsZW1lbnQpPy50YWdOYW1lICE9PSAnQ0VNLVBGLVY2LVRSRUUtSVRFTScpIHJldHVybjtcblxuICAgICAgY29uc3Qgbm9kZUlkID0gdGhpcy4jZ2V0VHJlZU5vZGVJZChlLnRhcmdldCBhcyBFbGVtZW50KTtcbiAgICAgIGNvbnN0IHRyZWVTdGF0ZSA9IFN0YXRlUGVyc2lzdGVuY2UuZ2V0VHJlZVN0YXRlKCk7XG4gICAgICBpZiAoIXRyZWVTdGF0ZS5leHBhbmRlZC5pbmNsdWRlcyhub2RlSWQpKSB7XG4gICAgICAgIHRyZWVTdGF0ZS5leHBhbmRlZC5wdXNoKG5vZGVJZCk7XG4gICAgICAgIFN0YXRlUGVyc2lzdGVuY2Uuc2V0VHJlZVN0YXRlKHRyZWVTdGF0ZSk7XG4gICAgICB9XG4gICAgfTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2V4cGFuZCcsIHRoaXMuI2hhbmRsZVRyZWVFeHBhbmQpO1xuXG4gICAgdGhpcy4jaGFuZGxlVHJlZUNvbGxhcHNlID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgICBpZiAoKGUudGFyZ2V0IGFzIEVsZW1lbnQpPy50YWdOYW1lICE9PSAnQ0VNLVBGLVY2LVRSRUUtSVRFTScpIHJldHVybjtcblxuICAgICAgY29uc3Qgbm9kZUlkID0gdGhpcy4jZ2V0VHJlZU5vZGVJZChlLnRhcmdldCBhcyBFbGVtZW50KTtcbiAgICAgIGNvbnN0IHRyZWVTdGF0ZSA9IFN0YXRlUGVyc2lzdGVuY2UuZ2V0VHJlZVN0YXRlKCk7XG4gICAgICBjb25zdCBpbmRleCA9IHRyZWVTdGF0ZS5leHBhbmRlZC5pbmRleE9mKG5vZGVJZCk7XG4gICAgICBpZiAoaW5kZXggPiAtMSkge1xuICAgICAgICB0cmVlU3RhdGUuZXhwYW5kZWQuc3BsaWNlKGluZGV4LCAxKTtcbiAgICAgICAgU3RhdGVQZXJzaXN0ZW5jZS5zZXRUcmVlU3RhdGUodHJlZVN0YXRlKTtcbiAgICAgIH1cbiAgICB9O1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignY29sbGFwc2UnLCB0aGlzLiNoYW5kbGVUcmVlQ29sbGFwc2UpO1xuXG4gICAgdGhpcy4jaGFuZGxlVHJlZVNlbGVjdCA9IChlOiBFdmVudCkgPT4ge1xuICAgICAgaWYgKChlLnRhcmdldCBhcyBFbGVtZW50KT8udGFnTmFtZSAhPT0gJ0NFTS1QRi1WNi1UUkVFLUlURU0nKSByZXR1cm47XG5cbiAgICAgIGNvbnN0IG5vZGVJZCA9IHRoaXMuI2dldFRyZWVOb2RlSWQoZS50YXJnZXQgYXMgRWxlbWVudCk7XG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVRyZWVTdGF0ZSh7IHNlbGVjdGVkOiBub2RlSWQgfSk7XG4gICAgfTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QpO1xuXG4gICAgdGhpcy4jYXBwbHlUcmVlU3RhdGUoKTtcbiAgfVxuXG4gICNhcHBseVRyZWVTdGF0ZSgpIHtcbiAgICBjb25zdCB0cmVlU3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFRyZWVTdGF0ZSgpO1xuXG4gICAgZm9yIChjb25zdCBub2RlSWQgb2YgdHJlZVN0YXRlLmV4cGFuZGVkKSB7XG4gICAgICBjb25zdCB0cmVlSXRlbSA9IHRoaXMuI2ZpbmRUcmVlSXRlbUJ5SWQobm9kZUlkKTtcbiAgICAgIGlmICh0cmVlSXRlbSAmJiAhdHJlZUl0ZW0uaGFzQXR0cmlidXRlKCdleHBhbmRlZCcpKSB7XG4gICAgICAgIHRyZWVJdGVtLnNldEF0dHJpYnV0ZSgnZXhwYW5kZWQnLCAnJyk7XG4gICAgICB9XG4gICAgfVxuXG4gICAgaWYgKHRyZWVTdGF0ZS5zZWxlY3RlZCkge1xuICAgICAgY29uc3QgdHJlZUl0ZW0gPSB0aGlzLiNmaW5kVHJlZUl0ZW1CeUlkKHRyZWVTdGF0ZS5zZWxlY3RlZCk7XG4gICAgICBpZiAodHJlZUl0ZW0gJiYgIXRyZWVJdGVtLmhhc0F0dHJpYnV0ZSgnY3VycmVudCcpKSB7XG4gICAgICAgIHRyZWVJdGVtLnNldEF0dHJpYnV0ZSgnY3VycmVudCcsICcnKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjc2V0dXBTaWRlYmFyU3RhdGVQZXJzaXN0ZW5jZSgpIHtcbiAgICBjb25zdCBwYWdlID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdjZW0tcGYtdjYtcGFnZScpO1xuXG4gICAgaWYgKCFwYWdlKSByZXR1cm47XG5cbiAgICBwYWdlLmFkZEV2ZW50TGlzdGVuZXIoJ3NpZGViYXItdG9nZ2xlJywgKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgICAgY29uc3QgY29sbGFwc2VkID0gIShldmVudCBhcyBhbnkpLmV4cGFuZGVkO1xuXG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHtcbiAgICAgICAgc2lkZWJhcjogeyBjb2xsYXBzZWQgfVxuICAgICAgfSk7XG4gICAgfSk7XG4gIH1cblxuICAjZmluZFRyZWVJdGVtQnlJZChub2RlSWQ6IHN0cmluZyk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICBjb25zdCBwYXJ0cyA9IG5vZGVJZC5zcGxpdCgnOicpO1xuICAgIGNvbnN0IFt0eXBlLCBtb2R1bGVQYXRoLCB0YWdOYW1lLCBuYW1lXSA9IHBhcnRzO1xuXG4gICAgbGV0IGF0dHJTdWZmaXggPSAnJztcbiAgICBpZiAodGFnTmFtZSkge1xuICAgICAgYXR0clN1ZmZpeCArPSBgW2RhdGEtdGFnLW5hbWU9XCIke0NTUy5lc2NhcGUodGFnTmFtZSl9XCJdYDtcbiAgICB9XG4gICAgaWYgKG5hbWUpIHtcbiAgICAgIGF0dHJTdWZmaXggKz0gYFtkYXRhLW5hbWU9XCIke0NTUy5lc2NhcGUobmFtZSl9XCJdYDtcbiAgICB9XG5cbiAgICBsZXQgc2VsZWN0b3IgPSBgY2VtLXBmLXY2LXRyZWUtaXRlbVtkYXRhLXR5cGU9XCIke0NTUy5lc2NhcGUodHlwZSl9XCJdYDtcbiAgICBpZiAobW9kdWxlUGF0aCkge1xuICAgICAgY29uc3QgZXNjYXBlZE1vZHVsZVBhdGggPSBDU1MuZXNjYXBlKG1vZHVsZVBhdGgpO1xuICAgICAgY29uc3QgZXNjYXBlZFR5cGUgPSBDU1MuZXNjYXBlKHR5cGUpO1xuICAgICAgY29uc3Qgc2VsZWN0b3IxID0gYGNlbS1wZi12Ni10cmVlLWl0ZW1bZGF0YS10eXBlPVwiJHtlc2NhcGVkVHlwZX1cIl1bZGF0YS1tb2R1bGUtcGF0aD1cIiR7ZXNjYXBlZE1vZHVsZVBhdGh9XCJdJHthdHRyU3VmZml4fWA7XG4gICAgICBjb25zdCBzZWxlY3RvcjIgPSBgY2VtLXBmLXY2LXRyZWUtaXRlbVtkYXRhLXR5cGU9XCIke2VzY2FwZWRUeXBlfVwiXVtkYXRhLXBhdGg9XCIke2VzY2FwZWRNb2R1bGVQYXRofVwiXSR7YXR0clN1ZmZpeH1gO1xuICAgICAgc2VsZWN0b3IgPSBgJHtzZWxlY3RvcjF9LCAke3NlbGVjdG9yMn1gO1xuICAgIH0gZWxzZSB7XG4gICAgICBzZWxlY3RvciArPSBhdHRyU3VmZml4O1xuICAgIH1cblxuICAgIHJldHVybiB0aGlzLnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpO1xuICB9XG5cbiAgI2dldFRyZWVOb2RlSWQodHJlZUl0ZW06IEVsZW1lbnQpOiBzdHJpbmcge1xuICAgIGNvbnN0IHR5cGUgPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtdHlwZScpO1xuICAgIGNvbnN0IG1vZHVsZVBhdGggPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtbW9kdWxlLXBhdGgnKSB8fCB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtcGF0aCcpO1xuICAgIGNvbnN0IHRhZ05hbWUgPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtdGFnLW5hbWUnKTtcbiAgICBjb25zdCBuYW1lID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLW5hbWUnKTtcbiAgICBjb25zdCBjYXRlZ29yeSA9IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS1jYXRlZ29yeScpO1xuXG4gICAgY29uc3QgcGFydHMgPSBbdHlwZV07XG4gICAgaWYgKG1vZHVsZVBhdGgpIHBhcnRzLnB1c2gobW9kdWxlUGF0aCk7XG4gICAgaWYgKHRhZ05hbWUpIHBhcnRzLnB1c2godGFnTmFtZSk7XG4gICAgaWYgKGNhdGVnb3J5KSB7XG4gICAgICBwYXJ0cy5wdXNoKGNhdGVnb3J5KTtcbiAgICB9IGVsc2UgaWYgKG5hbWUpIHtcbiAgICAgIHBhcnRzLnB1c2gobmFtZSk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHBhcnRzLmpvaW4oJzonKTtcbiAgfVxuXG4gIC8vIEV2ZW50IERpc2NvdmVyeSAmIENhcHR1cmUgTWV0aG9kc1xuXG4gIGFzeW5jICNkaXNjb3ZlckVsZW1lbnRFdmVudHMoKTogUHJvbWlzZTxNYXA8c3RyaW5nLCBFdmVudEluZm8+PiB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHJlc3BvbnNlID0gYXdhaXQgZmV0Y2goJy9jdXN0b20tZWxlbWVudHMuanNvbicpO1xuICAgICAgaWYgKCFyZXNwb25zZS5vaykge1xuICAgICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBObyBtYW5pZmVzdCBhdmFpbGFibGUgZm9yIGV2ZW50IGRpc2NvdmVyeScpO1xuICAgICAgICByZXR1cm4gbmV3IE1hcCgpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBtYW5pZmVzdCA9IGF3YWl0IHJlc3BvbnNlLmpzb24oKSBhcyBNYW5pZmVzdDtcbiAgICAgIHRoaXMuI21hbmlmZXN0ID0gbWFuaWZlc3Q7XG5cbiAgICAgIGNvbnN0IGV2ZW50TWFwID0gbmV3IE1hcDxzdHJpbmcsIEV2ZW50SW5mbz4oKTtcblxuICAgICAgZm9yIChjb25zdCBtb2R1bGUgb2YgbWFuaWZlc3QubW9kdWxlcyB8fCBbXSkge1xuICAgICAgICBmb3IgKGNvbnN0IGRlY2xhcmF0aW9uIG9mIG1vZHVsZS5kZWNsYXJhdGlvbnMgfHwgW10pIHtcbiAgICAgICAgICBpZiAoZGVjbGFyYXRpb24uY3VzdG9tRWxlbWVudCAmJiBkZWNsYXJhdGlvbi50YWdOYW1lKSB7XG4gICAgICAgICAgICBjb25zdCB0YWdOYW1lID0gZGVjbGFyYXRpb24udGFnTmFtZTtcbiAgICAgICAgICAgIGNvbnN0IGV2ZW50cyA9IGRlY2xhcmF0aW9uLmV2ZW50cyB8fCBbXTtcblxuICAgICAgICAgICAgaWYgKGV2ZW50cy5sZW5ndGggPiAwKSB7XG4gICAgICAgICAgICAgIGNvbnN0IGV2ZW50TmFtZXMgPSBuZXcgU2V0KGV2ZW50cy5tYXAoZSA9PiBlLm5hbWUpKTtcbiAgICAgICAgICAgICAgZXZlbnRNYXAuc2V0KHRhZ05hbWUsIHtcbiAgICAgICAgICAgICAgICBldmVudE5hbWVzLFxuICAgICAgICAgICAgICAgIGV2ZW50czogZXZlbnRzXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICByZXR1cm4gZXZlbnRNYXA7XG4gICAgfSBjYXRjaCAoZXJyb3IpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIEVycm9yIGxvYWRpbmcgbWFuaWZlc3QgZm9yIGV2ZW50IGRpc2NvdmVyeTonLCBlcnJvcik7XG4gICAgICByZXR1cm4gbmV3IE1hcCgpO1xuICAgIH1cbiAgfVxuXG4gIGFzeW5jICNzZXR1cEV2ZW50Q2FwdHVyZSgpIHtcbiAgICB0aGlzLiNlbGVtZW50RXZlbnRNYXAgPSBhd2FpdCB0aGlzLiNkaXNjb3ZlckVsZW1lbnRFdmVudHMoKTtcblxuICAgIGlmICh0aGlzLiNlbGVtZW50RXZlbnRNYXAuc2l6ZSA9PT0gMCkgcmV0dXJuO1xuXG4gICAgdGhpcy4jYXR0YWNoRXZlbnRMaXN0ZW5lcnMoKTtcbiAgICB0aGlzLiN1cGRhdGVFdmVudEZpbHRlcnMoKTtcbiAgICB0aGlzLiNvYnNlcnZlRGVtb011dGF0aW9ucygpO1xuICB9XG5cbiAgI2F0dGFjaEV2ZW50TGlzdGVuZXJzKCkge1xuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCByb290ID0gZGVtby5zaGFkb3dSb290ID8/IGRlbW87XG5cbiAgICBmb3IgKGNvbnN0IFt0YWdOYW1lLCBldmVudEluZm9dIG9mIHRoaXMuI2VsZW1lbnRFdmVudE1hcCEpIHtcbiAgICAgIGNvbnN0IGVsZW1lbnRzID0gcm9vdC5xdWVyeVNlbGVjdG9yQWxsKHRhZ05hbWUpO1xuXG4gICAgICBmb3IgKGNvbnN0IGVsZW1lbnQgb2YgZWxlbWVudHMpIHtcbiAgICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgZXZlbnRJbmZvLmV2ZW50TmFtZXMpIHtcbiAgICAgICAgICBlbGVtZW50LmFkZEV2ZW50TGlzdGVuZXIoZXZlbnROYW1lLCB0aGlzLiNoYW5kbGVFbGVtZW50RXZlbnQsIHsgY2FwdHVyZTogdHJ1ZSB9KTtcbiAgICAgICAgfVxuICAgICAgICAoZWxlbWVudCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5jZW1FdmVudHNBdHRhY2hlZCA9ICd0cnVlJztcbiAgICAgICAgdGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjb2JzZXJ2ZURlbW9NdXRhdGlvbnMoKSB7XG4gICAgY29uc3QgZGVtbyA9IHRoaXMuZGVtbztcbiAgICBpZiAoIWRlbW8pIHJldHVybjtcblxuICAgIGNvbnN0IHJvb3QgPSBkZW1vLnNoYWRvd1Jvb3QgPz8gZGVtbztcblxuICAgIHRoaXMuI29ic2VydmVyLm9ic2VydmUocm9vdCwge1xuICAgICAgY2hpbGRMaXN0OiB0cnVlLFxuICAgICAgc3VidHJlZTogdHJ1ZVxuICAgIH0pO1xuICB9XG5cbiAgI2hhbmRsZUVsZW1lbnRFdmVudCA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCBlbGVtZW50ID0gZXZlbnQuY3VycmVudFRhcmdldDtcbiAgICBpZiAoIShlbGVtZW50IGluc3RhbmNlb2YgSFRNTEVsZW1lbnQpKSByZXR1cm47XG5cbiAgICBjb25zdCB0YWdOYW1lID0gZWxlbWVudC50YWdOYW1lLnRvTG93ZXJDYXNlKCk7XG4gICAgY29uc3QgZXZlbnRJbmZvID0gdGhpcy4jZWxlbWVudEV2ZW50TWFwPy5nZXQodGFnTmFtZSk7XG5cbiAgICBpZiAoIWV2ZW50SW5mbyB8fCAhZXZlbnRJbmZvLmV2ZW50TmFtZXMuaGFzKGV2ZW50LnR5cGUpKSByZXR1cm47XG5cbiAgICB0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuYWRkKHRhZ05hbWUpO1xuICAgIHRoaXMuI2NhcHR1cmVFdmVudChldmVudCwgZWxlbWVudCwgdGFnTmFtZSwgZXZlbnRJbmZvKTtcbiAgfTtcblxuICAjZ2V0RXZlbnREb2N1bWVudGF0aW9uKG1hbmlmZXN0RXZlbnQ6IEV2ZW50SW5mb1snZXZlbnRzJ11bMF0gfCB1bmRlZmluZWQpIHtcbiAgICBpZiAoIW1hbmlmZXN0RXZlbnQpIHtcbiAgICAgIHJldHVybiB7IHN1bW1hcnk6IG51bGwsIGRlc2NyaXB0aW9uOiBudWxsIH07XG4gICAgfVxuXG4gICAgbGV0IHN1bW1hcnkgPSBtYW5pZmVzdEV2ZW50LnN1bW1hcnkgfHwgbnVsbDtcbiAgICBsZXQgZGVzY3JpcHRpb24gPSBtYW5pZmVzdEV2ZW50LmRlc2NyaXB0aW9uIHx8IG51bGw7XG5cbiAgICBpZiAobWFuaWZlc3RFdmVudC50eXBlPy50ZXh0ICYmIHRoaXMuI21hbmlmZXN0KSB7XG4gICAgICBjb25zdCB0eXBlTmFtZSA9IG1hbmlmZXN0RXZlbnQudHlwZS50ZXh0O1xuICAgICAgY29uc3QgdHlwZURlY2xhcmF0aW9uID0gdGhpcy4jZmluZFR5cGVEZWNsYXJhdGlvbih0eXBlTmFtZSk7XG5cbiAgICAgIGlmICh0eXBlRGVjbGFyYXRpb24pIHtcbiAgICAgICAgaWYgKCFzdW1tYXJ5ICYmIHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5KSB7XG4gICAgICAgICAgc3VtbWFyeSA9IHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5O1xuICAgICAgICB9IGVsc2UgaWYgKHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5ICYmIHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5ICE9PSBzdW1tYXJ5KSB7XG4gICAgICAgICAgc3VtbWFyeSA9IHN1bW1hcnkgPyBgJHtzdW1tYXJ5fVxcblxcbkZyb20gJHt0eXBlTmFtZX06ICR7dHlwZURlY2xhcmF0aW9uLnN1bW1hcnl9YCA6IHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5O1xuICAgICAgICB9XG5cbiAgICAgICAgaWYgKCFkZXNjcmlwdGlvbiAmJiB0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb24pIHtcbiAgICAgICAgICBkZXNjcmlwdGlvbiA9IHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbjtcbiAgICAgICAgfSBlbHNlIGlmICh0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb24gJiYgdHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uICE9PSBkZXNjcmlwdGlvbikge1xuICAgICAgICAgIGRlc2NyaXB0aW9uID0gZGVzY3JpcHRpb24gPyBgJHtkZXNjcmlwdGlvbn1cXG5cXG4ke3R5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbn1gIDogdHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIHsgc3VtbWFyeSwgZGVzY3JpcHRpb24gfTtcbiAgfVxuXG4gICNmaW5kVHlwZURlY2xhcmF0aW9uKHR5cGVOYW1lOiBzdHJpbmcpIHtcbiAgICBpZiAoIXRoaXMuI21hbmlmZXN0KSByZXR1cm4gbnVsbDtcblxuICAgIGZvciAoY29uc3QgbW9kdWxlIG9mIHRoaXMuI21hbmlmZXN0Lm1vZHVsZXMgfHwgW10pIHtcbiAgICAgIGZvciAoY29uc3QgZGVjbGFyYXRpb24gb2YgbW9kdWxlLmRlY2xhcmF0aW9ucyB8fCBbXSkge1xuICAgICAgICBpZiAoZGVjbGFyYXRpb24ubmFtZSA9PT0gdHlwZU5hbWUgJiZcbiAgICAgICAgICAgIChkZWNsYXJhdGlvbi5raW5kID09PSAnY2xhc3MnIHx8IGRlY2xhcmF0aW9uLmtpbmQgPT09ICdpbnRlcmZhY2UnKSkge1xuICAgICAgICAgIHJldHVybiBkZWNsYXJhdGlvbiBhcyB7IHN1bW1hcnk/OiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH07XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICByZXR1cm4gbnVsbDtcbiAgfVxuXG4gICNjYXB0dXJlRXZlbnQoZXZlbnQ6IEV2ZW50LCB0YXJnZXQ6IEhUTUxFbGVtZW50LCB0YWdOYW1lOiBzdHJpbmcsIGV2ZW50SW5mbzogRXZlbnRJbmZvKSB7XG4gICAgY29uc3QgbWFuaWZlc3RFdmVudCA9IGV2ZW50SW5mby5ldmVudHMuZmluZChlID0+IGUubmFtZSA9PT0gZXZlbnQudHlwZSk7XG5cbiAgICBjb25zdCBldmVudERvY3MgPSB0aGlzLiNnZXRFdmVudERvY3VtZW50YXRpb24obWFuaWZlc3RFdmVudCk7XG5cbiAgICBjb25zdCBjdXN0b21Qcm9wZXJ0aWVzID0gdGhpcy4jZXh0cmFjdEV2ZW50UHJvcGVydGllcyhldmVudCk7XG5cbiAgICBjb25zdCBldmVudFJlY29yZDogRXZlbnRSZWNvcmQgPSB7XG4gICAgICBpZDogYCR7RGF0ZS5ub3coKX0tJHtNYXRoLnJhbmRvbSgpfWAsXG4gICAgICB0aW1lc3RhbXA6IG5ldyBEYXRlKCksXG4gICAgICBldmVudE5hbWU6IGV2ZW50LnR5cGUsXG4gICAgICB0YWdOYW1lOiB0YWdOYW1lLFxuICAgICAgZWxlbWVudElkOiB0YXJnZXQuaWQgfHwgbnVsbCxcbiAgICAgIGVsZW1lbnRDbGFzczogdGFyZ2V0LmNsYXNzTmFtZSB8fCBudWxsLFxuICAgICAgY3VzdG9tUHJvcGVydGllczogY3VzdG9tUHJvcGVydGllcyxcbiAgICAgIG1hbmlmZXN0VHlwZTogbWFuaWZlc3RFdmVudD8udHlwZT8udGV4dCB8fCBudWxsLFxuICAgICAgc3VtbWFyeTogZXZlbnREb2NzLnN1bW1hcnksXG4gICAgICBkZXNjcmlwdGlvbjogZXZlbnREb2NzLmRlc2NyaXB0aW9uLFxuICAgICAgYnViYmxlczogZXZlbnQuYnViYmxlcyxcbiAgICAgIGNvbXBvc2VkOiBldmVudC5jb21wb3NlZCxcbiAgICAgIGNhbmNlbGFibGU6IGV2ZW50LmNhbmNlbGFibGUsXG4gICAgICBkZWZhdWx0UHJldmVudGVkOiBldmVudC5kZWZhdWx0UHJldmVudGVkXG4gICAgfTtcblxuICAgIHRoaXMuI2NhcHR1cmVkRXZlbnRzLnB1c2goZXZlbnRSZWNvcmQpO1xuXG4gICAgaWYgKHRoaXMuI2NhcHR1cmVkRXZlbnRzLmxlbmd0aCA+IHRoaXMuI21heENhcHR1cmVkRXZlbnRzKSB7XG4gICAgICB0aGlzLiNjYXB0dXJlZEV2ZW50cy5zaGlmdCgpO1xuICAgIH1cblxuICAgIHRoaXMuI3JlbmRlckV2ZW50KGV2ZW50UmVjb3JkKTtcbiAgfVxuXG4gICNleHRyYWN0RXZlbnRQcm9wZXJ0aWVzKGV2ZW50OiBFdmVudCk6IFJlY29yZDxzdHJpbmcsIHVua25vd24+IHtcbiAgICBjb25zdCBwcm9wZXJ0aWVzOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiA9IHt9O1xuICAgIGNvbnN0IGV2ZW50UHJvdG90eXBlS2V5cyA9IG5ldyBTZXQoT2JqZWN0LmdldE93blByb3BlcnR5TmFtZXMoRXZlbnQucHJvdG90eXBlKSk7XG5cbiAgICBjb25zdCBzZXJpYWxpemVWYWx1ZSA9ICh2YWx1ZTogdW5rbm93bik6IHVua25vd24gPT4ge1xuICAgICAgdHJ5IHtcbiAgICAgICAgcmV0dXJuIEpTT04ucGFyc2UoSlNPTi5zdHJpbmdpZnkodmFsdWUpKTtcbiAgICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgICAgdHJ5IHtcbiAgICAgICAgICByZXR1cm4gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgfSBjYXRjaCAoc3RyaW5nRXJyKSB7XG4gICAgICAgICAgcmV0dXJuICdbTm90IHNlcmlhbGl6YWJsZV0nO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfTtcblxuICAgIGlmIChldmVudCBpbnN0YW5jZW9mIEN1c3RvbUV2ZW50ICYmIGV2ZW50LmRldGFpbCAhPT0gdW5kZWZpbmVkKSB7XG4gICAgICBwcm9wZXJ0aWVzLmRldGFpbCA9IHNlcmlhbGl6ZVZhbHVlKGV2ZW50LmRldGFpbCk7XG4gICAgfVxuXG4gICAgZm9yIChjb25zdCBrZXkgb2YgT2JqZWN0LmdldE93blByb3BlcnR5TmFtZXMoZXZlbnQpKSB7XG4gICAgICBpZiAoIWV2ZW50UHJvdG90eXBlS2V5cy5oYXMoa2V5KSAmJiAha2V5LnN0YXJ0c1dpdGgoJ18nKSAmJiAhcHJvcGVydGllcy5oYXNPd25Qcm9wZXJ0eShrZXkpKSB7XG4gICAgICAgIHByb3BlcnRpZXNba2V5XSA9IHNlcmlhbGl6ZVZhbHVlKChldmVudCBhcyBhbnkpW2tleV0pO1xuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiBwcm9wZXJ0aWVzO1xuICB9XG5cbiAgI3JlbmRlckV2ZW50KGV2ZW50UmVjb3JkOiBFdmVudFJlY29yZCkge1xuICAgIGlmICghdGhpcy4jZXZlbnRMaXN0KSByZXR1cm47XG5cbiAgICBjb25zdCBmcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNldmVudEVudHJ5VGVtcGxhdGUuY29udGVudC5jbG9uZU5vZGUodHJ1ZSkgYXMgRG9jdW1lbnRGcmFnbWVudDtcblxuICAgIGNvbnN0IHRpbWUgPSBldmVudFJlY29yZC50aW1lc3RhbXAudG9Mb2NhbGVUaW1lU3RyaW5nKCk7XG5cbiAgICBjb25zdCBjb250YWluZXIgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNvbnRhaW5lclwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgIGNvbnRhaW5lci5kYXRhc2V0LmV2ZW50SWQgPSBldmVudFJlY29yZC5pZDtcbiAgICBjb250YWluZXIuZGF0YXNldC5ldmVudFR5cGUgPSBldmVudFJlY29yZC5ldmVudE5hbWU7XG4gICAgY29udGFpbmVyLmRhdGFzZXQuZWxlbWVudFR5cGUgPSBldmVudFJlY29yZC50YWdOYW1lO1xuXG4gICAgY29uc3QgdGV4dE1hdGNoID0gdGhpcy4jZXZlbnRNYXRjaGVzVGV4dEZpbHRlcihldmVudFJlY29yZCk7XG4gICAgY29uc3QgdHlwZU1hdGNoID0gdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5zaXplID09PSAwIHx8IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuaGFzKGV2ZW50UmVjb3JkLmV2ZW50TmFtZSk7XG4gICAgY29uc3QgZWxlbWVudE1hdGNoID0gdGhpcy4jZWxlbWVudEZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNlbGVtZW50RmlsdGVycy5oYXMoZXZlbnRSZWNvcmQudGFnTmFtZSk7XG5cbiAgICBpZiAoISh0ZXh0TWF0Y2ggJiYgdHlwZU1hdGNoICYmIGVsZW1lbnRNYXRjaCkpIHtcbiAgICAgIGNvbnRhaW5lci5zZXRBdHRyaWJ1dGUoJ2hpZGRlbicsICcnKTtcbiAgICB9XG5cbiAgICBjb25zdCBsYWJlbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibGFiZWxcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICBsYWJlbC50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLmV2ZW50TmFtZTtcbiAgICBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICdpbmZvJyk7XG5cbiAgICBjb25zdCB0aW1lRWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cInRpbWVcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICB0aW1lRWwuc2V0QXR0cmlidXRlKCdkYXRldGltZScsIGV2ZW50UmVjb3JkLnRpbWVzdGFtcC50b0lTT1N0cmluZygpKTtcbiAgICB0aW1lRWwudGV4dENvbnRlbnQgPSB0aW1lO1xuXG4gICAgY29uc3QgZWxlbWVudEVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJlbGVtZW50XCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgbGV0IGVsZW1lbnRUZXh0ID0gYDwke2V2ZW50UmVjb3JkLnRhZ05hbWV9PmA7XG4gICAgaWYgKGV2ZW50UmVjb3JkLmVsZW1lbnRJZCkge1xuICAgICAgZWxlbWVudFRleHQgKz0gYCMke2V2ZW50UmVjb3JkLmVsZW1lbnRJZH1gO1xuICAgIH1cbiAgICBlbGVtZW50RWwudGV4dENvbnRlbnQgPSBlbGVtZW50VGV4dDtcblxuICAgIHRoaXMuI2V2ZW50TGlzdC5hcHBlbmQoZnJhZ21lbnQpO1xuXG4gICAgaWYgKCF0aGlzLiNzZWxlY3RlZEV2ZW50SWQpIHtcbiAgICAgIHRoaXMuI3NlbGVjdEV2ZW50KGV2ZW50UmVjb3JkLmlkKTtcbiAgICB9XG5cbiAgICBpZiAodGhpcy4jZHJhd2VyT3BlbiAmJiB0aGlzLiNpc0V2ZW50c1RhYkFjdGl2ZSgpKSB7XG4gICAgICB0aGlzLiNzY3JvbGxFdmVudHNUb0JvdHRvbSgpO1xuICAgIH1cbiAgfVxuXG4gICNzZWxlY3RFdmVudChldmVudElkOiBzdHJpbmcpIHtcbiAgICBjb25zdCBldmVudFJlY29yZCA9IHRoaXMuI2dldEV2ZW50UmVjb3JkQnlJZChldmVudElkKTtcbiAgICBpZiAoIWV2ZW50UmVjb3JkKSByZXR1cm47XG5cbiAgICB0aGlzLiNzZWxlY3RlZEV2ZW50SWQgPSBldmVudElkO1xuXG4gICAgY29uc3QgYWxsSXRlbXMgPSB0aGlzLiNldmVudExpc3Q/LnF1ZXJ5U2VsZWN0b3JBbGwoJy5ldmVudC1saXN0LWl0ZW0nKTtcbiAgICBhbGxJdGVtcz8uZm9yRWFjaChpdGVtID0+IHtcbiAgICAgIGlmICgoaXRlbSBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5ldmVudElkID09PSBldmVudElkKSB7XG4gICAgICAgIGl0ZW0uY2xhc3NMaXN0LmFkZCgnc2VsZWN0ZWQnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2FyaWEtc2VsZWN0ZWQnLCAndHJ1ZScpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgaXRlbS5jbGFzc0xpc3QucmVtb3ZlKCdzZWxlY3RlZCcpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnYXJpYS1zZWxlY3RlZCcsICdmYWxzZScpO1xuICAgICAgfVxuICAgIH0pO1xuXG4gICAgaWYgKHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyKSB7XG4gICAgICB0aGlzLiNldmVudERldGFpbEhlYWRlci5pbm5lckhUTUwgPSAnJztcblxuICAgICAgY29uc3QgaGVhZGVyQ29udGVudCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuICAgICAgaGVhZGVyQ29udGVudC5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLWhlYWRlci1jb250ZW50JztcblxuICAgICAgY29uc3QgZXZlbnROYW1lID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnaDMnKTtcbiAgICAgIGV2ZW50TmFtZS50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLmV2ZW50TmFtZTtcbiAgICAgIGV2ZW50TmFtZS5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLW5hbWUnO1xuICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChldmVudE5hbWUpO1xuXG4gICAgICBpZiAoZXZlbnRSZWNvcmQuc3VtbWFyeSkge1xuICAgICAgICBjb25zdCBzdW1tYXJ5ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncCcpO1xuICAgICAgICBzdW1tYXJ5LnRleHRDb250ZW50ID0gZXZlbnRSZWNvcmQuc3VtbWFyeTtcbiAgICAgICAgc3VtbWFyeS5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXN1bW1hcnknO1xuICAgICAgICBoZWFkZXJDb250ZW50LmFwcGVuZENoaWxkKHN1bW1hcnkpO1xuICAgICAgfVxuXG4gICAgICBpZiAoZXZlbnRSZWNvcmQuZGVzY3JpcHRpb24pIHtcbiAgICAgICAgY29uc3QgZGVzY3JpcHRpb24gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwJyk7XG4gICAgICAgIGRlc2NyaXB0aW9uLnRleHRDb250ZW50ID0gZXZlbnRSZWNvcmQuZGVzY3JpcHRpb247XG4gICAgICAgIGRlc2NyaXB0aW9uLmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtZGVzY3JpcHRpb24nO1xuICAgICAgICBoZWFkZXJDb250ZW50LmFwcGVuZENoaWxkKGRlc2NyaXB0aW9uKTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgbWV0YSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuICAgICAgbWV0YS5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLW1ldGEnO1xuXG4gICAgICBjb25zdCB0aW1lRWwgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0aW1lJyk7XG4gICAgICB0aW1lRWwuc2V0QXR0cmlidXRlKCdkYXRldGltZScsIGV2ZW50UmVjb3JkLnRpbWVzdGFtcC50b0lTT1N0cmluZygpKTtcbiAgICAgIHRpbWVFbC50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLnRpbWVzdGFtcC50b0xvY2FsZVRpbWVTdHJpbmcoKTtcbiAgICAgIHRpbWVFbC5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXRpbWUnO1xuXG4gICAgICBjb25zdCBlbGVtZW50ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgbGV0IGVsZW1lbnRUZXh0ID0gYDwke2V2ZW50UmVjb3JkLnRhZ05hbWV9PmA7XG4gICAgICBpZiAoZXZlbnRSZWNvcmQuZWxlbWVudElkKSB7XG4gICAgICAgIGVsZW1lbnRUZXh0ICs9IGAjJHtldmVudFJlY29yZC5lbGVtZW50SWR9YDtcbiAgICAgIH1cbiAgICAgIGVsZW1lbnQudGV4dENvbnRlbnQgPSBlbGVtZW50VGV4dDtcbiAgICAgIGVsZW1lbnQuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1lbGVtZW50JztcblxuICAgICAgbWV0YS5hcHBlbmRDaGlsZCh0aW1lRWwpO1xuICAgICAgbWV0YS5hcHBlbmRDaGlsZChlbGVtZW50KTtcblxuICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChtZXRhKTtcblxuICAgICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIuYXBwZW5kQ2hpbGQoaGVhZGVyQ29udGVudCk7XG4gICAgfVxuXG4gICAgaWYgKHRoaXMuI2V2ZW50RGV0YWlsQm9keSkge1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5LmlubmVySFRNTCA9ICcnO1xuXG4gICAgICBjb25zdCBwcm9wZXJ0aWVzSGVhZGluZyA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2g0Jyk7XG4gICAgICBwcm9wZXJ0aWVzSGVhZGluZy50ZXh0Q29udGVudCA9ICdQcm9wZXJ0aWVzJztcbiAgICAgIHByb3BlcnRpZXNIZWFkaW5nLmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtcHJvcGVydGllcy1oZWFkaW5nJztcblxuICAgICAgY29uc3QgcHJvcGVydGllc0NvbnRhaW5lciA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuICAgICAgcHJvcGVydGllc0NvbnRhaW5lci5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMnO1xuXG4gICAgICBjb25zdCBldmVudFByb3BlcnRpZXMgPSB0aGlzLiNidWlsZFByb3BlcnRpZXNGb3JEaXNwbGF5KGV2ZW50UmVjb3JkKTtcbiAgICAgIGlmIChPYmplY3Qua2V5cyhldmVudFByb3BlcnRpZXMpLmxlbmd0aCA+IDApIHtcbiAgICAgICAgcHJvcGVydGllc0NvbnRhaW5lci5hcHBlbmRDaGlsZCh0aGlzLiNidWlsZFByb3BlcnR5VHJlZShldmVudFByb3BlcnRpZXMpKTtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHByb3BlcnRpZXNDb250YWluZXIudGV4dENvbnRlbnQgPSAnTm8gcHJvcGVydGllcyB0byBkaXNwbGF5JztcbiAgICAgIH1cblxuICAgICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5LmFwcGVuZENoaWxkKHByb3BlcnRpZXNIZWFkaW5nKTtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5hcHBlbmRDaGlsZChwcm9wZXJ0aWVzQ29udGFpbmVyKTtcbiAgICB9XG4gIH1cblxuICAjYnVpbGRQcm9wZXJ0aWVzRm9yRGlzcGxheShldmVudFJlY29yZDogRXZlbnRSZWNvcmQpOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiB7XG4gICAgY29uc3QgcHJvcGVydGllczogUmVjb3JkPHN0cmluZywgdW5rbm93bj4gPSB7fTtcblxuICAgIGlmIChldmVudFJlY29yZC5jdXN0b21Qcm9wZXJ0aWVzKSB7XG4gICAgICBPYmplY3QuYXNzaWduKHByb3BlcnRpZXMsIGV2ZW50UmVjb3JkLmN1c3RvbVByb3BlcnRpZXMpO1xuICAgIH1cblxuICAgIHByb3BlcnRpZXMuYnViYmxlcyA9IGV2ZW50UmVjb3JkLmJ1YmJsZXM7XG4gICAgcHJvcGVydGllcy5jYW5jZWxhYmxlID0gZXZlbnRSZWNvcmQuY2FuY2VsYWJsZTtcbiAgICBwcm9wZXJ0aWVzLmRlZmF1bHRQcmV2ZW50ZWQgPSBldmVudFJlY29yZC5kZWZhdWx0UHJldmVudGVkO1xuICAgIHByb3BlcnRpZXMuY29tcG9zZWQgPSBldmVudFJlY29yZC5jb21wb3NlZDtcblxuICAgIGlmIChldmVudFJlY29yZC5tYW5pZmVzdFR5cGUpIHtcbiAgICAgIHByb3BlcnRpZXMudHlwZSA9IGV2ZW50UmVjb3JkLm1hbmlmZXN0VHlwZTtcbiAgICB9XG5cbiAgICByZXR1cm4gcHJvcGVydGllcztcbiAgfVxuXG4gICNidWlsZFByb3BlcnR5VHJlZShvYmo6IFJlY29yZDxzdHJpbmcsIHVua25vd24+LCBkZXB0aCA9IDApOiBIVE1MVUxpc3RFbGVtZW50IHtcbiAgICBjb25zdCB1bCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3VsJyk7XG4gICAgdWwuY2xhc3NOYW1lID0gJ2V2ZW50LXByb3BlcnR5LXRyZWUnO1xuICAgIGlmIChkZXB0aCA+IDApIHtcbiAgICAgIHVsLmNsYXNzTGlzdC5hZGQoJ25lc3RlZCcpO1xuICAgIH1cblxuICAgIGZvciAoY29uc3QgW2tleSwgdmFsdWVdIG9mIE9iamVjdC5lbnRyaWVzKG9iaikpIHtcbiAgICAgIGNvbnN0IGxpID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnbGknKTtcbiAgICAgIGxpLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS1pdGVtJztcblxuICAgICAgY29uc3Qga2V5U3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgIGtleVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LWtleSc7XG4gICAgICBrZXlTcGFuLnRleHRDb250ZW50ID0ga2V5O1xuXG4gICAgICBjb25zdCBjb2xvblNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICBjb2xvblNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LWNvbG9uJztcbiAgICAgIGNvbG9uU3Bhbi50ZXh0Q29udGVudCA9ICc6ICc7XG5cbiAgICAgIGxpLmFwcGVuZENoaWxkKGtleVNwYW4pO1xuICAgICAgbGkuYXBwZW5kQ2hpbGQoY29sb25TcGFuKTtcblxuICAgICAgaWYgKHZhbHVlID09PSBudWxsIHx8IHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIG51bGwnO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUgYm9vbGVhbic7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ251bWJlcicpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIG51bWJlcic7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ3N0cmluZycpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIHN0cmluZyc7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IGBcIiR7dmFsdWV9XCJgO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmIChBcnJheS5pc0FycmF5KHZhbHVlKSkge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUgYXJyYXknO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBgQXJyYXkoJHt2YWx1ZS5sZW5ndGh9KWA7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG5cbiAgICAgICAgaWYgKHZhbHVlLmxlbmd0aCA+IDAgJiYgZGVwdGggPCAzKSB7XG4gICAgICAgICAgY29uc3QgbmVzdGVkT2JqOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiA9IHt9O1xuICAgICAgICAgIHZhbHVlLmZvckVhY2goKGl0ZW0sIGluZGV4KSA9PiB7XG4gICAgICAgICAgICBuZXN0ZWRPYmpbaW5kZXhdID0gaXRlbTtcbiAgICAgICAgICB9KTtcbiAgICAgICAgICBsaS5hcHBlbmRDaGlsZCh0aGlzLiNidWlsZFByb3BlcnR5VHJlZShuZXN0ZWRPYmosIGRlcHRoICsgMSkpO1xuICAgICAgICB9XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ29iamVjdCcpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIG9iamVjdCc7XG4gICAgICAgIGNvbnN0IGtleXMgPSBPYmplY3Qua2V5cyh2YWx1ZSBhcyBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPik7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IGtleXMubGVuZ3RoID4gMCA/IGBPYmplY3RgIDogYHt9YDtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcblxuICAgICAgICBpZiAoa2V5cy5sZW5ndGggPiAwICYmIGRlcHRoIDwgMykge1xuICAgICAgICAgIGxpLmFwcGVuZENoaWxkKHRoaXMuI2J1aWxkUHJvcGVydHlUcmVlKHZhbHVlIGFzIFJlY29yZDxzdHJpbmcsIHVua25vd24+LCBkZXB0aCArIDEpKTtcbiAgICAgICAgfVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlJztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH1cblxuICAgICAgdWwuYXBwZW5kQ2hpbGQobGkpO1xuICAgIH1cblxuICAgIHJldHVybiB1bDtcbiAgfVxuXG4gICNzY3JvbGxFdmVudHNUb0JvdHRvbSgpIHtcbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgcmVxdWVzdEFuaW1hdGlvbkZyYW1lKCgpID0+IHtcbiAgICAgIGNvbnN0IGxhc3RFdmVudCA9IHRoaXMuI2V2ZW50TGlzdCEubGFzdEVsZW1lbnRDaGlsZDtcbiAgICAgIGlmIChsYXN0RXZlbnQpIHtcbiAgICAgICAgbGFzdEV2ZW50LnNjcm9sbEludG9WaWV3KHsgYmVoYXZpb3I6ICdhdXRvJywgYmxvY2s6ICdlbmQnIH0pO1xuICAgICAgfVxuICAgIH0pO1xuICB9XG5cbiAgI2lzRXZlbnRzVGFiQWN0aXZlKCk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHRhYnMgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ2NlbS1wZi12Ni10YWJzJyk7XG4gICAgaWYgKCF0YWJzKSByZXR1cm4gZmFsc2U7XG5cbiAgICBjb25zdCBzZWxlY3RlZEluZGV4ID0gcGFyc2VJbnQodGFicy5nZXRBdHRyaWJ1dGUoJ3NlbGVjdGVkJykgfHwgJzAnLCAxMCk7XG4gICAgcmV0dXJuIHNlbGVjdGVkSW5kZXggPT09IDM7XG4gIH1cblxuICAjZmlsdGVyRXZlbnRzKHF1ZXJ5OiBzdHJpbmcpIHtcbiAgICB0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgZm9yIChjb25zdCBlbnRyeSBvZiB0aGlzLiNldmVudExpc3QuY2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IGV2ZW50UmVjb3JkID0gdGhpcy4jZ2V0RXZlbnRSZWNvcmRCeUlkKChlbnRyeSBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5ldmVudElkISk7XG5cbiAgICAgIGlmICghZXZlbnRSZWNvcmQpIGNvbnRpbnVlO1xuXG4gICAgICBjb25zdCB0ZXh0TWF0Y2ggPSB0aGlzLiNldmVudE1hdGNoZXNUZXh0RmlsdGVyKGV2ZW50UmVjb3JkKTtcbiAgICAgIGNvbnN0IHR5cGVNYXRjaCA9IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmhhcyhldmVudFJlY29yZC5ldmVudE5hbWUpO1xuICAgICAgY29uc3QgZWxlbWVudE1hdGNoID0gdGhpcy4jZWxlbWVudEZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNlbGVtZW50RmlsdGVycy5oYXMoZXZlbnRSZWNvcmQudGFnTmFtZSk7XG5cbiAgICAgIChlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuID0gISh0ZXh0TWF0Y2ggJiYgdHlwZU1hdGNoICYmIGVsZW1lbnRNYXRjaCk7XG4gICAgfVxuICB9XG5cbiAgI2V2ZW50TWF0Y2hlc1RleHRGaWx0ZXIoZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkKTogYm9vbGVhbiB7XG4gICAgaWYgKCF0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSkgcmV0dXJuIHRydWU7XG5cbiAgICBjb25zdCBzZWFyY2hUZXh0ID0gW1xuICAgICAgZXZlbnRSZWNvcmQudGFnTmFtZSxcbiAgICAgIGV2ZW50UmVjb3JkLmV2ZW50TmFtZSxcbiAgICAgIGV2ZW50UmVjb3JkLmVsZW1lbnRJZCB8fCAnJyxcbiAgICAgIEpTT04uc3RyaW5naWZ5KGV2ZW50UmVjb3JkLmN1c3RvbVByb3BlcnRpZXMgfHwge30pXG4gICAgXS5qb2luKCcgJykudG9Mb3dlckNhc2UoKTtcblxuICAgIHJldHVybiBzZWFyY2hUZXh0LmluY2x1ZGVzKHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKTtcbiAgfVxuXG4gICNnZXRFdmVudFJlY29yZEJ5SWQoaWQ6IHN0cmluZyk6IEV2ZW50UmVjb3JkIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy4jY2FwdHVyZWRFdmVudHMuZmluZChlID0+IGUuaWQgPT09IGlkKTtcbiAgfVxuXG4gICN1cGRhdGVFdmVudEZpbHRlcnMoKSB7XG4gICAgY29uc3Qgc2F2ZWRQcmVmZXJlbmNlcyA9IHRoaXMuI2xvYWRFdmVudEZpbHRlcnNGcm9tU3RvcmFnZSgpO1xuXG4gICAgY29uc3QgZXZlbnRUeXBlRmlsdGVyID0gdGhpcy4jJCgnZXZlbnQtdHlwZS1maWx0ZXInKTtcbiAgICBpZiAoZXZlbnRUeXBlRmlsdGVyICYmIHRoaXMuI2VsZW1lbnRFdmVudE1hcCkge1xuICAgICAgbGV0IG1lbnUgPSBldmVudFR5cGVGaWx0ZXIucXVlcnlTZWxlY3RvcignY2VtLXBmLXY2LW1lbnUnKTtcbiAgICAgIGlmICghbWVudSkge1xuICAgICAgICBtZW51ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnY2VtLXBmLXY2LW1lbnUnKTtcbiAgICAgICAgZXZlbnRUeXBlRmlsdGVyLmFwcGVuZENoaWxkKG1lbnUpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBleGlzdGluZ0l0ZW1zID0gbWVudS5xdWVyeVNlbGVjdG9yQWxsKCdjZW0tcGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICBleGlzdGluZ0l0ZW1zLmZvckVhY2goaXRlbSA9PiBpdGVtLnJlbW92ZSgpKTtcblxuICAgICAgY29uc3QgYWxsRXZlbnRUeXBlcyA9IG5ldyBTZXQ8c3RyaW5nPigpO1xuICAgICAgZm9yIChjb25zdCBbdGFnTmFtZSwgZXZlbnRJbmZvXSBvZiB0aGlzLiNlbGVtZW50RXZlbnRNYXApIHtcbiAgICAgICAgaWYgKHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5oYXModGFnTmFtZSkpIHtcbiAgICAgICAgICBmb3IgKGNvbnN0IGV2ZW50TmFtZSBvZiBldmVudEluZm8uZXZlbnROYW1lcykge1xuICAgICAgICAgICAgYWxsRXZlbnRUeXBlcy5hZGQoZXZlbnROYW1lKTtcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgaWYgKHNhdmVkUHJlZmVyZW5jZXMuZXZlbnRUeXBlcykge1xuICAgICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzID0gKHNhdmVkUHJlZmVyZW5jZXMuZXZlbnRUeXBlcyBhcyBTZXQ8c3RyaW5nPiAmIHsgaW50ZXJzZWN0aW9uOiAob3RoZXI6IFNldDxzdHJpbmc+KSA9PiBTZXQ8c3RyaW5nPiB9KS5pbnRlcnNlY3Rpb24oYWxsRXZlbnRUeXBlcyk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzID0gbmV3IFNldChhbGxFdmVudFR5cGVzKTtcbiAgICAgIH1cblxuICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgYWxsRXZlbnRUeXBlcykge1xuICAgICAgICBjb25zdCBpdGVtID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnY2VtLXBmLXY2LW1lbnUtaXRlbScpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgndmFyaWFudCcsICdjaGVja2JveCcpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgndmFsdWUnLCBldmVudE5hbWUpO1xuICAgICAgICBpZiAodGhpcy4jZXZlbnRUeXBlRmlsdGVycy5oYXMoZXZlbnROYW1lKSkge1xuICAgICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCdjaGVja2VkJywgJycpO1xuICAgICAgICB9XG4gICAgICAgIGl0ZW0udGV4dENvbnRlbnQgPSBldmVudE5hbWU7XG4gICAgICAgIG1lbnUuYXBwZW5kQ2hpbGQoaXRlbSk7XG4gICAgICB9XG4gICAgfVxuXG4gICAgY29uc3QgZWxlbWVudEZpbHRlciA9IHRoaXMuIyQoJ2VsZW1lbnQtZmlsdGVyJyk7XG4gICAgaWYgKGVsZW1lbnRGaWx0ZXIgJiYgdGhpcy4jZWxlbWVudEV2ZW50TWFwKSB7XG4gICAgICBsZXQgbWVudSA9IGVsZW1lbnRGaWx0ZXIucXVlcnlTZWxlY3RvcignY2VtLXBmLXY2LW1lbnUnKTtcbiAgICAgIGlmICghbWVudSkge1xuICAgICAgICBtZW51ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnY2VtLXBmLXY2LW1lbnUnKTtcbiAgICAgICAgZWxlbWVudEZpbHRlci5hcHBlbmRDaGlsZChtZW51KTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgZXhpc3RpbmdJdGVtcyA9IG1lbnUucXVlcnlTZWxlY3RvckFsbCgnY2VtLXBmLXY2LW1lbnUtaXRlbScpO1xuICAgICAgZXhpc3RpbmdJdGVtcy5mb3JFYWNoKGl0ZW0gPT4gaXRlbS5yZW1vdmUoKSk7XG5cbiAgICAgIGNvbnN0IGFsbEVsZW1lbnRzID0gbmV3IFNldDxzdHJpbmc+KCk7XG4gICAgICBmb3IgKGNvbnN0IHRhZ05hbWUgb2YgdGhpcy4jZWxlbWVudEV2ZW50TWFwLmtleXMoKSkge1xuICAgICAgICBpZiAodGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmhhcyh0YWdOYW1lKSkge1xuICAgICAgICAgIGFsbEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICBpZiAoc2F2ZWRQcmVmZXJlbmNlcy5lbGVtZW50cykge1xuICAgICAgICB0aGlzLiNlbGVtZW50RmlsdGVycyA9IChzYXZlZFByZWZlcmVuY2VzLmVsZW1lbnRzIGFzIFNldDxzdHJpbmc+ICYgeyBpbnRlcnNlY3Rpb246IChvdGhlcjogU2V0PHN0cmluZz4pID0+IFNldDxzdHJpbmc+IH0pLmludGVyc2VjdGlvbihhbGxFbGVtZW50cyk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICB0aGlzLiNlbGVtZW50RmlsdGVycyA9IG5ldyBTZXQoYWxsRWxlbWVudHMpO1xuICAgICAgfVxuXG4gICAgICBmb3IgKGNvbnN0IHRhZ05hbWUgb2YgYWxsRWxlbWVudHMpIHtcbiAgICAgICAgY29uc3QgaXRlbSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2NlbS1wZi12Ni1tZW51LWl0ZW0nKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3ZhcmlhbnQnLCAnY2hlY2tib3gnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3ZhbHVlJywgdGFnTmFtZSk7XG4gICAgICAgIGlmICh0aGlzLiNlbGVtZW50RmlsdGVycy5oYXModGFnTmFtZSkpIHtcbiAgICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnY2hlY2tlZCcsICcnKTtcbiAgICAgICAgfVxuICAgICAgICBpdGVtLnRleHRDb250ZW50ID0gYDwke3RhZ05hbWV9PmA7XG4gICAgICAgIG1lbnUuYXBwZW5kQ2hpbGQoaXRlbSk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI2hhbmRsZUV2ZW50VHlwZUZpbHRlckNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB7IHZhbHVlLCBjaGVja2VkIH0gPSBldmVudCBhcyBFdmVudCAmIHsgdmFsdWU6IHN0cmluZzsgY2hlY2tlZDogYm9vbGVhbiB9O1xuXG4gICAgaWYgKCF2YWx1ZSkgcmV0dXJuO1xuXG4gICAgaWYgKGNoZWNrZWQpIHtcbiAgICAgIHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuYWRkKHZhbHVlKTtcbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5kZWxldGUodmFsdWUpO1xuICAgIH1cblxuICAgIHRoaXMuI3NhdmVFdmVudEZpbHRlcnMoKTtcbiAgICB0aGlzLiNmaWx0ZXJFdmVudHModGhpcy4jZXZlbnRzRmlsdGVyVmFsdWUpO1xuICB9O1xuXG4gICNoYW5kbGVFbGVtZW50RmlsdGVyQ2hhbmdlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHsgdmFsdWUsIGNoZWNrZWQgfSA9IGV2ZW50IGFzIEV2ZW50ICYgeyB2YWx1ZTogc3RyaW5nOyBjaGVja2VkOiBib29sZWFuIH07XG5cbiAgICBpZiAoIXZhbHVlKSByZXR1cm47XG5cbiAgICBpZiAoY2hlY2tlZCkge1xuICAgICAgdGhpcy4jZWxlbWVudEZpbHRlcnMuYWRkKHZhbHVlKTtcbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy4jZWxlbWVudEZpbHRlcnMuZGVsZXRlKHZhbHVlKTtcbiAgICB9XG5cbiAgICB0aGlzLiNzYXZlRXZlbnRGaWx0ZXJzKCk7XG4gICAgdGhpcy4jZmlsdGVyRXZlbnRzKHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKTtcbiAgfTtcblxuICAjbG9hZEV2ZW50RmlsdGVyc0Zyb21TdG9yYWdlKCk6IHsgZXZlbnRUeXBlczogU2V0PHN0cmluZz4gfCBudWxsOyBlbGVtZW50czogU2V0PHN0cmluZz4gfCBudWxsIH0ge1xuICAgIGNvbnN0IHByZWZlcmVuY2VzOiB7IGV2ZW50VHlwZXM6IFNldDxzdHJpbmc+IHwgbnVsbDsgZWxlbWVudHM6IFNldDxzdHJpbmc+IHwgbnVsbCB9ID0ge1xuICAgICAgZXZlbnRUeXBlczogbnVsbCxcbiAgICAgIGVsZW1lbnRzOiBudWxsXG4gICAgfTtcblxuICAgIHRyeSB7XG4gICAgICBjb25zdCBzYXZlZEV2ZW50VHlwZXMgPSBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWV2ZW50LXR5cGUtZmlsdGVycycpO1xuICAgICAgaWYgKHNhdmVkRXZlbnRUeXBlcykge1xuICAgICAgICBwcmVmZXJlbmNlcy5ldmVudFR5cGVzID0gbmV3IFNldChKU09OLnBhcnNlKHNhdmVkRXZlbnRUeXBlcykpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBzYXZlZEVsZW1lbnRzID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1lbGVtZW50LWZpbHRlcnMnKTtcbiAgICAgIGlmIChzYXZlZEVsZW1lbnRzKSB7XG4gICAgICAgIHByZWZlcmVuY2VzLmVsZW1lbnRzID0gbmV3IFNldChKU09OLnBhcnNlKHNhdmVkRWxlbWVudHMpKTtcbiAgICAgIH1cbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICBjb25zb2xlLmRlYnVnKCdbY2VtLXNlcnZlLWNocm9tZV0gbG9jYWxTdG9yYWdlIHVuYXZhaWxhYmxlIGZvciBldmVudCBmaWx0ZXJzJyk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHByZWZlcmVuY2VzO1xuICB9XG5cbiAgI3NhdmVFdmVudEZpbHRlcnMoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdjZW0tc2VydmUtZXZlbnQtdHlwZS1maWx0ZXJzJyxcbiAgICAgICAgSlNPTi5zdHJpbmdpZnkoWy4uLnRoaXMuI2V2ZW50VHlwZUZpbHRlcnNdKSk7XG4gICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLWVsZW1lbnQtZmlsdGVycycsXG4gICAgICAgIEpTT04uc3RyaW5naWZ5KFsuLi50aGlzLiNlbGVtZW50RmlsdGVyc10pKTtcbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICAvLyBsb2NhbFN0b3JhZ2UgdW5hdmFpbGFibGUgKHByaXZhdGUgbW9kZSksIHNpbGVudGx5IGNvbnRpbnVlXG4gICAgfVxuICB9XG5cbiAgI2NsZWFyRXZlbnRzKCkge1xuICAgIHRoaXMuI2NhcHR1cmVkRXZlbnRzID0gW107XG4gICAgdGhpcy4jc2VsZWN0ZWRFdmVudElkID0gbnVsbDtcbiAgICBpZiAodGhpcy4jZXZlbnRMaXN0KSB7XG4gICAgICB0aGlzLiNldmVudExpc3QucmVwbGFjZUNoaWxkcmVuKCk7XG4gICAgfVxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEhlYWRlcikge1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIuaW5uZXJIVE1MID0gJyc7XG4gICAgfVxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEJvZHkpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5pbm5lckhUTUwgPSAnJztcbiAgICB9XG4gIH1cblxuICBhc3luYyAjY29weUV2ZW50cygpIHtcbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgY29uc3QgdmlzaWJsZUV2ZW50cyA9IEFycmF5LmZyb20odGhpcy4jZXZlbnRMaXN0LmNoaWxkcmVuKVxuICAgICAgLmZpbHRlcihlbnRyeSA9PiAhKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5oaWRkZW4pXG4gICAgICAubWFwKGVudHJ5ID0+IHtcbiAgICAgICAgY29uc3QgaWQgPSAoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuZXZlbnRJZCE7XG4gICAgICAgIHJldHVybiB0aGlzLiNnZXRFdmVudFJlY29yZEJ5SWQoaWQpO1xuICAgICAgfSlcbiAgICAgIC5maWx0ZXIoKGV2ZW50KTogZXZlbnQgaXMgRXZlbnRSZWNvcmQgPT4gISFldmVudClcbiAgICAgIC5tYXAoZXZlbnQgPT4ge1xuICAgICAgICBjb25zdCB0aW1lID0gZXZlbnQudGltZXN0YW1wLnRvTG9jYWxlVGltZVN0cmluZygpO1xuICAgICAgICBjb25zdCBlbGVtZW50ID0gZXZlbnQuZWxlbWVudElkID8gYCMke2V2ZW50LmVsZW1lbnRJZH1gIDogZXZlbnQudGFnTmFtZTtcbiAgICAgICAgY29uc3QgcHJvcHMgPSBldmVudC5jdXN0b21Qcm9wZXJ0aWVzICYmIE9iamVjdC5rZXlzKGV2ZW50LmN1c3RvbVByb3BlcnRpZXMpLmxlbmd0aCA+IDBcbiAgICAgICAgICA/IGAgLSAke0pTT04uc3RyaW5naWZ5KGV2ZW50LmN1c3RvbVByb3BlcnRpZXMpfWBcbiAgICAgICAgICA6ICcnO1xuICAgICAgICByZXR1cm4gYFske3RpbWV9XSA8JHtldmVudC50YWdOYW1lfT4gJHtlbGVtZW50fSBcXHUyMTkyICR7ZXZlbnQuZXZlbnROYW1lfSR7cHJvcHN9YDtcbiAgICAgIH0pXG4gICAgICAuam9pbignXFxuJyk7XG5cbiAgICBpZiAoIXZpc2libGVFdmVudHMpIHJldHVybjtcblxuICAgIHRyeSB7XG4gICAgICBhd2FpdCBuYXZpZ2F0b3IuY2xpcGJvYXJkLndyaXRlVGV4dCh2aXNpYmxlRXZlbnRzKTtcbiAgICAgIGNvbnN0IGJ0biA9IHRoaXMuIyQoJ2NvcHktZXZlbnRzJyk7XG4gICAgICBpZiAoYnRuKSB7XG4gICAgICAgIGNvbnN0IHRleHROb2RlID0gQXJyYXkuZnJvbShidG4uY2hpbGROb2RlcykuZmluZChcbiAgICAgICAgICBuID0+IG4ubm9kZVR5cGUgPT09IE5vZGUuVEVYVF9OT0RFICYmIChuLnRleHRDb250ZW50Py50cmltKCkubGVuZ3RoID8/IDApID4gMFxuICAgICAgICApO1xuICAgICAgICBpZiAodGV4dE5vZGUpIHtcbiAgICAgICAgICBjb25zdCBvcmlnaW5hbCA9IHRleHROb2RlLnRleHRDb250ZW50O1xuICAgICAgICAgIHRleHROb2RlLnRleHRDb250ZW50ID0gJ0NvcGllZCEnO1xuXG4gICAgICAgICAgaWYgKHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICB0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgICBpZiAodGhpcy5pc0Nvbm5lY3RlZCAmJiB0ZXh0Tm9kZS5wYXJlbnROb2RlKSB7XG4gICAgICAgICAgICAgIHRleHROb2RlLnRleHRDb250ZW50ID0gb3JpZ2luYWw7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgICB0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICAgICAgICB9LCAyMDAwKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gY2F0Y2ggKGVycikge1xuICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjb3B5IGV2ZW50czonLCBlcnIpO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cEV2ZW50TGlzdGVuZXJzKCkge1xuICAgIHRoaXMuI2V2ZW50TGlzdCA9IHRoaXMuIyQoJ2V2ZW50LWxpc3QnKTtcbiAgICB0aGlzLiNldmVudERldGFpbEhlYWRlciA9IHRoaXMuIyQoJ2V2ZW50LWRldGFpbC1oZWFkZXInKTtcbiAgICB0aGlzLiNldmVudERldGFpbEJvZHkgPSB0aGlzLiMkKCdldmVudC1kZXRhaWwtYm9keScpO1xuXG4gICAgaWYgKHRoaXMuI2V2ZW50TGlzdCkge1xuICAgICAgdGhpcy4jZXZlbnRMaXN0LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICAgIGNvbnN0IGxpc3RJdGVtID0gKGUudGFyZ2V0IGFzIEVsZW1lbnQpLmNsb3Nlc3QoJy5ldmVudC1saXN0LWl0ZW0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICAgICAgaWYgKGxpc3RJdGVtKSB7XG4gICAgICAgICAgY29uc3QgZXZlbnRJZCA9IGxpc3RJdGVtLmRhdGFzZXQuZXZlbnRJZDtcbiAgICAgICAgICBpZiAoZXZlbnRJZCkge1xuICAgICAgICAgICAgdGhpcy4jc2VsZWN0RXZlbnQoZXZlbnRJZCk7XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9KTtcbiAgICB9XG5cbiAgICBjb25zdCBldmVudHNGaWx0ZXIgPSB0aGlzLiMkKCdldmVudHMtZmlsdGVyJyk7XG4gICAgaWYgKGV2ZW50c0ZpbHRlcikge1xuICAgICAgZXZlbnRzRmlsdGVyLmFkZEV2ZW50TGlzdGVuZXIoJ2lucHV0JywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICAgIGNvbnN0IHsgdmFsdWUgPSAnJyB9ID0gZS50YXJnZXQgYXMgSFRNTElucHV0RWxlbWVudDtcbiAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2V2ZW50c0ZpbHRlckRlYm91bmNlVGltZXIhKTtcbiAgICAgICAgdGhpcy4jZXZlbnRzRmlsdGVyRGVib3VuY2VUaW1lciA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgIHRoaXMuI2ZpbHRlckV2ZW50cyh2YWx1ZSk7XG4gICAgICAgIH0sIDMwMCk7XG4gICAgICB9KTtcbiAgICB9XG5cbiAgICBjb25zdCBldmVudFR5cGVGaWx0ZXIgPSB0aGlzLiMkKCdldmVudC10eXBlLWZpbHRlcicpO1xuICAgIGlmIChldmVudFR5cGVGaWx0ZXIpIHtcbiAgICAgIGV2ZW50VHlwZUZpbHRlci5hZGRFdmVudExpc3RlbmVyKCdzZWxlY3QnLCB0aGlzLiNoYW5kbGVFdmVudFR5cGVGaWx0ZXJDaGFuZ2UgYXMgRXZlbnRMaXN0ZW5lcik7XG4gICAgfVxuXG4gICAgY29uc3QgZWxlbWVudEZpbHRlciA9IHRoaXMuIyQoJ2VsZW1lbnQtZmlsdGVyJyk7XG4gICAgaWYgKGVsZW1lbnRGaWx0ZXIpIHtcbiAgICAgIGVsZW1lbnRGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlRWxlbWVudEZpbHRlckNoYW5nZSBhcyBFdmVudExpc3RlbmVyKTtcbiAgICB9XG5cbiAgICB0aGlzLiMkKCdjbGVhci1ldmVudHMnKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB0aGlzLiNjbGVhckV2ZW50cygpO1xuICAgIH0pO1xuXG4gICAgdGhpcy4jJCgnY29weS1ldmVudHMnKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB0aGlzLiNjb3B5RXZlbnRzKCk7XG4gICAgfSk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuXG4gICAgLy8gQ2xlYW4gdXAga25vYiBsaXN0ZW5lcnNcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOnByb3BlcnR5LWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOmNzcy1wcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjphdHRyaWJ1dGUtY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOnByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG5cbiAgICAvLyBDbGVhbiB1cCB0cmVlIHN0YXRlIGxpc3RlbmVyc1xuICAgIGlmICh0aGlzLiNoYW5kbGVUcmVlRXhwYW5kKSB7XG4gICAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2V4cGFuZCcsIHRoaXMuI2hhbmRsZVRyZWVFeHBhbmQpO1xuICAgIH1cbiAgICBpZiAodGhpcy4jaGFuZGxlVHJlZUNvbGxhcHNlKSB7XG4gICAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2NvbGxhcHNlJywgdGhpcy4jaGFuZGxlVHJlZUNvbGxhcHNlKTtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QpIHtcbiAgICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlVHJlZVNlbGVjdCk7XG4gICAgfVxuXG4gICAgLy8gQ2xlYW4gdXAgd2luZG93IGxpc3RlbmVyXG4gICAgaWYgKHRoaXMuI2hhbmRsZUxvZ3NFdmVudCkge1xuICAgICAgd2luZG93LnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2NlbTpsb2dzJywgdGhpcy4jaGFuZGxlTG9nc0V2ZW50KTtcbiAgICB9XG5cbiAgICAvLyBDbGVhciBwZW5kaW5nIGZlZWRiYWNrIHRpbWVvdXRzXG4gICAgaWYgKHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgdGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgIH1cbiAgICBpZiAodGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgIHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgfVxuICAgIGlmICh0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICB0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICB9XG5cbiAgICAvLyBEaXNjb25uZWN0IG11dGF0aW9uIG9ic2VydmVyXG4gICAgdGhpcy4jb2JzZXJ2ZXIuZGlzY29ubmVjdCgpO1xuXG4gICAgLy8gQ2xvc2UgV2ViU29ja2V0IGNvbm5lY3Rpb25cbiAgICBpZiAodGhpcy4jd3NDbGllbnQpIHtcbiAgICAgIHRoaXMuI3dzQ2xpZW50LmRlc3Ryb3koKTtcbiAgICB9XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXNlcnZlLWNocm9tZSc6IENlbVNlcnZlQ2hyb21lO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBoZWlnaHQ6IDEwMHZoO1xcXFxuICBvdmVyZmxvdzogaGlkZGVuO1xcXFxuICAtLWNlbS1wZi12Ni1jLW1hc3RoZWFkX19sb2dvLS1XaWR0aDogbWF4LWNvbnRlbnQ7XFxcXG4gIC0tY2VtLXBmLXY2LWMtbWFzdGhlYWRfX3RvZ2dsZS0tU2l6ZTogMXJlbTtcXFxcbn1cXFxcblxcXFxuW2hpZGRlbl0ge1xcXFxuICBkaXNwbGF5OiBub25lICFpbXBvcnRhbnQ7XFxcXG59XFxcXG5cXFxcbi8qIE1hc3RoZWFkIGxvZ28gc3R5bGVzICovXFxcXG4ubWFzdGhlYWQtbG9nbyB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcbiAgY29sb3I6IGluaGVyaXQ7XFxcXG4gIG1heC1oZWlnaHQ6IHZhcigtLWNlbS1wZi12Ni1jLW1hc3RoZWFkX19sb2dvLS1NYXhIZWlnaHQpO1xcXFxuICBnYXA6IDRweDtcXFxcbiAgXFxcXHUwMDI2IGltZyB7XFxcXG4gICAgZGlzcGxheTogYmxvY2s7XFxcXG4gICAgbWF4LWhlaWdodDogdmFyKC0tY2VtLXBmLXY2LWMtbWFzdGhlYWRfX2xvZ28tLU1heEhlaWdodCk7XFxcXG4gICAgd2lkdGg6IGF1dG87XFxcXG4gIH1cXFxcbiAgXFxcXHUwMDI2IDo6c2xvdHRlZChbc2xvdD1cXFxcXFxcInRpdGxlXFxcXFxcXCJdKSB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICAgIGZvbnQtc2l6ZTogMS4xMjVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIH1cXFxcbiAgXFxcXHUwMDI2IGgxIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gICAgZm9udC1zaXplOiAxOHB4O1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIFRvb2xiYXIgZ3JvdXAgYWxpZ25tZW50ICovXFxcXG5jZW0tcGYtdjYtdG9vbGJhci1ncm91cFt2YXJpYW50PVxcXFxcXFwiYWN0aW9uLWdyb3VwXFxcXFxcXCJdIHtcXFxcbiAgbWFyZ2luLWlubGluZS1zdGFydDogYXV0bztcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbn1cXFxcblxcXFxuLmRlYnVnLXBhbmVsIHtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlcjogMXB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgcGFkZGluZzogMS41cmVtO1xcXFxuICBtYXgtd2lkdGg6IDYwMHB4O1xcXFxuICB3aWR0aDogOTAlO1xcXFxuICBtYXgtaGVpZ2h0OiA4MHZoO1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuXFxcXG4gIGgyIHtcXFxcbiAgICBtYXJnaW46IDAgMCAxcmVtIDA7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIGZvbnQtc2l6ZTogMS4xMjVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGRsIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gIH1cXFxcblxcXFxuICBkdCB7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBtYXJnaW4tdG9wOiAwLjVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDUwMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGRkIHtcXFxcbiAgICBtYXJnaW46IDAgMCAwLjVyZW0gMDtcXFxcbiAgICBmb250LWZhbWlseTogdWktbW9ub3NwYWNlLCAnQ2FzY2FkaWEgQ29kZScsICdTb3VyY2UgQ29kZSBQcm8nLCBNZW5sbywgQ29uc29sYXMsICdEZWphVnUgU2FucyBNb25vJywgbW9ub3NwYWNlO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICB9XFxcXG5cXFxcbiAgZGV0YWlscyB7XFxcXG4gICAgbWFyZ2luLXRvcDogMXJlbTtcXFxcblxcXFxuICAgIHN1bW1hcnkge1xcXFxuICAgICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG4gICAgICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICAgICAgZGlzcGxheTogZmxleDtcXFxcbiAgICAgIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gICAgICBnYXA6IDAuNXJlbTtcXFxcbiAgICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuXFxcXG4gICAgICBcXFxcdTAwMjY6Oi13ZWJraXQtZGV0YWlscy1tYXJrZXIge1xcXFxuICAgICAgICBkaXNwbGF5OiBub25lO1xcXFxuICAgICAgfVxcXFxuXFxcXG4gICAgICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgICAgIGNvbnRlbnQ6ICdcXFxcXFxcXDI1QjgnO1xcXFxuICAgICAgICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XFxcXG4gICAgICAgIHRyYW5zaXRpb246IHRyYW5zZm9ybSAxMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICAgIH1cXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjZbb3Blbl0gc3VtbWFyeTo6YmVmb3JlIHtcXFxcbiAgICAgIHRyYW5zZm9ybTogcm90YXRlKDkwZGVnKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBwcmUge1xcXFxuICAgICAgbWFyZ2luLXRvcDogMC41cmVtO1xcXFxuICAgICAgbWFyZ2luLWxlZnQ6IDEuNXJlbTtcXFxcbiAgICAgIHBhZGRpbmc6IDAuNXJlbTtcXFxcbiAgICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0tZGVmYXVsdCk7XFxcXG4gICAgICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gICAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgICAgb3ZlcmZsb3cteDogYXV0bztcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICAuYnV0dG9uLWNvbnRhaW5lciB7XFxcXG4gICAgZGlzcGxheTogZmxleDtcXFxcbiAgICBnYXA6IDAuNXJlbTtcXFxcbiAgICBtYXJnaW4tdG9wOiAxcmVtO1xcXFxuICB9XFxcXG5cXFxcbiAgcCB7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgfVxcXFxuXFxcXG4gIGJ1dHRvbiB7XFxcXG4gICAgbWFyZ2luLXRvcDogMXJlbTtcXFxcbiAgICBwYWRkaW5nOiAwLjVyZW0gMXJlbTtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLWJyYW5kKTtcXFxcbiAgICBib3JkZXI6IG5vbmU7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDQwMDtcXFxcbiAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gICAgdHJhbnNpdGlvbjogYWxsIDIwMG1zIGN1YmljLWJlemllcigwLjY0NSwgMC4wNDUsIDAuMzU1LCAxKTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0taG92ZXIpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBDb250ZW50IGFyZWEgcGFkZGluZyBmb3IgZGVtbyAqL1xcXFxuY2VtLXBmLXY2LXBhZ2UtbWFpbiB7XFxcXG4gIG1pbi1oZWlnaHQ6IGNhbGMoMTAwZHZoIC0gNzJweCAtIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1pbnNldC0tcGFnZS1jaHJvbWUpKTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgXFxcXHUwMDI2IFxcXFx1MDAzZSA6OnNsb3R0ZWQoOm5vdChbc2xvdD1rbm9ic10pKSB7XFxcXG4gICAgcGFkZGluZzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgICBmbGV4OiAxO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbmNlbS1kcmF3ZXIge1xcXFxuICBjZW0tcGYtdjYtdGFicyB7XFxcXG4gICAgY2VtLXBmLXY2LXRhYiB7XFxcXG4gICAgICBwYWRkaW5nLWJsb2NrLWVuZDogMDtcXFxcbiAgICB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogRWxlbWVudCBkZXNjcmlwdGlvbnMgaW4gbGlzdGluZyAqL1xcXFxuLmVsZW1lbnQtc3VtbWFyeSB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG59XFxcXG5cXFxcbi5lbGVtZW50LWRlc2NyaXB0aW9uIHtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtKTtcXFxcbn1cXFxcblxcXFxuLyogQ2FyZCBmb290ZXIgZGVtbyBuYXZpZ2F0aW9uICovXFxcXG4uY2FyZC1kZW1vcyB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtd3JhcDogd3JhcDtcXFxcbiAgZ2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tZ2FwLS1hY3Rpb24tdG8tYWN0aW9uLS1kZWZhdWx0KTtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgbWFyZ2luOiAwO1xcXFxufVxcXFxuXFxcXG4ucGFja2FnZS1uYW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG59XFxcXG5cXFxcbi8qIEtub2JzIGNvbnRhaW5lciAtIGZpbGxzIHRhYiBwYW5lbCBoZWlnaHQgKi9cXFxcbiNrbm9icy1jb250YWluZXIge1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIFxcXFx1MDAyNiA6OnNsb3R0ZWQoW3Nsb3Q9XFxcXFxcXCJrbm9ic1xcXFxcXFwiXSkge1xcXFxuICAgIGZsZXg6IDE7XFxcXG4gICAgbWluLWhlaWdodDogMDtcXFxcbiAgICBvdmVyZmxvdzogaGlkZGVuO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5rbm9icy1lbXB0eSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICB0ZXh0LWFsaWduOiBjZW50ZXI7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbGcpO1xcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBwYWRkaW5nOiAycHggNnB4O1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDNweDtcXFxcbiAgICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmluc3RhbmNlLXRhZyB7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLmluc3RhbmNlLWxhYmVsIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4ua25vYi1ncm91cCB7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbGcpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjpsYXN0LWNoaWxkIHtcXFxcbiAgICBtYXJnaW4tYm90dG9tOiAwO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIFBhdHRlcm5GbHkgdjYgZm9ybSAtIGhvcml6b250YWwgbGF5b3V0ICovXFxcXG5jZW0tcGYtdjYtZm9ybVtob3Jpem9udGFsXSBjZW0tcGYtdjYtZm9ybS1maWVsZC1ncm91cCB7XFxcXG4gIGdyaWQtY29sdW1uOiBzcGFuIDJcXFxcbn1cXFxcblxcXFxuLmtub2ItZ3JvdXAtdGl0bGUge1xcXFxuICBncmlkLWNvbHVtbjogMSAvIC0xO1xcXFxuICBtYXJnaW46IDAgMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtYmFzZSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIGJvcmRlci1ib3R0b206IDFweCBzb2xpZCB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItY29sb3IpO1xcXFxuICBwYWRkaW5nLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWNvbnRyb2wge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbn1cXFxcblxcXFxuLmtub2ItbGFiZWwge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG59XFxcXG5cXFxcbi5rbm9iLW5hbWUge1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWRlc2NyaXB0aW9uIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxuXFxcXG4gIHAge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiAzcHg7XFxcXG4gICAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICB9XFxcXG5cXFxcbiAgYSB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIHRleHQtZGVjb3JhdGlvbjogdW5kZXJsaW5lO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIHN0cm9uZyB7XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgfVxcXFxuXFxcXG4gIHVsLCBvbCB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSAwO1xcXFxuICAgIHBhZGRpbmctbGVmdDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuZm9vdGVyLnBmLW0tc3RpY2t5LWJvdHRvbSB7XFxcXG4gIHZpZXctdHJhbnNpdGlvbi1uYW1lOiBkZXYtc2VydmVyLWZvb3RlcjtcXFxcbiAgcG9zaXRpb246IHN0aWNreTtcXFxcbiAgYm90dG9tOiAwO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyLXRvcDogMXB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIHotaW5kZXg6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2UtLXNlY3Rpb24tLW0tc3RpY2t5LWJvdHRvbS0tWkluZGV4LCB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLW1kKSk7XFxcXG4gIGJveC1zaGFkb3c6IHZhcigtLWNlbS1wZi12Ni1jLXBhZ2UtLXNlY3Rpb24tLW0tc3RpY2t5LWJvdHRvbS0tQm94U2hhZG93LCB2YXIoLS1wZi10LS1nbG9iYWwtLWJveC1zaGFkb3ctLXNtLS10b3ApKTtcXFxcbn1cXFxcblxcXFxuLmZvb3Rlci1kZXNjcmlwdGlvbiB7XFxcXG4gIHBhZGRpbmc6IDEuNXJlbTtcXFxcblxcXFxuICBcXFxcdTAwMjYuZW1wdHkge1xcXFxuICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuZm9vdGVyIDo6c2xvdHRlZChbc2xvdD1cXFxcXFxcImRlc2NyaXB0aW9uXFxcXFxcXCJdKSB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjY7XFxcXG4gIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG5cXFxcbiAgY29kZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0taG92ZXIpO1xcXFxuICAgIHBhZGRpbmc6IDJweCA2cHg7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogM3B4O1xcXFxuICAgIGZvbnQtZmFtaWx5OiB1aS1tb25vc3BhY2UsICdDYXNjYWRpYSBDb2RlJywgJ1NvdXJjZSBDb2RlIFBybycsIE1lbmxvLCBDb25zb2xhcywgJ0RlamFWdSBTYW5zIE1vbm8nLCBtb25vc3BhY2U7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubG9ncy13cmFwcGVyIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4jbG9nLWNvbnRhaW5lciB7XFxcXG4gIGZsZXgtZ3JvdzogMTtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbn1cXFxcblxcXFxuLmxvZy1lbnRyeSB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxuICBhbGlnbi1pdGVtczogYmFzZWxpbmU7XFxcXG4gIGNlbS1wZi12Ni1sYWJlbCB7XFxcXG4gICAgZmxleC1zaHJpbms6IDA7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmxvZy10aW1lLFxcXFxuLmxvZy1tZXNzYWdlIHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5sb2ctdGltZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZmxleC1zaHJpbms6IDA7XFxcXG4gIGZvbnQtc2l6ZTogMTFweDtcXFxcbn1cXFxcblxcXFxuLmxvZy1tZXNzYWdlIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIHdvcmQtYnJlYWs6IGJyZWFrLXdvcmQ7XFxcXG59XFxcXG5cXFxcbi8qIE5hdmlnYXRpb24gY29udGVudCAobGlnaHQgRE9NIHNsb3R0ZWQgY29udGVudCBmb3IgY2VtLXBmLXY2LXBhZ2Utc2lkZWJhcikgKi9cXFxcbi5uYXYtcGFja2FnZSB7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNiBcXFxcdTAwM2Ugc3VtbWFyeSB7XFxcXG4gICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICAgIHBhZGRpbmc6IDAuNXJlbSAxcmVtO1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0tZGVmYXVsdCk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICBmb250LXdlaWdodDogNjAwO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kIDIwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgbWFyZ2luLWJvdHRvbTogMC41cmVtO1xcXFxuICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgICBnYXA6IDAuNXJlbTtcXFxcbiAgICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWhvdmVyKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6Oi13ZWJraXQtZGV0YWlscy1tYXJrZXIge1xcXFxuICAgICAgZGlzcGxheTogbm9uZTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgICBjb250ZW50OiAnXFxcXFxcXFwyNUI4JztcXFxcbiAgICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgICAgIHRyYW5zaXRpb246IHRyYW5zZm9ybSAxMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2W29wZW5dIFxcXFx1MDAzZSBzdW1tYXJ5OjpiZWZvcmUge1xcXFxuICAgIHRyYW5zZm9ybTogcm90YXRlKDkwZGVnKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubmF2LWVsZW1lbnQge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgbWFyZ2luLWlubGluZS1zdGFydDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG5cXFxcbiAgc3VtbWFyeSB7XFxcXG4gICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICAgIHBhZGRpbmc6IDAuNXJlbSAxcmVtO1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0tZGVmYXVsdCk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICBmb250LWZhbWlseTogdWktbW9ub3NwYWNlLCAnQ2FzY2FkaWEgQ29kZScsICdTb3VyY2UgQ29kZSBQcm8nLCBNZW5sbywgQ29uc29sYXMsICdEZWphVnUgU2FucyBNb25vJywgbW9ub3NwYWNlO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kIDIwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgZGlzcGxheTogZmxleDtcXFxcbiAgICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICAgIGdhcDogMC41cmVtO1xcXFxuICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0taG92ZXIpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNjo6LXdlYmtpdC1kZXRhaWxzLW1hcmtlciB7XFxcXG4gICAgICBkaXNwbGF5OiBub25lO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNjo6YmVmb3JlIHtcXFxcbiAgICAgIGNvbnRlbnQ6ICdcXFxcXFxcXDI1QjgnO1xcXFxuICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICAgICAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIDEwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjZbb3Blbl0gc3VtbWFyeTo6YmVmb3JlIHtcXFxcbiAgICB0cmFuc2Zvcm06IHJvdGF0ZSg5MGRlZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLm5hdi1lbGVtZW50LXRpdGxlIHtcXFxcbiAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG59XFxcXG5cXFxcbi5uYXYtZGVtby1saXN0IHtcXFxcbiAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSAwIDAgMDtcXFxcbiAgZGlzcGxheTogZ3JpZDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKTtcXFxcbn1cXFxcblxcXFxuLm5hdi1kZW1vLWxpbmsge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IGNhbGModmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgKiAyKTtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItcmFkaXVzKTtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAwLjJzLCBjb2xvciAwLjJzO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tYnJhbmQpO1xcXFxuXFxcXG4gICAgLm5hdi1wYWNrYWdlLW5hbWUge1xcXFxuICAgICAgY29sb3I6IHJnYmEoMjU1LCAyNTUsIDI1NSwgMC44KTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjZbYXJpYS1jdXJyZW50PVxcXFxcXFwicGFnZVxcXFxcXFwiXSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tYnJhbmQpO1xcXFxuXFxcXG4gICAgLm5hdi1wYWNrYWdlLW5hbWUge1xcXFxuICAgICAgY29sb3I6IHJnYmEoMjU1LCAyNTUsIDI1NSwgMC44KTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLm5hdi1wYWNrYWdlLW5hbWUge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi8qIEluZm8gYnV0dG9uIHBvcG92ZXIgdHJpZ2dlcnMgaW4ga25vYnMgLSBvdmVycmlkZSBwbGFpbiBidXR0b24gcGFkZGluZyAqL1xcXFxuY2VtLXBmLXY2LXBvcG92ZXIgY2VtLXBmLXY2LWJ1dHRvblt2YXJpYW50PVxcXFxcXFwicGxhaW5cXFxcXFxcIl0ge1xcXFxuICAtLWNlbS1wZi12Ni1jLWJ1dHRvbi0tbS1wbGFpbi0tUGFkZGluZ0lubGluZUVuZDogMDtcXFxcbiAgLS1jZW0tcGYtdjYtYy1idXR0b24tLW0tcGxhaW4tLVBhZGRpbmdJbmxpbmVTdGFydDogMDtcXFxcbiAgLS1jZW0tcGYtdjYtYy1idXR0b24tLU1pbldpZHRoOiBhdXRvO1xcXFxufVxcXFxuXFxcXG4vKiBLbm9iIGRlc2NyaXB0aW9uIGNvbnRlbnQgKHNsb3R0ZWQgaW4gZm9ybSBncm91cCBoZWxwZXIpICovXFxcXG5jZW0tcGYtdjYtZm9ybS1ncm91cCBbc2xvdD1cXFxcXFxcImhlbHBlclxcXFxcXFwiXSB7XFxcXG4gIHAge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiAzcHg7XFxcXG4gICAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICB9XFxcXG5cXFxcbiAgYSB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIHRleHQtZGVjb3JhdGlvbjogdW5kZXJsaW5lO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIHN0cm9uZyB7XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgfVxcXFxuXFxcXG4gIHVsLCBvbCB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSAwO1xcXFxuICAgIHBhZGRpbmctbGVmdDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogU3ludGF4IGhpZ2hsaWdodGluZyAoY2hyb21hIC0gdGhlbWFibGUgdmlhIENTUyBjdXN0b20gcHJvcGVydGllcykgKi9cXFxcbmNlbS1wZi12Ni1mb3JtLWdyb3VwIFtzbG90PVxcXFxcXFwiaGVscGVyXFxcXFxcXCJdIHtcXFxcbiAgLmNocm9tYSB7XFxcXG4gICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICAgIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1yYWRpdXMpO1xcXFxuICAgIG92ZXJmbG93LXg6IGF1dG87XFxcXG5cXFxcbiAgICBcXFxcdTAwMjYgLmxudGQgeyB2ZXJ0aWNhbC1hbGlnbjogdG9wOyBwYWRkaW5nOiAwOyBtYXJnaW46IDA7IGJvcmRlcjogMDsgfVxcXFxuICAgIFxcXFx1MDAyNiAubG50YWJsZSB7IGJvcmRlci1zcGFjaW5nOiAwOyBwYWRkaW5nOiAwOyBtYXJnaW46IDA7IGJvcmRlcjogMDsgfVxcXFxuICAgIFxcXFx1MDAyNiAuaGwgeyBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtaGlnaGxpZ2h0KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5sbnQsXFxcXG4gICAgXFxcXHUwMDI2IC5sbiB7XFxcXG4gICAgICB3aGl0ZS1zcGFjZTogcHJlO1xcXFxuICAgICAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG4gICAgICBtYXJnaW4tcmlnaHQ6IDAuNGVtO1xcXFxuICAgICAgcGFkZGluZzogMCAwLjRlbSAwIDAuNGVtO1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmxpbmUgeyBkaXNwbGF5OiBmbGV4OyB9XFxcXG5cXFxcbiAgICAvKiBLZXl3b3JkcyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuayxcXFxcbiAgICBcXFxcdTAwMjYgLmtjLFxcXFxuICAgIFxcXFx1MDAyNiAua2QsXFxcXG4gICAgXFxcXHUwMDI2IC5rbixcXFxcbiAgICBcXFxcdTAwMjYgLmtwLFxcXFxuICAgIFxcXFx1MDAyNiAua3Ige1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1rZXl3b3JkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmt0IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC10eXBlKTsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcblxcXFxuICAgIC8qIE5hbWVzICovXFxcXG4gICAgXFxcXHUwMDI2IC5uYSxcXFxcbiAgICBcXFxcdTAwMjYgLm5iLFxcXFxuICAgIFxcXFx1MDAyNiAubm8sXFxcXG4gICAgXFxcXHUwMDI2IC5udixcXFxcbiAgICBcXFxcdTAwMjYgLnZjLFxcXFxuICAgIFxcXFx1MDAyNiAudmcsXFxcXG4gICAgXFxcXHUwMDI2IC52aSB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LW5hbWUpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmJwIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5uYyB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtY2xhc3MpOyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuICAgIFxcXFx1MDAyNiAubmQgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWRlY29yYXRvcik7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5uaSxcXFxcbiAgICBcXFxcdTAwMjYgLnNzIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtc3BlY2lhbCk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAubmUsXFxcXG4gICAgXFxcXHUwMDI2IC5ubCB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWtleXdvcmQpO1xcXFxuICAgICAgZm9udC13ZWlnaHQ6IGJvbGQ7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAubmYgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWZ1bmN0aW9uKTsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLm5uIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5udCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtdGFnKSB9XFxcXG5cXFxcbiAgICAvKiBTdHJpbmdzICovXFxcXG4gICAgXFxcXHUwMDI2IC5zLFxcXFxuICAgIFxcXFx1MDAyNiAuc2EsXFxcXG4gICAgXFxcXHUwMDI2IC5zYixcXFxcbiAgICBcXFxcdTAwMjYgLnNjLFxcXFxuICAgIFxcXFx1MDAyNiAuZGwsXFxcXG4gICAgXFxcXHUwMDI2IC5zZCxcXFxcbiAgICBcXFxcdTAwMjYgLnMyLFxcXFxuICAgIFxcXFx1MDAyNiAuc2UsXFxcXG4gICAgXFxcXHUwMDI2IC5zaCxcXFxcbiAgICBcXFxcdTAwMjYgLnNpLFxcXFxuICAgIFxcXFx1MDAyNiAuc3gsXFxcXG4gICAgXFxcXHUwMDI2IC5zMSB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXN0cmluZyk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuc3IgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXRhZykgfVxcXFxuXFxcXG4gICAgLyogTnVtYmVycyAqL1xcXFxuICAgIFxcXFx1MDAyNiAubSxcXFxcbiAgICBcXFxcdTAwMjYgLm1iLFxcXFxuICAgIFxcXFx1MDAyNiAubWYsXFxcXG4gICAgXFxcXHUwMDI2IC5taCxcXFxcbiAgICBcXFxcdTAwMjYgLm1pLFxcXFxuICAgIFxcXFx1MDAyNiAuaWwsXFxcXG4gICAgXFxcXHUwMDI2IC5tbyB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LW51bWJlcik7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgLyogT3BlcmF0b3JzICovXFxcXG4gICAgXFxcXHUwMDI2IC5vLFxcXFxuICAgIFxcXFx1MDAyNiAub3cge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1rZXl3b3JkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIENvbW1lbnRzICovXFxcXG4gICAgXFxcXHUwMDI2IC5jLFxcXFxuICAgIFxcXFx1MDAyNiAuY2gsXFxcXG4gICAgXFxcXHUwMDI2IC5jbSxcXFxcbiAgICBcXFxcdTAwMjYgLmMxIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5jcyxcXFxcbiAgICBcXFxcdTAwMjYgLmNwLFxcXFxuICAgIFxcXFx1MDAyNiAuY3BmIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgICAgZm9udC1zdHlsZTogaXRhbGljO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIEVycm9ycyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuZXJyIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZXJyb3IpO1xcXFxuICAgICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWVycm9yLWJnKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBHZW5lcmljcyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuZ2Qge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1kZWxldGVkKTtcXFxcbiAgICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1kZWxldGVkLWJnKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5nZSB7IGZvbnQtc3R5bGU6IGl0YWxpYzsgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3IgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWVycm9yKSB9XFxcXG4gICAgXFxcXHUwMDI2IC5naCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2kge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1pbnNlcnRlZCk7XFxcXG4gICAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtaW5zZXJ0ZWQtYmcpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdvIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5ncCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3MgeyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3UgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmd0IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvcikgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2wgeyB0ZXh0LWRlY29yYXRpb246IHVuZGVybGluZTsgfVxcXFxuICAgIFxcXFx1MDAyNiAudyB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKSB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogRXZlbnRzIHRhYiBzdHlsaW5nIC0gUHJpbWFyeS1kZXRhaWwgbGF5b3V0ICovXFxcXG4uZXZlbnRzLXdyYXBwZXIge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbiNldmVudC1kcmF3ZXIge1xcXFxuICBmbGV4OiAxO1xcXFxuICBtaW4taGVpZ2h0OiAwO1xcXFxufVxcXFxuXFxcXG4vKiBFdmVudCBsaXN0IChwcmltYXJ5IHBhbmVsKSAqL1xcXFxuI2V2ZW50LWxpc3Qge1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1saXN0LWl0ZW0ge1xcXFxuICAvKiBSZXNldCBidXR0b24gc3R5bGVzICovXFxcXG4gIGFwcGVhcmFuY2U6IG5vbmU7XFxcXG4gIGJhY2tncm91bmQ6IG5vbmU7XFxcXG4gIGJvcmRlcjogbm9uZTtcXFxcbiAgYm9yZGVyLWxlZnQ6IDNweCBzb2xpZCB0cmFuc3BhcmVudDtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBmb250OiBpbmhlcml0O1xcXFxuICBjb2xvcjogaW5oZXJpdDtcXFxcbiAgdGV4dC1hbGlnbjogaW5oZXJpdDtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG5cXFxcbiAgLyogQ29tcG9uZW50IHN0eWxlcyAqL1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kIDEwMG1zIGVhc2UtaW4tb3V0LCBib3JkZXItY29sb3IgMTAwbXMgZWFzZS1pbi1vdXQ7XFxcXG5cXFxcbiAgY2VtLXBmLXY2LWxhYmVsIHtcXFxcbiAgICBmbGV4LXNocmluazogMDtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0taG92ZXIpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmZvY3VzIHtcXFxcbiAgICBvdXRsaW5lOiAycHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1jbGlja2VkKTtcXFxcbiAgICBvdXRsaW5lLW9mZnNldDogLTJweDtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5zZWxlY3RlZCB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0tc2VsZWN0ZWQpO1xcXFxuICAgIGJvcmRlci1sZWZ0LWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4uZXZlbnQtdGltZSxcXFxcbi5ldmVudC1lbGVtZW50IHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC10aW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmbGV4LXNocmluazogMDtcXFxcbiAgZm9udC1zaXplOiAxMXB4O1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZWxlbWVudCB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDQwMDtcXFxcbn1cXFxcblxcXFxuLyogRXZlbnQgZGV0YWlsIHBhbmVsICovXFxcXG4uZXZlbnQtZGV0YWlsLWhlYWRlci1jb250ZW50IHtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGJvcmRlci1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci13aWR0aCkgc29saWQgdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLWNvbG9yKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1uYW1lIHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtbGcpO1xcXFxuICBmb250LXdlaWdodDogNjAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1zdW1tYXJ5IHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxuICB3aGl0ZS1zcGFjZTogcHJlLXdyYXA7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtZGVzY3JpcHRpb24ge1xcXFxuICBtYXJnaW46IDAgMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSAwO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjU7XFxcXG4gIHdoaXRlLXNwYWNlOiBwcmUtd3JhcDtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1tZXRhIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLXRpbWUge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1lbGVtZW50IHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMtaGVhZGluZyB7XFxcXG4gIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLWJhc2UpO1xcXFxuICBmb250LXdlaWdodDogNjAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1wcm9wZXJ0aWVzIHtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXNlY29uZGFyeSk7XFxcXG4gIGJvcmRlcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXdpZHRoKSBzb2xpZCB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItY29sb3IpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItcmFkaXVzKTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IDEycHg7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjY7XFxcXG4gIG1hcmdpbjogMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LXByb3BlcnR5LXRyZWUge1xcXFxuICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBtYXJnaW46IDA7XFxcXG5cXFxcbiAgXFxcXHUwMDI2Lm5lc3RlZCB7XFxcXG4gICAgcGFkZGluZy1sZWZ0OiAxLjVlbTtcXFxcbiAgICBtYXJnaW4tdG9wOiAwLjI1ZW07XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LWl0ZW0ge1xcXFxuICBwYWRkaW5nOiAwLjEyNWVtIDA7XFxcXG59XFxcXG5cXFxcbi5wcm9wZXJ0eS1rZXkge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDUwMDtcXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LWNvbG9uIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxufVxcXFxuXFxcXG4ucHJvcGVydHktdmFsdWUge1xcXFxuICBcXFxcdTAwMjYubnVsbCxcXFxcbiAgXFxcXHUwMDI2LnVuZGVmaW5lZCB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5ib29sZWFuIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItY29sb3ItYm9vbGVhbik7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYubnVtYmVyIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItY29sb3ItbnVtYmVyKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5zdHJpbmcge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1jb2xvci1zdHJpbmcpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2LmFycmF5LFxcXFxuICBcXFxcdTAwMjYub2JqZWN0IHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNkZWJ1Zy1tb2RhbCB7XFxcXG4gIGNvbnRhaW5lci10eXBlOiBpbmxpbmUtc2l6ZTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLE1BQU0sZUFBZTtBQUMxQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjtBQUV6QixTQUFTLGlCQUFpQjs7O0FDSjFCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLDR0c0JBQTh1c0IsQ0FB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
