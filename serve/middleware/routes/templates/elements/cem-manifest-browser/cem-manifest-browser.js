import '/__cem/elements/pf-v6-badge/pf-v6-badge.js';
import '/__cem/elements/pf-v6-button/pf-v6-button.js';
import '/__cem/elements/pf-v6-drawer/pf-v6-drawer.js';
import '/__cem/elements/pf-v6-text-input-group/pf-v6-text-input-group.js';
import '/__cem/elements/pf-v6-toolbar/pf-v6-toolbar.js';
import '/__cem/elements/pf-v6-toolbar-group/pf-v6-toolbar-group.js';
import '/__cem/elements/pf-v6-toolbar-item/pf-v6-toolbar-item.js';
import '/__cem/elements/cem-virtual-tree/cem-virtual-tree.js';
import '/__cem/elements/cem-detail-panel/cem-detail-panel.js';
import { CemElement } from '/__cem/cem-element.js';

/**
 * Manifest Browser with tree navigation and detail drawer
 * @customElement cem-manifest-browser
 */
export class CemManifestBrowser extends CemElement {
  static is = 'cem-manifest-browser';

  #drawer;
  #virtualTree;
  #detailPanel;
  #searchDebounceTimer = null;
  #searchInput;
  #searchCount;
  #searchClear;

  afterTemplateLoaded() {
    this.#drawer = this.shadowRoot.getElementById('drawer');
    this.#virtualTree = this.shadowRoot.getElementById('virtual-tree');
    this.#detailPanel = this.shadowRoot.getElementById('detail-panel');
    this.#searchInput = this.shadowRoot.getElementById('search');
    this.#searchCount = this.shadowRoot.getElementById('search-count');
    this.#searchClear = this.shadowRoot.getElementById('search-clear');

    // Listen for item selection from virtual tree
    this.#virtualTree?.addEventListener('item-select', (e) => {
      this.#handleItemSelect(e.item);
    });

    // Listen for search input with debouncing
    if (this.#searchInput) {
      this.#searchInput.addEventListener('input', () => {
        const value = this.#searchInput.value || '';

        // Show/hide clear button based on whether there's a value
        if (this.#searchClear) {
          this.#searchClear.hidden = !value;
        }

        // Debounce search - wait 300ms after user stops typing
        clearTimeout(this.#searchDebounceTimer);
        this.#searchDebounceTimer = setTimeout(() => {
          this.#handleSearch(value);
        }, 300);
      });
    }

    // Listen for clear button
    if (this.#searchClear) {
      this.#searchClear.addEventListener('click', () => {
        // Clear the input
        if (this.#searchInput) {
          this.#searchInput.value = '';
        }
        // Hide clear button and badge
        this.#searchClear.hidden = true;
        if (this.#searchCount) {
          this.#searchCount.hidden = true;
        }
        // Reset search
        this.#handleSearch('');
      });
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

  async #handleItemSelect(item) {
    if (!item) return;

    // Render detail panel for this item
    const manifest = window.__CEM_MANIFEST__;
    if (this.#detailPanel && manifest) {
      await this.#detailPanel.renderItem(item, manifest);
      // Open drawer
      if (this.#drawer) {
        this.#drawer.expanded = true;
      }
    }
  }

  #handleSearch(query) {
    // Delegate search to virtual tree
    if (this.#virtualTree) {
      this.#virtualTree.search(query);
    }

    // Update search count badge (optional - could be handled by virtual tree)
    if (this.#searchCount) {
      // For now, hide the count badge as we're delegating to virtual tree
      // Virtual tree could emit events with count if needed
      this.#searchCount.hidden = !query;
    }
  }

  #expandAll() {
    if (this.#virtualTree) {
      this.#virtualTree.expandAll();
    }
  }

  #collapseAll() {
    if (this.#virtualTree) {
      this.#virtualTree.collapseAll();
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
