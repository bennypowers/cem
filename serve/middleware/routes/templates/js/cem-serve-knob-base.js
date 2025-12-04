// CEM Serve Knobs - Base class for interactive controls

// Base class for knob controls
export class CemServeKnobBase extends HTMLElement {
  #$ = id => this.shadowRoot.getElementById(id);

  get name() {
    return this.getAttribute('name') || '';
  }

  get type() {
    return this.getAttribute('type') || 'string';
  }

  get value() {
    return this.getAttribute('value') || '';
  }

  set value(val) {
    this.setAttribute('value', val);
  }

  get default() {
    return this.getAttribute('default') || '';
  }

  get enumValues() {
    const vals = this.getAttribute('enum-values');
    return vals ? vals.split(',') : [];
  }

  connectedCallback() {
    if (!this.shadowRoot) return;

    // Get the input container
    const container = this.#$('input-container');
    if (!container) {
      this.#reportSSRError('Missing input-container element');
      return;
    }

    // Input should always exist from SSR
    const input = container.querySelector('#input');
    if (!input) {
      this.#reportSSRError(
        `Missing SSR input for type="${this.type}"`,
        `Expected #input element to be server-rendered for ${this.constructor.is}`
      );
      return;
    }

    // Set up event listeners
    this.setupInputListeners(input);
  }

  #reportSSRError(title, message) {
    const errorOverlay = document.querySelector('cem-transform-error-overlay');
    if (errorOverlay) {
      errorOverlay.show(title, message || '', '');
    }
    console.error(`[${this.constructor.is}] ${title}`, message);
  }

  // Main setupInputListeners method that calls abstract createChangeEvent
  setupInputListeners(input) {
    const type = this.type;

    if (type === 'boolean') {
      input.addEventListener('change', () => {
        const value = input.checked;
        this.value = String(value);
        this.dispatchEvent(this.createChangeEvent(this.name, value));
      });
    } else if (type === 'number') {
      input.addEventListener('input', () => {
        let value;
        if (input.value === '' || Number.isNaN(input.valueAsNumber)) {
          value = null;
        } else {
          value = input.valueAsNumber;
        }
        this.value = value === null ? '' : String(value);
        this.dispatchEvent(this.createChangeEvent(this.name, value));
      });
    } else if (type === 'enum') {
      input.addEventListener('change', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(this.createChangeEvent(this.name, value));
      });
    } else if (type === 'color') {
      // Color input has both a color picker and text input that need to stay in sync
      const colorPicker = this.shadowRoot.getElementById('color-picker');

      // Sync color picker to text input
      if (colorPicker) {
        colorPicker.addEventListener('input', () => {
          input.value = colorPicker.value;
          input.dispatchEvent(new Event('input', { bubbles: true }));
        });

        // Sync text input to color picker (when valid hex)
        input.addEventListener('input', () => {
          if (/^#[0-9A-Fa-f]{6}$/.test(input.value)) {
            colorPicker.value = input.value;
          }
        });
      }

      // Standard input handler for color text input
      input.addEventListener('input', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(this.createChangeEvent(this.name, value));
      });
    } else {
      // string
      input.addEventListener('input', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(this.createChangeEvent(this.name, value));
      });
    }
  }

  // Abstract method to be implemented by subclasses
  createChangeEvent(name, value) {
    throw new Error('createChangeEvent() must be implemented by subclass');
  }
}
