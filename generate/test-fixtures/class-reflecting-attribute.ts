@customElement('class-reflecting-attr')
class ClassReflectingAttr extends LitElement {
  @property({
    type: Boolean,
    reflect: true,
    attribute: 'reflecting-attribute',
  }) reflectingAttribute = false;
}

