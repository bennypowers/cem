import { LitElement, customElement, css, html } from 'lit';

// a plain comment above an exported custom element
@customElement('commented-element')
export class CommentedElement extends LitElement {
  static styles = css`:host { --color: red; }`;

  render() {
    return html`<slot></slot>`;
  }
}

// a plain comment above a non-exported custom element
@customElement('commented-unexported')
class CommentedUnexported extends LitElement {
  render() {
    return html`<slot></slot>`;
  }
}

// a plain comment above a non-exported class
class CommentedPlain extends HTMLElement {
  static observedAttributes = ['foo', 'bar'];
}

// a plain comment above an exported class
export class CommentedExported extends HTMLElement {}
