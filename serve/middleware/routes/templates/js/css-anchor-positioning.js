/**
 * @oddbird/css-anchor-positioning v0.7.0
 * https://github.com/oddbird/css-anchor-positioning
 * 
 * Copyright (c) 2022–2024 OddBird LLC
 * Licensed under the BSD-3-Clause License
 * 
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 * 
 * 1. Redistributions of source code must retain the above copyright notice, this
 *    list of conditions and the following disclaimer.
 * 
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 * 
 * 3. Neither the name of the copyright holder nor the names of its contributors
 *    may be used to endorse or promote products derived from this software without
 *    specific prior written permission.
 * 
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
 * ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
 * ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

var yo = Object.defineProperty, bo = Object.defineProperties;
var xo = Object.getOwnPropertyDescriptors;
var $n = Object.getOwnPropertySymbols;
var Co = Object.prototype.hasOwnProperty, wo = Object.prototype.propertyIsEnumerable;
var Ln = (t, e, n) => e in t ? yo(t, e, { enumerable: !0, configurable: !0, writable: !0, value: n }) : t[e] = n, U = (t, e) => {
  for (var n in e || (e = {}))
    Co.call(e, n) && Ln(t, n, e[n]);
  if ($n)
    for (var n of $n(e))
      wo.call(e, n) && Ln(t, n, e[n]);
  return t;
}, q = (t, e) => bo(t, xo(e));
var To = (t, e) => () => (e || t((e = { exports: {} }).exports, e), e.exports);
var R = (t, e, n) => new Promise((s, r) => {
  var o = (u) => {
    try {
      l(n.next(u));
    } catch (i) {
      r(i);
    }
  }, a = (u) => {
    try {
      l(n.throw(u));
    } catch (i) {
      r(i);
    }
  }, l = (u) => u.done ? s(u.value) : Promise.resolve(u.value).then(o, a);
  l((n = n.apply(t, e)).next());
});
var rd = To((fe) => {
  const sn = Math.min, Gt = Math.max, Ae = Math.round, de = Math.floor, St = (t) => ({
    x: t,
    y: t
  });
  function Ao(t, e) {
    return typeof t == "function" ? t(e) : t;
  }
  function vo(t) {
    return U({
      top: 0,
      right: 0,
      bottom: 0,
      left: 0
    }, t);
  }
  function Eo(t) {
    return typeof t != "number" ? vo(t) : {
      top: t,
      right: t,
      bottom: t,
      left: t
    };
  }
  function ve(t) {
    const {
      x: e,
      y: n,
      width: s,
      height: r
    } = t;
    return {
      width: s,
      height: r,
      top: n,
      left: e,
      right: e + s,
      bottom: n + r,
      x: e,
      y: n
    };
  }
  function $o(t, e) {
    return R(this, null, function* () {
      var n;
      e === void 0 && (e = {});
      const {
        x: s,
        y: r,
        platform: o,
        rects: a,
        elements: l,
        strategy: u
      } = t, {
        boundary: i = "clippingAncestors",
        rootBoundary: c = "viewport",
        elementContext: h = "floating",
        altBoundary: f = !1,
        padding: p = 0
      } = Ao(e, t), d = Eo(p), k = l[f ? h === "floating" ? "reference" : "floating" : h], y = ve(yield o.getClippingRect({
        element: (n = yield o.isElement == null ? void 0 : o.isElement(k)) == null || n ? k : k.contextElement || (yield o.getDocumentElement == null ? void 0 : o.getDocumentElement(l.floating)),
        boundary: i,
        rootBoundary: c,
        strategy: u
      })), x = h === "floating" ? {
        x: s,
        y: r,
        width: a.floating.width,
        height: a.floating.height
      } : a.reference, A = yield o.getOffsetParent == null ? void 0 : o.getOffsetParent(l.floating), M = (yield o.isElement == null ? void 0 : o.isElement(A)) ? (yield o.getScale == null ? void 0 : o.getScale(A)) || {
        x: 1,
        y: 1
      } : {
        x: 1,
        y: 1
      }, et = ve(o.convertOffsetParentRelativeRectToViewportRelativeRect ? yield o.convertOffsetParentRelativeRectToViewportRelativeRect({
        elements: l,
        rect: x,
        offsetParent: A,
        strategy: u
      }) : x);
      return {
        top: (y.top - et.top + d.top) / M.y,
        bottom: (et.bottom - y.bottom + d.bottom) / M.y,
        left: (y.left - et.left + d.left) / M.x,
        right: (et.right - y.right + d.right) / M.x
      };
    });
  }
  function _e() {
    return typeof window != "undefined";
  }
  function te(t) {
    return Ds(t) ? (t.nodeName || "").toLowerCase() : "#document";
  }
  function it(t) {
    var e;
    return (t == null || (e = t.ownerDocument) == null ? void 0 : e.defaultView) || window;
  }
  function bt(t) {
    var e;
    return (e = (Ds(t) ? t.ownerDocument : t.document) || window.document) == null ? void 0 : e.documentElement;
  }
  function Ds(t) {
    return _e() ? t instanceof Node || t instanceof it(t).Node : !1;
  }
  function ct(t) {
    return _e() ? t instanceof Element || t instanceof it(t).Element : !1;
  }
  function yt(t) {
    return _e() ? t instanceof HTMLElement || t instanceof it(t).HTMLElement : !1;
  }
  function On(t) {
    return !_e() || typeof ShadowRoot == "undefined" ? !1 : t instanceof ShadowRoot || t instanceof it(t).ShadowRoot;
  }
  const Lo = /* @__PURE__ */ new Set(["inline", "contents"]);
  function he(t) {
    const {
      overflow: e,
      overflowX: n,
      overflowY: s,
      display: r
    } = ut(t);
    return /auto|scroll|overlay|hidden|clip/.test(e + s + n) && !Lo.has(r);
  }
  const Oo = /* @__PURE__ */ new Set(["table", "td", "th"]);
  function Po(t) {
    return Oo.has(te(t));
  }
  const Ro = [":popover-open", ":modal"];
  function Ie(t) {
    return Ro.some((e) => {
      try {
        return t.matches(e);
      } catch (n) {
        return !1;
      }
    });
  }
  const _o = ["transform", "translate", "scale", "rotate", "perspective"], Io = ["transform", "translate", "scale", "rotate", "perspective", "filter"], No = ["paint", "layout", "strict", "content"];
  function Sn(t) {
    const e = yn(), n = ct(t) ? ut(t) : t;
    return _o.some((s) => n[s] ? n[s] !== "none" : !1) || (n.containerType ? n.containerType !== "normal" : !1) || !e && (n.backdropFilter ? n.backdropFilter !== "none" : !1) || !e && (n.filter ? n.filter !== "none" : !1) || Io.some((s) => (n.willChange || "").includes(s)) || No.some((s) => (n.contain || "").includes(s));
  }
  function Do(t) {
    let e = Ot(t);
    for (; yt(e) && !Qt(e); ) {
      if (Sn(e))
        return e;
      if (Ie(e))
        return null;
      e = Ot(e);
    }
    return null;
  }
  function yn() {
    return typeof CSS == "undefined" || !CSS.supports ? !1 : CSS.supports("-webkit-backdrop-filter", "none");
  }
  const Fo = /* @__PURE__ */ new Set(["html", "body", "#document"]);
  function Qt(t) {
    return Fo.has(te(t));
  }
  function ut(t) {
    return it(t).getComputedStyle(t);
  }
  function Ne(t) {
    return ct(t) ? {
      scrollLeft: t.scrollLeft,
      scrollTop: t.scrollTop
    } : {
      scrollLeft: t.scrollX,
      scrollTop: t.scrollY
    };
  }
  function Ot(t) {
    if (te(t) === "html")
      return t;
    const e = (
      // Step into the shadow DOM of the parent of a slotted node.
      t.assignedSlot || // DOM Element detected.
      t.parentNode || // ShadowRoot detected.
      On(t) && t.host || // Fallback.
      bt(t)
    );
    return On(e) ? e.host : e;
  }
  function Fs(t) {
    const e = Ot(t);
    return Qt(e) ? t.ownerDocument ? t.ownerDocument.body : t.body : yt(e) && he(e) ? e : Fs(e);
  }
  function le(t, e, n) {
    var s;
    e === void 0 && (e = []), n === void 0 && (n = !0);
    const r = Fs(t), o = r === ((s = t.ownerDocument) == null ? void 0 : s.body), a = it(r);
    if (o) {
      const l = rn(a);
      return e.concat(a, a.visualViewport || [], he(r) ? r : [], l && n ? le(l) : []);
    }
    return e.concat(r, le(r, [], n));
  }
  function rn(t) {
    return t.parent && Object.getPrototypeOf(t.parent) ? t.frameElement : null;
  }
  function Ms(t) {
    const e = ut(t);
    let n = parseFloat(e.width) || 0, s = parseFloat(e.height) || 0;
    const r = yt(t), o = r ? t.offsetWidth : n, a = r ? t.offsetHeight : s, l = Ae(n) !== o || Ae(s) !== a;
    return l && (n = o, s = a), {
      width: n,
      height: s,
      $: l
    };
  }
  function bn(t) {
    return ct(t) ? t : t.contextElement;
  }
  function qt(t) {
    const e = bn(t);
    if (!yt(e))
      return St(1);
    const n = e.getBoundingClientRect(), {
      width: s,
      height: r,
      $: o
    } = Ms(e);
    let a = (o ? Ae(n.width) : n.width) / s, l = (o ? Ae(n.height) : n.height) / r;
    return (!a || !Number.isFinite(a)) && (a = 1), (!l || !Number.isFinite(l)) && (l = 1), {
      x: a,
      y: l
    };
  }
  const Mo = /* @__PURE__ */ St(0);
  function js(t) {
    const e = it(t);
    return !yn() || !e.visualViewport ? Mo : {
      x: e.visualViewport.offsetLeft,
      y: e.visualViewport.offsetTop
    };
  }
  function jo(t, e, n) {
    return e === void 0 && (e = !1), !n || e && n !== it(t) ? !1 : e;
  }
  function Dt(t, e, n, s) {
    e === void 0 && (e = !1), n === void 0 && (n = !1);
    const r = t.getBoundingClientRect(), o = bn(t);
    let a = St(1);
    e && (s ? ct(s) && (a = qt(s)) : a = qt(t));
    const l = jo(o, n, s) ? js(o) : St(0);
    let u = (r.left + l.x) / a.x, i = (r.top + l.y) / a.y, c = r.width / a.x, h = r.height / a.y;
    if (o) {
      const f = it(o), p = s && ct(s) ? it(s) : s;
      let d = f, m = rn(d);
      for (; m && s && p !== d; ) {
        const k = qt(m), y = m.getBoundingClientRect(), x = ut(m), A = y.left + (m.clientLeft + parseFloat(x.paddingLeft)) * k.x, M = y.top + (m.clientTop + parseFloat(x.paddingTop)) * k.y;
        u *= k.x, i *= k.y, c *= k.x, h *= k.y, u += A, i += M, d = it(m), m = rn(d);
      }
    }
    return ve({
      width: c,
      height: h,
      x: u,
      y: i
    });
  }
  function De(t, e) {
    const n = Ne(t).scrollLeft;
    return e ? e.left + n : Dt(bt(t)).left + n;
  }
  function Bs(t, e) {
    const n = t.getBoundingClientRect(), s = n.left + e.scrollLeft - De(t, n), r = n.top + e.scrollTop;
    return {
      x: s,
      y: r
    };
  }
  function Bo(t) {
    let {
      elements: e,
      rect: n,
      offsetParent: s,
      strategy: r
    } = t;
    const o = r === "fixed", a = bt(s), l = e ? Ie(e.floating) : !1;
    if (s === a || l && o)
      return n;
    let u = {
      scrollLeft: 0,
      scrollTop: 0
    }, i = St(1);
    const c = St(0), h = yt(s);
    if ((h || !h && !o) && ((te(s) !== "body" || he(a)) && (u = Ne(s)), yt(s))) {
      const p = Dt(s);
      i = qt(s), c.x = p.x + s.clientLeft, c.y = p.y + s.clientTop;
    }
    const f = a && !h && !o ? Bs(a, u) : St(0);
    return {
      width: n.width * i.x,
      height: n.height * i.y,
      x: n.x * i.x - u.scrollLeft * i.x + c.x + f.x,
      y: n.y * i.y - u.scrollTop * i.y + c.y + f.y
    };
  }
  function Wo(t) {
    return Array.from(t.getClientRects());
  }
  function Uo(t) {
    const e = bt(t), n = Ne(t), s = t.ownerDocument.body, r = Gt(e.scrollWidth, e.clientWidth, s.scrollWidth, s.clientWidth), o = Gt(e.scrollHeight, e.clientHeight, s.scrollHeight, s.clientHeight);
    let a = -n.scrollLeft + De(t);
    const l = -n.scrollTop;
    return ut(s).direction === "rtl" && (a += Gt(e.clientWidth, s.clientWidth) - r), {
      width: r,
      height: o,
      x: a,
      y: l
    };
  }
  const Pn = 25;
  function zo(t, e) {
    const n = it(t), s = bt(t), r = n.visualViewport;
    let o = s.clientWidth, a = s.clientHeight, l = 0, u = 0;
    if (r) {
      o = r.width, a = r.height;
      const c = yn();
      (!c || c && e === "fixed") && (l = r.offsetLeft, u = r.offsetTop);
    }
    const i = De(s);
    if (i <= 0) {
      const c = s.ownerDocument, h = c.body, f = getComputedStyle(h), p = c.compatMode === "CSS1Compat" && parseFloat(f.marginLeft) + parseFloat(f.marginRight) || 0, d = Math.abs(s.clientWidth - h.clientWidth - p);
      d <= Pn && (o -= d);
    } else i <= Pn && (o += i);
    return {
      width: o,
      height: a,
      x: l,
      y: u
    };
  }
  const Vo = /* @__PURE__ */ new Set(["absolute", "fixed"]);
  function Ho(t, e) {
    const n = Dt(t, !0, e === "fixed"), s = n.top + t.clientTop, r = n.left + t.clientLeft, o = yt(t) ? qt(t) : St(1), a = t.clientWidth * o.x, l = t.clientHeight * o.y, u = r * o.x, i = s * o.y;
    return {
      width: a,
      height: l,
      x: u,
      y: i
    };
  }
  function Rn(t, e, n) {
    let s;
    if (e === "viewport")
      s = zo(t, n);
    else if (e === "document")
      s = Uo(bt(t));
    else if (ct(e))
      s = Ho(e, n);
    else {
      const r = js(t);
      s = {
        x: e.x - r.x,
        y: e.y - r.y,
        width: e.width,
        height: e.height
      };
    }
    return ve(s);
  }
  function Ws(t, e) {
    const n = Ot(t);
    return n === e || !ct(n) || Qt(n) ? !1 : ut(n).position === "fixed" || Ws(n, e);
  }
  function Go(t, e) {
    const n = e.get(t);
    if (n)
      return n;
    let s = le(t, [], !1).filter((l) => ct(l) && te(l) !== "body"), r = null;
    const o = ut(t).position === "fixed";
    let a = o ? Ot(t) : t;
    for (; ct(a) && !Qt(a); ) {
      const l = ut(a), u = Sn(a);
      !u && l.position === "fixed" && (r = null), (o ? !u && !r : !u && l.position === "static" && !!r && Vo.has(r.position) || he(a) && !u && Ws(t, a)) ? s = s.filter((c) => c !== a) : r = l, a = Ot(a);
    }
    return e.set(t, s), s;
  }
  function qo(t) {
    let {
      element: e,
      boundary: n,
      rootBoundary: s,
      strategy: r
    } = t;
    const a = [...n === "clippingAncestors" ? Ie(e) ? [] : Go(e, this._c) : [].concat(n), s], l = a[0], u = a.reduce((i, c) => {
      const h = Rn(e, c, r);
      return i.top = Gt(h.top, i.top), i.right = sn(h.right, i.right), i.bottom = sn(h.bottom, i.bottom), i.left = Gt(h.left, i.left), i;
    }, Rn(e, l, r));
    return {
      width: u.right - u.left,
      height: u.bottom - u.top,
      x: u.left,
      y: u.top
    };
  }
  function Ko(t) {
    const {
      width: e,
      height: n
    } = Ms(t);
    return {
      width: e,
      height: n
    };
  }
  function Qo(t, e, n) {
    const s = yt(e), r = bt(e), o = n === "fixed", a = Dt(t, !0, o, e);
    let l = {
      scrollLeft: 0,
      scrollTop: 0
    };
    const u = St(0);
    function i() {
      u.x = De(r);
    }
    if (s || !s && !o)
      if ((te(e) !== "body" || he(r)) && (l = Ne(e)), s) {
        const p = Dt(e, !0, o, e);
        u.x = p.x + e.clientLeft, u.y = p.y + e.clientTop;
      } else r && i();
    o && !s && r && i();
    const c = r && !s && !o ? Bs(r, l) : St(0), h = a.left + l.scrollLeft - u.x - c.x, f = a.top + l.scrollTop - u.y - c.y;
    return {
      x: h,
      y: f,
      width: a.width,
      height: a.height
    };
  }
  function Me(t) {
    return ut(t).position === "static";
  }
  function _n(t, e) {
    if (!yt(t) || ut(t).position === "fixed")
      return null;
    if (e)
      return e(t);
    let n = t.offsetParent;
    return bt(t) === n && (n = n.ownerDocument.body), n;
  }
  function Us(t, e) {
    const n = it(t);
    if (Ie(t))
      return n;
    if (!yt(t)) {
      let r = Ot(t);
      for (; r && !Qt(r); ) {
        if (ct(r) && !Me(r))
          return r;
        r = Ot(r);
      }
      return n;
    }
    let s = _n(t, e);
    for (; s && Po(s) && Me(s); )
      s = _n(s, e);
    return s && Qt(s) && Me(s) && !Sn(s) ? n : s || Do(t) || n;
  }
  const Yo = function(t) {
    return R(this, null, function* () {
      const e = this.getOffsetParent || Us, n = this.getDimensions, s = yield n(t.floating);
      return {
        reference: Qo(t.reference, yield e(t.floating), t.strategy),
        floating: {
          x: 0,
          y: 0,
          width: s.width,
          height: s.height
        }
      };
    });
  };
  function Xo(t) {
    return ut(t).direction === "rtl";
  }
  const V = {
    convertOffsetParentRelativeRectToViewportRelativeRect: Bo,
    getDocumentElement: bt,
    getClippingRect: qo,
    getOffsetParent: Us,
    getElementRects: Yo,
    getClientRects: Wo,
    getDimensions: Ko,
    getScale: qt,
    isElement: ct,
    isRTL: Xo
  };
  function zs(t, e) {
    return t.x === e.x && t.y === e.y && t.width === e.width && t.height === e.height;
  }
  function Jo(t, e) {
    let n = null, s;
    const r = bt(t);
    function o() {
      var l;
      clearTimeout(s), (l = n) == null || l.disconnect(), n = null;
    }
    function a(l, u) {
      l === void 0 && (l = !1), u === void 0 && (u = 1), o();
      const i = t.getBoundingClientRect(), {
        left: c,
        top: h,
        width: f,
        height: p
      } = i;
      if (l || e(), !f || !p)
        return;
      const d = de(h), m = de(r.clientWidth - (c + f)), k = de(r.clientHeight - (h + p)), y = de(c), A = {
        rootMargin: -d + "px " + -m + "px " + -k + "px " + -y + "px",
        threshold: Gt(0, sn(1, u)) || 1
      };
      let M = !0;
      function et(w) {
        const C = w[0].intersectionRatio;
        if (C !== u) {
          if (!M)
            return a();
          C ? a(!1, C) : s = setTimeout(() => {
            a(!1, 1e-7);
          }, 1e3);
        }
        C === 1 && !zs(i, t.getBoundingClientRect()) && a(), M = !1;
      }
      try {
        n = new IntersectionObserver(et, q(U({}, A), {
          // Handle <iframe>s
          root: r.ownerDocument
        }));
      } catch (w) {
        n = new IntersectionObserver(et, A);
      }
      n.observe(t);
    }
    return a(!0), o;
  }
  function on(t, e, n, s) {
    s === void 0 && (s = {});
    const {
      ancestorScroll: r = !0,
      ancestorResize: o = !0,
      elementResize: a = typeof ResizeObserver == "function",
      layoutShift: l = typeof IntersectionObserver == "function",
      animationFrame: u = !1
    } = s, i = bn(t), c = r || o ? [...i ? le(i) : [], ...le(e)] : [];
    c.forEach((y) => {
      r && y.addEventListener("scroll", n, {
        passive: !0
      }), o && y.addEventListener("resize", n);
    });
    const h = i && l ? Jo(i, n) : null;
    let f = -1, p = null;
    a && (p = new ResizeObserver((y) => {
      let [x] = y;
      x && x.target === i && p && (p.unobserve(e), cancelAnimationFrame(f), f = requestAnimationFrame(() => {
        var A;
        (A = p) == null || A.observe(e);
      })), n();
    }), i && !u && p.observe(i), p.observe(e));
    let d, m = u ? Dt(t) : null;
    u && k();
    function k() {
      const y = Dt(t);
      m && !zs(m, y) && n(), m = y, d = requestAnimationFrame(k);
    }
    return n(), () => {
      var y;
      c.forEach((x) => {
        r && x.removeEventListener("scroll", n), o && x.removeEventListener("resize", n);
      }), h == null || h(), (y = p) == null || y.disconnect(), p = null, u && cancelAnimationFrame(d);
    };
  }
  const Zo = $o, { hasOwnProperty: xn } = Object.prototype, ne = function() {
  };
  function In(t) {
    return typeof t == "function" ? t : ne;
  }
  function Nn(t, e) {
    return function(n, s, r) {
      n.type === e && t.call(this, n, s, r);
    };
  }
  function ta(t, e) {
    const n = e.structure, s = [];
    for (const r in n) {
      if (xn.call(n, r) === !1)
        continue;
      let o = n[r];
      const a = {
        name: r,
        type: !1,
        nullable: !1
      };
      Array.isArray(o) || (o = [o]);
      for (const l of o)
        l === null ? a.nullable = !0 : typeof l == "string" ? a.type = "node" : Array.isArray(l) && (a.type = "list");
      a.type && s.push(a);
    }
    return s.length ? {
      context: e.walkContext,
      fields: s
    } : null;
  }
  function ea(t) {
    const e = {};
    for (const n in t.node)
      if (xn.call(t.node, n)) {
        const s = t.node[n];
        if (!s.structure)
          throw new Error("Missed `structure` field in `" + n + "` node type definition");
        e[n] = ta(n, s);
      }
    return e;
  }
  function Dn(t, e) {
    const n = t.fields.slice(), s = t.context, r = typeof s == "string";
    return e && n.reverse(), function(o, a, l, u) {
      let i;
      r && (i = a[s], a[s] = o);
      for (const c of n) {
        const h = o[c.name];
        if (!c.nullable || h) {
          if (c.type === "list") {
            if (e ? h.reduceRight(u, !1) : h.reduce(u, !1))
              return !0;
          } else if (l(h))
            return !0;
        }
      }
      r && (a[s] = i);
    };
  }
  function Fn({
    StyleSheet: t,
    Atrule: e,
    Rule: n,
    Block: s,
    DeclarationList: r
  }) {
    return {
      Atrule: {
        StyleSheet: t,
        Atrule: e,
        Rule: n,
        Block: s
      },
      Rule: {
        StyleSheet: t,
        Atrule: e,
        Rule: n,
        Block: s
      },
      Declaration: {
        StyleSheet: t,
        Atrule: e,
        Rule: n,
        Block: s,
        DeclarationList: r
      }
    };
  }
  function na(t) {
    const e = ea(t), n = {}, s = {}, r = Symbol("break-walk"), o = Symbol("skip-node");
    for (const i in e)
      xn.call(e, i) && e[i] !== null && (n[i] = Dn(e[i], !1), s[i] = Dn(e[i], !0));
    const a = Fn(n), l = Fn(s), u = function(i, c) {
      function h(y, x, A) {
        const M = f.call(k, y, x, A);
        return M === r ? !0 : M === o ? !1 : !!(d.hasOwnProperty(y.type) && d[y.type](y, k, h, m) || p.call(k, y, x, A) === r);
      }
      let f = ne, p = ne, d = n, m = (y, x, A, M) => y || h(x, A, M);
      const k = {
        break: r,
        skip: o,
        root: i,
        stylesheet: null,
        atrule: null,
        atrulePrelude: null,
        rule: null,
        selector: null,
        block: null,
        declaration: null,
        function: null
      };
      if (typeof c == "function")
        f = c;
      else if (c && (f = In(c.enter), p = In(c.leave), c.reverse && (d = s), c.visit)) {
        if (a.hasOwnProperty(c.visit))
          d = c.reverse ? l[c.visit] : a[c.visit];
        else if (!e.hasOwnProperty(c.visit))
          throw new Error("Bad value `" + c.visit + "` for `visit` option (should be: " + Object.keys(e).sort().join(", ") + ")");
        f = Nn(f, c.visit), p = Nn(p, c.visit);
      }
      if (f === ne && p === ne)
        throw new Error("Neither `enter` nor `leave` walker handler is set or both aren't a function");
      h(i);
    };
    return u.break = r, u.skip = o, u.find = function(i, c) {
      let h = null;
      return u(i, function(f, p, d) {
        if (c.call(this, f, p, d))
          return h = f, r;
      }), h;
    }, u.findLast = function(i, c) {
      let h = null;
      return u(i, {
        reverse: !0,
        enter(f, p, d) {
          if (c.call(this, f, p, d))
            return h = f, r;
        }
      }), h;
    }, u.findAll = function(i, c) {
      const h = [];
      return u(i, function(f, p, d) {
        c.call(this, f, p, d) && h.push(f);
      }), h;
    }, u;
  }
  const vt = 0, g = 1, T = 2, z = 3, I = 4, Tt = 5, sa = 6, Q = 7, at = 8, L = 9, b = 10, F = 11, E = 12, W = 13, Fe = 14, nt = 15, X = 16, tt = 17, ft = 18, ee = 19, ce = 20, _ = 21, S = 22, ht = 23, Yt = 24, Y = 25, ra = 0;
  function rt(t) {
    return t >= 48 && t <= 57;
  }
  function Xt(t) {
    return rt(t) || // 0 .. 9
    t >= 65 && t <= 70 || // A .. F
    t >= 97 && t <= 102;
  }
  function Cn(t) {
    return t >= 65 && t <= 90;
  }
  function ia(t) {
    return t >= 97 && t <= 122;
  }
  function oa(t) {
    return Cn(t) || ia(t);
  }
  function aa(t) {
    return t >= 128;
  }
  function Ee(t) {
    return oa(t) || aa(t) || t === 95;
  }
  function Vs(t) {
    return Ee(t) || rt(t) || t === 45;
  }
  function la(t) {
    return t >= 0 && t <= 8 || t === 11 || t >= 14 && t <= 31 || t === 127;
  }
  function $e(t) {
    return t === 10 || t === 13 || t === 12;
  }
  function Ft(t) {
    return $e(t) || t === 32 || t === 9;
  }
  function kt(t, e) {
    return !(t !== 92 || $e(e) || e === ra);
  }
  function je(t, e, n) {
    return t === 45 ? Ee(e) || e === 45 || kt(e, n) : Ee(t) ? !0 : t === 92 ? kt(t, e) : !1;
  }
  function Be(t, e, n) {
    return t === 43 || t === 45 ? rt(e) ? 2 : e === 46 && rt(n) ? 3 : 0 : t === 46 ? rt(e) ? 2 : 0 : rt(t) ? 1 : 0;
  }
  function Hs(t) {
    return t === 65279 || t === 65534 ? 1 : 0;
  }
  const an = new Array(128), ca = 128, be = 130, Gs = 131, wn = 132, qs = 133;
  for (let t = 0; t < an.length; t++)
    an[t] = Ft(t) && be || rt(t) && Gs || Ee(t) && wn || la(t) && qs || t || ca;
  function We(t) {
    return t < 128 ? an[t] : wn;
  }
  function Kt(t, e) {
    return e < t.length ? t.charCodeAt(e) : 0;
  }
  function ln(t, e, n) {
    return n === 13 && Kt(t, e + 1) === 10 ? 2 : 1;
  }
  function Ks(t, e, n) {
    let s = t.charCodeAt(e);
    return Cn(s) && (s = s | 32), s === n;
  }
  function Le(t, e, n, s) {
    if (n - e !== s.length || e < 0 || n > t.length)
      return !1;
    for (let r = e; r < n; r++) {
      const o = s.charCodeAt(r - e);
      let a = t.charCodeAt(r);
      if (Cn(a) && (a = a | 32), a !== o)
        return !1;
    }
    return !0;
  }
  function ua(t, e) {
    for (; e >= 0 && Ft(t.charCodeAt(e)); e--)
      ;
    return e + 1;
  }
  function ge(t, e) {
    for (; e < t.length && Ft(t.charCodeAt(e)); e++)
      ;
    return e;
  }
  function Ue(t, e) {
    for (; e < t.length && rt(t.charCodeAt(e)); e++)
      ;
    return e;
  }
  function Jt(t, e) {
    if (e += 2, Xt(Kt(t, e - 1))) {
      for (const s = Math.min(t.length, e + 5); e < s && Xt(Kt(t, e)); e++)
        ;
      const n = Kt(t, e);
      Ft(n) && (e += ln(t, e, n));
    }
    return e;
  }
  function me(t, e) {
    for (; e < t.length; e++) {
      const n = t.charCodeAt(e);
      if (!Vs(n)) {
        if (kt(n, Kt(t, e + 1))) {
          e = Jt(t, e) - 1;
          continue;
        }
        break;
      }
    }
    return e;
  }
  function Qs(t, e) {
    let n = t.charCodeAt(e);
    if ((n === 43 || n === 45) && (n = t.charCodeAt(e += 1)), rt(n) && (e = Ue(t, e + 1), n = t.charCodeAt(e)), n === 46 && rt(t.charCodeAt(e + 1)) && (e += 2, e = Ue(t, e)), Ks(
      t,
      e,
      101
      /* e */
    )) {
      let s = 0;
      n = t.charCodeAt(e + 1), (n === 45 || n === 43) && (s = 1, n = t.charCodeAt(e + 2)), rt(n) && (e = Ue(t, e + 1 + s + 1));
    }
    return e;
  }
  function ze(t, e) {
    for (; e < t.length; e++) {
      const n = t.charCodeAt(e);
      if (n === 41) {
        e++;
        break;
      }
      kt(n, Kt(t, e + 1)) && (e = Jt(t, e));
    }
    return e;
  }
  function Ys(t) {
    if (t.length === 1 && !Xt(t.charCodeAt(0)))
      return t[0];
    let e = parseInt(t, 16);
    return (e === 0 || // If this number is zero,
    e >= 55296 && e <= 57343 || // or is for a surrogate,
    e > 1114111) && (e = 65533), String.fromCodePoint(e);
  }
  const Xs = [
    "EOF-token",
    "ident-token",
    "function-token",
    "at-keyword-token",
    "hash-token",
    "string-token",
    "bad-string-token",
    "url-token",
    "bad-url-token",
    "delim-token",
    "number-token",
    "percentage-token",
    "dimension-token",
    "whitespace-token",
    "CDO-token",
    "CDC-token",
    "colon-token",
    "semicolon-token",
    "comma-token",
    "[-token",
    "]-token",
    "(-token",
    ")-token",
    "{-token",
    "}-token",
    "comment-token"
  ], ha = 16 * 1024;
  function Oe(t = null, e) {
    return t === null || t.length < e ? new Uint32Array(Math.max(e + 1024, ha)) : t;
  }
  const Mn = 10, fa = 12, jn = 13;
  function Bn(t) {
    const e = t.source, n = e.length, s = e.length > 0 ? Hs(e.charCodeAt(0)) : 0, r = Oe(t.lines, n), o = Oe(t.columns, n);
    let a = t.startLine, l = t.startColumn;
    for (let u = s; u < n; u++) {
      const i = e.charCodeAt(u);
      r[u] = a, o[u] = l++, (i === Mn || i === jn || i === fa) && (i === jn && u + 1 < n && e.charCodeAt(u + 1) === Mn && (u++, r[u] = a, o[u] = l), a++, l = 1);
    }
    r[n] = a, o[n] = l, t.lines = r, t.columns = o, t.computed = !0;
  }
  class pa {
    constructor(e, n, s, r) {
      this.setSource(e, n, s, r), this.lines = null, this.columns = null;
    }
    setSource(e = "", n = 0, s = 1, r = 1) {
      this.source = e, this.startOffset = n, this.startLine = s, this.startColumn = r, this.computed = !1;
    }
    getLocation(e, n) {
      return this.computed || Bn(this), {
        source: n,
        offset: this.startOffset + e,
        line: this.lines[e],
        column: this.columns[e]
      };
    }
    getLocationRange(e, n, s) {
      return this.computed || Bn(this), {
        source: s,
        start: {
          offset: this.startOffset + e,
          line: this.lines[e],
          column: this.columns[e]
        },
        end: {
          offset: this.startOffset + n,
          line: this.lines[n],
          column: this.columns[n]
        }
      };
    }
  }
  const dt = 16777215, gt = 24, Mt = new Uint8Array(32);
  Mt[T] = S;
  Mt[_] = S;
  Mt[ee] = ce;
  Mt[ht] = Yt;
  function Wn(t) {
    return Mt[t] !== 0;
  }
  class da {
    constructor(e, n) {
      this.setSource(e, n);
    }
    reset() {
      this.eof = !1, this.tokenIndex = -1, this.tokenType = 0, this.tokenStart = this.firstCharOffset, this.tokenEnd = this.firstCharOffset;
    }
    setSource(e = "", n = () => {
    }) {
      e = String(e || "");
      const s = e.length, r = Oe(this.offsetAndType, e.length + 1), o = Oe(this.balance, e.length + 1);
      let a = 0, l = -1, u = 0, i = e.length;
      this.offsetAndType = null, this.balance = null, o.fill(0), n(e, (c, h, f) => {
        const p = a++;
        if (r[p] = c << gt | f, l === -1 && (l = h), o[p] = i, c === u) {
          const d = o[i];
          o[i] = p, i = d, u = Mt[r[d] >> gt];
        } else Wn(c) && (i = p, u = Mt[c]);
      }), r[a] = vt << gt | s, o[a] = a;
      for (let c = 0; c < a; c++) {
        const h = o[c];
        if (h <= c) {
          const f = o[h];
          f !== c && (o[c] = f);
        } else h > a && (o[c] = a);
      }
      this.source = e, this.firstCharOffset = l === -1 ? 0 : l, this.tokenCount = a, this.offsetAndType = r, this.balance = o, this.reset(), this.next();
    }
    lookupType(e) {
      return e += this.tokenIndex, e < this.tokenCount ? this.offsetAndType[e] >> gt : vt;
    }
    lookupTypeNonSC(e) {
      for (let n = this.tokenIndex; n < this.tokenCount; n++) {
        const s = this.offsetAndType[n] >> gt;
        if (s !== W && s !== Y && e-- === 0)
          return s;
      }
      return vt;
    }
    lookupOffset(e) {
      return e += this.tokenIndex, e < this.tokenCount ? this.offsetAndType[e - 1] & dt : this.source.length;
    }
    lookupOffsetNonSC(e) {
      for (let n = this.tokenIndex; n < this.tokenCount; n++) {
        const s = this.offsetAndType[n] >> gt;
        if (s !== W && s !== Y && e-- === 0)
          return n - this.tokenIndex;
      }
      return vt;
    }
    lookupValue(e, n) {
      return e += this.tokenIndex, e < this.tokenCount ? Le(
        this.source,
        this.offsetAndType[e - 1] & dt,
        this.offsetAndType[e] & dt,
        n
      ) : !1;
    }
    getTokenStart(e) {
      return e === this.tokenIndex ? this.tokenStart : e > 0 ? e < this.tokenCount ? this.offsetAndType[e - 1] & dt : this.offsetAndType[this.tokenCount] & dt : this.firstCharOffset;
    }
    substrToCursor(e) {
      return this.source.substring(e, this.tokenStart);
    }
    isBalanceEdge(e) {
      return this.balance[this.tokenIndex] < e;
    }
    isDelim(e, n) {
      return n ? this.lookupType(n) === L && this.source.charCodeAt(this.lookupOffset(n)) === e : this.tokenType === L && this.source.charCodeAt(this.tokenStart) === e;
    }
    skip(e) {
      let n = this.tokenIndex + e;
      n < this.tokenCount ? (this.tokenIndex = n, this.tokenStart = this.offsetAndType[n - 1] & dt, n = this.offsetAndType[n], this.tokenType = n >> gt, this.tokenEnd = n & dt) : (this.tokenIndex = this.tokenCount, this.next());
    }
    next() {
      let e = this.tokenIndex + 1;
      e < this.tokenCount ? (this.tokenIndex = e, this.tokenStart = this.tokenEnd, e = this.offsetAndType[e], this.tokenType = e >> gt, this.tokenEnd = e & dt) : (this.eof = !0, this.tokenIndex = this.tokenCount, this.tokenType = vt, this.tokenStart = this.tokenEnd = this.source.length);
    }
    skipSC() {
      for (; this.tokenType === W || this.tokenType === Y; )
        this.next();
    }
    skipUntilBalanced(e, n) {
      let s = e, r = 0, o = 0;
      t:
        for (; s < this.tokenCount; s++) {
          if (r = this.balance[s], r < e)
            break t;
          switch (o = s > 0 ? this.offsetAndType[s - 1] & dt : this.firstCharOffset, n(this.source.charCodeAt(o))) {
            case 1:
              break t;
            case 2:
              s++;
              break t;
            default:
              Wn(this.offsetAndType[s] >> gt) && (s = r);
          }
        }
      this.skip(s - this.tokenIndex);
    }
    forEachToken(e) {
      for (let n = 0, s = this.firstCharOffset; n < this.tokenCount; n++) {
        const r = s, o = this.offsetAndType[n], a = o & dt, l = o >> gt;
        s = a, e(l, r, a, n);
      }
    }
    dump() {
      const e = new Array(this.tokenCount);
      return this.forEachToken((n, s, r, o) => {
        e[o] = {
          idx: o,
          type: Xs[n],
          chunk: this.source.substring(s, r),
          balance: this.balance[o]
        };
      }), e;
    }
  }
  function Js(t, e) {
    function n(h) {
      return h < l ? t.charCodeAt(h) : 0;
    }
    function s() {
      if (i = Qs(t, i), je(n(i), n(i + 1), n(i + 2))) {
        c = E, i = me(t, i);
        return;
      }
      if (n(i) === 37) {
        c = F, i++;
        return;
      }
      c = b;
    }
    function r() {
      const h = i;
      if (i = me(t, i), Le(t, h, i, "url") && n(i) === 40) {
        if (i = ge(t, i + 1), n(i) === 34 || n(i) === 39) {
          c = T, i = h + 4;
          return;
        }
        a();
        return;
      }
      if (n(i) === 40) {
        c = T, i++;
        return;
      }
      c = g;
    }
    function o(h) {
      for (h || (h = n(i++)), c = Tt; i < t.length; i++) {
        const f = t.charCodeAt(i);
        switch (We(f)) {
          // ending code point
          case h:
            i++;
            return;
          // EOF
          // case EofCategory:
          // This is a parse error. Return the <string-token>.
          // return;
          // newline
          case be:
            if ($e(f)) {
              i += ln(t, i, f), c = sa;
              return;
            }
            break;
          // U+005C REVERSE SOLIDUS (\)
          case 92:
            if (i === t.length - 1)
              break;
            const p = n(i + 1);
            $e(p) ? i += ln(t, i + 1, p) : kt(f, p) && (i = Jt(t, i) - 1);
            break;
        }
      }
    }
    function a() {
      for (c = Q, i = ge(t, i); i < t.length; i++) {
        const h = t.charCodeAt(i);
        switch (We(h)) {
          // U+0029 RIGHT PARENTHESIS ())
          case 41:
            i++;
            return;
          // EOF
          // case EofCategory:
          // This is a parse error. Return the <url-token>.
          // return;
          // whitespace
          case be:
            if (i = ge(t, i), n(i) === 41 || i >= t.length) {
              i < t.length && i++;
              return;
            }
            i = ze(t, i), c = at;
            return;
          // U+0022 QUOTATION MARK (")
          // U+0027 APOSTROPHE (')
          // U+0028 LEFT PARENTHESIS (()
          // non-printable code point
          case 34:
          case 39:
          case 40:
          case qs:
            i = ze(t, i), c = at;
            return;
          // U+005C REVERSE SOLIDUS (\)
          case 92:
            if (kt(h, n(i + 1))) {
              i = Jt(t, i) - 1;
              break;
            }
            i = ze(t, i), c = at;
            return;
        }
      }
    }
    t = String(t || "");
    const l = t.length;
    let u = Hs(n(0)), i = u, c;
    for (; i < l; ) {
      const h = t.charCodeAt(i);
      switch (We(h)) {
        // whitespace
        case be:
          c = W, i = ge(t, i + 1);
          break;
        // U+0022 QUOTATION MARK (")
        case 34:
          o();
          break;
        // U+0023 NUMBER SIGN (#)
        case 35:
          Vs(n(i + 1)) || kt(n(i + 1), n(i + 2)) ? (c = I, i = me(t, i + 1)) : (c = L, i++);
          break;
        // U+0027 APOSTROPHE (')
        case 39:
          o();
          break;
        // U+0028 LEFT PARENTHESIS (()
        case 40:
          c = _, i++;
          break;
        // U+0029 RIGHT PARENTHESIS ())
        case 41:
          c = S, i++;
          break;
        // U+002B PLUS SIGN (+)
        case 43:
          Be(h, n(i + 1), n(i + 2)) ? s() : (c = L, i++);
          break;
        // U+002C COMMA (,)
        case 44:
          c = ft, i++;
          break;
        // U+002D HYPHEN-MINUS (-)
        case 45:
          Be(h, n(i + 1), n(i + 2)) ? s() : n(i + 1) === 45 && n(i + 2) === 62 ? (c = nt, i = i + 3) : je(h, n(i + 1), n(i + 2)) ? r() : (c = L, i++);
          break;
        // U+002E FULL STOP (.)
        case 46:
          Be(h, n(i + 1), n(i + 2)) ? s() : (c = L, i++);
          break;
        // U+002F SOLIDUS (/)
        case 47:
          n(i + 1) === 42 ? (c = Y, i = t.indexOf("*/", i + 2), i = i === -1 ? t.length : i + 2) : (c = L, i++);
          break;
        // U+003A COLON (:)
        case 58:
          c = X, i++;
          break;
        // U+003B SEMICOLON (;)
        case 59:
          c = tt, i++;
          break;
        // U+003C LESS-THAN SIGN (<)
        case 60:
          n(i + 1) === 33 && n(i + 2) === 45 && n(i + 3) === 45 ? (c = Fe, i = i + 4) : (c = L, i++);
          break;
        // U+0040 COMMERCIAL AT (@)
        case 64:
          je(n(i + 1), n(i + 2), n(i + 3)) ? (c = z, i = me(t, i + 1)) : (c = L, i++);
          break;
        // U+005B LEFT SQUARE BRACKET ([)
        case 91:
          c = ee, i++;
          break;
        // U+005C REVERSE SOLIDUS (\)
        case 92:
          kt(h, n(i + 1)) ? r() : (c = L, i++);
          break;
        // U+005D RIGHT SQUARE BRACKET (])
        case 93:
          c = ce, i++;
          break;
        // U+007B LEFT CURLY BRACKET ({)
        case 123:
          c = ht, i++;
          break;
        // U+007D RIGHT CURLY BRACKET (})
        case 125:
          c = Yt, i++;
          break;
        // digit
        case Gs:
          s();
          break;
        // name-start code point
        case wn:
          r();
          break;
        // EOF
        // case EofCategory:
        // Return an <EOF-token>.
        // break;
        // anything else
        default:
          c = L, i++;
      }
      e(c, u, u = i);
    }
  }
  const mt = 43, st = 45, xe = 110, _t = !0, ga = !1;
  function Ce(t, e) {
    let n = this.tokenStart + t;
    const s = this.charCodeAt(n);
    for ((s === mt || s === st) && (e && this.error("Number sign is not allowed"), n++); n < this.tokenEnd; n++)
      rt(this.charCodeAt(n)) || this.error("Integer is expected", n);
  }
  function Vt(t) {
    return Ce.call(this, 0, t);
  }
  function At(t, e) {
    if (!this.cmpChar(this.tokenStart + t, e)) {
      let n = "";
      switch (e) {
        case xe:
          n = "N is expected";
          break;
        case st:
          n = "HyphenMinus is expected";
          break;
      }
      this.error(n, this.tokenStart + t);
    }
  }
  function Ve() {
    let t = 0, e = 0, n = this.tokenType;
    for (; n === W || n === Y; )
      n = this.lookupType(++t);
    if (n !== b)
      if (this.isDelim(mt, t) || this.isDelim(st, t)) {
        e = this.isDelim(mt, t) ? mt : st;
        do
          n = this.lookupType(++t);
        while (n === W || n === Y);
        n !== b && (this.skip(t), Vt.call(this, _t));
      } else
        return null;
    return t > 0 && this.skip(t), e === 0 && (n = this.charCodeAt(this.tokenStart), n !== mt && n !== st && this.error("Number sign is expected")), Vt.call(this, e !== 0), e === st ? "-" + this.consume(b) : this.consume(b);
  }
  const ma = "AnPlusB", ka = {
    a: [String, null],
    b: [String, null]
  };
  function Zs() {
    const t = this.tokenStart;
    let e = null, n = null;
    if (this.tokenType === b)
      Vt.call(this, ga), n = this.consume(b);
    else if (this.tokenType === g && this.cmpChar(this.tokenStart, st))
      switch (e = "-1", At.call(this, 1, xe), this.tokenEnd - this.tokenStart) {
        // -n
        // -n <signed-integer>
        // -n ['+' | '-'] <signless-integer>
        case 2:
          this.next(), n = Ve.call(this);
          break;
        // -n- <signless-integer>
        case 3:
          At.call(this, 2, st), this.next(), this.skipSC(), Vt.call(this, _t), n = "-" + this.consume(b);
          break;
        // <dashndashdigit-ident>
        default:
          At.call(this, 2, st), Ce.call(this, 3, _t), this.next(), n = this.substrToCursor(t + 2);
      }
    else if (this.tokenType === g || this.isDelim(mt) && this.lookupType(1) === g) {
      let s = 0;
      switch (e = "1", this.isDelim(mt) && (s = 1, this.next()), At.call(this, 0, xe), this.tokenEnd - this.tokenStart) {
        // '+'? n
        // '+'? n <signed-integer>
        // '+'? n ['+' | '-'] <signless-integer>
        case 1:
          this.next(), n = Ve.call(this);
          break;
        // '+'? n- <signless-integer>
        case 2:
          At.call(this, 1, st), this.next(), this.skipSC(), Vt.call(this, _t), n = "-" + this.consume(b);
          break;
        // '+'? <ndashdigit-ident>
        default:
          At.call(this, 1, st), Ce.call(this, 2, _t), this.next(), n = this.substrToCursor(t + s + 1);
      }
    } else if (this.tokenType === E) {
      const s = this.charCodeAt(this.tokenStart), r = s === mt || s === st;
      let o = this.tokenStart + r;
      for (; o < this.tokenEnd && rt(this.charCodeAt(o)); o++)
        ;
      o === this.tokenStart + r && this.error("Integer is expected", this.tokenStart + r), At.call(this, o - this.tokenStart, xe), e = this.substring(t, o), o + 1 === this.tokenEnd ? (this.next(), n = Ve.call(this)) : (At.call(this, o - this.tokenStart + 1, st), o + 2 === this.tokenEnd ? (this.next(), this.skipSC(), Vt.call(this, _t), n = "-" + this.consume(b)) : (Ce.call(this, o - this.tokenStart + 2, _t), this.next(), n = this.substrToCursor(o + 1)));
    } else
      this.error();
    return e !== null && e.charCodeAt(0) === mt && (e = e.substr(1)), n !== null && n.charCodeAt(0) === mt && (n = n.substr(1)), {
      type: "AnPlusB",
      loc: this.getLocation(t, this.tokenStart),
      a: e,
      b: n
    };
  }
  function tr(t) {
    if (t.a) {
      const e = t.a === "+1" && "n" || t.a === "1" && "n" || t.a === "-1" && "-n" || t.a + "n";
      if (t.b) {
        const n = t.b[0] === "-" || t.b[0] === "+" ? t.b : "+" + t.b;
        this.tokenize(e + n);
      } else
        this.tokenize(e);
    } else
      this.tokenize(t.b);
  }
  const Sa = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: tr,
    name: ma,
    parse: Zs,
    structure: ka
  }, Symbol.toStringTag, { value: "Module" }));
  function Un() {
    return this.Raw(this.consumeUntilLeftCurlyBracketOrSemicolon, !0);
  }
  function ya() {
    for (let t = 1, e; e = this.lookupType(t); t++) {
      if (e === Yt)
        return !0;
      if (e === ht || e === z)
        return !1;
    }
    return !1;
  }
  const ba = "Atrule", xa = "atrule", Ca = {
    name: String,
    prelude: ["AtrulePrelude", "Raw", null],
    block: ["Block", null]
  };
  function er(t = !1) {
    const e = this.tokenStart;
    let n, s, r = null, o = null;
    switch (this.eat(z), n = this.substrToCursor(e + 1), s = n.toLowerCase(), this.skipSC(), this.eof === !1 && this.tokenType !== ht && this.tokenType !== tt && (this.parseAtrulePrelude ? r = this.parseWithFallback(this.AtrulePrelude.bind(this, n, t), Un) : r = Un.call(this, this.tokenIndex), this.skipSC()), this.tokenType) {
      case tt:
        this.next();
        break;
      case ht:
        hasOwnProperty.call(this.atrule, s) && typeof this.atrule[s].block == "function" ? o = this.atrule[s].block.call(this, t) : o = this.Block(ya.call(this));
        break;
    }
    return {
      type: "Atrule",
      loc: this.getLocation(e, this.tokenStart),
      name: n,
      prelude: r,
      block: o
    };
  }
  function nr(t) {
    this.token(z, "@" + t.name), t.prelude !== null && this.node(t.prelude), t.block ? this.node(t.block) : this.token(tt, ";");
  }
  const wa = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: nr,
    name: ba,
    parse: er,
    structure: Ca,
    walkContext: xa
  }, Symbol.toStringTag, { value: "Module" })), Ta = "AtrulePrelude", Aa = "atrulePrelude", va = {
    children: [[]]
  };
  function sr(t) {
    let e = null;
    return t !== null && (t = t.toLowerCase()), this.skipSC(), hasOwnProperty.call(this.atrule, t) && typeof this.atrule[t].prelude == "function" ? e = this.atrule[t].prelude.call(this) : e = this.readSequence(this.scope.AtrulePrelude), this.skipSC(), this.eof !== !0 && this.tokenType !== ht && this.tokenType !== tt && this.error("Semicolon or block is expected"), {
      type: "AtrulePrelude",
      loc: this.getLocationFromList(e),
      children: e
    };
  }
  function rr(t) {
    this.children(t);
  }
  const Ea = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: rr,
    name: Ta,
    parse: sr,
    structure: va,
    walkContext: Aa
  }, Symbol.toStringTag, { value: "Module" })), $a = 36, ir = 42, we = 61, La = 94, cn = 124, Oa = 126;
  function Pa() {
    this.eof && this.error("Unexpected end of input");
    const t = this.tokenStart;
    let e = !1;
    return this.isDelim(ir) ? (e = !0, this.next()) : this.isDelim(cn) || this.eat(g), this.isDelim(cn) ? this.charCodeAt(this.tokenStart + 1) !== we ? (this.next(), this.eat(g)) : e && this.error("Identifier is expected", this.tokenEnd) : e && this.error("Vertical line is expected"), {
      type: "Identifier",
      loc: this.getLocation(t, this.tokenStart),
      name: this.substrToCursor(t)
    };
  }
  function Ra() {
    const t = this.tokenStart, e = this.charCodeAt(t);
    return e !== we && // =
    e !== Oa && // ~=
    e !== La && // ^=
    e !== $a && // $=
    e !== ir && // *=
    e !== cn && this.error("Attribute selector (=, ~=, ^=, $=, *=, |=) is expected"), this.next(), e !== we && (this.isDelim(we) || this.error("Equal sign is expected"), this.next()), this.substrToCursor(t);
  }
  const _a = "AttributeSelector", Ia = {
    name: "Identifier",
    matcher: [String, null],
    value: ["String", "Identifier", null],
    flags: [String, null]
  };
  function or() {
    const t = this.tokenStart;
    let e, n = null, s = null, r = null;
    return this.eat(ee), this.skipSC(), e = Pa.call(this), this.skipSC(), this.tokenType !== ce && (this.tokenType !== g && (n = Ra.call(this), this.skipSC(), s = this.tokenType === Tt ? this.String() : this.Identifier(), this.skipSC()), this.tokenType === g && (r = this.consume(g), this.skipSC())), this.eat(ce), {
      type: "AttributeSelector",
      loc: this.getLocation(t, this.tokenStart),
      name: e,
      matcher: n,
      value: s,
      flags: r
    };
  }
  function ar(t) {
    this.token(L, "["), this.node(t.name), t.matcher !== null && (this.tokenize(t.matcher), this.node(t.value)), t.flags !== null && this.token(g, t.flags), this.token(L, "]");
  }
  const Na = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: ar,
    name: _a,
    parse: or,
    structure: Ia
  }, Symbol.toStringTag, { value: "Module" })), Da = 38;
  function lr() {
    return this.Raw(null, !0);
  }
  function zn() {
    return this.parseWithFallback(this.Rule, lr);
  }
  function Vn() {
    return this.Raw(this.consumeUntilSemicolonIncluded, !0);
  }
  function Fa() {
    if (this.tokenType === tt)
      return Vn.call(this, this.tokenIndex);
    const t = this.parseWithFallback(this.Declaration, Vn);
    return this.tokenType === tt && this.next(), t;
  }
  const Ma = "Block", ja = "block", Ba = {
    children: [[
      "Atrule",
      "Rule",
      "Declaration"
    ]]
  };
  function cr(t) {
    const e = t ? Fa : zn, n = this.tokenStart;
    let s = this.createList();
    this.eat(ht);
    t:
      for (; !this.eof; )
        switch (this.tokenType) {
          case Yt:
            break t;
          case W:
          case Y:
            this.next();
            break;
          case z:
            s.push(this.parseWithFallback(this.Atrule.bind(this, t), lr));
            break;
          default:
            t && this.isDelim(Da) ? s.push(zn.call(this)) : s.push(e.call(this));
        }
    return this.eof || this.eat(Yt), {
      type: "Block",
      loc: this.getLocation(n, this.tokenStart),
      children: s
    };
  }
  function ur(t) {
    this.token(ht, "{"), this.children(t, (e) => {
      e.type === "Declaration" && this.token(tt, ";");
    }), this.token(Yt, "}");
  }
  const Wa = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: ur,
    name: Ma,
    parse: cr,
    structure: Ba,
    walkContext: ja
  }, Symbol.toStringTag, { value: "Module" })), Ua = "Brackets", za = {
    children: [[]]
  };
  function hr(t, e) {
    const n = this.tokenStart;
    let s = null;
    return this.eat(ee), s = t.call(this, e), this.eof || this.eat(ce), {
      type: "Brackets",
      loc: this.getLocation(n, this.tokenStart),
      children: s
    };
  }
  function fr(t) {
    this.token(L, "["), this.children(t), this.token(L, "]");
  }
  const Va = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: fr,
    name: Ua,
    parse: hr,
    structure: za
  }, Symbol.toStringTag, { value: "Module" })), Ha = "CDC", Ga = [];
  function pr() {
    const t = this.tokenStart;
    return this.eat(nt), {
      type: "CDC",
      loc: this.getLocation(t, this.tokenStart)
    };
  }
  function dr() {
    this.token(nt, "-->");
  }
  const qa = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: dr,
    name: Ha,
    parse: pr,
    structure: Ga
  }, Symbol.toStringTag, { value: "Module" })), Ka = "CDO", Qa = [];
  function gr() {
    const t = this.tokenStart;
    return this.eat(Fe), {
      type: "CDO",
      loc: this.getLocation(t, this.tokenStart)
    };
  }
  function mr() {
    this.token(Fe, "<!--");
  }
  const Ya = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: mr,
    name: Ka,
    parse: gr,
    structure: Qa
  }, Symbol.toStringTag, { value: "Module" })), Xa = 46, Ja = "ClassSelector", Za = {
    name: String
  };
  function kr() {
    return this.eatDelim(Xa), {
      type: "ClassSelector",
      loc: this.getLocation(this.tokenStart - 1, this.tokenEnd),
      name: this.consume(g)
    };
  }
  function Sr(t) {
    this.token(L, "."), this.token(g, t.name);
  }
  const tl = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Sr,
    name: Ja,
    parse: kr,
    structure: Za
  }, Symbol.toStringTag, { value: "Module" })), el = 43, Hn = 47, nl = 62, sl = 126, rl = "Combinator", il = {
    name: String
  };
  function yr() {
    const t = this.tokenStart;
    let e;
    switch (this.tokenType) {
      case W:
        e = " ";
        break;
      case L:
        switch (this.charCodeAt(this.tokenStart)) {
          case nl:
          case el:
          case sl:
            this.next();
            break;
          case Hn:
            this.next(), this.eatIdent("deep"), this.eatDelim(Hn);
            break;
          default:
            this.error("Combinator is expected");
        }
        e = this.substrToCursor(t);
        break;
    }
    return {
      type: "Combinator",
      loc: this.getLocation(t, this.tokenStart),
      name: e
    };
  }
  function br(t) {
    this.tokenize(t.name);
  }
  const ol = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: br,
    name: rl,
    parse: yr,
    structure: il
  }, Symbol.toStringTag, { value: "Module" })), al = 42, ll = 47, cl = "Comment", ul = {
    value: String
  };
  function xr() {
    const t = this.tokenStart;
    let e = this.tokenEnd;
    return this.eat(Y), e - t + 2 >= 2 && this.charCodeAt(e - 2) === al && this.charCodeAt(e - 1) === ll && (e -= 2), {
      type: "Comment",
      loc: this.getLocation(t, this.tokenStart),
      value: this.substring(t + 2, e)
    };
  }
  function Cr(t) {
    this.token(Y, "/*" + t.value + "*/");
  }
  const hl = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Cr,
    name: cl,
    parse: xr,
    structure: ul
  }, Symbol.toStringTag, { value: "Module" })), fl = /* @__PURE__ */ new Set([X, S, vt]), pl = "Condition", dl = {
    kind: String,
    children: [[
      "Identifier",
      "Feature",
      "FeatureFunction",
      "FeatureRange",
      "SupportsDeclaration"
    ]]
  };
  function Gn(t) {
    return this.lookupTypeNonSC(1) === g && fl.has(this.lookupTypeNonSC(2)) ? this.Feature(t) : this.FeatureRange(t);
  }
  const gl = {
    media: Gn,
    container: Gn,
    supports() {
      return this.SupportsDeclaration();
    }
  };
  function wr(t = "media") {
    const e = this.createList();
    t: for (; !this.eof; )
      switch (this.tokenType) {
        case Y:
        case W:
          this.next();
          continue;
        case g:
          e.push(this.Identifier());
          break;
        case _: {
          let n = this.parseWithFallback(
            () => gl[t].call(this, t),
            () => null
          );
          n || (n = this.parseWithFallback(
            () => {
              this.eat(_);
              const s = this.Condition(t);
              return this.eat(S), s;
            },
            () => this.GeneralEnclosed(t)
          )), e.push(n);
          break;
        }
        case T: {
          let n = this.parseWithFallback(
            () => this.FeatureFunction(t),
            () => null
          );
          n || (n = this.GeneralEnclosed(t)), e.push(n);
          break;
        }
        default:
          break t;
      }
    return e.isEmpty && this.error("Condition is expected"), {
      type: "Condition",
      loc: this.getLocationFromList(e),
      kind: t,
      children: e
    };
  }
  function Tr(t) {
    t.children.forEach((e) => {
      e.type === "Condition" ? (this.token(_, "("), this.node(e), this.token(S, ")")) : this.node(e);
    });
  }
  const ml = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Tr,
    name: pl,
    parse: wr,
    structure: dl
  }, Symbol.toStringTag, { value: "Module" })), qn = 45;
  function kl(t, e) {
    return e = e || 0, t.length - e >= 2 && t.charCodeAt(e) === qn && t.charCodeAt(e + 1) === qn;
  }
  const Ar = 33, Sl = 35, yl = 36, bl = 38, xl = 42, Cl = 43, Kn = 47;
  function wl() {
    return this.Raw(this.consumeUntilExclamationMarkOrSemicolon, !0);
  }
  function Tl() {
    return this.Raw(this.consumeUntilExclamationMarkOrSemicolon, !1);
  }
  function Al() {
    const t = this.tokenIndex, e = this.Value();
    return e.type !== "Raw" && this.eof === !1 && this.tokenType !== tt && this.isDelim(Ar) === !1 && this.isBalanceEdge(t) === !1 && this.error(), e;
  }
  const vl = "Declaration", El = "declaration", $l = {
    important: [Boolean, String],
    property: String,
    value: ["Value", "Raw"]
  };
  function vr() {
    const t = this.tokenStart, e = this.tokenIndex, n = Ll.call(this), s = kl(n), r = s ? this.parseCustomProperty : this.parseValue, o = s ? Tl : wl;
    let a = !1, l;
    this.skipSC(), this.eat(X);
    const u = this.tokenIndex;
    if (s || this.skipSC(), r ? l = this.parseWithFallback(Al, o) : l = o.call(this, this.tokenIndex), s && l.type === "Value" && l.children.isEmpty) {
      for (let i = u - this.tokenIndex; i <= 0; i++)
        if (this.lookupType(i) === W) {
          l.children.appendData({
            type: "WhiteSpace",
            loc: null,
            value: " "
          });
          break;
        }
    }
    return this.isDelim(Ar) && (a = Ol.call(this), this.skipSC()), this.eof === !1 && this.tokenType !== tt && this.isBalanceEdge(e) === !1 && this.error(), {
      type: "Declaration",
      loc: this.getLocation(t, this.tokenStart),
      important: a,
      property: n,
      value: l
    };
  }
  function Er(t) {
    this.token(g, t.property), this.token(X, ":"), this.node(t.value), t.important && (this.token(L, "!"), this.token(g, t.important === !0 ? "important" : t.important));
  }
  function Ll() {
    const t = this.tokenStart;
    if (this.tokenType === L)
      switch (this.charCodeAt(this.tokenStart)) {
        case xl:
        case yl:
        case Cl:
        case Sl:
        case bl:
          this.next();
          break;
        // TODO: not sure we should support this hack
        case Kn:
          this.next(), this.isDelim(Kn) && this.next();
          break;
      }
    return this.tokenType === I ? this.eat(I) : this.eat(g), this.substrToCursor(t);
  }
  function Ol() {
    this.eat(L), this.skipSC();
    const t = this.consume(g);
    return t === "important" ? !0 : t;
  }
  const Pl = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Er,
    name: vl,
    parse: vr,
    structure: $l,
    walkContext: El
  }, Symbol.toStringTag, { value: "Module" })), Rl = 38;
  function He() {
    return this.Raw(this.consumeUntilSemicolonIncluded, !0);
  }
  const _l = "DeclarationList", Il = {
    children: [[
      "Declaration",
      "Atrule",
      "Rule"
    ]]
  };
  function $r() {
    const t = this.createList();
    for (; !this.eof; )
      switch (this.tokenType) {
        case W:
        case Y:
        case tt:
          this.next();
          break;
        case z:
          t.push(this.parseWithFallback(this.Atrule.bind(this, !0), He));
          break;
        default:
          this.isDelim(Rl) ? t.push(this.parseWithFallback(this.Rule, He)) : t.push(this.parseWithFallback(this.Declaration, He));
      }
    return {
      type: "DeclarationList",
      loc: this.getLocationFromList(t),
      children: t
    };
  }
  function Lr(t) {
    this.children(t, (e) => {
      e.type === "Declaration" && this.token(tt, ";");
    });
  }
  const Nl = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Lr,
    name: _l,
    parse: $r,
    structure: Il
  }, Symbol.toStringTag, { value: "Module" })), Dl = "Dimension", Fl = {
    value: String,
    unit: String
  };
  function Or() {
    const t = this.tokenStart, e = this.consumeNumber(E);
    return {
      type: "Dimension",
      loc: this.getLocation(t, this.tokenStart),
      value: e,
      unit: this.substring(t + e.length, this.tokenStart)
    };
  }
  function Pr(t) {
    this.token(E, t.value + t.unit);
  }
  const Ml = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Pr,
    name: Dl,
    parse: Or,
    structure: Fl
  }, Symbol.toStringTag, { value: "Module" })), jl = 47, Bl = "Feature", Wl = {
    kind: String,
    name: String,
    value: ["Identifier", "Number", "Dimension", "Ratio", "Function", null]
  };
  function Rr(t) {
    const e = this.tokenStart;
    let n, s = null;
    if (this.eat(_), this.skipSC(), n = this.consume(g), this.skipSC(), this.tokenType !== S) {
      switch (this.eat(X), this.skipSC(), this.tokenType) {
        case b:
          this.lookupNonWSType(1) === L ? s = this.Ratio() : s = this.Number();
          break;
        case E:
          s = this.Dimension();
          break;
        case g:
          s = this.Identifier();
          break;
        case T:
          s = this.parseWithFallback(
            () => {
              const r = this.Function(this.readSequence, this.scope.Value);
              return this.skipSC(), this.isDelim(jl) && this.error(), r;
            },
            () => this.Ratio()
          );
          break;
        default:
          this.error("Number, dimension, ratio or identifier is expected");
      }
      this.skipSC();
    }
    return this.eof || this.eat(S), {
      type: "Feature",
      loc: this.getLocation(e, this.tokenStart),
      kind: t,
      name: n,
      value: s
    };
  }
  function _r(t) {
    this.token(_, "("), this.token(g, t.name), t.value !== null && (this.token(X, ":"), this.node(t.value)), this.token(S, ")");
  }
  const Ul = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: _r,
    name: Bl,
    parse: Rr,
    structure: Wl
  }, Symbol.toStringTag, { value: "Module" })), zl = "FeatureFunction", Vl = {
    kind: String,
    feature: String,
    value: ["Declaration", "Selector"]
  };
  function Hl(t, e) {
    const s = (this.features[t] || {})[e];
    return typeof s != "function" && this.error(`Unknown feature ${e}()`), s;
  }
  function Ir(t = "unknown") {
    const e = this.tokenStart, n = this.consumeFunctionName(), s = Hl.call(this, t, n.toLowerCase());
    this.skipSC();
    const r = this.parseWithFallback(
      () => {
        const o = this.tokenIndex, a = s.call(this);
        return this.eof === !1 && this.isBalanceEdge(o) === !1 && this.error(), a;
      },
      () => this.Raw(null, !1)
    );
    return this.eof || this.eat(S), {
      type: "FeatureFunction",
      loc: this.getLocation(e, this.tokenStart),
      kind: t,
      feature: n,
      value: r
    };
  }
  function Nr(t) {
    this.token(T, t.feature + "("), this.node(t.value), this.token(S, ")");
  }
  const Gl = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Nr,
    name: zl,
    parse: Ir,
    structure: Vl
  }, Symbol.toStringTag, { value: "Module" })), Qn = 47, ql = 60, Yn = 61, Kl = 62, Ql = "FeatureRange", Yl = {
    kind: String,
    left: ["Identifier", "Number", "Dimension", "Ratio", "Function"],
    leftComparison: String,
    middle: ["Identifier", "Number", "Dimension", "Ratio", "Function"],
    rightComparison: [String, null],
    right: ["Identifier", "Number", "Dimension", "Ratio", "Function", null]
  };
  function Ge() {
    switch (this.skipSC(), this.tokenType) {
      case b:
        return this.isDelim(Qn, this.lookupOffsetNonSC(1)) ? this.Ratio() : this.Number();
      case E:
        return this.Dimension();
      case g:
        return this.Identifier();
      case T:
        return this.parseWithFallback(
          () => {
            const t = this.Function(this.readSequence, this.scope.Value);
            return this.skipSC(), this.isDelim(Qn) && this.error(), t;
          },
          () => this.Ratio()
        );
      default:
        this.error("Number, dimension, ratio or identifier is expected");
    }
  }
  function Xn(t) {
    if (this.skipSC(), this.isDelim(ql) || this.isDelim(Kl)) {
      const e = this.source[this.tokenStart];
      return this.next(), this.isDelim(Yn) ? (this.next(), e + "=") : e;
    }
    if (this.isDelim(Yn))
      return "=";
    this.error(`Expected ${t ? '":", ' : ""}"<", ">", "=" or ")"`);
  }
  function Dr(t = "unknown") {
    const e = this.tokenStart;
    this.skipSC(), this.eat(_);
    const n = Ge.call(this), s = Xn.call(this, n.type === "Identifier"), r = Ge.call(this);
    let o = null, a = null;
    return this.lookupNonWSType(0) !== S && (o = Xn.call(this), a = Ge.call(this)), this.skipSC(), this.eat(S), {
      type: "FeatureRange",
      loc: this.getLocation(e, this.tokenStart),
      kind: t,
      left: n,
      leftComparison: s,
      middle: r,
      rightComparison: o,
      right: a
    };
  }
  function Fr(t) {
    this.token(_, "("), this.node(t.left), this.tokenize(t.leftComparison), this.node(t.middle), t.right && (this.tokenize(t.rightComparison), this.node(t.right)), this.token(S, ")");
  }
  const Xl = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Fr,
    name: Ql,
    parse: Dr,
    structure: Yl
  }, Symbol.toStringTag, { value: "Module" })), Jl = "Function", Zl = "function", tc = {
    name: String,
    children: [[]]
  };
  function Mr(t, e) {
    const n = this.tokenStart, s = this.consumeFunctionName(), r = s.toLowerCase();
    let o;
    return o = e.hasOwnProperty(r) ? e[r].call(this, e) : t.call(this, e), this.eof || this.eat(S), {
      type: "Function",
      loc: this.getLocation(n, this.tokenStart),
      name: s,
      children: o
    };
  }
  function jr(t) {
    this.token(T, t.name + "("), this.children(t), this.token(S, ")");
  }
  const ec = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: jr,
    name: Jl,
    parse: Mr,
    structure: tc,
    walkContext: Zl
  }, Symbol.toStringTag, { value: "Module" })), nc = "GeneralEnclosed", sc = {
    kind: String,
    function: [String, null],
    children: [[]]
  };
  function Br(t) {
    const e = this.tokenStart;
    let n = null;
    this.tokenType === T ? n = this.consumeFunctionName() : this.eat(_);
    const s = this.parseWithFallback(
      () => {
        const r = this.tokenIndex, o = this.readSequence(this.scope.Value);
        return this.eof === !1 && this.isBalanceEdge(r) === !1 && this.error(), o;
      },
      () => this.createSingleNodeList(
        this.Raw(null, !1)
      )
    );
    return this.eof || this.eat(S), {
      type: "GeneralEnclosed",
      loc: this.getLocation(e, this.tokenStart),
      kind: t,
      function: n,
      children: s
    };
  }
  function Wr(t) {
    t.function ? this.token(T, t.function + "(") : this.token(_, "("), this.children(t), this.token(S, ")");
  }
  const rc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Wr,
    name: nc,
    parse: Br,
    structure: sc
  }, Symbol.toStringTag, { value: "Module" })), ic = "XXX", oc = "Hash", ac = {
    value: String
  };
  function Ur() {
    const t = this.tokenStart;
    return this.eat(I), {
      type: "Hash",
      loc: this.getLocation(t, this.tokenStart),
      value: this.substrToCursor(t + 1)
    };
  }
  function zr(t) {
    this.token(I, "#" + t.value);
  }
  const lc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: zr,
    name: oc,
    parse: Ur,
    structure: ac,
    xxx: ic
  }, Symbol.toStringTag, { value: "Module" })), cc = "Identifier", uc = {
    name: String
  };
  function Vr() {
    return {
      type: "Identifier",
      loc: this.getLocation(this.tokenStart, this.tokenEnd),
      name: this.consume(g)
    };
  }
  function Hr(t) {
    this.token(g, t.name);
  }
  const hc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Hr,
    name: cc,
    parse: Vr,
    structure: uc
  }, Symbol.toStringTag, { value: "Module" })), fc = "IdSelector", pc = {
    name: String
  };
  function Gr() {
    const t = this.tokenStart;
    return this.eat(I), {
      type: "IdSelector",
      loc: this.getLocation(t, this.tokenStart),
      name: this.substrToCursor(t + 1)
    };
  }
  function qr(t) {
    this.token(L, "#" + t.name);
  }
  const dc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: qr,
    name: fc,
    parse: Gr,
    structure: pc
  }, Symbol.toStringTag, { value: "Module" })), gc = 46, mc = "Layer", kc = {
    name: String
  };
  function Kr() {
    let t = this.tokenStart, e = this.consume(g);
    for (; this.isDelim(gc); )
      this.eat(L), e += "." + this.consume(g);
    return {
      type: "Layer",
      loc: this.getLocation(t, this.tokenStart),
      name: e
    };
  }
  function Qr(t) {
    this.tokenize(t.name);
  }
  const Sc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Qr,
    name: mc,
    parse: Kr,
    structure: kc
  }, Symbol.toStringTag, { value: "Module" })), yc = "LayerList", bc = {
    children: [[
      "Layer"
    ]]
  };
  function Yr() {
    const t = this.createList();
    for (this.skipSC(); !this.eof && (t.push(this.Layer()), this.lookupTypeNonSC(0) === ft); )
      this.skipSC(), this.next(), this.skipSC();
    return {
      type: "LayerList",
      loc: this.getLocationFromList(t),
      children: t
    };
  }
  function Xr(t) {
    this.children(t, () => this.token(ft, ","));
  }
  const xc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Xr,
    name: yc,
    parse: Yr,
    structure: bc
  }, Symbol.toStringTag, { value: "Module" })), Cc = "MediaQuery", wc = {
    modifier: [String, null],
    mediaType: [String, null],
    condition: ["Condition", null]
  };
  function Jr() {
    const t = this.tokenStart;
    let e = null, n = null, s = null;
    if (this.skipSC(), this.tokenType === g && this.lookupTypeNonSC(1) !== _) {
      const r = this.consume(g), o = r.toLowerCase();
      switch (o === "not" || o === "only" ? (this.skipSC(), e = o, n = this.consume(g)) : n = r, this.lookupTypeNonSC(0)) {
        case g: {
          this.skipSC(), this.eatIdent("and"), s = this.Condition("media");
          break;
        }
        case ht:
        case tt:
        case ft:
        case vt:
          break;
        default:
          this.error("Identifier or parenthesis is expected");
      }
    } else
      switch (this.tokenType) {
        case g:
        case _:
        case T: {
          s = this.Condition("media");
          break;
        }
        case ht:
        case tt:
        case vt:
          break;
        default:
          this.error("Identifier or parenthesis is expected");
      }
    return {
      type: "MediaQuery",
      loc: this.getLocation(t, this.tokenStart),
      modifier: e,
      mediaType: n,
      condition: s
    };
  }
  function Zr(t) {
    t.mediaType ? (t.modifier && this.token(g, t.modifier), this.token(g, t.mediaType), t.condition && (this.token(g, "and"), this.node(t.condition))) : t.condition && this.node(t.condition);
  }
  const Tc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Zr,
    name: Cc,
    parse: Jr,
    structure: wc
  }, Symbol.toStringTag, { value: "Module" })), Ac = "MediaQueryList", vc = {
    children: [[
      "MediaQuery"
    ]]
  };
  function ti() {
    const t = this.createList();
    for (this.skipSC(); !this.eof && (t.push(this.MediaQuery()), this.tokenType === ft); )
      this.next();
    return {
      type: "MediaQueryList",
      loc: this.getLocationFromList(t),
      children: t
    };
  }
  function ei(t) {
    this.children(t, () => this.token(ft, ","));
  }
  const Ec = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: ei,
    name: Ac,
    parse: ti,
    structure: vc
  }, Symbol.toStringTag, { value: "Module" })), $c = 38, Lc = "NestingSelector", Oc = {};
  function ni() {
    const t = this.tokenStart;
    return this.eatDelim($c), {
      type: "NestingSelector",
      loc: this.getLocation(t, this.tokenStart)
    };
  }
  function si() {
    this.token(L, "&");
  }
  const Pc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: si,
    name: Lc,
    parse: ni,
    structure: Oc
  }, Symbol.toStringTag, { value: "Module" })), Rc = "Nth", _c = {
    nth: ["AnPlusB", "Identifier"],
    selector: ["SelectorList", null]
  };
  function ri() {
    this.skipSC();
    const t = this.tokenStart;
    let e = t, n = null, s;
    return this.lookupValue(0, "odd") || this.lookupValue(0, "even") ? s = this.Identifier() : s = this.AnPlusB(), e = this.tokenStart, this.skipSC(), this.lookupValue(0, "of") && (this.next(), n = this.SelectorList(), e = this.tokenStart), {
      type: "Nth",
      loc: this.getLocation(t, e),
      nth: s,
      selector: n
    };
  }
  function ii(t) {
    this.node(t.nth), t.selector !== null && (this.token(g, "of"), this.node(t.selector));
  }
  const Ic = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: ii,
    name: Rc,
    parse: ri,
    structure: _c
  }, Symbol.toStringTag, { value: "Module" })), Nc = "Number", Dc = {
    value: String
  };
  function oi() {
    return {
      type: "Number",
      loc: this.getLocation(this.tokenStart, this.tokenEnd),
      value: this.consume(b)
    };
  }
  function ai(t) {
    this.token(b, t.value);
  }
  const Fc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: ai,
    name: Nc,
    parse: oi,
    structure: Dc
  }, Symbol.toStringTag, { value: "Module" })), Mc = "Operator", jc = {
    value: String
  };
  function li() {
    const t = this.tokenStart;
    return this.next(), {
      type: "Operator",
      loc: this.getLocation(t, this.tokenStart),
      value: this.substrToCursor(t)
    };
  }
  function ci(t) {
    this.tokenize(t.value);
  }
  const Bc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: ci,
    name: Mc,
    parse: li,
    structure: jc
  }, Symbol.toStringTag, { value: "Module" })), Wc = "Parentheses", Uc = {
    children: [[]]
  };
  function ui(t, e) {
    const n = this.tokenStart;
    let s = null;
    return this.eat(_), s = t.call(this, e), this.eof || this.eat(S), {
      type: "Parentheses",
      loc: this.getLocation(n, this.tokenStart),
      children: s
    };
  }
  function hi(t) {
    this.token(_, "("), this.children(t), this.token(S, ")");
  }
  const zc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: hi,
    name: Wc,
    parse: ui,
    structure: Uc
  }, Symbol.toStringTag, { value: "Module" })), Vc = "Percentage", Hc = {
    value: String
  };
  function fi() {
    return {
      type: "Percentage",
      loc: this.getLocation(this.tokenStart, this.tokenEnd),
      value: this.consumeNumber(F)
    };
  }
  function pi(t) {
    this.token(F, t.value + "%");
  }
  const Gc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: pi,
    name: Vc,
    parse: fi,
    structure: Hc
  }, Symbol.toStringTag, { value: "Module" })), qc = "PseudoClassSelector", Kc = "function", Qc = {
    name: String,
    children: [["Raw"], null]
  };
  function di() {
    const t = this.tokenStart;
    let e = null, n, s;
    return this.eat(X), this.tokenType === T ? (n = this.consumeFunctionName(), s = n.toLowerCase(), this.lookupNonWSType(0) == S ? e = this.createList() : hasOwnProperty.call(this.pseudo, s) ? (this.skipSC(), e = this.pseudo[s].call(this), this.skipSC()) : (e = this.createList(), e.push(
      this.Raw(null, !1)
    )), this.eat(S)) : n = this.consume(g), {
      type: "PseudoClassSelector",
      loc: this.getLocation(t, this.tokenStart),
      name: n,
      children: e
    };
  }
  function gi(t) {
    this.token(X, ":"), t.children === null ? this.token(g, t.name) : (this.token(T, t.name + "("), this.children(t), this.token(S, ")"));
  }
  const Yc = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: gi,
    name: qc,
    parse: di,
    structure: Qc,
    walkContext: Kc
  }, Symbol.toStringTag, { value: "Module" })), Xc = "PseudoElementSelector", Jc = "function", Zc = {
    name: String,
    children: [["Raw"], null]
  };
  function mi() {
    const t = this.tokenStart;
    let e = null, n, s;
    return this.eat(X), this.eat(X), this.tokenType === T ? (n = this.consumeFunctionName(), s = n.toLowerCase(), this.lookupNonWSType(0) == S ? e = this.createList() : hasOwnProperty.call(this.pseudo, s) ? (this.skipSC(), e = this.pseudo[s].call(this), this.skipSC()) : (e = this.createList(), e.push(
      this.Raw(null, !1)
    )), this.eat(S)) : n = this.consume(g), {
      type: "PseudoElementSelector",
      loc: this.getLocation(t, this.tokenStart),
      name: n,
      children: e
    };
  }
  function ki(t) {
    this.token(X, ":"), this.token(X, ":"), t.children === null ? this.token(g, t.name) : (this.token(T, t.name + "("), this.children(t), this.token(S, ")"));
  }
  const tu = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: ki,
    name: Xc,
    parse: mi,
    structure: Zc,
    walkContext: Jc
  }, Symbol.toStringTag, { value: "Module" })), Jn = 47;
  function Zn() {
    switch (this.skipSC(), this.tokenType) {
      case b:
        return this.Number();
      case T:
        return this.Function(this.readSequence, this.scope.Value);
      default:
        this.error("Number of function is expected");
    }
  }
  const eu = "Ratio", nu = {
    left: ["Number", "Function"],
    right: ["Number", "Function", null]
  };
  function Si() {
    const t = this.tokenStart, e = Zn.call(this);
    let n = null;
    return this.skipSC(), this.isDelim(Jn) && (this.eatDelim(Jn), n = Zn.call(this)), {
      type: "Ratio",
      loc: this.getLocation(t, this.tokenStart),
      left: e,
      right: n
    };
  }
  function yi(t) {
    this.node(t.left), this.token(L, "/"), t.right ? this.node(t.right) : this.node(b, 1);
  }
  const su = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: yi,
    name: eu,
    parse: Si,
    structure: nu
  }, Symbol.toStringTag, { value: "Module" }));
  function ru() {
    return this.tokenIndex > 0 && this.lookupType(-1) === W ? this.tokenIndex > 1 ? this.getTokenStart(this.tokenIndex - 1) : this.firstCharOffset : this.tokenStart;
  }
  const iu = "Raw", ou = {
    value: String
  };
  function bi(t, e) {
    const n = this.getTokenStart(this.tokenIndex);
    let s;
    return this.skipUntilBalanced(this.tokenIndex, t || this.consumeUntilBalanceEnd), e && this.tokenStart > n ? s = ru.call(this) : s = this.tokenStart, {
      type: "Raw",
      loc: this.getLocation(n, s),
      value: this.substring(n, s)
    };
  }
  function xi(t) {
    this.tokenize(t.value);
  }
  const au = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: xi,
    name: iu,
    parse: bi,
    structure: ou
  }, Symbol.toStringTag, { value: "Module" }));
  function ts() {
    return this.Raw(this.consumeUntilLeftCurlyBracket, !0);
  }
  function lu() {
    const t = this.SelectorList();
    return t.type !== "Raw" && this.eof === !1 && this.tokenType !== ht && this.error(), t;
  }
  const cu = "Rule", uu = "rule", hu = {
    prelude: ["SelectorList", "Raw"],
    block: ["Block"]
  };
  function Ci() {
    const t = this.tokenIndex, e = this.tokenStart;
    let n, s;
    return this.parseRulePrelude ? n = this.parseWithFallback(lu, ts) : n = ts.call(this, t), s = this.Block(!0), {
      type: "Rule",
      loc: this.getLocation(e, this.tokenStart),
      prelude: n,
      block: s
    };
  }
  function wi(t) {
    this.node(t.prelude), this.node(t.block);
  }
  const fu = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: wi,
    name: cu,
    parse: Ci,
    structure: hu,
    walkContext: uu
  }, Symbol.toStringTag, { value: "Module" })), pu = "Scope", du = {
    root: ["SelectorList", "Raw", null],
    limit: ["SelectorList", "Raw", null]
  };
  function Ti() {
    let t = null, e = null;
    this.skipSC();
    const n = this.tokenStart;
    return this.tokenType === _ && (this.next(), this.skipSC(), t = this.parseWithFallback(
      this.SelectorList,
      () => this.Raw(!1, !0)
    ), this.skipSC(), this.eat(S)), this.lookupNonWSType(0) === g && (this.skipSC(), this.eatIdent("to"), this.skipSC(), this.eat(_), this.skipSC(), e = this.parseWithFallback(
      this.SelectorList,
      () => this.Raw(!1, !0)
    ), this.skipSC(), this.eat(S)), {
      type: "Scope",
      loc: this.getLocation(n, this.tokenStart),
      root: t,
      limit: e
    };
  }
  function Ai(t) {
    t.root && (this.token(_, "("), this.node(t.root), this.token(S, ")")), t.limit && (this.token(g, "to"), this.token(_, "("), this.node(t.limit), this.token(S, ")"));
  }
  const gu = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Ai,
    name: pu,
    parse: Ti,
    structure: du
  }, Symbol.toStringTag, { value: "Module" })), mu = "Selector", ku = {
    children: [[
      "TypeSelector",
      "IdSelector",
      "ClassSelector",
      "AttributeSelector",
      "PseudoClassSelector",
      "PseudoElementSelector",
      "Combinator"
    ]]
  };
  function vi() {
    const t = this.readSequence(this.scope.Selector);
    return this.getFirstListNode(t) === null && this.error("Selector is expected"), {
      type: "Selector",
      loc: this.getLocationFromList(t),
      children: t
    };
  }
  function Ei(t) {
    this.children(t);
  }
  const Su = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Ei,
    name: mu,
    parse: vi,
    structure: ku
  }, Symbol.toStringTag, { value: "Module" })), yu = "SelectorList", bu = "selector", xu = {
    children: [[
      "Selector",
      "Raw"
    ]]
  };
  function $i() {
    const t = this.createList();
    for (; !this.eof; ) {
      if (t.push(this.Selector()), this.tokenType === ft) {
        this.next();
        continue;
      }
      break;
    }
    return {
      type: "SelectorList",
      loc: this.getLocationFromList(t),
      children: t
    };
  }
  function Li(t) {
    this.children(t, () => this.token(ft, ","));
  }
  const Cu = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Li,
    name: yu,
    parse: $i,
    structure: xu,
    walkContext: bu
  }, Symbol.toStringTag, { value: "Module" })), un = 92, Oi = 34, wu = 39;
  function Pi(t) {
    const e = t.length, n = t.charCodeAt(0), s = n === Oi || n === wu ? 1 : 0, r = s === 1 && e > 1 && t.charCodeAt(e - 1) === n ? e - 2 : e - 1;
    let o = "";
    for (let a = s; a <= r; a++) {
      let l = t.charCodeAt(a);
      if (l === un) {
        if (a === r) {
          a !== e - 1 && (o = t.substr(a + 1));
          break;
        }
        if (l = t.charCodeAt(++a), kt(un, l)) {
          const u = a - 1, i = Jt(t, u);
          a = i - 1, o += Ys(t.substring(u + 1, i));
        } else
          l === 13 && t.charCodeAt(a + 1) === 10 && a++;
      } else
        o += t[a];
    }
    return o;
  }
  function Tu(t, e) {
    const s = Oi;
    let r = "", o = !1;
    for (let a = 0; a < t.length; a++) {
      const l = t.charCodeAt(a);
      if (l === 0) {
        r += "�";
        continue;
      }
      if (l <= 31 || l === 127) {
        r += "\\" + l.toString(16), o = !0;
        continue;
      }
      l === s || l === un ? (r += "\\" + t.charAt(a), o = !1) : (o && (Xt(l) || Ft(l)) && (r += " "), r += t.charAt(a), o = !1);
    }
    return '"' + r + '"';
  }
  const Au = "String", vu = {
    value: String
  };
  function Ri() {
    return {
      type: "String",
      loc: this.getLocation(this.tokenStart, this.tokenEnd),
      value: Pi(this.consume(Tt))
    };
  }
  function _i(t) {
    this.token(Tt, Tu(t.value));
  }
  const Eu = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: _i,
    name: Au,
    parse: Ri,
    structure: vu
  }, Symbol.toStringTag, { value: "Module" })), $u = 33;
  function es() {
    return this.Raw(null, !1);
  }
  const Lu = "StyleSheet", Ou = "stylesheet", Pu = {
    children: [[
      "Comment",
      "CDO",
      "CDC",
      "Atrule",
      "Rule",
      "Raw"
    ]]
  };
  function Ii() {
    const t = this.tokenStart, e = this.createList();
    let n;
    for (; !this.eof; ) {
      switch (this.tokenType) {
        case W:
          this.next();
          continue;
        case Y:
          if (this.charCodeAt(this.tokenStart + 2) !== $u) {
            this.next();
            continue;
          }
          n = this.Comment();
          break;
        case Fe:
          n = this.CDO();
          break;
        case nt:
          n = this.CDC();
          break;
        // CSS Syntax Module Level 3
        // §2.2 Error handling
        // At the "top level" of a stylesheet, an <at-keyword-token> starts an at-rule.
        case z:
          n = this.parseWithFallback(this.Atrule, es);
          break;
        // Anything else starts a qualified rule ...
        default:
          n = this.parseWithFallback(this.Rule, es);
      }
      e.push(n);
    }
    return {
      type: "StyleSheet",
      loc: this.getLocation(t, this.tokenStart),
      children: e
    };
  }
  function Ni(t) {
    this.children(t);
  }
  const Ru = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Ni,
    name: Lu,
    parse: Ii,
    structure: Pu,
    walkContext: Ou
  }, Symbol.toStringTag, { value: "Module" })), _u = "SupportsDeclaration", Iu = {
    declaration: "Declaration"
  };
  function Di() {
    const t = this.tokenStart;
    this.eat(_), this.skipSC();
    const e = this.Declaration();
    return this.eof || this.eat(S), {
      type: "SupportsDeclaration",
      loc: this.getLocation(t, this.tokenStart),
      declaration: e
    };
  }
  function Fi(t) {
    this.token(_, "("), this.node(t.declaration), this.token(S, ")");
  }
  const Nu = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Fi,
    name: _u,
    parse: Di,
    structure: Iu
  }, Symbol.toStringTag, { value: "Module" })), Du = 42, ns = 124;
  function qe() {
    this.tokenType !== g && this.isDelim(Du) === !1 && this.error("Identifier or asterisk is expected"), this.next();
  }
  const Fu = "TypeSelector", Mu = {
    name: String
  };
  function Mi() {
    const t = this.tokenStart;
    return this.isDelim(ns) ? (this.next(), qe.call(this)) : (qe.call(this), this.isDelim(ns) && (this.next(), qe.call(this))), {
      type: "TypeSelector",
      loc: this.getLocation(t, this.tokenStart),
      name: this.substrToCursor(t)
    };
  }
  function ji(t) {
    this.tokenize(t.name);
  }
  const ju = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: ji,
    name: Fu,
    parse: Mi,
    structure: Mu
  }, Symbol.toStringTag, { value: "Module" })), Bi = 43, Wi = 45, hn = 63;
  function se(t, e) {
    let n = 0;
    for (let s = this.tokenStart + t; s < this.tokenEnd; s++) {
      const r = this.charCodeAt(s);
      if (r === Wi && e && n !== 0)
        return se.call(this, t + n + 1, !1), -1;
      Xt(r) || this.error(
        e && n !== 0 ? "Hyphen minus" + (n < 6 ? " or hex digit" : "") + " is expected" : n < 6 ? "Hex digit is expected" : "Unexpected input",
        s
      ), ++n > 6 && this.error("Too many hex digits", s);
    }
    return this.next(), n;
  }
  function ke(t) {
    let e = 0;
    for (; this.isDelim(hn); )
      ++e > t && this.error("Too many question marks"), this.next();
  }
  function Bu(t) {
    this.charCodeAt(this.tokenStart) !== t && this.error((t === Bi ? "Plus sign" : "Hyphen minus") + " is expected");
  }
  function Wu() {
    let t = 0;
    switch (this.tokenType) {
      case b:
        if (t = se.call(this, 1, !0), this.isDelim(hn)) {
          ke.call(this, 6 - t);
          break;
        }
        if (this.tokenType === E || this.tokenType === b) {
          Bu.call(this, Wi), se.call(this, 1, !1);
          break;
        }
        break;
      case E:
        t = se.call(this, 1, !0), t > 0 && ke.call(this, 6 - t);
        break;
      default:
        if (this.eatDelim(Bi), this.tokenType === g) {
          t = se.call(this, 0, !0), t > 0 && ke.call(this, 6 - t);
          break;
        }
        if (this.isDelim(hn)) {
          this.next(), ke.call(this, 5);
          break;
        }
        this.error("Hex digit or question mark is expected");
    }
  }
  const Uu = "UnicodeRange", zu = {
    value: String
  };
  function Ui() {
    const t = this.tokenStart;
    return this.eatIdent("u"), Wu.call(this), {
      type: "UnicodeRange",
      loc: this.getLocation(t, this.tokenStart),
      value: this.substrToCursor(t)
    };
  }
  function zi(t) {
    this.tokenize(t.value);
  }
  const Vu = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: zi,
    name: Uu,
    parse: Ui,
    structure: zu
  }, Symbol.toStringTag, { value: "Module" })), Hu = 32, fn = 92, Gu = 34, qu = 39, Ku = 40, Vi = 41;
  function Qu(t) {
    const e = t.length;
    let n = 4, s = t.charCodeAt(e - 1) === Vi ? e - 2 : e - 1, r = "";
    for (; n < s && Ft(t.charCodeAt(n)); )
      n++;
    for (; n < s && Ft(t.charCodeAt(s)); )
      s--;
    for (let o = n; o <= s; o++) {
      let a = t.charCodeAt(o);
      if (a === fn) {
        if (o === s) {
          o !== e - 1 && (r = t.substr(o + 1));
          break;
        }
        if (a = t.charCodeAt(++o), kt(fn, a)) {
          const l = o - 1, u = Jt(t, l);
          o = u - 1, r += Ys(t.substring(l + 1, u));
        } else
          a === 13 && t.charCodeAt(o + 1) === 10 && o++;
      } else
        r += t[o];
    }
    return r;
  }
  function Yu(t) {
    let e = "", n = !1;
    for (let s = 0; s < t.length; s++) {
      const r = t.charCodeAt(s);
      if (r === 0) {
        e += "�";
        continue;
      }
      if (r <= 31 || r === 127) {
        e += "\\" + r.toString(16), n = !0;
        continue;
      }
      r === Hu || r === fn || r === Gu || r === qu || r === Ku || r === Vi ? (e += "\\" + t.charAt(s), n = !1) : (n && Xt(r) && (e += " "), e += t.charAt(s), n = !1);
    }
    return "url(" + e + ")";
  }
  const Xu = "Url", Ju = {
    value: String
  };
  function Hi() {
    const t = this.tokenStart;
    let e;
    switch (this.tokenType) {
      case Q:
        e = Qu(this.consume(Q));
        break;
      case T:
        this.cmpStr(this.tokenStart, this.tokenEnd, "url(") || this.error("Function name must be `url`"), this.eat(T), this.skipSC(), e = Pi(this.consume(Tt)), this.skipSC(), this.eof || this.eat(S);
        break;
      default:
        this.error("Url or Function is expected");
    }
    return {
      type: "Url",
      loc: this.getLocation(t, this.tokenStart),
      value: e
    };
  }
  function Gi(t) {
    this.token(Q, Yu(t.value));
  }
  const Zu = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Gi,
    name: Xu,
    parse: Hi,
    structure: Ju
  }, Symbol.toStringTag, { value: "Module" })), th = "Value", eh = {
    children: [[]]
  };
  function qi() {
    const t = this.tokenStart, e = this.readSequence(this.scope.Value);
    return {
      type: "Value",
      loc: this.getLocation(t, this.tokenStart),
      children: e
    };
  }
  function Ki(t) {
    this.children(t);
  }
  const nh = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Ki,
    name: th,
    parse: qi,
    structure: eh
  }, Symbol.toStringTag, { value: "Module" })), sh = Object.freeze({
    type: "WhiteSpace",
    loc: null,
    value: " "
  }), rh = "WhiteSpace", ih = {
    value: String
  };
  function Qi() {
    return this.eat(W), sh;
  }
  function Yi(t) {
    this.token(W, t.value);
  }
  const oh = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    generate: Yi,
    name: rh,
    parse: Qi,
    structure: ih
  }, Symbol.toStringTag, { value: "Module" })), ah = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    AnPlusB: Sa,
    Atrule: wa,
    AtrulePrelude: Ea,
    AttributeSelector: Na,
    Block: Wa,
    Brackets: Va,
    CDC: qa,
    CDO: Ya,
    ClassSelector: tl,
    Combinator: ol,
    Comment: hl,
    Condition: ml,
    Declaration: Pl,
    DeclarationList: Nl,
    Dimension: Ml,
    Feature: Ul,
    FeatureFunction: Gl,
    FeatureRange: Xl,
    Function: ec,
    GeneralEnclosed: rc,
    Hash: lc,
    IdSelector: dc,
    Identifier: hc,
    Layer: Sc,
    LayerList: xc,
    MediaQuery: Tc,
    MediaQueryList: Ec,
    NestingSelector: Pc,
    Nth: Ic,
    Number: Fc,
    Operator: Bc,
    Parentheses: zc,
    Percentage: Gc,
    PseudoClassSelector: Yc,
    PseudoElementSelector: tu,
    Ratio: su,
    Raw: au,
    Rule: fu,
    Scope: gu,
    Selector: Su,
    SelectorList: Cu,
    String: Eu,
    StyleSheet: Ru,
    SupportsDeclaration: Nu,
    TypeSelector: ju,
    UnicodeRange: Vu,
    Url: Zu,
    Value: nh,
    WhiteSpace: oh
  }, Symbol.toStringTag, { value: "Module" })), lh = {
    node: ah
  }, $t = na(lh), Tn = [
    "left",
    "right",
    "top",
    "bottom",
    "inset-block-start",
    "inset-block-end",
    "inset-inline-start",
    "inset-inline-end",
    "inset-block",
    "inset-inline",
    "inset"
  ];
  function ue(t) {
    return Tn.includes(t);
  }
  const An = [
    "margin-block-start",
    "margin-block-end",
    "margin-block",
    "margin-inline-start",
    "margin-inline-end",
    "margin-inline",
    "margin-bottom",
    "margin-left",
    "margin-right",
    "margin-top",
    "margin"
  ];
  function ch(t) {
    return An.includes(t);
  }
  const vn = [
    "width",
    "height",
    "min-width",
    "min-height",
    "max-width",
    "max-height",
    "block-size",
    "inline-size",
    "min-block-size",
    "min-inline-size",
    "max-block-size",
    "max-inline-size"
  ];
  function Xi(t) {
    return vn.includes(t);
  }
  const Ji = [
    "justify-self",
    "align-self",
    "place-self"
  ];
  function uh(t) {
    return Ji.includes(t);
  }
  const Zi = [
    ...Tn,
    ...An,
    ...vn,
    ...Ji,
    "position-anchor",
    "position-area"
  ], hh = [
    ...vn,
    ...Tn,
    ...An
  ];
  function to(t) {
    return hh.includes(
      t
    );
  }
  const fh = [
    "top",
    "left",
    "right",
    "bottom",
    "start",
    "end",
    "self-start",
    "self-end",
    "center",
    "inside",
    "outside"
  ];
  function eo(t) {
    return fh.includes(t);
  }
  const ph = [
    "width",
    "height",
    "block",
    "inline",
    "self-block",
    "self-inline"
  ];
  function dh(t) {
    return ph.includes(t);
  }
  const ss = /* @__PURE__ */ new Set(["Atrule", "Selector", "Declaration"]);
  function gh(t) {
    const e = new SourceMapGenerator(), n = {
      line: 1,
      column: 0
    }, s = {
      line: 0,
      // should be zero to add first mapping
      column: 0
    }, r = {
      line: 1,
      column: 0
    }, o = {
      generated: r
    };
    let a = 1, l = 0, u = !1;
    const i = t.node;
    t.node = function(f) {
      if (f.loc && f.loc.start && ss.has(f.type)) {
        const p = f.loc.start.line, d = f.loc.start.column - 1;
        (s.line !== p || s.column !== d) && (s.line = p, s.column = d, n.line = a, n.column = l, u && (u = !1, (n.line !== r.line || n.column !== r.column) && e.addMapping(o)), u = !0, e.addMapping({
          source: f.loc.source,
          original: s,
          generated: n
        }));
      }
      i.call(this, f), u && ss.has(f.type) && (r.line = a, r.column = l);
    };
    const c = t.emit;
    t.emit = function(f, p, d) {
      for (let m = 0; m < f.length; m++)
        f.charCodeAt(m) === 10 ? (a++, l = 0) : l++;
      c(f, p, d);
    };
    const h = t.result;
    return t.result = function() {
      return u && e.addMapping(o), {
        css: h(),
        map: e
      };
    }, t;
  }
  const mh = 43, kh = 45, Ke = (t, e) => {
    if (t === L && (t = e), typeof t == "string") {
      const n = t.charCodeAt(0);
      return n > 127 ? 32768 : n << 8;
    }
    return t;
  }, no = [
    [g, g],
    [g, T],
    [g, Q],
    [g, at],
    [g, "-"],
    [g, b],
    [g, F],
    [g, E],
    [g, nt],
    [g, _],
    [z, g],
    [z, T],
    [z, Q],
    [z, at],
    [z, "-"],
    [z, b],
    [z, F],
    [z, E],
    [z, nt],
    [I, g],
    [I, T],
    [I, Q],
    [I, at],
    [I, "-"],
    [I, b],
    [I, F],
    [I, E],
    [I, nt],
    [E, g],
    [E, T],
    [E, Q],
    [E, at],
    [E, "-"],
    [E, b],
    [E, F],
    [E, E],
    [E, nt],
    ["#", g],
    ["#", T],
    ["#", Q],
    ["#", at],
    ["#", "-"],
    ["#", b],
    ["#", F],
    ["#", E],
    ["#", nt],
    // https://github.com/w3c/csswg-drafts/pull/6874
    ["-", g],
    ["-", T],
    ["-", Q],
    ["-", at],
    ["-", "-"],
    ["-", b],
    ["-", F],
    ["-", E],
    ["-", nt],
    // https://github.com/w3c/csswg-drafts/pull/6874
    [b, g],
    [b, T],
    [b, Q],
    [b, at],
    [b, b],
    [b, F],
    [b, E],
    [b, "%"],
    [b, nt],
    // https://github.com/w3c/csswg-drafts/pull/6874
    ["@", g],
    ["@", T],
    ["@", Q],
    ["@", at],
    ["@", "-"],
    ["@", nt],
    // https://github.com/w3c/csswg-drafts/pull/6874
    [".", b],
    [".", F],
    [".", E],
    ["+", b],
    ["+", F],
    ["+", E],
    ["/", "*"]
  ], Sh = no.concat([
    [g, I],
    [E, I],
    [I, I],
    [z, _],
    [z, Tt],
    [z, X],
    [F, F],
    [F, E],
    [F, T],
    [F, "-"],
    [S, g],
    [S, T],
    [S, F],
    [S, E],
    [S, I],
    [S, "-"]
  ]);
  function so(t) {
    const e = new Set(
      t.map(([n, s]) => Ke(n) << 16 | Ke(s))
    );
    return function(n, s, r) {
      const o = Ke(s, r), a = r.charCodeAt(0);
      return (a === kh && s !== g && s !== T && s !== nt || a === mh ? e.has(n << 16 | a << 8) : e.has(n << 16 | o)) && this.emit(" ", W, !0), o;
    };
  }
  const yh = so(no), ro = so(Sh), rs = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    safe: ro,
    spec: yh
  }, Symbol.toStringTag, { value: "Module" })), bh = 92;
  function xh(t, e) {
    if (typeof e == "function") {
      let n = null;
      t.children.forEach((s) => {
        n !== null && e.call(this, n), this.node(s), n = s;
      });
      return;
    }
    t.children.forEach(this.node, this);
  }
  function Ch(t) {
    Js(t, (e, n, s) => {
      this.token(e, t.slice(n, s));
    });
  }
  function wh(t) {
    const e = /* @__PURE__ */ new Map();
    for (let [n, s] of Object.entries(t.node))
      typeof (s.generate || s) == "function" && e.set(n, s.generate || s);
    return function(n, s) {
      let r = "", o = 0, a = {
        node(u) {
          if (e.has(u.type))
            e.get(u.type).call(l, u);
          else
            throw new Error("Unknown node type: " + u.type);
        },
        tokenBefore: ro,
        token(u, i) {
          o = this.tokenBefore(o, u, i), this.emit(i, u, !1), u === L && i.charCodeAt(0) === bh && this.emit(`
`, W, !0);
        },
        emit(u) {
          r += u;
        },
        result() {
          return r;
        }
      };
      s && (typeof s.decorator == "function" && (a = s.decorator(a)), s.sourceMap && (a = gh(a)), s.mode in rs && (a.tokenBefore = rs[s.mode]));
      const l = {
        node: (u) => a.node(u),
        children: xh,
        token: (u, i) => a.token(u, i),
        tokenize: Ch
      };
      return a.node(n), a.result();
    };
  }
  const Th = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    AnPlusB: tr,
    Atrule: nr,
    AtrulePrelude: rr,
    AttributeSelector: ar,
    Block: ur,
    Brackets: fr,
    CDC: dr,
    CDO: mr,
    ClassSelector: Sr,
    Combinator: br,
    Comment: Cr,
    Condition: Tr,
    Declaration: Er,
    DeclarationList: Lr,
    Dimension: Pr,
    Feature: _r,
    FeatureFunction: Nr,
    FeatureRange: Fr,
    Function: jr,
    GeneralEnclosed: Wr,
    Hash: zr,
    IdSelector: qr,
    Identifier: Hr,
    Layer: Qr,
    LayerList: Xr,
    MediaQuery: Zr,
    MediaQueryList: ei,
    NestingSelector: si,
    Nth: ii,
    Number: ai,
    Operator: ci,
    Parentheses: hi,
    Percentage: pi,
    PseudoClassSelector: gi,
    PseudoElementSelector: ki,
    Ratio: yi,
    Raw: xi,
    Rule: wi,
    Scope: Ai,
    Selector: Ei,
    SelectorList: Li,
    String: _i,
    StyleSheet: Ni,
    SupportsDeclaration: Fi,
    TypeSelector: ji,
    UnicodeRange: zi,
    Url: Gi,
    Value: Ki,
    WhiteSpace: Yi
  }, Symbol.toStringTag, { value: "Module" })), Ah = {
    node: Th
  }, vh = wh(Ah);
  let zt = null;
  class K {
    static createItem(e) {
      return {
        prev: null,
        next: null,
        data: e
      };
    }
    constructor() {
      this.head = null, this.tail = null, this.cursor = null;
    }
    createItem(e) {
      return K.createItem(e);
    }
    // cursor helpers
    allocateCursor(e, n) {
      let s;
      return zt !== null ? (s = zt, zt = zt.cursor, s.prev = e, s.next = n, s.cursor = this.cursor) : s = {
        prev: e,
        next: n,
        cursor: this.cursor
      }, this.cursor = s, s;
    }
    releaseCursor() {
      const { cursor: e } = this;
      this.cursor = e.cursor, e.prev = null, e.next = null, e.cursor = zt, zt = e;
    }
    updateCursors(e, n, s, r) {
      let { cursor: o } = this;
      for (; o !== null; )
        o.prev === e && (o.prev = n), o.next === s && (o.next = r), o = o.cursor;
    }
    *[Symbol.iterator]() {
      for (let e = this.head; e !== null; e = e.next)
        yield e.data;
    }
    // getters
    get size() {
      let e = 0;
      for (let n = this.head; n !== null; n = n.next)
        e++;
      return e;
    }
    get isEmpty() {
      return this.head === null;
    }
    get first() {
      return this.head && this.head.data;
    }
    get last() {
      return this.tail && this.tail.data;
    }
    // convertors
    fromArray(e) {
      let n = null;
      this.head = null;
      for (let s of e) {
        const r = K.createItem(s);
        n !== null ? n.next = r : this.head = r, r.prev = n, n = r;
      }
      return this.tail = n, this;
    }
    toArray() {
      return [...this];
    }
    toJSON() {
      return [...this];
    }
    // array-like methods
    forEach(e, n = this) {
      const s = this.allocateCursor(null, this.head);
      for (; s.next !== null; ) {
        const r = s.next;
        s.next = r.next, e.call(n, r.data, r, this);
      }
      this.releaseCursor();
    }
    forEachRight(e, n = this) {
      const s = this.allocateCursor(this.tail, null);
      for (; s.prev !== null; ) {
        const r = s.prev;
        s.prev = r.prev, e.call(n, r.data, r, this);
      }
      this.releaseCursor();
    }
    reduce(e, n, s = this) {
      let r = this.allocateCursor(null, this.head), o = n, a;
      for (; r.next !== null; )
        a = r.next, r.next = a.next, o = e.call(s, o, a.data, a, this);
      return this.releaseCursor(), o;
    }
    reduceRight(e, n, s = this) {
      let r = this.allocateCursor(this.tail, null), o = n, a;
      for (; r.prev !== null; )
        a = r.prev, r.prev = a.prev, o = e.call(s, o, a.data, a, this);
      return this.releaseCursor(), o;
    }
    some(e, n = this) {
      for (let s = this.head; s !== null; s = s.next)
        if (e.call(n, s.data, s, this))
          return !0;
      return !1;
    }
    map(e, n = this) {
      const s = new K();
      for (let r = this.head; r !== null; r = r.next)
        s.appendData(e.call(n, r.data, r, this));
      return s;
    }
    filter(e, n = this) {
      const s = new K();
      for (let r = this.head; r !== null; r = r.next)
        e.call(n, r.data, r, this) && s.appendData(r.data);
      return s;
    }
    nextUntil(e, n, s = this) {
      if (e === null)
        return;
      const r = this.allocateCursor(null, e);
      for (; r.next !== null; ) {
        const o = r.next;
        if (r.next = o.next, n.call(s, o.data, o, this))
          break;
      }
      this.releaseCursor();
    }
    prevUntil(e, n, s = this) {
      if (e === null)
        return;
      const r = this.allocateCursor(e, null);
      for (; r.prev !== null; ) {
        const o = r.prev;
        if (r.prev = o.prev, n.call(s, o.data, o, this))
          break;
      }
      this.releaseCursor();
    }
    // mutation
    clear() {
      this.head = null, this.tail = null;
    }
    copy() {
      const e = new K();
      for (let n of this)
        e.appendData(n);
      return e;
    }
    prepend(e) {
      return this.updateCursors(null, e, this.head, e), this.head !== null ? (this.head.prev = e, e.next = this.head) : this.tail = e, this.head = e, this;
    }
    prependData(e) {
      return this.prepend(K.createItem(e));
    }
    append(e) {
      return this.insert(e);
    }
    appendData(e) {
      return this.insert(K.createItem(e));
    }
    insert(e, n = null) {
      if (n !== null)
        if (this.updateCursors(n.prev, e, n, e), n.prev === null) {
          if (this.head !== n)
            throw new Error("before doesn't belong to list");
          this.head = e, n.prev = e, e.next = n, this.updateCursors(null, e);
        } else
          n.prev.next = e, e.prev = n.prev, n.prev = e, e.next = n;
      else
        this.updateCursors(this.tail, e, null, e), this.tail !== null ? (this.tail.next = e, e.prev = this.tail) : this.head = e, this.tail = e;
      return this;
    }
    insertData(e, n) {
      return this.insert(K.createItem(e), n);
    }
    remove(e) {
      if (this.updateCursors(e, e.prev, e, e.next), e.prev !== null)
        e.prev.next = e.next;
      else {
        if (this.head !== e)
          throw new Error("item doesn't belong to list");
        this.head = e.next;
      }
      if (e.next !== null)
        e.next.prev = e.prev;
      else {
        if (this.tail !== e)
          throw new Error("item doesn't belong to list");
        this.tail = e.prev;
      }
      return e.prev = null, e.next = null, e;
    }
    push(e) {
      this.insert(K.createItem(e));
    }
    pop() {
      return this.tail !== null ? this.remove(this.tail) : null;
    }
    unshift(e) {
      this.prepend(K.createItem(e));
    }
    shift() {
      return this.head !== null ? this.remove(this.head) : null;
    }
    prependList(e) {
      return this.insertList(e, this.head);
    }
    appendList(e) {
      return this.insertList(e);
    }
    insertList(e, n) {
      return e.head === null ? this : (n != null ? (this.updateCursors(n.prev, e.tail, n, e.head), n.prev !== null ? (n.prev.next = e.head, e.head.prev = n.prev) : this.head = e.head, n.prev = e.tail, e.tail.next = n) : (this.updateCursors(this.tail, e.tail, null, e.head), this.tail !== null ? (this.tail.next = e.head, e.head.prev = this.tail) : this.head = e.head, this.tail = e.tail), e.head = null, e.tail = null, this);
    }
    replace(e, n) {
      "head" in n ? this.insertList(n, e) : this.insert(n, e), this.remove(e);
    }
  }
  function Eh(t, e) {
    const n = Object.create(SyntaxError.prototype), s = new Error();
    return Object.assign(n, {
      name: t,
      message: e,
      get stack() {
        return (s.stack || "").replace(/^(.+\n){1,3}/, `${t}: ${e}
`);
      }
    });
  }
  const Qe = 100, is = 60, os = "    ";
  function as({ source: t, line: e, column: n, baseLine: s, baseColumn: r }, o) {
    function a(d, m) {
      return i.slice(d, m).map(
        (k, y) => String(d + y + 1).padStart(f) + " |" + k
      ).join(`
`);
    }
    const l = `
`.repeat(Math.max(s - 1, 0)), u = " ".repeat(Math.max(r - 1, 0)), i = (l + u + t).split(/\r\n?|\n|\f/), c = Math.max(1, e - o) - 1, h = Math.min(e + o, i.length + 1), f = Math.max(4, String(h).length) + 1;
    let p = 0;
    n += (os.length - 1) * (i[e - 1].substr(0, n - 1).match(/\t/g) || []).length, n > Qe && (p = n - is + 3, n = is - 2);
    for (let d = c; d <= h; d++)
      d >= 0 && d < i.length && (i[d] = i[d].replace(/\t/g, os), i[d] = (p > 0 && i[d].length > p ? "…" : "") + i[d].substr(p, Qe - 2) + (i[d].length > p + Qe - 1 ? "…" : ""));
    return [
      a(c, e),
      new Array(n + f + 2).join("-") + "^",
      a(e, h)
    ].filter(Boolean).join(`
`).replace(/^(\s+\d+\s+\|\n)+/, "").replace(/\n(\s+\d+\s+\|)+$/, "");
  }
  function ls(t, e, n, s, r, o = 1, a = 1) {
    return Object.assign(Eh("SyntaxError", t), {
      source: e,
      offset: n,
      line: s,
      column: r,
      sourceFragment(u) {
        return as({ source: e, line: s, column: r, baseLine: o, baseColumn: a }, isNaN(u) ? 0 : u);
      },
      get formattedMessage() {
        return `Parse error: ${t}
` + as({ source: e, line: s, column: r, baseLine: o, baseColumn: a }, 2);
      }
    });
  }
  function $h(t) {
    const e = this.createList();
    let n = !1;
    const s = {
      recognizer: t
    };
    for (; !this.eof; ) {
      switch (this.tokenType) {
        case Y:
          this.next();
          continue;
        case W:
          n = !0, this.next();
          continue;
      }
      let r = t.getNode.call(this, s);
      if (r === void 0)
        break;
      n && (t.onWhiteSpace && t.onWhiteSpace.call(this, r, e, s), n = !1), e.push(r);
    }
    return n && t.onWhiteSpace && t.onWhiteSpace.call(this, null, e, s), e;
  }
  const cs = () => {
  }, Lh = 33, Oh = 35, Ye = 59, us = 123, hs = 0;
  function Ph(t) {
    return function() {
      return this[t]();
    };
  }
  function Xe(t) {
    const e = /* @__PURE__ */ Object.create(null);
    for (const n of Object.keys(t)) {
      const s = t[n], r = s.parse || s;
      r && (e[n] = r);
    }
    return e;
  }
  function Rh(t) {
    const e = {
      context: /* @__PURE__ */ Object.create(null),
      features: Object.assign(/* @__PURE__ */ Object.create(null), t.features),
      scope: Object.assign(/* @__PURE__ */ Object.create(null), t.scope),
      atrule: Xe(t.atrule),
      pseudo: Xe(t.pseudo),
      node: Xe(t.node)
    };
    for (const [n, s] of Object.entries(t.parseContext))
      switch (typeof s) {
        case "function":
          e.context[n] = s;
          break;
        case "string":
          e.context[n] = Ph(s);
          break;
      }
    return U(U({
      config: e
    }, e), e.node);
  }
  function _h(t) {
    let e = "", n = "<unknown>", s = !1, r = cs, o = !1;
    const a = new pa(), l = Object.assign(new da(), Rh(t || {}), {
      parseAtrulePrelude: !0,
      parseRulePrelude: !0,
      parseValue: !0,
      parseCustomProperty: !1,
      readSequence: $h,
      consumeUntilBalanceEnd: () => 0,
      consumeUntilLeftCurlyBracket(i) {
        return i === us ? 1 : 0;
      },
      consumeUntilLeftCurlyBracketOrSemicolon(i) {
        return i === us || i === Ye ? 1 : 0;
      },
      consumeUntilExclamationMarkOrSemicolon(i) {
        return i === Lh || i === Ye ? 1 : 0;
      },
      consumeUntilSemicolonIncluded(i) {
        return i === Ye ? 2 : 0;
      },
      createList() {
        return new K();
      },
      createSingleNodeList(i) {
        return new K().appendData(i);
      },
      getFirstListNode(i) {
        return i && i.first;
      },
      getLastListNode(i) {
        return i && i.last;
      },
      parseWithFallback(i, c) {
        const h = this.tokenIndex;
        try {
          return i.call(this);
        } catch (f) {
          if (o)
            throw f;
          this.skip(h - this.tokenIndex);
          const p = c.call(this);
          return o = !0, r(f, p), o = !1, p;
        }
      },
      lookupNonWSType(i) {
        let c;
        do
          if (c = this.lookupType(i++), c !== W && c !== Y)
            return c;
        while (c !== hs);
        return hs;
      },
      charCodeAt(i) {
        return i >= 0 && i < e.length ? e.charCodeAt(i) : 0;
      },
      substring(i, c) {
        return e.substring(i, c);
      },
      substrToCursor(i) {
        return this.source.substring(i, this.tokenStart);
      },
      cmpChar(i, c) {
        return Ks(e, i, c);
      },
      cmpStr(i, c, h) {
        return Le(e, i, c, h);
      },
      consume(i) {
        const c = this.tokenStart;
        return this.eat(i), this.substrToCursor(c);
      },
      consumeFunctionName() {
        const i = e.substring(this.tokenStart, this.tokenEnd - 1);
        return this.eat(T), i;
      },
      consumeNumber(i) {
        const c = e.substring(this.tokenStart, Qs(e, this.tokenStart));
        return this.eat(i), c;
      },
      eat(i) {
        if (this.tokenType !== i) {
          const c = Xs[i].slice(0, -6).replace(/-/g, " ").replace(/^./, (p) => p.toUpperCase());
          let h = `${/[[\](){}]/.test(c) ? `"${c}"` : c} is expected`, f = this.tokenStart;
          switch (i) {
            case g:
              this.tokenType === T || this.tokenType === Q ? (f = this.tokenEnd - 1, h = "Identifier is expected but function found") : h = "Identifier is expected";
              break;
            case I:
              this.isDelim(Oh) && (this.next(), f++, h = "Name is expected");
              break;
            case F:
              this.tokenType === b && (f = this.tokenEnd, h = "Percent sign is expected");
              break;
          }
          this.error(h, f);
        }
        this.next();
      },
      eatIdent(i) {
        (this.tokenType !== g || this.lookupValue(0, i) === !1) && this.error(`Identifier "${i}" is expected`), this.next();
      },
      eatDelim(i) {
        this.isDelim(i) || this.error(`Delim "${String.fromCharCode(i)}" is expected`), this.next();
      },
      getLocation(i, c) {
        return s ? a.getLocationRange(
          i,
          c,
          n
        ) : null;
      },
      getLocationFromList(i) {
        if (s) {
          const c = this.getFirstListNode(i), h = this.getLastListNode(i);
          return a.getLocationRange(
            c !== null ? c.loc.start.offset - a.startOffset : this.tokenStart,
            h !== null ? h.loc.end.offset - a.startOffset : this.tokenStart,
            n
          );
        }
        return null;
      },
      error(i, c) {
        const h = typeof c != "undefined" && c < e.length ? a.getLocation(c) : this.eof ? a.getLocation(ua(e, e.length - 1)) : a.getLocation(this.tokenStart);
        throw new ls(
          i || "Unexpected input",
          e,
          h.offset,
          h.line,
          h.column,
          a.startLine,
          a.startColumn
        );
      }
    });
    return Object.assign(function(i, c) {
      e = i, c = c || {}, l.setSource(e, Js), a.setSource(
        e,
        c.offset,
        c.line,
        c.column
      ), n = c.filename || "<unknown>", s = !!c.positions, r = typeof c.onParseError == "function" ? c.onParseError : cs, o = !1, l.parseAtrulePrelude = "parseAtrulePrelude" in c ? !!c.parseAtrulePrelude : !0, l.parseRulePrelude = "parseRulePrelude" in c ? !!c.parseRulePrelude : !0, l.parseValue = "parseValue" in c ? !!c.parseValue : !0, l.parseCustomProperty = "parseCustomProperty" in c ? !!c.parseCustomProperty : !1;
      const { context: h = "default", onComment: f } = c;
      if (!(h in l.context))
        throw new Error("Unknown context `" + h + "`");
      typeof f == "function" && l.forEachToken((d, m, k) => {
        if (d === Y) {
          const y = l.getLocation(m, k), x = Le(e, k - 2, k, "*/") ? e.slice(m + 2, k - 2) : e.slice(m + 2, k);
          f(x, y);
        }
      });
      const p = l.context[h].call(l, c);
      return l.eof || l.error(), p;
    }, {
      SyntaxError: ls,
      config: l.config
    });
  }
  const Ih = 35, Nh = 42, fs = 43, Dh = 45, Fh = 47, Mh = 117;
  function io(t) {
    switch (this.tokenType) {
      case I:
        return this.Hash();
      case ft:
        return this.Operator();
      case _:
        return this.Parentheses(this.readSequence, t.recognizer);
      case ee:
        return this.Brackets(this.readSequence, t.recognizer);
      case Tt:
        return this.String();
      case E:
        return this.Dimension();
      case F:
        return this.Percentage();
      case b:
        return this.Number();
      case T:
        return this.cmpStr(this.tokenStart, this.tokenEnd, "url(") ? this.Url() : this.Function(this.readSequence, t.recognizer);
      case Q:
        return this.Url();
      case g:
        return this.cmpChar(this.tokenStart, Mh) && this.cmpChar(this.tokenStart + 1, fs) ? this.UnicodeRange() : this.Identifier();
      case L: {
        const e = this.charCodeAt(this.tokenStart);
        if (e === Fh || e === Nh || e === fs || e === Dh)
          return this.Operator();
        e === Ih && this.error("Hex or identifier is expected", this.tokenStart + 1);
        break;
      }
    }
  }
  const jh = {
    getNode: io
  }, Bh = 35, Wh = 38, Uh = 42, zh = 43, Vh = 47, ps = 46, Hh = 62, Gh = 124, qh = 126;
  function Kh(t, e) {
    e.last !== null && e.last.type !== "Combinator" && t !== null && t.type !== "Combinator" && e.push({
      // FIXME: this.Combinator() should be used instead
      type: "Combinator",
      loc: null,
      name: " "
    });
  }
  function Qh() {
    switch (this.tokenType) {
      case ee:
        return this.AttributeSelector();
      case I:
        return this.IdSelector();
      case X:
        return this.lookupType(1) === X ? this.PseudoElementSelector() : this.PseudoClassSelector();
      case g:
        return this.TypeSelector();
      case b:
      case F:
        return this.Percentage();
      case E:
        this.charCodeAt(this.tokenStart) === ps && this.error("Identifier is expected", this.tokenStart + 1);
        break;
      case L: {
        switch (this.charCodeAt(this.tokenStart)) {
          case zh:
          case Hh:
          case qh:
          case Vh:
            return this.Combinator();
          case ps:
            return this.ClassSelector();
          case Uh:
          case Gh:
            return this.TypeSelector();
          case Bh:
            return this.IdSelector();
          case Wh:
            return this.NestingSelector();
        }
        break;
      }
    }
  }
  const Yh = {
    onWhiteSpace: Kh,
    getNode: Qh
  };
  function Xh() {
    return this.createSingleNodeList(
      this.Raw(null, !1)
    );
  }
  function Jh() {
    const t = this.createList();
    if (this.skipSC(), t.push(this.Identifier()), this.skipSC(), this.tokenType === ft) {
      t.push(this.Operator());
      const e = this.tokenIndex, n = this.parseCustomProperty ? this.Value(null) : this.Raw(this.consumeUntilExclamationMarkOrSemicolon, !1);
      if (n.type === "Value" && n.children.isEmpty) {
        for (let s = e - this.tokenIndex; s <= 0; s++)
          if (this.lookupType(s) === W) {
            n.children.appendData({
              type: "WhiteSpace",
              loc: null,
              value: " "
            });
            break;
          }
      }
      t.push(n);
    }
    return t;
  }
  function ds(t) {
    return t !== null && t.type === "Operator" && (t.value[t.value.length - 1] === "-" || t.value[t.value.length - 1] === "+");
  }
  const Zh = {
    getNode: io,
    onWhiteSpace(t, e) {
      ds(t) && (t.value = " " + t.value), ds(e.last) && (e.last.value += " ");
    },
    expression: Xh,
    var: Jh
  }, tf = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    AtrulePrelude: jh,
    Selector: Yh,
    Value: Zh
  }, Symbol.toStringTag, { value: "Module" })), ef = /* @__PURE__ */ new Set(["none", "and", "not", "or"]), nf = {
    parse: {
      prelude() {
        const t = this.createList();
        if (this.tokenType === g) {
          const e = this.substring(this.tokenStart, this.tokenEnd);
          ef.has(e.toLowerCase()) || t.push(this.Identifier());
        }
        return t.push(this.Condition("container")), t;
      },
      block(t = !1) {
        return this.Block(t);
      }
    }
  }, sf = {
    parse: {
      prelude: null,
      block() {
        return this.Block(!0);
      }
    }
  };
  function Je(t, e) {
    return this.parseWithFallback(
      () => {
        try {
          return t.call(this);
        } finally {
          this.skipSC(), this.lookupNonWSType(0) !== S && this.error();
        }
      },
      e || (() => this.Raw(null, !0))
    );
  }
  const gs = {
    layer() {
      this.skipSC();
      const t = this.createList(), e = Je.call(this, this.Layer);
      return (e.type !== "Raw" || e.value !== "") && t.push(e), t;
    },
    supports() {
      this.skipSC();
      const t = this.createList(), e = Je.call(
        this,
        this.Declaration,
        () => Je.call(this, () => this.Condition("supports"))
      );
      return (e.type !== "Raw" || e.value !== "") && t.push(e), t;
    }
  }, rf = {
    parse: {
      prelude() {
        const t = this.createList();
        switch (this.tokenType) {
          case Tt:
            t.push(this.String());
            break;
          case Q:
          case T:
            t.push(this.Url());
            break;
          default:
            this.error("String or url() is expected");
        }
        return this.skipSC(), this.tokenType === g && this.cmpStr(this.tokenStart, this.tokenEnd, "layer") ? t.push(this.Identifier()) : this.tokenType === T && this.cmpStr(this.tokenStart, this.tokenEnd, "layer(") && t.push(this.Function(null, gs)), this.skipSC(), this.tokenType === T && this.cmpStr(this.tokenStart, this.tokenEnd, "supports(") && t.push(this.Function(null, gs)), (this.lookupNonWSType(0) === g || this.lookupNonWSType(0) === _) && t.push(this.MediaQueryList()), t;
      },
      block: null
    }
  }, of = {
    parse: {
      prelude() {
        return this.createSingleNodeList(
          this.LayerList()
        );
      },
      block() {
        return this.Block(!1);
      }
    }
  }, af = {
    parse: {
      prelude() {
        return this.createSingleNodeList(
          this.MediaQueryList()
        );
      },
      block(t = !1) {
        return this.Block(t);
      }
    }
  }, lf = {
    parse: {
      prelude() {
        return this.createSingleNodeList(
          this.SelectorList()
        );
      },
      block() {
        return this.Block(!0);
      }
    }
  }, cf = {
    parse: {
      prelude() {
        return this.createSingleNodeList(
          this.SelectorList()
        );
      },
      block() {
        return this.Block(!0);
      }
    }
  }, uf = {
    parse: {
      prelude() {
        return this.createSingleNodeList(
          this.Scope()
        );
      },
      block(t = !1) {
        return this.Block(t);
      }
    }
  }, hf = {
    parse: {
      prelude: null,
      block(t = !1) {
        return this.Block(t);
      }
    }
  }, ff = {
    parse: {
      prelude() {
        return this.createSingleNodeList(
          this.Condition("supports")
        );
      },
      block(t = !1) {
        return this.Block(t);
      }
    }
  }, pf = {
    container: nf,
    "font-face": sf,
    import: rf,
    layer: of,
    media: af,
    nest: lf,
    page: cf,
    scope: uf,
    "starting-style": hf,
    supports: ff
  };
  function df() {
    const t = this.createList();
    this.skipSC();
    t: for (; !this.eof; ) {
      switch (this.tokenType) {
        case g:
          t.push(this.Identifier());
          break;
        case Tt:
          t.push(this.String());
          break;
        case ft:
          t.push(this.Operator());
          break;
        case S:
          break t;
        default:
          this.error("Identifier, string or comma is expected");
      }
      this.skipSC();
    }
    return t;
  }
  const Rt = {
    parse() {
      return this.createSingleNodeList(
        this.SelectorList()
      );
    }
  }, Ze = {
    parse() {
      return this.createSingleNodeList(
        this.Selector()
      );
    }
  }, gf = {
    parse() {
      return this.createSingleNodeList(
        this.Identifier()
      );
    }
  }, mf = {
    parse: df
  }, Se = {
    parse() {
      return this.createSingleNodeList(
        this.Nth()
      );
    }
  }, kf = {
    dir: gf,
    has: Rt,
    lang: mf,
    matches: Rt,
    is: Rt,
    "-moz-any": Rt,
    "-webkit-any": Rt,
    where: Rt,
    not: Rt,
    "nth-child": Se,
    "nth-last-child": Se,
    "nth-last-of-type": Se,
    "nth-of-type": Se,
    slotted: Ze,
    host: Ze,
    "host-context": Ze
  }, Sf = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
    __proto__: null,
    AnPlusB: Zs,
    Atrule: er,
    AtrulePrelude: sr,
    AttributeSelector: or,
    Block: cr,
    Brackets: hr,
    CDC: pr,
    CDO: gr,
    ClassSelector: kr,
    Combinator: yr,
    Comment: xr,
    Condition: wr,
    Declaration: vr,
    DeclarationList: $r,
    Dimension: Or,
    Feature: Rr,
    FeatureFunction: Ir,
    FeatureRange: Dr,
    Function: Mr,
    GeneralEnclosed: Br,
    Hash: Ur,
    IdSelector: Gr,
    Identifier: Vr,
    Layer: Kr,
    LayerList: Yr,
    MediaQuery: Jr,
    MediaQueryList: ti,
    NestingSelector: ni,
    Nth: ri,
    Number: oi,
    Operator: li,
    Parentheses: ui,
    Percentage: fi,
    PseudoClassSelector: di,
    PseudoElementSelector: mi,
    Ratio: Si,
    Raw: bi,
    Rule: Ci,
    Scope: Ti,
    Selector: vi,
    SelectorList: $i,
    String: Ri,
    StyleSheet: Ii,
    SupportsDeclaration: Di,
    TypeSelector: Mi,
    UnicodeRange: Ui,
    Url: Hi,
    Value: qi,
    WhiteSpace: Qi
  }, Symbol.toStringTag, { value: "Module" })), yf = {
    parseContext: {
      default: "StyleSheet",
      stylesheet: "StyleSheet",
      atrule: "Atrule",
      atrulePrelude(t) {
        return this.AtrulePrelude(t.atrule ? String(t.atrule) : null);
      },
      mediaQueryList: "MediaQueryList",
      mediaQuery: "MediaQuery",
      condition(t) {
        return this.Condition(t.kind);
      },
      rule: "Rule",
      selectorList: "SelectorList",
      selector: "Selector",
      block() {
        return this.Block(!0);
      },
      declarationList: "DeclarationList",
      declaration: "Declaration",
      value: "Value"
    },
    features: {
      supports: {
        selector() {
          return this.Selector();
        }
      },
      container: {
        style() {
          return this.Declaration();
        }
      }
    },
    scope: tf,
    atrule: pf,
    pseudo: kf,
    node: Sf
  }, bf = _h(yf);
  function Pe(t) {
    const e = {};
    for (const n of Object.keys(t)) {
      let s = t[n];
      s && (Array.isArray(s) || s instanceof K ? s = s.map(Pe) : s.constructor === Object && (s = Pe(s))), e[n] = s;
    }
    return e;
  }
  let xf = "useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict", lt = (t = 21) => {
    let e = "", n = t | 0;
    for (; n--; )
      e += xf[Math.random() * 64 | 0];
    return e;
  };
  const oo = lt(), re = /* @__PURE__ */ new Set();
  function Re(t) {
    return !!(t && t.type === "Function" && t.name === "anchor");
  }
  function Lt(t, e = !1) {
    return bf(t, {
      parseAtrulePrelude: !1,
      parseCustomProperty: !0,
      onParseError: (n) => {
        e && re.add(n);
      }
    });
  }
  function Z(t) {
    return vh(t, {
      // Default `safe` adds extra (potentially breaking) spaces for compatibility
      // with old browsers.
      mode: "spec"
    });
  }
  function Cf(t) {
    return t.type === "Declaration";
  }
  function wf(t) {
    return t.toArray().reduce(
      (e, n) => n.type === "Operator" && n.value === "," ? (e.push([]), e) : (n.type === "Identifier" && e[e.length - 1].push(n), e),
      [[]]
    );
  }
  function pn(t) {
    return t ? t.children.map((e) => {
      var r;
      let n;
      ((r = e.children.last) == null ? void 0 : r.type) === "PseudoElementSelector" && (e = Pe(e), n = Z(e.children.last), e.children.pop());
      const s = Z(e);
      return {
        selector: s + (n != null ? n : ""),
        elementPart: s,
        pseudoElementPart: n
      };
    }).toArray() : [];
  }
  function Tf() {
    re.size > 0 && (console.group(
      `The CSS anchor positioning polyfill was not applied due to ${re.size === 1 ? "a CSS parse error" : "CSS parse errors"}.`
    ), re.forEach((t) => {
      console.warn(t.formattedMessage);
    }), console.groupEnd());
  }
  function Af() {
    re.clear();
  }
  const dn = [
    ...Zi,
    "anchor-scope",
    "anchor-name"
  ].reduce(
    (t, e) => (t[e] = `--${e}-${oo}`, t),
    {}
  );
  function vf(t, e) {
    return Cf(t) && dn[t.property] && e ? (e.children.appendData(q(U({}, t), {
      property: dn[t.property]
    })), { updated: !0 }) : {};
  }
  function Ef(t) {
    for (const e of t) {
      let n = !1;
      const s = Lt(e.css, !0);
      $t(s, {
        visit: "Declaration",
        enter(r) {
          var l;
          const o = (l = this.rule) == null ? void 0 : l.block, { updated: a } = vf(r, o);
          a && (n = !0);
        }
      }), n && (e.css = Z(s), e.changed = !0);
    }
    return t.some((e) => e.changed === !0);
  }
  var ao = /* @__PURE__ */ ((t) => (t.All = "all", t.None = "none", t))(ao || {});
  function ot(t, e) {
    var s;
    return e = (s = dn[e]) != null ? s : e, (t instanceof HTMLElement ? getComputedStyle(t) : t.computedStyle).getPropertyValue(e).trim();
  }
  function Zt(t, e, n) {
    return ot(t, e) === n;
  }
  function $f(t, { selector: e, pseudoElementPart: n }) {
    const s = getComputedStyle(t, n), r = document.createElement("div"), o = document.createElement("style");
    r.id = `fake-pseudo-element-${lt()}`;
    for (const l of Array.from(s)) {
      const u = s.getPropertyValue(l);
      r.style.setProperty(l, u);
    }
    o.textContent += `#${r.id}${n} { content: ${s.content}; }`, o.textContent += `${e} { display: none !important; }`, document.head.append(o);
    const a = n === "::before" ? "afterbegin" : "beforeend";
    return t.insertAdjacentElement(a, r), { fakePseudoElement: r, sheet: o, computedStyle: s };
  }
  function Lf(t) {
    let e = t;
    for (; e; ) {
      if (Zt(e, "overflow", "scroll"))
        return e;
      e = e.parentElement;
    }
    return e;
  }
  function Of(t) {
    let e = Lf(t);
    return e === document.documentElement && (e = null), e != null ? e : { scrollTop: 0, scrollLeft: 0 };
  }
  function Pf(t, e) {
    const { elementPart: n, pseudoElementPart: s } = t, r = [];
    if (s && !(s === "::before" || s === "::after")) return r;
    const l = ie(e.roots, n);
    if (!s)
      return r.push(...l), r;
    for (const u of l) {
      const { fakePseudoElement: i, sheet: c, computedStyle: h } = $f(
        u,
        t
      ), f = i.getBoundingClientRect(), { scrollY: p, scrollX: d } = globalThis, m = Of(u);
      r.push({
        fakePseudoElement: i,
        computedStyle: h,
        removeFakePseudoElement() {
          i.remove(), c.remove();
        },
        // For https://floating-ui.com/docs/autoupdate#ancestorscroll to work on
        // `VirtualElement`s.
        contextElement: u,
        // https://floating-ui.com/docs/virtual-elements
        getBoundingClientRect() {
          const { scrollY: k, scrollX: y } = globalThis, { scrollTop: x, scrollLeft: A } = m;
          return DOMRect.fromRect({
            y: f.y + (p - k) + (m.scrollTop - x),
            x: f.x + (d - y) + (m.scrollLeft - A),
            width: f.width,
            height: f.height
          });
        }
      });
    }
    return r;
  }
  function Rf(t, e) {
    const n = ot(t, "anchor-name");
    return e ? n.split(",").map((s) => s.trim()).includes(e) : !n;
  }
  function _f(t, e) {
    const n = ot(t, "anchor-scope");
    return n === e || n === "all";
  }
  const En = (t) => R(null, null, function* () {
    var n, s, r;
    let e = yield (n = V.getOffsetParent) == null ? void 0 : n.call(V, t);
    return (yield (s = V.isElement) == null ? void 0 : s.call(V, e)) || (e = (yield (r = V.getDocumentElement) == null ? void 0 : r.call(V, t)) || window.document.documentElement), e;
  }), ie = (t, e) => t.flatMap(
    (n) => [...n.querySelectorAll(e)]
  ), ms = "InvalidMimeType";
  function If(t) {
    return !!((t.type === "text/css" || t.rel === "stylesheet") && t.href);
  }
  function Nf(t) {
    const e = new URL(t.href, document.baseURI);
    if (If(t) && e.origin === location.origin)
      return e;
  }
  function Df(t) {
    return R(this, null, function* () {
      return (yield Promise.all(
        t.map((n) => R(null, null, function* () {
          var s;
          if (!n.url)
            return n;
          if ((s = n.el) != null && s.disabled)
            return null;
          try {
            const r = yield fetch(n.url.toString()), o = r.headers.get("content-type");
            if (!(o != null && o.startsWith("text/css"))) {
              const l = new Error(
                `Error loading ${n.url}: expected content-type "text/css", got "${o}".`
              );
              throw l.name = ms, l;
            }
            const a = yield r.text();
            return q(U({}, n), { css: a });
          } catch (r) {
            if (r instanceof Error && r.name === ms)
              return console.warn(r), null;
            throw r;
          }
        }))
      )).filter((n) => n !== null);
    });
  }
  const ks = '[style*="anchor"]', Ss = '[style*="position-area"]';
  function Ff(t) {
    const e = t ? t.filter(
      (s) => s instanceof HTMLElement && (s.matches(ks) || s.matches(Ss))
    ) : Array.from(
      document.querySelectorAll(
        [
          ks,
          Ss
        ].join(",")
      )
    ), n = [];
    return e.filter((s) => s instanceof HTMLElement).forEach((s) => {
      const r = lt(12), o = "data-has-inline-styles";
      s.setAttribute(o, r);
      const a = s.getAttribute("style"), l = `[${o}="${r}"] { ${a} }`;
      n.push({ el: s, css: l });
    }), n;
  }
  function Mf(t) {
    return R(this, null, function* () {
      var o, a;
      const e = (o = t.elements) != null ? o : ie(t.roots, "link, style"), n = [];
      e.filter((l) => l instanceof HTMLElement).forEach((l) => {
        if (l.tagName.toLowerCase() === "link") {
          const u = Nf(l);
          u && n.push({ el: l, url: u });
        }
        l.tagName.toLowerCase() === "style" && n.push({ el: l, css: l.innerHTML });
      });
      const s = t.excludeInlineStyles ? (a = t.elements) != null ? a : [] : void 0, r = Ff(s);
      return yield Df([...n, ...r]);
    });
  }
  const jf = "useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict";
  let lo = (t = 21) => {
    let e = "", n = crypto.getRandomValues(new Uint8Array(t |= 0));
    for (; t--; )
      e += jf[n[t] & 63];
    return e;
  };
  const co = "--pa-cascade-property", uo = "data-anchor-position-wrapper", ho = "data-pa-wrapper-for-", ys = "POLYFILL-POSITION-AREA", Bf = [
    "left",
    "center",
    "right",
    "span-left",
    "span-right",
    "x-start",
    "x-end",
    "span-x-start",
    "span-x-end",
    "x-self-start",
    "x-self-end",
    "span-x-self-start",
    "span-x-self-end",
    "span-all",
    "top",
    "bottom",
    "span-top",
    "span-bottom",
    "y-start",
    "y-end",
    "span-y-start",
    "span-y-end",
    "y-self-start",
    "y-self-end",
    "span-y-self-start",
    "span-y-self-end",
    "block-start",
    "block-end",
    "span-block-start",
    "span-block-end",
    "inline-start",
    "inline-end",
    "span-inline-start",
    "span-inline-end",
    "self-block-start",
    "self-block-end",
    "span-self-block-start",
    "span-self-block-end",
    "self-inline-start",
    "self-inline-end",
    "span-self-inline-start",
    "span-self-inline-end",
    "start",
    "end",
    "span-start",
    "span-end",
    "self-start",
    "self-end",
    "span-self-start",
    "span-self-end"
  ];
  function fo(t) {
    return Bf.includes(t);
  }
  const bs = {
    left: [
      0,
      1,
      "Irrelevant"
      /* Irrelevant */
    ],
    center: [
      1,
      2,
      "Irrelevant"
      /* Irrelevant */
    ],
    right: [
      2,
      3,
      "Irrelevant"
      /* Irrelevant */
    ],
    "span-left": [
      0,
      2,
      "Irrelevant"
      /* Irrelevant */
    ],
    "span-right": [
      1,
      3,
      "Irrelevant"
      /* Irrelevant */
    ],
    "x-start": [
      0,
      1,
      "Physical"
      /* Physical */
    ],
    "x-end": [
      2,
      3,
      "Physical"
      /* Physical */
    ],
    "span-x-start": [
      0,
      2,
      "Physical"
      /* Physical */
    ],
    "span-x-end": [
      1,
      3,
      "Physical"
      /* Physical */
    ],
    "x-self-start": [
      0,
      1,
      "PhysicalSelf"
      /* PhysicalSelf */
    ],
    "x-self-end": [
      2,
      3,
      "PhysicalSelf"
      /* PhysicalSelf */
    ],
    "span-x-self-start": [
      0,
      2,
      "PhysicalSelf"
      /* PhysicalSelf */
    ],
    "span-x-self-end": [
      1,
      3,
      "PhysicalSelf"
      /* PhysicalSelf */
    ],
    "span-all": [
      0,
      3,
      "Irrelevant"
      /* Irrelevant */
    ],
    top: [
      0,
      1,
      "Irrelevant"
      /* Irrelevant */
    ],
    bottom: [
      2,
      3,
      "Irrelevant"
      /* Irrelevant */
    ],
    "span-top": [
      0,
      2,
      "Irrelevant"
      /* Irrelevant */
    ],
    "span-bottom": [
      1,
      3,
      "Irrelevant"
      /* Irrelevant */
    ],
    "y-start": [
      0,
      1,
      "Physical"
      /* Physical */
    ],
    "y-end": [
      2,
      3,
      "Physical"
      /* Physical */
    ],
    "span-y-start": [
      0,
      2,
      "Physical"
      /* Physical */
    ],
    "span-y-end": [
      1,
      3,
      "Physical"
      /* Physical */
    ],
    "y-self-start": [
      0,
      1,
      "PhysicalSelf"
      /* PhysicalSelf */
    ],
    "y-self-end": [
      2,
      3,
      "PhysicalSelf"
      /* PhysicalSelf */
    ],
    "span-y-self-start": [
      0,
      2,
      "PhysicalSelf"
      /* PhysicalSelf */
    ],
    "span-y-self-end": [
      1,
      3,
      "PhysicalSelf"
      /* PhysicalSelf */
    ],
    "block-start": [
      0,
      1,
      "Logical"
      /* Logical */
    ],
    "block-end": [
      2,
      3,
      "Logical"
      /* Logical */
    ],
    "span-block-start": [
      0,
      2,
      "Logical"
      /* Logical */
    ],
    "span-block-end": [
      1,
      3,
      "Logical"
      /* Logical */
    ],
    "inline-start": [
      0,
      1,
      "Logical"
      /* Logical */
    ],
    "inline-end": [
      2,
      3,
      "Logical"
      /* Logical */
    ],
    "span-inline-start": [
      0,
      2,
      "Logical"
      /* Logical */
    ],
    "span-inline-end": [
      1,
      3,
      "Logical"
      /* Logical */
    ],
    "self-block-start": [
      0,
      1,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "self-block-end": [
      2,
      3,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "span-self-block-start": [
      0,
      2,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "span-self-block-end": [
      1,
      3,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "self-inline-start": [
      0,
      1,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "self-inline-end": [
      2,
      3,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "span-self-inline-start": [
      0,
      2,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "span-self-inline-end": [
      1,
      3,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    start: [
      0,
      1,
      "Logical"
      /* Logical */
    ],
    end: [
      2,
      3,
      "Logical"
      /* Logical */
    ],
    "span-start": [
      0,
      2,
      "Logical"
      /* Logical */
    ],
    "span-end": [
      1,
      3,
      "Logical"
      /* Logical */
    ],
    "self-start": [
      0,
      1,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "self-end": [
      2,
      3,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "span-self-start": [
      0,
      2,
      "LogicalSelf"
      /* LogicalSelf */
    ],
    "span-self-end": [
      1,
      3,
      "LogicalSelf"
      /* LogicalSelf */
    ]
  }, Wf = [
    "left",
    "center",
    "right",
    "span-left",
    "span-right",
    "x-start",
    "x-end",
    "span-x-start",
    "span-x-end",
    "x-self-start",
    "x-self-end",
    "span-x-self-start",
    "span-x-self-end",
    "span-all"
  ], Uf = [
    "top",
    "center",
    "bottom",
    "span-top",
    "span-bottom",
    "y-start",
    "y-end",
    "span-y-start",
    "span-y-end",
    "y-self-start",
    "y-self-end",
    "span-y-self-start",
    "span-y-self-end",
    "span-all"
  ], zf = [
    "block-start",
    "center",
    "block-end",
    "span-block-start",
    "span-block-end",
    "span-all"
  ], Vf = [
    "inline-start",
    "center",
    "inline-end",
    "span-inline-start",
    "span-inline-end",
    "span-all"
  ], Hf = [
    "self-block-start",
    "center",
    "self-block-end",
    "span-self-block-start",
    "span-self-block-end",
    "span-all"
  ], Gf = [
    "self-inline-start",
    "center",
    "self-inline-end",
    "span-self-inline-start",
    "span-self-inline-end",
    "span-all"
  ], xs = [
    "start",
    "center",
    "end",
    "span-start",
    "span-end",
    "span-all"
  ], Cs = [
    "self-start",
    "center",
    "self-end",
    "span-self-start",
    "span-self-end",
    "span-all"
  ], qf = ["block", "top", "bottom", "y"], Kf = ["inline", "left", "right", "x"];
  function gn(t) {
    const e = t.split("-");
    for (const n of e) {
      if (qf.includes(n)) return "block";
      if (Kf.includes(n)) return "inline";
    }
    return "ambiguous";
  }
  function Qf(t, e) {
    return e[0].includes(t[0]) && e[1].includes(t[1]) || e[0].includes(t[1]) && e[1].includes(t[0]);
  }
  const Yf = [
    [Wf, Uf],
    [zf, Vf],
    [Hf, Gf],
    [xs, xs],
    [Cs, Cs]
  ];
  function Xf(t) {
    for (const e of Yf)
      if (Qf(t, e)) return !0;
    return !1;
  }
  const ws = (t) => {
    const e = getComputedStyle(t);
    return {
      writingMode: e.writingMode,
      direction: e.direction
    };
  }, Jf = (t, e) => R(null, null, function* () {
    const n = yield En(t);
    switch (e) {
      case "Logical":
      case "Physical":
        return ws(n);
      case "LogicalSelf":
      case "PhysicalSelf":
        return ws(t);
      default:
        return null;
    }
  }), tn = (t) => t.reverse().map((e) => 3 - e), po = (t, e) => t === "Irrelevant" ? e : t, Zf = (s, r) => R(null, [s, r], function* ({
    block: t,
    inline: e
  }, n) {
    const o = po(t[2], e[2]), a = yield Jf(n, o), l = {
      block: [t[0], t[1]],
      inline: [e[0], e[1]]
    };
    if (a) {
      if (a.direction === "rtl" && (l.inline = tn(l.inline)), a.writingMode.startsWith("vertical")) {
        const u = l.block;
        l.block = l.inline, l.inline = u;
      }
      if (a.writingMode.startsWith("sideways")) {
        const u = l.block;
        l.block = l.inline, l.inline = u, a.writingMode.endsWith("lr") && (l.block = tn(l.block));
      }
      a.writingMode.endsWith("rl") && (l.inline = tn(l.inline));
    }
    return l;
  }), tp = ({
    block: t,
    inline: e
  }) => {
    const n = [0, "top", "bottom", 0], s = [0, "left", "right", 0];
    return {
      block: [n[t[0]], n[t[1]]],
      inline: [s[e[0]], s[e[1]]]
    };
  };
  function Ts([t, e]) {
    return t === 0 && e === 3 ? "center" : t === 0 ? "end" : e === 3 ? "start" : "center";
  }
  function ep(t) {
    return t.type === "Declaration" && t.property === "position-area";
  }
  function np(t) {
    const e = t.value.children.toArray().map(({ name: n }) => n);
    return e.length === 1 && (gn(e[0]) === "ambiguous" ? e.push(e[0]) : e.push("span-all")), e;
  }
  function sp(t) {
    if (!ep(t)) return;
    const e = np(t);
    if (!Xf(e)) return;
    const n = {};
    switch (gn(e[0])) {
      case "block":
        n.block = e[0], n.inline = e[1];
        break;
      case "inline":
        n.inline = e[0], n.block = e[1];
        break;
      case "ambiguous":
        gn(e[1]) == "block" ? (n.block = e[1], n.inline = e[0]) : (n.inline = e[1], n.block = e[0]);
        break;
    }
    const s = {
      block: bs[n.block],
      inline: bs[n.inline]
    }, r = `--pa-declaration-${lo(12)}`;
    return {
      values: n,
      grid: s,
      selectorUUID: r
    };
  }
  function rp(t, e) {
    [
      // Insets are applied to a wrapping element
      "justify-self",
      "align-self"
    ].forEach((n) => {
      e.children.appendData({
        type: "Declaration",
        property: n,
        value: { type: "Raw", value: `var(--pa-value-${n})` },
        important: !1
      });
    }), e.children.appendData({
      type: "Declaration",
      property: co,
      value: { type: "Raw", value: t.selectorUUID },
      important: !1
    });
  }
  function ip(t, e) {
    var s, r;
    let n;
    if (((s = t.parentElement) == null ? void 0 : s.tagName) === ys)
      n = t.parentElement;
    else {
      n = document.createElement(ys), n.style.display = "grid", n.style.position = "absolute";
      const o = getComputedStyle(t).pointerEvents;
      n.style.pointerEvents = "none", t.style.pointerEvents = o, ["top", "left", "right", "bottom"].forEach((a) => {
        n.style.setProperty(a, `var(--pa-value-${a})`);
      }), (r = t.parentElement) == null || r.insertBefore(n, t), n.appendChild(t);
    }
    return n.setAttribute(
      `${ho}${e}`,
      ""
    ), n;
  }
  function op(t, e, n) {
    return R(this, null, function* () {
      const s = `--pa-target-${lo(12)}`, r = yield Zf(
        e.grid,
        t
      ), o = tp(r), a = po(
        e.grid.block[2],
        e.grid.inline[2]
      ), l = [
        "LogicalSelf",
        "PhysicalSelf"
        /* PhysicalSelf */
      ].includes(a) ? r : e.grid, u = {
        block: Ts([l.block[0], l.block[1]]),
        inline: Ts([
          l.inline[0],
          l.inline[1]
        ])
      };
      return {
        insets: o,
        alignments: u,
        targetUUID: s,
        targetEl: t,
        anchorEl: n,
        wrapperEl: ip(t, s),
        values: e.values,
        grid: e.grid,
        selectorUUID: e.selectorUUID
      };
    });
  }
  function ap(t, e) {
    return `
    [${uo}="${e}"][${ho}${t}] {
      --pa-value-top: var(${t}-top);
      --pa-value-left: var(${t}-left);
      --pa-value-right: var(${t}-right);
      --pa-value-bottom: var(${t}-bottom);
      --pa-value-justify-self: var(${t}-justify-self);
      --pa-value-align-self: var(${t}-align-self);
    }
  `.replaceAll(`
`, "");
  }
  const lp = [
    "normal",
    "most-width",
    "most-height",
    "most-block-size",
    "most-inline-size"
  ], cp = [
    "flip-block",
    "flip-inline",
    "flip-start"
  ];
  function up(t) {
    return t.type === "Declaration";
  }
  function hp(t) {
    return t.type === "Declaration" && t.property === "position-try-fallbacks";
  }
  function fp(t) {
    return t.type === "Declaration" && t.property === "position-try-order";
  }
  function pp(t) {
    return t.type === "Declaration" && t.property === "position-try";
  }
  function dp(t) {
    return t.type === "Atrule" && t.name === "position-try";
  }
  function gp(t) {
    return cp.includes(t);
  }
  function mp(t) {
    return lp.includes(t);
  }
  function kp(t, e) {
    const n = document.querySelector(t);
    if (n) {
      let s = yp(n);
      return e.forEach((r) => {
        s = go(s, r);
      }), s;
    }
  }
  function Sp(t, e) {
    let n = t.declarations;
    return e.forEach((s) => {
      n = go(n, s);
    }), n;
  }
  function yp(t) {
    const e = {};
    return Zi.forEach((n) => {
      const s = ot(t, `--${n}-${oo}`);
      s && (e[n] = s);
    }), e;
  }
  const bp = {
    "flip-block": {
      top: "bottom",
      bottom: "top",
      "inset-block-start": "inset-block-end",
      "inset-block-end": "inset-block-start",
      "margin-top": "margin-bottom",
      "margin-bottom": "margin-top"
    },
    "flip-inline": {
      left: "right",
      right: "left",
      "inset-inline-start": "inset-inline-end",
      "inset-inline-end": "inset-inline-start",
      "margin-left": "margin-right",
      "margin-right": "margin-left"
    },
    "flip-start": {
      left: "top",
      right: "bottom",
      top: "left",
      bottom: "right",
      "inset-block-start": "inset-block-end",
      "inset-block-end": "inset-block-start",
      "inset-inline-start": "inset-inline-end",
      "inset-inline-end": "inset-inline-start",
      "inset-block": "inset-inline",
      "inset-inline": "inset-block"
    }
  }, xp = {
    "flip-block": {
      top: "bottom",
      bottom: "top",
      start: "end",
      end: "start",
      "self-end": "self-start",
      "self-start": "self-end"
    },
    "flip-inline": {
      left: "right",
      right: "left",
      start: "end",
      end: "start",
      "self-end": "self-start",
      "self-start": "self-end"
    },
    "flip-start": {
      top: "left",
      left: "top",
      right: "bottom",
      bottom: "right"
    }
  }, Cp = {
    "flip-block": {
      top: "bottom",
      bottom: "top",
      start: "end",
      end: "start"
    },
    "flip-inline": {
      left: "right",
      right: "left",
      start: "end",
      end: "start"
    },
    "flip-start": {
      // TODO: Requires fuller logic
    }
  };
  function wp(t, e) {
    return bp[e][t] || t;
  }
  function Tp(t, e) {
    return xp[e][t] || t;
  }
  function Ap(t, e) {
    if (e === "flip-start")
      return t;
    {
      const n = Cp[e];
      return t.split("-").map((s) => n[s] || s).join("-");
    }
  }
  function vp(t, e, n) {
    if (t === "margin") {
      const [s, r, o, a] = e.children.toArray();
      n === "flip-block" ? a ? e.children.fromArray([o, r, s, a]) : o && e.children.fromArray([o, r, s]) : n === "flip-inline" && a && e.children.fromArray([s, a, o, r]);
    } else if (t === "margin-block") {
      const [s, r] = e.children.toArray();
      n === "flip-block" && r && e.children.fromArray([r, s]);
    } else if (t === "margin-inline") {
      const [s, r] = e.children.toArray();
      n === "flip-inline" && r && e.children.fromArray([r, s]);
    }
  }
  const Ep = (t, e) => {
    var r;
    return ((r = Lt(`#id{${t}: ${e};}`).children.first) == null ? void 0 : r.block.children.first).value;
  };
  function go(t, e) {
    const n = {};
    return Object.entries(t).forEach(([s, r]) => {
      var u;
      const o = s, a = Ep(o, r), l = wp(o, e);
      l !== o && ((u = n[o]) != null || (n[o] = "revert")), $t(a, {
        visit: "Function",
        enter(i) {
          Re(i) && i.children.forEach((c) => {
            oe(c) && eo(c.name) && (c.name = Tp(c.name, e));
          });
        }
      }), o === "position-area" && a.children.forEach((i) => {
        oe(i) && fo(i.name) && (i.name = Ap(i.name, e));
      }), o.startsWith("margin") && vp(o, a, e), n[l] = Z(a);
    }), n;
  }
  function mo(t) {
    const e = wf(t), n = [];
    return e.forEach((s) => {
      const r = {
        atRules: [],
        tactics: [],
        positionAreas: []
      };
      s.forEach((o) => {
        gp(o.name) ? r.tactics.push(o.name) : o.name.startsWith("--") ? r.atRules.push(o.name) : fo(o.name) && r.positionAreas.push(o.name);
      }), r.positionAreas.length ? n.push({
        positionArea: r.positionAreas[0],
        type: "position-area"
      }) : r.atRules.length && r.tactics.length ? n.push({
        tactics: r.tactics,
        atRule: r.atRules[0],
        type: "at-rule-with-try-tactic"
      }) : r.atRules.length ? n.push({
        atRule: r.atRules[0],
        type: "at-rule"
      }) : r.tactics.length && n.push({
        tactics: r.tactics,
        type: "try-tactic"
      });
    }), n;
  }
  function $p(t) {
    return hp(t) && t.value.children.first ? mo(t.value.children) : [];
  }
  function Lp(t) {
    if (pp(t) && t.value.children.first) {
      const e = Pe(t);
      let n;
      const s = e.value.children.first.name;
      s && mp(s) && (n = s, e.value.children.shift());
      const r = mo(e.value.children);
      return { order: n, options: r };
    }
    return {};
  }
  function Op(t) {
    return fp(t) && t.value.children.first ? {
      order: t.value.children.first.name
    } : {};
  }
  function Pp(t) {
    const { order: e, options: n } = Lp(t);
    if (e || n)
      return { order: e, options: n };
    const { order: s } = Op(t), r = $p(t);
    return s || r ? { order: s, options: r } : {};
  }
  function Rp(t) {
    return ue(t.property) || ch(t.property) || Xi(t.property) || uh(t.property) || ["position-anchor", "position-area"].includes(t.property);
  }
  function _p(t) {
    var e, n;
    if (dp(t) && ((e = t.prelude) != null && e.value) && ((n = t.block) != null && n.children)) {
      const s = t.prelude.value, r = t.block.children.filter(
        (a) => up(a) && Rp(a)
      ), o = {
        uuid: `${s}-try-${lt(12)}`,
        declarations: Object.fromEntries(
          r.map((a) => [a.property, Z(a.value)])
        )
      };
      return { name: s, tryBlock: o };
    }
    return {};
  }
  function Ip(t) {
    const e = {}, n = {}, s = {};
    for (const r of t) {
      const o = Lt(r.css);
      $t(o, {
        visit: "Atrule",
        enter(a) {
          const { name: l, tryBlock: u } = _p(a);
          l && u && (e[l] = u);
        }
      });
    }
    for (const r of t) {
      let o = !1;
      const a = /* @__PURE__ */ new Set(), l = Lt(r.css);
      $t(l, {
        visit: "Declaration",
        enter(u) {
          var d;
          const i = (d = this.rule) == null ? void 0 : d.prelude, c = pn(i);
          if (!c.length) return;
          const { order: h, options: f } = Pp(u), p = {};
          h && (p.order = h), c.forEach(({ selector: m }) => {
            var k, y;
            f == null || f.forEach((x) => {
              var M, et, w;
              let A;
              if (x.type === "at-rule")
                A = x.atRule;
              else if (x.type === "try-tactic") {
                A = `${m}-${x.tactics.join("-")}`;
                const C = kp(
                  m,
                  x.tactics
                );
                C && (e[A] = {
                  uuid: `${m}-${x.tactics.join("-")}-try-${lt(12)}`,
                  declarations: C
                });
              } else if (x.type === "at-rule-with-try-tactic") {
                A = `${m}-${x.atRule}-${x.tactics.join("-")}`;
                const C = e[x.atRule], P = Sp(
                  C,
                  x.tactics
                );
                P && (e[A] = {
                  uuid: `${m}-${x.atRule}-${x.tactics.join("-")}-try-${lt(12)}`,
                  declarations: P
                });
              }
              if (A && e[A]) {
                const C = `[data-anchor-polyfill="${e[A].uuid}"]`;
                (M = n[C]) != null || (n[C] = []), n[C].push(m), a.has(A) || ((et = p.fallbacks) != null || (p.fallbacks = []), p.fallbacks.push(e[A]), a.add(A), (w = this.stylesheet) == null || w.children.prependData({
                  type: "Rule",
                  prelude: {
                    type: "Raw",
                    value: C
                  },
                  block: {
                    type: "Block",
                    children: new K().fromArray(
                      Object.entries(e[A].declarations).map(
                        ([P, O]) => ({
                          type: "Declaration",
                          important: !0,
                          property: P,
                          value: {
                            type: "Raw",
                            value: O
                          }
                        })
                      )
                    )
                  }
                }), o = !0);
              }
            }), Object.keys(p).length > 0 && (s[m] ? (p.order && (s[m].order = p.order), p.fallbacks && ((y = (k = s[m]).fallbacks) != null || (k.fallbacks = []), s[m].fallbacks.push(
              ...p.fallbacks
            ))) : s[m] = p);
          });
        }
      }), o && (r.css = Z(l), r.changed = !0);
    }
    return { fallbackTargets: n, validPositions: s };
  }
  function Np(t, e) {
    return !t || t === e ? !1 : ko(t) ? t.document.contains(e) : t.contains(e);
  }
  function ko(t) {
    return !!(t && t === t.window);
  }
  function Dp(t) {
    return Zt(t, "position", "fixed");
  }
  function mn(t) {
    return !!(t && (Dp(t) || Zt(t, "position", "absolute")));
  }
  function As(t, e) {
    return t.compareDocumentPosition(e) & Node.DOCUMENT_POSITION_FOLLOWING;
  }
  function Fp(t) {
    return R(this, null, function* () {
      return yield V.getOffsetParent(t);
    });
  }
  function en(t) {
    return R(this, null, function* () {
      if (!["absolute", "fixed"].includes(ot(t, "position")))
        return yield Fp(t);
      let e = t.parentElement;
      for (; e; ) {
        if (!Zt(e, "position", "static") && Zt(e, "display", "block"))
          return e;
        e = e.parentElement;
      }
      return window;
    });
  }
  function Mp(t, e, n, s) {
    return R(this, null, function* () {
      const r = yield en(t), o = yield en(n);
      if (!(Np(o, t) || ko(o)) || r === o && !(!mn(t) || As(t, n)))
        return !1;
      if (r !== o) {
        let a;
        const l = [];
        for (a = r; a && a !== o && a !== window; )
          l.push(a), a = yield en(a);
        const u = l[l.length - 1];
        if (u instanceof HTMLElement && !(!mn(u) || As(u, n)))
          return !1;
      }
      {
        let a = t.parentElement;
        for (; a; ) {
          if (Zt(a, "content-visibility", "hidden"))
            return !1;
          a = a.parentElement;
        }
      }
      return !(e && s && vs(t, e, s) !== vs(n, e, s));
    });
  }
  function vs(t, e, n) {
    for (; !(t.matches(n) && _f(t, e)); ) {
      if (!t.parentElement)
        return null;
      t = t.parentElement;
    }
    return t;
  }
  function jp(t, e, n, s, r) {
    return R(this, null, function* () {
      if (!(t instanceof HTMLElement && n.length && mn(t)))
        return null;
      const o = n.flatMap((l) => Pf(l, r)).filter((l) => Rf(l, e)), a = s.map((l) => l.selector).join(",") || null;
      for (let l = o.length - 1; l >= 0; l--) {
        const u = o[l], i = "fakePseudoElement" in u;
        if (yield Mp(
          i ? u.fakePseudoElement : u,
          e,
          t,
          a
        ))
          return i && u.removeFakePseudoElement(), u;
      }
      return null;
    });
  }
  function Bp(t) {
    return t.type === "Declaration" && t.property === "anchor-name";
  }
  function Wp(t) {
    return t.type === "Declaration" && t.property === "anchor-scope";
  }
  function kn(t) {
    return !!(t && t.type === "Function" && t.name === "anchor-size");
  }
  function Te(t) {
    return !!(t && t.type === "Function" && t.name === "var");
  }
  function oe(t) {
    return !!(t.type === "Identifier" && t.name);
  }
  function Up(t) {
    return !!(t.type === "Percentage" && t.value);
  }
  function Es(t, e) {
    let n, s, r, o = "", a = !1, l;
    const u = [];
    t.children.toArray().forEach((f) => {
      if (a) {
        o = `${o}${Z(f)}`;
        return;
      }
      if (f.type === "Operator" && f.value === ",") {
        a = !0;
        return;
      }
      u.push(f);
    });
    let [i, c] = u;
    if (c || (c = i, i = void 0), i && (oe(i) && i.name.startsWith("--") ? n = i.name : Te(i) && i.children.first && (l = i.children.first.name)), c)
      if (Re(t)) {
        if (oe(c) && eo(c.name))
          s = c.name;
        else if (Up(c)) {
          const f = Number(c.value);
          s = Number.isNaN(f) ? void 0 : f;
        }
      } else kn(t) && oe(c) && dh(c.name) && (r = c.name);
    const h = `--anchor-${lt(12)}`;
    return Object.assign(t, {
      type: "Raw",
      value: `var(${h})`,
      children: null
    }), Reflect.deleteProperty(t, "name"), {
      anchorName: n,
      anchorSide: s,
      anchorSize: r,
      fallbackValue: o || "0px",
      customPropName: l,
      uuid: h
    };
  }
  function $s(t) {
    return t.value.children.map(({ name: e }) => e);
  }
  let Ht = {}, Et = {}, Nt = {}, ae = {}, It = {};
  function zp() {
    Ht = {}, Et = {}, Nt = {}, ae = {}, It = {};
  }
  function Vp(t, e) {
    var n;
    if ((Re(t) || kn(t)) && e) {
      if (e.property.startsWith("--")) {
        const s = Z(e.value), r = Es(t);
        return ae[r.uuid] = s, Nt[e.property] = [
          ...(n = Nt[e.property]) != null ? n : [],
          r
        ], { changed: !0 };
      }
      if (Re(t) && ue(e.property) || kn(t) && to(e.property)) {
        const s = Es(t);
        return { prop: e.property, data: s, changed: !0 };
      }
    }
    return {};
  }
  function Ls(t, e, n) {
    return R(this, null, function* () {
      let s = e == null ? void 0 : e.anchorName;
      const r = e == null ? void 0 : e.customPropName;
      if (t && !s) {
        const u = ot(
          t,
          "position-anchor"
        );
        u ? s = u : r && (s = ot(t, r));
      }
      const o = s ? Ht[s] || [] : [], a = s ? Et[ao.All] || [] : [], l = s ? Et[s] || [] : [];
      return yield jp(
        t,
        s || null,
        o,
        [...a, ...l],
        { roots: n.roots }
      );
    });
  }
  function Hp(t, e) {
    return R(this, null, function* () {
      var h, f, p, d, m, k, y, x, A, M, et;
      const n = {}, s = {};
      zp();
      const { fallbackTargets: r, validPositions: o } = Ip(t);
      for (const w of t) {
        let C = !1;
        const P = Lt(w.css);
        $t(P, function(O) {
          var J, xt, Pt, Ct, pt, wt;
          const N = (J = this.rule) == null ? void 0 : J.prelude, $ = pn(N);
          if (Bp(O) && $.length)
            for (const B of $s(O))
              (xt = Ht[B]) != null || (Ht[B] = []), Ht[B].push(...$);
          if (Wp(O) && $.length)
            for (const B of $s(O))
              (Pt = Et[B]) != null || (Et[B] = []), Et[B].push(...$);
          const {
            prop: v,
            data: j,
            changed: D
          } = Vp(O, this.declaration);
          if (v && j && $.length)
            for (const { selector: B } of $)
              n[B] = q(U({}, n[B]), {
                [v]: [...(pt = (Ct = n[B]) == null ? void 0 : Ct[v]) != null ? pt : [], j]
              });
          let H;
          if (this.block && (H = sp(O), H)) {
            rp(
              H,
              this.block
            );
            for (const { selector: B } of $)
              s[B] = [
                ...(wt = s[B]) != null ? wt : [],
                H
              ];
          }
          (D || H) && (C = !0);
        }), C && (w.css = Z(P), w.changed = !0);
      }
      const a = new Set(Object.keys(Nt)), l = {}, u = (w) => {
        var O, N, $, v, j;
        const C = [], P = new Set((N = (O = l[w]) == null ? void 0 : O.names) != null ? N : []);
        for (; P.size > 0; )
          for (const D of P)
            C.push(...($ = Nt[D]) != null ? $ : []), P.delete(D), (j = (v = l[D]) == null ? void 0 : v.names) != null && j.length && l[D].names.forEach((H) => P.add(H));
        return C;
      };
      for (; a.size > 0; ) {
        const w = [];
        for (const C of t) {
          let P = !1;
          const O = Lt(C.css);
          $t(O, {
            visit: "Function",
            enter(N) {
              var D, H;
              const $ = (D = this.rule) == null ? void 0 : D.prelude, v = this.declaration, j = v == null ? void 0 : v.property;
              if (($ == null ? void 0 : $.children.isEmpty) === !1 && Te(N) && v && j && N.children.first && a.has(N.children.first.name) && // For now, we only want assignments to other CSS custom properties
              j.startsWith("--")) {
                const J = N.children.first, xt = (H = Nt[J.name]) != null ? H : [], Pt = u(J.name);
                if (!(xt.length || Pt.length))
                  return;
                const Ct = `${J.name}-anchor-${lt(12)}`, pt = Z(v.value);
                ae[Ct] = pt, l[j] || (l[j] = { names: [], uuids: [] });
                const wt = l[j];
                wt.names.includes(J.name) || wt.names.push(J.name), wt.uuids.push(Ct), w.push(j), J.name = Ct, P = !0;
              }
            }
          }), P && (C.css = Z(O), C.changed = !0);
        }
        a.clear(), w.forEach((C) => a.add(C));
      }
      for (const w of t) {
        let C = !1;
        const P = Lt(w.css);
        $t(P, {
          visit: "Function",
          enter(O) {
            var j, D, H, J, xt, Pt, Ct;
            const N = (j = this.rule) == null ? void 0 : j.prelude, $ = this.declaration, v = $ == null ? void 0 : $.property;
            if ((N == null ? void 0 : N.children.isEmpty) === !1 && Te(O) && $ && v && O.children.first && // Now we only want assignments to inset/sizing properties
            (ue(v) || Xi(v))) {
              const pt = O.children.first, wt = (D = Nt[pt.name]) != null ? D : [], B = u(pt.name);
              if (!(wt.length || B.length))
                return;
              const pe = `${v}-${lt(12)}`;
              if (B.length) {
                const jt = /* @__PURE__ */ new Set([pt.name]);
                for (; jt.size > 0; )
                  for (const Bt of jt) {
                    const G = l[Bt];
                    if ((H = G == null ? void 0 : G.names) != null && H.length && ((J = G == null ? void 0 : G.uuids) != null && J.length))
                      for (const Wt of G.names)
                        for (const Ut of G.uuids)
                          It[Ut] = q(U({}, It[Ut]), {
                            // - `key` (`propUuid`) is the property-specific
                            //   uuid to append to the new custom property name
                            // - `value` is the new property-specific custom
                            //   property value to use
                            [pe]: `${Wt}-${pe}`
                          });
                    jt.delete(Bt), (xt = G == null ? void 0 : G.names) != null && xt.length && G.names.forEach((Wt) => jt.add(Wt));
                  }
              }
              const So = pn(N);
              for (const jt of [...wt, ...B]) {
                const Bt = U({}, jt), G = `--anchor-${lt(12)}-${v}`, Wt = Bt.uuid;
                Bt.uuid = G;
                for (const { selector: Ut } of So)
                  n[Ut] = q(U({}, n[Ut]), {
                    [v]: [...(Ct = (Pt = n[Ut]) == null ? void 0 : Pt[v]) != null ? Ct : [], Bt]
                  });
                It[Wt] = q(U({}, It[Wt]), {
                  // - `key` (`propUuid`) is the property-specific
                  //   uuid to append to the new custom property name
                  // - `value` is the new property-specific custom
                  //   property value to use
                  [pe]: G
                });
              }
              pt.name = `${pt.name}-${pe}`, C = !0;
            }
          }
        }), C && (w.css = Z(P), w.changed = !0);
      }
      if (Object.keys(It).length > 0)
        for (const w of t) {
          let C = !1;
          const P = Lt(w.css);
          $t(P, {
            visit: "Function",
            enter(O) {
              var N, $, v, j;
              if (Te(O) && (($ = (N = O.children.first) == null ? void 0 : N.name) != null && $.startsWith("--")) && ((j = (v = this.declaration) == null ? void 0 : v.property) != null && j.startsWith("--")) && this.block) {
                const D = O.children.first, H = It[D.name];
                if (H)
                  for (const [J, xt] of Object.entries(H))
                    this.block.children.appendData({
                      type: "Declaration",
                      important: !1,
                      property: `${this.declaration.property}-${J}`,
                      value: {
                        type: "Raw",
                        value: Z(this.declaration.value).replace(
                          `var(${D.name})`,
                          `var(${xt})`
                        )
                      }
                    }), C = !0;
                ae[D.name] && (this.declaration.value = {
                  type: "Raw",
                  value: ae[D.name]
                }, C = !0);
              }
            }
          }), C && (w.css = Z(P), w.changed = !0);
        }
      const i = /* @__PURE__ */ new Map();
      for (const [w, C] of Object.entries(n)) {
        let P;
        w.startsWith("[data-anchor-polyfill=") && ((h = r[w]) != null && h.length) ? P = ie(
          e.roots,
          r[w].join(",")
        ) : P = ie(e.roots, w);
        for (const [O, N] of Object.entries(C))
          for (const $ of N)
            for (const v of P) {
              const j = yield Ls(v, $, {
                roots: e.roots
              }), D = `--anchor-${lt(12)}`;
              i.set(v, q(U({}, (f = i.get(v)) != null ? f : {}), {
                [$.uuid]: D
              })), v.setAttribute(
                "style",
                `${$.uuid}: var(${D}); ${(p = v.getAttribute("style")) != null ? p : ""}`
              ), o[w] = q(U({}, o[w]), {
                declarations: q(U({}, (d = o[w]) == null ? void 0 : d.declarations), {
                  [O]: [
                    ...(y = (k = (m = o[w]) == null ? void 0 : m.declarations) == null ? void 0 : k[O]) != null ? y : [],
                    q(U({}, $), { anchorEl: j, targetEl: v, uuid: D })
                  ]
                })
              });
            }
      }
      const c = {
        el: document.createElement("link"),
        changed: !1,
        created: !0,
        css: ""
      };
      t.push(c);
      for (const [w, C] of Object.entries(s)) {
        const P = ie(e.roots, w);
        for (const O of P) {
          const N = yield Ls(O, null, {
            roots: e.roots
          });
          for (const $ of C) {
            const v = yield op(
              O,
              $,
              N
            );
            c.css += ap(
              v.targetUUID,
              $.selectorUUID
            ), c.changed = !0, o[w] = q(U({}, o[w]), {
              declarations: q(U({}, (x = o[w]) == null ? void 0 : x.declarations), {
                "position-area": [
                  ...(et = (M = (A = o[w]) == null ? void 0 : A.declarations) == null ? void 0 : M["position-area"]) != null ? et : [],
                  v
                ]
              })
            });
          }
        }
      }
      return { rules: o, inlineStyles: i, anchorScopes: Et };
    });
  }
  const Gp = [
    "as",
    "blocking",
    "crossorigin",
    // 'disabled' is not relevant for style elements, but this exclusion is
    // theoretical, as a <link rel=stylesheet disabled> will not be loaded, and
    // will not reach this part of the polyfill. See #246.
    "disabled",
    "fetchpriority",
    "href",
    "hreflang",
    "integrity",
    "referrerpolicy",
    "rel",
    "type"
  ];
  function Os(t, e, n = !1) {
    const s = [];
    for (const { el: r, css: o, changed: a, created: l = !1 } of t) {
      const u = { el: r, css: o, changed: !1 };
      if (a) {
        if (r.tagName.toLowerCase() === "style")
          r.innerHTML = o;
        else if (r instanceof HTMLLinkElement) {
          const i = document.createElement("style");
          i.textContent = o;
          for (const c of r.getAttributeNames())
            if (!c.startsWith("on") && !Gp.includes(c)) {
              const h = r.getAttribute(c);
              h !== null && i.setAttribute(c, h);
            }
          r.hasAttribute("href") && i.setAttribute("data-original-href", r.getAttribute("href")), l ? (i.setAttribute("data-generated-by-polyfill", "true"), document.head.insertAdjacentElement("beforeend", i)) : (r.insertAdjacentElement("beforebegin", i), r.remove()), u.el = i;
        } else if (r.hasAttribute("data-has-inline-styles")) {
          const i = r.getAttribute("data-has-inline-styles");
          if (i) {
            const c = `[data-has-inline-styles="${i}"]{`;
            let f = o.slice(c.length, 0 - "}".length);
            const p = e == null ? void 0 : e.get(r);
            if (p)
              for (const [d, m] of Object.entries(p))
                f = `${d}: var(${m}); ${f}`;
            r.setAttribute("style", f);
          }
        }
      }
      n && r.hasAttribute("data-has-inline-styles") && r.removeAttribute("data-has-inline-styles"), s.push(u);
    }
    return s;
  }
  const qp = q(U({}, V), { _c: /* @__PURE__ */ new Map() }), Kp = (t, e) => {
    let n;
    switch (t) {
      case "start":
      case "self-start":
        n = 0;
        break;
      case "end":
      case "self-end":
        n = 100;
        break;
      default:
        typeof t == "number" && !Number.isNaN(t) && (n = t);
    }
    if (n !== void 0)
      return e ? 100 - n : n;
  }, Qp = (t, e) => {
    let n;
    switch (t) {
      case "block":
      case "self-block":
        n = e ? "width" : "height";
        break;
      case "inline":
      case "self-inline":
        n = e ? "height" : "width";
        break;
    }
    return n;
  }, Ps = (t) => {
    switch (t) {
      case "top":
      case "bottom":
        return "y";
      case "left":
      case "right":
        return "x";
    }
    return null;
  }, Yp = (t) => {
    switch (t) {
      case "x":
        return "width";
      case "y":
        return "height";
    }
    return null;
  }, Rs = (t) => ot(t, "display") === "inline", _s = (t, e) => (e === "x" ? ["border-left-width", "border-right-width"] : ["border-top-width", "border-bottom-width"]).reduce(
    (s, r) => s + parseInt(ot(t, r), 10),
    0
  ) || 0, ye = (t, e) => parseInt(ot(t, `margin-${e}`), 10) || 0, Xp = (t) => ({
    top: ye(t, "top"),
    right: ye(t, "right"),
    bottom: ye(t, "bottom"),
    left: ye(t, "left")
  }), nn = (a) => R(null, [a], function* ({
    targetEl: t,
    targetProperty: e,
    anchorRect: n,
    anchorSide: s,
    anchorSize: r,
    fallback: o = null
  }) {
    var l;
    if (!((r || s !== void 0) && t && n))
      return o;
    if (r) {
      if (!to(e))
        return o;
      let u;
      switch (r) {
        case "width":
        case "height":
          u = r;
          break;
        default: {
          let i = !1;
          const c = ot(t, "writing-mode");
          i = c.startsWith("vertical-") || c.startsWith("sideways-"), u = Qp(r, i);
        }
      }
      return u ? `${n[u]}px` : o;
    }
    if (s !== void 0) {
      let u, i;
      const c = Ps(e);
      if (!(ue(e) && c && (!ue(s) || c === Ps(s))))
        return o;
      const h = ["top", "left"];
      switch (s) {
        case "left":
        case "top":
          u = 0;
          break;
        case "right":
        case "bottom":
          u = 100;
          break;
        case "center":
          u = 50;
          break;
        case "inside":
          u = h.includes(e) ? 0 : 100;
          break;
        case "outside":
          u = h.includes(e) ? 100 : 0;
          break;
        default:
          if (t) {
            const d = (yield (l = V.isRTL) == null ? void 0 : l.call(V, t)) || !1;
            u = Kp(s, d);
          }
      }
      const f = typeof u == "number" && !Number.isNaN(u), p = Yp(c);
      if (f && p) {
        (e === "bottom" || e === "right") && (i = yield En(t));
        let d = n[c] + n[p] * (u / 100);
        switch (e) {
          case "bottom": {
            if (!i)
              break;
            let m = i.clientHeight;
            if (m === 0 && Rs(i)) {
              const k = _s(i, c);
              m = i.offsetHeight - k;
            }
            d = m - d;
            break;
          }
          case "right": {
            if (!i)
              break;
            let m = i.clientWidth;
            if (m === 0 && Rs(i)) {
              const k = _s(i, c);
              m = i.offsetWidth - k;
            }
            d = m - d;
            break;
          }
        }
        return `${d}px`;
      }
    }
    return o;
  }), Jp = (t) => "wrapperEl" in t, Zp = (t) => "uuid" in t;
  function td(t, e = !1) {
    return R(this, null, function* () {
      const n = document.documentElement;
      for (const [s, r] of Object.entries(t))
        for (const o of r) {
          const a = o.anchorEl, l = o.targetEl;
          if (a && l)
            if (Jp(o)) {
              const u = o.wrapperEl, i = (c, h, f) => R(null, null, function* () {
                return c === 0 ? "0px" : yield nn({
                  targetEl: u,
                  targetProperty: h,
                  anchorRect: f,
                  anchorSide: c
                });
              });
              on(
                a,
                u,
                () => R(null, null, function* () {
                  const c = ot(
                    l,
                    co
                  );
                  u.setAttribute(uo, c);
                  const h = yield V.getElementRects({
                    reference: a,
                    floating: u,
                    strategy: "absolute"
                  }), f = o.insets, p = yield i(
                    f.block[0],
                    "top",
                    h.reference
                  ), d = yield i(
                    f.block[1],
                    "bottom",
                    h.reference
                  ), m = yield i(
                    f.inline[0],
                    "left",
                    h.reference
                  ), k = yield i(
                    f.inline[1],
                    "right",
                    h.reference
                  );
                  n.style.setProperty(
                    `${o.targetUUID}-top`,
                    p || null
                  ), n.style.setProperty(
                    `${o.targetUUID}-left`,
                    m || null
                  ), n.style.setProperty(
                    `${o.targetUUID}-right`,
                    k || null
                  ), n.style.setProperty(
                    `${o.targetUUID}-bottom`,
                    d || null
                  ), n.style.setProperty(
                    `${o.targetUUID}-justify-self`,
                    o.alignments.inline
                  ), n.style.setProperty(
                    `${o.targetUUID}-align-self`,
                    o.alignments.block
                  );
                }),
                { animationFrame: e }
              );
            } else
              on(
                a,
                l,
                () => R(null, null, function* () {
                  const u = yield V.getElementRects({
                    reference: a,
                    floating: l,
                    strategy: "absolute"
                  }), i = yield nn({
                    targetEl: l,
                    targetProperty: s,
                    anchorRect: u.reference,
                    anchorSide: o.anchorSide,
                    anchorSize: o.anchorSize,
                    fallback: o.fallbackValue
                  });
                  n.style.setProperty(o.uuid, i);
                }),
                { animationFrame: e }
              );
          else if (Zp(o)) {
            const u = yield nn({
              targetProperty: s,
              anchorSide: o.anchorSide,
              anchorSize: o.anchorSize,
              fallback: o.fallbackValue
            });
            n.style.setProperty(o.uuid, u);
          }
        }
    });
  }
  function Is(t, e) {
    return R(this, null, function* () {
      const n = yield V.getElementRects({
        reference: t,
        floating: t,
        strategy: "absolute"
      });
      return yield Zo(
        {
          x: t.offsetLeft,
          y: t.offsetTop,
          platform: qp,
          rects: n,
          elements: {
            floating: t,
            reference: e
          },
          strategy: "absolute"
        },
        {
          padding: Xp(t)
        }
      );
    });
  }
  function ed(t, e, n = !1) {
    return R(this, null, function* () {
      if (!e.length)
        return;
      const s = document.querySelectorAll(t);
      for (const r of s) {
        let o = !1;
        const a = yield En(r);
        on(
          // We're just checking whether the target element overflows, so we don't
          // care about the position of the anchor element in this case. Passing in
          // an empty object instead of a reference element avoids unnecessarily
          // watching for irrelevant changes.
          {},
          r,
          () => R(null, null, function* () {
            if (o)
              return;
            o = !0, r.removeAttribute("data-anchor-polyfill");
            const l = yield Is(r, a);
            if (Object.values(l).every((u) => u <= 0)) {
              r.removeAttribute("data-anchor-polyfill-last-successful"), o = !1;
              return;
            }
            for (const [u, { uuid: i }] of e.entries()) {
              r.setAttribute("data-anchor-polyfill", i);
              const c = yield Is(r, a);
              if (Object.values(c).every((h) => h <= 0)) {
                r.setAttribute("data-anchor-polyfill-last-successful", i), o = !1;
                break;
              }
              if (u === e.length - 1) {
                const h = r.getAttribute(
                  "data-anchor-polyfill-last-successful"
                );
                h ? r.setAttribute("data-anchor-polyfill", h) : r.removeAttribute("data-anchor-polyfill"), o = !1;
                break;
              }
            }
          }),
          { animationFrame: n, layoutShift: !1 }
        );
      }
    });
  }
  function nd(t, e = !1) {
    return R(this, null, function* () {
      var n, s;
      for (const r of Object.values(t))
        yield td((n = r.declarations) != null ? n : {}, e);
      for (const [r, o] of Object.entries(t))
        yield ed(
          r,
          (s = o.fallbacks) != null ? s : [],
          e
        );
    });
  }
  function sd(t = {}) {
    const e = typeof t == "boolean" ? { useAnimationFrame: t } : t, n = e.useAnimationFrame === void 0 ? !!window.UPDATE_ANCHOR_ON_ANIMATION_FRAME : e.useAnimationFrame;
    return Array.isArray(e.elements) || (e.elements = void 0), (!Array.isArray(e.roots) || e.roots.length === 0) && (e.roots = [document]), Object.assign(e, {
      useAnimationFrame: n
    });
  }
  function Ns(t) {
    return R(this, null, function* () {
      const e = sd(
        window.ANCHOR_POSITIONING_POLYFILL_OPTIONS
      );
      let n = yield Mf(e), s = {}, r;
      Af();
      try {
        Ef(n) && (n = Os(n));
        const a = yield Hp(n, { roots: e.roots });
        s = a.rules, r = a.inlineStyles;
      } catch (o) {
        throw Tf(), o;
      }
      return Object.values(s).length && (Os(n, r, !0), yield nd(s, e.useAnimationFrame)), s;
    });
  }
  document.readyState !== "complete" ? window.addEventListener("load", () => {
    Ns();
  }) : Ns();
});
export default rd();
//# sourceMappingURL=css-anchor-positioning.js.map
