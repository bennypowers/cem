import { CemElement } from '/__cem/cem-element.js';

// Custom event for toggle group item selection
export class ToggleGroupItemSelectEvent extends Event {
  constructor(item, selected, value) {
    super('pf-v6-toggle-group-item-select', { bubbles: true, composed: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
}

/**
 * @customElement pf-v6-toggle-group-item
 */
export class PfToggleGroupItem extends CemElement {
  static is = 'pf-v6-toggle-group-item';

  static observedAttributes = ['selected', 'disabled', 'value'];

  #internals = this.attachInternals();
  #wrapper;

  #$ = id => this.shadowRoot.getElementById(id);

  constructor() {
    super();
    // Set role immediately
    this.#internals.role = 'radio';

    // Set initial tabindex to -1 (will be updated in afterTemplateLoaded)
    if (!this.hasAttribute('tabindex')) {
      this.setAttribute('tabindex', '-1');
    }
  }

  connectedCallback() {
    super.connectedCallback?.();

    // Update tabindex when connected to ensure proper roving tabindex
    // Use setTimeout to ensure all siblings are in the DOM
    setTimeout(() => {
      const isSelected = this.hasAttribute('selected');
      const parent = this.parentElement;
      if (parent) {
        const items = Array.from(parent.querySelectorAll('pf-v6-toggle-group-item'));
        const isFirstItem = items[0] === this;
        if (isSelected || isFirstItem) {
          this.#updateRovingTabindex();
        }
      }
    }, 0);
  }

  async afterTemplateLoaded() {
    this.addEventListener('slotchange', this.#updateSlotVisibility);
    this.#wrapper = this.#$('wrapper');

    // Event listeners on HOST
    this.addEventListener('click', this.#handleClick);
    this.addEventListener('keydown', this.#handleKeydown);
    this.addEventListener('focus', this.#handleFocus);

    // Initial state
    this.#updateSlotVisibility();
    this.#updateAriaChecked();

    // Set initial tabindex immediately
    this.#updateTabindex();
  }

  disconnectedCallback() {
    this.removeEventListener('slotchange', this.#updateSlotVisibility);
    this.removeEventListener('click', this.#handleClick);
    this.removeEventListener('keydown', this.#handleKeydown);
    this.removeEventListener('focus', this.#handleFocus);
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'selected') {
      this.#updateAriaChecked();
      // When selection changes, update tabindex for all siblings
      this.#updateRovingTabindex();
    } else if (name === 'disabled') {
      this.#updateDisabled();
    }
  }

  #updateAriaChecked() {
    const isSelected = this.hasAttribute('selected');
    this.#internals.ariaChecked = isSelected.toString();
  }

  #updateTabindex() {
    // Roving tabindex on HOST
    const isSelected = this.hasAttribute('selected');
    const parent = this.parentElement;

    if (!parent) return;

    const items = Array.from(parent.querySelectorAll('pf-v6-toggle-group-item'));
    const isFirstItem = items[0] === this;
    const hasSelectedItem = items.some(item => item.hasAttribute('selected'));

    // This item should be tabbable if:
    // - It's selected, OR
    // - It's the first item and nothing is selected
    const shouldBeTabbable = isSelected || (isFirstItem && !hasSelectedItem);
    this.setAttribute('tabindex', shouldBeTabbable ? '0' : '-1');
  }

  #updateDisabled() {
    const isDisabled = this.hasAttribute('disabled');
    this.#internals.ariaDisabled = isDisabled ? 'true' : null;

    if (isDisabled) {
      this.setAttribute('tabindex', '-1');
    } else {
      this.#updateTabindex();
    }
  }

  #handleClick = (event) => {
    if (this.hasAttribute('disabled')) {
      event.preventDefault();
      event.stopImmediatePropagation();
      return;
    }

    // Focus this item first
    this.focus();
    this.#selectItem();
  }

  #handleKeydown = (event) => {
    if (this.hasAttribute('disabled')) {
      event.preventDefault();
      return;
    }

    switch (event.key) {
      case ' ':
      case 'Enter':
        event.preventDefault();
        this.#selectItem();
        break;

      case 'ArrowLeft':
      case 'ArrowRight':
        event.preventDefault();
        this.#navigateItems(event.key === 'ArrowLeft' ? -1 : 1);
        break;

      case 'Home':
        event.preventDefault();
        this.#navigateToEnd(true);
        break;

      case 'End':
        event.preventDefault();
        this.#navigateToEnd(false);
        break;
    }
  }

  #handleFocus = () => {
    // When this item receives focus, update roving tabindex
    this.#updateRovingTabindex();
  }

  #selectItem() {
    // Always select this item (parent will handle deselecting siblings)
    if (!this.hasAttribute('selected')) {
      this.setAttribute('selected', '');

      // Dispatch event to parent toggle group
      this.dispatchEvent(new ToggleGroupItemSelectEvent(
        this,
        true,
        this.getAttribute('value')
      ));
    }
  }

  #navigateItems(direction) {
    const parent = this.parentElement;
    if (!parent) return;

    const items = Array.from(parent.querySelectorAll('pf-v6-toggle-group-item'))
      .filter(item => !item.hasAttribute('disabled'));

    const currentIndex = items.indexOf(this);
    if (currentIndex === -1) return;

    let newIndex = currentIndex + direction;

    // Wrap around
    if (newIndex < 0) {
      newIndex = items.length - 1;
    } else if (newIndex >= items.length) {
      newIndex = 0;
    }

    const targetItem = items[newIndex];
    if (targetItem) {
      targetItem.focus(); // Focus HOST
      // Select the item when navigating (radio group behavior)
      targetItem.setAttribute('selected', '');
      targetItem.dispatchEvent(new ToggleGroupItemSelectEvent(
        targetItem,
        true,
        targetItem.getAttribute('value')
      ));
    }
  }

  #navigateToEnd(toStart) {
    const parent = this.parentElement;
    if (!parent) return;

    const items = Array.from(parent.querySelectorAll('pf-v6-toggle-group-item'))
      .filter(item => !item.hasAttribute('disabled'));

    const targetItem = toStart ? items[0] : items[items.length - 1];
    if (targetItem) {
      targetItem.focus(); // Focus HOST
      // Select the item when navigating (radio group behavior)
      targetItem.setAttribute('selected', '');
      targetItem.dispatchEvent(new ToggleGroupItemSelectEvent(
        targetItem,
        true,
        targetItem.getAttribute('value')
      ));
    }
  }

  #updateRovingTabindex() {
    const parent = this.parentElement;
    if (!parent) return;

    const items = Array.from(parent.querySelectorAll('pf-v6-toggle-group-item'));

    // Set all items to tabindex="-1" except this one
    items.forEach(item => {
      item.setAttribute('tabindex', item === this ? '0' : '-1');
    });
  }

  #updateSlotVisibility() {
    const iconSlot = this.shadowRoot.querySelector('slot[name="icon"]');
    const iconEndSlot = this.shadowRoot.querySelector('slot[name="icon-end"]');
    const textSlot = this.shadowRoot.querySelector('slot:not([name])');

    const iconStart = this.shadowRoot.getElementById('icon-start');
    const iconEnd = this.shadowRoot.getElementById('icon-end');
    const text = this.#$('text');

    // Hide icon wrappers if slots have no assigned nodes
    if (iconStart && iconSlot) {
      const hasContent = iconSlot.assignedNodes().length > 0;
      iconStart.hidden = !hasContent;
    }

    if (iconEnd && iconEndSlot) {
      const hasContent = iconEndSlot.assignedNodes().length > 0;
      iconEnd.hidden = !hasContent;
    }

    // Text slot - check for non-whitespace content
    if (text && textSlot) {
      const nodes = textSlot.assignedNodes();
      const hasContent = nodes.some(node => {
        if (node.nodeType === Node.TEXT_NODE) {
          return node.textContent.trim().length > 0;
        }
        return true; // Non-text nodes count as content
      });
      text.hidden = !hasContent;
    }
  }

  focus(options) {
    super.focus(options);
  }

  static {
    customElements.define(this.is, this);
  }
}
