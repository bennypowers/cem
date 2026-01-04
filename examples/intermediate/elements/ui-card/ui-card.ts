import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

import styles from './ui-card.css' with { type: 'css' };

/**
 * A card component for grouping related content.
 *
 * @summary Cards contain content and actions about a single subject
 */
@customElement('ui-card')
export class UiCard extends LitElement {
  static styles = [styles];

  /**
   * @summary The card heading text
   * The heading attribute's value is overrides by the `header` slot
   */
  @property()
  accessor heading = '';

  render() {
    return html`
      <!-- The card's base wrapper -->
      <div id="card" part="base">
        <!-- The header container -->
        <div id="header" part="header">
          ${!this.heading ? '' : html`<div>${this.heading}</div>`}
          <!-- Content for the card header -->
          <slot name="header"></slot>
        </div>
        <!-- The body/content container -->
        <div id="body" part="body">
          <!-- The card's main content -->
          <slot></slot>
        </div>
        <!-- The footer container -->
        <div id="footer" part="footer">
          <!-- Content for the card footer (e.g., actions) -->
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
