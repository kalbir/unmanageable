# 11+ Sprint

A single-file 11+ practice app built for a capable-but-reluctant test-taker preparing
for **Dame Alice Owen's**, **Henrietta Barnett** and similar London super-selectives.
Designed around **short bursts** and a **10-day plan of ~20-minute sessions**.

The look is a soft **kawaii-London** theme aimed at a 10-year-old girl: candy pastels,
rounded bubbly type, sparkle rewards on correct answers, and **Biscuit the corgi**
(with a tiny crown) as an encouraging coach, over a faint pastel London skyline
(Big Ben, the London Eye, the Shard, Tower Bridge). Full light + dark themes.

## Use it

- **Open `index.html`** in any browser (desktop or mobile) — it's fully self-contained
  and works offline after first load. Progress is saved in the browser's local storage.
- It is also published as a private, shareable Artifact (a link that opens on a phone).

## How it's tuned to the brief

- **Short rounds of 5.** The atomic unit is a 5-question round. A "day" is ~20 minutes
  (4 rounds of 20 questions), but you can stop after any round — progress is saved.
- **Momentum, not pressure.** Instant right/wrong feedback with a one-line explanation,
  a day streak, XP, and encouraging copy. Timing is visible but gentle; an optional
  **Timed exam mode** can be switched on later for exam stamina.
- **Good at individual questions.** Every question stands alone and is multiple-choice
  (tap-friendly on mobile). Difficulty adapts to recent accuracy (3 levels).
- **10-day arc.** Each day has a focus but interleaves topics (better for retention):
  warm-up → maths → verbal → vocab → problems → analogies → mini-mock → non-verbal →
  **weak-spot focus** (targets your lowest-accuracy skill) → confidence day.

## Coverage

| Skill | How it's generated | Examples |
|-------|--------------------|----------|
| **Maths** | Procedural (endless) | arithmetic, BODMAS, fractions/%, ratio, sequences, mean, area/perimeter, time, word problems |
| **Verbal reasoning** | Procedural + word bank | letter series, interleaved sequences, letter codes, synonyms/antonyms, analogies, number analogies, odd-one-out |
| **English** | Curated banks | short comprehension passages, vocabulary, spelling, grammar/punctuation |
| **Non-verbal** | Procedural (light) | shape-count sequences, rotation sequences |

Maths, verbal and non-verbal questions are generated with random parameters, so the
bank never runs out over the 10 days; English draws from curated passages and word banks.

## Tech notes

- One HTML file. No build step, no dependencies, no network calls except Google Fonts
  (Fredoka / Baloo 2 / Quicksand, with rounded system-font fallbacks if offline).
- The corgi mascot and London skyline are inline SVG (no image files).
- Light and dark themes (follows the device, with a manual toggle).
- `sprint.artifact.html` is a content-only copy generated from `index.html` for
  publishing as a Claude Artifact; `index.html` is the canonical standalone source.

## Set the exam date

Open **Progress → Set your exam date** to show a live days-to-exam countdown on the
home screen. (For these schools the 11+ typically falls in September of Year 6.)
