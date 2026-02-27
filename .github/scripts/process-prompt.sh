#!/bin/bash
set -euo pipefail

TEMPLATE_FILE="$1"
OUTPUT_FILE="$2"

if [ ! -f "$TEMPLATE_FILE" ]; then
  echo "Template file not found: $TEMPLATE_FILE"
  exit 1
fi

TEMP_FILE=$(mktemp)
envsubst < "$TEMPLATE_FILE" > "$TEMP_FILE"

{
  echo
  echo "## Repository Context"
  echo

  if [ -f "CLAUDE.md" ]; then
    echo "### Workflow Guide (CLAUDE.md)"
    echo '```'
    cat "CLAUDE.md"
    echo '```'
    echo
  fi

  if [ -f ".claude/principles.md" ]; then
    echo "### Engineering Principles (.claude/principles.md)"
    echo '```'
    cat ".claude/principles.md"
    echo '```'
    echo
  fi
} >> "$TEMP_FILE"

mv "$TEMP_FILE" "$OUTPUT_FILE"
echo "Processed prompt written to: $OUTPUT_FILE"
