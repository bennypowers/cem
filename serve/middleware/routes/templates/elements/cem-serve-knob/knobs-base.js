// CEM Serve Knobs - Interactive controls for element attributes, properties, and CSS properties

import {
  KnobAttributeChangeEvent,
  KnobPropertyChangeEvent,
  KnobCSSPropertyChangeEvent
} from '/__cem/knob-events.js';

// Base class for knob controls
class CemServeKnobBase extends HTMLElement {
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

  createInput() {
    // Override in subclasses to create appropriate input
    return document.createElement('input');
  }

  setupInputListeners(input) {
    // Override in subclasses to listen to input events
  }
}

// Attribute Knob Element
export class CemServeKnobAttribute extends CemServeKnobBase {
  static is = 'cem-serve-knob-attribute';
  static { customElements.define(this.is, this); }

  createInput() {
    const type = this.type;
    const value = this.value;
    const defaultValue = this.default;

    let input;

    // pfv6-switch handles its own label
    if (type === 'boolean') {
      input = document.createElement('pfv6-switch');
      input.id = 'input';
      input.checked = false;
      input.textContent = this.name;
      return input;
    }

    // All other inputs need a label wrapper
    const label = document.createElement('label');
    label.textContent = this.name;
    label.htmlFor = 'input';

    switch (type) {

      case 'enum': {
        input = document.createElement('select');
        input.id = 'input';
        const emptyOpt = document.createElement('option');
        emptyOpt.value = '';
        emptyOpt.textContent = '-- Select --';
        input.appendChild(emptyOpt);

        this.enumValues.forEach(val => {
          const opt = document.createElement('option');
          opt.value = val;
          opt.textContent = val;
          if (val === value) opt.selected = true;
          input.appendChild(opt);
        });
        break;
      }

      case 'number':
        input = document.createElement('pfv6-text-input');
        input.type = 'number';
        input.id = 'input';
        input.value = value;
        if (defaultValue) input.placeholder = defaultValue;
        break;

      case 'color': {
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

      default: // string
        input = document.createElement('pfv6-text-input');
        input.type = 'text';
        input.id = 'input';
        input.value = value;
        if (defaultValue) input.placeholder = defaultValue;
        break;
    }

    input.autocomplete = 'off';
    label.appendChild(input);
    return label;
  }

  setupInputListeners(input) {
    const type = this.type;

    if (type === 'boolean') {
      input.addEventListener('change', () => {
        const value = input.checked;
        this.value = String(value);
        this.dispatchEvent(new KnobAttributeChangeEvent(this.name, value));
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
        this.dispatchEvent(new KnobAttributeChangeEvent(this.name, value));
      });
    } else if (type === 'enum') {
      input.addEventListener('change', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(new KnobAttributeChangeEvent(this.name, value));
      });
    } else {
      // string, color
      input.addEventListener('input', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(new KnobAttributeChangeEvent(this.name, value));
      });
    }
  }
}

// Property Knob Element
export class CemServeKnobProperty extends CemServeKnobBase {
  static is = 'cem-serve-knob-property';
  static { customElements.define(this.is, this); }

  createInput() {
    const type = this.type;
    const value = this.value;
    const defaultValue = this.default;

    let input;

    // pfv6-switch handles its own label
    if (type === 'boolean') {
      input = document.createElement('pfv6-switch');
      input.id = 'input';
      input.checked = false;
      input.textContent = this.name;
      return input;
    }

    // All other inputs need a label wrapper
    const label = document.createElement('label');
    label.textContent = this.name;
    label.htmlFor = 'input';

    switch (type) {

      case 'enum': {
        input = document.createElement('select');
        input.id = 'input';
        const emptyOpt = document.createElement('option');
        emptyOpt.value = '';
        emptyOpt.textContent = '-- Select --';
        input.appendChild(emptyOpt);

        this.enumValues.forEach(val => {
          const opt = document.createElement('option');
          opt.value = val;
          opt.textContent = val;
          if (val === value) opt.selected = true;
          input.appendChild(opt);
        });
        break;
      }

      case 'number':
        input = document.createElement('pfv6-text-input');
        input.type = 'number';
        input.id = 'input';
        input.value = value;
        if (defaultValue) input.placeholder = defaultValue;
        break;

      case 'color': {
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

      default: // string
        input = document.createElement('pfv6-text-input');
        input.type = 'text';
        input.id = 'input';
        input.value = value;
        if (defaultValue) input.placeholder = defaultValue;
        break;
    }

    input.autocomplete = 'off';
    label.appendChild(input);
    return label;
  }

  setupInputListeners(input) {
    const type = this.type;

    if (type === 'boolean') {
      input.addEventListener('change', () => {
        const value = input.checked;
        this.value = String(value);
        this.dispatchEvent(new KnobPropertyChangeEvent(this.name, value));
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
        this.dispatchEvent(new KnobPropertyChangeEvent(this.name, value));
      });
    } else if (type === 'enum') {
      input.addEventListener('change', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(new KnobPropertyChangeEvent(this.name, value));
      });
    } else {
      // string, color
      input.addEventListener('input', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(new KnobPropertyChangeEvent(this.name, value));
      });
    }
  }
}

// CSS Property Knob Element
export class CemServeKnobCSSProperty extends CemServeKnobBase {
  static is = 'cem-serve-knob-css-property';
  static { customElements.define(this.is, this); }

  createInput() {
    const type = this.type;
    const value = this.value;
    const defaultValue = this.default;

    let input;

    // pfv6-switch handles its own label
    if (type === 'boolean') {
      input = document.createElement('pfv6-switch');
      input.id = 'input';
      input.checked = false;
      input.textContent = this.name;
      return input;
    }

    // All other inputs need a label wrapper
    const label = document.createElement('label');
    label.textContent = this.name;
    label.htmlFor = 'input';

    switch (type) {

      case 'enum': {
        input = document.createElement('select');
        input.id = 'input';
        const emptyOpt = document.createElement('option');
        emptyOpt.value = '';
        emptyOpt.textContent = '-- Select --';
        input.appendChild(emptyOpt);

        this.enumValues.forEach(val => {
          const opt = document.createElement('option');
          opt.value = val;
          opt.textContent = val;
          if (val === value) opt.selected = true;
          input.appendChild(opt);
        });
        break;
      }

      case 'number':
        input = document.createElement('pfv6-text-input');
        input.type = 'number';
        input.id = 'input';
        input.value = value;
        if (defaultValue) input.placeholder = defaultValue;
        break;

      case 'color': {
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

      default: // string
        input = document.createElement('pfv6-text-input');
        input.type = 'text';
        input.id = 'input';
        input.value = value;
        if (defaultValue) input.placeholder = defaultValue;
        break;
    }

    input.autocomplete = 'off';
    label.appendChild(input);
    return label;
  }

  setupInputListeners(input) {
    const type = this.type;

    if (type === 'boolean') {
      input.addEventListener('change', () => {
        const value = input.checked;
        this.value = String(value);
        this.dispatchEvent(new KnobCSSPropertyChangeEvent(this.name, value));
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
        this.dispatchEvent(new KnobCSSPropertyChangeEvent(this.name, value));
      });
    } else if (type === 'enum') {
      input.addEventListener('change', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(new KnobCSSPropertyChangeEvent(this.name, value));
      });
    } else {
      // string, color
      input.addEventListener('input', () => {
        const value = input.value;
        this.value = value;
        this.dispatchEvent(new KnobCSSPropertyChangeEvent(this.name, value));
      });
    }
  }
}
