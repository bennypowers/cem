import { html, LitElement, type TemplateResult } from 'lit';

/**
 * Foo element
 * @element my-foo
 */
export class Foo extends LitElement {
  render(): TemplateResult {
    return html`<p>Foo</p>`;
  }
}
