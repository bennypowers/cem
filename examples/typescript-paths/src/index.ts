/**
 * Main entry file demonstrating TypeScript path mappings.
 *
 * This file shows how to use the @components/* and @utils/* aliases
 * defined in tsconfig.json paths configuration.
 */

// Import component using @components/* path mapping
import { TsButton } from '@components/ts-button.js';

// Import utilities using @utils/* path mapping
import { capitalize, kebabCase, truncate } from '@utils/format.js';

// Export for use in demos
export { TsButton, capitalize, kebabCase, truncate };

// Example usage demonstrating the utilities work
console.log('TypeScript Path Mappings Demo:');
console.log('  capitalize("hello") =', capitalize('hello'));
console.log('  kebabCase("myButton") =', kebabCase('myButton'));
console.log('  truncate("very long text", 10) =', truncate('very long text', 10));
