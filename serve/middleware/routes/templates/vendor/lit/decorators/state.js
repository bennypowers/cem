// node_modules/@lit/reactive-element/css-tag.js
var t = globalThis, e = t.ShadowRoot && (t.ShadyCSS === void 0 || t.ShadyCSS.nativeShadow) && "adoptedStyleSheets" in Document.prototype && "replace" in CSSStyleSheet.prototype, s = /* @__PURE__ */ Symbol(), o = /* @__PURE__ */ new WeakMap(), n = class {
  constructor(t2, e3, o4) {
    if (this._$cssResult$ = !0, o4 !== s) throw Error("CSSResult is not constructable. Use `unsafeCSS` or `css` instead.");
    this.cssText = t2, this.t = e3;
  }
  get styleSheet() {
    let t2 = this.o, s2 = this.t;
    if (e && t2 === void 0) {
      let e3 = s2 !== void 0 && s2.length === 1;
      e3 && (t2 = o.get(s2)), t2 === void 0 && ((this.o = t2 = new CSSStyleSheet()).replaceSync(this.cssText), e3 && o.set(s2, t2));
    }
    return t2;
  }
  toString() {
    return this.cssText;
  }
}, r = (t2) => new n(typeof t2 == "string" ? t2 : t2 + "", void 0, s);
var S = (s2, o4) => {
  if (e) s2.adoptedStyleSheets = o4.map(((t2) => t2 instanceof CSSStyleSheet ? t2 : t2.styleSheet));
  else for (let e3 of o4) {
    let o5 = document.createElement("style"), n4 = t.litNonce;
    n4 !== void 0 && o5.setAttribute("nonce", n4), o5.textContent = e3.cssText, s2.appendChild(o5);
  }
}, c = e ? (t2) => t2 : (t2) => t2 instanceof CSSStyleSheet ? ((t3) => {
  let e3 = "";
  for (let s2 of t3.cssRules) e3 += s2.cssText;
  return r(e3);
})(t2) : t2;

// node_modules/@lit/reactive-element/reactive-element.js
var { is: i2, defineProperty: e2, getOwnPropertyDescriptor: h, getOwnPropertyNames: r2, getOwnPropertySymbols: o2, getPrototypeOf: n2 } = Object, a = globalThis, c2 = a.trustedTypes, l = c2 ? c2.emptyScript : "", p = a.reactiveElementPolyfillSupport, d = (t2, s2) => t2, u = { toAttribute(t2, s2) {
  switch (s2) {
    case Boolean:
      t2 = t2 ? l : null;
      break;
    case Object:
    case Array:
      t2 = t2 == null ? t2 : JSON.stringify(t2);
  }
  return t2;
}, fromAttribute(t2, s2) {
  let i3 = t2;
  switch (s2) {
    case Boolean:
      i3 = t2 !== null;
      break;
    case Number:
      i3 = t2 === null ? null : Number(t2);
      break;
    case Object:
    case Array:
      try {
        i3 = JSON.parse(t2);
      } catch {
        i3 = null;
      }
  }
  return i3;
} }, f = (t2, s2) => !i2(t2, s2), b = { attribute: !0, type: String, converter: u, reflect: !1, useDefault: !1, hasChanged: f };
Symbol.metadata ??= /* @__PURE__ */ Symbol("metadata"), a.litPropertyMetadata ??= /* @__PURE__ */ new WeakMap();
var y = class extends HTMLElement {
  static addInitializer(t2) {
    this._$Ei(), (this.l ??= []).push(t2);
  }
  static get observedAttributes() {
    return this.finalize(), this._$Eh && [...this._$Eh.keys()];
  }
  static createProperty(t2, s2 = b) {
    if (s2.state && (s2.attribute = !1), this._$Ei(), this.prototype.hasOwnProperty(t2) && ((s2 = Object.create(s2)).wrapped = !0), this.elementProperties.set(t2, s2), !s2.noAccessor) {
      let i3 = /* @__PURE__ */ Symbol(), h2 = this.getPropertyDescriptor(t2, i3, s2);
      h2 !== void 0 && e2(this.prototype, t2, h2);
    }
  }
  static getPropertyDescriptor(t2, s2, i3) {
    let { get: e3, set: r5 } = h(this.prototype, t2) ?? { get() {
      return this[s2];
    }, set(t3) {
      this[s2] = t3;
    } };
    return { get: e3, set(s3) {
      let h2 = e3?.call(this);
      r5?.call(this, s3), this.requestUpdate(t2, h2, i3);
    }, configurable: !0, enumerable: !0 };
  }
  static getPropertyOptions(t2) {
    return this.elementProperties.get(t2) ?? b;
  }
  static _$Ei() {
    if (this.hasOwnProperty(d("elementProperties"))) return;
    let t2 = n2(this);
    t2.finalize(), t2.l !== void 0 && (this.l = [...t2.l]), this.elementProperties = new Map(t2.elementProperties);
  }
  static finalize() {
    if (this.hasOwnProperty(d("finalized"))) return;
    if (this.finalized = !0, this._$Ei(), this.hasOwnProperty(d("properties"))) {
      let t3 = this.properties, s2 = [...r2(t3), ...o2(t3)];
      for (let i3 of s2) this.createProperty(i3, t3[i3]);
    }
    let t2 = this[Symbol.metadata];
    if (t2 !== null) {
      let s2 = litPropertyMetadata.get(t2);
      if (s2 !== void 0) for (let [t3, i3] of s2) this.elementProperties.set(t3, i3);
    }
    this._$Eh = /* @__PURE__ */ new Map();
    for (let [t3, s2] of this.elementProperties) {
      let i3 = this._$Eu(t3, s2);
      i3 !== void 0 && this._$Eh.set(i3, t3);
    }
    this.elementStyles = this.finalizeStyles(this.styles);
  }
  static finalizeStyles(s2) {
    let i3 = [];
    if (Array.isArray(s2)) {
      let e3 = new Set(s2.flat(1 / 0).reverse());
      for (let s3 of e3) i3.unshift(c(s3));
    } else s2 !== void 0 && i3.push(c(s2));
    return i3;
  }
  static _$Eu(t2, s2) {
    let i3 = s2.attribute;
    return i3 === !1 ? void 0 : typeof i3 == "string" ? i3 : typeof t2 == "string" ? t2.toLowerCase() : void 0;
  }
  constructor() {
    super(), this._$Ep = void 0, this.isUpdatePending = !1, this.hasUpdated = !1, this._$Em = null, this._$Ev();
  }
  _$Ev() {
    this._$ES = new Promise(((t2) => this.enableUpdating = t2)), this._$AL = /* @__PURE__ */ new Map(), this._$E_(), this.requestUpdate(), this.constructor.l?.forEach(((t2) => t2(this)));
  }
  addController(t2) {
    (this._$EO ??= /* @__PURE__ */ new Set()).add(t2), this.renderRoot !== void 0 && this.isConnected && t2.hostConnected?.();
  }
  removeController(t2) {
    this._$EO?.delete(t2);
  }
  _$E_() {
    let t2 = /* @__PURE__ */ new Map(), s2 = this.constructor.elementProperties;
    for (let i3 of s2.keys()) this.hasOwnProperty(i3) && (t2.set(i3, this[i3]), delete this[i3]);
    t2.size > 0 && (this._$Ep = t2);
  }
  createRenderRoot() {
    let t2 = this.shadowRoot ?? this.attachShadow(this.constructor.shadowRootOptions);
    return S(t2, this.constructor.elementStyles), t2;
  }
  connectedCallback() {
    this.renderRoot ??= this.createRenderRoot(), this.enableUpdating(!0), this._$EO?.forEach(((t2) => t2.hostConnected?.()));
  }
  enableUpdating(t2) {
  }
  disconnectedCallback() {
    this._$EO?.forEach(((t2) => t2.hostDisconnected?.()));
  }
  attributeChangedCallback(t2, s2, i3) {
    this._$AK(t2, i3);
  }
  _$ET(t2, s2) {
    let i3 = this.constructor.elementProperties.get(t2), e3 = this.constructor._$Eu(t2, i3);
    if (e3 !== void 0 && i3.reflect === !0) {
      let h2 = (i3.converter?.toAttribute !== void 0 ? i3.converter : u).toAttribute(s2, i3.type);
      this._$Em = t2, h2 == null ? this.removeAttribute(e3) : this.setAttribute(e3, h2), this._$Em = null;
    }
  }
  _$AK(t2, s2) {
    let i3 = this.constructor, e3 = i3._$Eh.get(t2);
    if (e3 !== void 0 && this._$Em !== e3) {
      let t3 = i3.getPropertyOptions(e3), h2 = typeof t3.converter == "function" ? { fromAttribute: t3.converter } : t3.converter?.fromAttribute !== void 0 ? t3.converter : u;
      this._$Em = e3;
      let r5 = h2.fromAttribute(s2, t3.type);
      this[e3] = r5 ?? this._$Ej?.get(e3) ?? r5, this._$Em = null;
    }
  }
  requestUpdate(t2, s2, i3) {
    if (t2 !== void 0) {
      let e3 = this.constructor, h2 = this[t2];
      if (i3 ??= e3.getPropertyOptions(t2), !((i3.hasChanged ?? f)(h2, s2) || i3.useDefault && i3.reflect && h2 === this._$Ej?.get(t2) && !this.hasAttribute(e3._$Eu(t2, i3)))) return;
      this.C(t2, s2, i3);
    }
    this.isUpdatePending === !1 && (this._$ES = this._$EP());
  }
  C(t2, s2, { useDefault: i3, reflect: e3, wrapped: h2 }, r5) {
    i3 && !(this._$Ej ??= /* @__PURE__ */ new Map()).has(t2) && (this._$Ej.set(t2, r5 ?? s2 ?? this[t2]), h2 !== !0 || r5 !== void 0) || (this._$AL.has(t2) || (this.hasUpdated || i3 || (s2 = void 0), this._$AL.set(t2, s2)), e3 === !0 && this._$Em !== t2 && (this._$Eq ??= /* @__PURE__ */ new Set()).add(t2));
  }
  async _$EP() {
    this.isUpdatePending = !0;
    try {
      await this._$ES;
    } catch (t3) {
      Promise.reject(t3);
    }
    let t2 = this.scheduleUpdate();
    return t2 != null && await t2, !this.isUpdatePending;
  }
  scheduleUpdate() {
    return this.performUpdate();
  }
  performUpdate() {
    if (!this.isUpdatePending) return;
    if (!this.hasUpdated) {
      if (this.renderRoot ??= this.createRenderRoot(), this._$Ep) {
        for (let [t4, s3] of this._$Ep) this[t4] = s3;
        this._$Ep = void 0;
      }
      let t3 = this.constructor.elementProperties;
      if (t3.size > 0) for (let [s3, i3] of t3) {
        let { wrapped: t4 } = i3, e3 = this[s3];
        t4 !== !0 || this._$AL.has(s3) || e3 === void 0 || this.C(s3, void 0, i3, e3);
      }
    }
    let t2 = !1, s2 = this._$AL;
    try {
      t2 = this.shouldUpdate(s2), t2 ? (this.willUpdate(s2), this._$EO?.forEach(((t3) => t3.hostUpdate?.())), this.update(s2)) : this._$EM();
    } catch (s3) {
      throw t2 = !1, this._$EM(), s3;
    }
    t2 && this._$AE(s2);
  }
  willUpdate(t2) {
  }
  _$AE(t2) {
    this._$EO?.forEach(((t3) => t3.hostUpdated?.())), this.hasUpdated || (this.hasUpdated = !0, this.firstUpdated(t2)), this.updated(t2);
  }
  _$EM() {
    this._$AL = /* @__PURE__ */ new Map(), this.isUpdatePending = !1;
  }
  get updateComplete() {
    return this.getUpdateComplete();
  }
  getUpdateComplete() {
    return this._$ES;
  }
  shouldUpdate(t2) {
    return !0;
  }
  update(t2) {
    this._$Eq &&= this._$Eq.forEach(((t3) => this._$ET(t3, this[t3]))), this._$EM();
  }
  updated(t2) {
  }
  firstUpdated(t2) {
  }
};
y.elementStyles = [], y.shadowRootOptions = { mode: "open" }, y[d("elementProperties")] = /* @__PURE__ */ new Map(), y[d("finalized")] = /* @__PURE__ */ new Map(), p?.({ ReactiveElement: y }), (a.reactiveElementVersions ??= []).push("2.1.1");

// node_modules/@lit/reactive-element/decorators/property.js
var o3 = { attribute: !0, type: String, converter: u, reflect: !1, hasChanged: f }, r3 = (t2 = o3, e3, r5) => {
  let { kind: n4, metadata: i3 } = r5, s2 = globalThis.litPropertyMetadata.get(i3);
  if (s2 === void 0 && globalThis.litPropertyMetadata.set(i3, s2 = /* @__PURE__ */ new Map()), n4 === "setter" && ((t2 = Object.create(t2)).wrapped = !0), s2.set(r5.name, t2), n4 === "accessor") {
    let { name: o4 } = r5;
    return { set(r6) {
      let n5 = e3.get.call(this);
      e3.set.call(this, r6), this.requestUpdate(o4, n5, t2);
    }, init(e4) {
      return e4 !== void 0 && this.C(o4, void 0, t2, e4), e4;
    } };
  }
  if (n4 === "setter") {
    let { name: o4 } = r5;
    return function(r6) {
      let n5 = this[o4];
      e3.call(this, r6), this.requestUpdate(o4, n5, t2);
    };
  }
  throw Error("Unsupported decorator location: " + n4);
};
function n3(t2) {
  return (e3, o4) => typeof o4 == "object" ? r3(t2, e3, o4) : ((t3, e4, o5) => {
    let r5 = e4.hasOwnProperty(o5);
    return e4.constructor.createProperty(o5, t3), r5 ? Object.getOwnPropertyDescriptor(e4, o5) : void 0;
  })(t2, e3, o4);
}

// node_modules/@lit/reactive-element/decorators/state.js
function r4(r5) {
  return n3({ ...r5, state: !0, attribute: !1 });
}
export {
  r4 as state
};
/*! Bundled license information:

@lit/reactive-element/css-tag.js:
  (**
   * @license
   * Copyright 2019 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

@lit/reactive-element/reactive-element.js:
@lit/reactive-element/decorators/property.js:
@lit/reactive-element/decorators/state.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=state.js.map
