import { expect, fixture, aTimeout, oneEvent } from '@open-wc/testing';
import { sendKeys } from '@web/test-runner-commands';
import './pf-v6-menu-item.js';
import { PfMenuItemSelectEvent } from './pf-v6-menu-item.js';

describe('pf-v6-menu-item', () => {
  describe('PfMenuItemSelectEvent', () => {
    it('has correct properties', () => {
      const event = new PfMenuItemSelectEvent('test-value', true);

      expect(event.type).to.equal('select');
      expect(event.value).to.equal('test-value');
      expect(event.checked).to.equal(true);
      expect(event.bubbles).to.be.true;
    });
  });

  describe('basic functionality', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-menu-item></pf-v6-menu-item>');
    });

    it('is defined as custom element', () => {
      expect(el).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('PfV6MenuItem');
      expect(el.shadowRoot).to.exist;
    });

    it('renders input element', () => {
      const input = el.shadowRoot.getElementById('input');
      expect(input).to.exist;
    });

    it('renders slot for content', () => {
      const slot = el.shadowRoot.querySelector('slot');
      expect(slot).to.exist;
    });

    it('observes all attributes', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('disabled');
      expect(attrs).to.include('checked');
      expect(attrs).to.include('variant');
      expect(attrs).to.include('value');
    });
  });

  describe('default variant (menuitem)', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-menu-item></pf-v6-menu-item>');
    });

    it('has default variant', () => {
      expect(el.variant).to.equal('default');
    });

    it('has default values', () => {
      expect(el.disabled).to.be.false;
      expect(el.checked).to.be.false;
      expect(el.value).to.equal('');
    });

    it('dispatches select event on click', async () => {
      el.value = 'test-item';

      setTimeout(() => el.click());
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('test-item');
      expect(event.checked).to.be.false;
    });

    it('dispatches select event on Enter key', async () => {
      el.value = 'test-item';
      el.focus();

      setTimeout(() => sendKeys({ press: 'Enter' }));
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('test-item');
    });

    it('dispatches select event on Space key', async () => {
      el.value = 'test-item';
      el.focus();

      setTimeout(() => sendKeys({ press: 'Space' }));
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('test-item');
    });

    it('does not toggle checked state on click', () => {
      el.checked = false;
      el.click();
      expect(el.checked).to.be.false;
    });
  });

  describe('checkbox variant (menuitemcheckbox)', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-menu-item variant="checkbox"></pf-v6-menu-item>');
      await el.rendered;
    });

    it('has checkbox variant', () => {
      expect(el.variant).to.equal('checkbox');
    });

    it('has checkbox input type', () => {
      const input = el.shadowRoot.getElementById('input');
      expect(input.type).to.equal('checkbox');
    });

    it('syncs checkbox input with checked property', async () => {
      el.checked = true;
      await el.updateComplete;

      const input = el.shadowRoot.getElementById('input');
      expect(input.checked).to.be.true;

      el.checked = false;
      await el.updateComplete;
      expect(input.checked).to.be.false;
    });

    it('toggles checked state on click', () => {
      el.checked = false;
      el.click();
      expect(el.checked).to.be.true;

      el.click();
      expect(el.checked).to.be.false;
    });

    it('dispatches select event with checked state', async () => {
      el.value = 'checkbox-item';
      el.checked = false;

      setTimeout(() => el.click());
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('checkbox-item');
      expect(event.checked).to.be.true;
    });

    it('toggles on Enter key', async () => {
      el.checked = false;
      el.focus();

      await sendKeys({ press: 'Enter' });
      await aTimeout(50);

      expect(el.checked).to.be.true;
    });

    it('toggles on Space key', async () => {
      el.checked = false;
      el.focus();

      await sendKeys({ press: 'Space' });
      await aTimeout(50);

      expect(el.checked).to.be.true;
    });

    it('handles rapid toggles', () => {
      el.checked = false;

      for (let i = 0; i < 10; i++) {
        el.click();
      }

      expect(el.checked).to.be.false; // Even number of toggles
    });
  });

  describe('disabled state', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-menu-item disabled></pf-v6-menu-item>');
      await el.rendered;
    });

    it('reflects disabled attribute', () => {
      expect(el.disabled).to.be.true;
    });

    it('sets tabindex to -1 when disabled', () => {
      expect(el.getAttribute('tabindex')).to.equal('-1');
    });

    it('does not dispatch select event when disabled', () => {
      let eventFired = false;
      el.addEventListener('select', () => eventFired = true);

      el.click();

      expect(eventFired).to.be.false;
    });

    it('does not toggle checkbox when disabled', async () => {
      el.variant = 'checkbox';
      el.checked = false;

      el.click();

      expect(el.checked).to.be.false;
    });

    it('disables native checkbox input', async () => {
      el = await fixture('<pf-v6-menu-item variant="checkbox" disabled></pf-v6-menu-item>');
      await el.rendered;

      const input = el.shadowRoot.getElementById('input');
      expect(input.disabled).to.be.true;
    });
  });

  describe('property setters', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-menu-item></pf-v6-menu-item>');
    });

    it('sets and removes disabled attribute', () => {
      el.disabled = true;
      expect(el.hasAttribute('disabled')).to.be.true;

      el.disabled = false;
      expect(el.hasAttribute('disabled')).to.be.false;
    });

    it('sets and removes checked attribute', () => {
      el.checked = true;
      expect(el.hasAttribute('checked')).to.be.true;

      el.checked = false;
      expect(el.hasAttribute('checked')).to.be.false;
    });

    it('sets and removes variant attribute', () => {
      el.variant = 'checkbox';
      expect(el.getAttribute('variant')).to.equal('checkbox');

      el.variant = '';
      expect(el.hasAttribute('variant')).to.be.false;
    });

    it('sets and removes value attribute', () => {
      el.value = 'test';
      expect(el.getAttribute('value')).to.equal('test');

      el.value = '';
      expect(el.hasAttribute('value')).to.be.false;
    });
  });

  describe('variant switching', () => {
    let el;

    beforeEach(async () => {
      el = await fixture('<pf-v6-menu-item></pf-v6-menu-item>');
    });

    it('maintains checked state when switching variants', async () => {
      el.checked = true;
      el.variant = 'checkbox';
      await el.updateComplete;

      expect(el.checked).to.be.true;
      const input = el.shadowRoot.getElementById('input');
      expect(input.checked).to.be.true;

      el.variant = 'default';
      await el.updateComplete;
      expect(el.checked).to.be.true;
    });
  });

  describe('edge cases', () => {
    it('handles special characters in value', async () => {
      const el = await fixture('<pf-v6-menu-item></pf-v6-menu-item>');
      el.value = '<>&"\'';
      expect(el.value).to.equal('<>&"\'');
    });

    it('handles very long values', async () => {
      const el = await fixture('<pf-v6-menu-item></pf-v6-menu-item>');
      const longValue = 'x'.repeat(10000);
      el.value = longValue;
      expect(el.value).to.equal(longValue);
    });

    it('handles initial checked state', async () => {
      const el = await fixture('<pf-v6-menu-item variant="checkbox" checked></pf-v6-menu-item>');
      await el.rendered;

      expect(el.checked).to.be.true;
      const input = el.shadowRoot.getElementById('input');
      expect(input.checked).to.be.true;
    });
  });

  describe('real-world scenarios', () => {
    it('simulates menu item selection', async () => {
      const el = await fixture('<pf-v6-menu-item value="save">Save</pf-v6-menu-item>');
      await el.rendered;

      setTimeout(() => el.click());
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('save');
    });

    it('simulates checkbox menu item toggle', async () => {
      const el = await fixture('<pf-v6-menu-item variant="checkbox" value="show-warnings" checked></pf-v6-menu-item>');
      await el.rendered;

      setTimeout(() => el.click());
      const event = await oneEvent(el, 'select');

      expect(el.checked).to.be.false;
      expect(event.value).to.equal('show-warnings');
      expect(event.checked).to.be.false;
    });

    it('simulates keyboard navigation and selection', async () => {
      const el = await fixture('<pf-v6-menu-item value="copy"></pf-v6-menu-item>');
      await el.rendered;
      el.focus();

      setTimeout(() => sendKeys({ press: 'Enter' }));
      const event = await oneEvent(el, 'select');

      expect(event.value).to.equal('copy');
    });

    it('simulates disabled menu item', async () => {
      const el = await fixture('<pf-v6-menu-item value="delete" disabled></pf-v6-menu-item>');
      await el.rendered;

      let eventFired = false;
      el.addEventListener('select', () => eventFired = true);
      el.click();

      expect(eventFired).to.be.false;
    });

    it('simulates multi-select checkbox group', async () => {
      const container = await fixture(`
        <div>
          <pf-v6-menu-item variant="checkbox" value="info" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="warn" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="error" checked></pf-v6-menu-item>
          <pf-v6-menu-item variant="checkbox" value="debug" checked></pf-v6-menu-item>
        </div>
      `);
      const items = [...container.querySelectorAll('pf-v6-menu-item')];

      // User unchecks 'debug'
      items[3].click();

      expect(items[0].checked).to.be.true;
      expect(items[1].checked).to.be.true;
      expect(items[2].checked).to.be.true;
      expect(items[3].checked).to.be.false;
    });
  });
});
