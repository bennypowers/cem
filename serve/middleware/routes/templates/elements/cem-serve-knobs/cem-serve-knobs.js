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

// elements/cem-serve-knobs/cem-serve-knobs.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-serve-knobs/cem-serve-knobs.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: flex;\\n  flex-direction: column;\\n  gap: 0;\\n  height: 100%;\\n}\\n\\ncem-pf-v6-navigation {\\n  flex-shrink: 0;\\n}\\n\\n#knobs {\\n  flex: 1;\\n  min-height: 0;\\n  overflow-y: auto;\\n  padding: var(--pf-t--global--spacer--lg);\\n}\\n\\n#knobs slot {\\n  display: grid;\\n  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));\\n  gap: var(--pf-t--global--spacer--gap--group-to-group--horizontal--default) var(--pf-t--global--spacer--gap--group-to-group--vertical--default);\\n  align-items: start;\\n}\\n\\n::slotted(cem-pf-v6-card.active) {\\n  --cem-pf-v6-c-card--BorderColor: var(--cem-pf-v6-c-card--m-selectable--m-selected--BorderColor);\\n  --cem-pf-v6-c-card--BorderWidth: var(--cem-pf-v6-c-card--m-selectable--m-selected--BorderWidth);\\n}\\n"'));
var cem_serve_knobs_default = s;

// elements/cem-serve-knobs/cem-serve-knobs.ts
import "../cem-pf-v6-navigation/cem-pf-v6-navigation.js";
import "../cem-pf-v6-nav-list/cem-pf-v6-nav-list.js";
import "../cem-pf-v6-nav-item/cem-pf-v6-nav-item.js";
import "../cem-pf-v6-nav-link/cem-pf-v6-nav-link.js";
import "../cem-pf-v6-card/cem-pf-v6-card.js";
import "../cem-pf-v6-form-field-group/cem-pf-v6-form-field-group.js";
var _CemServeKnobs_decorators, _init, _a;
_CemServeKnobs_decorators = [customElement("cem-serve-knobs")];
var CemServeKnobs = class extends (_a = LitElement) {
  static styles = cem_serve_knobs_default;
  #navList = null;
  #handleHashChange = () => {
    const hash = window.location.hash;
    if (!hash || hash === "#") return;
    const cardId = hash.substring(1);
    const cards = this.querySelectorAll("cem-pf-v6-card");
    const navLinks = this.#navList?.querySelectorAll("cem-pf-v6-nav-link");
    const knobsContainer = this.shadowRoot?.getElementById("knobs");
    if (navLinks) {
      for (const link of navLinks) {
        const linkHref = link.getAttribute("href");
        link.toggleAttribute("current", linkHref === hash);
      }
    }
    for (const card of cards) {
      const isActive = card.dataset.card === cardId;
      card.classList.toggle("active", isActive);
      if (isActive && knobsContainer) {
        const cardRect = card.getBoundingClientRect();
        const containerRect = knobsContainer.getBoundingClientRect();
        const scrollOffset = cardRect.top - containerRect.top + knobsContainer.scrollTop;
        knobsContainer.scrollTo({
          top: scrollOffset,
          behavior: "smooth"
        });
      }
    }
  };
  render() {
    return html`
      <cem-pf-v6-navigation horizontal
                        variant="horizontal-subnav">
        <cem-pf-v6-nav-list id="nav-list"
                        aria-label="Elements">
        </cem-pf-v6-nav-list>
      </cem-pf-v6-navigation>
      <div id="knobs">
        <slot @slotchange=${this.#onSlotChange}></slot>
      </div>
    `;
  }
  connectedCallback() {
    super.connectedCallback();
    window.addEventListener("hashchange", this.#handleHashChange);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener("hashchange", this.#handleHashChange);
  }
  firstUpdated() {
    this.#navList = this.shadowRoot?.getElementById("nav-list") ?? null;
    this.#updateNavigation();
    this.#handleHashChange();
  }
  #onSlotChange() {
    this.#updateNavigation();
  }
  #updateNavigation() {
    const slot = this.shadowRoot?.querySelector("slot");
    if (!slot || !this.#navList) return;
    const panels = slot.assignedElements().filter((el) => el.tagName === "CEM-PF-V6-CARD");
    if (panels.length === 0) return;
    this.#navList.innerHTML = "";
    panels.forEach((panel, index) => {
      const navItem = document.createElement("cem-pf-v6-nav-item");
      const navLink = document.createElement("cem-pf-v6-nav-link");
      const label = document.createElement("span");
      const panelId = panel.dataset.card || `instance-${index}`;
      const labelText = panel.dataset.label || `Instance ${index + 1}`;
      label.className = "instance-label";
      label.textContent = labelText;
      navLink.setAttribute("href", `#${panelId}`);
      if (index === 0) {
        navLink.setAttribute("current", "");
      }
      navLink.appendChild(label);
      navItem.appendChild(navLink);
      this.#navList.appendChild(navItem);
    });
  }
};
_init = __decoratorStart(_a);
CemServeKnobs = __decorateElement(_init, 0, "CemServeKnobs", _CemServeKnobs_decorators, CemServeKnobs);
__runInitializers(_init, 1, CemServeKnobs);
export {
  CemServeKnobs
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWtub2JzL2NlbS1zZXJ2ZS1rbm9icy50cyIsICJsaXQtY3NzOmVsZW1lbnRzL2NlbS1zZXJ2ZS1rbm9icy9jZW0tc2VydmUta25vYnMuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tc2VydmUta25vYnMuY3NzJztcblxuaW1wb3J0ICcuLi9jZW0tcGYtdjYtbmF2aWdhdGlvbi9jZW0tcGYtdjYtbmF2aWdhdGlvbi5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1uYXYtbGlzdC9jZW0tcGYtdjYtbmF2LWxpc3QuanMnO1xuaW1wb3J0ICcuLi9jZW0tcGYtdjYtbmF2LWl0ZW0vY2VtLXBmLXY2LW5hdi1pdGVtLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LW5hdi1saW5rL2NlbS1wZi12Ni1uYXYtbGluay5qcyc7XG5pbXBvcnQgJy4uL2NlbS1wZi12Ni1jYXJkL2NlbS1wZi12Ni1jYXJkLmpzJztcbmltcG9ydCAnLi4vY2VtLXBmLXY2LWZvcm0tZmllbGQtZ3JvdXAvY2VtLXBmLXY2LWZvcm0tZmllbGQtZ3JvdXAuanMnO1xuXG4vKipcbiAqIENFTSBTZXJ2ZSBLbm9icyAtIENvbnRhaW5lciBmb3IgbXVsdGktaW5zdGFuY2Uga25vYnMgcGFuZWxzXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIGNlbS1wZi12Ni1jYXJkIGtub2IgcGFuZWxzXG4gKi9cbkBjdXN0b21FbGVtZW50KCdjZW0tc2VydmUta25vYnMnKVxuZXhwb3J0IGNsYXNzIENlbVNlcnZlS25vYnMgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAjbmF2TGlzdDogRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjaGFuZGxlSGFzaENoYW5nZSA9ICgpID0+IHtcbiAgICBjb25zdCBoYXNoID0gd2luZG93LmxvY2F0aW9uLmhhc2g7XG4gICAgaWYgKCFoYXNoIHx8IGhhc2ggPT09ICcjJykgcmV0dXJuO1xuXG4gICAgY29uc3QgY2FyZElkID0gaGFzaC5zdWJzdHJpbmcoMSk7XG4gICAgY29uc3QgY2FyZHMgPSB0aGlzLnF1ZXJ5U2VsZWN0b3JBbGwoJ2NlbS1wZi12Ni1jYXJkJyk7XG4gICAgY29uc3QgbmF2TGlua3MgPSB0aGlzLiNuYXZMaXN0Py5xdWVyeVNlbGVjdG9yQWxsKCdjZW0tcGYtdjYtbmF2LWxpbmsnKTtcbiAgICBjb25zdCBrbm9ic0NvbnRhaW5lciA9IHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ2tub2JzJyk7XG5cbiAgICBpZiAobmF2TGlua3MpIHtcbiAgICAgIGZvciAoY29uc3QgbGluayBvZiBuYXZMaW5rcykge1xuICAgICAgICBjb25zdCBsaW5rSHJlZiA9IGxpbmsuZ2V0QXR0cmlidXRlKCdocmVmJyk7XG4gICAgICAgIGxpbmsudG9nZ2xlQXR0cmlidXRlKCdjdXJyZW50JywgbGlua0hyZWYgPT09IGhhc2gpO1xuICAgICAgfVxuICAgIH1cblxuICAgIGZvciAoY29uc3QgY2FyZCBvZiBjYXJkcykge1xuICAgICAgY29uc3QgaXNBY3RpdmUgPSAoY2FyZCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5jYXJkID09PSBjYXJkSWQ7XG4gICAgICBjYXJkLmNsYXNzTGlzdC50b2dnbGUoJ2FjdGl2ZScsIGlzQWN0aXZlKTtcbiAgICAgIGlmIChpc0FjdGl2ZSAmJiBrbm9ic0NvbnRhaW5lcikge1xuICAgICAgICBjb25zdCBjYXJkUmVjdCA9IGNhcmQuZ2V0Qm91bmRpbmdDbGllbnRSZWN0KCk7XG4gICAgICAgIGNvbnN0IGNvbnRhaW5lclJlY3QgPSBrbm9ic0NvbnRhaW5lci5nZXRCb3VuZGluZ0NsaWVudFJlY3QoKTtcbiAgICAgICAgY29uc3Qgc2Nyb2xsT2Zmc2V0ID0gY2FyZFJlY3QudG9wIC0gY29udGFpbmVyUmVjdC50b3AgKyBrbm9ic0NvbnRhaW5lci5zY3JvbGxUb3A7XG4gICAgICAgIGtub2JzQ29udGFpbmVyLnNjcm9sbFRvKHtcbiAgICAgICAgICB0b3A6IHNjcm9sbE9mZnNldCxcbiAgICAgICAgICBiZWhhdmlvcjogJ3Ntb290aCcsXG4gICAgICAgIH0pO1xuICAgICAgfVxuICAgIH1cbiAgfTtcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8Y2VtLXBmLXY2LW5hdmlnYXRpb24gaG9yaXpvbnRhbFxuICAgICAgICAgICAgICAgICAgICAgICAgdmFyaWFudD1cImhvcml6b250YWwtc3VibmF2XCI+XG4gICAgICAgIDxjZW0tcGYtdjYtbmF2LWxpc3QgaWQ9XCJuYXYtbGlzdFwiXG4gICAgICAgICAgICAgICAgICAgICAgICBhcmlhLWxhYmVsPVwiRWxlbWVudHNcIj5cbiAgICAgICAgPC9jZW0tcGYtdjYtbmF2LWxpc3Q+XG4gICAgICA8L2NlbS1wZi12Ni1uYXZpZ2F0aW9uPlxuICAgICAgPGRpdiBpZD1cImtub2JzXCI+XG4gICAgICAgIDxzbG90IEBzbG90Y2hhbmdlPSR7dGhpcy4jb25TbG90Q2hhbmdlfT48L3Nsb3Q+XG4gICAgICA8L2Rpdj5cbiAgICBgO1xuICB9XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB3aW5kb3cuYWRkRXZlbnRMaXN0ZW5lcignaGFzaGNoYW5nZScsIHRoaXMuI2hhbmRsZUhhc2hDaGFuZ2UpO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB3aW5kb3cucmVtb3ZlRXZlbnRMaXN0ZW5lcignaGFzaGNoYW5nZScsIHRoaXMuI2hhbmRsZUhhc2hDaGFuZ2UpO1xuICB9XG5cbiAgZmlyc3RVcGRhdGVkKCkge1xuICAgIHRoaXMuI25hdkxpc3QgPSB0aGlzLnNoYWRvd1Jvb3Q/LmdldEVsZW1lbnRCeUlkKCduYXYtbGlzdCcpID8/IG51bGw7XG4gICAgdGhpcy4jdXBkYXRlTmF2aWdhdGlvbigpO1xuICAgIHRoaXMuI2hhbmRsZUhhc2hDaGFuZ2UoKTtcbiAgfVxuXG4gICNvblNsb3RDaGFuZ2UoKSB7XG4gICAgdGhpcy4jdXBkYXRlTmF2aWdhdGlvbigpO1xuICB9XG5cbiAgI3VwZGF0ZU5hdmlnYXRpb24oKSB7XG4gICAgY29uc3Qgc2xvdCA9IHRoaXMuc2hhZG93Um9vdD8ucXVlcnlTZWxlY3Rvcignc2xvdCcpO1xuICAgIGlmICghc2xvdCB8fCAhdGhpcy4jbmF2TGlzdCkgcmV0dXJuO1xuXG4gICAgY29uc3QgcGFuZWxzID0gc2xvdC5hc3NpZ25lZEVsZW1lbnRzKCkuZmlsdGVyKGVsID0+IGVsLnRhZ05hbWUgPT09ICdDRU0tUEYtVjYtQ0FSRCcpO1xuICAgIGlmIChwYW5lbHMubGVuZ3RoID09PSAwKSByZXR1cm47XG5cbiAgICB0aGlzLiNuYXZMaXN0LmlubmVySFRNTCA9ICcnO1xuXG4gICAgcGFuZWxzLmZvckVhY2goKHBhbmVsLCBpbmRleCkgPT4ge1xuICAgICAgY29uc3QgbmF2SXRlbSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2NlbS1wZi12Ni1uYXYtaXRlbScpO1xuICAgICAgY29uc3QgbmF2TGluayA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2NlbS1wZi12Ni1uYXYtbGluaycpO1xuICAgICAgY29uc3QgbGFiZWwgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdzcGFuJyk7XG5cbiAgICAgIGNvbnN0IHBhbmVsSWQgPSAocGFuZWwgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQuY2FyZCB8fCBgaW5zdGFuY2UtJHtpbmRleH1gO1xuICAgICAgY29uc3QgbGFiZWxUZXh0ID0gKHBhbmVsIGFzIEhUTUxFbGVtZW50KS5kYXRhc2V0LmxhYmVsIHx8IGBJbnN0YW5jZSAke2luZGV4ICsgMX1gO1xuXG4gICAgICBsYWJlbC5jbGFzc05hbWUgPSAnaW5zdGFuY2UtbGFiZWwnO1xuICAgICAgbGFiZWwudGV4dENvbnRlbnQgPSBsYWJlbFRleHQ7XG5cbiAgICAgIG5hdkxpbmsuc2V0QXR0cmlidXRlKCdocmVmJywgYCMke3BhbmVsSWR9YCk7XG4gICAgICBpZiAoaW5kZXggPT09IDApIHtcbiAgICAgICAgbmF2TGluay5zZXRBdHRyaWJ1dGUoJ2N1cnJlbnQnLCAnJyk7XG4gICAgICB9XG5cbiAgICAgIG5hdkxpbmsuYXBwZW5kQ2hpbGQobGFiZWwpO1xuICAgICAgbmF2SXRlbS5hcHBlbmRDaGlsZChuYXZMaW5rKTtcbiAgICAgIHRoaXMuI25hdkxpc3QhLmFwcGVuZENoaWxkKG5hdkl0ZW0pO1xuICAgIH0pO1xuICB9XG59XG5cbmRlY2xhcmUgZ2xvYmFsIHtcbiAgaW50ZXJmYWNlIEhUTUxFbGVtZW50VGFnTmFtZU1hcCB7XG4gICAgJ2NlbS1zZXJ2ZS1rbm9icyc6IENlbVNlcnZlS25vYnM7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogZmxleDtcXFxcbiAgZmxleC1kaXJlY3Rpb246IGNvbHVtbjtcXFxcbiAgZ2FwOiAwO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG59XFxcXG5cXFxcbmNlbS1wZi12Ni1uYXZpZ2F0aW9uIHtcXFxcbiAgZmxleC1zaHJpbms6IDA7XFxcXG59XFxcXG5cXFxcbiNrbm9icyB7XFxcXG4gIGZsZXg6IDE7XFxcXG4gIG1pbi1oZWlnaHQ6IDA7XFxcXG4gIG92ZXJmbG93LXk6IGF1dG87XFxcXG4gIHBhZGRpbmc6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1sZyk7XFxcXG59XFxcXG5cXFxcbiNrbm9icyBzbG90IHtcXFxcbiAgZGlzcGxheTogZ3JpZDtcXFxcbiAgZ3JpZC10ZW1wbGF0ZS1jb2x1bW5zOiByZXBlYXQoYXV0by1maWxsLCBtaW5tYXgoNDAwcHgsIDFmcikpO1xcXFxuICBnYXA6IHZhcigtLXBmLXQtLWdsb2JhbC0tc3BhY2VyLS1nYXAtLWdyb3VwLXRvLWdyb3VwLS1ob3Jpem9udGFsLS1kZWZhdWx0KSB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tZ2FwLS1ncm91cC10by1ncm91cC0tdmVydGljYWwtLWRlZmF1bHQpO1xcXFxuICBhbGlnbi1pdGVtczogc3RhcnQ7XFxcXG59XFxcXG5cXFxcbjo6c2xvdHRlZChjZW0tcGYtdjYtY2FyZC5hY3RpdmUpIHtcXFxcbiAgLS1jZW0tcGYtdjYtYy1jYXJkLS1Cb3JkZXJDb2xvcjogdmFyKC0tY2VtLXBmLXY2LWMtY2FyZC0tbS1zZWxlY3RhYmxlLS1tLXNlbGVjdGVkLS1Cb3JkZXJDb2xvcik7XFxcXG4gIC0tY2VtLXBmLXY2LWMtY2FyZC0tQm9yZGVyV2lkdGg6IHZhcigtLWNlbS1wZi12Ni1jLWNhcmQtLW0tc2VsZWN0YWJsZS0tbS1zZWxlY3RlZC0tQm9yZGVyV2lkdGgpO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7OztBQ0Q5QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSx1eEJBQXl4QixDQUFDO0FBQUUsSUFBTywwQkFBUTs7O0FES2gyQixPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFWUDtBQWlCQSw2QkFBQyxjQUFjLGlCQUFpQjtBQUN6QixJQUFNLGdCQUFOLGVBQTRCLGlCQUFXO0FBQUEsRUFDNUMsT0FBTyxTQUFTO0FBQUEsRUFFaEIsV0FBMkI7QUFBQSxFQUMzQixv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
