import { LitElement, html, nothing } from 'lit';
import { customElement, property, query, state } from 'lit/decorators.js';
import { classMap } from 'lit/directives/class-map.js';
import styles from './demo-button.css' with { type: 'css' };
import '../demo-spinner/demo-spinner.js';

type ButtonType = 'button' | 'submit' | 'reset';

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
 * @csspart base - The component's base wrapper
 * @csspart start - The container that wraps the start slot
 * @csspart label - The button's label container
 * @csspart end - The container that wraps the end slot
 * @csspart spinner - The loading spinner element
 */
@customElement('demo-button')
export class DemoButton extends LitElement {
  static styles = [styles];

  @query('#button') button!: HTMLButtonElement;

  @state() private hasStartSlot = false;
  @state() private hasEndSlot = false;
  @state() private hasDefaultSlot = false;

  /**
   * The button's theme variant
   * @type {'neutral' | 'brand' | 'success' | 'warning' | 'danger'}
   */
  @property({ reflect: true }) variant: 'neutral' | 'brand' | 'success' | 'warning' | 'danger' = 'brand';

  /**
   * The button's visual appearance style
   * @type {'filled' | 'outlined' | 'plain'}
   */
  @property({ reflect: true }) appearance: 'filled' | 'outlined' | 'plain' = 'filled';

  /**
   * The button's size
   * @type {'small' | 'medium' | 'large'}
   */
  @property({ reflect: true }) size: 'small' | 'medium' | 'large' = 'medium';

  /**
   * Draws the button with a pill shape (fully rounded)
   */
  @property({ type: Boolean, reflect: true }) pill = false;

  /**
   * Disables the button
   */
  @property({ type: Boolean, reflect: true }) disabled = false;

  /**
   * Shows a loading spinner and disables interaction
   */
  @property({ type: Boolean, reflect: true }) loading = false;

  /**
   * The button type for form submission
   * @type {'button' | 'submit' | 'reset'}
   */
  @property() type: ButtonType = 'button';

  /**
   * The button's name when used in forms
   */
  @property() name = '';

  /**
   * The button's value when used in forms
   */
  @property() value = '';

  private onSlotchange() {
    this.hasStartSlot = this.querySelector('[slot="start"]') !== null;
    this.hasEndSlot = this.querySelector('[slot="end"]') !== null;
    this.hasDefaultSlot = Array.from(this.childNodes).some(
      node => !node.nodeValue?.trim() || (node as Element).slot === undefined
    );
  }

  /**
   * Simulates a click on the button
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

  render() {
    const { hasStartSlot, hasDefaultSlot, hasEndSlot, loading, disabled, pill, variant, appearance, size } = this;
    return html`
      <button id="button"
              part="base"
              class="${classMap({ [variant]: true, [appearance]: true, [size]: true, pill, disabled, loading, hasStartSlot, hasEndSlot, hasDefaultSlot })}"
              ?disabled="${disabled || loading}"
              type="${this.type}"
              name="${this.name || nothing}"
              value="${this.value || nothing}"
              aria-disabled="${String(disabled)}"
              tabindex="${disabled ? '-1' : '0'}">
        <slot name="start" part="start" @slotchange=${this.onSlotchange}></slot>
        <slot part="label" @slotchange=${this.onSlotchange}></slot>
        <slot name="end" part="end" @slotchange=${this.onSlotchange}></slot>
        ${loading ? html`<demo-spinner part="spinner"></demo-spinner>` : ''}
      </button>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-button': DemoButton;
  }
}
