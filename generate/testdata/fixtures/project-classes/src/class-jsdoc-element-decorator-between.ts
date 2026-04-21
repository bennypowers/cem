import { html, LitElement, type TemplateResult } from 'lit';

function localized() {
  return function(target: any) { return target; };
}

/**
 * A component to display a breadcrumb trail.
 *
 * @customElement sl-breadcrumbs
 */
@localized()
export class Breadcrumbs extends LitElement {
  render(): TemplateResult {
    return html`<p>breadcrumbs</p>`;
  }
}
