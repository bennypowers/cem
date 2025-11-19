import { CemElement } from '/__cem/cem-element.js';

// Custom event for toggle group item selection
export class ToggleGroupItemSelectEvent extends Event {
  constructor(item, selected, value) {
    super('pfv6-toggle-group-item-select', { bubbles: true, composed: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
}

export class PfToggleGroupItem extends CemElement {
  #button;

  static observedAttributes = ['selected', 'disabled', 'value'];

  async afterTemplateLoaded() {
    this.addEventListener('slotchange', this.#updateSlotVisibility);
    this.#button = this.shadowRoot.getElementById('button');

    // Set up event listeners
    this.#button.addEventListener('click', this.#handleClick);
    this.#button.addEventListener('keydown', this.#handleKeydown);
    this.#button.addEventListener('focus', this.#handleFocus);

    // Initial slot visibility update
    this.#updateSlotVisibility();

    // Initialize aria-checked and tabindex
    this.#updateAriaChecked();
    this.#updateTabindex();
  }

  disconnectedCallback() {
    this.removeEventListener('slotchange', this.#updateSlotVisibility);
    if (this.#button) {
      this.#button.removeEventListener('click', this.#handleClick);
      this.#button.removeEventListener('keydown', this.#handleKeydown);
      this.#button.removeEventListener('focus', this.#handleFocus);
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.#button) return;

    if (name === 'selected') {
      this.#updateAriaChecked();
      this.#updateTabindex();
    } else if (name === 'disabled') {
      this.#updateDisabled();
    }
  }

  #updateAriaChecked() {
    const isSelected = this.hasAttribute('selected');
    this.#button.setAttribute('aria-checked', isSelected.toString());
  }

  #updateTabindex() {
    // Roving tabindex: selected item or first item gets tabindex="0"
    const isSelected = this.hasAttribute('selected');
    const parent = this.parentElement;

    if (!parent) return;

    const items = Array.from(parent.querySelectorAll('pfv6-toggle-group-item'));
    const isFirstItem = items[0] === this;
    const hasSelectedItem = items.some(item => item.hasAttribute('selected'));

    // This item should be tabbable if:
    // - It's selected, OR
    // - It's the first item and nothing is selected
    const shouldBeTabbable = isSelected || (isFirstItem && !hasSelectedItem);
    this.#button.setAttribute('tabindex', shouldBeTabbable ? '0' : '-1');
  }

  #updateDisabled() {
    const isDisabled = this.hasAttribute('disabled');
    this.#button.disabled = isDisabled;
  }

  #handleClick = (event) => {
    event.preventDefault();
    if (this.hasAttribute('disabled')) return;

    this.#selectItem();
  }

  #handleKeydown = (event) => {
    if (this.hasAttribute('disabled')) return;

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

    const items = Array.from(parent.querySelectorAll('pfv6-toggle-group-item'))
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
      targetItem.shadowRoot.getElementById('button').focus();
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

    const items = Array.from(parent.querySelectorAll('pfv6-toggle-group-item'))
      .filter(item => !item.hasAttribute('disabled'));

    const targetItem = toStart ? items[0] : items[items.length - 1];
    if (targetItem) {
      targetItem.shadowRoot.getElementById('button').focus();
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

    const items = Array.from(parent.querySelectorAll('pfv6-toggle-group-item'));

    // Set all items to tabindex="-1" except this one
    items.forEach(item => {
      const button = item.shadowRoot?.getElementById('button');
      if (button) {
        button.setAttribute('tabindex', item === this ? '0' : '-1');
      }
    });
  }

  #updateSlotVisibility() {
    const iconSlot = this.shadowRoot.querySelector('slot[name="icon"]');
    const iconEndSlot = this.shadowRoot.querySelector('slot[name="icon-end"]');
    const textSlot = this.shadowRoot.querySelector('slot:not([name])');

    const iconStart = this.shadowRoot.getElementById('icon-start');
    const iconEnd = this.shadowRoot.getElementById('icon-end');
    const text = this.shadowRoot.getElementById('text');

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
}

customElements.define('pfv6-toggle-group-item', PfToggleGroupItem);
