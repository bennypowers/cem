import { expect, waitUntil } from '@open-wc/testing';
import './cem-transform-error-overlay.js';

describe('cem-transform-error-overlay', () => {
  let el;

  beforeEach(async () => {
    el = document.createElement('cem-transform-error-overlay');
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
      const element = document.createElement('cem-transform-error-overlay');
      expect(element).to.be.instanceOf(HTMLElement);
    });

    it('extends LitElement', () => {
      expect(el.constructor.name).to.equal('CemTransformErrorOverlay');
      expect(el.shadowRoot).to.exist;
    });

    it('renders overlay content', () => {
      const overlay = el.shadowRoot.getElementById('overlay-content');
      expect(overlay).to.exist;
    });

    it('renders title, message, and file elements', () => {
      const titleContainer = el.shadowRoot.getElementById('title-container');
      const message = el.shadowRoot.getElementById('message');
      const file = el.shadowRoot.getElementById('file');

      expect(titleContainer).to.exist;
      expect(message).to.exist;
      expect(file).to.exist;
    });

    it('renders close button', () => {
      const closeButton = el.shadowRoot.getElementById('close');
      expect(closeButton).to.exist;
    });

    it('starts with overlay closed', () => {
      expect(el.hasAttribute('open')).to.be.false;
    });
  });

  describe('show() method', () => {
    it('opens the overlay', async () => {
      el.show('Error Title', 'Error message');
      await el.updateComplete;

      expect(el.hasAttribute('open')).to.be.true;
    });

    it('displays title and message', async () => {
      el.show('Transform Error', 'Failed to parse TypeScript');
      await el.updateComplete;

      const title = el.shadowRoot.querySelector('#title-container span:last-child');
      const message = el.shadowRoot.getElementById('message');

      expect(title.textContent).to.equal('Transform Error');
      expect(message.textContent).to.equal('Failed to parse TypeScript');
    });

    it('displays file path when provided', async () => {
      el.show('Error', 'Message', '/src/component.ts');
      await el.updateComplete;

      const file = el.shadowRoot.getElementById('file');
      expect(file.textContent).to.equal('File: /src/component.ts');
      expect(file.textContent).to.not.equal('');
    });

    it('hides file element when file path not provided', async () => {
      el.show('Error', 'Message');
      await el.updateComplete;

      const file = el.shadowRoot.getElementById('file');
      expect(file.textContent).to.equal('');
    });

    it('hides file element when file path is empty string', async () => {
      el.show('Error', 'Message', '');
      await el.updateComplete;

      const file = el.shadowRoot.getElementById('file');
      expect(file.textContent).to.equal('');
    });

    it('updates content on subsequent calls', async () => {
      el.show('Error 1', 'Message 1', '/file1.ts');
      await el.updateComplete;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal('Error 1');

      el.show('Error 2', 'Message 2', '/file2.ts');
      await el.updateComplete;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal('Error 2');
      expect(el.shadowRoot.getElementById('message').textContent).to.equal('Message 2');
      expect(el.shadowRoot.getElementById('file').textContent).to.equal('File: /file2.ts');
    });

    it('handles long error messages', async () => {
      const longMessage = 'This is a very long error message that contains multiple lines\nand detailed stack trace information\nwith file paths and line numbers';
      el.show('Error', longMessage);
      await el.updateComplete;

      const message = el.shadowRoot.getElementById('message');
      expect(message.textContent).to.equal(longMessage);
    });

    it('handles special characters in content', async () => {
      el.show('Error <>&"\'', 'Message with <script>alert("xss")</script>');
      await el.updateComplete;

      const title = el.shadowRoot.querySelector('#title-container span:last-child');
      const message = el.shadowRoot.getElementById('message');

      expect(title.textContent).to.equal('Error <>&"\'');
      expect(message.textContent).to.equal('Message with <script>alert("xss")</script>');
    });

    it('handles HTML entities safely', async () => {
      el.show('&lt;Error&gt;', '&amp; message');
      await el.updateComplete;

      const title = el.shadowRoot.querySelector('#title-container span:last-child');
      expect(title.textContent).to.equal('&lt;Error&gt;');
    });
  });

  describe('hide() method', () => {
    beforeEach(async () => {
      el.show('Test Error', 'Test message');
      await el.updateComplete;
    });

    it('closes the overlay', async () => {
      el.hide();
      await el.updateComplete;

      expect(el.hasAttribute('open')).to.be.false;
    });

    it('can be called multiple times safely', () => {
      expect(() => {
        el.hide();
        el.hide();
        el.hide();
      }).to.not.throw();
    });

    it('can be called when overlay is already closed', async () => {
      el.hide();
      await el.updateComplete;
      expect(el.hasAttribute('open')).to.be.false;

      expect(() => el.hide()).to.not.throw();
    });
  });

  describe('open getter', () => {
    it('returns false when overlay is closed', () => {
      expect(el.open).to.be.false;
    });

    it('returns true when overlay is open', () => {
      el.show('Error', 'Message');
      expect(el.open).to.be.true;
    });

    it('returns false after closing', () => {
      el.show('Error', 'Message');
      expect(el.open).to.be.true;

      el.hide();
      expect(el.open).to.be.false;
    });

    it('reflects current overlay state', () => {
      expect(el.open).to.be.false;

      el.show('Error', 'Message');
      expect(el.open).to.be.true;

      el.hide();
      expect(el.open).to.be.false;

      el.show('Another Error', 'Another Message');
      expect(el.open).to.be.true;
    });
  });

  describe('close button', () => {
    beforeEach(async () => {
      el.show('Test Error', 'Test message');
      await el.updateComplete;
    });

    it('closes overlay when clicked', async () => {
      const closeButton = el.shadowRoot.getElementById('close');
      closeButton.click();
      await el.updateComplete;

      expect(el.hasAttribute('open')).to.be.false;
    });

    it('changes open getter to false', () => {
      expect(el.open).to.be.true;

      const closeButton = el.shadowRoot.getElementById('close');
      closeButton.click();

      expect(el.open).to.be.false;
    });

    it('can be clicked multiple times', () => {
      const closeButton = el.shadowRoot.getElementById('close');

      expect(() => {
        closeButton.click();
        closeButton.click();
      }).to.not.throw();
    });
  });

  describe('keyboard handling', () => {
    beforeEach(async () => {
      el.show('Test Error', 'Test message');
      await el.updateComplete;
    });

    it('closes on Escape key', async () => {
      expect(el.hasAttribute('open')).to.be.true;

      const event = new KeyboardEvent('keydown', { key: 'Escape' });
      document.dispatchEvent(event);
      await el.updateComplete;

      expect(el.hasAttribute('open')).to.be.false;
    });

    it('does not close on other keys', () => {
      const keys = ['Enter', 'Space', 'Tab', 'a', 'x', 'ArrowUp'];
      keys.forEach(key => {
        const event = new KeyboardEvent('keydown', { key });
        document.dispatchEvent(event);
        // open property should remain true (attribute check is async but property is sync)
        expect(el.open).to.be.true;
      });
    });

    it('Escape key changes open getter to false', () => {
      expect(el.open).to.be.true;

      const event = new KeyboardEvent('keydown', { key: 'Escape' });
      document.dispatchEvent(event);

      expect(el.open).to.be.false;
    });
  });

  describe('state transitions', () => {
    it('transitions from closed to open', () => {
      expect(el.open).to.be.false;

      el.show('Error', 'Message');
      expect(el.open).to.be.true;
    });

    it('transitions from open to closed', () => {
      el.show('Error', 'Message');
      expect(el.open).to.be.true;

      el.hide();
      expect(el.open).to.be.false;
    });

    it('can reopen after closing', () => {
      el.show('Error 1', 'Message 1');
      el.hide();
      expect(el.open).to.be.false;

      el.show('Error 2', 'Message 2');
      expect(el.open).to.be.true;
    });

    it('maintains correct state through rapid transitions', async () => {
      el.show('Error 1', 'Message 1');
      el.hide();
      el.show('Error 2', 'Message 2');
      el.hide();
      el.show('Error 3', 'Message 3');
      await el.updateComplete;

      expect(el.open).to.be.true;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal('Error 3');
    });
  });

  describe('accessibility', () => {
    it('has overlay content', () => {
      const overlay = el.shadowRoot.getElementById('overlay-content');
      expect(overlay).to.exist;
    });

    it('close button exists', () => {
      const closeButton = el.shadowRoot.getElementById('close');
      expect(closeButton).to.exist;
    });

    it('overlay is visible when opened', async () => {
      el.show('Error', 'Message');
      await el.updateComplete;
      expect(el.hasAttribute('open')).to.be.true;
    });

    it('provides semantic structure with heading and content', () => {
      const titleContainer = el.shadowRoot.getElementById('title-container');
      const message = el.shadowRoot.getElementById('message');

      expect(titleContainer).to.exist;
      expect(message).to.exist;
    });
  });

  describe('edge cases', () => {
    it('handles undefined title', async () => {
      el.show(undefined, 'Message');
      await el.updateComplete;
      const title = el.shadowRoot.querySelector('#title-container span:last-child');
      expect(title.textContent).to.equal('');
    });

    it('handles null title', async () => {
      el.show(null, 'Message');
      await el.updateComplete;
      const title = el.shadowRoot.querySelector('#title-container span:last-child');
      expect(title.textContent).to.equal('');
    });

    it('handles undefined message', async () => {
      el.show('Title', undefined);
      await el.updateComplete;
      const message = el.shadowRoot.getElementById('message');
      expect(message.textContent).to.equal('');
    });

    it('handles null message', async () => {
      el.show('Title', null);
      await el.updateComplete;
      const message = el.shadowRoot.getElementById('message');
      expect(message.textContent).to.equal('');
    });

    it('handles numeric title and message', async () => {
      el.show(404, 12345);
      await el.updateComplete;
      const title = el.shadowRoot.querySelector('#title-container span:last-child');
      const message = el.shadowRoot.getElementById('message');

      expect(title.textContent).to.equal('404');
      expect(message.textContent).to.equal('12345');
    });

    it('handles empty strings', async () => {
      el.show('', '');
      await el.updateComplete;
      const title = el.shadowRoot.querySelector('#title-container span:last-child');
      const message = el.shadowRoot.getElementById('message');

      expect(title.textContent).to.equal('');
      expect(message.textContent).to.equal('');
    });

    it('handles multiline messages with proper formatting', async () => {
      const multiline = 'Line 1\nLine 2\nLine 3';
      el.show('Error', multiline);
      await el.updateComplete;

      const message = el.shadowRoot.getElementById('message');
      expect(message.textContent).to.equal(multiline);
    });

    it('handles very long file paths', async () => {
      const longPath = '/very/long/path/to/some/deeply/nested/directory/structure/component.ts';
      el.show('Error', 'Message', longPath);
      await el.updateComplete;

      const file = el.shadowRoot.getElementById('file');
      expect(file.textContent).to.equal('File: ' + longPath);
    });

    it('handles rapid show() calls', async () => {
      el.show('Error 1', 'Message 1', '/file1.ts');
      el.show('Error 2', 'Message 2', '/file2.ts');
      el.show('Error 3', 'Message 3', '/file3.ts');
      await el.updateComplete;

      expect(el.open).to.be.true;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal('Error 3');
      expect(el.shadowRoot.getElementById('message').textContent).to.equal('Message 3');
      expect(el.shadowRoot.getElementById('file').textContent).to.equal('File: /file3.ts');
    });

    it('handles show() with only file path changing', async () => {
      el.show('Same Error', 'Same Message', '/file1.ts');
      await el.updateComplete;
      expect(el.shadowRoot.getElementById('file').textContent).to.equal('File: /file1.ts');

      el.show('Same Error', 'Same Message', '/file2.ts');
      await el.updateComplete;
      expect(el.shadowRoot.getElementById('file').textContent).to.equal('File: /file2.ts');
    });

    it('handles switching from file to no file', async () => {
      el.show('Error', 'Message', '/file.ts');
      await el.updateComplete;
      const file = el.shadowRoot.getElementById('file');
      expect(file.textContent).to.not.equal('');

      el.show('Error', 'Message');
      await el.updateComplete;
      expect(file.textContent).to.equal('');
    });

    it('handles switching from no file to file', async () => {
      el.show('Error', 'Message');
      await el.updateComplete;
      const file = el.shadowRoot.getElementById('file');
      expect(file.textContent).to.equal('');

      el.show('Error', 'Message', '/file.ts');
      await el.updateComplete;
      expect(file.textContent).to.not.equal('');
      expect(file.textContent).to.equal('File: /file.ts');
    });
  });

  describe('real-world usage', () => {
    it('simulates TypeScript transform error', async () => {
      const title = 'TypeScript Transform Error';
      const message = `error TS2304: Cannot find name 'Foo'.

src/component.ts(15,10): error TS2304: Cannot find name 'Foo'`;
      const file = '/home/user/project/src/component.ts';

      el.show(title, message, file);
      await el.updateComplete;

      expect(el.open).to.be.true;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal(title);
      expect(el.shadowRoot.getElementById('message').textContent).to.equal(message);
      expect(el.shadowRoot.getElementById('file').textContent).to.equal('File: ' + file);
    });

    it('simulates module resolution error', async () => {
      const title = 'Module Not Found';
      const message = 'Failed to resolve import "./missing-file" from "/src/app.ts"';
      const file = '/src/app.ts';

      el.show(title, message, file);
      await el.updateComplete;

      expect(el.open).to.be.true;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal(title);
      expect(el.shadowRoot.getElementById('message').textContent).to.equal(message);
    });

    it('simulates error cleared after fix', () => {
      // Show error
      el.show('Syntax Error', 'Unexpected token', '/src/bad.ts');
      expect(el.open).to.be.true;

      // User fixes file, error cleared
      el.hide();
      expect(el.open).to.be.false;
    });

    it('simulates multiple errors in sequence', async () => {
      // First error
      el.show('Error 1', 'First problem', '/file1.ts');
      await el.updateComplete;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal('Error 1');

      // Second error (replaces first)
      el.show('Error 2', 'Second problem', '/file2.ts');
      await el.updateComplete;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal('Error 2');

      // Third error
      el.show('Error 3', 'Third problem', '/file3.ts');
      await el.updateComplete;
      expect(el.shadowRoot.querySelector('#title-container span:last-child').textContent).to.equal('Error 3');

      expect(el.open).to.be.true;
    });

    it('simulates user dismissing error with close button', async () => {
      el.show('Warning', 'Deprecated API usage', '/src/old.ts');
      await el.updateComplete;
      expect(el.open).to.be.true;

      const closeButton = el.shadowRoot.getElementById('close');
      closeButton.click();

      expect(el.open).to.be.false;
    });

    it('simulates user dismissing error with Escape key', () => {
      el.show('Info', 'Build completed with warnings');
      expect(el.open).to.be.true;

      const event = new KeyboardEvent('keydown', { key: 'Escape' });
      document.dispatchEvent(event);

      expect(el.open).to.be.false;
    });
  });

  describe('Lit rendering', () => {
    it('renders template via Lit render()', async () => {
      const newEl = document.createElement('cem-transform-error-overlay');
      document.body.appendChild(newEl);
      await newEl.updateComplete;

      expect(newEl.shadowRoot).to.exist;
      expect(newEl.shadowRoot.getElementById('overlay-content')).to.exist;

      document.body.removeChild(newEl);
    });

    it('can be used immediately after first render', async () => {
      const newEl = document.createElement('cem-transform-error-overlay');
      document.body.appendChild(newEl);
      await newEl.updateComplete;

      // Should work immediately
      newEl.show('Test', 'Message');
      expect(newEl.open).to.be.true;

      document.body.removeChild(newEl);
    });
  });
});
