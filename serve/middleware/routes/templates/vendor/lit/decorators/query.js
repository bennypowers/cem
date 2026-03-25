// node_modules/@lit/reactive-element/decorators/base.js
var e = (e3, t, c) => (c.configurable = !0, c.enumerable = !0, Reflect.decorate && typeof t != "object" && Object.defineProperty(e3, t, c), c);

// node_modules/@lit/reactive-element/decorators/query.js
function e2(e3, r) {
  return (n, s, i) => {
    let o = (t) => t.renderRoot?.querySelector(e3) ?? null;
    if (r) {
      let { get: e4, set: r2 } = typeof s == "object" ? n : i ?? /* @__PURE__ */ (() => {
        let t = /* @__PURE__ */ Symbol();
        return { get() {
          return this[t];
        }, set(e5) {
          this[t] = e5;
        } };
      })();
      return e(n, s, { get() {
        let t = e4.call(this);
        return t === void 0 && (t = o(this), (t !== null || this.hasUpdated) && r2.call(this, t)), t;
      } });
    }
    return e(n, s, { get() {
      return o(this);
    } });
  };
}
export {
  e2 as query
};
/*! Bundled license information:

@lit/reactive-element/decorators/base.js:
@lit/reactive-element/decorators/query.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=query.js.map
