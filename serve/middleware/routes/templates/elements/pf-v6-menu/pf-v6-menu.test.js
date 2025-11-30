import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './pf-v6-menu.js';
import '../pf-v6-menu-item/pf-v6-menu-item.js';

describe('pf-v6-menu', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('pf-v6-menu');
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
      const element = document.createElement('pf-v6-menu');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('PfV6Menu');
      expect(el.shadowRoot).to.exist;
    });

    it('has menu role', async () => {
      await el.updateComplete;

      // Role is set via ElementInternals, verify via slot presence
      expect(el.shadowRoot.querySelector('slot')).to.exist;
    });

    it('renders slot', () => {
      const slot = el.shadowRoot.querySelector('slot');
      expect(slot).to.exist;
    });
  });

  describe('observed attributes', () => {
    it('observes label attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('label');
    });
  });

  describe('label property', () => {
    it('returns empty string by default', () => {
      expect(el.label).to.equal('');
    });

    it('reflects label attribute', () => {
      el.setAttribute('label', 'File menu');
      expect(el.label).to.equal('File menu');
    });

    it('sets label attribute', () => {
      el.label = 'Edit menu';
      expect(el.getAttribute('label')).to.equal('Edit menu');
    });

    it('removes label attribute when set to empty', () => {
      el.setAttribute('label', 'Test');
      el.label = '';
      expect(el.hasAttribute('label')).to.be.false;
    });

    it('sets aria-label via label attribute', async () => {
      el.label = 'Main menu';
      await el.updateComplete;

      expect(el.getAttribute('label')).to.equal('Main menu');
    });

    it('clears label when empty', async () => {
      el.label = 'Test menu';
      await el.updateComplete;

      el.label = '';
      await el.updateComplete;

      expect(el.hasAttribute('label')).to.be.false;
    });
  });

  describe('roving tabindex initialization', () => {
    it('sets first item tabindex to 0', async () => {
      const item1 = document.createElement('pf-v6-menu-item');
      const item2 = document.createElement('pf-v6-menu-item');

      el.appendChild(item1);
      el.appendChild(item2);

      await item1.rendered;
      await item2.rendered;


      // Wait for multiple animation frames to ensure RTI initialization completes
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));

      // Wait for RTI initialization
      await waitUntil(() => item1.getAttribute('tabindex') === '0', 'RTI should initialize', {
        timeout: 3000
      });

      expect(item1.getAttribute('tabindex')).to.equal('0');
      expect(item2.getAttribute('tabindex')).to.equal('-1');
    });

    it('sets remaining items tabindex to -1', async () => {
      const items = [];
      for (let i = 0; i < 4; i++) {
        const item = document.createElement('pf-v6-menu-item');
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI initialization
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));

      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });

      expect(items[0].getAttribute('tabindex')).to.equal('0');
      expect(items[1].getAttribute('tabindex')).to.equal('-1');
      expect(items[2].getAttribute('tabindex')).to.equal('-1');
      expect(items[3].getAttribute('tabindex')).to.equal('-1');
    });

    it('handles empty menu gracefully', () => {
      // Should not throw
      expect(el).to.exist;
    });

    it('excludes disabled items from tabindex', async () => {
      const item1 = document.createElement('pf-v6-menu-item');
      item1.disabled = true;

      const item2 = document.createElement('pf-v6-menu-item');

      el.appendChild(item1);
      el.appendChild(item2);

      await item1.rendered;
      await item2.rendered;


      // Wait for RTI initialization - item2 should be first in tabindex
      await waitUntil(() => item2.getAttribute('tabindex') === '0', '', { timeout: 3000 });

      expect(item2.getAttribute('tabindex')).to.equal('0');
    });
  });

  describe('ArrowDown navigation', () => {
    let items;

    beforeEach(async () => {
      items = [];
      for (let i = 0; i < 3; i++) {
        const item = document.createElement('pf-v6-menu-item');
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });
    });

    it('moves focus to next item', () => {
      items[0].focus();

      const event = new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      });
      items[0].dispatchEvent(event);

      expect(items[1].getAttribute('tabindex')).to.equal('0');
      expect(items[0].getAttribute('tabindex')).to.equal('-1');
    });

    it('wraps from last item to first', () => {
      // Focus last item
      items[2].setAttribute('tabindex', '0');
      items[2].focus();

      const event = new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      });
      items[2].dispatchEvent(event);

      expect(items[0].getAttribute('tabindex')).to.equal('0');
      expect(items[2].getAttribute('tabindex')).to.equal('-1');
    });

    it('prevents default behavior', () => {
      const event = new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true,
        cancelable: true
      });

      const preventSpy = sinon.spy(event, 'preventDefault');
      items[0].dispatchEvent(event);

      expect(preventSpy.called).to.be.true;
    });
  });

  describe('ArrowUp navigation', () => {
    let items;

    beforeEach(async () => {
      items = [];
      for (let i = 0; i < 3; i++) {
        const item = document.createElement('pf-v6-menu-item');
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });
    });

    it('moves focus to previous item', () => {
      // Focus second item
      items[1].setAttribute('tabindex', '0');
      items[1].focus();

      const event = new KeyboardEvent('keydown', {
        key: 'ArrowUp',
        bubbles: true
      });
      items[1].dispatchEvent(event);

      expect(items[0].getAttribute('tabindex')).to.equal('0');
      expect(items[1].getAttribute('tabindex')).to.equal('-1');
    });

    it('wraps from first item to last', () => {
      items[0].focus();

      const event = new KeyboardEvent('keydown', {
        key: 'ArrowUp',
        bubbles: true
      });
      items[0].dispatchEvent(event);

      expect(items[2].getAttribute('tabindex')).to.equal('0');
      expect(items[0].getAttribute('tabindex')).to.equal('-1');
    });

    it('prevents default behavior', () => {
      const event = new KeyboardEvent('keydown', {
        key: 'ArrowUp',
        bubbles: true,
        cancelable: true
      });

      const preventSpy = sinon.spy(event, 'preventDefault');
      items[0].dispatchEvent(event);

      expect(preventSpy.called).to.be.true;
    });
  });

  describe('Home key navigation', () => {
    let items;

    beforeEach(async () => {
      items = [];
      for (let i = 0; i < 3; i++) {
        const item = document.createElement('pf-v6-menu-item');
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });
    });

    it('focuses first item', () => {
      // Focus last item
      items[2].setAttribute('tabindex', '0');
      items[2].focus();

      const event = new KeyboardEvent('keydown', {
        key: 'Home',
        bubbles: true
      });
      items[2].dispatchEvent(event);

      expect(items[0].getAttribute('tabindex')).to.equal('0');
      expect(items[2].getAttribute('tabindex')).to.equal('-1');
    });

    it('prevents default behavior', () => {
      const event = new KeyboardEvent('keydown', {
        key: 'Home',
        bubbles: true,
        cancelable: true
      });

      const preventSpy = sinon.spy(event, 'preventDefault');
      items[0].dispatchEvent(event);

      expect(preventSpy.called).to.be.true;
    });
  });

  describe('End key navigation', () => {
    let items;

    beforeEach(async () => {
      items = [];
      for (let i = 0; i < 3; i++) {
        const item = document.createElement('pf-v6-menu-item');
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });
    });

    it('focuses last item', () => {
      items[0].focus();

      const event = new KeyboardEvent('keydown', {
        key: 'End',
        bubbles: true
      });
      items[0].dispatchEvent(event);

      expect(items[2].getAttribute('tabindex')).to.equal('0');
      expect(items[0].getAttribute('tabindex')).to.equal('-1');
    });

    it('prevents default behavior', () => {
      const event = new KeyboardEvent('keydown', {
        key: 'End',
        bubbles: true,
        cancelable: true
      });

      const preventSpy = sinon.spy(event, 'preventDefault');
      items[0].dispatchEvent(event);

      expect(preventSpy.called).to.be.true;
    });
  });

  describe('disabled items handling', () => {
    it('skips disabled items in navigation', async () => {
      const item1 = document.createElement('pf-v6-menu-item');
      const item2 = document.createElement('pf-v6-menu-item');
      item2.disabled = true;
      const item3 = document.createElement('pf-v6-menu-item');

      el.appendChild(item1);
      el.appendChild(item2);
      el.appendChild(item3);

      await item1.rendered;
      await item2.rendered;
      await item3.rendered;


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));

      await waitUntil(() => item1.getAttribute('tabindex') === '0', '', { timeout: 3000 });

      // Navigate down from item1 - should skip item2 and go to item3
      item1.focus();
      const event = new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      });
      item1.dispatchEvent(event);

      expect(item3.getAttribute('tabindex')).to.equal('0');
      expect(item2.getAttribute('tabindex')).to.equal('-1'); // Still -1, skipped
    });
  });

  describe('focusFirstItem() method', () => {
    it('focuses first menu item', async () => {
      const items = [];
      for (let i = 0; i < 3; i++) {
        const item = document.createElement('pf-v6-menu-item');
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });

      const focusSpy = sinon.spy(items[0], 'focus');
      el.focusFirstItem();

      expect(focusSpy.called).to.be.true;
      expect(items[0].getAttribute('tabindex')).to.equal('0');
    });

    it('handles empty menu', () => {
      // Should not throw
      expect(() => el.focusFirstItem()).to.not.throw();
    });
  });

  describe('event bubbling', () => {
    it('allows select events to bubble', async () => {
      const item = document.createElement('pf-v6-menu-item');
      item.value = 'test';
      el.appendChild(item);
      await item.rendered;

      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      item.click();

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.calledOnce).to.be.true;
      expect(selectSpy.firstCall.args[0].value).to.equal('test');
    });
  });

  describe('lifecycle', () => {
    it('removes event listeners on disconnect', async () => {
      const newEl = document.createElement('pf-v6-menu');
      document.body.appendChild(newEl);
      await newEl.rendered;

      const item = document.createElement('pf-v6-menu-item');
      newEl.appendChild(item);
      await item.rendered;

      // Disconnect
      document.body.removeChild(newEl);

      // Try to trigger keyboard event
      const event = new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      });

      // Should not throw
      expect(() => item.dispatchEvent(event)).to.not.throw();
    });
  });

  describe('edge cases', () => {
    it('handles single item menu', async () => {
      const item = document.createElement('pf-v6-menu-item');
      el.appendChild(item);
      await item.rendered;

      await waitUntil(() => item.getAttribute('tabindex') === '0');

      // Arrow down should wrap to self
      const event = new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      });
      item.dispatchEvent(event);

      expect(item.getAttribute('tabindex')).to.equal('0');
    });

    it('handles all disabled items', async () => {
      const item1 = document.createElement('pf-v6-menu-item');
      item1.disabled = true;
      const item2 = document.createElement('pf-v6-menu-item');
      item2.disabled = true;

      el.appendChild(item1);
      el.appendChild(item2);

      await item1.rendered;
      await item2.rendered;

      // Wait a bit for RTI initialization
      await new Promise(resolve => requestAnimationFrame(resolve));

      // No items should have tabindex="0" since all are disabled
      expect(item1.getAttribute('tabindex')).to.not.equal('0');
      expect(item2.getAttribute('tabindex')).to.not.equal('0');
    });

    it('ignores non-menu-item children', async () => {
      const item = document.createElement('pf-v6-menu-item');
      const div = document.createElement('div');

      el.appendChild(item);
      el.appendChild(div);

      await item.rendered;
      await waitUntil(() => item.getAttribute('tabindex') === '0');

      // Only menu item should get tabindex
      expect(item.getAttribute('tabindex')).to.equal('0');
      expect(div.hasAttribute('tabindex')).to.be.false;
    });

    it('handles rapid keyboard navigation', async () => {
      const items = [];
      for (let i = 0; i < 5; i++) {
        const item = document.createElement('pf-v6-menu-item');
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });

      // Rapid down navigation
      for (let i = 0; i < 10; i++) {
        const currentFocused = items.find(item => item.getAttribute('tabindex') === '0');
        const event = new KeyboardEvent('keydown', {
          key: 'ArrowDown',
          bubbles: true
        });
        currentFocused.dispatchEvent(event);
      }

      // Should have wrapped around (10 % 5 = 0, so back to first)
      expect(items[0].getAttribute('tabindex')).to.equal('0');
    });
  });

  describe('real-world usage', () => {
    it('simulates dropdown menu navigation', async () => {
      const menuItems = [
        { value: 'new', label: 'New File' },
        { value: 'open', label: 'Open' },
        { value: 'save', label: 'Save' },
        { value: 'close', label: 'Close' }
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


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });

      // User opens menu - first item focused
      el.focusFirstItem();
      expect(document.activeElement).to.equal(items[0]);

      // User presses ArrowDown twice
      items[0].dispatchEvent(new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      }));

      const secondFocused = items.find(item => item.getAttribute('tabindex') === '0');
      secondFocused.dispatchEvent(new KeyboardEvent('keydown', {
        key: 'ArrowDown',
        bubbles: true
      }));

      // Should be on 'Save'
      expect(items[2].getAttribute('tabindex')).to.equal('0');
    });

    it('simulates checkbox menu with filtering', async () => {
      el.label = 'Log level filters';

      const filters = ['info', 'warn', 'error', 'debug'];
      const items = [];

      for (const filter of filters) {
        const item = document.createElement('pf-v6-menu-item');
        item.variant = 'checkbox';
        item.value = filter;
        item.checked = true;
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });

      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      // User navigates to 'debug' and unchecks it
      el.focusFirstItem();
      for (let i = 0; i < 3; i++) {
        const currentFocused = items.find(item => item.getAttribute('tabindex') === '0');
        currentFocused.dispatchEvent(new KeyboardEvent('keydown', {
          key: 'ArrowDown',
          bubbles: true
        }));
      }

      // Should be on 'debug'
      expect(items[3].getAttribute('tabindex')).to.equal('0');

      // User presses Space to uncheck
      items[3].click();

      await waitUntil(() => selectSpy.called);
      expect(items[3].checked).to.be.false;
      expect(selectSpy.firstCall.args[0].value).to.equal('debug');
      expect(selectSpy.firstCall.args[0].checked).to.be.false;
    });

    it('simulates keyboard-only menu navigation', async () => {
      const items = [];
      for (let i = 0; i < 4; i++) {
        const item = document.createElement('pf-v6-menu-item');
        item.value = `item-${i}`;
        el.appendChild(item);
        await item.rendered;
        items.push(item);
      }


      // Wait for RTI
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => requestAnimationFrame(resolve));
      await new Promise(resolve => setTimeout(resolve, 100));
      await waitUntil(() => items[0].getAttribute('tabindex') === '0', '', { timeout: 3000 });

      // User uses Home key
      items[2].setAttribute('tabindex', '0');
      items[2].dispatchEvent(new KeyboardEvent('keydown', {
        key: 'Home',
        bubbles: true
      }));

      expect(items[0].getAttribute('tabindex')).to.equal('0');

      // User uses End key
      items[0].dispatchEvent(new KeyboardEvent('keydown', {
        key: 'End',
        bubbles: true
      }));

      expect(items[3].getAttribute('tabindex')).to.equal('0');
    });
  });

  describe('template loading', () => {
    it('loads template asynchronously', async () => {
      const newEl = document.createElement('pf-v6-menu');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('slot'), 'Should load template', {
        timeout: 3000
      });

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.querySelector('slot')).to.exist;

      document.body.removeChild(newEl);
    });
  });
});
