import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import styles from './demo-card.css' with { type: 'css' };

@customElement('demo-card')
export class DemoCard extends LitElement {
  static styles = [styles];

  render() {
    return html`
      <div id="header" part="header">
        <!--
          summary: Header slot
          description: |
            Place header content here, such as a title or actions.
            This section is visually separated from the body.
        -->
        <slot name="header"></slot>
      </div>
      <div id="body" part="body">
        <!-- The main content area of the card -->
        <slot></slot>
      </div>
      <div id="footer" part="footer">
        <!--
          summary: Footer slot
          description: Place footer content like actions or metadata here
        -->
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
