import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element field-label
 * @summary Label for form fields, with optional required marker.
 */
@customElement('field-label')
export class FieldLabel extends LitElement {
  @property({ type: String }) for = '';
  @property({ type: Boolean }) required = false;

  static styles = css`
    label { font-weight: bold; }
    .required { color: red; margin-left: 0.25em; }
  `;

  render() {
    return html`
      <label for=${this.for}>
        <slot></slot>
        ${this.required ? html`<span class="required">*</span>` : ''}
      </label>
    `;
  }
}