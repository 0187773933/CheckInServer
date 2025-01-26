"use strict";

function _getRequireWildcardCache(e) { if ("function" != typeof WeakMap) return null; var r = new WeakMap(), t = new WeakMap(); return (_getRequireWildcardCache = function _getRequireWildcardCache(e) { return e ? t : r; })(e); }
function _interopRequireWildcard(e, r) { if (!r && e && e.__esModule) return e; if (null === e || "object" != _typeof(e) && "function" != typeof e) return { "default": e }; var t = _getRequireWildcardCache(r); if (t && t.has(e)) return t.get(e); var n = { __proto__: null }, a = Object.defineProperty && Object.getOwnPropertyDescriptor; for (var u in e) if ("default" !== u && {}.hasOwnProperty.call(e, u)) { var i = a ? Object.getOwnPropertyDescriptor(e, u) : null; i && (i.get || i.set) ? Object.defineProperty(n, u, i) : n[u] = e[u]; } return n["default"] = e, t && t.set(e, n), n; }
function _regeneratorRuntime() { "use strict"; /*! regenerator-runtime -- Copyright (c) 2014-present, Facebook, Inc. -- license (MIT): https://github.com/facebook/regenerator/blob/main/LICENSE */ _regeneratorRuntime = function _regeneratorRuntime() { return e; }; var t, e = {}, r = Object.prototype, n = r.hasOwnProperty, o = Object.defineProperty || function (t, e, r) { t[e] = r.value; }, i = "function" == typeof Symbol ? Symbol : {}, a = i.iterator || "@@iterator", c = i.asyncIterator || "@@asyncIterator", u = i.toStringTag || "@@toStringTag"; function define(t, e, r) { return Object.defineProperty(t, e, { value: r, enumerable: !0, configurable: !0, writable: !0 }), t[e]; } try { define({}, ""); } catch (t) { define = function define(t, e, r) { return t[e] = r; }; } function wrap(t, e, r, n) { var i = e && e.prototype instanceof Generator ? e : Generator, a = Object.create(i.prototype), c = new Context(n || []); return o(a, "_invoke", { value: makeInvokeMethod(t, r, c) }), a; } function tryCatch(t, e, r) { try { return { type: "normal", arg: t.call(e, r) }; } catch (t) { return { type: "throw", arg: t }; } } e.wrap = wrap; var h = "suspendedStart", l = "suspendedYield", f = "executing", s = "completed", y = {}; function Generator() {} function GeneratorFunction() {} function GeneratorFunctionPrototype() {} var p = {}; define(p, a, function () { return this; }); var d = Object.getPrototypeOf, v = d && d(d(values([]))); v && v !== r && n.call(v, a) && (p = v); var g = GeneratorFunctionPrototype.prototype = Generator.prototype = Object.create(p); function defineIteratorMethods(t) { ["next", "throw", "return"].forEach(function (e) { define(t, e, function (t) { return this._invoke(e, t); }); }); } function AsyncIterator(t, e) { function invoke(r, o, i, a) { var c = tryCatch(t[r], t, o); if ("throw" !== c.type) { var u = c.arg, h = u.value; return h && "object" == _typeof(h) && n.call(h, "__await") ? e.resolve(h.__await).then(function (t) { invoke("next", t, i, a); }, function (t) { invoke("throw", t, i, a); }) : e.resolve(h).then(function (t) { u.value = t, i(u); }, function (t) { return invoke("throw", t, i, a); }); } a(c.arg); } var r; o(this, "_invoke", { value: function value(t, n) { function callInvokeWithMethodAndArg() { return new e(function (e, r) { invoke(t, n, e, r); }); } return r = r ? r.then(callInvokeWithMethodAndArg, callInvokeWithMethodAndArg) : callInvokeWithMethodAndArg(); } }); } function makeInvokeMethod(e, r, n) { var o = h; return function (i, a) { if (o === f) throw Error("Generator is already running"); if (o === s) { if ("throw" === i) throw a; return { value: t, done: !0 }; } for (n.method = i, n.arg = a;;) { var c = n.delegate; if (c) { var u = maybeInvokeDelegate(c, n); if (u) { if (u === y) continue; return u; } } if ("next" === n.method) n.sent = n._sent = n.arg;else if ("throw" === n.method) { if (o === h) throw o = s, n.arg; n.dispatchException(n.arg); } else "return" === n.method && n.abrupt("return", n.arg); o = f; var p = tryCatch(e, r, n); if ("normal" === p.type) { if (o = n.done ? s : l, p.arg === y) continue; return { value: p.arg, done: n.done }; } "throw" === p.type && (o = s, n.method = "throw", n.arg = p.arg); } }; } function maybeInvokeDelegate(e, r) { var n = r.method, o = e.iterator[n]; if (o === t) return r.delegate = null, "throw" === n && e.iterator["return"] && (r.method = "return", r.arg = t, maybeInvokeDelegate(e, r), "throw" === r.method) || "return" !== n && (r.method = "throw", r.arg = new TypeError("The iterator does not provide a '" + n + "' method")), y; var i = tryCatch(o, e.iterator, r.arg); if ("throw" === i.type) return r.method = "throw", r.arg = i.arg, r.delegate = null, y; var a = i.arg; return a ? a.done ? (r[e.resultName] = a.value, r.next = e.nextLoc, "return" !== r.method && (r.method = "next", r.arg = t), r.delegate = null, y) : a : (r.method = "throw", r.arg = new TypeError("iterator result is not an object"), r.delegate = null, y); } function pushTryEntry(t) { var e = { tryLoc: t[0] }; 1 in t && (e.catchLoc = t[1]), 2 in t && (e.finallyLoc = t[2], e.afterLoc = t[3]), this.tryEntries.push(e); } function resetTryEntry(t) { var e = t.completion || {}; e.type = "normal", delete e.arg, t.completion = e; } function Context(t) { this.tryEntries = [{ tryLoc: "root" }], t.forEach(pushTryEntry, this), this.reset(!0); } function values(e) { if (e || "" === e) { var r = e[a]; if (r) return r.call(e); if ("function" == typeof e.next) return e; if (!isNaN(e.length)) { var o = -1, i = function next() { for (; ++o < e.length;) if (n.call(e, o)) return next.value = e[o], next.done = !1, next; return next.value = t, next.done = !0, next; }; return i.next = i; } } throw new TypeError(_typeof(e) + " is not iterable"); } return GeneratorFunction.prototype = GeneratorFunctionPrototype, o(g, "constructor", { value: GeneratorFunctionPrototype, configurable: !0 }), o(GeneratorFunctionPrototype, "constructor", { value: GeneratorFunction, configurable: !0 }), GeneratorFunction.displayName = define(GeneratorFunctionPrototype, u, "GeneratorFunction"), e.isGeneratorFunction = function (t) { var e = "function" == typeof t && t.constructor; return !!e && (e === GeneratorFunction || "GeneratorFunction" === (e.displayName || e.name)); }, e.mark = function (t) { return Object.setPrototypeOf ? Object.setPrototypeOf(t, GeneratorFunctionPrototype) : (t.__proto__ = GeneratorFunctionPrototype, define(t, u, "GeneratorFunction")), t.prototype = Object.create(g), t; }, e.awrap = function (t) { return { __await: t }; }, defineIteratorMethods(AsyncIterator.prototype), define(AsyncIterator.prototype, c, function () { return this; }), e.AsyncIterator = AsyncIterator, e.async = function (t, r, n, o, i) { void 0 === i && (i = Promise); var a = new AsyncIterator(wrap(t, r, n, o), i); return e.isGeneratorFunction(r) ? a : a.next().then(function (t) { return t.done ? t.value : a.next(); }); }, defineIteratorMethods(g), define(g, u, "Generator"), define(g, a, function () { return this; }), define(g, "toString", function () { return "[object Generator]"; }), e.keys = function (t) { var e = Object(t), r = []; for (var n in e) r.push(n); return r.reverse(), function next() { for (; r.length;) { var t = r.pop(); if (t in e) return next.value = t, next.done = !1, next; } return next.done = !0, next; }; }, e.values = values, Context.prototype = { constructor: Context, reset: function reset(e) { if (this.prev = 0, this.next = 0, this.sent = this._sent = t, this.done = !1, this.delegate = null, this.method = "next", this.arg = t, this.tryEntries.forEach(resetTryEntry), !e) for (var r in this) "t" === r.charAt(0) && n.call(this, r) && !isNaN(+r.slice(1)) && (this[r] = t); }, stop: function stop() { this.done = !0; var t = this.tryEntries[0].completion; if ("throw" === t.type) throw t.arg; return this.rval; }, dispatchException: function dispatchException(e) { if (this.done) throw e; var r = this; function handle(n, o) { return a.type = "throw", a.arg = e, r.next = n, o && (r.method = "next", r.arg = t), !!o; } for (var o = this.tryEntries.length - 1; o >= 0; --o) { var i = this.tryEntries[o], a = i.completion; if ("root" === i.tryLoc) return handle("end"); if (i.tryLoc <= this.prev) { var c = n.call(i, "catchLoc"), u = n.call(i, "finallyLoc"); if (c && u) { if (this.prev < i.catchLoc) return handle(i.catchLoc, !0); if (this.prev < i.finallyLoc) return handle(i.finallyLoc); } else if (c) { if (this.prev < i.catchLoc) return handle(i.catchLoc, !0); } else { if (!u) throw Error("try statement without catch or finally"); if (this.prev < i.finallyLoc) return handle(i.finallyLoc); } } } }, abrupt: function abrupt(t, e) { for (var r = this.tryEntries.length - 1; r >= 0; --r) { var o = this.tryEntries[r]; if (o.tryLoc <= this.prev && n.call(o, "finallyLoc") && this.prev < o.finallyLoc) { var i = o; break; } } i && ("break" === t || "continue" === t) && i.tryLoc <= e && e <= i.finallyLoc && (i = null); var a = i ? i.completion : {}; return a.type = t, a.arg = e, i ? (this.method = "next", this.next = i.finallyLoc, y) : this.complete(a); }, complete: function complete(t, e) { if ("throw" === t.type) throw t.arg; return "break" === t.type || "continue" === t.type ? this.next = t.arg : "return" === t.type ? (this.rval = this.arg = t.arg, this.method = "return", this.next = "end") : "normal" === t.type && e && (this.next = e), y; }, finish: function finish(t) { for (var e = this.tryEntries.length - 1; e >= 0; --e) { var r = this.tryEntries[e]; if (r.finallyLoc === t) return this.complete(r.completion, r.afterLoc), resetTryEntry(r), y; } }, "catch": function _catch(t) { for (var e = this.tryEntries.length - 1; e >= 0; --e) { var r = this.tryEntries[e]; if (r.tryLoc === t) { var n = r.completion; if ("throw" === n.type) { var o = n.arg; resetTryEntry(r); } return o; } } throw Error("illegal catch attempt"); }, delegateYield: function delegateYield(e, r, n) { return this.delegate = { iterator: values(e), resultName: r, nextLoc: n }, "next" === this.method && (this.arg = t), y; } }, e; }
function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
function asyncGeneratorStep(n, t, e, r, o, a, c) { try { var i = n[a](c), u = i.value; } catch (n) { return void e(n); } i.done ? t(u) : Promise.resolve(u).then(r, o); }
function _asyncToGenerator(n) { return function () { var t = this, e = arguments; return new Promise(function (r, o) { var a = n.apply(t, e); function _next(n) { asyncGeneratorStep(a, r, o, _next, _throw, "next", n); } function _throw(n) { asyncGeneratorStep(a, r, o, _next, _throw, "throw", n); } _next(void 0); }); }; }
function _slicedToArray(r, e) { return _arrayWithHoles(r) || _iterableToArrayLimit(r, e) || _unsupportedIterableToArray(r, e) || _nonIterableRest(); }
function _nonIterableRest() { throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); }
function _iterableToArrayLimit(r, l) { var t = null == r ? null : "undefined" != typeof Symbol && r[Symbol.iterator] || r["@@iterator"]; if (null != t) { var e, n, i, u, a = [], f = !0, o = !1; try { if (i = (t = t.call(r)).next, 0 === l) { if (Object(t) !== t) return; f = !1; } else for (; !(f = (e = i.call(t)).done) && (a.push(e.value), a.length !== l); f = !0); } catch (r) { o = !0, n = r; } finally { try { if (!f && null != t["return"] && (u = t["return"](), Object(u) !== u)) return; } finally { if (o) throw n; } } return a; } }
function _arrayWithHoles(r) { if (Array.isArray(r)) return r; }
function _defineProperties(e, r) { for (var t = 0; t < r.length; t++) { var o = r[t]; o.enumerable = o.enumerable || !1, o.configurable = !0, "value" in o && (o.writable = !0), Object.defineProperty(e, _toPropertyKey(o.key), o); } }
function _createClass(e, r, t) { return r && _defineProperties(e.prototype, r), t && _defineProperties(e, t), Object.defineProperty(e, "prototype", { writable: !1 }), e; }
function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == _typeof(i) ? i : i + ""; }
function _toPrimitive(t, r) { if ("object" != _typeof(t) || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != _typeof(i)) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
function _classCallCheck(a, n) { if (!(a instanceof n)) throw new TypeError("Cannot call a class as a function"); }
function _callSuper(t, o, e) { return o = _getPrototypeOf(o), _possibleConstructorReturn(t, _isNativeReflectConstruct() ? Reflect.construct(o, e || [], _getPrototypeOf(t).constructor) : o.apply(t, e)); }
function _possibleConstructorReturn(t, e) { if (e && ("object" == _typeof(e) || "function" == typeof e)) return e; if (void 0 !== e) throw new TypeError("Derived constructors may only return object or undefined"); return _assertThisInitialized(t); }
function _assertThisInitialized(e) { if (void 0 === e) throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); return e; }
function _inherits(t, e) { if ("function" != typeof e && null !== e) throw new TypeError("Super expression must either be null or a function"); t.prototype = Object.create(e && e.prototype, { constructor: { value: t, writable: !0, configurable: !0 } }), Object.defineProperty(t, "prototype", { writable: !1 }), e && _setPrototypeOf(t, e); }
function _wrapNativeSuper(t) { var r = "function" == typeof Map ? new Map() : void 0; return _wrapNativeSuper = function _wrapNativeSuper(t) { if (null === t || !_isNativeFunction(t)) return t; if ("function" != typeof t) throw new TypeError("Super expression must either be null or a function"); if (void 0 !== r) { if (r.has(t)) return r.get(t); r.set(t, Wrapper); } function Wrapper() { return _construct(t, arguments, _getPrototypeOf(this).constructor); } return Wrapper.prototype = Object.create(t.prototype, { constructor: { value: Wrapper, enumerable: !1, writable: !0, configurable: !0 } }), _setPrototypeOf(Wrapper, t); }, _wrapNativeSuper(t); }
function _construct(t, e, r) { if (_isNativeReflectConstruct()) return Reflect.construct.apply(null, arguments); var o = [null]; o.push.apply(o, e); var p = new (t.bind.apply(t, o))(); return r && _setPrototypeOf(p, r.prototype), p; }
function _isNativeReflectConstruct() { try { var t = !Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], function () {})); } catch (t) {} return (_isNativeReflectConstruct = function _isNativeReflectConstruct() { return !!t; })(); }
function _isNativeFunction(t) { try { return -1 !== Function.toString.call(t).indexOf("[native code]"); } catch (n) { return "function" == typeof t; } }
function _setPrototypeOf(t, e) { return _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function (t, e) { return t.__proto__ = e, t; }, _setPrototypeOf(t, e); }
function _getPrototypeOf(t) { return _getPrototypeOf = Object.setPrototypeOf ? Object.getPrototypeOf.bind() : function (t) { return t.__proto__ || Object.getPrototypeOf(t); }, _getPrototypeOf(t); }
function _createForOfIteratorHelper(r, e) { var t = "undefined" != typeof Symbol && r[Symbol.iterator] || r["@@iterator"]; if (!t) { if (Array.isArray(r) || (t = _unsupportedIterableToArray(r)) || e && r && "number" == typeof r.length) { t && (r = t); var _n = 0, F = function F() {}; return { s: F, n: function n() { return _n >= r.length ? { done: !0 } : { done: !1, value: r[_n++] }; }, e: function e(r) { throw r; }, f: F }; } throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); } var o, a = !0, u = !1; return { s: function s() { t = t.call(r); }, n: function n() { var r = t.next(); return a = r.done, r; }, e: function e(r) { u = !0, o = r; }, f: function f() { try { a || null == t["return"] || t["return"](); } finally { if (u) throw o; } } }; }
function _unsupportedIterableToArray(r, a) { if (r) { if ("string" == typeof r) return _arrayLikeToArray(r, a); var t = {}.toString.call(r).slice(8, -1); return "Object" === t && r.constructor && (t = r.constructor.name), "Map" === t || "Set" === t ? Array.from(r) : "Arguments" === t || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t) ? _arrayLikeToArray(r, a) : void 0; } }
function _arrayLikeToArray(r, a) { (null == a || a > r.length) && (a = r.length); for (var e = 0, n = Array(a); e < a; e++) n[e] = r[e]; return n; }
function _typeof(o) { "@babel/helpers - typeof"; return _typeof = "function" == typeof Symbol && "symbol" == typeof Symbol.iterator ? function (o) { return typeof o; } : function (o) { return o && "function" == typeof Symbol && o.constructor === Symbol && o !== Symbol.prototype ? "symbol" : typeof o; }, _typeof(o); }
var Kyber = function () {
  var __create = Object.create;
  var __defProp = Object.defineProperty;
  var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
  var __getOwnPropNames = Object.getOwnPropertyNames;
  var __getProtoOf = Object.getPrototypeOf;
  var __hasOwnProp = Object.prototype.hasOwnProperty;
  var __require = /* @__PURE__ */function (x) {
    return typeof require !== "undefined" ? require : typeof Proxy !== "undefined" ? new Proxy(x, {
      get: function get(a, b) {
        return (typeof require !== "undefined" ? require : a)[b];
      }
    }) : x;
  }(function (x) {
    if (typeof require !== "undefined") return require.apply(this, arguments);
    throw Error('Dynamic require of "' + x + '" is not supported');
  });
  var __export = function __export(target, all) {
    for (var name in all) __defProp(target, name, {
      get: all[name],
      enumerable: true
    });
  };
  var __copyProps = function __copyProps(to, from, except, desc) {
    if (from && _typeof(from) === "object" || typeof from === "function") {
      var _iterator = _createForOfIteratorHelper(__getOwnPropNames(from)),
        _step;
      try {
        var _loop = function _loop() {
          var key = _step.value;
          if (!__hasOwnProp.call(to, key) && key !== except) __defProp(to, key, {
            get: function get() {
              return from[key];
            },
            enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable
          });
        };
        for (_iterator.s(); !(_step = _iterator.n()).done;) {
          _loop();
        }
      } catch (err) {
        _iterator.e(err);
      } finally {
        _iterator.f();
      }
    }
    return to;
  };
  var __toESM = function __toESM(mod, isNodeMode, target) {
    return target = mod != null ? __create(__getProtoOf(mod)) : {}, __copyProps(
    // If the importer is in node compatibility mode or this is not an ESM
    // file that has been converted to a CommonJS file using a Babel-
    // compatible transform (i.e. "__esModule" has not been set), then set
    // "default" to the CommonJS "module.exports" for node compatibility.
    isNodeMode || !mod || !mod.__esModule ? __defProp(target, "default", {
      value: mod,
      enumerable: true
    }) : target, mod);
  };
  var __toCommonJS = function __toCommonJS(mod) {
    return __copyProps(__defProp({}, "__esModule", {
      value: true
    }), mod);
  };

  // mod.ts
  var mod_exports = {};
  __export(mod_exports, {
    MlKem1024: function MlKem1024() {
      return _MlKem;
    },
    MlKem512: function MlKem512() {
      return _MlKem2;
    },
    MlKem768: function MlKem768() {
      return _MlKem3;
    },
    MlKemError: function MlKemError() {
      return _MlKemError;
    }
  });

  // src/errors.ts
  var _MlKemError = /*#__PURE__*/function (_Error) {
    function _MlKemError(e) {
      var _this;
      _classCallCheck(this, _MlKemError);
      var message;
      if (e instanceof Error) {
        message = e.message;
      } else if (typeof e === "string") {
        message = e;
      } else {
        message = "";
      }
      _this = _callSuper(this, _MlKemError, [message]);
      _this.name = _this.constructor.name;
      return _this;
    }
    _inherits(_MlKemError, _Error);
    return _createClass(_MlKemError);
  }(/*#__PURE__*/_wrapNativeSuper(Error));

  // src/consts.ts
  var N = 256;
  var Q = 3329;
  var Q_INV = 62209;
  var NTT_ZETAS = [2285, 2571, 2970, 1812, 1493, 1422, 287, 202, 3158, 622, 1577, 182, 962, 2127, 1855, 1468, 573, 2004, 264, 383, 2500, 1458, 1727, 3199, 2648, 1017, 732, 608, 1787, 411, 3124, 1758, 1223, 652, 2777, 1015, 2036, 1491, 3047, 1785, 516, 3321, 3009, 2663, 1711, 2167, 126, 1469, 2476, 3239, 3058, 830, 107, 1908, 3082, 2378, 2931, 961, 1821, 2604, 448, 2264, 677, 2054, 2226, 430, 555, 843, 2078, 871, 1550, 105, 422, 587, 177, 3094, 3038, 2869, 1574, 1653, 3083, 778, 1159, 3182, 2552, 1483, 2727, 1119, 1739, 644, 2457, 349, 418, 329, 3173, 3254, 817, 1097, 603, 610, 1322, 2044, 1864, 384, 2114, 3193, 1218, 1994, 2455, 220, 2142, 1670, 2144, 1799, 2051, 794, 1819, 2475, 2459, 478, 3221, 3021, 996, 991, 958, 1869, 1522, 1628];
  var NTT_ZETAS_INV = [1701, 1807, 1460, 2371, 2338, 2333, 308, 108, 2851, 870, 854, 1510, 2535, 1278, 1530, 1185, 1659, 1187, 3109, 874, 1335, 2111, 136, 1215, 2945, 1465, 1285, 2007, 2719, 2726, 2232, 2512, 75, 156, 3e3, 2911, 2980, 872, 2685, 1590, 2210, 602, 1846, 777, 147, 2170, 2551, 246, 1676, 1755, 460, 291, 235, 3152, 2742, 2907, 3224, 1779, 2458, 1251, 2486, 2774, 2899, 1103, 1275, 2652, 1065, 2881, 725, 1508, 2368, 398, 951, 247, 1421, 3222, 2499, 271, 90, 853, 1860, 3203, 1162, 1618, 666, 320, 8, 2813, 1544, 282, 1838, 1293, 2314, 552, 2677, 2106, 1571, 205, 2918, 1542, 2721, 2597, 2312, 681, 130, 1602, 1871, 829, 2946, 3065, 1325, 2756, 1861, 1474, 1202, 2367, 3147, 1752, 2707, 171, 3127, 3042, 1907, 1836, 1517, 359, 758, 1441];

  // node_modules/@noble/hashes/esm/_assert.js
  function anumber(n) {
    if (!Number.isSafeInteger(n) || n < 0) throw new Error("positive integer expected, got " + n);
  }
  function isBytes(a) {
    return a instanceof Uint8Array || ArrayBuffer.isView(a) && a.constructor.name === "Uint8Array";
  }
  function abytes(b) {
    if (!isBytes(b)) throw new Error("Uint8Array expected");
    for (var _len = arguments.length, lengths = new Array(_len > 1 ? _len - 1 : 0), _key = 1; _key < _len; _key++) {
      lengths[_key - 1] = arguments[_key];
    }
    if (lengths.length > 0 && !lengths.includes(b.length)) throw new Error("Uint8Array expected of length " + lengths + ", got length=" + b.length);
  }
  function aexists(instance) {
    var checkFinished = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : true;
    if (instance.destroyed) throw new Error("Hash instance has been destroyed");
    if (checkFinished && instance.finished) throw new Error("Hash#digest() has already been called");
  }
  function aoutput(out, instance) {
    abytes(out);
    var min = instance.outputLen;
    if (out.length < min) {
      throw new Error("digestInto() expects output buffer of length at least " + min);
    }
  }

  // node_modules/@noble/hashes/esm/_u64.js
  var U32_MASK64 = /* @__PURE__ */BigInt(Math.pow(2, 32) - 1);
  var _32n = /* @__PURE__ */BigInt(32);
  function fromBig(n) {
    var le = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : false;
    if (le) return {
      h: Number(n & U32_MASK64),
      l: Number(n >> _32n & U32_MASK64)
    };
    return {
      h: Number(n >> _32n & U32_MASK64) | 0,
      l: Number(n & U32_MASK64) | 0
    };
  }
  function split(lst) {
    var le = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : false;
    var Ah = new Uint32Array(lst.length);
    var Al = new Uint32Array(lst.length);
    for (var i = 0; i < lst.length; i++) {
      var _fromBig = fromBig(lst[i], le),
        h2 = _fromBig.h,
        l = _fromBig.l;
      var _ref = [h2, l];
      Ah[i] = _ref[0];
      Al[i] = _ref[1];
    }
    return [Ah, Al];
  }
  var rotlSH = function rotlSH(h2, l, s) {
    return h2 << s | l >>> 32 - s;
  };
  var rotlSL = function rotlSL(h2, l, s) {
    return l << s | h2 >>> 32 - s;
  };
  var rotlBH = function rotlBH(h2, l, s) {
    return l << s - 32 | h2 >>> 64 - s;
  };
  var rotlBL = function rotlBL(h2, l, s) {
    return h2 << s - 32 | l >>> 64 - s;
  };

  // node_modules/@noble/hashes/esm/utils.js
  function u32(arr) {
    return new Uint32Array(arr.buffer, arr.byteOffset, Math.floor(arr.byteLength / 4));
  }
  var isLE = /* @__PURE__ */function () {
    return new Uint8Array(new Uint32Array([287454020]).buffer)[0] === 68;
  }();
  function byteSwap(word) {
    return word << 24 & 4278190080 | word << 8 & 16711680 | word >>> 8 & 65280 | word >>> 24 & 255;
  }
  function byteSwap32(arr) {
    for (var i = 0; i < arr.length; i++) {
      arr[i] = byteSwap(arr[i]);
    }
  }
  function utf8ToBytes(str) {
    if (typeof str !== "string") throw new Error("utf8ToBytes expected string, got " + _typeof(str));
    return new Uint8Array(new TextEncoder().encode(str));
  }
  function toBytes(data) {
    if (typeof data === "string") data = utf8ToBytes(data);
    abytes(data);
    return data;
  }
  var Hash = /*#__PURE__*/function () {
    function Hash() {
      _classCallCheck(this, Hash);
    }
    return _createClass(Hash, [{
      key: "clone",
      value:
      // Safe version that clones internal state
      function clone() {
        return this._cloneInto();
      }
    }]);
  }();
  function wrapConstructor(hashCons) {
    var hashC = function hashC(msg) {
      return hashCons().update(toBytes(msg)).digest();
    };
    var tmp = hashCons();
    hashC.outputLen = tmp.outputLen;
    hashC.blockLen = tmp.blockLen;
    hashC.create = function () {
      return hashCons();
    };
    return hashC;
  }
  function wrapXOFConstructorWithOpts(hashCons) {
    var hashC = function hashC(msg, opts) {
      return hashCons(opts).update(toBytes(msg)).digest();
    };
    var tmp = hashCons({});
    hashC.outputLen = tmp.outputLen;
    hashC.blockLen = tmp.blockLen;
    hashC.create = function (opts) {
      return hashCons(opts);
    };
    return hashC;
  }

  // node_modules/@noble/hashes/esm/sha3.js
  var SHA3_PI = [];
  var SHA3_ROTL = [];
  var _SHA3_IOTA = [];
  var _0n = /* @__PURE__ */BigInt(0);
  var _1n = /* @__PURE__ */BigInt(1);
  var _2n = /* @__PURE__ */BigInt(2);
  var _7n = /* @__PURE__ */BigInt(7);
  var _256n = /* @__PURE__ */BigInt(256);
  var _0x71n = /* @__PURE__ */BigInt(113);
  for (var round = 0, R = _1n, x = 1, y = 0; round < 24; round++) {
    var _ref2 = [y, (2 * x + 3 * y) % 5];
    x = _ref2[0];
    y = _ref2[1];
    SHA3_PI.push(2 * (5 * y + x));
    SHA3_ROTL.push((round + 1) * (round + 2) / 2 % 64);
    var t = _0n;
    for (var j = 0; j < 7; j++) {
      R = (R << _1n ^ (R >> _7n) * _0x71n) % _256n;
      if (R & _2n) t ^= _1n << (_1n << /* @__PURE__ */BigInt(j)) - _1n;
    }
    _SHA3_IOTA.push(t);
  }
  var _split = /* @__PURE__ */split(_SHA3_IOTA, true),
    _split2 = _slicedToArray(_split, 2),
    SHA3_IOTA_H = _split2[0],
    SHA3_IOTA_L = _split2[1];
  var rotlH = function rotlH(h2, l, s) {
    return s > 32 ? rotlBH(h2, l, s) : rotlSH(h2, l, s);
  };
  var rotlL = function rotlL(h2, l, s) {
    return s > 32 ? rotlBL(h2, l, s) : rotlSL(h2, l, s);
  };
  function keccakP(s) {
    var rounds = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 24;
    var B = new Uint32Array(5 * 2);
    for (var _round = 24 - rounds; _round < 24; _round++) {
      for (var _x = 0; _x < 10; _x++) B[_x] = s[_x] ^ s[_x + 10] ^ s[_x + 20] ^ s[_x + 30] ^ s[_x + 40];
      for (var _x2 = 0; _x2 < 10; _x2 += 2) {
        var idx1 = (_x2 + 8) % 10;
        var idx0 = (_x2 + 2) % 10;
        var B0 = B[idx0];
        var B1 = B[idx0 + 1];
        var Th = rotlH(B0, B1, 1) ^ B[idx1];
        var Tl = rotlL(B0, B1, 1) ^ B[idx1 + 1];
        for (var _y = 0; _y < 50; _y += 10) {
          s[_x2 + _y] ^= Th;
          s[_x2 + _y + 1] ^= Tl;
        }
      }
      var curH = s[2];
      var curL = s[3];
      for (var _t = 0; _t < 24; _t++) {
        var shift = SHA3_ROTL[_t];
        var _Th = rotlH(curH, curL, shift);
        var _Tl = rotlL(curH, curL, shift);
        var PI = SHA3_PI[_t];
        curH = s[PI];
        curL = s[PI + 1];
        s[PI] = _Th;
        s[PI + 1] = _Tl;
      }
      for (var _y2 = 0; _y2 < 50; _y2 += 10) {
        for (var _x3 = 0; _x3 < 10; _x3++) B[_x3] = s[_y2 + _x3];
        for (var _x4 = 0; _x4 < 10; _x4++) s[_y2 + _x4] ^= ~B[(_x4 + 2) % 10] & B[(_x4 + 4) % 10];
      }
      s[0] ^= SHA3_IOTA_H[_round];
      s[1] ^= SHA3_IOTA_L[_round];
    }
    B.fill(0);
  }
  var Keccak = /*#__PURE__*/function (_Hash) {
    // NOTE: we accept arguments in bytes instead of bits here.
    function _Keccak(blockLen, suffix, outputLen) {
      var _this2;
      var enableXOF = arguments.length > 3 && arguments[3] !== undefined ? arguments[3] : false;
      var rounds = arguments.length > 4 && arguments[4] !== undefined ? arguments[4] : 24;
      _classCallCheck(this, _Keccak);
      _this2 = _callSuper(this, _Keccak);
      _this2.blockLen = blockLen;
      _this2.suffix = suffix;
      _this2.outputLen = outputLen;
      _this2.enableXOF = enableXOF;
      _this2.rounds = rounds;
      _this2.pos = 0;
      _this2.posOut = 0;
      _this2.finished = false;
      _this2.destroyed = false;
      anumber(outputLen);
      if (0 >= _this2.blockLen || _this2.blockLen >= 200) throw new Error("Sha3 supports only keccak-f1600 function");
      _this2.state = new Uint8Array(200);
      _this2.state32 = u32(_this2.state);
      return _this2;
    }
    _inherits(_Keccak, _Hash);
    return _createClass(_Keccak, [{
      key: "keccak",
      value: function keccak() {
        if (!isLE) byteSwap32(this.state32);
        keccakP(this.state32, this.rounds);
        if (!isLE) byteSwap32(this.state32);
        this.posOut = 0;
        this.pos = 0;
      }
    }, {
      key: "update",
      value: function update(data) {
        aexists(this);
        var blockLen = this.blockLen,
          state = this.state;
        data = toBytes(data);
        var len = data.length;
        for (var pos = 0; pos < len;) {
          var take = Math.min(blockLen - this.pos, len - pos);
          for (var i = 0; i < take; i++) state[this.pos++] ^= data[pos++];
          if (this.pos === blockLen) this.keccak();
        }
        return this;
      }
    }, {
      key: "finish",
      value: function finish() {
        if (this.finished) return;
        this.finished = true;
        var state = this.state,
          suffix = this.suffix,
          pos = this.pos,
          blockLen = this.blockLen;
        state[pos] ^= suffix;
        if ((suffix & 128) !== 0 && pos === blockLen - 1) this.keccak();
        state[blockLen - 1] ^= 128;
        this.keccak();
      }
    }, {
      key: "writeInto",
      value: function writeInto(out) {
        aexists(this, false);
        abytes(out);
        this.finish();
        var bufferOut = this.state;
        var blockLen = this.blockLen;
        for (var pos = 0, len = out.length; pos < len;) {
          if (this.posOut >= blockLen) this.keccak();
          var take = Math.min(blockLen - this.posOut, len - pos);
          out.set(bufferOut.subarray(this.posOut, this.posOut + take), pos);
          this.posOut += take;
          pos += take;
        }
        return out;
      }
    }, {
      key: "xofInto",
      value: function xofInto(out) {
        if (!this.enableXOF) throw new Error("XOF is not possible for this instance");
        return this.writeInto(out);
      }
    }, {
      key: "xof",
      value: function xof(bytes) {
        anumber(bytes);
        return this.xofInto(new Uint8Array(bytes));
      }
    }, {
      key: "digestInto",
      value: function digestInto(out) {
        aoutput(out, this);
        if (this.finished) throw new Error("digest() was already called");
        this.writeInto(out);
        this.destroy();
        return out;
      }
    }, {
      key: "digest",
      value: function digest() {
        return this.digestInto(new Uint8Array(this.outputLen));
      }
    }, {
      key: "destroy",
      value: function destroy() {
        this.destroyed = true;
        this.state.fill(0);
      }
    }, {
      key: "_cloneInto",
      value: function _cloneInto(to) {
        var blockLen = this.blockLen,
          suffix = this.suffix,
          outputLen = this.outputLen,
          rounds = this.rounds,
          enableXOF = this.enableXOF;
        to || (to = new _Keccak(blockLen, suffix, outputLen, enableXOF, rounds));
        to.state32.set(this.state32);
        to.pos = this.pos;
        to.posOut = this.posOut;
        to.finished = this.finished;
        to.rounds = rounds;
        to.suffix = suffix;
        to.outputLen = outputLen;
        to.enableXOF = enableXOF;
        to.destroyed = this.destroyed;
        return to;
      }
    }]);
  }(Hash);
  var gen = function gen(suffix, blockLen, outputLen) {
    return wrapConstructor(function () {
      return new Keccak(blockLen, suffix, outputLen);
    });
  };
  var sha3_224 = /* @__PURE__ */gen(6, 144, 224 / 8);
  var sha3_256 = /* @__PURE__ */gen(6, 136, 256 / 8);
  var sha3_384 = /* @__PURE__ */gen(6, 104, 384 / 8);
  var sha3_512 = /* @__PURE__ */gen(6, 72, 512 / 8);
  var keccak_224 = /* @__PURE__ */gen(1, 144, 224 / 8);
  var keccak_256 = /* @__PURE__ */gen(1, 136, 256 / 8);
  var keccak_384 = /* @__PURE__ */gen(1, 104, 384 / 8);
  var keccak_512 = /* @__PURE__ */gen(1, 72, 512 / 8);
  var genShake = function genShake(suffix, blockLen, outputLen) {
    return wrapXOFConstructorWithOpts(function () {
      var opts = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
      return new Keccak(blockLen, suffix, opts.dkLen === void 0 ? outputLen : opts.dkLen, true);
    });
  };
  var shake128 = /* @__PURE__ */genShake(31, 168, 128 / 8);
  var shake256 = /* @__PURE__ */genShake(31, 136, 256 / 8);

  // src/utils.ts
  function _byte(n) {
    return n % 256;
  }
  function int16(n) {
    var end = -32768;
    var start = 32767;
    if (n >= end && n <= start) {
      return n;
    }
    if (n < end) {
      n = n + 32769;
      n = n % 65536;
      return start + n;
    }
    n = n - 32768;
    n = n % 65536;
    return end + n;
  }
  function uint16(n) {
    return n % 65536;
  }
  function int32(n) {
    var end = -2147483648;
    var start = 2147483647;
    if (n >= end && n <= start) {
      return n;
    }
    if (n < end) {
      n = n + 2147483649;
      n = n % 4294967296;
      return start + n;
    }
    n = n - 2147483648;
    n = n % 4294967296;
    return end + n;
  }
  function uint32(n) {
    return n % 4294967296;
  }
  function constantTimeCompare(x, y) {
    if (x.length != y.length) {
      return 0;
    }
    var v = new Uint8Array([0]);
    for (var i = 0; i < x.length; i++) {
      v[0] |= x[i] ^ y[i];
    }
    var z = new Uint8Array([0]);
    z[0] = ~(v[0] ^ z[0]);
    z[0] &= z[0] >> 4;
    z[0] &= z[0] >> 2;
    z[0] &= z[0] >> 1;
    return z[0];
  }
  function equalUint8Array(x, y) {
    if (x.length != y.length) {
      return false;
    }
    for (var i = 0; i < x.length; i++) {
      if (x[i] !== y[i]) {
        return false;
      }
    }
    return true;
  }
  function loadCrypto() {
    return _loadCrypto.apply(this, arguments);
  }
  function _loadCrypto() {
    _loadCrypto = _asyncToGenerator(/*#__PURE__*/_regeneratorRuntime().mark(function _callee6() {
      var _yield$import, webcrypto;
      return _regeneratorRuntime().wrap(function _callee6$(_context6) {
        while (1) switch (_context6.prev = _context6.next) {
          case 0:
            if (!(typeof globalThis !== "undefined" && globalThis.crypto !== void 0)) {
              _context6.next = 2;
              break;
            }
            return _context6.abrupt("return", globalThis.crypto);
          case 2:
            _context6.prev = 2;
            _context6.next = 5;
            return Promise.resolve().then(function () {
              return _interopRequireWildcard(require("crypto"));
            });
          case 5:
            _yield$import = _context6.sent;
            webcrypto = _yield$import.webcrypto;
            return _context6.abrupt("return", webcrypto);
          case 10:
            _context6.prev = 10;
            _context6.t0 = _context6["catch"](2);
            throw new Error("failed to load Crypto");
          case 13:
          case "end":
            return _context6.stop();
        }
      }, _callee6, null, [[2, 10]]);
    }));
    return _loadCrypto.apply(this, arguments);
  }
  function prf(len, seed, nonce) {
    return shake256.create({
      dkLen: len
    }).update(seed).update(new Uint8Array([nonce])).digest();
  }
  function byteopsLoad24(x) {
    var r = uint32(x[0]);
    r |= uint32(x[1]) << 8;
    r |= uint32(x[2]) << 16;
    return r;
  }
  function byteopsLoad32(x) {
    var r = uint32(x[0]);
    r |= uint32(x[1]) << 8;
    r |= uint32(x[2]) << 16;
    r |= uint32(x[3]) << 24;
    return uint32(r);
  }

  // src/mlKemBase.ts
  var MlKemBase = /*#__PURE__*/function () {
    /**
     * Creates a new instance of the MlKemBase class.
     */
    function MlKemBase() {
      _classCallCheck(this, MlKemBase);
      _defineProperty(this, "_api", void 0);
      _defineProperty(this, "_k", 0);
      _defineProperty(this, "_du", 0);
      _defineProperty(this, "_dv", 0);
      _defineProperty(this, "_eta1", 0);
      _defineProperty(this, "_eta2", 0);
      _defineProperty(this, "_skSize", 0);
      _defineProperty(this, "_pkSize", 0);
      _defineProperty(this, "_compressedUSize", 0);
      _defineProperty(this, "_compressedVSize", 0);
    }
    /**
     * Generates a keypair [publicKey, privateKey].
     *
     * If an error occurred, throws {@link MlKemError}.
     *
     * @returns A kaypair [publicKey, privateKey].
     * @throws {@link MlKemError}
     *
     * @example Generates a {@link MlKem768} keypair.
     *
     * ```ts
     * // Using jsr:
     * import { MlKem768 } from "@dajiaji/mlkem";
     * // Using npm:
     * // import { MlKem768 } from "mlkem"; // or "crystals-kyber-js"
     *
     * const kyber = new MlKem768();
     * const [pk, sk] = await kyber.generateKeyPair();
     * ```
     */
    return _createClass(MlKemBase, [{
      key: "generateKeyPair",
      value: (function () {
        var _generateKeyPair = _asyncToGenerator(/*#__PURE__*/_regeneratorRuntime().mark(function _callee() {
          var rnd;
          return _regeneratorRuntime().wrap(function _callee$(_context) {
            while (1) switch (_context.prev = _context.next) {
              case 0:
                _context.next = 2;
                return this._setup();
              case 2:
                _context.prev = 2;
                rnd = new Uint8Array(64);
                this._api.getRandomValues(rnd);
                return _context.abrupt("return", this._deriveKeyPair(rnd));
              case 8:
                _context.prev = 8;
                _context.t0 = _context["catch"](2);
                throw new _MlKemError(_context.t0);
              case 11:
              case "end":
                return _context.stop();
            }
          }, _callee, this, [[2, 8]]);
        }));
        function generateKeyPair() {
          return _generateKeyPair.apply(this, arguments);
        }
        return generateKeyPair;
      }()
      /**
       * Derives a keypair [publicKey, privateKey] deterministically from a 64-octet seed.
       *
       * If an error occurred, throws {@link MlKemError}.
       *
       * @param seed A 64-octet seed for the deterministic key generation.
       * @returns A kaypair [publicKey, privateKey].
       * @throws {@link MlKemError}
       *
       * @example Derives a {@link MlKem768} keypair deterministically.
       *
       * ```ts
       * // Using jsr:
       * import { MlKem768 } from "@dajiaji/mlkem";
       * // Using npm:
       * // import { MlKem768 } from "mlkem"; // or "crystals-kyber-js"
       *
       * const kyber = new MlKem768();
       * const seed = new Uint8Array(64);
       * globalThis.crypto.getRandomValues(seed);
       * const [pk, sk] = await kyber.deriveKeyPair(seed);
       * ```
       */
      )
    }, {
      key: "deriveKeyPair",
      value: (function () {
        var _deriveKeyPair2 = _asyncToGenerator(/*#__PURE__*/_regeneratorRuntime().mark(function _callee2(seed) {
          return _regeneratorRuntime().wrap(function _callee2$(_context2) {
            while (1) switch (_context2.prev = _context2.next) {
              case 0:
                _context2.next = 2;
                return this._setup();
              case 2:
                _context2.prev = 2;
                if (!(seed.byteLength !== 64)) {
                  _context2.next = 5;
                  break;
                }
                throw new Error("seed must be 64 bytes in length");
              case 5:
                return _context2.abrupt("return", this._deriveKeyPair(seed));
              case 8:
                _context2.prev = 8;
                _context2.t0 = _context2["catch"](2);
                throw new _MlKemError(_context2.t0);
              case 11:
              case "end":
                return _context2.stop();
            }
          }, _callee2, this, [[2, 8]]);
        }));
        function deriveKeyPair(_x5) {
          return _deriveKeyPair2.apply(this, arguments);
        }
        return deriveKeyPair;
      }()
      /**
       * Generates a shared secret from the encapsulated ciphertext and the private key.
       *
       * If an error occurred, throws {@link MlKemError}.
       *
       * @param pk A public key.
       * @param seed An optional 32-octet seed for the deterministic shared secret generation.
       * @returns A ciphertext (encapsulated public key) and a shared secret.
       * @throws {@link MlKemError}
       *
       * @example The {@link MlKem768} encapsulation.
       *
       * ```ts
       * // Using jsr:
       * import { MlKem768 } from "@dajiaji/mlkem";
       * // Using npm:
       * // import { MlKem768 } from "mlkem"; // or "crystals-kyber-js"
       *
       * const kyber = new MlKem768();
       * const [pk, sk] = await kyber.generateKeyPair();
       * const [ct, ss] = await kyber.encap(pk);
       * ```
       */
      )
    }, {
      key: "encap",
      value: (function () {
        var _encap2 = _asyncToGenerator(/*#__PURE__*/_regeneratorRuntime().mark(function _callee3(pk, seed) {
          var m, _g, _g2, k, r, ct;
          return _regeneratorRuntime().wrap(function _callee3$(_context3) {
            while (1) switch (_context3.prev = _context3.next) {
              case 0:
                _context3.next = 2;
                return this._setup();
              case 2:
                _context3.prev = 2;
                if (!(pk.length !== 384 * this._k + 32)) {
                  _context3.next = 5;
                  break;
                }
                throw new Error("invalid encapsulation key");
              case 5:
                m = this._getSeed(seed);
                _g = g(m, h(pk)), _g2 = _slicedToArray(_g, 2), k = _g2[0], r = _g2[1];
                ct = this._encap(pk, m, r);
                return _context3.abrupt("return", [ct, k]);
              case 11:
                _context3.prev = 11;
                _context3.t0 = _context3["catch"](2);
                throw new _MlKemError(_context3.t0);
              case 14:
              case "end":
                return _context3.stop();
            }
          }, _callee3, this, [[2, 11]]);
        }));
        function encap(_x6, _x7) {
          return _encap2.apply(this, arguments);
        }
        return encap;
      }()
      /**
       * Generates a ciphertext for the public key and a shared secret.
       *
       * If an error occurred, throws {@link MlKemError}.
       *
       * @param ct A ciphertext generated by {@link encap}.
       * @param sk A private key.
       * @returns A shared secret.
       * @throws {@link MlKemError}
       *
       * @example The {@link MlKem768} decapsulation.
       *
       * ```ts
       * // Using jsr:
       * import { MlKem768 } from "@dajiaji/mlkem";
       * // Using npm:
       * // import { MlKem768 } from "mlkem"; // or "crystals-kyber-js"
       *
       * const kyber = new MlKem768();
       * const [pk, sk] = await kyber.generateKeyPair();
       * const [ct, ssS] = await kyber.encap(pk);
       * const ssR = await kyber.decap(ct, sk);
       * // ssS === ssR
       * ```
       */
      )
    }, {
      key: "decap",
      value: (function () {
        var _decap2 = _asyncToGenerator(/*#__PURE__*/_regeneratorRuntime().mark(function _callee4(ct, sk) {
          var sk2, pk, hpk, z, m2, _g3, _g4, k2, r2, kBar, ct2;
          return _regeneratorRuntime().wrap(function _callee4$(_context4) {
            while (1) switch (_context4.prev = _context4.next) {
              case 0:
                _context4.next = 2;
                return this._setup();
              case 2:
                _context4.prev = 2;
                if (!(ct.byteLength !== this._compressedUSize + this._compressedVSize)) {
                  _context4.next = 5;
                  break;
                }
                throw new Error("Invalid ct size");
              case 5:
                if (!(sk.length !== 768 * this._k + 96)) {
                  _context4.next = 7;
                  break;
                }
                throw new Error("Invalid decapsulation key");
              case 7:
                sk2 = sk.subarray(0, this._skSize);
                pk = sk.subarray(this._skSize, this._skSize + this._pkSize);
                hpk = sk.subarray(this._skSize + this._pkSize, this._skSize + this._pkSize + 32);
                z = sk.subarray(this._skSize + this._pkSize + 32, this._skSize + this._pkSize + 64);
                m2 = this._decap(ct, sk2);
                _g3 = g(m2, hpk), _g4 = _slicedToArray(_g3, 2), k2 = _g4[0], r2 = _g4[1];
                kBar = kdf(z, ct);
                ct2 = this._encap(pk, m2, r2);
                return _context4.abrupt("return", constantTimeCompare(ct, ct2) === 1 ? k2 : kBar);
              case 18:
                _context4.prev = 18;
                _context4.t0 = _context4["catch"](2);
                throw new _MlKemError(_context4.t0);
              case 21:
              case "end":
                return _context4.stop();
            }
          }, _callee4, this, [[2, 18]]);
        }));
        function decap(_x8, _x9) {
          return _decap2.apply(this, arguments);
        }
        return decap;
      }()
      /**
       * Sets up the MlKemBase instance by loading the necessary crypto library.
       * If the crypto library is already loaded, this method does nothing.
       * @returns {Promise<void>} A promise that resolves when the setup is complete.
       */
      )
    }, {
      key: "_setup",
      value: (function () {
        var _setup2 = _asyncToGenerator(/*#__PURE__*/_regeneratorRuntime().mark(function _callee5() {
          return _regeneratorRuntime().wrap(function _callee5$(_context5) {
            while (1) switch (_context5.prev = _context5.next) {
              case 0:
                if (!(this._api !== void 0)) {
                  _context5.next = 2;
                  break;
                }
                return _context5.abrupt("return");
              case 2:
                _context5.next = 4;
                return loadCrypto();
              case 4:
                this._api = _context5.sent;
              case 5:
              case "end":
                return _context5.stop();
            }
          }, _callee5, this);
        }));
        function _setup() {
          return _setup2.apply(this, arguments);
        }
        return _setup;
      }()
      /**
       * Returns a Uint8Array seed for cryptographic operations.
       * If no seed is provided, a random seed of length 32 bytes is generated.
       * If a seed is provided, it must be exactly 32 bytes in length.
       *
       * @param seed - Optional seed for cryptographic operations.
       * @returns A Uint8Array seed.
       * @throws Error if the provided seed is not 32 bytes in length.
       */
      )
    }, {
      key: "_getSeed",
      value: function _getSeed(seed) {
        if (seed == void 0) {
          var s = new Uint8Array(32);
          this._api.getRandomValues(s);
          return s;
        }
        if (seed.byteLength !== 32) {
          throw new Error("seed must be 32 bytes in length");
        }
        return seed;
      }
      /**
       * Derives a key pair from a given seed.
       *
       * @param seed - The seed used for key derivation.
       * @returns An array containing the public key and secret key.
       */
    }, {
      key: "_deriveKeyPair",
      value: function _deriveKeyPair(seed) {
        var cpaSeed = seed.subarray(0, 32);
        var z = seed.subarray(32, 64);
        var _this$_deriveCpaKeyPa = this._deriveCpaKeyPair(cpaSeed),
          _this$_deriveCpaKeyPa2 = _slicedToArray(_this$_deriveCpaKeyPa, 2),
          pk = _this$_deriveCpaKeyPa2[0],
          skBody = _this$_deriveCpaKeyPa2[1];
        var pkh = h(pk);
        var sk = new Uint8Array(this._skSize + this._pkSize + 64);
        sk.set(skBody, 0);
        sk.set(pk, this._skSize);
        sk.set(pkh, this._skSize + this._pkSize);
        sk.set(z, this._skSize + this._pkSize + 32);
        return [pk, sk];
      }
      // indcpaKeyGen generates public and private keys for the CPA-secure
      // public-key encryption scheme underlying ML-KEM.
      /**
       * Derives a CPA key pair using the provided CPA seed.
       *
       * @param cpaSeed - The CPA seed used for key derivation.
       * @returns An array containing the public key and private key.
       */
    }, {
      key: "_deriveCpaKeyPair",
      value: function _deriveCpaKeyPair(cpaSeed) {
        var _g5 = g(cpaSeed, new Uint8Array([this._k])),
          _g6 = _slicedToArray(_g5, 2),
          publicSeed = _g6[0],
          noiseSeed = _g6[1];
        var a = this._sampleMatrix(publicSeed, false);
        var s = this._sampleNoise1(noiseSeed, 0, this._k);
        var e = this._sampleNoise1(noiseSeed, this._k, this._k);
        for (var i = 0; i < this._k; i++) {
          s[i] = ntt(s[i]);
          s[i] = reduce(s[i]);
          e[i] = ntt(e[i]);
        }
        var pk = new Array(this._k);
        for (var _i = 0; _i < this._k; _i++) {
          pk[_i] = polyToMont(multiply(a[_i], s));
          pk[_i] = add(pk[_i], e[_i]);
          pk[_i] = reduce(pk[_i]);
        }
        var pubKey = new Uint8Array(this._pkSize);
        for (var _i2 = 0; _i2 < this._k; _i2++) {
          pubKey.set(polyToBytes(pk[_i2]), _i2 * 384);
        }
        pubKey.set(publicSeed, this._skSize);
        var privKey = new Uint8Array(this._skSize);
        for (var _i3 = 0; _i3 < this._k; _i3++) {
          privKey.set(polyToBytes(s[_i3]), _i3 * 384);
        }
        return [pubKey, privKey];
      }
      // _encap is the encapsulation function of the CPA-secure
      // public-key encryption scheme underlying ML-KEM.
      /**
       * Encapsulates a message using the ML-KEM encryption scheme.
       *
       * @param pk - The public key.
       * @param msg - The message to be encapsulated.
       * @param seed - The seed used for generating random values.
       * @returns The encapsulated message as a Uint8Array.
       */
    }, {
      key: "_encap",
      value: function _encap(pk, msg, seed) {
        var tHat = new Array(this._k);
        var pkCheck = new Uint8Array(384 * this._k);
        for (var i = 0; i < this._k; i++) {
          tHat[i] = polyFromBytes(pk.subarray(i * 384, (i + 1) * 384));
          pkCheck.set(polyToBytes(tHat[i]), i * 384);
        }
        if (!equalUint8Array(pk.subarray(0, pkCheck.length), pkCheck)) {
          throw new Error("invalid encapsulation key");
        }
        var rho = pk.subarray(this._skSize);
        var a = this._sampleMatrix(rho, true);
        var r = this._sampleNoise1(seed, 0, this._k);
        var e1 = this._sampleNoise2(seed, this._k, this._k);
        var e2 = this._sampleNoise2(seed, this._k * 2, 1)[0];
        for (var _i4 = 0; _i4 < this._k; _i4++) {
          r[_i4] = ntt(r[_i4]);
          r[_i4] = reduce(r[_i4]);
        }
        var u = new Array(this._k);
        for (var _i5 = 0; _i5 < this._k; _i5++) {
          u[_i5] = multiply(a[_i5], r);
          u[_i5] = nttInverse(u[_i5]);
          u[_i5] = add(u[_i5], e1[_i5]);
          u[_i5] = reduce(u[_i5]);
        }
        var m = polyFromMsg(msg);
        var v = multiply(tHat, r);
        v = nttInverse(v);
        v = add(v, e2);
        v = add(v, m);
        v = reduce(v);
        var ret = new Uint8Array(this._compressedUSize + this._compressedVSize);
        this._compressU(ret.subarray(0, this._compressedUSize), u);
        this._compressV(ret.subarray(this._compressedUSize), v);
        return ret;
      }
      // indcpaDecrypt is the decryption function of the CPA-secure
      // public-key encryption scheme underlying ML-KEM.
      /**
       * Decapsulates the ciphertext using the provided secret key.
       *
       * @param ct - The ciphertext to be decapsulated.
       * @param sk - The secret key used for decapsulation.
       * @returns The decapsulated message as a Uint8Array.
       */
    }, {
      key: "_decap",
      value: function _decap(ct, sk) {
        var u = this._decompressU(ct.subarray(0, this._compressedUSize));
        var v = this._decompressV(ct.subarray(this._compressedUSize));
        var privateKeyPolyvec = this._polyvecFromBytes(sk);
        for (var i = 0; i < this._k; i++) {
          u[i] = ntt(u[i]);
        }
        var mp = multiply(privateKeyPolyvec, u);
        mp = nttInverse(mp);
        mp = subtract(v, mp);
        mp = reduce(mp);
        return polyToMsg(mp);
      }
      // generateMatrixA deterministically generates a matrix `A` (or the transpose of `A`)
      // from a seed. Entries of the matrix are polynomials that look uniformly random.
      // Performs rejection sampling on the output of an extendable-output function (XOF).
      /**
       * Generates a sample matrix based on the provided seed and transposition flag.
       *
       * @param seed - The seed used for generating the matrix.
       * @param transposed - A flag indicating whether the matrix should be transposed or not.
       * @returns The generated sample matrix.
       */
    }, {
      key: "_sampleMatrix",
      value: function _sampleMatrix(seed, transposed) {
        var a = new Array(this._k);
        var transpose = new Uint8Array(2);
        for (var ctr = 0, i = 0; i < this._k; i++) {
          a[i] = new Array(this._k);
          for (var _j = 0; _j < this._k; _j++) {
            if (transposed) {
              transpose[0] = i;
              transpose[1] = _j;
            } else {
              transpose[0] = _j;
              transpose[1] = i;
            }
            var output = xof(seed, transpose);
            var result = indcpaRejUniform(output.subarray(0, 504), 504, N);
            a[i][_j] = result[0];
            ctr = result[1];
            while (ctr < N) {
              var outputn = output.subarray(504, 672);
              var result1 = indcpaRejUniform(outputn, 168, N - ctr);
              var missing = result1[0];
              var ctrn = result1[1];
              for (var k = ctr; k < N; k++) {
                a[i][_j][k] = missing[k - ctr];
              }
              ctr = ctr + ctrn;
            }
          }
        }
        return a;
      }
      /**
       * Generates a 2D array of noise samples.
       *
       * @param sigma - The noise parameter.
       * @param offset - The offset value.
       * @param size - The size of the array.
       * @returns The generated 2D array of noise samples.
       */
    }, {
      key: "_sampleNoise1",
      value: function _sampleNoise1(sigma, offset, size) {
        var r = new Array(size);
        for (var i = 0; i < size; i++) {
          r[i] = byteopsCbd(prf(this._eta1 * N / 4, sigma, offset), this._eta1);
          offset++;
        }
        return r;
      }
      /**
       * Generates a 2-dimensional array of noise samples.
       *
       * @param sigma - The noise parameter.
       * @param offset - The offset value.
       * @param size - The size of the array.
       * @returns The generated 2-dimensional array of noise samples.
       */
    }, {
      key: "_sampleNoise2",
      value: function _sampleNoise2(sigma, offset, size) {
        var r = new Array(size);
        for (var i = 0; i < size; i++) {
          r[i] = byteopsCbd(prf(this._eta2 * N / 4, sigma, offset), this._eta2);
          offset++;
        }
        return r;
      }
      // polyvecFromBytes deserializes a vector of polynomials.
      /**
       * Converts a Uint8Array to a 2D array of numbers representing a polynomial vector.
       * Each element in the resulting array represents a polynomial.
       * @param a The Uint8Array to convert.
       * @returns The 2D array of numbers representing the polynomial vector.
       */
    }, {
      key: "_polyvecFromBytes",
      value: function _polyvecFromBytes(a) {
        var r = new Array(this._k);
        for (var i = 0; i < this._k; i++) {
          r[i] = polyFromBytes(a.subarray(i * 384, (i + 1) * 384));
        }
        return r;
      }
      // compressU lossily compresses and serializes a vector of polynomials.
      /**
       * Compresses the given array of coefficients into a Uint8Array.
       *
       * @param r - The output Uint8Array.
       * @param u - The array of coefficients.
       * @returns The compressed Uint8Array.
       */
    }, {
      key: "_compressU",
      value: function _compressU(r, u) {
        var t = new Array(4);
        for (var rr = 0, i = 0; i < this._k; i++) {
          for (var _j2 = 0; _j2 < N / 4; _j2++) {
            for (var k = 0; k < 4; k++) {
              t[k] = ((u[i][4 * _j2 + k] << 10) + Q / 2) / Q & 1023;
            }
            r[rr++] = _byte(t[0] >> 0);
            r[rr++] = _byte(t[0] >> 8 | t[1] << 2);
            r[rr++] = _byte(t[1] >> 6 | t[2] << 4);
            r[rr++] = _byte(t[2] >> 4 | t[3] << 6);
            r[rr++] = _byte(t[3] >> 2);
          }
        }
        return r;
      }
      // compressV lossily compresses and subsequently serializes a polynomial.
      /**
       * Compresses the given array of numbers into a Uint8Array.
       *
       * @param r - The Uint8Array to store the compressed values.
       * @param v - The array of numbers to compress.
       * @returns The compressed Uint8Array.
       */
    }, {
      key: "_compressV",
      value: function _compressV(r, v) {
        var t = new Uint8Array(8);
        for (var rr = 0, i = 0; i < N / 8; i++) {
          for (var _j3 = 0; _j3 < 8; _j3++) {
            t[_j3] = _byte(((v[8 * i + _j3] << 4) + Q / 2) / Q) & 15;
          }
          r[rr++] = t[0] | t[1] << 4;
          r[rr++] = t[2] | t[3] << 4;
          r[rr++] = t[4] | t[5] << 4;
          r[rr++] = t[6] | t[7] << 4;
        }
        return r;
      }
      // decompressU de-serializes and decompresses a vector of polynomials and
      // represents the approximate inverse of compress1. Since compression is lossy,
      // the results of decompression will may not match the original vector of polynomials.
      /**
       * Decompresses a Uint8Array into a two-dimensional array of numbers.
       *
       * @param a The Uint8Array to decompress.
       * @returns The decompressed two-dimensional array.
       */
    }, {
      key: "_decompressU",
      value: function _decompressU(a) {
        var r = new Array(this._k);
        for (var i = 0; i < this._k; i++) {
          r[i] = new Array(384);
        }
        var t = new Array(4);
        for (var aa = 0, _i6 = 0; _i6 < this._k; _i6++) {
          for (var _j4 = 0; _j4 < N / 4; _j4++) {
            t[0] = uint16(a[aa + 0]) >> 0 | uint16(a[aa + 1]) << 8;
            t[1] = uint16(a[aa + 1]) >> 2 | uint16(a[aa + 2]) << 6;
            t[2] = uint16(a[aa + 2]) >> 4 | uint16(a[aa + 3]) << 4;
            t[3] = uint16(a[aa + 3]) >> 6 | uint16(a[aa + 4]) << 2;
            aa = aa + 5;
            for (var k = 0; k < 4; k++) {
              r[_i6][4 * _j4 + k] = int16(uint32(t[k] & 1023) * uint32(Q) + 512 >> 10);
            }
          }
        }
        return r;
      }
      // decompressV de-serializes and subsequently decompresses a polynomial,
      // representing the approximate inverse of compress2.
      // Note that compression is lossy, and thus decompression will not match the
      // original input.
      /**
       * Decompresses a Uint8Array into an array of numbers.
       *
       * @param a - The Uint8Array to decompress.
       * @returns An array of numbers.
       */
    }, {
      key: "_decompressV",
      value: function _decompressV(a) {
        var r = new Array(384);
        for (var aa = 0, i = 0; i < N / 2; i++, aa++) {
          r[2 * i + 0] = int16(uint16(a[aa] & 15) * uint16(Q) + 8 >> 4);
          r[2 * i + 1] = int16(uint16(a[aa] >> 4) * uint16(Q) + 8 >> 4);
        }
        return r;
      }
    }]);
  }();
  function g(a, b) {
    var hash = sha3_512.create().update(a);
    if (b !== void 0) {
      hash.update(b);
    }
    var res = hash.digest();
    return [res.subarray(0, 32), res.subarray(32, 64)];
  }
  function h(msg) {
    return sha3_256.create().update(msg).digest();
  }
  function kdf(a, b) {
    var hash = shake256.create({
      dkLen: 32
    }).update(a);
    if (b !== void 0) {
      hash.update(b);
    }
    return hash.digest();
  }
  function xof(seed, transpose) {
    return shake128.create({
      dkLen: 672
    }).update(seed).update(transpose).digest();
  }
  function polyToBytes(a) {
    var t0 = 0;
    var t1 = 0;
    var r = new Uint8Array(384);
    var a2 = subtractQ(a);
    for (var i = 0; i < N / 2; i++) {
      t0 = uint16(a2[2 * i]);
      t1 = uint16(a2[2 * i + 1]);
      r[3 * i + 0] = _byte(t0 >> 0);
      r[3 * i + 1] = _byte(t0 >> 8) | _byte(t1 << 4);
      r[3 * i + 2] = _byte(t1 >> 4);
    }
    return r;
  }
  function polyFromBytes(a) {
    var r = new Array(384).fill(0);
    for (var i = 0; i < N / 2; i++) {
      r[2 * i] = int16((uint16(a[3 * i + 0]) >> 0 | uint16(a[3 * i + 1]) << 8) & 4095);
      r[2 * i + 1] = int16((uint16(a[3 * i + 1]) >> 4 | uint16(a[3 * i + 2]) << 4) & 4095);
    }
    return r;
  }
  function polyToMsg(a) {
    var msg = new Uint8Array(32);
    var t;
    var a2 = subtractQ(a);
    for (var i = 0; i < N / 8; i++) {
      msg[i] = 0;
      for (var _j5 = 0; _j5 < 8; _j5++) {
        t = ((uint16(a2[8 * i + _j5]) << 1) + uint16(Q / 2)) / uint16(Q) & 1;
        msg[i] |= _byte(t << _j5);
      }
    }
    return msg;
  }
  function polyFromMsg(msg) {
    var r = new Array(384).fill(0);
    var mask;
    for (var i = 0; i < N / 8; i++) {
      for (var _j6 = 0; _j6 < 8; _j6++) {
        mask = -1 * int16(msg[i] >> _j6 & 1);
        r[8 * i + _j6] = mask & int16((Q + 1) / 2);
      }
    }
    return r;
  }
  function indcpaRejUniform(buf, bufl, len) {
    var r = new Array(384).fill(0);
    var ctr = 0;
    var val0, val1;
    for (var pos = 0; ctr < len && pos + 3 <= bufl;) {
      val0 = (uint16(buf[pos] >> 0) | uint16(buf[pos + 1]) << 8) & 4095;
      val1 = (uint16(buf[pos + 1] >> 4) | uint16(buf[pos + 2]) << 4) & 4095;
      pos = pos + 3;
      if (val0 < Q) {
        r[ctr] = val0;
        ctr = ctr + 1;
      }
      if (ctr < len && val1 < Q) {
        r[ctr] = val1;
        ctr = ctr + 1;
      }
    }
    return [r, ctr];
  }
  function byteopsCbd(buf, eta) {
    var t, d;
    var a, b;
    var r = new Array(384).fill(0);
    for (var i = 0; i < N / 8; i++) {
      t = byteopsLoad32(buf.subarray(4 * i, buf.length));
      d = t & 1431655765;
      d = d + (t >> 1 & 1431655765);
      for (var _j7 = 0; _j7 < 8; _j7++) {
        a = int16(d >> 4 * _j7 + 0 & 3);
        b = int16(d >> 4 * _j7 + eta & 3);
        r[8 * i + _j7] = a - b;
      }
    }
    return r;
  }
  function ntt(r) {
    for (var _j8 = 0, k = 1, l = 128; l >= 2; l >>= 1) {
      for (var start = 0; start < 256; start = _j8 + l) {
        var zeta = NTT_ZETAS[k];
        k = k + 1;
        for (_j8 = start; _j8 < start + l; _j8++) {
          var _t2 = nttFqMul(zeta, r[_j8 + l]);
          r[_j8 + l] = r[_j8] - _t2;
          r[_j8] = r[_j8] + _t2;
        }
      }
    }
    return r;
  }
  function nttFqMul(a, b) {
    return byteopsMontgomeryReduce(a * b);
  }
  function reduce(r) {
    for (var i = 0; i < N; i++) {
      r[i] = barrett(r[i]);
    }
    return r;
  }
  function barrett(a) {
    var v = ((1 << 24) + Q / 2) / Q;
    var t = v * a >> 24;
    t = t * Q;
    return a - t;
  }
  function byteopsMontgomeryReduce(a) {
    var u = int16(int32(a) * Q_INV);
    var t = u * Q;
    t = a - t;
    t >>= 16;
    return int16(t);
  }
  function polyToMont(r) {
    var f = 1353;
    for (var i = 0; i < N; i++) {
      r[i] = byteopsMontgomeryReduce(int32(r[i]) * int32(f));
    }
    return r;
  }
  function multiply(a, b) {
    var r = polyBaseMulMontgomery(a[0], b[0]);
    var t;
    for (var i = 1; i < a.length; i++) {
      t = polyBaseMulMontgomery(a[i], b[i]);
      r = add(r, t);
    }
    return reduce(r);
  }
  function polyBaseMulMontgomery(a, b) {
    var rx, ry;
    for (var i = 0; i < N / 4; i++) {
      rx = nttBaseMul(a[4 * i + 0], a[4 * i + 1], b[4 * i + 0], b[4 * i + 1], NTT_ZETAS[64 + i]);
      ry = nttBaseMul(a[4 * i + 2], a[4 * i + 3], b[4 * i + 2], b[4 * i + 3], -NTT_ZETAS[64 + i]);
      a[4 * i + 0] = rx[0];
      a[4 * i + 1] = rx[1];
      a[4 * i + 2] = ry[0];
      a[4 * i + 3] = ry[1];
    }
    return a;
  }
  function nttBaseMul(a0, a1, b0, b1, zeta) {
    var r = new Array(2);
    r[0] = nttFqMul(a1, b1);
    r[0] = nttFqMul(r[0], zeta);
    r[0] += nttFqMul(a0, b0);
    r[1] = nttFqMul(a0, b1);
    r[1] += nttFqMul(a1, b0);
    return r;
  }
  function add(a, b) {
    var c = new Array(384);
    for (var i = 0; i < N; i++) {
      c[i] = a[i] + b[i];
    }
    return c;
  }
  function subtract(a, b) {
    for (var i = 0; i < N; i++) {
      a[i] -= b[i];
    }
    return a;
  }
  function nttInverse(r) {
    var j = 0;
    for (var k = 0, l = 2; l <= 128; l <<= 1) {
      for (var start = 0; start < 256; start = j + l) {
        var zeta = NTT_ZETAS_INV[k];
        k = k + 1;
        for (j = start; j < start + l; j++) {
          var _t3 = r[j];
          r[j] = barrett(_t3 + r[j + l]);
          r[j + l] = _t3 - r[j + l];
          r[j + l] = nttFqMul(zeta, r[j + l]);
        }
      }
    }
    for (j = 0; j < 256; j++) {
      r[j] = nttFqMul(r[j], NTT_ZETAS_INV[127]);
    }
    return r;
  }
  function subtractQ(r) {
    for (var i = 0; i < N; i++) {
      r[i] -= Q;
      r[i] += r[i] >> 31 & Q;
    }
    return r;
  }

  // src/mlKem512.ts
  var _MlKem2 = /*#__PURE__*/function (_MlKemBase) {
    /**
     * Constructs a new instance of the MlKem512 class.
     */
    function _MlKem2() {
      var _this3;
      _classCallCheck(this, _MlKem2);
      _this3 = _callSuper(this, _MlKem2);
      _defineProperty(_this3, "_k", 2);
      _defineProperty(_this3, "_du", 10);
      _defineProperty(_this3, "_dv", 4);
      _defineProperty(_this3, "_eta1", 3);
      _defineProperty(_this3, "_eta2", 2);
      _this3._skSize = 12 * _this3._k * N / 8;
      _this3._pkSize = _this3._skSize + 32;
      _this3._compressedUSize = _this3._k * _this3._du * N / 8;
      _this3._compressedVSize = _this3._dv * N / 8;
      return _this3;
    }
    /**
     * Samples a vector of polynomials from a seed.
     * @internal
     * @param sigma - The seed.
     * @param offset - The offset.
     * @param size - The size.
     * @returns The sampled vector of polynomials.
     */
    _inherits(_MlKem2, _MlKemBase);
    return _createClass(_MlKem2, [{
      key: "_sampleNoise1",
      value: function _sampleNoise1(sigma, offset, size) {
        var r = new Array(size);
        for (var i = 0; i < size; i++) {
          r[i] = byteopsCbd2(prf(this._eta1 * N / 4, sigma, offset), this._eta1);
          offset++;
        }
        return r;
      }
    }]);
  }(MlKemBase);
  function byteopsCbd2(buf, eta) {
    var t, d;
    var a, b;
    var r = new Array(384).fill(0);
    for (var i = 0; i < N / 4; i++) {
      t = byteopsLoad24(buf.subarray(3 * i, buf.length));
      d = t & 2396745;
      d = d + (t >> 1 & 2396745);
      d = d + (t >> 2 & 2396745);
      for (var _j9 = 0; _j9 < 4; _j9++) {
        a = int16(d >> 6 * _j9 + 0 & 7);
        b = int16(d >> 6 * _j9 + eta & 7);
        r[4 * i + _j9] = a - b;
      }
    }
    return r;
  }

  // src/mlKem768.ts
  var _MlKem3 = /*#__PURE__*/function (_MlKemBase2) {
    function _MlKem3() {
      var _this4;
      _classCallCheck(this, _MlKem3);
      _this4 = _callSuper(this, _MlKem3);
      _defineProperty(_this4, "_k", 3);
      _defineProperty(_this4, "_du", 10);
      _defineProperty(_this4, "_dv", 4);
      _defineProperty(_this4, "_eta1", 2);
      _defineProperty(_this4, "_eta2", 2);
      _this4._skSize = 12 * _this4._k * N / 8;
      _this4._pkSize = _this4._skSize + 32;
      _this4._compressedUSize = _this4._k * _this4._du * N / 8;
      _this4._compressedVSize = _this4._dv * N / 8;
      return _this4;
    }
    _inherits(_MlKem3, _MlKemBase2);
    return _createClass(_MlKem3);
  }(MlKemBase);

  // src/mlKem1024.ts
  var _MlKem = /*#__PURE__*/function (_MlKemBase3) {
    /**
     * Constructs a new instance of the MlKem1024 class.
     */
    function _MlKem() {
      var _this5;
      _classCallCheck(this, _MlKem);
      _this5 = _callSuper(this, _MlKem);
      _defineProperty(_this5, "_k", 4);
      _defineProperty(_this5, "_du", 11);
      _defineProperty(_this5, "_dv", 5);
      _defineProperty(_this5, "_eta1", 2);
      _defineProperty(_this5, "_eta2", 2);
      _this5._skSize = 12 * _this5._k * N / 8;
      _this5._pkSize = _this5._skSize + 32;
      _this5._compressedUSize = _this5._k * _this5._du * N / 8;
      _this5._compressedVSize = _this5._dv * N / 8;
      return _this5;
    }
    // compressU lossily compresses and serializes a vector of polynomials.
    /**
     * Lossily compresses and serializes a vector of polynomials.
     *
     * @param u - The vector of polynomials to compress.
     * @returns The compressed and serialized data as a Uint8Array.
     */
    _inherits(_MlKem, _MlKemBase3);
    return _createClass(_MlKem, [{
      key: "_compressU",
      value: function _compressU(r, u) {
        var t = new Array(8);
        for (var rr = 0, i = 0; i < this._k; i++) {
          for (var _j10 = 0; _j10 < N / 8; _j10++) {
            for (var k = 0; k < 8; k++) {
              t[k] = uint16(((uint32(u[i][8 * _j10 + k]) << 11) + uint32(Q / 2)) / uint32(Q) & 2047);
            }
            r[rr++] = _byte(t[0] >> 0);
            r[rr++] = _byte(t[0] >> 8 | t[1] << 3);
            r[rr++] = _byte(t[1] >> 5 | t[2] << 6);
            r[rr++] = _byte(t[2] >> 2);
            r[rr++] = _byte(t[2] >> 10 | t[3] << 1);
            r[rr++] = _byte(t[3] >> 7 | t[4] << 4);
            r[rr++] = _byte(t[4] >> 4 | t[5] << 7);
            r[rr++] = _byte(t[5] >> 1);
            r[rr++] = _byte(t[5] >> 9 | t[6] << 2);
            r[rr++] = _byte(t[6] >> 6 | t[7] << 5);
            r[rr++] = _byte(t[7] >> 3);
          }
        }
        return r;
      }
      // compressV lossily compresses and subsequently serializes a polynomial.
      /**
       * Lossily compresses and serializes a polynomial.
       *
       * @param r - The output buffer to store the compressed data.
       * @param v - The polynomial to compress.
       * @returns The compressed and serialized data as a Uint8Array.
       */
    }, {
      key: "_compressV",
      value: function _compressV(r, v) {
        var t = new Uint8Array(8);
        for (var rr = 0, i = 0; i < N / 8; i++) {
          for (var _j11 = 0; _j11 < 8; _j11++) {
            t[_j11] = _byte(((uint32(v[8 * i + _j11]) << 5) + uint32(Q / 2)) / uint32(Q)) & 31;
          }
          r[rr++] = _byte(t[0] >> 0 | t[1] << 5);
          r[rr++] = _byte(t[1] >> 3 | t[2] << 2 | t[3] << 7);
          r[rr++] = _byte(t[3] >> 1 | t[4] << 4);
          r[rr++] = _byte(t[4] >> 4 | t[5] << 1 | t[6] << 6);
          r[rr++] = _byte(t[6] >> 2 | t[7] << 3);
        }
        return r;
      }
      // decompressU de-serializes and decompresses a vector of polynomials and
      // represents the approximate inverse of compress1. Since compression is lossy,
      // the results of decompression will may not match the original vector of polynomials.
      /**
       * Deserializes and decompresses a vector of polynomials.
       * This is the approximate inverse of the `_compressU` method.
       * Since compression is lossy, the decompressed data may not match the original vector of polynomials.
       *
       * @param a - The compressed and serialized data as a Uint8Array.
       * @returns The decompressed vector of polynomials.
       */
    }, {
      key: "_decompressU",
      value: function _decompressU(a) {
        var r = new Array(this._k);
        for (var i = 0; i < this._k; i++) {
          r[i] = new Array(384);
        }
        var t = new Array(8);
        for (var aa = 0, _i7 = 0; _i7 < this._k; _i7++) {
          for (var _j12 = 0; _j12 < N / 8; _j12++) {
            t[0] = uint16(a[aa + 0]) >> 0 | uint16(a[aa + 1]) << 8;
            t[1] = uint16(a[aa + 1]) >> 3 | uint16(a[aa + 2]) << 5;
            t[2] = uint16(a[aa + 2]) >> 6 | uint16(a[aa + 3]) << 2 | uint16(a[aa + 4]) << 10;
            t[3] = uint16(a[aa + 4]) >> 1 | uint16(a[aa + 5]) << 7;
            t[4] = uint16(a[aa + 5]) >> 4 | uint16(a[aa + 6]) << 4;
            t[5] = uint16(a[aa + 6]) >> 7 | uint16(a[aa + 7]) << 1 | uint16(a[aa + 8]) << 9;
            t[6] = uint16(a[aa + 8]) >> 2 | uint16(a[aa + 9]) << 6;
            t[7] = uint16(a[aa + 9]) >> 5 | uint16(a[aa + 10]) << 3;
            aa = aa + 11;
            for (var k = 0; k < 8; k++) {
              r[_i7][8 * _j12 + k] = uint32(t[k] & 2047) * Q + 1024 >> 11;
            }
          }
        }
        return r;
      }
      // decompressV de-serializes and subsequently decompresses a polynomial,
      // representing the approximate inverse of compress2.
      // Note that compression is lossy, and thus decompression will not match the
      // original input.
      /**
       * Decompresses a given polynomial, representing the approximate inverse of
       * compress2, in Uint8Array into an array of numbers.
       *
       * Note that compression is lossy, and thus decompression will not match the
       * original input.
       *
       * @param a - The Uint8Array to decompress.
       * @returns An array of numbers obtained from the decompression process.
       */
    }, {
      key: "_decompressV",
      value: function _decompressV(a) {
        var r = new Array(384);
        var t = new Array(8);
        for (var aa = 0, i = 0; i < N / 8; i++) {
          t[0] = a[aa + 0] >> 0;
          t[1] = a[aa + 0] >> 5 | a[aa + 1] << 3;
          t[2] = a[aa + 1] >> 2;
          t[3] = a[aa + 1] >> 7 | a[aa + 2] << 1;
          t[4] = a[aa + 2] >> 4 | a[aa + 3] << 4;
          t[5] = a[aa + 3] >> 1;
          t[6] = a[aa + 3] >> 6 | a[aa + 4] << 2;
          t[7] = a[aa + 4] >> 3;
          aa = aa + 5;
          for (var _j13 = 0; _j13 < 8; _j13++) {
            r[8 * i + _j13] = int16(uint32(t[_j13] & 31) * uint32(Q) + 16 >> 5);
          }
        }
        return r;
      }
    }]);
  }(MlKemBase);
  return __toCommonJS(mod_exports);
}();
/*! Bundled license information:

@noble/hashes/esm/utils.js:
  (*! noble-hashes - MIT License (c) 2022 Paul Miller (paulmillr.com) *)
*/
