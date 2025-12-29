class MyElement extends HTMLElement {
  connectedCallback() {
    this.innerHTML = '<p>Hello from my-element!</p>';
  }
}

customElements.define('my-element', MyElement);
