import '/__cem/elements/pf-v6-drawer/pf-v6-drawer.js';
import { CemElement } from '/__cem/cem-element.js';

/**
 * Manifest Browser with tree navigation and detail drawer
 * @customElement cem-manifest-browser
 */
export class CemManifestBrowser extends CemElement {
  static is = 'cem-manifest-browser';

  #panelTitle;
  #drawer;
  #currentDetail;

  afterTemplateLoaded() {
    this.#panelTitle = this.shadowRoot.getElementById('panel-title');
    this.#drawer = this.shadowRoot.getElementById('drawer');

    // Debug: Check how many detail panels exist
    const allDetailPanels = this.querySelectorAll('[slot="manifest-details"]');
    console.log('Total detail panels found: ' + allDetailPanels.length);
    if (allDetailPanels.length > 0) {
      const first = allDetailPanels[0];
      console.log('First panel tagName: ' + first.tagName);
      console.log('First panel outerHTML: ' + first.outerHTML.substring(0, 200));
      console.log('First panel data-type: ' + first.getAttribute('data-type'));
      console.log('First panel data-module-path: ' + first.getAttribute('data-module-path'));
      console.log('First panel data-tag-name: ' + first.getAttribute('data-tag-name'));
      console.log('First panel data-name: ' + first.getAttribute('data-name'));
    }

    // Also check all elements with data-type
    const allDataTypeElements = this.querySelectorAll('[data-type]');
    console.log('Total elements with data-type: ' + allDataTypeElements.length);
    if (allDataTypeElements.length > 0) {
      console.log('First data-type element: ' + allDataTypeElements[0].tagName + ' with data-type=' + allDataTypeElements[0].getAttribute('data-type'));
    }

    // Listen for tree item selections
    this.addEventListener('select', (e) => {
      const treeItem = e.target;
      if (treeItem.tagName !== 'PF-V6-TREE-ITEM') return;

      this.#handleItemSelect(treeItem);
    });
  }

  #handleItemSelect(treeItem) {
    console.log('handleItemSelect called');
    const type = treeItem.getAttribute('data-type');
    console.log('type: ' + type);

    if (!type || type === 'category') {
      console.log('Skipping category or no type');
      return;
    }

    // Build selector from tree item attributes
    const selector = this.#buildSelector(type, treeItem);
    console.log('Built selector: ' + selector);

    // Find all matching elements (both h3 and dl)
    const matchingElements = document.body.querySelectorAll(selector);
    console.log('Found matching elements: ' + matchingElements.length);

    if (matchingElements.length === 0) {
      console.warn('No elements found for selector: ' + selector);
      return;
    }

    // Hide current detail elements if exist
    if (this.#currentDetail && this.#currentDetail.length > 0) {
      console.log('Hiding previous detail elements: ' + this.#currentDetail.length);
      this.#currentDetail.forEach(el => el.hidden = true);
    }

    // Show new detail elements
    console.log('Showing new detail elements');
    matchingElements.forEach(el => el.hidden = false);
    this.#currentDetail = matchingElements;

    // Open drawer
    this.#drawer.expanded = true;
    console.log('Drawer expanded');
  }

  #buildSelector(type, treeItem) {
    const selectors = [`[data-type="${type}"]`];

    const modulePath = treeItem.getAttribute('data-module-path') || treeItem.getAttribute('data-path');
    const tagName = treeItem.getAttribute('data-tag-name');
    const name = treeItem.getAttribute('data-name');

    console.log('Building selector - modulePath: ' + modulePath + ', tagName: ' + tagName + ', name: ' + name);

    if (modulePath) {
      selectors.push(`[data-module-path="${modulePath}"]`);
    }
    if (tagName) {
      selectors.push(`[data-tag-name="${tagName}"]`);
    }
    if (name) {
      selectors.push(`[data-name="${name}"]`);
    }

    return selectors.join('');
  }

  static {
    customElements.define(this.is, this);
  }
}
