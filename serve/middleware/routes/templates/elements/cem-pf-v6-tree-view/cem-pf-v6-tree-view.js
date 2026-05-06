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

// elements/cem-pf-v6-tree-view/cem-pf-v6-tree-view.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-pf-v6-tree-view/cem-pf-v6-tree-view.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\n#tree {\\n  margin: 0;\\n  padding: 0;\\n  list-style: none;\\n}\\n"'));
var cem_pf_v6_tree_view_default = s;

// elements/cem-pf-v6-tree-view/cem-pf-v6-tree-view.ts
var _PfV6TreeView_decorators, _init, _a;
_PfV6TreeView_decorators = [customElement("cem-pf-v6-tree-view")];
var PfV6TreeView = class extends (_a = LitElement) {
  static styles = cem_pf_v6_tree_view_default;
  #currentSelection = null;
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("select", this.#onSelect);
    this.addEventListener("keydown", this.#onKeydown);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("select", this.#onSelect);
    this.removeEventListener("keydown", this.#onKeydown);
  }
  firstUpdated() {
    this.#initializeTabindex();
  }
  render() {
    return html`
      <ul id="tree"
          role="tree"
          part="tree"
          aria-label=${this.getAttribute("aria-label") ?? ""}>
        <slot></slot>
      </ul>
    `;
  }
  #onSelect = (event) => {
    const item = event.target;
    if (this.#currentSelection && this.#currentSelection !== item) {
      this.#currentSelection.deselect?.();
    }
    this.#currentSelection = item;
  };
  #onKeydown = (event) => {
    const target = event.target;
    if (target.tagName !== "CEM-PF-V6-TREE-ITEM") return;
    const treeItem = target;
    const items = this.#getAllVisibleItems();
    const currentIndex = items.indexOf(target);
    switch (event.key) {
      case "ArrowDown":
        event.preventDefault();
        if (currentIndex < items.length - 1) {
          this.#focusItem(items[currentIndex + 1]);
        }
        break;
      case "ArrowUp":
        event.preventDefault();
        if (currentIndex > 0) {
          this.#focusItem(items[currentIndex - 1]);
        }
        break;
      case "ArrowRight":
        event.preventDefault();
        if (treeItem.hasChildren) {
          if (!treeItem.expanded) {
            treeItem.expand?.();
          } else {
            const children = this.#getDirectChildren(target);
            if (children.length > 0) {
              this.#focusItem(children[0]);
            }
          }
        }
        break;
      case "ArrowLeft":
        event.preventDefault();
        if (treeItem.hasChildren && treeItem.expanded) {
          treeItem.collapse?.();
        } else {
          const parent = this.#getParentItem(target);
          if (parent) {
            this.#focusItem(parent);
          }
        }
        break;
      case "Home":
        event.preventDefault();
        if (items.length > 0) {
          this.#focusItem(items[0]);
        }
        break;
      case "End":
        event.preventDefault();
        if (items.length > 0) {
          this.#focusItem(items[items.length - 1]);
        }
        break;
      case "Enter":
      case " ":
        event.preventDefault();
        treeItem.select?.();
        if (treeItem.hasChildren) {
          treeItem.toggle?.();
        }
        break;
      case "*": {
        event.preventDefault();
        const parent = this.#getParentItem(target);
        const siblings = parent ? this.#getDirectChildren(parent) : this.#getTopLevelItems();
        siblings.forEach((item) => {
          if (item.hasChildren) {
            item.expand?.();
          }
        });
        break;
      }
    }
  };
  #getAllVisibleItems() {
    const visible = [];
    const walk = (parent) => {
      const children = Array.from(parent.children).filter((el) => el.tagName === "CEM-PF-V6-TREE-ITEM");
      for (const item of children) {
        visible.push(item);
        const treeItem = item;
        if (treeItem.expanded && treeItem.hasChildren) {
          walk(item);
        }
      }
    };
    walk(this);
    return visible;
  }
  #getDirectChildren(item) {
    return Array.from(item.children).filter((el) => el.tagName === "CEM-PF-V6-TREE-ITEM");
  }
  #getTopLevelItems() {
    return Array.from(this.children).filter((el) => el.tagName === "CEM-PF-V6-TREE-ITEM");
  }
  #getParentItem(item) {
    let current = item.parentElement;
    while (current && current !== this) {
      if (current.tagName === "CEM-PF-V6-TREE-ITEM") {
        return current;
      }
      current = current.parentElement;
    }
    return null;
  }
  #initializeTabindex() {
    const topLevelItems = this.#getTopLevelItems();
    if (topLevelItems.length === 0) return;
    const firstItem = topLevelItems[0];
    firstItem.setTabindex?.(0);
    const allItems = this.querySelectorAll("cem-pf-v6-tree-item");
    allItems.forEach((item) => {
      if (item !== firstItem) {
        item.setTabindex?.(-1);
      }
    });
  }
  #focusItem(item) {
    if (!item) return;
    const allItems = this.querySelectorAll("cem-pf-v6-tree-item");
    allItems.forEach((i) => i.setTabindex?.(-1));
    item.setTabindex?.(0);
    item.focusItem?.();
  }
  expandAll() {
    this.querySelectorAll("cem-pf-v6-tree-item").forEach(
      (item) => item.expand?.()
    );
  }
  collapseAll() {
    this.querySelectorAll("cem-pf-v6-tree-item").forEach(
      (item) => item.collapse?.()
    );
  }
};
_init = __decoratorStart(_a);
PfV6TreeView = __decorateElement(_init, 0, "PfV6TreeView", _PfV6TreeView_decorators, PfV6TreeView);
__runInitializers(_init, 1, PfV6TreeView);
export {
  PfV6TreeView
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXBmLXY2LXRyZWUtdmlldy9jZW0tcGYtdjYtdHJlZS12aWV3LnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXBmLXY2LXRyZWUtdmlldy9jZW0tcGYtdjYtdHJlZS12aWV3LmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXBmLXY2LXRyZWUtdmlldy5jc3MnO1xuXG4vKipcbiAqIFBhdHRlcm5GbHkgdjYgVHJlZSBWaWV3XG4gKlxuICogQ29udGFpbmVyIGZvciBjZW0tcGYtdjYtdHJlZS1pdGVtIGVsZW1lbnRzLiBDb29yZGluYXRlcyBzZWxlY3Rpb24gYW5kIGtleWJvYXJkIG5hdmlnYXRpb24uXG4gKlxuICogQHNsb3QgLSBUcmVlIGl0ZW1zIChjZW0tcGYtdjYtdHJlZS1pdGVtIGVsZW1lbnRzKVxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXBmLXY2LXRyZWUtdmlldycpXG5leHBvcnQgY2xhc3MgUGZWNlRyZWVWaWV3IGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgI2N1cnJlbnRTZWxlY3Rpb246IEVsZW1lbnQgfCBudWxsID0gbnVsbDtcblxuICBjb25uZWN0ZWRDYWxsYmFjaygpIHtcbiAgICBzdXBlci5jb25uZWN0ZWRDYWxsYmFjaygpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcignc2VsZWN0JywgdGhpcy4jb25TZWxlY3QgYXMgRXZlbnRMaXN0ZW5lcik7XG4gICAgdGhpcy5hZGRFdmVudExpc3RlbmVyKCdrZXlkb3duJywgdGhpcy4jb25LZXlkb3duIGFzIEV2ZW50TGlzdGVuZXIpO1xuICB9XG5cbiAgZGlzY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuZGlzY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI29uU2VsZWN0IGFzIEV2ZW50TGlzdGVuZXIpO1xuICAgIHRoaXMucmVtb3ZlRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI29uS2V5ZG93biBhcyBFdmVudExpc3RlbmVyKTtcbiAgfVxuXG4gIGZpcnN0VXBkYXRlZCgpIHtcbiAgICB0aGlzLiNpbml0aWFsaXplVGFiaW5kZXgoKTtcbiAgfVxuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDx1bCBpZD1cInRyZWVcIlxuICAgICAgICAgIHJvbGU9XCJ0cmVlXCJcbiAgICAgICAgICBwYXJ0PVwidHJlZVwiXG4gICAgICAgICAgYXJpYS1sYWJlbD0ke3RoaXMuZ2V0QXR0cmlidXRlKCdhcmlhLWxhYmVsJykgPz8gJyd9PlxuICAgICAgICA8c2xvdD48L3Nsb3Q+XG4gICAgICA8L3VsPlxuICAgIGA7XG4gIH1cblxuICAjb25TZWxlY3QgPSAoZXZlbnQ6IEV2ZW50KSA9PiB7XG4gICAgY29uc3QgaXRlbSA9IGV2ZW50LnRhcmdldCBhcyBUcmVlSXRlbUxpa2U7XG4gICAgaWYgKHRoaXMuI2N1cnJlbnRTZWxlY3Rpb24gJiYgdGhpcy4jY3VycmVudFNlbGVjdGlvbiAhPT0gaXRlbSkge1xuICAgICAgKHRoaXMuI2N1cnJlbnRTZWxlY3Rpb24gYXMgVHJlZUl0ZW1MaWtlKS5kZXNlbGVjdD8uKCk7XG4gICAgfVxuICAgIHRoaXMuI2N1cnJlbnRTZWxlY3Rpb24gPSBpdGVtO1xuICB9O1xuXG4gICNvbktleWRvd24gPSAoZXZlbnQ6IEtleWJvYXJkRXZlbnQpID0+IHtcbiAgICBjb25zdCB0YXJnZXQgPSBldmVudC50YXJnZXQgYXMgRWxlbWVudDtcbiAgICBpZiAodGFyZ2V0LnRhZ05hbWUgIT09ICdDRU0tUEYtVjYtVFJFRS1JVEVNJykgcmV0dXJuO1xuXG4gICAgY29uc3QgdHJlZUl0ZW0gPSB0YXJnZXQgYXMgVHJlZUl0ZW1MaWtlO1xuICAgIGNvbnN0IGl0ZW1zID0gdGhpcy4jZ2V0QWxsVmlzaWJsZUl0ZW1zKCk7XG4gICAgY29uc3QgY3VycmVudEluZGV4ID0gaXRlbXMuaW5kZXhPZih0YXJnZXQpO1xuXG4gICAgc3dpdGNoIChldmVudC5rZXkpIHtcbiAgICAgIGNhc2UgJ0Fycm93RG93bic6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmIChjdXJyZW50SW5kZXggPCBpdGVtcy5sZW5ndGggLSAxKSB7XG4gICAgICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zW2N1cnJlbnRJbmRleCArIDFdKTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnQXJyb3dVcCc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmIChjdXJyZW50SW5kZXggPiAwKSB7XG4gICAgICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zW2N1cnJlbnRJbmRleCAtIDFdKTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnQXJyb3dSaWdodCc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmICh0cmVlSXRlbS5oYXNDaGlsZHJlbikge1xuICAgICAgICAgIGlmICghdHJlZUl0ZW0uZXhwYW5kZWQpIHtcbiAgICAgICAgICAgIHRyZWVJdGVtLmV4cGFuZD8uKCk7XG4gICAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgIGNvbnN0IGNoaWxkcmVuID0gdGhpcy4jZ2V0RGlyZWN0Q2hpbGRyZW4odGFyZ2V0KTtcbiAgICAgICAgICAgIGlmIChjaGlsZHJlbi5sZW5ndGggPiAwKSB7XG4gICAgICAgICAgICAgIHRoaXMuI2ZvY3VzSXRlbShjaGlsZHJlblswXSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICAgIGJyZWFrO1xuXG4gICAgICBjYXNlICdBcnJvd0xlZnQnOlxuICAgICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBpZiAodHJlZUl0ZW0uaGFzQ2hpbGRyZW4gJiYgdHJlZUl0ZW0uZXhwYW5kZWQpIHtcbiAgICAgICAgICB0cmVlSXRlbS5jb2xsYXBzZT8uKCk7XG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgY29uc3QgcGFyZW50ID0gdGhpcy4jZ2V0UGFyZW50SXRlbSh0YXJnZXQpO1xuICAgICAgICAgIGlmIChwYXJlbnQpIHtcbiAgICAgICAgICAgIHRoaXMuI2ZvY3VzSXRlbShwYXJlbnQpO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnSG9tZSc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmIChpdGVtcy5sZW5ndGggPiAwKSB7XG4gICAgICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zWzBdKTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnRW5kJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgaWYgKGl0ZW1zLmxlbmd0aCA+IDApIHtcbiAgICAgICAgICB0aGlzLiNmb2N1c0l0ZW0oaXRlbXNbaXRlbXMubGVuZ3RoIC0gMV0pO1xuICAgICAgICB9XG4gICAgICAgIGJyZWFrO1xuXG4gICAgICBjYXNlICdFbnRlcic6XG4gICAgICBjYXNlICcgJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgdHJlZUl0ZW0uc2VsZWN0Py4oKTtcbiAgICAgICAgaWYgKHRyZWVJdGVtLmhhc0NoaWxkcmVuKSB7XG4gICAgICAgICAgdHJlZUl0ZW0udG9nZ2xlPy4oKTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnKic6IHtcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgY29uc3QgcGFyZW50ID0gdGhpcy4jZ2V0UGFyZW50SXRlbSh0YXJnZXQpO1xuICAgICAgICBjb25zdCBzaWJsaW5ncyA9IHBhcmVudCA/IHRoaXMuI2dldERpcmVjdENoaWxkcmVuKHBhcmVudCkgOiB0aGlzLiNnZXRUb3BMZXZlbEl0ZW1zKCk7XG4gICAgICAgIHNpYmxpbmdzLmZvckVhY2goaXRlbSA9PiB7XG4gICAgICAgICAgaWYgKChpdGVtIGFzIFRyZWVJdGVtTGlrZSkuaGFzQ2hpbGRyZW4pIHtcbiAgICAgICAgICAgIChpdGVtIGFzIFRyZWVJdGVtTGlrZSkuZXhwYW5kPy4oKTtcbiAgICAgICAgICB9XG4gICAgICAgIH0pO1xuICAgICAgICBicmVhaztcbiAgICAgIH1cbiAgICB9XG4gIH07XG5cbiAgI2dldEFsbFZpc2libGVJdGVtcygpOiBFbGVtZW50W10ge1xuICAgIGNvbnN0IHZpc2libGU6IEVsZW1lbnRbXSA9IFtdO1xuICAgIGNvbnN0IHdhbGsgPSAocGFyZW50OiBFbGVtZW50KSA9PiB7XG4gICAgICBjb25zdCBjaGlsZHJlbiA9IEFycmF5LmZyb20ocGFyZW50LmNoaWxkcmVuKS5maWx0ZXIoZWwgPT4gZWwudGFnTmFtZSA9PT0gJ0NFTS1QRi1WNi1UUkVFLUlURU0nKTtcbiAgICAgIGZvciAoY29uc3QgaXRlbSBvZiBjaGlsZHJlbikge1xuICAgICAgICB2aXNpYmxlLnB1c2goaXRlbSk7XG4gICAgICAgIGNvbnN0IHRyZWVJdGVtID0gaXRlbSBhcyBUcmVlSXRlbUxpa2U7XG4gICAgICAgIGlmICh0cmVlSXRlbS5leHBhbmRlZCAmJiB0cmVlSXRlbS5oYXNDaGlsZHJlbikge1xuICAgICAgICAgIHdhbGsoaXRlbSk7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9O1xuICAgIHdhbGsodGhpcyk7XG4gICAgcmV0dXJuIHZpc2libGU7XG4gIH1cblxuICAjZ2V0RGlyZWN0Q2hpbGRyZW4oaXRlbTogRWxlbWVudCk6IEVsZW1lbnRbXSB7XG4gICAgcmV0dXJuIEFycmF5LmZyb20oaXRlbS5jaGlsZHJlbikuZmlsdGVyKGVsID0+IGVsLnRhZ05hbWUgPT09ICdDRU0tUEYtVjYtVFJFRS1JVEVNJyk7XG4gIH1cblxuICAjZ2V0VG9wTGV2ZWxJdGVtcygpOiBFbGVtZW50W10ge1xuICAgIHJldHVybiBBcnJheS5mcm9tKHRoaXMuY2hpbGRyZW4pLmZpbHRlcihlbCA9PiBlbC50YWdOYW1lID09PSAnQ0VNLVBGLVY2LVRSRUUtSVRFTScpO1xuICB9XG5cbiAgI2dldFBhcmVudEl0ZW0oaXRlbTogRWxlbWVudCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICBsZXQgY3VycmVudCA9IGl0ZW0ucGFyZW50RWxlbWVudDtcbiAgICB3aGlsZSAoY3VycmVudCAmJiBjdXJyZW50ICE9PSB0aGlzKSB7XG4gICAgICBpZiAoY3VycmVudC50YWdOYW1lID09PSAnQ0VNLVBGLVY2LVRSRUUtSVRFTScpIHtcbiAgICAgICAgcmV0dXJuIGN1cnJlbnQ7XG4gICAgICB9XG4gICAgICBjdXJyZW50ID0gY3VycmVudC5wYXJlbnRFbGVtZW50O1xuICAgIH1cbiAgICByZXR1cm4gbnVsbDtcbiAgfVxuXG4gICNpbml0aWFsaXplVGFiaW5kZXgoKSB7XG4gICAgY29uc3QgdG9wTGV2ZWxJdGVtcyA9IHRoaXMuI2dldFRvcExldmVsSXRlbXMoKTtcbiAgICBpZiAodG9wTGV2ZWxJdGVtcy5sZW5ndGggPT09IDApIHJldHVybjtcbiAgICBjb25zdCBmaXJzdEl0ZW0gPSB0b3BMZXZlbEl0ZW1zWzBdIGFzIFRyZWVJdGVtTGlrZTtcbiAgICBmaXJzdEl0ZW0uc2V0VGFiaW5kZXg/LigwKTtcbiAgICBjb25zdCBhbGxJdGVtcyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCgnY2VtLXBmLXY2LXRyZWUtaXRlbScpO1xuICAgIGFsbEl0ZW1zLmZvckVhY2goaXRlbSA9PiB7XG4gICAgICBpZiAoaXRlbSAhPT0gZmlyc3RJdGVtKSB7XG4gICAgICAgIChpdGVtIGFzIHVua25vd24gYXMgVHJlZUl0ZW1MaWtlKS5zZXRUYWJpbmRleD8uKC0xKTtcbiAgICAgIH1cbiAgICB9KTtcbiAgfVxuXG4gICNmb2N1c0l0ZW0oaXRlbTogRWxlbWVudCkge1xuICAgIGlmICghaXRlbSkgcmV0dXJuO1xuICAgIGNvbnN0IGFsbEl0ZW1zID0gdGhpcy5xdWVyeVNlbGVjdG9yQWxsKCdjZW0tcGYtdjYtdHJlZS1pdGVtJyk7XG4gICAgYWxsSXRlbXMuZm9yRWFjaChpID0+IChpIGFzIHVua25vd24gYXMgVHJlZUl0ZW1MaWtlKS5zZXRUYWJpbmRleD8uKC0xKSk7XG4gICAgKGl0ZW0gYXMgVHJlZUl0ZW1MaWtlKS5zZXRUYWJpbmRleD8uKDApO1xuICAgIChpdGVtIGFzIFRyZWVJdGVtTGlrZSkuZm9jdXNJdGVtPy4oKTtcbiAgfVxuXG4gIGV4cGFuZEFsbCgpIHtcbiAgICB0aGlzLnF1ZXJ5U2VsZWN0b3JBbGwoJ2NlbS1wZi12Ni10cmVlLWl0ZW0nKS5mb3JFYWNoKGl0ZW0gPT5cbiAgICAgIChpdGVtIGFzIHVua25vd24gYXMgVHJlZUl0ZW1MaWtlKS5leHBhbmQ/LigpXG4gICAgKTtcbiAgfVxuXG4gIGNvbGxhcHNlQWxsKCkge1xuICAgIHRoaXMucXVlcnlTZWxlY3RvckFsbCgnY2VtLXBmLXY2LXRyZWUtaXRlbScpLmZvckVhY2goaXRlbSA9PlxuICAgICAgKGl0ZW0gYXMgdW5rbm93biBhcyBUcmVlSXRlbUxpa2UpLmNvbGxhcHNlPy4oKVxuICAgICk7XG4gIH1cbn1cblxuLyoqIEludGVyZmFjZSBmb3IgY2VtLXBmLXY2LXRyZWUtaXRlbSBtZXRob2RzIHVzZWQgYnkgdHJlZS12aWV3ICovXG5pbnRlcmZhY2UgVHJlZUl0ZW1MaWtlIGV4dGVuZHMgRWxlbWVudCB7XG4gIGhhc0NoaWxkcmVuPzogYm9vbGVhbjtcbiAgZXhwYW5kZWQ/OiBib29sZWFuO1xuICBleHBhbmQ/KCk6IHZvaWQ7XG4gIGNvbGxhcHNlPygpOiB2b2lkO1xuICB0b2dnbGU/KCk6IHZvaWQ7XG4gIHNlbGVjdD8oKTogdm9pZDtcbiAgZGVzZWxlY3Q/KCk6IHZvaWQ7XG4gIHNldFRhYmluZGV4Pyh2YWx1ZTogbnVtYmVyKTogdm9pZDtcbiAgZm9jdXNJdGVtPygpOiB2b2lkO1xufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdjZW0tcGYtdjYtdHJlZS12aWV3JzogUGZWNlRyZWVWaWV3O1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCI6aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxufVxcXFxuXFxcXG4jdHJlZSB7XFxcXG4gIG1hcmdpbjogMDtcXFxcbiAgcGFkZGluZzogMDtcXFxcbiAgbGlzdC1zdHlsZTogbm9uZTtcXFxcbn1cXFxcblxcXCJcIikpO2V4cG9ydCBkZWZhdWx0IHM7Il0sCiAgIm1hcHBpbmdzIjogIjs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQUFBLFNBQVMsWUFBWSxZQUFZO0FBQ2pDLFNBQVMscUJBQXFCOzs7QUNEOUIsSUFBTSxJQUFFLElBQUksY0FBYztBQUFFLEVBQUUsWUFBWSxLQUFLLE1BQU0sNEdBQThHLENBQUM7QUFBRSxJQUFPLDhCQUFROzs7QURBckw7QUFZQSw0QkFBQyxjQUFjLHFCQUFxQjtBQUM3QixJQUFNLGVBQU4sZUFBMkIsaUJBQVc7QUFBQSxFQUMzQyxPQUFPLFNBQVM7QUFBQSxFQUVoQixvQkFBb0M7QUFBQSxFQUVwQyxvQkFBb0I7QUFDbEIsVUFBTSxrQkFBa0I7QUFDeEIsU0FBSyxpQkFBaUIsVUFBVSxLQUFLLFNBQTBCO0FBQy9ELFNBQUssaUJBQWlCLFdBQVcsS0FBSyxVQUEyQjtBQUFBLEVBQ25FO0FBQUEsRUFFQSx1QkFBdUI7QUFDckIsVUFBTSxxQkFBcUI7QUFDM0IsU0FBSyxvQkFB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
