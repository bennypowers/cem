class Pfv6TextInput extends HTMLElement {
  #input = null;

  connectedCallback() {
    this.#input = this.shadowRoot?.getElementById('text-input');
    if (!this.#input) return;

    // Sync initial state
    this.#syncAttributes();

    // Forward change events from internal input and sync value
    this.#input.addEventListener('input', () => {
      // Sync the value attribute to match internal input
      this.setAttribute('value', this.#input.value);
      this.dispatchEvent(new Event('input', { bubbles: true, composed: true }));
    });

    this.#input.addEventListener('change', () => {
      this.setAttribute('value', this.#input.value);
      this.dispatchEvent(new Event('change', { bubbles: true, composed: true }));
    });
  }

  static get observedAttributes() {
    return ['value', 'type', 'placeholder', 'disabled', 'readonly', 'invalid', 'min', 'max', 'step', 'aria-label', 'aria-labelledby'];
  }

  attributeChangedCallback() {
    this.#syncAttributes();
  }

  #syncAttributes() {
    if (!this.#input) return;

    // Sync value
    const value = this.getAttribute('value') || '';
    if (this.#input.value !== value) {
      this.#input.value = value;
    }

    // Sync type
    const type = this.getAttribute('type') || 'text';
    this.#input.type = type;

    // Sync placeholder
    const placeholder = this.getAttribute('placeholder');
    if (placeholder) {
      this.#input.placeholder = placeholder;
    } else {
      this.#input.removeAttribute('placeholder');
    }

    // Sync disabled state
    this.#input.disabled = this.hasAttribute('disabled');

    // Sync readonly state
    this.#input.readOnly = this.hasAttribute('readonly');

    // Sync number input attributes
    if (type === 'number') {
      const min = this.getAttribute('min');
      const max = this.getAttribute('max');
      const step = this.getAttribute('step');

      if (min !== null) {
        this.#input.setAttribute('min', min);
      } else {
        this.#input.removeAttribute('min');
      }

      if (max !== null) {
        this.#input.setAttribute('max', max);
      } else {
        this.#input.removeAttribute('max');
      }

      if (step !== null) {
        this.#input.setAttribute('step', step);
      } else {
        this.#input.removeAttribute('step');
      }
    }

    // Sync aria attributes
    const ariaLabel = this.getAttribute('aria-label');
    if (ariaLabel) {
      this.#input.setAttribute('aria-label', ariaLabel);
    } else {
      this.#input.removeAttribute('aria-label');
    }

    const ariaLabelledby = this.getAttribute('aria-labelledby');
    if (ariaLabelledby) {
      this.#input.setAttribute('aria-labelledby', ariaLabelledby);
    } else {
      this.#input.removeAttribute('aria-labelledby');
    }

    // Sync invalid state
    if (this.hasAttribute('invalid')) {
      this.#input.setAttribute('aria-invalid', 'true');
    } else {
      this.#input.removeAttribute('aria-invalid');
    }
  }

  // Public API
  get value() {
    return this.getAttribute('value') || '';
  }

  set value(val) {
    this.setAttribute('value', val);
  }

  get valueAsNumber() {
    return this.#input?.valueAsNumber ?? NaN;
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

  get type() {
    return this.getAttribute('type') || 'text';
  }

  set type(value) {
    this.setAttribute('type', value);
  }

  focus() {
    this.#input?.focus();
  }

  blur() {
    this.#input?.blur();
  }

  select() {
    this.#input?.select();
  }
}

customElements.define('pfv6-text-input', Pfv6TextInput);
