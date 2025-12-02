/**
 * Demo wrapper component for knobs integration
 */
export class CemServeDemo extends HTMLElement {
  static is = 'cem-serve-demo';
  static { customElements.define(this.is, this); }

  /**
   * Find the Nth instance of an element by tag name
   */
  #getElementInstance(tagName, instanceIndex = 0) {
    const root = this.shadowRoot ?? this;
    // Search entire subtree for elements (not just direct children)
    const elements = root.querySelectorAll(tagName);
    return elements[instanceIndex] || null;
  }

  /**
   * Apply a knob change to an element in the demo
   * Called by parent chrome element when knob events occur
   * @param {string} type - 'attribute', 'property', or 'css-property'
   * @param {string} name - Attribute/property/CSS name
   * @param {*} value - Value to apply
   * @param {string} tagName - Target element tag name
   * @param {number} instanceIndex - Which instance of the element (0-based)
   * @returns {boolean} - Whether the operation succeeded
   */
  applyKnobChange(type, name, value, tagName, instanceIndex = 0) {
    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn('[cem-serve-demo] Element not found:', tagName, 'at index', instanceIndex);
      return false;
    }

    switch (type) {
      case 'attribute':
        return this.#applyAttributeChange(element, name, value);
      case 'property':
        return this.#applyPropertyChange(element, name, value);
      case 'css-property':
        return this.#applyCSSPropertyChange(element, name, value);
      default:
        console.warn('[cem-serve-demo] Unknown knob type:', type);
        return false;
    }
  }

  #applyAttributeChange(element, name, value) {
    if (typeof value === 'boolean') {
      element.toggleAttribute(name, value);
    } else if (value === '' || value === null || value === undefined) {
      element.removeAttribute(name);
    } else {
      element.setAttribute(name, value);
    }
    return true;
  }

  #applyPropertyChange(element, name, value) {
    if (value === undefined) {
      // Try to delete the property, fall back to setting undefined if that fails
      try {
        delete element[name];
      } catch (e) {
        // Some properties can't be deleted, just set to undefined
        element[name] = undefined;
      }
    } else {
      element[name] = value;
    }
    return true;
  }

  #applyCSSPropertyChange(element, name, value) {
    const propertyName = name.startsWith('--') ? name : `--${name}`;
    if (value === '' || value === null || value === undefined) {
      element.style.removeProperty(propertyName);
    } else {
      element.style.setProperty(propertyName, value);
    }
    return true;
  }

  /**
   * Set an attribute on an element in the demo
   * @param {string} selector - CSS selector for target element
   * @param {string} attribute - Attribute name
   * @param {string|boolean} value - Attribute value (boolean for presence/absence)
   * @returns {boolean} - Whether the operation succeeded
   */
  setDemoAttribute(selector, attribute, value) {
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
    if (!target) return false;

    // Handle boolean attributes (presence/absence)
    if (typeof value === 'boolean') {
      if (value) {
        target.setAttribute(attribute, '');
      } else {
        target.removeAttribute(attribute);
      }
    } else if (value === '' || value === null || value === undefined) {
      target.removeAttribute(attribute);
    } else {
      target.setAttribute(attribute, value);
    }

    return true;
  }

  /**
   * Set a property on an element in the demo
   * @param {string} selector - CSS selector for target element
   * @param {string} property - Property name
   * @param {*} value - Property value
   * @returns {boolean} - Whether the operation succeeded
   */
  setDemoProperty(selector, property, value) {
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
    if (target) {
      target[property] = value;
      return true;
    }
    return false;
  }

  /**
   * Set a CSS custom property on an element in the demo
   * @param {string} selector - CSS selector for target element
   * @param {string} cssProperty - CSS custom property name (with or without --)
   * @param {string} value - CSS property value
   * @returns {boolean} - Whether the operation succeeded
   */
  setDemoCssCustomProperty(selector, cssProperty, value) {
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
    if (target) {
      const propertyName = cssProperty.startsWith('--') ? cssProperty : `--${cssProperty}`;
      target.style.setProperty(propertyName, value);
      return true;
    }
    return false;
  }
}
