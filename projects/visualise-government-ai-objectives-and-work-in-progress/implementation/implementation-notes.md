# Implementation Notes

## Summary

Implemented a Next.js 14 web application that visualises UK government AI objectives and work in progress, with Chart.js charts, department/category/status/date filters, and a structured JSON data store.

## Stack

| Concern | Choice | Reason |
|---------|--------|--------|
| Framework | Next.js 14 (App Router) | Static export (`output: "export"`) for Vercel, server components by default |
| Charts | Chart.js + react-chartjs-2 | Stakeholder preference |
| Data | JSON (`src/data/initiatives.json`) | MVP-appropriate; can migrate to graph DB later |
| Language | TypeScript | Type-safe schema enforces data integrity |
| Deployment | Vercel (static export) | Stakeholder requirement |
| Tests | Jest + ts-jest | Fast, isolated unit tests; no browser needed for logic |

## File Structure

```
src/
  app/
    layout.tsx        — HTML shell, metadata
    page.tsx          — Entry point; loads JSON data
    globals.css       — All styles (single file for MVP)
  components/
    Dashboard.tsx     — State management (filters, view toggle), layout
    FilterBar.tsx     — Filter controls (department, category, status, date, search)
    InitiativeCard.tsx — Single initiative display
    charts/
      StatusDonut.tsx   — Doughnut chart: breakdown by status
      DepartmentBar.tsx — Horizontal bar: initiatives per department
      CategoryBar.tsx   — Horizontal bar: initiatives per category
      TimelineBar.tsx   — Stacked bar: announcements per year by status
  lib/
    filterInitiatives.ts — Pure filter/aggregation functions (no React deps)
    chartData.ts         — Chart.js dataset builders
  types/
    initiative.ts     — TypeScript interfaces for Initiative, FilterState, etc.
  data/
    initiatives.json  — 15 seed initiatives from gov.uk, GDS, i.AI sources

tests/
  filterInitiatives.test.ts — 14 tests covering all filter logic and edge cases
  chartData.test.ts         — 8 tests covering chart dataset builders
  initiativesData.test.ts   — 6 tests validating JSON data integrity
```

## Key Decisions

### Data schema flexibility
The `Initiative` type includes `tags[]`, `sources[]`, and a `category` enum. This allows future graph visualisations (initiatives as nodes, departments/categories as edges). The JSON is versioned with a `version` field for schema migration tracking.

### Filter architecture
All filtering is pure TypeScript in `lib/filterInitiatives.ts` — no React state logic. This makes the filters fast, testable, and reusable if the data layer changes.

### Static export
`next.config.js` sets `output: "export"`. This means the app builds to a `/out` directory of static HTML/JS/CSS with no server runtime, suitable for Vercel's free tier with zero cold starts.

### Chart data builders are separate from components
`lib/chartData.ts` is framework-agnostic. Chart components just call these builders. This keeps components thin and makes chart logic independently testable.

## Assumptions

1. **No backend required for MVP** — data is bundled at build time from `initiatives.json`.
2. **Data updates via PR** — the weekly pipeline (see below) will update the JSON file and trigger a Vercel redeploy.
3. **No authentication** — the site is public read-only.
4. **Chart.js 4.x** — requires manual registration of chart components (done in each chart component).
5. **15 seed initiatives** — representative sample covering DSIT, DHSC, DfE, Home Office, MoD, HMRC, DWP, GDS, i.AI.

## Categories Proposed

Based on the UK government AI landscape, ten primary categories:

| Category | Example initiatives |
|----------|-------------------|
| Healthcare | NHS AI Lab, medical imaging AI |
| Education | DfE generative AI framework, AI skills |
| Defence & Security | MoD AI strategy, facial recognition |
| Public Services | HMRC tax compliance, DWP benefits |
| Economic Growth | AI Opportunities Action Plan |
| Infrastructure | AI compute strategy, growth zones |
| Research & Innovation | National AI Strategy research pillar |
| Regulation & Safety | AI Safety Institute, pro-innovation regulation |
| Digital Government | GDS AI Studio, i.AI incubator |
| Other | Catch-all for cross-cutting or unclassified |

These were included in the `Category` type in `types/initiative.ts`.

## Deployment Instructions (Vercel)

### First deploy

1. Push this branch to GitHub (or merge to `main`).
2. Go to [vercel.com](https://vercel.com) → New Project → Import the repository.
3. Set **Root Directory** to `projects/visualise-government-ai-objectives-and-work-in-progress/src`.
4. Framework preset: **Next.js** (auto-detected).
5. Build command: `npm run build` (default).
6. Output directory: `out` (set manually, or Next.js tells Vercel automatically).
7. Click Deploy.

Vercel will assign a URL such as `uk-gov-ai-tracker.vercel.app`.

### Subsequent deploys

Every push to `main` will trigger a Vercel rebuild automatically.

### Custom domain

In Vercel project settings → Domains, add your custom domain and follow the DNS instructions.

## Weekly Data Pipeline

A GitHub Actions workflow can automate weekly data refresh:

```yaml
# .github/workflows/refresh-data.yml  (add when ready)
name: Weekly data refresh
on:
  schedule:
    - cron: "0 8 * * 1"   # Every Monday at 08:00 UTC
  workflow_dispatch:

jobs:
  refresh:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run data collection script
        run: node scripts/collect-data.js
      - name: Commit updated data
        run: |
          git config user.email "actions@github.com"
          git config user.name "GitHub Actions"
          git add projects/visualise-government-ai-objectives-and-work-in-progress/src/data/initiatives.json
          git commit -m "chore: weekly data refresh" || echo "No changes"
          git push
```

A `scripts/collect-data.js` script would fetch from gov.uk RSS feeds, department blogs, and known i.AI/GDS sources, then merge new items into `initiatives.json`. This is a follow-on task.

## Running Tests

```bash
cd projects/visualise-government-ai-objectives-and-work-in-progress/tests
npm install
npx jest
```

Tests do not require a running Next.js app — they test pure TypeScript logic only.

## Running the App Locally

```bash
cd projects/visualise-government-ai-objectives-and-work-in-progress/src
npm install
npm run dev
# Open http://localhost:3000
```
