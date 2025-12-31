import { LitElement, css } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * Element that has empty comments with design tokens
 * @customElement empty-comment
 */
@customElement('empty-comment')
export class EmptyComment extends LitElement {
  static styles = css`
    :host {
      /** */
      color: var(--cem-color-primary);
      /**  */
      margin: var(--cem-spacing-base);
    }
  `;
}
