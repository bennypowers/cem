import { expect, waitUntil } from '@open-wc/testing';
import sinon from 'sinon';
import './cem-serve-knobs.js';

describe('cem-serve-knobs', () => {
  let el;
  let originalHash;

  beforeEach(async () => {
    // Save original hash
    originalHash = window.location.hash;
    window.location.hash = '';

    el = document.createElement('cem-serve-knobs');
    document.body.appendChild(el);

    // Wait for CemElement to load template from real server
    await el.rendered;
  });

  afterEach(() => {
    if (el && el.parentNode) {
      el.parentNode.removeChild(el);
    }
    // Restore original hash
    window.location.hash = originalHash;
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const element = document.createElement('cem-serve-knobs');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends CemElement', () => {
      expect(el.constructor.name).to.equal('CemServeKnobs');
      expect(el.shadowRoot).to.exist;
    });

    it('renders navigation list', () => {
      const navList = el.shadowRoot.getElementById('nav-list');
      expect(navList).to.exist;
    });

    it('renders slot for knobs content', () => {
      const slot = el.shadowRoot.querySelector('slot');
      expect(slot).to.exist;
    });

    it('renders knobs container', () => {
      const knobsContainer = el.shadowRoot.getElementById('knobs');
      expect(knobsContainer).to.exist;
    });
  });

  describe('navigation generation', () => {
    it('creates nav items from slotted cards', async () => {
      const card1 = document.createElement('pf-v6-card');
      card1.dataset.card = 'instance-1';
      card1.dataset.label = 'First Instance';

      const card2 = document.createElement('pf-v6-card');
      card2.dataset.card = 'instance-2';
      card2.dataset.label = 'Second Instance';

      el.appendChild(card1);
      el.appendChild(card2);

      // Wait for slotchange event to fire and navigation to update
      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      }, 'Should create 2 nav links', { timeout: 1000 });

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks).to.have.lengthOf(2);

      expect(navLinks[0].getAttribute('href')).to.equal('#instance-1');
      expect(navLinks[0].querySelector('.instance-label').textContent).to.equal('First Instance');

      expect(navLinks[1].getAttribute('href')).to.equal('#instance-2');
      expect(navLinks[1].querySelector('.instance-label').textContent).to.equal('Second Instance');
    });

    it('marks first nav item as current by default', async () => {
      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'test-instance';
      el.appendChild(card);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const navLink = el.shadowRoot.querySelector('pf-v6-nav-link');
      expect(navLink.hasAttribute('current')).to.be.true;
    });

    it('generates default IDs when dataset.card is missing', async () => {
      const card1 = document.createElement('pf-v6-card');
      const card2 = document.createElement('pf-v6-card');

      el.appendChild(card1);
      el.appendChild(card2);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks[0].getAttribute('href')).to.equal('#instance-0');
      expect(navLinks[1].getAttribute('href')).to.equal('#instance-1');
    });

    it('generates default labels when dataset.label is missing', async () => {
      const card1 = document.createElement('pf-v6-card');
      const card2 = document.createElement('pf-v6-card');

      el.appendChild(card1);
      el.appendChild(card2);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks[0].querySelector('.instance-label').textContent).to.equal('Instance 1');
      expect(navLinks[1].querySelector('.instance-label').textContent).to.equal('Instance 2');
    });

    it('ignores non-card elements in slot', async () => {
      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'instance-1';

      const div = document.createElement('div');
      const span = document.createElement('span');

      el.appendChild(card);
      el.appendChild(div);
      el.appendChild(span);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks).to.have.lengthOf(1);
    });

    it('updates navigation when cards are added dynamically', async () => {
      const card1 = document.createElement('pf-v6-card');
      card1.dataset.card = 'instance-1';
      el.appendChild(card1);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      // Add another card
      const card2 = document.createElement('pf-v6-card');
      card2.dataset.card = 'instance-2';
      el.appendChild(card2);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks).to.have.lengthOf(2);
    });

    it('updates navigation when cards are removed', async () => {
      const card1 = document.createElement('pf-v6-card');
      card1.dataset.card = 'instance-1';
      const card2 = document.createElement('pf-v6-card');
      card2.dataset.card = 'instance-2';

      el.appendChild(card1);
      el.appendChild(card2);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });

      // Remove a card
      el.removeChild(card1);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks).to.have.lengthOf(1);
      expect(navLinks[0].getAttribute('href')).to.equal('#instance-2');
    });

    it('handles no slotted cards gracefully', () => {
      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks).to.have.lengthOf(0);
    });
  });

  describe('hash-based navigation', () => {
    let card1, card2;

    beforeEach(async () => {
      card1 = document.createElement('pf-v6-card');
      card1.dataset.card = 'instance-1';
      card1.dataset.label = 'First';

      card2 = document.createElement('pf-v6-card');
      card2.dataset.card = 'instance-2';
      card2.dataset.label = 'Second';

      el.appendChild(card1);
      el.appendChild(card2);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });
    });

    it('updates active card when hash changes', () => {
      window.location.hash = '#instance-2';

      // Trigger hashchange event manually since jsdom doesn't auto-trigger
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      expect(card1.classList.contains('active')).to.be.false;
      expect(card2.classList.contains('active')).to.be.true;
    });

    it('updates current nav link when hash changes', () => {
      window.location.hash = '#instance-2';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks[0].hasAttribute('current')).to.be.false;
      expect(navLinks[1].hasAttribute('current')).to.be.true;
    });

    it('handles hash change to first instance', () => {
      window.location.hash = '#instance-1';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      expect(card1.classList.contains('active')).to.be.true;
      expect(card2.classList.contains('active')).to.be.false;

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks[0].hasAttribute('current')).to.be.true;
      expect(navLinks[1].hasAttribute('current')).to.be.false;
    });

    it('handles multiple hash changes', () => {
      window.location.hash = '#instance-1';
      window.dispatchEvent(new HashChangeEvent('hashchange'));
      expect(card1.classList.contains('active')).to.be.true;

      window.location.hash = '#instance-2';
      window.dispatchEvent(new HashChangeEvent('hashchange'));
      expect(card2.classList.contains('active')).to.be.true;

      window.location.hash = '#instance-1';
      window.dispatchEvent(new HashChangeEvent('hashchange'));
      expect(card1.classList.contains('active')).to.be.true;
    });

    it('handles empty hash gracefully', () => {
      window.location.hash = '#instance-1';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      window.location.hash = '';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      // Should not throw or cause errors
      expect(el).to.exist;
    });

    it('handles # only hash gracefully', () => {
      window.location.hash = '#';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      // Should not throw
      expect(el).to.exist;
    });

    it('handles non-existent card ID in hash', () => {
      window.location.hash = '#non-existent';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      // Should not throw, just no cards active
      expect(card1.classList.contains('active')).to.be.false;
      expect(card2.classList.contains('active')).to.be.false;
    });
  });

  describe('scroll behavior', () => {
    let card1, card2, knobsContainer;

    beforeEach(async () => {
      card1 = document.createElement('pf-v6-card');
      card1.dataset.card = 'instance-1';

      card2 = document.createElement('pf-v6-card');
      card2.dataset.card = 'instance-2';

      el.appendChild(card1);
      el.appendChild(card2);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });

      knobsContainer = el.shadowRoot.getElementById('knobs');
    });

    it('scrolls to active card when hash changes', () => {
      const scrollToSpy = sinon.spy(knobsContainer, 'scrollTo');

      window.location.hash = '#instance-2';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      expect(scrollToSpy.called).to.be.true;

      scrollToSpy.restore();
    });

    it('uses smooth scrolling behavior', () => {
      const scrollToSpy = sinon.spy(knobsContainer, 'scrollTo');

      window.location.hash = '#instance-1';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      expect(scrollToSpy.calledOnce).to.be.true;
      const args = scrollToSpy.firstCall.args[0];
      expect(args.behavior).to.equal('smooth');

      scrollToSpy.restore();
    });

    it('does not scroll when no knobs container exists', () => {
      // Remove knobs container temporarily
      const container = el.shadowRoot.getElementById('knobs');
      const parent = container.parentNode;
      parent.removeChild(container);

      // Should not throw
      expect(() => {
        window.location.hash = '#instance-1';
        window.dispatchEvent(new HashChangeEvent('hashchange'));
      }).to.not.throw();
    });
  });

  describe('lifecycle', () => {
    it('sets up hashchange listener on connect', async () => {
      const addEventListenerSpy = sinon.spy(window, 'addEventListener');

      const newEl = document.createElement('cem-serve-knobs');
      document.body.appendChild(newEl);

      // Wait for template to load
      await waitUntil(() => newEl.shadowRoot?.querySelector('#nav-list'), '', {
        timeout: 3000
      });

      // Should have added hashchange listener
      const hashchangeCalls = addEventListenerSpy.getCalls().filter(
        call => call.args[0] === 'hashchange'
      );
      expect(hashchangeCalls.length).to.be.greaterThan(0);

      document.body.removeChild(newEl);
      addEventListenerSpy.restore();
    });

    it('removes hashchange listener on disconnect', async () => {
      const newEl = document.createElement('cem-serve-knobs');
      document.body.appendChild(newEl);

      // Wait for template to load
      await waitUntil(() => newEl.shadowRoot?.querySelector('#nav-list'), '', {
        timeout: 3000
      });

      const removeEventListenerSpy = sinon.spy(window, 'removeEventListener');

      document.body.removeChild(newEl);

      const hashchangeCalls = removeEventListenerSpy.getCalls().filter(
        call => call.args[0] === 'hashchange'
      );
      expect(hashchangeCalls.length).to.be.greaterThan(0);

      removeEventListenerSpy.restore();
    });

    it('handles initial hash on load', async () => {
      window.location.hash = '#test-instance';

      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'test-instance';

      const newEl = document.createElement('cem-serve-knobs');
      newEl.appendChild(card);
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('#nav-list'), '', {
        timeout: 3000
      });

      // Should handle initial hash
      await waitUntil(() => card.classList.contains('active'), 'Card should be active', {
        timeout: 1000
      });

      expect(card.classList.contains('active')).to.be.true;

      document.body.removeChild(newEl);
    });
  });

  describe('navigation structure', () => {
    it('creates pf-v6-nav-item elements', async () => {
      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'instance-1';
      el.appendChild(card);

      await waitUntil(() => {
        const navItems = el.shadowRoot.querySelectorAll('pf-v6-nav-item');
        return navItems.length === 1;
      });

      const navItems = el.shadowRoot.querySelectorAll('pf-v6-nav-item');
      expect(navItems).to.have.lengthOf(1);
    });

    it('nests pf-v6-nav-link inside pf-v6-nav-item', async () => {
      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'instance-1';
      el.appendChild(card);

      await waitUntil(() => {
        const navItems = el.shadowRoot.querySelectorAll('pf-v6-nav-item');
        return navItems.length === 1;
      });

      const navItem = el.shadowRoot.querySelector('pf-v6-nav-item');
      const navLink = navItem.querySelector('pf-v6-nav-link');
      expect(navLink).to.exist;
    });

    it('adds label span inside nav link', async () => {
      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'instance-1';
      card.dataset.label = 'Test Label';
      el.appendChild(card);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const navLink = el.shadowRoot.querySelector('pf-v6-nav-link');
      const label = navLink.querySelector('span.instance-label');

      expect(label).to.exist;
      expect(label.textContent).to.equal('Test Label');
      expect(label.className).to.equal('instance-label');
    });
  });

  describe('edge cases', () => {
    it('handles rapid card additions and removals', async () => {
      const cards = [];
      for (let i = 0; i < 5; i++) {
        const card = document.createElement('pf-v6-card');
        card.dataset.card = `instance-${i}`;
        cards.push(card);
        el.appendChild(card);
      }

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 5;
      });

      // Remove some cards
      cards.slice(0, 3).forEach(card => el.removeChild(card));

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks).to.have.lengthOf(2);
    });

    it('handles special characters in labels', async () => {
      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'instance-1';
      card.dataset.label = 'Label with <>&"\'';
      el.appendChild(card);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const label = el.shadowRoot.querySelector('.instance-label');
      expect(label.textContent).to.equal('Label with <>&"\'');
    });

    it('handles special characters in card IDs', async () => {
      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'instance-with-dashes-and_underscores';
      el.appendChild(card);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const navLink = el.shadowRoot.querySelector('pf-v6-nav-link');
      expect(navLink.getAttribute('href')).to.equal('#instance-with-dashes-and_underscores');
    });

    it('handles very long labels', async () => {
      const longLabel = 'This is a very long label that might overflow the navigation area and needs to be handled gracefully';
      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'instance-1';
      card.dataset.label = longLabel;
      el.appendChild(card);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const label = el.shadowRoot.querySelector('.instance-label');
      expect(label.textContent).to.equal(longLabel);
    });

    it('handles numeric card IDs and labels', async () => {
      const card = document.createElement('pf-v6-card');
      card.dataset.card = '123';
      card.dataset.label = '456';
      el.appendChild(card);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const navLink = el.shadowRoot.querySelector('pf-v6-nav-link');
      const label = navLink.querySelector('.instance-label');

      expect(navLink.getAttribute('href')).to.equal('#123');
      expect(label.textContent).to.equal('456');
    });
  });

  describe('real-world usage', () => {
    it('simulates multi-instance component demo', async () => {
      // Setup three instances like a real demo
      const instances = [
        { id: 'default', label: 'Default' },
        { id: 'custom-theme', label: 'Custom Theme' },
        { id: 'dark-mode', label: 'Dark Mode' }
      ];

      const cards = instances.map(({ id, label }) => {
        const card = document.createElement('pf-v6-card');
        card.dataset.card = id;
        card.dataset.label = label;
        return card;
      });

      cards.forEach(card => el.appendChild(card));

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 3;
      });

      // Navigate between instances
      window.location.hash = '#default';
      window.dispatchEvent(new HashChangeEvent('hashchange'));
      expect(cards[0].classList.contains('active')).to.be.true;

      window.location.hash = '#custom-theme';
      window.dispatchEvent(new HashChangeEvent('hashchange'));
      expect(cards[1].classList.contains('active')).to.be.true;

      window.location.hash = '#dark-mode';
      window.dispatchEvent(new HashChangeEvent('hashchange'));
      expect(cards[2].classList.contains('active')).to.be.true;
    });

    it('simulates user clicking navigation links', async () => {
      const card1 = document.createElement('pf-v6-card');
      card1.dataset.card = 'instance-1';
      const card2 = document.createElement('pf-v6-card');
      card2.dataset.card = 'instance-2';

      el.appendChild(card1);
      el.appendChild(card2);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });

      const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');

      // Simulate clicking second nav link
      window.location.hash = navLinks[1].getAttribute('href');
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      expect(card2.classList.contains('active')).to.be.true;
      expect(navLinks[1].hasAttribute('current')).to.be.true;
    });

    it('handles browser back/forward navigation', async () => {
      const card1 = document.createElement('pf-v6-card');
      card1.dataset.card = 'instance-1';
      const card2 = document.createElement('pf-v6-card');
      card2.dataset.card = 'instance-2';

      el.appendChild(card1);
      el.appendChild(card2);

      await waitUntil(() => {
        const navLinks = el.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 2;
      });

      // Navigate forward
      window.location.hash = '#instance-1';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      window.location.hash = '#instance-2';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      expect(card2.classList.contains('active')).to.be.true;

      // Simulate back button
      window.location.hash = '#instance-1';
      window.dispatchEvent(new HashChangeEvent('hashchange'));

      expect(card1.classList.contains('active')).to.be.true;
      expect(card2.classList.contains('active')).to.be.false;
    });
  });

  describe('template loading', () => {
    it('loads template asynchronously', async () => {
      const newEl = document.createElement('cem-serve-knobs');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('#nav-list'), 'Should load template', {
        timeout: 3000
      });

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.querySelector('#nav-list')).to.exist;

      document.body.removeChild(newEl);
    });

    it('can accept slotted content after template loads', async () => {
      const newEl = document.createElement('cem-serve-knobs');
      document.body.appendChild(newEl);

      await waitUntil(() => newEl.shadowRoot?.querySelector('slot'), '', {
        timeout: 3000
      });

      const card = document.createElement('pf-v6-card');
      card.dataset.card = 'test';
      newEl.appendChild(card);

      await waitUntil(() => {
        const navLinks = newEl.shadowRoot.querySelectorAll('pf-v6-nav-link');
        return navLinks.length === 1;
      });

      const navLinks = newEl.shadowRoot.querySelectorAll('pf-v6-nav-link');
      expect(navLinks).to.have.lengthOf(1);

      document.body.removeChild(newEl);
    });
  });
});
