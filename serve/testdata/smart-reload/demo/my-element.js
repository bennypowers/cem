export class MyElement extends HTMLElement {
  connectedCallback() {
    this.textContent = 'Hello World';
  }
}

customElements.define('my-element', MyElement);
