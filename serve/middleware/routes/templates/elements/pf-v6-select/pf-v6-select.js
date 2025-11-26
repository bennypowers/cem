import { CemFormControl } from '/__cem/cem-form-control.js';

/**
 * @customElement pf-v6-select
 */
class PfV6Select extends CemFormControl {
  static is = 'pf-v6-select';

  static observedAttributes = [
    'value',
    'disabled',
    'invalid',
  ];

  #select;
  #internals = this.attachInternals();
  #observer = new MutationObserver(() => this.#cloneOptions());

  #onChange = () => {
    // Sync the value attribute to match internal select
    this.setAttribute('value', this.#select.value);

    // Update form value via ElementInternals
    this.#internals.setFormValue(this.#select.value);

    this.dispatchEvent(new Event('change', { bubbles: true, composed: true }));
  };

  #onInput = () => {
    this.setAttribute('value', this.#select.value);
    this.#internals.setFormValue(this.#select.value);
    this.dispatchEvent(new Event('input', { bubbles: true, composed: true }));
  };

  /**
   * Returns the internal select element for CemFormControl API.
   * @protected
   * @returns {HTMLSelectElement|null}
   */
  get formControlElement() { return this.#select; }

  get disabled() { return this.hasAttribute('disabled'); }
  set disabled(b) { this.toggleAttribute('disabled', !!b); }

  get invalid() { return this.hasAttribute('invalid'); }
  set invalid(b) { this.toggleAttribute('invalid', !!b); }

  get value() { return this.#select?.value ?? ''; }
  set value(v) {
    if (!this.#select) return;
    this.#select.value = v ?? '';
    this.#internals.setFormValue(this.#select.value);
  }

  afterTemplateLoaded() {
    this.#select = this.shadowRoot.getElementById('select');
    if (!this.#select) return;

    // Clone options from light DOM into shadow select
    this.#cloneOptions();

    // Sync initial attributes from host to internal select
    // (attributeChangedCallback may have run before this.#select was set)
    const initialValue = this.getAttribute('value');
    if (initialValue) {
      this.#select.value = initialValue;
    }

    if (this.disabled) {
      this.#select.disabled = true;
    }

    if (this.invalid) {
      this.#select.setAttribute('aria-invalid', 'true');
    }

    // Set initial form value
    this.#internals.setFormValue(this.#select.value || null);

    // Watch for changes to light DOM options
    this.#observer.observe(this, {
      childList: true,
      subtree: true,
      attributes: true,
    });

    // Forward change events from internal select and sync value
    this.#select.addEventListener('change', this.#onChange);
    this.#select.addEventListener('input', this.#onInput);
  }

  disconnectedCallback() {
    super.disconnectedCallback?.();
    this.#observer?.disconnect();

    // Clean up event listeners
    if (this.#select) {
      this.#select.removeEventListener('change', this.#onChange);
      this.#select.removeEventListener('input', this.#onInput);
    }
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

    // Restore previous value (including empty string for placeholder options)
    this.#select.value = currentValue;
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

  // Note: focus() and blur() delegated automatically via CemFormControl

  static {
    customElements.define(this.is, this);
  }
}
