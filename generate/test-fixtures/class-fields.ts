@customElement('class-fields')
class ClassFields extends LitElement {
  static readonly styles = ["MUST_OMIT"] // must omit
  static formAssociated = "MUST_OMIT";   // must omit
  static #hehe = "MUST_OMIT";            // must omit
  #hehe = "MUST_OMIT"                    // must omit

  // INCLUDE
  static readonly rstat = 0;

  // INCLUDE
  public static readonly rstatpub = 0;

  // INCLUDE
  private static readonly rstatpriv: Float = 0.1;

  // INCLUDE
  protected static readonly rstatprot = 0;

  // INCLUDE
  static stat = 0;

  // INCLUDE
  public static statpriv = 0;

  // INCLUDE
  private static statpriv = 0;

  // INCLUDE
  protected static statprot = 0;

  // INCLUDE
  @property() string: string;

  // INCLUDE
  /** string description field */
  @property() stringDescription: string;

  // INCLUDE
  /**
   * multiline
   * description
   * field
   */
  @property() stringMultilineDescription: string;

  // INCLUDE
  /** @summary summary field */
  @property() summary: string;

  // INCLUDE
  /**
   * @summary multiline
   *          summary
   *          field
   */
  @property() multilineSummary: string;

  // INCLUDE
  /**
   * camelcase
   * multiline
   * description
   * field
   * @summary camelcase field
   */
  @property({ attribute: 'camel-case' }) camelCase: string;

  // INCLUDE
  @property() initialized = 'string'

  // INCLUDE
  @property({ reflects: true, type: Boolean }) reflects: boolean;

  // INCLUDE
  @property({ attribute: 'attr-reflects', reflects: true, type: Boolean }) attrReflects: boolean;

  // INCLUDE
  @property({ attribute: true, reflects: true }) truetrue: string;

  // INCLUDE
  @property({ attribute: false }) nonAttr: number;

  // INCLUDE
  /** highly decorated field */
  @squishy
  @icky()
  @tickly({ giggly: true })
  @property({ attribute: 'multi-decorator', reflects: true }) multiDecorator = 0;

  // INCLUDE
  /** @type {number} */
  @property() typejsdoc: string;

  // INCLUDE
  /**
   * @deprecated
   */
  @property() deprecated: string;

  // INCLUDE
  /**
   * @deprecated reason
   */
  @property() reason: string;

  // INCLUDE
  @property({ attribute: false }) protected protected: string;
}
