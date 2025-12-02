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
 * Custom event fired when a knob attribute is cleared
 */
export class KnobAttributeClearEvent extends Event {
  constructor(name) {
    super('knob:attribute-clear', { bubbles: true });
    this.name = name;
  }
}

/**
 * Custom event fired when a knob property is cleared
 */
export class KnobPropertyClearEvent extends Event {
  constructor(name) {
    super('knob:property-clear', { bubbles: true });
    this.name = name;
  }
}

/**
 * Custom event fired when a knob CSS property is cleared
 */
export class KnobCssPropertyClearEvent extends Event {
  constructor(name) {
    super('knob:css-property-clear', { bubbles: true });
    this.name = name;
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
export class CemServeKnobGroup extends CemElement {
  static is = 'cem-serve-knob-group';

  static observedAttributes = ['for'];

  #debounceTimers = new Map();
  #debounceDelay = 250; // milliseconds
  #colorButtonListeners = new WeakMap(); // Track click listeners for cleanup
  #clearButtonListeners = new WeakMap(); // Track clear button click listeners

  async afterTemplateLoaded() {
    // Event delegation for input/change events
    this.addEventListener('input', this.#handleInput);
    this.addEventListener('change', this.#handleChange);

    // Attach click listeners to color picker buttons
    this.#attachColorButtonListeners();

    // Attach click listeners to clear buttons
    this.#attachClearButtonListeners();

    // Re-attach when slot content changes
    const slot = this.shadowRoot?.querySelector('slot');
    if (slot) {
      slot.addEventListener('slotchange', () => {
        this.#attachColorButtonListeners();
        this.#attachClearButtonListeners();
      });
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

    // Remove click listeners from clear buttons
    this.#removeClearButtonListeners();
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

  #attachClearButtonListeners() {
    const buttons = this.querySelectorAll('.knob-clear-button');
    for (const button of buttons) {
      // Skip if already has listener
      if (this.#clearButtonListeners.has(button)) continue;

      // Create and store the bound handler
      const handler = (e) => this.#handleClearButtonClick(e, button);
      this.#clearButtonListeners.set(button, handler);

      button.addEventListener('click', handler);
    }
  }

  #removeClearButtonListeners() {
    const buttons = this.querySelectorAll('.knob-clear-button');
    for (const button of buttons) {
      const handler = this.#clearButtonListeners.get(button);
      if (handler) {
        button.removeEventListener('click', handler);
        this.#clearButtonListeners.delete(button);
      }
    }
  }

  #handleClearButtonClick = (e, button) => {
    e.preventDefault();

    const knobType = button.dataset.knobType;
    const knobName = button.dataset.knobName;

    if (!knobType || !knobName) return;

    // Find the associated control (input/select) within the same form-group
    const formGroup = button.closest('pf-v6-form-group');
    if (!formGroup) return;

    // Find the control with matching knob type and name
    const control = formGroup.querySelector(`[data-knob-type="${knobType}"][data-knob-name="${knobName}"]`);
    if (!control) return;

    // Clear the control value
    if (this.#isBooleanControl(control)) {
      control.checked = false;
    } else {
      control.value = '';
    }

    // Hide the clear button
    button.hidden = true;

    // Dispatch input event so the change is processed
    control.dispatchEvent(new Event('input', { bubbles: true }));

    // Dispatch the appropriate clear event
    switch (knobType) {
      case 'attribute':
        this.dispatchEvent(new KnobAttributeClearEvent(knobName));
        break;
      case 'property':
        this.dispatchEvent(new KnobPropertyClearEvent(knobName));
        break;
      case 'css-property':
        this.dispatchEvent(new KnobCssPropertyClearEvent(knobName));
        break;
    }
  }

  #updateClearButtonVisibility(control) {
    const knobType = control.dataset.knobType;
    const knobName = control.dataset.knobName;

    if (!knobType || !knobName) return;

    // Find the associated clear button
    const formGroup = control.closest('pf-v6-form-group');
    if (!formGroup) return;

    const clearButton = formGroup.querySelector(`.knob-clear-button[data-knob-type="${knobType}"][data-knob-name="${knobName}"]`);
    if (!clearButton) return;

    // Show/hide based on whether control has a value
    const hasValue = this.#isBooleanControl(control)
      ? control.checked
      : control.value !== '';

    clearButton.hidden = !hasValue;
  }

  #handleInput = (e) => {
    const control = e.target;
    const knobType = control.dataset.knobType;
    const knobName = control.dataset.knobName;

    if (!knobType || !knobName) return;

    // Update clear button visibility
    this.#updateClearButtonVisibility(control);

    // Boolean controls (checkboxes, switches) should be handled immediately
    // Selects use change event only, skip them here
    if (this.#isBooleanControl(control)) {
      return this.#applyChange(knobType, knobName, control.checked);
    }

    if (control.tagName === 'SELECT') {
      return; // Let change event handle selects
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

    // Update clear button visibility
    this.#updateClearButtonVisibility(control);

    // Immediate update on change (for select, checkbox)
    const key = `${knobType}-${knobName}`;
    clearTimeout(this.#debounceTimers.get(key));
    const value = this.#isBooleanControl(control) ? control.checked : control.value;
    this.#applyChange(
      knobType,
      knobName,
      value,
    );
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

  #applyChange(type, name, value) {
    switch (type) {
      case 'attribute':
        this.dispatchEvent(new KnobAttributeChangeEvent(name, value));
        break;
      case 'property':
        // Parse value for properties
        this.dispatchEvent(new KnobPropertyChangeEvent(name, this.#parseValue(value)));
        break;
      case 'css-property':
        this.dispatchEvent(new KnobCssPropertyChangeEvent(name, value));
        break;
      default:
        console.warn(`[KnobGroup] Unknown knob type: ${type}`);
        return;
    }
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
