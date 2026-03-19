import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';

import styles from './cem-serve-demo.css' with { type: 'css' };

/**
 * Demo wrapper component for knobs integration.
 *
 * @slot - Default slot for demo content
 * @customElement cem-serve-demo
 */
@customElement('cem-serve-demo')
export class CemServeDemo extends LitElement {
  static styles = styles;

  render() {
    return html`<slot></slot>`;
  }

  /**
   * Find the Nth instance of an element by tag name
   */
  #getElementInstance(tagName: string, instanceIndex = 0): Element | null {
    const root = this.shadowRoot ?? this;
    const elements = root.querySelectorAll(tagName);
    return elements[instanceIndex] || null;
  }

  /**
   * Apply a knob change to an element in the demo.
   * Called by parent chrome element when knob events occur.
   * @param type - 'attribute', 'property', or 'css-property'
   * @param name - Attribute/property/CSS name
   * @param value - Value to apply
   * @param tagName - Target element tag name
   * @param instanceIndex - Which instance of the element (0-based)
   * @returns Whether the operation succeeded
   */
  applyKnobChange(
    type: string,
    name: string,
    value: unknown,
    tagName: string,
    instanceIndex = 0,
  ): boolean {
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
        return this.#applyCSSPropertyChange(element as HTMLElement, name, value);
      default:
        console.warn('[cem-serve-demo] Unknown knob type:', type);
        return false;
    }
  }

  #applyAttributeChange(element: Element, name: string, value: unknown): boolean {
    if (typeof value === 'boolean') {
      element.toggleAttribute(name, value);
    } else if (value === '' || value === null || value === undefined) {
      element.removeAttribute(name);
    } else {
      element.setAttribute(name, String(value));
    }
    return true;
  }

  #applyPropertyChange(element: Element, name: string, value: unknown): boolean {
    if (value === undefined) {
      try {
        delete (element as Record<string, unknown>)[name];
      } catch {
        (element as Record<string, unknown>)[name] = undefined;
      }
    } else {
      (element as Record<string, unknown>)[name] = value;
    }
    return true;
  }

  #applyCSSPropertyChange(element: HTMLElement, name: string, value: unknown): boolean {
    const propertyName = name.startsWith('--') ? name : `--${name}`;
    if (value === '' || value === null || value === undefined) {
      element.style.removeProperty(propertyName);
    } else {
      element.style.setProperty(propertyName, String(value));
    }
    return true;
  }

  /**
   * Set an attribute on an element in the demo
   * @param selector - CSS selector for target element
   * @param attribute - Attribute name
   * @param value - Attribute value (boolean for presence/absence)
   * @returns Whether the operation succeeded
   */
  setDemoAttribute(selector: string, attribute: string, value: string | boolean): boolean {
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
    if (!target) return false;

    if (typeof value === 'boolean') {
      target.toggleAttribute(attribute, value);
    } else if (value === '' || value === null || value === undefined) {
      target.removeAttribute(attribute);
    } else {
      target.setAttribute(attribute, value);
    }

    return true;
  }

  /**
   * Set a property on an element in the demo
   * @param selector - CSS selector for target element
   * @param property - Property name
   * @param value - Property value
   * @returns Whether the operation succeeded
   */
  setDemoProperty(selector: string, property: string, value: unknown): boolean {
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector);
    if (target) {
      (target as Record<string, unknown>)[property] = value;
      return true;
    }
    return false;
  }

  /**
   * Set a CSS custom property on an element in the demo
   * @param selector - CSS selector for target element
   * @param cssProperty - CSS custom property name (with or without --)
   * @param value - CSS property value
   * @returns Whether the operation succeeded
   */
  setDemoCssCustomProperty(selector: string, cssProperty: string, value: string): boolean {
    const root = this.shadowRoot ?? this;
    const target = root.querySelector(selector) as HTMLElement | null;
    if (target) {
      const propertyName = cssProperty.startsWith('--') ? cssProperty : `--${cssProperty}`;
      target.style.setProperty(propertyName, value);
      return true;
    }
    return false;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-serve-demo': CemServeDemo;
  }
}
