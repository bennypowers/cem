import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-button.css' with { type: 'css' };

type ButtonVariant = 'primary' | 'secondary' | 'tertiary' | 'danger' | 'warning' | 'plain' | 'link';
type ButtonSize = 'sm' | 'md' | 'lg';

/**
 * PatternFly v6 inspired button component
 *
 * Renders as a host-based button (role="button") by default,
 * or as a shadow `<a>` when `href` is set.
 *
 * @slot - Default slot for button text/content
 * @slot icon-start - Slot for icon before text
 * @slot icon-end - Slot for icon after text
 *
 * @fires click - Bubbles click events from internal button
 */
@customElement('pf-v6-button')
export class PfV6Button extends LitElement {
  static readonly formAssociated = true;

  static shadowRootOptions: ShadowRootInit = {
    ...LitElement.shadowRootOptions,
    delegatesFocus: true,
  };

  static styles = styles;

  #internals = this.attachInternals();

  @property({ reflect: true })
  accessor variant?: ButtonVariant;

  @property({ reflect: true })
  accessor size?: ButtonSize;

  @property({ type: Boolean, reflect: true })
  accessor block = false;

  @property({ type: Boolean, reflect: true })
  accessor disabled = false;

  @property({ type: Boolean, reflect: true, attribute: 'no-padding' })
  accessor noPadding = false;

  @property({ reflect: true })
  accessor href?: string;

  @property()
  accessor type: string = 'button';

  connectedCallback() {
    super.connectedCallback();
    this.#updateMode();
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('href')) {
      this.#updateMode();
    }
    if (changed.has('disabled')) {
      this.#updateDisabled();
    }
  }

  render() {
    const slots = html`
      <slot name="icon-start"></slot>
      <slot></slot>
      <slot name="icon-end"></slot>
    `;

    if (this.href) {
      return html`<a id="link"
                      href=${this.href}
                      @click=${this.#onLinkClick}>${slots}</a>`;
    }

    return slots;
  }

  #updateMode() {
    if (this.href) {
      // Link mode: remove button semantics from host
      this.#internals.role = null;
      this.removeEventListener('click', this.#onClick);
      this.removeEventListener('keydown', this.#onKeydown);
      if (this.getAttribute('tabindex') === '0') {
        this.removeAttribute('tabindex');
      }
    } else {
      // Button mode: set role and make focusable
      this.#internals.role = 'button';
      if (!this.hasAttribute('tabindex')) {
        this.setAttribute('tabindex', '0');
      }
      this.addEventListener('click', this.#onClick);
      this.addEventListener('keydown', this.#onKeydown);
    }
    this.#updateDisabled();
  }

  #updateDisabled() {
    if (this.href) {
      // For links, set aria-disabled on the shadow <a>
      const link = this.shadowRoot?.getElementById('link');
      if (link) {
        if (this.disabled) {
          link.setAttribute('aria-disabled', 'true');
          link.style.pointerEvents = 'none';
        } else {
          link.removeAttribute('aria-disabled');
          link.style.pointerEvents = '';
        }
      }
    } else {
      // For host-based buttons, use ElementInternals
      if (this.disabled) {
        this.#internals.ariaDisabled = 'true';
        this.setAttribute('tabindex', '-1');
      } else {
        this.#internals.ariaDisabled = null;
        if (!this.href) {
          this.setAttribute('tabindex', '0');
        }
      }
    }
  }

  #onClick = (event: Event) => {
    if (this.disabled) {
      event.preventDefault();
      event.stopImmediatePropagation();
    }
  };

  #onKeydown = (event: KeyboardEvent) => {
    if (this.disabled) {
      event.preventDefault();
      return;
    }
    if (event.key === ' ' || event.key === 'Enter') {
      event.preventDefault();
      this.click();
    }
  };

  #onLinkClick(event: Event) {
    if (this.disabled) {
      event.preventDefault();
      event.stopPropagation();
    }
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('click', this.#onClick);
    this.removeEventListener('keydown', this.#onKeydown);
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-button': PfV6Button;
  }
}
