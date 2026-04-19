# Code Project Workflow Monorepo

A GitHub issue-driven workflow system for building generic code projects in a monorepo.

Each issue can be progressed through fixed stages by labels.

### Engineering Projects (`projects/`)

1. `status:plan`
2. `status:design`
3. `status:test`
4. `status:implement`
5. `status:revise`

### Non-Engineering Work (`work/`)

1. `action:triage`
2. `action:draft`

Engineering projects live under `projects/<issue-title-kebab-case>/`.
Work items live under `work/<issue-title-kebab-case>/`.

## How It Works

### Engineering Projects

1. Create an issue describing the project or feature.
2. Add `status:plan` to generate a concrete implementation plan.
3. Add `status:design` to generate architecture/design specs.
4. Add `status:test` to generate a test strategy and gaps checklist.
5. Add `status:implement` to implement code in that project directory and open/update a PR.
6. Add `status:revise` (or comment `/revise`) to apply feedback and refine outputs.

### Non-Engineering Work

1. Create an issue describing the work (e.g. draft an email, write a document, prepare a brief).
2. Add `action:triage` — Claude analyses the issue and produces an action brief with recommended output format.
3. Add `action:draft` — Claude produces the artifact(s) and opens a PR.
4. Comment on the issue to iterate — Claude revises the drafts based on your feedback.

## Project Directory Layout

### Engineering Projects

```text
projects/<project-slug>/
  context/
    issue.md
  planning/
    plan.md
  design/
    architecture.md
  testing/
    test-plan.md
    regression-checklist.md
  implementation/
    implementation-notes.md
  review/
    revision-log.md
  comments/
    <timestamp>.md
  src/
  tests/
  state.json
```

### Non-Engineering Work

```text
work/<work-slug>/
  context/
    issue.md
  triage/
    action-brief.md
  drafts/
    <descriptive-name>.md
    README.md
  comments/
    <timestamp>.md
  usage.json
  state.json
```

## Pipeline Behavior

- The workflow listens to issue label events and issue comments.
- Only one workflow status label is kept at a time.
- All automation output for an issue is constrained to `projects/<slug>/`.
- `status:implement` and `action:draft` create/update a branch `codex/issue-<number>-<slug>` and open/update a PR against your default branch.
- When `action:draft` is active, any user comment on the issue re-triggers the draft stage with the feedback incorporated.

## Answering Plan Questions

If `projects/<slug>/planning/plan.md` includes open questions:

1. Add a `## Decisions / Answers` section to the GitHub issue body.
2. List each answer as numbered items that map to the plan questions.
3. Re-run planning by removing and re-adding `status:plan`.

This is the most reliable way to get your answers incorporated, since the plan stage reads from the current issue body.

## Cost Tracking

Each workflow run appends a record to `usage.json` in the project/work directory with the model used, stage, and timestamp. Token counts and cost breakdowns are recorded when available.

## Key Files

- Workflow: `.github/workflows/claude-code-projects.yml`
- Prompt templates: `.github/prompts/*.md` (`plan`, `design`, `test`, `implement`, `revise`, `triage`, `draft`)
- Prompt processor: `.github/scripts/process-prompt.sh`
- Workflow guide: `CLAUDE.md`
- Engineering principles: `.claude/principles.md`
- Label guide: `README_CLAUDE_LABELS.md`
