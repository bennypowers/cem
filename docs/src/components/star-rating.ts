import { LitElement, html, property, customElement } from 'lit-element';

@customElement('star-rating')
export class StarRating extends LitElement {
  @property({ type: Number }) rating = 0;

  render() {
    return html`
      ${[1, 2, 3, 4, 5].map(
        n => html`<span @click=${() => this._setRating(n)}>
          ${n <= this.rating ? '★' : '☆'}
        </span>`
      )}
    `;
  }

  private _setRating(n: number) {
    this.rating = n;
    this.dispatchEvent(new CustomEvent('rating-changed', { detail: n }));
  }
}
