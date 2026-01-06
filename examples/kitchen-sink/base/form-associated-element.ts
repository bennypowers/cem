import { LitElement } from 'lit';

/**
 * Base class for form-associated custom elements (FACE).
 *
 * Provides form integration capabilities for custom elements,
 * allowing them to participate in HTML form submission and validation.
 *
 * @summary Form-associated base class for custom elements
 */
export class FormAssociatedElement extends LitElement {
  /**
   * Indicates this element is form-associated
   */
  static formAssociated = true;

  private internals: ElementInternals;
  private _value: string = '';

  constructor() {
    super();
    this.internals = this.attachInternals();
  }

  /**
   * Get the form this element is associated with
   */
  get form(): HTMLFormElement | null {
    return this.internals.form;
  }

  /**
   * Get the current form value
   */
  get value(): string {
    return this._value;
  }

  /**
   * Set the form value
   *
   * @param value - The new form value
   */
  set value(value: string) {
    this._value = value;
    this.internals.setFormValue(value);
  }

  /**
   * Get the current validation message
   */
  get validationMessage(): string {
    return this.internals.validationMessage;
  }

  /**
   * Get the validity state
   */
  get validity(): ValidityState {
    return this.internals.validity;
  }

  /**
   * Check if the element will be validated
   */
  get willValidate(): boolean {
    return this.internals.willValidate;
  }

  /**
   * Check if the form is valid
   *
   * @returns {boolean} True if valid
   */
  checkValidity(): boolean {
    return this.internals.checkValidity();
  }

  /**
   * Report form validity to the user
   *
   * @returns {boolean} True if valid
   */
  reportValidity(): boolean {
    return this.internals.reportValidity();
  }

  /**
   * Set a custom validation message
   *
   * @param message - The validation message
   */
  setCustomValidity(message: string): void {
    if (message) {
      this.internals.setValidity({ customError: true }, message);
    } else {
      this.internals.setValidity({});
    }
  }

  /**
   * Called when the form is reset
   */
  formResetCallback(): void {
    this._value = '';
    this.internals.setFormValue(this._value);
  }

  /**
   * Called when the form is disabled
   *
   * @param disabled - Whether the form is disabled
   */
  formDisabledCallback(disabled: boolean): void {
    this.toggleAttribute('disabled', disabled);
  }
}
