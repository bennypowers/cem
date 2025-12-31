import { customElement } from 'lit/decorators.js';

class WebAwesomeFormAssociatedElement {}

/**
 * @summary Buttons represent actions that are available to the user.
 */
@customElement('wa-button')
export default class WaButton extends WebAwesomeFormAssociatedElement {
  render() {
    return 'button';
  }
}
