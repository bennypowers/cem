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
var __accessCheck = (obj, member, msg) => member.has(obj) || __typeError("Cannot " + msg);
var __privateIn = (member, obj) => Object(obj) !== obj ? __typeError('Cannot use the "in" operator on this value') : member.has(obj);
var __privateGet = (obj, member, getter) => (__accessCheck(obj, member, "read from private field"), getter ? getter.call(obj) : member.get(obj));
var __privateSet = (obj, member, value, setter) => (__accessCheck(obj, member, "write to private field"), setter ? setter.call(obj, value) : member.set(obj, value), value);
var __privateMethod = (obj, member, method) => (__accessCheck(obj, member, "access private method"), method);

// elements/cem-color-scheme-toggle/cem-color-scheme-toggle.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-color-scheme-toggle/cem-color-scheme-toggle.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\n#toggle-group {\\n  display: flex;\\n  gap: 2px;\\n  background: var(--pf-t--global--background--color--secondary--hover);\\n  border-radius: 6px;\\n  padding: 2px;\\n  border: 1px solid var(--pf-t--global--border--color--default);\\n}\\n\\n.toggle-option {\\n  position: relative;\\n  cursor: pointer;\\n  margin: 0;\\n}\\n\\n.toggle-option input[type=\\"radio\\"] {\\n  position: absolute;\\n  opacity: 0;\\n  width: 0;\\n  height: 0;\\n}\\n\\n.toggle-label {\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n  width: 32px;\\n  height: 32px;\\n  border-radius: 4px;\\n  color: var(--pf-t--global--icon--color--subtle);\\n  transition: all 200ms cubic-bezier(0.4, 0, 0.2, 1);\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--background--color--primary--hover);\\n    color: var(--pf-t--global--icon--color--regular);\\n  }\\n}\\n\\n.toggle-option input[type=\\"radio\\"]:checked + .toggle-label {\\n  background: var(--pf-t--global--color--brand--default);\\n  color: var(--pf-t--global--icon--color--on-brand--default);\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--color--brand--hover);\\n    color: var(--pf-t--global--icon--color--on-brand--hover);\\n  }\\n}\\n\\n.toggle-option input[type=\\"radio\\"]:focus-visible + .toggle-label {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: 2px;\\n}\\n\\n.sr-only {\\n  position: absolute;\\n  width: 1px;\\n  height: 1px;\\n  padding: 0;\\n  margin: -1px;\\n  overflow: hidden;\\n  clip: rect(0, 0, 0, 0);\\n  white-space: nowrap;\\n  border: 0;\\n}\\n"'));
var cem_color_scheme_toggle_default = s;

// elements/cem-color-scheme-toggle/cem-color-scheme-toggle.ts
var _CemColorSchemeToggle_decorators, _init, _a;
_CemColorSchemeToggle_decorators = [customElement("cem-color-scheme-toggle")];
var CemColorSchemeToggle = class extends (_a = LitElement) {
  static styles = cem_color_scheme_toggle_default;
  /** Storage access gatekeeper - localStorage can throw in Safari private mode */
  #getStorageItem(key, defaultValue) {
    try {
      return localStorage.getItem(key) ?? defaultValue;
    } catch {
      return defaultValue;
    }
  }
  #setStorageItem(key, value) {
    try {
      localStorage.setItem(key, value);
    } catch {
    }
  }
  connectedCallback() {
    super.connectedCallback();
    const saved = this.#getStorageItem("cem-serve-color-scheme", "system");
    this.#applyColorScheme(saved);
  }
  render() {
    const saved = this.#getStorageItem("cem-serve-color-scheme", "system");
    return html`
      <div id="toggle-group"
           role="radiogroup"
           aria-label="Color scheme">
        <label class="toggle-option">
          <input type="radio"
                 name="color-scheme"
                 value="light"
                 .checked=${saved === "light"}
                 @change=${this.#onChange}>
          <span class="toggle-label">
            <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
              <path d="M8 11a3 3 0 1 1 0-6 3 3 0 0 1 0 6zm0 1a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z"/>
            </svg>
            <span class="sr-only">Light</span>
          </span>
        </label>
        <label class="toggle-option">
          <input type="radio"
                 name="color-scheme"
                 value="system"
                 .checked=${saved === "system"}
                 @change=${this.#onChange}>
          <span class="toggle-label">
            <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
              <path d="M1 2.5A1.5 1.5 0 0 1 2.5 1h3A1.5 1.5 0 0 1 7 2.5v3A1.5 1.5 0 0 1 5.5 7h-3A1.5 1.5 0 0 1 1 5.5v-3zm8 0A1.5 1.5 0 0 1 10.5 1h3A1.5 1.5 0 0 1 15 2.5v3A1.5 1.5 0 0 1 13.5 7h-3A1.5 1.5 0 0 1 9 5.5v-3zm-8 8A1.5 1.5 0 0 1 2.5 9h3A1.5 1.5 0 0 1 7 10.5v3A1.5 1.5 0 0 1 5.5 15h-3A1.5 1.5 0 0 1 1 13.5v-3zm8 0A1.5 1.5 0 0 1 10.5 9h3a1.5 1.5 0 0 1 1.5 1.5v3a1.5 1.5 0 0 1-1.5 1.5h-3A1.5 1.5 0 0 1 9 13.5v-3z"/>
            </svg>
            <span class="sr-only">System</span>
          </span>
        </label>
        <label class="toggle-option">
          <input type="radio"
                 name="color-scheme"
                 value="dark"
                 .checked=${saved === "dark"}
                 @change=${this.#onChange}>
          <span class="toggle-label">
            <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
              <path d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278z"/>
            </svg>
            <span class="sr-only">Dark</span>
          </span>
        </label>
      </div>
    `;
  }
  #onChange(e) {
    const target = e.target;
    if (target.checked) {
      this.#applyColorScheme(target.value);
    }
  }
  #applyColorScheme(scheme) {
    switch (scheme) {
      case "light":
        document.body.style.colorScheme = "light";
        break;
      case "dark":
        document.body.style.colorScheme = "dark";
        break;
      case "system":
      default:
        document.body.style.colorScheme = "light dark";
        break;
    }
    this.#setStorageItem("cem-serve-color-scheme", scheme);
  }
};
_init = __decoratorStart(_a);
CemColorSchemeToggle = __decorateElement(_init, 0, "CemColorSchemeToggle", _CemColorSchemeToggle_decorators, CemColorSchemeToggle);
__runInitializers(_init, 1, CemColorSchemeToggle);
export {
  CemColorSchemeToggle
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLWNvbG9yLXNjaGVtZS10b2dnbGUvY2VtLWNvbG9yLXNjaGVtZS10b2dnbGUudHMiLCAibGl0LWNzczplbGVtZW50cy9jZW0tY29sb3Itc2NoZW1lLXRvZ2dsZS9jZW0tY29sb3Itc2NoZW1lLXRvZ2dsZS5jc3MiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImltcG9ydCB7IExpdEVsZW1lbnQsIGh0bWwgfSBmcm9tICdsaXQnO1xuaW1wb3J0IHsgY3VzdG9tRWxlbWVudCB9IGZyb20gJ2xpdC9kZWNvcmF0b3JzL2N1c3RvbS1lbGVtZW50LmpzJztcblxuaW1wb3J0IHN0eWxlcyBmcm9tICcuL2NlbS1jb2xvci1zY2hlbWUtdG9nZ2xlLmNzcyc7XG5cbnR5cGUgQ29sb3JTY2hlbWUgPSAnbGlnaHQnIHwgJ2RhcmsnIHwgJ3N5c3RlbSc7XG5cbi8qKlxuICogVG9nZ2xlIGJldHdlZW4gbGlnaHQsIGRhcmssIGFuZCBzeXN0ZW0gY29sb3Igc2NoZW1lcy5cbiAqXG4gKiBAY3VzdG9tRWxlbWVudCBjZW0tY29sb3Itc2NoZW1lLXRvZ2dsZVxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLWNvbG9yLXNjaGVtZS10b2dnbGUnKVxuZXhwb3J0IGNsYXNzIENlbUNvbG9yU2NoZW1lVG9nZ2xlIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgLyoqIFN0b3JhZ2UgYWNjZXNzIGdhdGVrZWVwZXIgLSBsb2NhbFN0b3JhZ2UgY2FuIHRocm93IGluIFNhZmFyaSBwcml2YXRlIG1vZGUgKi9cbiAgI2dldFN0b3JhZ2VJdGVtKGtleTogc3RyaW5nLCBkZWZhdWx0VmFsdWU6IHN0cmluZyk6IHN0cmluZyB7XG4gICAgdHJ5IHtcbiAgICAgIHJldHVybiBsb2NhbFN0b3JhZ2UuZ2V0SXRlbShrZXkpID8/IGRlZmF1bHRWYWx1ZTtcbiAgICB9IGNhdGNoIHtcbiAgICAgIHJldHVybiBkZWZhdWx0VmFsdWU7XG4gICAgfVxuICB9XG5cbiAgI3NldFN0b3JhZ2VJdGVtKGtleTogc3RyaW5nLCB2YWx1ZTogc3RyaW5nKTogdm9pZCB7XG4gICAgdHJ5IHtcbiAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKGtleSwgdmFsdWUpO1xuICAgIH0gY2F0Y2gge1xuICAgICAgLy8gU3RvcmFnZSB1bmF2YWlsYWJsZSAocHJpdmF0ZSBtb2RlKSwgc2lsZW50bHkgY29udGludWVcbiAgICB9XG4gIH1cblxuICBjb25uZWN0ZWRDYWxsYmFjaygpOiB2b2lkIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIC8vIFJlc3RvcmUgc2F2ZWQgY29sb3Igc2NoZW1lIHByZWZlcmVuY2VcbiAgICBjb25zdCBzYXZlZCA9IHRoaXMuI2dldFN0b3JhZ2VJdGVtKCdjZW0tc2VydmUtY29sb3Itc2NoZW1lJywgJ3N5c3RlbScpIGFzIENvbG9yU2NoZW1lO1xuICAgIHRoaXMuI2FwcGx5Q29sb3JTY2hlbWUoc2F2ZWQpO1xuICB9XG5cbiAgcmVuZGVyKCkge1xuICAgIGNvbnN0IHNhdmVkID0gdGhpcy4jZ2V0U3RvcmFnZUl0ZW0oJ2NlbS1zZXJ2ZS1jb2xvci1zY2hlbWUnLCAnc3lzdGVtJyk7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8ZGl2IGlkPVwidG9nZ2xlLWdyb3VwXCJcbiAgICAgICAgICAgcm9sZT1cInJhZGlvZ3JvdXBcIlxuICAgICAgICAgICBhcmlhLWxhYmVsPVwiQ29sb3Igc2NoZW1lXCI+XG4gICAgICAgIDxsYWJlbCBjbGFzcz1cInRvZ2dsZS1vcHRpb25cIj5cbiAgICAgICAgICA8aW5wdXQgdHlwZT1cInJhZGlvXCJcbiAgICAgICAgICAgICAgICAgbmFtZT1cImNvbG9yLXNjaGVtZVwiXG4gICAgICAgICAgICAgICAgIHZhbHVlPVwibGlnaHRcIlxuICAgICAgICAgICAgICAgICAuY2hlY2tlZD0ke3NhdmVkID09PSAnbGlnaHQnfVxuICAgICAgICAgICAgICAgICBAY2hhbmdlPSR7dGhpcy4jb25DaGFuZ2V9PlxuICAgICAgICAgIDxzcGFuIGNsYXNzPVwidG9nZ2xlLWxhYmVsXCI+XG4gICAgICAgICAgICA8c3ZnIHdpZHRoPVwiMTZcIiBoZWlnaHQ9XCIxNlwiIHZpZXdCb3g9XCIwIDAgMTYgMTZcIiBmaWxsPVwiY3VycmVudENvbG9yXCIgYXJpYS1oaWRkZW49XCJ0cnVlXCI+XG4gICAgICAgICAgICAgIDxwYXRoIGQ9XCJNOCAxMWEzIDMgMCAxIDEgMC02IDMgMyAwIDAgMSAwIDZ6bTAgMWE0IDQgMCAxIDAgMC04IDQgNCAwIDAgMCAwIDh6TTggMGEuNS41IDAgMCAxIC41LjV2MmEuNS41IDAgMCAxLTEgMHYtMkEuNS41IDAgMCAxIDggMHptMCAxM2EuNS41IDAgMCAxIC41LjV2MmEuNS41IDAgMCAxLTEgMHYtMkEuNS41IDAgMCAxIDggMTN6bTgtNWEuNS41IDAgMCAxLS41LjVoLTJhLjUuNSAwIDAgMSAwLTFoMmEuNS41IDAgMCAxIC41LjV6TTMgOGEuNS41IDAgMCAxLS41LjVoLTJhLjUuNSAwIDAgMSAwLTFoMkEuNS41IDAgMCAxIDMgOHptMTAuNjU3LTUuNjU3YS41LjUgMCAwIDEgMCAuNzA3bC0xLjQxNCAxLjQxNWEuNS41IDAgMSAxLS43MDctLjcwOGwxLjQxNC0xLjQxNGEuNS41IDAgMCAxIC43MDcgMHptLTkuMTkzIDkuMTkzYS41LjUgMCAwIDEgMCAuNzA3TDMuMDUgMTMuNjU3YS41LjUgMCAwIDEtLjcwNy0uNzA3bDEuNDE0LTEuNDE0YS41LjUgMCAwIDEgLjcwNyAwem05LjE5MyAyLjEyMWEuNS41IDAgMCAxLS43MDcgMGwtMS40MTQtMS40MTRhLjUuNSAwIDAgMSAuNzA3LS43MDdsMS40MTQgMS40MTRhLjUuNSAwIDAgMSAwIC43MDd6TTQuNDY0IDQuNDY1YS41LjUgMCAwIDEtLjcwNyAwTDIuMzQzIDMuMDVhLjUuNSAwIDEgMSAuNzA3LS43MDdsMS40MTQgMS40MTRhLjUuNSAwIDAgMSAwIC43MDh6XCIvPlxuICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICA8c3BhbiBjbGFzcz1cInNyLW9ubHlcIj5MaWdodDwvc3Bhbj5cbiAgICAgICAgICA8L3NwYW4+XG4gICAgICAgIDwvbGFiZWw+XG4gICAgICAgIDxsYWJlbCBjbGFzcz1cInRvZ2dsZS1vcHRpb25cIj5cbiAgICAgICAgICA8aW5wdXQgdHlwZT1cInJhZGlvXCJcbiAgICAgICAgICAgICAgICAgbmFtZT1cImNvbG9yLXNjaGVtZVwiXG4gICAgICAgICAgICAgICAgIHZhbHVlPVwic3lzdGVtXCJcbiAgICAgICAgICAgICAgICAgLmNoZWNrZWQ9JHtzYXZlZCA9PT0gJ3N5c3RlbSd9XG4gICAgICAgICAgICAgICAgIEBjaGFuZ2U9JHt0aGlzLiNvbkNoYW5nZX0+XG4gICAgICAgICAgPHNwYW4gY2xhc3M9XCJ0b2dnbGUtbGFiZWxcIj5cbiAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIiBhcmlhLWhpZGRlbj1cInRydWVcIj5cbiAgICAgICAgICAgICAgPHBhdGggZD1cIk0xIDIuNUExLjUgMS41IDAgMCAxIDIuNSAxaDNBMS41IDEuNSAwIDAgMSA3IDIuNXYzQTEuNSAxLjUgMCAwIDEgNS41IDdoLTNBMS41IDEuNSAwIDAgMSAxIDUuNXYtM3ptOCAwQTEuNSAxLjUgMCAwIDEgMTAuNSAxaDNBMS41IDEuNSAwIDAgMSAxNSAyLjV2M0ExLjUgMS41IDAgMCAxIDEzLjUgN2gtM0ExLjUgMS41IDAgMCAxIDkgNS41di0zem0tOCA4QTEuNSAxLjUgMCAwIDEgMi41IDloM0ExLjUgMS41IDAgMCAxIDcgMTAuNXYzQTEuNSAxLjUgMCAwIDEgNS41IDE1aC0zQTEuNSAxLjUgMCAwIDEgMSAxMy41di0zem04IDBBMS41IDEuNSAwIDAgMSAxMC41IDloM2ExLjUgMS41IDAgMCAxIDEuNSAxLjV2M2ExLjUgMS41IDAgMCAxLTEuNSAxLjVoLTNBMS41IDEuNSAwIDAgMSA5IDEzLjV2LTN6XCIvPlxuICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICA8c3BhbiBjbGFzcz1cInNyLW9ubHlcIj5TeXN0ZW08L3NwYW4+XG4gICAgICAgICAgPC9zcGFuPlxuICAgICAgICA8L2xhYmVsPlxuICAgICAgICA8bGFiZWwgY2xhc3M9XCJ0b2dnbGUtb3B0aW9uXCI+XG4gICAgICAgICAgPGlucHV0IHR5cGU9XCJyYWRpb1wiXG4gICAgICAgICAgICAgICAgIG5hbWU9XCJjb2xvci1zY2hlbWVcIlxuICAgICAgICAgICAgICAgICB2YWx1ZT1cImRhcmtcIlxuICAgICAgICAgICAgICAgICAuY2hlY2tlZD0ke3NhdmVkID09PSAnZGFyayd9XG4gICAgICAgICAgICAgICAgIEBjaGFuZ2U9JHt0aGlzLiNvbkNoYW5nZX0+XG4gICAgICAgICAgPHNwYW4gY2xhc3M9XCJ0b2dnbGUtbGFiZWxcIj5cbiAgICAgICAgICAgIDxzdmcgd2lkdGg9XCIxNlwiIGhlaWdodD1cIjE2XCIgdmlld0JveD1cIjAgMCAxNiAxNlwiIGZpbGw9XCJjdXJyZW50Q29sb3JcIiBhcmlhLWhpZGRlbj1cInRydWVcIj5cbiAgICAgICAgICAgICAgPHBhdGggZD1cIk02IC4yNzhhLjc2OC43NjggMCAwIDEgLjA4Ljg1OCA3LjIwOCA3LjIwOCAwIDAgMC0uODc4IDMuNDZjMCA0LjAyMSAzLjI3OCA3LjI3NyA3LjMxOCA3LjI3Ny41MjcgMCAxLjA0LS4wNTUgMS41MzMtLjE2YS43ODcuNzg3IDAgMCAxIC44MS4zMTYuNzMzLjczMyAwIDAgMS0uMDMxLjg5M0E4LjM0OSA4LjM0OSAwIDAgMSA4LjM0NCAxNkMzLjczNCAxNiAwIDEyLjI4NiAwIDcuNzEgMCA0LjI2NiAyLjExNCAxLjMxMiA1LjEyNC4wNkEuNzUyLjc1MiAwIDAgMSA2IC4yNzh6XCIvPlxuICAgICAgICAgICAgPC9zdmc+XG4gICAgICAgICAgICA8c3BhbiBjbGFzcz1cInNyLW9ubHlcIj5EYXJrPC9zcGFuPlxuICAgICAgICAgIDwvc3Bhbj5cbiAgICAgICAgPC9sYWJlbD5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cblxuICAjb25DaGFuZ2UoZTogRXZlbnQpOiB2b2lkIHtcbiAgICBjb25zdCB0YXJnZXQgPSBlLnRhcmdldCBhcyBIVE1MSW5wdXRFbGVtZW50O1xuICAgIGlmICh0YXJnZXQuY2hlY2tlZCkge1xuICAgICAgdGhpcy4jYXBwbHlDb2xvclNjaGVtZSh0YXJnZXQudmFsdWUgYXMgQ29sb3JTY2hlbWUpO1xuICAgIH1cbiAgfVxuXG4gICNhcHBseUNvbG9yU2NoZW1lKHNjaGVtZTogQ29sb3JTY2hlbWUpOiB2b2lkIHtcbiAgICBzd2l0Y2ggKHNjaGVtZSkge1xuICAgICAgY2FzZSAnbGlnaHQnOlxuICAgICAgICBkb2N1bWVudC5ib2R5LnN0eWxlLmNvbG9yU2NoZW1lID0gJ2xpZ2h0JztcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdkYXJrJzpcbiAgICAgICAgZG9jdW1lbnQuYm9keS5zdHlsZS5jb2xvclNjaGVtZSA9ICdkYXJrJztcbiAgICAgICAgYnJlYWs7XG4gICAgICBjYXNlICdzeXN0ZW0nOlxuICAgICAgZGVmYXVsdDpcbiAgICAgICAgZG9jdW1lbnQuYm9keS5zdHlsZS5jb2xvclNjaGVtZSA9ICdsaWdodCBkYXJrJztcbiAgICAgICAgYnJlYWs7XG4gICAgfVxuICAgIHRoaXMuI3NldFN0b3JhZ2VJdGVtKCdjZW0tc2VydmUtY29sb3Itc2NoZW1lJywgc2NoZW1lKTtcbiAgfVxufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tY29sb3Itc2NoZW1lLXRvZ2dsZSc6IENlbUNvbG9yU2NoZW1lVG9nZ2xlO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxufVxcXFxuXFxcXG4jdG9nZ2xlLWdyb3VwIHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZ2FwOiAycHg7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0taG92ZXIpO1xcXFxuICBib3JkZXItcmFkaXVzOiA2cHg7XFxcXG4gIHBhZGRpbmc6IDJweDtcXFxcbiAgYm9yZGVyOiAxcHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1ib3JkZXItLWNvbG9yLS1kZWZhdWx0KTtcXFxcbn1cXFxcblxcXFxuLnRvZ2dsZS1vcHRpb24ge1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG4gIGN1cnNvcjogcG9pbnRlcjtcXFxcbiAgbWFyZ2luOiAwO1xcXFxufVxcXFxuXFxcXG4udG9nZ2xlLW9wdGlvbiBpbnB1dFt0eXBlPVxcXFxcXFwicmFkaW9cXFxcXFxcIl0ge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIG9wYWNpdHk6IDA7XFxcXG4gIHdpZHRoOiAwO1xcXFxuICBoZWlnaHQ6IDA7XFxcXG59XFxcXG5cXFxcbi50b2dnbGUtbGFiZWwge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBhbGlnbi1pdGVtczogY2VudGVyO1xcXFxuICBqdXN0aWZ5LWNvbnRlbnQ6IGNlbnRlcjtcXFxcbiAgd2lkdGg6IDMycHg7XFxcXG4gIGhlaWdodDogMzJweDtcXFxcbiAgYm9yZGVyLXJhZGl1czogNHB4O1xcXFxuICBjb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1pY29uLS1jb2xvci0tc3VidGxlKTtcXFxcbiAgdHJhbnNpdGlvbjogYWxsIDIwMG1zIGN1YmljLWJlemllcigwLjQsIDAsIDAuMiwgMSk7XFxcXG5cXFxcbiAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICBiYWNrZ3JvdW5kOiB2YXIoLS1wZi10LS1nbG9iYWwtLWJhY2tncm91bmQtLWNvbG9yLS1wcmltYXJ5LS1ob3Zlcik7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLXJlZ3VsYXIpO1xcXFxuICB9XFxcXG59XFxcXG5cXFxcbi50b2dnbGUtb3B0aW9uIGlucHV0W3R5cGU9XFxcXFxcXCJyYWRpb1xcXFxcXFwiXTpjaGVja2VkICsgLnRvZ2dsZS1sYWJlbCB7XFxcXG4gIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLW9uLWJyYW5kLS1kZWZhdWx0KTtcXFxcblxcXFxuICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgIGJhY2tncm91bmQ6IHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1ob3Zlcik7XFxcXG4gICAgY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0taWNvbi0tY29sb3ItLW9uLWJyYW5kLS1ob3Zlcik7XFxcXG4gIH1cXFxcbn1cXFxcblxcXFxuLnRvZ2dsZS1vcHRpb24gaW5wdXRbdHlwZT1cXFxcXFxcInJhZGlvXFxcXFxcXCJdOmZvY3VzLXZpc2libGUgKyAudG9nZ2xlLWxhYmVsIHtcXFxcbiAgb3V0bGluZTogMnB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0KTtcXFxcbiAgb3V0bGluZS1vZmZzZXQ6IDJweDtcXFxcbn1cXFxcblxcXFxuLnNyLW9ubHkge1xcXFxuICBwb3NpdGlvbjogYWJzb2x1dGU7XFxcXG4gIHdpZHRoOiAxcHg7XFxcXG4gIGhlaWdodDogMXB4O1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBtYXJnaW46IC0xcHg7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG4gIGNsaXA6IHJlY3QoMCwgMCwgMCwgMCk7XFxcXG4gIHdoaXRlLXNwYWNlOiBub3dyYXA7XFxcXG4gIGJvcmRlcjogMDtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCOzs7QUNEOUIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0saW1EQUF5bUQsQ0FBQztBQUFFLElBQU8sa0NBQVE7OztBREFockQ7QUFZQSxvQ0FBQyxjQUFjLHlCQUF5QjtBQUNqQyxJQUFNLHVCQUFOLGVBQW1DLGlCQUFXO0FBQUEsRUFDbkQsT0FBTyxTQUFTO0FBQUE7QUFBQSxFQUdoQixnQkFBZ0IsS0FBYSxjQUE4QjtBQUN6RCxRQUFJO0FBQ0YsYUFBTyxhQUFhLFFBQVEsR0FBRyxLQUFLO0FBQUEsSUFDdEMsUUFBUTtBQUNOLGFBQU87QUFBQSxJQUNUO0FBQUEsRUFDRjtBQUFBLEVBRUEsZ0JBQWdCLEtBQWEsT0FBcUI7QUFDaEQsUUFBSTtBQUNGLG1CQUFhLFFBQVEsS0FBSyxLQUFLO0FBQUEsSUFDakMsUUFBUTtBQUFBLElBRVI7QUFBQSxFQUNGO0FBQUEsRUFFQSxvQkFBMEI7QUFDeEIsVUFBTSxrQkFBa0I7QUFFeEIsVUFBTSxRQUFRLEtBQUssZ0JBQWdCLDBCQUEwQixRQUFRO0FBQ3JFLFNBQUssa0JBQWtCLEtBQUs7QUFBQSxFQUM5QjtBQUFBLEVBRUEsU0FBUztBQUNQLFVBQU0sUUFBUSxLQUFLLGdCQUFnQiwwQkFBMEIsUUFBUTtBQUNyRSxXQUFPO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSw0QkFRaUIsVUFBVSxPQUFPO0FBQUEsMkJBQ2xCLEtBQUssU0FBUztBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSw0QkFZYixVQUFVLFFBQVE7QUFBQSwyQkFDbkIsS0FBSyxTQUFTO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLDRCQVliLFVBQVUsTUFBTTtBQUFBLDJCQUNqQixLQUFLLFNBQVM7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxFQVV2QztBQUFBLEVBRUEsVUFBVSxHQUFnQjtBQUN4QixVQUFNLFNBQVMsRUFBRTtBQUNqQixRQUFJLE9BQU8sU0FBUztBQUNsQixXQUFLLGtCQUFrQixPQUFPLEtBQW9CO0FBQUEsSUFDcEQ7QUFBQSxFQUNGO0FBQUEsRUFFQSxrQkFBa0IsUUFBMkI7QUFDM0MsWUFBUSxRQUFRO0FBQUEsTUFDZCxLQUFLO0FBQ0gsaUJBQVMsS0FBSyxNQUFNLGNBQWM7QUFDbEM7QUFBQSxNQUNGLEtBQUs7QUFDSCxpQkFBUyxLQUFLLE1BQU0sY0FBYztBQUNsQztBQUFBLE1BQ0YsS0FBSztBQUFBLE1BQ0w7QUFDRSxpQkFBUyxLQUFLLE1BQU0sY0FBYztBQUNsQztBQUFBLElBQ0o7QUFDQSxTQUFLLGdCQUFnQiwwQkFBMEIsTUFBTTtBQUFBLEVBQ3ZEO0FBQ0Y7QUFsR087QUFBTSx1QkFBTixvREFEUCxrQ0FDYTtBQUFOLDRCQUFNOyIsCiAgIm5hbWVzIjogW10KfQo=
