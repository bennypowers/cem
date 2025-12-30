import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A versatile card component with header, footer, and content slots.
 * Perfect for displaying structured content with optional theming.
 *
 * @fires expand - Fired when the card is expanded
 * @fires collapse - Fired when the card is collapsed
 */
@customElement('card-element')
export class CardElement extends LitElement {
  static styles = css`
    :host {
      display: block;
      /** Background color of the card */
      background: var(--card-background, #ffffff);
      /** Border radius of the card */
      border-radius: var(--card-border-radius, 8px);
      /** Internal padding */
      padding: var(--card-padding, 16px);
      /** Box shadow effect */
      box-shadow: var(--card-shadow, 0 2px 8px rgba(0,0,0,0.1));
    }

    .card {
      display: flex;
      flex-direction: column;
      height: 100%;
    }

    .header {
      border-bottom: 1px solid #eee;
      padding-bottom: 12px;
      margin-bottom: 16px;
    }

    .content {
      flex: 1;
    }

    .footer {
      border-top: 1px solid #eee;
      padding-top: 12px;
      margin-top: 16px;
    }

    .actions {
      display: flex;
      gap: 8px;
      justify-content: flex-end;
      margin-top: 16px;
    }
  `;

  /**
   * The visual theme variant of the card
   */
  @property({ type: String }) theme: 'default' | 'primary' | 'success' | 'warning' | 'danger' = 'default';

  /**
   * Whether the card should have an elevated appearance
   */
  @property({ type: Boolean }) elevated = false;

  /**
   * Whether the card content should be centered
   */
  @property({ type: Boolean }) centered = false;

  render() {
    return html`
      <!-- The main card container -->
      <div 
        class="card" 
        part="card"
        data-theme="${this.theme}"
        ?elevated="${this.elevated}"
        ?centered="${this.centered}"
      >
        <!-- The header section -->
        <div class="header" part="header">
          <!-- Header content area -->
          <slot name="header"></slot>
        </div>
        
        <!-- The main content area -->
        <div class="content" part="content">
          <!-- Default slot for main content -->
          <slot></slot>
        </div>
        
        <!-- The footer section -->
        <div class="footer" part="footer">
          <!-- Footer content area -->
          <slot name="footer"></slot>
        </div>
        
        <!-- Action buttons area -->
        <div class="actions" part="actions">
          <slot name="actions"></slot>
        </div>
      </div>
    `;
  }
}