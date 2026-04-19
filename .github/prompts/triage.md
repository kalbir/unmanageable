# Triage Phase

## Issue Context
Processing issue #${ISSUE_NUMBER}: ${ISSUE_TITLE}

### Issue Description
${ISSUE_BODY}

## Working Directory for This Issue
`${PROJECT_DIR}`

## Your Task
You are in the **Triage Phase** of the action workflow (non-engineering work).

Read the issue carefully and determine what concrete action should be taken next.

Create or update:
1. `${PROJECT_DIR}/triage/action-brief.md`
2. `${PROJECT_DIR}/context/issue.md`
3. `${PROJECT_DIR}/state.json` (if missing fields)

## Action Brief Requirements

The action brief must answer:
- **What type of work is this?** (e.g. document, email, presentation, summary, analysis, proposal, brief, meeting notes)
- **Who is the audience?** (e.g. internal team, external stakeholder, public)
- **What is the desired outcome?** What should the reader/recipient do or understand after receiving the output?
- **What key information from the issue should be included?** Extract and organise the relevant facts, context, and data.
- **What is the recommended output format?** (e.g. markdown document, email draft, slide outline, structured notes)
- **What tone and style is appropriate?** (e.g. formal, conversational, technical, executive summary)

## Action Brief Structure

```markdown
# Action Brief

## Work Type
<type of output to produce>

## Audience
<who this is for>

## Desired Outcome
<what success looks like>

## Key Information
<extracted and organised facts from the issue>

## Recommended Output
- **Format:** <format>
- **Tone:** <tone>
- **Approximate length:** <short/medium/long>

## Open Questions
<anything ambiguous that needs human input before drafting>
```

## Rules
- Be specific and concrete — avoid vague recommendations.
- If the issue contains enough information to act on, say so clearly.
- If critical information is missing, list it under Open Questions.
- Do not produce the draft itself — only the brief.
- Keep this scoped to the current issue only.

## Cost Tracking
At the end of your work, append a run entry to `${PROJECT_DIR}/usage.json`.
If the file does not exist, create it with this structure:
```json
{
  "runs": [
    {
      "stage": "triage",
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
