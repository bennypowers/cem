import { customElement, property, query, state } from 'lit/decorators.js';

class WebAwesomeFormAssociatedElement {}

/**
 * @summary Test with multiple decorator imports
 */
@customElement('wa-button-multi')
export default class WaButtonMulti extends WebAwesomeFormAssociatedElement {
  @property() name = '';
  
  render() {
    return 'button';
  }
}
