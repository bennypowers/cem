import { LitElement, html, property, customElement } from 'lit-element';

@customElement('image-gallery')
export class ImageGallery extends LitElement {
  @property({ type: Array }) images: string[] = [];
  @property({ type: Number }) current = 0;

  render() {
    return html`
      <img src="${this.images[this.current] ?? ''}" alt="Image" />
      <button @click=${this.prev}>Previous</button>
      <button @click=${this.next}>Next</button>
    `;
  }

  prev = () => { if (this.current > 0) this.current--; }
  next = () => { if (this.current < this.images.length - 1) this.current++; }
}
