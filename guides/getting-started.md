# Getting Started with APEX
### A Week-Zero Playbook for Adopting APEX

---

## Why This Document

You've read the manifesto. APEX makes sense. But "how do I actually start?" is a different question. This guide walks through concrete first steps: what to stop doing immediately, what to set up, what conversations to have, and how to know if it's taking hold.

This is not a lengthy transition plan. APEX adoption should itself be expedient. Week one is about disrupting the wrong patterns. The rest is reinforcement.

---

## Week Zero: Four Actions

### 1. Name Owners

Stop having team members without explicit responsibility. In a meeting with your team:

**Immediately do this:**
- List every piece of in-flight work (projects, features, investigations, initiatives)
- Assign a single owner to each
- Write it down and share it publicly
- Make it clear: the owner is not a point of contact or a coordinator. They make decisions. They own the outcome.

**What this does:**
- Shifts accountability from "the team" to individuals
- Removes the invisible coordination overhead of consensus
- Creates clarity about who to escalate to and why

**For a 2-person team:** Your owner assignments might be trivial (you own the backend, I own the frontend). The point is not complexity — it is naming it explicitly.

**For a 10-person team:** You might have 8-10 owners for different systems, features, or initiatives. The number is less important than clarity: every person knows what they own and who owns what they touch.

---

### 2. Fill Out a Project Manifesto

Take the [project-manifesto.md](../templates/project-manifesto.md) template and complete it for each significant initiative currently in flight. You do not need to do this for every task — only for projects large enough to justify documentation. If something fits in a couple of paragraphs, that's probably too small to document this way.

**In a meeting, collectively:**
- **Define your user:** Not a job title. The actual person. What does their day look like? What are they trying to accomplish?
- **Name the core:** What two to four things must work exceptionally well? What would you cut everything else to protect?
- **State constraints:** Hard deadlines, fixed technical requirements, decisions already made that are not negotiable
- **List risks:** What are you assuming that could break? What do you not yet understand?

**Why this matters:**
- Stops you from building for assumptions instead of understanding
- Creates a shared record that prevents scope creep from relitigating resolved questions
- Takes an hour. It should take an hour. If it takes a day, you are not yet ready to start.

**What happens to it:**
The project manifesto lives in your repo under `docs/project-manifesto.md`. Not in a shared drive, not in Notion — in the repo, where it is version-controlled, PR-reviewable, and readable by AI tools alongside the code. The team reads it before asking questions about scope. It changes only when scope actually changes. Update the change log when you do.

See [guides/repo-as-os.md](repo-as-os.md) for why everything that informs the work lives in the repo.

---

### 3. Identify What to Stop

Most teams adopting APEX are coming from Agile or similar. Identify specific ceremonies that exist without clear purpose, and stop them.

**Common candidates to eliminate immediately:**
- **Velocity tracking:** You are not measuring progress by story points. You are measuring it by delivery. Stop recording and planning to velocity.
- **Backlog refinement:** If you have a 200-item backlog, the problem is not that it is not refined — it is that it is 200 items. You do not need a ceremony to make that worse. Delete the meeting. Go fix your prioritization.
- **Sprint planning:** If sprints were a useful forcing function for your team's work cycle, keep them. If you are doing two-week sprints because the calendar demands it, stop. Plan to delivery, not to a sprint boundary.
- **Standups (maybe):** If standups are working hard to create artificial busyness reports, replace them with a lightweight async standup (one sentence per person, posted to a channel, takes 10 seconds to read). If standups are genuinely enabling coordination, keep them but make them fast — 5 minutes maximum, one person talking at a time.
- **Retrospectives focused on process:** If your retrospectives spend 30 minutes discussing how the retrospective should work, you have a process problem, not a retrospective problem. Replace a lengthy retro with a 30-minute "what did we learn and what changes," run by the owner of the work, invite whoever cares.

**For teams smaller than 5 people:** You probably don't have many ceremonies to cut. The overhead was not yet visible. Do not create new ones.

**For teams of 10+:** You will probably find 2-4 meetings that are pure overhead. Delete them this week.

---

### 4. Have the Ownership Conversation

Your team needs to understand what ownership means under APEX, because it is different from what "ownership" meant in their prior experience.

**Explicitly say:**
- Ownership means you make the call. You are not coordinating consensus. You listen to input and you decide.
- Ownership is not antisocial. You communicate. You share context. You ask for input. Then you decide.
- If you are wrong, you own it. You are responsible for the quality of the decision. Not for the outcome necessarily — outcomes have variables — but for the quality of the reasoning.
- Shared responsibility is responsibility that belongs to no one. We do not do that.
- The owner is the escalation point. If the owner and a stakeholder disagree, the owner decides — unless the stakeholder is the owner's owner, in which case they also own it.

**What to listen for:**
- Discomfort with decisiveness ("but what if we get it wrong?") — normalize being wrong and owning it as a success, not a failure
- Fear of isolation ("but I can't decide this alone") — reinforce that they are not alone in gathering input, only alone in making the call
- Confusion about who the owner is of shared systems — clarify: one person owns each system. When two systems interact, there is a protocol for how the owners resolve disagreements.

---

### 5. Set Up Git-Native Tracking

Cancel the project board setup. You do not need a separate tracking tool. Before the first PR is opened:

- Agree on outcome slugs for your current work cycle (e.g., `auth`, `reporting`, `onboarding`) — these become your branch prefixes and milestone names
- Add the standard label set to your GitHub repo (`status: in-progress`, `status: blocked`, `status: in-review`, plus type and outcome labels)
- Create one GitHub Milestone per shared outcome
- Define WIP limits per team and write them in `IMPLEMENTATION.md`

From that point, tracking is automatic. PRs opened against milestones, labeled by status, with one named owner — that is the board. See [guides/git-native-tracking.md](git-native-tracking.md) for the full setup.

---

## Week One: Reinforce and Observe

You do not need new actions. You need to notice whether the first four actions are taking hold:

- Are people naming themselves as the owner when they talk about work?
- Is the project manifesto being used to answer scope questions?
- Are the ceremonies you cut staying cut, or are new ones reappearing?
- Are owners making decisions, or are they convening meetings to make decisions?

If the answers are yes, momentum is real. If they are no, there is a specific block — probably one of:

1. **Ownership is still diffuse in people's minds.** Have the conversation again. More specifically. More directly.
2. **The ceremonies you cut were solving a real problem.** Replace them with something smaller. A decision log (see the template). A 10-minute async check-in. Something lightweight that actually addresses the problem without importing ceremony.
3. **People do not understand the project manifesto.** Rewrite the sections people keep asking questions about. If the language is unclear, that is a documentation failure, not a team failure.

---

## Weeks Two and Beyond: Iterations and Anti-Patterns

### Iteration Pattern

Every two weeks (or on a natural delivery boundary), have a 30-minute check-in:
- What happened last period?
- What do we need to adjust?
- Are there ownership or prioritization questions?

This is not a sprint review, not a retrospective, not a standup. It is a sync. One owner runs it. It is recorded (one paragraph summary) and posted. Done.

### Watch for These Anti-Patterns

**"We decided as a team"** — If a decision is being attributed to the team, it is not owned. Call it out. Who decided? If it was truly collective, name why that was necessary and who is accountable if it turns out to be wrong.

**Decision meetings expanding** — If a decision that should take 15 minutes is getting an hour, the owner does not understand the problem well enough. The owner's job is to decide, not to gather perfect information. Send the owner to investigate solo, then decide.

**New meetings appearing** — Meetings are the default way teams respond to problems. When you notice a new recurring meeting, ask: what problem does this solve? Why can it not be solved asynchronously? If the answer is "it is more efficient" or "it feels like better communication," delete it.

**Backlog growing** — If your backlog is consistently growing faster than it shrinks, the problem is not in delivery — it is in intake. The owner (usually a product person) needs to decide what actually matters and say no to what does not. You cannot have a functioning prioritization system when the backlog is everything.

**"That is the AI's decision"** — If you are seeing AI output accepted without the owner evaluating it, stop. That violates Expert-led. The owner must be capable of reviewing and judging the output independently. If they are not, they need to learn or the task needs to be different.

---

## Scaling Considerations

### For 2-3 Person Teams

You probably do not need a project manifesto for every task. Use it for significant work only. You probably do not need async standups. Sync twice a week and call it done. The biggest win for small teams is removing pretense about consensus — someone decides, everyone else can focus on their work.

### For 5-10 Person Teams

This is where APEX starts creating the most value. The overhead of consensus gets expensive at this size. Ownership is crisp. Project manifestos are used. Ceremonies drop dramatically. The risk: teams can inadvertently create silos. Explicitly create coordination points where different owners sync (weekly if there are dependencies, monthly if there are not).

### For 10+ Person Teams

At this size, you have enough complexity that APEX discipline becomes essential. Ownership is not negotiable — it is the only way to avoid decision paralysis. Project manifestos are written formally. The decision log becomes critical. You need explicit protocols for how owners of interdependent systems resolve disagreements. You probably need a lightweight weekly sync (30 minutes maximum) organized by owner, where significant decisions and blockers surface.

---

## How to Know It Is Working

You will see these signals:

**Early (week one-two):**
- People are using the word "owner" naturally in conversation
- Questions about scope are being asked to the owner, not to the team
- The ceremonies you cut have not come back

**Medium term (week three-six):**
- Decisions are being made faster (because there is not a waiting period for consensus)
- Late-stage surprises decrease (because the project manifesto was actually used)
- Fewer meetings, shorter meetings
- Less email about "alignment" and more email about actual work

**Long term (week six+):**
- New team members are told "you own this" and understand what that means
- Ownership survives vacations and job transitions (it is a structural fact, not a person-dependent fact)
- The language of the team reflects APEX values (outcomes not activity, precision not completeness, etc.)

---

## Common Mistakes

**Mistake 1: Announcing APEX as an organizational change program.** You are not implementing a methodology. You are fixing a specific operating problem. Frame it that way: "We are going to make decisions faster by being clear about who owns what." That is a concrete problem statement. "We are adopting a new agile methodology" sounds like a process change and people resist it.

**Mistake 2: Expecting adoption to be smooth.** Teams that have spent years in consensus-based decision-making will find decisiveness uncomfortable at first. That is normal. The discomfort lasts a few weeks. Do not abandon the change because of it.

**Mistake 3: Treating ownership as a title.** Ownership is a responsibility structure, not an organizational chart. One person can own multiple things. Ownership can move week to week if the work changes. What matters is that it is named and clear.

**Mistake 4: Forgetting that APEX still requires communication.** APEX is not "decide in isolation." It is "decide without waiting for consensus." Owners communicate, solicit input, make themselves available. The difference is the owner makes the final call and owns the result.

---

## What to Do If It Is Not Working

If after three weeks you are seeing none of the "early signals" above, something is broken.

- **If people still say "we decided":** The ownership conversation did not land. Have it again, more directly. Find one concrete decision from the last week, name the owner, and reframe that decision explicitly: "[Owner] made this decision, consulted with [people], and committed to it."

- **If the ceremonies you cut came back:** Something they were solving for is now unsolved. Identify the actual problem (usually a coordination gap or a decision that needs to be made) and address it directly, not with a ceremony. A decision log often fixes this.

- **If decisions are slower, not faster:** The owner is trying to be too perfect before deciding. Coach them: what is the minimum you need to know before you can commit? Decide with that. Adjust if you are wrong.

- **If there is organizational resistance:** This usually means APEX conflicts with existing power structures (someone benefits from consensus-based slowness). Address it directly. If the existing structure is sacrosanct, APEX cannot work. Say that plainly.

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
