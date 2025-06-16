import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element tag-list
 * @summary Displays a list of tags.
 */
@customElement('tag-list')
export class TagList extends LitElement {
  @property({ type: Array }) tags: string[] = [];

  static styles = css`
    .tag { background: #eee; border-radius: 2em; padding: 0.2em 1em; margin: 0 0.2em; display: inline-block; }
  `;

  addTag(tag: string) {
    this.tags = [...this.tags, tag];
  }

  removeTag(tag: string) {
    this.tags = this.tags.filter(t => t !== tag);
  }

  render() {
    return html`
      ${this.tags.map(tag => html`
        <span class="tag" @click=${() => this.removeTag(tag)}>${tag}</span>
      `)}
    `;
  }
}