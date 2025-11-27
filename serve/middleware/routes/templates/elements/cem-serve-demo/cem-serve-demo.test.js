import { expect } from '@open-wc/testing';
import './cem-serve-demo.js';

describe('cem-serve-demo', () => {
  let el;

  beforeEach(() => {
    el = document.createElement('cem-serve-demo');
    document.body.appendChild(el);
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('cem-serve-demo');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('does not have shadow root by default', () => {
      expect(el.shadowRoot).to.be.null;
    });
  });

  describe('applyKnobChange() - attribute changes', () => {
    let targetElement;

    beforeEach(() => {
      targetElement = document.createElement('test-element');
      el.appendChild(targetElement);
    });

    it('sets string attribute', () => {
      const result = el.applyKnobChange('attribute', 'label', 'Submit', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.getAttribute('label')).to.equal('Submit');
    });

    it('toggles boolean attribute to true', () => {
      const result = el.applyKnobChange('attribute', 'disabled', true, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('disabled')).to.be.true;
    });

    it('toggles boolean attribute to false', () => {
      targetElement.setAttribute('disabled', '');

      const result = el.applyKnobChange('attribute', 'disabled', false, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('disabled')).to.be.false;
    });

    it('removes attribute when value is empty string', () => {
      targetElement.setAttribute('label', 'Test');

      const result = el.applyKnobChange('attribute', 'label', '', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('label')).to.be.false;
    });

    it('removes attribute when value is null', () => {
      targetElement.setAttribute('label', 'Test');

      const result = el.applyKnobChange('attribute', 'label', null, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('label')).to.be.false;
    });

    it('removes attribute when value is undefined', () => {
      targetElement.setAttribute('label', 'Test');

      const result = el.applyKnobChange('attribute', 'label', undefined, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('label')).to.be.false;
    });

    it('returns false when element not found', () => {
      const result = el.applyKnobChange('attribute', 'label', 'Test', 'non-existent', 0);

      expect(result).to.be.false;
    });

    it('targets specific instance by index', () => {
      // Remove the targetElement from beforeEach to test clean indexing
      targetElement.remove();

      const element1 = document.createElement('test-element');
      const element2 = document.createElement('test-element');
      el.appendChild(element1);
      el.appendChild(element2);

      el.applyKnobChange('attribute', 'label', 'First', 'test-element', 0);
      el.applyKnobChange('attribute', 'label', 'Second', 'test-element', 1);

      expect(element1.getAttribute('label')).to.equal('First');
      expect(element2.getAttribute('label')).to.equal('Second');
    });
  });

  describe('applyKnobChange() - property changes', () => {
    let targetElement;

    beforeEach(() => {
      targetElement = document.createElement('test-element');
      el.appendChild(targetElement);
    });

    it('sets string property', () => {
      const result = el.applyKnobChange('property', 'variant', 'primary', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.variant).to.equal('primary');
    });

    it('sets number property', () => {
      const result = el.applyKnobChange('property', 'count', 42, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.count).to.equal(42);
    });

    it('sets boolean property', () => {
      const result = el.applyKnobChange('property', 'active', true, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.active).to.equal(true);
    });

    it('sets null property', () => {
      targetElement.data = 'something';

      const result = el.applyKnobChange('property', 'data', null, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.data).to.equal(null);
    });

    it('sets object property', () => {
      const obj = { foo: 'bar' };
      const result = el.applyKnobChange('property', 'config', obj, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.config).to.equal(obj);
    });

    it('returns false when element not found', () => {
      const result = el.applyKnobChange('property', 'variant', 'primary', 'non-existent', 0);

      expect(result).to.be.false;
    });
  });

  describe('applyKnobChange() - CSS property changes', () => {
    let targetElement;

    beforeEach(() => {
      targetElement = document.createElement('test-element');
      el.appendChild(targetElement);
    });

    it('sets CSS custom property with -- prefix', () => {
      const result = el.applyKnobChange('css-property', '--color', 'red', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--color')).to.equal('red');
    });

    it('sets CSS custom property without -- prefix', () => {
      const result = el.applyKnobChange('css-property', 'color', 'blue', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--color')).to.equal('blue');
    });

    it('removes CSS property when value is empty string', () => {
      targetElement.style.setProperty('--color', 'red');

      const result = el.applyKnobChange('css-property', '--color', '', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--color')).to.equal('');
    });

    it('removes CSS property when value is null', () => {
      targetElement.style.setProperty('--color', 'red');

      const result = el.applyKnobChange('css-property', '--color', null, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--color')).to.equal('');
    });

    it('removes CSS property when value is undefined', () => {
      targetElement.style.setProperty('--color', 'red');

      const result = el.applyKnobChange('css-property', '--color', undefined, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--color')).to.equal('');
    });

    it('returns false when element not found', () => {
      const result = el.applyKnobChange('css-property', '--color', 'red', 'non-existent', 0);

      expect(result).to.be.false;
    });
  });

  describe('applyKnobChange() - unknown type', () => {
    it('returns false for unknown type', () => {
      const targetElement = document.createElement('test-element');
      el.appendChild(targetElement);

      const result = el.applyKnobChange('unknown-type', 'name', 'value', 'test-element', 0);

      expect(result).to.be.false;
    });
  });

  describe('setDemoAttribute() method', () => {
    let targetElement;

    beforeEach(() => {
      targetElement = document.createElement('test-element');
      targetElement.id = 'target';
      el.appendChild(targetElement);
    });

    it('sets string attribute using selector', () => {
      const result = el.setDemoAttribute('#target', 'label', 'Test');

      expect(result).to.be.true;
      expect(targetElement.getAttribute('label')).to.equal('Test');
    });

    it('sets boolean attribute to true', () => {
      const result = el.setDemoAttribute('#target', 'disabled', true);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('disabled')).to.be.true;
      expect(targetElement.getAttribute('disabled')).to.equal('');
    });

    it('removes boolean attribute when false', () => {
      targetElement.setAttribute('disabled', '');

      const result = el.setDemoAttribute('#target', 'disabled', false);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('disabled')).to.be.false;
    });

    it('removes attribute when value is empty string', () => {
      targetElement.setAttribute('label', 'Test');

      const result = el.setDemoAttribute('#target', 'label', '');

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('label')).to.be.false;
    });

    it('removes attribute when value is null', () => {
      targetElement.setAttribute('label', 'Test');

      const result = el.setDemoAttribute('#target', 'label', null);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('label')).to.be.false;
    });

    it('removes attribute when value is undefined', () => {
      targetElement.setAttribute('label', 'Test');

      const result = el.setDemoAttribute('#target', 'label', undefined);

      expect(result).to.be.true;
      expect(targetElement.hasAttribute('label')).to.be.false;
    });

    it('returns false when selector does not match', () => {
      const result = el.setDemoAttribute('#non-existent', 'label', 'Test');

      expect(result).to.be.false;
    });
  });

  describe('setDemoProperty() method', () => {
    let targetElement;

    beforeEach(() => {
      targetElement = document.createElement('test-element');
      targetElement.id = 'target';
      el.appendChild(targetElement);
    });

    it('sets property using selector', () => {
      const result = el.setDemoProperty('#target', 'variant', 'primary');

      expect(result).to.be.true;
      expect(targetElement.variant).to.equal('primary');
    });

    it('sets number property', () => {
      const result = el.setDemoProperty('#target', 'count', 10);

      expect(result).to.be.true;
      expect(targetElement.count).to.equal(10);
    });

    it('sets boolean property', () => {
      const result = el.setDemoProperty('#target', 'active', false);

      expect(result).to.be.true;
      expect(targetElement.active).to.equal(false);
    });

    it('returns false when selector does not match', () => {
      const result = el.setDemoProperty('#non-existent', 'variant', 'primary');

      expect(result).to.be.false;
    });
  });

  describe('setDemoCssCustomProperty() method', () => {
    let targetElement;

    beforeEach(() => {
      targetElement = document.createElement('test-element');
      targetElement.id = 'target';
      el.appendChild(targetElement);
    });

    it('sets CSS property with -- prefix using selector', () => {
      const result = el.setDemoCssCustomProperty('#target', '--color', 'red');

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--color')).to.equal('red');
    });

    it('sets CSS property without -- prefix', () => {
      const result = el.setDemoCssCustomProperty('#target', 'background', 'blue');

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--background')).to.equal('blue');
    });

    it('returns false when selector does not match', () => {
      const result = el.setDemoCssCustomProperty('#non-existent', '--color', 'red');

      expect(result).to.be.false;
    });
  });

  describe('shadow DOM support', () => {
    beforeEach(() => {
      // Create element with shadow DOM
      el.attachShadow({ mode: 'open' });
    });

    it('applies changes to elements in shadow DOM', () => {
      const targetElement = document.createElement('test-element');
      el.shadowRoot.appendChild(targetElement);

      const result = el.applyKnobChange('attribute', 'label', 'Test', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.getAttribute('label')).to.equal('Test');
    });

    it('setDemoAttribute works with shadow DOM', () => {
      const targetElement = document.createElement('test-element');
      targetElement.id = 'target';
      el.shadowRoot.appendChild(targetElement);

      const result = el.setDemoAttribute('#target', 'label', 'Test');

      expect(result).to.be.true;
      expect(targetElement.getAttribute('label')).to.equal('Test');
    });

    it('setDemoProperty works with shadow DOM', () => {
      const targetElement = document.createElement('test-element');
      targetElement.id = 'target';
      el.shadowRoot.appendChild(targetElement);

      const result = el.setDemoProperty('#target', 'variant', 'primary');

      expect(result).to.be.true;
      expect(targetElement.variant).to.equal('primary');
    });

    it('setDemoCssCustomProperty works with shadow DOM', () => {
      const targetElement = document.createElement('test-element');
      targetElement.id = 'target';
      el.shadowRoot.appendChild(targetElement);

      const result = el.setDemoCssCustomProperty('#target', '--color', 'red');

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--color')).to.equal('red');
    });
  });

  describe('edge cases', () => {
    it('handles very long attribute values', () => {
      const targetElement = document.createElement('test-element');
      el.appendChild(targetElement);

      const longValue = 'x'.repeat(10000);
      const result = el.applyKnobChange('attribute', 'data', longValue, 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.getAttribute('data')).to.equal(longValue);
    });

    it('handles special characters in attribute values', () => {
      const targetElement = document.createElement('test-element');
      el.appendChild(targetElement);

      const result = el.applyKnobChange('attribute', 'label', '<>&"\'', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.getAttribute('label')).to.equal('<>&"\'');
    });

    it('handles numeric attribute names', function() {
      // Skip on Firefox - numeric attribute names are invalid per HTML spec
      // Firefox correctly throws InvalidCharacterError, while Chrome allows it
      if (navigator.userAgent.toLowerCase().includes('firefox')) {
        this.skip();
        return;
      }

      const targetElement = document.createElement('test-element');
      el.appendChild(targetElement);

      const result = el.applyKnobChange('attribute', '123', 'value', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.getAttribute('123')).to.equal('value');
    });

    it('handles multiple elements with same tag name', () => {
      const elements = [];
      for (let i = 0; i < 5; i++) {
        const element = document.createElement('test-element');
        elements.push(element);
        el.appendChild(element);
      }

      // Target specific instances
      el.applyKnobChange('attribute', 'index', '0', 'test-element', 0);
      el.applyKnobChange('attribute', 'index', '2', 'test-element', 2);
      el.applyKnobChange('attribute', 'index', '4', 'test-element', 4);

      expect(elements[0].getAttribute('index')).to.equal('0');
      expect(elements[1].hasAttribute('index')).to.be.false;
      expect(elements[2].getAttribute('index')).to.equal('2');
      expect(elements[3].hasAttribute('index')).to.be.false;
      expect(elements[4].getAttribute('index')).to.equal('4');
    });

    it('handles instance index beyond available elements', () => {
      const targetElement = document.createElement('test-element');
      el.appendChild(targetElement);

      const result = el.applyKnobChange('attribute', 'label', 'Test', 'test-element', 5);

      expect(result).to.be.false;
    });

    it('handles negative instance index', () => {
      const targetElement = document.createElement('test-element');
      el.appendChild(targetElement);

      const result = el.applyKnobChange('attribute', 'label', 'Test', 'test-element', -1);

      expect(result).to.be.false;
    });

    it('handles complex CSS property values', () => {
      const targetElement = document.createElement('test-element');
      el.appendChild(targetElement);

      const result = el.applyKnobChange('css-property', '--shadow', '0 2px 4px rgba(0,0,0,0.1)', 'test-element', 0);

      expect(result).to.be.true;
      expect(targetElement.style.getPropertyValue('--shadow')).to.equal('0 2px 4px rgba(0,0,0,0.1)');
    });
  });

  describe('real-world usage', () => {
    it('simulates knobs updating button label', () => {
      const button = document.createElement('button');
      el.appendChild(button);

      el.applyKnobChange('attribute', 'aria-label', 'Click me', 'button', 0);

      expect(button.getAttribute('aria-label')).to.equal('Click me');
    });

    it('simulates knobs toggling disabled state', () => {
      const button = document.createElement('button');
      el.appendChild(button);

      el.applyKnobChange('attribute', 'disabled', true, 'button', 0);
      expect(button.hasAttribute('disabled')).to.be.true;

      el.applyKnobChange('attribute', 'disabled', false, 'button', 0);
      expect(button.hasAttribute('disabled')).to.be.false;
    });

    it('simulates knobs updating component property', () => {
      const customElement = document.createElement('custom-element');
      el.appendChild(customElement);

      el.applyKnobChange('property', 'variant', 'primary', 'custom-element', 0);
      expect(customElement.variant).to.equal('primary');

      el.applyKnobChange('property', 'variant', 'secondary', 'custom-element', 0);
      expect(customElement.variant).to.equal('secondary');
    });

    it('simulates knobs updating CSS custom properties', () => {
      const customElement = document.createElement('custom-element');
      el.appendChild(customElement);

      el.applyKnobChange('css-property', '--button-color', '#ff0000', 'custom-element', 0);
      expect(customElement.style.getPropertyValue('--button-color')).to.equal('#ff0000');

      el.applyKnobChange('css-property', '--button-color', '#00ff00', 'custom-element', 0);
      expect(customElement.style.getPropertyValue('--button-color')).to.equal('#00ff00');
    });

    it('simulates multi-instance demo', () => {
      // Create 3 instances of same component
      const instances = [];
      for (let i = 0; i < 3; i++) {
        const instance = document.createElement('my-component');
        instances.push(instance);
        el.appendChild(instance);
      }

      // Update each instance independently
      el.applyKnobChange('attribute', 'label', 'First', 'my-component', 0);
      el.applyKnobChange('attribute', 'label', 'Second', 'my-component', 1);
      el.applyKnobChange('attribute', 'label', 'Third', 'my-component', 2);

      expect(instances[0].getAttribute('label')).to.equal('First');
      expect(instances[1].getAttribute('label')).to.equal('Second');
      expect(instances[2].getAttribute('label')).to.equal('Third');
    });

    it('simulates updating complex component state', () => {
      const component = document.createElement('complex-component');
      el.appendChild(component);

      // Set multiple properties and attributes
      el.applyKnobChange('attribute', 'variant', 'primary', 'complex-component', 0);
      el.applyKnobChange('property', 'count', 10, 'complex-component', 0);
      el.applyKnobChange('css-property', '--spacing', '16px', 'complex-component', 0);

      expect(component.getAttribute('variant')).to.equal('primary');
      expect(component.count).to.equal(10);
      expect(component.style.getPropertyValue('--spacing')).to.equal('16px');
    });
  });
});
