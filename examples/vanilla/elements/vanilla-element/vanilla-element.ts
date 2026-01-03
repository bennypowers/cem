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
 * @csspart container - The container div element
 * @csspart message - The message paragraph element
 *
 * @cssprop --vanilla-bg - Background color of the container
 * @cssprop --vanilla-color - Text color
 * @cssprop --vanilla-padding - Padding around the content
 */
export class VanillaElement extends HTMLElement {
  private static template = document.createElement('template');
  static get observedAttributes() {
    return ['message', 'reversed'];
  }

  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
  }

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback(name: string, oldValue: string | null, newValue: string | null) {
    if (oldValue !== newValue) {
      if (name === 'message') {
        this.dispatchEvent(new MessageChangedEvent(newValue || ''));
      }
      this.render();
    }
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
    if (value) {
      this.setAttribute('reversed', '');
    } else {
      this.removeAttribute('reversed');
    }
  }

  render() {
    const message = this.reversed
      ? this.message.split('').reverse().join('')
      : this.message;

    if (!this.shadowRoot) return;

    VanillaElement.template.innerHTML = `
      <style>
        :host {
          display: block;
        }
        [part="container"] {
          background: var(--vanilla-bg, #f5f5f5);
          color: var(--vanilla-color, #333);
          padding: var(--vanilla-padding, 1rem);
          border-radius: 4px;
        }
        [part="message"] {
          margin: 0 0 0.5rem 0;
          font-size: 1.1em;
        }
      </style>
      <div part="container">
        <p part="message">${message}</p>
        <slot></slot>
      </div>
    `;

    this.shadowRoot.innerHTML = '';
    this.shadowRoot.appendChild(VanillaElement.template.content.cloneNode(true));
  }
}

customElements.define('vanilla-element', VanillaElement);

declare global {
  interface HTMLElementTagNameMap {
    'vanilla-element': VanillaElement;
  }
}
