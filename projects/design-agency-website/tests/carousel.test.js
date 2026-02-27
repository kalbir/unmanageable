/**
 * Unit tests for the values carousel logic.
 * Run with: node tests/carousel.test.js
 * No external dependencies required.
 */

var assert = require('assert');

/* ------------------------------------------------------------------ */
/* Minimal DOM stub                                                     */
/* ------------------------------------------------------------------ */

function makeEl(tag, attrs) {
  var el = {
    _tag: tag,
    _children: [],
    _attrs: Object.assign({}, attrs || {}),
    _classes: [],
    _listeners: {},
    style: {},
    disabled: false,
    textContent: '',
    getAttribute: function (k) { return this._attrs[k] !== undefined ? String(this._attrs[k]) : null; },
    setAttribute: function (k, v) { this._attrs[k] = v; },
    classList: {
      _el: null,
      add: function () { Array.prototype.forEach.call(arguments, function (c) { if (this._el._classes.indexOf(c) === -1) this._el._classes.push(c); }.bind(this)); },
      remove: function () { Array.prototype.forEach.call(arguments, function (c) { var i = this._el._classes.indexOf(c); if (i > -1) this._el._classes.splice(i, 1); }.bind(this)); },
      toggle: function (c, force) {
        var has = this._el._classes.indexOf(c) > -1;
        if (force === true) { if (!has) this._el._classes.push(c); }
        else if (force === false) { var i = this._el._classes.indexOf(c); if (i > -1) this._el._classes.splice(i, 1); }
        else { if (has) { var j = this._el._classes.indexOf(c); this._el._classes.splice(j, 1); } else { this._el._classes.push(c); } }
      },
      contains: function (c) { return this._el._classes.indexOf(c) > -1; }
    },
    addEventListener: function (ev, fn) {
      if (!this._listeners[ev]) this._listeners[ev] = [];
      this._listeners[ev].push(fn);
    },
    _trigger: function (ev, detail) {
      (this._listeners[ev] || []).forEach(function (fn) { fn(detail || {}); });
    }
  };
  el.classList._el = el;
  return el;
}

/* ------------------------------------------------------------------ */
/* Carousel factory (mirrors main.js logic)                            */
/* ------------------------------------------------------------------ */

function createCarousel(count) {
  var slides = [];
  for (var i = 0; i < count; i++) {
    var s = makeEl('div', { 'aria-hidden': i === 0 ? 'false' : 'true' });
    slides.push(s);
  }

  var prevBtn = makeEl('button');
  var nextBtn = makeEl('button');
  var currentLabel = makeEl('span');
  currentLabel.textContent = '01';

  var dots = slides.map(function (_, i) {
    return makeEl('button', { 'data-target': i, 'aria-current': i === 0 ? 'true' : 'false' });
  });

  var current = 0;
  var total = count;
  var isAnimating = false;

  function pad(n) { return String(n + 1).padStart(2, '0'); }

  function showSlide(index, direction) {
    if (isAnimating || index === current) return;
    isAnimating = true;

    var outgoing = slides[current];
    var incoming = slides[index];

    outgoing.classList.add('is-exiting');
    outgoing.setAttribute('aria-hidden', 'true');
    incoming.classList.add('is-active');
    incoming.setAttribute('aria-hidden', 'false');

    setTimeout(function () {
      outgoing.classList.remove('is-active', 'is-exiting');
      isAnimating = false;
    }, 0);

    current = index;
    currentLabel.textContent = pad(current);

    dots.forEach(function (dot, i) {
      dot.classList.toggle('values__dot--active', i === current);
      dot.setAttribute('aria-current', i === current ? 'true' : 'false');
    });

    prevBtn.disabled = current === 0;
    nextBtn.disabled = current === total - 1;
  }

  slides[0].classList.add('is-active');
  prevBtn.disabled = true;

  prevBtn.addEventListener('click', function () { if (current > 0) showSlide(current - 1, -1); });
  nextBtn.addEventListener('click', function () { if (current < total - 1) showSlide(current + 1, 1); });

  dots.forEach(function (dot) {
    dot.addEventListener('click', function () {
      var target = parseInt(dot.getAttribute('data-target'), 10);
      if (!isNaN(target)) showSlide(target, target > current ? 1 : -1);
    });
  });

  return { slides: slides, prevBtn: prevBtn, nextBtn: nextBtn, currentLabel: currentLabel, dots: dots, getCurrent: function () { return current; } };
}

/* ------------------------------------------------------------------ */
/* Tests                                                                */
/* ------------------------------------------------------------------ */

var passed = 0;
var failed = 0;

function test(name, fn) {
  try {
    fn();
    console.log('  \u2713 ' + name);
    passed++;
  } catch (e) {
    console.error('  \u2717 ' + name);
    console.error('    ' + e.message);
    failed++;
  }
}

console.log('\nValues Carousel Tests\n');

test('first slide starts active', function () {
  var c = createCarousel(5);
  assert.ok(c.slides[0].classList.contains('is-active'), 'slide 0 should be active');
});

test('prev button starts disabled', function () {
  var c = createCarousel(5);
  assert.strictEqual(c.prevBtn.disabled, true);
});

test('next button starts enabled when count > 1', function () {
  var c = createCarousel(5);
  assert.strictEqual(c.nextBtn.disabled, false);
});

test('clicking next advances to slide 1', function () {
  var c = createCarousel(5);
  c.nextBtn._trigger('click');
  assert.strictEqual(c.getCurrent(), 1);
});

test('current label updates to 02 after next click', function () {
  var c = createCarousel(5);
  c.nextBtn._trigger('click');
  assert.strictEqual(c.currentLabel.textContent, '02');
});

test('dot aria-current updates on slide change', function () {
  var c = createCarousel(5);
  c.nextBtn._trigger('click');
  assert.strictEqual(c.dots[0].getAttribute('aria-current'), 'false');
  assert.strictEqual(c.dots[1].getAttribute('aria-current'), 'true');
});

test('prev button enabled after moving to slide 1', function () {
  var c = createCarousel(5);
  c.nextBtn._trigger('click');
  assert.strictEqual(c.prevBtn.disabled, false);
});

test('next button disabled on last slide', function () {
  var c = createCarousel(3);
  c.nextBtn._trigger('click');
  c.nextBtn._trigger('click');
  assert.strictEqual(c.getCurrent(), 2);
  assert.strictEqual(c.nextBtn.disabled, true);
});

test('clicking prev from slide 1 returns to slide 0', function () {
  var c = createCarousel(5);
  c.nextBtn._trigger('click');
  c.prevBtn._trigger('click');
  assert.strictEqual(c.getCurrent(), 0);
});

test('clicking dot navigates directly to target', function () {
  var c = createCarousel(5);
  c.dots[3]._trigger('click');
  assert.strictEqual(c.getCurrent(), 3);
});

test('outgoing slide gets aria-hidden set to true', function () {
  var c = createCarousel(5);
  c.nextBtn._trigger('click');
  assert.strictEqual(c.slides[0].getAttribute('aria-hidden'), 'true');
});

test('incoming slide gets aria-hidden set to false', function () {
  var c = createCarousel(5);
  c.nextBtn._trigger('click');
  assert.strictEqual(c.slides[1].getAttribute('aria-hidden'), 'false');
});

test('active dot class updates correctly', function () {
  var c = createCarousel(5);
  c.dots[2]._trigger('click');
  assert.ok(c.dots[2].classList.contains('values__dot--active'));
  assert.ok(!c.dots[0].classList.contains('values__dot--active'));
});

test('single-slide carousel: both buttons disabled', function () {
  var c = createCarousel(1);
  assert.strictEqual(c.prevBtn.disabled, true);
  assert.strictEqual(c.nextBtn.disabled, true);
});

console.log('\n' + passed + ' passed, ' + failed + ' failed\n');

if (failed > 0) process.exit(1);
