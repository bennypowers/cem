import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
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
 */
@customElement('jsdoc-attr')
export class JsdocAttr extends LitElement {
  @property() classAttr: string;
}
