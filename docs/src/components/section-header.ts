import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * @element section-header
 * @summary Header section with a slot for actions.
 */
@customElement('section-header')
export class SectionHeader extends LitElement {
  @property({ type: String }) heading = '';

  static styles = css`
    .header { display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #ccc; padding: 0.25em 0; }
    .title { font-size: 1.2em; font-weight: bold; }
  `;

  render() {
    return html`
      <div class="header">
        <span class="title">${this.heading}</span>
        <slot name="actions"></slot>
      </div>
    `;
  }
}