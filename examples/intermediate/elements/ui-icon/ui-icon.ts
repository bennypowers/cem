import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import { styleMap } from 'lit/directives/style-map.js';

import styles from './ui-icon.css' with { type: 'css' };

export type IconName =
  | 'check'
  | 'chevron-right'
  | 'heart'
  | 'star'
  | 'x';

/**
 * An icon component for displaying simple SVG icons.
 *
 * @summary A simple icon component with size control
 */
@customElement('ui-icon')
export class UiIcon extends LitElement {
  private static readonly ICONS = {
    check: 'M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z',
    'chevron-right': 'M10 6L8.59 7.41 13.17 12l-4.58 4.59L10 18l6-6z',
    heart: 'M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z',
    star: 'M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z',
    x: 'M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z',
  };

  static styles = [styles];

  private get iconPath(): string {
    return UiIcon.ICONS[this.name] || UiIcon.ICONS.check;
  }

  /**
   * The name of the icon to display
   * @type {'check' | 'x' | 'chevron-right' | 'heart' | 'star'}
   */
  @property() name: IconName = 'check';

  /**
   * Accessible label for the icon (required for accessibility)
   */
  @property() label = '';

  /**
   * The size of the icon in pixels
   */
  @property({ type: Number }) size?: number;

  connectedCallback() {
    super.connectedCallback();
    // Warn if no label is provided (accessibility concern)
    if (!this.label) {
      console.warn('ui-icon: Missing "label" attribute. Icons should have labels for screen readers.');
    }
  }

  render() {
    const { label, name, size, iconPath } = this;
    return html`
      <!-- The icon's base SVG element -->
      <svg part="base"
           viewBox="0 0 24 24"
           style="${styleMap({ '--icon-size': size })}"
           aria-label="${label || name}">
        <path d="${iconPath}"></path>
      </svg>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'ui-icon': UiIcon;
  }
}
