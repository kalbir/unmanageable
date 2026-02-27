# CLAUDE.md

This repository runs a generic code-project workflow using issue labels.

## Workflow Stages

1. `status:plan`
2. `status:design`
3. `status:test`
4. `status:implement`
5. `status:revise`

## Core Rules

- All issue artifacts must live under `projects/<issue-title-kebab-case>/`.
- Do not create an `issues/` artifact directory.
- Keep phase outputs deterministic and explicit.
- Preserve project-local context in `state.json` and `comments/`.
- In `implement`, produce code and tests in project-local `src/` and `tests/`.

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

## Quality Bar

- State assumptions directly in artifacts.
- Keep edits minimal and in-scope.
- Prefer reproducibility over open-ended outputs.
- Do not leave unresolved ambiguity in stage outputs.
