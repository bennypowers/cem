import { customElement } from 'lit/decorators.js';

class CustomBase {}

/**
 * @summary A custom element with custom base class
 */
@customElement('class-custom-base-ce')
export class ClassCustomBaseCe extends CustomBase {
  render() {
    return 'custom base';
  }
}
