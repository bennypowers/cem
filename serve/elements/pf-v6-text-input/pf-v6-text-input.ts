import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { live } from 'lit/directives/live.js';
import { ifDefined } from 'lit/directives/if-defined.js';

import styles from './pf-v6-text-input.css' with { type: 'css' };

/**
 * PatternFly v6 Text Input
 *
 * A form-associated text input control.
 *
 * @fires input - Fired when the input value changes
 * @fires change - Fired when the input value is committed
 */
@customElement('pf-v6-text-input')
export class PfV6TextInput extends LitElement {
  static readonly formAssociated = true;

  static shadowRootOptions: ShadowRootInit = {
    ...LitElement.shadowRootOptions,
    delegatesFocus: true,
  };

  static styles = styles;

  #internals = this.attachInternals();

  @property({ reflect: true })
  accessor type: string = 'text';

  @property()
  accessor value: string = '';

  @property()
  accessor placeholder?: string;

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ type: Boolean, reflect: true })
  accessor readonly = false;

  @property({ type: Boolean, reflect: true })
  accessor invalid = false;

  @property()
  accessor min?: string;

  @property()
  accessor max?: string;

  @property()
  accessor step?: string;

  get valueAsNumber(): number {
    const input = this.shadowRoot?.getElementById('text-input') as HTMLInputElement | null;
    return input?.valueAsNumber ?? NaN;
  }

  render() {
    return html`
      <input id="text-input"
             .type=${this.type}
             .value=${live(this.value)}
             placeholder=${ifDefined(this.placeholder)}
             ?disabled=${this.disabled}
             ?readonly=${this.readonly}
             aria-invalid=${ifDefined(this.invalid ? 'true' : undefined)}
             min=${ifDefined(this.type === 'number' ? this.min : undefined)}
             max=${ifDefined(this.type === 'number' ? this.max : undefined)}
             step=${ifDefined(this.type === 'number' ? this.step : undefined)}
             @input=${this.#onInput}
             @change=${this.#onChange}>
    `;
  }

  #onInput() {
    const input = this.shadowRoot?.getElementById('text-input') as HTMLInputElement | null;
    if (input) {
      this.value = input.value;
      this.#internals.setFormValue(input.value);
      this.dispatchEvent(new Event('input', { bubbles: true }));
    }
  }

  #onChange() {
    const input = this.shadowRoot?.getElementById('text-input') as HTMLInputElement | null;
    if (input) {
      this.value = input.value;
      this.#internals.setFormValue(input.value);
      this.dispatchEvent(new Event('change', { bubbles: true }));
    }
  }

  /**
   * Sets the accessible label on the internal input.
   * Called by pf-v6-form-label when label text changes.
   */
  setAccessibleLabel(labelText: string) {
    const input = this.shadowRoot?.getElementById('text-input') as HTMLInputElement | null;
    if (input && labelText) {
      input.setAttribute('aria-label', labelText);
    }
  }

  select() {
    const input = this.shadowRoot?.getElementById('text-input') as HTMLInputElement | null;
    input?.select();
  }

  checkValidity(): boolean {
    const input = this.shadowRoot?.getElementById('text-input') as HTMLInputElement | null;
    return input?.checkValidity() ?? true;
  }

  reportValidity(): boolean {
    const input = this.shadowRoot?.getElementById('text-input') as HTMLInputElement | null;
    return input?.reportValidity() ?? true;
  }

  setCustomValidity(message: string) {
    const input = this.shadowRoot?.getElementById('text-input') as HTMLInputElement | null;
    input?.setCustomValidity(message);
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-text-input': PfV6TextInput;
  }
}
