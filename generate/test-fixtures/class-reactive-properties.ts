@customElement('class-reactive-properties')
class ClassReactiveProperties extends LitElement {
  @property() string: string;
  @property({ attribute: false }) protected protected: string;
  @property() initialized = 'string'
  @property({ reflect: true, type: Boolean }) reflect: boolean;
  @property({ attribute: 'attr-reflects', reflect: true, type: Boolean }) attrReflects: boolean;
  @property({ attribute: true, reflect: true }) truetrue: string;
  @property({ attribute: false }) nonAttr: number;
  @squishy
  @icky()
  @tickly({ giggly: true })
  @property({ attribute: 'multi-decorator', reflect: true }) multiDecorator = 0;
}
