import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './cem-pf-v6-toggle-group.css' with { type: 'css' };

/**
 * Custom event for toggle group selection changes
 */
export class ToggleGroupChangeEvent extends Event {
  item: Element;
  selected: boolean;
  value: string | null;
  constructor(item: Element, selected: boolean, value: string | null) {
    super('cem-pf-v6-toggle-group-change', { bubbles: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
}

/**
 * PatternFly v6 Toggle Group
 *
 * A group of toggle items that behave like a radio group.
 *
 * @slot - Toggle group items (cem-pf-v6-toggle-group-item elements)
 *
 * @fires cem-pf-v6-toggle-group-change - Fires when selection changes
 */
@customElement('cem-pf-v6-toggle-group')
export class PfV6ToggleGroup extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor compact = false;

  #internals = this.attachInternals();

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('cem-pf-v6-toggle-group-item-select', this.#handleItemSelect as EventListener);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('cem-pf-v6-toggle-group-item-select', this.#handleItemSelect as EventListener);
  }

  render() {
    return html`
      <div id="container"
           role="radiogroup"
           part="container"
           aria-label=${this.getAttribute('aria-label') ?? ''}>
        <slot></slot>
      </div>
    `;
  }

  #handleItemSelect = (event: Event) => {
    const detail = event as ToggleGroupItemSelectEvent;
    const selectedItem = detail.item;
    const isNowSelected = detail.selected;

    if (isNowSelected) {
      const items = this.querySelectorAll('cem-pf-v6-toggle-group-item');
      items.forEach(item => {
        if (item !== selectedItem && item.hasAttribute('selected')) {
          item.removeAttribute('selected');
        }
      });
    }

    this.dispatchEvent(new ToggleGroupChangeEvent(
      selectedItem,
      isNowSelected,
      selectedItem.getAttribute('value')
    ));
  };
}

// Import type for the event from toggle-group-item
interface ToggleGroupItemSelectEvent extends Event {
  item: Element;
  selected: boolean;
  value: string | null;
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-pf-v6-toggle-group': PfV6ToggleGroup;
  }
}
