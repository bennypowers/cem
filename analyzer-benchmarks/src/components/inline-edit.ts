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

  _startEdit() { this.editing = true; }
  _onInput(e: Event) {
    this.value = (e.target as HTMLInputElement).value;
  }
  _onBlur() { this.editing = false; }

  render() {
    return this.editing
      ? html`<input class="edit" .value=${this.value} @input=${this._onInput} @blur=${this._onBlur}>`
      : html`<span @click=${this._startEdit}>${this.displayValue}</span>`;
  }
}