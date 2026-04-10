import { LitElement, css } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * Element that demonstrates comments before declarations with calc/min/max wrapping var()
 * @customElement calc-outer-comment
 */
@customElement('calc-outer-comment')
export class CalcOuterComment extends LitElement {
  static styles = css`
    :host {
      /** @summary Outer comment on calc with no fallback */
      padding: calc(var(--cem-spacing-base) / 2);

      /** @summary Outer comment on calc with fallback */
      width: calc(var(--cem-length-xl, 24px) + 4px);
    }

    .inner {
      /** @summary Non-host calc outer comment */
      margin: calc(var(--cem-color-secondary) + 0);
    }

    @container style(--cem-color-primary) {
      :host {
        /** @summary Token inside container query */
        color: var(--cem-color-primary);
      }
    }

    @container sidebar style(--cem-length-xl) {
      .content {
        /** @summary Non-host token inside named container query */
        padding: calc(var(--cem-spacing-base) * 3);
      }
    }
  `;
}
