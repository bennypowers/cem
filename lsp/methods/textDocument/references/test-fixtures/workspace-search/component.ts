import { html, LitElement } from 'lit';

export class MyComponent extends LitElement {
  render() {
    return html`
      <div>
        <rh-card variant="primary">
          <h2>TypeScript Card</h2>
        </rh-card>
      </div>
    `;
  }

  anotherTemplate() {
    return html`<rh-button>TS Button</rh-button>`;
  }
}
