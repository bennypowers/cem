import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';

import styles from './cem-manifest-browser.css' with { type: 'css' };

import '../pf-v6-badge/pf-v6-badge.js';
import '../pf-v6-button/pf-v6-button.js';
import '../pf-v6-drawer/pf-v6-drawer.js';
import '../pf-v6-text-input-group/pf-v6-text-input-group.js';
import '../pf-v6-toolbar/pf-v6-toolbar.js';
import '../pf-v6-toolbar-group/pf-v6-toolbar-group.js';
import '../pf-v6-toolbar-item/pf-v6-toolbar-item.js';

// cem-detail-panel and cem-virtual-tree are still CemElement-based,
// imported as side-effect JS from their template locations.
import '../cem-detail-panel/cem-detail-panel.js';
import '../cem-virtual-tree/cem-virtual-tree.js';

// Re-export for type usage; CemVirtualTree provides loadManifest()
// @ts-ignore -- CemElement-based JS, no types available
import { CemVirtualTree } from '../cem-virtual-tree/cem-virtual-tree.js';

interface TreeItem {
  id: number;
  type: string;
  label: string;
  [key: string]: unknown;
}

interface DetailPanel extends HTMLElement {
  renderItem(item: TreeItem, manifest: unknown): Promise<void>;
}

interface VirtualTree extends HTMLElement {
  search(query: string): void;
  expandAll(): void;
  collapseAll(): void;
}

interface ItemSelectEvent extends Event {
  item: TreeItem;
}

/**
 * Manifest Browser with tree navigation and detail drawer
 *
 * @slot - Default slot (unused)
 */
@customElement('cem-manifest-browser')
export class CemManifestBrowser extends LitElement {
  static styles = styles;

  #searchDebounceTimer: ReturnType<typeof setTimeout> | null = null;

  disconnectedCallback() {
    super.disconnectedCallback();
    if (this.#searchDebounceTimer != null) {
      clearTimeout(this.#searchDebounceTimer);
      this.#searchDebounceTimer = null;
    }
  }

  render() {
    return html`
      <div id="drawer-content" slot="content">
        <pf-v6-toolbar sticky>
          <pf-v6-toolbar-group>
            <pf-v6-toolbar-item>
              <pf-v6-text-input-group id="search"
                                      type="search"
                                      placeholder="Search manifest..."
                                      aria-label="Search manifest"
                                      icon
                                      @input=${this.#onSearchInput}>
                <svg slot="icon"
                     role="presentation"
                     fill="currentColor"
                     height="1em"
                     width="1em"
                     viewBox="0 0 512 512">
                  <path d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                </svg>
                <pf-v6-badge slot="utilities"
                             id="search-count"
                             compact
                             hidden>
                  0<span class="pf-v6-screen-reader"> results</span>
                </pf-v6-badge>
                <pf-v6-button slot="utilities"
                              id="search-clear"
                              variant="plain"
                              aria-label="Clear search"
                              hidden
                              @click=${this.#onSearchClear}>
                  <svg width="1em"
                       height="1em"
                       viewBox="0 0 352 512"
                       fill="currentColor"
                       aria-hidden="true">
                    <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
                  </svg>
                </pf-v6-button>
              </pf-v6-text-input-group>
            </pf-v6-toolbar-item>
          </pf-v6-toolbar-group>
          <pf-v6-toolbar-group variant="action-group">
            <pf-v6-toolbar-item>
              <pf-v6-button id="expand-all"
                            variant="tertiary"
                            size="small"
                            aria-label="Expand all tree items"
                            @click=${this.#onExpandAll}>
                Expand all
              </pf-v6-button>
            </pf-v6-toolbar-item>
            <pf-v6-toolbar-item>
              <pf-v6-button id="collapse-all"
                            variant="tertiary"
                            size="small"
                            aria-label="Collapse all tree items"
                            @click=${this.#onCollapseAll}>
                Collapse all
              </pf-v6-button>
            </pf-v6-toolbar-item>
          </pf-v6-toolbar-group>
        </pf-v6-toolbar>
        <pf-v6-drawer id="drawer">
          <div id="tree-wrapper">
            <cem-virtual-tree id="virtual-tree"
                              @item-select=${this.#onItemSelect}></cem-virtual-tree>
          </div>
          <cem-detail-panel id="detail-panel"
                            slot="panel-body"></cem-detail-panel>
        </pf-v6-drawer>
      </div>
    `;
  }

  get #drawer() {
    return this.shadowRoot?.getElementById('drawer') as HTMLElement & { expanded: boolean } | null;
  }

  get #virtualTree(): VirtualTree | null {
    return this.shadowRoot?.getElementById('virtual-tree') as VirtualTree | null;
  }

  get #detailPanel(): DetailPanel | null {
    return this.shadowRoot?.getElementById('detail-panel') as DetailPanel | null;
  }

  get #searchInput() {
    return this.shadowRoot?.getElementById('search') as HTMLElement & { value: string } | null;
  }

  get #searchCount() {
    return this.shadowRoot?.getElementById('search-count') as HTMLElement | null;
  }

  get #searchClear() {
    return this.shadowRoot?.getElementById('search-clear') as HTMLElement | null;
  }

  async #onItemSelect(e: ItemSelectEvent) {
    const item = e.item;
    if (!item) return;

    const manifest = await CemVirtualTree.loadManifest();
    if (this.#detailPanel && manifest) {
      await this.#detailPanel.renderItem(item, manifest);
      if (this.#drawer) {
        this.#drawer.expanded = true;
      }
    }
  }

  #onSearchInput() {
    const value = this.#searchInput?.value || '';

    // Show/hide clear button based on whether there's a value
    if (this.#searchClear) {
      this.#searchClear.hidden = !value;
    }

    // Debounce search - wait 300ms after user stops typing
    if (this.#searchDebounceTimer != null) {
      clearTimeout(this.#searchDebounceTimer);
    }
    this.#searchDebounceTimer = setTimeout(() => {
      this.#handleSearch(value);
    }, 300);
  }

  #onSearchClear() {
    if (this.#searchInput) {
      this.#searchInput.value = '';
    }
    if (this.#searchClear) {
      this.#searchClear.hidden = true;
    }
    if (this.#searchCount) {
      this.#searchCount.hidden = true;
    }
    this.#handleSearch('');
  }

  #handleSearch(query: string) {
    this.#virtualTree?.search(query);
    if (this.#searchCount) {
      this.#searchCount.hidden = !query;
    }
  }

  #onExpandAll() {
    this.#virtualTree?.expandAll();
  }

  #onCollapseAll() {
    this.#virtualTree?.collapseAll();
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-manifest-browser': CemManifestBrowser;
  }
}
