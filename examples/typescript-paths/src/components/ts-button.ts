import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A button component demonstrating TypeScript path mappings.
 *
 * This component shows how CEM works with TypeScript projects that use:
 * - Path mappings in tsconfig.json (@components/*, @utils/*, etc.)
 * - src/ directory for source files
 * - dist/ directory for compiled output
 * - Dev server URL rewrites
 *
 * @summary Button with TypeScript path mapping demonstration
 * @status stable
 *
 * @slot - Button label
 *
 * @csspart base - The button element
 *
 * @cssprop --ts-button-bg - Background color
 * @cssprop --ts-button-color - Text color
 */
@customElement('ts-button')
export class TsButton extends LitElement {
  static styles = css`
    :host {
      display: inline-block;
    }

    button {
      padding: 0.5rem 1rem;
      background: var(--ts-button-bg, #4f46e5);
      color: var(--ts-button-color, white);
      border: none;
      border-radius: 6px;
      font-family: inherit;
      font-size: 1rem;
      cursor: pointer;
      transition: all 0.2s;
    }

    button:hover {
      opacity: 0.9;
      transform: translateY(-1px);
    }

    button:active {
      transform: translateY(0);
    }

    button:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }
  `;

  /**
   * Whether the button is disabled
   */
  @property({ type: Boolean }) disabled = false;

  /**
   * Button label (alternative to slot)
   */
  @property() label = '';

  render() {
    return html`
      <button part="base" ?disabled=${this.disabled}>
        ${this.label ? this.label : html`<slot></slot>`}
      </button>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'ts-button': TsButton;
  }
}
