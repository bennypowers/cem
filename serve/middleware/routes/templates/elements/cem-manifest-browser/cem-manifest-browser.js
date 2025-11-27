import { CemElement } from '/__cem/cem-element.js';

/**
 * Placeholder for CemManifestBrowser
 * @customElement cem-manifest-browser
 */
export class CemManifestBrowser extends CemElement {
  static is = 'cem-manifest-browser';

  static {
    customElements.define(this.is, this);
  }
}

