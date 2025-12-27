import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './cem-manifest-browser.js';
import { CemVirtualTree } from '/__cem/elements/cem-virtual-tree/cem-virtual-tree.js';

describe('cem-manifest-browser', () => {
  let el;
  let fetchStub;

  beforeEach(async () => {
    // Clear static cache before each test
    CemVirtualTree.clearCache();

    // Store original fetch BEFORE creating element
    const originalFetch = window.fetch.bind(window);

    // Stub fetch BEFORE creating element (so child components get mocked fetch)
    fetchStub = sinon.stub(window, 'fetch').callsFake((url, ...args) => {
      if (url === '/custom-elements.json') {
        return Promise.resolve({
          ok: true,
          json: async () => ({ modules: [] })
        });
      }
      return originalFetch(url, ...args);
    });

    el = document.createElement('cem-manifest-browser');
    document.body.appendChild(el);

    // Wait for CemElement to load template from real server
    await el.rendered;
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
    if (fetchStub) {
      fetchStub.restore();
    }
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('cem-manifest-browser');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('CemManifestBrowser');
      expect(el.shadowRoot).to.exist;
    });

    it('renders search input', () => {
      const searchInput = el.shadowRoot.getElementById('search');
      expect(searchInput).to.exist;
    });

    it('renders search count badge', () => {
      const searchCount = el.shadowRoot.getElementById('search-count');
      expect(searchCount).to.exist;
    });

    it('renders search clear button', () => {
      const searchClear = el.shadowRoot.getElementById('search-clear');
      expect(searchClear).to.exist;
    });

    it('renders drawer element', () => {
      const drawer = el.shadowRoot.getElementById('drawer');
      expect(drawer).to.exist;
    });

    it('renders virtual tree component', () => {
      const virtualTree = el.shadowRoot.getElementById('virtual-tree');
      expect(virtualTree).to.exist;
      expect(virtualTree.tagName.toLowerCase()).to.equal('cem-virtual-tree');
    });

    it('renders detail panel component', () => {
      const detailPanel = el.shadowRoot.getElementById('detail-panel');
      expect(detailPanel).to.exist;
      expect(detailPanel.tagName.toLowerCase()).to.equal('cem-detail-panel');
    });

    it('renders expand all button', () => {
      const expandAllBtn = el.shadowRoot.getElementById('expand-all');
      expect(expandAllBtn).to.exist;
    });

    it('renders collapse all button', () => {
      const collapseAllBtn = el.shadowRoot.getElementById('collapse-all');
      expect(collapseAllBtn).to.exist;
    });

    it('hides clear button initially', () => {
      const searchClear = el.shadowRoot.getElementById('search-clear');
      expect(searchClear.hidden).to.be.true;
    });

    it('hides search count badge initially', () => {
      const searchCount = el.shadowRoot.getElementById('search-count');
      expect(searchCount.hidden).to.be.true;
    });
  });

  describe('virtual tree integration', () => {
    let virtualTree, detailPanel, drawer, localFetchStub;

    const testManifest = {
      modules: [{
        path: './test.js',
        declarations: [{
          kind: 'class',
          name: 'TestElement',
          customElement: true,
          tagName: 'test-element'
        }]
      }]
    };

    beforeEach(async () => {
      // Restore the outer stub and replace with test manifest stub
      fetchStub.restore();

      // Store original fetch before stubbing
      const originalFetch = window.fetch.bind(window);

      // Mock fetch to return test manifest, but call through for other URLs
      localFetchStub = sinon.stub(window, 'fetch').callsFake((url, ...args) => {
        if (url === '/custom-elements.json') {
          return Promise.resolve({
            ok: true,
            json: async () => testManifest
          });
        }
        return originalFetch(url, ...args);
      });

      virtualTree = el.shadowRoot.getElementById('virtual-tree');
      detailPanel = el.shadowRoot.getElementById('detail-panel');
      drawer = el.shadowRoot.getElementById('drawer');

      // Wait for virtual tree to load the manifest
      await virtualTree.rendered;
    });

    afterEach(() => {
      if (localFetchStub) {
        localFetchStub.restore();
      }
      // Restore the outer stub for other tests
      const originalFetch = window.fetch.bind(window);
      fetchStub = sinon.stub(window, 'fetch').callsFake((url, ...args) => {
        if (url === '/custom-elements.json') {
          return Promise.resolve({
            ok: true,
            json: async () => ({ modules: [] })
          });
        }
        return originalFetch(url, ...args);
      });
    });

    it('listens for item-select events from virtual tree', async () => {
      const renderSpy = sinon.spy(detailPanel, 'renderItem');

      const itemData = {
        id: 1,
        type: 'custom-element',
        tagName: 'test-element',
        modulePath: './test.js'
      };

      // Simulate virtual tree emitting item-select event
      const event = new CustomEvent('item-select', { bubbles: true, composed: true });
      event.item = itemData;
      virtualTree.dispatchEvent(event);

      // Wait for async render and manifest fetch
      await new Promise(resolve => setTimeout(resolve, 50));

      expect(renderSpy.calledOnce).to.be.true;
      expect(renderSpy.calledWith(itemData, testManifest)).to.be.true;

      renderSpy.restore();
    });

    it('opens drawer when item is selected', async () => {
      const itemData = {
        id: 1,
        type: 'custom-element',
        tagName: 'test-element',
        modulePath: './test.js'
      };

      expect(drawer.expanded).to.be.false;

      const event = new CustomEvent('item-select', { bubbles: true, composed: true });
      event.item = itemData;
      virtualTree.dispatchEvent(event);

      await new Promise(resolve => setTimeout(resolve, 10));

      expect(drawer.expanded).to.be.true;
    });
  });

  describe('search functionality', () => {
    let searchInput, virtualTree;

    beforeEach(async () => {
      searchInput = el.shadowRoot.getElementById('search');
      virtualTree = el.shadowRoot.getElementById('virtual-tree');
      await virtualTree.rendered;
    });

    it('shows clear button when search has value', () => {
      const searchClear = el.shadowRoot.getElementById('search-clear');

      searchInput.value = 'test';
      searchInput.dispatchEvent(new Event('input'));

      expect(searchClear.hidden).to.be.false;
    });

    it('hides clear button when search is empty', () => {
      const searchClear = el.shadowRoot.getElementById('search-clear');

      searchInput.value = 'test';
      searchInput.dispatchEvent(new Event('input'));

      searchInput.value = '';
      searchInput.dispatchEvent(new Event('input'));

      expect(searchClear.hidden).to.be.true;
    });

    it('debounces search input', async () => {
      const searchSpy = sinon.spy(virtualTree, 'search');

      searchInput.value = 'B';
      searchInput.dispatchEvent(new Event('input'));

      searchInput.value = 'Bu';
      searchInput.dispatchEvent(new Event('input'));

      searchInput.value = 'But';
      searchInput.dispatchEvent(new Event('input'));

      // Should not have called search yet
      expect(searchSpy.called).to.be.false;

      // Wait for debounce (300ms)
      await new Promise(resolve => setTimeout(resolve, 350));

      // Should have called search once with final value
      expect(searchSpy.calledOnce).to.be.true;
      expect(searchSpy.calledWith('But')).to.be.true;

      searchSpy.restore();
    });

    it('delegates search to virtual tree', async () => {
      const searchSpy = sinon.spy(virtualTree, 'search');

      searchInput.value = 'test';
      searchInput.dispatchEvent(new Event('input'));

      await new Promise(resolve => setTimeout(resolve, 350));

      expect(searchSpy.calledOnce).to.be.true;
      expect(searchSpy.calledWith('test')).to.be.true;

      searchSpy.restore();
    });

    it('clears search when clear button clicked', async () => {
      const searchClear = el.shadowRoot.getElementById('search-clear');
      const searchSpy = sinon.spy(virtualTree, 'search');

      searchInput.value = 'test';
      searchInput.dispatchEvent(new Event('input'));

      await new Promise(resolve => setTimeout(resolve, 350));

      searchClear.click();

      await new Promise(resolve => setTimeout(resolve, 350));

      expect(searchInput.value).to.equal('');
      expect(searchClear.hidden).to.be.true;
      expect(searchSpy.calledWith('')).to.be.true;

      searchSpy.restore();
    });

    it('hides search count when search is cleared', async () => {
      const searchClear = el.shadowRoot.getElementById('search-clear');
      const searchCount = el.shadowRoot.getElementById('search-count');

      searchInput.value = 'test';
      searchInput.dispatchEvent(new Event('input'));

      await new Promise(resolve => setTimeout(resolve, 350));

      searchClear.click();

      await new Promise(resolve => setTimeout(resolve, 350));

      expect(searchCount.hidden).to.be.true;
    });

    it('handles very long search queries', async () => {
      const longQuery = 'x'.repeat(1000);

      searchInput.value = longQuery;
      searchInput.dispatchEvent(new Event('input'));

      await new Promise(resolve => setTimeout(resolve, 350));

      // Should not throw or crash
      expect(el).to.exist;
    });
  });

  describe('expand/collapse functionality', () => {
    let virtualTree;

    beforeEach(async () => {
      virtualTree = el.shadowRoot.getElementById('virtual-tree');
      await virtualTree.rendered;
    });

    it('delegates expandAll to virtual tree', () => {
      const expandSpy = sinon.spy(virtualTree, 'expandAll');
      const expandAllBtn = el.shadowRoot.getElementById('expand-all');

      expandAllBtn.click();

      expect(expandSpy.calledOnce).to.be.true;

      expandSpy.restore();
    });

    it('delegates collapseAll to virtual tree', () => {
      const collapseSpy = sinon.spy(virtualTree, 'collapseAll');
      const collapseAllBtn = el.shadowRoot.getElementById('collapse-all');

      collapseAllBtn.click();

      expect(collapseSpy.calledOnce).to.be.true;

      collapseSpy.restore();
    });
  });

  describe('template loading', () => {
    it('loads template asynchronously', async () => {
      const newEl = document.createElement('cem-manifest-browser');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.getElementById('search'), 'Should load template', {
        timeout: 3000
      });

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.getElementById('search')).to.exist;
      expect(newEl.shadowRoot.getElementById('virtual-tree')).to.exist;
      expect(newEl.shadowRoot.getElementById('detail-panel')).to.exist;

      document.body.removeChild(newEl);
    });
  });
});
