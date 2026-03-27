# Specifications: Report Generator

This directory contains all feature specifications for the Report Generator CLI tool.

## Spec Index

| Spec | Status | Owner | Description |
|------|--------|-------|-------------|
| [core-features.md](core-features.md) | Complete | Alex, Jake, Morgan | All user-facing features: init, run, list commands; configuration; CSV export |

## Reading the Specs

Each spec includes:
- **User Workflow:** Example command and expected output
- **Behavior:** Detailed requirements and error cases
- **Success Criteria:** How we know the feature works
- **Edge Cases:** What we handle and what we explicitly don't handle
- **Open Questions:** Questions for the user, resolved before implementation

## Relationship to Code

Code implements specs. Connection points:

| Spec Section | Code Location | How Connected |
|--------------|---------------|----------------|
| Init command | [src/cmd/init/](../../src/cmd/init/init.go) | Comment at top of file references this spec |
| Run command | [src/cmd/run/](../../src/cmd/run/) | Comments reference this spec |
| Config | [src/config/](../../src/config/) | Loader/validator implements spec behavior |
| CSV export | [src/export/](../../src/export/) | RFC 4180 CSV format per spec decision |

## When Specs Change

If a spec changes after implementation has begun:

1. Document the change in the spec file itself (add a note: "Updated [date]: [what changed]")
2. Add an entry to [../decisions/decision-log.md](../decisions/decision-log.md) explaining why
3. Update code to match new spec
4. Update status in IMPLEMENTATION.md if needed

## Status Definitions

- **Complete:** Implemented, tested, meets all success criteria
- **In Progress:** Active development, not yet ready for use
- **Blocked:** Can't proceed without external input (user decision, design clarity)
- **Planned:** Not yet started, scheduled for future
