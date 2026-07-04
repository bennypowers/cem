import { html } from 'lit';
import './other-element.js';

const tpl = html`
  <script type="module">
    import './inner-element.js';
  </script>
`;

export class MyComponent {
  render() {
    return html`<my-element></my-element>`;
  }
}
