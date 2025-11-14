class Pfv6Select extends HTMLElement {
  #select = null;

  connectedCallback() {
    this.#select = this.shadowRoot?.getElementById('select');
    if (!this.#select) return;

    // Populate options first (in case attributeChangedCallback fired before connectedCallback)
    this.#populateOptions();

    // Then sync other attributes
    this.#syncAttributes();

    // Forward change events from internal select and sync value
    this.#select.addEventListener('change', () => {
      // Sync the value attribute to match internal select
      this.setAttribute('value', this.#select.value);
      this.dispatchEvent(new Event('change', { bubbles: true, composed: true }));
    });

    this.#select.addEventListener('input', () => {
      this.setAttribute('value', this.#select.value);
      this.dispatchEvent(new Event('input', { bubbles: true, composed: true }));
    });
  }

  static get observedAttributes() {
    return ['value', 'options', 'disabled', 'invalid', 'aria-label', 'aria-labelledby'];
  }

  attributeChangedCallback(name) {
    if (name === 'options') {
      this.#populateOptions();
    }
    this.#syncAttributes();
  }

  #populateOptions() {
    if (!this.#select) return;

    const optionsAttr = this.getAttribute('options');
    if (!optionsAttr) return;

    // Clear existing options
    this.#select.innerHTML = '';

    // Add empty option
    const emptyOpt = document.createElement('option');
    emptyOpt.value = '';
    emptyOpt.textContent = '-- Select --';
    this.#select.appendChild(emptyOpt);

    // Parse comma-separated options
    const options = optionsAttr.split(',').map(o => o.trim());
    options.forEach(opt => {
      const option = document.createElement('option');
      option.value = opt;
      option.textContent = opt;
      this.#select.appendChild(option);
    });

    // Set selected value if present
    const value = this.getAttribute('value');
    if (value) {
      this.#select.value = value;
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

    // Sync aria attributes
    const ariaLabel = this.getAttribute('aria-label');
    if (ariaLabel) {
      this.#select.setAttribute('aria-label', ariaLabel);
    } else {
      this.#select.removeAttribute('aria-label');
    }

    const ariaLabelledby = this.getAttribute('aria-labelledby');
    if (ariaLabelledby) {
      this.#select.setAttribute('aria-labelledby', ariaLabelledby);
    } else {
      this.#select.removeAttribute('aria-labelledby');
    }

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

  get options() {
    return this.getAttribute('options') || '';
  }

  set options(value) {
    this.setAttribute('options', value);
  }

  focus() {
    this.#select?.focus();
  }

  blur() {
    this.#select?.blur();
  }
}

customElements.define('pfv6-select', Pfv6Select);
