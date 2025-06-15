import { LitElement, html, property, customElement } from 'lit-element';

@customElement('expandable-panel')
export class ExpandablePanel extends LitElement {
  @property({ type: Boolean }) expanded = false;

  render() {
    return html`
      <button @click=${() => (this.expanded = !this.expanded)}>
        ${this.expanded ? 'Collapse' : 'Expand'}
      </button>
      <div ?hidden=${!this.expanded}>
        <slot></slot>
      </div>
    `;
  }
}
