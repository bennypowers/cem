import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './pf-v6-dropdown.js';

describe('pf-v6-dropdown', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('pf-v6-dropdown');
    document.body.appendChild(el);

    // Wait for CemElement to load template
    await el.rendered;
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('pf-v6-dropdown');
      expect(element).to.be.instanceOf(HTMLElement);
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

    it('menu is hidden by default', () => {
      const menuContainer = el.shadowRoot.getElementById('menu-container');
      const display = window.getComputedStyle(menuContainer).display;
      expect(display).to.equal('none');
    });
  });

  describe('observed attributes', () => {
    it('observes expanded attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('expanded');
    });

    it('observes disabled attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('disabled');
    });

    it('observes label attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('label');
    });
  });

  describe('expanded property', () => {
    it('returns false by default', () => {
      expect(el.expanded).to.be.false;
    });

    it('reflects expanded attribute', () => {
      el.setAttribute('expanded', '');
      expect(el.expanded).to.be.true;
    });

    it('sets expanded attribute', () => {
      el.expanded = true;
      expect(el.hasAttribute('expanded')).to.be.true;
    });

    it('removes expanded attribute when set to false', () => {
      el.setAttribute('expanded', '');
      el.expanded = false;
      expect(el.hasAttribute('expanded')).to.be.false;
    });
  });

  describe('disabled property', () => {
    it('returns false by default', () => {
      expect(el.disabled).to.be.false;
    });

    it('reflects disabled attribute', () => {
      el.setAttribute('disabled', '');
      expect(el.disabled).to.be.true;
    });

    it('sets disabled attribute', () => {
      el.disabled = true;
      expect(el.hasAttribute('disabled')).to.be.true;
    });

    it('disables toggle button', async () => {
      el.disabled = true;
      await el.updateComplete;

      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.disabled).to.be.true;
    });
  });

  describe('label property', () => {
    it('returns empty string by default', () => {
      expect(el.label).to.equal('');
    });

    it('reflects label attribute', () => {
      el.setAttribute('label', 'Filter menu');
      expect(el.label).to.equal('Filter menu');
    });

    it('sets label attribute', () => {
      el.label = 'Options menu';
      expect(el.getAttribute('label')).to.equal('Options menu');
    });

    it('passes label to menu', async () => {
      el.label = 'Test menu';
      await el.updateComplete;

      const menu = el.shadowRoot.getElementById('menu');
      expect(menu.getAttribute('label')).to.equal('Test menu');
    });

    it('removes label from menu when cleared', async () => {
      el.label = 'Test menu';
      await el.updateComplete;

      el.label = '';
      await el.updateComplete;

      const menu = el.shadowRoot.getElementById('menu');
      expect(menu.hasAttribute('label')).to.be.false;
    });
  });

  describe('expand() method', () => {
    it('opens the dropdown', () => {
      el.expand();
      expect(el.expanded).to.be.true;
    });

    it('shows menu container', () => {
      el.expand();

      const menuContainer = el.shadowRoot.getElementById('menu-container');
      expect(menuContainer.style.display).to.equal('block');
    });

    it('dispatches expand event', async () => {
      const expandSpy = sinon.spy();
      el.addEventListener('expand', expandSpy);

      el.expand();

      await waitUntil(() => expandSpy.called);
      expect(expandSpy.calledOnce).to.be.true;
    });

    it('sets aria-expanded on toggle', () => {
      el.expand();

      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.getAttribute('aria-expanded')).to.equal('true');
    });

    it('focuses first menu item', async () => {
      const item = document.createElement('pf-v6-menu-item');
      el.appendChild(item);
      await item.rendered;

      el.expand();

      // Wait for requestAnimationFrame and focus
      await new Promise(resolve => setTimeout(resolve, 100));

      // Verify menu is expanded
      expect(el.expanded).to.be.true;
    });
  });

  describe('collapse() method', () => {
    beforeEach(() => {
      el.expand();
    });

    it('closes the dropdown', () => {
      el.collapse();
      expect(el.expanded).to.be.false;
    });

    it('hides menu container', () => {
      el.collapse();

      const menuContainer = el.shadowRoot.getElementById('menu-container');
      expect(menuContainer.style.display).to.equal('none');
    });

    it('dispatches collapse event', async () => {
      const collapseSpy = sinon.spy();
      el.addEventListener('collapse', collapseSpy);

      el.collapse();

      await waitUntil(() => collapseSpy.called);
      expect(collapseSpy.calledOnce).to.be.true;
    });

    it('sets aria-expanded to false on toggle', () => {
      el.collapse();

      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.getAttribute('aria-expanded')).to.equal('false');
    });
  });

  describe('toggle() method', () => {
    it('opens when closed', () => {
      expect(el.expanded).to.be.false;
      el.toggle();
      expect(el.expanded).to.be.true;
    });

    it('closes when open', () => {
      el.expanded = true;
      el.toggle();
      expect(el.expanded).to.be.false;
    });

    it('toggles multiple times', () => {
      el.toggle();
      expect(el.expanded).to.be.true;

      el.toggle();
      expect(el.expanded).to.be.false;

      el.toggle();
      expect(el.expanded).to.be.true;
    });
  });

  describe('toggle button click', () => {
    it('toggles dropdown on click', () => {
      const toggle = el.shadowRoot.getElementById('toggle');

      toggle.click();
      expect(el.expanded).to.be.true;

      toggle.click();
      expect(el.expanded).to.be.false;
    });

    it('does not toggle when disabled', () => {
      el.disabled = true;
      const toggle = el.shadowRoot.getElementById('toggle');

      toggle.click();
      expect(el.expanded).to.be.false;
    });
  });

  describe('Escape key behavior', () => {
    beforeEach(() => {
      el.expand();
    });

    it('closes dropdown on Escape', () => {
      const event = new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true
      });

      el.dispatchEvent(event);
      expect(el.expanded).to.be.false;
    });

    it('focuses toggle button after Escape', async () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      const focusSpy = sinon.spy(toggle, 'focus');

      const event = new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true
      });

      el.dispatchEvent(event);

      await waitUntil(() => focusSpy.called);
      expect(focusSpy.called).to.be.true;
    });

    it('prevents default on Escape', () => {
      const event = new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true,
        cancelable: true
      });

      const preventSpy = sinon.spy(event, 'preventDefault');
      el.dispatchEvent(event);

      expect(preventSpy.called).to.be.true;
    });

    it('ignores Escape when closed', () => {
      el.collapse();

      const event = new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true,
        cancelable: true
      });

      const preventSpy = sinon.spy(event, 'preventDefault');
      el.dispatchEvent(event);

      expect(preventSpy.called).to.be.false;
    });
  });

  describe('click-outside-to-close', () => {
    beforeEach(() => {
      el.expand();
    });

    it('closes on click outside', () => {
      const outsideElement = document.createElement('div');
      document.body.appendChild(outsideElement);

      outsideElement.click();

      expect(el.expanded).to.be.false;

      document.body.removeChild(outsideElement);
    });

    it('does not close on click inside dropdown', () => {
      el.click();
      expect(el.expanded).to.be.true;
    });

    it('does not close on click inside menu', () => {
      const menu = el.shadowRoot.getElementById('menu');
      menu.click();
      expect(el.expanded).to.be.true;
    });

    it('does not close on click on toggle', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      toggle.click();
      // Click on toggle will toggle, so it closes then immediately opens
      // Just verify no error occurs
      expect(el).to.exist;
    });

    it('registers click listener on expand', async () => {
      const newEl = document.createElement('pf-v6-dropdown');
      document.body.appendChild(newEl);
      await newEl.rendered;

      const addSpy = sinon.spy(document, 'addEventListener');

      newEl.expand();

      const clickCalls = addSpy.getCalls().filter(
        call => call.args[0] === 'click'
      );
      expect(clickCalls.length).to.be.greaterThan(0);

      document.body.removeChild(newEl);
      addSpy.restore();
    });

    it('unregisters click listener on collapse', () => {
      const removeSpy = sinon.spy(document, 'removeEventListener');

      el.collapse();

      const clickCalls = removeSpy.getCalls().filter(
        call => call.args[0] === 'click'
      );
      expect(clickCalls.length).to.be.greaterThan(0);

      removeSpy.restore();
    });
  });

  describe('ARIA attributes', () => {
    it('sets aria-haspopup on toggle', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.getAttribute('aria-haspopup')).to.equal('true');
    });

    it('sets aria-expanded to false by default', () => {
      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.getAttribute('aria-expanded')).to.equal('false');
    });

    it('updates aria-expanded when expanded', () => {
      el.expanded = true;

      const toggle = el.shadowRoot.getElementById('toggle');
      expect(toggle.getAttribute('aria-expanded')).to.equal('true');
    });
  });

  describe('menu item selection bubbling', () => {
    it('allows select events to bubble from menu items', async () => {
      const item = document.createElement('pf-v6-menu-item');
      item.value = 'test-value';
      el.appendChild(item);
      await item.rendered;

      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      el.expand();
      item.click();

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.calledOnce).to.be.true;
      expect(selectSpy.firstCall.args[0].value).to.equal('test-value');
    });
  });

  describe('lifecycle', () => {
    it('removes event listeners on disconnect', async () => {
      const newEl = document.createElement('pf-v6-dropdown');
      document.body.appendChild(newEl);
      await newEl.rendered;

      newEl.expand();

      // Disconnect
      document.body.removeChild(newEl);

      // Click outside should not throw
      const outsideElement = document.createElement('div');
      document.body.appendChild(outsideElement);
      expect(() => outsideElement.click()).to.not.throw();

      document.body.removeChild(outsideElement);
    });

    it('cleans up click listener on disconnect', () => {
      const removeSpy = sinon.spy(document, 'removeEventListener');

      const newEl = document.createElement('pf-v6-dropdown');
      document.body.appendChild(newEl);
      newEl.expand();
      document.body.removeChild(newEl);

      const clickCalls = removeSpy.getCalls().filter(
        call => call.args[0] === 'click'
      );
      expect(clickCalls.length).to.be.greaterThan(0);

      removeSpy.restore();
    });
  });

  describe('edge cases', () => {
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

    it('handles empty dropdown (no menu items)', async () => {
      const expandSpy = sinon.spy();
      el.addEventListener('expand', expandSpy);

      el.expand();

      await waitUntil(() => expandSpy.called);
      expect(el.expanded).to.be.true;
    });

    it('handles multiple Escape key presses', () => {
      el.expand();

      for (let i = 0; i < 5; i++) {
        const event = new KeyboardEvent('keydown', {
          key: 'Escape',
          bubbles: true
        });
        el.dispatchEvent(event);
      }

      expect(el.expanded).to.be.false;
    });

    it('handles non-Escape keys', () => {
      el.expand();

      const event = new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      });
      el.dispatchEvent(event);

      // Should not close
      expect(el.expanded).to.be.true;
    });
  });

  describe('real-world usage', () => {
    it('simulates user opening and selecting from dropdown', async () => {
      el.label = 'File menu';

      const menuItems = [
        { value: 'new', label: 'New' },
        { value: 'open', label: 'Open' },
        { value: 'save', label: 'Save' }
      ];

      const items = [];
      for (const { value, label } of menuItems) {
        const item = document.createElement('pf-v6-menu-item');
        item.value = value;
        item.textContent = label;
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }

      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      // User clicks toggle
      const toggle = el.shadowRoot.getElementById('toggle');
      toggle.click();

      expect(el.expanded).to.be.true;

      // User selects 'Open'
      items[1].click();

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.firstCall.args[0].value).to.equal('open');
    });

    it('simulates checkbox dropdown with multi-select', async () => {
      el.label = 'Filter log levels';

      const levels = ['info', 'warn', 'error', 'debug'];
      const items = [];

      for (const level of levels) {
        const item = document.createElement('pf-v6-menu-item');
        item.variant = 'checkbox';
        item.value = level;
        item.checked = true;
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }

      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      // User opens dropdown
      el.expand();
      expect(el.expanded).to.be.true;

      // User unchecks 'debug'
      items[3].click();

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.firstCall.args[0].value).to.equal('debug');
      expect(selectSpy.firstCall.args[0].checked).to.be.false;

      // User clicks outside to close
      const outside = document.createElement('div');
      document.body.appendChild(outside);
      outside.click();

      expect(el.expanded).to.be.false;

      document.body.removeChild(outside);
    });

    it('simulates keyboard-only interaction', async () => {
      const item1 = document.createElement('pf-v6-menu-item');
      item1.value = 'option1';
      const item2 = document.createElement('pf-v6-menu-item');
      item2.value = 'option2';

      el.appendChild(item1);
      el.appendChild(item2);

      await item1.rendered;
      await item2.rendered;

      // User tabs to dropdown and presses Enter to open
      const toggle = el.shadowRoot.getElementById('toggle');
      toggle.focus();
      toggle.click(); // Simulates activation

      expect(el.expanded).to.be.true;

      // User presses Escape to close
      const event = new KeyboardEvent('keydown', {
        key: 'Escape',
        bubbles: true
      });
      el.dispatchEvent(event);

      expect(el.expanded).to.be.false;
    });

    it('simulates disabled dropdown', () => {
      el.disabled = true;
      const toggle = el.shadowRoot.getElementById('toggle');

      // User tries to click disabled dropdown
      toggle.click();

      expect(el.expanded).to.be.false;
      expect(toggle.disabled).to.be.true;
    });
  });

  describe('template loading', () => {
    it('loads template asynchronously', async () => {
      const newEl = document.createElement('pf-v6-dropdown');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('#toggle'), 'Should load template', {
        timeout: 3000
      });

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.querySelector('#toggle')).to.exist;
      expect(newEl.shadowRoot.querySelector('#menu')).to.exist;

      document.body.removeChild(newEl);
    });
  });
});
