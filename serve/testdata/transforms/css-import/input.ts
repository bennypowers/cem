import './component.css';

export class TestComponent extends HTMLElement {
  connectedCallback() {
    this.textContent = 'Hello';
  }
}
