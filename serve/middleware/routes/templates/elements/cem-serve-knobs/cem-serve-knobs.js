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

// lit-css:/home/bennyp/Developer/cem/serve/elements/cem-serve-knobs/cem-serve-knobs.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: flex;\\n  flex-direction: column;\\n  gap: 0;\\n  height: 100%;\\n}\\n\\npf-v6-navigation {\\n  flex-shrink: 0;\\n}\\n\\n#knobs {\\n  flex: 1;\\n  min-height: 0;\\n  overflow-y: auto;\\n  padding: var(--pf-t--global--spacer--lg);\\n}\\n\\n#knobs slot {\\n  display: grid;\\n  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));\\n  gap: var(--pf-t--global--spacer--gap--group-to-group--horizontal--default) var(--pf-t--global--spacer--gap--group-to-group--vertical--default);\\n  align-items: start;\\n}\\n\\n::slotted(pf-v6-card.active) {\\n  --pf-v6-c-card--BorderColor: var(--pf-v6-c-card--m-selectable--m-selected--BorderColor);\\n  --pf-v6-c-card--BorderWidth: var(--pf-v6-c-card--m-selectable--m-selected--BorderWidth);\\n}\\n"'));
var cem_serve_knobs_default = s;

// elements/cem-serve-knobs/cem-serve-knobs.ts
import "../pf-v6-navigation/pf-v6-navigation.js";
import "../pf-v6-nav-list/pf-v6-nav-list.js";
import "../pf-v6-nav-item/pf-v6-nav-item.js";
import "../pf-v6-nav-link/pf-v6-nav-link.js";
import "../pf-v6-card/pf-v6-card.js";
import "../pf-v6-form-field-group/pf-v6-form-field-group.js";
var _CemServeKnobs_decorators, _init, _a;
_CemServeKnobs_decorators = [customElement("cem-serve-knobs")];
var CemServeKnobs = class extends (_a = LitElement) {
  static styles = cem_serve_knobs_default;
  #navList = null;
  #handleHashChange = () => {
    const hash = window.location.hash;
    if (!hash || hash === "#") return;
    const cardId = hash.substring(1);
    const cards = this.querySelectorAll("pf-v6-card");
    const navLinks = this.#navList?.querySelectorAll("pf-v6-nav-link");
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
      <pf-v6-navigation horizontal
                        variant="horizontal-subnav">
        <pf-v6-nav-list id="nav-list"
                        aria-label="Elements">
        </pf-v6-nav-list>
      </pf-v6-navigation>
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
    const panels = slot.assignedElements().filter((el) => el.tagName === "PF-V6-CARD");
    if (panels.length === 0) return;
    this.#navList.innerHTML = "";
    panels.forEach((panel, index) => {
      const navItem = document.createElement("pf-v6-nav-item");
      const navLink = document.createElement("pf-v6-nav-link");
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXNlcnZlLWtub2JzL2NlbS1zZXJ2ZS1rbm9icy50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL2NlbS1zZXJ2ZS1rbm9icy9jZW0tc2VydmUta25vYnMuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9jZW0tc2VydmUta25vYnMuY3NzJztcblxuaW1wb3J0ICcuLi9wZi12Ni1uYXZpZ2F0aW9uL3BmLXY2LW5hdmlnYXRpb24uanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1uYXYtbGlzdC9wZi12Ni1uYXYtbGlzdC5qcyc7XG5pbXBvcnQgJy4uL3BmLXY2LW5hdi1pdGVtL3BmLXY2LW5hdi1pdGVtLmpzJztcbmltcG9ydCAnLi4vcGYtdjYtbmF2LWxpbmsvcGYtdjYtbmF2LWxpbmsuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1jYXJkL3BmLXY2LWNhcmQuanMnO1xuaW1wb3J0ICcuLi9wZi12Ni1mb3JtLWZpZWxkLWdyb3VwL3BmLXY2LWZvcm0tZmllbGQtZ3JvdXAuanMnO1xuXG4vKipcbiAqIENFTSBTZXJ2ZSBLbm9icyAtIENvbnRhaW5lciBmb3IgbXVsdGktaW5zdGFuY2Uga25vYnMgcGFuZWxzXG4gKlxuICogQHNsb3QgLSBEZWZhdWx0IHNsb3QgZm9yIHBmLXY2LWNhcmQga25vYiBwYW5lbHNcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS1zZXJ2ZS1rbm9icycpXG5leHBvcnQgY2xhc3MgQ2VtU2VydmVLbm9icyBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gICNuYXZMaXN0OiBFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNoYW5kbGVIYXNoQ2hhbmdlID0gKCkgPT4ge1xuICAgIGNvbnN0IGhhc2ggPSB3aW5kb3cubG9jYXRpb24uaGFzaDtcbiAgICBpZiAoIWhhc2ggfHwgaGFzaCA9PT0gJyMnKSByZXR1cm47XG5cbiAgICBjb25zdCBjYXJkSWQgPSBoYXNoLnN1YnN0cmluZygxKTtcbiAgICBjb25zdCBjYXJkcyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtY2FyZCcpO1xuICAgIGNvbnN0IG5hdkxpbmtzID0gdGhpcy4jbmF2TGlzdD8ucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtbmF2LWxpbmsnKTtcbiAgICBjb25zdCBrbm9ic0NvbnRhaW5lciA9IHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ2tub2JzJyk7XG5cbiAgICBpZiAobmF2TGlua3MpIHtcbiAgICAgIGZvciAoY29uc3QgbGluayBvZiBuYXZMaW5rcykge1xuICAgICAgICBjb25zdCBsaW5rSHJlZiA9IGxpbmsuZ2V0QXR0cmlidXRlKCdocmVmJyk7XG4gICAgICAgIGxpbmsudG9nZ2xlQXR0cmlidXRlKCdjdXJyZW50JywgbGlua0hyZWYgPT09IGhhc2gpO1xuICAgICAgfVxuICAgIH1cblxuICAgIGZvciAoY29uc3QgY2FyZCBvZiBjYXJkcykge1xuICAgICAgY29uc3QgaXNBY3RpdmUgPSAoY2FyZCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5jYXJkID09PSBjYXJkSWQ7XG4gICAgICBjYXJkLmNsYXNzTGlzdC50b2dnbGUoJ2FjdGl2ZScsIGlzQWN0aXZlKTtcbiAgICAgIGlmIChpc0FjdGl2ZSAmJiBrbm9ic0NvbnRhaW5lcikge1xuICAgICAgICBjb25zdCBjYXJkUmVjdCA9IGNhcmQuZ2V0Qm91bmRpbmdDbGllbnRSZWN0KCk7XG4gICAgICAgIGNvbnN0IGNvbnRhaW5lclJlY3QgPSBrbm9ic0NvbnRhaW5lci5nZXRCb3VuZGluZ0NsaWVudFJlY3QoKTtcbiAgICAgICAgY29uc3Qgc2Nyb2xsT2Zmc2V0ID0gY2FyZFJlY3QudG9wIC0gY29udGFpbmVyUmVjdC50b3AgKyBrbm9ic0NvbnRhaW5lci5zY3JvbGxUb3A7XG4gICAgICAgIGtub2JzQ29udGFpbmVyLnNjcm9sbFRvKHtcbiAgICAgICAgICB0b3A6IHNjcm9sbE9mZnNldCxcbiAgICAgICAgICBiZWhhdmlvcjogJ3Ntb290aCcsXG4gICAgICAgIH0pO1xuICAgICAgfVxuICAgIH1cbiAgfTtcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8cGYtdjYtbmF2aWdhdGlvbiBob3Jpem9udGFsXG4gICAgICAgICAgICAgICAgICAgICAgICB2YXJpYW50PVwiaG9yaXpvbnRhbC1zdWJuYXZcIj5cbiAgICAgICAgPHBmLXY2LW5hdi1saXN0IGlkPVwibmF2LWxpc3RcIlxuICAgICAgICAgICAgICAgICAgICAgICAgYXJpYS1sYWJlbD1cIkVsZW1lbnRzXCI+XG4gICAgICAgIDwvcGYtdjYtbmF2LWxpc3Q+XG4gICAgICA8L3BmLXY2LW5hdmlnYXRpb24+XG4gICAgICA8ZGl2IGlkPVwia25vYnNcIj5cbiAgICAgICAgPHNsb3QgQHNsb3RjaGFuZ2U9JHt0aGlzLiNvblNsb3RDaGFuZ2V9Pjwvc2xvdD5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHdpbmRvdy5hZGRFdmVudExpc3RlbmVyKCdoYXNoY2hhbmdlJywgdGhpcy4jaGFuZGxlSGFzaENoYW5nZSk7XG4gIH1cblxuICBkaXNjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5kaXNjb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHdpbmRvdy5yZW1vdmVFdmVudExpc3RlbmVyKCdoYXNoY2hhbmdlJywgdGhpcy4jaGFuZGxlSGFzaENoYW5nZSk7XG4gIH1cblxuICBmaXJzdFVwZGF0ZWQoKSB7XG4gICAgdGhpcy4jbmF2TGlzdCA9IHRoaXMuc2hhZG93Um9vdD8uZ2V0RWxlbWVudEJ5SWQoJ25hdi1saXN0JykgPz8gbnVsbDtcbiAgICB0aGlzLiN1cGRhdGVOYXZpZ2F0aW9uKCk7XG4gICAgdGhpcy4jaGFuZGxlSGFzaENoYW5nZSgpO1xuICB9XG5cbiAgI29uU2xvdENoYW5nZSgpIHtcbiAgICB0aGlzLiN1cGRhdGVOYXZpZ2F0aW9uKCk7XG4gIH1cblxuICAjdXBkYXRlTmF2aWdhdGlvbigpIHtcbiAgICBjb25zdCBzbG90ID0gdGhpcy5zaGFkb3dSb290Py5xdWVyeVNlbGVjdG9yKCdzbG90Jyk7XG4gICAgaWYgKCFzbG90IHx8ICF0aGlzLiNuYXZMaXN0KSByZXR1cm47XG5cbiAgICBjb25zdCBwYW5lbHMgPSBzbG90LmFzc2lnbmVkRWxlbWVudHMoKS5maWx0ZXIoZWwgPT4gZWwudGFnTmFtZSA9PT0gJ1BGLVY2LUNBUkQnKTtcbiAgICBpZiAocGFuZWxzLmxlbmd0aCA9PT0gMCkgcmV0dXJuO1xuXG4gICAgdGhpcy4jbmF2TGlzdC5pbm5lckhUTUwgPSAnJztcblxuICAgIHBhbmVscy5mb3JFYWNoKChwYW5lbCwgaW5kZXgpID0+IHtcbiAgICAgIGNvbnN0IG5hdkl0ZW0gPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdwZi12Ni1uYXYtaXRlbScpO1xuICAgICAgY29uc3QgbmF2TGluayA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3BmLXY2LW5hdi1saW5rJyk7XG4gICAgICBjb25zdCBsYWJlbCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3NwYW4nKTtcblxuICAgICAgY29uc3QgcGFuZWxJZCA9IChwYW5lbCBhcyBIVE1MRWxlbWVudCkuZGF0YXNldC5jYXJkIHx8IGBpbnN0YW5jZS0ke2luZGV4fWA7XG4gICAgICBjb25zdCBsYWJlbFRleHQgPSAocGFuZWwgYXMgSFRNTEVsZW1lbnQpLmRhdGFzZXQubGFiZWwgfHwgYEluc3RhbmNlICR7aW5kZXggKyAxfWA7XG5cbiAgICAgIGxhYmVsLmNsYXNzTmFtZSA9ICdpbnN0YW5jZS1sYWJlbCc7XG4gICAgICBsYWJlbC50ZXh0Q29udGVudCA9IGxhYmVsVGV4dDtcblxuICAgICAgbmF2TGluay5zZXRBdHRyaWJ1dGUoJ2hyZWYnLCBgIyR7cGFuZWxJZH1gKTtcbiAgICAgIGlmIChpbmRleCA9PT0gMCkge1xuICAgICAgICBuYXZMaW5rLnNldEF0dHJpYnV0ZSgnY3VycmVudCcsICcnKTtcbiAgICAgIH1cblxuICAgICAgbmF2TGluay5hcHBlbmRDaGlsZChsYWJlbCk7XG4gICAgICBuYXZJdGVtLmFwcGVuZENoaWxkKG5hdkxpbmspO1xuICAgICAgdGhpcy4jbmF2TGlzdCEuYXBwZW5kQ2hpbGQobmF2SXRlbSk7XG4gICAgfSk7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXNlcnZlLWtub2JzJzogQ2VtU2VydmVLbm9icztcbiAgfVxufVxuIiwgImNvbnN0IHM9bmV3IENTU1N0eWxlU2hlZXQoKTtzLnJlcGxhY2VTeW5jKEpTT04ucGFyc2UoXCJcXFwiOmhvc3Qge1xcXFxuICBkaXNwbGF5OiBmbGV4O1xcXFxuICBmbGV4LWRpcmVjdGlvbjogY29sdW1uO1xcXFxuICBnYXA6IDA7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbn1cXFxcblxcXFxucGYtdjYtbmF2aWdhdGlvbiB7XFxcXG4gIGZsZXgtc2hyaW5rOiAwO1xcXFxufVxcXFxuXFxcXG4ja25vYnMge1xcXFxuICBmbGV4OiAxO1xcXFxuICBtaW4taGVpZ2h0OiAwO1xcXFxuICBvdmVyZmxvdy15OiBhdXRvO1xcXFxuICBwYWRkaW5nOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tbGcpO1xcXFxufVxcXFxuXFxcXG4ja25vYnMgc2xvdCB7XFxcXG4gIGRpc3BsYXk6IGdyaWQ7XFxcXG4gIGdyaWQtdGVtcGxhdGUtY29sdW1uczogcmVwZWF0KGF1dG8tZmlsbCwgbWlubWF4KDQwMHB4LCAxZnIpKTtcXFxcbiAgZ2FwOiB2YXIoLS1wZi10LS1nbG9iYWwtLXNwYWNlci0tZ2FwLS1ncm91cC10by1ncm91cC0taG9yaXpvbnRhbC0tZGVmYXVsdCkgdmFyKC0tcGYtdC0tZ2xvYmFsLS1zcGFjZXItLWdhcC0tZ3JvdXAtdG8tZ3JvdXAtLXZlcnRpY2FsLS1kZWZhdWx0KTtcXFxcbiAgYWxpZ24taXRlbXM6IHN0YXJ0O1xcXFxufVxcXFxuXFxcXG46OnNsb3R0ZWQocGYtdjYtY2FyZC5hY3RpdmUpIHtcXFxcbiAgLS1wZi12Ni1jLWNhcmQtLUJvcmRlckNvbG9yOiB2YXIoLS1wZi12Ni1jLWNhcmQtLW0tc2VsZWN0YWJsZS0tbS1zZWxlY3RlZC0tQm9yZGVyQ29sb3IpO1xcXFxuICAtLXBmLXY2LWMtY2FyZC0tQm9yZGVyV2lkdGg6IHZhcigtLXBmLXY2LWMtY2FyZC0tbS1zZWxlY3RhYmxlLS1tLXNlbGVjdGVkLS1Cb3JkZXJXaWR0aCk7XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjs7O0FDRDlCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLCt2QkFBaXdCLENBQUM7QUFBRSxJQUFPLDBCQUFROzs7QURLeDBCLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQUNQLE9BQU87QUFDUCxPQUFPO0FBQ1AsT0FBTztBQVZQO0FBaUJBLDZCQUFDLGNBQWMsaUJBQWlCO0FBQ3pCLElBQU0sZ0JBQU4sZUFBNEIsaUJBQVc7QUFBQSxFQUM1QyxPQUFPLFNBQVM7QUFBQSxFQUVoQixXQUEyQjtBQUFBLEVBQzNCLG9CQUFvQixNQUFNO0FBQ3hCLFVBQU0sT0FBTyxPQUFPLFNBQVM7QUFDN0IsUUFBSSxDQUFDLFFBQVEsU0FBUyxJQUFLO0FBRTNCLFVBQU0s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
