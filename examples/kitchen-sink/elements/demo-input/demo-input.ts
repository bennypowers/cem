import { LitElement, html } from 'lit';
import { customElement, property, query } from 'lit/decorators.js';
import { classMap } from 'lit/directives/class-map.js';
import { ifDefined } from 'lit/directives/if-defined.js';
import styles from './demo-input.css' with { type: 'css' };

export type InputType =
  | 'text'
  | 'email'
  | 'password'
  | 'number'
  | 'tel'
  | 'url';

/**
 * Custom event for demo-input component value changes.
 */
export class DemoInputEvent extends Event {
  value: string;
  constructor(type: string, value: string) {
    super(type, { bubbles: true, composed: true });
    this.value = value;
  }
}

/**
 * A comprehensive text input component with validation and helper text.
 *
 * @summary Text input with label, validation, and accessibility features
 *
 * @fires demo-input - Emitted when the input value changes (on input event)
 * @fires demo-change - Emitted when the input value is committed (on change event)
 */
@customElement('demo-input')
export class DemoInput extends LitElement {
  static styles = [styles];

  @query('#input') accessor input!: HTMLInputElement;

  /**
   * The input's label text
   */
  @property() accessor label = '';

  /**
   * The current value of the input
   */
  @property() accessor value = '';

  /**
   * Placeholder text shown when the input is empty
   */
  @property() accessor placeholder = '';

  /**
   * The input type
   * @type {'text' | 'email' | 'password' | 'number' | 'tel' | 'url'}
   */
  @property() accessor type: InputType = 'text';

  /**
   * Disables the input
   */
  @property({ type: Boolean, reflect: true }) accessor disabled = false;

  /**
   * Marks the input as required for form validation
   */
  @property({ type: Boolean, reflect: true }) accessor required = false;

  /**
   * Makes the input read-only
   */
  @property({ type: Boolean, reflect: true }) accessor readonly = false;

  /**
   * The input's name when used in forms
   */
  @property() accessor name = '';

  /**
   * Autocomplete hint for the browser
   */
  @property() accessor autocomplete = '';

  /**
   * Helper text displayed below the input
   */
  @property({ attribute: 'helper-text' }) accessor helperText = '';

  /**
   * Error text displayed when the input is in an error state
   */
  @property({ attribute: 'error-text' }) accessor errorText = '';

  /**
   * Shows the input in an error state
   */
  @property({ type: Boolean, reflect: true }) accessor error = false;

  #onInput(e: Event) {
    this.value = (e.target as HTMLInputElement).value;
    this.dispatchEvent(new DemoInputEvent('demo-input', this.value));
  }

  #onChange(e: Event) {
    this.value = (e.target as HTMLInputElement).value;
    this.dispatchEvent(new DemoInputEvent('demo-change', this.value));
  }

  focus(options?: FocusOptions) {
    this.input?.focus(options);
  }

  blur() {
    this.input?.blur();
  }

  select() {
    this.input?.select();
  }

  render() {
    const {
      autocomplete,
      disabled,
      helperText,
      error,
      errorText,
      label,
      name,
      placeholder,
      readonly,
      required,
      type,
      value,
    } = this;
    const displayHelperText = error && errorText ? errorText : helperText;

    return html`
      <div id="wrapper">${!label ? '' : html`
        <!-- The input's label element -->
        <label id="label" part="label">${label}${!required ? '' : html`
          <span aria-label="required">*</span>`}
        </label>`}
        <!-- The native input element -->
        <input id="input"
               part="input"
               class="${classMap({ error })}"
               type="${type}"
               .value="${value}"
               placeholder="${ifDefined(placeholder || undefined)}"
               ?disabled="${disabled}"
               ?required="${required}"
               ?readonly="${readonly}"
               name="${ifDefined(name || undefined)}"
               autocomplete="${ifDefined(autocomplete || undefined)}"
               @input="${this.#onInput}"
               @change="${this.#onChange}">${!displayHelperText ? '' : html`
        <!-- The helper or error text container -->
        <div id="helper-text"
             part="helper-text"
             class="${classMap({ error })}">
          ${displayHelperText}
        </div>`}
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-input': DemoInput;
  }
}
