import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * @csspart part
 * @csspart part-description - part description
 * @csspart part-description-multiline - multiline
 *                                       part with description
 */
@customElement('jsdoc-css-part')
export class JsdocCssPart extends LitElement {
}
