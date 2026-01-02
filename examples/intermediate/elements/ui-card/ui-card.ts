import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A card component for grouping related content.
 *
 * @summary Cards contain content and actions about a single subject
 * @status stable
 *
 * @slot - The card's main content
 * @slot header - Content for the card header
 * @slot footer - Content for the card footer (e.g., actions)
 *
 * @csspart base - The card's base wrapper
 * @csspart header - The header container
 * @csspart body - The body/content container
 * @csspart footer - The footer container
 *
 * @cssprop --card-bg - Background color of the card
 * @cssprop --card-border - Border style
 * @cssprop --card-radius - Border radius
 * @cssprop --card-padding - Internal padding
 * @cssprop --card-shadow - Box shadow for elevation
 */
@customElement('ui-card')
export class UiCard extends LitElement {
  static styles = css`
    :host {
      display: block;
    }

    .card {
      background: var(--card-bg, white);
      border: var(--card-border, 1px solid #e5e5e5);
      border-radius: var(--card-radius, 8px);
      box-shadow: var(--card-shadow, 0 2px 4px rgba(0, 0, 0, 0.1));
      overflow: hidden;
    }

    .header,
    .body,
    .footer {
      padding: var(--card-padding, 1rem);
    }

    .header {
      border-bottom: 1px solid #e5e5e5;
      font-weight: 600;
    }

    .header:empty {
      display: none;
    }

    .body {
      flex: 1;
    }

    .footer {
      border-top: 1px solid #e5e5e5;
      background: #f9f9f9;
      display: flex;
      gap: 0.5rem;
      align-items: center;
    }

    .footer:empty {
      display: none;
    }
  `;

  /**
   * The card heading text (alternative to header slot)
   */
  @property() heading?: string;

  render() {
    return html`
      <div part="base" class="card">
        <div part="header" class="header">
          ${this.heading ? html`<div>${this.heading}</div>` : ''}
          <slot name="header"></slot>
        </div>
        <div part="body" class="body">
          <slot></slot>
        </div>
        <div part="footer" class="footer">
          <slot name="footer"></slot>
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'ui-card': UiCard;
  }
}
