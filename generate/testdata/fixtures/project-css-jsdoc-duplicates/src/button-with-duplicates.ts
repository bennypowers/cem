import { LitElement } from 'lit';
import { customElement } from 'lit/decorators.js';
import styles from './button-with-duplicates.css';

/**
 * @cssprop --button-bg - Background color of the button
 * @cssprop --button-color - Text color
 * @cssprop --button-border - Border style
 * @cssprop --button-radius - Border radius
 * @cssprop --button-padding - Internal padding
 */
@customElement('button-with-duplicates')
class ButtonWithDuplicates extends LitElement {
  static styles = styles;
}
