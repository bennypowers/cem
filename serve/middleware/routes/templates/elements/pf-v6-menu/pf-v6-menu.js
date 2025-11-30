import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Menu component
 *
 * Container for pf-v6-menu-item elements. Implements keyboard navigation
 * using roving tabindex pattern.
 *
 * @slot - Menu items (pf-v6-menu-item elements)
 *
 * @attr {string} label - Accessible label for the menu (sets aria-label via ElementInternals)
 *
 * @customElement pf-v6-menu
 */
class PfV6Menu extends CemElement {
  static is = 'pf-v6-menu';

  static observedAttributes = ['label'];

  #internals = this.attachInternals();
  #slot;
  #initialized = false;

  get label() {
    return this.getAttribute('label') || '';
  }

  set label(value) {
    if (value) {
      this.setAttribute('label', value);
    } else {
      this.removeAttribute('label');
    }
  }

  async afterTemplateLoaded() {
    this.#slot = this.shadowRoot.querySelector('slot');
    if (!this.#slot) return;

    // Set role via ElementInternals
    this.#internals.role = 'menu';
    this.#updateAriaLabel();

    // Listen for keyboard events
    this.addEventListener('keydown', this.#handleKeydown);

    // Listen for select events from menu items
    this.addEventListener('select', this.#handleItemSelect);

    // Initialize roving tabindex after a delay to allow menu items to be added
    requestAnimationFrame(() => {
      this.#initializeTabindex();
      this.#initialized = true;
    });
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.shadowRoot?.firstChild || oldValue === newValue) {
      return;
    }

    if (name === 'label') {
      this.#updateAriaLabel();
    }
  }

  #updateAriaLabel() {
    if (this.label) {
      this.#internals.ariaLabel = this.label;
    } else {
      this.#internals.ariaLabel = null;
    }
  }

  #handleItemSelect = (event) => {
    // Just let the event bubble naturally
    // Event already has value and checked properties
  };

  #handleKeydown = (event) => {
    const target = event.target;

    // Only handle if target is a menu item
    if (target.tagName !== 'PF-V6-MENU-ITEM') return;

    const items = this.#getMenuItems();
    const currentIndex = items.indexOf(target);

    switch (event.key) {
      case 'ArrowDown':
        event.preventDefault();
        if (currentIndex < items.length - 1) {
          this.#focusItem(items[currentIndex + 1]);
        } else {
          // Wrap to first item
          this.#focusItem(items[0]);
        }
        break;

      case 'ArrowUp':
        event.preventDefault();
        if (currentIndex > 0) {
          this.#focusItem(items[currentIndex - 1]);
        } else {
          // Wrap to last item
          this.#focusItem(items[items.length - 1]);
        }
        break;

      case 'Home':
        event.preventDefault();
        if (items.length > 0) {
          this.#focusItem(items[0]);
        }
        break;

      case 'End':
        event.preventDefault();
        if (items.length > 0) {
          this.#focusItem(items[items.length - 1]);
        }
        break;

      // Space and Enter are handled by menu items themselves
    }
  };

  /**
   * Get all menu items in the default slot
   */
  #getMenuItems() {
    if (!this.#slot) return [];

    const assigned = this.#slot.assignedElements();
    return assigned.filter(el => el.tagName === 'PF-V6-MENU-ITEM' && !el.disabled);
  }

  /**
   * Initialize roving tabindex pattern
   * First item gets tabindex="0", all others get tabindex="-1"
   */
  #initializeTabindex() {
    const items = this.#getMenuItems();
    if (items.length === 0) return;

    items.forEach((item, index) => {
      item.setAttribute('tabindex', index === 0 ? '0' : '-1');
    });
  }

  /**
   * Focus a menu item and update roving tabindex
   */
  #focusItem(item) {
    if (!item) return;

    const items = this.#getMenuItems();

    // Update tabindex: only focused item should have tabindex="0"
    items.forEach(i => {
      i.setAttribute('tabindex', i === item ? '0' : '-1');
    });

    // Focus the item
    item.focus();
  }

  /**
   * Focus the first menu item
   * Called when dropdown opens
   */
  focusFirstItem() {
    const items = this.#getMenuItems();
    if (items.length > 0) {
      this.#focusItem(items[0]);
    }
  }

  disconnectedCallback() {
    this.removeEventListener('keydown', this.#handleKeydown);
    this.removeEventListener('select', this.#handleItemSelect);
  }

  static {
    customElements.define(this.is, this);
  }
}
