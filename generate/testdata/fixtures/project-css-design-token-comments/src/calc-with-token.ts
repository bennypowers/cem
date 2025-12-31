import { LitElement, css } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * Element that demonstrates CSS comments in calc() with design tokens
 * @customElement calc-with-token
 */
@customElement('calc-with-token')
export class CalcWithToken extends LitElement {
  static styles = css`
    :host {
      --_step-template-rows:
        calc(
          /** @summary Base icon container size */
          var(--cem-length-xl, 24px) + 4px
        ) auto 1fr;
    }
  `;
}
