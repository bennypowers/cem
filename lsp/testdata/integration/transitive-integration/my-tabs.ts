import { LitElement, customElement } from 'lit';
import './my-tab.js';

@customElement('my-tabs')
export class MyTabs extends LitElement {
  connectedCallback() {
    this.innerHTML = '<div class="tabs">Tabs container</div>';
  }
}
