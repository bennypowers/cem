import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import { classMap } from 'lit/directives/class-map.js';

import './ui-spinner.js';

import styles from './ui-button.css' with { type: 'css' };

export type ButtonSize = 'small' | 'medium' | 'large';

export type ButtonVariant = 'primary' | 'secondary' | 'success' | 'danger';

/**
 * A button component with multiple variants and sizes.
 *
 * @summary Buttons represent actions available to the user
 * @status stable
 *
 * @event click - Emitted when the button is clicked
 * @event focus - Emitted when the button gains focus
 * @event blur - Emitted when the button loses focus
 *
 */
@customElement('ui-button')
export class UiButton extends LitElement {
  static styles = [styles];

  /**
   * The button variant
   */
  @property({ reflect: true }) variant: ButtonVariant = 'primary';

  /**
   * The button size
   */
  @property({ reflect: true }) size: ButtonSize = 'medium';

  /**
   * Whether the button is disabled
   */
  @property({ type: Boolean, reflect: true }) disabled = false;

  /**
   * Whether the button is in a loading state
   */
  @property({ type: Boolean, reflect: true }) loading = false;

  render() {
    const { disabled, loading, size, variant } = this;
    return html`
      <!-- The button's base wrapper element -->
      <button id="button"
              part="base"
              class="${classMap({ [variant]: !!variant, [size]: !!size, disabled, loading })}"
              ?disabled="${disabled || loading}">
        <!-- slot:
               summary: Content placed before the label (e.g., an icon)
             part:
               summary: The start slot container -->
        <slot name="start" part="start"></slot>
        <!-- slot:
               summary: The button's label
             part:
               summary: The button's label -->
        <slot part="label"></slot>
        <!-- slot:
               summary: Content placed after the label (e.g., an icon)
             part:
               summary: The end slot container -->
        <slot name="end" part="end"></slot>
        <ui-spinner id="spinner" ?hidden="${!loading}"></ui-spinner>
      </button>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'ui-button': UiButton;
  }
}
