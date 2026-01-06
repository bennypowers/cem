import { html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import { FormAssociatedElement } from '../../base/form-associated-element.js';

/**
 * A form-associated custom element demonstrating FACE pattern.
 *
 * This element extends FormAssociatedElement to participate in HTML forms,
 * providing form validation, submission, and reset capabilities.
 *
 * @summary Form-associated text field
 * @fires input - Emitted when the value changes
 * @fires change - Emitted when the input loses focus
 * @csspart input - The native input element
 * @csspart label - The label element
 * @cssprop --field-border-color - Border color (default: #ddd)
 * @cssprop --field-focus-color - Focus border color (default: #0066cc)
 */
@customElement('demo-field')
export class DemoField extends FormAssociatedElement {
  static styles = css`
    :host {
      display: block;
      margin-bottom: 1rem;
    }

    .field {
      display: flex;
      flex-direction: column;
      gap: 0.5rem;
    }

    label {
      font-weight: 600;
      font-size: 0.875rem;
    }

    input {
      padding: 0.5rem;
      border: 1px solid var(--field-border-color, #ddd);
      border-radius: 4px;
      font-size: 1rem;
    }

    input:focus {
      outline: none;
      border-color: var(--field-focus-color, #0066cc);
    }

    input:disabled {
      background: #f5f5f5;
      cursor: not-allowed;
    }

    .error {
      color: #d32f2f;
      font-size: 0.75rem;
    }
  `;

  /**
   * The field label
   */
  @property() accessor label = '';

  /**
   * The field name for form submission
   */
  @property() accessor name = '';

  /**
   * The placeholder text
   */
  @property() accessor placeholder = '';

  /**
   * Whether the field is required
   */
  @property({ type: Boolean }) accessor required = false;

  /**
   * Whether the field is disabled
   */
  @property({ type: Boolean }) accessor disabled = false;

  /**
   * The minimum length
   */
  @property({ type: Number }) accessor minlength?: number;

  /**
   * The maximum length
   */
  @property({ type: Number }) accessor maxlength?: number;

  /**
   * The input pattern
   */
  @property() accessor pattern?: string;

  private handleInput(e: Event) {
    const input = e.target as HTMLInputElement;
    this.value = input.value;
    this.dispatchEvent(new Event('input', { bubbles: true, composed: true }));
  }

  private handleChange(e: Event) {
    this.dispatchEvent(new Event('change', { bubbles: true, composed: true }));
  }

  render() {
    return html`
      <div class="field">
        ${this.label ? html`<label part="label">${this.label}</label>` : ''}
        <input
          part="input"
          type="text"
          .value=${this.value}
          .name=${this.name}
          .placeholder=${this.placeholder}
          ?required=${this.required}
          ?disabled=${this.disabled}
          .minLength=${this.minlength ?? 0}
          .maxLength=${this.maxlength ?? 524288}
          .pattern=${this.pattern ?? ''}
          @input=${this.handleInput}
          @change=${this.handleChange}
        />
        ${!this.validity.valid
          ? html`<div class="error">${this.validationMessage}</div>`
          : ''}
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-field': DemoField;
  }
}
