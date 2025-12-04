// TypeScript custom element with modern features
import { LitElement, html, css } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';

interface GreetingOptions {
  prefix?: string;
  suffix?: string;
}

@customElement('my-greeting')
export class MyGreeting extends LitElement {
  static styles = css`
    :host {
      display: block;
      padding: 16px;
    }
  `;

  @property({ type: String })
  name: string = 'World';

  @property({ type: Number })
  count: number = 0;

  @state()
  private _greeting: string = 'Hello';

  // ECMAScript private field
  #internalState: Map<string, unknown> = new Map();

  // Private method
  #formatGreeting(options: GreetingOptions = {}): string {
    const { prefix = '', suffix = '!' } = options;
    return `${prefix}${this._greeting}, ${this.name}${suffix}`;
  }

  connectedCallback(): void {
    super.connectedCallback();
    this.#internalState.set('connected', true);
  }

  render() {
    return html`
      <div>
        <p>${this.#formatGreeting()}</p>
        <button @click=${this._handleClick}>
          Clicked ${this.count} times
        </button>
      </div>
    `;
  }

  private _handleClick(): void {
    this.count++;
    this.dispatchEvent(new CustomEvent('count-changed', {
      detail: { count: this.count },
      bubbles: true,
      composed: true
    }));
  }
}
