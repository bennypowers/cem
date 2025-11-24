import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Form Label Component
 *
 * Automatically associates with control in parent pf-v6-form-group by:
 * - Setting aria-label on the control from label text content
 * - Forwarding clicks to focus the control
 *
 * @slot - Label text content
 * @customElement pf-v6-form-label
 */
class PfV6FormLabel extends CemElement {
  static is = 'pf-v6-form-label';

  #labelElement;
  #slot;

  async afterTemplateLoaded() {
    this.#labelElement = this.shadowRoot.querySelector('label');
    this.#slot = this.shadowRoot.querySelector('slot');

    // Set up label click handler
    this.#labelElement.addEventListener('click', this.#handleClick);

    // Set aria-label on control when slot content changes
    this.#slot.addEventListener('slotchange', this.#updateControlLabel);

    // Initial setup
    this.#updateControlLabel();
  }

  disconnectedCallback() {
    this.#labelElement?.removeEventListener('click', this.#handleClick);
    this.#slot?.removeEventListener('slotchange', this.#updateControlLabel);
  }

  #handleClick = (e) => {
    e.preventDefault();
    const control = this.#getControl();
    if (control && typeof control.focus === 'function') {
      control.focus();
    }
  }

  #updateControlLabel = () => {
    const control = this.#getControl();
    const labelText = this.textContent?.trim() || '';

    if (control && labelText) {
      // Access the shadow DOM of the control to set aria-label on internal input
      const shadowRoot = control.shadowRoot;
      if (!shadowRoot) {
        console.warn('Control has no shadow root');
        return;
      }

      // Find the internal input/select element
      // Different controls use different IDs: text-input, select, switch-input
      const internalInput = shadowRoot.querySelector('input, select');
      if (internalInput) {
        internalInput.setAttribute('aria-label', labelText);
      } else {
        console.warn('Could not find internal input in control shadow DOM');
      }
    }
  }

  #getControl() {
    // Find the control element in the parent form group
    const formGroup = this.closest('pf-v6-form-group');
    if (!formGroup) {
      console.warn('pf-v6-form-label must be inside pf-v6-form-group');
      return null;
    }

    // Get the control slot
    const control = formGroup.querySelector('[slot="control"]');
    if (!control) {
      console.warn('No control found in pf-v6-form-group');
      return null;
    }

    return control;
  }

  static {
    customElements.define(this.is, this);
  }
}
