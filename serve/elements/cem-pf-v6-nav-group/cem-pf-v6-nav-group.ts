import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './cem-pf-v6-nav-group.css' with { type: 'css' };

/**
 * PatternFly v6 Navigation Group (Nested/Subnav)
 *
 * @slot - Default slot for nav-list containing nav-items
 */
@customElement('cem-pf-v6-nav-group')
export class PfV6NavGroup extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  constructor() {
    super();
    this.#internals.role = 'region';
  }

  render() {
    return html`
      <section id="subnav"
               part="subnav">
        <div id="list"
             role="list">
          <slot></slot>
        </div>
      </section>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-pf-v6-nav-group': PfV6NavGroup;
  }
}
