import { css, html, LitElement } from 'lit';
import {
	customElement, // imported from lit/decorators
	property,
	query,
} from 'lit/decorators.js';
import './other-element.js';

export class MyComponent {
  render() {
    return html`<my-element></my-element>`;
  }
}
