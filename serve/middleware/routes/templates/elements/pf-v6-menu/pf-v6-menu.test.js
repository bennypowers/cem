import { expect, fixture, aTimeout, oneEvent } from '@open-wc/testing';
import { sendKeys } from '@web/test-runner-commands';
import './pf-v6-menu.js';
import '../pf-v6-menu-item/pf-v6-menu-item.js';

describe('pf-v6-menu', () => {
  describe('basic functionality', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-menu></pf-v6-menu>');
      await el.rendered;
    });

    it('is defined as custom element', () => {
      expect(el).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('PfV6Menu');
      expect(el.shadowRoot).to.exist;
    });

    it('renders slot', () => {
      const slot = el.shadowRoot.querySelector('slot');
      expect(slot).to.exist;
    });

    it('observes label attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('label');
    });

    it('handles empty menu gracefully', () => {
      expect(el).to.exist;
    });

    it('focusFirstItem() handles empty menu', () => {
      expect(() => el.focusFirstItem()).to.not.throw();
    });
  });

  describe('label property', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-menu></pf-v6-menu>');
    });

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
  });

  describe('with one item', () => {
    let el, item;

    beforeEach(async () => {
      el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item></pf-v6-menu-item>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      item = el.querySelector('pf-v6-menu-item');
    });

    it('sets tabindex to 0', () => {
      expect(item.getAttribute('tabindex')).to.equal('0');
    });

    it('wraps to self on ArrowDown', async () => {
      item.focus();
      await sendKeys({ press: 'ArrowDown' });

      expect(item.getAttribute('tabindex')).to.equal('0');
    });

    it('wraps to self on ArrowUp', async () => {
      item.focus();
      await sendKeys({ press: 'ArrowUp' });

      expect(item.getAttribute('tabindex')).to.equal('0');
    });
  });

  describe('with three items', () => {
    let el, items;

    beforeEach(async () => {
      el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item></pf-v6-menu-item>
          <pf-v6-menu-item></pf-v6-menu-item>
          <pf-v6-menu-item></pf-v6-menu-item>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      items = [...el.querySelectorAll('pf-v6-menu-item')];
    });

    describe('roving tabindex initialization', () => {
      it('sets first item tabindex to 0, rest to -1', () => {
        expect(items[0].getAttribute('tabindex')).to.equal('0');
        expect(items[1].getAttribute('tabindex')).to.equal('-1');
        expect(items[2].getAttribute('tabindex')).to.equal('-1');
      });
    });

    describe('ArrowDown navigation', () => {
      it('moves focus to next item', async () => {
        items[0].focus();
        await sendKeys({ press: 'ArrowDown' });

        expect(items[1].getAttribute('tabindex')).to.equal('0');
        expect(items[0].getAttribute('tabindex')).to.equal('-1');
      });

      it('wraps from last item to first', async () => {
        items[2].setAttribute('tabindex', '0');
        items[2].focus();
        await sendKeys({ press: 'ArrowDown' });

        expect(items[0].getAttribute('tabindex')).to.equal('0');
        expect(items[2].getAttribute('tabindex')).to.equal('-1');
      });
    });

    describe('ArrowUp navigation', () => {
      it('moves focus to previous item', async () => {
        items[1].setAttribute('tabindex', '0');
        items[1].focus();
        await sendKeys({ press: 'ArrowUp' });

        expect(items[0].getAttribute('tabindex')).to.equal('0');
        expect(items[1].getAttribute('tabindex')).to.equal('-1');
      });

      it('wraps from first item to last', async () => {
        items[0].focus();
        await sendKeys({ press: 'ArrowUp' });

        expect(items[2].getAttribute('tabindex')).to.equal('0');
        expect(items[0].getAttribute('tabindex')).to.equal('-1');
      });
    });

    describe('Home key navigation', () => {
      it('focuses first item', async () => {
        items[2].setAttribute('tabindex', '0');
        items[2].focus();
        await sendKeys({ press: 'Home' });

        expect(items[0].getAttribute('tabindex')).to.equal('0');
        expect(items[2].getAttribute('tabindex')).to.equal('-1');
      });
    });

    describe('End key navigation', () => {
      it('focuses last item', async () => {
        items[0].focus();
        await sendKeys({ press: 'End' });

        expect(items[2].getAttribute('tabindex')).to.equal('0');
        expect(items[0].getAttribute('tabindex')).to.equal('-1');
      });
    });

    describe('focusFirstItem() method', () => {
      it('focuses first menu item', () => {
        el.focusFirstItem();

        expect(document.activeElement).to.equal(items[0]);
        expect(items[0].getAttribute('tabindex')).to.equal('0');
      });
    });

    describe('event bubbling', () => {
      it('allows select events to bubble', async () => {
        items[0].value = 'test';

        setTimeout(() => items[0].click());
        const event = await oneEvent(el, 'select');

        expect(event.value).to.equal('test');
      });
    });
  });

  describe('with one disabled item', () => {
    let el, items;

    beforeEach(async () => {
      el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item></pf-v6-menu-item>
          <pf-v6-menu-item disabled></pf-v6-menu-item>
          <pf-v6-menu-item></pf-v6-menu-item>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      items = [...el.querySelectorAll('pf-v6-menu-item')];
    });

    it('excludes disabled item from initial tabindex', () => {
      expect(items[0].getAttribute('tabindex')).to.equal('0');
      expect(items[1].getAttribute('tabindex')).to.equal('-1');
      expect(items[2].getAttribute('tabindex')).to.equal('-1');
    });

    it('skips disabled items in navigation', async () => {
      items[0].focus();
      await sendKeys({ press: 'ArrowDown' });

      expect(items[2].getAttribute('tabindex')).to.equal('0');
      expect(items[1].getAttribute('tabindex')).to.equal('-1');
    });
  });

  describe('with all disabled items', () => {
    let el, items;

    beforeEach(async () => {
      el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item disabled></pf-v6-menu-item>
          <pf-v6-menu-item disabled></pf-v6-menu-item>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      items = [...el.querySelectorAll('pf-v6-menu-item')];
    });

    it('does not set tabindex="0" on any item', () => {
      expect(items[0].getAttribute('tabindex')).to.not.equal('0');
      expect(items[1].getAttribute('tabindex')).to.not.equal('0');
    });
  });

  describe('with mixed content', () => {
    let el, item, div;

    beforeEach(async () => {
      el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item></pf-v6-menu-item>
          <div>Not a menu item</div>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      item = el.querySelector('pf-v6-menu-item');
      div = el.querySelector('div');
    });

    it('ignores non-menu-item children', () => {
      expect(item.getAttribute('tabindex')).to.equal('0');
      expect(div.hasAttribute('tabindex')).to.be.false;
    });
  });

  describe('lifecycle', () => {
    it('removes event listeners on disconnect', async () => {
      const el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item></pf-v6-menu-item>
        </pf-v6-menu>
      `);
      const item = el.querySelector('pf-v6-menu-item');

      el.remove();

      // Should not throw after disconnect
      expect(() => {
        item.focus();
        item.dispatchEvent(new KeyboardEvent('keydown', {
          key: 'ArrowDown',
          bubbles: true
        }));
      }).to.not.throw();
    });
  });

  describe('complex scenarios', () => {
    it('handles rapid keyboard navigation', async () => {
      const el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item></pf-v6-menu-item>
          <pf-v6-menu-item></pf-v6-menu-item>
          <pf-v6-menu-item></pf-v6-menu-item>
          <pf-v6-menu-item></pf-v6-menu-item>
          <pf-v6-menu-item></pf-v6-menu-item>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      const items = [...el.querySelectorAll('pf-v6-menu-item')];

      items[0].focus();
      for (let i = 0; i < 10; i++) {
        await sendKeys({ press: 'ArrowDown' });
      }

      // Should have wrapped around (10 % 5 = 0, so back to first)
      expect(items[0].getAttribute('tabindex')).to.equal('0');
    });

    it('simulates dropdown menu navigation', async () => {
      const el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item value="new">New File</pf-v6-menu-item>
          <pf-v6-menu-item value="open">Open</pf-v6-menu-item>
          <pf-v6-menu-item value="save">Save</pf-v6-menu-item>
          <pf-v6-menu-item value="close">Close</pf-v6-menu-item>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      const items = [...el.querySelectorAll('pf-v6-menu-item')];

      el.focusFirstItem();
      expect(document.activeElement).to.equal(items[0]);

      await sendKeys({ press: 'ArrowDown' });
      await sendKeys({ press: 'ArrowDown' });

      expect(items[2].getAttribute('tabindex')).to.equal('0');
    });

    it.only('simulates checkbox menu interaction', async () => {
      const el = await fixture(`
        <pf-v6-menu label="Log level filters">
          <pf-v6-menu-item variant="checkbox" value="info" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="warn" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="error" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="debug" checked></pf-v6-menu-item>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      const items = [...el.querySelectorAll('pf-v6-menu-item')];

      el.focusFirstItem();
      await sendKeys({ press: 'ArrowDown' });
      await sendKeys({ press: 'ArrowDown' });
      await sendKeys({ press: 'ArrowDown' });

      expect(items[3].getAttribute('tabindex')).to.equal('0');

      setTimeout(() => items[3].click());
      const event = await oneEvent(el, 'select');

      expect(items[3].checked).to.be.false;
      expect(event.value).to.equal('debug');
      expect(event.checked).to.be.false;
    });

    it('simulates keyboard-only navigation with Home and End', async () => {
      const el = await fixture(`
        <pf-v6-menu>
          <pf-v6-menu-item value="item-0"></pf-v6-menu-item>
          <pf-v6-menu-item value="item-1"></pf-v6-menu-item>
          <pf-v6-menu-item value="item-2"></pf-v6-menu-item>
          <pf-v6-menu-item value="item-3"></pf-v6-menu-item>
        </pf-v6-menu>
      `);
      await aTimeout(100);
      const items = [...el.querySelectorAll('pf-v6-menu-item')];

      items[2].setAttribute('tabindex', '0');
      items[2].focus();
      await sendKeys({ press: 'Home' });

      expect(items[0].getAttribute('tabindex')).to.equal('0');

      await sendKeys({ press: 'End' });

      expect(items[3].getAttribute('tabindex')).to.equal('0');
    });
  });
});
