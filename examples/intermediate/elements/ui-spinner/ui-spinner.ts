import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import { styleMap } from 'lit/directives/style-map.js';

import styles from './ui-spinner.css' with { type: 'css' };

/**
 * A loading spinner component for indicating progress.
 *
 * @summary Spinners provide visual feedback that an action is processing
 */
@customElement('ui-spinner')
export class UiSpinner extends LitElement {
  static styles = [styles];

  /**
   * Accessible label for the spinner
   */
  @property() label = 'Loading';

  /**
   * The size of the spinner in pixels
   */
  @property({ type: Number }) size?: number;

  render() {
    const { size, label } = this;
    const spinnerSize = size != null ? `${size}px` : undefined;

    return html`
      <!-- The spinner's base element -->
      <div id="spinner"
           part="base"
           role="status"
           aria-label="${label}"
           style="${styleMap({ '--spinner-size': spinnerSize })}"></div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'ui-spinner': UiSpinner;
  }
}
