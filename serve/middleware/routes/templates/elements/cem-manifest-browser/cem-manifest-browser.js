import '/__cem/elements/pf-v6-button/pf-v6-button.js';
import '/__cem/elements/pf-v6-drawer/pf-v6-drawer.js';
import '/__cem/elements/pf-v6-text-input-group/pf-v6-text-input-group.js';
import '/__cem/elements/pf-v6-toolbar/pf-v6-toolbar.js';
import '/__cem/elements/pf-v6-toolbar-group/pf-v6-toolbar-group.js';
import '/__cem/elements/pf-v6-toolbar-item/pf-v6-toolbar-item.js';
import { CemElement } from '/__cem/cem-element.js';

/**
 * Manifest Browser with tree navigation and detail drawer
 * @customElement cem-manifest-browser
 */
export class CemManifestBrowser extends CemElement {
  static is = 'cem-manifest-browser';

  #panelTitle;
  #drawer;
  #currentDetail;
  #searchDebounceTimer = null;

  afterTemplateLoaded() {
    this.#panelTitle = this.shadowRoot.getElementById('panel-title');
    this.#drawer = this.shadowRoot.getElementById('drawer');

    // Listen for tree item selections
    this.addEventListener('select', (e) => {
      const treeItem = e.target;
      if (treeItem.tagName !== 'PF-V6-TREE-ITEM') return;

      this.#handleItemSelect(treeItem);
    });

    // Listen for search input with debouncing
    const searchInput = this.shadowRoot.getElementById('search');
    if (searchInput) {
      searchInput.addEventListener('input', () => {
        // Use getAttribute since the value property reads from the attribute
        // and might not be updated yet in the event handler
        const value = searchInput.getAttribute('value') || '';

        // Debounce search - wait 300ms after user stops typing
        clearTimeout(this.#searchDebounceTimer);
        this.#searchDebounceTimer = setTimeout(() => {
          this.#handleSearch(value);
        }, 300);
      });
    } else {
      console.warn('[manifest search] search input not found');
    }

    // Listen for expand/collapse all buttons
    const expandAllBtn = this.shadowRoot.getElementById('expand-all');
    const collapseAllBtn = this.shadowRoot.getElementById('collapse-all');

    if (expandAllBtn) {
      expandAllBtn.addEventListener('click', () => {
        this.#expandAll();
      });
    }

    if (collapseAllBtn) {
      collapseAllBtn.addEventListener('click', () => {
        this.#collapseAll();
      });
    }
  }

  #handleItemSelect(treeItem) {
    const type = treeItem.getAttribute('data-type');

    if (!type || type === 'category') {
      return;
    }

    // Build selector from tree item attributes
    const selector = this.#buildSelector(type, treeItem);

    // Find all matching elements (both h3 and dl)
    const matchingElements = document.body.querySelectorAll(selector);

    if (matchingElements.length === 0) {
      return;
    }

    // Hide current detail elements if exist
    if (this.#currentDetail && this.#currentDetail.length > 0) {
      this.#currentDetail.forEach(el => el.hidden = true);
    }

    // Show new detail elements
    matchingElements.forEach(el => el.hidden = false);
    this.#currentDetail = matchingElements;

    // Open drawer
    this.#drawer.expanded = true;
  }

  #buildSelector(type, treeItem) {
    const selectors = [`[data-type="${type}"]`];

    const modulePath = treeItem.getAttribute('data-module-path') || treeItem.getAttribute('data-path');
    const tagName = treeItem.getAttribute('data-tag-name');
    const name = treeItem.getAttribute('data-name');

    if (modulePath) {
      selectors.push(`[data-module-path="${modulePath}"]`);
    }
    if (tagName) {
      selectors.push(`[data-tag-name="${tagName}"]`);
    }
    if (name) {
      selectors.push(`[data-name="${name}"]`);
    }

    return selectors.join('');
  }

  #getTreeView() {
    // Get the slot in shadow DOM and access its assigned elements
    const slot = this.shadowRoot.querySelector('slot[name="manifest-tree"]');
    if (!slot) {
      console.warn('[manifest browser] manifest-tree slot not found');
      return null;
    }

    // Get the slotted element (might be another slot due to slot forwarding)
    let assignedElements = slot.assignedElements();
    let element = assignedElements[0];

    // If we got a slot element (slot forwarding), get ITS assigned elements
    if (element && element.tagName === 'SLOT') {
      assignedElements = element.assignedElements();
      element = assignedElements[0];
    }

    if (!element) {
      console.warn('[manifest browser] no tree-view assigned to slot');
      return null;
    }

    return element;
  }

  #handleSearch(query) {
    const normalizedQuery = query.toLowerCase().trim();
    const treeView = this.#getTreeView();
    if (!treeView) return;

    // Query tree items from the slotted tree-view
    const items = treeView.querySelectorAll('pf-v6-tree-item');

    if (!normalizedQuery) {
      // Show all items when search is empty
      items.forEach(item => {
        item.hidden = false;
      });
      return;
    }

    // First pass: hide all items and find matches
    const matchingItems = new Set();
    items.forEach(item => {
      const label = item.getAttribute('label') || '';
      const matches = label.toLowerCase().includes(normalizedQuery);
      item.hidden = true;
      if (matches) {
        matchingItems.add(item);
      }
    });

    // Second pass: for each matching item, show it and all ancestors
    matchingItems.forEach(item => {
      // Show the matching item
      item.hidden = false;

      // Walk up the tree, showing and expanding all ancestors
      let parent = item.parentElement;
      while (parent && parent.tagName === 'PF-V6-TREE-ITEM') {
        parent.hidden = false;
        parent.expanded = true;
        parent = parent.parentElement;
      }
    });
  }

  #expandAll() {
    const treeView = this.#getTreeView();
    if (!treeView) return;

    const items = treeView.querySelectorAll('pf-v6-tree-item');
    items.forEach(item => {
      item.expanded = true;
    });
  }

  #collapseAll() {
    const treeView = this.#getTreeView();
    if (!treeView) return;

    const items = treeView.querySelectorAll('pf-v6-tree-item');
    items.forEach(item => {
      item.expanded = false;
    });
  }

  static {
    customElements.define(this.is, this);
  }
}
