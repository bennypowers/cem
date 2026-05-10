/**
 * A simple button component
 * @summary A button for clicking
 * @customElement example-button
 * @csspart button - The button element
 */
export class ExampleButton extends HTMLElement {
  /** The button label */
  label = 'Click me';

  /** Whether the button is disabled */
  disabled = false;
}

customElements.define('example-button', ExampleButton);
