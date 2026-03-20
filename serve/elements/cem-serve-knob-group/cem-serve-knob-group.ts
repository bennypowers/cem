import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './cem-serve-knob-group.css' with { type: 'css' };

import '../pf-v6-text-input-group/pf-v6-text-input-group.js';

/**
 * Custom event fired when a knob attribute changes
 */
export class KnobAttributeChangeEvent extends Event {
  name: string;
  value: unknown;
  constructor(name: string, value: unknown) {
    super('knob:attribute-change', { bubbles: true });
    this.name = name;
    this.value = value;
  }
}

/**
 * Custom event fired when a knob property changes
 */
export class KnobPropertyChangeEvent extends Event {
  name: string;
  value: unknown;
  constructor(name: string, value: unknown) {
    super('knob:property-change', { bubbles: true });
    this.name = name;
    this.value = value;
  }
}

/**
 * Custom event fired when a knob CSS property changes
 */
export class KnobCssPropertyChangeEvent extends Event {
  name: string;
  value: string;
  constructor(name: string, value: string) {
    super('knob:css-property-change', { bubbles: true });
    this.name = name;
    this.value = value;
  }
}

/**
 * Custom event fired when a knob attribute is cleared
 */
export class KnobAttributeClearEvent extends Event {
  name: string;
  constructor(name: string) {
    super('knob:attribute-clear', { bubbles: true });
    this.name = name;
  }
}

/**
 * Custom event fired when a knob property is cleared
 */
export class KnobPropertyClearEvent extends Event {
  name: string;
  constructor(name: string) {
    super('knob:property-clear', { bubbles: true });
    this.name = name;
  }
}

/**
 * Custom event fired when a knob CSS property is cleared
 */
export class KnobCssPropertyClearEvent extends Event {
  name: string;
  constructor(name: string) {
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
 * @slot - Default slot for knob controls
 */
@customElement('cem-serve-knob-group')
export class CemServeKnobGroup extends LitElement {
  static styles = styles;

  @property({ reflect: true, attribute: 'for' })
  accessor htmlFor: string | undefined;

  #debounceTimers = new Map<string, ReturnType<typeof setTimeout>>();
  #debounceDelay = 250;
  #colorButtonListeners = new WeakMap<Element, EventListener>();
  #clearButtonListeners = new WeakMap<Element, EventListener>();

  render() {
    return html`<slot @slotchange=${this.#onSlotChange}></slot>`;
  }

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('input', this.#handleInput);
    this.addEventListener('change', this.#handleChange);
  }

  firstUpdated() {
    this.#attachColorButtonListeners();
    this.#attachClearButtonListeners();
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    for (const timer of this.#debounceTimers.values()) {
      clearTimeout(timer);
    }
    this.#debounceTimers.clear();
    this.removeEventListener('input', this.#handleInput);
    this.removeEventListener('change', this.#handleChange);
    this.#removeColorButtonListeners();
    this.#removeClearButtonListeners();
  }

  #onSlotChange() {
    this.#attachColorButtonListeners();
    this.#attachClearButtonListeners();
  }

  #attachColorButtonListeners() {
    const buttons = this.querySelectorAll('.color-picker-button');
    for (const button of buttons) {
      if (this.#colorButtonListeners.has(button)) continue;
      const handler = (e: Event) => this.#handleColorButtonClick(e, button);
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
      if (this.#clearButtonListeners.has(button)) continue;
      const handler = (e: Event) => this.#handleClearButtonClick(e, button as HTMLElement);
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

  #handleClearButtonClick = (e: Event, button: HTMLElement) => {
    e.preventDefault();

    const knobType = button.dataset.knobType;
    const knobName = button.dataset.knobName;

    if (!knobType || !knobName) return;

    const formGroup = button.closest('pf-v6-form-group');
    if (!formGroup) return;

    const control = formGroup.querySelector(
      `[data-knob-type="${knobType}"][data-knob-name="${knobName}"]`,
    ) as HTMLInputElement | null;
    if (!control) return;

    if (this.#isBooleanControl(control)) {
      control.checked = false;
    } else {
      control.value = '';
    }

    button.hidden = true;

    control.dispatchEvent(new Event('input', { bubbles: true }));

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
  };

  #updateClearButtonVisibility(control: HTMLElement) {
    const knobType = control.dataset.knobType;
    const knobName = control.dataset.knobName;

    if (!knobType || !knobName) return;

    const formGroup = control.closest('pf-v6-form-group');
    if (!formGroup) return;

    const clearButton = formGroup.querySelector(
      `.knob-clear-button[data-knob-type="${knobType}"][data-knob-name="${knobName}"]`,
    ) as HTMLElement | null;
    if (!clearButton) return;

    const hasValue = this.#isBooleanControl(control)
      ? (control as HTMLInputElement).checked
      : (control as HTMLInputElement).value !== '';

    clearButton.hidden = !hasValue;
  }

  #handleInput = (e: Event) => {
    const control = e.target as HTMLInputElement;
    const knobType = control.dataset?.knobType;
    const knobName = control.dataset?.knobName;

    if (!knobType || !knobName) return;

    this.#updateClearButtonVisibility(control);

    if (this.#isBooleanControl(control)) {
      return this.#applyChange(knobType, knobName, control.checked);
    }

    if (control.tagName === 'SELECT') {
      return;
    }

    const key = `${knobType}-${knobName}`;
    clearTimeout(this.#debounceTimers.get(key));

    this.#debounceTimers.set(key, setTimeout(() => {
      this.#applyChange(knobType, knobName, control.value);
    }, this.#debounceDelay));
  };

  #handleChange = (e: Event) => {
    const control = e.target as HTMLInputElement;
    const knobType = control.dataset?.knobType;
    const knobName = control.dataset?.knobName;
    if (!knobType || !knobName) return;

    this.#updateClearButtonVisibility(control);

    const key = `${knobType}-${knobName}`;
    clearTimeout(this.#debounceTimers.get(key));
    const value = this.#isBooleanControl(control) ? control.checked : control.value;
    this.#applyChange(knobType, knobName, value);
  };

  #isBooleanControl(control: HTMLElement): boolean {
    if (control.tagName === 'PF-V6-SWITCH') {
      return true;
    }
    if (control.tagName === 'INPUT' && (control as HTMLInputElement).type === 'checkbox') {
      return true;
    }
    return false;
  }

  #handleColorButtonClick = async (e: Event, button: Element) => {
    e.preventDefault();

    const textInputGroup = button.closest('pf-v6-text-input-group') as HTMLElement & { value?: string } | null;
    if (!textInputGroup) return;

    const currentValue = textInputGroup.value || '#000000';

    if ('EyeDropper' in window) {
      try {
        const eyeDropper = new (window as any).EyeDropper();
        const result = await eyeDropper.open();

        textInputGroup.value = result.sRGBHex;
        textInputGroup.dispatchEvent(new Event('input', { bubbles: true }));
      } catch (err: any) {
        if (err.name !== 'AbortError') {
          console.warn('[KnobGroup] EyeDropper error:', err);
        }
      }
    } else {
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

      colorInput.addEventListener('blur', () => {
        if (colorInput.parentNode) {
          document.body.removeChild(colorInput);
        }
      });

      colorInput.click();
    }
  };

  #applyChange(type: string, name: string, value: unknown) {
    switch (type) {
      case 'attribute':
        this.dispatchEvent(new KnobAttributeChangeEvent(name, value));
        break;
      case 'property':
        this.dispatchEvent(new KnobPropertyChangeEvent(name, this.#parseValue(value)));
        break;
      case 'css-property':
        this.dispatchEvent(new KnobCssPropertyChangeEvent(name, value as string));
        break;
      default:
        console.warn(`[KnobGroup] Unknown knob type: ${type}`);
        return;
    }
  }

  #parseValue(value: unknown): unknown {
    if (typeof value === 'boolean') {
      return value;
    }

    if (value === 'true') return true;
    if (value === 'false') return false;
    if (value === 'null') return null;
    if (value === '') return '';

    const num = Number(value);
    if (!isNaN(num) && value !== '') {
      return num;
    }

    return value;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-serve-knob-group': CemServeKnobGroup;
  }
}
