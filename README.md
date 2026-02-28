# Code Project Workflow Monorepo

A GitHub issue-driven workflow system for building generic code projects in a monorepo.

Each issue can be progressed through fixed stages by labels:

1. `status:plan`
2. `status:design`
3. `status:test`
4. `status:implement`
5. `status:revise`

Each project lives in its own directory under `projects/<issue-title-kebab-case>/`.

## How It Works

1. Create an issue describing the project or feature.
2. Add `status:plan` to generate a concrete implementation plan.
3. Add `status:design` to generate architecture/design specs.
4. Add `status:test` to generate a test strategy and gaps checklist.
5. Add `status:implement` to implement code in that project directory and open/update a PR.
6. Add `status:revise` (or comment `/revise`) to apply feedback and refine outputs.

## Project Directory Layout

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

## Pipeline Behavior

- The workflow listens to issue label events and issue comments.
- Only one workflow status label is kept at a time.
- All automation output for an issue is constrained to `projects/<slug>/`.
- `status:implement` creates/updates a branch `codex/issue-<number>-<slug>` and opens/updates a PR against your default branch.

## Answering Plan Questions

If `projects/<slug>/planning/plan.md` includes open questions:

1. Add a `## Decisions / Answers` section to the GitHub issue body.
2. List each answer as numbered items that map to the plan questions.
3. Re-run planning by removing and re-adding `status:plan`.

This is the most reliable way to get your answers incorporated, since the plan stage reads from the current issue body.

## Key Files

- Workflow: `.github/workflows/claude-code-projects.yml`
- Prompt templates: `.github/prompts/*.md`
- Prompt processor: `.github/scripts/process-prompt.sh`
- Workflow guide: `CLAUDE.md`
- Engineering principles: `.claude/principles.md`
- Label guide: `README_CLAUDE_LABELS.md`
