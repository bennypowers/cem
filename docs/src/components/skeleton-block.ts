import { LitElement, html, css, customElement } from 'lit-element';

/**
 * @element skeleton-block
 * @summary Placeholder skeleton for loading content.
 */
@customElement('skeleton-block')
export class SkeletonBlock extends LitElement {
  static styles = css`
    .skeleton {
      background: linear-gradient(90deg,#eee 25%,#f5f5f5 50%,#eee 75%);
      height: 1.5em;
      border-radius: 0.5em;
      animation: shimmer 1.2s infinite linear;
      margin: 0.2em 0;
    }
    @keyframes shimmer {
      0% { background-position: -100px 0; }
      100% { background-position: 100px 0; }
    }
  `;

  render() {
    return html`<div class="skeleton"></div>`;
  }
}