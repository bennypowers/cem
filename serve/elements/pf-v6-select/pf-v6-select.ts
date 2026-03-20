import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-select.css' with { type: 'css' };

/**
 * PatternFly v6 Select
 *
 * A form-associated select control that clones light DOM `<option>` elements
 * into an internal shadow `<select>`.
 *
 * @fires input - Fired when the select value changes
 * @fires change - Fired when the select value is committed
 */
@customElement('pf-v6-select')
export class PfV6Select extends LitElement {
  static readonly formAssociated = true;

  static shadowRootOptions: ShadowRootInit = {
    ...LitElement.shadowRootOptions,
    delegatesFocus: true,
  };

  static styles = styles;

  #internals = this.attachInternals();
  #observer = new MutationObserver(() => this.#cloneOptions());

  @property({ reflect: true })
  accessor value = '';

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ type: Boolean, reflect: true })
  accessor invalid = false;

  render() {
    return html`
      <select id="select"
              ?disabled=${this.disabled}
              aria-invalid=${this.invalid ? 'true' : 'false'}
              @change=${this.#onChange}
              @input=${this.#onInput}></select>
    `;
  }

  firstUpdated() {
    this.#cloneOptions();

    const select = this.#selectEl;
    if (select && this.value) {
      select.value = this.value;
    }

    this.#internals.setFormValue(select?.value || null);

    this.#observer.observe(this, {
      childList: true,
      subtree: true,
      attributes: true,
    });
  }

  get #selectEl(): HTMLSelectElement | null {
    return this.shadowRoot?.getElementById('select') as HTMLSelectElement | null;
  }

  #cloneOptions() {
    const select = this.#selectEl;
    if (!select) return;

    const currentValue = select.value;
    select.innerHTML = '';

    const options = this.querySelectorAll('option');
    options.forEach(option => {
      select.appendChild(option.cloneNode(true));
    });

    select.value = currentValue;
  }

  #onChange = () => {
    const select = this.#selectEl;
    if (!select) return;
    this.value = select.value;
    this.#internals.setFormValue(select.value);
    this.dispatchEvent(new Event('change', { bubbles: true }));
  };

  #onInput = () => {
    const select = this.#selectEl;
    if (!select) return;
    this.value = select.value;
    this.#internals.setFormValue(select.value);
    this.dispatchEvent(new Event('input', { bubbles: true }));
  };

  updated(changed: Map<string, unknown>) {
    if (changed.has('value')) {
      const select = this.#selectEl;
      if (select && select.value !== this.value) {
        select.value = this.value;
      }
      this.#internals.setFormValue(this.value || null);
    }
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.#observer.disconnect();
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-select': PfV6Select;
  }
}
