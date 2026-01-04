import { LitElement, html, svg } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import styles from './demo-icon.css' with { type: 'css' };

export type IconName =
  | 'alert'
  | 'check'
  | 'chevron'
  | 'heart'
  | 'info'
  | 'minus'
  | 'plus'
  | 'search'
  | 'star'
  | 'x';

const ICONS = {
  alert:   svg`<circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/>`,
  check:   svg`<path d="M20 6L9 17l-5-5"/>`,
  chevron: svg`<polyline points="6 9 12 15 18 9"/>`,
  heart:   svg`<path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>`,
  info:    svg`<circle cx="12" cy="12" r="10"/><line x1="12" y1="16" x2="12" y2="12"/><line x1="12" y1="8" x2="12.01" y2="8"/>`,
  minus:   svg`<line x1="5" y1="12" x2="19" y2="12"/>`,
  plus:    svg`<line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/>`,
  search:  svg`<circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/>`,
  star:    svg`<polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/>`,
  x:       svg`<path d="M18 6L6 18M6 6l12 12"/>`,
};

/**
 * An icon component for displaying simple SVG icons.
 *
 * @summary Displays SVG icons with customizable appearance
 * @status stable
 * @since 1.0
 *
 * @csspart svg - The icon's base SVG element
 */
@customElement('demo-icon')
export class DemoIcon extends LitElement {
  static styles = [styles];

  /**
   * The name of the icon to display
   * @type {'alert' | 'check' | 'chevron' | 'heart' | 'info' | 'minus' | 'plus' | 'search' | 'star' | 'x'}
   */
  @property() accessor name: keyof typeof ICONS = 'check';

  /**
   * Accessible label for the icon (required for accessibility)
   */
  @property() accessor label = '';

  connectedCallback() {
    super.connectedCallback();
    if (!this.label) {
      console.warn('demo-icon: Missing "label" attribute. Icons should have labels for screen readers.');
    }
  }

  render() {
    const { label, name } = this;
    return html`
      <!-- The icon's base SVG element -->
      <svg part="svg"
           viewBox="0 0 24 24"
           fill="none"
           stroke="currentColor"
           stroke-width="2"
           stroke-linecap="round"
           stroke-linejoin="round"
           role="img"
           aria-label="${label || name}">${ICONS[name] || ICONS.check}</svg>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-icon': DemoIcon;
  }
}
