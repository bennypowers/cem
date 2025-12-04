import { CemFormControl } from '/__cem/cem-form-control.js';

/**
 * @customElement pf-v6-text-input
 */
class PfV6TextInput extends CemFormControl {
  static is = 'pf-v6-text-input';

  static observedAttributes = [
    'disabled',
    'invalid',
    'max',
    'min',
    'placeholder',
    'readonly',
    'step',
    'type',
    'value',
  ];

  #input;
  #internals;

  /**
   * Returns the internal input element for CemFormControl API.
   * @protected
   * @returns {HTMLInputElement|null}
   */
  get formControlElement() { return this.#input; }

  // Public API
  get value() { return this.#input.value; }
  set value(val) { this.#input.value = val; }

  get valueAsNumber() { return this.#input?.valueAsNumber ?? NaN; }

  get disabled() { return this.hasAttribute('disabled'); }
  set disabled(value) { this.toggleAttribute('disabled', !!value); }

  get type() { return this.getAttribute('type') || 'text'; }

  set type(value) { this.setAttribute('type', value); }

  constructor() {
    super();
    // Attach ElementInternals for form association
    this.#internals = this.attachInternals();
  }

  afterTemplateLoaded() {
    this.#input = this.shadowRoot.getElementById('text-input');
    if (!this.#input) return;

    // Forward change events from internal input and sync value
    this.#input.addEventListener('input', () => {
      // Sync the value attribute to match internal input
      this.setAttribute('value', this.#input.value);

      // Update form value via ElementInternals
      this.#internals.setFormValue(this.#input.value);

      this.dispatchEvent(new Event('input', { bubbles: true }));
    });

    this.#input.addEventListener('change', () => {
      this.setAttribute('value', this.#input.value);
      this.#internals.setFormValue(this.#input.value);
      this.dispatchEvent(new Event('change', { bubbles: true }));
    });
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (!this.#input) return;
    if (oldValue === newValue) return;

    switch (name) {
      case 'value': {
        const value = newValue || '';
        if (this.#input.value !== value) {
          this.#input.value = value;
        }
        break;
      }

      case 'type': {
        const type = newValue || 'text';
        this.#input.type = type;

        // When changing away from 'number', remove number-specific attributes
        if (oldValue === 'number' && type !== 'number') {
          this.#input.removeAttribute('min');
          this.#input.removeAttribute('max');
          this.#input.removeAttribute('step');
        }
        break;
      }

      case 'placeholder':
        if (newValue) {
          this.#input.placeholder = newValue;
        } else {
          this.#input.removeAttribute('placeholder');
        }
        break;

      case 'disabled':
        this.#input.disabled = this.hasAttribute('disabled');
        break;

      case 'readonly':
        this.#input.readOnly = this.hasAttribute('readonly');
        break;

      case 'min':
      case 'max':
      case 'step': {
        // Only apply these attributes when input type is 'number'
        const type = this.getAttribute('type') || 'text';
        if (type === 'number') {
          if (newValue !== null) {
            this.#input.setAttribute(name, newValue);
          } else {
            this.#input.removeAttribute(name);
          }
        }
        break;
      }

      case 'invalid':
        if (this.hasAttribute('invalid')) {
          this.#input.setAttribute('aria-invalid', 'true');
        } else {
          this.#input.removeAttribute('aria-invalid');
        }
        break;
    }
  }

  select() {
    this.#input?.select();
  }

  static {
    customElements.define(this.is, this);
  }
}
