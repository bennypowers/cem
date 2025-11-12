/**
 * Demo wrapper component for knobs integration
 */
export class CemServeDemo extends HTMLElement {
  static is = 'cem-serve-demo';
  static { customElements.define(this.is, this); }

  /**
   * Set an attribute on an element in the demo
   * @param {string} selector - CSS selector for target element
   * @param {string} attribute - Attribute name
   * @param {string} value - Attribute value
   * @returns {boolean} - Whether the operation succeeded
   */
  setDemoAttribute(selector, attribute, value) {
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
    if (target) {
      target.setAttribute(attribute, value);
      return true;
    }
    return false;
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
