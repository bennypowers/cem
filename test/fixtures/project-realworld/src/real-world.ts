import styles from './real-world.css';
@customElement('real-world')
export class RealWorld extends LitElement {
  static styles = [styles];

  @property({ reflect: true, attribute: 'color-palette' }) colorPalette?: ColorPalette;

  /**
   * variant
   */
  @property({ reflect: true }) variant?: 'promo';

  /**
   * is full-width
   */
  @property({ reflect: true, attribute: 'full-width', type: Boolean }) fullWidth = false;

  willUpdate() {
  }
}

