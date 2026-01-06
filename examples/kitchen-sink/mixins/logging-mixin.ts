interface CustomElement extends HTMLElement {
  connectedCallback?(): void
  disconnectedCallback?(): void
}

export type Constructor<T = object> = new (...args: any[]) => T;

/**
 * Mixin that adds logging capabilities to custom elements.
 *
 * @summary Provides lifecycle logging and debug utilities
 */
export function LoggingMixin<T extends Constructor<CustomElement>>(Base: T) {
  return class LoggingMixinClass extends Base {
    static properties = {
      enableLogging: { type: Boolean, attribute: 'enable-logging' }
    };

    /**
     * Enable or disable logging for this element
     */
    enableLogging = false;

    /**
     * Log a message to the console if logging is enabled
     *
     * @param message - The message to log
     * @param data - Optional data to log
     */
    log(message: string, ...data: any[]) {
      if (this.enableLogging) {
        console.log(`[${this.tagName.toLowerCase()}]`, message, ...data);
      }
    }

    /**
     * Log a warning to the console if logging is enabled
     *
     * @param message - The warning message
     * @param data - Optional data to log
     */
    warn(message: string, ...data: any[]) {
      if (this.enableLogging) {
        console.warn(`[${this.tagName.toLowerCase()}]`, message, ...data);
      }
    }

    /**
     * Log an error to the console if logging is enabled
     *
     * @param message - The error message
     * @param data - Optional data to log
     */
    error(message: string, ...data: any[]) {
      if (this.enableLogging) {
        console.error(`[${this.tagName.toLowerCase()}]`, message, ...data);
      }
    }

    connectedCallback() {
      super.connectedCallback?.();
      this.log('Connected to DOM');
    }

    disconnectedCallback() {
      super.disconnectedCallback?.();
      this.log('Disconnected from DOM');
    }
  };
}
