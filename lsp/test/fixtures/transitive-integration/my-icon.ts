import { LitElement, customElement } from 'lit';

@customElement('my-icon')
export class MyIcon extends LitElement {
  connectedCallback() {
    this.innerHTML = '<svg><!-- icon --></svg>';
  }
}
