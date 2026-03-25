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

// elements/pf-v6-tree-view/pf-v6-tree-view.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:/home/bennyp/Developer/cem/serve/elements/pf-v6-tree-view/pf-v6-tree-view.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\n#tree {\\n  margin: 0;\\n  padding: 0;\\n  list-style: none;\\n}\\n"'));
var pf_v6_tree_view_default = s;

// elements/pf-v6-tree-view/pf-v6-tree-view.ts
var _PfV6TreeView_decorators, _init, _a;
_PfV6TreeView_decorators = [customElement("pf-v6-tree-view")];
var PfV6TreeView = class extends (_a = LitElement) {
  static styles = pf_v6_tree_view_default;
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
    if (target.tagName !== "PF-V6-TREE-ITEM") return;
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
      const children = Array.from(parent.children).filter((el) => el.tagName === "PF-V6-TREE-ITEM");
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
    return Array.from(item.children).filter((el) => el.tagName === "PF-V6-TREE-ITEM");
  }
  #getTopLevelItems() {
    return Array.from(this.children).filter((el) => el.tagName === "PF-V6-TREE-ITEM");
  }
  #getParentItem(item) {
    let current = item.parentElement;
    while (current && current !== this) {
      if (current.tagName === "PF-V6-TREE-ITEM") {
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
    const allItems = this.querySelectorAll("pf-v6-tree-item");
    allItems.forEach((item) => {
      if (item !== firstItem) {
        item.setTabindex?.(-1);
      }
    });
  }
  #focusItem(item) {
    if (!item) return;
    const allItems = this.querySelectorAll("pf-v6-tree-item");
    allItems.forEach((i) => i.setTabindex?.(-1));
    item.setTabindex?.(0);
    item.focusItem?.();
  }
  expandAll() {
    this.querySelectorAll("pf-v6-tree-item").forEach(
      (item) => item.expand?.()
    );
  }
  collapseAll() {
    this.querySelectorAll("pf-v6-tree-item").forEach(
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvcGYtdjYtdHJlZS12aWV3L3BmLXY2LXRyZWUtdmlldy50cyIsICJsaXQtY3NzOi9ob21lL2Jlbm55cC9EZXZlbG9wZXIvY2VtL3NlcnZlL2VsZW1lbnRzL3BmLXY2LXRyZWUtdmlldy9wZi12Ni10cmVlLXZpZXcuY3NzIl0sCiAgInNvdXJjZXNDb250ZW50IjogWyJpbXBvcnQgeyBMaXRFbGVtZW50LCBodG1sIH0gZnJvbSAnbGl0JztcbmltcG9ydCB7IGN1c3RvbUVsZW1lbnQgfSBmcm9tICdsaXQvZGVjb3JhdG9ycy9jdXN0b20tZWxlbWVudC5qcyc7XG5cbmltcG9ydCBzdHlsZXMgZnJvbSAnLi9wZi12Ni10cmVlLXZpZXcuY3NzJztcblxuLyoqXG4gKiBQYXR0ZXJuRmx5IHY2IFRyZWUgVmlld1xuICpcbiAqIENvbnRhaW5lciBmb3IgcGYtdjYtdHJlZS1pdGVtIGVsZW1lbnRzLiBDb29yZGluYXRlcyBzZWxlY3Rpb24gYW5kIGtleWJvYXJkIG5hdmlnYXRpb24uXG4gKlxuICogQHNsb3QgLSBUcmVlIGl0ZW1zIChwZi12Ni10cmVlLWl0ZW0gZWxlbWVudHMpXG4gKi9cbkBjdXN0b21FbGVtZW50KCdwZi12Ni10cmVlLXZpZXcnKVxuZXhwb3J0IGNsYXNzIFBmVjZUcmVlVmlldyBleHRlbmRzIExpdEVsZW1lbnQge1xuICBzdGF0aWMgc3R5bGVzID0gc3R5bGVzO1xuXG4gICNjdXJyZW50U2VsZWN0aW9uOiBFbGVtZW50IHwgbnVsbCA9IG51bGw7XG5cbiAgY29ubmVjdGVkQ2FsbGJhY2soKSB7XG4gICAgc3VwZXIuY29ubmVjdGVkQ2FsbGJhY2soKTtcbiAgICB0aGlzLmFkZEV2ZW50TGlzdGVuZXIoJ3NlbGVjdCcsIHRoaXMuI29uU2VsZWN0IGFzIEV2ZW50TGlzdGVuZXIpO1xuICAgIHRoaXMuYWRkRXZlbnRMaXN0ZW5lcigna2V5ZG93bicsIHRoaXMuI29uS2V5ZG93biBhcyBFdmVudExpc3RlbmVyKTtcbiAgfVxuXG4gIGRpc2Nvbm5lY3RlZENhbGxiYWNrKCkge1xuICAgIHN1cGVyLmRpc2Nvbm5lY3RlZENhbGxiYWNrKCk7XG4gICAgdGhpcy5yZW1vdmVFdmVudExpc3RlbmVyKCdzZWxlY3QnLCB0aGlzLiNvblNlbGVjdCBhcyBFdmVudExpc3RlbmVyKTtcbiAgICB0aGlzLnJlbW92ZUV2ZW50TGlzdGVuZXIoJ2tleWRvd24nLCB0aGlzLiNvbktleWRvd24gYXMgRXZlbnRMaXN0ZW5lcik7XG4gIH1cblxuICBmaXJzdFVwZGF0ZWQoKSB7XG4gICAgdGhpcy4jaW5pdGlhbGl6ZVRhYmluZGV4KCk7XG4gIH1cblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8dWwgaWQ9XCJ0cmVlXCJcbiAgICAgICAgICByb2xlPVwidHJlZVwiXG4gICAgICAgICAgcGFydD1cInRyZWVcIlxuICAgICAgICAgIGFyaWEtbGFiZWw9JHt0aGlzLmdldEF0dHJpYnV0ZSgnYXJpYS1sYWJlbCcpID8/ICcnfT5cbiAgICAgICAgPHNsb3Q+PC9zbG90PlxuICAgICAgPC91bD5cbiAgICBgO1xuICB9XG5cbiAgI29uU2VsZWN0ID0gKGV2ZW50OiBFdmVudCkgPT4ge1xuICAgIGNvbnN0IGl0ZW0gPSBldmVudC50YXJnZXQgYXMgVHJlZUl0ZW1MaWtlO1xuICAgIGlmICh0aGlzLiNjdXJyZW50U2VsZWN0aW9uICYmIHRoaXMuI2N1cnJlbnRTZWxlY3Rpb24gIT09IGl0ZW0pIHtcbiAgICAgICh0aGlzLiNjdXJyZW50U2VsZWN0aW9uIGFzIFRyZWVJdGVtTGlrZSkuZGVzZWxlY3Q/LigpO1xuICAgIH1cbiAgICB0aGlzLiNjdXJyZW50U2VsZWN0aW9uID0gaXRlbTtcbiAgfTtcblxuICAjb25LZXlkb3duID0gKGV2ZW50OiBLZXlib2FyZEV2ZW50KSA9PiB7XG4gICAgY29uc3QgdGFyZ2V0ID0gZXZlbnQudGFyZ2V0IGFzIEVsZW1lbnQ7XG4gICAgaWYgKHRhcmdldC50YWdOYW1lICE9PSAnUEYtVjYtVFJFRS1JVEVNJykgcmV0dXJuO1xuXG4gICAgY29uc3QgdHJlZUl0ZW0gPSB0YXJnZXQgYXMgVHJlZUl0ZW1MaWtlO1xuICAgIGNvbnN0IGl0ZW1zID0gdGhpcy4jZ2V0QWxsVmlzaWJsZUl0ZW1zKCk7XG4gICAgY29uc3QgY3VycmVudEluZGV4ID0gaXRlbXMuaW5kZXhPZih0YXJnZXQpO1xuXG4gICAgc3dpdGNoIChldmVudC5rZXkpIHtcbiAgICAgIGNhc2UgJ0Fycm93RG93bic6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmIChjdXJyZW50SW5kZXggPCBpdGVtcy5sZW5ndGggLSAxKSB7XG4gICAgICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zW2N1cnJlbnRJbmRleCArIDFdKTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnQXJyb3dVcCc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmIChjdXJyZW50SW5kZXggPiAwKSB7XG4gICAgICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zW2N1cnJlbnRJbmRleCAtIDFdKTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnQXJyb3dSaWdodCc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmICh0cmVlSXRlbS5oYXNDaGlsZHJlbikge1xuICAgICAgICAgIGlmICghdHJlZUl0ZW0uZXhwYW5kZWQpIHtcbiAgICAgICAgICAgIHRyZWVJdGVtLmV4cGFuZD8uKCk7XG4gICAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgIGNvbnN0IGNoaWxkcmVuID0gdGhpcy4jZ2V0RGlyZWN0Q2hpbGRyZW4odGFyZ2V0KTtcbiAgICAgICAgICAgIGlmIChjaGlsZHJlbi5sZW5ndGggPiAwKSB7XG4gICAgICAgICAgICAgIHRoaXMuI2ZvY3VzSXRlbShjaGlsZHJlblswXSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICAgIGJyZWFrO1xuXG4gICAgICBjYXNlICdBcnJvd0xlZnQnOlxuICAgICAgICBldmVudC5wcmV2ZW50RGVmYXVsdCgpO1xuICAgICAgICBpZiAodHJlZUl0ZW0uaGFzQ2hpbGRyZW4gJiYgdHJlZUl0ZW0uZXhwYW5kZWQpIHtcbiAgICAgICAgICB0cmVlSXRlbS5jb2xsYXBzZT8uKCk7XG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgY29uc3QgcGFyZW50ID0gdGhpcy4jZ2V0UGFyZW50SXRlbSh0YXJnZXQpO1xuICAgICAgICAgIGlmIChwYXJlbnQpIHtcbiAgICAgICAgICAgIHRoaXMuI2ZvY3VzSXRlbShwYXJlbnQpO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnSG9tZSc6XG4gICAgICAgIGV2ZW50LnByZXZlbnREZWZhdWx0KCk7XG4gICAgICAgIGlmIChpdGVtcy5sZW5ndGggPiAwKSB7XG4gICAgICAgICAgdGhpcy4jZm9jdXNJdGVtKGl0ZW1zWzBdKTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnRW5kJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgaWYgKGl0ZW1zLmxlbmd0aCA+IDApIHtcbiAgICAgICAgICB0aGlzLiNmb2N1c0l0ZW0oaXRlbXNbaXRlbXMubGVuZ3RoIC0gMV0pO1xuICAgICAgICB9XG4gICAgICAgIGJyZWFrO1xuXG4gICAgICBjYXNlICdFbnRlcic6XG4gICAgICBjYXNlICcgJzpcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgdHJlZUl0ZW0uc2VsZWN0Py4oKTtcbiAgICAgICAgaWYgKHRyZWVJdGVtLmhhc0NoaWxkcmVuKSB7XG4gICAgICAgICAgdHJlZUl0ZW0udG9nZ2xlPy4oKTtcbiAgICAgICAgfVxuICAgICAgICBicmVhaztcblxuICAgICAgY2FzZSAnKic6IHtcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICAgICAgY29uc3QgcGFyZW50ID0gdGhpcy4jZ2V0UGFyZW50SXRlbSh0YXJnZXQpO1xuICAgICAgICBjb25zdCBzaWJsaW5ncyA9IHBhcmVudCA/IHRoaXMuI2dldERpcmVjdENoaWxkcmVuKHBhcmVudCkgOiB0aGlzLiNnZXRUb3BMZXZlbEl0ZW1zKCk7XG4gICAgICAgIHNpYmxpbmdzLmZvckVhY2goaXRlbSA9PiB7XG4gICAgICAgICAgaWYgKChpdGVtIGFzIFRyZWVJdGVtTGlrZSkuaGFzQ2hpbGRyZW4pIHtcbiAgICAgICAgICAgIChpdGVtIGFzIFRyZWVJdGVtTGlrZSkuZXhwYW5kPy4oKTtcbiAgICAgICAgICB9XG4gICAgICAgIH0pO1xuICAgICAgICBicmVhaztcbiAgICAgIH1cbiAgICB9XG4gIH07XG5cbiAgI2dldEFsbFZpc2libGVJdGVtcygpOiBFbGVtZW50W10ge1xuICAgIGNvbnN0IHZpc2libGU6IEVsZW1lbnRbXSA9IFtdO1xuICAgIGNvbnN0IHdhbGsgPSAocGFyZW50OiBFbGVtZW50KSA9PiB7XG4gICAgICBjb25zdCBjaGlsZHJlbiA9IEFycmF5LmZyb20ocGFyZW50LmNoaWxkcmVuKS5maWx0ZXIoZWwgPT4gZWwudGFnTmFtZSA9PT0gJ1BGLVY2LVRSRUUtSVRFTScpO1xuICAgICAgZm9yIChjb25zdCBpdGVtIG9mIGNoaWxkcmVuKSB7XG4gICAgICAgIHZpc2libGUucHVzaChpdGVtKTtcbiAgICAgICAgY29uc3QgdHJlZUl0ZW0gPSBpdGVtIGFzIFRyZWVJdGVtTGlrZTtcbiAgICAgICAgaWYgKHRyZWVJdGVtLmV4cGFuZGVkICYmIHRyZWVJdGVtLmhhc0NoaWxkcmVuKSB7XG4gICAgICAgICAgd2FsayhpdGVtKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH07XG4gICAgd2Fsayh0aGlzKTtcbiAgICByZXR1cm4gdmlzaWJsZTtcbiAgfVxuXG4gICNnZXREaXJlY3RDaGlsZHJlbihpdGVtOiBFbGVtZW50KTogRWxlbWVudFtdIHtcbiAgICByZXR1cm4gQXJyYXkuZnJvbShpdGVtLmNoaWxkcmVuKS5maWx0ZXIoZWwgPT4gZWwudGFnTmFtZSA9PT0gJ1BGLVY2LVRSRUUtSVRFTScpO1xuICB9XG5cbiAgI2dldFRvcExldmVsSXRlbXMoKTogRWxlbWVudFtdIHtcbiAgICByZXR1cm4gQXJyYXkuZnJvbSh0aGlzLmNoaWxkcmVuKS5maWx0ZXIoZWwgPT4gZWwudGFnTmFtZSA9PT0gJ1BGLVY2LVRSRUUtSVRFTScpO1xuICB9XG5cbiAgI2dldFBhcmVudEl0ZW0oaXRlbTogRWxlbWVudCk6IEVsZW1lbnQgfCBudWxsIHtcbiAgICBsZXQgY3VycmVudCA9IGl0ZW0ucGFyZW50RWxlbWVudDtcbiAgICB3aGlsZSAoY3VycmVudCAmJiBjdXJyZW50ICE9PSB0aGlzKSB7XG4gICAgICBpZiAoY3VycmVudC50YWdOYW1lID09PSAnUEYtVjYtVFJFRS1JVEVNJykge1xuICAgICAgICByZXR1cm4gY3VycmVudDtcbiAgICAgIH1cbiAgICAgIGN1cnJlbnQgPSBjdXJyZW50LnBhcmVudEVsZW1lbnQ7XG4gICAgfVxuICAgIHJldHVybiBudWxsO1xuICB9XG5cbiAgI2luaXRpYWxpemVUYWJpbmRleCgpIHtcbiAgICBjb25zdCB0b3BMZXZlbEl0ZW1zID0gdGhpcy4jZ2V0VG9wTGV2ZWxJdGVtcygpO1xuICAgIGlmICh0b3BMZXZlbEl0ZW1zLmxlbmd0aCA9PT0gMCkgcmV0dXJuO1xuICAgIGNvbnN0IGZpcnN0SXRlbSA9IHRvcExldmVsSXRlbXNbMF0gYXMgVHJlZUl0ZW1MaWtlO1xuICAgIGZpcnN0SXRlbS5zZXRUYWJpbmRleD8uKDApO1xuICAgIGNvbnN0IGFsbEl0ZW1zID0gdGhpcy5xdWVyeVNlbGVjdG9yQWxsKCdwZi12Ni10cmVlLWl0ZW0nKTtcbiAgICBhbGxJdGVtcy5mb3JFYWNoKGl0ZW0gPT4ge1xuICAgICAgaWYgKGl0ZW0gIT09IGZpcnN0SXRlbSkge1xuICAgICAgICAoaXRlbSBhcyB1bmtub3duIGFzIFRyZWVJdGVtTGlrZSkuc2V0VGFiaW5kZXg/LigtMSk7XG4gICAgICB9XG4gICAgfSk7XG4gIH1cblxuICAjZm9jdXNJdGVtKGl0ZW06IEVsZW1lbnQpIHtcbiAgICBpZiAoIWl0ZW0pIHJldHVybjtcbiAgICBjb25zdCBhbGxJdGVtcyA9IHRoaXMucXVlcnlTZWxlY3RvckFsbCgncGYtdjYtdHJlZS1pdGVtJyk7XG4gICAgYWxsSXRlbXMuZm9yRWFjaChpID0+IChpIGFzIHVua25vd24gYXMgVHJlZUl0ZW1MaWtlKS5zZXRUYWJpbmRleD8uKC0xKSk7XG4gICAgKGl0ZW0gYXMgVHJlZUl0ZW1MaWtlKS5zZXRUYWJpbmRleD8uKDApO1xuICAgIChpdGVtIGFzIFRyZWVJdGVtTGlrZSkuZm9jdXNJdGVtPy4oKTtcbiAgfVxuXG4gIGV4cGFuZEFsbCgpIHtcbiAgICB0aGlzLnF1ZXJ5U2VsZWN0b3JBbGwoJ3BmLXY2LXRyZWUtaXRlbScpLmZvckVhY2goaXRlbSA9PlxuICAgICAgKGl0ZW0gYXMgdW5rbm93biBhcyBUcmVlSXRlbUxpa2UpLmV4cGFuZD8uKClcbiAgICApO1xuICB9XG5cbiAgY29sbGFwc2VBbGwoKSB7XG4gICAgdGhpcy5xdWVyeVNlbGVjdG9yQWxsKCdwZi12Ni10cmVlLWl0ZW0nKS5mb3JFYWNoKGl0ZW0gPT5cbiAgICAgIChpdGVtIGFzIHVua25vd24gYXMgVHJlZUl0ZW1MaWtlKS5jb2xsYXBzZT8uKClcbiAgICApO1xuICB9XG59XG5cbi8qKiBJbnRlcmZhY2UgZm9yIHBmLXY2LXRyZWUtaXRlbSBtZXRob2RzIHVzZWQgYnkgdHJlZS12aWV3ICovXG5pbnRlcmZhY2UgVHJlZUl0ZW1MaWtlIGV4dGVuZHMgRWxlbWVudCB7XG4gIGhhc0NoaWxkcmVuPzogYm9vbGVhbjtcbiAgZXhwYW5kZWQ/OiBib29sZWFuO1xuICBleHBhbmQ/KCk6IHZvaWQ7XG4gIGNvbGxhcHNlPygpOiB2b2lkO1xuICB0b2dnbGU/KCk6IHZvaWQ7XG4gIHNlbGVjdD8oKTogdm9pZDtcbiAgZGVzZWxlY3Q/KCk6IHZvaWQ7XG4gIHNldFRhYmluZGV4Pyh2YWx1ZTogbnVtYmVyKTogdm9pZDtcbiAgZm9jdXNJdGVtPygpOiB2b2lkO1xufVxuXG5kZWNsYXJlIGdsb2JhbCB7XG4gIGludGVyZmFjZSBIVE1MRWxlbWVudFRhZ05hbWVNYXAge1xuICAgICdwZi12Ni10cmVlLXZpZXcnOiBQZlY2VHJlZVZpZXc7XG4gIH1cbn1cbiIsICJjb25zdCBzPW5ldyBDU1NTdHlsZVNoZWV0KCk7cy5yZXBsYWNlU3luYyhKU09OLnBhcnNlKFwiXFxcIjpob3N0IHtcXFxcbiAgZGlzcGxheTogYmxvY2s7XFxcXG59XFxcXG5cXFxcbiN0cmVlIHtcXFxcbiAgbWFyZ2luOiAwO1xcXFxuICBwYWRkaW5nOiAwO1xcXFxuICBsaXN0LXN0eWxlOiBub25lO1xcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7OztBQ0Q5QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSw0R0FBOEcsQ0FBQztBQUFFLElBQU8sMEJBQVE7OztBREFyTDtBQVlBLDRCQUFDLGNBQWMsaUJBQWlCO0FBQ3pCLElBQU0sZUFBTixlQUEyQixpQkFBVztBQUFBLEVBQzNDLE9BQU8sU0FBUztBQUFBLEVBRWhCLG9CQUFvQztBQUFBLEVBRXBDLG9CQUFvQjtBQUNsQixVQUFNLGtCQUFrQjtBQUN4QixTQUFLLGlCQUFpQixVQUFVLEtBQUssU0FBMEI7QUFDL0QsU0FBSyxpQkFBaUIsV0FBVyxLQUFLLFVBQTJCO0FBQUEsRUFDbkU7QUFBQSxFQUVBLHVCQUF1QjtBQUNyQixVQUFNLHFCQUFxQjtBQUMzQixTQUFLLG9CQUFvQixVQUFVLEtBQUssU0FBMEI7QUFDbEUsU0FBSyxvQkFBb0IsV0FBVyxL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
