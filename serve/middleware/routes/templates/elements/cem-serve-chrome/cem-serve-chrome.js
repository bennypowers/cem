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
import { CEMReloadClient } from "/__cem/websocket-client.js";
import { StatePersistence } from "/__cem/state-persistence.js";
import "/__cem/health-badges.js";
var CemLogsEvent = class extends Event {
  logs;
  constructor(logs) {
    super("cem:logs", { bubbles: true });
    this.logs = logs;
  }
};
var _hasDescription_dec, _sidebar_dec, _tabsSelected_dec, _drawerHeight_dec, _drawer_dec, _knobs_dec, _sourceURL_dec, _canonicalURL_dec, _packageName_dec, _demoTitle_dec, _primaryTagName_dec, _a, _CemServeChrome_decorators, _demoInfoTemplate, _demoGroupTemplate, _demoListTemplate, _logEntryTemplate, _eventEntryTemplate, _init, _primaryTagName, _demoTitle, _packageName, _canonicalURL, _sourceURL, _knobs, _drawer, _drawerHeight, _tabsSelected, _sidebar, _hasDescription, _CemServeChrome_instances, $_fn, $$_fn, _logContainer, _drawerOpen, _initialLogsFetched, _isInitialLoad, _debugData, _elementEventMap, _manifest, _capturedEvents, _maxCapturedEvents, _eventList, _eventDetailHeader, _eventDetailBody, _selectedEventId, _eventsFilterValue, _eventsFilterDebounceTimer, _eventTypeFilters, _elementFilters, _discoveredElements, _handleLogsEvent, _handleTreeExpand, _handleTreeCollapse, _handleTreeSelect, _copyLogsFeedbackTimeout, _copyDebugFeedbackTimeout, _copyEventsFeedbackTimeout, _logsFilterValue, _logsFilterDebounceTimer, _logLevelFilters, _logLevelDropdown, _observer, _wsClient, renderSourceButton_fn, fetchDebugInfo_fn, populateDemoUrls_fn, setupLogListener_fn, filterLogs_fn, getLogTypeFromEntry_fn, loadLogFilterState_fn, syncCheckboxStates_fn, saveLogFilterState_fn, _handleLogFilterChange, copyLogs_fn, setupDebugOverlay_fn, setupFooterDrawer_fn, detectBrowser_fn, copyDebugInfo_fn, renderLogs_fn, getLogBadge_fn, applyLogLabelAttrs_fn, scrollLatestIntoView_fn, scrollLogsToBottom_fn, migrateFromLocalStorageIfNeeded_fn, setupColorSchemeToggle_fn, applyColorScheme_fn, setupKnobCoordination_fn, _onKnobChange, _onKnobClear, getKnobTarget_fn, getKnobTypeFromEvent_fn, getKnobTypeFromClearEvent_fn, setupTreeStatePersistence_fn, applyTreeState_fn, setupSidebarStatePersistence_fn, findTreeItemById_fn, getTreeNodeId_fn, discoverElementEvents_fn, setupEventCapture_fn, attachEventListeners_fn, observeDemoMutations_fn, _handleElementEvent, getEventDocumentation_fn, findTypeDeclaration_fn, captureEvent_fn, extractEventProperties_fn, renderEvent_fn, selectEvent_fn, buildPropertiesForDisplay_fn, buildPropertyTree_fn, scrollEventsToBottom_fn, isEventsTabActive_fn, filterEvents_fn, eventMatchesTextFilter_fn, getEventRecordById_fn, updateEventFilters_fn, _handleEventTypeFilterChange, _handleElementFilterChange, loadEventFiltersFromStorage_fn, saveEventFilters_fn, clearEvents_fn, copyEvents_fn, setupEventListeners_fn;
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
    __privateAdd(this, _wsClient, new CEMReloadClient({
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
  connectedCallback() {
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvY2VtLXNlcnZlLWNocm9tZS9jZW0tc2VydmUtY2hyb21lLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1zZXJ2ZS1jaHJvbWUuY3NzJztcblxuaW1wb3J0ICcuLi9jZW0tY29sb3Itc2NoZW1lLXRvZ2dsZS9jZW0tY29sb3Itc2NoZW1lLXRvZ2dsZS5qcyc7XG5pbXBvcnQgJy4uL2NlbS1kcmF3ZXIvY2VtLWRyYXdlci5qcyc7XG5pbXBvcnQgJy4uL2NlbS1oZWFsdGgtcGFuZWwvY2VtLWhlYWx0aC1wYW5lbC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1tYW5pZmVzdC1icm93c2VyL2NlbS1tYW5pZmVzdC1icm93c2VyLmpzJztcbmltcG9ydCAnLi4vY2VtLXJlY29ubmVjdGlvbi1jb250ZW50L2NlbS1yZWNvbm5lY3Rpb24tY29udGVudC5qcyc7XG5pbXBvcnQgJy4uL2NlbS1zZXJ2ZS1kZW1vL2NlbS1zZXJ2ZS1kZW1vLmpzJztcbmltcG9ydCAnLi4vY2VtLXNlcnZlLWtub2ItZ3JvdXAvY2VtLXNlcnZlLWtub2ItZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9jZW0tc2VydmUta25vYnMvY2VtLXNlcnZlLWtub2JzLmpzJztcbmltcG9ydCAnLi4vY2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5L2NlbS10cmFuc2Zvcm0tZXJyb3Itb3ZlcmxheS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWFsZXJ0L3BmLXY2LWFsZXJ0LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtYWxlcnQtZ3JvdXAvcGYtdjYtYWxlcnQtZ3JvdXAuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1idXR0b24vcGYtdjYtYnV0dG9uLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtY2FyZC9wZi12Ni1jYXJkLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtYmFkZ2UvcGYtdjYtYmFkZ2UuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1kcm9wZG93bi9wZi12Ni1kcm9wZG93bi5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWV4cGFuZGFibGUtc2VjdGlvbi9wZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1sYWJlbC9wZi12Ni1sYWJlbC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW1hc3RoZWFkL3BmLXY2LW1hc3RoZWFkLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbW9kYWwvcGYtdjYtbW9kYWwuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXYtZ3JvdXAvcGYtdjYtbmF2LWdyb3VwLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbmF2LWl0ZW0vcGYtdjYtbmF2LWl0ZW0uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXYtbGluay9wZi12Ni1uYXYtbGluay5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW5hdi1saXN0L3BmLXY2LW5hdi1saXN0LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbmF2aWdhdGlvbi9wZi12Ni1uYXZpZ2F0aW9uLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtcGFnZS1tYWluL3BmLXY2LXBhZ2UtbWFpbi5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXBhZ2Utc2lkZWJhci9wZi12Ni1wYWdlLXNpZGViYXIuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1wYWdlL3BmLXY2LXBhZ2UuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1wb3BvdmVyL3BmLXY2LXBvcG92ZXIuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1zZWxlY3QvcGYtdjYtc2VsZWN0LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtc2tpcC10by1jb250ZW50L3BmLXY2LXNraXAtdG8tY29udGVudC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXN3aXRjaC9wZi12Ni1zd2l0Y2guanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10YWIvcGYtdjYtdGFiLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdGFicy9wZi12Ni10YWJzLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdGV4dC1pbnB1dC1ncm91cC9wZi12Ni10ZXh0LWlucHV0LWdyb3VwLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdGV4dC1pbnB1dC9wZi12Ni10ZXh0LWlucHV0LmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0vcGYtdjYtdG9nZ2xlLWdyb3VwLWl0ZW0uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10b2dnbGUtZ3JvdXAvcGYtdjYtdG9nZ2xlLWdyb3VwLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdG9vbGJhci1ncm91cC9wZi12Ni10b29sYmFyLWdyb3VwLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtdG9vbGJhci1pdGVtL3BmLXY2LXRvb2xiYXItaXRlbS5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRvb2xiYXIvcGYtdjYtdG9vbGJhci5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LXRyZWUtaXRlbS9wZi12Ni10cmVlLWl0ZW0uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni10cmVlLXZpZXcvcGYtdjYtdHJlZS12aWV3LmpzJztcblxuLy8gQHRzLWlnbm9yZSAtLSBwbGFpbiBKUyBtb2R1bGUgc2VydmVkIGF0IHJ1bnRpbWUgYnkgR28gc2VydmVyXG5pbXBvcnQgeyBDRU1SZWxvYWRDbGllbnQgfSBmcm9tICcvX19jZW0vd2Vic29ja2V0LWNsaWVudC5qcyc7XG4vLyBAdHMtaWdub3JlIC0tIHBsYWluIEpTIG1vZHVsZSBzZXJ2ZWQgYXQgcnVudGltZSBieSBHbyBzZXJ2ZXJcbmltcG9ydCB7IFN0YXRlUGVyc2lzdGVuY2UgfSBmcm9tICcvX19jZW0vc3RhdGUtcGVyc2lzdGVuY2UuanMnO1xuLy8gQHRzLWlnbm9yZSAtLSBwbGFpbiBKUyBtb2R1bGUgc2VydmVkIGF0IHJ1bnRpbWUgYnkgR28gc2VydmVyXG5pbXBvcnQgJy9fX2NlbS9oZWFsdGgtYmFkZ2VzLmpzJztcblxuaW50ZXJmYWNlIEV2ZW50SW5mbyB7XG4gIGV2ZW50TmFtZXM6IFNldDxzdHJpbmc+O1xuICBldmVudHM6IEFycmF5PHsgbmFtZTogc3RyaW5nOyB0eXBlPzogeyB0ZXh0OiBzdHJpbmcgfTsgc3VtbWFyeT86IHN0cmluZzsgZGVzY3JpcHRpb24/OiBzdHJpbmcgfT47XG59XG5cbmludGVyZmFjZSBFdmVudFJlY29yZCB7XG4gIGlkOiBzdHJpbmc7XG4gIHRpbWVzdGFtcDogRGF0ZTtcbiAgZXZlbnROYW1lOiBzdHJpbmc7XG4gIHRhZ05hbWU6IHN0cmluZztcbiAgZWxlbWVudElkOiBzdHJpbmcgfCBudWxsO1xuICBlbGVtZW50Q2xhc3M6IHN0cmluZyB8IG51bGw7XG4gIGN1c3RvbVByb3BlcnRpZXM6IFJlY29yZDxzdHJpbmcsIHVua25vd24+O1xuICBtYW5pZmVzdFR5cGU6IHN0cmluZyB8IG51bGw7XG4gIHN1bW1hcnk6IHN0cmluZyB8IG51bGw7XG4gIGRlc2NyaXB0aW9uOiBzdHJpbmcgfCBudWxsO1xuICBidWJibGVzOiBib29sZWFuO1xuICBjb21wb3NlZDogYm9vbGVhbjtcbiAgY2FuY2VsYWJsZTogYm9vbGVhbjtcbiAgZGVmYXVsdFByZXZlbnRlZDogYm9vbGVhbjtcbn1cblxuaW50ZXJmYWNlIERlYnVnRGF0YSB7XG4gIHZlcnNpb24/OiBzdHJpbmc7XG4gIG9zPzogc3RyaW5nO1xuICB3YXRjaERpcj86IHN0cmluZztcbiAgbWFuaWZlc3RTaXplPzogc3RyaW5nO1xuICBkZW1vQ291bnQ/OiBudW1iZXI7XG4gIGRlbW9zPzogQXJyYXk8e1xuICAgIHRhZ05hbWU6IHN0cmluZztcbiAgICBkZXNjcmlwdGlvbj86IHN0cmluZztcbiAgICBmaWxlcGF0aD86IHN0cmluZztcbiAgICBjYW5vbmljYWxVUkw6IHN0cmluZztcbiAgICBsb2NhbFJvdXRlOiBzdHJpbmc7XG4gIH0+O1xuICBpbXBvcnRNYXA/OiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPjtcbiAgaW1wb3J0TWFwSlNPTj86IHN0cmluZztcbn1cblxuaW50ZXJmYWNlIE1hbmlmZXN0IHtcbiAgbW9kdWxlcz86IEFycmF5PHtcbiAgICBkZWNsYXJhdGlvbnM/OiBBcnJheTx7XG4gICAgICBjdXN0b21FbGVtZW50PzogYm9vbGVhbjtcbiAgICAgIHRhZ05hbWU/OiBzdHJpbmc7XG4gICAgICBuYW1lPzogc3RyaW5nO1xuICAgICAga2luZD86IHN0cmluZztcbiAgICAgIGV2ZW50cz86IEFycmF5PHsgbmFtZTogc3RyaW5nOyB0eXBlPzogeyB0ZXh0OiBzdHJpbmcgfTsgc3VtbWFyeT86IHN0cmluZzsgZGVzY3JpcHRpb24/OiBzdHJpbmcgfT47XG4gICAgfT47XG4gIH0+O1xufVxuXG4vKipcbiAqIEN1c3RvbSBldmVudCBmaXJlZCB3aGVuIGxvZ3MgYXJlIHJlY2VpdmVkXG4gKi9cbmV4cG9ydCBjbGFzcyBDZW1Mb2dzRXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIGxvZ3M6IEFycmF5PHsgdHlwZTogc3RyaW5nOyBkYXRlOiBzdHJpbmc7IG1lc3NhZ2U6IHN0cmluZyB9PjtcbiAgY29uc3RydWN0b3IobG9nczogQXJyYXk8eyB0eXBlOiBzdHJpbmc7IGRhdGU6IHN0cmluZzsgbWVzc2FnZTogc3RyaW5nIH0+KSB7XG4gICAgc3VwZXIoJ2NlbTpsb2dzJywgeyBidWJibGVzOiB0cnVlIH0pO1xuICAgIHRoaXMubG9ncyA9IGxvZ3M7XG4gIH1cbn1cblxuLyoqXG4gKiBDRU0gU2VydmUgQ2hyb21lIC0gTWFpbiBkZW1vIHdyYXBwZXIgY29tcG9uZW50XG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGRlbW8gY29udGVudFxuICogQHNsb3QgbmF2aWdhdGlvbiAtIE5hdmlnYXRpb24gc2lkZWJhciBjb250ZW50XG4gKiBAc2xvdCBrbm9icyAtIEtub2IgY29udHJvbHNcbiAqIEBzbG90IGRlc2NyaXB0aW9uIC0gRGVtbyBkZXNjcmlwdGlvblxuICogQHNsb3QgbWFuaWZlc3QtdHJlZSAtIE1hbmlmZXN0IHRyZWUgdmlld1xuICogQHNsb3QgbWFuaWZlc3QtbmFtZSAtIE1hbmlmZXN0IG5hbWUgZGlzcGxheVxuICogQHNsb3QgbWFuaWZlc3QtZGV0YWlscyAtIE1hbmlmZXN0IGRldGFpbHMgZGlzcGxheVxuICpcbiAqIEBjdXN0b21FbGVtZW50IGNlbS1zZXJ2ZS1jaHJvbWVcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1zZXJ2ZS1jaHJvbWUnKVxuZXhwb3J0IGNsYXNzIENlbVNlcnZlQ2hyb21lIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgLy8gU3RhdGljIHRlbXBsYXRlcyBmb3IgRE9NIGNsb25pbmcgKGxvZ3MsIGV2ZW50cywgZGVidWcgaW5mbylcbiAgc3RhdGljICNkZW1vSW5mb1RlbXBsYXRlID0gKCgpID0+IHtcbiAgICBjb25zdCB0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGVtcGxhdGUnKTtcbiAgICB0LmlubmVySFRNTCA9IGBcbiAgICAgIDxoMz5EZW1vIEluZm9ybWF0aW9uPC9oMz5cbiAgICAgIDxkbCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdCBwZi1tLWhvcml6b250YWwgcGYtbS1jb21wYWN0XCI+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+VGFnIE5hbWU8L2R0PlxuICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBkYXRhLWZpZWxkPVwidGFnTmFtZVwiPjwvZGQ+XG4gICAgICAgIDwvZGl2PlxuICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiIGRhdGEtZmllbGQtZ3JvdXA9XCJkZXNjcmlwdGlvblwiPlxuICAgICAgICAgIDxkdCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fdGVybVwiPkRlc2NyaXB0aW9uPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cImRlc2NyaXB0aW9uXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCIgZGF0YS1maWVsZC1ncm91cD1cImZpbGVwYXRoXCI+XG4gICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+RmlsZSBQYXRoPC9kdD5cbiAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgZGF0YS1maWVsZD1cImZpbGVwYXRoXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+Q2Fub25pY2FsIFVSTDwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJjYW5vbmljYWxVUkxcIj48L2RkPlxuICAgICAgICA8L2Rpdj5cbiAgICAgICAgPGRpdiBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZ3JvdXBcIj5cbiAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5Mb2NhbCBSb3V0ZTwvZHQ+XG4gICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGRhdGEtZmllbGQ9XCJsb2NhbFJvdXRlXCI+PC9kZD5cbiAgICAgICAgPC9kaXY+XG4gICAgICA8L2RsPmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgc3RhdGljICNkZW1vR3JvdXBUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIiBkYXRhLWZpZWxkPVwidGFnTmFtZVwiPjwvZHQ+XG4gICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIj5cbiAgICAgICAgICA8c3BhbiBkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIj48L3NwYW4+PGJyPlxuICAgICAgICAgIDxzbWFsbCBkYXRhLWZpZWxkLWdyb3VwPVwiZmlsZXBhdGhcIj5GaWxlOiA8c3BhbiBkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIj48L3NwYW4+PC9zbWFsbD5cbiAgICAgICAgICA8c21hbGw+Q2Fub25pY2FsOiA8c3BhbiBkYXRhLWZpZWxkPVwiY2Fub25pY2FsVVJMXCI+PC9zcGFuPjwvc21hbGw+PGJyPlxuICAgICAgICAgIDxzbWFsbD5Mb2NhbDogPHNwYW4gZGF0YS1maWVsZD1cImxvY2FsUm91dGVcIj48L3NwYW4+PC9zbWFsbD5cbiAgICAgICAgPC9kZD5cbiAgICAgIDwvZGl2PmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgc3RhdGljICNkZW1vTGlzdFRlbXBsYXRlID0gKCgpID0+IHtcbiAgICBjb25zdCB0ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndGVtcGxhdGUnKTtcbiAgICB0LmlubmVySFRNTCA9IGBcbiAgICAgIDxwZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24gaWQ9XCJkZWJ1Zy1kZW1vcy1zZWN0aW9uXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdG9nZ2xlLXRleHQ9XCJTaG93IERlbW9zIEluZm9cIj5cbiAgICAgICAgPGRsIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0IHBmLW0taG9yaXpvbnRhbCBwZi1tLWNvbXBhY3RcIiBkYXRhLWNvbnRhaW5lcj1cImdyb3Vwc1wiPjwvZGw+XG4gICAgICA8L3BmLXY2LWV4cGFuZGFibGUtc2VjdGlvbj5gO1xuICAgIHJldHVybiB0O1xuICB9KSgpO1xuXG4gIHN0YXRpYyAjbG9nRW50cnlUZW1wbGF0ZSA9ICgoKSA9PiB7XG4gICAgY29uc3QgdCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RlbXBsYXRlJyk7XG4gICAgdC5pbm5lckhUTUwgPSBgXG4gICAgICA8ZGl2IGNsYXNzPVwibG9nLWVudHJ5XCIgZGF0YS1maWVsZD1cImNvbnRhaW5lclwiPlxuICAgICAgICA8cGYtdjYtbGFiZWwgY29tcGFjdCBkYXRhLWZpZWxkPVwibGFiZWxcIj48L3BmLXY2LWxhYmVsPlxuICAgICAgICA8dGltZSBjbGFzcz1cImxvZy10aW1lXCIgZGF0YS1maWVsZD1cInRpbWVcIj48L3RpbWU+XG4gICAgICAgIDxzcGFuIGNsYXNzPVwibG9nLW1lc3NhZ2VcIiBkYXRhLWZpZWxkPVwibWVzc2FnZVwiPjwvc3Bhbj5cbiAgICAgIDwvZGl2PmA7XG4gICAgcmV0dXJuIHQ7XG4gIH0pKCk7XG5cbiAgc3RhdGljICNldmVudEVudHJ5VGVtcGxhdGUgPSAoKCkgPT4ge1xuICAgIGNvbnN0IHQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCd0ZW1wbGF0ZScpO1xuICAgIHQuaW5uZXJIVE1MID0gYFxuICAgICAgPGJ1dHRvbiBjbGFzcz1cImV2ZW50LWxpc3QtaXRlbVwiIGRhdGEtZmllbGQ9XCJjb250YWluZXJcIj5cbiAgICAgICAgPHBmLXY2LWxhYmVsIGNvbXBhY3QgZGF0YS1maWVsZD1cImxhYmVsXCI+PC9wZi12Ni1sYWJlbD5cbiAgICAgICAgPHRpbWUgY2xhc3M9XCJldmVudC10aW1lXCIgZGF0YS1maWVsZD1cInRpbWVcIj48L3RpbWU+XG4gICAgICAgIDxzcGFuIGNsYXNzPVwiZXZlbnQtZWxlbWVudFwiIGRhdGEtZmllbGQ9XCJlbGVtZW50XCI+PC9zcGFuPlxuICAgICAgPC9idXR0b24+YDtcbiAgICByZXR1cm4gdDtcbiAgfSkoKTtcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdwcmltYXJ5LXRhZy1uYW1lJyB9KVxuICBhY2Nlc3NvciBwcmltYXJ5VGFnTmFtZSA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ2RlbW8tdGl0bGUnIH0pXG4gIGFjY2Vzc29yIGRlbW9UaXRsZSA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IGF0dHJpYnV0ZTogJ3BhY2thZ2UtbmFtZScgfSlcbiAgYWNjZXNzb3IgcGFja2FnZU5hbWUgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdjYW5vbmljYWwtdXJsJyB9KVxuICBhY2Nlc3NvciBjYW5vbmljYWxVUkwgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdzb3VyY2UtdXJsJyB9KVxuICBhY2Nlc3NvciBzb3VyY2VVUkwgPSAnJztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBrbm9icyA9ICcnO1xuXG4gIEBwcm9wZXJ0eSgpXG4gIGFjY2Vzc29yIGRyYXdlcjogJ2V4cGFuZGVkJyB8ICdjb2xsYXBzZWQnIHwgJycgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICdkcmF3ZXItaGVpZ2h0JyB9KVxuICBhY2Nlc3NvciBkcmF3ZXJIZWlnaHQgPSAnJztcblxuICBAcHJvcGVydHkoeyBhdHRyaWJ1dGU6ICd0YWJzLXNlbGVjdGVkJyB9KVxuICBhY2Nlc3NvciB0YWJzU2VsZWN0ZWQgPSAnJztcblxuICBAcHJvcGVydHkoKVxuICBhY2Nlc3NvciBzaWRlYmFyOiAnZXhwYW5kZWQnIHwgJ2NvbGxhcHNlZCcgfCAnJyA9ICcnO1xuXG4gIEBwcm9wZXJ0eSh7IHR5cGU6IEJvb2xlYW4sIGF0dHJpYnV0ZTogJ2hhcy1kZXNjcmlwdGlvbicgfSlcbiAgYWNjZXNzb3IgaGFzRGVzY3JpcHRpb24gPSBmYWxzZTtcblxuICAjJChpZDogc3RyaW5nKSB7XG4gICAgcmV0dXJuIHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoaWQpO1xuICB9XG5cbiAgIyQkKHNlbGVjdG9yOiBzdHJpbmcpIHtcbiAgICByZXR1cm4gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yQWxsKHNlbGVjdG9yKSA/PyBbXTtcbiAgfVxuXG4gICNsb2dDb250YWluZXI6IEhUTUxFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNkcmF3ZXJPcGVuID0gZmFsc2U7XG4gICNpbml0aWFsTG9nc0ZldGNoZWQgPSBmYWxzZTtcbiAgI2lzSW5pdGlhbExvYWQgPSB0cnVlO1xuICAjZGVidWdEYXRhOiBEZWJ1Z0RhdGEgfCBudWxsID0gbnVsbDtcblxuICAvLyBFbGVtZW50IGV2ZW50IHRyYWNraW5nXG4gICNlbGVtZW50RXZlbnRNYXA6IE1hcDxzdHJpbmcsIEV2ZW50SW5mbz4gfCBudWxsID0gbnVsbDtcbiAgI21hbmlmZXN0OiBNYW5pZmVzdCB8IG51bGwgPSBudWxsO1xuICAjY2FwdHVyZWRFdmVudHM6IEV2ZW50UmVjb3JkW10gPSBbXTtcbiAgI21heENhcHR1cmVkRXZlbnRzID0gMTAwMDtcbiAgI2V2ZW50TGlzdDogSFRNTEVsZW1lbnQgfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50RGV0YWlsSGVhZGVyOiBIVE1MRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjZXZlbnREZXRhaWxCb2R5OiBIVE1MRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjc2VsZWN0ZWRFdmVudElkOiBzdHJpbmcgfCBudWxsID0gbnVsbDtcbiAgI2V2ZW50c0ZpbHRlclZhbHVlID0gJyc7XG4gICNldmVudHNGaWx0ZXJEZWJvdW5jZVRpbWVyOiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuICAjZXZlbnRUeXBlRmlsdGVycyA9IG5ldyBTZXQ8c3RyaW5nPigpO1xuICAjZWxlbWVudEZpbHRlcnMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgI2Rpc2NvdmVyZWRFbGVtZW50cyA9IG5ldyBTZXQ8c3RyaW5nPigpO1xuXG4gIC8vIEV2ZW50IGxpc3RlbmVyIHJlZmVyZW5jZXMgZm9yIGNsZWFudXBcbiAgI2hhbmRsZUxvZ3NFdmVudDogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG4gICNoYW5kbGVUcmVlRXhwYW5kOiAoKGV2ZW50OiBFdmVudCkgPT4gdm9pZCkgfCBudWxsID0gbnVsbDtcbiAgI2hhbmRsZVRyZWVDb2xsYXBzZTogKChldmVudDogRXZlbnQpID0+IHZvaWQpIHwgbnVsbCA9IG51bGw7XG4gICNoYW5kbGVUcmVlU2VsZWN0OiAoKGV2ZW50OiBFdmVudCkgPT4gdm9pZCkgfCBudWxsID0gbnVsbDtcblxuICAvLyBUaW1lb3V0IElEcyBmb3IgY2xlYW51cFxuICAjY29weUxvZ3NGZWVkYmFja1RpbWVvdXQ6IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+IHwgbnVsbCA9IG51bGw7XG4gICNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQ6IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+IHwgbnVsbCA9IG51bGw7XG4gICNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0OiBSZXR1cm5UeXBlPHR5cGVvZiBzZXRUaW1lb3V0PiB8IG51bGwgPSBudWxsO1xuXG4gIC8vIExvZyBmaWx0ZXIgc3RhdGVcbiAgI2xvZ3NGaWx0ZXJWYWx1ZSA9ICcnO1xuICAjbG9nc0ZpbHRlckRlYm91bmNlVGltZXI6IFJldHVyblR5cGU8dHlwZW9mIHNldFRpbWVvdXQ+IHwgbnVsbCA9IG51bGw7XG4gICNsb2dMZXZlbEZpbHRlcnMgPSBuZXcgU2V0KFsnaW5mbycsICd3YXJuJywgJ2Vycm9yJywgJ2RlYnVnJ10pO1xuICAjbG9nTGV2ZWxEcm9wZG93bjogRWxlbWVudCB8IG51bGwgPSBudWxsO1xuXG4gIC8vIFdhdGNoIGZvciBkeW5hbWljYWxseSBhZGRlZCBlbGVtZW50c1xuICAvKiBjOCBpZ25vcmUgc3RhcnQgLSBNdXRhdGlvbk9ic2VydmVyIGNhbGxiYWNrIHRlc3RlZCB2aWEgaW50ZWdyYXRpb24gKi9cbiAgI29ic2VydmVyID0gbmV3IE11dGF0aW9uT2JzZXJ2ZXIoKG11dGF0aW9ucykgPT4ge1xuICAgIGxldCBuZWVkc1VwZGF0ZSA9IGZhbHNlO1xuXG4gICAgZm9yIChjb25zdCBtdXRhdGlvbiBvZiBtdXRhdGlvbnMpIHtcbiAgICAgIGZvciAoY29uc3Qgbm9kZSBvZiBtdXRhdGlvbi5hZGRlZE5vZGVzKSB7XG4gICAgICAgIGlmIChub2RlIGluc3RhbmNlb2YgSFRNTEVsZW1lbnQpIHtcbiAgICAgICAgICBjb25zdCB0YWdOYW1lID0gbm9kZS50YWdOYW1lLnRvTG93ZXJDYXNlKCk7XG4gICAgICAgICAgaWYgKHRoaXMuI2VsZW1lbnRFdmVudE1hcD8uaGFzKHRhZ05hbWUpICYmICFub2RlLmRhdGFzZXQuY2VtRXZlbnRzQXR0YWNoZWQpIHtcbiAgICAgICAgICAgIGNvbnN0IGV2ZW50SW5mbyA9IHRoaXMuI2VsZW1lbnRFdmVudE1hcC5nZXQodGFnTmFtZSkhO1xuICAgICAgICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgZXZlbnRJbmZvLmV2ZW50TmFtZXMpIHtcbiAgICAgICAgICAgICAgbm9kZS5hZGRFdmVudExpc3RlbmVyKGV2ZW50TmFtZSwgdGhpcy4jaGFuZGxlRWxlbWVudEV2ZW50LCB7IGNhcHR1cmU6IHRydWUgfSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgICBub2RlLmRhdGFzZXQuY2VtRXZlbnRzQXR0YWNoZWQgPSAndHJ1ZSc7XG4gICAgICAgICAgICB0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuYWRkKHRhZ05hbWUpO1xuICAgICAgICAgICAgbmVlZHNVcGRhdGUgPSB0cnVlO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIGlmIChuZWVkc1VwZGF0ZSkge1xuICAgICAgdGhpcy4jdXBkYXRlRXZlbnRGaWx0ZXJzKCk7XG4gICAgfVxuICB9KTtcbiAgLyogYzggaWdub3JlIHN0b3AgKi9cblxuICAjd3NDbGllbnQgPSBuZXcgQ0VNUmVsb2FkQ2xpZW50KHtcbiAgICBqaXR0ZXJNYXg6IDEwMDAsXG4gICAgb3ZlcmxheVRocmVzaG9sZDogMTUsXG4gICAgYmFkZ2VGYWRlRGVsYXk6IDIwMDAsXG4gICAgLyogYzggaWdub3JlIHN0YXJ0IC0gV2ViU29ja2V0IGNhbGxiYWNrcyB0ZXN0ZWQgdmlhIGludGVncmF0aW9uICovXG4gICAgY2FsbGJhY2tzOiB7XG4gICAgICBvbk9wZW46ICgpID0+IHtcbiAgICAgICAgdGhpcy4jJCgncmVjb25uZWN0aW9uLW1vZGFsJyk/LmNsb3NlKCk7XG4gICAgICB9LFxuICAgICAgb25FcnJvcjogKGVycm9yRGF0YTogeyB0aXRsZT86IHN0cmluZzsgbWVzc2FnZT86IHN0cmluZzsgZmlsZT86IHN0cmluZyB9KSA9PiB7XG4gICAgICAgIGlmIChlcnJvckRhdGE/LnRpdGxlICYmIGVycm9yRGF0YT8ubWVzc2FnZSkge1xuICAgICAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmVdIFNlcnZlciBlcnJvcjonLCBlcnJvckRhdGEpO1xuICAgICAgICAgICh0aGlzLiMkKCdlcnJvci1vdmVybGF5JykgYXMgYW55KT8uc2hvdyhlcnJvckRhdGEudGl0bGUsIGVycm9yRGF0YS5tZXNzYWdlLCBlcnJvckRhdGEuZmlsZSk7XG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgY29uc29sZS5lcnJvcignW2NlbS1zZXJ2ZV0gV2ViU29ja2V0IGVycm9yOicsIGVycm9yRGF0YSk7XG4gICAgICAgIH1cbiAgICAgIH0sXG4gICAgICBvblJlY29ubmVjdGluZzogKHsgYXR0ZW1wdCwgZGVsYXkgfTogeyBhdHRlbXB0OiBudW1iZXI7IGRlbGF5OiBudW1iZXIgfSkgPT4ge1xuICAgICAgICBpZiAoYXR0ZW1wdCA+PSAxNSkge1xuICAgICAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tbW9kYWwnKSBhcyBhbnkpPy5zaG93TW9kYWwoKTtcbiAgICAgICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLWNvbnRlbnQnKSBhcyBhbnkpPy51cGRhdGVSZXRyeUluZm8oYXR0ZW1wdCwgZGVsYXkpO1xuICAgICAgICB9XG4gICAgICB9LFxuICAgICAgb25SZWxvYWQ6ICgpID0+IHtcbiAgICAgICAgY29uc3QgZXJyb3JPdmVybGF5ID0gdGhpcy4jJCgnZXJyb3Itb3ZlcmxheScpO1xuICAgICAgICBpZiAoZXJyb3JPdmVybGF5Py5oYXNBdHRyaWJ1dGUoJ29wZW4nKSkge1xuICAgICAgICAgIChlcnJvck92ZXJsYXkgYXMgYW55KS5oaWRlKCk7XG4gICAgICAgIH1cbiAgICAgICAgd2luZG93LmxvY2F0aW9uLnJlbG9hZCgpO1xuICAgICAgfSxcbiAgICAgIG9uU2h1dGRvd246ICgpID0+IHtcbiAgICAgICAgKHRoaXMuIyQoJ3JlY29ubmVjdGlvbi1tb2RhbCcpIGFzIGFueSk/LnNob3dNb2RhbCgpO1xuICAgICAgICAodGhpcy4jJCgncmVjb25uZWN0aW9uLWNvbnRlbnQnKSBhcyBhbnkpPy51cGRhdGVSZXRyeUluZm8oMzAsIDMwMDAwKTtcbiAgICAgIH0sXG4gICAgICBvbkxvZ3M6IChsb2dzOiBBcnJheTx7IHR5cGU6IHN0cmluZzsgZGF0ZTogc3RyaW5nOyBtZXNzYWdlOiBzdHJpbmcgfT4pID0+IHtcbiAgICAgICAgd2luZG93LmRpc3BhdGNoRXZlbnQobmV3IENlbUxvZ3NFdmVudChsb2dzKSk7XG4gICAgICB9XG4gICAgfVxuICAgIC8qIGM4IGlnbm9yZSBzdG9wICovXG4gIH0pO1xuXG4gIGdldCBkZW1vKCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICByZXR1cm4gdGhpcy5xdWVyeVNlbGVjdG9yKCdjZW0tc2VydmUtZGVtbycpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIHJldHVybiBodG1sYFxuICAgICAgPGxpbmsgcmVsPVwic3R5bGVzaGVldFwiIGhyZWY9XCIvX19jZW0vcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0LmNzc1wiPlxuICAgICAgPGxpbmsgcmVsPVwic3R5bGVzaGVldFwiIGhyZWY9XCIvX19jZW0vcGYtbGlnaHRkb20uY3NzXCI+XG5cbiAgICAgIDxwZi12Ni1wYWdlID9zaWRlYmFyLWNvbGxhcHNlZD0ke3RoaXMuc2lkZWJhciA9PT0gJ2NvbGxhcHNlZCd9PlxuICAgICAgICA8cGYtdjYtc2tpcC10by1jb250ZW50IHNsb3Q9XCJza2lwLXRvLWNvbnRlbnRcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhyZWY9XCIjbWFpbi1jb250ZW50XCI+XG4gICAgICAgICAgU2tpcCB0byBjb250ZW50XG4gICAgICAgIDwvcGYtdjYtc2tpcC10by1jb250ZW50PlxuXG4gICAgICAgIDxwZi12Ni1tYXN0aGVhZCBzbG90PVwibWFzdGhlYWRcIj5cbiAgICAgICAgICA8YSBjbGFzcz1cIm1hc3RoZWFkLWxvZ29cIlxuICAgICAgICAgICAgIGhyZWY9XCIvXCJcbiAgICAgICAgICAgICBzbG90PVwibG9nb1wiPlxuICAgICAgICAgICAgPGltZyBhbHQ9XCJDRU0gRGV2IFNlcnZlclwiXG4gICAgICAgICAgICAgICAgIHNyYz1cIi9fX2NlbS9sb2dvLnN2Z1wiPlxuICAgICAgICAgICAgJHt0aGlzLnBhY2thZ2VOYW1lID8gaHRtbGA8aDE+JHt0aGlzLnBhY2thZ2VOYW1lfTwvaDE+YCA6IG5vdGhpbmd9XG4gICAgICAgICAgPC9hPlxuICAgICAgICAgIDxwZi12Ni10b29sYmFyIHNsb3Q9XCJ0b29sYmFyXCI+XG4gICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cCB2YXJpYW50PVwiYWN0aW9uLWdyb3VwXCI+XG4gICAgICAgICAgICAgICR7dGhpcy4jcmVuZGVyU291cmNlQnV0dG9uKCl9XG4gICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImRlYnVnLWluZm9cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cInBsYWluXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGFyaWEtbGFiZWw9XCJEZWJ1ZyBpbmZvXCI+XG4gICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICBoZWlnaHQ9XCIxNlwiXG4gICAgICAgICAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIlxuICAgICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTggMTVBNyA3IDAgMSAxIDggMWE3IDcgMCAwIDEgMCAxNHptMCAxQTggOCAwIDEgMCA4IDBhOCA4IDAgMCAwIDAgMTZ6XCIvPlxuICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwibTguOTMgNi41ODgtMi4yOS4yODctLjA4Mi4zOC40NS4wODNjLjI5NC4wNy4zNTIuMTc2LjI4OC40NjlsLS43MzggMy40NjhjLS4xOTQuODk3LjEwNSAxLjMxOS44MDggMS4zMTkuNTQ1IDAgMS4xNzgtLjI1MiAxLjQ2NS0uNTk4bC4wODgtLjQxNmMtLjIuMTc2LS40OTIuMjQ2LS42ODYuMjQ2LS4yNzUgMC0uMzc1LS4xOTMtLjMwNC0uNTMzTDguOTMgNi41ODh6TTkgNC41YTEgMSAwIDEgMS0yIDAgMSAxIDAgMCAxIDIgMHpcIi8+XG4gICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgPHBmLXY2LXRvZ2dsZS1ncm91cCBjbGFzcz1cImNvbG9yLXNjaGVtZS10b2dnbGVcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkNvbG9yIHNjaGVtZVwiPlxuICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvZ2dsZS1ncm91cC1pdGVtIHZhbHVlPVwibGlnaHRcIj5cbiAgICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiIGFyaWEtbGFiZWw9XCJMaWdodCBtb2RlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk04IDExYTMgMyAwIDEgMSAwLTYgMyAzIDAgMCAxIDAgNnptMCAxYTQgNCAwIDEgMCAwLTggNCA0IDAgMCAwIDAgOHpNOCAwYS41LjUgMCAwIDEgLjUuNXYyYS41LjUgMCAwIDEtMSAwdi0yQS41LjUgMCAwIDEgOCAwem0wIDEzYS41LjUgMCAwIDEgLjUuNXYyYS41LjUgMCAwIDEtMSAwdi0yQS41LjUgMCAwIDEgOCAxM3ptOC01YS41LjUgMCAwIDEtLjUuNWgtMmEuNS41IDAgMCAxIDAtMWgyYS41LjUgMCAwIDEgLjUuNXpNMyA4YS41LjUgMCAwIDEtLjUuNWgtMmEuNS41IDAgMCAxIDAtMWgyQS41LjUgMCAwIDEgMyA4em0xMC42NTctNS42NTdhLjUuNSAwIDAgMSAwIC43MDdsLTEuNDE0IDEuNDE1YS41LjUgMCAxIDEtLjcwNy0uNzA4bDEuNDE0LTEuNDE0YS41LjUgMCAwIDEgLjcwNyAwem0tOS4xOTMgOS4xOTNhLjUuNSAwIDAgMSAwIC43MDdMMy4wNSAxMy42NTdhLjUuNSAwIDAgMS0uNzA3LS43MDdsMS40MTQtMS40MTRhLjUuNSAwIDAgMSAuNzA3IDB6bTkuMTkzIDIuMTIxYS41LjUgMCAwIDEtLjcwNyAwbC0xLjQxNC0xLjQxNGEuNS41IDAgMCAxIC43MDctLjcwN2wxLjQxNCAxLjQxNGEuNS41IDAgMCAxIDAgLjcwN3pNNC40NjQgNC40NjVhLjUuNSAwIDAgMS0uNzA3IDBMMi4zNDMgMy4wNWEuNS41IDAgMSAxIC43MDctLjcwN2wxLjQxNCAxLjQxNGEuNS41IDAgMCAxIDAgLjcwOHpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICAgIDxwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cImRhcmtcIj5cbiAgICAgICAgICAgICAgICAgICAgPHN2ZyB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiIGFyaWEtbGFiZWw9XCJEYXJrIG1vZGVcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTYgLjI3OGEuNzY4Ljc2OCAwIDAgMSAuMDguODU4IDcuMjA4IDcuMjA4IDAgMCAwLS44NzggMy40NmMwIDQuMDIxIDMuMjc4IDcuMjc3IDcuMzE4IDcuMjc3LjUyNyAwIDEuMDQtLjA1NSAxLjUzMy0uMTZhLjc4Ny43ODcgMCAwIDEgLjgxLjMxNi43MzMuNzMzIDAgMCAxLS4wMzEuODkzQTguMzQ5IDguMzQ5IDAgMCAxIDguMzQ0IDE2QzMuNzM0IDE2IDAgMTIuMjg2IDAgNy43MSAwIDQuMjY2IDIuMTE0IDEuMzEyIDUuMTI0LjA2QS43NTIuNzUyIDAgMCAxIDYgLjI3OHpNNC44NTggMS4zMTFBNy4yNjkgNy4yNjkgMCAwIDAgMS4wMjUgNy43MWMwIDQuMDIgMy4yNzkgNy4yNzYgNy4zMTkgNy4yNzZhNy4zMTYgNy4zMTYgMCAwIDAgNS4yMDUtMi4xNjJjLS4zMzcuMDQyLS42OC4wNjMtMS4wMjkuMDYzLTQuNjEgMC04LjM0My0zLjcxNC04LjM0My04LjI5IDAtMS4xNjcuMjQyLTIuMjc4LjY4MS0zLjI4NnpcIi8+XG4gICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b2dnbGUtZ3JvdXAtaXRlbT5cbiAgICAgICAgICAgICAgICAgIDxwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbSB2YWx1ZT1cInN5c3RlbVwiPlxuICAgICAgICAgICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCIgYXJpYS1sYWJlbD1cIlN5c3RlbSBwcmVmZXJlbmNlXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0wIDEuNUExLjUgMS41IDAgMCAxIDEuNSAwaDEzQTEuNSAxLjUgMCAwIDEgMTYgMS41djhhMS41IDEuNSAwIDAgMS0xLjUgMS41aC0xM0ExLjUgMS41IDAgMCAxIDAgOS41di04ek0xLjUgMWEuNS41IDAgMCAwLS41LjV2OGEuNS41IDAgMCAwIC41LjVoMTNhLjUuNSAwIDAgMCAuNS0uNXYtOGEuNS41IDAgMCAwLS41LS41aC0xM3pcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0yLjUgMTJoMTFhLjUuNSAwIDAgMSAwIDFoLTExYS41LjUgMCAwIDEgMC0xem0wIDJoMTFhLjUuNSAwIDAgMSAwIDFoLTExYS41LjUgMCAwIDEgMC0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICA8L3BmLXY2LXRvZ2dsZS1ncm91cC1pdGVtPlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9nZ2xlLWdyb3VwPlxuICAgICAgICAgICAgICA8L3BmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICA8L3BmLXY2LXRvb2xiYXI+XG4gICAgICAgIDwvcGYtdjYtbWFzdGhlYWQ+XG5cbiAgICAgICAgPHBmLXY2LXBhZ2Utc2lkZWJhciBzbG90PVwic2lkZWJhclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgP2V4cGFuZGVkPSR7dGhpcy5zaWRlYmFyID09PSAnZXhwYW5kZWQnfVxuICAgICAgICAgICAgICAgICAgICAgICAgICAgID9jb2xsYXBzZWQ9JHt0aGlzLnNpZGViYXIgIT09ICdleHBhbmRlZCd9PlxuICAgICAgICAgIDxzbG90IG5hbWU9XCJuYXZpZ2F0aW9uXCI+PC9zbG90PlxuICAgICAgICA8L3BmLXY2LXBhZ2Utc2lkZWJhcj5cblxuICAgICAgICA8cGYtdjYtcGFnZS1tYWluIHNsb3Q9XCJtYWluXCIgaWQ9XCJtYWluLWNvbnRlbnRcIj5cbiAgICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICAgICAgPGZvb3RlciBjbGFzcz1cInBmLW0tc3RpY2t5LWJvdHRvbVwiPlxuICAgICAgICAgICAgPGRpdiBjbGFzcz1cImZvb3Rlci1kZXNjcmlwdGlvbiR7dGhpcy5oYXNEZXNjcmlwdGlvbiA/ICcnIDogJyBlbXB0eSd9XCI+XG4gICAgICAgICAgICAgIDxzbG90IG5hbWU9XCJkZXNjcmlwdGlvblwiPjwvc2xvdD5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgPGNlbS1kcmF3ZXIgP29wZW49JHt0aGlzLmRyYXdlciA9PT0gJ2V4cGFuZGVkJ31cbiAgICAgICAgICAgICAgICAgICAgICAgIGRyYXdlci1oZWlnaHQ9XCIke3RoaXMuZHJhd2VySGVpZ2h0IHx8ICc0MDAnfVwiPlxuICAgICAgICAgICAgICA8cGYtdjYtdGFicyBzZWxlY3RlZD1cIiR7dGhpcy50YWJzU2VsZWN0ZWQgfHwgJzAnfVwiPlxuICAgICAgICAgICAgICAgIDxwZi12Ni10YWIgdGl0bGU9XCJLbm9ic1wiPlxuICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImtub2JzLWNvbnRhaW5lclwiPlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwia25vYnNcIj5cbiAgICAgICAgICAgICAgICAgICAgICA8cCBjbGFzcz1cImtub2JzLWVtcHR5XCI+Tm8ga25vYnMgYXZhaWxhYmxlIGZvciB0aGlzIGVsZW1lbnQuPC9wPlxuICAgICAgICAgICAgICAgICAgICA8L3Nsb3Q+XG4gICAgICAgICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRhYj5cbiAgICAgICAgICAgICAgICA8cGYtdjYtdGFiIHRpdGxlPVwiTWFuaWZlc3QgQnJvd3NlclwiPlxuICAgICAgICAgICAgICAgICAgPGNlbS1tYW5pZmVzdC1icm93c2VyPlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtdHJlZVwiIHNsb3Q9XCJtYW5pZmVzdC10cmVlXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtbmFtZVwiIHNsb3Q9XCJtYW5pZmVzdC1uYW1lXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgICA8c2xvdCBuYW1lPVwibWFuaWZlc3QtZGV0YWlsc1wiIHNsb3Q9XCJtYW5pZmVzdC1kZXRhaWxzXCI+PC9zbG90PlxuICAgICAgICAgICAgICAgICAgPC9jZW0tbWFuaWZlc3QtYnJvd3Nlcj5cbiAgICAgICAgICAgICAgICA8L3BmLXY2LXRhYj5cbiAgICAgICAgICAgICAgICA8cGYtdjYtdGFiIHRpdGxlPVwiU2VydmVyIExvZ3NcIj5cbiAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJsb2dzLXdyYXBwZXJcIj5cbiAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXIgc3RpY2t5PlxuICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRleHQtaW5wdXQtZ3JvdXAgaWQ9XCJsb2dzLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHBsYWNlaG9sZGVyPVwiRmlsdGVyIGxvZ3MuLi5cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBpY29uPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgcm9sZT1cInByZXNlbnRhdGlvblwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGhlaWdodD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB3aWR0aD1cIjFlbVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2aWV3Qm94PVwiMCAwIDUxMiA1MTJcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGQ9XCJNNTA1IDQ0Mi43TDQwNS4zIDM0M2MtNC41LTQuNS0xMC42LTctMTctN0gzNzJjMjcuNi0zNS4zIDQ0LTc5LjcgNDQtMTI4QzQxNiA5My4xIDMyMi45IDAgMjA4IDBTMCA5My4xIDAgMjA4czkzLjEgMjA4IDIwOCAyMDhjNDguMyAwIDkyLjctMTYuNCAxMjgtNDR2MTYuM2MwIDYuNCAyLjUgMTIuNSA3IDE3bDk5LjcgOTkuN2M5LjQgOS40IDI0LjYgOS40IDMzLjkgMGwyOC4zLTI4LjNjOS40LTkuNCA5LjQtMjQuNi4xLTM0ek0yMDggMzM2Yy03MC43IDAtMTI4LTU3LjItMTI4LTEyOCAwLTcwLjcgNTcuMi0xMjggMTI4LTEyOCA3MC43IDAgMTI4IDU3LjIgMTI4IDEyOCAwIDcwLjctNTcuMiAxMjgtMTI4IDEyOHpcIj48L3BhdGg+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdGV4dC1pbnB1dC1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwibG9nLWxldmVsLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBsYWJlbD1cIkZpbHRlciBsb2cgbGV2ZWxzXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHNwYW4gc2xvdD1cInRvZ2dsZS10ZXh0XCI+TG9nIExldmVsczwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtbWVudS1pdGVtIHZhcmlhbnQ9XCJjaGVja2JveFwiIHZhbHVlPVwiaW5mb1wiIGNoZWNrZWQ+SW5mbzwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJ3YXJuXCIgY2hlY2tlZD5XYXJuaW5nczwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1tZW51LWl0ZW0gdmFyaWFudD1cImNoZWNrYm94XCIgdmFsdWU9XCJlcnJvclwiIGNoZWNrZWQ+RXJyb3JzPC9wZi12Ni1tZW51LWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LW1lbnUtaXRlbSB2YXJpYW50PVwiY2hlY2tib3hcIiB2YWx1ZT1cImRlYnVnXCIgY2hlY2tlZD5EZWJ1ZzwvcGYtdjYtbWVudS1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWRyb3Bkb3duPlxuICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyLWdyb3VwPlxuICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWdyb3VwIHZhcmlhbnQ9XCJhY3Rpb24tZ3JvdXBcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJjb3B5LWxvZ3NcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJ0ZXJ0aWFyeVwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgc2l6ZT1cInNtYWxsXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHN2ZyBzbG90PVwiaWNvblwiIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTEzIDBINmEyIDIgMCAwIDAtMiAyIDIgMiAwIDAgMC0yIDJ2MTBhMiAyIDAgMCAwIDIgMmg3YTIgMiAwIDAgMCAyLTIgMiAyIDAgMCAwIDItMlYyYTIgMiAwIDAgMC0yLTJ6bTAgMTNWNGEyIDIgMCAwIDAtMi0ySDVhMSAxIDAgMCAxIDEtMWg3YTEgMSAwIDAgMSAxIDF2MTBhMSAxIDAgMCAxLTEgMXpNMyAxM1Y0YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjlhMSAxIDAgMCAxLTEgMUg0YTEgMSAwIDAgMS0xLTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIENvcHkgTG9nc1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyPlxuICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwibG9nLWNvbnRhaW5lclwiPjwvZGl2PlxuICAgICAgICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICAgICAgPC9wZi12Ni10YWI+XG4gICAgICAgICAgICAgICAgPHBmLXY2LXRhYiB0aXRsZT1cIkV2ZW50c1wiPlxuICAgICAgICAgICAgICAgICAgPGRpdiBjbGFzcz1cImV2ZW50cy13cmFwcGVyXCI+XG4gICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyIHN0aWNreT5cbiAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDxwZi12Ni10ZXh0LWlucHV0LWdyb3VwIGlkPVwiZXZlbnRzLWZpbHRlclwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHBsYWNlaG9sZGVyPVwiRmlsdGVyIGV2ZW50cy4uLlwiXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIGljb24+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPHN2ZyBzbG90PVwiaWNvblwiIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8cGF0aCBkPVwiTTExLjc0MiAxMC4zNDRhNi41IDYuNSAwIDEgMC0xLjM5NyAxLjM5OGgtLjAwMWMuMDMuMDQuMDYyLjA3OC4wOTguMTE1bDMuODUgMy44NWExIDEgMCAwIDAgMS40MTUtMS40MTRsLTMuODUtMy44NWExLjAwNyAxLjAwNyAwIDAgMC0uMTE1LS4xek0xMiA2LjVhNS41IDUuNSAwIDEgMS0xMSAwIDUuNSA1LjUgMCAwIDEgMTEgMHpcIi8+XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdGV4dC1pbnB1dC1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwiZXZlbnQtdHlwZS1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgZXZlbnQgdHlwZXNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5FdmVudCBUeXBlczwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1kcm9wZG93bj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWRyb3Bkb3duIGlkPVwiZWxlbWVudC1maWx0ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgbGFiZWw9XCJGaWx0ZXIgZWxlbWVudHNcIj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8c3BhbiBzbG90PVwidG9nZ2xlLXRleHRcIj5FbGVtZW50czwvc3Bhbj5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni1kcm9wZG93bj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1ncm91cCB2YXJpYW50PVwiYWN0aW9uLWdyb3VwXCI+XG4gICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgICA8cGYtdjYtYnV0dG9uIGlkPVwiY2xlYXItZXZlbnRzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk01LjUgNS41QS41LjUgMCAwIDEgNiA2djZhLjUuNSAwIDAgMS0xIDBWNmEuNS41IDAgMCAxIC41LS41em0yLjUgMGEuNS41IDAgMCAxIC41LjV2NmEuNS41IDAgMCAxLTEgMFY2YS41LjUgMCAwIDEgLjUtLjV6bTMgLjVhLjUuNSAwIDAgMC0xIDB2NmEuNS41IDAgMCAwIDEgMFY2elwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxwYXRoIGZpbGwtcnVsZT1cImV2ZW5vZGRcIiBkPVwiTTE0LjUgM2ExIDEgMCAwIDEtMSAxSDEzdjlhMiAyIDAgMCAxLTIgMkg1YTIgMiAwIDAgMS0yLTJWNGgtLjVhMSAxIDAgMCAxLTEtMVYyYTEgMSAwIDAgMSAxLTFINmExIDEgMCAwIDEgMS0xaDJhMSAxIDAgMCAxIDEgMWgzLjVhMSAxIDAgMCAxIDEgMXYxek00LjExOCA0IDQgNC4wNTlWMTNhMSAxIDAgMCAwIDEgMWg2YTEgMSAwIDAgMCAxLTFWNC4wNTlMMTEuODgyIDRINC4xMTh6TTIuNSAzVjJoMTF2MWgtMTF6XCIvPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvc3ZnPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIENsZWFyIEV2ZW50c1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LXRvb2xiYXItaXRlbT5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgPHBmLXY2LWJ1dHRvbiBpZD1cImNvcHktZXZlbnRzXCJcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwidGVydGlhcnlcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJzbWFsbFwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxzdmcgc2xvdD1cImljb25cIiB3aWR0aD1cIjE2XCIgaGVpZ2h0PVwiMTZcIiB2aWV3Qm94PVwiMCAwIDE2IDE2XCIgZmlsbD1cImN1cnJlbnRDb2xvclwiPlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPHBhdGggZD1cIk0xMyAwSDZhMiAyIDAgMCAwLTIgMiAyIDIgMCAwIDAtMiAydjEwYTIgMiAwIDAgMCAyIDJoN2EyIDIgMCAwIDAgMi0yIDIgMiAwIDAgMCAyLTJWMmEyIDIgMCAwIDAtMi0yem0wIDEzVjRhMiAyIDAgMCAwLTItMkg1YTEgMSAwIDAgMSAxLTFoN2ExIDEgMCAwIDEgMSAxdjEwYTEgMSAwIDAgMS0xIDF6TTMgMTNWNGExIDEgMCAwIDEgMS0xaDdhMSAxIDAgMCAxIDEgMXY5YTEgMSAwIDAgMS0xIDFINGExIDEgMCAwIDEtMS0xelwiLz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L3N2Zz5cbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBDb3B5IEV2ZW50c1xuICAgICAgICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgICAgICAgICAgICAgICAgICAgIDwvcGYtdjYtdG9vbGJhci1ncm91cD5cbiAgICAgICAgICAgICAgICAgICAgPC9wZi12Ni10b29sYmFyPlxuICAgICAgICAgICAgICAgICAgICA8cGYtdjYtZHJhd2VyIGlkPVwiZXZlbnQtZHJhd2VyXCIgZXhwYW5kZWQ+XG4gICAgICAgICAgICAgICAgICAgICAgPGRpdiBpZD1cImV2ZW50LWxpc3RcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwiZXZlbnQtZGV0YWlsLWhlYWRlclwiIHNsb3Q9XCJwYW5lbC1oZWFkZXJcIj48L2Rpdj5cbiAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGlkPVwiZXZlbnQtZGV0YWlsLWJvZHlcIiBzbG90PVwicGFuZWwtYm9keVwiPjwvZGl2PlxuICAgICAgICAgICAgICAgICAgICA8L3BmLXY2LWRyYXdlcj5cbiAgICAgICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtdGFiPlxuICAgICAgICAgICAgICAgIDxwZi12Ni10YWIgdGl0bGU9XCJIZWFsdGhcIj5cbiAgICAgICAgICAgICAgICAgIDxjZW0taGVhbHRoLXBhbmVsICR7dGhpcy5wcmltYXJ5VGFnTmFtZSA/IGh0bWxgY29tcG9uZW50PVwiJHt0aGlzLnByaW1hcnlUYWdOYW1lfVwiYCA6IG5vdGhpbmd9PlxuICAgICAgICAgICAgICAgICAgPC9jZW0taGVhbHRoLXBhbmVsPlxuICAgICAgICAgICAgICAgIDwvcGYtdjYtdGFiPlxuICAgICAgICAgICAgICA8L3BmLXY2LXRhYnM+XG4gICAgICAgICAgICA8L2NlbS1kcmF3ZXI+XG4gICAgICAgICAgPC9mb290ZXI+XG4gICAgICAgIDwvcGYtdjYtcGFnZS1tYWluPlxuICAgICAgPC9wZi12Ni1wYWdlPlxuXG4gICAgICA8cGYtdjYtbW9kYWwgaWQ9XCJkZWJ1Zy1tb2RhbFwiIHZhcmlhbnQ9XCJsYXJnZVwiPlxuICAgICAgICA8aDIgc2xvdD1cImhlYWRlclwiPkRlYnVnIEluZm9ybWF0aW9uPC9oMj5cbiAgICAgICAgPGRsIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0IHBmLW0taG9yaXpvbnRhbCBwZi1tLWNvbXBhY3RcIj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+U2VydmVyIFZlcnNpb248L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctdmVyc2lvblwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5TZXJ2ZXIgT1M8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctb3NcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+V2F0Y2ggRGlyZWN0b3J5PC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLXdhdGNoLWRpclwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5NYW5pZmVzdCBTaXplPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLW1hbmlmZXN0LXNpemVcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+RGVtb3MgRm91bmQ8L2R0PlxuICAgICAgICAgICAgPGRkIGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiIGlkPVwiZGVidWctZGVtby1jb3VudFwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5UYWcgTmFtZTwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCI+JHt0aGlzLnByaW1hcnlUYWdOYW1lIHx8ICctJ308L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5EZW1vIFRpdGxlPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIj4ke3RoaXMuZGVtb1RpdGxlIHx8ICctJ308L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj5Ccm93c2VyPC9kdD5cbiAgICAgICAgICAgIDxkZCBjbGFzcz1cInBmLXY2LWMtZGVzY3JpcHRpb24tbGlzdF9fZGVzY3JpcHRpb25cIiBpZD1cImRlYnVnLWJyb3dzZXJcIj4tPC9kZD5cbiAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICAgICAgPGR0IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X190ZXJtXCI+VXNlciBBZ2VudDwvZHQ+XG4gICAgICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCIgaWQ9XCJkZWJ1Zy11YVwiPi08L2RkPlxuICAgICAgICAgIDwvZGl2PlxuICAgICAgICA8L2RsPlxuICAgICAgICA8ZGl2IGlkPVwiZGVtby11cmxzLWNvbnRhaW5lclwiPjwvZGl2PlxuICAgICAgICA8cGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uIGlkPVwiZGVidWctaW1wb3J0bWFwLWRldGFpbHNcIlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIHRvZ2dsZS10ZXh0PVwiU2hvdyBJbXBvcnQgTWFwXCI+XG4gICAgICAgICAgPHByZSBpZD1cImRlYnVnLWltcG9ydG1hcFwiPi08L3ByZT5cbiAgICAgICAgPC9wZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24+XG4gICAgICAgIDxkaXYgc2xvdD1cImZvb3RlclwiIGNsYXNzPVwiYnV0dG9uLWNvbnRhaW5lclwiPlxuICAgICAgICAgIDxwZi12Ni1idXR0b24gY2xhc3M9XCJkZWJ1Zy1jb3B5XCIgdmFyaWFudD1cInByaW1hcnlcIj5cbiAgICAgICAgICAgIENvcHkgRGVidWcgSW5mb1xuICAgICAgICAgIDwvcGYtdjYtYnV0dG9uPlxuICAgICAgICAgIDxwZi12Ni1idXR0b24gY2xhc3M9XCJkZWJ1Zy1jbG9zZVwiIHZhcmlhbnQ9XCJzZWNvbmRhcnlcIj5cbiAgICAgICAgICAgIENsb3NlXG4gICAgICAgICAgPC9wZi12Ni1idXR0b24+XG4gICAgICAgIDwvZGl2PlxuICAgICAgPC9wZi12Ni1tb2RhbD5cblxuICAgICAgPCEtLSBSZWNvbm5lY3Rpb24gbW9kYWwgLS0+XG4gICAgICA8cGYtdjYtbW9kYWwgaWQ9XCJyZWNvbm5lY3Rpb24tbW9kYWxcIiB2YXJpYW50PVwibGFyZ2VcIj5cbiAgICAgICAgPGgyIHNsb3Q9XCJoZWFkZXJcIj5EZXZlbG9wbWVudCBTZXJ2ZXIgRGlzY29ubmVjdGVkPC9oMj5cbiAgICAgICAgPGNlbS1yZWNvbm5lY3Rpb24tY29udGVudCBpZD1cInJlY29ubmVjdGlvbi1jb250ZW50XCI+PC9jZW0tcmVjb25uZWN0aW9uLWNvbnRlbnQ+XG4gICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJyZWxvYWQtYnV0dG9uXCJcbiAgICAgICAgICAgICAgICAgICAgICBzbG90PVwiZm9vdGVyXCJcbiAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwicHJpbWFyeVwiPlJlbG9hZCBQYWdlPC9wZi12Ni1idXR0b24+XG4gICAgICAgIDxwZi12Ni1idXR0b24gaWQ9XCJyZXRyeS1idXR0b25cIlxuICAgICAgICAgICAgICAgICAgICAgIHNsb3Q9XCJmb290ZXJcIlxuICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJzZWNvbmRhcnlcIj5SZXRyeSBOb3c8L3BmLXY2LWJ1dHRvbj5cbiAgICAgIDwvcGYtdjYtbW9kYWw+XG5cbiAgICAgIDwhLS0gVHJhbnNmb3JtIGVycm9yIG92ZXJsYXkgLS0+XG4gICAgICA8Y2VtLXRyYW5zZm9ybS1lcnJvci1vdmVybGF5IGlkPVwiZXJyb3Itb3ZlcmxheVwiPlxuICAgICAgPC9jZW0tdHJhbnNmb3JtLWVycm9yLW92ZXJsYXk+XG4gICAgYDtcbiAgfVxuXG4gICNyZW5kZXJTb3VyY2VCdXR0b24oKSB7XG4gICAgaWYgKCF0aGlzLnNvdXJjZVVSTCkgcmV0dXJuIG5vdGhpbmc7XG5cbiAgICBsZXQgbGFiZWwgPSAnVmVyc2lvbiBDb250cm9sJztcbiAgICBsZXQgcGF0aCA9ICdNNS44NTQgNC44NTRhLjUuNSAwIDEgMC0uNzA4LS43MDhsLTMuNSAzLjVhLjUuNSAwIDAgMCAwIC43MDhsMy41IDMuNWEuNS41IDAgMCAwIC43MDgtLjcwOEwyLjcwNyA4bDMuMTQ3LTMuMTQ2em00LjI5MiAwYS41LjUgMCAwIDEgLjcwOC0uNzA4bDMuNSAzLjVhLjUuNSAwIDAgMSAwIC43MDhsLTMuNSAzLjVhLjUuNSAwIDAgMS0uNzA4LS43MDhMMTMuMjkzIDhsLTMuMTQ3LTMuMTQ2eic7XG5cbiAgICBpZiAodGhpcy5zb3VyY2VVUkwuaW5jbHVkZXMoJ2dpdGh1Yi5jb20nKSkge1xuICAgICAgbGFiZWwgPSAnR2l0SHViLmNvbSc7XG4gICAgICBwYXRoID0gJ004IDBDMy41OCAwIDAgMy41OCAwIDhjMCAzLjU0IDIuMjkgNi41MyA1LjQ3IDcuNTkuNC4wNy41NS0uMTcuNTUtLjM4IDAtLjE5LS4wMS0uODItLjAxLTEuNDktMi4wMS4zNy0yLjUzLS40OS0yLjY5LS45NC0uMDktLjIzLS40OC0uOTQtLjgyLTEuMTMtLjI4LS4xNS0uNjgtLjUyLS4wMS0uNTMuNjMtLjAxIDEuMDguNTggMS4yMy44Mi43MiAxLjIxIDEuODcuODcgMi4zMy42Ni4wNy0uNTIuMjgtLjg3LjUxLTEuMDctMS43OC0uMi0zLjY0LS44OS0zLjY0LTMuOTUgMC0uODcuMzEtMS41OS44Mi0yLjE1LS4wOC0uMi0uMzYtMS4wMi4wOC0yLjEyIDAgMCAuNjctLjIxIDIuMi44Mi42NC0uMTggMS4zMi0uMjcgMi0uMjcuNjggMCAxLjM2LjA5IDIgLjI3IDEuNTMtMS4wNCAyLjItLjgyIDIuMi0uODIuNDQgMS4xLjE2IDEuOTIuMDggMi4xMi41MS41Ni44MiAxLjI3LjgyIDIuMTUgMCAzLjA3LTEuODcgMy43NS0zLjY1IDMuOTUuMjkuMjUuNTQuNzMuNTQgMS40OCAwIDEuMDctLjAxIDEuOTMtLjAxIDIuMiAwIC4yMS4xNS40Ni41NS4zOEE4LjAxMyA4LjAxMyAwIDAwMTYgOGMwLTQuNDItMy41OC04LTgtOHonO1xuICAgIH0gZWxzZSBpZiAodGhpcy5zb3VyY2VVUkwuaW5jbHVkZXMoJ2dpdGxhYi5jb20nKSkge1xuICAgICAgbGFiZWwgPSAnR2l0TGFiJztcbiAgICAgIHBhdGggPSAnbTE1LjczNCA2LjEtLjAyMi0uMDU4TDEzLjUzNC4zNThhLjU2OC41NjggMCAwIDAtLjU2My0uMzU2LjU4My41ODMgMCAwIDAtLjMyOC4xMjIuNTgyLjU4MiAwIDAgMC0uMTkzLjI5NGwtMS40NyA0LjQ5OUg1LjAyNWwtMS40Ny00LjVBLjU3Mi41NzIgMCAwIDAgMy4zNjAuMTc0YS41NzIuNTcyIDAgMCAwLS4zMjgtLjE3Mi41ODIuNTgyIDAgMCAwLS41NjMuMzU3TC4yOSA2LjA0bC0uMDIyLjA1N0E0LjA0NCA0LjA0NCAwIDAgMCAxLjYxIDEwLjc3bC4wMDcuMDA2LjAyLjAxNCAzLjMxOCAyLjQ4NSAxLjY0IDEuMjQyIDEgLjc1NWEuNjczLjY3MyAwIDAgMCAuODE0IDBsMS0uNzU1IDEuNjQtMS4yNDIgMy4zMzgtMi41LjAwOS0uMDA3YTQuMDUgNC4wNSAwIDAgMCAxLjM0LTQuNjY4Wic7XG4gICAgfSBlbHNlIGlmICh0aGlzLnNvdXJjZVVSTC5pbmNsdWRlcygnYml0YnVja2V0Lm9yZycpKSB7XG4gICAgICBsYWJlbCA9ICdCaXRidWNrZXQnO1xuICAgICAgcGF0aCA9ICdNMCAxLjVBMS41IDEuNSAwIDAgMSAxLjUgMGgxM0ExLjUgMS41IDAgMCAxIDE2IDEuNXYxM2ExLjUgMS41IDAgMCAxLTEuNSAxLjVoLTEzQTEuNSAxLjUgMCAwIDEgMCAxNC41di0xM3pNMi41IDNhLjUuNSAwIDAgMC0uNS41djlhLjUuNSAwIDAgMCAuNS41aDExYS41LjUgMCAwIDAgLjUtLjV2LTlhLjUuNSAwIDAgMC0uNS0uNWgtMTF6bTUuMDM4IDEuNDM1YS41LjUgMCAwIDEgLjkyNCAwbDEuNDIgMy4zN0g4Ljc4bC0uMjQzLjYwOC0uMjQzLS42MDhINS4wODJsMS40Mi0zLjM3ek04IDkuMTQzbC0uNzQzIDEuODU3SDQuNzQzTDYuMDc2IDcuNjA4IDggOS4xNDN6JztcbiAgICB9XG5cbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxwZi12Ni10b29sYmFyLWl0ZW0+XG4gICAgICAgIDxwZi12Ni1idXR0b24gaHJlZj1cIiR7dGhpcy5zb3VyY2VVUkx9XCJcbiAgICAgICAgICAgICAgICAgICAgICB0YXJnZXQ9XCJfYmxhbmtcIlxuICAgICAgICAgICAgICAgICAgICAgIHZhcmlhbnQ9XCJwbGFpblwiXG4gICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIlZpZXcgc291cmNlIGZpbGVcIj5cbiAgICAgICAgICA8c3ZnIGFyaWEtbGFiZWw9XCIke2xhYmVsfVwiXG4gICAgICAgICAgICAgICB3aWR0aD1cIjFyZW1cIlxuICAgICAgICAgICAgICAgaGVpZ2h0PVwiMXJlbVwiXG4gICAgICAgICAgICAgICBmaWxsPVwiY3VycmVudENvbG9yXCJcbiAgICAgICAgICAgICAgIHZpZXdCb3g9XCIwIDAgMTYgMTZcIj5cbiAgICAgICAgICAgIDxwYXRoIGQ9XCIke3BhdGh9XCIvPlxuICAgICAgICAgIDwvc3ZnPlxuICAgICAgICA8L3BmLXY2LWJ1dHRvbj5cbiAgICAgIDwvcGYtdjYtdG9vbGJhci1pdGVtPlxuICAgIGA7XG4gIH1cblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuXG4gICAgLy8gQ2hlY2sgaWYgd2UgbmVlZCB0byBtaWdyYXRlIGZyb20gbG9jYWxTdG9yYWdlXG4gICAgdGhpcy4jbWlncmF0ZUZyb21Mb2NhbFN0b3JhZ2VJZk5lZWRlZCgpO1xuICB9XG5cbiAgZmlyc3RVcGRhdGVkKCkge1xuICAgIC8vIFNldCB1cCBkZWJ1ZyBvdmVybGF5XG4gICAgdGhpcy4jc2V0dXBEZWJ1Z092ZXJsYXkoKTtcblxuICAgIC8vIFNldCB1cCBjb2xvciBzY2hlbWUgdG9nZ2xlXG4gICAgdGhpcy4jc2V0dXBDb2xvclNjaGVtZVRvZ2dsZSgpO1xuXG4gICAgLy8gU2V0IHVwIGZvb3RlciBkcmF3ZXIgYW5kIHRhYnNcbiAgICB0aGlzLiNzZXR1cEZvb3RlckRyYXdlcigpO1xuXG4gICAgLy8gTGlzdGVuIGZvciBzZXJ2ZXIgbG9nIG1lc3NhZ2VzIGZyb20gV2ViU29ja2V0XG4gICAgdGhpcy4jc2V0dXBMb2dMaXN0ZW5lcigpO1xuXG4gICAgLy8gU2V0IHVwIGtub2IgZXZlbnQgY29vcmRpbmF0aW9uXG4gICAgdGhpcy4jc2V0dXBLbm9iQ29vcmRpbmF0aW9uKCk7XG5cbiAgICAvLyBTZXQgdXAgdHJlZSBzdGF0ZSBwZXJzaXN0ZW5jZVxuICAgIHRoaXMuI3NldHVwVHJlZVN0YXRlUGVyc2lzdGVuY2UoKTtcblxuICAgIC8vIFNldCB1cCBzaWRlYmFyIHN0YXRlIHBlcnNpc3RlbmNlXG4gICAgdGhpcy4jc2V0dXBTaWRlYmFyU3RhdGVQZXJzaXN0ZW5jZSgpO1xuXG4gICAgLy8gU2V0IHVwIGVsZW1lbnQgZXZlbnQgY2FwdHVyZVxuICAgIHRoaXMuI3NldHVwRXZlbnRDYXB0dXJlKCkudGhlbigoKSA9PiB7XG4gICAgICB0aGlzLiNzZXR1cEV2ZW50TGlzdGVuZXJzKCk7XG4gICAgfSk7XG5cbiAgICAvLyBTZXQgdXAgcmVjb25uZWN0aW9uIG1vZGFsIGJ1dHRvbiBoYW5kbGVyc1xuICAgIC8qIGM4IGlnbm9yZSBzdGFydCAtIHdpbmRvdy5sb2NhdGlvbi5yZWxvYWQgd291bGQgcmVsb2FkIHRlc3QgcGFnZSAqL1xuICAgIHRoaXMuIyQoJ3JlbG9hZC1idXR0b24nKT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICB3aW5kb3cubG9jYXRpb24ucmVsb2FkKCk7XG4gICAgfSk7XG4gICAgLyogYzggaWdub3JlIHN0b3AgKi9cblxuICAgIHRoaXMuIyQoJ3JldHJ5LWJ1dHRvbicpPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgICh0aGlzLiMkKCdyZWNvbm5lY3Rpb24tbW9kYWwnKSBhcyBhbnkpPy5jbG9zZSgpO1xuICAgICAgdGhpcy4jd3NDbGllbnQucmV0cnkoKTtcbiAgICB9KTtcblxuICAgIC8vIEluaXRpYWxpemUgV2ViU29ja2V0IGNvbm5lY3Rpb25cbiAgICB0aGlzLiN3c0NsaWVudC5pbml0KCk7XG4gIH1cblxuICBhc3luYyAjZmV0Y2hEZWJ1Z0luZm8oKSB7XG4gICAgLy8gUG9wdWxhdGUgYnJvd3NlciBpbmZvIGltbWVkaWF0ZWx5XG4gICAgY29uc3QgYnJvd3NlckVsID0gdGhpcy4jJCgnZGVidWctYnJvd3NlcicpO1xuICAgIGNvbnN0IHVhRWwgPSB0aGlzLiMkKCdkZWJ1Zy11YScpO1xuICAgIGlmIChicm93c2VyRWwpIHtcbiAgICAgIGNvbnN0IGJyb3dzZXIgPSB0aGlzLiNkZXRlY3RCcm93c2VyKCk7XG4gICAgICBicm93c2VyRWwudGV4dENvbnRlbnQgPSBicm93c2VyO1xuICAgIH1cbiAgICBpZiAodWFFbCkge1xuICAgICAgdWFFbC50ZXh0Q29udGVudCA9IG5hdmlnYXRvci51c2VyQWdlbnQ7XG4gICAgfVxuXG4gICAgLy8gRmV0Y2ggc2VydmVyIGRlYnVnIGluZm9cbiAgICBjb25zdCBkYXRhID0gYXdhaXQgZmV0Y2goJy9fX2NlbS9kZWJ1ZycpXG4gICAgICAudGhlbihyZXMgPT4gcmVzLmpzb24oKSlcbiAgICAgIC5jYXRjaCgoZXJyOiBFcnJvcikgPT4ge1xuICAgICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGZldGNoIGRlYnVnIGluZm86JywgZXJyKTtcbiAgICAgIH0pO1xuXG4gICAgaWYgKCFkYXRhKSByZXR1cm47XG4gICAgY29uc3QgdmVyc2lvbkVsID0gdGhpcy4jJCgnZGVidWctdmVyc2lvbicpO1xuICAgIGNvbnN0IG9zRWwgPSB0aGlzLiMkKCdkZWJ1Zy1vcycpO1xuICAgIGNvbnN0IHdhdGNoRGlyRWwgPSB0aGlzLiMkKCdkZWJ1Zy13YXRjaC1kaXInKTtcbiAgICBjb25zdCBtYW5pZmVzdFNpemVFbCA9IHRoaXMuIyQoJ2RlYnVnLW1hbmlmZXN0LXNpemUnKTtcbiAgICBjb25zdCBkZW1vQ291bnRFbCA9IHRoaXMuIyQoJ2RlYnVnLWRlbW8tY291bnQnKTtcbiAgICBjb25zdCBkZW1vVXJsc0NvbnRhaW5lciA9IHRoaXMuIyQoJ2RlbW8tdXJscy1jb250YWluZXInKTtcbiAgICBjb25zdCBpbXBvcnRNYXBFbCA9IHRoaXMuIyQoJ2RlYnVnLWltcG9ydG1hcCcpO1xuXG4gICAgaWYgKHZlcnNpb25FbCkgdmVyc2lvbkVsLnRleHRDb250ZW50ID0gZGF0YS52ZXJzaW9uIHx8ICctJztcbiAgICBpZiAob3NFbCkgb3NFbC50ZXh0Q29udGVudCA9IGRhdGEub3MgfHwgJy0nO1xuICAgIGlmICh3YXRjaERpckVsKSB3YXRjaERpckVsLnRleHRDb250ZW50ID0gZGF0YS53YXRjaERpciB8fCAnLSc7XG4gICAgaWYgKG1hbmlmZXN0U2l6ZUVsKSBtYW5pZmVzdFNpemVFbC50ZXh0Q29udGVudCA9IGRhdGEubWFuaWZlc3RTaXplIHx8ICctJztcbiAgICBpZiAoZGVtb0NvdW50RWwpIGRlbW9Db3VudEVsLnRleHRDb250ZW50ID0gZGF0YS5kZW1vQ291bnQgfHwgJzAnO1xuXG4gICAgaWYgKGRlbW9VcmxzQ29udGFpbmVyKSB7XG4gICAgICB0aGlzLiNwb3B1bGF0ZURlbW9VcmxzKGRlbW9VcmxzQ29udGFpbmVyLCBkYXRhLmRlbW9zKTtcbiAgICB9XG5cbiAgICBpZiAoaW1wb3J0TWFwRWwgJiYgZGF0YS5pbXBvcnRNYXApIHtcbiAgICAgIGNvbnN0IGltcG9ydE1hcEpTT04gPSBKU09OLnN0cmluZ2lmeShkYXRhLmltcG9ydE1hcCwgbnVsbCwgMik7XG4gICAgICBpbXBvcnRNYXBFbC50ZXh0Q29udGVudCA9IGltcG9ydE1hcEpTT047XG4gICAgICBkYXRhLmltcG9ydE1hcEpTT04gPSBpbXBvcnRNYXBKU09OO1xuICAgIH0gZWxzZSBpZiAoaW1wb3J0TWFwRWwpIHtcbiAgICAgIGltcG9ydE1hcEVsLnRleHRDb250ZW50ID0gJ05vIGltcG9ydCBtYXAgZ2VuZXJhdGVkJztcbiAgICB9XG5cbiAgICB0aGlzLiNkZWJ1Z0RhdGEgPSBkYXRhO1xuICB9XG5cbiAgI3BvcHVsYXRlRGVtb1VybHMoY29udGFpbmVyOiBIVE1MRWxlbWVudCwgZGVtb3M6IERlYnVnRGF0YVsnZGVtb3MnXSkge1xuICAgIGlmICghZGVtb3M/Lmxlbmd0aCkge1xuICAgICAgY29udGFpbmVyLnRleHRDb250ZW50ID0gJ05vIGRlbW9zIGZvdW5kIGluIG1hbmlmZXN0JztcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICBjb25zdCBjdXJyZW50VGFnTmFtZSA9IHRoaXMucHJpbWFyeVRhZ05hbWUgfHwgJyc7XG4gICAgY29uc3QgaXNPbkRlbW9QYWdlID0gISFjdXJyZW50VGFnTmFtZTtcblxuICAgIGlmIChpc09uRGVtb1BhZ2UpIHtcbiAgICAgIGNvbnN0IGN1cnJlbnREZW1vID0gZGVtb3MuZmluZChkZW1vID0+IGRlbW8udGFnTmFtZSA9PT0gY3VycmVudFRhZ05hbWUpO1xuICAgICAgaWYgKCFjdXJyZW50RGVtbykge1xuICAgICAgICBjb250YWluZXIudGV4dENvbnRlbnQgPSAnQ3VycmVudCBkZW1vIG5vdCBmb3VuZCBpbiBtYW5pZmVzdCc7XG4gICAgICAgIHJldHVybjtcbiAgICAgIH1cblxuICAgICAgY29uc3QgZnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZGVtb0luZm9UZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cInRhZ05hbWVcIl0nKSEudGV4dENvbnRlbnQgPSBjdXJyZW50RGVtby50YWdOYW1lO1xuICAgICAgZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJjYW5vbmljYWxVUkxcIl0nKSEudGV4dENvbnRlbnQgPSBjdXJyZW50RGVtby5jYW5vbmljYWxVUkw7XG4gICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImxvY2FsUm91dGVcIl0nKSEudGV4dENvbnRlbnQgPSBjdXJyZW50RGVtby5sb2NhbFJvdXRlO1xuXG4gICAgICBjb25zdCBkZXNjcmlwdGlvbkdyb3VwID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQtZ3JvdXA9XCJkZXNjcmlwdGlvblwiXScpO1xuICAgICAgaWYgKGN1cnJlbnREZW1vLmRlc2NyaXB0aW9uKSB7XG4gICAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZGVzY3JpcHRpb25cIl0nKSEudGV4dENvbnRlbnQgPSBjdXJyZW50RGVtby5kZXNjcmlwdGlvbjtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGRlc2NyaXB0aW9uR3JvdXA/LnJlbW92ZSgpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBmaWxlcGF0aEdyb3VwID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQtZ3JvdXA9XCJmaWxlcGF0aFwiXScpO1xuICAgICAgaWYgKGN1cnJlbnREZW1vLmZpbGVwYXRoKSB7XG4gICAgICAgIGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiZmlsZXBhdGhcIl0nKSEudGV4dENvbnRlbnQgPSBjdXJyZW50RGVtby5maWxlcGF0aDtcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGZpbGVwYXRoR3JvdXA/LnJlbW92ZSgpO1xuICAgICAgfVxuXG4gICAgICBjb250YWluZXIucmVwbGFjZUNoaWxkcmVuKGZyYWdtZW50KTtcbiAgICB9IGVsc2Uge1xuICAgICAgY29uc3QgbGlzdEZyYWdtZW50ID0gQ2VtU2VydmVDaHJvbWUuI2RlbW9MaXN0VGVtcGxhdGUuY29udGVudC5jbG9uZU5vZGUodHJ1ZSkgYXMgRG9jdW1lbnRGcmFnbWVudDtcblxuICAgICAgY29uc3QgZ3JvdXBzQ29udGFpbmVyID0gbGlzdEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWNvbnRhaW5lcj1cImdyb3Vwc1wiXScpITtcblxuICAgICAgZm9yIChjb25zdCBkZW1vIG9mIGRlbW9zKSB7XG4gICAgICAgIGNvbnN0IGdyb3VwRnJhZ21lbnQgPSBDZW1TZXJ2ZUNocm9tZS4jZGVtb0dyb3VwVGVtcGxhdGUuY29udGVudC5jbG9uZU5vZGUodHJ1ZSkgYXMgRG9jdW1lbnRGcmFnbWVudDtcblxuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGFnTmFtZVwiXScpIS50ZXh0Q29udGVudCA9IGRlbW8udGFnTmFtZTtcbiAgICAgICAgZ3JvdXBGcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImRlc2NyaXB0aW9uXCJdJykhLnRleHRDb250ZW50ID1cbiAgICAgICAgICBkZW1vLmRlc2NyaXB0aW9uIHx8ICcobm8gZGVzY3JpcHRpb24pJztcbiAgICAgICAgZ3JvdXBGcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNhbm9uaWNhbFVSTFwiXScpIS50ZXh0Q29udGVudCA9IGRlbW8uY2Fub25pY2FsVVJMO1xuICAgICAgICBncm91cEZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibG9jYWxSb3V0ZVwiXScpIS50ZXh0Q29udGVudCA9IGRlbW8ubG9jYWxSb3V0ZTtcblxuICAgICAgICBjb25zdCBmaWxlcGF0aEdyb3VwID0gZ3JvdXBGcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZC1ncm91cD1cImZpbGVwYXRoXCJdJyk7XG4gICAgICAgIGlmIChkZW1vLmZpbGVwYXRoKSB7XG4gICAgICAgICAgZ3JvdXBGcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImZpbGVwYXRoXCJdJykhLnRleHRDb250ZW50ID0gZGVtby5maWxlcGF0aDtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBmaWxlcGF0aEdyb3VwPy5yZW1vdmUoKTtcbiAgICAgICAgfVxuXG4gICAgICAgIGdyb3Vwc0NvbnRhaW5lci5hcHBlbmRDaGlsZChncm91cEZyYWdtZW50KTtcbiAgICAgIH1cblxuICAgICAgY29udGFpbmVyLnJlcGxhY2VDaGlsZHJlbihsaXN0RnJhZ21lbnQpO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cExvZ0xpc3RlbmVyKCkge1xuICAgIHRoaXMuI2xvZ0NvbnRhaW5lciA9IHRoaXMuIyQoJ2xvZy1jb250YWluZXInKTtcblxuICAgIGNvbnN0IGxvZ3NGaWx0ZXIgPSB0aGlzLiMkKCdsb2dzLWZpbHRlcicpIGFzIEhUTUxJbnB1dEVsZW1lbnQgfCBudWxsO1xuICAgIGlmIChsb2dzRmlsdGVyKSB7XG4gICAgICBsb2dzRmlsdGVyLmFkZEV2ZW50TGlzdGVuZXIoJ2lucHV0JywgKCkgPT4ge1xuICAgICAgICBjb25zdCB7IHZhbHVlID0gJycgfSA9IGxvZ3NGaWx0ZXI7XG4gICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNsb2dzRmlsdGVyRGVib3VuY2VUaW1lciEpO1xuICAgICAgICB0aGlzLiNsb2dzRmlsdGVyRGVib3VuY2VUaW1lciA9IHNldFRpbWVvdXQoKCkgPT4ge1xuICAgICAgICAgIHRoaXMuI2ZpbHRlckxvZ3ModmFsdWUpO1xuICAgICAgICB9LCAzMDApO1xuICAgICAgfSk7XG4gICAgfVxuXG4gICAgdGhpcy4jbG9nTGV2ZWxEcm9wZG93biA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignI2xvZy1sZXZlbC1maWx0ZXInKSA/PyBudWxsO1xuICAgIGlmICh0aGlzLiNsb2dMZXZlbERyb3Bkb3duKSB7XG4gICAgICByZXF1ZXN0QW5pbWF0aW9uRnJhbWUoKCkgPT4ge1xuICAgICAgICB0aGlzLiNsb2FkTG9nRmlsdGVyU3RhdGUoKTtcbiAgICAgIH0pO1xuICAgICAgdGhpcy4jbG9nTGV2ZWxEcm9wZG93bi5hZGRFdmVudExpc3RlbmVyKCdzZWxlY3QnLCB0aGlzLiNoYW5kbGVMb2dGaWx0ZXJDaGFuZ2UgYXMgRXZlbnRMaXN0ZW5lcik7XG4gICAgfVxuXG4gICAgdGhpcy4jJCgnY29weS1sb2dzJyk/LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4ge1xuICAgICAgdGhpcy4jY29weUxvZ3MoKTtcbiAgICB9KTtcblxuICAgIHRoaXMuI2hhbmRsZUxvZ3NFdmVudCA9ICgoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgICBjb25zdCBsb2dzID0gKGV2ZW50IGFzIENlbUxvZ3NFdmVudCkubG9ncztcbiAgICAgIGlmIChsb2dzKSB7XG4gICAgICAgIHRoaXMuI3JlbmRlckxvZ3MobG9ncyk7XG4gICAgICB9XG4gICAgfSk7XG4gICAgd2luZG93LmFkZEV2ZW50TGlzdGVuZXIoJ2NlbTpsb2dzJywgdGhpcy4jaGFuZGxlTG9nc0V2ZW50KTtcbiAgfVxuXG4gICNmaWx0ZXJMb2dzKHF1ZXJ5OiBzdHJpbmcpIHtcbiAgICB0aGlzLiNsb2dzRmlsdGVyVmFsdWUgPSBxdWVyeS50b0xvd2VyQ2FzZSgpO1xuXG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGZvciAoY29uc3QgZW50cnkgb2YgdGhpcy4jbG9nQ29udGFpbmVyLmNoaWxkcmVuKSB7XG4gICAgICBjb25zdCB0ZXh0ID0gZW50cnkudGV4dENvbnRlbnQ/LnRvTG93ZXJDYXNlKCkgPz8gJyc7XG4gICAgICBjb25zdCB0ZXh0TWF0Y2ggPSAhdGhpcy4jbG9nc0ZpbHRlclZhbHVlIHx8IHRleHQuaW5jbHVkZXModGhpcy4jbG9nc0ZpbHRlclZhbHVlKTtcblxuICAgICAgY29uc3QgbG9nVHlwZSA9IHRoaXMuI2dldExvZ1R5cGVGcm9tRW50cnkoZW50cnkpO1xuICAgICAgY29uc3QgbGV2ZWxNYXRjaCA9IHRoaXMuI2xvZ0xldmVsRmlsdGVycy5oYXMobG9nVHlwZSk7XG5cbiAgICAgIChlbnRyeSBhcyBIVE1MRWxlbWVudCkuaGlkZGVuID0gISh0ZXh0TWF0Y2ggJiYgbGV2ZWxNYXRjaCk7XG4gICAgfVxuICB9XG5cbiAgI2dldExvZ1R5cGVGcm9tRW50cnkoZW50cnk6IEVsZW1lbnQpOiBzdHJpbmcge1xuICAgIGZvciAoY29uc3QgY2xzIG9mIGVudHJ5LmNsYXNzTGlzdCkge1xuICAgICAgaWYgKFsnaW5mbycsICd3YXJuaW5nJywgJ2Vycm9yJywgJ2RlYnVnJ10uaW5jbHVkZXMoY2xzKSkge1xuICAgICAgICByZXR1cm4gY2xzID09PSAnd2FybmluZycgPyAnd2FybicgOiBjbHM7XG4gICAgICB9XG4gICAgfVxuICAgIHJldHVybiAnaW5mbyc7XG4gIH1cblxuICAjbG9hZExvZ0ZpbHRlclN0YXRlKCkge1xuICAgIHRyeSB7XG4gICAgICBjb25zdCBzYXZlZCA9IGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtbG9nLWZpbHRlcnMnKTtcbiAgICAgIGlmIChzYXZlZCkge1xuICAgICAgICB0aGlzLiNsb2dMZXZlbEZpbHRlcnMgPSBuZXcgU2V0KEpTT04ucGFyc2Uoc2F2ZWQpKTtcbiAgICAgIH1cbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICBjb25zb2xlLmRlYnVnKCdbY2VtLXNlcnZlLWNocm9tZV0gbG9jYWxTdG9yYWdlIHVuYXZhaWxhYmxlLCB1c2luZyBkZWZhdWx0IGxvZyBmaWx0ZXJzJyk7XG4gICAgfVxuICAgIHRoaXMuI3N5bmNDaGVja2JveFN0YXRlcygpO1xuICB9XG5cbiAgI3N5bmNDaGVja2JveFN0YXRlcygpIHtcbiAgICBpZiAoIXRoaXMuI2xvZ0xldmVsRHJvcGRvd24pIHJldHVybjtcbiAgICBjb25zdCBtZW51SXRlbXMgPSB0aGlzLiNsb2dMZXZlbERyb3Bkb3duLnF1ZXJ5U2VsZWN0b3JBbGwoJ3BmLXY2LW1lbnUtaXRlbScpO1xuICAgIG1lbnVJdGVtcy5mb3JFYWNoKGl0ZW0gPT4ge1xuICAgICAgY29uc3QgdmFsdWUgPSAoaXRlbSBhcyBhbnkpLnZhbHVlO1xuICAgICAgKGl0ZW0gYXMgYW55KS5jaGVja2VkID0gdGhpcy4jbG9nTGV2ZWxGaWx0ZXJzLmhhcyh2YWx1ZSk7XG4gICAgfSk7XG4gIH1cblxuICAjc2F2ZUxvZ0ZpbHRlclN0YXRlKCkge1xuICAgIHRyeSB7XG4gICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSgnY2VtLXNlcnZlLWxvZy1maWx0ZXJzJyxcbiAgICAgICAgSlNPTi5zdHJpbmdpZnkoWy4uLnRoaXMuI2xvZ0xldmVsRmlsdGVyc10pKTtcbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICAvLyBsb2NhbFN0b3JhZ2UgdW5hdmFpbGFibGUgKHByaXZhdGUgbW9kZSksIHNpbGVudGx5IGNvbnRpbnVlXG4gICAgfVxuICB9XG5cbiAgI2hhbmRsZUxvZ0ZpbHRlckNoYW5nZSA9IChldmVudDogRXZlbnQpID0+IHtcbiAgICBjb25zdCB7IHZhbHVlLCBjaGVja2VkIH0gPSBldmVudCBhcyBFdmVudCAmIHsgdmFsdWU6IHN0cmluZzsgY2hlY2tlZDogYm9vbGVhbiB9O1xuXG4gICAgaWYgKGNoZWNrZWQpIHtcbiAgICAgIHRoaXMuI2xvZ0xldmVsRmlsdGVycy5hZGQodmFsdWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuZGVsZXRlKHZhbHVlKTtcbiAgICB9XG5cbiAgICB0aGlzLiNzYXZlTG9nRmlsdGVyU3RhdGUoKTtcbiAgICB0aGlzLiNmaWx0ZXJMb2dzKHRoaXMuI2xvZ3NGaWx0ZXJWYWx1ZSk7XG4gIH07XG5cbiAgYXN5bmMgI2NvcHlMb2dzKCkge1xuICAgIGlmICghdGhpcy4jbG9nQ29udGFpbmVyKSByZXR1cm47XG5cbiAgICBjb25zdCBsb2dzID0gQXJyYXkuZnJvbSh0aGlzLiNsb2dDb250YWluZXIuY2hpbGRyZW4pXG4gICAgICAuZmlsdGVyKGVudHJ5ID0+ICEoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmhpZGRlbilcbiAgICAgIC5tYXAoZW50cnkgPT4ge1xuICAgICAgICBjb25zdCB0eXBlID0gZW50cnkucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsYWJlbFwiXScpPy50ZXh0Q29udGVudD8udHJpbSgpIHx8ICdJTkZPJztcbiAgICAgICAgY29uc3QgdGltZSA9IGVudHJ5LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGltZVwiXScpPy50ZXh0Q29udGVudD8udHJpbSgpIHx8ICcnO1xuICAgICAgICBjb25zdCBtZXNzYWdlID0gZW50cnkucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJtZXNzYWdlXCJdJyk/LnRleHRDb250ZW50Py50cmltKCkgfHwgJyc7XG4gICAgICAgIHJldHVybiBgWyR7dHlwZX1dICR7dGltZX0gJHttZXNzYWdlfWA7XG4gICAgICB9KS5qb2luKCdcXG4nKTtcblxuICAgIGlmICghbG9ncykgcmV0dXJuO1xuXG4gICAgdHJ5IHtcbiAgICAgIGF3YWl0IG5hdmlnYXRvci5jbGlwYm9hcmQud3JpdGVUZXh0KGxvZ3MpO1xuICAgICAgY29uc3QgYnRuID0gdGhpcy4jJCgnY29weS1sb2dzJyk7XG4gICAgICBpZiAoYnRuKSB7XG4gICAgICAgIGNvbnN0IHRleHROb2RlID0gQXJyYXkuZnJvbShidG4uY2hpbGROb2RlcykuZmluZChcbiAgICAgICAgICBuID0+IG4ubm9kZVR5cGUgPT09IE5vZGUuVEVYVF9OT0RFICYmIChuLnRleHRDb250ZW50Py50cmltKCkubGVuZ3RoID8/IDApID4gMFxuICAgICAgICApO1xuICAgICAgICBpZiAodGV4dE5vZGUpIHtcbiAgICAgICAgICBjb25zdCBvcmlnaW5hbCA9IHRleHROb2RlLnRleHRDb250ZW50O1xuICAgICAgICAgIHRleHROb2RlLnRleHRDb250ZW50ID0gJ0NvcGllZCEnO1xuXG4gICAgICAgICAgaWYgKHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0KSB7XG4gICAgICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIHRoaXMuI2NvcHlMb2dzRmVlZGJhY2tUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgICBpZiAodGhpcy5pc0Nvbm5lY3RlZCAmJiB0ZXh0Tm9kZS5wYXJlbnROb2RlKSB7XG4gICAgICAgICAgICAgIHRleHROb2RlLnRleHRDb250ZW50ID0gb3JpZ2luYWw7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgICB0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgICAgICAgfSwgMjAwMCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9IGNhdGNoIChlcnIpIHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY29weSBsb2dzOicsIGVycik7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwRGVidWdPdmVybGF5KCkge1xuICAgIGNvbnN0IGRlYnVnQnV0dG9uID0gdGhpcy4jJCgnZGVidWctaW5mbycpO1xuICAgIGNvbnN0IGRlYnVnTW9kYWwgPSB0aGlzLiMkKCdkZWJ1Zy1tb2RhbCcpO1xuICAgIGNvbnN0IGRlYnVnQ2xvc2UgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJy5kZWJ1Zy1jbG9zZScpO1xuICAgIGNvbnN0IGRlYnVnQ29weSA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignLmRlYnVnLWNvcHknKTtcblxuICAgIGlmIChkZWJ1Z0J1dHRvbiAmJiBkZWJ1Z01vZGFsKSB7XG4gICAgICBkZWJ1Z0J1dHRvbi5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgICAgdGhpcy4jZmV0Y2hEZWJ1Z0luZm8oKTtcbiAgICAgICAgKGRlYnVnTW9kYWwgYXMgYW55KS5zaG93TW9kYWwoKTtcbiAgICAgIH0pO1xuXG4gICAgICBkZWJ1Z0Nsb3NlPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IChkZWJ1Z01vZGFsIGFzIGFueSkuY2xvc2UoKSk7XG5cbiAgICAgIGRlYnVnQ29weT8uYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoKSA9PiB7XG4gICAgICAgIHRoaXMuI2NvcHlEZWJ1Z0luZm8oKTtcbiAgICAgIH0pO1xuICAgIH1cbiAgfVxuXG4gICNzZXR1cEZvb3RlckRyYXdlcigpIHtcbiAgICBjb25zdCBkcmF3ZXIgPSB0aGlzLnNoYWRvd1Jvb3Q/LnF1ZXJ5U2VsZWN0b3IoJ2NlbS1kcmF3ZXInKTtcbiAgICBjb25zdCB0YWJzID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdwZi12Ni10YWJzJyk7XG5cbiAgICBpZiAoIWRyYXdlciB8fCAhdGFicykgcmV0dXJuO1xuXG4gICAgdGhpcy4jZHJhd2VyT3BlbiA9IChkcmF3ZXIgYXMgYW55KS5vcGVuO1xuXG4gICAgZHJhd2VyLmFkZEV2ZW50TGlzdGVuZXIoJ2NoYW5nZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgdGhpcy4jZHJhd2VyT3BlbiA9IChlIGFzIGFueSkub3BlbjtcblxuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIGRyYXdlcjogeyBvcGVuOiAoZSBhcyBhbnkpLm9wZW4gfVxuICAgICAgfSk7XG5cbiAgICAgIGlmICgoZSBhcyBhbnkpLm9wZW4pIHtcbiAgICAgICAgdGhpcy4jc2Nyb2xsTG9nc1RvQm90dG9tKCk7XG4gICAgICB9XG4gICAgfSk7XG5cbiAgICBkcmF3ZXIuYWRkRXZlbnRMaXN0ZW5lcigncmVzaXplJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICAoZHJhd2VyIGFzIGFueSkuc2V0QXR0cmlidXRlKCdkcmF3ZXItaGVpZ2h0JywgKGUgYXMgYW55KS5oZWlnaHQpO1xuXG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVN0YXRlKHtcbiAgICAgICAgZHJhd2VyOiB7IGhlaWdodDogKGUgYXMgYW55KS5oZWlnaHQgfVxuICAgICAgfSk7XG4gICAgfSk7XG5cbiAgICB0YWJzLmFkZEV2ZW50TGlzdGVuZXIoJ2NoYW5nZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgU3RhdGVQZXJzaXN0ZW5jZS51cGRhdGVTdGF0ZSh7XG4gICAgICAgIHRhYnM6IHsgc2VsZWN0ZWRJbmRleDogKGUgYXMgYW55KS5zZWxlY3RlZEluZGV4IH1cbiAgICAgIH0pO1xuXG4gICAgICBpZiAoKGUgYXMgYW55KS5zZWxlY3RlZEluZGV4ID09PSAyICYmIChkcmF3ZXIgYXMgYW55KS5vcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExvZ3NUb0JvdHRvbSgpO1xuICAgICAgfVxuXG4gICAgICBpZiAoKGUgYXMgYW55KS5zZWxlY3RlZEluZGV4ID09PSAzICYmIChkcmF3ZXIgYXMgYW55KS5vcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbEV2ZW50c1RvQm90dG9tKCk7XG4gICAgICB9XG4gICAgfSk7XG4gIH1cblxuICAjZGV0ZWN0QnJvd3NlcigpOiBzdHJpbmcge1xuICAgIGNvbnN0IHVhID0gbmF2aWdhdG9yLnVzZXJBZ2VudDtcbiAgICBpZiAodWEuaW5jbHVkZXMoJ0ZpcmVmb3gvJykpIHtcbiAgICAgIGNvbnN0IG1hdGNoID0gdWEubWF0Y2goL0ZpcmVmb3hcXC8oXFxkKykvKTtcbiAgICAgIHJldHVybiBtYXRjaCA/IGBGaXJlZm94ICR7bWF0Y2hbMV19YCA6ICdGaXJlZm94JztcbiAgICB9IGVsc2UgaWYgKHVhLmluY2x1ZGVzKCdFZGcvJykpIHtcbiAgICAgIGNvbnN0IG1hdGNoID0gdWEubWF0Y2goL0VkZ1xcLyhcXGQrKS8pO1xuICAgICAgcmV0dXJuIG1hdGNoID8gYEVkZ2UgJHttYXRjaFsxXX1gIDogJ0VkZ2UnO1xuICAgIH0gZWxzZSBpZiAodWEuaW5jbHVkZXMoJ0Nocm9tZS8nKSkge1xuICAgICAgY29uc3QgbWF0Y2ggPSB1YS5tYXRjaCgvQ2hyb21lXFwvKFxcZCspLyk7XG4gICAgICByZXR1cm4gbWF0Y2ggPyBgQ2hyb21lICR7bWF0Y2hbMV19YCA6ICdDaHJvbWUnO1xuICAgIH0gZWxzZSBpZiAodWEuaW5jbHVkZXMoJ1NhZmFyaS8nKSAmJiAhdWEuaW5jbHVkZXMoJ0Nocm9tZScpKSB7XG4gICAgICBjb25zdCBtYXRjaCA9IHVhLm1hdGNoKC9WZXJzaW9uXFwvKFxcZCspLyk7XG4gICAgICByZXR1cm4gbWF0Y2ggPyBgU2FmYXJpICR7bWF0Y2hbMV19YCA6ICdTYWZhcmknO1xuICAgIH1cbiAgICByZXR1cm4gJ1Vua25vd24nO1xuICB9XG5cbiAgYXN5bmMgI2NvcHlEZWJ1Z0luZm8oKSB7XG4gICAgY29uc3QgaW5mbyA9IEFycmF5LmZyb20odGhpcy4jJCQoJyNkZWJ1Zy1tb2RhbCBkbCBkdCcpKS5tYXAoZHQgPT4ge1xuICAgICAgY29uc3QgZGQgPSBkdC5uZXh0RWxlbWVudFNpYmxpbmc7XG4gICAgICBpZiAoZGQgJiYgZGQudGFnTmFtZSA9PT0gJ0REJykge1xuICAgICAgICByZXR1cm4gYCR7ZHQudGV4dENvbnRlbnR9OiAke2RkLnRleHRDb250ZW50fWA7XG4gICAgICB9XG4gICAgICByZXR1cm4gJyc7XG4gICAgfSkuam9pbignXFxuJyk7XG5cbiAgICBsZXQgaW1wb3J0TWFwU2VjdGlvbiA9ICcnO1xuICAgIGlmICh0aGlzLiNkZWJ1Z0RhdGE/LmltcG9ydE1hcEpTT04pIHtcbiAgICAgIGltcG9ydE1hcFNlY3Rpb24gPSBgXFxuJHsnPScucmVwZWF0KDQwKX1cXG5JbXBvcnQgTWFwOlxcbiR7Jz0nLnJlcGVhdCg0MCl9XFxuJHt0aGlzLiNkZWJ1Z0RhdGEuaW1wb3J0TWFwSlNPTn1cXG5gO1xuICAgIH1cblxuICAgIGNvbnN0IGRlYnVnVGV4dCA9IGBDRU0gU2VydmUgRGVidWcgSW5mb3JtYXRpb25cbiR7Jz0nLnJlcGVhdCg0MCl9XG4ke2luZm99JHtpbXBvcnRNYXBTZWN0aW9ufVxuJHsnPScucmVwZWF0KDQwKX1cbkdlbmVyYXRlZDogJHtuZXcgRGF0ZSgpLnRvSVNPU3RyaW5nKCl9YDtcblxuICAgIHRyeSB7XG4gICAgICBhd2FpdCBuYXZpZ2F0b3IuY2xpcGJvYXJkLndyaXRlVGV4dChkZWJ1Z1RleHQpO1xuICAgICAgY29uc3QgY29weUJ1dHRvbiA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignLmRlYnVnLWNvcHknKTtcbiAgICAgIGlmIChjb3B5QnV0dG9uKSB7XG4gICAgICAgIGNvbnN0IG9yaWdpbmFsVGV4dCA9IGNvcHlCdXR0b24udGV4dENvbnRlbnQ7XG4gICAgICAgIGNvcHlCdXR0b24udGV4dENvbnRlbnQgPSAnQ29waWVkISc7XG5cbiAgICAgICAgaWYgKHRoaXMuI2NvcHlEZWJ1Z0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgICB9XG5cbiAgICAgICAgdGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0ID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgaWYgKHRoaXMuaXNDb25uZWN0ZWQgJiYgY29weUJ1dHRvbi5wYXJlbnROb2RlKSB7XG4gICAgICAgICAgICBjb3B5QnV0dG9uLnRleHRDb250ZW50ID0gb3JpZ2luYWxUZXh0O1xuICAgICAgICAgIH1cbiAgICAgICAgICB0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgICAgICB9LCAyMDAwKTtcbiAgICAgIH1cbiAgICB9IGNhdGNoIChlcnIpIHtcbiAgICAgIGNvbnNvbGUuZXJyb3IoJ1tjZW0tc2VydmUtY2hyb21lXSBGYWlsZWQgdG8gY29weSBkZWJ1ZyBpbmZvOicsIGVycik7XG4gICAgfVxuICB9XG5cbiAgI3JlbmRlckxvZ3MobG9nczogQXJyYXk8eyB0eXBlOiBzdHJpbmc7IGRhdGU6IHN0cmluZzsgbWVzc2FnZTogc3RyaW5nIH0+KSB7XG4gICAgaWYgKCF0aGlzLiNsb2dDb250YWluZXIpIHJldHVybjtcblxuICAgIGNvbnN0IGxvZ0VsZW1lbnRzID0gbG9ncy5tYXAobG9nID0+IHtcbiAgICAgIGNvbnN0IGZyYWdtZW50ID0gQ2VtU2VydmVDaHJvbWUuI2xvZ0VudHJ5VGVtcGxhdGUuY29udGVudC5jbG9uZU5vZGUodHJ1ZSkgYXMgRG9jdW1lbnRGcmFnbWVudDtcblxuICAgICAgY29uc3QgZGF0ZSA9IG5ldyBEYXRlKGxvZy5kYXRlKTtcbiAgICAgIGNvbnN0IHRpbWUgPSBkYXRlLnRvTG9jYWxlVGltZVN0cmluZygpO1xuXG4gICAgICBjb25zdCBjb250YWluZXIgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImNvbnRhaW5lclwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgICAgY29udGFpbmVyLmNsYXNzTGlzdC5hZGQobG9nLnR5cGUpO1xuICAgICAgY29udGFpbmVyLnNldEF0dHJpYnV0ZSgnZGF0YS1sb2ctaWQnLCBsb2cuZGF0ZSk7XG5cbiAgICAgIGNvbnN0IHR5cGVMYWJlbCA9IHRoaXMuI2dldExvZ0JhZGdlKGxvZy50eXBlKTtcbiAgICAgIGNvbnN0IHNlYXJjaFRleHQgPSBgJHt0eXBlTGFiZWx9ICR7dGltZX0gJHtsb2cubWVzc2FnZX1gLnRvTG93ZXJDYXNlKCk7XG4gICAgICBjb25zdCB0ZXh0TWF0Y2ggPSAhdGhpcy4jbG9nc0ZpbHRlclZhbHVlIHx8IHNlYXJjaFRleHQuaW5jbHVkZXModGhpcy4jbG9nc0ZpbHRlclZhbHVlKTtcblxuICAgICAgY29uc3QgbG9nVHlwZUZvckZpbHRlciA9IGxvZy50eXBlID09PSAnd2FybmluZycgPyAnd2FybicgOiBsb2cudHlwZTtcbiAgICAgIGNvbnN0IGxldmVsTWF0Y2ggPSB0aGlzLiNsb2dMZXZlbEZpbHRlcnMuaGFzKGxvZ1R5cGVGb3JGaWx0ZXIpO1xuXG4gICAgICBpZiAoISh0ZXh0TWF0Y2ggJiYgbGV2ZWxNYXRjaCkpIHtcbiAgICAgICAgY29udGFpbmVyLnNldEF0dHJpYnV0ZSgnaGlkZGVuJywgJycpO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBsYWJlbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwibGFiZWxcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICAgIGxhYmVsLnRleHRDb250ZW50ID0gdGhpcy4jZ2V0TG9nQmFkZ2UobG9nLnR5cGUpO1xuICAgICAgdGhpcy4jYXBwbHlMb2dMYWJlbEF0dHJzKGxhYmVsLCBsb2cudHlwZSk7XG5cbiAgICAgIGNvbnN0IHRpbWVFbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGltZVwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgICAgdGltZUVsLnNldEF0dHJpYnV0ZSgnZGF0ZXRpbWUnLCBsb2cuZGF0ZSk7XG4gICAgICB0aW1lRWwudGV4dENvbnRlbnQgPSB0aW1lO1xuXG4gICAgICAoZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJtZXNzYWdlXCJdJykgYXMgSFRNTEVsZW1lbnQpLnRleHRDb250ZW50ID0gbG9nLm1lc3NhZ2U7XG5cbiAgICAgIHJldHVybiBmcmFnbWVudDtcbiAgICB9KTtcblxuICAgIGlmICghdGhpcy4jaW5pdGlhbExvZ3NGZXRjaGVkKSB7XG4gICAgICB0aGlzLiNsb2dDb250YWluZXIucmVwbGFjZUNoaWxkcmVuKC4uLmxvZ0VsZW1lbnRzKTtcbiAgICAgIHRoaXMuI2luaXRpYWxMb2dzRmV0Y2hlZCA9IHRydWU7XG5cbiAgICAgIGlmICh0aGlzLiNkcmF3ZXJPcGVuKSB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExhdGVzdEludG9WaWV3KCk7XG4gICAgICB9XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMuI2xvZ0NvbnRhaW5lci5hcHBlbmQoLi4ubG9nRWxlbWVudHMpO1xuXG4gICAgICBpZiAodGhpcy4jZHJhd2VyT3Blbikge1xuICAgICAgICB0aGlzLiNzY3JvbGxMYXRlc3RJbnRvVmlldygpO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNnZXRMb2dCYWRnZSh0eXBlOiBzdHJpbmcpOiBzdHJpbmcge1xuICAgIHN3aXRjaCAodHlwZSkge1xuICAgICAgY2FzZSAnaW5mbyc6IHJldHVybiAnSW5mbyc7XG4gICAgICBjYXNlICd3YXJuaW5nJzogcmV0dXJuICdXYXJuJztcbiAgICAgIGNhc2UgJ2Vycm9yJzogcmV0dXJuICdFcnJvcic7XG4gICAgICBjYXNlICdkZWJ1Zyc6IHJldHVybiAnRGVidWcnO1xuICAgICAgZGVmYXVsdDogcmV0dXJuIHR5cGUudG9VcHBlckNhc2UoKTtcbiAgICB9XG4gIH1cblxuICAjYXBwbHlMb2dMYWJlbEF0dHJzKGxhYmVsOiBIVE1MRWxlbWVudCwgdHlwZTogc3RyaW5nKSB7XG4gICAgc3dpdGNoICh0eXBlKSB7XG4gICAgICBjYXNlICdpbmZvJzpcbiAgICAgICAgcmV0dXJuIGxhYmVsLnNldEF0dHJpYnV0ZSgnc3RhdHVzJywgJ2luZm8nKTtcbiAgICAgIGNhc2UgJ3dhcm5pbmcnOlxuICAgICAgICByZXR1cm4gbGFiZWwuc2V0QXR0cmlidXRlKCdzdGF0dXMnLCAnd2FybmluZycpO1xuICAgICAgY2FzZSAnZXJyb3InOlxuICAgICAgICByZXR1cm4gbGFiZWwuc2V0QXR0cmlidXRlKCdzdGF0dXMnLCAnZGFuZ2VyJyk7XG4gICAgICBjYXNlICdkZWJ1Zyc6XG4gICAgICAgIHJldHVybiBsYWJlbC5zZXRBdHRyaWJ1dGUoJ2NvbG9yJywgJ3B1cnBsZScpO1xuICAgICAgZGVmYXVsdDpcbiAgICAgICAgbGFiZWwuc2V0QXR0cmlidXRlKCdjb2xvcicsICdncmV5Jyk7XG4gICAgfVxuICB9XG5cbiAgI3Njcm9sbExhdGVzdEludG9WaWV3KCkge1xuICAgIGlmICghdGhpcy4jbG9nQ29udGFpbmVyKSByZXR1cm47XG5cbiAgICByZXF1ZXN0QW5pbWF0aW9uRnJhbWUoKCkgPT4ge1xuICAgICAgY29uc3QgbGFzdExvZyA9IHRoaXMuI2xvZ0NvbnRhaW5lciEubGFzdEVsZW1lbnRDaGlsZDtcbiAgICAgIGlmIChsYXN0TG9nKSB7XG4gICAgICAgIGxhc3RMb2cuc2Nyb2xsSW50b1ZpZXcoeyBiZWhhdmlvcjogJ2F1dG8nLCBibG9jazogJ2VuZCcgfSk7XG4gICAgICB9XG4gICAgfSk7XG4gIH1cblxuICAjc2Nyb2xsTG9nc1RvQm90dG9tKCkge1xuICAgIGlmICghdGhpcy4jbG9nQ29udGFpbmVyKSByZXR1cm47XG5cbiAgICBpZiAodGhpcy4jaXNJbml0aWFsTG9hZCkge1xuICAgICAgdGhpcy4jc2Nyb2xsTGF0ZXN0SW50b1ZpZXcoKTtcbiAgICB9IGVsc2Uge1xuICAgICAgc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgIHRoaXMuI3Njcm9sbExhdGVzdEludG9WaWV3KCk7XG4gICAgICB9LCAzNTApO1xuICAgIH1cbiAgfVxuXG4gICNtaWdyYXRlRnJvbUxvY2FsU3RvcmFnZUlmTmVlZGVkKCkge1xuICAgIHRyeSB7XG4gICAgICBjb25zdCBoYXNMb2NhbFN0b3JhZ2UgPVxuICAgICAgICBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWNvbG9yLXNjaGVtZScpICE9PSBudWxsIHx8XG4gICAgICAgIGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtZHJhd2VyLW9wZW4nKSAhPT0gbnVsbCB8fFxuICAgICAgICBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWRyYXdlci1oZWlnaHQnKSAhPT0gbnVsbCB8fFxuICAgICAgICBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWFjdGl2ZS10YWInKSAhPT0gbnVsbDtcblxuICAgICAgaWYgKGhhc0xvY2FsU3RvcmFnZSkge1xuICAgICAgICBjb25zdCBtaWdyYXRlZCA9IGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtbWlncmF0ZWQtdG8tY29va2llcycpO1xuICAgICAgICBpZiAoIW1pZ3JhdGVkKSB7XG4gICAgICAgICAgU3RhdGVQZXJzaXN0ZW5jZS5taWdyYXRlRnJvbUxvY2FsU3RvcmFnZSgpO1xuICAgICAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdjZW0tc2VydmUtbWlncmF0ZWQtdG8tY29va2llcycsICd0cnVlJyk7XG4gICAgICAgICAgc2V0VGltZW91dCgoKSA9PiB3aW5kb3cubG9jYXRpb24ucmVsb2FkKCksIDEwMCk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICAvLyBsb2NhbFN0b3JhZ2Ugbm90IGF2YWlsYWJsZSwgc2tpcCBtaWdyYXRpb25cbiAgICB9XG4gIH1cblxuICAjc2V0dXBDb2xvclNjaGVtZVRvZ2dsZSgpIHtcbiAgICBjb25zdCB0b2dnbGVHcm91cCA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcignLmNvbG9yLXNjaGVtZS10b2dnbGUnKTtcbiAgICBpZiAoIXRvZ2dsZUdyb3VwKSByZXR1cm47XG5cbiAgICBjb25zdCBzdGF0ZSA9IFN0YXRlUGVyc2lzdGVuY2UuZ2V0U3RhdGUoKTtcblxuICAgIHRoaXMuI2FwcGx5Q29sb3JTY2hlbWUoc3RhdGUuY29sb3JTY2hlbWUpO1xuXG4gICAgY29uc3QgaXRlbXMgPSB0b2dnbGVHcm91cC5xdWVyeVNlbGVjdG9yQWxsKCdwZi12Ni10b2dnbGUtZ3JvdXAtaXRlbScpO1xuICAgIGl0ZW1zLmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBpZiAoKGl0ZW0gYXMgYW55KS52YWx1ZSA9PT0gc3RhdGUuY29sb3JTY2hlbWUpIHtcbiAgICAgICAgaXRlbS5zZXRBdHRyaWJ1dGUoJ3NlbGVjdGVkJywgJycpO1xuICAgICAgfVxuICAgIH0pO1xuXG4gICAgdG9nZ2xlR3JvdXAuYWRkRXZlbnRMaXN0ZW5lcigncGYtdjYtdG9nZ2xlLWdyb3VwLWNoYW5nZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgY29uc3Qgc2NoZW1lID0gKGUgYXMgYW55KS52YWx1ZTtcbiAgICAgIHRoaXMuI2FwcGx5Q29sb3JTY2hlbWUoc2NoZW1lKTtcbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoeyBjb2xvclNjaGVtZTogc2NoZW1lIH0pO1xuICAgIH0pO1xuICB9XG5cbiAgI2FwcGx5Q29sb3JTY2hlbWUoc2NoZW1lOiBzdHJpbmcpIHtcbiAgICBjb25zdCBib2R5ID0gZG9jdW1lbnQuYm9keTtcblxuICAgIHN3aXRjaCAoc2NoZW1lKSB7XG4gICAgICBjYXNlICdsaWdodCc6XG4gICAgICAgIGJvZHkuc3R5bGUuY29sb3JTY2hlbWUgPSAnbGlnaHQnO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ2RhcmsnOlxuICAgICAgICBib2R5LnN0eWxlLmNvbG9yU2NoZW1lID0gJ2RhcmsnO1xuICAgICAgICBicmVhaztcbiAgICAgIGNhc2UgJ3N5c3RlbSc6XG4gICAgICBkZWZhdWx0OlxuICAgICAgICBib2R5LnN0eWxlLmNvbG9yU2NoZW1lID0gJ2xpZ2h0IGRhcmsnO1xuICAgICAgICBicmVhaztcbiAgICB9XG4gIH1cblxuICAjc2V0dXBLbm9iQ29vcmRpbmF0aW9uKCkge1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna25vYjphdHRyaWJ1dGUtY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6cHJvcGVydHktY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOmF0dHJpYnV0ZS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2tub2I6cHJvcGVydHktY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrbm9iOmNzcy1wcm9wZXJ0eS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgfVxuXG4gICNvbktub2JDaGFuZ2UgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gdGhpcy4jZ2V0S25vYlRhcmdldChldmVudCk7XG4gICAgaWYgKCF0YXJnZXQpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIENvdWxkIG5vdCBmaW5kIGtub2IgdGFyZ2V0IGluZm8gaW4gZXZlbnQgcGF0aCcpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIGNvbnN0IHsgdGFnTmFtZSwgaW5zdGFuY2VJbmRleCB9ID0gdGFyZ2V0O1xuXG4gICAgY29uc3QgZGVtbyA9IHRoaXMuZGVtbztcbiAgICBpZiAoIWRlbW8pIHJldHVybjtcblxuICAgIGNvbnN0IGtub2JUeXBlID0gdGhpcy4jZ2V0S25vYlR5cGVGcm9tRXZlbnQoZXZlbnQpO1xuXG4gICAgY29uc3Qgc3VjY2VzcyA9IChkZW1vIGFzIGFueSkuYXBwbHlLbm9iQ2hhbmdlKFxuICAgICAga25vYlR5cGUsXG4gICAgICAoZXZlbnQgYXMgYW55KS5uYW1lLFxuICAgICAgKGV2ZW50IGFzIGFueSkudmFsdWUsXG4gICAgICB0YWdOYW1lLFxuICAgICAgaW5zdGFuY2VJbmRleFxuICAgICk7XG5cbiAgICBpZiAoIXN1Y2Nlc3MpIHtcbiAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIEZhaWxlZCB0byBhcHBseSBrbm9iIGNoYW5nZTonLCB7XG4gICAgICAgIHR5cGU6IGtub2JUeXBlLFxuICAgICAgICBuYW1lOiAoZXZlbnQgYXMgYW55KS5uYW1lLFxuICAgICAgICB0YWdOYW1lLFxuICAgICAgICBpbnN0YW5jZUluZGV4XG4gICAgICB9KTtcbiAgICB9XG4gIH07XG5cbiAgI29uS25vYkNsZWFyID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHRhcmdldCA9IHRoaXMuI2dldEtub2JUYXJnZXQoZXZlbnQpO1xuICAgIGlmICghdGFyZ2V0KSB7XG4gICAgICBjb25zb2xlLndhcm4oJ1tjZW0tc2VydmUtY2hyb21lXSBDb3VsZCBub3QgZmluZCBrbm9iIHRhcmdldCBpbmZvIGluIGV2ZW50IHBhdGgnKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICBjb25zdCB7IHRhZ05hbWUsIGluc3RhbmNlSW5kZXggfSA9IHRhcmdldDtcblxuICAgIGNvbnN0IGRlbW8gPSB0aGlzLmRlbW87XG4gICAgaWYgKCFkZW1vKSByZXR1cm47XG5cbiAgICBjb25zdCBrbm9iVHlwZSA9IHRoaXMuI2dldEtub2JUeXBlRnJvbUNsZWFyRXZlbnQoZXZlbnQpO1xuICAgIGNvbnN0IGNsZWFyVmFsdWUgPSBrbm9iVHlwZSA9PT0gJ3Byb3BlcnR5JyA/IHVuZGVmaW5lZCA6ICcnO1xuXG4gICAgY29uc3Qgc3VjY2VzcyA9IChkZW1vIGFzIGFueSkuYXBwbHlLbm9iQ2hhbmdlKFxuICAgICAga25vYlR5cGUsXG4gICAgICAoZXZlbnQgYXMgYW55KS5uYW1lLFxuICAgICAgY2xlYXJWYWx1ZSxcbiAgICAgIHRhZ05hbWUsXG4gICAgICBpbnN0YW5jZUluZGV4XG4gICAgKTtcblxuICAgIGlmICghc3VjY2Vzcykge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGNsZWFyIGtub2I6Jywge1xuICAgICAgICB0eXBlOiBrbm9iVHlwZSxcbiAgICAgICAgbmFtZTogKGV2ZW50IGFzIGFueSkubmFtZSxcbiAgICAgICAgdGFnTmFtZSxcbiAgICAgICAgaW5zdGFuY2VJbmRleFxuICAgICAgfSk7XG4gICAgfVxuICB9O1xuXG4gICNnZXRLbm9iVGFyZ2V0KGV2ZW50OiBFdmVudCk6IHsgdGFnTmFtZTogc3RyaW5nOyBpbnN0YW5jZUluZGV4OiBudW1iZXIgfSB8IG51bGwge1xuICAgIGNvbnN0IGRlZmF1bHRUYWdOYW1lID0gdGhpcy5wcmltYXJ5VGFnTmFtZSB8fCAnJztcblxuICAgIGlmIChldmVudC5jb21wb3NlZFBhdGgpIHtcbiAgICAgIGZvciAoY29uc3QgZWxlbWVudCBvZiBldmVudC5jb21wb3NlZFBhdGgoKSkge1xuICAgICAgICBpZiAoIShlbGVtZW50IGluc3RhbmNlb2YgRWxlbWVudCkpIGNvbnRpbnVlO1xuXG4gICAgICAgIGlmICgoZWxlbWVudCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldD8uaXNFbGVtZW50S25vYiA9PT0gJ3RydWUnKSB7XG4gICAgICAgICAgY29uc3QgdGFnTmFtZSA9IChlbGVtZW50IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LnRhZ05hbWUgfHwgZGVmYXVsdFRhZ05hbWU7XG4gICAgICAgICAgbGV0IGluc3RhbmNlSW5kZXggPSBOdW1iZXIucGFyc2VJbnQoKGVsZW1lbnQgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuaW5zdGFuY2VJbmRleCA/PyAnJywgMTApO1xuICAgICAgICAgIGlmIChOdW1iZXIuaXNOYU4oaW5zdGFuY2VJbmRleCkpIGluc3RhbmNlSW5kZXggPSAwO1xuICAgICAgICAgIHJldHVybiB7IHRhZ05hbWUsIGluc3RhbmNlSW5kZXggfTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiB7IHRhZ05hbWU6IGRlZmF1bHRUYWdOYW1lLCBpbnN0YW5jZUluZGV4OiAwIH07XG4gIH1cblxuICAjZ2V0S25vYlR5cGVGcm9tRXZlbnQoZXZlbnQ6IEV2ZW50KTogc3RyaW5nIHtcbiAgICBzd2l0Y2ggKGV2ZW50LnR5cGUpIHtcbiAgICAgIGNhc2UgJ2tub2I6YXR0cmlidXRlLWNoYW5nZSc6XG4gICAgICAgIHJldHVybiAnYXR0cmlidXRlJztcbiAgICAgIGNhc2UgJ2tub2I6cHJvcGVydHktY2hhbmdlJzpcbiAgICAgICAgcmV0dXJuICdwcm9wZXJ0eSc7XG4gICAgICBjYXNlICdrbm9iOmNzcy1wcm9wZXJ0eS1jaGFuZ2UnOlxuICAgICAgICByZXR1cm4gJ2Nzcy1wcm9wZXJ0eSc7XG4gICAgICBkZWZhdWx0OlxuICAgICAgICByZXR1cm4gJ3Vua25vd24nO1xuICAgIH1cbiAgfVxuXG4gICNnZXRLbm9iVHlwZUZyb21DbGVhckV2ZW50KGV2ZW50OiBFdmVudCk6IHN0cmluZyB7XG4gICAgc3dpdGNoIChldmVudC50eXBlKSB7XG4gICAgICBjYXNlICdrbm9iOmF0dHJpYnV0ZS1jbGVhcic6XG4gICAgICAgIHJldHVybiAnYXR0cmlidXRlJztcbiAgICAgIGNhc2UgJ2tub2I6cHJvcGVydHktY2xlYXInOlxuICAgICAgICByZXR1cm4gJ3Byb3BlcnR5JztcbiAgICAgIGNhc2UgJ2tub2I6Y3NzLXByb3BlcnR5LWNsZWFyJzpcbiAgICAgICAgcmV0dXJuICdjc3MtcHJvcGVydHknO1xuICAgICAgZGVmYXVsdDpcbiAgICAgICAgcmV0dXJuICd1bmtub3duJztcbiAgICB9XG4gIH1cblxuICAjc2V0dXBUcmVlU3RhdGVQZXJzaXN0ZW5jZSgpIHtcbiAgICB0aGlzLiNoYW5kbGVUcmVlRXhwYW5kID0gKGU6IEV2ZW50KSA9PiB7XG4gICAgICBpZiAoKGUudGFyZ2V0IGFzIEVsZW1lbnQpPy50YWdOYW1lICE9PSAnUEYtVjYtVFJFRS1JVEVNJykgcmV0dXJuO1xuXG4gICAgICBjb25zdCBub2RlSWQgPSB0aGlzLiNnZXRUcmVlTm9kZUlkKGUudGFyZ2V0IGFzIEVsZW1lbnQpO1xuICAgICAgY29uc3QgdHJlZVN0YXRlID0gU3RhdGVQZXJzaXN0ZW5jZS5nZXRUcmVlU3RhdGUoKTtcbiAgICAgIGlmICghdHJlZVN0YXRlLmV4cGFuZGVkLmluY2x1ZGVzKG5vZGVJZCkpIHtcbiAgICAgICAgdHJlZVN0YXRlLmV4cGFuZGVkLnB1c2gobm9kZUlkKTtcbiAgICAgICAgU3RhdGVQZXJzaXN0ZW5jZS5zZXRUcmVlU3RhdGUodHJlZVN0YXRlKTtcbiAgICAgIH1cbiAgICB9O1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignZXhwYW5kJywgdGhpcy4jaGFuZGxlVHJlZUV4cGFuZCk7XG5cbiAgICB0aGlzLiNoYW5kbGVUcmVlQ29sbGFwc2UgPSAoZTogRXZlbnQpID0+IHtcbiAgICAgIGlmICgoZS50YXJnZXQgYXMgRWxlbWVudCk/LnRhZ05hbWUgIT09ICdQRi1WNi1UUkVFLUlURU0nKSByZXR1cm47XG5cbiAgICAgIGNvbnN0IG5vZGVJZCA9IHRoaXMuI2dldFRyZWVOb2RlSWQoZS50YXJnZXQgYXMgRWxlbWVudCk7XG4gICAgICBjb25zdCB0cmVlU3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFRyZWVTdGF0ZSgpO1xuICAgICAgY29uc3QgaW5kZXggPSB0cmVlU3RhdGUuZXhwYW5kZWQuaW5kZXhPZihub2RlSWQpO1xuICAgICAgaWYgKGluZGV4ID4gLTEpIHtcbiAgICAgICAgdHJlZVN0YXRlLmV4cGFuZGVkLnNwbGljZShpbmRleCwgMSk7XG4gICAgICAgIFN0YXRlUGVyc2lzdGVuY2Uuc2V0VHJlZVN0YXRlKHRyZWVTdGF0ZSk7XG4gICAgICB9XG4gICAgfTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ2NvbGxhcHNlJywgdGhpcy4jaGFuZGxlVHJlZUNvbGxhcHNlKTtcblxuICAgIHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QgPSAoZTogRXZlbnQpID0+IHtcbiAgICAgIGlmICgoZS50YXJnZXQgYXMgRWxlbWVudCk/LnRhZ05hbWUgIT09ICdQRi1WNi1UUkVFLUlURU0nKSByZXR1cm47XG5cbiAgICAgIGNvbnN0IG5vZGVJZCA9IHRoaXMuI2dldFRyZWVOb2RlSWQoZS50YXJnZXQgYXMgRWxlbWVudCk7XG4gICAgICBTdGF0ZVBlcnNpc3RlbmNlLnVwZGF0ZVRyZWVTdGF0ZSh7IHNlbGVjdGVkOiBub2RlSWQgfSk7XG4gICAgfTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZVRyZWVTZWxlY3QpO1xuXG4gICAgdGhpcy4jYXBwbHlUcmVlU3RhdGUoKTtcbiAgfVxuXG4gICNhcHBseVRyZWVTdGF0ZSgpIHtcbiAgICBjb25zdCB0cmVlU3RhdGUgPSBTdGF0ZVBlcnNpc3RlbmNlLmdldFRyZWVTdGF0ZSgpO1xuXG4gICAgZm9yIChjb25zdCBub2RlSWQgb2YgdHJlZVN0YXRlLmV4cGFuZGVkKSB7XG4gICAgICBjb25zdCB0cmVlSXRlbSA9IHRoaXMuI2ZpbmRUcmVlSXRlbUJ5SWQobm9kZUlkKTtcbiAgICAgIGlmICh0cmVlSXRlbSAmJiAhdHJlZUl0ZW0uaGFzQXR0cmlidXRlKCdleHBhbmRlZCcpKSB7XG4gICAgICAgIHRyZWVJdGVtLnNldEF0dHJpYnV0ZSgnZXhwYW5kZWQnLCAnJyk7XG4gICAgICB9XG4gICAgfVxuXG4gICAgaWYgKHRyZWVTdGF0ZS5zZWxlY3RlZCkge1xuICAgICAgY29uc3QgdHJlZUl0ZW0gPSB0aGlzLiNmaW5kVHJlZUl0ZW1CeUlkKHRyZWVTdGF0ZS5zZWxlY3RlZCk7XG4gICAgICBpZiAodHJlZUl0ZW0gJiYgIXRyZWVJdGVtLmhhc0F0dHJpYnV0ZSgnY3VycmVudCcpKSB7XG4gICAgICAgIHRyZWVJdGVtLnNldEF0dHJpYnV0ZSgnY3VycmVudCcsICcnKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjc2V0dXBTaWRlYmFyU3RhdGVQZXJzaXN0ZW5jZSgpIHtcbiAgICBjb25zdCBwYWdlID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdwZi12Ni1wYWdlJyk7XG5cbiAgICBpZiAoIXBhZ2UpIHJldHVybjtcblxuICAgIHBhZ2UuYWRkRXZlbnRMaXN0ZW5lcignc2lkZWJhci10b2dnbGUnLCAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgICBjb25zdCBjb2xsYXBzZWQgPSAhKGV2ZW50IGFzIGFueSkuZXhwYW5kZWQ7XG5cbiAgICAgIFN0YXRlUGVyc2lzdGVuY2UudXBkYXRlU3RhdGUoe1xuICAgICAgICBzaWRlYmFyOiB7IGNvbGxhcHNlZCB9XG4gICAgICB9KTtcbiAgICB9KTtcbiAgfVxuXG4gICNmaW5kVHJlZUl0ZW1CeUlkKG5vZGVJZDogc3RyaW5nKTogRWxlbWVudCB8IG51bGwge1xuICAgIGNvbnN0IHBhcnRzID0gbm9kZUlkLnNwbGl0KCc6Jyk7XG4gICAgY29uc3QgW3R5cGUsIG1vZHVsZVBhdGgsIHRhZ05hbWUsIG5hbWVdID0gcGFydHM7XG5cbiAgICBsZXQgYXR0clN1ZmZpeCA9ICcnO1xuICAgIGlmICh0YWdOYW1lKSB7XG4gICAgICBhdHRyU3VmZml4ICs9IGBbZGF0YS10YWctbmFtZT1cIiR7Q1NTLmVzY2FwZSh0YWdOYW1lKX1cIl1gO1xuICAgIH1cbiAgICBpZiAobmFtZSkge1xuICAgICAgYXR0clN1ZmZpeCArPSBgW2RhdGEtbmFtZT1cIiR7Q1NTLmVzY2FwZShuYW1lKX1cIl1gO1xuICAgIH1cblxuICAgIGxldCBzZWxlY3RvciA9IGBwZi12Ni10cmVlLWl0ZW1bZGF0YS10eXBlPVwiJHtDU1MuZXNjYXBlKHR5cGUpfVwiXWA7XG4gICAgaWYgKG1vZHVsZVBhdGgpIHtcbiAgICAgIGNvbnN0IGVzY2FwZWRNb2R1bGVQYXRoID0gQ1NTLmVzY2FwZShtb2R1bGVQYXRoKTtcbiAgICAgIGNvbnN0IGVzY2FwZWRUeXBlID0gQ1NTLmVzY2FwZSh0eXBlKTtcbiAgICAgIGNvbnN0IHNlbGVjdG9yMSA9IGBwZi12Ni10cmVlLWl0ZW1bZGF0YS10eXBlPVwiJHtlc2NhcGVkVHlwZX1cIl1bZGF0YS1tb2R1bGUtcGF0aD1cIiR7ZXNjYXBlZE1vZHVsZVBhdGh9XCJdJHthdHRyU3VmZml4fWA7XG4gICAgICBjb25zdCBzZWxlY3RvcjIgPSBgcGYtdjYtdHJlZS1pdGVtW2RhdGEtdHlwZT1cIiR7ZXNjYXBlZFR5cGV9XCJdW2RhdGEtcGF0aD1cIiR7ZXNjYXBlZE1vZHVsZVBhdGh9XCJdJHthdHRyU3VmZml4fWA7XG4gICAgICBzZWxlY3RvciA9IGAke3NlbGVjdG9yMX0sICR7c2VsZWN0b3IyfWA7XG4gICAgfSBlbHNlIHtcbiAgICAgIHNlbGVjdG9yICs9IGF0dHJTdWZmaXg7XG4gICAgfVxuXG4gICAgcmV0dXJuIHRoaXMucXVlcnlTZWxlY3RvcihzZWxlY3Rvcik7XG4gIH1cblxuICAjZ2V0VHJlZU5vZGVJZCh0cmVlSXRlbTogRWxlbWVudCk6IHN0cmluZyB7XG4gICAgY29uc3QgdHlwZSA9IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS10eXBlJyk7XG4gICAgY29uc3QgbW9kdWxlUGF0aCA9IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS1tb2R1bGUtcGF0aCcpIHx8IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS1wYXRoJyk7XG4gICAgY29uc3QgdGFnTmFtZSA9IHRyZWVJdGVtLmdldEF0dHJpYnV0ZSgnZGF0YS10YWctbmFtZScpO1xuICAgIGNvbnN0IG5hbWUgPSB0cmVlSXRlbS5nZXRBdHRyaWJ1dGUoJ2RhdGEtbmFtZScpO1xuICAgIGNvbnN0IGNhdGVnb3J5ID0gdHJlZUl0ZW0uZ2V0QXR0cmlidXRlKCdkYXRhLWNhdGVnb3J5Jyk7XG5cbiAgICBjb25zdCBwYXJ0cyA9IFt0eXBlXTtcbiAgICBpZiAobW9kdWxlUGF0aCkgcGFydHMucHVzaChtb2R1bGVQYXRoKTtcbiAgICBpZiAodGFnTmFtZSkgcGFydHMucHVzaCh0YWdOYW1lKTtcbiAgICBpZiAoY2F0ZWdvcnkpIHtcbiAgICAgIHBhcnRzLnB1c2goY2F0ZWdvcnkpO1xuICAgIH0gZWxzZSBpZiAobmFtZSkge1xuICAgICAgcGFydHMucHVzaChuYW1lKTtcbiAgICB9XG5cbiAgICByZXR1cm4gcGFydHMuam9pbignOicpO1xuICB9XG5cbiAgLy8gRXZlbnQgRGlzY292ZXJ5ICYgQ2FwdHVyZSBNZXRob2RzXG5cbiAgYXN5bmMgI2Rpc2NvdmVyRWxlbWVudEV2ZW50cygpOiBQcm9taXNlPE1hcDxzdHJpbmcsIEV2ZW50SW5mbz4+IHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgcmVzcG9uc2UgPSBhd2FpdCBmZXRjaCgnL2N1c3RvbS1lbGVtZW50cy5qc29uJyk7XG4gICAgICBpZiAoIXJlc3BvbnNlLm9rKSB7XG4gICAgICAgIGNvbnNvbGUud2FybignW2NlbS1zZXJ2ZS1jaHJvbWVdIE5vIG1hbmlmZXN0IGF2YWlsYWJsZSBmb3IgZXZlbnQgZGlzY292ZXJ5Jyk7XG4gICAgICAgIHJldHVybiBuZXcgTWFwKCk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IG1hbmlmZXN0ID0gYXdhaXQgcmVzcG9uc2UuanNvbigpIGFzIE1hbmlmZXN0O1xuICAgICAgdGhpcy4jbWFuaWZlc3QgPSBtYW5pZmVzdDtcblxuICAgICAgY29uc3QgZXZlbnRNYXAgPSBuZXcgTWFwPHN0cmluZywgRXZlbnRJbmZvPigpO1xuXG4gICAgICBmb3IgKGNvbnN0IG1vZHVsZSBvZiBtYW5pZmVzdC5tb2R1bGVzIHx8IFtdKSB7XG4gICAgICAgIGZvciAoY29uc3QgZGVjbGFyYXRpb24gb2YgbW9kdWxlLmRlY2xhcmF0aW9ucyB8fCBbXSkge1xuICAgICAgICAgIGlmIChkZWNsYXJhdGlvbi5jdXN0b21FbGVtZW50ICYmIGRlY2xhcmF0aW9uLnRhZ05hbWUpIHtcbiAgICAgICAgICAgIGNvbnN0IHRhZ05hbWUgPSBkZWNsYXJhdGlvbi50YWdOYW1lO1xuICAgICAgICAgICAgY29uc3QgZXZlbnRzID0gZGVjbGFyYXRpb24uZXZlbnRzIHx8IFtdO1xuXG4gICAgICAgICAgICBpZiAoZXZlbnRzLmxlbmd0aCA+IDApIHtcbiAgICAgICAgICAgICAgY29uc3QgZXZlbnROYW1lcyA9IG5ldyBTZXQoZXZlbnRzLm1hcChlID0+IGUubmFtZSkpO1xuICAgICAgICAgICAgICBldmVudE1hcC5zZXQodGFnTmFtZSwge1xuICAgICAgICAgICAgICAgIGV2ZW50TmFtZXMsXG4gICAgICAgICAgICAgICAgZXZlbnRzOiBldmVudHNcbiAgICAgICAgICAgICAgfSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIHJldHVybiBldmVudE1hcDtcbiAgICB9IGNhdGNoIChlcnJvcikge1xuICAgICAgY29uc29sZS53YXJuKCdbY2VtLXNlcnZlLWNocm9tZV0gRXJyb3IgbG9hZGluZyBtYW5pZmVzdCBmb3IgZXZlbnQgZGlzY292ZXJ5OicsIGVycm9yKTtcbiAgICAgIHJldHVybiBuZXcgTWFwKCk7XG4gICAgfVxuICB9XG5cbiAgYXN5bmMgI3NldHVwRXZlbnRDYXB0dXJlKCkge1xuICAgIHRoaXMuI2VsZW1lbnRFdmVudE1hcCA9IGF3YWl0IHRoaXMuI2Rpc2NvdmVyRWxlbWVudEV2ZW50cygpO1xuXG4gICAgaWYgKHRoaXMuI2VsZW1lbnRFdmVudE1hcC5zaXplID09PSAwKSByZXR1cm47XG5cbiAgICB0aGlzLiNhdHRhY2hFdmVudExpc3RlbmVycygpO1xuICAgIHRoaXMuI3VwZGF0ZUV2ZW50RmlsdGVycygpO1xuICAgIHRoaXMuI29ic2VydmVEZW1vTXV0YXRpb25zKCk7XG4gIH1cblxuICAjYXR0YWNoRXZlbnRMaXN0ZW5lcnMoKSB7XG4gICAgY29uc3QgZGVtbyA9IHRoaXMuZGVtbztcbiAgICBpZiAoIWRlbW8pIHJldHVybjtcblxuICAgIGNvbnN0IHJvb3QgPSBkZW1vLnNoYWRvd1Jvb3QgPz8gZGVtbztcblxuICAgIGZvciAoY29uc3QgW3RhZ05hbWUsIGV2ZW50SW5mb10gb2YgdGhpcy4jZWxlbWVudEV2ZW50TWFwISkge1xuICAgICAgY29uc3QgZWxlbWVudHMgPSByb290LnF1ZXJ5U2VsZWN0b3JBbGwodGFnTmFtZSk7XG5cbiAgICAgIGZvciAoY29uc3QgZWxlbWVudCBvZiBlbGVtZW50cykge1xuICAgICAgICBmb3IgKGNvbnN0IGV2ZW50TmFtZSBvZiBldmVudEluZm8uZXZlbnROYW1lcykge1xuICAgICAgICAgIGVsZW1lbnQuYWRkRXZlbnRMaXN0ZW5lcihldmVudE5hbWUsIHRoaXMuI2hhbmRsZUVsZW1lbnRFdmVudCwgeyBjYXB0dXJlOiB0cnVlIH0pO1xuICAgICAgICB9XG4gICAgICAgIChlbGVtZW50IGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LmNlbUV2ZW50c0F0dGFjaGVkID0gJ3RydWUnO1xuICAgICAgICB0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuYWRkKHRhZ05hbWUpO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gICNvYnNlcnZlRGVtb011dGF0aW9ucygpIHtcbiAgICBjb25zdCBkZW1vID0gdGhpcy5kZW1vO1xuICAgIGlmICghZGVtbykgcmV0dXJuO1xuXG4gICAgY29uc3Qgcm9vdCA9IGRlbW8uc2hhZG93Um9vdCA/PyBkZW1vO1xuXG4gICAgdGhpcy4jb2JzZXJ2ZXIub2JzZXJ2ZShyb290LCB7XG4gICAgICBjaGlsZExpc3Q6IHRydWUsXG4gICAgICBzdWJ0cmVlOiB0cnVlXG4gICAgfSk7XG4gIH1cblxuICAjaGFuZGxlRWxlbWVudEV2ZW50ID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IGVsZW1lbnQgPSBldmVudC5jdXJyZW50VGFyZ2V0O1xuICAgIGlmICghKGVsZW1lbnQgaW5zdGFuY2VvZiBIVE1MRWxlbWVudCkpIHJldHVybjtcblxuICAgIGNvbnN0IHRhZ05hbWUgPSBlbGVtZW50LnRhZ05hbWUudG9Mb3dlckNhc2UoKTtcbiAgICBjb25zdCBldmVudEluZm8gPSB0aGlzLiNlbGVtZW50RXZlbnRNYXA/LmdldCh0YWdOYW1lKTtcblxuICAgIGlmICghZXZlbnRJbmZvIHx8ICFldmVudEluZm8uZXZlbnROYW1lcy5oYXMoZXZlbnQudHlwZSkpIHJldHVybjtcblxuICAgIHRoaXMuI2Rpc2NvdmVyZWRFbGVtZW50cy5hZGQodGFnTmFtZSk7XG4gICAgdGhpcy4jY2FwdHVyZUV2ZW50KGV2ZW50LCBlbGVtZW50LCB0YWdOYW1lLCBldmVudEluZm8pO1xuICB9O1xuXG4gICNnZXRFdmVudERvY3VtZW50YXRpb24obWFuaWZlc3RFdmVudDogRXZlbnRJbmZvWydldmVudHMnXVswXSB8IHVuZGVmaW5lZCkge1xuICAgIGlmICghbWFuaWZlc3RFdmVudCkge1xuICAgICAgcmV0dXJuIHsgc3VtbWFyeTogbnVsbCwgZGVzY3JpcHRpb246IG51bGwgfTtcbiAgICB9XG5cbiAgICBsZXQgc3VtbWFyeSA9IG1hbmlmZXN0RXZlbnQuc3VtbWFyeSB8fCBudWxsO1xuICAgIGxldCBkZXNjcmlwdGlvbiA9IG1hbmlmZXN0RXZlbnQuZGVzY3JpcHRpb24gfHwgbnVsbDtcblxuICAgIGlmIChtYW5pZmVzdEV2ZW50LnR5cGU/LnRleHQgJiYgdGhpcy4jbWFuaWZlc3QpIHtcbiAgICAgIGNvbnN0IHR5cGVOYW1lID0gbWFuaWZlc3RFdmVudC50eXBlLnRleHQ7XG4gICAgICBjb25zdCB0eXBlRGVjbGFyYXRpb24gPSB0aGlzLiNmaW5kVHlwZURlY2xhcmF0aW9uKHR5cGVOYW1lKTtcblxuICAgICAgaWYgKHR5cGVEZWNsYXJhdGlvbikge1xuICAgICAgICBpZiAoIXN1bW1hcnkgJiYgdHlwZURlY2xhcmF0aW9uLnN1bW1hcnkpIHtcbiAgICAgICAgICBzdW1tYXJ5ID0gdHlwZURlY2xhcmF0aW9uLnN1bW1hcnk7XG4gICAgICAgIH0gZWxzZSBpZiAodHlwZURlY2xhcmF0aW9uLnN1bW1hcnkgJiYgdHlwZURlY2xhcmF0aW9uLnN1bW1hcnkgIT09IHN1bW1hcnkpIHtcbiAgICAgICAgICBzdW1tYXJ5ID0gc3VtbWFyeSA/IGAke3N1bW1hcnl9XFxuXFxuRnJvbSAke3R5cGVOYW1lfTogJHt0eXBlRGVjbGFyYXRpb24uc3VtbWFyeX1gIDogdHlwZURlY2xhcmF0aW9uLnN1bW1hcnk7XG4gICAgICAgIH1cblxuICAgICAgICBpZiAoIWRlc2NyaXB0aW9uICYmIHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbikge1xuICAgICAgICAgIGRlc2NyaXB0aW9uID0gdHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9uO1xuICAgICAgICB9IGVsc2UgaWYgKHR5cGVEZWNsYXJhdGlvbi5kZXNjcmlwdGlvbiAmJiB0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb24gIT09IGRlc2NyaXB0aW9uKSB7XG4gICAgICAgICAgZGVzY3JpcHRpb24gPSBkZXNjcmlwdGlvbiA/IGAke2Rlc2NyaXB0aW9ufVxcblxcbiR7dHlwZURlY2xhcmF0aW9uLmRlc2NyaXB0aW9ufWAgOiB0eXBlRGVjbGFyYXRpb24uZGVzY3JpcHRpb247XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICByZXR1cm4geyBzdW1tYXJ5LCBkZXNjcmlwdGlvbiB9O1xuICB9XG5cbiAgI2ZpbmRUeXBlRGVjbGFyYXRpb24odHlwZU5hbWU6IHN0cmluZykge1xuICAgIGlmICghdGhpcy4jbWFuaWZlc3QpIHJldHVybiBudWxsO1xuXG4gICAgZm9yIChjb25zdCBtb2R1bGUgb2YgdGhpcy4jbWFuaWZlc3QubW9kdWxlcyB8fCBbXSkge1xuICAgICAgZm9yIChjb25zdCBkZWNsYXJhdGlvbiBvZiBtb2R1bGUuZGVjbGFyYXRpb25zIHx8IFtdKSB7XG4gICAgICAgIGlmIChkZWNsYXJhdGlvbi5uYW1lID09PSB0eXBlTmFtZSAmJlxuICAgICAgICAgICAgKGRlY2xhcmF0aW9uLmtpbmQgPT09ICdjbGFzcycgfHwgZGVjbGFyYXRpb24ua2luZCA9PT0gJ2ludGVyZmFjZScpKSB7XG4gICAgICAgICAgcmV0dXJuIGRlY2xhcmF0aW9uIGFzIHsgc3VtbWFyeT86IHN0cmluZzsgZGVzY3JpcHRpb24/OiBzdHJpbmcgfTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiBudWxsO1xuICB9XG5cbiAgI2NhcHR1cmVFdmVudChldmVudDogRXZlbnQsIHRhcmdldDogSFRNTEVsZW1lbnQsIHRhZ05hbWU6IHN0cmluZywgZXZlbnRJbmZvOiBFdmVudEluZm8pIHtcbiAgICBjb25zdCBtYW5pZmVzdEV2ZW50ID0gZXZlbnRJbmZvLmV2ZW50cy5maW5kKGUgPT4gZS5uYW1lID09PSBldmVudC50eXBlKTtcblxuICAgIGNvbnN0IGV2ZW50RG9jcyA9IHRoaXMuI2dldEV2ZW50RG9jdW1lbnRhdGlvbihtYW5pZmVzdEV2ZW50KTtcblxuICAgIGNvbnN0IGN1c3RvbVByb3BlcnRpZXMgPSB0aGlzLiNleHRyYWN0RXZlbnRQcm9wZXJ0aWVzKGV2ZW50KTtcblxuICAgIGNvbnN0IGV2ZW50UmVjb3JkOiBFdmVudFJlY29yZCA9IHtcbiAgICAgIGlkOiBgJHtEYXRlLm5vdygpfS0ke01hdGgucmFuZG9tKCl9YCxcbiAgICAgIHRpbWVzdGFtcDogbmV3IERhdGUoKSxcbiAgICAgIGV2ZW50TmFtZTogZXZlbnQudHlwZSxcbiAgICAgIHRhZ05hbWU6IHRhZ05hbWUsXG4gICAgICBlbGVtZW50SWQ6IHRhcmdldC5pZCB8fCBudWxsLFxuICAgICAgZWxlbWVudENsYXNzOiB0YXJnZXQuY2xhc3NOYW1lIHx8IG51bGwsXG4gICAgICBjdXN0b21Qcm9wZXJ0aWVzOiBjdXN0b21Qcm9wZXJ0aWVzLFxuICAgICAgbWFuaWZlc3RUeXBlOiBtYW5pZmVzdEV2ZW50Py50eXBlPy50ZXh0IHx8IG51bGwsXG4gICAgICBzdW1tYXJ5OiBldmVudERvY3Muc3VtbWFyeSxcbiAgICAgIGRlc2NyaXB0aW9uOiBldmVudERvY3MuZGVzY3JpcHRpb24sXG4gICAgICBidWJibGVzOiBldmVudC5idWJibGVzLFxuICAgICAgY29tcG9zZWQ6IGV2ZW50LmNvbXBvc2VkLFxuICAgICAgY2FuY2VsYWJsZTogZXZlbnQuY2FuY2VsYWJsZSxcbiAgICAgIGRlZmF1bHRQcmV2ZW50ZWQ6IGV2ZW50LmRlZmF1bHRQcmV2ZW50ZWRcbiAgICB9O1xuXG4gICAgdGhpcy4jY2FwdHVyZWRFdmVudHMucHVzaChldmVudFJlY29yZCk7XG5cbiAgICBpZiAodGhpcy4jY2FwdHVyZWRFdmVudHMubGVuZ3RoID4gdGhpcy4jbWF4Q2FwdHVyZWRFdmVudHMpIHtcbiAgICAgIHRoaXMuI2NhcHR1cmVkRXZlbnRzLnNoaWZ0KCk7XG4gICAgfVxuXG4gICAgdGhpcy4jcmVuZGVyRXZlbnQoZXZlbnRSZWNvcmQpO1xuICB9XG5cbiAgI2V4dHJhY3RFdmVudFByb3BlcnRpZXMoZXZlbnQ6IEV2ZW50KTogUmVjb3JkPHN0cmluZywgdW5rbm93bj4ge1xuICAgIGNvbnN0IHByb3BlcnRpZXM6IFJlY29yZDxzdHJpbmcsIHVua25vd24+ID0ge307XG4gICAgY29uc3QgZXZlbnRQcm90b3R5cGVLZXlzID0gbmV3IFNldChPYmplY3QuZ2V0T3duUHJvcGVydHlOYW1lcyhFdmVudC5wcm90b3R5cGUpKTtcblxuICAgIGNvbnN0IHNlcmlhbGl6ZVZhbHVlID0gKHZhbHVlOiB1bmtub3duKTogdW5rbm93biA9PiB7XG4gICAgICB0cnkge1xuICAgICAgICByZXR1cm4gSlNPTi5wYXJzZShKU09OLnN0cmluZ2lmeSh2YWx1ZSkpO1xuICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICB0cnkge1xuICAgICAgICAgIHJldHVybiBTdHJpbmcodmFsdWUpO1xuICAgICAgICB9IGNhdGNoIChzdHJpbmdFcnIpIHtcbiAgICAgICAgICByZXR1cm4gJ1tOb3Qgc2VyaWFsaXphYmxlXSc7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9O1xuXG4gICAgaWYgKGV2ZW50IGluc3RhbmNlb2YgQ3VzdG9tRXZlbnQgJiYgZXZlbnQuZGV0YWlsICE9PSB1bmRlZmluZWQpIHtcbiAgICAgIHByb3BlcnRpZXMuZGV0YWlsID0gc2VyaWFsaXplVmFsdWUoZXZlbnQuZGV0YWlsKTtcbiAgICB9XG5cbiAgICBmb3IgKGNvbnN0IGtleSBvZiBPYmplY3QuZ2V0T3duUHJvcGVydHlOYW1lcyhldmVudCkpIHtcbiAgICAgIGlmICghZXZlbnRQcm90b3R5cGVLZXlzLmhhcyhrZXkpICYmICFrZXkuc3RhcnRzV2l0aCgnXycpICYmICFwcm9wZXJ0aWVzLmhhc093blByb3BlcnR5KGtleSkpIHtcbiAgICAgICAgcHJvcGVydGllc1trZXldID0gc2VyaWFsaXplVmFsdWUoKGV2ZW50IGFzIGFueSlba2V5XSk7XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIHByb3BlcnRpZXM7XG4gIH1cblxuICAjcmVuZGVyRXZlbnQoZXZlbnRSZWNvcmQ6IEV2ZW50UmVjb3JkKSB7XG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIGNvbnN0IGZyYWdtZW50ID0gQ2VtU2VydmVDaHJvbWUuI2V2ZW50RW50cnlUZW1wbGF0ZS5jb250ZW50LmNsb25lTm9kZSh0cnVlKSBhcyBEb2N1bWVudEZyYWdtZW50O1xuXG4gICAgY29uc3QgdGltZSA9IGV2ZW50UmVjb3JkLnRpbWVzdGFtcC50b0xvY2FsZVRpbWVTdHJpbmcoKTtcblxuICAgIGNvbnN0IGNvbnRhaW5lciA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwiY29udGFpbmVyXCJdJykgYXMgSFRNTEVsZW1lbnQ7XG4gICAgY29udGFpbmVyLmRhdGFzZXQuZXZlbnRJZCA9IGV2ZW50UmVjb3JkLmlkO1xuICAgIGNvbnRhaW5lci5kYXRhc2V0LmV2ZW50VHlwZSA9IGV2ZW50UmVjb3JkLmV2ZW50TmFtZTtcbiAgICBjb250YWluZXIuZGF0YXNldC5lbGVtZW50VHlwZSA9IGV2ZW50UmVjb3JkLnRhZ05hbWU7XG5cbiAgICBjb25zdCB0ZXh0TWF0Y2ggPSB0aGlzLiNldmVudE1hdGNoZXNUZXh0RmlsdGVyKGV2ZW50UmVjb3JkKTtcbiAgICBjb25zdCB0eXBlTWF0Y2ggPSB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLnNpemUgPT09IDAgfHwgdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5oYXMoZXZlbnRSZWNvcmQuZXZlbnROYW1lKTtcbiAgICBjb25zdCBlbGVtZW50TWF0Y2ggPSB0aGlzLiNlbGVtZW50RmlsdGVycy5zaXplID09PSAwIHx8IHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmhhcyhldmVudFJlY29yZC50YWdOYW1lKTtcblxuICAgIGlmICghKHRleHRNYXRjaCAmJiB0eXBlTWF0Y2ggJiYgZWxlbWVudE1hdGNoKSkge1xuICAgICAgY29udGFpbmVyLnNldEF0dHJpYnV0ZSgnaGlkZGVuJywgJycpO1xuICAgIH1cblxuICAgIGNvbnN0IGxhYmVsID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignW2RhdGEtZmllbGQ9XCJsYWJlbFwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgIGxhYmVsLnRleHRDb250ZW50ID0gZXZlbnRSZWNvcmQuZXZlbnROYW1lO1xuICAgIGxhYmVsLnNldEF0dHJpYnV0ZSgnc3RhdHVzJywgJ2luZm8nKTtcblxuICAgIGNvbnN0IHRpbWVFbCA9IGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3IoJ1tkYXRhLWZpZWxkPVwidGltZVwiXScpIGFzIEhUTUxFbGVtZW50O1xuICAgIHRpbWVFbC5zZXRBdHRyaWJ1dGUoJ2RhdGV0aW1lJywgZXZlbnRSZWNvcmQudGltZXN0YW1wLnRvSVNPU3RyaW5nKCkpO1xuICAgIHRpbWVFbC50ZXh0Q29udGVudCA9IHRpbWU7XG5cbiAgICBjb25zdCBlbGVtZW50RWwgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbZGF0YS1maWVsZD1cImVsZW1lbnRcIl0nKSBhcyBIVE1MRWxlbWVudDtcbiAgICBsZXQgZWxlbWVudFRleHQgPSBgPCR7ZXZlbnRSZWNvcmQudGFnTmFtZX0+YDtcbiAgICBpZiAoZXZlbnRSZWNvcmQuZWxlbWVudElkKSB7XG4gICAgICBlbGVtZW50VGV4dCArPSBgIyR7ZXZlbnRSZWNvcmQuZWxlbWVudElkfWA7XG4gICAgfVxuICAgIGVsZW1lbnRFbC50ZXh0Q29udGVudCA9IGVsZW1lbnRUZXh0O1xuXG4gICAgdGhpcy4jZXZlbnRMaXN0LmFwcGVuZChmcmFnbWVudCk7XG5cbiAgICBpZiAoIXRoaXMuI3NlbGVjdGVkRXZlbnRJZCkge1xuICAgICAgdGhpcy4jc2VsZWN0RXZlbnQoZXZlbnRSZWNvcmQuaWQpO1xuICAgIH1cblxuICAgIGlmICh0aGlzLiNkcmF3ZXJPcGVuICYmIHRoaXMuI2lzRXZlbnRzVGFiQWN0aXZlKCkpIHtcbiAgICAgIHRoaXMuI3Njcm9sbEV2ZW50c1RvQm90dG9tKCk7XG4gICAgfVxuICB9XG5cbiAgI3NlbGVjdEV2ZW50KGV2ZW50SWQ6IHN0cmluZykge1xuICAgIGNvbnN0IGV2ZW50UmVjb3JkID0gdGhpcy4jZ2V0RXZlbnRSZWNvcmRCeUlkKGV2ZW50SWQpO1xuICAgIGlmICghZXZlbnRSZWNvcmQpIHJldHVybjtcblxuICAgIHRoaXMuI3NlbGVjdGVkRXZlbnRJZCA9IGV2ZW50SWQ7XG5cbiAgICBjb25zdCBhbGxJdGVtcyA9IHRoaXMuI2V2ZW50TGlzdD8ucXVlcnlTZWxlY3RvckFsbCgnLmV2ZW50LWxpc3QtaXRlbScpO1xuICAgIGFsbEl0ZW1zPy5mb3JFYWNoKGl0ZW0gPT4ge1xuICAgICAgaWYgKChpdGVtIGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LmV2ZW50SWQgPT09IGV2ZW50SWQpIHtcbiAgICAgICAgaXRlbS5jbGFzc0xpc3QuYWRkKCdzZWxlY3RlZCcpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgnYXJpYS1zZWxlY3RlZCcsICd0cnVlJyk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBpdGVtLmNsYXNzTGlzdC5yZW1vdmUoJ3NlbGVjdGVkJyk7XG4gICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCdhcmlhLXNlbGVjdGVkJywgJ2ZhbHNlJyk7XG4gICAgICB9XG4gICAgfSk7XG5cbiAgICBpZiAodGhpcy4jZXZlbnREZXRhaWxIZWFkZXIpIHtcbiAgICAgIHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyLmlubmVySFRNTCA9ICcnO1xuXG4gICAgICBjb25zdCBoZWFkZXJDb250ZW50ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG4gICAgICBoZWFkZXJDb250ZW50LmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtaGVhZGVyLWNvbnRlbnQnO1xuXG4gICAgICBjb25zdCBldmVudE5hbWUgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdoMycpO1xuICAgICAgZXZlbnROYW1lLnRleHRDb250ZW50ID0gZXZlbnRSZWNvcmQuZXZlbnROYW1lO1xuICAgICAgZXZlbnROYW1lLmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtbmFtZSc7XG4gICAgICBoZWFkZXJDb250ZW50LmFwcGVuZENoaWxkKGV2ZW50TmFtZSk7XG5cbiAgICAgIGlmIChldmVudFJlY29yZC5zdW1tYXJ5KSB7XG4gICAgICAgIGNvbnN0IHN1bW1hcnkgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwJyk7XG4gICAgICAgIHN1bW1hcnkudGV4dENvbnRlbnQgPSBldmVudFJlY29yZC5zdW1tYXJ5O1xuICAgICAgICBzdW1tYXJ5LmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtc3VtbWFyeSc7XG4gICAgICAgIGhlYWRlckNvbnRlbnQuYXBwZW5kQ2hpbGQoc3VtbWFyeSk7XG4gICAgICB9XG5cbiAgICAgIGlmIChldmVudFJlY29yZC5kZXNjcmlwdGlvbikge1xuICAgICAgICBjb25zdCBkZXNjcmlwdGlvbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3AnKTtcbiAgICAgICAgZGVzY3JpcHRpb24udGV4dENvbnRlbnQgPSBldmVudFJlY29yZC5kZXNjcmlwdGlvbjtcbiAgICAgICAgZGVzY3JpcHRpb24uY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1kZXNjcmlwdGlvbic7XG4gICAgICAgIGhlYWRlckNvbnRlbnQuYXBwZW5kQ2hpbGQoZGVzY3JpcHRpb24pO1xuICAgICAgfVxuXG4gICAgICBjb25zdCBtZXRhID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG4gICAgICBtZXRhLmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtbWV0YSc7XG5cbiAgICAgIGNvbnN0IHRpbWVFbCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RpbWUnKTtcbiAgICAgIHRpbWVFbC5zZXRBdHRyaWJ1dGUoJ2RhdGV0aW1lJywgZXZlbnRSZWNvcmQudGltZXN0YW1wLnRvSVNPU3RyaW5nKCkpO1xuICAgICAgdGltZUVsLnRleHRDb250ZW50ID0gZXZlbnRSZWNvcmQudGltZXN0YW1wLnRvTG9jYWxlVGltZVN0cmluZygpO1xuICAgICAgdGltZUVsLmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtdGltZSc7XG5cbiAgICAgIGNvbnN0IGVsZW1lbnQgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICBsZXQgZWxlbWVudFRleHQgPSBgPCR7ZXZlbnRSZWNvcmQudGFnTmFtZX0+YDtcbiAgICAgIGlmIChldmVudFJlY29yZC5lbGVtZW50SWQpIHtcbiAgICAgICAgZWxlbWVudFRleHQgKz0gYCMke2V2ZW50UmVjb3JkLmVsZW1lbnRJZH1gO1xuICAgICAgfVxuICAgICAgZWxlbWVudC50ZXh0Q29udGVudCA9IGVsZW1lbnRUZXh0O1xuICAgICAgZWxlbWVudC5jbGFzc05hbWUgPSAnZXZlbnQtZGV0YWlsLWVsZW1lbnQnO1xuXG4gICAgICBtZXRhLmFwcGVuZENoaWxkKHRpbWVFbCk7XG4gICAgICBtZXRhLmFwcGVuZENoaWxkKGVsZW1lbnQpO1xuXG4gICAgICBoZWFkZXJDb250ZW50LmFwcGVuZENoaWxkKG1ldGEpO1xuXG4gICAgICB0aGlzLiNldmVudERldGFpbEhlYWRlci5hcHBlbmRDaGlsZChoZWFkZXJDb250ZW50KTtcbiAgICB9XG5cbiAgICBpZiAodGhpcy4jZXZlbnREZXRhaWxCb2R5KSB7XG4gICAgICB0aGlzLiNldmVudERldGFpbEJvZHkuaW5uZXJIVE1MID0gJyc7XG5cbiAgICAgIGNvbnN0IHByb3BlcnRpZXNIZWFkaW5nID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnaDQnKTtcbiAgICAgIHByb3BlcnRpZXNIZWFkaW5nLnRleHRDb250ZW50ID0gJ1Byb3BlcnRpZXMnO1xuICAgICAgcHJvcGVydGllc0hlYWRpbmcuY2xhc3NOYW1lID0gJ2V2ZW50LWRldGFpbC1wcm9wZXJ0aWVzLWhlYWRpbmcnO1xuXG4gICAgICBjb25zdCBwcm9wZXJ0aWVzQ29udGFpbmVyID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG4gICAgICBwcm9wZXJ0aWVzQ29udGFpbmVyLmNsYXNzTmFtZSA9ICdldmVudC1kZXRhaWwtcHJvcGVydGllcyc7XG5cbiAgICAgIGNvbnN0IGV2ZW50UHJvcGVydGllcyA9IHRoaXMuI2J1aWxkUHJvcGVydGllc0ZvckRpc3BsYXkoZXZlbnRSZWNvcmQpO1xuICAgICAgaWYgKE9iamVjdC5rZXlzKGV2ZW50UHJvcGVydGllcykubGVuZ3RoID4gMCkge1xuICAgICAgICBwcm9wZXJ0aWVzQ29udGFpbmVyLmFwcGVuZENoaWxkKHRoaXMuI2J1aWxkUHJvcGVydHlUcmVlKGV2ZW50UHJvcGVydGllcykpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgcHJvcGVydGllc0NvbnRhaW5lci50ZXh0Q29udGVudCA9ICdObyBwcm9wZXJ0aWVzIHRvIGRpc3BsYXknO1xuICAgICAgfVxuXG4gICAgICB0aGlzLiNldmVudERldGFpbEJvZHkuYXBwZW5kQ2hpbGQocHJvcGVydGllc0hlYWRpbmcpO1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5LmFwcGVuZENoaWxkKHByb3BlcnRpZXNDb250YWluZXIpO1xuICAgIH1cbiAgfVxuXG4gICNidWlsZFByb3BlcnRpZXNGb3JEaXNwbGF5KGV2ZW50UmVjb3JkOiBFdmVudFJlY29yZCk6IFJlY29yZDxzdHJpbmcsIHVua25vd24+IHtcbiAgICBjb25zdCBwcm9wZXJ0aWVzOiBSZWNvcmQ8c3RyaW5nLCB1bmtub3duPiA9IHt9O1xuXG4gICAgaWYgKGV2ZW50UmVjb3JkLmN1c3RvbVByb3BlcnRpZXMpIHtcbiAgICAgIE9iamVjdC5hc3NpZ24ocHJvcGVydGllcywgZXZlbnRSZWNvcmQuY3VzdG9tUHJvcGVydGllcyk7XG4gICAgfVxuXG4gICAgcHJvcGVydGllcy5idWJibGVzID0gZXZlbnRSZWNvcmQuYnViYmxlcztcbiAgICBwcm9wZXJ0aWVzLmNhbmNlbGFibGUgPSBldmVudFJlY29yZC5jYW5jZWxhYmxlO1xuICAgIHByb3BlcnRpZXMuZGVmYXVsdFByZXZlbnRlZCA9IGV2ZW50UmVjb3JkLmRlZmF1bHRQcmV2ZW50ZWQ7XG4gICAgcHJvcGVydGllcy5jb21wb3NlZCA9IGV2ZW50UmVjb3JkLmNvbXBvc2VkO1xuXG4gICAgaWYgKGV2ZW50UmVjb3JkLm1hbmlmZXN0VHlwZSkge1xuICAgICAgcHJvcGVydGllcy50eXBlID0gZXZlbnRSZWNvcmQubWFuaWZlc3RUeXBlO1xuICAgIH1cblxuICAgIHJldHVybiBwcm9wZXJ0aWVzO1xuICB9XG5cbiAgI2J1aWxkUHJvcGVydHlUcmVlKG9iajogUmVjb3JkPHN0cmluZywgdW5rbm93bj4sIGRlcHRoID0gMCk6IEhUTUxVTGlzdEVsZW1lbnQge1xuICAgIGNvbnN0IHVsID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgndWwnKTtcbiAgICB1bC5jbGFzc05hbWUgPSAnZXZlbnQtcHJvcGVydHktdHJlZSc7XG4gICAgaWYgKGRlcHRoID4gMCkge1xuICAgICAgdWwuY2xhc3NMaXN0LmFkZCgnbmVzdGVkJyk7XG4gICAgfVxuXG4gICAgZm9yIChjb25zdCBba2V5LCB2YWx1ZV0gb2YgT2JqZWN0LmVudHJpZXMob2JqKSkge1xuICAgICAgY29uc3QgbGkgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdsaScpO1xuICAgICAgbGkuY2xhc3NOYW1lID0gJ3Byb3BlcnR5LWl0ZW0nO1xuXG4gICAgICBjb25zdCBrZXlTcGFuID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnc3BhbicpO1xuICAgICAga2V5U3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHkta2V5JztcbiAgICAgIGtleVNwYW4udGV4dENvbnRlbnQgPSBrZXk7XG5cbiAgICAgIGNvbnN0IGNvbG9uU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgIGNvbG9uU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktY29sb24nO1xuICAgICAgY29sb25TcGFuLnRleHRDb250ZW50ID0gJzogJztcblxuICAgICAgbGkuYXBwZW5kQ2hpbGQoa2V5U3Bhbik7XG4gICAgICBsaS5hcHBlbmRDaGlsZChjb2xvblNwYW4pO1xuXG4gICAgICBpZiAodmFsdWUgPT09IG51bGwgfHwgdmFsdWUgPT09IHVuZGVmaW5lZCkge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUgbnVsbCc7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IFN0cmluZyh2YWx1ZSk7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiB2YWx1ZSA9PT0gJ2Jvb2xlYW4nKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBib29sZWFuJztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH0gZWxzZSBpZiAodHlwZW9mIHZhbHVlID09PSAnbnVtYmVyJykge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUgbnVtYmVyJztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gU3RyaW5nKHZhbHVlKTtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcbiAgICAgIH0gZWxzZSBpZiAodHlwZW9mIHZhbHVlID09PSAnc3RyaW5nJykge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUgc3RyaW5nJztcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0gYFwiJHt2YWx1ZX1cImA7XG4gICAgICAgIGxpLmFwcGVuZENoaWxkKHZhbHVlU3Bhbik7XG4gICAgICB9IGVsc2UgaWYgKEFycmF5LmlzQXJyYXkodmFsdWUpKSB7XG4gICAgICAgIGNvbnN0IHZhbHVlU3BhbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcbiAgICAgICAgdmFsdWVTcGFuLmNsYXNzTmFtZSA9ICdwcm9wZXJ0eS12YWx1ZSBhcnJheSc7XG4gICAgICAgIHZhbHVlU3Bhbi50ZXh0Q29udGVudCA9IGBBcnJheSgke3ZhbHVlLmxlbmd0aH0pYDtcbiAgICAgICAgbGkuYXBwZW5kQ2hpbGQodmFsdWVTcGFuKTtcblxuICAgICAgICBpZiAodmFsdWUubGVuZ3RoID4gMCAmJiBkZXB0aCA8IDMpIHtcbiAgICAgICAgICBjb25zdCBuZXN0ZWRPYmo6IFJlY29yZDxzdHJpbmcsIHVua25vd24+ID0ge307XG4gICAgICAgICAgdmFsdWUuZm9yRWFjaCgoaXRlbSwgaW5kZXgpID0+IHtcbiAgICAgICAgICAgIG5lc3RlZE9ialtpbmRleF0gPSBpdGVtO1xuICAgICAgICAgIH0pO1xuICAgICAgICAgIGxpLmFwcGVuZENoaWxkKHRoaXMuI2J1aWxkUHJvcGVydHlUcmVlKG5lc3RlZE9iaiwgZGVwdGggKyAxKSk7XG4gICAgICAgIH1cbiAgICAgIH0gZWxzZSBpZiAodHlwZW9mIHZhbHVlID09PSAnb2JqZWN0Jykge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUgb2JqZWN0JztcbiAgICAgICAgY29uc3Qga2V5cyA9IE9iamVjdC5rZXlzKHZhbHVlIGFzIFJlY29yZDxzdHJpbmcsIHVua25vd24+KTtcbiAgICAgICAgdmFsdWVTcGFuLnRleHRDb250ZW50ID0ga2V5cy5sZW5ndGggPiAwID8gYE9iamVjdGAgOiBge31gO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuXG4gICAgICAgIGlmIChrZXlzLmxlbmd0aCA+IDAgJiYgZGVwdGggPCAzKSB7XG4gICAgICAgICAgbGkuYXBwZW5kQ2hpbGQodGhpcy4jYnVpbGRQcm9wZXJ0eVRyZWUodmFsdWUgYXMgUmVjb3JkPHN0cmluZywgdW5rbm93bj4sIGRlcHRoICsgMSkpO1xuICAgICAgICB9XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBjb25zdCB2YWx1ZVNwYW4gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG4gICAgICAgIHZhbHVlU3Bhbi5jbGFzc05hbWUgPSAncHJvcGVydHktdmFsdWUnO1xuICAgICAgICB2YWx1ZVNwYW4udGV4dENvbnRlbnQgPSBTdHJpbmcodmFsdWUpO1xuICAgICAgICBsaS5hcHBlbmRDaGlsZCh2YWx1ZVNwYW4pO1xuICAgICAgfVxuXG4gICAgICB1bC5hcHBlbmRDaGlsZChsaSk7XG4gICAgfVxuXG4gICAgcmV0dXJuIHVsO1xuICB9XG5cbiAgI3Njcm9sbEV2ZW50c1RvQm90dG9tKCkge1xuICAgIGlmICghdGhpcy4jZXZlbnRMaXN0KSByZXR1cm47XG5cbiAgICByZXF1ZXN0QW5pbWF0aW9uRnJhbWUoKCkgPT4ge1xuICAgICAgY29uc3QgbGFzdEV2ZW50ID0gdGhpcy4jZXZlbnRMaXN0IS5sYXN0RWxlbWVudENoaWxkO1xuICAgICAgaWYgKGxhc3RFdmVudCkge1xuICAgICAgICBsYXN0RXZlbnQuc2Nyb2xsSW50b1ZpZXcoeyBiZWhhdmlvcjogJ2F1dG8nLCBibG9jazogJ2VuZCcgfSk7XG4gICAgICB9XG4gICAgfSk7XG4gIH1cblxuICAjaXNFdmVudHNUYWJBY3RpdmUoKTogYm9vbGVhbiB7XG4gICAgY29uc3QgdGFicyA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3RvcigncGYtdjYtdGFicycpO1xuICAgIGlmICghdGFicykgcmV0dXJuIGZhbHNlO1xuXG4gICAgY29uc3Qgc2VsZWN0ZWRJbmRleCA9IHBhcnNlSW50KHRhYnMuZ2V0QXR0cmlidXRlKCdzZWxlY3RlZCcpIHx8ICcwJywgMTApO1xuICAgIHJldHVybiBzZWxlY3RlZEluZGV4ID09PSAzO1xuICB9XG5cbiAgI2ZpbHRlckV2ZW50cyhxdWVyeTogc3RyaW5nKSB7XG4gICAgdGhpcy4jZXZlbnRzRmlsdGVyVmFsdWUgPSBxdWVyeS50b0xvd2VyQ2FzZSgpO1xuXG4gICAgaWYgKCF0aGlzLiNldmVudExpc3QpIHJldHVybjtcblxuICAgIGZvciAoY29uc3QgZW50cnkgb2YgdGhpcy4jZXZlbnRMaXN0LmNoaWxkcmVuKSB7XG4gICAgICBjb25zdCBldmVudFJlY29yZCA9IHRoaXMuI2dldEV2ZW50UmVjb3JkQnlJZCgoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuZXZlbnRJZCEpO1xuXG4gICAgICBpZiAoIWV2ZW50UmVjb3JkKSBjb250aW51ZTtcblxuICAgICAgY29uc3QgdGV4dE1hdGNoID0gdGhpcy4jZXZlbnRNYXRjaGVzVGV4dEZpbHRlcihldmVudFJlY29yZCk7XG4gICAgICBjb25zdCB0eXBlTWF0Y2ggPSB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLnNpemUgPT09IDAgfHwgdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5oYXMoZXZlbnRSZWNvcmQuZXZlbnROYW1lKTtcbiAgICAgIGNvbnN0IGVsZW1lbnRNYXRjaCA9IHRoaXMuI2VsZW1lbnRGaWx0ZXJzLnNpemUgPT09IDAgfHwgdGhpcy4jZWxlbWVudEZpbHRlcnMuaGFzKGV2ZW50UmVjb3JkLnRhZ05hbWUpO1xuXG4gICAgICAoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmhpZGRlbiA9ICEodGV4dE1hdGNoICYmIHR5cGVNYXRjaCAmJiBlbGVtZW50TWF0Y2gpO1xuICAgIH1cbiAgfVxuXG4gICNldmVudE1hdGNoZXNUZXh0RmlsdGVyKGV2ZW50UmVjb3JkOiBFdmVudFJlY29yZCk6IGJvb2xlYW4ge1xuICAgIGlmICghdGhpcy4jZXZlbnRzRmlsdGVyVmFsdWUpIHJldHVybiB0cnVlO1xuXG4gICAgY29uc3Qgc2VhcmNoVGV4dCA9IFtcbiAgICAgIGV2ZW50UmVjb3JkLnRhZ05hbWUsXG4gICAgICBldmVudFJlY29yZC5ldmVudE5hbWUsXG4gICAgICBldmVudFJlY29yZC5lbGVtZW50SWQgfHwgJycsXG4gICAgICBKU09OLnN0cmluZ2lmeShldmVudFJlY29yZC5jdXN0b21Qcm9wZXJ0aWVzIHx8IHt9KVxuICAgIF0uam9pbignICcpLnRvTG93ZXJDYXNlKCk7XG5cbiAgICByZXR1cm4gc2VhcmNoVGV4dC5pbmNsdWRlcyh0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSk7XG4gIH1cblxuICAjZ2V0RXZlbnRSZWNvcmRCeUlkKGlkOiBzdHJpbmcpOiBFdmVudFJlY29yZCB8IHVuZGVmaW5lZCB7XG4gICAgcmV0dXJuIHRoaXMuI2NhcHR1cmVkRXZlbnRzLmZpbmQoZSA9PiBlLmlkID09PSBpZCk7XG4gIH1cblxuICAjdXBkYXRlRXZlbnRGaWx0ZXJzKCkge1xuICAgIGNvbnN0IHNhdmVkUHJlZmVyZW5jZXMgPSB0aGlzLiNsb2FkRXZlbnRGaWx0ZXJzRnJvbVN0b3JhZ2UoKTtcblxuICAgIGNvbnN0IGV2ZW50VHlwZUZpbHRlciA9IHRoaXMuIyQoJ2V2ZW50LXR5cGUtZmlsdGVyJyk7XG4gICAgaWYgKGV2ZW50VHlwZUZpbHRlciAmJiB0aGlzLiNlbGVtZW50RXZlbnRNYXApIHtcbiAgICAgIGxldCBtZW51ID0gZXZlbnRUeXBlRmlsdGVyLnF1ZXJ5U2VsZWN0b3IoJ3BmLXY2LW1lbnUnKTtcbiAgICAgIGlmICghbWVudSkge1xuICAgICAgICBtZW51ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgncGYtdjYtbWVudScpO1xuICAgICAgICBldmVudFR5cGVGaWx0ZXIuYXBwZW5kQ2hpbGQobWVudSk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGV4aXN0aW5nSXRlbXMgPSBtZW51LnF1ZXJ5U2VsZWN0b3JBbGwoJ3BmLXY2LW1lbnUtaXRlbScpO1xuICAgICAgZXhpc3RpbmdJdGVtcy5mb3JFYWNoKGl0ZW0gPT4gaXRlbS5yZW1vdmUoKSk7XG5cbiAgICAgIGNvbnN0IGFsbEV2ZW50VHlwZXMgPSBuZXcgU2V0PHN0cmluZz4oKTtcbiAgICAgIGZvciAoY29uc3QgW3RhZ05hbWUsIGV2ZW50SW5mb10gb2YgdGhpcy4jZWxlbWVudEV2ZW50TWFwKSB7XG4gICAgICAgIGlmICh0aGlzLiNkaXNjb3ZlcmVkRWxlbWVudHMuaGFzKHRhZ05hbWUpKSB7XG4gICAgICAgICAgZm9yIChjb25zdCBldmVudE5hbWUgb2YgZXZlbnRJbmZvLmV2ZW50TmFtZXMpIHtcbiAgICAgICAgICAgIGFsbEV2ZW50VHlwZXMuYWRkKGV2ZW50TmFtZSk7XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIGlmIChzYXZlZFByZWZlcmVuY2VzLmV2ZW50VHlwZXMpIHtcbiAgICAgICAgdGhpcy4jZXZlbnRUeXBlRmlsdGVycyA9IChzYXZlZFByZWZlcmVuY2VzLmV2ZW50VHlwZXMgYXMgU2V0PHN0cmluZz4gJiB7IGludGVyc2VjdGlvbjogKG90aGVyOiBTZXQ8c3RyaW5nPikgPT4gU2V0PHN0cmluZz4gfSkuaW50ZXJzZWN0aW9uKGFsbEV2ZW50VHlwZXMpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgdGhpcy4jZXZlbnRUeXBlRmlsdGVycyA9IG5ldyBTZXQoYWxsRXZlbnRUeXBlcyk7XG4gICAgICB9XG5cbiAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIG9mIGFsbEV2ZW50VHlwZXMpIHtcbiAgICAgICAgY29uc3QgaXRlbSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3BmLXY2LW1lbnUtaXRlbScpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgndmFyaWFudCcsICdjaGVja2JveCcpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgndmFsdWUnLCBldmVudE5hbWUpO1xuICAgICAgICBpZiAodGhpcy4jZXZlbnRUeXBlRmlsdGVycy5oYXMoZXZlbnROYW1lKSkge1xuICAgICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCdjaGVja2VkJywgJycpO1xuICAgICAgICB9XG4gICAgICAgIGl0ZW0udGV4dENvbnRlbnQgPSBldmVudE5hbWU7XG4gICAgICAgIG1lbnUuYXBwZW5kQ2hpbGQoaXRlbSk7XG4gICAgICB9XG4gICAgfVxuXG4gICAgY29uc3QgZWxlbWVudEZpbHRlciA9IHRoaXMuIyQoJ2VsZW1lbnQtZmlsdGVyJyk7XG4gICAgaWYgKGVsZW1lbnRGaWx0ZXIgJiYgdGhpcy4jZWxlbWVudEV2ZW50TWFwKSB7XG4gICAgICBsZXQgbWVudSA9IGVsZW1lbnRGaWx0ZXIucXVlcnlTZWxlY3RvcigncGYtdjYtbWVudScpO1xuICAgICAgaWYgKCFtZW51KSB7XG4gICAgICAgIG1lbnUgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwZi12Ni1tZW51Jyk7XG4gICAgICAgIGVsZW1lbnRGaWx0ZXIuYXBwZW5kQ2hpbGQobWVudSk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IGV4aXN0aW5nSXRlbXMgPSBtZW51LnF1ZXJ5U2VsZWN0b3JBbGwoJ3BmLXY2LW1lbnUtaXRlbScpO1xuICAgICAgZXhpc3RpbmdJdGVtcy5mb3JFYWNoKGl0ZW0gPT4gaXRlbS5yZW1vdmUoKSk7XG5cbiAgICAgIGNvbnN0IGFsbEVsZW1lbnRzID0gbmV3IFNldDxzdHJpbmc+KCk7XG4gICAgICBmb3IgKGNvbnN0IHRhZ05hbWUgb2YgdGhpcy4jZWxlbWVudEV2ZW50TWFwLmtleXMoKSkge1xuICAgICAgICBpZiAodGhpcy4jZGlzY292ZXJlZEVsZW1lbnRzLmhhcyh0YWdOYW1lKSkge1xuICAgICAgICAgIGFsbEVsZW1lbnRzLmFkZCh0YWdOYW1lKTtcbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICBpZiAoc2F2ZWRQcmVmZXJlbmNlcy5lbGVtZW50cykge1xuICAgICAgICB0aGlzLiNlbGVtZW50RmlsdGVycyA9IChzYXZlZFByZWZlcmVuY2VzLmVsZW1lbnRzIGFzIFNldDxzdHJpbmc+ICYgeyBpbnRlcnNlY3Rpb246IChvdGhlcjogU2V0PHN0cmluZz4pID0+IFNldDxzdHJpbmc+IH0pLmludGVyc2VjdGlvbihhbGxFbGVtZW50cyk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICB0aGlzLiNlbGVtZW50RmlsdGVycyA9IG5ldyBTZXQoYWxsRWxlbWVudHMpO1xuICAgICAgfVxuXG4gICAgICBmb3IgKGNvbnN0IHRhZ05hbWUgb2YgYWxsRWxlbWVudHMpIHtcbiAgICAgICAgY29uc3QgaXRlbSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3BmLXY2LW1lbnUtaXRlbScpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgndmFyaWFudCcsICdjaGVja2JveCcpO1xuICAgICAgICBpdGVtLnNldEF0dHJpYnV0ZSgndmFsdWUnLCB0YWdOYW1lKTtcbiAgICAgICAgaWYgKHRoaXMuI2VsZW1lbnRGaWx0ZXJzLmhhcyh0YWdOYW1lKSkge1xuICAgICAgICAgIGl0ZW0uc2V0QXR0cmlidXRlKCdjaGVja2VkJywgJycpO1xuICAgICAgICB9XG4gICAgICAgIGl0ZW0udGV4dENvbnRlbnQgPSBgPCR7dGFnTmFtZX0+YDtcbiAgICAgICAgbWVudS5hcHBlbmRDaGlsZChpdGVtKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAjaGFuZGxlRXZlbnRUeXBlRmlsdGVyQ2hhbmdlID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IHsgdmFsdWUsIGNoZWNrZWQgfSA9IGV2ZW50IGFzIEV2ZW50ICYgeyB2YWx1ZTogc3RyaW5nOyBjaGVja2VkOiBib29sZWFuIH07XG5cbiAgICBpZiAoIXZhbHVlKSByZXR1cm47XG5cbiAgICBpZiAoY2hlY2tlZCkge1xuICAgICAgdGhpcy4jZXZlbnRUeXBlRmlsdGVycy5hZGQodmFsdWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICB0aGlzLiNldmVudFR5cGVGaWx0ZXJzLmRlbGV0ZSh2YWx1ZSk7XG4gICAgfVxuXG4gICAgdGhpcy4jc2F2ZUV2ZW50RmlsdGVycygpO1xuICAgIHRoaXMuI2ZpbHRlckV2ZW50cyh0aGlzLiNldmVudHNGaWx0ZXJWYWx1ZSk7XG4gIH07XG5cbiAgI2hhbmRsZUVsZW1lbnRGaWx0ZXJDaGFuZ2UgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgeyB2YWx1ZSwgY2hlY2tlZCB9ID0gZXZlbnQgYXMgRXZlbnQgJiB7IHZhbHVlOiBzdHJpbmc7IGNoZWNrZWQ6IGJvb2xlYW4gfTtcblxuICAgIGlmICghdmFsdWUpIHJldHVybjtcblxuICAgIGlmIChjaGVja2VkKSB7XG4gICAgICB0aGlzLiNlbGVtZW50RmlsdGVycy5hZGQodmFsdWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICB0aGlzLiNlbGVtZW50RmlsdGVycy5kZWxldGUodmFsdWUpO1xuICAgIH1cblxuICAgIHRoaXMuI3NhdmVFdmVudEZpbHRlcnMoKTtcbiAgICB0aGlzLiNmaWx0ZXJFdmVudHModGhpcy4jZXZlbnRzRmlsdGVyVmFsdWUpO1xuICB9O1xuXG4gICNsb2FkRXZlbnRGaWx0ZXJzRnJvbVN0b3JhZ2UoKTogeyBldmVudFR5cGVzOiBTZXQ8c3RyaW5nPiB8IG51bGw7IGVsZW1lbnRzOiBTZXQ8c3RyaW5nPiB8IG51bGwgfSB7XG4gICAgY29uc3QgcHJlZmVyZW5jZXM6IHsgZXZlbnRUeXBlczogU2V0PHN0cmluZz4gfCBudWxsOyBlbGVtZW50czogU2V0PHN0cmluZz4gfCBudWxsIH0gPSB7XG4gICAgICBldmVudFR5cGVzOiBudWxsLFxuICAgICAgZWxlbWVudHM6IG51bGxcbiAgICB9O1xuXG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHNhdmVkRXZlbnRUeXBlcyA9IGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdjZW0tc2VydmUtZXZlbnQtdHlwZS1maWx0ZXJzJyk7XG4gICAgICBpZiAoc2F2ZWRFdmVudFR5cGVzKSB7XG4gICAgICAgIHByZWZlcmVuY2VzLmV2ZW50VHlwZXMgPSBuZXcgU2V0KEpTT04ucGFyc2Uoc2F2ZWRFdmVudFR5cGVzKSk7XG4gICAgICB9XG5cbiAgICAgIGNvbnN0IHNhdmVkRWxlbWVudHMgPSBsb2NhbFN0b3JhZ2UuZ2V0SXRlbSgnY2VtLXNlcnZlLWVsZW1lbnQtZmlsdGVycycpO1xuICAgICAgaWYgKHNhdmVkRWxlbWVudHMpIHtcbiAgICAgICAgcHJlZmVyZW5jZXMuZWxlbWVudHMgPSBuZXcgU2V0KEpTT04ucGFyc2Uoc2F2ZWRFbGVtZW50cykpO1xuICAgICAgfVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIGNvbnNvbGUuZGVidWcoJ1tjZW0tc2VydmUtY2hyb21lXSBsb2NhbFN0b3JhZ2UgdW5hdmFpbGFibGUgZm9yIGV2ZW50IGZpbHRlcnMnKTtcbiAgICB9XG5cbiAgICByZXR1cm4gcHJlZmVyZW5jZXM7XG4gIH1cblxuICAjc2F2ZUV2ZW50RmlsdGVycygpIHtcbiAgICB0cnkge1xuICAgICAgbG9jYWxTdG9yYWdlLnNldEl0ZW0oJ2NlbS1zZXJ2ZS1ldmVudC10eXBlLWZpbHRlcnMnLFxuICAgICAgICBKU09OLnN0cmluZ2lmeShbLi4udGhpcy4jZXZlbnRUeXBlRmlsdGVyc10pKTtcbiAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdjZW0tc2VydmUtZWxlbWVudC1maWx0ZXJzJyxcbiAgICAgICAgSlNPTi5zdHJpbmdpZnkoWy4uLnRoaXMuI2VsZW1lbnRGaWx0ZXJzXSkpO1xuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIC8vIGxvY2FsU3RvcmFnZSB1bmF2YWlsYWJsZSAocHJpdmF0ZSBtb2RlKSwgc2lsZW50bHkgY29udGludWVcbiAgICB9XG4gIH1cblxuICAjY2xlYXJFdmVudHMoKSB7XG4gICAgdGhpcy4jY2FwdHVyZWRFdmVudHMgPSBbXTtcbiAgICB0aGlzLiNzZWxlY3RlZEV2ZW50SWQgPSBudWxsO1xuICAgIGlmICh0aGlzLiNldmVudExpc3QpIHtcbiAgICAgIHRoaXMuI2V2ZW50TGlzdC5yZXBsYWNlQ2hpbGRyZW4oKTtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyKSB7XG4gICAgICB0aGlzLiNldmVudERldGFpbEhlYWRlci5pbm5lckhUTUwgPSAnJztcbiAgICB9XG4gICAgaWYgKHRoaXMuI2V2ZW50RGV0YWlsQm9keSkge1xuICAgICAgdGhpcy4jZXZlbnREZXRhaWxCb2R5LmlubmVySFRNTCA9ICcnO1xuICAgIH1cbiAgfVxuXG4gIGFzeW5jICNjb3B5RXZlbnRzKCkge1xuICAgIGlmICghdGhpcy4jZXZlbnRMaXN0KSByZXR1cm47XG5cbiAgICBjb25zdCB2aXNpYmxlRXZlbnRzID0gQXJyYXkuZnJvbSh0aGlzLiNldmVudExpc3QuY2hpbGRyZW4pXG4gICAgICAuZmlsdGVyKGVudHJ5ID0+ICEoZW50cnkgYXMgSFRNTEVsZW1lbnQpLmhpZGRlbilcbiAgICAgIC5tYXAoZW50cnkgPT4ge1xuICAgICAgICBjb25zdCBpZCA9IChlbnRyeSBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5ldmVudElkITtcbiAgICAgICAgcmV0dXJuIHRoaXMuI2dldEV2ZW50UmVjb3JkQnlJZChpZCk7XG4gICAgICB9KVxuICAgICAgLmZpbHRlcigoZXZlbnQpOiBldmVudCBpcyBFdmVudFJlY29yZCA9PiAhIWV2ZW50KVxuICAgICAgLm1hcChldmVudCA9PiB7XG4gICAgICAgIGNvbnN0IHRpbWUgPSBldmVudC50aW1lc3RhbXAudG9Mb2NhbGVUaW1lU3RyaW5nKCk7XG4gICAgICAgIGNvbnN0IGVsZW1lbnQgPSBldmVudC5lbGVtZW50SWQgPyBgIyR7ZXZlbnQuZWxlbWVudElkfWAgOiBldmVudC50YWdOYW1lO1xuICAgICAgICBjb25zdCBwcm9wcyA9IGV2ZW50LmN1c3RvbVByb3BlcnRpZXMgJiYgT2JqZWN0LmtleXMoZXZlbnQuY3VzdG9tUHJvcGVydGllcykubGVuZ3RoID4gMFxuICAgICAgICAgID8gYCAtICR7SlNPTi5zdHJpbmdpZnkoZXZlbnQuY3VzdG9tUHJvcGVydGllcyl9YFxuICAgICAgICAgIDogJyc7XG4gICAgICAgIHJldHVybiBgWyR7dGltZX1dIDwke2V2ZW50LnRhZ05hbWV9PiAke2VsZW1lbnR9IFxcdTIxOTIgJHtldmVudC5ldmVudE5hbWV9JHtwcm9wc31gO1xuICAgICAgfSlcbiAgICAgIC5qb2luKCdcXG4nKTtcblxuICAgIGlmICghdmlzaWJsZUV2ZW50cykgcmV0dXJuO1xuXG4gICAgdHJ5IHtcbiAgICAgIGF3YWl0IG5hdmlnYXRvci5jbGlwYm9hcmQud3JpdGVUZXh0KHZpc2libGVFdmVudHMpO1xuICAgICAgY29uc3QgYnRuID0gdGhpcy4jJCgnY29weS1ldmVudHMnKTtcbiAgICAgIGlmIChidG4pIHtcbiAgICAgICAgY29uc3QgdGV4dE5vZGUgPSBBcnJheS5mcm9tKGJ0bi5jaGlsZE5vZGVzKS5maW5kKFxuICAgICAgICAgIG4gPT4gbi5ub2RlVHlwZSA9PT0gTm9kZS5URVhUX05PREUgJiYgKG4udGV4dENvbnRlbnQ/LnRyaW0oKS5sZW5ndGggPz8gMCkgPiAwXG4gICAgICAgICk7XG4gICAgICAgIGlmICh0ZXh0Tm9kZSkge1xuICAgICAgICAgIGNvbnN0IG9yaWdpbmFsID0gdGV4dE5vZGUudGV4dENvbnRlbnQ7XG4gICAgICAgICAgdGV4dE5vZGUudGV4dENvbnRlbnQgPSAnQ29waWVkISc7XG5cbiAgICAgICAgICBpZiAodGhpcy4jY29weUV2ZW50c0ZlZWRiYWNrVGltZW91dCkge1xuICAgICAgICAgICAgY2xlYXJUaW1lb3V0KHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgICAgIH1cblxuICAgICAgICAgIHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQgPSBzZXRUaW1lb3V0KCgpID0+IHtcbiAgICAgICAgICAgIGlmICh0aGlzLmlzQ29ubmVjdGVkICYmIHRleHROb2RlLnBhcmVudE5vZGUpIHtcbiAgICAgICAgICAgICAgdGV4dE5vZGUudGV4dENvbnRlbnQgPSBvcmlnaW5hbDtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgICAgICAgIH0sIDIwMDApO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfSBjYXRjaCAoZXJyKSB7XG4gICAgICBjb25zb2xlLmVycm9yKCdbY2VtLXNlcnZlLWNocm9tZV0gRmFpbGVkIHRvIGNvcHkgZXZlbnRzOicsIGVycik7XG4gICAgfVxuICB9XG5cbiAgI3NldHVwRXZlbnRMaXN0ZW5lcnMoKSB7XG4gICAgdGhpcy4jZXZlbnRMaXN0ID0gdGhpcy4jJCgnZXZlbnQtbGlzdCcpO1xuICAgIHRoaXMuI2V2ZW50RGV0YWlsSGVhZGVyID0gdGhpcy4jJCgnZXZlbnQtZGV0YWlsLWhlYWRlcicpO1xuICAgIHRoaXMuI2V2ZW50RGV0YWlsQm9keSA9IHRoaXMuIyQoJ2V2ZW50LWRldGFpbC1ib2R5Jyk7XG5cbiAgICBpZiAodGhpcy4jZXZlbnRMaXN0KSB7XG4gICAgICB0aGlzLiNldmVudExpc3QuYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCAoZTogRXZlbnQpID0+IHtcbiAgICAgICAgY29uc3QgbGlzdEl0ZW0gPSAoZS50YXJnZXQgYXMgRWxlbWVudCkuY2xvc2VzdCgnLmV2ZW50LWxpc3QtaXRlbScpIGFzIEhUTUxFbGVtZW50O1xuICAgICAgICBpZiAobGlzdEl0ZW0pIHtcbiAgICAgICAgICBjb25zdCBldmVudElkID0gbGlzdEl0ZW0uZGF0YXNldC5ldmVudElkO1xuICAgICAgICAgIGlmIChldmVudElkKSB7XG4gICAgICAgICAgICB0aGlzLiNzZWxlY3RFdmVudChldmVudElkKTtcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH0pO1xuICAgIH1cblxuICAgIGNvbnN0IGV2ZW50c0ZpbHRlciA9IHRoaXMuIyQoJ2V2ZW50cy1maWx0ZXInKTtcbiAgICBpZiAoZXZlbnRzRmlsdGVyKSB7XG4gICAgICBldmVudHNGaWx0ZXIuYWRkRXZlbnRMaXN0ZW5lcignaW5wdXQnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgICAgY29uc3QgeyB2YWx1ZSA9ICcnIH0gPSBlLnRhcmdldCBhcyBIVE1MSW5wdXRFbGVtZW50O1xuICAgICAgICBjbGVhclRpbWVvdXQodGhpcy4jZXZlbnRzRmlsdGVyRGVib3VuY2VUaW1lciEpO1xuICAgICAgICB0aGlzLiNldmVudHNGaWx0ZXJEZWJvdW5jZVRpbWVyID0gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgICAgdGhpcy4jZmlsdGVyRXZlbnRzKHZhbHVlKTtcbiAgICAgICAgfSwgMzAwKTtcbiAgICAgIH0pO1xuICAgIH1cblxuICAgIGNvbnN0IGV2ZW50VHlwZUZpbHRlciA9IHRoaXMuIyQoJ2V2ZW50LXR5cGUtZmlsdGVyJyk7XG4gICAgaWYgKGV2ZW50VHlwZUZpbHRlcikge1xuICAgICAgZXZlbnRUeXBlRmlsdGVyLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI2hhbmRsZUV2ZW50VHlwZUZpbHRlckNoYW5nZSBhcyBFdmVudExpc3RlbmVyKTtcbiAgICB9XG5cbiAgICBjb25zdCBlbGVtZW50RmlsdGVyID0gdGhpcy4jJCgnZWxlbWVudC1maWx0ZXInKTtcbiAgICBpZiAoZWxlbWVudEZpbHRlcikge1xuICAgICAgZWxlbWVudEZpbHRlci5hZGRFdmVudExpc3RlbmVyKCdzZWxlY3QnLCB0aGlzLiNoYW5kbGVFbGVtZW50RmlsdGVyQ2hhbmdlIGFzIEV2ZW50TGlzdGVuZXIpO1xuICAgIH1cblxuICAgIHRoaXMuIyQoJ2NsZWFyLWV2ZW50cycpPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgIHRoaXMuI2NsZWFyRXZlbnRzKCk7XG4gICAgfSk7XG5cbiAgICB0aGlzLiMkKCdjb3B5LWV2ZW50cycpPy5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IHtcbiAgICAgIHRoaXMuI2NvcHlFdmVudHMoKTtcbiAgICB9KTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG5cbiAgICAvLyBDbGVhbiB1cCBrbm9iIGxpc3RlbmVyc1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna25vYjphdHRyaWJ1dGUtY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6cHJvcGVydHktY2hhbmdlJywgdGhpcy4jb25Lbm9iQ2hhbmdlKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6Y3NzLXByb3BlcnR5LWNoYW5nZScsIHRoaXMuI29uS25vYkNoYW5nZSk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOmF0dHJpYnV0ZS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tub2I6cHJvcGVydHktY2xlYXInLCB0aGlzLiNvbktub2JDbGVhcik7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdrbm9iOmNzcy1wcm9wZXJ0eS1jbGVhcicsIHRoaXMuI29uS25vYkNsZWFyKTtcblxuICAgIC8vIENsZWFuIHVwIHRyZWUgc3RhdGUgbGlzdGVuZXJzXG4gICAgaWYgKHRoaXMuI2hhbmRsZVRyZWVFeHBhbmQpIHtcbiAgICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcignZXhwYW5kJywgdGhpcy4jaGFuZGxlVHJlZUV4cGFuZCk7XG4gICAgfVxuICAgIGlmICh0aGlzLiNoYW5kbGVUcmVlQ29sbGFwc2UpIHtcbiAgICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcignY29sbGFwc2UnLCB0aGlzLiNoYW5kbGVUcmVlQ29sbGFwc2UpO1xuICAgIH1cbiAgICBpZiAodGhpcy4jaGFuZGxlVHJlZVNlbGVjdCkge1xuICAgICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdzZWxlY3QnLCB0aGlzLiNoYW5kbGVUcmVlU2VsZWN0KTtcbiAgICB9XG5cbiAgICAvLyBDbGVhbiB1cCB3aW5kb3cgbGlzdGVuZXJcbiAgICBpZiAodGhpcy4jaGFuZGxlTG9nc0V2ZW50KSB7XG4gICAgICB3aW5kb3cucmVtb3ZlRXZlbnRMaXN0ZW5lcignY2VtOmxvZ3MnLCB0aGlzLiNoYW5kbGVMb2dzRXZlbnQpO1xuICAgIH1cblxuICAgIC8vIENsZWFyIHBlbmRpbmcgZmVlZGJhY2sgdGltZW91dHNcbiAgICBpZiAodGhpcy4jY29weUxvZ3NGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCk7XG4gICAgICB0aGlzLiNjb3B5TG9nc0ZlZWRiYWNrVGltZW91dCA9IG51bGw7XG4gICAgfVxuICAgIGlmICh0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5RGVidWdGZWVkYmFja1RpbWVvdXQpO1xuICAgICAgdGhpcy4jY29weURlYnVnRmVlZGJhY2tUaW1lb3V0ID0gbnVsbDtcbiAgICB9XG4gICAgaWYgKHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQpIHtcbiAgICAgIGNsZWFyVGltZW91dCh0aGlzLiNjb3B5RXZlbnRzRmVlZGJhY2tUaW1lb3V0KTtcbiAgICAgIHRoaXMuI2NvcHlFdmVudHNGZWVkYmFja1RpbWVvdXQgPSBudWxsO1xuICAgIH1cblxuICAgIC8vIERpc2Nvbm5lY3QgbXV0YXRpb24gb2JzZXJ2ZXJcbiAgICB0aGlzLiNvYnNlcnZlci5kaXNjb25uZWN0KCk7XG5cbiAgICAvLyBDbG9zZSBXZWJTb2NrZXQgY29ubmVjdGlvblxuICAgIGlmICh0aGlzLiN3c0NsaWVudCkge1xuICAgICAgdGhpcy4jd3NDbGllbnQuZGVzdHJveSgpO1xuICAgIH1cbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tc2VydmUtY2hyb21lJzogQ2VtU2VydmVDaHJvbWU7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIGhlaWdodDogMTAwdmg7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG4gIC0tcGYtdjYtYy1tYXN0aGVhZF9fbG9nby0tV2lkdGg6IG1heC1jb250ZW50O1xcXFxuICAtLXBmLXY2LWMtbWFzdGhlYWRfX3RvZ2dsZS0tU2l6ZTogMXJlbTtcXFxcbn1cXFxcblxcXFxuW2hpZGRlbl0ge1xcXFxuICBkaXNwbGF5OiBub25lICFpbXBvcnRhbnQ7XFxcXG59XFxcXG5cXFxcbi8qIE1hc3RoZWFkIGxvZ28gc3R5bGVzICovXFxcXG4ubWFzdGhlYWQtbG9nbyB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcbiAgY29sb3I6IGluaGVyaXQ7XFxcXG4gIG1heC1oZWlnaHQ6IHZhcigtLXBmLXY2LWMtbWFzdGhlYWRfX2xvZ28tLU1heEhlaWdodCk7XFxcXG4gIGdhcDogNHB4O1xcXFxuICBcXFxcdTAwMjYgaW1nIHtcXFxcbiAgICBkaXNwbGF5OiBibG9jaztcXFxcbiAgICBtYXgtaGVpZ2h0OiB2YXIoLS1wZi12Ni1jLW1hc3RoZWFkX19sb2dvLS1NYXhIZWlnaHQpO1xcXFxuICAgIHdpZHRoOiBhdXRvO1xcXFxuICB9XFxcXG4gIFxcXFx1MDAyNiA6OnNsb3R0ZWQoW3Nsb3Q9XFxcXFxcXCJ0aXRsZVxcXFxcXFwiXSkge1xcXFxuICAgIG1hcmdpbjogMDtcXFxcbiAgICBmb250LXNpemU6IDEuMTI1cmVtO1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICB9XFxcXG4gIFxcXFx1MDAyNiBoMSB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICAgIGZvbnQtc2l6ZTogMThweDtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBUb29sYmFyIGdyb3VwIGFsaWdubWVudCAqL1xcXFxucGYtdjYtdG9vbGJhci1ncm91cFt2YXJpYW50PVxcXFxcXFwiYWN0aW9uLWdyb3VwXFxcXFxcXCJdIHtcXFxcbiAgbWFyZ2luLWlubGluZS1zdGFydDogYXV0bztcXFxcbiAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbn1cXFxcblxcXFxuLmRlYnVnLXBhbmVsIHtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlcjogMXB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlci1yYWRpdXM6IDZweDtcXFxcbiAgcGFkZGluZzogMS41cmVtO1xcXFxuICBtYXgtd2lkdGg6IDYwMHB4O1xcXFxuICB3aWR0aDogOTAlO1xcXFxuICBtYXgtaGVpZ2h0OiA4MHZoO1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuXFxcXG4gIGgyIHtcXFxcbiAgICBtYXJnaW46IDAgMCAxcmVtIDA7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgIGZvbnQtc2l6ZTogMS4xMjVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGRsIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gIH1cXFxcblxcXFxuICBkdCB7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICBtYXJnaW4tdG9wOiAwLjVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDUwMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGRkIHtcXFxcbiAgICBtYXJnaW46IDAgMCAwLjVyZW0gMDtcXFxcbiAgICBmb250LWZhbWlseTogdWktbW9ub3NwYWNlLCAnQ2FzY2FkaWEgQ29kZScsICdTb3VyY2UgQ29kZSBQcm8nLCBNZW5sbywgQ29uc29sYXMsICdEZWphVnUgU2FucyBNb25vJywgbW9ub3NwYWNlO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICB9XFxcXG5cXFxcbiAgZGV0YWlscyB7XFxcXG4gICAgbWFyZ2luLXRvcDogMXJlbTtcXFxcblxcXFxuICAgIHN1bW1hcnkge1xcXFxuICAgICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICAgICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgICAgIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG4gICAgICBsaXN0LXN0eWxlOiBub25lO1xcXFxuICAgICAgZGlzcGxheTogZmxleDtcXFxcbiAgICAgIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gICAgICBnYXA6IDAuNXJlbTtcXFxcbiAgICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuXFxcXG4gICAgICBcXFxcdTAwMjY6Oi13ZWJraXQtZGV0YWlscy1tYXJrZXIge1xcXFxuICAgICAgICBkaXNwbGF5OiBub25lO1xcXFxuICAgICAgfVxcXFxuXFxcXG4gICAgICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgICAgIGNvbnRlbnQ6ICdcXFxcXFxcXDI1QjgnO1xcXFxuICAgICAgICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7XFxcXG4gICAgICAgIHRyYW5zaXRpb246IHRyYW5zZm9ybSAxMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICAgIH1cXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjZbb3Blbl0gc3VtbWFyeTo6YmVmb3JlIHtcXFxcbiAgICAgIHRyYW5zZm9ybTogcm90YXRlKDkwZGVnKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBwcmUge1xcXFxuICAgICAgbWFyZ2luLXRvcDogMC41cmVtO1xcXFxuICAgICAgbWFyZ2luLWxlZnQ6IDEuNXJlbTtcXFxcbiAgICAgIHBhZGRpbmc6IDAuNXJlbTtcXFxcbiAgICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0tZGVmYXVsdCk7XFxcXG4gICAgICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gICAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICAgICAgb3ZlcmZsb3cteDogYXV0bztcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICAuYnV0dG9uLWNvbnRhaW5lciB7XFxcXG4gICAgZGlzcGxheTogZmxleDtcXFxcbiAgICBnYXA6IDAuNXJlbTtcXFxcbiAgICBtYXJnaW4tdG9wOiAxcmVtO1xcXFxuICB9XFxcXG5cXFxcbiAgcCB7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcbiAgfVxcXFxuXFxcXG4gIGJ1dHRvbiB7XFxcXG4gICAgbWFyZ2luLXRvcDogMXJlbTtcXFxcbiAgICBwYWRkaW5nOiAwLjVyZW0gMXJlbTtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLW9uLWJyYW5kKTtcXFxcbiAgICBib3JkZXI6IG5vbmU7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgZm9udC13ZWlnaHQ6IDQwMDtcXFxcbiAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gICAgdHJhbnNpdGlvbjogYWxsIDIwMG1zIGN1YmljLWJlemllcigwLjY0NSwgMC4wNDUsIDAuMzU1LCAxKTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1icmFuZC0taG92ZXIpO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBDb250ZW50IGFyZWEgcGFkZGluZyBmb3IgZGVtbyAqL1xcXFxucGYtdjYtcGFnZS1tYWluIHtcXFxcbiAgbWluLWhlaWdodDogY2FsYygxMDBkdmggLSA3MnB4IC0gdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWluc2V0LS1wYWdlLWNocm9tZSkpO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBcXFxcdTAwMjYgXFxcXHUwMDNlIDo6c2xvdHRlZCg6bm90KFtzbG90PWtub2JzXSkpIHtcXFxcbiAgICBwYWRkaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAgIGZsZXg6IDE7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuY2VtLWRyYXdlciB7XFxcXG4gIHBmLXY2LXRhYnMge1xcXFxuICAgIHBmLXY2LXRhYiB7XFxcXG4gICAgICBwYWRkaW5nLWJsb2NrLWVuZDogMDtcXFxcbiAgICB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogRWxlbWVudCBkZXNjcmlwdGlvbnMgaW4gbGlzdGluZyAqL1xcXFxuLmVsZW1lbnQtc3VtbWFyeSB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG59XFxcXG5cXFxcbi5lbGVtZW50LWRlc2NyaXB0aW9uIHtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtKTtcXFxcbn1cXFxcblxcXFxuLyogQ2FyZCBmb290ZXIgZGVtbyBuYXZpZ2F0aW9uICovXFxcXG4uY2FyZC1kZW1vcyB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtd3JhcDogd3JhcDtcXFxcbiAgZ2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tZ2FwLS1hY3Rpb24tdG8tYWN0aW9uLS1kZWZhdWx0KTtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgbWFyZ2luOiAwO1xcXFxufVxcXFxuXFxcXG4ucGFja2FnZS1uYW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSk7XFxcXG59XFxcXG5cXFxcbi8qIEtub2JzIGNvbnRhaW5lciAtIGZpbGxzIHRhYiBwYW5lbCBoZWlnaHQgKi9cXFxcbiNrbm9icy1jb250YWluZXIge1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIFxcXFx1MDAyNiA6OnNsb3R0ZWQoW3Nsb3Q9XFxcXFxcXCJrbm9ic1xcXFxcXFwiXSkge1xcXFxuICAgIGZsZXg6IDE7XFxcXG4gICAgbWluLWhlaWdodDogMDtcXFxcbiAgICBvdmVyZmxvdzogaGlkZGVuO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5rbm9icy1lbXB0eSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICB0ZXh0LWFsaWduOiBjZW50ZXI7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbGcpO1xcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBwYWRkaW5nOiAycHggNnB4O1xcXFxuICAgIGJvcmRlci1yYWRpdXM6IDNweDtcXFxcbiAgICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmluc3RhbmNlLXRhZyB7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLmluc3RhbmNlLWxhYmVsIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4ua25vYi1ncm91cCB7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbGcpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjpsYXN0LWNoaWxkIHtcXFxcbiAgICBtYXJnaW4tYm90dG9tOiAwO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi8qIFBhdHRlcm5GbHkgdjYgZm9ybSAtIGhvcml6b250YWwgbGF5b3V0ICovXFxcXG5wZi12Ni1mb3JtW2hvcml6b250YWxdIHBmLXY2LWZvcm0tZmllbGQtZ3JvdXAge1xcXFxuICBncmlkLWNvbHVtbjogc3BhbiAyXFxcXG59XFxcXG5cXFxcbi5rbm9iLWdyb3VwLXRpdGxlIHtcXFxcbiAgZ3JpZC1jb2x1bW46IDEgLyAtMTtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgMDtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLWJhc2UpO1xcXFxuICBmb250LXdlaWdodDogNjAwO1xcXFxuICBib3JkZXItYm90dG9tOiAxcHggc29saWQgdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLWNvbG9yKTtcXFxcbiAgcGFkZGluZy1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxufVxcXFxuXFxcXG4ua25vYi1jb250cm9sIHtcXFxcbiAgbWFyZ2luLWJvdHRvbTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG59XFxcXG5cXFxcbi5rbm9iLWxhYmVsIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKTtcXFxcbiAgY3Vyc29yOiBwb2ludGVyO1xcXFxufVxcXFxuXFxcXG4ua25vYi1uYW1lIHtcXFxcbiAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxuICBmb250LXdlaWdodDogNTAwO1xcXFxufVxcXFxuXFxcXG4ua25vYi1kZXNjcmlwdGlvbiB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgbGluZS1oZWlnaHQ6IDEuNTtcXFxcblxcXFxuICBwIHtcXFxcbiAgICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmcteHMpIDA7XFxcXG4gIH1cXFxcblxcXFxuICBjb2RlIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogM3B4O1xcXFxuICAgIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgfVxcXFxuXFxcXG4gIGEge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1hY2NlbnQtY29sb3IpO1xcXFxuICAgIHRleHQtZGVjb3JhdGlvbjogbm9uZTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICB0ZXh0LWRlY29yYXRpb246IHVuZGVybGluZTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBzdHJvbmcge1xcXFxuICAgIGZvbnQtd2VpZ2h0OiA2MDA7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG4gIH1cXFxcblxcXFxuICB1bCwgb2wge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgICBwYWRkaW5nLWxlZnQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbGcpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbmZvb3Rlci5wZi1tLXN0aWNreS1ib3R0b20ge1xcXFxuICB2aWV3LXRyYW5zaXRpb24tbmFtZTogZGV2LXNlcnZlci1mb290ZXI7XFxcXG4gIHBvc2l0aW9uOiBzdGlja3k7XFxcXG4gIGJvdHRvbTogMDtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tcHJpbWFyeS0tZGVmYXVsdCk7XFxcXG4gIGJvcmRlci10b3A6IDFweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQpO1xcXFxuICB6LWluZGV4OiB2YXIoLS1wZi12Ni1jLXBhZ2UtLXNlY3Rpb24tLW0tc3RpY2t5LWJvdHRvbS0tWkluZGV4LCB2YXIoLS1wZi10LS1nbG9iYWwtLXotaW5kZXgtLW1kKSk7XFxcXG4gIGJveC1zaGFkb3c6IHZhcigtLXBmLXY2LWMtcGFnZS0tc2VjdGlvbi0tbS1zdGlja3ktYm90dG9tLS1Cb3hTaGFkb3csIHZhcigtLXBmLXQtLWdsb2JhbC0tYm94LXNoYWRvdy0tc20tLXRvcCkpO1xcXFxufVxcXFxuXFxcXG4uZm9vdGVyLWRlc2NyaXB0aW9uIHtcXFxcbiAgcGFkZGluZzogMS41cmVtO1xcXFxuXFxcXG4gIFxcXFx1MDAyNi5lbXB0eSB7XFxcXG4gICAgZGlzcGxheTogbm9uZTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG5mb290ZXIgOjpzbG90dGVkKFtzbG90PVxcXFxcXFwiZGVzY3JpcHRpb25cXFxcXFxcIl0pIHtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgbGluZS1oZWlnaHQ6IDEuNjtcXFxcbiAgZm9udC1zaXplOiAwLjg3NXJlbTtcXFxcblxcXFxuICBjb2RlIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1ob3Zlcik7XFxcXG4gICAgcGFkZGluZzogMnB4IDZweDtcXFxcbiAgICBib3JkZXItcmFkaXVzOiAzcHg7XFxcXG4gICAgZm9udC1mYW1pbHk6IHVpLW1vbm9zcGFjZSwgJ0Nhc2NhZGlhIENvZGUnLCAnU291cmNlIENvZGUgUHJvJywgTWVubG8sIENvbnNvbGFzLCAnRGVqYVZ1IFNhbnMgTW9ubycsIG1vbm9zcGFjZTtcXFxcbiAgICBmb250LXNpemU6IDAuODc1cmVtO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5sb2dzLXdyYXBwZXIge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbiNsb2ctY29udGFpbmVyIHtcXFxcbiAgZmxleC1ncm93OiAxO1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxufVxcXFxuXFxcXG4ubG9nLWVudHJ5IHtcXFxcbiAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG4gIGFsaWduLWl0ZW1zOiBiYXNlbGluZTtcXFxcbiAgcGYtdjYtbGFiZWwge1xcXFxuICAgIGZsZXgtc2hyaW5rOiAwO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5sb2ctdGltZSxcXFxcbi5sb2ctbWVzc2FnZSB7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4ubG9nLXRpbWUge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZsZXgtc2hyaW5rOiAwO1xcXFxuICBmb250LXNpemU6IDExcHg7XFxcXG59XFxcXG5cXFxcbi5sb2ctbWVzc2FnZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXByaW1hcnkpO1xcXFxuICB3b3JkLWJyZWFrOiBicmVhay13b3JkO1xcXFxufVxcXFxuXFxcXG4vKiBOYXZpZ2F0aW9uIGNvbnRlbnQgKGxpZ2h0IERPTSBzbG90dGVkIGNvbnRlbnQgZm9yIHBmLXY2LXBhZ2Utc2lkZWJhcikgKi9cXFxcbi5uYXYtcGFja2FnZSB7XFxcXG4gIG1hcmdpbi1ib3R0b206IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNiBcXFxcdTAwM2Ugc3VtbWFyeSB7XFxcXG4gICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICAgIHBhZGRpbmc6IDAuNXJlbSAxcmVtO1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0tZGVmYXVsdCk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICBmb250LXdlaWdodDogNjAwO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kIDIwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgbWFyZ2luLWJvdHRvbTogMC41cmVtO1xcXFxuICAgIGRpc3BsYXk6IGZsZXg7XFxcXG4gICAgYWxpZ24taXRlbXM6IGNlbnRlcjtcXFxcbiAgICBnYXA6IDAuNXJlbTtcXFxcbiAgICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcblxcXFxuICAgIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1zZWNvbmRhcnktLWhvdmVyKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6Oi13ZWJraXQtZGV0YWlscy1tYXJrZXIge1xcXFxuICAgICAgZGlzcGxheTogbm9uZTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6OmJlZm9yZSB7XFxcXG4gICAgICBjb250ZW50OiAnXFxcXFxcXFwyNUI4JztcXFxcbiAgICAgIGRpc3BsYXk6IGlubGluZS1ibG9jaztcXFxcbiAgICAgIHRyYW5zaXRpb246IHRyYW5zZm9ybSAxMDBtcyBjdWJpYy1iZXppZXIoMC40LCAwLCAwLjIsIDEpO1xcXFxuICAgICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tdGV4dC0tY29sb3ItLXN1YnRsZSk7XFxcXG4gICAgfVxcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2W29wZW5dIFxcXFx1MDAzZSBzdW1tYXJ5OjpiZWZvcmUge1xcXFxuICAgIHRyYW5zZm9ybTogcm90YXRlKDkwZGVnKTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4ubmF2LWVsZW1lbnQge1xcXFxuICBtYXJnaW4tYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKTtcXFxcbiAgbWFyZ2luLWlubGluZS1zdGFydDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG5cXFxcbiAgc3VtbWFyeSB7XFxcXG4gICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuICAgIHBhZGRpbmc6IDAuNXJlbSAxcmVtO1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0tZGVmYXVsdCk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogNnB4O1xcXFxuICAgIGNvbG9yOiB2YXIoLS1wZi10LS1nbG9iYWwtLXRleHQtLWNvbG9yLS1yZWd1bGFyKTtcXFxcbiAgICBmb250LWZhbWlseTogdWktbW9ub3NwYWNlLCAnQ2FzY2FkaWEgQ29kZScsICdTb3VyY2UgQ29kZSBQcm8nLCBNZW5sbywgQ29uc29sYXMsICdEZWphVnUgU2FucyBNb25vJywgbW9ub3NwYWNlO1xcXFxuICAgIGZvbnQtc2l6ZTogMC44NzVyZW07XFxcXG4gICAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgICB0cmFuc2l0aW9uOiBiYWNrZ3JvdW5kIDIwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgZGlzcGxheTogZmxleDtcXFxcbiAgICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICAgIGdhcDogMC41cmVtO1xcXFxuICAgIHVzZXItc2VsZWN0OiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0taG92ZXIpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNjo6LXdlYmtpdC1kZXRhaWxzLW1hcmtlciB7XFxcXG4gICAgICBkaXNwbGF5OiBub25lO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIFxcXFx1MDAyNjo6YmVmb3JlIHtcXFxcbiAgICAgIGNvbnRlbnQ6ICdcXFxcXFxcXDI1QjgnO1xcXFxuICAgICAgZGlzcGxheTogaW5saW5lLWJsb2NrO1xcXFxuICAgICAgdHJhbnNpdGlvbjogdHJhbnNmb3JtIDEwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tc3VidGxlKTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjZbb3Blbl0gc3VtbWFyeTo6YmVmb3JlIHtcXFxcbiAgICB0cmFuc2Zvcm06IHJvdGF0ZSg5MGRlZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLm5hdi1lbGVtZW50LXRpdGxlIHtcXFxcbiAgdXNlci1zZWxlY3Q6IG5vbmU7XFxcXG59XFxcXG5cXFxcbi5uYXYtZGVtby1saXN0IHtcXFxcbiAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSAwIDAgMDtcXFxcbiAgZGlzcGxheTogZ3JpZDtcXFxcbiAgZ2FwOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKTtcXFxcbn1cXFxcblxcXFxuLm5hdi1kZW1vLWxpbmsge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1wcmltYXJ5KTtcXFxcbiAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuICBwYWRkaW5nOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXNtKSB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLW1kKTtcXFxcbiAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IGNhbGModmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgKiAyKTtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYmctdGVydGlhcnkpO1xcXFxuICBib3JkZXItcmFkaXVzOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItcmFkaXVzKTtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgdHJhbnNpdGlvbjogYmFja2dyb3VuZCAwLjJzLCBjb2xvciAwLjJzO1xcXFxuXFxcXG4gIFxcXFx1MDAyNjpob3ZlciB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tYnJhbmQpO1xcXFxuXFxcXG4gICAgLm5hdi1wYWNrYWdlLW5hbWUge1xcXFxuICAgICAgY29sb3I6IHJnYmEoMjU1LCAyNTUsIDI1NSwgMC44KTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjZbYXJpYS1jdXJyZW50PVxcXFxcXFwicGFnZVxcXFxcXFwiXSB7XFxcXG4gICAgYmFja2dyb3VuZDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYWNjZW50LWNvbG9yKTtcXFxcbiAgICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS10ZXh0LS1jb2xvci0tb24tYnJhbmQpO1xcXFxuXFxcXG4gICAgLm5hdi1wYWNrYWdlLW5hbWUge1xcXFxuICAgICAgY29sb3I6IHJnYmEoMjU1LCAyNTUsIDI1NSwgMC44KTtcXFxcbiAgICB9XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLm5hdi1wYWNrYWdlLW5hbWUge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1zbSk7XFxcXG59XFxcXG5cXFxcbi8qIEluZm8gYnV0dG9uIHBvcG92ZXIgdHJpZ2dlcnMgaW4ga25vYnMgLSBvdmVycmlkZSBwbGFpbiBidXR0b24gcGFkZGluZyAqL1xcXFxucGYtdjYtcG9wb3ZlciBwZi12Ni1idXR0b25bdmFyaWFudD1cXFxcXFxcInBsYWluXFxcXFxcXCJdIHtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tbS1wbGFpbi0tUGFkZGluZ0lubGluZUVuZDogMDtcXFxcbiAgLS1wZi12Ni1jLWJ1dHRvbi0tbS1wbGFpbi0tUGFkZGluZ0lubGluZVN0YXJ0OiAwO1xcXFxuICAtLXBmLXY2LWMtYnV0dG9uLS1NaW5XaWR0aDogYXV0bztcXFxcbn1cXFxcblxcXFxuLyogS25vYiBkZXNjcmlwdGlvbiBjb250ZW50IChzbG90dGVkIGluIGZvcm0gZ3JvdXAgaGVscGVyKSAqL1xcXFxucGYtdjYtZm9ybS1ncm91cCBbc2xvdD1cXFxcXFxcImhlbHBlclxcXFxcXFwiXSB7XFxcXG4gIHAge1xcXFxuICAgIG1hcmdpbjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy14cykgMDtcXFxcbiAgfVxcXFxuXFxcXG4gIGNvZGUge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJnLXRlcnRpYXJ5KTtcXFxcbiAgICBib3JkZXItcmFkaXVzOiAzcHg7XFxcXG4gICAgZm9udC1mYW1pbHk6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtZmFtaWx5LW1vbm8pO1xcXFxuICB9XFxcXG5cXFxcbiAgYSB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gICAgdGV4dC1kZWNvcmF0aW9uOiBub25lO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIHRleHQtZGVjb3JhdGlvbjogdW5kZXJsaW5lO1xcXFxuICAgIH1cXFxcbiAgfVxcXFxuXFxcXG4gIHN0cm9uZyB7XFxcXG4gICAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgfVxcXFxuXFxcXG4gIHVsLCBvbCB7XFxcXG4gICAgbWFyZ2luOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zcGFjaW5nLXhzKSAwO1xcXFxuICAgIHBhZGRpbmctbGVmdDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1sZyk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLyogU3ludGF4IGhpZ2hsaWdodGluZyAoY2hyb21hIC0gdGhlbWFibGUgdmlhIENTUyBjdXN0b20gcHJvcGVydGllcykgKi9cXFxcbnBmLXY2LWZvcm0tZ3JvdXAgW3Nsb3Q9XFxcXFxcXCJoZWxwZXJcXFxcXFxcIl0ge1xcXFxuICAuY2hyb21hIHtcXFxcbiAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy10ZXJ0aWFyeSk7XFxcXG4gICAgcGFkZGluZzogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSk7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXJhZGl1cyk7XFxcXG4gICAgb3ZlcmZsb3cteDogYXV0bztcXFxcblxcXFxuICAgIFxcXFx1MDAyNiAubG50ZCB7IHZlcnRpY2FsLWFsaWduOiB0b3A7IHBhZGRpbmc6IDA7IG1hcmdpbjogMDsgYm9yZGVyOiAwOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5sbnRhYmxlIHsgYm9yZGVyLXNwYWNpbmc6IDA7IHBhZGRpbmc6IDA7IG1hcmdpbjogMDsgYm9yZGVyOiAwOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5obCB7IGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1oaWdobGlnaHQpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmxudCxcXFxcbiAgICBcXFxcdTAwMjYgLmxuIHtcXFxcbiAgICAgIHdoaXRlLXNwYWNlOiBwcmU7XFxcXG4gICAgICB1c2VyLXNlbGVjdDogbm9uZTtcXFxcbiAgICAgIG1hcmdpbi1yaWdodDogMC40ZW07XFxcXG4gICAgICBwYWRkaW5nOiAwIDAuNGVtIDAgMC40ZW07XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAubGluZSB7IGRpc3BsYXk6IGZsZXg7IH1cXFxcblxcXFxuICAgIC8qIEtleXdvcmRzICovXFxcXG4gICAgXFxcXHUwMDI2IC5rLFxcXFxuICAgIFxcXFx1MDAyNiAua2MsXFxcXG4gICAgXFxcXHUwMDI2IC5rZCxcXFxcbiAgICBcXFxcdTAwMjYgLmtuLFxcXFxuICAgIFxcXFx1MDAyNiAua3AsXFxcXG4gICAgXFxcXHUwMDI2IC5rciB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWtleXdvcmQpO1xcXFxuICAgICAgZm9udC13ZWlnaHQ6IGJvbGQ7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAua3QgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LXR5cGUpOyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuXFxcXG4gICAgLyogTmFtZXMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLm5hLFxcXFxuICAgIFxcXFx1MDAyNiAubmIsXFxcXG4gICAgXFxcXHUwMDI2IC5ubyxcXFxcbiAgICBcXFxcdTAwMjYgLm52LFxcXFxuICAgIFxcXFx1MDAyNiAudmMsXFxcXG4gICAgXFxcXHUwMDI2IC52ZyxcXFxcbiAgICBcXFxcdTAwMjYgLnZpIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtbmFtZSk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuYnAgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLm5jIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1jbGFzcyk7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5uZCB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZGVjb3JhdG9yKTsgZm9udC13ZWlnaHQ6IGJvbGQ7IH1cXFxcbiAgICBcXFxcdTAwMjYgLm5pLFxcXFxuICAgIFxcXFx1MDAyNiAuc3Mge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1zcGVjaWFsKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5uZSxcXFxcbiAgICBcXFxcdTAwMjYgLm5sIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgta2V5d29yZCk7XFxcXG4gICAgICBmb250LXdlaWdodDogYm9sZDtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5uZiB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZnVuY3Rpb24pOyBmb250LXdlaWdodDogYm9sZDsgfVxcXFxuICAgIFxcXFx1MDAyNiAubm4geyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLm50IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC10YWcpIH1cXFxcblxcXFxuICAgIC8qIFN0cmluZ3MgKi9cXFxcbiAgICBcXFxcdTAwMjYgLnMsXFxcXG4gICAgXFxcXHUwMDI2IC5zYSxcXFxcbiAgICBcXFxcdTAwMjYgLnNiLFxcXFxuICAgIFxcXFx1MDAyNiAuc2MsXFxcXG4gICAgXFxcXHUwMDI2IC5kbCxcXFxcbiAgICBcXFxcdTAwMjYgLnNkLFxcXFxuICAgIFxcXFx1MDAyNiAuczIsXFxcXG4gICAgXFxcXHUwMDI2IC5zZSxcXFxcbiAgICBcXFxcdTAwMjYgLnNoLFxcXFxuICAgIFxcXFx1MDAyNiAuc2ksXFxcXG4gICAgXFxcXHUwMDI2IC5zeCxcXFxcbiAgICBcXFxcdTAwMjYgLnMxIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtc3RyaW5nKTtcXFxcbiAgICB9XFxcXG4gICAgXFxcXHUwMDI2IC5zciB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtdGFnKSB9XFxcXG5cXFxcbiAgICAvKiBOdW1iZXJzICovXFxcXG4gICAgXFxcXHUwMDI2IC5tLFxcXFxuICAgIFxcXFx1MDAyNiAubWIsXFxcXG4gICAgXFxcXHUwMDI2IC5tZixcXFxcbiAgICBcXFxcdTAwMjYgLm1oLFxcXFxuICAgIFxcXFx1MDAyNiAubWksXFxcXG4gICAgXFxcXHUwMDI2IC5pbCxcXFxcbiAgICBcXFxcdTAwMjYgLm1vIHtcXFxcbiAgICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtbnVtYmVyKTtcXFxcbiAgICB9XFxcXG5cXFxcbiAgICAvKiBPcGVyYXRvcnMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLm8sXFxcXG4gICAgXFxcXHUwMDI2IC5vdyB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWtleXdvcmQpO1xcXFxuICAgICAgZm9udC13ZWlnaHQ6IGJvbGQ7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgLyogQ29tbWVudHMgKi9cXFxcbiAgICBcXFxcdTAwMjYgLmMsXFxcXG4gICAgXFxcXHUwMDI2IC5jaCxcXFxcbiAgICBcXFxcdTAwMjYgLmNtLFxcXFxuICAgIFxcXFx1MDAyNiAuYzEge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICAgICAgZm9udC1zdHlsZTogaXRhbGljO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmNzLFxcXFxuICAgIFxcXFx1MDAyNiAuY3AsXFxcXG4gICAgXFxcXHUwMDI2IC5jcGYge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICAgICAgZm9udC13ZWlnaHQ6IGJvbGQ7XFxcXG4gICAgICBmb250LXN0eWxlOiBpdGFsaWM7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgLyogRXJyb3JzICovXFxcXG4gICAgXFxcXHUwMDI2IC5lcnIge1xcXFxuICAgICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1lcnJvcik7XFxcXG4gICAgICBiYWNrZ3JvdW5kLWNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZXJyb3ItYmcpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIEdlbmVyaWNzICovXFxcXG4gICAgXFxcXHUwMDI2IC5nZCB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWRlbGV0ZWQpO1xcXFxuICAgICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWRlbGV0ZWQtYmcpO1xcXFxuICAgIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdlIHsgZm9udC1zdHlsZTogaXRhbGljOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5nciB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1zeW50YXgtZXJyb3IpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdoIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5naSB7XFxcXG4gICAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWluc2VydGVkKTtcXFxcbiAgICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXN5bnRheC1pbnNlcnRlZC1iZyk7XFxcXG4gICAgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ28geyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpIH1cXFxcbiAgICBcXFxcdTAwMjYgLmdwIHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KSB9XFxcXG4gICAgXFxcXHUwMDI2IC5ncyB7IGZvbnQtd2VpZ2h0OiBib2xkOyB9XFxcXG4gICAgXFxcXHUwMDI2IC5ndSB7IGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSkgfVxcXFxuICAgIFxcXFx1MDAyNiAuZ3QgeyBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3ludGF4LWVycm9yKSB9XFxcXG4gICAgXFxcXHUwMDI2IC5nbCB7IHRleHQtZGVjb3JhdGlvbjogdW5kZXJsaW5lOyB9XFxcXG4gICAgXFxcXHUwMDI2IC53IHsgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpIH1cXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4vKiBFdmVudHMgdGFiIHN0eWxpbmcgLSBQcmltYXJ5LWRldGFpbCBsYXlvdXQgKi9cXFxcbi5ldmVudHMtd3JhcHBlciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGZsZXgtZGlyZWN0aW9uOiBjb2x1bW47XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxuI2V2ZW50LWRyYXdlciB7XFxcXG4gIGZsZXg6IDE7XFxcXG4gIG1pbi1oZWlnaHQ6IDA7XFxcXG59XFxcXG5cXFxcbi8qIEV2ZW50IGxpc3QgKHByaW1hcnkgcGFuZWwpICovXFxcXG4jZXZlbnQtbGlzdCB7XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWxpc3QtaXRlbSB7XFxcXG4gIC8qIFJlc2V0IGJ1dHRvbiBzdHlsZXMgKi9cXFxcbiAgYXBwZWFyYW5jZTogbm9uZTtcXFxcbiAgYmFja2dyb3VuZDogbm9uZTtcXFxcbiAgYm9yZGVyOiBub25lO1xcXFxuICBib3JkZXItbGVmdDogM3B4IHNvbGlkIHRyYW5zcGFyZW50O1xcXFxuICBtYXJnaW46IDA7XFxcXG4gIGZvbnQ6IGluaGVyaXQ7XFxcXG4gIGNvbG9yOiBpbmhlcml0O1xcXFxuICB0ZXh0LWFsaWduOiBpbmhlcml0O1xcXFxuICB3aWR0aDogMTAwJTtcXFxcblxcXFxuICAvKiBDb21wb25lbnQgc3R5bGVzICovXFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBnYXA6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pO1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBjdXJzb3I6IHBvaW50ZXI7XFxcXG4gIHRyYW5zaXRpb246IGJhY2tncm91bmQgMTAwbXMgZWFzZS1pbi1vdXQsIGJvcmRlci1jb2xvciAxMDBtcyBlYXNlLWluLW91dDtcXFxcblxcXFxuICBwZi12Ni1sYWJlbCB7XFxcXG4gICAgZmxleC1zaHJpbms6IDA7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXByaW1hcnktLWhvdmVyKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNjpmb2N1cyB7XFxcXG4gICAgb3V0bGluZTogMnB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1jb2xvci0tY2xpY2tlZCk7XFxcXG4gICAgb3V0bGluZS1vZmZzZXQ6IC0ycHg7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYuc2VsZWN0ZWQge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLWFjdGlvbi0tcGxhaW4tLXNlbGVjdGVkKTtcXFxcbiAgICBib3JkZXItbGVmdC1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1icmFuZC0tZGVmYXVsdCk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmV2ZW50LXRpbWUsXFxcXG4uZXZlbnQtZWxlbWVudCB7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtdGltZSB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgZmxleC1zaHJpbms6IDA7XFxcXG4gIGZvbnQtc2l6ZTogMTFweDtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWVsZW1lbnQge1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1tdXRlZCk7XFxcXG4gIGZvbnQtd2VpZ2h0OiA0MDA7XFxcXG59XFxcXG5cXFxcbi8qIEV2ZW50IGRldGFpbCBwYW5lbCAqL1xcXFxuLmV2ZW50LWRldGFpbC1oZWFkZXItY29udGVudCB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBib3JkZXItYm90dG9tOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1ib3JkZXItd2lkdGgpIHNvbGlkIHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci1jb2xvcik7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtbmFtZSB7XFxcXG4gIG1hcmdpbjogMCAwIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIDA7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLWxnKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtc3VtbWFyeSB7XFxcXG4gIG1hcmdpbjogMCAwIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIDA7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgbGluZS1oZWlnaHQ6IDEuNTtcXFxcbiAgd2hpdGUtc3BhY2U6IHByZS13cmFwO1xcXFxufVxcXFxuXFxcXG4uZXZlbnQtZGV0YWlsLWRlc2NyaXB0aW9uIHtcXFxcbiAgbWFyZ2luOiAwIDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1zbSkgMDtcXFxcbiAgZm9udC1zaXplOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LXNpemUtc20pO1xcXFxuICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItdGV4dC1zZWNvbmRhcnkpO1xcXFxuICBsaW5lLWhlaWdodDogMS41O1xcXFxuICB3aGl0ZS1zcGFjZTogcHJlLXdyYXA7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtbWV0YSB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGdhcDogdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1zaXplLXNtKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC10aW1lIHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtbXV0ZWQpO1xcXFxuICBmb250LWZhbWlseTogdmFyKC0tY2VtLWRldi1zZXJ2ZXItZm9udC1mYW1pbHktbW9ubyk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtZWxlbWVudCB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LXNlY29uZGFyeSk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbn1cXFxcblxcXFxuLmV2ZW50LWRldGFpbC1wcm9wZXJ0aWVzLWhlYWRpbmcge1xcXFxuICBtYXJnaW46IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBmb250LXNpemU6IHZhcigtLWNlbS1kZXYtc2VydmVyLWZvbnQtc2l6ZS1iYXNlKTtcXFxcbiAgZm9udC13ZWlnaHQ6IDYwMDtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtcHJpbWFyeSk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1kZXRhaWwtcHJvcGVydGllcyB7XFxcXG4gIHBhZGRpbmc6IHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctc20pIHZhcigtLWNlbS1kZXYtc2VydmVyLXNwYWNpbmctbWQpO1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1jZW0tZGV2LXNlcnZlci1iZy1zZWNvbmRhcnkpO1xcXFxuICBib3JkZXI6IHZhcigtLWNlbS1kZXYtc2VydmVyLWJvcmRlci13aWR0aCkgc29saWQgdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLWNvbG9yKTtcXFxcbiAgYm9yZGVyLXJhZGl1czogdmFyKC0tY2VtLWRldi1zZXJ2ZXItYm9yZGVyLXJhZGl1cyk7XFxcXG4gIGZvbnQtZmFtaWx5OiB2YXIoLS1jZW0tZGV2LXNlcnZlci1mb250LWZhbWlseS1tb25vKTtcXFxcbiAgZm9udC1zaXplOiAxMnB4O1xcXFxuICBsaW5lLWhlaWdodDogMS42O1xcXFxuICBtYXJnaW46IDAgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCkgdmFyKC0tY2VtLWRldi1zZXJ2ZXItc3BhY2luZy1tZCk7XFxcXG59XFxcXG5cXFxcbi5ldmVudC1wcm9wZXJ0eS10cmVlIHtcXFxcbiAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuXFxcXG4gIFxcXFx1MDAyNi5uZXN0ZWQge1xcXFxuICAgIHBhZGRpbmctbGVmdDogMS41ZW07XFxcXG4gICAgbWFyZ2luLXRvcDogMC4yNWVtO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi5wcm9wZXJ0eS1pdGVtIHtcXFxcbiAgcGFkZGluZzogMC4xMjVlbSAwO1xcXFxufVxcXFxuXFxcXG4ucHJvcGVydHkta2V5IHtcXFxcbiAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWFjY2VudC1jb2xvcik7XFxcXG4gIGZvbnQtd2VpZ2h0OiA1MDA7XFxcXG59XFxcXG5cXFxcbi5wcm9wZXJ0eS1jb2xvbiB7XFxcXG4gIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbn1cXFxcblxcXFxuLnByb3BlcnR5LXZhbHVlIHtcXFxcbiAgXFxcXHUwMDI2Lm51bGwsXFxcXG4gIFxcXFx1MDAyNi51bmRlZmluZWQge1xcXFxuICAgIGNvbG9yOiB2YXIoLS1jZW0tZGV2LXNlcnZlci10ZXh0LW11dGVkKTtcXFxcbiAgICBmb250LXN0eWxlOiBpdGFsaWM7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYuYm9vbGVhbiB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWNvbG9yLWJvb2xlYW4pO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2Lm51bWJlciB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLWNvbG9yLW51bWJlcik7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYuc3RyaW5nIHtcXFxcbiAgICBjb2xvcjogdmFyKC0tY2VtLWRldi1zZXJ2ZXItY29sb3Itc3RyaW5nKTtcXFxcbiAgfVxcXFxuXFxcXG4gIFxcXFx1MDAyNi5hcnJheSxcXFxcbiAgXFxcXHUwMDI2Lm9iamVjdCB7XFxcXG4gICAgY29sb3I6IHZhcigtLWNlbS1kZXYtc2VydmVyLXRleHQtc2Vjb25kYXJ5KTtcXFxcbiAgfVxcXFxufVxcXFxuXFxcXG4jZGVidWctbW9kYWwge1xcXFxuICBjb250YWluZXItdHlwZTogaW5saW5lLXNpemU7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxNQUFNLGVBQWU7QUFDMUMsU0FBUyxxQkFBcUI7QUFDOUIsU0FBUyxnQkFBZ0I7OztBQ0Z6QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFB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
