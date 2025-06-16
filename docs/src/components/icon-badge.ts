import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element icon-badge
 * @summary Displays an icon with an overlaid badge counter.
 */
@customElement('icon-badge')
export class IconBadge extends LitElement {
  @property({ type: String }) icon = '';
  @property({ type: Number }) count = 0;

  static styles = css`
    .icon-wrap { position: relative; display: inline-block; }
    .badge {
      position: absolute; top: -8px; right: -8px;
      background: red; color: white; border-radius: 50%; font-size: 0.75em;
      padding: 0 0.5em;
    }
  `;

  render() {
    return html`
      <span class="icon-wrap">
        <span>${this.icon}</span>
        ${this.count > 0 ? html`<span class="badge">${this.count}</span>` : ''}
      </span>
    `;
  }
}