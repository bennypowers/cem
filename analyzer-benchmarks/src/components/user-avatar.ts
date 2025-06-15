import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element user-avatar
 * @summary Displays a user's avatar image and name.
 */
@customElement('user-avatar')
export class UserAvatar extends LitElement {
  @property({ type: String }) name = '';
  @property({ type: String }) src = '';

  static styles = css`
    .avatar { border-radius: 50%; width: 40px; height: 40px; vertical-align: middle; }
    .username { margin-left: 0.5em; }
  `;

  render() {
    return html`
      <img class="avatar" src=${this.src} alt=${this.name}>
      <span class="username">${this.name}</span>
    `;
  }
}