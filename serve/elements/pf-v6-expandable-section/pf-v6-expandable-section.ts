import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-expandable-section.css' with { type: 'css' };

import '../pf-v6-button/pf-v6-button.js';

/**
 * Custom event fired when expandable section toggles
 */
export class PfExpandableSectionToggleEvent extends Event {
  expanded: boolean;
  constructor(expanded: boolean) {
    super('toggle', { bubbles: true });
    this.expanded = expanded;
  }
}

/**
 * PatternFly v6 Expandable Section
 *
 * A disclosure component built on native `<details>`/`<summary>` elements,
 * providing show/hide content with smooth animations.
 *
 * @slot toggle-icon - Custom toggle icon (defaults to chevron)
 * @slot toggle-text - Text for the toggle button
 * @slot - Content to show/hide
 *
 * @fires toggle - Fires when expansion state changes
 *
 * @csspart toggle - The summary/toggle container
 * @csspart content - The content container
 */
@customElement('pf-v6-expandable-section')
export class PfV6ExpandableSection extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  @property({ type: Boolean, reflect: true, attribute: 'expand-top' })
  accessor expandTop = false;

  @property({ type: Boolean, reflect: true, attribute: 'expand-bottom' })
  accessor expandBottom = false;

  @property({ type: Boolean, reflect: true, attribute: 'display-lg' })
  accessor displayLg = false;

  @property({ type: Boolean, reflect: true })
  accessor indented = false;

  @property({ type: Boolean, reflect: true, attribute: 'limit-width' })
  accessor limitWidth = false;

  @property({ type: Boolean, reflect: true })
  accessor truncate = false;

  @property({ attribute: 'toggle-text' })
  accessor toggleText = '';

  toggle() {
    this.expanded = !this.expanded;
  }

  show() {
    this.expanded = true;
  }

  hide() {
    this.expanded = false;
  }

  render() {
    return html`
      <details ?open=${this.expanded}
               @toggle=${this.#onDetailsToggle}>
        <summary part="toggle">
          <pf-v6-button id="toggle-button"
                        part="toggle-button"
                        variant="link"
                        tabindex="-1">
            <span id="toggle-icon"
                  slot="icon-start">
              <slot name="toggle-icon">
                <svg viewBox="0 0 256 512"
                     fill="currentColor"
                     aria-hidden="true"
                     role="img"
                     width="1em"
                     height="1em">
                  <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
                </svg>
              </slot>
            </span>
            <slot name="toggle-text">${this.toggleText || nothing}</slot>
          </pf-v6-button>
        </summary>
        <div id="content"
             part="content">
          <slot></slot>
        </div>
      </details>
    `;
  }

  #onDetailsToggle(event: Event) {
    const details = event.target as HTMLDetailsElement;
    this.expanded = details.open;
    this.dispatchEvent(new PfExpandableSectionToggleEvent(this.expanded));
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-expandable-section': PfV6ExpandableSection;
  }
}
