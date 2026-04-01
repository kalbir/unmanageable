# Implementation Plan: UK Government AI Visualisation

## Problem Statement

Build a web page that visualises:
1. **Objectives** — publicly stated AI commitments (announcements, parliamentary statements).
2. **Work in Progress** — publicly stated ongoing AI work (announcements, blogs, news).

Data must be stored flexibly enough to support future visualisations and graph-based exploration.

## Scope

### In Scope (MVP)
- Static web page with Chart.js visualisations
- Two-category data model: objectives / work-in-progress
- JSON data files as the single source of truth
- Filters: department, timeframe, plus 1–2 additional fields (defined in design phase)
- Deployment to Vercel with instructions for the project owner
- Weekly automated data refresh via GitHub Actions pipeline

### Out of Scope (post-MVP)
- Graph database backend
- Authentication or admin UI
- Real-time data fetching from live sources
- Full-text search

## Data Sources

| Source | Type | Notes |
|--------|------|-------|
| gov.uk | Official announcements, press releases | Primary source |
| Department blogs (e.g. DSIT, DHSC, MOD) | Blog posts, news | gov.uk-hosted and standalone |
| GDS AI Studio | Blog posts, GitHub | Not on main gov.uk |
| i.AI (Incubator for Artificial Intelligence) | Blog posts, publications | Not on main gov.uk |

## Technology Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Data format | JSON | Simple for MVP; schema can mirror a node/edge graph structure for future migration |
| Visualisation library | Chart.js | Higher-level than D3.js; sufficient for MVP bar, timeline, and donut charts |
| Deployment | Vercel | Project owner has account; automatic redeploy on git push |
| Update pipeline | GitHub Actions (scheduled) | Commits refreshed JSON weekly; triggers Vercel redeploy |

## Implementation Tasks

1. **Data schema design** — Define JSON schema for initiatives, including fields for category, department, date, source URL, status, and tags.
2. **Seed data** — Populate initial JSON with 20–30 curated UK government AI items across both categories.
3. **Category suggestions** — Propose theme taxonomy (e.g. healthcare, transport, justice, public services, defence) in design phase.
4. **Web page scaffold** — HTML/CSS/JS static site; load JSON, render with Chart.js.
5. **Filters** — Implement department filter, timeframe filter, and up to 2 more (e.g. category/theme, initiative type).
6. **Deployment** — Configure Vercel project, write deployment instructions for project owner.
7. **Update pipeline** — GitHub Actions workflow on weekly schedule to update data JSON and commit.
8. **Tests** — Data schema validation tests; basic rendering smoke tests.

## Risks and Dependencies

| Risk | Mitigation |
|------|-----------|
| Government data is scattered across many sources | Start with a curated seed set; pipeline handles ongoing discovery |
| Category taxonomy may not fit all initiatives | Treat taxonomy as v1; allow `tags` array for overflow |
| Vercel deployment requires owner action | Provide step-by-step instructions; pipeline handles subsequent deploys |
| Chart.js may not support graph/network views | JSON schema designed for graph migration; graph view deferred to post-MVP |

## Open Questions (Resolved)

All 7 planning open questions resolved — see `review/revision-log.md`.
