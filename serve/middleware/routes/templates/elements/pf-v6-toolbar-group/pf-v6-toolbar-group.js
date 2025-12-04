/**
 * PatternFly v6 Toolbar Group Component
 * @customElement pf-v6-toolbar-group
 */

import { CemElement } from '/__cem/cem-element.js';

export class PfV6ToolbarGroup extends CemElement {
  static is = 'pf-v6-toolbar-group';

  static {
    customElements.define(this.is, this);
  }
}

