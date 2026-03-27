# Decision Log Template

---

## Purpose

This template tracks significant decisions made during a project — what was decided, who decided it, when, the alternatives considered, and why one was chosen. It complements the [project-manifesto.md](project-manifesto.md), which captures decisions made *before* work begins. This log captures decisions made *during* delivery.

The distinction is important: the project manifesto prevents relitigating foundational questions. The decision log prevents relitigating delivery-phase trade-offs. Together, they create a record that stops scope creep and scope disputes.

---

## When to Log a Decision

**Log it if:** The decision could plausibly be reversed without trivial cost, someone was consulted who disagreed, or a significant alternative was ruled out.

**Do not log it if:** It is a routine implementation choice, the owner made it unilaterally without consultation, or it is clearly reversible with zero friction.

**The test:** If a decision dispute could surface three months from now, it warrants a log entry.

---

## The Log Template

Use this format. Copy it into a Markdown file (`DECISIONS.md` or `decision-log.md`) at the root of your project or in your shared documentation.

```markdown
# Decision Log

| Date | Owner | Decision | Alternatives Considered | Why This One |
|------|-------|----------|------------------------|--------------|
| | | | | |

```

---

## Filling It Out

**Date:** When the decision was made. Use YYYY-MM-DD format.

**Owner:** The person who made the decision. One name. Not "the team." Not "we." One owner.

**Decision:** What was decided, stated as simply as possible. One sentence. Examples:
- "We will use PostgreSQL instead of MySQL for the database"
- "We are deprioritizing the admin export feature to ship on March 15"
- "We will use Redis for caching, not in-memory"
- "We are refactoring the auth system now, not post-launch"

**Alternatives Considered:** What other options were seriously evaluated? If only one option was considered, say "none seriously considered." Examples:
- "MySQL (rejected: licensing concerns), SQLite (rejected: not suitable for scale)"
- "Firebase (rejected: vendor lock-in), DynamoDB (rejected: cost structure)"

**Why This One:** Why did the owner choose this option over the alternatives? One or two sentences. The reasoning should be clear enough that someone reading it six months from now understands the logic, even if they disagree.

---

## Example Log

```markdown
# Decision Log

| Date | Owner | Decision | Alternatives Considered | Why This One |
|------|-------|----------|------------------------|--------------|
| 2026-03-15 | Alex | Use Stripe for payments instead of implementing custom billing | PayPal, custom integration | Stripe's API is simpler, reduces scope by 2 weeks, fully compatible with our architecture. The small fee is worth the execution speed. |
| 2026-03-18 | Jordan | Ship v1 without the dashboard analytics feature | Include it, defer it to v1.1 | Analytics is nice-to-have. The core task (report generation) works without it. Shipping without it lets customers give feedback on the core. |
| 2026-03-22 | Casey | Revert to synchronous database queries during setup, not async | Keep async pattern throughout | Async added complexity during setup for no benefit (users are willing to wait 5 seconds for a one-time operation). Simplifies the code path. Will revisit if init time becomes a problem. |
| 2026-03-25 | Alex | Extend the deadline to March 30 | Keep March 28 deadline | Two critical bugs surfaced in testing that are simple to fix but need time. Better to ship clean than on time but broken. Stakeholders agreed this trade-off serves them better. |

```

---

## How to Use This Log

1. **Create it at kickoff.** Start the log when the project manifesto is filled out. It is a sibling document.

2. **Update it when decisions are made.** The owner logs it same day or next day. It should not be a burden — one minute to add an entry.

3. **Refer to it when scope questions arise.** When someone asks "should we include feature X?" and you said no during delivery, point them to the decision log entry. This is its main value.

4. **Use it during the final stretch.** Late-stage decisions are made under pressure. The log prevents them from being challenged or forgotten. "We logged this on March 22. Here is why. We are not relitigating."

5. **Keep it through post-launch.** Questions about "why did we build it this way?" are easiest answered by referring to the decision made during building. This is good for knowledge transfer and for planning the next version.

---

## What NOT to Log

- Routine implementation choices ("we used this design pattern," "we optimized this function," "we chose this variable name")
- Decisions that were made by explicit process (e.g., "we always use this database for this class of problem")
- Things that are obviously reversible ("we shipped the UI without animations")
- Decisions the owner made without consultation (unless the owner actively chose to consult and someone disagreed)

The goal is to capture *significant* decisions that someone might later question or that could have gone a different direction. Logging everything becomes clutter. Logging nothing defeats the purpose.

---

## Common Patterns

### Scope Trade-offs
These are the most important to log.

| Date | Owner | Decision | Alternatives Considered | Why This One |
|------|-------|----------|------------------------|--------------|
| 2026-03-20 | Morgan | Remove the bulk export feature from v1 | Scope delay to include it, cut something else | Bulk export would add 10 days. Core feature (one-off export) works fine. Export can ship in v1.1. Customers need the core now. |

### Technical Decisions
Log these when the alternatives had real trade-offs (speed vs. maintainability, flexibility vs. simplicity, etc.).

| Date | Owner | Decision | Alternatives Considered | Why This One |
|------|-------|----------|------------------------|--------------|
| 2026-03-12 | Taylor | Use a monolithic backend, not microservices | Microservices architecture | We have 3 engineers. Microservices would require more infrastructure overhead than execution overhead. Monolith lets us move faster. Revisit in 12 months if scale demands it. |

### Schedule Changes
Log changes to milestones or deadlines, especially if they were disputed.

| Date | Owner | Decision | Alternatives Considered | Why This One |
|------|-------|----------|------------------------|--------------|
| 2026-03-25 | Riley | Extend final testing to April 2 (slip 4 days) | Keep March 29 deadline | Quality gap in the purchase flow is unacceptable to launch with. Fix takes 4 days. Slipping is better than launching with a broken core feature. |

---

## Review the Log as Part of Your Health Check

Periodically (quarterly or at project close), scan the decision log and ask:
- Are decisions being logged or buried?
- Are controversial decisions being named and explained, or are they reappearing as disputes?
- Is the log being used to settle scope questions, or are the same debates happening repeatedly?

If the answer to any of these is "no," something is broken. Usually it is ownership — decisions are not being made clearly, or they are being made by committee and then the owner is afraid to enforce them.

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
