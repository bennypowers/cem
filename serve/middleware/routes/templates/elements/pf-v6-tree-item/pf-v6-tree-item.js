import '/__cem/elements/pf-v6-badge/pf-v6-badge.js'
import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event classes for tree item interactions
 */
export class PfTreeItemSelectEvent extends Event {
  constructor() {
    super('select', { bubbles: true });
  }
}

export class PfTreeItemExpandEvent extends Event {
  constructor() {
    super('expand', { bubbles: true });
  }
}

export class PfTreeItemCollapseEvent extends Event {
  constructor() {
    super('collapse', { bubbles: true });
  }
}

/**
 * PatternFly v6 Tree Item
 *
 * Individual tree node that can be nested.
 *
 * @slot - Child tree items
 * @slot label - Item label text
 *
 * @attr {string} label - Item label (alternative to slot)
 * @attr {string} icon - SVG icon markup
 * @attr {string} badge - Badge text/number
 * @attr {boolean} expanded - Whether item is expanded
 * @attr {boolean} has-children - Whether item has children (auto-detected)
 * @attr {boolean} current - Whether item is currently selected
 *
 * @fires select - Fires when item is selected
 * @fires expand - Fires when item is expanded
 * @fires collapse - Fires when item is collapsed
 *
 * @customElement pf-v6-tree-item
 */
export class PfV6TreeItem extends CemElement {
  static is = 'pf-v6-tree-item';

  static observedAttributes = [
    'label',
    'icon',
    'badge',
    'expanded',
    'current',
  ];

  #$ = id => this.shadowRoot.getElementById(id);
  #item;
  #toggle;
  #node;
  #icon;
  #badge;
  #children;
  #childrenSlot;

  get label() { return this.getAttribute('label'); }
  set label(value) {
    if (value) {
      this.setAttribute('label', value);
    } else {
      this.removeAttribute('label');
    }
  }

  get icon() { return this.getAttribute('icon'); }
  set icon(value) {
    if (value) {
      this.setAttribute('icon', value);
    } else {
      this.removeAttribute('icon');
    }
  }

  get badge() { return this.getAttribute('badge'); }
  set badge(value) {
    if (value) {
      this.setAttribute('badge', value);
    } else {
      this.removeAttribute('badge');
    }
  }

  get expanded() { return this.hasAttribute('expanded'); }
  set expanded(value) { this.toggleAttribute('expanded', !!value); }

  get current() { return this.hasAttribute('current'); }
  set current(value) { this.toggleAttribute('current', !!value); }

  get hasChildren() { return this.hasAttribute('has-children'); }

  // Track if has-children was explicitly set by user
  #explicitHasChildren = false;

  afterTemplateLoaded() {
    this.#item = this.#$('item');
    this.#toggle = this.#$('toggle');
    this.#node = this.#$('node');
    this.#icon = this.#$('icon');
    this.#badge = this.#$('badge');
    this.#children = this.#$('children');

    if (!this.#item || !this.#toggle || !this.#node) return;

    // Check if has-children was explicitly set before we started observing
    this.#explicitHasChildren = this.hasAttribute('has-children');

    // Set up ARIA
    this.#updateAriaAttributes();

    // Toggle handler
    this.#toggle.addEventListener('click', (e) => {
      e.stopPropagation();
      this.toggle();
    });

    // Node click handler
    this.#node.addEventListener('click', () => {
      if (this.hasChildren) {
        this.toggle();
      }
      this.select();
    });

    // Watch for children via slot
    this.#childrenSlot = this.#children.querySelector('slot');
    this.#childrenSlot?.addEventListener('slotchange', () => this.#updateHasChildren());
    this.#updateHasChildren();

    // Update icon and badge
    this.#updateIcon();
    this.#updateBadge();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue === newValue) return;

    switch (name) {
      case 'expanded':
        this.#updateAriaAttributes();
        if (newValue !== null) {
          this.dispatchEvent(new PfTreeItemExpandEvent());
        } else {
          this.dispatchEvent(new PfTreeItemCollapseEvent());
        }
        break;

      case 'current':
        if (this.#node) {
          this.#node.classList.toggle('current', newValue !== null);
        }
        this.#updateAriaAttributes();
        break;

      case 'icon':
        this.#updateIcon();
        break;

      case 'badge':
        this.#updateBadge();
        break;

      case 'label':
        // Label is handled via slot, but we update aria-label
        this.#updateAriaAttributes();
        break;
    }
  }

  #updateAriaAttributes() {
    if (!this.#item) return;

    // Set aria-expanded if has children
    if (this.hasChildren) {
      this.#item.setAttribute('aria-expanded', String(this.expanded));
    } else {
      this.#item.removeAttribute('aria-expanded');
    }

    // Set aria-selected
    this.#item.setAttribute('aria-selected', String(this.current));

    // Set aria-label from label attribute if present
    if (this.label) {
      this.#item.setAttribute('aria-label', this.label);
    }
  }

  #updateHasChildren() {
    if (!this.#childrenSlot) return;

    // If has-children was explicitly set, don't auto-update it
    // This allows lazy rendering where children aren't in DOM until expanded
    if (this.#explicitHasChildren) return;

    const children = this.#childrenSlot.assignedElements();
    const hasChildren = children.length > 0;

    this.toggleAttribute('has-children', hasChildren);
    this.#updateAriaAttributes();
  }

  #updateIcon() {
    if (!this.#icon) return;

    const iconMarkup = this.icon;
    if (iconMarkup) {
      this.#icon.innerHTML = iconMarkup;
    } else {
      this.#icon.innerHTML = '';
    }
  }

  #updateBadge() {
    if (!this.#badge) return;

    const badgeText = this.badge;
    if (badgeText) {
      // Check if badge slot has content
      const badgeSlot = this.#badge.querySelector('slot[name="badge"]');
      if (badgeSlot && badgeSlot.assignedNodes().length === 0) {
        // No slotted content - set text content directly on badge element
        this.#badge.textContent = badgeText;
      }
    } else {
      // Clear badge if no attribute and no slotted content
      const badgeSlot = this.#badge.querySelector('slot[name="badge"]');
      if (badgeSlot && badgeSlot.assignedNodes().length === 0) {
        this.#badge.textContent = '';
      }
    }
  }

  /**
   * Toggle expanded state
   */
  toggle() {
    if (!this.hasChildren) return;
    this.expanded = !this.expanded;
  }

  /**
   * Expand the item
   */
  expand() {
    if (!this.hasChildren) return;
    this.expanded = true;
  }

  /**
   * Collapse the item
   */
  collapse() {
    if (!this.hasChildren) return;
    this.expanded = false;
  }

  /**
   * Select this item
   */
  select() {
    this.current = true;
    this.dispatchEvent(new PfTreeItemSelectEvent());
  }

  /**
   * Deselect this item
   */
  deselect() {
    this.current = false;
  }

  /**
   * Set tabindex for roving tabindex pattern
   */
  setTabindex(value) {
    if (!this.#item) return;
    this.#item.setAttribute('tabindex', String(value));
  }

  /**
   * Focus this item
   */
  focusItem() {
    if (!this.#item) return;
    this.#item.focus();
  }

  static {
    customElements.define(this.is, this);
  }
}
