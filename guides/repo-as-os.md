# The Repository as the Team's Operating System

## The Problem With Five Tools

A typical modern team runs on five or more tools simultaneously. Design intent lives in Figma. Planning lives in Notion. Task status lives in Jira. Decisions made in meetings get buried in Slack threads. Code lives in GitHub.

None of these tools talk to each other in any meaningful way. The result is a team that spends significant energy just *locating* context — and frequently can't find it at all.

This fragmentation has always been expensive. In the AI era, it becomes a fundamental blocker.

---

## Why This Matters Differently Now

When an AI coding tool helps a developer build something, it can only see what's in the repository. It sees the code. It does not see the Figma file that explains why the UI works the way it does. It does not see the Notion doc where the product decision was made. It does not see the Slack thread where the team decided to abandon approach A in favor of approach B.

This is why AI assistance often feels shallow. The AI is working with a slice of the picture. It can write code, but it cannot reason about intent — because intent isn't in the repo.

**The fix is simple: put everything in the repo.**

Not as a process burden. As a deliberate architectural decision about where the team's knowledge lives.

---

## The Repo-as-OS Principle

In APEX, the repository is the team's operating system. Every artifact that informs the work lives there:

- **Specs** — what we're building and why (see [spec-driven-development.md](spec-driven-development.md))
- **Design rationale** — the thinking behind design decisions, not just the pixels
- **Research** — findings, competitive notes, user feedback that shaped decisions
- **Architecture decisions** — why the system is built the way it is (see [templates/decision-log.md](../templates/decision-log.md))
- **Planning** — outcomes, priorities, and the reasoning behind them
- **Status** — what's done, what's in progress, what's blocked

Code is one artifact among many. It is not the only thing that matters, and it is not the only thing that should be versioned.

---

## What This Is Not

This is not a call to abandon design tools or stop using specialized software. Designers should use design tools. The point is not where you *create* the artifact — it is where you *store the reasoning* behind it.

A Figma file exported to `/design/feature-name/` tells the AI what the screen looks like. A `rationale.md` in the same folder tells the AI *why* it looks that way, what constraints existed, what alternatives were rejected. The Figma file without the rationale is half the picture.

Similarly, this is not about creating documentation for its own sake. Every artifact in the repo should exist because it informs the work — either for the humans doing it now or the AI assisting them or the team members joining later.

---

## Connecting to Spec-Driven Development

APEX already holds that specs live in the repo before code is written. The Repo-as-OS principle extends this naturally:

If the spec lives in the repo, the design rationale that informed the spec should too. If the design rationale lives in the repo, the research that shaped the design should too. If the research lives in the repo, the decisions about what to build and why should too.

**The repo becomes the complete record of how the team thinks, not just what the team ships.**

This is not overhead. This is what prevents the same conversations from happening three times, what prevents new team members from repeating mistakes that were already made, and what allows AI tools to give genuinely useful, contextually grounded assistance.

---

## What Belongs in the Repo

### Always
- Code and tests
- Specs (before and during implementation)
- Architecture and technical decisions
- Design rationale (the *why* behind design choices)
- `IMPLEMENTATION.md` — current status, in-progress work, known limitations

### When It Informed a Decision
- Research findings that shaped scope or approach
- Competitive analysis that influenced design choices
- User feedback that changed direction
- Meeting outcomes where significant decisions were made

### What Doesn't Belong
- Raw working files that have no explanatory value (100-layer Figma source files, unless the team reviews them in PRs)
- Communications that didn't result in a decision
- Duplicates of information already captured elsewhere in the repo

The test: *could a new team member or an AI assistant use this to understand why something was built the way it was?* If yes, it belongs. If it's just noise, it doesn't.

---

## The PR-as-Review Workflow

When everything lives in the repo, the pull request becomes the team's universal review mechanism — not just for code.

A designer opening a PR with updated design files and a `design/rationale.md` explaining the changes goes through the same review workflow as an engineer opening a PR with new code. The discussion happens in the PR. The approval is explicit. The merge timestamp is the decision record.

This has a significant benefit: it eliminates the distinction between "code review" and "design review" and "planning review." There is one workflow. Everyone uses it. Nothing gets siloed.

---

## Suggested Repo Structure Extension

This extends the structure defined in [project-structure.md](project-structure.md):

```
project-name/
├── docs/
│   ├── project-manifesto.md
│   ├── specs/
│   ├── decisions/
│   │   └── decision-log.md
│   ├── research/              # Findings that informed decisions
│   │   └── README.md          # Index of research, what each influenced
│   └── planning/              # Outcome briefs, roadmap items
│       └── README.md
├── design/
│   ├── README.md              # Overview of design approach and current state
│   ├── [feature-name]/
│   │   ├── rationale.md       # Why it looks/works this way (use design-brief template)
│   │   ├── assets/            # Exported SVGs, PNGs, design tokens
│   │   └── iterations/        # Prior versions if they explain the evolution
│   └── system/
│       ├── colors.md          # Design system decisions
│       ├── typography.md
│       └── components.md
├── src/
├── IMPLEMENTATION.md
└── README.md
```

---

## What the AI Gets

When this structure is followed, an AI assistant working on the codebase can:

- Read the spec to understand what it's supposed to build
- Read the design brief to understand why the UI is structured the way it is
- Read the decision log to understand why the architecture was chosen
- Read the research to understand what user problems motivated the work
- Read `IMPLEMENTATION.md` to understand what's complete and what's in progress

This is not theoretical. The difference between an AI that has access to this full context and an AI that only sees the code is the difference between a collaborator and a syntax completer.

---

## For Multi-Team Environments

When multiple small teams work toward a shared goal, the repo-as-OS principle scales naturally. Each team's work lives in its own area of the repo, but shared context — architecture decisions, cross-team specs, research that affects multiple workstreams — lives at the root level in `docs/`.

This means any team can see what any other team is building and why, without asking, without a meeting, without waiting for someone to update a project board. The repo *is* the shared context.

---

## The Discipline Required

This only works if the team is consistent about it. A design rationale that lives in someone's head is worth nothing to the AI or to the team member who joins three months later. A decision that got made on a call but never made it into the decision log is a decision that will get relitigated.

The habit to build: when something gets decided, it goes in the repo before the day ends. Not as a long document. A few sentences in the right file. The bar is "good enough that someone else can understand what was decided and why" — not "comprehensive enough to satisfy a documentation audit."

APEX teams treat the repo as the authoritative record. If it's not in the repo, it didn't happen.

---

## Next Steps

1. Read [guides/project-structure.md](project-structure.md) for the base repo structure APEX projects follow
2. Add `/design` and `/docs/research` and `/docs/planning` directories to your project
3. Use [templates/design-brief.md](../templates/design-brief.md) for design and planning artifacts
4. Review [guides/git-native-tracking.md](git-native-tracking.md) for how progress tracking works from this structure
5. When onboarding new team members or AI tools, point them at the repo root — not at five different tools

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
