import { LitElement, html, css } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';
import { classMap } from 'lit/directives/class-map.js';

/**
 * A button component with multiple variants and sizes.
 *
 * @summary Buttons represent actions available to the user
 * @status stable
 *
 * @dependency ui-spinner
 *
 * @slot - The button's label
 * @slot start - Content placed before the label (e.g., an icon)
 * @slot end - Content placed after the label (e.g., an icon)
 *
 * @event click - Emitted when the button is clicked
 * @event focus - Emitted when the button gains focus
 * @event blur - Emitted when the button loses focus
 *
 * @csspart base - The button's base wrapper element
 * @csspart label - The button's label
 * @csspart start - The start slot container
 * @csspart end - The end slot container
 *
 * @cssprop --button-bg - Background color of the button
 * @cssprop --button-color - Text color
 * @cssprop --button-border - Border style
 * @cssprop --button-radius - Border radius
 * @cssprop --button-padding - Internal padding
 */
@customElement('ui-button')
export class UiButton extends LitElement {
  static styles = css`
    :host {
      display: inline-block;
    }

    .button {
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
      padding: var(--button-padding, 0.5rem 1rem);
      background: var(--button-bg, #0070f3);
      color: var(--button-color, white);
      border: var(--button-border, none);
      border-radius: var(--button-radius, 4px);
      font-family: inherit;
      font-size: 1rem;
      cursor: pointer;
      transition: all 0.2s ease;
    }

    .button:hover:not(.disabled) {
      opacity: 0.9;
    }

    .button:active:not(.disabled) {
      transform: scale(0.98);
    }

    .button.disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }

    /* Variants */
    .button.variant-primary {
      --button-bg: #0070f3;
      --button-color: white;
    }

    .button.variant-secondary {
      --button-bg: #666;
      --button-color: white;
    }

    .button.variant-success {
      --button-bg: #10b981;
      --button-color: white;
    }

    .button.variant-danger {
      --button-bg: #ef4444;
      --button-color: white;
    }

    /* Sizes */
    .button.size-small {
      --button-padding: 0.25rem 0.75rem;
      font-size: 0.875rem;
    }

    .button.size-medium {
      --button-padding: 0.5rem 1rem;
      font-size: 1rem;
    }

    .button.size-large {
      --button-padding: 0.75rem 1.5rem;
      font-size: 1.125rem;
    }

    .button.loading {
      position: relative;
      color: transparent;
    }

    .spinner {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 1em;
      height: 1em;
      border: 2px solid currentColor;
      border-top-color: transparent;
      border-radius: 50%;
      animation: spin 0.6s linear infinite;
      color: var(--button-color);
    }

    @keyframes spin {
      to { transform: translate(-50%, -50%) rotate(360deg); }
    }

    [hidden] {
      display: none;
    }
  `;

  /**
   * The button variant
   */
  @property({ reflect: true }) variant: 'primary' | 'secondary' | 'success' | 'danger' = 'primary';

  /**
   * The button size
   */
  @property({ reflect: true }) size: 'small' | 'medium' | 'large' = 'medium';

  /**
   * Whether the button is disabled
   */
  @property({ type: Boolean, reflect: true }) disabled = false;

  /**
   * Whether the button is in a loading state
   */
  @property({ type: Boolean, reflect: true }) loading = false;

  render() {
    return html`
      <button
        part="base"
        class=${classMap({
          button: true,
          [`variant-${this.variant}`]: true,
          [`size-${this.size}`]: true,
          disabled: this.disabled,
          loading: this.loading,
        })}
        ?disabled=${this.disabled || this.loading}
      >
        <slot name="start" part="start"></slot>
        <slot part="label"></slot>
        <slot name="end" part="end"></slot>
        <div class="spinner" ?hidden=${!this.loading}></div>
      </button>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'ui-button': UiButton;
  }
}
