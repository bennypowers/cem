/**
 * A button
 * @summary Test button
 * @customElement test-button
 */
export class TestButton extends HTMLElement {
  /** Button label */
  label = 'Click';
}

customElements.define('test-button', TestButton);
