import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('class-fields')
export class ClassFields extends LitElement {
  @property() string: string;

  /** string description */
  @property() stringDescription: string;

  /**
   * multiline
   * description
   */
  @property() stringDescription: string;

  /** @summary string */
  @property() summary: string;

  /**
   * @summary multiline
   *          summary
   */
  @property() multilineSummary: string;

  /**
   * camelcase attr
   * has multiline description
   * @summary camelcase
   */
  @property({ attribute: 'camel-case' }) camelCase: string;

  @property() initialized = 'string'

  @property({ reflects: true, type: Boolean }) reflects: boolean;

  @property({ attribute: 'attr-reflects', reflects: true, type: Boolean }) attrReflects: boolean;

  @property({ attribute: true, reflects: true }) truetrue: string;

  @property({ attribute: false }) nonAttr: string;

  /** highly decorated */
  @squishy
  @icky()
  @tickly({ giggly: true })
  @property({ attribute: 'multi-decorator', reflects: true }) multiDecorator = 0;

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

  @property({ attribute: false }) protected protected: string;
}
