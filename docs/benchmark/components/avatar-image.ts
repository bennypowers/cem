import { LitElement, html, css } from 'lit';
import { property, customElement } from 'lit/decorators.js';

/**
 * Shows an avatar image with fallback.
 */
@customElement('avatar-image')
export class AvatarImage extends LitElement {
  @property({ type: String }) src = '';
  @property({ type: String }) alt = 'Avatar';

  static styles = css`
    img { border-radius: 50%; width: 56px; height: 56px; }
  `;

  render() {
    return html`<img src="${this.src}" alt="${this.alt}" @error="${this.#onError}" />`;
  }

  #onError(e: Event) {
    (e.target as HTMLImageElement).src = 'default-avatar.png';
  }
}
