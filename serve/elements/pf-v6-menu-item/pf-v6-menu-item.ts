import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-menu-item.css' with { type: 'css' };

/**
 * Custom event fired when menu item is selected/checked
 */
export class PfMenuItemSelectEvent extends Event {
  value: string;
  checked: boolean;
  constructor(value: string, checked: boolean) {
    super('select', { bubbles: true });
    this.value = value;
    this.checked = checked;
  }
}

/**
 * PatternFly v6 Menu Item
 *
 * @slot - Default slot for item text
 *
 * @fires select - Fires when item is selected/checked
 */
@customElement('pf-v6-menu-item')
export class PfV6MenuItem extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ type: Boolean, reflect: true })
  accessor checked = false;

  @property({ reflect: true })
  accessor variant: 'default' | 'checkbox' = 'default';

  @property({ reflect: true })
  accessor value = '';

  @property({ reflect: true })
  accessor description?: string;

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('click', this.#handleClick);
    this.addEventListener('keydown', this.#handleKeydown);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('click', this.#handleClick);
    this.removeEventListener('keydown', this.#handleKeydown);
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('variant')) {
      this.#internals.role = this.variant === 'checkbox' ? 'menuitemcheckbox' : 'menuitem';
    }
    if (changed.has('checked') || changed.has('variant')) {
      this.#internals.ariaChecked = this.variant === 'checkbox' ? String(this.checked) : null;
    }
    if (changed.has('disabled')) {
      this.#internals.ariaDisabled = this.disabled ? 'true' : null;
      this.setAttribute('tabindex', '-1');
    }
  }

  render() {
    return html`
      ${this.variant === 'checkbox' ? html`
        <span id="check"
              class="pf-v6-c-check pf-m-standalone">
          <input id="input"
                 class="pf-v6-c-check__input"
                 role="presentation"
                 tabindex="-1"
                 type="checkbox"
                 .checked=${this.checked}
                 ?disabled=${this.disabled}>
        </span>
      ` : nothing}
      <span id="text"><slot></slot></span>
      ${this.description ? html`
        <span id="description">${this.description}</span>
      ` : nothing}
    `;
  }

  #handleClick = (event: Event) => {
    if (this.disabled) {
      event.preventDefault();
      event.stopImmediatePropagation();
      return;
    }
    if (this.variant === 'checkbox') {
      this.checked = !this.checked;
    }
    this.dispatchEvent(new PfMenuItemSelectEvent(this.value, this.checked));
  };

  #handleKeydown = (event: KeyboardEvent) => {
    if (this.disabled) return;
    if (event.key === ' ' || event.key === 'Enter') {
      event.preventDefault();
      this.#handleClick(event);
    }
  };
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-menu-item': PfV6MenuItem;
  }
}
