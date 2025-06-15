import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element message-bubble
 * @summary A chat message bubble.
 */
@customElement('message-bubble')
export class MessageBubble extends LitElement {
  @property({ type: String }) text = '';
  @property({ type: String }) align: 'left' | 'right' = 'left';

  static styles = css`
    .bubble {
      display: inline-block;
      max-width: 60%;
      padding: 0.5em 1em;
      border-radius: 1em;
      background: #e0e0e0;
      margin: 0.2em;
    }
    .right { background: #2196f3; color: white; float: right; }
    .left { float: left; }
  `;

  render() {
    return html`
      <div class="bubble ${this.align}">${this.text}</div>
    `;
  }
}