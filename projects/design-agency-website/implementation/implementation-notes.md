# Implementation Notes — Design Agency Website

## Overview

Static HTML/CSS/JS demo for a fictional design studio called **FORMA**. No build tooling, no backend, no external JS libraries. The three source files (`src/index.html`, `src/main.css`, `src/main.js`) can be opened directly in a browser.

## Design Decisions

### Colour Scheme
- Background: `#0A0A0A` (near-black) — maximises contrast and creates depth.
- Accent: `#FF006E` (electric pink) — bold, singular, used sparingly for critical UI and active states.
- Text: `#F0EDE8` (warm off-white) — softer than pure white, reduces eye strain.
- Surfaces and borders are layered using low-opacity whites so the palette stays coherent.

### Typography
- **Display / headers**: Inter Tight (variable, 100–900 weight) — high contrast between thin and black weights enables expressive headline composition.
- **Body / UI**: Inter (300, 400) — neutral and legible at small sizes.
- All type sizes use `clamp()` fluid values so headlines scale continuously from mobile to wide-screen.

### Layout and Space
- No imagery; all visual weight comes from type scale, colour contrast, and generous whitespace.
- The hero uses near-full-viewport height with content pinned to the bottom, creating a sense of scale before the user scrolls.
- A ghost-number behind each values slide (`-webkit-text-stroke`) provides depth through typographic texture without image assets.
- The work grid uses a 1px border-collapse trick (background on the grid, gap of 1px, tiles have solid bg) to draw a grid without hard-coded border logic.
- The contact section inverts the colour scheme (accent pink as background) — a full-bleed colour block acting as a punctuation mark before the footer.

### Values Carousel
- Slides are absolutely positioned within a fixed-height container. Transitions use `translateX` + `opacity` for a slide-in/slide-out effect.
- Prev/next buttons, dot indicators, and keyboard arrow navigation all work.
- Touch swipe support (threshold: 50 px) for mobile.
- `aria-hidden`, `aria-current`, `aria-live`, and `aria-label` attributes kept in sync on every transition.
- `prefers-reduced-motion` disables CSS animations and shortens JS transition durations.

### Marquee
- Pure CSS `animation: marquee` on a doubled-up track — no JS scroll listener.
- Paused via `animation: none` under `prefers-reduced-motion`.

## Assumptions

- No design/architecture file was present in `projects/design-agency-website/design/`; this implementation was built directly from the issue description.
- Google Fonts (Inter Tight, Inter) are loaded from CDN. An offline fallback to system sans-serif is specified.
- No IE11 support required; CSS custom properties and `clamp()` are used freely.
- Accessibility target: WCAG 2.1 AA for keyboard navigation and colour contrast.

## Testing

`tests/carousel.test.js` — 14 unit tests covering:
- Initial state (active slide, button states)
- Forward and backward navigation
- Direct dot navigation
- `aria-hidden` / `aria-current` attribute sync
- Edge cases: single-slide carousel, last-slide boundary

Run with:
```
node tests/carousel.test.js
```

No npm install required.

## Out of Scope

- Backend, CMS, or form submission.
- Internationalisation.
- Analytics.
- Dark/light mode toggle (page is dark-mode only by design).
