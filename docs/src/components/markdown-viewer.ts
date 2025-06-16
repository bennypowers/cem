import { LitElement, html, property, customElement } from 'lit-element';

@customElement('markdown-viewer')
export class MarkdownViewer extends LitElement {
  @property({ type: String }) markdown = '';

  render() {
    return html`
      <div class="markdown">${this.markdown}</div>
    `;
  }
}
