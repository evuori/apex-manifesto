# Report Generator CLI — Example APEX Project

This is a complete, worked example of an APEX project using SpecDriven development. It shows:

- How to structure a small project (CLI tool)
- What a project manifesto looks like
- How to write specs for features
- How to track decisions
- How specs connect to code

**This is a realistic example meant for learning.** You can study the artifacts, adapt them for your own project, or even fork this as a template.

---

## The Project

**Report Generator** is a CLI tool that lets analysts generate reports from log data without writing code or waiting for engineering.

- Users write SQL templates once
- Run reports on demand with date parameters
- Export results as CSV
- Deploy to any machine (single binary)

This is the kind of project that benefits from APEX:
- ✅ Small, focused scope
- ✅ User is clear (analysts)
- ✅ Success is measurable (time to first report)
- ✅ AI can implement well from a clear spec

---

## How to Use This Example

1. **Read the docs in order:**
   - [docs/project-manifesto.md](docs/project-manifesto.md) — Understand the user and problem
   - [docs/specs/core-features.md](docs/specs/core-features.md) — See what we're building
   - [docs/decisions/decision-log.md](docs/decisions/decision-log.md) — Understand why we made certain choices

2. **Look at the code:**
   - [src/cmd/](src/cmd/) — How commands are organized
   - Notice the comments linking back to specs
   - [IMPLEMENTATION.md](IMPLEMENTATION.md) — Current status

3. **Adapt for your project:**
   - Use [docs/project-manifesto.md](docs/project-manifesto.md) as a template for your own
   - Copy [docs/specs/core-features.md](docs/specs/core-features.md) and edit for your features
   - Start a decision-log.md as you make architectural choices

---

## Project Structure

```
cli-tool-example/
├── README.md                        # You are here
├── docs/
│   ├── project-manifesto.md        # User, problem, scope
│   ├── specs/
│   │   ├── README.md               # Index of specs
│   │   └── core-features.md        # Feature specifications
│   ├── decisions/
│   │   └── decision-log.md         # Why we chose Go, Docker, etc.
│   └── user-guide.md               # How to use the tool (for analysts)
├── src/
│   ├── cmd/
│   │   ├── init/
│   │   │   └── init.go             # Implements spec: init command
│   │   ├── run/
│   │   │   └── run.go              # Implements spec: run command
│   │   └── root.go
│   ├── config/
│   │   ├── loader.go
│   │   └── validator.go            # Implements spec: config validation
│   ├── main.go
│   └── version.go
├── tests/
│   ├── integration/
│   │   └── report_test.go          # Tests against actual runner
│   └── fixtures/
│       ├── sample-template.sql
│       └── sample-config.yaml
├── Dockerfile                       # Build as single binary
├── go.mod
├── go.sum
├── IMPLEMENTATION.md                # Current status and next priorities
└── .github/
    └── CODEOWNERS
```

---

## What Makes This an APEX Project

### Accountable
- **One owner per component:** Jake owns the SQL runner, Alex owns the CLI commands, Morgan owns the data warehouse connection
- **CODEOWNERS file:** Shows explicitly who decides about what
- **Decision log:** Shows who made each key decision

### Precise
- **Project manifesto:** User, problem, scope are crystal clear before development started
- **Specs for each feature:** Init command, run command, config file format — all specified before code
- **Success criteria:** "First report in < 15 minutes" — we know when we're done

### Expert-led
- **Analyst feedback incorporated early:** Spec was reviewed with actual users before code started
- **Technical decisions made by domain experts:** Data warehouse connection handled by person with warehouse expertise
- **AI assistance:** Specs were clear enough to use Claude for implementation

### eXpedient
- **Minimal process:** No standups, no weekly status meetings, no story points
- **Decision log instead of meetings:** Decisions are documented, not discussed in alignment sessions
- **Done means done:** When all success criteria pass, feature is shipped

---

## How to Read the Specs

Start with [docs/specs/README.md](docs/specs/README.md) which lists:
- What each spec covers
- Status (complete, in progress, planned)
- Who implemented it

Then read [docs/specs/core-features.md](docs/specs/core-features.md) to see:
- How the tool works (user workflows)
- What success looks like
- Edge cases we handle
- Configuration format

---

## Real Implementation Details

The code here is **real, working code** that:
- ✅ Actually parses YAML configs
- ✅ Actually runs SQL and handles errors
- ✅ Actually validates inputs against the spec
- ✅ Actually exports to CSV

You can study it to see how specs map to implementation. For example:
- The success criterion "First report in < 15 minutes" drove the choice to use templates (no SQL writing needed)
- The constraint "Runs on analysts' laptops" drove the choice to compile to a single binary
- The spec on config validation drove the error messages users see

---

## Using This as a Template

To use this as a starting point for your own project:

1. Copy the entire `cli-tool-example/` directory
2. Update [docs/project-manifesto.md](docs/project-manifesto.md) with your user, problem, and scope
3. Replace [docs/specs/core-features.md](docs/specs/core-features.md) with your actual features
4. Replace the code in `src/` with your implementation
5. Update [docs/decisions/decision-log.md](docs/decisions/decision-log.md) with your choices
6. Keep [IMPLEMENTATION.md](IMPLEMENTATION.md) updated as you work

---

## Questions?

Each artifact in this example has comments explaining what it is and why. Read them in this order:

1. [docs/project-manifesto.md](docs/project-manifesto.md) — 5 min read
2. [docs/specs/core-features.md](docs/specs/core-features.md) — 10 min read
3. [docs/decisions/decision-log.md](docs/decisions/decision-log.md) — 5 min read
4. [src/cmd/init/init.go](src/cmd/init/init.go) — See spec→code connection
5. [IMPLEMENTATION.md](IMPLEMENTATION.md) — Current status
