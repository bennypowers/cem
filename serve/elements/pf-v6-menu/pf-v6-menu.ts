import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-menu.css' with { type: 'css' };

/**
 * PatternFly v6 Menu component
 *
 * Container for pf-v6-menu-item elements. Implements keyboard navigation
 * using roving tabindex pattern.
 *
 * @slot - Menu items (pf-v6-menu-item elements)
 *
 * @attr {string} label - Accessible label for the menu
 */
@customElement('pf-v6-menu')
export class PfV6Menu extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  @property()
  accessor label = '';

  constructor() {
    super();
    this.#internals.role = 'menu';
  }

  render() {
    return html`<slot @slotchange=${this.#onSlotChange}></slot>`;
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('label')) {
      this.#internals.ariaLabel = this.label || null;
    }
  }

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('keydown', this.#onKeydown);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('keydown', this.#onKeydown);
  }

  #onSlotChange() {
    requestAnimationFrame(() => {
      this.#initializeTabindex();
    });
  }

  #onKeydown = (event: KeyboardEvent) => {
    const { key, target } = event;

    if ((target as Element).tagName !== 'PF-V6-MENU-ITEM') return;

    const items = this.#getMenuItems();
    const currentIndex = items.indexOf(target as Element);

    switch (key) {
      case 'ArrowDown':
        event.preventDefault();
        this.#focusItem(items[currentIndex < items.length - 1 ? currentIndex + 1 : 0]);
        break;

      case 'ArrowUp':
        event.preventDefault();
        this.#focusItem(items[currentIndex > 0 ? currentIndex - 1 : items.length - 1]);
        break;

      case 'Home':
        event.preventDefault();
        if (items.length > 0) this.#focusItem(items[0]);
        break;

      case 'End':
        event.preventDefault();
        if (items.length > 0) this.#focusItem(items[items.length - 1]);
        break;
    }
  };

  #getMenuItems(): Element[] {
    const slot = this.shadowRoot?.querySelector('slot') as HTMLSlotElement | null;
    if (!slot) return [];
    return slot.assignedElements().filter(
      el => el.tagName === 'PF-V6-MENU-ITEM' && !(el as HTMLElement & { disabled?: boolean }).disabled
    );
  }

  #initializeTabindex() {
    const items = this.#getMenuItems();
    if (items.length === 0) return;
    items.forEach((item, index) => {
      item.setAttribute('tabindex', index === 0 ? '0' : '-1');
    });
  }

  #focusItem(item: Element) {
    if (!item) return;
    const items = this.#getMenuItems();
    items.forEach(i => {
      i.setAttribute('tabindex', i === item ? '0' : '-1');
    });
    (item as HTMLElement).focus();
  }

  /**
   * Focus the first menu item.
   * Called when dropdown opens.
   */
  focusFirstItem() {
    const items = this.#getMenuItems();
    if (items.length > 0) {
      this.#focusItem(items[0]);
    }
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-menu': PfV6Menu;
  }
}
