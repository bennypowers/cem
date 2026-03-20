// node_modules/lit-html/lit-html.js
var t = globalThis, i = t.trustedTypes, s = i ? i.createPolicy("lit-html", { createHTML: (t4) => t4 }) : void 0, e = "$lit$", h = `lit$${Math.random().toFixed(9).slice(2)}$`, o = "?" + h, n = `<${o}>`, r = document, l = () => r.createComment(""), c = (t4) => t4 === null || typeof t4 != "object" && typeof t4 != "function", a = Array.isArray, u = (t4) => a(t4) || typeof t4?.[Symbol.iterator] == "function", d = `[ 	
\f\r]`, f = /<(?:(!--|\/[^a-zA-Z])|(\/?[a-zA-Z][^>\s]*)|(\/?$))/g, v = /-->/g, _ = />/g, m = RegExp(`>|${d}(?:([^\\s"'>=/]+)(${d}*=${d}*(?:[^ 	
\f\r"'\`<>=]|("|')|))|$)`, "g"), p = /'/g, g = /"/g, $ = /^(?:script|style|textarea|title)$/i, y = (t4) => (i3, ...s2) => ({ _$litType$: t4, strings: i3, values: s2 }), x = y(1), b = y(2), w = y(3), T = Symbol.for("lit-noChange"), E = Symbol.for("lit-nothing"), A = /* @__PURE__ */ new WeakMap(), C = r.createTreeWalker(r, 129);
function P(t4, i3) {
  if (!a(t4) || !t4.hasOwnProperty("raw")) throw Error("invalid template strings array");
  return s !== void 0 ? s.createHTML(i3) : i3;
}
var V = (t4, i3) => {
  let s2 = t4.length - 1, o2 = [], r2, l3 = i3 === 2 ? "<svg>" : i3 === 3 ? "<math>" : "", c2 = f;
  for (let i4 = 0; i4 < s2; i4++) {
    let s3 = t4[i4], a2, u3, d2 = -1, y2 = 0;
    for (; y2 < s3.length && (c2.lastIndex = y2, u3 = c2.exec(s3), u3 !== null); ) y2 = c2.lastIndex, c2 === f ? u3[1] === "!--" ? c2 = v : u3[1] !== void 0 ? c2 = _ : u3[2] !== void 0 ? ($.test(u3[2]) && (r2 = RegExp("</" + u3[2], "g")), c2 = m) : u3[3] !== void 0 && (c2 = m) : c2 === m ? u3[0] === ">" ? (c2 = r2 ?? f, d2 = -1) : u3[1] === void 0 ? d2 = -2 : (d2 = c2.lastIndex - u3[2].length, a2 = u3[1], c2 = u3[3] === void 0 ? m : u3[3] === '"' ? g : p) : c2 === g || c2 === p ? c2 = m : c2 === v || c2 === _ ? c2 = f : (c2 = m, r2 = void 0);
    let x2 = c2 === m && t4[i4 + 1].startsWith("/>") ? " " : "";
    l3 += c2 === f ? s3 + n : d2 >= 0 ? (o2.push(a2), s3.slice(0, d2) + e + s3.slice(d2) + h + x2) : s3 + h + (d2 === -2 ? i4 : x2);
  }
  return [P(t4, l3 + (t4[s2] || "<?>") + (i3 === 2 ? "</svg>" : i3 === 3 ? "</math>" : "")), o2];
}, N = class _N {
  constructor({ strings: t4, _$litType$: s2 }, n2) {
    let r2;
    this.parts = [];
    let c2 = 0, a2 = 0, u3 = t4.length - 1, d2 = this.parts, [f3, v2] = V(t4, s2);
    if (this.el = _N.createElement(f3, n2), C.currentNode = this.el.content, s2 === 2 || s2 === 3) {
      let t5 = this.el.content.firstChild;
      t5.replaceWith(...t5.childNodes);
    }
    for (; (r2 = C.nextNode()) !== null && d2.length < u3; ) {
      if (r2.nodeType === 1) {
        if (r2.hasAttributes()) for (let t5 of r2.getAttributeNames()) if (t5.endsWith(e)) {
          let i3 = v2[a2++], s3 = r2.getAttribute(t5).split(h), e3 = /([.?@])?(.*)/.exec(i3);
          d2.push({ type: 1, index: c2, name: e3[2], strings: s3, ctor: e3[1] === "." ? H : e3[1] === "?" ? I : e3[1] === "@" ? L : k }), r2.removeAttribute(t5);
        } else t5.startsWith(h) && (d2.push({ type: 6, index: c2 }), r2.removeAttribute(t5));
        if ($.test(r2.tagName)) {
          let t5 = r2.textContent.split(h), s3 = t5.length - 1;
          if (s3 > 0) {
            r2.textContent = i ? i.emptyScript : "";
            for (let i3 = 0; i3 < s3; i3++) r2.append(t5[i3], l()), C.nextNode(), d2.push({ type: 2, index: ++c2 });
            r2.append(t5[s3], l());
          }
        }
      } else if (r2.nodeType === 8) if (r2.data === o) d2.push({ type: 2, index: c2 });
      else {
        let t5 = -1;
        for (; (t5 = r2.data.indexOf(h, t5 + 1)) !== -1; ) d2.push({ type: 7, index: c2 }), t5 += h.length - 1;
      }
      c2++;
    }
  }
  static createElement(t4, i3) {
    let s2 = r.createElement("template");
    return s2.innerHTML = t4, s2;
  }
};
function S(t4, i3, s2 = t4, e3) {
  if (i3 === T) return i3;
  let h2 = e3 !== void 0 ? s2._$Co?.[e3] : s2._$Cl, o2 = c(i3) ? void 0 : i3._$litDirective$;
  return h2?.constructor !== o2 && (h2?._$AO?.(!1), o2 === void 0 ? h2 = void 0 : (h2 = new o2(t4), h2._$AT(t4, s2, e3)), e3 !== void 0 ? (s2._$Co ??= [])[e3] = h2 : s2._$Cl = h2), h2 !== void 0 && (i3 = S(t4, h2._$AS(t4, i3.values), h2, e3)), i3;
}
var M = class {
  constructor(t4, i3) {
    this._$AV = [], this._$AN = void 0, this._$AD = t4, this._$AM = i3;
  }
  get parentNode() {
    return this._$AM.parentNode;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  u(t4) {
    let { el: { content: i3 }, parts: s2 } = this._$AD, e3 = (t4?.creationScope ?? r).importNode(i3, !0);
    C.currentNode = e3;
    let h2 = C.nextNode(), o2 = 0, n2 = 0, l3 = s2[0];
    for (; l3 !== void 0; ) {
      if (o2 === l3.index) {
        let i4;
        l3.type === 2 ? i4 = new R(h2, h2.nextSibling, this, t4) : l3.type === 1 ? i4 = new l3.ctor(h2, l3.name, l3.strings, this, t4) : l3.type === 6 && (i4 = new z(h2, this, t4)), this._$AV.push(i4), l3 = s2[++n2];
      }
      o2 !== l3?.index && (h2 = C.nextNode(), o2++);
    }
    return C.currentNode = r, e3;
  }
  p(t4) {
    let i3 = 0;
    for (let s2 of this._$AV) s2 !== void 0 && (s2.strings !== void 0 ? (s2._$AI(t4, s2, i3), i3 += s2.strings.length - 2) : s2._$AI(t4[i3])), i3++;
  }
}, R = class _R {
  get _$AU() {
    return this._$AM?._$AU ?? this._$Cv;
  }
  constructor(t4, i3, s2, e3) {
    this.type = 2, this._$AH = E, this._$AN = void 0, this._$AA = t4, this._$AB = i3, this._$AM = s2, this.options = e3, this._$Cv = e3?.isConnected ?? !0;
  }
  get parentNode() {
    let t4 = this._$AA.parentNode, i3 = this._$AM;
    return i3 !== void 0 && t4?.nodeType === 11 && (t4 = i3.parentNode), t4;
  }
  get startNode() {
    return this._$AA;
  }
  get endNode() {
    return this._$AB;
  }
  _$AI(t4, i3 = this) {
    t4 = S(this, t4, i3), c(t4) ? t4 === E || t4 == null || t4 === "" ? (this._$AH !== E && this._$AR(), this._$AH = E) : t4 !== this._$AH && t4 !== T && this._(t4) : t4._$litType$ !== void 0 ? this.$(t4) : t4.nodeType !== void 0 ? this.T(t4) : u(t4) ? this.k(t4) : this._(t4);
  }
  O(t4) {
    return this._$AA.parentNode.insertBefore(t4, this._$AB);
  }
  T(t4) {
    this._$AH !== t4 && (this._$AR(), this._$AH = this.O(t4));
  }
  _(t4) {
    this._$AH !== E && c(this._$AH) ? this._$AA.nextSibling.data = t4 : this.T(r.createTextNode(t4)), this._$AH = t4;
  }
  $(t4) {
    let { values: i3, _$litType$: s2 } = t4, e3 = typeof s2 == "number" ? this._$AC(t4) : (s2.el === void 0 && (s2.el = N.createElement(P(s2.h, s2.h[0]), this.options)), s2);
    if (this._$AH?._$AD === e3) this._$AH.p(i3);
    else {
      let t5 = new M(e3, this), s3 = t5.u(this.options);
      t5.p(i3), this.T(s3), this._$AH = t5;
    }
  }
  _$AC(t4) {
    let i3 = A.get(t4.strings);
    return i3 === void 0 && A.set(t4.strings, i3 = new N(t4)), i3;
  }
  k(t4) {
    a(this._$AH) || (this._$AH = [], this._$AR());
    let i3 = this._$AH, s2, e3 = 0;
    for (let h2 of t4) e3 === i3.length ? i3.push(s2 = new _R(this.O(l()), this.O(l()), this, this.options)) : s2 = i3[e3], s2._$AI(h2), e3++;
    e3 < i3.length && (this._$AR(s2 && s2._$AB.nextSibling, e3), i3.length = e3);
  }
  _$AR(t4 = this._$AA.nextSibling, i3) {
    for (this._$AP?.(!1, !0, i3); t4 !== this._$AB; ) {
      let i4 = t4.nextSibling;
      t4.remove(), t4 = i4;
    }
  }
  setConnected(t4) {
    this._$AM === void 0 && (this._$Cv = t4, this._$AP?.(t4));
  }
}, k = class {
  get tagName() {
    return this.element.tagName;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  constructor(t4, i3, s2, e3, h2) {
    this.type = 1, this._$AH = E, this._$AN = void 0, this.element = t4, this.name = i3, this._$AM = e3, this.options = h2, s2.length > 2 || s2[0] !== "" || s2[1] !== "" ? (this._$AH = Array(s2.length - 1).fill(new String()), this.strings = s2) : this._$AH = E;
  }
  _$AI(t4, i3 = this, s2, e3) {
    let h2 = this.strings, o2 = !1;
    if (h2 === void 0) t4 = S(this, t4, i3, 0), o2 = !c(t4) || t4 !== this._$AH && t4 !== T, o2 && (this._$AH = t4);
    else {
      let e4 = t4, n2, r2;
      for (t4 = h2[0], n2 = 0; n2 < h2.length - 1; n2++) r2 = S(this, e4[s2 + n2], i3, n2), r2 === T && (r2 = this._$AH[n2]), o2 ||= !c(r2) || r2 !== this._$AH[n2], r2 === E ? t4 = E : t4 !== E && (t4 += (r2 ?? "") + h2[n2 + 1]), this._$AH[n2] = r2;
    }
    o2 && !e3 && this.j(t4);
  }
  j(t4) {
    t4 === E ? this.element.removeAttribute(this.name) : this.element.setAttribute(this.name, t4 ?? "");
  }
}, H = class extends k {
  constructor() {
    super(...arguments), this.type = 3;
  }
  j(t4) {
    this.element[this.name] = t4 === E ? void 0 : t4;
  }
}, I = class extends k {
  constructor() {
    super(...arguments), this.type = 4;
  }
  j(t4) {
    this.element.toggleAttribute(this.name, !!t4 && t4 !== E);
  }
}, L = class extends k {
  constructor(t4, i3, s2, e3, h2) {
    super(t4, i3, s2, e3, h2), this.type = 5;
  }
  _$AI(t4, i3 = this) {
    if ((t4 = S(this, t4, i3, 0) ?? E) === T) return;
    let s2 = this._$AH, e3 = t4 === E && s2 !== E || t4.capture !== s2.capture || t4.once !== s2.once || t4.passive !== s2.passive, h2 = t4 !== E && (s2 === E || e3);
    e3 && this.element.removeEventListener(this.name, this, s2), h2 && this.element.addEventListener(this.name, this, t4), this._$AH = t4;
  }
  handleEvent(t4) {
    typeof this._$AH == "function" ? this._$AH.call(this.options?.host ?? this.element, t4) : this._$AH.handleEvent(t4);
  }
}, z = class {
  constructor(t4, i3, s2) {
    this.element = t4, this.type = 6, this._$AN = void 0, this._$AM = i3, this.options = s2;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AI(t4) {
    S(this, t4);
  }
}, Z = { M: e, P: h, A: o, C: 1, L: V, R: M, D: u, V: S, I: R, H: k, N: I, U: L, B: H, F: z }, j = t.litHtmlPolyfillSupport;
j?.(N, R), (t.litHtmlVersions ??= []).push("3.3.1");

// node_modules/lit-html/directive.js
var t2 = { ATTRIBUTE: 1, CHILD: 2, PROPERTY: 3, BOOLEAN_ATTRIBUTE: 4, EVENT: 5, ELEMENT: 6 }, e2 = (t4) => (...e3) => ({ _$litDirective$: t4, values: e3 }), i2 = class {
  constructor(t4) {
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AT(t4, e3, i3) {
    this._$Ct = t4, this._$AM = e3, this._$Ci = i3;
  }
  _$AS(t4, e3) {
    return this.update(t4, e3);
  }
  update(t4, e3) {
    return this.render(...e3);
  }
};

// node_modules/lit-html/directive-helpers.js
var { I: t3 } = Z;
var f2 = (o2) => o2.strings === void 0;
var u2 = {}, m2 = (o2, t4 = u2) => o2._$AH = t4;

// node_modules/lit-html/directives/live.js
var l2 = e2(class extends i2 {
  constructor(r2) {
    if (super(r2), r2.type !== t2.PROPERTY && r2.type !== t2.ATTRIBUTE && r2.type !== t2.BOOLEAN_ATTRIBUTE) throw Error("The `live` directive is not allowed on child or event bindings");
    if (!f2(r2)) throw Error("`live` bindings can only contain a single expression");
  }
  render(r2) {
    return r2;
  }
  update(i3, [t4]) {
    if (t4 === T || t4 === E) return t4;
    let o2 = i3.element, l3 = i3.name;
    if (i3.type === t2.PROPERTY) {
      if (t4 === o2[l3]) return T;
    } else if (i3.type === t2.BOOLEAN_ATTRIBUTE) {
      if (!!t4 === o2.hasAttribute(l3)) return T;
    } else if (i3.type === t2.ATTRIBUTE && o2.getAttribute(l3) === t4 + "") return T;
    return m2(i3), t4;
  }
});
export {
  l2 as live
};
/*! Bundled license information:

lit-html/lit-html.js:
lit-html/directive.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/directive-helpers.js:
lit-html/directives/live.js:
  (**
   * @license
   * Copyright 2020 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=live.js.map
