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
                  <cem-health-panel ${this.primaryTagName ? html`component="${this.primaryTagName}"` : nothing}>
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
      __privateMethod(this, _CemServeChrome_instances, initWsClient_fn).call(this);
    }
    super.connectedCallback();
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1zZXJ2ZS1jaHJvbWUuY3NzJztcblxuaW1wb3J0ICcuLi9jZW0tY29sb3Itc2NoZW1lLXRvZ2dsZS9jZW0tY29sb3Itc2NoZW1lLXRvZ2dsZS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1kcmF3ZXIvY2VtLWRyYXdlci5qcyc7XG5pbXBvcnQgJy4uL2NlbS1oZWFsdGgtcGFuZWwvY2VtLWhlYWx0aC1wYW5lbC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1tYW5pZmVzdC1icm93c2VyL2NlbS1tYW5pZmVzdC1icm93c2VyLmpzJztcbmltcG9ydCAnLi4vY2VtLXJlY29ubmVjdGlvbi1jb250ZW50L2NlbS1yZWNvbm5lY3Rpb24tY29udGVudC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1zZXJ2ZS1kZW1vL2NlbS1zZXJ2ZS1kZW1vLmpzJztcbmltcG9ydCAnLi4vY2VtLXNlcnZlLWtub2ItZ3JvdXAvY2VtLXNlcnZlLWtub2ItZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9jZW0tc2VydmUta25vYnMvY2VtLXNlcnZlLWtub2JzLmpzJztcbmltcG9ydCAnLi4vY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5L2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWFsZXJ0L3BmLXY2LWFsZXJ0LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtYWxlcnQtZ3JvdXAvcGYtdjYtYWxlcnQtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtY2FyZC9wZi12Ni1jYXJkLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtYmFkZ2UvcGYtdjYtYmFkZ2UuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1kcm9wZG93bi9wZi12Ni1kcm9wZG93bi5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWV4cGFuZGFibGUtc2VjdGlvbi9wZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1sYWJlbC9wZi12Ni1sYWJlbC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW1hc3RoZWFkL3BmLXY2LW1hc3RoZWFkLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbW9kYWwvcGYtdjYtbW9kYWwuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXYtZ3JvdXAvcGYtdjYtbmF2LWdyb3VwLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbmF2LWl0ZW0vcGYtdjYtbmF2LWl0ZW0uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXYtbGluay9wZi12Ni1uYXYtbGluay5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW5hdi1saXN0L3BmLXY2LW5hdi1saXN0LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbmF2aWdhdGlvbi9wZi12Ni1uYXZpZ2F0aW9uLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtcGFnZS1tYWluL3BmLXY2LXBhZ2UtbWFpbi5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXBhZ2Utc2lkZWJhci9wZi12Ni1wYWdlLXNpZGViYXIuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1wYWdlL3BmLXY2LXBhZ2UuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1wb3BvdmVyL3BmLXY2LXBvcG92ZXIuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1zZWxlY3QvcGYtdjYtc2VsZWN0LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtc2tpcC10by1jb250ZW50L3BmLXY2LXNraXAtdG8tY29udGVudC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXN3aXRjaC9wZi12Ni1zd2l0Y2guanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10YWIvcGYtdjYtdGFiLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdGFicy9wZi12Ni10YWJzLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdGV4dC1pbnB1dC1ncm91cC9wZi12Ni10ZXh0LWlucHV0LWdyb3VwLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdGV4dC1pbnB1dC9wZi12Ni10ZXh0LWlucHV0LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0vcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10b2dnbGUtZ3JvdXAvcGYtdjYtdG9nZ2xlLWdyb3VwLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdG9vbGJhci1ncm91cC9wZi12Ni10b29sYmFyLWdyb3VwLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdG9vbGJhci1pdGVtL3BmLXY2LXRvb2xiYXItaXRlbS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvb2xiYXIvcGYtdjYtdG9vbGJhci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRyZWUtaXRlbS9wZi12Ni10cmVlLWl0ZW0uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10cmVlLXZpZXcvcGYtdjYtdHJlZS12aWV3LmpzJztcblxuLy8gQ2xpZW50LW9ubHkgbW9kdWxlcyBsb2FkZWQgZHluYW1pY2FsbHkgdG8gYXZvaWQgYnJlYWtpbmcgU1NSLlxuLy8gVGhlc2UgYXJlIHBsYWluIEpTIG1vZHVsZXMgc2VydmVkIGF0IHJ1bnRpbWUgYnkgdGhlIEdvIHNlcnZlci5cbnR5cGUgQ0VNUmVsb2FkQ2xpZW50VHlwZSA9IHsgbmV3KG9wdHM6IGFueSk6IGFueSB9O1xudHlwZSBTdGF0ZVBlcnNpc3RlbmNlVHlwZSA9IHtcbiAgZ2V0U3RhdGUoKTogYW55O1xuICB1cGRhdGVTdGF0ZShzOiBhbnkpOiB2b2lkO1xuICBnZXRUcmVlU3RhdGUoKTogYW55O1xuICBzZXRUcmVlU3RhdGUoczogYW55KTogdm9pZDtcbiAgbWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2UoKTogdm9pZDtcbn07XG5sZXQgQ0VNUmVsb2FkQ2xpZW50OiBDRU1SZWxvYWRDbGllbnRUeXBlO1xubGV0IFN0YXRlUGVyc2lzdGVuY2U6IFN0YXRlUGVyc2lzdGVuY2VUeXBlO1xuXG5pbnRlcmZhY2UgRXZlbnRJbmZvIHtcbiAgZXZlbnROYW1lczogU2V0PHN0cmluZz47XG4gIGV2ZW50czogQXJyYXk8eyBuYW1lOiBzdHJpbmc7IHR5cGU/OiB7IHRleHQ6IHN0cmluZyB9OyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9Pjtcbn1cblxuaW50ZXJmYWNlIEV2ZW50UmVjb3JkIHtcbiAgaWQ6IHN0cmluZztcbiAgdGltZXN0YW1wOiBEYXRlO1xuICBldmVudE5hbWU6IHN0cmluZztcbiAgdGFnTmFtZTogc3RyaW5nO1xuICBlbGVtZW50SWQ6IHN0cmluZyB8IG51bGw7XG4gIGVsZW1lbnRDbGFzczogc3RyaW5nIHwgbnVsbDtcbiAgY3VzdG9tUHJvcGVydGllczogUmVjb3JkPHN0cmluZywgdW5rbm93bj47XG4gIG1hbmlmZXN0VHlwZTogc3RyaW5nIHwgbnVsbDtcbiAgc3VtbWFyeTogc3RyaW5nIHwgbnVsbDtcbiAgZGVzY3JpcHRpb246IHN0cmluZyB8IG51bGw7XG4gIGJ1YmJsZXM6IGJvb2xlYW47XG4gIGNvbXBvc2VkOiBib29sZWFuO1xuICBjYW5jZWxhYmxlOiBib29sZWFuO1xuICBkZWZhdWx0UHJldmVudGVkOiBib29sZWFuO1xufVxuXG5pbnRlcmZhY2UgRGVidWdEYXRhIHtcbiAgdmVyc2lvbj86IHN0cmluZztcbiAgb3M/OiBzdHJpbmc7XG4gIHdhdGNoRGlyPzogc3RyaW5nO1xuICBtYW5pZmVzdFNpemU/OiBzdHJpbmc7XG4gIGRlbW9Db3VudD86IG51bWJlcjtcbiAgZGVtb3M/OiBBcnJheTx7XG4gICAgdGFnTmFtZTogc3RyaW5nO1xuICAgIGRlc2NyaXB0aW9uPzogc3RyaW5nO1xuICAgIGZpbGVwYXRoPzogc3RyaW5nO1xuICAgIGNhbm9uaWNhbFVSTDogc3RyaW5nO1xuICAgIGxvY2FsUm91dGU6IHN0cmluZztcbiAgfT47XG4gIGltcG9ydE1hcD86IFJlY29yZDxzdHJpbmcsIHVua25vd24+O1xuICBpbXBvcnRNYXBKU09OPzogc3RyaW5nO1xufVxuXG5pbnRlcmZhY2UgTWFuaWZlc3Qge1xuICBtb2R1bGVzPzogQXJyYXk8e1xuICAgIGRlY2xhcmF0aW9ucz86IEFycmF5PHtcbiAgICAgIGN1c3RvbUVsZW1lbnQ/OiBib29sZWFuO1xuICAgICAgdGFnTmFtZT86IHN0cmluZztcbiAgICAgIG5hbWU/OiBzdHJpbmc7XG4gICAgICBraW5kPzogc3RyaW5nO1xuICAgICAgZXZlbnRzPzogQXJyYXk8eyBuYW1lOiBzdHJpbmc7IHR5cGU/OiB7IHRleHQ6IHN0cmluZyB9OyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9PjtcbiAgICB9PjtcbiAgfT47XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZpcmVkIHdoZW4gbG9ncyBhcmUgcmVjZWl2ZWRcbiAqL1xuZXhwb3J0IGNsYXNzIENlbUxvZ3NFdmVudCBleHRlbmRzIEV2ZW50IHtcbiAgbG9nczogQXJyYXk8eyB0eXBlOiBzdHJpbmc7IGRhdGU6IHN0cmluZzsgbWVzc2FnZTogc3RyaW5nIH0+O1xuICBjb25zdHJ1Y3Rvcihsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT4pIHtcbiAgICBzdXBlcignY2VtOmxvZ3MnLCB7IGJ1YmJsZXM6IHRydWUgfSk7XG4gICAgdGhpcy5sb2dzID0gbG9ncztcbiAgfVxufVxuXG4vKipcbiAqIENFTSBTZXJ2ZSBDaHJvbWUgLSBNYWluIGRlbW8gd3JhcHBlciBjb21wb25lbnRcbiAqXG4gKiBAc2xvdCAtIERlZmF1bHQgc2xvdCBmb3IgZGVtbyBjb250ZW50XG4gKiBAc2xvdCBuYXZpZ2F0aW9uIC0gTmF2aWdhdGlvbiBzaWRlYmFyIGNvbnRlbnRcbiAqIEBzbG90IGtub2JzIC0gS25vYiBjb250cm9sc1xuICogQHNsb3QgZGVzY3JpcHRpb24gLSBEZW1vIGRlc2NyaXB0aW9uXG4gKiBAc2xvdCBtYW5pZmVzdC10cmVlIC0gTWFuaWZlc3QgdHJlZSB2aWV3XG4gKiBAc2xvdCBtYW5pZmVzdC1uYW1lIC0gTWFuaWZlc3QgbmFtZSBkaXNwbGF5XG4gKiBAc2xvdCBtYW5pZmVzdC1kZXRhaWxzIC0gTWFuaWZlc3QgZGV0YWlscyBkaXNwbGF5XG4gKlxuICogQGN1c3RvbUVsZW1lbnQgY2VtLXNlcnZlLWNocm9tZVxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXNlcnZlLWNocm9tZScpXG5leHBvcnQgY2xhc3MgQ2VtU2VydmVDaHJvbWUgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAvLyBTdGF0aWMgdGVtcGxhdGVzIGZvciBET00gY2xvbmluZyAobG9ncywgZXZlbnRzLCBkZWJ1ZyBpbmZvKVxuICBzdGF0aWMgI2RlbW9JbmZvVGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPGgzPkRlbW8gSW5mb3JtYXRpb248L2gzPlxuICAgICAgPGRsIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0IHBmLW0taG9yaXpvbnRhbCBwZi1tLWNvbXBhY3RcIj5cbiAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5UYWcgTmFtZTwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJ0YWdOYW1lXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCIgZGF0YS1maWVsZC1ncm91cD1cImRlc2NyaXB0aW9uXCI+XG4gICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+RGVzY3JpcHRpb248L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIj48L2RkPlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIiBkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIj5cbiAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5GaWxlIFBhdGg8L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIj48L2RkPlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5DYW5vbmljYWwgVVJMPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cImNhbm9uaWNhbFVSTFwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkxvY2FsIFJvdXRlPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cImxvY2FsUm91dGVcIj48L2RkPlxuICAgICAgICA8L2Rpdj5cbiAgICAgIDwvZGw+YDtcbiAgICByZXR1cm4gdDtcbiAgfSkoKTtcblxuICBzdGF0aWMgI2RlbW9Hcm91cFRlbXBsYXRlID0gKCgpID0+IHtcbiAgICBjb25zdCB0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGVtcGxhdGUnKTtcbiAgICB0LmlubmVySFRNTCA9IGBcbiAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiIGRhdGEtZmllbGQ9XCJ0YWdOYW1lXCI+PC9kdD5cbiAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiPlxuICAgICAgICAgIDxzcGFuIGRhdGEtZmllbGQ9XCJkZXNjcmlwdGlvblwiPjwvc3Bhbj48YnI+XG4gICAgICAgICAgPHNtYWxsIGRhdGEtZmllbGQtZ3JvdXA9XCJmaWxlcGF0aFwiPkZpbGU6IDxzcGFuIGRhdGEtZmllbGQ9XCJmaWxlcGF0aFwiPjwvc3Bhbj48L3NtYWxsPlxuICAgICAgICAgIDxzbWFsbD5DYW5vbmljYWw6IDxzcGFuIGRhdGEtZmllbGQ9XCJjYW5vbmljYWxVUkxcIj48L3NwYW4+PC9zbWFsbD48YnI+XG4gICAgICAgICAgPHNtYWxsPkxvY2FsOiA8c3BhbiBkYXRhLWZpZWxkPVwibG9jYWxSb3V0ZVwiPjwvc3Bhbj48L3NtYWxsPlxuICAgICAgICA8L2RkPlxuICAgICAgPC9kaXY+YDtcbiAgICByZXR1cm4gdDtcbiAgfSkoKTtcblxuICBzdGF0aWMgI2RlbW9MaXN0VGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPHBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbiBpZD1cImRlYnVnLWRlbW9zLXNlY3Rpb25cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB0b2dnbGUtdGV4dD1cIlNob3cgRGVtb3MgSW5mb1wiPlxuICAgICAgICA8ZGwgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3QgcGYtbS1ob3Jpem9udGFsIHBmLW0tY29tcGFjdFwiIGRhdGEtY29udGFpbmVyPVwiZ3JvdXBzXCI+PC9kbD5cbiAgICAgIDwvcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uPmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgc3RhdGljICNsb2dFbnRyeVRlbXBsYXRlID0gKCgpID0+IHtcbiAgICBjb25zdCB0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGVtcGxhdGUnKTtcbiAgICB0LmlubmVySFRNTCA9IGBcbiAgICAgIDxkaXYgY2xhc3M9XCJsb2ctZW50cnlcIiBkYXRhLWZpZWxkPVwiY29udGFpbmVyXCI+XG4gICAgICAgIDxwZi12Ni1sYWJlbCBjb21wYWN0IGRhdGEtZmllbGQ9XCJsYWJlbFwiPjwvcGYtdjYtbGFiZWw+XG4gICAgICAgIDx0aW1lIGNsYXNzPVwibG9nLXRpbWVcIiBkYXRhLWZpZWxkPVwidGltZVwiPjwvdGltZT5cbiAgICAgICAgPHNwYW4gY2xhc3M9XCJsb2ctbWVzc2FnZVwiIGRhdGEtZmllbGQ9XCJtZXNzYWdlXCI+PC9zcGFuPlxuICAgICAgPC9kaXY+YDtcbiAgICByZXR1cm4gdDtcbiAgfSkoKTtcblxuICBzdGF0aWMgI2V2ZW50RW50cnlUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8YnV0dG9uIGNsYXNzPVwiZXZlbnQtbGlzdC1pdGVtXCIgZGF0YS1maWVsZD1cImNvbnRhaW5lclwiPlxuICAgICAgICA8cGYtdjYtbGFiZWwgY29tcGFjdCBkYXRhLWZpZWxkPVwibGFiZWxcIj48L3BmLXY2LWxhYmVsPlxuICAgICAgICA8dGltZSBjbGFzcz1cImV2ZW50LXRpbWVcIiBkYXRhLWZpZWxkPVwidGltZVwiPjwvdGltZT5cbiAgICAgICAgPHNwYW4gY2xhc3M9XCJldmVudC1lbGVtZW50XCIgZGF0YS1maWVsZD1cImVsZW1lbnRcIj48L3NwYW4+XG4gICAgICA8L2J1dHRvbj5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ3ByaW1hcnktdGFnLW5hbWUnIH0pXG4gIGFjY2Vzc29yIHByaW1hcnlUYWdOYW1lID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAnZGVtby10aXRsZScgfSlcbiAgYWNjZXNzb3IgZGVtb1RpdGxlID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgYXR0cmlidXRlOiAncGFja2FnZS1uYW1lJyB9KVxuICBhY2Nlc3NvciBwYWNrYWdlTmFtZSA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ2Nhbm9uaWNhbC11cmwnIH0pXG4gIGFjY2Vzc29yIGNhbm9uaWNhbFVSTCA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ3NvdXJjZS11cmwnIH0pXG4gIGFjY2Vzc29yIHNvdXJjZVVSTCA9ICcnO1xuXG4gIEBwcm9wZXJ0eSgpXG4gIGFjY2Vzc29yIGtub2JzID0gJyc7XG5cbiAgQHByb3BlcnR5KClcbiAgYWNjZXNzb3IgZHJhd2VyOiAnZXhwYW5kZWQnIHwgJ2NvbGxhcHNlZCcgfCAnJyA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ2RyYXdlci1oZWlnaHQnIH0pXG4gIGFjY2Vzc29yIGRyYXdlckhlaWdodCA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ3RhYnMtc2VsZWN0ZWQnIH0pXG4gIGFjY2Vzc29yIHRhYnNTZWxlY3RlZCA9ICcnO1xuXG4gIEBwcm9wZXJ0eSgpXG4gIGFjY2Vzc29yIHNpZGViYXI6ICdleHBhbmRlZCcgfCAnY29sbGFwc2VkJyB8ICcnID0gJyc7XG5cbiAgQHByb3BlcnR5KHsgdHlwZTogQm9vbGVhbiwgYXR0cmlidXRlOiAnaGFzLWRlc2NyaXB0aW9uJyB9KVxuICBhY2Nlc3NvciBoYXNEZXNjcmlwdGlvbiA9IGZhbHNlO1xuXG4gICMkKGlkOiBzdHJpbmcpIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5nZXRFbGVtZW50QnlJZChpZCk7XG4gIH1cblxuICAjJCQoc2VsZWN0b3I6IHN0cmluZykge1xuICAgIHJldHVybiB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3JBbGwoc2VsZWN0b3IpID8/IFtdO1xuICB9XG5cbiAgI2xvZ0NvbnRhaW5lcjogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI2RyYXdlck9wZW4gPSBmYWxzZTtcbiAgI2luaXRpYWxMb2dzRmV0Y2hlZCA9IGZhbHNlO1xuICAjaXNJbml0aWFsTG9hZCA9IHRydWU7XG4gICNkZWJ1Z0RhdGE6IERlYnVnRGF0YSB8IG51bGwgPSBudWxsO1xuXG4gIC8vIEVsZW1lbnQgZXZlbnQgdHJhY2tpbmdcbiAgI2VsZW1lbnRFdmVudE1hcDogTWFwPHN0cmluZywgRXZlbnRJbmZvPiB8IG51bGwgPSBudWxsO1xuICAjbWFuaWZlc3Q6IE1hbmlmZXN0IHwgbnVsbCA9IG51bGw7XG4gICNjYXB0dXJlZEV2ZW50czogRXZlbnRSZWNvcmRbXSA9IFtdO1xuICAjbWF4Q2FwdHVyZWRFdmVudHMgPSAxMDAwO1xuICAjZXZlbnRMaXN0OiBIVE1MRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjZXZlbnREZXRhaWxIZWFkZXI6IEhUTUxFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNldmVudERldGFpbEJvZHk6IEhUTUxFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNzZWxlY3RlZEV2ZW50SWQ6IHN0cmluZyB8IG51bGwgPSBudWxsO1xuICAjZXZlbnRzRmlsdGVyVmFsdWUgPSAnJztcbiAgI2V2ZW50c0ZpbHRlckRlYm91bmNlVGltZXI6IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+IHwgbnVsbCA9IG51bGw7XG4gICNldmVudFR5cGVGaWx0ZXJzID0gbmV3IFNldDxzdHJpbmc+KCk7XG4gICNlbGVtZW50RmlsdGVycyA9IG5ldyBTZXQ8c3RyaW5nPigpO1xuICAjZGlzY292ZXJlZEVsZW1lbnRzID0gbmV3IFNldDxzdHJpbmc+KCk7XG5cbiAgLy8gRXZlbnQgbGlzdGVuZXIgcmVmZXJlbmNlcyBmb3IgY2xlYW51cFxuICAjaGFuZGxlTG9nc0V2ZW50OiAoKGV2ZW50OiBFdmVudCkgPT4gdm9pZCkgfCBudWxsID0gbnVsbDtcbiAgI2hhbmRsZVRyZWVFeHBhbmQ6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlVHJlZUNvbGxhcHNlOiAoKGV2ZW50OiBFdmVudCkgPT4gdm9pZCkgfCBudWxsID0gbnVsbDtcbiAgI2hhbmRsZVRyZWVTZWxlY3Q6ICgoZXZlbnQ6IEV2ZW50KSA9PiB2b2lkKSB8IG51bGwgPSBudWxsO1xuXG4gIC8vIFRpbWVvdXQgSURzIGZvciBjbGVhbnVwXG4gICNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dDogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcbiAgI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dDogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcbiAgI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQ6IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+IHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gTG9nIGZpbHRlciBzdGF0ZVxuICAjbG9nc0ZpbHRlclZhbHVlID0gJyc7XG4gICNsb2dzRmlsdGVyRGVib3VuY2VUaW1lcjogUmV0dXJuVHlwZTx0eXBlb2Ygc2V0VGltZW91dD4gfCBudWxsID0gbnVsbDtcbiAgI2xvZ0xldmVsRmlsdGVycyA9IG5ldyBTZXQoWydpbmZvJywgJ3dhcm4nLCAnZXJyb3InLCAnZGVidWcnXSk7XG4gICNsb2dMZXZlbERyb3Bkb3duOiBFbGVtZW50IHwgbnVsbCA9IG51bGw7XG5cbiAgLy8gV2F0Y2ggZm9yIGR5bmFtaWNhbGx5IGFkZGVkIGVsZW1lbnRzXG4gIC8qIGM4IGlnbm9yZSBzdGFydCAtIE11dGF0aW9uT2JzZXJ2ZXIgY2FsbGJhY2sgdGVzdGVkIHZpYSBpbnRlZ3JhdGlvbiAqL1xuICAjb2JzZXJ2ZXIgPSBuZXcgTXV0YXRpb25PYnNlcnZlcigobXV0YXRpb25zKSA9PiB7XG4gICAgbGV0IG5lZWRzVXBkYXRlID0gZmFsc2U7XG5cbiAgICBmb3IgKGNvbnN0IG11dGF0aW9uIG9mIG11dGF0aW9ucykge1xuICAgICAgZm9yIChjb25zdCBub2RlIG9mIG11dGF0aW9uLmFkZGVkTm9kZXMpIHtcbiAgICAgICAgaWYgKG5vZGUgaW5zdGFuY2VvZiBIVE1MRWxlbWVudCkge1xuICAgICAgICAgIGNvbnN0IHRhZ05hbWUgPSBub2RlLnRhZ05hbWUudG9Mb3dlckNhc2UoKTtcbiAgICAgICAgICBpZiAodGhpcy4jZWxlbWVudEV2ZW50TWFwPy5oYXModGFnTmFtZSkgJiYgIW5vZGUuZGF0YXNldC5jZW1FdmVudHNBdHRhY2hlZCkge1xuICAgICAgICAgICAgY29uc3QgZXZlbnRJbmZvID0gdGhpcy4jZWxlbWVudEV2ZW50TWFwLmdldCh0YWdOYW1lKSE7XG4gICAgICAgICAgICBmb3IgKGNvbnN0IGV2ZW50TmFtZSBvZiBldmVudEluZm8uZXZlbnROYW1lcykge1xuICAgICAgICAgICAgICBub2RlLmFkZEV2ZW50TGlzdGVuZXIoZXZlbnROYW1lLCB0aGlzLiNoYW5kbGVFbGVtZW50RXZlbnQsIHsgY2FwdHVyZTogdHJ1ZSB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIG5vZGUuZGF0YXNldC5jZW1FdmVudHNBdHRhY2hlZCA9ICd0cnVlJztcbiAgICAgICAgICAgIHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5hZGQodGFnTmFtZSk7XG4gICAgICAgICAgICBuZWVkc1VwZGF0ZSA9IHRydWU7XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgaWYgKG5lZWRzVXBkYXRlKSB7XG4gICAgICB0aGlzLiN1cGRhdGVFdmVudEZpbHRlcnMoKTtcbiAgICB9XG4gIH0pO1xuICAvKiBjOCBpZ25vcmUgc3RvcCAqL1xuXG4gICN3c0NsaWVudDogYW55O1xuICAjY2xpZW50TW9kdWxlc0xvYWRlZCA9IGZhbHNlO1xuXG4gICNpbml0V3NDbGllbnQoKSB7XG4gICAgdGhpcy4jd3NDbGllbnQgPSBuZXcgQ0VNUmVsb2FkQ2xpZW50KHtcbiAgICBqaXR0ZXJNYXg6IDEwMDAsXG4gICAgb3ZlcmxheVRocmVzaG9sZDogMTUsXG4gICAgYmFkZ2VGYWRlRGVsYXk6IDIwMDAsXG4gICAgLyogYzggaWdub3JlIHN0YXJ0IC0gV2ViU29ja2V0IGNhbGxiYWNrcyB0ZXN0ZWQgdmlhIGludGVncmF0aW9uICovXG4gICAgY2FsbGJhY2tzOiB7XG4gICAgICBvbk9wZW46ICgpID0+IHtcbiAgICAgICAgdGhpcy4jJCgncmVjb25uZWN0aW9uLW1vZGFsJyk/LmNsb3NlKCk7XG4gICAgICB9LFxuICAgICAgb25FcnJvcjogKGVycm9yRGF0YTogeyB0aXRsZT86IHN0cmluZzsgbWVzc2FnZT86IHN0cmluZzsgZmlsZT86IHN0cmluZyB9KSA9PiB7XG4gICAgICAgIGlmIChlcnJvckRhdGE/LnRpdGxlICYmIGVycm9yRGF0YT8ubWVzc2FnZSkge1xuICAgICAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmVdIFNlcnZlciBlcnJvcjonLCBlcnJvckRhdGEpO1xuICAgICAgICAgICh0aGlzLiMkKCdlcnJvci1vdmVybGF5JykgYXMgYW55KT8uc2hvdyhlcnJvckRhdGEudGl0bGUsIGVycm9yRGF0YS5tZXNzYWdlLCBlcnJvckRhdGEuZmlsZSk7XG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZV0gV2ViU29ja2V0IGVycm9yOicsIGVycm9yRGF0YSk7XG4gICAgICAgIH1cbiAgICAgIH0sXG4gICAgICBvblJlY29ubmVjdGluZzogKHsgYXR0ZW1wdCwgZGVsYXkgfTogeyBhdHRlbXB0OiBudW1iZXI7IGRlbGF5OiBudW1iZXIgfSkgPT4ge1xuICAgICAgICBpZiAoYXR0ZW1wdCA+PSAxNSkge1xuICAgICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tbW9kYWwnKSBhcyBhbnkpPy5zaG93TW9kYWwoKTtcbiAgICAgICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLWNvbnRlbnQnKSBhcyBhbnkpPy51cGRhdGVSZXRyeUluZm8oYXR0ZW1wdCwgZGVsYXkpO1xuICAgICAgICB9XG4gICAgICB9LFxuICAgICAgb25SZWxvYWQ6ICgpID0+IHtcbiAgICAgICAgY29uc3QgZXJyb3JPdmVybGF5ID0gdGhpcy4jJCgnZXJyb3Itb3ZlcmxheScpO1xuICAgICAgICBpZiAoZXJyb3JPdmVybGF5Py5oYXNBdHRyaWJ1dGUoJ29wZW4nKSkge1xuICAgICAgICAgIChlcnJvck92ZXJsYXkgYXMgYW55KS5oaWRlKCk7XG4gICAgICAgIH1cbiAgICAgICAgd2luZG93LmxvY2F0aW9uLnJlbG9hZCgpO1xuICAgICAgfSxcbiAgICAgIG9uU2h1dGRvd246ICgpID0+IHtcbiAgICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LnNob3dNb2RhbCgpO1xuICAgICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLWNvbnRlbnQnKSBhcyBhbnkpPy51cGRhdGVSZXRyeUluZm8oMzAsIDMwMDAwKTtcbiAgICAgIH0sXG4gICAgICBvbkxvZ3M6IChsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT4pID0+IHtcbiAgICAgICAgd2luZG93LmRpc3BhdGNoRXZlbnQobmV3IENlbUxvZ3NFdmVudChsb2dzKSk7XG4gICAgICB9XG4gICAgfVxuICAgIC8qIGM4IGlnbm9yZSBzdG9wICovXG4gICAgfSk7XG4gIH1cblxuICBnZXQgZGVtbygpOiBFbGVtZW50IHwgbnVsbCB7XG4gICAgcmV0dXJuIHRoaXMucXVlcnlTZWxlY3RvcignY2VtLXNlcnZlLWRlbW8nKTtcbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxsaW5rIHJlbD1cInN0eWxlc2hlZXRcIiBocmVmPVwiL19fY2VtL3BmLXY2LWMtZGVzY3JpcHRpb24tbGlzdC5jc3NcIj5cbiAgICAgIDxsaW5rIHJlbD1cInN0eWxlc2hlZXRcIiBocmVmPVwiL19fY2VtL3BmLWxpZ2h0ZG9tLmNzc1wiPlxuXG4gICAgICA8cGYtdjYtcGFnZSA/c2lkZWJhci1jb2xsYXBzZWQ9JHt0aGlzLnNpZGViYXIgPT09ICdjb2xsYXBzZWQnfT5cbiAgICAgICAgPHBmLXY2LXNraXAtdG8tY29udGVudCBzbG90PVwic2tpcC10by1jb250ZW50XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBocmVmPVwiI21haW4tY29udGVudFwiPlxuICAgICAgICAgIFNraXAgdG8gY29udGVudFxuICAgICAgICA8L3BmLXY2LXNraXAtdG8tY29udGVudD5cblxuICAgICAgICA8cGYtdjYtbWFzdGhlYWQgc2xvdD1cIm1hc3RoZWFkXCI+XG4gICAgICAgICAgPGEgY2xhc3M9XCJtYXN0aGVhZC1sb2dvXCJcbiAgICAgICAgICAgICBocmVmPVwiL1wiXG4gICAgICAgICAgICAgc2xvdD1cImxvZ29cIj5cbiAgICAgICAgICAgIDxpbWcgYWx0PVwiQ0VNIERldiBTZXJ2ZXJcIlxuICAgICAgICAgICAgICAgICBzcmM9XCIvX19jZW0vbG9nby5zdmdcIj5cbiAgICAgICAgICAgICR7dGhpcy5wYWNrYWdlTmFtZSA/IGh0bWxgPGgxPiR7dGhpcy5wYWNrYWdlTmFtZX08L2gxPmAgOiBub3RoaW5nfVxuICAgICAgICAgIDwvYT5cbiAgICAgICAgICA8cGYtdjYtdG9vbGJhciBzbG90PVwidG9vbGJhclwiPlxuICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItZ3JvdXAgdmFyaWFudD1cImFjdGlvbi1ncm91cFwiPlxuICAgICAgICAgICAgICAke3RoaXMuI3JlbmRlclNvdXJjZUJ1dHRvbigpfVxuICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJkZWJ1Zy1pbmZvXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiRGVidWcgaW5mb1wiPlxuICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCJcbiAgICAgICAgICAgICAgICAgICAgICAgaGVpZ2h0PVwiMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDE2IDE2XCJcbiAgICAgICAgICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIj5cbiAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk04IDE1QTcgNyAwIDEgMSA4IDFhNyA3IDAgMCAxIDAgMTR6bTAgMUE4IDggMCAxIDAgOCAwYTggOCAwIDAgMCAwIDE2elwiLz5cbiAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIm04LjkzIDYuNTg4LTIuMjkuMjg3LS4wODIuMzguNDUuMDgzYy4yOTQuMDcuMzUyLjE3Ni4yODguNDY5bC0uNzM4IDMuNDY4Yy0uMTk0Ljg5Ny4xMDUgMS4zMTkuODA4IDEuMzE5LjU0NSAwIDEuMTc4LS4yNTIgMS40NjUtLjU5OGwuMDg4LS40MTZjLS4yLjE3Ni0uNDkyLjI0Ni0uNjg2LjI0Ni0uMjc1IDAtLjM3NS0uMTkzLS4zMDQtLjUzM0w4LjkzIDYuNTg4ek05IDQuNWExIDEgMCAxIDEtMiAwIDEgMSAwIDAgMSAyIDB6XCIvPlxuICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgIDxwZi12Ni10b2dnbGUtZ3JvdXAgY2xhc3M9XCJjb2xvci1zY2hlbWUtdG9nZ2xlXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJDb2xvciBzY2hlbWVcIj5cbiAgICAgICAgICAgICAgICAgIDxwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cImxpZ2h0XCI+XG4gICAgICAgICAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIiBhcmlhLWxhYmVsPVwiTGlnaHQgbW9kZVwiPlxuICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNOCAxMWEzIDMgMCAxIDEgMC02IDMgMyAwIDAgMSAwIDZ6bTAgMWE0IDQgMCAxIDAgMC04IDQgNCAwIDAgMCAwIDh6TTggMGEuNS41IDAgMCAxIC41LjV2MmEuNS41IDAgMCAxLTEgMHYtMkEuNS41IDAgMCAxIDggMHptMCAxM2EuNS41IDAgMCAxIC41LjV2MmEuNS41IDAgMCAxLTEgMHYtMkEuNS41IDAgMCAxIDggMTN6bTgtNWEuNS41IDAgMCAxLS41LjVoLTJhLjUuNSAwIDAgMSAwLTFoMmEuNS41IDAgMCAxIC41LjV6TTMgOGEuNS41IDAgMCAxLS41LjVoLTJhLjUuNSAwIDAgMSAwLTFoMkEuNS41IDAgMCAxIDMgOHptMTAuNjU3LTUuNjU3YS41LjUgMCAwIDEgMCAuNzA3bC0xLjQxNCAxLjQxNWEuNS41IDAgMSAxLS43MDctLjcwOGwxLjQxNC0xLjQxNGEuNS41IDAgMCAxIC43MDcgMHptLTkuMTkzIDkuMTkzYS41LjUgMCAwIDEgMCAuNzA3TDMuMDUgMTMuNjU3YS41LjUgMCAwIDEtLjcwNy0uNzA3bDEuNDE0LTEuNDE0YS41LjUgMCAwIDEgLjcwNyAwem05LjE5MyAyLjEyMWEuNS41IDAgMCAxLS43MDcgMGwtMS40MTQtMS40MTRhLjUuNSAwIDAgMSAuNzA3LS43MDdsMS40MTQgMS40MTRhLjUuNSAwIDAgMSAwIC43MDd6TTQuNDY0IDQuNDY1YS41LjUgMCAwIDEtLjcwNyAwTDIuMzQzIDMuMDVhLjUuNSAwIDEgMSAuNzA3LS43MDdsMS40MTQgMS40MTRhLjUuNSAwIDAgMSAwIC43MDh6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0+XG4gICAgICAgICAgICAgICAgICA8cGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0gdmFsdWU9XCJkYXJrXCI+XG4gICAgICAgICAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIiBhcmlhLWxhYmVsPVwiRGFyayBtb2RlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk02IC4yNzhhLjc2OC43NjggMCAwIDEgLjA4Ljg1OCA3LjIwOCA3LjIwOCAwIDAgMC0uODc4IDMuNDZjMCA0LjAyMSAzLjI3OCA3LjI3NyA3LjMxOCA3LjI3Ny41MjcgMCAxLjA0LS4wNTUgMS41MzMtLjE2YS43ODcuNzg3IDAgMCAxIC44MS4zMTYuNzMzLjczMyAwIDAgMS0uMDMxLjg5M0E4LjM0OSA4LjM0OSAwIDAgMSA4LjM0NCAxNkMzLjczNCAxNiAwIDEyLjI4NiAwIDcuNzEgMCA0LjI2NiAyLjExNCAxLjMxMiA1LjEyNC4wNkEuNzUyLjc1MiAwIDAgMSA2IC4yNzh6TTQuODU4IDEuMzExQTcuMjY5IDcuMjY5IDAgMCAwIDEuMDI1IDcuNzFjMCA0LjAyIDMuMjc5IDcuMjc2IDcuMzE5IDcuMjc2YTcuMzE2IDcuMzE2IDAgMCAwIDUuMjA1LTIuMTYyYy0uMzM3LjA0Mi0uNjguMDYzLTEuMDI5LjA2My00LjYxIDAtOC4zNDMtMy43MTQtOC4zNDMtOC4yOSAwLTEuMTY3LjI0Mi0yLjI3OC42ODEtMy4yODZ6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0+XG4gICAgICAgICAgICAgICAgICA8cGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0gdmFsdWU9XCJzeXN0ZW1cIj5cbiAgICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiIGFyaWEtbGFiZWw9XCJTeXN0ZW0gcHJlZmVyZW5jZVwiPlxuICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNMCAxLjVBMS41IDEuNSAwIDAgMSAxLjUgMGgxM0ExLjUgMS41IDAgMCAxIDE2IDEuNXY4YTEuNSAxLjUgMCAwIDEtMS41IDEuNWgtMTNBMS41IDEuNSAwIDAgMSAwIDkuNXYtOHpNMS41IDFhLjUuNSAwIDAgMC0uNS41djhhLjUuNSAwIDAgMCAuNS41aDEzYS41LjUgMCAwIDAgLjUtLjV2LThhLjUuNSAwIDAgMC0uNS0uNWgtMTN6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNMi41IDEyaDExYS41LjUgMCAwIDEgMCAxaC0xMWEuNS41IDAgMCAxIDAtMXptMCAyaDExYS41LjUgMCAwIDEgMCAxaC0xMWEuNS41IDAgMCAxIDAtMXpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRvZ2dsZS1ncm91cD5cbiAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgPC9wZi12Ni10b29sYmFyPlxuICAgICAgICA8L3BmLXY2LW1hc3RoZWFkPlxuXG4gICAgICAgIDxwZi12Ni1wYWdlLXNpZGViYXIgc2xvdD1cInNpZGViYXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgID9leHBhbmRlZD0ke3RoaXMuc2lkZWJhciA9PT0gJ2V4cGFuZGVkJ31cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA/Y29sbGFwc2VkPSR7dGhpcy5zaWRlYmFyICE9PSAnZXhwYW5kZWQnfT5cbiAgICAgICAgICA8c2xvdCBuYW1lPVwibmF2aWdhdGlvblwiPjwvc2xvdD5cbiAgICAgICAgPC9wZi12Ni1wYWdlLXNpZGViYXI+XG5cbiAgICAgICAgPHBmLXY2LXBhZ2UtbWFpbiBzbG90PVwibWFpblwiIGlkPVwibWFpbi1jb250ZW50XCI+XG4gICAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgICAgIDxmb290ZXIgY2xhc3M9XCJwZi1tLXN0aWNreS1ib3R0b21cIj5cbiAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJmb290ZXItZGVzY3JpcHRpb24ke3RoaXMuaGFzRGVzY3JpcHRpb24gPyAnJyA6ICcgZW1wdHknfVwiPlxuICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwiZGVzY3JpcHRpb25cIj48L3Nsb3Q+XG4gICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgIDxjZW0tZHJhd2VyID9vcGVuPSR7dGhpcy5kcmF3ZXIgPT09ICdleHBhbmRlZCd9XG4gICAgICAgICAgICAgICAgICAgICAgICBkcmF3ZXItaGVpZ2h0PVwiJHt0aGlzLmRyYXdlckhlaWdodCB8fCAnNDAwJ31cIj5cbiAgICAgICAgICAgICAgPHBmLXY2LXRhYnMgc2VsZWN0ZWQ9XCIke3RoaXMudGFic1NlbGVjdGVkIHx8ICcwJ31cIj5cbiAgICAgICAgICAgICAgICA8cGYtdjYtdGFiIHRpdGxlPVwiS25vYnNcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgaWQ9XCJrbm9icy1jb250YWluZXJcIj5cbiAgICAgICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cImtub2JzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHAgY2xhc3M9XCJrbm9icy1lbXB0eVwiPk5vIGtub2JzIGF2YWlsYWJsZSBmb3IgdGhpcyBlbGVtZW50LjwvcD5cbiAgICAgICAgICAgICAgICAgICAgPC9zbG90PlxuICAgICAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICAgICAgPC9wZi12Ni10YWI+XG4gICAgICAgICAgICAgICAgPHBmLXY2LXRhYiB0aXRsZT1cIk1hbmlmZXN0IEJyb3dzZXJcIj5cbiAgICAgICAgICAgICAgICAgIDxjZW0tbWFuaWZlc3QtYnJvd3Nlcj5cbiAgICAgICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cIm1hbmlmZXN0LXRyZWVcIiBzbG90PVwibWFuaWZlc3QtdHJlZVwiPjwvc2xvdD5cbiAgICAgICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cIm1hbmlmZXN0LW5hbWVcIiBzbG90PVwibWFuaWZlc3QtbmFtZVwiPjwvc2xvdD5cbiAgICAgICAgICAgICAgICAgICAgPHNsb3QgbmFtZT1cIm1hbmlmZXN0LWRldGFpbHNcIiBzbG90PVwibWFuaWZlc3QtZGV0YWlsc1wiPjwvc2xvdD5cbiAgICAgICAgICAgICAgICAgIDwvY2VtLW1hbmlmZXN0LWJyb3dzZXI+XG4gICAgICAgICAgICAgICAgPC9wZi12Ni10YWI+XG4gICAgICAgICAgICAgICAgPHBmLXY2LXRhYiB0aXRsZT1cIlNlcnZlciBMb2dzXCI+XG4gICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPVwibG9ncy13cmFwcGVyXCI+XG4gICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyIHN0aWNreT5cbiAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10ZXh0LWlucHV0LWdyb3VwIGlkPVwibG9ncy1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBwbGFjZWhvbGRlcj1cIkZpbHRlciBsb2dzLi4uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgaWNvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHJvbGU9XCJwcmVzZW50YXRpb25cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxZW1cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgd2lkdGg9XCIxZW1cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmlld0JveD1cIjAgMCA1MTIgNTEyXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTUwNSA0NDIuN0w0MDUuMyAzNDNjLTQuNS00LjUtMTAuNi03LTE3LTdIMzcyYzI3LjYtMzUuMyA0NC03OS43IDQ0LTEyOEM0MTYgOTMuMSAzMjIuOSAwIDIwOCAwUzAgOTMuMSAwIDIwOHM5My4xIDIwOCAyMDggMjA4YzQ4LjMgMCA5Mi43LTE2LjQgMTI4LTQ0djE2LjNjMCA2LjQgMi41IDEyLjUgNyAxN2w5OS43IDk5LjdjOS40IDkuNCAyNC42IDkuNCAzMy45IDBsMjguMy0yOC4zYzkuNC05LjQgOS40LTI0LjYuMS0zNHpNMjA4IDMzNmMtNzAuNyAwLTEyOC01Ny4yLTEyOC0xMjggMC03MC43IDU3LjItMTI4IDEyOC0xMjggNzAuNyAwIDEyOCA1Ny4yIDEyOCAxMjggMCA3MC43LTU3LjIgMTI4LTEyOCAxMjh6XCI+PC9wYXRoPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRleHQtaW5wdXQtZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1kcm9wZG93biBpZD1cImxvZy1sZXZlbC1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgbG9nIGxldmVsc1wiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzcGFuIHNsb3Q9XCJ0b2dnbGUtdGV4dFwiPkxvZyBMZXZlbHM8L3NwYW4+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LW1lbnUtaXRlbSB2YXJpYW50PVwiY2hlY2tib3hcIiB2YWx1ZT1cImluZm9cIiBjaGVja2VkPkluZm88L3BmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwid2FyblwiIGNoZWNrZWQ+V2FybmluZ3M8L3BmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwiZXJyb3JcIiBjaGVja2VkPkVycm9yczwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJkZWJ1Z1wiIGNoZWNrZWQ+RGVidWc8L3BmLXY2LW1lbnUtaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1kcm9wZG93bj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cCB2YXJpYW50PVwiYWN0aW9uLWdyb3VwXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtYnV0dG9uIGlkPVwiY29weS1sb2dzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0xMyAwSDZhMiAyIDAgMCAwLTIgMiAyIDIgMCAwIDAtMiAydjEwYTIgMiAwIDAgMCAyIDJoN2EyIDIgMCAwIDAgMi0yIDIgMiAwIDAgMCAyLTJWMmEyIDIgMCAwIDAtMi0yem0wIDEzVjRhMiAyIDAgMCAwLTItMkg1YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjEwYTEgMSAwIDAgMS0xIDF6TTMgMTNWNGExIDEgMCAwIDEgMS0xaDdhMSAxIDAgMCAxIDEgMXY5YTEgMSAwIDAgMS0xIDFINGExIDEgMCAwIDEtMS0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDb3B5IExvZ3NcbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhcj5cbiAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImxvZy1jb250YWluZXJcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtdGFiPlxuICAgICAgICAgICAgICAgIDxwZi12Ni10YWIgdGl0bGU9XCJFdmVudHNcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJldmVudHMtd3JhcHBlclwiPlxuICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhciBzdGlja3k+XG4gICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdGV4dC1pbnB1dC1ncm91cCBpZD1cImV2ZW50cy1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBwbGFjZWhvbGRlcj1cIkZpbHRlciBldmVudHMuLi5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpY29uPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0xMS43NDIgMTAuMzQ0YTYuNSA2LjUgMCAxIDAtMS4zOTcgMS4zOThoLS4wMDFjLjAzLjA0LjA2Mi4wNzguMDk4LjExNWwzLjg1IDMuODVhMSAxIDAgMCAwIDEuNDE1LTEuNDE0bC0zLjg1LTMuODVhMS4wMDcgMS4wMDcgMCAwIDAtLjExNS0uMXpNMTIgNi41YTUuNSA1LjUgMCAxIDEtMTEgMCA1LjUgNS41IDAgMCAxIDExIDB6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRleHQtaW5wdXQtZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1kcm9wZG93biBpZD1cImV2ZW50LXR5cGUtZmlsdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGxhYmVsPVwiRmlsdGVyIGV2ZW50IHR5cGVzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHNwYW4gc2xvdD1cInRvZ2dsZS10ZXh0XCI+RXZlbnQgVHlwZXM8L3NwYW4+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtZHJvcGRvd24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1kcm9wZG93biBpZD1cImVsZW1lbnQtZmlsdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGxhYmVsPVwiRmlsdGVyIGVsZW1lbnRzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHNwYW4gc2xvdD1cInRvZ2dsZS10ZXh0XCI+RWxlbWVudHM8L3NwYW4+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtZHJvcGRvd24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItZ3JvdXAgdmFyaWFudD1cImFjdGlvbi1ncm91cFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImNsZWFyLWV2ZW50c1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInRlcnRpYXJ5XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBzaXplPVwic21hbGxcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCIgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNS41IDUuNUEuNS41IDAgMCAxIDYgNnY2YS41LjUgMCAwIDEtMSAwVjZhLjUuNSAwIDAgMSAuNS0uNXptMi41IDBhLjUuNSAwIDAgMSAuNS41djZhLjUuNSAwIDAgMS0xIDBWNmEuNS41IDAgMCAxIC41LS41em0zIC41YS41LjUgMCAwIDAtMSAwdjZhLjUuNSAwIDAgMCAxIDBWNnpcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBmaWxsLXJ1bGU9XCJldmVub2RkXCIgZD1cIk0xNC41IDNhMSAxIDAgMCAxLTEgMUgxM3Y5YTIgMiAwIDAgMS0yIDJINWEyIDIgMCAwIDEtMi0yVjRoLS41YTEgMSAwIDAgMS0xLTFWMmExIDEgMCAwIDEgMS0xSDZhMSAxIDAgMCAxIDEtMWgyYTEgMSAwIDAgMSAxIDFoMy41YTEgMSAwIDAgMSAxIDF2MXpNNC4xMTggNCA0IDQuMDU5VjEzYTEgMSAwIDAgMCAxIDFoNmExIDEgMCAwIDAgMS0xVjQuMDU5TDExLjg4MiA0SDQuMTE4ek0yLjUgM1YyaDExdjFoLTExelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDbGVhciBFdmVudHNcbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJjb3B5LWV2ZW50c1wiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInRlcnRpYXJ5XCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBzaXplPVwic21hbGxcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3ZnIHNsb3Q9XCJpY29uXCIgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNMTMgMEg2YTIgMiAwIDAgMC0yIDIgMiAyIDAgMCAwLTIgMnYxMGEyIDIgMCAwIDAgMiAyaDdhMiAyIDAgMCAwIDItMiAyIDIgMCAwIDAgMi0yVjJhMiAyIDAgMCAwLTItMnptMCAxM1Y0YTIgMiAwIDAgMC0yLTJINWExIDEgMCAwIDEgMS0xaDdhMSAxIDAgMCAxIDEgMXYxMGExIDEgMCAwIDEtMSAxek0zIDEzVjRhMSAxIDAgMCAxIDEtMWg3YTEgMSAwIDAgMSAxIDF2OWExIDEgMCAwIDEtMSAxSDRhMSAxIDAgMCAxLTEtMXpcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgQ29weSBFdmVudHNcbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItZ3JvdXA+XG4gICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhcj5cbiAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyYXdlciBpZD1cImV2ZW50LWRyYXdlclwiIGV4cGFuZGVkPlxuICAgICAgICAgICAgICAgICAgICAgIDxkaXYgaWQ9XCJldmVudC1saXN0XCI+PC9kaXY+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWRldGFpbC1oZWFkZXJcIiBzbG90PVwicGFuZWwtaGVhZGVyXCI+PC9kaXY+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWRldGFpbC1ib2R5XCIgc2xvdD1cInBhbmVsLWJvZHlcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1kcmF3ZXI+XG4gICAgICAgICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRhYj5cbiAgICAgICAgICAgICAgICA8cGYtdjYtdGFiIHRpdGxlPVwiSGVhbHRoXCI+XG4gICAgICAgICAgICAgICAgICA8Y2VtLWhlYWx0aC1wYW5lbCAke3RoaXMucHJpbWFyeVRhZ05hbWUgPyBodG1sYGNvbXBvbmVudD1cIiR7dGhpcy5wcmltYXJ5VGFnTmFtZX1cImAgOiBub3RoaW5nfT5cbiAgICAgICAgICAgICAgICAgIDwvY2VtLWhlYWx0aC1wYW5lbD5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRhYj5cbiAgICAgICAgICAgICAgPC9wZi12Ni10YWJzPlxuICAgICAgICAgICAgPC9jZW0tZHJhd2VyPlxuICAgICAgICAgIDwvZm9vdGVyPlxuICAgICAgICA8L3BmLXY2LXBhZ2UtbWFpbj5cbiAgICAgIDwvcGYtdjYtcGFnZT5cblxuICAgICAgPHBmLXY2LW1vZGFsIGlkPVwiZGVidWctbW9kYWxcIiB2YXJpYW50PVwibGFyZ2VcIj5cbiAgICAgICAgPGgyIHNsb3Q9XCJoZWFkZXJcIj5EZWJ1ZyBJbmZvcm1hdGlvbjwvaDI+XG4gICAgICAgIDxkbCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdCBwZi1tLWhvcml6b250YWwgcGYtbS1jb21wYWN0XCI+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlNlcnZlciBWZXJzaW9uPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLXZlcnNpb25cIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+U2VydmVyIE9TPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLW9zXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPldhdGNoIERpcmVjdG9yeTwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy13YXRjaC1kaXJcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+TWFuaWZlc3QgU2l6ZTwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy1tYW5pZmVzdC1zaXplXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkRlbW9zIEZvdW5kPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLWRlbW8tY291bnRcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+VGFnIE5hbWU8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiPiR7dGhpcy5wcmltYXJ5VGFnTmFtZSB8fCAnLSd9PC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+RGVtbyBUaXRsZTwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCI+JHt0aGlzLmRlbW9UaXRsZSB8fCAnLSd9PC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+QnJvd3NlcjwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy1icm93c2VyXCI+LTwvZGQ+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPlVzZXIgQWdlbnQ8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctdWFcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgPC9kbD5cbiAgICAgICAgPGRpdiBpZD1cImRlbW8tdXJscy1jb250YWluZXJcIj48L2Rpdj5cbiAgICAgICAgPHBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbiBpZD1cImRlYnVnLWltcG9ydG1hcC1kZXRhaWxzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB0b2dnbGUtdGV4dD1cIlNob3cgSW1wb3J0IE1hcFwiPlxuICAgICAgICAgIDxwcmUgaWQ9XCJkZWJ1Zy1pbXBvcnRtYXBcIj4tPC9wcmU+XG4gICAgICAgIDwvcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uPlxuICAgICAgICA8ZGl2IHNsb3Q9XCJmb290ZXJcIiBjbGFzcz1cImJ1dHRvbi1jb250YWluZXJcIj5cbiAgICAgICAgICA8cGYtdjYtYnV0dG9uIGNsYXNzPVwiZGVidWctY29weVwiIHZhcmlhbnQ9XCJwcmltYXJ5XCI+XG4gICAgICAgICAgICBDb3B5IERlYnVnIEluZm9cbiAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICA8cGYtdjYtYnV0dG9uIGNsYXNzPVwiZGVidWctY2xvc2VcIiB2YXJpYW50PVwic2Vjb25kYXJ5XCI+XG4gICAgICAgICAgICBDbG9zZVxuICAgICAgICAgIDwvcGYtdjYtYnV0dG9uPlxuICAgICAgICA8L2Rpdj5cbiAgICAgIDwvcGYtdjYtbW9kYWw+XG5cbiAgICAgIDwhLS0gUmVjb25uZWN0aW9uIG1vZGFsIC0tPlxuICAgICAgPHBmLXY2LW1vZGFsIGlkPVwicmVjb25uZWN0aW9uLW1vZGFsXCIgdmFyaWFudD1cImxhcmdlXCI+XG4gICAgICAgIDxoMiBzbG90PVwiaGVhZGVyXCI+RGV2ZWxvcG1lbnQgU2VydmVyIERpc2Nvbm5lY3RlZDwvaDI+XG4gICAgICAgIDxjZW0tcmVjb25uZWN0aW9uLWNvbnRlbnQgaWQ9XCJyZWNvbm5lY3Rpb24tY29udGVudFwiPjwvY2VtLXJlY29ubmVjdGlvbi1jb250ZW50PlxuICAgICAgICA8cGYtdjYtYnV0dG9uIGlkPVwicmVsb2FkLWJ1dHRvblwiXG4gICAgICAgICAgICAgICAgICAgICAgc2xvdD1cImZvb3RlclwiXG4gICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInByaW1hcnlcIj5SZWxvYWQgUGFnZTwvcGYtdjYtYnV0dG9uPlxuICAgICAgICA8cGYtdjYtYnV0dG9uIGlkPVwicmV0cnktYnV0dG9uXCJcbiAgICAgICAgICAgICAgICAgICAgICBzbG90PVwiZm9vdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwic2Vjb25kYXJ5XCI+UmV0cnkgTm93PC9wZi12Ni1idXR0b24+XG4gICAgICA8L3BmLXY2LW1vZGFsPlxuXG4gICAgICA8IS0tIFRyYW5zZm9ybSBlcnJvciBvdmVybGF5IC0tPlxuICAgICAgPGNlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheSBpZD1cImVycm9yLW92ZXJsYXlcIj5cbiAgICAgIDwvY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5PlxuICAgIGA7XG4gIH1cblxuICAjcmVuZGVyU291cmNlQnV0dG9uKCkge1xuICAgIGlmICghdGhpcy5zb3VyY2VVUkwpIHJldHVybiBub3RoaW5nO1xuXG4gICAgbGV0IGxhYmVsID0gJ1ZlcnNpb24gQ29udHJvbCc7XG4gICAgbGV0IHBhdGggPSAnTTUuODU0IDQuODU0YS41LjUgMCAxIDAtLjcwOC0uNzA4bC0zLjUgMy41YS41LjUgMCAwIDAgMCAuNzA4bDMuNSAzLjVhLjUuNSAwIDAgMCAuNzA4LS43MDhMMi43MDcgOGwzLjE0Ny0zLjE0NnptNC4yOTIgMGEuNS41IDAgMCAxIC43MDgtLjcwOGwzLjUgMy41YS41LjUgMCAwIDEgMCAuNzA4bC0zLjUgMy41YS41LjUgMCAwIDEtLjcwOC0uNzA4TDEzLjI5MyA4bC0zLjE0Ny0zLjE0NnonO1xuXG4gICAgaWYgKHRoaXMuc291cmNlVVJMLmluY2x1ZGVzKCdnaXRodWIuY29tJykpIHtcbiAgICAgIGxhYmVsID0gJ0dpdEh1Yi5jb20nO1xuICAgICAgcGF0aCA9ICdNOCAwQzMuNTggMCAwIDMuNTggMCA4YzAgMy41NCAyLjI5IDYuNTMgNS40NyA3LjU5LjQuMDcuNTUtLjE3LjU1LS4zOCAwLS4xOS0uMDEtLjgyLS4wMS0xLjQ5LTIuMDEuMzctMi41My0uNDktMi42OS0uOTQtLjA5LS4yMy0uNDgtLjk0LS44Mi0xLjEzLS4yOC0uMTUtLjY4LS41Mi0uMDEtLjUzLjYzLS4wMSAxLjA4LjU4IDEuMjMuODIuNzIgMS4yMSAxLjg3Ljg3IDIuMzMuNjYuMDctLjUyLjI4LS44Ny41MS0xLjA3LTEuNzgtLjItMy42NC0uODktMy42NC0zLjk1IDAtLjg3LjMxLTEuNTkuODItMi4xNS0uMDgtLjItLjM2LTEuMDIuMDgtMi4xMiAwIDAgLjY3LS4yMSAyLjIuODIuNjQtLjE4IDEuMzItLjI3IDItLjI3LjY4IDAgMS4zNi4wOSAyIC4yNyAxLjUzLTEuMDQgMi4yLS44MiAyLjItLjgyLjQ0IDEuMS4xNiAxLjkyLjA4IDIuMTIuNTEuNTYuODIgMS4yNy44MiAyLjE1IDAgMy4wNy0xLjg3IDMuNzUtMy42NSAzLjk1LjI5LjI1LjU0LjczLjU0IDEuNDggMCAxLjA3LS4wMSAxLjkzLS4wMSAyLjIgMCAuMjEuMTUuNDYuNTUuMzhBOC4wMTMgOC4wMTMgMCAwMDE2IDhjMC00LjQyLTMuNTgtOC04LTh6JztcbiAgICB9IGVsc2UgaWYgKHRoaXMuc291cmNlVVJMLmluY2x1ZGVzKCdnaXRsYWIuY29tJykpIHtcbiAgICAgIGxhYmVsID0gJ0dpdExhYic7XG4gICAgICBwYXRoID0gJ20xNS43MzQgNi4xLS4wMjItLjA1OEwxMy41MzQuMzU4YS41NjguNTY4IDAgMCAwLS41NjMtLjM1Ni41ODMuNTgzIDAgMCAwLS4zMjguMTIyLjU4Mi41ODIgMCAwIDAtLjE5My4yOTRsLTEuNDcgNC40OTlINS4wMjVsLTEuNDctNC41QS41NzIuNTcyIDAgMCAwIDMuMzYwLjE3NGEuNTcyLjU3MiAwIDAgMC0uMzI4LS4xNzIuNTgyLjU4MiAwIDAgMC0uNTYzLjM1N0wuMjkgNi4wNGwtLjAyMi4wNTdBNC4wNDQgNC4wNDQgMCAwIDAgMS42MSAxMC43N2wuMDA3LjAwNi4wMi4wMTQgMy4zMTggMi40ODUgMS42NCAxLjI0MiAxIC43NTVhLjY3My42NzMgMCAwIDAgLjgxNCAwbDEtLjc1NSAxLjY0LTEuMjQyIDMuMzM4LTIuNS4wMDktLjAwN2E0LjA1IDQuMDUgMCAwIDAgMS4zNC00LjY2OFonO1xuICAgIH0gZWxzZSBpZiAodGhpcy5zb3VyY2VVUkwuaW5jbHVkZXMoJ2JpdGJ1Y2tldC5vcmcnKSkge1xuICAgICAgbGFiZWwgPSAnQml0YnVja2V0JztcbiAgICAgIHBhdGggPSAnTTAgMS41QTEuNSAxLjUgMCAwIDEgMS41IDBoMTNBMS41IDEuNSAwIDAgMSAxNiAxLjV2MTNhMS41IDEuNSAwIDAgMS0xLjUgMS41aC0xM0ExLjUgMS41IDAgMCAxIDAgMTQuNXYtMTN6TTIuNSAzYS41LjUgMCAwIDAtLjUuNXY5YS41LjUgMCAwIDAgLjUuNWgxMWEuNS41IDAgMCAwIC41LS41di05YS41LjUgMCAwIDAtLjUtLjVoLTExem01LjAzOCAxLjQzNWEuNS41IDAgMCAxIC45MjQgMGwxLjQyIDMuMzdIOC43OGwtLjI0My42MDgtLjI0My0uNjA4SDUuMDgybDEuNDItMy4zN3pNOCA5LjE0M2wtLjc0MyAxLjg1N0g0Ljc0M0w2LjA3NiA3LjYwOCA4IDkuMTQzeic7XG4gICAgfVxuXG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICA8cGYtdjYtYnV0dG9uIGhyZWY9XCIke3RoaXMuc291cmNlVVJMfVwiXG4gICAgICAgICAgICAgICAgICAgICAgdGFyZ2V0PVwiX2JsYW5rXCJcbiAgICAgICAgICAgICAgICAgICAgICByZWw9XCJub29wZW5lciBub3JlZmVycmVyXCJcbiAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwicGxhaW5cIlxuICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJWaWV3IHNvdXJjZSBmaWxlXCI+XG4gICAgICAgICAgPHN2ZyBhcmlhLWxhYmVsPVwiJHtsYWJlbH1cIlxuICAgICAgICAgICAgICAgd2lkdGg9XCIxcmVtXCJcbiAgICAgICAgICAgICAgIGhlaWdodD1cIjFyZW1cIlxuICAgICAgICAgICAgICAgZmlsbD1cImN1cnJlbnRDb2xvclwiXG4gICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDE2IDE2XCI+XG4gICAgICAgICAgICA8cGF0aCBkPVwiJHtwYXRofVwiLz5cbiAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICBgO1xuICB9XG5cbiAgYXN5bmMgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgLy8gTG9hZCBjbGllbnQtb25seSBtb2R1bGVzIGJlZm9yZSBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpIHNvIHRoZXkncmVcbiAgICAvLyBhdmFpbGFibGUgd2hlbiBMaXQncyB1cGRhdGUgY3ljbGUgcnVucyAoZmlyc3RVcGRhdGVkLCBldGMuKVxuICAgIGlmICghdGhpcy4jY2xpZW50TW9kdWxlc0xvYWRlZCkge1xuICAgICAgW3sgQ0VNUmVsb2FkQ2xpZW50IH0sIHsgU3RhdGVQZXJzaXN0ZW5jZSB9XSA9IGF3YWl0IFByb21pc2UuYWxsKFtcbiAgICAgICAgLy8gQHRzLWlnbm9yZSAtLSBwbGFpbiBKUyBtb2R1bGVzIHNlcnZlZCBhdCBydW50aW1lIGJ5IEdvIHNlcnZlclxuICAgICAgICBpbXBvcnQoJy9fX2NlbS93ZWJzb2NrZXQtY2xpZW50LmpzJyksXG4gICAgICAgIC8vIEB0cy1pZ25vcmVcbiAgICAgICAgaW1wb3J0KCcvX19jZW0vc3RhdGUtcGVyc2lzdGVuY2UuanMnKSxcbiAgICAgIF0pO1xuICAgICAgLy8gQHRzLWlnbm9yZVxuICAgICAgaW1wb3J0KCcvX19jZW0vaGVhbHRoLWJhZGdlcy5qcycpLmNhdGNoKChlOiB1bmtub3duKSA9PlxuICAgICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlXSBGYWlsZWQgdG8gbG9hZCBoZWFsdGgtYmFkZ2VzOicsIGUpKTtcbiAgICAgIHRoaXMuI2NsaWVudE1vZHVsZXNMb2FkZWQgPSB0cnVlO1xuICAgICAgdGhpcy4jaW5pdFdzQ2xpZW50KCk7XG4gICAgfVxuXG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLiNtaWdyYXRlRnJvbUxvY2FsU3RvcmFnZUlmTmVlZGVkKCk7XG4gIH1cblxuICBmaXJzdFVwZGF0ZWQoKSB7XG4gICAgLy8gU2V0IHVwIGRlYnVnIG92ZXJsYXlcbiAgICB0aGlzLiNzZXR1cERlYnVnT3ZlcmxheSgpO1xuXG4gICAgLy8gU2V0IHVwIGNvbG9yIHNjaGVtZSB0b2dnbGVcbiAgICB0aGlzLiNzZXR1cENvbG9yU2NoZW1lVG9nZ2xlKCk7XG5cbiAgICAvLyBTZXQgdXAgZm9vdGVyIGRyYXdlciBhbmQgdGFic1xuICAgIHRoaXMuI3NldHVwRm9vdGVyRHJhd2VyKCk7XG5cbiAgICAvLyBMaXN0ZW4gZm9yIHNlcnZlciBsb2cgbWVzc2FnZXMgZnJvbSBXZWJTb2NrZXRcbiAgICB0aGlzLiNzZXR1cExvZ0xpc3RlbmVyKCk7XG5cbiAgICAvLyBTZXQgdXAga25vYiBldmVudCBjb29yZGluYXRpb25cbiAgICB0aGlzLiNzZXR1cEtub2JDb29yZGluYXRpb24oKTtcblxuICAgIC8vIFNldCB1cCB0cmVlIHN0YXRlIHBlcnNpc3RlbmNlXG4gICAgdGhpcy4jc2V0dXBUcmVlU3RhdGVQZXJzaXN0ZW5jZSgpO1xuXG4gICAgLy8gU2V0IHVwIHNpZGViYXIgc3RhdGUgcGVyc2lzdGVuY2VcbiAgICB0aGlzLiNzZXR1cFNpZGViYXJTdGF0ZVBlcnNpc3RlbmNlKCk7XG5cbiAgICAvLyBTZXQgdXAgZWxlbWVudCBldmVudCBjYXB0dXJlXG4gICAgdGhpcy4jc2V0dXBFdmVudENhcHR1cmUoKS50aGVuKCgpID0+IHtcbiAgICAgIHRoaXMuI3NldHVwRXZlbnRMaXN0ZW5lcnMoKTtcbiAgICB9KTtcblxuICAgIC8vIFNldCB1cCByZWNvbm5lY3Rpb24gbW9kYWwgYnV0dG9uIGhhbmRsZXJzXG4gICAgLyogYzggaWdub3JlIHN0YXJ0IC0gd2luZG93LmxvY2F0aW9uLnJlbG9hZCB3b3VsZCByZWxvYWQgdGVzdCBwYWdlICovXG4gICAgdGhpcy4jJCgncmVsb2FkLWJ1dHRvbicpPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgIHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKTtcbiAgICB9KTtcbiAgICAvKiBjOCBpZ25vcmUgc3RvcCAqL1xuXG4gICAgdGhpcy4jJCgncmV0cnktYnV0dG9uJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LmNsb3NlKCk7XG4gICAgICB0aGlzLiN3c0NsaWVudC5yZXRyeSgpO1xuICAgIH0pO1xuXG4gICAgLy8gSW5pdGlhbGl6ZSBXZWJTb2NrZXQgY29ubmVjdGlvblxuICAgIHRoaXMuI3dzQ2xpZW50LmluaXQoKTtcbiAgfVxuXG4gIGFzeW5jICNmZXRjaERlYnVnSW5mbygpIHtcbiAgICAvLyBQb3B1bGF0ZSBicm93c2VyIGluZm8gaW1tZWRpYXRlbHlcbiAgICBjb25zdCBicm93c2VyRWwgPSB0aGlzLiMkKCdkZWJ1Zy1icm93c2VyJyk7XG4gICAgY29uc3QgdWFFbCA9IHRoaXMuIyQoJ2RlYnVnLXVhJyk7XG4gICAgaWYgKGJyb3dzZXJFbCkge1xuICAgICAgY29uc3QgYnJvd3NlciA9IHRoaXMuI2RldGVjdEJyb3dzZXIoKTtcbiAgICAgIGJyb3dzZXJFbC50ZXh0Q29udGVudCA9IGJyb3dzZXI7XG4gICAgfVxuICAgIGlmICh1YUVsKSB7XG4gICAgICB1YUVsLnRleHRDb250ZW50ID0gbmF2aWdhdG9yLnVzZXJBZ2VudDtcbiAgICB9XG5cbiAgICAvLyBGZXRjaCBzZXJ2ZXIgZGVidWcgaW5mb1xuICAgIGNvbnN0IGRhdGEgPSBhd2FpdCBmZXRjaCgnL19fY2VtL2RlYnVnJylcbiAgICAgIC50aGVuKHJlcyA9PiByZXMuanNvbigpKVxuICAgICAgLmNhdGNoKChlcnI6IEVycm9yKSA9PiB7XG4gICAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gZmV0Y2ggZGVidWcgaW5mbzonLCBlcnIpO1xuICAgICAgfSk7XG5cbiAgICBpZiAoIWRhdGEpIHJldHVybjtcbiAgICBjb25zdCB2ZXJzaW9uRWwgPSB0aGlzLiMkKCdkZWJ1Zy12ZXJzaW9uJyk7XG4gICAgY29uc3Qgb3NFbCA9IHRoaXMuIyQoJ2RlYnVnLW9zJyk7XG4gICAgY29uc3Qgd2F0Y2hEaXJFbCA9IHRoaXMuIyQoJ2RlYnVnLXdhdGNoLWRpcicpO1xuICAgIGNvbnN0IG1hbmlmZXN0U2l6ZUVsID0gdGhpcy4jJCgnZGVidWctbWFuaWZlc3Qtc2l6ZScpO1xuICAgIGNvbnN0IGRlbW9Db3VudEVsID0gdGhpcy4jJCgnZGVidWctZGVtby1jb3VudCcpO1xuICAgIGNvbnN0IGRlbW9VcmxzQ29udGFpbmVyID0gdGhpcy4jJCgnZGVtby11cmxzLWNvbnRhaW5lcicpO1xuICAgIGNvbnN0IGltcG9ydE1hcEVsID0gdGhpcy4jJCgnZGVidWctaW1wb3J0bWFwJyk7XG5cbiAgICBpZiAodmVyc2lvbkVsKSB2ZXJzaW9uRWwudGV4dENvbnRlbnQgPSBkYXRhLnZlcnNpb24gfHwgJy0nO1xuICAgIGlmIChvc0VsKSBvc0VsLnRleHRDb250ZW50ID0gZGF0YS5vcyB8fCAnLSc7XG4gICAgaWYgKHdhdGNoRGlyRWwpIHdhdGNoRGlyRWwudGV4dENvbnRlbnQgPSBkYXRhLndhdGNoRGlyIHx8ICctJztcbiAgICBpZiAobWFuaWZlc3RTaXplRWwpIG1hbmlmZXN0U2l6ZUVsLnRleHRDb250ZW50ID0gZGF0YS5tYW5pZmVzdFNpemUgfHwgJy0nO1xuICAgIGlmIChkZW1vQ291bnRFbCkgZGVtb0NvdW50RWwudGV4dENvbnRlbnQgPSBkYXRhLmRlbW9Db3VudCB8fCAnMCc7XG5cbiAgICBpZiAoZGVtb1VybHNDb250YWluZXIpIHtcbiAgICAgIHRoaXMuI3BvcHVsYXRlRGVtb1VybHMoZGVtb1VybHNDb250YWluZXIsIGRhdGEuZGVtb3MpO1xuICAgIH1cblxuICAgIGlmIChpbXBvcnRNYXBFbCAmJiBkYXRhLmltcG9ydE1hcCkge1xuICAgICAgY29uc3QgaW1wb3J0TWFwSlNPTiA9IEpTT04uc3RyaW5naWZ5KGRhdGEuaW1wb3J0TWFwLCBudWxsLCAyKTtcbiAgICAgIGltcG9ydE1hcEVsLnRleHRDb250ZW50ID0gaW1wb3J0TWFwSlNPTjtcbiAgICAgIGRhdGEuaW1wb3J0TWFwSlNPTiA9IGltcG9ydE1hcEpTT047XG4gICAgfSBlbHNlIGlmIChpbXBvcnRNYXBFbCkge1xuICAgICAgaW1wb3J0TWFwRWwudGV4dENvbnRlbnQgPSAnTm8gaW1wb3J0IG1hcCBnZW5lcmF0ZWQnO1xuICAgIH1cblxuICAgIHRoaXMuI2RlYnVnRGF0YSA9IGRhdGE7XG4gIH1cblxuICAjcG9wdWxhdGVEZW1vVXJscyhjb250YWluZXI6IEhUTUxFbGVtZW50LCBkZW1vczogRGVidWdEYXRhWydkZW1vcyddKSB7XG4gICAgaWYgKCFkZW1vcz8ubGVuZ3RoKSB7XG4gICAgICBjb250YWluZXIudGV4dENvbnRlbnQgPSAnTm8gZGVtb3MgZm91bmQgaW4gbWFuaWZlc3QnO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IGN1cnJlbnRUYWdOYW1lID0gdGhpcy5wcmltYXJ5VGFnTmFtZSB8fCAnJztcbiAgICBjb25zdCBpc09uRGVtb1BhZ2UgPSAhIWN1cnJlbnRUYWdOYW1lO1xuXG4gICAgaWYgKGlzT25EZW1vUGFnZSkge1xuICAgICAgY29uc3QgY3VycmVudERlbW8gPSBkZW1vcy5maW5kKGRlbW8gPT4gZGVtby50YWdOYW1lID09PSBjdXJyZW50VGFnTmFtZSk7XG4gICAgICBpZiAoIWN1cnJlbnREZW1vKSB7XG4gICAgICAgIGNvbnRhaW5lci50ZXh0Q29udGVudCA9ICdDdXJyZW50IGRlbW8gbm90IGZvdW5kIGluIG1hbmlmZXN0JztcbiAgICAgICAgcmV0dXJuO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBmcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vSW5mb1RlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGFnTmFtZVwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLnRhZ05hbWU7XG4gICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNhbm9uaWNhbFVSTFwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmNhbm9uaWNhbFVSTDtcbiAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibG9jYWxSb3V0ZVwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmxvY2FsUm91dGU7XG5cbiAgICAgIGNvbnN0IGRlc2NyaXB0aW9uR3JvdXAgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImRlc2NyaXB0aW9uXCJdJyk7XG4gICAgICBpZiAoY3VycmVudERlbW8uZGVzY3JpcHRpb24pIHtcbiAgICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJkZXNjcmlwdGlvblwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmRlc2NyaXB0aW9uO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgZGVzY3JpcHRpb25Hcm91cD8ucmVtb3ZlKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGZpbGVwYXRoR3JvdXAgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImZpbGVwYXRoXCJdJyk7XG4gICAgICBpZiAoY3VycmVudERlbW8uZmlsZXBhdGgpIHtcbiAgICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJmaWxlcGF0aFwiXScpIS50ZXh0Q29udGVudCA9IGN1cnJlbnREZW1vLmZpbGVwYXRoO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgZmlsZXBhdGhHcm91cD8ucmVtb3ZlKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnRhaW5lci5yZXBsYWNlQ2hpbGRyZW4oZnJhZ21lbnQpO1xuICAgIH0gZWxzZSB7XG4gICAgICBjb25zdCBsaXN0RnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZGVtb0xpc3RUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICBjb25zdCBncm91cHNDb250YWluZXIgPSBsaXN0RnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtY29udGFpbmVyPVwiZ3JvdXBzXCJdJykhO1xuXG4gICAgICBmb3IgKGNvbnN0IGRlbW8gb2YgZGVtb3MpIHtcbiAgICAgICAgY29uc3QgZ3JvdXBGcmFnbWVudCA9IENlbVNlcnZlQ2hyb21lLiNkZW1vR3JvdXBUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0YWdOYW1lXCJdJykhLnRleHRDb250ZW50ID0gZGVtby50YWdOYW1lO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIl0nKSEudGV4dENvbnRlbnQgPVxuICAgICAgICAgIGRlbW8uZGVzY3JpcHRpb24gfHwgJyhubyBkZXNjcmlwdGlvbiknO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5jYW5vbmljYWxVUkw7XG4gICAgICAgIGdyb3VwRnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5sb2NhbFJvdXRlO1xuXG4gICAgICAgIGNvbnN0IGZpbGVwYXRoR3JvdXAgPSBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIl0nKTtcbiAgICAgICAgaWYgKGRlbW8uZmlsZXBhdGgpIHtcbiAgICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIl0nKSEudGV4dENvbnRlbnQgPSBkZW1vLmZpbGVwYXRoO1xuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIGZpbGVwYXRoR3JvdXA/LnJlbW92ZSgpO1xuICAgICAgICB9XG5cbiAgICAgICAgZ3JvdXBzQ29udGFpbmVyLmFwcGVuZENoaWxkKGdyb3VwRnJhZ21lbnQpO1xuICAgICAgfVxuXG4gICAgICBjb250YWluZXIucmVwbGFjZUNoaWxkcmVuKGxpc3RGcmFnbWVudCk7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwTG9nTGlzdGVuZXIoKSB7XG4gICAgdGhpcy4jbG9nQ29udGFpbmVyID0gdGhpcy4jJCgnbG9nLWNvbnRhaW5lcicpO1xuXG4gICAgY29uc3QgbG9nc0ZpbHRlciA9IHRoaXMuIyQoJ2xvZ3MtZmlsdGVyJykgYXMgSFRNTElucHV0RWxlbWVudCB8IG51bGw7XG4gICAgaWYgKGxvZ3NGaWx0ZXIpIHtcbiAgICAgIGxvZ3NGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignaW5wdXQnLCAoKSA9PiB7XG4gICAgICAgIGNvbnN0IHsgdmFsdWUgPSAnJyB9ID0gbG9nc0ZpbHRlcjtcbiAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyISk7XG4gICAgICAgIHRoaXMuI2xvZ3NGaWx0ZXJEZWJvdW5jZVRpbWVyID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgdGhpcy4jZmlsdGVyTG9ncyh2YWx1ZSk7XG4gICAgICAgIH0sIDMwMCk7XG4gICAgICB9KTtcbiAgICB9XG5cbiAgICB0aGlzLiNsb2dMZXZlbERyb3Bkb3duID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcjbG9nLWxldmVsLWZpbHRlcicpID8/IG51bGw7XG4gICAgaWYgKHRoaXMuI2xvZ0xldmVsRHJvcGRvd24pIHtcbiAgICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICAgIHRoaXMuI2xvYWRMb2dGaWx0ZXJTdGF0ZSgpO1xuICAgICAgfSk7XG4gICAgICB0aGlzLiNsb2dMZXZlbERyb3Bkb3duLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZUxvZ0ZpbHRlckNoYW5nZSBhcyBFdmVudExpc3RlbmVyKTtcbiAgICB9XG5cbiAgICB0aGlzLiMkKCdjb3B5LWxvZ3MnKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB0aGlzLiNjb3B5TG9ncygpO1xuICAgIH0pO1xuXG4gICAgdGhpcy4jaGFuZGxlTG9nc0V2ZW50ID0gKChldmVudDogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IGxvZ3MgPSAoZXZlbnQgYXMgQ2VtTG9nc0V2ZW50KS5sb2dzO1xuICAgICAgaWYgKGxvZ3MpIHtcbiAgICAgICAgdGhpcy4jcmVuZGVyTG9ncyhsb2dzKTtcbiAgICAgIH1cbiAgICB9KTtcbiAgICB3aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcignY2VtOmxvZ3MnLCB0aGlzLiNoYW5kbGVMb2dzRXZlbnQpO1xuICB9XG5cbiAgI2ZpbHRlckxvZ3MocXVlcnk6IHN0cmluZykge1xuICAgIHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgZm9yIChjb25zdCBlbnRyeSBvZiB0aGlzLiNsb2dDb250YWluZXIuY2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IHRleHQgPSBlbnRyeS50ZXh0Q29udGVudD8udG9Mb3dlckNhc2UoKSA/PyAnJztcbiAgICAgIGNvbnN0IHRleHRNYXRjaCA9ICF0aGlzLiNsb2dzRmlsdGVyVmFsdWUgfHwgdGV4dC5pbmNsdWRlcyh0aGlzLiNsb2dzRmlsdGVyVmFsdWUpO1xuXG4gICAgICBjb25zdCBsb2dUeXBlID0gdGhpcy4jZ2V0TG9nVHlwZUZyb21FbnRyeShlbnRyeSk7XG4gICAgICBjb25zdCBsZXZlbE1hdGNoID0gdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmhhcyhsb2dUeXBlKTtcblxuICAgICAgKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5oaWRkZW4gPSAhKHRleHRNYXRjaCAmJiBsZXZlbE1hdGNoKTtcbiAgICB9XG4gIH1cblxuICAjZ2V0TG9nVHlwZUZyb21FbnRyeShlbnRyeTogRWxlbWVudCk6IHN0cmluZyB7XG4gICAgZm9yIChjb25zdCBjbHMgb2YgZW50cnkuY2xhc3NMaXN0KSB7XG4gICAgICBpZiAoWydpbmZvJywgJ3dhcm5pbmcnLCAnZXJyb3InLCAnZGVidWcnXS5pbmNsdWRlcyhjbHMpKSB7XG4gICAgICAgIHJldHVybiBjbHMgPT09ICd3YXJuaW5nJyA/ICd3YXJuJyA6IGNscztcbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuICdpbmZvJztcbiAgfVxuXG4gICNsb2FkTG9nRmlsdGVyU3RhdGUoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHNhdmVkID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1sb2ctZmlsdGVycycpO1xuICAgICAgaWYgKHNhdmVkKSB7XG4gICAgICAgIHRoaXMuI2xvZ0xldmVsRmlsdGVycyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZCkpO1xuICAgICAgfVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIGNvbnNvbGUuZGVidWcoJ1tjZW0tc2VydmUtY2hyb21lXSBsb2NhbFN0b3JhZ2UgdW5hdmFpbGFibGUsIHVzaW5nIGRlZmF1bHQgbG9nIGZpbHRlcnMnKTtcbiAgICB9XG4gICAgdGhpcy4jc3luY0NoZWNrYm94U3RhdGVzKCk7XG4gIH1cblxuICAjc3luY0NoZWNrYm94U3RhdGVzKCkge1xuICAgIGlmICghdGhpcy4jbG9nTGV2ZWxEcm9wZG93bikgcmV0dXJuO1xuICAgIGNvbnN0IG1lbnVJdGVtcyA9IHRoaXMuI2xvZ0xldmVsRHJvcGRvd24ucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgbWVudUl0ZW1zLmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBjb25zdCB2YWx1ZSA9IChpdGVtIGFzIGFueSkudmFsdWU7XG4gICAgICAoaXRlbSBhcyBhbnkpLmNoZWNrZWQgPSB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuaGFzKHZhbHVlKTtcbiAgICB9KTtcbiAgfVxuXG4gICNzYXZlTG9nRmlsdGVyU3RhdGUoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdjZW0tc2VydmUtbG9nLWZpbHRlcnMnLFxuICAgICAgICBKU09OLnN0cmluZ2lmeShbLi4udGhpcy4jbG9nTGV2ZWxGaWx0ZXJzXSkpO1xuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIC8vIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSAocHJpdmF0ZSBtb2RlKSwgc2lsZW50bHkgY29udGludWVcbiAgICB9XG4gIH1cblxuICAjaGFuZGxlTG9nRmlsdGVyQ2hhbmdlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHsgdmFsdWUsIGNoZWNrZWQgfSA9IGV2ZW50IGFzIEV2ZW50ICYgeyB2YWx1ZTogc3RyaW5nOyBjaGVja2VkOiBib29sZWFuIH07XG5cbiAgICBpZiAoY2hlY2tlZCkge1xuICAgICAgdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2xvZ0xldmVsRmlsdGVycy5kZWxldGUodmFsdWUpO1xuICAgIH1cblxuICAgIHRoaXMuI3NhdmVMb2dGaWx0ZXJTdGF0ZSgpO1xuICAgIHRoaXMuI2ZpbHRlckxvZ3ModGhpcy4jbG9nc0ZpbHRlclZhbHVlKTtcbiAgfTtcblxuICBhc3luYyAjY29weUxvZ3MoKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGNvbnN0IGxvZ3MgPSBBcnJheS5mcm9tKHRoaXMuI2xvZ0NvbnRhaW5lci5jaGlsZHJlbilcbiAgICAgIC5maWx0ZXIoZW50cnkgPT4gIShlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuKVxuICAgICAgLm1hcChlbnRyeSA9PiB7XG4gICAgICAgIGNvbnN0IHR5cGUgPSBlbnRyeS5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxhYmVsXCJdJyk/LnRleHRDb250ZW50Py50cmltKCkgfHwgJ0lORk8nO1xuICAgICAgICBjb25zdCB0aW1lID0gZW50cnkucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0aW1lXCJdJyk/LnRleHRDb250ZW50Py50cmltKCkgfHwgJyc7XG4gICAgICAgIGNvbnN0IG1lc3NhZ2UgPSBlbnRyeS5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cIm1lc3NhZ2VcIl0nKT8udGV4dENvbnRlbnQ/LnRyaW0oKSB8fCAnJztcbiAgICAgICAgcmV0dXJuIGBbJHt0eXBlfV0gJHt0aW1lfSAke21lc3NhZ2V9YDtcbiAgICAgIH0pLmpvaW4oJ1xcbicpO1xuXG4gICAgaWYgKCFsb2dzKSByZXR1cm47XG5cbiAgICB0cnkge1xuICAgICAgYXdhaXQgbmF2aWdhdG9yLmNsaXBib2FyZC53cml0ZVRleHQobG9ncyk7XG4gICAgICBjb25zdCBidG4gPSB0aGlzLiMkKCdjb3B5LWxvZ3MnKTtcbiAgICAgIGlmIChidG4pIHtcbiAgICAgICAgY29uc3QgdGV4dE5vZGUgPSBBcnJheS5mcm9tKGJ0bi5jaGlsZE5vZGVzKS5maW5kKFxuICAgICAgICAgIG4gPT4gbi5ub2RlVHlwZSA9PT0gTm9kZS5URVhUX05PREUgJiYgKG4udGV4dENvbnRlbnQ/LnRyaW0oKS5sZW5ndGggPz8gMCkgPiAwXG4gICAgICAgICk7XG4gICAgICAgIGlmICh0ZXh0Tm9kZSkge1xuICAgICAgICAgIGNvbnN0IG9yaWdpbmFsID0gdGV4dE5vZGUudGV4dENvbnRlbnQ7XG4gICAgICAgICAgdGV4dE5vZGUudGV4dENvbnRlbnQgPSAnQ29waWVkISc7XG5cbiAgICAgICAgICBpZiAodGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgdGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICAgIGlmICh0aGlzLmlzQ29ubmVjdGVkICYmIHRleHROb2RlLnBhcmVudE5vZGUpIHtcbiAgICAgICAgICAgICAgdGV4dE5vZGUudGV4dENvbnRlbnQgPSBvcmlnaW5hbDtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICAgICAgICB9LCAyMDAwKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gY2F0Y2ggKGVycikge1xuICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjb3B5IGxvZ3M6JywgZXJyKTtcbiAgICB9XG4gIH1cblxuICAjc2V0dXBEZWJ1Z092ZXJsYXkoKSB7XG4gICAgY29uc3QgZGVidWdCdXR0b24gPSB0aGlzLiMkKCdkZWJ1Zy1pbmZvJyk7XG4gICAgY29uc3QgZGVidWdNb2RhbCA9IHRoaXMuIyQoJ2RlYnVnLW1vZGFsJyk7XG4gICAgY29uc3QgZGVidWdDbG9zZSA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignLmRlYnVnLWNsb3NlJyk7XG4gICAgY29uc3QgZGVidWdDb3B5ID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuZGVidWctY29weScpO1xuXG4gICAgaWYgKGRlYnVnQnV0dG9uICYmIGRlYnVnTW9kYWwpIHtcbiAgICAgIGRlYnVnQnV0dG9uLmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgICB0aGlzLiNmZXRjaERlYnVnSW5mbygpO1xuICAgICAgICAoZGVidWdNb2RhbCBhcyBhbnkpLnNob3dNb2RhbCgpO1xuICAgICAgfSk7XG5cbiAgICAgIGRlYnVnQ2xvc2U/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4gKGRlYnVnTW9kYWwgYXMgYW55KS5jbG9zZSgpKTtcblxuICAgICAgZGVidWdDb3B5Py5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgICAgdGhpcy4jY29weURlYnVnSW5mbygpO1xuICAgICAgfSk7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwRm9vdGVyRHJhd2VyKCkge1xuICAgIGNvbnN0IGRyYXdlciA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignY2VtLWRyYXdlcicpO1xuICAgIGNvbnN0IHRhYnMgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ3BmLXY2LXRhYnMnKTtcblxuICAgIGlmICghZHJhd2VyIHx8ICF0YWJzKSByZXR1cm47XG5cbiAgICB0aGlzLiNkcmF3ZXJPcGVuID0gKGRyYXdlciBhcyBhbnkpLm9wZW47XG5cbiAgICBkcmF3ZXIuYWRkRXZlbnRMaXN0ZW5lcignY2hhbmdlJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICB0aGlzLiNkcmF3ZXJPcGVuID0gKGUgYXMgYW55KS5vcGVuO1xuXG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHtcbiAgICAgICAgZHJhd2VyOiB7IG9wZW46IChlIGFzIGFueSkub3BlbiB9XG4gICAgICB9KTtcblxuICAgICAgaWYgKChlIGFzIGFueSkub3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxMb2dzVG9Cb3R0b20oKTtcbiAgICAgIH1cbiAgICB9KTtcblxuICAgIGRyYXdlci5hZGRFdmVudExpc3RlbmVyKCdyZXNpemUnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgIChkcmF3ZXIgYXMgYW55KS5zZXRBdHRyaWJ1dGUoJ2RyYXdlci1oZWlnaHQnLCAoZSBhcyBhbnkpLmhlaWdodCk7XG5cbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoe1xuICAgICAgICBkcmF3ZXI6IHsgaGVpZ2h0OiAoZSBhcyBhbnkpLmhlaWdodCB9XG4gICAgICB9KTtcbiAgICB9KTtcblxuICAgIHRhYnMuYWRkRXZlbnRMaXN0ZW5lcignY2hhbmdlJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHtcbiAgICAgICAgdGFiczogeyBzZWxlY3RlZEluZGV4OiAoZSBhcyBhbnkpLnNlbGVjdGVkSW5kZXggfVxuICAgICAgfSk7XG5cbiAgICAgIGlmICgoZSBhcyBhbnkpLnNlbGVjdGVkSW5kZXggPT09IDIgJiYgKGRyYXdlciBhcyBhbnkpLm9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTG9nc1RvQm90dG9tKCk7XG4gICAgICB9XG5cbiAgICAgIGlmICgoZSBhcyBhbnkpLnNlbGVjdGVkSW5kZXggPT09IDMgJiYgKGRyYXdlciBhcyBhbnkpLm9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsRXZlbnRzVG9Cb3R0b20oKTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNkZXRlY3RCcm93c2VyKCk6IHN0cmluZyB7XG4gICAgY29uc3QgdWEgPSBuYXZpZ2F0b3IudXNlckFnZW50O1xuICAgIGlmICh1YS5pbmNsdWRlcygnRmlyZWZveC8nKSkge1xuICAgICAgY29uc3QgbWF0Y2ggPSB1YS5tYXRjaCgvRmlyZWZveFxcLyhcXGQrKS8pO1xuICAgICAgcmV0dXJuIG1hdGNoID8gYEZpcmVmb3ggJHttYXRjaFsxXX1gIDogJ0ZpcmVmb3gnO1xuICAgIH0gZWxzZSBpZiAodWEuaW5jbHVkZXMoJ0VkZy8nKSkge1xuICAgICAgY29uc3QgbWF0Y2ggPSB1YS5tYXRjaCgvRWRnXFwvKFxcZCspLyk7XG4gICAgICByZXR1cm4gbWF0Y2ggPyBgRWRnZSAke21hdGNoWzFdfWAgOiAnRWRnZSc7XG4gICAgfSBlbHNlIGlmICh1YS5pbmNsdWRlcygnQ2hyb21lLycpKSB7XG4gICAgICBjb25zdCBtYXRjaCA9IHVhLm1hdGNoKC9DaHJvbWVcXC8oXFxkKykvKTtcbiAgICAgIHJldHVybiBtYXRjaCA/IGBDaHJvbWUgJHttYXRjaFsxXX1gIDogJ0Nocm9tZSc7XG4gICAgfSBlbHNlIGlmICh1YS5pbmNsdWRlcygnU2FmYXJpLycpICYmICF1YS5pbmNsdWRlcygnQ2hyb21lJykpIHtcbiAgICAgIGNvbnN0IG1hdGNoID0gdWEubWF0Y2goL1ZlcnNpb25cXC8oXFxkKykvKTtcbiAgICAgIHJldHVybiBtYXRjaCA/IGBTYWZhcmkgJHttYXRjaFsxXX1gIDogJ1NhZmFyaSc7XG4gICAgfVxuICAgIHJldHVybiAnVW5rbm93bic7XG4gIH1cblxuICBhc3luYyAjY29weURlYnVnSW5mbygpIHtcbiAgICBjb25zdCBpbmZvID0gQXJyYXkuZnJvbSh0aGlzLiMkJCgnI2RlYnVnLW1vZGFsIGRsIGR0JykpLm1hcChkdCA9PiB7XG4gICAgICBjb25zdCBkZCA9IGR0Lm5leHRFbGVtZW50U2libGluZztcbiAgICAgIGlmIChkZCAmJiBkZC50YWdOYW1lID09PSAnREQnKSB7XG4gICAgICAgIHJldHVybiBgJHtkdC50ZXh0Q29udGVudH06ICR7ZGQudGV4dENvbnRlbnR9YDtcbiAgICAgIH1cbiAgICAgIHJldHVybiAnJztcbiAgICB9KS5qb2luKCdcXG4nKTtcblxuICAgIGxldCBpbXBvcnRNYXBTZWN0aW9uID0gJyc7XG4gICAgaWYgKHRoaXMuI2RlYnVnRGF0YT8uaW1wb3J0TWFwSlNPTikge1xuICAgICAgaW1wb3J0TWFwU2VjdGlvbiA9IGBcXG4keyc9Jy5yZXBlYXQoNDApfVxcbkltcG9ydCBNYXA6XFxuJHsnPScucmVwZWF0KDQwKX1cXG4ke3RoaXMuI2RlYnVnRGF0YS5pbXBvcnRNYXBKU09OfVxcbmA7XG4gICAgfVxuXG4gICAgY29uc3QgZGVidWdUZXh0ID0gYENFTSBTZXJ2ZSBEZWJ1ZyBJbmZvcm1hdGlvblxuJHsnPScucmVwZWF0KDQwKX1cbiR7aW5mb30ke2ltcG9ydE1hcFNlY3Rpb259XG4keyc9Jy5yZXBlYXQoNDApfVxuR2VuZXJhdGVkOiAke25ldyBEYXRlKCkudG9JU09TdHJpbmcoKX1gO1xuXG4gICAgdHJ5IHtcbiAgICAgIGF3YWl0IG5hdmlnYXRvci5jbGlwYm9hcmQud3JpdGVUZXh0KGRlYnVnVGV4dCk7XG4gICAgICBjb25zdCBjb3B5QnV0dG9uID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuZGVidWctY29weScpO1xuICAgICAgaWYgKGNvcHlCdXR0b24pIHtcbiAgICAgICAgY29uc3Qgb3JpZ2luYWxUZXh0ID0gY29weUJ1dHRvbi50ZXh0Q29udGVudDtcbiAgICAgICAgY29weUJ1dHRvbi50ZXh0Q29udGVudCA9ICdDb3BpZWQhJztcblxuICAgICAgICBpZiAodGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICAgIH1cblxuICAgICAgICB0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICBpZiAodGhpcy5pc0Nvbm5lY3RlZCAmJiBjb3B5QnV0dG9uLnBhcmVudE5vZGUpIHtcbiAgICAgICAgICAgIGNvcHlCdXR0b24udGV4dENvbnRlbnQgPSBvcmlnaW5hbFRleHQ7XG4gICAgICAgICAgfVxuICAgICAgICAgIHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgICAgIH0sIDIwMDApO1xuICAgICAgfVxuICAgIH0gY2F0Y2ggKGVycikge1xuICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBjb3B5IGRlYnVnIGluZm86JywgZXJyKTtcbiAgICB9XG4gIH1cblxuICAjcmVuZGVyTG9ncyhsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT4pIHtcbiAgICBpZiAoIXRoaXMuI2xvZ0NvbnRhaW5lcikgcmV0dXJuO1xuXG4gICAgY29uc3QgbG9nRWxlbWVudHMgPSBsb2dzLm1hcChsb2cgPT4ge1xuICAgICAgY29uc3QgZnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jbG9nRW50cnlUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICBjb25zdCBkYXRlID0gbmV3IERhdGUobG9nLmRhdGUpO1xuICAgICAgY29uc3QgdGltZSA9IGRhdGUudG9Mb2NhbGVUaW1lU3RyaW5nKCk7XG5cbiAgICAgIGNvbnN0IGNvbnRhaW5lciA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiY29udGFpbmVyXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICBjb250YWluZXIuY2xhc3NMaXN0LmFkZChsb2cudHlwZSk7XG4gICAgICBjb250YWluZXIuc2V0QXR0cmlidXRlKCdkYXRhLWxvZy1pZCcsIGxvZy5kYXRlKTtcblxuICAgICAgY29uc3QgdHlwZUxhYmVsID0gdGhpcy4jZ2V0TG9nQmFkZ2UobG9nLnR5cGUpO1xuICAgICAgY29uc3Qgc2VhcmNoVGV4dCA9IGAke3R5cGVMYWJlbH0gJHt0aW1lfSAke2xvZy5tZXNzYWdlfWAudG9Mb3dlckNhc2UoKTtcbiAgICAgIGNvbnN0IHRleHRNYXRjaCA9ICF0aGlzLiNsb2dzRmlsdGVyVmFsdWUgfHwgc2VhcmNoVGV4dC5pbmNsdWRlcyh0aGlzLiNsb2dzRmlsdGVyVmFsdWUpO1xuXG4gICAgICBjb25zdCBsb2dUeXBlRm9yRmlsdGVyID0gbG9nLnR5cGUgPT09ICd3YXJuaW5nJyA/ICd3YXJuJyA6IGxvZy50eXBlO1xuICAgICAgY29uc3QgbGV2ZWxNYXRjaCA9IHRoaXMuI2xvZ0xldmVsRmlsdGVycy5oYXMobG9nVHlwZUZvckZpbHRlcik7XG5cbiAgICAgIGlmICghKHRleHRNYXRjaCAmJiBsZXZlbE1hdGNoKSkge1xuICAgICAgICBjb250YWluZXIuc2V0QXR0cmlidXRlKCdoaWRkZW4nLCAnJyk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGxhYmVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsYWJlbFwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgICAgbGFiZWwudGV4dENvbnRlbnQgPSB0aGlzLiNnZXRMb2dCYWRnZShsb2cudHlwZSk7XG4gICAgICB0aGlzLiNhcHBseUxvZ0xhYmVsQXR0cnMobGFiZWwsIGxvZy50eXBlKTtcblxuICAgICAgY29uc3QgdGltZUVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0aW1lXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICB0aW1lRWwuc2V0QXR0cmlidXRlKCdkYXRldGltZScsIGxvZy5kYXRlKTtcbiAgICAgIHRpbWVFbC50ZXh0Q29udGVudCA9IHRpbWU7XG5cbiAgICAgIChmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cIm1lc3NhZ2VcIl0nKSBhcyBIVE1MRWxlbWVudCkudGV4dENvbnRlbnQgPSBsb2cubWVzc2FnZTtcblxuICAgICAgcmV0dXJuIGZyYWdtZW50O1xuICAgIH0pO1xuXG4gICAgaWYgKCF0aGlzLiNpbml0aWFsTG9nc0ZldGNoZWQpIHtcbiAgICAgIHRoaXMuI2xvZ0NvbnRhaW5lci5yZXBsYWNlQ2hpbGRyZW4oLi4ubG9nRWxlbWVudHMpO1xuICAgICAgdGhpcy4jaW5pdGlhbExvZ3NGZXRjaGVkID0gdHJ1ZTtcblxuICAgICAgaWYgKHRoaXMuI2RyYXdlck9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKTtcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgdGhpcy4jbG9nQ29udGFpbmVyLmFwcGVuZCguLi5sb2dFbGVtZW50cyk7XG5cbiAgICAgIGlmICh0aGlzLiNkcmF3ZXJPcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExhdGVzdEludG9WaWV3KCk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI2dldExvZ0JhZGdlKHR5cGU6IHN0cmluZyk6IHN0cmluZyB7XG4gICAgc3dpdGNoICh0eXBlKSB7XG4gICAgICBjYXNlICdpbmZvJzogcmV0dXJuICdJbmZvJztcbiAgICAgIGNhc2UgJ3dhcm5pbmcnOiByZXR1cm4gJ1dhcm4nO1xuICAgICAgY2FzZSAnZXJyb3InOiByZXR1cm4gJ0Vycm9yJztcbiAgICAgIGNhc2UgJ2RlYnVnJzogcmV0dXJuICdEZWJ1Zyc7XG4gICAgICBkZWZhdWx0OiByZXR1cm4gdHlwZS50b1VwcGVyQ2FzZSgpO1xuICAgIH1cbiAgfVxuXG4gICNhcHBseUxvZ0xhYmVsQXR0cnMobGFiZWw6IEhUTUxFbGVtZW50LCB0eXBlOiBzdHJpbmcpIHtcbiAgICBzd2l0Y2ggKHR5cGUpIHtcbiAgICAgIGNhc2UgJ2luZm8nOlxuICAgICAgICByZXR1cm4gbGFiZWwuc2V0QXR0cmlidXRlKCdzdGF0dXMnLCAnaW5mbycpO1xuICAgICAgY2FzZSAnd2FybmluZyc6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICd3YXJuaW5nJyk7XG4gICAgICBjYXNlICdlcnJvcic6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ3N0YXR1cycsICdkYW5nZXInKTtcbiAgICAgIGNhc2UgJ2RlYnVnJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnY29sb3InLCAncHVycGxlJyk7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICBsYWJlbC5zZXRBdHRyaWJ1dGUoJ2NvbG9yJywgJ2dyZXknKTtcbiAgICB9XG4gIH1cblxuICAjc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICBjb25zdCBsYXN0TG9nID0gdGhpcy4jbG9nQ29udGFpbmVyIS5sYXN0RWxlbWVudENoaWxkO1xuICAgICAgaWYgKGxhc3RMb2cpIHtcbiAgICAgICAgbGFzdExvZy5zY3JvbGxJbnRvVmlldyh7IGJlaGF2aW9yOiAnYXV0bycsIGJsb2NrOiAnZW5kJyB9KTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNzY3JvbGxMb2dzVG9Cb3R0b20oKSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGlmICh0aGlzLiNpc0luaXRpYWxMb2FkKSB7XG4gICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgIH0gZWxzZSB7XG4gICAgICBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKTtcbiAgICAgIH0sIDM1MCk7XG4gICAgfVxuICB9XG5cbiAgI21pZ3JhdGVGcm9tTG9jYWxTdG9yYWdlSWZOZWVkZWQoKSB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IGhhc0xvY2FsU3RvcmFnZSA9XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtY29sb3Itc2NoZW1lJykgIT09IG51bGwgfHxcbiAgICAgICAgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1kcmF3ZXItb3BlbicpICE9PSBudWxsIHx8XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtZHJhd2VyLWhlaWdodCcpICE9PSBudWxsIHx8XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtYWN0aXZlLXRhYicpICE9PSBudWxsO1xuXG4gICAgICBpZiAoaGFzTG9jYWxTdG9yYWdlKSB7XG4gICAgICAgIGNvbnN0IG1pZ3JhdGVkID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1taWdyYXRlZC10by1jb29raWVzJyk7XG4gICAgICAgIGlmICghbWlncmF0ZWQpIHtcbiAgICAgICAgICBTdGF0ZVBlcnNpc3RlbmNlLm1pZ3JhdGVGcm9tTG9jYWxTdG9yYWdlKCk7XG4gICAgICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oJ2NlbS1zZXJ2ZS1taWdyYXRlZC10by1jb29raWVzJywgJ3RydWUnKTtcbiAgICAgICAgICBzZXRUaW1lb3V0KCgpID0+IHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQoKSwgMTAwKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIC8vIGxvY2FsU3RvcmFnZSBub3QgYXZhaWxhYmxlLCBza2lwIG1pZ3JhdGlvblxuICAgIH1cbiAgfVxuXG4gICNzZXR1cENvbG9yU2NoZW1lVG9nZ2xlKCkge1xuICAgIGNvbnN0IHRvZ2dsZUdyb3VwID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCcuY29sb3Itc2NoZW1lLXRvZ2dsZScpO1xuICAgIGlmICghdG9nZ2xlR3JvdXApIHJldHVybjtcblxuICAgIGNvbnN0IHN0YXRlID0gU3RhdGVQZXJzaXN0ZW5jZS5nZXRTdGF0ZSgpO1xuXG4gICAgdGhpcy4jYXBwbHlDb2xvclNjaGVtZShzdGF0ZS5jb2xvclNjaGVtZSk7XG5cbiAgICBjb25zdCBpdGVtcyA9IHRvZ2dsZUdyb3VwLnF1ZXJ5U2VsZWN0b3JBbGwoJ3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtJyk7XG4gICAgaXRlbXMuZm9yRWFjaChpdGVtID0+IHtcbiAgICAgIGlmICgoaXRlbSBhcyBhbnkpLnZhbHVlID09PSBzdGF0ZS5jb2xvclNjaGVtZSkge1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnc2VsZWN0ZWQnLCAnJyk7XG4gICAgICB9XG4gICAgfSk7XG5cbiAgICB0b2dnbGVHcm91cC5hZGRFdmVudExpc3RlbmVyKCdwZi12Ni10b2dnbGUtZ3JvdXAtY2hhbmdlJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICBjb25zdCBzY2hlbWUgPSAoZSBhcyBhbnkpLnZhbHVlO1xuICAgICAgdGhpcy4jYXBwbHlDb2xvclNjaGVtZShzY2hlbWUpO1xuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7IGNvbG9yU2NoZW1lOiBzY2hlbWUgfSk7XG4gICAgfSk7XG4gIH1cblxuICAjYXBwbHlDb2xvclNjaGVtZShzY2hlbWU6IHN0cmluZykge1xuICAgIGNvbnN0IGJvZHkgPSBkb2N1bWVudC5ib2R5O1xuXG4gICAgc3dpdGNoIChzY2hlbWUpIHtcbiAgICAgIGNhc2UgJ2xpZ2h0JzpcbiAgICAgICAgYm9keS5zdHlsZS5jb2xvclNjaGVtZSA9ICdsaWdodCc7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnZGFyayc6XG4gICAgICAgIGJvZHkuc3R5bGUuY29sb3JTY2hlbWUgPSAnZGFyayc7XG4gICAgICAgIGJyZWFrO1xuICAgICAgY2FzZSAnc3lzdGVtJzpcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIGJvZHkuc3R5bGUuY29sb3JTY2hlbWUgPSAnbGlnaHQgZGFyayc7XG4gICAgICAgIGJyZWFrO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cEtub2JDb29yZGluYXRpb24oKSB7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOmF0dHJpYnV0ZS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICB9XG5cbiAgI29uS25vYkNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB0YXJnZXQgPSB0aGlzLiNnZXRLbm9iVGFyZ2V0KGV2ZW50KTtcbiAgICBpZiAoIXRhcmdldCkge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gQ291bGQgbm90IGZpbmQga25vYiB0YXJnZXQgaW5mbyBpbiBldmVudCBwYXRoJyk7XG4gICAgICByZXR1cm47XG4gICAgfVxuXG4gICAgY29uc3QgeyB0YWdOYW1lLCBpbnN0YW5jZUluZGV4IH0gPSB0YXJnZXQ7XG5cbiAgICBjb25zdCBkZW1vID0gdGhpcy5kZW1vO1xuICAgIGlmICghZGVtbykgcmV0dXJuO1xuXG4gICAgY29uc3Qga25vYlR5cGUgPSB0aGlzLiNnZXRLbm9iVHlwZUZyb21FdmVudChldmVudCk7XG5cbiAgICBjb25zdCBzdWNjZXNzID0gKGRlbW8gYXMgYW55KS5hcHBseUtub2JDaGFuZ2UoXG4gICAgICBrbm9iVHlwZSxcbiAgICAgIChldmVudCBhcyBhbnkpLm5hbWUsXG4gICAgICAoZXZlbnQgYXMgYW55KS52YWx1ZSxcbiAgICAgIHRhZ05hbWUsXG4gICAgICBpbnN0YW5jZUluZGV4XG4gICAgKTtcblxuICAgIGlmICghc3VjY2Vzcykge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGFwcGx5IGtub2IgY2hhbmdlOicsIHtcbiAgICAgICAgdHlwZToga25vYlR5cGUsXG4gICAgICAgIG5hbWU6IChldmVudCBhcyBhbnkpLm5hbWUsXG4gICAgICAgIHRhZ05hbWUsXG4gICAgICAgIGluc3RhbmNlSW5kZXhcbiAgICAgIH0pO1xuICAgIH1cbiAgfTtcblxuICAjb25Lbm9iQ2xlYXIgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gdGhpcy4jZ2V0S25vYlRhcmdldChldmVudCk7XG4gICAgaWYgKCF0YXJnZXQpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIENvdWxkIG5vdCBmaW5kIGtub2IgdGFyZ2V0IGluZm8gaW4gZXZlbnQgcGF0aCcpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IHsgdGFnTmFtZSwgaW5zdGFuY2VJbmRleCB9ID0gdGFyZ2V0O1xuXG4gICAgY29uc3QgZGVtbyA9IHRoaXMuZGVtbztcbiAgICBpZiAoIWRlbW8pIHJldHVybjtcblxuICAgIGNvbnN0IGtub2JUeXBlID0gdGhpcy4jZ2V0S25vYlR5cGVGcm9tQ2xlYXJFdmVudChldmVudCk7XG4gICAgY29uc3QgY2xlYXJWYWx1ZSA9IGtub2JUeXBlID09PSAncHJvcGVydHknID8gdW5kZWZpbmVkIDogJyc7XG5cbiAgICBjb25zdCBzdWNjZXNzID0gKGRlbW8gYXMgYW55KS5hcHBseUtub2JDaGFuZ2UoXG4gICAgICBrbm9iVHlwZSxcbiAgICAgIChldmVudCBhcyBhbnkpLm5hbWUsXG4gICAgICBjbGVhclZhbHVlLFxuICAgICAgdGFnTmFtZSxcbiAgICAgIGluc3RhbmNlSW5kZXhcbiAgICApO1xuXG4gICAgaWYgKCFzdWNjZXNzKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY2xlYXIga25vYjonLCB7XG4gICAgICAgIHR5cGU6IGtub2JUeXBlLFxuICAgICAgICBuYW1lOiAoZXZlbnQgYXMgYW55KS5uYW1lLFxuICAgICAgICB0YWdOYW1lLFxuICAgICAgICBpbnN0YW5jZUluZGV4XG4gICAgICB9KTtcbiAgICB9XG4gIH07XG5cbiAgI2dldEtub2JUYXJnZXQoZXZlbnQ6IEV2ZW50KTogeyB0YWdOYW1lOiBzdHJpbmc7IGluc3RhbmNlSW5kZXg6IG51bWJlciB9IHwgbnVsbCB7XG4gICAgY29uc3QgZGVmYXVsdFRhZ05hbWUgPSB0aGlzLnByaW1hcnlUYWdOYW1lIHx8ICcnO1xuXG4gICAgaWYgKGV2ZW50LmNvbXBvc2VkUGF0aCkge1xuICAgICAgZm9yIChjb25zdCBlbGVtZW50IG9mIGV2ZW50LmNvbXBvc2VkUGF0aCgpKSB7XG4gICAgICAgIGlmICghKGVsZW1lbnQgaW5zdGFuY2VvZiBFbGVtZW50KSkgY29udGludWU7XG5cbiAgICAgICAgaWYgKChlbGVtZW50IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0Py5pc0VsZW1lbnRLbm9iID09PSAndHJ1ZScpIHtcbiAgICAgICAgICBjb25zdCB0YWdOYW1lID0gKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQudGFnTmFtZSB8fCBkZWZhdWx0VGFnTmFtZTtcbiAgICAgICAgICBsZXQgaW5zdGFuY2VJbmRleCA9IE51bWJlci5wYXJzZUludCgoZWxlbWVudCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5pbnN0YW5jZUluZGV4ID8/ICcnLCAxMCk7XG4gICAgICAgICAgaWYgKE51bWJlci5pc05hTihpbnN0YW5jZUluZGV4KSkgaW5zdGFuY2VJbmRleCA9IDA7XG4gICAgICAgICAgcmV0dXJuIHsgdGFnTmFtZSwgaW5zdGFuY2VJbmRleCB9O1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIHsgdGFnTmFtZTogZGVmYXVsdFRhZ05hbWUsIGluc3RhbmNlSW5kZXg6IDAgfTtcbiAgfVxuXG4gICNnZXRLbm9iVHlwZUZyb21FdmVudChldmVudDogRXZlbnQpOiBzdHJpbmcge1xuICAgIHN3aXRjaCAoZXZlbnQudHlwZSkge1xuICAgICAgY2FzZSAna25vYjphdHRyaWJ1dGUtY2hhbmdlJzpcbiAgICAgICAgcmV0dXJuICdhdHRyaWJ1dGUnO1xuICAgICAgY2FzZSAna25vYjpwcm9wZXJ0eS1jaGFuZ2UnOlxuICAgICAgICByZXR1cm4gJ3Byb3BlcnR5JztcbiAgICAgIGNhc2UgJ2tub2I6Y3NzLXByb3BlcnR5LWNoYW5nZSc6XG4gICAgICAgIHJldHVybiAnY3NzLXByb3BlcnR5JztcbiAgICAgIGRlZmF1bHQ6XG4gICAgICAgIHJldHVybiAndW5rbm93bic7XG4gICAgfVxuICB9XG5cbiAgI2dldEtub2JUeXBlRnJvbUNsZWFyRXZlbnQoZXZlbnQ6IEV2ZW50KTogc3RyaW5nIHtcbiAgICBzd2l0Y2ggKGV2ZW50LnR5cGUpIHtcbiAgICAgIGNhc2UgJ2tub2I6YXR0cmlidXRlLWNsZWFyJzpcbiAgICAgICAgcmV0dXJuICdhdHRyaWJ1dGUnO1xuICAgICAgY2FzZSAna25vYjpwcm9wZXJ0eS1jbGVhcic6XG4gICAgICAgIHJldHVybiAncHJvcGVydHknO1xuICAgICAgY2FzZSAna25vYjpjc3MtcHJvcGVydHktY2xlYXInOlxuICAgICAgICByZXR1cm4gJ2Nzcy1wcm9wZXJ0eSc7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICByZXR1cm4gJ3Vua25vd24nO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cFRyZWVTdGF0ZVBlcnNpc3RlbmNlKCkge1xuICAgIHRoaXMuI2hhbmRsZVRyZWVFeHBhbmQgPSAoZTogRXZlbnQpID0+IHtcbiAgICAgIGlmICgoZS50YXJnZXQgYXMgRWxlbWVudCk/LnRhZ05hbWUgIT09ICdQRi1WNi1UUkVFLUlURU0nKSByZXR1cm47XG5cbiAgICAgIGNvbnN0IG5vZGVJZCA9IHRoaXMuI2dldFRyZWVOb2RlSWQoZS50YXJnZXQgYXMgRWxlbWVudCk7XG4gICAgICBjb25zdCB0cmVlU3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFRyZWVTdGF0ZSgpO1xuICAgICAgaWYgKCF0cmVlU3RhdGUuZXhwYW5kZWQuaW5jbHVkZXMobm9kZUlkKSkge1xuICAgICAgICB0cmVlU3RhdGUuZXhwYW5kZWQucHVzaChub2RlSWQpO1xuICAgICAgICBTdGF0ZVBlcnNpc3RlbmNlLnNldFRyZWVTdGF0ZSh0cmVlU3RhdGUpO1xuICAgICAgfVxuICAgIH07XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdleHBhbmQnLCB0aGlzLiNoYW5kbGVUcmVlRXhwYW5kKTtcblxuICAgIHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSA9IChlOiBFdmVudCkgPT4ge1xuICAgICAgaWYgKChlLnRhcmdldCBhcyBFbGVtZW50KT8udGFnTmFtZSAhPT0gJ1BGLVY2LVRSRUUtSVRFTScpIHJldHVybjtcblxuICAgICAgY29uc3Qgbm9kZUlkID0gdGhpcy4jZ2V0VHJlZU5vZGVJZChlLnRhcmdldCBhcyBFbGVtZW50KTtcbiAgICAgIGNvbnN0IHRyZWVTdGF0ZSA9IFN0YXRlUGVyc2lzdGVuY2UuZ2V0VHJlZVN0YXRlKCk7XG4gICAgICBjb25zdCBpbmRleCA9IHRyZWVTdGF0ZS5leHBhbmRlZC5pbmRleE9mKG5vZGVJZCk7XG4gICAgICBpZiAoaW5kZXggPiAtMSkge1xuICAgICAgICB0cmVlU3RhdGUuZXhwYW5kZWQuc3BsaWNlKGluZGV4LCAxKTtcbiAgICAgICAgU3RhdGVQZXJzaXN0ZW5jZS5zZXRUcmVlU3RhdGUodHJlZVN0YXRlKTtcbiAgICAgIH1cbiAgICB9O1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignY29sbGFwc2UnLCB0aGlzLiNoYW5kbGVUcmVlQ29sbGFwc2UpO1xuXG4gICAgdGhpcy4jaGFuZGxlVHJlZVNlbGVjdCA9IChlOiBFdmVudCkgPT4ge1xuICAgICAgaWYgKChlLnRhcmdldCBhcyBFbGVtZW50KT8udGFnTmFtZSAhPT0gJ1BGLVY2LVRSRUUtSVRFTScpIHJldHVybjtcblxuICAgICAgY29uc3Qgbm9kZUlkID0gdGhpcy4jZ2V0VHJlZU5vZGVJZChlLnRhcmdldCBhcyBFbGVtZW50KTtcbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlVHJlZVN0YXRlKHsgc2VsZWN0ZWQ6IG5vZGVJZCB9KTtcbiAgICB9O1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlVHJlZVNlbGVjdCk7XG5cbiAgICB0aGlzLiNhcHBseVRyZWVTdGF0ZSgpO1xuICB9XG5cbiAgI2FwcGx5VHJlZVN0YXRlKCkge1xuICAgIGNvbnN0IHRyZWVTdGF0ZSA9IFN0YXRlUGVyc2lzdGVuY2UuZ2V0VHJlZVN0YXRlKCk7XG5cbiAgICBmb3IgKGNvbnN0IG5vZGVJZCBvZiB0cmVlU3RhdGUuZXhwYW5kZWQpIHtcbiAgICAgIGNvbnN0IHRyZWVJdGVtID0gdGhpcy4jZmluZFRyZWVJdGVtQnlJZChub2RlSWQpO1xuICAgICAgaWYgKHRyZWVJdGVtICYmICF0cmVlSXRlbS5oYXNBdHRyaWJ1dGUoJ2V4cGFuZGVkJykpIHtcbiAgICAgICAgdHJlZUl0ZW0uc2V0QXR0cmlidXRlKCdleHBhbmRlZCcsICcnKTtcbiAgICAgIH1cbiAgICB9XG5cbiAgICBpZiAodHJlZVN0YXRlLnNlbGVjdGVkKSB7XG4gICAgICBjb25zdCB0cmVlSXRlbSA9IHRoaXMuI2ZpbmRUcmVlSXRlbUJ5SWQodHJlZVN0YXRlLnNlbGVjdGVkKTtcbiAgICAgIGlmICh0cmVlSXRlbSAmJiAhdHJlZUl0ZW0uaGFzQXR0cmlidXRlKCdjdXJyZW50JykpIHtcbiAgICAgICAgdHJlZUl0ZW0uc2V0QXR0cmlidXRlKCdjdXJyZW50JywgJycpO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNzZXR1cFNpZGViYXJTdGF0ZVBlcnNpc3RlbmNlKCkge1xuICAgIGNvbnN0IHBhZ2UgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ3BmLXY2LXBhZ2UnKTtcblxuICAgIGlmICghcGFnZSkgcmV0dXJuO1xuXG4gICAgcGFnZS5hZGRFdmVudExpc3RlbmVyKCdzaWRlYmFyLXRvZ2dsZScsIChldmVudDogRXZlbnQpID0+IHtcbiAgICAgIGNvbnN0IGNvbGxhcHNlZCA9ICEoZXZlbnQgYXMgYW55KS5leHBhbmRlZDtcblxuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIHNpZGViYXI6IHsgY29sbGFwc2VkIH1cbiAgICAgIH0pO1xuICAgIH0pO1xuICB9XG5cbiAgI2ZpbmRUcmVlSXRlbUJ5SWQobm9kZUlkOiBzdHJpbmcpOiBFbGVtZW50IHwgbnVsbCB7XG4gICAgY29uc3QgcGFydHMgPSBub2RlSWQuc3BsaXQoJzonKTtcbiAgICBjb25zdCBbdHlwZSwgbW9kdWxlUGF0aCwgdGFnTmFtZSwgbmFtZV0gPSBwYXJ0cztcblxuICAgIGxldCBhdHRyU3VmZml4ID0gJyc7XG4gICAgaWYgKHRhZ05hbWUpIHtcbiAgICAgIGF0dHJTdWZmaXggKz0gYFtkYXRhLXRhZy1uYW1lPVwiJHtDU1MuZXNjYXBlKHRhZ05hbWUpfVwiXWA7XG4gICAgfVxuICAgIGlmIChuYW1lKSB7XG4gICAgICBhdHRyU3VmZml4ICs9IGBbZGF0YS1uYW1lPVwiJHtDU1MuZXNjYXBlKG5hbWUpfVwiXWA7XG4gICAgfVxuXG4gICAgbGV0IHNlbGVjdG9yID0gYHBmLXY2LXRyZWUtaXRlbVtkYXRhLXR5cGU9XCIke0NTUy5lc2NhcGUodHlwZSl9XCJdYDtcbiAgICBpZiAobW9kdWxlUGF0aCkge1xuICAgICAgY29uc3QgZXNjYXBlZE1vZHVsZVBhdGggPSBDU1MuZXNjYXBlKG1vZHVsZVBhdGgpO1xuICAgICAgY29uc3QgZXNjYXBlZFR5cGUgPSBDU1MuZXNjYXBlKHR5cGUpO1xuICAgICAgY29uc3Qgc2VsZWN0b3IxID0gYHBmLXY2LXRyZWUtaXRlbVtkYXRhLXR5cGU9XCIke2VzY2FwZWRUeXBlfVwiXVtkYXRhLW1vZHVsZS1wYXRoPVwiJHtlc2NhcGVkTW9kdWxlUGF0aH1cIl0ke2F0dHJTdWZmaXh9YDtcbiAgICAgIGNvbnN0IHNlbGVjdG9yMiA9IGBwZi12Ni10cmVlLWl0ZW1bZGF0YS10eXBlPVwiJHtlc2NhcGVkVHlwZX1cIl1bZGF0YS1wYXRoPVwiJHtlc2NhcGVkTW9kdWxlUGF0aH1cIl0ke2F0dHJTdWZmaXh9YDtcbiAgICAgIHNlbGVjdG9yID0gYCR7c2VsZWN0b3IxfSwgJHtzZWxlY3RvcjJ9YDtcbiAgICB9IGVsc2Uge1xuICAgICAgc2VsZWN0b3IgKz0gYXR0clN1ZmZpeDtcbiAgICB9XG5cbiAgICByZXR1cm4gdGhpcy5xdWVyeVNlbGVjdG9yKHNlbGVjdG9yKTtcbiAgfVxuXG4gICNnZXRUcmVlTm9kZUlkKHRyZWVJdGVtOiBFbGVtZW50KTogc3RyaW5nIHtcbiAgICBjb25zdCB0eXBlID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXR5cGUnKTtcbiAgICBjb25zdCBtb2R1bGVQYXRoID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLW1vZHVsZS1wYXRoJykgfHwgdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXBhdGgnKTtcbiAgICBjb25zdCB0YWdOYW1lID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLXRhZy1uYW1lJyk7XG4gICAgY29uc3QgbmFtZSA9IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS1uYW1lJyk7XG4gICAgY29uc3QgY2F0ZWdvcnkgPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtY2F0ZWdvcnknKTtcblxuICAgIGNvbnN0IHBhcnRzID0gW3R5cGVdO1xuICAgIGlmIChtb2R1bGVQYXRoKSBwYXJ0cy5wdXNoKG1vZHVsZVBhdGgpO1xuICAgIGlmICh0YWdOYW1lKSBwYXJ0cy5wdXNoKHRhZ05hbWUpO1xuICAgIGlmIChjYXRlZ29yeSkge1xuICAgICAgcGFydHMucHVzaChjYXRlZ29yeSk7XG4gICAgfSBlbHNlIGlmIChuYW1lKSB7XG4gICAgICBwYXJ0cy5wdXNoKG5hbWUpO1xuICAgIH1cblxuICAgIHJldHVybiBwYXJ0cy5qb2luKCc6Jyk7XG4gIH1cblxuICAvLyBFdmVudCBEaXNjb3ZlcnkgJiBDYXB0dXJlIE1ldGhvZHNcblxuICBhc3luYyAjZGlzY292ZXJFbGVtZW50RXZlbnRzKCk6IFByb21pc2U8TWFwPHN0cmluZywgRXZlbnRJbmZvPj4ge1xuICAgIHRyeSB7XG4gICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IGZldGNoKCcvY3VzdG9tLWVsZW1lbnRzLmpzb24nKTtcbiAgICAgIGlmICghcmVzcG9uc2Uub2spIHtcbiAgICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gTm8gbWFuaWZlc3QgYXZhaWxhYmxlIGZvciBldmVudCBkaXNjb3ZlcnknKTtcbiAgICAgICAgcmV0dXJuIG5ldyBNYXAoKTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgbWFuaWZlc3QgPSBhd2FpdCByZXNwb25zZS5qc29uKCkgYXMgTWFuaWZlc3Q7XG4gICAgICB0aGlzLiNtYW5pZmVzdCA9IG1hbmlmZXN0O1xuXG4gICAgICBjb25zdCBldmVudE1hcCA9IG5ldyBNYXA8c3RyaW5nLCBFdmVudEluZm8+KCk7XG5cbiAgICAgIGZvciAoY29uc3QgbW9kdWxlIG9mIG1hbmlmZXN0Lm1vZHVsZXMgfHwgW10pIHtcbiAgICAgICAgZm9yIChjb25zdCBkZWNsYXJhdGlvbiBvZiBtb2R1bGUuZGVjbGFyYXRpb25zIHx8IFtdKSB7XG4gICAgICAgICAgaWYgKGRlY2xhcmF0aW9uLmN1c3RvbUVsZW1lbnQgJiYgZGVjbGFyYXRpb24udGFnTmFtZSkge1xuICAgICAgICAgICAgY29uc3QgdGFnTmFtZSA9IGRlY2xhcmF0aW9uLnRhZ05hbWU7XG4gICAgICAgICAgICBjb25zdCBldmVudHMgPSBkZWNsYXJhdGlvbi5ldmVudHMgfHwgW107XG5cbiAgICAgICAgICAgIGlmIChldmVudHMubGVuZ3RoID4gMCkge1xuICAgICAgICAgICAgICBjb25zdCBldmVudE5hbWVzID0gbmV3IFNldChldmVudHMubWFwKGUgPT4gZS5uYW1lKSk7XG4gICAgICAgICAgICAgIGV2ZW50TWFwLnNldCh0YWdOYW1lLCB7XG4gICAgICAgICAgICAgICAgZXZlbnROYW1lcyxcbiAgICAgICAgICAgICAgICBldmVudHM6IGV2ZW50c1xuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgcmV0dXJuIGV2ZW50TWFwO1xuICAgIH0gY2F0Y2ggKGVycm9yKSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBFcnJvciBsb2FkaW5nIG1hbmlmZXN0IGZvciBldmVudCBkaXNjb3Zlcnk6JywgZXJyb3IpO1xuICAgICAgcmV0dXJuIG5ldyBNYXAoKTtcbiAgICB9XG4gIH1cblxuICBhc3luYyAjc2V0dXBFdmVudENhcHR1cmUoKSB7XG4gICAgdGhpcy4jZWxlbWVudEV2ZW50TWFwID0gYXdhaXQgdGhpcy4jZGlzY292ZXJFbGVtZW50RXZlbnRzKCk7XG5cbiAgICBpZiAodGhpcy4jZWxlbWVudEV2ZW50TWFwLnNpemUgPT09IDApIHJldHVybjtcblxuICAgIHRoaXMuI2F0dGFjaEV2ZW50TGlzdGVuZXJzKCk7XG4gICAgdGhpcy4jdXBkYXRlRXZlbnRGaWx0ZXJzKCk7XG4gICAgdGhpcy4jb2JzZXJ2ZURlbW9NdXRhdGlvbnMoKTtcbiAgfVxuXG4gICNhdHRhY2hFdmVudExpc3RlbmVycygpIHtcbiAgICBjb25zdCBkZW1vID0gdGhpcy5kZW1vO1xuICAgIGlmICghZGVtbykgcmV0dXJuO1xuXG4gICAgY29uc3Qgcm9vdCA9IGRlbW8uc2hhZG93Um9vdCA/PyBkZW1vO1xuXG4gICAgZm9yIChjb25zdCBbdGFnTmFtZSwgZXZlbnRJbmZvXSBvZiB0aGlzLiNlbGVtZW50RXZlbnRNYXAhKSB7XG4gICAgICBjb25zdCBlbGVtZW50cyA9IHJvb3QucXVlcnlTZWxlY3RvckFsbCh0YWdOYW1lKTtcblxuICAgICAgZm9yIChjb25zdCBlbGVtZW50IG9mIGVsZW1lbnRzKSB7XG4gICAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGV2ZW50SW5mby5ldmVudE5hbWVzKSB7XG4gICAgICAgICAgZWxlbWVudC5hZGRFdmVudExpc3RlbmVyKGV2ZW50TmFtZSwgdGhpcy4jaGFuZGxlRWxlbWVudEV2ZW50LCB7IGNhcHR1cmU6IHRydWUgfSk7XG4gICAgICAgIH1cbiAgICAgICAgKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuY2VtRXZlbnRzQXR0YWNoZWQgPSAndHJ1ZSc7XG4gICAgICAgIHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5hZGQodGFnTmFtZSk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgI29ic2VydmVEZW1vTXV0YXRpb25zKCkge1xuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCByb290ID0gZGVtby5zaGFkb3dSb290ID8/IGRlbW87XG5cbiAgICB0aGlzLiNvYnNlcnZlci5vYnNlcnZlKHJvb3QsIHtcbiAgICAgIGNoaWxkTGlzdDogdHJ1ZSxcbiAgICAgIHN1YnRyZWU6IHRydWVcbiAgICB9KTtcbiAgfVxuXG4gICNoYW5kbGVFbGVtZW50RXZlbnQgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgZWxlbWVudCA9IGV2ZW50LmN1cnJlbnRUYXJnZXQ7XG4gICAgaWYgKCEoZWxlbWVudCBpbnN0YW5jZW9mIEhUTUxFbGVtZW50KSkgcmV0dXJuO1xuXG4gICAgY29uc3QgdGFnTmFtZSA9IGVsZW1lbnQudGFnTmFtZS50b0xvd2VyQ2FzZSgpO1xuICAgIGNvbnN0IGV2ZW50SW5mbyA9IHRoaXMuI2VsZW1lbnRFdmVudE1hcD8uZ2V0KHRhZ05hbWUpO1xuXG4gICAgaWYgKCFldmVudEluZm8gfHwgIWV2ZW50SW5mby5ldmVudE5hbWVzLmhhcyhldmVudC50eXBlKSkgcmV0dXJuO1xuXG4gICAgdGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICB0aGlzLiNjYXB0dXJlRXZlbnQoZXZlbnQsIGVsZW1lbnQsIHRhZ05hbWUsIGV2ZW50SW5mbyk7XG4gIH07XG5cbiAgI2dldEV2ZW50RG9jdW1lbnRhdGlvbihtYW5pZmVzdEV2ZW50OiBFdmVudEluZm9bJ2V2ZW50cyddWzBdIHwgdW5kZWZpbmVkKSB7XG4gICAgaWYgKCFtYW5pZmVzdEV2ZW50KSB7XG4gICAgICByZXR1cm4geyBzdW1tYXJ5OiBudWxsLCBkZXNjcmlwdGlvbjogbnVsbCB9O1xuICAgIH1cblxuICAgIGxldCBzdW1tYXJ5ID0gbWFuaWZlc3RFdmVudC5zdW1tYXJ5IHx8IG51bGw7XG4gICAgbGV0IGRlc2NyaXB0aW9uID0gbWFuaWZlc3RFdmVudC5kZXNjcmlwdGlvbiB8fCBudWxsO1xuXG4gICAgaWYgKG1hbmlmZXN0RXZlbnQudHlwZT8udGV4dCAmJiB0aGlzLiNtYW5pZmVzdCkge1xuICAgICAgY29uc3QgdHlwZU5hbWUgPSBtYW5pZmVzdEV2ZW50LnR5cGUudGV4dDtcbiAgICAgIGNvbnN0IHR5cGVEZWNsYXJhdGlvbiA9IHRoaXMuI2ZpbmRUeXBlRGVjbGFyYXRpb24odHlwZU5hbWUpO1xuXG4gICAgICBpZiAodHlwZURlY2xhcmF0aW9uKSB7XG4gICAgICAgIGlmICghc3VtbWFyeSAmJiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSkge1xuICAgICAgICAgIHN1bW1hcnkgPSB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeTtcbiAgICAgICAgfSBlbHNlIGlmICh0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSAmJiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeSAhPT0gc3VtbWFyeSkge1xuICAgICAgICAgIHN1bW1hcnkgPSBzdW1tYXJ5ID8gYCR7c3VtbWFyeX1cXG5cXG5Gcm9tICR7dHlwZU5hbWV9OiAke3R5cGVEZWNsYXJhdGlvbi5zdW1tYXJ5fWAgOiB0eXBlRGVjbGFyYXRpb24uc3VtbWFyeTtcbiAgICAgICAgfVxuXG4gICAgICAgIGlmICghZGVzY3JpcHRpb24gJiYgdHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uKSB7XG4gICAgICAgICAgZGVzY3JpcHRpb24gPSB0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb247XG4gICAgICAgIH0gZWxzZSBpZiAodHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uICYmIHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbiAhPT0gZGVzY3JpcHRpb24pIHtcbiAgICAgICAgICBkZXNjcmlwdGlvbiA9IGRlc2NyaXB0aW9uID8gYCR7ZGVzY3JpcHRpb259XFxuXFxuJHt0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb259YCA6IHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbjtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiB7IHN1bW1hcnksIGRlc2NyaXB0aW9uIH07XG4gIH1cblxuICAjZmluZFR5cGVEZWNsYXJhdGlvbih0eXBlTmFtZTogc3RyaW5nKSB7XG4gICAgaWYgKCF0aGlzLiNtYW5pZmVzdCkgcmV0dXJuIG51bGw7XG5cbiAgICBmb3IgKGNvbnN0IG1vZHVsZSBvZiB0aGlzLiNtYW5pZmVzdC5tb2R1bGVzIHx8IFtdKSB7XG4gICAgICBmb3IgKGNvbnN0IGRlY2xhcmF0aW9uIG9mIG1vZHVsZS5kZWNsYXJhdGlvbnMgfHwgW10pIHtcbiAgICAgICAgaWYgKGRlY2xhcmF0aW9uLm5hbWUgPT09IHR5cGVOYW1lICYmXG4gICAgICAgICAgICAoZGVjbGFyYXRpb24ua2luZCA9PT0gJ2NsYXNzJyB8fCBkZWNsYXJhdGlvbi5raW5kID09PSAnaW50ZXJmYWNlJykpIHtcbiAgICAgICAgICByZXR1cm4gZGVjbGFyYXRpb24gYXMgeyBzdW1tYXJ5Pzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9O1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIG51bGw7XG4gIH1cblxuICAjY2FwdHVyZUV2ZW50KGV2ZW50OiBFdmVudCwgdGFyZ2V0OiBIVE1MRWxlbWVudCwgdGFnTmFtZTogc3RyaW5nLCBldmVudEluZm86IEV2ZW50SW5mbykge1xuICAgIGNvbnN0IG1hbmlmZXN0RXZlbnQgPSBldmVudEluZm8uZXZlbnRzLmZpbmQoZSA9PiBlLm5hbWUgPT09IGV2ZW50LnR5cGUpO1xuXG4gICAgY29uc3QgZXZlbnREb2NzID0gdGhpcy4jZ2V0RXZlbnREb2N1bWVudGF0aW9uKG1hbmlmZXN0RXZlbnQpO1xuXG4gICAgY29uc3QgY3VzdG9tUHJvcGVydGllcyA9IHRoaXMuI2V4dHJhY3RFdmVudFByb3BlcnRpZXMoZXZlbnQpO1xuXG4gICAgY29uc3QgZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkID0ge1xuICAgICAgaWQ6IGAke0RhdGUubm93KCl9LSR7TWF0aC5yYW5kb20oKX1gLFxuICAgICAgdGltZXN0YW1wOiBuZXcgRGF0ZSgpLFxuICAgICAgZXZlbnROYW1lOiBldmVudC50eXBlLFxuICAgICAgdGFnTmFtZTogdGFnTmFtZSxcbiAgICAgIGVsZW1lbnRJZDogdGFyZ2V0LmlkIHx8IG51bGwsXG4gICAgICBlbGVtZW50Q2xhc3M6IHRhcmdldC5jbGFzc05hbWUgfHwgbnVsbCxcbiAgICAgIGN1c3RvbVByb3BlcnRpZXM6IGN1c3RvbVByb3BlcnRpZXMsXG4gICAgICBtYW5pZmVzdFR5cGU6IG1hbmlmZXN0RXZlbnQ/LnR5cGU/LnRleHQgfHwgbnVsbCxcbiAgICAgIHN1bW1hcnk6IGV2ZW50RG9jcy5zdW1tYXJ5LFxuICAgICAgZGVzY3JpcHRpb246IGV2ZW50RG9jcy5kZXNjcmlwdGlvbixcbiAgICAgIGJ1YmJsZXM6IGV2ZW50LmJ1YmJsZXMsXG4gICAgICBjb21wb3NlZDogZXZlbnQuY29tcG9zZWQsXG4gICAgICBjYW5jZWxhYmxlOiBldmVudC5jYW5jZWxhYmxlLFxuICAgICAgZGVmYXVsdFByZXZlbnRlZDogZXZlbnQuZGVmYXVsdFByZXZlbnRlZFxuICAgIH07XG5cbiAgICB0aGlzLiNjYXB0dXJlZEV2ZW50cy5wdXNoKGV2ZW50UmVjb3JkKTtcblxuICAgIGlmICh0aGlzLiNjYXB0dXJlZEV2ZW50cy5sZW5ndGggPiB0aGlzLiNtYXhDYXB0dXJlZEV2ZW50cykge1xuICAgICAgdGhpcy4jY2FwdHVyZWRFdmVudHMuc2hpZnQoKTtcbiAgICB9XG5cbiAgICB0aGlzLiNyZW5kZXJFdmVudChldmVudFJlY29yZCk7XG4gIH1cblxuICAjZXh0cmFjdEV2ZW50UHJvcGVydGllcyhldmVudDogRXZlbnQpOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiB7XG4gICAgY29uc3QgcHJvcGVydGllczogUmVjb3JkPHN0cmluZywgdW5rbm93bj4gPSB7fTtcbiAgICBjb25zdCBldmVudFByb3RvdHlwZUtleXMgPSBuZXcgU2V0KE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKEV2ZW50LnByb3RvdHlwZSkpO1xuXG4gICAgY29uc3Qgc2VyaWFsaXplVmFsdWUgPSAodmFsdWU6IHVua25vd24pOiB1bmtub3duID0+IHtcbiAgICAgIHRyeSB7XG4gICAgICAgIHJldHVybiBKU09OLnBhcnNlKEpTT04uc3RyaW5naWZ5KHZhbHVlKSk7XG4gICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgIHRyeSB7XG4gICAgICAgICAgcmV0dXJuIFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIH0gY2F0Y2ggKHN0cmluZ0Vycikge1xuICAgICAgICAgIHJldHVybiAnW05vdCBzZXJpYWxpemFibGVdJztcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH07XG5cbiAgICBpZiAoZXZlbnQgaW5zdGFuY2VvZiBDdXN0b21FdmVudCAmJiBldmVudC5kZXRhaWwgIT09IHVuZGVmaW5lZCkge1xuICAgICAgcHJvcGVydGllcy5kZXRhaWwgPSBzZXJpYWxpemVWYWx1ZShldmVudC5kZXRhaWwpO1xuICAgIH1cblxuICAgIGZvciAoY29uc3Qga2V5IG9mIE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKGV2ZW50KSkge1xuICAgICAgaWYgKCFldmVudFByb3RvdHlwZUtleXMuaGFzKGtleSkgJiYgIWtleS5zdGFydHNXaXRoKCdfJykgJiYgIXByb3BlcnRpZXMuaGFzT3duUHJvcGVydHkoa2V5KSkge1xuICAgICAgICBwcm9wZXJ0aWVzW2tleV0gPSBzZXJpYWxpemVWYWx1ZSgoZXZlbnQgYXMgYW55KVtrZXldKTtcbiAgICAgIH1cbiAgICB9XG5cbiAgICByZXR1cm4gcHJvcGVydGllcztcbiAgfVxuXG4gICNyZW5kZXJFdmVudChldmVudFJlY29yZDogRXZlbnRSZWNvcmQpIHtcbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgY29uc3QgZnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZXZlbnRFbnRyeVRlbXBsYXRlLmNvbnRlbnQuY2xvbmVOb2RlKHRydWUpIGFzIERvY3VtZW50RnJhZ21lbnQ7XG5cbiAgICBjb25zdCB0aW1lID0gZXZlbnRSZWNvcmQudGltZXN0YW1wLnRvTG9jYWxlVGltZVN0cmluZygpO1xuXG4gICAgY29uc3QgY29udGFpbmVyID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJjb250YWluZXJcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICBjb250YWluZXIuZGF0YXNldC5ldmVudElkID0gZXZlbnRSZWNvcmQuaWQ7XG4gICAgY29udGFpbmVyLmRhdGFzZXQuZXZlbnRUeXBlID0gZXZlbnRSZWNvcmQuZXZlbnROYW1lO1xuICAgIGNvbnRhaW5lci5kYXRhc2V0LmVsZW1lbnRUeXBlID0gZXZlbnRSZWNvcmQudGFnTmFtZTtcblxuICAgIGNvbnN0IHRleHRNYXRjaCA9IHRoaXMuI2V2ZW50TWF0Y2hlc1RleHRGaWx0ZXIoZXZlbnRSZWNvcmQpO1xuICAgIGNvbnN0IHR5cGVNYXRjaCA9IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmhhcyhldmVudFJlY29yZC5ldmVudE5hbWUpO1xuICAgIGNvbnN0IGVsZW1lbnRNYXRjaCA9IHRoaXMuI2VsZW1lbnRGaWx0ZXJzLnNpemUgPT09IDAgfHwgdGhpcy4jZWxlbWVudEZpbHRlcnMuaGFzKGV2ZW50UmVjb3JkLnRhZ05hbWUpO1xuXG4gICAgaWYgKCEodGV4dE1hdGNoICYmIHR5cGVNYXRjaCAmJiBlbGVtZW50TWF0Y2gpKSB7XG4gICAgICBjb250YWluZXIuc2V0QXR0cmlidXRlKCdoaWRkZW4nLCAnJyk7XG4gICAgfVxuXG4gICAgY29uc3QgbGFiZWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxhYmVsXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgbGFiZWwudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC5ldmVudE5hbWU7XG4gICAgbGFiZWwuc2V0QXR0cmlidXRlKCdzdGF0dXMnLCAnaW5mbycpO1xuXG4gICAgY29uc3QgdGltZUVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJ0aW1lXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgdGltZUVsLnNldEF0dHJpYnV0ZSgnZGF0ZXRpbWUnLCBldmVudFJlY29yZC50aW1lc3RhbXAudG9JU09TdHJpbmcoKSk7XG4gICAgdGltZUVsLnRleHRDb250ZW50ID0gdGltZTtcblxuICAgIGNvbnN0IGVsZW1lbnRFbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZWxlbWVudFwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgIGxldCBlbGVtZW50VGV4dCA9IGA8JHtldmVudFJlY29yZC50YWdOYW1lfT5gO1xuICAgIGlmIChldmVudFJlY29yZC5lbGVtZW50SWQpIHtcbiAgICAgIGVsZW1lbnRUZXh0ICs9IGAjJHtldmVudFJlY29yZC5lbGVtZW50SWR9YDtcbiAgICB9XG4gICAgZWxlbWVudEVsLnRleHRDb250ZW50ID0gZWxlbWVudFRleHQ7XG5cbiAgICB0aGlzLiNldmVudExpc3QuYXBwZW5kKGZyYWdtZW50KTtcblxuICAgIGlmICghdGhpcy4jc2VsZWN0ZWRFdmVudElkKSB7XG4gICAgICB0aGlzLiNzZWxlY3RFdmVudChldmVudFJlY29yZC5pZCk7XG4gICAgfVxuXG4gICAgaWYgKHRoaXMuI2RyYXdlck9wZW4gJiYgdGhpcy4jaXNFdmVudHNUYWJBY3RpdmUoKSkge1xuICAgICAgdGhpcy4jc2Nyb2xsRXZlbnRzVG9Cb3R0b20oKTtcbiAgICB9XG4gIH1cblxuICAjc2VsZWN0RXZlbnQoZXZlbnRJZDogc3RyaW5nKSB7XG4gICAgY29uc3QgZXZlbnRSZWNvcmQgPSB0aGlzLiNnZXRFdmVudFJlY29yZEJ5SWQoZXZlbnRJZCk7XG4gICAgaWYgKCFldmVudFJlY29yZCkgcmV0dXJuO1xuXG4gICAgdGhpcy4jc2VsZWN0ZWRFdmVudElkID0gZXZlbnRJZDtcblxuICAgIGNvbnN0IGFsbEl0ZW1zID0gdGhpcy4jZXZlbnRMaXN0Py5xdWVyeVNlbGVjdG9yQWxsKCcuZXZlbnQtbGlzdC1pdGVtJyk7XG4gICAgYWxsSXRlbXM/LmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBpZiAoKGl0ZW0gYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuZXZlbnRJZCA9PT0gZXZlbnRJZCkge1xuICAgICAgICBpdGVtLmNsYXNzTGlzdC5hZGQoJ3NlbGVjdGVkJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCdhcmlhLXNlbGVjdGVkJywgJ3RydWUnKTtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGl0ZW0uY2xhc3NMaXN0LnJlbW92ZSgnc2VsZWN0ZWQnKTtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2FyaWEtc2VsZWN0ZWQnLCAnZmFsc2UnKTtcbiAgICAgIH1cbiAgICB9KTtcblxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEhlYWRlcikge1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIuaW5uZXJIVE1MID0gJyc7XG5cbiAgICAgIGNvbnN0IGhlYWRlckNvbnRlbnQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIGhlYWRlckNvbnRlbnQuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1oZWFkZXItY29udGVudCc7XG5cbiAgICAgIGNvbnN0IGV2ZW50TmFtZSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2gzJyk7XG4gICAgICBldmVudE5hbWUudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC5ldmVudE5hbWU7XG4gICAgICBldmVudE5hbWUuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1uYW1lJztcbiAgICAgIGhlYWRlckNvbnRlbnQuYXBwZW5kQ2hpbGQoZXZlbnROYW1lKTtcblxuICAgICAgaWYgKGV2ZW50UmVjb3JkLnN1bW1hcnkpIHtcbiAgICAgICAgY29uc3Qgc3VtbWFyeSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3AnKTtcbiAgICAgICAgc3VtbWFyeS50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLnN1bW1hcnk7XG4gICAgICAgIHN1bW1hcnkuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1zdW1tYXJ5JztcbiAgICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChzdW1tYXJ5KTtcbiAgICAgIH1cblxuICAgICAgaWYgKGV2ZW50UmVjb3JkLmRlc2NyaXB0aW9uKSB7XG4gICAgICAgIGNvbnN0IGRlc2NyaXB0aW9uID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncCcpO1xuICAgICAgICBkZXNjcmlwdGlvbi50ZXh0Q29udGVudCA9IGV2ZW50UmVjb3JkLmRlc2NyaXB0aW9uO1xuICAgICAgICBkZXNjcmlwdGlvbi5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLWRlc2NyaXB0aW9uJztcbiAgICAgICAgaGVhZGVyQ29udGVudC5hcHBlbmRDaGlsZChkZXNjcmlwdGlvbik7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IG1ldGEgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIG1ldGEuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1tZXRhJztcblxuICAgICAgY29uc3QgdGltZUVsID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGltZScpO1xuICAgICAgdGltZUVsLnNldEF0dHJpYnV0ZSgnZGF0ZXRpbWUnLCBldmVudFJlY29yZC50aW1lc3RhbXAudG9JU09TdHJpbmcoKSk7XG4gICAgICB0aW1lRWwudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC50aW1lc3RhbXAudG9Mb2NhbGVUaW1lU3RyaW5nKCk7XG4gICAgICB0aW1lRWwuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC10aW1lJztcblxuICAgICAgY29uc3QgZWxlbWVudCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgIGxldCBlbGVtZW50VGV4dCA9IGA8JHtldmVudFJlY29yZC50YWdOYW1lfT5gO1xuICAgICAgaWYgKGV2ZW50UmVjb3JkLmVsZW1lbnRJZCkge1xuICAgICAgICBlbGVtZW50VGV4dCArPSBgIyR7ZXZlbnRSZWNvcmQuZWxlbWVudElkfWA7XG4gICAgICB9XG4gICAgICBlbGVtZW50LnRleHRDb250ZW50ID0gZWxlbWVudFRleHQ7XG4gICAgICBlbGVtZW50LmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtZWxlbWVudCc7XG5cbiAgICAgIG1ldGEuYXBwZW5kQ2hpbGQodGltZUVsKTtcbiAgICAgIG1ldGEuYXBwZW5kQ2hpbGQoZWxlbWVudCk7XG5cbiAgICAgIGhlYWRlckNvbnRlbnQuYXBwZW5kQ2hpbGQobWV0YSk7XG5cbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyLmFwcGVuZENoaWxkKGhlYWRlckNvbnRlbnQpO1xuICAgIH1cblxuICAgIGlmICh0aGlzLiNldmVudERldGFpbEJvZHkpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5pbm5lckhUTUwgPSAnJztcblxuICAgICAgY29uc3QgcHJvcGVydGllc0hlYWRpbmcgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdoNCcpO1xuICAgICAgcHJvcGVydGllc0hlYWRpbmcudGV4dENvbnRlbnQgPSAnUHJvcGVydGllcyc7XG4gICAgICBwcm9wZXJ0aWVzSGVhZGluZy5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMtaGVhZGluZyc7XG5cbiAgICAgIGNvbnN0IHByb3BlcnRpZXNDb250YWluZXIgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgIHByb3BlcnRpZXNDb250YWluZXIuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1wcm9wZXJ0aWVzJztcblxuICAgICAgY29uc3QgZXZlbnRQcm9wZXJ0aWVzID0gdGhpcy4jYnVpbGRQcm9wZXJ0aWVzRm9yRGlzcGxheShldmVudFJlY29yZCk7XG4gICAgICBpZiAoT2JqZWN0LmtleXMoZXZlbnRQcm9wZXJ0aWVzKS5sZW5ndGggPiAwKSB7XG4gICAgICAgIHByb3BlcnRpZXNDb250YWluZXIuYXBwZW5kQ2hpbGQodGhpcy4jYnVpbGRQcm9wZXJ0eVRyZWUoZXZlbnRQcm9wZXJ0aWVzKSk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBwcm9wZXJ0aWVzQ29udGFpbmVyLnRleHRDb250ZW50ID0gJ05vIHByb3BlcnRpZXMgdG8gZGlzcGxheSc7XG4gICAgICB9XG5cbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keS5hcHBlbmRDaGlsZChwcm9wZXJ0aWVzSGVhZGluZyk7XG4gICAgICB0aGlzLiNldmVudERldGFpbEJvZHkuYXBwZW5kQ2hpbGQocHJvcGVydGllc0NvbnRhaW5lcik7XG4gICAgfVxuICB9XG5cbiAgI2J1aWxkUHJvcGVydGllc0ZvckRpc3BsYXkoZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkKTogUmVjb3JkPHN0cmluZywgdW5rbm93bj4ge1xuICAgIGNvbnN0IHByb3BlcnRpZXM6IFJlY29yZDxzdHJpbmcsIHVua25vd24+ID0ge307XG5cbiAgICBpZiAoZXZlbnRSZWNvcmQuY3VzdG9tUHJvcGVydGllcykge1xuICAgICAgT2JqZWN0LmFzc2lnbihwcm9wZXJ0aWVzLCBldmVudFJlY29yZC5jdXN0b21Qcm9wZXJ0aWVzKTtcbiAgICB9XG5cbiAgICBwcm9wZXJ0aWVzLmJ1YmJsZXMgPSBldmVudFJlY29yZC5idWJibGVzO1xuICAgIHByb3BlcnRpZXMuY2FuY2VsYWJsZSA9IGV2ZW50UmVjb3JkLmNhbmNlbGFibGU7XG4gICAgcHJvcGVydGllcy5kZWZhdWx0UHJldmVudGVkID0gZXZlbnRSZWNvcmQuZGVmYXVsdFByZXZlbnRlZDtcbiAgICBwcm9wZXJ0aWVzLmNvbXBvc2VkID0gZXZlbnRSZWNvcmQuY29tcG9zZWQ7XG5cbiAgICBpZiAoZXZlbnRSZWNvcmQubWFuaWZlc3RUeXBlKSB7XG4gICAgICBwcm9wZXJ0aWVzLnR5cGUgPSBldmVudFJlY29yZC5tYW5pZmVzdFR5cGU7XG4gICAgfVxuXG4gICAgcmV0dXJuIHByb3BlcnRpZXM7XG4gIH1cblxuICAjYnVpbGRQcm9wZXJ0eVRyZWUob2JqOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiwgZGVwdGggPSAwKTogSFRNTFVMaXN0RWxlbWVudCB7XG4gICAgY29uc3QgdWwgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd1bCcpO1xuICAgIHVsLmNsYXNzTmFtZSA9ICdldmVudC1wcm9wZXJ0eS10cmVlJztcbiAgICBpZiAoZGVwdGggPiAwKSB7XG4gICAgICB1bC5jbGFzc0xpc3QuYWRkKCduZXN0ZWQnKTtcbiAgICB9XG5cbiAgICBmb3IgKGNvbnN0IFtrZXksIHZhbHVlXSBvZiBPYmplY3QuZW50cmllcyhvYmopKSB7XG4gICAgICBjb25zdCBsaSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2xpJyk7XG4gICAgICBsaS5jbGFzc05hbWUgPSAncHJvcGVydHktaXRlbSc7XG5cbiAgICAgIGNvbnN0IGtleVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICBrZXlTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS1rZXknO1xuICAgICAga2V5U3Bhbi50ZXh0Q29udGVudCA9IGtleTtcblxuICAgICAgY29uc3QgY29sb25TcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgY29sb25TcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS1jb2xvbic7XG4gICAgICBjb2xvblNwYW4udGV4dENvbnRlbnQgPSAnOiAnO1xuXG4gICAgICBsaS5hcHBlbmRDaGlsZChrZXlTcGFuKTtcbiAgICAgIGxpLmFwcGVuZENoaWxkKGNvbG9uU3Bhbik7XG5cbiAgICAgIGlmICh2YWx1ZSA9PT0gbnVsbCB8fCB2YWx1ZSA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBudWxsJztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH0gZWxzZSBpZiAodHlwZW9mIHZhbHVlID09PSAnYm9vbGVhbicpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIGJvb2xlYW4nO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdudW1iZXInKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBudW1iZXInO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdzdHJpbmcnKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBzdHJpbmcnO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBgXCIke3ZhbHVlfVwiYDtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH0gZWxzZSBpZiAoQXJyYXkuaXNBcnJheSh2YWx1ZSkpIHtcbiAgICAgICAgY29uc3QgdmFsdWVTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAgICB2YWx1ZVNwYW4uY2xhc3NOYW1lID0gJ3Byb3BlcnR5LXZhbHVlIGFycmF5JztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gYEFycmF5KCR7dmFsdWUubGVuZ3RofSlgO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuXG4gICAgICAgIGlmICh2YWx1ZS5sZW5ndGggPiAwICYmIGRlcHRoIDwgMykge1xuICAgICAgICAgIGNvbnN0IG5lc3RlZE9iajogUmVjb3JkPHN0cmluZywgdW5rbm93bj4gPSB7fTtcbiAgICAgICAgICB2YWx1ZS5mb3JFYWNoKChpdGVtLCBpbmRleCkgPT4ge1xuICAgICAgICAgICAgbmVzdGVkT2JqW2luZGV4XSA9IGl0ZW07XG4gICAgICAgICAgfSk7XG4gICAgICAgICAgbGkuYXBwZW5kQ2hpbGQodGhpcy4jYnVpbGRQcm9wZXJ0eVRyZWUobmVzdGVkT2JqLCBkZXB0aCArIDEpKTtcbiAgICAgICAgfVxuICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdvYmplY3QnKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBvYmplY3QnO1xuICAgICAgICBjb25zdCBrZXlzID0gT2JqZWN0LmtleXModmFsdWUgYXMgUmVjb3JkPHN0cmluZywgdW5rbm93bj4pO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBrZXlzLmxlbmd0aCA+IDAgPyBgT2JqZWN0YCA6IGB7fWA7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG5cbiAgICAgICAgaWYgKGtleXMubGVuZ3RoID4gMCAmJiBkZXB0aCA8IDMpIHtcbiAgICAgICAgICBsaS5hcHBlbmRDaGlsZCh0aGlzLiNidWlsZFByb3BlcnR5VHJlZSh2YWx1ZSBhcyBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiwgZGVwdGggKyAxKSk7XG4gICAgICAgIH1cbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSc7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9XG5cbiAgICAgIHVsLmFwcGVuZENoaWxkKGxpKTtcbiAgICB9XG5cbiAgICByZXR1cm4gdWw7XG4gIH1cblxuICAjc2Nyb2xsRXZlbnRzVG9Cb3R0b20oKSB7XG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIHJlcXVlc3RBbmltYXRpb25GcmFtZSgoKSA9PiB7XG4gICAgICBjb25zdCBsYXN0RXZlbnQgPSB0aGlzLiNldmVudExpc3QhLmxhc3RFbGVtZW50Q2hpbGQ7XG4gICAgICBpZiAobGFzdEV2ZW50KSB7XG4gICAgICAgIGxhc3RFdmVudC5zY3JvbGxJbnRvVmlldyh7IGJlaGF2aW9yOiAnYXV0bycsIGJsb2NrOiAnZW5kJyB9KTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNpc0V2ZW50c1RhYkFjdGl2ZSgpOiBib29sZWFuIHtcbiAgICBjb25zdCB0YWJzID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdwZi12Ni10YWJzJyk7XG4gICAgaWYgKCF0YWJzKSByZXR1cm4gZmFsc2U7XG5cbiAgICBjb25zdCBzZWxlY3RlZEluZGV4ID0gcGFyc2VJbnQodGFicy5nZXRBdHRyaWJ1dGUoJ3NlbGVjdGVkJykgfHwgJzAnLCAxMCk7XG4gICAgcmV0dXJuIHNlbGVjdGVkSW5kZXggPT09IDM7XG4gIH1cblxuICAjZmlsdGVyRXZlbnRzKHF1ZXJ5OiBzdHJpbmcpIHtcbiAgICB0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoIXRoaXMuI2V2ZW50TGlzdCkgcmV0dXJuO1xuXG4gICAgZm9yIChjb25zdCBlbnRyeSBvZiB0aGlzLiNldmVudExpc3QuY2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IGV2ZW50UmVjb3JkID0gdGhpcy4jZ2V0RXZlbnRSZWNvcmRCeUlkKChlbnRyeSBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5ldmVudElkISk7XG5cbiAgICAgIGlmICghZXZlbnRSZWNvcmQpIGNvbnRpbnVlO1xuXG4gICAgICBjb25zdCB0ZXh0TWF0Y2ggPSB0aGlzLiNldmVudE1hdGNoZXNUZXh0RmlsdGVyKGV2ZW50UmVjb3JkKTtcbiAgICAgIGNvbnN0IHR5cGVNYXRjaCA9IHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmhhcyhldmVudFJlY29yZC5ldmVudE5hbWUpO1xuICAgICAgY29uc3QgZWxlbWVudE1hdGNoID0gdGhpcy4jZWxlbWVudEZpbHRlcnMuc2l6ZSA9PT0gMCB8fCB0aGlzLiNlbGVtZW50RmlsdGVycy5oYXMoZXZlbnRSZWNvcmQudGFnTmFtZSk7XG5cbiAgICAgIChlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuID0gISh0ZXh0TWF0Y2ggJiYgdHlwZU1hdGNoICYmIGVsZW1lbnRNYXRjaCk7XG4gICAgfVxuICB9XG5cbiAgI2V2ZW50TWF0Y2hlc1RleHRGaWx0ZXIoZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkKTogYm9vbGVhbiB7XG4gICAgaWYgKCF0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSkgcmV0dXJuIHRydWU7XG5cbiAgICBjb25zdCBzZWFyY2hUZXh0ID0gW1xuICAgICAgZXZlbnRSZWNvcmQudGFnTmFtZSxcbiAgICAgIGV2ZW50UmVjb3JkLmV2ZW50TmFtZSxcbiAgICAgIGV2ZW50UmVjb3JkLmVsZW1lbnRJZCB8fCAnJyxcbiAgICAgIEpTT04uc3RyaW5naWZ5KGV2ZW50UmVjb3JkLmN1c3RvbVByb3BlcnRpZXMgfHwge30pXG4gICAgXS5qb2luKCcgJykudG9Mb3dlckNhc2UoKTtcblxuICAgIHJldHVybiBzZWFyY2hUZXh0LmluY2x1ZGVzKHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKTtcbiAgfVxuXG4gICNnZXRFdmVudFJlY29yZEJ5SWQoaWQ6IHN0cmluZyk6IEV2ZW50UmVjb3JkIHwgdW5kZWZpbmVkIHtcbiAgICByZXR1cm4gdGhpcy4jY2FwdHVyZWRFdmVudHMuZmluZChlID0+IGUuaWQgPT09IGlkKTtcbiAgfVxuXG4gICN1cGRhdGVFdmVudEZpbHRlcnMoKSB7XG4gICAgY29uc3Qgc2F2ZWRQcmVmZXJlbmNlcyA9IHRoaXMuI2xvYWRFdmVudEZpbHRlcnNGcm9tU3RvcmFnZSgpO1xuXG4gICAgY29uc3QgZXZlbnRUeXBlRmlsdGVyID0gdGhpcy4jJCgnZXZlbnQtdHlwZS1maWx0ZXInKTtcbiAgICBpZiAoZXZlbnRUeXBlRmlsdGVyICYmIHRoaXMuI2VsZW1lbnRFdmVudE1hcCkge1xuICAgICAgbGV0IG1lbnUgPSBldmVudFR5cGVGaWx0ZXIucXVlcnlTZWxlY3RvcigncGYtdjYtbWVudScpO1xuICAgICAgaWYgKCFtZW51KSB7XG4gICAgICAgIG1lbnUgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwZi12Ni1tZW51Jyk7XG4gICAgICAgIGV2ZW50VHlwZUZpbHRlci5hcHBlbmRDaGlsZChtZW51KTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgZXhpc3RpbmdJdGVtcyA9IG1lbnUucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICBleGlzdGluZ0l0ZW1zLmZvckVhY2goaXRlbSA9PiBpdGVtLnJlbW92ZSgpKTtcblxuICAgICAgY29uc3QgYWxsRXZlbnRUeXBlcyA9IG5ldyBTZXQ8c3RyaW5nPigpO1xuICAgICAgZm9yIChjb25zdCBbdGFnTmFtZSwgZXZlbnRJbmZvXSBvZiB0aGlzLiNlbGVtZW50RXZlbnRNYXApIHtcbiAgICAgICAgaWYgKHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5oYXModGFnTmFtZSkpIHtcbiAgICAgICAgICBmb3IgKGNvbnN0IGV2ZW50TmFtZSBvZiBldmVudEluZm8uZXZlbnROYW1lcykge1xuICAgICAgICAgICAgYWxsRXZlbnRUeXBlcy5hZGQoZXZlbnROYW1lKTtcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgaWYgKHNhdmVkUHJlZmVyZW5jZXMuZXZlbnRUeXBlcykge1xuICAgICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzID0gKHNhdmVkUHJlZmVyZW5jZXMuZXZlbnRUeXBlcyBhcyBTZXQ8c3RyaW5nPiAmIHsgaW50ZXJzZWN0aW9uOiAob3RoZXI6IFNldDxzdHJpbmc+KSA9PiBTZXQ8c3RyaW5nPiB9KS5pbnRlcnNlY3Rpb24oYWxsRXZlbnRUeXBlcyk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzID0gbmV3IFNldChhbGxFdmVudFR5cGVzKTtcbiAgICAgIH1cblxuICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgYWxsRXZlbnRUeXBlcykge1xuICAgICAgICBjb25zdCBpdGVtID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YXJpYW50JywgJ2NoZWNrYm94Jyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YWx1ZScsIGV2ZW50TmFtZSk7XG4gICAgICAgIGlmICh0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmhhcyhldmVudE5hbWUpKSB7XG4gICAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2NoZWNrZWQnLCAnJyk7XG4gICAgICAgIH1cbiAgICAgICAgaXRlbS50ZXh0Q29udGVudCA9IGV2ZW50TmFtZTtcbiAgICAgICAgbWVudS5hcHBlbmRDaGlsZChpdGVtKTtcbiAgICAgIH1cbiAgICB9XG5cbiAgICBjb25zdCBlbGVtZW50RmlsdGVyID0gdGhpcy4jJCgnZWxlbWVudC1maWx0ZXInKTtcbiAgICBpZiAoZWxlbWVudEZpbHRlciAmJiB0aGlzLiNlbGVtZW50RXZlbnRNYXApIHtcbiAgICAgIGxldCBtZW51ID0gZWxlbWVudEZpbHRlci5xdWVyeVNlbGVjdG9yKCdwZi12Ni1tZW51Jyk7XG4gICAgICBpZiAoIW1lbnUpIHtcbiAgICAgICAgbWVudSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3BmLXY2LW1lbnUnKTtcbiAgICAgICAgZWxlbWVudEZpbHRlci5hcHBlbmRDaGlsZChtZW51KTtcbiAgICAgIH1cblxuICAgICAgY29uc3QgZXhpc3RpbmdJdGVtcyA9IG1lbnUucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICBleGlzdGluZ0l0ZW1zLmZvckVhY2goaXRlbSA9PiBpdGVtLnJlbW92ZSgpKTtcblxuICAgICAgY29uc3QgYWxsRWxlbWVudHMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgICAgIGZvciAoY29uc3QgdGFnTmFtZSBvZiB0aGlzLiNlbGVtZW50RXZlbnRNYXAua2V5cygpKSB7XG4gICAgICAgIGlmICh0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuaGFzKHRhZ05hbWUpKSB7XG4gICAgICAgICAgYWxsRWxlbWVudHMuYWRkKHRhZ05hbWUpO1xuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIGlmIChzYXZlZFByZWZlcmVuY2VzLmVsZW1lbnRzKSB7XG4gICAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzID0gKHNhdmVkUHJlZmVyZW5jZXMuZWxlbWVudHMgYXMgU2V0PHN0cmluZz4gJiB7IGludGVyc2VjdGlvbjogKG90aGVyOiBTZXQ8c3RyaW5nPikgPT4gU2V0PHN0cmluZz4gfSkuaW50ZXJzZWN0aW9uKGFsbEVsZW1lbnRzKTtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzID0gbmV3IFNldChhbGxFbGVtZW50cyk7XG4gICAgICB9XG5cbiAgICAgIGZvciAoY29uc3QgdGFnTmFtZSBvZiBhbGxFbGVtZW50cykge1xuICAgICAgICBjb25zdCBpdGVtID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncGYtdjYtbWVudS1pdGVtJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YXJpYW50JywgJ2NoZWNrYm94Jyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCd2YWx1ZScsIHRhZ05hbWUpO1xuICAgICAgICBpZiAodGhpcy4jZWxlbWVudEZpbHRlcnMuaGFzKHRhZ05hbWUpKSB7XG4gICAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ2NoZWNrZWQnLCAnJyk7XG4gICAgICAgIH1cbiAgICAgICAgaXRlbS50ZXh0Q29udGVudCA9IGA8JHt0YWdOYW1lfT5gO1xuICAgICAgICBtZW51LmFwcGVuZENoaWxkKGl0ZW0pO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNoYW5kbGVFdmVudFR5cGVGaWx0ZXJDaGFuZ2UgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgeyB2YWx1ZSwgY2hlY2tlZCB9ID0gZXZlbnQgYXMgRXZlbnQgJiB7IHZhbHVlOiBzdHJpbmc7IGNoZWNrZWQ6IGJvb2xlYW4gfTtcblxuICAgIGlmICghdmFsdWUpIHJldHVybjtcblxuICAgIGlmIChjaGVja2VkKSB7XG4gICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2V2ZW50VHlwZUZpbHRlcnMuZGVsZXRlKHZhbHVlKTtcbiAgICB9XG5cbiAgICB0aGlzLiNzYXZlRXZlbnRGaWx0ZXJzKCk7XG4gICAgdGhpcy4jZmlsdGVyRXZlbnRzKHRoaXMuI2V2ZW50c0ZpbHRlclZhbHVlKTtcbiAgfTtcblxuICAjaGFuZGxlRWxlbWVudEZpbHRlckNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB7IHZhbHVlLCBjaGVja2VkIH0gPSBldmVudCBhcyBFdmVudCAmIHsgdmFsdWU6IHN0cmluZzsgY2hlY2tlZDogYm9vbGVhbiB9O1xuXG4gICAgaWYgKCF2YWx1ZSkgcmV0dXJuO1xuXG4gICAgaWYgKGNoZWNrZWQpIHtcbiAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmFkZCh2YWx1ZSk7XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmRlbGV0ZSh2YWx1ZSk7XG4gICAgfVxuXG4gICAgdGhpcy4jc2F2ZUV2ZW50RmlsdGVycygpO1xuICAgIHRoaXMuI2ZpbHRlckV2ZW50cyh0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSk7XG4gIH07XG5cbiAgI2xvYWRFdmVudEZpbHRlcnNGcm9tU3RvcmFnZSgpOiB7IGV2ZW50VHlwZXM6IFNldDxzdHJpbmc+IHwgbnVsbDsgZWxlbWVudHM6IFNldDxzdHJpbmc+IHwgbnVsbCB9IHtcbiAgICBjb25zdCBwcmVmZXJlbmNlczogeyBldmVudFR5cGVzOiBTZXQ8c3RyaW5nPiB8IG51bGw7IGVsZW1lbnRzOiBTZXQ8c3RyaW5nPiB8IG51bGwgfSA9IHtcbiAgICAgIGV2ZW50VHlwZXM6IG51bGwsXG4gICAgICBlbGVtZW50czogbnVsbFxuICAgIH07XG5cbiAgICB0cnkge1xuICAgICAgY29uc3Qgc2F2ZWRFdmVudFR5cGVzID0gbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2NlbS1zZXJ2ZS1ldmVudC10eXBlLWZpbHRlcnMnKTtcbiAgICAgIGlmIChzYXZlZEV2ZW50VHlwZXMpIHtcbiAgICAgICAgcHJlZmVyZW5jZXMuZXZlbnRUeXBlcyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZEV2ZW50VHlwZXMpKTtcbiAgICAgIH1cblxuICAgICAgY29uc3Qgc2F2ZWRFbGVtZW50cyA9IGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtZWxlbWVudC1maWx0ZXJzJyk7XG4gICAgICBpZiAoc2F2ZWRFbGVtZW50cykge1xuICAgICAgICBwcmVmZXJlbmNlcy5lbGVtZW50cyA9IG5ldyBTZXQoSlNPTi5wYXJzZShzYXZlZEVsZW1lbnRzKSk7XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgY29uc29sZS5kZWJ1ZygnW2NlbS1zZXJ2ZS1jaHJvbWVdIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSBmb3IgZXZlbnQgZmlsdGVycycpO1xuICAgIH1cblxuICAgIHJldHVybiBwcmVmZXJlbmNlcztcbiAgfVxuXG4gICNzYXZlRXZlbnRGaWx0ZXJzKCkge1xuICAgIHRyeSB7XG4gICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLWV2ZW50LXR5cGUtZmlsdGVycycsXG4gICAgICAgIEpTT04uc3RyaW5naWZ5KFsuLi50aGlzLiNldmVudFR5cGVGaWx0ZXJzXSkpO1xuICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oJ2NlbS1zZXJ2ZS1lbGVtZW50LWZpbHRlcnMnLFxuICAgICAgICBKU09OLnN0cmluZ2lmeShbLi4udGhpcy4jZWxlbWVudEZpbHRlcnNdKSk7XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgLy8gbG9jYWxTdG9yYWdlIHVuYXZhaWxhYmxlIChwcml2YXRlIG1vZGUpLCBzaWxlbnRseSBjb250aW51ZVxuICAgIH1cbiAgfVxuXG4gICNjbGVhckV2ZW50cygpIHtcbiAgICB0aGlzLiNjYXB0dXJlZEV2ZW50cyA9IFtdO1xuICAgIHRoaXMuI3NlbGVjdGVkRXZlbnRJZCA9IG51bGw7XG4gICAgaWYgKHRoaXMuI2V2ZW50TGlzdCkge1xuICAgICAgdGhpcy4jZXZlbnRMaXN0LnJlcGxhY2VDaGlsZHJlbigpO1xuICAgIH1cbiAgICBpZiAodGhpcy4jZXZlbnREZXRhaWxIZWFkZXIpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyLmlubmVySFRNTCA9ICcnO1xuICAgIH1cbiAgICBpZiAodGhpcy4jZXZlbnREZXRhaWxCb2R5KSB7XG4gICAgICB0aGlzLiNldmVudERldGFpbEJvZHkuaW5uZXJIVE1MID0gJyc7XG4gICAgfVxuICB9XG5cbiAgYXN5bmMgI2NvcHlFdmVudHMoKSB7XG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIGNvbnN0IHZpc2libGVFdmVudHMgPSBBcnJheS5mcm9tKHRoaXMuI2V2ZW50TGlzdC5jaGlsZHJlbilcbiAgICAgIC5maWx0ZXIoZW50cnkgPT4gIShlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuKVxuICAgICAgLm1hcChlbnRyeSA9PiB7XG4gICAgICAgIGNvbnN0IGlkID0gKGVudHJ5IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LmV2ZW50SWQhO1xuICAgICAgICByZXR1cm4gdGhpcy4jZ2V0RXZlbnRSZWNvcmRCeUlkKGlkKTtcbiAgICAgIH0pXG4gICAgICAuZmlsdGVyKChldmVudCk6IGV2ZW50IGlzIEV2ZW50UmVjb3JkID0+ICEhZXZlbnQpXG4gICAgICAubWFwKGV2ZW50ID0+IHtcbiAgICAgICAgY29uc3QgdGltZSA9IGV2ZW50LnRpbWVzdGFtcC50b0xvY2FsZVRpbWVTdHJpbmcoKTtcbiAgICAgICAgY29uc3QgZWxlbWVudCA9IGV2ZW50LmVsZW1lbnRJZCA/IGAjJHtldmVudC5lbGVtZW50SWR9YCA6IGV2ZW50LnRhZ05hbWU7XG4gICAgICAgIGNvbnN0IHByb3BzID0gZXZlbnQuY3VzdG9tUHJvcGVydGllcyAmJiBPYmplY3Qua2V5cyhldmVudC5jdXN0b21Qcm9wZXJ0aWVzKS5sZW5ndGggPiAwXG4gICAgICAgICAgPyBgIC0gJHtKU09OLnN0cmluZ2lmeShldmVudC5jdXN0b21Qcm9wZXJ0aWVzKX1gXG4gICAgICAgICAgOiAnJztcbiAgICAgICAgcmV0dXJuIGBbJHt0aW1lfV0gPCR7ZXZlbnQudGFnTmFtZX0+ICR7ZWxlbWVudH0gXFx1MjE5MiAke2V2ZW50LmV2ZW50TmFtZX0ke3Byb3BzfWA7XG4gICAgICB9KVxuICAgICAgLmpvaW4oJ1xcbicpO1xuXG4gICAgaWYgKCF2aXNpYmxlRXZlbnRzKSByZXR1cm47XG5cbiAgICB0cnkge1xuICAgICAgYXdhaXQgbmF2aWdhdG9yLmNsaXBib2FyZC53cml0ZVRleHQodmlzaWJsZUV2ZW50cyk7XG4gICAgICBjb25zdCBidG4gPSB0aGlzLiMkKCdjb3B5LWV2ZW50cycpO1xuICAgICAgaWYgKGJ0bikge1xuICAgICAgICBjb25zdCB0ZXh0Tm9kZSA9IEFycmF5LmZyb20oYnRuLmNoaWxkTm9kZXMpLmZpbmQoXG4gICAgICAgICAgbiA9PiBuLm5vZGVUeXBlID09PSBOb2RlLlRFWFRfTk9ERSAmJiAobi50ZXh0Q29udGVudD8udHJpbSgpLmxlbmd0aCA/PyAwKSA+IDBcbiAgICAgICAgKTtcbiAgICAgICAgaWYgKHRleHROb2RlKSB7XG4gICAgICAgICAgY29uc3Qgb3JpZ2luYWwgPSB0ZXh0Tm9kZS50ZXh0Q29udGVudDtcbiAgICAgICAgICB0ZXh0Tm9kZS50ZXh0Q29udGVudCA9ICdDb3BpZWQhJztcblxuICAgICAgICAgIGlmICh0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgICAgaWYgKHRoaXMuaXNDb25uZWN0ZWQgJiYgdGV4dE5vZGUucGFyZW50Tm9kZSkge1xuICAgICAgICAgICAgICB0ZXh0Tm9kZS50ZXh0Q29udGVudCA9IG9yaWdpbmFsO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgICAgICAgfSwgMjAwMCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9IGNhdGNoIChlcnIpIHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY29weSBldmVudHM6JywgZXJyKTtcbiAgICB9XG4gIH1cblxuICAjc2V0dXBFdmVudExpc3RlbmVycygpIHtcbiAgICB0aGlzLiNldmVudExpc3QgPSB0aGlzLiMkKCdldmVudC1saXN0Jyk7XG4gICAgdGhpcy4jZXZlbnREZXRhaWxIZWFkZXIgPSB0aGlzLiMkKCdldmVudC1kZXRhaWwtaGVhZGVyJyk7XG4gICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5ID0gdGhpcy4jJCgnZXZlbnQtZGV0YWlsLWJvZHknKTtcblxuICAgIGlmICh0aGlzLiNldmVudExpc3QpIHtcbiAgICAgIHRoaXMuI2V2ZW50TGlzdC5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsIChlOiBFdmVudCkgPT4ge1xuICAgICAgICBjb25zdCBsaXN0SXRlbSA9IChlLnRhcmdldCBhcyBFbGVtZW50KS5jbG9zZXN0KCcuZXZlbnQtbGlzdC1pdGVtJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgICAgIGlmIChsaXN0SXRlbSkge1xuICAgICAgICAgIGNvbnN0IGV2ZW50SWQgPSBsaXN0SXRlbS5kYXRhc2V0LmV2ZW50SWQ7XG4gICAgICAgICAgaWYgKGV2ZW50SWQpIHtcbiAgICAgICAgICAgIHRoaXMuI3NlbGVjdEV2ZW50KGV2ZW50SWQpO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfSk7XG4gICAgfVxuXG4gICAgY29uc3QgZXZlbnRzRmlsdGVyID0gdGhpcy4jJCgnZXZlbnRzLWZpbHRlcicpO1xuICAgIGlmIChldmVudHNGaWx0ZXIpIHtcbiAgICAgIGV2ZW50c0ZpbHRlci5hZGRFdmVudExpc3RlbmVyKCdpbnB1dCcsIChlOiBFdmVudCkgPT4ge1xuICAgICAgICBjb25zdCB7IHZhbHVlID0gJycgfSA9IGUudGFyZ2V0IGFzIEhUTUxJbnB1dEVsZW1lbnQ7XG4gICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNldmVudHNGaWx0ZXJEZWJvdW5jZVRpbWVyISk7XG4gICAgICAgIHRoaXMuI2V2ZW50c0ZpbHRlckRlYm91bmNlVGltZXIgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICB0aGlzLiNmaWx0ZXJFdmVudHModmFsdWUpO1xuICAgICAgICB9LCAzMDApO1xuICAgICAgfSk7XG4gICAgfVxuXG4gICAgY29uc3QgZXZlbnRUeXBlRmlsdGVyID0gdGhpcy4jJCgnZXZlbnQtdHlwZS1maWx0ZXInKTtcbiAgICBpZiAoZXZlbnRUeXBlRmlsdGVyKSB7XG4gICAgICBldmVudFR5cGVGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jaGFuZGxlRXZlbnRUeXBlRmlsdGVyQ2hhbmdlIGFzIEV2ZW50TGlzdGVuZXIpO1xuICAgIH1cblxuICAgIGNvbnN0IGVsZW1lbnRGaWx0ZXIgPSB0aGlzLiMkKCdlbGVtZW50LWZpbHRlcicpO1xuICAgIGlmIChlbGVtZW50RmlsdGVyKSB7XG4gICAgICBlbGVtZW50RmlsdGVyLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZUVsZW1lbnRGaWx0ZXJDaGFuZ2UgYXMgRXZlbnRMaXN0ZW5lcik7XG4gICAgfVxuXG4gICAgdGhpcy4jJCgnY2xlYXItZXZlbnRzJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgdGhpcy4jY2xlYXJFdmVudHMoKTtcbiAgICB9KTtcblxuICAgIHRoaXMuIyQoJ2NvcHktZXZlbnRzJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgdGhpcy4jY29weUV2ZW50cygpO1xuICAgIH0pO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcblxuICAgIC8vIENsZWFuIHVwIGtub2IgbGlzdGVuZXJzXG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOmF0dHJpYnV0ZS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jaGFuZ2UnLCB0aGlzLiNvbktub2JDaGFuZ2UpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpjc3MtcHJvcGVydHktY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6YXR0cmlidXRlLWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjpwcm9wZXJ0eS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNsZWFyJywgdGhpcy4jb25Lbm9iQ2xlYXIpO1xuXG4gICAgLy8gQ2xlYW4gdXAgdHJlZSBzdGF0ZSBsaXN0ZW5lcnNcbiAgICBpZiAodGhpcy4jaGFuZGxlVHJlZUV4cGFuZCkge1xuICAgICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdleHBhbmQnLCB0aGlzLiNoYW5kbGVUcmVlRXhwYW5kKTtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSkge1xuICAgICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdjb2xsYXBzZScsIHRoaXMuI2hhbmRsZVRyZWVDb2xsYXBzZSk7XG4gICAgfVxuICAgIGlmICh0aGlzLiNoYW5kbGVUcmVlU2VsZWN0KSB7XG4gICAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QpO1xuICAgIH1cblxuICAgIC8vIENsZWFuIHVwIHdpbmRvdyBsaXN0ZW5lclxuICAgIGlmICh0aGlzLiNoYW5kbGVMb2dzRXZlbnQpIHtcbiAgICAgIHdpbmRvdy5yZW1vdmVFdmVudExpc3RlbmVyKCdjZW06bG9ncycsIHRoaXMuI2hhbmRsZUxvZ3NFdmVudCk7XG4gICAgfVxuXG4gICAgLy8gQ2xlYXIgcGVuZGluZyBmZWVkYmFjayB0aW1lb3V0c1xuICAgIGlmICh0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgIHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICB0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgIH1cbiAgICBpZiAodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgdGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgfVxuXG4gICAgLy8gRGlzY29ubmVjdCBtdXRhdGlvbiBvYnNlcnZlclxuICAgIHRoaXMuI29ic2VydmVyLmRpc2Nvbm5lY3QoKTtcblxuICAgIC8vIENsb3NlIFdlYlNvY2tldCBjb25uZWN0aW9uXG4gICAgaWYgKHRoaXMuI3dzQ2xpZW50KSB7XG4gICAgICB0aGlzLiN3c0NsaWVudC5kZXN0cm95KCk7XG4gICAgfVxuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1zZXJ2ZS1jaHJvbWUnOiBDZW1TZXJ2ZUNocm9tZTtcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbiAgaGVpZ2h0OiAxMDB2aDtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcbiAgLS1wZi12Ni1jLW1hc3RoZWFkX19sb2dvLS1XaWR0aDogbWF4LWNvbnRlbnQ7XFxcXG4gIC0tcGYtdjYtYy1tYXN0aGVhZF9fdG9nZ2xlLS1TaXplOiAxcmVtO1xcXFxufVxcXFxuXFxcXG5baGlkZGVuXSB7XFxcXG4gIGRpc3BsYXk6IG5vbmUgIWltcG9ydGFudDtcXFxcbn1cXFxcblxcXFxuLyogTWFzdGhlYWQgbG9nbyBzdHlsZXMgKi9cXFxcbi5tYXN0aGVhZC1sb2dvIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuICBjb2xvcjogaW5oZXJpdDtcXFxcbiAgbWF4LWhlaWdodDogdmFyKC0tcGYtdjYtYy1tYXN0aGVhZF9fbG9nby0tTWF4SGVpZ2h0KTtcXFxcbiAgZ2FwOiA0cHg7XFxcXG4gIFxcXFx1MDAyNiBpbWcge1xcXFxuICAgIGRpc3BsYXk6IGJsb2NrO1xcXFxuICAgIG1heC1oZWlnaHQ6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX2xvZ28tLU1heEhlaWdodCk7XFxcXG4gICAgd2lkdGg6IGF1dG87XFxcXG4gIH1cXFxcbiAgXFxcXHUwMDI2IDo6c2xvdHRlZChbc2xvdD1cXFxcXFxcInRpdGxlXFxcXFxcXCJdKSB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICAgIGZvbnQtc2l6ZTogMS4xMjVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIH1cXFxcbiAgXFxcXHUwMDI2IGgxIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gICAgZm9udC1zaXplOiAxOHB4O1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIFRvb2xiYXIgZ3JvdXAgYWxpZ25tZW50ICovXFxcXG5wZi12Ni10b29sYmFyLWdyb3VwW3ZhcmlhbnQ9XFxcXFxcXCJhY3Rpb24tZ3JvdXBcXFxcXFxcIl0ge1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiBhdXRvO1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxufVxcXFxuXFxcXG4uZGVidWctcGFuZWwge1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyOiAxcHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICBwYWRkaW5nOiAxLjVyZW07XFxcXG4gIG1heC13aWR0aDogNjAwcHg7XFxcXG4gIHdpZHRoOiA5MCU7XFxcXG4gIG1heC1oZWlnaHQ6IDgwdmg7XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG5cXFxcbiAgaDIge1xcXFxuICAgIG1hcmdpbjogMCAwIDFyZW0gMDtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgZm9udC1zaXplOiAxLjEyNXJlbTtcXFxcbiAgICBmb250LXdlaWdodDogNjAwO1xcXFxuICB9XFxcXG5cXFxcbiAgZGwge1xcXFxuICAgIG1hcmdpbjogMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGR0IHtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgIG1hcmdpbi10b3A6IDAuNXJlbTtcXFxcbiAgICBmb250LXdlaWdodDogNTAwO1xcXFxuICB9XFxcXG5cXFxcbiAgZGQge1xcXFxuICAgIG1hcmdpbjogMCAwIDAuNXJlbSAwO1xcXFxuICAgIGZvbnQtZmFtaWx5OiB1aS1tb25vc3BhY2UsICdDYXNjYWRpYSBDb2RlJywgJ1NvdXJjZSBDb2RlIFBybycsIE1lbmxvLCBDb25zb2xhcywgJ0RlamFWdSBTYW5zIE1vbm8nLCBtb25vc3BhY2U7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gIH1cXFxcblxcXFxuICBkZXRhaWxzIHtcXFxcbiAgICBtYXJnaW4tdG9wOiAxcmVtO1xcXFxuXFxcXG4gICAgc3VtbWFyeSB7XFxcXG4gICAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tcmVndWxhcik7XFxcXG4gICAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgICAgZm9udC13ZWlnaHQ6IDUwMDtcXFxcbiAgICAgIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gICAgICBkaXNwbGF5OiBmbGV4O1xcXFxuICAgICAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgICAgIGdhcDogMC41cmVtO1xcXFxuICAgICAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG5cXFxcbiAgICAgIFxcXFx1MDAyNjo6LXdlYmtpdC1kZXRhaWxzLW1hcmtlciB7XFxcXG4gICAgICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gICAgICB9XFxcXG5cXFxcbiAgICAgIFxcXFx1MDAyNjo6YmVmb3JlIHtcXFxcbiAgICAgICAgY29udGVudDogJ1xcXFxcXFxcMjVCOCc7XFxcXG4gICAgICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgICAgICAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIDEwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgICAgfVxcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNltvcGVuXSBzdW1tYXJ5OjpiZWZvcmUge1xcXFxuICAgICAgdHJhbnNmb3JtOiByb3RhdGUoOTBkZWcpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIHByZSB7XFxcXG4gICAgICBtYXJnaW4tdG9wOiAwLjVyZW07XFxcXG4gICAgICBtYXJnaW4tbGVmdDogMS41cmVtO1xcXFxuICAgICAgcGFkZGluZzogMC41cmVtO1xcXFxuICAgICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0KTtcXFxcbiAgICAgIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgICBvdmVyZmxvdy14OiBhdXRvO1xcXFxuICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIC5idXR0b24tY29udGFpbmVyIHtcXFxcbiAgICBkaXNwbGF5OiBmbGV4O1xcXFxuICAgIGdhcDogMC41cmVtO1xcXFxuICAgIG1hcmdpbi10b3A6IDFyZW07XFxcXG4gIH1cXFxcblxcXFxuICBwIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICB9XFxcXG5cXFxcbiAgYnV0dG9uIHtcXFxcbiAgICBtYXJnaW4tdG9wOiAxcmVtO1xcXFxuICAgIHBhZGRpbmc6IDAuNXJlbSAxcmVtO1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tYnJhbmQpO1xcXFxuICAgIGJvcmRlcjogbm9uZTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBmb250LXdlaWdodDogNDAwO1xcXFxuICAgIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgICB0cmFuc2l0aW9uOiBhbGwgMjAwbXMgY3ViaWMtYmV6aWVyKDAuNjQ1LCAwLjA0NSwgMC4zNTUsIDEpO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1ob3Zlcik7XFxcXG4gICAgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIENvbnRlbnQgYXJlYSBwYWRkaW5nIGZvciBkZW1vICovXFxcXG5wZi12Ni1wYWdlLW1haW4ge1xcXFxuICBtaW4taGVpZ2h0OiBjYWxjKDEwMGR2aCAtIDcycHggLSB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0taW5zZXQtLXBhZ2UtY2hyb21lKSk7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIFxcXFx1MDAyNiBcXFxcdTAwM2UgOjpzbG90dGVkKDpub3QoW3Nsb3Q9a25vYnNdKSkge1xcXFxuICAgIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gICAgZmxleDogMTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG5jZW0tZHJhd2VyIHtcXFxcbiAgcGYtdjYtdGFicyB7XFxcXG4gICAgcGYtdjYtdGFiIHtcXFxcbiAgICAgIHBhZGRpbmctYmxvY2stZW5kOiAwO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBFbGVtZW50IGRlc2NyaXB0aW9ucyBpbiBsaXN0aW5nICovXFxcXG4uZWxlbWVudC1zdW1tYXJ5IHtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtKTtcXFxcbn1cXFxcblxcXFxuLmVsZW1lbnQtZGVzY3JpcHRpb24ge1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tc20pO1xcXFxufVxcXFxuXFxcXG4vKiBDYXJkIGZvb3RlciBkZW1vIG5hdmlnYXRpb24gKi9cXFxcbi5jYXJkLWRlbW9zIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC13cmFwOiB3cmFwO1xcXFxuICBnYXA6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1nYXAtLWFjdGlvbi10by1hY3Rpb24tLWRlZmF1bHQpO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBtYXJnaW46IDA7XFxcXG59XFxcXG5cXFxcbi5wYWNrYWdlLW5hbWUge1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtKTtcXFxcbn1cXFxcblxcXFxuLyogS25vYnMgY29udGFpbmVyIC0gZmlsbHMgdGFiIHBhbmVsIGhlaWdodCAqL1xcXFxuI2tub2JzLWNvbnRhaW5lciB7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgXFxcXHUwMDI2IDo6c2xvdHRlZChbc2xvdD1cXFxcXFxcImtub2JzXFxcXFxcXCJdKSB7XFxcXG4gICAgZmxleDogMTtcXFxcbiAgICBtaW4taGVpZ2h0OiAwO1xcXFxuICAgIG92ZXJmbG93OiBoaWRkZW47XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmtub2JzLWVtcHR5IHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIHRleHQtYWxpZ246IGNlbnRlcjtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG5cXFxcbiAgY29kZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICAgIHBhZGRpbmc6IDJweCA2cHg7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogM3B4O1xcXFxuICAgIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4uaW5zdGFuY2UtdGFnIHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4uaW5zdGFuY2UtbGFiZWwge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWdyb3VwIHtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2Omxhc3QtY2hpbGQge1xcXFxuICAgIG1hcmdpbi1ib3R0b206IDA7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogUGF0dGVybkZseSB2NiBmb3JtIC0gaG9yaXpvbnRhbCBsYXlvdXQgKi9cXFxcbnBmLXY2LWZvcm1baG9yaXpvbnRhbF0gcGYtdjYtZm9ybS1maWVsZC1ncm91cCB7XFxcXG4gIGdyaWQtY29sdW1uOiBzcGFuIDJcXFxcbn1cXFxcblxcXFxuLmtub2ItZ3JvdXAtdGl0bGUge1xcXFxuICBncmlkLWNvbHVtbjogMSAvIC0xO1xcXFxuICBtYXJnaW46IDAgMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtYmFzZSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gIGJvcmRlci1ib3R0b206IDFweCBzb2xpZCB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItY29sb3IpO1xcXFxuICBwYWRkaW5nLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWNvbnRyb2wge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbn1cXFxcblxcXFxuLmtub2ItbGFiZWwge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG59XFxcXG5cXFxcbi5rbm9iLW5hbWUge1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWRlc2NyaXB0aW9uIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxuXFxcXG4gIHAge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiAzcHg7XFxcXG4gICAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICB9XFxcXG5cXFxcbiAgYSB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIHRleHQtZGVjb3JhdGlvbjogdW5kZXJsaW5lO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIHN0cm9uZyB7XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgfVxcXFxuXFxcXG4gIHVsLCBvbCB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSAwO1xcXFxuICAgIHBhZGRpbmctbGVmdDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuZm9vdGVyLnBmLW0tc3RpY2t5LWJvdHRvbSB7XFxcXG4gIHZpZXctdHJhbnNpdGlvbi1uYW1lOiBkZXYtc2VydmVyLWZvb3RlcjtcXFxcbiAgcG9zaXRpb246IHN0aWNreTtcXFxcbiAgYm90dG9tOiAwO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1kZWZhdWx0KTtcXFxcbiAgYm9yZGVyLXRvcDogMXB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIHotaW5kZXg6IHZhcigtLXBmLXY2LWMtcGFnZS0tc2VjdGlvbi0tbS1zdGlja3ktYm90dG9tLS1aSW5kZXgsIHZhcigtLXBmLXQtLWdsb2JhbC0tei1pbmRleC0tbWQpKTtcXFxcbiAgYm94LXNoYWRvdzogdmFyKC0tcGYtdjYtYy1wYWdlLS1zZWN0aW9uLS1tLXN0aWNreS1ib3R0b20tLUJveFNoYWRvdywgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3gtc2hhZG93LS1zbS0tdG9wKSk7XFxcXG59XFxcXG5cXFxcbi5mb290ZXItZGVzY3JpcHRpb24ge1xcXFxuICBwYWRkaW5nOiAxLjVyZW07XFxcXG5cXFxcbiAgXFxcXHUwMDI2LmVtcHR5IHtcXFxcbiAgICBkaXNwbGF5OiBub25lO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbmZvb3RlciA6OnNsb3R0ZWQoW3Nsb3Q9XFxcXFxcXCJkZXNjcmlwdGlvblxcXFxcXFwiXSkge1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICBsaW5lLWhlaWdodDogMS42O1xcXFxuICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWhvdmVyKTtcXFxcbiAgICBwYWRkaW5nOiAycHggNnB4O1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDNweDtcXFxcbiAgICBmb250LWZhbWlseTogdWktbW9ub3NwYWNlLCAnQ2FzY2FkaWEgQ29kZScsICdTb3VyY2UgQ29kZSBQcm8nLCBNZW5sbywgQ29uc29sYXMsICdEZWphVnUgU2FucyBNb25vJywgbW9ub3NwYWNlO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmxvZ3Mtd3JhcHBlciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxuI2xvZy1jb250YWluZXIge1xcXFxuICBmbGV4LWdyb3c6IDE7XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG59XFxcXG5cXFxcbi5sb2ctZW50cnkge1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgYWxpZ24taXRlbXM6IGJhc2VsaW5lO1xcXFxuICBwZi12Ni1sYWJlbCB7XFxcXG4gICAgZmxleC1zaHJpbms6IDA7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmxvZy10aW1lLFxcXFxuLmxvZy1tZXNzYWdlIHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5sb2ctdGltZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZmxleC1zaHJpbms6IDA7XFxcXG4gIGZvbnQtc2l6ZTogMTFweDtcXFxcbn1cXFxcblxcXFxuLmxvZy1tZXNzYWdlIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIHdvcmQtYnJlYWs6IGJyZWFrLXdvcmQ7XFxcXG59XFxcXG5cXFxcbi8qIE5hdmlnYXRpb24gY29udGVudCAobGlnaHQgRE9NIHNsb3R0ZWQgY29udGVudCBmb3IgcGYtdjYtcGFnZS1zaWRlYmFyKSAqL1xcXFxuLm5hdi1wYWNrYWdlIHtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2IFxcXFx1MDAzZSBzdW1tYXJ5IHtcXFxcbiAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gICAgcGFkZGluZzogMC41cmVtIDFyZW07XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICAgIHRyYW5zaXRpb246IGJhY2tncm91bmQgMjAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICBtYXJnaW4tYm90dG9tOiAwLjVyZW07XFxcXG4gICAgZGlzcGxheTogZmxleDtcXFxcbiAgICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICAgIGdhcDogMC41cmVtO1xcXFxuICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0taG92ZXIpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNjo6LXdlYmtpdC1kZXRhaWxzLW1hcmtlciB7XFxcXG4gICAgICBkaXNwbGF5OiBub25lO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNjo6YmVmb3JlIHtcXFxcbiAgICAgIGNvbnRlbnQ6ICdcXFxcXFxcXDI1QjgnO1xcXFxuICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICAgICAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIDEwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjZbb3Blbl0gXFxcXHUwMDNlIHN1bW1hcnk6OmJlZm9yZSB7XFxcXG4gICAgdHJhbnNmb3JtOiByb3RhdGUoOTBkZWcpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5uYXYtZWxlbWVudCB7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxuICBtYXJnaW4taW5saW5lLXN0YXJ0OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcblxcXFxuICBzdW1tYXJ5IHtcXFxcbiAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gICAgcGFkZGluZzogMC41cmVtIDFyZW07XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIGZvbnQtZmFtaWx5OiB1aS1tb25vc3BhY2UsICdDYXNjYWRpYSBDb2RlJywgJ1NvdXJjZSBDb2RlIFBybycsIE1lbmxvLCBDb25zb2xhcywgJ0RlamFWdSBTYW5zIE1vbm8nLCBtb25vc3BhY2U7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICAgIHRyYW5zaXRpb246IGJhY2tncm91bmQgMjAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICBkaXNwbGF5OiBmbGV4O1xcXFxuICAgIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gICAgZ2FwOiAwLjVyZW07XFxcXG4gICAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1ob3Zlcik7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2Ojotd2Via2l0LWRldGFpbHMtbWFya2VyIHtcXFxcbiAgICAgIGRpc3BsYXk6IG5vbmU7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgXFxcXHUwMDI2OjpiZWZvcmUge1xcXFxuICAgICAgY29udGVudDogJ1xcXFxcXFxcMjVCOCc7XFxcXG4gICAgICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XFxcXG4gICAgICB0cmFuc2l0aW9uOiB0cmFuc2Zvcm0gMTAwbXMgY3ViaWMtYmV6aWVyKDAuNCwgMCwgMC4yLCAxKTtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1zdWJ0bGUpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNltvcGVuXSBzdW1tYXJ5OjpiZWZvcmUge1xcXFxuICAgIHRyYW5zZm9ybTogcm90YXRlKDkwZGVnKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubmF2LWVsZW1lbnQtdGl0bGUge1xcXFxuICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcbn1cXFxcblxcXFxuLm5hdi1kZW1vLWxpc3Qge1xcXFxuICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIDAgMCAwO1xcXFxuICBkaXNwbGF5OiBncmlkO1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpO1xcXFxufVxcXFxuXFxcXG4ubmF2LWRlbW8tbGluayB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxuICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBwYWRkaW5nLWlubGluZS1zdGFydDogY2FsYyh2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSAqIDIpO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1yYWRpdXMpO1xcXFxuICBkaXNwbGF5OiBibG9jaztcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kIDAuMnMsIGNvbG9yIDAuMnM7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1icmFuZCk7XFxcXG5cXFxcbiAgICAubmF2LXBhY2thZ2UtbmFtZSB7XFxcXG4gICAgICBjb2xvcjogcmdiYSgyNTUsIDI1NSwgMjU1LCAwLjgpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNlthcmlhLWN1cnJlbnQ9XFxcXFxcXCJwYWdlXFxcXFxcXCJdIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1vbi1icmFuZCk7XFxcXG5cXFxcbiAgICAubmF2LXBhY2thZ2UtbmFtZSB7XFxcXG4gICAgICBjb2xvcjogcmdiYSgyNTUsIDI1NSwgMjU1LCAwLjgpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubmF2LXBhY2thZ2UtbmFtZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLyogSW5mbyBidXR0b24gcG9wb3ZlciB0cmlnZ2VycyBpbiBrbm9icyAtIG92ZXJyaWRlIHBsYWluIGJ1dHRvbiBwYWRkaW5nICovXFxcXG5wZi12Ni1wb3BvdmVyIHBmLXY2LWJ1dHRvblt2YXJpYW50PVxcXFxcXFwicGxhaW5cXFxcXFxcIl0ge1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1tLXBsYWluLS1QYWRkaW5nSW5saW5lRW5kOiAwO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1tLXBsYWluLS1QYWRkaW5nSW5saW5lU3RhcnQ6IDA7XFxcXG4gIC0tcGYtdjYtYy1idXR0b24tLU1pbldpZHRoOiBhdXRvO1xcXFxufVxcXFxuXFxcXG4vKiBLbm9iIGRlc2NyaXB0aW9uIGNvbnRlbnQgKHNsb3R0ZWQgaW4gZm9ybSBncm91cCBoZWxwZXIpICovXFxcXG5wZi12Ni1mb3JtLWdyb3VwIFtzbG90PVxcXFxcXFwiaGVscGVyXFxcXFxcXCJdIHtcXFxcbiAgcCB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSAwO1xcXFxuICB9XFxcXG5cXFxcbiAgY29kZSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDNweDtcXFxcbiAgICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIH1cXFxcblxcXFxuICBhIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgICB0ZXh0LWRlY29yYXRpb246IG5vbmU7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgdGV4dC1kZWNvcmF0aW9uOiB1bmRlcmxpbmU7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgc3Ryb25nIHtcXFxcbiAgICBmb250LXdlaWdodDogNjAwO1xcXFxuICB9XFxcXG5cXFxcbiAgdWwsIG9sIHtcXFxcbiAgICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIDA7XFxcXG4gICAgcGFkZGluZy1sZWZ0OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLWxnKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBTeW50YXggaGlnaGxpZ2h0aW5nIChjaHJvbWEgLSB0aGVtYWJsZSB2aWEgQ1NTIGN1c3RvbSBwcm9wZXJ0aWVzKSAqL1xcXFxucGYtdjYtZm9ybS1ncm91cCBbc2xvdD1cXFxcXFxcImhlbHBlclxcXFxcXFwiXSB7XFxcXG4gIC5jaHJvbWEge1xcXFxuICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItcmFkaXVzKTtcXFxcbiAgICBvdmVyZmxvdy14OiBhdXRvO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2IC5sbnRkIHsgdmVydGljYWwtYWxpZ246IHRvcDsgcGFkZGluZzogMDsgbWFyZ2luOiAwOyBib3JkZXI6IDA7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmxudGFibGUgeyBib3JkZXItc3BhY2luZzogMDsgcGFkZGluZzogMDsgbWFyZ2luOiAwOyBib3JkZXI6IDA7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmhsIHsgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWhpZ2hsaWdodCkgfVxcXFxuICAgIFxcXFx1MDAyNiAubG50LFxcXFxuICAgIFxcXFx1MDAyNiAubG4ge1xcXFxuICAgICAgd2hpdGUtc3BhY2U6IHByZTtcXFxcbiAgICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuICAgICAgbWFyZ2luLXJpZ2h0OiAwLjRlbTtcXFxcbiAgICAgIHBhZGRpbmc6IDAgMC40ZW0gMCAwLjRlbTtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5saW5lIHsgZGlzcGxheTogZmxleDsgfVxcXFxuXFxcXG4gICAgLyogS2V5d29yZHMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmssXFxcXG4gICAgXFxcXHUwMDI2IC5rYyxcXFxcbiAgICBcXFxcdTAwMjYgLmtkLFxcXFxuICAgIFxcXFx1MDAyNiAua24sXFxcXG4gICAgXFxcXHUwMDI2IC5rcCxcXFxcbiAgICBcXFxcdTAwMjYgLmtyIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgta2V5d29yZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5rdCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtdHlwZSk7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG5cXFxcbiAgICAvKiBOYW1lcyAqL1xcXFxuICAgIFxcXFx1MDAyNiAubmEsXFxcXG4gICAgXFxcXHUwMDI2IC5uYixcXFxcbiAgICBcXFxcdTAwMjYgLm5vLFxcXFxuICAgIFxcXFx1MDAyNiAubnYsXFxcXG4gICAgXFxcXHUwMDI2IC52YyxcXFxcbiAgICBcXFxcdTAwMjYgLnZnLFxcXFxuICAgIFxcXFx1MDAyNiAudmkge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1uYW1lKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5icCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAubmMgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWNsYXNzKTsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLm5kIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1kZWNvcmF0b3IpOyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuICAgIFxcXFx1MDAyNiAubmksXFxcXG4gICAgXFxcXHUwMDI2IC5zcyB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXNwZWNpYWwpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLm5lLFxcXFxuICAgIFxcXFx1MDAyNiAubmwge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1rZXl3b3JkKTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiBib2xkO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLm5mIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1mdW5jdGlvbik7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5ubiB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAubnQgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXRhZykgfVxcXFxuXFxcXG4gICAgLyogU3RyaW5ncyAqL1xcXFxuICAgIFxcXFx1MDAyNiAucyxcXFxcbiAgICBcXFxcdTAwMjYgLnNhLFxcXFxuICAgIFxcXFx1MDAyNiAuc2IsXFxcXG4gICAgXFxcXHUwMDI2IC5zYyxcXFxcbiAgICBcXFxcdTAwMjYgLmRsLFxcXFxuICAgIFxcXFx1MDAyNiAuc2QsXFxcXG4gICAgXFxcXHUwMDI2IC5zMixcXFxcbiAgICBcXFxcdTAwMjYgLnNlLFxcXFxuICAgIFxcXFx1MDAyNiAuc2gsXFxcXG4gICAgXFxcXHUwMDI2IC5zaSxcXFxcbiAgICBcXFxcdTAwMjYgLnN4LFxcXFxuICAgIFxcXFx1MDAyNiAuczEge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1zdHJpbmcpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLnNyIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC10YWcpIH1cXFxcblxcXFxuICAgIC8qIE51bWJlcnMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLm0sXFxcXG4gICAgXFxcXHUwMDI2IC5tYixcXFxcbiAgICBcXFxcdTAwMjYgLm1mLFxcXFxuICAgIFxcXFx1MDAyNiAubWgsXFxcXG4gICAgXFxcXHUwMDI2IC5taSxcXFxcbiAgICBcXFxcdTAwMjYgLmlsLFxcXFxuICAgIFxcXFx1MDAyNiAubW8ge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1udW1iZXIpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIE9wZXJhdG9ycyAqL1xcXFxuICAgIFxcXFx1MDAyNiAubyxcXFxcbiAgICBcXFxcdTAwMjYgLm93IHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgta2V5d29yZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBDb21tZW50cyAqL1xcXFxuICAgIFxcXFx1MDAyNiAuYyxcXFxcbiAgICBcXFxcdTAwMjYgLmNoLFxcXFxuICAgIFxcXFx1MDAyNiAuY20sXFxcXG4gICAgXFxcXHUwMDI2IC5jMSB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gICAgICBmb250LXN0eWxlOiBpdGFsaWM7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuY3MsXFxcXG4gICAgXFxcXHUwMDI2IC5jcCxcXFxcbiAgICBcXFxcdTAwMjYgLmNwZiB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBFcnJvcnMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmVyciB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWVycm9yKTtcXFxcbiAgICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvci1iZyk7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgLyogR2VuZXJpY3MgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmdkIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZGVsZXRlZCk7XFxcXG4gICAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZGVsZXRlZC1iZyk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2UgeyBmb250LXN0eWxlOiBpdGFsaWM7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmdyIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvcikgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ2ggeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdpIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtaW5zZXJ0ZWQpO1xcXFxuICAgICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWluc2VydGVkLWJnKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5nbyB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3AgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdzIHsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLmd1IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5ndCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZXJyb3IpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdsIHsgdGV4dC1kZWNvcmF0aW9uOiB1bmRlcmxpbmU7IH1cXFxcbiAgICBcXFxcdTAwMjYgLncgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCkgfVxcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIEV2ZW50cyB0YWIgc3R5bGluZyAtIFByaW1hcnktZGV0YWlsIGxheW91dCAqL1xcXFxuLmV2ZW50cy13cmFwcGVyIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4jZXZlbnQtZHJhd2VyIHtcXFxcbiAgZmxleDogMTtcXFxcbiAgbWluLWhlaWdodDogMDtcXFxcbn1cXFxcblxcXFxuLyogRXZlbnQgbGlzdCAocHJpbWFyeSBwYW5lbCkgKi9cXFxcbiNldmVudC1saXN0IHtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbiAgaGVpZ2h0OiAxMDAlO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtbGlzdC1pdGVtIHtcXFxcbiAgLyogUmVzZXQgYnV0dG9uIHN0eWxlcyAqL1xcXFxuICBhcHBlYXJhbmNlOiBub25lO1xcXFxuICBiYWNrZ3JvdW5kOiBub25lO1xcXFxuICBib3JkZXI6IG5vbmU7XFxcXG4gIGJvcmRlci1sZWZ0OiAzcHggc29saWQgdHJhbnNwYXJlbnQ7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgZm9udDogaW5oZXJpdDtcXFxcbiAgY29sb3I6IGluaGVyaXQ7XFxcXG4gIHRleHQtYWxpZ246IGluaGVyaXQ7XFxcXG4gIHdpZHRoOiAxMDAlO1xcXFxuXFxcXG4gIC8qIENvbXBvbmVudCBzdHlsZXMgKi9cXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAxMDBtcyBlYXNlLWluLW91dCwgYm9yZGVyLWNvbG9yIDEwMG1zIGVhc2UtaW4tb3V0O1xcXFxuXFxcXG4gIHBmLXY2LWxhYmVsIHtcXFxcbiAgICBmbGV4LXNocmluazogMDtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0taG92ZXIpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmZvY3VzIHtcXFxcbiAgICBvdXRsaW5lOiAycHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1jbGlja2VkKTtcXFxcbiAgICBvdXRsaW5lLW9mZnNldDogLTJweDtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5zZWxlY3RlZCB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tYWN0aW9uLS1wbGFpbi0tc2VsZWN0ZWQpO1xcXFxuICAgIGJvcmRlci1sZWZ0LWNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4uZXZlbnQtdGltZSxcXFxcbi5ldmVudC1lbGVtZW50IHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC10aW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmbGV4LXNocmluazogMDtcXFxcbiAgZm9udC1zaXplOiAxMXB4O1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZWxlbWVudCB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDQwMDtcXFxcbn1cXFxcblxcXFxuLyogRXZlbnQgZGV0YWlsIHBhbmVsICovXFxcXG4uZXZlbnQtZGV0YWlsLWhlYWRlci1jb250ZW50IHtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGJvcmRlci1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci13aWR0aCkgc29saWQgdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLWNvbG9yKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1uYW1lIHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtbGcpO1xcXFxuICBmb250LXdlaWdodDogNjAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1zdW1tYXJ5IHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxuICB3aGl0ZS1zcGFjZTogcHJlLXdyYXA7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtZGVzY3JpcHRpb24ge1xcXFxuICBtYXJnaW46IDAgMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSAwO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjU7XFxcXG4gIHdoaXRlLXNwYWNlOiBwcmUtd3JhcDtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1tZXRhIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLXRpbWUge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1lbGVtZW50IHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLXByb3BlcnRpZXMtaGVhZGluZyB7XFxcXG4gIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLWJhc2UpO1xcXFxuICBmb250LXdlaWdodDogNjAwO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1wcm9wZXJ0aWVzIHtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXNlY29uZGFyeSk7XFxcXG4gIGJvcmRlcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXdpZHRoKSBzb2xpZCB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItY29sb3IpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItcmFkaXVzKTtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IDEycHg7XFxcXG4gIGxpbmUtaGVpZ2h0OiAxLjY7XFxcXG4gIG1hcmdpbjogMCB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LXByb3BlcnR5LXRyZWUge1xcXFxuICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBtYXJnaW46IDA7XFxcXG5cXFxcbiAgXFxcXHUwMDI2Lm5lc3RlZCB7XFxcXG4gICAgcGFkZGluZy1sZWZ0OiAxLjVlbTtcXFxcbiAgICBtYXJnaW4tdG9wOiAwLjI1ZW07XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LWl0ZW0ge1xcXFxuICBwYWRkaW5nOiAwLjEyNWVtIDA7XFxcXG59XFxcXG5cXFxcbi5wcm9wZXJ0eS1rZXkge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDUwMDtcXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LWNvbG9uIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxufVxcXFxuXFxcXG4ucHJvcGVydHktdmFsdWUge1xcXFxuICBcXFxcdTAwMjYubnVsbCxcXFxcbiAgXFxcXHUwMDI2LnVuZGVmaW5lZCB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICAgIGZvbnQtc3R5bGU6IGl0YWxpYztcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5ib29sZWFuIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItY29sb3ItYm9vbGVhbik7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYubnVtYmVyIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItY29sb3ItbnVtYmVyKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5zdHJpbmcge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1jb2xvci1zdHJpbmcpO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2LmFycmF5LFxcXFxuICBcXFxcdTAwMjYub2JqZWN0IHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbiNkZWJ1Zy1tb2RhbCB7XFxcXG4gIGNvbnRhaW5lci10eXBlOiBpbmxpbmUtc2l6ZTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLE1BQU0sZUFBZTtBQUMxQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjs7O0FDRnpCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLG9vc0JB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
