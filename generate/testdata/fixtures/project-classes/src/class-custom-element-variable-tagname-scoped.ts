import { LitElement, customElement } from 'lit';

function getTagName() {
  const tagName = 'scoped-wrong-element';
  return tagName;
}

const tagName = 'scoped-right-element';

@customElement(tagName)
export class ScopedElement extends LitElement {
}
