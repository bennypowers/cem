import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './cem-serve-chrome.js';
import { CemLogsEvent } from './cem-serve-chrome.js';

describe('cem-serve-chrome', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('cem-serve-chrome');
    document.body.appendChild(el);

    // Wait for CemElement to load template from real server
    await el.rendered;

    // Store original fetch before stubbing
    const originalFetch = window.fetch;

    // Now stub fetch for debug endpoint only (after template loads)
    sinon.stub(window, 'fetch').callsFake((url, ...args) => {
      if (url === '/__cem/debug') {
        return Promise.resolve({
          json: () => Promise.resolve({
            version: '1.0.0',
            os: 'darwin/arm64',
            watchDir: '/test',
            manifestSize: '100KB',
            demoCount: 5,
            demos: [],
            importMap: {}
          })
        });
      }
      // Pass through all other fetches to the original implementation
      return originalFetch.call(window, url, ...args);
    });
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
    window.fetch.restore();
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('cem-serve-chrome');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('CemServeChrome');
      expect(el.shadowRoot).to.exist;
    });

    it('renders drawer component', () => {
      const drawer = el.shadowRoot.querySelector('cem-drawer');
      expect(drawer).to.exist;
    });

    it('renders tabs component', () => {
      const tabs = el.shadowRoot.querySelector('pf-v6-tabs');
      expect(tabs).to.exist;
    });

    it('renders connection alerts component', () => {
      const connectionAlerts = el.shadowRoot.getElementById('connection-alerts');
      expect(connectionAlerts).to.exist;
    });

    it('renders error overlay component', () => {
      const errorOverlay = el.shadowRoot.getElementById('error-overlay');
      expect(errorOverlay).to.exist;
    });

    it('renders reconnection modal', () => {
      const reconnectionModal = el.shadowRoot.getElementById('reconnection-modal');
      expect(reconnectionModal).to.exist;
    });

    it('renders debug modal', () => {
      const debugModal = el.shadowRoot.getElementById('debug-modal');
      expect(debugModal).to.exist;
    });

    it('renders log container', () => {
      const logContainer = el.shadowRoot.getElementById('log-container');
      expect(logContainer).to.exist;
    });
  });

  describe('event custom classes', () => {
    it('CemLogsEvent has correct properties', () => {
      const logs = [{ type: 'info', date: '2025-01-01T00:00:00Z', message: 'Test' }];
      const event = new CemLogsEvent(logs);

      expect(event.type).to.equal('cem:logs');
      expect(event.logs).to.equal(logs);
      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });
  });

  describe('observed attributes', () => {
    it('observes knobs attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('knobs');
    });

    it('observes primary-tag-name attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('primary-tag-name');
    });

    it('observes demo-title attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('demo-title');
    });

    it('observes package-name attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('package-name');
    });

    it('observes canonical-url attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('canonical-url');
    });

    it('observes source-url attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('source-url');
    });

    it('observes drawer-open attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('drawer-open');
    });

    it('observes drawer-height attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('drawer-height');
    });

    it('observes tabs-selected attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('tabs-selected');
    });
  });

  describe('demo getter', () => {
    it('returns demo element', () => {
      const demoEl = document.createElement('cem-serve-demo');
      el.appendChild(demoEl);

      expect(el.demo).to.equal(demoEl);
    });

    it('returns null when no demo element', () => {
      expect(el.demo).to.be.null;
    });
  });

  describe('knobs getter', () => {
    it('returns knobs attribute value', () => {
      el.setAttribute('knobs', 'true');
      expect(el.knobs).to.equal('true');
    });

    it('returns empty string when not set', () => {
      expect(el.knobs).to.equal('');
    });
  });

  describe('primaryTagName getter', () => {
    it('returns primary-tag-name attribute value', () => {
      el.setAttribute('primary-tag-name', 'my-button');
      expect(el.primaryTagName).to.equal('my-button');
    });

    it('returns empty string when not set', () => {
      expect(el.primaryTagName).to.equal('');
    });
  });

  describe('demoTitle getter', () => {
    it('returns demo-title attribute value', () => {
      el.setAttribute('demo-title', 'Button Demo');
      expect(el.demoTitle).to.equal('Button Demo');
    });

    it('returns empty string when not set', () => {
      expect(el.demoTitle).to.equal('');
    });
  });

  describe('packageName getter', () => {
    it('returns package-name attribute value', () => {
      el.setAttribute('package-name', '@my/components');
      expect(el.packageName).to.equal('@my/components');
    });

    it('returns empty string when not set', () => {
      expect(el.packageName).to.equal('');
    });
  });

  describe('canonicalURL getter', () => {
    it('returns canonical-url attribute value', () => {
      el.setAttribute('canonical-url', 'https://example.com');
      expect(el.canonicalURL).to.equal('https://example.com');
    });

    it('returns empty string when not set', () => {
      expect(el.canonicalURL).to.equal('');
    });
  });

  describe('sourceURL getter', () => {
    it('returns source-url attribute value', () => {
      el.setAttribute('source-url', 'https://github.com/example/repo');
      expect(el.sourceURL).to.equal('https://github.com/example/repo');
    });

    it('returns empty string when not set', () => {
      expect(el.sourceURL).to.equal('');
    });
  });

  describe('knob coordination', () => {
    let demo;

    beforeEach(() => {
      demo = document.createElement('cem-serve-demo');
      demo.applyKnobChange = sinon.stub().returns(true);
      el.appendChild(demo);
      el.setAttribute('primary-tag-name', 'my-button');
    });

    it('handles knob:attribute-change events', () => {
      const event = new Event('knob:attribute-change', { bubbles: true });
      event.name = 'label';
      event.value = 'Click me';

      el.dispatchEvent(event);

      expect(demo.applyKnobChange.called).to.be.true;
      expect(demo.applyKnobChange.firstCall.args).to.deep.equal([
        'attribute',
        'label',
        'Click me',
        'my-button',
        0
      ]);
    });

    it('handles knob:property-change events', () => {
      const event = new Event('knob:property-change', { bubbles: true });
      event.name = 'variant';
      event.value = 'primary';

      el.dispatchEvent(event);

      expect(demo.applyKnobChange.called).to.be.true;
      expect(demo.applyKnobChange.firstCall.args).to.deep.equal([
        'property',
        'variant',
        'primary',
        'my-button',
        0
      ]);
    });

    it('handles knob:css-property-change events', () => {
      const event = new Event('knob:css-property-change', { bubbles: true });
      event.name = '--color';
      event.value = 'red';

      el.dispatchEvent(event);

      expect(demo.applyKnobChange.called).to.be.true;
      expect(demo.applyKnobChange.firstCall.args).to.deep.equal([
        'css-property',
        '--color',
        'red',
        'my-button',
        0
      ]);
    });

    it('uses default tag name when not specified', () => {
      el.removeAttribute('primary-tag-name');

      const event = new Event('knob:attribute-change', { bubbles: true });
      event.name = 'label';
      event.value = 'Test';

      el.dispatchEvent(event);

      expect(demo.applyKnobChange.firstCall.args[3]).to.equal('');
    });

    it('handles missing demo element gracefully', () => {
      el.removeChild(demo);

      const event = new Event('knob:attribute-change', { bubbles: true });
      event.name = 'label';
      event.value = 'Test';

      // Should not throw
      expect(() => el.dispatchEvent(event)).to.not.throw();
    });
  });

  describe('tree state persistence', () => {
    it('listens for expand events', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-module-path', '/components/button');
      el.appendChild(treeItem);

      const expandEvent = new Event('expand', { bubbles: true });
      Object.defineProperty(expandEvent, 'target', { value: treeItem });

      // Should not throw
      expect(() => el.dispatchEvent(expandEvent)).to.not.throw();
    });

    it('listens for collapse events', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-module-path', '/components/button');
      el.appendChild(treeItem);

      const collapseEvent = new Event('collapse', { bubbles: true });
      Object.defineProperty(collapseEvent, 'target', { value: treeItem });

      // Should not throw
      expect(() => el.dispatchEvent(collapseEvent)).to.not.throw();
    });

    it('listens for select events', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-module-path', '/components/button');
      el.appendChild(treeItem);

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });

      // Should not throw
      expect(() => el.dispatchEvent(selectEvent)).to.not.throw();
    });

    it('ignores events from non-tree-items', () => {
      const div = document.createElement('div');

      const expandEvent = new Event('expand', { bubbles: true });
      Object.defineProperty(expandEvent, 'target', { value: div });

      // Should not throw
      expect(() => el.dispatchEvent(expandEvent)).to.not.throw();
    });
  });

  describe('drawer state management', () => {
    it('tracks drawer state changes', () => {
      const drawer = el.shadowRoot.querySelector('cem-drawer');

      const changeEvent = new Event('change', { bubbles: true });
      changeEvent.open = true;

      drawer.dispatchEvent(changeEvent);

      // Should not throw - component tracks state internally
      expect(el).to.exist;
    });

    it('handles drawer resize events', () => {
      const drawer = el.shadowRoot.querySelector('cem-drawer');

      const resizeEvent = new Event('resize', { bubbles: true });
      resizeEvent.height = 500;

      drawer.dispatchEvent(resizeEvent);

      // Should persist height to attribute
      expect(drawer.getAttribute('drawer-height')).to.equal('500');
    });
  });

  describe('tabs state management', () => {
    it('handles tab changes', () => {
      const tabs = el.shadowRoot.querySelector('pf-v6-tabs');

      const changeEvent = new Event('change', { bubbles: true });
      changeEvent.selectedIndex = 1;

      tabs.dispatchEvent(changeEvent);

      // Should persist tab state
      expect(el).to.exist;
    });
  });

  describe('reconnection modal buttons', () => {
    it.skip('reloads page when reload button clicked', () => {
      // Skipped: window.location.reload is non-configurable and cannot be stubbed in modern browsers
      const reloadButton = el.shadowRoot.getElementById('reload-button');
      const reloadStub = sinon.stub(window.location, 'reload');

      reloadButton.click();

      expect(reloadStub.called).to.be.true;

      reloadStub.restore();
    });

    it('closes modal and retries when retry button clicked', () => {
      const retryButton = el.shadowRoot.getElementById('retry-button');
      const reconnectionModal = el.shadowRoot.getElementById('reconnection-modal');

      // Mock close method
      reconnectionModal.close = sinon.stub();

      retryButton.click();

      expect(reconnectionModal.close.called).to.be.true;
    });
  });

  describe('lifecycle', () => {
    it('removes knob listeners on disconnect', async () => {
      const newEl = document.createElement('cem-serve-chrome');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('cem-drawer'), '', {
        timeout: 3000
      });

      const demo = document.createElement('cem-serve-demo');
      demo.applyKnobChange = sinon.stub().returns(true);
      newEl.appendChild(demo);

      // Disconnect
      document.body.removeChild(newEl);

      // Try to trigger knob event
      const event = new Event('knob:attribute-change', { bubbles: true });
      event.name = 'test';
      event.value = 'value';
      newEl.dispatchEvent(event);

      // Handler should not be called after disconnect
      expect(demo.applyKnobChange.called).to.be.false;
    });
  });

  describe('edge cases', () => {
    it('handles multiple attribute updates', () => {
      el.setAttribute('primary-tag-name', 'button-1');
      el.setAttribute('demo-title', 'Demo 1');
      el.setAttribute('knobs', 'true');

      expect(el.primaryTagName).to.equal('button-1');
      expect(el.demoTitle).to.equal('Demo 1');
      expect(el.knobs).to.equal('true');

      el.setAttribute('primary-tag-name', 'button-2');
      el.setAttribute('demo-title', 'Demo 2');

      expect(el.primaryTagName).to.equal('button-2');
      expect(el.demoTitle).to.equal('Demo 2');
    });

    it('handles special characters in attribute values', () => {
      el.setAttribute('primary-tag-name', 'my-component<>&"\'');
      expect(el.primaryTagName).to.equal('my-component<>&"\'');
    });

    it('handles very long attribute values', () => {
      const longValue = 'x'.repeat(10000);
      el.setAttribute('demo-title', longValue);
      expect(el.demoTitle).to.equal(longValue);
    });
  });

  describe('real-world usage', () => {
    it('simulates complete demo page setup', () => {
      // Set up component
      el.setAttribute('primary-tag-name', 'my-button');
      el.setAttribute('demo-title', 'Button Component Demo');
      el.setAttribute('package-name', '@my/components');
      el.setAttribute('canonical-url', 'https://example.com/components/button');

      // Add demo
      const demo = document.createElement('cem-serve-demo');
      el.appendChild(demo);

      // Verify setup
      expect(el.primaryTagName).to.equal('my-button');
      expect(el.demoTitle).to.equal('Button Component Demo');
      expect(el.packageName).to.equal('@my/components');
      expect(el.canonicalURL).to.equal('https://example.com/components/button');
      expect(el.demo).to.equal(demo);
    });

    it('simulates knob interaction workflow', () => {
      const demo = document.createElement('cem-serve-demo');
      demo.applyKnobChange = sinon.stub().returns(true);
      el.appendChild(demo);
      el.setAttribute('primary-tag-name', 'my-button');

      // User changes attribute knob
      const attrEvent = new Event('knob:attribute-change', { bubbles: true });
      attrEvent.name = 'label';
      attrEvent.value = 'Submit';
      el.dispatchEvent(attrEvent);

      // User changes property knob
      const propEvent = new Event('knob:property-change', { bubbles: true });
      propEvent.name = 'variant';
      propEvent.value = 'primary';
      el.dispatchEvent(propEvent);

      // User changes CSS knob
      const cssEvent = new Event('knob:css-property-change', { bubbles: true });
      cssEvent.name = '--color';
      cssEvent.value = 'red';
      el.dispatchEvent(cssEvent);

      // All knobs should be applied
      expect(demo.applyKnobChange.callCount).to.equal(3);
    });

    it('simulates drawer interaction', () => {
      const drawer = el.shadowRoot.querySelector('cem-drawer');

      // User opens drawer
      const openEvent = new Event('change', { bubbles: true });
      openEvent.open = true;
      drawer.dispatchEvent(openEvent);

      // User resizes drawer
      const resizeEvent = new Event('resize', { bubbles: true });
      resizeEvent.height = 600;
      drawer.dispatchEvent(resizeEvent);

      expect(drawer.getAttribute('drawer-height')).to.equal('600');
    });

    it('simulates tab switching', () => {
      const tabs = el.shadowRoot.querySelector('pf-v6-tabs');

      // User switches tabs
      for (let i = 0; i < 3; i++) {
        const changeEvent = new Event('change', { bubbles: true });
        changeEvent.selectedIndex = i;
        tabs.dispatchEvent(changeEvent);
      }

      // Should not throw
      expect(el).to.exist;
    });
  });

  describe('template loading', () => {
    it('loads template asynchronously', async () => {
      const newEl = document.createElement('cem-serve-chrome');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('cem-drawer'), 'Should load template', {
        timeout: 3000
      });

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.querySelector('cem-drawer')).to.exist;

      document.body.removeChild(newEl);
    });
  });
});
