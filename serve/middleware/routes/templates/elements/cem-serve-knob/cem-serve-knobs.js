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

    // Populate the template
    this.#renderContent();
    this.#setupEventListeners();

    // Sync after page load and browser form restoration
    this.#waitForPageLoad();
  }

  #waitForPageLoad() {
    // Sync on next tick to ensure demo element is ready
    setTimeout(() => {
      this.#syncFromDemoElement();
    }, 0);
  }

  #renderContent() {
    // Set name
    const nameEl = this.#$('knob-name');
    if (nameEl) nameEl.textContent = this.name;

    // Set type
    const typeEl = this.#$('knob-type');
    if (typeEl && this.type) {
      typeEl.textContent = this.type;
    }

    // Create and append input element
    const container = this.#$('input-container');
    if (container) {
      const input = this.#createInput();
      container.appendChild(input);
    }
  }

  #createInput() {
    const type = this.type;
    const value = this.value;
    const defaultValue = this.default;

    let input;

    switch (type) {
      case 'boolean':
        input = document.createElement('input');
        input.type = 'checkbox';
        // Don't set initial checked state - let sync handle it
        input.checked = false;
        break;

      case 'enum': {
        input = document.createElement('select');
        // Add empty option
        const emptyOpt = document.createElement('option');
        emptyOpt.value = '';
        emptyOpt.textContent = '-- Select --';
        input.appendChild(emptyOpt);

        // Add enum values
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
        input = document.createElement('input');
        input.type = 'number';
        input.value = value;
        if (defaultValue) input.placeholder = defaultValue;
        break;

      case 'color': {
        // Create a wrapper for color picker + text input
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

        // If value is a valid hex color, set color picker
        if (/^#[0-9A-Fa-f]{6}$/.test(textInput.value)) {
          colorPicker.value = textInput.value;
        }

        // Sync text input when color picker changes
        colorPicker.addEventListener('input', () => {
          textInput.value = colorPicker.value;
          textInput.dispatchEvent(new Event('input', { bubbles: true }));
        });

        // Update color picker when text input has valid hex color
        textInput.addEventListener('input', () => {
          if (/^#[0-9A-Fa-f]{6}$/.test(textInput.value)) {
            colorPicker.value = textInput.value;
          }
        });

        wrapper.appendChild(colorPicker);
        wrapper.appendChild(textInput);

        // Return wrapper instead of single input
        return wrapper;
      }

      default: // string
        input = document.createElement('input');
        input.type = 'text';
        input.value = value;
        if (defaultValue) input.placeholder = defaultValue;
        break;
    }

    input.classList.add('knob-input');
    input.id = 'input';
    // Don't set name attribute to prevent browser form state restoration
    // input.name = this.name;
    input.autocomplete = 'off';

    return input;
  }

  #setupEventListeners() {
    const input = this.#$('input');
    if (!input) return;

    input.addEventListener('input', () => {
      this.#handleChange(input);
    });

    // For select/checkbox, also listen to change event
    if (input.tagName === 'SELECT' || input.type === 'checkbox') {
      input.addEventListener('change', () => {
        this.#handleChange(input);
      });
    }
  }

  #handleChange(input) {
    let value;

    // Get the value based on input type
    if (input.type === 'checkbox') {
      value = input.checked;
    } else if (input.type === 'number') {
      // Handle cleared/invalid number inputs
      if (input.value === '' || Number.isNaN(input.valueAsNumber)) {
        value = null;
      } else {
        value = input.valueAsNumber;
      }
    } else {
      value = input.value;
    }

    // Update attribute (convert null to empty string for attributes)
    this.value = value === null ? '' : String(value);

    // Let subclasses dispatch their specific event
    this.dispatchChangeEvent(this.name, value);
  }

  dispatchChangeEvent(name, value) {
    // Override in subclasses
  }

  #syncFromDemoElement() {
    // Find the chrome element and demo element
    const chrome = this.closest('cem-serve-chrome');
    if (!chrome) return;

    const tagName = chrome.getAttribute('tag-name');
    if (!tagName) return;

    const demo = chrome.querySelector('cem-serve-demo');
    if (!demo) return;

    const demoElement = demo.querySelector(tagName);
    if (!demoElement) return;

    // Read current value from demo element (subclasses override)
    const currentValue = this.readValueFromElement(demoElement);
    if (currentValue === null || currentValue === undefined) return;

    // Update the input to reflect current value
    const input = this.#$('input');
    if (!input) return;

    if (input.type === 'checkbox') {
      input.checked = currentValue === true || currentValue === '';
    } else {
      input.value = currentValue;
    }
  }

  readValueFromElement(element) {
    // Override in subclasses
    return null;
  }
}

// Attribute Knob Element
export class CemServeKnobAttribute extends CemServeKnobBase {
  static is = 'cem-serve-knob-attribute';
  static { customElements.define(this.is, this); }

  dispatchChangeEvent(name, value) {
    this.dispatchEvent(new KnobAttributeChangeEvent(name, value));
  }

  readValueFromElement(element) {
    // Read attribute value
    if (this.type === 'boolean') {
      return element.hasAttribute(this.name);
    }
    return element.getAttribute(this.name) || '';
  }
}

// Property Knob Element
export class CemServeKnobProperty extends CemServeKnobBase {
  static is = 'cem-serve-knob-property';
  static { customElements.define(this.is, this); }

  dispatchChangeEvent(name, value) {
    this.dispatchEvent(new KnobPropertyChangeEvent(name, value));
  }

  readValueFromElement(element) {
    // Read property value
    const value = element[this.name];
    if (value === undefined || value === null) return '';
    return String(value);
  }
}

// CSS Property Knob Element
export class CemServeKnobCSSProperty extends CemServeKnobBase {
  static is = 'cem-serve-knob-css-property';
  static { customElements.define(this.is, this); }

  dispatchChangeEvent(name, value) {
    this.dispatchEvent(new KnobCSSPropertyChangeEvent(name, value));
  }

  readValueFromElement(element) {
    // Read CSS custom property value
    return getComputedStyle(element).getPropertyValue(this.name).trim() || '';
  }
}
