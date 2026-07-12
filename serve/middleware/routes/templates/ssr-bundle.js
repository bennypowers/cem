var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __knownSymbol = (name, symbol) => (symbol = Symbol[name]) ? symbol : /* @__PURE__ */ Symbol.for("Symbol." + name);
var __typeError = (msg) => {
  throw TypeError(msg);
};
var __defNormalProp = (obj, key, value) => key in obj ? __defProp(obj, key, { enumerable: true, configurable: true, writable: true, value }) : obj[key] = value;
var __name = (target, value) => __defProp(target, "name", { value, configurable: true });
var __esm = (fn, res) => function __init() {
  return fn && (res = (0, fn[__getOwnPropNames(fn)[0]])(fn = 0)), res;
};
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __decoratorStart = (base) => [, , , __create(base?.[__knownSymbol("metadata")] ?? null)];
var __decoratorStrings = ["class", "method", "getter", "setter", "accessor", "field", "value", "get", "set"];
var __expectFn = (fn) => fn !== void 0 && typeof fn !== "function" ? __typeError("Function expected") : fn;
var __decoratorContext = (kind, name, done, metadata, fns) => ({ kind: __decoratorStrings[kind], name, metadata, addInitializer: (fn) => done._ ? __typeError("Already initialized") : fns.push(__expectFn(fn || null)) });
var __decoratorMetadata = (array, target) => __defNormalProp(target, __knownSymbol("metadata"), array[3]);
var __runInitializers = (array, flags, self, value) => {
  for (var i6 = 0, fns = array[flags >> 1], n6 = fns && fns.length; i6 < n6; i6++) flags & 1 ? fns[i6].call(self) : value = fns[i6].call(self, value);
  return value;
};
var __decorateElement = (array, flags, name, decorators, target, extra) => {
  var fn, it, done, ctx, access, k2 = flags & 7, s58 = !!(flags & 8), p4 = !!(flags & 16);
  var j2 = k2 > 3 ? array.length + 1 : k2 ? s58 ? 1 : 2 : 0, key = __decoratorStrings[k2 + 5];
  var initializers = k2 > 3 && (array[j2 - 1] = []), extraInitializers = array[j2] || (array[j2] = []);
  var desc = k2 && (!p4 && !s58 && (target = target.prototype), k2 < 5 && (k2 > 3 || !p4) && __getOwnPropDesc(k2 < 4 ? target : { get [name]() {
    return __privateGet(this, extra);
  }, set [name](x2) {
    return __privateSet(this, extra, x2);
  } }, name));
  k2 ? p4 && k2 < 4 && __name(extra, (k2 > 2 ? "set " : k2 > 1 ? "get " : "") + name) : __name(target, name);
  for (var i6 = decorators.length - 1; i6 >= 0; i6--) {
    ctx = __decoratorContext(k2, name, done = {}, array[3], extraInitializers);
    if (k2) {
      ctx.static = s58, ctx.private = p4, access = ctx.access = { has: p4 ? (x2) => __privateIn(target, x2) : (x2) => name in x2 };
      if (k2 ^ 3) access.get = p4 ? (x2) => (k2 ^ 1 ? __privateGet : __privateMethod)(x2, target, k2 ^ 4 ? extra : desc.get) : (x2) => x2[name];
      if (k2 > 2) access.set = p4 ? (x2, y3) => __privateSet(x2, target, y3, k2 ^ 4 ? extra : desc.set) : (x2, y3) => x2[name] = y3;
    }
    it = (0, decorators[i6])(k2 ? k2 < 4 ? p4 ? extra : desc[key] : k2 > 4 ? void 0 : { get: desc.get, set: desc.set } : target, ctx), done._ = 1;
    if (k2 ^ 4 || it === void 0) __expectFn(it) && (k2 > 4 ? initializers.unshift(it) : k2 ? p4 ? extra = it : desc[key] = it : target = it);
    else if (typeof it !== "object" || it === null) __typeError("Object expected");
    else __expectFn(fn = it.get) && (desc.get = fn), __expectFn(fn = it.set) && (desc.set = fn), __expectFn(fn = it.init) && initializers.unshift(fn);
  }
  return k2 || __decoratorMetadata(array, target), desc && __defProp(target, name, desc), p4 ? k2 ^ 4 ? extra : desc : target;
};
var __publicField = (obj, key, value) => __defNormalProp(obj, typeof key !== "symbol" ? key + "" : key, value);
var __accessCheck = (obj, member, msg) => member.has(obj) || __typeError("Cannot " + msg);
var __privateIn = (member, obj) => Object(obj) !== obj ? __typeError('Cannot use the "in" operator on this value') : member.has(obj);
var __privateGet = (obj, member, getter) => (__accessCheck(obj, member, "read from private field"), getter ? getter.call(obj) : member.get(obj));
var __privateAdd = (obj, member, value) => member.has(obj) ? __typeError("Cannot add the same private member more than once") : member instanceof WeakSet ? member.add(obj) : member.set(obj, value);
var __privateSet = (obj, member, value, setter) => (__accessCheck(obj, member, "write to private field"), setter ? setter.call(obj, value) : member.set(obj, value), value);
var __privateMethod = (obj, member, method) => (__accessCheck(obj, member, "access private method"), method);

// cem-stub:/__cem/websocket-client.js
var websocket_client_exports = {};
__export(websocket_client_exports, {
  CEMReloadClient: () => CEMReloadClient,
  StatePersistence: () => StatePersistence,
  default: () => websocket_client_default
});
var websocket_client_default, CEMReloadClient, StatePersistence;
var init_websocket_client = __esm({
  "cem-stub:/__cem/websocket-client.js"() {
    websocket_client_default = {};
    CEMReloadClient = class {
    };
    StatePersistence = class {
    };
  }
});

// cem-stub:/__cem/state-persistence.js
var state_persistence_exports = {};
__export(state_persistence_exports, {
  CEMReloadClient: () => CEMReloadClient2,
  StatePersistence: () => StatePersistence2,
  default: () => state_persistence_default
});
var state_persistence_default, CEMReloadClient2, StatePersistence2;
var init_state_persistence = __esm({
  "cem-stub:/__cem/state-persistence.js"() {
    state_persistence_default = {};
    CEMReloadClient2 = class {
    };
    StatePersistence2 = class {
    };
  }
});

// cem-stub:/__cem/health-badges.js
var health_badges_exports = {};
__export(health_badges_exports, {
  CEMReloadClient: () => CEMReloadClient3,
  StatePersistence: () => StatePersistence3,
  default: () => health_badges_default
});
var health_badges_default, CEMReloadClient3, StatePersistence3;
var init_health_badges = __esm({
  "cem-stub:/__cem/health-badges.js"() {
    health_badges_default = {};
    CEMReloadClient3 = class {
    };
    StatePersistence3 = class {
    };
  }
});

// lit-ssr-wasm-shim:@lit-labs/ssr-dom-shim
var customElements2 = globalThis.customElements;
var HTMLElement2 = globalThis.HTMLElement;
var Element2 = globalThis.Element;
var CSSStyleSheet2 = globalThis.CSSStyleSheet;
var CustomElementRegistry = globalThis.CustomElementRegistry;
var Event2 = globalThis.Event;
var CustomEvent2 = globalThis.CustomEvent;
var EventTarget = globalThis.EventTarget;
var ariaMixinAttributes = globalThis.ariaMixinAttributes ?? {};
var HYDRATE_INTERNALS_ATTR_PREFIX = globalThis.HYDRATE_INTERNALS_ATTR_PREFIX ?? "internals-";
var ElementInternals = globalThis.ElementInternals ?? class ElementInternals2 {
};

// elements/node_modules/@lit/reactive-element/node/css-tag.js
var t = globalThis;
var e = t.ShadowRoot && (void 0 === t.ShadyCSS || t.ShadyCSS.nativeShadow) && "adoptedStyleSheets" in Document.prototype && "replace" in CSSStyleSheet.prototype;
var s = /* @__PURE__ */ Symbol();
var o = /* @__PURE__ */ new WeakMap();
var n = class {
  constructor(t6, e6, o9) {
    if (this._$cssResult$ = true, o9 !== s) throw Error("CSSResult is not constructable. Use `unsafeCSS` or `css` instead.");
    this.cssText = t6, this.t = e6;
  }
  get styleSheet() {
    let t6 = this.o;
    const s58 = this.t;
    if (e && void 0 === t6) {
      const e6 = void 0 !== s58 && 1 === s58.length;
      e6 && (t6 = o.get(s58)), void 0 === t6 && ((this.o = t6 = new CSSStyleSheet()).replaceSync(this.cssText), e6 && o.set(s58, t6));
    }
    return t6;
  }
  toString() {
    return this.cssText;
  }
};
var r = (t6) => new n("string" == typeof t6 ? t6 : t6 + "", void 0, s);
var S = (s58, o9) => {
  if (e) s58.adoptedStyleSheets = o9.map((t6) => t6 instanceof CSSStyleSheet ? t6 : t6.styleSheet);
  else for (const e6 of o9) {
    const o10 = document.createElement("style"), n6 = t.litNonce;
    void 0 !== n6 && o10.setAttribute("nonce", n6), o10.textContent = e6.cssText, s58.appendChild(o10);
  }
};
var c = e || void 0 === t.CSSStyleSheet ? (t6) => t6 : (t6) => t6 instanceof CSSStyleSheet ? ((t7) => {
  let e6 = "";
  for (const s58 of t7.cssRules) e6 += s58.cssText;
  return r(e6);
})(t6) : t6;

// elements/node_modules/@lit/reactive-element/node/reactive-element.js
var { is: h, defineProperty: r2, getOwnPropertyDescriptor: o2, getOwnPropertyNames: n2, getOwnPropertySymbols: a, getPrototypeOf: c2 } = Object;
var l = globalThis;
l.customElements ??= customElements2;
var p = l.trustedTypes;
var d = p ? p.emptyScript : "";
var u = l.reactiveElementPolyfillSupport;
var f = (t6, s58) => t6;
var b = { toAttribute(t6, s58) {
  switch (s58) {
    case Boolean:
      t6 = t6 ? d : null;
      break;
    case Object:
    case Array:
      t6 = null == t6 ? t6 : JSON.stringify(t6);
  }
  return t6;
}, fromAttribute(t6, s58) {
  let i6 = t6;
  switch (s58) {
    case Boolean:
      i6 = null !== t6;
      break;
    case Number:
      i6 = null === t6 ? null : Number(t6);
      break;
    case Object:
    case Array:
      try {
        i6 = JSON.parse(t6);
      } catch (t7) {
        i6 = null;
      }
  }
  return i6;
} };
var m = (t6, s58) => !h(t6, s58);
var y = { attribute: true, type: String, converter: b, reflect: false, useDefault: false, hasChanged: m };
Symbol.metadata ??= /* @__PURE__ */ Symbol("metadata"), l.litPropertyMetadata ??= /* @__PURE__ */ new WeakMap();
var g = class extends (globalThis.HTMLElement ?? HTMLElement2) {
  static addInitializer(t6) {
    this._$Ei(), (this.l ??= []).push(t6);
  }
  static get observedAttributes() {
    return this.finalize(), this._$Eh && [...this._$Eh.keys()];
  }
  static createProperty(t6, s58 = y) {
    if (s58.state && (s58.attribute = false), this._$Ei(), this.prototype.hasOwnProperty(t6) && ((s58 = Object.create(s58)).wrapped = true), this.elementProperties.set(t6, s58), !s58.noAccessor) {
      const i6 = /* @__PURE__ */ Symbol(), e6 = this.getPropertyDescriptor(t6, i6, s58);
      void 0 !== e6 && r2(this.prototype, t6, e6);
    }
  }
  static getPropertyDescriptor(t6, s58, i6) {
    const { get: e6, set: h3 } = o2(this.prototype, t6) ?? { get() {
      return this[s58];
    }, set(t7) {
      this[s58] = t7;
    } };
    return { get: e6, set(s59) {
      const r7 = e6?.call(this);
      h3?.call(this, s59), this.requestUpdate(t6, r7, i6);
    }, configurable: true, enumerable: true };
  }
  static getPropertyOptions(t6) {
    return this.elementProperties.get(t6) ?? y;
  }
  static _$Ei() {
    if (this.hasOwnProperty(f("elementProperties"))) return;
    const t6 = c2(this);
    t6.finalize(), void 0 !== t6.l && (this.l = [...t6.l]), this.elementProperties = new Map(t6.elementProperties);
  }
  static finalize() {
    if (this.hasOwnProperty(f("finalized"))) return;
    if (this.finalized = true, this._$Ei(), this.hasOwnProperty(f("properties"))) {
      const t7 = this.properties, s58 = [...n2(t7), ...a(t7)];
      for (const i6 of s58) this.createProperty(i6, t7[i6]);
    }
    const t6 = this[Symbol.metadata];
    if (null !== t6) {
      const s58 = litPropertyMetadata.get(t6);
      if (void 0 !== s58) for (const [t7, i6] of s58) this.elementProperties.set(t7, i6);
    }
    this._$Eh = /* @__PURE__ */ new Map();
    for (const [t7, s58] of this.elementProperties) {
      const i6 = this._$Eu(t7, s58);
      void 0 !== i6 && this._$Eh.set(i6, t7);
    }
    this.elementStyles = this.finalizeStyles(this.styles);
  }
  static finalizeStyles(t6) {
    const s58 = [];
    if (Array.isArray(t6)) {
      const e6 = new Set(t6.flat(1 / 0).reverse());
      for (const t7 of e6) s58.unshift(c(t7));
    } else void 0 !== t6 && s58.push(c(t6));
    return s58;
  }
  static _$Eu(t6, s58) {
    const i6 = s58.attribute;
    return false === i6 ? void 0 : "string" == typeof i6 ? i6 : "string" == typeof t6 ? t6.toLowerCase() : void 0;
  }
  constructor() {
    super(), this._$Ep = void 0, this.isUpdatePending = false, this.hasUpdated = false, this._$Em = null, this._$Ev();
  }
  _$Ev() {
    this._$ES = new Promise((t6) => this.enableUpdating = t6), this._$AL = /* @__PURE__ */ new Map(), this._$E_(), this.requestUpdate(), this.constructor.l?.forEach((t6) => t6(this));
  }
  addController(t6) {
    (this._$EO ??= /* @__PURE__ */ new Set()).add(t6), void 0 !== this.renderRoot && this.isConnected && t6.hostConnected?.();
  }
  removeController(t6) {
    this._$EO?.delete(t6);
  }
  _$E_() {
    const t6 = /* @__PURE__ */ new Map(), s58 = this.constructor.elementProperties;
    for (const i6 of s58.keys()) this.hasOwnProperty(i6) && (t6.set(i6, this[i6]), delete this[i6]);
    t6.size > 0 && (this._$Ep = t6);
  }
  createRenderRoot() {
    const t6 = this.shadowRoot ?? this.attachShadow(this.constructor.shadowRootOptions);
    return S(t6, this.constructor.elementStyles), t6;
  }
  connectedCallback() {
    this.renderRoot ??= this.createRenderRoot(), this.enableUpdating(true), this._$EO?.forEach((t6) => t6.hostConnected?.());
  }
  enableUpdating(t6) {
  }
  disconnectedCallback() {
    this._$EO?.forEach((t6) => t6.hostDisconnected?.());
  }
  attributeChangedCallback(t6, s58, i6) {
    this._$AK(t6, i6);
  }
  _$ET(t6, s58) {
    const i6 = this.constructor.elementProperties.get(t6), e6 = this.constructor._$Eu(t6, i6);
    if (void 0 !== e6 && true === i6.reflect) {
      const h3 = (void 0 !== i6.converter?.toAttribute ? i6.converter : b).toAttribute(s58, i6.type);
      this._$Em = t6, null == h3 ? this.removeAttribute(e6) : this.setAttribute(e6, h3), this._$Em = null;
    }
  }
  _$AK(t6, s58) {
    const i6 = this.constructor, e6 = i6._$Eh.get(t6);
    if (void 0 !== e6 && this._$Em !== e6) {
      const t7 = i6.getPropertyOptions(e6), h3 = "function" == typeof t7.converter ? { fromAttribute: t7.converter } : void 0 !== t7.converter?.fromAttribute ? t7.converter : b;
      this._$Em = e6;
      const r7 = h3.fromAttribute(s58, t7.type);
      this[e6] = r7 ?? this._$Ej?.get(e6) ?? r7, this._$Em = null;
    }
  }
  requestUpdate(t6, s58, i6, e6 = false, h3) {
    if (void 0 !== t6) {
      const r7 = this.constructor;
      if (false === e6 && (h3 = this[t6]), i6 ??= r7.getPropertyOptions(t6), !((i6.hasChanged ?? m)(h3, s58) || i6.useDefault && i6.reflect && h3 === this._$Ej?.get(t6) && !this.hasAttribute(r7._$Eu(t6, i6)))) return;
      this.C(t6, s58, i6);
    }
    false === this.isUpdatePending && (this._$ES = this._$EP());
  }
  C(t6, s58, { useDefault: i6, reflect: e6, wrapped: h3 }, r7) {
    i6 && !(this._$Ej ??= /* @__PURE__ */ new Map()).has(t6) && (this._$Ej.set(t6, r7 ?? s58 ?? this[t6]), true !== h3 || void 0 !== r7) || (this._$AL.has(t6) || (this.hasUpdated || i6 || (s58 = void 0), this._$AL.set(t6, s58)), true === e6 && this._$Em !== t6 && (this._$Eq ??= /* @__PURE__ */ new Set()).add(t6));
  }
  async _$EP() {
    this.isUpdatePending = true;
    try {
      await this._$ES;
    } catch (t7) {
      Promise.reject(t7);
    }
    const t6 = this.scheduleUpdate();
    return null != t6 && await t6, !this.isUpdatePending;
  }
  scheduleUpdate() {
    return this.performUpdate();
  }
  performUpdate() {
    if (!this.isUpdatePending) return;
    if (!this.hasUpdated) {
      if (this.renderRoot ??= this.createRenderRoot(), this._$Ep) {
        for (const [t8, s59] of this._$Ep) this[t8] = s59;
        this._$Ep = void 0;
      }
      const t7 = this.constructor.elementProperties;
      if (t7.size > 0) for (const [s59, i6] of t7) {
        const { wrapped: t8 } = i6, e6 = this[s59];
        true !== t8 || this._$AL.has(s59) || void 0 === e6 || this.C(s59, void 0, i6, e6);
      }
    }
    let t6 = false;
    const s58 = this._$AL;
    try {
      t6 = this.shouldUpdate(s58), t6 ? (this.willUpdate(s58), this._$EO?.forEach((t7) => t7.hostUpdate?.()), this.update(s58)) : this._$EM();
    } catch (s59) {
      throw t6 = false, this._$EM(), s59;
    }
    t6 && this._$AE(s58);
  }
  willUpdate(t6) {
  }
  _$AE(t6) {
    this._$EO?.forEach((t7) => t7.hostUpdated?.()), this.hasUpdated || (this.hasUpdated = true, this.firstUpdated(t6)), this.updated(t6);
  }
  _$EM() {
    this._$AL = /* @__PURE__ */ new Map(), this.isUpdatePending = false;
  }
  get updateComplete() {
    return this.getUpdateComplete();
  }
  getUpdateComplete() {
    return this._$ES;
  }
  shouldUpdate(t6) {
    return true;
  }
  update(t6) {
    this._$Eq &&= this._$Eq.forEach((t7) => this._$ET(t7, this[t7])), this._$EM();
  }
  updated(t6) {
  }
  firstUpdated(t6) {
  }
};
g.elementStyles = [], g.shadowRootOptions = { mode: "open" }, g[f("elementProperties")] = /* @__PURE__ */ new Map(), g[f("finalized")] = /* @__PURE__ */ new Map(), u?.({ ReactiveElement: g }), (l.reactiveElementVersions ??= []).push("2.1.2");

// elements/node_modules/lit-html/node/lit-html.js
var t2 = globalThis;
var i2 = (t6) => t6;
var s2 = t2.trustedTypes;
var e2 = s2 ? s2.createPolicy("lit-html", { createHTML: (t6) => t6 }) : void 0;
var h2 = "$lit$";
var o3 = `lit$${Math.random().toFixed(9).slice(2)}$`;
var n3 = "?" + o3;
var r3 = `<${n3}>`;
var l2 = void 0 === t2.document ? { createTreeWalker: () => ({}) } : document;
var c3 = () => l2.createComment("");
var a2 = (t6) => null === t6 || "object" != typeof t6 && "function" != typeof t6;
var u2 = Array.isArray;
var d2 = (t6) => u2(t6) || "function" == typeof t6?.[Symbol.iterator];
var f2 = "[ 	\n\f\r]";
var v = /<(?:(!--|\/[^a-zA-Z])|(\/?[a-zA-Z][^>\s]*)|(\/?$))/g;
var _ = /-->/g;
var m2 = />/g;
var p2 = RegExp(`>|${f2}(?:([^\\s"'>=/]+)(${f2}*=${f2}*(?:[^ 	
\f\r"'\`<>=]|("|')|))|$)`, "g");
var g2 = /'/g;
var $ = /"/g;
var y2 = /^(?:script|style|textarea|title)$/i;
var x = (t6) => (i6, ...s58) => ({ _$litType$: t6, strings: i6, values: s58 });
var T = x(1);
var b2 = x(2);
var w = x(3);
var E = /* @__PURE__ */ Symbol.for("lit-noChange");
var A = /* @__PURE__ */ Symbol.for("lit-nothing");
var C = /* @__PURE__ */ new WeakMap();
var P = l2.createTreeWalker(l2, 129);
function V(t6, i6) {
  if (!u2(t6) || !t6.hasOwnProperty("raw")) throw Error("invalid template strings array");
  return void 0 !== e2 ? e2.createHTML(i6) : i6;
}
var N = (t6, i6) => {
  const s58 = t6.length - 1, e6 = [];
  let n6, l4 = 2 === i6 ? "<svg>" : 3 === i6 ? "<math>" : "", c4 = v;
  for (let i7 = 0; i7 < s58; i7++) {
    const s59 = t6[i7];
    let a3, u3, d3 = -1, f3 = 0;
    for (; f3 < s59.length && (c4.lastIndex = f3, u3 = c4.exec(s59), null !== u3); ) f3 = c4.lastIndex, c4 === v ? "!--" === u3[1] ? c4 = _ : void 0 !== u3[1] ? c4 = m2 : void 0 !== u3[2] ? (y2.test(u3[2]) && (n6 = RegExp("</" + u3[2], "g")), c4 = p2) : void 0 !== u3[3] && (c4 = p2) : c4 === p2 ? ">" === u3[0] ? (c4 = n6 ?? v, d3 = -1) : void 0 === u3[1] ? d3 = -2 : (d3 = c4.lastIndex - u3[2].length, a3 = u3[1], c4 = void 0 === u3[3] ? p2 : '"' === u3[3] ? $ : g2) : c4 === $ || c4 === g2 ? c4 = p2 : c4 === _ || c4 === m2 ? c4 = v : (c4 = p2, n6 = void 0);
    const x2 = c4 === p2 && t6[i7 + 1].startsWith("/>") ? " " : "";
    l4 += c4 === v ? s59 + r3 : d3 >= 0 ? (e6.push(a3), s59.slice(0, d3) + h2 + s59.slice(d3) + o3 + x2) : s59 + o3 + (-2 === d3 ? i7 : x2);
  }
  return [V(t6, l4 + (t6[s58] || "<?>") + (2 === i6 ? "</svg>" : 3 === i6 ? "</math>" : "")), e6];
};
var S2 = class _S {
  constructor({ strings: t6, _$litType$: i6 }, e6) {
    let r7;
    this.parts = [];
    let l4 = 0, a3 = 0;
    const u3 = t6.length - 1, d3 = this.parts, [f3, v2] = N(t6, i6);
    if (this.el = _S.createElement(f3, e6), P.currentNode = this.el.content, 2 === i6 || 3 === i6) {
      const t7 = this.el.content.firstChild;
      t7.replaceWith(...t7.childNodes);
    }
    for (; null !== (r7 = P.nextNode()) && d3.length < u3; ) {
      if (1 === r7.nodeType) {
        if (r7.hasAttributes()) for (const t7 of r7.getAttributeNames()) if (t7.endsWith(h2)) {
          const i7 = v2[a3++], s58 = r7.getAttribute(t7).split(o3), e7 = /([.?@])?(.*)/.exec(i7);
          d3.push({ type: 1, index: l4, name: e7[2], strings: s58, ctor: "." === e7[1] ? I : "?" === e7[1] ? L : "@" === e7[1] ? z : H }), r7.removeAttribute(t7);
        } else t7.startsWith(o3) && (d3.push({ type: 6, index: l4 }), r7.removeAttribute(t7));
        if (y2.test(r7.tagName)) {
          const t7 = r7.textContent.split(o3), i7 = t7.length - 1;
          if (i7 > 0) {
            r7.textContent = s2 ? s2.emptyScript : "";
            for (let s58 = 0; s58 < i7; s58++) r7.append(t7[s58], c3()), P.nextNode(), d3.push({ type: 2, index: ++l4 });
            r7.append(t7[i7], c3());
          }
        }
      } else if (8 === r7.nodeType) if (r7.data === n3) d3.push({ type: 2, index: l4 });
      else {
        let t7 = -1;
        for (; -1 !== (t7 = r7.data.indexOf(o3, t7 + 1)); ) d3.push({ type: 7, index: l4 }), t7 += o3.length - 1;
      }
      l4++;
    }
  }
  static createElement(t6, i6) {
    const s58 = l2.createElement("template");
    return s58.innerHTML = t6, s58;
  }
};
function M(t6, i6, s58 = t6, e6) {
  if (i6 === E) return i6;
  let h3 = void 0 !== e6 ? s58._$Co?.[e6] : s58._$Cl;
  const o9 = a2(i6) ? void 0 : i6._$litDirective$;
  return h3?.constructor !== o9 && (h3?._$AO?.(false), void 0 === o9 ? h3 = void 0 : (h3 = new o9(t6), h3._$AT(t6, s58, e6)), void 0 !== e6 ? (s58._$Co ??= [])[e6] = h3 : s58._$Cl = h3), void 0 !== h3 && (i6 = M(t6, h3._$AS(t6, i6.values), h3, e6)), i6;
}
var k = class {
  constructor(t6, i6) {
    this._$AV = [], this._$AN = void 0, this._$AD = t6, this._$AM = i6;
  }
  get parentNode() {
    return this._$AM.parentNode;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  u(t6) {
    const { el: { content: i6 }, parts: s58 } = this._$AD, e6 = (t6?.creationScope ?? l2).importNode(i6, true);
    P.currentNode = e6;
    let h3 = P.nextNode(), o9 = 0, n6 = 0, r7 = s58[0];
    for (; void 0 !== r7; ) {
      if (o9 === r7.index) {
        let i7;
        2 === r7.type ? i7 = new R(h3, h3.nextSibling, this, t6) : 1 === r7.type ? i7 = new r7.ctor(h3, r7.name, r7.strings, this, t6) : 6 === r7.type && (i7 = new W(h3, this, t6)), this._$AV.push(i7), r7 = s58[++n6];
      }
      o9 !== r7?.index && (h3 = P.nextNode(), o9++);
    }
    return P.currentNode = l2, e6;
  }
  p(t6) {
    let i6 = 0;
    for (const s58 of this._$AV) void 0 !== s58 && (void 0 !== s58.strings ? (s58._$AI(t6, s58, i6), i6 += s58.strings.length - 2) : s58._$AI(t6[i6])), i6++;
  }
};
var R = class _R {
  get _$AU() {
    return this._$AM?._$AU ?? this._$Cv;
  }
  constructor(t6, i6, s58, e6) {
    this.type = 2, this._$AH = A, this._$AN = void 0, this._$AA = t6, this._$AB = i6, this._$AM = s58, this.options = e6, this._$Cv = e6?.isConnected ?? true;
  }
  get parentNode() {
    let t6 = this._$AA.parentNode;
    const i6 = this._$AM;
    return void 0 !== i6 && 11 === t6?.nodeType && (t6 = i6.parentNode), t6;
  }
  get startNode() {
    return this._$AA;
  }
  get endNode() {
    return this._$AB;
  }
  _$AI(t6, i6 = this) {
    t6 = M(this, t6, i6), a2(t6) ? t6 === A || null == t6 || "" === t6 ? (this._$AH !== A && this._$AR(), this._$AH = A) : t6 !== this._$AH && t6 !== E && this._(t6) : void 0 !== t6._$litType$ ? this.$(t6) : void 0 !== t6.nodeType ? this.T(t6) : d2(t6) ? this.k(t6) : this._(t6);
  }
  O(t6) {
    return this._$AA.parentNode.insertBefore(t6, this._$AB);
  }
  T(t6) {
    this._$AH !== t6 && (this._$AR(), this._$AH = this.O(t6));
  }
  _(t6) {
    this._$AH !== A && a2(this._$AH) ? this._$AA.nextSibling.data = t6 : this.T(l2.createTextNode(t6)), this._$AH = t6;
  }
  $(t6) {
    const { values: i6, _$litType$: s58 } = t6, e6 = "number" == typeof s58 ? this._$AC(t6) : (void 0 === s58.el && (s58.el = S2.createElement(V(s58.h, s58.h[0]), this.options)), s58);
    if (this._$AH?._$AD === e6) this._$AH.p(i6);
    else {
      const t7 = new k(e6, this), s59 = t7.u(this.options);
      t7.p(i6), this.T(s59), this._$AH = t7;
    }
  }
  _$AC(t6) {
    let i6 = C.get(t6.strings);
    return void 0 === i6 && C.set(t6.strings, i6 = new S2(t6)), i6;
  }
  k(t6) {
    u2(this._$AH) || (this._$AH = [], this._$AR());
    const i6 = this._$AH;
    let s58, e6 = 0;
    for (const h3 of t6) e6 === i6.length ? i6.push(s58 = new _R(this.O(c3()), this.O(c3()), this, this.options)) : s58 = i6[e6], s58._$AI(h3), e6++;
    e6 < i6.length && (this._$AR(s58 && s58._$AB.nextSibling, e6), i6.length = e6);
  }
  _$AR(t6 = this._$AA.nextSibling, s58) {
    for (this._$AP?.(false, true, s58); t6 !== this._$AB; ) {
      const s59 = i2(t6).nextSibling;
      i2(t6).remove(), t6 = s59;
    }
  }
  setConnected(t6) {
    void 0 === this._$AM && (this._$Cv = t6, this._$AP?.(t6));
  }
};
var H = class {
  get tagName() {
    return this.element.tagName;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  constructor(t6, i6, s58, e6, h3) {
    this.type = 1, this._$AH = A, this._$AN = void 0, this.element = t6, this.name = i6, this._$AM = e6, this.options = h3, s58.length > 2 || "" !== s58[0] || "" !== s58[1] ? (this._$AH = Array(s58.length - 1).fill(new String()), this.strings = s58) : this._$AH = A;
  }
  _$AI(t6, i6 = this, s58, e6) {
    const h3 = this.strings;
    let o9 = false;
    if (void 0 === h3) t6 = M(this, t6, i6, 0), o9 = !a2(t6) || t6 !== this._$AH && t6 !== E, o9 && (this._$AH = t6);
    else {
      const e7 = t6;
      let n6, r7;
      for (t6 = h3[0], n6 = 0; n6 < h3.length - 1; n6++) r7 = M(this, e7[s58 + n6], i6, n6), r7 === E && (r7 = this._$AH[n6]), o9 ||= !a2(r7) || r7 !== this._$AH[n6], r7 === A ? t6 = A : t6 !== A && (t6 += (r7 ?? "") + h3[n6 + 1]), this._$AH[n6] = r7;
    }
    o9 && !e6 && this.j(t6);
  }
  j(t6) {
    t6 === A ? this.element.removeAttribute(this.name) : this.element.setAttribute(this.name, t6 ?? "");
  }
};
var I = class extends H {
  constructor() {
    super(...arguments), this.type = 3;
  }
  j(t6) {
    this.element[this.name] = t6 === A ? void 0 : t6;
  }
};
var L = class extends H {
  constructor() {
    super(...arguments), this.type = 4;
  }
  j(t6) {
    this.element.toggleAttribute(this.name, !!t6 && t6 !== A);
  }
};
var z = class extends H {
  constructor(t6, i6, s58, e6, h3) {
    super(t6, i6, s58, e6, h3), this.type = 5;
  }
  _$AI(t6, i6 = this) {
    if ((t6 = M(this, t6, i6, 0) ?? A) === E) return;
    const s58 = this._$AH, e6 = t6 === A && s58 !== A || t6.capture !== s58.capture || t6.once !== s58.once || t6.passive !== s58.passive, h3 = t6 !== A && (s58 === A || e6);
    e6 && this.element.removeEventListener(this.name, this, s58), h3 && this.element.addEventListener(this.name, this, t6), this._$AH = t6;
  }
  handleEvent(t6) {
    "function" == typeof this._$AH ? this._$AH.call(this.options?.host ?? this.element, t6) : this._$AH.handleEvent(t6);
  }
};
var W = class {
  constructor(t6, i6, s58) {
    this.element = t6, this.type = 6, this._$AN = void 0, this._$AM = i6, this.options = s58;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AI(t6) {
    M(this, t6);
  }
};
var Z = { M: h2, P: o3, A: n3, C: 1, L: N, R: k, D: d2, V: M, I: R, H, N: L, U: z, B: I, F: W };
var j = t2.litHtmlPolyfillSupport;
j?.(S2, R), (t2.litHtmlVersions ??= []).push("3.3.2");
var B = (t6, i6, s58) => {
  const e6 = s58?.renderBefore ?? i6;
  let h3 = e6._$litPart$;
  if (void 0 === h3) {
    const t7 = s58?.renderBefore ?? null;
    e6._$litPart$ = h3 = new R(i6.insertBefore(c3(), t7), t7, void 0, s58 ?? {});
  }
  return h3._$AI(t6), h3;
};

// elements/node_modules/lit-element/lit-element.js
var s3 = globalThis;
var i3 = class extends g {
  constructor() {
    super(...arguments), this.renderOptions = { host: this }, this._$Do = void 0;
  }
  createRenderRoot() {
    const t6 = super.createRenderRoot();
    return this.renderOptions.renderBefore ??= t6.firstChild, t6;
  }
  update(t6) {
    const r7 = this.render();
    this.hasUpdated || (this.renderOptions.isConnected = this.isConnected), super.update(t6), this._$Do = B(r7, this.renderRoot, this.renderOptions);
  }
  connectedCallback() {
    super.connectedCallback(), this._$Do?.setConnected(true);
  }
  disconnectedCallback() {
    super.disconnectedCallback(), this._$Do?.setConnected(false);
  }
  render() {
    return E;
  }
};
i3._$litElement$ = true, i3["finalized"] = true, s3.litElementHydrateSupport?.({ LitElement: i3 });
var o4 = s3.litElementPolyfillSupport;
o4?.({ LitElement: i3 });
(s3.litElementVersions ??= []).push("4.2.2");

// elements/node_modules/@lit/reactive-element/node/decorators/custom-element.js
var t3 = (t6) => (e6, o9) => {
  void 0 !== o9 ? o9.addInitializer(() => {
    customElements.define(t6, e6);
  }) : customElements.define(t6, e6);
};

// lit-css:elements/cem-color-scheme-toggle/cem-color-scheme-toggle.css
var s4 = new CSSStyleSheet();
s4.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\n#toggle-group {\\n  display: flex;\\n  gap: 2px;\\n  background: var(--pf-t--global--background--color--secondary--hover);\\n  border-radius: 6px;\\n  padding: 2px;\\n  border: 1px solid var(--pf-t--global--border--color--default);\\n}\\n\\n.toggle-option {\\n  position: relative;\\n  cursor: pointer;\\n  margin: 0;\\n}\\n\\n.toggle-option input[type=\\"radio\\"] {\\n  position: absolute;\\n  opacity: 0;\\n  width: 0;\\n  height: 0;\\n}\\n\\n.toggle-label {\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n  width: 32px;\\n  height: 32px;\\n  border-radius: 4px;\\n  color: var(--pf-t--global--icon--color--subtle);\\n  transition: all 200ms cubic-bezier(0.4, 0, 0.2, 1);\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--background--color--primary--hover);\\n    color: var(--pf-t--global--icon--color--regular);\\n  }\\n}\\n\\n.toggle-option input[type=\\"radio\\"]:checked + .toggle-label {\\n  background: var(--pf-t--global--color--brand--default);\\n  color: var(--pf-t--global--icon--color--on-brand--default);\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--color--brand--hover);\\n    color: var(--pf-t--global--icon--color--on-brand--hover);\\n  }\\n}\\n\\n.toggle-option input[type=\\"radio\\"]:focus-visible + .toggle-label {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: 2px;\\n}\\n\\n.sr-only {\\n  position: absolute;\\n  width: 1px;\\n  height: 1px;\\n  padding: 0;\\n  margin: -1px;\\n  overflow: hidden;\\n  clip: rect(0, 0, 0, 0);\\n  white-space: nowrap;\\n  border: 0;\\n}\\n"'));
var cem_color_scheme_toggle_default = s4;

// elements/cem-color-scheme-toggle/cem-color-scheme-toggle.ts
var _CemColorSchemeToggle_decorators, _init, _a;
_CemColorSchemeToggle_decorators = [t3("cem-color-scheme-toggle")];
var CemColorSchemeToggle = class extends (_a = i3) {
  static styles = cem_color_scheme_toggle_default;
  /** Storage access gatekeeper - localStorage can throw in Safari private mode */
  #getStorageItem(key, defaultValue) {
    try {
      return localStorage.getItem(key) ?? defaultValue;
    } catch {
      return defaultValue;
    }
  }
  #setStorageItem(key, value) {
    try {
      localStorage.setItem(key, value);
    } catch {
    }
  }
  connectedCallback() {
    super.connectedCallback();
    const saved = this.#getStorageItem("cem-serve-color-scheme", "system");
    this.#applyColorScheme(saved);
  }
  render() {
    const saved = this.#getStorageItem("cem-serve-color-scheme", "system");
    return T`
      <div id="toggle-group"
           role="radiogroup"
           aria-label="Color scheme">
        <label class="toggle-option">
          <input type="radio"
                 name="color-scheme"
                 value="light"
                 .checked=${saved === "light"}
                 @change=${this.#onChange}>
          <span class="toggle-label">
            <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
              <path d="M8 11a3 3 0 1 1 0-6 3 3 0 0 1 0 6zm0 1a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z"/>
            </svg>
            <span class="sr-only">Light</span>
          </span>
        </label>
        <label class="toggle-option">
          <input type="radio"
                 name="color-scheme"
                 value="system"
                 .checked=${saved === "system"}
                 @change=${this.#onChange}>
          <span class="toggle-label">
            <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
              <path d="M1 2.5A1.5 1.5 0 0 1 2.5 1h3A1.5 1.5 0 0 1 7 2.5v3A1.5 1.5 0 0 1 5.5 7h-3A1.5 1.5 0 0 1 1 5.5v-3zm8 0A1.5 1.5 0 0 1 10.5 1h3A1.5 1.5 0 0 1 15 2.5v3A1.5 1.5 0 0 1 13.5 7h-3A1.5 1.5 0 0 1 9 5.5v-3zm-8 8A1.5 1.5 0 0 1 2.5 9h3A1.5 1.5 0 0 1 7 10.5v3A1.5 1.5 0 0 1 5.5 15h-3A1.5 1.5 0 0 1 1 13.5v-3zm8 0A1.5 1.5 0 0 1 10.5 9h3a1.5 1.5 0 0 1 1.5 1.5v3a1.5 1.5 0 0 1-1.5 1.5h-3A1.5 1.5 0 0 1 9 13.5v-3z"/>
            </svg>
            <span class="sr-only">System</span>
          </span>
        </label>
        <label class="toggle-option">
          <input type="radio"
                 name="color-scheme"
                 value="dark"
                 .checked=${saved === "dark"}
                 @change=${this.#onChange}>
          <span class="toggle-label">
            <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true">
              <path d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278z"/>
            </svg>
            <span class="sr-only">Dark</span>
          </span>
        </label>
      </div>
    `;
  }
  #onChange(e6) {
    const target = e6.target;
    if (target.checked) {
      this.#applyColorScheme(target.value);
    }
  }
  #applyColorScheme(scheme) {
    switch (scheme) {
      case "light":
        document.body.style.colorScheme = "light";
        break;
      case "dark":
        document.body.style.colorScheme = "dark";
        break;
      case "system":
      default:
        document.body.style.colorScheme = "light dark";
        break;
    }
    this.#setStorageItem("cem-serve-color-scheme", scheme);
  }
};
_init = __decoratorStart(_a);
CemColorSchemeToggle = __decorateElement(_init, 0, "CemColorSchemeToggle", _CemColorSchemeToggle_decorators, CemColorSchemeToggle);
__runInitializers(_init, 1, CemColorSchemeToggle);

// elements/node_modules/@lit/reactive-element/node/decorators/property.js
var o5 = { attribute: true, type: String, converter: b, reflect: false, hasChanged: m };
var r4 = (t6 = o5, e6, r7) => {
  const { kind: n6, metadata: i6 } = r7;
  let s58 = globalThis.litPropertyMetadata.get(i6);
  if (void 0 === s58 && globalThis.litPropertyMetadata.set(i6, s58 = /* @__PURE__ */ new Map()), "setter" === n6 && ((t6 = Object.create(t6)).wrapped = true), s58.set(r7.name, t6), "accessor" === n6) {
    const { name: o9 } = r7;
    return { set(r8) {
      const n7 = e6.get.call(this);
      e6.set.call(this, r8), this.requestUpdate(o9, n7, t6, true, r8);
    }, init(e7) {
      return void 0 !== e7 && this.C(o9, void 0, t6, e7), e7;
    } };
  }
  if ("setter" === n6) {
    const { name: o9 } = r7;
    return function(r8) {
      const n7 = this[o9];
      e6.call(this, r8), this.requestUpdate(o9, n7, t6, true, r8);
    };
  }
  throw Error("Unsupported decorator location: " + n6);
};
function n4(t6) {
  return (e6, o9) => "object" == typeof o9 ? r4(t6, e6, o9) : ((t7, e7, o10) => {
    const r7 = e7.hasOwnProperty(o10);
    return e7.constructor.createProperty(o10, t7), r7 ? Object.getOwnPropertyDescriptor(e7, o10) : void 0;
  })(t6, e6, o9);
}

// lit-css:elements/cem-connection-status/cem-connection-status.css
var s5 = new CSSStyleSheet();
s5.replaceSync(JSON.parse('":host {\\n  position: fixed;\\n  bottom: 20px;\\n  right: 20px;\\n  padding: 8px 12px;\\n  border-radius: 6px;\\n  font-family: system-ui, -apple-system, sans-serif;\\n  font-size: 12px;\\n  font-weight: 500;\\n  z-index: 999999;\\n  display: flex;\\n  align-items: center;\\n  gap: 8px;\\n  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);\\n  transition: opacity 0.3s ease;\\n}\\n\\n:host([state=\\"connected\\"]) {\\n  color: #10b981;\\n  background-color: #064e3b;\\n  border: 1px solid #10b981;\\n}\\n\\n:host([state=\\"reconnecting\\"]) {\\n  color: #fbbf24;\\n  background-color: #78350f;\\n  border: 1px solid #fbbf24;\\n}\\n\\n:host([state=\\"disconnected\\"]) {\\n  color: #ef4444;\\n  background-color: #7f1d1d;\\n  border: 1px solid #ef4444;\\n}\\n\\n:host([state=\\"connected\\"][faded]) {\\n  opacity: 0.3;\\n}\\n"'));
var cem_connection_status_default = s5;

// elements/cem-connection-status/cem-connection-status.ts
var ICONS = {
  connected: "\u{1F7E2}",
  reconnecting: "\u{1F7E1}",
  disconnected: "\u{1F534}"
};
var _message_dec, _faded_dec, _state_dec, _a2, _CemConnectionStatus_decorators, _init2, _state, _faded, _message, _fadeTimeout, _hideTimeout;
_CemConnectionStatus_decorators = [t3("cem-connection-status")];
var CemConnectionStatus = class extends (_a2 = i3, _state_dec = [n4({ reflect: true })], _faded_dec = [n4({ type: Boolean, reflect: true })], _message_dec = [n4()], _a2) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _state, __runInitializers(_init2, 8, this)), __runInitializers(_init2, 11, this);
    __privateAdd(this, _faded, __runInitializers(_init2, 12, this, false)), __runInitializers(_init2, 15, this);
    __privateAdd(this, _message, __runInitializers(_init2, 16, this, "")), __runInitializers(_init2, 19, this);
    __privateAdd(this, _fadeTimeout, null);
    __privateAdd(this, _hideTimeout, null);
  }
  render() {
    return T`
      <span id="icon">${this.state ? ICONS[this.state] : ""}</span>
      <span id="message">${this.message}</span>
    `;
  }
  show(state, message, { fadeDelay = 0 } = {}) {
    if (__privateGet(this, _fadeTimeout) != null) {
      clearTimeout(__privateGet(this, _fadeTimeout));
      __privateSet(this, _fadeTimeout, null);
    }
    if (__privateGet(this, _hideTimeout) != null) {
      clearTimeout(__privateGet(this, _hideTimeout));
      __privateSet(this, _hideTimeout, null);
    }
    this.state = state;
    this.faded = false;
    this.message = message;
    if (fadeDelay > 0 && state === "connected") {
      __privateSet(this, _fadeTimeout, setTimeout(() => {
        this.faded = true;
      }, fadeDelay));
    }
  }
  hide() {
    this.style.opacity = "0";
    __privateSet(this, _hideTimeout, setTimeout(() => this.remove(), 300));
  }
};
_init2 = __decoratorStart(_a2);
_state = new WeakMap();
_faded = new WeakMap();
_message = new WeakMap();
_fadeTimeout = new WeakMap();
_hideTimeout = new WeakMap();
__decorateElement(_init2, 4, "state", _state_dec, CemConnectionStatus, _state);
__decorateElement(_init2, 4, "faded", _faded_dec, CemConnectionStatus, _faded);
__decorateElement(_init2, 4, "message", _message_dec, CemConnectionStatus, _message);
CemConnectionStatus = __decorateElement(_init2, 0, "CemConnectionStatus", _CemConnectionStatus_decorators, CemConnectionStatus);
__publicField(CemConnectionStatus, "styles", cem_connection_status_default);
__runInitializers(_init2, 1, CemConnectionStatus);

// elements/node_modules/@lit/reactive-element/node/decorators/state.js
function r5(r7) {
  return n4({ ...r7, state: true, attribute: false });
}

// elements/node_modules/lit-html/node/directive.js
var t4 = { ATTRIBUTE: 1, CHILD: 2, PROPERTY: 3, BOOLEAN_ATTRIBUTE: 4, EVENT: 5, ELEMENT: 6 };
var e3 = (t6) => (...e6) => ({ _$litDirective$: t6, values: e6 });
var i4 = class {
  constructor(t6) {
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AT(t6, e6, i6) {
    this._$Ct = t6, this._$AM = e6, this._$Ci = i6;
  }
  _$AS(t6, e6) {
    return this.update(t6, e6);
  }
  update(t6, e6) {
    return this.render(...e6);
  }
};

// elements/node_modules/lit-html/node/directives/unsafe-html.js
var e4 = class extends i4 {
  constructor(i6) {
    if (super(i6), this.it = A, i6.type !== t4.CHILD) throw Error(this.constructor.directiveName + "() can only be used in child bindings");
  }
  render(r7) {
    if (r7 === A || null == r7) return this._t = void 0, this.it = r7;
    if (r7 === E) return r7;
    if ("string" != typeof r7) throw Error(this.constructor.directiveName + "() called with a non-string value");
    if (r7 === this.it) return this._t;
    this.it = r7;
    const s58 = [r7];
    return s58.raw = s58, this._t = { _$litType$: this.constructor.resultType, strings: s58, values: [] };
  }
};
e4.directiveName = "unsafeHTML", e4.resultType = 1;
var o6 = e3(e4);

// lit-css:elements/cem-detail-panel/cem-detail-panel.css
var s6 = new CSSStyleSheet();
s6.replaceSync(JSON.parse('"/* Detail Panel - Displays manifest item details on-demand */\\n\\n:host {\\n  display: block;\\n  padding: var(--pf-t--global--spacer--md, 1rem);\\n}\\n\\n#details-content {\\n  h3 {\\n    margin-top: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n    font-size: var(--pf-t--global--font--size--heading--lg, 1.25rem);\\n  }\\n\\n  dl {\\n    margin: 0;\\n  }\\n\\n  dt {\\n    font-weight: var(--pf-t--global--font--weight--body--bold, 600);\\n    margin-top: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  dd {\\n    margin-left: 0;\\n    margin-bottom: var(--cem-pf-v6-c-description-list--m-horizontal__description--MarginBlockEnd, 0.5rem);\\n  }\\n\\n  .empty-state {\\n    color: var(--pf-t--global--text--color--subtle, #6c757d);\\n    font-style: italic;\\n  }\\n\\n  code {\\n    background-color: var(--pf-t--global--background--color--secondary--default, #f5f5f5);\\n    padding: 0.125rem 0.25rem;\\n    border-radius: 0.25rem;\\n    font-family: var(--pf-t--global--font--family--mono, monospace);\\n    font-size: 0.875em;\\n  }\\n}\\n"'));
var cem_detail_panel_default = s6;

// elements/cem-detail-panel/cem-detail-panel.ts
var _contentHTML_dec, _a3, _CemDetailPanel_decorators, _init3, _contentHTML, _b, contentHTML_get, contentHTML_set, _CemDetailPanel_instances, _cache, getItemId_fn, buildDetailHTML_fn, buildPackageDetails_fn, buildModuleDetails_fn, buildCustomElementDetails_fn, buildAttributeDetails_fn, buildPropertyDetails_fn, buildMethodDetails_fn, buildEventDetails_fn, buildSlotDetails_fn, buildCSSPropertyDetails_fn, buildCSSPartDetails_fn, buildCSSStateDetails_fn, buildDemoDetails_fn, buildClassDetails_fn, buildFunctionDetails_fn, buildVariableDetails_fn, buildMixinDetails_fn, renderMarkdown_fn, findModule_fn, findCustomElement_fn, findDeclaration_fn, escapeHtml_fn;
_CemDetailPanel_decorators = [t3("cem-detail-panel")];
var CemDetailPanel = class extends (_a3 = i3, _contentHTML_dec = [r5()], _a3) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemDetailPanel_instances);
    __privateAdd(this, _contentHTML, __runInitializers(_init3, 8, this, '<div class="empty-state">Select an item to view details</div>')), __runInitializers(_init3, 11, this);
    __privateAdd(this, _cache, /* @__PURE__ */ new Map());
  }
  render() {
    return T`
      <div id="details-content">
        ${o6(__privateGet(this, _CemDetailPanel_instances, contentHTML_get))}
      </div>
    `;
  }
  /**
   * Render details for a manifest item
   */
  async renderItem(item, manifest) {
    if (!item) {
      __privateSet(this, _CemDetailPanel_instances, '<div class="empty-state">Select an item to view details</div>', contentHTML_set);
      return;
    }
    const itemId = __privateMethod(this, _CemDetailPanel_instances, getItemId_fn).call(this, item);
    if (__privateGet(this, _cache).has(itemId)) {
      __privateSet(this, _CemDetailPanel_instances, __privateGet(this, _cache).get(itemId), contentHTML_set);
      return;
    }
    const contentHTML = await __privateMethod(this, _CemDetailPanel_instances, buildDetailHTML_fn).call(this, item, manifest);
    __privateGet(this, _cache).set(itemId, contentHTML);
    __privateSet(this, _CemDetailPanel_instances, contentHTML, contentHTML_set);
  }
};
_init3 = __decoratorStart(_a3);
_contentHTML = new WeakMap();
_CemDetailPanel_instances = new WeakSet();
_cache = new WeakMap();
/**
 * Generate a unique ID for a manifest item
 */
getItemId_fn = function(item) {
  const parts = [item.type];
  if (item.modulePath) parts.push(item.modulePath);
  if (item.tagName) parts.push(item.tagName);
  if (item.name) parts.push(item.name);
  return parts.join(":");
};
buildDetailHTML_fn = async function(item, manifest) {
  const type = item.type;
  switch (type) {
    case "package":
      return __privateMethod(this, _CemDetailPanel_instances, buildPackageDetails_fn).call(this, item, manifest);
    case "module":
      return __privateMethod(this, _CemDetailPanel_instances, buildModuleDetails_fn).call(this, item, manifest);
    case "custom-element":
      return __privateMethod(this, _CemDetailPanel_instances, buildCustomElementDetails_fn).call(this, item, manifest);
    case "category":
    case "group":
      return '<div class="empty-state">Select an item to view details</div>';
    case "attribute":
      return __privateMethod(this, _CemDetailPanel_instances, buildAttributeDetails_fn).call(this, item, manifest);
    case "property":
      return __privateMethod(this, _CemDetailPanel_instances, buildPropertyDetails_fn).call(this, item, manifest);
    case "method":
      return __privateMethod(this, _CemDetailPanel_instances, buildMethodDetails_fn).call(this, item, manifest);
    case "event":
      return __privateMethod(this, _CemDetailPanel_instances, buildEventDetails_fn).call(this, item, manifest);
    case "slot":
      return __privateMethod(this, _CemDetailPanel_instances, buildSlotDetails_fn).call(this, item, manifest);
    case "css-property":
      return __privateMethod(this, _CemDetailPanel_instances, buildCSSPropertyDetails_fn).call(this, item, manifest);
    case "css-part":
      return __privateMethod(this, _CemDetailPanel_instances, buildCSSPartDetails_fn).call(this, item, manifest);
    case "css-state":
      return __privateMethod(this, _CemDetailPanel_instances, buildCSSStateDetails_fn).call(this, item, manifest);
    case "demo":
      return __privateMethod(this, _CemDetailPanel_instances, buildDemoDetails_fn).call(this, item, manifest);
    case "class":
      return __privateMethod(this, _CemDetailPanel_instances, buildClassDetails_fn).call(this, item, manifest);
    case "function":
      return __privateMethod(this, _CemDetailPanel_instances, buildFunctionDetails_fn).call(this, item, manifest);
    case "variable":
      return __privateMethod(this, _CemDetailPanel_instances, buildVariableDetails_fn).call(this, item, manifest);
    case "mixin":
      return __privateMethod(this, _CemDetailPanel_instances, buildMixinDetails_fn).call(this, item, manifest);
    default:
      return `<div class="empty-state">No details available for ${type}</div>`;
  }
};
buildPackageDetails_fn = async function(item, manifest) {
  const pkg = manifest.packages?.find((p4) => p4.name === item.packageName);
  if (!pkg) return '<div class="empty-state">Package not found</div>';
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, pkg.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Type</dt><dd class="cem-pf-v6-c-description-list__description">Package</dd></div>
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Modules</dt><dd class="cem-pf-v6-c-description-list__description">${pkg.modules?.length || 0}</dd></div>
        ${pkg.schemaVersion ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Schema Version</dt><dd class="cem-pf-v6-c-description-list__description">${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, pkg.schemaVersion)}</dd></div>` : ""}
      </dl>
    `;
};
buildModuleDetails_fn = async function(item, manifest) {
  const module = __privateMethod(this, _CemDetailPanel_instances, findModule_fn).call(this, manifest, item.modulePath);
  if (!module) return '<div class="empty-state">Module not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = module.summary ? `modules.#(path=="${escapedPath}").summary` : "";
  const descriptionPath = module.description ? `modules.#(path=="${escapedPath}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, module.path)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Declarations</dt><dd class="cem-pf-v6-c-description-list__description">${module.declarations?.length || 0}</dd></div>
      </dl>
    `;
};
buildCustomElementDetails_fn = async function(item, manifest) {
  if (!item.tagName) {
    return '<div class="empty-state">Invalid custom element (missing tagName)</div>';
  }
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = ce.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").summary` : "";
  const descriptionPath = ce.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Name</dt><dd class="cem-pf-v6-c-description-list__description">${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.name)}</dd></div>
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        ${ce.superclass?.name ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Extends</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.superclass.name)}</code></dd></div>` : ""}
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Attributes</dt><dd class="cem-pf-v6-c-description-list__description">${ce.attributes?.length || 0}</dd></div>
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Events</dt><dd class="cem-pf-v6-c-description-list__description">${ce.events?.length || 0}</dd></div>
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Slots</dt><dd class="cem-pf-v6-c-description-list__description">${ce.slots?.length || 0}</dd></div>
      </dl>
    `;
};
buildAttributeDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const attr = ce.attributes?.find((a3) => a3.name === item.name);
  if (!attr) return '<div class="empty-state">Attribute not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = attr.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").attributes.#(name=="${escapedName}").summary` : "";
  const descriptionPath = attr.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").attributes.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, attr.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${attr.type?.text ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Type</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, attr.type.text)}</code></dd></div>` : ""}
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        ${attr.default ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Default</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, attr.default)}</code></dd></div>` : ""}
        ${attr.fieldName ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Reflects to</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, attr.fieldName)}</code></dd></div>` : ""}
      </dl>
    `;
};
buildPropertyDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const prop = ce.members?.find((m4) => m4.kind === "field" && m4.name === item.name);
  if (!prop) return '<div class="empty-state">Property not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = prop.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").members.#(name=="${escapedName}").summary` : "";
  const descriptionPath = prop.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").members.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, prop.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${prop.type?.text ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Type</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, prop.type.text)}</code></dd></div>` : ""}
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        ${prop.default ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Default</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, prop.default)}</code></dd></div>` : ""}
        ${prop.privacy ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Privacy</dt><dd class="cem-pf-v6-c-description-list__description">${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, prop.privacy)}</dd></div>` : ""}
      </dl>
    `;
};
buildMethodDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const method = ce.members?.find((m4) => m4.kind === "method" && m4.name === item.name);
  if (!method) return '<div class="empty-state">Method not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = method.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").members.#(name=="${escapedName}").summary` : "";
  const descriptionPath = method.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").members.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, method.name)}()</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        ${method.return?.type?.text ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Returns</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, method.return.type.text)}</code></dd></div>` : ""}
        ${method.privacy ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Privacy</dt><dd class="cem-pf-v6-c-description-list__description">${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, method.privacy)}</dd></div>` : ""}
      </dl>
    `;
};
buildEventDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const event = ce.events?.find((e6) => e6.name === item.name);
  if (!event) return '<div class="empty-state">Event not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = event.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").events.#(name=="${escapedName}").summary` : "";
  const descriptionPath = event.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").events.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, event.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${event.type?.text ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Detail Type</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, event.type.text)}</code></dd></div>` : ""}
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
      </dl>
    `;
};
buildSlotDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const slot = ce.slots?.find((s58) => s58.name === item.name);
  if (!slot) return '<div class="empty-state">Slot not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = slot.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").slots.#(name=="${escapedName}").summary` : "";
  const descriptionPath = slot.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").slots.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, slot.name) || "(default)"}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
      </dl>
    `;
};
buildCSSPropertyDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const cssProp = ce.cssProperties?.find((p4) => p4.name === item.name);
  if (!cssProp) return '<div class="empty-state">CSS property not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = cssProp.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssProperties.#(name=="${escapedName}").summary` : "";
  const descriptionPath = cssProp.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssProperties.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, cssProp.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${cssProp.syntax ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Syntax</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, cssProp.syntax)}</code></dd></div>` : ""}
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        ${cssProp.default ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Default</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, cssProp.default)}</code></dd></div>` : ""}
      </dl>
    `;
};
buildCSSPartDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const cssPart = ce.cssParts?.find((p4) => p4.name === item.name);
  if (!cssPart) return '<div class="empty-state">CSS part not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = cssPart.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssParts.#(name=="${escapedName}").summary` : "";
  const descriptionPath = cssPart.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssParts.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, cssPart.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
      </dl>
    `;
};
buildCSSStateDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const cssState = ce.cssStates?.find((s58) => s58.name === item.name);
  if (!cssState) return '<div class="empty-state">CSS state not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = cssState.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssStates.#(name=="${escapedName}").summary` : "";
  const descriptionPath = cssState.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssStates.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>:--${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, cssState.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
      </dl>
    `;
};
buildDemoDetails_fn = async function(item, manifest) {
  const ce = __privateMethod(this, _CemDetailPanel_instances, findCustomElement_fn).call(this, manifest, item.modulePath, item.tagName);
  if (!ce) return '<div class="empty-state">Custom element not found</div>';
  const demo = ce.demos?.find((d3) => d3.url === item.url);
  if (!demo) return '<div class="empty-state">Demo not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedTagName = item.tagName.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedUrl = demo.url.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const descriptionPath = demo.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").demos.#(url=="${escapedUrl}").description` : "";
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>Demo</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Element</dt><dd class="cem-pf-v6-c-description-list__description"><code>&lt;${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, ce.tagName)}&gt;</code></dd></div>
        ${demo.url ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">URL</dt><dd class="cem-pf-v6-c-description-list__description"><a href="${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, demo.url)}" target="_blank" rel="noopener noreferrer">${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, demo.url)}</a></dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
      </dl>
    `;
};
buildClassDetails_fn = async function(item, manifest) {
  const cls = __privateMethod(this, _CemDetailPanel_instances, findDeclaration_fn).call(this, manifest, item.modulePath, item.name, "class");
  if (!cls) return '<div class="empty-state">Class not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = cls.summary ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").summary` : "";
  const descriptionPath = cls.description ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, cls.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Type</dt><dd class="cem-pf-v6-c-description-list__description">Class</dd></div>
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        ${cls.superclass?.name ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Extends</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, cls.superclass.name)}</code></dd></div>` : ""}
      </dl>
    `;
};
buildFunctionDetails_fn = async function(item, manifest) {
  const func = __privateMethod(this, _CemDetailPanel_instances, findDeclaration_fn).call(this, manifest, item.modulePath, item.name, "function");
  if (!func) return '<div class="empty-state">Function not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = func.summary ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").summary` : "";
  const descriptionPath = func.description ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, func.name)}()</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Type</dt><dd class="cem-pf-v6-c-description-list__description">Function</dd></div>
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        ${func.return?.type?.text ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Returns</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, func.return.type.text)}</code></dd></div>` : ""}
      </dl>
    `;
};
buildVariableDetails_fn = async function(item, manifest) {
  const variable = __privateMethod(this, _CemDetailPanel_instances, findDeclaration_fn).call(this, manifest, item.modulePath, item.name, "variable");
  if (!variable) return '<div class="empty-state">Variable not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = variable.summary ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").summary` : "";
  const descriptionPath = variable.description ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, variable.name)}</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Type</dt><dd class="cem-pf-v6-c-description-list__description">Variable</dd></div>
        ${variable.type?.text ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Value Type</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, variable.type.text)}</code></dd></div>` : ""}
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
        ${variable.default ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Default</dt><dd class="cem-pf-v6-c-description-list__description"><code>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, variable.default)}</code></dd></div>` : ""}
      </dl>
    `;
};
buildMixinDetails_fn = async function(item, manifest) {
  const mixin = __privateMethod(this, _CemDetailPanel_instances, findDeclaration_fn).call(this, manifest, item.modulePath, item.name, "mixin");
  if (!mixin) return '<div class="empty-state">Mixin not found</div>';
  const escapedPath = item.modulePath.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const escapedName = item.name.replace(/\\/g, "\\\\").replace(/"/g, '\\"');
  const summaryPath = mixin.summary ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").summary` : "";
  const descriptionPath = mixin.description ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").description` : "";
  const summary = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, summaryPath);
  const description = await __privateMethod(this, _CemDetailPanel_instances, renderMarkdown_fn).call(this, descriptionPath);
  return `
      <h3>${__privateMethod(this, _CemDetailPanel_instances, escapeHtml_fn).call(this, mixin.name)}()</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Type</dt><dd class="cem-pf-v6-c-description-list__description">Mixin</dd></div>
        ${summary ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Summary</dt><dd class="cem-pf-v6-c-description-list__description">${summary}</dd></div>` : ""}
        ${description ? `<div class="cem-pf-v6-c-description-list__group"><dt class="cem-pf-v6-c-description-list__term">Description</dt><dd class="cem-pf-v6-c-description-list__description">${description}</dd></div>` : ""}
      </dl>
    `;
};
renderMarkdown_fn = async function(path) {
  if (!path) return "";
  try {
    const response = await fetch("/__cem/api/markdown", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ path })
    });
    if (!response.ok) {
      console.error("Markdown API returned error:", response.status);
      return "";
    }
    const data = await response.json();
    return data.html;
  } catch (error) {
    console.error("Failed to render markdown:", error);
    return "";
  }
};
/**
 * Find a module in the manifest.
 * Handles both single-package format (manifest.modules) and
 * workspace format (manifest.packages[*].modules)
 */
findModule_fn = function(manifest, modulePath) {
  if (manifest.modules) {
    const module = manifest.modules.find((m4) => m4.path === modulePath);
    if (module) return module;
  }
  if (manifest.packages) {
    for (const pkg of manifest.packages) {
      const module = pkg.modules?.find((m4) => m4.path === modulePath);
      if (module) return module;
    }
  }
  return null;
};
/**
 * Find a custom element in the manifest
 */
findCustomElement_fn = function(manifest, modulePath, tagName) {
  const module = __privateMethod(this, _CemDetailPanel_instances, findModule_fn).call(this, manifest, modulePath);
  if (!module) return null;
  return module.declarations?.find((d3) => d3.kind === "class" && d3.customElement && d3.tagName === tagName) ?? null;
};
/**
 * Find a declaration (class, function, variable, mixin) in the manifest
 */
findDeclaration_fn = function(manifest, modulePath, name, kind) {
  const module = __privateMethod(this, _CemDetailPanel_instances, findModule_fn).call(this, manifest, modulePath);
  if (!module) return null;
  return module.declarations?.find((d3) => d3.kind === kind && d3.name === name) ?? null;
};
/**
 * Escape HTML to prevent XSS
 */
escapeHtml_fn = function(str) {
  if (!str) return "";
  return String(str).replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;").replace(/'/g, "&#039;");
};
_b = __decorateElement(_init3, 20, "#contentHTML", _contentHTML_dec, _CemDetailPanel_instances, _contentHTML), contentHTML_get = _b.get, contentHTML_set = _b.set;
CemDetailPanel = __decorateElement(_init3, 0, "CemDetailPanel", _CemDetailPanel_decorators, CemDetailPanel);
__publicField(CemDetailPanel, "styles", cem_detail_panel_default);
__runInitializers(_init3, 1, CemDetailPanel);

// lit-css:elements/cem-drawer/cem-drawer.css
var s7 = new CSSStyleSheet();
s7.replaceSync(JSON.parse(`":host {\\n  display: flex;\\n  flex-direction: column;\\n}\\n\\n#resize-handle {\\n  height: 4px;\\n  width: 100%;\\n  cursor: ns-resize;\\n  background: transparent;\\n  position: relative;\\n  display: none;\\n\\n  :host([open]) \\u0026 {\\n    display: block;\\n  }\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    top: 50%;\\n    left: 50%;\\n    transform: translate(-50%, -50%);\\n    width: 40px;\\n    height: 4px;\\n    background: var(--pf-t--global--border--color--default);\\n    border-radius: var(--pf-t--global--border--radius--small);\\n    transition: background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n  }\\n\\n  \\u0026:hover::before,\\n  \\u0026:active::before {\\n    background: var(--pf-t--global--color--brand--default);\\n  }\\n\\n  \\u0026:focus-visible {\\n    outline: 2px solid var(--pf-t--global--color--brand--default);\\n    outline-offset: 2px;\\n  }\\n\\n  \\u0026:focus-visible::before {\\n    background: var(--pf-t--global--color--brand--default);\\n  }\\n}\\n\\n#toggle {\\n  background: var(--pf-t--global--background--color--action--plain--default);\\n  border: var(--pf-t--global--border--width--action--plain--hover) solid transparent;\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  color: var(--pf-t--global--text--color--subtle);\\n  padding: var(--pf-t--global--spacer--xs) var(--pf-t--global--spacer--sm);\\n  cursor: pointer;\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n  width: 100%;\\n  transition:\\n    background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default),\\n    color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default),\\n    transform var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--background--color--action--plain--hover);\\n    color: var(--pf-t--global--text--color--regular);\\n    border-color: var(--pf-t--global--border--color--high-contrast);\\n  }\\n\\n  \\u0026 svg {\\n    rotate: 180deg;\\n    transition: rotate var(--pf-t--global--motion--duration--md) var(--pf-t--global--motion--timing-function--decelerate);\\n  }\\n\\n  :host([open]) \\u0026 svg {\\n    rotate: 0deg;\\n  }\\n}\\n\\n#content {\\n  view-transition-name: dev-server-drawer;\\n  height: 0;\\n  overflow: hidden;\\n  display: flex;\\n  flex-direction: column;\\n  opacity: 0;\\n  contain: layout style paint;\\n\\n  \\u0026.transitions-enabled {\\n    transition:\\n      height var(--pf-t--global--motion--duration--slide-in--short) var(--pf-t--global--motion--timing-function--decelerate),\\n      opacity var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default);\\n  }\\n\\n  \\u0026.resizing {\\n    will-change: height;\\n    contain: layout style paint size;\\n  }\\n\\n  :host([open]) \\u0026 {\\n    opacity: 1;\\n  }\\n}\\n"`));
var cem_drawer_default = s7;

// lit-css:elements/cem-pf-v6-button/cem-pf-v6-button.css
var s8 = new CSSStyleSheet();
s8.replaceSync(JSON.parse('":host {\\n  display: var(--cem-pf-v6-c-button--Display, inline-flex);\\n\\n  /* Base button variables */\\n  --cem-pf-v6-c-button--Display: inline-flex;\\n  --cem-pf-v6-c-button--AlignItems: baseline;\\n  --cem-pf-v6-c-button--JustifyContent: center;\\n  --cem-pf-v6-c-button--Gap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-button--MinWidth: calc(1em * var(--cem-pf-v6-c-button--LineHeight) + var(--cem-pf-v6-c-button--PaddingBlockStart) + var(--cem-pf-v6-c-button--PaddingBlockEnd));\\n  --cem-pf-v6-c-button--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--pf-t--global--spacer--action--horizontal--default);\\n  --cem-pf-v6-c-button--PaddingBlockEnd: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--pf-t--global--spacer--action--horizontal--default);\\n  --cem-pf-v6-c-button--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-button--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-button--FontWeight: var(--pf-t--global--font--weight--body--default);\\n  --cem-pf-v6-c-button--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-button--BackgroundColor: transparent;\\n  --cem-pf-v6-c-button--BorderColor: transparent;\\n  --cem-pf-v6-c-button--BorderWidth: var(--pf-t--global--border--width--action--default);\\n  --cem-pf-v6-c-button--BorderRadius: var(--pf-t--global--border--radius--pill);\\n  --cem-pf-v6-c-button--TextDecorationLine: none;\\n  --cem-pf-v6-c-button--TextDecorationStyle: none;\\n  --cem-pf-v6-c-button--TextDecorationColor: currentcolor;\\n  --cem-pf-v6-c-button--TransitionDelay: 0s;\\n  --cem-pf-v6-c-button--TransitionDuration: var(--pf-t--global--motion--duration--fade--default);\\n  --cem-pf-v6-c-button--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--accelerate);\\n  --cem-pf-v6-c-button--TransitionProperty: color, background, border-color;\\n  --cem-pf-v6-c-button--ScaleX: 1;\\n  --cem-pf-v6-c-button--ScaleY: 1;\\n  --cem-pf-v6-c-button--border--offset: 0;\\n  --cem-pf-v6-c-button--hover--BackgroundColor: transparent;\\n  --cem-pf-v6-c-button--hover--BorderColor: transparent;\\n  --cem-pf-v6-c-button--hover--BorderWidth: var(--pf-t--global--border--width--action--hover);\\n  --cem-pf-v6-c-button--hover--ScaleX: 1;\\n  --cem-pf-v6-c-button--hover--ScaleY: 1;\\n  --cem-pf-v6-c-button--m-clicked--BackgroundColor: transparent;\\n  --cem-pf-v6-c-button--m-clicked--BorderColor: transparent;\\n  --cem-pf-v6-c-button--m-clicked--BorderWidth: var(--pf-t--global--border--width--action--clicked);\\n  --cem-pf-v6-c-button--disabled--Color: var(--pf-t--global--text--color--on-disabled);\\n  --cem-pf-v6-c-button--disabled--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n  --cem-pf-v6-c-button--disabled--BorderColor: transparent;\\n  --cem-pf-v6-c-button__icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-button--hover__icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-button--m-clicked__icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-button--disabled__icon--Color: var(--pf-t--global--icon--color--on-disabled);\\n\\n  /* Primary variant variables */\\n  --cem-pf-v6-c-button--m-primary--Color: var(--pf-t--global--text--color--on-brand--default);\\n  --cem-pf-v6-c-button--m-primary--BackgroundColor: var(--pf-t--global--color--brand--default);\\n  --cem-pf-v6-c-button--m-primary__icon--Color: var(--pf-t--global--icon--color--on-brand--default);\\n  --cem-pf-v6-c-button--m-primary--hover--Color: var(--pf-t--global--text--color--on-brand--hover);\\n  --cem-pf-v6-c-button--m-primary--hover--BackgroundColor: var(--pf-t--global--color--brand--hover);\\n  --cem-pf-v6-c-button--m-primary--hover__icon--Color: var(--pf-t--global--icon--color--on-brand--hover);\\n  --cem-pf-v6-c-button--m-primary--m-clicked--Color: var(--pf-t--global--text--color--on-brand--clicked);\\n  --cem-pf-v6-c-button--m-primary--m-clicked--BackgroundColor: var(--pf-t--global--color--brand--clicked);\\n  --cem-pf-v6-c-button--m-primary--m-clicked__icon--Color: var(--pf-t--global--icon--color--on-brand--clicked);\\n\\n  /* Secondary variant variables */\\n  --cem-pf-v6-c-button--m-secondary--Color: var(--pf-t--global--text--color--brand--default);\\n  --cem-pf-v6-c-button--m-secondary--BorderColor: var(--pf-t--global--border--color--brand--default);\\n  --cem-pf-v6-c-button--m-secondary--TransitionDuration: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-button--m-secondary__icon--Color: var(--pf-t--global--icon--color--brand--default);\\n  --cem-pf-v6-c-button--m-secondary--hover--Color: var(--pf-t--global--text--color--brand--hover);\\n  --cem-pf-v6-c-button--m-secondary--hover--BorderColor: var(--pf-t--global--border--color--brand--hover);\\n  --cem-pf-v6-c-button--m-secondary--hover__icon--Color: var(--pf-t--global--icon--color--brand--hover);\\n  --cem-pf-v6-c-button--m-secondary--m-clicked--Color: var(--pf-t--global--text--color--brand--clicked);\\n  --cem-pf-v6-c-button--m-secondary--m-clicked--BorderColor: var(--pf-t--global--border--color--brand--clicked);\\n  --cem-pf-v6-c-button--m-secondary--m-clicked__icon--Color: var(--pf-t--global--icon--color--brand--clicked);\\n\\n  /* Tertiary variant variables */\\n  --cem-pf-v6-c-button--m-tertiary--Color: var(--pf-t--global--text--color--brand--default);\\n  --cem-pf-v6-c-button--m-tertiary--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-button--m-tertiary--TransitionDuration: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-button--m-tertiary__icon--Color: var(--pf-t--global--icon--color--brand--default);\\n  --cem-pf-v6-c-button--m-tertiary--hover--Color: var(--pf-t--global--text--color--brand--hover);\\n  --cem-pf-v6-c-button--m-tertiary--hover--BorderColor: var(--pf-t--global--border--color--hover);\\n  --cem-pf-v6-c-button--m-tertiary--hover__icon--Color: var(--pf-t--global--icon--color--brand--hover);\\n  --cem-pf-v6-c-button--m-tertiary--m-clicked--Color: var(--pf-t--global--text--color--brand--clicked);\\n  --cem-pf-v6-c-button--m-tertiary--m-clicked--BorderColor: var(--pf-t--global--border--color--clicked);\\n  --cem-pf-v6-c-button--m-tertiary--m-clicked__icon--Color: var(--pf-t--global--icon--color--brand--clicked);\\n\\n  /* Plain variant variables */\\n  --cem-pf-v6-c-button--m-plain--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-button--m-plain--PaddingInlineEnd: var(--pf-t--global--spacer--action--horizontal--plain--default);\\n  --cem-pf-v6-c-button--m-plain--PaddingInlineStart: var(--pf-t--global--spacer--action--horizontal--plain--default);\\n  --cem-pf-v6-c-button--m-plain--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-button--m-plain__icon--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-button--m-plain--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-button--m-plain--hover--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-button--m-plain--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --cem-pf-v6-c-button--m-plain--m-clicked--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-button--m-plain--m-clicked--BackgroundColor: var(--pf-t--global--background--color--action--plain--clicked);\\n  --cem-pf-v6-c-button--m-plain--disabled--Color: var(--pf-t--global--text--color--disabled);\\n  --cem-pf-v6-c-button--m-plain--disabled__icon--Color: var(--pf-t--global--icon--color--disabled);\\n  --cem-pf-v6-c-button--m-plain--disabled--BackgroundColor: transparent;\\n  --cem-pf-v6-c-button--m-plain--m-small--PaddingInlineEnd: var(--pf-t--global--spacer--action--horizontal--plain--compact);\\n  --cem-pf-v6-c-button--m-plain--m-small--PaddingInlineStart: var(--pf-t--global--spacer--action--horizontal--plain--compact);\\n  --cem-pf-v6-c-button--m-plain--BorderWidth: var(--pf-t--global--border--width--action--plain--default);\\n  --cem-pf-v6-c-button--m-plain--m-clicked--BorderWidth: var(--pf-t--global--border--width--action--plain--clicked);\\n  --cem-pf-v6-c-button--m-plain--hover--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --cem-pf-v6-c-button--m-plain--BorderColor: transparent;\\n  --cem-pf-v6-c-button--m-plain--m-clicked--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-button--m-plain--hover--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-button--m-plain--m-no-padding--MinWidth: auto;\\n  --cem-pf-v6-c-button--m-plain--m-no-padding--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-button--m-plain--m-no-padding--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-button--m-plain--m-no-padding--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-button--m-plain--m-no-padding--PaddingInlineStart: 0;\\n\\n  /* Link variant variables */\\n  --cem-pf-v6-c-button--m-link--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-button--m-link--PaddingInlineEnd: var(--pf-t--global--spacer--action--horizontal--plain--default);\\n  --cem-pf-v6-c-button--m-link--PaddingInlineStart: var(--pf-t--global--spacer--action--horizontal--plain--default);\\n  --cem-pf-v6-c-button--m-link--Color: var(--pf-t--global--text--color--brand--default);\\n  --cem-pf-v6-c-button--m-link--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-button--m-link__icon--Color: var(--pf-t--global--icon--color--brand--default);\\n  --cem-pf-v6-c-button--m-link--hover--Color: var(--pf-t--global--text--color--brand--hover);\\n  --cem-pf-v6-c-button--m-link--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --cem-pf-v6-c-button--m-link--hover__icon--Color: var(--pf-t--global--text--color--brand--hover);\\n  --cem-pf-v6-c-button--m-link--m-clicked--Color: var(--pf-t--global--text--color--brand--clicked);\\n  --cem-pf-v6-c-button--m-link--m-clicked--BackgroundColor: var(--pf-t--global--background--color--action--plain--clicked);\\n  --cem-pf-v6-c-button--m-link--m-clicked__icon--Color: var(--pf-t--global--text--color--brand--clicked);\\n\\n  /* Danger variant variables */\\n  --cem-pf-v6-c-button--m-danger--Color: var(--pf-t--global--text--color--status--on-danger--default);\\n  --cem-pf-v6-c-button--m-danger--BackgroundColor: var(--pf-t--global--color--status--danger--default);\\n  --cem-pf-v6-c-button--m-danger__icon--Color: var(--pf-t--global--icon--color--status--on-danger--default);\\n  --cem-pf-v6-c-button--m-danger--hover--Color: var(--pf-t--global--text--color--status--on-danger--hover);\\n  --cem-pf-v6-c-button--m-danger--hover--BackgroundColor: var(--pf-t--global--color--status--danger--hover);\\n  --cem-pf-v6-c-button--m-danger--hover__icon--Color: var(--pf-t--global--icon--color--status--on-danger--hover);\\n  --cem-pf-v6-c-button--m-danger--m-clicked--Color: var(--pf-t--global--text--color--status--on-danger--clicked);\\n  --cem-pf-v6-c-button--m-danger--m-clicked--BackgroundColor: var(--pf-t--global--color--status--danger--clicked);\\n  --cem-pf-v6-c-button--m-danger--m-clicked__icon--Color: var(--pf-t--global--icon--color--status--on-danger--clicked);\\n\\n  /* Warning variant variables */\\n  --cem-pf-v6-c-button--m-warning--Color: var(--pf-t--global--text--color--status--on-warning--default);\\n  --cem-pf-v6-c-button--m-warning--BackgroundColor: var(--pf-t--global--color--status--warning--default);\\n  --cem-pf-v6-c-button--m-warning__icon--Color: var(--pf-t--global--icon--color--status--on-warning--default);\\n  --cem-pf-v6-c-button--m-warning--hover--Color: var(--pf-t--global--text--color--status--on-warning--hover);\\n  --cem-pf-v6-c-button--m-warning--hover--BackgroundColor: var(--pf-t--global--color--status--warning--hover);\\n  --cem-pf-v6-c-button--m-warning--hover__icon--Color: var(--pf-t--global--icon--color--status--on-warning--hover);\\n  --cem-pf-v6-c-button--m-warning--m-clicked--Color: var(--pf-t--global--text--color--status--on-warning--clicked);\\n  --cem-pf-v6-c-button--m-warning--m-clicked--BackgroundColor: var(--pf-t--global--color--status--warning--clicked);\\n  --cem-pf-v6-c-button--m-warning--m-clicked__icon--Color: var(--pf-t--global--icon--color--status--on-warning--clicked);\\n\\n  /* Small size variant variables */\\n  --cem-pf-v6-c-button--m-small--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--compact);\\n  --cem-pf-v6-c-button--m-small--PaddingInlineEnd: var(--pf-t--global--spacer--action--horizontal--compact);\\n  --cem-pf-v6-c-button--m-small--PaddingBlockEnd: var(--pf-t--global--spacer--control--vertical--compact);\\n  --cem-pf-v6-c-button--m-small--PaddingInlineStart: var(--pf-t--global--spacer--action--horizontal--compact);\\n\\n  /* Large size variant variables */\\n  --cem-pf-v6-c-button--m-display-lg--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--spacious);\\n  --cem-pf-v6-c-button--m-display-lg--PaddingInlineEnd: var(--pf-t--global--spacer--action--horizontal--spacious);\\n  --cem-pf-v6-c-button--m-display-lg--PaddingBlockEnd: var(--pf-t--global--spacer--control--vertical--spacious);\\n  --cem-pf-v6-c-button--m-display-lg--PaddingInlineStart: var(--pf-t--global--spacer--action--horizontal--spacious);\\n  --cem-pf-v6-c-button--m-display-lg--FontWeight: var(--pf-t--global--font--weight--body--bold);\\n\\n  /* Block variant variables */\\n  --cem-pf-v6-c-button--m-block--Display: flex;\\n  --cem-pf-v6-c-button--m-block--Width: 100%;\\n}\\n\\n/* Shared styles for both button (host) and link (shadow a) */\\n:host(:not([href])),\\na {\\n  position: relative;\\n  display: var(--cem-pf-v6-c-button--Display);\\n  box-sizing: border-box;\\n  gap: var(--cem-pf-v6-c-button--Gap);\\n  align-items: var(--cem-pf-v6-c-button--AlignItems);\\n  justify-content: var(--cem-pf-v6-c-button--JustifyContent);\\n  min-width: var(--cem-pf-v6-c-button--MinWidth);\\n  padding-block-start: var(--cem-pf-v6-c-button--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-button--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-button--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-button--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-button--FontSize);\\n  font-weight: var(--cem-pf-v6-c-button--FontWeight);\\n  line-height: var(--cem-pf-v6-c-button--LineHeight);\\n  color: var(--cem-pf-v6-c-button--Color);\\n  text-align: center;\\n  text-decoration-line: var(--cem-pf-v6-c-button--TextDecorationLine);\\n  text-decoration-style: var(--cem-pf-v6-c-button--TextDecorationStyle);\\n  text-decoration-color: var(--cem-pf-v6-c-button--TextDecorationColor);\\n  white-space: nowrap;\\n  cursor: pointer;\\n  user-select: none;\\n  background: var(--cem-pf-v6-c-button--BackgroundColor) radial-gradient(circle, transparent 1%, color-mix(in srgb, currentcolor 15%, transparent) 2%) center/15000% 15000%;\\n  border: 0;\\n  border-start-start-radius: var(--cem-pf-v6-c-button--BorderRadius);\\n  border-start-end-radius: var(--cem-pf-v6-c-button--BorderRadius);\\n  border-end-start-radius: var(--cem-pf-v6-c-button--BorderRadius);\\n  border-end-end-radius: var(--cem-pf-v6-c-button--BorderRadius);\\n  transition-delay: var(--cem-pf-v6-c-button--TransitionDelay);\\n  transition-timing-function: var(--cem-pf-v6-c-button--TransitionTimingFunction);\\n  transition-duration: var(--cem-pf-v6-c-button--TransitionDuration);\\n  transition-property: var(--cem-pf-v6-c-button--TransitionProperty);\\n  scale: var(--cem-pf-v6-c-button--ScaleX) var(--cem-pf-v6-c-button--ScaleY);\\n}\\n\\n:host(:not([href]))::after,\\na::after {\\n  position: absolute;\\n  inset: var(--cem-pf-v6-c-button--border--offset, 0);\\n  pointer-events: none;\\n  content: \\"\\";\\n  border: var(--cem-pf-v6-c-button--BorderWidth) solid var(--cem-pf-v6-c-button--BorderColor);\\n  border-radius: inherit;\\n  transition: inherit;\\n}\\n\\n:host(:not([href]):hover:not([aria-disabled=\\"true\\"])),\\na:hover:not([aria-disabled=\\"true\\"]) {\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--hover--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-button--hover--BackgroundColor);\\n  --cem-pf-v6-c-button--BorderColor: var(--cem-pf-v6-c-button--hover--BorderColor);\\n  --cem-pf-v6-c-button--BorderWidth: var(--cem-pf-v6-c-button--hover--BorderWidth);\\n  --cem-pf-v6-c-button--ScaleX: var(--cem-pf-v6-c-button--hover--ScaleX);\\n  --cem-pf-v6-c-button--ScaleY: var(--cem-pf-v6-c-button--hover--ScaleY);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--hover__icon--Color);\\n}\\n\\n:host(:not([href]):active:not([aria-disabled=\\"true\\"])),\\n:host(:not([href]).pf-m-clicked),\\na:active:not([aria-disabled=\\"true\\"]),\\na.pf-m-clicked {\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--m-clicked--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-button--m-clicked--BackgroundColor);\\n  --cem-pf-v6-c-button--BorderColor: var(--cem-pf-v6-c-button--m-clicked--BorderColor);\\n  --cem-pf-v6-c-button--BorderWidth: var(--cem-pf-v6-c-button--m-clicked--BorderWidth);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--m-clicked__icon--Color);\\n}\\n\\n:host(:not([href]):active:not([aria-disabled=\\"true\\"])),\\na:active:not([aria-disabled=\\"true\\"]) {\\n  background-size: 100% 100%;\\n  transition: background 0s;\\n}\\n\\n:host(:not([href]):focus-visible),\\na:focus-visible {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: 2px;\\n}\\n\\n:host(:not([href])[aria-disabled=\\"true\\"]),\\na[aria-disabled=\\"true\\"] {\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--disabled--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-button--disabled--BackgroundColor);\\n  --cem-pf-v6-c-button--BorderColor: var(--cem-pf-v6-c-button--disabled--BorderColor);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--disabled__icon--Color);\\n  cursor: not-allowed;\\n  pointer-events: none;\\n}\\n\\n/* Primary variant */\\n:host([variant=\\"primary\\"]) {\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--m-primary--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-button--m-primary--BackgroundColor);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--m-primary__icon--Color);\\n  --cem-pf-v6-c-button--hover--Color: var(--cem-pf-v6-c-button--m-primary--hover--Color);\\n  --cem-pf-v6-c-button--hover--BackgroundColor: var(--cem-pf-v6-c-button--m-primary--hover--BackgroundColor);\\n  --cem-pf-v6-c-button--hover__icon--Color: var(--cem-pf-v6-c-button--m-primary--hover__icon--Color);\\n  --cem-pf-v6-c-button--m-clicked--Color: var(--cem-pf-v6-c-button--m-primary--m-clicked--Color);\\n  --cem-pf-v6-c-button--m-clicked--BackgroundColor: var(--cem-pf-v6-c-button--m-primary--m-clicked--BackgroundColor);\\n  --cem-pf-v6-c-button--m-clicked__icon--Color: var(--cem-pf-v6-c-button--m-primary--m-clicked__icon--Color);\\n}\\n\\n/* Secondary variant */\\n:host([variant=\\"secondary\\"]) {\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--m-secondary--Color);\\n  --cem-pf-v6-c-button--BorderColor: var(--cem-pf-v6-c-button--m-secondary--BorderColor);\\n  --cem-pf-v6-c-button--TransitionDuration: var(--cem-pf-v6-c-button--m-secondary--TransitionDuration);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--m-secondary__icon--Color);\\n  --cem-pf-v6-c-button--hover--Color: var(--cem-pf-v6-c-button--m-secondary--hover--Color);\\n  --cem-pf-v6-c-button--hover--BorderColor: var(--cem-pf-v6-c-button--m-secondary--hover--BorderColor);\\n  --cem-pf-v6-c-button--hover__icon--Color: var(--cem-pf-v6-c-button--m-secondary--hover__icon--Color);\\n  --cem-pf-v6-c-button--m-clicked--Color: var(--cem-pf-v6-c-button--m-secondary--m-clicked--Color);\\n  --cem-pf-v6-c-button--m-clicked--BorderColor: var(--cem-pf-v6-c-button--m-secondary--m-clicked--BorderColor);\\n  --cem-pf-v6-c-button--m-clicked__icon--Color: var(--cem-pf-v6-c-button--m-secondary--m-clicked__icon--Color);\\n}\\n\\n/* Tertiary variant */\\n:host([variant=\\"tertiary\\"]) {\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--m-tertiary--Color);\\n  --cem-pf-v6-c-button--BorderColor: var(--cem-pf-v6-c-button--m-tertiary--BorderColor);\\n  --cem-pf-v6-c-button--TransitionDuration: var(--cem-pf-v6-c-button--m-tertiary--TransitionDuration);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--m-tertiary__icon--Color);\\n  --cem-pf-v6-c-button--hover--Color: var(--cem-pf-v6-c-button--m-tertiary--hover--Color);\\n  --cem-pf-v6-c-button--hover--BorderColor: var(--cem-pf-v6-c-button--m-tertiary--hover--BorderColor);\\n  --cem-pf-v6-c-button--hover__icon--Color: var(--cem-pf-v6-c-button--m-tertiary--hover__icon--Color);\\n  --cem-pf-v6-c-button--m-clicked--Color: var(--cem-pf-v6-c-button--m-tertiary--m-clicked--Color);\\n  --cem-pf-v6-c-button--m-clicked--BorderColor: var(--cem-pf-v6-c-button--m-tertiary--m-clicked--BorderColor);\\n  --cem-pf-v6-c-button--m-clicked__icon--Color: var(--cem-pf-v6-c-button--m-tertiary--m-clicked__icon--Color);\\n}\\n\\n/* Plain variant (icon-only) */\\n:host([variant=\\"plain\\"]) {\\n  --cem-pf-v6-c-button--BorderRadius: var(--cem-pf-v6-c-button--m-plain--BorderRadius);\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--cem-pf-v6-c-button--m-plain--PaddingInlineEnd);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--cem-pf-v6-c-button--m-plain--PaddingInlineStart);\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--m-plain--Color);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--m-plain__icon--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-button--m-plain--BackgroundColor);\\n  --cem-pf-v6-c-button--hover--Color: var(--cem-pf-v6-c-button--m-plain--hover--Color);\\n  --cem-pf-v6-c-button--hover--BackgroundColor: var(--cem-pf-v6-c-button--m-plain--hover--BackgroundColor);\\n  --cem-pf-v6-c-button--m-clicked--Color: var(--cem-pf-v6-c-button--m-plain--m-clicked--Color);\\n  --cem-pf-v6-c-button--m-clicked--BackgroundColor: var(--cem-pf-v6-c-button--m-plain--m-clicked--BackgroundColor);\\n  --cem-pf-v6-c-button--disabled--Color: var(--cem-pf-v6-c-button--m-plain--disabled--Color);\\n  --cem-pf-v6-c-button--disabled__icon--Color: var(--cem-pf-v6-c-button--m-plain--disabled__icon--Color);\\n  --cem-pf-v6-c-button--disabled--BackgroundColor: var(--cem-pf-v6-c-button--m-plain--disabled--BackgroundColor);\\n  --cem-pf-v6-c-button--BorderWidth: var(--cem-pf-v6-c-button--m-plain--BorderWidth);\\n  --cem-pf-v6-c-button--m-clicked--BorderWidth: var(--cem-pf-v6-c-button--m-plain--m-clicked--BorderWidth);\\n  --cem-pf-v6-c-button--hover--BorderWidth: var(--cem-pf-v6-c-button--m-plain--hover--BorderWidth);\\n  --cem-pf-v6-c-button--BorderColor: var(--cem-pf-v6-c-button--m-plain--BorderColor);\\n  --cem-pf-v6-c-button--m-clicked--BorderColor: var(--cem-pf-v6-c-button--m-plain--m-clicked--BorderColor);\\n  --cem-pf-v6-c-button--hover--BorderColor: var(--cem-pf-v6-c-button--m-plain--hover--BorderColor);\\n}\\n\\n/* Link variant */\\n:host([variant=\\"link\\"]) {\\n  --cem-pf-v6-c-button--BorderRadius: var(--cem-pf-v6-c-button--m-link--BorderRadius);\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--cem-pf-v6-c-button--m-link--PaddingInlineEnd);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--cem-pf-v6-c-button--m-link--PaddingInlineStart);\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--m-link--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-button--m-link--BackgroundColor);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--m-link__icon--Color);\\n  --cem-pf-v6-c-button--hover--Color: var(--cem-pf-v6-c-button--m-link--hover--Color);\\n  --cem-pf-v6-c-button--hover--BackgroundColor: var(--cem-pf-v6-c-button--m-link--hover--BackgroundColor);\\n  --cem-pf-v6-c-button--hover__icon--Color: var(--cem-pf-v6-c-button--m-link--hover__icon--Color);\\n  --cem-pf-v6-c-button--m-clicked--Color: var(--cem-pf-v6-c-button--m-link--m-clicked--Color);\\n  --cem-pf-v6-c-button--m-clicked--BackgroundColor: var(--cem-pf-v6-c-button--m-link--m-clicked--BackgroundColor);\\n  --cem-pf-v6-c-button--m-clicked__icon--Color: var(--cem-pf-v6-c-button--m-link--m-clicked__icon--Color);\\n}\\n\\n/* Danger variant */\\n:host([variant=\\"danger\\"]) {\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--m-danger--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-button--m-danger--BackgroundColor);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--m-danger__icon--Color);\\n  --cem-pf-v6-c-button--hover--Color: var(--cem-pf-v6-c-button--m-danger--hover--Color);\\n  --cem-pf-v6-c-button--hover--BackgroundColor: var(--cem-pf-v6-c-button--m-danger--hover--BackgroundColor);\\n  --cem-pf-v6-c-button--hover__icon--Color: var(--cem-pf-v6-c-button--m-danger--hover__icon--Color);\\n  --cem-pf-v6-c-button--m-clicked--Color: var(--cem-pf-v6-c-button--m-danger--m-clicked--Color);\\n  --cem-pf-v6-c-button--m-clicked--BackgroundColor: var(--cem-pf-v6-c-button--m-danger--m-clicked--BackgroundColor);\\n  --cem-pf-v6-c-button--m-clicked__icon--Color: var(--cem-pf-v6-c-button--m-danger--m-clicked__icon--Color);\\n}\\n\\n/* Warning variant */\\n:host([variant=\\"warning\\"]) {\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-button--m-warning--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-button--m-warning--BackgroundColor);\\n  --cem-pf-v6-c-button__icon--Color: var(--cem-pf-v6-c-button--m-warning__icon--Color);\\n  --cem-pf-v6-c-button--hover--Color: var(--cem-pf-v6-c-button--m-warning--hover--Color);\\n  --cem-pf-v6-c-button--hover--BackgroundColor: var(--cem-pf-v6-c-button--m-warning--hover--BackgroundColor);\\n  --cem-pf-v6-c-button--hover__icon--Color: var(--cem-pf-v6-c-button--m-warning--hover__icon--Color);\\n  --cem-pf-v6-c-button--m-clicked--Color: var(--cem-pf-v6-c-button--m-warning--m-clicked--Color);\\n  --cem-pf-v6-c-button--m-clicked--BackgroundColor: var(--cem-pf-v6-c-button--m-warning--m-clicked--BackgroundColor);\\n  --cem-pf-v6-c-button--m-clicked__icon--Color: var(--cem-pf-v6-c-button--m-warning--m-clicked__icon--Color);\\n}\\n\\n/* Size variants */\\n:host([size=\\"sm\\"]) {\\n  --cem-pf-v6-c-button--PaddingBlockStart: var(--cem-pf-v6-c-button--m-small--PaddingBlockStart);\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--cem-pf-v6-c-button--m-small--PaddingInlineEnd);\\n  --cem-pf-v6-c-button--PaddingBlockEnd: var(--cem-pf-v6-c-button--m-small--PaddingBlockEnd);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--cem-pf-v6-c-button--m-small--PaddingInlineStart);\\n}\\n\\n:host([size=\\"sm\\"][variant=\\"plain\\"]) {\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--cem-pf-v6-c-button--m-plain--m-small--PaddingInlineEnd);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--cem-pf-v6-c-button--m-plain--m-small--PaddingInlineStart);\\n}\\n\\n/* Plain no-padding modifier (for icon-only buttons) */\\n:host([variant=\\"plain\\"][no-padding]) {\\n  --cem-pf-v6-c-button--MinWidth: var(--cem-pf-v6-c-button--m-plain--m-no-padding--MinWidth);\\n  --cem-pf-v6-c-button--PaddingBlockStart: var(--cem-pf-v6-c-button--m-plain--m-no-padding--PaddingBlockStart);\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--cem-pf-v6-c-button--m-plain--m-no-padding--PaddingInlineEnd);\\n  --cem-pf-v6-c-button--PaddingBlockEnd: var(--cem-pf-v6-c-button--m-plain--m-no-padding--PaddingBlockEnd);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--cem-pf-v6-c-button--m-plain--m-no-padding--PaddingInlineStart);\\n}\\n\\n:host([size=\\"lg\\"]) {\\n  --cem-pf-v6-c-button--PaddingBlockStart: var(--cem-pf-v6-c-button--m-display-lg--PaddingBlockStart);\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--cem-pf-v6-c-button--m-display-lg--PaddingInlineEnd);\\n  --cem-pf-v6-c-button--PaddingBlockEnd: var(--cem-pf-v6-c-button--m-display-lg--PaddingBlockEnd);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--cem-pf-v6-c-button--m-display-lg--PaddingInlineStart);\\n  --cem-pf-v6-c-button--FontWeight: var(--cem-pf-v6-c-button--m-display-lg--FontWeight);\\n}\\n\\n/* Block variant */\\n:host([block]) {\\n  display: var(--cem-pf-v6-c-button--m-block--Display);\\n  width: var(--cem-pf-v6-c-button--m-block--Width);\\n}\\n\\n:host([block]) button {\\n  width: 100%;\\n}\\n\\nslot {\\n  display: inline;\\n}\\n\\n/* Slot visibility controlled by custom properties set from light DOM */\\nslot[name=\\"icon-start\\"] {\\n  display: var(--_has-icon-start, inline);\\n}\\n\\nslot[name=\\"icon-end\\"] {\\n  display: var(--_has-icon-end, inline);\\n}\\n\\n/* Icon slot styling */\\n::slotted([slot=icon-start]),\\n::slotted([slot=icon-end]) {\\n  width: 1em;\\n  height: 1em;\\n  fill: currentColor;\\n  color: var(--cem-pf-v6-c-button__icon--Color);\\n}\\n"'));
var cem_pf_v6_button_default = s8;

// elements/cem-pf-v6-button/cem-pf-v6-button.ts
var _type_dec, _href_dec, _noPadding_dec, _disabled_dec, _block_dec, _size_dec, _variant_dec, _a4, _PfV6Button_decorators, _internals, _init4, _variant, _size, _block, _disabled, _noPadding, _href, _type, _PfV6Button_instances, updateMode_fn, updateDisabled_fn, _onClick, _onKeydown, onLinkClick_fn;
_PfV6Button_decorators = [t3("cem-pf-v6-button")];
var PfV6Button = class extends (_a4 = i3, _variant_dec = [n4({ reflect: true })], _size_dec = [n4({ reflect: true })], _block_dec = [n4({ type: Boolean, reflect: true })], _disabled_dec = [n4({ type: Boolean, reflect: true })], _noPadding_dec = [n4({ type: Boolean, reflect: true, attribute: "no-padding" })], _href_dec = [n4({ reflect: true })], _type_dec = [n4()], _a4) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Button_instances);
    __privateAdd(this, _internals, this.attachInternals());
    __privateAdd(this, _variant, __runInitializers(_init4, 8, this)), __runInitializers(_init4, 11, this);
    __privateAdd(this, _size, __runInitializers(_init4, 12, this)), __runInitializers(_init4, 15, this);
    __privateAdd(this, _block, __runInitializers(_init4, 16, this, false)), __runInitializers(_init4, 19, this);
    __privateAdd(this, _disabled, __runInitializers(_init4, 20, this, false)), __runInitializers(_init4, 23, this);
    __privateAdd(this, _noPadding, __runInitializers(_init4, 24, this, false)), __runInitializers(_init4, 27, this);
    __privateAdd(this, _href, __runInitializers(_init4, 28, this)), __runInitializers(_init4, 31, this);
    __privateAdd(this, _type, __runInitializers(_init4, 32, this, "button")), __runInitializers(_init4, 35, this);
    __privateAdd(this, _onClick, (event) => {
      if (this.disabled) {
        event.preventDefault();
        event.stopImmediatePropagation();
      }
    });
    __privateAdd(this, _onKeydown, (event) => {
      if (this.disabled) {
        event.preventDefault();
        return;
      }
      if (event.key === " " || event.key === "Enter") {
        event.preventDefault();
        this.click();
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    __privateMethod(this, _PfV6Button_instances, updateMode_fn).call(this);
  }
  updated(changed) {
    if (changed.has("href")) {
      __privateMethod(this, _PfV6Button_instances, updateMode_fn).call(this);
    }
    if (changed.has("disabled")) {
      __privateMethod(this, _PfV6Button_instances, updateDisabled_fn).call(this);
    }
  }
  render() {
    const slots = T`
      <slot name="icon-start"></slot>
      <slot></slot>
      <slot name="icon-end"></slot>
    `;
    if (this.href) {
      return T`<a id="link"
                      href=${this.href}
                      @click=${__privateMethod(this, _PfV6Button_instances, onLinkClick_fn)}>${slots}</a>`;
    }
    return slots;
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("click", __privateGet(this, _onClick));
    this.removeEventListener("keydown", __privateGet(this, _onKeydown));
  }
};
_init4 = __decoratorStart(_a4);
_internals = new WeakMap();
_variant = new WeakMap();
_size = new WeakMap();
_block = new WeakMap();
_disabled = new WeakMap();
_noPadding = new WeakMap();
_href = new WeakMap();
_type = new WeakMap();
_PfV6Button_instances = new WeakSet();
updateMode_fn = function() {
  if (this.href) {
    __privateGet(this, _internals).role = null;
    this.removeEventListener("click", __privateGet(this, _onClick));
    this.removeEventListener("keydown", __privateGet(this, _onKeydown));
    if (this.getAttribute("tabindex") === "0") {
      this.removeAttribute("tabindex");
    }
  } else {
    __privateGet(this, _internals).role = "button";
    if (!this.hasAttribute("tabindex")) {
      this.setAttribute("tabindex", "0");
    }
    this.addEventListener("click", __privateGet(this, _onClick));
    this.addEventListener("keydown", __privateGet(this, _onKeydown));
  }
  __privateMethod(this, _PfV6Button_instances, updateDisabled_fn).call(this);
};
updateDisabled_fn = function() {
  if (this.href) {
    const link = this.shadowRoot?.getElementById("link");
    if (link) {
      if (this.disabled) {
        link.setAttribute("aria-disabled", "true");
        link.style.pointerEvents = "none";
      } else {
        link.removeAttribute("aria-disabled");
        link.style.pointerEvents = "";
      }
    }
  } else {
    if (this.disabled) {
      __privateGet(this, _internals).ariaDisabled = "true";
      this.setAttribute("tabindex", "-1");
    } else {
      __privateGet(this, _internals).ariaDisabled = null;
      if (!this.href) {
        this.setAttribute("tabindex", "0");
      }
    }
  }
};
_onClick = new WeakMap();
_onKeydown = new WeakMap();
onLinkClick_fn = function(event) {
  if (this.disabled) {
    event.preventDefault();
    event.stopPropagation();
  }
};
__decorateElement(_init4, 4, "variant", _variant_dec, PfV6Button, _variant);
__decorateElement(_init4, 4, "size", _size_dec, PfV6Button, _size);
__decorateElement(_init4, 4, "block", _block_dec, PfV6Button, _block);
__decorateElement(_init4, 4, "disabled", _disabled_dec, PfV6Button, _disabled);
__decorateElement(_init4, 4, "noPadding", _noPadding_dec, PfV6Button, _noPadding);
__decorateElement(_init4, 4, "href", _href_dec, PfV6Button, _href);
__decorateElement(_init4, 4, "type", _type_dec, PfV6Button, _type);
PfV6Button = __decorateElement(_init4, 0, "PfV6Button", _PfV6Button_decorators, PfV6Button);
__publicField(PfV6Button, "formAssociated", true);
__publicField(PfV6Button, "shadowRootOptions", {
  ...i3.shadowRootOptions,
  delegatesFocus: true
});
__publicField(PfV6Button, "styles", cem_pf_v6_button_default);
__runInitializers(_init4, 1, PfV6Button);

// elements/cem-drawer/cem-drawer.ts
var CemDrawerChangeEvent = class extends Event {
  open;
  constructor(open) {
    super("change", { bubbles: true });
    this.open = open;
  }
};
var CemDrawerResizeEvent = class extends Event {
  height;
  constructor(height) {
    super("resize", { bubbles: true });
    this.height = height;
  }
};
var _drawerHeight_dec, _open_dec, _a5, _CemServeDrawer_decorators, _init5, _open, _drawerHeight, _isDragging, _startY, _startHeight, _maxHeight, _rafId, _pendingHeight, _CemServeDrawer_instances, $_fn, getMaxHeight_fn, _onToggleClick, _startResize, _handleResize, _applyResize, _handleKeydown, updateAriaValueNow_fn, _stopResize, _handleWindowResize;
_CemServeDrawer_decorators = [t3("cem-drawer")];
var CemServeDrawer = class extends (_a5 = i3, _open_dec = [n4({ type: Boolean, reflect: true })], _drawerHeight_dec = [n4({ type: Number, reflect: true, attribute: "drawer-height" })], _a5) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeDrawer_instances);
    __privateAdd(this, _open, __runInitializers(_init5, 8, this, false)), __runInitializers(_init5, 11, this);
    __privateAdd(this, _drawerHeight, __runInitializers(_init5, 12, this, 400)), __runInitializers(_init5, 15, this);
    __privateAdd(this, _isDragging, false);
    __privateAdd(this, _startY, 0);
    __privateAdd(this, _startHeight, 0);
    __privateAdd(this, _maxHeight, null);
    __privateAdd(this, _rafId, null);
    __privateAdd(this, _pendingHeight, null);
    __privateAdd(this, _onToggleClick, () => {
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (content) {
        content.classList.add("transitions-enabled");
      }
      this.toggle();
    });
    __privateAdd(this, _startResize, (e6) => {
      __privateSet(this, _isDragging, true);
      __privateSet(this, _startY, e6.clientY);
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (!content) return;
      __privateSet(this, _startHeight, content.offsetHeight);
      __privateSet(this, _maxHeight, __privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this));
      content.classList.remove("transitions-enabled");
      content.classList.add("resizing");
      document.addEventListener("mousemove", __privateGet(this, _handleResize), { passive: true });
      document.addEventListener("mouseup", __privateGet(this, _stopResize));
      e6.preventDefault();
    });
    __privateAdd(this, _handleResize, (e6) => {
      if (!__privateGet(this, _isDragging)) return;
      const deltaY = __privateGet(this, _startY) - e6.clientY;
      const newHeight = Math.max(100, Math.min(__privateGet(this, _maxHeight), __privateGet(this, _startHeight) + deltaY));
      __privateSet(this, _pendingHeight, newHeight);
      if (!__privateGet(this, _rafId)) {
        __privateSet(this, _rafId, requestAnimationFrame(__privateGet(this, _applyResize)));
      }
    });
    __privateAdd(this, _applyResize, () => {
      if (__privateGet(this, _pendingHeight) === null) return;
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (content) {
        content.style.height = `${__privateGet(this, _pendingHeight)}px`;
      }
      __privateSet(this, _pendingHeight, null);
      __privateSet(this, _rafId, null);
    });
    __privateAdd(this, _handleKeydown, (e6) => {
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (!content) return;
      const step = e6.shiftKey ? 50 : 10;
      let currentHeight = parseInt(content.style.height, 10) || 400;
      let newHeight = currentHeight;
      switch (e6.key) {
        case "ArrowUp":
          e6.preventDefault();
          newHeight = currentHeight + step;
          break;
        case "ArrowDown":
          e6.preventDefault();
          newHeight = currentHeight - step;
          break;
        case "Home":
          e6.preventDefault();
          newHeight = 100;
          break;
        case "End":
          e6.preventDefault();
          newHeight = __privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this);
          break;
        default:
          return;
      }
      newHeight = Math.max(100, Math.min(__privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this), newHeight));
      content.style.height = `${newHeight}px`;
      __privateMethod(this, _CemServeDrawer_instances, updateAriaValueNow_fn).call(this, newHeight);
      this.dispatchEvent(new CemDrawerResizeEvent(newHeight));
    });
    __privateAdd(this, _stopResize, () => {
      __privateSet(this, _isDragging, false);
      if (__privateGet(this, _rafId)) {
        cancelAnimationFrame(__privateGet(this, _rafId));
        __privateSet(this, _rafId, null);
      }
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (!content) return;
      content.classList.add("transitions-enabled");
      content.classList.remove("resizing");
      const height = parseInt(content.style.height, 10);
      __privateMethod(this, _CemServeDrawer_instances, updateAriaValueNow_fn).call(this, height);
      this.dispatchEvent(new CemDrawerResizeEvent(height));
      document.removeEventListener("mousemove", __privateGet(this, _handleResize));
      document.removeEventListener("mouseup", __privateGet(this, _stopResize));
    });
    __privateAdd(this, _handleWindowResize, () => {
      if (!this.open) return;
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (!content) return;
      const currentHeight = parseInt(content.style.height) || 0;
      const maxHeight = __privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this);
      if (currentHeight > maxHeight) {
        content.style.height = `${maxHeight}px`;
        this.drawerHeight = maxHeight;
        this.dispatchEvent(new CemDrawerResizeEvent(maxHeight));
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    window.addEventListener("resize", __privateGet(this, _handleWindowResize));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener("resize", __privateGet(this, _handleWindowResize));
  }
  render() {
    return T`
      <div id="resize-handle"
           role="separator"
           aria-orientation="horizontal"
           aria-label="Resize drawer"
           tabindex="0"
           aria-controls="content"
           aria-valuemin="100"
           aria-valuemax="1000"
           aria-valuenow="400"
           @mousedown=${__privateGet(this, _startResize)}
           @keydown=${__privateGet(this, _handleKeydown)}></div>
      <cem-pf-v6-button id="toggle"
                     variant="plain"
                     aria-label="Toggle drawer"
                     aria-expanded="${this.open}"
                     aria-controls="content"
                     @click=${__privateGet(this, _onToggleClick)}>
        <svg width="16"
             height="16"
             viewBox="0 0 16 16"
             fill="currentColor"
             role="presentation">
          <path fill-rule="evenodd"
                d="M1.646 4.646a.5.5 0 0 1 .708 0L8 10.293l5.646-5.647a.5.5 0 0 1 .708.708l-6 6a.5.5 0 0 1-.708 0l-6-6a.5.5 0 0 1 0-.708z"/>
        </svg>
      </cem-pf-v6-button>
      <div id="content"
           style="height: ${this.open ? this.drawerHeight : 0}px;">
        <slot></slot>
      </div>
    `;
  }
  updated(changed) {
    if (changed.has("open")) {
      const content = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "content");
      if (content && this.open) {
        let height = this.drawerHeight;
        let needsPersist = false;
        if (height <= 0 || isNaN(height)) {
          height = 400;
          needsPersist = true;
        }
        const maxHeight = __privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this);
        if (height > maxHeight) {
          height = maxHeight;
          needsPersist = true;
        }
        if (needsPersist) {
          this.drawerHeight = height;
          this.dispatchEvent(new CemDrawerResizeEvent(height));
        }
      }
      this.dispatchEvent(new CemDrawerChangeEvent(this.open));
    }
  }
  toggle() {
    this.open = !this.open;
  }
  openDrawer() {
    this.open = true;
  }
  close() {
    this.open = false;
  }
};
_init5 = __decoratorStart(_a5);
_open = new WeakMap();
_drawerHeight = new WeakMap();
_isDragging = new WeakMap();
_startY = new WeakMap();
_startHeight = new WeakMap();
_maxHeight = new WeakMap();
_rafId = new WeakMap();
_pendingHeight = new WeakMap();
_CemServeDrawer_instances = new WeakSet();
$_fn = function(id) {
  return this.shadowRoot?.getElementById(id);
};
/**
 * Returns the maximum safe height for the drawer content in pixels.
 *
 * The toggle button must always remain visible below the masthead so the
 * user can close or resize the drawer at any time.
 */
getMaxHeight_fn = function() {
  const toggle = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "toggle");
  const handle = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "resize-handle");
  const mastheadHeight = 56;
  const toggleHeight = toggle?.offsetHeight ?? 32;
  const handleHeight = handle?.offsetHeight ?? 4;
  return Math.max(100, window.innerHeight - mastheadHeight - toggleHeight - handleHeight);
};
_onToggleClick = new WeakMap();
_startResize = new WeakMap();
_handleResize = new WeakMap();
_applyResize = new WeakMap();
_handleKeydown = new WeakMap();
updateAriaValueNow_fn = function(height) {
  const resizeHandle = __privateMethod(this, _CemServeDrawer_instances, $_fn).call(this, "resize-handle");
  if (resizeHandle) {
    resizeHandle.setAttribute("aria-valuenow", String(Math.round(height)));
    resizeHandle.setAttribute("aria-valuemax", String(Math.round(__privateMethod(this, _CemServeDrawer_instances, getMaxHeight_fn).call(this))));
  }
};
_stopResize = new WeakMap();
_handleWindowResize = new WeakMap();
__decorateElement(_init5, 4, "open", _open_dec, CemServeDrawer, _open);
__decorateElement(_init5, 4, "drawerHeight", _drawerHeight_dec, CemServeDrawer, _drawerHeight);
CemServeDrawer = __decorateElement(_init5, 0, "CemServeDrawer", _CemServeDrawer_decorators, CemServeDrawer);
__publicField(CemServeDrawer, "styles", cem_drawer_default);
__runInitializers(_init5, 1, CemServeDrawer);

// lit-css:elements/cem-health-panel/cem-health-panel.css
var s9 = new CSSStyleSheet();
s9.replaceSync(JSON.parse('":host {\\n  display: block;\\n  padding: var(--pf-t--global--spacer--md);\\n}\\n\\n#overall {\\n  margin-block-end: var(--pf-t--global--spacer--md);\\n}\\n\\n.category-bar {\\n  display: flex;\\n  align-items: center;\\n  gap: var(--pf-t--global--spacer--sm);\\n}\\n\\n.category-meter {\\n  flex: 1;\\n  height: 8px;\\n  border-radius: var(--pf-t--global--border--radius--pill, 999px);\\n  background: var(--pf-t--global--background--color--secondary--default, #f0f0f0);\\n  overflow: hidden;\\n\\n  \\u0026 \\u003e .category-fill {\\n    height: 100%;\\n    border-radius: inherit;\\n    transition: width 0.3s ease;\\n  }\\n}\\n\\n.fill-pass {\\n  background: var(--pf-t--global--color--status--success--default, #3e8635);\\n}\\n\\n.fill-warn {\\n  background: var(--pf-t--global--color--status--warning--default, #f0ab00);\\n}\\n\\n.fill-fail {\\n  background: var(--pf-t--global--color--status--danger--default, #c9190b);\\n}\\n\\n.category-score {\\n  min-width: 4ch;\\n  text-align: end;\\n  font-variant-numeric: tabular-nums;\\n}\\n\\n#recommendations {\\n  margin-block-start: var(--pf-t--global--spacer--md);\\n  border-block-start: 1px solid var(--pf-t--global--border--color--default, #d2d2d2);\\n  padding-block-start: var(--pf-t--global--spacer--md);\\n\\n  \\u0026 h4 {\\n    margin: 0 0 var(--pf-t--global--spacer--sm);\\n    font-size: var(--pf-t--global--font--size--body--default, 1rem);\\n    font-weight: var(--pf-t--global--font--weight--body--bold, 700);\\n  }\\n\\n  \\u0026 ul {\\n    margin: 0;\\n    padding-inline-start: var(--pf-t--global--spacer--lg);\\n  }\\n\\n  \\u0026 li {\\n    margin-block-end: var(--pf-t--global--spacer--xs);\\n    font-size: var(--pf-t--global--font--size--body--sm, 0.875rem);\\n  }\\n}\\n\\n.finding-details {\\n  font-size: var(--pf-t--global--font--size--body--sm, 0.875rem);\\n  padding-block-start: var(--pf-t--global--spacer--xs);\\n\\n  \\u0026 ul {\\n    margin: 0;\\n    padding-inline-start: var(--pf-t--global--spacer--lg);\\n    list-style: none;\\n  }\\n\\n  \\u0026 li::before {\\n    content: \\"\\\\2192 \\";\\n  }\\n}\\n"'));
var cem_health_panel_default = s9;

// elements/node_modules/lit-html/node/directives/if-defined.js
var o7 = (o9) => o9 ?? A;

// lit-css:elements/cem-pf-v6-label/cem-pf-v6-label.css
var s10 = new CSSStyleSheet();
s10.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-label--PaddingBlockStart: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-label--PaddingInlineEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-label--PaddingBlockEnd: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-label--PaddingInlineStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-label--MaxWidth: 100%;\\n  --cem-pf-v6-c-label--MinWidth: calc((var(--cem-pf-v6-c-label--FontSize) * var(--pf-t--global--font--line-height--body) + var(--cem-pf-v6-c-label--PaddingBlockStart) + var(--cem-pf-v6-c-label--PaddingBlockEnd)));\\n  --cem-pf-v6-c-label--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-label--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-label--BorderRadius: var(--pf-t--global--border--radius--pill);\\n  --cem-pf-v6-c-label--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --cem-pf-v6-c-label--Gap: var(--pf-t--global--spacer--gap--text-to-element--compact);\\n  --cem-pf-v6-c-label--BackgroundColor: var(--pf-t--global--color--nonstatus--gray--default);\\n  --cem-pf-v6-c-label--Color: var(--pf-t--global--text--color--nonstatus--on-gray--default);\\n  --cem-pf-v6-c-label__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-gray--default);\\n  --cem-pf-v6-c-label--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--gray--hover);\\n  --cem-pf-v6-c-label--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-gray--hover);\\n  --cem-pf-v6-c-label--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-gray--hover);\\n  --cem-pf-v6-c-label--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--gray--default);\\n  --cem-pf-v6-c-label--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--gray--hover);\\n\\n  /* Color variants */\\n  --cem-pf-v6-c-label--m-blue--BackgroundColor: var(--pf-t--global--color--nonstatus--blue--default);\\n  --cem-pf-v6-c-label--m-blue--Color: var(--pf-t--global--text--color--nonstatus--on-blue--default);\\n  --cem-pf-v6-c-label--m-blue__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-blue--default);\\n  --cem-pf-v6-c-label--m-blue--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--blue--hover);\\n  --cem-pf-v6-c-label--m-blue--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-blue--hover);\\n  --cem-pf-v6-c-label--m-blue--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-blue--hover);\\n  --cem-pf-v6-c-label--m-blue--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--blue--default);\\n  --cem-pf-v6-c-label--m-blue--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--blue--hover);\\n\\n  --cem-pf-v6-c-label--m-red--BackgroundColor: var(--pf-t--global--color--nonstatus--red--default);\\n  --cem-pf-v6-c-label--m-red--Color: var(--pf-t--global--text--color--nonstatus--on-red--default);\\n  --cem-pf-v6-c-label--m-red__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-red--default);\\n  --cem-pf-v6-c-label--m-red--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--red--hover);\\n  --cem-pf-v6-c-label--m-red--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-red--hover);\\n  --cem-pf-v6-c-label--m-red--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-red--hover);\\n  --cem-pf-v6-c-label--m-red--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--red--default);\\n  --cem-pf-v6-c-label--m-red--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--red--hover);\\n\\n  --cem-pf-v6-c-label--m-orange--BackgroundColor: var(--pf-t--global--color--nonstatus--orange--default);\\n  --cem-pf-v6-c-label--m-orange--Color: var(--pf-t--global--text--color--nonstatus--on-orange--default);\\n  --cem-pf-v6-c-label--m-orange__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-orange--default);\\n  --cem-pf-v6-c-label--m-orange--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--orange--hover);\\n  --cem-pf-v6-c-label--m-orange--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-orange--hover);\\n  --cem-pf-v6-c-label--m-orange--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-orange--hover);\\n  --cem-pf-v6-c-label--m-orange--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--orange--default);\\n  --cem-pf-v6-c-label--m-orange--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--orange--hover);\\n\\n  --cem-pf-v6-c-label--m-orangered--BackgroundColor: var(--pf-t--global--color--nonstatus--orangered--default);\\n  --cem-pf-v6-c-label--m-orangered--Color: var(--pf-t--global--text--color--nonstatus--on-orangered--default);\\n  --cem-pf-v6-c-label--m-orangered__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-orangered--default);\\n  --cem-pf-v6-c-label--m-orangered--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--orangered--hover);\\n  --cem-pf-v6-c-label--m-orangered--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-orangered--hover);\\n  --cem-pf-v6-c-label--m-orangered--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-orangered--hover);\\n  --cem-pf-v6-c-label--m-orangered--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--orangered--default);\\n  --cem-pf-v6-c-label--m-orangered--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--orangered--hover);\\n\\n  --cem-pf-v6-c-label--m-yellow--BackgroundColor: var(--pf-t--global--color--nonstatus--yellow--default);\\n  --cem-pf-v6-c-label--m-yellow--Color: var(--pf-t--global--text--color--nonstatus--on-yellow--default);\\n  --cem-pf-v6-c-label--m-yellow__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-yellow--default);\\n  --cem-pf-v6-c-label--m-yellow--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--yellow--hover);\\n  --cem-pf-v6-c-label--m-yellow--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-yellow--hover);\\n  --cem-pf-v6-c-label--m-yellow--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-yellow--hover);\\n  --cem-pf-v6-c-label--m-yellow--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--yellow--default);\\n  --cem-pf-v6-c-label--m-yellow--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--yellow--hover);\\n\\n  --cem-pf-v6-c-label--m-green--BackgroundColor: var(--pf-t--global--color--nonstatus--green--default);\\n  --cem-pf-v6-c-label--m-green--Color: var(--pf-t--global--text--color--nonstatus--on-green--default);\\n  --cem-pf-v6-c-label--m-green__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-green--default);\\n  --cem-pf-v6-c-label--m-green--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--green--hover);\\n  --cem-pf-v6-c-label--m-green--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-green--hover);\\n  --cem-pf-v6-c-label--m-green--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-green--hover);\\n  --cem-pf-v6-c-label--m-green--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--green--default);\\n  --cem-pf-v6-c-label--m-green--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--green--hover);\\n\\n  --cem-pf-v6-c-label--m-teal--BackgroundColor: var(--pf-t--global--color--nonstatus--teal--default);\\n  --cem-pf-v6-c-label--m-teal--Color: var(--pf-t--global--text--color--nonstatus--on-teal--default);\\n  --cem-pf-v6-c-label--m-teal__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-teal--default);\\n  --cem-pf-v6-c-label--m-teal--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--teal--hover);\\n  --cem-pf-v6-c-label--m-teal--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-teal--hover);\\n  --cem-pf-v6-c-label--m-teal--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-teal--hover);\\n  --cem-pf-v6-c-label--m-teal--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--teal--default);\\n  --cem-pf-v6-c-label--m-teal--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--teal--hover);\\n\\n  --cem-pf-v6-c-label--m-purple--BackgroundColor: var(--pf-t--global--color--nonstatus--purple--default);\\n  --cem-pf-v6-c-label--m-purple--Color: var(--pf-t--global--text--color--nonstatus--on-purple--default);\\n  --cem-pf-v6-c-label--m-purple__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-purple--default);\\n  --cem-pf-v6-c-label--m-purple--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--nonstatus--purple--hover);\\n  --cem-pf-v6-c-label--m-purple--m-clickable--hover--Color: var(--pf-t--global--text--color--nonstatus--on-purple--hover);\\n  --cem-pf-v6-c-label--m-purple--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--nonstatus--on-purple--hover);\\n  --cem-pf-v6-c-label--m-purple--m-outline--BorderColor: var(--pf-t--global--border--color--nonstatus--purple--default);\\n  --cem-pf-v6-c-label--m-purple--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--nonstatus--purple--hover);\\n\\n  /* Status variants */\\n  --cem-pf-v6-c-label--m-success--BackgroundColor: var(--pf-t--global--color--status--success--default);\\n  --cem-pf-v6-c-label--m-success--Color: var(--pf-t--global--text--color--status--on-success--default);\\n  --cem-pf-v6-c-label--m-success__icon--Color: var(--pf-t--global--icon--color--status--on-success--default);\\n  --cem-pf-v6-c-label--m-success--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--status--success--hover);\\n  --cem-pf-v6-c-label--m-success--m-clickable--hover--Color: var(--pf-t--global--text--color--status--on-success--hover);\\n  --cem-pf-v6-c-label--m-success--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--on-success--hover);\\n  --cem-pf-v6-c-label--m-success--m-outline__icon--Color: var(--pf-t--global--icon--color--status--success--default);\\n  --cem-pf-v6-c-label--m-success--m-outline--BorderColor: var(--pf-t--global--border--color--status--success--default);\\n  --cem-pf-v6-c-label--m-success--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--status--success--hover);\\n  --cem-pf-v6-c-label--m-success--m-outline--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--success--hover);\\n\\n  --cem-pf-v6-c-label--m-warning--BackgroundColor: var(--pf-t--global--color--status--warning--default);\\n  --cem-pf-v6-c-label--m-warning--Color: var(--pf-t--global--text--color--status--on-warning--default);\\n  --cem-pf-v6-c-label--m-warning__icon--Color: var(--pf-t--global--icon--color--status--on-warning--default);\\n  --cem-pf-v6-c-label--m-warning--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--status--warning--hover);\\n  --cem-pf-v6-c-label--m-warning--m-clickable--hover--Color: var(--pf-t--global--text--color--status--on-warning--hover);\\n  --cem-pf-v6-c-label--m-warning--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--on-warning--hover);\\n  --cem-pf-v6-c-label--m-warning--m-outline__icon--Color: var(--pf-t--global--icon--color--status--warning--default);\\n  --cem-pf-v6-c-label--m-warning--m-outline--BorderColor: var(--pf-t--global--border--color--status--warning--default);\\n  --cem-pf-v6-c-label--m-warning--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--status--warning--hover);\\n  --cem-pf-v6-c-label--m-warning--m-outline--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--warning--hover);\\n\\n  --cem-pf-v6-c-label--m-danger--BackgroundColor: var(--pf-t--global--color--status--danger--default);\\n  --cem-pf-v6-c-label--m-danger--Color: var(--pf-t--global--text--color--status--on-danger--default);\\n  --cem-pf-v6-c-label--m-danger__icon--Color: var(--pf-t--global--icon--color--status--on-danger--default);\\n  --cem-pf-v6-c-label--m-danger--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--status--danger--hover);\\n  --cem-pf-v6-c-label--m-danger--m-clickable--hover--Color: var(--pf-t--global--text--color--status--on-danger--hover);\\n  --cem-pf-v6-c-label--m-danger--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--on-danger--hover);\\n  --cem-pf-v6-c-label--m-danger--m-outline__icon--Color: var(--pf-t--global--icon--color--status--danger--default);\\n  --cem-pf-v6-c-label--m-danger--m-outline--BorderColor: var(--pf-t--global--border--color--status--danger--default);\\n  --cem-pf-v6-c-label--m-danger--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--status--danger--hover);\\n  --cem-pf-v6-c-label--m-danger--m-outline--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--danger--hover);\\n\\n  --cem-pf-v6-c-label--m-info--BackgroundColor: var(--pf-t--global--color--status--info--default);\\n  --cem-pf-v6-c-label--m-info--Color: var(--pf-t--global--text--color--status--on-info--default);\\n  --cem-pf-v6-c-label--m-info__icon--Color: var(--pf-t--global--icon--color--status--on-info--default);\\n  --cem-pf-v6-c-label--m-info--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--status--info--hover);\\n  --cem-pf-v6-c-label--m-info--m-clickable--hover--Color: var(--pf-t--global--text--color--status--on-info--hover);\\n  --cem-pf-v6-c-label--m-info--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--on-info--hover);\\n  --cem-pf-v6-c-label--m-info--m-outline__icon--Color: var(--pf-t--global--icon--color--status--info--default);\\n  --cem-pf-v6-c-label--m-info--m-outline--BorderColor: var(--pf-t--global--border--color--status--info--default);\\n  --cem-pf-v6-c-label--m-info--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--status--info--hover);\\n  --cem-pf-v6-c-label--m-info--m-outline--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--info--hover);\\n\\n  --cem-pf-v6-c-label--m-custom--BackgroundColor: var(--pf-t--global--color--status--custom--default);\\n  --cem-pf-v6-c-label--m-custom--Color: var(--pf-t--global--text--color--status--on-custom--default);\\n  --cem-pf-v6-c-label--m-custom__icon--Color: var(--pf-t--global--icon--color--status--on-custom--default);\\n  --cem-pf-v6-c-label--m-custom--m-clickable--hover--BackgroundColor: var(--pf-t--global--color--status--custom--hover);\\n  --cem-pf-v6-c-label--m-custom--m-clickable--hover--Color: var(--pf-t--global--text--color--status--on-custom--hover);\\n  --cem-pf-v6-c-label--m-custom--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--on-custom--hover);\\n  --cem-pf-v6-c-label--m-custom--m-outline__icon--Color: var(--pf-t--global--icon--color--status--custom--default);\\n  --cem-pf-v6-c-label--m-custom--m-outline--BorderColor: var(--pf-t--global--border--color--status--custom--default);\\n  --cem-pf-v6-c-label--m-custom--m-outline--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--status--custom--hover);\\n  --cem-pf-v6-c-label--m-custom--m-outline--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--status--custom--hover);\\n\\n  /* Modifier variants */\\n  --cem-pf-v6-c-label--m-clickable--hover--BorderWidth: var(--pf-t--global--border--width--action--hover);\\n  --cem-pf-v6-c-label--m-clickable--hover--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-label--m-clickable__content--Cursor: pointer;\\n  --cem-pf-v6-c-label--m-filled__actions--c-button__icon--Color: var(--cem-pf-v6-c-label__icon--Color);\\n  --cem-pf-v6-c-label--m-outline--BorderWidth: var(--pf-t--global--border--width--action--default);\\n  --cem-pf-v6-c-label--m-outline--BackgroundColor: transparent;\\n  --cem-pf-v6-c-label--m-outline--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-label--m-outline__icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-label--m-outline--m-clickable--hover--BackgroundColor: transparent;\\n  --cem-pf-v6-c-label--m-outline--m-clickable--hover--BorderWidth: var(--pf-t--global--border--width--action--hover);\\n  --cem-pf-v6-c-label--m-outline--m-clickable--hover--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-label--m-outline--m-clickable--hover__icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-label--m-overflow--Color: var(--pf-t--global--text--color--brand--default);\\n  --cem-pf-v6-c-label--m-overflow--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-label--m-overflow--BorderWidth: var(--pf-t--global--border--width--action--plain--default);\\n  --cem-pf-v6-c-label--m-overflow--hover--Color: var(--pf-t--global--text--color--brand--hover);\\n  --cem-pf-v6-c-label--m-overflow--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --cem-pf-v6-c-label--m-overflow--hover--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --cem-pf-v6-c-label--m-add--Color: var(--pf-t--global--text--color--brand--default);\\n  --cem-pf-v6-c-label--m-add--BackgroundColor: transparent;\\n  --cem-pf-v6-c-label--m-add--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-label--m-add--BorderWidth: var(--pf-t--global--border--width--action--default);\\n  --cem-pf-v6-c-label--m-add--hover--Color: var(--pf-t--global--text--color--brand--hover);\\n  --cem-pf-v6-c-label--m-add--hover--BackgroundColor: transparent;\\n  --cem-pf-v6-c-label--m-add--hover--BorderColor: var(--pf-t--global--border--color--hover);\\n  --cem-pf-v6-c-label--m-add--hover--BorderWidth: var(--pf-t--global--border--width--action--hover);\\n  --cem-pf-v6-c-label--m-compact--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-label--m-compact--PaddingInlineEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-label--m-compact--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-label--m-compact--PaddingInlineStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-label--m-compact--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --cem-pf-v6-c-label--m-compact--m-editable--TextUnderlineOffset: 0.0625rem;\\n  --cem-pf-v6-c-label--m-compact--MinWidth: calc((var(--cem-pf-v6-c-label--m-compact--FontSize) * var(--pf-t--global--font--line-height--body)) + (var(--cem-pf-v6-c-label--m-compact--PaddingBlockStart) + var(--cem-pf-v6-c-label--m-compact--PaddingBlockEnd)));\\n  --cem-pf-v6-c-label__content--MaxWidth: 100%;\\n  --cem-pf-v6-c-label__content--Gap: var(--pf-t--global--spacer--gap--text-to-element--compact);\\n  --cem-pf-v6-c-label__content--Cursor: initial;\\n  --cem-pf-v6-c-label__icon--FontSize: var(--pf-t--global--icon--size--font--body--sm);\\n  --cem-pf-v6-c-label__text--MaxWidth: 100%;\\n  --cem-pf-v6-c-label__actions--MarginInlineEnd: 0;\\n  --cem-pf-v6-c-label__actions--c-button--FontSize: var(--pf-t--global--icon--size--font--body--sm);\\n  --cem-pf-v6-c-label__actions--c-button--OutlineOffset: -0.1875rem;\\n  --cem-pf-v6-c-label__actions--c-button--MarginBlockStart: 0;\\n  --cem-pf-v6-c-label__actions--c-button--MarginBlockEnd: 0;\\n  --cem-pf-v6-c-label__actions--c-button--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-label__actions--c-button--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-label__actions--c-button--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-label__actions--c-button--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-label--m-editable--TextDecorationLine: var(--pf-t--global--text-decoration--editable-text--line--default);\\n  --cem-pf-v6-c-label--m-editable--TextDecorationStyle: var(--pf-t--global--text-decoration--editable-text--style--default);\\n  --cem-pf-v6-c-label--m-editable--hover--TextDecorationLine: var(--pf-t--global--text-decoration--editable-text--line--hover);\\n  --cem-pf-v6-c-label--m-editable--hover--TextDecorationStyle: var(--pf-t--global--text-decoration--editable-text--style--hover);\\n  --cem-pf-v6-c-label--m-editable--TextUnderlineOffset: 0.25rem;\\n  --cem-pf-v6-c-label--m-editable__content--MaxWidth: 16ch;\\n  --cem-pf-v6-c-label--m-editable__content--Cursor: pointer;\\n  --cem-pf-v6-c-label--m-editable--m-editable-active--TextDecorationLine: none;\\n  --cem-pf-v6-c-label--m-editable--m-editable-active--BackgroundColor: transparent;\\n  --cem-pf-v6-c-label--m-editable--m-editable-active--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-label--m-editable--m-editable-active__content--Cursor: initial;\\n  --cem-pf-v6-c-label--m-disabled--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n  --cem-pf-v6-c-label--m-disabled--Color: var(--pf-t--global--text--color--on-disabled);\\n  --cem-pf-v6-c-label--m-disabled__icon--Color: var(--pf-t--global--icon--color--on-disabled);\\n  --cem-pf-v6-c-label--m-disabled--c-button--disabled__icon--Color: var(--pf-t--global--icon--color--on-disabled);\\n\\n  position: relative;\\n  gap: var(--cem-pf-v6-c-label--Gap);\\n  min-width: var(--cem-pf-v6-c-label--MinWidth);\\n  max-width: var(--cem-pf-v6-c-label--MaxWidth);\\n  padding-block-start: var(--cem-pf-v6-c-label--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-label--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-label--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-label--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-label--FontSize);\\n  white-space: nowrap;\\n  background: transparent;\\n  isolation: isolate;\\n  border: 0;\\n  border-radius: var(--cem-pf-v6-c-label--BorderRadius);\\n  display: inline-flex;\\n  align-items: center;\\n  justify-content: center;\\n}\\n\\n/* Nonstatus backgrounds are pastel in both schemes. The text color tokens\\n   all chain down to light-dark(), which resolves to white in dark mode,\\n   producing poor contrast on pastel backgrounds. Bypass the token chain\\n   and use the light-mode base palette color directly. */\\n:host([color]) {\\n  --cem-pf-v6-c-label--Color: var(--pf-t--color--gray--95);\\n  --cem-pf-v6-c-label__icon--Color: var(--pf-t--color--gray--95);\\n}\\n\\n/* Color variants - only set background; text/icon color handled by :host([color]) */\\n:host([color=\\"blue\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-blue--BackgroundColor);\\n}\\n\\n:host([color=\\"blue\\"][variant=\\"outline\\"]) {\\n  --cem-pf-v6-c-label--m-outline--BorderColor: var(--cem-pf-v6-c-label--m-blue--m-outline--BorderColor);\\n}\\n\\n:host([color=\\"red\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-red--BackgroundColor);\\n}\\n\\n:host([color=\\"red\\"][variant=\\"outline\\"]) {\\n  --cem-pf-v6-c-label--m-outline--BorderColor: var(--cem-pf-v6-c-label--m-red--m-outline--BorderColor);\\n}\\n\\n:host([color=\\"orange\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-orange--BackgroundColor);\\n}\\n\\n:host([color=\\"orangered\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-orangered--BackgroundColor);\\n}\\n\\n:host([color=\\"yellow\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-yellow--BackgroundColor);\\n}\\n\\n:host([color=\\"green\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-green--BackgroundColor);\\n}\\n\\n:host([color=\\"teal\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-teal--BackgroundColor);\\n}\\n\\n:host([color=\\"purple\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-purple--BackgroundColor);\\n}\\n\\n/* Status variants */\\n:host([status=\\"success\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-success--BackgroundColor);\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-success--Color);\\n  --cem-pf-v6-c-label__icon--Color: var(--cem-pf-v6-c-label--m-success__icon--Color);\\n}\\n\\n:host([status=\\"warning\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-warning--BackgroundColor);\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-warning--Color);\\n  --cem-pf-v6-c-label__icon--Color: var(--cem-pf-v6-c-label--m-warning__icon--Color);\\n}\\n\\n:host([status=\\"danger\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-danger--BackgroundColor);\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-danger--Color);\\n  --cem-pf-v6-c-label__icon--Color: var(--cem-pf-v6-c-label--m-danger__icon--Color);\\n}\\n\\n:host([status=\\"info\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-info--BackgroundColor);\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-info--Color);\\n  --cem-pf-v6-c-label__icon--Color: var(--cem-pf-v6-c-label--m-info__icon--Color);\\n}\\n\\n:host([status=\\"custom\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-custom--BackgroundColor);\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-custom--Color);\\n  --cem-pf-v6-c-label__icon--Color: var(--cem-pf-v6-c-label--m-custom__icon--Color);\\n}\\n\\n/* Variant modifiers */\\n:host([variant=\\"outline\\"]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-outline--BackgroundColor);\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-outline--Color);\\n  --cem-pf-v6-c-label__icon--Color: var(--cem-pf-v6-c-label--m-outline__icon--Color);\\n  --cem-pf-v6-c-label--BorderWidth: var(--cem-pf-v6-c-label--m-outline--BorderWidth);\\n}\\n\\n:host([variant=\\"overflow\\"]) {\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-overflow--Color);\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-overflow--BackgroundColor);\\n  --cem-pf-v6-c-label--BorderWidth: var(--cem-pf-v6-c-label--m-overflow--BorderWidth);\\n}\\n\\n:host([variant=\\"add\\"]) {\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-add--Color);\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-add--BackgroundColor);\\n  --cem-pf-v6-c-label--BorderColor: var(--cem-pf-v6-c-label--m-add--BorderColor);\\n  --cem-pf-v6-c-label--BorderWidth: var(--cem-pf-v6-c-label--m-add--BorderWidth);\\n}\\n\\n:host([compact]) {\\n  --cem-pf-v6-c-label--PaddingBlockStart: var(--cem-pf-v6-c-label--m-compact--PaddingBlockStart);\\n  --cem-pf-v6-c-label--PaddingBlockEnd: var(--cem-pf-v6-c-label--m-compact--PaddingBlockEnd);\\n  --cem-pf-v6-c-label--PaddingInlineStart: var(--cem-pf-v6-c-label--m-compact--PaddingInlineStart);\\n  --cem-pf-v6-c-label--PaddingInlineEnd: var(--cem-pf-v6-c-label--m-compact--PaddingInlineEnd);\\n  --cem-pf-v6-c-label--FontSize: var(--cem-pf-v6-c-label--m-compact--FontSize);\\n  --cem-pf-v6-c-label--MinWidth: var(--cem-pf-v6-c-label--m-compact--MinWidth);\\n}\\n\\n:host([disabled]) {\\n  --cem-pf-v6-c-label--BackgroundColor: var(--cem-pf-v6-c-label--m-disabled--BackgroundColor);\\n  --cem-pf-v6-c-label--Color: var(--cem-pf-v6-c-label--m-disabled--Color);\\n  --cem-pf-v6-c-label__icon--Color: var(--cem-pf-v6-c-label--m-disabled__icon--Color);\\n}\\n\\n#content {\\n  display: inline-flex;\\n  align-items: center;\\n  justify-content: center;\\n  gap: var(--cem-pf-v6-c-label__content--Gap);\\n  max-width: var(--cem-pf-v6-c-label__content--MaxWidth);\\n  padding: 0;\\n  overflow: hidden;\\n  color: var(--cem-pf-v6-c-label--Color);\\n  cursor: var(--cem-pf-v6-c-label__content--Cursor, revert);\\n  background: transparent;\\n  border-width: 0;\\n  text-decoration: none;\\n\\n  \\u0026::before {\\n    position: absolute;\\n    inset: 0;\\n    z-index: -1;\\n    pointer-events: none;\\n    content: \\"\\";\\n    background-color: var(--cem-pf-v6-c-label--BackgroundColor);\\n    border: var(--cem-pf-v6-c-label--BorderWidth) solid var(--cem-pf-v6-c-label--BorderColor);\\n    border-radius: var(--cem-pf-v6-c-label--BorderRadius);\\n  }\\n}\\n\\n::slotted([slot=\\"icon\\"]) {\\n  color: var(--cem-pf-v6-c-label__icon--Color);\\n  font-size: var(--cem-pf-v6-c-label__icon--FontSize);\\n}\\n\\n#text {\\n  overflow: hidden;\\n  text-overflow: ellipsis;\\n  white-space: nowrap;\\n  max-width: var(--cem-pf-v6-c-label__text--MaxWidth);\\n}\\n\\n#actions {\\n  margin-inline-end: var(--cem-pf-v6-c-label__actions--MarginInlineEnd);\\n\\n  \\u0026[hidden] {\\n    display: none;\\n  }\\n}\\n"'));
var cem_pf_v6_label_default = s10;

// elements/cem-pf-v6-label/cem-pf-v6-label.ts
var _href_dec2, _editable_dec, _disabled_dec2, _compact_dec, _status_dec, _variant_dec2, _color_dec, _a6, _PfV6Label_decorators, _init6, _color, _variant2, _status, _compact, _disabled2, _editable, _href2, _hasActions, _PfV6Label_instances, onContentClick_fn, onActionsSlotChange_fn;
_PfV6Label_decorators = [t3("cem-pf-v6-label")];
var PfV6Label = class extends (_a6 = i3, _color_dec = [n4({ reflect: true })], _variant_dec2 = [n4({ reflect: true })], _status_dec = [n4({ reflect: true })], _compact_dec = [n4({ type: Boolean, reflect: true })], _disabled_dec2 = [n4({ type: Boolean, reflect: true })], _editable_dec = [n4({ type: Boolean, reflect: true })], _href_dec2 = [n4({ reflect: true })], _a6) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Label_instances);
    __privateAdd(this, _color, __runInitializers(_init6, 8, this)), __runInitializers(_init6, 11, this);
    __privateAdd(this, _variant2, __runInitializers(_init6, 12, this)), __runInitializers(_init6, 15, this);
    __privateAdd(this, _status, __runInitializers(_init6, 16, this)), __runInitializers(_init6, 19, this);
    __privateAdd(this, _compact, __runInitializers(_init6, 20, this, false)), __runInitializers(_init6, 23, this);
    __privateAdd(this, _disabled2, __runInitializers(_init6, 24, this, false)), __runInitializers(_init6, 27, this);
    __privateAdd(this, _editable, __runInitializers(_init6, 28, this, false)), __runInitializers(_init6, 31, this);
    __privateAdd(this, _href2, __runInitializers(_init6, 32, this)), __runInitializers(_init6, 35, this);
    __privateAdd(this, _hasActions, false);
  }
  render() {
    const inner = T`
      <slot name="icon"></slot>
      <span id="text" part="text">
        <slot></slot>
      </span>
      <span id="actions"
            part="actions"
            ?hidden=${!__privateGet(this, _hasActions)}>
        <slot name="actions"
              @slotchange=${__privateMethod(this, _PfV6Label_instances, onActionsSlotChange_fn)}></slot>
      </span>
    `;
    if (this.href) {
      return T`
        <a id="content"
           part="content"
           href=${o7(this.disabled ? void 0 : this.href)}
           @click=${__privateMethod(this, _PfV6Label_instances, onContentClick_fn)}>${inner}</a>
      `;
    }
    return T`
      <span id="content"
            part="content"
            @click=${__privateMethod(this, _PfV6Label_instances, onContentClick_fn)}>${inner}</span>
    `;
  }
};
_init6 = __decoratorStart(_a6);
_color = new WeakMap();
_variant2 = new WeakMap();
_status = new WeakMap();
_compact = new WeakMap();
_disabled2 = new WeakMap();
_editable = new WeakMap();
_href2 = new WeakMap();
_hasActions = new WeakMap();
_PfV6Label_instances = new WeakSet();
onContentClick_fn = function(event) {
  if (this.disabled) {
    event.preventDefault();
    event.stopPropagation();
  }
};
onActionsSlotChange_fn = function(event) {
  const slot = event.target;
  const hadActions = __privateGet(this, _hasActions);
  __privateSet(this, _hasActions, slot.assignedElements().length > 0);
  if (hadActions !== __privateGet(this, _hasActions)) {
    this.requestUpdate();
  }
};
__decorateElement(_init6, 4, "color", _color_dec, PfV6Label, _color);
__decorateElement(_init6, 4, "variant", _variant_dec2, PfV6Label, _variant2);
__decorateElement(_init6, 4, "status", _status_dec, PfV6Label, _status);
__decorateElement(_init6, 4, "compact", _compact_dec, PfV6Label, _compact);
__decorateElement(_init6, 4, "disabled", _disabled_dec2, PfV6Label, _disabled2);
__decorateElement(_init6, 4, "editable", _editable_dec, PfV6Label, _editable);
__decorateElement(_init6, 4, "href", _href_dec2, PfV6Label, _href2);
PfV6Label = __decorateElement(_init6, 0, "PfV6Label", _PfV6Label_decorators, PfV6Label);
__publicField(PfV6Label, "styles", cem_pf_v6_label_default);
__runInitializers(_init6, 1, PfV6Label);

// lit-css:elements/cem-pf-v6-expandable-section/cem-pf-v6-expandable-section.css
var s11 = new CSSStyleSheet();
s11.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-expandable-section--Gap: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-expandable-section__toggle-icon--MinWidth: 1em;\\n  --cem-pf-v6-c-expandable-section__toggle-icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-expandable-section__toggle-icon--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-expandable-section__toggle-icon--TransitionDuration: var(--pf-t--global--motion--duration--icon--default);\\n  --cem-pf-v6-c-expandable-section__toggle-icon--Transition: transform var(--cem-pf-v6-c-expandable-section__toggle-icon--TransitionDuration) var(--cem-pf-v6-c-expandable-section__toggle-icon--TransitionTimingFunction);\\n  --cem-pf-v6-c-expandable-section__toggle-icon--Rotate: 0;\\n  --cem-pf-v6-c-expandable-section__toggle-icon--m-expand-top--Rotate: 0;\\n  --cem-pf-v6-c-expandable-section--m-expanded__toggle-icon--Rotate: 90deg;\\n  --cem-pf-v6-c-expandable-section--m-expanded__toggle-icon--m-expand-top--Rotate: -90deg;\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--collapse--slide: 0s;\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--collapse--fade: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--expand--slide: 0s;\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--expand--fade: var(--pf-t--global--motion--duration--fade--default);\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--slide: var(--cem-pf-v6-c-expandable-section__content--TransitionDuration--collapse--slide);\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--fade: var(--cem-pf-v6-c-expandable-section__content--TransitionDuration--collapse--fade);\\n  --cem-pf-v6-c-expandable-section__content--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-expandable-section__content--TransitionDelay--hide: var(--cem-pf-v6-c-expandable-section__content--TransitionDuration--fade);\\n  --cem-pf-v6-c-expandable-section__content--Opacity: 0;\\n  --cem-pf-v6-c-expandable-section__content--TranslateY: 0;\\n  --cem-pf-v6-c-expandable-section__content--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-expandable-section--m-expand-top__content--TranslateY: 0;\\n  --cem-pf-v6-c-expandable-section--m-expand-bottom__content--TranslateY: 0;\\n  --cem-pf-v6-c-expandable-section--m-expanded__content--Opacity: 1;\\n  --cem-pf-v6-c-expandable-section--m-expanded__content--TranslateY: 0;\\n  --cem-pf-v6-c-expandable-section__content--MaxWidth: auto;\\n  --cem-pf-v6-c-expandable-section--m-limit-width__content--MaxWidth: 46.875rem;\\n  --cem-pf-v6-c-expandable-section--m-display-lg--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--m-expanded--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--BackgroundColor: var(--pf-t--global--background--color--secondary--default);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--BorderWidth: var(--pf-t--global--border--width--box--default);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--BorderRadius: var(--pf-t--global--border--radius--medium);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--TransitionDelay: var(--cem-pf-v6-c-expandable-section__content--TransitionDuration--collapse--fade);\\n  --cem-pf-v6-c-expandable-section--m-indented__content--PaddingInlineStart: calc(var(--pf-t--global--spacer--action--horizontal--plain--default) + var(--pf-t--global--spacer--gap--text-to-element--default) + var(--cem-pf-v6-c-expandable-section__toggle-icon--MinWidth));\\n  --cem-pf-v6-c-expandable-section--m-truncate__content--LineClamp: 3;\\n  --cem-pf-v6-c-expandable-section--m-truncate--Gap: var(--pf-t--global--spacer--xs);\\n\\n  display: flex;\\n  flex-direction: column;\\n  gap: 0;\\n  transition-delay: var(--cem-pf-v6-c-expandable-section__content--TransitionDelay--hide);\\n  transition-duration: 0s;\\n  transition-property: gap, padding-block-end;\\n}\\n\\n@media screen and (prefers-reduced-motion: no-preference) {\\n  :host {\\n    --cem-pf-v6-c-expandable-section__content--TransitionDuration--collapse--slide: var(--pf-t--global--motion--duration--fade--short);\\n    --cem-pf-v6-c-expandable-section__content--TransitionDuration--expand--slide: var(--pf-t--global--motion--duration--fade--default);\\n    --cem-pf-v6-c-expandable-section__content--TranslateY: -.5rem;\\n    --cem-pf-v6-c-expandable-section--m-expand-top__content--TranslateY: .5rem;\\n    --cem-pf-v6-c-expandable-section--m-expand-bottom__content--TranslateY: -.5rem;\\n  }\\n}\\n\\n:host([expanded]) {\\n  --cem-pf-v6-c-expandable-section__toggle-icon--Rotate: var(--cem-pf-v6-c-expandable-section--m-expanded__toggle-icon--Rotate);\\n  --cem-pf-v6-c-expandable-section__toggle-icon--m-expand-top--Rotate: var(--cem-pf-v6-c-expandable-section--m-expanded__toggle-icon--m-expand-top--Rotate);\\n  --cem-pf-v6-c-expandable-section--m-display-lg--PaddingBlockEnd: var(--cem-pf-v6-c-expandable-section--m-display-lg--m-expanded--PaddingBlockEnd);\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--slide: var(--cem-pf-v6-c-expandable-section__content--TransitionDuration--expand--slide);\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--fade: var(--cem-pf-v6-c-expandable-section__content--TransitionDuration--expand--fade);\\n  --cem-pf-v6-c-expandable-section__content--Opacity: var(--cem-pf-v6-c-expandable-section--m-expanded__content--Opacity);\\n  --cem-pf-v6-c-expandable-section__content--TranslateY: var(--cem-pf-v6-c-expandable-section--m-expanded__content--TranslateY);\\n  --cem-pf-v6-c-expandable-section__content--Visibility: auto;\\n  --cem-pf-v6-c-expandable-section__content--Overflow: visible;\\n  --cem-pf-v6-c-expandable-section__content--MaxHeight: 99999px;\\n  --cem-pf-v6-c-expandable-section__content--TransitionDelay--hide: 0s;\\n\\n  gap: var(--cem-pf-v6-c-expandable-section--Gap);\\n}\\n\\n:host([expand-top]) {\\n  --cem-pf-v6-c-expandable-section__toggle-icon--Rotate: var(--cem-pf-v6-c-expandable-section__toggle-icon--m-expand-top--Rotate);\\n}\\n\\n:host(:has(#content:only-child):not([expand-top], [expand-bottom])) {\\n  --cem-pf-v6-c-expandable-section__content--TranslateY: 0;\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--expand--fade: 0s;\\n  --cem-pf-v6-c-expandable-section__content--TransitionDuration--collapse--fade: 0s;\\n}\\n\\n:host(:has(#content:only-child)[expand-top]:not([expanded])) {\\n  --cem-pf-v6-c-expandable-section__content--TranslateY: var(--cem-pf-v6-c-expandable-section--m-expand-top__content--TranslateY);\\n}\\n\\n:host(:has(#content:only-child)[expand-bottom]:not([expanded])) {\\n  --cem-pf-v6-c-expandable-section__content--TranslateY: var(--cem-pf-v6-c-expandable-section--m-expand-bottom__content--TranslateY);\\n}\\n\\n:host([limit-width]) {\\n  --cem-pf-v6-c-expandable-section__content--MaxWidth: var(--cem-pf-v6-c-expandable-section--m-limit-width__content--MaxWidth);\\n}\\n\\n:host([display-lg]) {\\n  padding-block-start: var(--cem-pf-v6-c-expandable-section--m-display-lg--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-expandable-section--m-display-lg--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-expandable-section--m-display-lg--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-expandable-section--m-display-lg--PaddingInlineEnd);\\n  background-color: var(--cem-pf-v6-c-expandable-section--m-display-lg--BackgroundColor);\\n  border: var(--cem-pf-v6-c-expandable-section--m-display-lg--BorderWidth) solid var(--cem-pf-v6-c-expandable-section--m-display-lg--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-expandable-section--m-display-lg--BorderRadius);\\n}\\n\\n:host([indented]) {\\n  --cem-pf-v6-c-expandable-section__content--PaddingInlineStart: var(--cem-pf-v6-c-expandable-section--m-indented__content--PaddingInlineStart);\\n}\\n\\n:host([truncate]) {\\n  --cem-pf-v6-c-expandable-section--Gap: var(--cem-pf-v6-c-expandable-section--m-truncate--Gap);\\n  --cem-pf-v6-c-expandable-section__content--TransitionDelay--hide: 0s;\\n}\\n\\n:host([truncate]:not([expanded])) #content {\\n  display: -webkit-box;\\n  -webkit-box-orient: vertical;\\n  -webkit-line-clamp: var(--cem-pf-v6-c-expandable-section--m-truncate__content--LineClamp);\\n  overflow: hidden;\\n}\\n\\n/* Details element reset */\\ndetails {\\n  display: contents;\\n}\\n\\nsummary {\\n  display: flex;\\n  align-items: center;\\n  list-style: none;\\n  cursor: pointer;\\n}\\n\\n/* Remove default marker in all browsers */\\nsummary::marker,\\nsummary::-webkit-details-marker {\\n  display: none;\\n}\\n\\n/* Toggle icon */\\n#toggle-icon {\\n  display: inline-block;\\n  min-width: var(--cem-pf-v6-c-expandable-section__toggle-icon--MinWidth);\\n  color: var(--cem-pf-v6-c-expandable-section__toggle-icon--Color);\\n  transition: var(--cem-pf-v6-c-expandable-section__toggle-icon--Transition);\\n  transform: rotate(var(--cem-pf-v6-c-expandable-section__toggle-icon--Rotate));\\n}\\n\\n:host(:dir(rtl)) #toggle-icon {\\n  scale: -1 1;\\n}\\n\\n:host([expand-top]) #toggle-icon {\\n  --cem-pf-v6-c-expandable-section__toggle-icon--Rotate: var(--cem-pf-v6-c-expandable-section__toggle-icon--m-expand-top--Rotate);\\n}\\n\\n/* Content */\\n#content {\\n  max-width: var(--cem-pf-v6-c-expandable-section__content--MaxWidth);\\n  padding-block-end: var(--cem-pf-v6-c-expandable-section__content--PaddingBlockEnd, 0);\\n  padding-inline-start: var(--cem-pf-v6-c-expandable-section__content--PaddingInlineStart);\\n}\\n\\n:host(:not([truncate])) #content {\\n  max-height: var(--cem-pf-v6-c-expandable-section__content--MaxHeight, 0);\\n  overflow: var(--cem-pf-v6-c-expandable-section__content--Overflow, hidden);\\n  visibility: var(--cem-pf-v6-c-expandable-section__content--Visibility, hidden);\\n  opacity: var(--cem-pf-v6-c-expandable-section__content--Opacity);\\n  transition-delay: 0s, 0s, var(--cem-pf-v6-c-expandable-section__content--TransitionDelay--hide, 0s), var(--cem-pf-v6-c-expandable-section__content--TransitionDelay--hide, 0s);\\n  transition-timing-function: var(--cem-pf-v6-c-expandable-section__content--TransitionTimingFunction);\\n  transition-duration: var(--cem-pf-v6-c-expandable-section__content--TransitionDuration--fade), var(--cem-pf-v6-c-expandable-section__content--TransitionDuration--slide), 0s, 0s;\\n  transition-property: opacity, translate, visibility, max-height;\\n  translate: 0 var(--cem-pf-v6-c-expandable-section__content--TranslateY);\\n}\\n"'));
var cem_pf_v6_expandable_section_default = s11;

// elements/cem-pf-v6-expandable-section/cem-pf-v6-expandable-section.ts
var PfExpandableSectionToggleEvent = class extends Event {
  expanded;
  constructor(expanded) {
    super("toggle", { bubbles: true });
    this.expanded = expanded;
  }
};
var _toggleText_dec, _truncate_dec, _limitWidth_dec, _indented_dec, _displayLg_dec, _expandBottom_dec, _expandTop_dec, _expanded_dec, _a7, _PfV6ExpandableSection_decorators, _init7, _expanded, _expandTop, _expandBottom, _displayLg, _indented, _limitWidth, _truncate, _toggleText, _PfV6ExpandableSection_instances, onDetailsToggle_fn;
_PfV6ExpandableSection_decorators = [t3("cem-pf-v6-expandable-section")];
var PfV6ExpandableSection = class extends (_a7 = i3, _expanded_dec = [n4({ type: Boolean, reflect: true })], _expandTop_dec = [n4({ type: Boolean, reflect: true, attribute: "expand-top" })], _expandBottom_dec = [n4({ type: Boolean, reflect: true, attribute: "expand-bottom" })], _displayLg_dec = [n4({ type: Boolean, reflect: true, attribute: "display-lg" })], _indented_dec = [n4({ type: Boolean, reflect: true })], _limitWidth_dec = [n4({ type: Boolean, reflect: true, attribute: "limit-width" })], _truncate_dec = [n4({ type: Boolean, reflect: true })], _toggleText_dec = [n4({ attribute: "toggle-text" })], _a7) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6ExpandableSection_instances);
    __privateAdd(this, _expanded, __runInitializers(_init7, 8, this, false)), __runInitializers(_init7, 11, this);
    __privateAdd(this, _expandTop, __runInitializers(_init7, 12, this, false)), __runInitializers(_init7, 15, this);
    __privateAdd(this, _expandBottom, __runInitializers(_init7, 16, this, false)), __runInitializers(_init7, 19, this);
    __privateAdd(this, _displayLg, __runInitializers(_init7, 20, this, false)), __runInitializers(_init7, 23, this);
    __privateAdd(this, _indented, __runInitializers(_init7, 24, this, false)), __runInitializers(_init7, 27, this);
    __privateAdd(this, _limitWidth, __runInitializers(_init7, 28, this, false)), __runInitializers(_init7, 31, this);
    __privateAdd(this, _truncate, __runInitializers(_init7, 32, this, false)), __runInitializers(_init7, 35, this);
    __privateAdd(this, _toggleText, __runInitializers(_init7, 36, this, "")), __runInitializers(_init7, 39, this);
  }
  toggle() {
    this.expanded = !this.expanded;
  }
  show() {
    this.expanded = true;
  }
  hide() {
    this.expanded = false;
  }
  render() {
    return T`
      <details ?open=${this.expanded}
               @toggle=${__privateMethod(this, _PfV6ExpandableSection_instances, onDetailsToggle_fn)}>
        <summary part="toggle">
          <cem-pf-v6-button id="toggle-button"
                        part="toggle-button"
                        variant="link"
                        tabindex="-1">
            <span id="toggle-icon"
                  slot="icon-start">
              <slot name="toggle-icon">
                <svg viewBox="0 0 256 512"
                     fill="currentColor"
                     aria-hidden="true"
                     role="img"
                     width="1em"
                     height="1em">
                  <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
                </svg>
              </slot>
            </span>
            <slot name="toggle-text">${this.toggleText || A}</slot>
          </cem-pf-v6-button>
        </summary>
        <div id="content"
             part="content">
          <slot></slot>
        </div>
      </details>
    `;
  }
};
_init7 = __decoratorStart(_a7);
_expanded = new WeakMap();
_expandTop = new WeakMap();
_expandBottom = new WeakMap();
_displayLg = new WeakMap();
_indented = new WeakMap();
_limitWidth = new WeakMap();
_truncate = new WeakMap();
_toggleText = new WeakMap();
_PfV6ExpandableSection_instances = new WeakSet();
onDetailsToggle_fn = function(event) {
  const details = event.target;
  this.expanded = details.open;
  this.dispatchEvent(new PfExpandableSectionToggleEvent(this.expanded));
};
__decorateElement(_init7, 4, "expanded", _expanded_dec, PfV6ExpandableSection, _expanded);
__decorateElement(_init7, 4, "expandTop", _expandTop_dec, PfV6ExpandableSection, _expandTop);
__decorateElement(_init7, 4, "expandBottom", _expandBottom_dec, PfV6ExpandableSection, _expandBottom);
__decorateElement(_init7, 4, "displayLg", _displayLg_dec, PfV6ExpandableSection, _displayLg);
__decorateElement(_init7, 4, "indented", _indented_dec, PfV6ExpandableSection, _indented);
__decorateElement(_init7, 4, "limitWidth", _limitWidth_dec, PfV6ExpandableSection, _limitWidth);
__decorateElement(_init7, 4, "truncate", _truncate_dec, PfV6ExpandableSection, _truncate);
__decorateElement(_init7, 4, "toggleText", _toggleText_dec, PfV6ExpandableSection, _toggleText);
PfV6ExpandableSection = __decorateElement(_init7, 0, "PfV6ExpandableSection", _PfV6ExpandableSection_decorators, PfV6ExpandableSection);
__publicField(PfV6ExpandableSection, "styles", cem_pf_v6_expandable_section_default);
__runInitializers(_init7, 1, PfV6ExpandableSection);

// elements/cem-health-panel/cem-health-panel.ts
var STATUS_COLORS = {
  pass: "green",
  warn: "orange",
  fail: "red"
};
var _result_dec, _loading_dec, _component_dec, _a8, _CemHealthPanel_decorators, _init8, _component, _loading, _b2, loading_get, loading_set, _CemHealthPanel_instances, _result, _c, result_get, result_set, _abortController, findDeclaration_fn2, renderCategory_fn, renderRecommendations_fn;
_CemHealthPanel_decorators = [t3("cem-health-panel")];
var CemHealthPanel = class extends (_a8 = i3, _component_dec = [n4({ reflect: true })], _loading_dec = [r5()], _result_dec = [r5()], _a8) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemHealthPanel_instances);
    __privateAdd(this, _component, __runInitializers(_init8, 8, this, null)), __runInitializers(_init8, 11, this);
    __privateAdd(this, _loading, __runInitializers(_init8, 12, this, true)), __runInitializers(_init8, 15, this);
    __privateAdd(this, _result, __runInitializers(_init8, 16, this, null)), __runInitializers(_init8, 19, this);
    __privateAdd(this, _abortController, null);
  }
  connectedCallback() {
    super.connectedCallback();
    this.fetchHealth();
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(this, _abortController)?.abort();
  }
  async fetchHealth() {
    __privateGet(this, _abortController)?.abort();
    __privateSet(this, _abortController, new AbortController());
    __privateSet(this, _CemHealthPanel_instances, true, loading_set);
    __privateSet(this, _CemHealthPanel_instances, null, result_set);
    const url = new URL("/__cem/api/health", location.origin);
    if (this.component) {
      url.searchParams.set("component", this.component);
    }
    try {
      const response = await fetch(url, { signal: __privateGet(this, _abortController).signal });
      if (!response.ok) throw new Error(`HTTP ${response.status}`);
      __privateSet(this, _CemHealthPanel_instances, await response.json(), result_set);
    } catch (e6) {
      if (e6.name === "AbortError") return;
      console.warn("cem-health-panel: failed to fetch health data", e6);
    } finally {
      __privateSet(this, _CemHealthPanel_instances, false, loading_set);
    }
  }
  render() {
    if (__privateGet(this, _CemHealthPanel_instances, loading_get)) {
      return T`<p>Analyzing documentation health...</p>`;
    }
    const decl = __privateMethod(this, _CemHealthPanel_instances, findDeclaration_fn2).call(this);
    if (!decl) {
      return T`<p>No health data available.</p>`;
    }
    const pct = decl.maxScore > 0 ? Math.round(decl.score / decl.maxScore * 100) : 0;
    const overallStatus = pct >= 80 ? "pass" : pct >= 40 ? "warn" : "fail";
    return T`
      <div id="overall">
        <cem-pf-v6-label color=${STATUS_COLORS[overallStatus]}
                     size="lg">${pct}% -- ${decl.score}/${decl.maxScore}</cem-pf-v6-label>
      </div>
      <dl id="categories"
          class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        ${decl.categories.map((cat) => __privateMethod(this, _CemHealthPanel_instances, renderCategory_fn).call(this, cat))}
      </dl>
      ${__privateMethod(this, _CemHealthPanel_instances, renderRecommendations_fn).call(this)}
    `;
  }
};
_init8 = __decoratorStart(_a8);
_component = new WeakMap();
_loading = new WeakMap();
_CemHealthPanel_instances = new WeakSet();
_result = new WeakMap();
_abortController = new WeakMap();
findDeclaration_fn2 = function() {
  if (!__privateGet(this, _CemHealthPanel_instances, result_get)) return null;
  for (const mod of __privateGet(this, _CemHealthPanel_instances, result_get).modules) {
    for (const decl of mod.declarations) {
      if (!this.component || decl.tagName === this.component || decl.name === this.component) {
        return decl;
      }
    }
  }
  return null;
};
renderCategory_fn = function(cat) {
  const pct = cat.maxPoints > 0 ? Math.round(cat.points / cat.maxPoints * 100) : 0;
  const findings = cat.findings?.filter((f3) => f3.message) ?? [];
  return T`
      <div class="cem-pf-v6-c-description-list__group">
        <dt class="cem-pf-v6-c-description-list__term">${cat.category}</dt>
        <dd class="cem-pf-v6-c-description-list__description">
          <div class="category-bar">
            <div class="category-meter">
              <div class="category-fill fill-${cat.status}"
                   style="width: ${pct}%"></div>
            </div>
            <span class="category-score">${cat.points}/${cat.maxPoints}</span>
            <cem-pf-v6-label color=${STATUS_COLORS[cat.status]}>${cat.status}</cem-pf-v6-label>
          </div>
          ${findings.length > 0 ? T`
            <div class="finding-details">
              <ul>
                ${findings.map((f3) => T`<li>${f3.message}</li>`)}
              </ul>
            </div>
          ` : A}
        </dd>
      </div>
    `;
};
renderRecommendations_fn = function() {
  const recs = __privateGet(this, _CemHealthPanel_instances, result_get)?.recommendations;
  if (!recs?.length) return A;
  return T`
      <div id="recommendations">
        <h4>Recommendations</h4>
        <ul>
          ${recs.map((rec) => T`<li>${rec}</li>`)}
        </ul>
      </div>
    `;
};
__decorateElement(_init8, 4, "component", _component_dec, CemHealthPanel, _component);
_b2 = __decorateElement(_init8, 20, "#loading", _loading_dec, _CemHealthPanel_instances, _loading), loading_get = _b2.get, loading_set = _b2.set;
_c = __decorateElement(_init8, 20, "#result", _result_dec, _CemHealthPanel_instances, _result), result_get = _c.get, result_set = _c.set;
CemHealthPanel = __decorateElement(_init8, 0, "CemHealthPanel", _CemHealthPanel_decorators, CemHealthPanel);
__publicField(CemHealthPanel, "styles", cem_health_panel_default);
__runInitializers(_init8, 1, CemHealthPanel);

// lit-css:elements/cem-manifest-browser/cem-manifest-browser.css
var s12 = new CSSStyleSheet();
s12.replaceSync(JSON.parse('"/* Manifest Browser - Component API navigation tree */\\n\\n:host {\\n  display: block;\\n  height: 100%;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n#drawer {\\n  height: 100%;\\n}\\n\\n#drawer-content {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n  overflow: hidden;\\n}\\n\\n#tree-wrapper {\\n  flex: 1;\\n  overflow: auto;\\n  min-height: 0;\\n}\\n\\n::slotted(cem-pf-v6-tree-view) {\\n  margin: var(--pf-t--global--spacer--md, 1rem);\\n}\\n\\n#panel-title {\\n  margin: 0;\\n}\\n\\n#details-content {\\n  padding: var(--pf-t--global--spacer--md, 1rem);\\n\\n  h3 {\\n    margin-top: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n    font-size: var(--pf-t--global--font--size--heading--lg, 1.25rem);\\n  }\\n\\n  dl {\\n    margin: 0;\\n  }\\n\\n  dt {\\n    font-weight: var(--pf-t--global--font--weight--body--bold, 600);\\n    margin-top: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  dd {\\n    margin-left: 0;\\n    margin-bottom: var(--pf-t--global--spacer--sm, 0.5rem);\\n  }\\n\\n  .empty-state {\\n    color: var(--pf-t--global--text--color--subtle, #6c757d);\\n    font-style: italic;\\n  }\\n}\\n"'));
var cem_manifest_browser_default = s12;

// lit-css:elements/cem-pf-v6-badge/cem-pf-v6-badge.css
var s13 = new CSSStyleSheet();
s13.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-badge--BorderColor: transparent;\\n  --cem-pf-v6-c-badge--BorderWidth: var(--pf-t--global--border--width--regular);\\n  --cem-pf-v6-c-badge--BorderRadius: var(--pf-t--global--border--radius--pill);\\n  --cem-pf-v6-c-badge--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --cem-pf-v6-c-badge--FontWeight: var(--pf-t--global--font--weight--body--bold);\\n  --cem-pf-v6-c-badge--PaddingInlineEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-badge--PaddingInlineStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-badge--Color: var(--pf-t--global--text--color--nonstatus--on-gray--default);\\n  --cem-pf-v6-c-badge--MinWidth: var(--pf-t--global--spacer--xl);\\n  --cem-pf-v6-c-badge--BackgroundColor: var(--pf-t--global--color--nonstatus--gray--default);\\n\\n  position: relative;\\n  display: inline-block;\\n  min-width: var(--cem-pf-v6-c-badge--MinWidth);\\n  padding-inline-start: var(--cem-pf-v6-c-badge--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-badge--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-badge--FontSize);\\n  font-weight: var(--cem-pf-v6-c-badge--FontWeight);\\n  color: var(--cem-pf-v6-c-badge--Color);\\n  text-align: center;\\n  white-space: nowrap;\\n  background-color: var(--cem-pf-v6-c-badge--BackgroundColor);\\n  border-radius: var(--cem-pf-v6-c-badge--BorderRadius);\\n}\\n\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border: var(--cem-pf-v6-c-badge--BorderWidth) solid var(--cem-pf-v6-c-badge--BorderColor);\\n  border-radius: inherit;\\n}\\n\\n:host([read]) {\\n  --cem-pf-v6-c-badge--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-badge--Color: var(--pf-t--global--text--color--nonstatus--on-gray--default);\\n  --cem-pf-v6-c-badge--BackgroundColor: var(--pf-t--global--color--nonstatus--gray--default);\\n}\\n\\n:host([unread]) {\\n  --cem-pf-v6-c-badge--Color: var(--pf-t--global--text--color--on-brand--default);\\n  --cem-pf-v6-c-badge--BackgroundColor: var(--pf-t--global--color--brand--default);\\n}\\n\\n:host([disabled]) {\\n  --cem-pf-v6-c-badge--Color: var(--pf-t--global--text--color--on-disabled);\\n  --cem-pf-v6-c-badge--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n}\\n\\n:host([disabled])::after {\\n  border-color: var(--pf-t--global--border--color--disabled);\\n}\\n\\n#sr-text {\\n  position: absolute;\\n  width: 1px;\\n  height: 1px;\\n  padding: 0;\\n  margin: -1px;\\n  overflow: hidden;\\n  clip: rect(0, 0, 0, 0);\\n  white-space: nowrap;\\n  border: 0;\\n}\\n"'));
var cem_pf_v6_badge_default = s13;

// elements/cem-pf-v6-badge/cem-pf-v6-badge.ts
var _screenReaderText_dec, _disabled_dec3, _unread_dec, _read_dec, _a9, _PfV6Badge_decorators, _init9, _read, _unread, _disabled3, _screenReaderText;
_PfV6Badge_decorators = [t3("cem-pf-v6-badge")];
var PfV6Badge = class extends (_a9 = i3, _read_dec = [n4({ type: Boolean, reflect: true })], _unread_dec = [n4({ type: Boolean, reflect: true })], _disabled_dec3 = [n4({ type: Boolean, reflect: true })], _screenReaderText_dec = [n4({ attribute: "screen-reader-text" })], _a9) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _read, __runInitializers(_init9, 8, this, false)), __runInitializers(_init9, 11, this);
    __privateAdd(this, _unread, __runInitializers(_init9, 12, this, false)), __runInitializers(_init9, 15, this);
    __privateAdd(this, _disabled3, __runInitializers(_init9, 16, this, false)), __runInitializers(_init9, 19, this);
    __privateAdd(this, _screenReaderText, __runInitializers(_init9, 20, this)), __runInitializers(_init9, 23, this);
  }
  render() {
    return T`
      <slot></slot>
      ${this.screenReaderText ? T`
        <span id="sr-text">${this.screenReaderText}</span>
      ` : ""}
    `;
  }
};
_init9 = __decoratorStart(_a9);
_read = new WeakMap();
_unread = new WeakMap();
_disabled3 = new WeakMap();
_screenReaderText = new WeakMap();
__decorateElement(_init9, 4, "read", _read_dec, PfV6Badge, _read);
__decorateElement(_init9, 4, "unread", _unread_dec, PfV6Badge, _unread);
__decorateElement(_init9, 4, "disabled", _disabled_dec3, PfV6Badge, _disabled3);
__decorateElement(_init9, 4, "screenReaderText", _screenReaderText_dec, PfV6Badge, _screenReaderText);
PfV6Badge = __decorateElement(_init9, 0, "PfV6Badge", _PfV6Badge_decorators, PfV6Badge);
__publicField(PfV6Badge, "styles", cem_pf_v6_badge_default);
__runInitializers(_init9, 1, PfV6Badge);

// lit-css:elements/cem-pf-v6-drawer/cem-pf-v6-drawer.css
var s14 = new CSSStyleSheet();
s14.replaceSync(JSON.parse('":host {\\n\\n  /* CSS Custom Properties from PatternFly design tokens */\\n  --cem-pf-v6-c-drawer__section--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-drawer__section--m-secondary--BackgroundColor: var(--pf-t--global--background--color--secondary--default);\\n  --cem-pf-v6-c-drawer__content--FlexBasis: 100%;\\n  --cem-pf-v6-c-drawer__content--BackgroundColor: transparent;\\n  --cem-pf-v6-c-drawer__content--m-primary--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-drawer__content--m-secondary--BackgroundColor: var(--pf-t--global--background--color--secondary--default);\\n  --cem-pf-v6-c-drawer__content--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-drawer__panel--MinWidth: 50%;\\n  --cem-pf-v6-c-drawer__panel--MaxHeight: auto;\\n  --cem-pf-v6-c-drawer__panel--ZIndex: var(--pf-t--global--z-index--sm);\\n  --cem-pf-v6-c-drawer__panel--BorderColor: var(--cem-pf-v6-c-drawer__panel--after--BackgroundColor);\\n  --cem-pf-v6-c-drawer__panel--BorderBlockStartWidth: 0;\\n  --cem-pf-v6-c-drawer__panel--BorderBlockEndWidth: 0;\\n  --cem-pf-v6-c-drawer__panel--BorderInlineStartWidth: var(--cem-pf-v6-c-drawer__panel--after--Width);\\n  --cem-pf-v6-c-drawer__panel--BorderInlineEndWidth: 0;\\n  --cem-pf-v6-c-drawer__panel--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --cem-pf-v6-c-drawer__panel--m-inline--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-drawer__panel--m-secondary--BackgroundColor: var(--pf-t--global--background--color--secondary--default);\\n  --cem-pf-v6-c-drawer__panel--RowGap: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-drawer__panel--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-drawer__panel--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-drawer__panel--TransitionDelay: 0s, var(--cem-pf-v6-c-drawer__panel--TransitionDelay--focus), var(--cem-pf-v6-c-drawer__panel--TransitionDelay--focus);\\n  --cem-pf-v6-c-drawer__panel--TransitionDelay--focus: var(--cem-pf-v6-c-drawer__panel--TransitionDuration--fade);\\n  --cem-pf-v6-c-drawer__panel--TransitionDelay--expand--focus: 0s;\\n  --cem-pf-v6-c-drawer__panel--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--decelerate);\\n  --cem-pf-v6-c-drawer__panel--TransitionDuration--fade: var(--pf-t--global--motion--duration--fade--default);\\n  --cem-pf-v6-c-drawer__panel--TransitionDuration: var(--cem-pf-v6-c-drawer__panel--TransitionDuration--fade), 0s, 0s;\\n  --cem-pf-v6-c-drawer__panel--TransitionProperty: opacity, visibility, transform;\\n  --cem-pf-v6-c-drawer__panel--Opacity: 0;\\n  --cem-pf-v6-c-drawer--m-expanded__panel--Opacity: 1;\\n  --cem-pf-v6-c-drawer__panel--FlexBasis: 100%;\\n  --cem-pf-v6-c-drawer__panel--md--FlexBasis--min: 1.5rem;\\n  --cem-pf-v6-c-drawer__panel--md--FlexBasis: 50%;\\n  --cem-pf-v6-c-drawer__panel--md--FlexBasis--max: 100%;\\n  --cem-pf-v6-c-drawer__panel--xl--MinWidth: 28.125rem;\\n  --cem-pf-v6-c-drawer__panel--xl--FlexBasis: 28.125rem;\\n  --cem-pf-v6-c-drawer--m-panel-bottom__panel--md--MinHeight: 50%;\\n  --cem-pf-v6-c-drawer--m-panel-bottom__panel--xl--MinHeight: 18.75rem;\\n  --cem-pf-v6-c-drawer--m-panel-bottom__panel--xl--FlexBasis: 18.75rem;\\n  --cem-pf-v6-c-drawer__panel--m-resizable--FlexDirection: row;\\n  --cem-pf-v6-c-drawer__panel--m-resizable--md--FlexBasis--min: var(--cem-pf-v6-c-drawer__splitter--m-vertical--Width);\\n  --cem-pf-v6-c-drawer__panel--m-resizable--MinWidth: 1.5rem;\\n  --cem-pf-v6-c-drawer--m-panel-bottom__panel--m-resizable--FlexDirection: column;\\n  --cem-pf-v6-c-drawer--m-panel-bottom__panel--m-resizable--md--FlexBasis--min: 1.5rem;\\n  --cem-pf-v6-c-drawer--m-panel-bottom__panel--m-resizable--MinHeight: 1.5rem;\\n  --cem-pf-v6-c-drawer__head--ColumnGap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-drawer__head--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-drawer__head--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-drawer__head--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__head--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__description--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__description--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__body--PaddingBlockStart--base: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__body--PaddingInlineEnd--base: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__body--PaddingBlockEnd--base: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__body--PaddingInlineStart--base: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__content__body--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-drawer__content__body--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-drawer__content__body--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-drawer__content__body--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-drawer__panel__body--PaddingBlockStart: var(--cem-pf-v6-c-drawer__body--PaddingBlockStart--base);\\n  --cem-pf-v6-c-drawer__panel__body--PaddingInlineEnd: var(--cem-pf-v6-c-drawer__body--PaddingInlineEnd--base);\\n  --cem-pf-v6-c-drawer__panel__body--PaddingBlockEnd: var(--cem-pf-v6-c-drawer__body--PaddingBlockEnd--base);\\n  --cem-pf-v6-c-drawer__panel__body--PaddingInlineStart: var(--cem-pf-v6-c-drawer__body--PaddingInlineStart--base);\\n  --cem-pf-v6-c-drawer__body--m-padding--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__body--m-padding--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__body--m-padding--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__body--m-padding--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-drawer__actions--MarginBlockStart: calc(var(--pf-t--global--spacer--control--vertical--compact) * -1.5);\\n  --cem-pf-v6-c-drawer__actions--MarginInlineEnd: calc(var(--pf-t--global--spacer--control--horizontal--compact) * -1.5);\\n  --cem-pf-v6-c-drawer__panel--BoxShadow: none;\\n  --cem-pf-v6-c-drawer--m-expanded--m-panel-bottom__panel--BoxShadow: var(--pf-t--global--box-shadow--md--top);\\n  --cem-pf-v6-c-drawer__panel--after--Width: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-drawer--m-inline__panel--after--Width: 0;\\n  --cem-pf-v6-c-drawer--m-inline__panel--after--md--Width: var(--pf-t--global--border--width--divider--default);\\n  --cem-pf-v6-c-drawer--m-panel-bottom__panel--after--Height: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-drawer__panel--after--BackgroundColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-drawer--m-inline--m-expanded__panel--after--BackgroundColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-drawer--m-inline__panel--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-drawer--m-panel-left--m-inline__panel--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-drawer--m-panel-bottom--m-inline__panel--PaddingBlockStart: 0;\\n\\n  /* Box shadow variations */\\n  --cem-pf-v6-c-drawer--m-expanded__panel--BoxShadow: var(--pf-t--global--box-shadow--md--left);\\n  --cem-pf-v6-c-drawer--m-expanded--m-panel-left__panel--BoxShadow: var(--pf-t--global--box-shadow--md--right);\\n\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n  overflow-x: hidden;\\n}\\n\\n@media screen and (prefers-reduced-motion: no-preference) {\\n  :host {\\n    --cem-pf-v6-c-drawer__panel--TransitionDuration--slide: var(--pf-t--global--motion--duration--slide-in--short);\\n    --cem-pf-v6-c-drawer__panel--TransitionDelay--focus: var(--cem-pf-v6-c-drawer__panel--TransitionDuration--slide);\\n    --cem-pf-v6-c-drawer__panel--TransitionDelay: 0s, 0s, 0s, 0s, var(--cem-pf-v6-c-drawer__panel--TransitionDelay--focus);\\n    --cem-pf-v6-c-drawer__panel--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--decelerate);\\n    --cem-pf-v6-c-drawer__panel--TransitionDuration: 0s, var(--cem-pf-v6-c-drawer__panel--TransitionDuration--slide), 0s, 0s, 0s;\\n    --cem-pf-v6-c-drawer__panel--TransitionProperty: margin, transform, box-shadow, flex-basis, visibility;\\n    --cem-pf-v6-c-drawer__panel--Opacity: 1;\\n  }\\n}\\n\\n@media screen and (min-width: 75rem) {\\n  :host {\\n    --cem-pf-v6-c-drawer__panel--MinWidth: var(--cem-pf-v6-c-drawer__panel--xl--MinWidth);\\n  }\\n\\n  :host([panel-position=\\"bottom\\"]) {\\n    --cem-pf-v6-c-drawer__panel--MinWidth: auto;\\n    --cem-pf-v6-c-drawer__panel--MinHeight: var(--cem-pf-v6-c-drawer--m-panel-bottom__panel--xl--MinHeight);\\n  }\\n}\\n\\n/* Inline and static variants */\\n:host([variant=\\"inline\\"]),\\n:host([variant=\\"static\\"]) {\\n  --cem-pf-v6-c-drawer__panel--BackgroundColor: var(--cem-pf-v6-c-drawer__panel--m-inline--BackgroundColor);\\n  --cem-pf-v6-c-drawer__panel--BorderInlineStartWidth: var(--cem-pf-v6-c-drawer--m-inline__panel--after--Width);\\n}\\n\\n:host([variant=\\"inline\\"]) #panel:not([no-border]),\\n:host([variant=\\"static\\"]) #panel:not([no-border]) {\\n  padding-inline-start: var(--cem-pf-v6-c-drawer--m-inline__panel--PaddingInlineStart);\\n}\\n\\n/* Panel position: start (left) */\\n:host([panel-position=\\"start\\"]) {\\n  --cem-pf-v6-c-drawer--m-expanded__panel--BoxShadow: var(--cem-pf-v6-c-drawer--m-expanded--m-panel-left__panel--BoxShadow);\\n}\\n\\n:host([panel-position=\\"start\\"]) #panel {\\n  order: 0;\\n  margin-inline-end: calc(var(--cem-pf-v6-c-drawer__panel--FlexBasis) * -1);\\n  transform: translateX(-100%);\\n}\\n\\n:host([panel-position=\\"start\\"]) #content {\\n  order: 1;\\n}\\n\\n/* Panel position: bottom */\\n:host([panel-position=\\"bottom\\"]) #main {\\n  flex-direction: column;\\n}\\n\\n/* Expanded state */\\n:host([expanded]) {\\n  --cem-pf-v6-c-drawer__panel--TransitionDelay--focus: var(--cem-pf-v6-c-drawer__panel--TransitionDelay--expand--focus);\\n}\\n\\n:host([expanded]) #panel {\\n  --cem-pf-v6-c-drawer__panel--Opacity: var(--cem-pf-v6-c-drawer--m-expanded__panel--Opacity);\\n  transform: translateX(-100%);\\n  visibility: visible;\\n}\\n\\n:host([expanded][panel-position=\\"start\\"]) #panel {\\n  transform: translateX(0);\\n}\\n\\n:host([expanded][panel-position=\\"bottom\\"]) #panel {\\n  transform: translate(0, -100%);\\n}\\n\\n/* Main container */\\n#main {\\n  display: flex;\\n  flex: 1;\\n  overflow: hidden;\\n}\\n\\n/* Content area */\\n#content {\\n  display: flex;\\n  flex-direction: column;\\n  flex-shrink: 0;\\n  flex-basis: var(--cem-pf-v6-c-drawer__content--FlexBasis);\\n  order: 0;\\n  overflow: auto;\\n  z-index: var(--cem-pf-v6-c-drawer__content--ZIndex);\\n  background-color: var(--cem-pf-v6-c-drawer__content--BackgroundColor);\\n}\\n\\n#content-body {\\n  padding-block-start: var(--cem-pf-v6-c-drawer__content__body--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-drawer__content__body--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-drawer__content__body--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-drawer__content__body--PaddingInlineEnd);\\n}\\n\\n/* Panel */\\n#panel {\\n  position: relative;\\n  display: flex;\\n  flex-direction: column;\\n  flex-shrink: 0;\\n  flex-basis: var(--cem-pf-v6-c-drawer__panel--FlexBasis);\\n  row-gap: var(--cem-pf-v6-c-drawer__panel--RowGap);\\n  order: 1;\\n  max-height: var(--cem-pf-v6-c-drawer__panel--MaxHeight);\\n  overflow: auto;\\n  visibility: hidden;\\n  z-index: var(--cem-pf-v6-c-drawer__panel--ZIndex);\\n  background-color: var(--cem-pf-v6-c-drawer__panel--BackgroundColor);\\n  border: solid var(--cem-pf-v6-c-drawer__panel--BorderColor);\\n  border-block-start-width: var(--cem-pf-v6-c-drawer__panel--BorderBlockStartWidth);\\n  border-block-end-width: var(--cem-pf-v6-c-drawer__panel--BorderBlockEndWidth);\\n  border-inline-start-width: var(--cem-pf-v6-c-drawer__panel--BorderInlineStartWidth);\\n  border-inline-end-width: var(--cem-pf-v6-c-drawer__panel--BorderInlineEndWidth);\\n  box-shadow: var(--cem-pf-v6-c-drawer__panel--BoxShadow);\\n  opacity: var(--cem-pf-v6-c-drawer__panel--Opacity);\\n  transition-delay: var(--cem-pf-v6-c-drawer__panel--TransitionDelay);\\n  transition-timing-function: var(--cem-pf-v6-c-drawer__panel--TransitionTimingFunction);\\n  transition-duration: var(--cem-pf-v6-c-drawer__panel--TransitionDuration);\\n  transition-property: var(--cem-pf-v6-c-drawer__panel--TransitionProperty);\\n  -webkit-overflow-scrolling: touch;\\n  padding-block-start: var(--cem-pf-v6-c-drawer__panel--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-drawer__panel--PaddingBlockEnd);\\n}\\n\\n@media screen and (min-width: 48rem) {\\n  #panel {\\n    --cem-pf-v6-c-drawer__panel--FlexBasis: max(\\n      var(--cem-pf-v6-c-drawer__panel--md--FlexBasis--min),\\n      min(var(--cem-pf-v6-c-drawer__panel--md--FlexBasis), var(--cem-pf-v6-c-drawer__panel--md--FlexBasis--max))\\n    );\\n    box-shadow: var(--cem-pf-v6-c-drawer--m-expanded__panel--BoxShadow);\\n  }\\n}\\n\\n@media screen and (min-width: 75rem) {\\n  #panel {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: var(--cem-pf-v6-c-drawer__panel--xl--FlexBasis);\\n  }\\n\\n  :host([panel-position=\\"bottom\\"]) #panel {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: var(--cem-pf-v6-c-drawer--m-panel-bottom__panel--xl--FlexBasis);\\n  }\\n}\\n\\n/* Panel head */\\n#panel-head {\\n  display: grid;\\n  grid-template-columns: auto;\\n  grid-auto-columns: max-content;\\n  column-gap: var(--cem-pf-v6-c-drawer__head--ColumnGap);\\n  padding-block-start: var(--cem-pf-v6-c-drawer__head--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-drawer__head--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-drawer__head--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-drawer__head--PaddingInlineEnd);\\n}\\n\\n#panel-head \\u003e * {\\n  grid-column: 1;\\n}\\n\\n#actions {\\n  display: flex;\\n  grid-row: 1;\\n  grid-column: 2;\\n  align-self: baseline;\\n  margin-block-start: var(--cem-pf-v6-c-drawer__actions--MarginBlockStart);\\n  margin-inline-end: var(--cem-pf-v6-c-drawer__actions--MarginInlineEnd);\\n}\\n\\n/* Panel description */\\n#panel-description {\\n  padding-inline-start: var(--cem-pf-v6-c-drawer__description--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-drawer__description--PaddingInlineEnd);\\n}\\n\\n/* Panel body */\\n#panel-body {\\n  padding-block-start: var(--cem-pf-v6-c-drawer__panel__body--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-drawer__panel__body--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-drawer__panel__body--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-drawer__panel__body--PaddingInlineEnd);\\n}\\n\\n#panel-body:last-child {\\n  flex: 1 1;\\n}\\n\\n/* No border modifier */\\n:host([no-border]) #panel {\\n  --cem-pf-v6-c-drawer--m-expanded__panel--BoxShadow: none;\\n  --cem-pf-v6-c-drawer__panel--BorderBlockStartWidth: 0;\\n  --cem-pf-v6-c-drawer__panel--BorderBlockEndWidth: 0;\\n  --cem-pf-v6-c-drawer__panel--BorderInlineStartWidth: 0;\\n  --cem-pf-v6-c-drawer__panel--BorderInlineEndWidth: 0;\\n}\\n\\n/* Panel width modifiers */\\n@media (min-width: 48rem) {\\n  :host([panel-width=\\"25\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 25%;\\n  }\\n\\n  :host([panel-width=\\"33\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 33%;\\n  }\\n\\n  :host([panel-width=\\"50\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 50%;\\n  }\\n\\n  :host([panel-width=\\"66\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 66%;\\n  }\\n\\n  :host([panel-width=\\"75\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 75%;\\n  }\\n\\n  :host([panel-width=\\"100\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 100%;\\n  }\\n}\\n\\n@media (min-width: 62rem) {\\n  :host([panel-width-on-lg=\\"25\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 25%;\\n  }\\n\\n  :host([panel-width-on-lg=\\"33\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 33%;\\n  }\\n\\n  :host([panel-width-on-lg=\\"50\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 50%;\\n  }\\n\\n  :host([panel-width-on-lg=\\"66\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 66%;\\n  }\\n\\n  :host([panel-width-on-lg=\\"75\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 75%;\\n  }\\n\\n  :host([panel-width-on-lg=\\"100\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 100%;\\n  }\\n}\\n\\n@media (min-width: 75rem) {\\n  :host([panel-width-on-xl=\\"25\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 25%;\\n  }\\n\\n  :host([panel-width-on-xl=\\"33\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 33%;\\n  }\\n\\n  :host([panel-width-on-xl=\\"50\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 50%;\\n  }\\n\\n  :host([panel-width-on-xl=\\"66\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 66%;\\n  }\\n\\n  :host([panel-width-on-xl=\\"75\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 75%;\\n  }\\n\\n  :host([panel-width-on-xl=\\"100\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 100%;\\n  }\\n}\\n\\n@media (min-width: 90.625rem) {\\n  :host([panel-width-on-2xl=\\"25\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 25%;\\n  }\\n\\n  :host([panel-width-on-2xl=\\"33\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 33%;\\n  }\\n\\n  :host([panel-width-on-2xl=\\"50\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 50%;\\n  }\\n\\n  :host([panel-width-on-2xl=\\"66\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 66%;\\n  }\\n\\n  :host([panel-width-on-2xl=\\"75\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 75%;\\n  }\\n\\n  :host([panel-width-on-2xl=\\"100\\"]) {\\n    --cem-pf-v6-c-drawer__panel--md--FlexBasis: 100%;\\n  }\\n}\\n\\n/* Inline variant media queries */\\n@media (min-width: 48rem) {\\n  :host {\\n    min-width: var(--cem-pf-v6-c-drawer__panel--MinWidth);\\n    --cem-pf-v6-c-drawer--m-inline__panel--after--Width: var(--cem-pf-v6-c-drawer--m-inline__panel--after--md--Width);\\n  }\\n\\n  :host([panel-position=\\"start\\"]) {\\n    --cem-pf-v6-c-drawer__panel--BorderInlineEndWidth: var(--cem-pf-v6-c-drawer__panel--after--Width);\\n    --cem-pf-v6-c-drawer__panel--BorderInlineStartWidth: 0;\\n  }\\n\\n  :host([panel-position=\\"start\\"][variant=\\"inline\\"]) #panel:not([no-border]),\\n  :host([panel-position=\\"start\\"][variant=\\"static\\"]) #panel:not([no-border]) {\\n    --cem-pf-v6-c-drawer__panel--BorderInlineEndWidth: var(--cem-pf-v6-c-drawer--m-inline__panel--after--Width);\\n    padding-inline-start: 0;\\n    padding-inline-end: var(--cem-pf-v6-c-drawer--m-panel-left--m-inline__panel--PaddingInlineEnd);\\n  }\\n\\n  :host([expanded][panel-position=\\"start\\"]) #panel {\\n    transform: translateX(0);\\n  }\\n\\n  :host([panel-position=\\"bottom\\"]) {\\n    --cem-pf-v6-c-drawer--m-expanded__panel--BoxShadow: var(--cem-pf-v6-c-drawer--m-expanded--m-panel-bottom__panel--BoxShadow);\\n    --cem-pf-v6-c-drawer__panel--MaxHeight: 100%;\\n    --cem-pf-v6-c-drawer__panel--FlexBasis--min: var(--cem-pf-v6-c-drawer--m-panel-bottom__panel--FlexBasis--min);\\n    --cem-pf-v6-c-drawer__panel--BorderInlineStartWidth: 0;\\n    --cem-pf-v6-c-drawer__panel--BorderBlockStartWidth: var(--cem-pf-v6-c-drawer--m-panel-bottom__panel--after--Height);\\n    min-width: auto;\\n    min-height: var(--cem-pf-v6-c-drawer--m-panel-bottom__panel--md--MinHeight);\\n  }\\n\\n  :host([panel-position=\\"bottom\\"][variant=\\"inline\\"]) #panel:not([no-border]),\\n  :host([panel-position=\\"bottom\\"][variant=\\"static\\"]) #panel:not([no-border]) {\\n    --cem-pf-v6-c-drawer__panel--BorderBlockStartWidth: var(--cem-pf-v6-c-drawer--m-inline__panel--after--Width);\\n    padding-block-start: var(--cem-pf-v6-c-drawer--m-panel-bottom--m-inline__panel--PaddingBlockStart);\\n    padding-inline-start: 0;\\n  }\\n\\n  /* Inline and static variants at md+ */\\n  :host([variant=\\"inline\\"]) #content,\\n  :host([variant=\\"static\\"]) #content {\\n    flex-shrink: 1;\\n  }\\n\\n  :host([variant=\\"inline\\"]) #panel,\\n  :host([variant=\\"static\\"]) #panel {\\n    --cem-pf-v6-c-drawer--m-expanded__panel--BoxShadow: none;\\n  }\\n\\n  :host([variant=\\"inline\\"]) #panel:not([no-border]),\\n  :host([variant=\\"static\\"]) #panel:not([no-border]) {\\n    --cem-pf-v6-c-drawer__panel--BorderColor: var(--cem-pf-v6-c-drawer--m-inline--m-expanded__panel--after--BackgroundColor);\\n  }\\n\\n  :host([variant=\\"inline\\"]) #content {\\n    overflow-x: auto;\\n  }\\n\\n  :host([variant=\\"inline\\"]) #panel {\\n    margin-inline-start: calc(var(--cem-pf-v6-c-drawer__panel--FlexBasis) * -1);\\n    transform: translateX(100%);\\n  }\\n\\n  :host([variant=\\"inline\\"][expanded]) #panel {\\n    margin-inline-start: 0;\\n    transform: translateX(0);\\n  }\\n\\n  :host([variant=\\"inline\\"][panel-position=\\"start\\"]) #panel {\\n    margin-inline-start: 0;\\n    margin-inline-end: calc(var(--cem-pf-v6-c-drawer__panel--FlexBasis) * -1);\\n    transform: translateX(-100%);\\n  }\\n\\n  :host([variant=\\"inline\\"][panel-position=\\"start\\"][expanded]) #panel {\\n    margin-inline-end: 0;\\n    transform: translateX(0);\\n  }\\n\\n  :host([variant=\\"inline\\"][panel-position=\\"bottom\\"]) #panel {\\n    margin-block-end: calc(var(--cem-pf-v6-c-drawer__panel--FlexBasis) * -1);\\n    transform: translateY(100%);\\n  }\\n\\n  :host([variant=\\"inline\\"][panel-position=\\"bottom\\"][expanded]) #panel {\\n    margin-block-end: 0;\\n    transform: translateY(0);\\n  }\\n\\n  :host([variant=\\"static\\"]) #panel {\\n    transform: translateX(0);\\n  }\\n\\n  :host([variant=\\"static\\"][panel-position=\\"start\\"]) #panel {\\n    margin-inline-end: 0;\\n    transform: translateX(0);\\n  }\\n\\n  :host([variant=\\"static\\"][panel-position=\\"bottom\\"]) #panel {\\n    transform: translateX(0);\\n  }\\n}\\n"'));
var cem_pf_v6_drawer_default = s14;

// elements/cem-pf-v6-drawer/cem-pf-v6-drawer.ts
var PfDrawerExpandEvent = class extends Event {
  constructor() {
    super("expand", { bubbles: true });
  }
};
var PfDrawerCollapseEvent = class extends Event {
  constructor() {
    super("collapse", { bubbles: true });
  }
};
var PfDrawerCloseEvent = class extends Event {
  constructor() {
    super("close", { bubbles: true });
  }
};
var _resizable_dec, _noBorder_dec, _panelWidthOn2xl_dec, _panelWidthOnXl_dec, _panelWidthOnLg_dec, _panelWidth_dec, _variant_dec3, _panelPosition_dec, _expanded_dec2, _a10, _PfV6Drawer_decorators, _init10, _expanded2, _panelPosition, _variant3, _panelWidth, _panelWidthOnLg, _panelWidthOnXl, _panelWidthOn2xl, _noBorder, _resizable, _hasHeaderContent, _hasDescriptionContent, _hasBodyContent, _PfV6Drawer_instances, onHeaderSlotchange_fn, onDescriptionSlotchange_fn, onBodySlotchange_fn;
_PfV6Drawer_decorators = [t3("cem-pf-v6-drawer")];
var PfV6Drawer = class extends (_a10 = i3, _expanded_dec2 = [n4({ type: Boolean, reflect: true })], _panelPosition_dec = [n4({ reflect: true, attribute: "panel-position" })], _variant_dec3 = [n4({ reflect: true })], _panelWidth_dec = [n4({ reflect: true, attribute: "panel-width" })], _panelWidthOnLg_dec = [n4({ reflect: true, attribute: "panel-width-on-lg" })], _panelWidthOnXl_dec = [n4({ reflect: true, attribute: "panel-width-on-xl" })], _panelWidthOn2xl_dec = [n4({ reflect: true, attribute: "panel-width-on-2xl" })], _noBorder_dec = [n4({ type: Boolean, reflect: true, attribute: "no-border" })], _resizable_dec = [n4({ type: Boolean, reflect: true })], _a10) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Drawer_instances);
    __privateAdd(this, _expanded2, __runInitializers(_init10, 8, this, false)), __runInitializers(_init10, 11, this);
    __privateAdd(this, _panelPosition, __runInitializers(_init10, 12, this, "end")), __runInitializers(_init10, 15, this);
    __privateAdd(this, _variant3, __runInitializers(_init10, 16, this, "default")), __runInitializers(_init10, 19, this);
    __privateAdd(this, _panelWidth, __runInitializers(_init10, 20, this)), __runInitializers(_init10, 23, this);
    __privateAdd(this, _panelWidthOnLg, __runInitializers(_init10, 24, this)), __runInitializers(_init10, 27, this);
    __privateAdd(this, _panelWidthOnXl, __runInitializers(_init10, 28, this)), __runInitializers(_init10, 31, this);
    __privateAdd(this, _panelWidthOn2xl, __runInitializers(_init10, 32, this)), __runInitializers(_init10, 35, this);
    __privateAdd(this, _noBorder, __runInitializers(_init10, 36, this, false)), __runInitializers(_init10, 39, this);
    __privateAdd(this, _resizable, __runInitializers(_init10, 40, this, false)), __runInitializers(_init10, 43, this);
    __privateAdd(this, _hasHeaderContent, false);
    __privateAdd(this, _hasDescriptionContent, false);
    __privateAdd(this, _hasBodyContent, false);
  }
  /** Opens the drawer panel */
  open() {
    this.expanded = true;
  }
  /** Closes the drawer panel and fires close event */
  close() {
    this.expanded = false;
    this.dispatchEvent(new PfDrawerCloseEvent());
  }
  /** Toggles the drawer panel */
  toggle() {
    this.expanded = !this.expanded;
  }
  updated(changed) {
    if (changed.has("expanded")) {
      if (this.expanded) {
        this.dispatchEvent(new PfDrawerExpandEvent());
      } else if (changed.get("expanded") !== void 0) {
        this.dispatchEvent(new PfDrawerCollapseEvent());
      }
    }
  }
  render() {
    return T`
      <div id="main">
        <div id="content">
          <div id="content-body">
            <slot></slot>
          </div>
        </div>
        <div id="panel">
          <div id="panel-head"
               ?hidden=${!__privateGet(this, _hasHeaderContent)}>
            <div id="panel-header-text">
              <slot name="panel-header"
                    @slotchange=${__privateMethod(this, _PfV6Drawer_instances, onHeaderSlotchange_fn)}></slot>
            </div>
            <div id="actions">
              <div id="close-wrapper">
                <cem-pf-v6-button id="close"
                              variant="plain"
                              aria-label="Close drawer panel"
                              @click=${this.close}>
                  <svg viewBox="0 0 352 512"
                       fill="currentColor"
                       role="presentation"
                       width="1em"
                       height="1em">
                    <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
                  </svg>
                </cem-pf-v6-button>
              </div>
            </div>
          </div>
          <div id="panel-description"
               part="panel-description"
               ?hidden=${!__privateGet(this, _hasDescriptionContent)}>
            <slot name="panel-description"
                  @slotchange=${__privateMethod(this, _PfV6Drawer_instances, onDescriptionSlotchange_fn)}></slot>
          </div>
          <div id="panel-body"
               part="panel-body"
               ?hidden=${!__privateGet(this, _hasBodyContent)}>
            <slot name="panel-body"
                  @slotchange=${__privateMethod(this, _PfV6Drawer_instances, onBodySlotchange_fn)}></slot>
          </div>
        </div>
      </div>
    `;
  }
};
_init10 = __decoratorStart(_a10);
_expanded2 = new WeakMap();
_panelPosition = new WeakMap();
_variant3 = new WeakMap();
_panelWidth = new WeakMap();
_panelWidthOnLg = new WeakMap();
_panelWidthOnXl = new WeakMap();
_panelWidthOn2xl = new WeakMap();
_noBorder = new WeakMap();
_resizable = new WeakMap();
_hasHeaderContent = new WeakMap();
_hasDescriptionContent = new WeakMap();
_hasBodyContent = new WeakMap();
_PfV6Drawer_instances = new WeakSet();
onHeaderSlotchange_fn = function(e6) {
  const slot = e6.target;
  __privateSet(this, _hasHeaderContent, slot.assignedNodes().length > 0);
  this.requestUpdate();
};
onDescriptionSlotchange_fn = function(e6) {
  const slot = e6.target;
  __privateSet(this, _hasDescriptionContent, slot.assignedNodes().length > 0);
  this.requestUpdate();
};
onBodySlotchange_fn = function(e6) {
  const slot = e6.target;
  __privateSet(this, _hasBodyContent, slot.assignedNodes().length > 0);
  this.requestUpdate();
};
__decorateElement(_init10, 4, "expanded", _expanded_dec2, PfV6Drawer, _expanded2);
__decorateElement(_init10, 4, "panelPosition", _panelPosition_dec, PfV6Drawer, _panelPosition);
__decorateElement(_init10, 4, "variant", _variant_dec3, PfV6Drawer, _variant3);
__decorateElement(_init10, 4, "panelWidth", _panelWidth_dec, PfV6Drawer, _panelWidth);
__decorateElement(_init10, 4, "panelWidthOnLg", _panelWidthOnLg_dec, PfV6Drawer, _panelWidthOnLg);
__decorateElement(_init10, 4, "panelWidthOnXl", _panelWidthOnXl_dec, PfV6Drawer, _panelWidthOnXl);
__decorateElement(_init10, 4, "panelWidthOn2xl", _panelWidthOn2xl_dec, PfV6Drawer, _panelWidthOn2xl);
__decorateElement(_init10, 4, "noBorder", _noBorder_dec, PfV6Drawer, _noBorder);
__decorateElement(_init10, 4, "resizable", _resizable_dec, PfV6Drawer, _resizable);
PfV6Drawer = __decorateElement(_init10, 0, "PfV6Drawer", _PfV6Drawer_decorators, PfV6Drawer);
__publicField(PfV6Drawer, "styles", cem_pf_v6_drawer_default);
__runInitializers(_init10, 1, PfV6Drawer);

// elements/node_modules/lit-html/node/directive-helpers.js
var { I: t5 } = Z;
var r6 = (o9) => void 0 === o9.strings;
var m3 = {};
var p3 = (o9, t6 = m3) => o9._$AH = t6;

// elements/node_modules/lit-html/node/directives/live.js
var l3 = e3(class extends i4 {
  constructor(r7) {
    if (super(r7), r7.type !== t4.PROPERTY && r7.type !== t4.ATTRIBUTE && r7.type !== t4.BOOLEAN_ATTRIBUTE) throw Error("The `live` directive is not allowed on child or event bindings");
    if (!r6(r7)) throw Error("`live` bindings can only contain a single expression");
  }
  render(r7) {
    return r7;
  }
  update(i6, [t6]) {
    if (t6 === E || t6 === A) return t6;
    const o9 = i6.element, l4 = i6.name;
    if (i6.type === t4.PROPERTY) {
      if (t6 === o9[l4]) return E;
    } else if (i6.type === t4.BOOLEAN_ATTRIBUTE) {
      if (!!t6 === o9.hasAttribute(l4)) return E;
    } else if (i6.type === t4.ATTRIBUTE && o9.getAttribute(l4) === t6 + "") return E;
    return p3(i6), t6;
  }
});

// lit-css:elements/cem-pf-v6-text-input-group/cem-pf-v6-text-input-group.css
var s15 = new CSSStyleSheet();
s15.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-text-input-group--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-text-input-group--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-text-input-group--m-success--BorderColor: var(--pf-t--global--border--color--status--success--default);\\n  --cem-pf-v6-c-text-input-group--m-warning--BorderColor: var(--pf-t--global--border--color--status--warning--default);\\n  --cem-pf-v6-c-text-input-group--m-error--BorderColor: var(--pf-t--global--border--color--status--danger--default);\\n  --cem-pf-v6-c-text-input-group--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-text-input-group__LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-text-input-group__FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-text-input-group--m-hover--BorderColor: var(--pf-t--global--border--color--hover);\\n  --cem-pf-v6-c-text-input-group--m-hover--m-success--BorderColor: var(--pf-t--global--border--color--status--success--hover);\\n  --cem-pf-v6-c-text-input-group--m-hover--m-warning--BorderColor: var(--pf-t--global--border--color--status--warning--hover);\\n  --cem-pf-v6-c-text-input-group--m-hover--m-error--BorderColor: var(--pf-t--global--border--color--status--danger--hover);\\n  --cem-pf-v6-c-text-input-group__main--first-child--not--text-input--MarginInlineStart: calc(var(--pf-t--global--spacer--control--horizontal--plain) / 2);\\n  --cem-pf-v6-c-text-input-group__main--m-icon__text-input--PaddingInlineStart: calc(var(--pf-t--global--spacer--control--horizontal--default) + var(--cem-pf-v6-c-text-input-group__icon--FontSize) + var(--pf-t--global--spacer--gap--text-to-element--default));\\n  --cem-pf-v6-c-text-input-group--status__text-input--PaddingInlineEnd: calc(var(--pf-t--global--spacer--control--horizontal--default) + var(--cem-pf-v6-c-text-input-group__icon--FontSize) + var(--pf-t--global--spacer--gap--text-to-element--default));\\n  --cem-pf-v6-c-text-input-group--utilities--status__text-input--PaddingInlineEnd: calc(var(--pf-t--global--spacer--sm) + var(--cem-pf-v6-c-text-input-group__icon--FontSize) + var(--pf-t--global--spacer--gap--text-to-element--default));\\n  --cem-pf-v6-c-text-input-group__main--RowGap: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-text-input-group__main--ColumnGap: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-text-input-group__text--BorderRadius--base: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-text-input-group__text--BorderStartStartRadius: var(--cem-pf-v6-c-text-input-group__text--BorderRadius--base);\\n  --cem-pf-v6-c-text-input-group__text--BorderStartEndRadius: var(--cem-pf-v6-c-text-input-group__text--BorderRadius--base);\\n  --cem-pf-v6-c-text-input-group__text--BorderEndEndRadius: var(--cem-pf-v6-c-text-input-group__text--BorderRadius--base);\\n  --cem-pf-v6-c-text-input-group__text--BorderEndStartRadius: var(--cem-pf-v6-c-text-input-group__text--BorderRadius--base);\\n  --cem-pf-v6-c-text-input-group__text-input--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-text-input-group__text-input--PaddingInlineEnd: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-text-input-group__text-input--PaddingBlockEnd: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-text-input-group__text-input--PaddingInlineStart: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-text-input-group__text-input--MinWidth: 12ch;\\n  --cem-pf-v6-c-text-input-group__text-input--m-hint--Color: var(--pf-t--global--text--color--placeholder);\\n  --cem-pf-v6-c-text-input-group__text-input--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-text-input-group__text-input--placeholder--Color: var(--pf-t--global--text--color--placeholder);\\n  --cem-pf-v6-c-text-input-group__text-input--BackgroundColor: transparent;\\n  --cem-pf-v6-c-text-input-group__text-input--OutlineOffset: -6px;\\n  --cem-pf-v6-c-text-input-group__icon--FontSize: var(--pf-t--global--icon--size--md);\\n  --cem-pf-v6-c-text-input-group__icon--InsetInlineStart: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-text-input-group__icon--m-status--InsetInlineEnd: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-text-input-group--utilities--icon--m-status--InsetInlineEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-text-input-group__icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-text-input-group--m-disabled__icon--Color: var(--pf-t--global--icon--color--on-disabled);\\n  --cem-pf-v6-c-text-input-group__icon--m-status--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-text-input-group--m-disabled__icon--m-status--Color: var(--pf-t--global--icon--color--on-disabled);\\n  --cem-pf-v6-c-text-input-group__main--m-success__icon--m-status--Color: var(--pf-t--global--icon--color--status--success--default);\\n  --cem-pf-v6-c-text-input-group__main--m-warning__icon--m-status--Color: var(--pf-t--global--icon--color--status--warning--default);\\n  --cem-pf-v6-c-text-input-group__main--m-error__icon--m-status--Color: var(--pf-t--global--icon--color--status--danger--default);\\n  --cem-pf-v6-c-text-input-group__icon--TranslateY: -50%;\\n  --cem-pf-v6-c-text-input-group__utilities--child--MarginInlineStart: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-text-input-group--m-disabled--Color: var(--pf-t--global--text--color--on-disabled);\\n  --cem-pf-v6-c-text-input-group--m-disabled--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n\\n  position: relative;\\n  display: flex;\\n  width: 100%;\\n  font-size: var(--cem-pf-v6-c-text-input-group__FontSize);\\n  line-height: var(--cem-pf-v6-c-text-input-group__LineHeight);\\n  color: var(--cem-pf-v6-c-text-input-group--Color, inherit);\\n  background-color: var(--cem-pf-v6-c-text-input-group--BackgroundColor);\\n  border-start-start-radius: var(--cem-pf-v6-c-text-input-group__text--BorderStartStartRadius);\\n  border-start-end-radius: var(--cem-pf-v6-c-text-input-group__text--BorderStartEndRadius);\\n  border-end-start-radius: var(--cem-pf-v6-c-text-input-group__text--BorderEndStartRadius);\\n  border-end-end-radius: var(--cem-pf-v6-c-text-input-group__text--BorderEndEndRadius);\\n}\\n\\n:host::before {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border: var(--cem-pf-v6-c-text-input-group--BorderWidth) solid var(--cem-pf-v6-c-text-input-group--BorderColor);\\n  border-start-start-radius: var(--cem-pf-v6-c-text-input-group__text--BorderStartStartRadius);\\n  border-start-end-radius: var(--cem-pf-v6-c-text-input-group__text--BorderStartEndRadius);\\n  border-end-start-radius: var(--cem-pf-v6-c-text-input-group__text--BorderEndStartRadius);\\n  border-end-end-radius: var(--cem-pf-v6-c-text-input-group__text--BorderEndEndRadius);\\n}\\n\\n:host([disabled]) {\\n  --cem-pf-v6-c-text-input-group--Color: var(--cem-pf-v6-c-text-input-group--m-disabled--Color);\\n  --cem-pf-v6-c-text-input-group__text-input--Color: var(--cem-pf-v6-c-text-input-group--m-disabled--Color);\\n  --cem-pf-v6-c-text-input-group__text-input--placeholder--Color: var(--cem-pf-v6-c-text-input-group--m-disabled--Color);\\n  --cem-pf-v6-c-text-input-group__icon--Color: var(--cem-pf-v6-c-text-input-group--m-disabled__icon--Color);\\n  --cem-pf-v6-c-text-input-group__icon--m-status--Color: var(--cem-pf-v6-c-text-input-group--m-disabled__icon--m-status--Color);\\n  --cem-pf-v6-c-text-input-group--BackgroundColor: var(--cem-pf-v6-c-text-input-group--m-disabled--BackgroundColor);\\n  pointer-events: none;\\n}\\n\\n:host([plain])::before {\\n  border: 0;\\n}\\n\\n:host([plain]) {\\n  --cem-pf-v6-c-text-input-group--BackgroundColor: transparent;\\n}\\n\\n:host([status=\\"success\\"]) {\\n  --cem-pf-v6-c-text-input-group--BorderColor: var(--cem-pf-v6-c-text-input-group--m-success--BorderColor);\\n  --cem-pf-v6-c-text-input-group--m-hover--BorderColor: var(--cem-pf-v6-c-text-input-group--m-hover--m-success--BorderColor);\\n  --cem-pf-v6-c-text-input-group__icon--m-status--Color: var(--cem-pf-v6-c-text-input-group__main--m-success__icon--m-status--Color);\\n}\\n\\n:host([status=\\"warning\\"]) {\\n  --cem-pf-v6-c-text-input-group--BorderColor: var(--cem-pf-v6-c-text-input-group--m-warning--BorderColor);\\n  --cem-pf-v6-c-text-input-group--m-hover--BorderColor: var(--cem-pf-v6-c-text-input-group--m-hover--m-warning--BorderColor);\\n  --cem-pf-v6-c-text-input-group__icon--m-status--Color: var(--cem-pf-v6-c-text-input-group__main--m-warning__icon--m-status--Color);\\n}\\n\\n:host([status=\\"error\\"]) {\\n  --cem-pf-v6-c-text-input-group--BorderColor: var(--cem-pf-v6-c-text-input-group--m-error--BorderColor);\\n  --cem-pf-v6-c-text-input-group--m-hover--BorderColor: var(--cem-pf-v6-c-text-input-group--m-hover--m-error--BorderColor);\\n  --cem-pf-v6-c-text-input-group__icon--m-status--Color: var(--cem-pf-v6-c-text-input-group__main--m-error__icon--m-status--Color);\\n}\\n\\n:host(:hover) {\\n  --cem-pf-v6-c-text-input-group--BorderColor: var(--cem-pf-v6-c-text-input-group--m-hover--BorderColor);\\n}\\n\\n:host([icon]) {\\n  --cem-pf-v6-c-text-input-group__text--Position: relative;\\n  --cem-pf-v6-c-text-input-group__text-input--PaddingInlineStart: var(--cem-pf-v6-c-text-input-group__main--m-icon__text-input--PaddingInlineStart);\\n}\\n\\n:host([status]) {\\n  --cem-pf-v6-c-text-input-group__text--Position: relative;\\n  --cem-pf-v6-c-text-input-group__text-input--PaddingInlineEnd: var(--cem-pf-v6-c-text-input-group--status__text-input--PaddingInlineEnd);\\n}\\n\\n:host(:has(#utilities slot[name=\\"utilities\\"]::slotted(*))) {\\n  --cem-pf-v6-c-text-input-group__icon--m-status--InsetInlineEnd: var(--cem-pf-v6-c-text-input-group--utilities--icon--m-status--InsetInlineEnd);\\n  --cem-pf-v6-c-text-input-group--status__text-input--PaddingInlineEnd: var(--cem-pf-v6-c-text-input-group--utilities--status__text-input--PaddingInlineEnd);\\n}\\n\\n#main {\\n  display: flex;\\n  flex: 1;\\n  flex-wrap: wrap;\\n  gap: var(--cem-pf-v6-c-text-input-group__main--RowGap) var(--cem-pf-v6-c-text-input-group__main--ColumnGap);\\n  min-width: 0;\\n}\\n\\n:host([icon]) #main {\\n  display: inline-flex;\\n  align-items: center;\\n  justify-content: center;\\n  min-width: calc(var(--cem-pf-v6-c-text-input-group__LineHeight) * var(--cem-pf-v6-c-text-input-group__FontSize));\\n}\\n\\n#text {\\n  position: var(--cem-pf-v6-c-text-input-group__text--Position, revert);\\n  display: inline-grid;\\n  flex: 1;\\n  grid-template-areas: \\"text-input\\";\\n  grid-template-columns: 1fr;\\n}\\n\\n#icon {\\n  display: none;\\n  position: absolute;\\n  inset-block-start: 50%;\\n  inset-inline-start: var(--cem-pf-v6-c-text-input-group__icon--InsetInlineStart);\\n  font-size: var(--cem-pf-v6-c-text-input-group__icon--FontSize);\\n  color: var(--cem-pf-v6-c-text-input-group__icon--Color);\\n  transform: translateY(var(--cem-pf-v6-c-text-input-group__icon--TranslateY));\\n  pointer-events: none;\\n}\\n\\n:host([icon]) #icon {\\n  display: block;\\n}\\n\\n#status-icon {\\n  display: none;\\n  position: absolute;\\n  inset-block-start: 50%;\\n  inset-inline-start: auto;\\n  inset-inline-end: var(--cem-pf-v6-c-text-input-group__icon--m-status--InsetInlineEnd);\\n  font-size: var(--cem-pf-v6-c-text-input-group__icon--FontSize);\\n  color: var(--cem-pf-v6-c-text-input-group__icon--m-status--Color);\\n  transform: translateY(var(--cem-pf-v6-c-text-input-group__icon--TranslateY));\\n  pointer-events: none;\\n}\\n\\n:host([status]) #status-icon {\\n  display: block;\\n}\\n\\n#input {\\n  overflow: hidden;\\n  text-overflow: ellipsis;\\n  white-space: nowrap;\\n  position: relative;\\n  width: 100%;\\n  min-width: var(--cem-pf-v6-c-text-input-group__text-input--MinWidth);\\n  box-sizing: border-box;\\n  padding-block-start: var(--cem-pf-v6-c-text-input-group__text-input--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-text-input-group__text-input--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-text-input-group__text-input--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-text-input-group__text-input--PaddingInlineEnd);\\n  color: var(--cem-pf-v6-c-text-input-group__text-input--Color);\\n  background-color: var(--cem-pf-v6-c-text-input-group__text-input--BackgroundColor);\\n  border: 0;\\n  outline-offset: var(--cem-pf-v6-c-text-input-group__text-input--OutlineOffset);\\n  grid-area: text-input;\\n}\\n\\n#input::placeholder {\\n  color: var(--cem-pf-v6-c-text-input-group__text-input--placeholder--Color);\\n}\\n\\n#utilities {\\n  display: flex;\\n  align-items: center;\\n  margin-inline-start: var(--cem-pf-v6-c-text-input-group__utilities--MarginInlineStart);\\n  margin-inline-end: var(--cem-pf-v6-c-text-input-group__utilities--MarginInlineEnd);\\n}\\n\\n#utilities::slotted(*) {\\n  margin-inline-start: var(--cem-pf-v6-c-text-input-group__utilities--child--MarginInlineStart);\\n}\\n\\n#utilities::slotted(:first-child) {\\n  margin-inline-start: 0;\\n}\\n"'));
var cem_pf_v6_text_input_group_default = s15;

// elements/cem-pf-v6-text-input-group/cem-pf-v6-text-input-group.ts
var _status_dec2, _icon_dec, _plain_dec, _required_dec, _readonly_dec, _disabled_dec4, _placeholder_dec, _type_dec2, _value_dec, _a11, _PfV6TextInputGroup_decorators, _internals2, _init11, _value, _type2, _placeholder, _disabled4, _readonly, _required, _plain, _icon, _status2, _PfV6TextInputGroup_instances, onInput_fn, onChange_fn;
_PfV6TextInputGroup_decorators = [t3("cem-pf-v6-text-input-group")];
var PfV6TextInputGroup = class extends (_a11 = i3, _value_dec = [n4({ reflect: true })], _type_dec2 = [n4({ reflect: true })], _placeholder_dec = [n4({ reflect: true })], _disabled_dec4 = [n4({ type: Boolean, reflect: true })], _readonly_dec = [n4({ type: Boolean, reflect: true })], _required_dec = [n4({ type: Boolean, reflect: true })], _plain_dec = [n4({ type: Boolean, reflect: true })], _icon_dec = [n4({ type: Boolean, reflect: true })], _status_dec2 = [n4({ reflect: true })], _a11) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6TextInputGroup_instances);
    __privateAdd(this, _internals2, this.attachInternals());
    __privateAdd(this, _value, __runInitializers(_init11, 8, this, "")), __runInitializers(_init11, 11, this);
    __privateAdd(this, _type2, __runInitializers(_init11, 12, this, "text")), __runInitializers(_init11, 15, this);
    __privateAdd(this, _placeholder, __runInitializers(_init11, 16, this)), __runInitializers(_init11, 19, this);
    __privateAdd(this, _disabled4, __runInitializers(_init11, 20, this, false)), __runInitializers(_init11, 23, this);
    __privateAdd(this, _readonly, __runInitializers(_init11, 24, this, false)), __runInitializers(_init11, 27, this);
    __privateAdd(this, _required, __runInitializers(_init11, 28, this, false)), __runInitializers(_init11, 31, this);
    __privateAdd(this, _plain, __runInitializers(_init11, 32, this, false)), __runInitializers(_init11, 35, this);
    __privateAdd(this, _icon, __runInitializers(_init11, 36, this, false)), __runInitializers(_init11, 39, this);
    __privateAdd(this, _status2, __runInitializers(_init11, 40, this)), __runInitializers(_init11, 43, this);
  }
  render() {
    return T`
      <div id="main">
        <span id="text">
          <span id="icon"><slot name="icon"></slot></span>
          <span id="status-icon"><slot name="status-icon"></slot></span>
          <input id="input"
                 type=${this.type}
                 .value=${l3(this.value)}
                 placeholder=${o7(this.placeholder)}
                 ?disabled=${this.disabled}
                 ?readonly=${this.readonly}
                 ?required=${this.required}
                 @input=${__privateMethod(this, _PfV6TextInputGroup_instances, onInput_fn)}
                 @change=${__privateMethod(this, _PfV6TextInputGroup_instances, onChange_fn)}>
        </span>
      </div>
      <div id="utilities">
        <slot name="utilities"></slot>
      </div>
    `;
  }
  focus() {
    this.shadowRoot?.getElementById("input")?.focus();
  }
  blur() {
    this.shadowRoot?.getElementById("input")?.blur();
  }
  select() {
    this.shadowRoot?.getElementById("input")?.select();
  }
};
_init11 = __decoratorStart(_a11);
_internals2 = new WeakMap();
_value = new WeakMap();
_type2 = new WeakMap();
_placeholder = new WeakMap();
_disabled4 = new WeakMap();
_readonly = new WeakMap();
_required = new WeakMap();
_plain = new WeakMap();
_icon = new WeakMap();
_status2 = new WeakMap();
_PfV6TextInputGroup_instances = new WeakSet();
onInput_fn = function(e6) {
  const input = e6.target;
  this.value = input.value;
  this.dispatchEvent(new InputEvent("input", {
    bubbles: true,
    data: e6.data,
    inputType: e6.inputType
  }));
};
onChange_fn = function() {
  this.dispatchEvent(new Event("change", { bubbles: true }));
};
__decorateElement(_init11, 4, "value", _value_dec, PfV6TextInputGroup, _value);
__decorateElement(_init11, 4, "type", _type_dec2, PfV6TextInputGroup, _type2);
__decorateElement(_init11, 4, "placeholder", _placeholder_dec, PfV6TextInputGroup, _placeholder);
__decorateElement(_init11, 4, "disabled", _disabled_dec4, PfV6TextInputGroup, _disabled4);
__decorateElement(_init11, 4, "readonly", _readonly_dec, PfV6TextInputGroup, _readonly);
__decorateElement(_init11, 4, "required", _required_dec, PfV6TextInputGroup, _required);
__decorateElement(_init11, 4, "plain", _plain_dec, PfV6TextInputGroup, _plain);
__decorateElement(_init11, 4, "icon", _icon_dec, PfV6TextInputGroup, _icon);
__decorateElement(_init11, 4, "status", _status_dec2, PfV6TextInputGroup, _status2);
PfV6TextInputGroup = __decorateElement(_init11, 0, "PfV6TextInputGroup", _PfV6TextInputGroup_decorators, PfV6TextInputGroup);
__publicField(PfV6TextInputGroup, "formAssociated", true);
__publicField(PfV6TextInputGroup, "styles", cem_pf_v6_text_input_group_default);
__runInitializers(_init11, 1, PfV6TextInputGroup);

// lit-css:elements/cem-pf-v6-toolbar/cem-pf-v6-toolbar.css
var s16 = new CSSStyleSheet();
s16.replaceSync(JSON.parse('":host {\\n\\n  position: relative;\\n  display: grid;\\n  row-gap: var(--cem-pf-v6-c-toolbar--RowGap);\\n  width: var(--cem-pf-v6-c-toolbar--Width, auto);\\n  padding-block-start: var(--cem-pf-v6-c-toolbar--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-toolbar--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-toolbar--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-toolbar--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-toolbar--FontSize);\\n  line-height: var(--cem-pf-v6-c-toolbar--LineHeight);\\n  background-color: var(--cem-pf-v6-c-toolbar--BackgroundColor);\\n}\\n\\n:host([sticky]) {\\n  --cem-pf-v6-c-toolbar--BackgroundColor: var(--cem-pf-v6-c-toolbar--m-sticky--BackgroundColor);\\n\\n  position: sticky;\\n  inset-block-start: 0;\\n  z-index: var(--cem-pf-v6-c-toolbar--m-sticky--ZIndex);\\n  padding-block-start: var(--cem-pf-v6-c-toolbar--m-sticky--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-toolbar--m-sticky--PaddingBlockEnd);\\n  border-block-end: var(--cem-pf-v6-c-toolbar--m-sticky--BorderBlockEndWidth) solid var(--cem-pf-v6-c-toolbar--m-sticky--BorderBlockEndColor);\\n  box-shadow: var(--cem-pf-v6-c-toolbar--m-sticky--BoxShadow);\\n}\\n\\n:host([static]),\\n:host([static]) #content {\\n  position: static;\\n}\\n\\n:host([static]) #expandable-content {\\n  position: absolute;\\n}\\n\\n:host([full-height]) {\\n  --cem-pf-v6-c-toolbar--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-toolbar--PaddingBlockEnd: 0;\\n}\\n\\n:host([color-variant=\\"primary\\"]) {\\n  --cem-pf-v6-c-toolbar--BackgroundColor: var(--cem-pf-v6-c-toolbar--m-primary--BackgroundColor);\\n}\\n\\n:host([color-variant=\\"secondary\\"]) {\\n  --cem-pf-v6-c-toolbar--BackgroundColor: var(--cem-pf-v6-c-toolbar--m-secondary--BackgroundColor);\\n}\\n\\n:host([color-variant=\\"no-background\\"]) {\\n  --cem-pf-v6-c-toolbar--BackgroundColor: var(--cem-pf-v6-c-toolbar--m-no-background--BackgroundColor);\\n}\\n\\n#content {\\n  --cem-pf-v6-hidden-visible--visible--Display: grid;\\n  --cem-pf-v6-hidden-visible--hidden--Display: none;\\n  --cem-pf-v6-hidden-visible--Display: var(--cem-pf-v6-hidden-visible--visible--Display);\\n\\n  position: relative;\\n  display: var(--cem-pf-v6-hidden-visible--Display);\\n  row-gap: var(--cem-pf-v6-c-toolbar__content--RowGap);\\n  min-width: var(--cem-pf-v6-c-toolbar__content--MinWidth, auto);\\n  padding-inline-start: var(--cem-pf-v6-c-toolbar__content--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-toolbar__content--PaddingInlineEnd);\\n}\\n\\n#content-section {\\n  display: flex;\\n  flex-wrap: var(--cem-pf-v6-c-toolbar__content-section--FlexWrap, wrap);\\n  row-gap: var(--cem-pf-v6-c-toolbar__content-section--RowGap);\\n  column-gap: var(--cem-pf-v6-c-toolbar__content-section--ColumnGap);\\n  align-items: start;\\n}\\n\\n#expandable-content {\\n  position: absolute;\\n  inset-block-start: 100%;\\n  inset-inline-start: 0;\\n  z-index: var(--cem-pf-v6-c-toolbar__expandable-content--ZIndex);\\n  display: none;\\n  row-gap: var(--cem-pf-v6-c-toolbar__expandable-content--RowGap);\\n  width: 100%;\\n  padding-block-start: var(--cem-pf-v6-c-toolbar__expandable-content--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-toolbar__expandable-content--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-toolbar__expandable-content--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-toolbar__expandable-content--PaddingInlineEnd);\\n  background-color: var(--cem-pf-v6-c-toolbar__expandable-content--BackgroundColor);\\n  border-block-end: var(--cem-pf-v6-c-toolbar__expandable-content--BorderBlockEndWidth) solid var(--cem-pf-v6-c-toolbar__expandable-content--BorderBlockEndColor);\\n  box-shadow: var(--cem-pf-v6-c-toolbar__expandable-content--BoxShadow);\\n}\\n\\n@media screen and (min-width: 992px) {\\n  #expandable-content {\\n    position: static;\\n    padding: 0;\\n    border-block-end: 0;\\n    box-shadow: none;\\n  }\\n}\\n\\n:host([expandable][expanded]) #expandable-content {\\n  display: grid;\\n}\\n"'));
var cem_pf_v6_toolbar_default = s16;

// elements/cem-pf-v6-toolbar/cem-pf-v6-toolbar.ts
var _colorVariant_dec, _fullHeight_dec, _sticky_dec, _expanded_dec3, _expandable_dec, _a12, _PfV6Toolbar_decorators, _init12, _expandable, _expanded3, _sticky, _fullHeight, _colorVariant;
_PfV6Toolbar_decorators = [t3("cem-pf-v6-toolbar")];
var PfV6Toolbar = class extends (_a12 = i3, _expandable_dec = [n4({ type: Boolean, reflect: true })], _expanded_dec3 = [n4({ type: Boolean, reflect: true })], _sticky_dec = [n4({ type: Boolean, reflect: true })], _fullHeight_dec = [n4({ type: Boolean, reflect: true, attribute: "full-height" })], _colorVariant_dec = [n4({ reflect: true, attribute: "color-variant" })], _a12) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _expandable, __runInitializers(_init12, 8, this, false)), __runInitializers(_init12, 11, this);
    __privateAdd(this, _expanded3, __runInitializers(_init12, 12, this, false)), __runInitializers(_init12, 15, this);
    __privateAdd(this, _sticky, __runInitializers(_init12, 16, this, false)), __runInitializers(_init12, 19, this);
    __privateAdd(this, _fullHeight, __runInitializers(_init12, 20, this, false)), __runInitializers(_init12, 23, this);
    __privateAdd(this, _colorVariant, __runInitializers(_init12, 24, this)), __runInitializers(_init12, 27, this);
  }
  render() {
    return T`
      <div id="content">
        <div id="content-section">
          <slot></slot>
        </div>
        <div id="expandable-content"
             part="expandable-content">
          <slot name="expandable"></slot>
        </div>
      </div>
    `;
  }
};
_init12 = __decoratorStart(_a12);
_expandable = new WeakMap();
_expanded3 = new WeakMap();
_sticky = new WeakMap();
_fullHeight = new WeakMap();
_colorVariant = new WeakMap();
__decorateElement(_init12, 4, "expandable", _expandable_dec, PfV6Toolbar, _expandable);
__decorateElement(_init12, 4, "expanded", _expanded_dec3, PfV6Toolbar, _expanded3);
__decorateElement(_init12, 4, "sticky", _sticky_dec, PfV6Toolbar, _sticky);
__decorateElement(_init12, 4, "fullHeight", _fullHeight_dec, PfV6Toolbar, _fullHeight);
__decorateElement(_init12, 4, "colorVariant", _colorVariant_dec, PfV6Toolbar, _colorVariant);
PfV6Toolbar = __decorateElement(_init12, 0, "PfV6Toolbar", _PfV6Toolbar_decorators, PfV6Toolbar);
__publicField(PfV6Toolbar, "styles", cem_pf_v6_toolbar_default);
__runInitializers(_init12, 1, PfV6Toolbar);

// lit-css:elements/cem-pf-v6-toolbar-group/cem-pf-v6-toolbar-group.css
var s17 = new CSSStyleSheet();
s17.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-hidden-visible--visible--Display: flex;\\n  --cem-pf-v6-hidden-visible--hidden--Display: none;\\n  --cem-pf-v6-hidden-visible--Display: var(--cem-pf-v6-hidden-visible--visible--Display);\\n\\n  display: var(--cem-pf-v6-hidden-visible--Display);\\n  row-gap: var(--cem-pf-v6-c-toolbar__group--RowGap);\\n  column-gap: var(--cem-pf-v6-c-toolbar__group--ColumnGap);\\n  align-items: baseline;\\n}\\n\\n:host([variant=\\"filter-group\\"]) {\\n  column-gap: var(--cem-pf-v6-c-toolbar__group--m-filter-group--ColumnGap);\\n}\\n\\n:host([variant=\\"label-group\\"]) {\\n  flex: 1;\\n  flex-wrap: wrap;\\n  column-gap: var(--cem-pf-v6-c-toolbar__group--m-label-group--ColumnGap);\\n}\\n\\n:host([variant=\\"action-group\\"]) {\\n  column-gap: var(--cem-pf-v6-c-toolbar__group--m-action-group--ColumnGap);\\n}\\n\\n:host([variant=\\"action-group-plain\\"]) {\\n  column-gap: var(--cem-pf-v6-c-toolbar__group--m-action-group-plain--ColumnGap);\\n}\\n\\n:host([variant=\\"action-group-inline\\"]) {\\n  flex-wrap: wrap;\\n  column-gap: var(--cem-pf-v6-c-toolbar__group--m-action-group-inline--ColumnGap);\\n}\\n\\n:host([variant=\\"overflow-container\\"]) {\\n  flex: 1;\\n  min-width: var(--cem-pf-v6-c-toolbar__group--m-overflow-container--MinWidth);\\n  overflow: auto;\\n}\\n"'));
var cem_pf_v6_toolbar_group_default = s17;

// elements/cem-pf-v6-toolbar-group/cem-pf-v6-toolbar-group.ts
var _variant_dec4, _a13, _PfV6ToolbarGroup_decorators, _init13, _variant4;
_PfV6ToolbarGroup_decorators = [t3("cem-pf-v6-toolbar-group")];
var PfV6ToolbarGroup = class extends (_a13 = i3, _variant_dec4 = [n4({ reflect: true })], _a13) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _variant4, __runInitializers(_init13, 8, this)), __runInitializers(_init13, 11, this);
  }
  render() {
    return T`<slot></slot>`;
  }
};
_init13 = __decoratorStart(_a13);
_variant4 = new WeakMap();
__decorateElement(_init13, 4, "variant", _variant_dec4, PfV6ToolbarGroup, _variant4);
PfV6ToolbarGroup = __decorateElement(_init13, 0, "PfV6ToolbarGroup", _PfV6ToolbarGroup_decorators, PfV6ToolbarGroup);
__publicField(PfV6ToolbarGroup, "styles", cem_pf_v6_toolbar_group_default);
__runInitializers(_init13, 1, PfV6ToolbarGroup);

// lit-css:elements/cem-pf-v6-toolbar-item/cem-pf-v6-toolbar-item.css
var s18 = new CSSStyleSheet();
s18.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-hidden-visible--visible--Display: flex;\\n  --cem-pf-v6-hidden-visible--hidden--Display: none;\\n  --cem-pf-v6-hidden-visible--Display: var(--cem-pf-v6-hidden-visible--visible--Display);\\n\\n  display: var(--cem-pf-v6-hidden-visible--Display);\\n  row-gap: var(--cem-pf-v6-c-toolbar__item--RowGap);\\n  column-gap: var(--cem-pf-v6-c-toolbar__item--ColumnGap);\\n  width: var(--cem-pf-v6-c-toolbar__item--Width);\\n  min-width: var(--cem-pf-v6-c-toolbar__item--MinWidth);\\n  align-items: baseline;\\n}\\n\\n:host([variant=\\"label\\"]) {\\n  font-weight: var(--cem-pf-v6-c-toolbar__item--m-label--FontWeight);\\n}\\n\\n:host([variant=\\"pagination\\"]) {\\n  margin-inline-start: auto;\\n}\\n\\n:host([variant=\\"overflow-container\\"]) {\\n  flex: 1;\\n  min-width: var(--cem-pf-v6-c-toolbar__item--m-overflow-container--MinWidth);\\n  overflow: auto;\\n}\\n\\n:host([variant=\\"expand-all\\"][expanded]) ::slotted(.cem-pf-v6-c-toolbar__expand-all-icon) {\\n  transform: rotate(var(--cem-pf-v6-c-toolbar__item--m-expand-all--m-expanded__expand-all-icon--Rotate));\\n}\\n"'));
var cem_pf_v6_toolbar_item_default = s18;

// elements/cem-pf-v6-toolbar-item/cem-pf-v6-toolbar-item.ts
var _expanded_dec4, _variant_dec5, _a14, _PfV6ToolbarItem_decorators, _init14, _variant5, _expanded4;
_PfV6ToolbarItem_decorators = [t3("cem-pf-v6-toolbar-item")];
var PfV6ToolbarItem = class extends (_a14 = i3, _variant_dec5 = [n4({ reflect: true })], _expanded_dec4 = [n4({ type: Boolean, reflect: true })], _a14) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _variant5, __runInitializers(_init14, 8, this)), __runInitializers(_init14, 11, this);
    __privateAdd(this, _expanded4, __runInitializers(_init14, 12, this, false)), __runInitializers(_init14, 15, this);
  }
  render() {
    return T`<slot></slot>`;
  }
};
_init14 = __decoratorStart(_a14);
_variant5 = new WeakMap();
_expanded4 = new WeakMap();
__decorateElement(_init14, 4, "variant", _variant_dec5, PfV6ToolbarItem, _variant5);
__decorateElement(_init14, 4, "expanded", _expanded_dec4, PfV6ToolbarItem, _expanded4);
PfV6ToolbarItem = __decorateElement(_init14, 0, "PfV6ToolbarItem", _PfV6ToolbarItem_decorators, PfV6ToolbarItem);
__publicField(PfV6ToolbarItem, "styles", cem_pf_v6_toolbar_item_default);
__runInitializers(_init14, 1, PfV6ToolbarItem);

// lit-css:elements/cem-virtual-tree/cem-virtual-tree.css
var s19 = new CSSStyleSheet();
s19.replaceSync(JSON.parse('"/* Virtual Tree - Efficiently renders large tree structures */\\n\\n:host {\\n  display: block;\\n  height: 100%;\\n  overflow: hidden;\\n}\\n\\n#tree-container {\\n  position: relative;\\n  width: 100%;\\n}\\n\\n#viewport {\\n  height: 100%;\\n  overflow-y: auto;\\n  overflow-x: hidden;\\n  position: relative;\\n\\n  cem-pf-v6-tree-item {\\n    cursor: pointer;\\n\\n    \\u0026:hover {\\n      background-color: var(--pf-t--global--background--color--secondary--hover, #f5f5f5);\\n    }\\n\\n    /* Only outline the actually focused item, not ancestors */\\n    \\u0026:focus-visible {\\n      outline: 2px solid var(--pf-t--global--color--brand--default, #0066cc);\\n      outline-offset: -2px;\\n    }\\n  }\\n}\\n"'));
var cem_virtual_tree_default = s19;

// elements/node_modules/lit-html/node/directives/class-map.js
var e5 = e3(class extends i4 {
  constructor(t6) {
    if (super(t6), t6.type !== t4.ATTRIBUTE || "class" !== t6.name || t6.strings?.length > 2) throw Error("`classMap()` can only be used in the `class` attribute and must be the only part in the attribute.");
  }
  render(t6) {
    return " " + Object.keys(t6).filter((s58) => t6[s58]).join(" ") + " ";
  }
  update(s58, [i6]) {
    if (void 0 === this.st) {
      this.st = /* @__PURE__ */ new Set(), void 0 !== s58.strings && (this.nt = new Set(s58.strings.join(" ").split(/\s/).filter((t6) => "" !== t6)));
      for (const t6 in i6) i6[t6] && !this.nt?.has(t6) && this.st.add(t6);
      return this.render(i6);
    }
    const r7 = s58.element.classList;
    for (const t6 of this.st) t6 in i6 || (r7.remove(t6), this.st.delete(t6));
    for (const t6 in i6) {
      const s59 = !!i6[t6];
      s59 === this.st.has(t6) || this.nt?.has(t6) || (s59 ? (r7.add(t6), this.st.add(t6)) : (r7.remove(t6), this.st.delete(t6)));
    }
    return E;
  }
});

// lit-css:elements/cem-pf-v6-tree-item/cem-pf-v6-tree-item.css
var s20 = new CSSStyleSheet();
s20.replaceSync(JSON.parse('":host {\\n\\n  --_node-padding-block: var(--pf-t--global--spacer--sm, 0.5rem);\\n  --_node-padding-inline: var(--pf-t--global--spacer--sm, 0.5rem);\\n  --_indent: var(--pf-t--global--spacer--lg, 1rem);\\n\\n  --_node-color: var(--pf-t--global--text--color--subtle, #6a6e73);\\n  --_node-bg: transparent;\\n  --_node-hover-bg: var(--pf-t--global--background--color--primary--hover, #f0f0f0);\\n  --_node-current-color: var(--pf-t--global--text--color--regular, #151515);\\n  --_node-current-bg: var(--pf-t--global--background--color--primary--clicked, #e0e0e0);\\n\\n  --_node-border-radius: var(--pf-t--global--border--radius--small, 0.1875rem);\\n  --_node-border-color: var(--pf-t--global--border--color--high-contrast, transparent);\\n  --_node-border-width: var(--pf-t--global--border--width--action--plain--default, 0);\\n  --_node-hover-border-width: var(--pf-t--global--border--width--action--plain--hover, 0.0625rem);\\n  --_node-current-border-width: var(--pf-t--global--border--width--action--plain--clicked, 0.125rem);\\n\\n  --_toggle-color: var(--pf-t--global--icon--color--subtle, #6a6e73);\\n  --_toggle-expanded-color: var(--pf-t--global--icon--color--regular, #151515);\\n  --_toggle-padding: var(--pf-t--global--spacer--sm, 0.5rem) var(--pf-t--global--spacer--md, 1rem);\\n\\n  --_icon-color: var(--pf-t--global--icon--color--subtle, #6a6e73);\\n  --_icon-spacing: var(--pf-t--global--spacer--sm, 0.5rem);\\n  --_badge-spacing: var(--pf-t--global--spacer--sm, 0.5rem);\\n\\n  --_transition-duration: var(--pf-t--global--motion--duration--fade--short, 150ms);\\n  --_transition-timing: var(--pf-t--global--motion--timing-function--default, ease);\\n\\n  display: block;\\n}\\n\\n:host([hidden]) {\\n  display: none !important;\\n}\\n\\n#item {\\n  position: relative;\\n  margin: 0;\\n  padding: 0;\\n  list-style: none;\\n}\\n\\n#content {\\n  display: flex;\\n  align-items: center;\\n  background-color: var(--_node-bg);\\n  border-radius: var(--_node-border-radius);\\n\\n  \\u0026:hover {\\n    background-color: var(--_node-hover-bg);\\n  }\\n}\\n\\n#content:has(\\u003e #node.current) {\\n  background-color: var(--_node-current-bg);\\n}\\n\\n#node {\\n  position: relative;\\n  display: flex;\\n  flex: 1;\\n  align-items: flex-start;\\n  min-width: 0;\\n  padding-block: var(--_node-padding-block);\\n  padding-inline: var(--_node-padding-inline);\\n  color: var(--_node-color);\\n  background: transparent;\\n  border: 0;\\n  text-align: start;\\n  cursor: pointer;\\n\\n  \\u0026::after {\\n    position: absolute;\\n    inset: 0;\\n    pointer-events: none;\\n    content: \\"\\";\\n    border: var(--_node-border-width) solid var(--_node-border-color);\\n    border-radius: var(--_node-border-radius);\\n  }\\n\\n  \\u0026.current {\\n    color: var(--_node-current-color);\\n\\n    \\u0026::after {\\n      border-width: var(--_node-current-border-width);\\n    }\\n  }\\n}\\n\\n#content:hover #node::after {\\n  border-width: var(--_node-hover-border-width);\\n}\\n\\n#node-container {\\n  display: contents;\\n  flex-grow: 1;\\n}\\n\\n#toggle {\\n  display: none;\\n  align-items: center;\\n  justify-content: center;\\n  padding: var(--_toggle-padding);\\n  margin-block: calc(var(--_node-padding-block) * -1);\\n  color: var(--_toggle-color);\\n  background: transparent;\\n  border: 0;\\n  cursor: pointer;\\n}\\n\\n:host([has-children]) #toggle {\\n  display: inline-flex;\\n}\\n\\n:host([expanded]) #toggle {\\n  color: var(--_toggle-expanded-color);\\n}\\n\\n#toggle-icon {\\n  display: inline-block;\\n  min-width: 1em;\\n  text-align: center;\\n  transition: transform var(--_transition-duration) var(--_transition-timing);\\n  transform: rotate(0deg);\\n}\\n\\n:host([expanded]) #toggle-icon {\\n  transform: rotate(90deg);\\n}\\n\\n:host(:dir(rtl)) #toggle-icon {\\n  scale: -1 1;\\n}\\n\\n#icon {\\n  padding-inline-end: var(--_icon-spacing);\\n  color: var(--_icon-color);\\n\\n  \\u0026:empty {\\n    display: none;\\n  }\\n}\\n\\n#node-text {\\n  font-weight: inherit;\\n  color: inherit;\\n}\\n\\n#badge-container {\\n  margin-inline-start: var(--_badge-spacing);\\n\\n  \\u0026:empty {\\n    display: none;\\n  }\\n}\\n\\n#children {\\n  margin: 0;\\n  padding: 0;\\n  padding-inline-start: var(--_indent);\\n  list-style: none;\\n  max-height: 0;\\n  overflow: hidden;\\n  opacity: 0;\\n  transition:\\n    opacity var(--_transition-duration) var(--_transition-timing),\\n    max-height var(--_transition-duration) var(--_transition-timing);\\n}\\n\\n:host([expanded]) #children {\\n  max-height: 99999px;\\n  opacity: 1;\\n}\\n\\n@media (prefers-reduced-motion: reduce) {\\n  * {\\n    transition-duration: 0.01ms !important;\\n  }\\n}\\n"'));
var cem_pf_v6_tree_item_default = s20;

// elements/cem-pf-v6-tree-item/cem-pf-v6-tree-item.ts
var PfTreeItemSelectEvent = class extends Event {
  constructor() {
    super("select", { bubbles: true });
  }
};
var PfTreeItemExpandEvent = class extends Event {
  constructor() {
    super("expand", { bubbles: true });
  }
};
var PfTreeItemCollapseEvent = class extends Event {
  constructor() {
    super("collapse", { bubbles: true });
  }
};
var _hasChildren_dec, _current_dec, _expanded_dec5, _badge_dec, _icon_dec2, _label_dec, _a15, _PfV6TreeItem_decorators, _init15, _label, _icon2, _badge, _expanded5, _current, _hasChildren, _explicitHasChildren, _PfV6TreeItem_instances, onToggleClick_fn, onNodeClick_fn, onSlotChange_fn, updateIcon_fn;
_PfV6TreeItem_decorators = [t3("cem-pf-v6-tree-item")];
var PfV6TreeItem = class extends (_a15 = i3, _label_dec = [n4({ reflect: true })], _icon_dec2 = [n4()], _badge_dec = [n4()], _expanded_dec5 = [n4({ type: Boolean, reflect: true })], _current_dec = [n4({ type: Boolean, reflect: true })], _hasChildren_dec = [n4({ type: Boolean, reflect: true, attribute: "has-children" })], _a15) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6TreeItem_instances);
    __privateAdd(this, _label, __runInitializers(_init15, 8, this)), __runInitializers(_init15, 11, this);
    __privateAdd(this, _icon2, __runInitializers(_init15, 12, this)), __runInitializers(_init15, 15, this);
    __privateAdd(this, _badge, __runInitializers(_init15, 16, this)), __runInitializers(_init15, 19, this);
    __privateAdd(this, _expanded5, __runInitializers(_init15, 20, this, false)), __runInitializers(_init15, 23, this);
    __privateAdd(this, _current, __runInitializers(_init15, 24, this, false)), __runInitializers(_init15, 27, this);
    __privateAdd(this, _hasChildren, __runInitializers(_init15, 28, this, false)), __runInitializers(_init15, 31, this);
    /** Track if has-children was explicitly set by the user */
    __privateAdd(this, _explicitHasChildren, false);
  }
  connectedCallback() {
    super.connectedCallback();
    if (this.hasAttribute("has-children")) {
      __privateSet(this, _explicitHasChildren, true);
    }
  }
  render() {
    return T`
      <li id="item"
          role="treeitem"
          tabindex="-1"
          aria-expanded=${this.hasChildren ? String(this.expanded) : A}
          aria-selected=${String(this.current)}
          aria-label=${this.label ?? A}>
        <div id="content">
          <button id="toggle"
                  type="button"
                  aria-label="Toggle"
                  @click=${__privateMethod(this, _PfV6TreeItem_instances, onToggleClick_fn)}>
            <span id="toggle-icon">
              <svg class="cem-pf-v6-svg"
                   viewBox="0 0 256 512"
                   fill="currentColor"
                   role="presentation"
                   width="1em"
                   height="1em">
                <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
              </svg>
            </span>
          </button>
          <button id="node"
                  class=${e5({ current: this.current })}
                  @click=${__privateMethod(this, _PfV6TreeItem_instances, onNodeClick_fn)}>
            <span id="node-container">
              <span id="icon"></span>
              <span id="node-text">
                <slot name="label">${this.label}</slot>
              </span>
              <span id="badge-container">
                <cem-pf-v6-badge id="badge">
                  <slot name="badge">${this.badge ?? A}</slot>
                </cem-pf-v6-badge>
              </span>
            </span>
          </button>
        </div>
        <ul id="children" role="group">
          <slot @slotchange=${__privateMethod(this, _PfV6TreeItem_instances, onSlotChange_fn)}></slot>
        </ul>
      </li>
    `;
  }
  updated(changed) {
    if (changed.has("icon")) {
      __privateMethod(this, _PfV6TreeItem_instances, updateIcon_fn).call(this);
    }
    if (changed.has("expanded") && changed.get("expanded") !== void 0) {
      this.dispatchEvent(
        this.expanded ? new PfTreeItemExpandEvent() : new PfTreeItemCollapseEvent()
      );
    }
  }
  /** Toggle expanded state */
  toggle() {
    if (!this.hasChildren) return;
    this.expanded = !this.expanded;
  }
  /** Expand the item */
  expand() {
    if (!this.hasChildren) return;
    this.expanded = true;
  }
  /** Collapse the item */
  collapse() {
    if (!this.hasChildren) return;
    this.expanded = false;
  }
  /** Select this item */
  select() {
    this.current = true;
    this.dispatchEvent(new PfTreeItemSelectEvent());
  }
  /** Deselect this item */
  deselect() {
    this.current = false;
  }
  /** Set tabindex for roving tabindex pattern */
  setTabindex(value) {
    const item = this.shadowRoot?.getElementById("item");
    item?.setAttribute("tabindex", String(value));
  }
  /** Focus this item */
  focusItem() {
    const item = this.shadowRoot?.getElementById("item");
    item?.focus();
  }
};
_init15 = __decoratorStart(_a15);
_label = new WeakMap();
_icon2 = new WeakMap();
_badge = new WeakMap();
_expanded5 = new WeakMap();
_current = new WeakMap();
_hasChildren = new WeakMap();
_explicitHasChildren = new WeakMap();
_PfV6TreeItem_instances = new WeakSet();
onToggleClick_fn = function(e6) {
  e6.stopPropagation();
  this.toggle();
};
onNodeClick_fn = function() {
  if (this.hasChildren) {
    this.toggle();
  }
  this.select();
};
onSlotChange_fn = function() {
  if (__privateGet(this, _explicitHasChildren)) return;
  const slot = this.shadowRoot?.querySelector("#children slot");
  if (!slot) return;
  const children = slot.assignedElements();
  this.hasChildren = children.length > 0;
};
updateIcon_fn = function() {
  const iconEl = this.shadowRoot?.getElementById("icon");
  if (!iconEl) return;
  iconEl.innerHTML = this.icon ?? "";
};
__decorateElement(_init15, 4, "label", _label_dec, PfV6TreeItem, _label);
__decorateElement(_init15, 4, "icon", _icon_dec2, PfV6TreeItem, _icon2);
__decorateElement(_init15, 4, "badge", _badge_dec, PfV6TreeItem, _badge);
__decorateElement(_init15, 4, "expanded", _expanded_dec5, PfV6TreeItem, _expanded5);
__decorateElement(_init15, 4, "current", _current_dec, PfV6TreeItem, _current);
__decorateElement(_init15, 4, "hasChildren", _hasChildren_dec, PfV6TreeItem, _hasChildren);
PfV6TreeItem = __decorateElement(_init15, 0, "PfV6TreeItem", _PfV6TreeItem_decorators, PfV6TreeItem);
__publicField(PfV6TreeItem, "styles", cem_pf_v6_tree_item_default);
__runInitializers(_init15, 1, PfV6TreeItem);

// elements/cem-virtual-tree/cem-virtual-tree.ts
var ItemSelectEvent = class extends Event {
  item;
  constructor(item) {
    super("item-select", { bubbles: true, composed: true });
    this.item = item;
  }
};
var _CemVirtualTree_decorators, _init16, _a16;
_CemVirtualTree_decorators = [t3("cem-virtual-tree")];
var _CemVirtualTree = class _CemVirtualTree extends (_a16 = i3) {
  static styles = cem_virtual_tree_default;
  // Static cache for manifest (shared across all instances)
  static #manifestCache = null;
  static #manifestPromise = null;
  #manifest = null;
  #flatItems = [];
  #visibleItems = [];
  #searchQuery = "";
  #viewport = null;
  #currentSelectedElement = null;
  #currentSelectedItemId = null;
  render() {
    return T`
      <div id="tree-container">
        <div id="viewport" role="tree"></div>
      </div>
    `;
  }
  async firstUpdated() {
    this.#viewport = this.shadowRoot.getElementById("viewport");
    this.#manifest = await this.#loadManifest();
    if (!this.#manifest) {
      console.warn("[virtual-tree] Failed to load manifest");
      return;
    }
    this.#buildFlatList();
    this.#renderTree();
  }
  /**
   * Public static method to load manifest with caching
   * Can be called by other components to reuse the same cached manifest
   */
  static async loadManifest() {
    if (_CemVirtualTree.#manifestCache) {
      return _CemVirtualTree.#manifestCache;
    }
    if (_CemVirtualTree.#manifestPromise) {
      return _CemVirtualTree.#manifestPromise;
    }
    _CemVirtualTree.#manifestPromise = fetch("/custom-elements.json").then(async (response) => {
      if (!response.ok) {
        throw new Error(`Failed to fetch manifest: ${response.status}`);
      }
      const manifest = await response.json();
      _CemVirtualTree.#manifestCache = manifest;
      return manifest;
    }).catch((error) => {
      console.error("[virtual-tree] Error loading manifest:", error);
      _CemVirtualTree.#manifestPromise = null;
      return null;
    });
    return _CemVirtualTree.#manifestPromise;
  }
  /**
   * Clear the static manifest cache (for testing)
   */
  static clearCache() {
    _CemVirtualTree.#manifestCache = null;
    _CemVirtualTree.#manifestPromise = null;
  }
  /**
   * Load manifest from server with static caching
   */
  async #loadManifest() {
    return _CemVirtualTree.loadManifest();
  }
  /**
   * Build flat list from hierarchical manifest
   */
  #buildFlatList() {
    this.#flatItems = [];
    let id = 0;
    const hasMultiplePackages = this.#manifest.packages && this.#manifest.packages.length > 1;
    if (hasMultiplePackages) {
      for (const pkg of this.#manifest.packages) {
        const packageId = id++;
        const packageItem = {
          id: packageId,
          type: "package",
          label: pkg.name,
          depth: 0,
          hasChildren: (pkg.modules?.length ?? 0) > 0,
          expanded: false,
          visible: true,
          packageName: pkg.name,
          badge: pkg.modules?.length ?? 0
        };
        this.#flatItems.push(packageItem);
        if (!pkg.modules) continue;
        id = this.#buildModulesForPackage(pkg.modules, packageId, 1, id);
      }
    } else {
      const modules = this.#manifest.packages?.[0]?.modules ?? this.#manifest.modules;
      if (!modules) return;
      this.#buildModulesForPackage(modules, void 0, 0, id);
    }
    this.#updateVisibleItems();
  }
  /**
   * Build modules and their declarations for a package
   */
  #buildModulesForPackage(modules, parentId, depth, startId) {
    let id = startId;
    for (const module of modules) {
      const moduleId = id++;
      const moduleItem = {
        id: moduleId,
        type: "module",
        label: module.path,
        depth,
        parentId,
        hasChildren: (module.declarations?.length ?? 0) > 0,
        expanded: false,
        visible: parentId === void 0,
        // Visible if no parent, otherwise hidden until parent expands
        modulePath: module.path,
        badge: module.declarations?.length ?? 0
      };
      this.#flatItems.push(moduleItem);
      if (!module.declarations) continue;
      for (const decl of module.declarations) {
        if (decl.kind === "class" && decl.customElement && decl.tagName) {
          const ceId = id++;
          const properties = decl.members?.filter((m4) => m4.kind === "field") ?? [];
          const methods = decl.members?.filter((m4) => m4.kind === "method") ?? [];
          const hasChildren = [
            decl.attributes,
            properties,
            methods,
            decl.events,
            decl.slots,
            decl.cssProperties,
            decl.cssParts,
            decl.cssStates,
            decl.demos
          ].some((x2) => (x2?.length ?? 0) > 0);
          const ceItem = {
            id: ceId,
            type: "custom-element",
            label: `<${decl.tagName}>`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren,
            expanded: false,
            visible: false,
            modulePath: module.path,
            tagName: decl.tagName
          };
          this.#flatItems.push(ceItem);
          if (decl.attributes?.length) {
            const attrCatId = id++;
            this.#flatItems.push({
              id: attrCatId,
              type: "category",
              label: "Attributes",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.attributes.length,
              category: "attributes"
            });
            for (const attr of decl.attributes) {
              this.#flatItems.push({
                id: id++,
                type: "attribute",
                label: attr.name,
                depth: depth + 3,
                parentId: attrCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: attr.name
              });
            }
          }
          if (properties.length) {
            const propCatId = id++;
            this.#flatItems.push({
              id: propCatId,
              type: "category",
              label: "Properties",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: properties.length,
              category: "properties"
            });
            for (const prop of properties) {
              this.#flatItems.push({
                id: id++,
                type: "property",
                label: prop.name,
                depth: depth + 3,
                parentId: propCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: prop.name
              });
            }
          }
          if (methods.length) {
            const methodCatId = id++;
            this.#flatItems.push({
              id: methodCatId,
              type: "category",
              label: "Methods",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: methods.length,
              category: "methods"
            });
            for (const method of methods) {
              this.#flatItems.push({
                id: id++,
                type: "method",
                label: `${method.name}()`,
                depth: depth + 3,
                parentId: methodCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: method.name
              });
            }
          }
          if (decl.events?.length) {
            const eventCatId = id++;
            this.#flatItems.push({
              id: eventCatId,
              type: "category",
              label: "Events",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.events.length,
              category: "events"
            });
            for (const event of decl.events) {
              this.#flatItems.push({
                id: id++,
                type: "event",
                label: event.name,
                depth: depth + 3,
                parentId: eventCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: event.name
              });
            }
          }
          if (decl.slots?.length) {
            const slotCatId = id++;
            this.#flatItems.push({
              id: slotCatId,
              type: "category",
              label: "Slots",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.slots.length,
              category: "slots"
            });
            for (const slot of decl.slots) {
              this.#flatItems.push({
                id: id++,
                type: "slot",
                label: slot.name || "(default)",
                depth: depth + 3,
                parentId: slotCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: slot.name
              });
            }
          }
          if (decl.cssProperties?.length) {
            const cssPropCatId = id++;
            this.#flatItems.push({
              id: cssPropCatId,
              type: "category",
              label: "CSS Properties",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssProperties.length,
              category: "css-properties"
            });
            for (const cssProp of decl.cssProperties) {
              this.#flatItems.push({
                id: id++,
                type: "css-property",
                label: cssProp.name,
                depth: depth + 3,
                parentId: cssPropCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssProp.name
              });
            }
          }
          if (decl.cssParts?.length) {
            const cssPartCatId = id++;
            this.#flatItems.push({
              id: cssPartCatId,
              type: "category",
              label: "CSS Parts",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssParts.length,
              category: "css-parts"
            });
            for (const cssPart of decl.cssParts) {
              this.#flatItems.push({
                id: id++,
                type: "css-part",
                label: cssPart.name,
                depth: depth + 3,
                parentId: cssPartCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssPart.name
              });
            }
          }
          if (decl.cssStates?.length) {
            const cssStateCatId = id++;
            this.#flatItems.push({
              id: cssStateCatId,
              type: "category",
              label: "CSS States",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.cssStates.length,
              category: "css-states"
            });
            for (const cssState of decl.cssStates) {
              this.#flatItems.push({
                id: id++,
                type: "css-state",
                label: cssState.name,
                depth: depth + 3,
                parentId: cssStateCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                name: cssState.name
              });
            }
          }
          if (decl.demos?.length) {
            const demoCatId = id++;
            this.#flatItems.push({
              id: demoCatId,
              type: "category",
              label: "Demos",
              depth: depth + 2,
              parentId: ceId,
              hasChildren: true,
              expanded: false,
              visible: false,
              badge: decl.demos.length,
              category: "demos"
            });
            for (const demo of decl.demos) {
              this.#flatItems.push({
                id: id++,
                type: "demo",
                label: demo.name || demo.url || "(demo)",
                depth: depth + 3,
                parentId: demoCatId,
                hasChildren: false,
                expanded: false,
                visible: false,
                modulePath: module.path,
                tagName: decl.tagName,
                url: demo.url
              });
            }
          }
        } else if (decl.kind === "class") {
          this.#flatItems.push({
            id: id++,
            type: "class",
            label: decl.name,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name
          });
        } else if (decl.kind === "function") {
          this.#flatItems.push({
            id: id++,
            type: "function",
            label: `${decl.name}()`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name
          });
        } else if (decl.kind === "variable") {
          this.#flatItems.push({
            id: id++,
            type: "variable",
            label: decl.name,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name
          });
        } else if (decl.kind === "mixin") {
          this.#flatItems.push({
            id: id++,
            type: "mixin",
            label: `${decl.name}()`,
            depth: depth + 1,
            parentId: moduleId,
            hasChildren: false,
            expanded: false,
            visible: false,
            modulePath: module.path,
            name: decl.name
          });
        }
      }
    }
    return id;
  }
  /**
   * Update visible items based on expanded state
   */
  #updateVisibleItems() {
    this.#visibleItems = this.#flatItems.filter((item) => item.visible);
  }
  /**
   * Render visible tree items with proper nesting
   */
  #renderTree() {
    if (!this.#viewport) return;
    this.#viewport.innerHTML = "";
    const rootItems = this.#visibleItems.filter((item) => item.depth === 0);
    for (const rootItem of rootItems) {
      const treeItemEl = this.#createTreeItemWithChildren(rootItem);
      if (treeItemEl) {
        this.#viewport.appendChild(treeItemEl);
      }
    }
  }
  /**
   * Create a tree item element with its nested children
   */
  #createTreeItemWithChildren(item) {
    const el = document.createElement("cem-pf-v6-tree-item");
    el.setAttribute("label", item.label);
    if (item.badge) {
      el.setAttribute("badge", String(item.badge));
    }
    el.dataset.itemId = String(item.id);
    el.dataset.type = item.type;
    if (this.#currentSelectedItemId === item.id) {
      el.setAttribute("current", "");
      this.#currentSelectedElement = el;
    }
    if (item.hasChildren) {
      el.setAttribute("has-children", "");
      el.expanded = item.expanded;
      el.addEventListener("expand", (e6) => {
        e6.stopPropagation();
        this.#handleToggle(item, true);
      });
      el.addEventListener("collapse", (e6) => {
        e6.stopPropagation();
        this.#handleToggle(item, false);
      });
    }
    if (item.type !== "category") {
      el.addEventListener("select", (e6) => {
        e6.stopPropagation();
        this.#handleSelect(item, el);
      });
    }
    if (item.expanded && item.hasChildren) {
      const children = this.#visibleItems.filter((child) => child.parentId === item.id && child.visible);
      for (const child of children) {
        const childEl = this.#createTreeItemWithChildren(child);
        if (childEl) {
          el.appendChild(childEl);
        }
      }
    }
    return el;
  }
  /**
   * Handle tree item toggle (expand/collapse)
   */
  #handleToggle(item, expanded) {
    item.expanded = expanded;
    this.#updateChildVisibility(item);
    this.#updateVisibleItems();
    this.#renderTree();
  }
  /**
   * Update visibility of children when parent is toggled
   */
  #updateChildVisibility(parentItem) {
    const isExpanded = parentItem.expanded;
    for (const item of this.#flatItems) {
      if (item.parentId === parentItem.id) {
        item.visible = isExpanded;
        if (!isExpanded) {
          item.expanded = false;
          this.#updateChildVisibility(item);
        }
      }
    }
  }
  /**
   * Handle tree item selection
   */
  #handleSelect(item, element) {
    if (this.#currentSelectedElement) {
      this.#currentSelectedElement.removeAttribute("current");
    }
    element.setAttribute("current", "");
    this.#currentSelectedElement = element;
    this.#currentSelectedItemId = item.id;
    this.dispatchEvent(new ItemSelectEvent(item));
  }
  /**
   * Search/filter tree items
   */
  search(query) {
    this.#searchQuery = query.toLowerCase().trim();
    if (!this.#searchQuery) {
      for (const item of this.#flatItems) {
        item.visible = item.depth === 0 || item.parentId !== void 0 && this.#isParentExpanded(item);
      }
    } else {
      const matchingItems = /* @__PURE__ */ new Set();
      for (const item of this.#flatItems) {
        if (item.label.toLowerCase().includes(this.#searchQuery)) {
          matchingItems.add(item);
          this.#showAncestors(item);
        }
      }
      for (const item of this.#flatItems) {
        if (!matchingItems.has(item) && !this.#hasMatchingDescendant(item, matchingItems)) {
          item.visible = this.#isAncestorOfMatch(item, matchingItems);
        } else {
          item.visible = true;
        }
      }
    }
    this.#updateVisibleItems();
    this.#renderTree();
  }
  /**
   * Check if parent is expanded
   */
  #isParentExpanded(item) {
    if (item.parentId === void 0) return true;
    const parent = this.#flatItems.find((i6) => i6.id === item.parentId);
    return !!parent && parent.expanded && this.#isParentExpanded(parent);
  }
  /**
   * Show all ancestors of an item
   */
  #showAncestors(item) {
    if (item.parentId === void 0) return;
    const parent = this.#flatItems.find((i6) => i6.id === item.parentId);
    if (parent) {
      parent.visible = true;
      parent.expanded = true;
      this.#showAncestors(parent);
    }
  }
  /**
   * Check if item is ancestor of a match
   */
  #isAncestorOfMatch(item, matchingItems) {
    for (const match of matchingItems) {
      let current = match;
      while (current.parentId !== void 0) {
        const parent = this.#flatItems.find((i6) => i6.id === current.parentId);
        if (!parent) break;
        if (parent === item) return true;
        current = parent;
      }
    }
    return false;
  }
  /**
   * Check if item has matching descendant
   */
  #hasMatchingDescendant(item, matchingItems) {
    for (const match of matchingItems) {
      if (this.#isDescendantOf(match, item)) return true;
    }
    return false;
  }
  /**
   * Check if item is descendant of another
   */
  #isDescendantOf(item, ancestor) {
    let current = item;
    while (current.parentId !== void 0) {
      const parent = this.#flatItems.find((i6) => i6.id === current.parentId);
      if (!parent) return false;
      if (parent === ancestor) return true;
      current = parent;
    }
    return false;
  }
  /**
   * Expand all items
   */
  expandAll() {
    for (const item of this.#flatItems) {
      if (item.hasChildren) {
        item.expanded = true;
      }
      if (!this.#searchQuery || item.label.toLowerCase().includes(this.#searchQuery)) {
        item.visible = true;
      }
    }
    this.#updateVisibleItems();
    this.#renderTree();
  }
  /**
   * Collapse all items
   */
  collapseAll() {
    for (const item of this.#flatItems) {
      if (item.hasChildren) {
        item.expanded = false;
      }
      item.visible = item.depth === 0;
    }
    this.#updateVisibleItems();
    this.#renderTree();
  }
};
_init16 = __decoratorStart(_a16);
_CemVirtualTree = __decorateElement(_init16, 0, "CemVirtualTree", _CemVirtualTree_decorators, _CemVirtualTree);
__runInitializers(_init16, 1, _CemVirtualTree);
var CemVirtualTree = _CemVirtualTree;

// elements/cem-manifest-browser/cem-manifest-browser.ts
var _CemManifestBrowser_decorators, _init17, _a17;
_CemManifestBrowser_decorators = [t3("cem-manifest-browser")];
var CemManifestBrowser = class extends (_a17 = i3) {
  static styles = cem_manifest_browser_default;
  #searchDebounceTimer = null;
  disconnectedCallback() {
    super.disconnectedCallback();
    if (this.#searchDebounceTimer != null) {
      clearTimeout(this.#searchDebounceTimer);
      this.#searchDebounceTimer = null;
    }
  }
  render() {
    return T`
      <div id="drawer-content" slot="content">
        <cem-pf-v6-toolbar sticky>
          <cem-pf-v6-toolbar-group>
            <cem-pf-v6-toolbar-item>
              <cem-pf-v6-text-input-group id="search"
                                      type="search"
                                      placeholder="Search manifest..."
                                      aria-label="Search manifest"
                                      icon
                                      @input=${this.#onSearchInput}>
                <svg slot="icon"
                     role="presentation"
                     fill="currentColor"
                     height="1em"
                     width="1em"
                     viewBox="0 0 512 512">
                  <path d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                </svg>
                <cem-pf-v6-badge slot="utilities"
                             id="search-count"
                             compact
                             hidden>
                  0<span class="cem-pf-v6-screen-reader"> results</span>
                </cem-pf-v6-badge>
                <cem-pf-v6-button slot="utilities"
                              id="search-clear"
                              variant="plain"
                              aria-label="Clear search"
                              hidden
                              @click=${this.#onSearchClear}>
                  <svg width="1em"
                       height="1em"
                       viewBox="0 0 352 512"
                       fill="currentColor"
                       aria-hidden="true">
                    <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
                  </svg>
                </cem-pf-v6-button>
              </cem-pf-v6-text-input-group>
            </cem-pf-v6-toolbar-item>
          </cem-pf-v6-toolbar-group>
          <cem-pf-v6-toolbar-group variant="action-group">
            <cem-pf-v6-toolbar-item>
              <cem-pf-v6-button id="expand-all"
                            variant="tertiary"
                            size="small"
                            aria-label="Expand all tree items"
                            @click=${this.#onExpandAll}>
                Expand all
              </cem-pf-v6-button>
            </cem-pf-v6-toolbar-item>
            <cem-pf-v6-toolbar-item>
              <cem-pf-v6-button id="collapse-all"
                            variant="tertiary"
                            size="small"
                            aria-label="Collapse all tree items"
                            @click=${this.#onCollapseAll}>
                Collapse all
              </cem-pf-v6-button>
            </cem-pf-v6-toolbar-item>
          </cem-pf-v6-toolbar-group>
        </cem-pf-v6-toolbar>
        <cem-pf-v6-drawer id="drawer">
          <div id="tree-wrapper">
            <cem-virtual-tree id="virtual-tree"
                              @item-select=${this.#onItemSelect}></cem-virtual-tree>
          </div>
          <cem-detail-panel id="detail-panel"
                            slot="panel-body"></cem-detail-panel>
        </cem-pf-v6-drawer>
      </div>
    `;
  }
  get #drawer() {
    return this.shadowRoot?.getElementById("drawer");
  }
  get #virtualTree() {
    return this.shadowRoot?.getElementById("virtual-tree");
  }
  get #detailPanel() {
    return this.shadowRoot?.getElementById("detail-panel");
  }
  get #searchInput() {
    return this.shadowRoot?.getElementById("search");
  }
  get #searchCount() {
    return this.shadowRoot?.getElementById("search-count");
  }
  get #searchClear() {
    return this.shadowRoot?.getElementById("search-clear");
  }
  async #onItemSelect(e6) {
    const item = e6.item;
    if (!item) return;
    const manifest = await CemVirtualTree.loadManifest();
    if (this.#detailPanel && manifest) {
      await this.#detailPanel.renderItem(item, manifest);
      if (this.#drawer) {
        this.#drawer.expanded = true;
      }
    }
  }
  #onSearchInput() {
    const value = this.#searchInput?.value || "";
    if (this.#searchClear) {
      this.#searchClear.hidden = !value;
    }
    if (this.#searchDebounceTimer != null) {
      clearTimeout(this.#searchDebounceTimer);
    }
    this.#searchDebounceTimer = setTimeout(() => {
      this.#handleSearch(value);
    }, 300);
  }
  #onSearchClear() {
    if (this.#searchInput) {
      this.#searchInput.value = "";
    }
    if (this.#searchClear) {
      this.#searchClear.hidden = true;
    }
    if (this.#searchCount) {
      this.#searchCount.hidden = true;
    }
    this.#handleSearch("");
  }
  #handleSearch(query) {
    this.#virtualTree?.search(query);
    if (this.#searchCount) {
      this.#searchCount.hidden = !query;
    }
  }
  #onExpandAll() {
    this.#virtualTree?.expandAll();
  }
  #onCollapseAll() {
    this.#virtualTree?.collapseAll();
  }
};
_init17 = __decoratorStart(_a17);
CemManifestBrowser = __decorateElement(_init17, 0, "CemManifestBrowser", _CemManifestBrowser_decorators, CemManifestBrowser);
__runInitializers(_init17, 1, CemManifestBrowser);

// lit-css:elements/cem-pf-v6-alert/cem-pf-v6-alert.css
var s21 = new CSSStyleSheet();
s21.replaceSync(JSON.parse('":host {\\n  display: grid;\\n  position: relative;\\n  grid-template-areas: var(--cem-pf-v6-c-alert--GridTemplateAreas);\\n  grid-template-columns: var(--cem-pf-v6-c-alert--GridTemplateColumns);\\n  padding-block-start: var(--cem-pf-v6-c-alert--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-alert--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-alert--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-alert--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-alert--FontSize);\\n  background-color: var(--cem-pf-v6-c-alert--BackgroundColor);\\n  border: var(--cem-pf-v6-c-alert--BorderWidth) solid var(--cem-pf-v6-c-alert--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-alert--BorderRadius);\\n  box-shadow: var(--cem-pf-v6-c-alert--BoxShadow);\\n\\n  /* CSS variables from cem-pf-v6-c-alert */\\n  --cem-pf-v6-c-alert--BoxShadow: var(--pf-t--global--box-shadow--lg);\\n  --cem-pf-v6-c-alert--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --cem-pf-v6-c-alert--GridTemplateColumns: max-content 1fr max-content;\\n  --cem-pf-v6-c-alert--GridTemplateAreas:\\n    \\"icon title action\\"\\n    \\". description description\\"\\n    \\". actiongroup actiongroup\\";\\n  --cem-pf-v6-c-alert--BorderWidth: var(--pf-t--global--border--width--box--status--default);\\n  --cem-pf-v6-c-alert--BorderColor: transparent;\\n  --cem-pf-v6-c-alert--BorderRadius: var(--pf-t--global--border--radius--medium);\\n  --cem-pf-v6-c-alert--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-alert--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-alert--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-alert--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-alert--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-alert__toggle--MarginBlockStart: calc(-1 * var(--pf-t--global--spacer--control--vertical--default));\\n  --cem-pf-v6-c-alert__toggle--MarginBlockEnd: calc(-1 * var(--pf-t--global--spacer--control--vertical--default));\\n  --cem-pf-v6-c-alert__toggle--MarginInlineStart: calc(-1 * var(--pf-t--global--spacer--action--horizontal--plain--default));\\n  --cem-pf-v6-c-alert__toggle--MarginInlineEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-alert__toggle-icon--Rotate: 0;\\n  --cem-pf-v6-c-alert__toggle-icon--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-alert__toggle-icon--TransitionDuration: var(--pf-t--global--motion--duration--icon--long);\\n  --cem-pf-v6-c-alert__icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-alert__icon--MarginBlockStart: 0.25rem;\\n  --cem-pf-v6-c-alert__icon--MarginInlineEnd: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-alert__icon--FontSize: var(--pf-t--global--icon--size--md);\\n  --cem-pf-v6-c-alert__title--FontWeight: var(--pf-t--global--font--weight--body--bold);\\n  --cem-pf-v6-c-alert__title--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-alert__title--max-lines: 1;\\n  --cem-pf-v6-c-alert__action--MarginBlockStart: calc(var(--pf-t--global--spacer--control--vertical--default) * -1);\\n  --cem-pf-v6-c-alert__action--MarginBlockEnd: calc(var(--pf-t--global--spacer--control--vertical--default) * -1);\\n  --cem-pf-v6-c-alert__action--TranslateY: 0.125rem;\\n  --cem-pf-v6-c-alert__action--MarginInlineEnd: calc(var(--pf-t--global--spacer--sm) * -1);\\n  --cem-pf-v6-c-alert__description--PaddingBlockStart: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-alert__action-group--PaddingBlockStart-base: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-alert__action-group--PaddingBlockStart: var(--cem-pf-v6-c-alert__action-group--PaddingBlockStart-base);\\n  --cem-pf-v6-c-alert__description--action-group--PaddingBlockStart-base: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-alert__description--action-group--PaddingBlockStart: var(--cem-pf-v6-c-alert__description--action-group--PaddingBlockStart-base);\\n  --cem-pf-v6-c-alert__action-group__c-button--not-last-child--MarginInlineEnd: var(--pf-t--global--spacer--gap--action-to-action--default);\\n  --cem-pf-v6-c-alert--m-custom--BorderColor: var(--pf-t--global--border--color--status--custom--default);\\n  --cem-pf-v6-c-alert--m-custom__icon--Color: var(--pf-t--global--icon--color--status--custom--default);\\n  --cem-pf-v6-c-alert--m-custom__title--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-alert--m-success--BorderColor: var(--pf-t--global--border--color--status--success--default);\\n  --cem-pf-v6-c-alert--m-success__icon--Color: var(--pf-t--global--icon--color--status--success--default);\\n  --cem-pf-v6-c-alert--m-success__title--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-alert--m-danger--BorderColor: var(--pf-t--global--border--color--status--danger--default);\\n  --cem-pf-v6-c-alert--m-danger__icon--Color: var(--pf-t--global--icon--color--status--danger--default);\\n  --cem-pf-v6-c-alert--m-danger__title--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-alert--m-warning--BorderColor: var(--pf-t--global--border--color--status--warning--default);\\n  --cem-pf-v6-c-alert--m-warning__icon--Color: var(--pf-t--global--icon--color--status--warning--default);\\n  --cem-pf-v6-c-alert--m-warning__title--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-alert--m-info--BorderColor: var(--pf-t--global--border--color--status--info--default);\\n  --cem-pf-v6-c-alert--m-info__icon--Color: var(--pf-t--global--icon--color--status--info--default);\\n  --cem-pf-v6-c-alert--m-info__title--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-alert--m-inline--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-alert--m-inline--BoxShadow: none;\\n  --cem-pf-v6-c-alert--m-inline--m-plain--BorderWidth: 0;\\n  --cem-pf-v6-c-alert--m-inline--m-plain--BackgroundColor: transparent;\\n  --cem-pf-v6-c-alert--m-inline--m-plain--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-alert--m-inline--m-plain--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-alert--m-inline--m-plain--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-alert--m-inline--m-plain--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-alert--m-expandable--GridTemplateColumns: auto max-content 1fr max-content;\\n  --cem-pf-v6-c-alert--m-expandable--GridTemplateAreas:\\n    \\"toggle icon title action\\"\\n    \\". . description description\\"\\n    \\". . actiongroup actiongroup\\";\\n  --cem-pf-v6-c-alert--m-expandable__description--action-group--PaddingBlockStart: var(--cem-pf-v6-c-alert__action-group--PaddingBlockStart-base);\\n  --cem-pf-v6-c-alert--m-expanded__toggle-icon--Rotate: 90deg;\\n  --cem-pf-v6-c-alert--m-expanded__description--action-group--PaddingBlockStart: var(--cem-pf-v6-c-alert__description--action-group--PaddingBlockStart-base);\\n}\\n\\n/* Variant modifiers */\\n:host([variant=\\"custom\\"]) {\\n  --cem-pf-v6-c-alert--BorderColor: var(--cem-pf-v6-c-alert--m-custom--BorderColor);\\n  --cem-pf-v6-c-alert__icon--Color: var(--cem-pf-v6-c-alert--m-custom__icon--Color);\\n  --cem-pf-v6-c-alert__title--Color: var(--cem-pf-v6-c-alert--m-custom__title--Color);\\n}\\n\\n:host([variant=\\"success\\"]) {\\n  --cem-pf-v6-c-alert--BorderColor: var(--cem-pf-v6-c-alert--m-success--BorderColor);\\n  --cem-pf-v6-c-alert__icon--Color: var(--cem-pf-v6-c-alert--m-success__icon--Color);\\n  --cem-pf-v6-c-alert__title--Color: var(--cem-pf-v6-c-alert--m-success__title--Color);\\n}\\n\\n:host([variant=\\"danger\\"]) {\\n  --cem-pf-v6-c-alert--BorderColor: var(--cem-pf-v6-c-alert--m-danger--BorderColor);\\n  --cem-pf-v6-c-alert__icon--Color: var(--cem-pf-v6-c-alert--m-danger__icon--Color);\\n  --cem-pf-v6-c-alert__title--Color: var(--cem-pf-v6-c-alert--m-danger__title--Color);\\n}\\n\\n:host([variant=\\"warning\\"]) {\\n  --cem-pf-v6-c-alert--BorderColor: var(--cem-pf-v6-c-alert--m-warning--BorderColor);\\n  --cem-pf-v6-c-alert__icon--Color: var(--cem-pf-v6-c-alert--m-warning__icon--Color);\\n  --cem-pf-v6-c-alert__title--Color: var(--cem-pf-v6-c-alert--m-warning__title--Color);\\n}\\n\\n:host([variant=\\"info\\"]) {\\n  --cem-pf-v6-c-alert--BorderColor: var(--cem-pf-v6-c-alert--m-info--BorderColor);\\n  --cem-pf-v6-c-alert__icon--Color: var(--cem-pf-v6-c-alert--m-info__icon--Color);\\n  --cem-pf-v6-c-alert__title--Color: var(--cem-pf-v6-c-alert--m-info__title--Color);\\n}\\n\\n/* Inline modifier */\\n:host([inline]) {\\n  --cem-pf-v6-c-alert--BoxShadow: var(--cem-pf-v6-c-alert--m-inline--BoxShadow);\\n  --cem-pf-v6-c-alert--BackgroundColor: var(--cem-pf-v6-c-alert--m-inline--BackgroundColor);\\n}\\n\\n/* Plain modifier */\\n:host([plain]) {\\n  --cem-pf-v6-c-alert--BorderWidth: var(--cem-pf-v6-c-alert--m-inline--m-plain--BorderWidth);\\n  --cem-pf-v6-c-alert--BackgroundColor: var(--cem-pf-v6-c-alert--m-inline--m-plain--BackgroundColor);\\n  --cem-pf-v6-c-alert--PaddingBlockStart: var(--cem-pf-v6-c-alert--m-inline--m-plain--PaddingBlockStart);\\n  --cem-pf-v6-c-alert--PaddingInlineEnd: var(--cem-pf-v6-c-alert--m-inline--m-plain--PaddingInlineEnd);\\n  --cem-pf-v6-c-alert--PaddingBlockEnd: var(--cem-pf-v6-c-alert--m-inline--m-plain--PaddingBlockEnd);\\n  --cem-pf-v6-c-alert--PaddingInlineStart: var(--cem-pf-v6-c-alert--m-inline--m-plain--PaddingInlineStart);\\n}\\n\\n/* Expandable modifier */\\n:host([expandable]) {\\n  --cem-pf-v6-c-alert--GridTemplateColumns: var(--cem-pf-v6-c-alert--m-expandable--GridTemplateColumns);\\n  --cem-pf-v6-c-alert--GridTemplateAreas: var(--cem-pf-v6-c-alert--m-expandable--GridTemplateAreas);\\n  --cem-pf-v6-c-alert__description--action-group--PaddingBlockStart: var(--cem-pf-v6-c-alert--m-expandable__description--action-group--PaddingBlockStart);\\n}\\n\\n/* Expanded modifier */\\n:host([expanded]) {\\n  --cem-pf-v6-c-alert__toggle-icon--Rotate: var(--cem-pf-v6-c-alert--m-expanded__toggle-icon--Rotate);\\n  --cem-pf-v6-c-alert__description--action-group--PaddingBlockStart: var(--cem-pf-v6-c-alert--m-expanded__description--action-group--PaddingBlockStart);\\n}\\n\\n/* Internal element styles */\\n#toggle {\\n  grid-area: toggle;\\n  margin-block-start: var(--cem-pf-v6-c-alert__toggle--MarginBlockStart);\\n  margin-block-end: var(--cem-pf-v6-c-alert__toggle--MarginBlockEnd);\\n  margin-inline-start: var(--cem-pf-v6-c-alert__toggle--MarginInlineStart);\\n  margin-inline-end: var(--cem-pf-v6-c-alert__toggle--MarginInlineEnd);\\n}\\n\\n#toggle-icon {\\n  display: inline-block;\\n  transition: transform var(--cem-pf-v6-c-alert__toggle-icon--TransitionDuration) var(--cem-pf-v6-c-alert__toggle-icon--TransitionTimingFunction);\\n  transform: rotate(var(--cem-pf-v6-c-alert__toggle-icon--Rotate));\\n}\\n\\n:host(:dir(rtl)) #toggle-icon {\\n  scale: -1 1;\\n}\\n\\n#icon-container {\\n  grid-area: icon;\\n  margin-inline-end: var(--cem-pf-v6-c-alert__icon--MarginInlineEnd);\\n  font-size: var(--cem-pf-v6-c-alert__icon--FontSize);\\n  color: var(--cem-pf-v6-c-alert__icon--Color);\\n}\\n\\n#title {\\n  display: var(--_has-title, block);\\n  grid-area: title;\\n  font-weight: var(--cem-pf-v6-c-alert__title--FontWeight);\\n  color: var(--cem-pf-v6-c-alert__title--Color);\\n  word-break: break-word;\\n}\\n\\n:host([truncate]) #title {\\n  display: var(--_has-title, -webkit-box);\\n  -webkit-box-orient: vertical;\\n  -webkit-line-clamp: var(--cem-pf-v6-c-alert__title--max-lines);\\n  overflow: hidden;\\n}\\n\\n#description {\\n  display: var(--_has-description, block);\\n  grid-area: description;\\n  padding-block-start: var(--cem-pf-v6-c-alert__description--PaddingBlockStart);\\n  word-break: break-word;\\n}\\n\\n#description + #action-group {\\n  --cem-pf-v6-c-alert__action-group--PaddingBlockStart: var(--cem-pf-v6-c-alert__description--action-group--PaddingBlockStart);\\n}\\n\\n#action {\\n  display: var(--_has-action, block);\\n  grid-area: action;\\n  margin-block-start: var(--cem-pf-v6-c-alert__action--MarginBlockStart);\\n  margin-block-end: var(--cem-pf-v6-c-alert__action--MarginBlockEnd);\\n  margin-inline-end: var(--cem-pf-v6-c-alert__action--MarginInlineEnd);\\n}\\n\\n#action-group {\\n  display: var(--_has-actions, block);\\n  grid-area: actiongroup;\\n  padding-block-start: var(--cem-pf-v6-c-alert__action-group--PaddingBlockStart);\\n}\\n\\n#action-group \\u003e :is(cem-pf-v6-button, .cem-pf-v6-c-button):not(:last-child) {\\n  margin-inline-end: var(--cem-pf-v6-c-alert__action-group__c-button--not-last-child--MarginInlineEnd);\\n}\\n\\n#sr-variant {\\n  position: absolute;\\n  overflow: hidden;\\n  clip: rect(0, 0, 0, 0);\\n  white-space: nowrap;\\n  border: 0;\\n}\\n\\n:is(h1, h2, h3, h4, h5, h6),\\n::slotted(:is(h1, h2, h3, h4, h5, h6)) {\\n  margin: 0 !important;\\n}\\n"'));
var cem_pf_v6_alert_default = s21;

// elements/cem-pf-v6-alert/cem-pf-v6-alert.ts
var PfAlertExpandEvent = class extends Event {
  constructor() {
    super("expand", { bubbles: true });
  }
};
var PfAlertCollapseEvent = class extends Event {
  constructor() {
    super("collapse", { bubbles: true });
  }
};
var PfAlertCloseEvent = class extends Event {
  constructor() {
    super("close", { bubbles: true });
  }
};
var ICON_PATHS = {
  success: "M504 256c0 136.967-111.033 248-248 248S8 392.967 8 256 119.033 8 256 8s248 111.033 248 248zM227.314 387.314l184-184c6.248-6.248 6.248-16.379 0-22.627l-22.627-22.627c-6.248-6.249-16.379-6.249-22.628 0L216 308.118l-70.059-70.059c-6.248-6.248-16.379-6.248-22.628 0l-22.627 22.627c-6.248 6.248-6.248 16.379 0 22.627l104 104c6.249 6.249 16.379 6.249 22.628.001z",
  danger: "M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 110c23.196 0 42 18.804 42 42s-18.804 42-42 42-42-18.804-42-42 18.804-42 42-42zm56 254c0 6.627-5.373 12-12 12h-88c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h12v-64h-12c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h64c6.627 0 12 5.373 12 12v100h12c6.627 0 12 5.373 12 12v24z",
  warning: "M569.517 440.013C587.975 472.007 564.806 512 527.94 512H48.054c-36.937 0-59.999-40.055-41.577-71.987L246.423 23.985c18.467-32.009 64.72-31.951 83.154 0l239.94 416.028zM288 354c-25.405 0-46 20.595-46 46s20.595 46 46 46 46-20.595 46-46-20.595-46-46-46zm-43.673-165.346l7.418 136c.347 6.364 5.609 11.346 11.982 11.346h48.546c6.373 0 11.635-4.982 11.982-11.346l7.418-136c.375-6.874-5.098-12.654-11.982-12.654h-63.383c-6.884 0-12.356 5.78-11.981 12.654z",
  info: "M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 110c23.196 0 42 18.804 42 42s-18.804 42-42 42-42-18.804-42-42 18.804-42 42-42zm56 254c0 6.627-5.373 12-12 12h-88c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h12v-64h-12c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h64c6.627 0 12 5.373 12 12v100h12c6.627 0 12 5.373 12 12v24z",
  custom: "M464 32H48C21.49 32 0 53.49 0 80v352c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V80c0-26.51-21.49-48-48-48zm-6 400H54a6 6 0 0 1-6-6V86a6 6 0 0 1 6-6h404a6 6 0 0 1 6 6v340a6 6 0 0 1-6 6zm-42-92v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm0-96v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm0-96v24c0 6.627-5.373 12-12 12H204c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h200c6.627 0 12 5.373 12 12zm-252 12c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36zm0 96c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36zm0 96c0 19.882-16.118 36-36 36s-36-16.118-36-36 16.118-36 36-36 36 16.118 36 36z"
};
var VARIANT_LABELS = {
  success: "Success alert:",
  danger: "Danger alert:",
  warning: "Warning alert:",
  info: "Info alert:",
  custom: "Custom alert:"
};
var _liveRegion_dec, _dismissable_dec, _truncate_dec2, _expanded_dec6, _expandable_dec2, _plain_dec2, _inline_dec, _variant_dec6, _a18, _PfV6Alert_decorators, _init18, _variant6, _inline, _plain2, _expandable2, _expanded6, _truncate2, _dismissable, _liveRegion, _PfV6Alert_instances, iconPath_get, variantLabel_get, closeButtonLabel_get, onClose_fn;
_PfV6Alert_decorators = [t3("cem-pf-v6-alert")];
var PfV6Alert = class extends (_a18 = i3, _variant_dec6 = [n4({ reflect: true })], _inline_dec = [n4({ type: Boolean, reflect: true })], _plain_dec2 = [n4({ type: Boolean, reflect: true })], _expandable_dec2 = [n4({ type: Boolean, reflect: true })], _expanded_dec6 = [n4({ type: Boolean, reflect: true })], _truncate_dec2 = [n4({ type: Number, reflect: true })], _dismissable_dec = [n4({ type: Boolean, reflect: true })], _liveRegion_dec = [n4({ type: Boolean, reflect: true, attribute: "live-region" })], _a18) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Alert_instances);
    __privateAdd(this, _variant6, __runInitializers(_init18, 8, this, "custom")), __runInitializers(_init18, 11, this);
    __privateAdd(this, _inline, __runInitializers(_init18, 12, this, false)), __runInitializers(_init18, 15, this);
    __privateAdd(this, _plain2, __runInitializers(_init18, 16, this, false)), __runInitializers(_init18, 19, this);
    __privateAdd(this, _expandable2, __runInitializers(_init18, 20, this, false)), __runInitializers(_init18, 23, this);
    __privateAdd(this, _expanded6, __runInitializers(_init18, 24, this, false)), __runInitializers(_init18, 27, this);
    __privateAdd(this, _truncate2, __runInitializers(_init18, 28, this)), __runInitializers(_init18, 31, this);
    __privateAdd(this, _dismissable, __runInitializers(_init18, 32, this, false)), __runInitializers(_init18, 35, this);
    __privateAdd(this, _liveRegion, __runInitializers(_init18, 36, this, false)), __runInitializers(_init18, 39, this);
  }
  render() {
    return T`
      ${this.expandable ? T`
        <div id="toggle-container">
          <cem-pf-v6-button id="toggle"
                        variant="plain"
                        aria-expanded="${this.expanded}"
                        aria-label="Toggle alert"
                        @click=${this.toggle}>
            <svg id="toggle-icon"
                 aria-hidden="true"
                 fill="currentColor"
                 height="1em"
                 width="1em"
                 viewBox="0 0 320 512">
              <path d="M285.476 272.971L91.132 467.314c-9.373 9.373-24.569 9.373-33.941 0l-22.667-22.667c-9.357-9.357-9.375-24.522-.04-33.901L188.505 256 34.484 101.255c-9.335-9.379-9.317-24.544.04-33.901l22.667-22.667c9.373-9.373 24.569-9.373 33.941 0L285.475 239.03c9.373 9.372 9.373 24.568.001 33.941z"></path>
            </svg>
          </cem-pf-v6-button>
        </div>
      ` : A}

      <div id="icon-container">
        <slot name="icon">
          <svg viewBox="0 0 ${this.variant === "warning" ? "576" : "512"} 512"
               fill="currentColor"
               aria-hidden="true"
               role="img"
               width="1em"
               height="1em">
            <path d="${__privateGet(this, _PfV6Alert_instances, iconPath_get)}"></path>
          </svg>
        </slot>
      </div>

      <h4 id="title"
          style="${this.truncate ? `--cem-pf-v6-c-alert__title--max-lines: ${this.truncate}` : ""}">
        <span id="sr-variant">${__privateGet(this, _PfV6Alert_instances, variantLabel_get)}</span>
        <slot></slot>
      </h4>

      <div id="action">
        ${this.dismissable ? T`
          <cem-pf-v6-button id="close"
                        variant="plain"
                        aria-label="${__privateGet(this, _PfV6Alert_instances, closeButtonLabel_get)}"
                        @click=${__privateMethod(this, _PfV6Alert_instances, onClose_fn)}>
            <svg viewBox="0 0 352 512"
                 fill="currentColor"
                 aria-hidden="true"
                 role="img"
                 width="1em"
                 height="1em">
              <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
            </svg>
          </cem-pf-v6-button>
        ` : T`
          <slot name="action"></slot>
        `}
      </div>

      ${!this.expandable || this.expanded ? T`
        <div id="description">
          <slot name="description"></slot>
        </div>
      ` : A}

      <div id="action-group">
        <slot name="actions"></slot>
      </div>
    `;
  }
  /** Toggle expanded state (for expandable alerts) */
  toggle() {
    if (!this.expandable) return;
    this.expanded = !this.expanded;
    this.dispatchEvent(this.expanded ? new PfAlertExpandEvent() : new PfAlertCollapseEvent());
  }
  /** Expand the alert (for expandable alerts) */
  expand() {
    if (!this.expandable || this.expanded) return;
    this.expanded = true;
    this.dispatchEvent(new PfAlertExpandEvent());
  }
  /** Collapse the alert (for expandable alerts) */
  collapse() {
    if (!this.expandable || !this.expanded) return;
    this.expanded = false;
    this.dispatchEvent(new PfAlertCollapseEvent());
  }
};
_init18 = __decoratorStart(_a18);
_variant6 = new WeakMap();
_inline = new WeakMap();
_plain2 = new WeakMap();
_expandable2 = new WeakMap();
_expanded6 = new WeakMap();
_truncate2 = new WeakMap();
_dismissable = new WeakMap();
_liveRegion = new WeakMap();
_PfV6Alert_instances = new WeakSet();
iconPath_get = function() {
  return ICON_PATHS[this.variant] ?? ICON_PATHS.custom;
};
variantLabel_get = function() {
  return VARIANT_LABELS[this.variant] ?? VARIANT_LABELS.custom;
};
closeButtonLabel_get = function() {
  const titleText = this.textContent?.trim() ?? "";
  return titleText ? `Close ${__privateGet(this, _PfV6Alert_instances, variantLabel_get)} ${titleText}` : `Close ${__privateGet(this, _PfV6Alert_instances, variantLabel_get).replace(":", "")}`;
};
onClose_fn = function() {
  if (this.dispatchEvent(new PfAlertCloseEvent())) {
    this.remove();
  }
};
__decorateElement(_init18, 4, "variant", _variant_dec6, PfV6Alert, _variant6);
__decorateElement(_init18, 4, "inline", _inline_dec, PfV6Alert, _inline);
__decorateElement(_init18, 4, "plain", _plain_dec2, PfV6Alert, _plain2);
__decorateElement(_init18, 4, "expandable", _expandable_dec2, PfV6Alert, _expandable2);
__decorateElement(_init18, 4, "expanded", _expanded_dec6, PfV6Alert, _expanded6);
__decorateElement(_init18, 4, "truncate", _truncate_dec2, PfV6Alert, _truncate2);
__decorateElement(_init18, 4, "dismissable", _dismissable_dec, PfV6Alert, _dismissable);
__decorateElement(_init18, 4, "liveRegion", _liveRegion_dec, PfV6Alert, _liveRegion);
PfV6Alert = __decorateElement(_init18, 0, "PfV6Alert", _PfV6Alert_decorators, PfV6Alert);
__publicField(PfV6Alert, "styles", cem_pf_v6_alert_default);
__runInitializers(_init18, 1, PfV6Alert);

// lit-css:elements/cem-pf-v6-alert-group/cem-pf-v6-alert-group.css
var s22 = new CSSStyleSheet();
s22.replaceSync(JSON.parse('":host {\\n  display: block;\\n\\n  --cem-pf-v6-c-alert-group__item--MarginBlockStart: var(--pf-t--global--spacer--gap--group--vertical);\\n  --cem-pf-v6-c-alert-group--m-toast--InsetBlockStart: var(--pf-t--global--spacer--2xl);\\n  --cem-pf-v6-c-alert-group--m-toast--InsetInlineEnd: var(--pf-t--global--spacer--xl);\\n  --cem-pf-v6-c-alert-group--m-toast--MaxWidth: 37.5rem;\\n  --cem-pf-v6-c-alert-group--m-toast--ZIndex: var(--pf-t--global--z-index--2xl);\\n}\\n\\n:host([toast]) {\\n  position: fixed;\\n  inset-block-start: var(--cem-pf-v6-c-alert-group--m-toast--InsetBlockStart);\\n  inset-inline-end: var(--cem-pf-v6-c-alert-group--m-toast--InsetInlineEnd);\\n  z-index: var(--cem-pf-v6-c-alert-group--m-toast--ZIndex);\\n  width: calc(100% - var(--cem-pf-v6-c-alert-group--m-toast--InsetInlineEnd) * 2);\\n  max-width: var(--cem-pf-v6-c-alert-group--m-toast--MaxWidth);\\n}\\n\\n:host ::slotted(*:not(:first-child)),\\n:host ::slotted(.cem-pf-v6-c-alert-group__item:not(:first-child)) {\\n  margin-block-start: var(--cem-pf-v6-c-alert-group__item--MarginBlockStart);\\n}\\n"'));
var cem_pf_v6_alert_group_default = s22;

// elements/cem-pf-v6-alert-group/cem-pf-v6-alert-group.ts
var _liveRegion_dec2, _toast_dec, _a19, _PfAlertGroup_decorators, _init19, _toast, _liveRegion2, _onClose, _PfAlertGroup_instances, removeAlert_fn;
_PfAlertGroup_decorators = [t3("cem-pf-v6-alert-group")];
var PfAlertGroup = class extends (_a19 = i3, _toast_dec = [n4({ type: Boolean, reflect: true })], _liveRegion_dec2 = [n4({ type: Boolean, reflect: true, attribute: "live-region" })], _a19) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfAlertGroup_instances);
    __privateAdd(this, _toast, __runInitializers(_init19, 8, this, false)), __runInitializers(_init19, 11, this);
    __privateAdd(this, _liveRegion2, __runInitializers(_init19, 12, this, false)), __runInitializers(_init19, 15, this);
    __privateAdd(this, _onClose, (e6) => {
      const alert = e6.target;
      if (alert?.tagName === "CEM-PF-V6-ALERT") {
        __privateMethod(this, _PfAlertGroup_instances, removeAlert_fn).call(this, alert);
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("close", __privateGet(this, _onClose));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("close", __privateGet(this, _onClose));
  }
  firstUpdated() {
    if (this.liveRegion) {
      this.setAttribute("aria-live", "polite");
      this.setAttribute("aria-atomic", "false");
    }
  }
  render() {
    return T`<slot></slot>`;
  }
  /**
   * Add an alert to the group
   */
  addAlert(alert) {
    if (this.toast) {
      alert.classList.add("pf-m-incoming");
      this.appendChild(alert);
      requestAnimationFrame(() => {
        alert.classList.remove("pf-m-incoming");
      });
    } else {
      this.appendChild(alert);
    }
  }
};
_init19 = __decoratorStart(_a19);
_toast = new WeakMap();
_liveRegion2 = new WeakMap();
_onClose = new WeakMap();
_PfAlertGroup_instances = new WeakSet();
removeAlert_fn = function(alert) {
  if (!alert) return;
  if (this.toast) {
    alert.classList.add("pf-m-outgoing");
    alert.addEventListener("transitionend", () => {
      alert.remove();
    }, { once: true });
  } else {
    alert.remove();
  }
};
__decorateElement(_init19, 4, "toast", _toast_dec, PfAlertGroup, _toast);
__decorateElement(_init19, 4, "liveRegion", _liveRegion_dec2, PfAlertGroup, _liveRegion2);
PfAlertGroup = __decorateElement(_init19, 0, "PfAlertGroup", _PfAlertGroup_decorators, PfAlertGroup);
__publicField(PfAlertGroup, "styles", cem_pf_v6_alert_group_default);
__runInitializers(_init19, 1, PfAlertGroup);

// lit-css:elements/cem-pf-v6-card/cem-pf-v6-card.css
var s23 = new CSSStyleSheet();
s23.replaceSync(JSON.parse('":host {\\n\\n  display: flex;\\n  flex-direction: column;\\n  overflow: hidden;\\n\\n  --cem-pf-v6-c-card--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-card--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-card--BorderStyle: solid;\\n  --cem-pf-v6-c-card--BorderWidth: var(--pf-t--global--border--width--box--default);\\n  --cem-pf-v6-c-card--BorderRadius: var(--pf-t--global--border--radius--medium);\\n  --cem-pf-v6-c-card--first-child--PaddingBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-card--child--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-card--child--PaddingBlockEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-card--child--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-card__title--not--last-child--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-card__title-text--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-card__title-text--FontFamily: var(--pf-t--global--font--family--heading);\\n  --cem-pf-v6-c-card__title-text--FontSize: var(--pf-t--global--font--size--heading--xs);\\n  --cem-pf-v6-c-card__title-text--FontWeight: var(--pf-t--global--font--weight--heading--default);\\n  --cem-pf-v6-c-card__title-text--LineHeight: var(--pf-t--global--font--line-height--heading);\\n  --cem-pf-v6-c-card__body--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-card__footer--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-card__footer--Color: var(--pf-t--global--text--color--subtle);\\n\\n  background-color: var(--cem-pf-v6-c-card--BackgroundColor);\\n  border: var(--cem-pf-v6-c-card--BorderWidth) var(--cem-pf-v6-c-card--BorderStyle) var(--cem-pf-v6-c-card--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-card--BorderRadius);\\n}\\n\\nheader,\\n#body,\\nfooter {\\n  padding-inline-end: var(--cem-pf-v6-c-card--child--PaddingInlineEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-card--child--PaddingInlineStart);\\n  padding-block-end: var(--cem-pf-v6-c-card--child--PaddingBlockEnd);\\n}\\n\\nheader:first-child,\\n#body:first-child {\\n  padding-block-start: var(--cem-pf-v6-c-card--first-child--PaddingBlockStart);\\n}\\n\\n#title {\\n  display: flex;\\n  align-items: center;\\n  gap: var(--pf-t--global--spacer--sm);\\n  flex-wrap: wrap;\\n}\\n\\n#body {\\n  font-size: var(--cem-pf-v6-c-card__body--FontSize);\\n  flex: 1 1 auto;\\n}\\n\\nfooter {\\n  display: flex;\\n  flex-wrap: wrap;\\n  gap: var(--pf-t--global--spacer--gap--action-to-action--default);\\n  font-size: var(--cem-pf-v6-c-card__footer--FontSize);\\n  color: var(--cem-pf-v6-c-card__footer--Color);\\n}\\n\\n#title ::slotted(h1),\\n#title ::slotted(h2),\\n#title ::slotted(h3),\\n#title ::slotted(h4),\\n#title ::slotted(h5),\\n#title ::slotted(h6) {\\n  margin: 0;\\n  color: var(--cem-pf-v6-c-card__title-text--Color);\\n  font-family: var(--cem-pf-v6-c-card__title-text--FontFamily);\\n  font-size: var(--cem-pf-v6-c-card__title-text--FontSize);\\n  font-weight: var(--cem-pf-v6-c-card__title-text--FontWeight);\\n  line-height: var(--cem-pf-v6-c-card__title-text--LineHeight);\\n  padding-block-start: var(--cem-pf-v6-c-card--first-child--PaddingBlockStart);\\n}\\n\\n#title ::slotted(cem-pf-v6-popover) {\\n  margin-inline-start: auto;\\n  margin-block-start: var(--cem-pf-v6-c-card--first-child--PaddingBlockStart);\\n}\\n\\n#title ::slotted(.element-summary) {\\n  display: inline;\\n  font-family: var(--pf-t--global--font--family--body);\\n  font-size: var(--pf-t--global--font--size--body--default);\\n  font-weight: var(--pf-t--global--font--weight--body--default);\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--regular);\\n}\\n\\nheader:not([hidden]):not(:last-child) {\\n  padding-block-end: var(--cem-pf-v6-c-card__title--not--last-child--PaddingBlockEnd);\\n}\\n\\n:host([compact]) {\\n  --cem-pf-v6-c-card--first-child--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-card--child--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-card--child--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-card--child--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-card__title--not--last-child--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n}\\n\\n:host([variant=\\"secondary\\"]) {\\n  background-color: var(--pf-t--global--background--color--secondary--default);\\n  border-color: var(--pf-t--global--border--color--high-contrast);\\n}\\n\\n:host([variant=\\"plain\\"]) {\\n  background-color: transparent;\\n  border-color: transparent;\\n}\\n\\n:host([full-height]) {\\n  height: 100%;\\n}\\n"'));
var cem_pf_v6_card_default = s23;

// elements/cem-pf-v6-card/cem-pf-v6-card.ts
var _fullHeight_dec2, _compact_dec2, _variant_dec7, _a20, _PfV6Card_decorators, _init20, _variant7, _compact2, _fullHeight2;
_PfV6Card_decorators = [t3("cem-pf-v6-card")];
var PfV6Card = class extends (_a20 = i3, _variant_dec7 = [n4({ reflect: true })], _compact_dec2 = [n4({ type: Boolean, reflect: true })], _fullHeight_dec2 = [n4({ type: Boolean, reflect: true, attribute: "full-height" })], _a20) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _variant7, __runInitializers(_init20, 8, this)), __runInitializers(_init20, 11, this);
    __privateAdd(this, _compact2, __runInitializers(_init20, 12, this, false)), __runInitializers(_init20, 15, this);
    __privateAdd(this, _fullHeight2, __runInitializers(_init20, 16, this, false)), __runInitializers(_init20, 19, this);
  }
  render() {
    return T`
      <header id="header">
        <slot name="header"></slot>
        <div id="title">
          <slot name="title"></slot>
        </div>
      </header>
      <div id="body">
        <slot></slot>
      </div>
      <footer id="footer">
        <slot name="footer"></slot>
      </footer>
    `;
  }
};
_init20 = __decoratorStart(_a20);
_variant7 = new WeakMap();
_compact2 = new WeakMap();
_fullHeight2 = new WeakMap();
__decorateElement(_init20, 4, "variant", _variant_dec7, PfV6Card, _variant7);
__decorateElement(_init20, 4, "compact", _compact_dec2, PfV6Card, _compact2);
__decorateElement(_init20, 4, "fullHeight", _fullHeight_dec2, PfV6Card, _fullHeight2);
PfV6Card = __decorateElement(_init20, 0, "PfV6Card", _PfV6Card_decorators, PfV6Card);
__publicField(PfV6Card, "styles", cem_pf_v6_card_default);
__runInitializers(_init20, 1, PfV6Card);

// lit-css:elements/cem-pf-v6-dropdown/cem-pf-v6-dropdown.css
var s24 = new CSSStyleSheet();
s24.replaceSync(JSON.parse('":host {\\n  display: inline-block;\\n  position: relative;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n#toggle {\\n  --cem-pf-v6-c-menu-toggle--Gap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-menu-toggle--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-menu-toggle--PaddingInlineEnd: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-menu-toggle--PaddingBlockEnd: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-menu-toggle--PaddingInlineStart: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-menu-toggle--MinWidth: calc(var(--cem-pf-v6-c-menu-toggle--FontSize) * var(--cem-pf-v6-c-menu-toggle--LineHeight) + var(--cem-pf-v6-c-menu-toggle--PaddingBlockStart) + var(--cem-pf-v6-c-menu-toggle--PaddingBlockEnd));\\n  --cem-pf-v6-c-menu-toggle--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-menu-toggle--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu-toggle--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-menu-toggle--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-menu-toggle--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-menu-toggle--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-menu-toggle--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-menu-toggle--border--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-menu-toggle--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-menu-toggle--TransitionDuration: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-menu-toggle--TransitionProperty: color, background-color, border-width, border-color;\\n  --cem-pf-v6-c-menu-toggle--hover--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu-toggle--hover--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-menu-toggle--hover--BorderWidth: var(--pf-t--global--border--width--control--hover);\\n  --cem-pf-v6-c-menu-toggle--hover--BorderColor: var(--pf-t--global--border--color--hover);\\n  --cem-pf-v6-c-menu-toggle--expanded--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu-toggle--expanded--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-menu-toggle--expanded--BorderWidth: var(--pf-t--global--border--width--control--clicked);\\n  --cem-pf-v6-c-menu-toggle--expanded--BorderColor: var(--pf-t--global--border--color--clicked);\\n  --cem-pf-v6-c-menu-toggle__toggle-icon--Color: var(--pf-t--global--icon--color--regular);\\n\\n  --cem-pf-v6-c-button--Color: var(--cem-pf-v6-c-menu-toggle--Color);\\n  --cem-pf-v6-c-button--BackgroundColor: var(--cem-pf-v6-c-menu-toggle--BackgroundColor);\\n  --cem-pf-v6-c-button--BorderRadius: var(--cem-pf-v6-c-menu-toggle--BorderRadius);\\n  --cem-pf-v6-c-button--PaddingBlockStart: var(--cem-pf-v6-c-menu-toggle--PaddingBlockStart);\\n  --cem-pf-v6-c-button--PaddingBlockEnd: var(--cem-pf-v6-c-menu-toggle--PaddingBlockEnd);\\n  --cem-pf-v6-c-button--PaddingInlineStart: var(--cem-pf-v6-c-menu-toggle--PaddingInlineStart);\\n  --cem-pf-v6-c-button--PaddingInlineEnd: var(--cem-pf-v6-c-menu-toggle--PaddingInlineEnd);\\n  --cem-pf-v6-c-button--FontSize: var(--cem-pf-v6-c-menu-toggle--FontSize);\\n  --cem-pf-v6-c-button--LineHeight: var(--cem-pf-v6-c-menu-toggle--LineHeight);\\n  --cem-pf-v6-c-button--MinWidth: var(--cem-pf-v6-c-menu-toggle--MinWidth);\\n\\n  \\u0026::after {\\n    position: absolute;\\n    inset: 0;\\n    z-index: var(--cem-pf-v6-c-menu-toggle--border--ZIndex);\\n    pointer-events: none;\\n    content: \\"\\";\\n    border: var(--cem-pf-v6-c-menu-toggle--BorderWidth) solid var(--cem-pf-v6-c-menu-toggle--BorderColor);\\n    border-radius: var(--cem-pf-v6-c-menu-toggle--BorderRadius);\\n    transition: inherit;\\n  }\\n\\n  \\u0026:hover:not([disabled]) {\\n    --cem-pf-v6-c-menu-toggle--Color: var(--cem-pf-v6-c-menu-toggle--hover--Color);\\n    --cem-pf-v6-c-menu-toggle--BackgroundColor: var(--cem-pf-v6-c-menu-toggle--hover--BackgroundColor);\\n    --cem-pf-v6-c-menu-toggle--BorderWidth: var(--cem-pf-v6-c-menu-toggle--hover--BorderWidth);\\n    --cem-pf-v6-c-menu-toggle--BorderColor: var(--cem-pf-v6-c-menu-toggle--hover--BorderColor);\\n  }\\n}\\n\\n:host([expanded]) #toggle {\\n  --cem-pf-v6-c-menu-toggle--Color: var(--cem-pf-v6-c-menu-toggle--expanded--Color);\\n  --cem-pf-v6-c-menu-toggle--BackgroundColor: var(--cem-pf-v6-c-menu-toggle--expanded--BackgroundColor);\\n  --cem-pf-v6-c-menu-toggle--BorderWidth: var(--cem-pf-v6-c-menu-toggle--expanded--BorderWidth);\\n  --cem-pf-v6-c-menu-toggle--BorderColor: var(--cem-pf-v6-c-menu-toggle--expanded--BorderColor);\\n}\\n\\n#toggle svg[slot=\\"icon-end\\"] {\\n  width: 0.75em;\\n  height: 0.75em;\\n  fill: currentColor;\\n  color: var(--cem-pf-v6-c-menu-toggle__toggle-icon--Color);\\n  transition: transform 150ms ease-in-out;\\n}\\n\\n:host([expanded]) #toggle svg[slot=\\"icon-end\\"] {\\n  transform: rotate(180deg);\\n}\\n\\n#menu-container {\\n  position: absolute;\\n  top: 100%;\\n  left: 0;\\n  margin-top: var(--pf-t--global--spacer--xs);\\n  z-index: var(--pf-t--global--z-index--sm, 1000);\\n}\\n\\n:host([disabled]) {\\n  pointer-events: none;\\n  opacity: 0.6;\\n}\\n"'));
var cem_pf_v6_dropdown_default = s24;

// lit-css:elements/cem-pf-v6-menu/cem-pf-v6-menu.css
var s25 = new CSSStyleSheet();
s25.replaceSync(JSON.parse('":host {\\n  --cem-pf-v6-c-menu--RowGap: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu--Width: auto;\\n  --cem-pf-v6-c-menu--MinWidth: auto;\\n  --cem-pf-v6-c-menu--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --cem-pf-v6-c-menu--BoxShadow: var(--pf-t--global--box-shadow--md);\\n  --cem-pf-v6-c-menu--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-menu--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-menu--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-menu--OutlineOffset: calc(var(--pf-t--global--border--width--control--default) * -3);\\n  --cem-pf-v6-c-menu--ZIndex: var(--pf-t--global--z-index--sm);\\n  --cem-pf-v6-c-menu--TransitionDuration: 0s;\\n  --cem-pf-v6-c-menu--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n\\n  @media (prefers-reduced-motion: no-preference) {\\n    --cem-pf-v6-c-menu--TransitionDuration: var(--pf-t--global--motion--duration--fade--default);\\n  }\\n\\n  display: grid;\\n  row-gap: var(--cem-pf-v6-c-menu--RowGap);\\n  width: var(--cem-pf-v6-c-menu--Width);\\n  min-width: var(--cem-pf-v6-c-menu--MinWidth);\\n  padding-block-start: var(--cem-pf-v6-c-menu--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-menu--PaddingBlockEnd);\\n  overflow: hidden;\\n  color: var(--cem-pf-v6-c-menu--Color);\\n  background-color: var(--cem-pf-v6-c-menu--BackgroundColor);\\n  border: var(--cem-pf-v6-c-menu--BorderWidth) solid var(--cem-pf-v6-c-menu--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-menu--BorderRadius);\\n  box-shadow: var(--cem-pf-v6-c-menu--BoxShadow);\\n  transition-timing-function: var(--cem-pf-v6-c-menu--TransitionTimingFunction) !important;\\n  transition-duration: var(--cem-pf-v6-c-menu--TransitionDuration) !important;\\n}\\n\\nslot {\\n  display: contents;\\n}\\n"'));
var cem_pf_v6_menu_default = s25;

// elements/cem-pf-v6-menu/cem-pf-v6-menu.ts
var _label_dec2, _a21, _PfV6Menu_decorators, _internals3, _init21, _label2, _PfV6Menu_instances, onSlotChange_fn2, _onKeydown2, getMenuItems_fn, initializeTabindex_fn, focusItem_fn;
_PfV6Menu_decorators = [t3("cem-pf-v6-menu")];
var PfV6Menu = class extends (_a21 = i3, _label_dec2 = [n4()], _a21) {
  constructor() {
    super();
    __privateAdd(this, _PfV6Menu_instances);
    __privateAdd(this, _internals3, this.attachInternals());
    __privateAdd(this, _label2, __runInitializers(_init21, 8, this, "")), __runInitializers(_init21, 11, this);
    __privateAdd(this, _onKeydown2, (event) => {
      const { key, target } = event;
      if (target.tagName !== "CEM-PF-V6-MENU-ITEM") return;
      const items = __privateMethod(this, _PfV6Menu_instances, getMenuItems_fn).call(this);
      const currentIndex = items.indexOf(target);
      switch (key) {
        case "ArrowDown":
          event.preventDefault();
          __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[currentIndex < items.length - 1 ? currentIndex + 1 : 0]);
          break;
        case "ArrowUp":
          event.preventDefault();
          __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[currentIndex > 0 ? currentIndex - 1 : items.length - 1]);
          break;
        case "Home":
          event.preventDefault();
          if (items.length > 0) __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[0]);
          break;
        case "End":
          event.preventDefault();
          if (items.length > 0) __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[items.length - 1]);
          break;
      }
    });
    __privateGet(this, _internals3).role = "menu";
  }
  render() {
    return T`<slot @slotchange=${__privateMethod(this, _PfV6Menu_instances, onSlotChange_fn2)}></slot>`;
  }
  updated(changed) {
    if (changed.has("label")) {
      __privateGet(this, _internals3).ariaLabel = this.label || null;
    }
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("keydown", __privateGet(this, _onKeydown2));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("keydown", __privateGet(this, _onKeydown2));
  }
  /**
   * Focus the first menu item.
   * Called when dropdown opens.
   */
  focusFirstItem() {
    const items = __privateMethod(this, _PfV6Menu_instances, getMenuItems_fn).call(this);
    if (items.length > 0) {
      __privateMethod(this, _PfV6Menu_instances, focusItem_fn).call(this, items[0]);
    }
  }
};
_init21 = __decoratorStart(_a21);
_internals3 = new WeakMap();
_label2 = new WeakMap();
_PfV6Menu_instances = new WeakSet();
onSlotChange_fn2 = function() {
  requestAnimationFrame(() => {
    __privateMethod(this, _PfV6Menu_instances, initializeTabindex_fn).call(this);
  });
};
_onKeydown2 = new WeakMap();
getMenuItems_fn = function() {
  const slot = this.shadowRoot?.querySelector("slot");
  if (!slot) return [];
  return slot.assignedElements().filter(
    (el) => el.tagName === "CEM-PF-V6-MENU-ITEM" && !el.disabled
  );
};
initializeTabindex_fn = function() {
  const items = __privateMethod(this, _PfV6Menu_instances, getMenuItems_fn).call(this);
  if (items.length === 0) return;
  items.forEach((item, index) => {
    item.setAttribute("tabindex", index === 0 ? "0" : "-1");
  });
};
focusItem_fn = function(item) {
  if (!item) return;
  const items = __privateMethod(this, _PfV6Menu_instances, getMenuItems_fn).call(this);
  items.forEach((i6) => {
    i6.setAttribute("tabindex", i6 === item ? "0" : "-1");
  });
  item.focus();
};
__decorateElement(_init21, 4, "label", _label_dec2, PfV6Menu, _label2);
PfV6Menu = __decorateElement(_init21, 0, "PfV6Menu", _PfV6Menu_decorators, PfV6Menu);
__publicField(PfV6Menu, "styles", cem_pf_v6_menu_default);
__runInitializers(_init21, 1, PfV6Menu);

// lit-css:elements/cem-pf-v6-menu-item/cem-pf-v6-menu-item.css
var s26 = new CSSStyleSheet();
s26.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-menu__list-item--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu__list-item--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-menu__list-item--BorderWidth: var(--pf-t--global--border--width--action--plain--default);\\n  --cem-pf-v6-c-menu__list-item--hover--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --cem-pf-v6-c-menu__list-item--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-menu__list-item--TransitionDuration: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-menu__list-item--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-menu__list-item--TransitionProperty: background-color;\\n  --cem-pf-v6-c-menu__list-item--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --cem-pf-v6-c-menu__item--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu__item--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu__item--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-menu__item--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-menu__item--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-menu__item--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-menu__item--FontWeight: var(--pf-t--global--font--weight--body--default);\\n  --cem-pf-v6-c-menu__item--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-menu__item--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-menu__item--m-disabled--Color: var(--pf-t--global--text--color--disabled);\\n  --cem-pf-v6-c-menu__item-main--ColumnGap: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-menu__item-description--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --cem-pf-v6-c-menu__item-description--Color: var(--pf-t--global--text--color--subtle);\\n  --cem-pf-v6-c-menu--OutlineOffset: calc(var(--pf-t--global--border--width--control--default) * -3);\\n\\n  position: relative;\\n  display: flex;\\n  align-items: baseline;\\n  min-width: 0;\\n  padding-block-start: var(--cem-pf-v6-c-menu__item--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-menu__item--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-menu__item--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-menu__item--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-menu__item--FontSize);\\n  font-weight: var(--cem-pf-v6-c-menu__item--FontWeight);\\n  line-height: var(--cem-pf-v6-c-menu__item--LineHeight);\\n  color: var(--cem-pf-v6-c-menu__item--Color);\\n  text-align: start;\\n  text-decoration-line: none;\\n  background-color: transparent;\\n  border: 0;\\n  outline-offset: var(--cem-pf-v6-c-menu--OutlineOffset);\\n  cursor: pointer;\\n  user-select: none;\\n  transition-timing-function: var(--cem-pf-v6-c-menu__list-item--TransitionTimingFunction);\\n  transition-duration: var(--cem-pf-v6-c-menu__list-item--TransitionDuration);\\n  transition-property: var(--cem-pf-v6-c-menu__list-item--TransitionProperty);\\n  gap: var(--cem-pf-v6-c-menu__item-main--ColumnGap);\\n}\\n\\n:host::before {\\n  position: absolute;\\n  inset: 0;\\n  content: \\"\\";\\n  background-color: var(--cem-pf-v6-c-menu__list-item--BackgroundColor);\\n  border-block-start: var(--cem-pf-v6-c-menu__list-item--BorderWidth) solid var(--cem-pf-v6-c-menu__list-item--BorderColor);\\n  border-block-end: var(--cem-pf-v6-c-menu__list-item--BorderWidth) solid var(--cem-pf-v6-c-menu__list-item--BorderColor);\\n  transition: inherit;\\n}\\n\\n:host(:hover:not([disabled])),\\n:host(:focus-visible) {\\n  --cem-pf-v6-c-menu__list-item--BackgroundColor: var(--cem-pf-v6-c-menu__list-item--hover--BackgroundColor);\\n  --cem-pf-v6-c-menu__list-item--BorderWidth: var(--cem-pf-v6-c-menu__list-item--hover--BorderWidth);\\n}\\n\\n:host(:focus-visible) {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: -2px;\\n}\\n\\n:host([disabled]) {\\n  --cem-pf-v6-c-menu__item--Color: var(--cem-pf-v6-c-menu__item--m-disabled--Color);\\n  --cem-pf-v6-c-menu__list-item--BackgroundColor: transparent;\\n  --cem-pf-v6-c-menu__list-item--hover--BackgroundColor: transparent;\\n  cursor: not-allowed;\\n  pointer-events: none;\\n}\\n\\n#check,\\n#text,\\n#description {\\n  position: relative;\\n  z-index: 1;\\n}\\n\\n#check {\\n  display: flex;\\n  align-items: center;\\n  flex-shrink: 0;\\n\\n  \\u0026 .cem-pf-v6-c-check {\\n    --cem-pf-v6-c-check__input--TranslateY: none;\\n  }\\n}\\n\\n#text {\\n  overflow: hidden;\\n  text-overflow: ellipsis;\\n  white-space: nowrap;\\n  flex-grow: 1;\\n}\\n\\n#description {\\n  display: block;\\n  font-size: var(--cem-pf-v6-c-menu__item-description--FontSize);\\n  color: var(--cem-pf-v6-c-menu__item-description--Color);\\n  word-break: break-word;\\n  margin-block-start: 0.25rem;\\n}\\n\\nslot {\\n  display: inline;\\n}\\n"'));
var cem_pf_v6_menu_item_default = s26;

// elements/cem-pf-v6-menu-item/cem-pf-v6-menu-item.ts
var PfMenuItemSelectEvent = class extends Event {
  value;
  checked;
  constructor(value, checked) {
    super("select", { bubbles: true });
    this.value = value;
    this.checked = checked;
  }
};
var _description_dec, _value_dec2, _variant_dec8, _checked_dec, _disabled_dec5, _a22, _PfV6MenuItem_decorators, _internals4, _init22, _disabled5, _checked, _variant8, _value2, _description, _handleClick, _handleKeydown2;
_PfV6MenuItem_decorators = [t3("cem-pf-v6-menu-item")];
var PfV6MenuItem = class extends (_a22 = i3, _disabled_dec5 = [n4({ type: Boolean, reflect: true })], _checked_dec = [n4({ type: Boolean, reflect: true })], _variant_dec8 = [n4({ reflect: true })], _value_dec2 = [n4({ reflect: true })], _description_dec = [n4({ reflect: true })], _a22) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _internals4, this.attachInternals());
    __privateAdd(this, _disabled5, __runInitializers(_init22, 8, this, false)), __runInitializers(_init22, 11, this);
    __privateAdd(this, _checked, __runInitializers(_init22, 12, this, false)), __runInitializers(_init22, 15, this);
    __privateAdd(this, _variant8, __runInitializers(_init22, 16, this, "default")), __runInitializers(_init22, 19, this);
    __privateAdd(this, _value2, __runInitializers(_init22, 20, this, "")), __runInitializers(_init22, 23, this);
    __privateAdd(this, _description, __runInitializers(_init22, 24, this)), __runInitializers(_init22, 27, this);
    __privateAdd(this, _handleClick, (event) => {
      if (this.disabled) {
        event.preventDefault();
        event.stopImmediatePropagation();
        return;
      }
      if (this.variant === "checkbox") {
        this.checked = !this.checked;
      }
      this.dispatchEvent(new PfMenuItemSelectEvent(this.value, this.checked));
    });
    __privateAdd(this, _handleKeydown2, (event) => {
      if (this.disabled) return;
      if (event.key === " " || event.key === "Enter") {
        event.preventDefault();
        __privateGet(this, _handleClick).call(this, event);
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("click", __privateGet(this, _handleClick));
    this.addEventListener("keydown", __privateGet(this, _handleKeydown2));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("click", __privateGet(this, _handleClick));
    this.removeEventListener("keydown", __privateGet(this, _handleKeydown2));
  }
  updated(changed) {
    if (changed.has("variant")) {
      __privateGet(this, _internals4).role = this.variant === "checkbox" ? "menuitemcheckbox" : "menuitem";
    }
    if (changed.has("checked") || changed.has("variant")) {
      __privateGet(this, _internals4).ariaChecked = this.variant === "checkbox" ? String(this.checked) : null;
    }
    if (changed.has("disabled")) {
      __privateGet(this, _internals4).ariaDisabled = this.disabled ? "true" : null;
      this.setAttribute("tabindex", "-1");
    }
  }
  render() {
    return T`
      ${this.variant === "checkbox" ? T`
        <span id="check"
              class="cem-pf-v6-c-check pf-m-standalone">
          <input id="input"
                 class="cem-pf-v6-c-check__input"
                 role="presentation"
                 tabindex="-1"
                 type="checkbox"
                 .checked=${this.checked}
                 ?disabled=${this.disabled}>
        </span>
      ` : A}
      <span id="text"><slot></slot></span>
      ${this.description ? T`
        <span id="description">${this.description}</span>
      ` : A}
    `;
  }
};
_init22 = __decoratorStart(_a22);
_internals4 = new WeakMap();
_disabled5 = new WeakMap();
_checked = new WeakMap();
_variant8 = new WeakMap();
_value2 = new WeakMap();
_description = new WeakMap();
_handleClick = new WeakMap();
_handleKeydown2 = new WeakMap();
__decorateElement(_init22, 4, "disabled", _disabled_dec5, PfV6MenuItem, _disabled5);
__decorateElement(_init22, 4, "checked", _checked_dec, PfV6MenuItem, _checked);
__decorateElement(_init22, 4, "variant", _variant_dec8, PfV6MenuItem, _variant8);
__decorateElement(_init22, 4, "value", _value_dec2, PfV6MenuItem, _value2);
__decorateElement(_init22, 4, "description", _description_dec, PfV6MenuItem, _description);
PfV6MenuItem = __decorateElement(_init22, 0, "PfV6MenuItem", _PfV6MenuItem_decorators, PfV6MenuItem);
__publicField(PfV6MenuItem, "styles", cem_pf_v6_menu_item_default);
__runInitializers(_init22, 1, PfV6MenuItem);

// elements/cem-pf-v6-dropdown/cem-pf-v6-dropdown.ts
var _label_dec3, _disabled_dec6, _expanded_dec7, _a23, _PfV6Dropdown_decorators, _instances, _init23, _expanded7, _disabled6, _label3, _onKeydown3, _PfV6Dropdown_instances, onToggleClick_fn2;
_PfV6Dropdown_decorators = [t3("cem-pf-v6-dropdown")];
var _PfV6Dropdown = class _PfV6Dropdown extends (_a23 = i3, _expanded_dec7 = [n4({ type: Boolean, reflect: true })], _disabled_dec6 = [n4({ type: Boolean, reflect: true })], _label_dec3 = [n4()], _a23) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Dropdown_instances);
    __privateAdd(this, _expanded7, __runInitializers(_init23, 8, this, false)), __runInitializers(_init23, 11, this);
    __privateAdd(this, _disabled6, __runInitializers(_init23, 12, this, false)), __runInitializers(_init23, 15, this);
    __privateAdd(this, _label3, __runInitializers(_init23, 16, this, "")), __runInitializers(_init23, 19, this);
    __privateAdd(this, _onKeydown3, (event) => {
      if (event.key === "Escape" && this.expanded) {
        event.preventDefault();
        this.collapse();
        this.shadowRoot?.getElementById("toggle")?.focus();
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    __privateGet(_PfV6Dropdown, _instances).add(this);
    this.addEventListener("keydown", __privateGet(this, _onKeydown3));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(_PfV6Dropdown, _instances).delete(this);
    this.removeEventListener("keydown", __privateGet(this, _onKeydown3));
  }
  render() {
    return T`
      <cem-pf-v6-button id="toggle"
                     variant="tertiary"
                     aria-haspopup="true"
                     aria-expanded="${this.expanded}"
                     ?disabled=${this.disabled}
                     @click=${__privateMethod(this, _PfV6Dropdown_instances, onToggleClick_fn2)}>
        <slot name="toggle-text">Toggle</slot>
        <svg slot="icon-end"
             viewBox="0 0 320 512"
             aria-hidden="true">
          <path d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z"/>
        </svg>
      </cem-pf-v6-button>

      <div id="menu-container"
           ?hidden=${!this.expanded}>
        <cem-pf-v6-menu id="menu"
                     label=${this.label}>
          <slot></slot>
        </cem-pf-v6-menu>
      </div>
    `;
  }
  updated(changed) {
    if (changed.has("expanded")) {
      this.dispatchEvent(new Event(this.expanded ? "expand" : "collapse", { bubbles: true }));
      if (this.expanded) {
        requestAnimationFrame(() => {
          const menu = this.shadowRoot?.getElementById("menu");
          menu?.focusFirstItem();
        });
      }
    }
  }
  /** Toggle expanded state */
  toggle() {
    this.expanded = !this.expanded;
  }
  /** Expand the dropdown */
  expand() {
    this.expanded = true;
  }
  /** Collapse the dropdown */
  collapse() {
    this.expanded = false;
  }
};
_init23 = __decoratorStart(_a23);
_instances = new WeakMap();
_expanded7 = new WeakMap();
_disabled6 = new WeakMap();
_label3 = new WeakMap();
_onKeydown3 = new WeakMap();
_PfV6Dropdown_instances = new WeakSet();
onToggleClick_fn2 = function() {
  if (this.disabled) return;
  this.toggle();
};
__decorateElement(_init23, 4, "expanded", _expanded_dec7, _PfV6Dropdown, _expanded7);
__decorateElement(_init23, 4, "disabled", _disabled_dec6, _PfV6Dropdown, _disabled6);
__decorateElement(_init23, 4, "label", _label_dec3, _PfV6Dropdown, _label3);
_PfV6Dropdown = __decorateElement(_init23, 0, "PfV6Dropdown", _PfV6Dropdown_decorators, _PfV6Dropdown);
__privateAdd(_PfV6Dropdown, _instances, /* @__PURE__ */ new Set());
document?.addEventListener?.("click", (event) => {
  for (const instance of __privateGet(_PfV6Dropdown, _instances)) {
    if (instance.expanded && !event.composedPath().includes(instance)) {
      instance.collapse();
    }
  }
});
__publicField(_PfV6Dropdown, "styles", cem_pf_v6_dropdown_default);
__runInitializers(_init23, 1, _PfV6Dropdown);
var PfV6Dropdown = _PfV6Dropdown;

// lit-css:elements/cem-pf-v6-form/cem-pf-v6-form.css
var s27 = new CSSStyleSheet();
s27.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\n#form {\\n  display: block;\\n}\\n\\n:host([horizontal]) #form {\\n  --_form-column-gap: var(--pf-t--global--spacer--md);\\n  display: grid;\\n  grid-template-columns: minmax(0, auto) 1fr;\\n  column-gap: var(--_form-column-gap);\\n  row-gap: var(--pf-t--global--spacer--sm);\\n\\n  \\u0026 ::slotted(cem-pf-v6-form-field-group) {\\n    grid-column: 1 / -1;\\n  }\\n}\\n"'));
var cem_pf_v6_form_default = s27;

// elements/cem-pf-v6-form/cem-pf-v6-form.ts
var _horizontal_dec, _a24, _PfV6Form_decorators, _init24, _horizontal, _PfV6Form_instances, onSubmit_fn, formElement_get;
_PfV6Form_decorators = [t3("cem-pf-v6-form")];
var PfV6Form = class extends (_a24 = i3, _horizontal_dec = [n4({ type: Boolean, reflect: true })], _a24) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Form_instances);
    __privateAdd(this, _horizontal, __runInitializers(_init24, 8, this, false)), __runInitializers(_init24, 11, this);
  }
  render() {
    return T`
      <form id="form"
            @submit=${__privateMethod(this, _PfV6Form_instances, onSubmit_fn)}>
        <slot></slot>
      </form>
    `;
  }
  /** Submits the form programmatically */
  submit() {
    __privateGet(this, _PfV6Form_instances, formElement_get)?.submit();
  }
  /**
   * Requests form submission with validation
   * @param submitter - Optional submitter element
   */
  requestSubmit(submitter) {
    return __privateGet(this, _PfV6Form_instances, formElement_get)?.requestSubmit(submitter);
  }
  /** Resets the form to default values */
  reset() {
    __privateGet(this, _PfV6Form_instances, formElement_get)?.reset();
  }
  /** Checks form validity */
  checkValidity() {
    return __privateGet(this, _PfV6Form_instances, formElement_get)?.checkValidity() ?? true;
  }
  /** Reports form validity (shows validation messages) */
  reportValidity() {
    return __privateGet(this, _PfV6Form_instances, formElement_get)?.reportValidity() ?? true;
  }
};
_init24 = __decoratorStart(_a24);
_horizontal = new WeakMap();
_PfV6Form_instances = new WeakSet();
onSubmit_fn = function(e6) {
  this.dispatchEvent(new SubmitEvent("submit", {
    bubbles: true,
    cancelable: true,
    submitter: e6.submitter
  }));
};
formElement_get = function() {
  return this.shadowRoot?.getElementById("form");
};
__decorateElement(_init24, 4, "horizontal", _horizontal_dec, PfV6Form, _horizontal);
PfV6Form = __decorateElement(_init24, 0, "PfV6Form", _PfV6Form_decorators, PfV6Form);
__publicField(PfV6Form, "styles", cem_pf_v6_form_default);
__runInitializers(_init24, 1, PfV6Form);

// lit-css:elements/cem-pf-v6-form-field-group/cem-pf-v6-form-field-group.css
var s28 = new CSSStyleSheet();
s28.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-form__field-group--border-width-base: var(--pf-t--global--border--width--divider--default);\\n  --cem-pf-v6-c-form__field-group--BorderBlockStartWidth: var(--cem-pf-v6-c-form__field-group--border-width-base);\\n  --cem-pf-v6-c-form__field-group--BorderBlockStartColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-form__field-group--BorderBlockEndWidth: 0;\\n  --cem-pf-v6-c-form__field-group--BorderBlockEndColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-form__field-group--field-group--MarginBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-form__field-group--GridTemplateColumns--toggle: var(--pf-t--global--spacer--action--horizontal--plain--default);\\n  --cem-pf-v6-c-form__field-group-toggle--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-form__field-group-toggle--PaddingInlineEnd: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-form__field-group-toggle--field-group-header--GridColumn: 2 / 3;\\n  --cem-pf-v6-c-form__field-group-header--GridColumn: 1 / -1;\\n  --cem-pf-v6-c-form__field-group-header--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-form__field-group-header--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-form__field-group-header-description--MarginBlockStart: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-form__field-group-header-description--Color: var(--pf-t--global--text--color--subtle);\\n  --cem-pf-v6-c-form__field-group-header-actions--MarginBlockStart: calc(var(--pf-t--global--spacer--control--vertical--default) * -1);\\n  --cem-pf-v6-c-form__field-group-header-actions--MarginBlockEnd: calc(var(--pf-t--global--spacer--control--vertical--default) * -1);\\n  --cem-pf-v6-c-form__field-group-header-actions--MarginInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-form__field-group-body--GridColumn: 1 / -1;\\n  --cem-pf-v6-c-form__field-group-body--Gap: var(--pf-t--global--spacer--gap--group-to-group--vertical--default);\\n  --cem-pf-v6-c-form__field-group-body--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-form__field-group-body--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-form__field-group-body--TranslateY: -.5rem;\\n  --cem-pf-v6-c-form__field-group-body--TransitionDuration--collapse--fade: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-form__field-group-body--TransitionDuration--collapse--slide: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-form__field-group-body--TransitionDuration--expand--fade: var(--pf-t--global--motion--duration--fade--default);\\n  --cem-pf-v6-c-form__field-group-body--TransitionDuration--expand--slide: var(--pf-t--global--motion--duration--fade--default);\\n  --cem-pf-v6-c-form__field-group--m-expanded__field-group-body--TranslateY: 0;\\n  --cem-pf-v6-c-form__field-group-toggle-icon--MinWidth: 1em;\\n  --cem-pf-v6-c-form__field-group-toggle-icon--Rotate: 0;\\n  --cem-pf-v6-c-form__field-group--m-expanded__toggle-icon--Rotate: 90deg;\\n  --cem-pf-v6-c-form__field-group-toggle-icon--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-form__field-group-toggle-icon--TransitionDuration: var(--pf-t--global--motion--duration--icon--default);\\n  --cem-pf-v6-c-form__field-group-toggle-button--MarginBlockStart: calc(var(--pf-t--global--spacer--control--vertical--default) * -1);\\n  --cem-pf-v6-c-form__field-group-toggle-button--MarginBlockEnd: calc(var(--pf-t--global--spacer--control--vertical--default) * -1);\\n  --cem-pf-v6-c-form__field-group__field-group__field-group-body--GridColumn: 1 / -1;\\n  --cem-pf-v6-c-form__field-group__field-group__field-group-toggle--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-form__field-group__field-group__field-group-header--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-form__field-group__field-group__field-group-header--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-form__field-group__field-group__field-group-toggle--field-group-body--GridColumn: 2 / -1;\\n  --cem-pf-v6-c-form__field-group-body__field-group--last-child--MarginBlockEnd: calc(var(--pf-t--global--spacer--md) * -1);\\n\\n  display: grid;\\n  grid-template-columns: minmax(var(--cem-pf-v6-c-form__field-group--GridTemplateColumns--toggle), max-content) 1fr;\\n  border-block-start: var(--cem-pf-v6-c-form__field-group--BorderBlockStartWidth) solid var(--cem-pf-v6-c-form__field-group--BorderBlockStartColor);\\n  border-block-end: var(--cem-pf-v6-c-form__field-group--BorderBlockEndWidth) solid var(--cem-pf-v6-c-form__field-group--BorderBlockEndColor);\\n}\\n\\n:host(:last-child) {\\n  --cem-pf-v6-c-form__field-group--BorderBlockEndWidth: 0;\\n}\\n\\n:host + :host,\\n:host(:first-child) {\\n  --cem-pf-v6-c-form__field-group--BorderBlockStartWidth: 0;\\n}\\n\\n:host + :host {\\n  margin-block-start: var(--cem-pf-v6-c-form__field-group--field-group--MarginBlockStart);\\n}\\n\\n/* Nested field groups - parent sets these via CSS vars on slotted children */\\n:host {\\n  --cem-pf-v6-c-form__field-group-body--GridColumn: var(--_nested-body-grid-column, 1 / -1);\\n  --cem-pf-v6-c-form__field-group-toggle--PaddingBlockStart: var(--_nested-toggle-padding-block-start, var(--pf-t--global--spacer--md));\\n  --cem-pf-v6-c-form__field-group-header--PaddingBlockStart: var(--_nested-header-padding-block-start, var(--pf-t--global--spacer--md));\\n  --cem-pf-v6-c-form__field-group-header--PaddingBlockEnd: var(--_nested-header-padding-block-end, var(--pf-t--global--spacer--md));\\n  --cem-pf-v6-c-form__field-group-body--PaddingBlockStart: var(--_nested-body-padding-block-start, 0);\\n}\\n\\n:host([expanded]) {\\n  --cem-pf-v6-c-form__field-group-toggle-icon--Rotate: var(--cem-pf-v6-c-form__field-group--m-expanded__toggle-icon--Rotate);\\n}\\n\\n:host([expanded][expandable]) #body {\\n  max-height: 99999px;\\n  padding-block-start: var(--cem-pf-v6-c-form__field-group-body--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-form__field-group-body--PaddingBlockEnd);\\n  visibility: visible;\\n  opacity: 1;\\n  transition-delay: 0s;\\n  transition-duration: var(--cem-pf-v6-c-form__field-group-body--TransitionDuration--expand--fade), var(--cem-pf-v6-c-form__field-group-body--TransitionDuration--expand--slide), 0s, 0s, 0s, 0s;\\n  translate: 0 var(--cem-pf-v6-c-form__field-group--m-expanded__field-group-body--TranslateY);\\n}\\n\\n/* Toggle button */\\n#toggle-button {\\n  grid-row: 1/2;\\n  grid-column: 1/2;\\n  padding: auto auto;\\n  margin: auto auto;\\n}\\n\\n#toggle-button ~ #header {\\n  --cem-pf-v6-c-form__field-group-header--GridColumn: var(--cem-pf-v6-c-form__field-group-toggle--field-group-header--GridColumn);\\n}\\n\\n#toggle-button ~ #header ~ #body {\\n  padding-inline-start: calc(2 * var(--pf-t--global--spacer--action--horizontal--plain--default) + 1em * var(--pf-t--global--font--line-height--body));\\n}\\n\\n#toggle-icon {\\n  display: inline-flex;\\n  align-items: center;\\n  justify-content: center;\\n  min-width: var(--cem-pf-v6-c-form__field-group-toggle-icon--MinWidth);\\n  transition-timing-function: var(--cem-pf-v6-c-form__field-group-toggle-icon--TransitionTimingFunction);\\n  transition-duration: var(--cem-pf-v6-c-form__field-group-toggle-icon--TransitionDuration);\\n  transition-property: transform;\\n  transform: rotate(var(--cem-pf-v6-c-form__field-group-toggle-icon--Rotate));\\n  transform-origin: center;\\n}\\n\\n:host(:dir(rtl)) #toggle-icon {\\n  scale: -1 1;\\n}\\n\\n/* Header */\\n#header {\\n  display: flex;\\n  grid-row: 1/2;\\n  grid-column: var(--cem-pf-v6-c-form__field-group-header--GridColumn);\\n  align-items: flex-start;\\n  padding-block-start: var(--cem-pf-v6-c-form__field-group-header--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-form__field-group-header--PaddingBlockEnd);\\n}\\n\\n#header-main {\\n  display: flex;\\n  flex-direction: column;\\n  flex-grow: 1;\\n}\\n\\n#header-title {\\n  display: flex;\\n}\\n\\n#header-title-text {\\n  flex-grow: 1;\\n}\\n\\n#header-description {\\n  margin-block-start: var(--cem-pf-v6-c-form__field-group-header-description--MarginBlockStart);\\n  color: var(--cem-pf-v6-c-form__field-group-header-description--Color);\\n}\\n\\n#header-actions {\\n  margin-block-start: var(--cem-pf-v6-c-form__field-group-header-actions--MarginBlockStart);\\n  margin-block-end: var(--cem-pf-v6-c-form__field-group-header-actions--MarginBlockEnd);\\n  margin-inline-start: var(--cem-pf-v6-c-form__field-group-header-actions--MarginInlineStart);\\n  white-space: nowrap;\\n}\\n\\n/* Body */\\n#body {\\n  /* Set private vars for nested field groups */\\n  --_nested-body-grid-column: var(--cem-pf-v6-c-form__field-group__field-group__field-group-body--GridColumn);\\n  --_nested-toggle-padding-block-start: var(--cem-pf-v6-c-form__field-group__field-group__field-group-toggle--PaddingBlockStart);\\n  --_nested-header-padding-block-start: var(--cem-pf-v6-c-form__field-group__field-group__field-group-header--PaddingBlockStart);\\n  --_nested-header-padding-block-end: var(--cem-pf-v6-c-form__field-group__field-group__field-group-header--PaddingBlockEnd);\\n  --_nested-body-padding-block-start: 0;\\n\\n  display: grid;\\n  grid-column: var(--cem-pf-v6-c-form__field-group-body--GridColumn);\\n  padding-block-start: var(--cem-pf-v6-c-form__field-group-body--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-form__field-group-body--PaddingBlockEnd);\\n  grid-template-columns: inherit;\\n  column-gap: var(--_form-column-gap, var(--cem-pf-v6-c-form__field-group-body--Gap));\\n  row-gap: var(--pf-t--global--spacer--sm);\\n}\\n\\n:host([expandable]) #body {\\n  max-height: 0;\\n  padding-block-start: 0;\\n  padding-block-end: 0;\\n  visibility: hidden;\\n  opacity: 0;\\n  transition-delay: 0s, 0s, var(--cem-pf-v6-c-form__field-group-body--TransitionDuration--collapse--fade), var(--cem-pf-v6-c-form__field-group-body--TransitionDuration--collapse--fade), var(--cem-pf-v6-c-form__field-group-body--TransitionDuration--collapse--fade), var(--cem-pf-v6-c-form__field-group-body--TransitionDuration--collapse--fade);\\n  transition-duration: var(--cem-pf-v6-c-form__field-group-body--TransitionDuration--collapse--fade), var(--cem-pf-v6-c-form__field-group-body--TransitionDuration--collapse--slide), 0s, 0s, 0s, 0s;\\n  transition-property: opacity, translate, visibility, max-height, padding-block-start, padding-block-end;\\n  translate: 0 var(--cem-pf-v6-c-form__field-group-body--TranslateY);\\n}\\n\\n/* Accessibility: disable animations for users who prefer reduced motion */\\n@media (prefers-reduced-motion: reduce) {\\n  :host([expandable]) #body,\\n  :host([expanded][expandable]) #body {\\n    transition-duration: 0s;\\n    transition-delay: 0s;\\n  }\\n\\n  #toggle-icon {\\n    transition-duration: 0s;\\n  }\\n}\\n"'));
var cem_pf_v6_form_field_group_default = s28;

// elements/cem-pf-v6-form-field-group/cem-pf-v6-form-field-group.ts
var PfFormFieldGroupToggleEvent = class extends Event {
  expanded;
  constructor(expanded) {
    super("toggle", { bubbles: true });
    this.expanded = expanded;
  }
};
var _description_dec2, _toggleText_dec2, _expanded_dec8, _expandable_dec3, _a25, _PfV6FormFieldGroup_decorators, _init25, _expandable3, _expanded8, _toggleText2, _description2;
_PfV6FormFieldGroup_decorators = [t3("cem-pf-v6-form-field-group")];
var PfV6FormFieldGroup = class extends (_a25 = i3, _expandable_dec3 = [n4({ type: Boolean, reflect: true })], _expanded_dec8 = [n4({ type: Boolean, reflect: true })], _toggleText_dec2 = [n4({ reflect: true, attribute: "toggle-text" })], _description_dec2 = [n4({ reflect: true })], _a25) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _expandable3, __runInitializers(_init25, 8, this, false)), __runInitializers(_init25, 11, this);
    __privateAdd(this, _expanded8, __runInitializers(_init25, 12, this, false)), __runInitializers(_init25, 15, this);
    __privateAdd(this, _toggleText2, __runInitializers(_init25, 16, this)), __runInitializers(_init25, 19, this);
    __privateAdd(this, _description2, __runInitializers(_init25, 20, this)), __runInitializers(_init25, 23, this);
  }
  render() {
    const hasToggle = this.expandable || this.toggleText;
    const hasHeader = this.toggleText || this.description;
    return T`
      ${hasToggle ? T`
        <cem-pf-v6-button id="toggle-button"
                      part="toggle-button"
                      variant="plain"
                      aria-expanded=${String(this.expanded)}
                      aria-controls="body"
                      aria-label="Details"
                      @click=${this.toggle}>
          <svg id="toggle-icon"
               viewBox="0 0 256 512"
               fill="currentColor"
               role="presentation"
               width="1em"
               height="1em">
            <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
          </svg>
        </cem-pf-v6-button>
      ` : A}
      ${hasHeader ? T`
        <div id="header"
             part="header">
          <div id="header-main">
            ${this.toggleText ? T`
              <div id="header-title">
                <div id="header-title-text">
                  <span id="toggle-text-default">${this.toggleText}</span>
                </div>
              </div>
            ` : A}
            ${this.description ? T`
              <div id="header-description">${this.description}</div>
            ` : A}
          </div>
          <div id="header-actions">
            <slot name="header-actions"></slot>
          </div>
        </div>
      ` : A}
      <div id="body"
           part="body"
           ?inert=${this.expandable && !this.expanded}>
        <slot></slot>
      </div>
    `;
  }
  /** Toggles the expanded state */
  toggle() {
    if (this.expandable) {
      this.expanded = !this.expanded;
      this.dispatchEvent(new PfFormFieldGroupToggleEvent(this.expanded));
    }
  }
  /** Expands the field group */
  show() {
    if (this.expandable) {
      this.expanded = true;
      this.dispatchEvent(new PfFormFieldGroupToggleEvent(true));
    }
  }
  /** Collapses the field group */
  hide() {
    if (this.expandable) {
      this.expanded = false;
      this.dispatchEvent(new PfFormFieldGroupToggleEvent(false));
    }
  }
};
_init25 = __decoratorStart(_a25);
_expandable3 = new WeakMap();
_expanded8 = new WeakMap();
_toggleText2 = new WeakMap();
_description2 = new WeakMap();
__decorateElement(_init25, 4, "expandable", _expandable_dec3, PfV6FormFieldGroup, _expandable3);
__decorateElement(_init25, 4, "expanded", _expanded_dec8, PfV6FormFieldGroup, _expanded8);
__decorateElement(_init25, 4, "toggleText", _toggleText_dec2, PfV6FormFieldGroup, _toggleText2);
__decorateElement(_init25, 4, "description", _description_dec2, PfV6FormFieldGroup, _description2);
PfV6FormFieldGroup = __decorateElement(_init25, 0, "PfV6FormFieldGroup", _PfV6FormFieldGroup_decorators, PfV6FormFieldGroup);
__publicField(PfV6FormFieldGroup, "styles", cem_pf_v6_form_field_group_default);
__runInitializers(_init25, 1, PfV6FormFieldGroup);

// lit-css:elements/cem-pf-v6-form-group/cem-pf-v6-form-group.css
var s29 = new CSSStyleSheet();
s29.replaceSync(JSON.parse('":host {\\n  display: grid;\\n  grid-template-columns: subgrid;\\n  grid-column: 1 / -1;\\n  align-items: baseline;\\n}\\n\\n#label {\\n  grid-column: 1 / 2;\\n}\\n\\n#control {\\n  grid-column: 2 / 3;\\n  display: flex;\\n  gap: var(--pf-t--global--spacer--sm);\\n  align-items: start;\\n}\\n\\n#helper-text {\\n  grid-column: 1 / -1;\\n  font-size: var(--pf-t--global--font--size--sm);\\n  color: var(--pf-t--global--text--color--subtle);\\n  margin-top: var(--pf-t--global--spacer--xs);\\n\\n  \\u0026:empty {\\n    display: none;\\n  }\\n}\\n\\n::slotted([slot=control]:first-child) {\\n  flex: 1;\\n}\\n"'));
var cem_pf_v6_form_group_default = s29;

// elements/cem-pf-v6-form-group/cem-pf-v6-form-group.ts
var _PfV6FormGroup_decorators, _init26, _a26;
_PfV6FormGroup_decorators = [t3("cem-pf-v6-form-group")];
var PfV6FormGroup = class extends (_a26 = i3) {
  static styles = cem_pf_v6_form_group_default;
  render() {
    return T`
      <div id="label">
        <slot name="label"></slot>
      </div>
      <div id="control">
        <slot></slot>
      </div>
      <div id="helper-text">
        <slot name="helper"></slot>
      </div>
    `;
  }
};
_init26 = __decoratorStart(_a26);
PfV6FormGroup = __decorateElement(_init26, 0, "PfV6FormGroup", _PfV6FormGroup_decorators, PfV6FormGroup);
__runInitializers(_init26, 1, PfV6FormGroup);

// lit-css:elements/cem-pf-v6-form-label/cem-pf-v6-form-label.css
var s30 = new CSSStyleSheet();
s30.replaceSync(JSON.parse('":host {\\n}\\n\\nlabel {\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  font-size: var(--pf-t--global--font--size--body--default);\\n  color: var(--pf-t--global--text--color--regular);\\n  margin: 0;\\n  cursor: pointer;\\n}\\n\\nspan {\\n  display: inline;\\n}\\n"'));
var cem_pf_v6_form_label_default = s30;

// elements/cem-pf-v6-form-label/cem-pf-v6-form-label.ts
var _PfV6FormLabel_decorators, _init27, _a27;
_PfV6FormLabel_decorators = [t3("cem-pf-v6-form-label")];
var PfV6FormLabel = class extends (_a27 = i3) {
  static styles = cem_pf_v6_form_label_default;
  render() {
    return T`
      <label @click=${this.#handleClick}>
        <span><slot @slotchange=${this.#updateControlLabel}></slot></span>
      </label>
    `;
  }
  connectedCallback() {
    super.connectedCallback();
    requestAnimationFrame(() => this.#updateControlLabel());
  }
  #handleClick(e6) {
    e6.preventDefault();
    const control = this.#getControl();
    if (control && typeof control.focus === "function") {
      control.focus();
    }
  }
  /**
   * Updates the aria-label on the control's internal input element.
   *
   * Supports two approaches:
   * 1. Preferred: Control implements setAccessibleLabel(text) method (CemFormControl API)
   * 2. Fallback: Direct shadowRoot access for legacy controls
   *
   * Compatible controls should extend CemFormControl or implement setAccessibleLabel().
   */
  #updateControlLabel = () => {
    const control = this.#getControl();
    const labelText = this.textContent?.trim() || "";
    if (!control || !labelText) {
      return;
    }
    if (typeof control.setAccessibleLabel === "function") {
      control.setAccessibleLabel(labelText);
      return;
    }
    const shadowRoot = control.shadowRoot;
    if (!shadowRoot) {
      console.warn("Control has no shadow root and no setAccessibleLabel method");
      return;
    }
    const internalInput = shadowRoot.querySelector("input, select, textarea");
    if (internalInput) {
      internalInput.setAttribute("aria-label", labelText);
    } else {
      console.warn("Could not find internal input in control shadow DOM");
    }
  };
  #getControl() {
    const formGroup = this.closest("cem-pf-v6-form-group");
    if (!formGroup) {
      console.warn("cem-pf-v6-form-label must be inside cem-pf-v6-form-group");
      return null;
    }
    const control = formGroup.querySelector(":not([slot])");
    if (!control) {
      console.warn("No control found in cem-pf-v6-form-group");
      return null;
    }
    return control;
  }
};
_init27 = __decoratorStart(_a27);
PfV6FormLabel = __decorateElement(_init27, 0, "PfV6FormLabel", _PfV6FormLabel_decorators, PfV6FormLabel);
__runInitializers(_init27, 1, PfV6FormLabel);

// lit-css:elements/cem-pf-v6-masthead/cem-pf-v6-masthead.css
var s31 = new CSSStyleSheet();
s31.replaceSync(JSON.parse('":host {\\n  --cem-pf-v6-c-masthead--ColumnGap: var(--cem-pf-v6-c-masthead--m-display-stack--ColumnGap);\\n  --cem-pf-v6-c-masthead--GridTemplateColumns: var(--cem-pf-v6-c-masthead--m-display-stack--GridTemplateColumns);\\n  --cem-pf-v6-c-masthead__toggle--GridColumn: var(--cem-pf-v6-c-masthead--m-display-stack__toggle--GridColumn);\\n  --cem-pf-v6-c-masthead__brand--GridColumn: var(--cem-pf-v6-c-masthead--m-display-stack__brand--GridColumn);\\n  --cem-pf-v6-c-masthead__brand--Order: var(--cem-pf-v6-c-masthead--m-display-stack__brand--Order);\\n  --cem-pf-v6-c-masthead__brand--PaddingBlockEnd: var(--cem-pf-v6-c-masthead--m-display-stack__brand--PaddingBlockEnd);\\n  --cem-pf-v6-c-masthead__brand--BorderBlockEnd: var(--cem-pf-v6-c-masthead--m-display-stack__brand--BorderBlockEnd);\\n  --cem-pf-v6-c-masthead__main--GridColumn: unset;\\n  --cem-pf-v6-c-masthead__content--GridColumn: var(--cem-pf-v6-c-masthead--m-display-stack__content--GridColumn);\\n  --cem-pf-v6-c-masthead__content--Order: var(--cem-pf-v6-c-masthead--m-display-stack__content--Order);\\n  --cem-pf-v6-c-masthead__main--Display: var(--cem-pf-v6-c-masthead--m-display-stack__main--Display);\\n  --cem-pf-v6-c-masthead__main--ColumnGap: var(--cem-pf-v6-c-masthead--m-display-stack__main--ColumnGap);\\n\\n  position: relative;\\n  z-index: var(--cem-pf-v6-c-page--c-masthead--ZIndex);\\n  grid-area: header;\\n  display: grid;\\n  grid-template-columns: var(--cem-pf-v6-c-masthead--GridTemplateColumns);\\n  row-gap: var(--cem-pf-v6-c-masthead--RowGap);\\n  column-gap: var(--cem-pf-v6-c-masthead--ColumnGap);\\n  align-items: start;\\n  min-width: 0;\\n  padding-block-start: var(--cem-pf-v6-c-masthead--PaddingBlockStart, var(--cem-pf-v6-c-masthead--PaddingBlock));\\n  padding-block-end: var(--cem-pf-v6-c-masthead--PaddingBlockEnd, var(--cem-pf-v6-c-masthead--PaddingBlock));\\n  padding-inline-start: var(--cem-pf-v6-c-masthead--PaddingInlineStart, var(--cem-pf-v6-c-masthead--PaddingInline));\\n  padding-inline-end: var(--cem-pf-v6-c-masthead--PaddingInlineEnd, var(--cem-pf-v6-c-masthead--PaddingInline));\\n  background-color: var(--cem-pf-v6-c-masthead--BackgroundColor);\\n}\\n\\n#main {\\n  display: var(--cem-pf-v6-c-masthead__main--Display);\\n  grid-column: var(--cem-pf-v6-c-masthead__main--GridColumn);\\n  column-gap: var(--cem-pf-v6-c-masthead__main--ColumnGap);\\n  align-items: center;\\n}\\n\\n#toggle {\\n  grid-column: var(--cem-pf-v6-c-masthead__toggle--GridColumn);\\n  display: flex;\\n  align-items: center;\\n}\\n\\n#brand {\\n  position: relative;\\n  grid-column: var(--cem-pf-v6-c-masthead__brand--GridColumn);\\n  order: var(--cem-pf-v6-c-masthead__brand--Order);\\n  padding-block-end: var(--cem-pf-v6-c-masthead__brand--PaddingBlockEnd);\\n  border-block-end: var(--cem-pf-v6-c-masthead__brand--BorderBlockEnd);\\n}\\n\\n#content {\\n  grid-column: var(--cem-pf-v6-c-masthead__content--GridColumn);\\n  order: var(--cem-pf-v6-c-masthead__content--Order);\\n  column-gap: var(--cem-pf-v6-c-masthead__content--ColumnGap);\\n}\\n\\n::slotted(a[slot=\\"logo\\"]),\\n::slotted(img[slot=\\"logo\\"]),\\n::slotted(svg[slot=\\"logo\\"]) {\\n  position: relative;\\n  width: var(--cem-pf-v6-c-masthead__logo--Width);\\n  max-height: var(--cem-pf-v6-c-masthead__logo--MaxHeight);\\n}\\n\\n#toggle cem-pf-v6-button {\\n  font-size: var(--cem-pf-v6-c-masthead__toggle--Size);\\n}\\n\\n::slotted(cem-pf-v6-toolbar) {\\n  --cem-pf-v6-c-toolbar--Width: var(--cem-pf-v6-c-masthead--c-toolbar--Width);\\n  --cem-pf-v6-c-toolbar--PaddingBlockEnd: var(--cem-pf-v6-c-masthead--c-toolbar--PaddingBlockEnd);\\n  --cem-pf-v6-c-toolbar__content--MinWidth: var(--cem-pf-v6-c-masthead--c-toolbar__content--MinWidth);\\n  --cem-pf-v6-c-toolbar__content--PaddingInlineStart: var(--cem-pf-v6-c-masthead--c-toolbar__content--PaddingInlineStart);\\n  --cem-pf-v6-c-toolbar__content--PaddingInlineEnd: var(--cem-pf-v6-c-masthead--c-toolbar__content--PaddingInlineEnd);\\n  --cem-pf-v6-c-toolbar__content-section--FlexWrap: var(--cem-pf-v6-c-masthead--c-toolbar__content-section--FlexWrap);\\n}\\n"'));
var cem_pf_v6_masthead_default = s31;

// elements/cem-pf-v6-masthead/cem-pf-v6-masthead.ts
var SidebarToggleEvent = class extends Event {
  expanded;
  constructor(expanded) {
    super("sidebar-toggle", { bubbles: true });
    this.expanded = expanded;
  }
};
var _sidebarExpanded_dec, _a28, _PfV6Masthead_decorators, _internals5, _init28, _sidebarExpanded, _PfV6Masthead_instances, onToggle_fn;
_PfV6Masthead_decorators = [t3("cem-pf-v6-masthead")];
var PfV6Masthead = class extends (_a28 = i3, _sidebarExpanded_dec = [n4({ type: Boolean, reflect: true, attribute: "sidebar-expanded" })], _a28) {
  constructor() {
    super();
    __privateAdd(this, _PfV6Masthead_instances);
    __privateAdd(this, _internals5, this.attachInternals());
    __privateAdd(this, _sidebarExpanded, __runInitializers(_init28, 8, this, false)), __runInitializers(_init28, 11, this);
    __privateGet(this, _internals5).role = "banner";
  }
  render() {
    return T`
      <div id="main">
        <div id="toggle">
          <cem-pf-v6-button id="toggle-button"
                        variant="plain"
                        aria-label="Toggle global navigation"
                        aria-expanded=${String(this.sidebarExpanded)}
                        @click=${__privateMethod(this, _PfV6Masthead_instances, onToggle_fn)}>
            <svg id="hamburger"
                 role="presentation"
                 viewBox="0 0 10 10"
                 width="1em"
                 height="1em"
                 fill="none"
                 stroke="currentColor"
                 stroke-width="1.5"
                 stroke-linecap="round">
              <path d="M1,1 L9,1" />
              <path d="M1,5 L9,5" />
              <path d="M9,9 L1,9" />
            </svg>
          </cem-pf-v6-button>
        </div>
        <div id="brand">
          <slot name="logo"></slot>
        </div>
      </div>
      <div id="content">
        <slot name="toolbar"></slot>
      </div>
    `;
  }
};
_init28 = __decoratorStart(_a28);
_internals5 = new WeakMap();
_sidebarExpanded = new WeakMap();
_PfV6Masthead_instances = new WeakSet();
onToggle_fn = function() {
  this.dispatchEvent(new SidebarToggleEvent(!this.sidebarExpanded));
};
__decorateElement(_init28, 4, "sidebarExpanded", _sidebarExpanded_dec, PfV6Masthead, _sidebarExpanded);
PfV6Masthead = __decorateElement(_init28, 0, "PfV6Masthead", _PfV6Masthead_decorators, PfV6Masthead);
__publicField(PfV6Masthead, "styles", cem_pf_v6_masthead_default);
__runInitializers(_init28, 1, PfV6Masthead);

// lit-css:elements/cem-pf-v6-modal/cem-pf-v6-modal.css
var s32 = new CSSStyleSheet();
s32.replaceSync(JSON.parse('":host {\\n  display: block;\\n\\n  /* Modal box */\\n  --cem-pf-v6-c-modal-box--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --cem-pf-v6-c-modal-box--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-modal-box--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-modal-box--BorderRadius: var(--pf-t--global--border--radius--large);\\n  --cem-pf-v6-c-modal-box--BoxShadow: var(--pf-t--global--box-shadow--lg);\\n  --cem-pf-v6-c-modal-box--ZIndex: var(--pf-t--global--z-index--xl);\\n  --cem-pf-v6-c-modal-box--Width: 100%;\\n  --cem-pf-v6-c-modal-box--MaxWidth: calc(100% - var(--pf-t--global--spacer--xl));\\n  --cem-pf-v6-c-modal-box--m-sm--MaxWidth: 35rem;\\n  --cem-pf-v6-c-modal-box--m-md--Width: 52.5rem;\\n  --cem-pf-v6-c-modal-box--m-lg--MaxWidth: 70rem;\\n  --cem-pf-v6-c-modal-box--MaxHeight: calc(100% - var(--pf-t--global--spacer--2xl));\\n\\n  /* Align top position */\\n  --cem-pf-v6-c-modal-box--m-align-top--spacer: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-modal-box--m-align-top--InsetBlockStart: var(--cem-pf-v6-c-modal-box--m-align-top--spacer);\\n  --cem-pf-v6-c-modal-box--m-align-top--MaxHeight: calc(100% - min(var(--cem-pf-v6-c-modal-box--m-align-top--spacer), var(--pf-t--global--spacer--2xl)) - var(--cem-pf-v6-c-modal-box--m-align-top--spacer));\\n  --cem-pf-v6-c-modal-box--m-align-top--MaxWidth: calc(100% - min(var(--cem-pf-v6-c-modal-box--m-align-top--spacer) * 2, var(--pf-t--global--spacer--xl)));\\n\\n  /* Header */\\n  --cem-pf-v6-c-modal-box__header--PaddingBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__header--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-modal-box__header--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__header--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__header--Gap: var(--pf-t--global--spacer--md);\\n\\n  /* Header main */\\n  --cem-pf-v6-c-modal-box__header-main--Gap: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-modal-box__header-main--PaddingBlockStart: var(--pf-t--global--spacer--control--vertical--default);\\n\\n  /* Title */\\n  --cem-pf-v6-c-modal-box__title--LineHeight: var(--pf-t--global--font--line-height--heading);\\n  --cem-pf-v6-c-modal-box__title--FontFamily: var(--pf-t--global--font--family--heading);\\n  --cem-pf-v6-c-modal-box__title--FontWeight: var(--pf-t--global--font--weight--heading--default);\\n  --cem-pf-v6-c-modal-box__title--FontSize: var(--pf-t--global--font--size--heading--md);\\n\\n  /* Description */\\n  --cem-pf-v6-c-modal-box__description--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --cem-pf-v6-c-modal-box__description--Color: var(--pf-t--global--text--color--subtle);\\n\\n  /* Body */\\n  --cem-pf-v6-c-modal-box__body--MinHeight: calc(var(--pf-t--global--font--size--body--default) * var(--pf-t--global--font--line-height--body));\\n  --cem-pf-v6-c-modal-box__body--PaddingBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__body--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__body--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__body--PaddingBlockEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__header--body--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n\\n  /* Close button */\\n  --cem-pf-v6-c-modal-box__close--InsetBlockStart: var(--cem-pf-v6-c-modal-box__header--PaddingBlockStart);\\n  --cem-pf-v6-c-modal-box__close--InsetInlineEnd: var(--cem-pf-v6-c-modal-box__header--PaddingInlineEnd);\\n\\n  /* Footer */\\n  --cem-pf-v6-c-modal-box__footer--PaddingBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__footer--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__footer--PaddingBlockEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__footer--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-modal-box__footer--Gap: var(--pf-t--global--spacer--md);\\n}\\n\\n@media (min-width: 1200px) {\\n  :host {\\n    --cem-pf-v6-c-modal-box--m-align-top--spacer: var(--pf-t--global--spacer--xl);\\n  }\\n}\\n\\ndialog {\\n  position: fixed;\\n  inset: 0;\\n  z-index: var(--cem-pf-v6-c-modal-box--ZIndex);\\n  width: var(--cem-pf-v6-c-modal-box--Width);\\n  max-width: min(var(--cem-pf-v6-c-modal-box--MaxWidth), 100vw);\\n  max-height: min(var(--cem-pf-v6-c-modal-box--MaxHeight), 100vh);\\n  margin: auto;\\n  padding: 0;\\n  overflow: hidden;\\n  background-color: var(--cem-pf-v6-c-modal-box--BackgroundColor);\\n  border: var(--cem-pf-v6-c-modal-box--BorderWidth) solid var(--cem-pf-v6-c-modal-box--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-modal-box--BorderRadius);\\n  box-shadow: var(--cem-pf-v6-c-modal-box--BoxShadow);\\n\\n  \\u0026[open] {\\n    display: flex;\\n    flex-direction: column;\\n  }\\n\\n  \\u0026::backdrop {\\n    background-color: var(--cem-pf-v6-c-backdrop--BackgroundColor, var(--pf-t--global--background--color--backdrop--default));\\n  }\\n\\n  :host([variant=\\"small\\"]) \\u0026 {\\n    --cem-pf-v6-c-modal-box--Width: var(--cem-pf-v6-c-modal-box--m-sm--MaxWidth);\\n  }\\n\\n  :host([variant=\\"medium\\"]) \\u0026 {\\n    --cem-pf-v6-c-modal-box--Width: var(--cem-pf-v6-c-modal-box--m-md--Width);\\n  }\\n\\n  :host([variant=\\"large\\"]) \\u0026 {\\n    --cem-pf-v6-c-modal-box--Width: var(--cem-pf-v6-c-modal-box--m-lg--MaxWidth);\\n  }\\n\\n  :host([position=\\"top\\"]) \\u0026 {\\n    margin-block-start: var(--cem-pf-v6-c-modal-box--m-align-top--InsetBlockStart);\\n    margin-block-end: auto;\\n    --cem-pf-v6-c-modal-box--MaxWidth: var(--cem-pf-v6-c-modal-box--m-align-top--MaxWidth);\\n    --cem-pf-v6-c-modal-box--MaxHeight: var(--cem-pf-v6-c-modal-box--m-align-top--MaxHeight);\\n  }\\n}\\n\\n#header {\\n  display: flex;\\n  flex-direction: column;\\n  flex-shrink: 0;\\n  padding-block-start: var(--cem-pf-v6-c-modal-box__header--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-modal-box__header--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-modal-box__header--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-modal-box__header--PaddingInlineEnd);\\n  gap: var(--cem-pf-v6-c-modal-box__header--Gap);\\n}\\n\\n#header-main {\\n  display: flex;\\n  flex-direction: column;\\n  gap: var(--cem-pf-v6-c-modal-box__header-main--Gap);\\n  padding-block-start: var(--cem-pf-v6-c-modal-box__header-main--PaddingBlockStart);\\n  flex-grow: 1;\\n  min-width: 0;\\n}\\n\\n::slotted([slot=\\"header\\"]) {\\n  margin: 0 !important;\\n  overflow: hidden;\\n  text-overflow: ellipsis;\\n  white-space: nowrap;\\n  font-family: var(--cem-pf-v6-c-modal-box__title--FontFamily);\\n  font-size: var(--cem-pf-v6-c-modal-box__title--FontSize);\\n  font-weight: var(--cem-pf-v6-c-modal-box__title--FontWeight);\\n  line-height: var(--cem-pf-v6-c-modal-box__title--LineHeight);\\n}\\n\\n#description {\\n  font-size: var(--cem-pf-v6-c-modal-box__description--FontSize);\\n  color: var(--cem-pf-v6-c-modal-box__description--Color);\\n}\\n\\n#body {\\n  flex: 1 1 auto;\\n  min-height: var(--cem-pf-v6-c-modal-box__body--MinHeight);\\n  padding-block-start: var(--cem-pf-v6-c-modal-box__body--PaddingBlockStart);\\n  padding-inline-end: var(--cem-pf-v6-c-modal-box__body--PaddingInlineEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-modal-box__body--PaddingInlineStart);\\n  padding-block-end: var(--cem-pf-v6-c-modal-box__body--PaddingBlockEnd);\\n  overflow-x: hidden;\\n  overflow-y: auto;\\n  overscroll-behavior: contain;\\n  word-break: break-word;\\n  -webkit-overflow-scrolling: touch;\\n}\\n\\n#header:not([hidden]) ~ #body {\\n  --cem-pf-v6-c-modal-box__body--PaddingBlockStart: var(--cem-pf-v6-c-modal-box__header--body--PaddingBlockStart);\\n}\\n\\n#footer {\\n  display: flex;\\n  flex: 0 0 auto;\\n  align-items: center;\\n  gap: var(--cem-pf-v6-c-modal-box__footer--Gap);\\n  padding-block-start: var(--cem-pf-v6-c-modal-box__footer--PaddingBlockStart);\\n  padding-inline-end: var(--cem-pf-v6-c-modal-box__footer--PaddingInlineEnd);\\n  padding-block-end: var(--cem-pf-v6-c-modal-box__footer--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-modal-box__footer--PaddingInlineStart);\\n}\\n\\n#close {\\n  position: absolute;\\n  inset-block-start: var(--cem-pf-v6-c-modal-box__close--InsetBlockStart);\\n  inset-inline-end: var(--cem-pf-v6-c-modal-box__close--InsetInlineEnd);\\n\\n  \\u0026 svg {\\n    width: 1rem;\\n    height: 1rem;\\n    fill: currentColor;\\n  }\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n"'));
var cem_pf_v6_modal_default = s32;

// elements/cem-pf-v6-modal/cem-pf-v6-modal.ts
var PfModalOpenEvent = class extends Event {
  constructor() {
    super("open", { bubbles: true });
  }
};
var PfModalCloseEvent = class extends Event {
  returnValue;
  constructor(returnValue) {
    super("close", { bubbles: true });
    this.returnValue = returnValue;
  }
};
var PfModalCancelEvent = class extends Event {
  constructor() {
    super("cancel", { bubbles: true });
  }
};
var _open_dec2, _position_dec, _variant_dec9, _a29, _PfV6Modal_decorators, _dialog, _cancelling, _hasHeader, _hasDescription, _hasFooter, _init29, _variant9, _position, _open2, _PfV6Modal_instances, syncDialogState_fn, _onDialogClose, _onDialogCancel, _onCloseClick, _onHeaderSlotChange, _onDescriptionSlotChange, _onFooterSlotChange;
_PfV6Modal_decorators = [t3("cem-pf-v6-modal")];
var PfV6Modal = class extends (_a29 = i3, _variant_dec9 = [n4({ reflect: true })], _position_dec = [n4({ reflect: true })], _open_dec2 = [n4({ type: Boolean, reflect: true })], _a29) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Modal_instances);
    __privateAdd(this, _dialog, null);
    __privateAdd(this, _cancelling, false);
    __privateAdd(this, _hasHeader, false);
    __privateAdd(this, _hasDescription, false);
    __privateAdd(this, _hasFooter, false);
    /** @see https://developer.mozilla.org/en-US/docs/Web/API/HTMLDialogElement/returnValue */
    __publicField(this, "returnValue", "");
    __privateAdd(this, _variant9, __runInitializers(_init29, 8, this)), __runInitializers(_init29, 11, this);
    __privateAdd(this, _position, __runInitializers(_init29, 12, this)), __runInitializers(_init29, 15, this);
    __privateAdd(this, _open2, __runInitializers(_init29, 16, this, false)), __runInitializers(_init29, 19, this);
    __privateAdd(this, _onDialogClose, () => {
      if (this.open) {
        this.open = false;
      }
      if (__privateGet(this, _cancelling)) {
        this.dispatchEvent(new PfModalCancelEvent());
        __privateSet(this, _cancelling, false);
      } else {
        this.dispatchEvent(new PfModalCloseEvent(this.returnValue));
      }
    });
    __privateAdd(this, _onDialogCancel, (e6) => {
      e6.preventDefault();
      __privateSet(this, _cancelling, true);
      this.close();
    });
    __privateAdd(this, _onCloseClick, () => {
      __privateSet(this, _cancelling, true);
      this.close("cancel");
    });
    __privateAdd(this, _onHeaderSlotChange, (e6) => {
      const slot = e6.target;
      const had = __privateGet(this, _hasHeader);
      __privateSet(this, _hasHeader, slot.assignedNodes().length > 0);
      if (had !== __privateGet(this, _hasHeader)) this.requestUpdate();
    });
    __privateAdd(this, _onDescriptionSlotChange, (e6) => {
      const slot = e6.target;
      const had = __privateGet(this, _hasDescription);
      __privateSet(this, _hasDescription, slot.assignedNodes().length > 0);
      if (had !== __privateGet(this, _hasDescription)) this.requestUpdate();
    });
    __privateAdd(this, _onFooterSlotChange, (e6) => {
      const slot = e6.target;
      const had = __privateGet(this, _hasFooter);
      __privateSet(this, _hasFooter, slot.assignedNodes().length > 0);
      if (had !== __privateGet(this, _hasFooter)) this.requestUpdate();
    });
  }
  render() {
    return T`
      <dialog id="dialog"
              part="dialog"
              @close=${__privateGet(this, _onDialogClose)}
              @cancel=${__privateGet(this, _onDialogCancel)}>
        <cem-pf-v6-button id="close"
                      part="close"
                      variant="plain"
                      aria-label="Close"
                      @click=${__privateGet(this, _onCloseClick)}>
          <svg viewBox="0 0 352 512"
               role="presentation">
            <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
          </svg>
        </cem-pf-v6-button>
        <header id="header"
                part="header"
                ?hidden=${!__privateGet(this, _hasHeader)}>
          <div id="header-main">
            <slot name="header"
                  @slotchange=${__privateGet(this, _onHeaderSlotChange)}></slot>
            <div id="description"
                 part="description"
                 ?hidden=${!__privateGet(this, _hasDescription)}>
              <slot name="description"
                    @slotchange=${__privateGet(this, _onDescriptionSlotChange)}></slot>
            </div>
          </div>
        </header>
        <div id="body"
             part="body">
          <slot></slot>
        </div>
        <footer id="footer"
                part="footer"
                ?hidden=${!__privateGet(this, _hasFooter)}>
          <slot name="footer"
                @slotchange=${__privateGet(this, _onFooterSlotChange)}></slot>
        </footer>
      </dialog>
    `;
  }
  updated(changed) {
    if (changed.has("open")) {
      __privateMethod(this, _PfV6Modal_instances, syncDialogState_fn).call(this);
    }
  }
  firstUpdated() {
    __privateSet(this, _dialog, this.shadowRoot.getElementById("dialog"));
    __privateMethod(this, _PfV6Modal_instances, syncDialogState_fn).call(this);
  }
  /**
   * Opens the modal
   */
  show() {
    this.showModal();
  }
  /**
   * Opens the modal as a modal dialog (with backdrop)
   */
  showModal() {
    this.open = true;
  }
  /**
   * Closes the modal
   * @param returnValue - Optional return value
   */
  close(returnValue) {
    if (typeof returnValue === "string") {
      this.returnValue = returnValue;
      if (__privateGet(this, _dialog)) {
        __privateGet(this, _dialog).returnValue = returnValue;
      }
    }
    this.open = false;
  }
  /**
   * Toggles the modal open/closed
   */
  toggle() {
    if (this.open) {
      this.close();
    } else {
      this.showModal();
    }
  }
};
_init29 = __decoratorStart(_a29);
_dialog = new WeakMap();
_cancelling = new WeakMap();
_hasHeader = new WeakMap();
_hasDescription = new WeakMap();
_hasFooter = new WeakMap();
_variant9 = new WeakMap();
_position = new WeakMap();
_open2 = new WeakMap();
_PfV6Modal_instances = new WeakSet();
syncDialogState_fn = function() {
  if (!__privateGet(this, _dialog)) return;
  if (this.open) {
    if (!__privateGet(this, _dialog).open) {
      __privateGet(this, _dialog).showModal();
      this.dispatchEvent(new PfModalOpenEvent());
    }
  } else {
    if (__privateGet(this, _dialog).open) {
      __privateGet(this, _dialog).close();
    }
  }
};
_onDialogClose = new WeakMap();
_onDialogCancel = new WeakMap();
_onCloseClick = new WeakMap();
_onHeaderSlotChange = new WeakMap();
_onDescriptionSlotChange = new WeakMap();
_onFooterSlotChange = new WeakMap();
__decorateElement(_init29, 4, "variant", _variant_dec9, PfV6Modal, _variant9);
__decorateElement(_init29, 4, "position", _position_dec, PfV6Modal, _position);
__decorateElement(_init29, 4, "open", _open_dec2, PfV6Modal, _open2);
PfV6Modal = __decorateElement(_init29, 0, "PfV6Modal", _PfV6Modal_decorators, PfV6Modal);
__publicField(PfV6Modal, "shadowRootOptions", {
  ...i3.shadowRootOptions,
  delegatesFocus: true
});
__publicField(PfV6Modal, "styles", cem_pf_v6_modal_default);
__runInitializers(_init29, 1, PfV6Modal);

// lit-css:elements/cem-pf-v6-nav-group/cem-pf-v6-nav-group.css
var s33 = new CSSStyleSheet();
s33.replaceSync(JSON.parse('":host {\\n\\n  display: grid;\\n  max-height: 99999px;\\n  padding-block-start: var(--cem-pf-v6-c-nav__subnav--PaddingBlockStart, var(--pf-t--global--spacer--sm));\\n  padding-block-end: var(--cem-pf-v6-c-nav__subnav--PaddingBlockEnd, var(--pf-t--global--spacer--sm));\\n  padding-inline-start: var(--cem-pf-v6-c-nav__subnav--PaddingInlineStart, var(--pf-t--global--spacer--md));\\n  overflow-y: clip;\\n  visibility: visible;\\n  opacity: 1;\\n  translate: 0 var(--cem-pf-v6-c-nav__subnav--TranslateY, 0);\\n  transition: opacity var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default),\\n              translate var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default),\\n              visibility 0s 0s,\\n              max-height 0s 0s,\\n              padding-block-start 0s 0s,\\n              padding-block-end 0s 0s;\\n}\\n\\n@media (prefers-reduced-motion: no-preference) {\\n  :host {\\n    --cem-pf-v6-c-nav__subnav--TranslateY: 0;\\n  }\\n\\n  :host([hidden]) {\\n    --cem-pf-v6-c-nav__subnav--TranslateY: var(--cem-pf-v6-c-nav__subnav--hidden--TranslateY, -.5rem);\\n  }\\n}\\n\\n:host([hidden]) {\\n  --cem-pf-v6-c-nav__subnav--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-nav__subnav--PaddingBlockEnd: 0;\\n\\n  display: grid;\\n  max-height: 0;\\n  visibility: hidden;\\n  opacity: 0;\\n  transition-delay: var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default), var(--pf-t--global--motion--duration--fade--default);\\n  transition-duration: var(--pf-t--global--motion--duration--fade--short), var(--pf-t--global--motion--duration--fade--short), 0s, 0s, 0s, 0s;\\n}\\n\\n#list {\\n  display: grid;\\n  row-gap: var(--cem-pf-v6-c-nav__subnav--RowGap, var(--pf-t--global--border--width--extra-strong));\\n}\\n"'));
var cem_pf_v6_nav_group_default = s33;

// elements/cem-pf-v6-nav-group/cem-pf-v6-nav-group.ts
var _PfV6NavGroup_decorators, _init30, _a30;
_PfV6NavGroup_decorators = [t3("cem-pf-v6-nav-group")];
var PfV6NavGroup = class extends (_a30 = i3) {
  static styles = cem_pf_v6_nav_group_default;
  #internals = this.attachInternals();
  constructor() {
    super();
    this.#internals.role = "region";
  }
  render() {
    return T`
      <section id="subnav"
               part="subnav">
        <div id="list"
             role="list">
          <slot></slot>
        </div>
      </section>
    `;
  }
};
_init30 = __decoratorStart(_a30);
PfV6NavGroup = __decorateElement(_init30, 0, "PfV6NavGroup", _PfV6NavGroup_decorators, PfV6NavGroup);
__runInitializers(_init30, 1, PfV6NavGroup);

// lit-css:elements/cem-pf-v6-nav-item/cem-pf-v6-nav-item.css
var s34 = new CSSStyleSheet();
s34.replaceSync(JSON.parse('":host {\\n\\n  display: grid;\\n}\\n\\n:host([expanded]:last-child) ::slotted(cem-pf-v6-nav-group) {\\n  margin-block-end: calc(var(--cem-pf-v6-c-nav__subnav--PaddingBlockEnd, var(--pf-t--global--spacer--sm)) * -1);\\n}\\n"'));
var cem_pf_v6_nav_item_default = s34;

// elements/cem-pf-v6-nav-item/cem-pf-v6-nav-item.ts
var _expanded_dec9, _a31, _PfV6NavItem_decorators, _init31, _expanded9, _internals6, _handleToggle, _PfV6NavItem_instances, updateExpandedState_fn;
_PfV6NavItem_decorators = [t3("cem-pf-v6-nav-item")];
var PfV6NavItem = class extends (_a31 = i3, _expanded_dec9 = [n4({ type: Boolean, reflect: true })], _a31) {
  constructor() {
    super();
    __privateAdd(this, _PfV6NavItem_instances);
    __privateAdd(this, _expanded9, __runInitializers(_init31, 8, this, false)), __runInitializers(_init31, 11, this);
    __privateAdd(this, _internals6, this.attachInternals());
    __privateAdd(this, _handleToggle, (event) => {
      event.stopPropagation();
      this.expanded = !this.expanded;
    });
    __privateGet(this, _internals6).role = "listitem";
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("pf-nav-toggle", __privateGet(this, _handleToggle));
    const navLink = this.querySelector(":scope > cem-pf-v6-nav-link[expandable]");
    if (navLink?.hasAttribute("expanded") && !this.hasAttribute("expanded")) {
      this.expanded = true;
    }
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("pf-nav-toggle", __privateGet(this, _handleToggle));
  }
  updated(changed) {
    if (changed.has("expanded")) {
      __privateMethod(this, _PfV6NavItem_instances, updateExpandedState_fn).call(this);
    }
  }
  render() {
    return T`<slot></slot>`;
  }
};
_init31 = __decoratorStart(_a31);
_expanded9 = new WeakMap();
_internals6 = new WeakMap();
_handleToggle = new WeakMap();
_PfV6NavItem_instances = new WeakSet();
updateExpandedState_fn = function() {
  const isExpanded = this.expanded;
  this.querySelector(":scope > cem-pf-v6-nav-group")?.toggleAttribute("hidden", !isExpanded);
  this.querySelector(":scope > cem-pf-v6-nav-link[expandable]")?.toggleAttribute("expanded", isExpanded);
};
__decorateElement(_init31, 4, "expanded", _expanded_dec9, PfV6NavItem, _expanded9);
PfV6NavItem = __decorateElement(_init31, 0, "PfV6NavItem", _PfV6NavItem_decorators, PfV6NavItem);
__publicField(PfV6NavItem, "styles", cem_pf_v6_nav_item_default);
__runInitializers(_init31, 1, PfV6NavItem);

// lit-css:elements/cem-pf-v6-nav-link/cem-pf-v6-nav-link.css
var s35 = new CSSStyleSheet();
s35.replaceSync(JSON.parse('":host {\\n\\n  --cem-pf-v6-c-nav__toggle--TranslateY: 0;\\n\\n  position: relative;\\n  display: flex;\\n  column-gap: var(--cem-pf-v6-c-nav__link--ColumnGap, var(--pf-t--global--spacer--sm));\\n  align-items: baseline;\\n}\\n\\n#link {\\n  position: relative;\\n  display: flex;\\n  column-gap: inherit;\\n  align-items: inherit;\\n  padding-block-start: var(--cem-pf-v6-c-nav__link--PaddingBlockStart, var(--pf-t--global--spacer--sm));\\n  padding-block-end: var(--cem-pf-v6-c-nav__link--PaddingBlockEnd, var(--pf-t--global--spacer--sm));\\n  padding-inline-start: var(--cem-pf-v6-c-nav__link--PaddingInlineStart, var(--pf-t--global--spacer--md));\\n  padding-inline-end: var(--cem-pf-v6-c-nav__link--PaddingInlineEnd, var(--pf-t--global--spacer--md));\\n  color: var(--cem-pf-v6-c-nav__link--Color, var(--pf-t--global--text--color--subtle));\\n  text-align: start;\\n  text-decoration: none;\\n  background-color: var(--cem-pf-v6-c-nav__link--BackgroundColor, var(--pf-t--global--background--color--action--plain--default));\\n  border: none;\\n  border-radius: var(--cem-pf-v6-c-nav__link--BorderRadius, var(--pf-t--global--border--radius--small));\\n  transition: background-color var(--pf-t--global--motion--duration--fade--default) var(--pf-t--global--motion--timing-function--default),\\n              color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n  cursor: pointer;\\n  width: 100%;\\n  font: inherit;\\n}\\n\\n#link::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border: var(--cem-pf-v6-c-nav__link--BorderWidth, var(--pf-t--global--border--width--action--plain--default)) solid var(--cem-pf-v6-c-nav__link--BorderColor, var(--pf-t--global--border--color--high-contrast));\\n  border-radius: inherit;\\n}\\n\\n#link:hover,\\n#link:focus {\\n  --cem-pf-v6-c-nav__link--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  color: var(--cem-pf-v6-c-nav__link--hover--Color, var(--pf-t--global--text--color--regular));\\n  background-color: var(--cem-pf-v6-c-nav__link--hover--BackgroundColor, var(--pf-t--global--background--color--action--plain--alt--hover));\\n}\\n\\n:host([current]) #link,\\n:host([current]) #link:hover {\\n  --cem-pf-v6-c-nav__link--BorderWidth: var(--pf-t--global--border--width--action--plain--clicked);\\n  --cem-pf-v6-c-nav__link-icon--Color: var(--pf-t--global--icon--color--regular);\\n  color: var(--cem-pf-v6-c-nav__link--m-current--Color, var(--pf-t--global--text--color--regular));\\n  background-color: var(--cem-pf-v6-c-nav__link--m-current--BackgroundColor, var(--pf-t--global--background--color--action--plain--alt--clicked));\\n}\\n\\n::slotted(svg) {\\n  color: var(--cem-pf-v6-c-nav__link-icon--Color, var(--pf-t--global--icon--color--subtle));\\n}\\n\\n/* Toggle wrapper for expandable items */\\n#toggle {\\n  flex: none;\\n  align-self: start;\\n  margin-inline-start: auto;\\n  transform: translateY(var(--cem-pf-v6-c-nav__toggle--TranslateY));\\n}\\n\\n/* Hide toggle wrapper when not expandable */\\n:host(:not([expandable])) #toggle {\\n  display: none;\\n}\\n\\n#toggle-icon {\\n  display: inline-flex;\\n  transition: transform var(--pf-t--global--motion--duration--icon--default, 0.2s) var(--pf-t--global--motion--timing-function--default, ease);\\n  transform: rotate(var(--cem-pf-v6-c-nav__toggle-icon--Rotate, 0));\\n}\\n\\n#link:where([aria-expanded=\\"true\\"]) #toggle-icon {\\n  transform: rotate(var(--cem-pf-v6-c-nav__item--m-expanded__toggle-icon--Rotate, 90deg));\\n}\\n\\n.cem-pf-v6-svg {\\n  width: 1em;\\n  height: 1em;\\n  vertical-align: -0.125em;\\n}\\n"'));
var cem_pf_v6_nav_link_default = s35;

// elements/cem-pf-v6-nav-link/cem-pf-v6-nav-link.ts
var PfNavToggleEvent = class extends Event {
  expanded;
  constructor(expanded) {
    super("pf-nav-toggle", { bubbles: true });
    this.expanded = expanded;
  }
};
var _expanded_dec10, _expandable_dec4, _current_dec2, _label_dec4, _href_dec3, _a32, _PfV6NavLink_decorators, _init32, _href3, _label4, _current2, _expandable4, _expanded10, _PfV6NavLink_instances, onClick_fn, markCurrentIfMatches_fn, renderLinkContent_fn;
_PfV6NavLink_decorators = [t3("cem-pf-v6-nav-link")];
var PfV6NavLink = class extends (_a32 = i3, _href_dec3 = [n4({ reflect: true })], _label_dec4 = [n4()], _current_dec2 = [n4({ type: Boolean, reflect: true })], _expandable_dec4 = [n4({ type: Boolean, reflect: true })], _expanded_dec10 = [n4({ type: Boolean, reflect: true })], _a32) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6NavLink_instances);
    __privateAdd(this, _href3, __runInitializers(_init32, 8, this)), __runInitializers(_init32, 11, this);
    __privateAdd(this, _label4, __runInitializers(_init32, 12, this)), __runInitializers(_init32, 15, this);
    __privateAdd(this, _current2, __runInitializers(_init32, 16, this, false)), __runInitializers(_init32, 19, this);
    __privateAdd(this, _expandable4, __runInitializers(_init32, 20, this, false)), __runInitializers(_init32, 23, this);
    __privateAdd(this, _expanded10, __runInitializers(_init32, 24, this, false)), __runInitializers(_init32, 27, this);
  }
  connectedCallback() {
    super.connectedCallback();
    __privateMethod(this, _PfV6NavLink_instances, markCurrentIfMatches_fn).call(this);
  }
  render() {
    if (this.href) {
      return T`
        <a id="link"
           part="link"
           href=${this.href}
           aria-label=${this.label ?? A}
           aria-current=${this.current ? "page" : A}
           @click=${__privateMethod(this, _PfV6NavLink_instances, onClick_fn)}>
          ${__privateMethod(this, _PfV6NavLink_instances, renderLinkContent_fn).call(this)}
        </a>
      `;
    }
    return T`
      <button id="link"
              part="link"
              type="button"
              aria-label=${this.label ?? A}
              aria-expanded=${this.expandable ? String(this.expanded) : A}
              @click=${__privateMethod(this, _PfV6NavLink_instances, onClick_fn)}>
        ${__privateMethod(this, _PfV6NavLink_instances, renderLinkContent_fn).call(this)}
      </button>
    `;
  }
};
_init32 = __decoratorStart(_a32);
_href3 = new WeakMap();
_label4 = new WeakMap();
_current2 = new WeakMap();
_expandable4 = new WeakMap();
_expanded10 = new WeakMap();
_PfV6NavLink_instances = new WeakSet();
onClick_fn = function(e6) {
  if (this.expandable) {
    e6.preventDefault();
    this.dispatchEvent(new PfNavToggleEvent(!this.expanded));
  }
};
markCurrentIfMatches_fn = function() {
  if (this.href && this.href === window.location.pathname) {
    this.current = true;
  }
};
renderLinkContent_fn = function() {
  return T`
      <slot name="icon-start"></slot>
      <slot></slot>
      <slot name="icon-end"></slot>
      <span id="toggle"
            part="toggle">
        <span id="toggle-icon"
              part="toggle-icon">
          <svg class="cem-pf-v6-svg"
               viewBox="0 0 256 512"
               fill="currentColor"
               role="presentation">
            <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
          </svg>
        </span>
      </span>
    `;
};
__decorateElement(_init32, 4, "href", _href_dec3, PfV6NavLink, _href3);
__decorateElement(_init32, 4, "label", _label_dec4, PfV6NavLink, _label4);
__decorateElement(_init32, 4, "current", _current_dec2, PfV6NavLink, _current2);
__decorateElement(_init32, 4, "expandable", _expandable_dec4, PfV6NavLink, _expandable4);
__decorateElement(_init32, 4, "expanded", _expanded_dec10, PfV6NavLink, _expanded10);
PfV6NavLink = __decorateElement(_init32, 0, "PfV6NavLink", _PfV6NavLink_decorators, PfV6NavLink);
__publicField(PfV6NavLink, "styles", cem_pf_v6_nav_link_default);
__runInitializers(_init32, 1, PfV6NavLink);

// lit-css:elements/cem-pf-v6-nav-list/cem-pf-v6-nav-list.css
var s36 = new CSSStyleSheet();
s36.replaceSync(JSON.parse('":host {\\n\\n  display: grid;\\n  row-gap: var(--cem-pf-v6-c-nav__list--RowGap, var(--pf-t--global--spacer--sm));\\n  column-gap: var(--cem-pf-v6-c-nav__list--ColumnGap, var(--pf-t--global--spacer--xs));\\n  min-height: 0;\\n}\\n"'));
var cem_pf_v6_nav_list_default = s36;

// elements/cem-pf-v6-nav-list/cem-pf-v6-nav-list.ts
var _PfV6NavList_decorators, _init33, _a33;
_PfV6NavList_decorators = [t3("cem-pf-v6-nav-list")];
var PfV6NavList = class extends (_a33 = i3) {
  static styles = cem_pf_v6_nav_list_default;
  #internals = this.attachInternals();
  constructor() {
    super();
    this.#internals.role = "list";
  }
  render() {
    return T`<slot></slot>`;
  }
};
_init33 = __decoratorStart(_a33);
PfV6NavList = __decorateElement(_init33, 0, "PfV6NavList", _PfV6NavList_decorators, PfV6NavList);
__runInitializers(_init33, 1, PfV6NavList);

// lit-css:elements/cem-pf-v6-navigation/cem-pf-v6-navigation.css
var s37 = new CSSStyleSheet();
s37.replaceSync(JSON.parse('":host {\\n  --cem-pf-v6-c-nav--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-nav--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-nav--RowGap: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-nav--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-nav--FontWeight: var(--pf-t--global--font--weight--body--default);\\n  --cem-pf-v6-c-nav--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-nav--BackgroundColor: transparent;\\n  --cem-pf-v6-c-nav__list--RowGap: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-nav__list--ColumnGap: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-nav__link--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-nav__link--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-nav__link--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav__link--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav__link--ColumnGap: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-nav__link--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-nav__link--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-nav__link--Color: var(--pf-t--global--text--color--subtle);\\n  --cem-pf-v6-c-nav__link--hover--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-nav__link--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--alt--hover);\\n  --cem-pf-v6-c-nav__link--m-current--BackgroundColor: var(--pf-t--global--background--color--action--plain--alt--clicked);\\n  --cem-pf-v6-c-nav__link--m-current--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-nav__link--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-nav__link--BorderWidth: var(--pf-t--global--border--width--action--plain--default);\\n  --cem-pf-v6-c-nav__link--hover--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --cem-pf-v6-c-nav__link--m-current--BorderWidth: var(--pf-t--global--border--width--action--plain--clicked);\\n  --cem-pf-v6-c-nav__link-icon--Color: var(--pf-t--global--icon--color--subtle);\\n  --cem-pf-v6-c-nav__link--m-current__link-icon--Color: var(--pf-t--global--icon--color--regular);\\n  --cem-pf-v6-c-nav__toggle-icon--Rotate: 0;\\n  --cem-pf-v6-c-nav__item--m-expanded__toggle-icon--Rotate: 90deg;\\n  --cem-pf-v6-c-nav--ColumnGap: 0;\\n  --cem-pf-v6-c-nav--AlignItems: baseline;\\n  --cem-pf-v6-c-nav--OutlineOffset: calc(var(--pf-t--global--spacer--xs) * -1);\\n  --cem-pf-v6-c-nav__list--ScrollSnapTypeAxis: x;\\n  --cem-pf-v6-c-nav__list--ScrollSnapTypeStrictness: proximity;\\n  --cem-pf-v6-c-nav__list--ScrollSnapType: var(--cem-pf-v6-c-nav__list--ScrollSnapTypeAxis) var(--cem-pf-v6-c-nav__list--ScrollSnapTypeStrictness);\\n  --cem-pf-v6-c-nav__item--ScrollSnapAlign: end;\\n  --cem-pf-v6-c-nav__section-title--FontWeight: var(--pf-t--global--font--weight--body--bold);\\n  --cem-pf-v6-c-nav__section-title--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-nav__section-title--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-nav__section-title--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-nav__section-title--PaddingInlineStart: var(--cem-pf-v6-c-nav__link--PaddingInlineStart);\\n  --cem-pf-v6-c-nav__section-title--PaddingInlineEnd: var(--cem-pf-v6-c-nav__link--PaddingInlineEnd);\\n  --cem-pf-v6-c-nav__item--RowGap: var(--cem-pf-v6-c-nav__list--RowGap);\\n  --cem-pf-v6-c-nav__item__toggle-icon--Rotate: 0;\\n  --cem-pf-v6-c-nav__item__toggle-icon--TransitionDuration--transform: var(--pf-t--global--motion--duration--icon--default);\\n  --cem-pf-v6-c-nav__item__toggle-icon--TransitionTimingFunction--transform: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-nav__link--AlignItems: baseline;\\n  --cem-pf-v6-c-nav__link--WhiteSpace: normal;\\n  --cem-pf-v6-c-nav__link--TransitionDuration--background-color: var(--pf-t--global--motion--duration--fade--default);\\n  --cem-pf-v6-c-nav__link--TransitionTimingFunction--background-color: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-nav__link--m-current--TransitionDuration--color: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-nav__link--m-current--TransitionTimingFunction--color: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-nav__subnav--RowGap: var(--pf-t--global--border--width--extra-strong);\\n  --cem-pf-v6-c-nav__subnav--PaddingBlockStart: var(--cem-pf-v6-c-nav__list--RowGap);\\n  --cem-pf-v6-c-nav__subnav--PaddingBlockEnd: var(--cem-pf-v6-c-nav__list--RowGap);\\n  --cem-pf-v6-c-nav__subnav--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav__subnav--TranslateY: 0;\\n  --cem-pf-v6-c-nav__subnav--hidden--TranslateY: -.5rem;\\n  --cem-pf-v6-c-nav__subnav--TransitionDuration--expand: var(--pf-t--global--motion--duration--fade--default);\\n  --cem-pf-v6-c-nav__subnav--TransitionDuration--expand--slide: 0s;\\n  --cem-pf-v6-c-nav__subnav--TransitionDuration--collapse: var(--pf-t--global--motion--duration--fade--short);\\n  --cem-pf-v6-c-nav__subnav--TransitionDuration--collapse--slide: 0s;\\n  --cem-pf-v6-c-nav__subnav--TransitionTimingFunction--expand: var(--pf-t--global--motion--timing-function--default);\\n  --cem-pf-v6-c-nav__scroll-button--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-nav__scroll-button--BorderWidth: var(--pf-t--global--border--width--divider--default);\\n  --cem-pf-v6-c-nav__scroll-button--first-of-type--c-button--BorderStartStartRadius: var(--pf-t--global--border--radius--pill);\\n  --cem-pf-v6-c-nav__scroll-button--first-of-type--c-button--BorderEndStartRadius: var(--pf-t--global--border--radius--pill);\\n  --cem-pf-v6-c-nav__scroll-button--last-of-type--c-button--BorderStartEndRadius: var(--pf-t--global--border--radius--pill);\\n  --cem-pf-v6-c-nav__scroll-button--last-of-type--c-button--BorderEndEndRadius: var(--pf-t--global--border--radius--pill);\\n  --cem-pf-v6-c-nav__toggle--PaddingInlineStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-nav__toggle--PaddingInlineEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-nav__toggle--TranslateY: calc((var(--cem-pf-v6-c-nav--LineHeight) * var(--cem-pf-v6-c-nav--FontSize) / 2) - 50%);\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-nav--m-horizontal--m-scrollable__list--PaddingInlineStart: var(--cem-pf-v6-c-nav__list--ColumnGap);\\n  --cem-pf-v6-c-nav--m-horizontal--m-scrollable__list--PaddingInlineEnd: var(--cem-pf-v6-c-nav__list--ColumnGap);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav--BackgroundColor: var(--pf-t--global--background--color--secondary--default);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderRadius: var(--pf-t--global--border--radius--pill);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav__list--PaddingBlockStart: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav__list--PaddingBlockEnd: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav__list--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav__list--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav--m-scrollable__list--PaddingInlineStart: var(--cem-pf-v6-c-nav__list--ColumnGap);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav--m-scrollable__list--PaddingInlineEnd: var(--cem-pf-v6-c-nav__list--ColumnGap);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav__link--PaddingBlockStart: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav__link--PaddingBlockEnd: var(--pf-t--global--spacer--xs);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav__link--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav--m-horizontal--m-subnav__link--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n\\n  display: grid;\\n  position: relative;\\n  row-gap: var(--cem-pf-v6-c-nav--RowGap);\\n  max-width: 100%;\\n  padding-block-start: var(--cem-pf-v6-c-nav--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-nav--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-nav--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-nav--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-nav--FontSize);\\n  font-weight: var(--cem-pf-v6-c-nav--FontWeight);\\n  line-height: var(--cem-pf-v6-c-nav--LineHeight);\\n  background-color: var(--cem-pf-v6-c-nav--BackgroundColor);\\n}\\n\\n:host([inset]) {\\n  --cem-pf-v6-c-nav--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-nav--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n}\\n\\n:host([horizontal]) {\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-nav--m-horizontal--m-scrollable__list--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-nav--m-horizontal--m-scrollable__list--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-nav__list--ScrollSnapType: x mandatory;\\n  --cem-pf-v6-c-nav__item--ScrollSnapAlign: start;\\n  --cem-pf-v6-c-nav__scroll-button--BorderWidth: var(--pf-t--global--border--width--regular);\\n  --cem-pf-v6-c-nav__scroll-button--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-nav__scroll-button--first-of-type--c-button--BorderStartStartRadius: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderRadius, var(--pf-t--global--border--radius--small));\\n  --cem-pf-v6-c-nav__scroll-button--first-of-type--c-button--BorderEndStartRadius: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderRadius, var(--pf-t--global--border--radius--small));\\n  --cem-pf-v6-c-nav__scroll-button--last-of-type--c-button--BorderStartEndRadius: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderRadius, var(--pf-t--global--border--radius--small));\\n  --cem-pf-v6-c-nav__scroll-button--last-of-type--c-button--BorderEndEndRadius: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderRadius, var(--pf-t--global--border--radius--small));\\n  --cem-pf-v6-c-nav--OutlineOffset: calc(var(--pf-t--global--spacer--control-vertical--default) * -1);\\n\\n  display: flex;\\n  padding: 0;\\n  overflow: hidden;\\n}\\n\\n:host([horizontal]) nav {\\n  flex: 1;\\n  min-width: 0;\\n  overflow-x: auto;\\n  scrollbar-width: none;\\n  -ms-overflow-style: -ms-autohiding-scrollbar;\\n}\\n\\n:host([horizontal]) nav::-webkit-scrollbar {\\n  display: none;\\n}\\n\\n:host([horizontal]) ::slotted(cem-pf-v6-nav-list) {\\n  display: flex;\\n  padding-block-start: var(--cem-pf-v6-c-nav--m-horizontal__list--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-nav--m-horizontal__list--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineEnd);\\n  white-space: nowrap;\\n}\\n\\n:host([horizontal]) ::slotted(cem-pf-v6-nav-item) {\\n  scroll-snap-align: var(--cem-pf-v6-c-nav__item--ScrollSnapAlign);\\n  flex-shrink: 0;\\n}\\n\\n:host([horizontal]) ::slotted(cem-pf-v6-nav-link) {\\n  outline-offset: var(--cem-pf-v6-c-nav--OutlineOffset);\\n}\\n\\ncem-pf-v6-button.scroll-button {\\n  position: relative;\\n  outline-offset: var(--cem-pf-v6-c-nav--OutlineOffset);\\n}\\n\\ncem-pf-v6-button.scroll-button::before {\\n  position: absolute;\\n  inset: 0;\\n  content: \\"\\";\\n}\\n\\ncem-pf-v6-button.scroll-button:first-of-type {\\n  --cem-pf-v6-c-button--BorderStartStartRadius: var(--cem-pf-v6-c-nav__scroll-button--first-of-type--c-button--BorderStartStartRadius);\\n  --cem-pf-v6-c-button--BorderEndStartRadius: var(--cem-pf-v6-c-nav__scroll-button--first-of-type--c-button--BorderEndStartRadius);\\n}\\n\\ncem-pf-v6-button.scroll-button:first-of-type::before {\\n  border-inline-end: var(--cem-pf-v6-c-nav__scroll-button--BorderWidth) solid var(--cem-pf-v6-c-nav__scroll-button--BorderColor);\\n}\\n\\ncem-pf-v6-button.scroll-button:last-of-type {\\n  --cem-pf-v6-c-button--BorderStartEndRadius: var(--cem-pf-v6-c-nav__scroll-button--last-of-type--c-button--BorderStartEndRadius);\\n  --cem-pf-v6-c-button--BorderEndEndRadius: var(--cem-pf-v6-c-nav__scroll-button--last-of-type--c-button--BorderEndEndRadius);\\n}\\n\\ncem-pf-v6-button.scroll-button:last-of-type::before {\\n  border-inline-start: var(--cem-pf-v6-c-nav__scroll-button--BorderWidth) solid var(--cem-pf-v6-c-nav__scroll-button--BorderColor);\\n}\\n\\n:host(:not([scrollable])) cem-pf-v6-button.scroll-button {\\n  display: none;\\n}\\n\\n:host([horizontal][scrollable]) {\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineStart: var(--cem-pf-v6-c-nav--m-horizontal--m-scrollable__list--PaddingInlineStart);\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineEnd: var(--cem-pf-v6-c-nav--m-horizontal--m-scrollable__list--PaddingInlineEnd);\\n}\\n\\n:host([horizontal][variant=\\"horizontal-subnav\\"]) {\\n  --cem-pf-v6-c-nav--BackgroundColor: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--BackgroundColor);\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingBlockStart: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav__list--PaddingBlockStart);\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingBlockEnd: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav__list--PaddingBlockEnd);\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineStart: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav__list--PaddingInlineStart);\\n  --cem-pf-v6-c-nav--m-horizontal__list--PaddingInlineEnd: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav__list--PaddingInlineEnd);\\n  --cem-pf-v6-c-nav--m-horizontal--m-scrollable__list--PaddingInlineStart: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--m-scrollable__list--PaddingInlineStart);\\n  --cem-pf-v6-c-nav--m-horizontal--m-scrollable__list--PaddingInlineEnd: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--m-scrollable__list--PaddingInlineEnd);\\n  --cem-pf-v6-c-nav__link--PaddingBlockStart: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav__link--PaddingBlockStart);\\n  --cem-pf-v6-c-nav__link--PaddingBlockEnd: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav__link--PaddingBlockEnd);\\n  border: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderWidth) solid var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-nav--m-horizontal--m-subnav--BorderRadius);\\n}\\n\\nsvg {\\n  width: 1em;\\n  height: 1em;\\n  vertical-align: -0.125em;\\n  fill: currentColor;\\n}\\n"'));
var cem_pf_v6_navigation_default = s37;

// elements/cem-pf-v6-navigation/cem-pf-v6-navigation.ts
var _inset_dec, _scrollable_dec, _horizontal_dec2, _variant_dec10, _a34, _PfV6Navigation_decorators, _init34, _variant10, _horizontal2, _scrollable, _inset, _nav, _navList, _resizeObserver, _PfV6Navigation_instances, onSlotChange_fn3, setupScrollObserver_fn, _handleScrollButtons, scrollBack_fn, scrollForward_fn, teardownScrollObserver_fn;
_PfV6Navigation_decorators = [t3("cem-pf-v6-navigation")];
var PfV6Navigation = class extends (_a34 = i3, _variant_dec10 = [n4({ reflect: true })], _horizontal_dec2 = [n4({ type: Boolean, reflect: true })], _scrollable_dec = [n4({ type: Boolean, reflect: true })], _inset_dec = [n4({ type: Boolean, reflect: true })], _a34) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Navigation_instances);
    __privateAdd(this, _variant10, __runInitializers(_init34, 8, this)), __runInitializers(_init34, 11, this);
    __privateAdd(this, _horizontal2, __runInitializers(_init34, 12, this, false)), __runInitializers(_init34, 15, this);
    __privateAdd(this, _scrollable, __runInitializers(_init34, 16, this, false)), __runInitializers(_init34, 19, this);
    __privateAdd(this, _inset, __runInitializers(_init34, 20, this, false)), __runInitializers(_init34, 23, this);
    __privateAdd(this, _nav);
    __privateAdd(this, _navList);
    __privateAdd(this, _resizeObserver);
    __privateAdd(this, _handleScrollButtons, () => {
      if (!__privateGet(this, _navList) || !this.horizontal || !__privateGet(this, _nav)) return;
      const { scrollLeft, scrollWidth, clientWidth } = __privateGet(this, _nav);
      const isOverflowing = scrollWidth > clientWidth;
      const scrollViewAtStart = scrollLeft <= 1;
      const scrollViewAtEnd = scrollLeft + clientWidth >= scrollWidth - 1;
      this.scrollable = isOverflowing && (!scrollViewAtStart || !scrollViewAtEnd);
      const scrollBack = this.shadowRoot?.getElementById("scroll-back");
      const scrollForward = this.shadowRoot?.getElementById("scroll-forward");
      scrollBack?.toggleAttribute("disabled", scrollViewAtStart);
      scrollForward?.toggleAttribute("disabled", scrollViewAtEnd);
    });
  }
  render() {
    return T`
      <cem-pf-v6-button id="scroll-back"
                    class="scroll-button"
                    variant="plain"
                    aria-label="Scroll back"
                    @click=${__privateMethod(this, _PfV6Navigation_instances, scrollBack_fn)}>
        <svg viewBox="0 0 256 512" role="presentation">
          <path d="M31.7 239l136-136c9.4-9.4 24.6-9.4 33.9 0l22.6 22.6c9.4 9.4 9.4 24.6 0 33.9L127.9 256l96.4 96.4c9.4 9.4 9.4 24.6 0 33.9L201.7 409c-9.4 9.4-24.6 9.4-33.9 0l-136-136c-9.5-9.4-9.5-24.6-.1-34z"></path>
        </svg>
      </cem-pf-v6-button>
      <nav part="nav"
           aria-label=${this.getAttribute("aria-label") || ""}
           @scroll=${__privateGet(this, _handleScrollButtons)}>
        <slot @slotchange=${__privateMethod(this, _PfV6Navigation_instances, onSlotChange_fn3)}></slot>
      </nav>
      <cem-pf-v6-button id="scroll-forward"
                    class="scroll-button"
                    variant="plain"
                    aria-label="Scroll forward"
                    @click=${__privateMethod(this, _PfV6Navigation_instances, scrollForward_fn)}>
        <svg viewBox="0 0 256 512" role="presentation">
          <path d="M224.3 273l-136 136c-9.4 9.4-24.6 9.4-33.9 0l-22.6-22.6c-9.4-9.4-9.4-24.6 0-33.9l96.4-96.4-96.4-96.4c-9.4-9.4-9.4-24.6 0-33.9L54.3 103c9.4-9.4 24.6-9.4 33.9 0l136 136c9.5 9.4 9.5 24.6.1 34z"></path>
        </svg>
      </cem-pf-v6-button>
    `;
  }
  firstUpdated() {
    __privateSet(this, _nav, this.shadowRoot?.querySelector("nav"));
  }
  updated(changed) {
    if (changed.has("horizontal")) {
      if (this.horizontal && __privateGet(this, _navList)) {
        __privateMethod(this, _PfV6Navigation_instances, setupScrollObserver_fn).call(this);
      } else {
        __privateMethod(this, _PfV6Navigation_instances, teardownScrollObserver_fn).call(this);
      }
    }
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateMethod(this, _PfV6Navigation_instances, teardownScrollObserver_fn).call(this);
  }
};
_init34 = __decoratorStart(_a34);
_variant10 = new WeakMap();
_horizontal2 = new WeakMap();
_scrollable = new WeakMap();
_inset = new WeakMap();
_nav = new WeakMap();
_navList = new WeakMap();
_resizeObserver = new WeakMap();
_PfV6Navigation_instances = new WeakSet();
onSlotChange_fn3 = function(e6) {
  const slot = e6.target;
  const elements = slot.assignedElements();
  __privateSet(this, _navList, elements.find((el) => el.tagName === "CEM-PF-V6-NAV-LIST"));
  if (__privateGet(this, _navList) && this.horizontal) {
    __privateMethod(this, _PfV6Navigation_instances, setupScrollObserver_fn).call(this);
  }
};
setupScrollObserver_fn = function() {
  if (!__privateGet(this, _navList) || !__privateGet(this, _nav)) return;
  if (__privateGet(this, _resizeObserver)) return;
  __privateSet(this, _resizeObserver, new ResizeObserver(() => {
    __privateGet(this, _handleScrollButtons).call(this);
  }));
  __privateGet(this, _resizeObserver).observe(__privateGet(this, _nav));
  __privateGet(this, _resizeObserver).observe(__privateGet(this, _navList));
  requestAnimationFrame(() => {
    __privateGet(this, _handleScrollButtons).call(this);
  });
};
_handleScrollButtons = new WeakMap();
scrollBack_fn = function() {
  __privateGet(this, _nav)?.scrollBy({ left: -200, behavior: "smooth" });
};
scrollForward_fn = function() {
  __privateGet(this, _nav)?.scrollBy({ left: 200, behavior: "smooth" });
};
teardownScrollObserver_fn = function() {
  __privateGet(this, _resizeObserver)?.disconnect();
  __privateSet(this, _resizeObserver, void 0);
};
__decorateElement(_init34, 4, "variant", _variant_dec10, PfV6Navigation, _variant10);
__decorateElement(_init34, 4, "horizontal", _horizontal_dec2, PfV6Navigation, _horizontal2);
__decorateElement(_init34, 4, "scrollable", _scrollable_dec, PfV6Navigation, _scrollable);
__decorateElement(_init34, 4, "inset", _inset_dec, PfV6Navigation, _inset);
PfV6Navigation = __decorateElement(_init34, 0, "PfV6Navigation", _PfV6Navigation_decorators, PfV6Navigation);
__publicField(PfV6Navigation, "styles", cem_pf_v6_navigation_default);
__runInitializers(_init34, 1, PfV6Navigation);

// lit-css:elements/cem-pf-v6-page/cem-pf-v6-page.css
var s38 = new CSSStyleSheet();
s38.replaceSync(JSON.parse('":host {\\n  container-type: inline-size;\\n  display: grid;\\n  grid-template-areas:\\n    \\"header\\"\\n    \\"main\\";\\n  grid-template-rows: max-content 1fr;\\n  grid-template-columns: 1fr;\\n  height: 100vh;\\n  --cem-pf-v6-c-page--BackgroundColor: var(--pf-t--global--background--color--secondary--default);\\n  --cem-pf-v6-c-page--inset: var(--pf-t--global--spacer--inset--page-chrome);\\n  --cem-pf-v6-c-page--c-masthead--ZIndex: var(--pf-t--global--z-index--md);\\n  --cem-pf-v6-c-page__sidebar--ZIndex: var(--pf-t--global--z-index--sm);\\n  --cem-pf-v6-c-page__sidebar--Width--base: 18.125rem;\\n  --cem-pf-v6-c-page__sidebar--Width: var(--cem-pf-v6-c-page__sidebar--Width--base);\\n  --cem-pf-v6-c-page__sidebar--xl--Width: var(--cem-pf-v6-c-page__sidebar--Width--base);\\n  --cem-pf-v6-c-page__sidebar--BackgroundColor: var(--pf-t--global--background--color--secondary--default);\\n  --cem-pf-v6-c-page__sidebar--BoxShadow: var(--pf-t--global--box-shadow--md--right);\\n  --cem-pf-v6-c-page__sidebar--TransitionProperty: opacity;\\n  --cem-pf-v6-c-page__sidebar--TransitionDuration: var(--pf-t--global--motion--duration--fade--default);\\n  --cem-pf-v6-c-page__sidebar--TransitionTimingFunction: var(--pf-t--global--motion--timing-function--decelerate);\\n  --cem-pf-v6-c-page__sidebar--TranslateX: -100%;\\n  --cem-pf-v6-c-page__sidebar--Opacity: 0;\\n  --cem-pf-v6-c-page__sidebar--m-expanded--Opacity: 1;\\n  --cem-pf-v6-c-page__sidebar--xl--Opacity: 1;\\n  --cem-pf-v6-c-page__sidebar--TranslateZ: 0;\\n  --cem-pf-v6-c-page__sidebar--m-expanded--TranslateX: 0;\\n  --cem-pf-v6-c-page__sidebar--xl--TranslateX: 0;\\n  --cem-pf-v6-c-page__sidebar--MarginInlineEnd: calc(var(--pf-t--global--spacer--inset--page-chrome) / 2);\\n  --cem-pf-v6-c-page__sidebar--PaddingBlockEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-page__sidebar--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-page__sidebar--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-page__sidebar--BorderInlineEndWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-page__sidebar--BorderInlineEndColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-page__sidebar-header--BorderBlockEndWidth: var(--pf-t--global--border--width--divider--default);\\n  --cem-pf-v6-c-page__sidebar-header--BorderBlockEndColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-page__sidebar-header--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-page__sidebar-header--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__sidebar-title--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-page__sidebar-title--FontSize: var(--pf-t--global--font--size--heading--xs);\\n  --cem-pf-v6-c-page__sidebar-title--LineHeight: var(--pf-t--global--font--line-height--heading);\\n  --cem-pf-v6-c-page__sidebar-title--FontFamily: var(--pf-t--global--font--family--heading);\\n  --cem-pf-v6-c-page__sidebar-title--FontWeight: var(--pf-t--global--font--weight--heading--default);\\n  --cem-pf-v6-c-page__sidebar-body--PaddingInlineEnd: calc(var(--pf-t--global--spacer--inset--page-chrome) / 2);\\n  --cem-pf-v6-c-page__sidebar-body--PaddingInlineStart: var(--pf-t--global--spacer--inset--page-chrome);\\n  --cem-pf-v6-c-page__sidebar-body--m-page-insets--PaddingInlineEnd: var(--cem-pf-v6-c-page--inset);\\n  --cem-pf-v6-c-page__sidebar-body--m-page-insets--PaddingInlineStart: var(--cem-pf-v6-c-page--inset);\\n  --cem-pf-v6-c-page__sidebar-body--m-context-selector--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__sidebar-body--m-context-selector--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__main-container--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-page__main-container--GridArea: main;\\n  --cem-pf-v6-c-page--masthead--main-container--GridArea: sidebar / sidebar / main / main;\\n  --cem-pf-v6-c-page__main-container--MaxHeight: calc(100% - var(--pf-t--global--spacer--lg));\\n  --cem-pf-v6-c-page__main-container--AlignSelf: start;\\n  --cem-pf-v6-c-page__main-container--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-page__main-container--MarginInlineStart: var(--cem-pf-v6-c-page--inset);\\n  --cem-pf-v6-c-page__main-container--MarginInlineEnd: var(--cem-pf-v6-c-page--inset);\\n  --cem-pf-v6-c-page__main-container--BorderRadius: var(--pf-t--global--border--radius--medium);\\n  --cem-pf-v6-c-page__main-container--BorderBlockStartWidth: var(--pf-t--global--border--width--main--default);\\n  --cem-pf-v6-c-page__main-container--BorderBlockEndWidth: var(--pf-t--global--border--width--main--default);\\n  --cem-pf-v6-c-page__main-container--BorderInlineStartWidth: var(--pf-t--global--border--width--main--default);\\n  --cem-pf-v6-c-page__main-container--BorderInlineEndWidth: var(--pf-t--global--border--width--main--default);\\n  --cem-pf-v6-c-page__main-container--BorderColor: var(--pf-t--global--border--color--main--default);\\n  --cem-pf-v6-c-page__main-container--xs--AlignSelf: stretch;\\n  --cem-pf-v6-c-page__main-container--xs--BorderRadius: 0;\\n  --cem-pf-v6-c-page__main-container--xs--MarginInlineStart: 0;\\n  --cem-pf-v6-c-page__main-container--xs--MaxHeight: 100%;\\n  --cem-pf-v6-c-page__main-container--xs--MarginInlineEnd: 0;\\n  --cem-pf-v6-c-page__main-container--xs--BorderBlockEndWidth: 0;\\n  --cem-pf-v6-c-page__main-container--xs--BorderInlineStartWidth: 0px;\\n  --cem-pf-v6-c-page__main-container--xs--BorderInlineEndWidth: 0px;\\n  --cem-pf-v6-c-page__main-section--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__main-section--PaddingInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-page__main-section--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__main-section--PaddingInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-page__main-section--RowGap: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-page__main-breadcrumb--main-section--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__main-section--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-page__main-section--m-secondary--BackgroundColor: var(--pf-t--global--background--color--secondary--default);\\n  --cem-pf-v6-c-page__main-breadcrumb--page__main-tabs--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page--section--m-limit-width--MaxWidth: calc(125rem - var(--cem-pf-v6-c-page__sidebar--Width));\\n  --cem-pf-v6-c-page--section--m-sticky-top--ZIndex: var(--pf-t--global--z-index--md);\\n  --cem-pf-v6-c-page--section--m-sticky-top--BoxShadow: var(--pf-t--global--box-shadow--sm--bottom);\\n  --cem-pf-v6-c-page--section--m-sticky-top--BorderBlockEndWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-page--section--m-sticky-top--BorderBlockEndColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-page--section--m-sticky-bottom--ZIndex: var(--pf-t--global--z-index--md);\\n  --cem-pf-v6-c-page--section--m-sticky-bottom--BoxShadow: var(--pf-t--global--box-shadow--sm--top);\\n  --cem-pf-v6-c-page--section--m-sticky-bottom--BorderBlockStartWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-page--section--m-sticky-bottom--BorderBlockStartColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-page--section--m-shadow-bottom--BoxShadow: var(--pf-t--global--box-shadow--sm--bottom);\\n  --cem-pf-v6-c-page--section--m-shadow-bottom--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-page--section--m-shadow-bottom--BorderBlockEndWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-page--section--m-shadow-bottom--BorderBlockEndColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-page--section--m-shadow-top--BoxShadow: var(--pf-t--global--box-shadow--sm--top);\\n  --cem-pf-v6-c-page--section--m-shadow-top--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-page--section--m-shadow-top--BorderBlockStartWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-page--section--m-shadow-top--BorderBlockStartColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-page__main-subnav--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-page__main-subnav--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__main-subnav--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-page__main-subnav--PaddingInlineStart: calc(var(--pf-t--global--spacer--lg) - var(--cem-pf-v6-c-page__main-container--BorderInlineStartWidth));\\n  --cem-pf-v6-c-page__main-subnav--PaddingInlineEnd: calc(var(--pf-t--global--spacer--lg) - var(--cem-pf-v6-c-page__main-container--BorderInlineEndWidth));\\n  --cem-pf-v6-c-page__main-subnav--m-sticky-top--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__main-breadcrumb--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__main-breadcrumb--PaddingInlineEnd: calc(var(--pf-t--global--spacer--lg) - var(--cem-pf-v6-c-page__main-container--BorderInlineEndWidth));\\n  --cem-pf-v6-c-page__main-breadcrumb--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-page__main-breadcrumb--PaddingInlineStart: calc(var(--pf-t--global--spacer--lg) - var(--cem-pf-v6-c-page__main-container--BorderInlineStartWidth));\\n  --cem-pf-v6-c-page__main-breadcrumb--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-page__main-breadcrumb--m-sticky-top--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-page__main-tabs--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-page__main-tabs--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-page__main-tabs--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-page__main-tabs--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-page__main-tabs--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-page__main-wizard--BackgroundColor: var(--pf-t--global--background--color--primary--default);\\n  --cem-pf-v6-c-page__main-wizard--BorderBlockStartColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-page__main-wizard--BorderBlockStartWidth: var(--pf-t--global--border--width--action--default);\\n  --cem-pf-v6-c-page__drawer--c-drawer--BorderBlockStartWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-page__drawer--c-drawer--BorderBlockStartColor: var(--pf-t--global--border--color--high-contrast);\\n  background-color: var(--cem-pf-v6-c-page--BackgroundColor);\\n}\\n\\n@media screen and (max-width: calc(-1px + 48rem)) {\\n  :host {\\n    --cem-pf-v6-c-page__main-container--MarginInlineStart: var(--cem-pf-v6-c-page__main-container--xs--MarginInlineStart);\\n    --cem-pf-v6-c-page__main-container--MarginInlineEnd: var(--cem-pf-v6-c-page__main-container--xs--MarginInlineEnd);\\n    --cem-pf-v6-c-page__main-container--BorderBlockEndWidth: var(--cem-pf-v6-c-page__main-container--xs--BorderBlockEndWidth);\\n    --cem-pf-v6-c-page__main-container--BorderInlineStartWidth: var(--cem-pf-v6-c-page__main-container--xs--BorderInlineStartWidth);\\n    --cem-pf-v6-c-page__main-container--BorderInlineEndWidth: var(--cem-pf-v6-c-page__main-container--xs--BorderInlineEndWidth);\\n  }\\n}\\n\\n@media screen and (min-width: 75rem) {\\n  :host {\\n    --cem-pf-v6-c-page__sidebar--Width: var(--cem-pf-v6-c-page__sidebar--xl--Width);\\n    grid-template-areas:\\n      \\"header header\\"\\n      \\"sidebar main\\";\\n    grid-template-columns: var(--cem-pf-v6-c-page__sidebar--Width) 1fr;\\n    --cem-pf-v6-c-page__sidebar--TranslateX: var(--cem-pf-v6-c-page__sidebar--xl--TranslateX);\\n    --cem-pf-v6-c-page__sidebar--Opacity: var(--cem-pf-v6-c-page__sidebar--xl--Opacity);\\n    --cem-pf-v6-c-page__sidebar--BorderInlineEndWidth: 0;\\n  }\\n\\n  :host([sidebar-collapsed]) {\\n    grid-template-columns: 0 1fr;\\n\\n    \\u0026 ::slotted(cem-pf-v6-page-sidebar) {\\n      pointer-events: none;\\n    }\\n  }\\n\\n  :host(:not([sidebar-collapsed])) ::slotted(cem-pf-v6-page-main) {\\n    --cem-pf-v6-c-page__main-container--MarginInlineStart: 0;\\n  }\\n\\n  ::slotted(cem-pf-v6-masthead) {\\n    --cem-pf-v6-c-masthead--ColumnGap: var(--cem-pf-v6-c-masthead--m-display-inline--ColumnGap);\\n    --cem-pf-v6-c-masthead--GridTemplateColumns: var(--cem-pf-v6-c-masthead--m-display-inline--GridTemplateColumns);\\n    --cem-pf-v6-c-masthead__toggle--GridColumn: var(--cem-pf-v6-c-masthead--m-display-inline__toggle--GridColumn);\\n    --cem-pf-v6-c-masthead__brand--GridColumn: var(--cem-pf-v6-c-masthead--m-display-inline__brand--GridColumn);\\n    --cem-pf-v6-c-masthead__brand--Order: var(--cem-pf-v6-c-masthead--m-display-inline__brand--Order);\\n    --cem-pf-v6-c-masthead__brand--PaddingBlockEnd: var(--cem-pf-v6-c-masthead--m-display-inline__brand--PaddingBlockEnd);\\n    --cem-pf-v6-c-masthead__brand--BorderBlockEnd: var(--cem-pf-v6-c-masthead--m-display-inline__brand--BorderBlockEnd);\\n    --cem-pf-v6-c-masthead__main--GridColumn: var(--cem-pf-v6-c-masthead--m-display-inline__main--GridColumn);\\n    --cem-pf-v6-c-masthead__content--GridColumn: var(--cem-pf-v6-c-masthead--m-display-inline__content--GridColumn);\\n    --cem-pf-v6-c-masthead__content--Order: var(--cem-pf-v6-c-masthead--m-display-inline__content--Order);\\n    --cem-pf-v6-c-masthead__main--Display: var(--cem-pf-v6-c-masthead--m-display-inline__main--Display);\\n    --cem-pf-v6-c-masthead__main--ColumnGap: var(--cem-pf-v6-c-masthead--m-display-inline__main--ColumnGap);\\n  }\\n}\\n"'));
var cem_pf_v6_page_default = s38;

// elements/cem-pf-v6-page/cem-pf-v6-page.ts
var _sidebarCollapsed_dec, _a35, _PfV6Page_decorators, _init35, _sidebarCollapsed, _clickOutsideHandler, _onSidebarToggle, _PfV6Page_instances, updateSidebarState_fn;
_PfV6Page_decorators = [t3("cem-pf-v6-page")];
var _PfV6Page = class _PfV6Page extends (_a35 = i3, _sidebarCollapsed_dec = [n4({ type: Boolean, reflect: true, attribute: "sidebar-collapsed" })], _a35) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Page_instances);
    __privateAdd(this, _sidebarCollapsed, __runInitializers(_init35, 8, this, false)), __runInitializers(_init35, 11, this);
    __privateAdd(this, _clickOutsideHandler, (event) => {
      if (!_PfV6Page.match?.matches && !this.sidebarCollapsed) {
        const sidebar = this.querySelector("cem-pf-v6-page-sidebar");
        const mastheadToggle = this.querySelector("cem-pf-v6-masthead")?.shadowRoot?.getElementById("toggle-button");
        if (!event.composedPath().some((node) => node === sidebar || node === mastheadToggle)) {
          this.sidebarCollapsed = true;
          this.dispatchEvent(new SidebarToggleEvent(!this.sidebarCollapsed));
        }
      }
    });
    __privateAdd(this, _onSidebarToggle, (event) => {
      const toggleEvent = event;
      this.sidebarCollapsed = !toggleEvent.expanded;
    });
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("sidebar-toggle", __privateGet(this, _onSidebarToggle));
    document.body.addEventListener("click", __privateGet(this, _clickOutsideHandler));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("sidebar-toggle", __privateGet(this, _onSidebarToggle));
    document.body.removeEventListener("click", __privateGet(this, _clickOutsideHandler));
  }
  render() {
    return T`
      <slot name="skip-to-content"></slot>
      <slot name="masthead"></slot>
      <slot name="sidebar"></slot>
      <slot name="main"></slot>
    `;
  }
  updated(changed) {
    if (changed.has("sidebarCollapsed")) {
      __privateMethod(this, _PfV6Page_instances, updateSidebarState_fn).call(this);
    }
  }
};
_init35 = __decoratorStart(_a35);
_sidebarCollapsed = new WeakMap();
_clickOutsideHandler = new WeakMap();
_onSidebarToggle = new WeakMap();
_PfV6Page_instances = new WeakSet();
updateSidebarState_fn = function() {
  const isCollapsed = this.sidebarCollapsed;
  const sidebar = this.querySelector("cem-pf-v6-page-sidebar");
  sidebar?.toggleAttribute("collapsed", !!isCollapsed);
  sidebar?.toggleAttribute("expanded", !isCollapsed);
  const masthead = this.querySelector("cem-pf-v6-masthead");
  masthead?.toggleAttribute("sidebar-expanded", !isCollapsed);
};
__decorateElement(_init35, 4, "sidebarCollapsed", _sidebarCollapsed_dec, _PfV6Page, _sidebarCollapsed);
_PfV6Page = __decorateElement(_init35, 0, "PfV6Page", _PfV6Page_decorators, _PfV6Page);
__publicField(_PfV6Page, "match", globalThis.window?.matchMedia?.("(min-width: 75rem)"));
__publicField(_PfV6Page, "styles", cem_pf_v6_page_default);
__runInitializers(_init35, 1, _PfV6Page);
var PfV6Page = _PfV6Page;

// lit-css:elements/cem-pf-v6-page-main/cem-pf-v6-page-main.css
var s39 = new CSSStyleSheet();
s39.replaceSync(JSON.parse('":host {\\n  z-index: var(--cem-pf-v6-c-page__main-container--ZIndex);\\n  display: flex;\\n  flex-direction: column;\\n  grid-area: var(--cem-pf-v6-c-page__main-container--GridArea);\\n  align-self: var(--cem-pf-v6-c-page__main-container--AlignSelf);\\n  max-height: var(--cem-pf-v6-c-page__main-container--MaxHeight);\\n  margin-inline-start: var(--cem-pf-v6-c-page__main-container--MarginInlineStart);\\n  margin-inline-end: var(--cem-pf-v6-c-page__main-container--MarginInlineEnd);\\n  overflow-x: hidden;\\n  overflow-y: auto;\\n  -webkit-overflow-scrolling: touch;\\n  background: var(--cem-pf-v6-c-page__main-container--BackgroundColor);\\n  border: solid var(--cem-pf-v6-c-page__main-container--BorderColor);\\n  border-block-start-width: var(--cem-pf-v6-c-page__main-container--BorderBlockStartWidth);\\n  border-block-end-width: var(--cem-pf-v6-c-page__main-container--BorderBlockEndWidth);\\n  border-inline-start-width: var(--cem-pf-v6-c-page__main-container--BorderInlineStartWidth);\\n  border-inline-end-width: var(--cem-pf-v6-c-page__main-container--BorderInlineEndWidth);\\n  border-radius: var(--cem-pf-v6-c-page__main-container--BorderRadius);\\n}\\n\\nslot {\\n  display: flex;\\n  flex: 1;\\n  flex-direction: column;\\n  outline: 0;\\n}\\n"'));
var cem_pf_v6_page_main_default = s39;

// elements/cem-pf-v6-page-main/cem-pf-v6-page-main.ts
var _PfV6PageMain_decorators, _init36, _a36;
_PfV6PageMain_decorators = [t3("cem-pf-v6-page-main")];
var PfV6PageMain = class extends (_a36 = i3) {
  static styles = cem_pf_v6_page_main_default;
  #internals = this.attachInternals();
  constructor() {
    super();
    this.#internals.role = "main";
  }
  render() {
    return T`<slot></slot>`;
  }
};
_init36 = __decoratorStart(_a36);
PfV6PageMain = __decorateElement(_init36, 0, "PfV6PageMain", _PfV6PageMain_decorators, PfV6PageMain);
__runInitializers(_init36, 1, PfV6PageMain);

// lit-css:elements/cem-pf-v6-page-sidebar/cem-pf-v6-page-sidebar.css
var s40 = new CSSStyleSheet();
s40.replaceSync(JSON.parse('":host {\\n  z-index: var(--cem-pf-v6-c-page__sidebar--ZIndex);\\n  display: flex;\\n  flex-direction: column;\\n  grid-area: sidebar;\\n  grid-row-start: 2;\\n  grid-column-start: 1;\\n  width: var(--cem-pf-v6-c-page__sidebar--Width);\\n  padding-block-start: 0;\\n  padding-block-end: var(--cem-pf-v6-c-page__sidebar--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-page__sidebar--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-page__sidebar--PaddingInlineEnd);\\n  margin-inline-end: var(--cem-pf-v6-c-page__sidebar--MarginInlineEnd);\\n  overflow-x: hidden;\\n  overflow-y: auto;\\n  -webkit-overflow-scrolling: touch;\\n  background-color: var(--cem-pf-v6-c-page__sidebar--BackgroundColor);\\n  border-inline-end: var(--cem-pf-v6-c-page__sidebar--BorderInlineEndWidth) solid var(--cem-pf-v6-c-page__sidebar--BorderInlineEndColor);\\n  opacity: var(--cem-pf-v6-c-page__sidebar--Opacity);\\n  transition: var(--cem-pf-v6-c-page__sidebar--TransitionProperty) var(--cem-pf-v6-c-page__sidebar--TransitionDuration) var(--cem-pf-v6-c-page__sidebar--TransitionTimingFunction);\\n  transform: translateX(var(--cem-pf-v6-c-page__sidebar--TranslateX)) translateZ(var(--cem-pf-v6-c-page__sidebar--TranslateZ));\\n}\\n\\n#body {\\n  padding-inline-start: var(--cem-pf-v6-c-page__sidebar-body--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-page__sidebar-body--PaddingInlineEnd);\\n  flex-grow: 1;\\n}\\n\\n:host([expanded]) {\\n  --cem-pf-v6-c-page__sidebar--Opacity: var(--cem-pf-v6-c-page__sidebar--m-expanded--Opacity);\\n  --cem-pf-v6-c-page__sidebar--TranslateX: var(--cem-pf-v6-c-page__sidebar--m-expanded--TranslateX);\\n  box-shadow: var(--cem-pf-v6-c-page__sidebar--BoxShadow);\\n\\n  @media screen and (min-width: 75rem) {\\n    --cem-pf-v6-c-page__sidebar--BoxShadow: 0;\\n  }\\n}\\n\\n:host([collapsed]) {\\n  --cem-pf-v6-c-page__sidebar--Opacity: 0;\\n  --cem-pf-v6-c-page__sidebar--TranslateX: -100%;\\n  pointer-events: none;\\n}\\n"'));
var cem_pf_v6_page_sidebar_default = s40;

// elements/cem-pf-v6-page-sidebar/cem-pf-v6-page-sidebar.ts
var _expanded_dec11, _collapsed_dec, _a37, _PfV6PageSidebar_decorators, _internals7, _init37, _collapsed, _expanded11;
_PfV6PageSidebar_decorators = [t3("cem-pf-v6-page-sidebar")];
var PfV6PageSidebar = class extends (_a37 = i3, _collapsed_dec = [n4({ type: Boolean, reflect: true })], _expanded_dec11 = [n4({ type: Boolean, reflect: true })], _a37) {
  constructor() {
    super();
    __privateAdd(this, _internals7, this.attachInternals());
    __privateAdd(this, _collapsed, __runInitializers(_init37, 8, this, false)), __runInitializers(_init37, 11, this);
    __privateAdd(this, _expanded11, __runInitializers(_init37, 12, this, false)), __runInitializers(_init37, 15, this);
    __privateGet(this, _internals7).role = "navigation";
  }
  updated(changed) {
    if (changed.has("collapsed")) {
      this.expanded = !this.collapsed;
    } else if (changed.has("expanded")) {
      this.collapsed = !this.expanded;
    }
  }
  render() {
    return T`
      <div id="body">
        <slot></slot>
      </div>
    `;
  }
};
_init37 = __decoratorStart(_a37);
_internals7 = new WeakMap();
_collapsed = new WeakMap();
_expanded11 = new WeakMap();
__decorateElement(_init37, 4, "collapsed", _collapsed_dec, PfV6PageSidebar, _collapsed);
__decorateElement(_init37, 4, "expanded", _expanded_dec11, PfV6PageSidebar, _expanded11);
PfV6PageSidebar = __decorateElement(_init37, 0, "PfV6PageSidebar", _PfV6PageSidebar_decorators, PfV6PageSidebar);
__publicField(PfV6PageSidebar, "styles", cem_pf_v6_page_sidebar_default);
__runInitializers(_init37, 1, PfV6PageSidebar);

// lit-css:elements/cem-pf-v6-popover/cem-pf-v6-popover.css
var s41 = new CSSStyleSheet();
s41.replaceSync(JSON.parse('":host {\\n  display: inline-block;\\n  position: relative;\\n  align-self: start;\\n  justify-self: start;\\n\\n  --distance: 8px;\\n  --min-width: 100px;\\n  --max-width: 300px;\\n\\n  --cem-pf-v6-c-popover--FontSize: var(--pf-t--global--font--size--body--sm);\\n  --cem-pf-v6-c-popover--BoxShadow: var(--pf-t--global--box-shadow--md);\\n  --cem-pf-v6-c-popover--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-popover--BorderColor: var(--pf-t--global--border--color--high-contrast);\\n  --cem-pf-v6-c-popover--BorderRadius: var(--pf-t--global--border--radius--medium);\\n  --cem-pf-v6-c-popover__content--BackgroundColor: var(--pf-t--global--background--color--floating--default);\\n  --cem-pf-v6-c-popover__content--PaddingBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__content--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__content--PaddingBlockEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__content--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__close--InsetBlockStart: calc(var(--cem-pf-v6-c-popover__content--PaddingBlockStart) - var(--pf-t--global--spacer--control--vertical--default));\\n  --cem-pf-v6-c-popover__close--InsetInlineEnd: var(--cem-pf-v6-c-popover__content--PaddingInlineEnd);\\n  --cem-pf-v6-c-popover__close--sibling--PaddingInlineEnd: var(--pf-t--global--spacer--2xl);\\n  --cem-pf-v6-c-popover__header--MarginBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-popover__title-text--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-popover__title-text--FontWeight: var(--pf-t--global--font--weight--body--bold);\\n  --cem-pf-v6-c-popover__title-text--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-popover__footer--MarginBlockStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-popover__arrow--Width: 1.5rem;\\n  --cem-pf-v6-c-popover__arrow--Height: 1.5rem;\\n  --cem-pf-v6-c-popover__arrow--BackgroundColor: var(--cem-pf-v6-c-popover__content--BackgroundColor);\\n  --cem-pf-v6-c-popover__arrow--BoxShadow: var(--cem-pf-v6-c-popover--BoxShadow);\\n}\\n\\n#trigger {\\n  anchor-name: --trigger;\\n  display: inline-block;\\n}\\n\\n#content {\\n  background-color: transparent;\\n  padding: 0;\\n  overflow: initial;\\n  position: fixed;\\n  position-anchor: --trigger;\\n  min-width: var(--min-width);\\n  max-width: var(--max-width);\\n  margin: 0;\\n  font-size: var(--cem-pf-v6-c-popover--FontSize);\\n  border: var(--cem-pf-v6-c-popover--BorderWidth) solid var(--cem-pf-v6-c-popover--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-popover--BorderRadius);\\n  box-shadow: var(--cem-pf-v6-c-popover--BoxShadow);\\n  top: auto;\\n  bottom: auto;\\n  left: auto;\\n  right: auto;\\n}\\n\\n#popover {\\n  position: relative;\\n  padding-block-start: var(--cem-pf-v6-c-popover__content--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-popover__content--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-popover__content--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-popover__content--PaddingInlineEnd);\\n  background-color: var(--cem-pf-v6-c-popover__content--BackgroundColor);\\n  border-radius: var(--cem-pf-v6-c-popover--BorderRadius);\\n}\\n\\n#arrow {\\n  position: absolute;\\n  top: var(--cem-pf-v6-c-popover__arrow--InsetBlockStart, auto);\\n  right: var(--cem-pf-v6-c-popover__arrow--InsetInlineEnd, auto);\\n  bottom: var(--cem-pf-v6-c-popover__arrow--InsetBlockEnd, auto);\\n  left: var(--cem-pf-v6-c-popover__arrow--InsetInlineStart, auto);\\n  width: var(--cem-pf-v6-c-popover__arrow--Width);\\n  height: var(--cem-pf-v6-c-popover__arrow--Height);\\n  pointer-events: none;\\n  background-color: var(--cem-pf-v6-c-popover__arrow--BackgroundColor);\\n  border: var(--cem-pf-v6-c-popover--BorderWidth) solid var(--cem-pf-v6-c-popover--BorderColor);\\n  box-shadow: var(--cem-pf-v6-c-popover__arrow--BoxShadow);\\n  transform: translateX(var(--cem-pf-v6-c-popover__arrow--TranslateX, 0)) translateY(var(--cem-pf-v6-c-popover__arrow--TranslateY, 0)) rotate(var(--cem-pf-v6-c-popover__arrow--Rotate, 0));\\n}\\n\\n#close {\\n  position: absolute;\\n  inset-block-start: var(--cem-pf-v6-c-popover__close--InsetBlockStart);\\n  inset-inline-end: var(--cem-pf-v6-c-popover__close--InsetInlineEnd);\\n  padding: var(--pf-t--global--spacer--control--vertical--default) var(--pf-t--global--spacer--control--horizontal--plain--default);\\n  background: transparent;\\n  border: none;\\n  cursor: pointer;\\n  color: var(--pf-t--global--icon--color--regular);\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n\\n  \\u0026:hover {\\n    background-color: var(--pf-t--global--background--color--action--plain--hover);\\n  }\\n}\\n\\n:host(:not([closeable])) #close,\\n:host([closeable=\\"false\\"]) #close {\\n  display: none;\\n}\\n\\n:host([closeable]) #header,\\n:host([closeable=\\"\\"]) #header {\\n  padding-inline-end: var(--cem-pf-v6-c-popover__close--sibling--PaddingInlineEnd);\\n}\\n\\n#header {\\n  margin-block-end: 0;\\n\\n  \\u0026.has-content {\\n    margin-block-end: var(--cem-pf-v6-c-popover__header--MarginBlockEnd);\\n  }\\n}\\n\\n#header-content {\\n  display: flex;\\n  flex: 0 0 auto;\\n\\n  ::slotted(h1),\\n  ::slotted(h2),\\n  ::slotted(h3),\\n  ::slotted(h4),\\n  ::slotted(h5),\\n  ::slotted(h6) {\\n    min-width: 0 !important;\\n    font-size: var(--cem-pf-v6-c-popover__title-text--FontSize) !important;\\n    font-weight: var(--cem-pf-v6-c-popover__title-text--FontWeight) !important;\\n    color: var(--cem-pf-v6-c-popover__title-text--Color) !important;\\n    overflow-wrap: break-word !important;\\n    margin: 0 !important;\\n  }\\n}\\n\\n#body {\\n  word-wrap: break-word;\\n}\\n\\n#footer {\\n  margin-block-start: 0;\\n\\n  \\u0026.has-content {\\n    margin-block-start: var(--cem-pf-v6-c-popover__footer--MarginBlockStart);\\n  }\\n}\\n\\n:host([has-auto-width]) #content {\\n  min-width: unset;\\n  max-width: unset;\\n  width: max-content;\\n}\\n\\n/* Position: top (default) */\\n#content.position-top {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-bottom: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: 50%;\\n  --cem-pf-v6-c-popover__arrow--TranslateX: -50%;\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom {\\n  top: anchor(--trigger bottom);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-top: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: 50%;\\n  --cem-pf-v6-c-popover__arrow--TranslateX: -50%;\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left {\\n  top: anchor(--trigger center);\\n  right: anchor(--trigger left);\\n  translate: 0 -50%;\\n  margin-right: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: 50%;\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--TranslateY: -50%;\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right {\\n  top: anchor(--trigger center);\\n  left: anchor(--trigger right);\\n  translate: 0 -50%;\\n  margin-left: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: 50%;\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--TranslateY: -50%;\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-top-start {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger left);\\n  margin-bottom: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-top-end {\\n  bottom: anchor(--trigger top);\\n  right: anchor(--trigger right);\\n  margin-bottom: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom-start {\\n  top: anchor(--trigger bottom);\\n  left: anchor(--trigger left);\\n  margin-top: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-bottom-end {\\n  top: anchor(--trigger bottom);\\n  right: anchor(--trigger right);\\n  margin-top: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: calc(var(--cem-pf-v6-c-popover__arrow--Height) / -2);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left-start {\\n  top: anchor(--trigger top);\\n  right: anchor(--trigger left);\\n  margin-right: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-left-end {\\n  bottom: anchor(--trigger bottom);\\n  right: anchor(--trigger left);\\n  margin-right: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineEnd: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right-start {\\n  top: anchor(--trigger top);\\n  left: anchor(--trigger right);\\n  margin-left: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockStart: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-right-end {\\n  bottom: anchor(--trigger bottom);\\n  left: anchor(--trigger right);\\n  margin-left: var(--distance);\\n  --cem-pf-v6-c-popover__arrow--InsetBlockEnd: var(--pf-t--global--spacer--lg);\\n  --cem-pf-v6-c-popover__arrow--InsetInlineStart: calc(var(--cem-pf-v6-c-popover__arrow--Width) / -2);\\n  --cem-pf-v6-c-popover__arrow--Rotate: 45deg;\\n}\\n\\n#content.position-auto {\\n  bottom: anchor(--trigger top);\\n  left: anchor(--trigger center);\\n  translate: -50% 0;\\n  margin-bottom: var(--distance);\\n}\\n"'));
var cem_pf_v6_popover_default = s41;

// elements/cem-pf-v6-popover/cem-pf-v6-popover.ts
var PfPopoverShowEvent = class extends Event {
  open = true;
  constructor() {
    super("pf-popover-show", { bubbles: true });
  }
};
var PfPopoverHideEvent = class extends Event {
  open = false;
  constructor() {
    super("pf-popover-hide", { bubbles: true });
  }
};
var supportsAnchorPositioning = globalThis.CSS?.supports?.("anchor-name: --test") ?? false;
var _hasAutoWidth_dec, _maxWidth_dec, _minWidth_dec, _closeButtonLabel_dec, _closeable_dec, _triggerAction_dec, _distance_dec, _position_dec2, _open_dec3, _a38, _PfV6Popover_decorators, _init38, _open3, _position2, _distance, _triggerAction, _closeable, _closeButtonLabel, _minWidth, _maxWidth, _hasAutoWidth, _hoverShowTimeout, _hoverHideTimeout, _hasHeaderContent2, _hasFooterContent, _PfV6Popover_instances, contentStyle_get, showPopover_fn, hidePopover_fn, _handleTriggerClick, _handlePointerEnter, _handlePointerLeave, _handlePopoverPointerEnter, _handlePopoverPointerLeave, _handlePopoverToggle, _handleClose, onHeaderSlotChange_fn, onFooterSlotChange_fn;
_PfV6Popover_decorators = [t3("cem-pf-v6-popover")];
var PfV6Popover = class extends (_a38 = i3, _open_dec3 = [n4({ type: Boolean, reflect: true })], _position_dec2 = [n4({ reflect: true })], _distance_dec = [n4({ type: Number, reflect: true })], _triggerAction_dec = [n4({ reflect: true, attribute: "trigger-action" })], _closeable_dec = [n4({ type: Boolean, reflect: true })], _closeButtonLabel_dec = [n4({ attribute: "close-button-label" })], _minWidth_dec = [n4({ attribute: "min-width" })], _maxWidth_dec = [n4({ attribute: "max-width" })], _hasAutoWidth_dec = [n4({ type: Boolean, reflect: true, attribute: "has-auto-width" })], _a38) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Popover_instances);
    __privateAdd(this, _open3, __runInitializers(_init38, 8, this, false)), __runInitializers(_init38, 11, this);
    __privateAdd(this, _position2, __runInitializers(_init38, 12, this, "top")), __runInitializers(_init38, 15, this);
    __privateAdd(this, _distance, __runInitializers(_init38, 16, this, 8)), __runInitializers(_init38, 19, this);
    __privateAdd(this, _triggerAction, __runInitializers(_init38, 20, this, "click")), __runInitializers(_init38, 23, this);
    __privateAdd(this, _closeable, __runInitializers(_init38, 24, this, true)), __runInitializers(_init38, 27, this);
    __privateAdd(this, _closeButtonLabel, __runInitializers(_init38, 28, this, "Close")), __runInitializers(_init38, 31, this);
    __privateAdd(this, _minWidth, __runInitializers(_init38, 32, this)), __runInitializers(_init38, 35, this);
    __privateAdd(this, _maxWidth, __runInitializers(_init38, 36, this)), __runInitializers(_init38, 39, this);
    __privateAdd(this, _hasAutoWidth, __runInitializers(_init38, 40, this, false)), __runInitializers(_init38, 43, this);
    __privateAdd(this, _hoverShowTimeout);
    __privateAdd(this, _hoverHideTimeout);
    __privateAdd(this, _hasHeaderContent2, false);
    __privateAdd(this, _hasFooterContent, false);
    __privateAdd(this, _handleTriggerClick, (e6) => {
      const triggerSlot = this.shadowRoot?.querySelector('slot[name="trigger"]');
      const assigned = triggerSlot?.assignedElements();
      if (assigned && assigned.length > 0) {
        e6.stopPropagation();
        this.toggle();
      }
    });
    __privateAdd(this, _handlePointerEnter, () => {
      clearTimeout(__privateGet(this, _hoverHideTimeout));
      __privateSet(this, _hoverShowTimeout, setTimeout(() => {
        __privateMethod(this, _PfV6Popover_instances, showPopover_fn).call(this);
      }, 150));
    });
    __privateAdd(this, _handlePointerLeave, () => {
      clearTimeout(__privateGet(this, _hoverShowTimeout));
      __privateSet(this, _hoverHideTimeout, setTimeout(() => {
        __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
      }, 300));
    });
    __privateAdd(this, _handlePopoverPointerEnter, () => {
      clearTimeout(__privateGet(this, _hoverHideTimeout));
    });
    __privateAdd(this, _handlePopoverPointerLeave, () => {
      __privateSet(this, _hoverHideTimeout, setTimeout(() => {
        __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
      }, 300));
    });
    __privateAdd(this, _handlePopoverToggle, (e6) => {
      const isOpen = e6.newState === "open";
      if (isOpen !== this.open) {
        this.open = isOpen;
      }
      const triggerSlot = this.shadowRoot?.querySelector('slot[name="trigger"]');
      const triggerButton = triggerSlot?.assignedElements()?.[0];
      if (triggerButton) {
        triggerButton.setAttribute("aria-expanded", String(isOpen));
      }
      this.dispatchEvent(isOpen ? new PfPopoverShowEvent() : new PfPopoverHideEvent());
    });
    __privateAdd(this, _handleClose, (e6) => {
      e6.stopPropagation();
      __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
    });
  }
  render() {
    return T`
      <div id="trigger"
           @click=${this.triggerAction === "click" ? __privateGet(this, _handleTriggerClick) : A}
           @pointerenter=${this.triggerAction === "hover" ? __privateGet(this, _handlePointerEnter) : A}
           @pointerleave=${this.triggerAction === "hover" ? __privateGet(this, _handlePointerLeave) : A}>
        <slot name="trigger"></slot>
      </div>

      <div id="content"
           class="position-${this.position}"
           popover="manual"
           part="popover"
           style=${__privateGet(this, _PfV6Popover_instances, contentStyle_get)}
           @pointerenter=${this.triggerAction === "hover" ? __privateGet(this, _handlePopoverPointerEnter) : A}
           @pointerleave=${this.triggerAction === "hover" ? __privateGet(this, _handlePopoverPointerLeave) : A}
           @toggle=${__privateGet(this, _handlePopoverToggle)}>
        <div id="arrow"></div>
        <div id="popover"
             part="content">
          <button id="close"
                  part="close-button"
                  aria-label=${this.closeButtonLabel}
                  @click=${__privateGet(this, _handleClose)}>
            <svg class="cem-pf-v6-svg"
                 viewBox="0 0 352 512"
                 fill="currentColor"
                 role="presentation"
                 width="1em"
                 height="1em">
              <path d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path>
            </svg>
          </button>

          <header id="header"
                  class=${__privateGet(this, _hasHeaderContent2) ? "has-content" : ""}
                  part="header">
            <div id="title">
              <div id="header-content">
                <slot name="header"
                      @slotchange=${__privateMethod(this, _PfV6Popover_instances, onHeaderSlotChange_fn)}></slot>
              </div>
            </div>
          </header>

          <div id="body"
               part="body">
            <slot></slot>
          </div>

          <footer id="footer"
                  class=${__privateGet(this, _hasFooterContent) ? "has-content" : ""}
                  part="footer">
            <slot name="footer"
                  @slotchange=${__privateMethod(this, _PfV6Popover_instances, onFooterSlotChange_fn)}></slot>
          </footer>
        </div>
      </div>
    `;
  }
  updated(changed) {
    if (changed.has("open")) {
      const contentEl = this.shadowRoot?.getElementById("content");
      if (!contentEl) return;
      if (this.open && !contentEl.matches(":popover-open")) {
        __privateMethod(this, _PfV6Popover_instances, showPopover_fn).call(this);
      } else if (!this.open && contentEl.matches(":popover-open")) {
        __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
      }
    }
  }
  toggle() {
    if (this.open) {
      __privateMethod(this, _PfV6Popover_instances, hidePopover_fn).call(this);
    } else {
      __privateMethod(this, _PfV6Popover_instances, showPopover_fn).call(this);
    }
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    clearTimeout(__privateGet(this, _hoverShowTimeout));
    clearTimeout(__privateGet(this, _hoverHideTimeout));
  }
};
_init38 = __decoratorStart(_a38);
_open3 = new WeakMap();
_position2 = new WeakMap();
_distance = new WeakMap();
_triggerAction = new WeakMap();
_closeable = new WeakMap();
_closeButtonLabel = new WeakMap();
_minWidth = new WeakMap();
_maxWidth = new WeakMap();
_hasAutoWidth = new WeakMap();
_hoverShowTimeout = new WeakMap();
_hoverHideTimeout = new WeakMap();
_hasHeaderContent2 = new WeakMap();
_hasFooterContent = new WeakMap();
_PfV6Popover_instances = new WeakSet();
contentStyle_get = function() {
  const parts = [];
  if (this.minWidth) parts.push(`--min-width: ${this.minWidth}`);
  if (this.maxWidth) parts.push(`--max-width: ${this.maxWidth}`);
  parts.push(`--distance: ${this.distance}px`);
  return parts.join("; ");
};
showPopover_fn = async function() {
  const contentEl = this.shadowRoot?.getElementById("content");
  if (!contentEl) return;
  try {
    contentEl.showPopover();
    if (!supportsAnchorPositioning) {
      await new Promise((resolve) => requestAnimationFrame(resolve));
      const trigger = this.shadowRoot?.getElementById("trigger");
      if (!trigger) return;
      const triggerRect = trigger.getBoundingClientRect();
      const popoverRect = contentEl.getBoundingClientRect();
      const dist = this.distance;
      let top, left;
      switch (this.position) {
        case "top":
          top = triggerRect.top - popoverRect.height - dist;
          left = triggerRect.left + triggerRect.width / 2 - popoverRect.width / 2;
          break;
        case "top-start":
          top = triggerRect.top - popoverRect.height - dist;
          left = triggerRect.left;
          break;
        case "top-end":
          top = triggerRect.top - popoverRect.height - dist;
          left = triggerRect.right - popoverRect.width;
          break;
        case "bottom":
          top = triggerRect.bottom + dist;
          left = triggerRect.left + triggerRect.width / 2 - popoverRect.width / 2;
          break;
        case "bottom-start":
          top = triggerRect.bottom + dist;
          left = triggerRect.left;
          break;
        case "bottom-end":
          top = triggerRect.bottom + dist;
          left = triggerRect.right - popoverRect.width;
          break;
        case "left":
          top = triggerRect.top + triggerRect.height / 2 - popoverRect.height / 2;
          left = triggerRect.left - popoverRect.width - dist;
          break;
        case "left-start":
          top = triggerRect.top;
          left = triggerRect.left - popoverRect.width - dist;
          break;
        case "left-end":
          top = triggerRect.bottom - popoverRect.height;
          left = triggerRect.left - popoverRect.width - dist;
          break;
        case "right":
          top = triggerRect.top + triggerRect.height / 2 - popoverRect.height / 2;
          left = triggerRect.right + dist;
          break;
        case "right-start":
          top = triggerRect.top;
          left = triggerRect.right + dist;
          break;
        case "right-end":
          top = triggerRect.bottom - popoverRect.height;
          left = triggerRect.right + dist;
          break;
        default:
          top = triggerRect.top - popoverRect.height - dist;
          left = triggerRect.left + triggerRect.width / 2 - popoverRect.width / 2;
      }
      contentEl.style.top = `${top}px`;
      contentEl.style.left = `${left}px`;
      contentEl.style.bottom = "auto";
      contentEl.style.right = "auto";
      contentEl.style.translate = "none";
    }
  } catch (e6) {
    console.warn("Failed to show popover:", e6);
  }
};
hidePopover_fn = function() {
  const contentEl = this.shadowRoot?.getElementById("content");
  try {
    contentEl?.hidePopover();
  } catch (e6) {
    console.warn("Failed to hide popover:", e6);
  }
};
_handleTriggerClick = new WeakMap();
_handlePointerEnter = new WeakMap();
_handlePointerLeave = new WeakMap();
_handlePopoverPointerEnter = new WeakMap();
_handlePopoverPointerLeave = new WeakMap();
_handlePopoverToggle = new WeakMap();
_handleClose = new WeakMap();
onHeaderSlotChange_fn = function(e6) {
  const slot = e6.target;
  __privateSet(this, _hasHeaderContent2, slot.assignedNodes().length > 0);
  this.requestUpdate();
};
onFooterSlotChange_fn = function(e6) {
  const slot = e6.target;
  __privateSet(this, _hasFooterContent, slot.assignedNodes().length > 0);
  this.requestUpdate();
};
__decorateElement(_init38, 4, "open", _open_dec3, PfV6Popover, _open3);
__decorateElement(_init38, 4, "position", _position_dec2, PfV6Popover, _position2);
__decorateElement(_init38, 4, "distance", _distance_dec, PfV6Popover, _distance);
__decorateElement(_init38, 4, "triggerAction", _triggerAction_dec, PfV6Popover, _triggerAction);
__decorateElement(_init38, 4, "closeable", _closeable_dec, PfV6Popover, _closeable);
__decorateElement(_init38, 4, "closeButtonLabel", _closeButtonLabel_dec, PfV6Popover, _closeButtonLabel);
__decorateElement(_init38, 4, "minWidth", _minWidth_dec, PfV6Popover, _minWidth);
__decorateElement(_init38, 4, "maxWidth", _maxWidth_dec, PfV6Popover, _maxWidth);
__decorateElement(_init38, 4, "hasAutoWidth", _hasAutoWidth_dec, PfV6Popover, _hasAutoWidth);
PfV6Popover = __decorateElement(_init38, 0, "PfV6Popover", _PfV6Popover_decorators, PfV6Popover);
__publicField(PfV6Popover, "styles", cem_pf_v6_popover_default);
__runInitializers(_init38, 1, PfV6Popover);

// elements/node_modules/lit-html/node/directives/style-map.js
var n5 = "important";
var i5 = " !" + n5;
var o8 = e3(class extends i4 {
  constructor(t6) {
    if (super(t6), t6.type !== t4.ATTRIBUTE || "style" !== t6.name || t6.strings?.length > 2) throw Error("The `styleMap` directive must be used in the `style` attribute and must be the only part in the attribute.");
  }
  render(t6) {
    return Object.keys(t6).reduce((e6, r7) => {
      const s58 = t6[r7];
      return null == s58 ? e6 : e6 + `${r7 = r7.includes("-") ? r7 : r7.replace(/(?:^(webkit|moz|ms|o)|)(?=[A-Z])/g, "-$&").toLowerCase()}:${s58};`;
    }, "");
  }
  update(e6, [r7]) {
    const { style: s58 } = e6.element;
    if (void 0 === this.ft) return this.ft = new Set(Object.keys(r7)), this.render(r7);
    for (const t6 of this.ft) null == r7[t6] && (this.ft.delete(t6), t6.includes("-") ? s58.removeProperty(t6) : s58[t6] = null);
    for (const t6 in r7) {
      const e7 = r7[t6];
      if (null != e7) {
        this.ft.add(t6);
        const r8 = "string" == typeof e7 && e7.endsWith(i5);
        t6.includes("-") || r8 ? s58.setProperty(t6, r8 ? e7.slice(0, -11) : e7, r8 ? n5 : "") : s58[t6] = e7;
      }
    }
    return E;
  }
});

// lit-css:elements/cem-pf-v6-progress/cem-pf-v6-progress.css
var s42 = new CSSStyleSheet();
s42.replaceSync(JSON.parse("\"*,\\n*::before,\\n*::after {\\n  box-sizing: border-box;\\n}\\n\\n:host {\\n  display: block;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n#container {\\n  display: grid;\\n  grid-template-rows: 1fr auto;\\n  grid-template-columns: auto auto;\\n  /** Gap between the description, bar, and helper-text rows. Defaults to `--pf-t--global--spacer--md`. */\\n  gap: var(--pf-v6-c-progress--GridGap, var(--pf-t--global--spacer--md, 1rem));\\n  align-items: end;\\n\\n  /* Size modifiers */\\n\\n  \\u0026.sm {\\n    /** Height of the bar at small size. Defaults to `--pf-t--global--spacer--sm`. */\\n    --_bar-height: var(--pf-v6-c-progress--m-sm__bar--Height,\\n      var(--pf-t--global--spacer--sm, 0.5rem));\\n\\n    \\u0026 #measure {\\n      /** Font size of the measure text at small size. Defaults to `--pf-t--global--font--size--body--sm`. */\\n      font-size: var(--pf-v6-c-progress--m-sm__measure--FontSize,\\n        var(--pf-t--global--font--size--body--sm, 0.75rem));\\n    }\\n  }\\n\\n  \\u0026.lg {\\n    /** Height of the bar at large size. Defaults to `--pf-t--global--spacer--lg`. */\\n    --_bar-height: var(--pf-v6-c-progress--m-lg__bar--Height,\\n      var(--pf-t--global--spacer--lg, 1.5rem));\\n  }\\n\\n  /* Measure location modifiers */\\n\\n  \\u0026.outside,\\n  \\u0026.singleline {\\n    grid-template-columns: 1fr fit-content(50%);\\n  }\\n\\n  \\u0026.outside {\\n    \\u0026 #description {\\n      grid-column: 1 / 3;\\n    }\\n\\n    \\u0026 #status {\\n      grid-row: 2 / 3;\\n      grid-column: 2 / 3;\\n      align-self: center;\\n    }\\n\\n    \\u0026 #measure {\\n      display: inline-block;\\n      /** Font size of the measure text in outside layout. Defaults to `--pf-t--global--font--size--sm`. */\\n      font-size: var(--pf-v6-c-progress--m-outside__measure--FontSize,\\n        var(--pf-t--global--font--size--sm, 0.875rem));\\n    }\\n\\n    \\u0026 #bar {\\n      grid-column: 1 / 2;\\n    }\\n  }\\n\\n  \\u0026.singleline {\\n    grid-template-rows: 1fr;\\n\\n    \\u0026 #description {\\n      display: none;\\n    }\\n\\n    \\u0026 #bar {\\n      grid-row: 1 / 2;\\n      grid-column: 1 / 2;\\n    }\\n\\n    \\u0026 #status {\\n      grid-row: 1 / 2;\\n      grid-column: 2 / 3;\\n    }\\n  }\\n\\n  \\u0026.inside {\\n    \\u0026 #indicator {\\n      display: flex;\\n      align-items: center;\\n      justify-content: center;\\n      /** Minimum width of the indicator when measure is inside. Defaults to `--pf-t--global--spacer--xl`. */\\n      min-width: var(--pf-v6-c-progress--m-inside__indicator--MinWidth,\\n        var(--pf-t--global--spacer--xl, 2rem));\\n    }\\n\\n    \\u0026 #measure {\\n      /** Font size of the measure text inside the bar. Defaults to `--pf-t--global--font--size--sm`. */\\n      font-size: var(--pf-v6-c-progress--m-inside__measure--FontSize,\\n        var(--pf-t--global--font--size--sm, 0.875rem));\\n      /** Color of the measure text inside the bar. Defaults to `--pf-t--global--text--color--on-brand--default`. */\\n      color: var(--_inside-measure-color,\\n        var(--pf-v6-c-progress--m-inside__measure--Color,\\n          var(--pf-t--global--text--color--on-brand--default, #fff)));\\n      text-align: center;\\n    }\\n  }\\n\\n  /* Variant modifiers */\\n\\n  \\u0026.success {\\n    /** Indicator color for the success variant. Defaults to `--pf-t--global--color--status--success--default`. */\\n    --_indicator-bg: var(--pf-v6-c-progress--m-success__indicator--BackgroundColor,\\n      var(--pf-t--global--color--status--success--default, #3e8635));\\n    /** Status icon color for the success variant. Defaults to `--pf-t--global--icon--color--status--success--default`. */\\n    --_status-icon-color: var(--pf-v6-c-progress--m-success__status-icon--Color,\\n      var(--pf-t--global--icon--color--status--success--default, #3e8635));\\n    /** Inside measure text color for the success variant. Defaults to `--pf-t--global--text--color--status--on-success--default`. */\\n    --_inside-measure-color: var(--pf-v6-c-progress--m-success--m-inside__measure--Color,\\n      var(--pf-t--global--text--color--status--on-success--default, #fff));\\n  }\\n\\n  \\u0026.warning {\\n    /** Indicator color for the warning variant. Defaults to `--pf-t--global--color--status--warning--default`. */\\n    --_indicator-bg: var(--pf-v6-c-progress--m-warning__indicator--BackgroundColor,\\n      var(--pf-t--global--color--status--warning--default, #f0ab00));\\n    /** Status icon color for the warning variant. Defaults to `--pf-t--global--icon--color--status--warning--default`. */\\n    --_status-icon-color: var(--pf-v6-c-progress--m-warning__status-icon--Color,\\n      var(--pf-t--global--icon--color--status--warning--default, #f0ab00));\\n    /** Inside measure text color for the warning variant. Defaults to `--pf-t--global--text--color--status--on-warning--default`. */\\n    --_inside-measure-color: var(--pf-v6-c-progress--m-warning--m-inside__measure--Color,\\n      var(--pf-t--global--text--color--status--on-warning--default, #151515));\\n  }\\n\\n  \\u0026.danger {\\n    /** Indicator color for the danger variant. Defaults to `--pf-t--global--color--status--danger--default`. */\\n    --_indicator-bg: var(--pf-v6-c-progress--m-danger__indicator--BackgroundColor,\\n      var(--pf-t--global--color--status--danger--default, #c9190b));\\n    /** Status icon color for the danger variant. Defaults to `--pf-t--global--icon--color--status--danger--default`. */\\n    --_status-icon-color: var(--pf-v6-c-progress--m-danger__status-icon--Color,\\n      var(--pf-t--global--icon--color--status--danger--default, #c9190b));\\n    /** Inside measure text color for the danger variant. Defaults to `--pf-t--global--text--color--status--on-danger--default`. */\\n    --_inside-measure-color: var(--pf-v6-c-progress--m-danger--m-inside__measure--Color,\\n      var(--pf-t--global--text--color--status--on-danger--default, #fff));\\n  }\\n\\n  /* Description truncation */\\n\\n  \\u0026.truncated #description {\\n    overflow: hidden;\\n    text-overflow: ellipsis;\\n    white-space: nowrap;\\n  }\\n\\n}\\n\\n/* Danger animations */\\n\\n@media (prefers-reduced-motion: no-preference) {\\n  .danger #bar {\\n    animation: danger-jiggle\\n      var(--pf-t--global--motion--duration--fade--default, 0.3s)\\n      var(--pf-t--global--motion--timing-function--default, cubic-bezier(0.25, 0.1, 0.25, 1))\\n      both;\\n  }\\n\\n  .danger #status-icon {\\n    animation: fade-in\\n      var(--pf-t--global--motion--duration--fade--default, 0.3s)\\n      var(--pf-t--global--motion--timing-function--default, cubic-bezier(0.25, 0.1, 0.25, 1));\\n  }\\n}\\n\\n/* Elements */\\n\\n#description {\\n  grid-column: 1 / 2;\\n  overflow-wrap: anywhere;\\n  word-break: normal;\\n}\\n\\n#status {\\n  display: flex;\\n  grid-row: 1 / 2;\\n  grid-column: 2 / 3;\\n  /** Gap between the measure text and status icon. Defaults to `--pf-t--global--spacer--sm`. */\\n  gap: var(--pf-v6-c-progress__status--Gap,\\n    var(--pf-t--global--spacer--sm, 0.5rem));\\n  align-items: start;\\n  justify-content: end;\\n  text-align: end;\\n  overflow-wrap: anywhere;\\n  word-break: normal;\\n}\\n\\n/** Color of the status icon. Overridden by variant. Defaults to `--pf-t--global--icon--color--regular`. */\\n#status-icon {\\n  color: var(--_status-icon-color,\\n    var(--pf-v6-c-progress__status-icon--Color,\\n      var(--pf-t--global--icon--color--regular, #151515)));\\n}\\n\\n#bar {\\n  position: relative;\\n  grid-row: 2 / 3;\\n  grid-column: 1 / 3;\\n  align-self: center;\\n  /** Height of the progress bar track. Defaults to `--pf-t--global--spacer--md`. */\\n  height: var(--_bar-height,\\n    var(--pf-v6-c-progress__bar--Height,\\n      var(--pf-t--global--spacer--md, 1rem)));\\n  overflow: hidden;\\n  /** Background color of the bar track. Defaults to `--pf-t--global--color--nonstatus--gray--default`. */\\n  background-color: var(--pf-v6-c-progress__bar--BackgroundColor,\\n    var(--pf-t--global--color--nonstatus--gray--default, #f0f0f0));\\n  /** Border radius of the bar track. Defaults to `--pf-t--global--border--radius--medium`. */\\n  border-radius: var(--pf-v6-c-progress__bar--BorderRadius,\\n    var(--pf-t--global--border--radius--medium, 6px));\\n\\n  \\u0026::before {\\n    position: absolute;\\n    inset: 0;\\n    pointer-events: none;\\n    content: '';\\n    /** Border width of the bar track for high-contrast mode. Defaults to `--pf-t--global--border--width--high-contrast--regular`. */\\n    border: var(--pf-v6-c-progress__bar--BorderWidth,\\n        var(--pf-t--global--border--width--high-contrast--regular, 0))\\n      solid\\n      /** Border color of the bar track for high-contrast mode. Defaults to `--pf-t--global--border--color--high-contrast`. */\\n      var(--pf-v6-c-progress__bar--BorderColor,\\n        var(--pf-t--global--border--color--high-contrast, transparent));\\n    border-radius: inherit;\\n  }\\n}\\n\\n#indicator {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-inline-start: 0;\\n  /** Height of the filled indicator. Defaults to `--pf-v6-c-progress__bar--Height`. */\\n  height: var(--_bar-height,\\n    var(--pf-v6-c-progress__indicator--Height,\\n      var(--pf-v6-c-progress__bar--Height,\\n        var(--pf-t--global--spacer--md, 1rem))));\\n  /** Background color of the filled indicator. Overridden by variant. Defaults to `--pf-t--global--color--brand--default`. */\\n  background-color: var(--_indicator-bg,\\n    var(--pf-v6-c-progress__indicator--BackgroundColor,\\n      var(--pf-t--global--color--brand--default, #0066cc)));\\n\\n  \\u0026::before {\\n    position: absolute;\\n    inset: 0;\\n    content: '';\\n    /** Border width of the indicator for high-contrast mode. Defaults to `--pf-t--global--border--width--high-contrast--extra-strong`. */\\n    border: var(--pf-v6-c-progress__indicator--BorderWidth,\\n        var(--pf-t--global--border--width--high-contrast--extra-strong, 0))\\n      solid\\n      /** Border color of the indicator for high-contrast mode. */\\n      var(--pf-v6-c-progress__indicator--BorderColor, transparent);\\n    border-radius: var(--pf-v6-c-progress__bar--BorderRadius,\\n      var(--pf-t--global--border--radius--medium, 6px));\\n  }\\n}\\n\\n#helper-text {\\n  grid-row: 3 / 4;\\n  grid-column: 1 / 3;\\n  /** Negative offset to bring helper text closer to the bar. */\\n  margin-block-start: var(--pf-v6-c-progress__helper-text--MarginBlockStart,\\n    calc(var(--pf-t--global--spacer--sm, 0.5rem) - var(--pf-v6-c-progress--GridGap,\\n      var(--pf-t--global--spacer--md, 1rem))));\\n}\\n\\n/** Minimum width of the measure display for visual alignment across stacked progress bars. Set via CSS custom property. */\\n#measure {\\n  min-width: var(--pf-v6-c-progress__measure--m-static-width--MinWidth, 0);\\n  font-variant-numeric: tabular-nums;\\n}\\n\\nsvg {\\n  width: 1em;\\n  height: 1em;\\n  fill: currentcolor;\\n}\\n\\n@keyframes danger-jiggle {\\n  33% {\\n    translate: -2px;\\n  }\\n\\n  66% {\\n    translate: 3px;\\n  }\\n\\n  100% {\\n    translate: 0;\\n  }\\n}\\n\\n@keyframes fade-in {\\n  from {\\n    opacity: 0;\\n  }\\n\\n  to {\\n    opacity: 1;\\n  }\\n}\\n\""));
var cem_pf_v6_progress_default = s42;

// elements/cem-pf-v6-progress/cem-pf-v6-progress.ts
var checkCircleIcon = T`<svg id="status-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true"><path d="M256 512A256 256 0 1 0 256 0a256 256 0 1 0 0 512zM369 209L241 337c-9.4 9.4-24.6 9.4-33.9 0l-64-64c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l47 47L335 175c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9z"/></svg>`;
var triangleExclamationIcon = T`<svg id="status-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true"><path d="M256 32c14.2 0 27.3 7.5 34.5 19.8l216 368c7.3 12.4 7.3 27.7 .2 40.1S486.3 480 472 480H40c-14.3 0-27.6-7.7-34.7-20.1s-7-27.8 .2-40.1l216-368C228.7 39.5 241.8 32 256 32zm0 128c-13.3 0-24 10.7-24 24V296c0 13.3 10.7 24 24 24s24-10.7 24-24V184c0-13.3-10.7-24-24-24zm32 224a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"/></svg>`;
var circleExclamationIcon = T`<svg id="status-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true"><path d="M256 512A256 256 0 1 0 256 0a256 256 0 1 0 0 512zm0-384c13.3 0 24 10.7 24 24V264c0 13.3-10.7 24-24 24s-24-10.7-24-24V152c0-13.3 10.7-24 24-24zM224 352a32 32 0 1 1 64 0 32 32 0 1 1 -64 0z"/></svg>`;
var VARIANT_ICONS = /* @__PURE__ */ new Map([
  ["success", checkCircleIcon],
  ["warning", triangleExclamationIcon],
  ["danger", circleExclamationIcon]
]);
var _valueText_dec, _hideStatusIcon_dec, _variant_dec11, _measureLocation_dec, _size_dec2, _min_dec, _max_dec, _accessibleLabel_dec, _truncated_dec, _description_dec3, _value_dec3, _a39, _CemPfV6Progress_decorators, _internals8, _init39, _value3, _description3, _truncated, _accessibleLabel, _max, _min, _size2, _measureLocation, _variant11, _hideStatusIcon, _valueText, _hasHelperText, _CemPfV6Progress_instances, calculatedPercentage_get, displayText_get, icon_get, onHelperTextSlotchange_fn;
_CemPfV6Progress_decorators = [t3("cem-pf-v6-progress")];
var CemPfV6Progress = class extends (_a39 = i3, _value_dec3 = [n4({ type: Number })], _description_dec3 = [n4()], _truncated_dec = [n4({ type: Boolean })], _accessibleLabel_dec = [n4({ attribute: "accessible-label" })], _max_dec = [n4({ type: Number })], _min_dec = [n4({ type: Number })], _size_dec2 = [n4()], _measureLocation_dec = [n4({ attribute: "measure-location" })], _variant_dec11 = [n4()], _hideStatusIcon_dec = [n4({ type: Boolean, attribute: "hide-status-icon" })], _valueText_dec = [n4({ attribute: "value-text" })], _a39) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemPfV6Progress_instances);
    __privateAdd(this, _internals8, this.attachInternals());
    __privateAdd(this, _value3, __runInitializers(_init39, 8, this, 0)), __runInitializers(_init39, 11, this);
    __privateAdd(this, _description3, __runInitializers(_init39, 12, this)), __runInitializers(_init39, 15, this);
    __privateAdd(this, _truncated, __runInitializers(_init39, 16, this, false)), __runInitializers(_init39, 19, this);
    __privateAdd(this, _accessibleLabel, __runInitializers(_init39, 20, this)), __runInitializers(_init39, 23, this);
    __privateAdd(this, _max, __runInitializers(_init39, 24, this, 100)), __runInitializers(_init39, 27, this);
    __privateAdd(this, _min, __runInitializers(_init39, 28, this, 0)), __runInitializers(_init39, 31, this);
    __privateAdd(this, _size2, __runInitializers(_init39, 32, this)), __runInitializers(_init39, 35, this);
    __privateAdd(this, _measureLocation, __runInitializers(_init39, 36, this)), __runInitializers(_init39, 39, this);
    __privateAdd(this, _variant11, __runInitializers(_init39, 40, this)), __runInitializers(_init39, 43, this);
    __privateAdd(this, _hideStatusIcon, __runInitializers(_init39, 44, this, false)), __runInitializers(_init39, 47, this);
    __privateAdd(this, _valueText, __runInitializers(_init39, 48, this)), __runInitializers(_init39, 51, this);
    __privateAdd(this, _hasHelperText, false);
  }
  updated(changed) {
    if (changed.has("value") || changed.has("min") || changed.has("max")) {
      __privateGet(this, _internals8).ariaValueNow = __privateGet(this, _CemPfV6Progress_instances, calculatedPercentage_get).toString();
      __privateGet(this, _internals8).ariaValueMin = "0";
      __privateGet(this, _internals8).ariaValueMax = "100";
    }
    if (changed.has("valueText")) {
      __privateGet(this, _internals8).ariaValueText = this.valueText ?? null;
    }
    if (changed.has("accessibleLabel") || changed.has("description")) {
      __privateGet(this, _internals8).ariaLabel = this.accessibleLabel ?? this.description ?? "Progress status";
    }
  }
  connectedCallback() {
    super.connectedCallback();
    __privateGet(this, _internals8).role = "progressbar";
    __privateGet(this, _internals8).ariaValueNow = __privateGet(this, _CemPfV6Progress_instances, calculatedPercentage_get).toString();
    __privateGet(this, _internals8).ariaValueMin = "0";
    __privateGet(this, _internals8).ariaValueMax = "100";
    __privateGet(this, _internals8).ariaLabel = this.accessibleLabel ?? this.description ?? "Progress status";
  }
  render() {
    const pct = __privateGet(this, _CemPfV6Progress_instances, calculatedPercentage_get);
    const displayText = __privateGet(this, _CemPfV6Progress_instances, displayText_get);
    const icon = __privateGet(this, _CemPfV6Progress_instances, icon_get);
    const noMeasure = this.measureLocation === "none";
    const inside = this.measureLocation === "inside";
    const hasDescription = this.description != null;
    const hasIcon = this.variant != null && !this.hideStatusIcon;
    const singleline = !hasDescription;
    const classes = {
      [this.size ?? ""]: !!this.size,
      [this.measureLocation ?? ""]: !!this.measureLocation && this.measureLocation !== "none",
      [this.variant ?? ""]: !!this.variant,
      singleline,
      truncated: this.truncated
    };
    return T`
      <div id="container" class="${e5(classes)}">
        <div id="description"
             ?hidden="${!hasDescription}"
>${this.description ?? ""}</div>

        <div id="status"
             aria-hidden="true"
             ?hidden="${noMeasure && !hasIcon}">
          ${!inside && !noMeasure ? T`<span id="measure">${displayText}</span>` : A}
          ${icon}
        </div>

        <div id="bar">
          <div id="indicator"
               style="${o8({ width: `${pct}%` })}">
            ${inside && !noMeasure ? T`<span id="measure">${displayText}</span>` : A}
          </div>
        </div>

        <div id="helper-text" ?hidden="${!__privateGet(this, _hasHelperText)}">
          <slot id="helper-text-slot"
                name="helper-text"
                @slotchange="${__privateMethod(this, _CemPfV6Progress_instances, onHelperTextSlotchange_fn)}"></slot>
        </div>
      </div>
    `;
  }
};
_init39 = __decoratorStart(_a39);
_internals8 = new WeakMap();
_value3 = new WeakMap();
_description3 = new WeakMap();
_truncated = new WeakMap();
_accessibleLabel = new WeakMap();
_max = new WeakMap();
_min = new WeakMap();
_size2 = new WeakMap();
_measureLocation = new WeakMap();
_variant11 = new WeakMap();
_hideStatusIcon = new WeakMap();
_valueText = new WeakMap();
_hasHelperText = new WeakMap();
_CemPfV6Progress_instances = new WeakSet();
calculatedPercentage_get = function() {
  const { value, min, max } = this;
  const percentage = Math.round((value - min) / (max - min) * 100);
  if (Number.isNaN(percentage) || percentage < 0) {
    return 0;
  }
  return Math.min(percentage, 100);
};
displayText_get = function() {
  return this.valueText ?? `${__privateGet(this, _CemPfV6Progress_instances, calculatedPercentage_get)}%`;
};
icon_get = function() {
  if (this.hideStatusIcon) {
    return A;
  }
  return VARIANT_ICONS.get(this.variant) ?? A;
};
onHelperTextSlotchange_fn = function(event) {
  const slot = event.target;
  const elements = slot?.assignedElements() ?? [];
  __privateSet(this, _hasHelperText, elements.length > 0);
  __privateGet(this, _internals8).ariaDescribedByElements = elements.length ? elements : null;
  this.requestUpdate();
};
__decorateElement(_init39, 4, "value", _value_dec3, CemPfV6Progress, _value3);
__decorateElement(_init39, 4, "description", _description_dec3, CemPfV6Progress, _description3);
__decorateElement(_init39, 4, "truncated", _truncated_dec, CemPfV6Progress, _truncated);
__decorateElement(_init39, 4, "accessibleLabel", _accessibleLabel_dec, CemPfV6Progress, _accessibleLabel);
__decorateElement(_init39, 4, "max", _max_dec, CemPfV6Progress, _max);
__decorateElement(_init39, 4, "min", _min_dec, CemPfV6Progress, _min);
__decorateElement(_init39, 4, "size", _size_dec2, CemPfV6Progress, _size2);
__decorateElement(_init39, 4, "measureLocation", _measureLocation_dec, CemPfV6Progress, _measureLocation);
__decorateElement(_init39, 4, "variant", _variant_dec11, CemPfV6Progress, _variant11);
__decorateElement(_init39, 4, "hideStatusIcon", _hideStatusIcon_dec, CemPfV6Progress, _hideStatusIcon);
__decorateElement(_init39, 4, "valueText", _valueText_dec, CemPfV6Progress, _valueText);
CemPfV6Progress = __decorateElement(_init39, 0, "CemPfV6Progress", _CemPfV6Progress_decorators, CemPfV6Progress);
__publicField(CemPfV6Progress, "styles", cem_pf_v6_progress_default);
__runInitializers(_init39, 1, CemPfV6Progress);

// lit-css:elements/cem-pf-v6-select/cem-pf-v6-select.css
var s43 = new CSSStyleSheet();
s43.replaceSync(JSON.parse(`":host {\\n\\n  /* Form control custom properties */\\n  --cem-pf-v6-c-form-control--ColumnGap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-form-control--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-form-control--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-form-control--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-form-control--Resize: none;\\n  --cem-pf-v6-c-form-control--OutlineOffset: -6px;\\n  --cem-pf-v6-c-form-control--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-form-control--before--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-form-control--before--BorderStyle: solid;\\n  --cem-pf-v6-c-form-control--before--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-form-control--before--BorderRadius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  --cem-pf-v6-c-form-control--after--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-form-control--after--BorderStyle: solid;\\n  --cem-pf-v6-c-form-control--after--BorderColor: transparent;\\n  --cem-pf-v6-c-form-control--after--BorderRadius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  --cem-pf-v6-c-form-control--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-form-control--Width: 100%;\\n  --cem-pf-v6-c-form-control--inset--base: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-form-control--PaddingBlockStart--base: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-form-control--PaddingBlockEnd--base: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-form-control--PaddingInlineEnd--base: var(--cem-pf-v6-c-form-control--inset--base);\\n  --cem-pf-v6-c-form-control--PaddingInlineStart--base: var(--cem-pf-v6-c-form-control--inset--base);\\n  --cem-pf-v6-c-form-control--PaddingBlockStart: var(--cem-pf-v6-c-form-control__select--PaddingBlockStart);\\n  --cem-pf-v6-c-form-control--PaddingBlockEnd: var(--cem-pf-v6-c-form-control__select--PaddingBlockEnd);\\n  --cem-pf-v6-c-form-control--PaddingInlineStart: var(--cem-pf-v6-c-form-control__select--PaddingInlineStart);\\n  --cem-pf-v6-c-form-control--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__select--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__select--PaddingBlockStart: var(--cem-pf-v6-c-form-control--PaddingBlockStart--base);\\n  --cem-pf-v6-c-form-control__select--PaddingBlockEnd: var(--cem-pf-v6-c-form-control--PaddingBlockEnd--base);\\n  --cem-pf-v6-c-form-control__select--PaddingInlineEnd: var(--cem-pf-v6-c-form-control--PaddingInlineEnd--base);\\n  --cem-pf-v6-c-form-control__select--PaddingInlineStart: var(--cem-pf-v6-c-form-control--PaddingInlineStart--base);\\n  --cem-pf-v6-c-form-control__utilities--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__utilities--select--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__utilities--select--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__select--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control--hover--after--BorderWidth: var(--pf-t--global--border--width--control--hover);\\n  --cem-pf-v6-c-form-control--hover--after--BorderColor: var(--pf-t--global--border--color--hover);\\n\\n  position: relative;\\n  display: grid;\\n  grid-template-columns: 1fr auto;\\n  column-gap: var(--cem-pf-v6-c-form-control--ColumnGap);\\n  align-items: start;\\n  width: var(--cem-pf-v6-c-form-control--Width);\\n  font-size: var(--cem-pf-v6-c-form-control--FontSize);\\n  line-height: var(--cem-pf-v6-c-form-control--LineHeight);\\n  resize: var(--cem-pf-v6-c-form-control--Resize);\\n  background-color: var(--cem-pf-v6-c-form-control--BackgroundColor);\\n  border-radius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n}\\n\\n:host::before,\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n}\\n\\n:host::before {\\n  border-color: var(--cem-pf-v6-c-form-control--before--BorderColor);\\n  border-style: var(--cem-pf-v6-c-form-control--before--BorderStyle);\\n  border-width: var(--cem-pf-v6-c-form-control--before--BorderWidth);\\n  border-radius: var(--cem-pf-v6-c-form-control--before--BorderRadius);\\n}\\n\\n:host::after {\\n  border: var(--cem-pf-v6-c-form-control--after--BorderWidth) var(--cem-pf-v6-c-form-control--after--BorderStyle) var(--cem-pf-v6-c-form-control--after--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-form-control--before--BorderRadius);\\n}\\n\\n:host(:hover) {\\n  --cem-pf-v6-c-form-control--after--BorderColor: var(--cem-pf-v6-c-form-control--hover--after--BorderColor);\\n  --cem-pf-v6-c-form-control--after--BorderWidth: var(--cem-pf-v6-c-form-control--hover--after--BorderWidth);\\n}\\n\\n#select {\\n  grid-row: 1 / 2;\\n  grid-column: 1 / -1;\\n  padding-block-start: var(--cem-pf-v6-c-form-control--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-form-control--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-form-control--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-form-control--PaddingInlineEnd);\\n  margin: 0;\\n  font-family: inherit;\\n  font-size: 100%;\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--cem-pf-v6-c-form-control--Color);\\n  appearance: none;\\n  background-color: var(--cem-pf-v6-c-form-control--BackgroundColor);\\n  border: none;\\n  border-radius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  outline-offset: var(--cem-pf-v6-c-form-control--OutlineOffset);\\n  cursor: pointer;\\n  background-image: url(\\"data:image/svg+xml,%3Csvg width='10' height='6' viewBox='0 0 10 6' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M1 1L5 5L9 1' stroke='%236a6e73' stroke-width='1.5' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E\\");\\n  background-repeat: no-repeat;\\n  background-position: right var(--pf-t--global--spacer--md) center;\\n}\\n\\n#select * {\\n  color: var(--cem-pf-v6-c-form-control--Color);\\n}\\n"`));
var cem_pf_v6_select_default = s43;

// elements/cem-pf-v6-select/cem-pf-v6-select.ts
var _invalid_dec, _disabled_dec7, _value_dec4, _a40, _PfV6Select_decorators, _internals9, _observer, _init40, _value4, _disabled7, _invalid, _PfV6Select_instances, selectEl_get, cloneOptions_fn, _onChange, _onInput;
_PfV6Select_decorators = [t3("cem-pf-v6-select")];
var PfV6Select = class extends (_a40 = i3, _value_dec4 = [n4({ reflect: true })], _disabled_dec7 = [n4({ type: Boolean, reflect: true })], _invalid_dec = [n4({ type: Boolean, reflect: true })], _a40) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Select_instances);
    __privateAdd(this, _internals9, this.attachInternals());
    __privateAdd(this, _observer, new MutationObserver(() => __privateMethod(this, _PfV6Select_instances, cloneOptions_fn).call(this)));
    __privateAdd(this, _value4, __runInitializers(_init40, 8, this, "")), __runInitializers(_init40, 11, this);
    __privateAdd(this, _disabled7, __runInitializers(_init40, 12, this, false)), __runInitializers(_init40, 15, this);
    __privateAdd(this, _invalid, __runInitializers(_init40, 16, this, false)), __runInitializers(_init40, 19, this);
    __privateAdd(this, _onChange, () => {
      const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
      if (!select) return;
      this.value = select.value;
      __privateGet(this, _internals9).setFormValue(select.value);
      this.dispatchEvent(new Event("change", { bubbles: true }));
    });
    __privateAdd(this, _onInput, () => {
      const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
      if (!select) return;
      this.value = select.value;
      __privateGet(this, _internals9).setFormValue(select.value);
      this.dispatchEvent(new Event("input", { bubbles: true }));
    });
  }
  render() {
    return T`
      <select id="select"
              ?disabled=${this.disabled}
              aria-invalid=${this.invalid ? "true" : "false"}
              @change=${__privateGet(this, _onChange)}
              @input=${__privateGet(this, _onInput)}></select>
    `;
  }
  firstUpdated() {
    __privateMethod(this, _PfV6Select_instances, cloneOptions_fn).call(this);
    const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
    if (select && this.value) {
      select.value = this.value;
    }
    __privateGet(this, _internals9).setFormValue(select?.value || null);
    __privateGet(this, _observer).observe(this, {
      childList: true,
      subtree: true,
      attributes: true
    });
  }
  updated(changed) {
    if (changed.has("value")) {
      const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
      if (select && select.value !== this.value) {
        select.value = this.value;
      }
      __privateGet(this, _internals9).setFormValue(this.value || null);
    }
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(this, _observer).disconnect();
  }
};
_init40 = __decoratorStart(_a40);
_internals9 = new WeakMap();
_observer = new WeakMap();
_value4 = new WeakMap();
_disabled7 = new WeakMap();
_invalid = new WeakMap();
_PfV6Select_instances = new WeakSet();
selectEl_get = function() {
  return this.shadowRoot?.getElementById("select");
};
cloneOptions_fn = function() {
  const select = __privateGet(this, _PfV6Select_instances, selectEl_get);
  if (!select) return;
  const currentValue = select.value;
  select.innerHTML = "";
  const options = this.querySelectorAll("option");
  options.forEach((option) => {
    select.appendChild(option.cloneNode(true));
  });
  select.value = currentValue;
};
_onChange = new WeakMap();
_onInput = new WeakMap();
__decorateElement(_init40, 4, "value", _value_dec4, PfV6Select, _value4);
__decorateElement(_init40, 4, "disabled", _disabled_dec7, PfV6Select, _disabled7);
__decorateElement(_init40, 4, "invalid", _invalid_dec, PfV6Select, _invalid);
PfV6Select = __decorateElement(_init40, 0, "PfV6Select", _PfV6Select_decorators, PfV6Select);
__publicField(PfV6Select, "formAssociated", true);
__publicField(PfV6Select, "shadowRootOptions", {
  ...i3.shadowRootOptions,
  delegatesFocus: true
});
__publicField(PfV6Select, "styles", cem_pf_v6_select_default);
__runInitializers(_init40, 1, PfV6Select);

// lit-css:elements/cem-pf-v6-skip-to-content/cem-pf-v6-skip-to-content.css
var s44 = new CSSStyleSheet();
s44.replaceSync(JSON.parse('":host {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-inline-start: 50%;\\n  z-index: var(--pf-t--global--z-index--xl);\\n  transform: translateX(-50%);\\n}\\n\\n#wrapper {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-inline-start: 50%;\\n  z-index: var(--pf-t--global--z-index--xl);\\n  transform: translate(-50%, -100%);\\n  transition: transform var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n\\n  \\u0026:focus-within {\\n    transform: translate(-50%, 0);\\n  }\\n}\\n\\n#link {\\n  display: inline-flex;\\n  align-items: center;\\n  justify-content: center;\\n  padding-block-start: var(--pf-t--global--spacer--control--vertical--default);\\n  padding-block-end: var(--pf-t--global--spacer--control--vertical--default);\\n  padding-inline-start: var(--pf-t--global--spacer--control--horizontal--default);\\n  padding-inline-end: var(--pf-t--global--spacer--control--horizontal--default);\\n  font-family: var(--pf-t--global--font--family--body);\\n  font-size: var(--pf-t--global--font--size--body--default);\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--inverse);\\n  text-align: center;\\n  text-decoration: none;\\n  white-space: nowrap;\\n  background-color: var(--pf-t--global--color--brand--default);\\n  border: 0;\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  cursor: pointer;\\n\\n  \\u0026:hover {\\n    background-color: var(--pf-t--global--color--brand--hover);\\n  }\\n\\n  \\u0026:active {\\n    background-color: var(--pf-t--global--color--brand--clicked);\\n  }\\n\\n  \\u0026:focus {\\n    outline: var(--pf-t--global--border--width--action--default) solid var(--pf-t--global--focus-ring--color--default);\\n    outline-offset: var(--pf-t--global--focus-ring--position--offset);\\n  }\\n}\\n"'));
var cem_pf_v6_skip_to_content_default = s44;

// elements/cem-pf-v6-skip-to-content/cem-pf-v6-skip-to-content.ts
var _href_dec4, _a41, _PfV6SkipToContent_decorators, _init41, _href4;
_PfV6SkipToContent_decorators = [t3("cem-pf-v6-skip-to-content")];
var PfV6SkipToContent = class extends (_a41 = i3, _href_dec4 = [n4({ reflect: true })], _a41) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _href4, __runInitializers(_init41, 8, this, "#main-content")), __runInitializers(_init41, 11, this);
  }
  render() {
    return T`
      <div id="wrapper">
        <a id="link"
           part="link"
           href="${this.href}">
          <span id="text">
            <slot>Skip to content</slot>
          </span>
        </a>
      </div>
    `;
  }
};
_init41 = __decoratorStart(_a41);
_href4 = new WeakMap();
__decorateElement(_init41, 4, "href", _href_dec4, PfV6SkipToContent, _href4);
PfV6SkipToContent = __decorateElement(_init41, 0, "PfV6SkipToContent", _PfV6SkipToContent_decorators, PfV6SkipToContent);
__publicField(PfV6SkipToContent, "styles", cem_pf_v6_skip_to_content_default);
__runInitializers(_init41, 1, PfV6SkipToContent);

// lit-css:elements/cem-pf-v6-switch/cem-pf-v6-switch.css
var s45 = new CSSStyleSheet();
s45.replaceSync(JSON.parse('":host {\\n  display: inline-flex;\\n}\\n\\n#switch-label {\\n  display: inline-flex;\\n  align-items: center;\\n  gap: 0.5rem;\\n  cursor: pointer;\\n  user-select: none;\\n}\\n\\n#switch-input {\\n  position: absolute;\\n  opacity: 0;\\n  pointer-events: none;\\n}\\n\\n#switch-toggle {\\n  position: relative;\\n  display: inline-block;\\n  width: calc(calc(var(--pf-t--global--font--size--sm) * var(--pf-t--global--font--line-height--body)) + 0.125rem + calc(var(--pf-t--global--font--size--sm) - 0.125rem));\\n  height: calc(var(--pf-t--global--font--size--sm) * var(--pf-t--global--font--line-height--body));\\n  background-color: var(--pf-t--global--background--color--control--default);\\n  border-radius: var(--pf-t--global--border--radius--pill);\\n\\n  \\u0026::before {\\n    position: absolute;\\n    inset-block-start: 50%;\\n    inset-inline-start: calc((calc(var(--pf-t--global--font--size--sm) * var(--pf-t--global--font--line-height--body)) - calc(var(--pf-t--global--font--size--sm) - 0.125rem)) / 2);\\n    display: block;\\n    width: calc(var(--pf-t--global--font--size--sm) - 0.125rem);\\n    height: calc(var(--pf-t--global--font--size--sm) - 0.125rem);\\n    content: \\"\\";\\n    background-color: var(--pf-t--global--icon--color--subtle);\\n    border: var(--pf-t--global--border--width--regular) solid transparent;\\n    border-radius: var(--pf-t--global--border--radius--pill);\\n    transition: transform var(--pf-t--global--motion--duration--md) var(--pf-t--global--motion--timing-function--default),\\n                background-color var(--pf-t--global--motion--duration--md) var(--pf-t--global--motion--timing-function--default);\\n    transform: translateY(-50%);\\n  }\\n\\n  \\u0026::after {\\n    position: absolute;\\n    inset: 0;\\n    content: \\"\\";\\n    border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n    border-radius: var(--pf-t--global--border--radius--pill);\\n  }\\n}\\n\\n.switch-toggle-icon {\\n  position: absolute;\\n  inset-block-start: 0;\\n  inset-block-end: 0;\\n  inset-inline-start: calc(var(--pf-t--global--font--size--sm) * 0.4);\\n  display: none;\\n  align-items: center;\\n  font-size: calc(var(--pf-t--global--font--size--sm) * 0.625);\\n  color: var(--pf-t--global--icon--color--on-brand--default);\\n\\n  \\u0026 svg {\\n    width: 1em;\\n    height: 1em;\\n  }\\n}\\n\\n#switch-input:checked + #switch-toggle {\\n  background-color: var(--pf-t--global--color--brand--default);\\n\\n  \\u0026::before {\\n    transform: translate(calc(100% + 0.125rem), -50%);\\n    background-color: var(--pf-t--global--icon--color--inverse);\\n  }\\n\\n  \\u0026::after {\\n    border-color: transparent;\\n  }\\n\\n  \\u0026 .switch-toggle-icon {\\n    display: flex;\\n  }\\n}\\n\\n#switch-input:not(:checked) + #switch-toggle .switch-toggle-icon {\\n  display: none;\\n}\\n\\n#switch-input:focus-visible + #switch-toggle {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: 2px;\\n}\\n\\n#switch-input:disabled + #switch-toggle {\\n  background-color: var(--pf-t--global--background--color--disabled--default);\\n  cursor: not-allowed;\\n\\n  \\u0026::before {\\n    background-color: var(--pf-t--global--text--color--disabled);\\n  }\\n\\n  \\u0026::after {\\n    border-color: var(--pf-t--global--border--color--disabled);\\n  }\\n\\n  \\u0026 .switch-toggle-icon {\\n    color: var(--pf-t--global--icon--color--disabled);\\n  }\\n}\\n\\n#switch-input:disabled ~ #switch-text {\\n  color: var(--pf-t--global--text--color--disabled);\\n  cursor: not-allowed;\\n}\\n\\n:host([disabled]) #switch-label {\\n  cursor: not-allowed;\\n}\\n\\n#switch-text {\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: 0.875rem;\\n  line-height: 1.5;\\n}\\n\\n#switch-input:checked ~ #switch-text {\\n  color: var(--pf-t--global--text--color--regular);\\n}\\n\\n:host([label-position=\\"start\\"]) #switch-label {\\n  flex-direction: row-reverse;\\n  justify-content: flex-end;\\n}\\n"'));
var cem_pf_v6_switch_default = s45;

// elements/cem-pf-v6-switch/cem-pf-v6-switch.ts
var _labelPosition_dec, _disabled_dec8, _checked_dec2, _a42, _PfV6Switch_decorators, _internals10, _init42, _checked2, _disabled8, _labelPosition, _PfV6Switch_instances, onInput_fn2;
_PfV6Switch_decorators = [t3("cem-pf-v6-switch")];
var PfV6Switch = class extends (_a42 = i3, _checked_dec2 = [n4({ type: Boolean, reflect: true })], _disabled_dec8 = [n4({ type: Boolean, reflect: true })], _labelPosition_dec = [n4({ reflect: true, attribute: "label-position" })], _a42) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Switch_instances);
    __privateAdd(this, _internals10, this.attachInternals());
    __privateAdd(this, _checked2, __runInitializers(_init42, 8, this, false)), __runInitializers(_init42, 11, this);
    __privateAdd(this, _disabled8, __runInitializers(_init42, 12, this, false)), __runInitializers(_init42, 15, this);
    __privateAdd(this, _labelPosition, __runInitializers(_init42, 16, this)), __runInitializers(_init42, 19, this);
  }
  render() {
    return T`
      <label id="switch-label">
        <input type="checkbox"
               id="switch-input"
               role="switch"
               .checked=${l3(this.checked)}
               ?disabled=${this.disabled}
               @input=${__privateMethod(this, _PfV6Switch_instances, onInput_fn2)}>
        <span id="switch-toggle"
              aria-hidden="true">
          <span class="switch-toggle-icon">
            <svg width="10"
                 height="10"
                 viewBox="0 0 512 512"
                 fill="currentColor">
              <path d="M173.898 439.404l-166.4-166.4c-9.997-9.997-9.997-26.206 0-36.204l36.203-36.204c9.997-9.998 26.207-9.998 36.204 0L192 312.69 432.095 72.596c9.997-9.997 26.207-9.997 36.204 0l36.203 36.204c9.997 9.997 9.997 26.206 0 36.204l-294.4 294.401c-9.998 9.997-26.207 9.997-36.204-.001z"/>
            </svg>
          </span>
        </span>
        <span id="switch-text"><slot></slot></span>
      </label>
    `;
  }
};
_init42 = __decoratorStart(_a42);
_internals10 = new WeakMap();
_checked2 = new WeakMap();
_disabled8 = new WeakMap();
_labelPosition = new WeakMap();
_PfV6Switch_instances = new WeakSet();
onInput_fn2 = function() {
  const input = this.shadowRoot?.getElementById("switch-input");
  if (input) {
    this.checked = input.checked;
    __privateGet(this, _internals10).setFormValue(input.checked ? "on" : null);
  }
};
__decorateElement(_init42, 4, "checked", _checked_dec2, PfV6Switch, _checked2);
__decorateElement(_init42, 4, "disabled", _disabled_dec8, PfV6Switch, _disabled8);
__decorateElement(_init42, 4, "labelPosition", _labelPosition_dec, PfV6Switch, _labelPosition);
PfV6Switch = __decorateElement(_init42, 0, "PfV6Switch", _PfV6Switch_decorators, PfV6Switch);
__publicField(PfV6Switch, "formAssociated", true);
__publicField(PfV6Switch, "styles", cem_pf_v6_switch_default);
__runInitializers(_init42, 1, PfV6Switch);

// lit-css:elements/cem-pf-v6-tab/cem-pf-v6-tab.css
var s46 = new CSSStyleSheet();
s46.replaceSync(JSON.parse('":host {\\n  display: block;\\n  padding: var(--pf-t--global--spacer--md);\\n  height: 100%;\\n  box-sizing: border-box;\\n}\\n"'));
var cem_pf_v6_tab_default = s46;

// elements/cem-pf-v6-tab/cem-pf-v6-tab.ts
var _title_dec, _a43, _PfV6Tab_decorators, _init43, _title;
_PfV6Tab_decorators = [t3("cem-pf-v6-tab")];
var PfV6Tab = class extends (_a43 = i3, _title_dec = [n4({ reflect: true })], _a43) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _title, __runInitializers(_init43, 8, this, "")), __runInitializers(_init43, 11, this);
  }
  connectedCallback() {
    super.connectedCallback();
    this.dispatchEvent(new Event("cem-pf-v6-tab-connected", { bubbles: true }));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.dispatchEvent(new Event("cem-pf-v6-tab-disconnected", { bubbles: true }));
  }
  updated(changed) {
    if (changed.has("title")) {
      this.dispatchEvent(new Event("cem-pf-v6-tab-title-changed", { bubbles: true }));
    }
  }
  render() {
    return T`<slot></slot>`;
  }
};
_init43 = __decoratorStart(_a43);
_title = new WeakMap();
__decorateElement(_init43, 4, "title", _title_dec, PfV6Tab, _title);
PfV6Tab = __decorateElement(_init43, 0, "PfV6Tab", _PfV6Tab_decorators, PfV6Tab);
__publicField(PfV6Tab, "styles", cem_pf_v6_tab_default);
__runInitializers(_init43, 1, PfV6Tab);

// lit-css:elements/cem-pf-v6-tabs/cem-pf-v6-tabs.css
var s47 = new CSSStyleSheet();
s47.replaceSync(JSON.parse(`":host {\\n  display: block;\\n  height: 100%;\\n  position: relative;\\n  overflow: hidden;\\n}\\n\\n#tabs-container {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n@property --_cem-pf-v6-tabs--link-accent--length {\\n  syntax: \\"\\u003clength\\u003e\\";\\n  inherits: true;\\n  initial-value: 0px;\\n}\\n\\n#tabs {\\n  --_tab--Color: var(--pf-t--global--text--color--subtle);\\n  --_tab--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --_tab--BorderWidth: var(--pf-t--global--border--width--action--plain--hover);\\n  --_tab--BorderColor: transparent;\\n  --_tab--hover--Color: var(--pf-t--global--text--color--regular);\\n  --_tab--hover--BackgroundColor: var(--pf-t--global--background--color--action--plain--hover);\\n  --_tab--current--Color: var(--pf-t--global--text--color--regular);\\n  --_tab--current--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --_tab--disabled--Color: var(--pf-t--global--text--color--on-disabled);\\n  --_tab--disabled--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n\\n  --_cem-pf-v6-tabs--link-accent--start: 0px;\\n  --_cem-pf-v6-tabs--link-accent--length: 0px;\\n  --_cem-pf-v6-tabs--link-accent--color: var(--pf-t--global--border--color--clicked);\\n  --_cem-pf-v6-tabs--link-accent--border-size: var(--pf-t--global--border--width--extra-strong);\\n\\n  position: sticky;\\n  inset-block-start: 0;\\n  z-index: var(--pf-t--global--z-index--md);\\n  display: flex;\\n  overflow-x: auto;\\n  overflow-y: hidden;\\n  -webkit-overflow-scrolling: touch;\\n  scroll-behavior: smooth;\\n  flex-shrink: 0;\\n  background-color: var(--pf-t--global--background--color--primary--default);\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    inset-block-end: 0;\\n    inset-inline-start: 0;\\n    inset-inline-end: 0;\\n    border-block-end: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n  }\\n\\n  \\u0026::after {\\n    content: '';\\n    position: absolute;\\n    inset-block-start: auto;\\n    inset-block-end: 0;\\n    inset-inline-start: 0;\\n    width: var(--_cem-pf-v6-tabs--link-accent--length);\\n    height: 0;\\n    border: 0 solid var(--_cem-pf-v6-tabs--link-accent--color);\\n    border-block-end-width: var(--_cem-pf-v6-tabs--link-accent--border-size);\\n    transform-origin: 0 center;\\n    translate: var(--_cem-pf-v6-tabs--link-accent--start) 0;\\n    transition-property: width, translate;\\n    transition-duration: var(--pf-t--global--motion--duration--md);\\n    transition-timing-function: var(--pf-t--global--motion--timing-function--decelerate);\\n  }\\n}\\n\\n#panels {\\n  flex: 1;\\n  min-height: 0;\\n  overflow: hidden;\\n}\\n\\n.panel {\\n  display: block;\\n  height: 100%;\\n  overflow-y: auto;\\n  box-sizing: border-box;\\n  contain: layout style paint;\\n\\n  \\u0026[hidden] {\\n    display: none;\\n    content-visibility: hidden;\\n  }\\n}\\n\\n.tab {\\n  position: relative;\\n  display: flex;\\n  flex: none;\\n  align-items: center;\\n  padding: calc(var(--pf-t--global--spacer--sm) + var(--pf-t--global--spacer--xs))\\n           calc(var(--pf-t--global--spacer--sm) + var(--pf-t--global--spacer--sm));\\n  font-family: inherit;\\n  font-size: var(--pf-t--global--font--size--sm);\\n  font-weight: 400;\\n  color: var(--_tab--Color);\\n  text-decoration: none;\\n  background-color: transparent;\\n  border: 0;\\n  cursor: pointer;\\n\\n  \\u0026::before {\\n    content: '';\\n    position: absolute;\\n    inset-block-start: 0;\\n    inset-inline-start: 0;\\n    width: calc(100% - calc(var(--pf-t--global--spacer--sm) * 2));\\n    height: calc(100% - calc(var(--pf-t--global--spacer--sm) * 2));\\n    translate: var(--pf-t--global--spacer--sm) var(--pf-t--global--spacer--sm);\\n    background-color: var(--_tab--BackgroundColor);\\n    border: var(--_tab--BorderWidth) solid var(--_tab--BorderColor);\\n    border-radius: var(--pf-t--global--border--radius--small);\\n    transition: background-color var(--pf-t--global--motion--duration--fade--short) var(--pf-t--global--motion--timing-function--default);\\n    z-index: -1;\\n  }\\n\\n  \\u0026:where(:hover, :focus) {\\n    --_tab--Color: var(--_tab--hover--Color);\\n    --_tab--BackgroundColor: var(--_tab--hover--BackgroundColor);\\n  }\\n\\n  \\u0026[aria-selected=\\"true\\"] {\\n    --_tab--Color: var(--_tab--current--Color);\\n    --_tab--BackgroundColor: var(--_tab--current--BackgroundColor);\\n  }\\n\\n  \\u0026:disabled {\\n    --_tab--Color: var(--_tab--disabled--Color);\\n    --_tab--BackgroundColor: var(--_tab--disabled--BackgroundColor);\\n    cursor: not-allowed;\\n    pointer-events: none;\\n  }\\n}\\n"`));
var cem_pf_v6_tabs_default = s47;

// elements/cem-pf-v6-tabs/cem-pf-v6-tabs.ts
var PfTabsChangeEvent = class extends Event {
  selectedIndex;
  constructor(selectedIndex) {
    super("change", { bubbles: true });
    this.selectedIndex = selectedIndex;
  }
};
var __tabs_dec, _selected_dec, _a44, _PfV6Tabs_decorators, _init44, _selected, __tabs, _mutationObserver, _onTabsChanged, _PfV6Tabs_instances, syncTabs_fn, handleTabClick_fn, handleKeyDown_fn, updateAccentLine_fn;
_PfV6Tabs_decorators = [t3("cem-pf-v6-tabs")];
var PfV6Tabs = class extends (_a44 = i3, _selected_dec = [n4({ type: Number, reflect: true })], __tabs_dec = [r5()], _a44) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6Tabs_instances);
    __privateAdd(this, _selected, __runInitializers(_init44, 8, this, 0)), __runInitializers(_init44, 11, this);
    __privateAdd(this, __tabs, __runInitializers(_init44, 12, this, [])), __runInitializers(_init44, 15, this);
    __privateAdd(this, _mutationObserver);
    __privateAdd(this, _onTabsChanged, () => {
      __privateMethod(this, _PfV6Tabs_instances, syncTabs_fn).call(this);
    });
  }
  get selectedIndex() {
    return this.selected;
  }
  set selectedIndex(index) {
    this.selected = Math.max(0, Math.min(index, this._tabs.length - 1));
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("cem-pf-v6-tab-title-changed", __privateGet(this, _onTabsChanged));
    this.addEventListener("cem-pf-v6-tab-connected", __privateGet(this, _onTabsChanged));
    this.addEventListener("cem-pf-v6-tab-disconnected", __privateGet(this, _onTabsChanged));
    __privateSet(this, _mutationObserver, new MutationObserver(__privateGet(this, _onTabsChanged)));
    __privateGet(this, _mutationObserver).observe(this, { childList: true });
  }
  firstUpdated() {
    __privateMethod(this, _PfV6Tabs_instances, syncTabs_fn).call(this);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    __privateGet(this, _mutationObserver)?.disconnect();
    this.removeEventListener("cem-pf-v6-tab-title-changed", __privateGet(this, _onTabsChanged));
    this.removeEventListener("cem-pf-v6-tab-connected", __privateGet(this, _onTabsChanged));
    this.removeEventListener("cem-pf-v6-tab-disconnected", __privateGet(this, _onTabsChanged));
  }
  updated(changed) {
    if (changed.has("selected") || changed.has("_tabs")) {
      this._tabs.forEach((tab, i6) => {
        tab.setAttribute("slot", `panel-${i6}`);
      });
      requestAnimationFrame(() => __privateMethod(this, _PfV6Tabs_instances, updateAccentLine_fn).call(this));
    }
  }
  render() {
    return T`
      <div id="tabs-container">
        <div id="tabs"
             role="tablist"
             aria-label="Tabs"
             part="tabs"
             @keydown=${__privateMethod(this, _PfV6Tabs_instances, handleKeyDown_fn)}>
          ${this._tabs.map((tab, index) => T`
            <button class="tab"
                    type="button"
                    role="tab"
                    id="tab-${index}"
                    aria-controls="panel-${index}"
                    aria-selected="${index === this.selected ? "true" : "false"}"
                    tabindex="${index === this.selected ? 0 : -1}"
                    @click=${() => __privateMethod(this, _PfV6Tabs_instances, handleTabClick_fn).call(this, index)}>
              ${tab.title || `Tab ${index + 1}`}
            </button>
          `)}
        </div>
        <div id="panels"
             part="panels">
          ${this._tabs.map((_2, index) => T`
            <div class="panel"
                 id="panel-${index}"
                 role="tabpanel"
                 aria-labelledby="tab-${index}"
                 ?hidden=${index !== this.selected}>
              <slot name="panel-${index}"></slot>
            </div>
          `)}
        </div>
      </div>
    `;
  }
};
_init44 = __decoratorStart(_a44);
_selected = new WeakMap();
__tabs = new WeakMap();
_mutationObserver = new WeakMap();
_onTabsChanged = new WeakMap();
_PfV6Tabs_instances = new WeakSet();
syncTabs_fn = function() {
  this._tabs = Array.from(this.querySelectorAll("cem-pf-v6-tab"));
  this.requestUpdate();
};
handleTabClick_fn = function(index) {
  const oldIndex = this.selected;
  this.selected = index;
  if (oldIndex !== index) {
    this.dispatchEvent(new PfTabsChangeEvent(index));
  }
};
handleKeyDown_fn = function(e6) {
  const tabsEl = this.shadowRoot?.getElementById("tabs");
  if (!tabsEl) return;
  const buttons = Array.from(tabsEl.querySelectorAll(".tab"));
  const currentIndex = buttons.findIndex((btn) => btn === this.shadowRoot?.activeElement);
  if (currentIndex === -1) return;
  let nextIndex = currentIndex;
  switch (e6.key) {
    case "ArrowLeft":
      e6.preventDefault();
      nextIndex = currentIndex > 0 ? currentIndex - 1 : buttons.length - 1;
      break;
    case "ArrowRight":
      e6.preventDefault();
      nextIndex = currentIndex < buttons.length - 1 ? currentIndex + 1 : 0;
      break;
    case "Home":
      e6.preventDefault();
      nextIndex = 0;
      break;
    case "End":
      e6.preventDefault();
      nextIndex = buttons.length - 1;
      break;
    default:
      return;
  }
  __privateMethod(this, _PfV6Tabs_instances, handleTabClick_fn).call(this, nextIndex);
  buttons[nextIndex].focus();
};
updateAccentLine_fn = function() {
  const tabsEl = this.shadowRoot?.getElementById("tabs");
  if (!tabsEl) return;
  const buttons = Array.from(tabsEl.querySelectorAll(".tab"));
  const activeButton = buttons[this.selected];
  if (!activeButton) return;
  const containerRect = tabsEl.getBoundingClientRect();
  const buttonRect = activeButton.getBoundingClientRect();
  const inset = 8;
  const start = buttonRect.left - containerRect.left + inset;
  const length = buttonRect.width - inset * 2;
  tabsEl.style.setProperty("--_cem-pf-v6-tabs--link-accent--start", `${start}px`);
  tabsEl.style.setProperty("--_cem-pf-v6-tabs--link-accent--length", `${length}px`);
};
__decorateElement(_init44, 4, "selected", _selected_dec, PfV6Tabs, _selected);
__decorateElement(_init44, 4, "_tabs", __tabs_dec, PfV6Tabs, __tabs);
PfV6Tabs = __decorateElement(_init44, 0, "PfV6Tabs", _PfV6Tabs_decorators, PfV6Tabs);
__publicField(PfV6Tabs, "styles", cem_pf_v6_tabs_default);
__runInitializers(_init44, 1, PfV6Tabs);

// lit-css:elements/cem-pf-v6-text-input/cem-pf-v6-text-input.css
var s48 = new CSSStyleSheet();
s48.replaceSync(JSON.parse('":host {\\n\\n  /* Form control custom properties */\\n  --cem-pf-v6-c-form-control--ColumnGap: var(--pf-t--global--spacer--gap--text-to-element--default);\\n  --cem-pf-v6-c-form-control--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-form-control--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-form-control--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-form-control--Resize: none;\\n  --cem-pf-v6-c-form-control--OutlineOffset: -6px;\\n  --cem-pf-v6-c-form-control--BorderRadius: var(--pf-t--global--border--radius--small);\\n  --cem-pf-v6-c-form-control--before--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-form-control--before--BorderStyle: solid;\\n  --cem-pf-v6-c-form-control--before--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-form-control--before--BorderRadius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  --cem-pf-v6-c-form-control--after--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-form-control--after--BorderStyle: solid;\\n  --cem-pf-v6-c-form-control--after--BorderColor: transparent;\\n  --cem-pf-v6-c-form-control--after--BorderRadius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  --cem-pf-v6-c-form-control--BackgroundColor: var(--pf-t--global--background--color--control--default);\\n  --cem-pf-v6-c-form-control--Width: 100%;\\n  --cem-pf-v6-c-form-control--inset--base: var(--pf-t--global--spacer--control--horizontal--default);\\n  --cem-pf-v6-c-form-control--PaddingBlockStart--base: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-form-control--PaddingBlockEnd--base: var(--pf-t--global--spacer--control--vertical--default);\\n  --cem-pf-v6-c-form-control--PaddingInlineEnd--base: var(--cem-pf-v6-c-form-control--inset--base);\\n  --cem-pf-v6-c-form-control--PaddingInlineStart--base: var(--cem-pf-v6-c-form-control--inset--base);\\n  --cem-pf-v6-c-form-control--PaddingBlockStart: var(--cem-pf-v6-c-form-control__input--PaddingBlockStart);\\n  --cem-pf-v6-c-form-control--PaddingBlockEnd: var(--cem-pf-v6-c-form-control__input--PaddingBlockEnd);\\n  --cem-pf-v6-c-form-control--PaddingInlineStart: var(--cem-pf-v6-c-form-control__input--PaddingInlineStart);\\n  --cem-pf-v6-c-form-control--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__input--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__utilities--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__utilities--input--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__utilities--input--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__input--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__utilities--select--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__select--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__utilities--textarea--PaddingInlineEnd: var(--cem-pf-v6-c-form-control__textarea--PaddingInlineEnd);\\n  --cem-pf-v6-c-form-control__input--PaddingBlockStart: var(--cem-pf-v6-c-form-control--PaddingBlockStart--base);\\n  --cem-pf-v6-c-form-control__input--PaddingBlockEnd: var(--cem-pf-v6-c-form-control--PaddingBlockEnd--base);\\n  --cem-pf-v6-c-form-control__input--PaddingInlineEnd: var(--cem-pf-v6-c-form-control--PaddingInlineEnd--base);\\n  --cem-pf-v6-c-form-control__input--PaddingInlineStart: var(--cem-pf-v6-c-form-control--PaddingInlineStart--base);\\n  --cem-pf-v6-c-form-control__select--PaddingBlockStart: var(--cem-pf-v6-c-form-control--PaddingBlockStart--base);\\n  --cem-pf-v6-c-form-control__select--PaddingBlockEnd: var(--cem-pf-v6-c-form-control--PaddingBlockEnd--base);\\n  --cem-pf-v6-c-form-control__select--PaddingInlineEnd: var(--cem-pf-v6-c-form-control--PaddingInlineEnd--base);\\n  --cem-pf-v6-c-form-control__select--PaddingInlineStart: var(--cem-pf-v6-c-form-control--PaddingInlineStart--base);\\n  --cem-pf-v6-c-form-control__textarea--PaddingBlockStart: var(--cem-pf-v6-c-form-control--PaddingBlockStart--base);\\n  --cem-pf-v6-c-form-control__textarea--PaddingBlockEnd: var(--cem-pf-v6-c-form-control--PaddingBlockEnd--base);\\n  --cem-pf-v6-c-form-control__textarea--PaddingInlineEnd: var(--cem-pf-v6-c-form-control--PaddingInlineEnd--base);\\n  --cem-pf-v6-c-form-control__textarea--PaddingInlineStart: var(--cem-pf-v6-c-form-control--PaddingInlineStart--base);\\n  --cem-pf-v6-c-form-control--hover--after--BorderWidth: var(--pf-t--global--border--width--control--hover);\\n  --cem-pf-v6-c-form-control--hover--after--BorderColor: var(--pf-t--global--border--color--hover);\\n  --cem-pf-v6-c-form-control--m-expanded--after--BorderWidth: var(--pf-t--global--border--width--control--clicked);\\n  --cem-pf-v6-c-form-control--m-expanded--after--BorderColor: var(--pf-t--global--border--color--clicked);\\n  --cem-pf-v6-c-form-control--m-placeholder--Color: var(--pf-t--global--text--color--placeholder);\\n  --cem-pf-v6-c-form-control--m-placeholder--child--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-form-control--m-disabled--Color: var(--pf-t--global--text--color--on-disabled);\\n  --cem-pf-v6-c-form-control--m-disabled--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n  --cem-pf-v6-c-form-control--m-disabled--after--BorderColor: transparent;\\n  --cem-pf-v6-c-form-control--m-readonly--BackgroundColor: var(--pf-t--global--background--color--control--read-only);\\n  --cem-pf-v6-c-form-control--m-readonly--BorderColor: var(--pf-t--global--border--color--control--read-only);\\n  --cem-pf-v6-c-form-control--m-readonly--hover--after--BorderColor: revert;\\n  --cem-pf-v6-c-form-control--m-readonly--m-plain--BackgroundColor: transparent;\\n  --cem-pf-v6-c-form-control--m-readonly--m-plain--BorderColor: transparent;\\n  --cem-pf-v6-c-form-control--m-readonly--m-plain--inset--base: 0;\\n  --cem-pf-v6-c-form-control--m-readonly--m-plain--OutlineOffset: 0;\\n\\n  /* Structural styles */\\n  position: relative;\\n  display: grid;\\n  grid-template-columns: 1fr auto;\\n  column-gap: var(--cem-pf-v6-c-form-control--ColumnGap);\\n  align-items: start;\\n  width: var(--cem-pf-v6-c-form-control--Width);\\n  font-size: var(--cem-pf-v6-c-form-control--FontSize);\\n  line-height: var(--cem-pf-v6-c-form-control--LineHeight);\\n  resize: var(--cem-pf-v6-c-form-control--Resize);\\n  background-color: var(--cem-pf-v6-c-form-control--BackgroundColor);\\n  border-radius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n}\\n\\n:host::before,\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n}\\n\\n:host::before {\\n  border-color: var(--cem-pf-v6-c-form-control--before--BorderColor);\\n  border-style: var(--cem-pf-v6-c-form-control--before--BorderStyle);\\n  border-width: var(--cem-pf-v6-c-form-control--before--BorderWidth);\\n  border-radius: var(--cem-pf-v6-c-form-control--before--BorderRadius);\\n}\\n\\n:host::after {\\n  border: var(--cem-pf-v6-c-form-control--after--BorderWidth) var(--cem-pf-v6-c-form-control--after--BorderStyle) var(--cem-pf-v6-c-form-control--after--BorderColor);\\n  border-radius: var(--cem-pf-v6-c-form-control--before--BorderRadius);\\n}\\n\\n:host(:hover) {\\n  --cem-pf-v6-c-form-control--after--BorderColor: var(--cem-pf-v6-c-form-control--hover--after--BorderColor);\\n  --cem-pf-v6-c-form-control--after--BorderWidth: var(--cem-pf-v6-c-form-control--hover--after--BorderWidth);\\n}\\n\\n#text-input {\\n  grid-row: 1 / 2;\\n  grid-column: 1 / -1;\\n  padding-block-start: var(--cem-pf-v6-c-form-control--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-form-control--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-form-control--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-form-control--PaddingInlineEnd);\\n  margin: 0;\\n  font-family: inherit;\\n  font-size: 100%;\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--cem-pf-v6-c-form-control--Color);\\n  appearance: none;\\n  background-color: transparent;\\n  border: none;\\n  border-radius: var(--cem-pf-v6-c-form-control--BorderRadius);\\n  outline-offset: var(--cem-pf-v6-c-form-control--OutlineOffset);\\n  text-overflow: ellipsis;\\n}\\n\\n#text-input::placeholder {\\n  color: var(--cem-pf-v6-c-form-control--m-placeholder--Color);\\n}\\n"'));
var cem_pf_v6_text_input_default = s48;

// elements/cem-pf-v6-text-input/cem-pf-v6-text-input.ts
var _step_dec, _max_dec2, _min_dec2, _invalid_dec2, _readonly_dec2, _disabled_dec9, _placeholder_dec2, _value_dec5, _type_dec3, _a45, _PfV6TextInput_decorators, _internals11, _init45, _type3, _value5, _placeholder2, _disabled9, _readonly2, _invalid2, _min2, _max2, _step, _PfV6TextInput_instances, onInput_fn3, onChange_fn2;
_PfV6TextInput_decorators = [t3("cem-pf-v6-text-input")];
var PfV6TextInput = class extends (_a45 = i3, _type_dec3 = [n4({ reflect: true })], _value_dec5 = [n4()], _placeholder_dec2 = [n4()], _disabled_dec9 = [n4({ type: Boolean, reflect: true })], _readonly_dec2 = [n4({ type: Boolean, reflect: true })], _invalid_dec2 = [n4({ type: Boolean, reflect: true })], _min_dec2 = [n4()], _max_dec2 = [n4()], _step_dec = [n4()], _a45) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _PfV6TextInput_instances);
    __privateAdd(this, _internals11, this.attachInternals());
    __privateAdd(this, _type3, __runInitializers(_init45, 8, this, "text")), __runInitializers(_init45, 11, this);
    __privateAdd(this, _value5, __runInitializers(_init45, 12, this, "")), __runInitializers(_init45, 15, this);
    __privateAdd(this, _placeholder2, __runInitializers(_init45, 16, this)), __runInitializers(_init45, 19, this);
    __privateAdd(this, _disabled9, __runInitializers(_init45, 20, this, false)), __runInitializers(_init45, 23, this);
    __privateAdd(this, _readonly2, __runInitializers(_init45, 24, this, false)), __runInitializers(_init45, 27, this);
    __privateAdd(this, _invalid2, __runInitializers(_init45, 28, this, false)), __runInitializers(_init45, 31, this);
    __privateAdd(this, _min2, __runInitializers(_init45, 32, this)), __runInitializers(_init45, 35, this);
    __privateAdd(this, _max2, __runInitializers(_init45, 36, this)), __runInitializers(_init45, 39, this);
    __privateAdd(this, _step, __runInitializers(_init45, 40, this)), __runInitializers(_init45, 43, this);
  }
  get valueAsNumber() {
    const input = this.shadowRoot?.getElementById("text-input");
    return input?.valueAsNumber ?? NaN;
  }
  render() {
    return T`
      <input id="text-input"
             .type=${this.type}
             .value=${l3(this.value)}
             placeholder=${o7(this.placeholder)}
             ?disabled=${this.disabled}
             ?readonly=${this.readonly}
             aria-invalid=${o7(this.invalid ? "true" : void 0)}
             min=${o7(this.type === "number" ? this.min : void 0)}
             max=${o7(this.type === "number" ? this.max : void 0)}
             step=${o7(this.type === "number" ? this.step : void 0)}
             @input=${__privateMethod(this, _PfV6TextInput_instances, onInput_fn3)}
             @change=${__privateMethod(this, _PfV6TextInput_instances, onChange_fn2)}>
    `;
  }
  /**
   * Sets the accessible label on the internal input.
   * Called by cem-pf-v6-form-label when label text changes.
   */
  setAccessibleLabel(labelText) {
    const input = this.shadowRoot?.getElementById("text-input");
    if (input && labelText) {
      input.setAttribute("aria-label", labelText);
    }
  }
  select() {
    const input = this.shadowRoot?.getElementById("text-input");
    input?.select();
  }
  checkValidity() {
    const input = this.shadowRoot?.getElementById("text-input");
    return input?.checkValidity() ?? true;
  }
  reportValidity() {
    const input = this.shadowRoot?.getElementById("text-input");
    return input?.reportValidity() ?? true;
  }
  setCustomValidity(message) {
    const input = this.shadowRoot?.getElementById("text-input");
    input?.setCustomValidity(message);
  }
};
_init45 = __decoratorStart(_a45);
_internals11 = new WeakMap();
_type3 = new WeakMap();
_value5 = new WeakMap();
_placeholder2 = new WeakMap();
_disabled9 = new WeakMap();
_readonly2 = new WeakMap();
_invalid2 = new WeakMap();
_min2 = new WeakMap();
_max2 = new WeakMap();
_step = new WeakMap();
_PfV6TextInput_instances = new WeakSet();
onInput_fn3 = function() {
  const input = this.shadowRoot?.getElementById("text-input");
  if (input) {
    this.value = input.value;
    __privateGet(this, _internals11).setFormValue(input.value);
    this.dispatchEvent(new Event("input", { bubbles: true }));
  }
};
onChange_fn2 = function() {
  const input = this.shadowRoot?.getElementById("text-input");
  if (input) {
    this.value = input.value;
    __privateGet(this, _internals11).setFormValue(input.value);
    this.dispatchEvent(new Event("change", { bubbles: true }));
  }
};
__decorateElement(_init45, 4, "type", _type_dec3, PfV6TextInput, _type3);
__decorateElement(_init45, 4, "value", _value_dec5, PfV6TextInput, _value5);
__decorateElement(_init45, 4, "placeholder", _placeholder_dec2, PfV6TextInput, _placeholder2);
__decorateElement(_init45, 4, "disabled", _disabled_dec9, PfV6TextInput, _disabled9);
__decorateElement(_init45, 4, "readonly", _readonly_dec2, PfV6TextInput, _readonly2);
__decorateElement(_init45, 4, "invalid", _invalid_dec2, PfV6TextInput, _invalid2);
__decorateElement(_init45, 4, "min", _min_dec2, PfV6TextInput, _min2);
__decorateElement(_init45, 4, "max", _max_dec2, PfV6TextInput, _max2);
__decorateElement(_init45, 4, "step", _step_dec, PfV6TextInput, _step);
PfV6TextInput = __decorateElement(_init45, 0, "PfV6TextInput", _PfV6TextInput_decorators, PfV6TextInput);
__publicField(PfV6TextInput, "formAssociated", true);
__publicField(PfV6TextInput, "shadowRootOptions", {
  ...i3.shadowRootOptions,
  delegatesFocus: true
});
__publicField(PfV6TextInput, "styles", cem_pf_v6_text_input_default);
__runInitializers(_init45, 1, PfV6TextInput);

// lit-css:elements/cem-pf-v6-toggle-group/cem-pf-v6-toggle-group.css
var s49 = new CSSStyleSheet();
s49.replaceSync(JSON.parse('":host {\\n  display: inline-flex;\\n  --cem-pf-v6-c-toggle-group__button--PaddingBlockStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-toggle-group__button--PaddingInlineEnd: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-toggle-group__button--PaddingBlockEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-toggle-group__button--PaddingInlineStart: var(--pf-t--global--spacer--md);\\n  --cem-pf-v6-c-toggle-group__button--FontSize: var(--pf-t--global--font--size--body--default);\\n  --cem-pf-v6-c-toggle-group__button--LineHeight: var(--pf-t--global--font--line-height--body);\\n  --cem-pf-v6-c-toggle-group__button--Color: var(--pf-t--global--text--color--regular);\\n  --cem-pf-v6-c-toggle-group__button--BackgroundColor: var(--pf-t--global--background--color--action--plain--default);\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: auto;\\n  --cem-pf-v6-c-toggle-group__button--hover--BackgroundColor: var(--pf-t--global--background--color--primary--hover);\\n  --cem-pf-v6-c-toggle-group__button--hover--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-toggle-group__button--hover--before--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-toggle-group__button--hover--after--BorderWidth: var(--pf-t--global--border--width--high-contrast--regular);\\n  --cem-pf-v6-c-toggle-group__button--before--BorderWidth: var(--pf-t--global--border--width--control--default);\\n  --cem-pf-v6-c-toggle-group__button--before--BorderColor: var(--pf-t--global--border--color--default);\\n  --cem-pf-v6-c-toggle-group__item--item--MarginInlineStart: calc(-1 * var(--pf-t--global--border--width--control--default));\\n  --cem-pf-v6-c-toggle-group__item--first-child__button--BorderStartStartRadius: var(--pf-t--global--border--radius--tiny);\\n  --cem-pf-v6-c-toggle-group__item--first-child__button--BorderEndStartRadius: var(--pf-t--global--border--radius--tiny);\\n  --cem-pf-v6-c-toggle-group__item--last-child__button--BorderStartEndRadius: var(--pf-t--global--border--radius--tiny);\\n  --cem-pf-v6-c-toggle-group__item--last-child__button--BorderEndEndRadius: var(--pf-t--global--border--radius--tiny);\\n  --cem-pf-v6-c-toggle-group__icon--text--MarginInlineStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-toggle-group__button--m-selected--BackgroundColor: var(--pf-t--global--color--brand--default);\\n  --cem-pf-v6-c-toggle-group__button--m-selected--Color: var(--pf-t--global--text--color--on-brand--default);\\n  --cem-pf-v6-c-toggle-group__button--m-selected--before--BorderColor: var(--pf-t--global--border--color--clicked);\\n  --cem-pf-v6-c-toggle-group__button--m-selected-selected--before--BorderInlineStartColor: var(--pf-t--global--border--color--alt);\\n  --cem-pf-v6-c-toggle-group__button--m-selected--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-toggle-group__button--m-selected--after--BorderWidth: var(--pf-t--global--border--width--high-contrast--strong);\\n  --cem-pf-v6-c-toggle-group__button--disabled--BackgroundColor: var(--pf-t--global--background--color--disabled--default);\\n  --cem-pf-v6-c-toggle-group__button--disabled--Color: var(--pf-t--global--text--color--on-disabled);\\n  --cem-pf-v6-c-toggle-group__button--disabled--before--BorderColor: var(--pf-t--global--border--color--disabled);\\n  --cem-pf-v6-c-toggle-group__button--disabled-disabled--before--BorderInlineStartColor: var(--pf-t--global--border--color--disabled);\\n  --cem-pf-v6-c-toggle-group__button--disabled--ZIndex: var(--pf-t--global--z-index--xs);\\n  --cem-pf-v6-c-toggle-group--m-compact__button--PaddingBlockStart: 0;\\n  --cem-pf-v6-c-toggle-group--m-compact__button--PaddingInlineEnd: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-toggle-group--m-compact__button--PaddingBlockEnd: 0;\\n  --cem-pf-v6-c-toggle-group--m-compact__button--PaddingInlineStart: var(--pf-t--global--spacer--sm);\\n  --cem-pf-v6-c-toggle-group--m-compact__button--FontSize: var(--pf-t--global--font--size--body--default);\\n}\\n\\n#container {\\n  display: flex;\\n}\\n\\n:host([compact]) {\\n  --cem-pf-v6-c-toggle-group__button--PaddingBlockStart: var(--cem-pf-v6-c-toggle-group--m-compact__button--PaddingBlockStart);\\n  --cem-pf-v6-c-toggle-group__button--PaddingInlineEnd: var(--cem-pf-v6-c-toggle-group--m-compact__button--PaddingInlineEnd);\\n  --cem-pf-v6-c-toggle-group__button--PaddingBlockEnd: var(--cem-pf-v6-c-toggle-group--m-compact__button--PaddingBlockEnd);\\n  --cem-pf-v6-c-toggle-group__button--PaddingInlineStart: var(--cem-pf-v6-c-toggle-group--m-compact__button--PaddingInlineStart);\\n  --cem-pf-v6-c-toggle-group__button--FontSize: var(--cem-pf-v6-c-toggle-group--m-compact__button--FontSize);\\n}\\n"'));
var cem_pf_v6_toggle_group_default = s49;

// elements/cem-pf-v6-toggle-group/cem-pf-v6-toggle-group.ts
var ToggleGroupChangeEvent = class extends Event {
  item;
  selected;
  value;
  constructor(item, selected, value) {
    super("cem-pf-v6-toggle-group-change", { bubbles: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
};
var _compact_dec3, _a46, _PfV6ToggleGroup_decorators, _init46, _compact3, _internals12, _handleItemSelect;
_PfV6ToggleGroup_decorators = [t3("cem-pf-v6-toggle-group")];
var PfV6ToggleGroup = class extends (_a46 = i3, _compact_dec3 = [n4({ type: Boolean, reflect: true })], _a46) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _compact3, __runInitializers(_init46, 8, this, false)), __runInitializers(_init46, 11, this);
    __privateAdd(this, _internals12, this.attachInternals());
    __privateAdd(this, _handleItemSelect, (event) => {
      const detail = event;
      const selectedItem = detail.item;
      const isNowSelected = detail.selected;
      if (isNowSelected) {
        const items = this.querySelectorAll("cem-pf-v6-toggle-group-item");
        items.forEach((item) => {
          if (item !== selectedItem && item.hasAttribute("selected")) {
            item.removeAttribute("selected");
          }
        });
      }
      this.dispatchEvent(new ToggleGroupChangeEvent(
        selectedItem,
        isNowSelected,
        selectedItem.getAttribute("value")
      ));
    });
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("cem-pf-v6-toggle-group-item-select", __privateGet(this, _handleItemSelect));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("cem-pf-v6-toggle-group-item-select", __privateGet(this, _handleItemSelect));
  }
  render() {
    return T`
      <div id="container"
           role="radiogroup"
           part="container"
           aria-label=${this.getAttribute("aria-label") ?? ""}>
        <slot></slot>
      </div>
    `;
  }
};
_init46 = __decoratorStart(_a46);
_compact3 = new WeakMap();
_internals12 = new WeakMap();
_handleItemSelect = new WeakMap();
__decorateElement(_init46, 4, "compact", _compact_dec3, PfV6ToggleGroup, _compact3);
PfV6ToggleGroup = __decorateElement(_init46, 0, "PfV6ToggleGroup", _PfV6ToggleGroup_decorators, PfV6ToggleGroup);
__publicField(PfV6ToggleGroup, "styles", cem_pf_v6_toggle_group_default);
__runInitializers(_init46, 1, PfV6ToggleGroup);

// lit-css:elements/cem-pf-v6-toggle-group-item/cem-pf-v6-toggle-group-item.css
var s50 = new CSSStyleSheet();
s50.replaceSync(JSON.parse('":host {\\n  position: relative;\\n  z-index: var(--cem-pf-v6-c-toggle-group__button--ZIndex);\\n  display: inline-flex;\\n  align-items: center;\\n  padding-block-start: var(--cem-pf-v6-c-toggle-group__button--PaddingBlockStart);\\n  padding-block-end: var(--cem-pf-v6-c-toggle-group__button--PaddingBlockEnd);\\n  padding-inline-start: var(--cem-pf-v6-c-toggle-group__button--PaddingInlineStart);\\n  padding-inline-end: var(--cem-pf-v6-c-toggle-group__button--PaddingInlineEnd);\\n  font-size: var(--cem-pf-v6-c-toggle-group__button--FontSize);\\n  line-height: var(--cem-pf-v6-c-toggle-group__button--LineHeight);\\n  color: var(--cem-pf-v6-c-toggle-group__button--Color);\\n  background-color: var(--cem-pf-v6-c-toggle-group__button--BackgroundColor);\\n  border: 0;\\n  cursor: pointer;\\n  user-select: none;\\n}\\n\\n:host(:not(:first-child)) {\\n  margin-inline-start: var(--cem-pf-v6-c-toggle-group__item--item--MarginInlineStart);\\n}\\n\\n:host::before,\\n:host::after {\\n  position: absolute;\\n  inset: 0;\\n  pointer-events: none;\\n  content: \\"\\";\\n  border-style: solid;\\n  border-radius: inherit;\\n}\\n\\n:host::before {\\n  border-width: var(--cem-pf-v6-c-toggle-group__button--before--BorderWidth);\\n  border-block-start-color: var(--cem-pf-v6-c-toggle-group__button--before--BorderBlockStartColor, var(--cem-pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-block-end-color: var(--cem-pf-v6-c-toggle-group__button--before--BorderBlockEndColor, var(--cem-pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-inline-start-color: var(--cem-pf-v6-c-toggle-group__button--before--BorderInlineStartColor, var(--cem-pf-v6-c-toggle-group__button--before--BorderColor));\\n  border-inline-end-color: var(--cem-pf-v6-c-toggle-group__button--before--BorderInlineEndColor, var(--cem-pf-v6-c-toggle-group__button--before--BorderColor));\\n}\\n\\n:host::after {\\n  inset: var(--cem-pf-v6-c-toggle-group__button--before--BorderWidth);\\n  border-color: var(--cem-pf-v6-c-toggle-group__button--after--BorderColor);\\n  border-width: var(--cem-pf-v6-c-toggle-group__button--after--BorderWidth);\\n}\\n\\n:host(:hover:not([disabled])) {\\n  --cem-pf-v6-c-toggle-group__button--BackgroundColor: var(--cem-pf-v6-c-toggle-group__button--hover--BackgroundColor);\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: var(--cem-pf-v6-c-toggle-group__button--hover--ZIndex);\\n  --cem-pf-v6-c-toggle-group__button--before--BorderColor: var(--cem-pf-v6-c-toggle-group__button--hover--before--BorderColor);\\n  --cem-pf-v6-c-toggle-group__button--after--BorderWidth: var(--cem-pf-v6-c-toggle-group__button--hover--after--BorderWidth);\\n}\\n\\n:host(:focus-visible) {\\n  outline: 2px solid var(--pf-t--global--color--brand--default);\\n  outline-offset: 2px;\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: var(--cem-pf-v6-c-toggle-group__button--hover--ZIndex);\\n}\\n\\n:host([selected]) {\\n  --cem-pf-v6-c-toggle-group__button--BackgroundColor: var(--cem-pf-v6-c-toggle-group__button--m-selected--BackgroundColor);\\n  --cem-pf-v6-c-toggle-group__button--Color: var(--cem-pf-v6-c-toggle-group__button--m-selected--Color);\\n  --cem-pf-v6-c-toggle-group__button--before--BorderColor: var(--cem-pf-v6-c-toggle-group__button--m-selected--before--BorderColor);\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: var(--cem-pf-v6-c-toggle-group__button--m-selected--ZIndex);\\n  --cem-pf-v6-c-toggle-group__button--after--BorderWidth: var(--cem-pf-v6-c-toggle-group__button--m-selected--after--BorderWidth);\\n}\\n\\n:host([selected]) + :host([selected]) {\\n  --cem-pf-v6-c-toggle-group__button--before--BorderInlineStartColor: var(--cem-pf-v6-c-toggle-group__button--m-selected-selected--before--BorderInlineStartColor);\\n}\\n\\n:host([disabled]) {\\n  --cem-pf-v6-c-toggle-group__button--BackgroundColor: var(--cem-pf-v6-c-toggle-group__button--disabled--BackgroundColor);\\n  --cem-pf-v6-c-toggle-group__button--Color: var(--cem-pf-v6-c-toggle-group__button--disabled--Color);\\n  --cem-pf-v6-c-toggle-group__button--before--BorderColor: var(--cem-pf-v6-c-toggle-group__button--disabled--before--BorderColor);\\n  --cem-pf-v6-c-toggle-group__button--ZIndex: var(--cem-pf-v6-c-toggle-group__button--disabled--ZIndex);\\n  cursor: not-allowed;\\n  pointer-events: none;\\n}\\n\\n:host([disabled]) + :host([disabled]) {\\n  --cem-pf-v6-c-toggle-group__button--before--BorderInlineStartColor: var(--cem-pf-v6-c-toggle-group__button--disabled-disabled--before--BorderInlineStartColor);\\n}\\n\\n:host(:first-child) {\\n  border-start-start-radius: var(--cem-pf-v6-c-toggle-group__item--first-child__button--BorderStartStartRadius);\\n  border-end-start-radius: var(--cem-pf-v6-c-toggle-group__item--first-child__button--BorderEndStartRadius);\\n}\\n\\n:host(:last-child) {\\n  border-start-end-radius: var(--cem-pf-v6-c-toggle-group__item--last-child__button--BorderStartEndRadius);\\n  border-end-end-radius: var(--cem-pf-v6-c-toggle-group__item--last-child__button--BorderEndEndRadius);\\n}\\n\\n#wrapper {\\n  display: var(--_has-content, inline-flex);\\n  gap: var(--cem-pf-v6-c-toggle-group__icon--text--MarginInlineStart);\\n  align-items: center;\\n}\\n\\n::slotted(svg) {\\n  width: 1em;\\n  height: 1em;\\n  vertical-align: -0.125em;\\n  fill: currentColor;\\n}\\n"'));
var cem_pf_v6_toggle_group_item_default = s50;

// elements/cem-pf-v6-toggle-group-item/cem-pf-v6-toggle-group-item.ts
var ToggleGroupItemSelectEvent = class extends Event {
  item;
  selected;
  value;
  constructor(item, selected, value) {
    super("cem-pf-v6-toggle-group-item-select", { bubbles: true });
    this.item = item;
    this.selected = selected;
    this.value = value;
  }
};
var _value_dec6, _disabled_dec10, _selected_dec2, _a47, _PfV6ToggleGroupItem_decorators, _internals13, _init47, _selected2, _disabled10, _value6, _PfV6ToggleGroupItem_instances, updateTabindex_fn, _handleClick2, _handleKeydown3, _handleFocus, selectItem_fn, focusAndSelect_fn, navigateItems_fn, navigateToEnd_fn, updateRovingTabindex_fn;
_PfV6ToggleGroupItem_decorators = [t3("cem-pf-v6-toggle-group-item")];
var PfV6ToggleGroupItem = class extends (_a47 = i3, _selected_dec2 = [n4({ type: Boolean, reflect: true })], _disabled_dec10 = [n4({ type: Boolean, reflect: true })], _value_dec6 = [n4({ reflect: true })], _a47) {
  constructor() {
    super();
    __privateAdd(this, _PfV6ToggleGroupItem_instances);
    __privateAdd(this, _internals13, this.attachInternals());
    __privateAdd(this, _selected2, __runInitializers(_init47, 8, this, false)), __runInitializers(_init47, 11, this);
    __privateAdd(this, _disabled10, __runInitializers(_init47, 12, this, false)), __runInitializers(_init47, 15, this);
    __privateAdd(this, _value6, __runInitializers(_init47, 16, this)), __runInitializers(_init47, 19, this);
    __privateAdd(this, _handleClick2, (event) => {
      if (this.disabled) {
        event.preventDefault();
        event.stopImmediatePropagation();
        return;
      }
      this.focus();
      __privateMethod(this, _PfV6ToggleGroupItem_instances, selectItem_fn).call(this);
    });
    __privateAdd(this, _handleKeydown3, (event) => {
      if (this.disabled) {
        event.preventDefault();
        return;
      }
      switch (event.key) {
        case " ":
        case "Enter":
          event.preventDefault();
          __privateMethod(this, _PfV6ToggleGroupItem_instances, selectItem_fn).call(this);
          break;
        case "ArrowLeft":
        case "ArrowRight":
          event.preventDefault();
          __privateMethod(this, _PfV6ToggleGroupItem_instances, navigateItems_fn).call(this, event.key === "ArrowLeft" ? -1 : 1);
          break;
        case "Home":
          event.preventDefault();
          __privateMethod(this, _PfV6ToggleGroupItem_instances, navigateToEnd_fn).call(this, true);
          break;
        case "End":
          event.preventDefault();
          __privateMethod(this, _PfV6ToggleGroupItem_instances, navigateToEnd_fn).call(this, false);
          break;
      }
    });
    __privateAdd(this, _handleFocus, () => {
      __privateMethod(this, _PfV6ToggleGroupItem_instances, updateRovingTabindex_fn).call(this);
    });
    __privateGet(this, _internals13).role = "radio";
    if (!this.hasAttribute("tabindex")) {
      this.setAttribute("tabindex", "-1");
    }
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("click", __privateGet(this, _handleClick2));
    this.addEventListener("keydown", __privateGet(this, _handleKeydown3));
    this.addEventListener("focus", __privateGet(this, _handleFocus));
    queueMicrotask(() => {
      const parent = this.parentElement;
      if (parent) {
        const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item"));
        const isFirstItem = items[0] === this;
        if (this.selected || isFirstItem) {
          __privateMethod(this, _PfV6ToggleGroupItem_instances, updateRovingTabindex_fn).call(this);
        }
      }
    });
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("click", __privateGet(this, _handleClick2));
    this.removeEventListener("keydown", __privateGet(this, _handleKeydown3));
    this.removeEventListener("focus", __privateGet(this, _handleFocus));
  }
  updated(changed) {
    if (changed.has("selected")) {
      __privateGet(this, _internals13).ariaChecked = String(this.selected);
      __privateMethod(this, _PfV6ToggleGroupItem_instances, updateRovingTabindex_fn).call(this);
    }
    if (changed.has("disabled")) {
      __privateGet(this, _internals13).ariaDisabled = String(this.disabled);
      __privateMethod(this, _PfV6ToggleGroupItem_instances, updateTabindex_fn).call(this);
    }
  }
  render() {
    return T`
      <span id="wrapper">
        <span id="icon-start"><slot name="icon"></slot></span>
        <span id="text"><slot></slot></span>
        <span id="icon-end"><slot name="icon-end"></slot></span>
      </span>
    `;
  }
};
_init47 = __decoratorStart(_a47);
_internals13 = new WeakMap();
_selected2 = new WeakMap();
_disabled10 = new WeakMap();
_value6 = new WeakMap();
_PfV6ToggleGroupItem_instances = new WeakSet();
updateTabindex_fn = function() {
  if (this.disabled) {
    this.setAttribute("tabindex", "-1");
    return;
  }
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item"));
  const isFirstItem = items[0] === this;
  const hasSelectedItem = items.some((item) => item.hasAttribute("selected"));
  const shouldBeTabbable = this.selected || isFirstItem && !hasSelectedItem;
  this.setAttribute("tabindex", shouldBeTabbable ? "0" : "-1");
};
_handleClick2 = new WeakMap();
_handleKeydown3 = new WeakMap();
_handleFocus = new WeakMap();
selectItem_fn = function() {
  if (!this.selected) {
    this.selected = true;
    this.dispatchEvent(new ToggleGroupItemSelectEvent(
      this,
      true,
      this.getAttribute("value")
    ));
  }
};
focusAndSelect_fn = function(item) {
  item.focus();
  item.selected = true;
  item.dispatchEvent(new ToggleGroupItemSelectEvent(
    item,
    true,
    item.getAttribute("value")
  ));
};
navigateItems_fn = function(direction) {
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item")).filter((item) => !item.disabled);
  const currentIndex = items.indexOf(this);
  if (currentIndex === -1) return;
  let newIndex = currentIndex + direction;
  if (newIndex < 0) newIndex = items.length - 1;
  else if (newIndex >= items.length) newIndex = 0;
  const targetItem = items[newIndex];
  if (targetItem) __privateMethod(this, _PfV6ToggleGroupItem_instances, focusAndSelect_fn).call(this, targetItem);
};
navigateToEnd_fn = function(toStart) {
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item")).filter((item) => !item.disabled);
  const targetItem = toStart ? items[0] : items[items.length - 1];
  if (targetItem) __privateMethod(this, _PfV6ToggleGroupItem_instances, focusAndSelect_fn).call(this, targetItem);
};
updateRovingTabindex_fn = function() {
  const parent = this.parentElement;
  if (!parent) return;
  const items = Array.from(parent.querySelectorAll("cem-pf-v6-toggle-group-item"));
  items.forEach((item) => {
    item.setAttribute("tabindex", item === this ? "0" : "-1");
  });
};
__decorateElement(_init47, 4, "selected", _selected_dec2, PfV6ToggleGroupItem, _selected2);
__decorateElement(_init47, 4, "disabled", _disabled_dec10, PfV6ToggleGroupItem, _disabled10);
__decorateElement(_init47, 4, "value", _value_dec6, PfV6ToggleGroupItem, _value6);
PfV6ToggleGroupItem = __decorateElement(_init47, 0, "PfV6ToggleGroupItem", _PfV6ToggleGroupItem_decorators, PfV6ToggleGroupItem);
__publicField(PfV6ToggleGroupItem, "styles", cem_pf_v6_toggle_group_item_default);
__runInitializers(_init47, 1, PfV6ToggleGroupItem);

// lit-css:elements/cem-pf-v6-tree-view/cem-pf-v6-tree-view.css
var s51 = new CSSStyleSheet();
s51.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\n#tree {\\n  margin: 0;\\n  padding: 0;\\n  list-style: none;\\n}\\n"'));
var cem_pf_v6_tree_view_default = s51;

// elements/cem-pf-v6-tree-view/cem-pf-v6-tree-view.ts
var _PfV6TreeView_decorators, _init48, _a48;
_PfV6TreeView_decorators = [t3("cem-pf-v6-tree-view")];
var PfV6TreeView = class extends (_a48 = i3) {
  static styles = cem_pf_v6_tree_view_default;
  #currentSelection = null;
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("select", this.#onSelect);
    this.addEventListener("keydown", this.#onKeydown);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("select", this.#onSelect);
    this.removeEventListener("keydown", this.#onKeydown);
  }
  firstUpdated() {
    this.#initializeTabindex();
  }
  render() {
    return T`
      <ul id="tree"
          role="tree"
          part="tree"
          aria-label=${this.getAttribute("aria-label") ?? ""}>
        <slot></slot>
      </ul>
    `;
  }
  #onSelect = (event) => {
    const item = event.target;
    if (this.#currentSelection && this.#currentSelection !== item) {
      this.#currentSelection.deselect?.();
    }
    this.#currentSelection = item;
  };
  #onKeydown = (event) => {
    const target = event.target;
    if (target.tagName !== "CEM-PF-V6-TREE-ITEM") return;
    const treeItem = target;
    const items = this.#getAllVisibleItems();
    const currentIndex = items.indexOf(target);
    switch (event.key) {
      case "ArrowDown":
        event.preventDefault();
        if (currentIndex < items.length - 1) {
          this.#focusItem(items[currentIndex + 1]);
        }
        break;
      case "ArrowUp":
        event.preventDefault();
        if (currentIndex > 0) {
          this.#focusItem(items[currentIndex - 1]);
        }
        break;
      case "ArrowRight":
        event.preventDefault();
        if (treeItem.hasChildren) {
          if (!treeItem.expanded) {
            treeItem.expand?.();
          } else {
            const children = this.#getDirectChildren(target);
            if (children.length > 0) {
              this.#focusItem(children[0]);
            }
          }
        }
        break;
      case "ArrowLeft":
        event.preventDefault();
        if (treeItem.hasChildren && treeItem.expanded) {
          treeItem.collapse?.();
        } else {
          const parent = this.#getParentItem(target);
          if (parent) {
            this.#focusItem(parent);
          }
        }
        break;
      case "Home":
        event.preventDefault();
        if (items.length > 0) {
          this.#focusItem(items[0]);
        }
        break;
      case "End":
        event.preventDefault();
        if (items.length > 0) {
          this.#focusItem(items[items.length - 1]);
        }
        break;
      case "Enter":
      case " ":
        event.preventDefault();
        treeItem.select?.();
        if (treeItem.hasChildren) {
          treeItem.toggle?.();
        }
        break;
      case "*": {
        event.preventDefault();
        const parent = this.#getParentItem(target);
        const siblings = parent ? this.#getDirectChildren(parent) : this.#getTopLevelItems();
        siblings.forEach((item) => {
          if (item.hasChildren) {
            item.expand?.();
          }
        });
        break;
      }
    }
  };
  #getAllVisibleItems() {
    const visible = [];
    const walk = (parent) => {
      const children = Array.from(parent.children).filter((el) => el.tagName === "CEM-PF-V6-TREE-ITEM");
      for (const item of children) {
        visible.push(item);
        const treeItem = item;
        if (treeItem.expanded && treeItem.hasChildren) {
          walk(item);
        }
      }
    };
    walk(this);
    return visible;
  }
  #getDirectChildren(item) {
    return Array.from(item.children).filter((el) => el.tagName === "CEM-PF-V6-TREE-ITEM");
  }
  #getTopLevelItems() {
    return Array.from(this.children).filter((el) => el.tagName === "CEM-PF-V6-TREE-ITEM");
  }
  #getParentItem(item) {
    let current = item.parentElement;
    while (current && current !== this) {
      if (current.tagName === "CEM-PF-V6-TREE-ITEM") {
        return current;
      }
      current = current.parentElement;
    }
    return null;
  }
  #initializeTabindex() {
    const topLevelItems = this.#getTopLevelItems();
    if (topLevelItems.length === 0) return;
    const firstItem = topLevelItems[0];
    firstItem.setTabindex?.(0);
    const allItems = this.querySelectorAll("cem-pf-v6-tree-item");
    allItems.forEach((item) => {
      if (item !== firstItem) {
        item.setTabindex?.(-1);
      }
    });
  }
  #focusItem(item) {
    if (!item) return;
    const allItems = this.querySelectorAll("cem-pf-v6-tree-item");
    allItems.forEach((i6) => i6.setTabindex?.(-1));
    item.setTabindex?.(0);
    item.focusItem?.();
  }
  expandAll() {
    this.querySelectorAll("cem-pf-v6-tree-item").forEach(
      (item) => item.expand?.()
    );
  }
  collapseAll() {
    this.querySelectorAll("cem-pf-v6-tree-item").forEach(
      (item) => item.collapse?.()
    );
  }
};
_init48 = __decoratorStart(_a48);
PfV6TreeView = __decorateElement(_init48, 0, "PfV6TreeView", _PfV6TreeView_decorators, PfV6TreeView);
__runInitializers(_init48, 1, PfV6TreeView);

// lit-css:elements/cem-reconnection-content/cem-reconnection-content.css
var s52 = new CSSStyleSheet();
s52.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\np {\\n  margin: 0 0 var(--pf-t--global--spacer--md) 0;\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--regular);\\n}\\n\\n#retry-info {\\n  background: var(--pf-t--global--background--color--secondary--default);\\n  padding: var(--pf-t--global--spacer--md);\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  margin: 0;\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  color: var(--pf-t--global--text--color--subtle);\\n  border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n}\\n"'));
var cem_reconnection_content_default = s52;

// elements/cem-reconnection-content/cem-reconnection-content.ts
var _retryText_dec, _a49, _CemReconnectionContent_decorators, _init49, _retryText, _b3, retryText_get, retryText_set, _CemReconnectionContent_instances, _countdownInterval, _remainingMs;
_CemReconnectionContent_decorators = [t3("cem-reconnection-content")];
var CemReconnectionContent = class extends (_a49 = i3, _retryText_dec = [r5()], _a49) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemReconnectionContent_instances);
    __privateAdd(this, _retryText, __runInitializers(_init49, 8, this, "")), __runInitializers(_init49, 11, this);
    __privateAdd(this, _countdownInterval, null);
    __privateAdd(this, _remainingMs, 0);
  }
  render() {
    return T`
      <p>
        The connection to the development server was lost.
        Automatically retrying connection...
      </p>
      <div id="retry-info">${__privateGet(this, _CemReconnectionContent_instances, retryText_get)}</div>
    `;
  }
  /** Clear the countdown interval */
  clearCountdown() {
    if (__privateGet(this, _countdownInterval)) {
      clearInterval(__privateGet(this, _countdownInterval));
      __privateSet(this, _countdownInterval, null);
    }
  }
  /**
   * Update the retry information with countdown
   * @param attempt - Current retry attempt number
   * @param delay - Delay in milliseconds until next retry
   */
  updateRetryInfo(attempt, delay) {
    this.clearCountdown();
    __privateSet(this, _remainingMs, delay);
    if (__privateGet(this, _remainingMs) <= 0) {
      __privateSet(this, _CemReconnectionContent_instances, `Attempt #${attempt}. Connecting...`, retryText_set);
      return;
    }
    const seconds = Math.ceil(__privateGet(this, _remainingMs) / 1e3);
    __privateSet(this, _CemReconnectionContent_instances, `Attempt #${attempt}. Retrying in ${seconds}s...`, retryText_set);
    __privateSet(this, _countdownInterval, setInterval(() => {
      __privateSet(this, _remainingMs, __privateGet(this, _remainingMs) - 100);
      if (__privateGet(this, _remainingMs) < 100) {
        this.clearCountdown();
        __privateSet(this, _CemReconnectionContent_instances, `Attempt #${attempt}. Connecting...`, retryText_set);
        return;
      }
      const seconds2 = Math.ceil(__privateGet(this, _remainingMs) / 1e3);
      __privateSet(this, _CemReconnectionContent_instances, `Attempt #${attempt}. Retrying in ${seconds2}s...`, retryText_set);
    }, 100));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.clearCountdown();
  }
};
_init49 = __decoratorStart(_a49);
_retryText = new WeakMap();
_CemReconnectionContent_instances = new WeakSet();
_countdownInterval = new WeakMap();
_remainingMs = new WeakMap();
_b3 = __decorateElement(_init49, 20, "#retryText", _retryText_dec, _CemReconnectionContent_instances, _retryText), retryText_get = _b3.get, retryText_set = _b3.set;
CemReconnectionContent = __decorateElement(_init49, 0, "CemReconnectionContent", _CemReconnectionContent_decorators, CemReconnectionContent);
__publicField(CemReconnectionContent, "styles", cem_reconnection_content_default);
__runInitializers(_init49, 1, CemReconnectionContent);

// lit-css:elements/cem-serve-chrome/cem-serve-chrome.css
var s53 = new CSSStyleSheet();
s53.replaceSync(JSON.parse(`":host {\\n  display: block;\\n  height: 100vh;\\n  overflow: hidden;\\n  --cem-pf-v6-c-masthead__logo--Width: max-content;\\n  --cem-pf-v6-c-masthead__toggle--Size: 1rem;\\n}\\n\\n[hidden] {\\n  display: none !important;\\n}\\n\\n/* Masthead logo styles */\\n.masthead-logo {\\n  display: flex;\\n  align-items: center;\\n  text-decoration: none;\\n  color: inherit;\\n  max-height: var(--cem-pf-v6-c-masthead__logo--MaxHeight);\\n  gap: 4px;\\n  \\u0026 img {\\n    display: block;\\n    max-height: var(--cem-pf-v6-c-masthead__logo--MaxHeight);\\n    width: auto;\\n  }\\n  \\u0026 ::slotted([slot=\\"title\\"]) {\\n    margin: 0;\\n    font-size: 1.125rem;\\n    font-weight: 600;\\n    color: var(--pf-t--global--text--color--regular);\\n  }\\n  \\u0026 h1 {\\n    margin: 0;\\n    font-size: 18px;\\n  }\\n}\\n\\n/* Toolbar group alignment */\\ncem-pf-v6-toolbar-group[variant=\\"action-group\\"] {\\n  margin-inline-start: auto;\\n  align-items: center;\\n}\\n\\n.debug-panel {\\n  background: var(--pf-t--global--background--color--primary--default);\\n  border: 1px solid var(--pf-t--global--border--color--default);\\n  border-radius: 6px;\\n  padding: 1.5rem;\\n  max-width: 600px;\\n  width: 90%;\\n  max-height: 80vh;\\n  overflow-y: auto;\\n\\n  h2 {\\n    margin: 0 0 1rem 0;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-size: 1.125rem;\\n    font-weight: 600;\\n  }\\n\\n  dl {\\n    margin: 0;\\n  }\\n\\n  dt {\\n    color: var(--pf-t--global--text--color--subtle);\\n    font-size: 0.875rem;\\n    margin-top: 0.5rem;\\n    font-weight: 500;\\n  }\\n\\n  dd {\\n    margin: 0 0 0.5rem 0;\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n    color: var(--pf-t--global--text--color--regular);\\n  }\\n\\n  details {\\n    margin-top: 1rem;\\n\\n    summary {\\n      cursor: pointer;\\n      color: var(--pf-t--global--text--color--regular);\\n      font-size: 0.875rem;\\n      font-weight: 500;\\n      list-style: none;\\n      display: flex;\\n      align-items: center;\\n      gap: 0.5rem;\\n      user-select: none;\\n\\n      \\u0026::-webkit-details-marker {\\n        display: none;\\n      }\\n\\n      \\u0026::before {\\n        content: '\\\\25B8';\\n        display: inline-block;\\n        transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n        color: var(--pf-t--global--text--color--subtle);\\n      }\\n    }\\n\\n    \\u0026[open] summary::before {\\n      transform: rotate(90deg);\\n    }\\n\\n    pre {\\n      margin-top: 0.5rem;\\n      margin-left: 1.5rem;\\n      padding: 0.5rem;\\n      background: var(--pf-t--global--background--color--secondary--default);\\n      border-radius: 6px;\\n      font-size: 0.875rem;\\n      overflow-x: auto;\\n      color: var(--pf-t--global--text--color--regular);\\n    }\\n  }\\n\\n  .button-container {\\n    display: flex;\\n    gap: 0.5rem;\\n    margin-top: 1rem;\\n  }\\n\\n  p {\\n    color: var(--pf-t--global--text--color--subtle);\\n    font-size: 0.875rem;\\n  }\\n\\n  button {\\n    margin-top: 1rem;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--color--brand--default);\\n    color: var(--pf-t--global--text--color--on-brand);\\n    border: none;\\n    border-radius: 6px;\\n    font-size: 0.875rem;\\n    font-weight: 400;\\n    cursor: pointer;\\n    transition: all 200ms cubic-bezier(0.645, 0.045, 0.355, 1);\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--color--brand--hover);\\n    }\\n  }\\n}\\n\\n/* Content area padding for demo */\\ncem-pf-v6-page-main {\\n  min-height: calc(100dvh - 72px - var(--pf-t--global--spacer--inset--page-chrome));\\n  display: flex;\\n  flex-direction: column;\\n  \\u0026 \\u003e ::slotted(:not([slot=knobs])) {\\n    padding: var(--pf-t--global--spacer--lg);\\n    flex: 1;\\n  }\\n}\\n\\ncem-drawer {\\n  cem-pf-v6-tabs {\\n    cem-pf-v6-tab {\\n      padding-block-end: 0;\\n    }\\n  }\\n}\\n\\n/* Element descriptions in listing */\\n.element-summary {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n.element-description {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n/* Card footer demo navigation */\\n.card-demos {\\n  display: flex;\\n  flex-wrap: wrap;\\n  gap: var(--pf-t--global--spacer--gap--action-to-action--default);\\n  padding: 0;\\n  margin: 0;\\n}\\n\\n.package-name {\\n  color: var(--pf-t--global--text--color--subtle);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n}\\n\\n/* Knobs container - fills tab panel height */\\n#knobs-container {\\n  height: 100%;\\n  display: flex;\\n  flex-direction: column;\\n  \\u0026 ::slotted([slot=\\"knobs\\"]) {\\n    flex: 1;\\n    min-height: 0;\\n    overflow: hidden;\\n  }\\n}\\n\\n.knobs-empty {\\n  color: var(--cem-dev-server-text-muted);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  text-align: center;\\n  padding: var(--cem-dev-server-spacing-lg);\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    padding: 2px 6px;\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n}\\n\\n.instance-tag {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  color: var(--cem-dev-server-accent-color);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.instance-label {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.knob-group {\\n  margin-bottom: var(--cem-dev-server-spacing-lg);\\n\\n  \\u0026:last-child {\\n    margin-bottom: 0;\\n  }\\n}\\n\\n/* PatternFly v6 form - horizontal layout */\\ncem-pf-v6-form[horizontal] cem-pf-v6-form-field-group {\\n  grid-column: span 2\\n}\\n\\n.knob-group-title {\\n  grid-column: 1 / -1;\\n  margin: 0 0 var(--cem-dev-server-spacing-md) 0;\\n  color: var(--cem-dev-server-text-primary);\\n  font-size: var(--cem-dev-server-font-size-base);\\n  font-weight: 600;\\n  border-bottom: 1px solid var(--cem-dev-server-border-color);\\n  padding-bottom: var(--cem-dev-server-spacing-sm);\\n}\\n\\n.knob-control {\\n  margin-bottom: var(--cem-dev-server-spacing-md);\\n}\\n\\n.knob-label {\\n  display: flex;\\n  flex-direction: column;\\n  gap: var(--cem-dev-server-spacing-xs);\\n  cursor: pointer;\\n}\\n\\n.knob-name {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-primary);\\n  font-weight: 500;\\n}\\n\\n.knob-description {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  line-height: 1.5;\\n\\n  p {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n  }\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n\\n  a {\\n    color: var(--cem-dev-server-accent-color);\\n    text-decoration: none;\\n\\n    \\u0026:hover {\\n      text-decoration: underline;\\n    }\\n  }\\n\\n  strong {\\n    font-weight: 600;\\n    color: var(--cem-dev-server-text-primary);\\n  }\\n\\n  ul, ol {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n    padding-left: var(--cem-dev-server-spacing-lg);\\n  }\\n}\\n\\nfooter.pf-m-sticky-bottom {\\n  view-transition-name: dev-server-footer;\\n  position: sticky;\\n  bottom: 0;\\n  background: var(--pf-t--global--background--color--primary--default);\\n  border-top: 1px solid var(--pf-t--global--border--color--default);\\n  z-index: var(--cem-pf-v6-c-page--section--m-sticky-bottom--ZIndex, var(--pf-t--global--z-index--md));\\n  box-shadow: var(--cem-pf-v6-c-page--section--m-sticky-bottom--BoxShadow, var(--pf-t--global--box-shadow--sm--top));\\n}\\n\\n.footer-description {\\n  padding: 1.5rem;\\n\\n  \\u0026.empty {\\n    display: none;\\n  }\\n}\\n\\nfooter ::slotted([slot=\\"description\\"]) {\\n  margin: 0;\\n  color: var(--pf-t--global--text--color--subtle);\\n  line-height: 1.6;\\n  font-size: 0.875rem;\\n\\n  code {\\n    background: var(--pf-t--global--background--color--primary--hover);\\n    padding: 2px 6px;\\n    border-radius: 3px;\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n  }\\n}\\n\\n.logs-wrapper {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n#log-container {\\n  flex-grow: 1;\\n  overflow-y: auto;\\n}\\n\\n.log-entry {\\n  padding: var(--cem-dev-server-spacing-xs) var(--cem-dev-server-spacing-md);\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-sm);\\n  align-items: baseline;\\n  cem-pf-v6-label {\\n    flex-shrink: 0;\\n  }\\n}\\n\\n.log-time,\\n.log-message {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.log-time {\\n  color: var(--cem-dev-server-text-muted);\\n  flex-shrink: 0;\\n  font-size: 11px;\\n}\\n\\n.log-message {\\n  color: var(--cem-dev-server-text-primary);\\n  word-break: break-word;\\n  flex: 1;\\n\\n  .progress-grid {\\n    display: grid;\\n    grid-template-columns: auto 1fr;\\n    column-gap: var(--pf-t--global--spacer--gap--group--horizontal);\\n  }\\n}\\n\\n/* Navigation content (light DOM slotted content for cem-pf-v6-page-sidebar) */\\n.nav-package {\\n  margin-bottom: var(--cem-dev-server-spacing-md);\\n\\n  \\u0026 \\u003e summary {\\n    cursor: pointer;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--background--color--secondary--default);\\n    border-radius: 6px;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-weight: 600;\\n    font-size: 0.875rem;\\n    list-style: none;\\n    transition: background 200ms cubic-bezier(0.4, 0, 0.2, 1);\\n    margin-bottom: 0.5rem;\\n    display: flex;\\n    align-items: center;\\n    gap: 0.5rem;\\n    user-select: none;\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--background--color--secondary--hover);\\n    }\\n\\n    \\u0026::-webkit-details-marker {\\n      display: none;\\n    }\\n\\n    \\u0026::before {\\n      content: '\\\\25B8';\\n      display: inline-block;\\n      transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n      color: var(--pf-t--global--text--color--subtle);\\n    }\\n  }\\n\\n  \\u0026[open] \\u003e summary::before {\\n    transform: rotate(90deg);\\n  }\\n}\\n\\n.nav-element {\\n  margin-bottom: var(--cem-dev-server-spacing-sm);\\n  margin-inline-start: var(--cem-dev-server-spacing-md);\\n\\n  summary {\\n    cursor: pointer;\\n    padding: 0.5rem 1rem;\\n    background: var(--pf-t--global--background--color--secondary--default);\\n    border-radius: 6px;\\n    color: var(--pf-t--global--text--color--regular);\\n    font-family: ui-monospace, 'Cascadia Code', 'Source Code Pro', Menlo, Consolas, 'DejaVu Sans Mono', monospace;\\n    font-size: 0.875rem;\\n    list-style: none;\\n    transition: background 200ms cubic-bezier(0.4, 0, 0.2, 1);\\n    display: flex;\\n    align-items: center;\\n    gap: 0.5rem;\\n    user-select: none;\\n\\n    \\u0026:hover {\\n      background: var(--pf-t--global--background--color--secondary--hover);\\n    }\\n\\n    \\u0026::-webkit-details-marker {\\n      display: none;\\n    }\\n\\n    \\u0026::before {\\n      content: '\\\\25B8';\\n      display: inline-block;\\n      transition: transform 100ms cubic-bezier(0.4, 0, 0.2, 1);\\n      color: var(--pf-t--global--text--color--subtle);\\n    }\\n  }\\n\\n  \\u0026[open] summary::before {\\n    transform: rotate(90deg);\\n  }\\n}\\n\\n.nav-element-title {\\n  user-select: none;\\n}\\n\\n.nav-demo-list {\\n  list-style: none;\\n  padding: 0;\\n  margin: var(--cem-dev-server-spacing-sm) 0 0 0;\\n  display: grid;\\n  gap: var(--cem-dev-server-spacing-xs);\\n}\\n\\n.nav-demo-link {\\n  color: var(--cem-dev-server-text-primary);\\n  text-decoration: none;\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  padding-inline-start: calc(var(--cem-dev-server-spacing-md) * 2);\\n  background: var(--cem-dev-server-bg-tertiary);\\n  border-radius: var(--cem-dev-server-border-radius);\\n  display: block;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  transition: background 0.2s, color 0.2s;\\n\\n  \\u0026:hover {\\n    background: var(--cem-dev-server-accent-color);\\n    color: var(--pf-t--global--text--color--on-brand);\\n\\n    .nav-package-name {\\n      color: rgba(255, 255, 255, 0.8);\\n    }\\n  }\\n\\n  \\u0026[aria-current=\\"page\\"] {\\n    background: var(--cem-dev-server-accent-color);\\n    color: var(--pf-t--global--text--color--on-brand);\\n\\n    .nav-package-name {\\n      color: rgba(255, 255, 255, 0.8);\\n    }\\n  }\\n}\\n\\n.nav-package-name {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n/* Info button popover triggers in knobs - override plain button padding */\\ncem-pf-v6-popover cem-pf-v6-button[variant=\\"plain\\"] {\\n  --cem-pf-v6-c-button--m-plain--PaddingInlineEnd: 0;\\n  --cem-pf-v6-c-button--m-plain--PaddingInlineStart: 0;\\n  --cem-pf-v6-c-button--MinWidth: auto;\\n}\\n\\n/* Knob description content (slotted in form group helper) */\\ncem-pf-v6-form-group [slot=\\"helper\\"] {\\n  p {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n  }\\n\\n  code {\\n    background: var(--cem-dev-server-bg-tertiary);\\n    border-radius: 3px;\\n    font-family: var(--cem-dev-server-font-family-mono);\\n  }\\n\\n  a {\\n    color: var(--cem-dev-server-accent-color);\\n    text-decoration: none;\\n\\n    \\u0026:hover {\\n      text-decoration: underline;\\n    }\\n  }\\n\\n  strong {\\n    font-weight: 600;\\n  }\\n\\n  ul, ol {\\n    margin: var(--cem-dev-server-spacing-xs) 0;\\n    padding-left: var(--cem-dev-server-spacing-lg);\\n  }\\n}\\n\\n/* Syntax highlighting (chroma - themable via CSS custom properties) */\\ncem-pf-v6-form-group [slot=\\"helper\\"] {\\n  .chroma {\\n    background-color: var(--cem-dev-server-bg-tertiary);\\n    padding: var(--cem-dev-server-spacing-sm);\\n    border-radius: var(--cem-dev-server-border-radius);\\n    overflow-x: auto;\\n\\n    \\u0026 .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }\\n    \\u0026 .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; }\\n    \\u0026 .hl { background-color: var(--cem-dev-server-syntax-highlight) }\\n    \\u0026 .lnt,\\n    \\u0026 .ln {\\n      white-space: pre;\\n      user-select: none;\\n      margin-right: 0.4em;\\n      padding: 0 0.4em 0 0.4em;\\n      color: var(--cem-dev-server-text-muted);\\n    }\\n    \\u0026 .line { display: flex; }\\n\\n    /* Keywords */\\n    \\u0026 .k,\\n    \\u0026 .kc,\\n    \\u0026 .kd,\\n    \\u0026 .kn,\\n    \\u0026 .kp,\\n    \\u0026 .kr {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n    \\u0026 .kt { color: var(--cem-dev-server-syntax-type); font-weight: bold; }\\n\\n    /* Names */\\n    \\u0026 .na,\\n    \\u0026 .nb,\\n    \\u0026 .no,\\n    \\u0026 .nv,\\n    \\u0026 .vc,\\n    \\u0026 .vg,\\n    \\u0026 .vi {\\n      color: var(--cem-dev-server-syntax-name);\\n    }\\n    \\u0026 .bp { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .nc { color: var(--cem-dev-server-syntax-class); font-weight: bold; }\\n    \\u0026 .nd { color: var(--cem-dev-server-syntax-decorator); font-weight: bold; }\\n    \\u0026 .ni,\\n    \\u0026 .ss {\\n      color: var(--cem-dev-server-syntax-special);\\n    }\\n    \\u0026 .ne,\\n    \\u0026 .nl {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n    \\u0026 .nf { color: var(--cem-dev-server-syntax-function); font-weight: bold; }\\n    \\u0026 .nn { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .nt { color: var(--cem-dev-server-syntax-tag) }\\n\\n    /* Strings */\\n    \\u0026 .s,\\n    \\u0026 .sa,\\n    \\u0026 .sb,\\n    \\u0026 .sc,\\n    \\u0026 .dl,\\n    \\u0026 .sd,\\n    \\u0026 .s2,\\n    \\u0026 .se,\\n    \\u0026 .sh,\\n    \\u0026 .si,\\n    \\u0026 .sx,\\n    \\u0026 .s1 {\\n      color: var(--cem-dev-server-syntax-string);\\n    }\\n    \\u0026 .sr { color: var(--cem-dev-server-syntax-tag) }\\n\\n    /* Numbers */\\n    \\u0026 .m,\\n    \\u0026 .mb,\\n    \\u0026 .mf,\\n    \\u0026 .mh,\\n    \\u0026 .mi,\\n    \\u0026 .il,\\n    \\u0026 .mo {\\n      color: var(--cem-dev-server-syntax-number);\\n    }\\n\\n    /* Operators */\\n    \\u0026 .o,\\n    \\u0026 .ow {\\n      color: var(--cem-dev-server-syntax-keyword);\\n      font-weight: bold;\\n    }\\n\\n    /* Comments */\\n    \\u0026 .c,\\n    \\u0026 .ch,\\n    \\u0026 .cm,\\n    \\u0026 .c1 {\\n      color: var(--cem-dev-server-text-muted);\\n      font-style: italic;\\n    }\\n    \\u0026 .cs,\\n    \\u0026 .cp,\\n    \\u0026 .cpf {\\n      color: var(--cem-dev-server-text-muted);\\n      font-weight: bold;\\n      font-style: italic;\\n    }\\n\\n    /* Errors */\\n    \\u0026 .err {\\n      color: var(--cem-dev-server-syntax-error);\\n      background-color: var(--cem-dev-server-syntax-error-bg);\\n    }\\n\\n    /* Generics */\\n    \\u0026 .gd {\\n      color: var(--cem-dev-server-syntax-deleted);\\n      background-color: var(--cem-dev-server-syntax-deleted-bg);\\n    }\\n    \\u0026 .ge { font-style: italic; }\\n    \\u0026 .gr { color: var(--cem-dev-server-syntax-error) }\\n    \\u0026 .gh { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gi {\\n      color: var(--cem-dev-server-syntax-inserted);\\n      background-color: var(--cem-dev-server-syntax-inserted-bg);\\n    }\\n    \\u0026 .go { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gp { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gs { font-weight: bold; }\\n    \\u0026 .gu { color: var(--cem-dev-server-text-secondary) }\\n    \\u0026 .gt { color: var(--cem-dev-server-syntax-error) }\\n    \\u0026 .gl { text-decoration: underline; }\\n    \\u0026 .w { color: var(--cem-dev-server-text-muted) }\\n  }\\n}\\n\\n/* Events tab styling - Primary-detail layout */\\n.events-wrapper {\\n  display: flex;\\n  flex-direction: column;\\n  height: 100%;\\n}\\n\\n#event-drawer {\\n  flex: 1;\\n  min-height: 0;\\n}\\n\\n/* Event list (primary panel) */\\n#event-list {\\n  overflow-y: auto;\\n  height: 100%;\\n}\\n\\n.event-list-item {\\n  /* Reset button styles */\\n  appearance: none;\\n  background: none;\\n  border: none;\\n  border-left: 3px solid transparent;\\n  margin: 0;\\n  font: inherit;\\n  color: inherit;\\n  text-align: inherit;\\n  width: 100%;\\n\\n  /* Component styles */\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-sm);\\n  align-items: center;\\n  cursor: pointer;\\n  transition: background 100ms ease-in-out, border-color 100ms ease-in-out;\\n\\n  cem-pf-v6-label {\\n    flex-shrink: 0;\\n  }\\n\\n  \\u0026:hover {\\n    background: var(--pf-t--global--background--color--primary--hover);\\n  }\\n\\n  \\u0026:focus {\\n    outline: 2px solid var(--pf-t--global--border--color--clicked);\\n    outline-offset: -2px;\\n  }\\n\\n  \\u0026.selected {\\n    background: var(--pf-t--global--background--color--action--plain--selected);\\n    border-left-color: var(--pf-t--global--border--color--brand--default);\\n  }\\n}\\n\\n.event-time,\\n.event-element {\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.event-time {\\n  color: var(--cem-dev-server-text-muted);\\n  flex-shrink: 0;\\n  font-size: 11px;\\n}\\n\\n.event-element {\\n  color: var(--cem-dev-server-text-muted);\\n  font-weight: 400;\\n}\\n\\n/* Event detail panel */\\n.event-detail-header-content {\\n  padding: var(--cem-dev-server-spacing-md);\\n  border-bottom: var(--cem-dev-server-border-width) solid var(--cem-dev-server-border-color);\\n}\\n\\n.event-detail-name {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-lg);\\n  font-weight: 600;\\n  color: var(--cem-dev-server-text-primary);\\n}\\n\\n.event-detail-summary {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-secondary);\\n  line-height: 1.5;\\n  white-space: pre-wrap;\\n}\\n\\n.event-detail-description {\\n  margin: 0 0 var(--cem-dev-server-spacing-sm) 0;\\n  font-size: var(--cem-dev-server-font-size-sm);\\n  color: var(--cem-dev-server-text-secondary);\\n  line-height: 1.5;\\n  white-space: pre-wrap;\\n}\\n\\n.event-detail-meta {\\n  display: flex;\\n  gap: var(--cem-dev-server-spacing-md);\\n  font-size: var(--cem-dev-server-font-size-sm);\\n}\\n\\n.event-detail-time {\\n  color: var(--cem-dev-server-text-muted);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n}\\n\\n.event-detail-element {\\n  color: var(--cem-dev-server-text-secondary);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n}\\n\\n.event-detail-properties-heading {\\n  margin: var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  font-size: var(--cem-dev-server-font-size-base);\\n  font-weight: 600;\\n  color: var(--cem-dev-server-text-primary);\\n}\\n\\n.event-detail-properties {\\n  padding: var(--cem-dev-server-spacing-sm) var(--cem-dev-server-spacing-md);\\n  background: var(--cem-dev-server-bg-secondary);\\n  border: var(--cem-dev-server-border-width) solid var(--cem-dev-server-border-color);\\n  border-radius: var(--cem-dev-server-border-radius);\\n  font-family: var(--cem-dev-server-font-family-mono);\\n  font-size: 12px;\\n  line-height: 1.6;\\n  margin: 0 var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md) var(--cem-dev-server-spacing-md);\\n}\\n\\n.event-property-tree {\\n  list-style: none;\\n  padding: 0;\\n  margin: 0;\\n\\n  \\u0026.nested {\\n    padding-left: 1.5em;\\n    margin-top: 0.25em;\\n  }\\n}\\n\\n.property-item {\\n  padding: 0.125em 0;\\n}\\n\\n.property-key {\\n  color: var(--cem-dev-server-accent-color);\\n  font-weight: 500;\\n}\\n\\n.property-colon {\\n  color: var(--cem-dev-server-text-muted);\\n}\\n\\n.property-value {\\n  \\u0026.null,\\n  \\u0026.undefined {\\n    color: var(--cem-dev-server-text-muted);\\n    font-style: italic;\\n  }\\n\\n  \\u0026.boolean {\\n    color: var(--cem-dev-server-color-boolean);\\n  }\\n\\n  \\u0026.number {\\n    color: var(--cem-dev-server-color-number);\\n  }\\n\\n  \\u0026.string {\\n    color: var(--cem-dev-server-color-string);\\n  }\\n\\n  \\u0026.array,\\n  \\u0026.object {\\n    color: var(--cem-dev-server-text-secondary);\\n  }\\n}\\n\\n#debug-modal {\\n  container-type: inline-size;\\n}\\n"`));
var cem_serve_chrome_default = s53;

// lit-css:elements/cem-serve-demo/cem-serve-demo.css
var s54 = new CSSStyleSheet();
s54.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n\\niframe {\\n  border: none;\\n  width: 100%;\\n  height: 100%;\\n}\\n"'));
var cem_serve_demo_default = s54;

// elements/cem-serve-demo/cem-serve-demo.ts
var _rendering_dec, _a50, _CemServeDemo_decorators, _init50, _rendering, _iframeReady, _pendingMessages, _CemServeDemo_instances, iframe_get, iframeSrc_fn, _onChildReady, getElementInstance_fn, postKnobChange_fn, applyAttributeChange_fn, applyPropertyChange_fn, applyCSSStateChange_fn, applyCSSPropertyChange_fn;
_CemServeDemo_decorators = [t3("cem-serve-demo")];
var CemServeDemo = class extends (_a50 = i3, _rendering_dec = [n4({ reflect: true })], _a50) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeDemo_instances);
    __privateAdd(this, _rendering, __runInitializers(_init50, 8, this)), __runInitializers(_init50, 11, this);
    __privateAdd(this, _iframeReady, false);
    __privateAdd(this, _pendingMessages, []);
    __privateAdd(this, _onChildReady, (event) => {
      if (event.origin !== window.location.origin) return;
      if (event.data?.type !== "cem-iframe-ready") return;
      __privateSet(this, _iframeReady, true);
      const cw = __privateGet(this, _CemServeDemo_instances, iframe_get)?.contentWindow;
      if (cw) {
        for (const msg of __privateGet(this, _pendingMessages))
          cw.postMessage(msg, window.location.origin);
      }
      __privateSet(this, _pendingMessages, []);
      this.dispatchEvent(new Event("iframe-ready"));
    });
  }
  connectedCallback() {
    super.connectedCallback();
    window.addEventListener("message", __privateGet(this, _onChildReady));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener("message", __privateGet(this, _onChildReady));
  }
  render() {
    if (this.rendering === "iframe") {
      __privateSet(this, _iframeReady, false);
      return T`<iframe part="iframe"
                          src="${__privateMethod(this, _CemServeDemo_instances, iframeSrc_fn).call(this)}"></iframe>`;
    }
    return T`<slot></slot>`;
  }
  /**
   * Apply a knob change to an element in the demo.
   * Called by parent chrome element when knob events occur.
   * In iframe mode, bridges via postMessage instead of direct DOM access.
   */
  applyKnobChange(type, name, value, tagName, instanceIndex = 0) {
    if (this.rendering === "iframe") {
      return __privateMethod(this, _CemServeDemo_instances, postKnobChange_fn).call(this, type, name, value, tagName, instanceIndex);
    }
    const element = __privateMethod(this, _CemServeDemo_instances, getElementInstance_fn).call(this, tagName, instanceIndex);
    if (!element) {
      console.warn("[cem-serve-demo] Element not found:", tagName, "at index", instanceIndex);
      return false;
    }
    switch (type) {
      case "attribute":
        return __privateMethod(this, _CemServeDemo_instances, applyAttributeChange_fn).call(this, element, name, value);
      case "property":
        return __privateMethod(this, _CemServeDemo_instances, applyPropertyChange_fn).call(this, element, name, value);
      case "css-property":
        return __privateMethod(this, _CemServeDemo_instances, applyCSSPropertyChange_fn).call(this, element, name, value);
      case "css-state":
        return __privateMethod(this, _CemServeDemo_instances, applyCSSStateChange_fn).call(this, element, name, value);
      default:
        console.warn("[cem-serve-demo] Unknown knob type:", type);
        return false;
    }
  }
  /**
   * Set an attribute on an element in the demo
   */
  setDemoAttribute(selector, attribute, value) {
    const target = this.querySelector(selector);
    if (!target) return false;
    if (typeof value === "boolean") {
      target.toggleAttribute(attribute, value);
    } else if (value === "" || value === null || value === void 0) {
      target.removeAttribute(attribute);
    } else {
      target.setAttribute(attribute, value);
    }
    return true;
  }
  /**
   * Set a property on an element in the demo
   */
  setDemoProperty(selector, property, value) {
    const target = this.querySelector(selector);
    if (target) {
      target[property] = value;
      return true;
    }
    return false;
  }
  /**
   * Set a CSS custom property on an element in the demo
   */
  setDemoCssCustomProperty(selector, cssProperty, value) {
    const target = this.querySelector(selector);
    if (target) {
      const propertyName = cssProperty.startsWith("--") ? cssProperty : `--${cssProperty}`;
      target.style.setProperty(propertyName, value);
      return true;
    }
    return false;
  }
};
_init50 = __decoratorStart(_a50);
_rendering = new WeakMap();
_iframeReady = new WeakMap();
_pendingMessages = new WeakMap();
_CemServeDemo_instances = new WeakSet();
iframe_get = function() {
  return this.renderRoot.querySelector("iframe");
};
iframeSrc_fn = function() {
  const url = new URL(window.location.href);
  url.searchParams.set("rendering", "chromeless");
  return url.toString();
};
_onChildReady = new WeakMap();
/**
 * Find the Nth instance of an element by tag name
 */
getElementInstance_fn = function(tagName, instanceIndex = 0) {
  const elements = this.querySelectorAll(tagName);
  return elements[instanceIndex] || null;
};
postKnobChange_fn = function(knobType, name, value, tagName, instanceIndex) {
  const msg = { type: "cem-knob-change", knobType, name, value, tagName, instanceIndex };
  if (!__privateGet(this, _iframeReady)) {
    __privateGet(this, _pendingMessages).push(msg);
    return true;
  }
  const iframe = __privateGet(this, _CemServeDemo_instances, iframe_get);
  if (!iframe?.contentWindow) {
    console.warn("[cem-serve-demo] Iframe not ready for postMessage");
    return false;
  }
  iframe.contentWindow.postMessage(msg, window.location.origin);
  return true;
};
applyAttributeChange_fn = function(element, name, value) {
  if (typeof value === "boolean") {
    element.toggleAttribute(name, value);
  } else if (value === "" || value === null || value === void 0) {
    element.removeAttribute(name);
  } else {
    element.setAttribute(name, String(value));
  }
  return true;
};
applyPropertyChange_fn = function(element, name, value) {
  if (value === void 0) {
    try {
      delete element[name];
    } catch {
      element[name] = value;
    }
  } else {
    element[name] = value;
  }
  return true;
};
applyCSSStateChange_fn = function(element, name, value) {
  const stateName = name.startsWith("--") ? name.slice(2) : name;
  const internals = globalThis._elementInternals?.get(element);
  const states = internals?.states;
  if (!states) return false;
  if (value) {
    states.add(stateName);
  } else {
    states.delete(stateName);
  }
  return true;
};
applyCSSPropertyChange_fn = function(element, name, value) {
  const propertyName = name.startsWith("--") ? name : `--${name}`;
  if (value === "" || value === null || value === void 0) {
    element.style.removeProperty(propertyName);
  } else {
    element.style.setProperty(propertyName, String(value));
  }
  return true;
};
__decorateElement(_init50, 4, "rendering", _rendering_dec, CemServeDemo, _rendering);
CemServeDemo = __decorateElement(_init50, 0, "CemServeDemo", _CemServeDemo_decorators, CemServeDemo);
__publicField(CemServeDemo, "styles", cem_serve_demo_default);
__runInitializers(_init50, 1, CemServeDemo);

// lit-css:elements/cem-serve-knob-group/cem-serve-knob-group.css
var s55 = new CSSStyleSheet();
s55.replaceSync(JSON.parse('":host {\\n  display: block;\\n}\\n"'));
var cem_serve_knob_group_default = s55;

// elements/cem-serve-knob-group/cem-serve-knob-group.ts
var KnobAttributeChangeEvent = class extends Event {
  name;
  value;
  constructor(name, value) {
    super("knob:attribute-change", { bubbles: true });
    this.name = name;
    this.value = value;
  }
};
var KnobPropertyChangeEvent = class extends Event {
  name;
  value;
  constructor(name, value) {
    super("knob:property-change", { bubbles: true });
    this.name = name;
    this.value = value;
  }
};
var KnobCssPropertyChangeEvent = class extends Event {
  name;
  value;
  constructor(name, value) {
    super("knob:css-property-change", { bubbles: true });
    this.name = name;
    this.value = value;
  }
};
var KnobAttributeClearEvent = class extends Event {
  name;
  constructor(name) {
    super("knob:attribute-clear", { bubbles: true });
    this.name = name;
  }
};
var KnobPropertyClearEvent = class extends Event {
  name;
  constructor(name) {
    super("knob:property-clear", { bubbles: true });
    this.name = name;
  }
};
var KnobCssPropertyClearEvent = class extends Event {
  name;
  constructor(name) {
    super("knob:css-property-clear", { bubbles: true });
    this.name = name;
  }
};
var KnobCssStateChangeEvent = class extends Event {
  name;
  value;
  constructor(name, value) {
    super("knob:css-state-change", { bubbles: true });
    this.name = name;
    this.value = value;
  }
};
var KnobCssStateClearEvent = class extends Event {
  name;
  constructor(name) {
    super("knob:css-state-clear", { bubbles: true });
    this.name = name;
  }
};
var _htmlFor_dec, _a51, _CemServeKnobGroup_decorators, _init51, _htmlFor, _debounceTimers, _debounceDelay, _colorButtonListeners, _clearButtonListeners, _CemServeKnobGroup_instances, onSlotChange_fn4, attachColorButtonListeners_fn, removeColorButtonListeners_fn, attachClearButtonListeners_fn, removeClearButtonListeners_fn, _handleClearButtonClick, updateClearButtonVisibility_fn, _handleInput, _handleChange, isBooleanControl_fn, _handleColorButtonClick, applyChange_fn, parseValue_fn;
_CemServeKnobGroup_decorators = [t3("cem-serve-knob-group")];
var CemServeKnobGroup = class extends (_a51 = i3, _htmlFor_dec = [n4({ reflect: true, attribute: "for" })], _a51) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeKnobGroup_instances);
    __privateAdd(this, _htmlFor, __runInitializers(_init51, 8, this)), __runInitializers(_init51, 11, this);
    __privateAdd(this, _debounceTimers, /* @__PURE__ */ new Map());
    __privateAdd(this, _debounceDelay, 250);
    __privateAdd(this, _colorButtonListeners, /* @__PURE__ */ new WeakMap());
    __privateAdd(this, _clearButtonListeners, /* @__PURE__ */ new WeakMap());
    __privateAdd(this, _handleClearButtonClick, (e6, button) => {
      e6.preventDefault();
      const knobType = button.dataset.knobType;
      const knobName = button.dataset.knobName;
      if (!knobType || !knobName) return;
      const formGroup = button.closest("cem-pf-v6-form-group");
      if (!formGroup) return;
      const control = formGroup.querySelector(
        `[data-knob-type="${knobType}"][data-knob-name="${knobName}"]`
      );
      if (!control) return;
      if (__privateMethod(this, _CemServeKnobGroup_instances, isBooleanControl_fn).call(this, control)) {
        control.checked = false;
      } else {
        control.value = "";
      }
      button.hidden = true;
      control.dispatchEvent(new Event("input", { bubbles: true }));
      switch (knobType) {
        case "attribute":
          this.dispatchEvent(new KnobAttributeClearEvent(knobName));
          break;
        case "property":
          this.dispatchEvent(new KnobPropertyClearEvent(knobName));
          break;
        case "css-property":
          this.dispatchEvent(new KnobCssPropertyClearEvent(knobName));
          break;
        case "css-state":
          this.dispatchEvent(new KnobCssStateClearEvent(knobName));
          break;
      }
    });
    __privateAdd(this, _handleInput, (e6) => {
      const control = e6.target;
      const knobType = control.dataset?.knobType;
      const knobName = control.dataset?.knobName;
      if (!knobType || !knobName) return;
      __privateMethod(this, _CemServeKnobGroup_instances, updateClearButtonVisibility_fn).call(this, control);
      if (__privateMethod(this, _CemServeKnobGroup_instances, isBooleanControl_fn).call(this, control)) {
        return __privateMethod(this, _CemServeKnobGroup_instances, applyChange_fn).call(this, knobType, knobName, control.checked);
      }
      if (control.tagName === "SELECT") {
        return;
      }
      const key = `${knobType}-${knobName}`;
      clearTimeout(__privateGet(this, _debounceTimers).get(key));
      __privateGet(this, _debounceTimers).set(key, setTimeout(() => {
        __privateMethod(this, _CemServeKnobGroup_instances, applyChange_fn).call(this, knobType, knobName, control.value);
      }, __privateGet(this, _debounceDelay)));
    });
    __privateAdd(this, _handleChange, (e6) => {
      const control = e6.target;
      const knobType = control.dataset?.knobType;
      const knobName = control.dataset?.knobName;
      if (!knobType || !knobName) return;
      __privateMethod(this, _CemServeKnobGroup_instances, updateClearButtonVisibility_fn).call(this, control);
      const key = `${knobType}-${knobName}`;
      clearTimeout(__privateGet(this, _debounceTimers).get(key));
      const value = __privateMethod(this, _CemServeKnobGroup_instances, isBooleanControl_fn).call(this, control) ? control.checked : control.value;
      __privateMethod(this, _CemServeKnobGroup_instances, applyChange_fn).call(this, knobType, knobName, value);
    });
    __privateAdd(this, _handleColorButtonClick, async (e6, button) => {
      e6.preventDefault();
      const textInputGroup = button.closest("cem-pf-v6-text-input-group");
      if (!textInputGroup) return;
      const currentValue = textInputGroup.value || "#000000";
      if ("EyeDropper" in window) {
        try {
          const eyeDropper = new window.EyeDropper();
          const result = await eyeDropper.open();
          textInputGroup.value = result.sRGBHex;
          textInputGroup.dispatchEvent(new Event("input", { bubbles: true }));
        } catch (err) {
          if (err.name !== "AbortError") {
            console.warn("[KnobGroup] EyeDropper error:", err);
          }
        }
      } else {
        const colorInput = document.createElement("input");
        colorInput.type = "color";
        colorInput.value = currentValue;
        colorInput.style.position = "absolute";
        colorInput.style.opacity = "0";
        colorInput.style.pointerEvents = "none";
        document.body.appendChild(colorInput);
        colorInput.addEventListener("input", () => {
          textInputGroup.value = colorInput.value;
          textInputGroup.dispatchEvent(new Event("input", { bubbles: true }));
        });
        colorInput.addEventListener("change", () => {
          textInputGroup.value = colorInput.value;
          textInputGroup.dispatchEvent(new Event("input", { bubbles: true }));
          if (colorInput.parentNode) {
            document.body.removeChild(colorInput);
          }
        });
        colorInput.addEventListener("blur", () => {
          if (colorInput.parentNode) {
            document.body.removeChild(colorInput);
          }
        });
        colorInput.click();
      }
    });
  }
  render() {
    return T`<slot @slotchange=${__privateMethod(this, _CemServeKnobGroup_instances, onSlotChange_fn4)}></slot>`;
  }
  connectedCallback() {
    super.connectedCallback();
    this.addEventListener("input", __privateGet(this, _handleInput));
    this.addEventListener("change", __privateGet(this, _handleChange));
  }
  firstUpdated() {
    __privateMethod(this, _CemServeKnobGroup_instances, attachColorButtonListeners_fn).call(this);
    __privateMethod(this, _CemServeKnobGroup_instances, attachClearButtonListeners_fn).call(this);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    for (const timer of __privateGet(this, _debounceTimers).values()) {
      clearTimeout(timer);
    }
    __privateGet(this, _debounceTimers).clear();
    this.removeEventListener("input", __privateGet(this, _handleInput));
    this.removeEventListener("change", __privateGet(this, _handleChange));
    __privateMethod(this, _CemServeKnobGroup_instances, removeColorButtonListeners_fn).call(this);
    __privateMethod(this, _CemServeKnobGroup_instances, removeClearButtonListeners_fn).call(this);
  }
};
_init51 = __decoratorStart(_a51);
_htmlFor = new WeakMap();
_debounceTimers = new WeakMap();
_debounceDelay = new WeakMap();
_colorButtonListeners = new WeakMap();
_clearButtonListeners = new WeakMap();
_CemServeKnobGroup_instances = new WeakSet();
onSlotChange_fn4 = function() {
  __privateMethod(this, _CemServeKnobGroup_instances, attachColorButtonListeners_fn).call(this);
  __privateMethod(this, _CemServeKnobGroup_instances, attachClearButtonListeners_fn).call(this);
};
attachColorButtonListeners_fn = function() {
  const buttons = this.querySelectorAll(".color-picker-button");
  for (const button of buttons) {
    if (__privateGet(this, _colorButtonListeners).has(button)) continue;
    const handler = (e6) => __privateGet(this, _handleColorButtonClick).call(this, e6, button);
    __privateGet(this, _colorButtonListeners).set(button, handler);
    button.addEventListener("click", handler);
  }
};
removeColorButtonListeners_fn = function() {
  const buttons = this.querySelectorAll(".color-picker-button");
  for (const button of buttons) {
    const handler = __privateGet(this, _colorButtonListeners).get(button);
    if (handler) {
      button.removeEventListener("click", handler);
      __privateGet(this, _colorButtonListeners).delete(button);
    }
  }
};
attachClearButtonListeners_fn = function() {
  const buttons = this.querySelectorAll(".knob-clear-button");
  for (const button of buttons) {
    if (__privateGet(this, _clearButtonListeners).has(button)) continue;
    const handler = (e6) => __privateGet(this, _handleClearButtonClick).call(this, e6, button);
    __privateGet(this, _clearButtonListeners).set(button, handler);
    button.addEventListener("click", handler);
  }
};
removeClearButtonListeners_fn = function() {
  const buttons = this.querySelectorAll(".knob-clear-button");
  for (const button of buttons) {
    const handler = __privateGet(this, _clearButtonListeners).get(button);
    if (handler) {
      button.removeEventListener("click", handler);
      __privateGet(this, _clearButtonListeners).delete(button);
    }
  }
};
_handleClearButtonClick = new WeakMap();
updateClearButtonVisibility_fn = function(control) {
  const knobType = control.dataset.knobType;
  const knobName = control.dataset.knobName;
  if (!knobType || !knobName) return;
  const formGroup = control.closest("cem-pf-v6-form-group");
  if (!formGroup) return;
  const clearButton = formGroup.querySelector(
    `.knob-clear-button[data-knob-type="${knobType}"][data-knob-name="${knobName}"]`
  );
  if (!clearButton) return;
  const hasValue = __privateMethod(this, _CemServeKnobGroup_instances, isBooleanControl_fn).call(this, control) ? control.checked : control.value !== "";
  clearButton.hidden = !hasValue;
};
_handleInput = new WeakMap();
_handleChange = new WeakMap();
isBooleanControl_fn = function(control) {
  if (control.tagName === "CEM-PF-V6-SWITCH") {
    return true;
  }
  if (control.tagName === "INPUT" && control.type === "checkbox") {
    return true;
  }
  return false;
};
_handleColorButtonClick = new WeakMap();
applyChange_fn = function(type, name, value) {
  switch (type) {
    case "attribute":
      this.dispatchEvent(new KnobAttributeChangeEvent(name, value));
      break;
    case "property":
      this.dispatchEvent(new KnobPropertyChangeEvent(name, __privateMethod(this, _CemServeKnobGroup_instances, parseValue_fn).call(this, value)));
      break;
    case "css-property":
      this.dispatchEvent(new KnobCssPropertyChangeEvent(name, value));
      break;
    case "css-state":
      this.dispatchEvent(new KnobCssStateChangeEvent(name, value));
      break;
    default:
      console.warn(`[KnobGroup] Unknown knob type: ${type}`);
      return;
  }
};
parseValue_fn = function(value) {
  if (typeof value === "boolean") {
    return value;
  }
  if (value === "true") return true;
  if (value === "false") return false;
  if (value === "null") return null;
  if (value === "") return "";
  const num = Number(value);
  if (!isNaN(num) && value !== "") {
    return num;
  }
  return value;
};
__decorateElement(_init51, 4, "htmlFor", _htmlFor_dec, CemServeKnobGroup, _htmlFor);
CemServeKnobGroup = __decorateElement(_init51, 0, "CemServeKnobGroup", _CemServeKnobGroup_decorators, CemServeKnobGroup);
__publicField(CemServeKnobGroup, "styles", cem_serve_knob_group_default);
__runInitializers(_init51, 1, CemServeKnobGroup);

// lit-css:elements/cem-serve-knobs/cem-serve-knobs.css
var s56 = new CSSStyleSheet();
s56.replaceSync(JSON.parse('":host {\\n  display: flex;\\n  flex-direction: column;\\n  gap: 0;\\n  height: 100%;\\n}\\n\\ncem-pf-v6-navigation {\\n  flex-shrink: 0;\\n}\\n\\n#knobs {\\n  flex: 1;\\n  min-height: 0;\\n  overflow-y: auto;\\n  padding: var(--pf-t--global--spacer--lg);\\n}\\n\\n#knobs slot {\\n  display: grid;\\n  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));\\n  gap: var(--pf-t--global--spacer--gap--group-to-group--horizontal--default) var(--pf-t--global--spacer--gap--group-to-group--vertical--default);\\n  align-items: start;\\n}\\n\\n::slotted(cem-pf-v6-card.active) {\\n  --cem-pf-v6-c-card--BorderColor: var(--cem-pf-v6-c-card--m-selectable--m-selected--BorderColor);\\n  --cem-pf-v6-c-card--BorderWidth: var(--cem-pf-v6-c-card--m-selectable--m-selected--BorderWidth);\\n}\\n"'));
var cem_serve_knobs_default = s56;

// elements/cem-serve-knobs/cem-serve-knobs.ts
var _CemServeKnobs_decorators, _init52, _a52;
_CemServeKnobs_decorators = [t3("cem-serve-knobs")];
var CemServeKnobs = class extends (_a52 = i3) {
  static styles = cem_serve_knobs_default;
  #navList = null;
  #handleHashChange = () => {
    const hash = window.location.hash;
    if (!hash || hash === "#") return;
    const cardId = hash.substring(1);
    const cards = this.querySelectorAll("cem-pf-v6-card");
    const navLinks = this.#navList?.querySelectorAll("cem-pf-v6-nav-link");
    const knobsContainer = this.shadowRoot?.getElementById("knobs");
    if (navLinks) {
      for (const link of navLinks) {
        const linkHref = link.getAttribute("href");
        link.toggleAttribute("current", linkHref === hash);
      }
    }
    for (const card of cards) {
      const isActive = card.dataset.card === cardId;
      card.classList.toggle("active", isActive);
      if (isActive && knobsContainer) {
        const cardRect = card.getBoundingClientRect();
        const containerRect = knobsContainer.getBoundingClientRect();
        const scrollOffset = cardRect.top - containerRect.top + knobsContainer.scrollTop;
        knobsContainer.scrollTo({
          top: scrollOffset,
          behavior: "smooth"
        });
      }
    }
  };
  render() {
    return T`
      <cem-pf-v6-navigation horizontal
                        variant="horizontal-subnav">
        <cem-pf-v6-nav-list id="nav-list"
                        aria-label="Elements">
        </cem-pf-v6-nav-list>
      </cem-pf-v6-navigation>
      <div id="knobs">
        <slot @slotchange=${this.#onSlotChange}></slot>
      </div>
    `;
  }
  connectedCallback() {
    super.connectedCallback();
    window.addEventListener("hashchange", this.#handleHashChange);
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener("hashchange", this.#handleHashChange);
  }
  firstUpdated() {
    this.#navList = this.shadowRoot?.getElementById("nav-list") ?? null;
    this.#updateNavigation();
    this.#handleHashChange();
  }
  #onSlotChange() {
    this.#updateNavigation();
  }
  #updateNavigation() {
    const slot = this.shadowRoot?.querySelector("slot");
    if (!slot || !this.#navList) return;
    const panels = slot.assignedElements().filter((el) => el.tagName === "CEM-PF-V6-CARD");
    if (panels.length === 0) return;
    this.#navList.innerHTML = "";
    panels.forEach((panel, index) => {
      const navItem = document.createElement("cem-pf-v6-nav-item");
      const navLink = document.createElement("cem-pf-v6-nav-link");
      const label = document.createElement("span");
      const panelId = panel.dataset.card || `instance-${index}`;
      const labelText = panel.dataset.label || `Instance ${index + 1}`;
      label.className = "instance-label";
      label.textContent = labelText;
      navLink.setAttribute("href", `#${panelId}`);
      if (index === 0) {
        navLink.setAttribute("current", "");
      }
      navLink.appendChild(label);
      navItem.appendChild(navLink);
      this.#navList.appendChild(navItem);
    });
  }
};
_init52 = __decoratorStart(_a52);
CemServeKnobs = __decorateElement(_init52, 0, "CemServeKnobs", _CemServeKnobs_decorators, CemServeKnobs);
__runInitializers(_init52, 1, CemServeKnobs);

// lit-css:elements/cem-transform-error-overlay/cem-transform-error-overlay.css
var s57 = new CSSStyleSheet();
s57.replaceSync(JSON.parse('"/* Transform Error Overlay - displays server-side compilation errors */\\n\\n:host {\\n  display: none;\\n  position: fixed;\\n  inset: 0;\\n  z-index: var(--pf-t--global--z-index--2xl, 1000);\\n  background: rgba(0, 0, 0, 0.9);\\n  backdrop-filter: blur(4px);\\n  animation: fadeIn 0.2s ease-out;\\n}\\n\\n:host([open]) {\\n  display: flex;\\n  align-items: center;\\n  justify-content: center;\\n  padding: var(--pf-t--global--spacer--lg);\\n}\\n\\n@keyframes fadeIn {\\n  from {\\n    opacity: 0;\\n  }\\n  to {\\n    opacity: 1;\\n  }\\n}\\n\\n#overlay-content {\\n  background: var(--pf-t--global--background--color--floating--default);\\n  color: var(--pf-t--global--text--color--regular);\\n  border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--color--status--danger--default);\\n  border-radius: var(--pf-t--global--border--radius--medium);\\n  max-width: 800px;\\n  width: 100%;\\n  max-height: 90vh;\\n  display: flex;\\n  flex-direction: column;\\n  box-shadow: var(--pf-t--global--box-shadow--xl);\\n  font-family: var(--pf-t--global--font--family--mono);\\n}\\n\\n#header {\\n  display: flex;\\n  align-items: center;\\n  justify-content: space-between;\\n  padding: var(--pf-t--global--spacer--md) var(--pf-t--global--spacer--lg);\\n  background: var(--pf-t--global--color--status--danger--default);\\n  color: var(--pf-t--global--text--color--on-status--danger);\\n  border-radius: var(--pf-t--global--border--radius--small) var(--pf-t--global--border--radius--small) 0 0;\\n}\\n\\n#title-container {\\n  display: flex;\\n  align-items: center;\\n  gap: var(--pf-t--global--spacer--sm);\\n  font-size: var(--pf-t--global--font--size--body--lg);\\n  font-weight: var(--pf-t--global--font--weight--body--bold);\\n  margin: 0;\\n}\\n\\n#error-icon {\\n  font-size: var(--pf-t--global--font--size--heading--sm);\\n}\\n\\n#close {\\n  --cem-pf-v6-c-button--Color: var(--pf-t--global--text--color--on-status--danger);\\n  --cem-pf-v6-c-button--BackgroundColor: rgba(255, 255, 255, 0.2);\\n  --cem-pf-v6-c-button--BorderColor: transparent;\\n  --cem-pf-v6-c-button--hover--BackgroundColor: rgba(255, 255, 255, 0.3);\\n  --cem-pf-v6-c-button--hover--Color: var(--pf-t--global--text--color--on-status--danger);\\n}\\n\\n#body {\\n  padding: var(--pf-t--global--spacer--lg);\\n  overflow-y: auto;\\n  flex: 1;\\n}\\n\\n#file {\\n  background: var(--pf-t--global--background--color--secondary--default);\\n  padding: var(--pf-t--global--spacer--sm) var(--pf-t--global--spacer--md);\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  margin-bottom: var(--pf-t--global--spacer--md);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  color: var(--pf-t--global--color--brand--default);\\n  border-inline-start: var(--pf-t--global--border--width--strong, 3px) solid var(--pf-t--global--color--status--danger--default);\\n}\\n\\n#file:empty {\\n  display: none;\\n}\\n\\n#message {\\n  background: var(--pf-t--global--background--color--primary--default);\\n  padding: var(--pf-t--global--spacer--md);\\n  border-radius: var(--pf-t--global--border--radius--small);\\n  white-space: pre-wrap;\\n  word-break: break-word;\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  line-height: var(--pf-t--global--font--line-height--body);\\n  color: var(--pf-t--global--text--color--regular);\\n  border: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n}\\n\\n#footer {\\n  padding: var(--pf-t--global--spacer--md) var(--pf-t--global--spacer--lg);\\n  background: var(--pf-t--global--background--color--secondary--default);\\n  border-block-start: var(--pf-t--global--border--width--regular) solid var(--pf-t--global--border--color--default);\\n  border-radius: 0 0 var(--pf-t--global--border--radius--small) var(--pf-t--global--border--radius--small);\\n  font-size: var(--pf-t--global--font--size--body--sm);\\n  color: var(--pf-t--global--text--color--subtle);\\n}\\n"'));
var cem_transform_error_overlay_default = s57;

// elements/cem-transform-error-overlay/cem-transform-error-overlay.ts
var _message_dec2, _file_dec, _title_dec2, _open_dec4, _a53, _CemTransformErrorOverlay_decorators, _init53, _open4, _title2, _file, _message2, _handleKeydown4;
_CemTransformErrorOverlay_decorators = [t3("cem-transform-error-overlay")];
var CemTransformErrorOverlay = class extends (_a53 = i3, _open_dec4 = [n4({ type: Boolean, reflect: true })], _title_dec2 = [n4()], _file_dec = [n4()], _message_dec2 = [n4()], _a53) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _open4, __runInitializers(_init53, 8, this, false)), __runInitializers(_init53, 11, this);
    __privateAdd(this, _title2, __runInitializers(_init53, 12, this, "")), __runInitializers(_init53, 15, this);
    __privateAdd(this, _file, __runInitializers(_init53, 16, this, "")), __runInitializers(_init53, 19, this);
    __privateAdd(this, _message2, __runInitializers(_init53, 20, this, "")), __runInitializers(_init53, 23, this);
    __privateAdd(this, _handleKeydown4, (e6) => {
      if (e6.key === "Escape" && this.open) {
        this.hide();
      }
    });
  }
  connectedCallback() {
    super.connectedCallback();
    document.addEventListener("keydown", __privateGet(this, _handleKeydown4));
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    document.removeEventListener("keydown", __privateGet(this, _handleKeydown4));
  }
  render() {
    return T`
      <div id="overlay-content">
        <div id="header">
          <h2 id="title-container">
            <span id="error-icon">\u26A0\uFE0F</span>
            <span>${this.title}</span>
          </h2>
          <cem-pf-v6-button id="close"
                        variant="plain"
                        @click=${this.hide}>Dismiss</cem-pf-v6-button>
        </div>
        <div id="body">
          <div id="file">${this.file ? `File: ${this.file}` : ""}</div>
          <div id="message">${this.message}</div>
        </div>
        <div id="footer">
          This error will automatically dismiss when the issue is fixed.
        </div>
      </div>
    `;
  }
  /**
   * Show the error overlay.
   * @param title - Error title
   * @param message - Error message
   * @param file - Optional file path where error occurred
   */
  show(title, message, file = "") {
    this.title = title;
    this.message = message;
    this.file = file;
    this.open = true;
  }
  /** Hide the error overlay */
  hide() {
    this.open = false;
  }
};
_init53 = __decoratorStart(_a53);
_open4 = new WeakMap();
_title2 = new WeakMap();
_file = new WeakMap();
_message2 = new WeakMap();
_handleKeydown4 = new WeakMap();
__decorateElement(_init53, 4, "open", _open_dec4, CemTransformErrorOverlay, _open4);
__decorateElement(_init53, 4, "title", _title_dec2, CemTransformErrorOverlay, _title2);
__decorateElement(_init53, 4, "file", _file_dec, CemTransformErrorOverlay, _file);
__decorateElement(_init53, 4, "message", _message_dec2, CemTransformErrorOverlay, _message2);
CemTransformErrorOverlay = __decorateElement(_init53, 0, "CemTransformErrorOverlay", _CemTransformErrorOverlay_decorators, CemTransformErrorOverlay);
__publicField(CemTransformErrorOverlay, "styles", cem_transform_error_overlay_default);
__runInitializers(_init53, 1, CemTransformErrorOverlay);

// elements/cem-serve-chrome/cem-serve-chrome.ts
var CEMReloadClient4;
var StatePersistence4;
var CemLogsEvent = class extends Event {
  logs;
  constructor(logs) {
    super("cem:logs", { bubbles: true });
    this.logs = logs;
  }
};
var _hasDescription_dec, _sidebar_dec, _tabsSelected_dec, _drawerHeight_dec2, _drawer_dec, _knobs_dec, _sourceURL_dec, _canonicalURL_dec, _packageName_dec, _demoTitle_dec, _primaryTagName_dec, _a54, _CemServeChrome_decorators, _demoInfoTemplate, _demoGroupTemplate, _demoListTemplate, _logEntryTemplate, _eventEntryTemplate, _init54, _primaryTagName, _demoTitle, _packageName, _canonicalURL, _sourceURL, _knobs, _drawer, _drawerHeight2, _tabsSelected, _sidebar, _hasDescription2, _CemServeChrome_instances, $_fn2, $$_fn, _logContainer, _drawerOpen, _initialLogsFetched, _isInitialLoad, _debugData, _elementEventMap, _manifest, _capturedEvents, _maxCapturedEvents, _eventList, _eventDetailHeader, _eventDetailBody, _selectedEventId, _eventsFilterValue, _eventsFilterDebounceTimer, _eventTypeFilters, _elementFilters, _discoveredElements, _handleLogsEvent, _handleTreeExpand, _handleTreeCollapse, _handleTreeSelect, _copyLogsFeedbackTimeout, _copyDebugFeedbackTimeout, _copyEventsFeedbackTimeout, _logsFilterValue, _logsFilterDebounceTimer, _logLevelFilters, _logLevelDropdown, _observer2, _wsClient, initWsClient_fn, renderSourceButton_fn, _modulesReady, loadClientModules_fn, fetchDebugInfo_fn, populateDemoUrls_fn, setupLogListener_fn, filterLogs_fn, getLogTypeFromEntry_fn, loadLogFilterState_fn, syncCheckboxStates_fn, saveLogFilterState_fn, _handleLogFilterChange, copyLogs_fn, setupDebugOverlay_fn, setupFooterDrawer_fn, detectBrowser_fn, copyDebugInfo_fn, renderLogs_fn, getLogBadge_fn, applyLogLabelAttrs_fn, scrollLatestIntoView_fn, scrollLogsToBottom_fn, migrateFromLocalStorageIfNeeded_fn, setupColorSchemeToggle_fn, applyColorScheme_fn, setupKnobCoordination_fn, _onKnobChange, _onKnobClear, disableCssStateKnob_fn, getKnobTarget_fn, getKnobTypeFromEvent_fn, getKnobTypeFromClearEvent_fn, setupTreeStatePersistence_fn, applyTreeState_fn, setupSidebarStatePersistence_fn, findTreeItemById_fn, getTreeNodeId_fn, discoverElementEvents_fn, setupEventCapture_fn, attachEventListeners_fn, observeDemoMutations_fn, _handleElementEvent, getEventDocumentation_fn, findTypeDeclaration_fn, captureEvent_fn, extractEventProperties_fn, renderEvent_fn, selectEvent_fn, buildPropertiesForDisplay_fn, buildPropertyTree_fn, scrollEventsToBottom_fn, isEventsTabActive_fn, filterEvents_fn, eventMatchesTextFilter_fn, getEventRecordById_fn, updateEventFilters_fn, _handleEventTypeFilterChange, _handleElementFilterChange, loadEventFiltersFromStorage_fn, saveEventFilters_fn, clearEvents_fn, copyEvents_fn, setupEventListeners_fn;
_CemServeChrome_decorators = [t3("cem-serve-chrome")];
var _CemServeChrome = class _CemServeChrome extends (_a54 = i3, _primaryTagName_dec = [n4({ attribute: "primary-tag-name" })], _demoTitle_dec = [n4({ attribute: "demo-title" })], _packageName_dec = [n4({ attribute: "package-name" })], _canonicalURL_dec = [n4({ attribute: "canonical-url" })], _sourceURL_dec = [n4({ attribute: "source-url" })], _knobs_dec = [n4()], _drawer_dec = [n4()], _drawerHeight_dec2 = [n4({ attribute: "drawer-height" })], _tabsSelected_dec = [n4({ attribute: "tabs-selected" })], _sidebar_dec = [n4()], _hasDescription_dec = [n4({ type: Boolean, attribute: "has-description" })], _a54) {
  constructor() {
    super(...arguments);
    __privateAdd(this, _CemServeChrome_instances);
    __privateAdd(this, _primaryTagName, __runInitializers(_init54, 8, this, "")), __runInitializers(_init54, 11, this);
    __privateAdd(this, _demoTitle, __runInitializers(_init54, 12, this, "")), __runInitializers(_init54, 15, this);
    __privateAdd(this, _packageName, __runInitializers(_init54, 16, this, "")), __runInitializers(_init54, 19, this);
    __privateAdd(this, _canonicalURL, __runInitializers(_init54, 20, this, "")), __runInitializers(_init54, 23, this);
    __privateAdd(this, _sourceURL, __runInitializers(_init54, 24, this, "")), __runInitializers(_init54, 27, this);
    __privateAdd(this, _knobs, __runInitializers(_init54, 28, this, "")), __runInitializers(_init54, 31, this);
    __privateAdd(this, _drawer, __runInitializers(_init54, 32, this, "")), __runInitializers(_init54, 35, this);
    __privateAdd(this, _drawerHeight2, __runInitializers(_init54, 36, this, "")), __runInitializers(_init54, 39, this);
    __privateAdd(this, _tabsSelected, __runInitializers(_init54, 40, this, "")), __runInitializers(_init54, 43, this);
    __privateAdd(this, _sidebar, __runInitializers(_init54, 44, this, "")), __runInitializers(_init54, 47, this);
    __privateAdd(this, _hasDescription2, __runInitializers(_init54, 48, this, false)), __runInitializers(_init54, 51, this);
    __privateAdd(this, _logContainer, null);
    __privateAdd(this, _drawerOpen, false);
    __privateAdd(this, _initialLogsFetched, false);
    __privateAdd(this, _isInitialLoad, true);
    __privateAdd(this, _debugData, null);
    // Element event tracking
    __privateAdd(this, _elementEventMap, null);
    __privateAdd(this, _manifest, null);
    __privateAdd(this, _capturedEvents, []);
    __privateAdd(this, _maxCapturedEvents, 1e3);
    __privateAdd(this, _eventList, null);
    __privateAdd(this, _eventDetailHeader, null);
    __privateAdd(this, _eventDetailBody, null);
    __privateAdd(this, _selectedEventId, null);
    __privateAdd(this, _eventsFilterValue, "");
    __privateAdd(this, _eventsFilterDebounceTimer, null);
    __privateAdd(this, _eventTypeFilters, /* @__PURE__ */ new Set());
    __privateAdd(this, _elementFilters, /* @__PURE__ */ new Set());
    __privateAdd(this, _discoveredElements, /* @__PURE__ */ new Set());
    // Event listener references for cleanup
    __privateAdd(this, _handleLogsEvent, null);
    __privateAdd(this, _handleTreeExpand, null);
    __privateAdd(this, _handleTreeCollapse, null);
    __privateAdd(this, _handleTreeSelect, null);
    // Timeout IDs for cleanup
    __privateAdd(this, _copyLogsFeedbackTimeout, null);
    __privateAdd(this, _copyDebugFeedbackTimeout, null);
    __privateAdd(this, _copyEventsFeedbackTimeout, null);
    // Log filter state
    __privateAdd(this, _logsFilterValue, "");
    __privateAdd(this, _logsFilterDebounceTimer, null);
    __privateAdd(this, _logLevelFilters, /* @__PURE__ */ new Set(["info", "warn", "error", "success", "debug", "trace"]));
    __privateAdd(this, _logLevelDropdown, null);
    // Watch for dynamically added elements
    /* c8 ignore start - MutationObserver callback tested via integration */
    __privateAdd(this, _observer2, new MutationObserver((mutations) => {
      let needsUpdate = false;
      for (const mutation of mutations) {
        for (const node of mutation.addedNodes) {
          if (node instanceof HTMLElement) {
            const tagName = node.tagName.toLowerCase();
            if (__privateGet(this, _elementEventMap)?.has(tagName) && !node.dataset.cemEventsAttached) {
              const eventInfo = __privateGet(this, _elementEventMap).get(tagName);
              for (const eventName of eventInfo.eventNames) {
                node.addEventListener(eventName, __privateGet(this, _handleElementEvent), { capture: true });
              }
              node.dataset.cemEventsAttached = "true";
              __privateGet(this, _discoveredElements).add(tagName);
              needsUpdate = true;
            }
          }
        }
      }
      if (needsUpdate) {
        __privateMethod(this, _CemServeChrome_instances, updateEventFilters_fn).call(this);
      }
    }));
    /* c8 ignore stop */
    __privateAdd(this, _wsClient);
    /** Resolves when client-only modules are loaded. */
    __privateAdd(this, _modulesReady, __privateMethod(this, _CemServeChrome_instances, loadClientModules_fn).call(this));
    __privateAdd(this, _handleLogFilterChange, (event) => {
      const { value, checked } = event;
      if (checked) {
        __privateGet(this, _logLevelFilters).add(value);
      } else {
        __privateGet(this, _logLevelFilters).delete(value);
      }
      __privateMethod(this, _CemServeChrome_instances, saveLogFilterState_fn).call(this);
      __privateMethod(this, _CemServeChrome_instances, filterLogs_fn).call(this, __privateGet(this, _logsFilterValue));
    });
    __privateAdd(this, _onKnobChange, (event) => {
      const target = __privateMethod(this, _CemServeChrome_instances, getKnobTarget_fn).call(this, event);
      if (!target) {
        console.warn("[cem-serve-chrome] Could not find knob target info in event path");
        return;
      }
      const { tagName, instanceIndex } = target;
      const demo = this.demo;
      if (!demo) return;
      const knobType = __privateMethod(this, _CemServeChrome_instances, getKnobTypeFromEvent_fn).call(this, event);
      const success = demo.applyKnobChange(
        knobType,
        event.name,
        event.value,
        tagName,
        instanceIndex
      );
      if (!success) {
        if (knobType === "css-state") {
          const elementExists = !!demo.querySelectorAll(tagName)[instanceIndex];
          if (elementExists) {
            __privateMethod(this, _CemServeChrome_instances, disableCssStateKnob_fn).call(this, event);
          }
        }
        console.warn("[cem-serve-chrome] Failed to apply knob change:", {
          type: knobType,
          name: event.name,
          tagName,
          instanceIndex
        });
      }
    });
    __privateAdd(this, _onKnobClear, (event) => {
      const target = __privateMethod(this, _CemServeChrome_instances, getKnobTarget_fn).call(this, event);
      if (!target) {
        console.warn("[cem-serve-chrome] Could not find knob target info in event path");
        return;
      }
      const { tagName, instanceIndex } = target;
      const demo = this.demo;
      if (!demo) return;
      const knobType = __privateMethod(this, _CemServeChrome_instances, getKnobTypeFromClearEvent_fn).call(this, event);
      const clearValue = knobType === "property" ? void 0 : knobType === "css-state" ? false : "";
      const success = demo.applyKnobChange(
        knobType,
        event.name,
        clearValue,
        tagName,
        instanceIndex
      );
      if (!success) {
        console.warn("[cem-serve-chrome] Failed to clear knob:", {
          type: knobType,
          name: event.name,
          tagName,
          instanceIndex
        });
      }
    });
    __privateAdd(this, _handleElementEvent, (event) => {
      const element = event.currentTarget;
      if (!(element instanceof HTMLElement)) return;
      const tagName = element.tagName.toLowerCase();
      const eventInfo = __privateGet(this, _elementEventMap)?.get(tagName);
      if (!eventInfo || !eventInfo.eventNames.has(event.type)) return;
      __privateGet(this, _discoveredElements).add(tagName);
      __privateMethod(this, _CemServeChrome_instances, captureEvent_fn).call(this, event, element, tagName, eventInfo);
    });
    __privateAdd(this, _handleEventTypeFilterChange, (event) => {
      const { value, checked } = event;
      if (!value) return;
      if (checked) {
        __privateGet(this, _eventTypeFilters).add(value);
      } else {
        __privateGet(this, _eventTypeFilters).delete(value);
      }
      __privateMethod(this, _CemServeChrome_instances, saveEventFilters_fn).call(this);
      __privateMethod(this, _CemServeChrome_instances, filterEvents_fn).call(this, __privateGet(this, _eventsFilterValue));
    });
    __privateAdd(this, _handleElementFilterChange, (event) => {
      const { value, checked } = event;
      if (!value) return;
      if (checked) {
        __privateGet(this, _elementFilters).add(value);
      } else {
        __privateGet(this, _elementFilters).delete(value);
      }
      __privateMethod(this, _CemServeChrome_instances, saveEventFilters_fn).call(this);
      __privateMethod(this, _CemServeChrome_instances, filterEvents_fn).call(this, __privateGet(this, _eventsFilterValue));
    });
  }
  get demo() {
    return this.querySelector("cem-serve-demo");
  }
  render() {
    return T`
      <link rel="stylesheet" href="/__cem/cem-pf-v6-c-description-list.css">
      <link rel="stylesheet" href="/__cem/pf-lightdom.css">

      <cem-pf-v6-page ?sidebar-collapsed=${this.sidebar === "collapsed"}>
        <cem-pf-v6-skip-to-content slot="skip-to-content"
                               href="#main-content">
          Skip to content
        </cem-pf-v6-skip-to-content>

        <cem-pf-v6-masthead slot="masthead">
          <a class="masthead-logo"
             href="/"
             slot="logo">
            <img alt="CEM Dev Server"
                 src="/__cem/logo.svg">
            ${this.packageName ? T`<h1>${this.packageName}</h1>` : A}
          </a>
          <cem-pf-v6-toolbar slot="toolbar">
            <cem-pf-v6-toolbar-group variant="action-group">
              ${__privateMethod(this, _CemServeChrome_instances, renderSourceButton_fn).call(this)}
              <cem-pf-v6-toolbar-item>
                <cem-pf-v6-button id="debug-info"
                              variant="plain"
                              aria-label="Debug info">
                  <svg width="16"
                       height="16"
                       viewBox="0 0 16 16"
                       fill="currentColor"
                       role="presentation">
                    <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
                    <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
                  </svg>
                </cem-pf-v6-button>
              </cem-pf-v6-toolbar-item>
              <cem-pf-v6-toolbar-item>
                <cem-pf-v6-toggle-group class="color-scheme-toggle"
                                    aria-label="Color scheme">
                  <cem-pf-v6-toggle-group-item value="light">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="Light mode">
                      <path d="M8 11a3 3 0 1 1 0-6 3 3 0 0 1 0 6zm0 1a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z"/>
                    </svg>
                  </cem-pf-v6-toggle-group-item>
                  <cem-pf-v6-toggle-group-item value="dark">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="Dark mode">
                      <path d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278zM4.858 1.311A7.269 7.269 0 0 0 1.025 7.71c0 4.02 3.279 7.276 7.319 7.276a7.316 7.316 0 0 0 5.205-2.162c-.337.042-.68.063-1.029.063-4.61 0-8.343-3.714-8.343-8.29 0-1.167.242-2.278.681-3.286z"/>
                    </svg>
                  </cem-pf-v6-toggle-group-item>
                  <cem-pf-v6-toggle-group-item value="system">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="System preference">
                      <path d="M0 1.5A1.5 1.5 0 0 1 1.5 0h13A1.5 1.5 0 0 1 16 1.5v8a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 9.5v-8zM1.5 1a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-8a.5.5 0 0 0-.5-.5h-13z"/>
                      <path d="M2.5 12h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1 0-1zm0 2h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1 0-1z"/>
                    </svg>
                  </cem-pf-v6-toggle-group-item>
                </cem-pf-v6-toggle-group>
              </cem-pf-v6-toolbar-item>
            </cem-pf-v6-toolbar-group>
          </cem-pf-v6-toolbar>
        </cem-pf-v6-masthead>

        <cem-pf-v6-page-sidebar slot="sidebar"
                            ?expanded=${this.sidebar === "expanded"}
                            ?collapsed=${this.sidebar !== "expanded"}>
          <slot name="navigation"></slot>
        </cem-pf-v6-page-sidebar>

        <cem-pf-v6-page-main slot="main" id="main-content">
          <slot></slot>
          <footer class="pf-m-sticky-bottom">
            <div class="footer-description${this.hasDescription ? "" : " empty"}">
              <slot name="description"></slot>
            </div>
            <cem-drawer ?open=${this.drawer === "expanded"}
                        drawer-height="${this.drawerHeight || "400"}">
              <cem-pf-v6-tabs selected="${this.tabsSelected || "0"}">
                <cem-pf-v6-tab title="Knobs">
                  <div id="knobs-container">
                    <slot name="knobs">
                      <p class="knobs-empty">No knobs available for this element.</p>
                    </slot>
                  </div>
                </cem-pf-v6-tab>
                <cem-pf-v6-tab title="Manifest Browser">
                  <cem-manifest-browser>
                    <slot name="manifest-tree" slot="manifest-tree"></slot>
                    <slot name="manifest-name" slot="manifest-name"></slot>
                    <slot name="manifest-details" slot="manifest-details"></slot>
                  </cem-manifest-browser>
                </cem-pf-v6-tab>
                <cem-pf-v6-tab title="Server Logs">
                  <div class="logs-wrapper">
                    <cem-pf-v6-toolbar sticky>
                      <cem-pf-v6-toolbar-group>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-text-input-group id="logs-filter"
                                                  placeholder="Filter logs..."
                                                  icon>
                            <svg slot="icon"
                                 role="presentation"
                                 fill="currentColor"
                                 height="1em"
                                 width="1em"
                                 viewBox="0 0 512 512">
                              <path d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                            </svg>
                          </cem-pf-v6-text-input-group>
                        </cem-pf-v6-toolbar-item>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-dropdown id="log-level-filter"
                                          label="Filter log levels">
                            <span slot="toggle-text">Log Levels</span>
                            <cem-pf-v6-menu-item variant="checkbox" value="error" checked>Errors</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="warn" checked>Warnings</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="success" checked>Success</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="info" checked>Info</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="debug" checked>Debug</cem-pf-v6-menu-item>
                            <cem-pf-v6-menu-item variant="checkbox" value="trace" checked>Trace</cem-pf-v6-menu-item>
                          </cem-pf-v6-dropdown>
                        </cem-pf-v6-toolbar-item>
                      </cem-pf-v6-toolbar-group>
                      <cem-pf-v6-toolbar-group variant="action-group">
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-button id="copy-logs"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M13 0H6a2 2 0 0 0-2 2 2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2 2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm0 13V4a2 2 0 0 0-2-2H5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1zM3 13V4a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>
                            </svg>
                            Copy Logs
                          </cem-pf-v6-button>
                        </cem-pf-v6-toolbar-item>
                      </cem-pf-v6-toolbar-group>
                    </cem-pf-v6-toolbar>
                    <div id="log-container"></div>
                  </div>
                </cem-pf-v6-tab>
                <cem-pf-v6-tab title="Events">
                  <div class="events-wrapper">
                    <cem-pf-v6-toolbar sticky>
                      <cem-pf-v6-toolbar-group>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-text-input-group id="events-filter"
                                                  placeholder="Filter events..."
                                                  icon>
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z"/>
                            </svg>
                          </cem-pf-v6-text-input-group>
                        </cem-pf-v6-toolbar-item>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-dropdown id="event-type-filter"
                                          label="Filter event types">
                            <span slot="toggle-text">Event Types</span>
                          </cem-pf-v6-dropdown>
                        </cem-pf-v6-toolbar-item>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-dropdown id="element-filter"
                                          label="Filter elements">
                            <span slot="toggle-text">Elements</span>
                          </cem-pf-v6-dropdown>
                        </cem-pf-v6-toolbar-item>
                      </cem-pf-v6-toolbar-group>
                      <cem-pf-v6-toolbar-group variant="action-group">
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-button id="clear-events"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                              <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                            </svg>
                            Clear Events
                          </cem-pf-v6-button>
                        </cem-pf-v6-toolbar-item>
                        <cem-pf-v6-toolbar-item>
                          <cem-pf-v6-button id="copy-events"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M13 0H6a2 2 0 0 0-2 2 2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2 2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm0 13V4a2 2 0 0 0-2-2H5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1zM3 13V4a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>
                            </svg>
                            Copy Events
                          </cem-pf-v6-button>
                        </cem-pf-v6-toolbar-item>
                      </cem-pf-v6-toolbar-group>
                    </cem-pf-v6-toolbar>
                    <cem-pf-v6-drawer id="event-drawer" expanded>
                      <div id="event-list"></div>
                      <div id="event-detail-header" slot="panel-header"></div>
                      <div id="event-detail-body" slot="panel-body"></div>
                    </cem-pf-v6-drawer>
                  </div>
                </cem-pf-v6-tab>
                <cem-pf-v6-tab title="Health">
                  <cem-health-panel component=${o7(this.primaryTagName)}>
                  </cem-health-panel>
                </cem-pf-v6-tab>
              </cem-pf-v6-tabs>
            </cem-drawer>
          </footer>
        </cem-pf-v6-page-main>
      </cem-pf-v6-page>

      <cem-pf-v6-modal id="debug-modal" variant="large">
        <h2 slot="header">Debug Information</h2>
        <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Server Version</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-version">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Server OS</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-os">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Watch Directory</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-watch-dir">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Manifest Size</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-manifest-size">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Demos Found</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-demo-count">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Tag Name</dt>
            <dd class="cem-pf-v6-c-description-list__description">${this.primaryTagName || "-"}</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Demo Title</dt>
            <dd class="cem-pf-v6-c-description-list__description">${this.demoTitle || "-"}</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">Browser</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-browser">-</dd>
          </div>
          <div class="cem-pf-v6-c-description-list__group">
            <dt class="cem-pf-v6-c-description-list__term">User Agent</dt>
            <dd class="cem-pf-v6-c-description-list__description" id="debug-ua">-</dd>
          </div>
        </dl>
        <div id="demo-urls-container"></div>
        <cem-pf-v6-expandable-section id="debug-importmap-details"
                                  toggle-text="Show Import Map">
          <pre id="debug-importmap">-</pre>
        </cem-pf-v6-expandable-section>
        <div slot="footer" class="button-container">
          <cem-pf-v6-button class="debug-copy" variant="primary">
            Copy Debug Info
          </cem-pf-v6-button>
          <cem-pf-v6-button class="debug-close" variant="secondary">
            Close
          </cem-pf-v6-button>
        </div>
      </cem-pf-v6-modal>

      <!-- Reconnection modal -->
      <cem-pf-v6-modal id="reconnection-modal" variant="large">
        <h2 slot="header">Development Server Disconnected</h2>
        <cem-reconnection-content id="reconnection-content"></cem-reconnection-content>
        <cem-pf-v6-button id="reload-button"
                      slot="footer"
                      variant="primary">Reload Page</cem-pf-v6-button>
        <cem-pf-v6-button id="retry-button"
                      slot="footer"
                      variant="secondary">Retry Now</cem-pf-v6-button>
      </cem-pf-v6-modal>

      <!-- Transform error overlay -->
      <cem-transform-error-overlay id="error-overlay">
      </cem-transform-error-overlay>
    `;
  }
  async connectedCallback() {
    await __privateGet(this, _modulesReady);
    super.connectedCallback();
    if (__privateGet(this, _wsClient) == null) {
      __privateMethod(this, _CemServeChrome_instances, initWsClient_fn).call(this);
    }
    __privateMethod(this, _CemServeChrome_instances, migrateFromLocalStorageIfNeeded_fn).call(this);
  }
  firstUpdated() {
    __privateMethod(this, _CemServeChrome_instances, setupDebugOverlay_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupColorSchemeToggle_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupFooterDrawer_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupLogListener_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupKnobCoordination_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupTreeStatePersistence_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupSidebarStatePersistence_fn).call(this);
    __privateMethod(this, _CemServeChrome_instances, setupEventCapture_fn).call(this).then(() => {
      __privateMethod(this, _CemServeChrome_instances, setupEventListeners_fn).call(this);
    });
    __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "reload-button")?.addEventListener("click", () => {
      window.location.reload();
    });
    __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "retry-button")?.addEventListener("click", () => {
      __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "reconnection-modal")?.close();
      __privateGet(this, _wsClient).retry();
    });
    __privateGet(this, _wsClient).init();
  }
  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener("knob:attribute-change", __privateGet(this, _onKnobChange));
    this.removeEventListener("knob:property-change", __privateGet(this, _onKnobChange));
    this.removeEventListener("knob:css-property-change", __privateGet(this, _onKnobChange));
    this.removeEventListener("knob:css-state-change", __privateGet(this, _onKnobChange));
    this.removeEventListener("knob:attribute-clear", __privateGet(this, _onKnobClear));
    this.removeEventListener("knob:property-clear", __privateGet(this, _onKnobClear));
    this.removeEventListener("knob:css-property-clear", __privateGet(this, _onKnobClear));
    this.removeEventListener("knob:css-state-clear", __privateGet(this, _onKnobClear));
    if (__privateGet(this, _handleTreeExpand)) {
      this.removeEventListener("expand", __privateGet(this, _handleTreeExpand));
    }
    if (__privateGet(this, _handleTreeCollapse)) {
      this.removeEventListener("collapse", __privateGet(this, _handleTreeCollapse));
    }
    if (__privateGet(this, _handleTreeSelect)) {
      this.removeEventListener("select", __privateGet(this, _handleTreeSelect));
    }
    if (__privateGet(this, _handleLogsEvent)) {
      window.removeEventListener("cem:logs", __privateGet(this, _handleLogsEvent));
    }
    if (__privateGet(this, _copyLogsFeedbackTimeout)) {
      clearTimeout(__privateGet(this, _copyLogsFeedbackTimeout));
      __privateSet(this, _copyLogsFeedbackTimeout, null);
    }
    if (__privateGet(this, _copyDebugFeedbackTimeout)) {
      clearTimeout(__privateGet(this, _copyDebugFeedbackTimeout));
      __privateSet(this, _copyDebugFeedbackTimeout, null);
    }
    if (__privateGet(this, _copyEventsFeedbackTimeout)) {
      clearTimeout(__privateGet(this, _copyEventsFeedbackTimeout));
      __privateSet(this, _copyEventsFeedbackTimeout, null);
    }
    __privateGet(this, _observer2).disconnect();
    if (__privateGet(this, _wsClient)) {
      __privateGet(this, _wsClient).destroy();
    }
  }
};
_init54 = __decoratorStart(_a54);
_demoInfoTemplate = new WeakMap();
_demoGroupTemplate = new WeakMap();
_demoListTemplate = new WeakMap();
_logEntryTemplate = new WeakMap();
_eventEntryTemplate = new WeakMap();
_primaryTagName = new WeakMap();
_demoTitle = new WeakMap();
_packageName = new WeakMap();
_canonicalURL = new WeakMap();
_sourceURL = new WeakMap();
_knobs = new WeakMap();
_drawer = new WeakMap();
_drawerHeight2 = new WeakMap();
_tabsSelected = new WeakMap();
_sidebar = new WeakMap();
_hasDescription2 = new WeakMap();
_CemServeChrome_instances = new WeakSet();
$_fn2 = function(id) {
  return this.shadowRoot?.getElementById(id);
};
$$_fn = function(selector) {
  return this.shadowRoot?.querySelectorAll(selector) ?? [];
};
_logContainer = new WeakMap();
_drawerOpen = new WeakMap();
_initialLogsFetched = new WeakMap();
_isInitialLoad = new WeakMap();
_debugData = new WeakMap();
_elementEventMap = new WeakMap();
_manifest = new WeakMap();
_capturedEvents = new WeakMap();
_maxCapturedEvents = new WeakMap();
_eventList = new WeakMap();
_eventDetailHeader = new WeakMap();
_eventDetailBody = new WeakMap();
_selectedEventId = new WeakMap();
_eventsFilterValue = new WeakMap();
_eventsFilterDebounceTimer = new WeakMap();
_eventTypeFilters = new WeakMap();
_elementFilters = new WeakMap();
_discoveredElements = new WeakMap();
_handleLogsEvent = new WeakMap();
_handleTreeExpand = new WeakMap();
_handleTreeCollapse = new WeakMap();
_handleTreeSelect = new WeakMap();
_copyLogsFeedbackTimeout = new WeakMap();
_copyDebugFeedbackTimeout = new WeakMap();
_copyEventsFeedbackTimeout = new WeakMap();
_logsFilterValue = new WeakMap();
_logsFilterDebounceTimer = new WeakMap();
_logLevelFilters = new WeakMap();
_logLevelDropdown = new WeakMap();
_observer2 = new WeakMap();
_wsClient = new WeakMap();
initWsClient_fn = function() {
  __privateSet(this, _wsClient, new CEMReloadClient4({
    jitterMax: 1e3,
    overlayThreshold: 15,
    badgeFadeDelay: 2e3,
    /* c8 ignore start - WebSocket callbacks tested via integration */
    callbacks: {
      onOpen: () => {
        __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "reconnection-modal")?.close();
      },
      onError: (errorData) => {
        if (errorData?.title && errorData?.message) {
          console.error("[cem-serve] Server error:", errorData);
          __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "error-overlay")?.show(errorData.title, errorData.message, errorData.file);
        } else {
          console.error("[cem-serve] WebSocket error:", errorData);
        }
      },
      onReconnecting: ({ attempt, delay }) => {
        if (attempt >= 15) {
          __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "reconnection-modal")?.showModal();
          __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "reconnection-content")?.updateRetryInfo(attempt, delay);
        }
      },
      onReload: () => {
        const errorOverlay = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "error-overlay");
        if (errorOverlay?.hasAttribute("open")) {
          errorOverlay.hide();
        }
        window.location.reload();
      },
      onShutdown: () => {
        __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "reconnection-modal")?.showModal();
        __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "reconnection-content")?.updateRetryInfo(30, 3e4);
      },
      onLogs: (logs) => {
        window.dispatchEvent(new CemLogsEvent(logs));
      }
    }
    /* c8 ignore stop */
  }));
};
renderSourceButton_fn = function() {
  if (!this.sourceURL) return A;
  let label = "Version Control";
  let path = "M5.854 4.854a.5.5 0 1 0-.708-.708l-3.5 3.5a.5.5 0 0 0 0 .708l3.5 3.5a.5.5 0 0 0 .708-.708L2.707 8l3.147-3.146zm4.292 0a.5.5 0 0 1 .708-.708l3.5 3.5a.5.5 0 0 1 0 .708l-3.5 3.5a.5.5 0 0 1-.708-.708L13.293 8l-3.147-3.146z";
  if (this.sourceURL.includes("github.com")) {
    label = "GitHub.com";
    path = "M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z";
  } else if (this.sourceURL.includes("gitlab.com")) {
    label = "GitLab";
    path = "m15.734 6.1-.022-.058L13.534.358a.568.568 0 0 0-.563-.356.583.583 0 0 0-.328.122.582.582 0 0 0-.193.294l-1.47 4.499H5.025l-1.47-4.5A.572.572 0 0 0 3.360.174a.572.572 0 0 0-.328-.172.582.582 0 0 0-.563.357L.29 6.04l-.022.057A4.044 4.044 0 0 0 1.61 10.77l.007.006.02.014 3.318 2.485 1.64 1.242 1 .755a.673.673 0 0 0 .814 0l1-.755 1.64-1.242 3.338-2.5.009-.007a4.05 4.05 0 0 0 1.34-4.668Z";
  } else if (this.sourceURL.includes("bitbucket.org")) {
    label = "Bitbucket";
    path = "M0 1.5A1.5 1.5 0 0 1 1.5 0h13A1.5 1.5 0 0 1 16 1.5v13a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 14.5v-13zM2.5 3a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-11zm5.038 1.435a.5.5 0 0 1 .924 0l1.42 3.37H8.78l-.243.608-.243-.608H5.082l1.42-3.37zM8 9.143l-.743 1.857H4.743L6.076 7.608 8 9.143z";
  }
  return T`
      <cem-pf-v6-toolbar-item>
        <cem-pf-v6-button href="${this.sourceURL}"
                      target="_blank"
                      rel="noopener noreferrer"
                      variant="plain"
                      aria-label="View source file">
          <svg aria-label="${label}"
               width="1rem"
               height="1rem"
               fill="currentColor"
               viewBox="0 0 16 16">
            <path d="${path}"/>
          </svg>
        </cem-pf-v6-button>
      </cem-pf-v6-toolbar-item>
    `;
};
_modulesReady = new WeakMap();
loadClientModules_fn = function() {
  return Promise.all([
    // @ts-ignore -- plain JS modules served at runtime by Go server
    Promise.resolve().then(() => (init_websocket_client(), websocket_client_exports)),
    // @ts-ignore
    Promise.resolve().then(() => (init_state_persistence(), state_persistence_exports))
  ]).then(([ws, sp]) => {
    CEMReloadClient4 = ws.CEMReloadClient;
    StatePersistence4 = sp.StatePersistence;
    Promise.resolve().then(() => (init_health_badges(), health_badges_exports)).catch((e6) => console.error("[cem-serve] Failed to load health-badges:", e6));
  }).catch((e6) => {
    console.error("[cem-serve] Failed to load client modules:", e6);
    CEMReloadClient4 ??= class {
      init() {
      }
      retry() {
      }
      destroy() {
      }
    };
    StatePersistence4 ??= {
      getState: () => ({ colorScheme: "system" }),
      updateState() {
      },
      getTreeState: () => ({ expanded: [], selected: null }),
      setTreeState() {
      },
      updateTreeState() {
      },
      migrateFromLocalStorage() {
      }
    };
  });
};
fetchDebugInfo_fn = async function() {
  const browserEl = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-browser");
  const uaEl = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-ua");
  if (browserEl) {
    const browser = __privateMethod(this, _CemServeChrome_instances, detectBrowser_fn).call(this);
    browserEl.textContent = browser;
  }
  if (uaEl) {
    uaEl.textContent = navigator.userAgent;
  }
  const data = await fetch("/__cem/debug").then((res) => res.json()).catch((err) => {
    console.error("[cem-serve-chrome] Failed to fetch debug info:", err);
  });
  if (!data) return;
  const versionEl = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-version");
  const osEl = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-os");
  const watchDirEl = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-watch-dir");
  const manifestSizeEl = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-manifest-size");
  const demoCountEl = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-demo-count");
  const demoUrlsContainer = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "demo-urls-container");
  const importMapEl = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-importmap");
  if (versionEl) versionEl.textContent = data.version || "-";
  if (osEl) osEl.textContent = data.os || "-";
  if (watchDirEl) watchDirEl.textContent = data.watchDir || "-";
  if (manifestSizeEl) manifestSizeEl.textContent = data.manifestSize || "-";
  if (demoCountEl) demoCountEl.textContent = data.demoCount || "0";
  if (demoUrlsContainer) {
    __privateMethod(this, _CemServeChrome_instances, populateDemoUrls_fn).call(this, demoUrlsContainer, data.demos);
  }
  if (importMapEl && data.importMap) {
    const importMapJSON = JSON.stringify(data.importMap, null, 2);
    importMapEl.textContent = importMapJSON;
    data.importMapJSON = importMapJSON;
  } else if (importMapEl) {
    importMapEl.textContent = "No import map generated";
  }
  __privateSet(this, _debugData, data);
};
populateDemoUrls_fn = function(container, demos) {
  if (!demos?.length) {
    container.textContent = "No demos found in manifest";
    return;
  }
  const currentTagName = this.primaryTagName || "";
  const isOnDemoPage = !!currentTagName;
  if (isOnDemoPage) {
    const currentDemo = demos.find((demo) => demo.tagName === currentTagName);
    if (!currentDemo) {
      container.textContent = "Current demo not found in manifest";
      return;
    }
    const fragment = __privateGet(_CemServeChrome, _demoInfoTemplate).content.cloneNode(true);
    fragment.querySelector('[data-field="tagName"]').textContent = currentDemo.tagName;
    fragment.querySelector('[data-field="canonicalURL"]').textContent = currentDemo.canonicalURL;
    fragment.querySelector('[data-field="localRoute"]').textContent = currentDemo.localRoute;
    const descriptionGroup = fragment.querySelector('[data-field-group="description"]');
    if (currentDemo.description) {
      fragment.querySelector('[data-field="description"]').textContent = currentDemo.description;
    } else {
      descriptionGroup?.remove();
    }
    const filepathGroup = fragment.querySelector('[data-field-group="filepath"]');
    if (currentDemo.filepath) {
      fragment.querySelector('[data-field="filepath"]').textContent = currentDemo.filepath;
    } else {
      filepathGroup?.remove();
    }
    container.replaceChildren(fragment);
  } else {
    const listFragment = __privateGet(_CemServeChrome, _demoListTemplate).content.cloneNode(true);
    const groupsContainer = listFragment.querySelector('[data-container="groups"]');
    for (const demo of demos) {
      const groupFragment = __privateGet(_CemServeChrome, _demoGroupTemplate).content.cloneNode(true);
      groupFragment.querySelector('[data-field="tagName"]').textContent = demo.tagName;
      groupFragment.querySelector('[data-field="description"]').textContent = demo.description || "(no description)";
      groupFragment.querySelector('[data-field="canonicalURL"]').textContent = demo.canonicalURL;
      groupFragment.querySelector('[data-field="localRoute"]').textContent = demo.localRoute;
      const filepathGroup = groupFragment.querySelector('[data-field-group="filepath"]');
      if (demo.filepath) {
        groupFragment.querySelector('[data-field="filepath"]').textContent = demo.filepath;
      } else {
        filepathGroup?.remove();
      }
      groupsContainer.appendChild(groupFragment);
    }
    container.replaceChildren(listFragment);
  }
};
setupLogListener_fn = function() {
  __privateSet(this, _logContainer, __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "log-container"));
  const logsFilter = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "logs-filter");
  if (logsFilter) {
    logsFilter.addEventListener("input", () => {
      const { value = "" } = logsFilter;
      clearTimeout(__privateGet(this, _logsFilterDebounceTimer));
      __privateSet(this, _logsFilterDebounceTimer, setTimeout(() => {
        __privateMethod(this, _CemServeChrome_instances, filterLogs_fn).call(this, value);
      }, 300));
    });
  }
  __privateSet(this, _logLevelDropdown, this.shadowRoot?.querySelector("#log-level-filter") ?? null);
  if (__privateGet(this, _logLevelDropdown)) {
    requestAnimationFrame(() => {
      __privateMethod(this, _CemServeChrome_instances, loadLogFilterState_fn).call(this);
    });
    __privateGet(this, _logLevelDropdown).addEventListener("select", __privateGet(this, _handleLogFilterChange));
  }
  __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "copy-logs")?.addEventListener("click", () => {
    __privateMethod(this, _CemServeChrome_instances, copyLogs_fn).call(this);
  });
  __privateSet(this, _handleLogsEvent, ((event) => {
    const logs = event.logs;
    if (logs) {
      __privateMethod(this, _CemServeChrome_instances, renderLogs_fn).call(this, logs);
    }
  }));
  window.addEventListener("cem:logs", __privateGet(this, _handleLogsEvent));
};
filterLogs_fn = function(query) {
  __privateSet(this, _logsFilterValue, query.toLowerCase());
  if (!__privateGet(this, _logContainer)) return;
  for (const entry of __privateGet(this, _logContainer).children) {
    const text = entry.textContent?.toLowerCase() ?? "";
    const textMatch = !__privateGet(this, _logsFilterValue) || text.includes(__privateGet(this, _logsFilterValue));
    const logType = __privateMethod(this, _CemServeChrome_instances, getLogTypeFromEntry_fn).call(this, entry);
    const levelMatch = __privateGet(this, _logLevelFilters).has(logType);
    entry.hidden = !(textMatch && levelMatch);
  }
};
getLogTypeFromEntry_fn = function(entry) {
  for (const cls of entry.classList) {
    if (["info", "warning", "error", "success", "debug", "trace"].includes(cls)) {
      return cls === "warning" ? "warn" : cls;
    }
  }
  return "info";
};
loadLogFilterState_fn = function() {
  try {
    const saved = localStorage.getItem("cem-serve-log-filters");
    if (saved) {
      __privateSet(this, _logLevelFilters, new Set(JSON.parse(saved)));
    }
  } catch (e6) {
    console.debug("[cem-serve-chrome] localStorage unavailable, using default log filters");
  }
  __privateMethod(this, _CemServeChrome_instances, syncCheckboxStates_fn).call(this);
};
syncCheckboxStates_fn = function() {
  if (!__privateGet(this, _logLevelDropdown)) return;
  const menuItems = __privateGet(this, _logLevelDropdown).querySelectorAll("cem-pf-v6-menu-item");
  menuItems.forEach((item) => {
    const value = item.value;
    item.checked = __privateGet(this, _logLevelFilters).has(value);
  });
};
saveLogFilterState_fn = function() {
  try {
    localStorage.setItem(
      "cem-serve-log-filters",
      JSON.stringify([...__privateGet(this, _logLevelFilters)])
    );
  } catch (e6) {
  }
};
_handleLogFilterChange = new WeakMap();
copyLogs_fn = async function() {
  if (!__privateGet(this, _logContainer)) return;
  const logs = Array.from(__privateGet(this, _logContainer).children).filter((entry) => !entry.hidden).map((entry) => {
    const type = entry.querySelector('[data-field="label"]')?.textContent?.trim() || "INFO";
    const time = entry.querySelector('[data-field="time"]')?.textContent?.trim() || "";
    const message = entry.querySelector('[data-field="message"]')?.textContent?.trim() || "";
    return `[${type}] ${time} ${message}`;
  }).join("\n");
  if (!logs) return;
  try {
    await navigator.clipboard.writeText(logs);
    const btn = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "copy-logs");
    if (btn) {
      const textNode = Array.from(btn.childNodes).find(
        (n6) => n6.nodeType === Node.TEXT_NODE && (n6.textContent?.trim().length ?? 0) > 0
      );
      if (textNode) {
        const original = textNode.textContent;
        textNode.textContent = "Copied!";
        if (__privateGet(this, _copyLogsFeedbackTimeout)) {
          clearTimeout(__privateGet(this, _copyLogsFeedbackTimeout));
        }
        __privateSet(this, _copyLogsFeedbackTimeout, setTimeout(() => {
          if (this.isConnected && textNode.parentNode) {
            textNode.textContent = original;
          }
          __privateSet(this, _copyLogsFeedbackTimeout, null);
        }, 2e3));
      }
    }
  } catch (err) {
    console.error("[cem-serve-chrome] Failed to copy logs:", err);
  }
};
setupDebugOverlay_fn = function() {
  const debugButton = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-info");
  const debugModal = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "debug-modal");
  const debugClose = this.shadowRoot?.querySelector(".debug-close");
  const debugCopy = this.shadowRoot?.querySelector(".debug-copy");
  if (debugButton && debugModal) {
    debugButton.addEventListener("click", () => {
      __privateMethod(this, _CemServeChrome_instances, fetchDebugInfo_fn).call(this);
      debugModal.showModal();
    });
    debugClose?.addEventListener("click", () => debugModal.close());
    debugCopy?.addEventListener("click", () => {
      __privateMethod(this, _CemServeChrome_instances, copyDebugInfo_fn).call(this);
    });
  }
};
setupFooterDrawer_fn = function() {
  const drawer = this.shadowRoot?.querySelector("cem-drawer");
  const tabs = this.shadowRoot?.querySelector("cem-pf-v6-tabs");
  if (!drawer || !tabs) return;
  __privateSet(this, _drawerOpen, drawer.open);
  drawer.addEventListener("change", (e6) => {
    __privateSet(this, _drawerOpen, e6.open);
    StatePersistence4.updateState({
      drawer: { open: e6.open }
    });
    if (e6.open) {
      __privateMethod(this, _CemServeChrome_instances, scrollLogsToBottom_fn).call(this);
    }
  });
  drawer.addEventListener("resize", (e6) => {
    drawer.setAttribute("drawer-height", e6.height);
    StatePersistence4.updateState({
      drawer: { height: e6.height }
    });
  });
  tabs.addEventListener("change", (e6) => {
    StatePersistence4.updateState({
      tabs: { selectedIndex: e6.selectedIndex }
    });
    if (e6.selectedIndex === 2 && drawer.open) {
      __privateMethod(this, _CemServeChrome_instances, scrollLogsToBottom_fn).call(this);
    }
    if (e6.selectedIndex === 3 && drawer.open) {
      __privateMethod(this, _CemServeChrome_instances, scrollEventsToBottom_fn).call(this);
    }
  });
};
detectBrowser_fn = function() {
  const ua = navigator.userAgent;
  if (ua.includes("Firefox/")) {
    const match = ua.match(/Firefox\/(\d+)/);
    return match ? `Firefox ${match[1]}` : "Firefox";
  } else if (ua.includes("Edg/")) {
    const match = ua.match(/Edg\/(\d+)/);
    return match ? `Edge ${match[1]}` : "Edge";
  } else if (ua.includes("Chrome/")) {
    const match = ua.match(/Chrome\/(\d+)/);
    return match ? `Chrome ${match[1]}` : "Chrome";
  } else if (ua.includes("Safari/") && !ua.includes("Chrome")) {
    const match = ua.match(/Version\/(\d+)/);
    return match ? `Safari ${match[1]}` : "Safari";
  }
  return "Unknown";
};
copyDebugInfo_fn = async function() {
  const info = Array.from(__privateMethod(this, _CemServeChrome_instances, $$_fn).call(this, "#debug-modal dl dt")).map((dt) => {
    const dd = dt.nextElementSibling;
    if (dd && dd.tagName === "DD") {
      return `${dt.textContent}: ${dd.textContent}`;
    }
    return "";
  }).join("\n");
  let importMapSection = "";
  if (__privateGet(this, _debugData)?.importMapJSON) {
    importMapSection = `
${"=".repeat(40)}
Import Map:
${"=".repeat(40)}
${__privateGet(this, _debugData).importMapJSON}
`;
  }
  const debugText = `CEM Serve Debug Information
${"=".repeat(40)}
${info}${importMapSection}
${"=".repeat(40)}
Generated: ${(/* @__PURE__ */ new Date()).toISOString()}`;
  try {
    await navigator.clipboard.writeText(debugText);
    const copyButton = this.shadowRoot?.querySelector(".debug-copy");
    if (copyButton) {
      const originalText = copyButton.textContent;
      copyButton.textContent = "Copied!";
      if (__privateGet(this, _copyDebugFeedbackTimeout)) {
        clearTimeout(__privateGet(this, _copyDebugFeedbackTimeout));
      }
      __privateSet(this, _copyDebugFeedbackTimeout, setTimeout(() => {
        if (this.isConnected && copyButton.parentNode) {
          copyButton.textContent = originalText;
        }
        __privateSet(this, _copyDebugFeedbackTimeout, null);
      }, 2e3));
    }
  } catch (err) {
    console.error("[cem-serve-chrome] Failed to copy debug info:", err);
  }
};
renderLogs_fn = function(logs) {
  if (!__privateGet(this, _logContainer)) return;
  const logElements = logs.map((log) => {
    const fragment = __privateGet(_CemServeChrome, _logEntryTemplate).content.cloneNode(true);
    const date = new Date(log.date);
    const time = date.toLocaleTimeString();
    const container = fragment.querySelector('[data-field="container"]');
    container.classList.add(log.type);
    container.setAttribute("data-log-id", log.date);
    const typeLabel = __privateMethod(this, _CemServeChrome_instances, getLogBadge_fn).call(this, log.type);
    const searchText = `${typeLabel} ${time} ${log.message}`.toLowerCase();
    const textMatch = !__privateGet(this, _logsFilterValue) || searchText.includes(__privateGet(this, _logsFilterValue));
    const logTypeForFilter = log.type === "warning" ? "warn" : log.type;
    const levelMatch = __privateGet(this, _logLevelFilters).has(logTypeForFilter);
    if (!(textMatch && levelMatch)) {
      container.setAttribute("hidden", "");
    }
    const label = fragment.querySelector('[data-field="label"]');
    label.textContent = __privateMethod(this, _CemServeChrome_instances, getLogBadge_fn).call(this, log.type);
    __privateMethod(this, _CemServeChrome_instances, applyLogLabelAttrs_fn).call(this, label, log.type);
    const timeEl = fragment.querySelector('[data-field="time"]');
    timeEl.setAttribute("datetime", log.date);
    timeEl.textContent = time;
    const messageEl = fragment.querySelector('[data-field="message"]');
    if (log.data?.kind === "durations" && log.data.durations?.length) {
      B(T`
           <div class="progress-grid">${log.data.durations.map((d3) => T`
            <span>${d3.name}</span>
            <cem-pf-v6-progress value="${Math.round(d3.percent)}"
                                value-text="${d3.duration}"
                                measure-location="outside"
                                size="sm"></cem-pf-v6-progress>`)}
          </div>`, messageEl);
    } else {
      messageEl.textContent = log.message;
    }
    return fragment;
  });
  if (!__privateGet(this, _initialLogsFetched)) {
    __privateGet(this, _logContainer).replaceChildren(...logElements);
    __privateSet(this, _initialLogsFetched, true);
    if (__privateGet(this, _drawerOpen)) {
      __privateMethod(this, _CemServeChrome_instances, scrollLatestIntoView_fn).call(this);
    }
  } else {
    __privateGet(this, _logContainer).append(...logElements);
    if (__privateGet(this, _drawerOpen)) {
      __privateMethod(this, _CemServeChrome_instances, scrollLatestIntoView_fn).call(this);
    }
  }
};
getLogBadge_fn = function(type) {
  switch (type) {
    case "info":
      return "Info";
    case "warning":
      return "Warn";
    case "error":
      return "Error";
    case "success":
      return "OK";
    case "debug":
      return "Debug";
    case "trace":
      return "Trace";
    default:
      return type.toUpperCase();
  }
};
applyLogLabelAttrs_fn = function(label, type) {
  switch (type) {
    case "info":
      return label.setAttribute("status", "info");
    case "warning":
      return label.setAttribute("status", "warning");
    case "error":
      return label.setAttribute("status", "danger");
    case "success":
      return label.setAttribute("status", "success");
    case "debug":
      return label.setAttribute("color", "purple");
    case "trace":
      return label.setAttribute("color", "grey");
    default:
      label.setAttribute("color", "grey");
  }
};
scrollLatestIntoView_fn = function() {
  if (!__privateGet(this, _logContainer)) return;
  requestAnimationFrame(() => {
    const lastLog = __privateGet(this, _logContainer).lastElementChild;
    if (lastLog) {
      lastLog.scrollIntoView({ behavior: "auto", block: "end" });
    }
  });
};
scrollLogsToBottom_fn = function() {
  if (!__privateGet(this, _logContainer)) return;
  if (__privateGet(this, _isInitialLoad)) {
    __privateMethod(this, _CemServeChrome_instances, scrollLatestIntoView_fn).call(this);
  } else {
    setTimeout(() => {
      __privateMethod(this, _CemServeChrome_instances, scrollLatestIntoView_fn).call(this);
    }, 350);
  }
};
migrateFromLocalStorageIfNeeded_fn = function() {
  try {
    const hasLocalStorage = localStorage.getItem("cem-serve-color-scheme") !== null || localStorage.getItem("cem-serve-drawer-open") !== null || localStorage.getItem("cem-serve-drawer-height") !== null || localStorage.getItem("cem-serve-active-tab") !== null;
    if (hasLocalStorage) {
      const migrated = localStorage.getItem("cem-serve-migrated-to-cookies");
      if (!migrated) {
        StatePersistence4.migrateFromLocalStorage();
        localStorage.setItem("cem-serve-migrated-to-cookies", "true");
        setTimeout(() => window.location.reload(), 100);
      }
    }
  } catch (e6) {
  }
};
setupColorSchemeToggle_fn = function() {
  const toggleGroup = this.shadowRoot?.querySelector(".color-scheme-toggle");
  if (!toggleGroup) return;
  const state = StatePersistence4.getState();
  __privateMethod(this, _CemServeChrome_instances, applyColorScheme_fn).call(this, state.colorScheme);
  const items = toggleGroup.querySelectorAll("cem-pf-v6-toggle-group-item");
  items.forEach((item) => {
    if (item.value === state.colorScheme) {
      item.setAttribute("selected", "");
    }
  });
  toggleGroup.addEventListener("cem-pf-v6-toggle-group-change", (e6) => {
    const scheme = e6.value;
    __privateMethod(this, _CemServeChrome_instances, applyColorScheme_fn).call(this, scheme);
    StatePersistence4.updateState({ colorScheme: scheme });
  });
};
applyColorScheme_fn = function(scheme) {
  const body = document.body;
  switch (scheme) {
    case "light":
      body.style.colorScheme = "light";
      break;
    case "dark":
      body.style.colorScheme = "dark";
      break;
    case "system":
    default:
      body.style.colorScheme = "light dark";
      break;
  }
};
setupKnobCoordination_fn = function() {
  this.addEventListener("knob:attribute-change", __privateGet(this, _onKnobChange));
  this.addEventListener("knob:property-change", __privateGet(this, _onKnobChange));
  this.addEventListener("knob:css-property-change", __privateGet(this, _onKnobChange));
  this.addEventListener("knob:css-state-change", __privateGet(this, _onKnobChange));
  this.addEventListener("knob:attribute-clear", __privateGet(this, _onKnobClear));
  this.addEventListener("knob:property-clear", __privateGet(this, _onKnobClear));
  this.addEventListener("knob:css-property-clear", __privateGet(this, _onKnobClear));
  this.addEventListener("knob:css-state-clear", __privateGet(this, _onKnobClear));
};
_onKnobChange = new WeakMap();
_onKnobClear = new WeakMap();
disableCssStateKnob_fn = function(event) {
  const name = event.name;
  for (const el of event.composedPath()) {
    if (!(el instanceof HTMLElement)) continue;
    if (el.dataset?.isElementKnob !== "true") continue;
    const control = el.querySelector(
      `[data-knob-type="css-state"][data-knob-name="${CSS.escape(name)}"]`
    );
    if (control) {
      control.checked = false;
      control.disabled = true;
      control.title = "Element does not expose ElementInternals";
      const clearBtn = control.closest("cem-pf-v6-form-group")?.querySelector(`.knob-clear-button[data-knob-type="css-state"][data-knob-name="${CSS.escape(name)}"]`);
      if (clearBtn) clearBtn.hidden = true;
    }
    break;
  }
};
getKnobTarget_fn = function(event) {
  const defaultTagName = this.primaryTagName || "";
  if (event.composedPath) {
    for (const element of event.composedPath()) {
      if (!(element instanceof Element)) continue;
      if (element.dataset?.isElementKnob === "true") {
        const tagName = element.dataset.tagName || defaultTagName;
        let instanceIndex = Number.parseInt(element.dataset.instanceIndex ?? "", 10);
        if (Number.isNaN(instanceIndex)) instanceIndex = 0;
        return { tagName, instanceIndex };
      }
    }
  }
  return { tagName: defaultTagName, instanceIndex: 0 };
};
getKnobTypeFromEvent_fn = function(event) {
  switch (event.type) {
    case "knob:attribute-change":
      return "attribute";
    case "knob:property-change":
      return "property";
    case "knob:css-property-change":
      return "css-property";
    case "knob:css-state-change":
      return "css-state";
    default:
      return "unknown";
  }
};
getKnobTypeFromClearEvent_fn = function(event) {
  switch (event.type) {
    case "knob:attribute-clear":
      return "attribute";
    case "knob:property-clear":
      return "property";
    case "knob:css-property-clear":
      return "css-property";
    case "knob:css-state-clear":
      return "css-state";
    default:
      return "unknown";
  }
};
setupTreeStatePersistence_fn = function() {
  __privateSet(this, _handleTreeExpand, (e6) => {
    if (e6.target?.tagName !== "CEM-PF-V6-TREE-ITEM") return;
    const nodeId = __privateMethod(this, _CemServeChrome_instances, getTreeNodeId_fn).call(this, e6.target);
    const treeState = StatePersistence4.getTreeState();
    if (!treeState.expanded.includes(nodeId)) {
      treeState.expanded.push(nodeId);
      StatePersistence4.setTreeState(treeState);
    }
  });
  this.addEventListener("expand", __privateGet(this, _handleTreeExpand));
  __privateSet(this, _handleTreeCollapse, (e6) => {
    if (e6.target?.tagName !== "CEM-PF-V6-TREE-ITEM") return;
    const nodeId = __privateMethod(this, _CemServeChrome_instances, getTreeNodeId_fn).call(this, e6.target);
    const treeState = StatePersistence4.getTreeState();
    const index = treeState.expanded.indexOf(nodeId);
    if (index > -1) {
      treeState.expanded.splice(index, 1);
      StatePersistence4.setTreeState(treeState);
    }
  });
  this.addEventListener("collapse", __privateGet(this, _handleTreeCollapse));
  __privateSet(this, _handleTreeSelect, (e6) => {
    if (e6.target?.tagName !== "CEM-PF-V6-TREE-ITEM") return;
    const nodeId = __privateMethod(this, _CemServeChrome_instances, getTreeNodeId_fn).call(this, e6.target);
    StatePersistence4.updateTreeState({ selected: nodeId });
  });
  this.addEventListener("select", __privateGet(this, _handleTreeSelect));
  __privateMethod(this, _CemServeChrome_instances, applyTreeState_fn).call(this);
};
applyTreeState_fn = function() {
  const treeState = StatePersistence4.getTreeState();
  for (const nodeId of treeState.expanded) {
    const treeItem = __privateMethod(this, _CemServeChrome_instances, findTreeItemById_fn).call(this, nodeId);
    if (treeItem && !treeItem.hasAttribute("expanded")) {
      treeItem.setAttribute("expanded", "");
    }
  }
  if (treeState.selected) {
    const treeItem = __privateMethod(this, _CemServeChrome_instances, findTreeItemById_fn).call(this, treeState.selected);
    if (treeItem && !treeItem.hasAttribute("current")) {
      treeItem.setAttribute("current", "");
    }
  }
};
setupSidebarStatePersistence_fn = function() {
  const page = this.shadowRoot?.querySelector("cem-pf-v6-page");
  if (!page) return;
  page.addEventListener("sidebar-toggle", (event) => {
    const collapsed = !event.expanded;
    StatePersistence4.updateState({
      sidebar: { collapsed }
    });
  });
};
findTreeItemById_fn = function(nodeId) {
  const parts = nodeId.split(":");
  const [type, modulePath, tagName, name] = parts;
  let attrSuffix = "";
  if (tagName) {
    attrSuffix += `[data-tag-name="${CSS.escape(tagName)}"]`;
  }
  if (name) {
    attrSuffix += `[data-name="${CSS.escape(name)}"]`;
  }
  let selector = `cem-pf-v6-tree-item[data-type="${CSS.escape(type)}"]`;
  if (modulePath) {
    const escapedModulePath = CSS.escape(modulePath);
    const escapedType = CSS.escape(type);
    const selector1 = `cem-pf-v6-tree-item[data-type="${escapedType}"][data-module-path="${escapedModulePath}"]${attrSuffix}`;
    const selector2 = `cem-pf-v6-tree-item[data-type="${escapedType}"][data-path="${escapedModulePath}"]${attrSuffix}`;
    selector = `${selector1}, ${selector2}`;
  } else {
    selector += attrSuffix;
  }
  return this.querySelector(selector);
};
getTreeNodeId_fn = function(treeItem) {
  const type = treeItem.getAttribute("data-type");
  const modulePath = treeItem.getAttribute("data-module-path") || treeItem.getAttribute("data-path");
  const tagName = treeItem.getAttribute("data-tag-name");
  const name = treeItem.getAttribute("data-name");
  const category = treeItem.getAttribute("data-category");
  const parts = [type];
  if (modulePath) parts.push(modulePath);
  if (tagName) parts.push(tagName);
  if (category) {
    parts.push(category);
  } else if (name) {
    parts.push(name);
  }
  return parts.join(":");
};
discoverElementEvents_fn = async function() {
  try {
    const response = await fetch("/custom-elements.json");
    if (!response.ok) {
      console.warn("[cem-serve-chrome] No manifest available for event discovery");
      return /* @__PURE__ */ new Map();
    }
    const manifest = await response.json();
    __privateSet(this, _manifest, manifest);
    const eventMap = /* @__PURE__ */ new Map();
    for (const module of manifest.modules || []) {
      for (const declaration of module.declarations || []) {
        if (declaration.customElement && declaration.tagName) {
          const tagName = declaration.tagName;
          const events = declaration.events || [];
          if (events.length > 0) {
            const eventNames = new Set(events.map((e6) => e6.name));
            eventMap.set(tagName, {
              eventNames,
              events
            });
          }
        }
      }
    }
    return eventMap;
  } catch (error) {
    console.warn("[cem-serve-chrome] Error loading manifest for event discovery:", error);
    return /* @__PURE__ */ new Map();
  }
};
setupEventCapture_fn = async function() {
  __privateSet(this, _elementEventMap, await __privateMethod(this, _CemServeChrome_instances, discoverElementEvents_fn).call(this));
  if (__privateGet(this, _elementEventMap).size === 0) return;
  __privateMethod(this, _CemServeChrome_instances, attachEventListeners_fn).call(this);
  __privateMethod(this, _CemServeChrome_instances, updateEventFilters_fn).call(this);
  __privateMethod(this, _CemServeChrome_instances, observeDemoMutations_fn).call(this);
};
attachEventListeners_fn = function() {
  const demo = this.demo;
  if (!demo) return;
  const root = demo.shadowRoot ?? demo;
  for (const [tagName, eventInfo] of __privateGet(this, _elementEventMap)) {
    const elements = root.querySelectorAll(tagName);
    for (const element of elements) {
      for (const eventName of eventInfo.eventNames) {
        element.addEventListener(eventName, __privateGet(this, _handleElementEvent), { capture: true });
      }
      element.dataset.cemEventsAttached = "true";
      __privateGet(this, _discoveredElements).add(tagName);
    }
  }
};
observeDemoMutations_fn = function() {
  const demo = this.demo;
  if (!demo) return;
  const root = demo.shadowRoot ?? demo;
  __privateGet(this, _observer2).observe(root, {
    childList: true,
    subtree: true
  });
};
_handleElementEvent = new WeakMap();
getEventDocumentation_fn = function(manifestEvent) {
  if (!manifestEvent) {
    return { summary: null, description: null };
  }
  let summary = manifestEvent.summary || null;
  let description = manifestEvent.description || null;
  if (manifestEvent.type?.text && __privateGet(this, _manifest)) {
    const typeName = manifestEvent.type.text;
    const typeDeclaration = __privateMethod(this, _CemServeChrome_instances, findTypeDeclaration_fn).call(this, typeName);
    if (typeDeclaration) {
      if (!summary && typeDeclaration.summary) {
        summary = typeDeclaration.summary;
      } else if (typeDeclaration.summary && typeDeclaration.summary !== summary) {
        summary = summary ? `${summary}

From ${typeName}: ${typeDeclaration.summary}` : typeDeclaration.summary;
      }
      if (!description && typeDeclaration.description) {
        description = typeDeclaration.description;
      } else if (typeDeclaration.description && typeDeclaration.description !== description) {
        description = description ? `${description}

${typeDeclaration.description}` : typeDeclaration.description;
      }
    }
  }
  return { summary, description };
};
findTypeDeclaration_fn = function(typeName) {
  if (!__privateGet(this, _manifest)) return null;
  for (const module of __privateGet(this, _manifest).modules || []) {
    for (const declaration of module.declarations || []) {
      if (declaration.name === typeName && (declaration.kind === "class" || declaration.kind === "interface")) {
        return declaration;
      }
    }
  }
  return null;
};
captureEvent_fn = function(event, target, tagName, eventInfo) {
  const manifestEvent = eventInfo.events.find((e6) => e6.name === event.type);
  const eventDocs = __privateMethod(this, _CemServeChrome_instances, getEventDocumentation_fn).call(this, manifestEvent);
  const customProperties = __privateMethod(this, _CemServeChrome_instances, extractEventProperties_fn).call(this, event);
  const eventRecord = {
    id: `${Date.now()}-${Math.random()}`,
    timestamp: /* @__PURE__ */ new Date(),
    eventName: event.type,
    tagName,
    elementId: target.id || null,
    elementClass: target.className || null,
    customProperties,
    manifestType: manifestEvent?.type?.text || null,
    summary: eventDocs.summary,
    description: eventDocs.description,
    bubbles: event.bubbles,
    composed: event.composed,
    cancelable: event.cancelable,
    defaultPrevented: event.defaultPrevented
  };
  __privateGet(this, _capturedEvents).push(eventRecord);
  if (__privateGet(this, _capturedEvents).length > __privateGet(this, _maxCapturedEvents)) {
    __privateGet(this, _capturedEvents).shift();
  }
  __privateMethod(this, _CemServeChrome_instances, renderEvent_fn).call(this, eventRecord);
};
extractEventProperties_fn = function(event) {
  const properties = {};
  const eventPrototypeKeys = new Set(Object.getOwnPropertyNames(Event.prototype));
  const serializeValue = (value) => {
    try {
      return JSON.parse(JSON.stringify(value));
    } catch (e6) {
      try {
        return String(value);
      } catch (stringErr) {
        return "[Not serializable]";
      }
    }
  };
  if (event instanceof CustomEvent && event.detail !== void 0) {
    properties.detail = serializeValue(event.detail);
  }
  for (const key of Object.getOwnPropertyNames(event)) {
    if (!eventPrototypeKeys.has(key) && !key.startsWith("_") && !properties.hasOwnProperty(key)) {
      properties[key] = serializeValue(event[key]);
    }
  }
  return properties;
};
renderEvent_fn = function(eventRecord) {
  if (!__privateGet(this, _eventList)) return;
  const fragment = __privateGet(_CemServeChrome, _eventEntryTemplate).content.cloneNode(true);
  const time = eventRecord.timestamp.toLocaleTimeString();
  const container = fragment.querySelector('[data-field="container"]');
  container.dataset.eventId = eventRecord.id;
  container.dataset.eventType = eventRecord.eventName;
  container.dataset.elementType = eventRecord.tagName;
  const textMatch = __privateMethod(this, _CemServeChrome_instances, eventMatchesTextFilter_fn).call(this, eventRecord);
  const typeMatch = __privateGet(this, _eventTypeFilters).size === 0 || __privateGet(this, _eventTypeFilters).has(eventRecord.eventName);
  const elementMatch = __privateGet(this, _elementFilters).size === 0 || __privateGet(this, _elementFilters).has(eventRecord.tagName);
  if (!(textMatch && typeMatch && elementMatch)) {
    container.setAttribute("hidden", "");
  }
  const label = fragment.querySelector('[data-field="label"]');
  label.textContent = eventRecord.eventName;
  label.setAttribute("status", "info");
  const timeEl = fragment.querySelector('[data-field="time"]');
  timeEl.setAttribute("datetime", eventRecord.timestamp.toISOString());
  timeEl.textContent = time;
  const elementEl = fragment.querySelector('[data-field="element"]');
  let elementText = `<${eventRecord.tagName}>`;
  if (eventRecord.elementId) {
    elementText += `#${eventRecord.elementId}`;
  }
  elementEl.textContent = elementText;
  __privateGet(this, _eventList).append(fragment);
  if (!__privateGet(this, _selectedEventId)) {
    __privateMethod(this, _CemServeChrome_instances, selectEvent_fn).call(this, eventRecord.id);
  }
  if (__privateGet(this, _drawerOpen) && __privateMethod(this, _CemServeChrome_instances, isEventsTabActive_fn).call(this)) {
    __privateMethod(this, _CemServeChrome_instances, scrollEventsToBottom_fn).call(this);
  }
};
selectEvent_fn = function(eventId) {
  const eventRecord = __privateMethod(this, _CemServeChrome_instances, getEventRecordById_fn).call(this, eventId);
  if (!eventRecord) return;
  __privateSet(this, _selectedEventId, eventId);
  const allItems = __privateGet(this, _eventList)?.querySelectorAll(".event-list-item");
  allItems?.forEach((item) => {
    if (item.dataset.eventId === eventId) {
      item.classList.add("selected");
      item.setAttribute("aria-selected", "true");
    } else {
      item.classList.remove("selected");
      item.setAttribute("aria-selected", "false");
    }
  });
  if (__privateGet(this, _eventDetailHeader)) {
    __privateGet(this, _eventDetailHeader).innerHTML = "";
    const headerContent = document.createElement("div");
    headerContent.className = "event-detail-header-content";
    const eventName = document.createElement("h3");
    eventName.textContent = eventRecord.eventName;
    eventName.className = "event-detail-name";
    headerContent.appendChild(eventName);
    if (eventRecord.summary) {
      const summary = document.createElement("p");
      summary.textContent = eventRecord.summary;
      summary.className = "event-detail-summary";
      headerContent.appendChild(summary);
    }
    if (eventRecord.description) {
      const description = document.createElement("p");
      description.textContent = eventRecord.description;
      description.className = "event-detail-description";
      headerContent.appendChild(description);
    }
    const meta = document.createElement("div");
    meta.className = "event-detail-meta";
    const timeEl = document.createElement("time");
    timeEl.setAttribute("datetime", eventRecord.timestamp.toISOString());
    timeEl.textContent = eventRecord.timestamp.toLocaleTimeString();
    timeEl.className = "event-detail-time";
    const element = document.createElement("span");
    let elementText = `<${eventRecord.tagName}>`;
    if (eventRecord.elementId) {
      elementText += `#${eventRecord.elementId}`;
    }
    element.textContent = elementText;
    element.className = "event-detail-element";
    meta.appendChild(timeEl);
    meta.appendChild(element);
    headerContent.appendChild(meta);
    __privateGet(this, _eventDetailHeader).appendChild(headerContent);
  }
  if (__privateGet(this, _eventDetailBody)) {
    __privateGet(this, _eventDetailBody).innerHTML = "";
    const propertiesHeading = document.createElement("h4");
    propertiesHeading.textContent = "Properties";
    propertiesHeading.className = "event-detail-properties-heading";
    const propertiesContainer = document.createElement("div");
    propertiesContainer.className = "event-detail-properties";
    const eventProperties = __privateMethod(this, _CemServeChrome_instances, buildPropertiesForDisplay_fn).call(this, eventRecord);
    if (Object.keys(eventProperties).length > 0) {
      propertiesContainer.appendChild(__privateMethod(this, _CemServeChrome_instances, buildPropertyTree_fn).call(this, eventProperties));
    } else {
      propertiesContainer.textContent = "No properties to display";
    }
    __privateGet(this, _eventDetailBody).appendChild(propertiesHeading);
    __privateGet(this, _eventDetailBody).appendChild(propertiesContainer);
  }
};
buildPropertiesForDisplay_fn = function(eventRecord) {
  const properties = {};
  if (eventRecord.customProperties) {
    Object.assign(properties, eventRecord.customProperties);
  }
  properties.bubbles = eventRecord.bubbles;
  properties.cancelable = eventRecord.cancelable;
  properties.defaultPrevented = eventRecord.defaultPrevented;
  properties.composed = eventRecord.composed;
  if (eventRecord.manifestType) {
    properties.type = eventRecord.manifestType;
  }
  return properties;
};
buildPropertyTree_fn = function(obj, depth = 0) {
  const ul = document.createElement("ul");
  ul.className = "event-property-tree";
  if (depth > 0) {
    ul.classList.add("nested");
  }
  for (const [key, value] of Object.entries(obj)) {
    const li = document.createElement("li");
    li.className = "property-item";
    const keySpan = document.createElement("span");
    keySpan.className = "property-key";
    keySpan.textContent = key;
    const colonSpan = document.createElement("span");
    colonSpan.className = "property-colon";
    colonSpan.textContent = ": ";
    li.appendChild(keySpan);
    li.appendChild(colonSpan);
    if (value === null || value === void 0) {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value null";
      valueSpan.textContent = String(value);
      li.appendChild(valueSpan);
    } else if (typeof value === "boolean") {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value boolean";
      valueSpan.textContent = String(value);
      li.appendChild(valueSpan);
    } else if (typeof value === "number") {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value number";
      valueSpan.textContent = String(value);
      li.appendChild(valueSpan);
    } else if (typeof value === "string") {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value string";
      valueSpan.textContent = `"${value}"`;
      li.appendChild(valueSpan);
    } else if (Array.isArray(value)) {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value array";
      valueSpan.textContent = `Array(${value.length})`;
      li.appendChild(valueSpan);
      if (value.length > 0 && depth < 3) {
        const nestedObj = {};
        value.forEach((item, index) => {
          nestedObj[index] = item;
        });
        li.appendChild(__privateMethod(this, _CemServeChrome_instances, buildPropertyTree_fn).call(this, nestedObj, depth + 1));
      }
    } else if (typeof value === "object") {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value object";
      const keys = Object.keys(value);
      valueSpan.textContent = keys.length > 0 ? `Object` : `{}`;
      li.appendChild(valueSpan);
      if (keys.length > 0 && depth < 3) {
        li.appendChild(__privateMethod(this, _CemServeChrome_instances, buildPropertyTree_fn).call(this, value, depth + 1));
      }
    } else {
      const valueSpan = document.createElement("span");
      valueSpan.className = "property-value";
      valueSpan.textContent = String(value);
      li.appendChild(valueSpan);
    }
    ul.appendChild(li);
  }
  return ul;
};
scrollEventsToBottom_fn = function() {
  if (!__privateGet(this, _eventList)) return;
  requestAnimationFrame(() => {
    const lastEvent = __privateGet(this, _eventList).lastElementChild;
    if (lastEvent) {
      lastEvent.scrollIntoView({ behavior: "auto", block: "end" });
    }
  });
};
isEventsTabActive_fn = function() {
  const tabs = this.shadowRoot?.querySelector("cem-pf-v6-tabs");
  if (!tabs) return false;
  const selectedIndex = parseInt(tabs.getAttribute("selected") || "0", 10);
  return selectedIndex === 3;
};
filterEvents_fn = function(query) {
  __privateSet(this, _eventsFilterValue, query.toLowerCase());
  if (!__privateGet(this, _eventList)) return;
  for (const entry of __privateGet(this, _eventList).children) {
    const eventRecord = __privateMethod(this, _CemServeChrome_instances, getEventRecordById_fn).call(this, entry.dataset.eventId);
    if (!eventRecord) continue;
    const textMatch = __privateMethod(this, _CemServeChrome_instances, eventMatchesTextFilter_fn).call(this, eventRecord);
    const typeMatch = __privateGet(this, _eventTypeFilters).size === 0 || __privateGet(this, _eventTypeFilters).has(eventRecord.eventName);
    const elementMatch = __privateGet(this, _elementFilters).size === 0 || __privateGet(this, _elementFilters).has(eventRecord.tagName);
    entry.hidden = !(textMatch && typeMatch && elementMatch);
  }
};
eventMatchesTextFilter_fn = function(eventRecord) {
  if (!__privateGet(this, _eventsFilterValue)) return true;
  const searchText = [
    eventRecord.tagName,
    eventRecord.eventName,
    eventRecord.elementId || "",
    JSON.stringify(eventRecord.customProperties || {})
  ].join(" ").toLowerCase();
  return searchText.includes(__privateGet(this, _eventsFilterValue));
};
getEventRecordById_fn = function(id) {
  return __privateGet(this, _capturedEvents).find((e6) => e6.id === id);
};
updateEventFilters_fn = function() {
  const savedPreferences = __privateMethod(this, _CemServeChrome_instances, loadEventFiltersFromStorage_fn).call(this);
  const eventTypeFilter = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "event-type-filter");
  if (eventTypeFilter && __privateGet(this, _elementEventMap)) {
    let menu = eventTypeFilter.querySelector("cem-pf-v6-menu");
    if (!menu) {
      menu = document.createElement("cem-pf-v6-menu");
      eventTypeFilter.appendChild(menu);
    }
    const existingItems = menu.querySelectorAll("cem-pf-v6-menu-item");
    existingItems.forEach((item) => item.remove());
    const allEventTypes = /* @__PURE__ */ new Set();
    for (const [tagName, eventInfo] of __privateGet(this, _elementEventMap)) {
      if (__privateGet(this, _discoveredElements).has(tagName)) {
        for (const eventName of eventInfo.eventNames) {
          allEventTypes.add(eventName);
        }
      }
    }
    if (savedPreferences.eventTypes) {
      __privateSet(this, _eventTypeFilters, savedPreferences.eventTypes.intersection(allEventTypes));
    } else {
      __privateSet(this, _eventTypeFilters, new Set(allEventTypes));
    }
    for (const eventName of allEventTypes) {
      const item = document.createElement("cem-pf-v6-menu-item");
      item.setAttribute("variant", "checkbox");
      item.setAttribute("value", eventName);
      if (__privateGet(this, _eventTypeFilters).has(eventName)) {
        item.setAttribute("checked", "");
      }
      item.textContent = eventName;
      menu.appendChild(item);
    }
  }
  const elementFilter = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "element-filter");
  if (elementFilter && __privateGet(this, _elementEventMap)) {
    let menu = elementFilter.querySelector("cem-pf-v6-menu");
    if (!menu) {
      menu = document.createElement("cem-pf-v6-menu");
      elementFilter.appendChild(menu);
    }
    const existingItems = menu.querySelectorAll("cem-pf-v6-menu-item");
    existingItems.forEach((item) => item.remove());
    const allElements = /* @__PURE__ */ new Set();
    for (const tagName of __privateGet(this, _elementEventMap).keys()) {
      if (__privateGet(this, _discoveredElements).has(tagName)) {
        allElements.add(tagName);
      }
    }
    if (savedPreferences.elements) {
      __privateSet(this, _elementFilters, savedPreferences.elements.intersection(allElements));
    } else {
      __privateSet(this, _elementFilters, new Set(allElements));
    }
    for (const tagName of allElements) {
      const item = document.createElement("cem-pf-v6-menu-item");
      item.setAttribute("variant", "checkbox");
      item.setAttribute("value", tagName);
      if (__privateGet(this, _elementFilters).has(tagName)) {
        item.setAttribute("checked", "");
      }
      item.textContent = `<${tagName}>`;
      menu.appendChild(item);
    }
  }
};
_handleEventTypeFilterChange = new WeakMap();
_handleElementFilterChange = new WeakMap();
loadEventFiltersFromStorage_fn = function() {
  const preferences = {
    eventTypes: null,
    elements: null
  };
  try {
    const savedEventTypes = localStorage.getItem("cem-serve-event-type-filters");
    if (savedEventTypes) {
      preferences.eventTypes = new Set(JSON.parse(savedEventTypes));
    }
    const savedElements = localStorage.getItem("cem-serve-element-filters");
    if (savedElements) {
      preferences.elements = new Set(JSON.parse(savedElements));
    }
  } catch (e6) {
    console.debug("[cem-serve-chrome] localStorage unavailable for event filters");
  }
  return preferences;
};
saveEventFilters_fn = function() {
  try {
    localStorage.setItem(
      "cem-serve-event-type-filters",
      JSON.stringify([...__privateGet(this, _eventTypeFilters)])
    );
    localStorage.setItem(
      "cem-serve-element-filters",
      JSON.stringify([...__privateGet(this, _elementFilters)])
    );
  } catch (e6) {
  }
};
clearEvents_fn = function() {
  __privateSet(this, _capturedEvents, []);
  __privateSet(this, _selectedEventId, null);
  if (__privateGet(this, _eventList)) {
    __privateGet(this, _eventList).replaceChildren();
  }
  if (__privateGet(this, _eventDetailHeader)) {
    __privateGet(this, _eventDetailHeader).innerHTML = "";
  }
  if (__privateGet(this, _eventDetailBody)) {
    __privateGet(this, _eventDetailBody).innerHTML = "";
  }
};
copyEvents_fn = async function() {
  if (!__privateGet(this, _eventList)) return;
  const visibleEvents = Array.from(__privateGet(this, _eventList).children).filter((entry) => !entry.hidden).map((entry) => {
    const id = entry.dataset.eventId;
    return __privateMethod(this, _CemServeChrome_instances, getEventRecordById_fn).call(this, id);
  }).filter((event) => !!event).map((event) => {
    const time = event.timestamp.toLocaleTimeString();
    const element = event.elementId ? `#${event.elementId}` : event.tagName;
    const props = event.customProperties && Object.keys(event.customProperties).length > 0 ? ` - ${JSON.stringify(event.customProperties)}` : "";
    return `[${time}] <${event.tagName}> ${element} \u2192 ${event.eventName}${props}`;
  }).join("\n");
  if (!visibleEvents) return;
  try {
    await navigator.clipboard.writeText(visibleEvents);
    const btn = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "copy-events");
    if (btn) {
      const textNode = Array.from(btn.childNodes).find(
        (n6) => n6.nodeType === Node.TEXT_NODE && (n6.textContent?.trim().length ?? 0) > 0
      );
      if (textNode) {
        const original = textNode.textContent;
        textNode.textContent = "Copied!";
        if (__privateGet(this, _copyEventsFeedbackTimeout)) {
          clearTimeout(__privateGet(this, _copyEventsFeedbackTimeout));
        }
        __privateSet(this, _copyEventsFeedbackTimeout, setTimeout(() => {
          if (this.isConnected && textNode.parentNode) {
            textNode.textContent = original;
          }
          __privateSet(this, _copyEventsFeedbackTimeout, null);
        }, 2e3));
      }
    }
  } catch (err) {
    console.error("[cem-serve-chrome] Failed to copy events:", err);
  }
};
setupEventListeners_fn = function() {
  __privateSet(this, _eventList, __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "event-list"));
  __privateSet(this, _eventDetailHeader, __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "event-detail-header"));
  __privateSet(this, _eventDetailBody, __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "event-detail-body"));
  if (__privateGet(this, _eventList)) {
    __privateGet(this, _eventList).addEventListener("click", (e6) => {
      const listItem = e6.target.closest(".event-list-item");
      if (listItem) {
        const eventId = listItem.dataset.eventId;
        if (eventId) {
          __privateMethod(this, _CemServeChrome_instances, selectEvent_fn).call(this, eventId);
        }
      }
    });
  }
  const eventsFilter = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "events-filter");
  if (eventsFilter) {
    eventsFilter.addEventListener("input", (e6) => {
      const { value = "" } = e6.target;
      clearTimeout(__privateGet(this, _eventsFilterDebounceTimer));
      __privateSet(this, _eventsFilterDebounceTimer, setTimeout(() => {
        __privateMethod(this, _CemServeChrome_instances, filterEvents_fn).call(this, value);
      }, 300));
    });
  }
  const eventTypeFilter = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "event-type-filter");
  if (eventTypeFilter) {
    eventTypeFilter.addEventListener("select", __privateGet(this, _handleEventTypeFilterChange));
  }
  const elementFilter = __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "element-filter");
  if (elementFilter) {
    elementFilter.addEventListener("select", __privateGet(this, _handleElementFilterChange));
  }
  __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "clear-events")?.addEventListener("click", () => {
    __privateMethod(this, _CemServeChrome_instances, clearEvents_fn).call(this);
  });
  __privateMethod(this, _CemServeChrome_instances, $_fn2).call(this, "copy-events")?.addEventListener("click", () => {
    __privateMethod(this, _CemServeChrome_instances, copyEvents_fn).call(this);
  });
};
__decorateElement(_init54, 4, "primaryTagName", _primaryTagName_dec, _CemServeChrome, _primaryTagName);
__decorateElement(_init54, 4, "demoTitle", _demoTitle_dec, _CemServeChrome, _demoTitle);
__decorateElement(_init54, 4, "packageName", _packageName_dec, _CemServeChrome, _packageName);
__decorateElement(_init54, 4, "canonicalURL", _canonicalURL_dec, _CemServeChrome, _canonicalURL);
__decorateElement(_init54, 4, "sourceURL", _sourceURL_dec, _CemServeChrome, _sourceURL);
__decorateElement(_init54, 4, "knobs", _knobs_dec, _CemServeChrome, _knobs);
__decorateElement(_init54, 4, "drawer", _drawer_dec, _CemServeChrome, _drawer);
__decorateElement(_init54, 4, "drawerHeight", _drawerHeight_dec2, _CemServeChrome, _drawerHeight2);
__decorateElement(_init54, 4, "tabsSelected", _tabsSelected_dec, _CemServeChrome, _tabsSelected);
__decorateElement(_init54, 4, "sidebar", _sidebar_dec, _CemServeChrome, _sidebar);
__decorateElement(_init54, 4, "hasDescription", _hasDescription_dec, _CemServeChrome, _hasDescription2);
_CemServeChrome = __decorateElement(_init54, 0, "CemServeChrome", _CemServeChrome_decorators, _CemServeChrome);
__publicField(_CemServeChrome, "styles", cem_serve_chrome_default);
// Static templates for DOM cloning (logs, events, debug info)
__privateAdd(_CemServeChrome, _demoInfoTemplate, (() => {
  const t6 = document.createElement("template");
  t6.innerHTML = `
      <h3>Demo Information</h3>
      <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="cem-pf-v6-c-description-list__group">
          <dt class="cem-pf-v6-c-description-list__term">Tag Name</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="tagName"></dd>
        </div>
        <div class="cem-pf-v6-c-description-list__group" data-field-group="description">
          <dt class="cem-pf-v6-c-description-list__term">Description</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="description"></dd>
        </div>
        <div class="cem-pf-v6-c-description-list__group" data-field-group="filepath">
          <dt class="cem-pf-v6-c-description-list__term">File Path</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="filepath"></dd>
        </div>
        <div class="cem-pf-v6-c-description-list__group">
          <dt class="cem-pf-v6-c-description-list__term">Canonical URL</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="canonicalURL"></dd>
        </div>
        <div class="cem-pf-v6-c-description-list__group">
          <dt class="cem-pf-v6-c-description-list__term">Local Route</dt>
          <dd class="cem-pf-v6-c-description-list__description" data-field="localRoute"></dd>
        </div>
      </dl>`;
  return t6;
})());
__privateAdd(_CemServeChrome, _demoGroupTemplate, (() => {
  const t6 = document.createElement("template");
  t6.innerHTML = `
      <div class="cem-pf-v6-c-description-list__group">
        <dt class="cem-pf-v6-c-description-list__term" data-field="tagName"></dt>
        <dd class="cem-pf-v6-c-description-list__description">
          <span data-field="description"></span><br>
          <small data-field-group="filepath">File: <span data-field="filepath"></span></small>
          <small>Canonical: <span data-field="canonicalURL"></span></small><br>
          <small>Local: <span data-field="localRoute"></span></small>
        </dd>
      </div>`;
  return t6;
})());
__privateAdd(_CemServeChrome, _demoListTemplate, (() => {
  const t6 = document.createElement("template");
  t6.innerHTML = `
      <cem-pf-v6-expandable-section id="debug-demos-section"
                                toggle-text="Show Demos Info">
        <dl class="cem-pf-v6-c-description-list pf-m-horizontal pf-m-compact" data-container="groups"></dl>
      </cem-pf-v6-expandable-section>`;
  return t6;
})());
__privateAdd(_CemServeChrome, _logEntryTemplate, (() => {
  const t6 = document.createElement("template");
  t6.innerHTML = `
      <div class="log-entry" data-field="container">
        <cem-pf-v6-label compact data-field="label"></cem-pf-v6-label>
        <time class="log-time" data-field="time"></time>
        <span class="log-message" data-field="message"></span>
      </div>`;
  return t6;
})());
__privateAdd(_CemServeChrome, _eventEntryTemplate, (() => {
  const t6 = document.createElement("template");
  t6.innerHTML = `
      <button class="event-list-item" data-field="container">
        <cem-pf-v6-label compact data-field="label"></cem-pf-v6-label>
        <time class="event-time" data-field="time"></time>
        <span class="event-element" data-field="element"></span>
      </button>`;
  return t6;
})());
__runInitializers(_init54, 1, _CemServeChrome);
var CemServeChrome = _CemServeChrome;
/*! Bundled license information:

@lit/reactive-element/node/css-tag.js:
  (**
   * @license
   * Copyright 2019 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/node/lit-html.js:
lit-element/lit-element.js:
@lit/reactive-element/node/decorators/custom-element.js:
@lit/reactive-element/node/decorators/property.js:
@lit/reactive-element/node/decorators/state.js:
lit-html/node/directive.js:
lit-html/node/directives/unsafe-html.js:
  (**
   * @license
   * Copyright 2017 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/node/is-server.js:
  (**
   * @license
   * Copyright 2022 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/node/directives/if-defined.js:
lit-html/node/directives/class-map.js:
lit-html/node/directives/style-map.js:
  (**
   * @license
   * Copyright 2018 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)

lit-html/node/directive-helpers.js:
lit-html/node/directives/live.js:
  (**
   * @license
   * Copyright 2020 Google LLC
   * SPDX-License-Identifier: BSD-3-Clause
   *)
*/
