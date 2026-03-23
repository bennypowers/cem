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

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-serve-chrome/cem-serve-chrome.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse(`":host {\\n  display: block;\\n  height: 100vh;\\n  overflow: hidden;\\n  --pf-v6-c-masthead__logo--Width: max-content;\\n  --pf-v6-c-masthead__toggle--Size: 1rem;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n/* Masthead logo styles */\\n.masthead-logo {\\n  display: flex;\\n  align-items: center;\\n  text-decoration: none;\\n  color: inherit;\\n  max-height: var(--pf-v6-c-masthead__logo--MaxHeight);\\n  gap: 4px;\\n  \\u0026 img {\\n    display: block;\\n    max-height: var(--pf-v6-c-masthead__logo--MaxHeight);\\n    width: auto;\\n  }\\n  \\u0026 ::slotted([slot=\\"title\\"]) {\\n    margin: 0;\\n    font-size: 1.125rem;\\n    font-weight: 600;\\n    color: var(--pf-t--global--text--color--regular);\\n  }\\n  \\u0026 h1 {\\n    margin: 0;\\n    font-size: 18px;\\n  }\\n}\\n\\n/* Toolbar group alignment */\\npf-v6-toolbar-group[variant=\\"action-group\\"] {\\n  margin-inline-start: auto;\\n  align-items: center;\\n}\\n\\n.debug-panel {\\n  background: var(--pf-t--global--background--color--primary--default);\\n  border: 1px solid var(--pf-t--global--border--color--default);\\n  border-radius: 6px;\\n  padding: 1.5rem;\\n  max-width: 600px;\\n  width: 90%;\\n  max-height: 80vh;\\n  overflow-y: auto;\\n\\n  h2 {\\n    margin: 0 0 1rem 0;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-size: 1.125rem;\\n    font-weight: 600;\\n  }\\n\\n  dl {\\n    margin: 0;\\n  }\\n\\n  dt {\\n    color: var(--pf-t--global--text--color--subtle);\\n    font-size: 0.875rem;\\n    margin-top: 0.5rem;\\n    font-weight: 500;\\n  }\\n\\n  dd {\\n    margin: 0 0 0.5rem 0;\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n    color: var(--pf-t--global--text--color--regular);\\n  }\\n\\n  details {\\n    margin-top: 1rem;\\n\\n    summary {\\n      cursor: pointer;\\n      color: var(--pf-t--global--text--color--regular);\\n      font-size: 0.875rem;\\n      font-weight: 500;\\n      list-style: none;\\n      display: flex;\\n      align-items: center;\\n      gap: 0.5rem;\\n      user-select: none;\\n\\n      \\u0026::-webkit-details-marker {\\n        display: none;\\n      }\\n\\n      \\u0026::before {\\n        content: '\\\\25B8';\\n        display: inline-block;\\n        transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n        color: var(--pf-t--global--text--color--subtle);\\n      }\\n    }\\n\\n    \\u0026[open] summary::before {\\n      transform: rotate(90deg);\\n    }\\n\\n    pre {\\n      margin-top: 0.5rem;\\n      margin-left: 1.5rem;\\n      padding: 0.5rem;\\n      background: var(--pf-t--global--background--color--secondary--default);\\n      border-radius: 6px;\\n      font-size: 0.875rem;\\n      overflow-x: auto;\\n      color: var(--pf-t--global--text--color--regular);\\n    }\\n  }\\n\\n  .button-container {\\n    display: flex;\\n    gap: 0.5rem;\\n    margin-top: 1rem;\\n  }\\n\\n  p {\\n    color: var(--pf-t--global--text--color--subtle);\\n    font-size: 0.875rem;\\n  }\\n\\n  button {\\n    margin-top: 1rem;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--color--brand--default);\\n    color: var(--pf-t--global--text--color--on-brand);\\n    border: none;\\n    border-radius: 6px;\\n    font-size: 0.875rem;\\n    font-weight: 400;\\n    cursor: pointer;\\n    transition: all 200ms cubic-bezier(0.645, 0.045, 0.355, 1);\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--color--brand--hover);\\n    }\\n  }\\n}\\n\\n/* Content area padding for demo */\\npf-v6-page-main {\\n  min-height: calc(100dvh - 72px - var(--pf-t--global--spacer--inset--page-chrome));\\n  display: flex;\\n  flex-direction: column;\\n  \\u0026 \\u003e ::slotted(:not([slot=knobs])) {\\n    padding: var(--pf-t--global--spacer--lg);\\n    flex: 1;\\n  }\\n}\\n\\ncem-drawer {\\n  pf-v6-tabs {\\n    pf-v6-tab {\\n      padding-block-end: 0;\\n    }\\n  }\\n}\\n\\n/* Element descriptions in listing */\\n.element-summary {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n.element-description {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n/* Card footer demo navigation */\\n.card-demos {\\n  display: flex;\\n  flex-wrap: wrap;\\n  gap: var(--pf-t--global--spacer--gap--action-to-action--default);\\n  padding: 0;\\n  margin: 0;\\n}\\n\\n.package-name {\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n/* Knobs container - fills tab panel height */\\n#knobs-container {\\n  height: 100%;\\n  display: flex;\\n  flex-direction: column;\\n  \\u0026 ::slotted([slot=\\"knobs\\"]) {\\n    flex: 1;\\n    min-height: 0;\\n    overflow: hidden;\\n  }\\n}\\n\\n.knobs-empty {\\n  color: var(--cem-dev-server-text-muted);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  text-align: center;\\n  padding: var(--cem-dev-server-spacing-lg);\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    padding: 2px 6px;\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n}\\n\\n.instance-tag {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  color: var(--cem-dev-server-accent-color);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.instance-label {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.knob-group {\\n  margin-bottom: var(--cem-dev-server-spacing-lg);\\n\\n  \\u0026:last-child {\\n    margin-bottom: 0;\\n  }\\n}\\n\\n/* PatternFly v6 form - horizontal layout */\\npf-v6-form[horizontal] pf-v6-form-field-group {\\n  grid-column: span 2\\n}\\n\\n.knob-group-title {\\n  grid-column: 1 / -1;\\n  margin: 0 0 var(--cem-dev-server-spacing-md) 0;\\n  color: var(--cem-dev-server-text-primary);\\n  font-size: var(--cem-dev-server-font-size-base);\\n  font-weight: 600;\\n  border-bottom: 1px solid var(--cem-dev-server-border-color);\\n  padding-bottom: var(--cem-dev-server-spacing-sm);\\n}\\n\\n.knob-control {\\n  margin-bottom: var(--cem-dev-server-spacing-md);\\n}\\n\\n.knob-label {\\n  display: flex;\\n  flex-direction: column;\\n  gap: var(--cem-dev-server-spacing-xs);\\n  cursor: pointer;\\n}\\n\\n.knob-name {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-primary);\\n  font-weight: 500;\\n}\\n\\n.knob-description {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  line-height: 1.5;\\n\\n  p {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n  }\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n\\n  a {\\n    color: var(--cem-dev-server-accent-color);\\n    text-decoration: none;\\n\\n    \\u0026:hover {\\n      text-decoration: underline;\\n    }\\n  }\\n\\n  strong {\\n    font-weight: 600;\\n    color: var(--cem-dev-server-text-primary);\\n  }\\n\\n  ul, ol {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n    padding-left: var(--cem-dev-server-spacing-lg);\\n  }\\n}\\n\\nfooter.pf-m-sticky-bottom {\\n  view-transition-name: dev-server-footer;\\n  position: sticky;\\n  bottom: 0;\\n  background: var(--pf-t--global--background--color--primary--default);\\n  border-top: 1px solid var(--pf-t--global--border--color--default);\\n  z-index: var(--pf-v6-c-page--section--m-sticky-bottom--ZIndex, var(--pf-t--global--z-index--md));\\n  box-shadow: var(--pf-v6-c-page--section--m-sticky-bottom--BoxShadow, var(--pf-t--global--box-shadow--sm--top));\\n}\\n\\n.footer-description {\\n  padding: 1.5rem;\\n\\n  \\u0026.empty {\\n    display: none;\\n  }\\n}\\n\\nfooter ::slotted([slot=\\"description\\"]) {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  line-height: 1.6;\\n  font-size: 0.875rem;\\n\\n  code {\\n    background: var(--pf-t--global--background--color--primary--hover);\\n    padding: 2px 6px;\\n    border-radius: 3px;\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n  }\\n}\\n\\n.logs-wrapper {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n#log-container {\\n  flex-grow: 1;\\n  overflow-y: auto;\\n}\\n\\n.log-entry {\\n  padding: var(--cem-dev-server-spacing-xs) var(--cem-dev-server-spacing-md);\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-sm);\\n  align-items: baseline;\\n  pf-v6-label {\\n    flex-shrink: 0;\\n  }\\n}\\n\\n.log-time,\\n.log-message {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.log-time {\\n  color: var(--cem-dev-server-text-muted);\\n  flex-shrink: 0;\\n  font-size: 11px;\\n}\\n\\n.log-message {\\n  color: var(--cem-dev-server-text-primary);\\n  word-break: break-word;\\n}\\n\\n/* Navigation content (light DOM slotted content for pf-v6-page-sidebar) */\\n.nav-package {\\n  margin-bottom: var(--cem-dev-server-spacing-md);\\n\\n  \\u0026 \\u003e summary {\\n    cursor: pointer;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--background--color--secondary--default);\\n    border-radius: 6px;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-weight: 600;\\n    font-size: 0.875rem;\\n    list-style: none;\\n    transition: background 200ms cubic-bezier(0.4, 0, 0.2, 1);\\n    margin-bottom: 0.5rem;\\n    display: flex;\\n    align-items: center;\\n    gap: 0.5rem;\\n    user-select: none;\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--background--color--secondary--hover);\\n    }\\n\\n    \\u0026::-webkit-details-marker {\\n      display: none;\\n    }\\n\\n    \\u0026::before {\\n      content: '\\\\25B8';\\n      display: inline-block;\\n      transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n      color: var(--pf-t--global--text--color--subtle);\\n    }\\n  }\\n\\n  \\u0026[open] \\u003e summary::before {\\n    transform: rotate(90deg);\\n  }\\n}\\n\\n.nav-element {\\n  margin-bottom: var(--cem-dev-server-spacing-sm);\\n  margin-inline-start: var(--cem-dev-server-spacing-md);\\n\\n  summary {\\n    cursor: pointer;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--background--color--secondary--default);\\n    border-radius: 6px;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n    list-style: none;\\n    transition: background 200ms cubic-bezier(0.4, 0, 0.2, 1);\\n    display: flex;\\n    align-items: center;\\n    gap: 0.5rem;\\n    user-select: none;\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--background--color--secondary--hover);\\n    }\\n\\n    \\u0026::-webkit-details-marker {\\n      display: none;\\n    }\\n\\n    \\u0026::before {\\n      content: '\\\\25B8';\\n      display: inline-block;\\n      transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n      color: var(--pf-t--global--text--color--subtle);\\n    }\\n  }\\n\\n  \\u0026[open] summary::before {\\n    transform: rotate(90deg);\\n  }\\n}\\n\\n.nav-element-title {\\n  user-select: none;\\n}\\n\\n.nav-demo-list {\\n  list-style: none;\\n  padding: 0;\\n  margin: var(--cem-dev-server-spacing-sm) 0 0 0;\\n  display: grid;\\n  gap: var(--cem-dev-server-spacing-xs);\\n}\\n\\n.nav-demo-link {\\n  color: var(--cem-dev-server-text-primary);\\n  text-decoration: none;\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  padding-inline-start: calc(var(--cem-dev-server-spacing-md) * 2);\\n  background: var(--cem-dev-server-bg-tertiary);\\n  border-radius: var(--cem-dev-server-border-radius);\\n  display: block;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  transition: background 0.2s, color 0.2s;\\n\\n  \\u0026:hover {\\n    background: var(--cem-dev-server-accent-color);\\n    color: var(--pf-t--global--text--color--on-brand);\\n\\n    .nav-package-name {\\n      color: rgba(255, 255, 255, 0.8);\\n    }\\n  }\\n\\n  \\u0026[aria-current=\\"page\\"] {\\n    background: var(--cem-dev-server-accent-color);\\n    color: var(--pf-t--global--text--color--on-brand);\\n\\n    .nav-package-name {\\n      color: rgba(255, 255, 255, 0.8);\\n    }\\n  }\\n}\\n\\n.nav-package-name {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n/* Info button popover triggers in knobs - override plain button padding */\\npf-v6-popover pf-v6-button[variant=\\"plain\\"] {\\n  --pf-v6-c-button--m-plain--PaddingInlineEnd: 0;\\n  --pf-v6-c-button--m-plain--PaddingInlineStart: 0;\\n  --pf-v6-c-button--MinWidth: auto;\\n}\\n\\n/* Knob description content (slotted in form group helper) */\\npf-v6-form-group [slot=\\"helper\\"] {\\n  p {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n  }\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n\\n  a {\\n    color: var(--cem-dev-server-accent-color);\\n    text-decoration: none;\\n\\n    \\u0026:hover {\\n      text-decoration: underline;\\n    }\\n  }\\n\\n  strong {\\n    font-weight: 600;\\n  }\\n\\n  ul, ol {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n    padding-left: var(--cem-dev-server-spacing-lg);\\n  }\\n}\\n\\n/* Syntax highlighting (chroma - themable via CSS custom properties) */\\npf-v6-form-group [slot=\\"helper\\"] {\\n  .chroma {\\n    background-color: var(--cem-dev-server-bg-tertiary);\\n    padding: var(--cem-dev-server-spacing-sm);\\n    border-radius: var(--cem-dev-server-border-radius);\\n    overflow-x: auto;\\n\\n    \\u0026 .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }\\n    \\u0026 .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; }\\n    \\u0026 .hl { background-color: var(--cem-dev-server-syntax-highlight) }\\n    \\u0026 .lnt,\\n    \\u0026 .ln {\\n      white-space: pre;\\n      user-select: none;\\n      margin-right: 0.4em;\\n      padding: 0 0.4em 0 0.4em;\\n      color: var(--cem-dev-server-text-muted);\\n    }\\n    \\u0026 .line { display: flex; }\\n\\n    /* Keywords */\\n    \\u0026 .k,\\n    \\u0026 .kc,\\n    \\u0026 .kd,\\n    \\u0026 .kn,\\n    \\u0026 .kp,\\n    \\u0026 .kr {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n    \\u0026 .kt { color: var(--cem-dev-server-syntax-type); font-weight: bold; }\\n\\n    /* Names */\\n    \\u0026 .na,\\n    \\u0026 .nb,\\n    \\u0026 .no,\\n    \\u0026 .nv,\\n    \\u0026 .vc,\\n    \\u0026 .vg,\\n    \\u0026 .vi {\\n      color: var(--cem-dev-server-syntax-name);\\n    }\\n    \\u0026 .bp { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .nc { color: var(--cem-dev-server-syntax-class); font-weight: bold; }\\n    \\u0026 .nd { color: var(--cem-dev-server-syntax-decorator); font-weight: bold; }\\n    \\u0026 .ni,\\n    \\u0026 .ss {\\n      color: var(--cem-dev-server-syntax-special);\\n    }\\n    \\u0026 .ne,\\n    \\u0026 .nl {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n    \\u0026 .nf { color: var(--cem-dev-server-syntax-function); font-weight: bold; }\\n    \\u0026 .nn { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .nt { color: var(--cem-dev-server-syntax-tag) }\\n\\n    /* Strings */\\n    \\u0026 .s,\\n    \\u0026 .sa,\\n    \\u0026 .sb,\\n    \\u0026 .sc,\\n    \\u0026 .dl,\\n    \\u0026 .sd,\\n    \\u0026 .s2,\\n    \\u0026 .se,\\n    \\u0026 .sh,\\n    \\u0026 .si,\\n    \\u0026 .sx,\\n    \\u0026 .s1 {\\n      color: var(--cem-dev-server-syntax-string);\\n    }\\n    \\u0026 .sr { color: var(--cem-dev-server-syntax-tag) }\\n\\n    /* Numbers */\\n    \\u0026 .m,\\n    \\u0026 .mb,\\n    \\u0026 .mf,\\n    \\u0026 .mh,\\n    \\u0026 .mi,\\n    \\u0026 .il,\\n    \\u0026 .mo {\\n      color: var(--cem-dev-server-syntax-number);\\n    }\\n\\n    /* Operators */\\n    \\u0026 .o,\\n    \\u0026 .ow {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n\\n    /* Comments */\\n    \\u0026 .c,\\n    \\u0026 .ch,\\n    \\u0026 .cm,\\n    \\u0026 .c1 {\\n      color: var(--cem-dev-server-text-muted);\\n      font-style: italic;\\n    }\\n    \\u0026 .cs,\\n    \\u0026 .cp,\\n    \\u0026 .cpf {\\n      color: var(--cem-dev-server-text-muted);\\n      font-weight: bold;\\n      font-style: italic;\\n    }\\n\\n    /* Errors */\\n    \\u0026 .err {\\n      color: var(--cem-dev-server-syntax-error);\\n      background-color: var(--cem-dev-server-syntax-error-bg);\\n    }\\n\\n    /* Generics */\\n    \\u0026 .gd {\\n      color: var(--cem-dev-server-syntax-deleted);\\n      background-color: var(--cem-dev-server-syntax-deleted-bg);\\n    }\\n    \\u0026 .ge { font-style: italic; }\\n    \\u0026 .gr { color: var(--cem-dev-server-syntax-error) }\\n    \\u0026 .gh { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gi {\\n      color: var(--cem-dev-server-syntax-inserted);\\n      background-color: var(--cem-dev-server-syntax-inserted-bg);\\n    }\\n    \\u0026 .go { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gp { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gs { font-weight: bold; }\\n    \\u0026 .gu { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gt { color: var(--cem-dev-server-syntax-error) }\\n    \\u0026 .gl { text-decoration: underline; }\\n    \\u0026 .w { color: var(--cem-dev-server-text-muted) }\\n  }\\n}\\n\\n/* Events tab styling - Primary-detail layout */\\n.events-wrapper {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n#event-drawer {\\n  flex: 1;\\n  min-height: 0;\\n}\\n\\n/* Event list (primary panel) */\\n#event-list {\\n  overflow-y: auto;\\n  height: 100%;\\n}\\n\\n.event-list-item {\\n  /* Reset button styles */\\n  appearance: none;\\n  background: none;\\n  border: none;\\n  border-left: 3px solid transparent;\\n  margin: 0;\\n  font: inherit;\\n  color: inherit;\\n  text-align: inherit;\\n  width: 100%;\\n\\n  /* Component styles */\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-sm);\\n  align-items: center;\\n  cursor: pointer;\\n  transition: background 100ms ease-in-out, border-color 100ms ease-in-out;\\n\\n  pf-v6-label {\\n    flex-shrink: 0;\\n  }\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--background--color--primary--hover);\\n  }\\n\\n  \\u0026:focus {\\n    outline: 2px solid var(--pf-t--global--border--color--clicked);\\n    outline-offset: -2px;\\n  }\\n\\n  \\u0026.selected {\\n    background: var(--pf-t--global--background--color--action--plain--selected);\\n    border-left-color: var(--pf-t--global--border--color--brand--default);\\n  }\\n}\\n\\n.event-time,\\n.event-element {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.event-time {\\n  color: var(--cem-dev-server-text-muted);\\n  flex-shrink: 0;\\n  font-size: 11px;\\n}\\n\\n.event-element {\\n  color: var(--cem-dev-server-text-muted);\\n  font-weight: 400;\\n}\\n\\n/* Event detail panel */\\n.event-detail-header-content {\\n  padding: var(--cem-dev-server-spacing-md);\\n  border-bottom: var(--cem-dev-server-border-width) solid var(--cem-dev-server-border-color);\\n}\\n\\n.event-detail-name {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-lg);\\n  font-weight: 600;\\n  color: var(--cem-dev-server-text-primary);\\n}\\n\\n.event-detail-summary {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-secondary);\\n  line-height: 1.5;\\n  white-space: pre-wrap;\\n}\\n\\n.event-detail-description {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-secondary);\\n  line-height: 1.5;\\n  white-space: pre-wrap;\\n}\\n\\n.event-detail-meta {\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-md);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.event-detail-time {\\n  color: var(--cem-dev-server-text-muted);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n}\\n\\n.event-detail-element {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n}\\n\\n.event-detail-properties-heading {\\n  margin: var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  font-size: var(--cem-dev-server-font-size-base);\\n  font-weight: 600;\\n  color: var(--cem-dev-server-text-primary);\\n}\\n\\n.event-detail-properties {\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  background: var(--cem-dev-server-bg-secondary);\\n  border: var(--cem-dev-server-border-width) solid var(--cem-dev-server-border-color);\\n  border-radius: var(--cem-dev-server-border-radius);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: 12px;\\n  line-height: 1.6;\\n  margin: 0 var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md);\\n}\\n\\n.event-property-tree {\\n  list-style: none;\\n  padding: 0;\\n  margin: 0;\\n\\n  \\u0026.nested {\\n    padding-left: 1.5em;\\n    margin-top: 0.25em;\\n  }\\n}\\n\\n.property-item {\\n  padding: 0.125em 0;\\n}\\n\\n.property-key {\\n  color: var(--cem-dev-server-accent-color);\\n  font-weight: 500;\\n}\\n\\n.property-colon {\\n  color: var(--cem-dev-server-text-muted);\\n}\\n\\n.property-value {\\n  \\u0026.null,\\n  \\u0026.undefined {\\n    color: var(--cem-dev-server-text-muted);\\n    font-style: italic;\\n  }\\n\\n  \\u0026.boolean {\\n    color: var(--cem-dev-server-color-boolean);\\n  }\\n\\n  \\u0026.number {\\n    color: var(--cem-dev-server-color-number);\\n  }\\n\\n  \\u0026.string {\\n    color: var(--cem-dev-server-color-string);\\n  }\\n\\n  \\u0026.array,\\n  \\u0026.object {\\n    color: var(--cem-dev-server-text-secondary);\\n  }\\n}\\n\\n#debug-modal {\\n  container-type: inline-size;\\n}\\n"`));
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
import "../pf-v6-alert/pf-v6-alert.js";
import "../pf-v6-alert-group/pf-v6-alert-group.js";
import "../pf-v6-button/pf-v6-button.js";
import "../pf-v6-card/pf-v6-card.js";
import "../pf-v6-badge/pf-v6-badge.js";
import "../pf-v6-dropdown/pf-v6-dropdown.js";
import "../pf-v6-expandable-section/pf-v6-expandable-section.js";
import "../pf-v6-label/pf-v6-label.js";
import "../pf-v6-masthead/pf-v6-masthead.js";
import "../pf-v6-modal/pf-v6-modal.js";
import "../pf-v6-nav-group/pf-v6-nav-group.js";
import "../pf-v6-nav-item/pf-v6-nav-item.js";
import "../pf-v6-nav-link/pf-v6-nav-link.js";
import "../pf-v6-nav-list/pf-v6-nav-list.js";
import "../pf-v6-navigation/pf-v6-navigation.js";
import "../pf-v6-page-main/pf-v6-page-main.js";
import "../pf-v6-page-sidebar/pf-v6-page-sidebar.js";
import "../pf-v6-page/pf-v6-page.js";
import "../pf-v6-popover/pf-v6-popover.js";
import "../pf-v6-select/pf-v6-select.js";
import "../pf-v6-skip-to-content/pf-v6-skip-to-content.js";
import "../pf-v6-switch/pf-v6-switch.js";
import "../pf-v6-tab/pf-v6-tab.js";
import "../pf-v6-tabs/pf-v6-tabs.js";
import "../pf-v6-text-input-group/pf-v6-text-input-group.js";
import "../pf-v6-text-input/pf-v6-text-input.js";
import "../pf-v6-toggle-group-item/pf-v6-toggle-group-item.js";
import "../pf-v6-toggle-group/pf-v6-toggle-group.js";
import "../pf-v6-toolbar-group/pf-v6-toolbar-group.js";
import "../pf-v6-toolbar-item/pf-v6-toolbar-item.js";
import "../pf-v6-toolbar/pf-v6-toolbar.js";
import "../pf-v6-tree-item/pf-v6-tree-item.js";
import "../pf-v6-tree-view/pf-v6-tree-view.js";
var CEMReloadClient;
var StatePersistence;
var CemLogsEvent = class extends Event {
  logs;
  constructor(logs) {
    super("cem:logs", { bubbles: true });
    this.logs = logs;
  }
};
var _hasDescription_dec, _sidebar_dec, _tabsSelected_dec, _drawerHeight_dec, _drawer_dec, _knobs_dec, _sourceURL_dec, _canonicalURL_dec, _packageName_dec, _demoTitle_dec, _primaryTagName_dec, _a, _CemServeChrome_decorators, _demoInfoTemplate, _demoGroupTemplate, _demoListTemplate, _logEntryTemplate, _eventEntryTemplate, _init, _primaryTagName, _demoTitle, _packageName, _canonicalURL, _sourceURL, _knobs, _drawer, _drawerHeight, _tabsSelected, _sidebar, _hasDescription, _CemServeChrome_instances, $_fn, $$_fn, _logContainer, _drawerOpen, _initialLogsFetched, _isInitialLoad, _debugData, _elementEventMap, _manifest, _capturedEvents, _maxCapturedEvents, _eventList, _eventDetailHeader, _eventDetailBody, _selectedEventId, _eventsFilterValue, _eventsFilterDebounceTimer, _eventTypeFilters, _elementFilters, _discoveredElements, _handleLogsEvent, _handleTreeExpand, _handleTreeCollapse, _handleTreeSelect, _copyLogsFeedbackTimeout, _copyDebugFeedbackTimeout, _copyEventsFeedbackTimeout, _logsFilterValue, _logsFilterDebounceTimer, _logLevelFilters, _logLevelDropdown, _observer, _wsClient, _clientModulesLoaded, initWsClient_fn, renderSourceButton_fn, fetchDebugInfo_fn, populateDemoUrls_fn, setupLogListener_fn, filterLogs_fn, getLogTypeFromEntry_fn, loadLogFilterState_fn, syncCheckboxStates_fn, saveLogFilterState_fn, _handleLogFilterChange, copyLogs_fn, setupDebugOverlay_fn, setupFooterDrawer_fn, detectBrowser_fn, copyDebugInfo_fn, renderLogs_fn, getLogBadge_fn, applyLogLabelAttrs_fn, scrollLatestIntoView_fn, scrollLogsToBottom_fn, migrateFromLocalStorageIfNeeded_fn, setupColorSchemeToggle_fn, applyColorScheme_fn, setupKnobCoordination_fn, _onKnobChange, _onKnobClear, getKnobTarget_fn, getKnobTypeFromEvent_fn, getKnobTypeFromClearEvent_fn, setupTreeStatePersistence_fn, applyTreeState_fn, setupSidebarStatePersistence_fn, findTreeItemById_fn, getTreeNodeId_fn, discoverElementEvents_fn, setupEventCapture_fn, attachEventListeners_fn, observeDemoMutations_fn, _handleElementEvent, getEventDocumentation_fn, findTypeDeclaration_fn, captureEvent_fn, extractEventProperties_fn, renderEvent_fn, selectEvent_fn, buildPropertiesForDisplay_fn, buildPropertyTree_fn, scrollEventsToBottom_fn, isEventsTabActive_fn, filterEvents_fn, eventMatchesTextFilter_fn, getEventRecordById_fn, updateEventFilters_fn, _handleEventTypeFilterChange, _handleElementFilterChange, loadEventFiltersFromStorage_fn, saveEventFilters_fn, clearEvents_fn, copyEvents_fn, setupEventListeners_fn;
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
    __privateAdd(this, _clientModulesLoaded, false);
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
      <link rel="stylesheet" href="/__cem/pf-v6-c-description-list.css">
      <link rel="stylesheet" href="/__cem/pf-lightdom.css">

      <pf-v6-page ?sidebar-collapsed=${this.sidebar === "collapsed"}>
        <pf-v6-skip-to-content slot="skip-to-content"
                               href="#main-content">
          Skip to content
        </pf-v6-skip-to-content>

        <pf-v6-masthead slot="masthead">
          <a class="masthead-logo"
             href="/"
             slot="logo">
            <img alt="CEM Dev Server"
                 src="/__cem/logo.svg">
            ${this.packageName ? html`<h1>${this.packageName}</h1>` : nothing}
          </a>
          <pf-v6-toolbar slot="toolbar">
            <pf-v6-toolbar-group variant="action-group">
              ${__privateMethod(this, _CemServeChrome_instances, renderSourceButton_fn).call(this)}
              <pf-v6-toolbar-item>
                <pf-v6-button id="debug-info"
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
                </pf-v6-button>
              </pf-v6-toolbar-item>
              <pf-v6-toolbar-item>
                <pf-v6-toggle-group class="color-scheme-toggle"
                                    aria-label="Color scheme">
                  <pf-v6-toggle-group-item value="light">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="Light mode">
                      <path d="M8 11a3 3 0 1 1 0-6 3 3 0 0 1 0 6zm0 1a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z"/>
                    </svg>
                  </pf-v6-toggle-group-item>
                  <pf-v6-toggle-group-item value="dark">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="Dark mode">
                      <path d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278zM4.858 1.311A7.269 7.269 0 0 0 1.025 7.71c0 4.02 3.279 7.276 7.319 7.276a7.316 7.316 0 0 0 5.205-2.162c-.337.042-.68.063-1.029.063-4.61 0-8.343-3.714-8.343-8.29 0-1.167.242-2.278.681-3.286z"/>
                    </svg>
                  </pf-v6-toggle-group-item>
                  <pf-v6-toggle-group-item value="system">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="System preference">
                      <path d="M0 1.5A1.5 1.5 0 0 1 1.5 0h13A1.5 1.5 0 0 1 16 1.5v8a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 9.5v-8zM1.5 1a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-8a.5.5 0 0 0-.5-.5h-13z"/>
                      <path d="M2.5 12h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1 0-1zm0 2h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1 0-1z"/>
                    </svg>
                  </pf-v6-toggle-group-item>
                </pf-v6-toggle-group>
              </pf-v6-toolbar-item>
            </pf-v6-toolbar-group>
          </pf-v6-toolbar>
        </pf-v6-masthead>

        <pf-v6-page-sidebar slot="sidebar"
                            ?expanded=${this.sidebar === "expanded"}
                            ?collapsed=${this.sidebar !== "expanded"}>
          <slot name="navigation"></slot>
        </pf-v6-page-sidebar>

        <pf-v6-page-main slot="main" id="main-content">
          <slot></slot>
          <footer class="pf-m-sticky-bottom">
            <div class="footer-description${this.hasDescription ? "" : " empty"}">
              <slot name="description"></slot>
            </div>
            <cem-drawer ?open=${this.drawer === "expanded"}
                        drawer-height="${this.drawerHeight || "400"}">
              <pf-v6-tabs selected="${this.tabsSelected || "0"}">
                <pf-v6-tab title="Knobs">
                  <div id="knobs-container">
                    <slot name="knobs">
                      <p class="knobs-empty">No knobs available for this element.</p>
                    </slot>
                  </div>
                </pf-v6-tab>
                <pf-v6-tab title="Manifest Browser">
                  <cem-manifest-browser>
                    <slot name="manifest-tree" slot="manifest-tree"></slot>
                    <slot name="manifest-name" slot="manifest-name"></slot>
                    <slot name="manifest-details" slot="manifest-details"></slot>
                  </cem-manifest-browser>
                </pf-v6-tab>
                <pf-v6-tab title="Server Logs">
                  <div class="logs-wrapper">
                    <pf-v6-toolbar sticky>
                      <pf-v6-toolbar-group>
                        <pf-v6-toolbar-item>
                          <pf-v6-text-input-group id="logs-filter"
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
                          </pf-v6-text-input-group>
                        </pf-v6-toolbar-item>
                        <pf-v6-toolbar-item>
                          <pf-v6-dropdown id="log-level-filter"
                                          label="Filter log levels">
                            <span slot="toggle-text">Log Levels</span>
                            <pf-v6-menu-item variant="checkbox" value="info" checked>Info</pf-v6-menu-item>
                            <pf-v6-menu-item variant="checkbox" value="warn" checked>Warnings</pf-v6-menu-item>
                            <pf-v6-menu-item variant="checkbox" value="error" checked>Errors</pf-v6-menu-item>
                            <pf-v6-menu-item variant="checkbox" value="debug" checked>Debug</pf-v6-menu-item>
                          </pf-v6-dropdown>
                        </pf-v6-toolbar-item>
                      </pf-v6-toolbar-group>
                      <pf-v6-toolbar-group variant="action-group">
                        <pf-v6-toolbar-item>
                          <pf-v6-button id="copy-logs"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M13 0H6a2 2 0 0 0-2 2 2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2 2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm0 13V4a2 2 0 0 0-2-2H5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1zM3 13V4a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>
                            </svg>
                            Copy Logs
                          </pf-v6-button>
                        </pf-v6-toolbar-item>
                      </pf-v6-toolbar-group>
                    </pf-v6-toolbar>
                    <div id="log-container"></div>
                  </div>
                </pf-v6-tab>
                <pf-v6-tab title="Events">
                  <div class="events-wrapper">
                    <pf-v6-toolbar sticky>
                      <pf-v6-toolbar-group>
                        <pf-v6-toolbar-item>
                          <pf-v6-text-input-group id="events-filter"
                                                  placeholder="Filter events..."
                                                  icon>
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z"/>
                            </svg>
                          </pf-v6-text-input-group>
                        </pf-v6-toolbar-item>
                        <pf-v6-toolbar-item>
                          <pf-v6-dropdown id="event-type-filter"
                                          label="Filter event types">
                            <span slot="toggle-text">Event Types</span>
                          </pf-v6-dropdown>
                        </pf-v6-toolbar-item>
                        <pf-v6-toolbar-item>
                          <pf-v6-dropdown id="element-filter"
                                          label="Filter elements">
                            <span slot="toggle-text">Elements</span>
                          </pf-v6-dropdown>
                        </pf-v6-toolbar-item>
                      </pf-v6-toolbar-group>
                      <pf-v6-toolbar-group variant="action-group">
                        <pf-v6-toolbar-item>
                          <pf-v6-button id="clear-events"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                              <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                            </svg>
                            Clear Events
                          </pf-v6-button>
                        </pf-v6-toolbar-item>
                        <pf-v6-toolbar-item>
                          <pf-v6-button id="copy-events"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M13 0H6a2 2 0 0 0-2 2 2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2 2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm0 13V4a2 2 0 0 0-2-2H5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1zM3 13V4a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>
                            </svg>
                            Copy Events
                          </pf-v6-button>
                        </pf-v6-toolbar-item>
                      </pf-v6-toolbar-group>
                    </pf-v6-toolbar>
                    <pf-v6-drawer id="event-drawer" expanded>
                      <div id="event-list"></div>
                      <div id="event-detail-header" slot="panel-header"></div>
                      <div id="event-detail-body" slot="panel-body"></div>
                    </pf-v6-drawer>
                  </div>
                </pf-v6-tab>
                <pf-v6-tab title="Health">
                  <cem-health-panel component=${ifDefined(this.primaryTagName)}>
                  </cem-health-panel>
                </pf-v6-tab>
              </pf-v6-tabs>
            </cem-drawer>
          </footer>
        </pf-v6-page-main>
      </pf-v6-page>

      <pf-v6-modal id="debug-modal" variant="large">
        <h2 slot="header">Debug Information</h2>
        <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Server Version</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-version">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Server OS</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-os">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Watch Directory</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-watch-dir">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Manifest Size</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-manifest-size">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Demos Found</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-demo-count">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Tag Name</dt>
            <dd class="pf-v6-c-description-list__description">${this.primaryTagName || "-"}</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Demo Title</dt>
            <dd class="pf-v6-c-description-list__description">${this.demoTitle || "-"}</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Browser</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-browser">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">User Agent</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-ua">-</dd>
          </div>
        </dl>
        <div id="demo-urls-container"></div>
        <pf-v6-expandable-section id="debug-importmap-details"
                                  toggle-text="Show Import Map">
          <pre id="debug-importmap">-</pre>
        </pf-v6-expandable-section>
        <div slot="footer" class="button-container">
          <pf-v6-button class="debug-copy" variant="primary">
            Copy Debug Info
          </pf-v6-button>
          <pf-v6-button class="debug-close" variant="secondary">
            Close
          </pf-v6-button>
        </div>
      </pf-v6-modal>

      <!-- Reconnection modal -->
      <pf-v6-modal id="reconnection-modal" variant="large">
        <h2 slot="header">Development Server Disconnected</h2>
        <cem-reconnection-content id="reconnection-content"></cem-reconnection-content>
        <pf-v6-button id="reload-button"
                      slot="footer"
                      variant="primary">Reload Page</pf-v6-button>
        <pf-v6-button id="retry-button"
                      slot="footer"
                      variant="secondary">Retry Now</pf-v6-button>
      </pf-v6-modal>

      <!-- Transform error overlay -->
      <cem-transform-error-overlay id="error-overlay">
      </cem-transform-error-overlay>
    `;
  }
  async connectedCallback() {
    if (!__privateGet(this, _clientModulesLoaded)) {
      [{ CEMReloadClient }, { StatePersistence }] = await Promise.all([
        // @ts-ignore -- plain JS modules served at runtime by Go server
        import("/__cem/websocket-client.js"),
        // @ts-ignore
        import("/__cem/state-persistence.js")
      ]);
      import("/__cem/health-badges.js").catch((e) => console.error("[cem-serve] Failed to load health-badges:", e));
      __privateSet(this, _clientModulesLoaded, true);
    }
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
_clientModulesLoaded = new WeakMap();
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
      <pf-v6-toolbar-item>
        <pf-v6-button href="${this.sourceURL}"
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
        </pf-v6-button>
      </pf-v6-toolbar-item>
    `;
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
  const menuItems = __privateGet(this, _logLevelDropdown).querySelectorAll("pf-v6-menu-item");
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
  const tabs = this.shadowRoot?.querySelector("pf-v6-tabs");
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
  const items = toggleGroup.querySelectorAll("pf-v6-toggle-group-item");
  items.forEach((item) => {
    if (item.value === state.colorScheme) {
      item.setAttribute("selected", "");
    }
  });
  toggleGroup.addEventListener("pf-v6-toggle-group-change", (e) => {
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
    if (e.target?.tagName !== "PF-V6-TREE-ITEM") return;
    const nodeId = __privateMethod(this, _CemServeChrome_instances, getTreeNodeId_fn).call(this, e.target);
    const treeState = StatePersistence.getTreeState();
    if (!treeState.expanded.includes(nodeId)) {
      treeState.expanded.push(nodeId);
      StatePersistence.setTreeState(treeState);
    }
  });
  this.addEventListener("expand", __privateGet(this, _handleTreeExpand));
  __privateSet(this, _handleTreeCollapse, (e) => {
    if (e.target?.tagName !== "PF-V6-TREE-ITEM") return;
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
    if (e.target?.tagName !== "PF-V6-TREE-ITEM") return;
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
  const page = this.shadowRoot?.querySelector("pf-v6-page");
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
  let selector = `pf-v6-tree-item[data-type="${CSS.escape(type)}"]`;
  if (modulePath) {
    const escapedModulePath = CSS.escape(modulePath);
    const escapedType = CSS.escape(type);
    const selector1 = `pf-v6-tree-item[data-type="${escapedType}"][data-module-path="${escapedModulePath}"]${attrSuffix}`;
    const selector2 = `pf-v6-tree-item[data-type="${escapedType}"][data-path="${escapedModulePath}"]${attrSuffix}`;
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
  const tabs = this.shadowRoot?.querySelector("pf-v6-tabs");
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
    let menu = eventTypeFilter.querySelector("pf-v6-menu");
    if (!menu) {
      menu = document.createElement("pf-v6-menu");
      eventTypeFilter.appendChild(menu);
    }
    const existingItems = menu.querySelectorAll("pf-v6-menu-item");
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
      const item = document.createElement("pf-v6-menu-item");
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
    let menu = elementFilter.querySelector("pf-v6-menu");
    if (!menu) {
      menu = document.createElement("pf-v6-menu");
      elementFilter.appendChild(menu);
    }
    const existingItems = menu.querySelectorAll("pf-v6-menu-item");
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
      const item = document.createElement("pf-v6-menu-item");
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
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group">
          <dt class="pf-v6-c-description-list__term">Tag Name</dt>
          <dd class="pf-v6-c-description-list__description" data-field="tagName"></dd>
        </div>
        <div class="pf-v6-c-description-list__group" data-field-group="description">
          <dt class="pf-v6-c-description-list__term">Description</dt>
          <dd class="pf-v6-c-description-list__description" data-field="description"></dd>
        </div>
        <div class="pf-v6-c-description-list__group" data-field-group="filepath">
          <dt class="pf-v6-c-description-list__term">File Path</dt>
          <dd class="pf-v6-c-description-list__description" data-field="filepath"></dd>
        </div>
        <div class="pf-v6-c-description-list__group">
          <dt class="pf-v6-c-description-list__term">Canonical URL</dt>
          <dd class="pf-v6-c-description-list__description" data-field="canonicalURL"></dd>
        </div>
        <div class="pf-v6-c-description-list__group">
          <dt class="pf-v6-c-description-list__term">Local Route</dt>
          <dd class="pf-v6-c-description-list__description" data-field="localRoute"></dd>
        </div>
      </dl>`;
  return t;
})());
__privateAdd(_CemServeChrome, _demoGroupTemplate, (() => {
  const t = document.createElement("template");
  t.innerHTML = `
      <div class="pf-v6-c-description-list__group">
        <dt class="pf-v6-c-description-list__term" data-field="tagName"></dt>
        <dd class="pf-v6-c-description-list__description">
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
      <pf-v6-expandable-section id="debug-demos-section"
                                toggle-text="Show Demos Info">
        <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact" data-container="groups"></dl>
      </pf-v6-expandable-section>`;
  return t;
})());
__privateAdd(_CemServeChrome, _logEntryTemplate, (() => {
  const t = document.createElement("template");
  t.innerHTML = `
      <div class="log-entry" data-field="container">
        <pf-v6-label compact data-field="label"></pf-v6-label>
        <time class="log-time" data-field="time"></time>
        <span class="log-message" data-field="message"></span>
      </div>`;
  return t;
})());
__privateAdd(_CemServeChrome, _eventEntryTemplate, (() => {
  const t = document.createElement("template");
  t.innerHTML = `
      <button class="event-list-item" data-field="container">
        <pf-v6-label compact data-field="label"></pf-v6-label>
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcbmltcG9ydCB7IGlmRGVmaW5lZCB9IGZyb20gJ2xpdC9kaXJlY3RpdmVzL2lmLWRlZmluZWQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWNocm9tZS5jc3MnO1xuXG5pbXBvcnQgJy4uL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlLmpzJztcbmltcG9ydCAnLi4vY2VtLWRyYXdlci9jZW0tZHJhd2VyLmpzJztcbmltcG9ydCAnLi4vY2VtLWhlYWx0aC1wYW5lbC9jZW0taGVhbHRoLXBhbmVsLmpzJztcbmltcG9ydCAnLi4vY2VtLW1hbmlmZXN0LWJyb3dzZXIvY2VtLW1hbmlmZXN0LWJyb3dzZXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcmVjb25uZWN0aW9uLWNvbnRlbnQvY2VtLXJlY29ubmVjdGlvbi1jb250ZW50LmpzJztcbmltcG9ydCAnLi4vY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8uanMnO1xuaW1wb3J0ICcuLi9jZW0tc2VydmUta25vYi1ncm91cC9jZW0tc2VydmUta25vYi1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1zZXJ2ZS1rbm9icy9jZW0tc2VydmUta25vYnMuanMnO1xuaW1wb3J0ICcuLi9jZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXkvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtYWxlcnQvcGYtdjYtYWxlcnQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1hbGVydC1ncm91cC9wZi12Ni1hbGVydC1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWJ1dHRvbi9wZi12Ni1idXR0b24uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1jYXJkL3BmLXY2LWNhcmQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1iYWRnZS9wZi12Ni1iYWRnZS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWRyb3Bkb3duL3BmLXY2LWRyb3Bkb3duLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uL3BmLXY2LWV4cGFuZGFibGUtc2VjdGlvbi5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWxhYmVsL3BmLXY2LWxhYmVsLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbWFzdGhlYWQvcGYtdjYtbWFzdGhlYWQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1tb2RhbC9wZi12Ni1tb2RhbC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW5hdi1ncm91cC9wZi12Ni1uYXYtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXYtaXRlbS9wZi12Ni1uYXYtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW5hdi1saW5rL3BmLXY2LW5hdi1saW5rLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbmF2LWxpc3QvcGYtdjYtbmF2LWxpc3QuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXZpZ2F0aW9uL3BmLXY2LW5hdmlnYXRpb24uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1wYWdlLW1haW4vcGYtdjYtcGFnZS1tYWluLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtcGFnZS1zaWRlYmFyL3BmLXY2LXBhZ2Utc2lkZWJhci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXBhZ2UvcGYtdjYtcGFnZS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXBvcG92ZXIvcGYtdjYtcG9wb3Zlci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXNlbGVjdC9wZi12Ni1zZWxlY3QuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1za2lwLXRvLWNvbnRlbnQvcGYtdjYtc2tpcC10by1jb250ZW50LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtc3dpdGNoL3BmLXY2LXN3aXRjaC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRhYi9wZi12Ni10YWIuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10YWJzL3BmLXY2LXRhYnMuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10ZXh0LWlucHV0LWdyb3VwL3BmLXY2LXRleHQtaW5wdXQtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10ZXh0LWlucHV0L3BmLXY2LXRleHQtaW5wdXQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvZ2dsZS1ncm91cC9wZi12Ni10b2dnbGUtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10b29sYmFyLWdyb3VwL3BmLXY2LXRvb2xiYXItZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10b29sYmFyLWl0ZW0vcGYtdjYtdG9vbGJhci1pdGVtLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdG9vbGJhci9wZi12Ni10b29sYmFyLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdHJlZS1pdGVtL3BmLXY2LXRyZWUtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRyZWUtdmlldy9wZi12Ni10cmVlLXZpZXcuanMnO1xuXG4vLyBDbGllbnQtb25seSBtb2R1bGVzIGxvYWRlZCBkeW5hbWljYWxseSB0byBhdm9pZCBicmVha2luZyBTU1IuXG4vLyBUaGVzZSBhcmUgcGxhaW4gSlMgbW9kdWxlcyBzZXJ2ZWQgYXQgcnVudGltZSBieSB0aGUgR28gc2VydmVyLlxudHlwZSBDRU1SZWxvYWRDbGllbnRUeXBlID0geyBuZXcob3B0czogYW55KTogYW55IH07XG50eXBlIFN0YXRlUGVyc2lzdGVuY2VUeXBlID0ge1xuICBnZXRTdGF0ZSgpOiBhbnk7XG4gIHVwZGF0ZVN0YXRlKHM6IGFueSk6IHZvaWQ7XG4gIGdldFRyZWVTdGF0ZSgpOiBhbnk7XG4gIHNldFRyZWVTdGF0ZShzOiBhbnkpOiB2b2lkO1xuICBtaWdyYXRlRnJvbUxvY2FsU3RvcmFnZSgpOiB2b2lkO1xufTtcbmxldCBDRU1SZWxvYWRDbGllbnQ6IENFTVJlbG9hZENsaWVudFR5cGU7XG5sZXQgU3RhdGVQZXJzaXN0ZW5jZTogU3RhdGVQZXJzaXN0ZW5jZVR5cGU7XG5cbmludGVyZmFjZSBFdmVudEluZm8ge1xuICBldmVudE5hbWVzOiBTZXQ8c3RyaW5nPjtcbiAgZXZlbnRzOiBBcnJheTx7IG5hbWU6IHN0cmluZzsgdHlwZT86IHsgdGV4dDogc3RyaW5nIH07IHN1bW1hcnk/OiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH0+O1xufVxuXG5pbnRlcmZhY2UgRXZlbnRSZWNvcmQge1xuICBpZDogc3RyaW5nO1xuICB0aW1lc3RhbXA6IERhdGU7XG4gIGV2ZW50TmFtZTogc3RyaW5nO1xuICB0YWdOYW1lOiBzdHJpbmc7XG4gIGVsZW1lbnRJZDogc3RyaW5nIHwgbnVsbDtcbiAgZWxlbWVudENsYXNzOiBzdHJpbmcgfCBudWxsO1xuICBjdXN0b21Qcm9wZXJ0aWVzOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPjtcbiAgbWFuaWZlc3RUeXBlOiBzdHJpbmcgfCBudWxsO1xuICBzdW1tYXJ5OiBzdHJpbmcgfCBudWxsO1xuICBkZXNjcmlwdGlvbjogc3RyaW5nIHwgbnVsbDtcbiAgYnViYmxlczogYm9vbGVhbjtcbiAgY29tcG9zZWQ6IGJvb2xlYW47XG4gIGNhbmNlbGFibGU6IGJvb2xlYW47XG4gIGRlZmF1bHRQcmV2ZW50ZWQ6IGJvb2xlYW47XG59XG5cbmludGVyZmFjZSBEZWJ1Z0RhdGEge1xuICB2ZXJzaW9uPzogc3RyaW5nO1xuICBvcz86IHN0cmluZztcbiAgd2F0Y2hEaXI/OiBzdHJpbmc7XG4gIG1hbmlmZXN0U2l6ZT86IHN0cmluZztcbiAgZGVtb0NvdW50PzogbnVtYmVyO1xuICBkZW1vcz86IEFycmF5PHtcbiAgICB0YWdOYW1lOiBzdHJpbmc7XG4gICAgZGVzY3JpcHRpb24/OiBzdHJpbmc7XG4gICAgZmlsZXBhdGg/OiBzdHJpbmc7XG4gICAgY2Fub25pY2FsVVJMOiBzdHJpbmc7XG4gICAgbG9jYWxSb3V0ZTogc3RyaW5nO1xuICB9PjtcbiAgaW1wb3J0TWFwPzogUmVjb3JkPHN0cmluZywgdW5rbm93bj47XG4gIGltcG9ydE1hcEpTT04/OiBzdHJpbmc7XG59XG5cbmludGVyZmFjZSBNYW5pZmVzdCB7XG4gIG1vZHVsZXM/OiBBcnJheTx7XG4gICAgZGVjbGFyYXRpb25zPzogQXJyYXk8e1xuICAgICAgY3VzdG9tRWxlbWVudD86IGJvb2xlYW47XG4gICAgICB0YWdOYW1lPzogc3RyaW5nO1xuICAgICAgbmFtZT86IHN0cmluZztcbiAgICAgIGtpbmQ/OiBzdHJpbmc7XG4gICAgICBldmVudHM/OiBBcnJheTx7IG5hbWU6IHN0cmluZzsgdHlwZT86IHsgdGV4dDogc3RyaW5nIH07IHN1bW1hcnk/OiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH0+O1xuICAgIH0+O1xuICB9Pjtcbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBsb2dzIGFyZSByZWNlaXZlZFxuICovXG5leHBvcnQgY2xhc3MgQ2VtTG9nc0V2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT47XG4gIGNvbnN0cnVjdG9yKGxvZ3M6IEFycmF5PHsgdHlwZTogc3RyaW5nOyBkYXRlOiBzdHJpbmc7IG1lc3NhZ2U6IHN0cmluZyB9Pikge1xuICAgIHN1cGVyKCdjZW06bG9ncycsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLmxvZ3MgPSBsb2dzO1xuICB9XG59XG5cbi8qKlxuICogQ0VNIFNlcnZlIENocm9tZSAtIE1haW4gZGVtbyB3cmFwcGVyIGNvbXBvbmVudFxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBkZW1vIGNvbnRlbnRcbiAqIEBzbG90IG5hdmlnYXRpb24gLSBOYXZpZ2F0aW9uIHNpZGViYXIgY29udGVudFxuICogQHNsb3Qga25vYnMgLSBLbm9iIGNvbnRyb2xzXG4gKiBAc2xvdCBkZXNjcmlwdGlvbiAtIERlbW8gZGVzY3JpcHRpb25cbiAqIEBzbG90IG1hbmlmZXN0LXRyZWUgLSBNYW5pZmVzdCB0cmVlIHZpZXdcbiAqIEBzbG90IG1hbmlmZXN0LW5hbWUgLSBNYW5pZmVzdCBuYW1lIGRpc3BsYXlcbiAqIEBzbG90IG1hbmlmZXN0LWRldGFpbHMgLSBNYW5pZmVzdCBkZXRhaWxzIGRpc3BsYXlcbiAqXG4gKiBAY3VzdG9tRWxlbWVudCBjZW0tc2VydmUtY2hyb21lXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tc2VydmUtY2hyb21lJylcbmV4cG9ydCBjbGFzcyBDZW1TZXJ2ZUNocm9tZSBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIC8vIFN0YXRpYyB0ZW1wbGF0ZXMgZm9yIERPTSBjbG9uaW5nIChsb2dzLCBldmVudHMsIGRlYnVnIGluZm8pXG4gIHN0YXRpYyAjZGVtb0luZm9UZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8aDM+RGVtbyBJbmZvcm1hdGlvbjwvaDM+XG4gICAgICA8ZGwgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3QgcGYtbS1ob3Jpem9udGFsIHBmLW0tY29tcGFjdFwiPlxuICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlRhZyBOYW1lPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cInRhZ05hbWVcIj48L2RkPlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIiBkYXRhLWZpZWxkLWdyb3VwPVwiZGVzY3JpcHRpb25cIj5cbiAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5EZXNjcmlwdGlvbjwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJkZXNjcmlwdGlvblwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiIGRhdGEtZmllbGQtZ3JvdXA9XCJmaWxlcGF0aFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkZpbGUgUGF0aDwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJmaWxlcGF0aFwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkNhbm9uaWNhbCBVUkw8L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+TG9jYWwgUm91dGU8L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwibG9jYWxSb3V0ZVwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9kbD5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjZGVtb0dyb3VwVGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCIgZGF0YS1maWVsZD1cInRhZ05hbWVcIj48L2R0PlxuICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCI+XG4gICAgICAgICAgPHNwYW4gZGF0YS1maWVsZD1cImRlc2NyaXB0aW9uXCI+PC9zcGFuPjxicj5cbiAgICAgICAgICA8c21hbGwgZGF0YS1maWVsZC1ncm91cD1cImZpbGVwYXRoXCI+RmlsZTogPHNwYW4gZGF0YS1maWVsZD1cImZpbGVwYXRoXCI+PC9zcGFuPjwvc21hbGw+XG4gICAgICAgICAgPHNtYWxsPkNhbm9uaWNhbDogPHNwYW4gZGF0YS1maWVsZD1cImNhbm9uaWNhbFVSTFwiPjwvc3Bhbj48L3NtYWxsPjxicj5cbiAgICAgICAgICA8c21hbGw+TG9jYWw6IDxzcGFuIGRhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCI+PC9zcGFuPjwvc21hbGw+XG4gICAgICAgIDwvZGQ+XG4gICAgICA8L2Rpdj5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjZGVtb0xpc3RUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8cGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uIGlkPVwiZGVidWctZGVtb3Mtc2VjdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHRvZ2dsZS10ZXh0PVwiU2hvdyBEZW1vcyBJbmZvXCI+XG4gICAgICAgIDxkbCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdCBwZi1tLWhvcml6b250YWwgcGYtbS1jb21wYWN0XCIgZGF0YS1jb250YWluZXI9XCJncm91cHNcIj48L2RsPlxuICAgICAgPC9wZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24+YDtcbiAgICByZXR1cm4gdDtcbiAgfSkoKTtcblxuICBzdGF0aWMgI2xvZ0VudHJ5VGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPGRpdiBjbGFzcz1cImxvZy1lbnRyeVwiIGRhdGEtZmllbGQ9XCJjb250YWluZXJcIj5cbiAgICAgICAgPHBmLXY2LWxhYmVsIGNvbXBhY3QgZGF0YS1maWVsZD1cImxhYmVsXCI+PC9wZi12Ni1sYWJlbD5cbiAgICAgICAgPHRpbWUgY2xhc3M9XCJsb2ctdGltZVwiIGRhdGEtZmllbGQ9XCJ0aW1lXCI+PC90aW1lPlxuICAgICAgICA8c3BhbiBjbGFzcz1cImxvZy1tZXNzYWdlXCIgZGF0YS1maWVsZD1cIm1lc3NhZ2VcIj48L3NwYW4+XG4gICAgICA8L2Rpdj5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjZXZlbnRFbnRyeVRlbXBsYXRlID0gKCgpID0+IHtcbiAgICBjb25zdCB0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGVtcGxhdGUnKTtcbiAgICB0LmlubmVySFRNTCA9IGBcbiAgICAgIDxidXR0b24gY2xhc3M9XCJldmVudC1saXN0LWl0ZW1cIiBkYXRhLWZpZWxkPVwiY29udGFpbmVyXCI+XG4gICAgICAgIDxwZi12Ni1sYWJlbCBjb21wYWN0IGRhdGEtZmllbGQ9XCJsYWJlbFwiPjwvcGYtdjYtbGFiZWw+XG4gICAgICAgIDx0aW1lIGNsYXNzPVwiZXZlbnQtdGltZVwiIGRhdGEtZmllbGQ9XCJ0aW1lXCI+PC90aW1lPlxuICAgICAgICA8c3BhbiBjbGFzcz1cImV2ZW50LWVsZW1lbnRcIiBkYXRhLWZpZWxkPVwiZWxlbWVudFwiPjwvc3Bhbj5cbiAgICAgIDwvYnV0dG9uPmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAncHJpbWFyeS10YWctbmFtZScgfSlcbiAgYWNjZXNzb3IgcHJpbWFyeVRhZ05hbWUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdkZW1vLXRpdGxlJyB9KVxuICBhY2Nlc3NvciBkZW1vVGl0bGUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdwYWNrYWdlLW5hbWUnIH0pXG4gIGFjY2Vzc29yIHBhY2thZ2VOYW1lID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnY2Fub25pY2FsLXVybCcgfSlcbiAgYWNjZXNzb3IgY2Fub25pY2FsVVJMID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnc291cmNlLXVybCcgfSlcbiAgYWNjZXNzb3Igc291cmNlVVJMID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3Iga25vYnMgPSAnJztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBkcmF3ZXI6ICdleHBhbmRlZCcgfCAnY29sbGFwc2VkJyB8ICcnID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnZHJhd2VyLWhlaWdodCcgfSlcbiAgYWNjZXNzb3IgZHJhd2VySGVpZ2h0ID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAndGFicy1zZWxlY3RlZCcgfSlcbiAgYWNjZXNzb3IgdGFic1NlbGVjdGVkID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3Igc2lkZWJhcjogJ2V4cGFuZGVkJyB8ICdjb2xsYXBzZWQnIHwgJycgPSAnJztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCBhdHRyaWJ1dGU6ICdoYXMtZGVzY3JpcHRpb24nIH0pXG4gIGFjY2Vzc29yIGhhc0Rlc2NyaXB0aW9uID0gZmFsc2U7XG5cbiAgIyQoaWQ6IHN0cmluZykge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKGlkKTtcbiAgfVxuXG4gICMkJChzZWxlY3Rvcjogc3RyaW5nKSB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvckFsbChzZWxlY3RvcikgPz8gW107XG4gIH1cblxuICAjbG9nQ29udGFpbmVyOiBIVE1MRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjZHJhd2VyT3BlbiA9IGZhbHNlO1xuICAjaW5pdGlhbExvZ3NGZXRjaGVkID0gZmFsc2U7XG4gICNpc0luaXRpYWxMb2FkID0gdHJ1ZTtcbiAgI2RlYnVnRGF0YTogRGVidWdEYXRhIHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gRWxlbWVudCBldmVudCB0cmFja2luZ1xuICAjZWxlbWVudEV2ZW50TWFwOiBNYXA8c3RyaW5nLCBFdmVudEluZm8+IHwgbnVsbCA9IG51bGw7XG4gICNtYW5pZmVzdDogTWFuaWZlc3QgfCBudWxsID0gbnVsbDtcbiAgI2NhcHR1cmVkRXZlbnRzOiBFdmVudFJlY29yZFtdID0gW107XG4gICNtYXhDYXB0dXJlZEV2ZW50cyA9IDEwMDA7XG4gICNldmVudExpc3Q6IEhUTUxFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNldmVudERldGFpbEhlYWRlcjogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50RGV0YWlsQm9keTogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI3NlbGVjdGVkRXZlbnRJZDogc3RyaW5nIHwgbnVsbCA9IG51bGw7XG4gICNldmVudHNGaWx0ZXJWYWx1ZSA9ICcnO1xuICAjZXZlbnRzRmlsdGVyRGVib3VuY2VUaW1lcjogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50VHlwZUZpbHRlcnMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgI2VsZW1lbnRGaWx0ZXJzID0gbmV3IFNldDxzdHJpbmc+KCk7XG4gICNkaXNjb3ZlcmVkRWxlbWVudHMgPSBuZXcgU2V0PHN0cmluZz4oKTtcblxuICAvLyBFdmVudCBsaXN0ZW5lciByZWZlcmVuY2VzIGZvciBjbGVhbnVwXG4gICNoYW5kbGVMb2dzRXZlbnQ6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZUV4cGFuZDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG4gICNoYW5kbGVUcmVlQ29sbGFwc2U6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZVNlbGVjdDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gVGltZW91dCBJRHMgZm9yIGNsZWFudXBcbiAgI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjY29weURlYnVnRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dDogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcblxuICAvLyBMb2cgZmlsdGVyIHN0YXRlXG4gICNsb2dzRmlsdGVyVmFsdWUgPSAnJztcbiAgI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyOiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjbG9nTGV2ZWxGaWx0ZXJzID0gbmV3IFNldChbJ2luZm8nLCAnd2FybicsICdlcnJvcicsICdkZWJ1ZyddKTtcbiAgI2xvZ0xldmVsRHJvcGRvd246IEVsZW1lbnQgfCBudWxsID0gbnVsbDtcblxuICAvLyBXYXRjaCBmb3IgZHluYW1pY2FsbHkgYWRkZWQgZWxlbWVudHNcbiAgLyogYzggaWdub3JlIHN0YXJ0IC0gTXV0YXRpb25PYnNlcnZlciBjYWxsYmFjayB0ZXN0ZWQgdmlhIGludGVncmF0aW9uICovXG4gICNvYnNlcnZlciA9IG5ldyBNdXRhdGlvbk9ic2VydmVyKChtdXRhdGlvbnMpID0+IHtcbiAgICBsZXQgbmVlZHNVcGRhdGUgPSBmYWxzZTtcblxuICAgIGZvciAoY29uc3QgbXV0YXRpb24gb2YgbXV0YXRpb25zKSB7XG4gICAgICBmb3IgKGNvbnN0IG5vZGUgb2YgbXV0YXRpb24uYWRkZWROb2Rlcykge1xuICAgICAgICBpZiAobm9kZSBpbnN0YW5jZW9mIEhUTUxFbGVtZW50KSB7XG4gICAgICAgICAgY29uc3QgdGFnTmFtZSA9IG5vZGUudGFnTmFtZS50b0xvd2VyQ2FzZSgpO1xuICAgICAgICAgIGlmICh0aGlzLiNlbGVtZW50RXZlbnRNYXA/Lmhhcyh0YWdOYW1lKSAmJiAhbm9kZS5kYXRhc2V0LmNlbUV2ZW50c0F0dGFjaGVkKSB7XG4gICAgICAgICAgICBjb25zdCBldmVudEluZm8gPSB0aGlzLiNlbGVtZW50RXZlbnRNYXAuZ2V0KHRhZ05hbWUpITtcbiAgICAgICAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGV2ZW50SW5mby5ldmVudE5hbWVzKSB7XG4gICAgICAgICAgICAgIG5vZGUuYWRkRXZlbnRMaXN0ZW5lcihldmVudE5hbWUsIHRoaXMuI2hhbmRsZUVsZW1lbnRFdmVudCwgeyBjYXB0dXJlOiB0cnVlIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgbm9kZS5kYXRhc2V0LmNlbUV2ZW50c0F0dGFjaGVkID0gJ3RydWUnO1xuICAgICAgICAgICAgdGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICAgICAgICAgIG5lZWRzVXBkYXRlID0gdHJ1ZTtcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICBpZiAobmVlZHNVcGRhdGUpIHtcbiAgICAgIHRoaXMuI3VwZGF0ZUV2ZW50RmlsdGVycygpO1xuICAgIH1cbiAgfSk7XG4gIC8qIGM4IGlnbm9yZSBzdG9wICovXG5cbiAgI3dzQ2xpZW50OiBhbnk7XG4gICNjbGllbnRNb2R1bGVzTG9hZGVkID0gZmFsc2U7XG5cbiAgI2luaXRXc0NsaWVudCgpIHtcbiAgICB0aGlzLiN3c0NsaWVudCA9IG5ldyBDRU1SZWxvYWRDbGllbnQoe1xuICAgIGppdHRlck1heDogMTAwMCxcbiAgICBvdmVybGF5VGhyZXNob2xkOiAxNSxcbiAgICBiYWRnZUZhZGVEZWxheTogMjAwMCxcbiAgICAvKiBjOCBpZ25vcmUgc3RhcnQgLSBXZWJTb2NrZXQgY2FsbGJhY2tzIHRlc3RlZCB2aWEgaW50ZWdyYXRpb24gKi9cbiAgICBjYWxsYmFja3M6IHtcbiAgICAgIG9uT3BlbjogKCkgPT4ge1xuICAgICAgICB0aGlzLiMkKCdyZWNvbm5lY3Rpb24tbW9kYWwnKT8uY2xvc2UoKTtcbiAgICAgIH0sXG4gICAgICBvbkVycm9yOiAoZXJyb3JEYXRhOiB7IHRpdGxlPzogc3RyaW5nOyBtZXNzYWdlPzogc3RyaW5nOyBmaWxlPzogc3RyaW5nIH0pID0+IHtcbiAgICAgICAgaWYgKGVycm9yRGF0YT8udGl0bGUgJiYgZXJyb3JEYXRhPy5tZXNzYWdlKSB7XG4gICAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZV0gU2VydmVyIGVycm9yOicsIGVycm9yRGF0YSk7XG4gICAgICAgICAgKHRoaXMuIyQoJ2Vycm9yLW92ZXJsYXknKSBhcyBhbnkpPy5zaG93KGVycm9yRGF0YS50aXRsZSwgZXJyb3JEYXRhLm1lc3NhZ2UsIGVycm9yRGF0YS5maWxlKTtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlXSBXZWJTb2NrZXQgZXJyb3I6JywgZXJyb3JEYXRhKTtcbiAgICAgICAgfVxuICAgICAgfSxcbiAgICAgIG9uUmVjb25uZWN0aW5nOiAoeyBhdHRlbXB0LCBkZWxheSB9OiB7IGF0dGVtcHQ6IG51bWJlcjsgZGVsYXk6IG51bWJlciB9KSA9PiB7XG4gICAgICAgIGlmIChhdHRlbXB0ID49IDE1KSB7XG4gICAgICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LnNob3dNb2RhbCgpO1xuICAgICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tY29udGVudCcpIGFzIGFueSk/LnVwZGF0ZVJldHJ5SW5mbyhhdHRlbXB0LCBkZWxheSk7XG4gICAgICAgIH1cbiAgICAgIH0sXG4gICAgICBvblJlbG9hZDogKCkgPT4ge1xuICAgICAgICBjb25zdCBlcnJvck92ZXJsYXkgPSB0aGlzLiMkKCdlcnJvci1vdmVybGF5Jyk7XG4gICAgICAgIGlmIChlcnJvck92ZXJsYXk/Lmhhc0F0dHJpYnV0ZSgnb3BlbicpKSB7XG4gICAgICAgICAgKGVycm9yT3ZlcmxheSBhcyBhbnkpLmhpZGUoKTtcbiAgICAgICAgfVxuICAgICAgICB3aW5kb3cubG9jYXRpb24ucmVsb2FkKCk7XG4gICAgICB9LFxuICAgICAgb25TaHV0ZG93bjogKCkgPT4ge1xuICAgICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLW1vZGFsJykgYXMgYW55KT8uc2hvd01vZGFsKCk7XG4gICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tY29udGVudCcpIGFzIGFueSk/LnVwZGF0ZVJldHJ5SW5mbygzMCwgMzAwMDApO1xuICAgICAgfSxcbiAgICAgIG9uTG9nczogKGxvZ3M6IEFycmF5PHsgdHlwZTogc3RyaW5nOyBkYXRlOiBzdHJpbmc7IG1lc3NhZ2U6IHN0cmluZyB9PikgPT4ge1xuICAgICAgICB3aW5kb3cuZGlzcGF0Y2hFdmVudChuZXcgQ2VtTG9nc0V2ZW50KGxvZ3MpKTtcbiAgICAgIH1cbiAgICB9XG4gICAgLyogYzggaWdub3JlIHN0b3AgKi9cbiAgICB9KTtcbiAgfVxuXG4gIGdldCBkZW1vKCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5xdWVyeVNlbGVjdG9yKCdjZW0tc2VydmUtZGVtbycpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGxpbmsgcmVsPVwic3R5bGVzaGVldFwiIGhyZWY9XCIvX19jZW0vcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0LmNzc1wiPlxuICAgICAgPGxpbmsgcmVsPVwic3R5bGVzaGVldFwiIGhyZWY9XCIvX19jZW0vcGYtbGlnaHRkb20uY3NzXCI+XG5cbiAgICAgIDxwZi12Ni1wYWdlID9zaWRlYmFyLWNvbGxhcHNlZD0ke3RoaXMuc2lkZWJhciA9PT0gJ2NvbGxhcHNlZCd9PlxuICAgICAgICA8cGYtdjYtc2tpcC10by1jb250ZW50IHNsb3Q9XCJza2lwLXRvLWNvbnRlbnRcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhyZWY9XCIjbWFpbi1jb250ZW50XCI+XG4gICAgICAgICAgU2tpcCB0byBjb250ZW50XG4gICAgICAgIDwvcGYtdjYtc2tpcC10by1jb250ZW50PlxuXG4gICAgICAgIDxwZi12Ni1tYXN0aGVhZCBzbG90PVwibWFzdGhlYWRcIj5cbiAgICAgICAgICA8YSBjbGFzcz1cIm1hc3RoZWFkLWxvZ29cIlxuICAgICAgICAgICAgIGhyZWY9XCIvXCJcbiAgICAgICAgICAgICBzbG90PVwibG9nb1wiPlxuICAgICAgICAgICAgPGltZyBhbHQ9XCJDRU0gRGV2IFNlcnZlclwiXG4gICAgICAgICAgICAgICAgIHNyYz1cIi9fX2NlbS9sb2dvLnN2Z1wiPlxuICAgICAgICAgICAgJHt0aGlzLnBhY2thZ2VOYW1lID8gaHRtbGA8aDE+JHt0aGlzLnBhY2thZ2VOYW1lfTwvaDE+YCA6IG5vdGhpbmd9XG4gICAgICAgICAgPC9hPlxuICAgICAgICAgIDxwZi12Ni10b29sYmFyIHNsb3Q9XCJ0b29sYmFyXCI+XG4gICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cCB2YXJpYW50PVwiYWN0aW9uLWdyb3VwXCI+XG4gICAgICAgICAgICAgICR7dGhpcy4jcmVuZGVyU291cmNlQnV0dG9uKCl9XG4gICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImRlYnVnLWluZm9cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInBsYWluXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJEZWJ1ZyBpbmZvXCI+XG4gICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxNlwiXG4gICAgICAgICAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTggMTVBNyA3IDAgMSAxIDggMWE3IDcgMCAwIDEgMCAxNHptMCAxQTggOCAwIDEgMCA4IDBhOCA4IDAgMCAwIDAgMTZ6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwibTguOTMgNi41ODgtMi4yOS4yODctLjA4Mi4zOC40NS4wODNjLjI5NC4wNy4zNTIuMTc2LjI4OC40NjlsLS43MzggMy40NjhjLS4xOTQuODk3LjEwNSAxLjMxOS44MDggMS4zMTkuNTQ1IDAgMS4xNzgtLjI1MiAxLjQ2NS0uNTk4bC4wODgtLjQxNmMtLjIuMTc2LS40OTIuMjQ2LS42ODYuMjQ2LS4yNzUgMC0uMzc1LS4xOTMtLjMwNC0uNTMzTDguOTMgNi41ODh6TTkgNC41YTEgMSAwIDEgMS0yIDAgMSAxIDAgMCAxIDIgMHpcIi8+XG4gICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgPHBmLXY2LXRvZ2dsZS1ncm91cCBjbGFzcz1cImNvbG9yLXNjaGVtZS10b2dnbGVcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkNvbG9yIHNjaGVtZVwiPlxuICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtIHZhbHVlPVwibGlnaHRcIj5cbiAgICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiIGFyaWEtbGFiZWw9XCJMaWdodCBtb2RlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk04IDExYTMgMyAwIDEgMSAwLTYgMyAzIDAgMCAxIDAgNnptMCAxYTQgNCAwIDEgMCAwLTggNCA0IDAgMCAwIDAgOHpNOCAwYS41LjUgMCAwIDEgLjUuNXYyYS41LjUgMCAwIDEtMSAwdi0yQS41LjUgMCAwIDEgOCAwem0wIDEzYS41LjUgMCAwIDEgLjUuNXYyYS41LjUgMCAwIDEtMSAwdi0yQS41LjUgMCAwIDEgOCAxM3ptOC01YS41LjUgMCAwIDEtLjUuNWgtMmEuNS41IDAgMCAxIDAtMWgyYS41LjUgMCAwIDEgLjUuNXpNMyA4YS41LjUgMCAwIDEtLjUuNWgtMmEuNS41IDAgMCAxIDAtMWgyQS41LjUgMCAwIDEgMyA4em0xMC42NTctNS42NTdhLjUuNSAwIDAgMSAwIC43MDdsLTEuNDE0IDEuNDE1YS41LjUgMCAxIDEtLjcwNy0uNzA4bDEuNDE0LTEuNDE0YS41LjUgMCAwIDEgLjcwNyAwem0tOS4xOTMgOS4xOTNhLjUuNSAwIDAgMSAwIC43MDdMMy4wNSAxMy42NTdhLjUuNSAwIDAgMS0uNzA3LS43MDdsMS40MTQtMS40MTRhLjUuNSAwIDAgMSAuNzA3IDB6bTkuMTkzIDIuMTIxYS41LjUgMCAwIDEtLjcwNyAwbC0xLjQxNC0xLjQxNGEuNS41IDAgMCAxIC43MDctLjcwN2wxLjQxNCAxLjQxNGEuNS41IDAgMCAxIDAgLjcwN3pNNC40NjQgNC40NjVhLjUuNSAwIDAgMS0uNzA3IDBMMi4zNDMgMy4wNWEuNS41IDAgMSAxIC43MDctLjcwN2wxLjQxNCAxLjQxNGEuNS41IDAgMCAxIDAgLjcwOHpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICAgIDxwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cImRhcmtcIj5cbiAgICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiIGFyaWEtbGFiZWw9XCJEYXJrIG1vZGVcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTYgLjI3OGEuNzY4Ljc2OCAwIDAgMSAuMDguODU4IDcuMjA4IDcuMjA4IDAgMCAwLS44NzggMy40NmMwIDQuMDIxIDMuMjc4IDcuMjc3IDcuMzE4IDcuMjc3LjUyNyAwIDEuMDQtLjA1NSAxLjUzMy0uMTZhLjc4Ny43ODcgMCAwIDEgLjgxLjMxNi43MzMuNzMzIDAgMCAxLS4wMzEuODkzQTguMzQ5IDguMzQ5IDAgMCAxIDguMzQ0IDE2QzMuNzM0IDE2IDAgMTIuMjg2IDAgNy43MSAwIDQuMjY2IDIuMTE0IDEuMzEyIDUuMTI0LjA2QS43NTIuNzUyIDAgMCAxIDYgLjI3OHpNNC44NTggMS4zMTFBNy4yNjkgNy4yNjkgMCAwIDAgMS4wMjUgNy43MWMwIDQuMDIgMy4yNzkgNy4yNzYgNy4zMTkgNy4yNzZhNy4zMTYgNy4zMTYgMCAwIDAgNS4yMDUtMi4xNjJjLS4zMzcuMDQyLS42OC4wNjMtMS4wMjkuMDYzLTQuNjEgMC04LjM0My0zLjcxNC04LjM0My04LjI5IDAtMS4xNjcuMjQyLTIuMjc4LjY4MS0zLjI4NnpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICAgIDxwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cInN5c3RlbVwiPlxuICAgICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCIgYXJpYS1sYWJlbD1cIlN5c3RlbSBwcmVmZXJlbmNlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0wIDEuNUExLjUgMS41IDAgMCAxIDEuNSAwaDEzQTEuNSAxLjUgMCAwIDEgMTYgMS41djhhMS41IDEuNSAwIDAgMS0xLjUgMS41aC0xM0ExLjUgMS41IDAgMCAxIDAgOS41di04ek0xLjUgMWEuNS41IDAgMCAwLS41LjV2OGEuNS41IDAgMCAwIC41LjVoMTNhLjUuNSAwIDAgMCAuNS0uNXYtOGEuNS41IDAgMCAwLS41LS41aC0xM3pcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0yLjUgMTJoMTFhLjUuNSAwIDAgMSAwIDFoLTExYS41LjUgMCAwIDEgMC0xem0wIDJoMTFhLjUuNSAwIDAgMSAwIDFoLTExYS41LjUgMCAwIDEgMC0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtPlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9nZ2xlLWdyb3VwPlxuICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICA8L3BmLXY2LXRvb2xiYXI+XG4gICAgICAgIDwvcGYtdjYtbWFzdGhlYWQ+XG5cbiAgICAgICAgPHBmLXY2LXBhZ2Utc2lkZWJhciBzbG90PVwic2lkZWJhclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgP2V4cGFuZGVkPSR7dGhpcy5zaWRlYmFyID09PSAnZXhwYW5kZWQnfVxuICAgICAgICAgICAgICAgICAgICAgICAgICAgID9jb2xsYXBzZWQ9JHt0aGlzLnNpZGViYXIgIT09ICdleHBhbmRlZCd9PlxuICAgICAgICAgIDxzbG90IG5hbWU9XCJuYXZpZ2F0aW9uXCI+PC9zbG90PlxuICAgICAgICA8L3BmLXY2LXBhZ2Utc2lkZWJhcj5cblxuICAgICAgICA8cGYtdjYtcGFnZS1tYWluIHNsb3Q9XCJtYWluXCIgaWQ9XCJtYWluLWNvbnRlbnRcIj5cbiAgICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICAgICAgPGZvb3RlciBjbGFzcz1cInBmLW0tc3RpY2t5LWJvdHRvbVwiPlxuICAgICAgICAgICAgPGRpdiBjbGFzcz1cImZvb3Rlci1kZXNjcmlwdGlvbiR7dGhpcy5oYXNEZXNjcmlwdGlvbiA/ICcnIDogJyBlbXB0eSd9XCI+XG4gICAgICAgICAgICAgIDxzbG90IG5hbWU9XCJkZXNjcmlwdGlvblwiPjwvc2xvdD5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgPGNlbS1kcmF3ZXIgP29wZW49JHt0aGlzLmRyYXdlciA9PT0gJ2V4cGFuZGVkJ31cbiAgICAgICAgICAgICAgICAgICAgICAgIGRyYXdlci1oZWlnaHQ9XCIke3RoaXMuZHJhd2VySGVpZ2h0IHx8ICc0MDAnfVwiPlxuICAgICAgICAgICAgICA8cGYtdjYtdGFicyBzZWxlY3RlZD1cIiR7dGhpcy50YWJzU2VsZWN0ZWQgfHwgJzAnfVwiPlxuICAgICAgICAgICAgICAgIDxwZi12Ni10YWIgdGl0bGU9XCJLbm9ic1wiPlxuICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImtub2JzLWNvbnRhaW5lclwiPlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwia25vYnNcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cCBjbGFzcz1cImtub2JzLWVtcHR5XCI+Tm8ga25vYnMgYXZhaWxhYmxlIGZvciB0aGlzIGVsZW1lbnQuPC9wPlxuICAgICAgICAgICAgICAgICAgICA8L3Nsb3Q+XG4gICAgICAgICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRhYj5cbiAgICAgICAgICAgICAgICA8cGYtdjYtdGFiIHRpdGxlPVwiTWFuaWZlc3QgQnJvd3NlclwiPlxuICAgICAgICAgICAgICAgICAgPGNlbS1tYW5pZmVzdC1icm93c2VyPlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtdHJlZVwiIHNsb3Q9XCJtYW5pZmVzdC10cmVlXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtbmFtZVwiIHNsb3Q9XCJtYW5pZmVzdC1uYW1lXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtZGV0YWlsc1wiIHNsb3Q9XCJtYW5pZmVzdC1kZXRhaWxzXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgPC9jZW0tbWFuaWZlc3QtYnJvd3Nlcj5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRhYj5cbiAgICAgICAgICAgICAgICA8cGYtdjYtdGFiIHRpdGxlPVwiU2VydmVyIExvZ3NcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJsb2dzLXdyYXBwZXJcIj5cbiAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXIgc3RpY2t5PlxuICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRleHQtaW5wdXQtZ3JvdXAgaWQ9XCJsb2dzLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHBsYWNlaG9sZGVyPVwiRmlsdGVyIGxvZ3MuLi5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpY29uPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDUxMiA1MTJcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNTA1IDQ0Mi43TDQwNS4zIDM0M2MtNC41LTQuNS0xMC42LTctMTctN0gzNzJjMjcuNi0zNS4zIDQ0LTc5LjcgNDQtMTI4QzQxNiA5My4xIDMyMi45IDAgMjA4IDBTMCA5My4xIDAgMjA4czkzLjEgMjA4IDIwOCAyMDhjNDguMyAwIDkyLjctMTYuNCAxMjgtNDR2MTYuM2MwIDYuNCAyLjUgMTIuNSA3IDE3bDk5LjcgOTkuN2M5LjQgOS40IDI0LjYgOS40IDMzLjkgMGwyOC4zLTI4LjNjOS40LTkuNCA5LjQtMjQuNi4xLTM0ek0yMDggMzM2Yy03MC43IDAtMTI4LTU3LjItMTI4LTEyOCAwLTcwLjcgNTcuMi0xMjggMTI4LTEyOCA3MC43IDAgMTI4IDU3LjIgMTI4IDEyOCAwIDcwLjctNTcuMiAxMjgtMTI4IDEyOHpcIj48L3BhdGg+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdGV4dC1pbnB1dC1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwibG9nLWxldmVsLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBsYWJlbD1cIkZpbHRlciBsb2cgbGV2ZWxzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHNwYW4gc2xvdD1cInRvZ2dsZS10ZXh0XCI+TG9nIExldmVsczwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwiaW5mb1wiIGNoZWNrZWQ+SW5mbzwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJ3YXJuXCIgY2hlY2tlZD5XYXJuaW5nczwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJlcnJvclwiIGNoZWNrZWQ+RXJyb3JzPC9wZi12Ni1tZW51LWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LW1lbnUtaXRlbSB2YXJpYW50PVwiY2hlY2tib3hcIiB2YWx1ZT1cImRlYnVnXCIgY2hlY2tlZD5EZWJ1ZzwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWRyb3Bkb3duPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWdyb3VwIHZhcmlhbnQ9XCJhY3Rpb24tZ3JvdXBcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJjb3B5LWxvZ3NcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJ0ZXJ0aWFyeVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgc2l6ZT1cInNtYWxsXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHN2ZyBzbG90PVwiaWNvblwiIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTEzIDBINmEyIDIgMCAwIDAtMiAyIDIgMiAwIDAgMC0yIDJ2MTBhMiAyIDAgMCAwIDIgMmg3YTIgMiAwIDAgMCAyLTIgMiAyIDAgMCAwIDItMlYyYTIgMiAwIDAgMC0yLTJ6bTAgMTNWNGEyIDIgMCAwIDAtMi0ySDVhMSAxIDAgMCAxIDEtMWg3YTEgMSAwIDAgMSAxIDF2MTBhMSAxIDAgMCAxLTEgMXpNMyAxM1Y0YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjlhMSAxIDAgMCAxLTEgMUg0YTEgMSAwIDAgMS0xLTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIENvcHkgTG9nc1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyPlxuICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwibG9nLWNvbnRhaW5lclwiPjwvZGl2PlxuICAgICAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICAgICAgPC9wZi12Ni10YWI+XG4gICAgICAgICAgICAgICAgPHBmLXY2LXRhYiB0aXRsZT1cIkV2ZW50c1wiPlxuICAgICAgICAgICAgICAgICAgPGRpdiBjbGFzcz1cImV2ZW50cy13cmFwcGVyXCI+XG4gICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyIHN0aWNreT5cbiAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10ZXh0LWlucHV0LWdyb3VwIGlkPVwiZXZlbnRzLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHBsYWNlaG9sZGVyPVwiRmlsdGVyIGV2ZW50cy4uLlwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGljb24+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHN2ZyBzbG90PVwiaWNvblwiIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTExLjc0MiAxMC4zNDRhNi41IDYuNSAwIDEgMC0xLjM5NyAxLjM5OGgtLjAwMWMuMDMuMDQuMDYyLjA3OC4wOTguMTE1bDMuODUgMy44NWExIDEgMCAwIDAgMS40MTUtMS40MTRsLTMuODUtMy44NWExLjAwNyAxLjAwNyAwIDAgMC0uMTE1LS4xek0xMiA2LjVhNS41IDUuNSAwIDEgMS0xMSAwIDUuNSA1LjUgMCAwIDEgMTEgMHpcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdGV4dC1pbnB1dC1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwiZXZlbnQtdHlwZS1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgZXZlbnQgdHlwZXNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5FdmVudCBUeXBlczwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1kcm9wZG93bj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwiZWxlbWVudC1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgZWxlbWVudHNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5FbGVtZW50czwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1kcm9wZG93bj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cCB2YXJpYW50PVwiYWN0aW9uLWdyb3VwXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtYnV0dG9uIGlkPVwiY2xlYXItZXZlbnRzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk01LjUgNS41QS41LjUgMCAwIDEgNiA2djZhLjUuNSAwIDAgMS0xIDBWNmEuNS41IDAgMCAxIC41LS41em0yLjUgMGEuNS41IDAgMCAxIC41LjV2NmEuNS41IDAgMCAxLTEgMFY2YS41LjUgMCAwIDEgLjUtLjV6bTMgLjVhLjUuNSAwIDAgMC0xIDB2NmEuNS41IDAgMCAwIDEgMFY2elwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGZpbGwtcnVsZT1cImV2ZW5vZGRcIiBkPVwiTTE0LjUgM2ExIDEgMCAwIDEtMSAxSDEzdjlhMiAyIDAgMCAxLTIgMkg1YTIgMiAwIDAgMS0yLTJWNGgtLjVhMSAxIDAgMCAxLTEtMVYyYTEgMSAwIDAgMSAxLTFINmExIDEgMCAwIDEgMS0xaDJhMSAxIDAgMCAxIDEgMWgzLjVhMSAxIDAgMCAxIDEgMXYxek00LjExOCA0IDQgNC4wNTlWMTNhMSAxIDAgMCAwIDEgMWg2YTEgMSAwIDAgMCAxLTFWNC4wNTlMMTEuODgyIDRINC4xMTh6TTIuNSAzVjJoMTF2MWgtMTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIENsZWFyIEV2ZW50c1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImNvcHktZXZlbnRzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0xMyAwSDZhMiAyIDAgMCAwLTIgMiAyIDIgMCAwIDAtMiAydjEwYTIgMiAwIDAgMCAyIDJoN2EyIDIgMCAwIDAgMi0yIDIgMiAwIDAgMCAyLTJWMmEyIDIgMCAwIDAtMi0yem0wIDEzVjRhMiAyIDAgMCAwLTItMkg1YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjEwYTEgMSAwIDAgMS0xIDF6TTMgMTNWNGExIDEgMCAwIDEgMS0xaDdhMSAxIDAgMCAxIDEgMXY5YTEgMSAwIDAgMS0xIDFINGExIDEgMCAwIDEtMS0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDb3B5IEV2ZW50c1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyPlxuICAgICAgICAgICAgICAgICAgICA8cGYtdjYtZHJhd2VyIGlkPVwiZXZlbnQtZHJhd2VyXCIgZXhwYW5kZWQ+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWxpc3RcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwiZXZlbnQtZGV0YWlsLWhlYWRlclwiIHNsb3Q9XCJwYW5lbC1oZWFkZXJcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwiZXZlbnQtZGV0YWlsLWJvZHlcIiBzbG90PVwicGFuZWwtYm9keVwiPjwvZGl2PlxuICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWRyYXdlcj5cbiAgICAgICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtdGFiPlxuICAgICAgICAgICAgICAgIDxwZi12Ni10YWIgdGl0bGU9XCJIZWFsdGhcIj5cbiAgICAgICAgICAgICAgICAgIDxjZW0taGVhbHRoLXBhbmVsIGNvbXBvbmVudD0ke2lmRGVmaW5lZCh0aGlzLnByaW1hcnlUYWdOYW1lKX0+XG4gICAgICAgICAgICAgICAgICA8L2NlbS1oZWFsdGgtcGFuZWw+XG4gICAgICAgICAgICAgICAgPC9wZi12Ni10YWI+XG4gICAgICAgICAgICAgIDwvcGYtdjYtdGFicz5cbiAgICAgICAgICAgIDwvY2VtLWRyYXdlcj5cbiAgICAgICAgICA8L2Zvb3Rlcj5cbiAgICAgICAgPC9wZi12Ni1wYWdlLW1haW4+XG4gICAgICA8L3BmLXY2LXBhZ2U+XG5cbiAgICAgIDxwZi12Ni1tb2RhbCBpZD1cImRlYnVnLW1vZGFsXCIgdmFyaWFudD1cImxhcmdlXCI+XG4gICAgICAgIDxoMiBzbG90PVwiaGVhZGVyXCI+RGVidWcgSW5mb3JtYXRpb248L2gyPlxuICAgICAgICA8ZGwgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3QgcGYtbS1ob3Jpem9udGFsIHBmLW0tY29tcGFjdFwiPlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5TZXJ2ZXIgVmVyc2lvbjwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy12ZXJzaW9uXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlNlcnZlciBPUzwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy1vc1wiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5XYXRjaCBEaXJlY3Rvcnk8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctd2F0Y2gtZGlyXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPk1hbmlmZXN0IFNpemU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctbWFuaWZlc3Qtc2l6ZVwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5EZW1vcyBGb3VuZDwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy1kZW1vLWNvdW50XCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlRhZyBOYW1lPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIj4ke3RoaXMucHJpbWFyeVRhZ05hbWUgfHwgJy0nfTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkRlbW8gVGl0bGU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiPiR7dGhpcy5kZW1vVGl0bGUgfHwgJy0nfTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkJyb3dzZXI8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctYnJvd3NlclwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5Vc2VyIEFnZW50PC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLXVhXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgIDwvZGw+XG4gICAgICAgIDxkaXYgaWQ9XCJkZW1vLXVybHMtY29udGFpbmVyXCI+PC9kaXY+XG4gICAgICAgIDxwZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24gaWQ9XCJkZWJ1Zy1pbXBvcnRtYXAtZGV0YWlsc1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdG9nZ2xlLXRleHQ9XCJTaG93IEltcG9ydCBNYXBcIj5cbiAgICAgICAgICA8cHJlIGlkPVwiZGVidWctaW1wb3J0bWFwXCI+LTwvcHJlPlxuICAgICAgICA8L3BmLXY2LWV4cGFuZGFibGUtc2VjdGlvbj5cbiAgICAgICAgPGRpdiBzbG90PVwiZm9vdGVyXCIgY2xhc3M9XCJidXR0b24tY29udGFpbmVyXCI+XG4gICAgICAgICAgPHBmLXY2LWJ1dHRvbiBjbGFzcz1cImRlYnVnLWNvcHlcIiB2YXJpYW50PVwicHJpbWFyeVwiPlxuICAgICAgICAgICAgQ29weSBEZWJ1ZyBJbmZvXG4gICAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICAgICAgPHBmLXY2LWJ1dHRvbiBjbGFzcz1cImRlYnVnLWNsb3NlXCIgdmFyaWFudD1cInNlY29uZGFyeVwiPlxuICAgICAgICAgICAgQ2xvc2VcbiAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgPC9kaXY+XG4gICAgICA8L3BmLXY2LW1vZGFsPlxuXG4gICAgICA8IS0tIFJlY29ubmVjdGlvbiBtb2RhbCAtLT5cbiAgICAgIDxwZi12Ni1tb2RhbCBpZD1cInJlY29ubmVjdGlvbi1tb2RhbFwiIHZhcmlhbnQ9XCJsYXJnZVwiPlxuICAgICAgICA8aDIgc2xvdD1cImhlYWRlclwiPkRldmVsb3BtZW50IFNlcnZlciBEaXNjb25uZWN0ZWQ8L2gyPlxuICAgICAgICA8Y2VtLXJlY29ubmVjdGlvbi1jb250ZW50IGlkPVwicmVjb25uZWN0aW9uLWNvbnRlbnRcIj48L2NlbS1yZWNvbm5lY3Rpb24tY29udGVudD5cbiAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cInJlbG9hZC1idXR0b25cIlxuICAgICAgICAgICAgICAgICAgICAgIHNsb3Q9XCJmb290ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwcmltYXJ5XCI+UmVsb2FkIFBhZ2U8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cInJldHJ5LWJ1dHRvblwiXG4gICAgICAgICAgICAgICAgICAgICAgc2xvdD1cImZvb3RlclwiXG4gICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInNlY29uZGFyeVwiPlJldHJ5IE5vdzwvcGYtdjYtYnV0dG9uPlxuICAgICAgPC9wZi12Ni1tb2RhbD5cblxuICAgICAgPCEtLSBUcmFuc2Zvcm0gZXJyb3Igb3ZlcmxheSAtLT5cbiAgICAgIDxjZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXkgaWQ9XCJlcnJvci1vdmVybGF5XCI+XG4gICAgICA8L2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheT5cbiAgICBgO1xuICB9XG5cbiAgI3JlbmRlclNvdXJjZUJ1dHRvbigpIHtcbiAgICBpZiAoIXRoaXMuc291cmNlVVJMKSByZXR1cm4gbm90aGluZztcblxuICAgIGxldCBsYWJlbCA9ICdWZXJzaW9uIENvbnRyb2wnO1xuICAgIGxldCBwYXRoID0gJ001Ljg1NCA0Ljg1NGEuNS41IDAgMSAwLS43MDgtLjcwOGwtMy41IDMuNWEuNS41IDAgMCAwIDAgLjcwOGwzLjUgMy41YS41LjUgMCAwIDAgLjcwOC0uNzA4TDIuNzA3IDhsMy4xNDctMy4xNDZ6bTQuMjkyIDBhLjUuNSAwIDAgMSAuNzA4LS43MDhsMy41IDMuNWEuNS41IDAgMCAxIDAgLjcwOGwtMy41IDMuNWEuNS41IDAgMCAxLS43MDgtLjcwOEwxMy4yOTMgOGwtMy4xNDctMy4xNDZ6JztcblxuICAgIGlmICh0aGlzLnNvdXJjZVVSTC5pbmNsdWRlcygnZ2l0aHViLmNvbScpKSB7XG4gICAgICBsYWJlbCA9ICdHaXRIdWIuY29tJztcbiAgICAgIHBhdGggPSAnTTggMEMzLjU4IDAgMCAzLjU4IDAgOGMwIDMuNTQgMi4yOSA2LjUzIDUuNDcgNy41OS40LjA3LjU1LS4xNy41NS0uMzggMC0uMTktLjAxLS44Mi0uMDEtMS40OS0yLjAxLjM3LTIuNTMtLjQ5LTIuNjktLjk0LS4wOS0uMjMtLjQ4LS45NC0uODItMS4xMy0uMjgtLjE1LS42OC0uNTItLjAxLS41My42My0uMDEgMS4wOC41OCAxLjIzLjgyLjcyIDEuMjEgMS44Ny44NyAyLjMzLjY2LjA3LS41Mi4yOC0uODcuNTEtMS4wNy0xLjc4LS4yLTMuNjQtLjg5LTMuNjQtMy45NSAwLS44Ny4zMS0xLjU5LjgyLTIuMTUtLjA4LS4yLS4zNi0xLjAyLjA4LTIuMTIgMCAwIC42Ny0uMjEgMi4yLjgyLjY0LS4xOCAxLjMyLS4yNyAyLS4yNy42OCAwIDEuMzYuMDkgMiAuMjcgMS41My0xLjA0IDIuMi0uODIgMi4yLS44Mi40NCAxLjEuMTYgMS45Mi4wOCAyLjEyLjUxLjU2LjgyIDEuMjcuODIgMi4xNSAwIDMuMDctMS44NyAzLjc1LTMuNjUgMy45NS4yOS4yNS41NC43My41NCAxLjQ4IDAgMS4wNy0uMDEgMS45My0uMDEgMi4yIDAgLjIxLjE1LjQ2LjU1LjM4QTguMDEzIDguMDEzIDAgMDAxNiA4YzAtNC40Mi0zLjU4LTgtOC04eic7XG4gICAgfSBlbHNlIGlmICh0aGlzLnNvdXJjZVVSTC5pbmNsdWRlcygnZ2l0bGFiLmNvbScpKSB7XG4gICAgICBsYWJlbCA9ICdHaXRMYWInO1xuICAgICAgcGF0aCA9ICdtMTUuNzM0IDYuMS0uMDIyLS4wNThMMTMuNTM0LjM1OGEuNTY4LjU2OCAwIDAgMC0uNTYzLS4zNTYuNTgzLjU4MyAwIDAgMC0uMzI4LjEyMi41ODIuNTgyIDAgMCAwLS4xOTMuMjk0bC0xLjQ3IDQuNDk5SDUuMDI1bC0xLjQ3LTQuNUEuNTcyLjU3MiAwIDAgMCAzLjM2MC4xNzRhLjU3Mi41NzIgMCAwIDAtLjMyOC0uMTcyLjU4Mi41ODIgMCAwIDAtLjU2My4zNTdMLjI5IDYuMDRsLS4wMjIuMDU3QTQuMDQ0IDQuMDQ0IDAgMCAwIDEuNjEgMTAuNzdsLjAwNy4wMDYuMDIuMDE0IDMuMzE4IDIuNDg1IDEuNjQgMS4yNDIgMSAuNzU1YS42NzMuNjczIDAgMCAwIC44MTQgMGwxLS43NTUgMS42NC0xLjI0MiAzLjMzOC0yLjUuMDA5LS4wMDdhNC4wNSA0LjA1IDAgMCAwIDEuMzQtNC42NjhaJztcbiAgICB9IGVsc2UgaWYgKHRoaXMuc291cmNlVVJMLmluY2x1ZGVzKCdiaXRidWNrZXQub3JnJykpIHtcbiAgICAgIGxhYmVsID0gJ0JpdGJ1Y2tldCc7XG4gICAgICBwYXRoID0gJ00wIDEuNUExLjUgMS41IDAgMCAxIDEuNSAwaDEzQTEuNSAxLjUgMCAwIDEgMTYgMS41djEzYTEuNSAxLjUgMCAwIDEtMS41IDEuNWgtMTNBMS41IDEuNSAwIDAgMSAwIDE0LjV2LTEzek0yLjUgM2EuNS41IDAgMCAwLS41LjV2OWEuNS41IDAgMCAwIC41LjVoMTFhLjUuNSAwIDAgMCAuNS0uNXYtOWEuNS41IDAgMCAwLS41LS41aC0xMXptNS4wMzggMS40MzVhLjUuNSAwIDAgMSAuOTI0IDBsMS40MiAzLjM3SDguNzhsLS4yNDMuNjA4LS4yNDMtLjYwOEg1LjA4MmwxLjQyLTMuMzd6TTggOS4xNDNsLS43NDMgMS44NTdINC43NDNMNi4wNzYgNy42MDggOCA5LjE0M3onO1xuICAgIH1cblxuICAgIHJldHVybiBodG1sYFxuICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgPHBmLXY2LWJ1dHRvbiBocmVmPVwiJHt0aGlzLnNvdXJjZVVSTH1cIlxuICAgICAgICAgICAgICAgICAgICAgIHRhcmdldD1cIl9ibGFua1wiXG4gICAgICAgICAgICAgICAgICAgICAgcmVsPVwibm9vcGVuZXIgbm9yZWZlcnJlclwiXG4gICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInBsYWluXCJcbiAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiVmlldyBzb3VyY2UgZmlsZVwiPlxuICAgICAgICAgIDxzdmcgYXJpYS1sYWJlbD1cIiR7bGFiZWx9XCJcbiAgICAgICAgICAgICAgIHdpZHRoPVwiMXJlbVwiXG4gICAgICAgICAgICAgICBoZWlnaHQ9XCIxcmVtXCJcbiAgICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCAxNiAxNlwiPlxuICAgICAgICAgICAgPHBhdGggZD1cIiR7cGF0aH1cIi8+XG4gICAgICAgICAgPC9zdmc+XG4gICAgICAgIDwvcGYtdjYtYnV0dG9uPlxuICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgYDtcbiAgfVxuXG4gIGFzeW5jIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIC8vIExvYWQgY2xpZW50LW9ubHkgbW9kdWxlcyBCRUZPUkUgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKSBzbyB0aGV5J3JlXG4gICAgLy8gYXZhaWxhYmxlIHdoZW4gTGl0J3MgdXBkYXRlIGN5Y2xlIHJ1bnMuIGZpcnN0VXBkYXRlZCgpIGNhbGxzXG4gICAgLy8gI3NldHVwVHJlZVN0YXRlUGVyc2lzdGVuY2UgYW5kICNzZXR1cFNpZGViYXJTdGF0ZVBlcnNpc3RlbmNlIHdoaWNoXG4gICAgLy8gcmVmZXJlbmNlIFN0YXRlUGVyc2lzdGVuY2UuXG4gICAgaWYgKCF0aGlzLiNjbGllbnRNb2R1bGVzTG9hZGVkKSB7XG4gICAgICBbeyBDRU1SZWxvYWRDbGllbnQgfSwgeyBTdGF0ZVBlcnNpc3RlbmNlIH1dID0gYXdhaXQgUHJvbWlzZS5hbGwoW1xuICAgICAgICAvLyBAdHMtaWdub3JlIC0tIHBsYWluIEpTIG1vZHVsZXMgc2VydmVkIGF0IHJ1bnRpbWUgYnkgR28gc2VydmVyXG4gICAgICAgIGltcG9ydCgnL19fY2VtL3dlYnNvY2tldC1jbGllbnQuanMnKSxcbiAgICAgICAgLy8gQHRzLWlnbm9yZVxuICAgICAgICBpbXBvcnQoJy9fX2NlbS9zdGF0ZS1wZXJzaXN0ZW5jZS5qcycpLFxuICAgICAgXSk7XG4gICAgICAvLyBAdHMtaWdub3JlXG4gICAgICBpbXBvcnQoJy9fX2NlbS9oZWFsdGgtYmFkZ2VzLmpzJykuY2F0Y2goKGU6IHVua25vd24pID0+XG4gICAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmVdIEZhaWxlZCB0byBsb2FkIGhlYWx0aC1iYWRnZXM6JywgZSkpO1xuICAgICAgdGhpcy4jY2xpZW50TW9kdWxlc0xvYWRlZCA9IHRydWU7XG4gICAgfVxuXG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcblxuICAgIC8vIEluaXQgYWZ0ZXIgc3VwZXIgc28gdGhlIGVsZW1lbnQgaXMgZnVsbHkgY29ubmVjdGVkXG4gICAgaWYgKHRoaXMuI3dzQ2xpZW50ID09IG51bGwpIHtcbiAgICAgIHRoaXMuI2luaXRXc0NsaWVudCgpO1xuICAgIH1cbiAgICB0aGlzLiNtaWdyYXRlRnJvbUxvY2FsU3RvcmFnZUlmTmVlZGVkKCk7XG4gIH1cblxuICBmaXJzdFVwZGF0ZWQoKSB7XG4gICAgLy8gU2V0IHVwIGRlYnVnIG92ZXJsYXlcbiAgICB0aGlzLiNzZXR1cERlYnVnT3ZlcmxheSgpO1xuXG4gICAgLy8gU2V0IHVwIGNvbG9yIHNjaGVtZSB0b2dnbGVcbiAgICB0aGlzLiNzZXR1cENvbG9yU2NoZW1lVG9nZ2xlKCk7XG5cbiAgICAvLyBTZXQgdXAgZm9vdGVyIGRyYXdlciBhbmQgdGFic1xuICAgIHRoaXMuI3NldHVwRm9vdGVyRHJhd2VyKCk7XG5cbiAgICAvLyBMaXN0ZW4gZm9yIHNlcnZlciBsb2cgbWVzc2FnZXMgZnJvbSBXZWJTb2NrZXRcbiAgICB0aGlzLiNzZXR1cExvZ0xpc3RlbmVyKCk7XG5cbiAgICAvLyBTZXQgdXAga25vYiBldmVudCBjb29yZGluYXRpb25cbiAgICB0aGlzLiNzZXR1cEtub2JDb29yZGluYXRpb24oKTtcblxuICAgIC8vIFNldCB1cCB0cmVlIHN0YXRlIHBlcnNpc3RlbmNlXG4gICAgdGhpcy4jc2V0dXBUcmVlU3RhdGVQZXJzaXN0ZW5jZSgpO1xuXG4gICAgLy8gU2V0IHVwIHNpZGViYXIgc3RhdGUgcGVyc2lzdGVuY2VcbiAgICB0aGlzLiNzZXR1cFNpZGViYXJTdGF0ZVBlcnNpc3RlbmNlKCk7XG5cbiAgICAvLyBTZXQgdXAgZWxlbWVudCBldmVudCBjYXB0dXJlXG4gICAgdGhpcy4jc2V0dXBFdmVudENhcHR1cmUoKS50aGVuKCgpID0+IHtcbiAgICAgIHRoaXMuI3NldHVwRXZlbnRMaXN0ZW5lcnMoKTtcbiAgICB9KTtcblxuICAgIC8vIFNldCB1cCByZWNvbm5lY3Rpb24gbW9kYWwgYnV0dG9uIGhhbmRsZXJzXG4gICAgLyogYzggaWdub3JlIHN0YXJ0IC0gd2luZG93LmxvY2F0aW9uLnJlbG9hZCB3b3VsZCByZWxvYWQgdGVzdCBwYWdlICovXG4gICAgdGhpcy4jJCgncmVsb2FkLWJ1dHRvbicpPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgIHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKTtcbiAgICB9KTtcbiAgICAvKiBjOCBpZ25vcmUgc3RvcCAqL1xuXG4gICAgdGhpcy4jJCgncmV0cnktYnV0dG9uJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LmNsb3NlKCk7XG4gICAgICB0aGlzLiN3c0NsaWVudC5yZXRyeSgpO1xuICAgIH0pO1xuXG4gICAgLy8gSW5pdGlhbGl6ZSBXZWJTb2NrZXQgY29ubmVjdGlvblxuICAgIHRoaXMuI3dzQ2xpZW50LmluaXQoKTtcbiAgfVxuXG4gIGFzeW5jICNmZXRjaERlYnVnSW5mbygpIHtcbiAgICAvLyBQb3B1bGF0ZSBicm93c2VyIGluZm8gaW1tZWRpYXRlbHlcbiAgICBjb25zdCBicm93c2VyRWwgPSB0aGlzLiMkKCdkZWJ1Zy1icm93c2VyJyk7XG4gICAgY29uc3QgdWFFbCA9IHRoaXMuIyQoJ2RlYnVnLXVhJyk7XG4gICAgaWYgKGJyb3dzZXJFbCkge1xuICAgICAgY29uc3QgYnJvd3NlciA9IHRoaXMuI2RldGVjdEJyb3dzZXIoKTtcbiAgICAgIGJyb3dzZXJFbC50ZXh0Q29udGVudCA9IGJyb3dzZXI7XG4gICAgfVxuICAgIGlmICh1YUVsKSB7XG4gICAgICB1YUVsLnRleHRDb250ZW50ID0gbmF2aWdhdG9yLnVzZXJBZ2VudDtcbiAgICB9XG5cbiAgICAvLyBGZXRjaCBzZXJ2ZXIgZGVidWcgaW5mb1xuICAgIGNvbnN0IGRhdGEgPSBhd2FpdCBmZXRjaCgnL19fY2VtL2RlYnVnJylcbiAgICAgIC50aGVuKHJlcyA9PiByZXMuanNvbigpKVxuICAgICAgLmNhdGNoKChlcnI6IEVycm9yKSA9PiB7XG4gICAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gZmV0Y2ggZGVidWcgaW5mbzonLCBlcnIpO1xuICAgICAgfSk7XG5cbiAgICBpZiAoIWRhdGEpIHJldHVybjtcbiAgICBjb25zdCB2ZXJzaW9uRWwgPSB0aGlzLiMkKCdkZWJ1Zy12ZXJzaW9uJyk7XG4gICAgY29uc3Qgb3NFbCA9IHRoaXMuIyQoJ2RlYnVnLW9zJyk7XG4gICAgY29uc3Qgd2F0Y2hEaXJFbCA9IHRoaXMuIyQoJ2RlYnVnLXdhdGNoLWRpcicpO1xuICAgIGNvbnN0IG1hbmlmZXN0U2l6ZUVsID0gdGhpcy4jJCgnZGVidWctbWFuaWZlc3Qtc2l6ZScpO1xuICAgIGNvbnN0IGRlbW9Db3VudEVsID0gdGhpcy4jJCgnZGVidWctZGVtby1jb3VudCcpO1xuICAgIGNvbnN0IGRlbW9VcmxzQ29udGFpbmVyID0gdGhpcy4jJCgnZGVtby11cmxzLWNvbnRhaW5lcicpO1xuICAgIGNvbnN0IGltcG9ydE1hcEVsID0gdGhpcy4jJCgnZGVidWctaW1wb3J0bWFwJyk7XG5cbiAgICBpZiAodmVyc2lvbkVsKSB2ZXJzaW9uRWwudGV4dENvbnRlbnQgPSBkYXRhLnZlcnNpb24gfHwgJy0nO1xuICAgIGlmIChvc0VsKSBvc0VsLnRleHRDb250ZW50ID0gZGF0YS5vcyB8fCAnLSc7XG4gICAgaWYgKHdhdGNoRGlyRWwpIHdhdGNoRGlyRWwudGV4dENvbnRlbnQgPSBkYXRhLndhdGNoRGlyIHx8ICctJztcbiAgICBpZiAobWFuaWZlc3RTaXplRWwpIG1hbmlmZXN0U2l6ZUVsLnRleHRDb250ZW50ID0gZGF0YS5tYW5pZmVzdFNpemUgfHwgJy0nO1xuICAgIGlmIChkZW1vQ291bnRFbCkgZGVtb0NvdW50RWwudGV4dENvbnRlbnQgPSBkYXRhLmRlbW9Db3VudCB8fCAnMCc7XG5cbiAgICBpZiAoZGVtb1VybHNDb250YWluZXIpIHtcbiAgICAgIHRoaXMuI3BvcHVsYXRlRGVtb1VybHMoZGVtb1VybHNDb250YWluZXIsIGRhdGEuZGVtb3MpO1xuICAgIH1cblxuICAgIGlmIChpbXBvcnRNYXBFbCAmJiBkYXRhLmltcG9ydE1hcCkge1xuICAgICAgY29uc3QgaW1wb3J0TWFwSlNPTiA9IEpTT04uc3RyaW5naWZ5KGRhdGEuaW1wb3J0TWFwLCBudWxsLCAyKTtcbiAgICAgIGltcG9ydE1hcEVsLnRleHRDb250ZW50ID0gaW1wb3J0TWFwSlNPTjtcbiAgICAgIGRhdGEuaW1wb3J0TWFwSlNPTiA9IGltcG9ydE1hcEpTT047XG4gICAgfSBlbHNlIGlmIChpbXBvcnRNYXBFbCkge1xuICAgICAgaW1wb3J0TWFwRWwudGV4dENvbnRlbnQgPSAnTm8gaW1wb3J0IG1hcCBnZW5lcmF0ZWQnO1xuICAgIH1cblxuICAgIHRoaXMuI2RlYnVnRGF0YSA9IGRhdGE7XG4gIH1cblxuICAjcG9wdWxhdGVEZW1vVXJscyhjb250YWluZXI6IEhUTUxFbGVtZW50LCBkZW1vczogRGVidWdEYXRhWydkZW1vcyddKSB7XG4gICAgaWYgKCFkZW1vcz8ubGVuZ3RoKSB7XG4gICAgICBjb250YWluZXIudGV4dENvbnRlbnQgPSAnTm8gZGVtb3MgZm91bmQgaW4gbWFuaWZlc3QnO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IGN1cnJlbnRUYWdOYW1lID0gdGhpcy5wcmltYXJ5VGFnTmFtZSB8fCAnJztcbiAgICBjb25zdCBpc09uRGVtb1BhZ2UgPSAhIWN1cnJlbnRUYWdOYW1lO1xuXG4gICAgaWYgKGlzT25EZW1vUGFnZSkge1xuICAgICAgY29uc3QgY3VycmVudERlbW8gPSBkZW1vcy5maW5kKGRlbW8gPT4gZGVtby50YWdOYW1lID09PSBjdXJyZW50VGFnTmFtZSk7XG4gICAgICBpZiAoIWN1cnJlbnREZW1vKSB7XG4gICAgICAgIGNvbnRhaW5lci50ZXh0Q29udGVudCA9ICdDdXJyZW50IGRlbW8gbm90IGZvdW5kIGluIG1hbmlmZXN0JztcbiAgICAgICAgcmV0dXJuO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBmcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vSW5mb1RlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGFnTmFtZVwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLnRhZ05hbWU7XG4gICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNhbm9uaWNhbFVSTFwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmNhbm9uaWNhbFVSTDtcbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibG9jYWxSb3V0ZVwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmxvY2FsUm91dGU7XG5cbiAgICAgIGNvbnN0IGRlc2NyaXB0aW9uR3JvdXAgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImRlc2NyaXB0aW9uXCJdJyk7XG4gICAgICBpZiAoY3VycmVudERlbW8uZGVzY3JpcHRpb24pIHtcbiAgICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJkZXNjcmlwdGlvblwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmRlc2NyaXB0aW9uO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgZGVzY3JpcHRpb25Hcm91cD8ucmVtb3ZlKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGZpbGVwYXRoR3JvdXAgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImZpbGVwYXRoXCJdJyk7XG4gICAgICBpZiAoY3VycmVudERlbW8uZmlsZXBhdGgpIHtcbiAgICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJmaWxlcGF0aFwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmZpbGVwYXRoO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgZmlsZXBhdGhHcm91cD8ucmVtb3ZlKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnRhaW5lci5yZXBsYWNlQ2hpbGRyZW4oZnJhZ21lbnQpO1xuICAgIH0gZWxzZSB7XG4gICAgICBjb25zdCBsaXN0RnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZGVtb0xpc3RUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICBjb25zdCBncm91cHNDb250YWluZXIgPSBsaXN0RnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtY29udGFpbmVyPVwiZ3JvdXBzXCJdJykhO1xuXG4gICAgICBmb3IgKGNvbnN0IGRlbW8gb2YgZGVtb3MpIHtcbiAgICAgICAgY29uc3QgZ3JvdXBGcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vR3JvdXBUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0YWdOYW1lXCJdJykhLnRleHRDb250ZW50ID0gZGVtby50YWdOYW1lO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIl0nKSEudGV4dENvbnRlbnQgPVxuICAgICAgICAgIGRlbW8uZGVzY3JpcHRpb24gfHwgJyhubyBkZXNjcmlwdGlvbiknO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5jYW5vbmljYWxVUkw7XG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5sb2NhbFJvdXRlO1xuXG4gICAgICAgIGNvbnN0IGZpbGVwYXRoR3JvdXAgPSBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIl0nKTtcbiAgICAgICAgaWYgKGRlbW8uZmlsZXBhdGgpIHtcbiAgICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIl0nKSEudGV4dENvbnRlbnQgPSBkZW1vLmZpbGVwYXRoO1xuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIGZpbGVwYXRoR3JvdXA/LnJlbW92ZSgpO1xuICAgICAgICB9XG5cbiAgICAgICAgZ3JvdXBzQ29udGFpbmVyLmFwcGVuZENoaWxkKGdyb3VwRnJhZ21lbnQpO1xuICAgICAgfVxuXG4gICAgICBjb250YWluZXIucmVwbGFjZUNoaWxkcmVuKGxpc3RGcmFnbWVudCk7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwTG9nTGlzdGVuZXIoKSB7XG4gICAgdGhpcy4jbG9nQ29udGFpbmVyID0gdGhpcy4jJCgnbG9nLWNvbnRhaW5lcicpO1xuXG4gICAgY29uc3QgbG9nc0ZpbHRlciA9IHRoaXMuIyQoJ2xvZ3MtZmlsdGVyJykgYXMgSFRNTElucHV0RWxlbWVudCB8IG51bGw7XG4gICAgaWYgKGxvZ3NGaWx0ZXIpIHtcbiAgICAgIGxvZ3NGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignaW5wdXQnLCAoKSA9PiB7XG4gICAgICAgIGNvbnN0IHsgdmFsdWUgPSAnJyB9ID0gbG9nc0ZpbHRlcjtcbiAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyISk7XG4gICAgICAgIHRoaXMuI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgdGhpcy4jZmlsdGVyTG9ncyh2YWx1ZSk7XG4gICAgICAgIH0sIDMwMCk7XG4gICAgICB9KTtcbiAgICB9XG5cbiAgICB0aGlzLiNsb2dMZXZlbERyb3Bkb3duID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcjbG9nLWxldmVsLWZpbHRlcicpID8/IG51bGw7XG4gICAgaWYgKHRoaXMuI2xvZ0xldmVsRHJvcGRvd24pIHtcbiAgICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICAgIHRoaXMuI2xvYWRMb2dGaWx0ZXJTdGF0ZSgpO1xuICAgICAgfSk7XG4gICAgICB0aGlzLiNsb2dMZXZlbERyb3Bkb3duLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZUxvZ0ZpbHRlckNoYW5nZSBhcyBFdmVudExpc3RlbmVyKTtcbiAgICB9XG5cbiAgICB0aGlzLiMkKCdjb3B5LWxvZ3MnKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB0aGlzLiNjb3B5TG9ncygpO1xuICAgIH0pO1xuXG4gICAgdGhpcy4jaGFuZGxlTG9nc0V2ZW50ID0gKChldmVudDogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IGxvZ3MgPSAoZXZlbnQgYXMgQ2VtTG9nc0V2ZW50KS5sb2dzO1xuICAgICAgaWYgKGxvZ3MpIHtcbiAgICAgICAgdGhpcy4jcmVuZGVyTG9ncyhsb2dzKTtcbiAgICAgIH1cbiAgICB9KTtcbiAgICB3aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcignY2VtOmxvZ3MnLCB0aGlzLiNoYW5kbGVMb2dzRXZlbnQpO1xuICB9XG5cbiAgI2ZpbHRlckxvZ3MocXVlcnk6IHN0cmluZykge1xuICAgIHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgZm9yIChjb25zdCBlbnRyeSBvZiB0aGlzLiNsb2dDb250YWluZXIuY2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IHRleHQgPSBlbnRyeS50ZXh0Q29udGVudD8udG9Mb3dlckNhc2UoKSA/PyAnJztcbiAgICAgIGNvbnN0IHRleHRNYXRjaCA9ICF0aGlzLiNsb2dzRmlsdGVyVmFsdWUgfHwgdGV4dC5pbmNsdWRlcyh0aGlzLiNsb2dzRmlsdGVyVmFsdWUpO1xuXG4gICAgICBjb25zdCBsb2dUeXBlID0gdGhpcy4jZ2V0TG9nVHlwZUZyb21FbnRyeShlbnRyeSk7XG4gICAgICBjb25zdCBsZXZlbE1hdGNoID0gdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmhhcyhsb2dUeXBlKTtcblxuICAgICAgKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5oaWRkZW4gPSAhKHRleHRNYXRjaCAmJiBsZXZlbE1hdGNoKTtcbiAgICB9XG4gIH1cblxuICAjZ2V0TG9nVHlwZUZyb21FbnRyeShlbnRyeTogRWxlbWVudCk6IHN0cmluZyB7XG4gICAgZm9yIChjb25zdCBjbHMgb2YgZW50cnkuY2xhc3NMaXN0KSB7XG4gICAgICBpZiAoWydpbmZvJywgJ3dhcm5pbmcnLCAnZXJyb3InLCAnZGVidWcnXS5pbmNsdWRlcyhjbHMpKSB7XG4gICAgICAgIHJldHVybiBjbHMgPT09ICd3YXJuaW5nJyA/ICd3YXJuJyA6IGNscztcbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuICdpbmZvJztcbiAgfVxuXG4gICNsb2FkTG9nRmlsdGVyU3RhdGUoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHNhdmVkID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1sb2ctZmlsdGVycycpO1xuICAgICAgaWYgKHNhdmVkKSB7XG4gICAgICAgIHRoaXMuI2xvZ0xldmVsRmlsdGVycyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZCkpO1xuICAgICAgfVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIGNvbnNvbGUuZGVidWcoJ1tjZW0tc2VydmUtY2hyb21lXSBsb2NhbFN0b3JhZ2UgdW5hdmFpbGFibGUsIHVzaW5nIGRlZmF1bHQgbG9nIGZpbHRlcnMnKTtcbiAgICB9XG4gICAgdGhpcy4jc3luY0NoZWNrYm94U3RhdGVzKCk7XG4gIH1cblxuICAjc3luY0NoZWNrYm94U3RhdGVzKCkge1xuICAgIGlmICghdGhpcy4jbG9nTGV2ZWxEcm9wZG93bikgcmV0dXJuO1xuICAgIGNvbnN0IG1lbnVJdGVtcyA9IHRoaXMuI2xvZ0xldmVsRHJvcGRvd24ucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgbWVudUl0ZW1zLmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBjb25zdCB2YWx1ZSA9IChpdGVtIGFzIGFueSkudmFsdWU7XG4gICAgICAoaXRlbSBhcyBhbnkpLmNoZWNrZWQgPSB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuaGFzKHZhbHVlKTtcbiAgICB9KTtcbiAgfVxuXG4gICNzYXZlTG9nRmlsdGVyU3RhdGUoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdjZW0tc2VydmUtbG9nLWZpbHRlcnMnLFxuICAgICAgICBKU09OLnN0cmluZ2lmeShbLi4udGhpcy4jbG9nTGV2ZWxGaWx0ZXJzXSkpO1xuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIC8vIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSAocHJpdmF0ZSBtb2RlKSwgc2lsZW50bHkgY29udGludWVcbiAgICB9XG4gIH1cblxuICAjaGFuZGxlTG9nRmlsdGVyQ2hhbmdlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHsgdmFsdWUsIGNoZWNrZWQgfSA9IGV2ZW50IGFzIEV2ZW50ICYgeyB2YWx1ZTogc3RyaW5nOyBjaGVja2VkOiBib29sZWFuIH07XG5cbiAgICBpZiAoY2hlY2tlZCkge1xuICAgICAgdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2xvZ0xldmVsRmlsdGVycy5kZWxldGUodmFsdWUpO1xuICAgIH1cblxuICAgIHRoaXMuI3NhdmVMb2dGaWx0ZXJTdGF0ZSgpO1xuICAgIHRoaXMuI2ZpbHRlckxvZ3ModGhpcy4jbG9nc0ZpbHRlclZhbHVlKTtcbiAgfTtcblxuICBhc3luYyAjY29weUxvZ3MoKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGNvbnN0IGxvZ3MgPSBBcnJheS5mcm9tKHRoaXMuI2xvZ0NvbnRhaW5lci5jaGlsZHJlbilcbiAgICAgIC5maWx0ZXIoZW50cnkgPT4gIShlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuKVxuICAgICAgLm1hcChlbnRyeSA9PiB7XG4gICAgICAgIGNvbnN0IHR5cGUgPSBlbnRyeS5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxhYmVsXCJdJyk/LnRleHRDb250ZW50Py50cmltKCkgfHwgJ0lORk8nO1xuICAgICAgICBjb25zdCB0aW1lID0gZW50cnkucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0aW1lXCJdJyk/LnRleHRDb250ZW50Py50cmltKCkgfHwgJyc7XG4gICAgICAgIGNvbnN0IG1lc3NhZ2UgPSBlbnRyeS5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cIm1lc3NhZ2VcIl0nKT8udGV4dENvbnRlbnQ/LnRyaW0oKSB8fCAnJztcbiAgICAgICAgcmV0dXJuIGBbJHt0eXBlfV0gJHt0aW1lfSAke21lc3NhZ2V9YDtcbiAgICAgIH0pLmpvaW4oJ1xcbicpO1xuXG4gICAgaWYgKCFsb2dzKSByZXR1cm47XG5cbiAgICB0cnkge1xuICAgICAgYXdhaXQgbmF2aWdhdG9yLmNsaXBib2FyZC53cml0ZVRleHQobG9ncyk7XG4gICAgICBjb25zdCBidG4gPSB0aGlzLiMkKCdjb3B5LWxvZ3MnKTtcbiAgICAgIGlmIChidG4pIHtcbiAgICAgICAgY29uc3QgdGV4dE5vZGUgPSBBcnJheS5mcm9tKGJ0bi5jaGlsZE5vZGVzKS5maW5kKFxuICAgICAgICAgIG4gPT4gbi5ub2RlVHlwZSA9PT0gTm9kZS5URVhUX05PREUgJiYgKG4udGV4dENvbnRlbnQ/LnRyaW0oKS5sZW5ndGggPz8gMCkgPiAwXG4gICAgICAgICk7XG4gICAgICAgIGlmICh0ZXh0Tm9kZSkge1xuICAgICAgICAgIGNvbnN0IG9yaWdpbmFsID0gdGV4dE5vZGUudGV4dENvbnRlbnQ7XG4gICAgICAgICAgdGV4dE5vZGUudGV4dENvbnRlbnQgPSAnQ29waWVkISc7XG5cbiAgICAgICAgICBpZiAodGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgdGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICAgIGlmICh0aGlzLmlzQ29ubmVjdGVkICYmIHRleHROb2RlLnBhcmVudE5vZGUpIHtcbiAgICAgICAgICAgICAgdGV4dE5vZGUudGV4dENvbnRlbnQgPSBvcmlnaW5hbDtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICAgICAgICB9LCAyMDAwKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gY2F0Y2ggKGVycikge1xuICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjb3B5IGxvZ3M6JywgZXJyKTtcbiAgICB9XG4gIH1cblxuICAjc2V0dXBEZWJ1Z092ZXJsYXkoKSB7XG4gICAgY29uc3QgZGVidWdCdXR0b24gPSB0aGlzLiMkKCdkZWJ1Zy1pbmZvJyk7XG4gICAgY29uc3QgZGVidWdNb2RhbCA9IHRoaXMuIyQoJ2RlYnVnLW1vZGFsJyk7XG4gICAgY29uc3QgZGVidWdDbG9zZSA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignLmRlYnVnLWNsb3NlJyk7XG4gICAgY29uc3QgZGVidWdDb3B5ID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuZGVidWctY29weScpO1xuXG4gICAgaWYgKGRlYnVnQnV0dG9uICYmIGRlYnVnTW9kYWwpIHtcbiAgICAgIGRlYnVnQnV0dG9uLmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgICB0aGlzLiNmZXRjaERlYnVnSW5mbygpO1xuICAgICAgICAoZGVidWdNb2RhbCBhcyBhbnkpLnNob3dNb2RhbCgpO1xuICAgICAgfSk7XG5cbiAgICAgIGRlYnVnQ2xvc2U/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4gKGRlYnVnTW9kYWwgYXMgYW55KS5jbG9zZSgpKTtcblxuICAgICAgZGVidWdDb3B5Py5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgICAgdGhpcy4jY29weURlYnVnSW5mbygpO1xuICAgICAgfSk7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwRm9vdGVyRHJhd2VyKCkge1xuICAgIGNvbnN0IGRyYXdlciA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignY2VtLWRyYXdlcicpO1xuICAgIGNvbnN0IHRhYnMgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ3BmLXY2LXRhYnMnKTtcblxuICAgIGlmICghZHJhd2VyIHx8ICF0YWJzKSByZXR1cm47XG5cbiAgICB0aGlzLiNkcmF3ZXJPcGVuID0gKGRyYXdlciBhcyBhbnkpLm9wZW47XG5cbiAgICBkcmF3ZXIuYWRkRXZlbnRMaXN0ZW5lcignY2hhbmdlJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICB0aGlzLiNkcmF3ZXJPcGVuID0gKGUgYXMgYW55KS5vcGVuO1xuXG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHtcbiAgICAgICAgZHJhd2VyOiB7IG9wZW46IChlIGFzIGFueSkub3BlbiB9XG4gICAgICB9KTtcblxuICAgICAgaWYgKChlIGFzIGFueSkub3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxMb2dzVG9Cb3R0b20oKTtcbiAgICAgIH1cbiAgICB9KTtcblxuICAgIGRyYXdlci5hZGRFdmVudExpc3RlbmVyKCdyZXNpemUnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgIChkcmF3ZXIgYXMgYW55KS5zZXRBdHRyaWJ1dGUoJ2RyYXdlci1oZWlnaHQnLCAoZSBhcyBhbnkpLmhlaWdodCk7XG5cbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoe1xuICAgICAgICBkcmF3ZXI6IHsgaGVpZ2h0OiAoZSBhcyBhbnkpLmhlaWdodCB9XG4gICAgICB9KTtcbiAgICB9KTtcblxuICAgIHRhYnMuYWRkRXZlbnRMaXN0ZW5lcignY2hhbmdlJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHtcbiAgICAgICAgdGFiczogeyBzZWxlY3RlZEluZGV4OiAoZSBhcyBhbnkpLnNlbGVjdGVkSW5kZXggfVxuICAgICAgfSk7XG5cbiAgICAgIGlmICgoZSBhcyBhbnkpLnNlbGVjdGVkSW5kZXggPT09IDIgJiYgKGRyYXdlciBhcyBhbnkpLm9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTG9nc1RvQm90dG9tKCk7XG4gICAgICB9XG5cbiAgICAgIGlmICgoZSBhcyBhbnkpLnNlbGVjdGVkSW5kZXggPT09IDMgJiYgKGRyYXdlciBhcyBhbnkpLm9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsRXZlbnRzVG9Cb3R0b20oKTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNkZXRlY3RCcm93c2VyKCk6IHN0cmluZyB7XG4gICAgY29uc3QgdWEgPSBuYXZpZ2F0b3IudXNlckFnZW50O1xuICAgIGlmICh1YS5pbmNsdWRlcygnRmlyZWZveC8nKSkge1xuICAgICAgY29uc3QgbWF0Y2ggPSB1YS5tYXRjaCgvRmlyZWZveFxcLyhcXGQrKS8pO1xuICAgICAgcmV0dXJuIG1hdGNoID8gYEZpcmVmb3ggJHttYXRjaFsxXX1gIDogJ0ZpcmVmb3gnO1xuICAgIH0gZWxzZSBpZiAodWEuaW5jbHVkZXMoJ0VkZy8nKSkge1xuICAgICAgY29uc3QgbWF0Y2ggPSB1YS5tYXRjaCgvRWRnXFwvKFxcZCspLyk7XG4gICAgICByZXR1cm4gbWF0Y2ggPyBgRWRnZSAke21hdGNoWzFdfWAgOiAnRWRnZSc7XG4gICAgfSBlbHNlIGlmICh1YS5pbmNsdWRlcygnQ2hyb21lLycpKSB7XG4gICAgICBjb25zdCBtYXRjaCA9IHVhLm1hdGNoKC9DaHJvbWVcXC8oXFxkKykvKTtcbiAgICAgIHJldHVybiBtYXRjaCA/IGBDaHJvbWUgJHttYXRjaFsxXX1gIDogJ0Nocm9tZSc7XG4gICAgfSBlbHNlIGlmICh1YS5pbmNsdWRlcygnU2FmYXJpLycpICYmICF1YS5pbmNsdWRlcygnQ2hyb21lJykpIHtcbiAgICAgIGNvbnN0IG1hdGNoID0gdWEubWF0Y2goL1ZlcnNpb25cXC8oXFxkKykvKTtcbiAgICAgIHJldHVybiBtYXRjaCA/IGBTYWZhcmkgJHttYXRjaFsxXX1gIDogJ1NhZmFyaSc7XG4gICAgfVxuICAgIHJldHVybiAnVW5rbm93bic7XG4gIH1cblxuICBhc3luYyAjY29weURlYnVnSW5mbygpIHtcbiAgICBjb25zdCBpbmZvID0gQXJyYXkuZnJvbSh0aGlzLiMkJCgnI2RlYnVnLW1vZGFsIGRsIGR0JykpLm1hcChkdCA9PiB7XG4gICAgICBjb25zdCBkZCA9IGR0Lm5leHRFbGVtZW50U2libGluZztcbiAgICAgIGlmIChkZCAmJiBkZC50YWdOYW1lID09PSAnREQnKSB7XG4gICAgICAgIHJldHVybiBgJHtkdC50ZXh0Q29udGVudH06ICR7ZGQudGV4dENvbnRlbnR9YDtcbiAgICAgIH1cbiAgICAgIHJldHVybiAnJztcbiAgICB9KS5qb2luKCdcXG4nKTtcblxuICAgIGxldCBpbXBvcnRNYXBTZWN0aW9uID0gJyc7XG4gICAgaWYgKHRoaXMuI2RlYnVnRGF0YT8uaW1wb3J0TWFwSlNPTikge1xuICAgICAgaW1wb3J0TWFwU2VjdGlvbiA9IGBcXG4keyc9Jy5yZXBlYXQoNDApfVxcbkltcG9ydCBNYXA6XFxuJHsnPScucmVwZWF0KDQwKX1cXG4ke3RoaXMuI2RlYnVnRGF0YS5pbXBvcnRNYXBKU09OfVxcbmA7XG4gICAgfVxuXG4gICAgY29uc3QgZGVidWdUZXh0ID0gYENFTSBTZXJ2ZSBEZWJ1ZyBJbmZvcm1hdGlvblxuJHsnPScucmVwZWF0KDQwKX1cbiR7aW5mb30ke2ltcG9ydE1hcFNlY3Rpb259XG4keyc9Jy5yZXBlYXQoNDApfVxuR2VuZXJhdGVkOiAke25ldyBEYXRlKCkudG9JU09TdHJpbmcoKX1gO1xuXG4gICAgdHJ5IHtcbiAgICAgIGF3YWl0IG5hdmlnYXRvci5jbGlwYm9hcmQud3JpdGVUZXh0KGRlYnVnVGV4dCk7XG4gICAgICBjb25zdCBjb3B5QnV0dG9uID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuZGVidWctY29weScpO1xuICAgICAgaWYgKGNvcHlCdXR0b24pIHtcbiAgICAgICAgY29uc3Qgb3JpZ2luYWxUZXh0ID0gY29weUJ1dHRvbi50ZXh0Q29udGVudDtcbiAgICAgICAgY29weUJ1dHRvbi50ZXh0Q29udGVudCA9ICdDb3BpZWQhJztcblxuICAgICAgICBpZiAodGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICAgIH1cblxuICAgICAgICB0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICBpZiAodGhpcy5pc0Nvbm5lY3RlZCAmJiBjb3B5QnV0dG9uLnBhcmVudE5vZGUpIHtcbiAgICAgICAgICAgIGNvcHlCdXR0b24udGV4dENvbnRlbnQgPSBvcmlnaW5hbFRleHQ7XG4gICAgICAgICAgfVxuICAgICAgICAgIHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgICAgIH0sIDIwMDApO1xuICAgICAgfVxuICAgIH0gY2F0Y2ggKGVycikge1xuICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjb3B5IGRlYnVnIGluZm86JywgZXJyKTtcbiAgICB9XG4gIH1cblxuICAjcmVuZGVyTG9ncyhsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT4pIHtcbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgY29uc3QgbG9nRWxlbWVudHMgPSBsb2dzLm1hcChsb2cgPT4ge1xuICAgICAgY29uc3QgZnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jbG9nRW50cnlUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICBjb25zdCBkYXRlID0gbmV3IERhdGUobG9nLmRhdGUpO1xuICAgICAgY29uc3QgdGltZSA9IGRhdGUudG9Mb2NhbGVUaW1lU3RyaW5nKCk7XG5cbiAgICAgIGNvbnN0IGNvbnRhaW5lciA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiY29udGFpbmVyXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICBjb250YWluZXIuY2xhc3NMaXN0LmFkZChsb2cudHlwZSk7XG4gICAgICBjb250YWluZXIuc2V0QXR0cmlidXRlKCdkYXRhLWxvZy1pZCcsIGxvZy5kYXRlKTtcblxuICAgICAgY29uc3QgdHlwZUxhYmVsID0gdGhpcy4jZ2V0TG9nQmFkZ2UobG9nLnR5cGUpO1xuICAgICAgY29uc3Qgc2VhcmNoVGV4dCA9IGAke3R5cGVMYWJlbH0gJHt0aW1lfSAke2xvZy5tZXNzYWdlfWAudG9Mb3dlckNhc2UoKTtcbiAgICAgIGNvbnN0IHRleHRNYXRjaCA9ICF0aGlzLiNsb2dzRmlsdGVyVmFsdWUgfHwgc2VhcmNoVGV4dC5pbmNsdWRlcyh0aGlzLiNsb2dzRmlsdGVyVmFsdWUpO1xuXG4gICAgICBjb25zdCBsb2dUeXBlRm9yRmlsdGVyID0gbG9nLnR5cGUgPT09ICd3YXJuaW5nJyA/ICd3YXJuJyA6IGxvZy50eXBlO1xuICAgICAgY29uc3QgbGV2ZWxNYXRjaCA9IHRoaXMuI2xvZ0xldmVsRmlsdGVycy5oYXMobG9nVHlwZUZvckZpbHRlcik7XG5cbiAgICAgIGlmICghKHRleHRNYXRjaCAmJiBsZXZlbE1hdGNoKSkge1xuICAgICAgICBjb250YWluZXIuc2V0QXR0cmlidXRlKCdoaWRkZW4nLCAnJyk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGxhYmVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsYWJlbFwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgICAgbGFiZWwudGV4dENvbnRlbnQgPSB0aGlzLiNnZXRMb2dCYWRnZShsb2cudHlwZSk7XG4gICAgICB0aGlzLiNhcHBseUxvZ0xhYmVsQXR0cnMobGFiZWwsIGxvZy50eXBlKTtcblxuICAgICAgY29uc3QgdGltZUVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0aW1lXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICB0aW1lRWwuc2V0QXR0cmlidXRlKCdkYXRldGltZScsIGxvZy5kYXRlKTtcbiAgICAgIHRpbWVFbC50ZXh0Q29udGVudCA9IHRpbWU7XG5cbiAgICAgIChmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cIm1lc3NhZ2VcIl0nKSBhcyBIVE1MRWxlbWVudCkudGV4dENvbnRlbnQgPSBsb2cubWVzc2FnZTtcblxuICAgICAgcmV0dXJuIGZyYWdtZW50O1xuICAgIH0pO1xuXG4gICAgaWYgKCF0aGlzLiNpbml0aWFsTG9nc0ZldGNoZWQpIHtcbiAgICAgIHRoaXMuI2xvZ0NvbnRhaW5lci5yZXBsYWNlQ2hpbGRyZW4oLi4ubG9nRWxlbWVudHMpO1xuICAgICAgdGhpcy4jaW5pdGlhbExvZ3NGZXRjaGVkID0gdHJ1ZTtcblxuICAgICAgaWYgKHRoaXMuI2RyYXdlck9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKTtcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy4jbG9nQ29udGFpbmVyLmFwcGVuZCguLi5sb2dFbGVtZW50cyk7XG5cbiAgICAgIGlmICh0aGlzLiNkcmF3ZXJPcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExhdGVzdEludG9WaWV3KCk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI2dldExvZ0JhZGdlKHR5cGU6IHN0cmluZyk6IHN0cmluZyB7XG4gICAgc3dpdGNoICh0eXBlKSB7XG4gICAgICBjYXNlICdpbmZvJzogcmV0dXJuICdJbmZvJztcbiAgICAgIGNhc2UgJ3dhcm5pbmcnOiByZXR1cm4gJ1dhcm4nO1xuICAgICAgY2FzZSAnZXJyb3InOiByZXR1cm4gJ0Vycm9yJztcbiAgICAgIGNhc2UgJ2RlYnVnJzogcmV0dXJuICdEZWJ1Zyc7XG4gICAgICBkZWZhdWx0OiByZXR1cm4gdHlwZS50b1VwcGVyQ2FzZSgpO1xuICAgIH1cbiAgfVxuXG4gICNhcHBseUxvZ0xhYmVsQXR0cnMobGFiZWw6IEhUTUxFbGVtZW50LCB0eXBlOiBzdHJpbmcpIHtcbiAgICBzd2l0Y2ggKHR5cGUpIHtcbiAgICAgIGNhc2UgJ2luZm8nOlxuICAgICAgICByZXR1cm4gbGFiZWwuc2V0QXR0cmlidXRlKCdzdGF0dXMnLCAnaW5mbycpO1xuICAgICAgY2FzZSAnd2FybmluZyc6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICd3YXJuaW5nJyk7XG4gICAgICBjYXNlICdlcnJvcic6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICdkYW5nZXInKTtcbiAgICAgIGNhc2UgJ2RlYnVnJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnY29sb3InLCAncHVycGxlJyk7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICBsYWJlbC5zZXRBdHRyaWJ1dGUoJ2NvbG9yJywgJ2dyZXknKTtcbiAgICB9XG4gIH1cblxuICAjc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICBjb25zdCBsYXN0TG9nID0gdGhpcy4jbG9nQ29udGFpbmVyIS5sYXN0RWxlbWVudENoaWxkO1xuICAgICAgaWYgKGxhc3RMb2cpIHtcbiAgICAgICAgbGFzdExvZy5zY3JvbGxJbnRvVmlldyh7IGJlaGF2aW9yOiAnYXV0bycsIGJsb2NrOiAnZW5kJyB9KTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNzY3JvbGxMb2dzVG9Cb3R0b20oKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGlmICh0aGlzLiNpc0luaXRpYWxMb2FkKSB7XG4gICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgIH0gZWxzZSB7XG4gICAgICBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKTtcbiAgICAgIH0sIDM1MCk7XG4gICAgfVxuICB9XG5cbiAgI21pZ3JhdGVGcm9tTG9jYWxTdG9yYWdlSWZOZWVkZWQoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IGhhc0xvY2FsU3RvcmFnZSA9XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtY29sb3Itc2NoZW1lJykgIT09IG51bGwgfHxcbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1kcmF3ZXItb3BlbicpICE9PSBudWxsIHx8XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtZHJhd2VyLWhlaWdodCcpICE9PSBudWxsIHx8XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtYWN0aXZlLXRhYicpICE9PSBudWxsO1xuXG4gICAgICBpZiAoaGFzTG9jYWxTdG9yYWdlKSB7XG4gICAgICAgIGNvbnN0IG1pZ3JhdGVkID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1taWdyYXRlZC10by1jb29raWVzJyk7XG4gICAgICAgIGlmICghbWlncmF0ZWQpIHtcbiAgICAgICAgICBTdGF0ZVBlcnNpc3RlbmNlLm1pZ3JhdGVGcm9tTG9jYWxTdG9yYWdlKCk7XG4gICAgICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oJ2NlbS1zZXJ2ZS1taWdyYXRlZC10by1jb29raWVzJywgJ3RydWUnKTtcbiAgICAgICAgICBzZXRUaW1lb3V0KCgpID0+IHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKSwgMTAwKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIC8vIGxvY2FsU3RvcmFnZSBub3QgYXZhaWxhYmxlLCBza2lwIG1pZ3JhdGlvblxuICAgIH1cbiAgfVxuXG4gICNzZXR1cENvbG9yU2NoZW1lVG9nZ2xlKCkge1xuICAgIGNvbnN0IHRvZ2dsZUdyb3VwID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuY29sb3Itc2NoZW1lLXRvZ2dsZScpO1xuICAgIGlmICghdG9nZ2xlR3JvdXApIHJldHVybjtcblxuICAgIGNvbnN0IHN0YXRlID0gU3RhdGVQZXJzaXN0ZW5jZS5nZXRTdGF0ZSgpO1xuXG4gICAgdGhpcy4jYXBwbHlDb2xvclNjaGVtZShzdGF0ZS5jb2xvclNjaGVtZSk7XG5cbiAgICBjb25zdCBpdGVtcyA9IHRvZ2dsZUdyb3VwLnF1ZXJ5U2VsZWN0b3JBbGwoJ3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJyk7XG4gICAgaXRlbXMuZm9yRWFjaChpdGVtID0+IHtcbiAgICAgIGlmICgoaXRlbSBhcyBhbnkpLnZhbHVlID09PSBzdGF0ZS5jb2xvclNjaGVtZSkge1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnc2VsZWN0ZWQnLCAnJyk7XG4gICAgICB9XG4gICAgfSk7XG5cbiAgICB0b2dnbGVHcm91cC5hZGRFdmVudExpc3RlbmVyKCdwZi12Ni10b2dnbGUtZ3JvdXAtY2hhbmdlJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICBjb25zdCBzY2hlbWUgPSAoZSBhcyBhbnkpLnZhbHVlO1xuICAgICAgdGhpcy4jYXBwbHlDb2xvclNjaGVtZShzY2hlbWUpO1xuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7IGNvbG9yU2NoZW1lOiBzY2hlbWUgfSk7XG4gICAgfSk7XG4gIH1cblxuICAjYXBwbHlDb2xvclNjaGVtZShzY2hlbWU6IHN0cmluZykge1xuICAgIGNvbnN0IGJvZHkgPSBkb2N1bWVudC5ib2R5O1xuXG4gICAgc3dpdGNoIChzY2hlbWUpIHtcbiAgICAgIGNhc2UgJ2xpZ2h0JzpcbiAgICAgICAgYm9keS5zdHlsZS5jb2xvclNjaGVtZSA9ICdsaWdodCc7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnZGFyayc6XG4gICAgICAgIGJvZHkuc3R5bGUuY29sb3JTY2hlbWUgPSAnZGFyayc7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnc3lzdGVtJzpcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIGJvZHkuc3R5bGUuY29sb3JTY2hlbWUgPSAnbGlnaHQgZGFyayc7XG4gICAgICAgIGJyZWFrO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cEtub2JDb29yZGluYXRpb24oKSB7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOmF0dHJpYnV0ZS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICB9XG5cbiAgI29uS25vYkNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLiNnZXRLbm9iVGFyZ2V0KGV2ZW50KTtcbiAgICBpZiAoIXRhcmdldCkge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gQ291bGQgbm90IGZpbmQga25vYiB0YXJnZXQgaW5mbyBpbiBldmVudCBwYXRoJyk7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgY29uc3QgeyB0YWdOYW1lLCBpbnN0YW5jZUluZGV4IH0gPSB0YXJnZXQ7XG5cbiAgICBjb25zdCBkZW1vID0gdGhpcy5kZW1vO1xuICAgIGlmICghZGVtbykgcmV0dXJuO1xuXG4gICAgY29uc3Qga25vYlR5cGUgPSB0aGlzLiNnZXRLbm9iVHlwZUZyb21FdmVudChldmVudCk7XG5cbiAgICBjb25zdCBzdWNjZXNzID0gKGRlbW8gYXMgYW55KS5hcHBseUtub2JDaGFuZ2UoXG4gICAgICBrbm9iVHlwZSxcbiAgICAgIChldmVudCBhcyBhbnkpLm5hbWUsXG4gICAgICAoZXZlbnQgYXMgYW55KS52YWx1ZSxcbiAgICAgIHRhZ05hbWUsXG4gICAgICBpbnN0YW5jZUluZGV4XG4gICAgKTtcblxuICAgIGlmICghc3VjY2Vzcykge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGFwcGx5IGtub2IgY2hhbmdlOicsIHtcbiAgICAgICAgdHlwZToga25vYlR5cGUsXG4gICAgICAgIG5hbWU6IChldmVudCBhcyBhbnkpLm5hbWUsXG4gICAgICAgIHRhZ05hbWUsXG4gICAgICAgIGluc3RhbmNlSW5kZXhcbiAgICAgIH0pO1xuICAgIH1cbiAgfTtcblxuICAjb25Lbm9iQ2xlYXIgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gdGhpcy4jZ2V0S25vYlRhcmdldChldmVudCk7XG4gICAgaWYgKCF0YXJnZXQpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIENvdWxkIG5vdCBmaW5kIGtub2IgdGFyZ2V0IGluZm8gaW4gZXZlbnQgcGF0aCcpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IHsgdGFnTmFtZSwgaW5zdGFuY2VJbmRleCB9ID0gdGFyZ2V0O1xuXG4gICAgY29uc3QgZGVtbyA9IHRoaXMuZGVtbztcbiAgICBpZiAoIWRlbW8pIHJldHVybjtcblxuICAgIGNvbnN0IGtub2JUeXBlID0gdGhpcy4jZ2V0S25vYlR5cGVGcm9tQ2xlYXJFdmVudChldmVudCk7XG4gICAgY29uc3QgY2xlYXJWYWx1ZSA9IGtub2JUeXBlID09PSAncHJvcGVydHknID8gdW5kZWZpbmVkIDogJyc7XG5cbiAgICBjb25zdCBzdWNjZXNzID0gKGRlbW8gYXMgYW55KS5hcHBseUtub2JDaGFuZ2UoXG4gICAgICBrbm9iVHlwZSxcbiAgICAgIChldmVudCBhcyBhbnkpLm5hbWUsXG4gICAgICBjbGVhclZhbHVlLFxuICAgICAgdGFnTmFtZSxcbiAgICAgIGluc3RhbmNlSW5kZXhcbiAgICApO1xuXG4gICAgaWYgKCFzdWNjZXNzKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY2xlYXIga25vYjonLCB7XG4gICAgICAgIHR5cGU6IGtub2JUeXBlLFxuICAgICAgICBuYW1lOiAoZXZlbnQgYXMgYW55KS5uYW1lLFxuICAgICAgICB0YWdOYW1lLFxuICAgICAgICBpbnN0YW5jZUluZGV4XG4gICAgICB9KTtcbiAgICB9XG4gIH07XG5cbiAgI2dldEtub2JUYXJnZXQoZXZlbnQ6IEV2ZW50KTogeyB0YWdOYW1lOiBzdHJpbmc7IGluc3RhbmNlSW5kZXg6IG51bWJlciB9IHwgbnVsbCB7XG4gICAgY29uc3QgZGVmYXVsdFRhZ05hbWUgPSB0aGlzLnByaW1hcnlUYWdOYW1lIHx8ICcnO1xuXG4gICAgaWYgKGV2ZW50LmNvbXBvc2VkUGF0aCkge1xuICAgICAgZm9yIChjb25zdCBlbGVtZW50IG9mIGV2ZW50LmNvbXBvc2VkUGF0aCgpKSB7XG4gICAgICAgIGlmICghKGVsZW1lbnQgaW5zdGFuY2VvZiBFbGVtZW50KSkgY29udGludWU7XG5cbiAgICAgICAgaWYgKChlbGVtZW50IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0Py5pc0VsZW1lbnRLbm9iID09PSAndHJ1ZScpIHtcbiAgICAgICAgICBjb25zdCB0YWdOYW1lID0gKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQudGFnTmFtZSB8fCBkZWZhdWx0VGFnTmFtZTtcbiAgICAgICAgICBsZXQgaW5zdGFuY2VJbmRleCA9IE51bWJlci5wYXJzZUludCgoZWxlbWVudCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5pbnN0YW5jZUluZGV4ID8/ICcnLCAxMCk7XG4gICAgICAgICAgaWYgKE51bWJlci5pc05hTihpbnN0YW5jZUluZGV4KSkgaW5zdGFuY2VJbmRleCA9IDA7XG4gICAgICAgICAgcmV0dXJuIHsgdGFnTmFtZSwgaW5zdGFuY2VJbmRleCB9O1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIHsgdGFnTmFtZTogZGVmYXVsdFRhZ05hbWUsIGluc3RhbmNlSW5kZXg6IDAgfTtcbiAgfVxuXG4gICNnZXRLbm9iVHlwZUZyb21FdmVudChldmVudDogRXZlbnQpOiBzdHJpbmcge1xuICAgIHN3aXRjaCAoZXZlbnQudHlwZSkge1xuICAgICAgY2FzZSAna25vYjphdHRyaWJ1dGUtY2hhbmdlJzpcbiAgICAgICAgcmV0dXJuICdhdHRyaWJ1dGUnO1xuICAgICAgY2FzZSAna25vYjpwcm9wZXJ0eS1jaGFuZ2UnOlxuICAgICAgICByZXR1cm4gJ3Byb3BlcnR5JztcbiAgICAgIGNhc2UgJ2tub2I6Y3NzLXByb3BlcnR5LWNoYW5nZSc6XG4gICAgICAgIHJldHVybiAnY3NzLXByb3BlcnR5JztcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIHJldHVybiAndW5rbm93bic7XG4gICAgfVxuICB9XG5cbiAgI2dldEtub2JUeXBlRnJvbUNsZWFyRXZlbnQoZXZlbnQ6IEV2ZW50KTogc3RyaW5nIHtcbiAgICBzd2l0Y2ggKGV2ZW50LnR5cGUpIHtcbiAgICAgIGNhc2UgJ2tub2I6YXR0cmlidXRlLWNsZWFyJzpcbiAgICAgICAgcmV0dXJuICdhdHRyaWJ1dGUnO1xuICAgICAgY2FzZSAna25vYjpwcm9wZXJ0eS1jbGVhcic6XG4gICAgICAgIHJldHVybiAncHJvcGVydHknO1xuICAgICAgY2FzZSAna25vYjpjc3MtcHJvcGVydHktY2xlYXInOlxuICAgICAgICByZXR1cm4gJ2Nzcy1wcm9wZXJ0eSc7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICByZXR1cm4gJ3Vua25vd24nO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cFRyZWVTdGF0ZVBlcnNpc3RlbmNlKCkge1xuICAgIHRoaXMuI2hhbmRsZVRyZWVFeHBhbmQgPSAoZTogRXZlbnQpID0+IHtcbiAgICAgIGlmICgoZS50YXJnZXQgYXMgRWxlbWVudCk/LnRhZ05hbWUgIT09ICdQRi1WNi1UUkVFLUlURU0nKSByZXR1cm47XG5cbiAgICAgIGNvbnN0IG5vZGVJZCA9IHRoaXMuI2dldFRyZWVOb2RlSWQoZS50YXJnZXQgYXMgRWxlbWVudCk7XG4gICAgICBjb25zdCB0cmVlU3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFRyZWVTdGF0ZSgpO1xuICAgICAgaWYgKCF0cmVlU3RhdGUuZXhwYW5kZWQuaW5jbHVkZXMobm9kZUlkKSkge1xuICAgICAgICB0cmVlU3RhdGUuZXhwYW5kZWQucHVzaChub2RlSWQpO1xuICAgICAgICBTdGF0ZVBlcnNpc3RlbmNlLnNldFRyZWVTdGF0ZSh0cmVlU3RhdGUpO1xuICAgICAgfVxuICAgIH07XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdleHBhbmQnLCB0aGlzLiNoYW5kbGVUcmVlRXhwYW5kKTtcblxuICAgIHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSA9IChlOiBFdmVudCkgPT4ge1xuICAgICAgaWYgKChlLnRhcmdldCBhcyBFbGVtZW50KT8udGFnTmFtZSAhPT0gJ1BGLVY2LVRSRUUtSVRFTScpIHJldHVybjtcblxuICAgICAgY29uc3Qgbm9kZUlkID0gdGhpcy4jZ2V0VHJlZU5vZGVJZChlLnRhcmdldCBhcyBFbGVtZW50KTtcbiAgICAgIGNvbnN0IHRyZWVTdGF0ZSA9IFN0YXRlUGVyc2lzdGVuY2UuZ2V0VHJlZVN0YXRlKCk7XG4gICAgICBjb25zdCBpbmRleCA9IHRyZWVTdGF0ZS5leHBhbmRlZC5pbmRleE9mKG5vZGVJZCk7XG4gICAgICBpZiAoaW5kZXggPiAtMSkge1xuICAgICAgICB0cmVlU3RhdGUuZXhwYW5kZWQuc3BsaWNlKGluZGV4LCAxKTtcbiAgICAgICAgU3RhdGVQZXJzaXN0ZW5jZS5zZXRUcmVlU3RhdGUodHJlZVN0YXRlKTtcbiAgICAgIH1cbiAgICB9O1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignY29sbGFwc2UnLCB0aGlzLiNoYW5kbGVUcmVlQ29sbGFwc2UpO1xuXG4gICAgdGhpcy4jaGFuZGxlVHJlZVNlbGVjdCA9IChlOiBFdmVudCkgPT4ge1xuICAgICAgaWYgKChlLnRhcmdldCBhcyBFbGVtZW50KT8udGFnTmFtZSAhPT0gJ1BGLVY2LVRSRUUtSVRFTScpIHJldHVybjtcblxuICAgICAgY29uc3Qgbm9kZUlkID0gdGhpcy4jZ2V0VHJlZU5vZGVJZChlLnRhcmdldCBhcyBFbGVtZW50KTtcbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlVHJlZVN0YXRlKHsgc2VsZWN0ZWQ6IG5vZGVJZCB9KTtcbiAgICB9O1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlVHJlZVNlbGVjdCk7XG5cbiAgICB0aGlzLiNhcHBseVRyZWVTdGF0ZSgpO1xuICB9XG5cbiAgI2FwcGx5VHJlZVN0YXRlKCkge1xuICAgIGNvbnN0IHRyZWVTdGF0ZSA9IFN0YXRlUGVyc2lzdGVuY2UuZ2V0VHJlZVN0YXRlKCk7XG5cbiAgICBmb3IgKGNvbnN0IG5vZGVJZCBvZiB0cmVlU3RhdGUuZXhwYW5kZWQpIHtcbiAgICAgIGNvbnN0IHRyZWVJdGVtID0gdGhpcy4jZmluZFRyZWVJdGVtQnlJZChub2RlSWQpO1xuICAgICAgaWYgKHRyZWVJdGVtICYmICF0cmVlSXRlbS5oYXNBdHRyaWJ1dGUoJ2V4cGFuZGVkJykpIHtcbiAgICAgICAgdHJlZUl0ZW0uc2V0QXR0cmlidXRlKCdleHBhbmRlZCcsICcnKTtcbiAgICAgIH1cbiAgICB9XG5cbiAgICBpZiAodHJlZVN0YXRlLnNlbGVjdGVkKSB7XG4gICAgICBjb25zdCB0cmVlSXRlbSA9IHRoaXMuI2ZpbmRUcmVlSXRlbUJ5SWQodHJlZVN0YXRlLnNlbGVjdGVkKTtcbiAgICAgIGlmICh0cmVlSXRlbSAmJiAhdHJlZUl0ZW0uaGFzQXR0cmlidXRlKCdjdXJyZW50JykpIHtcbiAgICAgICAgdHJlZUl0ZW0uc2V0QXR0cmlidXRlKCdjdXJyZW50JywgJycpO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNzZXR1cFNpZGViYXJTdGF0ZVBlcnNpc3RlbmNlKCkge1xuICAgIGNvbnN0IHBhZ2UgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ3BmLXY2LXBhZ2UnKTtcblxuICAgIGlmICghcGFnZSkgcmV0dXJuO1xuXG4gICAgcGFnZS5hZGRFdmVudExpc3RlbmVyKCdzaWRlYmFyLXRvZ2dsZScsIChldmVudDogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IGNvbGxhcHNlZCA9ICEoZXZlbnQgYXMgYW55KS5leHBhbmRlZDtcblxuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIHNpZGViYXI6IHsgY29sbGFwc2VkIH1cbiAgICAgIH0pO1xuICAgIH0pO1xuICB9XG5cbiAgI2ZpbmRUcmVlSXRlbUJ5SWQobm9kZUlkOiBzdHJpbmcpOiBFbGVtZW50IHwgbnVsbCB7XG4gICAgY29uc3QgcGFydHMgPSBub2RlSWQuc3BsaXQoJzonKTtcbiAgICBjb25zdCBbdHlwZSwgbW9kdWxlUGF0aCwgdGFnTmFtZSwgbmFtZV0gPSBwYXJ0cztcblxuICAgIGxldCBhdHRyU3VmZml4ID0gJyc7XG4gICAgaWYgKHRhZ05hbWUpIHtcbiAgICAgIGF0dHJTdWZmaXggKz0gYFtkYXRhLXRhZy1uYW1lPVwiJHtDU1MuZXNjYXBlKHRhZ05hbWUpfVwiXWA7XG4gICAgfVxuICAgIGlmIChuYW1lKSB7XG4gICAgICBhdHRyU3VmZml4ICs9IGBbZGF0YS1uYW1lPVwiJHtDU1MuZXNjYXBlKG5hbWUpfVwiXWA7XG4gICAgfVxuXG4gICAgbGV0IHNlbGVjdG9yID0gYHBmLXY2LXRyZWUtaXRlbVtkYXRhLXR5cGU9XCIke0NTUy5lc2NhcGUodHlwZSl9XCJdYDtcbiAgICBpZiAobW9kdWxlUGF0aCkge1xuICAgICAgY29uc3QgZXNjYXBlZE1vZHVsZVBhdGggPSBDU1MuZXNjYXBlKG1vZHVsZVBhdGgpO1xuICAgICAgY29uc3QgZXNjYXBlZFR5cGUgPSBDU1MuZXNjYXBlKHR5cGUpO1xuICAgICAgY29uc3Qgc2VsZWN0b3IxID0gYHBmLXY2LXRyZWUtaXRlbVtkYXRhLXR5cGU9XCIke2VzY2FwZWRUeXBlfVwiXVtkYXRhLW1vZHVsZS1wYXRoPVwiJHtlc2NhcGVkTW9kdWxlUGF0aH1cIl0ke2F0dHJTdWZmaXh9YDtcbiAgICAgIGNvbnN0IHNlbGVjdG9yMiA9IGBwZi12Ni10cmVlLWl0ZW1bZGF0YS10eXBlPVwiJHtlc2NhcGVkVHlwZX1cIl1bZGF0YS1wYXRoPVwiJHtlc2NhcGVkTW9kdWxlUGF0aH1cIl0ke2F0dHJTdWZmaXh9YDtcbiAgICAgIHNlbGVjdG9yID0gYCR7c2VsZWN0b3IxfSwgJHtzZWxlY3RvcjJ9YDtcbiAgICB9IGVsc2Uge1xuICAgICAgc2VsZWN0b3IgKz0gYXR0clN1ZmZpeDtcbiAgICB9XG5cbiAgICByZXR1cm4gdGhpcy5xdWVyeVNlbGVjdG9yKHNlbGVjdG9yKTtcbiAgfVxuXG4gICNnZXRUcmVlTm9kZUlkKHRyZWVJdGVtOiBFbGVtZW50KTogc3RyaW5nIHtcbiAgICBjb25zdCB0eXBlID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXR5cGUnKTtcbiAgICBjb25zdCBtb2R1bGVQYXRoID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLW1vZHVsZS1wYXRoJykgfHwgdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXBhdGgnKTtcbiAgICBjb25zdCB0YWdOYW1lID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXRhZy1uYW1lJyk7XG4gICAgY29uc3QgbmFtZSA9IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS1uYW1lJyk7XG4gICAgY29uc3QgY2F0ZWdvcnkgPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtY2F0ZWdvcnknKTtcblxuICAgIGNvbnN0IHBhcnRzID0gW3R5cGVdO1xuICAgIGlmIChtb2R1bGVQYXRoKSBwYXJ0cy5wdXNoKG1vZHVsZVBhdGgpO1xuICAgIGlmICh0YWdOYW1lKSBwYXJ0cy5wdXNoKHRhZ05hbWUpO1xuICAgIGlmIChjYXRlZ29yeSkge1xuICAgICAgcGFydHMucHVzaChjYXRlZ29yeSk7XG4gICAgfSBlbHNlIGlmIChuYW1lKSB7XG4gICAgICBwYXJ0cy5wdXNoKG5hbWUpO1xuICAgIH1cblxuICAgIHJldHVybiBwYXJ0cy5qb2luKCc6Jyk7XG4gIH1cblxuICAvLyBFdmVudCBEaXNjb3ZlcnkgJiBDYXB0dXJlIE1ldGhvZHNcblxuICBhc3luYyAjZGlzY292ZXJFbGVtZW50RXZlbnRzKCk6IFByb21pc2U8TWFwPHN0cmluZywgRXZlbnRJbmZvPj4ge1xuICAgIHRyeSB7XG4gICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IGZldGNoKCcvY3VzdG9tLWVsZW1lbnRzLmpzb24nKTtcbiAgICAgIGlmICghcmVzcG9uc2Uub2spIHtcbiAgICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gTm8gbWFuaWZlc3QgYXZhaWxhYmxlIGZvciBldmVudCBkaXNjb3ZlcnknKTtcbiAgICAgICAgcmV0dXJuIG5ldyBNYXAoKTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgbWFuaWZlc3QgPSBhd2FpdCByZXNwb25zZS5qc29uKCkgYXMgTWFuaWZlc3Q7XG4gICAgICB0aGlzLiNtYW5pZmVzdCA9IG1hbmlmZXN0O1xuXG4gICAgICBjb25zdCBldmVudE1hcCA9IG5ldyBNYXA8c3RyaW5nLCBFdmVudEluZm8+KCk7XG5cbiAgICAgIGZvciAoY29uc3QgbW9kdWxlIG9mIG1hbmlmZXN0Lm1vZHVsZXMgfHwgW10pIHtcbiAgICAgICAgZm9yIChjb25zdCBkZWNsYXJhdGlvbiBvZiBtb2R1bGUuZGVjbGFyYXRpb25zIHx8IFtdKSB7XG4gICAgICAgICAgaWYgKGRlY2xhcmF0aW9uLmN1c3RvbUVsZW1lbnQgJiYgZGVjbGFyYXRpb24udGFnTmFtZSkge1xuICAgICAgICAgICAgY29uc3QgdGFnTmFtZSA9IGRlY2xhcmF0aW9uLnRhZ05hbWU7XG4gICAgICAgICAgICBjb25zdCBldmVudHMgPSBkZWNsYXJhdGlvbi5ldmVudHMgfHwgW107XG5cbiAgICAgICAgICAgIGlmIChldmVudHMubGVuZ3RoID4gMCkge1xuICAgICAgICAgICAgICBjb25zdCBldmVudE5hbWVzID0gbmV3IFNldChldmVudHMubWFwKGUgPT4gZS5uYW1lKSk7XG4gICAgICAgICAgICAgIGV2ZW50TWFwLnNldCh0YWdOYW1lLCB7XG4gICAgICAgICAgICAgICAgZXZlbnROYW1lcyxcbiAgICAgICAgICAgICAgICBldmVudHM6IGV2ZW50c1xuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgcmV0dXJuIGV2ZW50TWFwO1xuICAgIH0gY2F0Y2ggKGVycm9yKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBFcnJvciBsb2FkaW5nIG1hbmlmZXN0IGZvciBldmVudCBkaXNjb3Zlcnk6JywgZXJyb3IpO1xuICAgICAgcmV0dXJuIG5ldyBNYXAoKTtcbiAgICB9XG4gIH1cblxuICBhc3luYyAjc2V0dXBFdmVudENhcHR1cmUoKSB7XG4gICAgdGhpcy4jZWxlbWVudEV2ZW50TWFwID0gYXdhaXQgdGhpcy4jZGlzY292ZXJFbGVtZW50RXZlbnRzKCk7XG5cbiAgICBpZiAodGhpcy4jZWxlbWVudEV2ZW50TWFwLnNpemUgPT09IDApIHJldHVybjtcblxuICAgIHRoaXMuI2F0dGFjaEV2ZW50TGlzdGVuZXJzKCk7XG4gICAgdGhpcy4jdXBkYXRlRXZlbnRGaWx0ZXJzKCk7XG4gICAgdGhpcy4jb2JzZXJ2ZURlbW9NdXRhdGlvbnMoKTtcbiAgfVxuXG4gICNhdHRhY2hFdmVudExpc3RlbmVycygpIHtcbiAgICBjb25zdCBkZW1vID0gdGhpcy5kZW1vO1xuICAgIGlmICghZGVtbykgcmV0dXJuO1xuXG4gICAgY29uc3Qgcm9vdCA9IGRlbW8uc2hhZG93Um9vdCA/PyBkZW1vO1xuXG4gICAgZm9yIChjb25zdCBbdGFnTmFtZSwgZXZlbnRJbmZvXSBvZiB0aGlzLiNlbGVtZW50RXZlbnRNYXAhKSB7XG4gICAgICBjb25zdCBlbGVtZW50cyA9IHJvb3QucXVlcnlTZWxlY3RvckFsbCh0YWdOYW1lKTtcblxuICAgICAgZm9yIChjb25zdCBlbGVtZW50IG9mIGVsZW1lbnRzKSB7XG4gICAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGV2ZW50SW5mby5ldmVudE5hbWVzKSB7XG4gICAgICAgICAgZWxlbWVudC5hZGRFdmVudExpc3RlbmVyKGV2ZW50TmFtZSwgdGhpcy4jaGFuZGxlRWxlbWVudEV2ZW50LCB7IGNhcHR1cmU6IHRydWUgfSk7XG4gICAgICAgIH1cbiAgICAgICAgKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuY2VtRXZlbnRzQXR0YWNoZWQgPSAndHJ1ZSc7XG4gICAgICAgIHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5hZGQodGFnTmFtZSk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI29ic2VydmVEZW1vTXV0YXRpb25zKCkge1xuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCByb290ID0gZGVtby5zaGFkb3dSb290ID8/IGRlbW87XG5cbiAgICB0aGlzLiNvYnNlcnZlci5vYnNlcnZlKHJvb3QsIHtcbiAgICAgIGNoaWxkTGlzdDogdHJ1ZSxcbiAgICAgIHN1YnRyZWU6IHRydWVcbiAgICB9KTtcbiAgfVxuXG4gICNoYW5kbGVFbGVtZW50RXZlbnQgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgZWxlbWVudCA9IGV2ZW50LmN1cnJlbnRUYXJnZXQ7XG4gICAgaWYgKCEoZWxlbWVudCBpbnN0YW5jZW9mIEhUTUxFbGVtZW50KSkgcmV0dXJuO1xuXG4gICAgY29uc3QgdGFnTmFtZSA9IGVsZW1lbnQudGFnTmFtZS50b0xvd2VyQ2FzZSgpO1xuICAgIGNvbnN0IGV2ZW50SW5mbyA9IHRoaXMuI2VsZW1lbnRFdmVudE1hcD8uZ2V0KHRhZ05hbWUpO1xuXG4gICAgaWYgKCFldmVudEluZm8gfHwgIWV2ZW50SW5mby5ldmVudE5hbWVzLmhhcyhldmVudC50eXBlKSkgcmV0dXJuO1xuXG4gICAgdGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICB0aGlzLiNjYXB0dXJlRXZlbnQoZXZlbnQsIGVsZW1lbnQsIHRhZ05hbWUsIGV2ZW50SW5mbyk7XG4gIH07XG5cbiAgI2dldEV2ZW50RG9jdW1lbnRhdGlvbihtYW5pZmVzdEV2ZW50OiBFdmVudEluZm9bJ2V2ZW50cyddWzBdIHwgdW5kZWZpbmVkKSB7XG4gICAgaWYgKCFtYW5pZmVzdEV2ZW50KSB7XG4gICAgICByZXR1cm4geyBzdW1tYXJ5OiBudWxsLCBkZXNjcmlwdGlvbjogbnVsbCB9O1xuICAgIH1cblxuICAgIGxldCBzdW1tYXJ5ID0gbWFuaWZlc3RFdmVudC5zdW1tYXJ5IHx8IG51bGw7XG4gICAgbGV0IGRlc2NyaXB0aW9uID0gbWFuaWZlc3RFdmVudC5kZXNjcmlwdGlvbiB8fCBudWxsO1xuXG4gICAgaWYgKG1hbmlmZXN0RXZlbnQudHlwZT8udGV4dCAmJiB0aGlzLiNtYW5pZmVzdCkge1xuICAgICAgY29uc3QgdHlwZU5hbWUgPSBtYW5pZmVzdEV2ZW50LnR5cGUudGV4dDtcbiAgICAgIGNvbnN0IHR5cGVEZWNsYXJhdGlvbiA9IHRoaXMuI2ZpbmRUeXBlRGVjbGFyYXRpb24odHlwZU5hbWUpO1xuXG4gICAgICBpZiAodHlwZURlY2xhcmF0aW9uKSB7XG4gICAgICAgIGlmICghc3VtbWFyeSAmJiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSkge1xuICAgICAgICAgIHN1bW1hcnkgPSB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeTtcbiAgICAgICAgfSBlbHNlIGlmICh0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSAmJiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSAhPT0gc3VtbWFyeSkge1xuICAgICAgICAgIHN1bW1hcnkgPSBzdW1tYXJ5ID8gYCR7c3VtbWFyeX1cXG5cXG5Gcm9tICR7dHlwZU5hbWV9OiAke3R5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5fWAgOiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeTtcbiAgICAgICAgfVxuXG4gICAgICAgIGlmICghZGVzY3JpcHRpb24gJiYgdHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uKSB7XG4gICAgICAgICAgZGVzY3JpcHRpb24gPSB0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb247XG4gICAgICAgIH0gZWxzZSBpZiAodHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uICYmIHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbiAhPT0gZGVzY3JpcHRpb24pIHtcbiAgICAgICAgICBkZXNjcmlwdGlvbiA9IGRlc2NyaXB0aW9uID8gYCR7ZGVzY3JpcHRpb259XFxuXFxuJHt0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb259YCA6IHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbjtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiB7IHN1bW1hcnksIGRlc2NyaXB0aW9uIH07XG4gIH1cblxuICAjZmluZFR5cGVEZWNsYXJhdGlvbih0eXBlTmFtZTogc3RyaW5nKSB7XG4gICAgaWYgKCF0aGlzLiNtYW5pZmVzdCkgcmV0dXJuIG51bGw7XG5cbiAgICBmb3IgKGNvbnN0IG1vZHVsZSBvZiB0aGlzLiNtYW5pZmVzdC5tb2R1bGVzIHx8IFtdKSB7XG4gICAgICBmb3IgKGNvbnN0IGRlY2xhcmF0aW9uIG9mIG1vZHVsZS5kZWNsYXJhdGlvbnMgfHwgW10pIHtcbiAgICAgICAgaWYgKGRlY2xhcmF0aW9uLm5hbWUgPT09IHR5cGVOYW1lICYmXG4gICAgICAgICAgICAoZGVjbGFyYXRpb24ua2luZCA9PT0gJ2NsYXNzJyB8fCBkZWNsYXJhdGlvbi5raW5kID09PSAnaW50ZXJmYWNlJykpIHtcbiAgICAgICAgICByZXR1cm4gZGVjbGFyYXRpb24gYXMgeyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9O1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIG51bGw7XG4gIH1cblxuICAjY2FwdHVyZUV2ZW50KGV2ZW50OiBFdmVudCwgdGFyZ2V0OiBIVE1MRWxlbWVudCwgdGFnTmFtZTogc3RyaW5nLCBldmVudEluZm86IEV2ZW50SW5mbykge1xuICAgIGNvbnN0IG1hbmlmZXN0RXZlbnQgPSBldmVudEluZm8uZXZlbnRzLmZpbmQoZSA9PiBlLm5hbWUgPT09IGV2ZW50LnR5cGUpO1xuXG4gICAgY29uc3QgZXZlbnREb2NzID0gdGhpcy4jZ2V0RXZlbnREb2N1bWVudGF0aW9uKG1hbmlmZXN0RXZlbnQpO1xuXG4gICAgY29uc3QgY3VzdG9tUHJvcGVydGllcyA9IHRoaXMuI2V4dHJhY3RFdmVudFByb3BlcnRpZXMoZXZlbnQpO1xuXG4gICAgY29uc3QgZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkID0ge1xuICAgICAgaWQ6IGAke0RhdGUubm93KCl9LSR7TWF0aC5yYW5kb20oKX1gLFxuICAgICAgdGltZXN0YW1wOiBuZXcgRGF0ZSgpLFxuICAgICAgZXZlbnROYW1lOiBldmVudC50eXBlLFxuICAgICAgdGFnTmFtZTogdGFnTmFtZSxcbiAgICAgIGVsZW1lbnRJZDogdGFyZ2V0LmlkIHx8IG51bGwsXG4gICAgICBlbGVtZW50Q2xhc3M6IHRhcmdldC5jbGFzc05hbWUgfHwgbnVsbCxcbiAgICAgIGN1c3RvbVByb3BlcnRpZXM6IGN1c3RvbVByb3BlcnRpZXMsXG4gICAgICBtYW5pZmVzdFR5cGU6IG1hbmlmZXN0RXZlbnQ/LnR5cGU/LnRleHQgfHwgbnVsbCxcbiAgICAgIHN1bW1hcnk6IGV2ZW50RG9jcy5zdW1tYXJ5LFxuICAgICAgZGVzY3JpcHRpb246IGV2ZW50RG9jcy5kZXNjcmlwdGlvbixcbiAgICAgIGJ1YmJsZXM6IGV2ZW50LmJ1YmJsZXMsXG4gICAgICBjb21wb3NlZDogZXZlbnQuY29tcG9zZWQsXG4gICAgICBjYW5jZWxhYmxlOiBldmVudC5jYW5jZWxhYmxlLFxuICAgICAgZGVmYXVsdFByZXZlbnRlZDogZXZlbnQuZGVmYXVsdFByZXZlbnRlZFxuICAgIH07XG5cbiAgICB0aGlzLiNjYXB0dXJlZEV2ZW50cy5wdXNoKGV2ZW50UmVjb3JkKTtcblxuICAgIGlmICh0aGlzLiNjYXB0dXJlZEV2ZW50cy5sZW5ndGggPiB0aGlzLiNtYXhDYXB0dXJlZEV2ZW50cykge1xuICAgICAgdGhpcy4jY2FwdHVyZWRFdmVudHMuc2hpZnQoKTtcbiAgICB9XG5cbiAgICB0aGlzLiNyZW5kZXJFdmVudChldmVudFJlY29yZCk7XG4gIH1cblxuICAjZXh0cmFjdEV2ZW50UHJvcGVydGllcyhldmVudDogRXZlbnQpOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiB7XG4gICAgY29uc3QgcHJvcGVydGllczogUmVjb3JkPHN0cmluZywgdW5rbm93bj4gPSB7fTtcbiAgICBjb25zdCBldmVudFByb3RvdHlwZUtleXMgPSBuZXcgU2V0KE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKEV2ZW50LnByb3RvdHlwZSkpO1xuXG4gICAgY29uc3Qgc2VyaWFsaXplVmFsdWUgPSAodmFsdWU6IHVua25vd24pOiB1bmtub3duID0+IHtcbiAgICAgIHRyeSB7XG4gICAgICAgIHJldHVybiBKU09OLnBhcnNlKEpTT04uc3RyaW5naWZ5KHZhbHVlKSk7XG4gICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgIHRyeSB7XG4gICAgICAgICAgcmV0dXJuIFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIH0gY2F0Y2ggKHN0cmluZ0Vycikge1xuICAgICAgICAgIHJldHVybiAnW05vdCBzZXJpYWxpemFibGVdJztcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH07XG5cbiAgICBpZiAoZXZlbnQgaW5zdGFuY2VvZiBDdXN0b21FdmVudCAmJiBldmVudC5kZXRhaWwgIT09IHVuZGVmaW5lZCkge1xuICAgICAgcHJvcGVydGllcy5kZXRhaWwgPSBzZXJpYWxpemVWYWx1ZShldmVudC5kZXRhaWwpO1xuICAgIH1cblxuICAgIGZvciAoY29uc3Qga2V5IG9mIE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKGV2ZW50KSkge1xuICAgICAgaWYgKCFldmVudFByb3RvdHlwZUtleXMuaGFzKGtleSkgJiYgIWtleS5zdGFydHNXaXRoKCdfJykgJiYgIXByb3BlcnRpZXMuaGFzT3duUHJvcGVydHkoa2V5KSkge1xuICAgICAgICBwcm9wZXJ0aWVzW2tleV0gPSBzZXJpYWxpemVWYWx1ZSgoZXZlbnQgYXMgYW55KVtrZXldKTtcbiAgICAgIH1cbiAgICB9XG5cbiAgICByZXR1cm4gcHJvcGVydGllcztcbiAgfVxuXG4gICNyZW5kZXJFdmVudChldmVudFJlY29yZDogRXZlbnRSZWNvcmQpIHtcbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgY29uc3QgZnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZXZlbnRFbnRyeVRlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICBjb25zdCB0aW1lID0gZXZlbnRSZWNvcmQudGltZXN0YW1wLnRvTG9jYWxlVGltZVN0cmluZygpO1xuXG4gICAgY29uc3QgY29udGFpbmVyID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJjb250YWluZXJcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICBjb250YWluZXIuZGF0YXNldC5ldmVudElkID0gZXZlbnRSZWNvcmQuaWQ7XG4gICAgY29udGFpbmVyLmRhdGFzZXQuZXZlbnRUeXBlID0gZXZlbnRSZWNvcmQuZXZlbnROYW1lO1xuICAgIGNvbnRhaW5lci5kYXRhc2V0LmVsZW1lbnRUeXBlID0gZXZlbnRSZWNvcmQudGFnTmFtZTtcblxuICAgIGNvbnN0IHRleHRNYXRjaCA9IHRoaXMuI2V2ZW50TWF0Y2hlc1RleHRGaWx0ZXIoZXZlbnRSZWNvcmQpO1xuICAgIGNvbnN0IHR5cGVNYXRjaCA9IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmhhcyhldmVudFJlY29yZC5ldmVudE5hbWUpO1xuICAgIGNvbnN0IGVsZW1lbnRNYXRjaCA9IHRoaXMuI2VsZW1lbnRGaWx0ZXJzLnNpemUgPT09IDAgfHwgdGhpcy4jZWxlbWVudEZpbHRlcnMuaGFzKGV2ZW50UmVjb3JkLnRhZ05hbWUpO1xuXG4gICAgaWYgKCEodGV4dE1hdGNoICYmIHR5cGVNYXRjaCAmJiBlbGVtZW50TWF0Y2gpKSB7XG4gICAgICBjb250YWluZXIuc2V0QXR0cmlidXRlKCdoaWRkZW4nLCAnJyk7XG4gICAgfVxuXG4gICAgY29uc3QgbGFiZWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxhYmVsXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgbGFiZWwudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC5ldmVudE5hbWU7XG4gICAgbGFiZWwuc2V0QXR0cmlidXRlKCdzdGF0dXMnLCAnaW5mbycpO1xuXG4gICAgY29uc3QgdGltZUVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0aW1lXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgdGltZUVsLnNldEF0dHJpYnV0ZSgnZGF0ZXRpbWUnLCBldmVudFJlY29yZC50aW1lc3RhbXAudG9JU09TdHJpbmcoKSk7XG4gICAgdGltZUVsLnRleHRDb250ZW50ID0gdGltZTtcblxuICAgIGNvbnN0IGVsZW1lbnRFbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZWxlbWVudFwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgIGxldCBlbGVtZW50VGV4dCA9IGA8JHtldmVudFJlY29yZC50YWdOYW1lfT5gO1xuICAgIGlmIChldmVudFJlY29yZC5lbGVtZW50SWQpIHtcbiAgICAgIGVsZW1lbnRUZXh0ICs9IGAjJHtldmVudFJlY29yZC5lbGVtZW50SWR9YDtcbiAgICB9XG4gICAgZWxlbWVudEVsLnRleHRDb250ZW50ID0gZWxlbWVudFRleHQ7XG5cbiAgICB0aGlzLiNldmVudExpc3QuYXBwZW5kKGZyYWdtZW50KTtcblxuICAgIGlmICghdGhpcy4jc2VsZWN0ZWRFdmVudElkKSB7XG4gICAgICB0aGlzLiNzZWxlY3RFdmVudChldmVudFJlY29yZC5pZCk7XG4gICAgfVxuXG4gICAgaWYgKHRoaXMuI2RyYXdlck9wZW4gJiYgdGhpcy4jaXNFdmVudHNUYWJBY3RpdmUoKSkge1xuICAgICAgdGhpcy4jc2Nyb2xsRXZlbnRzVG9Cb3R0b20oKTtcbiAgICB9XG4gIH1cblxuICAjc2VsZWN0RXZlbnQoZXZlbnRJZDogc3RyaW5nKSB7XG4gICAgY29uc3QgZXZlbnRSZWNvcmQgPSB0aGlzLiNnZXRFdmVudFJlY29yZEJ5SWQoZXZlbnRJZCk7XG4gICAgaWYgKCFldmVudFJlY29yZCkgcmV0dXJuO1xuXG4gICAgdGhpcy4jc2VsZWN0ZWRFdmVudElkID0gZXZlbnRJZDtcblxuICAgIGNvbnN0IGFsbEl0ZW1zID0gdGhpcy4jZXZlbnRMaXN0Py5xdWVyeVNlbGVjdG9yQWxsKCcuZXZlbnQtbGlzdC1pdGVtJyk7XG4gICAgYWxsSXRlbXM/LmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBpZiAoKGl0ZW0gYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuZXZlbnRJZCA9PT0gZXZlbnRJZCkge1xuICAgICAgICBpdGVtLmNsYXNzTGlzdC5hZGQoJ3NlbGVjdGVkJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCdhcmlhLXNlbGVjdGVkJywgJ3RydWUnKTtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGl0ZW0uY2xhc3NMaXN0LnJlbW92ZSgnc2VsZWN0ZWQnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2FyaWEtc2VsZWN0ZWQnLCAnZmFsc2UnKTtcbiAgICAgIH1cbiAgICB9KTtcblxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEhlYWRlcikge1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIuaW5uZXJIVE1MID0gJyc7XG5cbiAgICAgIGNvbnN0IGhlYWRlckNvbnRlbnQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIGhlYWRlckNvbnRlbnQuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1oZWFkZXItY29udGVudCc7XG5cbiAgICAgIGNvbnN0IGV2ZW50TmFtZSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2gzJyk7XG4gICAgICBldmVudE5hbWUudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC5ldmVudE5hbWU7XG4gICAgICBldmVudE5hbWUuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1uYW1lJztcbiAgICAgIGhlYWRlckNvbnRlbnQuYXBwZW5kQ2hpbGQoZXZlbnROYW1lKTtcblxuICAgICAgaWYgKGV2ZW50UmVjb3JkLnN1bW1hcnkpIHtcbiAgICAgICAgY29uc3Qgc3VtbWFyeSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3AnKTtcbiAgICAgICAgc3VtbWFyeS50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLnN1bW1hcnk7XG4gICAgICAgIHN1bW1hcnkuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1zdW1tYXJ5JztcbiAgICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChzdW1tYXJ5KTtcbiAgICAgIH1cblxuICAgICAgaWYgKGV2ZW50UmVjb3JkLmRlc2NyaXB0aW9uKSB7XG4gICAgICAgIGNvbnN0IGRlc2NyaXB0aW9uID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncCcpO1xuICAgICAgICBkZXNjcmlwdGlvbi50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLmRlc2NyaXB0aW9uO1xuICAgICAgICBkZXNjcmlwdGlvbi5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLWRlc2NyaXB0aW9uJztcbiAgICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChkZXNjcmlwdGlvbik7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IG1ldGEgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIG1ldGEuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1tZXRhJztcblxuICAgICAgY29uc3QgdGltZUVsID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGltZScpO1xuICAgICAgdGltZUVsLnNldEF0dHJpYnV0ZSgnZGF0ZXRpbWUnLCBldmVudFJlY29yZC50aW1lc3RhbXAudG9JU09TdHJpbmcoKSk7XG4gICAgICB0aW1lRWwudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC50aW1lc3RhbXAudG9Mb2NhbGVUaW1lU3RyaW5nKCk7XG4gICAgICB0aW1lRWwuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC10aW1lJztcblxuICAgICAgY29uc3QgZWxlbWVudCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgIGxldCBlbGVtZW50VGV4dCA9IGA8JHtldmVudFJlY29yZC50YWdOYW1lfT5gO1xuICAgICAgaWYgKGV2ZW50UmVjb3JkLmVsZW1lbnRJZCkge1xuICAgICAgICBlbGVtZW50VGV4dCArPSBgIyR7ZXZlbnRSZWNvcmQuZWxlbWVudElkfWA7XG4gICAgICB9XG4gICAgICBlbGVtZW50LnRleHRDb250ZW50ID0gZWxlbWVudFRleHQ7XG4gICAgICBlbGVtZW50LmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtZWxlbWVudCc7XG5cbiAgICAgIG1ldGEuYXBwZW5kQ2hpbGQodGltZUVsKTtcbiAgICAgIG1ldGEuYXBwZW5kQ2hpbGQoZWxlbWVudCk7XG5cbiAgICAgIGhlYWRlckNvbnRlbnQuYXBwZW5kQ2hpbGQobWV0YSk7XG5cbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyLmFwcGVuZENoaWxkKGhlYWRlckNvbnRlbnQpO1xuICAgIH1cblxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEJvZHkpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5pbm5lckhUTUwgPSAnJztcblxuICAgICAgY29uc3QgcHJvcGVydGllc0hlYWRpbmcgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdoNCcpO1xuICAgICAgcHJvcGVydGllc0hlYWRpbmcudGV4dENvbnRlbnQgPSAnUHJvcGVydGllcyc7XG4gICAgICBwcm9wZXJ0aWVzSGVhZGluZy5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMtaGVhZGluZyc7XG5cbiAgICAgIGNvbnN0IHByb3BlcnRpZXNDb250YWluZXIgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIHByb3BlcnRpZXNDb250YWluZXIuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1wcm9wZXJ0aWVzJztcblxuICAgICAgY29uc3QgZXZlbnRQcm9wZXJ0aWVzID0gdGhpcy4jYnVpbGRQcm9wZXJ0aWVzRm9yRGlzcGxheShldmVudFJlY29yZCk7XG4gICAgICBpZiAoT2JqZWN0LmtleXMoZXZlbnRQcm9wZXJ0aWVzKS5sZW5ndGggPiAwKSB7XG4gICAgICAgIHByb3BlcnRpZXNDb250YWluZXIuYXBwZW5kQ2hpbGQodGhpcy4jYnVpbGRQcm9wZXJ0eVRyZWUoZXZlbnRQcm9wZXJ0aWVzKSk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBwcm9wZXJ0aWVzQ29udGFpbmVyLnRleHRDb250ZW50ID0gJ05vIHByb3BlcnRpZXMgdG8gZGlzcGxheSc7XG4gICAgICB9XG5cbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5hcHBlbmRDaGlsZChwcm9wZXJ0aWVzSGVhZGluZyk7XG4gICAgICB0aGlzLiNldmVudERldGFpbEJvZHkuYXBwZW5kQ2hpbGQocHJvcGVydGllc0NvbnRhaW5lcik7XG4gICAgfVxuICB9XG5cbiAgI2J1aWxkUHJvcGVydGllc0ZvckRpc3BsYXkoZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkKTogUmVjb3JkPHN0cmluZywgdW5rbm93bj4ge1xuICAgIGNvbnN0IHByb3BlcnRpZXM6IFJlY29yZDxzdHJpbmcsIHVua25vd24+ID0ge307XG5cbiAgICBpZiAoZXZlbnRSZWNvcmQuY3VzdG9tUHJvcGVydGllcykge1xuICAgICAgT2JqZWN0LmFzc2lnbihwcm9wZXJ0aWVzLCBldmVudFJlY29yZC5jdXN0b21Qcm9wZXJ0aWVzKTtcbiAgICB9XG5cbiAgICBwcm9wZXJ0aWVzLmJ1YmJsZXMgPSBldmVudFJlY29yZC5idWJibGVzO1xuICAgIHByb3BlcnRpZXMuY2FuY2VsYWJsZSA9IGV2ZW50UmVjb3JkLmNhbmNlbGFibGU7XG4gICAgcHJvcGVydGllcy5kZWZhdWx0UHJldmVudGVkID0gZXZlbnRSZWNvcmQuZGVmYXVsdFByZXZlbnRlZDtcbiAgICBwcm9wZXJ0aWVzLmNvbXBvc2VkID0gZXZlbnRSZWNvcmQuY29tcG9zZWQ7XG5cbiAgICBpZiAoZXZlbnRSZWNvcmQubWFuaWZlc3RUeXBlKSB7XG4gICAgICBwcm9wZXJ0aWVzLnR5cGUgPSBldmVudFJlY29yZC5tYW5pZmVzdFR5cGU7XG4gICAgfVxuXG4gICAgcmV0dXJuIHByb3BlcnRpZXM7XG4gIH1cblxuICAjYnVpbGRQcm9wZXJ0eVRyZWUob2JqOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiwgZGVwdGggPSAwKTogSFRNTFVMaXN0RWxlbWVudCB7XG4gICAgY29uc3QgdWwgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd1bCcpO1xuICAgIHVsLmNsYXNzTmFtZSA9ICdldmVudC1wcm9wZXJ0eS10cmVlJztcbiAgICBpZiAoZGVwdGggPiAwKSB7XG4gICAgICB1bC5jbGFzc0xpc3QuYWRkKCduZXN0ZWQnKTtcbiAgICB9XG5cbiAgICBmb3IgKGNvbnN0IFtrZXksIHZhbHVlXSBvZiBPYmplY3QuZW50cmllcyhvYmopKSB7XG4gICAgICBjb25zdCBsaSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2xpJyk7XG4gICAgICBsaS5jbGFzc05hbWUgPSAncHJvcGVydHktaXRlbSc7XG5cbiAgICAgIGNvbnN0IGtleVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICBrZXlTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS1rZXknO1xuICAgICAga2V5U3Bhbi50ZXh0Q29udGVudCA9IGtleTtcblxuICAgICAgY29uc3QgY29sb25TcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgY29sb25TcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS1jb2xvbic7XG4gICAgICBjb2xvblNwYW4udGV4dENvbnRlbnQgPSAnOiAnO1xuXG4gICAgICBsaS5hcHBlbmRDaGlsZChrZXlTcGFuKTtcbiAgICAgIGxpLmFwcGVuZENoaWxkKGNvbG9uU3Bhbik7XG5cbiAgICAgIGlmICh2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBudWxsJztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH0gZWxzZSBpZiAodHlwZW9mIHZhbHVlID09PSAnYm9vbGVhbicpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIGJvb2xlYW4nO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdudW1iZXInKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBudW1iZXInO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdzdHJpbmcnKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBzdHJpbmcnO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBgXCIke3ZhbHVlfVwiYDtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH0gZWxzZSBpZiAoQXJyYXkuaXNBcnJheSh2YWx1ZSkpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIGFycmF5JztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gYEFycmF5KCR7dmFsdWUubGVuZ3RofSlgO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuXG4gICAgICAgIGlmICh2YWx1ZS5sZW5ndGggPiAwICYmIGRlcHRoIDwgMykge1xuICAgICAgICAgIGNvbnN0IG5lc3RlZE9iajogUmVjb3JkPHN0cmluZywgdW5rbm93bj4gPSB7fTtcbiAgICAgICAgICB2YWx1ZS5mb3JFYWNoKChpdGVtLCBpbmRleCkgPT4ge1xuICAgICAgICAgICAgbmVzdGVkT2JqW2luZGV4XSA9IGl0ZW07XG4gICAgICAgICAgfSk7XG4gICAgICAgICAgbGkuYXBwZW5kQ2hpbGQodGhpcy4jYnVpbGRQcm9wZXJ0eVRyZWUobmVzdGVkT2JqLCBkZXB0aCArIDEpKTtcbiAgICAgICAgfVxuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdvYmplY3QnKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBvYmplY3QnO1xuICAgICAgICBjb25zdCBrZXlzID0gT2JqZWN0LmtleXModmFsdWUgYXMgUmVjb3JkPHN0cmluZywgdW5rbm93bj4pO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBrZXlzLmxlbmd0aCA+IDAgPyBgT2JqZWN0YCA6IGB7fWA7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG5cbiAgICAgICAgaWYgKGtleXMubGVuZ3RoID4gMCAmJiBkZXB0aCA8IDMpIHtcbiAgICAgICAgICBsaS5hcHBlbmRDaGlsZCh0aGlzLiNidWlsZFByb3BlcnR5VHJlZSh2YWx1ZSBhcyBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiwgZGVwdGggKyAxKSk7XG4gICAgICAgIH1cbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSc7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9XG5cbiAgICAgIHVsLmFwcGVuZENoaWxkKGxpKTtcbiAgICB9XG5cbiAgICByZXR1cm4gdWw7XG4gIH1cblxuICAjc2Nyb2xsRXZlbnRzVG9Cb3R0b20oKSB7XG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICBjb25zdCBsYXN0RXZlbnQgPSB0aGlzLiNldmVudExpc3QhLmxhc3RFbGVtZW50Q2hpbGQ7XG4gICAgICBpZiAobGFzdEV2ZW50KSB7XG4gICAgICAgIGxhc3RFdmVudC5zY3JvbGxJbnRvVmlldyh7IGJlaGF2aW9yOiAnYXV0bycsIGJsb2NrOiAnZW5kJyB9KTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNpc0V2ZW50c1RhYkFjdGl2ZSgpOiBib29sZWFuIHtcbiAgICBjb25zdCB0YWJzID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdwZi12Ni10YWJzJyk7XG4gICAgaWYgKCF0YWJzKSByZXR1cm4gZmFsc2U7XG5cbiAgICBjb25zdCBzZWxlY3RlZEluZGV4ID0gcGFyc2VJbnQodGFicy5nZXRBdHRyaWJ1dGUoJ3NlbGVjdGVkJykgfHwgJzAnLCAxMCk7XG4gICAgcmV0dXJuIHNlbGVjdGVkSW5kZXggPT09IDM7XG4gIH1cblxuICAjZmlsdGVyRXZlbnRzKHF1ZXJ5OiBzdHJpbmcpIHtcbiAgICB0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgZm9yIChjb25zdCBlbnRyeSBvZiB0aGlzLiNldmVudExpc3QuY2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IGV2ZW50UmVjb3JkID0gdGhpcy4jZ2V0RXZlbnRSZWNvcmRCeUlkKChlbnRyeSBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5ldmVudElkISk7XG5cbiAgICAgIGlmICghZXZlbnRSZWNvcmQpIGNvbnRpbnVlO1xuXG4gICAgICBjb25zdCB0ZXh0TWF0Y2ggPSB0aGlzLiNldmVudE1hdGNoZXNUZXh0RmlsdGVyKGV2ZW50UmVjb3JkKTtcbiAgICAgIGNvbnN0IHR5cGVNYXRjaCA9IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmhhcyhldmVudFJlY29yZC5ldmVudE5hbWUpO1xuICAgICAgY29uc3QgZWxlbWVudE1hdGNoID0gdGhpcy4jZWxlbWVudEZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNlbGVtZW50RmlsdGVycy5oYXMoZXZlbnRSZWNvcmQudGFnTmFtZSk7XG5cbiAgICAgIChlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuID0gISh0ZXh0TWF0Y2ggJiYgdHlwZU1hdGNoICYmIGVsZW1lbnRNYXRjaCk7XG4gICAgfVxuICB9XG5cbiAgI2V2ZW50TWF0Y2hlc1RleHRGaWx0ZXIoZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkKTogYm9vbGVhbiB7XG4gICAgaWYgKCF0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSkgcmV0dXJuIHRydWU7XG5cbiAgICBjb25zdCBzZWFyY2hUZXh0ID0gW1xuICAgICAgZXZlbnRSZWNvcmQudGFnTmFtZSxcbiAgICAgIGV2ZW50UmVjb3JkLmV2ZW50TmFtZSxcbiAgICAgIGV2ZW50UmVjb3JkLmVsZW1lbnRJZCB8fCAnJyxcbiAgICAgIEpTT04uc3RyaW5naWZ5KGV2ZW50UmVjb3JkLmN1c3RvbVByb3BlcnRpZXMgfHwge30pXG4gICAgXS5qb2luKCcgJykudG9Mb3dlckNhc2UoKTtcblxuICAgIHJldHVybiBzZWFyY2hUZXh0LmluY2x1ZGVzKHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKTtcbiAgfVxuXG4gICNnZXRFdmVudFJlY29yZEJ5SWQoaWQ6IHN0cmluZyk6IEV2ZW50UmVjb3JkIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy4jY2FwdHVyZWRFdmVudHMuZmluZChlID0+IGUuaWQgPT09IGlkKTtcbiAgfVxuXG4gICN1cGRhdGVFdmVudEZpbHRlcnMoKSB7XG4gICAgY29uc3Qgc2F2ZWRQcmVmZXJlbmNlcyA9IHRoaXMuI2xvYWRFdmVudEZpbHRlcnNGcm9tU3RvcmFnZSgpO1xuXG4gICAgY29uc3QgZXZlbnRUeXBlRmlsdGVyID0gdGhpcy4jJCgnZXZlbnQtdHlwZS1maWx0ZXInKTtcbiAgICBpZiAoZXZlbnRUeXBlRmlsdGVyICYmIHRoaXMuI2VsZW1lbnRFdmVudE1hcCkge1xuICAgICAgbGV0IG1lbnUgPSBldmVudFR5cGVGaWx0ZXIucXVlcnlTZWxlY3RvcigncGYtdjYtbWVudScpO1xuICAgICAgaWYgKCFtZW51KSB7XG4gICAgICAgIG1lbnUgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwZi12Ni1tZW51Jyk7XG4gICAgICAgIGV2ZW50VHlwZUZpbHRlci5hcHBlbmRDaGlsZChtZW51KTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgZXhpc3RpbmdJdGVtcyA9IG1lbnUucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICBleGlzdGluZ0l0ZW1zLmZvckVhY2goaXRlbSA9PiBpdGVtLnJlbW92ZSgpKTtcblxuICAgICAgY29uc3QgYWxsRXZlbnRUeXBlcyA9IG5ldyBTZXQ8c3RyaW5nPigpO1xuICAgICAgZm9yIChjb25zdCBbdGFnTmFtZSwgZXZlbnRJbmZvXSBvZiB0aGlzLiNlbGVtZW50RXZlbnRNYXApIHtcbiAgICAgICAgaWYgKHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5oYXModGFnTmFtZSkpIHtcbiAgICAgICAgICBmb3IgKGNvbnN0IGV2ZW50TmFtZSBvZiBldmVudEluZm8uZXZlbnROYW1lcykge1xuICAgICAgICAgICAgYWxsRXZlbnRUeXBlcy5hZGQoZXZlbnROYW1lKTtcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgaWYgKHNhdmVkUHJlZmVyZW5jZXMuZXZlbnRUeXBlcykge1xuICAgICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzID0gKHNhdmVkUHJlZmVyZW5jZXMuZXZlbnRUeXBlcyBhcyBTZXQ8c3RyaW5nPiAmIHsgaW50ZXJzZWN0aW9uOiAob3RoZXI6IFNldDxzdHJpbmc+KSA9PiBTZXQ8c3RyaW5nPiB9KS5pbnRlcnNlY3Rpb24oYWxsRXZlbnRUeXBlcyk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzID0gbmV3IFNldChhbGxFdmVudFR5cGVzKTtcbiAgICAgIH1cblxuICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgYWxsRXZlbnRUeXBlcykge1xuICAgICAgICBjb25zdCBpdGVtID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YXJpYW50JywgJ2NoZWNrYm94Jyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YWx1ZScsIGV2ZW50TmFtZSk7XG4gICAgICAgIGlmICh0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmhhcyhldmVudE5hbWUpKSB7XG4gICAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2NoZWNrZWQnLCAnJyk7XG4gICAgICAgIH1cbiAgICAgICAgaXRlbS50ZXh0Q29udGVudCA9IGV2ZW50TmFtZTtcbiAgICAgICAgbWVudS5hcHBlbmRDaGlsZChpdGVtKTtcbiAgICAgIH1cbiAgICB9XG5cbiAgICBjb25zdCBlbGVtZW50RmlsdGVyID0gdGhpcy4jJCgnZWxlbWVudC1maWx0ZXInKTtcbiAgICBpZiAoZWxlbWVudEZpbHRlciAmJiB0aGlzLiNlbGVtZW50RXZlbnRNYXApIHtcbiAgICAgIGxldCBtZW51ID0gZWxlbWVudEZpbHRlci5xdWVyeVNlbGVjdG9yKCdwZi12Ni1tZW51Jyk7XG4gICAgICBpZiAoIW1lbnUpIHtcbiAgICAgICAgbWVudSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3BmLXY2LW1lbnUnKTtcbiAgICAgICAgZWxlbWVudEZpbHRlci5hcHBlbmRDaGlsZChtZW51KTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgZXhpc3RpbmdJdGVtcyA9IG1lbnUucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICBleGlzdGluZ0l0ZW1zLmZvckVhY2goaXRlbSA9PiBpdGVtLnJlbW92ZSgpKTtcblxuICAgICAgY29uc3QgYWxsRWxlbWVudHMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgICAgIGZvciAoY29uc3QgdGFnTmFtZSBvZiB0aGlzLiNlbGVtZW50RXZlbnRNYXAua2V5cygpKSB7XG4gICAgICAgIGlmICh0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuaGFzKHRhZ05hbWUpKSB7XG4gICAgICAgICAgYWxsRWxlbWVudHMuYWRkKHRhZ05hbWUpO1xuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIGlmIChzYXZlZFByZWZlcmVuY2VzLmVsZW1lbnRzKSB7XG4gICAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzID0gKHNhdmVkUHJlZmVyZW5jZXMuZWxlbWVudHMgYXMgU2V0PHN0cmluZz4gJiB7IGludGVyc2VjdGlvbjogKG90aGVyOiBTZXQ8c3RyaW5nPikgPT4gU2V0PHN0cmluZz4gfSkuaW50ZXJzZWN0aW9uKGFsbEVsZW1lbnRzKTtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzID0gbmV3IFNldChhbGxFbGVtZW50cyk7XG4gICAgICB9XG5cbiAgICAgIGZvciAoY29uc3QgdGFnTmFtZSBvZiBhbGxFbGVtZW50cykge1xuICAgICAgICBjb25zdCBpdGVtID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YXJpYW50JywgJ2NoZWNrYm94Jyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YWx1ZScsIHRhZ05hbWUpO1xuICAgICAgICBpZiAodGhpcy4jZWxlbWVudEZpbHRlcnMuaGFzKHRhZ05hbWUpKSB7XG4gICAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2NoZWNrZWQnLCAnJyk7XG4gICAgICAgIH1cbiAgICAgICAgaXRlbS50ZXh0Q29udGVudCA9IGA8JHt0YWdOYW1lfT5gO1xuICAgICAgICBtZW51LmFwcGVuZENoaWxkKGl0ZW0pO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNoYW5kbGVFdmVudFR5cGVGaWx0ZXJDaGFuZ2UgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgeyB2YWx1ZSwgY2hlY2tlZCB9ID0gZXZlbnQgYXMgRXZlbnQgJiB7IHZhbHVlOiBzdHJpbmc7IGNoZWNrZWQ6IGJvb2xlYW4gfTtcblxuICAgIGlmICghdmFsdWUpIHJldHVybjtcblxuICAgIGlmIChjaGVja2VkKSB7XG4gICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuZGVsZXRlKHZhbHVlKTtcbiAgICB9XG5cbiAgICB0aGlzLiNzYXZlRXZlbnRGaWx0ZXJzKCk7XG4gICAgdGhpcy4jZmlsdGVyRXZlbnRzKHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKTtcbiAgfTtcblxuICAjaGFuZGxlRWxlbWVudEZpbHRlckNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB7IHZhbHVlLCBjaGVja2VkIH0gPSBldmVudCBhcyBFdmVudCAmIHsgdmFsdWU6IHN0cmluZzsgY2hlY2tlZDogYm9vbGVhbiB9O1xuXG4gICAgaWYgKCF2YWx1ZSkgcmV0dXJuO1xuXG4gICAgaWYgKGNoZWNrZWQpIHtcbiAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmRlbGV0ZSh2YWx1ZSk7XG4gICAgfVxuXG4gICAgdGhpcy4jc2F2ZUV2ZW50RmlsdGVycygpO1xuICAgIHRoaXMuI2ZpbHRlckV2ZW50cyh0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSk7XG4gIH07XG5cbiAgI2xvYWRFdmVudEZpbHRlcnNGcm9tU3RvcmFnZSgpOiB7IGV2ZW50VHlwZXM6IFNldDxzdHJpbmc+IHwgbnVsbDsgZWxlbWVudHM6IFNldDxzdHJpbmc+IHwgbnVsbCB9IHtcbiAgICBjb25zdCBwcmVmZXJlbmNlczogeyBldmVudFR5cGVzOiBTZXQ8c3RyaW5nPiB8IG51bGw7IGVsZW1lbnRzOiBTZXQ8c3RyaW5nPiB8IG51bGwgfSA9IHtcbiAgICAgIGV2ZW50VHlwZXM6IG51bGwsXG4gICAgICBlbGVtZW50czogbnVsbFxuICAgIH07XG5cbiAgICB0cnkge1xuICAgICAgY29uc3Qgc2F2ZWRFdmVudFR5cGVzID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1ldmVudC10eXBlLWZpbHRlcnMnKTtcbiAgICAgIGlmIChzYXZlZEV2ZW50VHlwZXMpIHtcbiAgICAgICAgcHJlZmVyZW5jZXMuZXZlbnRUeXBlcyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZEV2ZW50VHlwZXMpKTtcbiAgICAgIH1cblxuICAgICAgY29uc3Qgc2F2ZWRFbGVtZW50cyA9IGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtZWxlbWVudC1maWx0ZXJzJyk7XG4gICAgICBpZiAoc2F2ZWRFbGVtZW50cykge1xuICAgICAgICBwcmVmZXJlbmNlcy5lbGVtZW50cyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZEVsZW1lbnRzKSk7XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgY29uc29sZS5kZWJ1ZygnW2NlbS1zZXJ2ZS1jaHJvbWVdIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSBmb3IgZXZlbnQgZmlsdGVycycpO1xuICAgIH1cblxuICAgIHJldHVybiBwcmVmZXJlbmNlcztcbiAgfVxuXG4gICNzYXZlRXZlbnRGaWx0ZXJzKCkge1xuICAgIHRyeSB7XG4gICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLWV2ZW50LXR5cGUtZmlsdGVycycsXG4gICAgICAgIEpTT04uc3RyaW5naWZ5KFsuLi50aGlzLiNldmVudFR5cGVGaWx0ZXJzXSkpO1xuICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oJ2NlbS1zZXJ2ZS1lbGVtZW50LWZpbHRlcnMnLFxuICAgICAgICBKU09OLnN0cmluZ2lmeShbLi4udGhpcy4jZWxlbWVudEZpbHRlcnNdKSk7XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgLy8gbG9jYWxTdG9yYWdlIHVuYXZhaWxhYmxlIChwcml2YXRlIG1vZGUpLCBzaWxlbnRseSBjb250aW51ZVxuICAgIH1cbiAgfVxuXG4gICNjbGVhckV2ZW50cygpIHtcbiAgICB0aGlzLiNjYXB0dXJlZEV2ZW50cyA9IFtdO1xuICAgIHRoaXMuI3NlbGVjdGVkRXZlbnRJZCA9IG51bGw7XG4gICAgaWYgKHRoaXMuI2V2ZW50TGlzdCkge1xuICAgICAgdGhpcy4jZXZlbnRMaXN0LnJlcGxhY2VDaGlsZHJlbigpO1xuICAgIH1cbiAgICBpZiAodGhpcy4jZXZlbnREZXRhaWxIZWFkZXIpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyLmlubmVySFRNTCA9ICcnO1xuICAgIH1cbiAgICBpZiAodGhpcy4jZXZlbnREZXRhaWxCb2R5KSB7XG4gICAgICB0aGlzLiNldmVudERldGFpbEJvZHkuaW5uZXJIVE1MID0gJyc7XG4gICAgfVxuICB9XG5cbiAgYXN5bmMgI2NvcHlFdmVudHMoKSB7XG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIGNvbnN0IHZpc2libGVFdmVudHMgPSBBcnJheS5mcm9tKHRoaXMuI2V2ZW50TGlzdC5jaGlsZHJlbilcbiAgICAgIC5maWx0ZXIoZW50cnkgPT4gIShlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuKVxuICAgICAgLm1hcChlbnRyeSA9PiB7XG4gICAgICAgIGNvbnN0IGlkID0gKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LmV2ZW50SWQhO1xuICAgICAgICByZXR1cm4gdGhpcy4jZ2V0RXZlbnRSZWNvcmRCeUlkKGlkKTtcbiAgICAgIH0pXG4gICAgICAuZmlsdGVyKChldmVudCk6IGV2ZW50IGlzIEV2ZW50UmVjb3JkID0+ICEhZXZlbnQpXG4gICAgICAubWFwKGV2ZW50ID0+IHtcbiAgICAgICAgY29uc3QgdGltZSA9IGV2ZW50LnRpbWVzdGFtcC50b0xvY2FsZVRpbWVTdHJpbmcoKTtcbiAgICAgICAgY29uc3QgZWxlbWVudCA9IGV2ZW50LmVsZW1lbnRJZCA/IGAjJHtldmVudC5lbGVtZW50SWR9YCA6IGV2ZW50LnRhZ05hbWU7XG4gICAgICAgIGNvbnN0IHByb3BzID0gZXZlbnQuY3VzdG9tUHJvcGVydGllcyAmJiBPYmplY3Qua2V5cyhldmVudC5jdXN0b21Qcm9wZXJ0aWVzKS5sZW5ndGggPiAwXG4gICAgICAgICAgPyBgIC0gJHtKU09OLnN0cmluZ2lmeShldmVudC5jdXN0b21Qcm9wZXJ0aWVzKX1gXG4gICAgICAgICAgOiAnJztcbiAgICAgICAgcmV0dXJuIGBbJHt0aW1lfV0gPCR7ZXZlbnQudGFnTmFtZX0+ICR7ZWxlbWVudH0gXFx1MjE5MiAke2V2ZW50LmV2ZW50TmFtZX0ke3Byb3BzfWA7XG4gICAgICB9KVxuICAgICAgLmpvaW4oJ1xcbicpO1xuXG4gICAgaWYgKCF2aXNpYmxlRXZlbnRzKSByZXR1cm47XG5cbiAgICB0cnkge1xuICAgICAgYXdhaXQgbmF2aWdhdG9yLmNsaXBib2FyZC53cml0ZVRleHQodmlzaWJsZUV2ZW50cyk7XG4gICAgICBjb25zdCBidG4gPSB0aGlzLiMkKCdjb3B5LWV2ZW50cycpO1xuICAgICAgaWYgKGJ0bikge1xuICAgICAgICBjb25zdCB0ZXh0Tm9kZSA9IEFycmF5LmZyb20oYnRuLmNoaWxkTm9kZXMpLmZpbmQoXG4gICAgICAgICAgbiA9PiBuLm5vZGVUeXBlID09PSBOb2RlLlRFWFRfTk9ERSAmJiAobi50ZXh0Q29udGVudD8udHJpbSgpLmxlbmd0aCA/PyAwKSA+IDBcbiAgICAgICAgKTtcbiAgICAgICAgaWYgKHRleHROb2RlKSB7XG4gICAgICAgICAgY29uc3Qgb3JpZ2luYWwgPSB0ZXh0Tm9kZS50ZXh0Q29udGVudDtcbiAgICAgICAgICB0ZXh0Tm9kZS50ZXh0Q29udGVudCA9ICdDb3BpZWQhJztcblxuICAgICAgICAgIGlmICh0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgICAgaWYgKHRoaXMuaXNDb25uZWN0ZWQgJiYgdGV4dE5vZGUucGFyZW50Tm9kZSkge1xuICAgICAgICAgICAgICB0ZXh0Tm9kZS50ZXh0Q29udGVudCA9IG9yaWdpbmFsO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgICAgICAgfSwgMjAwMCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9IGNhdGNoIChlcnIpIHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY29weSBldmVudHM6JywgZXJyKTtcbiAgICB9XG4gIH1cblxuICAjc2V0dXBFdmVudExpc3RlbmVycygpIHtcbiAgICB0aGlzLiNldmVudExpc3QgPSB0aGlzLiMkKCdldmVudC1saXN0Jyk7XG4gICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIgPSB0aGlzLiMkKCdldmVudC1kZXRhaWwtaGVhZGVyJyk7XG4gICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5ID0gdGhpcy4jJCgnZXZlbnQtZGV0YWlsLWJvZHknKTtcblxuICAgIGlmICh0aGlzLiNldmVudExpc3QpIHtcbiAgICAgIHRoaXMuI2V2ZW50TGlzdC5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsIChlOiBFdmVudCkgPT4ge1xuICAgICAgICBjb25zdCBsaXN0SXRlbSA9IChlLnRhcmdldCBhcyBFbGVtZW50KS5jbG9zZXN0KCcuZXZlbnQtbGlzdC1pdGVtJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICAgIGlmIChsaXN0SXRlbSkge1xuICAgICAgICAgIGNvbnN0IGV2ZW50SWQgPSBsaXN0SXRlbS5kYXRhc2V0LmV2ZW50SWQ7XG4gICAgICAgICAgaWYgKGV2ZW50SWQpIHtcbiAgICAgICAgICAgIHRoaXMuI3NlbGVjdEV2ZW50KGV2ZW50SWQpO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfSk7XG4gICAgfVxuXG4gICAgY29uc3QgZXZlbnRzRmlsdGVyID0gdGhpcy4jJCgnZXZlbnRzLWZpbHRlcicpO1xuICAgIGlmIChldmVudHNGaWx0ZXIpIHtcbiAgICAgIGV2ZW50c0ZpbHRlci5hZGRFdmVudExpc3RlbmVyKCdpbnB1dCcsIChlOiBFdmVudCkgPT4ge1xuICAgICAgICBjb25zdCB7IHZhbHVlID0gJycgfSA9IGUudGFyZ2V0IGFzIEhUTUxJbnB1dEVsZW1lbnQ7XG4gICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNldmVudHNGaWx0ZXJEZWJvdW5jZVRpbWVyISk7XG4gICAgICAgIHRoaXMuI2V2ZW50c0ZpbHRlckRlYm91bmNlVGltZXIgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICB0aGlzLiNmaWx0ZXJFdmVudHModmFsdWUpO1xuICAgICAgICB9LCAzMDApO1xuICAgICAgfSk7XG4gICAgfVxuXG4gICAgY29uc3QgZXZlbnRUeXBlRmlsdGVyID0gdGhpcy4jJCgnZXZlbnQtdHlwZS1maWx0ZXInKTtcbiAgICBpZiAoZXZlbnRUeXBlRmlsdGVyKSB7XG4gICAgICBldmVudFR5cGVGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlRXZlbnRUeXBlRmlsdGVyQ2hhbmdlIGFzIEV2ZW50TGlzdGVuZXIpO1xuICAgIH1cblxuICAgIGNvbnN0IGVsZW1lbnRGaWx0ZXIgPSB0aGlzLiMkKCdlbGVtZW50LWZpbHRlcicpO1xuICAgIGlmIChlbGVtZW50RmlsdGVyKSB7XG4gICAgICBlbGVtZW50RmlsdGVyLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZUVsZW1lbnRGaWx0ZXJDaGFuZ2UgYXMgRXZlbnRMaXN0ZW5lcik7XG4gICAgfVxuXG4gICAgdGhpcy4jJCgnY2xlYXItZXZlbnRzJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgdGhpcy4jY2xlYXJFdmVudHMoKTtcbiAgICB9KTtcblxuICAgIHRoaXMuIyQoJ2NvcHktZXZlbnRzJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgdGhpcy4jY29weUV2ZW50cygpO1xuICAgIH0pO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcblxuICAgIC8vIENsZWFuIHVwIGtub2IgbGlzdGVuZXJzXG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOmF0dHJpYnV0ZS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuXG4gICAgLy8gQ2xlYW4gdXAgdHJlZSBzdGF0ZSBsaXN0ZW5lcnNcbiAgICBpZiAodGhpcy4jaGFuZGxlVHJlZUV4cGFuZCkge1xuICAgICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdleHBhbmQnLCB0aGlzLiNoYW5kbGVUcmVlRXhwYW5kKTtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSkge1xuICAgICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdjb2xsYXBzZScsIHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSk7XG4gICAgfVxuICAgIGlmICh0aGlzLiNoYW5kbGVUcmVlU2VsZWN0KSB7XG4gICAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QpO1xuICAgIH1cblxuICAgIC8vIENsZWFuIHVwIHdpbmRvdyBsaXN0ZW5lclxuICAgIGlmICh0aGlzLiNoYW5kbGVMb2dzRXZlbnQpIHtcbiAgICAgIHdpbmRvdy5yZW1vdmVFdmVudExpc3RlbmVyKCdjZW06bG9ncycsIHRoaXMuI2hhbmRsZUxvZ3NFdmVudCk7XG4gICAgfVxuXG4gICAgLy8gQ2xlYXIgcGVuZGluZyBmZWVkYmFjayB0aW1lb3V0c1xuICAgIGlmICh0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgIHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICB0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgIH1cbiAgICBpZiAodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgfVxuXG4gICAgLy8gRGlzY29ubmVjdCBtdXRhdGlvbiBvYnNlcnZlclxuICAgIHRoaXMuI29ic2VydmVyLmRpc2Nvbm5lY3QoKTtcblxuICAgIC8vIENsb3NlIFdlYlNvY2tldCBjb25uZWN0aW9uXG4gICAgaWYgKHRoaXMuI3dzQ2xpZW50KSB7XG4gICAgICB0aGlzLiN3c0NsaWVudC5kZXN0cm95KCk7XG4gICAgfVxuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1zZXJ2ZS1jaHJvbWUnOiBDZW1TZXJ2ZUNocm9tZTtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbiAgaGVpZ2h0OiAxMDB2aDtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgLS1wZi12Ni1jLW1hc3RoZWFkX19sb2dvLS1XaWR0aDogbWF4LWNvbnRlbnQ7XFxcXG4gIC0tcGYtdjYtYy1tYXN0aGVhZF9fdG9nZ2xlLS1TaXplOiAxcmVtO1xcXFxufVxcXFxuXFxcXG5baGlkZGVuXSB7XFxcXG4gIGRpc3BsYXk6IG5vbmUgIWltcG9ydGFudDtcXFxcbn1cXFxcblxcXFxuLyogTWFzdGhlYWQgbG9nbyBzdHlsZXMgKi9cXFxcbi5tYXN0aGVhZC1sb2dvIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuICBjb2xvcjogaW5oZXJpdDtcXFxcbiAgbWF4LWhlaWdodDogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZF9fbG9nby0tTWF4SGVpZ2h0KTtcXFxcbiAgZ2FwOiA0cHg7XFxcXG4gIFxcXFx1MDAyNiBpbWcge1xcXFxuICAgIGRpc3BsYXk6IGJsb2NrO1xcXFxuICAgIG1heC1oZWlnaHQ6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX2xvZ28tLU1heEhlaWdodCk7XFxcXG4gICAgd2lkdGg6IGF1dG87XFxcXG4gIH1cXFxcbiAgXFxcXHUwMDI2IDo6c2xvdHRlZChbc2xvdD1cXFxcXFxcInRpdGxlXFxcXFxcXCJdKSB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICAgIGZvbnQtc2l6ZTogMS4xMjVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIH1cXFxcbiAgXFxcXHUwMDI2IGgxIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gICAgZm9udC1zaXplOiAxOHB4O1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIFRvb2xiYXIgZ3JvdXAgYWxpZ25tZW50ICovXFxcXG5wZi12Ni10b29sYmFyLWdyb3VwW3ZhcmlhbnQ9XFxcXFxcXCJhY3Rpb24tZ3JvdXBcXFxcXFxcIl0ge1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiBhdXRvO1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxufVxcXFxuXFxcXG4uZGVidWctcGFuZWwge1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyOiAxcHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICBwYWRkaW5nOiAxLjVyZW07XFxcXG4gIG1heC13aWR0aDogNjAwcHg7XFxcXG4gIHdpZHRoOiA5MCU7XFxcXG4gIG1heC1oZWlnaHQ6IDgwdmg7XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG5cXFxcbiAgaDIge1xcXFxuICAgIG1hcmdpbjogMCAwIDFyZW0gMDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgZm9udC1zaXplOiAxLjEyNXJlbTtcXFxcbiAgICBmb250LXdlaWdodDogNjAwO1xcXFxuICB9XFxcXG5cXFxcbiAgZGwge1xcXFxuICAgIG1hcmdpbjogMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGR0IHtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIG1hcmdpbi10b3A6IDAuNXJlbTtcXFxcbiAgICBmb250LXdlaWdodDogNTAwO1xcXFxuICB9XFxcXG5cXFxcbiAgZGQge1xcXFxuICAgIG1hcmdpbjogMCAwIDAuNXJlbSAwO1xcXFxuICAgIGZvbnQtZmFtaWx5OiB1aS1tb25vc3BhY2UsICdDYXNjYWRpYSBDb2RlJywgJ1NvdXJjZSBDb2RlIFBybycsIE1lbmxvLCBDb25zb2xhcywgJ0RlamFWdSBTYW5zIE1vbm8nLCBtb25vc3BhY2U7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIH1cXFxcblxcXFxuICBkZXRhaWxzIHtcXFxcbiAgICBtYXJnaW4tdG9wOiAxcmVtO1xcXFxuXFxcXG4gICAgc3VtbWFyeSB7XFxcXG4gICAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgICAgZm9udC13ZWlnaHQ6IDUwMDtcXFxcbiAgICAgIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gICAgICBkaXNwbGF5OiBmbGV4O1xcXFxuICAgICAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgICAgIGdhcDogMC41cmVtO1xcXFxuICAgICAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG5cXFxcbiAgICAgIFxcXFx1MDAyNjo6LXdlYmtpdC1kZXRhaWxzLW1hcmtlciB7XFxcXG4gICAgICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gICAgICB9XFxcXG5cXFxcbiAgICAgIFxcXFx1MDAyNjo6YmVmb3JlIHtcXFxcbiAgICAgICAgY29udGVudDogJ1xcXFxcXFxcMjVCOCc7XFxcXG4gICAgICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgICAgICAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIDEwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgICAgfVxcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNltvcGVuXSBzdW1tYXJ5OjpiZWZvcmUge1xcXFxuICAgICAgdHJhbnNmb3JtOiByb3RhdGUoOTBkZWcpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIHByZSB7XFxcXG4gICAgICBtYXJnaW4tdG9wOiAwLjVyZW07XFxcXG4gICAgICBtYXJnaW4tbGVmdDogMS41cmVtO1xcXFxuICAgICAgcGFkZGluZzogMC41cmVtO1xcXFxuICAgICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0KTtcXFxcbiAgICAgIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgICBvdmVyZmxvdy14OiBhdXRvO1xcXFxuICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIC5idXR0b24tY29udGFpbmVyIHtcXFxcbiAgICBkaXNwbGF5OiBmbGV4O1xcXFxuICAgIGdhcDogMC41cmVtO1xcXFxuICAgIG1hcmdpbi10b3A6IDFyZW07XFxcXG4gIH1cXFxcblxcXFxuICBwIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICB9XFxcXG5cXFxcbiAgYnV0dG9uIHtcXFxcbiAgICBtYXJnaW4tdG9wOiAxcmVtO1xcXFxuICAgIHBhZGRpbmc6IDAuNXJlbSAxcmVtO1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tYnJhbmQpO1xcXFxuICAgIGJvcmRlcjogbm9uZTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBmb250LXdlaWdodDogNDAwO1xcXFxuICAgIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgICB0cmFuc2l0aW9uOiBhbGwgMjAwbXMgY3ViaWMtYmV6aWVyKDAuNjQ1LCAwLjA0NSwgMC4zNTUsIDEpO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1ob3Zlcik7XFxcXG4gICAgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIENvbnRlbnQgYXJlYSBwYWRkaW5nIGZvciBkZW1vICovXFxcXG5wZi12Ni1wYWdlLW1haW4ge1xcXFxuICBtaW4taGVpZ2h0OiBjYWxjKDEwMGR2aCAtIDcycHggLSB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0taW5zZXQtLXBhZ2UtY2hyb21lKSk7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIFxcXFx1MDAyNiBcXFxcdTAwM2UgOjpzbG90dGVkKDpub3QoW3Nsb3Q9a25vYnNdKSkge1xcXFxuICAgIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gICAgZmxleDogMTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG5jZW0tZHJhd2VyIHtcXFxcbiAgcGYtdjYtdGFicyB7XFxcXG4gICAgcGYtdjYtdGFiIHtcXFxcbiAgICAgIHBhZGRpbmctYmxvY2stZW5kOiAwO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBFbGVtZW50IGRlc2NyaXB0aW9ucyBpbiBsaXN0aW5nICovXFxcXG4uZWxlbWVudC1zdW1tYXJ5IHtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtKTtcXFxcbn1cXFxcblxcXFxuLmVsZW1lbnQtZGVzY3JpcHRpb24ge1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxufVxcXFxuXFxcXG4vKiBDYXJkIGZvb3RlciBkZW1vIG5hdmlnYXRpb24gKi9cXFxcbi5jYXJkLWRlbW9zIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC13cmFwOiB3cmFwO1xcXFxuICBnYXA6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1nYXAtLWFjdGlvbi10by1hY3Rpb24tLWRlZmF1bHQpO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBtYXJnaW46IDA7XFxcXG59XFxcXG5cXFxcbi5wYWNrYWdlLW5hbWUge1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtKTtcXFxcbn1cXFxcblxcXFxuLyogS25vYnMgY29udGFpbmVyIC0gZmlsbHMgdGFiIHBhbmVsIGhlaWdodCAqL1xcXFxuI2tub2JzLWNvbnRhaW5lciB7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgXFxcXHUwMDI2IDo6c2xvdHRlZChbc2xvdD1cXFxcXFxcImtub2JzXFxcXFxcXCJdKSB7XFxcXG4gICAgZmxleDogMTtcXFxcbiAgICBtaW4taGVpZ2h0OiAwO1xcXFxuICAgIG92ZXJmbG93OiBoaWRkZW47XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmtub2JzLWVtcHR5IHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIHRleHQtYWxpZ246IGNlbnRlcjtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG5cXFxcbiAgY29kZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICAgIHBhZGRpbmc6IDJweCA2cHg7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogM3B4O1xcXFxuICAgIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4uaW5zdGFuY2UtdGFnIHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4uaW5zdGFuY2UtbGFiZWwge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWdyb3VwIHtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2Omxhc3QtY2hpbGQge1xcXFxuICAgIG1hcmdpbi1ib3R0b206IDA7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogUGF0dGVybkZseSB2NiBmb3JtIC0gaG9yaXpvbnRhbCBsYXlvdXQgKi9cXFxcbnBmLXY2LWZvcm1baG9yaXpvbnRhbF0gcGYtdjYtZm9ybS1maWVsZC1ncm91cCB7XFxcXG4gIGdyaWQtY29sdW1uOiBzcGFuIDJcXFxcbn1cXFxcblxcXFxuLmtub2ItZ3JvdXAtdGl0bGUge1xcXFxuICBncmlkLWNvbHVtbjogMSAvIC0xO1xcXFxuICBtYXJnaW46IDAgMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtYmFzZSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIGJvcmRlci1ib3R0b206IDFweCBzb2xpZCB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItY29sb3IpO1xcXFxuICBwYWRkaW5nLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWNvbnRyb2wge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbn1cXFxcblxcXFxuLmtub2ItbGFiZWwge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG59XFxcXG5cXFxcbi5rbm9iLW5hbWUge1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWRlc2NyaXB0aW9uIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxuXFxcXG4gIHAge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiAzcHg7XFxcXG4gICAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICB9XFxcXG5cXFxcbiAgYSB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIHRleHQtZGVjb3JhdGlvbjogdW5kZXJsaW5lO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIHN0cm9uZyB7XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgfVxcXFxuXFxcXG4gIHVsLCBvbCB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSAwO1xcXFxuICAgIHBhZGRpbmctbGVmdDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuZm9vdGVyLnBmLW0tc3RpY2t5LWJvdHRvbSB7XFxcXG4gIHZpZXctdHJhbnNpdGlvbi1uYW1lOiBkZXYtc2VydmVyLWZvb3RlcjtcXFxcbiAgcG9zaXRpb246IHN0aWNreTtcXFxcbiAgYm90dG9tOiAwO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyLXRvcDogMXB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIHotaW5kZXg6IHZhcigtLXBmLXY2LWMtcGFnZS0tc2VjdGlvbi0tbS1zdGlja3ktYm90dG9tLS1aSW5kZXgsIHZhcigtLXBmLXQtLWdsb2JhbC0tei1pbmRleC0tbWQpKTtcXFxcbiAgYm94LXNoYWRvdzogdmFyKC0tcGYtdjYtYy1wYWdlLS1zZWN0aW9uLS1tLXN0aWNreS1ib3R0b20tLUJveFNoYWRvdywgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3gtc2hhZG93LS1zbS0tdG9wKSk7XFxcXG59XFxcXG5cXFxcbi5mb290ZXItZGVzY3JpcHRpb24ge1xcXFxuICBwYWRkaW5nOiAxLjVyZW07XFxcXG5cXFxcbiAgXFxcXHUwMDI2LmVtcHR5IHtcXFxcbiAgICBkaXNwbGF5OiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbmZvb3RlciA6OnNsb3R0ZWQoW3Nsb3Q9XFxcXFxcXCJkZXNjcmlwdGlvblxcXFxcXFwiXSkge1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBsaW5lLWhlaWdodDogMS42O1xcXFxuICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWhvdmVyKTtcXFxcbiAgICBwYWRkaW5nOiAycHggNnB4O1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDNweDtcXFxcbiAgICBmb250LWZhbWlseTogdWktbW9ub3NwYWNlLCAnQ2FzY2FkaWEgQ29kZScsICdTb3VyY2UgQ29kZSBQcm8nLCBNZW5sbywgQ29uc29sYXMsICdEZWphVnUgU2FucyBNb25vJywgbW9ub3NwYWNlO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmxvZ3Mtd3JhcHBlciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxuI2xvZy1jb250YWluZXIge1xcXFxuICBmbGV4LWdyb3c6IDE7XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG59XFxcXG5cXFxcbi5sb2ctZW50cnkge1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgYWxpZ24taXRlbXM6IGJhc2VsaW5lO1xcXFxuICBwZi12Ni1sYWJlbCB7XFxcXG4gICAgZmxleC1zaHJpbms6IDA7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmxvZy10aW1lLFxcXFxuLmxvZy1tZXNzYWdlIHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5sb2ctdGltZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZmxleC1zaHJpbms6IDA7XFxcXG4gIGZvbnQtc2l6ZTogMTFweDtcXFxcbn1cXFxcblxcXFxuLmxvZy1tZXNzYWdlIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIHdvcmQtYnJlYWs6IGJyZWFrLXdvcmQ7XFxcXG59XFxcXG5cXFxcbi8qIE5hdmlnYXRpb24gY29udGVudCAobGlnaHQgRE9NIHNsb3R0ZWQgY29udGVudCBmb3IgcGYtdjYtcGFnZS1zaWRlYmFyKSAqL1xcXFxuLm5hdi1wYWNrYWdlIHtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2IFxcXFx1MDAzZSBzdW1tYXJ5IHtcXFxcbiAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gICAgcGFkZGluZzogMC41cmVtIDFyZW07XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICAgIHRyYW5zaXRpb246IGJhY2tncm91bmQgMjAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICBtYXJnaW4tYm90dG9tOiAwLjVyZW07XFxcXG4gICAgZGlzcGxheTogZmxleDtcXFxcbiAgICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICAgIGdhcDogMC41cmVtO1xcXFxuICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0taG92ZXIpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNjo6LXdlYmtpdC1kZXRhaWxzLW1hcmtlciB7XFxcXG4gICAgICBkaXNwbGF5OiBub25lO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNjo6YmVmb3JlIHtcXFxcbiAgICAgIGNvbnRlbnQ6ICdcXFxcXFxcXDI1QjgnO1xcXFxuICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICAgICAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIDEwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjZbb3Blbl0gXFxcXHUwMDNlIHN1bW1hcnk6OmJlZm9yZSB7XFxcXG4gICAgdHJhbnNmb3JtOiByb3RhdGUoOTBkZWcpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5uYXYtZWxlbWVudCB7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcblxcXFxuICBzdW1tYXJ5IHtcXFxcbiAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gICAgcGFkZGluZzogMC41cmVtIDFyZW07XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIGZvbnQtZmFtaWx5OiB1aS1tb25vc3BhY2UsICdDYXNjYWRpYSBDb2RlJywgJ1NvdXJjZSBDb2RlIFBybycsIE1lbmxvLCBDb25zb2xhcywgJ0RlamFWdSBTYW5zIE1vbm8nLCBtb25vc3BhY2U7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICAgIHRyYW5zaXRpb246IGJhY2tncm91bmQgMjAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICBkaXNwbGF5OiBmbGV4O1xcXFxuICAgIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gICAgZ2FwOiAwLjVyZW07XFxcXG4gICAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1ob3Zlcik7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2Ojotd2Via2l0LWRldGFpbHMtbWFya2VyIHtcXFxcbiAgICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2OjpiZWZvcmUge1xcXFxuICAgICAgY29udGVudDogJ1xcXFxcXFxcMjVCOCc7XFxcXG4gICAgICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XFxcXG4gICAgICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gMTAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNltvcGVuXSBzdW1tYXJ5OjpiZWZvcmUge1xcXFxuICAgIHRyYW5zZm9ybTogcm90YXRlKDkwZGVnKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubmF2LWVsZW1lbnQtdGl0bGUge1xcXFxuICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcbn1cXFxcblxcXFxuLm5hdi1kZW1vLWxpc3Qge1xcXFxuICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIDAgMCAwO1xcXFxuICBkaXNwbGF5OiBncmlkO1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpO1xcXFxufVxcXFxuXFxcXG4ubmF2LWRlbW8tbGluayB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxuICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogY2FsYyh2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSAqIDIpO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1yYWRpdXMpO1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kIDAuMnMsIGNvbG9yIDAuMnM7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1icmFuZCk7XFxcXG5cXFxcbiAgICAubmF2LXBhY2thZ2UtbmFtZSB7XFxcXG4gICAgICBjb2xvcjogcmdiYSgyNTUsIDI1NSwgMjU1LCAwLjgpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNlthcmlhLWN1cnJlbnQ9XFxcXFxcXCJwYWdlXFxcXFxcXCJdIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1icmFuZCk7XFxcXG5cXFxcbiAgICAubmF2LXBhY2thZ2UtbmFtZSB7XFxcXG4gICAgICBjb2xvcjogcmdiYSgyNTUsIDI1NSwgMjU1LCAwLjgpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubmF2LXBhY2thZ2UtbmFtZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLyogSW5mbyBidXR0b24gcG9wb3ZlciB0cmlnZ2VycyBpbiBrbm9icyAtIG92ZXJyaWRlIHBsYWluIGJ1dHRvbiBwYWRkaW5nICovXFxcXG5wZi12Ni1wb3BvdmVyIHBmLXY2LWJ1dHRvblt2YXJpYW50PVxcXFxcXFwicGxhaW5cXFxcXFxcIl0ge1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1tLXBsYWluLS1QYWRkaW5nSW5saW5lRW5kOiAwO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1tLXBsYWluLS1QYWRkaW5nSW5saW5lU3RhcnQ6IDA7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLU1pbldpZHRoOiBhdXRvO1xcXFxufVxcXFxuXFxcXG4vKiBLbm9iIGRlc2NyaXB0aW9uIGNvbnRlbnQgKHNsb3R0ZWQgaW4gZm9ybSBncm91cCBoZWxwZXIpICovXFxcXG5wZi12Ni1mb3JtLWdyb3VwIFtzbG90PVxcXFxcXFwiaGVscGVyXFxcXFxcXCJdIHtcXFxcbiAgcCB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSAwO1xcXFxuICB9XFxcXG5cXFxcbiAgY29kZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDNweDtcXFxcbiAgICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIH1cXFxcblxcXFxuICBhIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgdGV4dC1kZWNvcmF0aW9uOiB1bmRlcmxpbmU7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgc3Ryb25nIHtcXFxcbiAgICBmb250LXdlaWdodDogNjAwO1xcXFxuICB9XFxcXG5cXFxcbiAgdWwsIG9sIHtcXFxcbiAgICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIDA7XFxcXG4gICAgcGFkZGluZy1sZWZ0OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLWxnKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBTeW50YXggaGlnaGxpZ2h0aW5nIChjaHJvbWEgLSB0aGVtYWJsZSB2aWEgQ1NTIGN1c3RvbSBwcm9wZXJ0aWVzKSAqL1xcXFxucGYtdjYtZm9ybS1ncm91cCBbc2xvdD1cXFxcXFxcImhlbHBlclxcXFxcXFwiXSB7XFxcXG4gIC5jaHJvbWEge1xcXFxuICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItcmFkaXVzKTtcXFxcbiAgICBvdmVyZmxvdy14OiBhdXRvO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2IC5sbnRkIHsgdmVydGljYWwtYWxpZ246IHRvcDsgcGFkZGluZzogMDsgbWFyZ2luOiAwOyBib3JkZXI6IDA7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmxudGFibGUgeyBib3JkZXItc3BhY2luZzogMDsgcGFkZGluZzogMDsgbWFyZ2luOiAwOyBib3JkZXI6IDA7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmhsIHsgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWhpZ2hsaWdodCkgfVxcXFxuICAgIFxcXFx1MDAyNiAubG50LFxcXFxuICAgIFxcXFx1MDAyNiAubG4ge1xcXFxuICAgICAgd2hpdGUtc3BhY2U6IHByZTtcXFxcbiAgICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuICAgICAgbWFyZ2luLXJpZ2h0OiAwLjRlbTtcXFxcbiAgICAgIHBhZGRpbmc6IDAgMC40ZW0gMCAwLjRlbTtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5saW5lIHsgZGlzcGxheTogZmxleDsgfVxcXFxuXFxcXG4gICAgLyogS2V5d29yZHMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmssXFxcXG4gICAgXFxcXHUwMDI2IC5rYyxcXFxcbiAgICBcXFxcdTAwMjYgLmtkLFxcXFxuICAgIFxcXFx1MDAyNiAua24sXFxcXG4gICAgXFxcXHUwMDI2IC5rcCxcXFxcbiAgICBcXFxcdTAwMjYgLmtyIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgta2V5d29yZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5rdCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtdHlwZSk7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG5cXFxcbiAgICAvKiBOYW1lcyAqL1xcXFxuICAgIFxcXFx1MDAyNiAubmEsXFxcXG4gICAgXFxcXHUwMDI2IC5uYixcXFxcbiAgICBcXFxcdTAwMjYgLm5vLFxcXFxuICAgIFxcXFx1MDAyNiAubnYsXFxcXG4gICAgXFxcXHUwMDI2IC52YyxcXFxcbiAgICBcXFxcdTAwMjYgLnZnLFxcXFxuICAgIFxcXFx1MDAyNiAudmkge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1uYW1lKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5icCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAubmMgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWNsYXNzKTsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLm5kIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1kZWNvcmF0b3IpOyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuICAgIFxcXFx1MDAyNiAubmksXFxcXG4gICAgXFxcXHUwMDI2IC5zcyB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXNwZWNpYWwpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLm5lLFxcXFxuICAgIFxcXFx1MDAyNiAubmwge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1rZXl3b3JkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLm5mIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1mdW5jdGlvbik7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5ubiB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAubnQgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXRhZykgfVxcXFxuXFxcXG4gICAgLyogU3RyaW5ncyAqL1xcXFxuICAgIFxcXFx1MDAyNiAucyxcXFxcbiAgICBcXFxcdTAwMjYgLnNhLFxcXFxuICAgIFxcXFx1MDAyNiAuc2IsXFxcXG4gICAgXFxcXHUwMDI2IC5zYyxcXFxcbiAgICBcXFxcdTAwMjYgLmRsLFxcXFxuICAgIFxcXFx1MDAyNiAuc2QsXFxcXG4gICAgXFxcXHUwMDI2IC5zMixcXFxcbiAgICBcXFxcdTAwMjYgLnNlLFxcXFxuICAgIFxcXFx1MDAyNiAuc2gsXFxcXG4gICAgXFxcXHUwMDI2IC5zaSxcXFxcbiAgICBcXFxcdTAwMjYgLnN4LFxcXFxuICAgIFxcXFx1MDAyNiAuczEge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1zdHJpbmcpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLnNyIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC10YWcpIH1cXFxcblxcXFxuICAgIC8qIE51bWJlcnMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLm0sXFxcXG4gICAgXFxcXHUwMDI2IC5tYixcXFxcbiAgICBcXFxcdTAwMjYgLm1mLFxcXFxuICAgIFxcXFx1MDAyNiAubWgsXFxcXG4gICAgXFxcXHUwMDI2IC5taSxcXFxcbiAgICBcXFxcdTAwMjYgLmlsLFxcXFxuICAgIFxcXFx1MDAyNiAubW8ge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1udW1iZXIpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIE9wZXJhdG9ycyAqL1xcXFxuICAgIFxcXFx1MDAyNiAubyxcXFxcbiAgICBcXFxcdTAwMjYgLm93IHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgta2V5d29yZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBDb21tZW50cyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuYyxcXFxcbiAgICBcXFxcdTAwMjYgLmNoLFxcXFxuICAgIFxcXFx1MDAyNiAuY20sXFxcXG4gICAgXFxcXHUwMDI2IC5jMSB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gICAgICBmb250LXN0eWxlOiBpdGFsaWM7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuY3MsXFxcXG4gICAgXFxcXHUwMDI2IC5jcCxcXFxcbiAgICBcXFxcdTAwMjYgLmNwZiB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBFcnJvcnMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmVyciB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWVycm9yKTtcXFxcbiAgICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvci1iZyk7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgLyogR2VuZXJpY3MgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmdkIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZGVsZXRlZCk7XFxcXG4gICAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZGVsZXRlZC1iZyk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2UgeyBmb250LXN0eWxlOiBpdGFsaWM7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmdyIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvcikgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2ggeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdpIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtaW5zZXJ0ZWQpO1xcXFxuICAgICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWluc2VydGVkLWJnKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5nbyB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3AgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdzIHsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmd1IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5ndCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZXJyb3IpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdsIHsgdGV4dC1kZWNvcmF0aW9uOiB1bmRlcmxpbmU7IH1cXFxcbiAgICBcXFxcdTAwMjYgLncgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCkgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIEV2ZW50cyB0YWIgc3R5bGluZyAtIFByaW1hcnktZGV0YWlsIGxheW91dCAqL1xcXFxuLmV2ZW50cy13cmFwcGVyIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4jZXZlbnQtZHJhd2VyIHtcXFxcbiAgZmxleDogMTtcXFxcbiAgbWluLWhlaWdodDogMDtcXFxcbn1cXFxcblxcXFxuLyogRXZlbnQgbGlzdCAocHJpbWFyeSBwYW5lbCkgKi9cXFxcbiNldmVudC1saXN0IHtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtbGlzdC1pdGVtIHtcXFxcbiAgLyogUmVzZXQgYnV0dG9uIHN0eWxlcyAqL1xcXFxuICBhcHBlYXJhbmNlOiBub25lO1xcXFxuICBiYWNrZ3JvdW5kOiBub25lO1xcXFxuICBib3JkZXI6IG5vbmU7XFxcXG4gIGJvcmRlci1sZWZ0OiAzcHggc29saWQgdHJhbnNwYXJlbnQ7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgZm9udDogaW5oZXJpdDtcXFxcbiAgY29sb3I6IGluaGVyaXQ7XFxcXG4gIHRleHQtYWxpZ246IGluaGVyaXQ7XFxcXG4gIHdpZHRoOiAxMDAlO1xcXFxuXFxcXG4gIC8qIENvbXBvbmVudCBzdHlsZXMgKi9cXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAxMDBtcyBlYXNlLWluLW91dCwgYm9yZGVyLWNvbG9yIDEwMG1zIGVhc2UtaW4tb3V0O1xcXFxuXFxcXG4gIHBmLXY2LWxhYmVsIHtcXFxcbiAgICBmbGV4LXNocmluazogMDtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0taG92ZXIpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmZvY3VzIHtcXFxcbiAgICBvdXRsaW5lOiAycHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1jbGlja2VkKTtcXFxcbiAgICBvdXRsaW5lLW9mZnNldDogLTJweDtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5zZWxlY3RlZCB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0tc2VsZWN0ZWQpO1xcXFxuICAgIGJvcmRlci1sZWZ0LWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4uZXZlbnQtdGltZSxcXFxcbi5ldmVudC1lbGVtZW50IHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC10aW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmbGV4LXNocmluazogMDtcXFxcbiAgZm9udC1zaXplOiAxMXB4O1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZWxlbWVudCB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDQwMDtcXFxcbn1cXFxcblxcXFxuLyogRXZlbnQgZGV0YWlsIHBhbmVsICovXFxcXG4uZXZlbnQtZGV0YWlsLWhlYWRlci1jb250ZW50IHtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGJvcmRlci1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci13aWR0aCkgc29saWQgdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLWNvbG9yKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1uYW1lIHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtbGcpO1xcXFxuICBmb250LXdlaWdodDogNjAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1zdW1tYXJ5IHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxuICB3aGl0ZS1zcGFjZTogcHJlLXdyYXA7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtZGVzY3JpcHRpb24ge1xcXFxuICBtYXJnaW46IDAgMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSAwO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjU7XFxcXG4gIHdoaXRlLXNwYWNlOiBwcmUtd3JhcDtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1tZXRhIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLXRpbWUge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1lbGVtZW50IHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMtaGVhZGluZyB7XFxcXG4gIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLWJhc2UpO1xcXFxuICBmb250LXdlaWdodDogNjAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1wcm9wZXJ0aWVzIHtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXNlY29uZGFyeSk7XFxcXG4gIGJvcmRlcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXdpZHRoKSBzb2xpZCB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItY29sb3IpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItcmFkaXVzKTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IDEycHg7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjY7XFxcXG4gIG1hcmdpbjogMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LXByb3BlcnR5LXRyZWUge1xcXFxuICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBtYXJnaW46IDA7XFxcXG5cXFxcbiAgXFxcXHUwMDI2Lm5lc3RlZCB7XFxcXG4gICAgcGFkZGluZy1sZWZ0OiAxLjVlbTtcXFxcbiAgICBtYXJnaW4tdG9wOiAwLjI1ZW07XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LWl0ZW0ge1xcXFxuICBwYWRkaW5nOiAwLjEyNWVtIDA7XFxcXG59XFxcXG5cXFxcbi5wcm9wZXJ0eS1rZXkge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDUwMDtcXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LWNvbG9uIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxufVxcXFxuXFxcXG4ucHJvcGVydHktdmFsdWUge1xcXFxuICBcXFxcdTAwMjYubnVsbCxcXFxcbiAgXFxcXHUwMDI2LnVuZGVmaW5lZCB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5ib29sZWFuIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItY29sb3ItYm9vbGVhbik7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYubnVtYmVyIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItY29sb3ItbnVtYmVyKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5zdHJpbmcge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1jb2xvci1zdHJpbmcpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2LmFycmF5LFxcXFxuICBcXFxcdTAwMjYub2JqZWN0IHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNkZWJ1Zy1tb2RhbCB7XFxcXG4gIGNvbnRhaW5lci10eXBl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QTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsZ0RBc0hkLFVBQVUsS0FBSyxjQUFjLENBQUM7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxnRUFrQ2QsS0FBSyxrQkFBa0IsR0FBRztBQUFBO0FBQUE7QUFBQTtBQUFBLGdFQUkxQixLQUFLLGFBQWEsR0FBRztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQTBDbkY7QUFBQSxFQXNDQSxNQUFNLG9CQUFvQjtBQUt4QixRQUFJLENBQUMsbUJBQUssdUJBQXNCO0FBQzlCLE9BQUMsRUFBRSxnQkFBZ0IsR0FBRyxFQUFFLGlCQUFpQixDQUFDLElBQUksTUFBTSxRQUFRLElBQUk7QUFBQTtBQUFBLFFBRTlELE9BQU8sNEJBQTRCO0FBQUE7QUFBQSxRQUVuQyxPQUFPLDZCQUE2QjtBQUFBLE1BQ3RDLENBQUM7QUFFRCxhQUFPLHlCQUF5QixFQUFFLE1BQU0sQ0FBQyxNQUN2QyxRQUFRLE1BQU0sNkNBQTZDLENBQUMsQ0FBQztBQUMvRCx5QkFBSyxzQkFBdUI7QUFBQSxJQUM5QjtBQUVBLFVBQU0sa0JBQWtCO0FBR3hCLFFBQUksbUJBQUssY0FBYSxNQUFNO0FBQzFCLDRCQUFLLDRDQUFMO0FBQUEsSUFDRjtBQUNBLDBCQUFLLCtEQUFMO0FBQUEsRUFDRjtBQUFBLEVBRUEsZUFBZTtBQUViLDBCQUFLLGlEQUFMO0FBR0EsMEJBQUssc0RBQUw7QUFHQSwwQkFBSyxpREFBTDtBQUdBLDBCQUFLLGdEQUFMO0FBR0EsMEJBQUsscURBQUw7QUFHQSwwQkFBSyx5REFBTDtBQUdBLDBCQUFLLDREQUFMO0FBR0EsMEJBQUssaURBQUwsV0FBMEIsS0FBSyxNQUFNO0FBQ25DLDRCQUFLLG1EQUFMO0FBQUEsSUFDRixDQUFDO0FBSUQsMEJBQUssaUNBQUwsV0FBUSxrQkFBa0IsaUJBQWlCLFNBQVMsTUFBTTtBQUN4RCxhQUFPLFNBQVMsT0FBTztBQUFBLElBQ3pCLENBQUM7QUFHRCwwQkFBSyxpQ0FBTCxXQUFRLGlCQUFpQixpQkFBaUIsU0FBUyxNQUFNO0FBQ3ZELE1BQUMsc0JBQUssaUNBQUwsV0FBUSx1QkFBK0IsTUFBTTtBQUM5Qyx5QkFBSyxXQUFVLE1BQU07QUFBQSxJQUN2QixDQUFDO0FBR0QsdUJBQUssV0FBVSxLQUFLO0FBQUEsRUFDdEI7QUFBQSxFQTYrQ0EsdUJBQXVCO0FBQ3JCLFVBQU0scUJBQXFCO0FBRzNCLFNBQUssb0JBQW9CLHlCQUF5QixtQkFBSyxjQUFhO0FBQ3BFLFNBQUssb0JBQW9CLHdCQUF3QixtQkFBSyxjQUFhO0FBQ25FLFNBQUssb0JBQW9CLDRCQUE0QixtQkFBSyxjQUFhO0FBQ3ZFLFNBQUssb0JBQW9CLHdCQUF3QixtQkFBSyxhQUFZO0FBQ2xFLFNBQUssb0JBQW9CLHVCQUF1QixtQkFBSyxhQUFZO0FBQ2pFLFNBQUssb0JBQW9CLDJCQUEyQixtQkFBSyxhQUFZO0FBR3JFLFFBQUksbUJBQUssb0JBQW1CO0FBQzFCLFdBQUssb0JBQW9CLFVBQVUsbUJBQUssa0JBQWlCO0FBQUEsSUFDM0Q7QUFDQSxRQUFJLG1CQUFLLHNCQUFxQjtBQUM1QixXQUFLLG9CQUFvQixZQUFZLG1CQUFLLG9CQUFtQjtBQUFBLElBQy9EO0FBQ0EsUUFBSSxtQkFBSyxvQkFBbUI7QUFDMUIsV0FBSyxvQkFBb0IsVUFBVSxtQkFBSyxrQkFBaUI7QUFBQSxJQUMzRDtBQUdBLFFBQUksbUJBQUssbUJBQWtCO0FBQ3pCLGFBQU8sb0JBQW9CLFlBQVksbUJBQUssaUJBQWdCO0FBQUEsSUFDOUQ7QUFHQSxRQUFJLG1CQUFLLDJCQUEwQjtBQUNqQyxtQkFBYSxtQkFBSyx5QkFBd0I7QUFDMUMseUJBQUssMEJBQTJCO0FBQUEsSUFDbEM7QUFDQSxRQUFJLG1CQUFLLDRCQUEyQjtBQUNsQyxtQkFBYSxtQkFBSywwQkFBeUI7QUFDM0MseUJBQUssMkJBQTRCO0FBQUEsSUFDbkM7QUFDQSxRQUFJLG1CQUFLLDZCQUE0QjtBQUNuQyxtQkFBYSxtQkFBSywyQkFBMEI7QUFDNUMseUJBQUssNEJBQTZCO0FBQUEsSUFDcEM7QUFHQSx1QkFBSyxXQUFVLFdBQVc7QUFHMUIsUUFBSSxtQkFBSyxZQUFXO0FBQ2xCLHlCQUFLLFdBQVUsUUFBUTtBQUFBLElBQ3pCO0FBQUEsRUFDRjtBQUNGO0FBdG9FTztBQUlFO0FBNkJBO0FBZUE7QUFVQTtBQVdBO0FBWUU7QUFHQTtBQUdBO0FBR0E7QUFHQTtBQUdBO0FBR0E7QUFHQTtBQUdBO0FBR0E7QUFHQTtBQS9HSjtBQWlITCxPQUFFLFNBQUMsSUFBWTtBQUNiLFNBQU8sS0FBSyxZQUFZLGVBQWUsRUFBRTtBQUMzQztBQUVBLFFBQUcsU0FBQyxVQUFrQjtBQUNwQixTQUFPLEtBQUssWUFBWSxpQkFBaUIsUUFBUSxLQUFLLENBQUM7QUFDekQ7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBR0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFHQTtBQUNBO0FBQ0E7QUFDQTtBQUdBO0FBQ0E7QUFDQTtBQUdBO0FBQ0E7QUFDQTtBQUNBO0FBSUE7QUEwQkE7QUFDQTtBQUVBLGtCQUFhLFdBQUc7QUFDZCxxQkFBSyxXQUFZLElBQUksZ0JBQWdCO0FBQUEsSUFDckMsV0FBVztBQUFBLElBQ1gsa0JBQWtCO0FBQUEsSUFDbEIsZ0JBQWdCO0FBQUE7QUFBQSxJQUVoQixXQUFXO0FBQUEsTUFDVCxRQUFRLE1BQU07QUFDWiw4QkFBSyxpQ0FBTCxXQUFRLHVCQUF1QixNQUFNO0FBQUEsTUFDdkM7QUFBQSxNQUNBLFNBQVMsQ0FBQyxjQUFtRTtBQUMzRSxZQUFJLFdBQVcsU0FBUyxXQUFXLFNBQVM7QUFDMUMsa0JBQVEsTUFBTSw2QkFBNkIsU0FBUztBQUNwRCxVQUFDLHNCQUFLLGlDQUFMLFdBQVEsa0JBQTBCLEtBQUssVUFBVSxPQUFPLFVBQVUsU0FBUyxVQUFVLElBQUk7QUFBQSxRQUM1RixPQUFPO0FBQ0wsa0JBQVEsTUFBTSxnQ0FBZ0MsU0FBUztBQUFBLFFBQ3pEO0FBQUEsTUFDRjtBQUFBLE1BQ0EsZ0JBQWdCLENBQUMsRUFBRSxTQUFTLE1BQU0sTUFBMEM7QUFDMUUsWUFBSSxXQUFXLElBQUk7QUFDakIsVUFBQyxzQkFBSyxpQ0FBTCxXQUFRLHVCQUErQixVQUFVO0FBQ2xELFVBQUMsc0JBQUssaUNBQUwsV0FBUSx5QkFBaUMsZ0JBQWdCLFNBQVMsS0FBSztBQUFBLFFBQzFFO0FBQUEsTUFDRjtBQUFBLE1BQ0EsVUFBVSxNQUFNO0FBQ2QsY0FBTSxlQUFlLHNCQUFLLGlDQUFMLFdBQVE7QUFDN0IsWUFBSSxjQUFjLGFBQWEsTUFBTSxHQUFHO0FBQ3RDLFVBQUMsYUFBcUIsS0FBSztBQUFBLFFBQzdCO0FBQ0EsZUFBTyxTQUFTLE9BQU87QUFBQSxNQUN6QjtBQUFBLE1BQ0EsWUFBWSxNQUFNO0FBQ2hCLFFBQUMsc0JBQUssaUNBQUwsV0FBUSx1QkFBK0IsVUFBVTtBQUNsRCxRQUFDLHNCQUFLLGlDQUFMLFdBQVEseUJBQWlDLGdCQUFnQixJQUFJLEdBQUs7QUFBQSxNQUNyRTtBQUFBLE1BQ0EsUUFBUSxDQUFDLFNBQWlFO0FBQ3hFLGVBQU8sY0FBYyxJQUFJLGFBQWEsSUFBSSxDQUFDO0FBQUEsTUFDN0M7QUFBQSxJQUNGO0FBQUE7QUFBQSxFQUVBLENBQUM7QUFDSDtBQXlSQSx3QkFBbUIsV0FBRztBQUNwQixNQUFJLENBQUMsS0FBSyxVQUFXLFFBQU87QUFFNUIsTUFBSSxRQUFRO0FBQ1osTUFBSSxPQUFPO0FBRVgsTUFBSSxLQUFLLFVBQVUsU0FBUyxZQUFZLEdBQUc7QUFDekMsWUFBUTtBQUNSLFdBQU87QUFBQSxFQUNULFdBQVcsS0FBSyxVQUFVLFNBQVMsWUFBWSxHQUFHO0FBQ2hELFlBQVE7QUFDUixXQUFPO0FBQUEsRUFDVCxXQUFXLEtBQUssVUFBVSxTQUFTLGVBQWUsR0FBRztBQUNuRCxZQUFRO0FBQ1IsV0FBTztBQUFBLEVBQ1Q7QUFFQSxTQUFPO0FBQUE7QUFBQSw4QkFFbUIsS0FBSyxTQUFTO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSw2QkFLZixLQUFLO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSx1QkFLWCxJQUFJO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFLekI7QUF3RU0sb0JBQWUsaUJBQUc7QUFFdEIsUUFBTSxZQUFZLHNCQUFLLGlDQUFMLFdBQVE7QUFDMUIsUUFBTSxPQUFPLHNCQUFLLGlDQUFMLFdBQVE7QUFDckIsTUFBSSxXQUFXO0FBQ2IsVUFBTSxVQUFVLHNCQUFLLDZDQUFMO0FBQ2hCLGNBQVUsY0FBYztBQUFBLEVBQzFCO0FBQ0EsTUFBSSxNQUFNO0FBQ1IsU0FBSyxjQUFjLFVBQVU7QUFBQSxFQUMvQjtBQUdBLFFBQU0sT0FBTyxNQUFNLE1BQU0sY0FBYyxFQUNwQyxLQUFLLFNBQU8sSUFBSSxLQUFLLENBQUMsRUFDdEIsTUFBTSxDQUFDLFFBQWU7QUFDckIsWUFBUSxNQUFNLGtEQUFrRCxHQUFHO0FBQUEsRUFDckUsQ0FBQztBQUVILE1BQUksQ0FBQyxLQUFNO0FBQ1gsUUFBTSxZQUFZLHNCQUFLLGlDQUFMLFdBQVE7QUFDMUIsUUFBTSxPQUFPLHNCQUFLLGlDQUFMLFdBQVE7QUFDckIsUUFBTSxhQUFhLHNCQUFLLGlDQUFMLFdBQVE7QUFDM0IsUUFBTSxpQkFBaUIsc0JBQUssaUNBQUwsV0FBUTtBQUMvQixRQUFNLGNBQWMsc0JBQUssaUNBQUwsV0FBUTtBQUM1QixRQUFNLG9CQUFvQixzQkFBSyxpQ0FBTCxXQUFRO0FBQ2xDLFFBQU0sY0FBYyxzQkFBSyxpQ0FBTCxXQUFRO0FBRTVCLE1BQUksVUFBVyxXQUFVLGNBQWMsS0FBSyxXQUFXO0FBQ3ZELE1BQUksS0FBTSxNQUFLLGNBQWMsS0FBSyxNQUFNO0FBQ3hDLE1BQUksV0FBWSxZQUFXLGNBQWMsS0FBSyxZQUFZO0FBQzFELE1BQUksZUFBZ0IsZ0JBQWUsY0FBYyxLQUFLLGdCQUFnQjtBQUN0RSxNQUFJLFlBQWEsYUFBWSxjQUFjLEtBQUssYUFBYTtBQUU3RCxNQUFJLG1CQUFtQjtBQUNyQiwwQkFBSyxnREFBTCxXQUF1QixtQkFBbUIsS0FBSztBQUFBLEVBQ2pEO0FBRUEsTUFBSSxlQUFlLEtBQUssV0FBVztBQUNqQyxVQUFNLGdCQUFnQixLQUFLLFVBQVUsS0FBSyxXQUFXLE1BQU0sQ0FBQztBQUM1RCxnQkFBWSxjQUFjO0FBQzFCLFNBQUssZ0JBQWdCO0FBQUEsRUFDdkIsV0FBVyxhQUFhO0FBQ3RCLGdCQUFZLGNBQWM7QUFBQSxFQUM1QjtBQUVBLHFCQUFLLFlBQWE7QUFDcEI7QUFFQSxzQkFBaUIsU0FBQyxXQUF3QixPQUEyQjtBQUNuRSxNQUFJLENBQUMsT0FBTyxRQUFRO0FBQ2xCLGNBQVUsY0FBYztBQUN4QjtBQUFBLEVBQ0Y7QUFFQSxRQUFNLGlCQUFpQixLQUFLLGtCQUFrQjtBQUM5QyxRQUFNLGVBQWUsQ0FBQyxDQUFDO0FBRXZCLE1BQUksY0FBYztBQUNoQixVQUFNLGNBQWMsTUFBTSxLQUFLLFVBQVEsS0FBSyxZQUFZLGNBQWM7QUFDdEUsUUFBSSxDQUFDLGFBQWE7QUFDaEIsZ0JBQVUsY0FBYztBQUN4QjtBQUFBLElBQ0Y7QUFFQSxVQUFNLFdBQVcsOEJBQWUsbUJBQWtCLFFBQVEsVUFBVSxJQUFJO0FBRXhFLGFBQVMsY0FBYyx3QkFBd0IsRUFBRyxjQUFjLFlBQVk7QUFDNUUsYUFBUyxjQUFjLDZCQUE2QixFQUFHLGNBQWMsWUFBWTtBQUNqRixhQUFTLGNBQWMsMkJBQTJCLEVBQUcsY0FBYyxZQUFZO0FBRS9FLFVBQU0sbUJBQW1CLFNBQVMsY0FBYyxrQ0FBa0M7QUFDbEYsUUFBSSxZQUFZLGFBQWE7QUFDM0IsZUFBUyxjQUFjLDRCQUE0QixFQUFHLGNBQWMsWUFBWTtBQUFBLElBQ2xGLE9BQU87QUFDTCx3QkFBa0IsT0FBTztBQUFBLElBQzNCO0FBRUEsVUFBTSxnQkFBZ0IsU0FBUyxjQUFjLCtCQUErQjtBQUM1RSxRQUFJLFlBQVksVUFBVTtBQUN4QixlQUFTLGNBQWMseUJBQXlCLEVBQUcsY0FBYyxZQUFZO0FBQUEsSUFDL0UsT0FBTztBQUNMLHFCQUFlLE9BQU87QUFBQSxJQUN4QjtBQUVBLGNBQVUsZ0JBQWdCLFFBQVE7QUFBQSxFQUNwQyxPQUFPO0FBQ0wsVUFBTSxlQUFlLDhCQUFlLG1CQUFrQixRQUFRLFVBQVUsSUFBSTtBQUU1RSxVQUFNLGtCQUFrQixhQUFhLGNBQWMsMkJBQTJCO0FBRTlFLGVBQVcsUUFBUSxPQUFPO0FBQ3hCLFlBQU0sZ0JBQWdCLDhCQUFlLG9CQUFtQixRQUFRLFVBQVUsSUFBSTtBQUU5RSxvQkFBYyxjQUFjLHdCQUF3QixFQUFHLGNBQWMsS0FBSztBQUMxRSxvQkFBYyxjQUFjLDRCQUE0QixFQUFHLGNBQ3pELEtBQUssZUFBZTtBQUN0QixvQkFBYyxjQUFjLDZCQUE2QixFQUFHLGNBQWMsS0FBSztBQUMvRSxvQkFBYyxjQUFjLDJCQUEyQixFQUFHLGNBQWMsS0FBSztBQUU3RSxZQUFNLGdCQUFnQixjQUFjLGNBQWMsK0JBQStCO0FBQ2pGLFVBQUksS0FBSyxVQUFVO0FBQ2pCLHNCQUFjLGNBQWMseUJBQXlCLEVBQUcsY0FBYyxLQUFLO0FBQUEsTUFDN0UsT0FBTztBQUNMLHVCQUFlLE9BQU87QUFBQSxNQUN4QjtBQUVBLHNCQUFnQixZQUFZLGFBQWE7QUFBQSxJQUMzQztBQUVBLGNBQVUsZ0JBQWdCLFlBQVk7QUFBQSxFQUN4QztBQUNGO0FBRUEsc0JBQWlCLFdBQUc7QUFDbEIscUJBQUssZUFBZ0Isc0JBQUssaUNBQUwsV0FBUTtBQUU3QixRQUFNLGFBQWEsc0JBQUssaUNBQUwsV0FBUTtBQUMzQixNQUFJLFlBQVk7QUFDZCxlQUFXLGlCQUFpQixTQUFTLE1BQU07QUFDekMsWUFBTSxFQUFFLFFBQVEsR0FBRyxJQUFJO0FBQ3ZCLG1CQUFhLG1CQUFLLHlCQUF5QjtBQUMzQyx5QkFBSywwQkFBMkIsV0FBVyxNQUFNO0FBQy9DLDhCQUFLLDBDQUFMLFdBQWlCO0FBQUEsTUFDbkIsR0FBRyxHQUFHO0FBQUEsSUFDUixDQUFDO0FBQUEsRUFDSDtBQUVBLHFCQUFLLG1CQUFvQixLQUFLLFlBQVksY0FBYyxtQkFBbUIsS0FBSztBQUNoRixNQUFJLG1CQUFLLG9CQUFtQjtBQUMxQiwwQkFBc0IsTUFBTTtBQUMxQiw0QkFBSyxrREFBTDtBQUFBLElBQ0YsQ0FBQztBQUNELHVCQUFLLG1CQUFrQixpQkFBaUIsVUFBVSxtQkFBSyx1QkFBdUM7QUFBQSxFQUNoRztBQUVBLHdCQUFLLGlDQUFMLFdBQVEsY0FBYyxpQkFBaUIsU0FBUyxNQUFNO0FBQ3BELDBCQUFLLHdDQUFMO0FBQUEsRUFDRixDQUFDO0FBRUQscUJBQUssbUJBQW9CLENBQUMsVUFBaUI7QUFDekMsVUFBTSxPQUFRLE1BQXVCO0FBQ3JDLFFBQUksTUFBTTtBQUNSLDRCQUFLLDBDQUFMLFdBQWlCO0FBQUEsSUFDbkI7QUFBQSxFQUNGO0FBQ0EsU0FBTyxpQkFBaUIsWUFBWSxtQkFBSyxpQkFBZ0I7QUFDM0Q7QUFFQSxnQkFBVyxTQUFDLE9BQWU7QUFDekIscUJBQUssa0JBQW1CLE1BQU0sWUFBWTtBQUUxQyxNQUFJLENBQUMsbUJBQUssZUFBZTtBQUV6QixhQUFXLFNBQVMsbUJBQUssZUFBYyxVQUFVO0FBQy9DLFVBQU0sT0FBTyxNQUFNLGFBQWEsWUFBWSxLQUFLO0FBQ2pELFVBQU0sWUFBWSxDQUFDLG1CQUFLLHFCQUFvQixLQUFLLFNBQVMsbUJBQUssaUJBQWdCO0FBRS9FLFVBQU0sVUFBVSxzQkFBSyxtREFBTCxXQUEwQjtBQUMxQyxVQUFNLGFBQWEsbUJBQUssa0JBQWlCLElBQUksT0FBTztBQUVwRCxJQUFDLE1BQXNCLFNBQVMsRUFBRSxhQUFhO0FBQUEsRUFDakQ7QUFDRjtBQUVBLHlCQUFvQixTQUFDLE9BQXdCO0FBQzNDLGFBQVcsT0FBTyxNQUFNLFdBQVc7QUFDakMsUUFBSSxDQUFDLFFBQVEsV0FBVyxTQUFTLE9BQU8sRUFBRSxTQUFTLEdBQUcsR0FBRztBQUN2RCxhQUFPLFFBQVEsWUFBWSxTQUFTO0FBQUEsSUFDdEM7QUFBQSxFQUNGO0FBQ0EsU0FBTztBQUNUO0FBRUEsd0JBQW1CLFdBQUc7QUFDcEIsTUFBSTtBQUNGLFVBQU0sUUFBUSxhQUFhLFFBQVEsdUJBQXVCO0FBQzFELFFBQUksT0FBTztBQUNULHlCQUFLLGtCQUFtQixJQUFJLElBQUksS0FBSyxNQUFNLEtBQUssQ0FBQztBQUFBLElBQ25EO0FBQUEsRUFDRixTQUFTLEdBQUc7QUFDVixZQUFRLE1BQU0sd0VBQXdFO0FBQUEsRUFDeEY7QUFDQSx3QkFBSyxrREFBTDtBQUNGO0FBRUEsd0JBQW1CLFdBQUc7QUFDcEIsTUFBSSxDQUFDLG1CQUFLLG1CQUFtQjtBQUM3QixRQUFNLFlBQVksbUJBQUssbUJBQWtCLGlCQUFpQixpQkFBaUI7QUFDM0UsWUFBVSxRQUFRLFVBQVE7QUFDeEIsVUFBTSxRQUFTLEtBQWE7QUFDNUIsSUFBQyxLQUFhLFVBQVUsbUJBQUssa0JBQWlCLElBQUksS0FBSztBQUFBLEVBQ3pELENBQUM7QUFDSDtBQUVBLHdCQUFtQixXQUFHO0FBQ3BCLE1BQUk7QUFDRixpQkFBYTtBQUFBLE1BQVE7QUFBQSxNQUNuQixLQUFLLFVBQVUsQ0FBQyxHQUFHLG1CQUFLLGlCQUFnQixDQUFDO0FBQUEsSUFBQztBQUFBLEVBQzlDLFNBQVMsR0FBRztBQUFBLEVBRVo7QUFDRjtBQUVBO0FBYU0sY0FBUyxpQkFBRztBQUNoQixNQUFJLENBQUMsbUJBQUssZUFBZTtBQUV6QixRQUFNLE9BQU8sTUFBTSxLQUFLLG1CQUFLLGVBQWMsUUFBUSxFQUNoRCxPQUFPLFdBQVMsQ0FBRSxNQUFzQixNQUFNLEVBQzlDLElBQUksV0FBUztBQUNaLFVBQU0sT0FBTyxNQUFNLGNBQWMsc0JBQXNCLEdBQUcsYUFBYSxLQUFLLEtBQUs7QUFDakYsVUFBTSxPQUFPLE1BQU0sY0FBYyxxQkFBcUIsR0FBRyxhQUFhLEtBQUssS0FBSztBQUNoRixVQUFNLFVBQVUsTUFBTSxjQUFjLHdCQUF3QixHQUFHLGFBQWEsS0FBSyxLQUFLO0FBQ3RGLFdBQU8sSUFBSSxJQUFJLEtBQUssSUFBSSxJQUFJLE9BQU87QUFBQSxFQUNyQyxDQUFDLEVBQUUsS0FBSyxJQUFJO0FBRWQsTUFBSSxDQUFDLEtBQU07QUFFWCxNQUFJO0FBQ0YsVUFBTSxVQUFVLFVBQVUsVUFBVSxJQUFJO0FBQ3hDLFVBQU0sTUFBTSxzQkFBSyxpQ0FBTCxXQUFRO0FBQ3BCLFFBQUksS0FBSztBQUNQLFlBQU0sV0FBVyxNQUFNLEtBQUssSUFBSSxVQUFVLEVBQUU7QUFBQSxRQUMxQyxPQUFLLEVBQUUsYUFBYSxLQUFLLGNBQWMsRUFBRSxhQUFhLEtBQUssRUFBRSxVQUFVLEtBQUs7QUFBQSxNQUM5RTtBQUNBLFVBQUksVUFBVTtBQUNaLGNBQU0sV0FBVyxTQUFTO0FBQzFCLGlCQUFTLGNBQWM7QUFFdkIsWUFBSSxtQkFBSywyQkFBMEI7QUFDakMsdUJBQWEsbUJBQUsseUJBQXdCO0FBQUEsUUFDNUM7QUFFQSwyQkFBSywwQkFBMkIsV0FBVyxNQUFNO0FBQy9DLGNBQUksS0FBSyxlQUFlLFNBQVMsWUFBWTtBQUMzQyxxQkFBUyxjQUFjO0FBQUEsVUFDekI7QUFDQSw2QkFBSywwQkFBMkI7QUFBQSxRQUNsQyxHQUFHLEdBQUk7QUFBQSxNQUNUO0FBQUEsSUFDRjtBQUFBLEVBQ0YsU0FBUyxLQUFLO0FBQ1osWUFBUSxNQUFNLDJDQUEyQyxHQUFHO0FBQUEsRUFDOUQ7QUFDRjtBQUVBLHVCQUFrQixXQUFHO0FBQ25CLFFBQU0sY0FBYyxzQkFBSyxpQ0FBTCxXQUFRO0FBQzVCLFFBQU0sYUFBYSxzQkFBSyxpQ0FBTCxXQUFRO0FBQzNCLFFBQU0sYUFBYSxLQUFLLFlBQVksY0FBYyxjQUFjO0FBQ2hFLFFBQU0sWUFBWSxLQUFLLFlBQVksY0FBYyxhQUFhO0FBRTlELE1BQUksZUFBZSxZQUFZO0FBQzdCLGdCQUFZLGlCQUFpQixTQUFTLE1BQU07QUFDMUMsNEJBQUssOENBQUw7QUFDQSxNQUFDLFdBQW1CLFVBQVU7QUFBQSxJQUNoQyxDQUFDO0FBRUQsZ0JBQVksaUJBQWlCLFNBQVMsTUFBTyxXQUFtQixNQUFNLENBQUM7QUFFdkUsZUFBVyxpQkFBaUIsU0FBUyxNQUFNO0FBQ3pDLDRCQUFLLDZDQUFMO0FBQUEsSUFDRixDQUFDO0FBQUEsRUFDSDtBQUNGO0FBRUEsdUJBQWtCLFdBQUc7QUFDbkIsUUFBTSxTQUFTLEtBQUssWUFBWSxjQUFjLFlBQVk7QUFDMUQsUUFBTSxPQUFPLEtBQUssWUFBWSxjQUFjLFlBQVk7QUFFeEQsTUFBSSxDQUFDLFVBQVUsQ0FBQyxLQUFNO0FBRXRCLHFCQUFLLGFBQWUsT0FBZTtBQUVuQyxTQUFPLGlCQUFpQixVQUFVLENBQUMsTUFBYTtBQUM5Qyx1QkFBSyxhQUFlLEVBQVU7QUFFOUIscUJBQWlCLFlBQVk7QUFBQSxNQUMzQixRQUFRLEVBQUUsTUFBTyxFQUFVLEtBQUs7QUFBQSxJQUNsQyxDQUFDO0FBRUQsUUFBSyxFQUFVLE1BQU07QUFDbkIsNEJBQUssa0RBQUw7QUFBQSxJQUNGO0FBQUEsRUFDRixDQUFDO0FBRUQsU0FBTyxpQkFBaUIsVUFBVSxDQUFDLE1BQWE7QUFDOUMsSUFBQyxPQUFlLGFBQWEsaUJBQWtCLEVBQVUsTUFBTTtBQUUvRCxxQkFBaUIsWUFBWTtBQUFBLE1BQzNCLFFBQVEsRUFBRSxRQUFTLEVBQVUsT0FBTztBQUFBLElBQ3RDLENBQUM7QUFBQSxFQUNILENBQUM7QUFFRCxPQUFLLGlCQUFpQixVQUFVLENBQUMsTUFBYTtBQUM1QyxxQkFBaUIsWUFBWTtBQUFBLE1BQzNCLE1BQU0sRUFBRSxlQUFnQixFQUFVLGNBQWM7QUFBQSxJQUNsRCxDQUFDO0FBRUQsUUFBSyxFQUFVLGtCQUFrQixLQUFNLE9BQWUsTUFBTTtBQUMxRCw0QkFBSyxrREFBTDtBQUFBLElBQ0Y7QUFFQSxRQUFLLEVBQVUsa0JBQWtCLEtBQU0sT0FBZSxNQUFNO0FBQzFELDRCQUFLLG9EQUFMO0FBQUEsSUFDRjtBQUFBLEVBQ0YsQ0FBQztBQUNIO0FBRUEsbUJBQWMsV0FBVztBQUN2QixRQUFNLEtBQUssVUFBVTtBQUNyQixNQUFJLEdBQUcsU0FBUyxVQUFVLEdBQUc7QUFDM0IsVUFBTSxRQUFRLEdBQUcsTUFBTSxnQkFBZ0I7QUFDdkMsV0FBTyxRQUFRLFdBQVcsTUFBTSxDQUFDLENBQUMsS0FBSztBQUFBLEVBQ3pDLFdBQVcsR0FBRyxTQUFTLE1BQU0sR0FBRztBQUM5QixVQUFNLFFBQVEsR0FBRyxNQUFNLFlBQVk7QUFDbkMsV0FBTyxRQUFRLFFBQVEsTUFBTSxDQUFDLENBQUMsS0FBSztBQUFBLEVBQ3RDLFdBQVcsR0FBRyxTQUFTLFNBQVMsR0FBRztBQUNqQyxVQUFNLFFBQVEsR0FBRyxNQUFNLGVBQWU7QUFDdEMsV0FBTyxRQUFRLFVBQVUsTUFBTSxDQUFDLENBQUMsS0FBSztBQUFBLEVBQ3hDLFdBQVcsR0FBRyxTQUFTLFNBQVMsS0FBSyxDQUFDLEdBQUcsU0FBUyxRQUFRLEdBQUc7QUFDM0QsVUFBTSxRQUFRLEdBQUcsTUFBTSxnQkFBZ0I7QUFDdkMsV0FBTyxRQUFRLFVBQVUsTUFBTSxDQUFDLENBQUMsS0FBSztBQUFBLEVBQ3hDO0FBQ0EsU0FBTztBQUNUO0FBRU0sbUJBQWMsaUJBQUc7QUFDckIsUUFBTSxPQUFPLE1BQU0sS0FBSyxzQkFBSyxrQ0FBTCxXQUFTLHFCQUFxQixFQUFFLElBQUksUUFBTTtBQUNoRSxVQUFNLEtBQUssR0FBRztBQUNkLFFBQUksTUFBTSxHQUFHLFlBQVksTUFBTTtBQUM3QixhQUFPLEdBQUcsR0FBRyxXQUFXLEtBQUssR0FBRyxXQUFXO0FBQUEsSUFDN0M7QUFDQSxXQUFPO0FBQUEsRUFDVCxDQUFDLEVBQUUsS0FBSyxJQUFJO0FBRVosTUFBSSxtQkFBbUI7QUFDdkIsTUFBSSxtQkFBSyxhQUFZLGVBQWU7QUFDbEMsdUJBQW1CO0FBQUEsRUFBSyxJQUFJLE9BQU8sRUFBRSxDQUFDO0FBQUE7QUFBQSxFQUFrQixJQUFJLE9BQU8sRUFBRSxDQUFDO0FBQUEsRUFBSyxtQkFBSyxZQUFXLGFBQWE7QUFBQTtBQUFBLEVBQzFHO0FBRUEsUUFBTSxZQUFZO0FBQUEsRUFDcEIsSUFBSSxPQUFPLEVBQUUsQ0FBQztBQUFBLEVBQ2QsSUFBSSxHQUFHLGdCQUFnQjtBQUFBLEVBQ3ZCLElBQUksT0FBTyxFQUFFLENBQUM7QUFBQSxjQUNILG9CQUFJLEtBQUssR0FBRSxZQUFZLENBQUM7QUFFakMsTUFBSTtBQUNGLFVBQU0sVUFBVSxVQUFVLFVBQVUsU0FBUztBQUM3QyxVQUFNLGFBQWEsS0FBSyxZQUFZLGNBQWMsYUFBYTtBQUMvRCxRQUFJLFlBQVk7QUFDZCxZQUFNLGVBQWUsV0FBVztBQUNoQyxpQkFBVyxjQUFjO0FBRXpCLFVBQUksbUJBQUssNEJBQTJCO0FBQ2xDLHFCQUFhLG1CQUFLLDBCQUF5QjtBQUFBLE1BQzdDO0FBRUEseUJBQUssMkJBQTRCLFdBQVcsTUFBTTtBQUNoRCxZQUFJLEtBQUssZUFBZSxXQUFXLFlBQVk7QUFDN0MscUJBQVcsY0FBYztBQUFBLFFBQzNCO0FBQ0EsMkJBQUssMkJBQTRCO0FBQUEsTUFDbkMsR0FBRyxHQUFJO0FBQUEsSUFDVDtBQUFBLEVBQ0YsU0FBUyxLQUFLO0FBQ1osWUFBUSxNQUFNLGlEQUFpRCxHQUFHO0FBQUEsRUFDcEU7QUFDRjtBQUVBLGdCQUFXLFNBQUMsTUFBOEQ7QUFDeEUsTUFBSSxDQUFDLG1CQUFLLGVBQWU7QUFFekIsUUFBTSxjQUFjLEtBQUssSUFBSSxTQUFPO0FBQ2xDLFVBQU0sV0FBVyw4QkFBZSxtQkFBa0IsUUFBUSxVQUFVLElBQUk7QUFFeEUsVUFBTSxPQUFPLElBQUksS0FBSyxJQUFJLElBQUk7QUFDOUIsVUFBTSxPQUFPLEtBQUssbUJBQW1CO0FBRXJDLFVBQU0sWUFBWSxTQUFTLGNBQWMsMEJBQTBCO0FBQ25FLGNBQVUsVUFBVSxJQUFJLElBQUksSUFBSTtBQUNoQyxjQUFVLGFBQWEsZUFBZSxJQUFJLElBQUk7QUFFOUMsVUFBTSxZQUFZLHNCQUFLLDJDQUFMLFdBQWtCLElBQUk7QUFDeEMsVUFBTSxhQUFhLEdBQUcsU0FBUyxJQUFJLElBQUksSUFBSSxJQUFJLE9BQU8sR0FBRyxZQUFZO0FBQ3JFLFVBQU0sWUFBWSxDQUFDLG1CQUFLLHFCQUFvQixXQUFXLFNBQVMsbUJBQUssaUJBQWdCO0FBRXJGLFVBQU0sbUJBQW1CLElBQUksU0FBUyxZQUFZLFNBQVMsSUFBSTtBQUMvRCxVQUFNLGFBQWEsbUJBQUssa0JBQWlCLElBQUksZ0JBQWdCO0FBRTdELFFBQUksRUFBRSxhQUFhLGFBQWE7QUFDOUIsZ0JBQVUsYUFBYSxVQUFVLEVBQUU7QUFBQSxJQUNyQztBQUVBLFVBQU0sUUFBUSxTQUFTLGNBQWMsc0JBQXNCO0FBQzNELFVBQU0sY0FBYyxzQkFBSywyQ0FBTCxXQUFrQixJQUFJO0FBQzFDLDBCQUFLLGtEQUFMLFdBQXlCLE9BQU8sSUFBSTtBQUVwQyxVQUFNLFNBQVMsU0FBUyxjQUFjLHFCQUFxQjtBQUMzRCxXQUFPLGFBQWEsWUFBWSxJQUFJLElBQUk7QUFDeEMsV0FBTyxjQUFjO0FBRXJCLElBQUMsU0FBUyxjQUFjLHdCQUF3QixFQUFrQixjQUFjLElBQUk7QUFFcEYsV0FBTztBQUFBLEVBQ1QsQ0FBQztBQUVELE1BQUksQ0FBQyxtQkFBSyxzQkFBcUI7QUFDN0IsdUJBQUssZUFBYyxnQkFBZ0IsR0FBRyxXQUFXO0FBQ2pELHVCQUFLLHFCQUFzQjtBQUUzQixRQUFJLG1CQUFLLGNBQWE7QUFDcEIsNEJBQUssb0RBQUw7QUFBQSxJQUNGO0FBQUEsRUFDRixPQUFPO0FBQ0wsdUJBQUssZUFBYyxPQUFPLEdBQUcsV0FBVztBQUV4QyxRQUFJLG1CQUFLLGNBQWE7QUFDcEIsNEJBQUssb0RBQUw7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUNGO0FBRUEsaUJBQVksU0FBQyxNQUFzQjtBQUNqQyxVQUFRLE1BQU07QUFBQSxJQUNaLEtBQUs7QUFBUSxhQUFPO0FBQUEsSUFDcEIsS0FBSztBQUFXLGFBQU87QUFBQSxJQUN2QixLQUFLO0FBQVMsYUFBTztBQUFBLElBQ3JCLEtBQUs7QUFBUyxhQUFPO0FBQUEsSUFDckI7QUFBUyxhQUFPLEtBQUssWUFBWTtBQUFBLEVBQ25DO0FBQ0Y7QUFFQSx3QkFBbUIsU0FBQyxPQUFvQixNQUFjO0FBQ3BELFVBQVEsTUFBTTtBQUFBLElBQ1osS0FBSztBQUNILGFBQU8sTUFBTSxhQUFhLFVBQVUsTUFBTTtBQUFBLElBQzVDLEtBQUs7QUFDSCxhQUFPLE1BQU0sYUFBYSxVQUFVLFNBQVM7QUFBQSxJQUMvQyxLQUFLO0FBQ0gsYUFBTyxNQUFNLGFBQWEsVUFBVSxRQUFRO0FBQUEsSUFDOUMsS0FBSztBQUNILGFBQU8sTUFBTSxhQUFhLFNBQVMsUUFBUTtBQUFBLElBQzdDO0FBQ0UsWUFBTSxhQUFhLFNBQVMsTUFBTTtBQUFBLEVBQ3RDO0FBQ0Y7QUFFQSwwQkFBcUIsV0FBRztBQUN0QixNQUFJLENBQUMsbUJBQUssZUFBZTtBQUV6Qix3QkFBc0IsTUFBTTtBQUMxQixVQUFNLFVBQVUsbUJBQUssZUFBZTtBQUNwQyxRQUFJLFNBQVM7QUFDWCxjQUFRLGVBQWUsRUFBRSxVQUFVLFFBQVEsT0FBTyxNQUFNLENBQUM7QUFBQSxJQUMzRDtBQUFBLEVBQ0YsQ0FBQztBQUNIO0FBRUEsd0JBQW1CLFdBQUc7QUFDcEIsTUFBSSxDQUFDLG1CQUFLLGVBQWU7QUFFekIsTUFBSSxtQkFBSyxpQkFBZ0I7QUFDdkIsMEJBQUssb0RBQUw7QUFBQSxFQUNGLE9BQU87QUFDTCxlQUFXLE1BQU07QUFDZiw0QkFBSyxvREFBTDtBQUFBLElBQ0YsR0FBRyxHQUFHO0FBQUEsRUFDUjtBQUNGO0FBRUEscUNBQWdDLFdBQUc7QUFDakMsTUFBSTtBQUNGLFVBQU0sa0JBQ0osYUFBYSxRQUFRLHdCQUF3QixNQUFNLFFBQ25ELGFBQWEsUUFBUSx1QkFBdUIsTUFBTSxRQUNsRCxhQUFhLFFBQVEseUJBQXlCLE1BQU0sUUFDcEQsYUFBYSxRQUFRLHNCQUFzQixNQUFNO0FBRW5ELFFBQUksaUJBQWlCO0FBQ25CLFlBQU0sV0FBVyxhQUFhLFFBQVEsK0JBQStCO0FBQ3JFLFVBQUksQ0FBQyxVQUFVO0FBQ2IseUJBQWlCLHdCQUF3QjtBQUN6QyxxQkFBYSxRQUFRLGlDQUFpQyxNQUFNO0FBQzVELG1CQUFXLE1BQU0sT0FBTyxTQUFTLE9BQU8sR0FBRyxHQUFHO0FBQUEsTUFDaEQ7QUFBQSxJQUNGO0FBQUEsRUFDRixTQUFTLEdBQUc7QUFBQSxFQUVaO0FBQ0Y7QUFFQSw0QkFBdUIsV0FBRztBQUN4QixRQUFNLGNBQWMsS0FBSyxZQUFZLGNBQWMsc0JBQXNCO0FBQ3pFLE1BQUksQ0FBQyxZQUFhO0FBRWxCLFFBQU0sUUFBUSxpQkFBaUIsU0FBUztBQUV4Qyx3QkFBSyxnREFBTCxXQUF1QixNQUFNO0FBRTdCLFFBQU0sUUFBUSxZQUFZLGlCQUFpQix5QkFBeUI7QUFDcEUsUUFBTSxRQUFRLFVBQVE7QUFDcEIsUUFBSyxLQUFhLFVBQVUsTUFBTSxhQUFhO0FBQzdDLFdBQUssYUFBYSxZQUFZLEVBQUU7QUFBQSxJQUNsQztBQUFBLEVBQ0YsQ0FBQztBQUVELGNBQVksaUJBQWlCLDZCQUE2QixDQUFDLE1BQWE7QUFDdEUsVUFBTSxTQUFVLEVBQVU7QUFDMUIsMEJBQUssZ0RBQUwsV0FBdUI7QUFDdkIscUJBQWlCLFlBQVksRUFBRSxhQUFhLE9BQU8sQ0FBQztBQUFBLEVBQ3RELENBQUM7QUFDSDtBQUVBLHNCQUFpQixTQUFDLFFBQWdCO0FBQ2hDLFFBQU0sT0FBTyxTQUFTO0FBRXRCLFVBQVEsUUFBUTtBQUFBLElBQ2QsS0FBSztBQUNILFdBQUssTUFBTSxjQUFjO0FBQ3pCO0FBQUEsSUFDRixLQUFLO0FBQ0gsV0FBSyxNQUFNLGNBQWM7QUFDekI7QUFBQSxJQUNGLEtBQUs7QUFBQSxJQUNMO0FBQ0UsV0FBSyxNQUFNLGNBQWM7QUFDekI7QUFBQSxFQUNKO0FBQ0Y7QUFFQSwyQkFBc0IsV0FBRztBQUN2QixPQUFLLGlCQUFpQix5QkFBeUIsbUJBQUssY0FBYTtBQUNqRSxPQUFLLGlCQUFpQix3QkFBd0IsbUJBQUssY0FBYTtBQUNoRSxPQUFLLGlCQUFpQiw0QkFBNEIsbUJBQUssY0FBYTtBQUNwRSxPQUFLLGlCQUFpQix3QkFBd0IsbUJBQUssYUFBWTtBQUMvRCxPQUFLLGlCQUFpQix1QkFBdUIsbUJBQUssYUFBWTtBQUM5RCxPQUFLLGlCQUFpQiwyQkFBMkIsbUJBQUssYUFBWTtBQUNwRTtBQUVBO0FBZ0NBO0FBaUNBLG1CQUFjLFNBQUMsT0FBaUU7QUFDOUUsUUFBTSxpQkFBaUIsS0FBSyxrQkFBa0I7QUFFOUMsTUFBSSxNQUFNLGNBQWM7QUFDdEIsZUFBVyxXQUFXLE1BQU0sYUFBYSxHQUFHO0FBQzFDLFVBQUksRUFBRSxtQkFBbUIsU0FBVTtBQUVuQyxVQUFLLFFBQXdCLFNBQVMsa0JBQWtCLFFBQVE7QUFDOUQsY0FBTSxVQUFXLFFBQXdCLFFBQVEsV0FBVztBQUM1RCxZQUFJLGdCQUFnQixPQUFPLFNBQVUsUUFBd0IsUUFBUSxpQkFBaUIsSUFBSSxFQUFFO0FBQzVGLFlBQUksT0FBTyxNQUFNLGFBQWEsRUFBRyxpQkFBZ0I7QUFDakQsZUFBTyxFQUFFLFNBQVMsY0FBYztBQUFBLE1BQ2xDO0FBQUEsSUFDRjtBQUFBLEVBQ0Y7QUFFQSxTQUFPLEVBQUUsU0FBUyxnQkFBZ0IsZUFBZSxFQUFFO0FBQ3JEO0FBRUEsMEJBQXFCLFNBQUMsT0FBc0I7QUFDMUMsVUFBUSxNQUFNLE1BQU07QUFBQSxJQUNsQixLQUFLO0FBQ0gsYUFBTztBQUFBLElBQ1QsS0FBSztBQUNILGFBQU87QUFBQSxJQUNULEtBQUs7QUFDSCxhQUFPO0FBQUEsSUFDVDtBQUNFLGFBQU87QUFBQSxFQUNYO0FBQ0Y7QUFFQSwrQkFBMEIsU0FBQyxPQUFzQjtBQUMvQyxVQUFRLE1BQU0sTUFBTTtBQUFBLElBQ2xCLEtBQUs7QUFDSCxhQUFPO0FBQUEsSUFDVCxLQUFLO0FBQ0gsYUFBTztBQUFBLElBQ1QsS0FBSztBQUNILGFBQU87QUFBQSxJQUNUO0FBQ0UsYUFBTztBQUFBLEVBQ1g7QUFDRjtBQUVBLCtCQUEwQixXQUFHO0FBQzNCLHFCQUFLLG1CQUFvQixDQUFDLE1BQWE7QUFDckMsUUFBSyxFQUFFLFFBQW9CLFlBQVksa0JBQW1CO0FBRTFELFVBQU0sU0FBUyxzQkFBSyw2Q0FBTCxXQUFvQixFQUFFO0FBQ3JDLFVBQU0sWUFBWSxpQkFBaUIsYUFBYTtBQUNoRCxRQUFJLENBQUMsVUFBVSxTQUFTLFNBQVMsTUFBTSxHQUFHO0FBQ3hDLGdCQUFVLFNBQVMsS0FBSyxNQUFNO0FBQzlCLHVCQUFpQixhQUFhLFNBQVM7QUFBQSxJQUN6QztBQUFBLEVBQ0Y7QUFDQSxPQUFLLGlCQUFpQixVQUFVLG1CQUFLLGtCQUFpQjtBQUV0RCxxQkFBSyxxQkFBc0IsQ0FBQyxNQUFhO0FBQ3ZDLFFBQUssRUFBRSxRQUFvQixZQUFZLGtCQUFtQjtBQUUxRCxVQUFNLFNBQVMsc0JBQUssNkNBQUwsV0FBb0IsRUFBRTtBQUNyQyxVQUFNLFlBQVksaUJBQWlCLGFBQWE7QUFDaEQsVUFBTSxRQUFRLFVBQVUsU0FBUyxRQUFRLE1BQU07QUFDL0MsUUFBSSxRQUFRLElBQUk7QUFDZCxnQkFBVSxTQUFTLE9BQU8sT0FBTyxDQUFDO0FBQ2xDLHVCQUFpQixhQUFhLFNBQVM7QUFBQSxJQUN6QztBQUFBLEVBQ0Y7QUFDQSxPQUFLLGlCQUFpQixZQUFZLG1CQUFLLG9CQUFtQjtBQUUxRCxxQkFBSyxtQkFBb0IsQ0FBQyxNQUFhO0FBQ3JDLFFBQUssRUFBRSxRQUFvQixZQUFZLGtCQUFtQjtBQUUxRCxVQUFNLFNBQVMsc0JBQUssNkNBQUwsV0FBb0IsRUFBRTtBQUNyQyxxQkFBaUIsZ0JBQWdCLEVBQUUsVUFBVSxPQUFPLENBQUM7QUFBQSxFQUN2RDtBQUNBLE9BQUssaUJBQWlCLFVBQVUsbUJBQUssa0JBQWlCO0FBRXRELHdCQUFLLDhDQUFMO0FBQ0Y7QUFFQSxvQkFBZSxXQUFHO0FBQ2hCLFFBQU0sWUFBWSxpQkFBaUIsYUFBYTtBQUVoRCxhQUFXLFVBQVUsVUFBVSxVQUFVO0FBQ3ZDLFVBQU0sV0FBVyxzQkFBSyxnREFBTCxXQUF1QjtBQUN4QyxRQUFJLFlBQVksQ0FBQyxTQUFTLGFBQWEsVUFBVSxHQUFHO0FBQ2xELGVBQVMsYUFBYSxZQUFZLEVBQUU7QUFBQSxJQUN0QztBQUFBLEVBQ0Y7QUFFQSxNQUFJLFVBQVUsVUFBVTtBQUN0QixVQUFNLFdBQVcsc0JBQUssZ0RBQUwsV0FBdUIsVUFBVTtBQUNsRCxRQUFJLFlBQVksQ0FBQyxTQUFTLGFBQWEsU0FBUyxHQUFHO0FBQ2pELGVBQVMsYUFBYSxXQUFXLEVBQUU7QUFBQSxJQUNyQztBQUFBLEVBQ0Y7QUFDRjtBQUVBLGtDQUE2QixXQUFHO0FBQzlCLFFBQU0sT0FBTyxLQUFLLFlBQVksY0FBYyxZQUFZO0FBRXhELE1BQUksQ0FBQyxLQUFNO0FBRVgsT0FBSyxpQkFBaUIsa0JBQWtCLENBQUMsVUFBaUI7QUFDeEQsVUFBTSxZQUFZLENBQUUsTUFBYztBQUVsQyxxQkFBaUIsWUFBWTtBQUFBLE1BQzNCLFNBQVMsRUFBRSxVQUFVO0FBQUEsSUFDdkIsQ0FBQztBQUFBLEVBQ0gsQ0FBQztBQUNIO0FBRUEsc0JBQWlCLFNBQUMsUUFBZ0M7QUFDaEQsUUFBTSxRQUFRLE9BQU8sTUFBTSxHQUFHO0FBQzlCLFFBQU0sQ0FBQyxNQUFNLFlBQVksU0FBUyxJQUFJLElBQUk7QUFFMUMsTUFBSSxhQUFhO0FBQ2pCLE1BQUksU0FBUztBQUNYLGtCQUFjLG1CQUFtQixJQUFJLE9BQU8sT0FBTyxDQUFDO0FBQUEsRUFDdEQ7QUFDQSxNQUFJLE1BQU07QUFDUixrQkFBYyxlQUFlLElBQUksT0FBTyxJQUFJLENBQUM7QUFBQSxFQUMvQztBQUVBLE1BQUksV0FBVyw4QkFBOEIsSUFBSSxPQUFPLElBQUksQ0FBQztBQUM3RCxNQUFJLFlBQVk7QUFDZCxVQUFNLG9CQUFvQixJQUFJLE9BQU8sVUFBVTtBQUMvQyxVQUFNLGNBQWMsSUFBSSxPQUFPLElBQUk7QUFDbkMsVUFBTSxZQUFZLDhCQUE4QixXQUFXLHdCQUF3QixpQkFBaUIsS0FBSyxVQUFVO0FBQ25ILFVBQU0sWUFBWSw4QkFBOEIsV0FBVyxpQkFBaUIsaUJBQWlCLEtBQUssVUFBVTtBQUM1RyxlQUFXLEdBQUcsU0FBUyxLQUFLLFNBQVM7QUFBQSxFQUN2QyxPQUFPO0FBQ0wsZ0JBQVk7QUFBQSxFQUNkO0FBRUEsU0FBTyxLQUFLLGNBQWMsUUFBUTtBQUNwQztBQUVBLG1CQUFjLFNBQUMsVUFBMkI7QUFDeEMsUUFBTSxPQUFPLFNBQVMsYUFBYSxXQUFXO0FBQzlDLFFBQU0sYUFBYSxTQUFTLGFBQWEsa0JBQWtCLEtBQUssU0FBUyxhQUFhLFdBQVc7QUFDakcsUUFBTSxVQUFVLFNBQVMsYUFBYSxlQUFlO0FBQ3JELFFBQU0sT0FBTyxTQUFTLGFBQWEsV0FBVztBQUM5QyxRQUFNLFdBQVcsU0FBUyxhQUFhLGVBQWU7QUFFdEQsUUFBTSxRQUFRLENBQUMsSUFBSTtBQUNuQixNQUFJLFdBQVksT0FBTSxLQUFLLFVBQVU7QUFDckMsTUFBSSxRQUFTLE9BQU0sS0FBSyxPQUFPO0FBQy9CLE1BQUksVUFBVTtBQUNaLFVBQU0sS0FBSyxRQUFRO0FBQUEsRUFDckIsV0FBVyxNQUFNO0FBQ2YsVUFBTSxLQUFLLElBQUk7QUFBQSxFQUNqQjtBQUVBLFNBQU8sTUFBTSxLQUFLLEdBQUc7QUFDdkI7QUFJTSwyQkFBc0IsaUJBQW9DO0FBQzlELE1BQUk7QUFDRixVQUFNLFdBQVcsTUFBTSxNQUFNLHVCQUF1QjtBQUNwRCxRQUFJLENBQUMsU0FBUyxJQUFJO0FBQ2hCLGNBQVEsS0FBSyw4REFBOEQ7QUFDM0UsYUFBTyxvQkFBSSxJQUFJO0FBQUEsSUFDakI7QUFFQSxVQUFNLFdBQVcsTUFBTSxTQUFTLEtBQUs7QUFDckMsdUJBQUssV0FBWTtBQUVqQixVQUFNLFdBQVcsb0JBQUksSUFBdUI7QUFFNUMsZUFBVyxVQUFVLFNBQVMsV0FBVyxDQUFDLEdBQUc7QUFDM0MsaUJBQVcsZUFBZSxPQUFPLGdCQUFnQixDQUFDLEdBQUc7QUFDbkQsWUFBSSxZQUFZLGlCQUFpQixZQUFZLFNBQVM7QUFDcEQsZ0JBQU0sVUFBVSxZQUFZO0FBQzVCLGdCQUFNLFNBQVMsWUFBWSxVQUFVLENBQUM7QUFFdEMsY0FBSSxPQUFPLFNBQVMsR0FBRztBQUNyQixrQkFBTSxhQUFhLElBQUksSUFBSSxPQUFPLElBQUksT0FBSyxFQUFFLElBQUksQ0FBQztBQUNsRCxxQkFBUyxJQUFJLFNBQVM7QUFBQSxjQUNwQjtBQUFBLGNBQ0E7QUFBQSxZQUNGLENBQUM7QUFBQSxVQUNIO0FBQUEsUUFDRjtBQUFBLE1BQ0Y7QUFBQSxJQUNGO0FBRUEsV0FBTztBQUFBLEVBQ1QsU0FBUyxPQUFPO0FBQ2QsWUFBUSxLQUFLLGtFQUFrRSxLQUFLO0FBQ3BGLFdBQU8sb0JBQUksSUFBSTtBQUFBLEVBQ2pCO0FBQ0Y7QUFFTSx1QkFBa0IsaUJBQUc7QUFDekIscUJBQUssa0JBQW1CLE1BQU0sc0JBQUsscURBQUw7QUFFOUIsTUFBSSxtQkFBSyxrQkFBaUIsU0FBUyxFQUFHO0FBRXRDLHdCQUFLLG9EQUFMO0FBQ0Esd0JBQUssa0RBQUw7QUFDQSx3QkFBSyxvREFBTDtBQUNGO0FBRUEsMEJBQXFCLFdBQUc7QUFDdEIsUUFBTSxPQUFPLEtBQUs7QUFDbEIsTUFBSSxDQUFDLEtBQU07QUFFWCxRQUFNLE9BQU8sS0FBSyxjQUFjO0FBRWhDLGFBQVcsQ0FBQyxTQUFTLFNBQVMsS0FBSyxtQkFBSyxtQkFBbUI7QUFDekQsVUFBTSxXQUFXLEtBQUssaUJBQWlCLE9BQU87QUFFOUMsZUFBVyxXQUFXLFVBQVU7QUFDOUIsaUJBQVcsYUFBYSxVQUFVLFlBQVk7QUFDNUMsZ0JBQVEsaUJBQWlCLFdBQVcsbUJBQUssc0JBQXFCLEVBQUUsU0FBUyxLQUFLLENBQUM7QUFBQSxNQUNqRjtBQUNBLE1BQUMsUUFBd0IsUUFBUSxvQkFBb0I7QUFDckQseUJBQUsscUJBQW9CLElBQUksT0FBTztBQUFBLElBQ3RDO0FBQUEsRUFDRjtBQUNGO0FBRUEsMEJBQXFCLFdBQUc7QUFDdEIsUUFBTSxPQUFPLEtBQUs7QUFDbEIsTUFBSSxDQUFDLEtBQU07QUFFWCxRQUFNLE9BQU8sS0FBSyxjQUFjO0FBRWhDLHFCQUFLLFdBQVUsUUFBUSxNQUFNO0FBQUEsSUFDM0IsV0FBVztBQUFBLElBQ1gsU0FBUztBQUFBLEVBQ1gsQ0FBQztBQUNIO0FBRUE7QUFhQSwyQkFBc0IsU0FBQyxlQUFtRDtBQUN4RSxNQUFJLENBQUMsZUFBZTtBQUNsQixXQUFPLEVBQUUsU0FBUyxNQUFNLGFBQWEsS0FBSztBQUFBLEVBQzVDO0FBRUEsTUFBSSxVQUFVLGNBQWMsV0FBVztBQUN2QyxNQUFJLGNBQWMsY0FBYyxlQUFlO0FBRS9DLE1BQUksY0FBYyxNQUFNLFFBQVEsbUJBQUssWUFBVztBQUM5QyxVQUFNLFdBQVcsY0FBYyxLQUFLO0FBQ3BDLFVBQU0sa0JBQWtCLHNCQUFLLG1EQUFMLFdBQTBCO0FBRWxELFFBQUksaUJBQWlCO0FBQ25CLFVBQUksQ0FBQyxXQUFXLGdCQUFnQixTQUFTO0FBQ3ZDLGtCQUFVLGdCQUFnQjtBQUFBLE1BQzVCLFdBQVcsZ0JBQWdCLFdBQVcsZ0JBQWdCLFlBQVksU0FBUztBQUN6RSxrQkFBVSxVQUFVLEdBQUcsT0FBTztBQUFBO0FBQUEsT0FBWSxRQUFRLEtBQUssZ0JBQWdCLE9BQU8sS0FBSyxnQkFBZ0I7QUFBQSxNQUNyRztBQUVBLFVBQUksQ0FBQyxlQUFlLGdCQUFnQixhQUFhO0FBQy9DLHNCQUFjLGdCQUFnQjtBQUFBLE1BQ2hDLFdBQVcsZ0JBQWdCLGVBQWUsZ0JBQWdCLGdCQUFnQixhQUFhO0FBQ3JGLHNCQUFjLGNBQWMsR0FBRyxXQUFXO0FBQUE7QUFBQSxFQUFPLGdCQUFnQixXQUFXLEtBQUssZ0JBQWdCO0FBQUEsTUFDbkc7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUVBLFNBQU8sRUFBRSxTQUFTLFlBQVk7QUFDaEM7QUFFQSx5QkFBb0IsU0FBQyxVQUFrQjtBQUNyQyxNQUFJLENBQUMsbUJBQUssV0FBVyxRQUFPO0FBRTVCLGFBQVcsVUFBVSxtQkFBSyxXQUFVLFdBQVcsQ0FBQyxHQUFHO0FBQ2pELGVBQVcsZUFBZSxPQUFPLGdCQUFnQixDQUFDLEdBQUc7QUFDbkQsVUFBSSxZQUFZLFNBQVMsYUFDcEIsWUFBWSxTQUFTLFdBQVcsWUFBWSxTQUFTLGNBQWM7QUFDdEUsZUFBTztBQUFBLE1BQ1Q7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUVBLFNBQU87QUFDVDtBQUVBLGtCQUFhLFNBQUMsT0FBYyxRQUFxQixTQUFpQixXQUFzQjtBQUN0RixRQUFNLGdCQUFnQixVQUFVLE9BQU8sS0FBSyxPQUFLLEVBQUUsU0FBUyxNQUFNLElBQUk7QUFFdEUsUUFBTSxZQUFZLHNCQUFLLHFEQUFMLFdBQTRCO0FBRTlDLFFBQU0sbUJBQW1CLHNCQUFLLHNEQUFMLFdBQTZCO0FBRXRELFFBQU0sY0FBMkI7QUFBQSxJQUMvQixJQUFJLEdBQUcsS0FBSyxJQUFJLENBQUMsSUFBSSxLQUFLLE9BQU8sQ0FBQztBQUFBLElBQ2xDLFdBQVcsb0JBQUksS0FBSztBQUFBLElBQ3BCLFdBQVcsTUFBTTtBQUFBLElBQ2pCO0FBQUEsSUFDQSxXQUFXLE9BQU8sTUFBTTtBQUFBLElBQ3hCLGNBQWMsT0FBTyxhQUFhO0FBQUEsSUFDbEM7QUFBQSxJQUNBLGNBQWMsZUFBZSxNQUFNLFFBQVE7QUFBQSxJQUMzQyxTQUFTLFVBQVU7QUFBQSxJQUNuQixhQUFhLFVBQVU7QUFBQSxJQUN2QixTQUFTLE1BQU07QUFBQSxJQUNmLFVBQVUsTUFBTTtBQUFBLElBQ2hCLFlBQVksTUFBTTtBQUFBLElBQ2xCLGtCQUFrQixNQUFNO0FBQUEsRUFDMUI7QUFFQSxxQkFBSyxpQkFBZ0IsS0FBSyxXQUFXO0FBRXJDLE1BQUksbUJBQUssaUJBQWdCLFNBQVMsbUJBQUsscUJBQW9CO0FBQ3pELHVCQUFLLGlCQUFnQixNQUFNO0FBQUEsRUFDN0I7QUFFQSx3QkFBSywyQ0FBTCxXQUFrQjtBQUNwQjtBQUVBLDRCQUF1QixTQUFDLE9BQXVDO0FBQzdELFFBQU0sYUFBc0MsQ0FBQztBQUM3QyxRQUFNLHFCQUFxQixJQUFJLElBQUksT0FBTyxvQkFBb0IsTUFBTSxTQUFTLENBQUM7QUFFOUUsUUFBTSxpQkFBaUIsQ0FBQyxVQUE0QjtBQUNsRCxRQUFJO0FBQ0YsYUFBTyxLQUFLLE1BQU0sS0FBSyxVQUFVLEtBQUssQ0FBQztBQUFBLElBQ3pDLFNBQVMsR0FBRztBQUNWLFVBQUk7QUFDRixlQUFPLE9BQU8sS0FBSztBQUFBLE1BQ3JCLFNBQVMsV0FBVztBQUNsQixlQUFPO0FBQUEsTUFDVDtBQUFBLElBQ0Y7QUFBQSxFQUNGO0FBRUEsTUFBSSxpQkFBaUIsZUFBZSxNQUFNLFdBQVcsUUFBVztBQUM5RCxlQUFXLFNBQVMsZUFBZSxNQUFNLE1BQU07QUFBQSxFQUNqRDtBQUVBLGFBQVcsT0FBTyxPQUFPLG9CQUFvQixLQUFLLEdBQUc7QUFDbkQsUUFBSSxDQUFDLG1CQUFtQixJQUFJLEdBQUcsS0FBSyxDQUFDLElBQUksV0FBVyxHQUFHLEtBQUssQ0FBQyxXQUFXLGVBQWUsR0FBRyxHQUFHO0FBQzNGLGlCQUFXLEdBQUcsSUFBSSxlQUFnQixNQUFjLEdBQUcsQ0FBQztBQUFBLElBQ3REO0FBQUEsRUFDRjtBQUVBLFNBQU87QUFDVDtBQUVBLGlCQUFZLFNBQUMsYUFBMEI7QUFDckMsTUFBSSxDQUFDLG1CQUFLLFlBQVk7QUFFdEIsUUFBTSxXQUFXLDhCQUFlLHFCQUFvQixRQUFRLFVBQVUsSUFBSTtBQUUxRSxRQUFNLE9BQU8sWUFBWSxVQUFVLG1CQUFtQjtBQUV0RCxRQUFNLFlBQVksU0FBUyxjQUFjLDBCQUEwQjtBQUNuRSxZQUFVLFFBQVEsVUFBVSxZQUFZO0FBQ3hDLFlBQVUsUUFBUSxZQUFZLFlBQVk7QUFDMUMsWUFBVSxRQUFRLGNBQWMsWUFBWTtBQUU1QyxRQUFNLFlBQVksc0JBQUssc0RBQUwsV0FBNkI7QUFDL0MsUUFBTSxZQUFZLG1CQUFLLG1CQUFrQixTQUFTLEtBQUssbUJBQUssbUJBQWtCLElBQUksWUFBWSxTQUFTO0FBQ3ZHLFFBQU0sZUFBZSxtQkFBSyxpQkFBZ0IsU0FBUyxLQUFLLG1CQUFLLGlCQUFnQixJQUFJLFlBQVksT0FBTztBQUVwRyxNQUFJLEVBQUUsYUFBYSxhQUFhLGVBQWU7QUFDN0MsY0FBVSxhQUFhLFVBQVUsRUFBRTtBQUFBLEVBQ3JDO0FBRUEsUUFBTSxRQUFRLFNBQVMsY0FBYyxzQkFBc0I7QUFDM0QsUUFBTSxjQUFjLFlBQVk7QUFDaEMsUUFBTSxhQUFhLFVBQVUsTUFBTTtBQUVuQyxRQUFNLFNBQVMsU0FBUyxjQUFjLHFCQUFxQjtBQUMzRCxTQUFPLGFBQWEsWUFBWSxZQUFZLFVBQVUsWUFBWSxDQUFDO0FBQ25FLFNBQU8sY0FBYztBQUVyQixRQUFNLFlBQVksU0FBUyxjQUFjLHdCQUF3QjtBQUNqRSxNQUFJLGNBQWMsSUFBSSxZQUFZLE9BQU87QUFDekMsTUFBSSxZQUFZLFdBQVc7QUFDekIsbUJBQWUsSUFBSSxZQUFZLFNBQVM7QUFBQSxFQUMxQztBQUNBLFlBQVUsY0FBYztBQUV4QixxQkFBSyxZQUFXLE9BQU8sUUFBUTtBQUUvQixNQUFJLENBQUMsbUJBQUssbUJBQWtCO0FBQzFCLDBCQUFLLDJDQUFMLFdBQWtCLFlBQVk7QUFBQSxFQUNoQztBQUVBLE1BQUksbUJBQUssZ0JBQWUsc0JBQUssaURBQUwsWUFBMkI7QUFDakQsMEJBQUssb0RBQUw7QUFBQSxFQUNGO0FBQ0Y7QUFFQSxpQkFBWSxTQUFDLFNBQWlCO0FBQzVCLFFBQU0sY0FBYyxzQkFBSyxrREFBTCxXQUF5QjtBQUM3QyxNQUFJLENBQUMsWUFBYTtBQUVsQixxQkFBSyxrQkFBbUI7QUFFeEIsUUFBTSxXQUFXLG1CQUFLLGFBQVksaUJBQWlCLGtCQUFrQjtBQUNyRSxZQUFVLFFBQVEsVUFBUTtBQUN4QixRQUFLLEtBQXFCLFFBQVEsWUFBWSxTQUFTO0FBQ3JELFdBQUssVUFBVSxJQUFJLFVBQVU7QUFDN0IsV0FBSyxhQUFhLGlCQUFpQixNQUFNO0FBQUEsSUFDM0MsT0FBTztBQUNMLFdBQUssVUFBVSxPQUFPLFVBQVU7QUFDaEMsV0FBSyxhQUFhLGlCQUFpQixPQUFPO0FBQUEsSUFDNUM7QUFBQSxFQUNGLENBQUM7QUFFRCxNQUFJLG1CQUFLLHFCQUFvQjtBQUMzQix1QkFBSyxvQkFBbUIsWUFBWTtBQUVwQyxVQUFNLGdCQUFnQixTQUFTLGNBQWMsS0FBSztBQUNsRCxrQkFBYyxZQUFZO0FBRTFCLFVBQU0sWUFBWSxTQUFTLGNBQWMsSUFBSTtBQUM3QyxjQUFVLGNBQWMsWUFBWTtBQUNwQyxjQUFVLFlBQVk7QUFDdEIsa0JBQWMsWUFBWSxTQUFTO0FBRW5DLFFBQUksWUFBWSxTQUFTO0FBQ3ZCLFlBQU0sVUFBVSxTQUFTLGNBQWMsR0FBRztBQUMxQyxjQUFRLGNBQWMsWUFBWTtBQUNsQyxjQUFRLFlBQVk7QUFDcEIsb0JBQWMsWUFBWSxPQUFPO0FBQUEsSUFDbkM7QUFFQSxRQUFJLFlBQVksYUFBYTtBQUMzQixZQUFNLGNBQWMsU0FBUyxjQUFjLEdBQUc7QUFDOUMsa0JBQVksY0FBYyxZQUFZO0FBQ3RDLGtCQUFZLFlBQVk7QUFDeEIsb0JBQWMsWUFBWSxXQUFXO0FBQUEsSUFDdkM7QUFFQSxVQUFNLE9BQU8sU0FBUyxjQUFjLEtBQUs7QUFDekMsU0FBSyxZQUFZO0FBRWpCLFVBQU0sU0FBUyxTQUFTLGNBQWMsTUFBTTtBQUM1QyxXQUFPLGFBQWEsWUFBWSxZQUFZLFVBQVUsWUFBWSxDQUFDO0FBQ25FLFdBQU8sY0FBYyxZQUFZLFVBQVUsbUJBQW1CO0FBQzlELFdBQU8sWUFBWTtBQUVuQixVQUFNLFVBQVUsU0FBUyxjQUFjLE1BQU07QUFDN0MsUUFBSSxjQUFjLElBQUksWUFBWSxPQUFPO0FBQ3pDLFFBQUksWUFBWSxXQUFXO0FBQ3pCLHFCQUFlLElBQUksWUFBWSxTQUFTO0FBQUEsSUFDMUM7QUFDQSxZQUFRLGNBQWM7QUFDdEIsWUFBUSxZQUFZO0FBRXBCLFNBQUssWUFBWSxNQUFNO0FBQ3ZCLFNBQUssWUFBWSxPQUFPO0FBRXhCLGtCQUFjLFlBQVksSUFBSTtBQUU5Qix1QkFBSyxvQkFBbUIsWUFBWSxhQUFhO0FBQUEsRUFDbkQ7QUFFQSxNQUFJLG1CQUFLLG1CQUFrQjtBQUN6Qix1QkFBSyxrQkFBaUIsWUFBWTtBQUVsQyxVQUFNLG9CQUFvQixTQUFTLGNBQWMsSUFBSTtBQUNyRCxzQkFBa0IsY0FBYztBQUNoQyxzQkFBa0IsWUFBWTtBQUU5QixVQUFNLHNCQUFzQixTQUFTLGNBQWMsS0FBSztBQUN4RCx3QkFBb0IsWUFBWTtBQUVoQyxVQUFNLGtCQUFrQixzQkFBSyx5REFBTCxXQUFnQztBQUN4RCxRQUFJLE9BQU8sS0FBSyxlQUFlLEVBQUUsU0FBUyxHQUFHO0FBQzNDLDBCQUFvQixZQUFZLHNCQUFLLGlEQUFMLFdBQXdCLGdCQUFnQjtBQUFBLElBQzFFLE9BQU87QUFDTCwwQkFBb0IsY0FBYztBQUFBLElBQ3BDO0FBRUEsdUJBQUssa0JBQWlCLFlBQVksaUJBQWlCO0FBQ25ELHVCQUFLLGtCQUFpQixZQUFZLG1CQUFtQjtBQUFBLEVBQ3ZEO0FBQ0Y7QUFFQSwrQkFBMEIsU0FBQyxhQUFtRDtBQUM1RSxRQUFNLGFBQXNDLENBQUM7QUFFN0MsTUFBSSxZQUFZLGtCQUFrQjtBQUNoQyxXQUFPLE9BQU8sWUFBWSxZQUFZLGdCQUFnQjtBQUFBLEVBQ3hEO0FBRUEsYUFBVyxVQUFVLFlBQVk7QUFDakMsYUFBVyxhQUFhLFlBQVk7QUFDcEMsYUFBVyxtQkFBbUIsWUFBWTtBQUMxQyxhQUFXLFdBQVcsWUFBWTtBQUVsQyxNQUFJLFlBQVksY0FBYztBQUM1QixlQUFXLE9BQU8sWUFBWTtBQUFBLEVBQ2hDO0FBRUEsU0FBTztBQUNUO0FBRUEsdUJBQWtCLFNBQUMsS0FBOEIsUUFBUSxHQUFxQjtBQUM1RSxRQUFNLEtBQUssU0FBUyxjQUFjLElBQUk7QUFDdEMsS0FBRyxZQUFZO0FBQ2YsTUFBSSxRQUFRLEdBQUc7QUFDYixPQUFHLFVBQVUsSUFBSSxRQUFRO0FBQUEsRUFDM0I7QUFFQSxhQUFXLENBQUMsS0FBSyxLQUFLLEtBQUssT0FBTyxRQUFRLEdBQUcsR0FBRztBQUM5QyxVQUFNLEtBQUssU0FBUyxjQUFjLElBQUk7QUFDdEMsT0FBRyxZQUFZO0FBRWYsVUFBTSxVQUFVLFNBQVMsY0FBYyxNQUFNO0FBQzdDLFlBQVEsWUFBWTtBQUNwQixZQUFRLGNBQWM7QUFFdEIsVUFBTSxZQUFZLFNBQVMsY0FBYyxNQUFNO0FBQy9DLGNBQVUsWUFBWTtBQUN0QixjQUFVLGNBQWM7QUFFeEIsT0FBRyxZQUFZLE9BQU87QUFDdEIsT0FBRyxZQUFZLFNBQVM7QUFFeEIsUUFBSSxVQUFVLFFBQVEsVUFBVSxRQUFXO0FBQ3pDLFlBQU0sWUFBWSxTQUFTLGNBQWMsTUFBTTtBQUMvQyxnQkFBVSxZQUFZO0FBQ3RCLGdCQUFVLGNBQWMsT0FBTyxLQUFLO0FBQ3BDLFNBQUcsWUFBWSxTQUFTO0FBQUEsSUFDMUIsV0FBVyxPQUFPLFVBQVUsV0FBVztBQUNyQyxZQUFNLFlBQVksU0FBUyxjQUFjLE1BQU07QUFDL0MsZ0JBQVUsWUFBWTtBQUN0QixnQkFBVSxjQUFjLE9BQU8sS0FBSztBQUNwQyxTQUFHLFlBQVksU0FBUztBQUFBLElBQzFCLFdBQVcsT0FBTyxVQUFVLFVBQVU7QUFDcEMsWUFBTSxZQUFZLFNBQVMsY0FBYyxNQUFNO0FBQy9DLGdCQUFVLFlBQVk7QUFDdEIsZ0JBQVUsY0FBYyxPQUFPLEtBQUs7QUFDcEMsU0FBRyxZQUFZLFNBQVM7QUFBQSxJQUMxQixXQUFXLE9BQU8sVUFBVSxVQUFVO0FBQ3BDLFlBQU0sWUFBWSxTQUFTLGNBQWMsTUFBTTtBQUMvQyxnQkFBVSxZQUFZO0FBQ3RCLGdCQUFVLGNBQWMsSUFBSSxLQUFLO0FBQ2pDLFNBQUcsWUFBWSxTQUFTO0FBQUEsSUFDMUIsV0FBVyxNQUFNLFFBQVEsS0FBSyxHQUFHO0FBQy9CLFlBQU0sWUFBWSxTQUFTLGNBQWMsTUFBTTtBQUMvQyxnQkFBVSxZQUFZO0FBQ3RCLGdCQUFVLGNBQWMsU0FBUyxNQUFNLE1BQU07QUFDN0MsU0FBRyxZQUFZLFNBQVM7QUFFeEIsVUFBSSxNQUFNLFNBQVMsS0FBSyxRQUFRLEdBQUc7QUFDakMsY0FBTSxZQUFxQyxDQUFDO0FBQzVDLGNBQU0sUUFBUSxDQUFDLE1BQU0sVUFBVTtBQUM3QixvQkFBVSxLQUFLLElBQUk7QUFBQSxRQUNyQixDQUFDO0FBQ0QsV0FBRyxZQUFZLHNCQUFLLGlEQUFMLFdBQXdCLFdBQVcsUUFBUSxFQUFFO0FBQUEsTUFDOUQ7QUFBQSxJQUNGLFdBQVcsT0FBTyxVQUFVLFVBQVU7QUFDcEMsWUFBTSxZQUFZLFNBQVMsY0FBYyxNQUFNO0FBQy9DLGdCQUFVLFlBQVk7QUFDdEIsWUFBTSxPQUFPLE9BQU8sS0FBSyxLQUFnQztBQUN6RCxnQkFBVSxjQUFjLEtBQUssU0FBUyxJQUFJLFdBQVc7QUFDckQsU0FBRyxZQUFZLFNBQVM7QUFFeEIsVUFBSSxLQUFLLFNBQVMsS0FBSyxRQUFRLEdBQUc7QUFDaEMsV0FBRyxZQUFZLHNCQUFLLGlEQUFMLFdBQXdCLE9BQWtDLFFBQVEsRUFBRTtBQUFBLE1BQ3JGO0FBQUEsSUFDRixPQUFPO0FBQ0wsWUFBTSxZQUFZLFNBQVMsY0FBYyxNQUFNO0FBQy9DLGdCQUFVLFlBQVk7QUFDdEIsZ0JBQVUsY0FBYyxPQUFPLEtBQUs7QUFDcEMsU0FBRyxZQUFZLFNBQVM7QUFBQSxJQUMxQjtBQUVBLE9BQUcsWUFBWSxFQUFFO0FBQUEsRUFDbkI7QUFFQSxTQUFPO0FBQ1Q7QUFFQSwwQkFBcUIsV0FBRztBQUN0QixNQUFJLENBQUMsbUJBQUssWUFBWTtBQUV0Qix3QkFBc0IsTUFBTTtBQUMxQixVQUFNLFlBQVksbUJBQUssWUFBWTtBQUNuQyxRQUFJLFdBQVc7QUFDYixnQkFBVSxlQUFlLEVBQUUsVUFBVSxRQUFRLE9BQU8sTUFBTSxDQUFDO0FBQUEsSUFDN0Q7QUFBQSxFQUNGLENBQUM7QUFDSDtBQUVBLHVCQUFrQixXQUFZO0FBQzVCLFFBQU0sT0FBTyxLQUFLLFlBQVksY0FBYyxZQUFZO0FBQ3hELE1BQUksQ0FBQyxLQUFNLFFBQU87QUFFbEIsUUFBTSxnQkFBZ0IsU0FBUyxLQUFLLGFBQWEsVUFBVSxLQUFLLEtBQUssRUFBRTtBQUN2RSxTQUFPLGtCQUFrQjtBQUMzQjtBQUVBLGtCQUFhLFNBQUMsT0FBZTtBQUMzQixxQkFBSyxvQkFBcUIsTUFBTSxZQUFZO0FBRTVDLE1BQUksQ0FBQyxtQkFBSyxZQUFZO0FBRXRCLGFBQVcsU0FBUyxtQkFBSyxZQUFXLFVBQVU7QUFDNUMsVUFBTSxjQUFjLHNCQUFLLGtEQUFMLFdBQTBCLE1BQXNCLFFBQVE7QUFFNUUsUUFBSSxDQUFDLFlBQWE7QUFFbEIsVUFBTSxZQUFZLHNCQUFLLHNEQUFMLFdBQTZCO0FBQy9DLFVBQU0sWUFBWSxtQkFBSyxtQkFBa0IsU0FBUyxLQUFLLG1CQUFLLG1CQUFrQixJQUFJLFlBQVksU0FBUztBQUN2RyxVQUFNLGVBQWUsbUJBQUssaUJBQWdCLFNBQVMsS0FBSyxtQkFBSyxpQkFBZ0IsSUFBSSxZQUFZLE9BQU87QUFFcEcsSUFBQyxNQUFzQixTQUFTLEVBQUUsYUFBYSxhQUFhO0FBQUEsRUFDOUQ7QUFDRjtBQUVBLDRCQUF1QixTQUFDLGFBQW1DO0FBQ3pELE1BQUksQ0FBQyxtQkFBSyxvQkFBb0IsUUFBTztBQUVyQyxRQUFNLGFBQWE7QUFBQSxJQUNqQixZQUFZO0FBQUEsSUFDWixZQUFZO0FBQUEsSUFDWixZQUFZLGFBQWE7QUFBQSxJQUN6QixLQUFLLFVBQVUsWUFBWSxvQkFBb0IsQ0FBQyxDQUFDO0FBQUEsRUFDbkQsRUFBRSxLQUFLLEdBQUcsRUFBRSxZQUFZO0FBRXhCLFNBQU8sV0FBVyxTQUFTLG1CQUFLLG1CQUFrQjtBQUNwRDtBQUVBLHdCQUFtQixTQUFDLElBQXFDO0FBQ3ZELFNBQU8sbUJBQUssaUJBQWdCLEtBQUssT0FBSyxFQUFFLE9BQU8sRUFBRTtBQUNuRDtBQUVBLHdCQUFtQixXQUFHO0FBQ3BCLFFBQU0sbUJBQW1CLHNCQUFLLDJEQUFMO0FBRXpCLFFBQU0sa0JBQWtCLHNCQUFLLGlDQUFMLFdBQVE7QUFDaEMsTUFBSSxtQkFBbUIsbUJBQUssbUJBQWtCO0FBQzVDLFFBQUksT0FBTyxnQkFBZ0IsY0FBYyxZQUFZO0FBQ3JELFFBQUksQ0FBQyxNQUFNO0FBQ1QsYUFBTyxTQUFTLGNBQWMsWUFBWTtBQUMxQyxzQkFBZ0IsWUFBWSxJQUFJO0FBQUEsSUFDbEM7QUFFQSxVQUFNLGdCQUFnQixLQUFLLGlCQUFpQixpQkFBaUI7QUFDN0Qsa0JBQWMsUUFBUSxVQUFRLEtBQUssT0FBTyxDQUFDO0FBRTNDLFVBQU0sZ0JBQWdCLG9CQUFJLElBQVk7QUFDdEMsZUFBVyxDQUFDLFNBQVMsU0FBUyxLQUFLLG1CQUFLLG1CQUFrQjtBQUN4RCxVQUFJLG1CQUFLLHFCQUFvQixJQUFJLE9BQU8sR0FBRztBQUN6QyxtQkFBVyxhQUFhLFVBQVUsWUFBWTtBQUM1Qyx3QkFBYyxJQUFJLFNBQVM7QUFBQSxRQUM3QjtBQUFBLE1BQ0Y7QUFBQSxJQUNGO0FBRUEsUUFBSSxpQkFBaUIsWUFBWTtBQUMvQix5QkFBSyxtQkFBcUIsaUJBQWlCLFdBQW1GLGFBQWEsYUFBYTtBQUFBLElBQzFKLE9BQU87QUFDTCx5QkFBSyxtQkFBb0IsSUFBSSxJQUFJLGFBQWE7QUFBQSxJQUNoRDtBQUVBLGVBQVcsYUFBYSxlQUFlO0FBQ3JDLFlBQU0sT0FBTyxTQUFTLGNBQWMsaUJBQWlCO0FBQ3JELFdBQUssYUFBYSxXQUFXLFVBQVU7QUFDdkMsV0FBSyxhQUFhLFNBQVMsU0FBUztBQUNwQyxVQUFJLG1CQUFLLG1CQUFrQixJQUFJLFNBQVMsR0FBRztBQUN6QyxhQUFLLGFBQWEsV0FBVyxFQUFFO0FBQUEsTUFDakM7QUFDQSxXQUFLLGNBQWM7QUFDbkIsV0FBSyxZQUFZLElBQUk7QUFBQSxJQUN2QjtBQUFBLEVBQ0Y7QUFFQSxRQUFNLGdCQUFnQixzQkFBSyxpQ0FBTCxXQUFRO0FBQzlCLE1BQUksaUJBQWlCLG1CQUFLLG1CQUFrQjtBQUMxQyxRQUFJLE9BQU8sY0FBYyxjQUFjLFlBQVk7QUFDbkQsUUFBSSxDQUFDLE1BQU07QUFDVCxhQUFPLFNBQVMsY0FBYyxZQUFZO0FBQzFDLG9CQUFjLFlBQVksSUFBSTtBQUFBLElBQ2hDO0FBRUEsVUFBTSxnQkFBZ0IsS0FBSyxpQkFBaUIsaUJBQWlCO0FBQzdELGtCQUFjLFFBQVEsVUFBUSxLQUFLLE9BQU8sQ0FBQztBQUUzQyxVQUFNLGNBQWMsb0JBQUksSUFBWTtBQUNwQyxlQUFXLFdBQVcsbUJBQUssa0JBQWlCLEtBQUssR0FBRztBQUNsRCxVQUFJLG1CQUFLLHFCQUFvQixJQUFJLE9BQU8sR0FBRztBQUN6QyxvQkFBWSxJQUFJLE9BQU87QUFBQSxNQUN6QjtBQUFBLElBQ0Y7QUFFQSxRQUFJLGlCQUFpQixVQUFVO0FBQzdCLHlCQUFLLGlCQUFtQixpQkFBaUIsU0FBaUYsYUFBYSxXQUFXO0FBQUEsSUFDcEosT0FBTztBQUNMLHlCQUFLLGlCQUFrQixJQUFJLElBQUksV0FBVztBQUFBLElBQzVDO0FBRUEsZUFBVyxXQUFXLGFBQWE7QUFDakMsWUFBTSxPQUFPLFNBQVMsY0FBYyxpQkFBaUI7QUFDckQsV0FBSyxhQUFhLFdBQVcsVUFBVTtBQUN2QyxXQUFLLGFBQWEsU0FBUyxPQUFPO0FBQ2xDLFVBQUksbUJBQUssaUJBQWdCLElBQUksT0FBTyxHQUFHO0FBQ3JDLGFBQUssYUFBYSxXQUFXLEVBQUU7QUFBQSxNQUNqQztBQUNBLFdBQUssY0FBYyxJQUFJLE9BQU87QUFDOUIsV0FBSyxZQUFZLElBQUk7QUFBQSxJQUN2QjtBQUFBLEVBQ0Y7QUFDRjtBQUVBO0FBZUE7QUFlQSxpQ0FBNEIsV0FBcUU7QUFDL0YsUUFBTSxjQUFnRjtBQUFBLElBQ3BGLFlBQVk7QUFBQSxJQUNaLFVBQVU7QUFBQSxFQUNaO0FBRUEsTUFBSTtBQUNGLFVBQU0sa0JBQWtCLGFBQWEsUUFBUSw4QkFBOEI7QUFDM0UsUUFBSSxpQkFBaUI7QUFDbkIsa0JBQVksYUFBYSxJQUFJLElBQUksS0FBSyxNQUFNLGVBQWUsQ0FBQztBQUFBLElBQzlEO0FBRUEsVUFBTSxnQkFBZ0IsYUFBYSxRQUFRLDJCQUEyQjtBQUN0RSxRQUFJLGVBQWU7QUFDakIsa0JBQVksV0FBVyxJQUFJLElBQUksS0FBSyxNQUFNLGFBQWEsQ0FBQztBQUFBLElBQzFEO0FBQUEsRUFDRixTQUFTLEdBQUc7QUFDVixZQUFRLE1BQU0sK0RBQStEO0FBQUEsRUFDL0U7QUFFQSxTQUFPO0FBQ1Q7QUFFQSxzQkFBaUIsV0FBRztBQUNsQixNQUFJO0FBQ0YsaUJBQWE7QUFBQSxNQUFRO0FBQUEsTUFDbkIsS0FBSyxVQUFVLENBQUMsR0FBRyxtQkFBSyxrQkFBaUIsQ0FBQztBQUFBLElBQUM7QUFDN0MsaUJBQWE7QUFBQSxNQUFRO0FBQUEsTUFDbkIsS0FBSyxVQUFVLENBQUMsR0FBRyxtQkFBSyxnQkFBZSxDQUFDO0FBQUEsSUFBQztBQUFBLEVBQzdDLFNBQVMsR0FBRztBQUFBLEVBRVo7QUFDRjtBQUVBLGlCQUFZLFdBQUc7QUFDYixxQkFBSyxpQkFBa0IsQ0FBQztBQUN4QixxQkFBSyxrQkFBbUI7QUFDeEIsTUFBSSxtQkFBSyxhQUFZO0FBQ25CLHVCQUFLLFlBQVcsZ0JBQWdCO0FBQUEsRUFDbEM7QUFDQSxNQUFJLG1CQUFLLHFCQUFvQjtBQUMzQix1QkFBSyxvQkFBbUIsWUFBWTtBQUFBLEVBQ3RDO0FBQ0EsTUFBSSxtQkFBSyxtQkFBa0I7QUFDekIsdUJBQUssa0JBQWlCLFlBQVk7QUFBQSxFQUNwQztBQUNGO0FBRU0sZ0JBQVcsaUJBQUc7QUFDbEIsTUFBSSxDQUFDLG1CQUFLLFlBQVk7QUFFdEIsUUFBTSxnQkFBZ0IsTUFBTSxLQUFLLG1CQUFLLFlBQVcsUUFBUSxFQUN0RCxPQUFPLFdBQVMsQ0FBRSxNQUFzQixNQUFNLEVBQzlDLElBQUksV0FBUztBQUNaLFVBQU0sS0FBTSxNQUFzQixRQUFRO0FBQzFDLFdBQU8sc0JBQUssa0RBQUwsV0FBeUI7QUFBQSxFQUNsQyxDQUFDLEVBQ0EsT0FBTyxDQUFDLFVBQWdDLENBQUMsQ0FBQyxLQUFLLEVBQy9DLElBQUksV0FBUztBQUNaLFVBQU0sT0FBTyxNQUFNLFVBQVUsbUJBQW1CO0FBQ2hELFVBQU0sVUFBVSxNQUFNLFlBQVksSUFBSSxNQUFNLFNBQVMsS0FBSyxNQUFNO0FBQ2hFLFVBQU0sUUFBUSxNQUFNLG9CQUFvQixPQUFPLEtBQUssTUFBTSxnQkFBZ0IsRUFBRSxTQUFTLElBQ2pGLE1BQU0sS0FBSyxVQUFVLE1BQU0sZ0JBQWdCLENBQUMsS0FDNUM7QUFDSixXQUFPLElBQUksSUFBSSxNQUFNLE1BQU0sT0FBTyxLQUFLLE9BQU8sV0FBVyxNQUFNLFNBQVMsR0FBRyxLQUFLO0FBQUEsRUFDbEYsQ0FBQyxFQUNBLEtBQUssSUFBSTtBQUVaLE1BQUksQ0FBQyxjQUFlO0FBRXBCLE1BQUk7QUFDRixVQUFNLFVBQVUsVUFBVSxVQUFVLGFBQWE7QUFDakQsVUFBTSxNQUFNLHNCQUFLLGlDQUFMLFdBQVE7QUFDcEIsUUFBSSxLQUFLO0FBQ1AsWUFBTSxXQUFXLE1BQU0sS0FBSyxJQUFJLFVBQVUsRUFBRTtBQUFBLFFBQzFDLE9BQUssRUFBRSxhQUFhLEtBQUssY0FBYyxFQUFFLGFBQWEsS0FBSyxFQUFFLFVBQVUsS0FBSztBQUFBLE1BQzlFO0FBQ0EsVUFBSSxVQUFVO0FBQ1osY0FBTSxXQUFXLFNBQVM7QUFDMUIsaUJBQVMsY0FBYztBQUV2QixZQUFJLG1CQUFLLDZCQUE0QjtBQUNuQyx1QkFBYSxtQkFBSywyQkFBMEI7QUFBQSxRQUM5QztBQUVBLDJCQUFLLDRCQUE2QixXQUFXLE1BQU07QUFDakQsY0FBSSxLQUFLLGVBQWUsU0FBUyxZQUFZO0FBQzNDLHFCQUFTLGNBQWM7QUFBQSxVQUN6QjtBQUNBLDZCQUFLLDRCQUE2QjtBQUFBLFFBQ3BDLEdBQUcsR0FBSTtBQUFBLE1BQ1Q7QUFBQSxJQUNGO0FBQUEsRUFDRixTQUFTLEtBQUs7QUFDWixZQUFRLE1BQU0sNkNBQTZDLEdBQUc7QUFBQSxFQUNoRTtBQUNGO0FBRUEseUJBQW9CLFdBQUc7QUFDckIscUJBQUssWUFBYSxzQkFBSyxpQ0FBTCxXQUFRO0FBQzFCLHFCQUFLLG9CQUFxQixzQkFBSyxpQ0FBTCxXQUFRO0FBQ2xDLHFCQUFLLGtCQUFtQixzQkFBSyxpQ0FBTCxXQUFRO0FBRWhDLE1BQUksbUJBQUssYUFBWTtBQUNuQix1QkFBSyxZQUFXLGlCQUFpQixTQUFTLENBQUMsTUFBYTtBQUN0RCxZQUFNLFdBQVksRUFBRSxPQUFtQixRQUFRLGtCQUFrQjtBQUNqRSxVQUFJLFVBQVU7QUFDWixjQUFNLFVBQVUsU0FBUyxRQUFRO0FBQ2pDLFlBQUksU0FBUztBQUNYLGdDQUFLLDJDQUFMLFdBQWtCO0FBQUEsUUFDcEI7QUFBQSxNQUNGO0FBQUEsSUFDRixDQUFDO0FBQUEsRUFDSDtBQUVBLFFBQU0sZUFBZSxzQkFBSyxpQ0FBTCxXQUFRO0FBQzdCLE1BQUksY0FBYztBQUNoQixpQkFBYSxpQkFBaUIsU0FBUyxDQUFDLE1BQWE7QUFDbkQsWUFBTSxFQUFFLFFBQVEsR0FBRyxJQUFJLEVBQUU7QUFDekIsbUJBQWEsbUJBQUssMkJBQTJCO0FBQzdDLHlCQUFLLDRCQUE2QixXQUFXLE1BQU07QUFDakQsOEJBQUssNENBQUwsV0FBbUI7QUFBQSxNQUNyQixHQUFHLEdBQUc7QUFBQSxJQUNSLENBQUM7QUFBQSxFQUNIO0FBRUEsUUFBTSxrQkFBa0Isc0JBQUssaUNBQUwsV0FBUTtBQUNoQyxNQUFJLGlCQUFpQjtBQUNuQixvQkFBZ0IsaUJBQWlCLFVBQVUsbUJBQUssNkJBQTZDO0FBQUEsRUFDL0Y7QUFFQSxRQUFNLGdCQUFnQixzQkFBSyxpQ0FBTCxXQUFRO0FBQzlCLE1BQUksZUFBZTtBQUNqQixrQkFBYyxpQkFBaUIsVUFBVSxtQkFBSywyQkFBMkM7QUFBQSxFQUMzRjtBQUVBLHdCQUFLLGlDQUFMLFdBQVEsaUJBQWlCLGlCQUFpQixTQUFTLE1BQU07QUFDdkQsMEJBQUssMkNBQUw7QUFBQSxFQUNGLENBQUM7QUFFRCx3QkFBSyxpQ0FBTCxXQUFRLGdCQUFnQixpQkFBaUIsU0FBUyxNQUFNO0FBQ3RELDBCQUFLLDBDQUFMO0FBQUEsRUFDRixDQUFDO0FBQ0g7QUFsZ0VBLDRCQUFTLGtCQURULHFCQWhGVyxpQkFpRkY7QUFHVCw0QkFBUyxhQURULGdCQW5GVyxpQkFvRkY7QUFHVCw0QkFBUyxlQURULGtCQXRGVyxpQkF1RkY7QUFHVCw0QkFBUyxnQkFEVCxtQkF6RlcsaUJBMEZGO0FBR1QsNEJBQVMsYUFEVCxnQkE1RlcsaUJBNkZGO0FBR1QsNEJBQVMsU0FEVCxZQS9GVyxpQkFnR0Y7QUFHVCw0QkFBUyxVQURULGFBbEdXLGlCQW1HRjtBQUdULDRCQUFTLGdCQURULG1CQXJHVyxpQkFzR0Y7QUFHVCw0QkFBUyxnQkFEVCxtQkF4R1csaUJBeUdGO0FBR1QsNEJBQVMsV0FEVCxjQTNHVyxpQkE0R0Y7QUFHVCw0QkFBUyxrQkFEVCxxQkE5R1csaUJBK0dGO0FBL0dFLGtCQUFOLDhDQURQLDRCQUNhO0FBQ1gsY0FEVyxpQkFDSixVQUFTO0FBQUE7QUFHaEIsYUFKVyxpQkFJSixvQkFBcUIsTUFBTTtBQUNoQyxRQUFNLElBQUksU0FBUyxjQUFjLFVBQVU7QUFDM0MsSUFBRSxZQUFZO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQXdCZCxTQUFPO0FBQ1QsR0FBRztBQUVILGFBakNXLGlCQWlDSixxQkFBc0IsTUFBTTtBQUNqQyxRQUFNLElBQUksU0FBUyxjQUFjLFVBQVU7QUFDM0MsSUFBRSxZQUFZO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBVWQsU0FBTztBQUNULEdBQUc7QUFFSCxhQWhEVyxpQkFnREosb0JBQXFCLE1BQU07QUFDaEMsUUFBTSxJQUFJLFNBQVMsY0FBYyxVQUFVO0FBQzNDLElBQUUsWUFBWTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBS2QsU0FBTztBQUNULEdBQUc7QUFFSCxhQTFEVyxpQkEwREosb0JBQXFCLE1BQU07QUFDaEMsUUFBTSxJQUFJLFNBQVMsY0FBYyxVQUFVO0FBQzNDLElBQUUsWUFBWTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFNZCxTQUFPO0FBQ1QsR0FBRztBQUVILGFBckVXLGlCQXFFSixzQkFBdUIsTUFBTTtBQUNsQyxRQUFNLElBQUksU0FBUyxjQUFjLFVBQVU7QUFDM0MsSUFBRSxZQUFZO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQU1kLFNBQU87QUFDVCxHQUFHO0FBOUVFLDRCQUFNO0FBQU4sSUFBTSxpQkFBTjsiLAogICJuYW1lcyI6IFtdCn0K
