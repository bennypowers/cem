import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element inline-edit
 * @summary Allows inline editing of a value.
 */
@customElement('inline-edit')
export class InlineEdit extends LitElement {
  @property({ type: String }) value = '';
  @property({ type: Boolean }) editing = false;

  static styles = css`
    .edit { border: 1px solid #2196f3; }
    input { font-size: inherit; }
  `;

  get displayValue() {
    return this.value || '(empty)';
  }

  #startEdit() { this.editing = true; }
  #onInput(e: Event) {
    this.value = (e.target as HTMLInputElement).value;
  }
  #onBlur() { this.editing = false; }

  render() {
    return this.editing
      ? html`<input class="edit" .value=${this.value} @input=${this.#onInput} @blur=${this.#onBlur}>`
      : html`<span @click=${this.#startEdit}>${this.displayValue}</span>`;
  }
}