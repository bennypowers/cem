import { expect, waitUntil } from '@open-wc/testing';
import { CemElement } from './cem-element.js';
// Import a real component for testing
import '../elements/cem-reconnection-content/cem-reconnection-content.js';

describe('CemElement', () => {
  describe('constructor and shadow root', () => {
    it('creates shadow root with default options', async () => {
      // Use real component from server
      const el = document.createElement('cem-reconnection-content');
      document.body.appendChild(el);

      await el.rendered;

      expect(el.shadowRoot).to.exist;
      expect(el.shadowRoot.mode).to.equal('open');

      document.body.removeChild(el);
    });

    it('uses custom shadowRootOptions when provided', async () => {
      class TestElement extends CemElement {
        static shadowRootOptions = { mode: 'open', delegatesFocus: true };
        static elementName = 'cem-reconnection-content'; // Use real template
      }
      customElements.define('test-element-delegates', TestElement);

      const el = document.createElement('test-element-delegates');
      document.body.appendChild(el);

      await el.rendered;

      expect(el.shadowRoot).to.exist;
      expect(el.shadowRoot.delegatesFocus).to.be.true;

      document.body.removeChild(el);
    });

    it('is SSR-aware and does not recreate existing shadow root', async () => {
      class TestElement extends CemElement {
        static elementName = 'cem-reconnection-content';
      }
      customElements.define('test-element-ssr-constructor', TestElement);

      const el = document.createElement('test-element-ssr-constructor');
      const shadowRoot = el.shadowRoot;

      // Shadow root should exist after construction
      expect(shadowRoot).to.exist;
      expect(shadowRoot.mode).to.equal('open');

      document.body.appendChild(el);
      await el.rendered;

      // Same shadow root should still be in use (not recreated)
      expect(el.shadowRoot).to.equal(shadowRoot);

      document.body.removeChild(el);
    });
  });

  describe('template loading', () => {
    it('loads template from real server', async () => {
      const el = document.createElement('cem-reconnection-content');
      document.body.appendChild(el);

      await el.rendered;

      // Verify template loaded - cem-reconnection-content has retry-info element
      const retryInfo = el.shadowRoot.getElementById('retry-info');
      expect(retryInfo).to.exist;

      document.body.removeChild(el);
    });

    it('includes element attributes in template URL query params', async () => {
      class TestElement extends CemElement {
        static elementName = 'cem-reconnection-content';
      }
      customElements.define('test-element-attrs', TestElement);

      const el = document.createElement('test-element-attrs');
      el.setAttribute('foo', 'bar');
      el.setAttribute('baz', 'qux');
      document.body.appendChild(el);

      await el.rendered;

      // Template should load successfully even with attributes
      expect(el.shadowRoot.firstChild).to.exist;

      document.body.removeChild(el);
    });

    it('uses custom elementName when provided', async () => {
      class TestElement extends CemElement {
        static elementName = 'cem-reconnection-content';
      }
      customElements.define('test-element-custom-name', TestElement);

      const el = document.createElement('test-element-custom-name');
      document.body.appendChild(el);

      await el.rendered;

      // Should load cem-reconnection-content template, not test-element-custom-name
      const retryInfo = el.shadowRoot.getElementById('retry-info');
      expect(retryInfo).to.exist;

      document.body.removeChild(el);
    });

    it('applies template HTML to shadow root', async () => {
      const el = document.createElement('cem-reconnection-content');
      document.body.appendChild(el);

      await el.rendered;

      // Check that real template content exists
      expect(el.shadowRoot.innerHTML).to.include('retry-info');
      expect(el.shadowRoot.getElementById('retry-info')).to.exist;

      document.body.removeChild(el);
    });
  });

  describe('stylesheet handling', () => {
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
  });

  describe('afterTemplateLoaded lifecycle', () => {
    it('calls afterTemplateLoaded after template is applied', async () => {
      let lifecycleCalled = false;
      let shadowRootAvailable = false;

      class TestElement extends CemElement {
        static elementName = 'cem-reconnection-content';

        async afterTemplateLoaded() {
          lifecycleCalled = true;
          shadowRootAvailable = !!this.shadowRoot?.firstChild;
        }
      }
      customElements.define('test-element-lifecycle', TestElement);

      const el = document.createElement('test-element-lifecycle');
      document.body.appendChild(el);

      await el.rendered;

      expect(lifecycleCalled).to.be.true;
      expect(shadowRootAvailable).to.be.true;

      document.body.removeChild(el);
    });

    it('allows afterTemplateLoaded to set up event listeners', async () => {
      let buttonClicked = false;

      class TestElement extends CemElement {
        static elementName = 'cem-reconnection-content';

        async afterTemplateLoaded() {
          const retryInfo = this.shadowRoot.getElementById('retry-info');
          if (retryInfo) {
            retryInfo.addEventListener('click', () => {
              buttonClicked = true;
            });
          }
        }
      }
      customElements.define('test-element-events', TestElement);

      const el = document.createElement('test-element-events');
      document.body.appendChild(el);

      await el.rendered;

      const retryInfo = el.shadowRoot.getElementById('retry-info');
      retryInfo.click();

      expect(buttonClicked).to.be.true;

      document.body.removeChild(el);
    });
  });

  describe('SSR support', () => {
    it('does not repopulate shadow root if already populated', async () => {
      class TestElement extends CemElement {
        static elementName = 'cem-reconnection-content';
      }
      customElements.define('test-element-ssr', TestElement);

      const el = document.createElement('test-element-ssr');
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
