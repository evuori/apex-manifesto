# APEX Work Item Format

---

## Purpose

Under APEX, a work item (ticket, task, issue, or whatever your system calls it) should capture the essential information needed for the owner to execute and for others to understand the work. Nothing more.

This is not a story template with acceptance criteria, story points, and sub-tasks. It is a minimal format that enforces APEX principles at the task level.

---

## The Template

Copy this format for each work item. Use it whether you are tracking work in GitHub Issues, Jira, a Notion database, or a simple Markdown file.

```
## [Title of Work Item]

**Owner:** [Single name]

**Problem:** [One sentence describing what is broken, missing, or blocked]

**Scope:**
- What this includes: [Brief description of what this work covers]
- What this explicitly does not include: [Things that might be assumed but are out of bounds]

**Done Condition:** [How the owner will know this is complete. Usually one sentence.]

**Dependencies:** [What this is blocked on, if anything. "None" if there are no blockers.]
```

---

## Filling It Out

### Title
One phrase describing the work. Examples:
- "Implement password reset flow"
- "Fix checkout performance regression"
- "Investigate database indexing strategy"
- "Write user research report"

Do not write: "Implement password reset flow and update documentation and write tests" — that is multiple work items.

### Owner
One name. The person responsible for this work. Not a team. Not "TBD." One person who can be held accountable.

If you do not know who owns this yet, do not create the work item. Assign ownership first.

### Problem
One sentence. What is actually broken, missing, or needed? Why does this work exist?

Examples:
- "Users cannot reset their password if they forget it, losing access to their account"
- "Checkout page takes 8 seconds to load on 4G networks; competitors load in 2 seconds"
- "We don't know whether our current schema design will support projected user growth"
- "Product team needs evidence about whether users actually want feature X before we build it"

Not: "Implement password reset" — that is the solution, not the problem. Describe the actual friction.

### Scope

Two parts: what is in, and what is explicitly out.

**What this includes:**
- "Implement reset flow from login page through email confirmation to password change"
- "Covers web interface only (mobile app is separate)"
- "Includes basic error handling and email retry logic"

**What this explicitly does not include:**
- "Does not include SMS-based reset (email only for v1)"
- "Does not include audit logging (can add post-launch)"
- "Does not redesign the login page"

The second part is important. It prevents scope creep by naming what you are *not* doing.

### Done Condition

How will you know this is finished? Usually one sentence.

Examples:
- "User can reset their password, receive confirmation email, and regain access"
- "Checkout loads in under 2 seconds on 4G; performance test passes"
- "Report includes analysis of schema, projections, and recommendation; delivered to engineering leadership"
- "Research complete with 8 interviews conducted; findings documented and presented to product team"

The done condition should be observable and achievable. Not: "password reset is working well" (too vague). Better: "user can reset password and log back in successfully; tested on Chrome and Safari."

### Dependencies

What is blocking this work? What does it depend on?

- "None" if there are no blockers
- If it depends on another work item: "Blocked on [other item owner/title]"
- If it is blocked by a decision: "Blocked on [decision and who should make it]"
- If it needs information from someone: "Blocked on feedback from design team"

If there are dependencies, note the owner of the dependent work. "Blocked on [thing]" is not enough information. "Blocked on password reset database schema (Alice owns it)" is actionable.

---

## Examples

### Example 1: Feature Work

```
## Implement password reset flow

**Owner:** Marcus

**Problem:** Users who forget their password cannot recover access to their account, resulting in support tickets and lost user accounts.

**Scope:**
- What this includes: Implement password reset link flow (request form → email with token → password change). Works on web. Includes email sending, token validation, and basic error messages.
- What this explicitly does not include: SMS-based reset, password strength requirements, audit logging, or mobile app support (separate feature).

**Done Condition:** User can reset forgotten password via email link; flow tested end-to-end in staging and production databases.

**Dependencies:** None
```

### Example 2: Performance Work

```
## Optimize checkout page load time

**Owner:** Jae

**Problem:** Checkout page loads in 8 seconds on 4G networks, significantly slower than competitors (2 seconds typical). This causes cart abandonment and poor perceived performance.

**Scope:**
- What this includes: Profile page load performance. Optimize images, code splitting, and server response time. Target <2 seconds on 4G.
- What this explicitly does not include: Redesign of checkout flow or steps, new payment methods, or analytics instrumentation.

**Done Condition:** Checkout page loads in <2 seconds on 4G network (measured via Lighthouse and real-world testing); performance benchmarks documented.

**Dependencies:** Blocked on CDN configuration change (Alex owns that).
```

### Example 3: Investigation Work

```
## Research schema design for projected growth

**Owner:** Sam

**Problem:** We don't know if our current database schema will support the projected user growth over the next 18 months. We risk outgrowing our design mid-launch, requiring emergency refactoring.

**Scope:**
- What this includes: Analyze current schema. Project queries and data sizes at 3x, 5x, and 10x current user count. Identify bottlenecks. Recommend design changes if necessary.
- What this explicitly does not include: Implement recommended changes, load testing, or benchmarking alternate database systems.

**Done Condition:** Report delivered to engineering leadership with analysis of current capacity, growth projections, identified risks, and specific recommendations (with trade-offs noted).

**Dependencies:** None
```

### Example 4: Research Work

```
## User research: demand for bulk export

**Owner:** Teri

**Problem:** Product team is unsure whether users actually need bulk export functionality. We are considering adding it to v1, but it adds complexity. We need evidence of actual demand before committing to scope.

**Scope:**
- What this includes: Conduct 8-10 user interviews (30 min each) asking about export needs, frequency, and current workflows.
- What this explicitly does not include: Building a prototype, analyzing feature requests from support, or conducting a survey.

**Done Condition:** Interviews completed and analyzed. Report includes theme summary, evidence for/against including bulk export, and recommendation to product team.

**Dependencies:** None
```

---

## What This Format Is Not

**Not a story template.** No acceptance criteria lists, no story points, no child tasks.

**Not a project.** If the work is more than one owner should reasonably accomplish in a couple of weeks, it is actually a project. Fill out the [project-manifesto.md](project-manifesto.md) template instead.

**Not a task list.** If you are breaking a work item into sub-tasks, the item is too large. Rethink the scope.

**Not a specification.** The problem and scope define what matters. The owner figures out the *how*.

---

## Using This Format

### At Intake
When someone proposes work, they fill out Owner, Problem, and Scope. If they cannot fill these out clearly, the work is not yet ready. Say no or ask them to clarify.

### Before Starting
The owner reviews the work item and confirms they understand it. If the problem or scope is unclear, they clarify with whoever proposed it. They confirm the dependencies are accurate.

### During Work
The owner may discover that scope was wrong or that dependencies were underestimated. They update the work item and communicate the change. That is normal. Document it.

### At Completion
The owner confirms the done condition is met. If it is, they close the work item. If it is not, they identify what is still blocked and create a new work item for the remainder (with a new owner and clear scope).

---

## Common Mistakes

**Mistake 1: Vague problem statement.**
- Wrong: "Improve user experience"
- Right: "Users report confusion when selecting payment method; they frequently select the wrong option and have to restart"

**Mistake 2: Scope so large it needs decomposition.**
- Wrong: "Build payment system" (that is a project)
- Right: "Implement Stripe integration for credit card payments" (that is a work item)

**Mistake 3: Done condition that is unmeasurable.**
- Wrong: "Password reset is working well"
- Right: "User can reset password via email; tested in staging and production; no errors in logs"

**Mistake 4: Unassigned work.**
- Wrong: "Owner: TBD"
- Right: Don't create the work item until you know who owns it.

**Mistake 5: Hidden dependencies.**
- Wrong: "Dependencies: None" (but actually blocked on someone's decision)
- Right: "Dependencies: Blocked on API design decision (Ahmed should make this call)"

---

## Scaling to Your System

### GitHub Issues
Use this template in your GitHub issue template file. Required fields: Owner, Problem, Scope, Done Condition, Dependencies.

### Jira
Create a custom issue view or template with these five required fields. Keep other fields minimal.

### Notion / Spreadsheet
Create a database with columns: Owner, Problem, Scope, Done, Dependencies. Make Owner a required field. Scope can be a longer text field.

### Markdown files
Create a simple document in your repo listing current work items in this format. Low overhead, version-controlled, searchable.

Whatever system you use, the goal is the same: one owner, clear problem, bounded scope, observable done condition.

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
