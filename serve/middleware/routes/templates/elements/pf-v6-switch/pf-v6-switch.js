import { CemFormControl } from '/__cem/cem-form-control.js';

/**
 * @customElement pf-v6-switch
 */
export class PfV6Switch extends CemFormControl {
  static is = 'pf-v6-switch';

  static observedAttributes = [
    'checked',
    'disabled',
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

  constructor() {
    super()
    this.addEventListener('input', this.#onChange);
  }

  attributeChangedCallback() {
    if (!this.#input) return;
    this.#input.checked = this.hasAttribute('checked');
    this.#input.disabled = this.hasAttribute('disabled');
  }

  afterTemplateLoaded() {
    this.#input = this.shadowRoot.getElementById('switch-input');
  }

  #onChange() {
    this.checked = this.#input.checked;
    this.#internals.setFormValue(this.#input.checked ? 'on' : null);
  }

  static {
    customElements.define(this.is, this);
  }
}
