import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('class-render-markdown-comment')
export class ClassRenderMarkdownComment extends LitElement {
  render() {
    return html`
      <!-- this comment contains a colon: it could easily break naive parsing -->
      <slot></slot>
    `;
  }
}
