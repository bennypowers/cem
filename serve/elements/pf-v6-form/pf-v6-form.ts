import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-form.css' with { type: 'css' };

/**
 * PatternFly v6 Form Component
 *
 * Wraps form controls in a semantic `<form>` element with PatternFly styling.
 * Supports standard form submission and validation APIs.
 *
 * @slot - Default slot for form controls
 *
 * @attr {boolean} horizontal - Enables horizontal layout with grid
 * @fires submit - Forwarded from internal `<form>` element
 */
@customElement('pf-v6-form')
export class PfV6Form extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor horizontal = false;

  render() {
    return html`
      <form id="form"
            @submit=${this.#onSubmit}>
        <slot></slot>
      </form>
    `;
  }

  #onSubmit(e: SubmitEvent) {
    this.dispatchEvent(new SubmitEvent('submit', {
      bubbles: true,
      cancelable: true,
      submitter: e.submitter,
    }));
  }

  /** Submits the form programmatically */
  submit() {
    this.#formElement?.submit();
  }

  /**
   * Requests form submission with validation
   * @param submitter - Optional submitter element
   */
  requestSubmit(submitter?: HTMLElement) {
    return this.#formElement?.requestSubmit(submitter);
  }

  /** Resets the form to default values */
  reset() {
    this.#formElement?.reset();
  }

  /** Checks form validity */
  checkValidity(): boolean {
    return this.#formElement?.checkValidity() ?? true;
  }

  /** Reports form validity (shows validation messages) */
  reportValidity(): boolean {
    return this.#formElement?.reportValidity() ?? true;
  }

  get #formElement(): HTMLFormElement | null {
    return this.shadowRoot?.getElementById('form') as HTMLFormElement | null;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-form': PfV6Form;
  }
}
