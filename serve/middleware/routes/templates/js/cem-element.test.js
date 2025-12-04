import { expect } from '@open-wc/testing';
import { CemElement } from './cem-element.js';
// Import a real component for testing
import '../elements/cem-reconnection-content/cem-reconnection-content.js';

describe('CemElement', () => {
  describe('basic functionality', () => {
    it('creates shadow root and loads template', async () => {
      const el = document.createElement('cem-reconnection-content');
      document.body.appendChild(el);

      await el.rendered;

      expect(el.shadowRoot).to.exist;
      expect(el.shadowRoot.mode).to.equal('open');

      // Verify template loaded
      const retryInfo = el.shadowRoot.getElementById('retry-info');
      expect(retryInfo).to.exist;

      document.body.removeChild(el);
    });

    it('applies stylesheet via adoptedStyleSheets', async () => {
      const el = document.createElement('cem-reconnection-content');
      document.body.appendChild(el);

      await el.rendered;

      expect(el.shadowRoot.adoptedStyleSheets).to.have.lengthOf(1);
      const sheet = el.shadowRoot.adoptedStyleSheets[0];
      expect(sheet.cssRules.length).to.be.greaterThan(0);

      document.body.removeChild(el);
    });

    it('caches stylesheets across instances', async () => {
      // Create first instance
      const el1 = document.createElement('cem-reconnection-content');
      document.body.appendChild(el1);
      await el1.rendered;

      // Create second instance
      const el2 = document.createElement('cem-reconnection-content');
      document.body.appendChild(el2);
      await el2.rendered;

      // Both should share the same stylesheet object
      expect(el1.shadowRoot.adoptedStyleSheets[0]).to.equal(el2.shadowRoot.adoptedStyleSheets[0]);

      document.body.removeChild(el1);
      document.body.removeChild(el2);
    });

    it('does not repopulate shadow root on reconnect', async () => {
      const el = document.createElement('cem-reconnection-content');
      document.body.appendChild(el);

      await el.rendered;

      // Save reference to first child
      const firstChild = el.shadowRoot.firstChild;
      expect(firstChild).to.exist;

      // Disconnect and reconnect element
      document.body.removeChild(el);
      document.body.appendChild(el);

      // Wait a bit to ensure connectedCallback has run
      await new Promise(resolve => setTimeout(resolve, 50));

      // Shadow root should not be repopulated - same firstChild
      expect(el.shadowRoot.firstChild).to.equal(firstChild);

      document.body.removeChild(el);
    });
  });
});
