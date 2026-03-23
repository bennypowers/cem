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
    super.connectedCallback();
    if (!__privateGet(this, _clientModulesLoaded)) {
      [{ CEMReloadClient }, { StatePersistence }] = await Promise.all([
        // @ts-ignore -- plain JS modules served at runtime by Go server
        import("/__cem/websocket-client.js"),
        // @ts-ignore
        import("/__cem/state-persistence.js")
      ]);
      import("/__cem/health-badges.js").catch((e) => console.error("[cem-serve] Failed to load health-badges:", e));
      __privateSet(this, _clientModulesLoaded, true);
      __privateMethod(this, _CemServeChrome_instances, initWsClient_fn).call(this);
      __privateMethod(this, _CemServeChrome_instances, migrateFromLocalStorageIfNeeded_fn).call(this);
    }
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcbmltcG9ydCB7IGlmRGVmaW5lZCB9IGZyb20gJ2xpdC9kaXJlY3RpdmVzL2lmLWRlZmluZWQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXNlcnZlLWNocm9tZS5jc3MnO1xuXG5pbXBvcnQgJy4uL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlLmpzJztcbmltcG9ydCAnLi4vY2VtLWRyYXdlci9jZW0tZHJhd2VyLmpzJztcbmltcG9ydCAnLi4vY2VtLWhlYWx0aC1wYW5lbC9jZW0taGVhbHRoLXBhbmVsLmpzJztcbmltcG9ydCAnLi4vY2VtLW1hbmlmZXN0LWJyb3dzZXIvY2VtLW1hbmlmZXN0LWJyb3dzZXIuanMnO1xuaW1wb3J0ICcuLi9jZW0tcmVjb25uZWN0aW9uLWNvbnRlbnQvY2VtLXJlY29ubmVjdGlvbi1jb250ZW50LmpzJztcbmltcG9ydCAnLi4vY2VtLXNlcnZlLWRlbW8vY2VtLXNlcnZlLWRlbW8uanMnO1xuaW1wb3J0ICcuLi9jZW0tc2VydmUta25vYi1ncm91cC9jZW0tc2VydmUta25vYi1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1zZXJ2ZS1rbm9icy9jZW0tc2VydmUta25vYnMuanMnO1xuaW1wb3J0ICcuLi9jZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXkvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtYWxlcnQvcGYtdjYtYWxlcnQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1hbGVydC1ncm91cC9wZi12Ni1hbGVydC1ncm91cC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWJ1dHRvbi9wZi12Ni1idXR0b24uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1jYXJkL3BmLXY2LWNhcmQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1iYWRnZS9wZi12Ni1iYWRnZS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWRyb3Bkb3duL3BmLXY2LWRyb3Bkb3duLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uL3BmLXY2LWV4cGFuZGFibGUtc2VjdGlvbi5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWxhYmVsL3BmLXY2LWxhYmVsLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbWFzdGhlYWQvcGYtdjYtbWFzdGhlYWQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1tb2RhbC9wZi12Ni1tb2RhbC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW5hdi1ncm91cC9wZi12Ni1uYXYtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXYtaXRlbS9wZi12Ni1uYXYtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW5hdi1saW5rL3BmLXY2LW5hdi1saW5rLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbmF2LWxpc3QvcGYtdjYtbmF2LWxpc3QuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXZpZ2F0aW9uL3BmLXY2LW5hdmlnYXRpb24uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1wYWdlLW1haW4vcGYtdjYtcGFnZS1tYWluLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtcGFnZS1zaWRlYmFyL3BmLXY2LXBhZ2Utc2lkZWJhci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXBhZ2UvcGYtdjYtcGFnZS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXBvcG92ZXIvcGYtdjYtcG9wb3Zlci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXNlbGVjdC9wZi12Ni1zZWxlY3QuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1za2lwLXRvLWNvbnRlbnQvcGYtdjYtc2tpcC10by1jb250ZW50LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtc3dpdGNoL3BmLXY2LXN3aXRjaC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRhYi9wZi12Ni10YWIuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10YWJzL3BmLXY2LXRhYnMuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10ZXh0LWlucHV0LWdyb3VwL3BmLXY2LXRleHQtaW5wdXQtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10ZXh0LWlucHV0L3BmLXY2LXRleHQtaW5wdXQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvZ2dsZS1ncm91cC9wZi12Ni10b2dnbGUtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10b29sYmFyLWdyb3VwL3BmLXY2LXRvb2xiYXItZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10b29sYmFyLWl0ZW0vcGYtdjYtdG9vbGJhci1pdGVtLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdG9vbGJhci9wZi12Ni10b29sYmFyLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdHJlZS1pdGVtL3BmLXY2LXRyZWUtaXRlbS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRyZWUtdmlldy9wZi12Ni10cmVlLXZpZXcuanMnO1xuXG4vLyBDbGllbnQtb25seSBtb2R1bGVzIGxvYWRlZCBkeW5hbWljYWxseSB0byBhdm9pZCBicmVha2luZyBTU1IuXG4vLyBUaGVzZSBhcmUgcGxhaW4gSlMgbW9kdWxlcyBzZXJ2ZWQgYXQgcnVudGltZSBieSB0aGUgR28gc2VydmVyLlxudHlwZSBDRU1SZWxvYWRDbGllbnRUeXBlID0geyBuZXcob3B0czogYW55KTogYW55IH07XG50eXBlIFN0YXRlUGVyc2lzdGVuY2VUeXBlID0ge1xuICBnZXRTdGF0ZSgpOiBhbnk7XG4gIHVwZGF0ZVN0YXRlKHM6IGFueSk6IHZvaWQ7XG4gIGdldFRyZWVTdGF0ZSgpOiBhbnk7XG4gIHNldFRyZWVTdGF0ZShzOiBhbnkpOiB2b2lkO1xuICBtaWdyYXRlRnJvbUxvY2FsU3RvcmFnZSgpOiB2b2lkO1xufTtcbmxldCBDRU1SZWxvYWRDbGllbnQ6IENFTVJlbG9hZENsaWVudFR5cGU7XG5sZXQgU3RhdGVQZXJzaXN0ZW5jZTogU3RhdGVQZXJzaXN0ZW5jZVR5cGU7XG5cbmludGVyZmFjZSBFdmVudEluZm8ge1xuICBldmVudE5hbWVzOiBTZXQ8c3RyaW5nPjtcbiAgZXZlbnRzOiBBcnJheTx7IG5hbWU6IHN0cmluZzsgdHlwZT86IHsgdGV4dDogc3RyaW5nIH07IHN1bW1hcnk/OiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH0+O1xufVxuXG5pbnRlcmZhY2UgRXZlbnRSZWNvcmQge1xuICBpZDogc3RyaW5nO1xuICB0aW1lc3RhbXA6IERhdGU7XG4gIGV2ZW50TmFtZTogc3RyaW5nO1xuICB0YWdOYW1lOiBzdHJpbmc7XG4gIGVsZW1lbnRJZDogc3RyaW5nIHwgbnVsbDtcbiAgZWxlbWVudENsYXNzOiBzdHJpbmcgfCBudWxsO1xuICBjdXN0b21Qcm9wZXJ0aWVzOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPjtcbiAgbWFuaWZlc3RUeXBlOiBzdHJpbmcgfCBudWxsO1xuICBzdW1tYXJ5OiBzdHJpbmcgfCBudWxsO1xuICBkZXNjcmlwdGlvbjogc3RyaW5nIHwgbnVsbDtcbiAgYnViYmxlczogYm9vbGVhbjtcbiAgY29tcG9zZWQ6IGJvb2xlYW47XG4gIGNhbmNlbGFibGU6IGJvb2xlYW47XG4gIGRlZmF1bHRQcmV2ZW50ZWQ6IGJvb2xlYW47XG59XG5cbmludGVyZmFjZSBEZWJ1Z0RhdGEge1xuICB2ZXJzaW9uPzogc3RyaW5nO1xuICBvcz86IHN0cmluZztcbiAgd2F0Y2hEaXI/OiBzdHJpbmc7XG4gIG1hbmlmZXN0U2l6ZT86IHN0cmluZztcbiAgZGVtb0NvdW50PzogbnVtYmVyO1xuICBkZW1vcz86IEFycmF5PHtcbiAgICB0YWdOYW1lOiBzdHJpbmc7XG4gICAgZGVzY3JpcHRpb24/OiBzdHJpbmc7XG4gICAgZmlsZXBhdGg/OiBzdHJpbmc7XG4gICAgY2Fub25pY2FsVVJMOiBzdHJpbmc7XG4gICAgbG9jYWxSb3V0ZTogc3RyaW5nO1xuICB9PjtcbiAgaW1wb3J0TWFwPzogUmVjb3JkPHN0cmluZywgdW5rbm93bj47XG4gIGltcG9ydE1hcEpTT04/OiBzdHJpbmc7XG59XG5cbmludGVyZmFjZSBNYW5pZmVzdCB7XG4gIG1vZHVsZXM/OiBBcnJheTx7XG4gICAgZGVjbGFyYXRpb25zPzogQXJyYXk8e1xuICAgICAgY3VzdG9tRWxlbWVudD86IGJvb2xlYW47XG4gICAgICB0YWdOYW1lPzogc3RyaW5nO1xuICAgICAgbmFtZT86IHN0cmluZztcbiAgICAgIGtpbmQ/OiBzdHJpbmc7XG4gICAgICBldmVudHM/OiBBcnJheTx7IG5hbWU6IHN0cmluZzsgdHlwZT86IHsgdGV4dDogc3RyaW5nIH07IHN1bW1hcnk/OiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH0+O1xuICAgIH0+O1xuICB9Pjtcbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZmlyZWQgd2hlbiBsb2dzIGFyZSByZWNlaXZlZFxuICovXG5leHBvcnQgY2xhc3MgQ2VtTG9nc0V2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT47XG4gIGNvbnN0cnVjdG9yKGxvZ3M6IEFycmF5PHsgdHlwZTogc3RyaW5nOyBkYXRlOiBzdHJpbmc7IG1lc3NhZ2U6IHN0cmluZyB9Pikge1xuICAgIHN1cGVyKCdjZW06bG9ncycsIHsgYnViYmxlczogdHJ1ZSB9KTtcbiAgICB0aGlzLmxvZ3MgPSBsb2dzO1xuICB9XG59XG5cbi8qKlxuICogQ0VNIFNlcnZlIENocm9tZSAtIE1haW4gZGVtbyB3cmFwcGVyIGNvbXBvbmVudFxuICpcbiAqIEBzbG90IC0gRGVmYXVsdCBzbG90IGZvciBkZW1vIGNvbnRlbnRcbiAqIEBzbG90IG5hdmlnYXRpb24gLSBOYXZpZ2F0aW9uIHNpZGViYXIgY29udGVudFxuICogQHNsb3Qga25vYnMgLSBLbm9iIGNvbnRyb2xzXG4gKiBAc2xvdCBkZXNjcmlwdGlvbiAtIERlbW8gZGVzY3JpcHRpb25cbiAqIEBzbG90IG1hbmlmZXN0LXRyZWUgLSBNYW5pZmVzdCB0cmVlIHZpZXdcbiAqIEBzbG90IG1hbmlmZXN0LW5hbWUgLSBNYW5pZmVzdCBuYW1lIGRpc3BsYXlcbiAqIEBzbG90IG1hbmlmZXN0LWRldGFpbHMgLSBNYW5pZmVzdCBkZXRhaWxzIGRpc3BsYXlcbiAqXG4gKiBAY3VzdG9tRWxlbWVudCBjZW0tc2VydmUtY2hyb21lXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tc2VydmUtY2hyb21lJylcbmV4cG9ydCBjbGFzcyBDZW1TZXJ2ZUNocm9tZSBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIC8vIFN0YXRpYyB0ZW1wbGF0ZXMgZm9yIERPTSBjbG9uaW5nIChsb2dzLCBldmVudHMsIGRlYnVnIGluZm8pXG4gIHN0YXRpYyAjZGVtb0luZm9UZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8aDM+RGVtbyBJbmZvcm1hdGlvbjwvaDM+XG4gICAgICA8ZGwgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3QgcGYtbS1ob3Jpem9udGFsIHBmLW0tY29tcGFjdFwiPlxuICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlRhZyBOYW1lPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cInRhZ05hbWVcIj48L2RkPlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIiBkYXRhLWZpZWxkLWdyb3VwPVwiZGVzY3JpcHRpb25cIj5cbiAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5EZXNjcmlwdGlvbjwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJkZXNjcmlwdGlvblwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiIGRhdGEtZmllbGQtZ3JvdXA9XCJmaWxlcGF0aFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkZpbGUgUGF0aDwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJmaWxlcGF0aFwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkNhbm9uaWNhbCBVUkw8L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+TG9jYWwgUm91dGU8L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwibG9jYWxSb3V0ZVwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9kbD5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjZGVtb0dyb3VwVGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCIgZGF0YS1maWVsZD1cInRhZ05hbWVcIj48L2R0PlxuICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCI+XG4gICAgICAgICAgPHNwYW4gZGF0YS1maWVsZD1cImRlc2NyaXB0aW9uXCI+PC9zcGFuPjxicj5cbiAgICAgICAgICA8c21hbGwgZGF0YS1maWVsZC1ncm91cD1cImZpbGVwYXRoXCI+RmlsZTogPHNwYW4gZGF0YS1maWVsZD1cImZpbGVwYXRoXCI+PC9zcGFuPjwvc21hbGw+XG4gICAgICAgICAgPHNtYWxsPkNhbm9uaWNhbDogPHNwYW4gZGF0YS1maWVsZD1cImNhbm9uaWNhbFVSTFwiPjwvc3Bhbj48L3NtYWxsPjxicj5cbiAgICAgICAgICA8c21hbGw+TG9jYWw6IDxzcGFuIGRhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCI+PC9zcGFuPjwvc21hbGw+XG4gICAgICAgIDwvZGQ+XG4gICAgICA8L2Rpdj5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjZGVtb0xpc3RUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8cGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uIGlkPVwiZGVidWctZGVtb3Mtc2VjdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHRvZ2dsZS10ZXh0PVwiU2hvdyBEZW1vcyBJbmZvXCI+XG4gICAgICAgIDxkbCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdCBwZi1tLWhvcml6b250YWwgcGYtbS1jb21wYWN0XCIgZGF0YS1jb250YWluZXI9XCJncm91cHNcIj48L2RsPlxuICAgICAgPC9wZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24+YDtcbiAgICByZXR1cm4gdDtcbiAgfSkoKTtcblxuICBzdGF0aWMgI2xvZ0VudHJ5VGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPGRpdiBjbGFzcz1cImxvZy1lbnRyeVwiIGRhdGEtZmllbGQ9XCJjb250YWluZXJcIj5cbiAgICAgICAgPHBmLXY2LWxhYmVsIGNvbXBhY3QgZGF0YS1maWVsZD1cImxhYmVsXCI+PC9wZi12Ni1sYWJlbD5cbiAgICAgICAgPHRpbWUgY2xhc3M9XCJsb2ctdGltZVwiIGRhdGEtZmllbGQ9XCJ0aW1lXCI+PC90aW1lPlxuICAgICAgICA8c3BhbiBjbGFzcz1cImxvZy1tZXNzYWdlXCIgZGF0YS1maWVsZD1cIm1lc3NhZ2VcIj48L3NwYW4+XG4gICAgICA8L2Rpdj5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjZXZlbnRFbnRyeVRlbXBsYXRlID0gKCgpID0+IHtcbiAgICBjb25zdCB0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGVtcGxhdGUnKTtcbiAgICB0LmlubmVySFRNTCA9IGBcbiAgICAgIDxidXR0b24gY2xhc3M9XCJldmVudC1saXN0LWl0ZW1cIiBkYXRhLWZpZWxkPVwiY29udGFpbmVyXCI+XG4gICAgICAgIDxwZi12Ni1sYWJlbCBjb21wYWN0IGRhdGEtZmllbGQ9XCJsYWJlbFwiPjwvcGYtdjYtbGFiZWw+XG4gICAgICAgIDx0aW1lIGNsYXNzPVwiZXZlbnQtdGltZVwiIGRhdGEtZmllbGQ9XCJ0aW1lXCI+PC90aW1lPlxuICAgICAgICA8c3BhbiBjbGFzcz1cImV2ZW50LWVsZW1lbnRcIiBkYXRhLWZpZWxkPVwiZWxlbWVudFwiPjwvc3Bhbj5cbiAgICAgIDwvYnV0dG9uPmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAncHJpbWFyeS10YWctbmFtZScgfSlcbiAgYWNjZXNzb3IgcHJpbWFyeVRhZ05hbWUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdkZW1vLXRpdGxlJyB9KVxuICBhY2Nlc3NvciBkZW1vVGl0bGUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdwYWNrYWdlLW5hbWUnIH0pXG4gIGFjY2Vzc29yIHBhY2thZ2VOYW1lID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnY2Fub25pY2FsLXVybCcgfSlcbiAgYWNjZXNzb3IgY2Fub25pY2FsVVJMID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnc291cmNlLXVybCcgfSlcbiAgYWNjZXNzb3Igc291cmNlVVJMID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3Iga25vYnMgPSAnJztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBkcmF3ZXI6ICdleHBhbmRlZCcgfCAnY29sbGFwc2VkJyB8ICcnID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnZHJhd2VyLWhlaWdodCcgfSlcbiAgYWNjZXNzb3IgZHJhd2VySGVpZ2h0ID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAndGFicy1zZWxlY3RlZCcgfSlcbiAgYWNjZXNzb3IgdGFic1NlbGVjdGVkID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3Igc2lkZWJhcjogJ2V4cGFuZGVkJyB8ICdjb2xsYXBzZWQnIHwgJycgPSAnJztcblxuICBAcHJvcGVydHkoeyB0eXBlOiBCb29sZWFuLCBhdHRyaWJ1dGU6ICdoYXMtZGVzY3JpcHRpb24nIH0pXG4gIGFjY2Vzc29yIGhhc0Rlc2NyaXB0aW9uID0gZmFsc2U7XG5cbiAgIyQoaWQ6IHN0cmluZykge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKGlkKTtcbiAgfVxuXG4gICMkJChzZWxlY3Rvcjogc3RyaW5nKSB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvckFsbChzZWxlY3RvcikgPz8gW107XG4gIH1cblxuICAjbG9nQ29udGFpbmVyOiBIVE1MRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjZHJhd2VyT3BlbiA9IGZhbHNlO1xuICAjaW5pdGlhbExvZ3NGZXRjaGVkID0gZmFsc2U7XG4gICNpc0luaXRpYWxMb2FkID0gdHJ1ZTtcbiAgI2RlYnVnRGF0YTogRGVidWdEYXRhIHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gRWxlbWVudCBldmVudCB0cmFja2luZ1xuICAjZWxlbWVudEV2ZW50TWFwOiBNYXA8c3RyaW5nLCBFdmVudEluZm8+IHwgbnVsbCA9IG51bGw7XG4gICNtYW5pZmVzdDogTWFuaWZlc3QgfCBudWxsID0gbnVsbDtcbiAgI2NhcHR1cmVkRXZlbnRzOiBFdmVudFJlY29yZFtdID0gW107XG4gICNtYXhDYXB0dXJlZEV2ZW50cyA9IDEwMDA7XG4gICNldmVudExpc3Q6IEhUTUxFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNldmVudERldGFpbEhlYWRlcjogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50RGV0YWlsQm9keTogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI3NlbGVjdGVkRXZlbnRJZDogc3RyaW5nIHwgbnVsbCA9IG51bGw7XG4gICNldmVudHNGaWx0ZXJWYWx1ZSA9ICcnO1xuICAjZXZlbnRzRmlsdGVyRGVib3VuY2VUaW1lcjogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50VHlwZUZpbHRlcnMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgI2VsZW1lbnRGaWx0ZXJzID0gbmV3IFNldDxzdHJpbmc+KCk7XG4gICNkaXNjb3ZlcmVkRWxlbWVudHMgPSBuZXcgU2V0PHN0cmluZz4oKTtcblxuICAvLyBFdmVudCBsaXN0ZW5lciByZWZlcmVuY2VzIGZvciBjbGVhbnVwXG4gICNoYW5kbGVMb2dzRXZlbnQ6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZUV4cGFuZDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG4gICNoYW5kbGVUcmVlQ29sbGFwc2U6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZVNlbGVjdDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gVGltZW91dCBJRHMgZm9yIGNsZWFudXBcbiAgI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjY29weURlYnVnRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dDogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcblxuICAvLyBMb2cgZmlsdGVyIHN0YXRlXG4gICNsb2dzRmlsdGVyVmFsdWUgPSAnJztcbiAgI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyOiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjbG9nTGV2ZWxGaWx0ZXJzID0gbmV3IFNldChbJ2luZm8nLCAnd2FybicsICdlcnJvcicsICdkZWJ1ZyddKTtcbiAgI2xvZ0xldmVsRHJvcGRvd246IEVsZW1lbnQgfCBudWxsID0gbnVsbDtcblxuICAvLyBXYXRjaCBmb3IgZHluYW1pY2FsbHkgYWRkZWQgZWxlbWVudHNcbiAgLyogYzggaWdub3JlIHN0YXJ0IC0gTXV0YXRpb25PYnNlcnZlciBjYWxsYmFjayB0ZXN0ZWQgdmlhIGludGVncmF0aW9uICovXG4gICNvYnNlcnZlciA9IG5ldyBNdXRhdGlvbk9ic2VydmVyKChtdXRhdGlvbnMpID0+IHtcbiAgICBsZXQgbmVlZHNVcGRhdGUgPSBmYWxzZTtcblxuICAgIGZvciAoY29uc3QgbXV0YXRpb24gb2YgbXV0YXRpb25zKSB7XG4gICAgICBmb3IgKGNvbnN0IG5vZGUgb2YgbXV0YXRpb24uYWRkZWROb2Rlcykge1xuICAgICAgICBpZiAobm9kZSBpbnN0YW5jZW9mIEhUTUxFbGVtZW50KSB7XG4gICAgICAgICAgY29uc3QgdGFnTmFtZSA9IG5vZGUudGFnTmFtZS50b0xvd2VyQ2FzZSgpO1xuICAgICAgICAgIGlmICh0aGlzLiNlbGVtZW50RXZlbnRNYXA/Lmhhcyh0YWdOYW1lKSAmJiAhbm9kZS5kYXRhc2V0LmNlbUV2ZW50c0F0dGFjaGVkKSB7XG4gICAgICAgICAgICBjb25zdCBldmVudEluZm8gPSB0aGlzLiNlbGVtZW50RXZlbnRNYXAuZ2V0KHRhZ05hbWUpITtcbiAgICAgICAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGV2ZW50SW5mby5ldmVudE5hbWVzKSB7XG4gICAgICAgICAgICAgIG5vZGUuYWRkRXZlbnRMaXN0ZW5lcihldmVudE5hbWUsIHRoaXMuI2hhbmRsZUVsZW1lbnRFdmVudCwgeyBjYXB0dXJlOiB0cnVlIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgbm9kZS5kYXRhc2V0LmNlbUV2ZW50c0F0dGFjaGVkID0gJ3RydWUnO1xuICAgICAgICAgICAgdGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICAgICAgICAgIG5lZWRzVXBkYXRlID0gdHJ1ZTtcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICBpZiAobmVlZHNVcGRhdGUpIHtcbiAgICAgIHRoaXMuI3VwZGF0ZUV2ZW50RmlsdGVycygpO1xuICAgIH1cbiAgfSk7XG4gIC8qIGM4IGlnbm9yZSBzdG9wICovXG5cbiAgI3dzQ2xpZW50OiBhbnk7XG4gICNjbGllbnRNb2R1bGVzTG9hZGVkID0gZmFsc2U7XG5cbiAgI2luaXRXc0NsaWVudCgpIHtcbiAgICB0aGlzLiN3c0NsaWVudCA9IG5ldyBDRU1SZWxvYWRDbGllbnQoe1xuICAgIGppdHRlck1heDogMTAwMCxcbiAgICBvdmVybGF5VGhyZXNob2xkOiAxNSxcbiAgICBiYWRnZUZhZGVEZWxheTogMjAwMCxcbiAgICAvKiBjOCBpZ25vcmUgc3RhcnQgLSBXZWJTb2NrZXQgY2FsbGJhY2tzIHRlc3RlZCB2aWEgaW50ZWdyYXRpb24gKi9cbiAgICBjYWxsYmFja3M6IHtcbiAgICAgIG9uT3BlbjogKCkgPT4ge1xuICAgICAgICB0aGlzLiMkKCdyZWNvbm5lY3Rpb24tbW9kYWwnKT8uY2xvc2UoKTtcbiAgICAgIH0sXG4gICAgICBvbkVycm9yOiAoZXJyb3JEYXRhOiB7IHRpdGxlPzogc3RyaW5nOyBtZXNzYWdlPzogc3RyaW5nOyBmaWxlPzogc3RyaW5nIH0pID0+IHtcbiAgICAgICAgaWYgKGVycm9yRGF0YT8udGl0bGUgJiYgZXJyb3JEYXRhPy5tZXNzYWdlKSB7XG4gICAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZV0gU2VydmVyIGVycm9yOicsIGVycm9yRGF0YSk7XG4gICAgICAgICAgKHRoaXMuIyQoJ2Vycm9yLW92ZXJsYXknKSBhcyBhbnkpPy5zaG93KGVycm9yRGF0YS50aXRsZSwgZXJyb3JEYXRhLm1lc3NhZ2UsIGVycm9yRGF0YS5maWxlKTtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlXSBXZWJTb2NrZXQgZXJyb3I6JywgZXJyb3JEYXRhKTtcbiAgICAgICAgfVxuICAgICAgfSxcbiAgICAgIG9uUmVjb25uZWN0aW5nOiAoeyBhdHRlbXB0LCBkZWxheSB9OiB7IGF0dGVtcHQ6IG51bWJlcjsgZGVsYXk6IG51bWJlciB9KSA9PiB7XG4gICAgICAgIGlmIChhdHRlbXB0ID49IDE1KSB7XG4gICAgICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LnNob3dNb2RhbCgpO1xuICAgICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tY29udGVudCcpIGFzIGFueSk/LnVwZGF0ZVJldHJ5SW5mbyhhdHRlbXB0LCBkZWxheSk7XG4gICAgICAgIH1cbiAgICAgIH0sXG4gICAgICBvblJlbG9hZDogKCkgPT4ge1xuICAgICAgICBjb25zdCBlcnJvck92ZXJsYXkgPSB0aGlzLiMkKCdlcnJvci1vdmVybGF5Jyk7XG4gICAgICAgIGlmIChlcnJvck92ZXJsYXk/Lmhhc0F0dHJpYnV0ZSgnb3BlbicpKSB7XG4gICAgICAgICAgKGVycm9yT3ZlcmxheSBhcyBhbnkpLmhpZGUoKTtcbiAgICAgICAgfVxuICAgICAgICB3aW5kb3cubG9jYXRpb24ucmVsb2FkKCk7XG4gICAgICB9LFxuICAgICAgb25TaHV0ZG93bjogKCkgPT4ge1xuICAgICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLW1vZGFsJykgYXMgYW55KT8uc2hvd01vZGFsKCk7XG4gICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tY29udGVudCcpIGFzIGFueSk/LnVwZGF0ZVJldHJ5SW5mbygzMCwgMzAwMDApO1xuICAgICAgfSxcbiAgICAgIG9uTG9nczogKGxvZ3M6IEFycmF5PHsgdHlwZTogc3RyaW5nOyBkYXRlOiBzdHJpbmc7IG1lc3NhZ2U6IHN0cmluZyB9PikgPT4ge1xuICAgICAgICB3aW5kb3cuZGlzcGF0Y2hFdmVudChuZXcgQ2VtTG9nc0V2ZW50KGxvZ3MpKTtcbiAgICAgIH1cbiAgICB9XG4gICAgLyogYzggaWdub3JlIHN0b3AgKi9cbiAgICB9KTtcbiAgfVxuXG4gIGdldCBkZW1vKCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5xdWVyeVNlbGVjdG9yKCdjZW0tc2VydmUtZGVtbycpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGxpbmsgcmVsPVwic3R5bGVzaGVldFwiIGhyZWY9XCIvX19jZW0vcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0LmNzc1wiPlxuICAgICAgPGxpbmsgcmVsPVwic3R5bGVzaGVldFwiIGhyZWY9XCIvX19jZW0vcGYtbGlnaHRkb20uY3NzXCI+XG5cbiAgICAgIDxwZi12Ni1wYWdlID9zaWRlYmFyLWNvbGxhcHNlZD0ke3RoaXMuc2lkZWJhciA9PT0gJ2NvbGxhcHNlZCd9PlxuICAgICAgICA8cGYtdjYtc2tpcC10by1jb250ZW50IHNsb3Q9XCJza2lwLXRvLWNvbnRlbnRcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhyZWY9XCIjbWFpbi1jb250ZW50XCI+XG4gICAgICAgICAgU2tpcCB0byBjb250ZW50XG4gICAgICAgIDwvcGYtdjYtc2tpcC10by1jb250ZW50PlxuXG4gICAgICAgIDxwZi12Ni1tYXN0aGVhZCBzbG90PVwibWFzdGhlYWRcIj5cbiAgICAgICAgICA8YSBjbGFzcz1cIm1hc3RoZWFkLWxvZ29cIlxuICAgICAgICAgICAgIGhyZWY9XCIvXCJcbiAgICAgICAgICAgICBzbG90PVwibG9nb1wiPlxuICAgICAgICAgICAgPGltZyBhbHQ9XCJDRU0gRGV2IFNlcnZlclwiXG4gICAgICAgICAgICAgICAgIHNyYz1cIi9fX2NlbS9sb2dvLnN2Z1wiPlxuICAgICAgICAgICAgJHt0aGlzLnBhY2thZ2VOYW1lID8gaHRtbGA8aDE+JHt0aGlzLnBhY2thZ2VOYW1lfTwvaDE+YCA6IG5vdGhpbmd9XG4gICAgICAgICAgPC9hPlxuICAgICAgICAgIDxwZi12Ni10b29sYmFyIHNsb3Q9XCJ0b29sYmFyXCI+XG4gICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cCB2YXJpYW50PVwiYWN0aW9uLWdyb3VwXCI+XG4gICAgICAgICAgICAgICR7dGhpcy4jcmVuZGVyU291cmNlQnV0dG9uKCl9XG4gICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImRlYnVnLWluZm9cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInBsYWluXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJEZWJ1ZyBpbmZvXCI+XG4gICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxNlwiXG4gICAgICAgICAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTggMTVBNyA3IDAgMSAxIDggMWE3IDcgMCAwIDEgMCAxNHptMCAxQTggOCAwIDEgMCA4IDBhOCA4IDAgMCAwIDAgMTZ6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwibTguOTMgNi41ODgtMi4yOS4yODctLjA4Mi4zOC40NS4wODNjLjI5NC4wNy4zNTIuMTc2LjI4OC40NjlsLS43MzggMy40NjhjLS4xOTQuODk3LjEwNSAxLjMxOS44MDggMS4zMTkuNTQ1IDAgMS4xNzgtLjI1MiAxLjQ2NS0uNTk4bC4wODgtLjQxNmMtLjIuMTc2LS40OTIuMjQ2LS42ODYuMjQ2LS4yNzUgMC0uMzc1LS4xOTMtLjMwNC0uNTMzTDguOTMgNi41ODh6TTkgNC41YTEgMSAwIDEgMS0yIDAgMSAxIDAgMCAxIDIgMHpcIi8+XG4gICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgPHBmLXY2LXRvZ2dsZS1ncm91cCBjbGFzcz1cImNvbG9yLXNjaGVtZS10b2dnbGVcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkNvbG9yIHNjaGVtZVwiPlxuICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtIHZhbHVlPVwibGlnaHRcIj5cbiAgICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiIGFyaWEtbGFiZWw9XCJMaWdodCBtb2RlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk04IDExYTMgMyAwIDEgMSAwLTYgMyAzIDAgMCAxIDAgNnptMCAxYTQgNCAwIDEgMCAwLTggNCA0IDAgMCAwIDAgOHpNOCAwYS41LjUgMCAwIDEgLjUuNXYyYS41LjUgMCAwIDEtMSAwdi0yQS41LjUgMCAwIDEgOCAwem0wIDEzYS41LjUgMCAwIDEgLjUuNXYyYS41LjUgMCAwIDEtMSAwdi0yQS41LjUgMCAwIDEgOCAxM3ptOC01YS41LjUgMCAwIDEtLjUuNWgtMmEuNS41IDAgMCAxIDAtMWgyYS41LjUgMCAwIDEgLjUuNXpNMyA4YS41LjUgMCAwIDEtLjUuNWgtMmEuNS41IDAgMCAxIDAtMWgyQS41LjUgMCAwIDEgMyA4em0xMC42NTctNS42NTdhLjUuNSAwIDAgMSAwIC43MDdsLTEuNDE0IDEuNDE1YS41LjUgMCAxIDEtLjcwNy0uNzA4bDEuNDE0LTEuNDE0YS41LjUgMCAwIDEgLjcwNyAwem0tOS4xOTMgOS4xOTNhLjUuNSAwIDAgMSAwIC43MDdMMy4wNSAxMy42NTdhLjUuNSAwIDAgMS0uNzA3LS43MDdsMS40MTQtMS40MTRhLjUuNSAwIDAgMSAuNzA3IDB6bTkuMTkzIDIuMTIxYS41LjUgMCAwIDEtLjcwNyAwbC0xLjQxNC0xLjQxNGEuNS41IDAgMCAxIC43MDctLjcwN2wxLjQxNCAxLjQxNGEuNS41IDAgMCAxIDAgLjcwN3pNNC40NjQgNC40NjVhLjUuNSAwIDAgMS0uNzA3IDBMMi4zNDMgMy4wNWEuNS41IDAgMSAxIC43MDctLjcwN2wxLjQxNCAxLjQxNGEuNS41IDAgMCAxIDAgLjcwOHpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICAgIDxwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cImRhcmtcIj5cbiAgICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiIGFyaWEtbGFiZWw9XCJEYXJrIG1vZGVcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTYgLjI3OGEuNzY4Ljc2OCAwIDAgMSAuMDguODU4IDcuMjA4IDcuMjA4IDAgMCAwLS44NzggMy40NmMwIDQuMDIxIDMuMjc4IDcuMjc3IDcuMzE4IDcuMjc3LjUyNyAwIDEuMDQtLjA1NSAxLjUzMy0uMTZhLjc4Ny43ODcgMCAwIDEgLjgxLjMxNi43MzMuNzMzIDAgMCAxLS4wMzEuODkzQTguMzQ5IDguMzQ5IDAgMCAxIDguMzQ0IDE2QzMuNzM0IDE2IDAgMTIuMjg2IDAgNy43MSAwIDQuMjY2IDIuMTE0IDEuMzEyIDUuMTI0LjA2QS43NTIuNzUyIDAgMCAxIDYgLjI3OHpNNC44NTggMS4zMTFBNy4yNjkgNy4yNjkgMCAwIDAgMS4wMjUgNy43MWMwIDQuMDIgMy4yNzkgNy4yNzYgNy4zMTkgNy4yNzZhNy4zMTYgNy4zMTYgMCAwIDAgNS4yMDUtMi4xNjJjLS4zMzcuMDQyLS42OC4wNjMtMS4wMjkuMDYzLTQuNjEgMC04LjM0My0zLjcxNC04LjM0My04LjI5IDAtMS4xNjcuMjQyLTIuMjc4LjY4MS0zLjI4NnpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICAgIDxwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cInN5c3RlbVwiPlxuICAgICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCIgYXJpYS1sYWJlbD1cIlN5c3RlbSBwcmVmZXJlbmNlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0wIDEuNUExLjUgMS41IDAgMCAxIDEuNSAwaDEzQTEuNSAxLjUgMCAwIDEgMTYgMS41djhhMS41IDEuNSAwIDAgMS0xLjUgMS41aC0xM0ExLjUgMS41IDAgMCAxIDAgOS41di04ek0xLjUgMWEuNS41IDAgMCAwLS41LjV2OGEuNS41IDAgMCAwIC41LjVoMTNhLjUuNSAwIDAgMCAuNS0uNXYtOGEuNS41IDAgMCAwLS41LS41aC0xM3pcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0yLjUgMTJoMTFhLjUuNSAwIDAgMSAwIDFoLTExYS41LjUgMCAwIDEgMC0xem0wIDJoMTFhLjUuNSAwIDAgMSAwIDFoLTExYS41LjUgMCAwIDEgMC0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtPlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9nZ2xlLWdyb3VwPlxuICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICA8L3BmLXY2LXRvb2xiYXI+XG4gICAgICAgIDwvcGYtdjYtbWFzdGhlYWQ+XG5cbiAgICAgICAgPHBmLXY2LXBhZ2Utc2lkZWJhciBzbG90PVwic2lkZWJhclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgP2V4cGFuZGVkPSR7dGhpcy5zaWRlYmFyID09PSAnZXhwYW5kZWQnfVxuICAgICAgICAgICAgICAgICAgICAgICAgICAgID9jb2xsYXBzZWQ9JHt0aGlzLnNpZGViYXIgIT09ICdleHBhbmRlZCd9PlxuICAgICAgICAgIDxzbG90IG5hbWU9XCJuYXZpZ2F0aW9uXCI+PC9zbG90PlxuICAgICAgICA8L3BmLXY2LXBhZ2Utc2lkZWJhcj5cblxuICAgICAgICA8cGYtdjYtcGFnZS1tYWluIHNsb3Q9XCJtYWluXCIgaWQ9XCJtYWluLWNvbnRlbnRcIj5cbiAgICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICAgICAgPGZvb3RlciBjbGFzcz1cInBmLW0tc3RpY2t5LWJvdHRvbVwiPlxuICAgICAgICAgICAgPGRpdiBjbGFzcz1cImZvb3Rlci1kZXNjcmlwdGlvbiR7dGhpcy5oYXNEZXNjcmlwdGlvbiA/ICcnIDogJyBlbXB0eSd9XCI+XG4gICAgICAgICAgICAgIDxzbG90IG5hbWU9XCJkZXNjcmlwdGlvblwiPjwvc2xvdD5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgPGNlbS1kcmF3ZXIgP29wZW49JHt0aGlzLmRyYXdlciA9PT0gJ2V4cGFuZGVkJ31cbiAgICAgICAgICAgICAgICAgICAgICAgIGRyYXdlci1oZWlnaHQ9XCIke3RoaXMuZHJhd2VySGVpZ2h0IHx8ICc0MDAnfVwiPlxuICAgICAgICAgICAgICA8cGYtdjYtdGFicyBzZWxlY3RlZD1cIiR7dGhpcy50YWJzU2VsZWN0ZWQgfHwgJzAnfVwiPlxuICAgICAgICAgICAgICAgIDxwZi12Ni10YWIgdGl0bGU9XCJLbm9ic1wiPlxuICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImtub2JzLWNvbnRhaW5lclwiPlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwia25vYnNcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cCBjbGFzcz1cImtub2JzLWVtcHR5XCI+Tm8ga25vYnMgYXZhaWxhYmxlIGZvciB0aGlzIGVsZW1lbnQuPC9wPlxuICAgICAgICAgICAgICAgICAgICA8L3Nsb3Q+XG4gICAgICAgICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRhYj5cbiAgICAgICAgICAgICAgICA8cGYtdjYtdGFiIHRpdGxlPVwiTWFuaWZlc3QgQnJvd3NlclwiPlxuICAgICAgICAgICAgICAgICAgPGNlbS1tYW5pZmVzdC1icm93c2VyPlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtdHJlZVwiIHNsb3Q9XCJtYW5pZmVzdC10cmVlXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtbmFtZVwiIHNsb3Q9XCJtYW5pZmVzdC1uYW1lXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtZGV0YWlsc1wiIHNsb3Q9XCJtYW5pZmVzdC1kZXRhaWxzXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgPC9jZW0tbWFuaWZlc3QtYnJvd3Nlcj5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRhYj5cbiAgICAgICAgICAgICAgICA8cGYtdjYtdGFiIHRpdGxlPVwiU2VydmVyIExvZ3NcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJsb2dzLXdyYXBwZXJcIj5cbiAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXIgc3RpY2t5PlxuICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRleHQtaW5wdXQtZ3JvdXAgaWQ9XCJsb2dzLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHBsYWNlaG9sZGVyPVwiRmlsdGVyIGxvZ3MuLi5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpY29uPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDUxMiA1MTJcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNTA1IDQ0Mi43TDQwNS4zIDM0M2MtNC41LTQuNS0xMC42LTctMTctN0gzNzJjMjcuNi0zNS4zIDQ0LTc5LjcgNDQtMTI4QzQxNiA5My4xIDMyMi45IDAgMjA4IDBTMCA5My4xIDAgMjA4czkzLjEgMjA4IDIwOCAyMDhjNDguMyAwIDkyLjctMTYuNCAxMjgtNDR2MTYuM2MwIDYuNCAyLjUgMTIuNSA3IDE3bDk5LjcgOTkuN2M5LjQgOS40IDI0LjYgOS40IDMzLjkgMGwyOC4zLTI4LjNjOS40LTkuNCA5LjQtMjQuNi4xLTM0ek0yMDggMzM2Yy03MC43IDAtMTI4LTU3LjItMTI4LTEyOCAwLTcwLjcgNTcuMi0xMjggMTI4LTEyOCA3MC43IDAgMTI4IDU3LjIgMTI4IDEyOCAwIDcwLjctNTcuMiAxMjgtMTI4IDEyOHpcIj48L3BhdGg+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdGV4dC1pbnB1dC1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwibG9nLWxldmVsLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBsYWJlbD1cIkZpbHRlciBsb2cgbGV2ZWxzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHNwYW4gc2xvdD1cInRvZ2dsZS10ZXh0XCI+TG9nIExldmVsczwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwiaW5mb1wiIGNoZWNrZWQ+SW5mbzwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJ3YXJuXCIgY2hlY2tlZD5XYXJuaW5nczwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJlcnJvclwiIGNoZWNrZWQ+RXJyb3JzPC9wZi12Ni1tZW51LWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LW1lbnUtaXRlbSB2YXJpYW50PVwiY2hlY2tib3hcIiB2YWx1ZT1cImRlYnVnXCIgY2hlY2tlZD5EZWJ1ZzwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWRyb3Bkb3duPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWdyb3VwIHZhcmlhbnQ9XCJhY3Rpb24tZ3JvdXBcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJjb3B5LWxvZ3NcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJ0ZXJ0aWFyeVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgc2l6ZT1cInNtYWxsXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHN2ZyBzbG90PVwiaWNvblwiIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTEzIDBINmEyIDIgMCAwIDAtMiAyIDIgMiAwIDAgMC0yIDJ2MTBhMiAyIDAgMCAwIDIgMmg3YTIgMiAwIDAgMCAyLTIgMiAyIDAgMCAwIDItMlYyYTIgMiAwIDAgMC0yLTJ6bTAgMTNWNGEyIDIgMCAwIDAtMi0ySDVhMSAxIDAgMCAxIDEtMWg3YTEgMSAwIDAgMSAxIDF2MTBhMSAxIDAgMCAxLTEgMXpNMyAxM1Y0YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjlhMSAxIDAgMCAxLTEgMUg0YTEgMSAwIDAgMS0xLTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIENvcHkgTG9nc1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyPlxuICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwibG9nLWNvbnRhaW5lclwiPjwvZGl2PlxuICAgICAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICAgICAgPC9wZi12Ni10YWI+XG4gICAgICAgICAgICAgICAgPHBmLXY2LXRhYiB0aXRsZT1cIkV2ZW50c1wiPlxuICAgICAgICAgICAgICAgICAgPGRpdiBjbGFzcz1cImV2ZW50cy13cmFwcGVyXCI+XG4gICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyIHN0aWNreT5cbiAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10ZXh0LWlucHV0LWdyb3VwIGlkPVwiZXZlbnRzLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHBsYWNlaG9sZGVyPVwiRmlsdGVyIGV2ZW50cy4uLlwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGljb24+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHN2ZyBzbG90PVwiaWNvblwiIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTExLjc0MiAxMC4zNDRhNi41IDYuNSAwIDEgMC0xLjM5NyAxLjM5OGgtLjAwMWMuMDMuMDQuMDYyLjA3OC4wOTguMTE1bDMuODUgMy44NWExIDEgMCAwIDAgMS40MTUtMS40MTRsLTMuODUtMy44NWExLjAwNyAxLjAwNyAwIDAgMC0uMTE1LS4xek0xMiA2LjVhNS41IDUuNSAwIDEgMS0xMSAwIDUuNSA1LjUgMCAwIDEgMTEgMHpcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdGV4dC1pbnB1dC1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwiZXZlbnQtdHlwZS1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgZXZlbnQgdHlwZXNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5FdmVudCBUeXBlczwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1kcm9wZG93bj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwiZWxlbWVudC1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgZWxlbWVudHNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5FbGVtZW50czwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1kcm9wZG93bj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cCB2YXJpYW50PVwiYWN0aW9uLWdyb3VwXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtYnV0dG9uIGlkPVwiY2xlYXItZXZlbnRzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk01LjUgNS41QS41LjUgMCAwIDEgNiA2djZhLjUuNSAwIDAgMS0xIDBWNmEuNS41IDAgMCAxIC41LS41em0yLjUgMGEuNS41IDAgMCAxIC41LjV2NmEuNS41IDAgMCAxLTEgMFY2YS41LjUgMCAwIDEgLjUtLjV6bTMgLjVhLjUuNSAwIDAgMC0xIDB2NmEuNS41IDAgMCAwIDEgMFY2elwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGZpbGwtcnVsZT1cImV2ZW5vZGRcIiBkPVwiTTE0LjUgM2ExIDEgMCAwIDEtMSAxSDEzdjlhMiAyIDAgMCAxLTIgMkg1YTIgMiAwIDAgMS0yLTJWNGgtLjVhMSAxIDAgMCAxLTEtMVYyYTEgMSAwIDAgMSAxLTFINmExIDEgMCAwIDEgMS0xaDJhMSAxIDAgMCAxIDEgMWgzLjVhMSAxIDAgMCAxIDEgMXYxek00LjExOCA0IDQgNC4wNTlWMTNhMSAxIDAgMCAwIDEgMWg2YTEgMSAwIDAgMCAxLTFWNC4wNTlMMTEuODgyIDRINC4xMTh6TTIuNSAzVjJoMTF2MWgtMTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIENsZWFyIEV2ZW50c1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImNvcHktZXZlbnRzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0xMyAwSDZhMiAyIDAgMCAwLTIgMiAyIDIgMCAwIDAtMiAydjEwYTIgMiAwIDAgMCAyIDJoN2EyIDIgMCAwIDAgMi0yIDIgMiAwIDAgMCAyLTJWMmEyIDIgMCAwIDAtMi0yem0wIDEzVjRhMiAyIDAgMCAwLTItMkg1YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjEwYTEgMSAwIDAgMS0xIDF6TTMgMTNWNGExIDEgMCAwIDEgMS0xaDdhMSAxIDAgMCAxIDEgMXY5YTEgMSAwIDAgMS0xIDFINGExIDEgMCAwIDEtMS0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDb3B5IEV2ZW50c1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyPlxuICAgICAgICAgICAgICAgICAgICA8cGYtdjYtZHJhd2VyIGlkPVwiZXZlbnQtZHJhd2VyXCIgZXhwYW5kZWQ+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWxpc3RcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwiZXZlbnQtZGV0YWlsLWhlYWRlclwiIHNsb3Q9XCJwYW5lbC1oZWFkZXJcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwiZXZlbnQtZGV0YWlsLWJvZHlcIiBzbG90PVwicGFuZWwtYm9keVwiPjwvZGl2PlxuICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWRyYXdlcj5cbiAgICAgICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtdGFiPlxuICAgICAgICAgICAgICAgIDxwZi12Ni10YWIgdGl0bGU9XCJIZWFsdGhcIj5cbiAgICAgICAgICAgICAgICAgIDxjZW0taGVhbHRoLXBhbmVsIGNvbXBvbmVudD0ke2lmRGVmaW5lZCh0aGlzLnByaW1hcnlUYWdOYW1lKX0+XG4gICAgICAgICAgICAgICAgICA8L2NlbS1oZWFsdGgtcGFuZWw+XG4gICAgICAgICAgICAgICAgPC9wZi12Ni10YWI+XG4gICAgICAgICAgICAgIDwvcGYtdjYtdGFicz5cbiAgICAgICAgICAgIDwvY2VtLWRyYXdlcj5cbiAgICAgICAgICA8L2Zvb3Rlcj5cbiAgICAgICAgPC9wZi12Ni1wYWdlLW1haW4+XG4gICAgICA8L3BmLXY2LXBhZ2U+XG5cbiAgICAgIDxwZi12Ni1tb2RhbCBpZD1cImRlYnVnLW1vZGFsXCIgdmFyaWFudD1cImxhcmdlXCI+XG4gICAgICAgIDxoMiBzbG90PVwiaGVhZGVyXCI+RGVidWcgSW5mb3JtYXRpb248L2gyPlxuICAgICAgICA8ZGwgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3QgcGYtbS1ob3Jpem9udGFsIHBmLW0tY29tcGFjdFwiPlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5TZXJ2ZXIgVmVyc2lvbjwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy12ZXJzaW9uXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlNlcnZlciBPUzwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy1vc1wiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5XYXRjaCBEaXJlY3Rvcnk8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctd2F0Y2gtZGlyXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPk1hbmlmZXN0IFNpemU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctbWFuaWZlc3Qtc2l6ZVwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5EZW1vcyBGb3VuZDwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy1kZW1vLWNvdW50XCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlRhZyBOYW1lPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIj4ke3RoaXMucHJpbWFyeVRhZ05hbWUgfHwgJy0nfTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkRlbW8gVGl0bGU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiPiR7dGhpcy5kZW1vVGl0bGUgfHwgJy0nfTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkJyb3dzZXI8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctYnJvd3NlclwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5Vc2VyIEFnZW50PC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLXVhXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgIDwvZGw+XG4gICAgICAgIDxkaXYgaWQ9XCJkZW1vLXVybHMtY29udGFpbmVyXCI+PC9kaXY+XG4gICAgICAgIDxwZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24gaWQ9XCJkZWJ1Zy1pbXBvcnRtYXAtZGV0YWlsc1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdG9nZ2xlLXRleHQ9XCJTaG93IEltcG9ydCBNYXBcIj5cbiAgICAgICAgICA8cHJlIGlkPVwiZGVidWctaW1wb3J0bWFwXCI+LTwvcHJlPlxuICAgICAgICA8L3BmLXY2LWV4cGFuZGFibGUtc2VjdGlvbj5cbiAgICAgICAgPGRpdiBzbG90PVwiZm9vdGVyXCIgY2xhc3M9XCJidXR0b24tY29udGFpbmVyXCI+XG4gICAgICAgICAgPHBmLXY2LWJ1dHRvbiBjbGFzcz1cImRlYnVnLWNvcHlcIiB2YXJpYW50PVwicHJpbWFyeVwiPlxuICAgICAgICAgICAgQ29weSBEZWJ1ZyBJbmZvXG4gICAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICAgICAgPHBmLXY2LWJ1dHRvbiBjbGFzcz1cImRlYnVnLWNsb3NlXCIgdmFyaWFudD1cInNlY29uZGFyeVwiPlxuICAgICAgICAgICAgQ2xvc2VcbiAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgPC9kaXY+XG4gICAgICA8L3BmLXY2LW1vZGFsPlxuXG4gICAgICA8IS0tIFJlY29ubmVjdGlvbiBtb2RhbCAtLT5cbiAgICAgIDxwZi12Ni1tb2RhbCBpZD1cInJlY29ubmVjdGlvbi1tb2RhbFwiIHZhcmlhbnQ9XCJsYXJnZVwiPlxuICAgICAgICA8aDIgc2xvdD1cImhlYWRlclwiPkRldmVsb3BtZW50IFNlcnZlciBEaXNjb25uZWN0ZWQ8L2gyPlxuICAgICAgICA8Y2VtLXJlY29ubmVjdGlvbi1jb250ZW50IGlkPVwicmVjb25uZWN0aW9uLWNvbnRlbnRcIj48L2NlbS1yZWNvbm5lY3Rpb24tY29udGVudD5cbiAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cInJlbG9hZC1idXR0b25cIlxuICAgICAgICAgICAgICAgICAgICAgIHNsb3Q9XCJmb290ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwcmltYXJ5XCI+UmVsb2FkIFBhZ2U8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cInJldHJ5LWJ1dHRvblwiXG4gICAgICAgICAgICAgICAgICAgICAgc2xvdD1cImZvb3RlclwiXG4gICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInNlY29uZGFyeVwiPlJldHJ5IE5vdzwvcGYtdjYtYnV0dG9uPlxuICAgICAgPC9wZi12Ni1tb2RhbD5cblxuICAgICAgPCEtLSBUcmFuc2Zvcm0gZXJyb3Igb3ZlcmxheSAtLT5cbiAgICAgIDxjZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXkgaWQ9XCJlcnJvci1vdmVybGF5XCI+XG4gICAgICA8L2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheT5cbiAgICBgO1xuICB9XG5cbiAgI3JlbmRlclNvdXJjZUJ1dHRvbigpIHtcbiAgICBpZiAoIXRoaXMuc291cmNlVVJMKSByZXR1cm4gbm90aGluZztcblxuICAgIGxldCBsYWJlbCA9ICdWZXJzaW9uIENvbnRyb2wnO1xuICAgIGxldCBwYXRoID0gJ001Ljg1NCA0Ljg1NGEuNS41IDAgMSAwLS43MDgtLjcwOGwtMy41IDMuNWEuNS41IDAgMCAwIDAgLjcwOGwzLjUgMy41YS41LjUgMCAwIDAgLjcwOC0uNzA4TDIuNzA3IDhsMy4xNDctMy4xNDZ6bTQuMjkyIDBhLjUuNSAwIDAgMSAuNzA4LS43MDhsMy41IDMuNWEuNS41IDAgMCAxIDAgLjcwOGwtMy41IDMuNWEuNS41IDAgMCAxLS43MDgtLjcwOEwxMy4yOTMgOGwtMy4xNDctMy4xNDZ6JztcblxuICAgIGlmICh0aGlzLnNvdXJjZVVSTC5pbmNsdWRlcygnZ2l0aHViLmNvbScpKSB7XG4gICAgICBsYWJlbCA9ICdHaXRIdWIuY29tJztcbiAgICAgIHBhdGggPSAnTTggMEMzLjU4IDAgMCAzLjU4IDAgOGMwIDMuNTQgMi4yOSA2LjUzIDUuNDcgNy41OS40LjA3LjU1LS4xNy41NS0uMzggMC0uMTktLjAxLS44Mi0uMDEtMS40OS0yLjAxLjM3LTIuNTMtLjQ5LTIuNjktLjk0LS4wOS0uMjMtLjQ4LS45NC0uODItMS4xMy0uMjgtLjE1LS42OC0uNTItLjAxLS41My42My0uMDEgMS4wOC41OCAxLjIzLjgyLjcyIDEuMjEgMS44Ny44NyAyLjMzLjY2LjA3LS41Mi4yOC0uODcuNTEtMS4wNy0xLjc4LS4yLTMuNjQtLjg5LTMuNjQtMy45NSAwLS44Ny4zMS0xLjU5LjgyLTIuMTUtLjA4LS4yLS4zNi0xLjAyLjA4LTIuMTIgMCAwIC42Ny0uMjEgMi4yLjgyLjY0LS4xOCAxLjMyLS4yNyAyLS4yNy42OCAwIDEuMzYuMDkgMiAuMjcgMS41My0xLjA0IDIuMi0uODIgMi4yLS44Mi40NCAxLjEuMTYgMS45Mi4wOCAyLjEyLjUxLjU2LjgyIDEuMjcuODIgMi4xNSAwIDMuMDctMS44NyAzLjc1LTMuNjUgMy45NS4yOS4yNS41NC43My41NCAxLjQ4IDAgMS4wNy0uMDEgMS45My0uMDEgMi4yIDAgLjIxLjE1LjQ2LjU1LjM4QTguMDEzIDguMDEzIDAgMDAxNiA4YzAtNC40Mi0zLjU4LTgtOC04eic7XG4gICAgfSBlbHNlIGlmICh0aGlzLnNvdXJjZVVSTC5pbmNsdWRlcygnZ2l0bGFiLmNvbScpKSB7XG4gICAgICBsYWJlbCA9ICdHaXRMYWInO1xuICAgICAgcGF0aCA9ICdtMTUuNzM0IDYuMS0uMDIyLS4wNThMMTMuNTM0LjM1OGEuNTY4LjU2OCAwIDAgMC0uNTYzLS4zNTYuNTgzLjU4MyAwIDAgMC0uMzI4LjEyMi41ODIuNTgyIDAgMCAwLS4xOTMuMjk0bC0xLjQ3IDQuNDk5SDUuMDI1bC0xLjQ3LTQuNUEuNTcyLjU3MiAwIDAgMCAzLjM2MC4xNzRhLjU3Mi41NzIgMCAwIDAtLjMyOC0uMTcyLjU4Mi41ODIgMCAwIDAtLjU2My4zNTdMLjI5IDYuMDRsLS4wMjIuMDU3QTQuMDQ0IDQuMDQ0IDAgMCAwIDEuNjEgMTAuNzdsLjAwNy4wMDYuMDIuMDE0IDMuMzE4IDIuNDg1IDEuNjQgMS4yNDIgMSAuNzU1YS42NzMuNjczIDAgMCAwIC44MTQgMGwxLS43NTUgMS42NC0xLjI0MiAzLjMzOC0yLjUuMDA5LS4wMDdhNC4wNSA0LjA1IDAgMCAwIDEuMzQtNC42NjhaJztcbiAgICB9IGVsc2UgaWYgKHRoaXMuc291cmNlVVJMLmluY2x1ZGVzKCdiaXRidWNrZXQub3JnJykpIHtcbiAgICAgIGxhYmVsID0gJ0JpdGJ1Y2tldCc7XG4gICAgICBwYXRoID0gJ00wIDEuNUExLjUgMS41IDAgMCAxIDEuNSAwaDEzQTEuNSAxLjUgMCAwIDEgMTYgMS41djEzYTEuNSAxLjUgMCAwIDEtMS41IDEuNWgtMTNBMS41IDEuNSAwIDAgMSAwIDE0LjV2LTEzek0yLjUgM2EuNS41IDAgMCAwLS41LjV2OWEuNS41IDAgMCAwIC41LjVoMTFhLjUuNSAwIDAgMCAuNS0uNXYtOWEuNS41IDAgMCAwLS41LS41aC0xMXptNS4wMzggMS40MzVhLjUuNSAwIDAgMSAuOTI0IDBsMS40MiAzLjM3SDguNzhsLS4yNDMuNjA4LS4yNDMtLjYwOEg1LjA4MmwxLjQyLTMuMzd6TTggOS4xNDNsLS43NDMgMS44NTdINC43NDNMNi4wNzYgNy42MDggOCA5LjE0M3onO1xuICAgIH1cblxuICAgIHJldHVybiBodG1sYFxuICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgPHBmLXY2LWJ1dHRvbiBocmVmPVwiJHt0aGlzLnNvdXJjZVVSTH1cIlxuICAgICAgICAgICAgICAgICAgICAgIHRhcmdldD1cIl9ibGFua1wiXG4gICAgICAgICAgICAgICAgICAgICAgcmVsPVwibm9vcGVuZXIgbm9yZWZlcnJlclwiXG4gICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInBsYWluXCJcbiAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiVmlldyBzb3VyY2UgZmlsZVwiPlxuICAgICAgICAgIDxzdmcgYXJpYS1sYWJlbD1cIiR7bGFiZWx9XCJcbiAgICAgICAgICAgICAgIHdpZHRoPVwiMXJlbVwiXG4gICAgICAgICAgICAgICBoZWlnaHQ9XCIxcmVtXCJcbiAgICAgICAgICAgICAgIGZpbGw9XCJjdXJyZW50Q29sb3JcIlxuICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCAxNiAxNlwiPlxuICAgICAgICAgICAgPHBhdGggZD1cIiR7cGF0aH1cIi8+XG4gICAgICAgICAgPC9zdmc+XG4gICAgICAgIDwvcGYtdjYtYnV0dG9uPlxuICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgYDtcbiAgfVxuXG4gIGFzeW5jIGNvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmNvbm5lY3RlZENhbGxiYWNrKCk7XG5cbiAgICAvLyBMb2FkIGNsaWVudC1vbmx5IG1vZHVsZXMgZHluYW1pY2FsbHkgKG5vdCBhdmFpbGFibGUgZHVyaW5nIFNTUilcbiAgICBpZiAoIXRoaXMuI2NsaWVudE1vZHVsZXNMb2FkZWQpIHtcbiAgICAgIFt7IENFTVJlbG9hZENsaWVudCB9LCB7IFN0YXRlUGVyc2lzdGVuY2UgfV0gPSBhd2FpdCBQcm9taXNlLmFsbChbXG4gICAgICAgIC8vIEB0cy1pZ25vcmUgLS0gcGxhaW4gSlMgbW9kdWxlcyBzZXJ2ZWQgYXQgcnVudGltZSBieSBHbyBzZXJ2ZXJcbiAgICAgICAgaW1wb3J0KCcvX19jZW0vd2Vic29ja2V0LWNsaWVudC5qcycpLFxuICAgICAgICAvLyBAdHMtaWdub3JlXG4gICAgICAgIGltcG9ydCgnL19fY2VtL3N0YXRlLXBlcnNpc3RlbmNlLmpzJyksXG4gICAgICBdKTtcbiAgICAgIC8vIEB0cy1pZ25vcmVcbiAgICAgIGltcG9ydCgnL19fY2VtL2hlYWx0aC1iYWRnZXMuanMnKS5jYXRjaCgoZTogdW5rbm93bikgPT5cbiAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZV0gRmFpbGVkIHRvIGxvYWQgaGVhbHRoLWJhZGdlczonLCBlKSk7XG4gICAgICB0aGlzLiNjbGllbnRNb2R1bGVzTG9hZGVkID0gdHJ1ZTtcbiAgICAgIHRoaXMuI2luaXRXc0NsaWVudCgpO1xuICAgICAgdGhpcy4jbWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2VJZk5lZWRlZCgpO1xuICAgIH1cbiAgfVxuXG4gIGZpcnN0VXBkYXRlZCgpIHtcbiAgICAvLyBTZXQgdXAgZGVidWcgb3ZlcmxheVxuICAgIHRoaXMuI3NldHVwRGVidWdPdmVybGF5KCk7XG5cbiAgICAvLyBTZXQgdXAgY29sb3Igc2NoZW1lIHRvZ2dsZVxuICAgIHRoaXMuI3NldHVwQ29sb3JTY2hlbWVUb2dnbGUoKTtcblxuICAgIC8vIFNldCB1cCBmb290ZXIgZHJhd2VyIGFuZCB0YWJzXG4gICAgdGhpcy4jc2V0dXBGb290ZXJEcmF3ZXIoKTtcblxuICAgIC8vIExpc3RlbiBmb3Igc2VydmVyIGxvZyBtZXNzYWdlcyBmcm9tIFdlYlNvY2tldFxuICAgIHRoaXMuI3NldHVwTG9nTGlzdGVuZXIoKTtcblxuICAgIC8vIFNldCB1cCBrbm9iIGV2ZW50IGNvb3JkaW5hdGlvblxuICAgIHRoaXMuI3NldHVwS25vYkNvb3JkaW5hdGlvbigpO1xuXG4gICAgLy8gU2V0IHVwIHRyZWUgc3RhdGUgcGVyc2lzdGVuY2VcbiAgICB0aGlzLiNzZXR1cFRyZWVTdGF0ZVBlcnNpc3RlbmNlKCk7XG5cbiAgICAvLyBTZXQgdXAgc2lkZWJhciBzdGF0ZSBwZXJzaXN0ZW5jZVxuICAgIHRoaXMuI3NldHVwU2lkZWJhclN0YXRlUGVyc2lzdGVuY2UoKTtcblxuICAgIC8vIFNldCB1cCBlbGVtZW50IGV2ZW50IGNhcHR1cmVcbiAgICB0aGlzLiNzZXR1cEV2ZW50Q2FwdHVyZSgpLnRoZW4oKCkgPT4ge1xuICAgICAgdGhpcy4jc2V0dXBFdmVudExpc3RlbmVycygpO1xuICAgIH0pO1xuXG4gICAgLy8gU2V0IHVwIHJlY29ubmVjdGlvbiBtb2RhbCBidXR0b24gaGFuZGxlcnNcbiAgICAvKiBjOCBpZ25vcmUgc3RhcnQgLSB3aW5kb3cubG9jYXRpb24ucmVsb2FkIHdvdWxkIHJlbG9hZCB0ZXN0IHBhZ2UgKi9cbiAgICB0aGlzLiMkKCdyZWxvYWQtYnV0dG9uJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgd2luZG93LmxvY2F0aW9uLnJlbG9hZCgpO1xuICAgIH0pO1xuICAgIC8qIGM4IGlnbm9yZSBzdG9wICovXG5cbiAgICB0aGlzLiMkKCdyZXRyeS1idXR0b24nKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLW1vZGFsJykgYXMgYW55KT8uY2xvc2UoKTtcbiAgICAgIHRoaXMuI3dzQ2xpZW50LnJldHJ5KCk7XG4gICAgfSk7XG5cbiAgICAvLyBJbml0aWFsaXplIFdlYlNvY2tldCBjb25uZWN0aW9uXG4gICAgdGhpcy4jd3NDbGllbnQuaW5pdCgpO1xuICB9XG5cbiAgYXN5bmMgI2ZldGNoRGVidWdJbmZvKCkge1xuICAgIC8vIFBvcHVsYXRlIGJyb3dzZXIgaW5mbyBpbW1lZGlhdGVseVxuICAgIGNvbnN0IGJyb3dzZXJFbCA9IHRoaXMuIyQoJ2RlYnVnLWJyb3dzZXInKTtcbiAgICBjb25zdCB1YUVsID0gdGhpcy4jJCgnZGVidWctdWEnKTtcbiAgICBpZiAoYnJvd3NlckVsKSB7XG4gICAgICBjb25zdCBicm93c2VyID0gdGhpcy4jZGV0ZWN0QnJvd3NlcigpO1xuICAgICAgYnJvd3NlckVsLnRleHRDb250ZW50ID0gYnJvd3NlcjtcbiAgICB9XG4gICAgaWYgKHVhRWwpIHtcbiAgICAgIHVhRWwudGV4dENvbnRlbnQgPSBuYXZpZ2F0b3IudXNlckFnZW50O1xuICAgIH1cblxuICAgIC8vIEZldGNoIHNlcnZlciBkZWJ1ZyBpbmZvXG4gICAgY29uc3QgZGF0YSA9IGF3YWl0IGZldGNoKCcvX19jZW0vZGVidWcnKVxuICAgICAgLnRoZW4ocmVzID0+IHJlcy5qc29uKCkpXG4gICAgICAuY2F0Y2goKGVycjogRXJyb3IpID0+IHtcbiAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBmZXRjaCBkZWJ1ZyBpbmZvOicsIGVycik7XG4gICAgICB9KTtcblxuICAgIGlmICghZGF0YSkgcmV0dXJuO1xuICAgIGNvbnN0IHZlcnNpb25FbCA9IHRoaXMuIyQoJ2RlYnVnLXZlcnNpb24nKTtcbiAgICBjb25zdCBvc0VsID0gdGhpcy4jJCgnZGVidWctb3MnKTtcbiAgICBjb25zdCB3YXRjaERpckVsID0gdGhpcy4jJCgnZGVidWctd2F0Y2gtZGlyJyk7XG4gICAgY29uc3QgbWFuaWZlc3RTaXplRWwgPSB0aGlzLiMkKCdkZWJ1Zy1tYW5pZmVzdC1zaXplJyk7XG4gICAgY29uc3QgZGVtb0NvdW50RWwgPSB0aGlzLiMkKCdkZWJ1Zy1kZW1vLWNvdW50Jyk7XG4gICAgY29uc3QgZGVtb1VybHNDb250YWluZXIgPSB0aGlzLiMkKCdkZW1vLXVybHMtY29udGFpbmVyJyk7XG4gICAgY29uc3QgaW1wb3J0TWFwRWwgPSB0aGlzLiMkKCdkZWJ1Zy1pbXBvcnRtYXAnKTtcblxuICAgIGlmICh2ZXJzaW9uRWwpIHZlcnNpb25FbC50ZXh0Q29udGVudCA9IGRhdGEudmVyc2lvbiB8fCAnLSc7XG4gICAgaWYgKG9zRWwpIG9zRWwudGV4dENvbnRlbnQgPSBkYXRhLm9zIHx8ICctJztcbiAgICBpZiAod2F0Y2hEaXJFbCkgd2F0Y2hEaXJFbC50ZXh0Q29udGVudCA9IGRhdGEud2F0Y2hEaXIgfHwgJy0nO1xuICAgIGlmIChtYW5pZmVzdFNpemVFbCkgbWFuaWZlc3RTaXplRWwudGV4dENvbnRlbnQgPSBkYXRhLm1hbmlmZXN0U2l6ZSB8fCAnLSc7XG4gICAgaWYgKGRlbW9Db3VudEVsKSBkZW1vQ291bnRFbC50ZXh0Q29udGVudCA9IGRhdGEuZGVtb0NvdW50IHx8ICcwJztcblxuICAgIGlmIChkZW1vVXJsc0NvbnRhaW5lcikge1xuICAgICAgdGhpcy4jcG9wdWxhdGVEZW1vVXJscyhkZW1vVXJsc0NvbnRhaW5lciwgZGF0YS5kZW1vcyk7XG4gICAgfVxuXG4gICAgaWYgKGltcG9ydE1hcEVsICYmIGRhdGEuaW1wb3J0TWFwKSB7XG4gICAgICBjb25zdCBpbXBvcnRNYXBKU09OID0gSlNPTi5zdHJpbmdpZnkoZGF0YS5pbXBvcnRNYXAsIG51bGwsIDIpO1xuICAgICAgaW1wb3J0TWFwRWwudGV4dENvbnRlbnQgPSBpbXBvcnRNYXBKU09OO1xuICAgICAgZGF0YS5pbXBvcnRNYXBKU09OID0gaW1wb3J0TWFwSlNPTjtcbiAgICB9IGVsc2UgaWYgKGltcG9ydE1hcEVsKSB7XG4gICAgICBpbXBvcnRNYXBFbC50ZXh0Q29udGVudCA9ICdObyBpbXBvcnQgbWFwIGdlbmVyYXRlZCc7XG4gICAgfVxuXG4gICAgdGhpcy4jZGVidWdEYXRhID0gZGF0YTtcbiAgfVxuXG4gICNwb3B1bGF0ZURlbW9VcmxzKGNvbnRhaW5lcjogSFRNTEVsZW1lbnQsIGRlbW9zOiBEZWJ1Z0RhdGFbJ2RlbW9zJ10pIHtcbiAgICBpZiAoIWRlbW9zPy5sZW5ndGgpIHtcbiAgICAgIGNvbnRhaW5lci50ZXh0Q29udGVudCA9ICdObyBkZW1vcyBmb3VuZCBpbiBtYW5pZmVzdCc7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgY29uc3QgY3VycmVudFRhZ05hbWUgPSB0aGlzLnByaW1hcnlUYWdOYW1lIHx8ICcnO1xuICAgIGNvbnN0IGlzT25EZW1vUGFnZSA9ICEhY3VycmVudFRhZ05hbWU7XG5cbiAgICBpZiAoaXNPbkRlbW9QYWdlKSB7XG4gICAgICBjb25zdCBjdXJyZW50RGVtbyA9IGRlbW9zLmZpbmQoZGVtbyA9PiBkZW1vLnRhZ05hbWUgPT09IGN1cnJlbnRUYWdOYW1lKTtcbiAgICAgIGlmICghY3VycmVudERlbW8pIHtcbiAgICAgICAgY29udGFpbmVyLnRleHRDb250ZW50ID0gJ0N1cnJlbnQgZGVtbyBub3QgZm91bmQgaW4gbWFuaWZlc3QnO1xuICAgICAgICByZXR1cm47XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGZyYWdtZW50ID0gQ2VtU2VydmVDaHJvbWUuI2RlbW9JbmZvVGVtcGxhdGUuY29udGVudC5jbG9uZU5vZGUodHJ1ZSkgYXMgRG9jdW1lbnRGcmFnbWVudDtcblxuICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0YWdOYW1lXCJdJykhLnRleHRDb250ZW50ID0gY3VycmVudERlbW8udGFnTmFtZTtcbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCJdJykhLnRleHRDb250ZW50ID0gY3VycmVudERlbW8uY2Fub25pY2FsVVJMO1xuICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCJdJykhLnRleHRDb250ZW50ID0gY3VycmVudERlbW8ubG9jYWxSb3V0ZTtcblxuICAgICAgY29uc3QgZGVzY3JpcHRpb25Hcm91cCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkLWdyb3VwPVwiZGVzY3JpcHRpb25cIl0nKTtcbiAgICAgIGlmIChjdXJyZW50RGVtby5kZXNjcmlwdGlvbikge1xuICAgICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImRlc2NyaXB0aW9uXCJdJykhLnRleHRDb250ZW50ID0gY3VycmVudERlbW8uZGVzY3JpcHRpb247XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBkZXNjcmlwdGlvbkdyb3VwPy5yZW1vdmUoKTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgZmlsZXBhdGhHcm91cCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIl0nKTtcbiAgICAgIGlmIChjdXJyZW50RGVtby5maWxlcGF0aCkge1xuICAgICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImZpbGVwYXRoXCJdJykhLnRleHRDb250ZW50ID0gY3VycmVudERlbW8uZmlsZXBhdGg7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBmaWxlcGF0aEdyb3VwPy5yZW1vdmUoKTtcbiAgICAgIH1cblxuICAgICAgY29udGFpbmVyLnJlcGxhY2VDaGlsZHJlbihmcmFnbWVudCk7XG4gICAgfSBlbHNlIHtcbiAgICAgIGNvbnN0IGxpc3RGcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vTGlzdFRlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICAgIGNvbnN0IGdyb3Vwc0NvbnRhaW5lciA9IGxpc3RGcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1jb250YWluZXI9XCJncm91cHNcIl0nKSE7XG5cbiAgICAgIGZvciAoY29uc3QgZGVtbyBvZiBkZW1vcykge1xuICAgICAgICBjb25zdCBncm91cEZyYWdtZW50ID0gQ2VtU2VydmVDaHJvbWUuI2RlbW9Hcm91cFRlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICAgICAgZ3JvdXBGcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cInRhZ05hbWVcIl0nKSEudGV4dENvbnRlbnQgPSBkZW1vLnRhZ05hbWU7XG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJkZXNjcmlwdGlvblwiXScpIS50ZXh0Q29udGVudCA9XG4gICAgICAgICAgZGVtby5kZXNjcmlwdGlvbiB8fCAnKG5vIGRlc2NyaXB0aW9uKSc7XG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJjYW5vbmljYWxVUkxcIl0nKSEudGV4dENvbnRlbnQgPSBkZW1vLmNhbm9uaWNhbFVSTDtcbiAgICAgICAgZ3JvdXBGcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxvY2FsUm91dGVcIl0nKSEudGV4dENvbnRlbnQgPSBkZW1vLmxvY2FsUm91dGU7XG5cbiAgICAgICAgY29uc3QgZmlsZXBhdGhHcm91cCA9IGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQtZ3JvdXA9XCJmaWxlcGF0aFwiXScpO1xuICAgICAgICBpZiAoZGVtby5maWxlcGF0aCkge1xuICAgICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJmaWxlcGF0aFwiXScpIS50ZXh0Q29udGVudCA9IGRlbW8uZmlsZXBhdGg7XG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgZmlsZXBhdGhHcm91cD8ucmVtb3ZlKCk7XG4gICAgICAgIH1cblxuICAgICAgICBncm91cHNDb250YWluZXIuYXBwZW5kQ2hpbGQoZ3JvdXBGcmFnbWVudCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnRhaW5lci5yZXBsYWNlQ2hpbGRyZW4obGlzdEZyYWdtZW50KTtcbiAgICB9XG4gIH1cblxuICAjc2V0dXBMb2dMaXN0ZW5lcigpIHtcbiAgICB0aGlzLiNsb2dDb250YWluZXIgPSB0aGlzLiMkKCdsb2ctY29udGFpbmVyJyk7XG5cbiAgICBjb25zdCBsb2dzRmlsdGVyID0gdGhpcy4jJCgnbG9ncy1maWx0ZXInKSBhcyBIVE1MSW5wdXRFbGVtZW50IHwgbnVsbDtcbiAgICBpZiAobG9nc0ZpbHRlcikge1xuICAgICAgbG9nc0ZpbHRlci5hZGRFdmVudExpc3RlbmVyKCdpbnB1dCcsICgpID0+IHtcbiAgICAgICAgY29uc3QgeyB2YWx1ZSA9ICcnIH0gPSBsb2dzRmlsdGVyO1xuICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jbG9nc0ZpbHRlckRlYm91bmNlVGltZXIhKTtcbiAgICAgICAgdGhpcy4jbG9nc0ZpbHRlckRlYm91bmNlVGltZXIgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICB0aGlzLiNmaWx0ZXJMb2dzKHZhbHVlKTtcbiAgICAgICAgfSwgMzAwKTtcbiAgICAgIH0pO1xuICAgIH1cblxuICAgIHRoaXMuI2xvZ0xldmVsRHJvcGRvd24gPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJyNsb2ctbGV2ZWwtZmlsdGVyJykgPz8gbnVsbDtcbiAgICBpZiAodGhpcy4jbG9nTGV2ZWxEcm9wZG93bikge1xuICAgICAgcmVxdWVzdEFuaW1hdGlvbkZyYW1lKCgpID0+IHtcbiAgICAgICAgdGhpcy4jbG9hZExvZ0ZpbHRlclN0YXRlKCk7XG4gICAgICB9KTtcbiAgICAgIHRoaXMuI2xvZ0xldmVsRHJvcGRvd24uYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlTG9nRmlsdGVyQ2hhbmdlIGFzIEV2ZW50TGlzdGVuZXIpO1xuICAgIH1cblxuICAgIHRoaXMuIyQoJ2NvcHktbG9ncycpPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgIHRoaXMuI2NvcHlMb2dzKCk7XG4gICAgfSk7XG5cbiAgICB0aGlzLiNoYW5kbGVMb2dzRXZlbnQgPSAoKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgICAgY29uc3QgbG9ncyA9IChldmVudCBhcyBDZW1Mb2dzRXZlbnQpLmxvZ3M7XG4gICAgICBpZiAobG9ncykge1xuICAgICAgICB0aGlzLiNyZW5kZXJMb2dzKGxvZ3MpO1xuICAgICAgfVxuICAgIH0pO1xuICAgIHdpbmRvdy5hZGRFdmVudExpc3RlbmVyKCdjZW06bG9ncycsIHRoaXMuI2hhbmRsZUxvZ3NFdmVudCk7XG4gIH1cblxuICAjZmlsdGVyTG9ncyhxdWVyeTogc3RyaW5nKSB7XG4gICAgdGhpcy4jbG9nc0ZpbHRlclZhbHVlID0gcXVlcnkudG9Mb3dlckNhc2UoKTtcblxuICAgIGlmICghdGhpcy4jbG9nQ29udGFpbmVyKSByZXR1cm47XG5cbiAgICBmb3IgKGNvbnN0IGVudHJ5IG9mIHRoaXMuI2xvZ0NvbnRhaW5lci5jaGlsZHJlbikge1xuICAgICAgY29uc3QgdGV4dCA9IGVudHJ5LnRleHRDb250ZW50Py50b0xvd2VyQ2FzZSgpID8/ICcnO1xuICAgICAgY29uc3QgdGV4dE1hdGNoID0gIXRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSB8fCB0ZXh0LmluY2x1ZGVzKHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSk7XG5cbiAgICAgIGNvbnN0IGxvZ1R5cGUgPSB0aGlzLiNnZXRMb2dUeXBlRnJvbUVudHJ5KGVudHJ5KTtcbiAgICAgIGNvbnN0IGxldmVsTWF0Y2ggPSB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuaGFzKGxvZ1R5cGUpO1xuXG4gICAgICAoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmhpZGRlbiA9ICEodGV4dE1hdGNoICYmIGxldmVsTWF0Y2gpO1xuICAgIH1cbiAgfVxuXG4gICNnZXRMb2dUeXBlRnJvbUVudHJ5KGVudHJ5OiBFbGVtZW50KTogc3RyaW5nIHtcbiAgICBmb3IgKGNvbnN0IGNscyBvZiBlbnRyeS5jbGFzc0xpc3QpIHtcbiAgICAgIGlmIChbJ2luZm8nLCAnd2FybmluZycsICdlcnJvcicsICdkZWJ1ZyddLmluY2x1ZGVzKGNscykpIHtcbiAgICAgICAgcmV0dXJuIGNscyA9PT0gJ3dhcm5pbmcnID8gJ3dhcm4nIDogY2xzO1xuICAgICAgfVxuICAgIH1cbiAgICByZXR1cm4gJ2luZm8nO1xuICB9XG5cbiAgI2xvYWRMb2dGaWx0ZXJTdGF0ZSgpIHtcbiAgICB0cnkge1xuICAgICAgY29uc3Qgc2F2ZWQgPSBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWxvZy1maWx0ZXJzJyk7XG4gICAgICBpZiAoc2F2ZWQpIHtcbiAgICAgICAgdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzID0gbmV3IFNldChKU09OLnBhcnNlKHNhdmVkKSk7XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgY29uc29sZS5kZWJ1ZygnW2NlbS1zZXJ2ZS1jaHJvbWVdIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSwgdXNpbmcgZGVmYXVsdCBsb2cgZmlsdGVycycpO1xuICAgIH1cbiAgICB0aGlzLiNzeW5jQ2hlY2tib3hTdGF0ZXMoKTtcbiAgfVxuXG4gICNzeW5jQ2hlY2tib3hTdGF0ZXMoKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dMZXZlbERyb3Bkb3duKSByZXR1cm47XG4gICAgY29uc3QgbWVudUl0ZW1zID0gdGhpcy4jbG9nTGV2ZWxEcm9wZG93bi5xdWVyeVNlbGVjdG9yQWxsKCdwZi12Ni1tZW51LWl0ZW0nKTtcbiAgICBtZW51SXRlbXMuZm9yRWFjaChpdGVtID0+IHtcbiAgICAgIGNvbnN0IHZhbHVlID0gKGl0ZW0gYXMgYW55KS52YWx1ZTtcbiAgICAgIChpdGVtIGFzIGFueSkuY2hlY2tlZCA9IHRoaXMuI2xvZ0xldmVsRmlsdGVycy5oYXModmFsdWUpO1xuICAgIH0pO1xuICB9XG5cbiAgI3NhdmVMb2dGaWx0ZXJTdGF0ZSgpIHtcbiAgICB0cnkge1xuICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oJ2NlbS1zZXJ2ZS1sb2ctZmlsdGVycycsXG4gICAgICAgIEpTT04uc3RyaW5naWZ5KFsuLi50aGlzLiNsb2dMZXZlbEZpbHRlcnNdKSk7XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgLy8gbG9jYWxTdG9yYWdlIHVuYXZhaWxhYmxlIChwcml2YXRlIG1vZGUpLCBzaWxlbnRseSBjb250aW51ZVxuICAgIH1cbiAgfVxuXG4gICNoYW5kbGVMb2dGaWx0ZXJDaGFuZ2UgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgeyB2YWx1ZSwgY2hlY2tlZCB9ID0gZXZlbnQgYXMgRXZlbnQgJiB7IHZhbHVlOiBzdHJpbmc7IGNoZWNrZWQ6IGJvb2xlYW4gfTtcblxuICAgIGlmIChjaGVja2VkKSB7XG4gICAgICB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuYWRkKHZhbHVlKTtcbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmRlbGV0ZSh2YWx1ZSk7XG4gICAgfVxuXG4gICAgdGhpcy4jc2F2ZUxvZ0ZpbHRlclN0YXRlKCk7XG4gICAgdGhpcy4jZmlsdGVyTG9ncyh0aGlzLiNsb2dzRmlsdGVyVmFsdWUpO1xuICB9O1xuXG4gIGFzeW5jICNjb3B5TG9ncygpIHtcbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgY29uc3QgbG9ncyA9IEFycmF5LmZyb20odGhpcy4jbG9nQ29udGFpbmVyLmNoaWxkcmVuKVxuICAgICAgLmZpbHRlcihlbnRyeSA9PiAhKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5oaWRkZW4pXG4gICAgICAubWFwKGVudHJ5ID0+IHtcbiAgICAgICAgY29uc3QgdHlwZSA9IGVudHJ5LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibGFiZWxcIl0nKT8udGV4dENvbnRlbnQ/LnRyaW0oKSB8fCAnSU5GTyc7XG4gICAgICAgIGNvbnN0IHRpbWUgPSBlbnRyeS5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cInRpbWVcIl0nKT8udGV4dENvbnRlbnQ/LnRyaW0oKSB8fCAnJztcbiAgICAgICAgY29uc3QgbWVzc2FnZSA9IGVudHJ5LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibWVzc2FnZVwiXScpPy50ZXh0Q29udGVudD8udHJpbSgpIHx8ICcnO1xuICAgICAgICByZXR1cm4gYFske3R5cGV9XSAke3RpbWV9ICR7bWVzc2FnZX1gO1xuICAgICAgfSkuam9pbignXFxuJyk7XG5cbiAgICBpZiAoIWxvZ3MpIHJldHVybjtcblxuICAgIHRyeSB7XG4gICAgICBhd2FpdCBuYXZpZ2F0b3IuY2xpcGJvYXJkLndyaXRlVGV4dChsb2dzKTtcbiAgICAgIGNvbnN0IGJ0biA9IHRoaXMuIyQoJ2NvcHktbG9ncycpO1xuICAgICAgaWYgKGJ0bikge1xuICAgICAgICBjb25zdCB0ZXh0Tm9kZSA9IEFycmF5LmZyb20oYnRuLmNoaWxkTm9kZXMpLmZpbmQoXG4gICAgICAgICAgbiA9PiBuLm5vZGVUeXBlID09PSBOb2RlLlRFWFRfTk9ERSAmJiAobi50ZXh0Q29udGVudD8udHJpbSgpLmxlbmd0aCA/PyAwKSA+IDBcbiAgICAgICAgKTtcbiAgICAgICAgaWYgKHRleHROb2RlKSB7XG4gICAgICAgICAgY29uc3Qgb3JpZ2luYWwgPSB0ZXh0Tm9kZS50ZXh0Q29udGVudDtcbiAgICAgICAgICB0ZXh0Tm9kZS50ZXh0Q29udGVudCA9ICdDb3BpZWQhJztcblxuICAgICAgICAgIGlmICh0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICB0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgICAgaWYgKHRoaXMuaXNDb25uZWN0ZWQgJiYgdGV4dE5vZGUucGFyZW50Tm9kZSkge1xuICAgICAgICAgICAgICB0ZXh0Tm9kZS50ZXh0Q29udGVudCA9IG9yaWdpbmFsO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgdGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgICAgICAgIH0sIDIwMDApO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZXJyKSB7XG4gICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGNvcHkgbG9nczonLCBlcnIpO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cERlYnVnT3ZlcmxheSgpIHtcbiAgICBjb25zdCBkZWJ1Z0J1dHRvbiA9IHRoaXMuIyQoJ2RlYnVnLWluZm8nKTtcbiAgICBjb25zdCBkZWJ1Z01vZGFsID0gdGhpcy4jJCgnZGVidWctbW9kYWwnKTtcbiAgICBjb25zdCBkZWJ1Z0Nsb3NlID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuZGVidWctY2xvc2UnKTtcbiAgICBjb25zdCBkZWJ1Z0NvcHkgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJy5kZWJ1Zy1jb3B5Jyk7XG5cbiAgICBpZiAoZGVidWdCdXR0b24gJiYgZGVidWdNb2RhbCkge1xuICAgICAgZGVidWdCdXR0b24uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICAgIHRoaXMuI2ZldGNoRGVidWdJbmZvKCk7XG4gICAgICAgIChkZWJ1Z01vZGFsIGFzIGFueSkuc2hvd01vZGFsKCk7XG4gICAgICB9KTtcblxuICAgICAgZGVidWdDbG9zZT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiAoZGVidWdNb2RhbCBhcyBhbnkpLmNsb3NlKCkpO1xuXG4gICAgICBkZWJ1Z0NvcHk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgICB0aGlzLiNjb3B5RGVidWdJbmZvKCk7XG4gICAgICB9KTtcbiAgICB9XG4gIH1cblxuICAjc2V0dXBGb290ZXJEcmF3ZXIoKSB7XG4gICAgY29uc3QgZHJhd2VyID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdjZW0tZHJhd2VyJyk7XG4gICAgY29uc3QgdGFicyA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcigncGYtdjYtdGFicycpO1xuXG4gICAgaWYgKCFkcmF3ZXIgfHwgIXRhYnMpIHJldHVybjtcblxuICAgIHRoaXMuI2RyYXdlck9wZW4gPSAoZHJhd2VyIGFzIGFueSkub3BlbjtcblxuICAgIGRyYXdlci5hZGRFdmVudExpc3RlbmVyKCdjaGFuZ2UnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgIHRoaXMuI2RyYXdlck9wZW4gPSAoZSBhcyBhbnkpLm9wZW47XG5cbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoe1xuICAgICAgICBkcmF3ZXI6IHsgb3BlbjogKGUgYXMgYW55KS5vcGVuIH1cbiAgICAgIH0pO1xuXG4gICAgICBpZiAoKGUgYXMgYW55KS5vcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExvZ3NUb0JvdHRvbSgpO1xuICAgICAgfVxuICAgIH0pO1xuXG4gICAgZHJhd2VyLmFkZEV2ZW50TGlzdGVuZXIoJ3Jlc2l6ZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgKGRyYXdlciBhcyBhbnkpLnNldEF0dHJpYnV0ZSgnZHJhd2VyLWhlaWdodCcsIChlIGFzIGFueSkuaGVpZ2h0KTtcblxuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIGRyYXdlcjogeyBoZWlnaHQ6IChlIGFzIGFueSkuaGVpZ2h0IH1cbiAgICAgIH0pO1xuICAgIH0pO1xuXG4gICAgdGFicy5hZGRFdmVudExpc3RlbmVyKCdjaGFuZ2UnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoe1xuICAgICAgICB0YWJzOiB7IHNlbGVjdGVkSW5kZXg6IChlIGFzIGFueSkuc2VsZWN0ZWRJbmRleCB9XG4gICAgICB9KTtcblxuICAgICAgaWYgKChlIGFzIGFueSkuc2VsZWN0ZWRJbmRleCA9PT0gMiAmJiAoZHJhd2VyIGFzIGFueSkub3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxMb2dzVG9Cb3R0b20oKTtcbiAgICAgIH1cblxuICAgICAgaWYgKChlIGFzIGFueSkuc2VsZWN0ZWRJbmRleCA9PT0gMyAmJiAoZHJhd2VyIGFzIGFueSkub3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxFdmVudHNUb0JvdHRvbSgpO1xuICAgICAgfVxuICAgIH0pO1xuICB9XG5cbiAgI2RldGVjdEJyb3dzZXIoKTogc3RyaW5nIHtcbiAgICBjb25zdCB1YSA9IG5hdmlnYXRvci51c2VyQWdlbnQ7XG4gICAgaWYgKHVhLmluY2x1ZGVzKCdGaXJlZm94LycpKSB7XG4gICAgICBjb25zdCBtYXRjaCA9IHVhLm1hdGNoKC9GaXJlZm94XFwvKFxcZCspLyk7XG4gICAgICByZXR1cm4gbWF0Y2ggPyBgRmlyZWZveCAke21hdGNoWzFdfWAgOiAnRmlyZWZveCc7XG4gICAgfSBlbHNlIGlmICh1YS5pbmNsdWRlcygnRWRnLycpKSB7XG4gICAgICBjb25zdCBtYXRjaCA9IHVhLm1hdGNoKC9FZGdcXC8oXFxkKykvKTtcbiAgICAgIHJldHVybiBtYXRjaCA/IGBFZGdlICR7bWF0Y2hbMV19YCA6ICdFZGdlJztcbiAgICB9IGVsc2UgaWYgKHVhLmluY2x1ZGVzKCdDaHJvbWUvJykpIHtcbiAgICAgIGNvbnN0IG1hdGNoID0gdWEubWF0Y2goL0Nocm9tZVxcLyhcXGQrKS8pO1xuICAgICAgcmV0dXJuIG1hdGNoID8gYENocm9tZSAke21hdGNoWzFdfWAgOiAnQ2hyb21lJztcbiAgICB9IGVsc2UgaWYgKHVhLmluY2x1ZGVzKCdTYWZhcmkvJykgJiYgIXVhLmluY2x1ZGVzKCdDaHJvbWUnKSkge1xuICAgICAgY29uc3QgbWF0Y2ggPSB1YS5tYXRjaCgvVmVyc2lvblxcLyhcXGQrKS8pO1xuICAgICAgcmV0dXJuIG1hdGNoID8gYFNhZmFyaSAke21hdGNoWzFdfWAgOiAnU2FmYXJpJztcbiAgICB9XG4gICAgcmV0dXJuICdVbmtub3duJztcbiAgfVxuXG4gIGFzeW5jICNjb3B5RGVidWdJbmZvKCkge1xuICAgIGNvbnN0IGluZm8gPSBBcnJheS5mcm9tKHRoaXMuIyQkKCcjZGVidWctbW9kYWwgZGwgZHQnKSkubWFwKGR0ID0+IHtcbiAgICAgIGNvbnN0IGRkID0gZHQubmV4dEVsZW1lbnRTaWJsaW5nO1xuICAgICAgaWYgKGRkICYmIGRkLnRhZ05hbWUgPT09ICdERCcpIHtcbiAgICAgICAgcmV0dXJuIGAke2R0LnRleHRDb250ZW50fTogJHtkZC50ZXh0Q29udGVudH1gO1xuICAgICAgfVxuICAgICAgcmV0dXJuICcnO1xuICAgIH0pLmpvaW4oJ1xcbicpO1xuXG4gICAgbGV0IGltcG9ydE1hcFNlY3Rpb24gPSAnJztcbiAgICBpZiAodGhpcy4jZGVidWdEYXRhPy5pbXBvcnRNYXBKU09OKSB7XG4gICAgICBpbXBvcnRNYXBTZWN0aW9uID0gYFxcbiR7Jz0nLnJlcGVhdCg0MCl9XFxuSW1wb3J0IE1hcDpcXG4keyc9Jy5yZXBlYXQoNDApfVxcbiR7dGhpcy4jZGVidWdEYXRhLmltcG9ydE1hcEpTT059XFxuYDtcbiAgICB9XG5cbiAgICBjb25zdCBkZWJ1Z1RleHQgPSBgQ0VNIFNlcnZlIERlYnVnIEluZm9ybWF0aW9uXG4keyc9Jy5yZXBlYXQoNDApfVxuJHtpbmZvfSR7aW1wb3J0TWFwU2VjdGlvbn1cbiR7Jz0nLnJlcGVhdCg0MCl9XG5HZW5lcmF0ZWQ6ICR7bmV3IERhdGUoKS50b0lTT1N0cmluZygpfWA7XG5cbiAgICB0cnkge1xuICAgICAgYXdhaXQgbmF2aWdhdG9yLmNsaXBib2FyZC53cml0ZVRleHQoZGVidWdUZXh0KTtcbiAgICAgIGNvbnN0IGNvcHlCdXR0b24gPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJy5kZWJ1Zy1jb3B5Jyk7XG4gICAgICBpZiAoY29weUJ1dHRvbikge1xuICAgICAgICBjb25zdCBvcmlnaW5hbFRleHQgPSBjb3B5QnV0dG9uLnRleHRDb250ZW50O1xuICAgICAgICBjb3B5QnV0dG9uLnRleHRDb250ZW50ID0gJ0NvcGllZCEnO1xuXG4gICAgICAgIGlmICh0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgICAgfVxuXG4gICAgICAgIHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgIGlmICh0aGlzLmlzQ29ubmVjdGVkICYmIGNvcHlCdXR0b24ucGFyZW50Tm9kZSkge1xuICAgICAgICAgICAgY29weUJ1dHRvbi50ZXh0Q29udGVudCA9IG9yaWdpbmFsVGV4dDtcbiAgICAgICAgICB9XG4gICAgICAgICAgdGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICAgICAgfSwgMjAwMCk7XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZXJyKSB7XG4gICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGNvcHkgZGVidWcgaW5mbzonLCBlcnIpO1xuICAgIH1cbiAgfVxuXG4gICNyZW5kZXJMb2dzKGxvZ3M6IEFycmF5PHsgdHlwZTogc3RyaW5nOyBkYXRlOiBzdHJpbmc7IG1lc3NhZ2U6IHN0cmluZyB9Pikge1xuICAgIGlmICghdGhpcy4jbG9nQ29udGFpbmVyKSByZXR1cm47XG5cbiAgICBjb25zdCBsb2dFbGVtZW50cyA9IGxvZ3MubWFwKGxvZyA9PiB7XG4gICAgICBjb25zdCBmcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNsb2dFbnRyeVRlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICAgIGNvbnN0IGRhdGUgPSBuZXcgRGF0ZShsb2cuZGF0ZSk7XG4gICAgICBjb25zdCB0aW1lID0gZGF0ZS50b0xvY2FsZVRpbWVTdHJpbmcoKTtcblxuICAgICAgY29uc3QgY29udGFpbmVyID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJjb250YWluZXJcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICAgIGNvbnRhaW5lci5jbGFzc0xpc3QuYWRkKGxvZy50eXBlKTtcbiAgICAgIGNvbnRhaW5lci5zZXRBdHRyaWJ1dGUoJ2RhdGEtbG9nLWlkJywgbG9nLmRhdGUpO1xuXG4gICAgICBjb25zdCB0eXBlTGFiZWwgPSB0aGlzLiNnZXRMb2dCYWRnZShsb2cudHlwZSk7XG4gICAgICBjb25zdCBzZWFyY2hUZXh0ID0gYCR7dHlwZUxhYmVsfSAke3RpbWV9ICR7bG9nLm1lc3NhZ2V9YC50b0xvd2VyQ2FzZSgpO1xuICAgICAgY29uc3QgdGV4dE1hdGNoID0gIXRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSB8fCBzZWFyY2hUZXh0LmluY2x1ZGVzKHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSk7XG5cbiAgICAgIGNvbnN0IGxvZ1R5cGVGb3JGaWx0ZXIgPSBsb2cudHlwZSA9PT0gJ3dhcm5pbmcnID8gJ3dhcm4nIDogbG9nLnR5cGU7XG4gICAgICBjb25zdCBsZXZlbE1hdGNoID0gdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmhhcyhsb2dUeXBlRm9yRmlsdGVyKTtcblxuICAgICAgaWYgKCEodGV4dE1hdGNoICYmIGxldmVsTWF0Y2gpKSB7XG4gICAgICAgIGNvbnRhaW5lci5zZXRBdHRyaWJ1dGUoJ2hpZGRlbicsICcnKTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgbGFiZWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxhYmVsXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICBsYWJlbC50ZXh0Q29udGVudCA9IHRoaXMuI2dldExvZ0JhZGdlKGxvZy50eXBlKTtcbiAgICAgIHRoaXMuI2FwcGx5TG9nTGFiZWxBdHRycyhsYWJlbCwgbG9nLnR5cGUpO1xuXG4gICAgICBjb25zdCB0aW1lRWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cInRpbWVcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICAgIHRpbWVFbC5zZXRBdHRyaWJ1dGUoJ2RhdGV0aW1lJywgbG9nLmRhdGUpO1xuICAgICAgdGltZUVsLnRleHRDb250ZW50ID0gdGltZTtcblxuICAgICAgKGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibWVzc2FnZVwiXScpIGFzIEhUTUxFbGVtZW50KS50ZXh0Q29udGVudCA9IGxvZy5tZXNzYWdlO1xuXG4gICAgICByZXR1cm4gZnJhZ21lbnQ7XG4gICAgfSk7XG5cbiAgICBpZiAoIXRoaXMuI2luaXRpYWxMb2dzRmV0Y2hlZCkge1xuICAgICAgdGhpcy4jbG9nQ29udGFpbmVyLnJlcGxhY2VDaGlsZHJlbiguLi5sb2dFbGVtZW50cyk7XG4gICAgICB0aGlzLiNpbml0aWFsTG9nc0ZldGNoZWQgPSB0cnVlO1xuXG4gICAgICBpZiAodGhpcy4jZHJhd2VyT3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICB0aGlzLiNsb2dDb250YWluZXIuYXBwZW5kKC4uLmxvZ0VsZW1lbnRzKTtcblxuICAgICAgaWYgKHRoaXMuI2RyYXdlck9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjZ2V0TG9nQmFkZ2UodHlwZTogc3RyaW5nKTogc3RyaW5nIHtcbiAgICBzd2l0Y2ggKHR5cGUpIHtcbiAgICAgIGNhc2UgJ2luZm8nOiByZXR1cm4gJ0luZm8nO1xuICAgICAgY2FzZSAnd2FybmluZyc6IHJldHVybiAnV2Fybic7XG4gICAgICBjYXNlICdlcnJvcic6IHJldHVybiAnRXJyb3InO1xuICAgICAgY2FzZSAnZGVidWcnOiByZXR1cm4gJ0RlYnVnJztcbiAgICAgIGRlZmF1bHQ6IHJldHVybiB0eXBlLnRvVXBwZXJDYXNlKCk7XG4gICAgfVxuICB9XG5cbiAgI2FwcGx5TG9nTGFiZWxBdHRycyhsYWJlbDogSFRNTEVsZW1lbnQsIHR5cGU6IHN0cmluZykge1xuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnaW5mbyc6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICdpbmZvJyk7XG4gICAgICBjYXNlICd3YXJuaW5nJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnc3RhdHVzJywgJ3dhcm5pbmcnKTtcbiAgICAgIGNhc2UgJ2Vycm9yJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnc3RhdHVzJywgJ2RhbmdlcicpO1xuICAgICAgY2FzZSAnZGVidWcnOlxuICAgICAgICByZXR1cm4gbGFiZWwuc2V0QXR0cmlidXRlKCdjb2xvcicsICdwdXJwbGUnKTtcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIGxhYmVsLnNldEF0dHJpYnV0ZSgnY29sb3InLCAnZ3JleScpO1xuICAgIH1cbiAgfVxuXG4gICNzY3JvbGxMYXRlc3RJbnRvVmlldygpIHtcbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgcmVxdWVzdEFuaW1hdGlvbkZyYW1lKCgpID0+IHtcbiAgICAgIGNvbnN0IGxhc3RMb2cgPSB0aGlzLiNsb2dDb250YWluZXIhLmxhc3RFbGVtZW50Q2hpbGQ7XG4gICAgICBpZiAobGFzdExvZykge1xuICAgICAgICBsYXN0TG9nLnNjcm9sbEludG9WaWV3KHsgYmVoYXZpb3I6ICdhdXRvJywgYmxvY2s6ICdlbmQnIH0pO1xuICAgICAgfVxuICAgIH0pO1xuICB9XG5cbiAgI3Njcm9sbExvZ3NUb0JvdHRvbSgpIHtcbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgaWYgKHRoaXMuI2lzSW5pdGlhbExvYWQpIHtcbiAgICAgIHRoaXMuI3Njcm9sbExhdGVzdEludG9WaWV3KCk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgICAgfSwgMzUwKTtcbiAgICB9XG4gIH1cblxuICAjbWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2VJZk5lZWRlZCgpIHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgaGFzTG9jYWxTdG9yYWdlID1cbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1jb2xvci1zY2hlbWUnKSAhPT0gbnVsbCB8fFxuICAgICAgICBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWRyYXdlci1vcGVuJykgIT09IG51bGwgfHxcbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1kcmF3ZXItaGVpZ2h0JykgIT09IG51bGwgfHxcbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1hY3RpdmUtdGFiJykgIT09IG51bGw7XG5cbiAgICAgIGlmIChoYXNMb2NhbFN0b3JhZ2UpIHtcbiAgICAgICAgY29uc3QgbWlncmF0ZWQgPSBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLW1pZ3JhdGVkLXRvLWNvb2tpZXMnKTtcbiAgICAgICAgaWYgKCFtaWdyYXRlZCkge1xuICAgICAgICAgIFN0YXRlUGVyc2lzdGVuY2UubWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2UoKTtcbiAgICAgICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLW1pZ3JhdGVkLXRvLWNvb2tpZXMnLCAndHJ1ZScpO1xuICAgICAgICAgIHNldFRpbWVvdXQoKCkgPT4gd2luZG93LmxvY2F0aW9uLnJlbG9hZCgpLCAxMDApO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgLy8gbG9jYWxTdG9yYWdlIG5vdCBhdmFpbGFibGUsIHNraXAgbWlncmF0aW9uXG4gICAgfVxuICB9XG5cbiAgI3NldHVwQ29sb3JTY2hlbWVUb2dnbGUoKSB7XG4gICAgY29uc3QgdG9nZ2xlR3JvdXAgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJy5jb2xvci1zY2hlbWUtdG9nZ2xlJyk7XG4gICAgaWYgKCF0b2dnbGVHcm91cCkgcmV0dXJuO1xuXG4gICAgY29uc3Qgc3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFN0YXRlKCk7XG5cbiAgICB0aGlzLiNhcHBseUNvbG9yU2NoZW1lKHN0YXRlLmNvbG9yU2NoZW1lKTtcblxuICAgIGNvbnN0IGl0ZW1zID0gdG9nZ2xlR3JvdXAucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0nKTtcbiAgICBpdGVtcy5mb3JFYWNoKGl0ZW0gPT4ge1xuICAgICAgaWYgKChpdGVtIGFzIGFueSkudmFsdWUgPT09IHN0YXRlLmNvbG9yU2NoZW1lKSB7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCdzZWxlY3RlZCcsICcnKTtcbiAgICAgIH1cbiAgICB9KTtcblxuICAgIHRvZ2dsZUdyb3VwLmFkZEV2ZW50TGlzdGVuZXIoJ3BmLXY2LXRvZ2dsZS1ncm91cC1jaGFuZ2UnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IHNjaGVtZSA9IChlIGFzIGFueSkudmFsdWU7XG4gICAgICB0aGlzLiNhcHBseUNvbG9yU2NoZW1lKHNjaGVtZSk7XG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHsgY29sb3JTY2hlbWU6IHNjaGVtZSB9KTtcbiAgICB9KTtcbiAgfVxuXG4gICNhcHBseUNvbG9yU2NoZW1lKHNjaGVtZTogc3RyaW5nKSB7XG4gICAgY29uc3QgYm9keSA9IGRvY3VtZW50LmJvZHk7XG5cbiAgICBzd2l0Y2ggKHNjaGVtZSkge1xuICAgICAgY2FzZSAnbGlnaHQnOlxuICAgICAgICBib2R5LnN0eWxlLmNvbG9yU2NoZW1lID0gJ2xpZ2h0JztcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdkYXJrJzpcbiAgICAgICAgYm9keS5zdHlsZS5jb2xvclNjaGVtZSA9ICdkYXJrJztcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdzeXN0ZW0nOlxuICAgICAgZGVmYXVsdDpcbiAgICAgICAgYm9keS5zdHlsZS5jb2xvclNjaGVtZSA9ICdsaWdodCBkYXJrJztcbiAgICAgICAgYnJlYWs7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwS25vYkNvb3JkaW5hdGlvbigpIHtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOnByb3BlcnR5LWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOmNzcy1wcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjphdHRyaWJ1dGUtY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOnByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gIH1cblxuICAjb25Lbm9iQ2hhbmdlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMuI2dldEtub2JUYXJnZXQoZXZlbnQpO1xuICAgIGlmICghdGFyZ2V0KSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBDb3VsZCBub3QgZmluZCBrbm9iIHRhcmdldCBpbmZvIGluIGV2ZW50IHBhdGgnKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICBjb25zdCB7IHRhZ05hbWUsIGluc3RhbmNlSW5kZXggfSA9IHRhcmdldDtcblxuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCBrbm9iVHlwZSA9IHRoaXMuI2dldEtub2JUeXBlRnJvbUV2ZW50KGV2ZW50KTtcblxuICAgIGNvbnN0IHN1Y2Nlc3MgPSAoZGVtbyBhcyBhbnkpLmFwcGx5S25vYkNoYW5nZShcbiAgICAgIGtub2JUeXBlLFxuICAgICAgKGV2ZW50IGFzIGFueSkubmFtZSxcbiAgICAgIChldmVudCBhcyBhbnkpLnZhbHVlLFxuICAgICAgdGFnTmFtZSxcbiAgICAgIGluc3RhbmNlSW5kZXhcbiAgICApO1xuXG4gICAgaWYgKCFzdWNjZXNzKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gYXBwbHkga25vYiBjaGFuZ2U6Jywge1xuICAgICAgICB0eXBlOiBrbm9iVHlwZSxcbiAgICAgICAgbmFtZTogKGV2ZW50IGFzIGFueSkubmFtZSxcbiAgICAgICAgdGFnTmFtZSxcbiAgICAgICAgaW5zdGFuY2VJbmRleFxuICAgICAgfSk7XG4gICAgfVxuICB9O1xuXG4gICNvbktub2JDbGVhciA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLiNnZXRLbm9iVGFyZ2V0KGV2ZW50KTtcbiAgICBpZiAoIXRhcmdldCkge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gQ291bGQgbm90IGZpbmQga25vYiB0YXJnZXQgaW5mbyBpbiBldmVudCBwYXRoJyk7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgY29uc3QgeyB0YWdOYW1lLCBpbnN0YW5jZUluZGV4IH0gPSB0YXJnZXQ7XG5cbiAgICBjb25zdCBkZW1vID0gdGhpcy5kZW1vO1xuICAgIGlmICghZGVtbykgcmV0dXJuO1xuXG4gICAgY29uc3Qga25vYlR5cGUgPSB0aGlzLiNnZXRLbm9iVHlwZUZyb21DbGVhckV2ZW50KGV2ZW50KTtcbiAgICBjb25zdCBjbGVhclZhbHVlID0ga25vYlR5cGUgPT09ICdwcm9wZXJ0eScgPyB1bmRlZmluZWQgOiAnJztcblxuICAgIGNvbnN0IHN1Y2Nlc3MgPSAoZGVtbyBhcyBhbnkpLmFwcGx5S25vYkNoYW5nZShcbiAgICAgIGtub2JUeXBlLFxuICAgICAgKGV2ZW50IGFzIGFueSkubmFtZSxcbiAgICAgIGNsZWFyVmFsdWUsXG4gICAgICB0YWdOYW1lLFxuICAgICAgaW5zdGFuY2VJbmRleFxuICAgICk7XG5cbiAgICBpZiAoIXN1Y2Nlc3MpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjbGVhciBrbm9iOicsIHtcbiAgICAgICAgdHlwZToga25vYlR5cGUsXG4gICAgICAgIG5hbWU6IChldmVudCBhcyBhbnkpLm5hbWUsXG4gICAgICAgIHRhZ05hbWUsXG4gICAgICAgIGluc3RhbmNlSW5kZXhcbiAgICAgIH0pO1xuICAgIH1cbiAgfTtcblxuICAjZ2V0S25vYlRhcmdldChldmVudDogRXZlbnQpOiB7IHRhZ05hbWU6IHN0cmluZzsgaW5zdGFuY2VJbmRleDogbnVtYmVyIH0gfCBudWxsIHtcbiAgICBjb25zdCBkZWZhdWx0VGFnTmFtZSA9IHRoaXMucHJpbWFyeVRhZ05hbWUgfHwgJyc7XG5cbiAgICBpZiAoZXZlbnQuY29tcG9zZWRQYXRoKSB7XG4gICAgICBmb3IgKGNvbnN0IGVsZW1lbnQgb2YgZXZlbnQuY29tcG9zZWRQYXRoKCkpIHtcbiAgICAgICAgaWYgKCEoZWxlbWVudCBpbnN0YW5jZW9mIEVsZW1lbnQpKSBjb250aW51ZTtcblxuICAgICAgICBpZiAoKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQ/LmlzRWxlbWVudEtub2IgPT09ICd0cnVlJykge1xuICAgICAgICAgIGNvbnN0IHRhZ05hbWUgPSAoZWxlbWVudCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC50YWdOYW1lIHx8IGRlZmF1bHRUYWdOYW1lO1xuICAgICAgICAgIGxldCBpbnN0YW5jZUluZGV4ID0gTnVtYmVyLnBhcnNlSW50KChlbGVtZW50IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0Lmluc3RhbmNlSW5kZXggPz8gJycsIDEwKTtcbiAgICAgICAgICBpZiAoTnVtYmVyLmlzTmFOKGluc3RhbmNlSW5kZXgpKSBpbnN0YW5jZUluZGV4ID0gMDtcbiAgICAgICAgICByZXR1cm4geyB0YWdOYW1lLCBpbnN0YW5jZUluZGV4IH07XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICByZXR1cm4geyB0YWdOYW1lOiBkZWZhdWx0VGFnTmFtZSwgaW5zdGFuY2VJbmRleDogMCB9O1xuICB9XG5cbiAgI2dldEtub2JUeXBlRnJvbUV2ZW50KGV2ZW50OiBFdmVudCk6IHN0cmluZyB7XG4gICAgc3dpdGNoIChldmVudC50eXBlKSB7XG4gICAgICBjYXNlICdrbm9iOmF0dHJpYnV0ZS1jaGFuZ2UnOlxuICAgICAgICByZXR1cm4gJ2F0dHJpYnV0ZSc7XG4gICAgICBjYXNlICdrbm9iOnByb3BlcnR5LWNoYW5nZSc6XG4gICAgICAgIHJldHVybiAncHJvcGVydHknO1xuICAgICAgY2FzZSAna25vYjpjc3MtcHJvcGVydHktY2hhbmdlJzpcbiAgICAgICAgcmV0dXJuICdjc3MtcHJvcGVydHknO1xuICAgICAgZGVmYXVsdDpcbiAgICAgICAgcmV0dXJuICd1bmtub3duJztcbiAgICB9XG4gIH1cblxuICAjZ2V0S25vYlR5cGVGcm9tQ2xlYXJFdmVudChldmVudDogRXZlbnQpOiBzdHJpbmcge1xuICAgIHN3aXRjaCAoZXZlbnQudHlwZSkge1xuICAgICAgY2FzZSAna25vYjphdHRyaWJ1dGUtY2xlYXInOlxuICAgICAgICByZXR1cm4gJ2F0dHJpYnV0ZSc7XG4gICAgICBjYXNlICdrbm9iOnByb3BlcnR5LWNsZWFyJzpcbiAgICAgICAgcmV0dXJuICdwcm9wZXJ0eSc7XG4gICAgICBjYXNlICdrbm9iOmNzcy1wcm9wZXJ0eS1jbGVhcic6XG4gICAgICAgIHJldHVybiAnY3NzLXByb3BlcnR5JztcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIHJldHVybiAndW5rbm93bic7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwVHJlZVN0YXRlUGVyc2lzdGVuY2UoKSB7XG4gICAgdGhpcy4jaGFuZGxlVHJlZUV4cGFuZCA9IChlOiBFdmVudCkgPT4ge1xuICAgICAgaWYgKChlLnRhcmdldCBhcyBFbGVtZW50KT8udGFnTmFtZSAhPT0gJ1BGLVY2LVRSRUUtSVRFTScpIHJldHVybjtcblxuICAgICAgY29uc3Qgbm9kZUlkID0gdGhpcy4jZ2V0VHJlZU5vZGVJZChlLnRhcmdldCBhcyBFbGVtZW50KTtcbiAgICAgIGNvbnN0IHRyZWVTdGF0ZSA9IFN0YXRlUGVyc2lzdGVuY2UuZ2V0VHJlZVN0YXRlKCk7XG4gICAgICBpZiAoIXRyZWVTdGF0ZS5leHBhbmRlZC5pbmNsdWRlcyhub2RlSWQpKSB7XG4gICAgICAgIHRyZWVTdGF0ZS5leHBhbmRlZC5wdXNoKG5vZGVJZCk7XG4gICAgICAgIFN0YXRlUGVyc2lzdGVuY2Uuc2V0VHJlZVN0YXRlKHRyZWVTdGF0ZSk7XG4gICAgICB9XG4gICAgfTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2V4cGFuZCcsIHRoaXMuI2hhbmRsZVRyZWVFeHBhbmQpO1xuXG4gICAgdGhpcy4jaGFuZGxlVHJlZUNvbGxhcHNlID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgICBpZiAoKGUudGFyZ2V0IGFzIEVsZW1lbnQpPy50YWdOYW1lICE9PSAnUEYtVjYtVFJFRS1JVEVNJykgcmV0dXJuO1xuXG4gICAgICBjb25zdCBub2RlSWQgPSB0aGlzLiNnZXRUcmVlTm9kZUlkKGUudGFyZ2V0IGFzIEVsZW1lbnQpO1xuICAgICAgY29uc3QgdHJlZVN0YXRlID0gU3RhdGVQZXJzaXN0ZW5jZS5nZXRUcmVlU3RhdGUoKTtcbiAgICAgIGNvbnN0IGluZGV4ID0gdHJlZVN0YXRlLmV4cGFuZGVkLmluZGV4T2Yobm9kZUlkKTtcbiAgICAgIGlmIChpbmRleCA+IC0xKSB7XG4gICAgICAgIHRyZWVTdGF0ZS5leHBhbmRlZC5zcGxpY2UoaW5kZXgsIDEpO1xuICAgICAgICBTdGF0ZVBlcnNpc3RlbmNlLnNldFRyZWVTdGF0ZSh0cmVlU3RhdGUpO1xuICAgICAgfVxuICAgIH07XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdjb2xsYXBzZScsIHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSk7XG5cbiAgICB0aGlzLiNoYW5kbGVUcmVlU2VsZWN0ID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgICBpZiAoKGUudGFyZ2V0IGFzIEVsZW1lbnQpPy50YWdOYW1lICE9PSAnUEYtVjYtVFJFRS1JVEVNJykgcmV0dXJuO1xuXG4gICAgICBjb25zdCBub2RlSWQgPSB0aGlzLiNnZXRUcmVlTm9kZUlkKGUudGFyZ2V0IGFzIEVsZW1lbnQpO1xuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVUcmVlU3RhdGUoeyBzZWxlY3RlZDogbm9kZUlkIH0pO1xuICAgIH07XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdzZWxlY3QnLCB0aGlzLiNoYW5kbGVUcmVlU2VsZWN0KTtcblxuICAgIHRoaXMuI2FwcGx5VHJlZVN0YXRlKCk7XG4gIH1cblxuICAjYXBwbHlUcmVlU3RhdGUoKSB7XG4gICAgY29uc3QgdHJlZVN0YXRlID0gU3RhdGVQZXJzaXN0ZW5jZS5nZXRUcmVlU3RhdGUoKTtcblxuICAgIGZvciAoY29uc3Qgbm9kZUlkIG9mIHRyZWVTdGF0ZS5leHBhbmRlZCkge1xuICAgICAgY29uc3QgdHJlZUl0ZW0gPSB0aGlzLiNmaW5kVHJlZUl0ZW1CeUlkKG5vZGVJZCk7XG4gICAgICBpZiAodHJlZUl0ZW0gJiYgIXRyZWVJdGVtLmhhc0F0dHJpYnV0ZSgnZXhwYW5kZWQnKSkge1xuICAgICAgICB0cmVlSXRlbS5zZXRBdHRyaWJ1dGUoJ2V4cGFuZGVkJywgJycpO1xuICAgICAgfVxuICAgIH1cblxuICAgIGlmICh0cmVlU3RhdGUuc2VsZWN0ZWQpIHtcbiAgICAgIGNvbnN0IHRyZWVJdGVtID0gdGhpcy4jZmluZFRyZWVJdGVtQnlJZCh0cmVlU3RhdGUuc2VsZWN0ZWQpO1xuICAgICAgaWYgKHRyZWVJdGVtICYmICF0cmVlSXRlbS5oYXNBdHRyaWJ1dGUoJ2N1cnJlbnQnKSkge1xuICAgICAgICB0cmVlSXRlbS5zZXRBdHRyaWJ1dGUoJ2N1cnJlbnQnLCAnJyk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI3NldHVwU2lkZWJhclN0YXRlUGVyc2lzdGVuY2UoKSB7XG4gICAgY29uc3QgcGFnZSA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcigncGYtdjYtcGFnZScpO1xuXG4gICAgaWYgKCFwYWdlKSByZXR1cm47XG5cbiAgICBwYWdlLmFkZEV2ZW50TGlzdGVuZXIoJ3NpZGViYXItdG9nZ2xlJywgKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgICAgY29uc3QgY29sbGFwc2VkID0gIShldmVudCBhcyBhbnkpLmV4cGFuZGVkO1xuXG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHtcbiAgICAgICAgc2lkZWJhcjogeyBjb2xsYXBzZWQgfVxuICAgICAgfSk7XG4gICAgfSk7XG4gIH1cblxuICAjZmluZFRyZWVJdGVtQnlJZChub2RlSWQ6IHN0cmluZyk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICBjb25zdCBwYXJ0cyA9IG5vZGVJZC5zcGxpdCgnOicpO1xuICAgIGNvbnN0IFt0eXBlLCBtb2R1bGVQYXRoLCB0YWdOYW1lLCBuYW1lXSA9IHBhcnRzO1xuXG4gICAgbGV0IGF0dHJTdWZmaXggPSAnJztcbiAgICBpZiAodGFnTmFtZSkge1xuICAgICAgYXR0clN1ZmZpeCArPSBgW2RhdGEtdGFnLW5hbWU9XCIke0NTUy5lc2NhcGUodGFnTmFtZSl9XCJdYDtcbiAgICB9XG4gICAgaWYgKG5hbWUpIHtcbiAgICAgIGF0dHJTdWZmaXggKz0gYFtkYXRhLW5hbWU9XCIke0NTUy5lc2NhcGUobmFtZSl9XCJdYDtcbiAgICB9XG5cbiAgICBsZXQgc2VsZWN0b3IgPSBgcGYtdjYtdHJlZS1pdGVtW2RhdGEtdHlwZT1cIiR7Q1NTLmVzY2FwZSh0eXBlKX1cIl1gO1xuICAgIGlmIChtb2R1bGVQYXRoKSB7XG4gICAgICBjb25zdCBlc2NhcGVkTW9kdWxlUGF0aCA9IENTUy5lc2NhcGUobW9kdWxlUGF0aCk7XG4gICAgICBjb25zdCBlc2NhcGVkVHlwZSA9IENTUy5lc2NhcGUodHlwZSk7XG4gICAgICBjb25zdCBzZWxlY3RvcjEgPSBgcGYtdjYtdHJlZS1pdGVtW2RhdGEtdHlwZT1cIiR7ZXNjYXBlZFR5cGV9XCJdW2RhdGEtbW9kdWxlLXBhdGg9XCIke2VzY2FwZWRNb2R1bGVQYXRofVwiXSR7YXR0clN1ZmZpeH1gO1xuICAgICAgY29uc3Qgc2VsZWN0b3IyID0gYHBmLXY2LXRyZWUtaXRlbVtkYXRhLXR5cGU9XCIke2VzY2FwZWRUeXBlfVwiXVtkYXRhLXBhdGg9XCIke2VzY2FwZWRNb2R1bGVQYXRofVwiXSR7YXR0clN1ZmZpeH1gO1xuICAgICAgc2VsZWN0b3IgPSBgJHtzZWxlY3RvcjF9LCAke3NlbGVjdG9yMn1gO1xuICAgIH0gZWxzZSB7XG4gICAgICBzZWxlY3RvciArPSBhdHRyU3VmZml4O1xuICAgIH1cblxuICAgIHJldHVybiB0aGlzLnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpO1xuICB9XG5cbiAgI2dldFRyZWVOb2RlSWQodHJlZUl0ZW06IEVsZW1lbnQpOiBzdHJpbmcge1xuICAgIGNvbnN0IHR5cGUgPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtdHlwZScpO1xuICAgIGNvbnN0IG1vZHVsZVBhdGggPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtbW9kdWxlLXBhdGgnKSB8fCB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtcGF0aCcpO1xuICAgIGNvbnN0IHRhZ05hbWUgPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtdGFnLW5hbWUnKTtcbiAgICBjb25zdCBuYW1lID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLW5hbWUnKTtcbiAgICBjb25zdCBjYXRlZ29yeSA9IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS1jYXRlZ29yeScpO1xuXG4gICAgY29uc3QgcGFydHMgPSBbdHlwZV07XG4gICAgaWYgKG1vZHVsZVBhdGgpIHBhcnRzLnB1c2gobW9kdWxlUGF0aCk7XG4gICAgaWYgKHRhZ05hbWUpIHBhcnRzLnB1c2godGFnTmFtZSk7XG4gICAgaWYgKGNhdGVnb3J5KSB7XG4gICAgICBwYXJ0cy5wdXNoKGNhdGVnb3J5KTtcbiAgICB9IGVsc2UgaWYgKG5hbWUpIHtcbiAgICAgIHBhcnRzLnB1c2gobmFtZSk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHBhcnRzLmpvaW4oJzonKTtcbiAgfVxuXG4gIC8vIEV2ZW50IERpc2NvdmVyeSAmIENhcHR1cmUgTWV0aG9kc1xuXG4gIGFzeW5jICNkaXNjb3ZlckVsZW1lbnRFdmVudHMoKTogUHJvbWlzZTxNYXA8c3RyaW5nLCBFdmVudEluZm8+PiB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHJlc3BvbnNlID0gYXdhaXQgZmV0Y2goJy9jdXN0b20tZWxlbWVudHMuanNvbicpO1xuICAgICAgaWYgKCFyZXNwb25zZS5vaykge1xuICAgICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBObyBtYW5pZmVzdCBhdmFpbGFibGUgZm9yIGV2ZW50IGRpc2NvdmVyeScpO1xuICAgICAgICByZXR1cm4gbmV3IE1hcCgpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBtYW5pZmVzdCA9IGF3YWl0IHJlc3BvbnNlLmpzb24oKSBhcyBNYW5pZmVzdDtcbiAgICAgIHRoaXMuI21hbmlmZXN0ID0gbWFuaWZlc3Q7XG5cbiAgICAgIGNvbnN0IGV2ZW50TWFwID0gbmV3IE1hcDxzdHJpbmcsIEV2ZW50SW5mbz4oKTtcblxuICAgICAgZm9yIChjb25zdCBtb2R1bGUgb2YgbWFuaWZlc3QubW9kdWxlcyB8fCBbXSkge1xuICAgICAgICBmb3IgKGNvbnN0IGRlY2xhcmF0aW9uIG9mIG1vZHVsZS5kZWNsYXJhdGlvbnMgfHwgW10pIHtcbiAgICAgICAgICBpZiAoZGVjbGFyYXRpb24uY3VzdG9tRWxlbWVudCAmJiBkZWNsYXJhdGlvbi50YWdOYW1lKSB7XG4gICAgICAgICAgICBjb25zdCB0YWdOYW1lID0gZGVjbGFyYXRpb24udGFnTmFtZTtcbiAgICAgICAgICAgIGNvbnN0IGV2ZW50cyA9IGRlY2xhcmF0aW9uLmV2ZW50cyB8fCBbXTtcblxuICAgICAgICAgICAgaWYgKGV2ZW50cy5sZW5ndGggPiAwKSB7XG4gICAgICAgICAgICAgIGNvbnN0IGV2ZW50TmFtZXMgPSBuZXcgU2V0KGV2ZW50cy5tYXAoZSA9PiBlLm5hbWUpKTtcbiAgICAgICAgICAgICAgZXZlbnRNYXAuc2V0KHRhZ05hbWUsIHtcbiAgICAgICAgICAgICAgICBldmVudE5hbWVzLFxuICAgICAgICAgICAgICAgIGV2ZW50czogZXZlbnRzXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICByZXR1cm4gZXZlbnRNYXA7XG4gICAgfSBjYXRjaCAoZXJyb3IpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIEVycm9yIGxvYWRpbmcgbWFuaWZlc3QgZm9yIGV2ZW50IGRpc2NvdmVyeTonLCBlcnJvcik7XG4gICAgICByZXR1cm4gbmV3IE1hcCgpO1xuICAgIH1cbiAgfVxuXG4gIGFzeW5jICNzZXR1cEV2ZW50Q2FwdHVyZSgpIHtcbiAgICB0aGlzLiNlbGVtZW50RXZlbnRNYXAgPSBhd2FpdCB0aGlzLiNkaXNjb3ZlckVsZW1lbnRFdmVudHMoKTtcblxuICAgIGlmICh0aGlzLiNlbGVtZW50RXZlbnRNYXAuc2l6ZSA9PT0gMCkgcmV0dXJuO1xuXG4gICAgdGhpcy4jYXR0YWNoRXZlbnRMaXN0ZW5lcnMoKTtcbiAgICB0aGlzLiN1cGRhdGVFdmVudEZpbHRlcnMoKTtcbiAgICB0aGlzLiNvYnNlcnZlRGVtb011dGF0aW9ucygpO1xuICB9XG5cbiAgI2F0dGFjaEV2ZW50TGlzdGVuZXJzKCkge1xuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCByb290ID0gZGVtby5zaGFkb3dSb290ID8/IGRlbW87XG5cbiAgICBmb3IgKGNvbnN0IFt0YWdOYW1lLCBldmVudEluZm9dIG9mIHRoaXMuI2VsZW1lbnRFdmVudE1hcCEpIHtcbiAgICAgIGNvbnN0IGVsZW1lbnRzID0gcm9vdC5xdWVyeVNlbGVjdG9yQWxsKHRhZ05hbWUpO1xuXG4gICAgICBmb3IgKGNvbnN0IGVsZW1lbnQgb2YgZWxlbWVudHMpIHtcbiAgICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgZXZlbnRJbmZvLmV2ZW50TmFtZXMpIHtcbiAgICAgICAgICBlbGVtZW50LmFkZEV2ZW50TGlzdGVuZXIoZXZlbnROYW1lLCB0aGlzLiNoYW5kbGVFbGVtZW50RXZlbnQsIHsgY2FwdHVyZTogdHJ1ZSB9KTtcbiAgICAgICAgfVxuICAgICAgICAoZWxlbWVudCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5jZW1FdmVudHNBdHRhY2hlZCA9ICd0cnVlJztcbiAgICAgICAgdGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjb2JzZXJ2ZURlbW9NdXRhdGlvbnMoKSB7XG4gICAgY29uc3QgZGVtbyA9IHRoaXMuZGVtbztcbiAgICBpZiAoIWRlbW8pIHJldHVybjtcblxuICAgIGNvbnN0IHJvb3QgPSBkZW1vLnNoYWRvd1Jvb3QgPz8gZGVtbztcblxuICAgIHRoaXMuI29ic2VydmVyLm9ic2VydmUocm9vdCwge1xuICAgICAgY2hpbGRMaXN0OiB0cnVlLFxuICAgICAgc3VidHJlZTogdHJ1ZVxuICAgIH0pO1xuICB9XG5cbiAgI2hhbmRsZUVsZW1lbnRFdmVudCA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCBlbGVtZW50ID0gZXZlbnQuY3VycmVudFRhcmdldDtcbiAgICBpZiAoIShlbGVtZW50IGluc3RhbmNlb2YgSFRNTEVsZW1lbnQpKSByZXR1cm47XG5cbiAgICBjb25zdCB0YWdOYW1lID0gZWxlbWVudC50YWdOYW1lLnRvTG93ZXJDYXNlKCk7XG4gICAgY29uc3QgZXZlbnRJbmZvID0gdGhpcy4jZWxlbWVudEV2ZW50TWFwPy5nZXQodGFnTmFtZSk7XG5cbiAgICBpZiAoIWV2ZW50SW5mbyB8fCAhZXZlbnRJbmZvLmV2ZW50TmFtZXMuaGFzKGV2ZW50LnR5cGUpKSByZXR1cm47XG5cbiAgICB0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuYWRkKHRhZ05hbWUpO1xuICAgIHRoaXMuI2NhcHR1cmVFdmVudChldmVudCwgZWxlbWVudCwgdGFnTmFtZSwgZXZlbnRJbmZvKTtcbiAgfTtcblxuICAjZ2V0RXZlbnREb2N1bWVudGF0aW9uKG1hbmlmZXN0RXZlbnQ6IEV2ZW50SW5mb1snZXZlbnRzJ11bMF0gfCB1bmRlZmluZWQpIHtcbiAgICBpZiAoIW1hbmlmZXN0RXZlbnQpIHtcbiAgICAgIHJldHVybiB7IHN1bW1hcnk6IG51bGwsIGRlc2NyaXB0aW9uOiBudWxsIH07XG4gICAgfVxuXG4gICAgbGV0IHN1bW1hcnkgPSBtYW5pZmVzdEV2ZW50LnN1bW1hcnkgfHwgbnVsbDtcbiAgICBsZXQgZGVzY3JpcHRpb24gPSBtYW5pZmVzdEV2ZW50LmRlc2NyaXB0aW9uIHx8IG51bGw7XG5cbiAgICBpZiAobWFuaWZlc3RFdmVudC50eXBlPy50ZXh0ICYmIHRoaXMuI21hbmlmZXN0KSB7XG4gICAgICBjb25zdCB0eXBlTmFtZSA9IG1hbmlmZXN0RXZlbnQudHlwZS50ZXh0O1xuICAgICAgY29uc3QgdHlwZURlY2xhcmF0aW9uID0gdGhpcy4jZmluZFR5cGVEZWNsYXJhdGlvbih0eXBlTmFtZSk7XG5cbiAgICAgIGlmICh0eXBlRGVjbGFyYXRpb24pIHtcbiAgICAgICAgaWYgKCFzdW1tYXJ5ICYmIHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5KSB7XG4gICAgICAgICAgc3VtbWFyeSA9IHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5O1xuICAgICAgICB9IGVsc2UgaWYgKHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5ICYmIHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5ICE9PSBzdW1tYXJ5KSB7XG4gICAgICAgICAgc3VtbWFyeSA9IHN1bW1hcnkgPyBgJHtzdW1tYXJ5fVxcblxcbkZyb20gJHt0eXBlTmFtZX06ICR7dHlwZURlY2xhcmF0aW9uLnN1bW1hcnl9YCA6IHR5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5O1xuICAgICAgICB9XG5cbiAgICAgICAgaWYgKCFkZXNjcmlwdGlvbiAmJiB0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb24pIHtcbiAgICAgICAgICBkZXNjcmlwdGlvbiA9IHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbjtcbiAgICAgICAgfSBlbHNlIGlmICh0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb24gJiYgdHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uICE9PSBkZXNjcmlwdGlvbikge1xuICAgICAgICAgIGRlc2NyaXB0aW9uID0gZGVzY3JpcHRpb24gPyBgJHtkZXNjcmlwdGlvbn1cXG5cXG4ke3R5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbn1gIDogdHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIHsgc3VtbWFyeSwgZGVzY3JpcHRpb24gfTtcbiAgfVxuXG4gICNmaW5kVHlwZURlY2xhcmF0aW9uKHR5cGVOYW1lOiBzdHJpbmcpIHtcbiAgICBpZiAoIXRoaXMuI21hbmlmZXN0KSByZXR1cm4gbnVsbDtcblxuICAgIGZvciAoY29uc3QgbW9kdWxlIG9mIHRoaXMuI21hbmlmZXN0Lm1vZHVsZXMgfHwgW10pIHtcbiAgICAgIGZvciAoY29uc3QgZGVjbGFyYXRpb24gb2YgbW9kdWxlLmRlY2xhcmF0aW9ucyB8fCBbXSkge1xuICAgICAgICBpZiAoZGVjbGFyYXRpb24ubmFtZSA9PT0gdHlwZU5hbWUgJiZcbiAgICAgICAgICAgIChkZWNsYXJhdGlvbi5raW5kID09PSAnY2xhc3MnIHx8IGRlY2xhcmF0aW9uLmtpbmQgPT09ICdpbnRlcmZhY2UnKSkge1xuICAgICAgICAgIHJldHVybiBkZWNsYXJhdGlvbiBhcyB7IHN1bW1hcnk/OiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH07XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICByZXR1cm4gbnVsbDtcbiAgfVxuXG4gICNjYXB0dXJlRXZlbnQoZXZlbnQ6IEV2ZW50LCB0YXJnZXQ6IEhUTUxFbGVtZW50LCB0YWdOYW1lOiBzdHJpbmcsIGV2ZW50SW5mbzogRXZlbnRJbmZvKSB7XG4gICAgY29uc3QgbWFuaWZlc3RFdmVudCA9IGV2ZW50SW5mby5ldmVudHMuZmluZChlID0+IGUubmFtZSA9PT0gZXZlbnQudHlwZSk7XG5cbiAgICBjb25zdCBldmVudERvY3MgPSB0aGlzLiNnZXRFdmVudERvY3VtZW50YXRpb24obWFuaWZlc3RFdmVudCk7XG5cbiAgICBjb25zdCBjdXN0b21Qcm9wZXJ0aWVzID0gdGhpcy4jZXh0cmFjdEV2ZW50UHJvcGVydGllcyhldmVudCk7XG5cbiAgICBjb25zdCBldmVudFJlY29yZDogRXZlbnRSZWNvcmQgPSB7XG4gICAgICBpZDogYCR7RGF0ZS5ub3coKX0tJHtNYXRoLnJhbmRvbSgpfWAsXG4gICAgICB0aW1lc3RhbXA6IG5ldyBEYXRlKCksXG4gICAgICBldmVudE5hbWU6IGV2ZW50LnR5cGUsXG4gICAgICB0YWdOYW1lOiB0YWdOYW1lLFxuICAgICAgZWxlbWVudElkOiB0YXJnZXQuaWQgfHwgbnVsbCxcbiAgICAgIGVsZW1lbnRDbGFzczogdGFyZ2V0LmNsYXNzTmFtZSB8fCBudWxsLFxuICAgICAgY3VzdG9tUHJvcGVydGllczogY3VzdG9tUHJvcGVydGllcyxcbiAgICAgIG1hbmlmZXN0VHlwZTogbWFuaWZlc3RFdmVudD8udHlwZT8udGV4dCB8fCBudWxsLFxuICAgICAgc3VtbWFyeTogZXZlbnREb2NzLnN1bW1hcnksXG4gICAgICBkZXNjcmlwdGlvbjogZXZlbnREb2NzLmRlc2NyaXB0aW9uLFxuICAgICAgYnViYmxlczogZXZlbnQuYnViYmxlcyxcbiAgICAgIGNvbXBvc2VkOiBldmVudC5jb21wb3NlZCxcbiAgICAgIGNhbmNlbGFibGU6IGV2ZW50LmNhbmNlbGFibGUsXG4gICAgICBkZWZhdWx0UHJldmVudGVkOiBldmVudC5kZWZhdWx0UHJldmVudGVkXG4gICAgfTtcblxuICAgIHRoaXMuI2NhcHR1cmVkRXZlbnRzLnB1c2goZXZlbnRSZWNvcmQpO1xuXG4gICAgaWYgKHRoaXMuI2NhcHR1cmVkRXZlbnRzLmxlbmd0aCA+IHRoaXMuI21heENhcHR1cmVkRXZlbnRzKSB7XG4gICAgICB0aGlzLiNjYXB0dXJlZEV2ZW50cy5zaGlmdCgpO1xuICAgIH1cblxuICAgIHRoaXMuI3JlbmRlckV2ZW50KGV2ZW50UmVjb3JkKTtcbiAgfVxuXG4gICNleHRyYWN0RXZlbnRQcm9wZXJ0aWVzKGV2ZW50OiBFdmVudCk6IFJlY29yZDxzdHJpbmcsIHVua25vd24+IHtcbiAgICBjb25zdCBwcm9wZXJ0aWVzOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiA9IHt9O1xuICAgIGNvbnN0IGV2ZW50UHJvdG90eXBlS2V5cyA9IG5ldyBTZXQoT2JqZWN0LmdldE93blByb3BlcnR5TmFtZXMoRXZlbnQucHJvdG90eXBlKSk7XG5cbiAgICBjb25zdCBzZXJpYWxpemVWYWx1ZSA9ICh2YWx1ZTogdW5rbm93bik6IHVua25vd24gPT4ge1xuICAgICAgdHJ5IHtcbiAgICAgICAgcmV0dXJuIEpTT04ucGFyc2UoSlNPTi5zdHJpbmdpZnkodmFsdWUpKTtcbiAgICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgICAgdHJ5IHtcbiAgICAgICAgICByZXR1cm4gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgfSBjYXRjaCAoc3RyaW5nRXJyKSB7XG4gICAgICAgICAgcmV0dXJuICdbTm90IHNlcmlhbGl6YWJsZV0nO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfTtcblxuICAgIGlmIChldmVudCBpbnN0YW5jZW9mIEN1c3RvbUV2ZW50ICYmIGV2ZW50LmRldGFpbCAhPT0gdW5kZWZpbmVkKSB7XG4gICAgICBwcm9wZXJ0aWVzLmRldGFpbCA9IHNlcmlhbGl6ZVZhbHVlKGV2ZW50LmRldGFpbCk7XG4gICAgfVxuXG4gICAgZm9yIChjb25zdCBrZXkgb2YgT2JqZWN0LmdldE93blByb3BlcnR5TmFtZXMoZXZlbnQpKSB7XG4gICAgICBpZiAoIWV2ZW50UHJvdG90eXBlS2V5cy5oYXMoa2V5KSAmJiAha2V5LnN0YXJ0c1dpdGgoJ18nKSAmJiAhcHJvcGVydGllcy5oYXNPd25Qcm9wZXJ0eShrZXkpKSB7XG4gICAgICAgIHByb3BlcnRpZXNba2V5XSA9IHNlcmlhbGl6ZVZhbHVlKChldmVudCBhcyBhbnkpW2tleV0pO1xuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiBwcm9wZXJ0aWVzO1xuICB9XG5cbiAgI3JlbmRlckV2ZW50KGV2ZW50UmVjb3JkOiBFdmVudFJlY29yZCkge1xuICAgIGlmICghdGhpcy4jZXZlbnRMaXN0KSByZXR1cm47XG5cbiAgICBjb25zdCBmcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNldmVudEVudHJ5VGVtcGxhdGUuY29udGVudC5jbG9uZU5vZGUodHJ1ZSkgYXMgRG9jdW1lbnRGcmFnbWVudDtcblxuICAgIGNvbnN0IHRpbWUgPSBldmVudFJlY29yZC50aW1lc3RhbXAudG9Mb2NhbGVUaW1lU3RyaW5nKCk7XG5cbiAgICBjb25zdCBjb250YWluZXIgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNvbnRhaW5lclwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgIGNvbnRhaW5lci5kYXRhc2V0LmV2ZW50SWQgPSBldmVudFJlY29yZC5pZDtcbiAgICBjb250YWluZXIuZGF0YXNldC5ldmVudFR5cGUgPSBldmVudFJlY29yZC5ldmVudE5hbWU7XG4gICAgY29udGFpbmVyLmRhdGFzZXQuZWxlbWVudFR5cGUgPSBldmVudFJlY29yZC50YWdOYW1lO1xuXG4gICAgY29uc3QgdGV4dE1hdGNoID0gdGhpcy4jZXZlbnRNYXRjaGVzVGV4dEZpbHRlcihldmVudFJlY29yZCk7XG4gICAgY29uc3QgdHlwZU1hdGNoID0gdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5zaXplID09PSAwIHx8IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuaGFzKGV2ZW50UmVjb3JkLmV2ZW50TmFtZSk7XG4gICAgY29uc3QgZWxlbWVudE1hdGNoID0gdGhpcy4jZWxlbWVudEZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNlbGVtZW50RmlsdGVycy5oYXMoZXZlbnRSZWNvcmQudGFnTmFtZSk7XG5cbiAgICBpZiAoISh0ZXh0TWF0Y2ggJiYgdHlwZU1hdGNoICYmIGVsZW1lbnRNYXRjaCkpIHtcbiAgICAgIGNvbnRhaW5lci5zZXRBdHRyaWJ1dGUoJ2hpZGRlbicsICcnKTtcbiAgICB9XG5cbiAgICBjb25zdCBsYWJlbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibGFiZWxcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICBsYWJlbC50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLmV2ZW50TmFtZTtcbiAgICBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICdpbmZvJyk7XG5cbiAgICBjb25zdCB0aW1lRWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cInRpbWVcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICB0aW1lRWwuc2V0QXR0cmlidXRlKCdkYXRldGltZScsIGV2ZW50UmVjb3JkLnRpbWVzdGFtcC50b0lTT1N0cmluZygpKTtcbiAgICB0aW1lRWwudGV4dENvbnRlbnQgPSB0aW1lO1xuXG4gICAgY29uc3QgZWxlbWVudEVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJlbGVtZW50XCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgbGV0IGVsZW1lbnRUZXh0ID0gYDwke2V2ZW50UmVjb3JkLnRhZ05hbWV9PmA7XG4gICAgaWYgKGV2ZW50UmVjb3JkLmVsZW1lbnRJZCkge1xuICAgICAgZWxlbWVudFRleHQgKz0gYCMke2V2ZW50UmVjb3JkLmVsZW1lbnRJZH1gO1xuICAgIH1cbiAgICBlbGVtZW50RWwudGV4dENvbnRlbnQgPSBlbGVtZW50VGV4dDtcblxuICAgIHRoaXMuI2V2ZW50TGlzdC5hcHBlbmQoZnJhZ21lbnQpO1xuXG4gICAgaWYgKCF0aGlzLiNzZWxlY3RlZEV2ZW50SWQpIHtcbiAgICAgIHRoaXMuI3NlbGVjdEV2ZW50KGV2ZW50UmVjb3JkLmlkKTtcbiAgICB9XG5cbiAgICBpZiAodGhpcy4jZHJhd2VyT3BlbiAmJiB0aGlzLiNpc0V2ZW50c1RhYkFjdGl2ZSgpKSB7XG4gICAgICB0aGlzLiNzY3JvbGxFdmVudHNUb0JvdHRvbSgpO1xuICAgIH1cbiAgfVxuXG4gICNzZWxlY3RFdmVudChldmVudElkOiBzdHJpbmcpIHtcbiAgICBjb25zdCBldmVudFJlY29yZCA9IHRoaXMuI2dldEV2ZW50UmVjb3JkQnlJZChldmVudElkKTtcbiAgICBpZiAoIWV2ZW50UmVjb3JkKSByZXR1cm47XG5cbiAgICB0aGlzLiNzZWxlY3RlZEV2ZW50SWQgPSBldmVudElkO1xuXG4gICAgY29uc3QgYWxsSXRlbXMgPSB0aGlzLiNldmVudExpc3Q/LnF1ZXJ5U2VsZWN0b3JBbGwoJy5ldmVudC1saXN0LWl0ZW0nKTtcbiAgICBhbGxJdGVtcz8uZm9yRWFjaChpdGVtID0+IHtcbiAgICAgIGlmICgoaXRlbSBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5ldmVudElkID09PSBldmVudElkKSB7XG4gICAgICAgIGl0ZW0uY2xhc3NMaXN0LmFkZCgnc2VsZWN0ZWQnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2FyaWEtc2VsZWN0ZWQnLCAndHJ1ZScpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgaXRlbS5jbGFzc0xpc3QucmVtb3ZlKCdzZWxlY3RlZCcpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnYXJpYS1zZWxlY3RlZCcsICdmYWxzZScpO1xuICAgICAgfVxuICAgIH0pO1xuXG4gICAgaWYgKHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyKSB7XG4gICAgICB0aGlzLiNldmVudERldGFpbEhlYWRlci5pbm5lckhUTUwgPSAnJztcblxuICAgICAgY29uc3QgaGVhZGVyQ29udGVudCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuICAgICAgaGVhZGVyQ29udGVudC5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLWhlYWRlci1jb250ZW50JztcblxuICAgICAgY29uc3QgZXZlbnROYW1lID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnaDMnKTtcbiAgICAgIGV2ZW50TmFtZS50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLmV2ZW50TmFtZTtcbiAgICAgIGV2ZW50TmFtZS5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLW5hbWUnO1xuICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChldmVudE5hbWUpO1xuXG4gICAgICBpZiAoZXZlbnRSZWNvcmQuc3VtbWFyeSkge1xuICAgICAgICBjb25zdCBzdW1tYXJ5ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncCcpO1xuICAgICAgICBzdW1tYXJ5LnRleHRDb250ZW50ID0gZXZlbnRSZWNvcmQuc3VtbWFyeTtcbiAgICAgICAgc3VtbWFyeS5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXN1bW1hcnknO1xuICAgICAgICBoZWFkZXJDb250ZW50LmFwcGVuZENoaWxkKHN1bW1hcnkpO1xuICAgICAgfVxuXG4gICAgICBpZiAoZXZlbnRSZWNvcmQuZGVzY3JpcHRpb24pIHtcbiAgICAgICAgY29uc3QgZGVzY3JpcHRpb24gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwJyk7XG4gICAgICAgIGRlc2NyaXB0aW9uLnRleHRDb250ZW50ID0gZXZlbnRSZWNvcmQuZGVzY3JpcHRpb247XG4gICAgICAgIGRlc2NyaXB0aW9uLmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtZGVzY3JpcHRpb24nO1xuICAgICAgICBoZWFkZXJDb250ZW50LmFwcGVuZENoaWxkKGRlc2NyaXB0aW9uKTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgbWV0YSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuICAgICAgbWV0YS5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLW1ldGEnO1xuXG4gICAgICBjb25zdCB0aW1lRWwgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0aW1lJyk7XG4gICAgICB0aW1lRWwuc2V0QXR0cmlidXRlKCdkYXRldGltZScsIGV2ZW50UmVjb3JkLnRpbWVzdGFtcC50b0lTT1N0cmluZygpKTtcbiAgICAgIHRpbWVFbC50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLnRpbWVzdGFtcC50b0xvY2FsZVRpbWVTdHJpbmcoKTtcbiAgICAgIHRpbWVFbC5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXRpbWUnO1xuXG4gICAgICBjb25zdCBlbGVtZW50ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgbGV0IGVsZW1lbnRUZXh0ID0gYDwke2V2ZW50UmVjb3JkLnRhZ05hbWV9PmA7XG4gICAgICBpZiAoZXZlbnRSZWNvcmQuZWxlbWVudElkKSB7XG4gICAgICAgIGVsZW1lbnRUZXh0ICs9IGAjJHtldmVudFJlY29yZC5lbGVtZW50SWR9YDtcbiAgICAgIH1cbiAgICAgIGVsZW1lbnQudGV4dENvbnRlbnQgPSBlbGVtZW50VGV4dDtcbiAgICAgIGVsZW1lbnQuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1lbGVtZW50JztcblxuICAgICAgbWV0YS5hcHBlbmRDaGlsZCh0aW1lRWwpO1xuICAgICAgbWV0YS5hcHBlbmRDaGlsZChlbGVtZW50KTtcblxuICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChtZXRhKTtcblxuICAgICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIuYXBwZW5kQ2hpbGQoaGVhZGVyQ29udGVudCk7XG4gICAgfVxuXG4gICAgaWYgKHRoaXMuI2V2ZW50RGV0YWlsQm9keSkge1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5LmlubmVySFRNTCA9ICcnO1xuXG4gICAgICBjb25zdCBwcm9wZXJ0aWVzSGVhZGluZyA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2g0Jyk7XG4gICAgICBwcm9wZXJ0aWVzSGVhZGluZy50ZXh0Q29udGVudCA9ICdQcm9wZXJ0aWVzJztcbiAgICAgIHByb3BlcnRpZXNIZWFkaW5nLmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtcHJvcGVydGllcy1oZWFkaW5nJztcblxuICAgICAgY29uc3QgcHJvcGVydGllc0NvbnRhaW5lciA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuICAgICAgcHJvcGVydGllc0NvbnRhaW5lci5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMnO1xuXG4gICAgICBjb25zdCBldmVudFByb3BlcnRpZXMgPSB0aGlzLiNidWlsZFByb3BlcnRpZXNGb3JEaXNwbGF5KGV2ZW50UmVjb3JkKTtcbiAgICAgIGlmIChPYmplY3Qua2V5cyhldmVudFByb3BlcnRpZXMpLmxlbmd0aCA+IDApIHtcbiAgICAgICAgcHJvcGVydGllc0NvbnRhaW5lci5hcHBlbmRDaGlsZCh0aGlzLiNidWlsZFByb3BlcnR5VHJlZShldmVudFByb3BlcnRpZXMpKTtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHByb3BlcnRpZXNDb250YWluZXIudGV4dENvbnRlbnQgPSAnTm8gcHJvcGVydGllcyB0byBkaXNwbGF5JztcbiAgICAgIH1cblxuICAgICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5LmFwcGVuZENoaWxkKHByb3BlcnRpZXNIZWFkaW5nKTtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5hcHBlbmRDaGlsZChwcm9wZXJ0aWVzQ29udGFpbmVyKTtcbiAgICB9XG4gIH1cblxuICAjYnVpbGRQcm9wZXJ0aWVzRm9yRGlzcGxheShldmVudFJlY29yZDogRXZlbnRSZWNvcmQpOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiB7XG4gICAgY29uc3QgcHJvcGVydGllczogUmVjb3JkPHN0cmluZywgdW5rbm93bj4gPSB7fTtcblxuICAgIGlmIChldmVudFJlY29yZC5jdXN0b21Qcm9wZXJ0aWVzKSB7XG4gICAgICBPYmplY3QuYXNzaWduKHByb3BlcnRpZXMsIGV2ZW50UmVjb3JkLmN1c3RvbVByb3BlcnRpZXMpO1xuICAgIH1cblxuICAgIHByb3BlcnRpZXMuYnViYmxlcyA9IGV2ZW50UmVjb3JkLmJ1YmJsZXM7XG4gICAgcHJvcGVydGllcy5jYW5jZWxhYmxlID0gZXZlbnRSZWNvcmQuY2FuY2VsYWJsZTtcbiAgICBwcm9wZXJ0aWVzLmRlZmF1bHRQcmV2ZW50ZWQgPSBldmVudFJlY29yZC5kZWZhdWx0UHJldmVudGVkO1xuICAgIHByb3BlcnRpZXMuY29tcG9zZWQgPSBldmVudFJlY29yZC5jb21wb3NlZDtcblxuICAgIGlmIChldmVudFJlY29yZC5tYW5pZmVzdFR5cGUpIHtcbiAgICAgIHByb3BlcnRpZXMudHlwZSA9IGV2ZW50UmVjb3JkLm1hbmlmZXN0VHlwZTtcbiAgICB9XG5cbiAgICByZXR1cm4gcHJvcGVydGllcztcbiAgfVxuXG4gICNidWlsZFByb3BlcnR5VHJlZShvYmo6IFJlY29yZDxzdHJpbmcsIHVua25vd24+LCBkZXB0aCA9IDApOiBIVE1MVUxpc3RFbGVtZW50IHtcbiAgICBjb25zdCB1bCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3VsJyk7XG4gICAgdWwuY2xhc3NOYW1lID0gJ2V2ZW50LXByb3BlcnR5LXRyZWUnO1xuICAgIGlmIChkZXB0aCA+IDApIHtcbiAgICAgIHVsLmNsYXNzTGlzdC5hZGQoJ25lc3RlZCcpO1xuICAgIH1cblxuICAgIGZvciAoY29uc3QgW2tleSwgdmFsdWVdIG9mIE9iamVjdC5lbnRyaWVzKG9iaikpIHtcbiAgICAgIGNvbnN0IGxpID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnbGknKTtcbiAgICAgIGxpLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS1pdGVtJztcblxuICAgICAgY29uc3Qga2V5U3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgIGtleVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LWtleSc7XG4gICAgICBrZXlTcGFuLnRleHRDb250ZW50ID0ga2V5O1xuXG4gICAgICBjb25zdCBjb2xvblNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICBjb2xvblNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LWNvbG9uJztcbiAgICAgIGNvbG9uU3Bhbi50ZXh0Q29udGVudCA9ICc6ICc7XG5cbiAgICAgIGxpLmFwcGVuZENoaWxkKGtleVNwYW4pO1xuICAgICAgbGkuYXBwZW5kQ2hpbGQoY29sb25TcGFuKTtcblxuICAgICAgaWYgKHZhbHVlID09PSBudWxsIHx8IHZhbHVlID09PSB1bmRlZmluZWQpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIG51bGwnO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdib29sZWFuJykge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUgYm9vbGVhbic7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ251bWJlcicpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIG51bWJlcic7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ3N0cmluZycpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIHN0cmluZyc7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IGBcIiR7dmFsdWV9XCJgO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmIChBcnJheS5pc0FycmF5KHZhbHVlKSkge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUgYXJyYXknO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBgQXJyYXkoJHt2YWx1ZS5sZW5ndGh9KWA7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG5cbiAgICAgICAgaWYgKHZhbHVlLmxlbmd0aCA+IDAgJiYgZGVwdGggPCAzKSB7XG4gICAgICAgICAgY29uc3QgbmVzdGVkT2JqOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiA9IHt9O1xuICAgICAgICAgIHZhbHVlLmZvckVhY2goKGl0ZW0sIGluZGV4KSA9PiB7XG4gICAgICAgICAgICBuZXN0ZWRPYmpbaW5kZXhdID0gaXRlbTtcbiAgICAgICAgICB9KTtcbiAgICAgICAgICBsaS5hcHBlbmRDaGlsZCh0aGlzLiNidWlsZFByb3BlcnR5VHJlZShuZXN0ZWRPYmosIGRlcHRoICsgMSkpO1xuICAgICAgICB9XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ29iamVjdCcpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIG9iamVjdCc7XG4gICAgICAgIGNvbnN0IGtleXMgPSBPYmplY3Qua2V5cyh2YWx1ZSBhcyBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPik7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IGtleXMubGVuZ3RoID4gMCA/IGBPYmplY3RgIDogYHt9YDtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcblxuICAgICAgICBpZiAoa2V5cy5sZW5ndGggPiAwICYmIGRlcHRoIDwgMykge1xuICAgICAgICAgIGxpLmFwcGVuZENoaWxkKHRoaXMuI2J1aWxkUHJvcGVydHlUcmVlKHZhbHVlIGFzIFJlY29yZDxzdHJpbmcsIHVua25vd24+LCBkZXB0aCArIDEpKTtcbiAgICAgICAgfVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlJztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH1cblxuICAgICAgdWwuYXBwZW5kQ2hpbGQobGkpO1xuICAgIH1cblxuICAgIHJldHVybiB1bDtcbiAgfVxuXG4gICNzY3JvbGxFdmVudHNUb0JvdHRvbSgpIHtcbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgcmVxdWVzdEFuaW1hdGlvbkZyYW1lKCgpID0+IHtcbiAgICAgIGNvbnN0IGxhc3RFdmVudCA9IHRoaXMuI2V2ZW50TGlzdCEubGFzdEVsZW1lbnRDaGlsZDtcbiAgICAgIGlmIChsYXN0RXZlbnQpIHtcbiAgICAgICAgbGFzdEV2ZW50LnNjcm9sbEludG9WaWV3KHsgYmVoYXZpb3I6ICdhdXRvJywgYmxvY2s6ICdlbmQnIH0pO1xuICAgICAgfVxuICAgIH0pO1xuICB9XG5cbiAgI2lzRXZlbnRzVGFiQWN0aXZlKCk6IGJvb2xlYW4ge1xuICAgIGNvbnN0IHRhYnMgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ3BmLXY2LXRhYnMnKTtcbiAgICBpZiAoIXRhYnMpIHJldHVybiBmYWxzZTtcblxuICAgIGNvbnN0IHNlbGVjdGVkSW5kZXggPSBwYXJzZUludCh0YWJzLmdldEF0dHJpYnV0ZSgnc2VsZWN0ZWQnKSB8fCAnMCcsIDEwKTtcbiAgICByZXR1cm4gc2VsZWN0ZWRJbmRleCA9PT0gMztcbiAgfVxuXG4gICNmaWx0ZXJFdmVudHMocXVlcnk6IHN0cmluZykge1xuICAgIHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlID0gcXVlcnkudG9Mb3dlckNhc2UoKTtcblxuICAgIGlmICghdGhpcy4jZXZlbnRMaXN0KSByZXR1cm47XG5cbiAgICBmb3IgKGNvbnN0IGVudHJ5IG9mIHRoaXMuI2V2ZW50TGlzdC5jaGlsZHJlbikge1xuICAgICAgY29uc3QgZXZlbnRSZWNvcmQgPSB0aGlzLiNnZXRFdmVudFJlY29yZEJ5SWQoKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LmV2ZW50SWQhKTtcblxuICAgICAgaWYgKCFldmVudFJlY29yZCkgY29udGludWU7XG5cbiAgICAgIGNvbnN0IHRleHRNYXRjaCA9IHRoaXMuI2V2ZW50TWF0Y2hlc1RleHRGaWx0ZXIoZXZlbnRSZWNvcmQpO1xuICAgICAgY29uc3QgdHlwZU1hdGNoID0gdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5zaXplID09PSAwIHx8IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuaGFzKGV2ZW50UmVjb3JkLmV2ZW50TmFtZSk7XG4gICAgICBjb25zdCBlbGVtZW50TWF0Y2ggPSB0aGlzLiNlbGVtZW50RmlsdGVycy5zaXplID09PSAwIHx8IHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmhhcyhldmVudFJlY29yZC50YWdOYW1lKTtcblxuICAgICAgKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5oaWRkZW4gPSAhKHRleHRNYXRjaCAmJiB0eXBlTWF0Y2ggJiYgZWxlbWVudE1hdGNoKTtcbiAgICB9XG4gIH1cblxuICAjZXZlbnRNYXRjaGVzVGV4dEZpbHRlcihldmVudFJlY29yZDogRXZlbnRSZWNvcmQpOiBib29sZWFuIHtcbiAgICBpZiAoIXRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKSByZXR1cm4gdHJ1ZTtcblxuICAgIGNvbnN0IHNlYXJjaFRleHQgPSBbXG4gICAgICBldmVudFJlY29yZC50YWdOYW1lLFxuICAgICAgZXZlbnRSZWNvcmQuZXZlbnROYW1lLFxuICAgICAgZXZlbnRSZWNvcmQuZWxlbWVudElkIHx8ICcnLFxuICAgICAgSlNPTi5zdHJpbmdpZnkoZXZlbnRSZWNvcmQuY3VzdG9tUHJvcGVydGllcyB8fCB7fSlcbiAgICBdLmpvaW4oJyAnKS50b0xvd2VyQ2FzZSgpO1xuXG4gICAgcmV0dXJuIHNlYXJjaFRleHQuaW5jbHVkZXModGhpcy4jZXZlbnRzRmlsdGVyVmFsdWUpO1xuICB9XG5cbiAgI2dldEV2ZW50UmVjb3JkQnlJZChpZDogc3RyaW5nKTogRXZlbnRSZWNvcmQgfCB1bmRlZmluZWQge1xuICAgIHJldHVybiB0aGlzLiNjYXB0dXJlZEV2ZW50cy5maW5kKGUgPT4gZS5pZCA9PT0gaWQpO1xuICB9XG5cbiAgI3VwZGF0ZUV2ZW50RmlsdGVycygpIHtcbiAgICBjb25zdCBzYXZlZFByZWZlcmVuY2VzID0gdGhpcy4jbG9hZEV2ZW50RmlsdGVyc0Zyb21TdG9yYWdlKCk7XG5cbiAgICBjb25zdCBldmVudFR5cGVGaWx0ZXIgPSB0aGlzLiMkKCdldmVudC10eXBlLWZpbHRlcicpO1xuICAgIGlmIChldmVudFR5cGVGaWx0ZXIgJiYgdGhpcy4jZWxlbWVudEV2ZW50TWFwKSB7XG4gICAgICBsZXQgbWVudSA9IGV2ZW50VHlwZUZpbHRlci5xdWVyeVNlbGVjdG9yKCdwZi12Ni1tZW51Jyk7XG4gICAgICBpZiAoIW1lbnUpIHtcbiAgICAgICAgbWVudSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3BmLXY2LW1lbnUnKTtcbiAgICAgICAgZXZlbnRUeXBlRmlsdGVyLmFwcGVuZENoaWxkKG1lbnUpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBleGlzdGluZ0l0ZW1zID0gbWVudS5xdWVyeVNlbGVjdG9yQWxsKCdwZi12Ni1tZW51LWl0ZW0nKTtcbiAgICAgIGV4aXN0aW5nSXRlbXMuZm9yRWFjaChpdGVtID0+IGl0ZW0ucmVtb3ZlKCkpO1xuXG4gICAgICBjb25zdCBhbGxFdmVudFR5cGVzID0gbmV3IFNldDxzdHJpbmc+KCk7XG4gICAgICBmb3IgKGNvbnN0IFt0YWdOYW1lLCBldmVudEluZm9dIG9mIHRoaXMuI2VsZW1lbnRFdmVudE1hcCkge1xuICAgICAgICBpZiAodGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmhhcyh0YWdOYW1lKSkge1xuICAgICAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGV2ZW50SW5mby5ldmVudE5hbWVzKSB7XG4gICAgICAgICAgICBhbGxFdmVudFR5cGVzLmFkZChldmVudE5hbWUpO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICBpZiAoc2F2ZWRQcmVmZXJlbmNlcy5ldmVudFR5cGVzKSB7XG4gICAgICAgIHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMgPSAoc2F2ZWRQcmVmZXJlbmNlcy5ldmVudFR5cGVzIGFzIFNldDxzdHJpbmc+ICYgeyBpbnRlcnNlY3Rpb246IChvdGhlcjogU2V0PHN0cmluZz4pID0+IFNldDxzdHJpbmc+IH0pLmludGVyc2VjdGlvbihhbGxFdmVudFR5cGVzKTtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMgPSBuZXcgU2V0KGFsbEV2ZW50VHlwZXMpO1xuICAgICAgfVxuXG4gICAgICBmb3IgKGNvbnN0IGV2ZW50TmFtZSBvZiBhbGxFdmVudFR5cGVzKSB7XG4gICAgICAgIGNvbnN0IGl0ZW0gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwZi12Ni1tZW51LWl0ZW0nKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3ZhcmlhbnQnLCAnY2hlY2tib3gnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3ZhbHVlJywgZXZlbnROYW1lKTtcbiAgICAgICAgaWYgKHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuaGFzKGV2ZW50TmFtZSkpIHtcbiAgICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnY2hlY2tlZCcsICcnKTtcbiAgICAgICAgfVxuICAgICAgICBpdGVtLnRleHRDb250ZW50ID0gZXZlbnROYW1lO1xuICAgICAgICBtZW51LmFwcGVuZENoaWxkKGl0ZW0pO1xuICAgICAgfVxuICAgIH1cblxuICAgIGNvbnN0IGVsZW1lbnRGaWx0ZXIgPSB0aGlzLiMkKCdlbGVtZW50LWZpbHRlcicpO1xuICAgIGlmIChlbGVtZW50RmlsdGVyICYmIHRoaXMuI2VsZW1lbnRFdmVudE1hcCkge1xuICAgICAgbGV0IG1lbnUgPSBlbGVtZW50RmlsdGVyLnF1ZXJ5U2VsZWN0b3IoJ3BmLXY2LW1lbnUnKTtcbiAgICAgIGlmICghbWVudSkge1xuICAgICAgICBtZW51ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncGYtdjYtbWVudScpO1xuICAgICAgICBlbGVtZW50RmlsdGVyLmFwcGVuZENoaWxkKG1lbnUpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBleGlzdGluZ0l0ZW1zID0gbWVudS5xdWVyeVNlbGVjdG9yQWxsKCdwZi12Ni1tZW51LWl0ZW0nKTtcbiAgICAgIGV4aXN0aW5nSXRlbXMuZm9yRWFjaChpdGVtID0+IGl0ZW0ucmVtb3ZlKCkpO1xuXG4gICAgICBjb25zdCBhbGxFbGVtZW50cyA9IG5ldyBTZXQ8c3RyaW5nPigpO1xuICAgICAgZm9yIChjb25zdCB0YWdOYW1lIG9mIHRoaXMuI2VsZW1lbnRFdmVudE1hcC5rZXlzKCkpIHtcbiAgICAgICAgaWYgKHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5oYXModGFnTmFtZSkpIHtcbiAgICAgICAgICBhbGxFbGVtZW50cy5hZGQodGFnTmFtZSk7XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgaWYgKHNhdmVkUHJlZmVyZW5jZXMuZWxlbWVudHMpIHtcbiAgICAgICAgdGhpcy4jZWxlbWVudEZpbHRlcnMgPSAoc2F2ZWRQcmVmZXJlbmNlcy5lbGVtZW50cyBhcyBTZXQ8c3RyaW5nPiAmIHsgaW50ZXJzZWN0aW9uOiAob3RoZXI6IFNldDxzdHJpbmc+KSA9PiBTZXQ8c3RyaW5nPiB9KS5pbnRlcnNlY3Rpb24oYWxsRWxlbWVudHMpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgdGhpcy4jZWxlbWVudEZpbHRlcnMgPSBuZXcgU2V0KGFsbEVsZW1lbnRzKTtcbiAgICAgIH1cblxuICAgICAgZm9yIChjb25zdCB0YWdOYW1lIG9mIGFsbEVsZW1lbnRzKSB7XG4gICAgICAgIGNvbnN0IGl0ZW0gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwZi12Ni1tZW51LWl0ZW0nKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3ZhcmlhbnQnLCAnY2hlY2tib3gnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3ZhbHVlJywgdGFnTmFtZSk7XG4gICAgICAgIGlmICh0aGlzLiNlbGVtZW50RmlsdGVycy5oYXModGFnTmFtZSkpIHtcbiAgICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnY2hlY2tlZCcsICcnKTtcbiAgICAgICAgfVxuICAgICAgICBpdGVtLnRleHRDb250ZW50ID0gYDwke3RhZ05hbWV9PmA7XG4gICAgICAgIG1lbnUuYXBwZW5kQ2hpbGQoaXRlbSk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI2hhbmRsZUV2ZW50VHlwZUZpbHRlckNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB7IHZhbHVlLCBjaGVja2VkIH0gPSBldmVudCBhcyBFdmVudCAmIHsgdmFsdWU6IHN0cmluZzsgY2hlY2tlZDogYm9vbGVhbiB9O1xuXG4gICAgaWYgKCF2YWx1ZSkgcmV0dXJuO1xuXG4gICAgaWYgKGNoZWNrZWQpIHtcbiAgICAgIHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuYWRkKHZhbHVlKTtcbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5kZWxldGUodmFsdWUpO1xuICAgIH1cblxuICAgIHRoaXMuI3NhdmVFdmVudEZpbHRlcnMoKTtcbiAgICB0aGlzLiNmaWx0ZXJFdmVudHModGhpcy4jZXZlbnRzRmlsdGVyVmFsdWUpO1xuICB9O1xuXG4gICNoYW5kbGVFbGVtZW50RmlsdGVyQ2hhbmdlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHsgdmFsdWUsIGNoZWNrZWQgfSA9IGV2ZW50IGFzIEV2ZW50ICYgeyB2YWx1ZTogc3RyaW5nOyBjaGVja2VkOiBib29sZWFuIH07XG5cbiAgICBpZiAoIXZhbHVlKSByZXR1cm47XG5cbiAgICBpZiAoY2hlY2tlZCkge1xuICAgICAgdGhpcy4jZWxlbWVudEZpbHRlcnMuYWRkKHZhbHVlKTtcbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy4jZWxlbWVudEZpbHRlcnMuZGVsZXRlKHZhbHVlKTtcbiAgICB9XG5cbiAgICB0aGlzLiNzYXZlRXZlbnRGaWx0ZXJzKCk7XG4gICAgdGhpcy4jZmlsdGVyRXZlbnRzKHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKTtcbiAgfTtcblxuICAjbG9hZEV2ZW50RmlsdGVyc0Zyb21TdG9yYWdlKCk6IHsgZXZlbnRUeXBlczogU2V0PHN0cmluZz4gfCBudWxsOyBlbGVtZW50czogU2V0PHN0cmluZz4gfCBudWxsIH0ge1xuICAgIGNvbnN0IHByZWZlcmVuY2VzOiB7IGV2ZW50VHlwZXM6IFNldDxzdHJpbmc+IHwgbnVsbDsgZWxlbWVudHM6IFNldDxzdHJpbmc+IHwgbnVsbCB9ID0ge1xuICAgICAgZXZlbnRUeXBlczogbnVsbCxcbiAgICAgIGVsZW1lbnRzOiBudWxsXG4gICAgfTtcblxuICAgIHRyeSB7XG4gICAgICBjb25zdCBzYXZlZEV2ZW50VHlwZXMgPSBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWV2ZW50LXR5cGUtZmlsdGVycycpO1xuICAgICAgaWYgKHNhdmVkRXZlbnRUeXBlcykge1xuICAgICAgICBwcmVmZXJlbmNlcy5ldmVudFR5cGVzID0gbmV3IFNldChKU09OLnBhcnNlKHNhdmVkRXZlbnRUeXBlcykpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBzYXZlZEVsZW1lbnRzID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1lbGVtZW50LWZpbHRlcnMnKTtcbiAgICAgIGlmIChzYXZlZEVsZW1lbnRzKSB7XG4gICAgICAgIHByZWZlcmVuY2VzLmVsZW1lbnRzID0gbmV3IFNldChKU09OLnBhcnNlKHNhdmVkRWxlbWVudHMpKTtcbiAgICAgIH1cbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICBjb25zb2xlLmRlYnVnKCdbY2VtLXNlcnZlLWNocm9tZV0gbG9jYWxTdG9yYWdlIHVuYXZhaWxhYmxlIGZvciBldmVudCBmaWx0ZXJzJyk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHByZWZlcmVuY2VzO1xuICB9XG5cbiAgI3NhdmVFdmVudEZpbHRlcnMoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdjZW0tc2VydmUtZXZlbnQtdHlwZS1maWx0ZXJzJyxcbiAgICAgICAgSlNPTi5zdHJpbmdpZnkoWy4uLnRoaXMuI2V2ZW50VHlwZUZpbHRlcnNdKSk7XG4gICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLWVsZW1lbnQtZmlsdGVycycsXG4gICAgICAgIEpTT04uc3RyaW5naWZ5KFsuLi50aGlzLiNlbGVtZW50RmlsdGVyc10pKTtcbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICAvLyBsb2NhbFN0b3JhZ2UgdW5hdmFpbGFibGUgKHByaXZhdGUgbW9kZSksIHNpbGVudGx5IGNvbnRpbnVlXG4gICAgfVxuICB9XG5cbiAgI2NsZWFyRXZlbnRzKCkge1xuICAgIHRoaXMuI2NhcHR1cmVkRXZlbnRzID0gW107XG4gICAgdGhpcy4jc2VsZWN0ZWRFdmVudElkID0gbnVsbDtcbiAgICBpZiAodGhpcy4jZXZlbnRMaXN0KSB7XG4gICAgICB0aGlzLiNldmVudExpc3QucmVwbGFjZUNoaWxkcmVuKCk7XG4gICAgfVxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEhlYWRlcikge1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIuaW5uZXJIVE1MID0gJyc7XG4gICAgfVxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEJvZHkpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5pbm5lckhUTUwgPSAnJztcbiAgICB9XG4gIH1cblxuICBhc3luYyAjY29weUV2ZW50cygpIHtcbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgY29uc3QgdmlzaWJsZUV2ZW50cyA9IEFycmF5LmZyb20odGhpcy4jZXZlbnRMaXN0LmNoaWxkcmVuKVxuICAgICAgLmZpbHRlcihlbnRyeSA9PiAhKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5oaWRkZW4pXG4gICAgICAubWFwKGVudHJ5ID0+IHtcbiAgICAgICAgY29uc3QgaWQgPSAoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuZXZlbnRJZCE7XG4gICAgICAgIHJldHVybiB0aGlzLiNnZXRFdmVudFJlY29yZEJ5SWQoaWQpO1xuICAgICAgfSlcbiAgICAgIC5maWx0ZXIoKGV2ZW50KTogZXZlbnQgaXMgRXZlbnRSZWNvcmQgPT4gISFldmVudClcbiAgICAgIC5tYXAoZXZlbnQgPT4ge1xuICAgICAgICBjb25zdCB0aW1lID0gZXZlbnQudGltZXN0YW1wLnRvTG9jYWxlVGltZVN0cmluZygpO1xuICAgICAgICBjb25zdCBlbGVtZW50ID0gZXZlbnQuZWxlbWVudElkID8gYCMke2V2ZW50LmVsZW1lbnRJZH1gIDogZXZlbnQudGFnTmFtZTtcbiAgICAgICAgY29uc3QgcHJvcHMgPSBldmVudC5jdXN0b21Qcm9wZXJ0aWVzICYmIE9iamVjdC5rZXlzKGV2ZW50LmN1c3RvbVByb3BlcnRpZXMpLmxlbmd0aCA+IDBcbiAgICAgICAgICA/IGAgLSAke0pTT04uc3RyaW5naWZ5KGV2ZW50LmN1c3RvbVByb3BlcnRpZXMpfWBcbiAgICAgICAgICA6ICcnO1xuICAgICAgICByZXR1cm4gYFske3RpbWV9XSA8JHtldmVudC50YWdOYW1lfT4gJHtlbGVtZW50fSBcXHUyMTkyICR7ZXZlbnQuZXZlbnROYW1lfSR7cHJvcHN9YDtcbiAgICAgIH0pXG4gICAgICAuam9pbignXFxuJyk7XG5cbiAgICBpZiAoIXZpc2libGVFdmVudHMpIHJldHVybjtcblxuICAgIHRyeSB7XG4gICAgICBhd2FpdCBuYXZpZ2F0b3IuY2xpcGJvYXJkLndyaXRlVGV4dCh2aXNpYmxlRXZlbnRzKTtcbiAgICAgIGNvbnN0IGJ0biA9IHRoaXMuIyQoJ2NvcHktZXZlbnRzJyk7XG4gICAgICBpZiAoYnRuKSB7XG4gICAgICAgIGNvbnN0IHRleHROb2RlID0gQXJyYXkuZnJvbShidG4uY2hpbGROb2RlcykuZmluZChcbiAgICAgICAgICBuID0+IG4ubm9kZVR5cGUgPT09IE5vZGUuVEVYVF9OT0RFICYmIChuLnRleHRDb250ZW50Py50cmltKCkubGVuZ3RoID8/IDApID4gMFxuICAgICAgICApO1xuICAgICAgICBpZiAodGV4dE5vZGUpIHtcbiAgICAgICAgICBjb25zdCBvcmlnaW5hbCA9IHRleHROb2RlLnRleHRDb250ZW50O1xuICAgICAgICAgIHRleHROb2RlLnRleHRDb250ZW50ID0gJ0NvcGllZCEnO1xuXG4gICAgICAgICAgaWYgKHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgICAgICB9XG5cbiAgICAgICAgICB0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgICBpZiAodGhpcy5pc0Nvbm5lY3RlZCAmJiB0ZXh0Tm9kZS5wYXJlbnROb2RlKSB7XG4gICAgICAgICAgICAgIHRleHROb2RlLnRleHRDb250ZW50ID0gb3JpZ2luYWw7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgICB0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICAgICAgICB9LCAyMDAwKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gY2F0Y2ggKGVycikge1xuICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjb3B5IGV2ZW50czonLCBlcnIpO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cEV2ZW50TGlzdGVuZXJzKCkge1xuICAgIHRoaXMuI2V2ZW50TGlzdCA9IHRoaXMuIyQoJ2V2ZW50LWxpc3QnKTtcbiAgICB0aGlzLiNldmVudERldGFpbEhlYWRlciA9IHRoaXMuIyQoJ2V2ZW50LWRldGFpbC1oZWFkZXInKTtcbiAgICB0aGlzLiNldmVudERldGFpbEJvZHkgPSB0aGlzLiMkKCdldmVudC1kZXRhaWwtYm9keScpO1xuXG4gICAgaWYgKHRoaXMuI2V2ZW50TGlzdCkge1xuICAgICAgdGhpcy4jZXZlbnRMaXN0LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICAgIGNvbnN0IGxpc3RJdGVtID0gKGUudGFyZ2V0IGFzIEVsZW1lbnQpLmNsb3Nlc3QoJy5ldmVudC1saXN0LWl0ZW0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICAgICAgaWYgKGxpc3RJdGVtKSB7XG4gICAgICAgICAgY29uc3QgZXZlbnRJZCA9IGxpc3RJdGVtLmRhdGFzZXQuZXZlbnRJZDtcbiAgICAgICAgICBpZiAoZXZlbnRJZCkge1xuICAgICAgICAgICAgdGhpcy4jc2VsZWN0RXZlbnQoZXZlbnRJZCk7XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9KTtcbiAgICB9XG5cbiAgICBjb25zdCBldmVudHNGaWx0ZXIgPSB0aGlzLiMkKCdldmVudHMtZmlsdGVyJyk7XG4gICAgaWYgKGV2ZW50c0ZpbHRlcikge1xuICAgICAgZXZlbnRzRmlsdGVyLmFkZEV2ZW50TGlzdGVuZXIoJ2lucHV0JywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICAgIGNvbnN0IHsgdmFsdWUgPSAnJyB9ID0gZS50YXJnZXQgYXMgSFRNTElucHV0RWxlbWVudDtcbiAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2V2ZW50c0ZpbHRlckRlYm91bmNlVGltZXIhKTtcbiAgICAgICAgdGhpcy4jZXZlbnRzRmlsdGVyRGVib3VuY2VUaW1lciA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgIHRoaXMuI2ZpbHRlckV2ZW50cyh2YWx1ZSk7XG4gICAgICAgIH0sIDMwMCk7XG4gICAgICB9KTtcbiAgICB9XG5cbiAgICBjb25zdCBldmVudFR5cGVGaWx0ZXIgPSB0aGlzLiMkKCdldmVudC10eXBlLWZpbHRlcicpO1xuICAgIGlmIChldmVudFR5cGVGaWx0ZXIpIHtcbiAgICAgIGV2ZW50VHlwZUZpbHRlci5hZGRFdmVudExpc3RlbmVyKCdzZWxlY3QnLCB0aGlzLiNoYW5kbGVFdmVudFR5cGVGaWx0ZXJDaGFuZ2UgYXMgRXZlbnRMaXN0ZW5lcik7XG4gICAgfVxuXG4gICAgY29uc3QgZWxlbWVudEZpbHRlciA9IHRoaXMuIyQoJ2VsZW1lbnQtZmlsdGVyJyk7XG4gICAgaWYgKGVsZW1lbnRGaWx0ZXIpIHtcbiAgICAgIGVsZW1lbnRGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlRWxlbWVudEZpbHRlckNoYW5nZSBhcyBFdmVudExpc3RlbmVyKTtcbiAgICB9XG5cbiAgICB0aGlzLiMkKCdjbGVhci1ldmVudHMnKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB0aGlzLiNjbGVhckV2ZW50cygpO1xuICAgIH0pO1xuXG4gICAgdGhpcy4jJCgnY29weS1ldmVudHMnKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB0aGlzLiNjb3B5RXZlbnRzKCk7XG4gICAgfSk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuXG4gICAgLy8gQ2xlYW4gdXAga25vYiBsaXN0ZW5lcnNcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOnByb3BlcnR5LWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOmNzcy1wcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjphdHRyaWJ1dGUtY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOnByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG5cbiAgICAvLyBDbGVhbiB1cCB0cmVlIHN0YXRlIGxpc3RlbmVyc1xuICAgIGlmICh0aGlzLiNoYW5kbGVUcmVlRXhwYW5kKSB7XG4gICAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2V4cGFuZCcsIHRoaXMuI2hhbmRsZVRyZWVFeHBhbmQpO1xuICAgIH1cbiAgICBpZiAodGhpcy4jaGFuZGxlVHJlZUNvbGxhcHNlKSB7XG4gICAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2NvbGxhcHNlJywgdGhpcy4jaGFuZGxlVHJlZUNvbGxhcHNlKTtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QpIHtcbiAgICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlVHJlZVNlbGVjdCk7XG4gICAgfVxuXG4gICAgLy8gQ2xlYW4gdXAgd2luZG93IGxpc3RlbmVyXG4gICAgaWYgKHRoaXMuI2hhbmRsZUxvZ3NFdmVudCkge1xuICAgICAgd2luZG93LnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2NlbTpsb2dzJywgdGhpcy4jaGFuZGxlTG9nc0V2ZW50KTtcbiAgICB9XG5cbiAgICAvLyBDbGVhciBwZW5kaW5nIGZlZWRiYWNrIHRpbWVvdXRzXG4gICAgaWYgKHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgdGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgIH1cbiAgICBpZiAodGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgIHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgfVxuICAgIGlmICh0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICB0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICB9XG5cbiAgICAvLyBEaXNjb25uZWN0IG11dGF0aW9uIG9ic2VydmVyXG4gICAgdGhpcy4jb2JzZXJ2ZXIuZGlzY29ubmVjdCgpO1xuXG4gICAgLy8gQ2xvc2UgV2ViU29ja2V0IGNvbm5lY3Rpb25cbiAgICBpZiAodGhpcy4jd3NDbGllbnQpIHtcbiAgICAgIHRoaXMuI3dzQ2xpZW50LmRlc3Ryb3koKTtcbiAgICB9XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXNlcnZlLWNocm9tZSc6IENlbVNlcnZlQ2hyb21lO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBoZWlnaHQ6IDEwMHZoO1xcXFxuICBvdmVyZmxvdzogaGlkZGVuO1xcXFxuICAtLXBmLXY2LWMtbWFzdGhlYWRfX2xvZ28tLVdpZHRoOiBtYXgtY29udGVudDtcXFxcbiAgLS1wZi12Ni1jLW1hc3RoZWFkX190b2dnbGUtLVNpemU6IDFyZW07XFxcXG59XFxcXG5cXFxcbltoaWRkZW5dIHtcXFxcbiAgZGlzcGxheTogbm9uZSAhaW1wb3J0YW50O1xcXFxufVxcXFxuXFxcXG4vKiBNYXN0aGVhZCBsb2dvIHN0eWxlcyAqL1xcXFxuLm1hc3RoZWFkLWxvZ28ge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG4gIGNvbG9yOiBpbmhlcml0O1xcXFxuICBtYXgtaGVpZ2h0OiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkX19sb2dvLS1NYXhIZWlnaHQpO1xcXFxuICBnYXA6IDRweDtcXFxcbiAgXFxcXHUwMDI2IGltZyB7XFxcXG4gICAgZGlzcGxheTogYmxvY2s7XFxcXG4gICAgbWF4LWhlaWdodDogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZF9fbG9nby0tTWF4SGVpZ2h0KTtcXFxcbiAgICB3aWR0aDogYXV0bztcXFxcbiAgfVxcXFxuICBcXFxcdTAwMjYgOjpzbG90dGVkKFtzbG90PVxcXFxcXFwidGl0bGVcXFxcXFxcIl0pIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gICAgZm9udC1zaXplOiAxLjEyNXJlbTtcXFxcbiAgICBmb250LXdlaWdodDogNjAwO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgfVxcXFxuICBcXFxcdTAwMjYgaDEge1xcXFxuICAgIG1hcmdpbjogMDtcXFxcbiAgICBmb250LXNpemU6IDE4cHg7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogVG9vbGJhciBncm91cCBhbGlnbm1lbnQgKi9cXFxcbnBmLXY2LXRvb2xiYXItZ3JvdXBbdmFyaWFudD1cXFxcXFxcImFjdGlvbi1ncm91cFxcXFxcXFwiXSB7XFxcXG4gIG1hcmdpbi1pbmxpbmUtc3RhcnQ6IGF1dG87XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG59XFxcXG5cXFxcbi5kZWJ1Zy1wYW5lbCB7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWRlZmF1bHQpO1xcXFxuICBib3JkZXI6IDFweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxuICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gIHBhZGRpbmc6IDEuNXJlbTtcXFxcbiAgbWF4LXdpZHRoOiA2MDBweDtcXFxcbiAgd2lkdGg6IDkwJTtcXFxcbiAgbWF4LWhlaWdodDogODB2aDtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcblxcXFxuICBoMiB7XFxcXG4gICAgbWFyZ2luOiAwIDAgMXJlbSAwO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICBmb250LXNpemU6IDEuMTI1cmVtO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIH1cXFxcblxcXFxuICBkbCB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICB9XFxcXG5cXFxcbiAgZHQge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgbWFyZ2luLXRvcDogMC41cmVtO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG4gIH1cXFxcblxcXFxuICBkZCB7XFxcXG4gICAgbWFyZ2luOiAwIDAgMC41cmVtIDA7XFxcXG4gICAgZm9udC1mYW1pbHk6IHVpLW1vbm9zcGFjZSwgJ0Nhc2NhZGlhIENvZGUnLCAnU291cmNlIENvZGUgUHJvJywgTWVubG8sIENvbnNvbGFzLCAnRGVqYVZ1IFNhbnMgTW9ubycsIG1vbm9zcGFjZTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgfVxcXFxuXFxcXG4gIGRldGFpbHMge1xcXFxuICAgIG1hcmdpbi10b3A6IDFyZW07XFxcXG5cXFxcbiAgICBzdW1tYXJ5IHtcXFxcbiAgICAgIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgICBmb250LXdlaWdodDogNTAwO1xcXFxuICAgICAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICAgICAgZ2FwOiAwLjVyZW07XFxcXG4gICAgICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcblxcXFxuICAgICAgXFxcXHUwMDI2Ojotd2Via2l0LWRldGFpbHMtbWFya2VyIHtcXFxcbiAgICAgICAgZGlzcGxheTogbm9uZTtcXFxcbiAgICAgIH1cXFxcblxcXFxuICAgICAgXFxcXHUwMDI2OjpiZWZvcmUge1xcXFxuICAgICAgICBjb250ZW50OiAnXFxcXFxcXFwyNUI4JztcXFxcbiAgICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICAgICAgICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gMTAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgICB9XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2W29wZW5dIHN1bW1hcnk6OmJlZm9yZSB7XFxcXG4gICAgICB0cmFuc2Zvcm06IHJvdGF0ZSg5MGRlZyk7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgcHJlIHtcXFxcbiAgICAgIG1hcmdpbi10b3A6IDAuNXJlbTtcXFxcbiAgICAgIG1hcmdpbi1sZWZ0OiAxLjVyZW07XFxcXG4gICAgICBwYWRkaW5nOiAwLjVyZW07XFxcXG4gICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWRlZmF1bHQpO1xcXFxuICAgICAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICAgICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICAgIG92ZXJmbG93LXg6IGF1dG87XFxcXG4gICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgLmJ1dHRvbi1jb250YWluZXIge1xcXFxuICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgZ2FwOiAwLjVyZW07XFxcXG4gICAgbWFyZ2luLXRvcDogMXJlbTtcXFxcbiAgfVxcXFxuXFxcXG4gIHAge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gIH1cXFxcblxcXFxuICBidXR0b24ge1xcXFxuICAgIG1hcmdpbi10b3A6IDFyZW07XFxcXG4gICAgcGFkZGluZzogMC41cmVtIDFyZW07XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQpO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1icmFuZCk7XFxcXG4gICAgYm9yZGVyOiBub25lO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA0MDA7XFxcXG4gICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICAgIHRyYW5zaXRpb246IGFsbCAyMDBtcyBjdWJpYy1iZXppZXIoMC42NDUsIDAuMDQ1LCAwLjM1NSwgMSk7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWhvdmVyKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogQ29udGVudCBhcmVhIHBhZGRpbmcgZm9yIGRlbW8gKi9cXFxcbnBmLXY2LXBhZ2UtbWFpbiB7XFxcXG4gIG1pbi1oZWlnaHQ6IGNhbGMoMTAwZHZoIC0gNzJweCAtIHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1pbnNldC0tcGFnZS1jaHJvbWUpKTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgXFxcXHUwMDI2IFxcXFx1MDAzZSA6OnNsb3R0ZWQoOm5vdChbc2xvdD1rbm9ic10pKSB7XFxcXG4gICAgcGFkZGluZzogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWxnKTtcXFxcbiAgICBmbGV4OiAxO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbmNlbS1kcmF3ZXIge1xcXFxuICBwZi12Ni10YWJzIHtcXFxcbiAgICBwZi12Ni10YWIge1xcXFxuICAgICAgcGFkZGluZy1ibG9jay1lbmQ6IDA7XFxcXG4gICAgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIEVsZW1lbnQgZGVzY3JpcHRpb25zIGluIGxpc3RpbmcgKi9cXFxcbi5lbGVtZW50LXN1bW1hcnkge1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxufVxcXFxuXFxcXG4uZWxlbWVudC1kZXNjcmlwdGlvbiB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG59XFxcXG5cXFxcbi8qIENhcmQgZm9vdGVyIGRlbW8gbmF2aWdhdGlvbiAqL1xcXFxuLmNhcmQtZGVtb3Mge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LXdyYXA6IHdyYXA7XFxcXG4gIGdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWdhcC0tYWN0aW9uLXRvLWFjdGlvbi0tZGVmYXVsdCk7XFxcXG4gIHBhZGRpbmc6IDA7XFxcXG4gIG1hcmdpbjogMDtcXFxcbn1cXFxcblxcXFxuLnBhY2thZ2UtbmFtZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxufVxcXFxuXFxcXG4vKiBLbm9icyBjb250YWluZXIgLSBmaWxscyB0YWIgcGFuZWwgaGVpZ2h0ICovXFxcXG4ja25vYnMtY29udGFpbmVyIHtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBcXFxcdTAwMjYgOjpzbG90dGVkKFtzbG90PVxcXFxcXFwia25vYnNcXFxcXFxcIl0pIHtcXFxcbiAgICBmbGV4OiAxO1xcXFxuICAgIG1pbi1oZWlnaHQ6IDA7XFxcXG4gICAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ua25vYnMtZW1wdHkge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgdGV4dC1hbGlnbjogY2VudGVyO1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLWxnKTtcXFxcblxcXFxuICBjb2RlIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gICAgcGFkZGluZzogMnB4IDZweDtcXFxcbiAgICBib3JkZXItcmFkaXVzOiAzcHg7XFxcXG4gICAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5pbnN0YW5jZS10YWcge1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5pbnN0YW5jZS1sYWJlbCB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLmtub2ItZ3JvdXAge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLWxnKTtcXFxcblxcXFxuICBcXFxcdTAwMjY6bGFzdC1jaGlsZCB7XFxcXG4gICAgbWFyZ2luLWJvdHRvbTogMDtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBQYXR0ZXJuRmx5IHY2IGZvcm0gLSBob3Jpem9udGFsIGxheW91dCAqL1xcXFxucGYtdjYtZm9ybVtob3Jpem9udGFsXSBwZi12Ni1mb3JtLWZpZWxkLWdyb3VwIHtcXFxcbiAgZ3JpZC1jb2x1bW46IHNwYW4gMlxcXFxufVxcXFxuXFxcXG4ua25vYi1ncm91cC10aXRsZSB7XFxcXG4gIGdyaWQtY29sdW1uOiAxIC8gLTE7XFxcXG4gIG1hcmdpbjogMCAwIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpIDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1iYXNlKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgYm9yZGVyLWJvdHRvbTogMXB4IHNvbGlkIHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1jb2xvcik7XFxcXG4gIHBhZGRpbmctYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbn1cXFxcblxcXFxuLmtub2ItY29udHJvbCB7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxufVxcXFxuXFxcXG4ua25vYi1sYWJlbCB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cyk7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbn1cXFxcblxcXFxuLmtub2ItbmFtZSB7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgZm9udC13ZWlnaHQ6IDUwMDtcXFxcbn1cXFxcblxcXFxuLmtub2ItZGVzY3JpcHRpb24ge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjU7XFxcXG5cXFxcbiAgcCB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSAwO1xcXFxuICB9XFxcXG5cXFxcbiAgY29kZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDNweDtcXFxcbiAgICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIH1cXFxcblxcXFxuICBhIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgdGV4dC1kZWNvcmF0aW9uOiB1bmRlcmxpbmU7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgc3Ryb25nIHtcXFxcbiAgICBmb250LXdlaWdodDogNjAwO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxuICB9XFxcXG5cXFxcbiAgdWwsIG9sIHtcXFxcbiAgICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIDA7XFxcXG4gICAgcGFkZGluZy1sZWZ0OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLWxnKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG5mb290ZXIucGYtbS1zdGlja3ktYm90dG9tIHtcXFxcbiAgdmlldy10cmFuc2l0aW9uLW5hbWU6IGRldi1zZXJ2ZXItZm9vdGVyO1xcXFxuICBwb3NpdGlvbjogc3RpY2t5O1xcXFxuICBib3R0b206IDA7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWRlZmF1bHQpO1xcXFxuICBib3JkZXItdG9wOiAxcHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbiAgei1pbmRleDogdmFyKC0tcGYtdjYtYy1wYWdlLS1zZWN0aW9uLS1tLXN0aWNreS1ib3R0b20tLVpJbmRleCwgdmFyKC0tcGYtdC0tZ2xvYmFsLS16LWluZGV4LS1tZCkpO1xcXFxuICBib3gtc2hhZG93OiB2YXIoLS1wZi12Ni1jLXBhZ2UtLXNlY3Rpb24tLW0tc3RpY2t5LWJvdHRvbS0tQm94U2hhZG93LCB2YXIoLS1wZi10LS1nbG9iYWwtLWJveC1zaGFkb3ctLXNtLS10b3ApKTtcXFxcbn1cXFxcblxcXFxuLmZvb3Rlci1kZXNjcmlwdGlvbiB7XFxcXG4gIHBhZGRpbmc6IDEuNXJlbTtcXFxcblxcXFxuICBcXFxcdTAwMjYuZW1wdHkge1xcXFxuICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuZm9vdGVyIDo6c2xvdHRlZChbc2xvdD1cXFxcXFxcImRlc2NyaXB0aW9uXFxcXFxcXCJdKSB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjY7XFxcXG4gIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG5cXFxcbiAgY29kZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0taG92ZXIpO1xcXFxuICAgIHBhZGRpbmc6IDJweCA2cHg7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogM3B4O1xcXFxuICAgIGZvbnQtZmFtaWx5OiB1aS1tb25vc3BhY2UsICdDYXNjYWRpYSBDb2RlJywgJ1NvdXJjZSBDb2RlIFBybycsIE1lbmxvLCBDb25zb2xhcywgJ0RlamFWdSBTYW5zIE1vbm8nLCBtb25vc3BhY2U7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubG9ncy13cmFwcGVyIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4jbG9nLWNvbnRhaW5lciB7XFxcXG4gIGZsZXgtZ3JvdzogMTtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbn1cXFxcblxcXFxuLmxvZy1lbnRyeSB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxuICBhbGlnbi1pdGVtczogYmFzZWxpbmU7XFxcXG4gIHBmLXY2LWxhYmVsIHtcXFxcbiAgICBmbGV4LXNocmluazogMDtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubG9nLXRpbWUsXFxcXG4ubG9nLW1lc3NhZ2Uge1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLmxvZy10aW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmbGV4LXNocmluazogMDtcXFxcbiAgZm9udC1zaXplOiAxMXB4O1xcXFxufVxcXFxuXFxcXG4ubG9nLW1lc3NhZ2Uge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgd29yZC1icmVhazogYnJlYWstd29yZDtcXFxcbn1cXFxcblxcXFxuLyogTmF2aWdhdGlvbiBjb250ZW50IChsaWdodCBET00gc2xvdHRlZCBjb250ZW50IGZvciBwZi12Ni1wYWdlLXNpZGViYXIpICovXFxcXG4ubmF2LXBhY2thZ2Uge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcblxcXFxuICBcXFxcdTAwMjYgXFxcXHUwMDNlIHN1bW1hcnkge1xcXFxuICAgIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgICBwYWRkaW5nOiAwLjVyZW0gMXJlbTtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWRlZmF1bHQpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gICAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAyMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgIG1hcmdpbi1ib3R0b206IDAuNXJlbTtcXFxcbiAgICBkaXNwbGF5OiBmbGV4O1xcXFxuICAgIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gICAgZ2FwOiAwLjVyZW07XFxcXG4gICAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1ob3Zlcik7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2Ojotd2Via2l0LWRldGFpbHMtbWFya2VyIHtcXFxcbiAgICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2OjpiZWZvcmUge1xcXFxuICAgICAgY29udGVudDogJ1xcXFxcXFxcMjVCOCc7XFxcXG4gICAgICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XFxcXG4gICAgICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gMTAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNltvcGVuXSBcXFxcdTAwM2Ugc3VtbWFyeTo6YmVmb3JlIHtcXFxcbiAgICB0cmFuc2Zvcm06IHJvdGF0ZSg5MGRlZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLm5hdi1lbGVtZW50IHtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG4gIG1hcmdpbi1pbmxpbmUtc3RhcnQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuXFxcXG4gIHN1bW1hcnkge1xcXFxuICAgIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgICBwYWRkaW5nOiAwLjVyZW0gMXJlbTtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWRlZmF1bHQpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgZm9udC1mYW1pbHk6IHVpLW1vbm9zcGFjZSwgJ0Nhc2NhZGlhIENvZGUnLCAnU291cmNlIENvZGUgUHJvJywgTWVubG8sIENvbnNvbGFzLCAnRGVqYVZ1IFNhbnMgTW9ubycsIG1vbm9zcGFjZTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gICAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAyMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgICBnYXA6IDAuNXJlbTtcXFxcbiAgICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWhvdmVyKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6Oi13ZWJraXQtZGV0YWlscy1tYXJrZXIge1xcXFxuICAgICAgZGlzcGxheTogbm9uZTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgICBjb250ZW50OiAnXFxcXFxcXFwyNUI4JztcXFxcbiAgICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgICAgIHRyYW5zaXRpb246IHRyYW5zZm9ybSAxMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2W29wZW5dIHN1bW1hcnk6OmJlZm9yZSB7XFxcXG4gICAgdHJhbnNmb3JtOiByb3RhdGUoOTBkZWcpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5uYXYtZWxlbWVudC10aXRsZSB7XFxcXG4gIHVzZXItc2VsZWN0OiBub25lO1xcXFxufVxcXFxuXFxcXG4ubmF2LWRlbW8tbGlzdCB7XFxcXG4gIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gIHBhZGRpbmc6IDA7XFxcXG4gIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMCAwIDA7XFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cyk7XFxcXG59XFxcXG5cXFxcbi5uYXYtZGVtby1saW5rIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIHBhZGRpbmctaW5saW5lLXN0YXJ0OiBjYWxjKHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpICogMik7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXJhZGl1cyk7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIHRyYW5zaXRpb246IGJhY2tncm91bmQgMC4ycywgY29sb3IgMC4ycztcXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLWJyYW5kKTtcXFxcblxcXFxuICAgIC5uYXYtcGFja2FnZS1uYW1lIHtcXFxcbiAgICAgIGNvbG9yOiByZ2JhKDI1NSwgMjU1LCAyNTUsIDAuOCk7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2W2FyaWEtY3VycmVudD1cXFxcXFxcInBhZ2VcXFxcXFxcIl0ge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLWJyYW5kKTtcXFxcblxcXFxuICAgIC5uYXYtcGFja2FnZS1uYW1lIHtcXFxcbiAgICAgIGNvbG9yOiByZ2JhKDI1NSwgMjU1LCAyNTUsIDAuOCk7XFxcXG4gICAgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5uYXYtcGFja2FnZS1uYW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4vKiBJbmZvIGJ1dHRvbiBwb3BvdmVyIHRyaWdnZXJzIGluIGtub2JzIC0gb3ZlcnJpZGUgcGxhaW4gYnV0dG9uIHBhZGRpbmcgKi9cXFxcbnBmLXY2LXBvcG92ZXIgcGYtdjYtYnV0dG9uW3ZhcmlhbnQ9XFxcXFxcXCJwbGFpblxcXFxcXFwiXSB7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLW0tcGxhaW4tLVBhZGRpbmdJbmxpbmVFbmQ6IDA7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLW0tcGxhaW4tLVBhZGRpbmdJbmxpbmVTdGFydDogMDtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tTWluV2lkdGg6IGF1dG87XFxcXG59XFxcXG5cXFxcbi8qIEtub2IgZGVzY3JpcHRpb24gY29udGVudCAoc2xvdHRlZCBpbiBmb3JtIGdyb3VwIGhlbHBlcikgKi9cXFxcbnBmLXY2LWZvcm0tZ3JvdXAgW3Nsb3Q9XFxcXFxcXCJoZWxwZXJcXFxcXFxcIl0ge1xcXFxuICBwIHtcXFxcbiAgICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIDA7XFxcXG4gIH1cXFxcblxcXFxuICBjb2RlIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogM3B4O1xcXFxuICAgIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgfVxcXFxuXFxcXG4gIGEge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICAgIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICB0ZXh0LWRlY29yYXRpb246IHVuZGVybGluZTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBzdHJvbmcge1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIH1cXFxcblxcXFxuICB1bCwgb2wge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgICBwYWRkaW5nLWxlZnQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbGcpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIFN5bnRheCBoaWdobGlnaHRpbmcgKGNocm9tYSAtIHRoZW1hYmxlIHZpYSBDU1MgY3VzdG9tIHByb3BlcnRpZXMpICovXFxcXG5wZi12Ni1mb3JtLWdyb3VwIFtzbG90PVxcXFxcXFwiaGVscGVyXFxcXFxcXCJdIHtcXFxcbiAgLmNocm9tYSB7XFxcXG4gICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICAgIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1yYWRpdXMpO1xcXFxuICAgIG92ZXJmbG93LXg6IGF1dG87XFxcXG5cXFxcbiAgICBcXFxcdTAwMjYgLmxudGQgeyB2ZXJ0aWNhbC1hbGlnbjogdG9wOyBwYWRkaW5nOiAwOyBtYXJnaW46IDA7IGJvcmRlcjogMDsgfVxcXFxuICAgIFxcXFx1MDAyNiAubG50YWJsZSB7IGJvcmRlci1zcGFjaW5nOiAwOyBwYWRkaW5nOiAwOyBtYXJnaW46IDA7IGJvcmRlcjogMDsgfVxcXFxuICAgIFxcXFx1MDAyNiAuaGwgeyBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtaGlnaGxpZ2h0KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5sbnQsXFxcXG4gICAgXFxcXHUwMDI2IC5sbiB7XFxcXG4gICAgICB3aGl0ZS1zcGFjZTogcHJlO1xcXFxuICAgICAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG4gICAgICBtYXJnaW4tcmlnaHQ6IDAuNGVtO1xcXFxuICAgICAgcGFkZGluZzogMCAwLjRlbSAwIDAuNGVtO1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmxpbmUgeyBkaXNwbGF5OiBmbGV4OyB9XFxcXG5cXFxcbiAgICAvKiBLZXl3b3JkcyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuayxcXFxcbiAgICBcXFxcdTAwMjYgLmtjLFxcXFxuICAgIFxcXFx1MDAyNiAua2QsXFxcXG4gICAgXFxcXHUwMDI2IC5rbixcXFxcbiAgICBcXFxcdTAwMjYgLmtwLFxcXFxuICAgIFxcXFx1MDAyNiAua3Ige1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1rZXl3b3JkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmt0IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC10eXBlKTsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcblxcXFxuICAgIC8qIE5hbWVzICovXFxcXG4gICAgXFxcXHUwMDI2IC5uYSxcXFxcbiAgICBcXFxcdTAwMjYgLm5iLFxcXFxuICAgIFxcXFx1MDAyNiAubm8sXFxcXG4gICAgXFxcXHUwMDI2IC5udixcXFxcbiAgICBcXFxcdTAwMjYgLnZjLFxcXFxuICAgIFxcXFx1MDAyNiAudmcsXFxcXG4gICAgXFxcXHUwMDI2IC52aSB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LW5hbWUpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmJwIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5uYyB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtY2xhc3MpOyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuICAgIFxcXFx1MDAyNiAubmQgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWRlY29yYXRvcik7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5uaSxcXFxcbiAgICBcXFxcdTAwMjYgLnNzIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtc3BlY2lhbCk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAubmUsXFxcXG4gICAgXFxcXHUwMDI2IC5ubCB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWtleXdvcmQpO1xcXFxuICAgICAgZm9udC13ZWlnaHQ6IGJvbGQ7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAubmYgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWZ1bmN0aW9uKTsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLm5uIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5udCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtdGFnKSB9XFxcXG5cXFxcbiAgICAvKiBTdHJpbmdzICovXFxcXG4gICAgXFxcXHUwMDI2IC5zLFxcXFxuICAgIFxcXFx1MDAyNiAuc2EsXFxcXG4gICAgXFxcXHUwMDI2IC5zYixcXFxcbiAgICBcXFxcdTAwMjYgLnNjLFxcXFxuICAgIFxcXFx1MDAyNiAuZGwsXFxcXG4gICAgXFxcXHUwMDI2IC5zZCxcXFxcbiAgICBcXFxcdTAwMjYgLnMyLFxcXFxuICAgIFxcXFx1MDAyNiAuc2UsXFxcXG4gICAgXFxcXHUwMDI2IC5zaCxcXFxcbiAgICBcXFxcdTAwMjYgLnNpLFxcXFxuICAgIFxcXFx1MDAyNiAuc3gsXFxcXG4gICAgXFxcXHUwMDI2IC5zMSB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXN0cmluZyk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuc3IgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXRhZykgfVxcXFxuXFxcXG4gICAgLyogTnVtYmVycyAqL1xcXFxuICAgIFxcXFx1MDAyNiAubSxcXFxcbiAgICBcXFxcdTAwMjYgLm1iLFxcXFxuICAgIFxcXFx1MDAyNiAubWYsXFxcXG4gICAgXFxcXHUwMDI2IC5taCxcXFxcbiAgICBcXFxcdTAwMjYgLm1pLFxcXFxuICAgIFxcXFx1MDAyNiAuaWwsXFxcXG4gICAgXFxcXHUwMDI2IC5tbyB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LW51bWJlcik7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgLyogT3BlcmF0b3JzICovXFxcXG4gICAgXFxcXHUwMDI2IC5vLFxcXFxuICAgIFxcXFx1MDAyNiAub3cge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1rZXl3b3JkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIENvbW1lbnRzICovXFxcXG4gICAgXFxcXHUwMDI2IC5jLFxcXFxuICAgIFxcXFx1MDAyNiAuY2gsXFxcXG4gICAgXFxcXHUwMDI2IC5jbSxcXFxcbiAgICBcXFxcdTAwMjYgLmMxIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5jcyxcXFxcbiAgICBcXFxcdTAwMjYgLmNwLFxcXFxuICAgIFxcXFx1MDAyNiAuY3BmIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgICAgZm9udC1zdHlsZTogaXRhbGljO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIEVycm9ycyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuZXJyIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZXJyb3IpO1xcXFxuICAgICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWVycm9yLWJnKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBHZW5lcmljcyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuZ2Qge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1kZWxldGVkKTtcXFxcbiAgICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1kZWxldGVkLWJnKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5nZSB7IGZvbnQtc3R5bGU6IGl0YWxpYzsgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3IgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWVycm9yKSB9XFxcXG4gICAgXFxcXHUwMDI2IC5naCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2kge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1pbnNlcnRlZCk7XFxcXG4gICAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtaW5zZXJ0ZWQtYmcpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdvIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5ncCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3MgeyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3UgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmd0IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvcikgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2wgeyB0ZXh0LWRlY29yYXRpb246IHVuZGVybGluZTsgfVxcXFxuICAgIFxcXFx1MDAyNiAudyB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKSB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogRXZlbnRzIHRhYiBzdHlsaW5nIC0gUHJpbWFyeS1kZXRhaWwgbGF5b3V0ICovXFxcXG4uZXZlbnRzLXdyYXBwZXIge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbiNldmVudC1kcmF3ZXIge1xcXFxuICBmbGV4OiAxO1xcXFxuICBtaW4taGVpZ2h0OiAwO1xcXFxufVxcXFxuXFxcXG4vKiBFdmVudCBsaXN0IChwcmltYXJ5IHBhbmVsKSAqL1xcXFxuI2V2ZW50LWxpc3Qge1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1saXN0LWl0ZW0ge1xcXFxuICAvKiBSZXNldCBidXR0b24gc3R5bGVzICovXFxcXG4gIGFwcGVhcmFuY2U6IG5vbmU7XFxcXG4gIGJhY2tncm91bmQ6IG5vbmU7XFxcXG4gIGJvcmRlcjogbm9uZTtcXFxcbiAgYm9yZGVyLWxlZnQ6IDNweCBzb2xpZCB0cmFuc3BhcmVudDtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBmb250OiBpbmhlcml0O1xcXFxuICBjb2xvcjogaW5oZXJpdDtcXFxcbiAgdGV4dC1hbGlnbjogaW5oZXJpdDtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG5cXFxcbiAgLyogQ29tcG9uZW50IHN0eWxlcyAqL1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kIDEwMG1zIGVhc2UtaW4tb3V0LCBib3JkZXItY29sb3IgMTAwbXMgZWFzZS1pbi1vdXQ7XFxcXG5cXFxcbiAgcGYtdjYtbGFiZWwge1xcXFxuICAgIGZsZXgtc2hyaW5rOiAwO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1ob3Zlcik7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6Zm9jdXMge1xcXFxuICAgIG91dGxpbmU6IDJweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWNsaWNrZWQpO1xcXFxuICAgIG91dGxpbmUtb2Zmc2V0OiAtMnB4O1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2LnNlbGVjdGVkIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1hY3Rpb24tLXBsYWluLS1zZWxlY3RlZCk7XFxcXG4gICAgYm9yZGVyLWxlZnQtY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5ldmVudC10aW1lLFxcXFxuLmV2ZW50LWVsZW1lbnQge1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LXRpbWUge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZsZXgtc2hyaW5rOiAwO1xcXFxuICBmb250LXNpemU6IDExcHg7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1lbGVtZW50IHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmb250LXdlaWdodDogNDAwO1xcXFxufVxcXFxuXFxcXG4vKiBFdmVudCBkZXRhaWwgcGFuZWwgKi9cXFxcbi5ldmVudC1kZXRhaWwtaGVhZGVyLWNvbnRlbnQge1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgYm9yZGVyLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXdpZHRoKSBzb2xpZCB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItY29sb3IpO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLW5hbWUge1xcXFxuICBtYXJnaW46IDAgMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSAwO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1sZyk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLXN1bW1hcnkge1xcXFxuICBtYXJnaW46IDAgMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSAwO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjU7XFxcXG4gIHdoaXRlLXNwYWNlOiBwcmUtd3JhcDtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1kZXNjcmlwdGlvbiB7XFxcXG4gIG1hcmdpbjogMCAwIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIDA7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgbGluZS1oZWlnaHQ6IDEuNTtcXFxcbiAgd2hpdGUtc3BhY2U6IHByZS13cmFwO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLW1ldGEge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtdGltZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLWVsZW1lbnQge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtcHJvcGVydGllcy1oZWFkaW5nIHtcXFxcbiAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtYmFzZSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMge1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctc2Vjb25kYXJ5KTtcXFxcbiAgYm9yZGVyOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItd2lkdGgpIHNvbGlkIHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1jb2xvcik7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1yYWRpdXMpO1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIGZvbnQtc2l6ZTogMTJweDtcXFxcbiAgbGluZS1oZWlnaHQ6IDEuNjtcXFxcbiAgbWFyZ2luOiAwIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtcHJvcGVydHktdHJlZSB7XFxcXG4gIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gIHBhZGRpbmc6IDA7XFxcXG4gIG1hcmdpbjogMDtcXFxcblxcXFxuICBcXFxcdTAwMjYubmVzdGVkIHtcXFxcbiAgICBwYWRkaW5nLWxlZnQ6IDEuNWVtO1xcXFxuICAgIG1hcmdpbi10b3A6IDAuMjVlbTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ucHJvcGVydHktaXRlbSB7XFxcXG4gIHBhZGRpbmc6IDAuMTI1ZW0gMDtcXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LWtleSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICBmb250LXdlaWdodDogNTAwO1xcXFxufVxcXFxuXFxcXG4ucHJvcGVydHktY29sb24ge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG59XFxcXG5cXFxcbi5wcm9wZXJ0eS12YWx1ZSB7XFxcXG4gIFxcXFx1MDAyNi5udWxsLFxcXFxuICBcXFxcdTAwMjYudW5kZWZpbmVkIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gICAgZm9udC1zdHlsZTogaXRhbGljO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2LmJvb2xlYW4ge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1jb2xvci1ib29sZWFuKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5udW1iZXIge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1jb2xvci1udW1iZXIpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2LnN0cmluZyB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWNvbG9yLXN0cmluZyk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYuYXJyYXksXFxcXG4gIFxcXFx1MDAyNi5vYmplY3Qge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuI2RlYnVnLW1vZGFsIHtcXFxcbiAgY29udGFpbmVyLXR5cGU6IGlubGluZS1zaXplO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksTUFBTSxlQUFlO0FBQzFDLFNBQVMscUJBQXFCO0FBQzlCLFNBQVMsZ0JBQWdCO0FBRXpCLFNBQVMsaUJBQWlCOzs7QUNKMUIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sb29zQkFBc3BzQixDQUFDO0FBQUUsSUFBTywyQkFBUTs7O0FEUTd0c0IsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBWVAsSUFBSTtBQUNKLElBQUk7QUF3REcsSUFBTSxlQUFOLGNBQTJCLE1BQU07QUFBQSxFQUN0QztBQUFBLEVBQ0EsWUFBWSxNQUE4RDtBQUN4RSxVQUFNLFlBQVksRUFBRSxTQUFTLEtBQUssQ0FBQztBQUNuQyxTQUFLLE9BQU87QUFBQSxFQUNkO0FBQ0Y7QUE1SEE7QUEySUEsOEJBQUMsY0FBYyxrQkFBa0I7QUFDMUIsSUFBTSxrQkFBTixNQUFNLHlCQUF1QixpQkFnRmxDLHVCQUFDLFNBQVMsRUFBRSxXQUFXLG1CQUFtQixDQUFDLElBRzNDLGtCQUFDLFNBQVMsRUFBRSxXQUFXLGFBQWEsQ0FBQyxJQUdyQyxvQkFBQyxTQUFTLEVBQUUsV0FBVyxlQUFlLENBQUMsSUFHdkMscUJBQUMsU0FBUyxFQUFFLFdBQVcsZ0JBQWdCLENBQUMsSUFHeEMsa0JBQUMsU0FBUyxFQUFFLFdBQVcsYUFBYSxDQUFDLElBR3JDLGNBQUMsU0FBUyxJQUdWLGVBQUMsU0FBUyxJQUdWLHFCQUFDLFNBQVMsRUFBRSxXQUFXLGdCQUFnQixDQUFDLElBR3hDLHFCQUFDLFNBQVMsRUFBRSxXQUFXLGdCQUFnQixDQUFDLElBR3hDLGdCQUFDLFNBQVMsSUFHVix1QkFBQyxTQUFTLEVBQUUsTUFBTSxTQUFTLFdBQVcsa0JBQWtCLENBQUMsSUE5R3ZCLElBQVc7QUFBQSxFQUF4QztBQUFBO0FBQUE7QUFpRkwsdUJBQVMsaUJBQWlCLGtCQUExQixnQkFBMEIsTUFBMUI7QUFHQSx1QkFBUyxZQUFZLGtCQUFyQixpQkFBcUIsTUFBckI7QUFHQSx1QkFBUyxjQUFjLGtCQUF2QixpQkFBdUIsTUFBdkI7QUFHQSx1QkFBUyxlQUFlLGtCQUF4QixpQkFBd0IsTUFBeEI7QUFHQSx1QkFBUyxZQUFZLGtCQUFyQixpQkFBcUIsTUFBckI7QUFHQSx1QkFBUyxRQUFRLGtCQUFqQixpQkFBaUIsTUFBakI7QUFHQSx1QkFBUyxTQUF3QyxrQkFBakQsaUJBQWlELE1BQWpEO0FBR0EsdUJBQVMsZUFBZSxrQkFBeEIsaUJBQXdCLE1BQXhCO0FBR0EsdUJBQVMsZUFBZSxrQkFBeEIsaUJBQXdCLE1BQXhCO0FBR0EsdUJBQVMsVUFBeUMsa0JBQWxELGlCQUFrRCxNQUFsRDtBQUdBLHVCQUFTLGlCQUFpQixrQkFBMUIsaUJBQTBCLFNBQTFCO0FBVUEsc0NBQW9DO0FBQ3BDLG9DQUFjO0FBQ2QsNENBQXNCO0FBQ3RCLHVDQUFpQjtBQUNqQixtQ0FBK0I7QUFHL0I7QUFBQSx5Q0FBa0Q7QUFDbEQsa0NBQTZCO0FBQzdCLHdDQUFpQyxDQUFDO0FBQ2xDLDJDQUFxQjtBQUNyQixtQ0FBaUM7QUFDakMsMkNBQXlDO0FBQ3pDLHlDQUF1QztBQUN2Qyx5Q0FBa0M7QUFDbEMsMkNBQXFCO0FBQ3JCLG1EQUFtRTtBQUNuRSwwQ0FBb0Isb0JBQUksSUFBWTtBQUNwQyx3Q0FBa0Isb0JBQUksSUFBWTtBQUNsQyw0Q0FBc0Isb0JBQUksSUFBWTtBQUd0QztBQUFBLHlDQUFvRDtBQUNwRCwwQ0FBcUQ7QUFDckQsNENBQXVEO0FBQ3ZELDBDQUFxRDtBQUdyRDtBQUFBLGlEQUFpRTtBQUNqRSxrREFBa0U7QUFDbEUsbURBQW1FO0FBR25FO0FBQUEseUNBQW1CO0FBQ25CLGlEQUFpRTtBQUNqRSx5Q0FBbUIsb0JBQUksSUFBSSxDQUFDLFFBQVEsUUFBUSxTQUFTLE9BQU8sQ0FBQztBQUM3RCwwQ0FBb0M7QUFJcEM7QUFBQTtBQUFBLGtDQUFZLElBQUksaUJBQWlCLENBQUMsY0FBYztBQUM5QyxVQUFJLGNBQWM7QUFFbEIsaUJBQVcsWUFBWSxXQUFXO0FBQ2hDLG1CQUFXLFFBQVEsU0FBUyxZQUFZO0FBQ3RDLGNBQUksZ0JBQWdCLGFBQWE7QUFDL0Isa0JBQU0sVUFBVSxLQUFLLFFBQVEsWUFBWTtBQUN6QyxnQkFBSSxtQkFBSyxtQkFBa0IsSUFBSSxPQUFPLEtBQUssQ0FBQyxLQUFLLFFBQVEsbUJBQW1CO0FBQzFFLG9CQUFNLFlBQVksbUJBQUssa0JBQWlCLElBQUksT0FBTztBQUNuRCx5QkFBVyxhQUFhLFVBQVUsWUFBWTtBQUM1QyxxQkFBSyxpQkFBaUIsV0FBVyxtQkFBSyxzQkFBcUIsRUFBRSxTQUFTLEtBQUssQ0FBQztBQUFBLGNBQzlFO0FBQ0EsbUJBQUssUUFBUSxvQkFBb0I7QUFDakMsaUNBQUsscUJBQW9CLElBQUksT0FBTztBQUNwQyw0QkFBYztBQUFBLFlBQ2hCO0FBQUEsVUFDRjtBQUFBLFFBQ0Y7QUFBQSxNQUNGO0FBRUEsVUFBSSxhQUFhO0FBQ2YsOEJBQUssa0RBQUw7QUFBQSxNQUNGO0FBQUEsSUFDRixDQUFDO0FBR0Q7QUFBQTtBQUNBLDZDQUF1QjtBQW1uQnZCLCtDQUF5QixDQUFDLFVBQWlCO0FBQ3pDLFlBQU0sRUFBRSxPQUFPLFFBQVEsSUFBSTtBQUUzQixVQUFJLFNBQVM7QUFDWCwyQkFBSyxrQkFBaUIsSUFBSSxLQUFLO0FBQUEsTUFDakMsT0FBTztBQUNMLDJCQUFLLGtCQUFpQixPQUFPLEtBQUs7QUFBQSxNQUNwQztBQUVBLDRCQUFLLGtEQUFMO0FBQ0EsNEJBQUssMENBQUwsV0FBaUIsbUJBQUs7QUFBQSxJQUN4QjtBQWtWQSxzQ0FBZ0IsQ0FBQyxVQUFpQjtBQUNoQyxZQUFNLFNBQVMsc0JBQUssNkNBQUwsV0FBb0I7QUFDbkMsVUFBSSxDQUFDLFFBQVE7QUFDWCxnQkFBUSxLQUFLLGtFQUFrRTtBQUMvRTtBQUFBLE1BQ0Y7QUFFQSxZQUFNLEVBQUUsU0FBUyxjQUFjLElBQUk7QUFFbkMsWUFBTSxPQUFPLEtBQUs7QUFDbEIsVUFBSSxDQUFDLEtBQU07QUFFWCxZQUFNLFdBQVcsc0JBQUssb0RBQUwsV0FBMkI7QUFFNUMsWUFBTSxVQUFXLEtBQWE7QUFBQSxRQUM1QjtBQUFBLFFBQ0MsTUFBYztBQUFBLFFBQ2QsTUFBYztBQUFBLFFBQ2Y7QUFBQSxRQUNBO0FBQUEsTUFDRjtBQUVBLFVBQUksQ0FBQyxTQUFTO0FBQ1osZ0JBQVEsS0FBSyxtREFBbUQ7QUFBQSxVQUM5RCxNQUFNO0FBQUEsVUFDTixNQUFPLE1BQWM7QUFBQSxVQUNyQjtBQUFBLFVBQ0E7QUFBQSxRQUNGLENBQUM7QUFBQSxNQUNIO0FBQUEsSUFDRjtBQUVBLHFDQUFlLENBQUMsVUFBaUI7QUFDL0IsWUFBTSxTQUFTLHNCQUFLLDZDQUFMLFdBQW9CO0FBQ25DLFVBQUksQ0FBQyxRQUFRO0FBQ1gsZ0JBQVEsS0FBSyxrRUFBa0U7QUFDL0U7QUFBQSxNQUNGO0FBRUEsWUFBTSxFQUFFLFNBQVMsY0FBYyxJQUFJO0FBRW5DLFlBQU0sT0FBTyxLQUFLO0FBQ2xCLFVBQUksQ0FBQyxLQUFNO0FBRVgsWUFBTSxXQUFXLHNCQUFLLHlEQUFMLFdBQWdDO0FBQ2pELFlBQU0sYUFBYSxhQUFhLGFBQWEsU0FBWTtBQUV6RCxZQUFNLFVBQVcsS0FBYTtBQUFBLFFBQzVCO0FBQUEsUUFDQyxNQUFjO0FBQUEsUUFDZjtBQUFBLFFBQ0E7QUFBQSxRQUNBO0FBQUEsTUFDRjtBQUVBLFVBQUksQ0FBQyxTQUFTO0FBQ1osZ0JBQVEsS0FBSyw0Q0FBNEM7QUFBQSxVQUN2RCxNQUFNO0FBQUEsVUFDTixNQUFPLE1BQWM7QUFBQSxVQUNyQjtBQUFBLFVBQ0E7QUFBQSxRQUNGLENBQUM7QUFBQSxNQUNIO0FBQUEsSUFDRjtBQWlQQSw0Q0FBc0IsQ0FBQyxVQUFpQjtBQUN0QyxZQUFNLFVBQVUsTUFBTTtBQUN0QixVQUFJLEVBQUUsbUJBQW1CLGFBQWM7QUFFdkMsWUFBTSxVQUFVLFFBQVEsUUFBUSxZQUFZO0FBQzVDLFlBQU0sWUFBWSxtQkFBSyxtQkFBa0IsSUFBSSxPQUFPO0FBRXBELFVBQUksQ0FBQyxhQUFhLENBQUMsVUFBVSxXQUFXLElBQUksTUFBTSxJQUFJLEVBQUc7QUFFekQseUJBQUsscUJBQW9CLElBQUksT0FBTztBQUNwQyw0QkFBSyw0Q0FBTCxXQUFtQixPQUFPLFNBQVMsU0FBUztBQUFBLElBQzlDO0FBd2RBLHFEQUErQixDQUFDLFVBQWlCO0FBQy9DLFlBQU0sRUFBRSxPQUFPLFFBQVEsSUFBSTtBQUUzQixVQUFJLENBQUMsTUFBTztBQUVaLFVBQUksU0FBUztBQUNYLDJCQUFLLG1CQUFrQixJQUFJLEtBQUs7QUFBQSxNQUNsQyxPQUFPO0FBQ0wsMkJBQUssbUJBQWtCLE9BQU8sS0FBSztBQUFBLE1BQ3JDO0FBRUEsNEJBQUssZ0RBQUw7QUFDQSw0QkFBSyw0Q0FBTCxXQUFtQixtQkFBSztBQUFBLElBQzFCO0FBRUEsbURBQTZCLENBQUMsVUFBaUI7QUFDN0MsWUFBTSxFQUFFLE9BQU8sUUFBUSxJQUFJO0FBRTNCLFVBQUksQ0FBQyxNQUFPO0FBRVosVUFBSSxTQUFTO0FBQ1gsMkJBQUssaUJBQWdCLElBQUksS0FBSztBQUFBLE1BQ2hDLE9BQU87QUFDTCwyQkFBSyxpQkFBZ0IsT0FBTyxLQUFLO0FBQUEsTUFDbkM7QUFFQSw0QkFBSyxnREFBTDtBQUNBLDRCQUFLLDRDQUFMLFdBQW1CLG1CQUFLO0FBQUEsSUFDMUI7QUFBQTtBQUFBLEVBbHREQSxJQUFJLE9BQXVCO0FBQ3pCLFdBQU8sS0FBSyxjQUFjLGdCQUFnQjtBQUFBLEVBQzVDO0FBQUEsRUFFQSxTQUFTO0FBQ1AsV0FBTztBQUFBO0FBQUE7QUFBQTtBQUFBLHVDQUk0QixLQUFLLFlBQVksV0FBVztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxjQVlyRCxLQUFLLGNBQWMsV0FBVyxLQUFLLFdBQVcsVUFBVSxPQUFPO0FBQUE7QUFBQTtBQUFBO0FBQUEsZ0JBSTdELHNCQUFLLGtEQUFMLFVBQTBCO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSx3Q0F5Q0YsS0FBSyxZQUFZLFVBQVU7QUFBQSx5Q0FDMUIsS0FBSyxZQUFZLFVBQVU7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSw0Q0FPeEIsS0FBSyxpQkFBaUIsS0FBSyxRQUFRO0FBQUE7QUFBQTtBQUFBLGdDQUcvQyxLQUFLLFdBQVcsVUFBVTtBQUFBLHlDQUNqQixLQUFLLGdCQUFnQixLQUFLO0FBQUEsc0NBQzdCLEtBQUssZ0JBQWdCLEdBQUc7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxnREFzSGQsVUFBVSxLQUFLLGNBQWMsQ0FBQztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLGdFQWtDZCxLQUFLLGtCQUFrQixHQUFHO0FBQUE7QUFBQTtBQUFBO0FBQUEsZ0VBSTFCLEtBQUssYUFBYSxHQUFHO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBMENuRjtBQUFBLEVBc0NBLE1BQU0sb0JBQW9CO0FBQ3hCLFVBQU0sa0JBQWtCO0FBR3hCLFFBQUksQ0FBQyxtQkFBSyx1QkFBc0I7QUFDOUIsT0FBQyxFQUFFLGdCQUFnQixHQUFHLEVBQUUsaUJBQWlCLENBQUMsSUFBSSxNQUFNLFFBQVEsSUFBSTtBQUFBO0FBQUEsUUFFOUQsT0FBTyw0QkFBNEI7QUFBQTtBQUFBLFFBRW5DLE9BQU8sNkJBQTZCO0FBQUEsTUFDdEMsQ0FBQztBQUVELGFBQU8seUJBQXlCLEVBQUUsTUFBTSxDQUFDLE1BQ3ZDLFFBQVEsTUFBTSw2Q0FBNkMsQ0FBQyxDQUFDO0FBQy9ELHlCQUFLLHNCQUF1QjtBQUM1Qiw0QkFBSyw0Q0FBTDtBQUNBLDRCQUFLLCtEQUFMO0FBQUEsSUFDRjtBQUFBLEVBQ0Y7QUFBQSxFQUVBLGVBQWU7QUFFYiwwQkFBSyxpREFBTDtBQUdBLDBCQUFLLHNEQUFMO0FBR0EsMEJBQUssaURBQUw7QUFHQSwwQkFBSyxnREFBTDtBQUdBLDBCQUFLLHFEQUFMO0FBR0EsMEJBQUsseURBQUw7QUFHQSwwQkFBSyw0REFBTDtBQUdBLDBCQUFLLGlEQUFMLFdBQTBCLEtBQUssTUFBTTtBQUNuQyw0QkFBSyxtREFBTDtBQUFBLElBQ0YsQ0FBQztBQUlELDBCQUFLLGlDQUFMLFdBQVEsa0JBQWtCLGlCQUFpQixTQUFTLE1BQU07QUFDeEQsYUFBTyxTQUFTLE9BQU87QUFBQSxJQUN6QixDQUFDO0FBR0QsMEJBQUssaUNBQUwsV0FBUSxpQkFBaUIsaUJBQWlCLFNBQVMsTUFBTTtBQUN2RCxNQUFDLHNCQUFLLGlDQUFMLFdBQVEsdUJBQStCLE1BQU07QUFDOUMseUJBQUssV0FBVSxNQUFNO0FBQUEsSUFDdkIsQ0FBQztBQUdELHVCQUFLLFdBQVUsS0FBSztBQUFBLEVBQ3RCO0FBQUEsRUE2K0NBLHVCQUF1QjtBQUNyQixVQUFNLHFCQUFxQjtBQUczQixTQUFLLG9CQUFvQix5QkFBeUIsbUJBQUssY0FBYTtBQUNwRSxTQUFLLG9CQUFvQix3QkFBd0IsbUJBQUssY0FBYTtBQUNuRSxTQUFLLG9CQUFvQiw0QkFBNEIsbUJBQUssY0FBYTtBQUN2RSxTQUFLLG9CQUFvQix3QkFBd0IsbUJBQUssYUFBWTtBQUNsRSxTQUFLLG9CQUFvQix1QkFBdUIsbUJBQUssYUFBWTtBQUNqRSxTQUFLLG9CQUFvQiwyQkFBMkIsbUJBQUssYUFBWTtBQUdyRSxRQUFJLG1CQUFLLG9CQUFtQjtBQUMxQixXQUFLLG9CQUFvQixVQUFVLG1CQUFLLGtCQUFpQjtBQUFBLElBQzNEO0FBQ0EsUUFBSSxtQkFBSyxzQkFBcUI7QUFDNUIsV0FBSyxvQkFBb0IsWUFBWSxtQkFBSyxvQkFBbUI7QUFBQSxJQUMvRDtBQUNBLFFBQUksbUJBQUssb0JBQW1CO0FBQzFCLFdBQUssb0JBQW9CLFVBQVUsbUJBQUssa0JBQWlCO0FBQUEsSUFDM0Q7QUFHQSxRQUFJLG1CQUFLLG1CQUFrQjtBQUN6QixhQUFPLG9CQUFvQixZQUFZLG1CQUFLLGlCQUFnQjtBQUFBLElBQzlEO0FBR0EsUUFBSSxtQkFBSywyQkFBMEI7QUFDakMsbUJBQWEsbUJBQUsseUJBQXdCO0FBQzFDLHlCQUFLLDBCQUEyQjtBQUFBLElBQ2xDO0FBQ0EsUUFBSSxtQkFBSyw0QkFBMkI7QUFDbEMsbUJBQWEsbUJBQUssMEJBQXlCO0FBQzNDLHlCQUFLLDJCQUE0QjtBQUFBLElBQ25DO0FBQ0EsUUFBSSxtQkFBSyw2QkFBNEI7QUFDbkMsbUJBQWEsbUJBQUssMkJBQTBCO0FBQzVDLHlCQUFLLDRCQUE2QjtBQUFBLElBQ3BDO0FBR0EsdUJBQUssV0FBVSxXQUFXO0FBRzFCLFFBQUksbUJBQUssWUFBVztBQUNsQix5QkFBSyxXQUFVLFFBQVE7QUFBQSxJQUN6QjtBQUFBLEVBQ0Y7QUFDRjtBQS9uRU87QUFJRTtBQTZCQTtBQWVBO0FBVUE7QUFXQTtBQVlFO0FBR0E7QUFHQTtBQUdBO0FBR0E7QUFHQTtBQUdBO0FBR0E7QUFHQTtBQUdBO0FBR0E7QUEvR0o7QUFpSEwsT0FBRSxTQUFDLElBQVk7QUFDYixTQUFPLEtBQUssWUFBWSxlQUFlLEVBQUU7QUFDM0M7QUFFQSxRQUFHLFNBQUMsVUFBa0I7QUFDcEIsU0FBTyxLQUFLLFlBQVksaUJBQWlCLFFBQVEsS0FBSyxDQUFDO0FBQ3pEO0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUdBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBR0E7QUFDQTtBQUNBO0FBQ0E7QUFHQTtBQUNBO0FBQ0E7QUFHQTtBQUNBO0FBQ0E7QUFDQTtBQUlBO0FBMEJBO0FBQ0E7QUFFQSxrQkFBYSxXQUFHO0FBQ2QscUJBQUssV0FBWSxJQUFJLGdCQUFnQjtBQUFBLElBQ3JDLFdBQVc7QUFBQSxJQUNYLGtCQUFrQjtBQUFBLElBQ2xCLGdCQUFnQjtBQUFBO0FBQUEsSUFFaEIsV0FBVztBQUFBLE1BQ1QsUUFBUSxNQUFNO0FBQ1osOEJBQUssaUNBQUwsV0FBUSx1QkFBdUIsTUFBTTtBQUFBLE1BQ3ZDO0FBQUEsTUFDQSxTQUFTLENBQUMsY0FBbUU7QUFDM0UsWUFBSSxXQUFXLFNBQVMsV0FBVyxTQUFTO0FBQzFDLGtCQUFRLE1BQU0sNkJBQTZCLFNBQVM7QUFDcEQsVUFBQyxzQkFBSyxpQ0FBTCxXQUFRLGtCQUEwQixLQUFLLFVBQVUsT0FBTyxVQUFVLFNBQVMsVUFBVSxJQUFJO0FBQUEsUUFDNUYsT0FBTztBQUNMLGtCQUFRLE1BQU0sZ0NBQWdDLFNBQVM7QUFBQSxRQUN6RDtBQUFBLE1BQ0Y7QUFBQSxNQUNBLGdCQUFnQixDQUFDLEVBQUUsU0FBUyxNQUFNLE1BQTBDO0FBQzFFLFlBQUksV0FBVyxJQUFJO0FBQ2pCLFVBQUMsc0JBQUssaUNBQUwsV0FBUSx1QkFBK0IsVUFBVTtBQUNsRCxVQUFDLHNCQUFLLGlDQUFMLFdBQVEseUJBQWlDLGdCQUFnQixTQUFTLEtBQUs7QUFBQSxRQUMxRTtBQUFBLE1BQ0Y7QUFBQSxNQUNBLFVBQVUsTUFBTTtBQUNkLGNBQU0sZUFBZSxzQkFBSyxpQ0FBTCxXQUFRO0FBQzdCLFlBQUksY0FBYyxhQUFhLE1BQU0sR0FBRztBQUN0QyxVQUFDLGFBQXFCLEtBQUs7QUFBQSxRQUM3QjtBQUNBLGVBQU8sU0FBUyxPQUFPO0FBQUEsTUFDekI7QUFBQSxNQUNBLFlBQVksTUFBTTtBQUNoQixRQUFDLHNCQUFLLGlDQUFMLFdBQVEsdUJBQStCLFVBQVU7QUFDbEQsUUFBQyxzQkFBSyxpQ0FBTCxXQUFRLHlCQUFpQyxnQkFBZ0IsSUFBSSxHQUFLO0FBQUEsTUFDckU7QUFBQSxNQUNBLFFBQVEsQ0FBQyxTQUFpRTtBQUN4RSxlQUFPLGNBQWMsSUFBSSxhQUFhLElBQUksQ0FBQztBQUFBLE1BQzdDO0FBQUEsSUFDRjtBQUFBO0FBQUEsRUFFQSxDQUFDO0FBQ0g7QUF5UkEsd0JBQW1CLFdBQUc7QUFDcEIsTUFBSSxDQUFDLEtBQUssVUFBVyxRQUFPO0FBRTVCLE1BQUksUUFBUTtBQUNaLE1BQUksT0FBTztBQUVYLE1BQUksS0FBSyxVQUFVLFNBQVMsWUFBWSxHQUFHO0FBQ3pDLFlBQVE7QUFDUixXQUFPO0FBQUEsRUFDVCxXQUFXLEtBQUssVUFBVSxTQUFTLFlBQVksR0FBRztBQUNoRCxZQUFRO0FBQ1IsV0FBTztBQUFBLEVBQ1QsV0FBVyxLQUFLLFVBQVUsU0FBUyxlQUFlLEdBQUc7QUFDbkQsWUFBUTtBQUNSLFdBQU87QUFBQSxFQUNUO0FBRUEsU0FBTztBQUFBO0FBQUEsOEJBRW1CLEtBQUssU0FBUztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsNkJBS2YsS0FBSztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsdUJBS1gsSUFBSTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBS3pCO0FBaUVNLG9CQUFlLGlCQUFHO0FBRXRCLFFBQU0sWUFBWSxzQkFBSyxpQ0FBTCxXQUFRO0FBQzFCLFFBQU0sT0FBTyxzQkFBSyxpQ0FBTCxXQUFRO0FBQ3JCLE1BQUksV0FBVztBQUNiLFVBQU0sVUFBVSxzQkFBSyw2Q0FBTDtBQUNoQixjQUFVLGNBQWM7QUFBQSxFQUMxQjtBQUNBLE1BQUksTUFBTTtBQUNSLFNBQUssY0FBYyxVQUFVO0FBQUEsRUFDL0I7QUFHQSxRQUFNLE9BQU8sTUFBTSxNQUFNLGNBQWMsRUFDcEMsS0FBSyxTQUFPLElBQUksS0FBSyxDQUFDLEVBQ3RCLE1BQU0sQ0FBQyxRQUFlO0FBQ3JCLFlBQVEsTUFBTSxrREFBa0QsR0FBRztBQUFBLEVBQ3JFLENBQUM7QUFFSCxNQUFJLENBQUMsS0FBTTtBQUNYLFFBQU0sWUFBWSxzQkFBSyxpQ0FBTCxXQUFRO0FBQzFCLFFBQU0sT0FBTyxzQkFBSyxpQ0FBTCxXQUFRO0FBQ3JCLFFBQU0sYUFBYSxzQkFBSyxpQ0FBTCxXQUFRO0FBQzNCLFFBQU0saUJBQWlCLHNCQUFLLGlDQUFMLFdBQVE7QUFDL0IsUUFBTSxjQUFjLHNCQUFLLGlDQUFMLFdBQVE7QUFDNUIsUUFBTSxvQkFBb0Isc0JBQUssaUNBQUwsV0FBUTtBQUNsQyxRQUFNLGNBQWMsc0JBQUssaUNBQUwsV0FBUTtBQUU1QixNQUFJLFVBQVcsV0FBVSxjQUFjLEtBQUssV0FBVztBQUN2RCxNQUFJLEtBQU0sTUFBSyxjQUFjLEtBQUssTUFBTTtBQUN4QyxNQUFJLFdBQVksWUFBVyxjQUFjLEtBQUssWUFBWTtBQUMxRCxNQUFJLGVBQWdCLGdCQUFlLGNBQWMsS0FBSyxnQkFBZ0I7QUFDdEUsTUFBSSxZQUFhLGFBQVksY0FBYyxLQUFLLGFBQWE7QUFFN0QsTUFBSSxtQkFBbUI7QUFDckIsMEJBQUssZ0RBQUwsV0FBdUIsbUJBQW1CLEtBQUs7QUFBQSxFQUNqRDtBQUVBLE1BQUksZUFBZSxLQUFLLFdBQVc7QUFDakMsVUFBTSxnQkFBZ0IsS0FBSyxVQUFVLEtBQUssV0FBVyxNQUFNLENBQUM7QUFDNUQsZ0JBQVksY0FBYztBQUMxQixTQUFLLGdCQUFnQjtBQUFBLEVBQ3ZCLFdBQVcsYUFBYTtBQUN0QixnQkFBWSxjQUFjO0FBQUEsRUFDNUI7QUFFQSxxQkFBSyxZQUFhO0FBQ3BCO0FBRUEsc0JBQWlCLFNBQUMsV0FBd0IsT0FBMkI7QUFDbkUsTUFBSSxDQUFDLE9BQU8sUUFBUTtBQUNsQixjQUFVLGNBQWM7QUFDeEI7QUFBQSxFQUNGO0FBRUEsUUFBTSxpQkFBaUIsS0FBSyxrQkFBa0I7QUFDOUMsUUFBTSxlQUFlLENBQUMsQ0FBQztBQUV2QixNQUFJLGNBQWM7QUFDaEIsVUFBTSxjQUFjLE1BQU0sS0FBSyxVQUFRLEtBQUssWUFBWSxjQUFjO0FBQ3RFLFFBQUksQ0FBQyxhQUFhO0FBQ2hCLGdCQUFVLGNBQWM7QUFDeEI7QUFBQSxJQUNGO0FBRUEsVUFBTSxXQUFXLDhCQUFlLG1CQUFrQixRQUFRLFVBQVUsSUFBSTtBQUV4RSxhQUFTLGNBQWMsd0JBQXdCLEVBQUcsY0FBYyxZQUFZO0FBQzVFLGFBQVMsY0FBYyw2QkFBNkIsRUFBRyxjQUFjLFlBQVk7QUFDakYsYUFBUyxjQUFjLDJCQUEyQixFQUFHLGNBQWMsWUFBWTtBQUUvRSxVQUFNLG1CQUFtQixTQUFTLGNBQWMsa0NBQWtDO0FBQ2xGLFFBQUksWUFBWSxhQUFhO0FBQzNCLGVBQVMsY0FBYyw0QkFBNEIsRUFBRyxjQUFjLFlBQVk7QUFBQSxJQUNsRixPQUFPO0FBQ0wsd0JBQWtCLE9BQU87QUFBQSxJQUMzQjtBQUVBLFVBQU0sZ0JBQWdCLFNBQVMsY0FBYywrQkFBK0I7QUFDNUUsUUFBSSxZQUFZLFVBQVU7QUFDeEIsZUFBUyxjQUFjLHlCQUF5QixFQUFHLGNBQWMsWUFBWTtBQUFBLElBQy9FLE9BQU87QUFDTCxxQkFBZSxPQUFPO0FBQUEsSUFDeEI7QUFFQSxjQUFVLGdCQUFnQixRQUFRO0FBQUEsRUFDcEMsT0FBTztBQUNMLFVBQU0sZUFBZSw4QkFBZSxtQkFBa0IsUUFBUSxVQUFVLElBQUk7QUFFNUUsVUFBTSxrQkFBa0IsYUFBYSxjQUFjLDJCQUEyQjtBQUU5RSxlQUFXLFFBQVEsT0FBTztBQUN4QixZQUFNLGdCQUFnQiw4QkFBZSxvQkFBbUIsUUFBUSxVQUFVLElBQUk7QUFFOUUsb0JBQWMsY0FBYyx3QkFBd0IsRUFBRyxjQUFjLEtBQUs7QUFDMUUsb0JBQWMsY0FBYyw0QkFBNEIsRUFBRyxjQUN6RCxLQUFLLGVBQWU7QUFDdEIsb0JBQWMsY0FBYyw2QkFBNkIsRUFBRyxjQUFjLEtBQUs7QUFDL0Usb0JBQWMsY0FBYywyQkFBMkIsRUFBRyxjQUFjLEtBQUs7QUFFN0UsWUFBTSxnQkFBZ0IsY0FBYyxjQUFjLCtCQUErQjtBQUNqRixVQUFJLEtBQUssVUFBVTtBQUNqQixzQkFBYyxjQUFjLHlCQUF5QixFQUFHLGNBQWMsS0FBSztBQUFBLE1BQzdFLE9BQU87QUFDTCx1QkFBZSxPQUFPO0FBQUEsTUFDeEI7QUFFQSxzQkFBZ0IsWUFBWSxhQUFhO0FBQUEsSUFDM0M7QUFFQSxjQUFVLGdCQUFnQixZQUFZO0FBQUEsRUFDeEM7QUFDRjtBQUVBLHNCQUFpQixXQUFHO0FBQ2xCLHFCQUFLLGVBQWdCLHNCQUFLLGlDQUFMLFdBQVE7QUFFN0IsUUFBTSxhQUFhLHNCQUFLLGlDQUFMLFdBQVE7QUFDM0IsTUFBSSxZQUFZO0FBQ2QsZUFBVyxpQkFBaUIsU0FBUyxNQUFNO0FBQ3pDLFlBQU0sRUFBRSxRQUFRLEdBQUcsSUFBSTtBQUN2QixtQkFBYSxtQkFBSyx5QkFBeUI7QUFDM0MseUJBQUssMEJBQTJCLFdBQVcsTUFBTTtBQUMvQyw4QkFBSywwQ0FBTCxXQUFpQjtBQUFBLE1BQ25CLEdBQUcsR0FBRztBQUFBLElBQ1IsQ0FBQztBQUFBLEVBQ0g7QUFFQSxxQkFBSyxtQkFBb0IsS0FBSyxZQUFZLGNBQWMsbUJBQW1CLEtBQUs7QUFDaEYsTUFBSSxtQkFBSyxvQkFBbUI7QUFDMUIsMEJBQXNCLE1BQU07QUFDMUIsNEJBQUssa0RBQUw7QUFBQSxJQUNGLENBQUM7QUFDRCx1QkFBSyxtQkFBa0IsaUJBQWlCLFVBQVUsbUJBQUssdUJBQXVDO0FBQUEsRUFDaEc7QUFFQSx3QkFBSyxpQ0FBTCxXQUFRLGNBQWMsaUJBQWlCLFNBQVMsTUFBTTtBQUNwRCwwQkFBSyx3Q0FBTDtBQUFBLEVBQ0YsQ0FBQztBQUVELHFCQUFLLG1CQUFvQixDQUFDLFVBQWlCO0FBQ3pDLFVBQU0sT0FBUSxNQUF1QjtBQUNyQyxRQUFJLE1BQU07QUFDUiw0QkFBSywwQ0FBTCxXQUFpQjtBQUFBLElBQ25CO0FBQUEsRUFDRjtBQUNBLFNBQU8saUJBQWlCLFlBQVksbUJBQUssaUJBQWdCO0FBQzNEO0FBRUEsZ0JBQVcsU0FBQyxPQUFlO0FBQ3pCLHFCQUFLLGtCQUFtQixNQUFNLFlBQVk7QUFFMUMsTUFBSSxDQUFDLG1CQUFLLGVBQWU7QUFFekIsYUFBVyxTQUFTLG1CQUFLLGVBQWMsVUFBVTtBQUMvQyxVQUFNLE9BQU8sTUFBTSxhQUFhLFlBQVksS0FBSztBQUNqRCxVQUFNLFlBQVksQ0FBQyxtQkFBSyxxQkFBb0IsS0FBSyxTQUFTLG1CQUFLLGlCQUFnQjtBQUUvRSxVQUFNLFVBQVUsc0JBQUssbURBQUwsV0FBMEI7QUFDMUMsVUFBTSxhQUFhLG1CQUFLLGtCQUFpQixJQUFJLE9BQU87QUFFcEQsSUFBQyxNQUFzQixTQUFTLEVBQUUsYUFBYTtBQUFBLEVBQ2pEO0FBQ0Y7QUFFQSx5QkFBb0IsU0FBQyxPQUF3QjtBQUMzQyxhQUFXLE9BQU8sTUFBTSxXQUFXO0FBQ2pDLFFBQUksQ0FBQyxRQUFRLFdBQVcsU0FBUyxPQUFPLEVBQUUsU0FBUyxHQUFHLEdBQUc7QUFDdkQsYUFBTyxRQUFRLFlBQVksU0FBUztBQUFBLElBQ3RDO0FBQUEsRUFDRjtBQUNBLFNBQU87QUFDVDtBQUVBLHdCQUFtQixXQUFHO0FBQ3BCLE1BQUk7QUFDRixVQUFNLFFBQVEsYUFBYSxRQUFRLHVCQUF1QjtBQUMxRCxRQUFJLE9BQU87QUFDVCx5QkFBSyxrQkFBbUIsSUFBSSxJQUFJLEtBQUssTUFBTSxLQUFLLENBQUM7QUFBQSxJQUNuRDtBQUFBLEVBQ0YsU0FBUyxHQUFHO0FBQ1YsWUFBUSxNQUFNLHdFQUF3RTtBQUFBLEVBQ3hGO0FBQ0Esd0JBQUssa0RBQUw7QUFDRjtBQUVBLHdCQUFtQixXQUFHO0FBQ3BCLE1BQUksQ0FBQyxtQkFBSyxtQkFBbUI7QUFDN0IsUUFBTSxZQUFZLG1CQUFLLG1CQUFrQixpQkFBaUIsaUJBQWlCO0FBQzNFLFlBQVUsUUFBUSxVQUFRO0FBQ3hCLFVBQU0sUUFBUyxLQUFhO0FBQzVCLElBQUMsS0FBYSxVQUFVLG1CQUFLLGtCQUFpQixJQUFJLEtBQUs7QUFBQSxFQUN6RCxDQUFDO0FBQ0g7QUFFQSx3QkFBbUIsV0FBRztBQUNwQixNQUFJO0FBQ0YsaUJBQWE7QUFBQSxNQUFRO0FBQUEsTUFDbkIsS0FBSyxVQUFVLENBQUMsR0FBRyxtQkFBSyxpQkFBZ0IsQ0FBQztBQUFBLElBQUM7QUFBQSxFQUM5QyxTQUFTLEdBQUc7QUFBQSxFQUVaO0FBQ0Y7QUFFQTtBQWFNLGNBQVMsaUJBQUc7QUFDaEIsTUFBSSxDQUFDLG1CQUFLLGVBQWU7QUFFekIsUUFBTSxPQUFPLE1BQU0sS0FBSyxtQkFBSyxlQUFjLFFBQVEsRUFDaEQsT0FBTyxXQUFTLENBQUUsTUFBc0IsTUFBTSxFQUM5QyxJQUFJLFdBQVM7QUFDWixVQUFNLE9BQU8sTUFBTSxjQUFjLHNCQUFzQixHQUFHLGFBQWEsS0FBSyxLQUFLO0FBQ2pGLFVBQU0sT0FBTyxNQUFNLGNBQWMscUJBQXFCLEdBQUcsYUFBYSxLQUFLLEtBQUs7QUFDaEYsVUFBTSxVQUFVLE1BQU0sY0FBYyx3QkFBd0IsR0FBRyxhQUFhLEtBQUssS0FBSztBQUN0RixXQUFPLElBQUksSUFBSSxLQUFLLElBQUksSUFBSSxPQUFPO0FBQUEsRUFDckMsQ0FBQyxFQUFFLEtBQUssSUFBSTtBQUVkLE1BQUksQ0FBQyxLQUFNO0FBRVgsTUFBSTtBQUNGLFVBQU0sVUFBVSxVQUFVLFVBQVUsSUFBSTtBQUN4QyxVQUFNLE1BQU0sc0JBQUssaUNBQUwsV0FBUTtBQUNwQixRQUFJLEtBQUs7QUFDUCxZQUFNLFdBQVcsTUFBTSxLQUFLLElBQUksVUFBVSxFQUFFO0FBQUEsUUFDMUMsT0FBSyxFQUFFLGFBQWEsS0FBSyxjQUFjLEVBQUUsYUFBYSxLQUFLLEVBQUUsVUFBVSxLQUFLO0FBQUEsTUFDOUU7QUFDQSxVQUFJLFVBQVU7QUFDWixjQUFNLFdBQVcsU0FBUztBQUMxQixpQkFBUyxjQUFjO0FBRXZCLFlBQUksbUJBQUssMkJBQTBCO0FBQ2pDLHVCQUFhLG1CQUFLLHlCQUF3QjtBQUFBLFFBQzVDO0FBRUEsMkJBQUssMEJBQTJCLFdBQVcsTUFBTTtBQUMvQyxjQUFJLEtBQUssZUFBZSxTQUFTLFlBQVk7QUFDM0MscUJBQVMsY0FBYztBQUFBLFVBQ3pCO0FBQ0EsNkJBQUssMEJBQTJCO0FBQUEsUUFDbEMsR0FBRyxHQUFJO0FBQUEsTUFDVDtBQUFBLElBQ0Y7QUFBQSxFQUNGLFNBQVMsS0FBSztBQUNaLFlBQVEsTUFBTSwyQ0FBMkMsR0FBRztBQUFBLEVBQzlEO0FBQ0Y7QUFFQSx1QkFBa0IsV0FBRztBQUNuQixRQUFNLGNBQWMsc0JBQUssaUNBQUwsV0FBUTtBQUM1QixRQUFNLGFBQWEsc0JBQUssaUNBQUwsV0FBUTtBQUMzQixRQUFNLGFBQWEsS0FBSyxZQUFZLGNBQWMsY0FBYztBQUNoRSxRQUFNLFlBQVksS0FBSyxZQUFZLGNBQWMsYUFBYTtBQUU5RCxNQUFJLGVBQWUsWUFBWTtBQUM3QixnQkFBWSxpQkFBaUIsU0FBUyxNQUFNO0FBQzFDLDRCQUFLLDhDQUFMO0FBQ0EsTUFBQyxXQUFtQixVQUFVO0FBQUEsSUFDaEMsQ0FBQztBQUVELGdCQUFZLGlCQUFpQixTQUFTLE1BQU8sV0FBbUIsTUFBTSxDQUFDO0FBRXZFLGVBQVcsaUJBQWlCLFNBQVMsTUFBTTtBQUN6Qyw0QkFBSyw2Q0FBTDtBQUFBLElBQ0YsQ0FBQztBQUFBLEVBQ0g7QUFDRjtBQUVBLHVCQUFrQixXQUFHO0FBQ25CLFFBQU0sU0FBUyxLQUFLLFlBQVksY0FBYyxZQUFZO0FBQzFELFFBQU0sT0FBTyxLQUFLLFlBQVksY0FBYyxZQUFZO0FBRXhELE1BQUksQ0FBQyxVQUFVLENBQUMsS0FBTTtBQUV0QixxQkFBSyxhQUFlLE9BQWU7QUFFbkMsU0FBTyxpQkFBaUIsVUFBVSxDQUFDLE1BQWE7QUFDOUMsdUJBQUssYUFBZSxFQUFVO0FBRTlCLHFCQUFpQixZQUFZO0FBQUEsTUFDM0IsUUFBUSxFQUFFLE1BQU8sRUFBVSxLQUFLO0FBQUEsSUFDbEMsQ0FBQztBQUVELFFBQUssRUFBVSxNQUFNO0FBQ25CLDRCQUFLLGtEQUFMO0FBQUEsSUFDRjtBQUFBLEVBQ0YsQ0FBQztBQUVELFNBQU8saUJBQWlCLFVBQVUsQ0FBQyxNQUFhO0FBQzlDLElBQUMsT0FBZSxhQUFhLGlCQUFrQixFQUFVLE1BQU07QUFFL0QscUJBQWlCLFlBQVk7QUFBQSxNQUMzQixRQUFRLEVBQUUsUUFBUyxFQUFVLE9BQU87QUFBQSxJQUN0QyxDQUFDO0FBQUEsRUFDSCxDQUFDO0FBRUQsT0FBSyxpQkFBaUIsVUFBVSxDQUFDLE1BQWE7QUFDNUMscUJBQWlCLFlBQVk7QUFBQSxNQUMzQixNQUFNLEVBQUUsZUFBZ0IsRUFBVSxjQUFjO0FBQUEsSUFDbEQsQ0FBQztBQUVELFFBQUssRUFBVSxrQkFBa0IsS0FBTSxPQUFlLE1BQU07QUFDMUQsNEJBQUssa0RBQUw7QUFBQSxJQUNGO0FBRUEsUUFBSyxFQUFVLGtCQUFrQixLQUFNLE9BQWUsTUFBTTtBQUMxRCw0QkFBSyxvREFBTDtBQUFBLElBQ0Y7QUFBQSxFQUNGLENBQUM7QUFDSDtBQUVBLG1CQUFjLFdBQVc7QUFDdkIsUUFBTSxLQUFLLFVBQVU7QUFDckIsTUFBSSxHQUFHLFNBQVMsVUFBVSxHQUFHO0FBQzNCLFVBQU0sUUFBUSxHQUFHLE1BQU0sZ0JBQWdCO0FBQ3ZDLFdBQU8sUUFBUSxXQUFXLE1BQU0sQ0FBQyxDQUFDLEtBQUs7QUFBQSxFQUN6QyxXQUFXLEdBQUcsU0FBUyxNQUFNLEdBQUc7QUFDOUIsVUFBTSxRQUFRLEdBQUcsTUFBTSxZQUFZO0FBQ25DLFdBQU8sUUFBUSxRQUFRLE1BQU0sQ0FBQyxDQUFDLEtBQUs7QUFBQSxFQUN0QyxXQUFXLEdBQUcsU0FBUyxTQUFTLEdBQUc7QUFDakMsVUFBTSxRQUFRLEdBQUcsTUFBTSxlQUFlO0FBQ3RDLFdBQU8sUUFBUSxVQUFVLE1BQU0sQ0FBQyxDQUFDLEtBQUs7QUFBQSxFQUN4QyxXQUFXLEdBQUcsU0FBUyxTQUFTLEtBQUssQ0FBQyxHQUFHLFNBQVMsUUFBUSxHQUFHO0FBQzNELFVBQU0sUUFBUSxHQUFHLE1BQU0sZ0JBQWdCO0FBQ3ZDLFdBQU8sUUFBUSxVQUFVLE1BQU0sQ0FBQyxDQUFDLEtBQUs7QUFBQSxFQUN4QztBQUNBLFNBQU87QUFDVDtBQUVNLG1CQUFjLGlCQUFHO0FBQ3JCLFFBQU0sT0FBTyxNQUFNLEtBQUssc0JBQUssa0NBQUwsV0FBUyxxQkFBcUIsRUFBRSxJQUFJLFFBQU07QUFDaEUsVUFBTSxLQUFLLEdBQUc7QUFDZCxRQUFJLE1BQU0sR0FBRyxZQUFZLE1BQU07QUFDN0IsYUFBTyxHQUFHLEdBQUcsV0FBVyxLQUFLLEdBQUcsV0FBVztBQUFBLElBQzdDO0FBQ0EsV0FBTztBQUFBLEVBQ1QsQ0FBQyxFQUFFLEtBQUssSUFBSTtBQUVaLE1BQUksbUJBQW1CO0FBQ3ZCLE1BQUksbUJBQUssYUFBWSxlQUFlO0FBQ2xDLHVCQUFtQjtBQUFBLEVBQUssSUFBSSxPQUFPLEVBQUUsQ0FBQztBQUFBO0FBQUEsRUFBa0IsSUFBSSxPQUFPLEVBQUUsQ0FBQztBQUFBLEVBQUssbUJBQUssWUFBVyxhQUFhO0FBQUE7QUFBQSxFQUMxRztBQUVBLFFBQU0sWUFBWTtBQUFBLEVBQ3BCLElBQUksT0FBTyxFQUFFLENBQUM7QUFBQSxFQUNkLElBQUksR0FBRyxnQkFBZ0I7QUFBQSxFQUN2QixJQUFJLE9BQU8sRUFBRSxDQUFDO0FBQUEsY0FDSCxvQkFBSSxLQUFLLEdBQUUsWUFBWSxDQUFDO0FBRWpDLE1BQUk7QUFDRixVQUFNLFVBQVUsVUFBVSxVQUFVLFNBQVM7QUFDN0MsVUFBTSxhQUFhLEtBQUssWUFBWSxjQUFjLGFBQWE7QUFDL0QsUUFBSSxZQUFZO0FBQ2QsWUFBTSxlQUFlLFdBQVc7QUFDaEMsaUJBQVcsY0FBYztBQUV6QixVQUFJLG1CQUFLLDRCQUEyQjtBQUNsQyxxQkFBYSxtQkFBSywwQkFBeUI7QUFBQSxNQUM3QztBQUVBLHlCQUFLLDJCQUE0QixXQUFXLE1BQU07QUFDaEQsWUFBSSxLQUFLLGVBQWUsV0FBVyxZQUFZO0FBQzdDLHFCQUFXLGNBQWM7QUFBQSxRQUMzQjtBQUNBLDJCQUFLLDJCQUE0QjtBQUFBLE1BQ25DLEdBQUcsR0FBSTtBQUFBLElBQ1Q7QUFBQSxFQUNGLFNBQVMsS0FBSztBQUNaLFlBQVEsTUFBTSxpREFBaUQsR0FBRztBQUFBLEVBQ3BFO0FBQ0Y7QUFFQSxnQkFBVyxTQUFDLE1BQThEO0FBQ3hFLE1BQUksQ0FBQyxtQkFBSyxlQUFlO0FBRXpCLFFBQU0sY0FBYyxLQUFLLElBQUksU0FBTztBQUNsQyxVQUFNLFdBQVcsOEJBQWUsbUJBQWtCLFFBQVEsVUFBVSxJQUFJO0FBRXhFLFVBQU0sT0FBTyxJQUFJLEtBQUssSUFBSSxJQUFJO0FBQzlCLFVBQU0sT0FBTyxLQUFLLG1CQUFtQjtBQUVyQyxVQUFNLFlBQVksU0FBUyxjQUFjLDBCQUEwQjtBQUNuRSxjQUFVLFVBQVUsSUFBSSxJQUFJLElBQUk7QUFDaEMsY0FBVSxhQUFhLGVBQWUsSUFBSSxJQUFJO0FBRTlDLFVBQU0sWUFBWSxzQkFBSywyQ0FBTCxXQUFrQixJQUFJO0FBQ3hDLFVBQU0sYUFBYSxHQUFHLFNBQVMsSUFBSSxJQUFJLElBQUksSUFBSSxPQUFPLEdBQUcsWUFBWTtBQUNyRSxVQUFNLFlBQVksQ0FBQyxtQkFBSyxxQkFBb0IsV0FBVyxTQUFTLG1CQUFLLGlCQUFnQjtBQUVyRixVQUFNLG1CQUFtQixJQUFJLFNBQVMsWUFBWSxTQUFTLElBQUk7QUFDL0QsVUFBTSxhQUFhLG1CQUFLLGtCQUFpQixJQUFJLGdCQUFnQjtBQUU3RCxRQUFJLEVBQUUsYUFBYSxhQUFhO0FBQzlCLGdCQUFVLGFBQWEsVUFBVSxFQUFFO0FBQUEsSUFDckM7QUFFQSxVQUFNLFFBQVEsU0FBUyxjQUFjLHNCQUFzQjtBQUMzRCxVQUFNLGNBQWMsc0JBQUssMkNBQUwsV0FBa0IsSUFBSTtBQUMxQywwQkFBSyxrREFBTCxXQUF5QixPQUFPLElBQUk7QUFFcEMsVUFBTSxTQUFTLFNBQVMsY0FBYyxxQkFBcUI7QUFDM0QsV0FBTyxhQUFhLFlBQVksSUFBSSxJQUFJO0FBQ3hDLFdBQU8sY0FBYztBQUVyQixJQUFDLFNBQVMsY0FBYyx3QkFBd0IsRUFBa0IsY0FBYyxJQUFJO0FBRXBGLFdBQU87QUFBQSxFQUNULENBQUM7QUFFRCxNQUFJLENBQUMsbUJBQUssc0JBQXFCO0FBQzdCLHVCQUFLLGVBQWMsZ0JBQWdCLEdBQUcsV0FBVztBQUNqRCx1QkFBSyxxQkFBc0I7QUFFM0IsUUFBSSxtQkFBSyxjQUFhO0FBQ3BCLDRCQUFLLG9EQUFMO0FBQUEsSUFDRjtBQUFBLEVBQ0YsT0FBTztBQUNMLHVCQUFLLGVBQWMsT0FBTyxHQUFHLFdBQVc7QUFFeEMsUUFBSSxtQkFBSyxjQUFhO0FBQ3BCLDRCQUFLLG9EQUFMO0FBQUEsSUFDRjtBQUFBLEVBQ0Y7QUFDRjtBQUVBLGlCQUFZLFNBQUMsTUFBc0I7QUFDakMsVUFBUSxNQUFNO0FBQUEsSUFDWixLQUFLO0FBQVEsYUFBTztBQUFBLElBQ3BCLEtBQUs7QUFBVyxhQUFPO0FBQUEsSUFDdkIsS0FBSztBQUFTLGFBQU87QUFBQSxJQUNyQixLQUFLO0FBQVMsYUFBTztBQUFBLElBQ3JCO0FBQVMsYUFBTyxLQUFLLFlBQVk7QUFBQSxFQUNuQztBQUNGO0FBRUEsd0JBQW1CLFNBQUMsT0FBb0IsTUFBYztBQUNwRCxVQUFRLE1BQU07QUFBQSxJQUNaLEtBQUs7QUFDSCxhQUFPLE1BQU0sYUFBYSxVQUFVLE1BQU07QUFBQSxJQUM1QyxLQUFLO0FBQ0gsYUFBTyxNQUFNLGFBQWEsVUFBVSxTQUFTO0FBQUEsSUFDL0MsS0FBSztBQUNILGFBQU8sTUFBTSxhQUFhLFVBQVUsUUFBUTtBQUFBLElBQzlDLEtBQUs7QUFDSCxhQUFPLE1BQU0sYUFBYSxTQUFTLFFBQVE7QUFBQSxJQUM3QztBQUNFLFlBQU0sYUFBYSxTQUFTLE1BQU07QUFBQSxFQUN0QztBQUNGO0FBRUEsMEJBQXFCLFdBQUc7QUFDdEIsTUFBSSxDQUFDLG1CQUFLLGVBQWU7QUFFekIsd0JBQXNCLE1BQU07QUFDMUIsVUFBTSxVQUFVLG1CQUFLLGVBQWU7QUFDcEMsUUFBSSxTQUFTO0FBQ1gsY0FBUSxlQUFlLEVBQUUsVUFBVSxRQUFRLE9BQU8sTUFBTSxDQUFDO0FBQUEsSUFDM0Q7QUFBQSxFQUNGLENBQUM7QUFDSDtBQUVBLHdCQUFtQixXQUFHO0FBQ3BCLE1BQUksQ0FBQyxtQkFBSyxlQUFlO0FBRXpCLE1BQUksbUJBQUssaUJBQWdCO0FBQ3ZCLDBCQUFLLG9EQUFMO0FBQUEsRUFDRixPQUFPO0FBQ0wsZUFBVyxNQUFNO0FBQ2YsNEJBQUssb0RBQUw7QUFBQSxJQUNGLEdBQUcsR0FBRztBQUFBLEVBQ1I7QUFDRjtBQUVBLHFDQUFnQyxXQUFHO0FBQ2pDLE1BQUk7QUFDRixVQUFNLGtCQUNKLGFBQWEsUUFBUSx3QkFBd0IsTUFBTSxRQUNuRCxhQUFhLFFBQVEsdUJBQXVCLE1BQU0sUUFDbEQsYUFBYSxRQUFRLHlCQUF5QixNQUFNLFFBQ3BELGFBQWEsUUFBUSxzQkFBc0IsTUFBTTtBQUVuRCxRQUFJLGlCQUFpQjtBQUNuQixZQUFNLFdBQVcsYUFBYSxRQUFRLCtCQUErQjtBQUNyRSxVQUFJLENBQUMsVUFBVTtBQUNiLHlCQUFpQix3QkFBd0I7QUFDekMscUJBQWEsUUFBUSxpQ0FBaUMsTUFBTTtBQUM1RCxtQkFBVyxNQUFNLE9BQU8sU0FBUyxPQUFPLEdBQUcsR0FBRztBQUFBLE1BQ2hEO0FBQUEsSUFDRjtBQUFBLEVBQ0YsU0FBUyxHQUFHO0FBQUEsRUFFWjtBQUNGO0FBRUEsNEJBQXVCLFdBQUc7QUFDeEIsUUFBTSxjQUFjLEtBQUssWUFBWSxjQUFjLHNCQUFzQjtBQUN6RSxNQUFJLENBQUMsWUFBYTtBQUVsQixRQUFNLFFBQVEsaUJBQWlCLFNBQVM7QUFFeEMsd0JBQUssZ0RBQUwsV0FBdUIsTUFBTTtBQUU3QixRQUFNLFFBQVEsWUFBWSxpQkFBaUIseUJBQXlCO0FBQ3BFLFFBQU0sUUFBUSxVQUFRO0FBQ3BCLFFBQUssS0FBYSxVQUFVLE1BQU0sYUFBYTtBQUM3QyxXQUFLLGFBQWEsWUFBWSxFQUFFO0FBQUEsSUFDbEM7QUFBQSxFQUNGLENBQUM7QUFFRCxjQUFZLGlCQUFpQiw2QkFBNkIsQ0FBQyxNQUFhO0FBQ3RFLFVBQU0sU0FBVSxFQUFVO0FBQzFCLDBCQUFLLGdEQUFMLFdBQXVCO0FBQ3ZCLHFCQUFpQixZQUFZLEVBQUUsYUFBYSxPQUFPLENBQUM7QUFBQSxFQUN0RCxDQUFDO0FBQ0g7QUFFQSxzQkFBaUIsU0FBQyxRQUFnQjtBQUNoQyxRQUFNLE9BQU8sU0FBUztBQUV0QixVQUFRLFFBQVE7QUFBQSxJQUNkLEtBQUs7QUFDSCxXQUFLLE1BQU0sY0FBYztBQUN6QjtBQUFBLElBQ0YsS0FBSztBQUNILFdBQUssTUFBTSxjQUFjO0FBQ3pCO0FBQUEsSUFDRixLQUFLO0FBQUEsSUFDTDtBQUNFLFdBQUssTUFBTSxjQUFjO0FBQ3pCO0FBQUEsRUFDSjtBQUNGO0FBRUEsMkJBQXNCLFdBQUc7QUFDdkIsT0FBSyxpQkFBaUIseUJBQXlCLG1CQUFLLGNBQWE7QUFDakUsT0FBSyxpQkFBaUIsd0JBQXdCLG1CQUFLLGNBQWE7QUFDaEUsT0FBSyxpQkFBaUIsNEJBQTRCLG1CQUFLLGNBQWE7QUFDcEUsT0FBSyxpQkFBaUIsd0JBQXdCLG1CQUFLLGFBQVk7QUFDL0QsT0FBSyxpQkFBaUIsdUJBQXVCLG1CQUFLLGFBQVk7QUFDOUQsT0FBSyxpQkFBaUIsMkJBQTJCLG1CQUFLLGFBQVk7QUFDcEU7QUFFQTtBQWdDQTtBQWlDQSxtQkFBYyxTQUFDLE9BQWlFO0FBQzlFLFFBQU0saUJBQWlCLEtBQUssa0JBQWtCO0FBRTlDLE1BQUksTUFBTSxjQUFjO0FBQ3RCLGVBQVcsV0FBVyxNQUFNLGFBQWEsR0FBRztBQUMxQyxVQUFJLEVBQUUsbUJBQW1CLFNBQVU7QUFFbkMsVUFBSyxRQUF3QixTQUFTLGtCQUFrQixRQUFRO0FBQzlELGNBQU0sVUFBVyxRQUF3QixRQUFRLFdBQVc7QUFDNUQsWUFBSSxnQkFBZ0IsT0FBTyxTQUFVLFFBQXdCLFFBQVEsaUJBQWlCLElBQUksRUFBRTtBQUM1RixZQUFJLE9BQU8sTUFBTSxhQUFhLEVBQUcsaUJBQWdCO0FBQ2pELGVBQU8sRUFBRSxTQUFTLGNBQWM7QUFBQSxNQUNsQztBQUFBLElBQ0Y7QUFBQSxFQUNGO0FBRUEsU0FBTyxFQUFFLFNBQVMsZ0JBQWdCLGVBQWUsRUFBRTtBQUNyRDtBQUVBLDBCQUFxQixTQUFDLE9BQXNCO0FBQzFDLFVBQVEsTUFBTSxNQUFNO0FBQUEsSUFDbEIsS0FBSztBQUNILGFBQU87QUFBQSxJQUNULEtBQUs7QUFDSCxhQUFPO0FBQUEsSUFDVCxLQUFLO0FBQ0gsYUFBTztBQUFBLElBQ1Q7QUFDRSxhQUFPO0FBQUEsRUFDWDtBQUNGO0FBRUEsK0JBQTBCLFNBQUMsT0FBc0I7QUFDL0MsVUFBUSxNQUFNLE1BQU07QUFBQSxJQUNsQixLQUFLO0FBQ0gsYUFBTztBQUFBLElBQ1QsS0FBSztBQUNILGFBQU87QUFBQSxJQUNULEtBQUs7QUFDSCxhQUFPO0FBQUEsSUFDVDtBQUNFLGFBQU87QUFBQSxFQUNYO0FBQ0Y7QUFFQSwrQkFBMEIsV0FBRztBQUMzQixxQkFBSyxtQkFBb0IsQ0FBQyxNQUFhO0FBQ3JDLFFBQUssRUFBRSxRQUFvQixZQUFZLGtCQUFtQjtBQUUxRCxVQUFNLFNBQVMsc0JBQUssNkNBQUwsV0FBb0IsRUFBRTtBQUNyQyxVQUFNLFlBQVksaUJBQWlCLGFBQWE7QUFDaEQsUUFBSSxDQUFDLFVBQVUsU0FBUyxTQUFTLE1BQU0sR0FBRztBQUN4QyxnQkFBVSxTQUFTLEtBQUssTUFBTTtBQUM5Qix1QkFBaUIsYUFBYSxTQUFTO0FBQUEsSUFDekM7QUFBQSxFQUNGO0FBQ0EsT0FBSyxpQkFBaUIsVUFBVSxtQkFBSyxrQkFBaUI7QUFFdEQscUJBQUsscUJBQXNCLENBQUMsTUFBYTtBQUN2QyxRQUFLLEVBQUUsUUFBb0IsWUFBWSxrQkFBbUI7QUFFMUQsVUFBTSxTQUFTLHNCQUFLLDZDQUFMLFdBQW9CLEVBQUU7QUFDckMsVUFBTSxZQUFZLGlCQUFpQixhQUFhO0FBQ2hELFVBQU0sUUFBUSxVQUFVLFNBQVMsUUFBUSxNQUFNO0FBQy9DLFFBQUksUUFBUSxJQUFJO0FBQ2QsZ0JBQVUsU0FBUyxPQUFPLE9BQU8sQ0FBQztBQUNsQyx1QkFBaUIsYUFBYSxTQUFTO0FBQUEsSUFDekM7QUFBQSxFQUNGO0FBQ0EsT0FBSyxpQkFBaUIsWUFBWSxtQkFBSyxvQkFBbUI7QUFFMUQscUJBQUssbUJBQW9CLENBQUMsTUFBYTtBQUNyQyxRQUFLLEVBQUUsUUFBb0IsWUFBWSxrQkFBbUI7QUFFMUQsVUFBTSxTQUFTLHNCQUFLLDZDQUFMLFdBQW9CLEVBQUU7QUFDckMscUJBQWlCLGdCQUFnQixFQUFFLFVBQVUsT0FBTyxDQUFDO0FBQUEsRUFDdkQ7QUFDQSxPQUFLLGlCQUFpQixVQUFVLG1CQUFLLGtCQUFpQjtBQUV0RCx3QkFBSyw4Q0FBTDtBQUNGO0FBRUEsb0JBQWUsV0FBRztBQUNoQixRQUFNLFlBQVksaUJBQWlCLGFBQWE7QUFFaEQsYUFBVyxVQUFVLFVBQVUsVUFBVTtBQUN2QyxVQUFNLFdBQVcsc0JBQUssZ0RBQUwsV0FBdUI7QUFDeEMsUUFBSSxZQUFZLENBQUMsU0FBUyxhQUFhLFVBQVUsR0FBRztBQUNsRCxlQUFTLGFBQWEsWUFBWSxFQUFFO0FBQUEsSUFDdEM7QUFBQSxFQUNGO0FBRUEsTUFBSSxVQUFVLFVBQVU7QUFDdEIsVUFBTSxXQUFXLHNCQUFLLGdEQUFMLFdBQXVCLFVBQVU7QUFDbEQsUUFBSSxZQUFZLENBQUMsU0FBUyxhQUFhLFNBQVMsR0FBRztBQUNqRCxlQUFTLGFBQWEsV0FBVyxFQUFFO0FBQUEsSUFDckM7QUFBQSxFQUNGO0FBQ0Y7QUFFQSxrQ0FBNkIsV0FBRztBQUM5QixRQUFNLE9BQU8sS0FBSyxZQUFZLGNBQWMsWUFBWTtBQUV4RCxNQUFJLENBQUMsS0FBTTtBQUVYLE9BQUssaUJBQWlCLGtCQUFrQixDQUFDLFVBQWlCO0FBQ3hELFVBQU0sWUFBWSxDQUFFLE1BQWM7QUFFbEMscUJBQWlCLFlBQVk7QUFBQSxNQUMzQixTQUFTLEVBQUUsVUFBVTtBQUFBLElBQ3ZCLENBQUM7QUFBQSxFQUNILENBQUM7QUFDSDtBQUVBLHNCQUFpQixTQUFDLFFBQWdDO0FBQ2hELFFBQU0sUUFBUSxPQUFPLE1BQU0sR0FBRztBQUM5QixRQUFNLENBQUMsTUFBTSxZQUFZLFNBQVMsSUFBSSxJQUFJO0FBRTFDLE1BQUksYUFBYTtBQUNqQixNQUFJLFNBQVM7QUFDWCxrQkFBYyxtQkFBbUIsSUFBSSxPQUFPLE9BQU8sQ0FBQztBQUFBLEVBQ3REO0FBQ0EsTUFBSSxNQUFNO0FBQ1Isa0JBQWMsZUFBZSxJQUFJLE9BQU8sSUFBSSxDQUFDO0FBQUEsRUFDL0M7QUFFQSxNQUFJLFdBQVcsOEJBQThCLElBQUksT0FBTyxJQUFJLENBQUM7QUFDN0QsTUFBSSxZQUFZO0FBQ2QsVUFBTSxvQkFBb0IsSUFBSSxPQUFPLFVBQVU7QUFDL0MsVUFBTSxjQUFjLElBQUksT0FBTyxJQUFJO0FBQ25DLFVBQU0sWUFBWSw4QkFBOEIsV0FBVyx3QkFBd0IsaUJBQWlCLEtBQUssVUFBVTtBQUNuSCxVQUFNLFlBQVksOEJBQThCLFdBQVcsaUJBQWlCLGlCQUFpQixLQUFLLFVBQVU7QUFDNUcsZUFBVyxHQUFHLFNBQVMsS0FBSyxTQUFTO0FBQUEsRUFDdkMsT0FBTztBQUNMLGdCQUFZO0FBQUEsRUFDZDtBQUVBLFNBQU8sS0FBSyxjQUFjLFFBQVE7QUFDcEM7QUFFQSxtQkFBYyxTQUFDLFVBQTJCO0FBQ3hDLFFBQU0sT0FBTyxTQUFTLGFBQWEsV0FBVztBQUM5QyxRQUFNLGFBQWEsU0FBUyxhQUFhLGtCQUFrQixLQUFLLFNBQVMsYUFBYSxXQUFXO0FBQ2pHLFFBQU0sVUFBVSxTQUFTLGFBQWEsZUFBZTtBQUNyRCxRQUFNLE9BQU8sU0FBUyxhQUFhLFdBQVc7QUFDOUMsUUFBTSxXQUFXLFNBQVMsYUFBYSxlQUFlO0FBRXRELFFBQU0sUUFBUSxDQUFDLElBQUk7QUFDbkIsTUFBSSxXQUFZLE9BQU0sS0FBSyxVQUFVO0FBQ3JDLE1BQUksUUFBUyxPQUFNLEtBQUssT0FBTztBQUMvQixNQUFJLFVBQVU7QUFDWixVQUFNLEtBQUssUUFBUTtBQUFBLEVBQ3JCLFdBQVcsTUFBTTtBQUNmLFVBQU0sS0FBSyxJQUFJO0FBQUEsRUFDakI7QUFFQSxTQUFPLE1BQU0sS0FBSyxHQUFHO0FBQ3ZCO0FBSU0sMkJBQXNCLGlCQUFvQztBQUM5RCxNQUFJO0FBQ0YsVUFBTSxXQUFXLE1BQU0sTUFBTSx1QkFBdUI7QUFDcEQsUUFBSSxDQUFDLFNBQVMsSUFBSTtBQUNoQixjQUFRLEtBQUssOERBQThEO0FBQzNFLGFBQU8sb0JBQUksSUFBSTtBQUFBLElBQ2pCO0FBRUEsVUFBTSxXQUFXLE1BQU0sU0FBUyxLQUFLO0FBQ3JDLHVCQUFLLFdBQVk7QUFFakIsVUFBTSxXQUFXLG9CQUFJLElBQXVCO0FBRTVDLGVBQVcsVUFBVSxTQUFTLFdBQVcsQ0FBQyxHQUFHO0FBQzNDLGlCQUFXLGVBQWUsT0FBTyxnQkFBZ0IsQ0FBQyxHQUFHO0FBQ25ELFlBQUksWUFBWSxpQkFBaUIsWUFBWSxTQUFTO0FBQ3BELGdCQUFNLFVBQVUsWUFBWTtBQUM1QixnQkFBTSxTQUFTLFlBQVksVUFBVSxDQUFDO0FBRXRDLGNBQUksT0FBTyxTQUFTLEdBQUc7QUFDckIsa0JBQU0sYUFBYSxJQUFJLElBQUksT0FBTyxJQUFJLE9BQUssRUFBRSxJQUFJLENBQUM7QUFDbEQscUJBQVMsSUFBSSxTQUFTO0FBQUEsY0FDcEI7QUFBQSxjQUNBO0FBQUEsWUFDRixDQUFDO0FBQUEsVUFDSDtBQUFBLFFBQ0Y7QUFBQSxNQUNGO0FBQUEsSUFDRjtBQUVBLFdBQU87QUFBQSxFQUNULFNBQVMsT0FBTztBQUNkLFlBQVEsS0FBSyxrRUFBa0UsS0FBSztBQUNwRixXQUFPLG9CQUFJLElBQUk7QUFBQSxFQUNqQjtBQUNGO0FBRU0sdUJBQWtCLGlCQUFHO0FBQ3pCLHFCQUFLLGtCQUFtQixNQUFNLHNCQUFLLHFEQUFMO0FBRTlCLE1BQUksbUJBQUssa0JBQWlCLFNBQVMsRUFBRztBQUV0Qyx3QkFBSyxvREFBTDtBQUNBLHdCQUFLLGtEQUFMO0FBQ0Esd0JBQUssb0RBQUw7QUFDRjtBQUVBLDBCQUFxQixXQUFHO0FBQ3RCLFFBQU0sT0FBTyxLQUFLO0FBQ2xCLE1BQUksQ0FBQyxLQUFNO0FBRVgsUUFBTSxPQUFPLEtBQUssY0FBYztBQUVoQyxhQUFXLENBQUMsU0FBUyxTQUFTLEtBQUssbUJBQUssbUJBQW1CO0FBQ3pELFVBQU0sV0FBVyxLQUFLLGlCQUFpQixPQUFPO0FBRTlDLGVBQVcsV0FBVyxVQUFVO0FBQzlCLGlCQUFXLGFBQWEsVUFBVSxZQUFZO0FBQzVDLGdCQUFRLGlCQUFpQixXQUFXLG1CQUFLLHNCQUFxQixFQUFFLFNBQVMsS0FBSyxDQUFDO0FBQUEsTUFDakY7QUFDQSxNQUFDLFFBQXdCLFFBQVEsb0JBQW9CO0FBQ3JELHlCQUFLLHFCQUFvQixJQUFJLE9BQU87QUFBQSxJQUN0QztBQUFBLEVBQ0Y7QUFDRjtBQUVBLDBCQUFxQixXQUFHO0FBQ3RCLFFBQU0sT0FBTyxLQUFLO0FBQ2xCLE1BQUksQ0FBQyxLQUFNO0FBRVgsUUFBTSxPQUFPLEtBQUssY0FBYztBQUVoQyxxQkFBSyxXQUFVLFFBQVEsTUFBTTtBQUFBLElBQzNCLFdBQVc7QUFBQSxJQUNYLFNBQVM7QUFBQSxFQUNYLENBQUM7QUFDSDtBQUVBO0FBYUEsMkJBQXNCLFNBQUMsZUFBbUQ7QUFDeEUsTUFBSSxDQUFDLGVBQWU7QUFDbEIsV0FBTyxFQUFFLFNBQVMsTUFBTSxhQUFhLEtBQUs7QUFBQSxFQUM1QztBQUVBLE1BQUksVUFBVSxjQUFjLFdBQVc7QUFDdkMsTUFBSSxjQUFjLGNBQWMsZUFBZTtBQUUvQyxNQUFJLGNBQWMsTUFBTSxRQUFRLG1CQUFLLFlBQVc7QUFDOUMsVUFBTSxXQUFXLGNBQWMsS0FBSztBQUNwQyxVQUFNLGtCQUFrQixzQkFBSyxtREFBTCxXQUEwQjtBQUVsRCxRQUFJLGlCQUFpQjtBQUNuQixVQUFJLENBQUMsV0FBVyxnQkFBZ0IsU0FBUztBQUN2QyxrQkFBVSxnQkFBZ0I7QUFBQSxNQUM1QixXQUFXLGdCQUFnQixXQUFXLGdCQUFnQixZQUFZLFNBQVM7QUFDekUsa0JBQVUsVUFBVSxHQUFHLE9BQU87QUFBQTtBQUFBLE9BQVksUUFBUSxLQUFLLGdCQUFnQixPQUFPLEtBQUssZ0JBQWdCO0FBQUEsTUFDckc7QUFFQSxVQUFJLENBQUMsZUFBZSxnQkFBZ0IsYUFBYTtBQUMvQyxzQkFBYyxnQkFBZ0I7QUFBQSxNQUNoQyxXQUFXLGdCQUFnQixlQUFlLGdCQUFnQixnQkFBZ0IsYUFBYTtBQUNyRixzQkFBYyxjQUFjLEdBQUcsV0FBVztBQUFBO0FBQUEsRUFBTyxnQkFBZ0IsV0FBVyxLQUFLLGdCQUFnQjtBQUFBLE1BQ25HO0FBQUEsSUFDRjtBQUFBLEVBQ0Y7QUFFQSxTQUFPLEVBQUUsU0FBUyxZQUFZO0FBQ2hDO0FBRUEseUJBQW9CLFNBQUMsVUFBa0I7QUFDckMsTUFBSSxDQUFDLG1CQUFLLFdBQVcsUUFBTztBQUU1QixhQUFXLFVBQVUsbUJBQUssV0FBVSxXQUFXLENBQUMsR0FBRztBQUNqRCxlQUFXLGVBQWUsT0FBTyxnQkFBZ0IsQ0FBQyxHQUFHO0FBQ25ELFVBQUksWUFBWSxTQUFTLGFBQ3BCLFlBQVksU0FBUyxXQUFXLFlBQVksU0FBUyxjQUFjO0FBQ3RFLGVBQU87QUFBQSxNQUNUO0FBQUEsSUFDRjtBQUFBLEVBQ0Y7QUFFQSxTQUFPO0FBQ1Q7QUFFQSxrQkFBYSxTQUFDLE9BQWMsUUFBcUIsU0FBaUIsV0FBc0I7QUFDdEYsUUFBTSxnQkFBZ0IsVUFBVSxPQUFPLEtBQUssT0FBSyxFQUFFLFNBQVMsTUFBTSxJQUFJO0FBRXRFLFFBQU0sWUFBWSxzQkFBSyxxREFBTCxXQUE0QjtBQUU5QyxRQUFNLG1CQUFtQixzQkFBSyxzREFBTCxXQUE2QjtBQUV0RCxRQUFNLGNBQTJCO0FBQUEsSUFDL0IsSUFBSSxHQUFHLEtBQUssSUFBSSxDQUFDLElBQUksS0FBSyxPQUFPLENBQUM7QUFBQSxJQUNsQyxXQUFXLG9CQUFJLEtBQUs7QUFBQSxJQUNwQixXQUFXLE1BQU07QUFBQSxJQUNqQjtBQUFBLElBQ0EsV0FBVyxPQUFPLE1BQU07QUFBQSxJQUN4QixjQUFjLE9BQU8sYUFBYTtBQUFBLElBQ2xDO0FBQUEsSUFDQSxjQUFjLGVBQWUsTUFBTSxRQUFRO0FBQUEsSUFDM0MsU0FBUyxVQUFVO0FBQUEsSUFDbkIsYUFBYSxVQUFVO0FBQUEsSUFDdkIsU0FBUyxNQUFNO0FBQUEsSUFDZixVQUFVLE1BQU07QUFBQSxJQUNoQixZQUFZLE1BQU07QUFBQSxJQUNsQixrQkFBa0IsTUFBTTtBQUFBLEVBQzFCO0FBRUEscUJBQUssaUJBQWdCLEtBQUssV0FBVztBQUVyQyxNQUFJLG1CQUFLLGlCQUFnQixTQUFTLG1CQUFLLHFCQUFvQjtBQUN6RCx1QkFBSyxpQkFBZ0IsTUFBTTtBQUFBLEVBQzdCO0FBRUEsd0JBQUssMkNBQUwsV0FBa0I7QUFDcEI7QUFFQSw0QkFBdUIsU0FBQyxPQUF1QztBQUM3RCxRQUFNLGFBQXNDLENBQUM7QUFDN0MsUUFBTSxxQkFBcUIsSUFBSSxJQUFJLE9BQU8sb0JBQW9CLE1BQU0sU0FBUyxDQUFDO0FBRTlFLFFBQU0saUJBQWlCLENBQUMsVUFBNEI7QUFDbEQsUUFBSTtBQUNGLGFBQU8sS0FBSyxNQUFNLEtBQUssVUFBVSxLQUFLLENBQUM7QUFBQSxJQUN6QyxTQUFTLEdBQUc7QUFDVixVQUFJO0FBQ0YsZUFBTyxPQUFPLEtBQUs7QUFBQSxNQUNyQixTQUFTLFdBQVc7QUFDbEIsZUFBTztBQUFBLE1BQ1Q7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUVBLE1BQUksaUJBQWlCLGVBQWUsTUFBTSxXQUFXLFFBQVc7QUFDOUQsZUFBVyxTQUFTLGVBQWUsTUFBTSxNQUFNO0FBQUEsRUFDakQ7QUFFQSxhQUFXLE9BQU8sT0FBTyxvQkFBb0IsS0FBSyxHQUFHO0FBQ25ELFFBQUksQ0FBQyxtQkFBbUIsSUFBSSxHQUFHLEtBQUssQ0FBQyxJQUFJLFdBQVcsR0FBRyxLQUFLLENBQUMsV0FBVyxlQUFlLEdBQUcsR0FBRztBQUMzRixpQkFBVyxHQUFHLElBQUksZUFBZ0IsTUFBYyxHQUFHLENBQUM7QUFBQSxJQUN0RDtBQUFBLEVBQ0Y7QUFFQSxTQUFPO0FBQ1Q7QUFFQSxpQkFBWSxTQUFDLGFBQTBCO0FBQ3JDLE1BQUksQ0FBQyxtQkFBSyxZQUFZO0FBRXRCLFFBQU0sV0FBVyw4QkFBZSxxQkFBb0IsUUFBUSxVQUFVLElBQUk7QUFFMUUsUUFBTSxPQUFPLFlBQVksVUFBVSxtQkFBbUI7QUFFdEQsUUFBTSxZQUFZLFNBQVMsY0FBYywwQkFBMEI7QUFDbkUsWUFBVSxRQUFRLFVBQVUsWUFBWTtBQUN4QyxZQUFVLFFBQVEsWUFBWSxZQUFZO0FBQzFDLFlBQVUsUUFBUSxjQUFjLFlBQVk7QUFFNUMsUUFBTSxZQUFZLHNCQUFLLHNEQUFMLFdBQTZCO0FBQy9DLFFBQU0sWUFBWSxtQkFBSyxtQkFBa0IsU0FBUyxLQUFLLG1CQUFLLG1CQUFrQixJQUFJLFlBQVksU0FBUztBQUN2RyxRQUFNLGVBQWUsbUJBQUssaUJBQWdCLFNBQVMsS0FBSyxtQkFBSyxpQkFBZ0IsSUFBSSxZQUFZLE9BQU87QUFFcEcsTUFBSSxFQUFFLGFBQWEsYUFBYSxlQUFlO0FBQzdDLGNBQVUsYUFBYSxVQUFVLEVBQUU7QUFBQSxFQUNyQztBQUVBLFFBQU0sUUFBUSxTQUFTLGNBQWMsc0JBQXNCO0FBQzNELFFBQU0sY0FBYyxZQUFZO0FBQ2hDLFFBQU0sYUFBYSxVQUFVLE1BQU07QUFFbkMsUUFBTSxTQUFTLFNBQVMsY0FBYyxxQkFBcUI7QUFDM0QsU0FBTyxhQUFhLFlBQVksWUFBWSxVQUFVLFlBQVksQ0FBQztBQUNuRSxTQUFPLGNBQWM7QUFFckIsUUFBTSxZQUFZLFNBQVMsY0FBYyx3QkFBd0I7QUFDakUsTUFBSSxjQUFjLElBQUksWUFBWSxPQUFPO0FBQ3pDLE1BQUksWUFBWSxXQUFXO0FBQ3pCLG1CQUFlLElBQUksWUFBWSxTQUFTO0FBQUEsRUFDMUM7QUFDQSxZQUFVLGNBQWM7QUFFeEIscUJBQUssWUFBVyxPQUFPLFFBQVE7QUFFL0IsTUFBSSxDQUFDLG1CQUFLLG1CQUFrQjtBQUMxQiwwQkFBSywyQ0FBTCxXQUFrQixZQUFZO0FBQUEsRUFDaEM7QUFFQSxNQUFJLG1CQUFLLGdCQUFlLHNCQUFLLGlEQUFMLFlBQTJCO0FBQ2pELDBCQUFLLG9EQUFMO0FBQUEsRUFDRjtBQUNGO0FBRUEsaUJBQVksU0FBQyxTQUFpQjtBQUM1QixRQUFNLGNBQWMsc0JBQUssa0RBQUwsV0FBeUI7QUFDN0MsTUFBSSxDQUFDLFlBQWE7QUFFbEIscUJBQUssa0JBQW1CO0FBRXhCLFFBQU0sV0FBVyxtQkFBSyxhQUFZLGlCQUFpQixrQkFBa0I7QUFDckUsWUFBVSxRQUFRLFVBQVE7QUFDeEIsUUFBSyxLQUFxQixRQUFRLFlBQVksU0FBUztBQUNyRCxXQUFLLFVBQVUsSUFBSSxVQUFVO0FBQzdCLFdBQUssYUFBYSxpQkFBaUIsTUFBTTtBQUFBLElBQzNDLE9BQU87QUFDTCxXQUFLLFVBQVUsT0FBTyxVQUFVO0FBQ2hDLFdBQUssYUFBYSxpQkFBaUIsT0FBTztBQUFBLElBQzVDO0FBQUEsRUFDRixDQUFDO0FBRUQsTUFBSSxtQkFBSyxxQkFBb0I7QUFDM0IsdUJBQUssb0JBQW1CLFlBQVk7QUFFcEMsVUFBTSxnQkFBZ0IsU0FBUyxjQUFjLEtBQUs7QUFDbEQsa0JBQWMsWUFBWTtBQUUxQixVQUFNLFlBQVksU0FBUyxjQUFjLElBQUk7QUFDN0MsY0FBVSxjQUFjLFlBQVk7QUFDcEMsY0FBVSxZQUFZO0FBQ3RCLGtCQUFjLFlBQVksU0FBUztBQUVuQyxRQUFJLFlBQVksU0FBUztBQUN2QixZQUFNLFVBQVUsU0FBUyxjQUFjLEdBQUc7QUFDMUMsY0FBUSxjQUFjLFlBQVk7QUFDbEMsY0FBUSxZQUFZO0FBQ3BCLG9CQUFjLFlBQVksT0FBTztBQUFBLElBQ25DO0FBRUEsUUFBSSxZQUFZLGFBQWE7QUFDM0IsWUFBTSxjQUFjLFNBQVMsY0FBYyxHQUFHO0FBQzlDLGtCQUFZLGNBQWMsWUFBWTtBQUN0QyxrQkFBWSxZQUFZO0FBQ3hCLG9CQUFjLFlBQVksV0FBVztBQUFBLElBQ3ZDO0FBRUEsVUFBTSxPQUFPLFNBQVMsY0FBYyxLQUFLO0FBQ3pDLFNBQUssWUFBWTtBQUVqQixVQUFNLFNBQVMsU0FBUyxjQUFjLE1BQU07QUFDNUMsV0FBTyxhQUFhLFlBQVksWUFBWSxVQUFVLFlBQVksQ0FBQztBQUNuRSxXQUFPLGNBQWMsWUFBWSxVQUFVLG1CQUFtQjtBQUM5RCxXQUFPLFlBQVk7QUFFbkIsVUFBTSxVQUFVLFNBQVMsY0FBYyxNQUFNO0FBQzdDLFFBQUksY0FBYyxJQUFJLFlBQVksT0FBTztBQUN6QyxRQUFJLFlBQVksV0FBVztBQUN6QixxQkFBZSxJQUFJLFlBQVksU0FBUztBQUFBLElBQzFDO0FBQ0EsWUFBUSxjQUFjO0FBQ3RCLFlBQVEsWUFBWTtBQUVwQixTQUFLLFlBQVksTUFBTTtBQUN2QixTQUFLLFlBQVksT0FBTztBQUV4QixrQkFBYyxZQUFZLElBQUk7QUFFOUIsdUJBQUssb0JBQW1CLFlBQVksYUFBYTtBQUFBLEVBQ25EO0FBRUEsTUFBSSxtQkFBSyxtQkFBa0I7QUFDekIsdUJBQUssa0JBQWlCLFlBQVk7QUFFbEMsVUFBTSxvQkFBb0IsU0FBUyxjQUFjLElBQUk7QUFDckQsc0JBQWtCLGNBQWM7QUFDaEMsc0JBQWtCLFlBQVk7QUFFOUIsVUFBTSxzQkFBc0IsU0FBUyxjQUFjLEtBQUs7QUFDeEQsd0JBQW9CLFlBQVk7QUFFaEMsVUFBTSxrQkFBa0Isc0JBQUsseURBQUwsV0FBZ0M7QUFDeEQsUUFBSSxPQUFPLEtBQUssZUFBZSxFQUFFLFNBQVMsR0FBRztBQUMzQywwQkFBb0IsWUFBWSxzQkFBSyxpREFBTCxXQUF3QixnQkFBZ0I7QUFBQSxJQUMxRSxPQUFPO0FBQ0wsMEJBQW9CLGNBQWM7QUFBQSxJQUNwQztBQUVBLHVCQUFLLGtCQUFpQixZQUFZLGlCQUFpQjtBQUNuRCx1QkFBSyxrQkFBaUIsWUFBWSxtQkFBbUI7QUFBQSxFQUN2RDtBQUNGO0FBRUEsK0JBQTBCLFNBQUMsYUFBbUQ7QUFDNUUsUUFBTSxhQUFzQyxDQUFDO0FBRTdDLE1BQUksWUFBWSxrQkFBa0I7QUFDaEMsV0FBTyxPQUFPLFlBQVksWUFBWSxnQkFBZ0I7QUFBQSxFQUN4RDtBQUVBLGFBQVcsVUFBVSxZQUFZO0FBQ2pDLGFBQVcsYUFBYSxZQUFZO0FBQ3BDLGFBQVcsbUJBQW1CLFlBQVk7QUFDMUMsYUFBVyxXQUFXLFlBQVk7QUFFbEMsTUFBSSxZQUFZLGNBQWM7QUFDNUIsZUFBVyxPQUFPLFlBQVk7QUFBQSxFQUNoQztBQUVBLFNBQU87QUFDVDtBQUVBLHVCQUFrQixTQUFDLEtBQThCLFFBQVEsR0FBcUI7QUFDNUUsUUFBTSxLQUFLLFNBQVMsY0FBYyxJQUFJO0FBQ3RDLEtBQUcsWUFBWTtBQUNmLE1BQUksUUFBUSxHQUFHO0FBQ2IsT0FBRyxVQUFVLElBQUksUUFBUTtBQUFBLEVBQzNCO0FBRUEsYUFBVyxDQUFDLEtBQUssS0FBSyxLQUFLLE9BQU8sUUFBUSxHQUFHLEdBQUc7QUFDOUMsVUFBTSxLQUFLLFNBQVMsY0FBYyxJQUFJO0FBQ3RDLE9BQUcsWUFBWTtBQUVmLFVBQU0sVUFBVSxTQUFTLGNBQWMsTUFBTTtBQUM3QyxZQUFRLFlBQVk7QUFDcEIsWUFBUSxjQUFjO0FBRXRCLFVBQU0sWUFBWSxTQUFTLGNBQWMsTUFBTTtBQUMvQyxjQUFVLFlBQVk7QUFDdEIsY0FBVSxjQUFjO0FBRXhCLE9BQUcsWUFBWSxPQUFPO0FBQ3RCLE9BQUcsWUFBWSxTQUFTO0FBRXhCLFFBQUksVUFBVSxRQUFRLFVBQVUsUUFBVztBQUN6QyxZQUFNLFlBQVksU0FBUyxjQUFjLE1BQU07QUFDL0MsZ0JBQVUsWUFBWTtBQUN0QixnQkFBVSxjQUFjLE9BQU8sS0FBSztBQUNwQyxTQUFHLFlBQVksU0FBUztBQUFBLElBQzFCLFdBQVcsT0FBTyxVQUFVLFdBQVc7QUFDckMsWUFBTSxZQUFZLFNBQVMsY0FBYyxNQUFNO0FBQy9DLGdCQUFVLFlBQVk7QUFDdEIsZ0JBQVUsY0FBYyxPQUFPLEtBQUs7QUFDcEMsU0FBRyxZQUFZLFNBQVM7QUFBQSxJQUMxQixXQUFXLE9BQU8sVUFBVSxVQUFVO0FBQ3BDLFlBQU0sWUFBWSxTQUFTLGNBQWMsTUFBTTtBQUMvQyxnQkFBVSxZQUFZO0FBQ3RCLGdCQUFVLGNBQWMsT0FBTyxLQUFLO0FBQ3BDLFNBQUcsWUFBWSxTQUFTO0FBQUEsSUFDMUIsV0FBVyxPQUFPLFVBQVUsVUFBVTtBQUNwQyxZQUFNLFlBQVksU0FBUyxjQUFjLE1BQU07QUFDL0MsZ0JBQVUsWUFBWTtBQUN0QixnQkFBVSxjQUFjLElBQUksS0FBSztBQUNqQyxTQUFHLFlBQVksU0FBUztBQUFBLElBQzFCLFdBQVcsTUFBTSxRQUFRLEtBQUssR0FBRztBQUMvQixZQUFNLFlBQVksU0FBUyxjQUFjLE1BQU07QUFDL0MsZ0JBQVUsWUFBWTtBQUN0QixnQkFBVSxjQUFjLFNBQVMsTUFBTSxNQUFNO0FBQzdDLFNBQUcsWUFBWSxTQUFTO0FBRXhCLFVBQUksTUFBTSxTQUFTLEtBQUssUUFBUSxHQUFHO0FBQ2pDLGNBQU0sWUFBcUMsQ0FBQztBQUM1QyxjQUFNLFFBQVEsQ0FBQyxNQUFNLFVBQVU7QUFDN0Isb0JBQVUsS0FBSyxJQUFJO0FBQUEsUUFDckIsQ0FBQztBQUNELFdBQUcsWUFBWSxzQkFBSyxpREFBTCxXQUF3QixXQUFXLFFBQVEsRUFBRTtBQUFBLE1BQzlEO0FBQUEsSUFDRixXQUFXLE9BQU8sVUFBVSxVQUFVO0FBQ3BDLFlBQU0sWUFBWSxTQUFTLGNBQWMsTUFBTTtBQUMvQyxnQkFBVSxZQUFZO0FBQ3RCLFlBQU0sT0FBTyxPQUFPLEtBQUssS0FBZ0M7QUFDekQsZ0JBQVUsY0FBYyxLQUFLLFNBQVMsSUFBSSxXQUFXO0FBQ3JELFNBQUcsWUFBWSxTQUFTO0FBRXhCLFVBQUksS0FBSyxTQUFTLEtBQUssUUFBUSxHQUFHO0FBQ2hDLFdBQUcsWUFBWSxzQkFBSyxpREFBTCxXQUF3QixPQUFrQyxRQUFRLEVBQUU7QUFBQSxNQUNyRjtBQUFBLElBQ0YsT0FBTztBQUNMLFlBQU0sWUFBWSxTQUFTLGNBQWMsTUFBTTtBQUMvQyxnQkFBVSxZQUFZO0FBQ3RCLGdCQUFVLGNBQWMsT0FBTyxLQUFLO0FBQ3BDLFNBQUcsWUFBWSxTQUFTO0FBQUEsSUFDMUI7QUFFQSxPQUFHLFlBQVksRUFBRTtBQUFBLEVBQ25CO0FBRUEsU0FBTztBQUNUO0FBRUEsMEJBQXFCLFdBQUc7QUFDdEIsTUFBSSxDQUFDLG1CQUFLLFlBQVk7QUFFdEIsd0JBQXNCLE1BQU07QUFDMUIsVUFBTSxZQUFZLG1CQUFLLFlBQVk7QUFDbkMsUUFBSSxXQUFXO0FBQ2IsZ0JBQVUsZUFBZSxFQUFFLFVBQVUsUUFBUSxPQUFPLE1BQU0sQ0FBQztBQUFBLElBQzdEO0FBQUEsRUFDRixDQUFDO0FBQ0g7QUFFQSx1QkFBa0IsV0FBWTtBQUM1QixRQUFNLE9BQU8sS0FBSyxZQUFZLGNBQWMsWUFBWTtBQUN4RCxNQUFJLENBQUMsS0FBTSxRQUFPO0FBRWxCLFFBQU0sZ0JBQWdCLFNBQVMsS0FBSyxhQUFhLFVBQVUsS0FBSyxLQUFLLEVBQUU7QUFDdkUsU0FBTyxrQkFBa0I7QUFDM0I7QUFFQSxrQkFBYSxTQUFDLE9BQWU7QUFDM0IscUJBQUssb0JBQXFCLE1BQU0sWUFBWTtBQUU1QyxNQUFJLENBQUMsbUJBQUssWUFBWTtBQUV0QixhQUFXLFNBQVMsbUJBQUssWUFBVyxVQUFVO0FBQzVDLFVBQU0sY0FBYyxzQkFBSyxrREFBTCxXQUEwQixNQUFzQixRQUFRO0FBRTVFLFFBQUksQ0FBQyxZQUFhO0FBRWxCLFVBQU0sWUFBWSxzQkFBSyxzREFBTCxXQUE2QjtBQUMvQyxVQUFNLFlBQVksbUJBQUssbUJBQWtCLFNBQVMsS0FBSyxtQkFBSyxtQkFBa0IsSUFBSSxZQUFZLFNBQVM7QUFDdkcsVUFBTSxlQUFlLG1CQUFLLGlCQUFnQixTQUFTLEtBQUssbUJBQUssaUJBQWdCLElBQUksWUFBWSxPQUFPO0FBRXBHLElBQUMsTUFBc0IsU0FBUyxFQUFFLGFBQWEsYUFBYTtBQUFBLEVBQzlEO0FBQ0Y7QUFFQSw0QkFBdUIsU0FBQyxhQUFtQztBQUN6RCxNQUFJLENBQUMsbUJBQUssb0JBQW9CLFFBQU87QUFFckMsUUFBTSxhQUFhO0FBQUEsSUFDakIsWUFBWTtBQUFBLElBQ1osWUFBWTtBQUFBLElBQ1osWUFBWSxhQUFhO0FBQUEsSUFDekIsS0FBSyxVQUFVLFlBQVksb0JBQW9CLENBQUMsQ0FBQztBQUFBLEVBQ25ELEVBQUUsS0FBSyxHQUFHLEVBQUUsWUFBWTtBQUV4QixTQUFPLFdBQVcsU0FBUyxtQkFBSyxtQkFBa0I7QUFDcEQ7QUFFQSx3QkFBbUIsU0FBQyxJQUFxQztBQUN2RCxTQUFPLG1CQUFLLGlCQUFnQixLQUFLLE9BQUssRUFBRSxPQUFPLEVBQUU7QUFDbkQ7QUFFQSx3QkFBbUIsV0FBRztBQUNwQixRQUFNLG1CQUFtQixzQkFBSywyREFBTDtBQUV6QixRQUFNLGtCQUFrQixzQkFBSyxpQ0FBTCxXQUFRO0FBQ2hDLE1BQUksbUJBQW1CLG1CQUFLLG1CQUFrQjtBQUM1QyxRQUFJLE9BQU8sZ0JBQWdCLGNBQWMsWUFBWTtBQUNyRCxRQUFJLENBQUMsTUFBTTtBQUNULGFBQU8sU0FBUyxjQUFjLFlBQVk7QUFDMUMsc0JBQWdCLFlBQVksSUFBSTtBQUFBLElBQ2xDO0FBRUEsVUFBTSxnQkFBZ0IsS0FBSyxpQkFBaUIsaUJBQWlCO0FBQzdELGtCQUFjLFFBQVEsVUFBUSxLQUFLLE9BQU8sQ0FBQztBQUUzQyxVQUFNLGdCQUFnQixvQkFBSSxJQUFZO0FBQ3RDLGVBQVcsQ0FBQyxTQUFTLFNBQVMsS0FBSyxtQkFBSyxtQkFBa0I7QUFDeEQsVUFBSSxtQkFBSyxxQkFBb0IsSUFBSSxPQUFPLEdBQUc7QUFDekMsbUJBQVcsYUFBYSxVQUFVLFlBQVk7QUFDNUMsd0JBQWMsSUFBSSxTQUFTO0FBQUEsUUFDN0I7QUFBQSxNQUNGO0FBQUEsSUFDRjtBQUVBLFFBQUksaUJBQWlCLFlBQVk7QUFDL0IseUJBQUssbUJBQXFCLGlCQUFpQixXQUFtRixhQUFhLGFBQWE7QUFBQSxJQUMxSixPQUFPO0FBQ0wseUJBQUssbUJBQW9CLElBQUksSUFBSSxhQUFhO0FBQUEsSUFDaEQ7QUFFQSxlQUFXLGFBQWEsZUFBZTtBQUNyQyxZQUFNLE9BQU8sU0FBUyxjQUFjLGlCQUFpQjtBQUNyRCxXQUFLLGFBQWEsV0FBVyxVQUFVO0FBQ3ZDLFdBQUssYUFBYSxTQUFTLFNBQVM7QUFDcEMsVUFBSSxtQkFBSyxtQkFBa0IsSUFBSSxTQUFTLEdBQUc7QUFDekMsYUFBSyxhQUFhLFdBQVcsRUFBRTtBQUFBLE1BQ2pDO0FBQ0EsV0FBSyxjQUFjO0FBQ25CLFdBQUssWUFBWSxJQUFJO0FBQUEsSUFDdkI7QUFBQSxFQUNGO0FBRUEsUUFBTSxnQkFBZ0Isc0JBQUssaUNBQUwsV0FBUTtBQUM5QixNQUFJLGlCQUFpQixtQkFBSyxtQkFBa0I7QUFDMUMsUUFBSSxPQUFPLGNBQWMsY0FBYyxZQUFZO0FBQ25ELFFBQUksQ0FBQyxNQUFNO0FBQ1QsYUFBTyxTQUFTLGNBQWMsWUFBWTtBQUMxQyxvQkFBYyxZQUFZLElBQUk7QUFBQSxJQUNoQztBQUVBLFVBQU0sZ0JBQWdCLEtBQUssaUJBQWlCLGlCQUFpQjtBQUM3RCxrQkFBYyxRQUFRLFVBQVEsS0FBSyxPQUFPLENBQUM7QUFFM0MsVUFBTSxjQUFjLG9CQUFJLElBQVk7QUFDcEMsZUFBVyxXQUFXLG1CQUFLLGtCQUFpQixLQUFLLEdBQUc7QUFDbEQsVUFBSSxtQkFBSyxxQkFBb0IsSUFBSSxPQUFPLEdBQUc7QUFDekMsb0JBQVksSUFBSSxPQUFPO0FBQUEsTUFDekI7QUFBQSxJQUNGO0FBRUEsUUFBSSxpQkFBaUIsVUFBVTtBQUM3Qix5QkFBSyxpQkFBbUIsaUJBQWlCLFNBQWlGLGFBQWEsV0FBVztBQUFBLElBQ3BKLE9BQU87QUFDTCx5QkFBSyxpQkFBa0IsSUFBSSxJQUFJLFdBQVc7QUFBQSxJQUM1QztBQUVBLGVBQVcsV0FBVyxhQUFhO0FBQ2pDLFlBQU0sT0FBTyxTQUFTLGNBQWMsaUJBQWlCO0FBQ3JELFdBQUssYUFBYSxXQUFXLFVBQVU7QUFDdkMsV0FBSyxhQUFhLFNBQVMsT0FBTztBQUNsQyxVQUFJLG1CQUFLLGlCQUFnQixJQUFJLE9BQU8sR0FBRztBQUNyQyxhQUFLLGFBQWEsV0FBVyxFQUFFO0FBQUEsTUFDakM7QUFDQSxXQUFLLGNBQWMsSUFBSSxPQUFPO0FBQzlCLFdBQUssWUFBWSxJQUFJO0FBQUEsSUFDdkI7QUFBQSxFQUNGO0FBQ0Y7QUFFQTtBQWVBO0FBZUEsaUNBQTRCLFdBQXFFO0FBQy9GLFFBQU0sY0FBZ0Y7QUFBQSxJQUNwRixZQUFZO0FBQUEsSUFDWixVQUFVO0FBQUEsRUFDWjtBQUVBLE1BQUk7QUFDRixVQUFNLGtCQUFrQixhQUFhLFFBQVEsOEJBQThCO0FBQzNFLFFBQUksaUJBQWlCO0FBQ25CLGtCQUFZLGFBQWEsSUFBSSxJQUFJLEtBQUssTUFBTSxlQUFlLENBQUM7QUFBQSxJQUM5RDtBQUVBLFVBQU0sZ0JBQWdCLGFBQWEsUUFBUSwyQkFBMkI7QUFDdEUsUUFBSSxlQUFlO0FBQ2pCLGtCQUFZLFdBQVcsSUFBSSxJQUFJLEtBQUssTUFBTSxhQUFhLENBQUM7QUFBQSxJQUMxRDtBQUFBLEVBQ0YsU0FBUyxHQUFHO0FBQ1YsWUFBUSxNQUFNLCtEQUErRDtBQUFBLEVBQy9FO0FBRUEsU0FBTztBQUNUO0FBRUEsc0JBQWlCLFdBQUc7QUFDbEIsTUFBSTtBQUNGLGlCQUFhO0FBQUEsTUFBUTtBQUFBLE1BQ25CLEtBQUssVUFBVSxDQUFDLEdBQUcsbUJBQUssa0JBQWlCLENBQUM7QUFBQSxJQUFDO0FBQzdDLGlCQUFhO0FBQUEsTUFBUTtBQUFBLE1BQ25CLEtBQUssVUFBVSxDQUFDLEdBQUcsbUJBQUssZ0JBQWUsQ0FBQztBQUFBLElBQUM7QUFBQSxFQUM3QyxTQUFTLEdBQUc7QUFBQSxFQUVaO0FBQ0Y7QUFFQSxpQkFBWSxXQUFHO0FBQ2IscUJBQUssaUJBQWtCLENBQUM7QUFDeEIscUJBQUssa0JBQW1CO0FBQ3hCLE1BQUksbUJBQUssYUFBWTtBQUNuQix1QkFBSyxZQUFXLGdCQUFnQjtBQUFBLEVBQ2xDO0FBQ0EsTUFBSSxtQkFBSyxxQkFBb0I7QUFDM0IsdUJBQUssb0JBQW1CLFlBQVk7QUFBQSxFQUN0QztBQUNBLE1BQUksbUJBQUssbUJBQWtCO0FBQ3pCLHVCQUFLLGtCQUFpQixZQUFZO0FBQUEsRUFDcEM7QUFDRjtBQUVNLGdCQUFXLGlCQUFHO0FBQ2xCLE1BQUksQ0FBQyxtQkFBSyxZQUFZO0FBRXRCLFFBQU0sZ0JBQWdCLE1BQU0sS0FBSyxtQkFBSyxZQUFXLFFBQVEsRUFDdEQsT0FBTyxXQUFTLENBQUUsTUFBc0IsTUFBTSxFQUM5QyxJQUFJLFdBQVM7QUFDWixVQUFNLEtBQU0sTUFBc0IsUUFBUTtBQUMxQyxXQUFPLHNCQUFLLGtEQUFMLFdBQXlCO0FBQUEsRUFDbEMsQ0FBQyxFQUNBLE9BQU8sQ0FBQyxVQUFnQyxDQUFDLENBQUMsS0FBSyxFQUMvQyxJQUFJLFdBQVM7QUFDWixVQUFNLE9BQU8sTUFBTSxVQUFVLG1CQUFtQjtBQUNoRCxVQUFNLFVBQVUsTUFBTSxZQUFZLElBQUksTUFBTSxTQUFTLEtBQUssTUFBTTtBQUNoRSxVQUFNLFFBQVEsTUFBTSxvQkFBb0IsT0FBTyxLQUFLLE1BQU0sZ0JBQWdCLEVBQUUsU0FBUyxJQUNqRixNQUFNLEtBQUssVUFBVSxNQUFNLGdCQUFnQixDQUFDLEtBQzVDO0FBQ0osV0FBTyxJQUFJLElBQUksTUFBTSxNQUFNLE9BQU8sS0FBSyxPQUFPLFdBQVcsTUFBTSxTQUFTLEdBQUcsS0FBSztBQUFBLEVBQ2xGLENBQUMsRUFDQSxLQUFLLElBQUk7QUFFWixNQUFJLENBQUMsY0FBZTtBQUVwQixNQUFJO0FBQ0YsVUFBTSxVQUFVLFVBQVUsVUFBVSxhQUFhO0FBQ2pELFVBQU0sTUFBTSxzQkFBSyxpQ0FBTCxXQUFRO0FBQ3BCLFFBQUksS0FBSztBQUNQLFlBQU0sV0FBVyxNQUFNLEtBQUssSUFBSSxVQUFVLEVBQUU7QUFBQSxRQUMxQyxPQUFLLEVBQUUsYUFBYSxLQUFLLGNBQWMsRUFBRSxhQUFhLEtBQUssRUFBRSxVQUFVLEtBQUs7QUFBQSxNQUM5RTtBQUNBLFVBQUksVUFBVTtBQUNaLGNBQU0sV0FBVyxTQUFTO0FBQzFCLGlCQUFTLGNBQWM7QUFFdkIsWUFBSSxtQkFBSyw2QkFBNEI7QUFDbkMsdUJBQWEsbUJBQUssMkJBQTBCO0FBQUEsUUFDOUM7QUFFQSwyQkFBSyw0QkFBNkIsV0FBVyxNQUFNO0FBQ2pELGNBQUksS0FBSyxlQUFlLFNBQVMsWUFBWTtBQUMzQyxxQkFBUyxjQUFjO0FBQUEsVUFDekI7QUFDQSw2QkFBSyw0QkFBNkI7QUFBQSxRQUNwQyxHQUFHLEdBQUk7QUFBQSxNQUNUO0FBQUEsSUFDRjtBQUFBLEVBQ0YsU0FBUyxLQUFLO0FBQ1osWUFBUSxNQUFNLDZDQUE2QyxHQUFHO0FBQUEsRUFDaEU7QUFDRjtBQUVBLHlCQUFvQixXQUFHO0FBQ3JCLHFCQUFLLFlBQWEsc0JBQUssaUNBQUwsV0FBUTtBQUMxQixxQkFBSyxvQkFBcUIsc0JBQUssaUNBQUwsV0FBUTtBQUNsQyxxQkFBSyxrQkFBbUIsc0JBQUssaUNBQUwsV0FBUTtBQUVoQyxNQUFJLG1CQUFLLGFBQVk7QUFDbkIsdUJBQUssWUFBVyxpQkFBaUIsU0FBUyxDQUFDLE1BQWE7QUFDdEQsWUFBTSxXQUFZLEVBQUUsT0FBbUIsUUFBUSxrQkFBa0I7QUFDakUsVUFBSSxVQUFVO0FBQ1osY0FBTSxVQUFVLFNBQVMsUUFBUTtBQUNqQyxZQUFJLFNBQVM7QUFDWCxnQ0FBSywyQ0FBTCxXQUFrQjtBQUFBLFFBQ3BCO0FBQUEsTUFDRjtBQUFBLElBQ0YsQ0FBQztBQUFBLEVBQ0g7QUFFQSxRQUFNLGVBQWUsc0JBQUssaUNBQUwsV0FBUTtBQUM3QixNQUFJLGNBQWM7QUFDaEIsaUJBQWEsaUJBQWlCLFNBQVMsQ0FBQyxNQUFhO0FBQ25ELFlBQU0sRUFBRSxRQUFRLEdBQUcsSUFBSSxFQUFFO0FBQ3pCLG1CQUFhLG1CQUFLLDJCQUEyQjtBQUM3Qyx5QkFBSyw0QkFBNkIsV0FBVyxNQUFNO0FBQ2pELDhCQUFLLDRDQUFMLFdBQW1CO0FBQUEsTUFDckIsR0FBRyxHQUFHO0FBQUEsSUFDUixDQUFDO0FBQUEsRUFDSDtBQUVBLFFBQU0sa0JBQWtCLHNCQUFLLGlDQUFMLFdBQVE7QUFDaEMsTUFBSSxpQkFBaUI7QUFDbkIsb0JBQWdCLGlCQUFpQixVQUFVLG1CQUFLLDZCQUE2QztBQUFBLEVBQy9GO0FBRUEsUUFBTSxnQkFBZ0Isc0JBQUssaUNBQUwsV0FBUTtBQUM5QixNQUFJLGVBQWU7QUFDakIsa0JBQWMsaUJBQWlCLFVBQVUsbUJBQUssMkJBQTJDO0FBQUEsRUFDM0Y7QUFFQSx3QkFBSyxpQ0FBTCxXQUFRLGlCQUFpQixpQkFBaUIsU0FBUyxNQUFNO0FBQ3ZELDBCQUFLLDJDQUFMO0FBQUEsRUFDRixDQUFDO0FBRUQsd0JBQUssaUNBQUwsV0FBUSxnQkFBZ0IsaUJBQWlCLFNBQVMsTUFBTTtBQUN0RCwwQkFBSywwQ0FBTDtBQUFBLEVBQ0YsQ0FBQztBQUNIO0FBMy9EQSw0QkFBUyxrQkFEVCxxQkFoRlcsaUJBaUZGO0FBR1QsNEJBQVMsYUFEVCxnQkFuRlcsaUJBb0ZGO0FBR1QsNEJBQVMsZUFEVCxrQkF0RlcsaUJBdUZGO0FBR1QsNEJBQVMsZ0JBRFQsbUJBekZXLGlCQTBGRjtBQUdULDRCQUFTLGFBRFQsZ0JBNUZXLGlCQTZGRjtBQUdULDRCQUFTLFNBRFQsWUEvRlcsaUJBZ0dGO0FBR1QsNEJBQVMsVUFEVCxhQWxHVyxpQkFtR0Y7QUFHVCw0QkFBUyxnQkFEVCxtQkFyR1csaUJBc0dGO0FBR1QsNEJBQVMsZ0JBRFQsbUJBeEdXLGlCQXlHRjtBQUdULDRCQUFTLFdBRFQsY0EzR1csaUJBNEdGO0FBR1QsNEJBQVMsa0JBRFQscUJBOUdXLGlCQStHRjtBQS9HRSxrQkFBTiw4Q0FEUCw0QkFDYTtBQUNYLGNBRFcsaUJBQ0osVUFBUztBQUFBO0FBR2hCLGFBSlcsaUJBSUosb0JBQXFCLE1BQU07QUFDaEMsUUFBTSxJQUFJLFNBQVMsY0FBYyxVQUFVO0FBQzNDLElBQUUsWUFBWTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUF3QmQsU0FBTztBQUNULEdBQUc7QUFFSCxhQWpDVyxpQkFpQ0oscUJBQXNCLE1BQU07QUFDakMsUUFBTSxJQUFJLFNBQVMsY0FBYyxVQUFVO0FBQzNDLElBQUUsWUFBWTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQVVkLFNBQU87QUFDVCxHQUFHO0FBRUgsYUFoRFcsaUJBZ0RKLG9CQUFxQixNQUFNO0FBQ2hDLFFBQU0sSUFBSSxTQUFTLGNBQWMsVUFBVTtBQUMzQyxJQUFFLFlBQVk7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUtkLFNBQU87QUFDVCxHQUFHO0FBRUgsYUExRFcsaUJBMERKLG9CQUFxQixNQUFNO0FBQ2hDLFFBQU0sSUFBSSxTQUFTLGNBQWMsVUFBVTtBQUMzQyxJQUFFLFlBQVk7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBTWQsU0FBTztBQUNULEdBQUc7QUFFSCxhQXJFVyxpQkFxRUosc0JBQXVCLE1BQU07QUFDbEMsUUFBTSxJQUFJLFNBQVMsY0FBYyxVQUFVO0FBQzNDLElBQUUsWUFBWTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFNZCxTQUFPO0FBQ1QsR0FBRztBQTlFRSw0QkFBTTtBQUFOLElBQU0saUJBQU47IiwKICAibmFtZXMiOiBbXQp9Cg==
