import { LitElement, html, css, property, customElement } from 'lit-element';

@customElement('text-ticker')
export class TextTicker extends LitElement {
  @property({ type: String }) text = 'Scrolling Text';

  static styles = css`
    .ticker {
      overflow: hidden;
      white-space: nowrap;
      animation: scroll 10s linear infinite;
    }
    @keyframes scroll {
      0% { transform: translateX(100%) }
      100% { transform: translateX(-100%) }
    }
  `;

  render() {
    return html`<div class="ticker">${this.text}</div>`;
  }
}
