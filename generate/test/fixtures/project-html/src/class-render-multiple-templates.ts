import { LitElement, html } from 'lit-element';
import { customElement } from 'lit-element/decorators.js';

@customElement('class-render-multiple-templates')
class ClassRenderMultipleTemplates extends LitElement {
  render() {
    return [
      html`
        <!-- First Slot -->
        <slot name="first"></slot>
        <!-- Part One -->
        <div part="one"></div>
      `,
      html`
        <!-- Second Slot -->
        <slot name="second"></slot>
        <!-- Part Two -->
        <div part="two"></div>
      `
    ];
  }
}

@customElement('class-render-composed-templates')
class ClassRenderComposedTemplates extends LitElement {
  render() {
    const composedSlot = html`
      <!-- Composed Slot -->
      <slot name="composed"></slot>
    `;
    return html`
      ${composedSlot}
      <!-- Main Slot -->
      <slot name="main"></slot>
    `;
  }
}
