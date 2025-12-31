import { VanillaElement } from './vanilla-declare.js';
import { VanillaElement as NamedBindingElement } from './vanilla-declare.js';
import DefaultImportElement from './vanilla-declare.js';

customElements.define('vanilla-element', VanillaElement);
customElements.define('vanilla-element-named', NamedBindingElement);
customElements.define('vanilla-element-default', DefaultImportElement);
