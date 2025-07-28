import { LitElement, html, property, customElement } from 'lit-element';

@customElement('color-picker')
export class ColorPicker extends LitElement {
  @property({ type: String }) selected = '#ff0000';

  render() {
    return html`
      <input type="color" .value=${this.selected} @input=${this.#onInput} />
      <span>Selected: ${this.selected}</span>
    `;
  }

  #onInput(e: Event) {
    this.selected = (e.target as HTMLInputElement).value;
    this.dispatchEvent(new CustomEvent('color-changed', { detail: this.selected }));
  }
}
