import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './cem-manifest-browser.js';

describe('cem-manifest-browser', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('cem-manifest-browser');
    document.body.appendChild(el);

    // Wait for CemElement to load template from real server
    await el.rendered;
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
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

  describe('tree item selection', () => {
    let treeItem, detailElement;

    beforeEach(() => {
      // Create a mock tree item
      treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-module-path', '/components/button');
      treeItem.setAttribute('data-name', 'Button');

      // Create a matching detail element in body
      detailElement = document.createElement('div');
      detailElement.setAttribute('data-type', 'class');
      detailElement.setAttribute('data-module-path', '/components/button');
      detailElement.setAttribute('data-name', 'Button');
      detailElement.hidden = true;
      document.body.appendChild(detailElement);
    });

    afterEach(() => {
      if (detailElement && detailElement.parentNode) {
        document.body.removeChild(detailElement);
      }
    });

    it('opens drawer when tree item is selected', () => {
      const drawer = el.shadowRoot.getElementById('drawer');

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      expect(drawer.expanded).to.be.true;
    });

    it('shows detail element when tree item is selected', () => {
      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      expect(detailElement.hidden).to.be.false;
    });

    it('hides previous detail when selecting new item', () => {
      const detailElement2 = document.createElement('div');
      detailElement2.setAttribute('data-type', 'class');
      detailElement2.setAttribute('data-module-path', '/components/input');
      detailElement2.setAttribute('data-name', 'Input');
      detailElement2.hidden = true;
      document.body.appendChild(detailElement2);

      // Select first item
      const selectEvent1 = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent1, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent1);

      expect(detailElement.hidden).to.be.false;

      // Select second item
      const treeItem2 = document.createElement('pf-v6-tree-item');
      treeItem2.setAttribute('data-type', 'class');
      treeItem2.setAttribute('data-module-path', '/components/input');
      treeItem2.setAttribute('data-name', 'Input');

      const selectEvent2 = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent2, 'target', { value: treeItem2 });
      el.dispatchEvent(selectEvent2);

      // First should be hidden, second should be visible
      expect(detailElement.hidden).to.be.true;
      expect(detailElement2.hidden).to.be.false;

      document.body.removeChild(detailElement2);
    });

    it('ignores category items', () => {
      const categoryItem = document.createElement('pf-v6-tree-item');
      categoryItem.setAttribute('data-type', 'category');

      const drawer = el.shadowRoot.getElementById('drawer');
      const initialExpanded = drawer.expanded;

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: categoryItem });
      el.dispatchEvent(selectEvent);

      // Drawer state should not change
      expect(drawer.expanded).to.equal(initialExpanded);
    });

    it('ignores non-tree-item events', () => {
      const div = document.createElement('div');

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: div });

      // Should not throw
      expect(() => el.dispatchEvent(selectEvent)).to.not.throw();
    });
  });

  describe('search functionality', () => {
    let searchInput, treeView, treeItem1, treeItem2;

    beforeEach(() => {
      searchInput = el.shadowRoot.getElementById('search');

      // Create a mock tree view with items
      treeView = document.createElement('pf-v6-tree-view');
      treeView.setAttribute('slot', 'manifest-tree');

      treeItem1 = document.createElement('pf-v6-tree-item');
      treeItem1.setAttribute('label', 'Button Component');

      treeItem2 = document.createElement('pf-v6-tree-item');
      treeItem2.setAttribute('label', 'Input Component');

      treeView.appendChild(treeItem1);
      treeView.appendChild(treeItem2);

      el.appendChild(treeView);
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
      const spy = sinon.spy();

      // Spy on querySelector to detect when search runs
      const originalQuerySelectorAll = Element.prototype.querySelectorAll;
      sinon.stub(Element.prototype, 'querySelectorAll').callsFake(function(...args) {
        if (args[0] === 'pf-v6-tree-item') {
          spy();
        }
        return originalQuerySelectorAll.apply(this, args);
      });

      searchInput.value = 'B';
      searchInput.dispatchEvent(new Event('input'));

      searchInput.value = 'Bu';
      searchInput.dispatchEvent(new Event('input'));

      searchInput.value = 'But';
      searchInput.dispatchEvent(new Event('input'));

      // Should not have run search yet
      expect(spy.called).to.be.false;

      // Wait for debounce
      await new Promise(resolve => setTimeout(resolve, 350));

      // Should have run search once
      expect(spy.called).to.be.true;

      Element.prototype.querySelectorAll.restore();
    });

    it('clears search when clear button clicked', async () => {
      const searchClear = el.shadowRoot.getElementById('search-clear');

      searchInput.value = 'test';
      searchInput.dispatchEvent(new Event('input'));

      await new Promise(resolve => setTimeout(resolve, 10));

      searchClear.click();

      expect(searchInput.value).to.equal('');
      expect(searchClear.hidden).to.be.true;
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
  });

  describe('expand/collapse functionality', () => {
    let treeView, treeItem1, treeItem2;

    beforeEach(() => {
      // Create mock tree view
      treeView = document.createElement('pf-v6-tree-view');
      treeView.setAttribute('slot', 'manifest-tree');

      treeItem1 = document.createElement('pf-v6-tree-item');
      treeItem1.expanded = false;

      treeItem2 = document.createElement('pf-v6-tree-item');
      treeItem2.expanded = false;

      treeView.appendChild(treeItem1);
      treeView.appendChild(treeItem2);

      el.appendChild(treeView);
    });

    it('expands all items when expand all button clicked', () => {
      const expandAllBtn = el.shadowRoot.getElementById('expand-all');

      expandAllBtn.click();

      expect(treeItem1.expanded).to.be.true;
      expect(treeItem2.expanded).to.be.true;
    });

    it('collapses all items when collapse all button clicked', () => {
      const collapseAllBtn = el.shadowRoot.getElementById('collapse-all');

      treeItem1.expanded = true;
      treeItem2.expanded = true;

      collapseAllBtn.click();

      expect(treeItem1.expanded).to.be.false;
      expect(treeItem2.expanded).to.be.false;
    });

    it('handles missing tree view gracefully', () => {
      // Remove tree view
      if (treeView.parentNode) {
        treeView.parentNode.removeChild(treeView);
      }

      const expandAllBtn = el.shadowRoot.getElementById('expand-all');
      const collapseAllBtn = el.shadowRoot.getElementById('collapse-all');

      // Should not throw
      expect(() => expandAllBtn.click()).to.not.throw();
      expect(() => collapseAllBtn.click()).to.not.throw();
    });
  });

  describe('selector building', () => {
    it('builds selector with type only', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');

      const detailElement = document.createElement('div');
      detailElement.setAttribute('data-type', 'class');
      document.body.appendChild(detailElement);
      detailElement.hidden = true;

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      expect(detailElement.hidden).to.be.false;

      document.body.removeChild(detailElement);
    });

    it('builds selector with type and module-path', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-module-path', '/components/button');

      const detailElement = document.createElement('div');
      detailElement.setAttribute('data-type', 'class');
      detailElement.setAttribute('data-module-path', '/components/button');
      document.body.appendChild(detailElement);
      detailElement.hidden = true;

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      expect(detailElement.hidden).to.be.false;

      document.body.removeChild(detailElement);
    });

    it('builds selector with type and tag-name', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'custom-element');
      treeItem.setAttribute('data-tag-name', 'my-button');

      const detailElement = document.createElement('div');
      detailElement.setAttribute('data-type', 'custom-element');
      detailElement.setAttribute('data-tag-name', 'my-button');
      document.body.appendChild(detailElement);
      detailElement.hidden = true;

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      expect(detailElement.hidden).to.be.false;

      document.body.removeChild(detailElement);
    });

    it('builds selector with all attributes', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'method');
      treeItem.setAttribute('data-module-path', '/components/button');
      treeItem.setAttribute('data-tag-name', 'my-button');
      treeItem.setAttribute('data-name', 'onClick');

      const detailElement = document.createElement('div');
      detailElement.setAttribute('data-type', 'method');
      detailElement.setAttribute('data-module-path', '/components/button');
      detailElement.setAttribute('data-tag-name', 'my-button');
      detailElement.setAttribute('data-name', 'onClick');
      document.body.appendChild(detailElement);
      detailElement.hidden = true;

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      expect(detailElement.hidden).to.be.false;

      document.body.removeChild(detailElement);
    });

    it('handles data-path as fallback for data-module-path', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-path', '/components/input');

      const detailElement = document.createElement('div');
      detailElement.setAttribute('data-type', 'class');
      detailElement.setAttribute('data-module-path', '/components/input');
      document.body.appendChild(detailElement);
      detailElement.hidden = true;

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      expect(detailElement.hidden).to.be.false;

      document.body.removeChild(detailElement);
    });
  });

  describe('edge cases', () => {
    it('handles selecting item with no matching details', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-name', 'NonExistent');

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });

      // Should not throw
      expect(() => el.dispatchEvent(selectEvent)).to.not.throw();
    });

    it('handles multiple detail elements matching selector', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');

      const detail1 = document.createElement('div');
      detail1.setAttribute('data-type', 'class');
      detail1.hidden = true;
      document.body.appendChild(detail1);

      const detail2 = document.createElement('div');
      detail2.setAttribute('data-type', 'class');
      detail2.hidden = true;
      document.body.appendChild(detail2);

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      // Both should be shown
      expect(detail1.hidden).to.be.false;
      expect(detail2.hidden).to.be.false;

      document.body.removeChild(detail1);
      document.body.removeChild(detail2);
    });

    it('handles special characters in attribute values', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-name', 'Class<T>');

      const detailElement = document.createElement('div');
      detailElement.setAttribute('data-type', 'class');
      detailElement.setAttribute('data-name', 'Class<T>');
      detailElement.hidden = true;
      document.body.appendChild(detailElement);

      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      expect(detailElement.hidden).to.be.false;

      document.body.removeChild(detailElement);
    });

    it('handles very long search queries', async () => {
      const searchInput = el.shadowRoot.getElementById('search');
      const longQuery = 'x'.repeat(1000);

      searchInput.value = longQuery;
      searchInput.dispatchEvent(new Event('input'));

      await new Promise(resolve => setTimeout(resolve, 350));

      // Should not throw or crash
      expect(el).to.exist;
    });

    it('handles empty tree view', () => {
      const treeView = document.createElement('pf-v6-tree-view');

      const expandAllBtn = el.shadowRoot.getElementById('expand-all');
      const collapseAllBtn = el.shadowRoot.getElementById('collapse-all');

      // Should not throw
      expect(() => expandAllBtn.click()).to.not.throw();
      expect(() => collapseAllBtn.click()).to.not.throw();
    });
  });

  describe('real-world usage', () => {
    it('simulates browsing manifest and viewing details', () => {
      const treeItem = document.createElement('pf-v6-tree-item');
      treeItem.setAttribute('data-type', 'class');
      treeItem.setAttribute('data-module-path', '/components/button');
      treeItem.setAttribute('data-name', 'Button');

      const detailElement = document.createElement('div');
      detailElement.setAttribute('data-type', 'class');
      detailElement.setAttribute('data-module-path', '/components/button');
      detailElement.setAttribute('data-name', 'Button');
      detailElement.hidden = true;
      document.body.appendChild(detailElement);

      const drawer = el.shadowRoot.getElementById('drawer');

      // User clicks tree item
      const selectEvent = new Event('select', { bubbles: true });
      Object.defineProperty(selectEvent, 'target', { value: treeItem });
      el.dispatchEvent(selectEvent);

      // Drawer opens and detail is shown
      expect(drawer.expanded).to.be.true;
      expect(detailElement.hidden).to.be.false;

      document.body.removeChild(detailElement);
    });

    it('simulates searching for a component', async () => {
      const searchInput = el.shadowRoot.getElementById('search');
      const searchClear = el.shadowRoot.getElementById('search-clear');

      // User types in search
      searchInput.value = 'button';
      searchInput.dispatchEvent(new Event('input'));

      // Clear button appears
      expect(searchClear.hidden).to.be.false;

      // User clears search
      searchClear.click();

      // Input is cleared
      expect(searchInput.value).to.equal('');
      expect(searchClear.hidden).to.be.true;
    });

    it('simulates expanding all tree items', () => {
      const treeView = document.createElement('pf-v6-tree-view');
      treeView.setAttribute('slot', 'manifest-tree');
      const items = [];

      for (let i = 0; i < 5; i++) {
        const item = document.createElement('pf-v6-tree-item');
        item.expanded = false;
        items.push(item);
        treeView.appendChild(item);
      }

      el.appendChild(treeView);

      const expandAllBtn = el.shadowRoot.getElementById('expand-all');
      expandAllBtn.click();

      items.forEach(item => {
        expect(item.expanded).to.be.true;
      });
    });

    it('simulates collapsing all tree items', () => {
      const treeView = document.createElement('pf-v6-tree-view');
      treeView.setAttribute('slot', 'manifest-tree');
      const items = [];

      for (let i = 0; i < 5; i++) {
        const item = document.createElement('pf-v6-tree-item');
        item.expanded = true;
        items.push(item);
        treeView.appendChild(item);
      }

      el.appendChild(treeView);

      const collapseAllBtn = el.shadowRoot.getElementById('collapse-all');
      collapseAllBtn.click();

      items.forEach(item => {
        expect(item.expanded).to.be.false;
      });
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

      document.body.removeChild(newEl);
    });
  });
});
