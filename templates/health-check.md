# APEX Health Check
### A Diagnostic Tool for Detecting Drift

---

## Purpose

APEX is a working methodology. Like all working systems, it erodes over time. Good intentions and stated values are not enough — teams slip back into ceremony, diffused accountability, and comfortable process without noticing.

This health check is not a grade or a maturity assessment. It is a diagnostic tool: a set of specific questions designed to surface where your team's actual practice has drifted from APEX principles. The output is not a score. It is a list of specific things to fix.

Run this every quarter or at natural project boundaries. One person (usually the technical lead or project owner) answers the questions. Discuss the "no" answers as a team and agree on what to change.

---

## How to Use This

For each question, answer Yes, No, or Uncertain.

- **Yes** means this is actively true in your team's current work.
- **No** means it is clearly not happening, or it has eroded.
- **Uncertain** means you are not sure — treat it as a "no" and investigate.

After answering all questions, the team discusses:
1. Which "no" answers are most serious?
2. What specifically needs to change?
3. Who is accountable for making that change happen?
4. How will you know it has changed?

Do not try to fix everything at once. Pick the 2-3 most serious drifts and address those.

---

## The Questions

### A — Accountable

These questions check whether individual ownership is real or whether accountability has diffused back into the team.

**1. Is there a named owner for every significant piece of in-flight work?**
- Yes = Everyone can say "I own X" or "Alice owns X" for every active project/feature/initiative
- No = There are pieces of work where responsibility is unclear or shared

**2. When a decision is made about your work, can you name the person who made it?**
- Yes = Decisions are attributed to individuals; you can trace "Sarah decided this" or "I decided this"
- No = Decisions get attributed to "the team" or "we decided" without a clear owner

**3. Do work items (tickets, tasks, stories) have a single assigned owner?**
- Yes = Every open ticket has one owner. Not a team. Not "TBD." One person.
- No = Many tickets are unassigned, team-assigned, or shared

**4. When something goes wrong, do you identify the owner and work forward?**
- Yes = Mistakes surface the owner's name; the conversation is about what they will do to fix it
- No = You search for distributed blame, or people hide mistakes to avoid responsibility

**5. Are owners actively making decisions without waiting for consensus?**
- Yes = You observe owners deciding, communicating the decision, and moving forward
- No = You observe owners convening meetings to "align" or get buy-in before deciding

**Interpretation:** If you answer "no" to 2 or more of these questions, accountability has eroded. Revisit the ownership conversation and reestablish that one person owns each piece of work.

---

### P — Precise

These questions check whether your team understands what it is building and for whom, or whether you are building from assumptions.

**6. Can your team articulate: who is the user, what is their core task, and what specific friction are you removing?**
- Yes = There is a shared understanding. People can say it the same way without reading from a document.
- No = Answers vary. People talk past each other about who the user is or what problem you are solving.

**7. Have you explicitly separated "must work exceptionally well" from "nice to have"?**
- Yes = You have a core (2-4 things) that you are protecting. Everything else is secondary.
- No = All features are treated as equally important. Your "core" is actually your full feature list.

**8. When a feature request comes in, is it evaluated against the core task or accepted because it is requested?**
- Yes = You actively ask "does this serve the core?" If not, it is deprioritized or rejected.
- No = Feature requests get added to the backlog. Deprioritization feels like something that happens to features, not something you actively choose.

**9. Has your team *investigated* the user's problem or just *accepted* their requested solution?**
- Yes = You have a record of talking to users, understanding their task, and designing based on understanding rather than acceptance of their requests.
- No = You designed to the specification as given. You have not probed the gap between what they asked for and what they actually need.

**10. When you look at your scope, are you confident that every piece serves the core or is explicitly deferred?**
- Yes = Your scope is tight. Everything in it connects to the core. Deferred items are named and committed to explicitly.
- No = Your scope has grown. You have items that seem important but are not core. Scope creep feels inevitable.

**Interpretation:** If you answer "no" to 2 or more of these questions, you have built on assumptions rather than understanding. This usually surfaces late as wasted effort. Invest in one cycle of genuine user investigation before continuing.

---

### E — Expert-led

These questions check whether experts are directing work and evaluating output, or whether AI (or anyone else) is being trusted without expert judgment.

**11. Before any AI-generated output goes forward, is it reviewed by someone qualified to evaluate it?**
- Yes = There is a person with domain knowledge checking AI output for correctness, appropriateness, and quality.
- No = AI output is accepted because it looks good or because the person is busy. No one is actively verifying.

**12. Can the person using AI tools to complete a task independently evaluate whether the output is correct?**
- Yes = They have sufficient domain knowledge to judge. They could do the task themselves if needed; they are using AI to accelerate.
- No = They could not independently evaluate the output. They are using AI to substitute for understanding.

**13. Are decisions about the product (what to build, how to design it, what matters) being made by people with domain knowledge?**
- Yes = The person making decisions understands the user, the problem, and the technical constraints.
- No = Decisions are being made by people without sufficient context or understanding.

**14. When you look at who is deciding about your work, are domain experts in the room or is the work being decided about in their absence?**
- Yes = The expert is present and has a voice in the decision.
- No = Decisions are made in other rooms by people without the relevant knowledge.

**15. Is "the AI produced it" ever an acceptable explanation for quality or correctness?**
- Yes = You hear this phrase. It is accepted as an excuse.
- No = When quality is questioned, the person who shipped it explains their reasoning, not the AI's.

**Interpretation:** If you answer "yes" to any of questions 11-15, expertise is being bypassed. This is not a sustainable way to work. The fixes: increase the expert's time on the task, reduce the scope, or reassign to someone with the necessary knowledge.

---

### X — eXpedient

These questions check whether you are moving fast without unnecessary ceremony, or whether ceremony has accumulated.

**16. How many recurring meetings does your team have? Can you justify each one in 30 seconds?**
- Yes = Few meetings (probably 3-5 weekly, could go down). Each has a clear purpose. If you cannot justify one, you discuss eliminating it.
- No = 10+ recurring meetings. Several are "for alignment," "for status," or "for sync." Meetings seem inevitable.

**17. Are decisions being made at the lowest level possible (by the person closest to the work)?**
- Yes = Most decisions are made by the owner. Escalation is rare and necessary.
- No = Many decisions are made in leadership meetings or require consensus. Owners are waiting for permission.

**18. Is there a short feedback loop between discovering a problem and fixing it?**
- Yes = When something is wrong, the owner can usually decide to fix it and move. Fixes ship fast.
- No = Problems get logged and triaged. Fixes compete for resources. Waiting periods are long.

**19. Are you measuring progress primarily by delivery (things that shipped) rather than activity (things in progress)?**
- Yes = The primary metric is "what shipped this week/month." You do not track velocity or burn-down.
- No = You measure progress by velocity, story points completed, tickets closed. Delivery is one of many metrics.

**20. When you look at your process, could you describe it in 5 minutes to someone new?**
- Yes = Your process is simple enough to explain quickly. It is not a complex set of rules and gates.
- No = Your process has special cases, exceptions, and evolved rules. Explaining it takes a long document.

**Interpretation:** If you answer "no" to 2 or more of these questions, ceremony has accumulated. Pick one thing — usually "reduce meetings" or "move decision-making down" — and focus on it for 2-3 weeks. Watch it unclog your system.

---

## Scoring and Discussion

There is no total score. The health check is not a grade.

Instead, organize your answers by principle:

**Accountable (Qs 1-5):** How many nos?
**Precise (Qs 6-10):** How many nos?
**Expert-led (Qs 11-15):** How many nos?
**eXpedient (Qs 16-20):** How many nos?

The principle(s) where you have the most "nos" are where you have drifted. Start there.

---

## What Comes Next

**Step 1:** As a team, discuss the 2-3 most serious "no" answers.

**Step 2:** For each, ask: what specifically would need to change for this to become a "yes"?

**Step 3:** Assign an owner (one person) to drive that change for the next 2-3 weeks.

**Step 4:** Set a check-in date (2 weeks out) to see if it has actually changed.

**Step 5:** Run this health check again in 3 months.

---

## Example Results

**Example 1:** A team answers "no" to Accountable questions 2 and 3. Decisions are fuzzy. Tickets are unassigned.
- **What this means:** Accountability has diffused. People are waiting for permission instead of owning decisions.
- **What to do:** Run the ownership conversation again. Make assigning an owner to every ticket non-negotiable. For the next 2 weeks, the tech lead does a 5-minute daily check: "Is every ticket assigned?" If not, assign it.

**Example 2:** A team answers "no" to Precise questions 8 and 10. They are not investigating user problems, just accepting requests. Scope has grown.
- **What this means:** You are building to specification without understanding. You will ship the wrong thing.
- **What to do:** Pause feature work. Spend one week investigating the core user task. What are they actually trying to accomplish? What is the friction they are experiencing? Design from understanding, not acceptance.

**Example 3:** A team answers "no" to eXpedient questions 16 and 17. Too many meetings. Decisions are centralized.
- **What this means:** Ceremony has accumulated. The owner is slowed by layers of approval.
- **What to do:** List all recurring meetings. For each one, ask: what would break if we stopped this? If the answer is "nothing serious," kill it. For decision-making, identify which decisions require approval and which the owner can make alone. Push authority down.

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
