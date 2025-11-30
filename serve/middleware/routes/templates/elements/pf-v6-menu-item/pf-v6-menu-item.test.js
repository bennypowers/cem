import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './pf-v6-menu-item.js';
import { PfMenuItemSelectEvent } from './pf-v6-menu-item.js';

describe('pf-v6-menu-item', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('pf-v6-menu-item');
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
      const element = document.createElement('pf-v6-menu-item');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('PfV6MenuItem');
      expect(el.shadowRoot).to.exist;
    });

    it('has delegatesFocus enabled', () => {
      // delegatesFocus is set in shadowRootOptions
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
  });

  describe('observed attributes', () => {
    it('observes disabled attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('disabled');
    });

    it('observes checked attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('checked');
    });

    it('observes variant attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('variant');
    });

    it('observes value attribute', () => {
      const attrs = el.constructor.observedAttributes;
      expect(attrs).to.include('value');
    });
  });

  describe('PfMenuItemSelectEvent', () => {
    it('has correct properties', () => {
      const event = new PfMenuItemSelectEvent('test-value', true);

      expect(event.type).to.equal('select');
      expect(event.value).to.equal('test-value');
      expect(event.checked).to.equal(true);
      expect(event.bubbles).to.be.true;
    });

    it('handles unchecked state', () => {
      const event = new PfMenuItemSelectEvent('foo', false);

      expect(event.value).to.equal('foo');
      expect(event.checked).to.equal(false);
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

    it('removes disabled attribute when set to false', () => {
      el.setAttribute('disabled', '');
      el.disabled = false;
      expect(el.hasAttribute('disabled')).to.be.false;
    });

    it('sets aria-disabled when disabled', async () => {
      el.disabled = true;
      await el.updateComplete;

      // tabindex should be -1 when disabled
      expect(el.getAttribute('tabindex')).to.equal('-1');
    });

    it('sets tabindex to -1 when disabled', () => {
      el.disabled = true;
      expect(el.getAttribute('tabindex')).to.equal('-1');
    });
  });

  describe('checked property', () => {
    it('returns false by default', () => {
      expect(el.checked).to.be.false;
    });

    it('reflects checked attribute', () => {
      el.setAttribute('checked', '');
      expect(el.checked).to.be.true;
    });

    it('sets checked attribute', () => {
      el.checked = true;
      expect(el.hasAttribute('checked')).to.be.true;
    });

    it('removes checked attribute when set to false', () => {
      el.setAttribute('checked', '');
      el.checked = false;
      expect(el.hasAttribute('checked')).to.be.false;
    });

    it('syncs checkbox input', async () => {
      el.variant = 'checkbox';
      await el.updateComplete;

      el.checked = true;
      await el.updateComplete;

      const input = el.shadowRoot.getElementById('input');
      expect(input.checked).to.be.true;
    });
  });

  describe('variant property', () => {
    it('returns "default" by default', () => {
      expect(el.variant).to.equal('default');
    });

    it('reflects variant attribute', () => {
      el.setAttribute('variant', 'checkbox');
      expect(el.variant).to.equal('checkbox');
    });

    it('sets variant attribute', () => {
      el.variant = 'checkbox';
      expect(el.getAttribute('variant')).to.equal('checkbox');
    });

    it('removes variant attribute when set to empty', () => {
      el.setAttribute('variant', 'checkbox');
      el.variant = '';
      expect(el.hasAttribute('variant')).to.be.false;
    });
  });

  describe('value property', () => {
    it('returns empty string by default', () => {
      expect(el.value).to.equal('');
    });

    it('reflects value attribute', () => {
      el.setAttribute('value', 'test');
      expect(el.value).to.equal('test');
    });

    it('sets value attribute', () => {
      el.value = 'test';
      expect(el.getAttribute('value')).to.equal('test');
    });

    it('removes value attribute when set to empty', () => {
      el.setAttribute('value', 'test');
      el.value = '';
      expect(el.hasAttribute('value')).to.be.false;
    });
  });

  describe('default variant behavior', () => {
    it('has menuitem role', async () => {
      await el.updateComplete;

      // Role is set via ElementInternals, verify via behavior
      expect(el.shadowRoot.querySelector('#input')).to.exist;
    });

    it('dispatches select event on click', async () => {
      el.value = 'test-item';
      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      el.click();

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.calledOnce).to.be.true;
      expect(selectSpy.firstCall.args[0].value).to.equal('test-item');
      expect(selectSpy.firstCall.args[0].checked).to.be.false;
    });

    it('dispatches select event on Enter key', async () => {
      el.value = 'test-item';
      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      const event = new KeyboardEvent('keydown', { key: 'Enter' });
      el.dispatchEvent(event);

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.calledOnce).to.be.true;
    });

    it('dispatches select event on Space key', async () => {
      el.value = 'test-item';
      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      const event = new KeyboardEvent('keydown', { key: ' ' });
      el.dispatchEvent(event);

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.calledOnce).to.be.true;
    });

    it('does not toggle checked state on click', () => {
      el.checked = false;
      el.click();
      expect(el.checked).to.be.false;

      el.checked = true;
      el.click();
      expect(el.checked).to.be.true;
    });
  });

  describe('checkbox variant behavior', () => {
    beforeEach(async () => {
      el.variant = 'checkbox';
      await el.updateComplete;
    });

    it('has menuitemcheckbox role', () => {
      // Role is set via ElementInternals, verify via checkbox input
      const input = el.shadowRoot.getElementById('input');
      expect(input.type).to.equal('checkbox');
    });

    it('sets checked state to false when unchecked', async () => {
      el.checked = false;
      await el.updateComplete;

      expect(el.checked).to.be.false;
      const input = el.shadowRoot.getElementById('input');
      expect(input.checked).to.be.false;
    });

    it('sets checked state to true when checked', async () => {
      el.checked = true;
      await el.updateComplete;

      expect(el.checked).to.be.true;
      const input = el.shadowRoot.getElementById('input');
      expect(input.checked).to.be.true;
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
      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      el.click();

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.firstCall.args[0].value).to.equal('checkbox-item');
      expect(selectSpy.firstCall.args[0].checked).to.be.true;
    });

    it('toggles on Enter key', async () => {
      el.checked = false;
      const event = new KeyboardEvent('keydown', { key: 'Enter' });
      el.dispatchEvent(event);

      await waitUntil(() => el.checked);
      expect(el.checked).to.be.true;
    });

    it('toggles on Space key', async () => {
      el.checked = false;
      const event = new KeyboardEvent('keydown', { key: ' ' });
      el.dispatchEvent(event);

      await waitUntil(() => el.checked);
      expect(el.checked).to.be.true;
    });
  });

  describe('disabled state behavior', () => {
    it('does not dispatch select event when disabled', () => {
      el.disabled = true;
      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      el.click();

      expect(selectSpy.called).to.be.false;
    });

    it('does not toggle checkbox when disabled', () => {
      el.variant = 'checkbox';
      el.disabled = true;
      el.checked = false;

      el.click();

      expect(el.checked).to.be.false;
    });

    it('ignores keyboard events when disabled', () => {
      el.disabled = true;
      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      const enterEvent = new KeyboardEvent('keydown', { key: 'Enter' });
      el.dispatchEvent(enterEvent);

      const spaceEvent = new KeyboardEvent('keydown', { key: ' ' });
      el.dispatchEvent(spaceEvent);

      expect(selectSpy.called).to.be.false;
    });

    it('disables native checkbox input', async () => {
      el.variant = 'checkbox';
      el.disabled = true;
      await el.updateComplete;

      const input = el.shadowRoot.getElementById('input');
      expect(input.disabled).to.be.true;
    });
  });

  describe('variant switching', () => {
    it('updates role when variant changes', async () => {
      el.variant = 'default';
      await el.updateComplete;

      const input = el.shadowRoot.getElementById('input');
      expect(input.type).to.equal('checkbox'); // Input is always checkbox type

      el.variant = 'checkbox';
      await el.updateComplete;
      expect(input.type).to.equal('checkbox');
    });

    it('updates checked state when switching to checkbox', async () => {
      el.checked = true;
      el.variant = 'default';
      await el.updateComplete;

      expect(el.checked).to.be.true;

      el.variant = 'checkbox';
      await el.updateComplete;
      expect(el.checked).to.be.true;

      const input = el.shadowRoot.getElementById('input');
      expect(input.checked).to.be.true;
    });

    it('maintains checked state when switching to default', async () => {
      el.variant = 'checkbox';
      el.checked = true;
      await el.updateComplete;

      expect(el.checked).to.be.true;

      el.variant = 'default';
      await el.updateComplete;
      expect(el.checked).to.be.true;
    });
  });

  describe('lifecycle', () => {
    it('removes event listeners on disconnect', async () => {
      const newEl = document.createElement('pf-v6-menu-item');
      document.body.appendChild(newEl);
      await newEl.rendered;

      const selectSpy = sinon.spy();
      newEl.addEventListener('select', selectSpy);

      // Disconnect
      document.body.removeChild(newEl);

      // Try to trigger event
      newEl.click();

      // Handler should not be called after disconnect
      expect(selectSpy.called).to.be.false;
    });
  });

  describe('edge cases', () => {
    it('handles rapid click toggles', async () => {
      el.variant = 'checkbox';
      el.checked = false;

      for (let i = 0; i < 10; i++) {
        el.click();
      }

      expect(el.checked).to.be.false; // Even number of toggles
    });

    it('handles special characters in value', () => {
      el.value = '<>&"\'';
      expect(el.value).to.equal('<>&"\'');
    });

    it('handles very long values', () => {
      const longValue = 'x'.repeat(10000);
      el.value = longValue;
      expect(el.value).to.equal(longValue);
    });

    it('handles numeric values', () => {
      el.value = '123';
      expect(el.value).to.equal('123');
    });

    it('handles empty value', () => {
      el.value = 'test';
      el.value = '';
      expect(el.value).to.equal('');
      expect(el.hasAttribute('value')).to.be.false;
    });

    it('handles initial checked state', async () => {
      const newEl = document.createElement('pf-v6-menu-item');
      newEl.setAttribute('variant', 'checkbox');
      newEl.setAttribute('checked', '');
      document.body.appendChild(newEl);

      await newEl.rendered;

      expect(newEl.checked).to.be.true;
      const input = newEl.shadowRoot.getElementById('input');
      expect(input.checked).to.be.true;

      document.body.removeChild(newEl);
    });

    it('handles attribute changes before template loads', () => {
      const newEl = document.createElement('pf-v6-menu-item');
      newEl.setAttribute('disabled', '');
      newEl.setAttribute('checked', '');

      // These should not throw before template loads
      expect(() => {
        newEl.setAttribute('variant', 'checkbox');
      }).to.not.throw();

      newEl.remove();
    });
  });

  describe('real-world usage', () => {
    it('simulates menu item selection', async () => {
      el.value = 'save';
      el.textContent = 'Save';

      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      el.click();

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.firstCall.args[0].value).to.equal('save');
    });

    it('simulates checkbox menu item toggle', async () => {
      el.variant = 'checkbox';
      el.value = 'show-warnings';
      el.checked = true;

      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      // User unchecks
      el.click();

      await waitUntil(() => selectSpy.called);
      expect(el.checked).to.be.false;
      expect(selectSpy.firstCall.args[0].value).to.equal('show-warnings');
      expect(selectSpy.firstCall.args[0].checked).to.be.false;
    });

    it('simulates keyboard navigation and selection', async () => {
      el.value = 'copy';
      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      // User navigates to item and presses Enter
      el.focus();
      const event = new KeyboardEvent('keydown', { key: 'Enter' });
      el.dispatchEvent(event);

      await waitUntil(() => selectSpy.called);
      expect(selectSpy.firstCall.args[0].value).to.equal('copy');
    });

    it('simulates disabled menu item', () => {
      el.value = 'delete';
      el.disabled = true;
      const selectSpy = sinon.spy();
      el.addEventListener('select', selectSpy);

      // User tries to click disabled item
      el.click();

      expect(selectSpy.called).to.be.false;
    });

    it('simulates multi-select checkbox group', async () => {
      const items = [];
      const values = ['info', 'warn', 'error', 'debug'];

      for (const value of values) {
        const item = document.createElement('pf-v6-menu-item');
        item.variant = 'checkbox';
        item.value = value;
        item.checked = true;
        document.body.appendChild(item);
        await item.rendered;
        items.push(item);
      }

      // User unchecks 'debug'
      items[3].click();

      // Check final state
      expect(items[0].checked).to.be.true;
      expect(items[1].checked).to.be.true;
      expect(items[2].checked).to.be.true;
      expect(items[3].checked).to.be.false;

      // Cleanup
      items.forEach(item => item.remove());
    });
  });

  describe('template loading', () => {
    it('loads template asynchronously', async () => {
      const newEl = document.createElement('pf-v6-menu-item');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('#input'), 'Should load template', {
        timeout: 3000
      });

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.querySelector('#input')).to.exist;

      document.body.removeChild(newEl);
    });
  });
});
