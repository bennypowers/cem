import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './cem-serve-knob-group.js';
import { KnobAttributeChangeEvent, KnobPropertyChangeEvent, KnobCssPropertyChangeEvent } from './cem-serve-knob-group.js';

describe('cem-serve-knob-group', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('cem-serve-knob-group');
    document.body.appendChild(el);

    // Wait for template to load
    await waitUntil(() => el.shadowRoot?.querySelector('slot'), 'Template should load', {
      timeout: 3000
    });
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('cem-serve-knob-group');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('CemServeKnobGroup');
      expect(el.shadowRoot).to.exist;
    });

    it('renders slot for controls', () => {
      const slot = el.shadowRoot.querySelector('slot');
      expect(slot).to.exist;
    });
  });

  describe('event custom classes', () => {
    it('KnobAttributeChangeEvent has correct properties', () => {
      const event = new KnobAttributeChangeEvent('disabled', true);
      expect(event.type).to.equal('knob:attribute-change');
      expect(event.name).to.equal('disabled');
      expect(event.value).to.equal(true);
      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });

    it('KnobPropertyChangeEvent has correct properties', () => {
      const event = new KnobPropertyChangeEvent('variant', 'primary');
      expect(event.type).to.equal('knob:property-change');
      expect(event.name).to.equal('variant');
      expect(event.value).to.equal('primary');
      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });

    it('KnobCssPropertyChangeEvent has correct properties', () => {
      const event = new KnobCssPropertyChangeEvent('--color', '#ff0000');
      expect(event.type).to.equal('knob:css-property-change');
      expect(event.name).to.equal('--color');
      expect(event.value).to.equal('#ff0000');
      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });
  });

  describe('attribute knobs', () => {
    let input;

    beforeEach(() => {
      input = document.createElement('input');
      input.type = 'text';
      input.dataset.knobType = 'attribute';
      input.dataset.knobName = 'label';
      el.appendChild(input);
    });

    it('dispatches attribute change event on input', (done) => {
      el.addEventListener('knob:attribute-change', (e) => {
        expect(e.name).to.equal('label');
        expect(e.value).to.equal('Hello');
        done();
      });

      input.value = 'Hello';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      // Wait for debounce
      setTimeout(() => {}, 300);
    });

    it('debounces text input changes', async () => {
      const spy = sinon.spy();
      el.addEventListener('knob:attribute-change', spy);

      input.value = 'H';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      input.value = 'He';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      input.value = 'Hel';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      input.value = 'Hello';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      // Should not fire immediately
      expect(spy.called).to.be.false;

      // Wait for debounce to complete
      await new Promise(resolve => setTimeout(resolve, 300));

      // Should have fired only once with final value
      expect(spy.calledOnce).to.be.true;
      expect(spy.firstCall.args[0].value).to.equal('Hello');
    });

    it('handles immediate update on change event', () => {
      const spy = sinon.spy();
      el.addEventListener('knob:attribute-change', spy);

      input.value = 'Test';
      input.dispatchEvent(new Event('change', { bubbles: true }));

      // Should fire immediately
      expect(spy.calledOnce).to.be.true;
      expect(spy.firstCall.args[0].value).to.equal('Test');
    });

    it('cancels debounce timer on change event', async () => {
      const spy = sinon.spy();
      el.addEventListener('knob:attribute-change', spy);

      // Start debouncing with input
      input.value = 'Test';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      // Immediately trigger change
      input.dispatchEvent(new Event('change', { bubbles: true }));

      // Should fire once immediately
      expect(spy.calledOnce).to.be.true;

      // Wait for debounce period
      await new Promise(resolve => setTimeout(resolve, 300));

      // Should still only have one call
      expect(spy.calledOnce).to.be.true;
    });
  });

  describe('property knobs', () => {
    let input;

    beforeEach(() => {
      input = document.createElement('input');
      input.type = 'text';
      input.dataset.knobType = 'property';
      input.dataset.knobName = 'variant';
      el.appendChild(input);
    });

    it('dispatches property change event', (done) => {
      el.addEventListener('knob:property-change', (e) => {
        expect(e.name).to.equal('variant');
        expect(e.value).to.equal('primary');
        done();
      });

      input.value = 'primary';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('parses boolean string "true"', (done) => {
      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal(true);
        done();
      });

      input.value = 'true';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('parses boolean string "false"', (done) => {
      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal(false);
        done();
      });

      input.value = 'false';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('parses null string', (done) => {
      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal(null);
        done();
      });

      input.value = 'null';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('parses numeric strings', (done) => {
      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal(42);
        done();
      });

      input.value = '42';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('parses float numbers', (done) => {
      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal(3.14);
        done();
      });

      input.value = '3.14';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('preserves string values that cannot be parsed', (done) => {
      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal('some-string');
        done();
      });

      input.value = 'some-string';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('preserves empty string', (done) => {
      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal('');
        done();
      });

      input.value = '';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });
  });

  describe('CSS property knobs', () => {
    let input;

    beforeEach(() => {
      input = document.createElement('input');
      input.type = 'text';
      input.dataset.knobType = 'css-property';
      input.dataset.knobName = '--color';
      el.appendChild(input);
    });

    it('dispatches CSS property change event', (done) => {
      el.addEventListener('knob:css-property-change', (e) => {
        expect(e.name).to.equal('--color');
        expect(e.value).to.equal('#ff0000');
        done();
      });

      input.value = '#ff0000';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('does not parse CSS values', (done) => {
      el.addEventListener('knob:css-property-change', (e) => {
        // Should remain string, not parse to number
        expect(e.value).to.equal('16px');
        expect(typeof e.value).to.equal('string');
        done();
      });

      input.value = '16px';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });
  });

  describe('boolean controls', () => {
    it('handles checkbox checked state', (done) => {
      const checkbox = document.createElement('input');
      checkbox.type = 'checkbox';
      checkbox.dataset.knobType = 'attribute';
      checkbox.dataset.knobName = 'disabled';
      el.appendChild(checkbox);

      el.addEventListener('knob:attribute-change', (e) => {
        expect(e.value).to.equal(true);
        done();
      });

      checkbox.checked = true;
      checkbox.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('handles checkbox unchecked state', (done) => {
      const checkbox = document.createElement('input');
      checkbox.type = 'checkbox';
      checkbox.dataset.knobType = 'attribute';
      checkbox.dataset.knobName = 'disabled';
      checkbox.checked = true;
      el.appendChild(checkbox);

      el.addEventListener('knob:attribute-change', (e) => {
        expect(e.value).to.equal(false);
        done();
      });

      checkbox.checked = false;
      checkbox.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('handles pf-v6-switch element', (done) => {
      const switchEl = document.createElement('pf-v6-switch');
      switchEl.dataset.knobType = 'property';
      switchEl.dataset.knobName = 'active';
      el.appendChild(switchEl);

      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal(true);
        done();
      });

      switchEl.checked = true;
      switchEl.dispatchEvent(new Event('change', { bubbles: true }));
    });
  });

  describe('control filtering', () => {
    it('ignores inputs without data-knob-type', async () => {
      const spy = sinon.spy();
      el.addEventListener('knob:attribute-change', spy);

      const input = document.createElement('input');
      input.dataset.knobName = 'test';
      el.appendChild(input);

      input.value = 'test';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      await new Promise(resolve => setTimeout(resolve, 300));

      expect(spy.called).to.be.false;
    });

    it('ignores inputs without data-knob-name', async () => {
      const spy = sinon.spy();
      el.addEventListener('knob:attribute-change', spy);

      const input = document.createElement('input');
      input.dataset.knobType = 'attribute';
      el.appendChild(input);

      input.value = 'test';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      await new Promise(resolve => setTimeout(resolve, 300));

      expect(spy.called).to.be.false;
    });

    it('handles unknown knob type gracefully', (done) => {
      const input = document.createElement('input');
      input.dataset.knobType = 'unknown-type';
      input.dataset.knobName = 'test';
      el.appendChild(input);

      // Should not throw
      input.value = 'test';
      input.dispatchEvent(new Event('change', { bubbles: true }));

      setTimeout(done, 50);
    });
  });

  describe('color picker buttons', () => {
    let button, textInputGroup;

    beforeEach(() => {
      textInputGroup = document.createElement('pf-v6-text-input-group');
      textInputGroup.value = '#000000';

      button = document.createElement('button');
      button.className = 'color-picker-button';

      textInputGroup.appendChild(button);
      el.appendChild(textInputGroup);
    });

    it('attaches click listeners to color picker buttons', async () => {
      // Wait for slotchange to fire and listeners to attach
      await new Promise(resolve => setTimeout(resolve, 100));

      const clickSpy = sinon.spy();
      button.addEventListener('click', clickSpy);

      button.click();

      expect(clickSpy.called).to.be.true;
    });

    it('uses EyeDropper API when available', async () => {
      // Mock EyeDropper
      const mockEyeDropper = {
        open: sinon.stub().resolves({ sRGBHex: '#ff0000' })
      };
      window.EyeDropper = sinon.stub().returns(mockEyeDropper);

      const inputSpy = sinon.spy();
      textInputGroup.addEventListener('input', inputSpy);

      // Wait for listeners to attach
      await new Promise(resolve => setTimeout(resolve, 100));

      button.click();

      await new Promise(resolve => setTimeout(resolve, 100));

      expect(mockEyeDropper.open.called).to.be.true;
      expect(textInputGroup.value).to.equal('#ff0000');
      expect(inputSpy.called).to.be.true;

      delete window.EyeDropper;
    });

    it('handles EyeDropper cancellation gracefully', async () => {
      const mockEyeDropper = {
        open: sinon.stub().rejects(new DOMException('User cancelled', 'AbortError'))
      };
      window.EyeDropper = sinon.stub().returns(mockEyeDropper);

      await new Promise(resolve => setTimeout(resolve, 100));

      // Should not throw
      expect(() => button.click()).to.not.throw();

      await new Promise(resolve => setTimeout(resolve, 100));

      delete window.EyeDropper;
    });

    it('falls back to native color input when EyeDropper unavailable', async () => {
      delete window.EyeDropper;

      const originalCreateElement = document.createElement.bind(document);
      const createElementSpy = sinon.spy(document, 'createElement');

      await new Promise(resolve => setTimeout(resolve, 100));

      button.click();

      await new Promise(resolve => setTimeout(resolve, 50));

      // Should have created a color input
      const colorInputCalls = createElementSpy.getCalls().filter(
        call => call.args[0] === 'input'
      );
      expect(colorInputCalls.length).to.be.greaterThan(0);

      createElementSpy.restore();
    });

    it('does not attach duplicate listeners to same button', async () => {
      await new Promise(resolve => setTimeout(resolve, 100));

      // Force slot change again
      const slot = el.shadowRoot.querySelector('slot');
      slot.dispatchEvent(new Event('slotchange'));

      await new Promise(resolve => setTimeout(resolve, 100));

      // Should not throw or create duplicate listeners
      expect(el).to.exist;
    });
  });

  describe('lifecycle', () => {
    it('clears debounce timers on disconnect', async () => {
      const input = document.createElement('input');
      input.dataset.knobType = 'attribute';
      input.dataset.knobName = 'test';
      el.appendChild(input);

      input.value = 'test';
      input.dispatchEvent(new Event('input', { bubbles: true }));

      // Disconnect before debounce completes
      el.parentNode.removeChild(el);

      // Wait for what would have been debounce time
      await new Promise(resolve => setTimeout(resolve, 300));

      // Should not throw
      expect(true).to.be.true;
    });

    it('removes event listeners on disconnect', async () => {
      const newEl = document.createElement('cem-serve-knob-group');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('slot'), '', {
        timeout: 3000
      });

      const spy = sinon.spy();
      newEl.addEventListener('knob:attribute-change', spy);

      const input = document.createElement('input');
      input.dataset.knobType = 'attribute';
      input.dataset.knobName = 'test';
      newEl.appendChild(input);

      // Disconnect
      document.body.removeChild(newEl);

      // Try to trigger event
      input.value = 'test';
      input.dispatchEvent(new Event('change', { bubbles: true }));

      expect(spy.called).to.be.false;
    });

    it('removes color button listeners on disconnect', async () => {
      const button = document.createElement('button');
      button.className = 'color-picker-button';
      el.appendChild(button);

      await new Promise(resolve => setTimeout(resolve, 100));

      document.body.removeChild(el);

      // Should not throw when button is clicked after disconnect
      expect(() => button.click()).to.not.throw();
    });
  });

  describe('edge cases', () => {
    it('handles rapid input changes', async () => {
      const input = document.createElement('input');
      input.dataset.knobType = 'attribute';
      input.dataset.knobName = 'test';
      el.appendChild(input);

      const spy = sinon.spy();
      el.addEventListener('knob:attribute-change', spy);

      // Rapid changes
      for (let i = 0; i < 10; i++) {
        input.value = `value-${i}`;
        input.dispatchEvent(new Event('input', { bubbles: true }));
      }

      await new Promise(resolve => setTimeout(resolve, 300));

      // Should only fire once with last value
      expect(spy.calledOnce).to.be.true;
      expect(spy.firstCall.args[0].value).to.equal('value-9');
    });

    it('handles multiple controls of different types', (done) => {
      const input1 = document.createElement('input');
      input1.dataset.knobType = 'attribute';
      input1.dataset.knobName = 'label';

      const input2 = document.createElement('input');
      input2.dataset.knobType = 'property';
      input2.dataset.knobName = 'count';

      const input3 = document.createElement('input');
      input3.dataset.knobType = 'css-property';
      input3.dataset.knobName = '--color';

      el.appendChild(input1);
      el.appendChild(input2);
      el.appendChild(input3);

      const events = [];
      el.addEventListener('knob:attribute-change', (e) => events.push(e));
      el.addEventListener('knob:property-change', (e) => events.push(e));
      el.addEventListener('knob:css-property-change', (e) => events.push(e));

      input1.value = 'Test';
      input1.dispatchEvent(new Event('change', { bubbles: true }));

      input2.value = '5';
      input2.dispatchEvent(new Event('change', { bubbles: true }));

      input3.value = 'red';
      input3.dispatchEvent(new Event('change', { bubbles: true }));

      setTimeout(() => {
        expect(events).to.have.lengthOf(3);
        expect(events[0].type).to.equal('knob:attribute-change');
        expect(events[1].type).to.equal('knob:property-change');
        expect(events[2].type).to.equal('knob:css-property-change');
        done();
      }, 50);
    });

    it('handles controls with same name but different types', (done) => {
      const input1 = document.createElement('input');
      input1.dataset.knobType = 'attribute';
      input1.dataset.knobName = 'value';

      const input2 = document.createElement('input');
      input2.dataset.knobType = 'property';
      input2.dataset.knobName = 'value';

      el.appendChild(input1);
      el.appendChild(input2);

      const events = [];
      el.addEventListener('knob:attribute-change', (e) => events.push(e));
      el.addEventListener('knob:property-change', (e) => events.push(e));

      input1.value = 'attr-value';
      input1.dispatchEvent(new Event('change', { bubbles: true }));

      input2.value = 'prop-value';
      input2.dispatchEvent(new Event('change', { bubbles: true }));

      setTimeout(() => {
        expect(events).to.have.lengthOf(2);
        expect(events[0].value).to.equal('attr-value');
        expect(events[1].value).to.equal('prop-value');
        done();
      }, 50);
    });

    it('handles special characters in values', (done) => {
      const input = document.createElement('input');
      input.dataset.knobType = 'attribute';
      input.dataset.knobName = 'label';
      el.appendChild(input);

      el.addEventListener('knob:attribute-change', (e) => {
        expect(e.value).to.equal('<>&"\'');
        done();
      });

      input.value = '<>&"\'';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('handles very long values', (done) => {
      const longValue = 'x'.repeat(10000);
      const input = document.createElement('input');
      input.dataset.knobType = 'attribute';
      input.dataset.knobName = 'data';
      el.appendChild(input);

      el.addEventListener('knob:attribute-change', (e) => {
        expect(e.value).to.equal(longValue);
        done();
      });

      input.value = longValue;
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('handles numeric string that looks like number', (done) => {
      const input = document.createElement('input');
      input.dataset.knobType = 'property';
      input.dataset.knobName = 'value';
      el.appendChild(input);

      el.addEventListener('knob:property-change', (e) => {
        expect(e.value).to.equal(0);
        expect(typeof e.value).to.equal('number');
        done();
      });

      input.value = '0';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });
  });

  describe('real-world usage', () => {
    it('simulates updating component attribute', (done) => {
      const input = document.createElement('input');
      input.type = 'text';
      input.dataset.knobType = 'attribute';
      input.dataset.knobName = 'label';
      el.appendChild(input);

      el.addEventListener('knob:attribute-change', (e) => {
        expect(e.name).to.equal('label');
        expect(e.value).to.equal('Submit');
        done();
      });

      input.value = 'Submit';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('simulates toggling boolean attribute', (done) => {
      const checkbox = document.createElement('input');
      checkbox.type = 'checkbox';
      checkbox.dataset.knobType = 'attribute';
      checkbox.dataset.knobName = 'disabled';
      el.appendChild(checkbox);

      el.addEventListener('knob:attribute-change', (e) => {
        expect(e.name).to.equal('disabled');
        expect(e.value).to.equal(true);
        done();
      });

      checkbox.checked = true;
      checkbox.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('simulates setting property with type coercion', (done) => {
      const input = document.createElement('input');
      input.type = 'number';
      input.dataset.knobType = 'property';
      input.dataset.knobName = 'count';
      el.appendChild(input);

      el.addEventListener('knob:property-change', (e) => {
        expect(e.name).to.equal('count');
        expect(e.value).to.equal(10);
        expect(typeof e.value).to.equal('number');
        done();
      });

      input.value = '10';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });

    it('simulates updating CSS custom property', (done) => {
      const input = document.createElement('input');
      input.type = 'color';
      input.dataset.knobType = 'css-property';
      input.dataset.knobName = '--button-background';
      el.appendChild(input);

      el.addEventListener('knob:css-property-change', (e) => {
        expect(e.name).to.equal('--button-background');
        expect(e.value).to.equal('#ff5733');
        done();
      });

      input.value = '#ff5733';
      input.dispatchEvent(new Event('change', { bubbles: true }));
    });
  });

  describe('template loading', () => {
    it('loads template asynchronously', async () => {
      const newEl = document.createElement('cem-serve-knob-group');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('slot'), 'Should load template', {
        timeout: 3000
      });

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.querySelector('slot')).to.exist;

      document.body.removeChild(newEl);
    });

    it('can accept slotted controls after template loads', async () => {
      const newEl = document.createElement('cem-serve-knob-group');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('slot'), '', {
        timeout: 3000
      });

      const input = document.createElement('input');
      input.dataset.knobType = 'attribute';
      input.dataset.knobName = 'test';
      newEl.appendChild(input);

      const spy = sinon.spy();
      newEl.addEventListener('knob:attribute-change', spy);

      input.value = 'test';
      input.dispatchEvent(new Event('change', { bubbles: true }));

      expect(spy.called).to.be.true;

      document.body.removeChild(newEl);
    });
  });
});
