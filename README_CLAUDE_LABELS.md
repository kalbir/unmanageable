# Claude Code Label Workflow

Use labels to move an issue through the pipeline.

| Label | Stage | Primary Output |
|---|---|---|
| `status:plan` | Planning | `projects/<slug>/planning/plan.md` |
| `status:design` | Design/Architecture | `projects/<slug>/design/architecture.md` |
| `status:test` | Testing Strategy | `projects/<slug>/testing/test-plan.md` + `regression-checklist.md` |
| `status:implement` | Implementation | Code + tests in `projects/<slug>/src` and `projects/<slug>/tests`, PR raised |
| `status:revise` | Final Revision | Updated docs/code based on feedback + `review/revision-log.md` |

## Revision Trigger

- If issue has `status:revise`, any new user comment triggers revise stage.
- A comment starting with `/revise` also triggers revise stage.

## Single Active Label

When a new workflow status label is added, older workflow labels are removed automatically.
