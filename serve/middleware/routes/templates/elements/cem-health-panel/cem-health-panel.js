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

// lit-css:elements/cem-health-panel/cem-health-panel.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n  padding: var(--pf-t--global--spacer--md);\\n}\\n\\n#overall {\\n  margin-block-end: var(--pf-t--global--spacer--md);\\n}\\n\\n.category-bar {\\n  display: flex;\\n  align-items: center;\\n  gap: var(--pf-t--global--spacer--sm);\\n}\\n\\n.category-meter {\\n  flex: 1;\\n  height: 8px;\\n  border-radius: var(--pf-t--global--border--radius--pill, 999px);\\n  background: var(--pf-t--global--background--color--secondary--default, #f0f0f0);\\n  overflow: hidden;\\n\\n  \\u0026 \\u003e .category-fill {\\n    height: 100%;\\n    border-radius: inherit;\\n    transition: width 0.3s ease;\\n  }\\n}\\n\\n.fill-pass {\\n  background: var(--pf-t--global--color--status--success--default, #3e8635);\\n}\\n\\n.fill-warn {\\n  background: var(--pf-t--global--color--status--warning--default, #f0ab00);\\n}\\n\\n.fill-fail {\\n  background: var(--pf-t--global--color--status--danger--default, #c9190b);\\n}\\n\\n.category-score {\\n  min-width: 4ch;\\n  text-align: end;\\n  font-variant-numeric: tabular-nums;\\n}\\n\\n#recommendations {\\n  margin-block-start: var(--pf-t--global--spacer--md);\\n  border-block-start: 1px solid var(--pf-t--global--border--color--default, #d2d2d2);\\n  padding-block-start: var(--pf-t--global--spacer--md);\\n\\n  \\u0026 h4 {\\n    margin: 0 0 var(--pf-t--global--spacer--sm);\\n    font-size: var(--pf-t--global--font--size--body--default, 1rem);\\n    font-weight: var(--pf-t--global--font--weight--body--bold, 700);\\n  }\\n\\n  \\u0026 ul {\\n    margin: 0;\\n    padding-inline-start: var(--pf-t--global--spacer--lg);\\n  }\\n\\n  \\u0026 li {\\n    margin-block-end: var(--pf-t--global--spacer--xs);\\n    font-size: var(--pf-t--global--font--size--body--sm, 0.875rem);\\n  }\\n}\\n\\n.finding-details {\\n  font-size: var(--pf-t--global--font--size--body--sm, 0.875rem);\\n  padding-block-start: var(--pf-t--global--spacer--xs);\\n\\n  \\u0026 ul {\\n    margin: 0;\\n    padding-inline-start: var(--pf-t--global--spacer--lg);\\n    list-style: none;\\n  }\\n\\n  \\u0026 li::before {\\n    content: \\"\\\\2192 \\";\\n  }\\n}\\n"'));
var cem_health_panel_default = s;

// elements/cem-health-panel/cem-health-panel.ts
import "../cem-pf-v6-label/cem-pf-v6-label.js";
import "../cem-pf-v6-expandable-section/cem-pf-v6-expandable-section.js";
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
        <cem-pf-v6-label color=${STATUS_COLORS[overallStatus]}
                     size="lg">${pct}% -- ${decl.score}/${decl.maxScore}</cem-pf-v6-label>
      </div>
      <dl id="categories"
          class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
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
      <div class="cem-pf-v6-c-description-list__group">
        <dt class="cem-pf-v6-c-description-list__term">${cat.category}</dt>
        <dd class="cem-pf-v6-c-description-list__description">
          <div class="category-bar">
            <div class="category-meter">
              <div class="category-fill fill-${cat.status}"
                   style="width: ${pct}%"></div>
            </div>
            <span class="category-score">${cat.points}/${cat.maxPoints}</span>
            <cem-pf-v6-label color=${STATUS_COLORS[cat.status]}>${cat.status}</cem-pf-v6-label>
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLWhlYWx0aC1wYW5lbC9jZW0taGVhbHRoLXBhbmVsLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLWhlYWx0aC1wYW5lbC9jZW0taGVhbHRoLXBhbmVsLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCwgbm90aGluZyB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuaW1wb3J0IHsgcHJvcGVydHkgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9wcm9wZXJ0eS5qcyc7XG5pbXBvcnQgeyBzdGF0ZSB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL3N0YXRlLmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1oZWFsdGgtcGFuZWwuY3NzJztcblxuaW1wb3J0ICcuLi9jZW0tcGYtdjYtbGFiZWwvY2VtLXBmLXY2LWxhYmVsLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LWV4cGFuZGFibGUtc2VjdGlvbi9jZW0tcGYtdjYtZXhwYW5kYWJsZS1zZWN0aW9uLmpzJztcblxuY29uc3QgU1RBVFVTX0NPTE9SUzogUmVjb3JkPHN0cmluZywgc3RyaW5nPiA9IHtcbiAgcGFzczogJ2dyZWVuJyxcbiAgd2FybjogJ29yYW5nZScsXG4gIGZhaWw6ICdyZWQnLFxufTtcblxuaW50ZXJmYWNlIEhlYWx0aEZpbmRpbmcge1xuICBtZXNzYWdlPzogc3RyaW5nO1xufVxuXG5pbnRlcmZhY2UgSGVhbHRoQ2F0ZWdvcnkge1xuICBjYXRlZ29yeTogc3RyaW5nO1xuICBwb2ludHM6IG51bWJlcjtcbiAgbWF4UG9pbnRzOiBudW1iZXI7XG4gIHN0YXR1czogc3RyaW5nO1xuICBmaW5kaW5ncz86IEhlYWx0aEZpbmRpbmdbXTtcbn1cblxuaW50ZXJmYWNlIEhlYWx0aERlY2xhcmF0aW9uIHtcbiAgdGFnTmFtZT86IHN0cmluZztcbiAgbmFtZT86IHN0cmluZztcbiAgc2NvcmU6IG51bWJlcjtcbiAgbWF4U2NvcmU6IG51bWJlcjtcbiAgY2F0ZWdvcmllczogSGVhbHRoQ2F0ZWdvcnlbXTtcbn1cblxuaW50ZXJmYWNlIEhlYWx0aE1vZHVsZSB7XG4gIGRlY2xhcmF0aW9uczogSGVhbHRoRGVjbGFyYXRpb25bXTtcbn1cblxuaW50ZXJmYWNlIEhlYWx0aFJlc3VsdCB7XG4gIG1vZHVsZXM6IEhlYWx0aE1vZHVsZVtdO1xuICByZWNvbW1lbmRhdGlvbnM/OiBzdHJpbmdbXTtcbn1cblxuLyoqXG4gKiBIZWFsdGggcGFuZWwgc2hvd2luZyBkb2N1bWVudGF0aW9uIHF1YWxpdHkgc2NvcmVzXG4gKlxuICogQGF0dHIge3N0cmluZ30gY29tcG9uZW50IC0gQ29tcG9uZW50IHRhZyBuYW1lIG9yIGNsYXNzIG5hbWUgdG8gZGlzcGxheSBoZWFsdGggZm9yXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0taGVhbHRoLXBhbmVsJylcbmV4cG9ydCBjbGFzcyBDZW1IZWFsdGhQYW5lbCBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gIEBwcm9wZXJ0eSh7IHJlZmxlY3Q6IHRydWUgfSlcbiAgYWNjZXNzb3IgY29tcG9uZW50OiBzdHJpbmcgfCBudWxsID0gbnVsbDtcblxuICBAc3RhdGUoKVxuICBhY2Nlc3NvciAjbG9hZGluZyA9IHRydWU7XG5cbiAgQHN0YXRlKClcbiAgYWNjZXNzb3IgI3Jlc3VsdDogSGVhbHRoUmVzdWx0IHwgbnVsbCA9IG51bGw7XG5cbiAgI2Fib3J0Q29udHJvbGxlcjogQWJvcnRDb250cm9sbGVyIHwgbnVsbCA9IG51bGw7XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLmZldGNoSGVhbHRoKCk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuI2Fib3J0Q29udHJvbGxlcj8uYWJvcnQoKTtcbiAgfVxuXG4gIGFzeW5jIGZldGNoSGVhbHRoKCkge1xuICAgIHRoaXMuI2Fib3J0Q29udHJvbGxlcj8uYWJvcnQoKTtcbiAgICB0aGlzLiNhYm9ydENvbnRyb2xsZXIgPSBuZXcgQWJvcnRDb250cm9sbGVyKCk7XG5cbiAgICB0aGlzLiNsb2FkaW5nID0gdHJ1ZTtcbiAgICB0aGlzLiNyZXN1bHQgPSBudWxsO1xuXG4gICAgY29uc3QgdXJsID0gbmV3IFVSTCgnL19fY2VtL2FwaS9oZWFsdGgnLCBsb2NhdGlvbi5vcmlnaW4pO1xuICAgIGlmICh0aGlzLmNvbXBvbmVudCkge1xuICAgICAgdXJsLnNlYXJjaFBhcmFtcy5zZXQoJ2NvbXBvbmVudCcsIHRoaXMuY29tcG9uZW50KTtcbiAgICB9XG5cbiAgICB0cnkge1xuICAgICAgY29uc3QgcmVzcG9uc2UgPSBhd2FpdCBmZXRjaCh1cmwsIHsgc2lnbmFsOiB0aGlzLiNhYm9ydENvbnRyb2xsZXIuc2lnbmFsIH0pO1xuICAgICAgaWYgKCFyZXNwb25zZS5vaykgdGhyb3cgbmV3IEVycm9yKGBIVFRQICR7cmVzcG9uc2Uuc3RhdHVzfWApO1xuICAgICAgdGhpcy4jcmVzdWx0ID0gYXdhaXQgcmVzcG9uc2UuanNvbigpO1xuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgIGlmICgoZSBhcyBFcnJvcikubmFtZSA9PT0gJ0Fib3J0RXJyb3InKSByZXR1cm47XG4gICAgICBjb25zb2xlLndhcm4oJ2NlbS1oZWFsdGgtcGFuZWw6IGZhaWxlZCB0byBmZXRjaCBoZWFsdGggZGF0YScsIGUpO1xuICAgIH0gZmluYWxseSB7XG4gICAgICB0aGlzLiNsb2FkaW5nID0gZmFsc2U7XG4gICAgfVxuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIGlmICh0aGlzLiNsb2FkaW5nKSB7XG4gICAgICByZXR1cm4gaHRtbGA8cD5BbmFseXppbmcgZG9jdW1lbnRhdGlvbiBoZWFsdGguLi48L3A+YDtcbiAgICB9XG5cbiAgICBjb25zdCBkZWNsID0gdGhpcy4jZmluZERlY2xhcmF0aW9uKCk7XG4gICAgaWYgKCFkZWNsKSB7XG4gICAgICByZXR1cm4gaHRtbGA8cD5ObyBoZWFsdGggZGF0YSBhdmFpbGFibGUuPC9wPmA7XG4gICAgfVxuXG4gICAgY29uc3QgcGN0ID0gZGVjbC5tYXhTY29yZSA+IDBcbiAgICAgID8gTWF0aC5yb3VuZCgoZGVjbC5zY29yZSAvIGRlY2wubWF4U2NvcmUpICogMTAwKVxuICAgICAgOiAwO1xuICAgIGNvbnN0IG92ZXJhbGxTdGF0dXMgPSBwY3QgPj0gODAgPyAncGFzcycgOiBwY3QgPj0gNDAgPyAnd2FybicgOiAnZmFpbCc7XG5cbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJvdmVyYWxsXCI+XG4gICAgICAgIDxjZW0tcGYtdjYtbGFiZWwgY29sb3I9JHtTVEFUVVNfQ09MT1JTW292ZXJhbGxTdGF0dXNdfVxuICAgICAgICAgICAgICAgICAgICAgc2l6ZT1cImxnXCI+JHtwY3R9JSAtLSAke2RlY2wuc2NvcmV9LyR7ZGVjbC5tYXhTY29yZX08L2NlbS1wZi12Ni1sYWJlbD5cbiAgICAgIDwvZGl2PlxuICAgICAgPGRsIGlkPVwiY2F0ZWdvcmllc1wiXG4gICAgICAgICAgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0IHBmLW0taG9yaXpvbnRhbCBwZi1tLWNvbXBhY3RcIj5cbiAgICAgICAgJHtkZWNsLmNhdGVnb3JpZXMubWFwKGNhdCA9PiB0aGlzLiNyZW5kZXJDYXRlZ29yeShjYXQpKX1cbiAgICAgIDwvZGw+XG4gICAgICAke3RoaXMuI3JlbmRlclJlY29tbWVuZGF0aW9ucygpfVxuICAgIGA7XG4gIH1cblxuICAjZmluZERlY2xhcmF0aW9uKCk6IEhlYWx0aERlY2xhcmF0aW9uIHwgbnVsbCB7XG4gICAgaWYgKCF0aGlzLiNyZXN1bHQpIHJldHVybiBudWxsO1xuICAgIGZvciAoY29uc3QgbW9kIG9mIHRoaXMuI3Jlc3VsdC5tb2R1bGVzKSB7XG4gICAgICBmb3IgKGNvbnN0IGRlY2wgb2YgbW9kLmRlY2xhcmF0aW9ucykge1xuICAgICAgICBpZiAoIXRoaXMuY29tcG9uZW50XG4gICAgICAgICAgICB8fCBkZWNsLnRhZ05hbWUgPT09IHRoaXMuY29tcG9uZW50XG4gICAgICAgICAgICB8fCBkZWNsLm5hbWUgPT09IHRoaXMuY29tcG9uZW50KSB7XG4gICAgICAgICAgcmV0dXJuIGRlY2w7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuIG51bGw7XG4gIH1cblxuICAjcmVuZGVyQ2F0ZWdvcnkoY2F0OiBIZWFsdGhDYXRlZ29yeSkge1xuICAgIGNvbnN0IHBjdCA9IGNhdC5tYXhQb2ludHMgPiAwXG4gICAgICA/IE1hdGgucm91bmQoKGNhdC5wb2ludHMgLyBjYXQubWF4UG9pbnRzKSAqIDEwMClcbiAgICAgIDogMDtcbiAgICBjb25zdCBmaW5kaW5ncyA9IGNhdC5maW5kaW5ncz8uZmlsdGVyKGYgPT4gZi5tZXNzYWdlKSA/PyBbXTtcblxuICAgIHJldHVybiBodG1sYFxuICAgICAgPGRpdiBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX2dyb3VwXCI+XG4gICAgICAgIDxkdCBjbGFzcz1cImNlbS1wZi12Ni1jLWRlc2NyaXB0aW9uLWxpc3RfX3Rlcm1cIj4ke2NhdC5jYXRlZ29yeX08L2R0PlxuICAgICAgICA8ZGQgY2xhc3M9XCJjZW0tcGYtdjYtYy1kZXNjcmlwdGlvbi1saXN0X19kZXNjcmlwdGlvblwiPlxuICAgICAgICAgIDxkaXYgY2xhc3M9XCJjYXRlZ29yeS1iYXJcIj5cbiAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJjYXRlZ29yeS1tZXRlclwiPlxuICAgICAgICAgICAgICA8ZGl2IGNsYXNzPVwiY2F0ZWdvcnktZmlsbCBmaWxsLSR7Y2F0LnN0YXR1c31cIlxuICAgICAgICAgICAgICAgICAgIHN0eWxlPVwid2lkdGg6ICR7cGN0fSVcIj48L2Rpdj5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgPHNwYW4gY2xhc3M9XCJjYXRlZ29yeS1zY29yZVwiPiR7Y2F0LnBvaW50c30vJHtjYXQubWF4UG9pbnRzfTwvc3Bhbj5cbiAgICAgICAgICAgIDxjZW0tcGYtdjYtbGFiZWwgY29sb3I9JHtTVEFUVVNfQ09MT1JTW2NhdC5zdGF0dXNdfT4ke2NhdC5zdGF0dXN9PC9jZW0tcGYtdjYtbGFiZWw+XG4gICAgICAgICAgPC9kaXY+XG4gICAgICAgICAgJHtmaW5kaW5ncy5sZW5ndGggPiAwID8gaHRtbGBcbiAgICAgICAgICAgIDxkaXYgY2xhc3M9XCJmaW5kaW5nLWRldGFpbHNcIj5cbiAgICAgICAgICAgICAgPHVsPlxuICAgICAgICAgICAgICAgICR7ZmluZGluZ3MubWFwKGYgPT4gaHRtbGA8bGk+JHtmLm1lc3NhZ2V9PC9saT5gKX1cbiAgICAgICAgICAgICAgPC91bD5cbiAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgIGAgOiBub3RoaW5nfVxuICAgICAgICA8L2RkPlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gICNyZW5kZXJSZWNvbW1lbmRhdGlvbnMoKSB7XG4gICAgY29uc3QgcmVjcyA9IHRoaXMuI3Jlc3VsdD8ucmVjb21tZW5kYXRpb25zO1xuICAgIGlmICghcmVjcz8ubGVuZ3RoKSByZXR1cm4gbm90aGluZztcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJyZWNvbW1lbmRhdGlvbnNcIj5cbiAgICAgICAgPGg0PlJlY29tbWVuZGF0aW9uczwvaDQ+XG4gICAgICAgIDx1bD5cbiAgICAgICAgICAke3JlY3MubWFwKHJlYyA9PiBodG1sYDxsaT4ke3JlY308L2xpPmApfVxuICAgICAgICA8L3VsPlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0taGVhbHRoLXBhbmVsJzogQ2VtSGVhbHRoUGFuZWw7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1tZCk7XFxcXG59XFxcXG5cXFxcbiNvdmVyYWxsIHtcXFxcbiAgbWFyZ2luLWJsb2NrLWVuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLW1kKTtcXFxcbn1cXFxcblxcXFxuLmNhdGVnb3J5LWJhciB7XFxcXG4gIGRpc3BsYXk6IGZsZXg7XFxcXG4gIGFsaWduLWl0ZW1zOiBjZW50ZXI7XFxcXG4gIGdhcDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbn1cXFxcblxcXFxuLmNhdGVnb3J5LW1ldGVyIHtcXFxcbiAgZmxleDogMTtcXFxcbiAgaGVpZ2h0OiA4cHg7XFxcXG4gIGJvcmRlci1yYWRpdXM6IHZhcigtLXBmLXQtLWdsb2JhbC0tYm9yZGVyLS1yYWRpdXMtLXBpbGwsIDk5OXB4KTtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1kZWZhdWx0LCAjZjBmMGYwKTtcXFxcbiAgb3ZlcmZsb3c6IGhpZGRlbjtcXFxcblxcXFxuICBcXFxcdTAwMjYgXFxcXHUwMDNlIC5jYXRlZ29yeS1maWxsIHtcXFxcbiAgICBoZWlnaHQ6IDEwMCU7XFxcXG4gICAgYm9yZGVyLXJhZGl1czogaW5oZXJpdDtcXFxcbiAgICB0cmFuc2l0aW9uOiB3aWR0aCAwLjNzIGVhc2U7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmZpbGwtcGFzcyB7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLXN0YXR1cy0tc3VjY2Vzcy0tZGVmYXVsdCwgIzNlODYzNSk7XFxcXG59XFxcXG5cXFxcbi5maWxsLXdhcm4ge1xcXFxuICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWNvbG9yLS1zdGF0dXMtLXdhcm5pbmctLWRlZmF1bHQsICNmMGFiMDApO1xcXFxufVxcXFxuXFxcXG4uZmlsbC1mYWlsIHtcXFxcbiAgYmFja2dyb3VuZDogdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tc3RhdHVzLS1kYW5nZXItLWRlZmF1bHQsICNjOTE5MGIpO1xcXFxufVxcXFxuXFxcXG4uY2F0ZWdvcnktc2NvcmUge1xcXFxuICBtaW4td2lkdGg6IDRjaDtcXFxcbiAgdGV4dC1hbGlnbjogZW5kO1xcXFxuICBmb250LXZhcmlhbnQtbnVtZXJpYzogdGFidWxhci1udW1zO1xcXFxufVxcXFxuXFxcXG4jcmVjb21tZW5kYXRpb25zIHtcXFxcbiAgbWFyZ2luLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuICBib3JkZXItYmxvY2stc3RhcnQ6IDFweCBzb2xpZCB2YXIoLS1wZi10LS1nbG9iYWwtLWJvcmRlci0tY29sb3ItLWRlZmF1bHQsICNkMmQyZDIpO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbWQpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNiBoNCB7XFxcXG4gICAgbWFyZ2luOiAwIDAgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLXNtKTtcXFxcbiAgICBmb250LXNpemU6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0tc2l6ZS0tYm9keS0tZGVmYXVsdCwgMXJlbSk7XFxcXG4gICAgZm9udC13ZWlnaHQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tZm9udC0td2VpZ2h0LS1ib2R5LS1ib2xkLCA3MDApO1xcXFxuICB9XFxcXG5cXFxcbiAgXFxcXHUwMDI2IHVsIHtcXFxcbiAgICBtYXJnaW46IDA7XFxcXG4gICAgcGFkZGluZy1pbmxpbmUtc3RhcnQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYgbGkge1xcXFxuICAgIG1hcmdpbi1ibG9jay1lbmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS14cyk7XFxcXG4gICAgZm9udC1zaXplOiB2YXIoLS1wZi10LS1nbG9iYWwtLWZvbnQtLXNpemUtLWJvZHktLXNtLCAwLjg3NXJlbSk7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLmZpbmRpbmctZGV0YWlscyB7XFxcXG4gIGZvbnQtc2l6ZTogdmFyKC0tcGYtdC0tZ2xvYmFsLS1mb250LS1zaXplLS1ib2R5LS1zbSwgMC44NzVyZW0pO1xcXFxuICBwYWRkaW5nLWJsb2NrLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0teHMpO1xcXFxuXFxcXG4gIFxcXFx1MDAyNiB1bCB7XFxcXG4gICAgbWFyZ2luOiAwO1xcXFxuICAgIHBhZGRpbmctaW5saW5lLXN0YXJ0OiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxuICAgIGxpc3Qtc3R5bGU6IG5vbmU7XFxcXG4gIH1cXFxcblxcXFxuICBcXFxcdTAwMjYgbGk6OmJlZm9yZSB7XFxcXG4gICAgY29udGVudDogXFxcXFxcXCJcXFxcXFxcXDIxOTIgXFxcXFxcXCI7XFxcXG4gIH1cXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLE1BQU0sZUFBZTtBQUMxQyxTQUFTLHFCQUFxQjtBQUM5QixTQUFTLGdCQUFnQjtBQUN6QixTQUFTLGFBQWE7OztBQ0h0QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxtakVBQXVqRSxDQUFDO0FBQUUsSUFBTywyQkFBUTs7O0FETzluRSxPQUFPO0FBQ1AsT0FBTztBQUVQLElBQU0sZ0JB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
