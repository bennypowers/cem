import '/__cem/elements/pf-v6-tree-item/pf-v6-tree-item.js';
import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event for item selection
 */
export class ItemSelectEvent extends Event {
  constructor(item) {
    super('item-select', { bubbles: true, composed: true });
    this.item = item;
  }
}

/**
 * Virtual tree view with lazy loading
 * @customElement cem-virtual-tree
 */
export class CemVirtualTree extends CemElement {
  static is = 'cem-virtual-tree';

  // Static cache for manifest (shared across all instances)
  static #manifestCache = null;
  static #manifestPromise = null;

  #manifest = null;
  #flatItems = [];
  #visibleItems = [];
  #searchQuery = '';
  #viewport = null;
  #currentSelectedElement = null;
  #currentSelectedItemId = null;

  async afterTemplateLoaded() {
    this.#viewport = this.shadowRoot.getElementById('viewport');

    // Load manifest from server with caching
    this.#manifest = await this.#loadManifest();
    if (!this.#manifest) {
      console.warn('[virtual-tree] Failed to load manifest');
      return;
    }

    // Build flat list from manifest
    this.#buildFlatList();

    // Initial render
    this.#render();
  }

  /**
   * Public static method to load manifest with caching
   * Can be called by other components to reuse the same cached manifest
   * @returns {Promise<Object|null>} The manifest object or null on error
   */
  static async loadManifest() {
    // Return cached manifest if available
    if (CemVirtualTree.#manifestCache) {
      return CemVirtualTree.#manifestCache;
    }

    // If already fetching, wait for existing promise
    if (CemVirtualTree.#manifestPromise) {
      return CemVirtualTree.#manifestPromise;
    }

    // Fetch manifest
    CemVirtualTree.#manifestPromise = fetch('/custom-elements.json')
      .then(async (response) => {
        if (!response.ok) {
          throw new Error(`Failed to fetch manifest: ${response.status}`);
        }
        const manifest = await response.json();
        CemVirtualTree.#manifestCache = manifest;
        return manifest;
      })
      .catch((error) => {
        console.error('[virtual-tree] Error loading manifest:', error);
        CemVirtualTree.#manifestPromise = null;
        return null;
      });

    return CemVirtualTree.#manifestPromise;
  }

  /**
   * Clear the static manifest cache (for testing)
   */
  static clearCache() {
    CemVirtualTree.#manifestCache = null;
    CemVirtualTree.#manifestPromise = null;
  }

  /**
   * Load manifest from server with static caching
   * Ensures the manifest is only fetched once across all instances
   */
  async #loadManifest() {
    return CemVirtualTree.loadManifest();
  }

  /**
   * Build flat list from hierarchical manifest
   */
  #buildFlatList() {
    this.#flatItems = [];
    let id = 0;

    // Check if this is a workspace manifest with multiple packages
    const hasMultiplePackages = this.#manifest.packages && this.#manifest.packages.length > 1;

    if (hasMultiplePackages) {
      // Workspace mode with multiple packages: show packages at top level
      for (const pkg of this.#manifest.packages) {
        const packageId = id++;
        const packageItem = {
          id: packageId,
          type: 'package',
          label: pkg.name,
          depth: 0,
          hasChildren: pkg.modules?.length > 0,
          expanded: false,
          visible: true,
          packageName: pkg.name,
          badge: pkg.modules?.length || 0,
        };
        this.#flatItems.push(packageItem);

        if (!pkg.modules) continue;

        // Add modules under this package at depth 1
        this.#buildModulesForPackage(pkg.modules, packageId, 1, id);
        id = this.#flatItems[this.#flatItems.length - 1].id + 1;
      }
    } else {
      // Single package mode OR workspace with 1 package: show modules at top level
      const modules = this.#manifest.packages?.[0]?.modules || this.#manifest.modules;

      if (!modules) return;

      // Add modules at depth 0 (no package level)
      this.#buildModulesForPackage(modules, null, 0, id);
    }

    this.#updateVisibleItems();
  }

  /**
   * Build modules and their declarations for a package
   * @param {Array} modules - Array of modules
   * @param {number|null} parentId - ID of parent package (null if no package level)
   * @param {number} depth - Depth level for modules
   * @param {number} startId - Starting ID for items
   */
  #buildModulesForPackage(modules, parentId, depth, startId) {
    let id = startId;

    for (const module of modules) {
      const moduleId = id++;
      const moduleItem = {
        id: moduleId,
        type: 'module',
        label: module.path,
        depth: depth,
        parentId: parentId,
        hasChildren: module.declarations?.length > 0,
        expanded: false,
        visible: parentId === null, // Visible if no parent, otherwise hidden until parent expands
        modulePath: module.path,
        badge: module.declarations?.length || 0,
      };
      this.#flatItems.push(moduleItem);

      if (!module.declarations) continue;

      for (const decl of module.declarations) {
        // Custom Element - must have both customElement flag and tagName
        if (decl.kind === 'class' && decl.customElement && decl.tagName) {
          const ceId = id++;

          // Compute hasChildren based on presence of any child arrays
          const properties = decl.members?.filter(m => m.kind === 'field') || [];
          const methods = decl.members?.filter(m => m.kind === 'method') || [];
          const hasChildren = [
            decl.attributes,
            properties,
            methods,
            decl.events,
            decl.slots,
            decl.cssProperties,
            decl.cssParts,
            decl.cssStates,
            decl.demos,
          ].some(x => x?.length > 0)

          const ceItem = {
            id: ceId,
            type: 'custom-element',
            label: `<${decl.tagName}>`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren,
            expanded: false,
            visible: false,
            modulePath: module.path,
            tagName: decl.tagName,
          };
          this.#flatItems.push(ceItem);

          // Attributes
          if (decl.attributes?.length) {
            const attrCatId = id++;
            this.#flatItems.push({
              id: attrCatId,
              type: 'category',
              label: 'Attributes',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.attributes.length,
              category: 'attributes',
            });

            for (const attr of decl.attributes) {
              this.#flatItems.push({
                id: id++,
                type: 'attribute',
                label: attr.name,
                depth: depth + 3,
                parentId: attrCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: attr.name,
              });
            }
          }

          // Properties
          if (properties.length) {
            const propCatId = id++;
            this.#flatItems.push({
              id: propCatId,
              type: 'category',
              label: 'Properties',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: properties.length,
              category: 'properties',
            });

            for (const prop of properties) {
              this.#flatItems.push({
                id: id++,
                type: 'property',
                label: prop.name,
                depth: depth + 3,
                parentId: propCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: prop.name,
              });
            }
          }

          // Methods
          if (methods.length) {
            const methodCatId = id++;
            this.#flatItems.push({
              id: methodCatId,
              type: 'category',
              label: 'Methods',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: methods.length,
              category: 'methods',
            });

            for (const method of methods) {
              this.#flatItems.push({
                id: id++,
                type: 'method',
                label: `${method.name}()`,
                depth: depth + 3,
                parentId: methodCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: method.name,
              });
            }
          }

          // Events
          if (decl.events?.length) {
            const eventCatId = id++;
            this.#flatItems.push({
              id: eventCatId,
              type: 'category',
              label: 'Events',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.events.length,
              category: 'events',
            });

            for (const event of decl.events) {
              this.#flatItems.push({
                id: id++,
                type: 'event',
                label: event.name,
                depth: depth + 3,
                parentId: eventCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: event.name,
              });
            }
          }

          // Slots
          if (decl.slots?.length) {
            const slotCatId = id++;
            this.#flatItems.push({
              id: slotCatId,
              type: 'category',
              label: 'Slots',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.slots.length,
              category: 'slots',
            });

            for (const slot of decl.slots) {
              this.#flatItems.push({
                id: id++,
                type: 'slot',
                label: slot.name || '(default)',
                depth: depth + 3,
                parentId: slotCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: slot.name,
              });
            }
          }

          // CSS Properties
          if (decl.cssProperties?.length) {
            const cssPropCatId = id++;
            this.#flatItems.push({
              id: cssPropCatId,
              type: 'category',
              label: 'CSS Properties',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssProperties.length,
              category: 'css-properties',
            });

            for (const cssProp of decl.cssProperties) {
              this.#flatItems.push({
                id: id++,
                type: 'css-property',
                label: cssProp.name,
                depth: depth + 3,
                parentId: cssPropCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssProp.name,
              });
            }
          }

          // CSS Parts
          if (decl.cssParts?.length) {
            const cssPartCatId = id++;
            this.#flatItems.push({
              id: cssPartCatId,
              type: 'category',
              label: 'CSS Parts',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssParts.length,
              category: 'css-parts',
            });

            for (const cssPart of decl.cssParts) {
              this.#flatItems.push({
                id: id++,
                type: 'css-part',
                label: cssPart.name,
                depth: depth + 3,
                parentId: cssPartCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssPart.name,
              });
            }
          }

          // CSS States
          if (decl.cssStates?.length) {
            const cssStateCatId = id++;
            this.#flatItems.push({
              id: cssStateCatId,
              type: 'category',
              label: 'CSS States',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssStates.length,
              category: 'css-states',
            });

            for (const cssState of decl.cssStates) {
              this.#flatItems.push({
                id: id++,
                type: 'css-state',
                label: cssState.name,
                depth: depth + 3,
                parentId: cssStateCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssState.name,
              });
            }
          }

          // Demos
          if (decl.demos?.length) {
            const demoCatId = id++;
            this.#flatItems.push({
              id: demoCatId,
              type: 'category',
              label: 'Demos',
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.demos.length,
              category: 'demos',
            });

            for (const demo of decl.demos) {
              this.#flatItems.push({
                id: id++,
                type: 'demo',
                label: demo.name || demo.url || '(demo)',
                depth: depth + 3,
                parentId: demoCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                url: demo.url,
              });
            }
          }
        } else if (decl.kind === 'class') {
          // Regular class
          this.#flatItems.push({
            id: id++,
            type: 'class',
            label: decl.name,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name,
          });
        } else if (decl.kind === 'function') {
          this.#flatItems.push({
            id: id++,
            type: 'function',
            label: `${decl.name}()`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name,
          });
        } else if (decl.kind === 'variable') {
          this.#flatItems.push({
            id: id++,
            type: 'variable',
            label: decl.name,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name,
          });
        } else if (decl.kind === 'mixin') {
          this.#flatItems.push({
            id: id++,
            type: 'mixin',
            label: `${decl.name}()`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name,
          });
        }
      }
    }

    this.#updateVisibleItems();
  }

  /**
   * Update visible items based on expanded state
   */
  #updateVisibleItems() {
    this.#visibleItems = this.#flatItems.filter(item => item.visible);
  }


  /**
   * Render visible tree items with proper nesting
   */
  #render() {
    if (!this.#viewport) return;

    // Clear viewport
    this.#viewport.innerHTML = '';

    // Build nested tree structure from flat visible items
    const rootItems = this.#visibleItems.filter(item => item.depth === 0);

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
    const el = document.createElement('pf-v6-tree-item');
    el.setAttribute('label', item.label);

    if (item.badge) {
      el.setAttribute('badge', item.badge);
    }

    // Store item data
    el.dataset.itemId = item.id;
    el.dataset.type = item.type;

    // Restore selection state if this item was selected
    if (this.#currentSelectedItemId === item.id) {
      el.setAttribute('current', '');
      this.#currentSelectedElement = el;
    }

    // Set has-children attribute for items with children (shows toggle button)
    if (item.hasChildren) {
      el.setAttribute('has-children', '');
      el.expanded = item.expanded;

      // Listen for expand/collapse
      // Stop propagation to prevent nested collapses from affecting parent items
      el.addEventListener('expand', (e) => {
        e.stopPropagation();
        this.#handleToggle(item, true);
      });

      el.addEventListener('collapse', (e) => {
        e.stopPropagation();
        this.#handleToggle(item, false);
      });
    }

    // Listen for selection (non-category items)
    if (item.type !== 'category') {
      el.addEventListener('select', (e) => {
        e.stopPropagation();
        this.#handleSelect(item, el);
      });
    }

    // Add children if expanded and visible
    if (item.expanded && item.hasChildren) {
      const children = this.#visibleItems.filter(child => child.parentId === item.id && child.visible);
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
    this.#render();
  }

  /**
   * Update visibility of children when parent is toggled
   */
  #updateChildVisibility(parentItem) {
    const isExpanded = parentItem.expanded;

    for (const item of this.#flatItems) {
      if (item.parentId === parentItem.id) {
        item.visible = isExpanded;
        // If collapsing, also collapse and hide all descendants
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
    // Clear previous selection
    if (this.#currentSelectedElement) {
      this.#currentSelectedElement.removeAttribute('current');
    }

    // Set new selection
    element.setAttribute('current', '');
    this.#currentSelectedElement = element;
    this.#currentSelectedItemId = item.id;

    // Dispatch custom event
    this.dispatchEvent(new ItemSelectEvent(item));
  }

  /**
   * Search/filter tree items
   * @param {string} query - Search query
   */
  search(query) {
    this.#searchQuery = query.toLowerCase().trim();

    if (!this.#searchQuery) {
      // Reset to normal visibility
      for (const item of this.#flatItems) {
        item.visible = item.depth === 0 || (item.parentId !== undefined && this.#isParentExpanded(item));
      }
    } else {
      // Filter based on search
      const matchingItems = new Set();

      // Find matching items
      for (const item of this.#flatItems) {
        if (item.label.toLowerCase().includes(this.#searchQuery)) {
          matchingItems.add(item);
          // Show and expand all ancestors
          this.#showAncestors(item);
        }
      }

      // Hide non-matching items
      for (const item of this.#flatItems) {
        if (!matchingItems.has(item) && !this.#hasMatchingDescendant(item, matchingItems)) {
          item.visible = this.#isAncestorOfMatch(item, matchingItems);
        } else {
          item.visible = true;
        }
      }
    }

    this.#updateVisibleItems();
    this.#render();
  }

  /**
   * Check if parent is expanded
   */
  #isParentExpanded(item) {
    if (item.parentId === undefined) return true;
    const parent = this.#flatItems.find(i => i.id === item.parentId);
    return parent && parent.expanded && this.#isParentExpanded(parent);
  }

  /**
   * Show all ancestors of an item
   */
  #showAncestors(item) {
    if (item.parentId === undefined) return;
    const parent = this.#flatItems.find(i => i.id === item.parentId);
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
      while (current.parentId !== undefined) {
        const parent = this.#flatItems.find(i => i.id === current.parentId);
        if (!parent) break; // Parent not found in tree
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
    while (current.parentId !== undefined) {
      const parent = this.#flatItems.find(i => i.id === current.parentId);
      if (!parent) return false; // Parent not found in tree
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
      // Make all items visible except those filtered by search
      if (!this.#searchQuery || item.label.toLowerCase().includes(this.#searchQuery)) {
        item.visible = true;
      }
    }
    this.#updateVisibleItems();
    this.#render();
  }

  /**
   * Collapse all items
   */
  collapseAll() {
    for (const item of this.#flatItems) {
      if (item.hasChildren) {
        item.expanded = false;
      }
      // Only top-level items visible
      item.visible = item.depth === 0;
    }
    this.#updateVisibleItems();
    this.#render();
  }

  static {
    customElements.define(this.is, this);
  }
}
