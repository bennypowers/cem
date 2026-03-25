import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { ifDefined } from 'lit/directives/if-defined.js';

import styles from './pf-v6-label.css' with { type: 'css' };

type LabelColor = 'blue' | 'teal' | 'green' | 'orange' | 'purple' | 'red' | 'orangered' | 'grey' | 'yellow';
type LabelVariant = 'filled' | 'outline' | 'overflow' | 'add';
type LabelStatus = 'success' | 'warning' | 'danger' | 'info' | 'custom';

/**
 * PatternFly v6 Label
 *
 * @attr {string} color - Label color (blue, teal, green, orange, purple, red, orangered, grey, yellow)
 * @attr {'outline'|'filled'|'overflow'|'add'} variant - Label variant - default: filled
 * @attr {string} status - Status with icon (success, warning, danger, info, custom) - overrides color
 * @attr {boolean} compact - Compact size
 * @attr {boolean} disabled - Disabled state
 * @attr {boolean} editable - Editable label
 * @attr {string} href - Makes label a link
 *
 * @slot - Default slot for label text
 * @slot icon - Slot for icon (overrides status icon)
 * @slot actions - Slot for action buttons
 */
@customElement('pf-v6-label')
export class PfV6Label extends LitElement {
  static styles = styles;

  @property({ reflect: true })
  accessor color?: LabelColor;

  @property({ reflect: true })
  accessor variant?: LabelVariant;

  @property({ reflect: true })
  accessor status?: LabelStatus;

  @property({ type: Boolean, reflect: true })
  accessor compact = false;

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ type: Boolean, reflect: true })
  accessor editable = false;

  @property({ reflect: true })
  accessor href?: string;

  #hasActions = false;

  render() {
    const inner = html`
      <slot name="icon"></slot>
      <span id="text" part="text">
        <slot></slot>
      </span>
      <span id="actions"
            part="actions"
            ?hidden=${!this.#hasActions}>
        <slot name="actions"
              @slotchange=${this.#onActionsSlotChange}></slot>
      </span>
    `;

    if (this.href) {
      return html`
        <a id="content"
           part="content"
           href=${ifDefined(this.disabled ? undefined : this.href)}
           @click=${this.#onContentClick}>${inner}</a>
      `;
    }

    return html`
      <span id="content"
            part="content"
            @click=${this.#onContentClick}>${inner}</span>
    `;
  }

  #onContentClick(event: Event) {
    if (this.disabled) {
      event.preventDefault();
      event.stopPropagation();
    }
  }

  #onActionsSlotChange(event: Event) {
    const slot = event.target as HTMLSlotElement;
    const hadActions = this.#hasActions;
    this.#hasActions = slot.assignedElements().length > 0;
    if (hadActions !== this.#hasActions) {
      this.requestUpdate();
    }
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-label': PfV6Label;
  }
}
