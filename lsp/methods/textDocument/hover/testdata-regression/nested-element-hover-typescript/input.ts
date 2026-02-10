import { html, LitElement } from 'lit';

export class MyAlert extends LitElement {
  render() {
    return html`
      <outer-surface color-palette="lightest">
        <div id="left-column">
          <inner-icon id="icon" set="ui" icon="check"></inner-icon>
        </div>
      </outer-surface>
    `;
  }
}
