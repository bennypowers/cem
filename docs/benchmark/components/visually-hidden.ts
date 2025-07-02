import { LitElement, html, css, customElement } from 'lit-element';

/**
 * @element visually-hidden
 * @summary Content for screen readers only.
 */
@customElement('visually-hidden')
export class VisuallyHidden extends LitElement {
  static styles = css`
    :host { 
      position: absolute !important;
      width: 1px; height: 1px;
      padding: 0; margin: -1px;
      overflow: hidden;
      clip: rect(0,0,0,0);
      border: 0;
      white-space: nowrap;
    }
  `;
  render() {
    return html`<slot></slot>`;
  }
}