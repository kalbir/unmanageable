# Work Directory

Non-engineering work items processed through the `action:*` label workflow.

Each issue is mapped to one work folder:

- `work/<issue-title-kebab-case>/`

## Directory Layout

```text
work/<slug>/
  context/issue.md        # Original issue text
  triage/action-brief.md  # What to do and how (from triage stage)
  drafts/                 # Output artifacts (documents, emails, etc.)
  comments/               # Timestamped feedback from issue comments
  usage.json              # Cost and token tracking per run
  state.json              # Current stage and metadata
```

## Workflow

1. Create an issue describing the work.
2. Add `action:triage` — Claude analyses the issue and produces an action brief.
3. Add `action:draft` — Claude produces the output artifact(s) and opens a PR.
4. Comment on the issue to iterate — Claude revises the drafts based on your feedback.
