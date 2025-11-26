import { CemElement } from '/__cem/cem-element.js';

/**
 * Base class for form control elements.
 *
 * Provides standard interface for:
 * - Accessible label association
 * - Form participation
 * - Automatic focus delegation via delegatesFocus
 *
 * Subclasses should:
 * - Call super.afterTemplateLoaded() if they override it
 * - Implement formControlElement() getter to return the internal input/select element
 *
 * @example
 * class PfV6TextInput extends CemFormControl {
 *   get formControlElement() {
 *     return this.shadowRoot.getElementById('text-input');
 *   }
 * }
 */
export class CemFormControl extends CemElement {
  static shadowRootOptions = { mode: 'open', delegatesFocus: true };
  /**
   * Override in subclasses to return the internal input/select/textarea element.
   * This element will receive aria-label and be the target of focus() calls.
   *
   * @protected
   * @returns {HTMLInputElement|HTMLSelectElement|HTMLTextAreaElement|null}
   */
  get formControlElement() {
    // Default implementation tries to find common form controls
    return this.shadowRoot?.querySelector('input, select, textarea') ?? null;
  }

  /**
   * Sets the accessible label on the internal form control.
   * Called by pf-v6-form-label when label text changes.
   *
   * @param {string} labelText - The label text to set
   */
  setAccessibleLabel(labelText) {
    const control = this.formControlElement;
    if (control && labelText) {
      control.setAttribute('aria-label', labelText);
    }
  }

  // Note: focus() and blur() are handled automatically via delegatesFocus: true

  /**
   * Gets the current value of the form control.
   * @returns {any} The form control value
   */
  get value() {
    return this.formControlElement?.value;
  }

  /**
   * Sets the value of the form control.
   * @param {any} val - The value to set
   */
  set value(val) {
    const control = this.formControlElement;
    if (control) {
      control.value = val;
    }
  }

  /**
   * Checks if the form control is disabled.
   * @returns {boolean}
   */
  get disabled() {
    return this.hasAttribute('disabled');
  }

  /**
   * Sets the disabled state of the form control.
   * @param {boolean} val - True to disable
   */
  set disabled(val) {
    if (val) {
      this.setAttribute('disabled', '');
    } else {
      this.removeAttribute('disabled');
    }
  }

  /**
   * Checks validity of the internal form control.
   * @returns {boolean} True if valid
   */
  checkValidity() {
    return this.formControlElement?.checkValidity() ?? true;
  }

  /**
   * Reports validity of the internal form control (shows validation message).
   * @returns {boolean} True if valid
   */
  reportValidity() {
    return this.formControlElement?.reportValidity() ?? true;
  }

  /**
   * Sets custom validity message on the internal form control.
   * @param {string} message - The validation message
   */
  setCustomValidity(message) {
    this.formControlElement?.setCustomValidity(message);
  }
}
