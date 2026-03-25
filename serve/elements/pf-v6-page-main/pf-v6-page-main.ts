import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';

import styles from './pf-v6-page-main.css' with { type: 'css' };

/**
 * PatternFly v6 Page Main Content
 *
 * Wraps main content area with proper semantics and skip-to-content support.
 * The id attribute should be set on the host element for skip-to-content links.
 *
 * @slot - Default slot for main content
 */
@customElement('pf-v6-page-main')
export class PfV6PageMain extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  constructor() {
    super();
    this.#internals.role = 'main';
  }

  render() {
    return html`<slot></slot>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-page-main': PfV6PageMain;
  }
}
