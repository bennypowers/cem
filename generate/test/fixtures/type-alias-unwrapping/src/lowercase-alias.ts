type buttonType = 'submit' | 'reset' | 'button';
type _privateType = 'internal' | 'external';
type $dollarType = 'usd' | 'eur';

export class MyElement extends HTMLElement {
  type: buttonType = 'submit';
  visibility: _privateType = 'external';
  currency: $dollarType = 'usd';
}

customElements.define('my-element', MyElement);
