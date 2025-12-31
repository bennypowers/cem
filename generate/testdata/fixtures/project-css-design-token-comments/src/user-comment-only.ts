import { LitElement, css } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * Element that has CSS comments but no matching design tokens
 * @customElement user-comment-only
 */
@customElement('user-comment-only')
export class UserCommentOnly extends LitElement {
  static styles = css`
    :host {
      /** USER description for a custom property without design token */
      color: var(--my-custom-color);
      background: var(--_my-private,
          /** USER comment for custom background */
          var(--my-custom-background));
    }
  `;
}
