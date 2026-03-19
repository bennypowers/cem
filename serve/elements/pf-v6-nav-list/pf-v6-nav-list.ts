import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';

import styles from './pf-v6-nav-list.css' with { type: 'css' };

/**
 * PatternFly v6 Navigation List
 *
 * @slot - Default slot for nav-item or nav-group elements
 */
@customElement('pf-v6-nav-list')
export class PfV6NavList extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  constructor() {
    super();
    this.#internals.role = 'list';
  }

  render() {
    return html`<slot></slot>`;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-nav-list': PfV6NavList;
  }
}
