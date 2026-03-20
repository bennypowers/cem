import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { live } from 'lit/directives/live.js';

import styles from './pf-v6-switch.css' with { type: 'css' };

/**
 * PatternFly v6 Switch
 *
 * A toggle switch form control.
 *
 * @slot - Default slot for switch label text
 *
 * @fires input - Fired when the switch is toggled
 * @fires change - Fired when the switch is toggled
 */
@customElement('pf-v6-switch')
export class PfV6Switch extends LitElement {
  static readonly formAssociated = true;

  static styles = styles;

  #internals = this.attachInternals();

  @property({ type: Boolean, reflect: true })
  accessor checked = false;

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ reflect: true, attribute: 'label-position' })
  accessor labelPosition?: 'start';

  render() {
    return html`
      <label id="switch-label">
        <input type="checkbox"
               id="switch-input"
               role="switch"
               .checked=${live(this.checked)}
               ?disabled=${this.disabled}
               @input=${this.#onInput}>
        <span id="switch-toggle"
              aria-hidden="true">
          <span class="switch-toggle-icon">
            <svg width="10"
                 height="10"
                 viewBox="0 0 512 512"
                 fill="currentColor">
              <path d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"/>
            </svg>
          </span>
        </span>
        <span id="switch-text"><slot></slot></span>
      </label>
    `;
  }

  #onInput() {
    const input = this.shadowRoot?.getElementById('switch-input') as HTMLInputElement | null;
    if (input) {
      this.checked = input.checked;
      this.#internals.setFormValue(input.checked ? 'on' : null);
    }
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-switch': PfV6Switch;
  }
}
