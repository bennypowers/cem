// elements/node_modules/lit-html/lit-html.js
var t = globalThis, i = (t4) => t4, s = t.trustedTypes, e = s ? s.createPolicy("lit-html", { createHTML: (t4) => t4 }) : void 0, h = "$lit$", o = `lit$${Math.random().toFixed(9).slice(2)}$`, n = "?" + o, r = `<${n}>`, l = document, c = () => l.createComment(""), a = (t4) => t4 === null || typeof t4 != "object" && typeof t4 != "function", u = Array.isArray, d = (t4) => u(t4) || typeof t4?.[Symbol.iterator] == "function", f = `[ 	
\f\r]`, v = /<(?:(!--|\/[^a-zA-Z])|(\/?[a-zA-Z][^>\s]*)|(\/?$))/g, _ = /-->/g, m = />/g, p = RegExp(`>|${f}(?:([^\\s"'>=/]+)(${f}*=${f}*(?:[^ 	
\f\r"'\`<>=]|("|')|))|$)`, "g"), g = /'/g, $ = /"/g, y = /^(?:script|style|textarea|title)$/i, x = (t4) => (i3, ...s3) => ({ _$litType$: t4, strings: i3, values: s3 }), b = x(1), w = x(2), T = x(3), E = /* @__PURE__ */ Symbol.for("lit-noChange"), A = /* @__PURE__ */ Symbol.for("lit-nothing"), C = /* @__PURE__ */ new WeakMap(), P = l.createTreeWalker(l, 129);
function V(t4, i3) {
  if (!u(t4) || !t4.hasOwnProperty("raw")) throw Error("invalid template strings array");
  return e !== void 0 ? e.createHTML(i3) : i3;
}
var N = (t4, i3) => {
  let s3 = t4.length - 1, e2 = [], n3, l4 = i3 === 2 ? "<svg>" : i3 === 3 ? "<math>" : "", c3 = v;
  for (let i4 = 0; i4 < s3; i4++) {
    let s4 = t4[i4], a2, u3, d4 = -1, f3 = 0;
    for (; f3 < s4.length && (c3.lastIndex = f3, u3 = c3.exec(s4), u3 !== null); ) f3 = c3.lastIndex, c3 === v ? u3[1] === "!--" ? c3 = _ : u3[1] !== void 0 ? c3 = m : u3[2] !== void 0 ? (y.test(u3[2]) && (n3 = RegExp("</" + u3[2], "g")), c3 = p) : u3[3] !== void 0 && (c3 = p) : c3 === p ? u3[0] === ">" ? (c3 = n3 ?? v, d4 = -1) : u3[1] === void 0 ? d4 = -2 : (d4 = c3.lastIndex - u3[2].length, a2 = u3[1], c3 = u3[3] === void 0 ? p : u3[3] === '"' ? $ : g) : c3 === $ || c3 === g ? c3 = p : c3 === _ || c3 === m ? c3 = v : (c3 = p, n3 = void 0);
    let x2 = c3 === p && t4[i4 + 1].startsWith("/>") ? " " : "";
    l4 += c3 === v ? s4 + r : d4 >= 0 ? (e2.push(a2), s4.slice(0, d4) + h + s4.slice(d4) + o + x2) : s4 + o + (d4 === -2 ? i4 : x2);
  }
  return [V(t4, l4 + (t4[s3] || "<?>") + (i3 === 2 ? "</svg>" : i3 === 3 ? "</math>" : "")), e2];
}, S = class _S {
  constructor({ strings: t4, _$litType$: i3 }, e2) {
    let r4;
    this.parts = [];
    let l4 = 0, a2 = 0, u3 = t4.length - 1, d4 = this.parts, [f3, v3] = N(t4, i3);
    if (this.el = _S.createElement(f3, e2), P.currentNode = this.el.content, i3 === 2 || i3 === 3) {
      let t5 = this.el.content.firstChild;
      t5.replaceWith(...t5.childNodes);
    }
    for (; (r4 = P.nextNode()) !== null && d4.length < u3; ) {
      if (r4.nodeType === 1) {
        if (r4.hasAttributes()) for (let t5 of r4.getAttributeNames()) if (t5.endsWith(h)) {
          let i4 = v3[a2++], s3 = r4.getAttribute(t5).split(o), e3 = /([.?@])?(.*)/.exec(i4);
          d4.push({ type: 1, index: l4, name: e3[2], strings: s3, ctor: e3[1] === "." ? I : e3[1] === "?" ? L : e3[1] === "@" ? z : H }), r4.removeAttribute(t5);
        } else t5.startsWith(o) && (d4.push({ type: 6, index: l4 }), r4.removeAttribute(t5));
        if (y.test(r4.tagName)) {
          let t5 = r4.textContent.split(o), i4 = t5.length - 1;
          if (i4 > 0) {
            r4.textContent = s ? s.emptyScript : "";
            for (let s3 = 0; s3 < i4; s3++) r4.append(t5[s3], c()), P.nextNode(), d4.push({ type: 2, index: ++l4 });
            r4.append(t5[i4], c());
          }
        }
      } else if (r4.nodeType === 8) if (r4.data === n) d4.push({ type: 2, index: l4 });
      else {
        let t5 = -1;
        for (; (t5 = r4.data.indexOf(o, t5 + 1)) !== -1; ) d4.push({ type: 7, index: l4 }), t5 += o.length - 1;
      }
      l4++;
    }
  }
  static createElement(t4, i3) {
    let s3 = l.createElement("template");
    return s3.innerHTML = t4, s3;
  }
};
function M(t4, i3, s3 = t4, e2) {
  if (i3 === E) return i3;
  let h3 = e2 !== void 0 ? s3._$Co?.[e2] : s3._$Cl, o2 = a(i3) ? void 0 : i3._$litDirective$;
  return h3?.constructor !== o2 && (h3?._$AO?.(!1), o2 === void 0 ? h3 = void 0 : (h3 = new o2(t4), h3._$AT(t4, s3, e2)), e2 !== void 0 ? (s3._$Co ??= [])[e2] = h3 : s3._$Cl = h3), h3 !== void 0 && (i3 = M(t4, h3._$AS(t4, i3.values), h3, e2)), i3;
}
var R = class {
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
    let { el: { content: i3 }, parts: s3 } = this._$AD, e2 = (t4?.creationScope ?? l).importNode(i3, !0);
    P.currentNode = e2;
    let h3 = P.nextNode(), o2 = 0, n3 = 0, r4 = s3[0];
    for (; r4 !== void 0; ) {
      if (o2 === r4.index) {
        let i4;
        r4.type === 2 ? i4 = new k(h3, h3.nextSibling, this, t4) : r4.type === 1 ? i4 = new r4.ctor(h3, r4.name, r4.strings, this, t4) : r4.type === 6 && (i4 = new Z(h3, this, t4)), this._$AV.push(i4), r4 = s3[++n3];
      }
      o2 !== r4?.index && (h3 = P.nextNode(), o2++);
    }
    return P.currentNode = l, e2;
  }
  p(t4) {
    let i3 = 0;
    for (let s3 of this._$AV) s3 !== void 0 && (s3.strings !== void 0 ? (s3._$AI(t4, s3, i3), i3 += s3.strings.length - 2) : s3._$AI(t4[i3])), i3++;
  }
}, k = class _k {
  get _$AU() {
    return this._$AM?._$AU ?? this._$Cv;
  }
  constructor(t4, i3, s3, e2) {
    this.type = 2, this._$AH = A, this._$AN = void 0, this._$AA = t4, this._$AB = i3, this._$AM = s3, this.options = e2, this._$Cv = e2?.isConnected ?? !0;
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
    t4 = M(this, t4, i3), a(t4) ? t4 === A || t4 == null || t4 === "" ? (this._$AH !== A && this._$AR(), this._$AH = A) : t4 !== this._$AH && t4 !== E && this._(t4) : t4._$litType$ !== void 0 ? this.$(t4) : t4.nodeType !== void 0 ? this.T(t4) : d(t4) ? this.k(t4) : this._(t4);
  }
  O(t4) {
    return this._$AA.parentNode.insertBefore(t4, this._$AB);
  }
  T(t4) {
    this._$AH !== t4 && (this._$AR(), this._$AH = this.O(t4));
  }
  _(t4) {
    this._$AH !== A && a(this._$AH) ? this._$AA.nextSibling.data = t4 : this.T(l.createTextNode(t4)), this._$AH = t4;
  }
  $(t4) {
    let { values: i3, _$litType$: s3 } = t4, e2 = typeof s3 == "number" ? this._$AC(t4) : (s3.el === void 0 && (s3.el = S.createElement(V(s3.h, s3.h[0]), this.options)), s3);
    if (this._$AH?._$AD === e2) this._$AH.p(i3);
    else {
      let t5 = new R(e2, this), s4 = t5.u(this.options);
      t5.p(i3), this.T(s4), this._$AH = t5;
    }
  }
  _$AC(t4) {
    let i3 = C.get(t4.strings);
    return i3 === void 0 && C.set(t4.strings, i3 = new S(t4)), i3;
  }
  k(t4) {
    u(this._$AH) || (this._$AH = [], this._$AR());
    let i3 = this._$AH, s3, e2 = 0;
    for (let h3 of t4) e2 === i3.length ? i3.push(s3 = new _k(this.O(c()), this.O(c()), this, this.options)) : s3 = i3[e2], s3._$AI(h3), e2++;
    e2 < i3.length && (this._$AR(s3 && s3._$AB.nextSibling, e2), i3.length = e2);
  }
  _$AR(t4 = this._$AA.nextSibling, s3) {
    for (this._$AP?.(!1, !0, s3); t4 !== this._$AB; ) {
      let s4 = i(t4).nextSibling;
      i(t4).remove(), t4 = s4;
    }
  }
  setConnected(t4) {
    this._$AM === void 0 && (this._$Cv = t4, this._$AP?.(t4));
  }
}, H = class {
  get tagName() {
    return this.element.tagName;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  constructor(t4, i3, s3, e2, h3) {
    this.type = 1, this._$AH = A, this._$AN = void 0, this.element = t4, this.name = i3, this._$AM = e2, this.options = h3, s3.length > 2 || s3[0] !== "" || s3[1] !== "" ? (this._$AH = Array(s3.length - 1).fill(new String()), this.strings = s3) : this._$AH = A;
  }
  _$AI(t4, i3 = this, s3, e2) {
    let h3 = this.strings, o2 = !1;
    if (h3 === void 0) t4 = M(this, t4, i3, 0), o2 = !a(t4) || t4 !== this._$AH && t4 !== E, o2 && (this._$AH = t4);
    else {
      let e3 = t4, n3, r4;
      for (t4 = h3[0], n3 = 0; n3 < h3.length - 1; n3++) r4 = M(this, e3[s3 + n3], i3, n3), r4 === E && (r4 = this._$AH[n3]), o2 ||= !a(r4) || r4 !== this._$AH[n3], r4 === A ? t4 = A : t4 !== A && (t4 += (r4 ?? "") + h3[n3 + 1]), this._$AH[n3] = r4;
    }
    o2 && !e2 && this.j(t4);
  }
  j(t4) {
    t4 === A ? this.element.removeAttribute(this.name) : this.element.setAttribute(this.name, t4 ?? "");
  }
}, I = class extends H {
  constructor() {
    super(...arguments), this.type = 3;
  }
  j(t4) {
    this.element[this.name] = t4 === A ? void 0 : t4;
  }
}, L = class extends H {
  constructor() {
    super(...arguments), this.type = 4;
  }
  j(t4) {
    this.element.toggleAttribute(this.name, !!t4 && t4 !== A);
  }
}, z = class extends H {
  constructor(t4, i3, s3, e2, h3) {
    super(t4, i3, s3, e2, h3), this.type = 5;
  }
  _$AI(t4, i3 = this) {
    if ((t4 = M(this, t4, i3, 0) ?? A) === E) return;
    let s3 = this._$AH, e2 = t4 === A && s3 !== A || t4.capture !== s3.capture || t4.once !== s3.once || t4.passive !== s3.passive, h3 = t4 !== A && (s3 === A || e2);
    e2 && this.element.removeEventListener(this.name, this, s3), h3 && this.element.addEventListener(this.name, this, t4), this._$AH = t4;
  }
  handleEvent(t4) {
    typeof this._$AH == "function" ? this._$AH.call(this.options?.host ?? this.element, t4) : this._$AH.handleEvent(t4);
  }
}, Z = class {
  constructor(t4, i3, s3) {
    this.element = t4, this.type = 6, this._$AN = void 0, this._$AM = i3, this.options = s3;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AI(t4) {
    M(this, t4);
  }
}, j = { M: h, P: o, A: n, C: 1, L: N, R, D: d, V: M, I: k, H, N: L, U: z, B: I, F: Z }, B = t.litHtmlPolyfillSupport;
B?.(S, k), (t.litHtmlVersions ??= []).push("3.3.2");
var D = (t4, i3, s3) => {
  let e2 = s3?.renderBefore ?? i3, h3 = e2._$litPart$;
  if (h3 === void 0) {
    let t5 = s3?.renderBefore ?? null;
    e2._$litPart$ = h3 = new k(i3.insertBefore(c(), t5), t5, void 0, s3 ?? {});
  }
  return h3._$AI(t4), h3;
};

// elements/node_modules/lit-html/private-ssr-support.js
var r2 = null, i2 = { boundAttributeSuffix: j.M, marker: j.P, markerMatch: j.A, HTML_RESULT: j.C, getTemplateHtml: j.L, overrideDirectiveResolve: (e2, t4) => class extends e2 {
  _$AS(e3, r4) {
    return t4(this, r4);
  }
}, patchDirectiveResolve: (e2, t4) => {
  if (e2.prototype._$AS.name !== t4.name) {
    r2 ??= e2.prototype._$AS.name;
    for (let i3 = e2.prototype; i3 !== Object.prototype; i3 = Object.getPrototypeOf(i3)) if (i3.hasOwnProperty(r2)) return void (i3[r2] = t4);
    throw Error("Internal error: It is possible that both dev mode and production mode Lit was mixed together during SSR. Please comment on the issue: https://github.com/lit/lit/issues/4527");
  }
}, setDirectiveClass(e2, t4) {
  e2._$litDirective$ = t4;
}, getAttributePartCommittedValue: (e2, r4, i3) => {
  let o2 = E;
  return e2.j = (e3) => o2 = e3, e2._$AI(r4, e2, i3), o2;
}, connectedDisconnectable: (e2) => ({ ...e2, _$AU: !0 }), resolveDirective: j.V, AttributePart: j.H, PropertyPart: j.B, BooleanAttributePart: j.N, EventPart: j.U, ElementPart: j.F, TemplateInstance: j.R, isIterable: j.D, ChildPart: j.I };

// elements/node_modules/lit-html/directive.js
var t2 = { ATTRIBUTE: 1, CHILD: 2, PROPERTY: 3, BOOLEAN_ATTRIBUTE: 4, EVENT: 5, ELEMENT: 6 };

// elements/node_modules/lit-html/directive-helpers.js
var { I: t3 } = j;
var n2 = (o2) => o2 === null || typeof o2 != "object" && typeof o2 != "function";
var l2 = (o2, t4) => t4 === void 0 ? o2?._$litType$ !== void 0 : o2?._$litType$ === t4, d2 = (o2) => o2?._$litType$?.h != null;
var r3 = (o2) => o2.strings === void 0;

// elements/node_modules/@lit-labs/ssr-client/lib/hydrate-lit-html.js
var { TemplateInstance: l3, isIterable: s2, resolveDirective: d3, ChildPart: c2, ElementPart: p2 } = i2, f2 = (e2, t4, r4 = {}) => {
  if (t4._$litPart$ !== void 0) throw Error("container already contains a live render");
  let n3, o2, i3, a2 = [], l4 = document.createTreeWalker(t4, NodeFilter.SHOW_COMMENT), s3;
  for (; (s3 = l4.nextNode()) !== null; ) {
    let t5 = s3.data;
    if (t5.startsWith("lit-part")) {
      if (a2.length === 0 && n3 !== void 0) throw Error(`There must be only one root part per container. Found a part marker (${s3}) when we already have a root part marker (${o2})`);
      i3 = m2(e2, s3, a2, r4), n3 === void 0 && (n3 = i3), o2 ??= s3;
    } else if (t5.startsWith("lit-node")) h2(s3, a2, r4);
    else if (t5.startsWith("/lit-part")) {
      if (a2.length === 1 && i3 !== n3) throw Error("internal error");
      i3 = u2(s3, i3, a2);
    }
  }
  if (n3 === void 0) {
    let e3 = t4 instanceof ShadowRoot ? "{container.host.localName}'s shadow root" : t4 instanceof DocumentFragment ? "DocumentFragment" : t4.localName;
    console.error(`There should be exactly one root part in a render container, but we didn't find any in ${e3}.`);
  }
  t4._$litPart$ = n3;
}, m2 = (t4, r4, a2, p3) => {
  let f3, m3;
  if (a2.length === 0) m3 = new c2(r4, null, void 0, p3), f3 = t4;
  else {
    let e2 = a2[a2.length - 1];
    if (e2.type === "template-instance") m3 = new c2(r4, null, e2.instance, p3), e2.instance._$AV.push(m3), f3 = e2.result.values[e2.instancePartIndex++], e2.templatePartIndex++;
    else if (e2.type === "iterable") {
      m3 = new c2(r4, null, e2.part, p3);
      let t5 = e2.iterator.next();
      if (t5.done) throw f3 = void 0, e2.done = !0, Error("Unhandled shorter than expected iterable");
      f3 = t5.value, e2.part._$AH.push(m3);
    } else m3 = new c2(r4, null, e2.part, p3);
  }
  if (f3 = d3(m3, f3), f3 === E) a2.push({ part: m3, type: "leaf" });
  else if (n2(f3)) a2.push({ part: m3, type: "leaf" }), m3._$AH = f3;
  else if (l2(f3)) {
    if (d2(f3)) throw Error("compiled templates are not supported");
    let e2 = "lit-part " + v2(f3);
    if (r4.data !== e2) throw Error("Hydration value mismatch: Unexpected TemplateResult rendered to part");
    {
      let e3 = c2.prototype._$AC(f3), t5 = new l3(e3, m3);
      a2.push({ type: "template-instance", instance: t5, part: m3, templatePartIndex: 0, instancePartIndex: 0, result: f3 }), m3._$AH = t5;
    }
  } else s2(f3) ? (a2.push({ part: m3, type: "iterable", value: f3, iterator: f3[Symbol.iterator](), done: !1 }), m3._$AH = []) : (a2.push({ part: m3, type: "leaf" }), m3._$AH = f3 ?? "");
  return m3;
}, u2 = (e2, t4, r4) => {
  if (t4 === void 0) throw Error("unbalanced part marker");
  t4._$AB = e2;
  let n3 = r4.pop();
  if (n3.type === "iterable" && !n3.iterator.next().done) throw Error("unexpected longer than expected iterable");
  if (r4.length > 0) return r4[r4.length - 1].part;
}, h2 = (e2, t4, n3) => {
  let o2 = /lit-node (\d+)/.exec(e2.data), i3 = parseInt(o2[1]), l4 = e2.nextElementSibling;
  if (l4 === null) throw Error("could not find node for attribute parts");
  l4.removeAttribute("defer-hydration");
  let s3 = t4[t4.length - 1];
  if (s3.type !== "template-instance") throw Error("Hydration value mismatch: Primitive found where TemplateResult expected. This usually occurs due to conditional rendering that resulted in a different value or template being rendered between the server and client.");
  {
    let e3 = s3.instance;
    for (; ; ) {
      let t5 = e3._$AD.parts[s3.templatePartIndex];
      if (t5 === void 0 || t5.type !== t2.ATTRIBUTE && t5.type !== t2.ELEMENT || t5.index !== i3) break;
      if (t5.type === t2.ATTRIBUTE) {
        let o3 = new t5.ctor(l4, t5.name, t5.strings, s3.instance, n3), i4 = r3(o3) ? s3.result.values[s3.instancePartIndex] : s3.result.values, d4 = !(o3.type === t2.EVENT || o3.type === t2.PROPERTY);
        o3._$AI(i4, o3, s3.instancePartIndex, d4), s3.instancePartIndex += t5.strings.length - 1, e3._$AV.push(o3);
      } else {
        let t6 = new p2(l4, s3.instance, n3);
        d3(t6, s3.result.values[s3.instancePartIndex++]), e3._$AV.push(t6);
      }
      s3.templatePartIndex++;
    }
  }
}, w2 = /* @__PURE__ */ new WeakMap(), v2 = (e2) => {
  let t4 = w2.get(e2.strings);
  if (t4 !== void 0) return t4;
  let r4 = new Uint32Array(2).fill(5381);
  for (let t5 of e2.strings) for (let e3 = 0; e3 < t5.length; e3++) r4[e3 % 2] = 33 * r4[e3 % 2] ^ t5.charCodeAt(e3);
  let n3 = String.fromCharCode(...new Uint8Array(r4.buffer));
  return t4 = btoa(n3), w2.set(e2.strings, t4), t4;
};

// elements/node_modules/@lit-labs/ssr-client/lit-element-hydrate-support.js
globalThis.litElementHydrateSupport = ({ LitElement: s3 }) => {
  let h3 = Object.getOwnPropertyDescriptor(Object.getPrototypeOf(s3), "observedAttributes").get;
  Object.defineProperty(s3, "observedAttributes", { get() {
    return [...h3.call(this), "defer-hydration"];
  } });
  let e2 = s3.prototype.attributeChangedCallback;
  s3.prototype.attributeChangedCallback = function(t4, i3, s4) {
    t4 === "defer-hydration" && s4 === null && n3.call(this), e2.call(this, t4, i3, s4);
  };
  let n3 = s3.prototype.connectedCallback;
  s3.prototype.connectedCallback = function() {
    this.hasAttribute("defer-hydration") || n3.call(this);
  };
  let o2 = s3.prototype.createRenderRoot;
  s3.prototype.createRenderRoot = function() {
    return this.shadowRoot ? (this._$AG = !0, this.shadowRoot) : o2.call(this);
  };
  let r4 = Object.getPrototypeOf(s3.prototype).update;
  s3.prototype.update = function(s4) {
    let h4 = this.render();
    if (r4.call(this, s4), this._$AG) {
      this._$AG = !1;
      for (let t4 of this.getAttributeNames()) if (t4.startsWith("hydrate-internals-")) {
        let i3 = t4.slice(18);
        this.removeAttribute(i3), this.removeAttribute(t4);
      }
      f2(h4, this.renderRoot, this.renderOptions);
    } else D(h4, this.renderRoot, this.renderOptions);
  };
};
/*! Bundled license information:

lit-html/lit-html.js:
lit-html/directive.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/private-ssr-support.js:
@lit-labs/ssr-client/lib/hydrate-lit-html.js:
  (**
   * @license
   * Copyright 2019 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/directive-helpers.js:
  (**
   * @license
   * Copyright 2020 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
//# sourceMappingURL=lit-element-hydrate-support.js.map
