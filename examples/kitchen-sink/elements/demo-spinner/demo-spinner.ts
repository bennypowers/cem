import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';
import styles from './demo-spinner.css' with { type: 'css' };

/**
 * A loading spinner component.
 *
 * @summary Displays an animated loading indicator
 * @since 1.0
 * @status stable
 *
 * @csspart base - The spinner's base wrapper
 */
@customElement('demo-spinner')
export class DemoSpinner extends LitElement {
  static styles = [styles];

  render() {
    return html`<div id="spinner" part="base"></div>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-spinner': DemoSpinner;
  }
}
