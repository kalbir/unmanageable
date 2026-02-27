(function () {
  'use strict';

  /* === Nav scroll effect === */
  function initNavScroll() {
    var nav = document.querySelector('.nav');
    if (!nav) return;

    function onScroll() {
      nav.classList.toggle('is-scrolled', window.scrollY > 10);
    }

    window.addEventListener('scroll', onScroll, { passive: true });
    onScroll();
  }

  /* === Values carousel === */
  function initValuesCarousel() {
    var track = document.getElementById('valuesTrack');
    var prevBtn = document.getElementById('valuePrev');
    var nextBtn = document.getElementById('valueNext');
    var currentLabel = document.getElementById('valueCurrentLabel');
    var dots = document.querySelectorAll('.values__dot');
    var slides = track ? Array.from(track.querySelectorAll('.value-slide')) : [];

    if (!track || slides.length === 0) return;

    var current = 0;
    var total = slides.length;
    var isAnimating = false;

    function pad(n) {
      return String(n + 1).padStart(2, '0');
    }

    function showSlide(index, direction) {
      if (isAnimating || index === current) return;
      isAnimating = true;

      var outgoing = slides[current];
      var incoming = slides[index];

      outgoing.classList.add('is-exiting');
      outgoing.setAttribute('aria-hidden', 'true');

      incoming.style.transform = direction > 0 ? 'translateX(80px)' : 'translateX(-80px)';
      incoming.classList.add('is-active');
      incoming.setAttribute('aria-hidden', 'false');

      var duration = 350;

      setTimeout(function () {
        outgoing.classList.remove('is-active', 'is-exiting');
        incoming.style.transform = '';
        isAnimating = false;
      }, duration);

      current = index;

      if (currentLabel) currentLabel.textContent = pad(current);

      dots.forEach(function (dot, i) {
        dot.classList.toggle('values__dot--active', i === current);
        dot.setAttribute('aria-current', i === current ? 'true' : 'false');
      });

      if (prevBtn) prevBtn.disabled = current === 0;
      if (nextBtn) nextBtn.disabled = current === total - 1;
    }

    slides[0].classList.add('is-active');
    if (prevBtn) prevBtn.disabled = true;

    if (prevBtn) {
      prevBtn.addEventListener('click', function () {
        if (current > 0) showSlide(current - 1, -1);
      });
    }

    if (nextBtn) {
      nextBtn.addEventListener('click', function () {
        if (current < total - 1) showSlide(current + 1, 1);
      });
    }

    dots.forEach(function (dot) {
      dot.addEventListener('click', function () {
        var target = parseInt(dot.getAttribute('data-target'), 10);
        if (!isNaN(target)) showSlide(target, target > current ? 1 : -1);
      });
    });

    track.addEventListener('keydown', function (e) {
      if (e.key === 'ArrowRight' || e.key === 'ArrowDown') {
        e.preventDefault();
        if (current < total - 1) showSlide(current + 1, 1);
      } else if (e.key === 'ArrowLeft' || e.key === 'ArrowUp') {
        e.preventDefault();
        if (current > 0) showSlide(current - 1, -1);
      }
    });

    var touchStartX = 0;

    track.addEventListener('touchstart', function (e) {
      touchStartX = e.touches[0].clientX;
    }, { passive: true });

    track.addEventListener('touchend', function (e) {
      var dx = e.changedTouches[0].clientX - touchStartX;
      if (Math.abs(dx) > 50) {
        if (dx < 0 && current < total - 1) showSlide(current + 1, 1);
        if (dx > 0 && current > 0) showSlide(current - 1, -1);
      }
    }, { passive: true });
  }

  /* === Init === */
  document.addEventListener('DOMContentLoaded', function () {
    initNavScroll();
    initValuesCarousel();
  });
}());
