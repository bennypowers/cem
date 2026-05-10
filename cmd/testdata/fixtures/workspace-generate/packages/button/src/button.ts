/**
 * A button
 * @summary Test button
 * @tag test-button
 */
export class TestButton extends HTMLElement {
  /** Button label */
  label = 'Click';
}

customElements.define('test-button', TestButton);
