// node_modules/@lit/reactive-element/decorators/custom-element.js
var t = (t2) => (e, o) => {
  o !== void 0 ? o.addInitializer((() => {
    customElements.define(t2, e);
  })) : customElements.define(t2, e);
};
export {
  t as customElement
};
/*! Bundled license information:

@lit/reactive-element/decorators/custom-element.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=custom-element.js.map
