import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './cem-drawer.js';
import { CemDrawerChangeEvent, CemDrawerResizeEvent } from './cem-drawer.js';

// Helper to flush requestAnimationFrame
const flushRAF = () => new Promise(resolve => requestAnimationFrame(resolve));

describe('cem-drawer', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('cem-drawer');
    document.body.appendChild(el);
    await el.updateComplete;
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('cem-drawer');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('has shadow root', () => {
      expect(el.shadowRoot).to.exist;
    });

    it('starts closed by default', () => {
      expect(el.open).to.be.false;
    });

    it('has toggle button', () => {
      const toggleButton = el.shadowRoot.getElementById('toggle');
      expect(toggleButton).to.exist;
    });

    it('has resize handle', () => {
      const resizeHandle = el.shadowRoot.getElementById('resize-handle');
      expect(resizeHandle).to.exist;
    });

    it('has content element', () => {
      const content = el.shadowRoot.getElementById('content');
      expect(content).to.exist;
    });
  });

  describe('event custom classes', () => {
    it('CemDrawerChangeEvent has correct properties', () => {
      const event = new CemDrawerChangeEvent(true);
      expect(event.type).to.equal('change');
      expect(event.open).to.equal(true);
      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });

    it('CemDrawerResizeEvent has correct properties', () => {
      const event = new CemDrawerResizeEvent(500);
      expect(event.type).to.equal('resize');
      expect(event.height).to.equal(500);
      expect(event.bubbles).to.be.true;
      expect(event.composed).to.be.false;
    });
  });

  describe('open getter/setter', () => {
    it('returns false when not open', () => {
      expect(el.open).to.be.false;
    });

    it('returns true when open attribute is present', async () => {
      el.setAttribute('open', '');
      await el.updateComplete;
      expect(el.open).to.be.true;
    });

    it('sets open attribute when set to true', async () => {
      el.open = true;
      await el.updateComplete;
      expect(el.hasAttribute('open')).to.be.true;
    });

    it('removes open attribute when set to false', async () => {
      el.open = true;
      await el.updateComplete;
      el.open = false;
      await el.updateComplete;
      expect(el.hasAttribute('open')).to.be.false;
    });

    it('coerces truthy values to true', async () => {
      el.open = 'yes';
      await el.updateComplete;
      expect(el.hasAttribute('open')).to.be.true;
    });

    it('coerces falsy values to false', async () => {
      el.open = true;
      await el.updateComplete;
      el.open = '';
      await el.updateComplete;
      expect(el.hasAttribute('open')).to.be.false;
    });
  });

  describe('toggle() method', () => {
    it('opens when closed', () => {
      expect(el.open).to.be.false;
      el.toggle();
      expect(el.open).to.be.true;
    });

    it('closes when open', () => {
      el.open = true;
      el.toggle();
      expect(el.open).to.be.false;
    });

    it('toggles multiple times', () => {
      el.toggle();
      expect(el.open).to.be.true;
      el.toggle();
      expect(el.open).to.be.false;
      el.toggle();
      expect(el.open).to.be.true;
    });
  });

  describe('openDrawer() method', () => {
    it('opens the drawer', () => {
      el.openDrawer();
      expect(el.open).to.be.true;
    });

    it('is idempotent', () => {
      el.openDrawer();
      el.openDrawer();
      expect(el.open).to.be.true;
    });
  });

  describe('close() method', () => {
    it('closes the drawer', () => {
      el.open = true;
      el.close();
      expect(el.open).to.be.false;
    });

    it('is idempotent', () => {
      el.close();
      el.close();
      expect(el.open).to.be.false;
    });
  });

  describe('change events', () => {
    it('dispatches change event when opened via property', (done) => {
      el.addEventListener('change', (e) => {
        expect(e.open).to.be.true;
        done();
      });

      el.open = true;
    });

    it('dispatches change event when closed via property', async () => {
      el.open = true;
      await el.updateComplete;

      const promise = new Promise((resolve) => {
        el.addEventListener('change', (e) => {
          expect(e.open).to.be.false;
          resolve();
        });
      });

      el.open = false;
      await promise;
    });

    it('dispatches change event when toggled', (done) => {
      el.addEventListener('change', (e) => {
        expect(e.open).to.be.true;
        done();
      });

      el.toggle();
    });
  });

  describe('toggle button', () => {
    it('toggles drawer when clicked', () => {
      const toggleButton = el.shadowRoot.getElementById('toggle');

      expect(el.open).to.be.false;
      toggleButton.click();
      expect(el.open).to.be.true;
    });

    it('updates aria-expanded attribute', async () => {
      const toggleButton = el.shadowRoot.getElementById('toggle');

      toggleButton.click();
      await el.updateComplete;
      expect(toggleButton.getAttribute('aria-expanded')).to.equal('true');

      toggleButton.click();
      await el.updateComplete;
      expect(toggleButton.getAttribute('aria-expanded')).to.equal('false');
    });

    it('enables transitions on first click', async () => {
      const toggleButton = el.shadowRoot.getElementById('toggle');

      toggleButton.click();
      await el.updateComplete;

      const content = el.shadowRoot.getElementById('content');
      expect(content.classList.contains('transitions-enabled')).to.be.true;
    });
  });

  describe('height management', () => {
    it('sets default height of 400px when opened', async () => {
      el.open = true;
      await el.updateComplete;

      const content = el.shadowRoot.getElementById('content');
      expect(content.style.height).to.equal('400px');
    });

    it('respects drawer-height attribute when opened', async () => {
      el.setAttribute('drawer-height', '300');
      await el.updateComplete;
      el.open = true;
      await el.updateComplete;

      const content = el.shadowRoot.getElementById('content');
      expect(content.style.height).to.equal('300px');
    });

    it('sets height to 0px when closed', async () => {
      el.open = true;
      await el.updateComplete;
      el.open = false;
      await el.updateComplete;

      const content = el.shadowRoot.getElementById('content');
      expect(content.style.height).to.equal('0px');
    });
  });

  describe('mouse resizing', () => {
    let resizeHandle, content;

    beforeEach(async () => {
      el.open = true;
      await el.updateComplete;
      resizeHandle = el.shadowRoot.getElementById('resize-handle');
      content = el.shadowRoot.getElementById('content');
    });

    it('starts resize on mousedown', () => {
      const event = new MouseEvent('mousedown', { clientY: 500 });
      resizeHandle.dispatchEvent(event);

      // Should have started dragging
      expect(content.classList.contains('transitions-enabled')).to.be.false;
    });

    it('resizes drawer upward on mousemove', async () => {
      // Start dragging
      resizeHandle.dispatchEvent(new MouseEvent('mousedown', { clientY: 500 }));

      // Move mouse up (decrease Y)
      document.dispatchEvent(new MouseEvent('mousemove', { clientY: 400 }));

      // Wait for RAF to apply the resize
      await flushRAF();

      // Height should increase (drawer grows upward)
      const height = parseInt(content.style.height, 10);
      expect(height).to.be.greaterThan(400);
    });

    it('resizes drawer downward on mousemove', async () => {
      content.style.height = '500px';

      // Start dragging
      resizeHandle.dispatchEvent(new MouseEvent('mousedown', { clientY: 300 }));

      // Move mouse down (increase Y)
      document.dispatchEvent(new MouseEvent('mousemove', { clientY: 400 }));

      // Wait for RAF to apply the resize
      await flushRAF();

      // Height should decrease
      const height = parseInt(content.style.height, 10);
      expect(height).to.be.lessThan(500);
    });

    it('enforces minimum height of 100px', () => {
      content.style.height = '100px';

      resizeHandle.dispatchEvent(new MouseEvent('mousedown', { clientY: 300 }));
      document.dispatchEvent(new MouseEvent('mousemove', { clientY: 1000 }));

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(100);
    });

    it('enforces maximum height based on window size', () => {
      resizeHandle.dispatchEvent(new MouseEvent('mousedown', { clientY: 1000 }));
      document.dispatchEvent(new MouseEvent('mousemove', { clientY: 0 }));

      const height = parseInt(content.style.height, 10);
      const maxHeight = window.innerHeight - 56; // headerHeight
      expect(height).to.be.at.most(maxHeight);
    });

    it('stops resize on mouseup', () => {
      resizeHandle.dispatchEvent(new MouseEvent('mousedown', { clientY: 500 }));
      document.dispatchEvent(new MouseEvent('mouseup'));

      // Transitions should be re-enabled
      expect(content.classList.contains('transitions-enabled')).to.be.true;
    });

    it('dispatches resize event after mouseup with debounce', (done) => {
      el.addEventListener('resize', (e) => {
        expect(e.height).to.be.a('number');
        done();
      });

      resizeHandle.dispatchEvent(new MouseEvent('mousedown', { clientY: 500 }));
      document.dispatchEvent(new MouseEvent('mousemove', { clientY: 450 }));
      document.dispatchEvent(new MouseEvent('mouseup'));

      // Wait for debounce (300ms)
    });

    it('prevents text selection during drag', () => {
      const event = new MouseEvent('mousedown', { clientY: 500 });
      const preventDefaultSpy = sinon.spy(event, 'preventDefault');

      resizeHandle.dispatchEvent(event);

      expect(preventDefaultSpy.called).to.be.true;
    });
  });

  describe('keyboard resizing', () => {
    let resizeHandle, content;

    beforeEach(async () => {
      el.open = true;
      await el.updateComplete;
      resizeHandle = el.shadowRoot.getElementById('resize-handle');
      content = el.shadowRoot.getElementById('content');
      content.style.height = '400px';
    });

    it('increases height on ArrowUp', () => {
      const event = new KeyboardEvent('keydown', { key: 'ArrowUp' });
      resizeHandle.dispatchEvent(event);

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(410); // 400 + 10
    });

    it('decreases height on ArrowDown', () => {
      const event = new KeyboardEvent('keydown', { key: 'ArrowDown' });
      resizeHandle.dispatchEvent(event);

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(390); // 400 - 10
    });

    it('uses larger step with Shift key', () => {
      const event = new KeyboardEvent('keydown', { key: 'ArrowUp', shiftKey: true });
      resizeHandle.dispatchEvent(event);

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(450); // 400 + 50
    });

    it('sets to minimum height on Home key', () => {
      const event = new KeyboardEvent('keydown', { key: 'Home' });
      resizeHandle.dispatchEvent(event);

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(100);
    });

    it('sets to maximum height on End key', () => {
      const event = new KeyboardEvent('keydown', { key: 'End' });
      resizeHandle.dispatchEvent(event);

      const height = parseInt(content.style.height, 10);
      const toggleHeight = el.shadowRoot.getElementById('toggle')?.offsetHeight ?? 32;
      const handleHeight = el.shadowRoot.getElementById('resize-handle')?.offsetHeight ?? 4;
      const maxHeight = Math.max(100, window.innerHeight - 56 - toggleHeight - handleHeight);
      expect(height).to.equal(maxHeight);
    });

    it('enforces minimum height with keyboard', () => {
      content.style.height = '100px';

      const event = new KeyboardEvent('keydown', { key: 'ArrowDown' });
      resizeHandle.dispatchEvent(event);

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(100);
    });

    it('enforces maximum height with keyboard', () => {
      const toggleHeight = el.shadowRoot.getElementById('toggle')?.offsetHeight ?? 32;
      const handleHeight = el.shadowRoot.getElementById('resize-handle')?.offsetHeight ?? 4;
      const maxHeight = Math.max(100, window.innerHeight - 56 - toggleHeight - handleHeight);
      content.style.height = `${maxHeight}px`;

      const event = new KeyboardEvent('keydown', { key: 'ArrowUp' });
      resizeHandle.dispatchEvent(event);

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(maxHeight);
    });

    it('prevents default for arrow keys', () => {
      const event = new KeyboardEvent('keydown', { key: 'ArrowUp' });
      const preventDefaultSpy = sinon.spy(event, 'preventDefault');

      resizeHandle.dispatchEvent(event);

      expect(preventDefaultSpy.called).to.be.true;
    });

    it('ignores unhandled keys', () => {
      const initialHeight = parseInt(content.style.height, 10);

      const event = new KeyboardEvent('keydown', { key: 'a' });
      resizeHandle.dispatchEvent(event);

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(initialHeight);
    });

    it('dispatches resize event after keyboard resize with debounce', (done) => {
      el.addEventListener('resize', (e) => {
        expect(e.height).to.be.a('number');
        done();
      });

      const event = new KeyboardEvent('keydown', { key: 'ArrowUp' });
      resizeHandle.dispatchEvent(event);

      // Wait for debounce (300ms)
    });
  });

  describe('ARIA attributes', () => {
    let resizeHandle;

    beforeEach(async () => {
      el.open = true;
      await el.updateComplete;
      resizeHandle = el.shadowRoot.getElementById('resize-handle');
    });

    it('updates aria-valuenow on resize', () => {
      const content = el.shadowRoot.getElementById('content');
      content.style.height = '500px';

      const event = new KeyboardEvent('keydown', { key: 'ArrowUp' });
      resizeHandle.dispatchEvent(event);

      const ariaValueNow = parseInt(resizeHandle.getAttribute('aria-valuenow'), 10);
      expect(ariaValueNow).to.equal(510);
    });

    it('updates aria-valuemax based on window size', () => {
      const event = new KeyboardEvent('keydown', { key: 'ArrowUp' });
      resizeHandle.dispatchEvent(event);

      const toggleHeight = el.shadowRoot.getElementById('toggle')?.offsetHeight ?? 32;
      const handleHeight = resizeHandle?.offsetHeight ?? 4;
      const ariaValueMax = parseInt(resizeHandle.getAttribute('aria-valuemax'), 10);
      const expectedMax = Math.max(100, window.innerHeight - 56 - toggleHeight - handleHeight);
      expect(ariaValueMax).to.equal(expectedMax);
    });
  });

  describe('edge cases', () => {
    it('handles rapid toggle clicks', () => {
      const toggleButton = el.shadowRoot.getElementById('toggle');

      for (let i = 0; i < 10; i++) {
        toggleButton.click();
      }

      // Should end up closed (even number of clicks)
      expect(el.open).to.be.false;
    });

    it('handles opening with non-numeric drawer-height', async () => {
      el.setAttribute('drawer-height', 'invalid');
      el.open = true;
      await el.updateComplete;

      // Should still function even with NaN drawer-height
      expect(el.open).to.be.true;
    });

    it('handles resize without content element', async () => {
      el.open = true;
      await el.updateComplete;
      const resizeHandle = el.shadowRoot.getElementById('resize-handle');

      // Remove content element
      const content = el.shadowRoot.getElementById('content');
      content.parentNode.removeChild(content);

      // Should not throw
      expect(() => {
        const event = new KeyboardEvent('keydown', { key: 'ArrowUp' });
        resizeHandle.dispatchEvent(event);
      }).to.not.throw();
    });

    it('handles mouseup without mousedown', () => {
      // Should not throw
      expect(() => {
        document.dispatchEvent(new MouseEvent('mouseup'));
      }).to.not.throw();
    });

    it('handles mousemove without mousedown', async () => {
      el.open = true;
      await el.updateComplete;
      const content = el.shadowRoot.getElementById('content');
      const initialHeight = content.style.height;

      document.dispatchEvent(new MouseEvent('mousemove', { clientY: 100 }));

      // Height should not change
      expect(content.style.height).to.equal(initialHeight);
    });

    it('fires resize events immediately for keyboard', async () => {
      el.open = true;
      await el.updateComplete;
      const resizeHandle = el.shadowRoot.getElementById('resize-handle');
      const spy = sinon.spy();
      el.addEventListener('resize', spy);

      // Trigger multiple resizes rapidly with keyboard
      for (let i = 0; i < 5; i++) {
        const event = new KeyboardEvent('keydown', { key: 'ArrowUp' });
        resizeHandle.dispatchEvent(event);
      }

      // Keyboard resizing fires events immediately (not debounced)
      expect(spy.callCount).to.equal(5);
    });
  });

  describe('real-world usage', () => {
    it('simulates user opening drawer via toggle button', async () => {
      const toggleButton = el.shadowRoot.getElementById('toggle');

      expect(el.open).to.be.false;
      toggleButton.click();
      await el.updateComplete;
      expect(el.open).to.be.true;

      const content = el.shadowRoot.getElementById('content');
      expect(content.style.height).to.equal('400px');
    });

    it('simulates user resizing drawer with mouse', async () => {
      el.open = true;
      await el.updateComplete;
      const resizeHandle = el.shadowRoot.getElementById('resize-handle');
      const content = el.shadowRoot.getElementById('content');

      // User grabs handle
      resizeHandle.dispatchEvent(new MouseEvent('mousedown', { clientY: 500 }));

      // User drags up
      document.dispatchEvent(new MouseEvent('mousemove', { clientY: 300 }));

      // Wait for RAF to apply the resize
      await flushRAF();

      // User releases
      document.dispatchEvent(new MouseEvent('mouseup'));

      // Drawer should be taller
      const height = parseInt(content.style.height, 10);
      expect(height).to.be.greaterThan(400);
    });

    it('simulates user resizing drawer with keyboard', async () => {
      el.open = true;
      await el.updateComplete;
      const resizeHandle = el.shadowRoot.getElementById('resize-handle');
      const content = el.shadowRoot.getElementById('content');
      content.style.height = '400px';

      // User presses ArrowUp 5 times
      for (let i = 0; i < 5; i++) {
        resizeHandle.dispatchEvent(new KeyboardEvent('keydown', { key: 'ArrowUp' }));
      }

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(450); // 400 + (5 * 10)
    });

    it('simulates user quickly maximizing drawer', async () => {
      el.open = true;
      await el.updateComplete;
      const resizeHandle = el.shadowRoot.getElementById('resize-handle');
      const content = el.shadowRoot.getElementById('content');

      // User presses End key
      resizeHandle.dispatchEvent(new KeyboardEvent('keydown', { key: 'End' }));

      const height = parseInt(content.style.height, 10);
      const toggleHeight = el.shadowRoot.getElementById('toggle')?.offsetHeight ?? 32;
      const handleHeight = resizeHandle?.offsetHeight ?? 4;
      const maxHeight = Math.max(100, window.innerHeight - 56 - toggleHeight - handleHeight);
      expect(height).to.equal(maxHeight);
    });

    it('simulates user minimizing drawer', async () => {
      el.open = true;
      await el.updateComplete;
      const resizeHandle = el.shadowRoot.getElementById('resize-handle');
      const content = el.shadowRoot.getElementById('content');

      // User presses Home key
      resizeHandle.dispatchEvent(new KeyboardEvent('keydown', { key: 'Home' }));

      const height = parseInt(content.style.height, 10);
      expect(height).to.equal(100);
    });

    it('simulates user closing drawer', async () => {
      el.open = true;
      await el.updateComplete;

      const toggleButton = el.shadowRoot.getElementById('toggle');
      toggleButton.click();
      await el.updateComplete;

      expect(el.open).to.be.false;

      const content = el.shadowRoot.getElementById('content');
      expect(content.style.height).to.equal('0px');
    });
  });
});
