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
    if (!container) return;

    // Check if input already exists (from SSR)
    let input = container.querySelector('#input');

    if (!input) {
      // No SSR input, create dynamically
      input = this.createInput();
      container.appendChild(input);
    }

    // Let subclasses set up their own event listeners
    this.setupInputListeners(input);
  }

  // Factory method for creating pfv6-switch
  createPfv6Switch(name, value) {
    const input = document.createElement('pfv6-switch');
    input.id = 'input';
    input.checked = false;
    input.textContent = name;
    return input;
  }

  // Factory method for creating pfv6-select
  createPfv6Select(name, value, defaultValue, enumValues) {
    const label = document.createElement('label');
    label.textContent = name;
    label.htmlFor = 'input';

    const input = document.createElement('pfv6-select');
    input.id = 'input';
    input.setAttribute('options', enumValues.join(','));
    input.value = value;
    input.autocomplete = 'off';

    label.appendChild(input);
    return label;
  }

  // Factory method for creating pfv6-text-input (string or number)
  createPfv6TextInput(name, value, defaultValue, inputType = 'text') {
    const label = document.createElement('label');
    label.textContent = name;
    label.htmlFor = 'input';

    const input = document.createElement('pfv6-text-input');
    input.type = inputType;
    input.id = 'input';
    input.value = value;
    if (defaultValue) input.placeholder = defaultValue;
    input.autocomplete = 'off';

    label.appendChild(input);
    return label;
  }

  // Factory method for creating color input (with picker + text input)
  createColorInput(name, value, defaultValue) {
    const label = document.createElement('label');
    label.textContent = name;
    label.htmlFor = 'input';

    const wrapper = document.createElement('div');
    wrapper.classList.add('color-input-wrapper');

    const colorPicker = document.createElement('input');
    colorPicker.type = 'color';
    colorPicker.classList.add('color-picker');
    colorPicker.id = 'color-picker';
    colorPicker.autocomplete = 'off';

    const textInput = document.createElement('input');
    textInput.type = 'text';
    textInput.classList.add('color-text');
    textInput.id = 'input';
    textInput.autocomplete = 'off';
    textInput.value = value || defaultValue || '';
    textInput.placeholder = 'Color or CSS variable';

    if (/^#[0-9A-Fa-f]{6}$/.test(textInput.value)) {
      colorPicker.value = textInput.value;
    }

    colorPicker.addEventListener('input', () => {
      textInput.value = colorPicker.value;
      textInput.dispatchEvent(new Event('input', { bubbles: true }));
    });

    textInput.addEventListener('input', () => {
      if (/^#[0-9A-Fa-f]{6}$/.test(textInput.value)) {
        colorPicker.value = textInput.value;
      }
    });

    wrapper.appendChild(colorPicker);
    wrapper.appendChild(textInput);
    label.appendChild(wrapper);
    return label;
  }

  // Main createInput method that uses factory methods
  createInput() {
    const type = this.type;
    const value = this.value;
    const defaultValue = this.default;

    switch (type) {
      case 'boolean':
        return this.createPfv6Switch(this.name, value);
      case 'enum':
        return this.createPfv6Select(this.name, value, defaultValue, this.enumValues);
      case 'number':
        return this.createPfv6TextInput(this.name, value, defaultValue, 'number');
      case 'color':
        return this.createColorInput(this.name, value, defaultValue);
      default: // string
        return this.createPfv6TextInput(this.name, value, defaultValue, 'text');
    }
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
    } else {
      // string, color
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
