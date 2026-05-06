import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './cem-serve-demo.css' with { type: 'css' };

/**
 * Demo wrapper component for knobs integration.
 *
 * In light/shadow mode, renders demo content via a default slot.
 * In iframe mode, loads the demo in an isolated iframe and bridges
 * knob changes via postMessage.
 *
 * @slot - Default slot for demo content (light/shadow mode only)
 * @customElement cem-serve-demo
 */
@customElement('cem-serve-demo')
export class CemServeDemo extends LitElement {
  static styles = styles;

  @property({ reflect: true }) accessor rendering: string | undefined;

  #iframeReady = false;
  #pendingMessages: Array<Record<string, unknown>> = [];

  get #iframe(): HTMLIFrameElement | null {
    return this.renderRoot.querySelector('iframe');
  }

  render() {
    return (this.rendering === 'iframe') ?  html`
      <iframe part="iframe"
              src="${this.#iframeSrc()}"
              @load="${this.#onIframeLoad}"></iframe>`: html`
      <slot></slot>`;
  }

  /**
   * Apply a knob change to an element in the demo.
   * Called by parent chrome element when knob events occur.
   * In iframe mode, bridges via postMessage instead of direct DOM access.
   */
  applyKnobChange(
    type: string,
    name: string,
    value: unknown,
    tagName: string,
    instanceIndex = 0,
  ): boolean {
    if (this.rendering === 'iframe') {
      return this.#postKnobChange(type, name, value, tagName, instanceIndex);
    }

    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn('[cem-serve-demo] Element not found:', tagName, 'at index', instanceIndex);
      return false;
    }

    switch (type) {
      case 'attribute':
        return this.#applyAttributeChange(element, name, value);
      case 'property':
        return this.#applyPropertyChange(
          element,
          name as keyof Element,
          value as Element[keyof Element],
        );
      case 'css-property':
        return this.#applyCSSPropertyChange(element as HTMLElement, name, value);
      default:
        console.warn('[cem-serve-demo] Unknown knob type:', type);
        return false;
    }
  }

  /**
   * Set an attribute on an element in the demo
   */
  setDemoAttribute(selector: string, attribute: string, value: string | boolean | null): boolean {
    const target = this.querySelector(selector);
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
   */
  setDemoProperty<T extends Element>(
    selector: string,
    property: keyof T,
    value: T[keyof T],
  ): boolean {
    const target = this.querySelector<T>(selector);
    if (target) {
      (target)[property] = value;
      return true;
    }
    return false;
  }

  /**
   * Set a CSS custom property on an element in the demo
   */
  setDemoCssCustomProperty(selector: string, cssProperty: string, value: string): boolean {
    const target = this.querySelector(selector) as HTMLElement | null;
    if (target) {
      const propertyName = cssProperty.startsWith('--') ? cssProperty : `--${cssProperty}`;
      target.style.setProperty(propertyName, value);
      return true;
    }
    return false;
  }

  #iframeSrc(): string {
    const url = new URL(window.location.href);
    url.searchParams.set('rendering', 'chromeless');
    return url.toString();
  }

  #onIframeLoad() {
    this.#iframeReady = true;
    for (const msg of this.#pendingMessages)
      this.#iframe?.contentWindow?.postMessage(msg, window.location.origin);
    this.#pendingMessages = [];
    this.dispatchEvent(new Event('iframe-ready'));
  }

  /**
   * Find the Nth instance of an element by tag name
   */
  #getElementInstance(tagName: string, instanceIndex = 0): Element | null {
    const elements = this.querySelectorAll(tagName);
    return elements[instanceIndex] || null;
  }

  #postKnobChange(
    knobType: string,
    name: string,
    value: unknown,
    tagName: string,
    instanceIndex: number,
  ): boolean {
    const msg = { type: 'cem-knob-change', knobType, name, value, tagName, instanceIndex };
    if (!this.#iframeReady) {
      this.#pendingMessages.push(msg);
      return true;
    }
    const iframe = this.#iframe;
    if (!iframe?.contentWindow) {
      console.warn('[cem-serve-demo] Iframe not ready for postMessage');
      return false;
    }
    iframe.contentWindow.postMessage(msg, window.location.origin);
    return true;
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

  #applyPropertyChange<T extends Element>(
    element: T,
    name: keyof T,
    value: T[keyof T],
  ): boolean {
    if (value === undefined) {
      try {
        delete (element)[name];
      } catch {
        (element)[name] = value;
      }
    } else {
      (element)[name] = value;
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
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-serve-demo': CemServeDemo;
  }
}
