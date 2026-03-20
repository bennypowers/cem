import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './pf-v6-drawer.css' with { type: 'css' };

import '../pf-v6-button/pf-v6-button.js';

/**
 * Custom event fired when drawer expands
 */
export class PfDrawerExpandEvent extends Event {
  constructor() {
    super('expand', { bubbles: true });
  }
}

/**
 * Custom event fired when drawer collapses
 */
export class PfDrawerCollapseEvent extends Event {
  constructor() {
    super('collapse', { bubbles: true });
  }
}

/**
 * Custom event fired when drawer close button is clicked
 */
export class PfDrawerCloseEvent extends Event {
  constructor() {
    super('close', { bubbles: true });
  }
}

/**
 * PatternFly v6 Drawer Component
 *
 * A drawer is a sliding panel that enters from the edge of the screen to display
 * additional content without navigating away from the current page.
 *
 * @slot - Main content area
 * @slot panel-header - Panel header content
 * @slot panel-description - Panel description
 * @slot panel-body - Panel body content
 *
 * @fires {PfDrawerExpandEvent} expand - When panel expands
 * @fires {PfDrawerCollapseEvent} collapse - When panel collapses
 * @fires {PfDrawerCloseEvent} close - When close button is clicked
 */
@customElement('pf-v6-drawer')
export class PfV6Drawer extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor expanded = false;

  @property({ reflect: true, attribute: 'panel-position' })
  accessor panelPosition: 'end' | 'start' | 'bottom' = 'end';

  @property({ reflect: true })
  accessor variant: 'default' | 'inline' | 'static' = 'default';

  @property({ reflect: true, attribute: 'panel-width' })
  accessor panelWidth?: string;

  @property({ reflect: true, attribute: 'panel-width-on-lg' })
  accessor panelWidthOnLg?: string;

  @property({ reflect: true, attribute: 'panel-width-on-xl' })
  accessor panelWidthOnXl?: string;

  @property({ reflect: true, attribute: 'panel-width-on-2xl' })
  accessor panelWidthOn2xl?: string;

  @property({ type: Boolean, reflect: true, attribute: 'no-border' })
  accessor noBorder = false;

  @property({ type: Boolean, reflect: true })
  accessor resizable = false;

  #hasHeaderContent = false;
  #hasDescriptionContent = false;
  #hasBodyContent = false;

  /** Opens the drawer panel */
  open() {
    this.expanded = true;
  }

  /** Closes the drawer panel and fires close event */
  close() {
    this.expanded = false;
    this.dispatchEvent(new PfDrawerCloseEvent());
  }

  /** Toggles the drawer panel */
  toggle() {
    this.expanded = !this.expanded;
  }

  protected override updated(changed: Map<string, unknown>) {
    if (changed.has('expanded')) {
      if (this.expanded) {
        this.dispatchEvent(new PfDrawerExpandEvent());
      } else if (changed.get('expanded') !== undefined) {
        this.dispatchEvent(new PfDrawerCollapseEvent());
      }
    }
  }

  #onHeaderSlotchange(e: Event) {
    const slot = e.target as HTMLSlotElement;
    this.#hasHeaderContent = slot.assignedNodes().length > 0;
    this.requestUpdate();
  }

  #onDescriptionSlotchange(e: Event) {
    const slot = e.target as HTMLSlotElement;
    this.#hasDescriptionContent = slot.assignedNodes().length > 0;
    this.requestUpdate();
  }

  #onBodySlotchange(e: Event) {
    const slot = e.target as HTMLSlotElement;
    this.#hasBodyContent = slot.assignedNodes().length > 0;
    this.requestUpdate();
  }

  render() {
    return html`
      <div id="main">
        <div id="content">
          <div id="content-body">
            <slot></slot>
          </div>
        </div>
        <div id="panel">
          <div id="panel-head"
               ?hidden=${!this.#hasHeaderContent}>
            <div id="panel-header-text">
              <slot name="panel-header"
                    @slotchange=${this.#onHeaderSlotchange}></slot>
            </div>
            <div id="actions">
              <div id="close-wrapper">
                <pf-v6-button id="close"
                              variant="plain"
                              aria-label="Close drawer panel"
                              @click=${this.close}>
                  <svg viewBox="0 0 352 512"
                       fill="currentColor"
                       role="presentation"
                       width="1em"
                       height="1em">
                    <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
                  </svg>
                </pf-v6-button>
              </div>
            </div>
          </div>
          <div id="panel-description"
               part="panel-description"
               ?hidden=${!this.#hasDescriptionContent}>
            <slot name="panel-description"
                  @slotchange=${this.#onDescriptionSlotchange}></slot>
          </div>
          <div id="panel-body"
               part="panel-body"
               ?hidden=${!this.#hasBodyContent}>
            <slot name="panel-body"
                  @slotchange=${this.#onBodySlotchange}></slot>
          </div>
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-drawer': PfV6Drawer;
  }
}
