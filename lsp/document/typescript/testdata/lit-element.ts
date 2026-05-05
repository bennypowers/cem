import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('my-app')
class MyApp extends LitElement {
  @property()
  name = 'World';

  render() {
    return html`
      <my-header title="Hello ${this.name}"></my-header>
      <my-card variant="primary">
        <span slot="content">Welcome</span>
      </my-card>
    `;
  }
}
