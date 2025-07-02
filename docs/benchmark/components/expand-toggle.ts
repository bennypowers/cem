import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element expand-toggle
 * @summary Expand/collapse details panel.
 */
@customElement('expand-toggle')
export class ExpandToggle extends LitElement {
  @property({ type: Boolean, reflect: true }) expanded = false;

  static styles = css`
    .panel { overflow: hidden; transition: max-height 0.2s; }
    :host([expanded]) .panel { max-height: 200px; }
    .panel { max-height: 0; }
    button { margin: 0.5em 0; }
  `;

  toggle() {
    this.expanded = !this.expanded;
  }

  render() {
    return html`
      <button @click=${this.toggle}>
        ${this.expanded ? 'Collapse' : 'Expand'}
      </button>
      <div class="panel">
        <slot></slot>
      </div>
    `;
  }
}