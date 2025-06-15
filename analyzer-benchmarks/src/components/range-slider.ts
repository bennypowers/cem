import { LitElement, html, property, customElement } from 'lit-element';

@customElement('range-slider')
export class RangeSlider extends LitElement {
  @property({ type: Number }) value = 50;
  @property({ type: Number }) min = 0;
  @property({ type: Number }) max = 100;

  render() {
    return html`
      <input
        type="range"
        .value=${this.value}
        min=${this.min}
        max=${this.max}
        @input=${this._onInput}
      />
      <span>${this.value}</span>
    `;
  }

  private _onInput(e: Event) {
    this.value = Number((e.target as HTMLInputElement).value);
    this.dispatchEvent(new CustomEvent('range-changed', { detail: this.value }));
  }
}
