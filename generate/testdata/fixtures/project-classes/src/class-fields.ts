export class ClassFields {
  string: string;
  protected protected: string;
  private private: string;
  initialized = 'string';
  bool = false;
  num = 0;
  /**
   * The class is not a custom element, so even though this field has a @property decorator,
   * this member entry mustn't have an attribute
   */
  @squishy
  @icky()
  @tickly({ giggly: true })
  @property({ attribute: 'multi-decorator', reflect: true }) decorated = 0;
  /** must be classified as a method */
  fieldMethod = () => {}
}
// test change
// another change
// test change
