import { LitElement, html, property, customElement } from 'lit-element';

interface User { id: number; name: string; }

@customElement('user-list')
export class UserList extends LitElement {
  @property({ type: Array }) users: User[] = [];

  render() {
    return html`
      <ul>
        ${this.users.map(user => html`<li>${user.name}</li>`)}
      </ul>
    `;
  }
}
