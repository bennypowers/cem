import { LitElement, css } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * Component with design token properties
 * @customElement test-component
 */
@customElement('test-component')
export class TestComponent extends LitElement {
  static styles = css`
    :host {
      /** User description for primary color */
      color: var(--cem-color-primary);
      /** User description for base spacing */
      padding: var(--cem-spacing-base);
    }
  `;
}