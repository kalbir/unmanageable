# Draft Phase

## Issue Context
Processing issue #${ISSUE_NUMBER}: ${ISSUE_TITLE}

### Issue Description
${ISSUE_BODY}

## Working Directory for This Issue
`${PROJECT_DIR}`

## Your Task
You are in the **Draft Phase** of the action workflow (non-engineering work).

Read the action brief at `${PROJECT_DIR}/triage/action-brief.md` and produce the output artifact(s) it describes.

Create or update:
1. Artifact(s) in `${PROJECT_DIR}/drafts/` — use the format specified in the action brief
2. `${PROJECT_DIR}/drafts/README.md` — index of what was produced and why

## Drafting Rules
- Follow the action brief closely: match the specified format, tone, audience, and length.
- If the brief has unresolved Open Questions, make reasonable assumptions and state them clearly at the top of the draft.
- Produce complete, usable output — not outlines or placeholders (unless the brief specifically asks for an outline).
- Use markdown files by default. For presentations, produce a structured markdown outline with clear slide breaks (`---`).
- For email drafts, include To/Subject/Body sections.
- Name files descriptively (e.g. `proposal.md`, `stakeholder-email.md`, `q3-summary.md`), not generically.
- Do not modify files outside `${PROJECT_DIR}`.

## Revision Context
If `${PROJECT_DIR}/comments/` contains feedback files, incorporate that feedback into the draft.
If previous drafts exist in `${PROJECT_DIR}/drafts/`, update them rather than creating duplicates.

### Latest Comment
${LATEST_COMMENT_BODY}

## Quality Bar
- The draft should be ready for a human to review, lightly edit, and send/publish.
- Prefer clarity and brevity over comprehensiveness.
- Match the audience's expected level of detail.

## Cost Tracking
At the end of your work, append a run entry to `${PROJECT_DIR}/usage.json`.
If the file does not exist, create it with this structure:
```json
{
  "runs": [
    {
      "stage": "draft",
      "model": "${MODEL}",
      "timestamp": "<current UTC ISO 8601>",
      "input_tokens": null,
      "output_tokens": null,
      "total_cost_usd": null
    }
  ]
}
```
If the file exists, append to the `runs` array.
