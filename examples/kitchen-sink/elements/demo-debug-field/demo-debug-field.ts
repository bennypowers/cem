import { html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import { FormAssociatedElement } from '../../base/form-associated-element.js';
import { LoggingMixin } from '../../mixins/logging-mixin.js';

/**
 * A form-associated element with logging capabilities.
 *
 * Demonstrates combining a base class (FormAssociatedElement) with a
 * mixin (LoggingMixin) to create an element with both form integration
 * and debug logging capabilities.
 *
 * @summary Debuggable form field combining FACE and mixin patterns
 * @fires input - Emitted when the value changes
 * @fires change - Emitted when the input loses focus
 * @fires focus - Emitted when the field gains focus
 * @fires blur - Emitted when the field loses focus
 * @slot prefix - Content before the input
 * @slot suffix - Content after the input
 * @csspart container - The outer container
 * @csspart input - The native input element
 * @csspart debug - The debug info panel
 * @cssprop --debug-field-bg - Background color (default: #fff)
 * @cssprop --debug-field-border - Border color (default: #ccc)
 */
@customElement('demo-debug-field')
export class DemoDebugField extends LoggingMixin(FormAssociatedElement) {
  static styles = css`
    :host {
      display: block;
      margin-bottom: 1rem;
    }

    #container {
      padding: 1rem;
      background: var(--debug-field-bg, #fff);
      border: 1px solid var(--debug-field-border, #ccc);
      border-radius: 4px;
    }

    .input-wrapper {
      display: flex;
      align-items: center;
      gap: 0.5rem;
      margin-bottom: 0.5rem;
    }

    input {
      flex: 1;
      padding: 0.5rem;
      border: 1px solid #ddd;
      border-radius: 4px;
    }

    input:focus {
      outline: none;
      border-color: #0066cc;
    }

    #debug {
      margin-top: 0.5rem;
      padding: 0.5rem;
      background: #f5f5f5;
      border-radius: 4px;
      font-family: monospace;
      font-size: 0.75rem;
    }

    .debug-row {
      display: flex;
      gap: 0.5rem;
    }

    .debug-label {
      font-weight: 600;
      min-width: 100px;
    }
  `;

  /**
   * The field label
   */
  @property() accessor label = '';

  /**
   * The field name
   */
  @property() accessor name = '';

  /**
   * The placeholder text
   */
  @property() accessor placeholder = '';

  /**
   * Whether to show debug information
   */
  @property({ type: Boolean }) accessor showDebug = false;

  private handleInput(e: Event) {
    const input = e.target as HTMLInputElement;
    const oldValue = this.value;
    this.value = input.value;

    this.log('Input changed', {
      oldValue,
      newValue: this.value,
      length: this.value.length,
    });

    this.dispatchEvent(new Event('input', { bubbles: true, composed: true }));
    this.requestUpdate();
  }

  private handleChange(e: Event) {
    this.log('Input change event');
    this.dispatchEvent(new Event('change', { bubbles: true, composed: true }));
  }

  private handleFocus() {
    this.log('Input focused');
    this.dispatchEvent(new Event('focus', { bubbles: true, composed: true }));
  }

  private handleBlur() {
    this.log('Input blurred', {
      valid: this.validity.valid,
      value: this.value,
    });
    this.dispatchEvent(new Event('blur', { bubbles: true, composed: true }));
  }

  connectedCallback() {
    super.connectedCallback();
    this.log('Connected with initial value', { value: this.value });
  }

  render() {
    return html`
      <div id="container" part="container">
        ${this.label ? html`<label>${this.label}</label>` : ''}
        <div class="input-wrapper">
          <slot name="prefix"></slot>
          <input
            part="input"
            type="text"
            .value=${this.value}
            .name=${this.name}
            .placeholder=${this.placeholder}
            @input=${this.handleInput}
            @change=${this.handleChange}
            @focus=${this.handleFocus}
            @blur=${this.handleBlur}
          />
          <slot name="suffix"></slot>
        </div>
        <label>
          <input
            type="checkbox"
            .checked=${this.enableLogging}
            @change=${(e: Event) => {
              this.enableLogging = (e.target as HTMLInputElement).checked;
            }}
          />
          Enable logging
        </label>
        ${this.showDebug
          ? html`
              <div id="debug" part="debug">
                <div class="debug-row">
                  <span class="debug-label">Value:</span>
                  <span>${this.value || '(empty)'}</span>
                </div>
                <div class="debug-row">
                  <span class="debug-label">Valid:</span>
                  <span>${this.validity.valid}</span>
                </div>
                <div class="debug-row">
                  <span class="debug-label">Form:</span>
                  <span>${this.form?.id || '(none)'}</span>
                </div>
                <div class="debug-row">
                  <span class="debug-label">Logging:</span>
                  <span>${this.enableLogging ? 'enabled' : 'disabled'}</span>
                </div>
              </div>
            `
          : ''}
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-debug-field': DemoDebugField;
  }
}
