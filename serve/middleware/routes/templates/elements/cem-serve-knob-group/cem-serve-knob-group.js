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
    const knobType = control.dataset.knobType;
    const knobName = control.dataset.knobName;

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
    const knobType = control.dataset.knobType;
    const knobName = control.dataset.knobName;

    if (!knobType || !knobName) return;

    // Immediate update on change (for select, checkbox)
    const key = `${knobType}-${knobName}`;
    clearTimeout(this.#debounceTimers.get(key));
    this.#applyChange(knobType, knobName, control.value, control.checked);
  }

  #applyChange(type, name, value, checked) {
    // Dispatch custom event that KnobsManager will catch
    let eventName;
    let eventValue = value;

    // Handle boolean values from checkboxes
    if (typeof checked === 'boolean') {
      eventValue = checked;
    }

    switch (type) {
      case 'attribute':
        eventName = 'knob:attribute-change';
        break;

      case 'property':
        eventName = 'knob:property-change';
        // Parse value for properties
        eventValue = this.#parseValue(eventValue);
        break;

      case 'css-property':
        eventName = 'knob:css-property-change';
        break;

      default:
        console.warn(`[KnobGroup] Unknown knob type: ${type}`);
        return;
    }

    const event = new CustomEvent(eventName, {
      bubbles: true,
      composed: true,
      detail: { name, value: eventValue }
    });

    // Add name and value as properties for easier access
    event.name = name;
    event.value = eventValue;

    this.dispatchEvent(event);
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
