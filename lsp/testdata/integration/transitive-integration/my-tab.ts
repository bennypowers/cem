import { LitElement, customElement } from 'lit';
import './my-icon.js';

@customElement('my-tab')
export class MyTab extends LitElement {
  connectedCallback() {
    this.innerHTML = '<div class="tab">Tab content</div>';
  }
}
