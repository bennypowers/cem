import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * Test alert component for completion tests
 */
@customElement('test-alert')
export class TestAlert extends LitElement {
  /** The state of the alert */
  @property({ type: String })
  state: 'info' | 'success' | 'warning' | 'error' = 'info';

  /** The variant of the alert */
  @property({ type: String })
  variant: 'primary' | 'secondary' | 'ghost' = 'primary';

  /** Whether the alert is disabled */
  @property({ type: Boolean })
  disabled = false;

  render() {
    return html`
      <div class="alert alert-${this.state} alert-${this.variant}" ?disabled=${this.disabled}>
        <slot name="header"></slot>
        <slot></slot>
      </div>
    `;
  }
}