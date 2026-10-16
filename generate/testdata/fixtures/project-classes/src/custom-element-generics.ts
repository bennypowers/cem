import { LitElement } from 'lit';
import { customElement } from 'lit/decorators.js';

/**
 * A custom element whose data shape is generic.
 */
@customElement('data-view')
export class DataView<Data = unknown> extends LitElement {
  data?: Data;

  render() {
    return null;
  }
}
