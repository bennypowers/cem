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
s.replaceSync(JSON.parse('"/* Virtual Tree - Efficiently renders large tree structures */\\n\\n:host {\\n  display: block;\\n  height: 100%;\\n  overflow: hidden;\\n}\\n\\n#tree-container {\\n  position: relative;\\n  width: 100%;\\n}\\n\\n#viewport {\\n  height: 100%;\\n  overflow-y: auto;\\n  overflow-x: hidden;\\n  position: relative;\\n\\n  pf-v6-tree-item {\\n    cursor: pointer;\\n\\n    \\u0026:hover {\\n      background-color: var(--pf-t--global--background--color--secondary--hover, #f5f5f5);\\n    }\\n\\n    /* Only outline the actually focused item, not ancestors */\\n    \\u0026:focus-visible {\\n      outline: 2px solid var(--pf-t--global--color--brand--default, #0066cc);\\n      outline-offset: -2px;\\n    }\\n  }\\n}\\n"'));
var cem_virtual_tree_default = s;

// elements/cem-virtual-tree/cem-virtual-tree.ts
import "../pf-v6-tree-item/pf-v6-tree-item.js";
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
    const el = document.createElement("pf-v6-tree-item");
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsiLi4vLi4vLi4vLi4vLi4vZWxlbWVudHMvY2VtLXZpcnR1YWwtdHJlZS9jZW0tdmlydHVhbC10cmVlLnRzIiwgImxpdC1jc3M6ZWxlbWVudHMvY2VtLXZpcnR1YWwtdHJlZS9jZW0tdmlydHVhbC10cmVlLmNzcyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiaW1wb3J0IHsgTGl0RWxlbWVudCwgaHRtbCB9IGZyb20gJ2xpdCc7XG5pbXBvcnQgeyBjdXN0b21FbGVtZW50IH0gZnJvbSAnbGl0L2RlY29yYXRvcnMvY3VzdG9tLWVsZW1lbnQuanMnO1xuXG5pbXBvcnQgc3R5bGVzIGZyb20gJy4vY2VtLXZpcnR1YWwtdHJlZS5jc3MnO1xuXG5pbXBvcnQgJy4uL3BmLXY2LXRyZWUtaXRlbS9wZi12Ni10cmVlLWl0ZW0uanMnO1xuXG5pbnRlcmZhY2UgVHJlZUl0ZW0ge1xuICBpZDogbnVtYmVyO1xuICB0eXBlOiBzdHJpbmc7XG4gIGxhYmVsOiBzdHJpbmc7XG4gIGRlcHRoOiBudW1iZXI7XG4gIHBhcmVudElkPzogbnVtYmVyO1xuICBoYXNDaGlsZHJlbjogYm9vbGVhbjtcbiAgZXhwYW5kZWQ6IGJvb2xlYW47XG4gIHZpc2libGU6IGJvb2xlYW47XG4gIHBhY2thZ2VOYW1lPzogc3RyaW5nO1xuICBtb2R1bGVQYXRoPzogc3RyaW5nO1xuICB0YWdOYW1lPzogc3RyaW5nO1xuICBuYW1lPzogc3RyaW5nO1xuICB1cmw/OiBzdHJpbmc7XG4gIGJhZGdlPzogbnVtYmVyO1xuICBjYXRlZ29yeT86IHN0cmluZztcbn1cblxuaW50ZXJmYWNlIE1hbmlmZXN0IHtcbiAgcGFja2FnZXM/OiBNYW5pZmVzdFBhY2thZ2VbXTtcbiAgbW9kdWxlcz86IE1hbmlmZXN0TW9kdWxlW107XG59XG5cbmludGVyZmFjZSBNYW5pZmVzdFBhY2thZ2Uge1xuICBuYW1lOiBzdHJpbmc7XG4gIG1vZHVsZXM/OiBNYW5pZmVzdE1vZHVsZVtdO1xufVxuXG5pbnRlcmZhY2UgTWFuaWZlc3RNb2R1bGUge1xuICBwYXRoOiBzdHJpbmc7XG4gIGRlY2xhcmF0aW9ucz86IE1hbmlmZXN0RGVjbGFyYXRpb25bXTtcbn1cblxuaW50ZXJmYWNlIE1hbmlmZXN0RGVjbGFyYXRpb24ge1xuICBraW5kOiBzdHJpbmc7XG4gIG5hbWU6IHN0cmluZztcbiAgY3VzdG9tRWxlbWVudD86IGJvb2xlYW47XG4gIHRhZ05hbWU/OiBzdHJpbmc7XG4gIG1lbWJlcnM/OiBNYW5pZmVzdE1lbWJlcltdO1xuICBhdHRyaWJ1dGVzPzogeyBuYW1lOiBzdHJpbmcgfVtdO1xuICBldmVudHM/OiB7IG5hbWU6IHN0cmluZyB9W107XG4gIHNsb3RzPzogeyBuYW1lOiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH1bXTtcbiAgY3NzUHJvcGVydGllcz86IHsgbmFtZTogc3RyaW5nOyBkZXNjcmlwdGlvbj86IHN0cmluZyB9W107XG4gIGNzc1BhcnRzPzogeyBuYW1lOiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH1bXTtcbiAgY3NzU3RhdGVzPzogeyBuYW1lOiBzdHJpbmc7IGRlc2NyaXB0aW9uPzogc3RyaW5nIH1bXTtcbiAgZGVtb3M/OiB7IG5hbWU/OiBzdHJpbmc7IHVybD86IHN0cmluZzsgZGVzY3JpcHRpb24/OiBzdHJpbmcgfVtdO1xufVxuXG5pbnRlcmZhY2UgTWFuaWZlc3RNZW1iZXIge1xuICBraW5kOiBzdHJpbmc7XG4gIG5hbWU6IHN0cmluZztcbiAgdHlwZT86IHsgdGV4dDogc3RyaW5nIH07XG59XG5cbi8qKlxuICogQ3VzdG9tIGV2ZW50IGZvciBpdGVtIHNlbGVjdGlvblxuICovXG5leHBvcnQgY2xhc3MgSXRlbVNlbGVjdEV2ZW50IGV4dGVuZHMgRXZlbnQge1xuICBpdGVtOiBUcmVlSXRlbTtcbiAgY29uc3RydWN0b3IoaXRlbTogVHJlZUl0ZW0pIHtcbiAgICBzdXBlcignaXRlbS1zZWxlY3QnLCB7IGJ1YmJsZXM6IHRydWUsIGNvbXBvc2VkOiB0cnVlIH0pO1xuICAgIHRoaXMuaXRlbSA9IGl0ZW07XG4gIH1cbn1cblxuLyoqXG4gKiBWaXJ0dWFsIHRyZWUgdmlldyB3aXRoIGxhenkgbG9hZGluZ1xuICpcbiAqIEBmaXJlcyBpdGVtLXNlbGVjdCAtIEZpcmVkIHdoZW4gYSB0cmVlIGl0ZW0gaXMgc2VsZWN0ZWRcbiAqL1xuQGN1c3RvbUVsZW1lbnQoJ2NlbS12aXJ0dWFsLXRyZWUnKVxuZXhwb3J0IGNsYXNzIENlbVZpcnR1YWxUcmVlIGV4dGVuZHMgTGl0RWxlbWVudCB7XG4gIHN0YXRpYyBzdHlsZXMgPSBzdHlsZXM7XG5cbiAgLy8gU3RhdGljIGNhY2hlIGZvciBtYW5pZmVzdCAoc2hhcmVkIGFjcm9zcyBhbGwgaW5zdGFuY2VzKVxuICBzdGF0aWMgI21hbmlmZXN0Q2FjaGU6IE1hbmlmZXN0IHwgbnVsbCA9IG51bGw7XG4gIHN0YXRpYyAjbWFuaWZlc3RQcm9taXNlOiBQcm9taXNlPE1hbmlmZXN0IHwgbnVsbD4gfCBudWxsID0gbnVsbDtcblxuICAjbWFuaWZlc3Q6IE1hbmlmZXN0IHwgbnVsbCA9IG51bGw7XG4gICNmbGF0SXRlbXM6IFRyZWVJdGVtW10gPSBbXTtcbiAgI3Zpc2libGVJdGVtczogVHJlZUl0ZW1bXSA9IFtdO1xuICAjc2VhcmNoUXVlcnkgPSAnJztcbiAgI3ZpZXdwb3J0OiBIVE1MRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjY3VycmVudFNlbGVjdGVkRWxlbWVudDogRWxlbWVudCB8IG51bGwgPSBudWxsO1xuICAjY3VycmVudFNlbGVjdGVkSXRlbUlkOiBudW1iZXIgfCBudWxsID0gbnVsbDtcblxuICByZW5kZXIoKSB7XG4gICAgcmV0dXJuIGh0bWxgXG4gICAgICA8ZGl2IGlkPVwidHJlZS1jb250YWluZXJcIj5cbiAgICAgICAgPGRpdiBpZD1cInZpZXdwb3J0XCIgcm9sZT1cInRyZWVcIj48L2Rpdj5cbiAgICAgIDwvZGl2PlxuICAgIGA7XG4gIH1cblxuICBhc3luYyBmaXJzdFVwZGF0ZWQoKSB7XG4gICAgdGhpcy4jdmlld3BvcnQgPSB0aGlzLnNoYWRvd1Jvb3QhLmdldEVsZW1lbnRCeUlkKCd2aWV3cG9ydCcpO1xuXG4gICAgLy8gTG9hZCBtYW5pZmVzdCBmcm9tIHNlcnZlciB3aXRoIGNhY2hpbmdcbiAgICB0aGlzLiNtYW5pZmVzdCA9IGF3YWl0IHRoaXMuI2xvYWRNYW5pZmVzdCgpO1xuICAgIGlmICghdGhpcy4jbWFuaWZlc3QpIHtcbiAgICAgIGNvbnNvbGUud2FybignW3ZpcnR1YWwtdHJlZV0gRmFpbGVkIHRvIGxvYWQgbWFuaWZlc3QnKTtcbiAgICAgIHJldHVybjtcbiAgICB9XG5cbiAgICAvLyBCdWlsZCBmbGF0IGxpc3QgZnJvbSBtYW5pZmVzdFxuICAgIHRoaXMuI2J1aWxkRmxhdExpc3QoKTtcblxuICAgIC8vIEluaXRpYWwgcmVuZGVyXG4gICAgdGhpcy4jcmVuZGVyVHJlZSgpO1xuICB9XG5cbiAgLyoqXG4gICAqIFB1YmxpYyBzdGF0aWMgbWV0aG9kIHRvIGxvYWQgbWFuaWZlc3Qgd2l0aCBjYWNoaW5nXG4gICAqIENhbiBiZSBjYWxsZWQgYnkgb3RoZXIgY29tcG9uZW50cyB0byByZXVzZSB0aGUgc2FtZSBjYWNoZWQgbWFuaWZlc3RcbiAgICovXG4gIHN0YXRpYyBhc3luYyBsb2FkTWFuaWZlc3QoKTogUHJvbWlzZTxNYW5pZmVzdCB8IG51bGw+IHtcbiAgICAvLyBSZXR1cm4gY2FjaGVkIG1hbmlmZXN0IGlmIGF2YWlsYWJsZVxuICAgIGlmIChDZW1WaXJ0dWFsVHJlZS4jbWFuaWZlc3RDYWNoZSkge1xuICAgICAgcmV0dXJuIENlbVZpcnR1YWxUcmVlLiNtYW5pZmVzdENhY2hlO1xuICAgIH1cblxuICAgIC8vIElmIGFscmVhZHkgZmV0Y2hpbmcsIHdhaXQgZm9yIGV4aXN0aW5nIHByb21pc2VcbiAgICBpZiAoQ2VtVmlydHVhbFRyZWUuI21hbmlmZXN0UHJvbWlzZSkge1xuICAgICAgcmV0dXJuIENlbVZpcnR1YWxUcmVlLiNtYW5pZmVzdFByb21pc2U7XG4gICAgfVxuXG4gICAgLy8gRmV0Y2ggbWFuaWZlc3RcbiAgICBDZW1WaXJ0dWFsVHJlZS4jbWFuaWZlc3RQcm9taXNlID0gZmV0Y2goJy9jdXN0b20tZWxlbWVudHMuanNvbicpXG4gICAgICAudGhlbihhc3luYyAocmVzcG9uc2UpID0+IHtcbiAgICAgICAgaWYgKCFyZXNwb25zZS5vaykge1xuICAgICAgICAgIHRocm93IG5ldyBFcnJvcihgRmFpbGVkIHRvIGZldGNoIG1hbmlmZXN0OiAke3Jlc3BvbnNlLnN0YXR1c31gKTtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCBtYW5pZmVzdCA9IGF3YWl0IHJlc3BvbnNlLmpzb24oKTtcbiAgICAgICAgQ2VtVmlydHVhbFRyZWUuI21hbmlmZXN0Q2FjaGUgPSBtYW5pZmVzdDtcbiAgICAgICAgcmV0dXJuIG1hbmlmZXN0O1xuICAgICAgfSlcbiAgICAgIC5jYXRjaCgoZXJyb3IpID0+IHtcbiAgICAgICAgY29uc29sZS5lcnJvcignW3ZpcnR1YWwtdHJlZV0gRXJyb3IgbG9hZGluZyBtYW5pZmVzdDonLCBlcnJvcik7XG4gICAgICAgIENlbVZpcnR1YWxUcmVlLiNtYW5pZmVzdFByb21pc2UgPSBudWxsO1xuICAgICAgICByZXR1cm4gbnVsbDtcbiAgICAgIH0pO1xuXG4gICAgcmV0dXJuIENlbVZpcnR1YWxUcmVlLiNtYW5pZmVzdFByb21pc2U7XG4gIH1cblxuICAvKipcbiAgICogQ2xlYXIgdGhlIHN0YXRpYyBtYW5pZmVzdCBjYWNoZSAoZm9yIHRlc3RpbmcpXG4gICAqL1xuICBzdGF0aWMgY2xlYXJDYWNoZSgpIHtcbiAgICBDZW1WaXJ0dWFsVHJlZS4jbWFuaWZlc3RDYWNoZSA9IG51bGw7XG4gICAgQ2VtVmlydHVhbFRyZWUuI21hbmlmZXN0UHJvbWlzZSA9IG51bGw7XG4gIH1cblxuICAvKipcbiAgICogTG9hZCBtYW5pZmVzdCBmcm9tIHNlcnZlciB3aXRoIHN0YXRpYyBjYWNoaW5nXG4gICAqL1xuICBhc3luYyAjbG9hZE1hbmlmZXN0KCk6IFByb21pc2U8TWFuaWZlc3QgfCBudWxsPiB7XG4gICAgcmV0dXJuIENlbVZpcnR1YWxUcmVlLmxvYWRNYW5pZmVzdCgpO1xuICB9XG5cbiAgLyoqXG4gICAqIEJ1aWxkIGZsYXQgbGlzdCBmcm9tIGhpZXJhcmNoaWNhbCBtYW5pZmVzdFxuICAgKi9cbiAgI2J1aWxkRmxhdExpc3QoKSB7XG4gICAgdGhpcy4jZmxhdEl0ZW1zID0gW107XG4gICAgbGV0IGlkID0gMDtcblxuICAgIC8vIENoZWNrIGlmIHRoaXMgaXMgYSB3b3Jrc3BhY2UgbWFuaWZlc3Qgd2l0aCBtdWx0aXBsZSBwYWNrYWdlc1xuICAgIGNvbnN0IGhhc011bHRpcGxlUGFja2FnZXMgPSB0aGlzLiNtYW5pZmVzdCEucGFja2FnZXMgJiYgdGhpcy4jbWFuaWZlc3QhLnBhY2thZ2VzLmxlbmd0aCA+IDE7XG5cbiAgICBpZiAoaGFzTXVsdGlwbGVQYWNrYWdlcykge1xuICAgICAgLy8gV29ya3NwYWNlIG1vZGUgd2l0aCBtdWx0aXBsZSBwYWNrYWdlczogc2hvdyBwYWNrYWdlcyBhdCB0b3AgbGV2ZWxcbiAgICAgIGZvciAoY29uc3QgcGtnIG9mIHRoaXMuI21hbmlmZXN0IS5wYWNrYWdlcyEpIHtcbiAgICAgICAgY29uc3QgcGFja2FnZUlkID0gaWQrKztcbiAgICAgICAgY29uc3QgcGFja2FnZUl0ZW06IFRyZWVJdGVtID0ge1xuICAgICAgICAgIGlkOiBwYWNrYWdlSWQsXG4gICAgICAgICAgdHlwZTogJ3BhY2thZ2UnLFxuICAgICAgICAgIGxhYmVsOiBwa2cubmFtZSxcbiAgICAgICAgICBkZXB0aDogMCxcbiAgICAgICAgICBoYXNDaGlsZHJlbjogKHBrZy5tb2R1bGVzPy5sZW5ndGggPz8gMCkgPiAwLFxuICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICB2aXNpYmxlOiB0cnVlLFxuICAgICAgICAgIHBhY2thZ2VOYW1lOiBwa2cubmFtZSxcbiAgICAgICAgICBiYWRnZTogcGtnLm1vZHVsZXM/Lmxlbmd0aCA/PyAwLFxuICAgICAgICB9O1xuICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaChwYWNrYWdlSXRlbSk7XG5cbiAgICAgICAgaWYgKCFwa2cubW9kdWxlcykgY29udGludWU7XG5cbiAgICAgICAgLy8gQWRkIG1vZHVsZXMgdW5kZXIgdGhpcyBwYWNrYWdlIGF0IGRlcHRoIDFcbiAgICAgICAgaWQgPSB0aGlzLiNidWlsZE1vZHVsZXNGb3JQYWNrYWdlKHBrZy5tb2R1bGVzLCBwYWNrYWdlSWQsIDEsIGlkKTtcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgLy8gU2luZ2xlIHBhY2thZ2UgbW9kZSBPUiB3b3Jrc3BhY2Ugd2l0aCAxIHBhY2thZ2U6IHNob3cgbW9kdWxlcyBhdCB0b3AgbGV2ZWxcbiAgICAgIGNvbnN0IG1vZHVsZXMgPSB0aGlzLiNtYW5pZmVzdCEucGFja2FnZXM/LlswXT8ubW9kdWxlcyA/PyB0aGlzLiNtYW5pZmVzdCEubW9kdWxlcztcblxuICAgICAgaWYgKCFtb2R1bGVzKSByZXR1cm47XG5cbiAgICAgIC8vIEFkZCBtb2R1bGVzIGF0IGRlcHRoIDAgKG5vIHBhY2thZ2UgbGV2ZWwpXG4gICAgICB0aGlzLiNidWlsZE1vZHVsZXNGb3JQYWNrYWdlKG1vZHVsZXMsIHVuZGVmaW5lZCwgMCwgaWQpO1xuICAgIH1cblxuICAgIHRoaXMuI3VwZGF0ZVZpc2libGVJdGVtcygpO1xuICB9XG5cbiAgLyoqXG4gICAqIEJ1aWxkIG1vZHVsZXMgYW5kIHRoZWlyIGRlY2xhcmF0aW9ucyBmb3IgYSBwYWNrYWdlXG4gICAqL1xuICAjYnVpbGRNb2R1bGVzRm9yUGFja2FnZShtb2R1bGVzOiBNYW5pZmVzdE1vZHVsZVtdLCBwYXJlbnRJZDogbnVtYmVyIHwgdW5kZWZpbmVkLCBkZXB0aDogbnVtYmVyLCBzdGFydElkOiBudW1iZXIpOiBudW1iZXIge1xuICAgIGxldCBpZCA9IHN0YXJ0SWQ7XG5cbiAgICBmb3IgKGNvbnN0IG1vZHVsZSBvZiBtb2R1bGVzKSB7XG4gICAgICBjb25zdCBtb2R1bGVJZCA9IGlkKys7XG4gICAgICBjb25zdCBtb2R1bGVJdGVtOiBUcmVlSXRlbSA9IHtcbiAgICAgICAgaWQ6IG1vZHVsZUlkLFxuICAgICAgICB0eXBlOiAnbW9kdWxlJyxcbiAgICAgICAgbGFiZWw6IG1vZHVsZS5wYXRoLFxuICAgICAgICBkZXB0aDogZGVwdGgsXG4gICAgICAgIHBhcmVudElkOiBwYXJlbnRJZCxcbiAgICAgICAgaGFzQ2hpbGRyZW46IChtb2R1bGUuZGVjbGFyYXRpb25zPy5sZW5ndGggPz8gMCkgPiAwLFxuICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgIHZpc2libGU6IHBhcmVudElkID09PSB1bmRlZmluZWQsIC8vIFZpc2libGUgaWYgbm8gcGFyZW50LCBvdGhlcndpc2UgaGlkZGVuIHVudGlsIHBhcmVudCBleHBhbmRzXG4gICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICBiYWRnZTogbW9kdWxlLmRlY2xhcmF0aW9ucz8ubGVuZ3RoID8/IDAsXG4gICAgICB9O1xuICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2gobW9kdWxlSXRlbSk7XG5cbiAgICAgIGlmICghbW9kdWxlLmRlY2xhcmF0aW9ucykgY29udGludWU7XG5cbiAgICAgIGZvciAoY29uc3QgZGVjbCBvZiBtb2R1bGUuZGVjbGFyYXRpb25zKSB7XG4gICAgICAgIC8vIEN1c3RvbSBFbGVtZW50IC0gbXVzdCBoYXZlIGJvdGggY3VzdG9tRWxlbWVudCBmbGFnIGFuZCB0YWdOYW1lXG4gICAgICAgIGlmIChkZWNsLmtpbmQgPT09ICdjbGFzcycgJiYgZGVjbC5jdXN0b21FbGVtZW50ICYmIGRlY2wudGFnTmFtZSkge1xuICAgICAgICAgIGNvbnN0IGNlSWQgPSBpZCsrO1xuXG4gICAgICAgICAgLy8gQ29tcHV0ZSBoYXNDaGlsZHJlbiBiYXNlZCBvbiBwcmVzZW5jZSBvZiBhbnkgY2hpbGQgYXJyYXlzXG4gICAgICAgICAgY29uc3QgcHJvcGVydGllcyA9IGRlY2wubWVtYmVycz8uZmlsdGVyKG0gPT4gbS5raW5kID09PSAnZmllbGQnKSA/PyBbXTtcbiAgICAgICAgICBjb25zdCBtZXRob2RzID0gZGVjbC5tZW1iZXJzPy5maWx0ZXIobSA9PiBtLmtpbmQgPT09ICdtZXRob2QnKSA/PyBbXTtcbiAgICAgICAgICBjb25zdCBoYXNDaGlsZHJlbiA9IFtcbiAgICAgICAgICAgIGRlY2wuYXR0cmlidXRlcyxcbiAgICAgICAgICAgIHByb3BlcnRpZXMsXG4gICAgICAgICAgICBtZXRob2RzLFxuICAgICAgICAgICAgZGVjbC5ldmVudHMsXG4gICAgICAgICAgICBkZWNsLnNsb3RzLFxuICAgICAgICAgICAgZGVjbC5jc3NQcm9wZXJ0aWVzLFxuICAgICAgICAgICAgZGVjbC5jc3NQYXJ0cyxcbiAgICAgICAgICAgIGRlY2wuY3NzU3RhdGVzLFxuICAgICAgICAgICAgZGVjbC5kZW1vcyxcbiAgICAgICAgICBdLnNvbWUoeCA9PiAoeD8ubGVuZ3RoID8/IDApID4gMCk7XG5cbiAgICAgICAgICBjb25zdCBjZUl0ZW06IFRyZWVJdGVtID0ge1xuICAgICAgICAgICAgaWQ6IGNlSWQsXG4gICAgICAgICAgICB0eXBlOiAnY3VzdG9tLWVsZW1lbnQnLFxuICAgICAgICAgICAgbGFiZWw6IGA8JHtkZWNsLnRhZ05hbWV9PmAsXG4gICAgICAgICAgICBkZXB0aDogZGVwdGggKyAxLFxuICAgICAgICAgICAgcGFyZW50SWQ6IG1vZHVsZUlkLFxuICAgICAgICAgICAgaGFzQ2hpbGRyZW4sXG4gICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgdGFnTmFtZTogZGVjbC50YWdOYW1lLFxuICAgICAgICAgIH07XG4gICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goY2VJdGVtKTtcblxuICAgICAgICAgIC8vIEF0dHJpYnV0ZXNcbiAgICAgICAgICBpZiAoZGVjbC5hdHRyaWJ1dGVzPy5sZW5ndGgpIHtcbiAgICAgICAgICAgIGNvbnN0IGF0dHJDYXRJZCA9IGlkKys7XG4gICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgIGlkOiBhdHRyQ2F0SWQsXG4gICAgICAgICAgICAgIHR5cGU6ICdjYXRlZ29yeScsXG4gICAgICAgICAgICAgIGxhYmVsOiAnQXR0cmlidXRlcycsXG4gICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDIsXG4gICAgICAgICAgICAgIHBhcmVudElkOiBjZUlkLFxuICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogdHJ1ZSxcbiAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgYmFkZ2U6IGRlY2wuYXR0cmlidXRlcy5sZW5ndGgsXG4gICAgICAgICAgICAgIGNhdGVnb3J5OiAnYXR0cmlidXRlcycsXG4gICAgICAgICAgICB9KTtcblxuICAgICAgICAgICAgZm9yIChjb25zdCBhdHRyIG9mIGRlY2wuYXR0cmlidXRlcykge1xuICAgICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICAgICAgdHlwZTogJ2F0dHJpYnV0ZScsXG4gICAgICAgICAgICAgICAgbGFiZWw6IGF0dHIubmFtZSxcbiAgICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAzLFxuICAgICAgICAgICAgICAgIHBhcmVudElkOiBhdHRyQ2F0SWQsXG4gICAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgICAgICB0YWdOYW1lOiBkZWNsLnRhZ05hbWUsXG4gICAgICAgICAgICAgICAgbmFtZTogYXR0ci5uYW1lLFxuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG5cbiAgICAgICAgICAvLyBQcm9wZXJ0aWVzXG4gICAgICAgICAgaWYgKHByb3BlcnRpZXMubGVuZ3RoKSB7XG4gICAgICAgICAgICBjb25zdCBwcm9wQ2F0SWQgPSBpZCsrO1xuICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICBpZDogcHJvcENhdElkLFxuICAgICAgICAgICAgICB0eXBlOiAnY2F0ZWdvcnknLFxuICAgICAgICAgICAgICBsYWJlbDogJ1Byb3BlcnRpZXMnLFxuICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAyLFxuICAgICAgICAgICAgICBwYXJlbnRJZDogY2VJZCxcbiAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IHRydWUsXG4gICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgIGJhZGdlOiBwcm9wZXJ0aWVzLmxlbmd0aCxcbiAgICAgICAgICAgICAgY2F0ZWdvcnk6ICdwcm9wZXJ0aWVzJyxcbiAgICAgICAgICAgIH0pO1xuXG4gICAgICAgICAgICBmb3IgKGNvbnN0IHByb3Agb2YgcHJvcGVydGllcykge1xuICAgICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICAgICAgdHlwZTogJ3Byb3BlcnR5JyxcbiAgICAgICAgICAgICAgICBsYWJlbDogcHJvcC5uYW1lLFxuICAgICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDMsXG4gICAgICAgICAgICAgICAgcGFyZW50SWQ6IHByb3BDYXRJZCxcbiAgICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgICAgIHRhZ05hbWU6IGRlY2wudGFnTmFtZSxcbiAgICAgICAgICAgICAgICBuYW1lOiBwcm9wLm5hbWUsXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cblxuICAgICAgICAgIC8vIE1ldGhvZHNcbiAgICAgICAgICBpZiAobWV0aG9kcy5sZW5ndGgpIHtcbiAgICAgICAgICAgIGNvbnN0IG1ldGhvZENhdElkID0gaWQrKztcbiAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgaWQ6IG1ldGhvZENhdElkLFxuICAgICAgICAgICAgICB0eXBlOiAnY2F0ZWdvcnknLFxuICAgICAgICAgICAgICBsYWJlbDogJ01ldGhvZHMnLFxuICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAyLFxuICAgICAgICAgICAgICBwYXJlbnRJZDogY2VJZCxcbiAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IHRydWUsXG4gICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgIGJhZGdlOiBtZXRob2RzLmxlbmd0aCxcbiAgICAgICAgICAgICAgY2F0ZWdvcnk6ICdtZXRob2RzJyxcbiAgICAgICAgICAgIH0pO1xuXG4gICAgICAgICAgICBmb3IgKGNvbnN0IG1ldGhvZCBvZiBtZXRob2RzKSB7XG4gICAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgICAgICB0eXBlOiAnbWV0aG9kJyxcbiAgICAgICAgICAgICAgICBsYWJlbDogYCR7bWV0aG9kLm5hbWV9KClgLFxuICAgICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDMsXG4gICAgICAgICAgICAgICAgcGFyZW50SWQ6IG1ldGhvZENhdElkLFxuICAgICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgICAgbW9kdWxlUGF0aDogbW9kdWxlLnBhdGgsXG4gICAgICAgICAgICAgICAgdGFnTmFtZTogZGVjbC50YWdOYW1lLFxuICAgICAgICAgICAgICAgIG5hbWU6IG1ldGhvZC5uYW1lLFxuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG5cbiAgICAgICAgICAvLyBFdmVudHNcbiAgICAgICAgICBpZiAoZGVjbC5ldmVudHM/Lmxlbmd0aCkge1xuICAgICAgICAgICAgY29uc3QgZXZlbnRDYXRJZCA9IGlkKys7XG4gICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgIGlkOiBldmVudENhdElkLFxuICAgICAgICAgICAgICB0eXBlOiAnY2F0ZWdvcnknLFxuICAgICAgICAgICAgICBsYWJlbDogJ0V2ZW50cycsXG4gICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDIsXG4gICAgICAgICAgICAgIHBhcmVudElkOiBjZUlkLFxuICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogdHJ1ZSxcbiAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgYmFkZ2U6IGRlY2wuZXZlbnRzLmxlbmd0aCxcbiAgICAgICAgICAgICAgY2F0ZWdvcnk6ICdldmVudHMnLFxuICAgICAgICAgICAgfSk7XG5cbiAgICAgICAgICAgIGZvciAoY29uc3QgZXZlbnQgb2YgZGVjbC5ldmVudHMpIHtcbiAgICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICAgIGlkOiBpZCsrLFxuICAgICAgICAgICAgICAgIHR5cGU6ICdldmVudCcsXG4gICAgICAgICAgICAgICAgbGFiZWw6IGV2ZW50Lm5hbWUsXG4gICAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMyxcbiAgICAgICAgICAgICAgICBwYXJlbnRJZDogZXZlbnRDYXRJZCxcbiAgICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgICAgIHRhZ05hbWU6IGRlY2wudGFnTmFtZSxcbiAgICAgICAgICAgICAgICBuYW1lOiBldmVudC5uYW1lLFxuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG5cbiAgICAgICAgICAvLyBTbG90c1xuICAgICAgICAgIGlmIChkZWNsLnNsb3RzPy5sZW5ndGgpIHtcbiAgICAgICAgICAgIGNvbnN0IHNsb3RDYXRJZCA9IGlkKys7XG4gICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgIGlkOiBzbG90Q2F0SWQsXG4gICAgICAgICAgICAgIHR5cGU6ICdjYXRlZ29yeScsXG4gICAgICAgICAgICAgIGxhYmVsOiAnU2xvdHMnLFxuICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAyLFxuICAgICAgICAgICAgICBwYXJlbnRJZDogY2VJZCxcbiAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IHRydWUsXG4gICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICAgIGJhZGdlOiBkZWNsLnNsb3RzLmxlbmd0aCxcbiAgICAgICAgICAgICAgY2F0ZWdvcnk6ICdzbG90cycsXG4gICAgICAgICAgICB9KTtcblxuICAgICAgICAgICAgZm9yIChjb25zdCBzbG90IG9mIGRlY2wuc2xvdHMpIHtcbiAgICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICAgIGlkOiBpZCsrLFxuICAgICAgICAgICAgICAgIHR5cGU6ICdzbG90JyxcbiAgICAgICAgICAgICAgICBsYWJlbDogc2xvdC5uYW1lIHx8ICcoZGVmYXVsdCknLFxuICAgICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDMsXG4gICAgICAgICAgICAgICAgcGFyZW50SWQ6IHNsb3RDYXRJZCxcbiAgICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogZmFsc2UsXG4gICAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICAgIG1vZHVsZVBhdGg6IG1vZHVsZS5wYXRoLFxuICAgICAgICAgICAgICAgIHRhZ05hbWU6IGRlY2wudGFnTmFtZSxcbiAgICAgICAgICAgICAgICBuYW1lOiBzbG90Lm5hbWUsXG4gICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cblxuICAgICAgICAgIC8vIENTUyBQcm9wZXJ0aWVzXG4gICAgICAgICAgaWYgKGRlY2wuY3NzUHJvcGVydGllcz8ubGVuZ3RoKSB7XG4gICAgICAgICAgICBjb25zdCBjc3NQcm9wQ2F0SWQgPSBpZCsrO1xuICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICBpZDogY3NzUHJvcENhdElkLFxuICAgICAgICAgICAgICB0eXBlOiAnY2F0ZWdvcnknLFxuICAgICAgICAgICAgICBsYWJlbDogJ0NTUyBQcm9wZXJ0aWVzJyxcbiAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMixcbiAgICAgICAgICAgICAgcGFyZW50SWQ6IGNlSWQsXG4gICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiB0cnVlLFxuICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICBiYWRnZTogZGVjbC5jc3NQcm9wZXJ0aWVzLmxlbmd0aCxcbiAgICAgICAgICAgICAgY2F0ZWdvcnk6ICdjc3MtcHJvcGVydGllcycsXG4gICAgICAgICAgICB9KTtcblxuICAgICAgICAgICAgZm9yIChjb25zdCBjc3NQcm9wIG9mIGRlY2wuY3NzUHJvcGVydGllcykge1xuICAgICAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICAgICAgdHlwZTogJ2Nzcy1wcm9wZXJ0eScsXG4gICAgICAgICAgICAgICAgbGFiZWw6IGNzc1Byb3AubmFtZSxcbiAgICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAzLFxuICAgICAgICAgICAgICAgIHBhcmVudElkOiBjc3NQcm9wQ2F0SWQsXG4gICAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgICAgICB0YWdOYW1lOiBkZWNsLnRhZ05hbWUsXG4gICAgICAgICAgICAgICAgbmFtZTogY3NzUHJvcC5uYW1lLFxuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG5cbiAgICAgICAgICAvLyBDU1MgUGFydHNcbiAgICAgICAgICBpZiAoZGVjbC5jc3NQYXJ0cz8ubGVuZ3RoKSB7XG4gICAgICAgICAgICBjb25zdCBjc3NQYXJ0Q2F0SWQgPSBpZCsrO1xuICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICBpZDogY3NzUGFydENhdElkLFxuICAgICAgICAgICAgICB0eXBlOiAnY2F0ZWdvcnknLFxuICAgICAgICAgICAgICBsYWJlbDogJ0NTUyBQYXJ0cycsXG4gICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDIsXG4gICAgICAgICAgICAgIHBhcmVudElkOiBjZUlkLFxuICAgICAgICAgICAgICBoYXNDaGlsZHJlbjogdHJ1ZSxcbiAgICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgYmFkZ2U6IGRlY2wuY3NzUGFydHMubGVuZ3RoLFxuICAgICAgICAgICAgICBjYXRlZ29yeTogJ2Nzcy1wYXJ0cycsXG4gICAgICAgICAgICB9KTtcblxuICAgICAgICAgICAgZm9yIChjb25zdCBjc3NQYXJ0IG9mIGRlY2wuY3NzUGFydHMpIHtcbiAgICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICAgIGlkOiBpZCsrLFxuICAgICAgICAgICAgICAgIHR5cGU6ICdjc3MtcGFydCcsXG4gICAgICAgICAgICAgICAgbGFiZWw6IGNzc1BhcnQubmFtZSxcbiAgICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAzLFxuICAgICAgICAgICAgICAgIHBhcmVudElkOiBjc3NQYXJ0Q2F0SWQsXG4gICAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgICAgICB0YWdOYW1lOiBkZWNsLnRhZ05hbWUsXG4gICAgICAgICAgICAgICAgbmFtZTogY3NzUGFydC5uYW1lLFxuICAgICAgICAgICAgICB9KTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG5cbiAgICAgICAgICAvLyBDU1MgU3RhdGVzXG4gICAgICAgICAgaWYgKGRlY2wuY3NzU3RhdGVzPy5sZW5ndGgpIHtcbiAgICAgICAgICAgIGNvbnN0IGNzc1N0YXRlQ2F0SWQgPSBpZCsrO1xuICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICBpZDogY3NzU3RhdGVDYXRJZCxcbiAgICAgICAgICAgICAgdHlwZTogJ2NhdGVnb3J5JyxcbiAgICAgICAgICAgICAgbGFiZWw6ICdDU1MgU3RhdGVzJyxcbiAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMixcbiAgICAgICAgICAgICAgcGFyZW50SWQ6IGNlSWQsXG4gICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiB0cnVlLFxuICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICBiYWRnZTogZGVjbC5jc3NTdGF0ZXMubGVuZ3RoLFxuICAgICAgICAgICAgICBjYXRlZ29yeTogJ2Nzcy1zdGF0ZXMnLFxuICAgICAgICAgICAgfSk7XG5cbiAgICAgICAgICAgIGZvciAoY29uc3QgY3NzU3RhdGUgb2YgZGVjbC5jc3NTdGF0ZXMpIHtcbiAgICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICAgIGlkOiBpZCsrLFxuICAgICAgICAgICAgICAgIHR5cGU6ICdjc3Mtc3RhdGUnLFxuICAgICAgICAgICAgICAgIGxhYmVsOiBjc3NTdGF0ZS5uYW1lLFxuICAgICAgICAgICAgICAgIGRlcHRoOiBkZXB0aCArIDMsXG4gICAgICAgICAgICAgICAgcGFyZW50SWQ6IGNzc1N0YXRlQ2F0SWQsXG4gICAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgICAgICB0YWdOYW1lOiBkZWNsLnRhZ05hbWUsXG4gICAgICAgICAgICAgICAgbmFtZTogY3NzU3RhdGUubmFtZSxcbiAgICAgICAgICAgICAgfSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgLy8gRGVtb3NcbiAgICAgICAgICBpZiAoZGVjbC5kZW1vcz8ubGVuZ3RoKSB7XG4gICAgICAgICAgICBjb25zdCBkZW1vQ2F0SWQgPSBpZCsrO1xuICAgICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgICBpZDogZGVtb0NhdElkLFxuICAgICAgICAgICAgICB0eXBlOiAnY2F0ZWdvcnknLFxuICAgICAgICAgICAgICBsYWJlbDogJ0RlbW9zJyxcbiAgICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMixcbiAgICAgICAgICAgICAgcGFyZW50SWQ6IGNlSWQsXG4gICAgICAgICAgICAgIGhhc0NoaWxkcmVuOiB0cnVlLFxuICAgICAgICAgICAgICBleHBhbmRlZDogZmFsc2UsXG4gICAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgICBiYWRnZTogZGVjbC5kZW1vcy5sZW5ndGgsXG4gICAgICAgICAgICAgIGNhdGVnb3J5OiAnZGVtb3MnLFxuICAgICAgICAgICAgfSk7XG5cbiAgICAgICAgICAgIGZvciAoY29uc3QgZGVtbyBvZiBkZWNsLmRlbW9zKSB7XG4gICAgICAgICAgICAgIHRoaXMuI2ZsYXRJdGVtcy5wdXNoKHtcbiAgICAgICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgICAgICB0eXBlOiAnZGVtbycsXG4gICAgICAgICAgICAgICAgbGFiZWw6IGRlbW8ubmFtZSB8fCBkZW1vLnVybCB8fCAnKGRlbW8pJyxcbiAgICAgICAgICAgICAgICBkZXB0aDogZGVwdGggKyAzLFxuICAgICAgICAgICAgICAgIHBhcmVudElkOiBkZW1vQ2F0SWQsXG4gICAgICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgICAgICB2aXNpYmxlOiBmYWxzZSxcbiAgICAgICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgICAgICB0YWdOYW1lOiBkZWNsLnRhZ05hbWUsXG4gICAgICAgICAgICAgICAgdXJsOiBkZW1vLnVybCxcbiAgICAgICAgICAgICAgfSk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuICAgICAgICB9IGVsc2UgaWYgKGRlY2wua2luZCA9PT0gJ2NsYXNzJykge1xuICAgICAgICAgIC8vIFJlZ3VsYXIgY2xhc3NcbiAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgIHR5cGU6ICdjbGFzcycsXG4gICAgICAgICAgICBsYWJlbDogZGVjbC5uYW1lLFxuICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMSxcbiAgICAgICAgICAgIHBhcmVudElkOiBtb2R1bGVJZCxcbiAgICAgICAgICAgIGhhc0NoaWxkcmVuOiBmYWxzZSxcbiAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgbW9kdWxlUGF0aDogbW9kdWxlLnBhdGgsXG4gICAgICAgICAgICBuYW1lOiBkZWNsLm5hbWUsXG4gICAgICAgICAgfSk7XG4gICAgICAgIH0gZWxzZSBpZiAoZGVjbC5raW5kID09PSAnZnVuY3Rpb24nKSB7XG4gICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICB0eXBlOiAnZnVuY3Rpb24nLFxuICAgICAgICAgICAgbGFiZWw6IGAke2RlY2wubmFtZX0oKWAsXG4gICAgICAgICAgICBkZXB0aDogZGVwdGggKyAxLFxuICAgICAgICAgICAgcGFyZW50SWQ6IG1vZHVsZUlkLFxuICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgIG5hbWU6IGRlY2wubmFtZSxcbiAgICAgICAgICB9KTtcbiAgICAgICAgfSBlbHNlIGlmIChkZWNsLmtpbmQgPT09ICd2YXJpYWJsZScpIHtcbiAgICAgICAgICB0aGlzLiNmbGF0SXRlbXMucHVzaCh7XG4gICAgICAgICAgICBpZDogaWQrKyxcbiAgICAgICAgICAgIHR5cGU6ICd2YXJpYWJsZScsXG4gICAgICAgICAgICBsYWJlbDogZGVjbC5uYW1lLFxuICAgICAgICAgICAgZGVwdGg6IGRlcHRoICsgMSxcbiAgICAgICAgICAgIHBhcmVudElkOiBtb2R1bGVJZCxcbiAgICAgICAgICAgIGhhc0NoaWxkcmVuOiBmYWxzZSxcbiAgICAgICAgICAgIGV4cGFuZGVkOiBmYWxzZSxcbiAgICAgICAgICAgIHZpc2libGU6IGZhbHNlLFxuICAgICAgICAgICAgbW9kdWxlUGF0aDogbW9kdWxlLnBhdGgsXG4gICAgICAgICAgICBuYW1lOiBkZWNsLm5hbWUsXG4gICAgICAgICAgfSk7XG4gICAgICAgIH0gZWxzZSBpZiAoZGVjbC5raW5kID09PSAnbWl4aW4nKSB7XG4gICAgICAgICAgdGhpcy4jZmxhdEl0ZW1zLnB1c2goe1xuICAgICAgICAgICAgaWQ6IGlkKyssXG4gICAgICAgICAgICB0eXBlOiAnbWl4aW4nLFxuICAgICAgICAgICAgbGFiZWw6IGAke2RlY2wubmFtZX0oKWAsXG4gICAgICAgICAgICBkZXB0aDogZGVwdGggKyAxLFxuICAgICAgICAgICAgcGFyZW50SWQ6IG1vZHVsZUlkLFxuICAgICAgICAgICAgaGFzQ2hpbGRyZW46IGZhbHNlLFxuICAgICAgICAgICAgZXhwYW5kZWQ6IGZhbHNlLFxuICAgICAgICAgICAgdmlzaWJsZTogZmFsc2UsXG4gICAgICAgICAgICBtb2R1bGVQYXRoOiBtb2R1bGUucGF0aCxcbiAgICAgICAgICAgIG5hbWU6IGRlY2wubmFtZSxcbiAgICAgICAgICB9KTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiBpZDtcbiAgfVxuXG4gIC8qKlxuICAgKiBVcGRhdGUgdmlzaWJsZSBpdGVtcyBiYXNlZCBvbiBleHBhbmRlZCBzdGF0ZVxuICAgKi9cbiAgI3VwZGF0ZVZpc2libGVJdGVtcygpIHtcbiAgICB0aGlzLiN2aXNpYmxlSXRlbXMgPSB0aGlzLiNmbGF0SXRlbXMuZmlsdGVyKGl0ZW0gPT4gaXRlbS52aXNpYmxlKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBSZW5kZXIgdmlzaWJsZSB0cmVlIGl0ZW1zIHdpdGggcHJvcGVyIG5lc3RpbmdcbiAgICovXG4gICNyZW5kZXJUcmVlKCkge1xuICAgIGlmICghdGhpcy4jdmlld3BvcnQpIHJldHVybjtcblxuICAgIC8vIENsZWFyIHZpZXdwb3J0XG4gICAgdGhpcy4jdmlld3BvcnQuaW5uZXJIVE1MID0gJyc7XG5cbiAgICAvLyBCdWlsZCBuZXN0ZWQgdHJlZSBzdHJ1Y3R1cmUgZnJvbSBmbGF0IHZpc2libGUgaXRlbXNcbiAgICBjb25zdCByb290SXRlbXMgPSB0aGlzLiN2aXNpYmxlSXRlbXMuZmlsdGVyKGl0ZW0gPT4gaXRlbS5kZXB0aCA9PT0gMCk7XG5cbiAgICBmb3IgKGNvbnN0IHJvb3RJdGVtIG9mIHJvb3RJdGVtcykge1xuICAgICAgY29uc3QgdHJlZUl0ZW1FbCA9IHRoaXMuI2NyZWF0ZVRyZWVJdGVtV2l0aENoaWxkcmVuKHJvb3RJdGVtKTtcbiAgICAgIGlmICh0cmVlSXRlbUVsKSB7XG4gICAgICAgIHRoaXMuI3ZpZXdwb3J0LmFwcGVuZENoaWxkKHRyZWVJdGVtRWwpO1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBDcmVhdGUgYSB0cmVlIGl0ZW0gZWxlbWVudCB3aXRoIGl0cyBuZXN0ZWQgY2hpbGRyZW5cbiAgICovXG4gICNjcmVhdGVUcmVlSXRlbVdpdGhDaGlsZHJlbihpdGVtOiBUcmVlSXRlbSk6IEhUTUxFbGVtZW50IHtcbiAgICBjb25zdCBlbCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3BmLXY2LXRyZWUtaXRlbScpO1xuICAgIGVsLnNldEF0dHJpYnV0ZSgnbGFiZWwnLCBpdGVtLmxhYmVsKTtcblxuICAgIGlmIChpdGVtLmJhZGdlKSB7XG4gICAgICBlbC5zZXRBdHRyaWJ1dGUoJ2JhZGdlJywgU3RyaW5nKGl0ZW0uYmFkZ2UpKTtcbiAgICB9XG5cbiAgICAvLyBTdG9yZSBpdGVtIGRhdGFcbiAgICBlbC5kYXRhc2V0Lml0ZW1JZCA9IFN0cmluZyhpdGVtLmlkKTtcbiAgICBlbC5kYXRhc2V0LnR5cGUgPSBpdGVtLnR5cGU7XG5cbiAgICAvLyBSZXN0b3JlIHNlbGVjdGlvbiBzdGF0ZSBpZiB0aGlzIGl0ZW0gd2FzIHNlbGVjdGVkXG4gICAgaWYgKHRoaXMuI2N1cnJlbnRTZWxlY3RlZEl0ZW1JZCA9PT0gaXRlbS5pZCkge1xuICAgICAgZWwuc2V0QXR0cmlidXRlKCdjdXJyZW50JywgJycpO1xuICAgICAgdGhpcy4jY3VycmVudFNlbGVjdGVkRWxlbWVudCA9IGVsO1xuICAgIH1cblxuICAgIC8vIFNldCBoYXMtY2hpbGRyZW4gYXR0cmlidXRlIGZvciBpdGVtcyB3aXRoIGNoaWxkcmVuIChzaG93cyB0b2dnbGUgYnV0dG9uKVxuICAgIGlmIChpdGVtLmhhc0NoaWxkcmVuKSB7XG4gICAgICBlbC5zZXRBdHRyaWJ1dGUoJ2hhcy1jaGlsZHJlbicsICcnKTtcbiAgICAgIChlbCBhcyBhbnkpLmV4cGFuZGVkID0gaXRlbS5leHBhbmRlZDtcblxuICAgICAgLy8gTGlzdGVuIGZvciBleHBhbmQvY29sbGFwc2VcbiAgICAgIC8vIFN0b3AgcHJvcGFnYXRpb24gdG8gcHJldmVudCBuZXN0ZWQgY29sbGFwc2VzIGZyb20gYWZmZWN0aW5nIHBhcmVudCBpdGVtc1xuICAgICAgZWwuYWRkRXZlbnRMaXN0ZW5lcignZXhwYW5kJywgKGU6IEV2ZW50KSA9PiB7XG4gICAgICAgIGUuc3RvcFByb3BhZ2F0aW9uKCk7XG4gICAgICAgIHRoaXMuI2hhbmRsZVRvZ2dsZShpdGVtLCB0cnVlKTtcbiAgICAgIH0pO1xuXG4gICAgICBlbC5hZGRFdmVudExpc3RlbmVyKCdjb2xsYXBzZScsIChlOiBFdmVudCkgPT4ge1xuICAgICAgICBlLnN0b3BQcm9wYWdhdGlvbigpO1xuICAgICAgICB0aGlzLiNoYW5kbGVUb2dnbGUoaXRlbSwgZmFsc2UpO1xuICAgICAgfSk7XG4gICAgfVxuXG4gICAgLy8gTGlzdGVuIGZvciBzZWxlY3Rpb24gKG5vbi1jYXRlZ29yeSBpdGVtcylcbiAgICBpZiAoaXRlbS50eXBlICE9PSAnY2F0ZWdvcnknKSB7XG4gICAgICBlbC5hZGRFdmVudExpc3RlbmVyKCdzZWxlY3QnLCAoZTogRXZlbnQpID0+IHtcbiAgICAgICAgZS5zdG9wUHJvcGFnYXRpb24oKTtcbiAgICAgICAgdGhpcy4jaGFuZGxlU2VsZWN0KGl0ZW0sIGVsKTtcbiAgICAgIH0pO1xuICAgIH1cblxuICAgIC8vIEFkZCBjaGlsZHJlbiBpZiBleHBhbmRlZCBhbmQgdmlzaWJsZVxuICAgIGlmIChpdGVtLmV4cGFuZGVkICYmIGl0ZW0uaGFzQ2hpbGRyZW4pIHtcbiAgICAgIGNvbnN0IGNoaWxkcmVuID0gdGhpcy4jdmlzaWJsZUl0ZW1zLmZpbHRlcihjaGlsZCA9PiBjaGlsZC5wYXJlbnRJZCA9PT0gaXRlbS5pZCAmJiBjaGlsZC52aXNpYmxlKTtcbiAgICAgIGZvciAoY29uc3QgY2hpbGQgb2YgY2hpbGRyZW4pIHtcbiAgICAgICAgY29uc3QgY2hpbGRFbCA9IHRoaXMuI2NyZWF0ZVRyZWVJdGVtV2l0aENoaWxkcmVuKGNoaWxkKTtcbiAgICAgICAgaWYgKGNoaWxkRWwpIHtcbiAgICAgICAgICBlbC5hcHBlbmRDaGlsZChjaGlsZEVsKTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIHJldHVybiBlbDtcbiAgfVxuXG4gIC8qKlxuICAgKiBIYW5kbGUgdHJlZSBpdGVtIHRvZ2dsZSAoZXhwYW5kL2NvbGxhcHNlKVxuICAgKi9cbiAgI2hhbmRsZVRvZ2dsZShpdGVtOiBUcmVlSXRlbSwgZXhwYW5kZWQ6IGJvb2xlYW4pIHtcbiAgICBpdGVtLmV4cGFuZGVkID0gZXhwYW5kZWQ7XG4gICAgdGhpcy4jdXBkYXRlQ2hpbGRWaXNpYmlsaXR5KGl0ZW0pO1xuICAgIHRoaXMuI3VwZGF0ZVZpc2libGVJdGVtcygpO1xuICAgIHRoaXMuI3JlbmRlclRyZWUoKTtcbiAgfVxuXG4gIC8qKlxuICAgKiBVcGRhdGUgdmlzaWJpbGl0eSBvZiBjaGlsZHJlbiB3aGVuIHBhcmVudCBpcyB0b2dnbGVkXG4gICAqL1xuICAjdXBkYXRlQ2hpbGRWaXNpYmlsaXR5KHBhcmVudEl0ZW06IFRyZWVJdGVtKSB7XG4gICAgY29uc3QgaXNFeHBhbmRlZCA9IHBhcmVudEl0ZW0uZXhwYW5kZWQ7XG5cbiAgICBmb3IgKGNvbnN0IGl0ZW0gb2YgdGhpcy4jZmxhdEl0ZW1zKSB7XG4gICAgICBpZiAoaXRlbS5wYXJlbnRJZCA9PT0gcGFyZW50SXRlbS5pZCkge1xuICAgICAgICBpdGVtLnZpc2libGUgPSBpc0V4cGFuZGVkO1xuICAgICAgICAvLyBJZiBjb2xsYXBzaW5nLCBhbHNvIGNvbGxhcHNlIGFuZCBoaWRlIGFsbCBkZXNjZW5kYW50c1xuICAgICAgICBpZiAoIWlzRXhwYW5kZWQpIHtcbiAgICAgICAgICBpdGVtLmV4cGFuZGVkID0gZmFsc2U7XG4gICAgICAgICAgdGhpcy4jdXBkYXRlQ2hpbGRWaXNpYmlsaXR5KGl0ZW0pO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEhhbmRsZSB0cmVlIGl0ZW0gc2VsZWN0aW9uXG4gICAqL1xuICAjaGFuZGxlU2VsZWN0KGl0ZW06IFRyZWVJdGVtLCBlbGVtZW50OiBFbGVtZW50KSB7XG4gICAgLy8gQ2xlYXIgcHJldmlvdXMgc2VsZWN0aW9uXG4gICAgaWYgKHRoaXMuI2N1cnJlbnRTZWxlY3RlZEVsZW1lbnQpIHtcbiAgICAgIHRoaXMuI2N1cnJlbnRTZWxlY3RlZEVsZW1lbnQucmVtb3ZlQXR0cmlidXRlKCdjdXJyZW50Jyk7XG4gICAgfVxuXG4gICAgLy8gU2V0IG5ldyBzZWxlY3Rpb25cbiAgICBlbGVtZW50LnNldEF0dHJpYnV0ZSgnY3VycmVudCcsICcnKTtcbiAgICB0aGlzLiNjdXJyZW50U2VsZWN0ZWRFbGVtZW50ID0gZWxlbWVudDtcbiAgICB0aGlzLiNjdXJyZW50U2VsZWN0ZWRJdGVtSWQgPSBpdGVtLmlkO1xuXG4gICAgLy8gRGlzcGF0Y2ggY3VzdG9tIGV2ZW50XG4gICAgdGhpcy5kaXNwYXRjaEV2ZW50KG5ldyBJdGVtU2VsZWN0RXZlbnQoaXRlbSkpO1xuICB9XG5cbiAgLyoqXG4gICAqIFNlYXJjaC9maWx0ZXIgdHJlZSBpdGVtc1xuICAgKi9cbiAgc2VhcmNoKHF1ZXJ5OiBzdHJpbmcpIHtcbiAgICB0aGlzLiNzZWFyY2hRdWVyeSA9IHF1ZXJ5LnRvTG93ZXJDYXNlKCkudHJpbSgpO1xuXG4gICAgaWYgKCF0aGlzLiNzZWFyY2hRdWVyeSkge1xuICAgICAgLy8gUmVzZXQgdG8gbm9ybWFsIHZpc2liaWxpdHlcbiAgICAgIGZvciAoY29uc3QgaXRlbSBvZiB0aGlzLiNmbGF0SXRlbXMpIHtcbiAgICAgICAgaXRlbS52aXNpYmxlID0gaXRlbS5kZXB0aCA9PT0gMCB8fCAoaXRlbS5wYXJlbnRJZCAhPT0gdW5kZWZpbmVkICYmIHRoaXMuI2lzUGFyZW50RXhwYW5kZWQoaXRlbSkpO1xuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICAvLyBGaWx0ZXIgYmFzZWQgb24gc2VhcmNoXG4gICAgICBjb25zdCBtYXRjaGluZ0l0ZW1zID0gbmV3IFNldDxUcmVlSXRlbT4oKTtcblxuICAgICAgLy8gRmluZCBtYXRjaGluZyBpdGVtc1xuICAgICAgZm9yIChjb25zdCBpdGVtIG9mIHRoaXMuI2ZsYXRJdGVtcykge1xuICAgICAgICBpZiAoaXRlbS5sYWJlbC50b0xvd2VyQ2FzZSgpLmluY2x1ZGVzKHRoaXMuI3NlYXJjaFF1ZXJ5KSkge1xuICAgICAgICAgIG1hdGNoaW5nSXRlbXMuYWRkKGl0ZW0pO1xuICAgICAgICAgIC8vIFNob3cgYW5kIGV4cGFuZCBhbGwgYW5jZXN0b3JzXG4gICAgICAgICAgdGhpcy4jc2hvd0FuY2VzdG9ycyhpdGVtKTtcbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICAvLyBIaWRlIG5vbi1tYXRjaGluZyBpdGVtc1xuICAgICAgZm9yIChjb25zdCBpdGVtIG9mIHRoaXMuI2ZsYXRJdGVtcykge1xuICAgICAgICBpZiAoIW1hdGNoaW5nSXRlbXMuaGFzKGl0ZW0pICYmICF0aGlzLiNoYXNNYXRjaGluZ0Rlc2NlbmRhbnQoaXRlbSwgbWF0Y2hpbmdJdGVtcykpIHtcbiAgICAgICAgICBpdGVtLnZpc2libGUgPSB0aGlzLiNpc0FuY2VzdG9yT2ZNYXRjaChpdGVtLCBtYXRjaGluZ0l0ZW1zKTtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBpdGVtLnZpc2libGUgPSB0cnVlO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgdGhpcy4jdXBkYXRlVmlzaWJsZUl0ZW1zKCk7XG4gICAgdGhpcy4jcmVuZGVyVHJlZSgpO1xuICB9XG5cbiAgLyoqXG4gICAqIENoZWNrIGlmIHBhcmVudCBpcyBleHBhbmRlZFxuICAgKi9cbiAgI2lzUGFyZW50RXhwYW5kZWQoaXRlbTogVHJlZUl0ZW0pOiBib29sZWFuIHtcbiAgICBpZiAoaXRlbS5wYXJlbnRJZCA9PT0gdW5kZWZpbmVkKSByZXR1cm4gdHJ1ZTtcbiAgICBjb25zdCBwYXJlbnQgPSB0aGlzLiNmbGF0SXRlbXMuZmluZChpID0+IGkuaWQgPT09IGl0ZW0ucGFyZW50SWQpO1xuICAgIHJldHVybiAhIXBhcmVudCAmJiBwYXJlbnQuZXhwYW5kZWQgJiYgdGhpcy4jaXNQYXJlbnRFeHBhbmRlZChwYXJlbnQpO1xuICB9XG5cbiAgLyoqXG4gICAqIFNob3cgYWxsIGFuY2VzdG9ycyBvZiBhbiBpdGVtXG4gICAqL1xuICAjc2hvd0FuY2VzdG9ycyhpdGVtOiBUcmVlSXRlbSkge1xuICAgIGlmIChpdGVtLnBhcmVudElkID09PSB1bmRlZmluZWQpIHJldHVybjtcbiAgICBjb25zdCBwYXJlbnQgPSB0aGlzLiNmbGF0SXRlbXMuZmluZChpID0+IGkuaWQgPT09IGl0ZW0ucGFyZW50SWQpO1xuICAgIGlmIChwYXJlbnQpIHtcbiAgICAgIHBhcmVudC52aXNpYmxlID0gdHJ1ZTtcbiAgICAgIHBhcmVudC5leHBhbmRlZCA9IHRydWU7XG4gICAgICB0aGlzLiNzaG93QW5jZXN0b3JzKHBhcmVudCk7XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIENoZWNrIGlmIGl0ZW0gaXMgYW5jZXN0b3Igb2YgYSBtYXRjaFxuICAgKi9cbiAgI2lzQW5jZXN0b3JPZk1hdGNoKGl0ZW06IFRyZWVJdGVtLCBtYXRjaGluZ0l0ZW1zOiBTZXQ8VHJlZUl0ZW0+KTogYm9vbGVhbiB7XG4gICAgZm9yIChjb25zdCBtYXRjaCBvZiBtYXRjaGluZ0l0ZW1zKSB7XG4gICAgICBsZXQgY3VycmVudDogVHJlZUl0ZW0gPSBtYXRjaDtcbiAgICAgIHdoaWxlIChjdXJyZW50LnBhcmVudElkICE9PSB1bmRlZmluZWQpIHtcbiAgICAgICAgY29uc3QgcGFyZW50ID0gdGhpcy4jZmxhdEl0ZW1zLmZpbmQoaSA9PiBpLmlkID09PSBjdXJyZW50LnBhcmVudElkKTtcbiAgICAgICAgaWYgKCFwYXJlbnQpIGJyZWFrO1xuICAgICAgICBpZiAocGFyZW50ID09PSBpdGVtKSByZXR1cm4gdHJ1ZTtcbiAgICAgICAgY3VycmVudCA9IHBhcmVudDtcbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuIGZhbHNlO1xuICB9XG5cbiAgLyoqXG4gICAqIENoZWNrIGlmIGl0ZW0gaGFzIG1hdGNoaW5nIGRlc2NlbmRhbnRcbiAgICovXG4gICNoYXNNYXRjaGluZ0Rlc2NlbmRhbnQoaXRlbTogVHJlZUl0ZW0sIG1hdGNoaW5nSXRlbXM6IFNldDxUcmVlSXRlbT4pOiBib29sZWFuIHtcbiAgICBmb3IgKGNvbnN0IG1hdGNoIG9mIG1hdGNoaW5nSXRlbXMpIHtcbiAgICAgIGlmICh0aGlzLiNpc0Rlc2NlbmRhbnRPZihtYXRjaCwgaXRlbSkpIHJldHVybiB0cnVlO1xuICAgIH1cbiAgICByZXR1cm4gZmFsc2U7XG4gIH1cblxuICAvKipcbiAgICogQ2hlY2sgaWYgaXRlbSBpcyBkZXNjZW5kYW50IG9mIGFub3RoZXJcbiAgICovXG4gICNpc0Rlc2NlbmRhbnRPZihpdGVtOiBUcmVlSXRlbSwgYW5jZXN0b3I6IFRyZWVJdGVtKTogYm9vbGVhbiB7XG4gICAgbGV0IGN1cnJlbnQ6IFRyZWVJdGVtID0gaXRlbTtcbiAgICB3aGlsZSAoY3VycmVudC5wYXJlbnRJZCAhPT0gdW5kZWZpbmVkKSB7XG4gICAgICBjb25zdCBwYXJlbnQgPSB0aGlzLiNmbGF0SXRlbXMuZmluZChpID0+IGkuaWQgPT09IGN1cnJlbnQucGFyZW50SWQpO1xuICAgICAgaWYgKCFwYXJlbnQpIHJldHVybiBmYWxzZTtcbiAgICAgIGlmIChwYXJlbnQgPT09IGFuY2VzdG9yKSByZXR1cm4gdHJ1ZTtcbiAgICAgIGN1cnJlbnQgPSBwYXJlbnQ7XG4gICAgfVxuICAgIHJldHVybiBmYWxzZTtcbiAgfVxuXG4gIC8qKlxuICAgKiBFeHBhbmQgYWxsIGl0ZW1zXG4gICAqL1xuICBleHBhbmRBbGwoKSB7XG4gICAgZm9yIChjb25zdCBpdGVtIG9mIHRoaXMuI2ZsYXRJdGVtcykge1xuICAgICAgaWYgKGl0ZW0uaGFzQ2hpbGRyZW4pIHtcbiAgICAgICAgaXRlbS5leHBhbmRlZCA9IHRydWU7XG4gICAgICB9XG4gICAgICAvLyBNYWtlIGFsbCBpdGVtcyB2aXNpYmxlIGV4Y2VwdCB0aG9zZSBmaWx0ZXJlZCBieSBzZWFyY2hcbiAgICAgIGlmICghdGhpcy4jc2VhcmNoUXVlcnkgfHwgaXRlbS5sYWJlbC50b0xvd2VyQ2FzZSgpLmluY2x1ZGVzKHRoaXMuI3NlYXJjaFF1ZXJ5KSkge1xuICAgICAgICBpdGVtLnZpc2libGUgPSB0cnVlO1xuICAgICAgfVxuICAgIH1cbiAgICB0aGlzLiN1cGRhdGVWaXNpYmxlSXRlbXMoKTtcbiAgICB0aGlzLiNyZW5kZXJUcmVlKCk7XG4gIH1cblxuICAvKipcbiAgICogQ29sbGFwc2UgYWxsIGl0ZW1zXG4gICAqL1xuICBjb2xsYXBzZUFsbCgpIHtcbiAgICBmb3IgKGNvbnN0IGl0ZW0gb2YgdGhpcy4jZmxhdEl0ZW1zKSB7XG4gICAgICBpZiAoaXRlbS5oYXNDaGlsZHJlbikge1xuICAgICAgICBpdGVtLmV4cGFuZGVkID0gZmFsc2U7XG4gICAgICB9XG4gICAgICAvLyBPbmx5IHRvcC1sZXZlbCBpdGVtcyB2aXNpYmxlXG4gICAgICBpdGVtLnZpc2libGUgPSBpdGVtLmRlcHRoID09PSAwO1xuICAgIH1cbiAgICB0aGlzLiN1cGRhdGVWaXNpYmxlSXRlbXMoKTtcbiAgICB0aGlzLiNyZW5kZXJUcmVlKCk7XG4gIH1cbn1cblxuZGVjbGFyZSBnbG9iYWwge1xuICBpbnRlcmZhY2UgSFRNTEVsZW1lbnRUYWdOYW1lTWFwIHtcbiAgICAnY2VtLXZpcnR1YWwtdHJlZSc6IENlbVZpcnR1YWxUcmVlO1xuICB9XG59XG4iLCAiY29uc3Qgcz1uZXcgQ1NTU3R5bGVTaGVldCgpO3MucmVwbGFjZVN5bmMoSlNPTi5wYXJzZShcIlxcXCIvKiBWaXJ0dWFsIFRyZWUgLSBFZmZpY2llbnRseSByZW5kZXJzIGxhcmdlIHRyZWUgc3RydWN0dXJlcyAqL1xcXFxuXFxcXG46aG9zdCB7XFxcXG4gIGRpc3BsYXk6IGJsb2NrO1xcXFxuICBoZWlnaHQ6IDEwMCU7XFxcXG4gIG92ZXJmbG93OiBoaWRkZW47XFxcXG59XFxcXG5cXFxcbiN0cmVlLWNvbnRhaW5lciB7XFxcXG4gIHBvc2l0aW9uOiByZWxhdGl2ZTtcXFxcbiAgd2lkdGg6IDEwMCU7XFxcXG59XFxcXG5cXFxcbiN2aWV3cG9ydCB7XFxcXG4gIGhlaWdodDogMTAwJTtcXFxcbiAgb3ZlcmZsb3cteTogYXV0bztcXFxcbiAgb3ZlcmZsb3cteDogaGlkZGVuO1xcXFxuICBwb3NpdGlvbjogcmVsYXRpdmU7XFxcXG5cXFxcbiAgcGYtdjYtdHJlZS1pdGVtIHtcXFxcbiAgICBjdXJzb3I6IHBvaW50ZXI7XFxcXG5cXFxcbiAgICBcXFxcdTAwMjY6aG92ZXIge1xcXFxuICAgICAgYmFja2dyb3VuZC1jb2xvcjogdmFyKC0tcGYtdC0tZ2xvYmFsLS1iYWNrZ3JvdW5kLS1jb2xvci0tc2Vjb25kYXJ5LS1ob3ZlciwgI2Y1ZjVmNSk7XFxcXG4gICAgfVxcXFxuXFxcXG4gICAgLyogT25seSBvdXRsaW5lIHRoZSBhY3R1YWxseSBmb2N1c2VkIGl0ZW0sIG5vdCBhbmNlc3RvcnMgKi9cXFxcbiAgICBcXFxcdTAwMjY6Zm9jdXMtdmlzaWJsZSB7XFxcXG4gICAgICBvdXRsaW5lOiAycHggc29saWQgdmFyKC0tcGYtdC0tZ2xvYmFsLS1jb2xvci0tYnJhbmQtLWRlZmF1bHQsICMwMDY2Y2MpO1xcXFxuICAgICAgb3V0bGluZS1vZmZzZXQ6IC0ycHg7XFxcXG4gICAgfVxcXFxuICB9XFxcXG59XFxcXG5cXFwiXCIpKTtleHBvcnQgZGVmYXVsdCBzOyJdLAogICJtYXBwaW5ncyI6ICI7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFBQSxTQUFTLFlBQVksWUFBWTtBQUNqQyxTQUFTLHFCQUFxQjs7O0FDRDlCLElBQU0sSUFBRSxJQUFJLGNBQWM7QUFBRSxFQUFFLFlBQVksS0FBSyxNQUFNLGd0QkFBa3RCLENBQUM7QUFBRSxJQUFPLDJCQUFROzs7QURLenhCLE9BQU87QUEyREEsSUFBTSxrQkFBTixjQUE4QixNQUFNO0FBQUEsRUFDekM7QUFBQSxFQUNBLFlBQVksTUFBZ0I7QUFDMUIsVUFBTSxlQUFlLEVBQUUsU0FBUyxNQUFNLFVBQVUsS0FBSyxDQUFDO0FBQ3RELFNBQUssT0FBTztBQUFBLEVBQ2Q7QUFDRjtBQXRFQTtBQTZFQSw4QkFBQyxjQUFjLGtCQUFrQjtB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
