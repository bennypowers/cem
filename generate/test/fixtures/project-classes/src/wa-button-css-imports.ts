import { customElement } from 'lit/decorators.js';
import styles from './button.css';
import sizeStyles from '../../styles/utilities/size.css';

class WebAwesomeFormAssociatedElement {}

/**
 * @summary Test with CSS imports
 */
@customElement('wa-button-css')
export default class WaButtonCss extends WebAwesomeFormAssociatedElement {
  render() {
    return 'button';
  }
}
