export class ClassFields {
  string: string;
  protected protected: string;
  private private: string;
  initialized = 'string';
  bool = false;
  num = 0;
  // not a ce, should not produce attribute
  @squishy
  @icky()
  @tickly({ giggly: true })
  @property({ attribute: 'multi-decorator', reflect: true }) decorated = 0;
  fieldMethod = () => {}
}
