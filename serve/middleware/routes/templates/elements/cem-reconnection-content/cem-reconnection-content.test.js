import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './cem-reconnection-content.js';

describe('cem-reconnection-content', () => {
  let el;
  let clock;

  beforeEach(async () => {
    el = document.createElement('cem-reconnection-content');
    document.body.appendChild(el);

    // Wait for CemElement to load template from real server
    await el.rendered;
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
    if (clock) {
      clock.restore();
      clock = null;
    }
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('cem-reconnection-content');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('CemReconnectionContent');
      expect(el.shadowRoot).to.exist;
    });

    it('renders retry info element', () => {
      const retryInfo = el.shadowRoot.getElementById('retry-info');
      expect(retryInfo).to.exist;
    });
  });

  describe('updateRetryInfo() method', () => {
    let retryInfo;

    beforeEach(() => {
      clock = sinon.useFakeTimers();
      retryInfo = el.shadowRoot.getElementById('retry-info');
    });

    it('updates retry info immediately', () => {
      el.updateRetryInfo(1, 5000);

      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 5s...');
    });

    it('displays attempt number', () => {
      el.updateRetryInfo(10, 3000);

      expect(retryInfo.textContent).to.include('Attempt #10');
    });

    it('rounds up delay to seconds', () => {
      el.updateRetryInfo(1, 4100);

      expect(retryInfo.textContent).to.include('5s');
    });

    it('updates countdown every 100ms', () => {
      el.updateRetryInfo(1, 5000);

      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 5s...');

      clock.tick(100);
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 5s...');

      clock.tick(900); // Total: 1000ms
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 4s...');

      clock.tick(1000); // Total: 2000ms
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 3s...');
    });

    it('shows "Connecting..." when countdown reaches zero', () => {
      el.updateRetryInfo(1, 1000);

      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 1s...');

      clock.tick(1000);

      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });

    it('stops countdown when reaching zero', () => {
      el.updateRetryInfo(1, 1000);

      clock.tick(1000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');

      // Advance more time - should not change
      clock.tick(5000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });

    it('clears previous countdown when called again', () => {
      el.updateRetryInfo(1, 5000);

      clock.tick(1000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 4s...');

      // Start new countdown
      el.updateRetryInfo(2, 10000);

      expect(retryInfo.textContent).to.equal('Attempt #2. Retrying in 10s...');

      clock.tick(1000);
      expect(retryInfo.textContent).to.equal('Attempt #2. Retrying in 9s...');
    });

    it('handles very short delays', () => {
      el.updateRetryInfo(1, 100);

      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 1s...');

      clock.tick(100);

      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });

    it('handles zero delay', () => {
      el.updateRetryInfo(1, 0);

      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });

    it('handles negative delay', () => {
      el.updateRetryInfo(1, -1000);

      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });

    it('handles very long delays', () => {
      el.updateRetryInfo(1, 60000); // 60 seconds

      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 60s...');

      clock.tick(30000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 30s...');

      clock.tick(30000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });

    it('does nothing when retry info element is missing', () => {
      // Remove retry info element
      const retryInfoEl = el.shadowRoot.getElementById('retry-info');
      retryInfoEl.parentNode.removeChild(retryInfoEl);

      // Should not throw
      expect(() => el.updateRetryInfo(1, 5000)).to.not.throw();
    });
  });

  describe('clearCountdown() method', () => {
    beforeEach(() => {
      clock = sinon.useFakeTimers();
    });

    it('stops countdown updates', () => {
      const retryInfo = el.shadowRoot.getElementById('retry-info');

      el.updateRetryInfo(1, 5000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 5s...');

      clock.tick(1000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 4s...');

      el.clearCountdown();

      clock.tick(1000);
      // Should not have updated
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 4s...');
    });

    it('can be called multiple times safely', () => {
      el.updateRetryInfo(1, 5000);

      expect(() => {
        el.clearCountdown();
        el.clearCountdown();
        el.clearCountdown();
      }).to.not.throw();
    });

    it('can be called without starting countdown', () => {
      expect(() => el.clearCountdown()).to.not.throw();
    });
  });

  describe('lifecycle', () => {
    beforeEach(() => {
      clock = sinon.useFakeTimers();
    });

    it('clears countdown on disconnect', () => {
      el.updateRetryInfo(1, 5000);

      const retryInfo = el.shadowRoot.getElementById('retry-info');
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 5s...');

      // Disconnect
      document.body.removeChild(el);

      clock.tick(5000);

      // Should have stopped updating (still shows 5s)
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 5s...');

      // Re-add to body for cleanup
      document.body.appendChild(el);
    });
  });

  describe('edge cases', () => {
    let retryInfo;

    beforeEach(() => {
      clock = sinon.useFakeTimers();
      retryInfo = el.shadowRoot.getElementById('retry-info');
    });

    it('handles fractional delays', () => {
      el.updateRetryInfo(1, 2500);

      // Should round up to 3s
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 3s...');
    });

    it('handles very large attempt numbers', () => {
      el.updateRetryInfo(9999, 1000);

      expect(retryInfo.textContent).to.equal('Attempt #9999. Retrying in 1s...');
    });

    it('handles rapid updateRetryInfo calls', () => {
      el.updateRetryInfo(1, 5000);
      el.updateRetryInfo(2, 4000);
      el.updateRetryInfo(3, 3000);

      // Should show latest
      expect(retryInfo.textContent).to.equal('Attempt #3. Retrying in 3s...');

      clock.tick(1000);

      // Should continue countdown for latest
      expect(retryInfo.textContent).to.equal('Attempt #3. Retrying in 2s...');
    });

    it('handles countdown at boundary (1000ms)', () => {
      el.updateRetryInfo(1, 1000);

      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 1s...');

      clock.tick(900);
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 1s...');

      clock.tick(100);
      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });

    it('handles countdown at boundary (999ms)', () => {
      el.updateRetryInfo(1, 999);

      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 1s...');

      clock.tick(999);
      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });

    it('handles countdown at boundary (1001ms)', () => {
      el.updateRetryInfo(1, 1001);

      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 2s...');

      clock.tick(100);
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 1s...');

      clock.tick(900);
      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');
    });
  });

  describe('real-world usage', () => {
    let retryInfo;

    beforeEach(() => {
      clock = sinon.useFakeTimers();
      retryInfo = el.shadowRoot.getElementById('retry-info');
    });

    it('simulates connection retry with progressive backoff', () => {
      // First attempt: 1s delay
      el.updateRetryInfo(1, 1000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Retrying in 1s...');

      clock.tick(1000);
      expect(retryInfo.textContent).to.equal('Attempt #1. Connecting...');

      // Second attempt: 2s delay
      el.updateRetryInfo(2, 2000);
      expect(retryInfo.textContent).to.equal('Attempt #2. Retrying in 2s...');

      clock.tick(2000);
      expect(retryInfo.textContent).to.equal('Attempt #2. Connecting...');

      // Third attempt: 4s delay
      el.updateRetryInfo(3, 4000);
      expect(retryInfo.textContent).to.equal('Attempt #3. Retrying in 4s...');

      clock.tick(1000);
      expect(retryInfo.textContent).to.equal('Attempt #3. Retrying in 3s...');

      clock.tick(1000);
      expect(retryInfo.textContent).to.equal('Attempt #3. Retrying in 2s...');
    });

    it('simulates user cancelling reconnection', () => {
      el.updateRetryInfo(5, 10000);
      expect(retryInfo.textContent).to.equal('Attempt #5. Retrying in 10s...');

      clock.tick(3000);
      expect(retryInfo.textContent).to.equal('Attempt #5. Retrying in 7s...');

      // User cancels / component disconnects
      el.clearCountdown();

      clock.tick(10000);
      // Should not have updated
      expect(retryInfo.textContent).to.equal('Attempt #5. Retrying in 7s...');
    });

    it('simulates long reconnection attempts', () => {
      // Attempt 15 with 30s delay
      el.updateRetryInfo(15, 30000);
      expect(retryInfo.textContent).to.equal('Attempt #15. Retrying in 30s...');

      // Fast-forward through countdown
      for (let i = 29; i >= 1; i--) {
        clock.tick(1000);
        expect(retryInfo.textContent).to.equal(`Attempt #15. Retrying in ${i}s...`);
      }

      clock.tick(1000);
      expect(retryInfo.textContent).to.equal('Attempt #15. Connecting...');
    });

    it('simulates rapid reconnection attempts', () => {
      // Rapidly failing connections
      for (let i = 1; i <= 10; i++) {
        el.updateRetryInfo(i, 1000);
        expect(retryInfo.textContent).to.equal(`Attempt #${i}. Retrying in 1s...`);
        clock.tick(1000);
        expect(retryInfo.textContent).to.equal(`Attempt #${i}. Connecting...`);
      }
    });
  });

  describe('template loading', () => {
    it('loads template asynchronously', async () => {
      const newEl = document.createElement('cem-reconnection-content');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.getElementById('retry-info'), 'Should load template', {
        timeout: 3000
      });

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.getElementById('retry-info')).to.exist;

      document.body.removeChild(newEl);
    });
  });
});
