@customElement('class-reactive-properties')
class ClassReactiveProperties extends LitElement {
  @property() string: string;
  @property({ attribute: false }) protected protected: string;
  @property() initialized = 'string'
  @property({ reflects: true, type: Boolean }) reflects: boolean;
  @property({ attribute: 'attr-reflects', reflects: true, type: Boolean }) attrReflects: boolean;
  @property({ attribute: true, reflects: true }) truetrue: string;
  @property({ attribute: false }) nonAttr: number;
  @squishy
  @icky()
  @tickly({ giggly: true })
  @property({ attribute: 'multi-decorator', reflects: true }) multiDecorator = 0;
}
