# APEX vs. Agile/Scrum: A Practical Comparison

---

## Why This Document

Most teams adopting APEX are coming from Agile or Scrum. You have built habits around sprints, user stories, velocity, and standup ceremonies. When you switch to APEX, specific questions will emerge:

- "What replaces the two-week sprint?"
- "How do we do planning if not sprint planning?"
- "What do we do instead of standups?"
- "How do we estimate if not story points?"

This document answers those questions directly. The comparison is not a polemic — APEX and Agile both solve real problems. The distinction is what those problems are, and what constraints you face.

---

## The Fundamental Difference

| Agile | APEX |
|-------|------|
| Optimized for **coordination at scale** — ensuring a large team works in alignment | Optimized for **individual ownership** — ensuring experts move fast without coordination overhead |
| Assumes teams change roles and need standardized processes for consistency | Assumes teams are stable and expertise matters |
| Treats velocity and consistency as valuable metrics | Treats delivery as the only meaningful metric |
| Makes distributed decision-making safe through process | Makes centralized (owner-level) decisions safe through clarity |

If you have a 40-person engineering org with high turnover, Agile is probably more suitable. If you have a 5-10 person expert team, APEX is probably more suitable.

---

## Ceremony-by-Ceremony Comparison

### Sprint vs. Delivery Cycle

| Agile/Scrum | APEX |
|----------|------|
| **Two-week sprints.** Every feature, fix, and task is assigned to a sprint. Velocity is tracked. Commitment to sprint scope is a core metric. | **No enforced sprint.** Work ships when it is ready. Planning is continuous and tied to delivery priorities, not calendar events. |
| Planning happens at sprint boundaries. Scope is locked for 2 weeks. Changes require process. | Planning happens as work changes. Scope is fluid until the decision log makes it fixed. |
| Why: Consistency, predictability, and team alignment. Everyone knows what the team is doing this sprint. | Why: Expedience. Forcing two-week boundaries is overhead when delivery is independent. |

**If you are adopting APEX, here is what you actually do:**

- Stop doing sprint planning meetings
- Stop assigning work to sprint numbers
- Align planning with when things actually need to ship (which might be weekly, might be monthly)
- Owners commit to delivery dates, not to sprint completion
- If you have stakeholders who need "alignment visibility," provide a weekly or bi-weekly async update (see below) instead of a ceremony

---

### User Stories with Acceptance Criteria vs. Work Items

| Agile/Scrum | APEX |
|----------|------|
| **User story format:** "As a [role], I want [feature] so that [benefit]." Includes acceptance criteria (testable conditions that define done). Points estimated. | **Work item format:** Problem (what is broken or missing), scope (what is in/out), done condition (how you know it shipped). Owner assigned. No estimation. |
| Why: Stories capture user perspective. Acceptance criteria prevent scope creep. Estimation enables velocity tracking. | Why: A problem statement beats a role statement. A clear owner beats distributed accountability. Done condition beats criteria lists. Estimation is overhead. |

**If you are adopting APEX:**

- Replace your story template with the [work-item-template.md](work-item-template.md)
- Do not estimate story points
- Make owner assignment non-negotiable (do not create work without an owner)
- Replace "acceptance criteria" with one "done condition" statement

---

### Daily Standup vs. Async Status

| Agile/Scrum | APEX |
|----------|------|
| **Daily standup.** Team gathers. Each person says what they did, what they are doing, and what is blocking them. Scrum Master watches. | **Async status (optional).** One sentence per owner, posted to a shared channel or document. 10 seconds to read. No meeting. |
| Why: Daily synchronization. Early surface of blockers. Team cohesion. | Why: Synchronous standups do not actually surface blockers faster than async. They feel productive. They rarely are. |

**If you are adopting APEX:**

- If standups are working, keep them but make them fast (5 minutes, one person speaking at a time)
- If they feel like busywork, replace them with an async one-liner per owner posted end-of-day or start-of-day
- Alternatively: eliminate standups entirely and replace with async updates only if blocking surfaces

---

### Sprint Retrospective vs. Lightweight Review

| Agile/Scrum | APEX |
|----------|------|
| **Sprint retrospective.** Team meets to discuss what went well, what went poorly, and what to improve. Action items are captured. Next sprint retrospective will check on them. | **Lightweight review (optional).** 30-minute check-in at delivery boundaries. What shipped? What needs to change? Run by the owner of the work. One-paragraph summary recorded. Done. |
| Why: Continuous improvement. Safe space to surface problems. Process visibility. | Why: Retrospectives work when teams actively fix things. They atrophy when they do not. A lightweight review forces action (the owner is accountable). A long retro feels productive and accomplishes less. |

**If you are adopting APEX:**

- Replace lengthy retrospectives with a 30-minute check-in run by the project owner
- Ask three questions: What shipped? What do we need to adjust? Are there ownership or prioritization questions?
- Record a one-paragraph summary
- The owner is accountable for follow-through (not "we should improve communication" — "I am going to do X to improve communication by Y date")

---

### Sprint Backlog vs. Priority List

| Agile/Scrum | APEX |
|----------|------|
| **Backlog.** Prioritized list of all known work. Grows constantly. Backlog refinement is a ceremony (making sure items are ready for future sprints). Velocity determines what fits in a sprint. | **Priority list.** Top 5-10 items that matter now. Not a catalog of everything that has ever been requested. The product owner (or whoever owns prioritization) actively says no to everything else. |
| Why: Transparency. Stakeholders can see what is queued. | Why: Backlogs become junk drawers. A 200-item backlog is not a plan — it is a procrastination list. A ruthless top-5 is a real decision. |

**If you are adopting APEX:**

- Delete your 100+ item backlog
- Create a top-5 list of what actually matters
- Everything else is "not in scope for the next 6 months" (you can revisit it, but it is not in the queue)
- When new requests come in, the owner of prioritization (usually a product person) decides: top-5 or no
- "No" is a complete answer

---

### Velocity Tracking vs. Delivery Metrics

| Agile/Scrum | APEX |
|----------|------|
| **Velocity.** Measure story points completed per sprint. Optimize for consistency. | **Delivery.** Measure things shipped. Did the deadline happen? Is the feature live? That is all that matters. |
| Why: Predictability. Enables future planning. Shows productivity. | Why: Velocity measures activity, not progress. You can have high velocity and low delivery. Delivery is what matters. |

**If you are adopting APEX:**

- Stop tracking velocity
- Stop estimating story points
- The only metric that matters is: what shipped this week/month?
- If the owner wants to project timelines, have them estimate based on prior similar work, not on velocity curves

---

### Scrum Master Role vs. Owner Accountability

| Agile/Scrum | APEX |
|----------|------|
| **Scrum Master.** Faciliates ceremonies. Removes impediments. Watches process health. Reports on team metrics. | **Owner.** Makes decisions. Communicates. Calls out problems. Ships work. That is the role. |
| Why: A dedicated process facilitator can scale Agile to larger teams. | Why: A separate process role diffuses accountability. The person doing the work should own removing blockers. |

**If you are adopting APEX:**

- You do not have Scrum Masters under APEX. The owner does the work and owns the process
- If you have a "technical lead," their role is to unblock owners and help them make decisions, not to manage process
- If you need someone tracking "team health," that person is diagnosing a team problem. Fix it directly instead of instrumenting it.

---

### Product Owner Role

| Agile/Scrum | APEX |
|----------|------|
| Clearly defined role. Owns the backlog, prioritizes work, writes stories, represents the user. | Undefined role name. Whoever owns the product/domain makes decisions about what to build. That person owns prioritization and user understanding. |
| Why: Clear accountability for product decisions. | Why: The accountable person is already clear — you don't need a formal title to name them. |

**If you are adopting APEX:**

- Explicitly name who owns product/prioritization for each area
- That person is responsible for understanding the user, making hard calls about scope, and saying no
- They do not write stories (work items are self-service). They set direction.

---

## Adoption Path

If you are moving from Agile to APEX:

**Week 1-2: Eliminate the Obvious**
- Stop sprint planning. Stop estimating story points.
- Replace sprints with delivery cycles (whenever work actually ships)
- Run one lightweight review at a natural boundary

**Week 3: Restructure Work Items**
- Replace user stories with the work item format
- Make owner assignment non-negotiable

**Week 4: Rethink Synchronization**
- Evaluate whether daily standup is working. If not, replace with async
- Drastically shrink (or eliminate) sprint retrospectives

**Week 5+: Optimize**
- Use the [health-check.md](health-check.md) to identify where you have drifted or where process is still creating overhead
- Make targeted changes based on what you find

---

## What Stays the Same

**Version control.** Still good. Still use it.

**Code review.** Still good. Still do it. (APEX does not specify how, but owners usually do).

**Testing.** Still necessary. APEX does not comment on test strategy, but testing hard problems is still good.

**Documentation.** Still useful (though APEX values it less than Agile). Keep docs that matter. Delete docs that are out of date.

**Communication.** Still essential. APEX argues against *coordination-as-ceremony*, not against *shared context*.

---

## Common Fears About APEX

**"Won't we lose visibility into what the team is doing?"**

Not if you have named owners and a short priority list. You know: this person owns this thing. It ships on this date. That is more visibility than a 200-item backlog.

**"How do we ensure quality without acceptance criteria?"**

The owner ensures quality. They decide what "done" means. They own the output. The done condition is not less rigorous than acceptance criteria — it is just simpler to state.

**"What if people don't ship on time?"**

Same as Agile: you find out there is a problem. In APEX, you find out faster because you are not hiding it in a sprint. The owner either ships, or you have a conversation about what blocked them. That is more honest than a sprint where nominal commitments are made but not met.

**"Won't we end up with silos?"**

Only if owners do not communicate. APEX does not reduce the need for communication — it reduces the need for *synchronous ceremony*. Communication happens asynchronously or in small focused groups, not in all-hands meetings.

**"Can we still use Jira / Linear / GitHub Issues?"**

Yes. Use whatever tool you have. The structure matters (clear owner, clear scope), not the tool.

---

## The One-Month Test

Adopt APEX for one month. Use the [getting-started.md](getting-started.md) guide. At the end of the month, honestly answer:

- Are you making decisions faster?
- Did fewer meetings happen?
- Did less work get lost to coordination overhead?
- Is ownership clearer?

If the answer to 3+ of these is yes, APEX is probably a good fit for your team. If not, there is a structural reason. Identify it and decide whether to fix it or revert.

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
