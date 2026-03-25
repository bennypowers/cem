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

// elements/cem-health-panel/cem-health-panel.ts
import { LitElement, html, nothing } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";
import { property } from "/__cem/vendor/lit/decorators/property.js";
import { state } from "/__cem/vendor/lit/decorators/state.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-health-panel/cem-health-panel.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n  padding: var(--pf-t--global--spacer--md);\\n}\\n\\n#overall {\\n  margin-block-end: var(--pf-t--global--spacer--md);\\n}\\n\\n.category-bar {\\n  display: flex;\\n  align-items: center;\\n  gap: var(--pf-t--global--spacer--sm);\\n}\\n\\n.category-meter {\\n  flex: 1;\\n  height: 8px;\\n  border-radius: var(--pf-t--global--border--radius--pill, 999px);\\n  background: var(--pf-t--global--background--color--secondary--default, #f0f0f0);\\n  overflow: hidden;\\n\\n  \\u0026 \\u003e .category-fill {\\n    height: 100%;\\n    border-radius: inherit;\\n    transition: width 0.3s ease;\\n  }\\n}\\n\\n.fill-pass {\\n  background: var(--pf-t--global--color--status--success--default, #3e8635);\\n}\\n\\n.fill-warn {\\n  background: var(--pf-t--global--color--status--warning--default, #f0ab00);\\n}\\n\\n.fill-fail {\\n  background: var(--pf-t--global--color--status--danger--default, #c9190b);\\n}\\n\\n.category-score {\\n  min-width: 4ch;\\n  text-align: end;\\n  font-variant-numeric: tabular-nums;\\n}\\n\\n#recommendations {\\n  margin-block-start: var(--pf-t--global--spacer--md);\\n  border-block-start: 1px solid var(--pf-t--global--border--color--default, #d2d2d2);\\n  padding-block-start: var(--pf-t--global--spacer--md);\\n\\n  \\u0026 h4 {\\n    margin: 0 0 var(--pf-t--global--spacer--sm);\\n    font-size: var(--pf-t--global--font--size--body--default, 1rem);\\n    font-weight: var(--pf-t--global--font--weight--body--bold, 700);\\n  }\\n\\n  \\u0026 ul {\\n    margin: 0;\\n    padding-inline-start: var(--pf-t--global--spacer--lg);\\n  }\\n\\n  \\u0026 li {\\n    margin-block-end: var(--pf-t--global--spacer--xs);\\n    font-size: var(--pf-t--global--font--size--body--sm, 0.875rem);\\n  }\\n}\\n\\n.finding-details {\\n  font-size: var(--pf-t--global--font--size--body--sm, 0.875rem);\\n  padding-block-start: var(--pf-t--global--spacer--xs);\\n\\n  \\u0026 ul {\\n    margin: 0;\\n    padding-inline-start: var(--pf-t--global--spacer--lg);\\n    list-style: none;\\n  }\\n\\n  \\u0026 li::before {\\n    content: \\"\\\\2192 \\";\\n  }\\n}\\n"'));
var cem_health_panel_default = s;

// elements/cem-health-panel/cem-health-panel.ts
import "../pf-v6-label/pf-v6-label.js";
import "../pf-v6-expandable-section/pf-v6-expandable-section.js";
var STATUS_COLORS = {
  pass: "green",
  warn: "orange",
  fail: "red"
};
var _result_dec, _loading_dec, _component_dec, _a, _CemHealthPanel_decorators, _init, _component, _loading, _b, loading_get, loading_set, _CemHealthPanel_instances, _result, _c, result_get, result_set, _abortController, findDeclaration_fn, renderCategory_fn, renderRecommendations_fn;
_CemHealthPanel_decorators = [customElement("cem-health-panel")];
var CemHealthPanel = class extends (_a = LitElement, _component_dec = [property({ reflect: true })], _loading_dec = [state()], _result_dec = [state()], _a) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemHealthPanel_instances);
    __privateAdd(this, _component, __runInitializers(_init, 8, this, null)), __runInitializers(_init, 11, this);
    __privateAdd(this, _loading, __runInitializers(_init, 12, this, true)), __runInitializers(_init, 15, this);
    __privateAdd(this, _result, __runInitializers(_init, 16, this, null)), __runInitializers(_init, 19, this);
    __privateAdd(this, _abortController, null);
  }
  connectedCallback() {
    super.connectedCallback();
    this.fetchHealth();
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(this, _abortController)?.abort();
  }
  async fetchHealth() {
    __privateGet(this, _abortController)?.abort();
    __privateSet(this, _abortController, new AbortController());
    __privateSet(this, _CemHealthPanel_instances, true, loading_set);
    __privateSet(this, _CemHealthPanel_instances, null, result_set);
    const url = new URL("/__cem/api/health", location.origin);
    if (this.component) {
      url.searchParams.set("component", this.component);
    }
    try {
      const response = await fetch(url, { signal: __privateGet(this, _abortController).signal });
      if (!response.ok) throw new Error(`HTTP ${response.status}`);
      __privateSet(this, _CemHealthPanel_instances, await response.json(), result_set);
    } catch (e) {
      if (e.name === "AbortError") return;
      console.warn("cem-health-panel: failed to fetch health data", e);
    } finally {
      __privateSet(this, _CemHealthPanel_instances, false, loading_set);
    }
  }
  render() {
    if (__privateGet(this, _CemHealthPanel_instances, loading_get)) {
      return html`<p>Analyzing documentation health...</p>`;
    }
    const decl = __privateMethod(this, _CemHealthPanel_instances, findDeclaration_fn).call(this);
    if (!decl) {
      return html`<p>No health data available.</p>`;
    }
    const pct = decl.maxScore > 0 ? Math.round(decl.score / decl.maxScore * 100) : 0;
    const overallStatus = pct >= 80 ? "pass" : pct >= 40 ? "warn" : "fail";
    return html`
      <div id="overall">
        <pf-v6-label color=${STATUS_COLORS[overallStatus]}
                     size="lg">${pct}% -- ${decl.score}/${decl.maxScore}</pf-v6-label>
      </div>
      <dl id="categories"
          class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        ${decl.categories.map((cat) => __privateMethod(this, _CemHealthPanel_instances, renderCategory_fn).call(this, cat))}
      </dl>
      ${__privateMethod(this, _CemHealthPanel_instances, renderRecommendations_fn).call(this)}
    `;
  }
};
_init = __decoratorStart(_a);
_component = new WeakMap();
_loading = new WeakMap();
_CemHealthPanel_instances = new WeakSet();
_result = new WeakMap();
_abortController = new WeakMap();
findDeclaration_fn = function() {
  if (!__privateGet(this, _CemHealthPanel_instances, result_get)) return null;
  for (const mod of __privateGet(this, _CemHealthPanel_instances, result_get).modules) {
    for (const decl of mod.declarations) {
      if (!this.component || decl.tagName === this.component || decl.name === this.component) {
        return decl;
      }
    }
  }
  return null;
};
renderCategory_fn = function(cat) {
  const pct = cat.maxPoints > 0 ? Math.round(cat.points / cat.maxPoints * 100) : 0;
  const findings = cat.findings?.filter((f) => f.message) ?? [];
  return html`
      <div class="pf-v6-c-description-list__group">
        <dt class="pf-v6-c-description-list__term">${cat.category}</dt>
        <dd class="pf-v6-c-description-list__description">
          <div class="category-bar">
            <div class="category-meter">
              <div class="category-fill fill-${cat.status}"
                   style="width: ${pct}%"></div>
            </div>
            <span class="category-score">${cat.points}/${cat.maxPoints}</span>
            <pf-v6-label color=${STATUS_COLORS[cat.status]}>${cat.status}</pf-v6-label>
          </div>
          ${findings.length > 0 ? html`
            <div class="finding-details">
              <ul>
                ${findings.map((f) => html`<li>${f.message}</li>`)}
              </ul>
            </div>
          ` : nothing}
        </dd>
      </div>
    `;
};
renderRecommendations_fn = function() {
  const recs = __privateGet(this, _CemHealthPanel_instances, result_get)?.recommendations;
  if (!recs?.length) return nothing;
  return html`
      <div id="recommendations">
        <h4>Recommendations</h4>
        <ul>
          ${recs.map((rec) => html`<li>${rec}</li>`)}
        </ul>
      </div>
    `;
};
__decorateElement(_init, 4, "component", _component_dec, CemHealthPanel, _component);
_b = __decorateElement(_init, 20, "#loading", _loading_dec, _CemHealthPanel_instances, _loading), loading_get = _b.get, loading_set = _b.set;
_c = __decorateElement(_init, 20, "#result", _result_dec, _CemHealthPanel_instances, _result), result_get = _c.get, result_set = _c.set;
CemHealthPanel = __decorateElement(_init, 0, "CemHealthPanel", _CemHealthPanel_decorators, CemHealthPanel);
__publicField(CemHealthPanel, "styles", cem_health_panel_default);
__runInitializers(_init, 1, CemHealthPanel);
export {
  CemHealthPanel
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLWhlYWx0aC1wYW5lbC9jZW0taGVhbHRoLXBhbmVsLnRzIiwgImxpdC1jc3M6L2hvbWUvYmVubnlwL0RldmVsb3Blci9jZW0vc2VydmUvZWxlbWVudHMvY2VtLWhlYWx0aC1wYW5lbC9jZW0taGVhbHRoLXBhbmVsLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1oZWFsdGgtcGFuZWwuY3NzJztcblxuaW1wb3J0ICcuLi9wZi12Ni1sYWJlbC9wZi12Ni1sYWJlbC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LWV4cGFuZGFibGUtc2VjdGlvbi9wZi12Ni1leHBhbmRhYmxlLXNlY3Rpb24uanMnO1xuXG5jb25zdCBTVEFUVVNfQ09MT1JTOiBSZWNvcmQ8c3RyaW5nLCBzdHJpbmc+ID0ge1xuICBwYXNzOiAnZ3JlZW4nLFxuICB3YXJuOiAnb3JhbmdlJyxcbiAgZmFpbDogJ3JlZCcsXG59O1xuXG5pbnRlcmZhY2UgSGVhbHRoRmluZGluZyB7XG4gIG1lc3NhZ2U/OiBzdHJpbmc7XG59XG5cbmludGVyZmFjZSBIZWFsdGhDYXRlZ29yeSB7XG4gIGNhdGVnb3J5OiBzdHJpbmc7XG4gIHBvaW50czogbnVtYmVyO1xuICBtYXhQb2ludHM6IG51bWJlcjtcbiAgc3RhdHVzOiBzdHJpbmc7XG4gIGZpbmRpbmdzPzogSGVhbHRoRmluZGluZ1tdO1xufVxuXG5pbnRlcmZhY2UgSGVhbHRoRGVjbGFyYXRpb24ge1xuICB0YWdOYW1lPzogc3RyaW5nO1xuICBuYW1lPzogc3RyaW5nO1xuICBzY29yZTogbnVtYmVyO1xuICBtYXhTY29yZTogbnVtYmVyO1xuICBjYXRlZ29yaWVzOiBIZWFsdGhDYXRlZ29yeVtdO1xufVxuXG5pbnRlcmZhY2UgSGVhbHRoTW9kdWxlIHtcbiAgZGVjbGFyYXRpb25zOiBIZWFsdGhEZWNsYXJhdGlvbltdO1xufVxuXG5pbnRlcmZhY2UgSGVhbHRoUmVzdWx0IHtcbiAgbW9kdWxlczogSGVhbHRoTW9kdWxlW107XG4gIHJlY29tbWVuZGF0aW9ucz86IHN0cmluZ1tdO1xufVxuXG4vKipcbiAqIEhlYWx0aCBwYW5lbCBzaG93aW5nIGRvY3VtZW50YXRpb24gcXVhbGl0eSBzY29yZXNcbiAqXG4gKiBAYXR0ciB7c3RyaW5nfSBjb21wb25lbnQgLSBDb21wb25lbnQgdGFnIG5hbWUgb3IgY2xhc3MgbmFtZSB0byBkaXNwbGF5IGhlYWx0aCBmb3JcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1oZWFsdGgtcGFuZWwnKVxuZXhwb3J0IGNsYXNzIENlbUhlYWx0aFBhbmVsIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgQHByb3BlcnR5KHsgcmVmbGVjdDogdHJ1ZSB9KVxuICBhY2Nlc3NvciBjb21wb25lbnQ6IHN0cmluZyB8IG51bGwgPSBudWxsO1xuXG4gIEBzdGF0ZSgpXG4gIGFjY2Vzc29yICNsb2FkaW5nID0gdHJ1ZTtcblxuICBAc3RhdGUoKVxuICBhY2Nlc3NvciAjcmVzdWx0OiBIZWFsdGhSZXN1bHQgfCBudWxsID0gbnVsbDtcblxuICAjYWJvcnRDb250cm9sbGVyOiBBYm9ydENvbnRyb2xsZXIgfCBudWxsID0gbnVsbDtcblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuZmV0Y2hIZWFsdGgoKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy4jYWJvcnRDb250cm9sbGVyPy5hYm9ydCgpO1xuICB9XG5cbiAgYXN5bmMgZmV0Y2hIZWFsdGgoKSB7XG4gICAgdGhpcy4jYWJvcnRDb250cm9sbGVyPy5hYm9ydCgpO1xuICAgIHRoaXMuI2Fib3J0Q29udHJvbGxlciA9IG5ldyBBYm9ydENvbnRyb2xsZXIoKTtcblxuICAgIHRoaXMuI2xvYWRpbmcgPSB0cnVlO1xuICAgIHRoaXMuI3Jlc3VsdCA9IG51bGw7XG5cbiAgICBjb25zdCB1cmwgPSBuZXcgVVJMKCcvX19jZW0vYXBpL2hlYWx0aCcsIGxvY2F0aW9uLm9yaWdpbik7XG4gICAgaWYgKHRoaXMuY29tcG9uZW50KSB7XG4gICAgICB1cmwuc2VhcmNoUGFyYW1zLnNldCgnY29tcG9uZW50JywgdGhpcy5jb21wb25lbnQpO1xuICAgIH1cblxuICAgIHRyeSB7XG4gICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IGZldGNoKHVybCwgeyBzaWduYWw6IHRoaXMuI2Fib3J0Q29udHJvbGxlci5zaWduYWwgfSk7XG4gICAgICBpZiAoIXJlc3BvbnNlLm9rKSB0aHJvdyBuZXcgRXJyb3IoYEhUVFAgJHtyZXNwb25zZS5zdGF0dXN9YCk7XG4gICAgICB0aGlzLiNyZXN1bHQgPSBhd2FpdCByZXNwb25zZS5qc29uKCk7XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgaWYgKChlIGFzIEVycm9yKS5uYW1lID09PSAnQWJvcnRFcnJvcicpIHJldHVybjtcbiAgICAgIGNvbnNvbGUud2FybignY2VtLWhlYWx0aC1wYW5lbDogZmFpbGVkIHRvIGZldGNoIGhlYWx0aCBkYXRhJywgZSk7XG4gICAgfSBmaW5hbGx5IHtcbiAgICAgIHRoaXMuI2xvYWRpbmcgPSBmYWxzZTtcbiAgICB9XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgaWYgKHRoaXMuI2xvYWRpbmcpIHtcbiAgICAgIHJldHVybiBodG1sYDxwPkFuYWx5emluZyBkb2N1bWVudGF0aW9uIGhlYWx0aC4uLjwvcD5gO1xuICAgIH1cblxuICAgIGNvbnN0IGRlY2wgPSB0aGlzLiNmaW5kRGVjbGFyYXRpb24oKTtcbiAgICBpZiAoIWRlY2wpIHtcbiAgICAgIHJldHVybiBodG1sYDxwPk5vIGhlYWx0aCBkYXRhIGF2YWlsYWJsZS48L3A+YDtcbiAgICB9XG5cbiAgICBjb25zdCBwY3QgPSBkZWNsLm1heFNjb3JlID4gMFxuICAgICAgPyBNYXRoLnJvdW5kKChkZWNsLnNjb3JlIC8gZGVjbC5tYXhTY29yZSkgKiAxMDApXG4gICAgICA6IDA7XG4gICAgY29uc3Qgb3ZlcmFsbFN0YXR1cyA9IHBjdCA+PSA4MCA/ICdwYXNzJyA6IHBjdCA+PSA0MCA/ICd3YXJuJyA6ICdmYWlsJztcblxuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBpZD1cIm92ZXJhbGxcIj5cbiAgICAgICAgPHBmLXY2LWxhYmVsIGNvbG9yPSR7U1RBVFVTX0NPTE9SU1tvdmVyYWxsU3RhdHVzXX1cbiAgICAgICAgICAgICAgICAgICAgIHNpemU9XCJsZ1wiPiR7cGN0fSUgLS0gJHtkZWNsLnNjb3JlfS8ke2RlY2wubWF4U2NvcmV9PC9wZi12Ni1sYWJlbD5cbiAgICAgIDwvZGl2PlxuICAgICAgPGRsIGlkPVwiY2F0ZWdvcmllc1wiXG4gICAgICAgICAgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3QgcGYtbS1ob3Jpem9udGFsIHBmLW0tY29tcGFjdFwiPlxuICAgICAgICAke2RlY2wuY2F0ZWdvcmllcy5tYXAoY2F0ID0+IHRoaXMuI3JlbmRlckNhdGVnb3J5KGNhdCkpfVxuICAgICAgPC9kbD5cbiAgICAgICR7dGhpcy4jcmVuZGVyUmVjb21tZW5kYXRpb25zKCl9XG4gICAgYDtcbiAgfVxuXG4gICNmaW5kRGVjbGFyYXRpb24oKTogSGVhbHRoRGVjbGFyYXRpb24gfCBudWxsIHtcbiAgICBpZiAoIXRoaXMuI3Jlc3VsdCkgcmV0dXJuIG51bGw7XG4gICAgZm9yIChjb25zdCBtb2Qgb2YgdGhpcy4jcmVzdWx0Lm1vZHVsZXMpIHtcbiAgICAgIGZvciAoY29uc3QgZGVjbCBvZiBtb2QuZGVjbGFyYXRpb25zKSB7XG4gICAgICAgIGlmICghdGhpcy5jb21wb25lbnRcbiAgICAgICAgICAgIHx8IGRlY2wudGFnTmFtZSA9PT0gdGhpcy5jb21wb25lbnRcbiAgICAgICAgICAgIHx8IGRlY2wubmFtZSA9PT0gdGhpcy5jb21wb25lbnQpIHtcbiAgICAgICAgICByZXR1cm4gZGVjbDtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cbiAgICByZXR1cm4gbnVsbDtcbiAgfVxuXG4gICNyZW5kZXJDYXRlZ29yeShjYXQ6IEhlYWx0aENhdGVnb3J5KSB7XG4gICAgY29uc3QgcGN0ID0gY2F0Lm1heFBvaW50cyA+IDBcbiAgICAgID8gTWF0aC5yb3VuZCgoY2F0LnBvaW50cyAvIGNhdC5tYXhQb2ludHMpICogMTAwKVxuICAgICAgOiAwO1xuICAgIGNvbnN0IGZpbmRpbmdzID0gY2F0LmZpbmRpbmdzPy5maWx0ZXIoZiA9PiBmLm1lc3NhZ2UpID8/IFtdO1xuXG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8ZGl2IGNsYXNzPVwicGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19ncm91cFwiPlxuICAgICAgICA8ZHQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj4ke2NhdC5jYXRlZ29yeX08L2R0PlxuICAgICAgICA8ZGQgY2xhc3M9XCJwZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2Rlc2NyaXB0aW9uXCI+XG4gICAgICAgICAgPGRpdiBjbGFzcz1cImNhdGVnb3J5LWJhclwiPlxuICAgICAgICAgICAgPGRpdiBjbGFzcz1cImNhdGVnb3J5LW1ldGVyXCI+XG4gICAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJjYXRlZ29yeS1maWxsIGZpbGwtJHtjYXQuc3RhdHVzfVwiXG4gICAgICAgICAgICAgICAgICAgc3R5bGU9XCJ3aWR0aDogJHtwY3R9JVwiPjwvZGl2PlxuICAgICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgICA8c3BhbiBjbGFzcz1cImNhdGVnb3J5LXNjb3JlXCI+JHtjYXQucG9pbnRzfS8ke2NhdC5tYXhQb2ludHN9PC9zcGFuPlxuICAgICAgICAgICAgPHBmLXY2LWxhYmVsIGNvbG9yPSR7U1RBVFVTX0NPTE9SU1tjYXQuc3RhdHVzXX0+JHtjYXQuc3RhdHVzfTwvcGYtdjYtbGFiZWw+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgJHtmaW5kaW5ncy5sZW5ndGggPiAwID8gaHRtbGBcbiAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJmaW5kaW5nLWRldGFpbHNcIj5cbiAgICAgICAgICAgICAgPHVsPlxuICAgICAgICAgICAgICAgICR7ZmluZGluZ3MubWFwKGYgPT4gaHRtbGA8bGk+JHtmLm1lc3NhZ2V9PC9saT5gKX1cbiAgICAgICAgICAgICAgPC91bD5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIGAgOiBub3RoaW5nfVxuICAgICAgICA8L2RkPlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gICNyZW5kZXJSZWNvbW1lbmRhdGlvbnMoKSB7XG4gICAgY29uc3QgcmVjcyA9IHRoaXMuI3Jlc3VsdD8ucmVjb21tZW5kYXRpb25zO1xuICAgIGlmICghcmVjcz8ubGVuZ3RoKSByZXR1cm4gbm90aGluZztcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJyZWNvbW1lbmRhdGlvbnNcIj5cbiAgICAgICAgPGg0PlJlY29tbWVuZGF0aW9uczwvaDQ+XG4gICAgICAgIDx1bD5cbiAgICAgICAgICAke3JlY3MubWFwKHJlYyA9PiBodG1sYDxsaT4ke3JlY308L2xpPmApfVxuICAgICAgICA8L3VsPlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0taGVhbHRoLXBhbmVsJzogQ2VtSGVhbHRoUGFuZWw7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCk7XFxcXG59XFxcXG5cXFxcbiNvdmVyYWxsIHtcXFxcbiAgbWFyZ2luLWJsb2NrLWVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbn1cXFxcblxcXFxuLmNhdGVnb3J5LWJhciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbn1cXFxcblxcXFxuLmNhdGVnb3J5LW1ldGVyIHtcXFxcbiAgZmxleDogMTtcXFxcbiAgaGVpZ2h0OiA4cHg7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXBpbGwsIDk5OXB4KTtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0LCAjZjBmMGYwKTtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcblxcXFxuICBcXFxcdTAwMjYgXFxcXHUwMDNlIC5jYXRlZ29yeS1maWxsIHtcXFxcbiAgICBoZWlnaHQ6IDEwMCU7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogaW5oZXJpdDtcXFxcbiAgICB0cmFuc2l0aW9uOiB3aWR0aCAwLjNzIGVhc2U7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmZpbGwtcGFzcyB7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLXN0YXR1cy0tc3VjY2Vzcy0tZGVmYXVsdCwgIzNlODYzNSk7XFxcXG59XFxcXG5cXFxcbi5maWxsLXdhcm4ge1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1zdGF0dXMtLXdhcm5pbmctLWRlZmF1bHQsICNmMGFiMDApO1xcXFxufVxcXFxuXFxcXG4uZmlsbC1mYWlsIHtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tc3RhdHVzLS1kYW5nZXItLWRlZmF1bHQsICNjOTE5MGIpO1xcXFxufVxcXFxuXFxcXG4uY2F0ZWdvcnktc2NvcmUge1xcXFxuICBtaW4td2lkdGg6IDRjaDtcXFxcbiAgdGV4dC1hbGlnbjogZW5kO1xcXFxuICBmb250LXZhcmlhbnQtbnVtZXJpYzogdGFidWxhci1udW1zO1xcXFxufVxcXFxuXFxcXG4jcmVjb21tZW5kYXRpb25zIHtcXFxcbiAgbWFyZ2luLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICBib3JkZXItYmxvY2stc3RhcnQ6IDFweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQsICNkMmQyZDIpO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNiBoNCB7XFxcXG4gICAgbWFyZ2luOiAwIDAgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tZGVmYXVsdCwgMXJlbSk7XFxcXG4gICAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0td2VpZ2h0LS1ib2R5LS1ib2xkLCA3MDApO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2IHVsIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gICAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYgbGkge1xcXFxuICAgIG1hcmdpbi1ibG9jay1lbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS14cyk7XFxcXG4gICAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtLCAwLjg3NXJlbSk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmZpbmRpbmctZGV0YWlscyB7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSwgMC44NzVyZW0pO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teHMpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNiB1bCB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICAgIHBhZGRpbmctaW5saW5lLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAgIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYgbGk6OmJlZm9yZSB7XFxcXG4gICAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcXDIxOTIgXFxcXFxcXCI7XFxcXG4gIH1cXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLE1BQU0sZUFBZTtBQUMxQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjtBQUN6QixTQUFTLGFBQWE7OztBQ0h0QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxtakVBQXVqRSxDQUFDO0FBQUUsSUFBTywyQkFBUTs7O0FETzluRSxPQUFPO0FBQ1AsT0FBTztBQUVQLElBQU0sZ0JBQXdDO0FBQUEsRUFDNUMsTUFBTTtBQUFBLEVBQ04sTUFBTTtBQUFBLEVBQ04sTUFBTTtBQUNSO0FBZEE7QUFrREEsOEJBQUMsY0FBYyxrQkFBa0I7QUFDMUIsSUFBTSxpQkFBTixlQUE2QixpQkFHbEMsa0JBQUMsU0FBUyxFQUFFLFNBQVMsS0FBSyxDQUFDLElBRzNCLGdCQUFDLE1BQU0sSUFHUCxlQUFDLE1BQU0sSUFUMkIsSUFBVztBQUFBLEVBQXhDO0FBQUE7QUFBQTtBQUlMLHVCQUFTLFlBQTJCLGtCQUFwQyxnQkFBb0MsUUFBcEM7QUFHQSx1QkFBUyxVQUFXLGtCQUFwQixpQkFBb0IsUUFBcEI7QUFHQSx1QkFBUyxTQUErQixrQkFBeEMsaUJBQXdDLFFBQXhDO0FBRUEseUNBQTJDO0FBQUE7QUFBQSxFQUUzQyxvQkFBb0I7QUFDbEIsVUFBTSxrQkFBa0I7QUFDeEIsU0FBSyxZQUFZO0FBQUEsRUFDbkI7QUFBQSxFQUVBLHVCQUF1QjtBQUNyQixVQUFNLHFCQUFxQjtBQUMzQix1QkFBSyxtQkFBa0IsTUFBTTtBQUFBLEVBQy9CO0FBQUEsRUFFQSxNQUFNLGNBQWM7QUFDbEIsdUJBQUssbUJBQWtCLE1BQU07QUFDN0IsdUJBQUssa0JBQW1CLElBQUksZ0JBQWdCO0FBRTVDLHVCQUFLLDJCQUFXLE1BQVg7QUFDTCx1QkFBSywyQkFBVSxNQUFWO0FBRUwsVUFBTSxNQUFNLElBQUksSUFBSSxxQkFBcUIsU0FBUyxNQUFNO0FBQ3hELFFBQUksS0FBSyxXQUFXO0FBQ2xCLFVBQUksYUFBYSxJQUFJLGFBQWEsS0FBSyxTQUFTO0FBQUEsSUFDbEQ7QUFFQSxRQUFJO0FBQ0YsWUFBTSxXQUFXLE1BQU0sTUFBTSxLQUFLLEVBQUUsUUFBUSxtQkFBSyxrQkFBaUIsT0FBTyxDQUFDO0FBQzFFLFVBQUksQ0FBQyxTQUFTLEdBQUksT0FBTSxJQUFJLE1BQU0sUUFBUSxTQUFTLE1BQU0sRUFBRTtBQUMzRCx5QkFBSywyQkFBVSxNQUFNLFNBQVMsS0FBSyxHQUE5QjtBQUFBLElBQ1AsU0FBUyxHQUFHO0FBQ1YsVUFBSyxFQUFZLFNBQVMsYUFBYztBQUN4QyxjQUFRLEtBQUssaURBQWlELENBQUM7QUFBQSxJQUNqRSxVQUFFO0FBQ0EseUJBQUssMkJBQVcsT0FBWDtBQUFBLElBQ1A7QUFBQSxFQUNGO0FBQUEsRUFFQSxTQUFTO0FBQ1AsUUFBSSxtQkFBSyx5Q0FBVTtBQUNqQixhQUFPO0FBQUEsSUFDVDtBQUVBLFVBQU0sT0FBTyxzQkFBSywrQ0FBTDtBQUNiLFFBQUksQ0FBQyxNQUFNO0FBQ1QsYUFBTztBQUFBLElBQ1Q7QUFFQSxVQUFNLE1BQU0sS0FBSyxXQUFXLElBQ3hCLEtBQUssTUFBTyxLQUFLLFFBQVEsS0FBSyxXQUFZLEdBQUcsSUFDN0M7QUFDSixVQUFNLGdCQUFnQixPQUFPLEtBQUssU0FBUyxPQUFPLEtBQUssU0FBUztBQUVoRSxXQUFPO0FBQUE7QUFBQSw2QkFFa0IsY0FBYyxhQUFhLENBQUM7QUFBQSxpQ0FDeEIsR0FBRyxRQUFRLEtBQUssS0FBSyxJQUFJLEtBQUssUUFBUTtBQUFBO0FBQUE7QUFBQTtBQUFBLFVBSTdELEtBQUssV0FBVyxJQUFJLFNBQU8sc0JBQUssOENBQUwsV0FBcUIsSUFBSSxDQUFDO0FBQUE7QUFBQSxRQUV2RCxzQkFBSyxxREFBTCxVQUE2QjtBQUFBO0FBQUEsRUFFbkM7QUEwREY7QUFwSU87QUFJSTtBQUdBO0FBUEo7QUFVSTtBQUVUO0FBZ0VBLHFCQUFnQixXQUE2QjtBQUMzQyxNQUFJLENBQUMsbUJBQUssdUNBQVMsUUFBTztBQUMxQixhQUFXLE9BQU8sbUJBQUssdUNBQVEsU0FBUztBQUN0QyxlQUFXLFFBQVEsSUFBSSxjQUFjO0FBQ25DLFVBQUksQ0FBQyxLQUFLLGFBQ0gsS0FBSyxZQUFZLEtBQUssYUFDdEIsS0FBSyxTQUFTLEtBQUssV0FBVztBQUNuQyxlQUFPO0FBQUEsTUFDVDtBQUFBLElBQ0Y7QUFBQSxFQUNGO0FBQ0EsU0FBTztBQUNUO0FBRUEsb0JBQWUsU0FBQyxLQUFxQjtBQUNuQyxRQUFNLE1BQU0sSUFBSSxZQUFZLElBQ3hCLEtBQUssTUFBTyxJQUFJLFNBQVMsSUFBSSxZQUFhLEdBQUcsSUFDN0M7QUFDSixRQUFNLFdBQVcsSUFBSSxVQUFVLE9BQU8sT0FBSyxFQUFFLE9BQU8sS0FBSyxDQUFDO0FBRTFELFNBQU87QUFBQTtBQUFBLHFEQUUwQyxJQUFJLFFBQVE7QUFBQTtBQUFBO0FBQUE7QUFBQSwrQ0FJbEIsSUFBSSxNQUFNO0FBQUEsbUNBQ3RCLEdBQUc7QUFBQTtBQUFBLDJDQUVLLElBQUksTUFBTSxJQUFJLElBQUksU0FBUztBQUFBLGlDQUNyQyxjQUFjLElBQUksTUFBTSxDQUFDLElBQUksSUFBSSxNQUFNO0FBQUE7QUFBQSxZQUU1RCxTQUFTLFNBQVMsSUFBSTtBQUFBO0FBQUE7QUFBQSxrQkFHaEIsU0FBUyxJQUFJLE9BQUssV0FBVyxFQUFFLE9BQU8sT0FBTyxDQUFDO0FBQUE7QUFBQTtBQUFBLGNBR2xELE9BQU87QUFBQTtBQUFBO0FBQUE7QUFJbkI7QUFFQSwyQkFBc0IsV0FBRztBQUN2QixRQUFNLE9BQU8sbUJBQUssd0NBQVM7QUFDM0IsTUFBSSxDQUFDLE1BQU0sT0FBUSxRQUFPO0FBQzFCLFNBQU87QUFBQTtBQUFBO0FBQUE7QUFBQSxZQUlDLEtBQUssSUFBSSxTQUFPLFdBQVcsR0FBRyxPQUFPLENBQUM7QUFBQTtBQUFBO0FBQUE7QUFJaEQ7QUEvSEEsNEJBQVMsYUFEVCxnQkFIVyxnQkFJRjtBQUdULDhDQURBLGNBQ1Msb0RBQVQsUUFBUyxjQUFUO0FBR0EsNkNBREEsYUFDUyxrREFBVCxRQUFTLGFBQVQ7QUFWVyxpQkFBTiw4Q0FEUCw0QkFDYTtBQUNYLGNBRFcsZ0JBQ0osVUFBUztBQURYLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
