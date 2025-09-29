import { html } from 'lit';

export function renderCard() {
  return html`
    <my-card variant="primary">
      <span slot="header">Title</span>
    </my-card>
  `;
}