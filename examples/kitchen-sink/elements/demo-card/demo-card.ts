import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import styles from './demo-card.css' with { type: 'css' };

/**
 * A card component for grouping related content.
 *
 * @summary Cards contain content and actions about a single subject
 */
@customElement('demo-card')
export class DemoCard extends LitElement {
  static styles = [styles];

  render() {
    return html`
      <!-- The header container -->
      <div id="header" part="header">
        <!-- slot:
               summary: Content for the card header (e.g., title, actions)
             part:
               summary: The header container -->
        <slot name="header"></slot>
      </div>
      <!-- The body/content container -->
      <div id="body" part="body">
        <!-- slot:
               summary: The card's main content
             part:
               summary: The body/content container -->
        <slot></slot>
      </div>
      <!-- The footer container -->
      <div id="footer" part="footer">
        <!-- slot:
               summary: Content for the card footer (e.g., actions, metadata)
             part:
               summary: The footer container -->
        <slot name="footer"></slot>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-card': DemoCard;
  }
}
