import { CemElement } from '/__cem/cem-element.js';

/**
 * @customElement pf-v6-select
 */
class PfV6Select extends CemElement {
  static formAssociated = true;
  static observedAttributes = ['value', 'disabled', 'invalid'];
  static is = 'pf-v6-select';

  #select;
  #internals;
  #observer = new MutationObserver(() => this.#cloneOptions());

  constructor() {
    super();
    // Attach ElementInternals for form association
    this.#internals = this.attachInternals();
  }

  async afterTemplateLoaded() {
    this.#select = this.shadowRoot.getElementById('select');
    if (!this.#select) return;

    // Note: Initial state (disabled, invalid) is rendered server-side via template
    // Clone options from light DOM into shadow select
    this.#cloneOptions();

    // Watch for changes to light DOM options
    this.#observer.observe(this, {
      childList: true,
      subtree: true,
      attributes: true,
    });

    // Forward change events from internal select and sync value
    this.#select.addEventListener('change', () => {
      // Sync the value attribute to match internal select
      this.setAttribute('value', this.#select.value);

      // Update form value via ElementInternals
      this.#internals.setFormValue(this.#select.value);

      this.dispatchEvent(new Event('change', { bubbles: true, composed: true }));
    });

    this.#select.addEventListener('input', () => {
      this.setAttribute('value', this.#select.value);
      this.#internals.setFormValue(this.#select.value);
      this.dispatchEvent(new Event('input', { bubbles: true, composed: true }));
    });
  }

  disconnectedCallback() {
    this.#observer?.disconnect();
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  #cloneOptions() {
    if (!this.#select) return;

    // Get current value before clearing
    const currentValue = this.#select.value;

    // Clear existing options
    this.#select.innerHTML = '';

    // Clone all option elements from light DOM
    const options = this.querySelectorAll('option');
    options.forEach(option => {
      this.#select.appendChild(option.cloneNode(true));
    });

    // Restore value if it exists in new options
    if (currentValue) {
      this.#select.value = currentValue;
    }
  }

  #syncAttributes() {
    if (!this.#select) return;

    // Sync value
    const value = this.getAttribute('value') || '';
    if (this.#select.value !== value) {
      this.#select.value = value;
    }

    // Sync disabled state
    this.#select.disabled = this.hasAttribute('disabled');

    // Sync invalid state
    if (this.hasAttribute('invalid')) {
      this.#select.setAttribute('aria-invalid', 'true');
    } else {
      this.#select.removeAttribute('aria-invalid');
    }
  }

  // Public API
  get value() {
    return this.getAttribute('value') || '';
  }

  set value(val) {
    this.setAttribute('value', val);
  }

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(value) {
    if (value) {
      this.setAttribute('disabled', '');
    } else {
      this.removeAttribute('disabled');
    }
  }


  focus() {
    this.#select?.focus();
  }

  blur() {
    this.#select?.blur();
  }

  static {
    customElements.define(this.is, this);
  }
}
