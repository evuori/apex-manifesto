# Git-Native Tracking

## The Problem With Separate Tracking Tools

Most teams maintain two parallel systems: the code in GitHub and the status in Jira, Linear, or Notion. These systems start synchronized and drift apart within weeks. Status boards become stale. Cards stay "In Progress" long after work has moved on. The board is a lie that everyone politely ignores.

The root cause is that status updates require manual effort disconnected from the actual work. Engineers do the work in git. They have to separately remember to update the board. They often don't. Nobody blames them — it's friction with no direct payoff for the person doing it.

APEX teams don't maintain two systems. The work *is* the tracking.

---

## The Core Principle

**Git activity is the authoritative record of progress.** Commits, pull requests, merges, and branch states tell the story of what's happening without any manual reporting layer. The goal of git-native tracking is to make that story legible — to teams, to stakeholders, and to AI tools — without adding a separate tool.

This connects directly to the Repo-as-OS principle (see [guides/repo-as-os.md](repo-as-os.md)): if everything lives in the repo, tracking falls out of the repo's natural state rather than requiring its own maintenance.

---

## The Pull Request as the Unit of Work

In git-native tracking, the pull request is the kanban card. It has everything a work item needs:

- **Title** — what it is
- **Description** — why it exists, what spec it implements, what done looks like
- **Status** — open, in review, merged, closed
- **Assignee** — who owns it (APEX: one owner, not "the team")
- **Labels** — workstream, type, priority
- **Timeline** — when it was opened, when it moved, when it merged
- **Discussion** — review comments, blockers surfaced, decisions made

The PR is not a code artifact with a card attached. The PR *is* the card. The code is what's inside it.

---

## Branch Naming Convention

Branch names encode context that makes the board readable without a separate tool:

```
{team}/{outcome}/{short-description}

Examples:
platform/auth/add-sso-support
data/reporting/fix-export-timeout
mobile/onboarding/redesign-step-two
```

**`{team}`** — which team owns this work. In multi-team environments, this makes cross-team visibility immediate from the branch list.

**`{outcome}`** — which outcome this work contributes to. This is the equivalent of the epic or the swimlane. Teams should agree on outcome slugs at the start of a cycle and keep them consistent.

**`{short-description}`** — what the branch does, in a few words. Lowercase, hyphenated.

This convention means `git branch -a` tells you who is working on what and why, without opening a project board.

---

## Labels as Swimlanes

GitHub labels replace the kanban column configuration. Standardize a label set across repos so any team member can read any repo's board without relearning the system.

### Recommended Label Set

**Status**
- `status: in-progress` — actively being worked on
- `status: in-review` — PR open, waiting for review
- `status: blocked` — can't move forward, requires resolution
- `status: ready-to-merge` — approved, waiting for merge

**Type**
- `type: feature` — new capability
- `type: fix` — corrects a defect
- `type: spec` — design or planning artifact (non-code PR)
- `type: decision` — records a significant decision
- `type: infra` — infrastructure, tooling, not user-facing

**Outcome** (set per project/cycle)
- `outcome: auth`
- `outcome: reporting`
- `outcome: onboarding`
- *(define these at cycle start, keep them consistent with branch names)*

Labels are applied when a PR is opened. Status labels are updated as work moves. This is the only manual update required — and it's a single click, not a separate tool.

---

## WIP Limits

Work In Progress limits are the mechanism that makes kanban actually work rather than just look like a wall of cards. Without WIP limits, teams pile up "In Progress" work, context-switch constantly, and finish nothing.

In git-native tracking, WIP limits are enforced by convention:

**Per team:** No more than N branches with `status: in-progress` open at once. When a team hits the limit, new work cannot start until something finishes or moves to review. The standard starting point is WIP = team size. Adjust based on observed cycle time.

**Per person:** One owner per PR. One active PR per person is the ideal. Two is acceptable during handoffs. Three is a signal something is wrong.

The discipline: when someone opens a new branch and their team is already at the WIP limit, they pick up an existing in-progress item and move it forward instead of starting something new.

This feels uncomfortable. That discomfort is the point. It surfaces bottlenecks — usually in review, usually because ownership of review isn't clear. When review becomes the constraint, you fix the review process, not the WIP limit.

---

## Milestones as Shared Outcomes

In multi-team environments, GitHub Milestones serve as the shared coordination layer across teams. A milestone represents a meaningful outcome that multiple teams contribute to — not a sprint, not a date, but a state of the world that matters.

Each PR that contributes to a milestone is attached to it. This gives leadership and cross-team stakeholders a single view: how many PRs are merged toward this milestone, how many are in progress, how many are blocked.

**Milestone naming:** Outcomes, not dates.
- ✅ `auth-v1-complete`
- ✅ `reporting-pipeline-live`
- ❌ `sprint-14`
- ❌ `q2-work`

Dates can be added to milestones in GitHub as due dates. But the name should describe the outcome, not the deadline.

---

## Automated Status From Git Activity

The most powerful extension of git-native tracking is automating the status report. No engineer should be writing standup updates or weekly summaries. That information exists in git.

A GitHub Action running nightly (or on demand) can:

1. Collect all PRs merged in the last 24h across team repos
2. Collect all PRs currently open with `status: blocked`
3. Collect all PRs open > N days with no activity (stall signals)
4. Pass this to an AI model with a prompt to generate a plain-language summary

The output is a brief posted to a shared channel, a file committed to `docs/status/YYYY-MM-DD.md`, or both. The team wakes up to a status summary they didn't have to write.

This is not aspirational — it is achievable with GitHub Actions and any LLM API today. The APEX tooling layer will provide a reference implementation (see the repo's `/examples` directory as it evolves).

### What the Summary Contains

```
## Status — 2026-03-28

**Shipped yesterday**
- platform/auth: SSO support merged (PR #47)
- data/reporting: Export timeout fix merged (PR #51)

**In progress** (5 open PRs)
- mobile/onboarding: Redesign step two — in review, no blockers
- platform/auth: Refresh token rotation — in progress, opened 3 days ago
- data/reporting: CSV schema validation — in progress

**Blocked** (1)
- mobile/onboarding: Push notification spec (#53) — waiting on design decision
  → Decision needed: which OS permission flow to use

**Stalled** (1)
- platform/infra: Database migration (#44) — no activity in 5 days
  → Needs attention
```

No meeting required to produce this. No manual updates. Git generated it.

---

## Multi-Team Coordination

When multiple teams work toward a shared goal, coordination happens through three mechanisms — none of which require a separate tool:

**Shared milestones** — cross-team outcomes that any team's PRs can attach to. Gives a unified view of progress toward shared goals.

**Cross-repo specs** — when a feature requires coordination across teams, a spec lives in the shared `docs/specs/` directory at the org or project level. Both teams reference it from their PRs. Changes to the spec go through a PR reviewed by both teams.

**Blocking PR links** — when one team's work is blocked by another team's PR, the GitHub "linked issues/PRs" feature makes the dependency explicit. The automated status report surfaces these as blockers with the responsible team named.

What this eliminates: the coordination meeting that exists only to share status. Teams look at the milestone, the blocked PRs, and the automated summary. The meeting can be about actual decisions, not updates.

---

## What Gets Tracked That Isn't Code

Non-code work — design, research, planning — is tracked through the same mechanism: PRs. A designer submitting a design brief as a PR gets a PR number, an owner, a label, and a milestone attachment just like a code change. It appears in the milestone progress. It appears in the automated summary.

This is only possible if non-code artifacts live in the repo. See [guides/repo-as-os.md](repo-as-os.md) for why and how.

---

## Making the Switch

If your team currently uses a separate tracking tool, the transition is incremental:

1. **Start with branch naming.** This costs nothing and immediately adds visibility.
2. **Add the standard label set.** One setup task, then it's just clicks.
3. **Attach PRs to milestones.** Define one meaningful outcome per cycle and attach everything to it.
4. **Drop the manual standup update.** Let git activity speak for itself. Review open PRs in a brief daily sync instead of status monologues.
5. **Add automated summaries.** Once the label discipline is in place, the automation is straightforward.

The tracking tool can be sunset when the team stops looking at it — usually within a cycle.

---

## What This Does Not Replace

Git-native tracking handles delivery well. It does not handle:

- **Strategic roadmapping** — what to build over the next two quarters. That lives in `docs/planning/` in the repo, not on a kanban board.
- **Stakeholder dashboards** — external-facing progress views for non-technical stakeholders who aren't reading PRs. These can be generated from git data, but require a presentation layer.
- **Qualitative team health** — velocity and throughput metrics don't tell you if the team is burning out or misaligned. Use the [health-check template](../templates/health-check.md) for that.

---

## APEX Principles in Practice

- **Accountable:** Every PR has one owner. Every blocked item names who can unblock it. The automated summary makes accountability visible without a manager having to ask.
- **Precise:** Branch names and PR descriptions encode outcome and intent. "What are you working on and why?" is answerable from the branch list.
- **Expert-led:** Experts decide what goes into the repo, how it's structured, and what the WIP limits should be. The system serves them — not the other way around.
- **eXpedient:** No separate tool to maintain. No manual status updates. The overhead of tracking approaches zero.

---

## Next Steps

1. Agree on outcome slugs for your current cycle — these become your branch prefixes and milestone names
2. Set up the standard label set in your GitHub repos
3. Define WIP limits per team and post them in your `README.md` or `IMPLEMENTATION.md`
4. Open your next piece of work as a PR with a proper description referencing the relevant spec
5. Review [guides/repo-as-os.md](repo-as-os.md) to understand how non-code artifacts fit into the same tracking flow

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
