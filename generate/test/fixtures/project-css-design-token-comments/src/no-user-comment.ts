import { LitElement, css } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * Element that uses design tokens without user comments
 * @customElement no-user-comment
 */
@customElement('no-user-comment')
export class NoUserComment extends LitElement {
  static styles = css`
    :host {
      color: var(--cem-color-primary);
      padding: var(--cem-spacing-base);
    }
  `;
}
