import { LitElement, html, property, customElement } from 'lit-element';

@customElement('notification-banner')
export class NotificationBanner extends LitElement {
  @property({ type: String }) message = '';
  @property({ type: Boolean }) visible = false;

  render() {
    return this.visible
      ? html`<div class="banner">${this.message}</div>`
      : html``;
  }
}
