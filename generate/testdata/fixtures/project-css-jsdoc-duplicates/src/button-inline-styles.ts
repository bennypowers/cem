import { LitElement, css } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * @cssprop --button-bg - Background color of the button
 * @cssprop --button-color - Text color
 * @cssprop --button-border - Border style
 * @cssprop --button-radius - Border radius
 * @cssprop --button-padding - Internal padding
 */
@customElement('button-inline-styles')
class ButtonInlineStyles extends LitElement {
  static styles = css`
    .button {
      padding: var(--button-padding, 0.5rem 1rem);
      background: var(--button-bg, #0070f3);
      color: var(--button-color, white);
      border: var(--button-border, none);
      border-radius: var(--button-radius, 4px);
    }
  `;
}
