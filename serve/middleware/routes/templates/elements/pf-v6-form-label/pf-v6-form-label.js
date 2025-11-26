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
    super.disconnectedCallback?.();
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

  /**
   * Updates the aria-label on the control's internal input element.
   *
   * Supports two approaches:
   * 1. Preferred: Control implements setAccessibleLabel(text) method (CemFormControl API)
   * 2. Fallback: Direct shadowRoot access for legacy controls
   *
   * Compatible controls should extend CemFormControl or implement setAccessibleLabel().
   */
  #updateControlLabel = () => {
    const control = this.#getControl();
    const labelText = this.textContent?.trim() || '';

    if (!control || !labelText) {
      return;
    }

    // Preferred: Use standard API if available
    if (typeof control.setAccessibleLabel === 'function') {
      control.setAccessibleLabel(labelText);
      return;
    }

    // Fallback: Direct shadowRoot access for legacy controls
    const shadowRoot = control.shadowRoot;
    if (!shadowRoot) {
      console.warn('Control has no shadow root and no setAccessibleLabel method');
      return;
    }

    const internalInput = shadowRoot.querySelector('input, select, textarea');
    if (internalInput) {
      internalInput.setAttribute('aria-label', labelText);
    } else {
      console.warn('Could not find internal input in control shadow DOM');
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
