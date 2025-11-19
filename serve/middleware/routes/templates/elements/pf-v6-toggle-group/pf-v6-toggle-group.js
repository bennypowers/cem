import { CemElement } from '/__cem/cem-element.js';

// Custom event for toggle group selection changes
export class ToggleGroupChangeEvent extends Event {
  constructor(item, selected, value) {
    super('pf-v6-toggle-group-change', { bubbles: true, composed: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
}

export class PfToggleGroup extends CemElement {
  #container;

  static observedAttributes = ['aria-label', 'compact'];

  async afterTemplateLoaded() {
    this.#container = this.shadowRoot.getElementById('container');
    this.#updateAriaLabel();

    // Listen for selection events from items
    this.addEventListener('pf-v6-toggle-group-item-select', this.#handleItemSelect);
  }

  disconnectedCallback() {
    this.removeEventListener('pf-v6-toggle-group-item-select', this.#handleItemSelect);
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.#container) return;

    if (name === 'aria-label') {
      this.#updateAriaLabel();
    }
  }

  #updateAriaLabel() {
    const label = this.getAttribute('aria-label');
    if (label) {
      this.#container.setAttribute('aria-label', label);
    } else {
      this.#container.removeAttribute('aria-label');
    }
  }

  #handleItemSelect = (event) => {
    const selectedItem = event.item;
    const isNowSelected = event.selected;

    // Single-select mode: deselect all other items when one is selected
    if (isNowSelected) {
      const items = this.querySelectorAll('pf-v6-toggle-group-item');
      items.forEach(item => {
        if (item !== selectedItem && item.hasAttribute('selected')) {
          item.removeAttribute('selected');
        }
      });
    }

    // Dispatch change event from the group
    this.dispatchEvent(new ToggleGroupChangeEvent(
      selectedItem,
      isNowSelected,
      selectedItem.getAttribute('value')
    ));
  }
}

customElements.define('pf-v6-toggle-group', PfToggleGroup);
