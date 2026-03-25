// node_modules/lit-html/lit-html.js
var t = globalThis, i = t.trustedTypes, s = i ? i.createPolicy("lit-html", { createHTML: (t3) => t3 }) : void 0, e = "$lit$", h = `lit$${Math.random().toFixed(9).slice(2)}$`, o = "?" + h, n = `<${o}>`, r = document, l = () => r.createComment(""), c = (t3) => t3 === null || typeof t3 != "object" && typeof t3 != "function", a = Array.isArray, u = (t3) => a(t3) || typeof t3?.[Symbol.iterator] == "function", d = `[ 	
\f\r]`, f = /<(?:(!--|\/[^a-zA-Z])|(\/?[a-zA-Z][^>\s]*)|(\/?$))/g, v = /-->/g, _ = />/g, m = RegExp(`>|${d}(?:([^\\s"'>=/]+)(${d}*=${d}*(?:[^ 	
\f\r"'\`<>=]|("|')|))|$)`, "g"), p = /'/g, g = /"/g, $ = /^(?:script|style|textarea|title)$/i, y = (t3) => (i4, ...s2) => ({ _$litType$: t3, strings: i4, values: s2 }), x = y(1), b = y(2), w = y(3), T = /* @__PURE__ */ Symbol.for("lit-noChange"), E = /* @__PURE__ */ Symbol.for("lit-nothing"), A = /* @__PURE__ */ new WeakMap(), C = r.createTreeWalker(r, 129);
function P(t3, i4) {
  if (!a(t3) || !t3.hasOwnProperty("raw")) throw Error("invalid template strings array");
  return s !== void 0 ? s.createHTML(i4) : i4;
}
var V = (t3, i4) => {
  let s2 = t3.length - 1, o3 = [], r2, l2 = i4 === 2 ? "<svg>" : i4 === 3 ? "<math>" : "", c2 = f;
  for (let i5 = 0; i5 < s2; i5++) {
    let s3 = t3[i5], a2, u2, d2 = -1, y2 = 0;
    for (; y2 < s3.length && (c2.lastIndex = y2, u2 = c2.exec(s3), u2 !== null); ) y2 = c2.lastIndex, c2 === f ? u2[1] === "!--" ? c2 = v : u2[1] !== void 0 ? c2 = _ : u2[2] !== void 0 ? ($.test(u2[2]) && (r2 = RegExp("</" + u2[2], "g")), c2 = m) : u2[3] !== void 0 && (c2 = m) : c2 === m ? u2[0] === ">" ? (c2 = r2 ?? f, d2 = -1) : u2[1] === void 0 ? d2 = -2 : (d2 = c2.lastIndex - u2[2].length, a2 = u2[1], c2 = u2[3] === void 0 ? m : u2[3] === '"' ? g : p) : c2 === g || c2 === p ? c2 = m : c2 === v || c2 === _ ? c2 = f : (c2 = m, r2 = void 0);
    let x2 = c2 === m && t3[i5 + 1].startsWith("/>") ? " " : "";
    l2 += c2 === f ? s3 + n : d2 >= 0 ? (o3.push(a2), s3.slice(0, d2) + e + s3.slice(d2) + h + x2) : s3 + h + (d2 === -2 ? i5 : x2);
  }
  return [P(t3, l2 + (t3[s2] || "<?>") + (i4 === 2 ? "</svg>" : i4 === 3 ? "</math>" : "")), o3];
}, N = class _N {
  constructor({ strings: t3, _$litType$: s2 }, n3) {
    let r2;
    this.parts = [];
    let c2 = 0, a2 = 0, u2 = t3.length - 1, d2 = this.parts, [f2, v2] = V(t3, s2);
    if (this.el = _N.createElement(f2, n3), C.currentNode = this.el.content, s2 === 2 || s2 === 3) {
      let t4 = this.el.content.firstChild;
      t4.replaceWith(...t4.childNodes);
    }
    for (; (r2 = C.nextNode()) !== null && d2.length < u2; ) {
      if (r2.nodeType === 1) {
        if (r2.hasAttributes()) for (let t4 of r2.getAttributeNames()) if (t4.endsWith(e)) {
          let i4 = v2[a2++], s3 = r2.getAttribute(t4).split(h), e3 = /([.?@])?(.*)/.exec(i4);
          d2.push({ type: 1, index: c2, name: e3[2], strings: s3, ctor: e3[1] === "." ? H : e3[1] === "?" ? I : e3[1] === "@" ? L : k }), r2.removeAttribute(t4);
        } else t4.startsWith(h) && (d2.push({ type: 6, index: c2 }), r2.removeAttribute(t4));
        if ($.test(r2.tagName)) {
          let t4 = r2.textContent.split(h), s3 = t4.length - 1;
          if (s3 > 0) {
            r2.textContent = i ? i.emptyScript : "";
            for (let i4 = 0; i4 < s3; i4++) r2.append(t4[i4], l()), C.nextNode(), d2.push({ type: 2, index: ++c2 });
            r2.append(t4[s3], l());
          }
        }
      } else if (r2.nodeType === 8) if (r2.data === o) d2.push({ type: 2, index: c2 });
      else {
        let t4 = -1;
        for (; (t4 = r2.data.indexOf(h, t4 + 1)) !== -1; ) d2.push({ type: 7, index: c2 }), t4 += h.length - 1;
      }
      c2++;
    }
  }
  static createElement(t3, i4) {
    let s2 = r.createElement("template");
    return s2.innerHTML = t3, s2;
  }
};
function S(t3, i4, s2 = t3, e3) {
  if (i4 === T) return i4;
  let h2 = e3 !== void 0 ? s2._$Co?.[e3] : s2._$Cl, o3 = c(i4) ? void 0 : i4._$litDirective$;
  return h2?.constructor !== o3 && (h2?._$AO?.(!1), o3 === void 0 ? h2 = void 0 : (h2 = new o3(t3), h2._$AT(t3, s2, e3)), e3 !== void 0 ? (s2._$Co ??= [])[e3] = h2 : s2._$Cl = h2), h2 !== void 0 && (i4 = S(t3, h2._$AS(t3, i4.values), h2, e3)), i4;
}
var M = class {
  constructor(t3, i4) {
    this._$AV = [], this._$AN = void 0, this._$AD = t3, this._$AM = i4;
  }
  get parentNode() {
    return this._$AM.parentNode;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  u(t3) {
    let { el: { content: i4 }, parts: s2 } = this._$AD, e3 = (t3?.creationScope ?? r).importNode(i4, !0);
    C.currentNode = e3;
    let h2 = C.nextNode(), o3 = 0, n3 = 0, l2 = s2[0];
    for (; l2 !== void 0; ) {
      if (o3 === l2.index) {
        let i5;
        l2.type === 2 ? i5 = new R(h2, h2.nextSibling, this, t3) : l2.type === 1 ? i5 = new l2.ctor(h2, l2.name, l2.strings, this, t3) : l2.type === 6 && (i5 = new z(h2, this, t3)), this._$AV.push(i5), l2 = s2[++n3];
      }
      o3 !== l2?.index && (h2 = C.nextNode(), o3++);
    }
    return C.currentNode = r, e3;
  }
  p(t3) {
    let i4 = 0;
    for (let s2 of this._$AV) s2 !== void 0 && (s2.strings !== void 0 ? (s2._$AI(t3, s2, i4), i4 += s2.strings.length - 2) : s2._$AI(t3[i4])), i4++;
  }
}, R = class _R {
  get _$AU() {
    return this._$AM?._$AU ?? this._$Cv;
  }
  constructor(t3, i4, s2, e3) {
    this.type = 2, this._$AH = E, this._$AN = void 0, this._$AA = t3, this._$AB = i4, this._$AM = s2, this.options = e3, this._$Cv = e3?.isConnected ?? !0;
  }
  get parentNode() {
    let t3 = this._$AA.parentNode, i4 = this._$AM;
    return i4 !== void 0 && t3?.nodeType === 11 && (t3 = i4.parentNode), t3;
  }
  get startNode() {
    return this._$AA;
  }
  get endNode() {
    return this._$AB;
  }
  _$AI(t3, i4 = this) {
    t3 = S(this, t3, i4), c(t3) ? t3 === E || t3 == null || t3 === "" ? (this._$AH !== E && this._$AR(), this._$AH = E) : t3 !== this._$AH && t3 !== T && this._(t3) : t3._$litType$ !== void 0 ? this.$(t3) : t3.nodeType !== void 0 ? this.T(t3) : u(t3) ? this.k(t3) : this._(t3);
  }
  O(t3) {
    return this._$AA.parentNode.insertBefore(t3, this._$AB);
  }
  T(t3) {
    this._$AH !== t3 && (this._$AR(), this._$AH = this.O(t3));
  }
  _(t3) {
    this._$AH !== E && c(this._$AH) ? this._$AA.nextSibling.data = t3 : this.T(r.createTextNode(t3)), this._$AH = t3;
  }
  $(t3) {
    let { values: i4, _$litType$: s2 } = t3, e3 = typeof s2 == "number" ? this._$AC(t3) : (s2.el === void 0 && (s2.el = N.createElement(P(s2.h, s2.h[0]), this.options)), s2);
    if (this._$AH?._$AD === e3) this._$AH.p(i4);
    else {
      let t4 = new M(e3, this), s3 = t4.u(this.options);
      t4.p(i4), this.T(s3), this._$AH = t4;
    }
  }
  _$AC(t3) {
    let i4 = A.get(t3.strings);
    return i4 === void 0 && A.set(t3.strings, i4 = new N(t3)), i4;
  }
  k(t3) {
    a(this._$AH) || (this._$AH = [], this._$AR());
    let i4 = this._$AH, s2, e3 = 0;
    for (let h2 of t3) e3 === i4.length ? i4.push(s2 = new _R(this.O(l()), this.O(l()), this, this.options)) : s2 = i4[e3], s2._$AI(h2), e3++;
    e3 < i4.length && (this._$AR(s2 && s2._$AB.nextSibling, e3), i4.length = e3);
  }
  _$AR(t3 = this._$AA.nextSibling, i4) {
    for (this._$AP?.(!1, !0, i4); t3 !== this._$AB; ) {
      let i5 = t3.nextSibling;
      t3.remove(), t3 = i5;
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
  constructor(t3, i4, s2, e3, h2) {
    this.type = 1, this._$AH = E, this._$AN = void 0, this.element = t3, this.name = i4, this._$AM = e3, this.options = h2, s2.length > 2 || s2[0] !== "" || s2[1] !== "" ? (this._$AH = Array(s2.length - 1).fill(new String()), this.strings = s2) : this._$AH = E;
  }
  _$AI(t3, i4 = this, s2, e3) {
    let h2 = this.strings, o3 = !1;
    if (h2 === void 0) t3 = S(this, t3, i4, 0), o3 = !c(t3) || t3 !== this._$AH && t3 !== T, o3 && (this._$AH = t3);
    else {
      let e4 = t3, n3, r2;
      for (t3 = h2[0], n3 = 0; n3 < h2.length - 1; n3++) r2 = S(this, e4[s2 + n3], i4, n3), r2 === T && (r2 = this._$AH[n3]), o3 ||= !c(r2) || r2 !== this._$AH[n3], r2 === E ? t3 = E : t3 !== E && (t3 += (r2 ?? "") + h2[n3 + 1]), this._$AH[n3] = r2;
    }
    o3 && !e3 && this.j(t3);
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
  constructor(t3, i4, s2, e3, h2) {
    super(t3, i4, s2, e3, h2), this.type = 5;
  }
  _$AI(t3, i4 = this) {
    if ((t3 = S(this, t3, i4, 0) ?? E) === T) return;
    let s2 = this._$AH, e3 = t3 === E && s2 !== E || t3.capture !== s2.capture || t3.once !== s2.once || t3.passive !== s2.passive, h2 = t3 !== E && (s2 === E || e3);
    e3 && this.element.removeEventListener(this.name, this, s2), h2 && this.element.addEventListener(this.name, this, t3), this._$AH = t3;
  }
  handleEvent(t3) {
    typeof this._$AH == "function" ? this._$AH.call(this.options?.host ?? this.element, t3) : this._$AH.handleEvent(t3);
  }
}, z = class {
  constructor(t3, i4, s2) {
    this.element = t3, this.type = 6, this._$AN = void 0, this._$AM = i4, this.options = s2;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AI(t3) {
    S(this, t3);
  }
};
var j = t.litHtmlPolyfillSupport;
j?.(N, R), (t.litHtmlVersions ??= []).push("3.3.1");

// node_modules/lit-html/directive.js
var t2 = { ATTRIBUTE: 1, CHILD: 2, PROPERTY: 3, BOOLEAN_ATTRIBUTE: 4, EVENT: 5, ELEMENT: 6 }, e2 = (t3) => (...e3) => ({ _$litDirective$: t3, values: e3 }), i2 = class {
  constructor(t3) {
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AT(t3, e3, i4) {
    this._$Ct = t3, this._$AM = e3, this._$Ci = i4;
  }
  _$AS(t3, e3) {
    return this.update(t3, e3);
  }
  update(t3, e3) {
    return this.render(...e3);
  }
};

// node_modules/lit-html/directives/style-map.js
var n2 = "important", i3 = " !" + n2, o2 = e2(class extends i2 {
  constructor(t3) {
    if (super(t3), t3.type !== t2.ATTRIBUTE || t3.name !== "style" || t3.strings?.length > 2) throw Error("The `styleMap` directive must be used in the `style` attribute and must be the only part in the attribute.");
  }
  render(t3) {
    return Object.keys(t3).reduce(((e3, r2) => {
      let s2 = t3[r2];
      return s2 == null ? e3 : e3 + `${r2 = r2.includes("-") ? r2 : r2.replace(/(?:^(webkit|moz|ms|o)|)(?=[A-Z])/g, "-$&").toLowerCase()}:${s2};`;
    }), "");
  }
  update(e3, [r2]) {
    let { style: s2 } = e3.element;
    if (this.ft === void 0) return this.ft = new Set(Object.keys(r2)), this.render(r2);
    for (let t3 of this.ft) r2[t3] == null && (this.ft.delete(t3), t3.includes("-") ? s2.removeProperty(t3) : s2[t3] = null);
    for (let t3 in r2) {
      let e4 = r2[t3];
      if (e4 != null) {
        this.ft.add(t3);
        let r3 = typeof e4 == "string" && e4.endsWith(i3);
        t3.includes("-") || r3 ? s2.setProperty(t3, r3 ? e4.slice(0, -11) : e4, r3 ? n2 : "") : s2[t3] = e4;
      }
    }
    return T;
  }
});
export {
  o2 as styleMap
};
/*! Bundled license information:

lit-html/lit-html.js:
lit-html/directive.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/directives/style-map.js:
  (**
   * @license
   * Copyright 2018 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=style-map.js.map
