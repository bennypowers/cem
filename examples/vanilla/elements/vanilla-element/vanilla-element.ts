export class MessageChangedEvent extends Event {
  message: string;
  constructor(message: string) {
    super('message-changed', { bubbles: true, composed: true });
    this.message = message;
  }
}

/**
 * A vanilla custom element demonstrating Web Components without a framework.
 *
 * This component shows how to create custom elements using only native Web APIs,
 * with proper JSDoc annotations for CEM manifest generation.
 *
 * @element vanilla-element
 * @fires {MessageChangedEvent} message-changed - Dispatched when the message attribute changes
 *
 * @attr {string} message - The message to display
 * @attr {boolean} reversed - Whether to display the message in reverse
 *
 * @slot - Default slot for additional content
 *
 * @csspart message - The message paragraph element
 *
 * @cssprop --vanilla-bg - Background color of the container
 * @cssprop --vanilla-color - Text color
 * @cssprop --vanilla-padding - Padding around the content
 */
export class VanillaElement extends HTMLElement {
  private static template = document.createElement('template');

  static is = 'vanilla-element';

  static {
    this.template.innerHTML = `
      <style>
        :host {
          display: block;
        }
        div {
          background: var(--vanilla-bg, #f5f5f5);
          color: var(--vanilla-color, #333);
          padding: var(--vanilla-padding, 1rem);
          border-radius: 4px;
        }
        #message {
          margin: 0 0 0.5rem 0;
          font-size: 1.1em;
        }
      </style>
      <div>
        <p id="message" part="message"></p>
        <slot></slot>
      </div>
    `;
    customElements.define(this.is, this);
  }

  static readonly observedAttributes = ['message', 'reversed'];

  #message = 'Hello from vanilla!';

  #setMessage(value = this.message) {
    this.#message = this.reversed
      ? value.split('').reverse().join('')
      : value;
  }

  /**
   * Gets the message to display
   */
  get message(): string {
    return this.getAttribute('message') || 'Hello from vanilla!';
  }

  /**
   * Sets the message to display
   */
  set message(value: string) {
    this.setAttribute('message', value);
    this.#setMessage(value)
  }

  /**
   * Gets whether the message should be reversed
   */
  get reversed(): boolean {
    return this.hasAttribute('reversed');
  }

  /**
   * Sets whether the message should be reversed
   */
  set reversed(value: boolean) {
    this.toggleAttribute('reversed', !!value);
    this.#setMessage();
  }

  constructor() {
    super();
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' })
        .append(VanillaElement.template.content.cloneNode(true));
    }
  }

  attributeChangedCallback(name: string, oldValue: string | null, newValue: string | null) {
    if (oldValue !== newValue) {
      if (name === 'message') this.dispatchEvent(new MessageChangedEvent(newValue || ''));
      this.#setMessage(name === 'message' ? newValue ?? '' : this.message);
      if (this.isConnected) this.connectedCallback();
    }
  }

  connectedCallback() {
    const node =
      this.shadowRoot?.getElementById('message');
    if (node) node.textContent = this.#message;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'vanilla-element': VanillaElement;
  }
}
