import { customElement } from 'lit/decorators.js';

class CustomBase {}

/**
 * @summary Export default with custom base
 */
@customElement('class-export-default-custom-base-ce')
export default class ClassExportDefaultCustomBaseCe extends CustomBase {
  render() {
    return 'export default + custom base';
  }
}
