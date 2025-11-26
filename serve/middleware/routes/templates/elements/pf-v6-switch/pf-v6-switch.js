import { CemFormControl } from '/__cem/cem-form-control.js';

/**
 * @customElement pf-v6-switch
 */
class PfV6Switch extends CemFormControl {
  static is = 'pf-v6-switch';

  static observedAttributes = [
    'checked',
    'disabled',
    'aria-label',
    'aria-labelledby',
  ];

  #input;
  #internals = this.attachInternals();

  /**
   * Returns the internal checkbox input element for CemFormControl API.
   * @protected
   * @returns {HTMLInputElement|null}
   */
  get formControlElement() { return this.#input; }

  get checked() { return this.hasAttribute('checked'); }
  set checked(value) { this.toggleAttribute('checked', !!value); }

  get value() { return this.checked; }
  set value(val) { this.checked = val; }

  afterTemplateLoaded() {
    this.#input = this.shadowRoot.getElementById('switch-input');
    if (!this.#input) return;
    this.#input.addEventListener('change', () => {
      this.toggleAttribute('checked', this.#input.checked);
      this.#internals.setFormValue(this.#input.checked ? 'on' : null);
    });
  }

  attributeChangedCallback() {
    if (!this.#input) return;
    this.#input.checked = this.hasAttribute('checked');
    this.#input.disabled = this.hasAttribute('disabled');
  }

  static {
    customElements.define(this.is, this);
  }
}
