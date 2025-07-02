import { LitElement, html, property, customElement } from 'lit-element';

interface User { name: string; age: number; }

@customElement('complex-types')
export class ComplexTypes extends LitElement {
  @property({ type: Array }) users: User[] = [];
  @property({ type: Object }) config: Record<string, unknown> = {};

  render() {
    return html`${this.users.map(u => html`<li>${u.name} (${u.age})</li>`)}`;
  }
}
