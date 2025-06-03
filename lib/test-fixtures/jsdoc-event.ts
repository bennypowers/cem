import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * @event event
 * @event event-description - event description
 * @event event-description-multiline - multiline
 *                                      event description
 * @event {Event} event-typed
 * @event {Event} event-typed-description - event typed description
 * @event {Event} event-typed-description-multiline - multiline
 *                                                    event typed description
 * @fires fires
 * @fires fires-description - fires description
 * @fires fires-description-multiline - multiline
 *                                      fires description
 * @fires {Event} fires-typed
 * @fires {Event} fires-typed-description - fires typed description
 * @fires {Event} fires-typed-description-multiline - multiline
 *                                                    fires typed description
 */
@customElement('jsdoc-events')
export class JsdocEvents extends LitElement {
}
