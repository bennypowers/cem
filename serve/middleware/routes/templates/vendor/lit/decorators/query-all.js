// node_modules/@lit/reactive-element/decorators/base.js
var e = (e3, t, c) => (c.configurable = !0, c.enumerable = !0, Reflect.decorate && typeof t != "object" && Object.defineProperty(e3, t, c), c);

// node_modules/@lit/reactive-element/decorators/query-all.js
var e2;
function r(r2) {
  return (n, o) => e(n, o, { get() {
    return (this.renderRoot ?? (e2 ??= document.createDocumentFragment())).querySelectorAll(r2);
  } });
}
export {
  r as queryAll
};
/*! Bundled license information:

@lit/reactive-element/decorators/base.js:
@lit/reactive-element/decorators/query-all.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=query-all.js.map
