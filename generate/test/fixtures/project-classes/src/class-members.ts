/**
 * Tests that fields, methods, properties, attributes, and arrow methods are picked up
 * @attr jsdoc-attr
 */
@customElement('class-members')
@spoiler
@spoilerArgs('with', 'args')
export class ClassMembers extends LitElement {
  static styles = [styles];
  @property({ reflect: true, attribute: 'attr-reflect' }) attrReflect?: string;
  @property({ reflect: true }) reflect?: 'reflect';
  @property({ reflect: true, attribute: 'attr-reflect-bool', type: Boolean }) attrReflectBool = false;
  #private = 0;
  field = 0;
  arrow = () => 4;
  method() {}
  willUpdate() { }
  override render() { return '' }
}

