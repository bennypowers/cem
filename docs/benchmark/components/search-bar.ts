import { LitElement, html, property, customElement } from 'lit-element';

@customElement('search-bar')
export class SearchBar extends LitElement {
  @property({ type: String }) query = '';

  render() {
    return html`
      <input
        type="search"
        .value=${this.query}
        @input=${e => this._onInput(e)}
        placeholder="Search..."
      />
    `;
  }

  private _onInput(e: Event) {
    this.query = (e.target as HTMLInputElement).value;
    this.dispatchEvent(new CustomEvent('search', { detail: this.query }));
  }
}
