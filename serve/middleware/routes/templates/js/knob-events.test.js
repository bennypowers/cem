import { expect } from '@open-wc/testing';
import {
  KnobAttributeChangeEvent,
  KnobPropertyChangeEvent,
  KnobCSSPropertyChangeEvent
} from './knob-events.js';

describe('Knob Events', () => {
  describe('KnobAttributeChangeEvent', () => {
    it('creates event with correct type', () => {
      const event = new KnobAttributeChangeEvent('disabled', 'true');

      expect(event.type).to.equal('knob:attribute-change');
    });

    it('stores attribute name and value', () => {
      const event = new KnobAttributeChangeEvent('variant', 'primary');

      expect(event.name).to.equal('variant');
      expect(event.value).to.equal('primary');
    });

    it('bubbles but is not composed', () => {
      const event = new KnobAttributeChangeEvent('foo', 'bar');

      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });

    it('can be dispatched and caught', (done) => {
      const el = document.createElement('div');
      document.body.appendChild(el);

      el.addEventListener('knob:attribute-change', (e) => {
        expect(e.name).to.equal('color');
        expect(e.value).to.equal('red');
        document.body.removeChild(el);
        done();
      });

      const event = new KnobAttributeChangeEvent('color', 'red');
      el.dispatchEvent(event);
    });

    it('extends Event correctly', () => {
      const event = new KnobAttributeChangeEvent('test', 'value');

      expect(event).to.be.instanceOf(Event);
      expect(event).to.be.instanceOf(KnobAttributeChangeEvent);
    });
  });

  describe('KnobPropertyChangeEvent', () => {
    it('creates event with correct type', () => {
      const event = new KnobPropertyChangeEvent('count', 5);

      expect(event.type).to.equal('knob:property-change');
    });

    it('stores property name and value', () => {
      const event = new KnobPropertyChangeEvent('items', ['a', 'b', 'c']);

      expect(event.name).to.equal('items');
      expect(event.value).to.deep.equal(['a', 'b', 'c']);
    });

    it('supports different value types', () => {
      // String
      let event = new KnobPropertyChangeEvent('label', 'Click me');
      expect(event.value).to.equal('Click me');

      // Number
      event = new KnobPropertyChangeEvent('count', 42);
      expect(event.value).to.equal(42);

      // Boolean
      event = new KnobPropertyChangeEvent('active', true);
      expect(event.value).to.be.true;

      // Array
      event = new KnobPropertyChangeEvent('items', [1, 2, 3]);
      expect(event.value).to.deep.equal([1, 2, 3]);

      // Object
      event = new KnobPropertyChangeEvent('config', { foo: 'bar' });
      expect(event.value).to.deep.equal({ foo: 'bar' });

      // Null
      event = new KnobPropertyChangeEvent('value', null);
      expect(event.value).to.be.null;
    });

    it('bubbles but is not composed', () => {
      const event = new KnobPropertyChangeEvent('prop', 123);

      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });

    it('can be dispatched and caught', (done) => {
      const el = document.createElement('div');
      document.body.appendChild(el);

      el.addEventListener('knob:property-change', (e) => {
        expect(e.name).to.equal('enabled');
        expect(e.value).to.be.true;
        document.body.removeChild(el);
        done();
      });

      const event = new KnobPropertyChangeEvent('enabled', true);
      el.dispatchEvent(event);
    });

    it('extends Event correctly', () => {
      const event = new KnobPropertyChangeEvent('test', 'value');

      expect(event).to.be.instanceOf(Event);
      expect(event).to.be.instanceOf(KnobPropertyChangeEvent);
    });
  });

  describe('KnobCSSPropertyChangeEvent', () => {
    it('creates event with correct type', () => {
      const event = new KnobCSSPropertyChangeEvent('--primary-color', '#ff0000');

      expect(event.type).to.equal('knob:css-property-change');
    });

    it('stores CSS property name and value', () => {
      const event = new KnobCSSPropertyChangeEvent('--spacing-large', '24px');

      expect(event.name).to.equal('--spacing-large');
      expect(event.value).to.equal('24px');
    });

    it('handles various CSS value types', () => {
      // Color
      let event = new KnobCSSPropertyChangeEvent('--color', 'rgba(255, 0, 0, 0.5)');
      expect(event.value).to.equal('rgba(255, 0, 0, 0.5)');

      // Size
      event = new KnobCSSPropertyChangeEvent('--size', '16rem');
      expect(event.value).to.equal('16rem');

      // Custom value
      event = new KnobCSSPropertyChangeEvent('--shadow', '0 2px 4px rgba(0,0,0,0.1)');
      expect(event.value).to.equal('0 2px 4px rgba(0,0,0,0.1)');
    });

    it('bubbles but is not composed', () => {
      const event = new KnobCSSPropertyChangeEvent('--foo', 'bar');

      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });

    it('can be dispatched and caught', (done) => {
      const el = document.createElement('div');
      document.body.appendChild(el);

      el.addEventListener('knob:css-property-change', (e) => {
        expect(e.name).to.equal('--theme-color');
        expect(e.value).to.equal('#0066cc');
        document.body.removeChild(el);
        done();
      });

      const event = new KnobCSSPropertyChangeEvent('--theme-color', '#0066cc');
      el.dispatchEvent(event);
    });

    it('extends Event correctly', () => {
      const event = new KnobCSSPropertyChangeEvent('--test', 'value');

      expect(event).to.be.instanceOf(Event);
      expect(event).to.be.instanceOf(KnobCSSPropertyChangeEvent);
    });
  });

  describe('Event bubbling behavior', () => {
    it('events bubble within light DOM', (done) => {
      // Create parent-child structure in light DOM
      const parent = document.createElement('div');
      const child = document.createElement('div');
      parent.appendChild(child);
      document.body.appendChild(parent);

      // Listen at parent level
      parent.addEventListener('knob:property-change', (e) => {
        expect(e.name).to.equal('bubbled');
        expect(e.value).to.equal('success');
        document.body.removeChild(parent);
        done();
      });

      // Dispatch from child
      const event = new KnobPropertyChangeEvent('bubbled', 'success');
      child.dispatchEvent(event);
    });
  });

  describe('Event bubbling', () => {
    it('events bubble up the DOM tree', (done) => {
      const parent = document.createElement('div');
      const child = document.createElement('div');
      parent.appendChild(child);
      document.body.appendChild(parent);

      parent.addEventListener('knob:property-change', (e) => {
        expect(e.name).to.equal('bubbled');
        expect(e.value).to.equal('success');
        document.body.removeChild(parent);
        done();
      });

      const event = new KnobPropertyChangeEvent('bubbled', 'success');
      child.dispatchEvent(event);
    });
  });

  describe('Multiple event types', () => {
    it('distinguishes between different event types', (done) => {
      const el = document.createElement('div');
      document.body.appendChild(el);

      let attrEventReceived = false;
      let propEventReceived = false;
      let cssEventReceived = false;

      el.addEventListener('knob:attribute-change', () => {
        attrEventReceived = true;
      });

      el.addEventListener('knob:property-change', () => {
        propEventReceived = true;
      });

      el.addEventListener('knob:css-property-change', () => {
        cssEventReceived = true;
        // Check all events received after last one
        expect(attrEventReceived).to.be.true;
        expect(propEventReceived).to.be.true;
        expect(cssEventReceived).to.be.true;
        document.body.removeChild(el);
        done();
      });

      el.dispatchEvent(new KnobAttributeChangeEvent('attr', 'value'));
      el.dispatchEvent(new KnobPropertyChangeEvent('prop', 123));
      el.dispatchEvent(new KnobCSSPropertyChangeEvent('--css', 'value'));
    });
  });
});
