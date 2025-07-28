import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * Component to test privacy validation rules.
 * Contains methods with underscore prefixes but no privacy annotations.
 */
@customElement('privacy-test')
export class PrivacyTest extends LitElement {
  @property({ type: String }) value = '';

  static styles = css`
    :host { display: block; }
  `;

  render() {
    return html`
      <div>
        <button @click=${this._handleClick}>Click me</button>
        <span>${this.value}</span>
      </div>
    `;
  }

  /**
   * This method starts with underscore but has no privacy flag.
   * Should trigger a validation warning.
   */
  _handleClick() {
    this.value = 'clicked';
    this._updateState();
  }

  /**
   * Another underscore method without privacy flag.
   * Should also trigger a validation warning.
   */
  _updateState() {
    this.dispatchEvent(new CustomEvent('state-updated'));
  }

  /**
   * This method is properly marked as private.
   * Should NOT trigger a validation warning.
   * @private
   */
  _privateMethod() {
    console.log('This is properly private');
  }
}