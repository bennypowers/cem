import { LitElement, html, property } from 'lit-element';

class BaseItem extends LitElement {
  @property({ type: Boolean }) enabled = true;
}

export class SpecialItem extends BaseItem {
  @property({ type: String }) note = 'Special';

  render() {
    return html`<span>${this.note} - enabled: ${this.enabled}</span>`;
  }
}

customElements.define('special-item', SpecialItem);
