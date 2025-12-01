import '/__cem/elements/pf-v6-text-input-group/pf-v6-text-input-group.js';

import { CemElement } from '/__cem/cem-element.js';

/**
 * Custom event fired when a knob attribute changes
 */
export class KnobAttributeChangeEvent extends Event {
  constructor(name, value) {
    super('knob:attribute-change', { bubbles: true });
    this.name = name;
    this.value = value;
  }
}

/**
 * Custom event fired when a knob property changes
 */
export class KnobPropertyChangeEvent extends Event {
  constructor(name, value) {
    super('knob:property-change', { bubbles: true });
    this.name = name;
    this.value = value;
  }
}

/**
 * Custom event fired when a knob CSS property changes
 */
export class KnobCssPropertyChangeEvent extends Event {
  constructor(name, value) {
    super('knob:css-property-change', { bubbles: true });
    this.name = name;
    this.value = value;
  }
}

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
  #colorButtonListeners = new WeakMap(); // Track click listeners for cleanup

  async afterTemplateLoaded() {
    // Event delegation for input/change events
    this.addEventListener('input', this.#handleInput);
    this.addEventListener('change', this.#handleChange);

    // Attach click listeners to color picker buttons
    this.#attachColorButtonListeners();

    // Re-attach when slot content changes
    const slot = this.shadowRoot?.querySelector('slot');
    if (slot) {
      slot.addEventListener('slotchange', () => this.#attachColorButtonListeners());
    }
  }

  disconnectedCallback() {
    // Clear any pending timers
    for (const timer of this.#debounceTimers.values()) {
      clearTimeout(timer);
    }
    this.#debounceTimers.clear();

    this.removeEventListener('input', this.#handleInput);
    this.removeEventListener('change', this.#handleChange);

    // Remove click listeners from color buttons
    this.#removeColorButtonListeners();
  }

  #attachColorButtonListeners() {
    const buttons = this.querySelectorAll('.color-picker-button');
    for (const button of buttons) {
      // Skip if already has listener
      if (this.#colorButtonListeners.has(button)) continue;

      // Create and store the bound handler
      const handler = (e) => this.#handleColorButtonClick(e, button);
      this.#colorButtonListeners.set(button, handler);

      button.addEventListener('click', handler);
    }
  }

  #removeColorButtonListeners() {
    const buttons = this.querySelectorAll('.color-picker-button');
    for (const button of buttons) {
      const handler = this.#colorButtonListeners.get(button);
      if (handler) {
        button.removeEventListener('click', handler);
        this.#colorButtonListeners.delete(button);
      }
    }
  }

  #handleInput = (e) => {
    const control = e.target;
    const knobType = control.dataset.knobType;
    const knobName = control.dataset.knobName;

    if (!knobType || !knobName) return;

    // Only debounce text-like inputs (not checkboxes, selects, etc)
    // Checkboxes and selects should use the change event for immediate updates
    if (this.#isBooleanControl(control) || control.tagName === 'SELECT') {
      return; // Let change event handle these
    }

    // Debounce updates for text inputs only
    const key = `${knobType}-${knobName}`;
    clearTimeout(this.#debounceTimers.get(key));

    this.#debounceTimers.set(key, setTimeout(() => {
      this.#applyChange(knobType, knobName, control.value);
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
    const checked = this.#isBooleanControl(control) ? control.checked : undefined;
    this.#applyChange(knobType, knobName, control.value, checked);
  }

  /**
   * Check if a control is a boolean-type control (checkbox, switch, etc.)
   * Only these controls should pass their checked state to #applyChange
   * @param {HTMLElement} control
   * @returns {boolean}
   */
  #isBooleanControl(control) {
    // Check for pf-v6-switch custom element
    if (control.tagName === 'PF-V6-SWITCH') {
      return true;
    }

    // Check for native checkbox input
    if (control.tagName === 'INPUT' && control.type === 'checkbox') {
      return true;
    }

    // Add other boolean control types here as needed
    return false;
  }

  #handleColorButtonClick = async (e, button) => {
    e.preventDefault();

    // Find the associated text input group
    const textInputGroup = button.closest('pf-v6-text-input-group');
    if (!textInputGroup) return;

    const currentValue = textInputGroup.value || '#000000';

    // Try to use EyeDropper API if available
    if (window.EyeDropper) {
      try {
        const eyeDropper = new EyeDropper();
        const result = await eyeDropper.open();

        // Update the text input group value
        textInputGroup.value = result.sRGBHex;

        // Trigger input event so the knob system picks it up
        textInputGroup.dispatchEvent(new Event('input', { bubbles: true }));
      } catch (err) {
        // User cancelled or error occurred
        if (err.name !== 'AbortError') {
          console.warn('[KnobGroup] EyeDropper error:', err);
        }
      }
    } else {
      // Fallback: create a temporary native color input
      const colorInput = document.createElement('input');
      colorInput.type = 'color';
      colorInput.value = currentValue;
      colorInput.style.position = 'absolute';
      colorInput.style.opacity = '0';
      colorInput.style.pointerEvents = 'none';

      document.body.appendChild(colorInput);

      colorInput.addEventListener('input', () => {
        textInputGroup.value = colorInput.value;
        textInputGroup.dispatchEvent(new Event('input', { bubbles: true }));
      });

      colorInput.addEventListener('change', () => {
        textInputGroup.value = colorInput.value;
        textInputGroup.dispatchEvent(new Event('input', { bubbles: true }));
        if (colorInput.parentNode) {
          document.body.removeChild(colorInput);
        }
      });

      // Clean up if focus is lost without selection
      colorInput.addEventListener('blur', () => {
        if (colorInput.parentNode) {
          document.body.removeChild(colorInput);
        }
      });

      colorInput.click();
    }
  }

  #applyChange(type, name, value, checked) {
    // Dispatch custom event that KnobsManager will catch
    let eventValue = value;

    // Handle boolean values from checkboxes/switches
    // Only use checked if explicitly provided (not undefined)
    if (checked !== undefined) {
      eventValue = checked;
    }

    let event;
    switch (type) {
      case 'attribute':
        event = new KnobAttributeChangeEvent(name, eventValue);
        break;

      case 'property':
        // Parse value for properties
        eventValue = this.#parseValue(eventValue);
        event = new KnobPropertyChangeEvent(name, eventValue);
        break;

      case 'css-property':
        event = new KnobCssPropertyChangeEvent(name, eventValue);
        break;

      default:
        console.warn(`[KnobGroup] Unknown knob type: ${type}`);
        return;
    }

    this.dispatchEvent(event);
  }

  #parseValue(value) {
    // If already a boolean, return as-is (from checkbox/switch checked state)
    if (typeof value === 'boolean') {
      return value;
    }

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
