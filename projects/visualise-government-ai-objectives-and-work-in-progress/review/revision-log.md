# Revision Log

## Revision: 2026-04-01

### Source
Feedback from @kalbir in issue #4 comment dated 2026-04-01, answering 7 open questions from the planning phase.

### Feedback → Change Mapping

| # | Open Question | Answer | Change |
|---|---------------|--------|--------|
| 1 | Data format preference | JSON for MVP | plan.md: data storage specified as JSON files; graph structure deferred post-MVP |
| 2 | Visualisation library | Chart.js preferred | plan.md: Chart.js confirmed as primary visualisation library; D3.js removed from consideration |
| 3 | Deployment platform | Vercel; user has account | plan.md: deployment target is Vercel; deployment instructions to be included in implementation |
| 4 | Update frequency | Weekly | plan.md: weekly data update cadence; CI/CD pipeline to automate data refresh |
| 5 | Data sources | gov.uk, department blogs, GDS AI Studio, i.AI | plan.md: explicit source list added; i.AI and GDS AI Studio noted as non-gov.uk sources |
| 6 | Categories/themes | Suggest categories in design, iterate | plan.md: category design delegated to design phase with example suggestions |
| 7 | Interactivity level | Filter by department, timeframe, ~2 more | plan.md: filters scoped to department, timeframe, and 1–2 additional fields TBD in design |

### Assumptions

- "Some sort of pipeline" for weekly updates is interpreted as a GitHub Actions workflow that commits refreshed JSON data on a schedule; Vercel redeploys automatically on push.
- i.AI refers to the Prime Minister's AI unit (https://www.gov.uk/government/organisations/incubator-for-artificial-intelligence).
- GDS AI Studio refers to the Government Digital Service AI team and its public blog/GitHub output.
- Category suggestions will be proposed in the design phase and are not finalised here.
- "A couple more" filters beyond department and timeframe is treated as 1–2; exact fields deferred to design.

### Tradeoffs

- JSON chosen over graph DB for MVP simplicity; the schema will be designed to allow future migration to a graph format (e.g. nodes/edges structure within JSON).
- Chart.js is higher-level than D3.js, which speeds up development but limits bespoke visual forms. Acceptable for MVP.
- Static site on Vercel keeps infrastructure simple; no backend required for MVP data serving.

### No Changes

- Issue scope (two categories: objectives vs work-in-progress) unchanged.
- Web page delivery target unchanged.
- Flexible data backing requirement unchanged.
