import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A customizable button component with multiple variants and states.
 *
 * @fires click - Fired when the button is clicked
 * @fires focus - Fired when the button receives focus
 * @fires blur - Fired when the button loses focus
 */
@customElement('button-element')
export class ButtonElement extends LitElement {
  static styles = css`
    :host {
      display: inline-block;
    }

    button {
      display: flex;
      align-items: center;
      gap: 8px;
      /** Internal padding */
      padding: var(--button-padding, 8px 16px);
      border: none;
      /** Border radius */
      border-radius: var(--button-border-radius, 4px);
      /** Background color */
      background: var(--button-background, #007bff);
      /** Text color */
      color: var(--button-color, white);
      /** Font size */
      font-size: var(--button-font-size, 14px);
      font-family: inherit;
      cursor: pointer;
      transition: all 0.2s ease;
    }

    button:hover {
      filter: brightness(1.1); }

    button:active {
      transform: translateY(1px);
    }

    button:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }

    .icon {
      display: flex;
      align-items: center;
    }
  `;

  /**
   * The button variant style
   */
  @property() variant: 'primary' | 'secondary' | 'success' | 'danger' | 'warning' = 'primary';

  /**
   * The button size
   */
  @property() size: 'small' | 'medium' | 'large' = 'medium';

  /**
   * Whether the button is disabled
   */
  @property({ type: Boolean }) disabled = false;

  /**
   * Whether the button should take full width
   */
  @property({ type: Boolean, attribute: 'full-width' }) fullWidth = false;

  /**
   * Whether the button is in a loading state
   */
  @property({ type: Boolean }) loading = false;

  render() {
    return html`
      <!-- The button element -->
      <button part="button"
              ?disabled="${this.disabled || this.loading}"
              data-variant="${this.variant}"
              data-size="${this.size}"
              ?full-width="${this.fullWidth}"
              ?loading="${this.loading}"
              @click="${this.#onClick}">
        <!-- The icon container -->
        <div class="icon" part="icon">
          <!-- Icon slot positioned before text -->
          <slot name="icon"></slot>
        </div>
        <!-- The text content -->
        <span part="text">
          <!-- Button content (text, icons, etc.) -->
          <slot></slot>
        </span>
      </button>
    `;
  }

  #onClick(event: Event) {
    if (this.disabled || this.loading) {
      event.preventDefault();
      event.stopPropagation();
      return;
    }

    this.dispatchEvent(new CustomEvent('click', {
      bubbles: true,
      composed: true,
      detail: { originalEvent: event }
    }));
  }
}
