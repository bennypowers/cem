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

  #manifest = null;
  #flatItems = [];
  #visibleItems = [];
  #searchQuery = '';
  #viewport = null;
  #currentSelectedElement = null;

  afterTemplateLoaded() {
    this.#viewport = this.shadowRoot.getElementById('viewport');

    // Load manifest
    this.#manifest = window.__CEM_MANIFEST__;
    if (!this.#manifest) {
      console.warn('[virtual-tree] No manifest found in window.__CEM_MANIFEST__');
      return;
    }

    // Build flat list from manifest
    this.#buildFlatList();

    // Initial render
    this.#render();
  }

  /**
   * Build flat list from hierarchical manifest
   */
  #buildFlatList() {
    this.#flatItems = [];
    let id = 0;

    if (!this.#manifest.modules) return;

    for (const module of this.#manifest.modules) {
      const moduleId = id++;
      const moduleItem = {
        id: moduleId,
        type: 'module',
        label: module.path,
        depth: 0,
        hasChildren: module.declarations?.length > 0,
        expanded: false,
        visible: true,
        modulePath: module.path,
        badge: module.declarations?.length || 0,
      };
      this.#flatItems.push(moduleItem);

      if (!module.declarations) continue;

      for (const decl of module.declarations) {
        // Custom Element
        if (decl.kind === 'class' && decl.customElement) {
          const ceId = id++;
          const ceItem = {
            id: ceId,
            type: 'custom-element',
            label: `<${decl.tagName}>`,
            depth: 1,
            parentId: moduleId,
            hasChildren: true,
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
              depth: 2,
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
                depth: 3,
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
          const properties = decl.members?.filter(m => m.kind === 'field') || [];
          if (properties.length) {
            const propCatId = id++;
            this.#flatItems.push({
              id: propCatId,
              type: 'category',
              label: 'Properties',
              depth: 2,
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
                depth: 3,
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
          const methods = decl.members?.filter(m => m.kind === 'method') || [];
          if (methods.length) {
            const methodCatId = id++;
            this.#flatItems.push({
              id: methodCatId,
              type: 'category',
              label: 'Methods',
              depth: 2,
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
                depth: 3,
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
              depth: 2,
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
                depth: 3,
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
              depth: 2,
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
                depth: 3,
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
        } else if (decl.kind === 'class') {
          // Regular class
          this.#flatItems.push({
            id: id++,
            type: 'class',
            label: decl.name,
            depth: 1,
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
            depth: 1,
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
            depth: 1,
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
            depth: 1,
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
