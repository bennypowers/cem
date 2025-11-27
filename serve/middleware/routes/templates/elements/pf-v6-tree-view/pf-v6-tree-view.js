import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Tree View
 *
 * Container for pf-v6-tree-item elements. Coordinates selection and keyboard navigation.
 *
 * @slot - Tree items (pf-v6-tree-item elements)
 *
 * @attr {string} aria-label - Accessible label for the tree
 *
 * @customElement pf-v6-tree-view
 */
class PfV6TreeView extends CemElement {
  static is = 'pf-v6-tree-view';

  static observedAttributes = ['aria-label'];

  #$ = id => this.shadowRoot.getElementById(id);
  #tree;
  #currentSelection = null;

  get ariaLabel() { return this.getAttribute('aria-label'); }
  set ariaLabel(value) {
    if (value) {
      this.setAttribute('aria-label', value);
    } else {
      this.removeAttribute('aria-label');
    }
  }

  afterTemplateLoaded() {
    this.#tree = this.#$('tree');
    if (!this.#tree) return;

    this.#updateAriaAttributes();

    // Listen for select events from tree items
    this.addEventListener('select', (e) => {
      this.#handleItemSelect(e.target);
    });

    // Set up keyboard navigation at tree level
    this.#setupKeyboardNavigation();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue) return;

    if (name === 'aria-label') {
      this.#updateAriaAttributes();
    }
  }

  #updateAriaAttributes() {
    if (!this.#tree) return;

    const ariaLabel = this.ariaLabel;
    if (ariaLabel) {
      this.#tree.setAttribute('aria-label', ariaLabel);
    } else {
      this.#tree.removeAttribute('aria-label');
    }
  }

  /**
   * Handle item selection - ensure only one item is current
   */
  #handleItemSelect(item) {
    // Clear previous selection
    if (this.#currentSelection && this.#currentSelection !== item) {
      this.#currentSelection.deselect();
    }

    this.#currentSelection = item;
  }

  /**
   * Set up keyboard navigation at tree level
   */
  #setupKeyboardNavigation() {
    this.addEventListener('keydown', (e) => {
      const target = e.target;

      // Only handle if target is a tree item
      if (target.tagName !== 'PF-V6-TREE-ITEM') return;

      const items = this.#getAllVisibleItems();
      const currentIndex = items.indexOf(target);

      switch (e.key) {
        case 'ArrowDown':
          e.preventDefault();
          if (currentIndex < items.length - 1) {
            items[currentIndex + 1].focus();
          }
          break;

        case 'ArrowUp':
          e.preventDefault();
          if (currentIndex > 0) {
            items[currentIndex - 1].focus();
          }
          break;

        case 'Home':
          e.preventDefault();
          if (items.length > 0) {
            items[0].focus();
          }
          break;

        case 'End':
          e.preventDefault();
          if (items.length > 0) {
            items[items.length - 1].focus();
          }
          break;
      }
    });
  }

  /**
   * Get all visible tree items in document order
   */
  #getAllVisibleItems() {
    return Array.from(this.querySelectorAll('pf-v6-tree-item'));
  }

  /**
   * Expand all tree items
   */
  expandAll() {
    const items = this.querySelectorAll('pf-v6-tree-item');
    items.forEach(item => item.expand());
  }

  /**
   * Collapse all tree items
   */
  collapseAll() {
    const items = this.querySelectorAll('pf-v6-tree-item');
    items.forEach(item => item.collapse());
  }

  static {
    customElements.define(this.is, this);
  }
}
