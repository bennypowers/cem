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

    // Initialize roving tabindex: first item should be focusable
    this.#initializeTabindex();
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
            this.#focusItem(items[currentIndex + 1]);
          }
          break;

        case 'ArrowUp':
          e.preventDefault();
          if (currentIndex > 0) {
            this.#focusItem(items[currentIndex - 1]);
          }
          break;

        case 'ArrowRight':
          e.preventDefault();
          if (target.hasChildren) {
            if (!target.expanded) {
              // Collapsed with children: expand it
              target.expand();
            } else {
              // Already expanded: move to first child
              const children = this.#getDirectChildren(target);
              if (children.length > 0) {
                this.#focusItem(children[0]);
              }
            }
          }
          // If no children, do nothing
          break;

        case 'ArrowLeft':
          e.preventDefault();
          if (target.hasChildren && target.expanded) {
            // Expanded: collapse it
            target.collapse();
          } else {
            // Collapsed or no children: move to parent
            const parent = this.#getParentItem(target);
            if (parent) {
              this.#focusItem(parent);
            }
          }
          break;

        case 'Home':
          e.preventDefault();
          if (items.length > 0) {
            this.#focusItem(items[0]);
          }
          break;

        case 'End':
          e.preventDefault();
          if (items.length > 0) {
            this.#focusItem(items[items.length - 1]);
          }
          break;

        case 'Enter':
        case ' ':
          e.preventDefault();
          target.select();
          if (target.hasChildren) {
            target.toggle();
          }
          break;

        case '*': {
          e.preventDefault();
          // Expand all siblings at the same level
          const parent = this.#getParentItem(target);
          const siblings = parent ? this.#getDirectChildren(parent) : this.#getTopLevelItems();
          siblings.forEach(item => {
            if (item.hasChildren) {
              item.expand();
            }
          });
          break;
        }
      }
    });
  }

  /**
   * Get all visible tree items in document order
   * Only returns items that are visible (not inside collapsed parents)
   */
  #getAllVisibleItems() {
    const visible = [];
    const walk = (parent) => {
      const children = Array.from(parent.children).filter(el => el.tagName === 'PF-V6-TREE-ITEM');
      for (const item of children) {
        visible.push(item);
        // Only recurse into children if this item is expanded
        if (item.expanded && item.hasChildren) {
          walk(item);
        }
      }
    };
    walk(this);
    return visible;
  }

  /**
   * Get direct children of a tree item
   */
  #getDirectChildren(item) {
    return Array.from(item.children).filter(el => el.tagName === 'PF-V6-TREE-ITEM');
  }

  /**
   * Get top-level items (direct children of tree view)
   */
  #getTopLevelItems() {
    return Array.from(this.children).filter(el => el.tagName === 'PF-V6-TREE-ITEM');
  }

  /**
   * Get the parent tree item, or null if at top level
   */
  #getParentItem(item) {
    const parent = item.parentElement;
    if (!parent) return null;

    // Walk up to find parent tree item
    let current = parent;
    while (current && current !== this) {
      if (current.tagName === 'PF-V6-TREE-ITEM') {
        return current;
      }
      current = current.parentElement;
    }
    return null;
  }

  /**
   * Initialize roving tabindex pattern
   * First item gets tabindex="0", all others get tabindex="-1"
   */
  #initializeTabindex() {
    const topLevelItems = this.#getTopLevelItems();
    if (topLevelItems.length === 0) return;

    const firstItem = topLevelItems[0];
    // Set first item to tabindex="0"
    firstItem.setTabindex(0);

    // Set all other items (including nested) to tabindex="-1"
    const allItems = this.querySelectorAll('pf-v6-tree-item');
    allItems.forEach(item => {
      if (item !== firstItem) {
        item.setTabindex(-1);
      }
    });
  }

  /**
   * Focus an item and manage roving tabindex
   */
  #focusItem(item) {
    if (!item) return;

    // Update tabindex: only focused item should have tabindex="0"
    const allItems = this.querySelectorAll('pf-v6-tree-item');
    allItems.forEach(i => i.setTabindex(-1));
    item.setTabindex(0);

    // Focus the item
    item.focusItem();
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
