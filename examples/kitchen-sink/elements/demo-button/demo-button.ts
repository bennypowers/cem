import { LitElement, html, nothing } from 'lit';
import { customElement, property, query } from 'lit/decorators.js';
import styles from './demo-button.css' with { type: 'css' };
import '../demo-spinner/demo-spinner.js';

export type ButtonType = 'button' | 'submit' | 'reset';

export type ButtonAppearance = 'filled' | 'outlined' | 'plain';

export type ButtonSize = 'small' | 'medium' | 'large';

export type ButtonVariant =
  | 'neutral'
  | 'brand'
  | 'success'
  | 'warning'
  | 'danger';

/**
 * Comprehensive button component demonstrating all CEM features.
 *
 * @summary Buttons represent actions that are available to the user
 * @documentation https://example.com/docs/components/demo-button
 * @status stable
 * @since 1.0
 *
 * @fires click - Emitted when the button is clicked
 * @fires focus - Emitted when the button gains focus
 * @fires blur - Emitted when the button loses focus
 * @fires invalid - Emitted when form validation fails
 *
 * @slot - The button's label
 * @slot start - Content placed before the label (e.g., an icon)
 * @slot end - Content placed after the label (e.g., an icon)
 *
 * @example <caption>Basic usage</caption>
 * ```html
 * <demo-button variant="brand">Click me</demo-button>
 * ```
 *
 * @cssstate loading - Applied when the button is in a loading state
 * @cssstate active - Applied when the button is actively pressed
 * @cssstate disabled - Applied when the button is disabled
 *
 * @csspart base - The component's base wrapper
 * @csspart start - The container that wraps the start slot
 * @csspart label - The button's label container
 * @csspart end - The container that wraps the end slot
 * @csspart spinner - The loading spinner element
 */
@customElement('demo-button')
export class DemoButton extends LitElement {
  static styles = [styles];

  #internals = this.attachInternals();

  constructor() {
    super();
    (globalThis as any)._elementInternals?.set(this, this.#internals);
  }

  @query('#button')
  accessor button!: HTMLButtonElement;

  /**
   * The button's theme variant
   * @type {'neutral' | 'brand' | 'success' | 'warning' | 'danger'}
   */
  @property({ reflect: true }) accessor variant: ButtonVariant = 'brand';

  /**
   * The button's visual appearance style
   * @type {'filled' | 'outlined' | 'plain'}
   */
  @property({ reflect: true }) accessor appearance: ButtonAppearance = 'filled';

  /**
   * The button's size
   * @type {'small' | 'medium' | 'large'}
   */
  @property({ reflect: true }) accessor size: ButtonSize = 'medium';

  /**
   * Draws the button with a pill shape (fully rounded)
   * @deprecated Use CSS border-radius instead
   */
  @property({ type: Boolean, reflect: true }) accessor pill = false;

  /**
   * Disables the button
   */
  @property({ type: Boolean, reflect: true }) accessor disabled = false;

  /**
   * Shows a loading spinner and disables interaction
   */
  @property({ type: Boolean, reflect: true }) accessor loading = false;

  /**
   * The button type for form submission
   * @type {'button' | 'submit' | 'reset'}
   */
  @property() accessor type: ButtonType = 'button';

  /**
   * The button's name when used in forms
   */
  @property() accessor name = '';

  /**
   * The button's value when used in forms
   */
  @property() accessor value = '';

  /**
   * Simulates a click on the button
   * @example
   * ```js
   * document.querySelector('demo-button').click();
   * ```
   */
  click() {
    this.button?.click();
  }

  /**
   * Sets focus on the button
   */
  focus(options?: FocusOptions) {
    this.button?.focus(options);
  }

  /**
   * Removes focus from the button
   */
  blur() {
    this.button?.blur();
  }

  /**
   * Checks form validity and reports any validation errors
   * @returns {boolean} True if the form is valid
   */
  checkValidity(): boolean {
    return this.button?.checkValidity() ?? true;
  }

  /**
   * Reports form validity to the user
   * @returns {boolean} True if the form is valid
   */
  reportValidity(): boolean {
    return this.button?.reportValidity() ?? true;
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('loading')) {
      this.#internals.states[this.loading ? 'add' : 'delete']('loading');
    }
    if (changed.has('disabled')) {
      this.#internals.states[this.disabled ? 'add' : 'delete']('disabled');
    }
  }

  render() {
    return html`
      <button id="button"
              part="base"
              ?disabled="${this.disabled || this.loading}"
              type="${this.type}"
              name="${this.name || nothing}"
              value="${this.value || nothing}"
              aria-disabled="${String(this.disabled)}"
              tabindex="${this.disabled ? '-1' : '0'}">
        <slot name="start" part="start"></slot>
        <slot part="label"></slot>
        <slot name="end" part="end"></slot>
        <demo-spinner part="spinner"></demo-spinner>
      </button>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-button': DemoButton;
  }
}
