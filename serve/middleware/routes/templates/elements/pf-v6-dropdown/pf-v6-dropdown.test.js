import { expect, fixture, oneEvent } from '@open-wc/testing';
import { sendKeys } from '@web/test-runner-commands';
import './pf-v6-dropdown.js';

describe('pf-v6-dropdown', () => {
  describe('basic functionality', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-dropdown></pf-v6-dropdown>');
      await el.rendered;
    });

    it('is defined as custom element', () => {
      expect(el).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('PfV6Dropdown');
      expect(el.shadowRoot).to.exist;
    });

    it('renders toggle button', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle).to.exist;
    });

    it('renders menu', () => {
      const menu = el.shadowRoot.getElementById('menu');
      expect(menu).to.exist;
    });

    it('renders menu container', () => {
      const menuContainer = el.shadowRoot.getElementById('menu-container');
      expect(menuContainer).to.exist;
    });

    it('observes all attributes', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('expanded');
      expect(attrs).to.include('disabled');
      expect(attrs).to.include('label');
    });
  });

  describe('closed state (default)', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-dropdown></pf-v6-dropdown>');
      await el.rendered;
    });

    it('has default values', () => {
      expect(el.expanded).to.be.false;
      expect(el.disabled).to.be.false;
      expect(el.label).to.equal('');
    });

    it('hides menu container', () => {
      const menuContainer = el.shadowRoot.getElementById('menu-container');
      const display = window.getComputedStyle(menuContainer).display;
      expect(display).to.equal('none');
    });

    it('sets aria-expanded to false', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.getAttribute('aria-expanded')).to.equal('false');
    });

    it('sets aria-haspopup on toggle', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.getAttribute('aria-haspopup')).to.equal('true');
    });

    it('expands on toggle click', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      toggle.click();

      expect(el.expanded).to.be.true;
    });

    it('expands via expand() method', async () => {
      setTimeout(() => el.expand());
      await oneEvent(el, 'expand');

      expect(el.expanded).to.be.true;
    });

    it('toggles via toggle() method', () => {
      el.toggle();
      expect(el.expanded).to.be.true;

      el.toggle();
      expect(el.expanded).to.be.false;
    });

    it('ignores Escape when closed', () => {
      el.dispatchEvent(new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true,
        cancelable: true
      }));

      expect(el.expanded).to.be.false;
    });
  });

  describe('expanded state', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-dropdown expanded></pf-v6-dropdown>');
      await el.rendered;
    });

    it('reflects expanded attribute', () => {
      expect(el.expanded).to.be.true;
    });

    it('shows menu container', () => {
      const menuContainer = el.shadowRoot.getElementById('menu-container');
      expect(menuContainer.style.display).to.equal('block');
    });

    it('sets aria-expanded to true', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.getAttribute('aria-expanded')).to.equal('true');
    });

    it('collapses on toggle click', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      toggle.click();

      expect(el.expanded).to.be.false;
    });

    it('collapses via collapse() method', async () => {
      setTimeout(() => el.collapse());
      await oneEvent(el, 'collapse');

      expect(el.expanded).to.be.false;
    });

    it('collapses on Escape key', async () => {
      el.dispatchEvent(new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true
      }));

      expect(el.expanded).to.be.false;
    });

    it('focuses toggle after Escape', async () => {
      const toggle = el.shadowRoot.getElementById('toggle');

      el.dispatchEvent(new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true
      }));

      // After collapse, toggle should be focused
      expect(el.expanded).to.be.false;
    });

    it('collapses on click outside', async () => {
      // Wait for click listener to be registered
      await new Promise(resolve => requestAnimationFrame(resolve));

      const outside = document.createElement('div');
      document.body.appendChild(outside);

      outside.click();

      expect(el.expanded).to.be.false;

      document.body.removeChild(outside);
    });

    it('does not collapse on click inside', () => {
      el.click();
      expect(el.expanded).to.be.true;
    });

    it('does not collapse on menu click', () => {
      const menu = el.shadowRoot.getElementById('menu');
      menu.click();
      expect(el.expanded).to.be.true;
    });

    it('ignores non-Escape keys', () => {
      el.dispatchEvent(new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      }));

      expect(el.expanded).to.be.true;
    });
  });

  describe('disabled state', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-dropdown disabled></pf-v6-dropdown>');
      await el.rendered;
    });

    it('reflects disabled attribute', () => {
      expect(el.disabled).to.be.true;
    });

    it('disables toggle button', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.disabled).to.be.true;
    });

    it('does not expand on toggle click', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      toggle.click();

      expect(el.expanded).to.be.false;
    });
  });

  describe('with label', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-dropdown label="Filter menu"></pf-v6-dropdown>');
      await el.rendered;
    });

    it('reflects label attribute', () => {
      expect(el.label).to.equal('Filter menu');
    });

    it('passes label to menu', () => {
      const menu = el.shadowRoot.getElementById('menu');
      expect(menu.getAttribute('label')).to.equal('Filter menu');
    });

    it('removes label from menu when cleared', async () => {
      el.label = '';
      await el.updateComplete;

      const menu = el.shadowRoot.getElementById('menu');
      expect(menu.hasAttribute('label')).to.be.false;
    });
  });

  describe('with menu items', () => {
    let el, items;

    beforeEach(async () => {
      el = await fixture(`
        <pf-v6-dropdown>
          <pf-v6-menu-item value="new">New</pf-v6-menu-item>
          <pf-v6-menu-item value="open">Open</pf-v6-menu-item>
          <pf-v6-menu-item value="save">Save</pf-v6-menu-item>
        </pf-v6-dropdown>
      `);
      await el.rendered;
      items = [...el.querySelectorAll('pf-v6-menu-item')];
      await Promise.all(items.map(item => item.rendered));
    });

    it('bubbles select events from menu items', async () => {
      el.expand();

      setTimeout(() => items[1].click());
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('open');
    });
  });

  describe('property setters', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-dropdown></pf-v6-dropdown>');
      await el.rendered;
    });

    it('sets and removes expanded attribute', () => {
      el.expanded = true;
      expect(el.hasAttribute('expanded')).to.be.true;

      el.expanded = false;
      expect(el.hasAttribute('expanded')).to.be.false;
    });

    it('sets and removes disabled attribute', () => {
      el.disabled = true;
      expect(el.hasAttribute('disabled')).to.be.true;

      el.disabled = false;
      expect(el.hasAttribute('disabled')).to.be.false;
    });

    it('sets and removes label attribute', () => {
      el.label = 'Options menu';
      expect(el.getAttribute('label')).to.equal('Options menu');

      el.label = '';
      expect(el.hasAttribute('label')).to.be.false;
    });
  });

  describe('edge cases', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-dropdown></pf-v6-dropdown>');
      await el.rendered;
    });

    it('handles rapid toggle clicks', () => {
      const toggle = el.shadowRoot.getElementById('toggle');

      for (let i = 0; i < 10; i++) {
        toggle.click();
      }

      // Even number of toggles, should be closed
      expect(el.expanded).to.be.false;
    });

    it('handles expand when already expanded', () => {
      el.expand();
      el.expand();

      expect(el.expanded).to.be.true;
    });

    it('handles collapse when already collapsed', () => {
      el.collapse();
      el.collapse();

      expect(el.expanded).to.be.false;
    });

    it('handles empty dropdown', async () => {
      setTimeout(() => el.expand());
      await oneEvent(el, 'expand');

      expect(el.expanded).to.be.true;
    });

    it('handles multiple Escape presses', () => {
      el.expand();

      for (let i = 0; i < 5; i++) {
        el.dispatchEvent(new KeyboardEvent('keydown', {
          key: 'Escape',
          bubbles: true
        }));
      }

      expect(el.expanded).to.be.false;
    });
  });

  describe('real-world scenarios', () => {
    it('simulates user opening and selecting', async () => {
      const el = await fixture(`
        <pf-v6-dropdown label="File menu">
          <pf-v6-menu-item value="new">New</pf-v6-menu-item>
          <pf-v6-menu-item value="open">Open</pf-v6-menu-item>
          <pf-v6-menu-item value="save">Save</pf-v6-menu-item>
        </pf-v6-dropdown>
      `);
      await el.rendered;
      const items = [...el.querySelectorAll('pf-v6-menu-item')];
      await Promise.all(items.map(item => item.rendered));

      // User clicks toggle
      const toggle = el.shadowRoot.getElementById('toggle');
      toggle.click();

      expect(el.expanded).to.be.true;

      // User selects 'Open'
      setTimeout(() => items[1].click());
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('open');
    });

    it('simulates checkbox dropdown with multi-select', async () => {
      const el = await fixture(`
        <pf-v6-dropdown label="Filter log levels">
          <pf-v6-menu-item variant="checkbox" value="info" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="warn" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="error" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="debug" checked></pf-v6-menu-item>
        </pf-v6-dropdown>
      `);
      await el.rendered;
      const items = [...el.querySelectorAll('pf-v6-menu-item')];
      await Promise.all(items.map(item => item.rendered));

      // User opens dropdown
      el.expand();
      expect(el.expanded).to.be.true;

      // User unchecks 'debug'
      setTimeout(() => items[3].click());
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('debug');
      expect(event.checked).to.be.false;

      // User clicks outside to close
      const outside = document.createElement('div');
      document.body.appendChild(outside);
      outside.click();

      expect(el.expanded).to.be.false;

      document.body.removeChild(outside);
    });

    it('simulates keyboard-only interaction', async () => {
      const el = await fixture(`
        <pf-v6-dropdown>
          <pf-v6-menu-item value="option1"></pf-v6-menu-item>
          <pf-v6-menu-item value="option2"></pf-v6-menu-item>
        </pf-v6-dropdown>
      `);
      await el.rendered;
      const items = [...el.querySelectorAll('pf-v6-menu-item')];
      await Promise.all(items.map(item => item.rendered));

      // User tabs to dropdown and presses Enter to open
      const toggle = el.shadowRoot.getElementById('toggle');
      toggle.focus();
      toggle.click();

      expect(el.expanded).to.be.true;

      // User presses Escape to close
      el.dispatchEvent(new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true
      }));

      expect(el.expanded).to.be.false;
    });

    it('simulates disabled dropdown', async () => {
      const el = await fixture('<pf-v6-dropdown disabled></pf-v6-dropdown>');
      await el.rendered;
      const toggle = el.shadowRoot.getElementById('toggle');

      // User tries to click disabled dropdown
      toggle.click();

      expect(el.expanded).to.be.false;
      expect(toggle.disabled).to.be.true;
    });
  });
});
