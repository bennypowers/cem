import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { live } from 'lit/directives/live.js';
import { ifDefined } from 'lit/directives/if-defined.js';

import styles from './pf-v6-text-input-group.css' with { type: 'css' };

/**
 * PatternFly v6 Text Input Group
 *
 * A text input with optional icons and utility buttons.
 *
 * @slot icon - Leading icon
 * @slot status-icon - Trailing status icon (success, warning, error)
 * @slot utilities - Utility buttons/actions (e.g., eyedropper, clear)
 *
 * @fires input - Native input event
 * @fires change - Native change event
 */
@customElement('pf-v6-text-input-group')
export class PfV6TextInputGroup extends LitElement {
  static readonly formAssociated = true;

  static styles = styles;

  #internals = this.attachInternals();

  @property({ reflect: true })
  accessor value = '';

  @property({ reflect: true })
  accessor type = 'text';

  @property({ reflect: true })
  accessor placeholder?: string;

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ type: Boolean, reflect: true })
  accessor readonly = false;

  @property({ type: Boolean, reflect: true })
  accessor required = false;

  @property({ type: Boolean, reflect: true })
  accessor plain = false;

  @property({ type: Boolean, reflect: true })
  accessor icon = false;

  @property({ reflect: true })
  accessor status?: 'success' | 'warning' | 'error';

  render() {
    return html`
      <div id="main">
        <span id="text">
          <span id="icon"><slot name="icon"></slot></span>
          <span id="status-icon"><slot name="status-icon"></slot></span>
          <input id="input"
                 type=${this.type}
                 .value=${live(this.value)}
                 placeholder=${ifDefined(this.placeholder)}
                 ?disabled=${this.disabled}
                 ?readonly=${this.readonly}
                 ?required=${this.required}
                 @input=${this.#onInput}
                 @change=${this.#onChange}>
        </span>
      </div>
      <div id="utilities">
        <slot name="utilities"></slot>
      </div>
    `;
  }

  focus() {
    this.shadowRoot?.getElementById('input')?.focus();
  }

  blur() {
    (this.shadowRoot?.getElementById('input') as HTMLInputElement | null)?.blur();
  }

  select() {
    (this.shadowRoot?.getElementById('input') as HTMLInputElement | null)?.select();
  }

  #onInput(e: InputEvent) {
    const input = e.target as HTMLInputElement;
    this.value = input.value;
    this.dispatchEvent(new InputEvent('input', {
      bubbles: true,
      data: e.data,
      inputType: e.inputType,
    }));
  }

  #onChange() {
    this.dispatchEvent(new Event('change', { bubbles: true }));
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-text-input-group': PfV6TextInputGroup;
  }
}
