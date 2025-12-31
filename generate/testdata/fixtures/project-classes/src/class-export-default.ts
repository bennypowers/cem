import { LitElement } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * @summary A default exported custom element
 */
@customElement('class-default-ce')
export default class ClassDefaultCe extends LitElement {
  render() {
    return 'default export with decorator';
  }
}
