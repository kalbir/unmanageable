#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

required_files=(
  "$ROOT/.github/workflows/claude-code-projects.yml"
  "$ROOT/.github/scripts/process-prompt.sh"
  "$ROOT/.github/prompts/plan.md"
  "$ROOT/.github/prompts/design.md"
  "$ROOT/.github/prompts/test.md"
  "$ROOT/.github/prompts/implement.md"
  "$ROOT/.github/prompts/revise.md"
  "$ROOT/CLAUDE.md"
  "$ROOT/.claude/principles.md"
  "$ROOT/README.md"
)

for file in "${required_files[@]}"; do
  if [ ! -f "$file" ]; then
    echo "Missing required file: $file"
    exit 1
  fi
done

bash -n "$ROOT/.github/scripts/process-prompt.sh"

workflow="$ROOT/.github/workflows/claude-code-projects.yml"
for label in status:plan status:design status:test status:implement status:revise; do
  if ! rg -q "$label" "$workflow"; then
    echo "Workflow is missing label: $label"
    exit 1
  fi
done

echo "Workflow contract check passed."
