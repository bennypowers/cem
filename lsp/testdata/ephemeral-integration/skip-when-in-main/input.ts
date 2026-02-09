import { LitElement, customElement, property } from 'lit';

@customElement('test-button')
export class TestButton extends LitElement {
  @property()
  variant: string;
}
