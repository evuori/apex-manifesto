# Design Brief Template

---

## Purpose

This template captures the *thinking* behind a design decision — not the pixels, but the reasoning. It lives in the repository alongside the design assets it describes, making it reviewable via PR, readable by AI tools, and accessible to any team member without opening a separate design tool.

A design brief is not a spec (which describes *what* to build). It is not a Figma file (which shows *how* it looks). It is the bridge: *why* it looks and works the way it does, what constraints existed, what alternatives were considered, and what success looks like from the user's perspective.

**Where it lives:**
```
design/
└── {feature-name}/
    ├── rationale.md          ← this document
    ├── assets/               ← exported images, SVGs, tokens
    └── iterations/           ← prior versions that explain the evolution (optional)
```

---

## When to Write One

Write a design brief when:
- A new user-facing feature or flow is being designed
- An existing design is being significantly changed
- A design decision was debated and one approach was chosen over another
- The design is constrained by something non-obvious (technical, business, accessibility)

You do not need one for minor visual fixes or copy changes. The test: *would a new team member or an AI assistant need this to understand why the design works the way it does?*

---

## The Template

Copy and fill in what's relevant. Not every section applies to every brief. Delete sections that don't add value — a brief that's short and complete is better than one that's long and padded.

---

```markdown
# Design Brief: {Feature or Flow Name}

**Owner:** {One person. Not "the team."}
**Date:** YYYY-MM-DD
**Status:** Draft | In Review | Approved
**Related spec:** {link to docs/specs/feature-name.md if one exists}
**Related PR:** {link when this brief is submitted as a PR}

---

## What We're Designing

One paragraph. What is this feature or flow? Who uses it? What are they trying to do?

Example:
> The onboarding flow for new users completing their first project setup. Users arrive
> after signup and need to connect a repository, configure their settings, and run their
> first check — ideally without reading documentation.

---

## The User's Goal

What does the user actually want to accomplish? Not what the feature does — what the user
is trying to do, in their terms.

Example:
> Get to a working state as fast as possible. They signed up because they want the
> outcome, not because they want to configure software. Every step that isn't directly
> toward the outcome is friction.

---

## Constraints

What made this design hard? What limits existed that shaped the decisions?

- **Technical:** {e.g., "Can't access file system until permissions are granted — flow must handle the case where user declines"}
- **Business:** {e.g., "Must collect billing info before trial ends — can't skip this step"}
- **Accessibility:** {e.g., "Must work with screen readers; modal approach creates ARIA complexity"}
- **Time:** {e.g., "Shipping in 2 weeks — deferring advanced configuration to v2"}
- **Existing system:** {e.g., "Auth flow is owned by a third-party library — limited control over that screen"}

---

## Design Decisions

The specific choices made and why. This is the most important section.

For each significant decision:

**Decision:** {What was chosen}
**Why:** {The reasoning — what problem it solves, what constraint it satisfies}
**Alternatives considered:** {What else was tried or evaluated}

Example:

**Decision:** Single-page setup instead of a multi-step wizard
**Why:** User research showed wizard steps created false expectations of progress — users
abandoned when they hit a step that required something they didn't have ready. A single
page shows everything upfront so users can assess the full task before starting.
**Alternatives considered:** 3-step wizard (abandoned after usability test showed 40%
drop-off at step 2), progressive disclosure accordion (too much interaction for a one-time
flow).

---

## What Success Looks Like

How do we know this design works? Specific, observable outcomes.

Example:
> - A new user with a GitHub account can complete setup in under 3 minutes
> - Users don't need to leave the flow to find information (no "where do I find X?" questions in support)
> - The error state for declined permissions is understood without support intervention

---

## What's Deferred

What was explicitly decided to leave out of this version, and why.

Example:
> - Team invite flow: out of scope for v1. Solo setup is the primary path.
>   Will design separately once v1 ships and we see how teams are actually using it.
> - Dark mode: deferred until design system is updated. Current tokens don't support it cleanly.

---

## Open Questions

Things that are not yet decided and need resolution before or during implementation.

Example:
> - Do we show a progress indicator? Leaning no (it's a single page) but want to validate
>   with one more test before committing.
> - What happens if the user closes the browser mid-setup and returns? Does the state persist?
>   Engineering input needed on feasibility.

---

## Assets

Links or paths to design assets associated with this brief.

| Asset | Location | Notes |
|-------|----------|-------|
| Screens (PNG) | `assets/screens/` | Exported from Figma 2026-03-28 |
| Flow diagram | `assets/flow.svg` | Full user journey |
| Design tokens | `assets/tokens.json` | Colors and spacing used |

---

## Revision Notes

If this brief has been updated after initial approval, note what changed and why here.
This is not a full changelog — only changes that would affect implementation or that
overturned a previous decision.

| Date | Change | Reason |
|------|--------|--------|
| | | |
```

---

## How to Use This in Practice

**When opening the PR:** The design brief is the PR description's source of truth. Link to this file. Reviewers read the brief before looking at the assets — it tells them what to evaluate.

**When reviewing:** Evaluate the brief, not just the visuals. Does the reasoning hold? Are the constraints accurate? Are there alternatives not considered? The review comment goes in the PR discussion, but significant changes to the brief's conclusions go back into the file.

**When the brief changes:** If a design decision gets overturned during implementation, update this file and log the change in the revision notes. Don't leave the brief pointing at a decision that's no longer true — the next person to read it (including an AI) will be misled.

**When AI is building from this:** An AI implementing a design should be given this brief alongside the spec. The brief explains intent; the spec explains behavior. Together they give the AI enough context to make reasonable decisions without constant interruption.

---

## The Difference Between a Brief and a Spec

A spec (see [guides/spec-driven-development.md](../guides/spec-driven-development.md)) answers: *What does this feature do? What are the success criteria? What are the edge cases?*

A design brief answers: *Why does it look and work this way? What constraints shaped the choices? What was rejected and why?*

They reference each other but serve different readers. An engineer implementing from the spec should also read the brief. A designer updating the design should also read the spec.

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
