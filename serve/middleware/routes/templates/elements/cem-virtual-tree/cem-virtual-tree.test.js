import { expect } from '@open-wc/testing';
import './cem-virtual-tree.js';

describe('cem-virtual-tree', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('cem-virtual-tree');
    document.body.appendChild(el);

    // Set up a test manifest with nested structure
    window.__CEM_MANIFEST__ = {
      modules: [{
        path: './rh-accordion/rh-accordion-header.js',
        declarations: [{
          kind: 'class',
          name: 'RhAccordionHeader',
          customElement: true,
          tagName: 'rh-accordion-header',
          members: [
            { kind: 'field', name: 'expanded', type: { text: 'boolean' } },
            { kind: 'field', name: 'disabled', type: { text: 'boolean' } },
          ],
          events: [
            { name: 'expand', type: { text: 'Event' } },
          ],
        }]
      }]
    };

    // Wait for CemElement to load template from real server
    await el.rendered;

    // Wait for virtual tree to build and render
    await new Promise(resolve => setTimeout(resolve, 50));
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
    delete window.__CEM_MANIFEST__;
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('cem-virtual-tree');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('CemVirtualTree');
      expect(el.shadowRoot).to.exist;
    });

    it('renders viewport', () => {
      const viewport = el.shadowRoot.getElementById('viewport');
      expect(viewport).to.exist;
    });
  });

  describe('manifest building', () => {
    it('builds flat list from manifest', () => {
      const viewport = el.shadowRoot.getElementById('viewport');

      // Should have rendered the module item
      const moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      expect(moduleItem).to.exist;
      expect(moduleItem.getAttribute('label')).to.equal('./rh-accordion/rh-accordion-header.js');
    });

    it('renders module with badge showing declaration count', () => {
      const viewport = el.shadowRoot.getElementById('viewport');
      const moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');

      expect(moduleItem.getAttribute('badge')).to.equal('1');
    });

    it('marks items with children', () => {
      const viewport = el.shadowRoot.getElementById('viewport');
      const moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');

      expect(moduleItem.hasAttribute('has-children')).to.be.true;
    });
  });

  describe('nested expand/collapse', () => {
    it('expands module to show custom element', async () => {
      const viewport = el.shadowRoot.getElementById('viewport');
      let moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');

      // Initially collapsed - no nested items
      expect(moduleItem.expanded).to.be.false;
      expect(moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]')).to.not.exist;

      // Expand module
      moduleItem.expanded = true;

      // Wait for re-render
      await new Promise(resolve => setTimeout(resolve, 50));

      // Re-query after render
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');

      // Should now have custom element as child
      const customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      expect(customElement).to.exist;
      expect(customElement.getAttribute('label')).to.equal('<rh-accordion-header>');
    });

    it('expands custom element to show categories', async () => {
      const viewport = el.shadowRoot.getElementById('viewport');

      // Expand module first
      let moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      moduleItem.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Re-query and expand custom element
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      let customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      customElement.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Re-query after render
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');

      // Should have Properties and Events categories
      const propertiesCategory = customElement.querySelector('pf-v6-tree-item[data-type="category"]');
      expect(propertiesCategory).to.exist;
      expect(propertiesCategory.getAttribute('label')).to.equal('Properties');
    });

    it('collapsing nested item does NOT collapse parent items', async () => {
      const viewport = el.shadowRoot.getElementById('viewport');

      // Step 1: Expand module
      let moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      moduleItem.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Step 2: Expand custom element
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      let customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      customElement.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Step 3: Expand properties category
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      let propertiesCategory = customElement.querySelector('pf-v6-tree-item[data-type="category"]');
      propertiesCategory.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Step 4: Verify all three levels are expanded
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      propertiesCategory = customElement.querySelector('pf-v6-tree-item[data-type="category"]');

      expect(moduleItem.expanded, 'module should be expanded').to.be.true;
      expect(customElement.expanded, 'custom element should be expanded').to.be.true;
      expect(propertiesCategory.expanded, 'properties category should be expanded').to.be.true;

      // Step 5: Collapse the properties category
      propertiesCategory.expanded = false;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Step 6: Verify ONLY properties collapsed, parents stay expanded
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      propertiesCategory = customElement.querySelector('pf-v6-tree-item[data-type="category"]');

      expect(moduleItem.expanded, 'module should STILL be expanded').to.be.true;
      expect(customElement.expanded, 'custom element should STILL be expanded').to.be.true;
      expect(propertiesCategory.expanded, 'properties category should be collapsed').to.be.false;
    });

    it('collapsing parent hides all descendants', async () => {
      const viewport = el.shadowRoot.getElementById('viewport');

      // Expand all three levels
      let moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      moduleItem.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      let customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      customElement.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      let propertiesCategory = customElement.querySelector('pf-v6-tree-item[data-type="category"]');
      propertiesCategory.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Verify all expanded
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      propertiesCategory = customElement.querySelector('pf-v6-tree-item[data-type="category"]');

      expect(moduleItem.expanded).to.be.true;
      expect(customElement.expanded).to.be.true;
      expect(propertiesCategory.expanded).to.be.true;

      // Collapse the module (top level)
      moduleItem.expanded = false;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Re-query
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');

      // Module should be collapsed and have no visible children
      expect(moduleItem.expanded).to.be.false;
      expect(moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]')).to.not.exist;
    });
  });

  describe('search functionality', () => {
    it('provides search method', () => {
      expect(el.search).to.be.a('function');
    });

    it('filters items by label', async () => {
      const viewport = el.shadowRoot.getElementById('viewport');

      el.search('accordion');
      await new Promise(resolve => setTimeout(resolve, 50));

      const moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      expect(moduleItem).to.exist;

      // Should auto-expand to show matching items
      expect(moduleItem.expanded).to.be.true;
    });

    it('clears filter when search is empty', async () => {
      const viewport = el.shadowRoot.getElementById('viewport');

      el.search('accordion');
      await new Promise(resolve => setTimeout(resolve, 50));

      el.search('');
      await new Promise(resolve => setTimeout(resolve, 50));

      const moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      expect(moduleItem).to.exist;
    });
  });

  describe('selection state management', () => {
    it('only one item has current attribute at a time', async () => {
      const viewport = el.shadowRoot.getElementById('viewport');

      // Expand tree to access properties
      let moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      moduleItem.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      let customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      customElement.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      let propertiesCategory = customElement.querySelector('pf-v6-tree-item[data-type="category"]');
      propertiesCategory.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      // Get the two property items
      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      propertiesCategory = customElement.querySelector('pf-v6-tree-item[data-type="category"]');
      const properties = propertiesCategory.querySelectorAll('pf-v6-tree-item[data-type="property"]');

      expect(properties.length).to.be.at.least(2);

      const property1 = properties[0];
      const property2 = properties[1];

      // Select first property
      property1.dispatchEvent(new Event('select', { bubbles: true }));
      await new Promise(resolve => setTimeout(resolve, 50));

      // Check only first has current
      expect(property1.hasAttribute('current'), 'property1 should have current after selection').to.be.true;
      expect(property2.hasAttribute('current'), 'property2 should NOT have current').to.be.false;

      // Select second property
      property2.dispatchEvent(new Event('select', { bubbles: true }));
      await new Promise(resolve => setTimeout(resolve, 50));

      // Check only second has current now
      expect(property1.hasAttribute('current'), 'property1 should NOT have current after selecting property2').to.be.false;
      expect(property2.hasAttribute('current'), 'property2 should have current').to.be.true;
    });

    it('clears current from all items when searching', async () => {
      const viewport = el.shadowRoot.getElementById('viewport');

      // Expand and select an item
      let moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      moduleItem.expanded = true;
      await new Promise(resolve => setTimeout(resolve, 50));

      moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      let customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      customElement.dispatchEvent(new Event('select', { bubbles: true }));
      await new Promise(resolve => setTimeout(resolve, 50));

      // Verify current is set
      customElement = viewport.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      expect(customElement.hasAttribute('current')).to.be.true;

      // Search
      el.search('accordion');
      await new Promise(resolve => setTimeout(resolve, 50));

      // Current should still be maintained (search doesn't affect selection)
      customElement = viewport.querySelector('pf-v6-tree-item[data-type="custom-element"]');
      expect(customElement.hasAttribute('current')).to.be.true;
    });
  });

  describe('expand/collapse all', () => {
    it('provides expandAll method', () => {
      expect(el.expandAll).to.be.a('function');
    });

    it('provides collapseAll method', () => {
      expect(el.collapseAll).to.be.a('function');
    });

    it('expandAll expands all items', async () => {
      el.expandAll();
      await new Promise(resolve => setTimeout(resolve, 50));

      const viewport = el.shadowRoot.getElementById('viewport');
      const moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');
      const customElement = moduleItem.querySelector('pf-v6-tree-item[data-type="custom-element"]');

      expect(moduleItem.expanded).to.be.true;
      expect(customElement.expanded).to.be.true;
    });

    it('collapseAll collapses all items', async () => {
      // First expand all
      el.expandAll();
      await new Promise(resolve => setTimeout(resolve, 50));

      // Then collapse all
      el.collapseAll();
      await new Promise(resolve => setTimeout(resolve, 50));

      const viewport = el.shadowRoot.getElementById('viewport');
      const moduleItem = viewport.querySelector('pf-v6-tree-item[data-type="module"]');

      expect(moduleItem.expanded).to.be.false;
      expect(moduleItem.querySelector('pf-v6-tree-item')).to.not.exist;
    });
  });
});
