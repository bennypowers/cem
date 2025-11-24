import { CemElement } from '/__cem/cem-element.js';

/**
 * CEM Serve Knob Group Component
 *
 * Handles event delegation and debouncing for form controls that modify demo elements.
 * Controls must have data-knob-type and data-knob-name attributes.
 *
 * @attr {string} for - ID of the target element to control
 * @customElement cem-serve-knob-group
 */
class CemServeKnobGroup extends CemElement {
  static is = 'cem-serve-knob-group';
  static observedAttributes = ['for'];

  #debounceTimers = new Map();
  #debounceDelay = 250; // milliseconds

  async afterTemplateLoaded() {
    // Event delegation for input/change events
    this.addEventListener('input', this.#handleInput);
    this.addEventListener('change', this.#handleChange);
  }

  disconnectedCallback() {
    // Clear any pending timers
    for (const timer of this.#debounceTimers.values()) {
      clearTimeout(timer);
    }
    this.#debounceTimers.clear();

    this.removeEventListener('input', this.#handleInput);
    this.removeEventListener('change', this.#handleChange);
  }

  #handleInput = (e) => {
    const control = e.target;
    const knobType = control.getAttribute('data-knob-type');
    const knobName = control.getAttribute('data-knob-name');

    if (!knobType || !knobName) return;

    // Debounce updates (for text inputs)
    const key = `${knobType}-${knobName}`;
    clearTimeout(this.#debounceTimers.get(key));

    this.#debounceTimers.set(key, setTimeout(() => {
      this.#applyChange(knobType, knobName, control.value, control.checked);
    }, this.#debounceDelay));
  }

  #handleChange = (e) => {
    const control = e.target;
    const knobType = control.getAttribute('data-knob-type');
    const knobName = control.getAttribute('data-knob-name');

    if (!knobType || !knobName) return;

    // Immediate update on change (for select, checkbox)
    const key = `${knobType}-${knobName}`;
    clearTimeout(this.#debounceTimers.get(key));
    this.#applyChange(knobType, knobName, control.value, control.checked);
  }

  #applyChange(type, name, value, checked) {
    const targetId = this.getAttribute('for');
    if (!targetId) {
      console.warn('cem-serve-knob-group has no "for" attribute');
      return;
    }

    const element = document.getElementById(targetId);
    if (!element) {
      console.warn(`Target element not found: ${targetId}`);
      return;
    }

    switch (type) {
      case 'attribute':
        this.#applyAttributeChange(element, name, value, checked);
        break;

      case 'property':
        this.#applyPropertyChange(element, name, value, checked);
        break;

      case 'css-property':
        this.#applyCSSPropertyChange(element, name, value);
        break;

      default:
        console.warn(`Unknown knob type: ${type}`);
    }
  }

  #applyAttributeChange(element, name, value, checked) {
    // Handle boolean attributes (from switches)
    if (typeof checked === 'boolean') {
      if (checked) {
        element.setAttribute(name, '');
      } else {
        element.removeAttribute(name);
      }
      return;
    }

    // Handle value-based attributes
    if (value === '' || value === 'false') {
      element.removeAttribute(name);
    } else if (value === 'true') {
      element.setAttribute(name, '');
    } else {
      element.setAttribute(name, value);
    }
  }

  #applyPropertyChange(element, name, value, checked) {
    // Handle boolean properties (from switches)
    if (typeof checked === 'boolean') {
      element[name] = checked;
      return;
    }

    // Parse value and set property
    element[name] = this.#parseValue(value);
  }

  #applyCSSPropertyChange(element, name, value) {
    if (value === '') {
      element.style.removeProperty(name);
    } else {
      element.style.setProperty(name, value);
    }
  }

  #parseValue(value) {
    // Try to parse as JSON for complex types
    if (value === 'true') return true;
    if (value === 'false') return false;
    if (value === 'null') return null;
    if (value === '') return '';

    // Try to parse as number
    const num = Number(value);
    if (!isNaN(num) && value !== '') {
      return num;
    }

    return value;
  }

  static {
    customElements.define(this.is, this);
  }
}
