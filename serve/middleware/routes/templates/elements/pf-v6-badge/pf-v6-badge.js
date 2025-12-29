import { CemElement } from '/__cem/cem-element.js';

/**
 * PatternFly v6 Badge Component
 *
 * @attr {boolean} read - Sets the badge to read state (gray)
 * @attr {boolean} unread - Sets the badge to unread state (blue)
 * @attr {boolean} disabled - Sets the badge to disabled state
 *
 * @customElement pf-v6-badge
 */
export class PfV6Badge extends CemElement {
  static is = 'pf-v6-badge';

  static observedAttributes = [
    'read',
    'unread',
    'disabled'
  ];

  get read() { return this.hasAttribute('read'); }
  set read(value) { this.toggleAttribute('read', !!value); }

  get unread() { return this.hasAttribute('unread'); }
  set unread(value) { this.toggleAttribute('unread', !!value); }

  get disabled() { return this.hasAttribute('disabled'); }
  set disabled(value) { this.toggleAttribute('disabled', !!value); }

  static {
    customElements.define(this.is, this);
  }
}
