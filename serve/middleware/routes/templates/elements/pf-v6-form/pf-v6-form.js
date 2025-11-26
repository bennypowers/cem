import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Form Component
 *
 * Wraps form controls in a semantic <form> element with PatternFly styling.
 * Supports standard form submission and validation APIs.
 *
 * @attr {boolean} horizontal - Enables horizontal layout with grid
 * @fires submit - Forwarded from internal <form> element
 * @customElement pf-v6-form
 */
class PfV6Form extends CemElement {
  static is = 'pf-v6-form';

  #formElement;

  async afterTemplateLoaded() {
    this.#formElement = this.shadowRoot.getElementById('form');

    // Forward submit event to host
    this.#formElement.addEventListener('submit', (e) => {
      this.dispatchEvent(new SubmitEvent('submit', {
        bubbles: true,
        cancelable: true,
        composed: true,
        submitter: e.submitter
      }));
    });
  }

  disconnectedCallback() {
    super.disconnectedCallback?.();
  }

  /**
   * Submits the form programmatically
   */
  submit() {
    this.#formElement?.submit();
  }

  /**
   * Requests form submission with validation
   * @returns {boolean} True if validation passed
   */
  requestSubmit(submitter) {
    return this.#formElement?.requestSubmit(submitter);
  }

  /**
   * Resets the form to default values
   */
  reset() {
    this.#formElement?.reset();
  }

  /**
   * Checks form validity
   * @returns {boolean} True if all controls are valid
   */
  checkValidity() {
    return this.#formElement?.checkValidity() ?? true;
  }

  /**
   * Reports form validity (shows validation messages)
   * @returns {boolean} True if all controls are valid
   */
  reportValidity() {
    return this.#formElement?.reportValidity() ?? true;
  }

  static {
    customElements.define(this.is, this);
  }
}
