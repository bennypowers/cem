import { LitElement, customElement } from 'lit';

const tagName = 'unexported-variable-element';

@customElement(tagName)
class UnexportedVariableElement extends LitElement {
  render() {
    return 'hello';
  }
}
