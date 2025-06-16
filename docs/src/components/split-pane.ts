import { LitElement, html, property, customElement } from 'lit-element';

@customElement('split-pane')
export class SplitPane extends LitElement {
  @property({ type: Number }) ratio = 0.5;

  render() {
    return html`
      <div style="display: flex;">
        <div style="flex: ${this.ratio};"><slot name="left"></slot></div>
        <div style="flex: ${1 - this.ratio};"><slot name="right"></slot></div>
      </div>
    `;
  }
}
