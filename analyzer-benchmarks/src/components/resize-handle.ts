import { LitElement, html, css, customElement } from 'lit-element';

/**
 * @element resize-handle
 * @summary Drag handle for resizing layouts.
 */
@customElement('resize-handle')
export class ResizeHandle extends LitElement {
  static styles = css`
    .handle {
      width: 8px; height: 100%;
      background: #bbb;
      cursor: ew-resize;
      user-select: none;
      position: absolute; right: 0; top: 0;
    }
  `;

  render() {
    return html`<div class="handle"></div>`;
  }
}