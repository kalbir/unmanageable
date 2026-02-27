# Global Engineering Principles

Grounded in Extreme Programming. Favor simplicity, feedback, and disciplined iteration.

## Design
- Do the simplest thing that could possibly work.
- YAGNI: do not build speculative features.
- Delay abstractions until repeated concrete cases justify them.
- Prefer composition over inheritance.
- Choose clear, intention-revealing names.

## Working with Code
- Make one coherent change at a time.
- Refactor continuously, but stay in scope.
- If fixing a bug, add a failing test first when practical.
- Prefer updating existing files over creating new ones unless structure requires it.

## Testing
- Tests must be fast, isolated, and deterministic.
- Test behavior, not implementation details.
- Cover edge cases and error paths.
- Every test requires meaningful assertions.

## Communication
- Optimize for the next reader.
- Comments explain why, not what.
- If requirements are ambiguous, make assumptions explicit in project artifacts.

## Discipline
- Run relevant tests before reporting completion.
- Do not suppress failures to make pipelines pass.
- Separate behavior changes from pure refactors.
