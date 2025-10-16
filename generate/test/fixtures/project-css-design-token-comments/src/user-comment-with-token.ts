import { LitElement, css } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * Element that demonstrates CSS comments merging with design tokens
 * @customElement user-comment-with-token
 */
@customElement('user-comment-with-token')
export class UserCommentWithToken extends LitElement {
  static styles = css`
    :host {
      color: var(--_my-var,
        /** USER description for the primary color, specified on the token fallback for a _private var */
        var(--cem-color-primary));
      /** @summary comment on a private var */
      border-color: var(--_private);

      /**
       * USER description for secondary color
       * This is a multiline comment that explains
       * the usage of the secondary color token
       */
      background-color: var(--cem-color-secondary);

      /** @summary Outer comment on private var value assignment */
      --_private-with-outer-comment: var(--_another-private);

      font-weight: var(--_private-weight,
        /** @summary Nested comment on spacing token inside private var fallback */
        var(--cem-spacing-base));

      /** @summary Direct summary comment on design token */
      padding: var(--cem-spacing-base);
    }
  `;
}
