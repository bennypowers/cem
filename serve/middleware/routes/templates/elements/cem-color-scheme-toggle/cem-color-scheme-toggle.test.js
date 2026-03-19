import { expect, waitUntil } from '@open-wc/testing';
import './cem-color-scheme-toggle.js';

describe('cem-color-scheme-toggle', () => {
  // Helper to create a LitElement instance and wait for first render
  async function createElement() {
    const el = document.createElement('cem-color-scheme-toggle');
    document.body.appendChild(el);
    await el.updateComplete;
    return el;
  }

  beforeEach(() => {
    localStorage.clear();
    document.body.style.colorScheme = '';
  });

  afterEach(() => {
    localStorage.clear();
    document.body.style.colorScheme = '';
    // Clean up any elements left in the DOM
    document.querySelectorAll('cem-color-scheme-toggle').forEach(el => el.remove());
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const el = document.createElement('cem-color-scheme-toggle');
      expect(el).to.be.instanceOf(HTMLElement);
    });

    it('restores saved color scheme from localStorage on connect', async () => {
      localStorage.setItem('cem-serve-color-scheme', 'dark');

      const el = await createElement();

      expect(document.body.style.colorScheme).to.equal('dark');

      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      expect(darkRadio.checked).to.be.true;

      el.remove();
    });

    it('defaults to system scheme when no preference saved', async () => {
      const el = await createElement();

      expect(document.body.style.colorScheme).to.equal('light dark');

      const systemRadio = el.shadowRoot.querySelector('input[value="system"]');
      expect(systemRadio.checked).to.be.true;

      el.remove();
    });

    it('has shadow root with rendered content', async () => {
      const el = await createElement();

      expect(el.shadowRoot).to.exist;
      expect(el.shadowRoot.getElementById('toggle-group')).to.exist;

      el.remove();
    });
  });

  describe('color scheme application', () => {
    it('applies light color scheme to body', async () => {
      const el = await createElement();

      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      lightRadio.click();

      await waitUntil(() => document.body.style.colorScheme === 'light');

      expect(document.body.style.colorScheme).to.equal('light');

      el.remove();
    });

    it('applies dark color scheme to body', async () => {
      const el = await createElement();

      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      darkRadio.click();

      await waitUntil(() => document.body.style.colorScheme === 'dark');

      expect(document.body.style.colorScheme).to.equal('dark');

      el.remove();
    });

    it('applies system color scheme (both light and dark) to body', async () => {
      const el = await createElement();

      // First set to light
      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      lightRadio.click();
      await waitUntil(() => document.body.style.colorScheme === 'light');

      // Then switch to system
      const systemRadio = el.shadowRoot.querySelector('input[value="system"]');
      systemRadio.click();

      await waitUntil(() => document.body.style.colorScheme === 'light dark');

      expect(document.body.style.colorScheme).to.equal('light dark');

      el.remove();
    });
  });

  describe('localStorage persistence', () => {
    it('saves color scheme preference to localStorage', async () => {
      const el = await createElement();

      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      darkRadio.click();

      await waitUntil(() => localStorage.getItem('cem-serve-color-scheme') === 'dark');

      expect(localStorage.getItem('cem-serve-color-scheme')).to.equal('dark');

      el.remove();
    });

    it('persists across multiple selections', async () => {
      const el = await createElement();

      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      const systemRadio = el.shadowRoot.querySelector('input[value="system"]');

      lightRadio.click();
      await waitUntil(() => localStorage.getItem('cem-serve-color-scheme') === 'light');
      expect(localStorage.getItem('cem-serve-color-scheme')).to.equal('light');

      darkRadio.click();
      await waitUntil(() => localStorage.getItem('cem-serve-color-scheme') === 'dark');
      expect(localStorage.getItem('cem-serve-color-scheme')).to.equal('dark');

      systemRadio.click();
      await waitUntil(() => localStorage.getItem('cem-serve-color-scheme') === 'system');
      expect(localStorage.getItem('cem-serve-color-scheme')).to.equal('system');

      el.remove();
    });

    it('handles localStorage errors gracefully (Safari private mode)', async () => {
      // Stub localStorage to throw
      const originalSetItem = Storage.prototype.setItem;
      const originalGetItem = Storage.prototype.getItem;

      Storage.prototype.setItem = function() {
        throw new Error('QuotaExceededError');
      };
      Storage.prototype.getItem = function() {
        throw new Error('SecurityError');
      };

      const el = await createElement();

      // Should not throw and should still function
      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      expect(() => darkRadio.click()).to.not.throw();

      // Should apply color scheme even without localStorage
      await waitUntil(() => document.body.style.colorScheme === 'dark');
      expect(document.body.style.colorScheme).to.equal('dark');

      el.remove();

      // Restore
      Storage.prototype.setItem = originalSetItem;
      Storage.prototype.getItem = originalGetItem;
    });
  });

  describe('radio button state management', () => {
    it('updates radio button selection on change', async () => {
      const el = await createElement();

      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      const systemRadio = el.shadowRoot.querySelector('input[value="system"]');
      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');

      // Initially system should be checked
      expect(systemRadio.checked).to.be.true;
      expect(lightRadio.checked).to.be.false;
      expect(darkRadio.checked).to.be.false;

      // Click dark
      darkRadio.click();
      await waitUntil(() => darkRadio.checked);

      expect(darkRadio.checked).to.be.true;
      expect(lightRadio.checked).to.be.false;
      expect(systemRadio.checked).to.be.false;

      // Click light
      lightRadio.click();
      await waitUntil(() => lightRadio.checked);

      expect(lightRadio.checked).to.be.true;
      expect(darkRadio.checked).to.be.false;
      expect(systemRadio.checked).to.be.false;

      el.remove();
    });

    it('restores correct radio button state from localStorage', async () => {
      localStorage.setItem('cem-serve-color-scheme', 'light');

      const el = await createElement();

      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      const systemRadio = el.shadowRoot.querySelector('input[value="system"]');
      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');

      expect(lightRadio.checked).to.be.true;
      expect(systemRadio.checked).to.be.false;
      expect(darkRadio.checked).to.be.false;

      el.remove();
    });

    it('handles all radio buttons with same name', async () => {
      const el = await createElement();

      const radios = el.shadowRoot.querySelectorAll('input[type="radio"]');
      expect(radios).to.have.lengthOf(3);

      // All should have same name
      radios.forEach(radio => {
        expect(radio.name).to.equal('color-scheme');
      });

      el.remove();
    });
  });

  describe('accessibility', () => {
    it('has radiogroup role', async () => {
      const el = await createElement();

      const group = el.shadowRoot.getElementById('toggle-group');
      expect(group.getAttribute('role')).to.equal('radiogroup');

      el.remove();
    });

    it('has aria-label on radiogroup', async () => {
      const el = await createElement();

      const group = el.shadowRoot.getElementById('toggle-group');
      expect(group.getAttribute('aria-label')).to.equal('Color scheme');

      el.remove();
    });
  });

  describe('edge cases', () => {
    it('handles rapid successive clicks', async () => {
      const el = await createElement();

      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');

      // Rapid clicks
      lightRadio.click();
      darkRadio.click();
      lightRadio.click();
      darkRadio.click();

      await waitUntil(() => document.body.style.colorScheme === 'dark');

      expect(document.body.style.colorScheme).to.equal('dark');
      expect(localStorage.getItem('cem-serve-color-scheme')).to.equal('dark');

      el.remove();
    });

    it('handles clicking the same radio button multiple times', async () => {
      const el = await createElement();

      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');

      darkRadio.click();
      await waitUntil(() => document.body.style.colorScheme === 'dark');

      // Click again (should be idempotent)
      darkRadio.click();

      expect(document.body.style.colorScheme).to.equal('dark');
      expect(localStorage.getItem('cem-serve-color-scheme')).to.equal('dark');

      el.remove();
    });

    it('handles invalid localStorage values', async () => {
      localStorage.setItem('cem-serve-color-scheme', 'invalid-scheme');

      const el = await createElement();

      // Should default to system behavior for invalid values
      expect(document.body.style.colorScheme).to.equal('light dark');

      el.remove();
    });
  });

  describe('Lit rendering', () => {
    it('renders shadow DOM via Lit render()', async () => {
      const el = await createElement();

      expect(el.shadowRoot).to.exist;

      const radios = el.shadowRoot.querySelectorAll('input[type="radio"]');
      expect(radios).to.have.lengthOf(3);

      el.remove();
    });

    it('does not throw on connect and disconnect', () => {
      const el = document.createElement('cem-color-scheme-toggle');

      expect(() => {
        document.body.appendChild(el);
        document.body.removeChild(el);
      }).to.not.throw();
    });
  });
});
