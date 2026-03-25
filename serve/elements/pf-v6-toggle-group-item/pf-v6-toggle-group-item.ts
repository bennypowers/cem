import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-toggle-group-item.css' with { type: 'css' };

/**
 * Custom event for toggle group item selection
 */
export class ToggleGroupItemSelectEvent extends Event {
  item: Element;
  selected: boolean;
  value: string | null;
  constructor(item: Element, selected: boolean, value: string | null) {
    super('pf-v6-toggle-group-item-select', { bubbles: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
}

/**
 * PatternFly v6 Toggle Group Item
 *
 * An individual toggle button within a toggle group.
 *
 * @slot - Default slot for button text
 * @slot icon - Leading icon
 * @slot icon-end - Trailing icon
 *
 * @fires pf-v6-toggle-group-item-select - Fires when item is selected
 */
@customElement('pf-v6-toggle-group-item')
export class PfV6ToggleGroupItem extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  @property({ type: Boolean, reflect: true })
  accessor selected = false;

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ reflect: true })
  accessor value?: string;

  constructor() {
    super();
    this.#internals.role = 'radio';
    if (!this.hasAttribute('tabindex')) {
      this.setAttribute('tabindex', '-1');
    }
  }

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('click', this.#handleClick);
    this.addEventListener('keydown', this.#handleKeydown);
    this.addEventListener('focus', this.#handleFocus);

    queueMicrotask(() => {
      const parent = this.parentElement;
      if (parent) {
        const items = Array.from(parent.querySelectorAll('pf-v6-toggle-group-item'));
        const isFirstItem = items[0] === this;
        if (this.selected || isFirstItem) {
          this.#updateRovingTabindex();
        }
      }
    });
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('click', this.#handleClick);
    this.removeEventListener('keydown', this.#handleKeydown);
    this.removeEventListener('focus', this.#handleFocus);
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('selected')) {
      this.#internals.ariaChecked = String(this.selected);
      this.#updateRovingTabindex();
    }
    if (changed.has('disabled')) {
      this.#internals.ariaDisabled = String(this.disabled);
      this.#updateTabindex();
    }
  }

  render() {
    return html`
      <span id="wrapper">
        <span id="icon-start"><slot name="icon"></slot></span>
        <span id="text"><slot></slot></span>
        <span id="icon-end"><slot name="icon-end"></slot></span>
      </span>
    `;
  }

  #updateTabindex() {
    if (this.disabled) {
      this.setAttribute('tabindex', '-1');
      return;
    }
    const parent = this.parentElement;
    if (!parent) return;
    const items = Array.from(parent.querySelectorAll('pf-v6-toggle-group-item'));
    const isFirstItem = items[0] === this;
    const hasSelectedItem = items.some(item => item.hasAttribute('selected'));
    const shouldBeTabbable = this.selected || (isFirstItem && !hasSelectedItem);
    this.setAttribute('tabindex', shouldBeTabbable ? '0' : '-1');
  }

  #handleClick = (event: Event) => {
    if (this.disabled) {
      event.preventDefault();
      event.stopImmediatePropagation();
      return;
    }
    this.focus();
    this.#selectItem();
  };

  #handleKeydown = (event: KeyboardEvent) => {
    if (this.disabled) {
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
  };

  #handleFocus = () => {
    this.#updateRovingTabindex();
  };

  #selectItem() {
    if (!this.selected) {
      this.selected = true;
      this.dispatchEvent(new ToggleGroupItemSelectEvent(
        this,
        true,
        this.getAttribute('value')
      ));
    }
  }

  #focusAndSelect(item: PfV6ToggleGroupItem) {
    item.focus();
    item.selected = true;
    item.dispatchEvent(new ToggleGroupItemSelectEvent(
      item,
      true,
      item.getAttribute('value')
    ));
  }

  #navigateItems(direction: number) {
    const parent = this.parentElement;
    if (!parent) return;
    const items = Array.from(parent.querySelectorAll<PfV6ToggleGroupItem>('pf-v6-toggle-group-item'))
      .filter(item => !item.disabled);
    const currentIndex = items.indexOf(this);
    if (currentIndex === -1) return;
    let newIndex = currentIndex + direction;
    if (newIndex < 0) newIndex = items.length - 1;
    else if (newIndex >= items.length) newIndex = 0;
    const targetItem = items[newIndex];
    if (targetItem) this.#focusAndSelect(targetItem);
  }

  #navigateToEnd(toStart: boolean) {
    const parent = this.parentElement;
    if (!parent) return;
    const items = Array.from(parent.querySelectorAll<PfV6ToggleGroupItem>('pf-v6-toggle-group-item'))
      .filter(item => !item.disabled);
    const targetItem = toStart ? items[0] : items[items.length - 1];
    if (targetItem) this.#focusAndSelect(targetItem);
  }

  #updateRovingTabindex() {
    const parent = this.parentElement;
    if (!parent) return;
    const items = Array.from(parent.querySelectorAll('pf-v6-toggle-group-item'));
    items.forEach(item => {
      item.setAttribute('tabindex', item === this ? '0' : '-1');
    });
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-toggle-group-item': PfV6ToggleGroupItem;
  }
}
