import { LitElement, customElement } from 'lit';

export const tagName = 'variable-element';

@customElement(tagName)
export class VariableElement extends LitElement {
  render() {
    return 'hello';
  }
}
