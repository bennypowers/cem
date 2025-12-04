import { expect, waitUntil } from '@open-wc/testing';
import './cem-color-scheme-toggle.js';

describe('cem-color-scheme-toggle', () => {
  // Helper to create element with shadow DOM and content
  function createElementWithShadow() {
    const el = document.createElement('cem-color-scheme-toggle');
    const shadow = el.attachShadow({ mode: 'open' });
    shadow.innerHTML = `
      <div id="toggle-group" role="radiogroup" aria-label="Color scheme">
        <label class="toggle-option">
          <input type="radio" name="color-scheme" value="light">
        </label>
        <label class="toggle-option">
          <input type="radio" name="color-scheme" value="system" checked>
        </label>
        <label class="toggle-option">
          <input type="radio" name="color-scheme" value="dark">
        </label>
      </div>
    `;
    return el;
  }

  beforeEach(() => {
    localStorage.clear();
    document.body.style.colorScheme = '';
  });

  afterEach(() => {
    localStorage.clear();
    document.body.style.colorScheme = '';
  });

  describe('initialization', () => {
    it('is defined as custom element', () => {
      const el = document.createElement('cem-color-scheme-toggle');
      expect(el).to.be.instanceOf(HTMLElement);
    });

    it('restores saved color scheme from localStorage on connect', () => {
      localStorage.setItem('cem-serve-color-scheme', 'dark');

      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      expect(document.body.style.colorScheme).to.equal('dark');

      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      expect(darkRadio.checked).to.be.true;

      document.body.removeChild(el);
    });

    it('defaults to system scheme when no preference saved', () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      expect(document.body.style.colorScheme).to.equal('light dark');

      const systemRadio = el.shadowRoot.querySelector('input[value="system"]');
      expect(systemRadio.checked).to.be.true;

      document.body.removeChild(el);
    });

    it('handles missing shadow root gracefully', () => {
      const el = document.createElement('cem-color-scheme-toggle');
      document.body.appendChild(el);

      // Should not throw
      expect(() => el.connectedCallback()).to.not.throw();

      document.body.removeChild(el);
    });
  });

  describe('color scheme application', () => {
    it('applies light color scheme to body', async () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      lightRadio.click();

      await waitUntil(() => document.body.style.colorScheme === 'light');

      expect(document.body.style.colorScheme).to.equal('light');

      document.body.removeChild(el);
    });

    it('applies dark color scheme to body', async () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      darkRadio.click();

      await waitUntil(() => document.body.style.colorScheme === 'dark');

      expect(document.body.style.colorScheme).to.equal('dark');

      document.body.removeChild(el);
    });

    it('applies system color scheme (both light and dark) to body', async () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      // First set to light
      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      lightRadio.click();
      await waitUntil(() => document.body.style.colorScheme === 'light');

      // Then switch to system
      const systemRadio = el.shadowRoot.querySelector('input[value="system"]');
      systemRadio.click();

      await waitUntil(() => document.body.style.colorScheme === 'light dark');

      expect(document.body.style.colorScheme).to.equal('light dark');

      document.body.removeChild(el);
    });
  });

  describe('localStorage persistence', () => {
    it('saves color scheme preference to localStorage', async () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      darkRadio.click();

      await waitUntil(() => localStorage.getItem('cem-serve-color-scheme') === 'dark');

      expect(localStorage.getItem('cem-serve-color-scheme')).to.equal('dark');

      document.body.removeChild(el);
    });

    it('persists across multiple selections', async () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

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

      document.body.removeChild(el);
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

      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      // Should not throw and should still function
      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');
      expect(() => darkRadio.click()).to.not.throw();

      // Should apply color scheme even without localStorage
      await waitUntil(() => document.body.style.colorScheme === 'dark');
      expect(document.body.style.colorScheme).to.equal('dark');

      document.body.removeChild(el);

      // Restore
      Storage.prototype.setItem = originalSetItem;
      Storage.prototype.getItem = originalGetItem;
    });
  });

  describe('radio button state management', () => {
    it('updates radio button selection on change', async () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

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

      document.body.removeChild(el);
    });

    it('restores correct radio button state from localStorage', () => {
      localStorage.setItem('cem-serve-color-scheme', 'light');

      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      const lightRadio = el.shadowRoot.querySelector('input[value="light"]');
      const systemRadio = el.shadowRoot.querySelector('input[value="system"]');
      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');

      expect(lightRadio.checked).to.be.true;
      expect(systemRadio.checked).to.be.false;
      expect(darkRadio.checked).to.be.false;

      document.body.removeChild(el);
    });

    it('handles all radio buttons with same name', () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      const radios = el.shadowRoot.querySelectorAll('input[type="radio"]');
      expect(radios).to.have.lengthOf(3);

      // All should have same name
      radios.forEach(radio => {
        expect(radio.name).to.equal('color-scheme');
      });

      document.body.removeChild(el);
    });
  });

  describe('accessibility', () => {
    it('has radiogroup role', () => {
      const el = createElementWithShadow();

      const group = el.shadowRoot.getElementById('toggle-group');
      expect(group.getAttribute('role')).to.equal('radiogroup');
    });

    it('has aria-label on radiogroup', () => {
      const el = createElementWithShadow();

      const group = el.shadowRoot.getElementById('toggle-group');
      expect(group.getAttribute('aria-label')).to.equal('Color scheme');
    });
  });

  describe('edge cases', () => {
    it('handles rapid successive clicks', async () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

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

      document.body.removeChild(el);
    });

    it('handles clicking the same radio button multiple times', async () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      const darkRadio = el.shadowRoot.querySelector('input[value="dark"]');

      darkRadio.click();
      await waitUntil(() => document.body.style.colorScheme === 'dark');

      // Click again (should be idempotent)
      darkRadio.click();

      expect(document.body.style.colorScheme).to.equal('dark');
      expect(localStorage.getItem('cem-serve-color-scheme')).to.equal('dark');

      document.body.removeChild(el);
    });

    it('handles invalid localStorage values', () => {
      localStorage.setItem('cem-serve-color-scheme', 'invalid-scheme');

      const el = createElementWithShadow();
      document.body.appendChild(el);
      el.connectedCallback();

      // Should default to system behavior for invalid values
      expect(document.body.style.colorScheme).to.equal('light dark');

      document.body.removeChild(el);
    });
  });

  describe('SSR compatibility', () => {
    it('works with pre-existing shadow DOM', () => {
      const el = createElementWithShadow();
      document.body.appendChild(el);

      expect(el.shadowRoot).to.exist;

      const radios = el.shadowRoot.querySelectorAll('input[type="radio"]');
      expect(radios).to.have.lengthOf(3);

      document.body.removeChild(el);
    });

    it('does not throw when connectedCallback runs without shadow root', () => {
      const el = document.createElement('cem-color-scheme-toggle');

      expect(() => {
        document.body.appendChild(el);
        document.body.removeChild(el);
      }).to.not.throw();
    });
  });
});
