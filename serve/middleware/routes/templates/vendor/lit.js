// node_modules/@lit/reactive-element/css-tag.js
var t = globalThis, e = t.ShadowRoot && (t.ShadyCSS === void 0 || t.ShadyCSS.nativeShadow) && "adoptedStyleSheets" in Document.prototype && "replace" in CSSStyleSheet.prototype, s = Symbol(), o = /* @__PURE__ */ new WeakMap(), n = class {
  constructor(t3, e4, o6) {
    if (this._$cssResult$ = !0, o6 !== s) throw Error("CSSResult is not constructable. Use `unsafeCSS` or `css` instead.");
    this.cssText = t3, this.t = e4;
  }
  get styleSheet() {
    let t3 = this.o, s4 = this.t;
    if (e && t3 === void 0) {
      let e4 = s4 !== void 0 && s4.length === 1;
      e4 && (t3 = o.get(s4)), t3 === void 0 && ((this.o = t3 = new CSSStyleSheet()).replaceSync(this.cssText), e4 && o.set(s4, t3));
    }
    return t3;
  }
  toString() {
    return this.cssText;
  }
}, r = (t3) => new n(typeof t3 == "string" ? t3 : t3 + "", void 0, s), i = (t3, ...e4) => {
  let o6 = t3.length === 1 ? t3[0] : e4.reduce(((e5, s4, o7) => e5 + ((t4) => {
    if (t4._$cssResult$ === !0) return t4.cssText;
    if (typeof t4 == "number") return t4;
    throw Error("Value passed to 'css' function must be a 'css' function result: " + t4 + ". Use 'unsafeCSS' to pass non-literal values, but take care to ensure page security.");
  })(s4) + t3[o7 + 1]), t3[0]);
  return new n(o6, t3, s);
}, S = (s4, o6) => {
  if (e) s4.adoptedStyleSheets = o6.map(((t3) => t3 instanceof CSSStyleSheet ? t3 : t3.styleSheet));
  else for (let e4 of o6) {
    let o7 = document.createElement("style"), n5 = t.litNonce;
    n5 !== void 0 && o7.setAttribute("nonce", n5), o7.textContent = e4.cssText, s4.appendChild(o7);
  }
}, c = e ? (t3) => t3 : (t3) => t3 instanceof CSSStyleSheet ? ((t4) => {
  let e4 = "";
  for (let s4 of t4.cssRules) e4 += s4.cssText;
  return r(e4);
})(t3) : t3;

// node_modules/@lit/reactive-element/reactive-element.js
var { is: i2, defineProperty: e2, getOwnPropertyDescriptor: h, getOwnPropertyNames: r2, getOwnPropertySymbols: o2, getPrototypeOf: n2 } = Object, a = globalThis, c2 = a.trustedTypes, l = c2 ? c2.emptyScript : "", p = a.reactiveElementPolyfillSupport, d = (t3, s4) => t3, u = { toAttribute(t3, s4) {
  switch (s4) {
    case Boolean:
      t3 = t3 ? l : null;
      break;
    case Object:
    case Array:
      t3 = t3 == null ? t3 : JSON.stringify(t3);
  }
  return t3;
}, fromAttribute(t3, s4) {
  let i5 = t3;
  switch (s4) {
    case Boolean:
      i5 = t3 !== null;
      break;
    case Number:
      i5 = t3 === null ? null : Number(t3);
      break;
    case Object:
    case Array:
      try {
        i5 = JSON.parse(t3);
      } catch {
        i5 = null;
      }
  }
  return i5;
} }, f = (t3, s4) => !i2(t3, s4), b = { attribute: !0, type: String, converter: u, reflect: !1, useDefault: !1, hasChanged: f };
Symbol.metadata ??= Symbol("metadata"), a.litPropertyMetadata ??= /* @__PURE__ */ new WeakMap();
var y = class extends HTMLElement {
  static addInitializer(t3) {
    this._$Ei(), (this.l ??= []).push(t3);
  }
  static get observedAttributes() {
    return this.finalize(), this._$Eh && [...this._$Eh.keys()];
  }
  static createProperty(t3, s4 = b) {
    if (s4.state && (s4.attribute = !1), this._$Ei(), this.prototype.hasOwnProperty(t3) && ((s4 = Object.create(s4)).wrapped = !0), this.elementProperties.set(t3, s4), !s4.noAccessor) {
      let i5 = Symbol(), h3 = this.getPropertyDescriptor(t3, i5, s4);
      h3 !== void 0 && e2(this.prototype, t3, h3);
    }
  }
  static getPropertyDescriptor(t3, s4, i5) {
    let { get: e4, set: r4 } = h(this.prototype, t3) ?? { get() {
      return this[s4];
    }, set(t4) {
      this[s4] = t4;
    } };
    return { get: e4, set(s5) {
      let h3 = e4?.call(this);
      r4?.call(this, s5), this.requestUpdate(t3, h3, i5);
    }, configurable: !0, enumerable: !0 };
  }
  static getPropertyOptions(t3) {
    return this.elementProperties.get(t3) ?? b;
  }
  static _$Ei() {
    if (this.hasOwnProperty(d("elementProperties"))) return;
    let t3 = n2(this);
    t3.finalize(), t3.l !== void 0 && (this.l = [...t3.l]), this.elementProperties = new Map(t3.elementProperties);
  }
  static finalize() {
    if (this.hasOwnProperty(d("finalized"))) return;
    if (this.finalized = !0, this._$Ei(), this.hasOwnProperty(d("properties"))) {
      let t4 = this.properties, s4 = [...r2(t4), ...o2(t4)];
      for (let i5 of s4) this.createProperty(i5, t4[i5]);
    }
    let t3 = this[Symbol.metadata];
    if (t3 !== null) {
      let s4 = litPropertyMetadata.get(t3);
      if (s4 !== void 0) for (let [t4, i5] of s4) this.elementProperties.set(t4, i5);
    }
    this._$Eh = /* @__PURE__ */ new Map();
    for (let [t4, s4] of this.elementProperties) {
      let i5 = this._$Eu(t4, s4);
      i5 !== void 0 && this._$Eh.set(i5, t4);
    }
    this.elementStyles = this.finalizeStyles(this.styles);
  }
  static finalizeStyles(s4) {
    let i5 = [];
    if (Array.isArray(s4)) {
      let e4 = new Set(s4.flat(1 / 0).reverse());
      for (let s5 of e4) i5.unshift(c(s5));
    } else s4 !== void 0 && i5.push(c(s4));
    return i5;
  }
  static _$Eu(t3, s4) {
    let i5 = s4.attribute;
    return i5 === !1 ? void 0 : typeof i5 == "string" ? i5 : typeof t3 == "string" ? t3.toLowerCase() : void 0;
  }
  constructor() {
    super(), this._$Ep = void 0, this.isUpdatePending = !1, this.hasUpdated = !1, this._$Em = null, this._$Ev();
  }
  _$Ev() {
    this._$ES = new Promise(((t3) => this.enableUpdating = t3)), this._$AL = /* @__PURE__ */ new Map(), this._$E_(), this.requestUpdate(), this.constructor.l?.forEach(((t3) => t3(this)));
  }
  addController(t3) {
    (this._$EO ??= /* @__PURE__ */ new Set()).add(t3), this.renderRoot !== void 0 && this.isConnected && t3.hostConnected?.();
  }
  removeController(t3) {
    this._$EO?.delete(t3);
  }
  _$E_() {
    let t3 = /* @__PURE__ */ new Map(), s4 = this.constructor.elementProperties;
    for (let i5 of s4.keys()) this.hasOwnProperty(i5) && (t3.set(i5, this[i5]), delete this[i5]);
    t3.size > 0 && (this._$Ep = t3);
  }
  createRenderRoot() {
    let t3 = this.shadowRoot ?? this.attachShadow(this.constructor.shadowRootOptions);
    return S(t3, this.constructor.elementStyles), t3;
  }
  connectedCallback() {
    this.renderRoot ??= this.createRenderRoot(), this.enableUpdating(!0), this._$EO?.forEach(((t3) => t3.hostConnected?.()));
  }
  enableUpdating(t3) {
  }
  disconnectedCallback() {
    this._$EO?.forEach(((t3) => t3.hostDisconnected?.()));
  }
  attributeChangedCallback(t3, s4, i5) {
    this._$AK(t3, i5);
  }
  _$ET(t3, s4) {
    let i5 = this.constructor.elementProperties.get(t3), e4 = this.constructor._$Eu(t3, i5);
    if (e4 !== void 0 && i5.reflect === !0) {
      let h3 = (i5.converter?.toAttribute !== void 0 ? i5.converter : u).toAttribute(s4, i5.type);
      this._$Em = t3, h3 == null ? this.removeAttribute(e4) : this.setAttribute(e4, h3), this._$Em = null;
    }
  }
  _$AK(t3, s4) {
    let i5 = this.constructor, e4 = i5._$Eh.get(t3);
    if (e4 !== void 0 && this._$Em !== e4) {
      let t4 = i5.getPropertyOptions(e4), h3 = typeof t4.converter == "function" ? { fromAttribute: t4.converter } : t4.converter?.fromAttribute !== void 0 ? t4.converter : u;
      this._$Em = e4;
      let r4 = h3.fromAttribute(s4, t4.type);
      this[e4] = r4 ?? this._$Ej?.get(e4) ?? r4, this._$Em = null;
    }
  }
  requestUpdate(t3, s4, i5) {
    if (t3 !== void 0) {
      let e4 = this.constructor, h3 = this[t3];
      if (i5 ??= e4.getPropertyOptions(t3), !((i5.hasChanged ?? f)(h3, s4) || i5.useDefault && i5.reflect && h3 === this._$Ej?.get(t3) && !this.hasAttribute(e4._$Eu(t3, i5)))) return;
      this.C(t3, s4, i5);
    }
    this.isUpdatePending === !1 && (this._$ES = this._$EP());
  }
  C(t3, s4, { useDefault: i5, reflect: e4, wrapped: h3 }, r4) {
    i5 && !(this._$Ej ??= /* @__PURE__ */ new Map()).has(t3) && (this._$Ej.set(t3, r4 ?? s4 ?? this[t3]), h3 !== !0 || r4 !== void 0) || (this._$AL.has(t3) || (this.hasUpdated || i5 || (s4 = void 0), this._$AL.set(t3, s4)), e4 === !0 && this._$Em !== t3 && (this._$Eq ??= /* @__PURE__ */ new Set()).add(t3));
  }
  async _$EP() {
    this.isUpdatePending = !0;
    try {
      await this._$ES;
    } catch (t4) {
      Promise.reject(t4);
    }
    let t3 = this.scheduleUpdate();
    return t3 != null && await t3, !this.isUpdatePending;
  }
  scheduleUpdate() {
    return this.performUpdate();
  }
  performUpdate() {
    if (!this.isUpdatePending) return;
    if (!this.hasUpdated) {
      if (this.renderRoot ??= this.createRenderRoot(), this._$Ep) {
        for (let [t5, s5] of this._$Ep) this[t5] = s5;
        this._$Ep = void 0;
      }
      let t4 = this.constructor.elementProperties;
      if (t4.size > 0) for (let [s5, i5] of t4) {
        let { wrapped: t5 } = i5, e4 = this[s5];
        t5 !== !0 || this._$AL.has(s5) || e4 === void 0 || this.C(s5, void 0, i5, e4);
      }
    }
    let t3 = !1, s4 = this._$AL;
    try {
      t3 = this.shouldUpdate(s4), t3 ? (this.willUpdate(s4), this._$EO?.forEach(((t4) => t4.hostUpdate?.())), this.update(s4)) : this._$EM();
    } catch (s5) {
      throw t3 = !1, this._$EM(), s5;
    }
    t3 && this._$AE(s4);
  }
  willUpdate(t3) {
  }
  _$AE(t3) {
    this._$EO?.forEach(((t4) => t4.hostUpdated?.())), this.hasUpdated || (this.hasUpdated = !0, this.firstUpdated(t3)), this.updated(t3);
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
  shouldUpdate(t3) {
    return !0;
  }
  update(t3) {
    this._$Eq &&= this._$Eq.forEach(((t4) => this._$ET(t4, this[t4]))), this._$EM();
  }
  updated(t3) {
  }
  firstUpdated(t3) {
  }
};
y.elementStyles = [], y.shadowRootOptions = { mode: "open" }, y[d("elementProperties")] = /* @__PURE__ */ new Map(), y[d("finalized")] = /* @__PURE__ */ new Map(), p?.({ ReactiveElement: y }), (a.reactiveElementVersions ??= []).push("2.1.1");

// node_modules/lit-html/lit-html.js
var t2 = globalThis, i3 = t2.trustedTypes, s2 = i3 ? i3.createPolicy("lit-html", { createHTML: (t3) => t3 }) : void 0, e3 = "$lit$", h2 = `lit$${Math.random().toFixed(9).slice(2)}$`, o3 = "?" + h2, n3 = `<${o3}>`, r3 = document, l2 = () => r3.createComment(""), c3 = (t3) => t3 === null || typeof t3 != "object" && typeof t3 != "function", a2 = Array.isArray, u2 = (t3) => a2(t3) || typeof t3?.[Symbol.iterator] == "function", d2 = `[ 	
\f\r]`, f2 = /<(?:(!--|\/[^a-zA-Z])|(\/?[a-zA-Z][^>\s]*)|(\/?$))/g, v = /-->/g, _ = />/g, m = RegExp(`>|${d2}(?:([^\\s"'>=/]+)(${d2}*=${d2}*(?:[^ 	
\f\r"'\`<>=]|("|')|))|$)`, "g"), p2 = /'/g, g = /"/g, $ = /^(?:script|style|textarea|title)$/i, y2 = (t3) => (i5, ...s4) => ({ _$litType$: t3, strings: i5, values: s4 }), x = y2(1), b2 = y2(2), w = y2(3), T = Symbol.for("lit-noChange"), E = Symbol.for("lit-nothing"), A = /* @__PURE__ */ new WeakMap(), C = r3.createTreeWalker(r3, 129);
function P(t3, i5) {
  if (!a2(t3) || !t3.hasOwnProperty("raw")) throw Error("invalid template strings array");
  return s2 !== void 0 ? s2.createHTML(i5) : i5;
}
var V = (t3, i5) => {
  let s4 = t3.length - 1, o6 = [], r4, l3 = i5 === 2 ? "<svg>" : i5 === 3 ? "<math>" : "", c4 = f2;
  for (let i6 = 0; i6 < s4; i6++) {
    let s5 = t3[i6], a3, u3, d3 = -1, y3 = 0;
    for (; y3 < s5.length && (c4.lastIndex = y3, u3 = c4.exec(s5), u3 !== null); ) y3 = c4.lastIndex, c4 === f2 ? u3[1] === "!--" ? c4 = v : u3[1] !== void 0 ? c4 = _ : u3[2] !== void 0 ? ($.test(u3[2]) && (r4 = RegExp("</" + u3[2], "g")), c4 = m) : u3[3] !== void 0 && (c4 = m) : c4 === m ? u3[0] === ">" ? (c4 = r4 ?? f2, d3 = -1) : u3[1] === void 0 ? d3 = -2 : (d3 = c4.lastIndex - u3[2].length, a3 = u3[1], c4 = u3[3] === void 0 ? m : u3[3] === '"' ? g : p2) : c4 === g || c4 === p2 ? c4 = m : c4 === v || c4 === _ ? c4 = f2 : (c4 = m, r4 = void 0);
    let x2 = c4 === m && t3[i6 + 1].startsWith("/>") ? " " : "";
    l3 += c4 === f2 ? s5 + n3 : d3 >= 0 ? (o6.push(a3), s5.slice(0, d3) + e3 + s5.slice(d3) + h2 + x2) : s5 + h2 + (d3 === -2 ? i6 : x2);
  }
  return [P(t3, l3 + (t3[s4] || "<?>") + (i5 === 2 ? "</svg>" : i5 === 3 ? "</math>" : "")), o6];
}, N = class _N {
  constructor({ strings: t3, _$litType$: s4 }, n5) {
    let r4;
    this.parts = [];
    let c4 = 0, a3 = 0, u3 = t3.length - 1, d3 = this.parts, [f3, v2] = V(t3, s4);
    if (this.el = _N.createElement(f3, n5), C.currentNode = this.el.content, s4 === 2 || s4 === 3) {
      let t4 = this.el.content.firstChild;
      t4.replaceWith(...t4.childNodes);
    }
    for (; (r4 = C.nextNode()) !== null && d3.length < u3; ) {
      if (r4.nodeType === 1) {
        if (r4.hasAttributes()) for (let t4 of r4.getAttributeNames()) if (t4.endsWith(e3)) {
          let i5 = v2[a3++], s5 = r4.getAttribute(t4).split(h2), e4 = /([.?@])?(.*)/.exec(i5);
          d3.push({ type: 1, index: c4, name: e4[2], strings: s5, ctor: e4[1] === "." ? H : e4[1] === "?" ? I : e4[1] === "@" ? L : k }), r4.removeAttribute(t4);
        } else t4.startsWith(h2) && (d3.push({ type: 6, index: c4 }), r4.removeAttribute(t4));
        if ($.test(r4.tagName)) {
          let t4 = r4.textContent.split(h2), s5 = t4.length - 1;
          if (s5 > 0) {
            r4.textContent = i3 ? i3.emptyScript : "";
            for (let i5 = 0; i5 < s5; i5++) r4.append(t4[i5], l2()), C.nextNode(), d3.push({ type: 2, index: ++c4 });
            r4.append(t4[s5], l2());
          }
        }
      } else if (r4.nodeType === 8) if (r4.data === o3) d3.push({ type: 2, index: c4 });
      else {
        let t4 = -1;
        for (; (t4 = r4.data.indexOf(h2, t4 + 1)) !== -1; ) d3.push({ type: 7, index: c4 }), t4 += h2.length - 1;
      }
      c4++;
    }
  }
  static createElement(t3, i5) {
    let s4 = r3.createElement("template");
    return s4.innerHTML = t3, s4;
  }
};
function S2(t3, i5, s4 = t3, e4) {
  if (i5 === T) return i5;
  let h3 = e4 !== void 0 ? s4._$Co?.[e4] : s4._$Cl, o6 = c3(i5) ? void 0 : i5._$litDirective$;
  return h3?.constructor !== o6 && (h3?._$AO?.(!1), o6 === void 0 ? h3 = void 0 : (h3 = new o6(t3), h3._$AT(t3, s4, e4)), e4 !== void 0 ? (s4._$Co ??= [])[e4] = h3 : s4._$Cl = h3), h3 !== void 0 && (i5 = S2(t3, h3._$AS(t3, i5.values), h3, e4)), i5;
}
var M = class {
  constructor(t3, i5) {
    this._$AV = [], this._$AN = void 0, this._$AD = t3, this._$AM = i5;
  }
  get parentNode() {
    return this._$AM.parentNode;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  u(t3) {
    let { el: { content: i5 }, parts: s4 } = this._$AD, e4 = (t3?.creationScope ?? r3).importNode(i5, !0);
    C.currentNode = e4;
    let h3 = C.nextNode(), o6 = 0, n5 = 0, l3 = s4[0];
    for (; l3 !== void 0; ) {
      if (o6 === l3.index) {
        let i6;
        l3.type === 2 ? i6 = new R(h3, h3.nextSibling, this, t3) : l3.type === 1 ? i6 = new l3.ctor(h3, l3.name, l3.strings, this, t3) : l3.type === 6 && (i6 = new z(h3, this, t3)), this._$AV.push(i6), l3 = s4[++n5];
      }
      o6 !== l3?.index && (h3 = C.nextNode(), o6++);
    }
    return C.currentNode = r3, e4;
  }
  p(t3) {
    let i5 = 0;
    for (let s4 of this._$AV) s4 !== void 0 && (s4.strings !== void 0 ? (s4._$AI(t3, s4, i5), i5 += s4.strings.length - 2) : s4._$AI(t3[i5])), i5++;
  }
}, R = class _R {
  get _$AU() {
    return this._$AM?._$AU ?? this._$Cv;
  }
  constructor(t3, i5, s4, e4) {
    this.type = 2, this._$AH = E, this._$AN = void 0, this._$AA = t3, this._$AB = i5, this._$AM = s4, this.options = e4, this._$Cv = e4?.isConnected ?? !0;
  }
  get parentNode() {
    let t3 = this._$AA.parentNode, i5 = this._$AM;
    return i5 !== void 0 && t3?.nodeType === 11 && (t3 = i5.parentNode), t3;
  }
  get startNode() {
    return this._$AA;
  }
  get endNode() {
    return this._$AB;
  }
  _$AI(t3, i5 = this) {
    t3 = S2(this, t3, i5), c3(t3) ? t3 === E || t3 == null || t3 === "" ? (this._$AH !== E && this._$AR(), this._$AH = E) : t3 !== this._$AH && t3 !== T && this._(t3) : t3._$litType$ !== void 0 ? this.$(t3) : t3.nodeType !== void 0 ? this.T(t3) : u2(t3) ? this.k(t3) : this._(t3);
  }
  O(t3) {
    return this._$AA.parentNode.insertBefore(t3, this._$AB);
  }
  T(t3) {
    this._$AH !== t3 && (this._$AR(), this._$AH = this.O(t3));
  }
  _(t3) {
    this._$AH !== E && c3(this._$AH) ? this._$AA.nextSibling.data = t3 : this.T(r3.createTextNode(t3)), this._$AH = t3;
  }
  $(t3) {
    let { values: i5, _$litType$: s4 } = t3, e4 = typeof s4 == "number" ? this._$AC(t3) : (s4.el === void 0 && (s4.el = N.createElement(P(s4.h, s4.h[0]), this.options)), s4);
    if (this._$AH?._$AD === e4) this._$AH.p(i5);
    else {
      let t4 = new M(e4, this), s5 = t4.u(this.options);
      t4.p(i5), this.T(s5), this._$AH = t4;
    }
  }
  _$AC(t3) {
    let i5 = A.get(t3.strings);
    return i5 === void 0 && A.set(t3.strings, i5 = new N(t3)), i5;
  }
  k(t3) {
    a2(this._$AH) || (this._$AH = [], this._$AR());
    let i5 = this._$AH, s4, e4 = 0;
    for (let h3 of t3) e4 === i5.length ? i5.push(s4 = new _R(this.O(l2()), this.O(l2()), this, this.options)) : s4 = i5[e4], s4._$AI(h3), e4++;
    e4 < i5.length && (this._$AR(s4 && s4._$AB.nextSibling, e4), i5.length = e4);
  }
  _$AR(t3 = this._$AA.nextSibling, i5) {
    for (this._$AP?.(!1, !0, i5); t3 !== this._$AB; ) {
      let i6 = t3.nextSibling;
      t3.remove(), t3 = i6;
    }
  }
  setConnected(t3) {
    this._$AM === void 0 && (this._$Cv = t3, this._$AP?.(t3));
  }
}, k = class {
  get tagName() {
    return this.element.tagName;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  constructor(t3, i5, s4, e4, h3) {
    this.type = 1, this._$AH = E, this._$AN = void 0, this.element = t3, this.name = i5, this._$AM = e4, this.options = h3, s4.length > 2 || s4[0] !== "" || s4[1] !== "" ? (this._$AH = Array(s4.length - 1).fill(new String()), this.strings = s4) : this._$AH = E;
  }
  _$AI(t3, i5 = this, s4, e4) {
    let h3 = this.strings, o6 = !1;
    if (h3 === void 0) t3 = S2(this, t3, i5, 0), o6 = !c3(t3) || t3 !== this._$AH && t3 !== T, o6 && (this._$AH = t3);
    else {
      let e5 = t3, n5, r4;
      for (t3 = h3[0], n5 = 0; n5 < h3.length - 1; n5++) r4 = S2(this, e5[s4 + n5], i5, n5), r4 === T && (r4 = this._$AH[n5]), o6 ||= !c3(r4) || r4 !== this._$AH[n5], r4 === E ? t3 = E : t3 !== E && (t3 += (r4 ?? "") + h3[n5 + 1]), this._$AH[n5] = r4;
    }
    o6 && !e4 && this.j(t3);
  }
  j(t3) {
    t3 === E ? this.element.removeAttribute(this.name) : this.element.setAttribute(this.name, t3 ?? "");
  }
}, H = class extends k {
  constructor() {
    super(...arguments), this.type = 3;
  }
  j(t3) {
    this.element[this.name] = t3 === E ? void 0 : t3;
  }
}, I = class extends k {
  constructor() {
    super(...arguments), this.type = 4;
  }
  j(t3) {
    this.element.toggleAttribute(this.name, !!t3 && t3 !== E);
  }
}, L = class extends k {
  constructor(t3, i5, s4, e4, h3) {
    super(t3, i5, s4, e4, h3), this.type = 5;
  }
  _$AI(t3, i5 = this) {
    if ((t3 = S2(this, t3, i5, 0) ?? E) === T) return;
    let s4 = this._$AH, e4 = t3 === E && s4 !== E || t3.capture !== s4.capture || t3.once !== s4.once || t3.passive !== s4.passive, h3 = t3 !== E && (s4 === E || e4);
    e4 && this.element.removeEventListener(this.name, this, s4), h3 && this.element.addEventListener(this.name, this, t3), this._$AH = t3;
  }
  handleEvent(t3) {
    typeof this._$AH == "function" ? this._$AH.call(this.options?.host ?? this.element, t3) : this._$AH.handleEvent(t3);
  }
}, z = class {
  constructor(t3, i5, s4) {
    this.element = t3, this.type = 6, this._$AN = void 0, this._$AM = i5, this.options = s4;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AI(t3) {
    S2(this, t3);
  }
}, Z = { M: e3, P: h2, A: o3, C: 1, L: V, R: M, D: u2, V: S2, I: R, H: k, N: I, U: L, B: H, F: z }, j = t2.litHtmlPolyfillSupport;
j?.(N, R), (t2.litHtmlVersions ??= []).push("3.3.1");
var B = (t3, i5, s4) => {
  let e4 = s4?.renderBefore ?? i5, h3 = e4._$litPart$;
  if (h3 === void 0) {
    let t4 = s4?.renderBefore ?? null;
    e4._$litPart$ = h3 = new R(i5.insertBefore(l2(), t4), t4, void 0, s4 ?? {});
  }
  return h3._$AI(t3), h3;
};

// node_modules/lit-element/lit-element.js
var s3 = globalThis, i4 = class extends y {
  constructor() {
    super(...arguments), this.renderOptions = { host: this }, this._$Do = void 0;
  }
  createRenderRoot() {
    let t3 = super.createRenderRoot();
    return this.renderOptions.renderBefore ??= t3.firstChild, t3;
  }
  update(t3) {
    let r4 = this.render();
    this.hasUpdated || (this.renderOptions.isConnected = this.isConnected), super.update(t3), this._$Do = B(r4, this.renderRoot, this.renderOptions);
  }
  connectedCallback() {
    super.connectedCallback(), this._$Do?.setConnected(!0);
  }
  disconnectedCallback() {
    super.disconnectedCallback(), this._$Do?.setConnected(!1);
  }
  render() {
    return T;
  }
};
i4._$litElement$ = !0, i4.finalized = !0, s3.litElementHydrateSupport?.({ LitElement: i4 });
var o4 = s3.litElementPolyfillSupport;
o4?.({ LitElement: i4 });
var n4 = { _$AK: (t3, e4, r4) => {
  t3._$AK(e4, r4);
}, _$AL: (t3) => t3._$AL };
(s3.litElementVersions ??= []).push("4.2.1");

// node_modules/lit-html/is-server.js
var o5 = !1;
export {
  n as CSSResult,
  i4 as LitElement,
  y as ReactiveElement,
  n4 as _$LE,
  Z as _$LH,
  S as adoptStyles,
  i as css,
  u as defaultConverter,
  c as getCompatibleStyle,
  x as html,
  o5 as isServer,
  w as mathml,
  T as noChange,
  f as notEqual,
  E as nothing,
  B as render,
  e as supportsAdoptingStyleSheets,
  b2 as svg,
  r as unsafeCSS
};
/*! Bundled license information:

@lit/reactive-element/css-tag.js:
  (**
   * @license
   * Copyright 2019 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

@lit/reactive-element/reactive-element.js:
lit-html/lit-html.js:
lit-element/lit-element.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/is-server.js:
  (**
   * @license
   * Copyright 2022 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=lit.js.map
