@customElement('class-reactive-properties')
class ClassReactiveProperties extends LitElement {
  /** string description field */
  @property() stringDescription: string;

  /**
   * multiline
   * description
   * field
   */
  @property() stringMultilineDescription: string;

  /** @summary summary field */
  @property() summary: string;

  /**
   * @summary multiline
   *          summary
   *          field
   */
  @property() multilineSummary: string;

  /**
   * camelcase
   * multiline
   * description
   * field
   * @summary camelcase field
   */
  @property({ attribute: 'camel-case' }) camelCase: string;

  /** highly decorated field */
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
}
