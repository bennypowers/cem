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

// elements/cem-virtual-tree/cem-virtual-tree.ts
import { LitElement, html } from "/__cem/vendor/lit.js";
import { customElement } from "/__cem/vendor/lit/decorators/custom-element.js";

// lit-css:elements/cem-virtual-tree/cem-virtual-tree.css
var s = new CSSStyleSheet();
s.replaceSync(JSON.parse('"/* Virtual Tree - Efficiently renders large tree structures */\\n\\n:host {\\n  display: block;\\n  height: 100%;\\n  overflow: hidden;\\n}\\n\\n#tree-container {\\n  position: relative;\\n  width: 100%;\\n}\\n\\n#viewport {\\n  height: 100%;\\n  overflow-y: auto;\\n  overflow-x: hidden;\\n  position: relative;\\n\\n  cem-pf-v6-tree-item {\\n    cursor: pointer;\\n\\n    \\u0026:hover {\\n      background-color: var(--pf-t--global--background--color--secondary--hover, #f5f5f5);\\n    }\\n\\n    /* Only outline the actually focused item, not ancestors */\\n    \\u0026:focus-visible {\\n      outline: 2px solid var(--pf-t--global--color--brand--default, #0066cc);\\n      outline-offset: -2px;\\n    }\\n  }\\n}\\n"'));
var cem_virtual_tree_default = s;

// elements/cem-virtual-tree/cem-virtual-tree.ts
import "../cem-pf-v6-tree-item/cem-pf-v6-tree-item.js";
var ItemSelectEvent = class extends Event {
  item;
  constructor(item) {
    super("item-select", { bubbles: true, composed: true });
    this.item = item;
  }
};
var _CemVirtualTree_decorators, _init, _a;
_CemVirtualTree_decorators = [customElement("cem-virtual-tree")];
var _CemVirtualTree = class _CemVirtualTree extends (_a = LitElement) {
  static styles = cem_virtual_tree_default;
  // Static cache for manifest (shared across all instances)
  static #manifestCache = null;
  static #manifestPromise = null;
  #manifest = null;
  #flatItems = [];
  #visibleItems = [];
  #searchQuery = "";
  #viewport = null;
  #currentSelectedElement = null;
  #currentSelectedItemId = null;
  render() {
    return html`
      <div id="tree-container">
        <div id="viewport" role="tree"></div>
      </div>
    `;
  }
  async firstUpdated() {
    this.#viewport = this.shadowRoot.getElementById("viewport");
    this.#manifest = await this.#loadManifest();
    if (!this.#manifest) {
      console.warn("[virtual-tree] Failed to load manifest");
      return;
    }
    this.#buildFlatList();
    this.#renderTree();
  }
  /**
   * Public static method to load manifest with caching
   * Can be called by other components to reuse the same cached manifest
   */
  static async loadManifest() {
    if (_CemVirtualTree.#manifestCache) {
      return _CemVirtualTree.#manifestCache;
    }
    if (_CemVirtualTree.#manifestPromise) {
      return _CemVirtualTree.#manifestPromise;
    }
    _CemVirtualTree.#manifestPromise = fetch("/custom-elements.json").then(async (response) => {
      if (!response.ok) {
        throw new Error(`Failed to fetch manifest: ${response.status}`);
      }
      const manifest = await response.json();
      _CemVirtualTree.#manifestCache = manifest;
      return manifest;
    }).catch((error) => {
      console.error("[virtual-tree] Error loading manifest:", error);
      _CemVirtualTree.#manifestPromise = null;
      return null;
    });
    return _CemVirtualTree.#manifestPromise;
  }
  /**
   * Clear the static manifest cache (for testing)
   */
  static clearCache() {
    _CemVirtualTree.#manifestCache = null;
    _CemVirtualTree.#manifestPromise = null;
  }
  /**
   * Load manifest from server with static caching
   */
  async #loadManifest() {
    return _CemVirtualTree.loadManifest();
  }
  /**
   * Build flat list from hierarchical manifest
   */
  #buildFlatList() {
    this.#flatItems = [];
    let id = 0;
    const hasMultiplePackages = this.#manifest.packages && this.#manifest.packages.length > 1;
    if (hasMultiplePackages) {
      for (const pkg of this.#manifest.packages) {
        const packageId = id++;
        const packageItem = {
          id: packageId,
          type: "package",
          label: pkg.name,
          depth: 0,
          hasChildren: (pkg.modules?.length ?? 0) > 0,
          expanded: false,
          visible: true,
          packageName: pkg.name,
          badge: pkg.modules?.length ?? 0
        };
        this.#flatItems.push(packageItem);
        if (!pkg.modules) continue;
        id = this.#buildModulesForPackage(pkg.modules, packageId, 1, id);
      }
    } else {
      const modules = this.#manifest.packages?.[0]?.modules ?? this.#manifest.modules;
      if (!modules) return;
      this.#buildModulesForPackage(modules, void 0, 0, id);
    }
    this.#updateVisibleItems();
  }
  /**
   * Build modules and their declarations for a package
   */
  #buildModulesForPackage(modules, parentId, depth, startId) {
    let id = startId;
    for (const module of modules) {
      const moduleId = id++;
      const moduleItem = {
        id: moduleId,
        type: "module",
        label: module.path,
        depth,
        parentId,
        hasChildren: (module.declarations?.length ?? 0) > 0,
        expanded: false,
        visible: parentId === void 0,
        // Visible if no parent, otherwise hidden until parent expands
        modulePath: module.path,
        badge: module.declarations?.length ?? 0
      };
      this.#flatItems.push(moduleItem);
      if (!module.declarations) continue;
      for (const decl of module.declarations) {
        if (decl.kind === "class" && decl.customElement && decl.tagName) {
          const ceId = id++;
          const properties = decl.members?.filter((m) => m.kind === "field") ?? [];
          const methods = decl.members?.filter((m) => m.kind === "method") ?? [];
          const hasChildren = [
            decl.attributes,
            properties,
            methods,
            decl.events,
            decl.slots,
            decl.cssProperties,
            decl.cssParts,
            decl.cssStates,
            decl.demos
          ].some((x) => (x?.length ?? 0) > 0);
          const ceItem = {
            id: ceId,
            type: "custom-element",
            label: `<${decl.tagName}>`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren,
            expanded: false,
            visible: false,
            modulePath: module.path,
            tagName: decl.tagName
          };
          this.#flatItems.push(ceItem);
          if (decl.attributes?.length) {
            const attrCatId = id++;
            this.#flatItems.push({
              id: attrCatId,
              type: "category",
              label: "Attributes",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.attributes.length,
              category: "attributes"
            });
            for (const attr of decl.attributes) {
              this.#flatItems.push({
                id: id++,
                type: "attribute",
                label: attr.name,
                depth: depth + 3,
                parentId: attrCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: attr.name
              });
            }
          }
          if (properties.length) {
            const propCatId = id++;
            this.#flatItems.push({
              id: propCatId,
              type: "category",
              label: "Properties",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: properties.length,
              category: "properties"
            });
            for (const prop of properties) {
              this.#flatItems.push({
                id: id++,
                type: "property",
                label: prop.name,
                depth: depth + 3,
                parentId: propCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: prop.name
              });
            }
          }
          if (methods.length) {
            const methodCatId = id++;
            this.#flatItems.push({
              id: methodCatId,
              type: "category",
              label: "Methods",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: methods.length,
              category: "methods"
            });
            for (const method of methods) {
              this.#flatItems.push({
                id: id++,
                type: "method",
                label: `${method.name}()`,
                depth: depth + 3,
                parentId: methodCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: method.name
              });
            }
          }
          if (decl.events?.length) {
            const eventCatId = id++;
            this.#flatItems.push({
              id: eventCatId,
              type: "category",
              label: "Events",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.events.length,
              category: "events"
            });
            for (const event of decl.events) {
              this.#flatItems.push({
                id: id++,
                type: "event",
                label: event.name,
                depth: depth + 3,
                parentId: eventCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: event.name
              });
            }
          }
          if (decl.slots?.length) {
            const slotCatId = id++;
            this.#flatItems.push({
              id: slotCatId,
              type: "category",
              label: "Slots",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.slots.length,
              category: "slots"
            });
            for (const slot of decl.slots) {
              this.#flatItems.push({
                id: id++,
                type: "slot",
                label: slot.name || "(default)",
                depth: depth + 3,
                parentId: slotCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: slot.name
              });
            }
          }
          if (decl.cssProperties?.length) {
            const cssPropCatId = id++;
            this.#flatItems.push({
              id: cssPropCatId,
              type: "category",
              label: "CSS Properties",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssProperties.length,
              category: "css-properties"
            });
            for (const cssProp of decl.cssProperties) {
              this.#flatItems.push({
                id: id++,
                type: "css-property",
                label: cssProp.name,
                depth: depth + 3,
                parentId: cssPropCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssProp.name
              });
            }
          }
          if (decl.cssParts?.length) {
            const cssPartCatId = id++;
            this.#flatItems.push({
              id: cssPartCatId,
              type: "category",
              label: "CSS Parts",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssParts.length,
              category: "css-parts"
            });
            for (const cssPart of decl.cssParts) {
              this.#flatItems.push({
                id: id++,
                type: "css-part",
                label: cssPart.name,
                depth: depth + 3,
                parentId: cssPartCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssPart.name
              });
            }
          }
          if (decl.cssStates?.length) {
            const cssStateCatId = id++;
            this.#flatItems.push({
              id: cssStateCatId,
              type: "category",
              label: "CSS States",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssStates.length,
              category: "css-states"
            });
            for (const cssState of decl.cssStates) {
              this.#flatItems.push({
                id: id++,
                type: "css-state",
                label: cssState.name,
                depth: depth + 3,
                parentId: cssStateCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssState.name
              });
            }
          }
          if (decl.demos?.length) {
            const demoCatId = id++;
            this.#flatItems.push({
              id: demoCatId,
              type: "category",
              label: "Demos",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.demos.length,
              category: "demos"
            });
            for (const demo of decl.demos) {
              this.#flatItems.push({
                id: id++,
                type: "demo",
                label: demo.name || demo.url || "(demo)",
                depth: depth + 3,
                parentId: demoCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                url: demo.url
              });
            }
          }
        } else if (decl.kind === "class") {
          this.#flatItems.push({
            id: id++,
            type: "class",
            label: decl.name,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name
          });
        } else if (decl.kind === "function") {
          this.#flatItems.push({
            id: id++,
            type: "function",
            label: `${decl.name}()`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name
          });
        } else if (decl.kind === "variable") {
          this.#flatItems.push({
            id: id++,
            type: "variable",
            label: decl.name,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name
          });
        } else if (decl.kind === "mixin") {
          this.#flatItems.push({
            id: id++,
            type: "mixin",
            label: `${decl.name}()`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name
          });
        }
      }
    }
    return id;
  }
  /**
   * Update visible items based on expanded state
   */
  #updateVisibleItems() {
    this.#visibleItems = this.#flatItems.filter((item) => item.visible);
  }
  /**
   * Render visible tree items with proper nesting
   */
  #renderTree() {
    if (!this.#viewport) return;
    this.#viewport.innerHTML = "";
    const rootItems = this.#visibleItems.filter((item) => item.depth === 0);
    for (const rootItem of rootItems) {
      const treeItemEl = this.#createTreeItemWithChildren(rootItem);
      if (treeItemEl) {
        this.#viewport.appendChild(treeItemEl);
      }
    }
  }
  /**
   * Create a tree item element with its nested children
   */
  #createTreeItemWithChildren(item) {
    const el = document.createElement("cem-pf-v6-tree-item");
    el.setAttribute("label", item.label);
    if (item.badge) {
      el.setAttribute("badge", String(item.badge));
    }
    el.dataset.itemId = String(item.id);
    el.dataset.type = item.type;
    if (this.#currentSelectedItemId === item.id) {
      el.setAttribute("current", "");
      this.#currentSelectedElement = el;
    }
    if (item.hasChildren) {
      el.setAttribute("has-children", "");
      el.expanded = item.expanded;
      el.addEventListener("expand", (e) => {
        e.stopPropagation();
        this.#handleToggle(item, true);
      });
      el.addEventListener("collapse", (e) => {
        e.stopPropagation();
        this.#handleToggle(item, false);
      });
    }
    if (item.type !== "category") {
      el.addEventListener("select", (e) => {
        e.stopPropagation();
        this.#handleSelect(item, el);
      });
    }
    if (item.expanded && item.hasChildren) {
      const children = this.#visibleItems.filter((child) => child.parentId === item.id && child.visible);
      for (const child of children) {
        const childEl = this.#createTreeItemWithChildren(child);
        if (childEl) {
          el.appendChild(childEl);
        }
      }
    }
    return el;
  }
  /**
   * Handle tree item toggle (expand/collapse)
   */
  #handleToggle(item, expanded) {
    item.expanded = expanded;
    this.#updateChildVisibility(item);
    this.#updateVisibleItems();
    this.#renderTree();
  }
  /**
   * Update visibility of children when parent is toggled
   */
  #updateChildVisibility(parentItem) {
    const isExpanded = parentItem.expanded;
    for (const item of this.#flatItems) {
      if (item.parentId === parentItem.id) {
        item.visible = isExpanded;
        if (!isExpanded) {
          item.expanded = false;
          this.#updateChildVisibility(item);
        }
      }
    }
  }
  /**
   * Handle tree item selection
   */
  #handleSelect(item, element) {
    if (this.#currentSelectedElement) {
      this.#currentSelectedElement.removeAttribute("current");
    }
    element.setAttribute("current", "");
    this.#currentSelectedElement = element;
    this.#currentSelectedItemId = item.id;
    this.dispatchEvent(new ItemSelectEvent(item));
  }
  /**
   * Search/filter tree items
   */
  search(query) {
    this.#searchQuery = query.toLowerCase().trim();
    if (!this.#searchQuery) {
      for (const item of this.#flatItems) {
        item.visible = item.depth === 0 || item.parentId !== void 0 && this.#isParentExpanded(item);
      }
    } else {
      const matchingItems = /* @__PURE__ */ new Set();
      for (const item of this.#flatItems) {
        if (item.label.toLowerCase().includes(this.#searchQuery)) {
          matchingItems.add(item);
          this.#showAncestors(item);
        }
      }
      for (const item of this.#flatItems) {
        if (!matchingItems.has(item) && !this.#hasMatchingDescendant(item, matchingItems)) {
          item.visible = this.#isAncestorOfMatch(item, matchingItems);
        } else {
          item.visible = true;
        }
      }
    }
    this.#updateVisibleItems();
    this.#renderTree();
  }
  /**
   * Check if parent is expanded
   */
  #isParentExpanded(item) {
    if (item.parentId === void 0) return true;
    const parent = this.#flatItems.find((i) => i.id === item.parentId);
    return !!parent && parent.expanded && this.#isParentExpanded(parent);
  }
  /**
   * Show all ancestors of an item
   */
  #showAncestors(item) {
    if (item.parentId === void 0) return;
    const parent = this.#flatItems.find((i) => i.id === item.parentId);
    if (parent) {
      parent.visible = true;
      parent.expanded = true;
      this.#showAncestors(parent);
    }
  }
  /**
   * Check if item is ancestor of a match
   */
  #isAncestorOfMatch(item, matchingItems) {
    for (const match of matchingItems) {
      let current = match;
      while (current.parentId !== void 0) {
        const parent = this.#flatItems.find((i) => i.id === current.parentId);
        if (!parent) break;
        if (parent === item) return true;
        current = parent;
      }
    }
    return false;
  }
  /**
   * Check if item has matching descendant
   */
  #hasMatchingDescendant(item, matchingItems) {
    for (const match of matchingItems) {
      if (this.#isDescendantOf(match, item)) return true;
    }
    return false;
  }
  /**
   * Check if item is descendant of another
   */
  #isDescendantOf(item, ancestor) {
    let current = item;
    while (current.parentId !== void 0) {
      const parent = this.#flatItems.find((i) => i.id === current.parentId);
      if (!parent) return false;
      if (parent === ancestor) return true;
      current = parent;
    }
    return false;
  }
  /**
   * Expand all items
   */
  expandAll() {
    for (const item of this.#flatItems) {
      if (item.hasChildren) {
        item.expanded = true;
      }
      if (!this.#searchQuery || item.label.toLowerCase().includes(this.#searchQuery)) {
        item.visible = true;
      }
    }
    this.#updateVisibleItems();
    this.#renderTree();
  }
  /**
   * Collapse all items
   */
  collapseAll() {
    for (const item of this.#flatItems) {
      if (item.hasChildren) {
        item.expanded = false;
      }
      item.visible = item.depth === 0;
    }
    this.#updateVisibleItems();
    this.#renderTree();
  }
};
_init = __decoratorStart(_a);
_CemVirtualTree = __decorateElement(_init, 0, "CemVirtualTree", _CemVirtualTree_decorators, _CemVirtualTree);
__runInitializers(_init, 1, _CemVirtualTree);
var CemVirtualTree = _CemVirtualTree;
export {
  CemVirtualTree,
  ItemSelectEvent
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXZpcnR1YWwtdHJlZS9jZW0tdmlydHVhbC10cmVlLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXZpcnR1YWwtdHJlZS9jZW0tdmlydHVhbC10cmVlLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXZpcnR1YWwtdHJlZS5jc3MnO1xuXG5pbXBvcnQgJy4uL2NlbS1wZi12Ni10cmVlLWl0ZW0vY2VtLXBmLXY2LXRyZWUtaXRlbS5qcyc7XG5cbmludGVyZmFjZSBUcmVlSXRlbSB7XG4gIGlkOiBudW1iZXI7XG4gIHR5cGU6IHN0cmluZztcbiAgbGFiZWw6IHN0cmluZztcbiAgZGVwdGg6IG51bWJlcjtcbiAgcGFyZW50SWQ/OiBudW1iZXI7XG4gIGhhc0NoaWxkcmVuOiBib29sZWFuO1xuICBleHBhbmRlZDogYm9vbGVhbjtcbiAgdmlzaWJsZTogYm9vbGVhbjtcbiAgcGFja2FnZU5hbWU/OiBzdHJpbmc7XG4gIG1vZHVsZVBhdGg/OiBzdHJpbmc7XG4gIHRhZ05hbWU/OiBzdHJpbmc7XG4gIG5hbWU/OiBzdHJpbmc7XG4gIHVybD86IHN0cmluZztcbiAgYmFkZ2U/OiBudW1iZXI7XG4gIGNhdGVnb3J5Pzogc3RyaW5nO1xufVxuXG5pbnRlcmZhY2UgTWFuaWZlc3Qge1xuICBwYWNrYWdlcz86IE1hbmlmZXN0UGFja2FnZVtdO1xuICBtb2R1bGVzPzogTWFuaWZlc3RNb2R1bGVbXTtcbn1cblxuaW50ZXJmYWNlIE1hbmlmZXN0UGFja2FnZSB7XG4gIG5hbWU6IHN0cmluZztcbiAgbW9kdWxlcz86IE1hbmlmZXN0TW9kdWxlW107XG59XG5cbmludGVyZmFjZSBNYW5pZmVzdE1vZHVsZSB7XG4gIHBhdGg6IHN0cmluZztcbiAgZGVjbGFyYXRpb25zPzogTWFuaWZlc3REZWNsYXJhdGlvbltdO1xufVxuXG5pbnRlcmZhY2UgTWFuaWZlc3REZWNsYXJhdGlvbiB7XG4gIGtpbmQ6IHN0cmluZztcbiAgbmFtZTogc3RyaW5nO1xuICBjdXN0b21FbGVtZW50PzogYm9vbGVhbjtcbiAgdGFnTmFtZT86IHN0cmluZztcbiAgbWVtYmVycz86IE1hbmlmZXN0TWVtYmVyW107XG4gIGF0dHJpYnV0ZXM/OiB7IG5hbWU6IHN0cmluZyB9W107XG4gIGV2ZW50cz86IHsgbmFtZTogc3RyaW5nIH1bXTtcbiAgc2xvdHM/OiB7IG5hbWU6IHN0cmluZzsgZGVzY3JpcHRpb24/OiBzdHJpbmcgfVtdO1xuICBjc3NQcm9wZXJ0aWVzPzogeyBuYW1lOiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH1bXTtcbiAgY3NzUGFydHM/OiB7IG5hbWU6IHN0cmluZzsgZGVzY3JpcHRpb24/OiBzdHJpbmcgfVtdO1xuICBjc3NTdGF0ZXM/OiB7IG5hbWU6IHN0cmluZzsgZGVzY3JpcHRpb24/OiBzdHJpbmcgfVtdO1xuICBkZW1vcz86IHsgbmFtZT86IHN0cmluZzsgdXJsPzogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9W107XG59XG5cbmludGVyZmFjZSBNYW5pZmVzdE1lbWJlciB7XG4gIGtpbmQ6IHN0cmluZztcbiAgbmFtZTogc3RyaW5nO1xuICB0eXBlPzogeyB0ZXh0OiBzdHJpbmcgfTtcbn1cblxuLyoqXG4gKiBDdXN0b20gZXZlbnQgZm9yIGl0ZW0gc2VsZWN0aW9uXG4gKi9cbmV4cG9ydCBjbGFzcyBJdGVtU2VsZWN0RXZlbnQgZXh0ZW5kcyBFdmVudCB7XG4gIGl0ZW06IFRyZWVJdGVtO1xuICBjb25zdHJ1Y3RvcihpdGVtOiBUcmVlSXRlbSkge1xuICAgIHN1cGVyKCdpdGVtLXNlbGVjdCcsIHsgYnViYmxlczogdHJ1ZSwgY29tcG9zZWQ6IHRydWUgfSk7XG4gICAgdGhpcy5pdGVtID0gaXRlbTtcbiAgfVxufVxuXG4vKipcbiAqIFZpcnR1YWwgdHJlZSB2aWV3IHdpdGggbGF6eSBsb2FkaW5nXG4gKlxuICogQGZpcmVzIGl0ZW0tc2VsZWN0IC0gRmlyZWQgd2hlbiBhIHRyZWUgaXRlbSBpcyBzZWxlY3RlZFxuICovXG5AY3VzdG9tRWxlbWVudCgnY2VtLXZpcnR1YWwtdHJlZScpXG5leHBvcnQgY2xhc3MgQ2VtVmlydHVhbFRyZWUgZXh0ZW5kcyBMaXRFbGVtZW50IHtcbiAgc3RhdGljIHN0eWxlcyA9IHN0eWxlcztcblxuICAvLyBTdGF0aWMgY2FjaGUgZm9yIG1hbmlmZXN0IChzaGFyZWQgYWNyb3NzIGFsbCBpbnN0YW5jZXMpXG4gIHN0YXRpYyAjbWFuaWZlc3RDYWNoZTogTWFuaWZlc3QgfCBudWxsID0gbnVsbDtcbiAgc3RhdGljICNtYW5pZmVzdFByb21pc2U6IFByb21pc2U8TWFuaWZlc3QgfCBudWxsPiB8IG51bGwgPSBudWxsO1xuXG4gICNtYW5pZmVzdDogTWFuaWZlc3QgfCBudWxsID0gbnVsbDtcbiAgI2ZsYXRJdGVtczogVHJlZUl0ZW1bXSA9IFtdO1xuICAjdmlzaWJsZUl0ZW1zOiBUcmVlSXRlbVtdID0gW107XG4gICNzZWFyY2hRdWVyeSA9ICcnO1xuICAjdmlld3BvcnQ6IEhUTUxFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNjdXJyZW50U2VsZWN0ZWRFbGVtZW50OiBFbGVtZW50IHwgbnVsbCA9IG51bGw7XG4gICNjdXJyZW50U2VsZWN0ZWRJdGVtSWQ6IG51bWJlciB8IG51bGwgPSBudWxsO1xuXG4gIHJlbmRlcigpIHtcbiAgICByZXR1cm4gaHRtbGBcbiAgICAgIDxkaXYgaWQ9XCJ0cmVlLWNvbnRhaW5lclwiPlxuICAgICAgICA8ZGl2IGlkPVwidmlld3BvcnRcIiByb2xlPVwidHJlZVwiPjwvZGl2PlxuICAgICAgPC9kaXY+XG4gICAgYDtcbiAgfVxuXG4gIGFzeW5jIGZpcnN0VXBkYXRlZCgpIHtcbiAgICB0aGlzLiN2aWV3cG9ydCA9IHRoaXMuc2hhZG93Um9vdCEuZ2V0RWxlbWVudEJ5SWQoJ3ZpZXdwb3J0Jyk7XG5cbiAgICAvLyBMb2FkIG1hbmlmZXN0IGZyb20gc2VydmVyIHdpdGggY2FjaGluZ1xuICAgIHRoaXMuI21hbmlmZXN0ID0gYXdhaXQgdGhpcy4jbG9hZE1hbmlmZXN0KCk7XG4gICAgaWYgKCF0aGlzLiNtYW5pZmVzdCkge1xuICAgICAgY29uc29sZS53YXJuKCdbdmlydHVhbC10cmVlXSBGYWlsZWQgdG8gbG9hZCBtYW5pZmVzdCcpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cblxuICAgIC8vIEJ1aWxkIGZsYXQgbGlzdCBmcm9tIG1hbmlmZXN0XG4gICAgdGhpcy4jYnVpbGRGbGF0TGlzdCgpO1xuXG4gICAgLy8gSW5pdGlhbCByZW5kZXJcbiAgICB0aGlzLiNyZW5kZXJUcmVlKCk7XG4gIH1cblxuICAvKipcbiAgICogUHVibGljIHN0YXRpYyBtZXRob2QgdG8gbG9hZCBtYW5pZmVzdCB3aXRoIGNhY2hpbmdcbiAgICogQ2FuIGJlIGNhbGxlZCBieSBvdGhlciBjb21wb25lbnRzIHRvIHJldXNlIHRoZSBzYW1lIGNhY2hlZCBtYW5pZmVzdFxuICAgKi9cbiAgc3RhdGljIGFzeW5jIGxvYWRNYW5pZmVzdCgpOiBQcm9taXNlPE1hbmlmZXN0IHwgbnVsbD4ge1xuICAgIC8vIFJldHVybiBjYWNoZWQgbWFuaWZlc3QgaWYgYXZhaWxhYmxlXG4gICAgaWYgKENlbVZpcnR1YWxUcmVlLiNtYW5pZmVzdENhY2hlKSB7XG4gICAgICByZXR1cm4gQ2VtVmlydHVhbFRyZWUuI21hbmlmZXN0Q2FjaGU7XG4gICAgfVxuXG4gICAgLy8gSWYgYWxyZWFkeSBmZXRjaGluZywgd2FpdCBmb3IgZXhpc3RpbmcgcHJvbWlzZVxuICAgIGlmIChDZW1WaXJ0dWFsVHJlZS4jbWFuaWZlc3RQcm9taXNlKSB7XG4gICAgICByZXR1cm4gQ2VtVmlydHVhbFRyZWUuI21hbmlmZXN0UHJvbWlzZTtcbiAgICB9XG5cbiAgICAvLyBGZXRjaCBtYW5pZmVzdFxuICAgIENlbVZpcnR1YWxUcmVlLiNtYW5pZmVzdFByb21pc2UgPSBmZXRjaCgnL2N1c3RvbS1lbGVtZW50cy5qc29uJylcbiAgICAgIC50aGVuKGFzeW5jIChyZXNwb25zZSkgPT4ge1xuICAgICAgICBpZiAoIXJlc3BvbnNlLm9rKSB7XG4gICAgICAgICAgdGhyb3cgbmV3IEVycm9yKGBGYWlsZWQgdG8gZmV0Y2ggbWFuaWZlc3Q6ICR7cmVzcG9uc2Uuc3RhdHVzfWApO1xuICAgICAgICB9XG4gICAgICAgIGNvbnN0IG1hbmlmZXN0ID0gYXdhaXQgcmVzcG9uc2UuanNvbigpO1xuICAgICAgICBDZW1WaXJ0dWFsVHJlZS4jbWFuaWZlc3RDYWNoZSA9IG1hbmlmZXN0O1xuICAgICAgICByZXR1cm4gbWFuaWZlc3Q7XG4gICAgICB9KVxuICAgICAgLmNhdGNoKChlcnJvcikgPT4ge1xuICAgICAgICBjb25zb2xlLmVycm9yKCdbdmlydHVhbC10cmVlXSBFcnJvciBsb2FkaW5nIG1hbmlmZXN0OicsIGVycm9yKTtcbiAgICAgICAgQ2VtVmlydHVhbFRyZWUuI21hbmlmZXN0UHJvbWlzZSA9IG51bGw7XG4gICAgICAgIHJldHVybiBudWxsO1xuICAgICAgfSk7XG5cbiAgICByZXR1cm4gQ2VtVmlydHVhbFRyZWUuI21hbmlmZXN0UHJvbWlzZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBDbGVhciB0aGUgc3RhdGljIG1hbmlmZXN0IGNhY2hlIChmb3IgdGVzdGluZylcbiAgICovXG4gIHN0YXRpYyBjbGVhckNhY2hlKCkge1xuICAgIENlbVZpcnR1YWxUcmVlLiNtYW5pZmVzdENhY2hlID0gbnVsbDtcbiAgICBDZW1WaXJ0dWFsVHJlZS4jbWFuaWZlc3RQcm9taXNlID0gbnVsbDtcbiAgfVxuXG4gIC8qKlxuICAgKiBMb2FkIG1hbmlmZXN0IGZyb20gc2VydmVyIHdpdGggc3RhdGljIGNhY2hpbmdcbiAgICovXG4gIGFzeW5jICNsb2FkTWFuaWZlc3QoKTogUHJvbWlzZTxNYW5pZmVzdCB8IG51bGw+IHtcbiAgICByZXR1cm4gQ2VtVmlydHVhbFRyZWUubG9hZE1hbmlmZXN0KCk7XG4gIH1cblxuICAvKipcbiAgICogQnVpbGQgZmxhdCBsaXN0IGZyb20gaGllcmFyY2hpY2FsIG1hbmlmZXN0XG4gICAqL1xuICAjYnVpbGRGbGF0TGlzdCgpIHtcbiAgICB0aGlzLiNmbGF0SXRlbXMgPSBbXTtcbiAgICBsZXQgaWQgPSAwO1xuXG4gICAgLy8gQ2hlY2sgaWYgdGhpcyBpcyBhIHdvcmtzcGFjZSBtYW5pZmVzdCB3aXRoIG11bHRpcGxlIHBhY2thZ2VzXG4gICAgY29uc3QgaGFzTXVsdGlwbGVQYWNrYWdlcyA9IHRoaXMuI21hbmlmZXN0IS5wYWNrYWdlcyAmJiB0aGlzLiNtYW5pZmVzdCEucGFja2FnZXMubGVuZ3RoID4gMTtcblxuICAgIGlmIChoYXNNdWx0aXBsZVBhY2thZ2VzKSB7XG4gICAgICAvLyBXb3Jrc3BhY2UgbW9kZSB3aXRoIG11bHRpcGxlIHBhY2thZ2VzOiBzaG93IHBhY2thZ2VzIGF0IHRvcCBsZXZlbFxuICAgICAgZm9yIChjb25zdCBwa2cgb2YgdGhpcy4jbWFuaWZlc3QhLnBhY2thZ2VzISkge1xuICAgICAgICBjb25zdCBwYWNrYWdlSWQgPSBpZCsrO1xuICAgICAgICBjb25zdCBwYWNrYWdlSXRlbTogVHJlZUl0ZW0gPSB7XG4gICAgICAgICAgaWQ6IHBhY2thZ2VJZCxcbiAgICAgICAgICB0eXBlOiAncGFja2FnZScsXG4gICAgICAgICAgbGFiZWw6IHBrZy5uYW1lLFxuICAgICAgICAgIGRlcHRoOiAwLFxuICAgICAgICAgIGhhc0NoaWxkcmVuOiAocGtnLm1vZHVsZXM/Lmxlbmd0aCA/PyAwKSA+IDAsXG4gICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgIHZpc2libGU6IHRydWUsXG4gICAgICAgICAgcGFja2FnZU5hbWU6IHBrZy5uYW1lLFxuICAgICAgICAgIGJhZGdlOiBwa2cubW9kdWxlcz8ubGVuZ3RoID8/IDAsXG4gICAgICAgIH07XG4gICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHBhY2thZ2VJdGVtKTtcblxuICAgICAgICBpZiAoIXBrZy5tb2R1bGVzKSBjb250aW51ZTtcblxuICAgICAgICAvLyBBZGQgbW9kdWxlcyB1bmRlciB0aGlzIHBhY2thZ2UgYXQgZGVwdGggMVxuICAgICAgICBpZCA9IHRoaXMuI2J1aWxkTW9kdWxlc0ZvclBhY2thZ2UocGtnLm1vZHVsZXMsIHBhY2thZ2VJZCwgMSwgaWQpO1xuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICAvLyBTaW5nbGUgcGFja2FnZSBtb2RlIE9SIHdvcmtzcGFjZSB3aXRoIDEgcGFja2FnZTogc2hvdyBtb2R1bGVzIGF0IHRvcCBsZXZlbFxuICAgICAgY29uc3QgbW9kdWxlcyA9IHRoaXMuI21hbmlmZXN0IS5wYWNrYWdlcz8uWzBdPy5tb2R1bGVzID8/IHRoaXMuI21hbmlmZXN0IS5tb2R1bGVzO1xuXG4gICAgICBpZiAoIW1vZHVsZXMpIHJldHVybjtcblxuICAgICAgLy8gQWRkIG1vZHVsZXMgYXQgZGVwdGggMCAobm8gcGFja2FnZSBsZXZlbClcbiAgICAgIHRoaXMuI2J1aWxkTW9kdWxlc0ZvclBhY2thZ2UobW9kdWxlcywgdW5kZWZpbmVkLCAwLCBpZCk7XG4gICAgfVxuXG4gICAgdGhpcy4jdXBkYXRlVmlzaWJsZUl0ZW1zKCk7XG4gIH1cblxuICAvKipcbiAgICogQnVpbGQgbW9kdWxlcyBhbmQgdGhlaXIgZGVjbGFyYXRpb25zIGZvciBhIHBhY2thZ2VcbiAgICovXG4gICNidWlsZE1vZHVsZXNGb3JQYWNrYWdlKG1vZHVsZXM6IE1hbmlmZXN0TW9kdWxlW10sIHBhcmVudElkOiBudW1iZXIgfCB1bmRlZmluZWQsIGRlcHRoOiBudW1iZXIsIHN0YXJ0SWQ6IG51bWJlcik6IG51bWJlciB7XG4gICAgbGV0IGlkID0gc3RhcnRJZDtcblxuICAgIGZvciAoY29uc3QgbW9kdWxlIG9mIG1vZHVsZXMpIHtcbiAgICAgIGNvbnN0IG1vZHVsZUlkID0gaWQrKztcbiAgICAgIGNvbnN0IG1vZHVsZUl0ZW06IFRyZWVJdGVtID0ge1xuICAgICAgICBpZDogbW9kdWxlSWQsXG4gICAgICAgIHR5cGU6ICdtb2R1bGUnLFxuICAgICAgICBsYWJlbDogbW9kdWxlLnBhdGgsXG4gICAgICAgIGRlcHRoOiBkZXB0aCxcbiAgICAgICAgcGFyZW50SWQ6IHBhcmVudElkLFxuICAgICAgICBoYXNDaGlsZHJlbjogKG1vZHVsZS5kZWNsYXJhdGlvbnM/Lmxlbmd0aCA/PyAwKSA+IDAsXG4gICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgdmlzaWJsZTogcGFyZW50SWQgPT09IHVuZGVmaW5lZCwgLy8gVmlzaWJsZSBpZiBubyBwYXJlbnQsIG90aGVyd2lzZSBoaWRkZW4gdW50aWwgcGFyZW50IGV4cGFuZHNcbiAgICAgICAgbW9kdWxlUGF0aDogbW9kdWxlLnBhdGgsXG4gICAgICAgIGJhZGdlOiBtb2R1bGUuZGVjbGFyYXRpb25zPy5sZW5ndGggPz8gMCxcbiAgICAgIH07XG4gICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaChtb2R1bGVJdGVtKTtcblxuICAgICAgaWYgKCFtb2R1bGUuZGVjbGFyYXRpb25zKSBjb250aW51ZTtcblxuICAgICAgZm9yIChjb25zdCBkZWNsIG9mIG1vZHVsZS5kZWNsYXJhdGlvbnMpIHtcbiAgICAgICAgLy8gQ3VzdG9tIEVsZW1lbnQgLSBtdXN0IGhhdmUgYm90aCBjdXN0b21FbGVtZW50IGZsYWcgYW5kIHRhZ05hbWVcbiAgICAgICAgaWYgKGRlY2wua2luZCA9PT0gJ2NsYXNzJyAmJiBkZWNsLmN1c3RvbUVsZW1lbnQgJiYgZGVjbC50YWdOYW1lKSB7XG4gICAgICAgICAgY29uc3QgY2VJZCA9IGlkKys7XG5cbiAgICAgICAgICAvLyBDb21wdXRlIGhhc0NoaWxkcmVuIGJhc2VkIG9uIHByZXNlbmNlIG9mIGFueSBjaGlsZCBhcnJheXNcbiAgICAgICAgICBjb25zdCBwcm9wZXJ0aWVzID0gZGVjbC5tZW1iZXJzPy5maWx0ZXIobSA9PiBtLmtpbmQgPT09ICdmaWVsZCcpID8/IFtdO1xuICAgICAgICAgIGNvbnN0IG1ldGhvZHMgPSBkZWNsLm1lbWJlcnM/LmZpbHRlcihtID0+IG0ua2luZCA9PT0gJ21ldGhvZCcpID8/IFtdO1xuICAgICAgICAgIGNvbnN0IGhhc0NoaWxkcmVuID0gW1xuICAgICAgICAgICAgZGVjbC5hdHRyaWJ1dGVzLFxuICAgICAgICAgICAgcHJvcGVydGllcyxcbiAgICAgICAgICAgIG1ldGhvZHMsXG4gICAgICAgICAgICBkZWNsLmV2ZW50cyxcbiAgICAgICAgICAgIGRlY2wuc2xvdHMsXG4gICAgICAgICAgICBkZWNsLmNzc1Byb3BlcnRpZXMsXG4gICAgICAgICAgICBkZWNsLmNzc1BhcnRzLFxuICAgICAgICAgICAgZGVjbC5jc3NTdGF0ZXMsXG4gICAgICAgICAgICBkZWNsLmRlbW9zLFxuICAgICAgICAgIF0uc29tZSh4ID0+ICh4Py5sZW5ndGggPz8gMCkgPiAwKTtcblxuICAgICAgICAgIGNvbnN0IGNlSXRlbTogVHJlZUl0ZW0gPSB7XG4gICAgICAgICAgICBpZDogY2VJZCxcbiAgICAgICAgICAgIHR5cGU6ICdjdXN0b20tZWxlbWVudCcsXG4gICAgICAgICAgICBsYWJlbDogYDwke2RlY2wudGFnTmFtZX0+YCxcbiAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDEsXG4gICAgICAgICAgICBwYXJlbnRJZDogbW9kdWxlSWQsXG4gICAgICAgICAgICBoYXNDaGlsZHJlbixcbiAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgbW9kdWxlUGF0aDogbW9kdWxlLnBhdGgsXG4gICAgICAgICAgICB0YWdOYW1lOiBkZWNsLnRhZ05hbWUsXG4gICAgICAgICAgfTtcbiAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaChjZUl0ZW0pO1xuXG4gICAgICAgICAgLy8gQXR0cmlidXRlc1xuICAgICAgICAgIGlmIChkZWNsLmF0dHJpYnV0ZXM/Lmxlbmd0aCkge1xuICAgICAgICAgICAgY29uc3QgYXR0ckNhdElkID0gaWQrKztcbiAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgaWQ6IGF0dHJDYXRJZCxcbiAgICAgICAgICAgICAgdHlwZTogJ2NhdGVnb3J5JyxcbiAgICAgICAgICAgICAgbGFiZWw6ICdBdHRyaWJ1dGVzJyxcbiAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMixcbiAgICAgICAgICAgICAgcGFyZW50SWQ6IGNlSWQsXG4gICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiB0cnVlLFxuICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICBiYWRnZTogZGVjbC5hdHRyaWJ1dGVzLmxlbmd0aCxcbiAgICAgICAgICAgICAgY2F0ZWdvcnk6ICdhdHRyaWJ1dGVzJyxcbiAgICAgICAgICAgIH0pO1xuXG4gICAgICAgICAgICBmb3IgKGNvbnN0IGF0dHIgb2YgZGVjbC5hdHRyaWJ1dGVzKSB7XG4gICAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgICAgICB0eXBlOiAnYXR0cmlidXRlJyxcbiAgICAgICAgICAgICAgICBsYWJlbDogYXR0ci5uYW1lLFxuICAgICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDMsXG4gICAgICAgICAgICAgICAgcGFyZW50SWQ6IGF0dHJDYXRJZCxcbiAgICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgICAgIHRhZ05hbWU6IGRlY2wudGFnTmFtZSxcbiAgICAgICAgICAgICAgICBuYW1lOiBhdHRyLm5hbWUsXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cblxuICAgICAgICAgIC8vIFByb3BlcnRpZXNcbiAgICAgICAgICBpZiAocHJvcGVydGllcy5sZW5ndGgpIHtcbiAgICAgICAgICAgIGNvbnN0IHByb3BDYXRJZCA9IGlkKys7XG4gICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgIGlkOiBwcm9wQ2F0SWQsXG4gICAgICAgICAgICAgIHR5cGU6ICdjYXRlZ29yeScsXG4gICAgICAgICAgICAgIGxhYmVsOiAnUHJvcGVydGllcycsXG4gICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDIsXG4gICAgICAgICAgICAgIHBhcmVudElkOiBjZUlkLFxuICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogdHJ1ZSxcbiAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgYmFkZ2U6IHByb3BlcnRpZXMubGVuZ3RoLFxuICAgICAgICAgICAgICBjYXRlZ29yeTogJ3Byb3BlcnRpZXMnLFxuICAgICAgICAgICAgfSk7XG5cbiAgICAgICAgICAgIGZvciAoY29uc3QgcHJvcCBvZiBwcm9wZXJ0aWVzKSB7XG4gICAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgICAgICB0eXBlOiAncHJvcGVydHknLFxuICAgICAgICAgICAgICAgIGxhYmVsOiBwcm9wLm5hbWUsXG4gICAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMyxcbiAgICAgICAgICAgICAgICBwYXJlbnRJZDogcHJvcENhdElkLFxuICAgICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgICAgbW9kdWxlUGF0aDogbW9kdWxlLnBhdGgsXG4gICAgICAgICAgICAgICAgdGFnTmFtZTogZGVjbC50YWdOYW1lLFxuICAgICAgICAgICAgICAgIG5hbWU6IHByb3AubmFtZSxcbiAgICAgICAgICAgICAgfSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLy8gTWV0aG9kc1xuICAgICAgICAgIGlmIChtZXRob2RzLmxlbmd0aCkge1xuICAgICAgICAgICAgY29uc3QgbWV0aG9kQ2F0SWQgPSBpZCsrO1xuICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICBpZDogbWV0aG9kQ2F0SWQsXG4gICAgICAgICAgICAgIHR5cGU6ICdjYXRlZ29yeScsXG4gICAgICAgICAgICAgIGxhYmVsOiAnTWV0aG9kcycsXG4gICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDIsXG4gICAgICAgICAgICAgIHBhcmVudElkOiBjZUlkLFxuICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogdHJ1ZSxcbiAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgYmFkZ2U6IG1ldGhvZHMubGVuZ3RoLFxuICAgICAgICAgICAgICBjYXRlZ29yeTogJ21ldGhvZHMnLFxuICAgICAgICAgICAgfSk7XG5cbiAgICAgICAgICAgIGZvciAoY29uc3QgbWV0aG9kIG9mIG1ldGhvZHMpIHtcbiAgICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICAgIGlkOiBpZCsrLFxuICAgICAgICAgICAgICAgIHR5cGU6ICdtZXRob2QnLFxuICAgICAgICAgICAgICAgIGxhYmVsOiBgJHttZXRob2QubmFtZX0oKWAsXG4gICAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMyxcbiAgICAgICAgICAgICAgICBwYXJlbnRJZDogbWV0aG9kQ2F0SWQsXG4gICAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgICAgICB0YWdOYW1lOiBkZWNsLnRhZ05hbWUsXG4gICAgICAgICAgICAgICAgbmFtZTogbWV0aG9kLm5hbWUsXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cblxuICAgICAgICAgIC8vIEV2ZW50c1xuICAgICAgICAgIGlmIChkZWNsLmV2ZW50cz8ubGVuZ3RoKSB7XG4gICAgICAgICAgICBjb25zdCBldmVudENhdElkID0gaWQrKztcbiAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgaWQ6IGV2ZW50Q2F0SWQsXG4gICAgICAgICAgICAgIHR5cGU6ICdjYXRlZ29yeScsXG4gICAgICAgICAgICAgIGxhYmVsOiAnRXZlbnRzJyxcbiAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMixcbiAgICAgICAgICAgICAgcGFyZW50SWQ6IGNlSWQsXG4gICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiB0cnVlLFxuICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICBiYWRnZTogZGVjbC5ldmVudHMubGVuZ3RoLFxuICAgICAgICAgICAgICBjYXRlZ29yeTogJ2V2ZW50cycsXG4gICAgICAgICAgICB9KTtcblxuICAgICAgICAgICAgZm9yIChjb25zdCBldmVudCBvZiBkZWNsLmV2ZW50cykge1xuICAgICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICAgICAgdHlwZTogJ2V2ZW50JyxcbiAgICAgICAgICAgICAgICBsYWJlbDogZXZlbnQubmFtZSxcbiAgICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAzLFxuICAgICAgICAgICAgICAgIHBhcmVudElkOiBldmVudENhdElkLFxuICAgICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgICAgbW9kdWxlUGF0aDogbW9kdWxlLnBhdGgsXG4gICAgICAgICAgICAgICAgdGFnTmFtZTogZGVjbC50YWdOYW1lLFxuICAgICAgICAgICAgICAgIG5hbWU6IGV2ZW50Lm5hbWUsXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cblxuICAgICAgICAgIC8vIFNsb3RzXG4gICAgICAgICAgaWYgKGRlY2wuc2xvdHM/Lmxlbmd0aCkge1xuICAgICAgICAgICAgY29uc3Qgc2xvdENhdElkID0gaWQrKztcbiAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgaWQ6IHNsb3RDYXRJZCxcbiAgICAgICAgICAgICAgdHlwZTogJ2NhdGVnb3J5JyxcbiAgICAgICAgICAgICAgbGFiZWw6ICdTbG90cycsXG4gICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDIsXG4gICAgICAgICAgICAgIHBhcmVudElkOiBjZUlkLFxuICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogdHJ1ZSxcbiAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgYmFkZ2U6IGRlY2wuc2xvdHMubGVuZ3RoLFxuICAgICAgICAgICAgICBjYXRlZ29yeTogJ3Nsb3RzJyxcbiAgICAgICAgICAgIH0pO1xuXG4gICAgICAgICAgICBmb3IgKGNvbnN0IHNsb3Qgb2YgZGVjbC5zbG90cykge1xuICAgICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICAgICAgdHlwZTogJ3Nsb3QnLFxuICAgICAgICAgICAgICAgIGxhYmVsOiBzbG90Lm5hbWUgfHwgJyhkZWZhdWx0KScsXG4gICAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMyxcbiAgICAgICAgICAgICAgICBwYXJlbnRJZDogc2xvdENhdElkLFxuICAgICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgICAgbW9kdWxlUGF0aDogbW9kdWxlLnBhdGgsXG4gICAgICAgICAgICAgICAgdGFnTmFtZTogZGVjbC50YWdOYW1lLFxuICAgICAgICAgICAgICAgIG5hbWU6IHNsb3QubmFtZSxcbiAgICAgICAgICAgICAgfSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLy8gQ1NTIFByb3BlcnRpZXNcbiAgICAgICAgICBpZiAoZGVjbC5jc3NQcm9wZXJ0aWVzPy5sZW5ndGgpIHtcbiAgICAgICAgICAgIGNvbnN0IGNzc1Byb3BDYXRJZCA9IGlkKys7XG4gICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgIGlkOiBjc3NQcm9wQ2F0SWQsXG4gICAgICAgICAgICAgIHR5cGU6ICdjYXRlZ29yeScsXG4gICAgICAgICAgICAgIGxhYmVsOiAnQ1NTIFByb3BlcnRpZXMnLFxuICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAyLFxuICAgICAgICAgICAgICBwYXJlbnRJZDogY2VJZCxcbiAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IHRydWUsXG4gICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgIGJhZGdlOiBkZWNsLmNzc1Byb3BlcnRpZXMubGVuZ3RoLFxuICAgICAgICAgICAgICBjYXRlZ29yeTogJ2Nzcy1wcm9wZXJ0aWVzJyxcbiAgICAgICAgICAgIH0pO1xuXG4gICAgICAgICAgICBmb3IgKGNvbnN0IGNzc1Byb3Agb2YgZGVjbC5jc3NQcm9wZXJ0aWVzKSB7XG4gICAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgICAgICB0eXBlOiAnY3NzLXByb3BlcnR5JyxcbiAgICAgICAgICAgICAgICBsYWJlbDogY3NzUHJvcC5uYW1lLFxuICAgICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDMsXG4gICAgICAgICAgICAgICAgcGFyZW50SWQ6IGNzc1Byb3BDYXRJZCxcbiAgICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgICAgIHRhZ05hbWU6IGRlY2wudGFnTmFtZSxcbiAgICAgICAgICAgICAgICBuYW1lOiBjc3NQcm9wLm5hbWUsXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cblxuICAgICAgICAgIC8vIENTUyBQYXJ0c1xuICAgICAgICAgIGlmIChkZWNsLmNzc1BhcnRzPy5sZW5ndGgpIHtcbiAgICAgICAgICAgIGNvbnN0IGNzc1BhcnRDYXRJZCA9IGlkKys7XG4gICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgIGlkOiBjc3NQYXJ0Q2F0SWQsXG4gICAgICAgICAgICAgIHR5cGU6ICdjYXRlZ29yeScsXG4gICAgICAgICAgICAgIGxhYmVsOiAnQ1NTIFBhcnRzJyxcbiAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMixcbiAgICAgICAgICAgICAgcGFyZW50SWQ6IGNlSWQsXG4gICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiB0cnVlLFxuICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICBiYWRnZTogZGVjbC5jc3NQYXJ0cy5sZW5ndGgsXG4gICAgICAgICAgICAgIGNhdGVnb3J5OiAnY3NzLXBhcnRzJyxcbiAgICAgICAgICAgIH0pO1xuXG4gICAgICAgICAgICBmb3IgKGNvbnN0IGNzc1BhcnQgb2YgZGVjbC5jc3NQYXJ0cykge1xuICAgICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICAgICAgdHlwZTogJ2Nzcy1wYXJ0JyxcbiAgICAgICAgICAgICAgICBsYWJlbDogY3NzUGFydC5uYW1lLFxuICAgICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDMsXG4gICAgICAgICAgICAgICAgcGFyZW50SWQ6IGNzc1BhcnRDYXRJZCxcbiAgICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgICAgIHRhZ05hbWU6IGRlY2wudGFnTmFtZSxcbiAgICAgICAgICAgICAgICBuYW1lOiBjc3NQYXJ0Lm5hbWUsXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cblxuICAgICAgICAgIC8vIENTUyBTdGF0ZXNcbiAgICAgICAgICBpZiAoZGVjbC5jc3NTdGF0ZXM/Lmxlbmd0aCkge1xuICAgICAgICAgICAgY29uc3QgY3NzU3RhdGVDYXRJZCA9IGlkKys7XG4gICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgIGlkOiBjc3NTdGF0ZUNhdElkLFxuICAgICAgICAgICAgICB0eXBlOiAnY2F0ZWdvcnknLFxuICAgICAgICAgICAgICBsYWJlbDogJ0NTUyBTdGF0ZXMnLFxuICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAyLFxuICAgICAgICAgICAgICBwYXJlbnRJZDogY2VJZCxcbiAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IHRydWUsXG4gICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgIGJhZGdlOiBkZWNsLmNzc1N0YXRlcy5sZW5ndGgsXG4gICAgICAgICAgICAgIGNhdGVnb3J5OiAnY3NzLXN0YXRlcycsXG4gICAgICAgICAgICB9KTtcblxuICAgICAgICAgICAgZm9yIChjb25zdCBjc3NTdGF0ZSBvZiBkZWNsLmNzc1N0YXRlcykge1xuICAgICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICAgICAgdHlwZTogJ2Nzcy1zdGF0ZScsXG4gICAgICAgICAgICAgICAgbGFiZWw6IGNzc1N0YXRlLm5hbWUsXG4gICAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMyxcbiAgICAgICAgICAgICAgICBwYXJlbnRJZDogY3NzU3RhdGVDYXRJZCxcbiAgICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgICAgIHRhZ05hbWU6IGRlY2wudGFnTmFtZSxcbiAgICAgICAgICAgICAgICBuYW1lOiBjc3NTdGF0ZS5uYW1lLFxuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG5cbiAgICAgICAgICAvLyBEZW1vc1xuICAgICAgICAgIGlmIChkZWNsLmRlbW9zPy5sZW5ndGgpIHtcbiAgICAgICAgICAgIGNvbnN0IGRlbW9DYXRJZCA9IGlkKys7XG4gICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgIGlkOiBkZW1vQ2F0SWQsXG4gICAgICAgICAgICAgIHR5cGU6ICdjYXRlZ29yeScsXG4gICAgICAgICAgICAgIGxhYmVsOiAnRGVtb3MnLFxuICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAyLFxuICAgICAgICAgICAgICBwYXJlbnRJZDogY2VJZCxcbiAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IHRydWUsXG4gICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgIGJhZGdlOiBkZWNsLmRlbW9zLmxlbmd0aCxcbiAgICAgICAgICAgICAgY2F0ZWdvcnk6ICdkZW1vcycsXG4gICAgICAgICAgICB9KTtcblxuICAgICAgICAgICAgZm9yIChjb25zdCBkZW1vIG9mIGRlY2wuZGVtb3MpIHtcbiAgICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICAgIGlkOiBpZCsrLFxuICAgICAgICAgICAgICAgIHR5cGU6ICdkZW1vJyxcbiAgICAgICAgICAgICAgICBsYWJlbDogZGVtby5uYW1lIHx8IGRlbW8udXJsIHx8ICcoZGVtbyknLFxuICAgICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDMsXG4gICAgICAgICAgICAgICAgcGFyZW50SWQ6IGRlbW9DYXRJZCxcbiAgICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgICAgIHRhZ05hbWU6IGRlY2wudGFnTmFtZSxcbiAgICAgICAgICAgICAgICB1cmw6IGRlbW8udXJsLFxuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgIH0gZWxzZSBpZiAoZGVjbC5raW5kID09PSAnY2xhc3MnKSB7XG4gICAgICAgICAgLy8gUmVndWxhciBjbGFzc1xuICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgIGlkOiBpZCsrLFxuICAgICAgICAgICAgdHlwZTogJ2NsYXNzJyxcbiAgICAgICAgICAgIGxhYmVsOiBkZWNsLm5hbWUsXG4gICAgICAgICAgICBkZXB0aDogZGVwdGggKyAxLFxuICAgICAgICAgICAgcGFyZW50SWQ6IG1vZHVsZUlkLFxuICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgIG5hbWU6IGRlY2wubmFtZSxcbiAgICAgICAgICB9KTtcbiAgICAgICAgfSBlbHNlIGlmIChkZWNsLmtpbmQgPT09ICdmdW5jdGlvbicpIHtcbiAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgIHR5cGU6ICdmdW5jdGlvbicsXG4gICAgICAgICAgICBsYWJlbDogYCR7ZGVjbC5uYW1lfSgpYCxcbiAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDEsXG4gICAgICAgICAgICBwYXJlbnRJZDogbW9kdWxlSWQsXG4gICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgbmFtZTogZGVjbC5uYW1lLFxuICAgICAgICAgIH0pO1xuICAgICAgICB9IGVsc2UgaWYgKGRlY2wua2luZCA9PT0gJ3ZhcmlhYmxlJykge1xuICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgIGlkOiBpZCsrLFxuICAgICAgICAgICAgdHlwZTogJ3ZhcmlhYmxlJyxcbiAgICAgICAgICAgIGxhYmVsOiBkZWNsLm5hbWUsXG4gICAgICAgICAgICBkZXB0aDogZGVwdGggKyAxLFxuICAgICAgICAgICAgcGFyZW50SWQ6IG1vZHVsZUlkLFxuICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgIG5hbWU6IGRlY2wubmFtZSxcbiAgICAgICAgICB9KTtcbiAgICAgICAgfSBlbHNlIGlmIChkZWNsLmtpbmQgPT09ICdtaXhpbicpIHtcbiAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgIHR5cGU6ICdtaXhpbicsXG4gICAgICAgICAgICBsYWJlbDogYCR7ZGVjbC5uYW1lfSgpYCxcbiAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDEsXG4gICAgICAgICAgICBwYXJlbnRJZDogbW9kdWxlSWQsXG4gICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgbmFtZTogZGVjbC5uYW1lLFxuICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgcmV0dXJuIGlkO1xuICB9XG5cbiAgLyoqXG4gICAqIFVwZGF0ZSB2aXNpYmxlIGl0ZW1zIGJhc2VkIG9uIGV4cGFuZGVkIHN0YXRlXG4gICAqL1xuICAjdXBkYXRlVmlzaWJsZUl0ZW1zKCkge1xuICAgIHRoaXMuI3Zpc2libGVJdGVtcyA9IHRoaXMuI2ZsYXRJdGVtcy5maWx0ZXIoaXRlbSA9PiBpdGVtLnZpc2libGUpO1xuICB9XG5cbiAgLyoqXG4gICAqIFJlbmRlciB2aXNpYmxlIHRyZWUgaXRlbXMgd2l0aCBwcm9wZXIgbmVzdGluZ1xuICAgKi9cbiAgI3JlbmRlclRyZWUoKSB7XG4gICAgaWYgKCF0aGlzLiN2aWV3cG9ydCkgcmV0dXJuO1xuXG4gICAgLy8gQ2xlYXIgdmlld3BvcnRcbiAgICB0aGlzLiN2aWV3cG9ydC5pbm5lckhUTUwgPSAnJztcblxuICAgIC8vIEJ1aWxkIG5lc3RlZCB0cmVlIHN0cnVjdHVyZSBmcm9tIGZsYXQgdmlzaWJsZSBpdGVtc1xuICAgIGNvbnN0IHJvb3RJdGVtcyA9IHRoaXMuI3Zpc2libGVJdGVtcy5maWx0ZXIoaXRlbSA9PiBpdGVtLmRlcHRoID09PSAwKTtcblxuICAgIGZvciAoY29uc3Qgcm9vdEl0ZW0gb2Ygcm9vdEl0ZW1zKSB7XG4gICAgICBjb25zdCB0cmVlSXRlbUVsID0gdGhpcy4jY3JlYXRlVHJlZUl0ZW1XaXRoQ2hpbGRyZW4ocm9vdEl0ZW0pO1xuICAgICAgaWYgKHRyZWVJdGVtRWwpIHtcbiAgICAgICAgdGhpcy4jdmlld3BvcnQuYXBwZW5kQ2hpbGQodHJlZUl0ZW1FbCk7XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIENyZWF0ZSBhIHRyZWUgaXRlbSBlbGVtZW50IHdpdGggaXRzIG5lc3RlZCBjaGlsZHJlblxuICAgKi9cbiAgI2NyZWF0ZVRyZWVJdGVtV2l0aENoaWxkcmVuKGl0ZW06IFRyZWVJdGVtKTogSFRNTEVsZW1lbnQge1xuICAgIGNvbnN0IGVsID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnY2VtLXBmLXY2LXRyZWUtaXRlbScpO1xuICAgIGVsLnNldEF0dHJpYnV0ZSgnbGFiZWwnLCBpdGVtLmxhYmVsKTtcblxuICAgIGlmIChpdGVtLmJhZGdlKSB7XG4gICAgICBlbC5zZXRBdHRyaWJ1dGUoJ2JhZGdlJywgU3RyaW5nKGl0ZW0uYmFkZ2UpKTtcbiAgICB9XG5cbiAgICAvLyBTdG9yZSBpdGVtIGRhdGFcbiAgICBlbC5kYXRhc2V0Lml0ZW1JZCA9IFN0cmluZyhpdGVtLmlkKTtcbiAgICBlbC5kYXRhc2V0LnR5cGUgPSBpdGVtLnR5cGU7XG5cbiAgICAvLyBSZXN0b3JlIHNlbGVjdGlvbiBzdGF0ZSBpZiB0aGlzIGl0ZW0gd2FzIHNlbGVjdGVkXG4gICAgaWYgKHRoaXMuI2N1cnJlbnRTZWxlY3RlZEl0ZW1JZCA9PT0gaXRlbS5pZCkge1xuICAgICAgZWwuc2V0QXR0cmlidXRlKCdjdXJyZW50JywgJycpO1xuICAgICAgdGhpcy4jY3VycmVudFNlbGVjdGVkRWxlbWVudCA9IGVsO1xuICAgIH1cblxuICAgIC8vIFNldCBoYXMtY2hpbGRyZW4gYXR0cmlidXRlIGZvciBpdGVtcyB3aXRoIGNoaWxkcmVuIChzaG93cyB0b2dnbGUgYnV0dG9uKVxuICAgIGlmIChpdGVtLmhhc0NoaWxkcmVuKSB7XG4gICAgICBlbC5zZXRBdHRyaWJ1dGUoJ2hhcy1jaGlsZHJlbicsICcnKTtcbiAgICAgIChlbCBhcyBhbnkpLmV4cGFuZGVkID0gaXRlbS5leHBhbmRlZDtcblxuICAgICAgLy8gTGlzdGVuIGZvciBleHBhbmQvY29sbGFwc2VcbiAgICAgIC8vIFN0b3AgcHJvcGFnYXRpb24gdG8gcHJldmVudCBuZXN0ZWQgY29sbGFwc2VzIGZyb20gYWZmZWN0aW5nIHBhcmVudCBpdGVtc1xuICAgICAgZWwuYWRkRXZlbnRMaXN0ZW5lcignZXhwYW5kJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICAgIGUuc3RvcFByb3BhZ2F0aW9uKCk7XG4gICAgICAgIHRoaXMuI2hhbmRsZVRvZ2dsZShpdGVtLCB0cnVlKTtcbiAgICAgIH0pO1xuXG4gICAgICBlbC5hZGRFdmVudExpc3RlbmVyKCdjb2xsYXBzZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgICBlLnN0b3BQcm9wYWdhdGlvbigpO1xuICAgICAgICB0aGlzLiNoYW5kbGVUb2dnbGUoaXRlbSwgZmFsc2UpO1xuICAgICAgfSk7XG4gICAgfVxuXG4gICAgLy8gTGlzdGVuIGZvciBzZWxlY3Rpb24gKG5vbi1jYXRlZ29yeSBpdGVtcylcbiAgICBpZiAoaXRlbS50eXBlICE9PSAnY2F0ZWdvcnknKSB7XG4gICAgICBlbC5hZGRFdmVudExpc3RlbmVyKCdzZWxlY3QnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgICAgZS5zdG9wUHJvcGFnYXRpb24oKTtcbiAgICAgICAgdGhpcy4jaGFuZGxlU2VsZWN0KGl0ZW0sIGVsKTtcbiAgICAgIH0pO1xuICAgIH1cblxuICAgIC8vIEFkZCBjaGlsZHJlbiBpZiBleHBhbmRlZCBhbmQgdmlzaWJsZVxuICAgIGlmIChpdGVtLmV4cGFuZGVkICYmIGl0ZW0uaGFzQ2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IGNoaWxkcmVuID0gdGhpcy4jdmlzaWJsZUl0ZW1zLmZpbHRlcihjaGlsZCA9PiBjaGlsZC5wYXJlbnRJZCA9PT0gaXRlbS5pZCAmJiBjaGlsZC52aXNpYmxlKTtcbiAgICAgIGZvciAoY29uc3QgY2hpbGQgb2YgY2hpbGRyZW4pIHtcbiAgICAgICAgY29uc3QgY2hpbGRFbCA9IHRoaXMuI2NyZWF0ZVRyZWVJdGVtV2l0aENoaWxkcmVuKGNoaWxkKTtcbiAgICAgICAgaWYgKGNoaWxkRWwpIHtcbiAgICAgICAgICBlbC5hcHBlbmRDaGlsZChjaGlsZEVsKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiBlbDtcbiAgfVxuXG4gIC8qKlxuICAgKiBIYW5kbGUgdHJlZSBpdGVtIHRvZ2dsZSAoZXhwYW5kL2NvbGxhcHNlKVxuICAgKi9cbiAgI2hhbmRsZVRvZ2dsZShpdGVtOiBUcmVlSXRlbSwgZXhwYW5kZWQ6IGJvb2xlYW4pIHtcbiAgICBpdGVtLmV4cGFuZGVkID0gZXhwYW5kZWQ7XG4gICAgdGhpcy4jdXBkYXRlQ2hpbGRWaXNpYmlsaXR5KGl0ZW0pO1xuICAgIHRoaXMuI3VwZGF0ZVZpc2libGVJdGVtcygpO1xuICAgIHRoaXMuI3JlbmRlclRyZWUoKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBVcGRhdGUgdmlzaWJpbGl0eSBvZiBjaGlsZHJlbiB3aGVuIHBhcmVudCBpcyB0b2dnbGVkXG4gICAqL1xuICAjdXBkYXRlQ2hpbGRWaXNpYmlsaXR5KHBhcmVudEl0ZW06IFRyZWVJdGVtKSB7XG4gICAgY29uc3QgaXNFeHBhbmRlZCA9IHBhcmVudEl0ZW0uZXhwYW5kZWQ7XG5cbiAgICBmb3IgKGNvbnN0IGl0ZW0gb2YgdGhpcy4jZmxhdEl0ZW1zKSB7XG4gICAgICBpZiAoaXRlbS5wYXJlbnRJZCA9PT0gcGFyZW50SXRlbS5pZCkge1xuICAgICAgICBpdGVtLnZpc2libGUgPSBpc0V4cGFuZGVkO1xuICAgICAgICAvLyBJZiBjb2xsYXBzaW5nLCBhbHNvIGNvbGxhcHNlIGFuZCBoaWRlIGFsbCBkZXNjZW5kYW50c1xuICAgICAgICBpZiAoIWlzRXhwYW5kZWQpIHtcbiAgICAgICAgICBpdGVtLmV4cGFuZGVkID0gZmFsc2U7XG4gICAgICAgICAgdGhpcy4jdXBkYXRlQ2hpbGRWaXNpYmlsaXR5KGl0ZW0pO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEhhbmRsZSB0cmVlIGl0ZW0gc2VsZWN0aW9uXG4gICAqL1xuICAjaGFuZGxlU2VsZWN0KGl0ZW06IFRyZWVJdGVtLCBlbGVtZW50OiBFbGVtZW50KSB7XG4gICAgLy8gQ2xlYXIgcHJldmlvdXMgc2VsZWN0aW9uXG4gICAgaWYgKHRoaXMuI2N1cnJlbnRTZWxlY3RlZEVsZW1lbnQpIHtcbiAgICAgIHRoaXMuI2N1cnJlbnRTZWxlY3RlZEVsZW1lbnQucmVtb3ZlQXR0cmlidXRlKCdjdXJyZW50Jyk7XG4gICAgfVxuXG4gICAgLy8gU2V0IG5ldyBzZWxlY3Rpb25cbiAgICBlbGVtZW50LnNldEF0dHJpYnV0ZSgnY3VycmVudCcsICcnKTtcbiAgICB0aGlzLiNjdXJyZW50U2VsZWN0ZWRFbGVtZW50ID0gZWxlbWVudDtcbiAgICB0aGlzLiNjdXJyZW50U2VsZWN0ZWRJdGVtSWQgPSBpdGVtLmlkO1xuXG4gICAgLy8gRGlzcGF0Y2ggY3VzdG9tIGV2ZW50XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBJdGVtU2VsZWN0RXZlbnQoaXRlbSkpO1xuICB9XG5cbiAgLyoqXG4gICAqIFNlYXJjaC9maWx0ZXIgdHJlZSBpdGVtc1xuICAgKi9cbiAgc2VhcmNoKHF1ZXJ5OiBzdHJpbmcpIHtcbiAgICB0aGlzLiNzZWFyY2hRdWVyeSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCkudHJpbSgpO1xuXG4gICAgaWYgKCF0aGlzLiNzZWFyY2hRdWVyeSkge1xuICAgICAgLy8gUmVzZXQgdG8gbm9ybWFsIHZpc2liaWxpdHlcbiAgICAgIGZvciAoY29uc3QgaXRlbSBvZiB0aGlzLiNmbGF0SXRlbXMpIHtcbiAgICAgICAgaXRlbS52aXNpYmxlID0gaXRlbS5kZXB0aCA9PT0gMCB8fCAoaXRlbS5wYXJlbnRJZCAhPT0gdW5kZWZpbmVkICYmIHRoaXMuI2lzUGFyZW50RXhwYW5kZWQoaXRlbSkpO1xuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICAvLyBGaWx0ZXIgYmFzZWQgb24gc2VhcmNoXG4gICAgICBjb25zdCBtYXRjaGluZ0l0ZW1zID0gbmV3IFNldDxUcmVlSXRlbT4oKTtcblxuICAgICAgLy8gRmluZCBtYXRjaGluZyBpdGVtc1xuICAgICAgZm9yIChjb25zdCBpdGVtIG9mIHRoaXMuI2ZsYXRJdGVtcykge1xuICAgICAgICBpZiAoaXRlbS5sYWJlbC50b0xvd2VyQ2FzZSgpLmluY2x1ZGVzKHRoaXMuI3NlYXJjaFF1ZXJ5KSkge1xuICAgICAgICAgIG1hdGNoaW5nSXRlbXMuYWRkKGl0ZW0pO1xuICAgICAgICAgIC8vIFNob3cgYW5kIGV4cGFuZCBhbGwgYW5jZXN0b3JzXG4gICAgICAgICAgdGhpcy4jc2hvd0FuY2VzdG9ycyhpdGVtKTtcbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICAvLyBIaWRlIG5vbi1tYXRjaGluZyBpdGVtc1xuICAgICAgZm9yIChjb25zdCBpdGVtIG9mIHRoaXMuI2ZsYXRJdGVtcykge1xuICAgICAgICBpZiAoIW1hdGNoaW5nSXRlbXMuaGFzKGl0ZW0pICYmICF0aGlzLiNoYXNNYXRjaGluZ0Rlc2NlbmRhbnQoaXRlbSwgbWF0Y2hpbmdJdGVtcykpIHtcbiAgICAgICAgICBpdGVtLnZpc2libGUgPSB0aGlzLiNpc0FuY2VzdG9yT2ZNYXRjaChpdGVtLCBtYXRjaGluZ0l0ZW1zKTtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBpdGVtLnZpc2libGUgPSB0cnVlO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgdGhpcy4jdXBkYXRlVmlzaWJsZUl0ZW1zKCk7XG4gICAgdGhpcy4jcmVuZGVyVHJlZSgpO1xuICB9XG5cbiAgLyoqXG4gICAqIENoZWNrIGlmIHBhcmVudCBpcyBleHBhbmRlZFxuICAgKi9cbiAgI2lzUGFyZW50RXhwYW5kZWQoaXRlbTogVHJlZUl0ZW0pOiBib29sZWFuIHtcbiAgICBpZiAoaXRlbS5wYXJlbnRJZCA9PT0gdW5kZWZpbmVkKSByZXR1cm4gdHJ1ZTtcbiAgICBjb25zdCBwYXJlbnQgPSB0aGlzLiNmbGF0SXRlbXMuZmluZChpID0+IGkuaWQgPT09IGl0ZW0ucGFyZW50SWQpO1xuICAgIHJldHVybiAhIXBhcmVudCAmJiBwYXJlbnQuZXhwYW5kZWQgJiYgdGhpcy4jaXNQYXJlbnRFeHBhbmRlZChwYXJlbnQpO1xuICB9XG5cbiAgLyoqXG4gICAqIFNob3cgYWxsIGFuY2VzdG9ycyBvZiBhbiBpdGVtXG4gICAqL1xuICAjc2hvd0FuY2VzdG9ycyhpdGVtOiBUcmVlSXRlbSkge1xuICAgIGlmIChpdGVtLnBhcmVudElkID09PSB1bmRlZmluZWQpIHJldHVybjtcbiAgICBjb25zdCBwYXJlbnQgPSB0aGlzLiNmbGF0SXRlbXMuZmluZChpID0+IGkuaWQgPT09IGl0ZW0ucGFyZW50SWQpO1xuICAgIGlmIChwYXJlbnQpIHtcbiAgICAgIHBhcmVudC52aXNpYmxlID0gdHJ1ZTtcbiAgICAgIHBhcmVudC5leHBhbmRlZCA9IHRydWU7XG4gICAgICB0aGlzLiNzaG93QW5jZXN0b3JzKHBhcmVudCk7XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIENoZWNrIGlmIGl0ZW0gaXMgYW5jZXN0b3Igb2YgYSBtYXRjaFxuICAgKi9cbiAgI2lzQW5jZXN0b3JPZk1hdGNoKGl0ZW06IFRyZWVJdGVtLCBtYXRjaGluZ0l0ZW1zOiBTZXQ8VHJlZUl0ZW0+KTogYm9vbGVhbiB7XG4gICAgZm9yIChjb25zdCBtYXRjaCBvZiBtYXRjaGluZ0l0ZW1zKSB7XG4gICAgICBsZXQgY3VycmVudDogVHJlZUl0ZW0gPSBtYXRjaDtcbiAgICAgIHdoaWxlIChjdXJyZW50LnBhcmVudElkICE9PSB1bmRlZmluZWQpIHtcbiAgICAgICAgY29uc3QgcGFyZW50ID0gdGhpcy4jZmxhdEl0ZW1zLmZpbmQoaSA9PiBpLmlkID09PSBjdXJyZW50LnBhcmVudElkKTtcbiAgICAgICAgaWYgKCFwYXJlbnQpIGJyZWFrO1xuICAgICAgICBpZiAocGFyZW50ID09PSBpdGVtKSByZXR1cm4gdHJ1ZTtcbiAgICAgICAgY3VycmVudCA9IHBhcmVudDtcbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuIGZhbHNlO1xuICB9XG5cbiAgLyoqXG4gICAqIENoZWNrIGlmIGl0ZW0gaGFzIG1hdGNoaW5nIGRlc2NlbmRhbnRcbiAgICovXG4gICNoYXNNYXRjaGluZ0Rlc2NlbmRhbnQoaXRlbTogVHJlZUl0ZW0sIG1hdGNoaW5nSXRlbXM6IFNldDxUcmVlSXRlbT4pOiBib29sZWFuIHtcbiAgICBmb3IgKGNvbnN0IG1hdGNoIG9mIG1hdGNoaW5nSXRlbXMpIHtcbiAgICAgIGlmICh0aGlzLiNpc0Rlc2NlbmRhbnRPZihtYXRjaCwgaXRlbSkpIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICByZXR1cm4gZmFsc2U7XG4gIH1cblxuICAvKipcbiAgICogQ2hlY2sgaWYgaXRlbSBpcyBkZXNjZW5kYW50IG9mIGFub3RoZXJcbiAgICovXG4gICNpc0Rlc2NlbmRhbnRPZihpdGVtOiBUcmVlSXRlbSwgYW5jZXN0b3I6IFRyZWVJdGVtKTogYm9vbGVhbiB7XG4gICAgbGV0IGN1cnJlbnQ6IFRyZWVJdGVtID0gaXRlbTtcbiAgICB3aGlsZSAoY3VycmVudC5wYXJlbnRJZCAhPT0gdW5kZWZpbmVkKSB7XG4gICAgICBjb25zdCBwYXJlbnQgPSB0aGlzLiNmbGF0SXRlbXMuZmluZChpID0+IGkuaWQgPT09IGN1cnJlbnQucGFyZW50SWQpO1xuICAgICAgaWYgKCFwYXJlbnQpIHJldHVybiBmYWxzZTtcbiAgICAgIGlmIChwYXJlbnQgPT09IGFuY2VzdG9yKSByZXR1cm4gdHJ1ZTtcbiAgICAgIGN1cnJlbnQgPSBwYXJlbnQ7XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBFeHBhbmQgYWxsIGl0ZW1zXG4gICAqL1xuICBleHBhbmRBbGwoKSB7XG4gICAgZm9yIChjb25zdCBpdGVtIG9mIHRoaXMuI2ZsYXRJdGVtcykge1xuICAgICAgaWYgKGl0ZW0uaGFzQ2hpbGRyZW4pIHtcbiAgICAgICAgaXRlbS5leHBhbmRlZCA9IHRydWU7XG4gICAgICB9XG4gICAgICAvLyBNYWtlIGFsbCBpdGVtcyB2aXNpYmxlIGV4Y2VwdCB0aG9zZSBmaWx0ZXJlZCBieSBzZWFyY2hcbiAgICAgIGlmICghdGhpcy4jc2VhcmNoUXVlcnkgfHwgaXRlbS5sYWJlbC50b0xvd2VyQ2FzZSgpLmluY2x1ZGVzKHRoaXMuI3NlYXJjaFF1ZXJ5KSkge1xuICAgICAgICBpdGVtLnZpc2libGUgPSB0cnVlO1xuICAgICAgfVxuICAgIH1cbiAgICB0aGlzLiN1cGRhdGVWaXNpYmxlSXRlbXMoKTtcbiAgICB0aGlzLiNyZW5kZXJUcmVlKCk7XG4gIH1cblxuICAvKipcbiAgICogQ29sbGFwc2UgYWxsIGl0ZW1zXG4gICAqL1xuICBjb2xsYXBzZUFsbCgpIHtcbiAgICBmb3IgKGNvbnN0IGl0ZW0gb2YgdGhpcy4jZmxhdEl0ZW1zKSB7XG4gICAgICBpZiAoaXRlbS5oYXNDaGlsZHJlbikge1xuICAgICAgICBpdGVtLmV4cGFuZGVkID0gZmFsc2U7XG4gICAgICB9XG4gICAgICAvLyBPbmx5IHRvcC1sZXZlbCBpdGVtcyB2aXNpYmxlXG4gICAgICBpdGVtLnZpc2libGUgPSBpdGVtLmRlcHRoID09PSAwO1xuICAgIH1cbiAgICB0aGlzLiN1cGRhdGVWaXNpYmxlSXRlbXMoKTtcbiAgICB0aGlzLiNyZW5kZXJUcmVlKCk7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXZpcnR1YWwtdHJlZSc6IENlbVZpcnR1YWxUcmVlO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCIvKiBWaXJ0dWFsIFRyZWUgLSBFZmZpY2llbnRseSByZW5kZXJzIGxhcmdlIHRyZWUgc3RydWN0dXJlcyAqL1xcXFxuXFxcXG46aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG59XFxcXG5cXFxcbiN0cmVlLWNvbnRhaW5lciB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG59XFxcXG5cXFxcbiN2aWV3cG9ydCB7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbiAgb3ZlcmZsb3cteDogaGlkZGVuO1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG5cXFxcbiAgY2VtLXBmLXY2LXRyZWUtaXRlbSB7XFxcXG4gICAgY3Vyc29yOiBwb2ludGVyO1xcXFxuXFxcXG4gICAgXFxcXHUwMDI2OmhvdmVyIHtcXFxcbiAgICAgIGJhY2tncm91bmQtY29sb3I6IHZhcigtLXBmLXQtLWdsb2JhbC0tYmFja2dyb3VuZC0tY29sb3ItLXNlY29uZGFyeS0taG92ZXIsICNmNWY1ZjUpO1xcXFxuICAgIH1cXFxcblxcXFxuICAgIC8qIE9ubHkgb3V0bGluZSB0aGUgYWN0dWFsbHkgZm9jdXNlZCBpdGVtLCBub3QgYW5jZXN0b3JzICovXFxcXG4gICAgXFxcXHUwMDI2OmZvY3VzLXZpc2libGUge1xcXFxuICAgICAgb3V0bGluZTogMnB4IHNvbGlkIHZhcigtLXBmLXQtLWdsb2JhbC0tY29sb3ItLWJyYW5kLS1kZWZhdWx0LCAjMDA2NmNjKTtcXFxcbiAgICAgIG91dGxpbmUtb2Zmc2V0OiAtMnB4O1xcXFxuICAgIH1cXFxcbiAgfVxcXFxufVxcXFxuXFxcIlwiKSk7ZXhwb3J0IGRlZmF1bHQgczsiXSwKICAibWFwcGluZ3MiOiAiOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7O0FBQUEsU0FBUyxZQUFZLFlBQVk7QUFDakMsU0FBUyxxQkFBcUI7OztBQ0Q5QixJQUFNLElBQUUsSUFBSSxjQUFjO0FBQUUsRUFBRSxZQUFZLEtBQUssTUFBTSxvdEJBQXN0QixDQUFDO0FBQUUsSUFBTywyQkFBUTs7O0FESzd4QixPQUFPO0FBMkRBLElBQU0sa0JBQU4sY0FBOEIsTUFBTTtBQUFBLEVBQ3pDO0FBQUEsRUFDQSxZQUFZLE1BQWdCO0FBQzFCLFVBQU0sZUFBZSxFQUFFLFNBQVMsTUFBTSxVQUFVLEtBQUssQ0FBQztBQUN0RCxTQUFLLE9BQU87QUFBQSxFQUNkO0FBQ0Y7QUF0RUE7QUE2RUEsOEJB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
