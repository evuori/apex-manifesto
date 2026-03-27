# SpecDriven Development in the AI Era

## Why Specs Matter More Now

In traditional software development, specs were often overhead. Teams worked iteratively, adapted based on feedback, and hoped precision wasn't necessary. That trade-off made sense in a world where developers were expensive and iteration was unavoidable.

In the AI era, everything flips.

**With AI tools, precision is cheap. Rework is expensive.**

A skilled engineer with Claude can produce in hours what took days before. But that acceleration only works in one direction: toward the correct target. Without a precise specification of what "correct" means, you're accelerating toward the wrong destination faster.

Spec-driven development is not about documentation. It's about capturing **what the user actually needs, what success looks like, and what constraints matter**—before you write code. In APEX terms, it's operationalizing the "Precise" principle.

---

## What a Spec Is (And What It's Not)

### A spec is:
- **A contract between what we build and what solves the user's problem**. It answers: What is the user trying to do? What does correct behavior look like? What edge cases matter?
- **The input to AI-assisted development**. AI can write from a precise spec. It cannot guess a vague intent well.
- **The reference for decision-making**. When unexpected questions come up, you check the spec, not argue in a meeting.
- **Lightweight and specific**. Not a 200-page requirements document. A clear statement of scope, behavior, and success.

### A spec is not:
- A design document. It describes *what*, not *how*.
- A UI mockup (though it may reference one for clarity).
- A complete feature list. It describes a focused scope—one feature, one workflow, one piece of behavior.
- Implementation details. "Use Redis" is not a spec; "request latency < 200ms for reads" is.

---

## Spec Structure

A good spec has these parts. Adapt to context—a small feature might be 1-2 pages; a complex system might be longer.

### 1. **User & Problem**
Who is this for? What are they trying to do? Why does it matter?

*Example:*
> A data analyst needs to generate monthly reports from log data without waiting for engineering. They know SQL but not Python. Currently it takes a week and requires a request.

### 2. **Scope**
What are we building? What is explicitly out of scope?

*Example:*
> We're building a CLI tool that lets analysts write SQL templates, run them against the data warehouse, and export results to CSV. We're not building a UI, we're not handling real-time queries, we're not managing user permissions (assume they have warehouse access).

### 3. **Core Behavior**
How does the user interact with this? What are the key workflows?

*Example:*
> 1. User runs `report-gen init` to create a new report template
> 2. User edits the SQL file to define the query
> 3. User runs `report-gen run --date 2026-03` to execute for a specific month
> 4. Tool queries the warehouse, formats results, exports to CSV

### 4. **Success Criteria**
How do we know this works? What must be true when it's done?

*Example:*
> - A new analyst can create and run their first report in < 15 minutes with only CLI --help
> - Reports execute in < 30 seconds for typical queries
> - CSV output is importable directly into Excel without cleanup
> - Errors show clear, actionable messages (not stack traces)

### 5. **Constraints & Assumptions**
What are the hard boundaries? What are we assuming to be true?

*Example:*
> - Runs on analysts' laptops (not a server)
> - Python 3.10+
> - Warehouse credentials already configured locally
> - Reports under 100k rows (no pagination)
> - Analysts have SELECT access, no DROP/ALTER

### 6. **Edge Cases & Anti-Patterns**
What should we explicitly handle, and what should we explicitly not handle?

*Example:*
> - Handle: Empty result sets (report success, show 0 rows)
> - Handle: Malformed SQL (show error, don't execute)
> - Don't handle: Network timeouts (fail fast, let user retry)
> - Don't handle: Very large result sets (assume users write efficient queries)

### 7. **Open Questions**
Questions you're not sure about. Bring them to the user before building.

*Example:*
> - Should users be able to schedule recurring reports, or is manual runs enough?
> - What's the preferred error message format for non-technical analysts?
> - Do we need audit logging of queries run?

---

## The SpecDriven Workflow

Spec-driven development follows this cycle:

### 1. **Write the Spec (with the user)**
You and the person who cares about the outcome (user, product owner, stakeholder) align on the spec. This conversation often clarifies thinking that seemed clear before it was written.

- Start with problem and scope
- Get feedback on success criteria
- Identify unknowns explicitly
- Document constraints together

**Time spent here prevents days of rework later.**

### 2. **Review & Approve**
The person who will use this confirms the spec describes what they need. This is their chance to say "wait, we forgot X" or "actually we don't need Y."

### 3. **Implement from the Spec**
With a clear spec, you can:
- Use AI tools effectively (give them the spec and ask for implementation)
- Make decisions faster (when questions come up, check the spec)
- Know when to stop (when all success criteria are met, you're done)
- Validate your work (do the success criteria pass?)

### 4. **Validate Against the Spec**
Does the implementation meet the success criteria? Does it handle the documented edge cases? Are there open questions that need resolving?

### 5. **Update the Spec if Reality Changed**
Sometimes you discover the spec was incomplete or reality forced a change. Document what changed and why in the decision log. Update the spec to match.

---

## Writing Specs That AI Can Work From

When you're using AI tools, the spec becomes the primary input. A good spec for AI-assisted development has these qualities:

### Be Specific About Behavior
❌ Bad: "The system should be fast"
✅ Good: "List queries return results in < 200ms even with 1M records"

❌ Bad: "Users should find what they need easily"
✅ Good: "A new user can find the pricing page within 3 clicks from any main page"

### Describe the Happy Path Clearly
AI works best when the core flow is unambiguous. Edge cases are fine—but make the main path crystal clear.

❌ Bad: "Handle various input formats"
✅ Good: "Accept CSV or JSON input. CSV has headers in row 1, columns are: name, email, created_at. JSON is an array of objects with those fields."

### Include Examples
When you can show, not just tell, AI understands better.

❌ Bad: "The API should accept a query and return results"
✅ Good:
```
GET /api/search?q=python&limit=10
{
  "results": [
    {"id": 1, "title": "...", "relevance": 0.95}
  ]
}
```

### Name the Constraints
AI will optimize for whatever you optimize for. If you care about security, cost, latency, or compatibility—say so.

❌ Bad: "Build a data processor"
✅ Good: "Build a data processor that runs on under 512MB RAM, can process 1000 rows/second, and handles 99% uptime."

### List Your Non-Requirements Explicitly
What *aren't* you doing matters. It helps AI know where to stop adding features.

✅ Good: "This is a single-user CLI tool. Don't add multi-user features, API endpoints, or a database."

---

## Spec Scope: When to Spec What

### Micro-Feature (few hours to a day)
**Scope:** One specific behavior or fix
**Spec length:** 1-2 pages
**Example:** "Add a `--dry-run` flag to the deploy command that shows what would change without applying it"

### Feature (few days)
**Scope:** One user workflow with multiple components
**Spec length:** 2-4 pages
**Example:** "Build a report scheduling system so users can set queries to run weekly and email results"

### System (1-2 weeks or more)
**Scope:** Multiple features, integration points, architectural decisions
**Spec length:** 5-10 pages (or broken into multiple specs for each major component)
**Example:** "Build a real-time analytics dashboard with live data refresh, custom widgets, and team sharing"

### Rule of thumb:
If a single engineer could implement it in one sprint, it fits in one spec. If it requires multiple engineers or multiple sprints, break it into smaller specs.

---

## Common Patterns

### The Happy Path + Edge Cases Pattern
```
**Core Workflow:**
1. User provides input
2. System validates
3. System processes
4. System returns result

**Edge Cases:**
- Empty input → Error: "Input required"
- Invalid format → Error: "Expected format: ..."
- Timeout → Error: "Request took too long, try again"
```

### The API Contract Pattern
```
**Endpoint:** POST /api/users

**Request Body:**
```json
{
  "name": "string, required, 1-100 chars",
  "email": "string, required, valid email",
  "role": "enum: admin | user | viewer"
}
```

**Response Success (200):**
```json
{
  "id": "uuid",
  "created_at": "ISO 8601 timestamp"
}
```

**Response Errors:**
- 400: Validation failed (return field name and reason)
- 409: Email already exists
- 500: Unexpected error
```

### The Performance Pattern
```
**Behavior:** [What it does]

**Performance Requirements:**
- Typical case: < X ms
- Worst case: < Y ms
- Scales to: N items/users/requests
- Storage: < Z MB
```

---

## When to Call a Spec Done

A spec is done when:
- ✅ The user confirms it describes what they need
- ✅ Success criteria are clear enough to know when you're done
- ✅ Edge cases are identified (even if not all handled)
- ✅ Any open questions are documented
- ✅ Constraints are explicit

It's *not* done when you feel like you've covered everything—perfection is the enemy of progress. Done is when both parties can move forward confidently.

---

## Spec + APEX

Specs are how APEX operationalizes **Precision**:

- **Accountable:** The spec clarifies who decides what and when. The owner approves the spec before work starts.
- **Precise:** The spec *is* precision. It answers the hardest question early: "What does done look like?"
- **Expert-led:** Experts write the spec in conversation with users. AI implements from the spec.
- **eXpedient:** Specs prevent meetings about "wait, did you mean...?" and rework from misalignment.

---

## Next Steps

1. Use [templates/project-manifesto.md](../templates/project-manifesto.md) at project kickoff to align on user, scope, and requirements
2. Write a spec for your first feature before implementation
3. Reference [guides/project-structure.md](project-structure.md) to organize your specs alongside code
4. Use [templates/decision-log.md](../templates/decision-log.md) to document when specs change and why
