import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * description
 * multiline description
 * @summary summary
 *          multiline summary
 * @attr attr
 * @attr attr-description - attr description
 * @attr attr-description-multiline - multiline
 *                                    attr description
 * @attr {number} attr-typed
 * @attr {number} attr-typed-description - attr typed description
 * @attr {number} attr-typed-description-multiline - multiline
 *                                                   attr typed description
 * @attribute attribute
 * @attribute attribute-description - attribute description
 * @attribute attribute-description-multiline - multiline
 *                                              attribute description
 * @attribute {number} attribute-typed
 * @attribute {number} attribute-typed-description - attribute typed description
 * @attribute {number} attribute-typed-description-multiline - multiline
 *                                                             attribute typed description
 * @csspart part
 * @csspart part-description - part description
 * @csspart part-description-multiline - multiline
 *                                       part with description
 * @cssprop --prop
 * @cssprop --prop-description - prop description
 * @cssprop --prop-description-multiline - multiline
 *                                         prop description
 * @cssprop [--prop-default=none]
 * @cssprop [--prop-default-description=none] - prop default description
 * @cssprop [--prop-default-description-multiline=none] - multiline
 *                                              prop default description
 * @cssprop {<color>} --prop-typed
 * @cssprop {<color>} --prop-typed-description - prop typed description
 * @cssprop {<color>} --prop-typed-description-multiline - multiline
 *                                                         prop typed description
 * @cssprop {<color>} [--prop-typed-default=none]
 * @cssprop {<color>} [--prop-typed-default-description=none] - prop typed default description
 * @cssprop {<color>} [--prop-typed-default-description-multiline=none] - multiline
 *                                                                        prop typed default description
 * @cssproperty --property
 * @cssproperty --property-description - property description
 * @cssproperty --property-description-multiline - multiline
 *                                                 property description
 * @cssproperty [--property-default=none]
 * @cssproperty [--property-default-description=none] - property default description
 * @cssproperty {<color>} [--property-typed-default-description-multiline=none] - multiline
 *                                                                                property typed default description
 * @cssproperty {<color>} --property-typed
 * @cssproperty {<color>} --property-typed-description - property typed description
 * @cssproperty {<color>} --property-typed-description-multiline - multiline
 *                                                                 property typed description
 * @cssproperty {<color>} [--property-typed-default=none]
 * @cssproperty {<color>} [--property-typed-default-description=none] - property typed default description
 * @cssproperty {<color>} [--property-typed-default-description-multiline=none] - multiline
 *                                                                                property typed default description
 * @cssstate state
 * @cssstate state-description - state description
 * @cssstate state-description-multiline - multiline
 *                                         state description
 * @event event
 * @event event-description - event description
 * @event event-description-multiline - multiline
 *                                      event description
 * @event {Event} event-typed
 * @event {Event} event-typed-description - event typed description
 * @event {Event} event-typed-description-multiline - multiline
 *                                                    event typed description
 * @deprecated deprecation reason
 * @fires fires
 * @fires fires-description - fires description
 * @fires fires-description-multiline - multiline
 *                                      fires description
 * @fires {Event} fires-typed
 * @fires {Event} fires-typed-description - fires typed description
 * @fires {Event} fires-typed-description-multiline - multiline
 *                                                    fires typed description
 * @slot - anon description
 * @slot slot
 * @slot slot-description - slot description
 */
@customElement('c-e')
export class CE extends LitElement {
  @property() string: string;

  /** string description */
  @property() stringDescription: string;

  /**
   * string property with init
   * @summary string
   */
  @property() stringInit = 'string';

  @property({ reflects: true, type: Boolean }) reflects: boolean;

  @property({ attribute: true, reflects: true, type: Boolean }) attrReflects: boolean;

  @property({ attribute: false }) nonAttr: string;

  /** @type {number} */
  @property() typejsdoc: string;

  /**
   * @deprecated
   */
  @property() deprecated: string;

  /**
   * @deprecated reason
   */
  @property() reason: string;
}
