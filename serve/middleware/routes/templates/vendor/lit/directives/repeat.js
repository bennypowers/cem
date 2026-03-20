// node_modules/lit-html/lit-html.js
var t = globalThis, i = t.trustedTypes, s = i ? i.createPolicy("lit-html", { createHTML: (t4) => t4 }) : void 0, e = "$lit$", h = `lit$${Math.random().toFixed(9).slice(2)}$`, o = "?" + h, n = `<${o}>`, r = document, l = () => r.createComment(""), c = (t4) => t4 === null || typeof t4 != "object" && typeof t4 != "function", a = Array.isArray, u = (t4) => a(t4) || typeof t4?.[Symbol.iterator] == "function", d = `[ 	
\f\r]`, f = /<(?:(!--|\/[^a-zA-Z])|(\/?[a-zA-Z][^>\s]*)|(\/?$))/g, v = /-->/g, _ = />/g, m = RegExp(`>|${d}(?:([^\\s"'>=/]+)(${d}*=${d}*(?:[^ 	
\f\r"'\`<>=]|("|')|))|$)`, "g"), p = /'/g, g = /"/g, $ = /^(?:script|style|textarea|title)$/i, y = (t4) => (i3, ...s3) => ({ _$litType$: t4, strings: i3, values: s3 }), x = y(1), b = y(2), w = y(3), T = Symbol.for("lit-noChange"), E = Symbol.for("lit-nothing"), A = /* @__PURE__ */ new WeakMap(), C = r.createTreeWalker(r, 129);
function P(t4, i3) {
  if (!a(t4) || !t4.hasOwnProperty("raw")) throw Error("invalid template strings array");
  return s !== void 0 ? s.createHTML(i3) : i3;
}
var V = (t4, i3) => {
  let s3 = t4.length - 1, o2 = [], r3, l2 = i3 === 2 ? "<svg>" : i3 === 3 ? "<math>" : "", c3 = f;
  for (let i4 = 0; i4 < s3; i4++) {
    let s4 = t4[i4], a2, u4, d2 = -1, y2 = 0;
    for (; y2 < s4.length && (c3.lastIndex = y2, u4 = c3.exec(s4), u4 !== null); ) y2 = c3.lastIndex, c3 === f ? u4[1] === "!--" ? c3 = v : u4[1] !== void 0 ? c3 = _ : u4[2] !== void 0 ? ($.test(u4[2]) && (r3 = RegExp("</" + u4[2], "g")), c3 = m) : u4[3] !== void 0 && (c3 = m) : c3 === m ? u4[0] === ">" ? (c3 = r3 ?? f, d2 = -1) : u4[1] === void 0 ? d2 = -2 : (d2 = c3.lastIndex - u4[2].length, a2 = u4[1], c3 = u4[3] === void 0 ? m : u4[3] === '"' ? g : p) : c3 === g || c3 === p ? c3 = m : c3 === v || c3 === _ ? c3 = f : (c3 = m, r3 = void 0);
    let x2 = c3 === m && t4[i4 + 1].startsWith("/>") ? " " : "";
    l2 += c3 === f ? s4 + n : d2 >= 0 ? (o2.push(a2), s4.slice(0, d2) + e + s4.slice(d2) + h + x2) : s4 + h + (d2 === -2 ? i4 : x2);
  }
  return [P(t4, l2 + (t4[s3] || "<?>") + (i3 === 2 ? "</svg>" : i3 === 3 ? "</math>" : "")), o2];
}, N = class _N {
  constructor({ strings: t4, _$litType$: s3 }, n2) {
    let r3;
    this.parts = [];
    let c3 = 0, a2 = 0, u4 = t4.length - 1, d2 = this.parts, [f2, v3] = V(t4, s3);
    if (this.el = _N.createElement(f2, n2), C.currentNode = this.el.content, s3 === 2 || s3 === 3) {
      let t5 = this.el.content.firstChild;
      t5.replaceWith(...t5.childNodes);
    }
    for (; (r3 = C.nextNode()) !== null && d2.length < u4; ) {
      if (r3.nodeType === 1) {
        if (r3.hasAttributes()) for (let t5 of r3.getAttributeNames()) if (t5.endsWith(e)) {
          let i3 = v3[a2++], s4 = r3.getAttribute(t5).split(h), e3 = /([.?@])?(.*)/.exec(i3);
          d2.push({ type: 1, index: c3, name: e3[2], strings: s4, ctor: e3[1] === "." ? H : e3[1] === "?" ? I : e3[1] === "@" ? L : k }), r3.removeAttribute(t5);
        } else t5.startsWith(h) && (d2.push({ type: 6, index: c3 }), r3.removeAttribute(t5));
        if ($.test(r3.tagName)) {
          let t5 = r3.textContent.split(h), s4 = t5.length - 1;
          if (s4 > 0) {
            r3.textContent = i ? i.emptyScript : "";
            for (let i3 = 0; i3 < s4; i3++) r3.append(t5[i3], l()), C.nextNode(), d2.push({ type: 2, index: ++c3 });
            r3.append(t5[s4], l());
          }
        }
      } else if (r3.nodeType === 8) if (r3.data === o) d2.push({ type: 2, index: c3 });
      else {
        let t5 = -1;
        for (; (t5 = r3.data.indexOf(h, t5 + 1)) !== -1; ) d2.push({ type: 7, index: c3 }), t5 += h.length - 1;
      }
      c3++;
    }
  }
  static createElement(t4, i3) {
    let s3 = r.createElement("template");
    return s3.innerHTML = t4, s3;
  }
};
function S(t4, i3, s3 = t4, e3) {
  if (i3 === T) return i3;
  let h2 = e3 !== void 0 ? s3._$Co?.[e3] : s3._$Cl, o2 = c(i3) ? void 0 : i3._$litDirective$;
  return h2?.constructor !== o2 && (h2?._$AO?.(!1), o2 === void 0 ? h2 = void 0 : (h2 = new o2(t4), h2._$AT(t4, s3, e3)), e3 !== void 0 ? (s3._$Co ??= [])[e3] = h2 : s3._$Cl = h2), h2 !== void 0 && (i3 = S(t4, h2._$AS(t4, i3.values), h2, e3)), i3;
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
    let { el: { content: i3 }, parts: s3 } = this._$AD, e3 = (t4?.creationScope ?? r).importNode(i3, !0);
    C.currentNode = e3;
    let h2 = C.nextNode(), o2 = 0, n2 = 0, l2 = s3[0];
    for (; l2 !== void 0; ) {
      if (o2 === l2.index) {
        let i4;
        l2.type === 2 ? i4 = new R(h2, h2.nextSibling, this, t4) : l2.type === 1 ? i4 = new l2.ctor(h2, l2.name, l2.strings, this, t4) : l2.type === 6 && (i4 = new z(h2, this, t4)), this._$AV.push(i4), l2 = s3[++n2];
      }
      o2 !== l2?.index && (h2 = C.nextNode(), o2++);
    }
    return C.currentNode = r, e3;
  }
  p(t4) {
    let i3 = 0;
    for (let s3 of this._$AV) s3 !== void 0 && (s3.strings !== void 0 ? (s3._$AI(t4, s3, i3), i3 += s3.strings.length - 2) : s3._$AI(t4[i3])), i3++;
  }
}, R = class _R {
  get _$AU() {
    return this._$AM?._$AU ?? this._$Cv;
  }
  constructor(t4, i3, s3, e3) {
    this.type = 2, this._$AH = E, this._$AN = void 0, this._$AA = t4, this._$AB = i3, this._$AM = s3, this.options = e3, this._$Cv = e3?.isConnected ?? !0;
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
    let { values: i3, _$litType$: s3 } = t4, e3 = typeof s3 == "number" ? this._$AC(t4) : (s3.el === void 0 && (s3.el = N.createElement(P(s3.h, s3.h[0]), this.options)), s3);
    if (this._$AH?._$AD === e3) this._$AH.p(i3);
    else {
      let t5 = new M(e3, this), s4 = t5.u(this.options);
      t5.p(i3), this.T(s4), this._$AH = t5;
    }
  }
  _$AC(t4) {
    let i3 = A.get(t4.strings);
    return i3 === void 0 && A.set(t4.strings, i3 = new N(t4)), i3;
  }
  k(t4) {
    a(this._$AH) || (this._$AH = [], this._$AR());
    let i3 = this._$AH, s3, e3 = 0;
    for (let h2 of t4) e3 === i3.length ? i3.push(s3 = new _R(this.O(l()), this.O(l()), this, this.options)) : s3 = i3[e3], s3._$AI(h2), e3++;
    e3 < i3.length && (this._$AR(s3 && s3._$AB.nextSibling, e3), i3.length = e3);
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
  constructor(t4, i3, s3, e3, h2) {
    this.type = 1, this._$AH = E, this._$AN = void 0, this.element = t4, this.name = i3, this._$AM = e3, this.options = h2, s3.length > 2 || s3[0] !== "" || s3[1] !== "" ? (this._$AH = Array(s3.length - 1).fill(new String()), this.strings = s3) : this._$AH = E;
  }
  _$AI(t4, i3 = this, s3, e3) {
    let h2 = this.strings, o2 = !1;
    if (h2 === void 0) t4 = S(this, t4, i3, 0), o2 = !c(t4) || t4 !== this._$AH && t4 !== T, o2 && (this._$AH = t4);
    else {
      let e4 = t4, n2, r3;
      for (t4 = h2[0], n2 = 0; n2 < h2.length - 1; n2++) r3 = S(this, e4[s3 + n2], i3, n2), r3 === T && (r3 = this._$AH[n2]), o2 ||= !c(r3) || r3 !== this._$AH[n2], r3 === E ? t4 = E : t4 !== E && (t4 += (r3 ?? "") + h2[n2 + 1]), this._$AH[n2] = r3;
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
  constructor(t4, i3, s3, e3, h2) {
    super(t4, i3, s3, e3, h2), this.type = 5;
  }
  _$AI(t4, i3 = this) {
    if ((t4 = S(this, t4, i3, 0) ?? E) === T) return;
    let s3 = this._$AH, e3 = t4 === E && s3 !== E || t4.capture !== s3.capture || t4.once !== s3.once || t4.passive !== s3.passive, h2 = t4 !== E && (s3 === E || e3);
    e3 && this.element.removeEventListener(this.name, this, s3), h2 && this.element.addEventListener(this.name, this, t4), this._$AH = t4;
  }
  handleEvent(t4) {
    typeof this._$AH == "function" ? this._$AH.call(this.options?.host ?? this.element, t4) : this._$AH.handleEvent(t4);
  }
}, z = class {
  constructor(t4, i3, s3) {
    this.element = t4, this.type = 6, this._$AN = void 0, this._$AM = i3, this.options = s3;
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
var r2 = () => document.createComment(""), s2 = (o2, i3, n2) => {
  let e3 = o2._$AA.parentNode, l2 = i3 === void 0 ? o2._$AB : i3._$AA;
  if (n2 === void 0) {
    let i4 = e3.insertBefore(r2(), l2), d2 = e3.insertBefore(r2(), l2);
    n2 = new t3(i4, d2, o2, o2.options);
  } else {
    let t4 = n2._$AB.nextSibling, i4 = n2._$AM, d2 = i4 !== o2;
    if (d2) {
      let t5;
      n2._$AQ?.(o2), n2._$AM = o2, n2._$AP !== void 0 && (t5 = o2._$AU) !== i4._$AU && n2._$AP(t5);
    }
    if (t4 !== l2 || d2) {
      let o3 = n2._$AA;
      for (; o3 !== t4; ) {
        let t5 = o3.nextSibling;
        e3.insertBefore(o3, l2), o3 = t5;
      }
    }
  }
  return n2;
}, v2 = (o2, t4, i3 = o2) => (o2._$AI(t4, i3), o2), u2 = {}, m2 = (o2, t4 = u2) => o2._$AH = t4, p2 = (o2) => o2._$AH, M2 = (o2) => {
  o2._$AR(), o2._$AA.remove();
};

// node_modules/lit-html/directives/repeat.js
var u3 = (e3, s3, t4) => {
  let r3 = /* @__PURE__ */ new Map();
  for (let l2 = s3; l2 <= t4; l2++) r3.set(e3[l2], l2);
  return r3;
}, c2 = e2(class extends i2 {
  constructor(e3) {
    if (super(e3), e3.type !== t2.CHILD) throw Error("repeat() can only be used in text expressions");
  }
  dt(e3, s3, t4) {
    let r3;
    t4 === void 0 ? t4 = s3 : s3 !== void 0 && (r3 = s3);
    let l2 = [], o2 = [], i3 = 0;
    for (let s4 of e3) l2[i3] = r3 ? r3(s4, i3) : i3, o2[i3] = t4(s4, i3), i3++;
    return { values: o2, keys: l2 };
  }
  render(e3, s3, t4) {
    return this.dt(e3, s3, t4).values;
  }
  update(s3, [t4, r3, c3]) {
    let d2 = p2(s3), { values: p3, keys: a2 } = this.dt(t4, r3, c3);
    if (!Array.isArray(d2)) return this.ut = a2, p3;
    let h2 = this.ut ??= [], v3 = [], m3, y2, x2 = 0, j2 = d2.length - 1, k2 = 0, w2 = p3.length - 1;
    for (; x2 <= j2 && k2 <= w2; ) if (d2[x2] === null) x2++;
    else if (d2[j2] === null) j2--;
    else if (h2[x2] === a2[k2]) v3[k2] = v2(d2[x2], p3[k2]), x2++, k2++;
    else if (h2[j2] === a2[w2]) v3[w2] = v2(d2[j2], p3[w2]), j2--, w2--;
    else if (h2[x2] === a2[w2]) v3[w2] = v2(d2[x2], p3[w2]), s2(s3, v3[w2 + 1], d2[x2]), x2++, w2--;
    else if (h2[j2] === a2[k2]) v3[k2] = v2(d2[j2], p3[k2]), s2(s3, d2[x2], d2[j2]), j2--, k2++;
    else if (m3 === void 0 && (m3 = u3(a2, k2, w2), y2 = u3(h2, x2, j2)), m3.has(h2[x2])) if (m3.has(h2[j2])) {
      let e3 = y2.get(a2[k2]), t5 = e3 !== void 0 ? d2[e3] : null;
      if (t5 === null) {
        let e4 = s2(s3, d2[x2]);
        v2(e4, p3[k2]), v3[k2] = e4;
      } else v3[k2] = v2(t5, p3[k2]), s2(s3, d2[x2], t5), d2[e3] = null;
      k2++;
    } else M2(d2[j2]), j2--;
    else M2(d2[x2]), x2++;
    for (; k2 <= w2; ) {
      let e3 = s2(s3, v3[w2 + 1]);
      v2(e3, p3[k2]), v3[k2++] = e3;
    }
    for (; x2 <= j2; ) {
      let e3 = d2[x2++];
      e3 !== null && M2(e3);
    }
    return this.ut = a2, m2(s3, v3), T;
  }
});
export {
  c2 as repeat
};
/*! Bundled license information:

lit-html/lit-html.js:
lit-html/directive.js:
lit-html/directives/repeat.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/directive-helpers.js:
  (**
   * @license
   * Copyright 2020 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=repeat.js.map
