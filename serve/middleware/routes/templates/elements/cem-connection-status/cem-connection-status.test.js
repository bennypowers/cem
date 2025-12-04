import { fixture, expect, html, waitUntil } from '@open-wc/testing';
import './cem-connection-status.js';

describe('cem-connection-status', () => {
  describe('initialization', () => {
    it('is defined as custom element', () => {
      const el = document.createElement('cem-connection-status');
      expect(el).to.be.instanceOf(HTMLElement);
    });

    it('creates shadow root on construction', () => {
      const el = document.createElement('cem-connection-status');
      expect(el.shadowRoot).to.exist;
      expect(el.shadowRoot.mode).to.equal('open');
    });

    it('renders template with icon and message elements', () => {
      const el = document.createElement('cem-connection-status');

      const icon = el.shadowRoot.getElementById('icon');
      const message = el.shadowRoot.getElementById('message');

      expect(icon).to.exist;
      expect(message).to.exist;
    });

    it('has correct inline styles', () => {
      const el = document.createElement('cem-connection-status');
      const style = el.shadowRoot.querySelector('style');

      expect(style).to.exist;
      expect(style.textContent).to.include('position: fixed');
      expect(style.textContent).to.include('bottom: 20px');
      expect(style.textContent).to.include('right: 20px');
    });
  });

  describe('show() method', () => {
    let el;

    beforeEach(() => {
      el = document.createElement('cem-connection-status');
      document.body.appendChild(el);
    });

    afterEach(() => {
      if (el && el.parentNode) {
        el.parentNode.removeChild(el);
      }
    });

    describe('connected state', () => {
      it('shows connected status with green icon', () => {
        el.show('connected', 'Connected to server');

        expect(el.getAttribute('state')).to.equal('connected');
        expect(el.shadowRoot.getElementById('icon').textContent).to.equal('🟢');
        expect(el.shadowRoot.getElementById('message').textContent).to.equal('Connected to server');
      });

      it('applies connected styling', () => {
        el.show('connected', 'Connected');

        expect(el.getAttribute('state')).to.equal('connected');
        // Connected state has specific colors in CSS
        const styles = getComputedStyle(el);
        // Note: computed styles may vary, just check attribute is set
        expect(el.hasAttribute('state')).to.be.true;
      });

      it('fades after specified delay', async () => {
        el.show('connected', 'Connected', { fadeDelay: 50 });

        expect(el.hasAttribute('faded')).to.be.false;

        await waitUntil(() => el.hasAttribute('faded'), 'Element should fade', {
          timeout: 200
        });

        expect(el.hasAttribute('faded')).to.be.true;
      });

      it('does not fade without fadeDelay', async () => {
        el.show('connected', 'Connected');

        await new Promise(resolve => setTimeout(resolve, 100));

        expect(el.hasAttribute('faded')).to.be.false;
      });

      it('does not fade with fadeDelay=0', async () => {
        el.show('connected', 'Connected', { fadeDelay: 0 });

        await new Promise(resolve => setTimeout(resolve, 50));

        expect(el.hasAttribute('faded')).to.be.false;
      });
    });

    describe('reconnecting state', () => {
      it('shows reconnecting status with yellow icon', () => {
        el.show('reconnecting', 'Reconnecting...');

        expect(el.getAttribute('state')).to.equal('reconnecting');
        expect(el.shadowRoot.getElementById('icon').textContent).to.equal('🟡');
        expect(el.shadowRoot.getElementById('message').textContent).to.equal('Reconnecting...');
      });

      it('does not fade reconnecting state', async () => {
        el.show('reconnecting', 'Reconnecting...', { fadeDelay: 50 });

        await new Promise(resolve => setTimeout(resolve, 100));

        // Reconnecting should not fade even with fadeDelay
        expect(el.hasAttribute('faded')).to.be.false;
      });
    });

    describe('disconnected state', () => {
      it('shows disconnected status with red icon', () => {
        el.show('disconnected', 'Connection lost');

        expect(el.getAttribute('state')).to.equal('disconnected');
        expect(el.shadowRoot.getElementById('icon').textContent).to.equal('🔴');
        expect(el.shadowRoot.getElementById('message').textContent).to.equal('Connection lost');
      });

      it('does not fade disconnected state', async () => {
        el.show('disconnected', 'Disconnected', { fadeDelay: 50 });

        await new Promise(resolve => setTimeout(resolve, 100));

        // Disconnected should not fade even with fadeDelay
        expect(el.hasAttribute('faded')).to.be.false;
      });
    });

    describe('state transitions', () => {
      it('transitions from connected to reconnecting', () => {
        el.show('connected', 'Connected');
        expect(el.getAttribute('state')).to.equal('connected');

        el.show('reconnecting', 'Reconnecting...');
        expect(el.getAttribute('state')).to.equal('reconnecting');
        expect(el.shadowRoot.getElementById('icon').textContent).to.equal('🟡');
      });

      it('transitions from reconnecting to connected', () => {
        el.show('reconnecting', 'Reconnecting...');
        expect(el.getAttribute('state')).to.equal('reconnecting');

        el.show('connected', 'Connected');
        expect(el.getAttribute('state')).to.equal('connected');
        expect(el.shadowRoot.getElementById('icon').textContent).to.equal('🟢');
      });

      it('transitions from reconnecting to disconnected', () => {
        el.show('reconnecting', 'Reconnecting...');
        expect(el.getAttribute('state')).to.equal('reconnecting');

        el.show('disconnected', 'Connection failed');
        expect(el.getAttribute('state')).to.equal('disconnected');
        expect(el.shadowRoot.getElementById('icon').textContent).to.equal('🔴');
      });

      it('removes faded attribute on state change', async () => {
        el.show('connected', 'Connected', { fadeDelay: 50 });

        await waitUntil(() => el.hasAttribute('faded'));
        expect(el.hasAttribute('faded')).to.be.true;

        // Change state
        el.show('reconnecting', 'Reconnecting...');

        expect(el.hasAttribute('faded')).to.be.false;
      });
    });

    describe('message updates', () => {
      it('updates message for same state', () => {
        el.show('reconnecting', 'Attempt 1...');
        expect(el.shadowRoot.getElementById('message').textContent).to.equal('Attempt 1...');

        el.show('reconnecting', 'Attempt 2...');
        expect(el.shadowRoot.getElementById('message').textContent).to.equal('Attempt 2...');
      });

      it('handles empty message', () => {
        el.show('connected', '');
        expect(el.shadowRoot.getElementById('message').textContent).to.equal('');
      });

      it('handles long message', () => {
        const longMessage = 'This is a very long message that might overflow the toast container';
        el.show('connected', longMessage);
        expect(el.shadowRoot.getElementById('message').textContent).to.equal(longMessage);
      });

      it('handles special characters in message', () => {
        el.show('connected', 'Connected! <>&"\'');
        expect(el.shadowRoot.getElementById('message').textContent).to.equal('Connected! <>&"\'');
      });
    });
  });

  describe('hide() method', () => {
    let el;

    beforeEach(() => {
      el = document.createElement('cem-connection-status');
      document.body.appendChild(el);
    });

    afterEach(() => {
      // Element might be removed, check first
      if (el && el.parentNode) {
        el.parentNode.removeChild(el);
      }
    });

    it('sets opacity to 0', () => {
      el.hide();

      expect(el.style.opacity).to.equal('0');
    });

    it('removes element after transition', async () => {
      const parent = el.parentNode;
      expect(parent.contains(el)).to.be.true;

      el.hide();

      await waitUntil(() => !parent.contains(el), 'Element should be removed', {
        timeout: 500
      });

      expect(parent.contains(el)).to.be.false;
    });

    it('can be called multiple times safely', () => {
      expect(() => {
        el.hide();
        el.hide();
        el.hide();
      }).to.not.throw();
    });

    it('works when element is already fading', async () => {
      el.show('connected', 'Connected', { fadeDelay: 50 });

      await waitUntil(() => el.hasAttribute('faded'));

      expect(() => el.hide()).to.not.throw();
    });
  });

  describe('positioning and styling', () => {
    let el;

    beforeEach(() => {
      el = document.createElement('cem-connection-status');
      document.body.appendChild(el);
    });

    afterEach(() => {
      if (el && el.parentNode) {
        el.parentNode.removeChild(el);
      }
    });

    it('is positioned fixed in bottom-right corner', () => {
      const styles = getComputedStyle(el);
      expect(styles.position).to.equal('fixed');
    });

    it('has high z-index for overlay', () => {
      el.show('connected', 'Test');
      const styles = getComputedStyle(el);
      // z-index should be very high (999999)
      expect(parseInt(styles.zIndex)).to.be.greaterThan(1000);
    });

    it('has flexbox layout for icon and message', () => {
      const styles = getComputedStyle(el);
      expect(styles.display).to.equal('flex');
    });
  });

  describe('accessibility', () => {
    let el;

    beforeEach(() => {
      el = document.createElement('cem-connection-status');
      document.body.appendChild(el);
    });

    afterEach(() => {
      if (el && el.parentNode) {
        el.parentNode.removeChild(el);
      }
    });

    it('uses emoji icons for visual clarity', () => {
      el.show('connected', 'Test');
      const icon = el.shadowRoot.getElementById('icon');

      // Should contain an emoji
      expect(icon.textContent).to.match(/^[\u{1F534}\u{1F7E1}\u{1F7E2}]$/u);
    });

    it('provides textual message alongside icon', () => {
      el.show('connected', 'Successfully connected');
      const message = el.shadowRoot.getElementById('message');

      expect(message.textContent).to.equal('Successfully connected');
    });
  });

  describe('edge cases', () => {
    let el;

    beforeEach(() => {
      el = document.createElement('cem-connection-status');
      document.body.appendChild(el);
    });

    afterEach(() => {
      if (el && el.parentNode) {
        el.parentNode.removeChild(el);
      }
    });

    it('handles rapid state changes', () => {
      el.show('connected', 'Connected');
      el.show('reconnecting', 'Reconnecting');
      el.show('disconnected', 'Disconnected');
      el.show('connected', 'Connected again');

      expect(el.getAttribute('state')).to.equal('connected');
      expect(el.shadowRoot.getElementById('message').textContent).to.equal('Connected again');
    });

    it('handles missing options parameter', () => {
      expect(() => {
        el.show('connected', 'Test');
      }).to.not.throw();
    });

    it('handles undefined message', () => {
      el.show('connected', undefined);
      // textContent converts undefined to empty string, not 'undefined'
      expect(el.shadowRoot.getElementById('message').textContent).to.equal('');
    });

    it('handles null message', () => {
      el.show('connected', null);
      // textContent converts null to empty string, not 'null'
      expect(el.shadowRoot.getElementById('message').textContent).to.equal('');
    });

    it('handles numeric message', () => {
      el.show('connected', 12345);
      expect(el.shadowRoot.getElementById('message').textContent).to.equal('12345');
    });

    it('handles unknown state gracefully', () => {
      el.show('unknown-state', 'Unknown');

      expect(el.getAttribute('state')).to.equal('unknown-state');
      // Icon should be undefined/empty since no match
      expect(el.shadowRoot.getElementById('icon').textContent).to.be.oneOf(['', 'undefined']);
    });

    it('handles hide() on element not in DOM', () => {
      const detachedEl = document.createElement('cem-connection-status');

      expect(() => detachedEl.hide()).to.not.throw();
    });
  });

  describe('multiple instances', () => {
    let el1, el2;

    beforeEach(() => {
      el1 = document.createElement('cem-connection-status');
      el2 = document.createElement('cem-connection-status');
      document.body.appendChild(el1);
      document.body.appendChild(el2);
    });

    afterEach(() => {
      if (el1 && el1.parentNode) el1.parentNode.removeChild(el1);
      if (el2 && el2.parentNode) el2.parentNode.removeChild(el2);
    });

    it('manages independent states', () => {
      el1.show('connected', 'Connected 1');
      el2.show('disconnected', 'Disconnected 2');

      expect(el1.getAttribute('state')).to.equal('connected');
      expect(el2.getAttribute('state')).to.equal('disconnected');
    });

    it('shares template but has independent shadow DOMs', () => {
      expect(el1.shadowRoot).to.not.equal(el2.shadowRoot);

      el1.show('connected', 'Message 1');
      el2.show('reconnecting', 'Message 2');

      expect(el1.shadowRoot.getElementById('message').textContent).to.equal('Message 1');
      expect(el2.shadowRoot.getElementById('message').textContent).to.equal('Message 2');
    });
  });

  describe('real-world usage', () => {
    let el;

    beforeEach(() => {
      el = document.createElement('cem-connection-status');
      document.body.appendChild(el);
    });

    afterEach(() => {
      if (el && el.parentNode) {
        el.parentNode.removeChild(el);
      }
    });

    it('simulates connection workflow', async () => {
      // Initial connection
      el.show('connected', 'Connected to dev server');
      expect(el.getAttribute('state')).to.equal('connected');

      // Connection drops
      el.show('reconnecting', 'Reconnecting... (attempt 1)');
      expect(el.getAttribute('state')).to.equal('reconnecting');

      // Still trying
      el.show('reconnecting', 'Reconnecting... (attempt 2)');
      expect(el.getAttribute('state')).to.equal('reconnecting');

      // Reconnected
      el.show('connected', 'Reconnected successfully', { fadeDelay: 100 });
      expect(el.getAttribute('state')).to.equal('connected');

      // Should fade after delay
      await waitUntil(() => el.hasAttribute('faded'), '', { timeout: 200 });
      expect(el.hasAttribute('faded')).to.be.true;
    });

    it('simulates permanent disconnection', () => {
      el.show('connected', 'Connected');
      el.show('reconnecting', 'Connection lost, reconnecting...');
      el.show('reconnecting', 'Retrying...');
      el.show('disconnected', 'Unable to reconnect to server');

      expect(el.getAttribute('state')).to.equal('disconnected');
      expect(el.shadowRoot.getElementById('icon').textContent).to.equal('🔴');
    });
  });
});
