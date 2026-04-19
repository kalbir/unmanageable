# CLAUDE.md

This repository runs two issue-label workflows: one for engineering projects and one for non-engineering work.

## Engineering Workflow Stages

1. `status:plan`
2. `status:design`
3. `status:test`
4. `status:implement`
5. `status:revise`

## Action Workflow Stages

1. `action:triage`
2. `action:draft`

## Core Rules

- Engineering issue artifacts must live under `projects/<issue-title-kebab-case>/`.
- Action issue artifacts must live under `work/<issue-title-kebab-case>/`.
- Do not create an `issues/` artifact directory.
- Keep phase outputs deterministic and explicit.
- Preserve project-local context in `state.json` and `comments/`.
- In `implement`, produce code and tests in project-local `src/` and `tests/`.
- In `draft`, produce output artifacts in work-local `drafts/`.
- Track cost/usage in `usage.json` for every workflow run.

## Stage Intent

### Plan Stage

Goal: turn issue text into an explicit implementation plan and scope boundaries.

### Design Stage

Goal: produce a deterministic architecture/design spec at file and interface level.

### Test Stage

Goal: define/add tests to close coverage gaps and lock in expected behavior.

### Implement Stage

Goal: execute design + testing strategy, produce working code, and prepare PR-ready changes.

### Revise Stage

Goal: apply reviewer feedback and tighten deliverables before merge.

### Triage Stage

Goal: analyse the issue, determine what type of work it is, and produce a concrete action brief describing the recommended output.

### Draft Stage

Goal: execute the action brief — produce the output artifact(s) (document, email, presentation, etc.) ready for human review.

## Quality Bar

- State assumptions directly in artifacts.
- Keep edits minimal and in-scope.
- Prefer reproducibility over open-ended outputs.
- Do not leave unresolved ambiguity in stage outputs.
