/**
 * Demo wrapper component for knobs integration
 */
export class CemServeDemo extends HTMLElement {
  static is = 'cem-serve-demo';
  static { customElements.define(this.is, this); }

  connectedCallback() {
    // Listen for knob change events from parent chrome element
    // Events bubble up through shadow DOM (composed: true)
    const chrome = this.closest('cem-serve-chrome');
    if (!chrome) return;

    chrome.addEventListener('knob:attribute-change', this.#handleAttributeChange);
    chrome.addEventListener('knob:property-change', this.#handlePropertyChange);
    chrome.addEventListener('knob:css-property-change', this.#handleCSSPropertyChange);
  }

  /**
   * Extract target element info from knob event by traversing composed path
   * to find the knob group container with data-tag-name and data-instance-index
   */
  #getKnobTarget(event) {
    const chrome = this.closest('cem-serve-chrome');
    const defaultTagName = chrome?.getAttribute('tag-name') || '';

    if (event.composedPath) {
      for (const element of event.composedPath()) {
        if (!(element instanceof Element)) continue;

        // Look for element with data-is-element-knob marker
        if (element.dataset?.isElementKnob === 'true') {
          const tagName = element.dataset.tagName || defaultTagName;
          let instanceIndex = Number.parseInt(element.dataset.instanceIndex ?? '', 10);
          if (Number.isNaN(instanceIndex)) instanceIndex = 0;
          return { tagName, instanceIndex };
        }
      }
    }

    return { tagName: defaultTagName, instanceIndex: 0 };
  }

  /**
   * Find the Nth instance of an element by tag name
   */
  #getElementInstance(tagName, instanceIndex = 0) {
    const root = this.shadowRoot ?? this;
    const elements = root.querySelectorAll(tagName);
    return elements[instanceIndex] || null;
  }

  #handleAttributeChange = (e) => {
    const { tagName, instanceIndex } = this.#getKnobTarget(e);
    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn('[cem-serve-demo] Element not found:', tagName, 'at index', instanceIndex);
      return;
    }

    const { name, value } = e;
    if (typeof value === 'boolean') {
      element.toggleAttribute(name, value);
    } else if (value === '' || value === null || value === undefined) {
      element.removeAttribute(name);
    } else {
      element.setAttribute(name, value);
    }
  };

  #handlePropertyChange = (e) => {
    const { tagName, instanceIndex } = this.#getKnobTarget(e);
    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn('[cem-serve-demo] Element not found:', tagName, 'at index', instanceIndex);
      return;
    }

    const { name, value } = e;
    element[name] = value;
  };

  #handleCSSPropertyChange = (e) => {
    const { tagName, instanceIndex } = this.#getKnobTarget(e);
    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn('[cem-serve-demo] Element not found:', tagName, 'at index', instanceIndex);
      return;
    }

    const { name, value } = e;
    const propertyName = name.startsWith('--') ? name : `--${name}`;
    if (value === '' || value === null || value === undefined) {
      element.style.removeProperty(propertyName);
    } else {
      element.style.setProperty(propertyName, value);
    }
  };

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
